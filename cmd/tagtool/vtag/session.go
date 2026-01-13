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
	"github.com/ebfe/scard"
)

// A Session represents a communications session to a card applied to the reader
type Session struct {
	atr    string
	driver driver
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

// Write will write the provided data at the specified byte address
// start must be block aligned
// If data is not multiple of block size, it will be padded to block size with 0x00 bytes
func (s *Session) Write(start int, data []byte) error {
	return s.driver.write(s.card, uint16(start), data)
}

// Read reads length bytes from the specified start address byte
// No block alignment is required for reads
func (s *Session) Read(start int, length int) ([]byte, error) {
	return s.driver.read(s.card, uint16(start), uint16(length))
}
