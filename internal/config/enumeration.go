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
	"sync"

	"github.com/iancoleman/strcase"
	"gopkg.in/yaml.v3"
)

const tagCategoriesEnumFile = "tag_categories_enum.yaml"

// enumerationsYAML is the base decoder for all enum files
type enumerationYAML struct {
	Key          int    `yaml:"key"`
	Category     string `yaml:"category"`
	Name         string `yaml:"name"`
	DisplayName  string `yaml:"display_name"`
	Description  any    `yaml:"description"`
	Emoji        string `yaml:"emoji"`
	Deprecated   bool   `yaml:"deprecated"`
	Abbreviation string `yaml:"abbreviation"`
}

// Enumeration is the public representation of an enumerated set
// excluding tag categories
type Enumeration struct {
	yaml enumerationYAML
}

func (e Enumeration) String() string {
	return fmt.Sprintf(`Enumeration: %s (%d)
DisplayName: %s
Category: %s
Description: %s`, e.Name(), e.Key(), e.DisplayName(), e.Category(), e.Description())
}

// Key returns the key associated with an enumeration
func (e Enumeration) Key() int {
	return e.yaml.Key
}

func (e Enumeration) Abbreviation() string {
	return e.yaml.Abbreviation
}

// Category returns the category of an enumeration, currently only applicable for tags
func (e Enumeration) Category() string {
	return e.yaml.Category
}

// Name returns the name of an enum
func (e Enumeration) Name() string {
	return e.yaml.Name
}

// DisplayName returns the display name of an enum
func (e Enumeration) DisplayName() string {
	return e.yaml.DisplayName
}

func (e Enumeration) IsDeprecated() bool {
	return e.yaml.Deprecated
}

// Description returns the description of an enum
func (e Enumeration) Description() string {
	if str, ok := e.yaml.Description.(string); ok {
		return str
	}
	if arr, ok := e.yaml.Description.([]string); ok {
		return strings.Join(arr, " ")
	}
	return ""
}

func (e Enumeration) GetInternalEnumName(field Field) (string, bool) {
	var en string
	var isAbbrev = false
	if e.Abbreviation() != "" {
		en = strings.ToUpper(strcase.ToCamel(e.Abbreviation()))
		isAbbrev = true
	} else {
		en = strcase.ToCamel(e.Name())
	}

	// Special cases for a couple of things that need re-casing
	switch en {
	case "Fff", "Sla":
		en = strings.ToUpper(en)
	}
	enumName := field.GetInternalEnumType() + en
	return enumName, isAbbrev
}

// TagCategoriesEnumeration enumerates tag categories
type TagCategoriesEnumeration struct {
	yaml enumerationYAML
}

// Name returns the tag category name
func (e TagCategoriesEnumeration) Name() string {
	return e.yaml.Name
}

// DisplayName returns the tag category display name
func (e TagCategoriesEnumeration) DisplayName() string {
	return e.yaml.DisplayName
}

// Emoji returns the tag category emoji
func (e TagCategoriesEnumeration) Emoji() string {
	return e.yaml.Emoji
}

func (e TagCategoriesEnumeration) String() string {
	return fmt.Sprintf("%s (%s): %s", e.Name(), e.Emoji(), e.yaml.DisplayName)
}

func (e TagCategoriesEnumeration) GetInternalEnumName(field Field) (string, bool) {
	return field.GetInternalEnumType() + strcase.ToCamel(e.Name()), false
}

// Because we only need to load tag category enumerations once
var categoriesEnumerationSingleton []TagCategoriesEnumeration
var categoriesEnumerationOnce sync.Once

// loadEnumerations will load (non tag category) enumerations from filename
func loadEnumerations(filename string) ([]Enumeration, error) {
	f, err := os.Open(dataFileName(filename))
	if err != nil {
		return nil, fmt.Errorf("failed to load enumerations from %s: %w", filename, err)
	}
	defer f.Close()
	var enums []enumerationYAML
	err = yaml.NewDecoder(f).Decode(&enums)
	if err != nil {
		return nil, fmt.Errorf("failed to decode enumerations from %s: %w", filename, err)
	}

	finalized := make([]Enumeration, len(enums))

	for idx, item := range enums {
		finalized[idx] = Enumeration{yaml: item}
	}

	return finalized, nil
}

// TagCategories reutrns all possible tag categories
func TagCategories() []TagCategoriesEnumeration {
	categoriesEnumerationOnce.Do(func() {
		f, err := os.Open(dataFileName(tagCategoriesEnumFile))
		if err != nil {
			return
		}
		defer f.Close()
		var enums []enumerationYAML
		err = yaml.NewDecoder(f).Decode(&enums)
		if err != nil {
			return
		}

		categoriesEnumerationSingleton = make([]TagCategoriesEnumeration, len(enums))

		for idx, item := range enums {
			categoriesEnumerationSingleton[idx] = TagCategoriesEnumeration{yaml: item}
		}
	})

	return categoriesEnumerationSingleton
}
