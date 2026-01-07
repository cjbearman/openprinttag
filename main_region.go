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
	"bytes"
	"github.com/google/uuid"
	"time"
)

type mainInternal struct {
	InstanceUuid                *uuid.UUID                `cbor:"0,keyasint,omitempty" yaml:"instance_uuid,omitempty" opt:"name=instance_uuid,key=0"`
	PackageUuid                 *uuid.UUID                `cbor:"1,keyasint,omitempty" yaml:"package_uuid,omitempty" opt:"name=package_uuid,key=1"`
	MaterialUuid                *uuid.UUID                `cbor:"2,keyasint,omitempty" yaml:"material_uuid,omitempty" opt:"name=material_uuid,key=2"`
	BrandUuid                   *uuid.UUID                `cbor:"3,keyasint,omitempty" yaml:"brand_uuid,omitempty" opt:"name=brand_uuid,key=3"`
	Gtin                        *uint64                   `cbor:"4,keyasint,omitempty" yaml:"gtin,omitempty" opt:"name=gtin,key=4,recommended"`
	BrandSpecificInstanceId     *string                   `cbor:"5,keyasint,omitempty" yaml:"brand_specific_instance_id,omitempty" opt:"name=brand_specific_instance_id,key=5,max_length=16"`
	BrandSpecificPackageId      *string                   `cbor:"6,keyasint,omitempty" yaml:"brand_specific_package_id,omitempty" opt:"name=brand_specific_package_id,key=6,max_length=16"`
	BrandSpecificMaterialId     *string                   `cbor:"7,keyasint,omitempty" yaml:"brand_specific_material_id,omitempty" opt:"name=brand_specific_material_id,key=7,max_length=16"`
	MaterialClass               *MaterialClass            `cbor:"8,keyasint,omitempty" yaml:"material_class,omitempty" opt:"name=material_class,key=8,required"`
	MaterialType                *MaterialType             `cbor:"9,keyasint,omitempty" yaml:"material_type,omitempty" opt:"name=material_type,key=9,recommended"`
	MaterialName                *string                   `cbor:"10,keyasint,omitempty" yaml:"material_name,omitempty" opt:"name=material_name,key=10,recommended,max_length=31"`
	MaterialAbbreviation        *string                   `cbor:"52,keyasint,omitempty" yaml:"material_abbreviation,omitempty" opt:"name=material_abbreviation,key=52,max_length=7"`
	BrandName                   *string                   `cbor:"11,keyasint,omitempty" yaml:"brand_name,omitempty" opt:"name=brand_name,key=11,recommended,max_length=31"`
	WriteProtection             *WriteProtection          `cbor:"13,keyasint,omitempty" yaml:"write_protection,omitempty" opt:"name=write_protection,key=13"`
	ManufacturedDate            *uint64                   `cbor:"14,keyasint,omitempty" yaml:"manufactured_date,omitempty" opt:"name=manufactured_date,key=14,recommended"`
	CountryOfOrigin             *string                   `cbor:"55,keyasint,omitempty" yaml:"country_of_origin,omitempty" opt:"name=country_of_origin,key=55,max_length=2"`
	ExpirationDate              *uint64                   `cbor:"15,keyasint,omitempty" yaml:"expiration_date,omitempty" opt:"name=expiration_date,key=15"`
	NominalNettoFullWeight      *float64                  `cbor:"16,keyasint,omitempty" yaml:"nominal_netto_full_weight,omitempty" opt:"name=nominal_netto_full_weight,key=16,recommended"`
	ActualNettoFullWeight       *float64                  `cbor:"17,keyasint,omitempty" yaml:"actual_netto_full_weight,omitempty" opt:"name=actual_netto_full_weight,key=17,recommended"`
	NominalFullLength           *float64                  `cbor:"53,keyasint,omitempty" yaml:"nominal_full_length,omitempty" opt:"name=nominal_full_length,key=53"`
	ActualFullLength            *float64                  `cbor:"54,keyasint,omitempty" yaml:"actual_full_length,omitempty" opt:"name=actual_full_length,key=54"`
	EmptyContainerWeight        *float64                  `cbor:"18,keyasint,omitempty" yaml:"empty_container_weight,omitempty" opt:"name=empty_container_weight,key=18,recommended"`
	PrimaryColor                *ColorRGBA                `cbor:"19,keyasint,omitempty" yaml:"primary_color,flow,omitempty" opt:"name=primary_color,key=19,recommended,rgba"`
	SecondaryColor0             *ColorRGBA                `cbor:"20,keyasint,omitempty" yaml:"secondary_color_0,flow,omitempty" opt:"name=secondary_color_0,key=20,rgba"`
	SecondaryColor1             *ColorRGBA                `cbor:"21,keyasint,omitempty" yaml:"secondary_color_1,flow,omitempty" opt:"name=secondary_color_1,key=21,rgba"`
	SecondaryColor2             *ColorRGBA                `cbor:"22,keyasint,omitempty" yaml:"secondary_color_2,flow,omitempty" opt:"name=secondary_color_2,key=22,rgba"`
	SecondaryColor3             *ColorRGBA                `cbor:"23,keyasint,omitempty" yaml:"secondary_color_3,flow,omitempty" opt:"name=secondary_color_3,key=23,rgba"`
	SecondaryColor4             *ColorRGBA                `cbor:"24,keyasint,omitempty" yaml:"secondary_color_4,flow,omitempty" opt:"name=secondary_color_4,key=24,rgba"`
	TransmissionDistance        *float64                  `cbor:"27,keyasint,omitempty" yaml:"transmission_distance,omitempty" opt:"name=transmission_distance,key=27"`
	Tags                        *[]Tag                    `cbor:"28,keyasint,omitempty" yaml:"tags,omitempty" opt:"name=tags,key=28,recommended,max_length=16"`
	Certifications              *[]MaterialCertifications `cbor:"56,keyasint,omitempty" yaml:"certifications,omitempty" opt:"name=certifications,key=56,max_length=8"`
	Density                     *float64                  `cbor:"29,keyasint,omitempty" yaml:"density,omitempty" opt:"name=density,key=29,recommended"`
	FilamentDiameter            *float64                  `cbor:"30,keyasint,omitempty" yaml:"filament_diameter,omitempty" opt:"name=filament_diameter,key=30"`
	ShoreHardnessA              *int                      `cbor:"31,keyasint,omitempty" yaml:"shore_hardness_a,omitempty" opt:"name=shore_hardness_a,key=31"`
	ShoreHardnessD              *int                      `cbor:"32,keyasint,omitempty" yaml:"shore_hardness_d,omitempty" opt:"name=shore_hardness_d,key=32"`
	MinNozzleDiameter           *float64                  `cbor:"33,keyasint,omitempty" yaml:"min_nozzle_diameter,omitempty" opt:"name=min_nozzle_diameter,key=33"`
	MinPrintTemperature         *int                      `cbor:"34,keyasint,omitempty" yaml:"min_print_temperature,omitempty" opt:"name=min_print_temperature,key=34,recommended"`
	MaxPrintTemperature         *int                      `cbor:"35,keyasint,omitempty" yaml:"max_print_temperature,omitempty" opt:"name=max_print_temperature,key=35,recommended"`
	PreheatTemperature          *int                      `cbor:"36,keyasint,omitempty" yaml:"preheat_temperature,omitempty" opt:"name=preheat_temperature,key=36,recommended"`
	MinBedTemperature           *int                      `cbor:"37,keyasint,omitempty" yaml:"min_bed_temperature,omitempty" opt:"name=min_bed_temperature,key=37,recommended"`
	MaxBedTemperature           *int                      `cbor:"38,keyasint,omitempty" yaml:"max_bed_temperature,omitempty" opt:"name=max_bed_temperature,key=38,recommended"`
	MinChamberTemperature       *int                      `cbor:"39,keyasint,omitempty" yaml:"min_chamber_temperature,omitempty" opt:"name=min_chamber_temperature,key=39"`
	MaxChamberTemperature       *int                      `cbor:"40,keyasint,omitempty" yaml:"max_chamber_temperature,omitempty" opt:"name=max_chamber_temperature,key=40"`
	ChamberTemperature          *int                      `cbor:"41,keyasint,omitempty" yaml:"chamber_temperature,omitempty" opt:"name=chamber_temperature,key=41"`
	ContainerWidth              *int                      `cbor:"42,keyasint,omitempty" yaml:"container_width,omitempty" opt:"name=container_width,key=42"`
	ContainerOuterDiameter      *int                      `cbor:"43,keyasint,omitempty" yaml:"container_outer_diameter,omitempty" opt:"name=container_outer_diameter,key=43"`
	ContainerInnerDiameter      *int                      `cbor:"44,keyasint,omitempty" yaml:"container_inner_diameter,omitempty" opt:"name=container_inner_diameter,key=44"`
	ContainerHoleDiameter       *int                      `cbor:"45,keyasint,omitempty" yaml:"container_hole_diameter,omitempty" opt:"name=container_hole_diameter,key=45"`
	Viscosity18C                *float64                  `cbor:"46,keyasint,omitempty" yaml:"viscosity_18c,omitempty" opt:"name=viscosity_18c,key=46"`
	Viscosity25C                *float64                  `cbor:"47,keyasint,omitempty" yaml:"viscosity_25c,omitempty" opt:"name=viscosity_25c,key=47"`
	Viscosity40C                *float64                  `cbor:"48,keyasint,omitempty" yaml:"viscosity_40c,omitempty" opt:"name=viscosity_40c,key=48"`
	Viscosity60C                *float64                  `cbor:"49,keyasint,omitempty" yaml:"viscosity_60c,omitempty" opt:"name=viscosity_60c,key=49"`
	ContainerVolumetricCapacity *float64                  `cbor:"50,keyasint,omitempty" yaml:"container_volumetric_capacity,omitempty" opt:"name=container_volumetric_capacity,key=50"`
	CureWavelength              *int                      `cbor:"51,keyasint,omitempty" yaml:"cure_wavelength,omitempty" opt:"name=cure_wavelength,key=51"`
	DryingTemperature           *int                      `cbor:"57,keyasint,omitempty" yaml:"drying_temperature,omitempty" opt:"name=drying_temperature,key=57"`
	DryingTime                  *int                      `cbor:"58,keyasint,omitempty" yaml:"drying_time,omitempty" opt:"name=drying_time,key=58"`
	Unknowns                    map[any]any               `cbor:"-" yaml:"other,omitempty"`
}

type MainRegion struct {
	internal      mainInternal
	regionOptions *RegionOptions
}

// SetInstanceUuid Sets the value of instance_uuid (0)
func (s *MainRegion) SetInstanceUuid(value uuid.UUID) *MainRegion {
	s.internal.InstanceUuid = &value
	return s
}

// GetInstanceUuid Gets the value of instance_uuid (0)
func (s *MainRegion) GetInstanceUuid() (uuid.UUID, bool) {
	if s.internal.InstanceUuid != nil {
		return *s.internal.InstanceUuid, true
	}
	return uuid.UUID{}, false
}

// ClearInstanceUuid Clears the value of instance_uuid (0)
func (s *MainRegion) ClearInstanceUuid() *MainRegion {
	s.internal.InstanceUuid = nil
	return s
}

// SetPackageUuid Sets the value of package_uuid (1)
func (s *MainRegion) SetPackageUuid(value uuid.UUID) *MainRegion {
	s.internal.PackageUuid = &value
	return s
}

// GetPackageUuid Gets the value of package_uuid (1)
func (s *MainRegion) GetPackageUuid() (uuid.UUID, bool) {
	if s.internal.PackageUuid != nil {
		return *s.internal.PackageUuid, true
	}
	return uuid.UUID{}, false
}

// ClearPackageUuid Clears the value of package_uuid (1)
func (s *MainRegion) ClearPackageUuid() *MainRegion {
	s.internal.PackageUuid = nil
	return s
}

// SetMaterialUuid Sets the value of material_uuid (2)
func (s *MainRegion) SetMaterialUuid(value uuid.UUID) *MainRegion {
	s.internal.MaterialUuid = &value
	return s
}

// GetMaterialUuid Gets the value of material_uuid (2)
func (s *MainRegion) GetMaterialUuid() (uuid.UUID, bool) {
	if s.internal.MaterialUuid != nil {
		return *s.internal.MaterialUuid, true
	}
	return uuid.UUID{}, false
}

// ClearMaterialUuid Clears the value of material_uuid (2)
func (s *MainRegion) ClearMaterialUuid() *MainRegion {
	s.internal.MaterialUuid = nil
	return s
}

// SetBrandUuid Sets the value of brand_uuid (3)
func (s *MainRegion) SetBrandUuid(value uuid.UUID) *MainRegion {
	s.internal.BrandUuid = &value
	return s
}

// GetBrandUuid Gets the value of brand_uuid (3)
func (s *MainRegion) GetBrandUuid() (uuid.UUID, bool) {
	if s.internal.BrandUuid != nil {
		return *s.internal.BrandUuid, true
	}
	return uuid.UUID{}, false
}

// ClearBrandUuid Clears the value of brand_uuid (3)
func (s *MainRegion) ClearBrandUuid() *MainRegion {
	s.internal.BrandUuid = nil
	return s
}

// SetGtin Sets the value of gtin (4)
// Global Trade Item Number.
func (s *MainRegion) SetGtin(value uint64) *MainRegion {
	s.internal.Gtin = &value
	return s
}

// GetGtin Gets the value of gtin (4)
// Global Trade Item Number.
func (s *MainRegion) GetGtin() (uint64, bool) {
	if s.internal.Gtin != nil {
		return *s.internal.Gtin, true
	}
	return 0, false
}

// ClearGtin Clears the value of gtin (4)
// Global Trade Item Number.
func (s *MainRegion) ClearGtin() *MainRegion {
	s.internal.Gtin = nil
	return s
}

// SetBrandSpecificInstanceId Sets the value of brand_specific_instance_id (5)
func (s *MainRegion) SetBrandSpecificInstanceId(value string) *MainRegion {
	s.internal.BrandSpecificInstanceId = &value
	return s
}

// GetBrandSpecificInstanceId Gets the value of brand_specific_instance_id (5)
func (s *MainRegion) GetBrandSpecificInstanceId() (string, bool) {
	if s.internal.BrandSpecificInstanceId != nil {
		return *s.internal.BrandSpecificInstanceId, true
	}

	return "", false
}

// ClearBrandSpecificInstanceId Clears the value of brand_specific_instance_id (5)
func (s *MainRegion) ClearBrandSpecificInstanceId() *MainRegion {
	s.internal.BrandSpecificInstanceId = nil
	return s
}

// SetBrandSpecificPackageId Sets the value of brand_specific_package_id (6)
func (s *MainRegion) SetBrandSpecificPackageId(value string) *MainRegion {
	s.internal.BrandSpecificPackageId = &value
	return s
}

// GetBrandSpecificPackageId Gets the value of brand_specific_package_id (6)
func (s *MainRegion) GetBrandSpecificPackageId() (string, bool) {
	if s.internal.BrandSpecificPackageId != nil {
		return *s.internal.BrandSpecificPackageId, true
	}

	return "", false
}

// ClearBrandSpecificPackageId Clears the value of brand_specific_package_id (6)
func (s *MainRegion) ClearBrandSpecificPackageId() *MainRegion {
	s.internal.BrandSpecificPackageId = nil
	return s
}

// SetBrandSpecificMaterialId Sets the value of brand_specific_material_id (7)
func (s *MainRegion) SetBrandSpecificMaterialId(value string) *MainRegion {
	s.internal.BrandSpecificMaterialId = &value
	return s
}

// GetBrandSpecificMaterialId Gets the value of brand_specific_material_id (7)
func (s *MainRegion) GetBrandSpecificMaterialId() (string, bool) {
	if s.internal.BrandSpecificMaterialId != nil {
		return *s.internal.BrandSpecificMaterialId, true
	}

	return "", false
}

// ClearBrandSpecificMaterialId Clears the value of brand_specific_material_id (7)
func (s *MainRegion) ClearBrandSpecificMaterialId() *MainRegion {
	s.internal.BrandSpecificMaterialId = nil
	return s
}

// SetMaterialClass Sets the value of material_class (8)
func (s *MainRegion) SetMaterialClass(value MaterialClass) *MainRegion {
	s.internal.MaterialClass = &value
	return s
}

// GetMaterialClass Gets the value of material_class (8)
func (s *MainRegion) GetMaterialClass() (MaterialClass, bool) {
	if s.internal.MaterialClass != nil {
		return MaterialClass(*s.internal.MaterialClass), true
	}
	return MaterialClass(0), false
}

// ClearMaterialClass Clears the value of material_class (8)
func (s *MainRegion) ClearMaterialClass() *MainRegion {
	s.internal.MaterialClass = nil
	return s
}

// SetMaterialType Sets the value of material_type (9)
func (s *MainRegion) SetMaterialType(value MaterialType) *MainRegion {
	s.internal.MaterialType = &value
	return s
}

// GetMaterialType Gets the value of material_type (9)
func (s *MainRegion) GetMaterialType() (MaterialType, bool) {
	if s.internal.MaterialType != nil {
		return MaterialType(*s.internal.MaterialType), true
	}
	return MaterialType(0), false
}

// ClearMaterialType Clears the value of material_type (9)
func (s *MainRegion) ClearMaterialType() *MainRegion {
	s.internal.MaterialType = nil
	return s
}

// SetMaterialName Sets the value of material_name (10)
func (s *MainRegion) SetMaterialName(value string) *MainRegion {
	s.internal.MaterialName = &value
	return s
}

// GetMaterialName Gets the value of material_name (10)
func (s *MainRegion) GetMaterialName() (string, bool) {
	if s.internal.MaterialName != nil {
		return *s.internal.MaterialName, true
	}

	return "", false
}

// ClearMaterialName Clears the value of material_name (10)
func (s *MainRegion) ClearMaterialName() *MainRegion {
	s.internal.MaterialName = nil
	return s
}

// SetMaterialAbbreviation Sets the value of material_abbreviation (52)
func (s *MainRegion) SetMaterialAbbreviation(value string) *MainRegion {
	s.internal.MaterialAbbreviation = &value
	return s
}

// GetMaterialAbbreviation Gets the value of material_abbreviation (52)
func (s *MainRegion) GetMaterialAbbreviation() (string, bool) {
	if s.internal.MaterialAbbreviation != nil {
		return *s.internal.MaterialAbbreviation, true
	}

	return "", false
}

// ClearMaterialAbbreviation Clears the value of material_abbreviation (52)
func (s *MainRegion) ClearMaterialAbbreviation() *MainRegion {
	s.internal.MaterialAbbreviation = nil
	return s
}

// SetBrandName Sets the value of brand_name (11)
// Brand of the material.
func (s *MainRegion) SetBrandName(value string) *MainRegion {
	s.internal.BrandName = &value
	return s
}

// GetBrandName Gets the value of brand_name (11)
// Brand of the material.
func (s *MainRegion) GetBrandName() (string, bool) {
	if s.internal.BrandName != nil {
		return *s.internal.BrandName, true
	}

	return "", false
}

// ClearBrandName Clears the value of brand_name (11)
// Brand of the material.
func (s *MainRegion) ClearBrandName() *MainRegion {
	s.internal.BrandName = nil
	return s
}

// SetWriteProtection Sets the value of write_protection (13)
func (s *MainRegion) SetWriteProtection(value WriteProtection) *MainRegion {
	s.internal.WriteProtection = &value
	return s
}

// GetWriteProtection Gets the value of write_protection (13)
func (s *MainRegion) GetWriteProtection() (WriteProtection, bool) {
	if s.internal.WriteProtection != nil {
		return WriteProtection(*s.internal.WriteProtection), true
	}
	return WriteProtection(0), false
}

// ClearWriteProtection Clears the value of write_protection (13)
func (s *MainRegion) ClearWriteProtection() *MainRegion {
	s.internal.WriteProtection = nil
	return s
}

// SetManufacturedDate Sets the value of manufactured_date (14)
func (s *MainRegion) SetManufacturedDate(value time.Time) *MainRegion {
	ival := uint64(value.Unix())
	s.internal.ManufacturedDate = &ival
	return s
}

// GetManufacturedDate Gets the value of manufactured_date (14)
func (s *MainRegion) GetManufacturedDate() (time.Time, bool) {
	if s.internal.ManufacturedDate != nil {
		return time.Unix(int64(*s.internal.ManufacturedDate), 0), true
	}
	return time.Time{}, false
}

// ClearManufacturedDate Clears the value of manufactured_date (14)
func (s *MainRegion) ClearManufacturedDate() *MainRegion {
	s.internal.ManufacturedDate = nil
	return s
}

// SetCountryOfOrigin Sets the value of country_of_origin (55)
// Country the [MaterialPackageInstance](terminology) was produced in, encoded as a two-letter code according to [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
func (s *MainRegion) SetCountryOfOrigin(value string) *MainRegion {
	s.internal.CountryOfOrigin = &value
	return s
}

// GetCountryOfOrigin Gets the value of country_of_origin (55)
// Country the [MaterialPackageInstance](terminology) was produced in, encoded as a two-letter code according to [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
func (s *MainRegion) GetCountryOfOrigin() (string, bool) {
	if s.internal.CountryOfOrigin != nil {
		return *s.internal.CountryOfOrigin, true
	}

	return "", false
}

// ClearCountryOfOrigin Clears the value of country_of_origin (55)
// Country the [MaterialPackageInstance](terminology) was produced in, encoded as a two-letter code according to [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
func (s *MainRegion) ClearCountryOfOrigin() *MainRegion {
	s.internal.CountryOfOrigin = nil
	return s
}

// SetExpirationDate Sets the value of expiration_date (15)
func (s *MainRegion) SetExpirationDate(value time.Time) *MainRegion {
	ival := uint64(value.Unix())
	s.internal.ExpirationDate = &ival
	return s
}

// GetExpirationDate Gets the value of expiration_date (15)
func (s *MainRegion) GetExpirationDate() (time.Time, bool) {
	if s.internal.ExpirationDate != nil {
		return time.Unix(int64(*s.internal.ExpirationDate), 0), true
	}
	return time.Time{}, false
}

// ClearExpirationDate Clears the value of expiration_date (15)
func (s *MainRegion) ClearExpirationDate() *MainRegion {
	s.internal.ExpirationDate = nil
	return s
}

// SetNominalNettoFullWeight Sets the value of nominal_netto_full_weight (16)
func (s *MainRegion) SetNominalNettoFullWeight(value float64) *MainRegion {
	s.internal.NominalNettoFullWeight = &value
	return s
}

// GetNominalNettoFullWeight Gets the value of nominal_netto_full_weight (16)
func (s *MainRegion) GetNominalNettoFullWeight() (float64, bool) {
	if s.internal.NominalNettoFullWeight != nil {
		return *s.internal.NominalNettoFullWeight, true
	}
	return 0.0, false
}

// ClearNominalNettoFullWeight Clears the value of nominal_netto_full_weight (16)
func (s *MainRegion) ClearNominalNettoFullWeight() *MainRegion {
	s.internal.NominalNettoFullWeight = nil
	return s
}

// SetActualNettoFullWeight Sets the value of actual_netto_full_weight (17)
func (s *MainRegion) SetActualNettoFullWeight(value float64) *MainRegion {
	s.internal.ActualNettoFullWeight = &value
	return s
}

// GetActualNettoFullWeight Gets the value of actual_netto_full_weight (17)
func (s *MainRegion) GetActualNettoFullWeight() (float64, bool) {
	if s.internal.ActualNettoFullWeight != nil {
		return *s.internal.ActualNettoFullWeight, true
	}
	return 0.0, false
}

// ClearActualNettoFullWeight Clears the value of actual_netto_full_weight (17)
func (s *MainRegion) ClearActualNettoFullWeight() *MainRegion {
	s.internal.ActualNettoFullWeight = nil
	return s
}

// SetNominalFullLength Sets the value of nominal_full_length (53)
func (s *MainRegion) SetNominalFullLength(value float64) *MainRegion {
	s.internal.NominalFullLength = &value
	return s
}

// GetNominalFullLength Gets the value of nominal_full_length (53)
func (s *MainRegion) GetNominalFullLength() (float64, bool) {
	if s.internal.NominalFullLength != nil {
		return *s.internal.NominalFullLength, true
	}
	return 0.0, false
}

// ClearNominalFullLength Clears the value of nominal_full_length (53)
func (s *MainRegion) ClearNominalFullLength() *MainRegion {
	s.internal.NominalFullLength = nil
	return s
}

// SetActualFullLength Sets the value of actual_full_length (54)
func (s *MainRegion) SetActualFullLength(value float64) *MainRegion {
	s.internal.ActualFullLength = &value
	return s
}

// GetActualFullLength Gets the value of actual_full_length (54)
func (s *MainRegion) GetActualFullLength() (float64, bool) {
	if s.internal.ActualFullLength != nil {
		return *s.internal.ActualFullLength, true
	}
	return 0.0, false
}

// ClearActualFullLength Clears the value of actual_full_length (54)
func (s *MainRegion) ClearActualFullLength() *MainRegion {
	s.internal.ActualFullLength = nil
	return s
}

// SetEmptyContainerWeight Sets the value of empty_container_weight (18)
// Weight of the empty container.
func (s *MainRegion) SetEmptyContainerWeight(value float64) *MainRegion {
	s.internal.EmptyContainerWeight = &value
	return s
}

// GetEmptyContainerWeight Gets the value of empty_container_weight (18)
// Weight of the empty container.
func (s *MainRegion) GetEmptyContainerWeight() (float64, bool) {
	if s.internal.EmptyContainerWeight != nil {
		return *s.internal.EmptyContainerWeight, true
	}
	return 0.0, false
}

// ClearEmptyContainerWeight Clears the value of empty_container_weight (18)
// Weight of the empty container.
func (s *MainRegion) ClearEmptyContainerWeight() *MainRegion {
	s.internal.EmptyContainerWeight = nil
	return s
}

// SetPrimaryColor Sets the value of primary_color (19)
func (s *MainRegion) SetPrimaryColor(value ColorRGBA) *MainRegion {
	copy := ColorRGBA(bytes.Clone(value))
	s.internal.PrimaryColor = &copy
	return s
}

// GetPrimaryColor Gets the value of primary_color (19)
func (s *MainRegion) GetPrimaryColor() (ColorRGBA, bool) {
	if s.internal.PrimaryColor != nil {
		return *s.internal.PrimaryColor, true
	}
	return ColorRGBA{0, 0, 0}, false
}

// ClearPrimaryColor Clears the value of primary_color (19)
func (s *MainRegion) ClearPrimaryColor() *MainRegion {
	s.internal.PrimaryColor = nil
	return s
}

// SetSecondaryColor0 Sets the value of secondary_color_0 (20)
func (s *MainRegion) SetSecondaryColor0(value ColorRGBA) *MainRegion {
	copy := ColorRGBA(bytes.Clone(value))
	s.internal.SecondaryColor0 = &copy
	return s
}

// GetSecondaryColor0 Gets the value of secondary_color_0 (20)
func (s *MainRegion) GetSecondaryColor0() (ColorRGBA, bool) {
	if s.internal.SecondaryColor0 != nil {
		return *s.internal.SecondaryColor0, true
	}
	return ColorRGBA{0, 0, 0}, false
}

// ClearSecondaryColor0 Clears the value of secondary_color_0 (20)
func (s *MainRegion) ClearSecondaryColor0() *MainRegion {
	s.internal.SecondaryColor0 = nil
	return s
}

// SetSecondaryColor1 Sets the value of secondary_color_1 (21)
// See `secondary_color_0`.
func (s *MainRegion) SetSecondaryColor1(value ColorRGBA) *MainRegion {
	copy := ColorRGBA(bytes.Clone(value))
	s.internal.SecondaryColor1 = &copy
	return s
}

// GetSecondaryColor1 Gets the value of secondary_color_1 (21)
// See `secondary_color_0`.
func (s *MainRegion) GetSecondaryColor1() (ColorRGBA, bool) {
	if s.internal.SecondaryColor1 != nil {
		return *s.internal.SecondaryColor1, true
	}
	return ColorRGBA{0, 0, 0}, false
}

// ClearSecondaryColor1 Clears the value of secondary_color_1 (21)
// See `secondary_color_0`.
func (s *MainRegion) ClearSecondaryColor1() *MainRegion {
	s.internal.SecondaryColor1 = nil
	return s
}

// SetSecondaryColor2 Sets the value of secondary_color_2 (22)
// See `secondary_color_0`.
func (s *MainRegion) SetSecondaryColor2(value ColorRGBA) *MainRegion {
	copy := ColorRGBA(bytes.Clone(value))
	s.internal.SecondaryColor2 = &copy
	return s
}

// GetSecondaryColor2 Gets the value of secondary_color_2 (22)
// See `secondary_color_0`.
func (s *MainRegion) GetSecondaryColor2() (ColorRGBA, bool) {
	if s.internal.SecondaryColor2 != nil {
		return *s.internal.SecondaryColor2, true
	}
	return ColorRGBA{0, 0, 0}, false
}

// ClearSecondaryColor2 Clears the value of secondary_color_2 (22)
// See `secondary_color_0`.
func (s *MainRegion) ClearSecondaryColor2() *MainRegion {
	s.internal.SecondaryColor2 = nil
	return s
}

// SetSecondaryColor3 Sets the value of secondary_color_3 (23)
// See `secondary_color_0`.
func (s *MainRegion) SetSecondaryColor3(value ColorRGBA) *MainRegion {
	copy := ColorRGBA(bytes.Clone(value))
	s.internal.SecondaryColor3 = &copy
	return s
}

// GetSecondaryColor3 Gets the value of secondary_color_3 (23)
// See `secondary_color_0`.
func (s *MainRegion) GetSecondaryColor3() (ColorRGBA, bool) {
	if s.internal.SecondaryColor3 != nil {
		return *s.internal.SecondaryColor3, true
	}
	return ColorRGBA{0, 0, 0}, false
}

// ClearSecondaryColor3 Clears the value of secondary_color_3 (23)
// See `secondary_color_0`.
func (s *MainRegion) ClearSecondaryColor3() *MainRegion {
	s.internal.SecondaryColor3 = nil
	return s
}

// SetSecondaryColor4 Sets the value of secondary_color_4 (24)
// See `secondary_color_0`.
func (s *MainRegion) SetSecondaryColor4(value ColorRGBA) *MainRegion {
	copy := ColorRGBA(bytes.Clone(value))
	s.internal.SecondaryColor4 = &copy
	return s
}

// GetSecondaryColor4 Gets the value of secondary_color_4 (24)
// See `secondary_color_0`.
func (s *MainRegion) GetSecondaryColor4() (ColorRGBA, bool) {
	if s.internal.SecondaryColor4 != nil {
		return *s.internal.SecondaryColor4, true
	}
	return ColorRGBA{0, 0, 0}, false
}

// ClearSecondaryColor4 Clears the value of secondary_color_4 (24)
// See `secondary_color_0`.
func (s *MainRegion) ClearSecondaryColor4() *MainRegion {
	s.internal.SecondaryColor4 = nil
	return s
}

// SetTransmissionDistance Sets the value of transmission_distance (27)
func (s *MainRegion) SetTransmissionDistance(value float64) *MainRegion {
	s.internal.TransmissionDistance = &value
	return s
}

// GetTransmissionDistance Gets the value of transmission_distance (27)
func (s *MainRegion) GetTransmissionDistance() (float64, bool) {
	if s.internal.TransmissionDistance != nil {
		return *s.internal.TransmissionDistance, true
	}
	return 0.0, false
}

// ClearTransmissionDistance Clears the value of transmission_distance (27)
func (s *MainRegion) ClearTransmissionDistance() *MainRegion {
	s.internal.TransmissionDistance = nil
	return s
}

// SetTags Sets the value of tags (28)
// Properties of the material. Can have multiple tags at once.
func (s *MainRegion) SetTags(value []Tag) *MainRegion {
	s.internal.Tags = &value
	return s
}

// GetTags Gets the value of tags (28)
// Properties of the material. Can have multiple tags at once.
func (s *MainRegion) GetTags() ([]Tag, bool) {
	if s.internal.Tags != nil {
		value := make([]Tag, len(*s.internal.Tags))
		for idx, item := range *s.internal.Tags {
			value[idx] = Tag(item)
		}
		return value, true
	}
	return []Tag{}, false
}

// ClearTags Clears the value of tags (28)
// Properties of the material. Can have multiple tags at once.
func (s *MainRegion) ClearTags() *MainRegion {
	s.internal.Tags = nil
	return s
}

// SetCertifications Sets the value of certifications (56)
// Certifications the material has.
func (s *MainRegion) SetCertifications(value []MaterialCertifications) *MainRegion {
	s.internal.Certifications = &value
	return s
}

// GetCertifications Gets the value of certifications (56)
// Certifications the material has.
func (s *MainRegion) GetCertifications() ([]MaterialCertifications, bool) {
	if s.internal.Certifications != nil {
		value := make([]MaterialCertifications, len(*s.internal.Certifications))
		for idx, item := range *s.internal.Certifications {
			value[idx] = MaterialCertifications(item)
		}
		return value, true
	}
	return []MaterialCertifications{}, false
}

// ClearCertifications Clears the value of certifications (56)
// Certifications the material has.
func (s *MainRegion) ClearCertifications() *MainRegion {
	s.internal.Certifications = nil
	return s
}

// SetDensity Sets the value of density (29)
// Density of the material.
func (s *MainRegion) SetDensity(value float64) *MainRegion {
	s.internal.Density = &value
	return s
}

// GetDensity Gets the value of density (29)
// Density of the material.
func (s *MainRegion) GetDensity() (float64, bool) {
	if s.internal.Density != nil {
		return *s.internal.Density, true
	}
	return 0.0, false
}

// ClearDensity Clears the value of density (29)
// Density of the material.
func (s *MainRegion) ClearDensity() *MainRegion {
	s.internal.Density = nil
	return s
}

// SetFilamentDiameter Sets the value of filament_diameter (30)
func (s *MainRegion) SetFilamentDiameter(value float64) *MainRegion {
	s.internal.FilamentDiameter = &value
	return s
}

// GetFilamentDiameter Gets the value of filament_diameter (30)
func (s *MainRegion) GetFilamentDiameter() (float64, bool) {
	if s.internal.FilamentDiameter != nil {
		return *s.internal.FilamentDiameter, true
	}
	return 0.0, false
}

// ClearFilamentDiameter Clears the value of filament_diameter (30)
func (s *MainRegion) ClearFilamentDiameter() *MainRegion {
	s.internal.FilamentDiameter = nil
	return s
}

// SetShoreHardnessA Sets the value of shore_hardness_a (31)
func (s *MainRegion) SetShoreHardnessA(value int) *MainRegion {
	s.internal.ShoreHardnessA = &value
	return s
}

// GetShoreHardnessA Gets the value of shore_hardness_a (31)
func (s *MainRegion) GetShoreHardnessA() (int, bool) {
	if s.internal.ShoreHardnessA != nil {
		return *s.internal.ShoreHardnessA, true
	}
	return 0, false
}

// ClearShoreHardnessA Clears the value of shore_hardness_a (31)
func (s *MainRegion) ClearShoreHardnessA() *MainRegion {
	s.internal.ShoreHardnessA = nil
	return s
}

// SetShoreHardnessD Sets the value of shore_hardness_d (32)
func (s *MainRegion) SetShoreHardnessD(value int) *MainRegion {
	s.internal.ShoreHardnessD = &value
	return s
}

// GetShoreHardnessD Gets the value of shore_hardness_d (32)
func (s *MainRegion) GetShoreHardnessD() (int, bool) {
	if s.internal.ShoreHardnessD != nil {
		return *s.internal.ShoreHardnessD, true
	}
	return 0, false
}

// ClearShoreHardnessD Clears the value of shore_hardness_d (32)
func (s *MainRegion) ClearShoreHardnessD() *MainRegion {
	s.internal.ShoreHardnessD = nil
	return s
}

// SetMinNozzleDiameter Sets the value of min_nozzle_diameter (33)
func (s *MainRegion) SetMinNozzleDiameter(value float64) *MainRegion {
	s.internal.MinNozzleDiameter = &value
	return s
}

// GetMinNozzleDiameter Gets the value of min_nozzle_diameter (33)
func (s *MainRegion) GetMinNozzleDiameter() (float64, bool) {
	if s.internal.MinNozzleDiameter != nil {
		return *s.internal.MinNozzleDiameter, true
	}
	return 0.0, false
}

// ClearMinNozzleDiameter Clears the value of min_nozzle_diameter (33)
func (s *MainRegion) ClearMinNozzleDiameter() *MainRegion {
	s.internal.MinNozzleDiameter = nil
	return s
}

// SetMinPrintTemperature Sets the value of min_print_temperature (34)
func (s *MainRegion) SetMinPrintTemperature(value int) *MainRegion {
	s.internal.MinPrintTemperature = &value
	return s
}

// GetMinPrintTemperature Gets the value of min_print_temperature (34)
func (s *MainRegion) GetMinPrintTemperature() (int, bool) {
	if s.internal.MinPrintTemperature != nil {
		return *s.internal.MinPrintTemperature, true
	}
	return 0, false
}

// ClearMinPrintTemperature Clears the value of min_print_temperature (34)
func (s *MainRegion) ClearMinPrintTemperature() *MainRegion {
	s.internal.MinPrintTemperature = nil
	return s
}

// SetMaxPrintTemperature Sets the value of max_print_temperature (35)
func (s *MainRegion) SetMaxPrintTemperature(value int) *MainRegion {
	s.internal.MaxPrintTemperature = &value
	return s
}

// GetMaxPrintTemperature Gets the value of max_print_temperature (35)
func (s *MainRegion) GetMaxPrintTemperature() (int, bool) {
	if s.internal.MaxPrintTemperature != nil {
		return *s.internal.MaxPrintTemperature, true
	}
	return 0, false
}

// ClearMaxPrintTemperature Clears the value of max_print_temperature (35)
func (s *MainRegion) ClearMaxPrintTemperature() *MainRegion {
	s.internal.MaxPrintTemperature = nil
	return s
}

// SetPreheatTemperature Sets the value of preheat_temperature (36)
func (s *MainRegion) SetPreheatTemperature(value int) *MainRegion {
	s.internal.PreheatTemperature = &value
	return s
}

// GetPreheatTemperature Gets the value of preheat_temperature (36)
func (s *MainRegion) GetPreheatTemperature() (int, bool) {
	if s.internal.PreheatTemperature != nil {
		return *s.internal.PreheatTemperature, true
	}
	return 0, false
}

// ClearPreheatTemperature Clears the value of preheat_temperature (36)
func (s *MainRegion) ClearPreheatTemperature() *MainRegion {
	s.internal.PreheatTemperature = nil
	return s
}

// SetMinBedTemperature Sets the value of min_bed_temperature (37)
func (s *MainRegion) SetMinBedTemperature(value int) *MainRegion {
	s.internal.MinBedTemperature = &value
	return s
}

// GetMinBedTemperature Gets the value of min_bed_temperature (37)
func (s *MainRegion) GetMinBedTemperature() (int, bool) {
	if s.internal.MinBedTemperature != nil {
		return *s.internal.MinBedTemperature, true
	}
	return 0, false
}

// ClearMinBedTemperature Clears the value of min_bed_temperature (37)
func (s *MainRegion) ClearMinBedTemperature() *MainRegion {
	s.internal.MinBedTemperature = nil
	return s
}

// SetMaxBedTemperature Sets the value of max_bed_temperature (38)
func (s *MainRegion) SetMaxBedTemperature(value int) *MainRegion {
	s.internal.MaxBedTemperature = &value
	return s
}

// GetMaxBedTemperature Gets the value of max_bed_temperature (38)
func (s *MainRegion) GetMaxBedTemperature() (int, bool) {
	if s.internal.MaxBedTemperature != nil {
		return *s.internal.MaxBedTemperature, true
	}
	return 0, false
}

// ClearMaxBedTemperature Clears the value of max_bed_temperature (38)
func (s *MainRegion) ClearMaxBedTemperature() *MainRegion {
	s.internal.MaxBedTemperature = nil
	return s
}

// SetMinChamberTemperature Sets the value of min_chamber_temperature (39)
func (s *MainRegion) SetMinChamberTemperature(value int) *MainRegion {
	s.internal.MinChamberTemperature = &value
	return s
}

// GetMinChamberTemperature Gets the value of min_chamber_temperature (39)
func (s *MainRegion) GetMinChamberTemperature() (int, bool) {
	if s.internal.MinChamberTemperature != nil {
		return *s.internal.MinChamberTemperature, true
	}
	return 0, false
}

// ClearMinChamberTemperature Clears the value of min_chamber_temperature (39)
func (s *MainRegion) ClearMinChamberTemperature() *MainRegion {
	s.internal.MinChamberTemperature = nil
	return s
}

// SetMaxChamberTemperature Sets the value of max_chamber_temperature (40)
func (s *MainRegion) SetMaxChamberTemperature(value int) *MainRegion {
	s.internal.MaxChamberTemperature = &value
	return s
}

// GetMaxChamberTemperature Gets the value of max_chamber_temperature (40)
func (s *MainRegion) GetMaxChamberTemperature() (int, bool) {
	if s.internal.MaxChamberTemperature != nil {
		return *s.internal.MaxChamberTemperature, true
	}
	return 0, false
}

// ClearMaxChamberTemperature Clears the value of max_chamber_temperature (40)
func (s *MainRegion) ClearMaxChamberTemperature() *MainRegion {
	s.internal.MaxChamberTemperature = nil
	return s
}

// SetChamberTemperature Sets the value of chamber_temperature (41)
func (s *MainRegion) SetChamberTemperature(value int) *MainRegion {
	s.internal.ChamberTemperature = &value
	return s
}

// GetChamberTemperature Gets the value of chamber_temperature (41)
func (s *MainRegion) GetChamberTemperature() (int, bool) {
	if s.internal.ChamberTemperature != nil {
		return *s.internal.ChamberTemperature, true
	}
	return 0, false
}

// ClearChamberTemperature Clears the value of chamber_temperature (41)
func (s *MainRegion) ClearChamberTemperature() *MainRegion {
	s.internal.ChamberTemperature = nil
	return s
}

// SetContainerWidth Sets the value of container_width (42)
func (s *MainRegion) SetContainerWidth(value int) *MainRegion {
	s.internal.ContainerWidth = &value
	return s
}

// GetContainerWidth Gets the value of container_width (42)
func (s *MainRegion) GetContainerWidth() (int, bool) {
	if s.internal.ContainerWidth != nil {
		return *s.internal.ContainerWidth, true
	}
	return 0, false
}

// ClearContainerWidth Clears the value of container_width (42)
func (s *MainRegion) ClearContainerWidth() *MainRegion {
	s.internal.ContainerWidth = nil
	return s
}

// SetContainerOuterDiameter Sets the value of container_outer_diameter (43)
func (s *MainRegion) SetContainerOuterDiameter(value int) *MainRegion {
	s.internal.ContainerOuterDiameter = &value
	return s
}

// GetContainerOuterDiameter Gets the value of container_outer_diameter (43)
func (s *MainRegion) GetContainerOuterDiameter() (int, bool) {
	if s.internal.ContainerOuterDiameter != nil {
		return *s.internal.ContainerOuterDiameter, true
	}
	return 0, false
}

// ClearContainerOuterDiameter Clears the value of container_outer_diameter (43)
func (s *MainRegion) ClearContainerOuterDiameter() *MainRegion {
	s.internal.ContainerOuterDiameter = nil
	return s
}

// SetContainerInnerDiameter Sets the value of container_inner_diameter (44)
func (s *MainRegion) SetContainerInnerDiameter(value int) *MainRegion {
	s.internal.ContainerInnerDiameter = &value
	return s
}

// GetContainerInnerDiameter Gets the value of container_inner_diameter (44)
func (s *MainRegion) GetContainerInnerDiameter() (int, bool) {
	if s.internal.ContainerInnerDiameter != nil {
		return *s.internal.ContainerInnerDiameter, true
	}
	return 0, false
}

// ClearContainerInnerDiameter Clears the value of container_inner_diameter (44)
func (s *MainRegion) ClearContainerInnerDiameter() *MainRegion {
	s.internal.ContainerInnerDiameter = nil
	return s
}

// SetContainerHoleDiameter Sets the value of container_hole_diameter (45)
func (s *MainRegion) SetContainerHoleDiameter(value int) *MainRegion {
	s.internal.ContainerHoleDiameter = &value
	return s
}

// GetContainerHoleDiameter Gets the value of container_hole_diameter (45)
func (s *MainRegion) GetContainerHoleDiameter() (int, bool) {
	if s.internal.ContainerHoleDiameter != nil {
		return *s.internal.ContainerHoleDiameter, true
	}
	return 0, false
}

// ClearContainerHoleDiameter Clears the value of container_hole_diameter (45)
func (s *MainRegion) ClearContainerHoleDiameter() *MainRegion {
	s.internal.ContainerHoleDiameter = nil
	return s
}

// SetViscosity18C Sets the value of viscosity_18c (46)
// Viscosity of the material at 18 °C.
func (s *MainRegion) SetViscosity18C(value float64) *MainRegion {
	s.internal.Viscosity18C = &value
	return s
}

// GetViscosity18C Gets the value of viscosity_18c (46)
// Viscosity of the material at 18 °C.
func (s *MainRegion) GetViscosity18C() (float64, bool) {
	if s.internal.Viscosity18C != nil {
		return *s.internal.Viscosity18C, true
	}
	return 0.0, false
}

// ClearViscosity18C Clears the value of viscosity_18c (46)
// Viscosity of the material at 18 °C.
func (s *MainRegion) ClearViscosity18C() *MainRegion {
	s.internal.Viscosity18C = nil
	return s
}

// SetViscosity25C Sets the value of viscosity_25c (47)
// Viscosity of the material at 25 °C.
func (s *MainRegion) SetViscosity25C(value float64) *MainRegion {
	s.internal.Viscosity25C = &value
	return s
}

// GetViscosity25C Gets the value of viscosity_25c (47)
// Viscosity of the material at 25 °C.
func (s *MainRegion) GetViscosity25C() (float64, bool) {
	if s.internal.Viscosity25C != nil {
		return *s.internal.Viscosity25C, true
	}
	return 0.0, false
}

// ClearViscosity25C Clears the value of viscosity_25c (47)
// Viscosity of the material at 25 °C.
func (s *MainRegion) ClearViscosity25C() *MainRegion {
	s.internal.Viscosity25C = nil
	return s
}

// SetViscosity40C Sets the value of viscosity_40c (48)
// Viscosity of the material at 40 °C.
func (s *MainRegion) SetViscosity40C(value float64) *MainRegion {
	s.internal.Viscosity40C = &value
	return s
}

// GetViscosity40C Gets the value of viscosity_40c (48)
// Viscosity of the material at 40 °C.
func (s *MainRegion) GetViscosity40C() (float64, bool) {
	if s.internal.Viscosity40C != nil {
		return *s.internal.Viscosity40C, true
	}
	return 0.0, false
}

// ClearViscosity40C Clears the value of viscosity_40c (48)
// Viscosity of the material at 40 °C.
func (s *MainRegion) ClearViscosity40C() *MainRegion {
	s.internal.Viscosity40C = nil
	return s
}

// SetViscosity60C Sets the value of viscosity_60c (49)
// Viscosity of the material at 60 °C.
func (s *MainRegion) SetViscosity60C(value float64) *MainRegion {
	s.internal.Viscosity60C = &value
	return s
}

// GetViscosity60C Gets the value of viscosity_60c (49)
// Viscosity of the material at 60 °C.
func (s *MainRegion) GetViscosity60C() (float64, bool) {
	if s.internal.Viscosity60C != nil {
		return *s.internal.Viscosity60C, true
	}
	return 0.0, false
}

// ClearViscosity60C Clears the value of viscosity_60c (49)
// Viscosity of the material at 60 °C.
func (s *MainRegion) ClearViscosity60C() *MainRegion {
	s.internal.Viscosity60C = nil
	return s
}

// SetContainerVolumetricCapacity Sets the value of container_volumetric_capacity (50)
// Maximum amount of material the container can hold.
func (s *MainRegion) SetContainerVolumetricCapacity(value float64) *MainRegion {
	s.internal.ContainerVolumetricCapacity = &value
	return s
}

// GetContainerVolumetricCapacity Gets the value of container_volumetric_capacity (50)
// Maximum amount of material the container can hold.
func (s *MainRegion) GetContainerVolumetricCapacity() (float64, bool) {
	if s.internal.ContainerVolumetricCapacity != nil {
		return *s.internal.ContainerVolumetricCapacity, true
	}
	return 0.0, false
}

// ClearContainerVolumetricCapacity Clears the value of container_volumetric_capacity (50)
// Maximum amount of material the container can hold.
func (s *MainRegion) ClearContainerVolumetricCapacity() *MainRegion {
	s.internal.ContainerVolumetricCapacity = nil
	return s
}

// SetCureWavelength Sets the value of cure_wavelength (51)
func (s *MainRegion) SetCureWavelength(value int) *MainRegion {
	s.internal.CureWavelength = &value
	return s
}

// GetCureWavelength Gets the value of cure_wavelength (51)
func (s *MainRegion) GetCureWavelength() (int, bool) {
	if s.internal.CureWavelength != nil {
		return *s.internal.CureWavelength, true
	}
	return 0, false
}

// ClearCureWavelength Clears the value of cure_wavelength (51)
func (s *MainRegion) ClearCureWavelength() *MainRegion {
	s.internal.CureWavelength = nil
	return s
}

// SetDryingTemperature Sets the value of drying_temperature (57)
func (s *MainRegion) SetDryingTemperature(value int) *MainRegion {
	s.internal.DryingTemperature = &value
	return s
}

// GetDryingTemperature Gets the value of drying_temperature (57)
func (s *MainRegion) GetDryingTemperature() (int, bool) {
	if s.internal.DryingTemperature != nil {
		return *s.internal.DryingTemperature, true
	}
	return 0, false
}

// ClearDryingTemperature Clears the value of drying_temperature (57)
func (s *MainRegion) ClearDryingTemperature() *MainRegion {
	s.internal.DryingTemperature = nil
	return s
}

// SetDryingTime Sets the value of drying_time (58)
func (s *MainRegion) SetDryingTime(value int) *MainRegion {
	s.internal.DryingTime = &value
	return s
}

// GetDryingTime Gets the value of drying_time (58)
func (s *MainRegion) GetDryingTime() (int, bool) {
	if s.internal.DryingTime != nil {
		return *s.internal.DryingTime, true
	}
	return 0, false
}

// ClearDryingTime Clears the value of drying_time (58)
func (s *MainRegion) ClearDryingTime() *MainRegion {
	s.internal.DryingTime = nil
	return s
}

// GetUnknownFields returns a map of all unknown fields
// For the aux region, this will also include all vendor specific fields
func (s *MainRegion) GetUnknownFields() map[any]any {
	if s.internal.Unknowns == nil {
		s.internal.Unknowns = make(map[any]any)
	}
	return s.internal.Unknowns
}

func (s MainRegion) getInternal() any {
	return &s.internal
}

func (s MainRegion) getRegionName() string {
	return "main"
}

// RegionOptions accesses encoding options for this region
func (s MainRegion) RegionOptions() *RegionOptions {
	return s.regionOptions
}
