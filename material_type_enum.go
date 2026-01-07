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

type MaterialType uint64

const (
	// MaterialTypePLA
	// Polylactic Acid
	// Easy-to-print, biodegradable material. Ideal for beginners, prototypes, and models.
	MaterialTypePLA MaterialType = 0

	// MaterialTypePETG
	// Polyethylene Terephthalate Glycol
	// Durable, strong, and temperature-resistant. Great for mechanical parts and functional prints.
	MaterialTypePETG MaterialType = 1

	// MaterialTypeTPU
	// Thermoplastic Polyurethane
	// A flexible, rubber-like material. Used for phone cases, vibration dampeners, and other soft parts.
	MaterialTypeTPU MaterialType = 2

	// MaterialTypeABS
	// Acrylonitrile Butadiene Styrene
	// Strong, durable, and heat-resistant plastic. Used for functional parts like car interiors and LEGOs. Requires a heated bed and enclosure.
	MaterialTypeABS MaterialType = 3

	// MaterialTypeASA
	// Acrylonitrile Styrene Acrylate
	// Similar to ABS but with high UV and weather resistance, making it perfect for outdoor applications.
	MaterialTypeASA MaterialType = 4

	// MaterialTypePC
	// Polycarbonate
	// Extremely strong, impact-resistant, and heat-resistant. Used for demanding engineering applications.
	MaterialTypePC MaterialType = 5

	// MaterialTypePCTG
	// Polycyclohexylenedimethylene Terephthalate Glycol
	// A tougher alternative to PETG with higher impact and chemical resistance.
	MaterialTypePCTG MaterialType = 6

	// MaterialTypePP
	// Polypropylene
	// Lightweight, chemically resistant, and flexible. Used for creating living hinges and durable containers.
	MaterialTypePP MaterialType = 7

	// MaterialTypePA6
	// Polyamide 6
	// A type of Nylon that is tough and wear-resistant but absorbs more moisture than other nylons.
	MaterialTypePA6 MaterialType = 8

	// MaterialTypePA11
	// Polyamide 11
	// A flexible, bio-based Nylon with low moisture absorption and good chemical resistance.
	MaterialTypePA11 MaterialType = 9

	// MaterialTypePA12
	// Polyamide 12
	// The most common Nylon for 3D printing. Strong, tough, with low moisture absorption. Great for functional parts.
	MaterialTypePA12 MaterialType = 10

	// MaterialTypePA66
	// Polyamide 66
	// A stiffer and more heat-resistant Nylon compared to PA6, used for durable mechanical parts.
	MaterialTypePA66 MaterialType = 11

	// MaterialTypeCPE
	// Copolyester
	// A family of strong and dimensionally stable materials (including PETG) known for chemical resistance.
	MaterialTypeCPE MaterialType = 12

	// MaterialTypeTPE
	// Thermoplastic Elastomer
	// A general class of soft, rubbery materials. Softer and more flexible than TPU.
	MaterialTypeTPE MaterialType = 13

	// MaterialTypeHIPS
	// High Impact Polystyrene
	// A lightweight material often used as a dissolvable support material for ABS prints (dissolves in Limonene).
	MaterialTypeHIPS MaterialType = 14

	// MaterialTypePHA
	// Polyhydroxyalkanoate
	// A biodegradable material similar to PLA but with better toughness and flexibility.
	MaterialTypePHA MaterialType = 15

	// MaterialTypePET
	// Polyethylene Terephthalate
	// The same plastic used in water bottles. Strong and food-safe, but less common for printing than PETG.
	MaterialTypePET MaterialType = 16

	// MaterialTypePEI
	// Polyetherimide
	// A high-performance material (also known as Ultem) with excellent thermal and mechanical properties.
	MaterialTypePEI MaterialType = 17

	// MaterialTypePBT
	// Polybutylene Terephthalate
	// An engineering polymer with good heat resistance and electrical insulation properties.
	MaterialTypePBT MaterialType = 18

	// MaterialTypePVB
	// Polyvinyl Butyral
	// Easy to print and can be chemically smoothed with isopropyl alcohol for a glossy finish.
	MaterialTypePVB MaterialType = 19

	// MaterialTypePVA
	// Polyvinyl Alcohol
	// A water-soluble filament used exclusively as a support material for complex prints.
	MaterialTypePVA MaterialType = 20

	// MaterialTypePEKK
	// Polyetherketoneketone
	// An ultra-high-performance polymer with exceptional heat, chemical, and mechanical properties for industrial use.
	MaterialTypePEKK MaterialType = 21

	// MaterialTypePEEK
	// Polyether Ether Ketone
	// An ultra-high-performance polymer with exceptional mechanical, thermal, and chemical resistance. Used in demanding aerospace, medical, and industrial applications.
	MaterialTypePEEK MaterialType = 22

	// MaterialTypeBVOH
	// Butenediol Vinyl Alcohol Copolymer
	// A water-soluble support material that often dissolves faster and is easier to print than PVA.
	MaterialTypeBVOH MaterialType = 23

	// MaterialTypeTPC
	// Thermoplastic Copolyester
	// A flexible, TPE-like material with good thermal and chemical resistance.
	MaterialTypeTPC MaterialType = 24

	// MaterialTypePPS
	// Polyphenylene Sulfide
	// A high-performance polymer known for its thermal stability and chemical resistance, often used in automotive and electronics.
	MaterialTypePPS MaterialType = 25

	// MaterialTypePPSU
	// Polyphenylsulfone
	// A high-performance material with excellent heat and chemical resistance, often used in medical applications.
	MaterialTypePPSU MaterialType = 26

	// MaterialTypePVC
	// Polyvinyl Chloride
	// Strong and durable but rarely used in 3D printing due to the release of toxic fumes.
	MaterialTypePVC MaterialType = 27

	// MaterialTypePEBA
	// Polyether Block Amide
	// A flexible and lightweight TPE known for its excellent energy return, used in sports equipment.
	MaterialTypePEBA MaterialType = 28

	// MaterialTypePVDF
	// Polyvinylidene Fluoride
	// High-performance polymer with excellent resistance to chemicals and UV light.
	MaterialTypePVDF MaterialType = 29

	// MaterialTypePPA
	// Polyphthalamide
	// A high-performance Nylon with superior strength, stiffness, and heat resistance compared to standard Nylons.
	MaterialTypePPA MaterialType = 30

	// MaterialTypePCL
	// Polycaprolactone
	// A biodegradable polyester with a very low melting point (~60 °C), allowing it to be reshaped by hand in hot water.
	MaterialTypePCL MaterialType = 31

	// MaterialTypePES
	// Polyethersulfone
	// A high-temperature, amorphous polymer with good chemical and hydrolytic stability.
	MaterialTypePES MaterialType = 32

	// MaterialTypePMMA
	// Polymethyl Methacrylate
	// A rigid, transparent material also known as acrylic. Offers good optical clarity.
	MaterialTypePMMA MaterialType = 33

	// MaterialTypePOM
	// Polyoxymethylene
	// A low-friction, rigid material also known as Delrin. Excellent for gears, bearings, and moving parts.
	MaterialTypePOM MaterialType = 34

	// MaterialTypePPE
	// Polyphenylene Ether
	// An engineering thermoplastic with good temperature resistance and dimensional stability, often used in blends.
	MaterialTypePPE MaterialType = 35

	// MaterialTypePS
	// Polystyrene
	// A lightweight and brittle material. Not commonly used in its pure form for 3D printing.
	MaterialTypePS MaterialType = 36

	// MaterialTypePSU
	// Polysulfone
	// A high-temperature material with good thermal stability and chemical resistance.
	MaterialTypePSU MaterialType = 37

	// MaterialTypeTPI
	// Thermoplastic Polyimide
	// An ultra-high-performance polymer with one of the highest glass transition temperatures and excellent thermal stability.
	MaterialTypeTPI MaterialType = 38

	// MaterialTypeSBS
	// Styrene-Butadiene-Styrene
	// A flexible, rubber-like material (a type of TPE) known for good durability. It is relatively easy to print for a flexible filament.
	MaterialTypeSBS MaterialType = 39

	// MaterialTypeOBC
	// Olefin Block Copolymer
	// A lightweight flexible material that has good dimensional stability and is weather, UV, and chemical resistant.
	MaterialTypeOBC MaterialType = 40

	// MaterialTypeEVA
	// Ethylene Vinyl Acetate
	// A flexible, soft material with rubber-like properties, known for its toughness and resistance to UV radiation and stress cracking.
	MaterialTypeEVA MaterialType = 41
)

var MaterialTypeMap = map[uint64]string{
	0:  "PLA",
	1:  "PETG",
	2:  "TPU",
	3:  "ABS",
	4:  "ASA",
	5:  "PC",
	6:  "PCTG",
	7:  "PP",
	8:  "PA6",
	9:  "PA11",
	10: "PA12",
	11: "PA66",
	12: "CPE",
	13: "TPE",
	14: "HIPS",
	15: "PHA",
	16: "PET",
	17: "PEI",
	18: "PBT",
	19: "PVB",
	20: "PVA",
	21: "PEKK",
	22: "PEEK",
	23: "BVOH",
	24: "TPC",
	25: "PPS",
	26: "PPSU",
	27: "PVC",
	28: "PEBA",
	29: "PVDF",
	30: "PPA",
	31: "PCL",
	32: "PES",
	33: "PMMA",
	34: "POM",
	35: "PPE",
	36: "PS",
	37: "PSU",
	38: "TPI",
	39: "SBS",
	40: "OBC",
	41: "EVA",
}

func (e MaterialType) String() string {
	return MaterialTypeMap[uint64(e)]
}

func (e MaterialType) MarshalYAML() (any, error) {
	if str, ok := MaterialTypeMap[uint64(e)]; ok {
		return str, nil
	}
	return nil, fmt.Errorf("unknown enumeration: %d", e)
}

func (e *MaterialType) UnmarshalYAML(value *yaml.Node) error {
	var str string
	if err := value.Decode(&str); err != nil {
		return err
	}

	// Hardly efficient, but this is not critical here
	for key, name := range MaterialTypeMap {
		if name == str {
			*e = MaterialType(key)
			return nil
		}
	}
	return fmt.Errorf("unknown enumeration: %s", str)
}
