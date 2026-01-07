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
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

// CBORContainerType is an encode option that is used to specify
// definite or indefinite container usage
type CBORContainerType int

const (
	// CBORContainerTypeDefinite will result in containers in the region
	// being encoided in definite form
	CBORContainerTypeDefinite CBORContainerType = iota
	// CBORContainerTypeIndefinite will result in containers in the region
	// being encoided in indefinite form
	CBORContainerTypeIndefinite
	// CBORContainerTypeAuto will utilize indefinite containers unless the
	// container is empty, in which case it optimizes to definite for one byte
	// saving
	CBORContainerTypeAuto
)

// FloatMaxPrecision is an encode option that controls the maximum precision of floating point
// numbers within a region
type FloatMaxPrecision int

const (
	// FloatMaxPrecision64 sets the maximum allowable encoded
	// floating point number representation
	// to 64 bits (high precision)
	FloatMaxPrecision64 FloatMaxPrecision = iota

	// FloatMaxPrecision32 sets the maximum allowable encoded
	// floating point number representation
	// to 32 bits (medium precision, usually fine)
	FloatMaxPrecision32

	// FloatMaxPrecision16 sets the maximum allowable
	// encoded floating point number representation
	// to 16 bits (low precision, very compact but less precise)
	FloatMaxPrecision16
)

// RegionOption contains all options related to the encoding of a tag region
type RegionOptions struct {
	cborContainerType CBORContainerType
	floatMaxPrecision FloatMaxPrecision
}

// GetCBORContainerType returns the encoding type for CBOR containers
// within a region
func (e *RegionOptions) GetCBORContainerType() CBORContainerType {
	return e.cborContainerType
}

// SetCBORContainerType sets the encoding type for CBOR containers
// within a region
func (e *RegionOptions) SetCBORContainerType(cborContainerType CBORContainerType) *RegionOptions {
	e.cborContainerType = cborContainerType
	return e
}

// GetFloatMaxPrecision returns the maximum floating point precision used
// when encoding floating point numbers within a region
func (e *RegionOptions) GetFloatMaxPrecision() FloatMaxPrecision {
	return e.floatMaxPrecision
}

// SetFloatMaxPrecision sets the maximum floating point precision used
// when encoding floating point numbers within a region
func (e *RegionOptions) SetFloatMaxPrecision(floatMaxPrecision FloatMaxPrecision) *RegionOptions {
	e.floatMaxPrecision = floatMaxPrecision
	return e
}

// A Region represents one of the three regions (meta, main, aux) within the open
// print tag
type Region interface {
	GetUnknownFields() map[any]any
	RegionOptions() *RegionOptions
	getInternal() any
	getRegionName() string
}

// EncodeDecode option provides options related to binary encoding and decoding
type EncodeDecodeOption int

const (
	// WithoutCapabilityContainer instructs the encoder/decoder to encode/decode
	// without the NFC-V capability container, which is useful when writing/reading
	// tags to non-compliant earlier NFC versions
	WithoutCapabilityContainer EncodeDecodeOption = iota
)

// Used to validate color declarations
var colorRBGValidationRegex = regexp.MustCompile(`^[0-9a-fA-F]{6,8}$`)

// ColorRGBA represents an RGB (optional A) color and should always be 3 or 4 bytes
type ColorRGBA []byte

// MarshalYAML converts ColorRGBA to #RRGGBB(AA) string representation
func (c ColorRGBA) MarshalYAML() (any, error) {
	return "#" + hex.EncodeToString(c), nil
}

// UnmarshalYAML converts YAML form of ColorRBGA back to object
func (c *ColorRGBA) UnmarshalYAML(value *yaml.Node) error {
	var str string
	if err := value.Decode(&str); err != nil {
		return err
	}

	colorBytes, err := NewColor(str)
	if err != nil {
		return err
	}
	*c = colorBytes
	return nil
}

func (c ColorRGBA) String() string {
	return "#" + hex.EncodeToString(c)
}

// NewColor will generate a ColorRGBA value from a string
// in the form of:
// #RRGGBB or #RRGGBBAA or RRGGBB or RRGGBBAA
func NewColor(str string) (ColorRGBA, error) {
	// If starts with #, remove the #
	str = strings.TrimPrefix(str, "#")

	if !colorRBGValidationRegex.MatchString(str) {
		return ColorRGBA{}, fmt.Errorf("invalid color_rgba representation: %s. (should be #rrggbb(aa)).", str)
	}

	// Can ignore error, regex already guarantess correctness
	colorBytes, _ := hex.DecodeString(str)
	return ColorRGBA(colorBytes), nil
}

// MustNewColor creates a new color, and panics in the case of an error
func MustNewColor(str string) ColorRGBA {
	c, err := NewColor(str)
	if err != nil {
		panic(fmt.Sprintf("failed to create new color from string %s: %v", str, err))
	}
	return c
}
