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
	"fmt"
)

const (
	mimeType          = "application/vnd.openprinttag"
	maxMetaRegionSize = 8
)

// RecoverAssertions may be set to false for debugging
// which allow assertions to propagate as panics
// during encode and decode operations
var RecoverAssertions = true

var (
	// defaultMetaOptions are the defaults used for encode options for the meta region
	defaultMetaOptions = &RegionOptions{
		cborContainerType: CBORContainerTypeDefinite,
		floatMaxPrecision: FloatMaxPrecision32,
	}

	// defaultMainOptions are the defaults used for encode options for the main region
	defaultMainOptions = &RegionOptions{
		cborContainerType: CBORContainerTypeAuto,
		floatMaxPrecision: FloatMaxPrecision32,
	}
	// defaultAuxOptions are the defaults used for encode options for the aux region
	defaultAuxOptions = &RegionOptions{
		cborContainerType: CBORContainerTypeAuto,
		floatMaxPrecision: FloatMaxPrecision32,
	}
)

// OpenPrintTag is the primary type, representing an open print tag
// with the ability to encode, decode, modify and so forth
type OpenPrintTag struct {
	meta           *MetaRegion
	main           *MainRegion
	aux            *AuxRegion
	size           int
	blockSize      int
	uri            string
	metaRegionSize int
	auxRegionSize  int
	stats          *Stats
}

// NewOpenPrintTag creates a new, blank, open print tag
func NewOpenPrintTag() *OpenPrintTag {
	return &OpenPrintTag{
		meta:      newMetaRegion(),
		main:      newMainRegion(),
		blockSize: 4,
	}
}

// WithSize sets the print tag size and must be used when creating a blank tag
// prior to encoding
func (o *OpenPrintTag) WithSize(size int) *OpenPrintTag {
	o.size = size
	return o
}

// WithMetaRegionSize sets an optional size for the meta region
// if not set, the meta region will be encoded at the minimum size possible
func (o *OpenPrintTag) WithMetaRegionSize(size int) *OpenPrintTag {
	o.metaRegionSize = size
	return o
}

// WithAuxRegionSize will both trigger the presence of an aux region
// and set its size
func (o *OpenPrintTag) WithAuxRegionSize(size int) *OpenPrintTag {
	if o.aux == nil {
		o.aux = newAuxRegion()
	}
	o.auxRegionSize = size
	return o
}

// WithBlockSize will set an optional block size (default 4)
// which can be used to help align the aux region to a block
func (o *OpenPrintTag) WithBlockSize(blockSize int) *OpenPrintTag {
	o.blockSize = blockSize
	return o
}

// WithURIRecord adds an optional URI record to the tag
func (o *OpenPrintTag) WithURIRecord(uri string) *OpenPrintTag {
	o.uri = uri
	return o
}

// MetaRegion returns the meta region
func (o *OpenPrintTag) MetaRegion() *MetaRegion {
	return o.meta
}

// MainRegion returns the main region
func (o *OpenPrintTag) MainRegion() *MainRegion {
	return o.main
}

// Aux region returns the aux region and causes it to be initialized if not already done
func (o *OpenPrintTag) AuxRegion() *AuxRegion {
	if o.aux == nil {
		// aux is not guaranteed to have been initialized, initialize now
		o.aux = newAuxRegion()
	}
	return o.aux
}

// RemoveAuxRegion will remove the aux region from the tag
func (o *OpenPrintTag) RemoveAuxRegion() *OpenPrintTag {
	o.aux = nil
	o.auxRegionSize = 0
	return o
}

// String gives a string representation of the open print tag
// which is, in reality, just a full YAML dump
func (o *OpenPrintTag) String() string {
	str, err := o.ToYAML(IncludeAll)
	if err != nil {
		return fmt.Sprintf("unable to generate string (yaml) tag representation: %v", err)
	}
	return str
}

// IsValid returns true if the tag has no errors
func (o *OpenPrintTag) IsValid() bool {
	e, _ := o.Validate()
	return len(e) == 0
}

// newMetaRegion creates a new empty meta region, with appropriate default encoding options
func newMetaRegion() *MetaRegion {
	return &MetaRegion{
		internal:      metaInternal{},
		regionOptions: defaultMetaOptions,
	}
}

// newMainRegion creates a new empty main region, with appropriate default encoding options
func newMainRegion() *MainRegion {
	return &MainRegion{
		internal:      mainInternal{},
		regionOptions: defaultMainOptions,
	}
}

// newAuxRegion creates a new empty aux region, with appropriate default encoding options
func newAuxRegion() *AuxRegion {
	return &AuxRegion{
		internal:      auxInternal{},
		regionOptions: defaultAuxOptions,
	}
}
