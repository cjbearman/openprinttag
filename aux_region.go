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

import (
	"time"
)

type auxInternal struct {
	ConsumedWeight          *float64    `cbor:"0,keyasint,omitempty" yaml:"consumed_weight,omitempty" opt:"name=consumed_weight,key=0"`
	Workgroup               *string     `cbor:"1,keyasint,omitempty" yaml:"workgroup,omitempty" opt:"name=workgroup,key=1,max_length=8"`
	GeneralPurposeRangeUser *string     `cbor:"2,keyasint,omitempty" yaml:"general_purpose_range_user,omitempty" opt:"name=general_purpose_range_user,key=2,max_length=8"`
	LastStirTime            *uint64     `cbor:"3,keyasint,omitempty" yaml:"last_stir_time,omitempty" opt:"name=last_stir_time,key=3"`
	Unknowns                map[any]any `cbor:"-" yaml:"other,omitempty"`
}

type AuxRegion struct {
	internal      auxInternal
	regionOptions *RegionOptions
}

// SetConsumedWeight Sets the value of consumed_weight (0)
func (s *AuxRegion) SetConsumedWeight(value float64) *AuxRegion {
	s.internal.ConsumedWeight = &value
	return s
}

// GetConsumedWeight Gets the value of consumed_weight (0)
func (s *AuxRegion) GetConsumedWeight() (float64, bool) {
	if s.internal.ConsumedWeight != nil {
		return *s.internal.ConsumedWeight, true
	}
	return 0.0, false
}

// ClearConsumedWeight Clears the value of consumed_weight (0)
func (s *AuxRegion) ClearConsumedWeight() *AuxRegion {
	s.internal.ConsumedWeight = nil
	return s
}

// SetWorkgroup Sets the value of workgroup (1)
func (s *AuxRegion) SetWorkgroup(value string) *AuxRegion {
	s.internal.Workgroup = &value
	return s
}

// GetWorkgroup Gets the value of workgroup (1)
func (s *AuxRegion) GetWorkgroup() (string, bool) {
	if s.internal.Workgroup != nil {
		return *s.internal.Workgroup, true
	}

	return "", false
}

// ClearWorkgroup Clears the value of workgroup (1)
func (s *AuxRegion) ClearWorkgroup() *AuxRegion {
	s.internal.Workgroup = nil
	return s
}

// SetGeneralPurposeRangeUser Sets the value of general_purpose_range_user (2)
func (s *AuxRegion) SetGeneralPurposeRangeUser(value string) *AuxRegion {
	s.internal.GeneralPurposeRangeUser = &value
	return s
}

// GetGeneralPurposeRangeUser Gets the value of general_purpose_range_user (2)
func (s *AuxRegion) GetGeneralPurposeRangeUser() (string, bool) {
	if s.internal.GeneralPurposeRangeUser != nil {
		return *s.internal.GeneralPurposeRangeUser, true
	}

	return "", false
}

// ClearGeneralPurposeRangeUser Clears the value of general_purpose_range_user (2)
func (s *AuxRegion) ClearGeneralPurposeRangeUser() *AuxRegion {
	s.internal.GeneralPurposeRangeUser = nil
	return s
}

// SetLastStirTime Sets the value of last_stir_time (3)
func (s *AuxRegion) SetLastStirTime(value time.Time) *AuxRegion {
	ival := uint64(value.Unix())
	s.internal.LastStirTime = &ival
	return s
}

// GetLastStirTime Gets the value of last_stir_time (3)
func (s *AuxRegion) GetLastStirTime() (time.Time, bool) {
	if s.internal.LastStirTime != nil {
		return time.Unix(int64(*s.internal.LastStirTime), 0), true
	}
	return time.Time{}, false
}

// ClearLastStirTime Clears the value of last_stir_time (3)
func (s *AuxRegion) ClearLastStirTime() *AuxRegion {
	s.internal.LastStirTime = nil
	return s
}

// GetUnknownFields returns a map of all unknown fields
// For the aux region, this will also include all vendor specific fields
func (s *AuxRegion) GetUnknownFields() map[any]any {
	if s.internal.Unknowns == nil {
		s.internal.Unknowns = make(map[any]any)
	}
	return s.internal.Unknowns
}

func (s AuxRegion) getInternal() any {
	return &s.internal
}

func (s AuxRegion) getRegionName() string {
	return "aux"
}

// RegionOptions accesses encoding options for this region
func (s AuxRegion) RegionOptions() *RegionOptions {
	return s.regionOptions
}

// GetVendorSpecificFieldKey will return any value found in the tag
// with the specified key
// The value is returned as an any type and must be cast (with checking)
// for use.
// Returns nil if no such key exists
func (s *AuxRegion) GetVendorSpecificField(key uint32) any {
	if s.internal.Unknowns == nil {
		return nil
	}
	for k, v := range s.internal.Unknowns {
		if k == uint64(key) {
			return v
		}
	}
	return nil
}

// SetVendorSpecificField sets a vendor specific field
// The value must be encodable in the cbor payload per specification
// however this is not enforced, which could cause to encoding failure
// if proper type (string, []byte, int, float) is not used
func (s *AuxRegion) SetVendorSpecificField(key uint32, value any) *AuxRegion {
	if s.internal.Unknowns == nil {
		s.internal.Unknowns = make(map[any]any)
	}
	s.internal.Unknowns[uint64(key)] = value
	return s
}
