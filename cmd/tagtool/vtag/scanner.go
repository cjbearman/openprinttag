// MIT License
//
// # Copyright (c) 2026 Christopher J Bearman
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
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

	commonDriver := &commonDriver{}
	var driver driver

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
	commonDriver.start_transparent(card)
	defer commonDriver.end_transparent(card)

	// Set the appropriate protocol for the session
	commonDriver.set_protocol_iso15693_3(card)

	// Grab the card status, and harvest the ATR
	status, err := card.Status()
	if err != nil {
		return fmt.Errorf("status:%w", err)
	}
	atr := hex.EncodeToString(status.Atr)
	debug("ATR: %s\n", atr)

	// Now we have to ID the card
	uid, err := commonDriver.getUID(card)
	if err != nil {
		return fmt.Errorf("get uid: %w", err)
	}

	var tag *Tag
	if uid[0] == 0xe0 && uid[1] == 0x04 && uid[2] == 0x01 {
		// First three bytes E0 04 01, then this is NXP ICODE
		driver = &icode{commonDriver: *commonDriver}
		driver.(*icode).commonDriver.driver = driver
		// Bits 37, 36 identify the tag type 00 (SLI), 10 (2/SLIX), 10 (1/SLIX2)
		// This are the following bits of the fourth byte:
		subtype := ((uid[3] & 0x18) >> 3)
		switch subtype {
		case 0x00:
			tag = &Tag{
				ttype: ICODE_SLI,
				uid:   uid,
			}
		case 0x01:
			tag = &Tag{
				ttype: ICODE_SLIX2,
				uid:   uid,
			}
		case 0x02:
			tag = &Tag{
				ttype: ICODE_SLIX,
				uid:   uid,
			}
		default:
			return errors.New("unsupported ICODE tag type")
		}
		driver.(*icode).tag = tag

	} else if uid[0] == 0xe0 && uid[1] == 0x02 && (uid[2] >= 0x24 && uid[2] <= 0x27) {
		// ST25DV...
		driver = &st25dvxx{commonDriver: *commonDriver}
		driver.(*st25dvxx).commonDriver.driver = driver

		// Use extended system get info to get nblocks
		_, nblocks, err := driver.getCardInformation(card)
		if err != nil {
			return fmt.Errorf("failed extended system get info on ST25DVxx: %v", err)
		}
		switch nblocks {
		case 128:
			tag = &Tag{
				ttype: ST25DV04,
				uid:   uid,
			}

		case 512:
			tag = &Tag{
				ttype: ST25DV16,
				uid:   uid,
			}
		case 2048:
			tag = &Tag{
				ttype: ST25DV64,
				uid:   uid,
			}
		default:
			return errors.New("unsupported ST25DVxx tag type")
		}
		driver.(*st25dvxx).tag = tag
	}

	// Now we can fire up a session
	session := &Session{
		atr:    atr,
		card:   card,
		status: status,
		driver: driver,
		tag:    tag,
	}

	// and process the callback
	return do(session)
}
