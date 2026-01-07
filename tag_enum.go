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
	"fmt"
	"gopkg.in/yaml.v3"
)

type Tag uint64

const (
	// TagFiltrationRecommended
	// Releases a higher concentration of unsafe particles/fumes during printing so a HEPA and carbon filter is strongly recommended.
	TagFiltrationRecommended Tag = 0

	// TagBiocompatible
	// Certified biocompatibility (does not cause harmful effects when in contact with the body).
	TagBiocompatible Tag = 1

	// TagHomeCompostable
	// Decomposes into natural elements in a home compost system at ambient temperatures.
	TagHomeCompostable Tag = 61

	// TagIndustriallyCompostable
	// Decomposes into natural elements under specific temperature and microbial conditions in commercial composting facilities.
	TagIndustriallyCompostable Tag = 62

	// TagBioBased
	// Predominantly made from renewable biological resources, like plants.
	TagBioBased Tag = 63

	// TagAntibacterial
	// Has antibacterial properties.
	TagAntibacterial Tag = 2

	// TagAirFiltering
	// Has air filtering properties (absorbs/filters harmful compounds/particles from the air).
	TagAirFiltering Tag = 3

	// TagAbrasive
	// The material is abrasive and requires an abrasive-resistant nozzle.
	TagAbrasive Tag = 4

	// TagFoaming
	// The material increases its volume during extrusion.
	TagFoaming Tag = 5

	// TagCastable
	TagCastable Tag = 67

	// TagSelfExtinguishing
	TagSelfExtinguishing Tag = 6

	// TagParamagnetic
	// The material has paramagnetic properties, meaning that it is (weakly) attracted to magnets.
	TagParamagnetic Tag = 7

	// TagRadiationShielding
	// Has radiation shielding properties.
	TagRadiationShielding Tag = 8

	// TagHighTemperature
	TagHighTemperature Tag = 9

	// TagHighSpeed
	TagHighSpeed Tag = 71

	// TagEsdSafe
	TagEsdSafe Tag = 10

	// TagConductive
	TagConductive Tag = 11

	// TagEmiShielding
	TagEmiShielding Tag = 70

	// TagBlend
	// The material is a blend of multiple polymers or a base polymer with significant additives that alter its properties and may require a specific print profile.
	TagBlend Tag = 12

	// TagWaterSoluble
	// Can be dissolved in water.
	TagWaterSoluble Tag = 13

	// TagIpaSoluble
	// Can be dissolved in IPA (isopropyl alcohol).
	TagIpaSoluble Tag = 14

	// TagLimoneneSoluble
	// Can be dissolved in limonene.
	TagLimoneneSoluble Tag = 15

	// TagLowOutgassing
	// Releases only minimal gas (and vapor) when placed in a vacuum.
	TagLowOutgassing Tag = 64

	// TagMatte
	TagMatte Tag = 16

	// TagSilk
	TagSilk Tag = 17

	// TagTranslucent
	TagTranslucent Tag = 19

	// TagTransparent
	TagTransparent Tag = 20

	// TagWithoutPigments
	TagWithoutPigments Tag = 65

	// TagIridescent
	TagIridescent Tag = 21

	// TagPearlescent
	TagPearlescent Tag = 22

	// TagGlitter
	TagGlitter Tag = 23

	// TagGlowInTheDark
	TagGlowInTheDark Tag = 24

	// TagNeon
	TagNeon Tag = 25

	// TagIlluminescentColorChange
	TagIlluminescentColorChange Tag = 26

	// TagTemperatureColorChange
	TagTemperatureColorChange Tag = 27

	// TagGradualColorChange
	TagGradualColorChange Tag = 28

	// TagCoextruded
	TagCoextruded Tag = 29

	// TagContainsCarbon
	TagContainsCarbon Tag = 30

	// TagContainsCarbonFiber
	TagContainsCarbonFiber Tag = 31

	// TagContainsCarbonNanoTubes
	TagContainsCarbonNanoTubes Tag = 32

	// TagContainsGraphene
	TagContainsGraphene Tag = 72

	// TagContainsGlass
	TagContainsGlass Tag = 33

	// TagContainsGlassFiber
	TagContainsGlassFiber Tag = 34

	// TagContainsKevlar
	TagContainsKevlar Tag = 35

	// TagContainsPtfe
	TagContainsPtfe Tag = 68

	// TagContainsStone
	TagContainsStone Tag = 36

	// TagContainsMagnetite
	TagContainsMagnetite Tag = 37

	// TagContainsOrganicMaterial
	// Contains organic material.
	TagContainsOrganicMaterial Tag = 38

	// TagContainsCork
	// Contains cork.
	TagContainsCork Tag = 39

	// TagContainsWax
	// Contains wax.
	TagContainsWax Tag = 40

	// TagContainsWood
	// Contains wood.
	TagContainsWood Tag = 41

	// TagContainsAlgae
	// Contains algae.
	TagContainsAlgae Tag = 66

	// TagContainsBamboo
	// Contains bamboo.
	TagContainsBamboo Tag = 42

	// TagContainsPine
	// Contains pine.
	TagContainsPine Tag = 43

	// TagContainsCeramic
	// Contains ceramic.
	TagContainsCeramic Tag = 44

	// TagContainsBoronCarbide
	// Contains boron carbide (useful for radiation shielding).
	TagContainsBoronCarbide Tag = 45

	// TagContainsMetal
	// Contains metal. Specific type of metal contained can be expressed by an other tag.
	TagContainsMetal Tag = 46

	// TagContainsBronze
	// Contains bronze.
	TagContainsBronze Tag = 47

	// TagContainsIron
	// Contains iron.
	TagContainsIron Tag = 48

	// TagContainsSteel
	// Contains steel.
	TagContainsSteel Tag = 49

	// TagContainsSilver
	// Contains silver (useful for antibacterial properties).
	TagContainsSilver Tag = 50

	// TagContainsCopper
	// Contains copper.
	TagContainsCopper Tag = 51

	// TagContainsAluminium
	// Contains aluminium.
	TagContainsAluminium Tag = 52

	// TagContainsBrass
	// Contains brass.
	TagContainsBrass Tag = 53

	// TagContainsTungsten
	// Contains Tungsten (useful for radiation shielding).
	TagContainsTungsten Tag = 54

	// TagImitatesWood
	// Imitates wood.
	TagImitatesWood Tag = 55

	// TagImitatesMetal
	// Imitates metal.
	TagImitatesMetal Tag = 56

	// TagImitatesMarble
	// Imitates marble.
	TagImitatesMarble Tag = 57

	// TagImitatesStone
	// Imitates stone.
	TagImitatesStone Tag = 58

	// TagLithophane
	// Specifically designed for lithophaning.
	TagLithophane Tag = 59

	// TagRecycled
	// Part of the material is recycled.
	TagRecycled Tag = 60

	// TagLimitedEdition
	// The material is a limited edition run.
	TagLimitedEdition Tag = 69
)

var TagMap = map[uint64]string{
	0:  "filtration_recommended",
	1:  "biocompatible",
	61: "home_compostable",
	62: "industrially_compostable",
	63: "bio_based",
	2:  "antibacterial",
	3:  "air_filtering",
	4:  "abrasive",
	5:  "foaming",
	67: "castable",
	6:  "self_extinguishing",
	7:  "paramagnetic",
	8:  "radiation_shielding",
	9:  "high_temperature",
	71: "high_speed",
	10: "esd_safe",
	11: "conductive",
	70: "emi_shielding",
	12: "blend",
	13: "water_soluble",
	14: "ipa_soluble",
	15: "limonene_soluble",
	64: "low_outgassing",
	16: "matte",
	17: "silk",
	19: "translucent",
	20: "transparent",
	65: "without_pigments",
	21: "iridescent",
	22: "pearlescent",
	23: "glitter",
	24: "glow_in_the_dark",
	25: "neon",
	26: "illuminescent_color_change",
	27: "temperature_color_change",
	28: "gradual_color_change",
	29: "coextruded",
	30: "contains_carbon",
	31: "contains_carbon_fiber",
	32: "contains_carbon_nano_tubes",
	72: "contains_graphene",
	33: "contains_glass",
	34: "contains_glass_fiber",
	35: "contains_kevlar",
	68: "contains_ptfe",
	36: "contains_stone",
	37: "contains_magnetite",
	38: "contains_organic_material",
	39: "contains_cork",
	40: "contains_wax",
	41: "contains_wood",
	66: "contains_algae",
	42: "contains_bamboo",
	43: "contains_pine",
	44: "contains_ceramic",
	45: "contains_boron_carbide",
	46: "contains_metal",
	47: "contains_bronze",
	48: "contains_iron",
	49: "contains_steel",
	50: "contains_silver",
	51: "contains_copper",
	52: "contains_aluminium",
	53: "contains_brass",
	54: "contains_tungsten",
	55: "imitates_wood",
	56: "imitates_metal",
	57: "imitates_marble",
	58: "imitates_stone",
	59: "lithophane",
	60: "recycled",
	69: "limited_edition",
}

func (e Tag) String() string {
	return TagMap[uint64(e)]
}

func (e Tag) MarshalYAML() (any, error) {
	if str, ok := TagMap[uint64(e)]; ok {
		return str, nil
	}
	return nil, fmt.Errorf("unknown enumeration: %d", e)
}

func (e *Tag) UnmarshalYAML(value *yaml.Node) error {
	var str string
	if err := value.Decode(&str); err != nil {
		return err
	}

	// Hardly efficient, but this is not critical here
	for key, name := range TagMap {
		if name == str {
			*e = Tag(key)
			return nil
		}
	}
	return fmt.Errorf("unknown enumeration: %s", str)
}
