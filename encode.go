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
package openprinttag

import (
	"errors"
	"fmt"
	"slices"

	"github.com/hsanjuan/go-ndef"
	"github.com/hsanjuan/go-ndef/types/generic"
)

// Encode encodes an open print tag into it's binary format
// Errors returned indicate only a failure to encode the tag, but
// do not indicate whether or not the tag is valid
// It is perfectly possible to encode an empty tag without error
// but for the tag to not be valid (due to missing fields)
// To check tag validity, use IsValid and OptCheck
func (o *OpenPrintTag) Encode(opts ...EncodeDecodeOption) (result []byte, err error) {
	if RecoverAssertions {
		defer func() {
			if r := recover(); r != nil {
				switch x := r.(type) {
				case string:
					err = errors.New(x)
				case error:
					err = x
				default:
					err = errors.New("unknown panic occurred")
				}
			}
		}()
	}
	return o.encode(opts...)
}

// encode is the workhorse of the encoder, called by Encode
// It uses panic based assertions for consistency with the original
// python codebase, to help with readability and maintainability
// those panics are caught and transformed to errors in the main Encode
// function
func (o *OpenPrintTag) encode(opts ...EncodeDecodeOption) ([]byte, error) {
	if !slices.Contains(opts, WithoutCapabilityContainer) {
		// If we are not encoding a capability container, these checks are irrelevant
		assertTrue((o.size%8) == 0, "Tag size %d must be divisible by 8 (to be encodable in the CC)", o.size)
		assertTrue((o.size/8) <= 255, "Tag too big to be representable in the CC")
	}

	assertTrue(o.blockSize > 0, "Block size must be >0")

	capabilityContainer := []byte{
		0xE1,             // Magic Number (Version 1, Read/write access without restriction)
		0x40,             // Version
		byte(o.size / 8), // Size
		0x01,             // MBREAD
	}

	capabilityContainerSize := len(capabilityContainer)
	TLVTerminator := []byte{0xFE}
	ndefTLVHeaderSize := 2

	// Our NDEF record will be adjusted so that the message fills the whole available space
	ndefMessageLength := o.size - capabilityContainerSize - len(TLVTerminator) - ndefTLVHeaderSize

	if ndefMessageLength > 0xFE {
		// We need two more bytes to encode longer TLV lengths
		ndefTLVHeaderSize += 2
		ndefMessageLength -= 2
	}

	var ndefTLVHeader []byte
	// Do not merge with the previous if - the available space decrease might get us under this line
	if ndefMessageLength <= 0xFE {
		ndefTLVHeader = []byte{
			0x03, // NDEF message tag
			byte(ndefMessageLength),
		}
	} else {
		ndefTLVHeader = []byte{
			0x03, // NDEF message tag
			0xFF,
			byte(ndefMessageLength / 256),
			byte(ndefMessageLength % 256),
		}
	}

	assertTrue(len(ndefTLVHeader) == ndefTLVHeaderSize, "length of ndef TLV header not as expected, expected: %d, actual: %d", ndefTLVHeaderSize, len(ndefTLVHeader))

	// Set up preceeding NDEF regions
	var records = []*ndef.Record{}
	preceedingRecordsSize := 0
	if o.uri != "" {
		records = append(records, ndef.NewURIRecord(o.uri))
		preceedingRecords, err := ndef.NewMessageFromRecords(records...).Marshal()
		if err != nil {
			return nil, fmt.Errorf("failed to encode preceeding records: %w", err)
		}
		preceedingRecordsSize = len(preceedingRecords)
	}

	ndefHeaderSize := 3 + len(mimeType)
	ndefPayloadStart := capabilityContainerSize + ndefTLVHeaderSize + preceedingRecordsSize + ndefHeaderSize
	payloadSize := ndefMessageLength - ndefHeaderSize - preceedingRecordsSize

	assertTrue(payloadSize > maxMetaRegionSize, "there is not enough space even for the meta region")

	// If the NDEF payload size would exceed 255 bytes, its length cannot be stored in a single byte
	// and NDEF switches to storing the length into 4 bytes
	if payloadSize > 255 {
		ndefHeaderSize += 3
		ndefPayloadStart += 3
		payloadSize -= 3

		assertTrue(payloadSize > 255, "Unable to fill the NDEF message correctly")
	}

	payload := make([]byte, payloadSize)

	writeSection := func(offset int, data []byte) int {

		for idx, b := range data {
			payload[idx+offset] = b
		}
		return len(data)
	}

	alignRegionOffset := func(offset int, alignUp bool) int {
		misalignment := (ndefPayloadStart + offset) % o.blockSize
		if misalignment == 0 {
			return offset
		} else if alignUp {
			return offset + o.blockSize - misalignment
		} else {
			return offset - misalignment
		}
	}

	// Determine meta region offset
	mainRegionOffset := 0
	if o.metaRegionSize != 0 {
		mainRegionOffset = o.metaRegionSize
		o.MetaRegion().SetMainRegionOffset(mainRegionOffset)
	} // otherwise we're not aligning, we don't need to write the main region offset, it will be directly after the meta region

	// Prepare AUX region
	var auxRegionOffset int
	var auxRegionSizeForStats int
	var auxEncoded []byte
	if o.aux != nil {
		assertTrue(o.auxRegionSize > 4, "Aux region is too small")

		var auxRegionOffsetKnown bool
		auxRegionOffset, auxRegionOffsetKnown = o.meta.GetAuxRegionOffset()
		if !auxRegionOffsetKnown {
			auxRegionOffset = alignRegionOffset(payloadSize-o.auxRegionSize, false)
			o.meta.SetAuxRegionOffset(auxRegionOffset)
		}
		var err error
		auxEncoded, err = encodeToCBOR(o.aux)
		if err != nil {
			return nil, fmt.Errorf("failed to encode aux region: %w", err)
		}
		writeSection(auxRegionOffset, auxEncoded)
		auxRegionSizeForStats = len(auxEncoded)
	}

	// Prepare META section
	metaEncoded, err := encodeToCBOR(o.meta)
	if err != nil {
		return nil, fmt.Errorf("failed to encode meta region: %w", err)
	}
	metaSectionSize := writeSection(0, metaEncoded)

	// Prepare meta section
	// Indefinite containers take one extra byte, don't do that for the meta region - that one won't likely ever be updated
	if mainRegionOffset == 0 {
		mainRegionOffset = metaSectionSize
	}

	if o.auxRegionSize != 0 {
		assertTrue(auxRegionOffset-mainRegionOffset >= 4, "Main region is too small")
	} else {
		assertTrue(payloadSize-mainRegionOffset >= 8, "Main region is too small")
	}

	// Write the main section
	mainEncoded, err := encodeToCBOR(o.main)
	if err != nil {
		return nil, fmt.Errorf("failed to encode main region: %w", err)
	}
	writeSection(mainRegionOffset, mainEncoded)

	// Create NDEF record
	records = append(records, ndef.NewRecord(2, mimeType, "", &generic.Payload{Payload: payload}))
	// ndefData := []byte{}
	ndefMsg := ndef.NewMessageFromRecords(records...)
	ndefData, err := ndefMsg.Marshal()
	if err != nil {
		return nil, fmt.Errorf("failed to marshal ndef record: %w", err)
	}

	assertTrue(len(ndefData) == ndefMessageLength, "ndef message not expected length. Expected: %d, Actual: %d", ndefMessageLength, len(ndefData))

	// Check that we have deduced the ndef header size correctly
	expectedSize := preceedingRecordsSize + ndefHeaderSize + payloadSize
	if len(ndefData) != expectedSize {
		return nil, fmt.Errorf("NDEF record calculated incorrectly: expected size %d, (%d + %d + %d), but got %d", expectedSize, preceedingRecordsSize, ndefHeaderSize, payloadSize, len(ndefData))
	}

	fullData := []byte{}
	fullData = append(fullData, capabilityContainer...)
	fullData = append(fullData, ndefTLVHeader...)
	fullData = append(fullData, ndefData...)
	fullData = append(fullData, TLVTerminator...)

	// The full data can be slightly smaller because we might have decreased ndef_tlv_available_space by 2
	// to fit the bigger TLV header and then ended up not needing the bigger TLV header
	assertTrue(o.size-1 <= len(fullData) && len(fullData) <= o.size, "message length incorrect")

	// Check the payload is where we expect it to be
	assertTrue(slices.Equal(fullData[ndefPayloadStart:ndefPayloadStart+payloadSize], payload), "payload not in correct location")

	// All that is left is to fill in our stats object

	// root stats
	o.stats = &Stats{}
	o.stats.Root.DataSize = o.size
	o.stats.Root.PayloadSize = payloadSize
	o.stats.Root.Overhead = o.size - payloadSize
	o.stats.Root.PayloadUsedSize = len(metaEncoded) + len(mainEncoded) + auxRegionSizeForStats
	o.stats.Root.TotalUsedSize = o.stats.Root.PayloadUsedSize + o.stats.Root.Overhead

	// meta stats
	o.stats.Meta.PayloadOffset = 0
	o.stats.Meta.AbsoluteOffset = ndefPayloadStart
	o.stats.Meta.Size = len(metaEncoded) // Will always be the used size
	o.stats.Meta.UsedSize = len(metaEncoded)

	// main stats
	o.stats.Main.PayloadOffset = mainRegionOffset
	o.stats.Main.Size = payloadSize - o.stats.Meta.Size // Assumes no aux, which may be adjusted in a minute
	o.stats.Main.AbsoluteOffset = ndefPayloadStart + mainRegionOffset
	o.stats.Main.UsedSize = len(mainEncoded)

	// aux stats
	if o.aux != nil {
		o.stats.Aux = &RegionStat{}
		o.stats.Aux.UsedSize = len(auxEncoded)
		o.stats.Aux.Size = payloadSize - auxRegionOffset
		o.stats.Aux.PayloadOffset = auxRegionOffset
		o.stats.Aux.AbsoluteOffset = ndefPayloadStart + auxRegionOffset

		// With aux region present, we adjust the max available size of our main region
		o.stats.Main.Size = payloadSize - o.stats.Meta.Size - o.stats.Aux.Size
	}

	// Strip the CC header if requested
	if slices.Contains(opts, WithoutCapabilityContainer) {
		fullData = fullData[4:]
	}

	return fullData, nil
}
