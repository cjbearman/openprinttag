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
	"time"

	"github.com/cjbearman/openprinttag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestTagInitialization tests the ability to initialize a tag
// without any actual data set and records the tag output in raw
// and yaml formats
func TestTagInitialization(t *testing.T) {
	require := require.New(t)
	tag := openprinttag.NewOpenPrintTag().
		WithAuxRegionSize(32).
		WithSize(304)

	output, err := tag.Encode()
	require.NoError(err)
	require.NotNil(output)
	yaml, err := tag.ToYAML(openprinttag.IncludeAll)
	require.NoError(err)
	recordTestOutput(t, "tag", output)
	recordTestOutput(t, "yaml", []byte(yaml))
}

// TestTagWithMainFields will encode a tag with fields set in main region
// and records the tag output in raw and yaml formats
func TestTagWithMainFields(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)
	tag := openprinttag.NewOpenPrintTag().
		WithAuxRegionSize(32).
		WithSize(304)

	brand := openprinttag.BrandUUID("Prusament")
	material := openprinttag.MaterialUUID("PLA Prusa Galaxy Black", brand)

	tag.MainRegion().
		SetBrandName("Prusament").
		SetBrandUuid(brand).
		SetMaterialName("PLA Prusa Galaxy Black").
		SetMaterialUuid(material).
		SetChamberTemperature(50).
		SetMaterialClass(openprinttag.MaterialClassFFF)
	output, err := tag.Encode()
	require.NoError(err)
	require.NotNil(output)
	yaml, err := tag.ToYAML(openprinttag.IncludeAll)
	require.NoError(err)
	recordTestOutput(t, "tag", output)
	recordTestOutput(t, "yaml", []byte(yaml))

	// Now recreate the tag from the binary representation
	// and ensure that it is identical
	reconstituted, err := openprinttag.Decode(output)
	require.NoError(err)
	assert.Equal(firstReturn(tag.MainRegion().GetBrandName()), firstReturn(reconstituted.MainRegion().GetBrandName()))
	assert.Equal(firstReturn(tag.MainRegion().GetBrandUuid()), firstReturn(reconstituted.MainRegion().GetBrandUuid()))
	assert.Equal(firstReturn(tag.MainRegion().GetMaterialName()), firstReturn(reconstituted.MainRegion().GetMaterialName()))
	assert.Equal(firstReturn(tag.MainRegion().GetMaterialUuid()), firstReturn(reconstituted.MainRegion().GetMaterialUuid()))
	assert.Equal(firstReturn(tag.MainRegion().GetChamberTemperature()), firstReturn(reconstituted.MainRegion().GetChamberTemperature()))
	assert.Equal(firstReturn(tag.MainRegion().GetMaterialClass()), firstReturn(reconstituted.MainRegion().GetMaterialClass()))
	assert.Equal(firstReturn(tag.MainRegion().GetManufacturedDate()), firstReturn(reconstituted.MainRegion().GetManufacturedDate()))
	assert.Equal(firstReturn(tag.MetaRegion().GetAuxRegionOffset()), firstReturn(reconstituted.MetaRegion().GetAuxRegionOffset()))
	assert.InDelta(firstReturn(tag.AuxRegion().GetConsumedWeight()), firstReturn(reconstituted.AuxRegion().GetConsumedWeight()), 0.001)

}

// TestTagWithMainAndAuxFields will encode a tag with fields set in main and aux regions
// and records the tag output in raw and yaml formats
func TestTagWithMainAndAuxFields(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)

	tag := openprinttag.NewOpenPrintTag().
		WithAuxRegionSize(32).
		WithSize(304)

	brand := openprinttag.BrandUUID("Prusament")
	material := openprinttag.MaterialUUID("PLA Prusa Galaxy Black", brand)

	tag.MainRegion().
		SetBrandName("Prusament").
		SetBrandUuid(brand).
		SetMaterialName("PLA Prusa Galaxy Black").
		SetMaterialUuid(material).
		SetChamberTemperature(50).
		SetMaterialClass(openprinttag.MaterialClassFFF).
		SetManufacturedDate(time.Now())

	tag.AuxRegion().
		SetConsumedWeight(1.234)
	output, err := tag.Encode()
	require.NoError(err)
	require.NotNil(output)
	yaml, err := tag.ToYAML(openprinttag.IncludeAll)
	require.NoError(err)
	recordTestOutput(t, "tag", output)
	recordTestOutput(t, "yaml", []byte(yaml))

	// Now recreate the tag from the binary representation
	// and ensure that it is identical
	reconstituted, err := openprinttag.Decode(output)
	require.NoError(err)
	assert.Equal(firstReturn(tag.MainRegion().GetBrandName()), firstReturn(reconstituted.MainRegion().GetBrandName()))
	assert.Equal(firstReturn(tag.MainRegion().GetBrandUuid()), firstReturn(reconstituted.MainRegion().GetBrandUuid()))
	assert.Equal(firstReturn(tag.MainRegion().GetMaterialName()), firstReturn(reconstituted.MainRegion().GetMaterialName()))
	assert.Equal(firstReturn(tag.MainRegion().GetMaterialUuid()), firstReturn(reconstituted.MainRegion().GetMaterialUuid()))
	assert.Equal(firstReturn(tag.MainRegion().GetChamberTemperature()), firstReturn(reconstituted.MainRegion().GetChamberTemperature()))
	assert.Equal(firstReturn(tag.MainRegion().GetMaterialClass()), firstReturn(reconstituted.MainRegion().GetMaterialClass()))
	assert.Equal(firstReturn(tag.MainRegion().GetManufacturedDate()), firstReturn(reconstituted.MainRegion().GetManufacturedDate()))
	assert.Equal(firstReturn(tag.MetaRegion().GetAuxRegionOffset()), firstReturn(reconstituted.MetaRegion().GetAuxRegionOffset()))
	assert.InDelta(firstReturn(tag.AuxRegion().GetConsumedWeight()), firstReturn(reconstituted.AuxRegion().GetConsumedWeight()), 0.001)
}

func TestTagWithUnknowns(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)

	tag := openprinttag.NewOpenPrintTag().
		WithAuxRegionSize(32).
		WithSize(304)

	brand := openprinttag.BrandUUID("Prusament")
	material := openprinttag.MaterialUUID("PLA Prusa Galaxy Black", brand)

	tag.MainRegion().
		SetBrandName("Prusament").
		SetBrandUuid(brand).
		SetMaterialName("PLA Prusa Galaxy Black").
		SetMaterialUuid(material).
		SetChamberTemperature(50).
		SetMaterialClass(openprinttag.MaterialClassFFF).
		SetManufacturedDate(time.Now())

	tag.AuxRegion().
		SetConsumedWeight(1.234)

	tag.AuxRegion().GetUnknownFields()[uint64(99)] = "fred"
	tag.AuxRegion().GetUnknownFields()[uint64(98)] = "barney"

	output, err := tag.Encode()
	require.NoError(err)
	require.NotNil(output)
	yaml, err := tag.ToYAML(openprinttag.IncludeAll)
	require.NoError(err)
	recordTestOutput(t, "tag", output)
	recordTestOutput(t, "yaml", []byte(yaml))

	// Now recreate the tag from the binary representation
	// and ensure that it is identical
	reconstituted, err := openprinttag.Decode(output)
	require.NoError(err)

	assert.Equal(firstReturn(tag.MainRegion().GetBrandName()), firstReturn(reconstituted.MainRegion().GetBrandName()))
	assert.Equal(firstReturn(tag.MainRegion().GetBrandUuid()), firstReturn(reconstituted.MainRegion().GetBrandUuid()))
	assert.Equal(firstReturn(tag.MainRegion().GetMaterialName()), firstReturn(reconstituted.MainRegion().GetMaterialName()))
	assert.Equal(firstReturn(tag.MainRegion().GetMaterialUuid()), firstReturn(reconstituted.MainRegion().GetMaterialUuid()))
	assert.Equal(firstReturn(tag.MainRegion().GetChamberTemperature()), firstReturn(reconstituted.MainRegion().GetChamberTemperature()))
	assert.Equal(firstReturn(tag.MainRegion().GetMaterialClass()), firstReturn(reconstituted.MainRegion().GetMaterialClass()))
	assert.Equal(firstReturn(tag.MainRegion().GetManufacturedDate()), firstReturn(reconstituted.MainRegion().GetManufacturedDate()))
	assert.Equal(firstReturn(tag.MetaRegion().GetAuxRegionOffset()), firstReturn(reconstituted.MetaRegion().GetAuxRegionOffset()))
	assert.InDelta(firstReturn(tag.AuxRegion().GetConsumedWeight()), firstReturn(reconstituted.AuxRegion().GetConsumedWeight()), 0.001)
	assert.Equal("fred", reconstituted.AuxRegion().GetUnknownFields()[uint64(uint64(99))])
	assert.Equal("barney", reconstituted.AuxRegion().GetUnknownFields()[uint64(98)])
}

// TestVendorSpecificField ensures that a vendor specific field can be set, encoded, decoded and retrieved
// through the aux region
func TestVendorSpecificField(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)
	tag := openprinttag.NewOpenPrintTag().
		WithAuxRegionSize(32).
		WithSize(304)

	tag.MainRegion().
		SetBrandName("Prusament")

	tag.AuxRegion().
		SetVendorSpecificField(655300, "Vendor Specific")

	output, err := tag.Encode()
	require.NoError(err)
	require.NotNil(output)
	yaml, err := tag.ToYAML(openprinttag.IncludeAll)
	require.NoError(err)
	recordTestOutput(t, "tag", output)
	recordTestOutput(t, "yaml", []byte(yaml))

	// Now recreate the tag from the binary representation
	// and ensure that it is identical
	reconstituted, err := openprinttag.Decode(output)
	require.NoError(err)
	assert.Equal(firstReturn(reconstituted.MainRegion().GetBrandName()), firstReturn(reconstituted.MainRegion().GetBrandName()))
	require.NotNil(reconstituted.AuxRegion().GetVendorSpecificField(655300))
	assert.Equal("Vendor Specific", reconstituted.AuxRegion().GetVendorSpecificField(655300))
}

// TestOversizeRegion ensures that oversize region rejection works by encoding a massively oversized vendor specific
// field in the aux region
func TestOversizeRegion(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)
	tag := openprinttag.NewOpenPrintTag().
		WithAuxRegionSize(32).
		WithSize(304)

	tag.AuxRegion().
		SetVendorSpecificField(655300, make([]byte, 1024))

	_, err := tag.Encode()
	require.Error(err)
	assert.Equal("failed to encode aux region: region aux size of 1034 exceeds maximum permissable size of 512 bytes", err.Error())
}
