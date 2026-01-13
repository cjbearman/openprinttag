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
)

// TagType enumerates the supported tags
type TagType uint8

const (
	UNKNOWN   TagType = iota
	ICODE_SLI         // N.B. untested
	ICODE_SLIX
	ICODE_SLIX2
	ST25DV04 // N.B. untested
	ST25DV16
	ST25DV64 // N.B. untested
)

// TagType returns a string representation of the TagType
func (t TagType) String() string {
	switch t {
	case UNKNOWN:
		return "UNKNOWN"
	case ICODE_SLI:
		return "ICODE SLI"
	case ICODE_SLIX:
		return "ICODE SLIX"
	case ICODE_SLIX2:
		return "ICODE_SLIX2"
	case ST25DV04:
		return "ST25DV04"
	case ST25DV16:
		return "ST25DV16"
	case ST25DV64:
		return "ST25DV64"
	default:
		panic("unknown tag type")
	}
}

// Tag represents an instance of a tag
type Tag struct {
	ttype TagType
	uid   []byte
}

// GetTagType returns the type of tag
func (t *Tag) GetTagType() TagType {
	return t.ttype
}

// GetUID returns the UID for the tag
func (t *Tag) GetUID() []byte {
	return t.uid
}

// GetUIDHex returns the UID for the tag, hex formatted
func (t *Tag) GetUIDHex() string {
	return hex.EncodeToString(t.uid)
}

// NBlocks returns the number of blocks supported by the tag
func (t *Tag) NBlocks() uint16 {
	switch t.ttype {
	case ICODE_SLI:
		return 32
	case ICODE_SLIX:
		return 28
	case ICODE_SLIX2:
		return 79
	case ST25DV04:
		return 128
	case ST25DV16:
		return 512
	case ST25DV64:
		return 2048
	default:
		return 0
	}
}

// BlockSize returns the number of bytes per block supported by the tag
func (t *Tag) BlockSize() int {
	switch t.ttype {
	case ICODE_SLI, ICODE_SLIX, ICODE_SLIX2, ST25DV04, ST25DV16, ST25DV64:
		return 4
	default:
		return 0
	}
}

// GetAvailableBytes returns the number of bytes available for reading/writing user data
func (t *Tag) GetAvailableBytes() int {
	return int(t.NBlocks()) * int(t.BlockSize())
}
