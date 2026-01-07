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
package test

import (
	"testing"

	"github.com/cjbearman/openprinttag"
	"github.com/stretchr/testify/assert"
)

// TestMergeWithoutOverwrite merges one tag into another without the overwrite
// option and verifies that the merge is successful and that the brand name, which
// is present in both tags, has not been overwriten
func TestMergeWithoutOverwrite(t *testing.T) {
	assert := assert.New(t)

	origTag := openprinttag.NewOpenPrintTag().
		WithAuxRegionSize(32).
		WithSize(304)

	brand := openprinttag.BrandUUID("Prusament")
	material := openprinttag.MaterialUUID("PLA Prusa Galaxy Black", brand)

	origTag.MainRegion().
		SetBrandName("Prusament").
		SetBrandUuid(brand).
		SetMaterialName("PLA Prusa Galaxy Black").
		SetMaterialUuid(material).
		SetChamberTemperature(50).
		SetMaterialClass(openprinttag.MaterialClassFFF)
	origTag.AuxRegion().
		SetGeneralPurposeRangeUser("fred")

	mergeToTag := openprinttag.NewOpenPrintTag().
		WithAuxRegionSize(32).
		WithSize(304)

	mergeToTag.MainRegion().
		SetBrandName("Some other brand").
		SetMaterialType(openprinttag.MaterialTypeABS)

	mergeToTag.Merge(origTag, false)

	assert.Equal("Some other brand", firstReturn(mergeToTag.MainRegion().GetBrandName()))
	assert.Equal(brand, firstReturn(mergeToTag.MainRegion().GetBrandUuid()))
	assert.Equal("PLA Prusa Galaxy Black", firstReturn(mergeToTag.MainRegion().GetMaterialName()))
	assert.Equal(material, firstReturn(mergeToTag.MainRegion().GetMaterialUuid()))
	assert.Equal(50, firstReturn(mergeToTag.MainRegion().GetChamberTemperature()))
	assert.Equal(openprinttag.MaterialClassFFF, firstReturn(mergeToTag.MainRegion().GetMaterialClass()))
	assert.Equal(openprinttag.MaterialTypeABS, firstReturn(mergeToTag.MainRegion().GetMaterialType()))
	assert.Equal("fred", firstReturn(mergeToTag.AuxRegion().GetGeneralPurposeRangeUser()))
}

// TestMergeWithOverwrite merges one tag into another with the overwrite
// option and verifies that the merge is successful and that the brand name, which
// is present in both tags, has been overwriten
func TestMergeWithOverwrite(t *testing.T) {
	assert := assert.New(t)

	origTag := openprinttag.NewOpenPrintTag().
		WithAuxRegionSize(32).
		WithSize(304)

	brand := openprinttag.BrandUUID("Prusament")
	material := openprinttag.MaterialUUID("PLA Prusa Galaxy Black", brand)

	origTag.MainRegion().
		SetBrandName("Prusament").
		SetBrandUuid(brand).
		SetMaterialName("PLA Prusa Galaxy Black").
		SetMaterialUuid(material).
		SetChamberTemperature(50).
		SetMaterialClass(openprinttag.MaterialClassFFF)
	origTag.AuxRegion().
		SetGeneralPurposeRangeUser("fred")

	mergeToTag := openprinttag.NewOpenPrintTag().
		WithAuxRegionSize(32).
		WithSize(304)

	mergeToTag.MainRegion().
		SetBrandName("Some other brand").
		SetMaterialType(openprinttag.MaterialTypeABS)

	mergeToTag.Merge(origTag, true)

	assert.Equal("Prusament", firstReturn(mergeToTag.MainRegion().GetBrandName()))
	assert.Equal(brand, firstReturn(mergeToTag.MainRegion().GetBrandUuid()))
	assert.Equal("PLA Prusa Galaxy Black", firstReturn(mergeToTag.MainRegion().GetMaterialName()))
	assert.Equal(material, firstReturn(mergeToTag.MainRegion().GetMaterialUuid()))
	assert.Equal(50, firstReturn(mergeToTag.MainRegion().GetChamberTemperature()))
	assert.Equal(openprinttag.MaterialClassFFF, firstReturn(mergeToTag.MainRegion().GetMaterialClass()))
	assert.Equal(openprinttag.MaterialTypeABS, firstReturn(mergeToTag.MainRegion().GetMaterialType()))
	assert.Equal("fred", firstReturn(mergeToTag.AuxRegion().GetGeneralPurposeRangeUser()))
}
