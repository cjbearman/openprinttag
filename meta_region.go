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

// ** THIS FILE IS AUTO-GENERATED, DO NOT MODIFY **

type metaInternal struct {
	MainRegionOffset *int        `cbor:"0,keyasint,omitempty" yaml:"main_region_offset,omitempty" opt:"name=main_region_offset,key=0"`
	MainRegionSize   *int        `cbor:"1,keyasint,omitempty" yaml:"main_region_size,omitempty" opt:"name=main_region_size,key=1"`
	AuxRegionOffset  *int        `cbor:"2,keyasint,omitempty" yaml:"aux_region_offset,omitempty" opt:"name=aux_region_offset,key=2"`
	AuxRegionSize    *int        `cbor:"3,keyasint,omitempty" yaml:"aux_region_size,omitempty" opt:"name=aux_region_size,key=3"`
	Unknowns         map[any]any `cbor:"-" yaml:"other,omitempty"`
}

type MetaRegion struct {
	internal      metaInternal
	regionOptions *RegionOptions
}

// SetMainRegionOffset Sets the value of main_region_offset (0)
func (s *MetaRegion) SetMainRegionOffset(value int) *MetaRegion {
	s.internal.MainRegionOffset = &value
	return s
}

// GetMainRegionOffset Gets the value of main_region_offset (0)
func (s *MetaRegion) GetMainRegionOffset() (int, bool) {
	if s.internal.MainRegionOffset != nil {
		return *s.internal.MainRegionOffset, true
	}
	return 0, false
}

// ClearMainRegionOffset Clears the value of main_region_offset (0)
func (s *MetaRegion) ClearMainRegionOffset() *MetaRegion {
	s.internal.MainRegionOffset = nil
	return s
}

// SetMainRegionSize Sets the value of main_region_size (1)
func (s *MetaRegion) SetMainRegionSize(value int) *MetaRegion {
	s.internal.MainRegionSize = &value
	return s
}

// GetMainRegionSize Gets the value of main_region_size (1)
func (s *MetaRegion) GetMainRegionSize() (int, bool) {
	if s.internal.MainRegionSize != nil {
		return *s.internal.MainRegionSize, true
	}
	return 0, false
}

// ClearMainRegionSize Clears the value of main_region_size (1)
func (s *MetaRegion) ClearMainRegionSize() *MetaRegion {
	s.internal.MainRegionSize = nil
	return s
}

// SetAuxRegionOffset Sets the value of aux_region_offset (2)
func (s *MetaRegion) SetAuxRegionOffset(value int) *MetaRegion {
	s.internal.AuxRegionOffset = &value
	return s
}

// GetAuxRegionOffset Gets the value of aux_region_offset (2)
func (s *MetaRegion) GetAuxRegionOffset() (int, bool) {
	if s.internal.AuxRegionOffset != nil {
		return *s.internal.AuxRegionOffset, true
	}
	return 0, false
}

// ClearAuxRegionOffset Clears the value of aux_region_offset (2)
func (s *MetaRegion) ClearAuxRegionOffset() *MetaRegion {
	s.internal.AuxRegionOffset = nil
	return s
}

// SetAuxRegionSize Sets the value of aux_region_size (3)
func (s *MetaRegion) SetAuxRegionSize(value int) *MetaRegion {
	s.internal.AuxRegionSize = &value
	return s
}

// GetAuxRegionSize Gets the value of aux_region_size (3)
func (s *MetaRegion) GetAuxRegionSize() (int, bool) {
	if s.internal.AuxRegionSize != nil {
		return *s.internal.AuxRegionSize, true
	}
	return 0, false
}

// ClearAuxRegionSize Clears the value of aux_region_size (3)
func (s *MetaRegion) ClearAuxRegionSize() *MetaRegion {
	s.internal.AuxRegionSize = nil
	return s
}

// GetUnknownFields returns a map of all unknown fields
// For the aux region, this will also include all vendor specific fields
func (s *MetaRegion) GetUnknownFields() map[any]any {
	if s.internal.Unknowns == nil {
		s.internal.Unknowns = make(map[any]any)
	}
	return s.internal.Unknowns
}

func (s MetaRegion) getInternal() any {
	return &s.internal
}

func (s MetaRegion) getRegionName() string {
	return "meta"
}

// RegionOptions accesses encoding options for this region
func (s MetaRegion) RegionOptions() *RegionOptions {
	return s.regionOptions
}
