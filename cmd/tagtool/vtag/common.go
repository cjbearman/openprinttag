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
	"slices"

	"github.com/ebfe/scard"
)

// MaxReadWriteMultiDataSize is the maximum data size for read/write multiple blocks
// basically a limit that can be supported by the PC/SC and ACR1552-U
const MaxReadWriteMultiDataSize = 128

// APDUYError represents an error response from an APDU command
type APDUError struct {
	Status [2]byte
}

func (e APDUError) Error() string {
	return fmt.Sprintf("APDU error: %02x %02x", e.Status[0], e.Status[1])
}

// driver represents the interface that all tag drivers must implement
type driver interface {
	// All drivers inherit from commonOps, so include all commonOps functions
	raw(card *scard.Card, command ...byte) ([]byte, error)
	start_transparent(card *scard.Card)
	end_transparent(card *scard.Card)
	set_protocol_iso15693_3(card *scard.Card)
	direct(card *scard.Card, payload ...byte) (resp []byte, err error)
	standardReadSingleBlock(card *scard.Card, bs int, block uint8) ([]byte, error)
	standardWriteSingleBlock(card *scard.Card, bs int, block uint8, data []byte) error
	getUID(card *scard.Card) (uid []byte, err error)
	read(card *scard.Card, start uint16, length uint16) ([]byte, error)
	write(card *scard.Card, start uint16, data []byte) error

	// All drivers must implement these methods
	getCardInformation(card *scard.Card) (blocksize, nblocks uint16, err error)
	readSingleBlock(card *scard.Card, block uint16) ([]byte, error)
	writeSingleBlock(card *scard.Card, block uint16, data []byte) error
	readMultipleBlocks(card *scard.Card, bstartBlock, nBlocks uint16) ([]byte, error)
}

// commonDriver implements common standard commands used by all tag types
type commonDriver struct {
	// A copy of the tag data, mainly so common methods can retrieve block sizes
	tag *Tag
	// The commonDriver needs a back reference to the final driver implementation
	// so it can call methods provided by the specific driver implementations
	driver driver
}

// raw just sends raw bytes to the reader without any formatting
func (d *commonDriver) raw(card *scard.Card, command ...byte) ([]byte, error) {
	debug("RAW: %s\n", hex.Dump(command))
	resp, err := card.Transmit(command)
	if err != nil {
		return nil, err
	}
	debug("RAW RESP :%s", hex.Dump(resp))
	return resp, nil
}

// start_transparent puts the reader into a transparent session
// which is pretty much required for all ops
func (d *commonDriver) start_transparent(card *scard.Card) {
	debug("start transparent")
	d.raw(card, 0xff, 0xC2, 0x00, 0x00, 0x02, 0x81, 0x00)
}

// end_transparent ends the transparent session
func (d *commonDriver) end_transparent(card *scard.Card) {
	debug("end transparent")
	d.raw(card, 0xff, 0xC2, 0x00, 0x00, 0x02, 0x82, 0x00)
}

// set_protocol_iso15693_3 sets the appropriate transparent session
// protocol
func (d *commonDriver) set_protocol_iso15693_3(card *scard.Card) {
	debug("set protocol")
	d.raw(card, 0xff, 0xc2, 0x00, 0x02, 0x04, 0x8f, 0x02, 0x02, 0x03)
}

// direct formats a transparent exchange command to send the requested payload
// to the card, and returns the response from the card opr an appropriate error
func (d *commonDriver) direct(card *scard.Card, payload ...byte) (resp []byte, err error) {
	// REALLY important to read 5.3.5.4 (Transparent Exchange) in the ACR1552-U manual

	// We construct the actual command we send as follows:
	// Transparent Command Header (FF C2 00 01 LEN)
	// Timeout Command (5F 46 LEN 4BYTES_LSB(TIME_IN_US))
	// Transcieve Command (95 LEN PAYLOAD)

	// This is a timeout command (5f, 46, len)
	// remaining bytes are timeout in microseconds, LSB first 30d40 = 200,000 microsecs
	// or 200ms which should be way more than needed
	timeoutCommand := []byte{0x5f, 0x46, 0x04, 0x40, 0x0d, 0x03, 0x00}

	// This is the transcieve command that sends our payload to the device
	// 95 len <the command we are sending>
	transcieveCommand := append([]byte{0x95, byte(len(payload))}, payload...)

	// We can construct the final command as the transparent exchange
	// ff c2 00 01 LEN <remainder>
	// where LEN is the length of the remainder
	// and the remainder is the concatanation of timeout and transieve
	// and then we stick 00 on the end
	transparentExchange := []byte{0xff, 0xc2, 0x00, 0x01, byte(len(timeoutCommand) + len(transcieveCommand))}

	transparentExchange = append(transparentExchange, timeoutCommand...)
	transparentExchange = append(transparentExchange, transcieveCommand...)
	transparentExchange = append(transparentExchange, 0x00)

	debug("APDU >>:\n%s", hex.Dump(transparentExchange))

	resp, err = card.Transmit(transparentExchange)
	if err != nil {
		return nil, err
	}

	debug("APDU <<:\n%s", hex.Dump(resp))

	// Let's parse the response. In case we run out of bytes, we'll catch panics
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("probably ran out of bytes parsing the response: %v", r)
		}
	}()

	if len(resp) < 5 {
		return nil, fmt.Errorf("short direct apdu response of %d bytes", len(resp))
	}

	// Discard the first three bytes
	remainder := resp[3:]

	// Hopefully the next two are 90 00, else it's an error
	if remainder[0] != 0x90 || remainder[1] != 0x00 {
		return nil, APDUError{[2]byte{resp[0], resp[1]}}
	}
	remainder = remainder[2:]

	for {
		// First byte will be the response section
		code := remainder[0]
		// next the length
		length := remainder[1]
		// and then that amount of data
		data := remainder[2 : 2+length]
		debug("Got code %02xh len %d data: %s", code, length, hex.EncodeToString(data))
		remainder = remainder[2+length:]
		if code == 0x96 {
			// Code 96 = Response Status, should be 0x00 0x00
			// the first byte, last nibble is important
			if data[0]&0x0F != 0x00 {
				return nil, fmt.Errorf("direct command response status non zero: %02xh", data[0])
			}
		}
		if code == 0x97 {
			// Code 97 has the response from the card, and it's all we're after
			return data, nil
		}
		// There's also a 0x92, we just ignore that
		// If we have no more data, we're toast
		if len(remainder) == 0 {
			return nil, errors.New("did not find response 96 code")
		}
	}
}

// getUID retrieves and returns the UID of the card
func (d *commonDriver) getUID(card *scard.Card) (uid []byte, err error) {
	debug("get-uid")
	// Get inventory
	data, err := d.direct(card, 0x26, 0x01, 0x00)
	if err != nil {
		debug("CardID error: %v", err)
		return nil, nil
	}

	if len(data) != 10 {
		return nil, errors.New("expected 10 bytes")
	}

	// UID is last 8 of 10 bytes, needs to be reversed
	uid = data[2:]
	slices.Reverse(uid)
	debug("cardID UID: %s", hex.EncodeToString(uid))
	return
}

// standardReadSingleBlock reads a single block using the standard ISO15693 command
func (d *commonDriver) standardReadSingleBlock(card *scard.Card, bs int, block uint8) ([]byte, error) {
	debug("std-read-single %d", block)
	data, err := d.direct(card, 0x02, 0x20, block)
	if err != nil {
		return nil, fmt.Errorf("failed single block read: %v", err)
	}
	if len(data) != bs+1 {
		return nil, fmt.Errorf("unexpected single block read length: %d", len(data))
	}
	return data[1:], nil
}

// standardWriteSingleBlock writes a single block using the standard ISO15693 command
func (d *commonDriver) standardWriteSingleBlock(card *scard.Card, bs int, block uint8, data []byte) error {
	debug("std-write-single %d", block)
	if len(data) != bs {
		return fmt.Errorf("data length must be %d bytes, got %d", bs, len(data))
	}

	cmd := append([]byte{0x02, 0x21, block}, data...)

	data, err := d.direct(card, cmd...)
	if err != nil {
		return fmt.Errorf("failed single block write: %v", err)
	}

	return nil
}

// read reads data from the tag starting at the specified address for the specified lengthq
// No block alignment is required, and the function will handle reading across block boundaries
func (d *commonDriver) read(card *scard.Card, start uint16, length uint16) ([]byte, error) {
	bs := d.tag.BlockSize()
	var skip uint16
	bytesToRead := length
	for start%uint16(bs) != 0 {
		// Start is not aligned to a block, which cannot be allowed
		start--       // Read one byte earlier
		skip++        // Will need to skip this byte when returning
		bytesToRead++ // Need to now read one extra byte
	}

	// Align the number of bytes we'll read to block boundary
	for bytesToRead%uint16(bs) != 0 {
		bytesToRead++
	}

	data := []byte{}
	remaining := bytesToRead
	offset := start
	for remaining > 0 {
		toRead := remaining
		if toRead > MaxReadWriteMultiDataSize {
			toRead = MaxReadWriteMultiDataSize
		}
		startBlock := offset / uint16(bs)
		numBlocks := toRead / uint16(bs)

		chunk, err := d.driver.readMultipleBlocks(card, startBlock, numBlocks)
		if err != nil {
			return nil, fmt.Errorf("failed read at block %d: %v", startBlock, err)
		}
		data = append(data, chunk...)

		remaining -= uint16(len(chunk))
		offset += uint16(len(chunk))
	}
	return data[skip : skip+length], nil
}

// write writes data to the tag starting at the specified address
// The start address must be block aligned
// The data length will be padded to block size with 0x00 bytes if needed
func (d *commonDriver) write(card *scard.Card, start uint16, data []byte) error {
	bs := uint16(d.tag.BlockSize())
	if start%bs != 0 {
		return fmt.Errorf("start address must be a multiple of %d", bs)
	}

	// Pad data to exact block sizes
	for uint16(len(data))%bs != 0 {
		data = append(data, 0x00)
	}

	if int(start)+len(data) > d.tag.GetAvailableBytes() {
		return errors.New("data exceeds tag user memory space")
	}

	nextBlock := uint16(start / bs)

	for i := 0; i < len(data); i += int(bs) {
		err := d.driver.writeSingleBlock(card, nextBlock, data[i:i+int(bs)])
		if err != nil {
			return fmt.Errorf("at block %d: %v", nextBlock, err)
		}
		nextBlock++
	}
	return nil
}
