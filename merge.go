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
	"reflect"
)

// Merge will merge the data from another tag into this tag
// if overwrite is true, then properties in the source tag will be overwritten
// by properties set in the source tag, otherwise not
func (o *OpenPrintTag) Merge(other *OpenPrintTag, overwrite bool) {
	mergeRegion(reflect.ValueOf(&o.main.internal), reflect.ValueOf(&other.main.internal), overwrite)
	if other.aux != nil {
		if o.aux == nil {
			o.aux = &AuxRegion{internal: auxInternal{}}
			if _, found := o.meta.GetAuxRegionOffset(); !found {
				incomingMetaOffset, _ := other.meta.GetAuxRegionOffset()
				o.meta.SetAuxRegionOffset(incomingMetaOffset)
			}
		}
		mergeRegion(reflect.ValueOf(&o.aux.internal), reflect.ValueOf(&other.aux.internal), overwrite)
	}
}

// mergeRegion merges a specific region within the tag
func mergeRegion(orig, other reflect.Value, overwrite bool) {

	// It must have an otherInternal field, get that
	otherInternal := other.Elem()
	origInternal := orig.Elem()

	// Iterate through all fields in the internal struct
	for i := 0; i < otherInternal.NumField(); i++ {
		fieldValueOther := otherInternal.FieldByIndex([]int{i})
		fieldValueOrig := origInternal.FieldByIndex([]int{i})
		if !fieldValueOther.IsNil() {
			// This field is set in our other tag
			if !overwrite {
				// We are not overwriting, so it must not be set in our target tag
				if !fieldValueOrig.IsNil() {
					continue
				}
			}
			fieldValueOrig.Set(fieldValueOther)
		}
	}
}
