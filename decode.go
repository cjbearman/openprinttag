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
	"bytes"
	"errors"
	"fmt"
	"io"
	"reflect"
	"slices"

	"github.com/fxamacker/cbor/v2"
	"github.com/hsanjuan/go-ndef"
	"github.com/hsanjuan/go-ndef/types/media"
)

// Decode reads a raw tag binary and returns the OpenPrintTag struct representation
func Decode(data []byte, opts ...EncodeDecodeOption) (opt *OpenPrintTag, err error) {
	// Since we use assertTrue, we'll catch any assertion panics
	// and change them to errors
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
	return decode(data, opts...)
}

// decode loads the tag from the raw binary, with assertion panics
func decode(tagData []byte, opts ...EncodeDecodeOption) (*OpenPrintTag, error) {

	// Create an empty tag as a starting point
	opt := NewOpenPrintTag()

	// It's size is the amount of data we received
	opt.size = len(tagData)

	// We will use a reader to stream the bytes
	data := bytes.NewReader(tagData)

	// Check capability container

	if !slices.Contains(opts, WithoutCapabilityContainer) {
		cc := make([]byte, 4)
		n, err := data.Read(cc)
		assertTrue(err == nil && n == 4, "Failed to read 4 byte cc, read: %d, err: %v", n, err)
		assertTrue(cc[0] == 0xe1, "Capability container magic number does not match")
	}

	// Find NDEF TLV
	for {
		baseTLV := make([]byte, 2)
		n, err := data.Read(baseTLV)

		// Either gone out of range or hit a terminator TLV
		tag := baseTLV[0]
		if (err != nil || n != 2) || tag == 0xFE {
			return nil, errors.New("Did not find NDEF TLV")
		}

		TLVLen := int64(baseTLV[1])

		// 0xFF means that the length takes two bytes
		if TLVLen == 0xFF {
			extLenBuf := make([]byte, 2)
			n, err := data.Read(extLenBuf)
			assertTrue(err == nil && n == 2, "Failed to read extended length")
			TLVLen = int64(extLenBuf[0])*256 + int64(extLenBuf[1])
		}

		if tag == 0x03 {
			// 0x03 = NDEF TLV, found it
			break
		}
		// Skip this TLV block
		_, err = data.Seek(TLVLen, io.SeekCurrent)
		assertTrue(err == nil, "failed TLV block skip")
	}

	// Everything that is left should be the NDEF content
	// read it all into remainingBytes
	remainingLen := data.Len()
	assertTrue(remainingLen > 0, "no NDEF records found")
	remainingBytes := make([]byte, remainingLen)
	n, err := data.Read(remainingBytes)
	assertTrue(err == nil && n == remainingLen, "Failed to read NDEF records")

	msg := ndef.Message{}
	_, err = msg.Unmarshal(remainingBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal NDEF message: %w", err)
	}

	// There may be multiple records, got to process them all
	var uriRecord, optRecord *ndef.Record
	for _, rec := range msg.Records {

		// Type U is a URI record, which we need to capture
		if rec.Type() == "U" {
			uriRecord = rec
		}

		// If the type is our mime type, it's the OPT CBOR data
		if rec.Type() == mimeType {
			optRecord = rec
		}
	}

	// If we didn't find our mime type (opt) record, it's game over
	assertTrue(optRecord != nil, "Did not find an open print tag record")

	// If there is a URI record, propagate it into the finalized open print tag
	if uriRecord != nil {
		payload, err := uriRecord.Payload()
		assertTrue(err == nil, "failed to get payload from URI record")
		opt.WithURIRecord(payload.String())
	}

	// Get the raw byte content from the OPT record, this is the data containing the CBOR regions
	optDataPayload, err := optRecord.Payload()
	odp, ok := optDataPayload.(*media.Payload)
	assertTrue(ok, "incorrect media payload type")
	optPayload := odp.Payload

	// The meta region is first
	meta := metaInternal{}
	rest, err := cbor.UnmarshalFirst(optPayload, &meta)
	assertTrue(err == nil, "failed to decode meta region: %v", err)
	opt.meta = &MetaRegion{internal: meta}
	opt.meta.internal.Unknowns, _ = getUnknownFields(opt.meta.internal, optPayload)

	// If our meta region doesn't have a main offset, it's immediately following meta
	// Since the unmarshaller returns the remaining bytes, we can calculate the offset
	// which we will refactor if the meta region has a different offset
	mainRegionOffset := len(optPayload) - len(rest)

	// Check for offsets in the meta
	if meta.MainRegionOffset != nil {
		mainRegionOffset = *meta.MainRegionOffset
	}

	var auxRegionOffset int
	if meta.AuxRegionOffset != nil {
		auxRegionOffset = *meta.AuxRegionOffset
	}

	// Load main region
	main := mainInternal{}
	rest, err = cbor.UnmarshalFirst(optPayload[mainRegionOffset:], &main)
	assertTrue(err == nil, "invalid main region")
	opt.main = &MainRegion{internal: main}

	opt.main.internal.Unknowns, _ = getUnknownFields(opt.main.internal, optPayload[mainRegionOffset:])

	// If there is no aux region offset, there is no aux region (by spec)
	// load it if we have the offset
	if auxRegionOffset != 0 {
		aux := auxInternal{}
		rest, err = cbor.UnmarshalFirst(optPayload[auxRegionOffset:], &aux)
		assertTrue(err == nil, "failed to read aux region")
		// Got an aux region
		opt.aux = &AuxRegion{internal: aux}
		opt.aux.internal.Unknowns, _ = getUnknownFields(opt.aux.internal, optPayload[auxRegionOffset:])

		// The aux region size is basically all bytes in the tag from the aux region offset
		// with a -3 offset for the NDEF structures
		opt.auxRegionSize = len(optPayload[auxRegionOffset:]) - 3
	} // if there is no aux region offset in meta, it is not present

	if opt.meta != nil {
		opt.meta.regionOptions = defaultMetaOptions
	}
	if opt.main != nil {
		opt.main.regionOptions = defaultMainOptions
	}
	if opt.aux != nil {
		opt.aux.regionOptions = defaultAuxOptions
	}

	// Done.
	// N.B. the blockSize setting may be incorrect, but if we have aux we have an aux region offset
	// already, so this is somewhat irrelevant
	return opt, nil
}

// The main cbor parser skips unknown fields, which is not allowed by spec
// we have to reparse everything into a dynamic structure, identify unknown fields
// and include them separately
// Although we will return an error if this fails (unlikely), the above code
// likely just ignores these errors as sub-critical
func getUnknownFields(forInternal any, raw []byte) (unknownFieldsMap map[any]any, err error) {

	// This map will contain any and all unknown cbor elements
	unknownFields := make(map[any]any)

	// Reflect the internal struct to get a map of which keys are known and which are not
	var internal = reflect.ValueOf(forInternal)
	knownFields := mapRegion(&internal)

	var dynamicData map[any]interface{}
	_, err = cbor.UnmarshalFirst(raw, &dynamicData)
	if err != nil {
		return
	}

	for key, value := range dynamicData {
		if keyInt, ok := key.(uint64); ok {
			// This is an integer key, which are all that are used in open print tag
			// do we know it
			if _, found := knownFields[int(keyInt)]; found {
				continue
			}
		}

		// This is not a known key, it gets added to our unknown fields
		unknownFields[key] = value
	}

	if len(unknownFields) > 0 {
		unknownFieldsMap = unknownFields
	}
	return
}
