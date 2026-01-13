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
	"fmt"

	"github.com/ebfe/scard"
)

// st25dvxx implements specific commands for ST25DVxx tags
type st25dvxx struct {
	commonDriver
}

// getCardInformation retrieves the block size and number of blocks from the tag
func (d *st25dvxx) getCardInformation(card *scard.Card) (blocksize, nblocks uint16, err error) {
	data, err := d.direct(card, 0x02, 0x3b, 0x14)
	if err != nil {
		return 0, 0, fmt.Errorf("failed extended system get info: %v", err)
	}
	if len(data) < 12 {
		return 0, 0, fmt.Errorf("extended system get info short return: %v", err)
	}
	nBlocks := (int(data[11]) << 8) + int(data[10]) + 1
	return 4, uint16(nBlocks), nil
}

// readSingleBlock utilizes the extended read single block command to access all memory locations
func (d *st25dvxx) readSingleBlock(card *scard.Card, block uint16) ([]byte, error) {
	// Use extended read single block, to allow for addressing all memory locations
	data, err := d.direct(card, 0x02, 0x30, uint8(block&0xff), uint8((block>>8)&0xff))
	if err != nil {
		return nil, fmt.Errorf("failed extended single block read: %v", err)
	}
	if len(data) != d.tag.BlockSize()+1 {
		return nil, fmt.Errorf("unexpected extended single block read length: %d", len(data))
	}
	return data[1:], nil

}

// writeSingleBlock utilizes the extended write single block command to access all memory locations
func (d *st25dvxx) writeSingleBlock(card *scard.Card, block uint16, data []byte) error {
	// Use extended write single block, to allow for addressing all memory locations
	if len(data) != d.tag.BlockSize() {
		return fmt.Errorf("data length must be %d bytes, got %d", d.tag.BlockSize(), len(data))
	}

	cmd := append([]byte{0x02, 0x31, uint8(block & 0xff), uint8((block >> 8) & 0xff)}, data...)

	data, err := d.direct(card, cmd...)
	if err != nil {
		return fmt.Errorf("failed extended single block write: %v", err)
	}

	return nil
}

// readMultipleBlocks implements the extended read multiple blocks method to access all memory locations
func (d *st25dvxx) readMultipleBlocks(card *scard.Card, startBlock, nBlocks uint16) ([]byte, error) {
	if d.tag.BlockSize()*int(nBlocks) > MaxReadWriteMultiDataSize {
		return nil, fmt.Errorf("read extended multiple blocks exceeds max data size of %d bytes", MaxReadWriteMultiDataSize)
	}
	data, err := d.direct(card, 0x02, 0x33, uint8(startBlock&0xff), uint8((startBlock>>8)&0xff), uint8((nBlocks-1)&0xff), uint8(((nBlocks-1)>>8)&0xff))
	if err != nil {
		return nil, fmt.Errorf("failed to extended read multiple blocks: %v", err)
	}
	if len(data) != 1+(int(nBlocks)*d.tag.BlockSize()) {
		return nil, fmt.Errorf("unexpected extended read multiple blocks length: %d, expected: %d", len(data), 1+(int(nBlocks)*d.tag.BlockSize()))
	}
	return data[1:], nil
}
