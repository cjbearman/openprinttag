package vtag

import (
	"encoding/hex"
	"errors"
	"fmt"
	"slices"

	"github.com/ebfe/scard"
)

type APDUError struct {
	Status [2]byte
}

func (e APDUError) Error() string {
	return fmt.Sprintf("APDU error: %02x %02x", e.Status[0], e.Status[1])
}

// acr1552 implements specific transparent (direct) command handling for ACR1522-U
type acr1552 struct {
	reader string
}

// raw just sends raw bytes to the reader without any formatting
func (d *acr1552) raw(card *scard.Card, command ...byte) ([]byte, error) {
	debug("RAW: %s\n", hex.Dump(command))
	resp, err := card.Transmit(command)
	if err != nil {
		debug("ERR: %v\n", err)
		return nil, err
	}
	debug("RAW RESP :%s", hex.Dump(resp))
	return resp, nil
}

// start_transparent puts the reader into a transparent session
// which is pretty much required for all ops
func (d *acr1552) start_transparent(card *scard.Card) {
	debug("start transparent")
	d.raw(card, 0xff, 0xC2, 0x00, 0x00, 0x02, 0x81, 0x00)
}

// end_transparent ends the transparent session
func (d *acr1552) end_transparent(card *scard.Card) {
	debug("end transparent")
	d.raw(card, 0xff, 0xC2, 0x00, 0x00, 0x02, 0x82, 0x00)
}

// set_protocol_iso15693_3 sets the appropriate transparent session
// protocol
func (d *acr1552) set_protocol_iso15693_3(card *scard.Card) {
	debug("set protocol")
	d.raw(card, 0xff, 0xc2, 0x00, 0x02, 0x04, 0x8f, 0x02, 0x02, 0x03)
}

// direct formats a transparent exchange command to send the requested payload
// to the card, and returns the response from the card opr an appropriate error
func (d *acr1552) direct(card *scard.Card, payload ...byte) (resp []byte, err error) {
	// REALLY important to read 5.3.5.4 (Transparent Exchange) in the ACR1552-U manual

	// We construct the actual command we send as follows:
	// Transparent Command Header (FF C2 00 01 LEN)
	// Timeout Command (5F 46 LEN 4BYTES_LSB(TIME_IN_US))
	// Transcieve Command (95 LEN PAYLOAD)

	// This is a timeout command (5f, 46, len)
	// remaining bytes are timeout in microseconds, LSB first C350 = 50,000 microsecs
	// or 50ms which should be way more than needed
	timeoutCommand := []byte{0x5f, 0x46, 0x04, 0x50, 0xc3, 0x00, 0x00}

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

	debug("APDU TX:\n%s", hex.Dump(transparentExchange))

	resp, err = card.Transmit(transparentExchange)
	if err != nil {
		return nil, err
	}

	debug("APDU RX:\n%s", hex.Dump(resp))

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

// cardID will attempt to identify the card
func (d *acr1552) cardID(card *scard.Card) (*Tag, error) {
	debug("CARD ID")

	// Get inventory, to get the UID, which will tell us a lot
	data, err := d.direct(card, 0x26, 0x01, 0x00)
	if err != nil {
		debug("CardID error: %v", err)
		return nil, nil
	}

	if len(data) != 10 {
		return nil, errors.New("expected 10 bytes")
	}

	// UID is last 8 of 10 bytes, needs to be reversed
	uid := data[2:]
	slices.Reverse(uid)
	debug("cardID UID: %s", hex.EncodeToString(uid))

	// First three bytes E0 04 01, then this is NXP ICODE
	if uid[0] == 0xe0 && uid[1] == 0x04 && uid[2] == 0x01 {
		// Bits 37, 36 identify the tag type 00 (SLI), 10 (2/SLIX), 10 (1/SLIX2)
		// This are the following bits of the fourth byte:
		subtype := ((uid[3] & 0x18) >> 3)
		switch subtype {
		case 0x01:
			return &Tag{
				ttype: ICODE_SLIX2,
				uid:   uid,
			}, nil
		case 0x02:
			return &Tag{
				ttype: ICODE_SLIX,
				uid:   uid,
			}, nil
		default:
			return nil, errors.New("unsupported ICODE tag type")
		}
	}

	if uid[0] == 0xe0 && uid[1] == 0x02 && (uid[2] >= 0x24 && uid[2] <= 0x27) {
		// ST25DV...

		// Use extended system get info to get nblocks
		nblocks, err := d.extendedSystemGetInfo(card)
		if err != nil {
			return nil, fmt.Errorf("failed extended system get info on ST25DVxx: %v", err)
		}
		switch nblocks {
		case 128:
			return &Tag{
				ttype: ST25DV04,
				uid:   uid,
			}, nil

		case 512:
			return &Tag{
				ttype: ST25DV16,
				uid:   uid,
			}, nil
		case 2048:
			return &Tag{
				ttype: ST25DV64,
				uid:   uid,
			}, nil
		}

		return &Tag{
			ttype: ST25DV16,
			uid:   uid,
		}, nil
	}

	// So this is some card we don't know about
	return nil, errors.New("unsupported card")

}

// systemInfo returns the blocksize and nblocks, where supported, for a card
// N.B. this is an optional command and may not be supported (isn't on ST25DV16/64 where the
// nblocks cannot be encoded in the length allowed by the response)
func (d *acr1552) systemInfo(card *scard.Card) (uid []byte, blocksize, nblocks uint8, err error) {
	data, err := d.direct(card, 0x02, 0x2B)
	if err != nil {
		return nil, 0, 0, fmt.Errorf("failed system info: %v", err)
	}
	if len(data) < 2 {
		return nil, 0, 0, fmt.Errorf("unexpected system info length: %d", len(data))
	}
	flags := data[1] & 0x0F
	// Drop first two bytes (flags, info)
	data = data[2:]
	// Grab UID and discard from data, it's byte reversed so fix that
	uid = data[0:8]
	data = data[8:]
	slices.Reverse(uid)
	debug("System Info UID: %s", hex.EncodeToString(uid))
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
		blocksize = (data[1] & 0x1F) + 1
		debug("nblocks from %02xh", data[0])
		nblocks = data[0] + 1
	} // else block info not available and blocksize/nblocks report as zero

	return uid, blocksize, nblocks, nil
}

// readSingleBlock reads a single block
func (d *acr1552) readSingleBlock(card *scard.Card, bs int, block uint16) ([]byte, error) {
	if block < 256 {
		return d.standardReadSingleBlock(card, bs, uint8(block))
	} else {
		return d.extendedReadSingleBlock(card, bs, block)
	}
}

// standardReadSingleBlock reads a single block, per standard
func (d *acr1552) standardReadSingleBlock(card *scard.Card, bs int, block uint8) ([]byte, error) {
	data, err := d.direct(card, 0x02, 0x20, block)
	if err != nil {
		return nil, fmt.Errorf("failed single block read: %v", err)
	}
	if len(data) != bs+1 {
		return nil, fmt.Errorf("unexpected single block read length: %d", len(data))
	}
	return data[1:], nil
}

// extendedReadSingleBlock is an extension used by STKDV cards to address blocks >0xff
func (d *acr1552) extendedReadSingleBlock(card *scard.Card, bs int, block uint16) ([]byte, error) {
	data, err := d.direct(card, 0x02, 0x30, uint8(block&0xff), uint8((block>>8)&0xff))
	if err != nil {
		return nil, fmt.Errorf("failed extended single block read: %v", err)
	}
	if len(data) != bs+1 {
		return nil, fmt.Errorf("unexpected extended single block read length: %d", len(data))
	}
	return data[1:], nil

}

// writeSingleBlock writes a single block
func (d *acr1552) writeSingleBlock(card *scard.Card, bs int, block uint16, data []byte) error {
	if block < 256 {
		return d.standardWriteSingleBlock(card, bs, uint8(block), data)
	} else {
		return d.extendedWriteSingleBlock(card, bs, block, data)
	}
}

// standardWriteSingleBlock writes a single block in the standard per-spec method
func (d *acr1552) standardWriteSingleBlock(card *scard.Card, bs int, block uint8, data []byte) error {
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

// extendedWriteSingleBlock writes a single block with extended addressing for ST52DV.. cards, non-standard
func (d *acr1552) extendedWriteSingleBlock(card *scard.Card, bs int, block uint16, data []byte) error {
	if len(data) != bs {
		return fmt.Errorf("data length must be 4 bytes, got %d", len(data))
	}

	cmd := append([]byte{0x02, 0x31, uint8(block & 0xff), uint8((block >> 8) & 0xff)}, data...)

	data, err := d.direct(card, cmd...)
	if err != nil {
		return fmt.Errorf("failed extended single block write: %v", err)
	}

	return nil
}

// extendedSystemGetInfo is supported by STK25 cards instead of systemGetInfo, and can be used to get
// the number of blocks and is used by cardID
func (d *acr1552) extendedSystemGetInfo(card *scard.Card) (nBlocks int, err error) {
	data, err := d.direct(card, 0x02, 0x3b, 0x14)
	if err != nil {
		return 0, fmt.Errorf("failed extended system get info: %v", err)
	}
	if len(data) < 12 {
		return 0, fmt.Errorf("extended system get info short return: %v", err)
	}
	nBlocks = (int(data[11]) << 8) + int(data[10]) + 1
	debug("ESGI nblocks is %d", nBlocks)
	return nBlocks, nil
}
