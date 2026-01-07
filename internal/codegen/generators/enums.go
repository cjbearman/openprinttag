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
package generators

import (
	"bufio"
	"bytes"
	"fmt"
	"path/filepath"

	"github.com/cjbearman/openprinttag/internal/config"
)

const (
	enumsPreamble = "package openprinttag\n\n// ** THIS FILE IS AUTO-GENERATED, DO NOT MODIFY **\n\n"
)

// GenerateEnums generates all enumerations
func GenerateEnums(pkgDir string) {

	// Load the config
	cfg, err := config.GetConfig(config.ConfigNFCV)
	if err != nil {
		panic(fmt.Sprintf("Failed to load configuration: %v", err))
	}

	// Find the names of all fields in all region objects
	allFields := append(cfg.MetaFields(), cfg.MainFields()...)
	allFields = append(allFields, cfg.AuxFields()...)

	// It is possible that an enumeration may be used more than once, but
	// we should only generate it once
	// Use a set to keep track of which we have generated
	processedEnumerations := make(map[string]bool)

	// Go through all fields to find enumerations
	for _, field := range allFields {
		if !field.HasEnumeration() {
			continue
		}

		// Check to see if this enumeration has already been processed and skip if it has
		if processedEnumerations[field.GetInternalEnumType()] {
			continue
		}

		// Mark this as processed so it will be skipped next time
		processedEnumerations[field.GetInternalEnumType()] = true

		// Generate the enumeration
		if err := generateEnum(filepath.Join(pkgDir, enumFileName(field)), field, field.EnumeratedValues()); err != nil {
			panic(err)
		}
	}
}

// generateEnum generates an individual enumeration
func generateEnum(filename string, field config.Field, enumerations []config.Enumeration) error {

	// buf and wr will be used to write the enumeration
	buf := new(bytes.Buffer)
	wr := bufio.NewWriter(buf)
	writeCopyright(wr)

	// Preamble contains package and auto-generated file comment
	wr.WriteString(enumsPreamble)

	wr.WriteString("import (\n	\"fmt\"\n  \"gopkg.in/yaml.v3\"\n)\n\n")

	// Generate the enumeration type with integer base
	fmt.Fprintf(wr, "type %s uint64\n\n", field.GetInternalEnumType())

	// Generate const list for all individual enumerations
	fmt.Fprintf(wr, "const (\n")
	for _, enum := range enumerations {

		// Skip deprecated enumerations
		if enum.IsDeprecated() {

			continue
		}

		// Get the internal enumeration name, also an indication as to whether it has been abbreviated
		en, abbreviated := enum.GetInternalEnumName(field)

		// Generate comment for the name
		fmt.Fprintf(wr, "  // %s\n", en)
		if abbreviated {
			// If it's abbreviated, we'll include the full name first
			fmt.Fprintf(wr, " // %s\n", enum.Name())
		}
		// And always a description, if provided
		if enum.Description() != "" {
			fmt.Fprintf(wr, "  // %s\n", enum.Description())
		}

		// Write the actual enumerated name and value
		fmt.Fprintf(wr, "  %s %s = %d\n\n", en, field.GetInternalEnumType(), enum.Key())
	}

	// Close the const list
	fmt.Fprintf(wr, ")\n\n")

	// Add a map of all enum values
	fmt.Fprintf(wr, "var %s = map[uint64]string {\n", field.GetInternalEnumMapName())
	for _, enum := range enumerations {
		if enum.IsDeprecated() {
			continue
		}
		value := enum.Abbreviation()
		if value == "" {
			value = enum.Name()
		}
		key := enum.Key()
		fmt.Fprintf(wr, "  %d: \"%s\",\n", key, value)
	}
	fmt.Fprintf(wr, "}\n\n")

	// Generate a stringer to allow for reasonable debugging
	fmt.Fprintf(wr, "func (e %s) String() string {\n", field.GetInternalEnumType())
	fmt.Fprintf(wr, "  return %s[uint64(e)]\n", field.GetInternalEnumMapName())
	fmt.Fprintf(wr, "}\n\n")

	// Generate YAML marshal/unmarshal functions
	NewTemplater().
		WithEnumName(field.GetInternalEnumType()).
		WithMapName(field.GetInternalEnumMapName()).
		WithFilename("enum_marshallers.template").
		Generate(wr)

	// Complete, flush the writer
	wr.Flush()

	// Now we can format for consistency
	formatted := formatCode(filename, buf)

	// Output
	return writeCodeFile(filename, formatted)
}

// enumFileName returns the filename that we will use for a given enum
func enumFileName(f config.Field) string {
	return fmt.Sprintf("%s_enum.go", f.GetGeneratedEnumFilename())
}
