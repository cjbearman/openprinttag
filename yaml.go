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
	"slices"

	"github.com/google/uuid"
	"gopkg.in/yaml.v3"
)

type YAMLOption int

const (
	IncludeOptCheck YAMLOption = iota
	IncludeValidation
	IncludeRootStats
	IncludeRegionStats
	IncludeURI
	IncludeUUIDs
	IncludeAll
)

// data provides the encoder/decoder for the tag data sections
type data struct {
	Meta *metaInternal `yaml:"meta,omitempty"`
	Main *mainInternal `yaml:"main,omitempty"`
	Aux  *auxInternal  `yaml:"aux,omitempty"`
}

type validate struct {
	Warnings []string `yaml:"warnings"`
	Errors   []string `yaml:"errors"`
}

type optcheck struct {
	Warnings []string `yaml:"warnings"`
	Errors   []string `yaml:"errors"`
	Notes    []string `yaml:"notes"`
}

type regionStats struct {
	Meta RegionStat  `yaml:"meta"`
	Main RegionStat  `yaml:"main"`
	Aux  *RegionStat `yaml:"aux,omitempty"`
}

type uuids struct {
	Brand    *string `yaml:"brand_uuid"`
	Material *string `yaml:"material_uuid"`
	Package  *string `yaml:"package_uuid"`
	Instance *string `yaml:"instance_uuid"`
}

// yamlJsonEncoder provides an encoder/decoder for our open print tag
type YamlEncoder struct {
	Regions   *regionStats `yaml:"regions,omitempty"`
	Root      *RootStat    `yaml:"root,omitempty"`
	Data      data         `yaml:"data"`
	UriRecord *string      `yaml:"uri,omitempty"`
	Validate  *validate    `yaml:"validate,omitempty"`
	OptCheck  *optcheck    `yaml:"opt_check,omitempty"`
	UUIDS     *uuids       `yaml:"uuids,omitempty"`
}

// prepare will prepare an open print tag representation
// in yamlEncoder form
func (o *OpenPrintTag) prepare(opts ...YAMLOption) *YamlEncoder {

	// If user has requested region stats or root stats then we should run an encode
	// to ensure up to date stats
	failedEncode := false
	if slices.Contains(opts, IncludeRegionStats) || slices.Contains(opts, IncludeRootStats) || slices.Contains(opts, IncludeAll) {
		_, err := o.Encode()
		if err != nil {
			// We can only include the stats if the encode did not error
			// If it did error, we'll not include them
			failedEncode = true
		}
	}

	encoder := YamlEncoder{}
	if o.meta != nil {
		encoder.Data.Meta = &o.meta.internal
	}
	if o.main != nil {
		encoder.Data.Main = &o.main.internal
	}
	if o.aux != nil {
		encoder.Data.Aux = &o.aux.internal
	}

	if slices.Contains(opts, IncludeValidation) || slices.Contains(opts, IncludeAll) {
		encoder.Validate = &validate{}
		encoder.Validate.Errors, encoder.Validate.Warnings = o.Validate()
	}

	if slices.Contains(opts, IncludeOptCheck) || slices.Contains(opts, IncludeAll) {
		encoder.OptCheck = &optcheck{}
		encoder.OptCheck.Errors, encoder.OptCheck.Warnings = o.OptCheck()
	}

	if slices.Contains(opts, IncludeURI) || slices.Contains(opts, IncludeAll) {
		if o.uri != "" {
			encoder.UriRecord = &o.uri
		}
	}

	if stats, gotStats := o.GetStats(); gotStats && !failedEncode {
		if gotStats {
			if slices.Contains(opts, IncludeRegionStats) || slices.Contains(opts, IncludeAll) {
				encoder.Regions = &regionStats{Meta: stats.Meta, Main: stats.Main, Aux: stats.Aux}
			}
			if slices.Contains(opts, IncludeRootStats) || slices.Contains(opts, IncludeAll) {
				encoder.Root = &stats.Root
			}
		}
	}

	if slices.Contains(opts, IncludeUUIDs) || slices.Contains(opts, IncludeAll) {
		encoder.UUIDS = &uuids{}
		o.getUUIDInformation(encoder.UUIDS)
	}

	return &encoder
}

func (o *OpenPrintTag) getUUIDInformation(uuids *uuids) {

	fromUUIDPtr := func(getter func() (uuid.UUID, bool)) *string {
		val, found := getter()
		if !found {
			return nil
		}
		str := val.String()
		return &str
	}

	var brand uuid.UUID
	var brandKnown bool

	uuids.Brand = fromUUIDPtr(o.main.GetBrandUuid)
	if uuids.Brand != nil {
		brand, _ = o.main.GetBrandUuid()
		brandKnown = true
	}
	if brandName, found := o.main.GetBrandName(); uuids.Brand == nil && found {
		brand = BrandUUID(brandName)
		str := brand.String()
		uuids.Brand = &str
		brandKnown = true
	}

	uuids.Material = fromUUIDPtr(o.main.GetMaterialUuid)
	if materialName, found := o.main.GetMaterialName(); uuids.Material == nil && brandKnown && found {
		str := MaterialUUID(materialName, brand).String()
		uuids.Material = &str
	}

	uuids.Package = fromUUIDPtr(o.main.GetPackageUuid)
	if gtin, found := o.main.GetGtin(); uuids.Package == nil && brandKnown && found {
		str := MaterialPackageUUID(fmt.Sprintf("%d", gtin), brand).String()
		uuids.Package = &str
	}

	uuids.Instance = fromUUIDPtr(o.main.GetInstanceUuid)
}

// ToYAML returns a YAML representation of the tag
// with optional options consisting of:
// IncludeValidation - Includes output from validation
// IncludeOptCheck - Includes output from opt check
// IncludeAll - Includes everything
func (o *OpenPrintTag) ToYAML(opts ...YAMLOption) (string, error) {
	obj := o.prepare(opts...)
	output, err := yaml.Marshal(obj)
	return string(output), err
}

// FromYAML reads a tag from YAML representation
func FromYAML(yamlData string) (*OpenPrintTag, error) {
	obj := YamlEncoder{}
	err := yaml.Unmarshal([]byte(yamlData), &obj)
	if err != nil {
		return nil, err
	}

	return reconstruct(obj), nil
}

// reconstruct will reconstruct a tag from yamlEncoder form
func reconstruct(from YamlEncoder) *OpenPrintTag {
	opt := NewOpenPrintTag()
	if from.UriRecord != nil {
		opt.uri = *from.UriRecord
	}
	if from.Data.Meta != nil {
		opt.meta.internal = *from.Data.Meta
	}
	if from.Data.Main != nil {
		opt.main.internal = *from.Data.Main
	}
	if from.Data.Aux != nil {
		opt.aux = &AuxRegion{internal: *from.Data.Aux}
	}

	return opt
}
