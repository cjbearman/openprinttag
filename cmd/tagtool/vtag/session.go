package vtag

import (
	"errors"
	"fmt"

	"github.com/ebfe/scard"
)

// A Session represents a communications session to a card applied to the reader
type Session struct {
	atr    string
	driver *acr1552
	card   *scard.Card
	status *scard.CardStatus
	uid    string
	tag    *Tag
}

// GetATR returns the ATR for the card associated with the session
func (s *Session) GetATR() string {
	return s.atr
}

// GetTag returns tag information for the card associated with the session
func (s *Session) GetTag() *Tag {
	return s.tag
}

// ReadSingleBlock reads a single block
func (s *Session) ReadSingleBlock(block uint16) ([]byte, error) {
	return s.driver.readSingleBlock(s.card, int(s.GetTag().BlockSize()), block)
}

// WriteSingleBlock writes a single block
func (s *Session) WriteSingleBlock(block uint16, data []byte) error {
	return s.driver.writeSingleBlock(s.card, int(s.GetTag().BlockSize()), block, data)
}

// Write will write the provided data at the specified byte address
func (s *Session) Write(start int, data []byte) error {
	bs := int(s.GetTag().BlockSize())
	if start%bs != 0 {
		return fmt.Errorf("start address must be a multiple of %d", bs)
	}

	// Pad data to exact block sizes
	for len(data)%bs != 0 {
		data = append(data, 0x00)
	}

	if start+len(data) > s.tag.GetAvailableBytes() {
		return errors.New("data exceeds tag user memory space")
	}

	nextBlock := uint16(start / bs)

	for i := 0; i < len(data); i += bs {
		err := s.WriteSingleBlock(nextBlock, data[i:i+bs])
		if err != nil {
			return fmt.Errorf("at block %d: %v", nextBlock, err)
		}
		nextBlock++
	}
	return nil
}

// Read reads length bytes from the specified start address byte
func (s *Session) Read(start int, length int) ([]byte, error) {
	bs := int(s.GetTag().BlockSize())

	data := make([]byte, 0)

	if start%bs != 0 {
		return nil, fmt.Errorf("start address must be a multiple of %d", bs)
	}
	if start+length > s.tag.GetAvailableBytes() {
		return nil, errors.New("data exceeds tag user memory space")
	}
	nextBlock := uint16(start / bs)
	for len(data) < length {
		blockData, err := s.ReadSingleBlock(nextBlock)
		if err != nil {
			return nil, fmt.Errorf("at block %d: %v", nextBlock, err)
		}
		data = append(data, blockData...)
		nextBlock++
	}
	return data, nil
}
