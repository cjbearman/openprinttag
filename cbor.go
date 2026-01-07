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
	"cmp"
	"fmt"
	"reflect"
	"slices"
	"strconv"
	"strings"

	"github.com/fxamacker/cbor/v2"
	"github.com/x448/float16"
)

// maxRegionSize is set at 512 in the spec
const maxRegionSize = 512

const (
	emptyDefiniteMap        = byte(0xa0)
	emptyIndefiniteMapByte1 = byte(0xbf)
	emptyIndefiniteMapByte2 = byte(0xff)
)

type reflectionMapItem struct {
	field *reflect.Value
	name  string
	key   int
}

// encodeToCBOR encodes a specific region in cbor
// Depending on encoding options, we use definite or indefinite form for containers
func encodeToCBOR(r Region) (data []byte, err error) {
	if r.RegionOptions().cborContainerType == CBORContainerTypeDefinite {
		data, err = encodeAsDefiniteMap(r)
	} else {
		// CBORContainerTypeIndefinite or CBORContainerTypeAuto
		data, err = encodeAsIndefiniteMap(r)
		if len(data) == 2 {
			if data[0] == emptyIndefiniteMapByte1 && data[1] == emptyIndefiniteMapByte2 {
				// Well that is an empy indefinite container
				if r.RegionOptions().cborContainerType == CBORContainerTypeAuto {
					// With type auto, whilst we use indefinite, we have a free hand
					// to optimize it to definite
					data = []byte{emptyDefiniteMap}
				}
			}
		}
	}
	if len(data) > maxRegionSize {
		err = fmt.Errorf("region %s size of %d exceeds maximum permissable size of %d bytes", r.getRegionName(), len(data), maxRegionSize)
	}
	return
}

// encodeAsIndefiniteMap will encode a region using an indefinite map
func encodeAsIndefiniteMap(r Region) ([]byte, error) {
	encmode, err := cbor.EncOptions{ShortestFloat: cbor.ShortestFloat16}.EncMode()
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer

	enc := encmode.NewEncoder(&buf)
	internal := reflect.ValueOf(r.getInternal()).Elem()

	theMap := mapRegion(&internal)

	enc.StartIndefiniteMap()
	for _, info := range getSortedFieldsToEncode(theMap) {
		enc.Encode(info.key)
		kind := info.field.Elem().Type().Kind()
		switch kind {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			enc.Encode(info.field.Elem().Interface())
		case reflect.Float32, reflect.Float64:
			enc.Encode(compressFloat(info.field.Elem().Interface(), r.RegionOptions()))
		case reflect.String:
			enc.Encode(info.field.Elem().Interface())
		case reflect.Slice, reflect.Array:
			kind := info.field.Elem().Type().Elem().Kind()
			if kind == reflect.Uint8 {
				// This is a []byte and must be encoded as a byte string
				enc.Encode(info.field.Elem().Interface())
			} else {
				// This is a []..something else and must be encoded as an indefinite array
				enc.StartIndefiniteArray()
				for n := 0; n < info.field.Elem().Len(); n++ {
					sliceContent := info.field.Elem().Index(n).Interface()
					enc.Encode(sliceContent)
				}
				enc.EndIndefinite()
			}

		default:
			return nil, fmt.Errorf("cannot encode %s", kind)
		}
	}

	// Add any unknowns back in
	for key, value := range r.GetUnknownFields() {
		enc.Encode(key)
		enc.Encode(value)
	}

	enc.EndIndefinite()

	return buf.Bytes(), nil
}

// encodeAsDefiniteMap will encode a region using an definite map
func encodeAsDefiniteMap(r Region) ([]byte, error) {

	encmode, err := cbor.EncOptions{ShortestFloat: cbor.ShortestFloat16}.EncMode()
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer

	enc := encmode.NewEncoder(&buf)
	internal := reflect.ValueOf(r.getInternal()).Elem()

	theMap := mapRegion(&internal)

	mapToEncode := make(map[any]any)

	for _, info := range getSortedFieldsToEncode(theMap) {
		if info.field.Elem().Type().Kind() == reflect.Float32 {
			mapToEncode[info.key] = compressFloat(info.field.Elem().Interface(), r.RegionOptions())
		} else {
			mapToEncode[info.key] = info.field.Elem().Interface()
		}
	}

	// Add any unknowns back in
	for key, value := range r.GetUnknownFields() {
		mapToEncode[key] = value
	}

	err = enc.Encode(mapToEncode)
	return buf.Bytes(), nil
}

// getSortedFieldsToEncode takes our map of the region fields
// and returns just those that are not nil in strict order
// of key. This helps keep our encoded cbor easy to debug
func getSortedFieldsToEncode(theMap map[int]reflectionMapItem) (arr []reflectionMapItem) {
	for _, item := range theMap {
		if !item.field.IsNil() {
			arr = append(arr, item)
		}
	}
	slices.SortFunc(arr, func(a, b reflectionMapItem) int {
		return cmp.Compare(a.key, b.key)
	})
	return
}

// mapRegion will take a region's internal structure
// and use it to generate a map containing the cbor fields
func mapRegion(internal *reflect.Value) map[int]reflectionMapItem {

	theMap := make(map[int]reflectionMapItem, internal.NumField())

	for i := 0; i < internal.NumField(); i++ {
		field := internal.FieldByIndex([]int{i})
		name := internal.Type().Field(i).Name

		cborTag := internal.Type().Field(i).Tag.Get("cbor")
		if cborTag == "" {
			// Not a field of interest since it has no cbor tag
			continue
		}

		cborTagBits := strings.Split(cborTag, ",")
		// The first bit of the tag should be the integer key, or -
		if cborTagBits[0] == "-" {
			// Field skipped from cbor representation
			continue
		}
		key, err := strconv.Atoi(cborTagBits[0])
		if err != nil {
			panic(fmt.Sprintf("not a valid integer key at start of cbor tag on field %s", name))
		}

		theMap[key] = reflectionMapItem{
			field: &field,
			name:  name,
			key:   key,
		}
	}
	return theMap
}

// compressFloat attempts to comrpess a float32 down to an integer
// where possible
// cbor library will already compress to float16 where a lossless
// downsizing is possible
func compressFloat(orig any, opts *RegionOptions) any {

	// the original will be either a float32 or float64
	// Start by upconverting to float64 if we can
	var original float64
	if f64, ok := orig.(float64); ok {
		original = f64
	} else if f32, ok := orig.(float32); ok {
		original = float64(f32)
	} else {
		panic("not a valid float")
	}

	// First, we'll try and optimize the number down to smallest possible integer
	// without loss

	// Is the number negative
	if original < 0 {
		// It is negative, so try first with signed integers
		if float64(int8(original)) == original {
			return int8(original)
		}
		if float64(int16(original)) == original {
			return int16(original)
		}
	}

	// The number is positive, so do that again with unsigned integers
	if float64(uint8(original)) == original {
		return uint8(original)
	}

	if float64(uint16(original)) == original {
		return uint16(original)
	}

	// It doesn't matter that we've upscaled to a float64, because
	// the cbor library automatically optimizes to the smallest usable
	// float size without loss
	// However, depending on encode options we can do a lossy downscale
	switch opts.GetFloatMaxPrecision() {
	case FloatMaxPrecision16:
		return float16.Fromfloat32(float32(original)).Float32()
	case FloatMaxPrecision32:
		return float32(original)
	default:
		return original

	}
}
