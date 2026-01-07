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
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

// dataToFill is from prusa example
var dataToFill = `data:
  main:
    material_class: FFF
    brand_name: Prusament
    material_type: PLA
    material_name: PLA Galaxy Black
    brand_specific_material_id: 1
    tags: [glitter]
    primary_color: "#3d3e3d"
    manufactured_date: 1739371290
    nominal_netto_full_weight: 1000
    actual_netto_full_weight: 1012
    empty_container_weight: 100
    transmission_distance: 0.2
    min_print_temperature: 205
    max_print_temperature: 220
    preheat_temperature: 170
    min_bed_temperature: 40
    max_bed_temperature: 60
    max_chamber_temperature: 40
    chamber_temperature: 20
    container_width: 75
    instance_uuid: 473bb8cd-e129-45b8-9fcf-da1c3add9c47
`

func TestStats(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)
	tag, err := openprinttag.FromYAML(dataToFill)
	tag.WithSize(304).WithAuxRegionSize(32)
	require.NoError(err)

	tag.MainRegion().RegionOptions().SetFloatMaxPrecision(openprinttag.FloatMaxPrecision16)
	encoded, err := tag.Encode()
	require.NoError(err)

	y, err := tag.ToYAML(openprinttag.IncludeAll)
	require.NoError(err)

	// t.Logf("%s", y)

	// Reencode to object so we can assert on some values
	enc := openprinttag.YamlEncoder{}
	err = yaml.Unmarshal([]byte(y), &enc)
	require.NoError(err)

	require.NotNil(enc.Regions)

	require.NotNil(enc.Regions.Meta)
	assert.Equal(0, enc.Regions.Meta.PayloadOffset)
	assert.Equal(42, enc.Regions.Meta.AbsoluteOffset)
	assert.Equal(4, enc.Regions.Meta.Size)
	assert.Equal(4, enc.Regions.Meta.UsedSize)

	require.NotNil(enc.Regions.Main)
	assert.Equal(4, enc.Regions.Main.PayloadOffset)
	assert.Equal(46, enc.Regions.Main.AbsoluteOffset)
	assert.Equal(222, enc.Regions.Main.Size)
	assert.Equal(119, enc.Regions.Main.UsedSize)

	require.NotNil(enc.Regions.Aux)
	assert.Equal(226, enc.Regions.Aux.PayloadOffset)
	assert.Equal(268, enc.Regions.Aux.AbsoluteOffset)
	assert.Equal(35, enc.Regions.Aux.Size)
	assert.Equal(1, enc.Regions.Aux.UsedSize)

	require.NotNil(enc.Root)
	assert.Equal(304, enc.Root.DataSize)
	assert.Equal(261, enc.Root.PayloadSize)
	assert.Equal(43, enc.Root.Overhead)
	assert.Equal(124, enc.Root.PayloadUsedSize)
	assert.Equal(167, enc.Root.TotalUsedSize)

	recordTestOutput(t, "tag", encoded)

}
