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
package config

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"gopkg.in/yaml.v3"

	"github.com/iancoleman/strcase"
)

// fieldYAML is the internal decoder for fields
type fieldYAML struct {
	Key              int    `yaml:"key"`
	Name             string `yaml:"name"`
	Type             string `yaml:"type"`
	Unit             string `yaml:"unit"`
	Description      any    `yaml:"description"`
	Category         string `yaml:"category"`
	ItemsFile        string `yaml:"items_file"`
	NameField        string `yaml:"name_field"`
	DisplayNameField string `yaml:"display_name_field"`
	Required         string `yaml:"required"`
	Example          string `yaml:"example"`
	MaxLength        int    `yaml:"max_length"`
	Deprecated       bool   `yaml:"deprecated"`
}

// Field is the public representation of a field
type Field struct {
	yaml             fieldYAML
	enumeratedValues []Enumeration
}

// Key returns the field key
func (f Field) Key() int {
	return f.yaml.Key
}

// Name returns the field name
func (f Field) Name() string {
	return f.yaml.Name
}

// Type returns the field type
func (f Field) Type() string {
	return f.yaml.Type
}

// Unit returns the field unit
func (f Field) Unit() string {
	return f.yaml.Unit
}

// Description returns the field description
func (f Field) Description() string {
	if str, ok := f.yaml.Description.(string); ok {
		return str
	}
	if arr, ok := f.yaml.Description.([]string); ok {
		return strings.Join(arr, " ")
	}
	return ""
}

// Category returns the field category
func (f Field) Category() string {
	return f.yaml.Category
}

func (f Field) NameField() string {
	return f.yaml.NameField
}

func (f Field) DisplayNameField() string {
	return f.yaml.DisplayNameField
}

func (f Field) Required() string {
	return f.yaml.Required
}

func (f Field) Example() string {
	return f.yaml.Example
}

func (f Field) MaxLength() int {
	return f.yaml.MaxLength
}

func (f Field) HasEnumeration() bool {
	return f.yaml.ItemsFile != ""
}

func (f Field) EnumeratedValues() []Enumeration {
	return f.enumeratedValues
}

func (f Field) IsDeprecated() bool {
	return f.yaml.Deprecated
}

func (f Field) EnumItemsFile() string {
	return f.yaml.ItemsFile
}

func (f Field) String() string {
	return fmt.Sprintf(`Field: %s (%d)
Type: %s
Unit: %s
Description: %s
Category: %s
NameField: %s
DisplayNameField: %s
Required: %s
Exmaple: %s
MaxLength: %d
HasEnumeration: %t
Deprecated: %t`, f.Name(), f.Key(), f.Type(), f.Unit(), f.Description(),
		f.Category(), f.NameField(), f.DisplayNameField(), f.Required(), f.Example(),
		f.MaxLength(), f.HasEnumeration(), f.IsDeprecated())
}

// loadFields loads fields from the specified yaml file
func loadFields(filename string) ([]Field, error) {
	f, err := os.Open(dataFileName(filename))
	if err != nil {
		return nil, fmt.Errorf("failed to load fields from %s: %w", filename, err)
	}
	defer f.Close()
	var fields []fieldYAML
	err = yaml.NewDecoder(f).Decode(&fields)
	if err != nil {
		return nil, fmt.Errorf("failed to decode fields from %s: %w", filename, err)
	}

	finalized := make([]Field, len(fields))
	for idx, item := range fields {
		finalized[idx] = Field{yaml: item}

		if item.Name == "gtin" {
			// gtin is problematic. It's a number but needs to be strictly treated as uint64
			finalized[idx].yaml.Type = "uint64"
		}

		if item.ItemsFile != "" {
			finalized[idx].enumeratedValues, err = loadEnumerations(item.ItemsFile)
			if err != nil {
				return nil, err
			}
		}
	}

	return finalized, nil
}

// GetInternalTypeAndImports returns the internal type that we will use for the
// specified schema type, and a list of any imports required to use that type
func (f Field) GetInternalTypeAndImports() (internalType string, requiredImports []string) {
	switch f.Type() {
	case "uint64":
		return "uint64", nil
	case "color_rgba":
		return "ColorRGBA", []string{"bytes"}
	case "string":
		return "string", nil
	case "enum":
		return f.GetInternalEnumType(), nil
	case "enum_array":
		return "[]" + f.GetInternalEnumType(), nil
	case "timestamp":
		return "uint64", []string{"time"}
	case "uuid":
		return "uuid.UUID", []string{"github.com/google/uuid"}
	case "int":
		return "int", nil
	case "number":
		return "float64", nil
	default:
		panic(fmt.Sprintf("unknown type: -%s-", f.Type()))
	}
}

// GetInternalFieldName returns the internal field name that we will use for this field
// it is derived from refactoring the snake_case native field name using golang standard
// upper camel case
func (f Field) GetInternalFieldName() string {
	return strcase.ToCamel(f.yaml.Name)
}

// GetInternalEnumType will derive a camel case enum type
// from the enum_items_file entry in the field
func (f Field) GetInternalEnumType() string {
	if f.yaml.ItemsFile == "" {
		panic(fmt.Sprintf("Field %s has no items_file, so cannot generate an enum type", f.Name()))
	}

	// Strip the _enum.yaml
	base := strings.ReplaceAll(f.yaml.ItemsFile, "_enum.yaml", "")

	// Modify what is left, for known problem cases
	switch base {
	case "tags":
		// Much better if tags enumeration is actually "tag" since the enumeration references individual tags
		base = "tag"
	}

	return strcase.ToCamel(base)
}

// GetInternalEnumMapName will return a suitable name for our internal enum map
// which is used for YAML encode/decode
func (f Field) GetInternalEnumMapName() string {
	base := []rune(f.GetInternalEnumType())
	return string(append([]rune{unicode.ToUpper(base[0])}, base[1:]...)) + "Map"
}

func (f Field) GetGeneratedEnumFilename() string {
	if f.yaml.ItemsFile == "" {
		panic(fmt.Sprintf("Field %s has no items_file, so cannot generate an enum filename", f.Name()))
	}

	// Strip the _enum.yaml
	base := strings.ReplaceAll(f.yaml.ItemsFile, "_enum.yaml", "")

	// Modify what is left, for known problem cases
	switch base {
	case "tags":
		// Much better if tags enumeration is actually "tag" since the enumeration references individual tags
		base = "tag"
	}

	return strcase.ToSnake(base)

}
