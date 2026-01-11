package vtag

import (
	_ "embed"
	"encoding/hex"
)

// TagType enumerates the supported tags
type TagType uint8

const (
	UNKNOWN TagType = iota
	ICODE_SLIX
	ICODE_SLIX2
	ST25DV04
	ST25DV16
	ST25DV64
)

// TagType returns a string representation of the TagType
func (t TagType) String() string {
	switch t {
	case UNKNOWN:
		return "UNKNOWN"
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
func (t *Tag) BlockSize() uint8 {
	switch t.ttype {
	case ICODE_SLIX, ICODE_SLIX2, ST25DV04, ST25DV16, ST25DV64:
		return 4
	default:
		return 0
	}
}

// GetAvailableBytes returns the number of bytes available for reading/writing user data
func (t *Tag) GetAvailableBytes() int {
	return int(t.NBlocks()) * int(t.BlockSize())
}
