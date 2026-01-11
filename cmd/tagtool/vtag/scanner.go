package vtag

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/ebfe/scard"
)

// Scanner implements the interface to the PCSC interface
type Scanner struct {
	ctx     *scard.Context
	readers []string
	ready   sync.WaitGroup
	close   sync.WaitGroup
	err     error
	isReady bool
}

// NewScanner creates a new scanner connection
func NewScanner() *Scanner {
	r := &Scanner{}
	r.ready.Add(1)
	r.close.Add(1)
	go r.run()
	return r
}

// OnCardFunc must be implemented by the caller of the OnCard callback
type OnCardFunc func(*Session) error

// run is a goroutine that maintains communication with the PCSC service
func (s *Scanner) run() {
	// Establish a context
	ctx, err := scard.EstablishContext()
	if err != nil {
		s.err = fmt.Errorf("failed to establish context: %w", err)
		s.ready.Done()
		return
	}
	defer ctx.Release()

	// List available readers
	readers, err := ctx.ListReaders()
	if err != nil {
		s.err = fmt.Errorf("failed to list readers: %w", err)
		s.ready.Done()
		return
	}
	if len(readers) < 1 {
		s.err = errors.New("no readers found")
		s.ready.Done()
		return
	}
	acceptedReaders := []string{}
	for _, rdr := range readers {
		if strings.Contains(rdr, "ACR1552") && strings.Contains(rdr, "PICC") {
			// This is the only reader we are supporting at this time
			acceptedReaders = append(acceptedReaders, rdr)
		}
		debug("found reader: %s", rdr)
	}

	s.ctx = ctx
	s.readers = acceptedReaders
	s.isReady = true
	s.ready.Done()
	debug("init complete")
	s.close.Wait()
	s.isReady = false
}

// Close must be called to close down our goroutine
func (s *Scanner) Close() {
	debug("closed")
	s.close.Done()
}

// waitForTag waits for a tag to be applied to the reader
func (r *Scanner) waitForTag() (string, error) {
	rs := make([]scard.ReaderState, len(r.readers))
	for i := range rs {
		rs[i].Reader = r.readers[i]
		rs[i].CurrentState = scard.StateUnaware
	}

	for {
		for i := range rs {
			if rs[i].EventState&scard.StatePresent != 0 {
				debug("tag found on reader %s", r.readers[i])
				return r.readers[i], nil
			}
			rs[i].CurrentState = rs[i].EventState
		}
		err := r.ctx.GetStatusChange(rs, -1)
		if err != nil {
			return "", err
		}
	}
}

// OnCard should be called to provide a callback function which will be called whenever a card
// is applied to the reader
func (s *Scanner) OnCard(do OnCardFunc) error {
	// Wait to ensure everything is initialized
	debug("waiting for ready")
	s.ready.Wait()
	debug("ready")
	if s.err != nil {
		return s.err
	}

	// Wait for a tag
	reader, err := s.waitForTag()
	if err != nil {
		return fmt.Errorf("waiting for tag: %w", err)
	}
	debug("tag found on reader %s", reader)

	driver := &acr1552{reader}
	if driver == nil {
		return errors.New("no scanner driver found")
	}

	// Establish a context with the reader
	debug("Connecting to reader %s", reader)
	card, err := s.ctx.Connect(reader, scard.ShareExclusive, scard.ProtocolAny)
	if err != nil {
		return fmt.Errorf("connect:%w", err)
	}
	defer card.Disconnect(scard.ResetCard)

	// We do all our communications in a transparent session
	// start it now and defer the ending of the session
	debug("start-transparent")
	driver.start_transparent(card)
	defer driver.end_transparent(card)

	// Set the appropriate protocol for the session
	driver.set_protocol_iso15693_3(card)

	// Grab the card status, and harvest the ATR
	status, err := card.Status()
	if err != nil {
		return fmt.Errorf("status:%w", err)
	}
	atr := hex.EncodeToString(status.Atr)
	debug("ATR: %s\n", atr)

	// Now we can fire up a session
	session := &Session{
		atr:    atr,
		driver: driver,
		card:   card,
		status: status,
	}

	// Try and ID the tag
	tag, err := driver.cardID(card)
	if err != nil {
		return fmt.Errorf("failed to find an acceptable card: %v", err)
	}
	session.tag = tag

	// and process the callback
	return do(session)
}
