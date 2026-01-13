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

// icode implements specific commands for ICODE tags
type icode struct {
	commonDriver
}

// getCardInformation retrieves the block size and number of blocks from the tag
func (d *icode) getCardInformation(card *scard.Card) (blocksize, nblocks uint16, err error) {
	data, err := d.direct(card, 0x02, 0x2B)
	if err != nil {
		return 0, 0, fmt.Errorf("failed system info: %v", err)
	}
	if len(data) < 2 {
		return 0, 0, fmt.Errorf("unexpected system info length: %d", len(data))
	}
	flags := data[1] & 0x0F
	// Drop first two bytes (flags, info)
	data = data[2:]
	// Discard UID from data,
	data = data[8:]
	if flags&0x01 > 0 {
		debug("Discard DSFID")
		// DSFID present, discard
		data = data[1:]
	}
	if flags&0x02 > 0 {
		debug("Discard AFI")
		// AFI present, discard
		data = data[1:]
	}
	if flags&0x04 > 0 {
		debug("Block size from %02xh", data[1])
		blocksize = uint16((data[1] & 0x1F) + 1)
		debug("nblocks from %02xh", data[0])
		nblocks = uint16(data[0] + 1)
	} // else block info not available and blocksize/nblocks report as zero

	return blocksize, nblocks, nil
}

// readSingleBlock just uses the standard IOS15693 method
func (d *icode) readSingleBlock(card *scard.Card, block uint16) ([]byte, error) {
	return d.standardReadSingleBlock(card, 4, uint8(block))
}

// writeSingleBlock just uses the standard ISO15693 method
func (d *icode) writeSingleBlock(card *scard.Card, block uint16, data []byte) error {
	return d.standardWriteSingleBlock(card, 4, uint8(block), data)
}

// readMultipleBlocks implements the standard ISO15693 method
func (d *icode) readMultipleBlocks(card *scard.Card, startBlock, nBlocks uint16) ([]byte, error) {
	if d.tag.BlockSize()*int(nBlocks) > MaxReadWriteMultiDataSize {
		return nil, fmt.Errorf("read multiple blocks exceeds max data size of %d bytes", MaxReadWriteMultiDataSize)
	}
	data, err := d.direct(card, 0x02, 0x23, uint8(startBlock), uint8(nBlocks-1))
	if err != nil {
		return nil, fmt.Errorf("failed to read multiple blocks: %v", err)
	}
	if len(data) != 1+(int(nBlocks)*d.tag.BlockSize()) {
		return nil, fmt.Errorf("unexpected read multiple blocks length: %d, expected: %d", len(data), 1+(int(nBlocks)*d.tag.BlockSize()))
	}
	return data[1:], nil
}
