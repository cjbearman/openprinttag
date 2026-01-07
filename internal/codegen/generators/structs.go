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
	"strings"

	"github.com/cjbearman/openprinttag/internal/config"
	st "github.com/cjbearman/openprinttag/structtags"
)

// GenerateStruct generates structures for the meta, main and aux regions
func GenerateStructs(pkgDir string) {
	cfg, err := config.GetConfig(config.ConfigNFCV)
	if err != nil {
		panic(fmt.Sprintf("Failed to load configuration: %v", err))
	}

	if err := generateStruct(filepath.Join(pkgDir, "meta_region.go"), "Meta", "meta", cfg.MetaFields(), false); err != nil {
		panic(err)
	}
	if err := generateStruct(filepath.Join(pkgDir, "main_region.go"), "Main", "main", cfg.MainFields(), false); err != nil {
		panic(err)
	}
	if err := generateStruct(filepath.Join(pkgDir, "aux_region.go"), "Aux", "aux", cfg.AuxFields(), true); err != nil {
		panic(err)
	}
}

// generateStruct generates an individual struct into the specified filename
// as <prefix>Region and utilizing the specified field set
func generateStruct(filename string, prefix, name string, fields []config.Field, isAux bool) error {

	prefixLC := strings.ToLower(prefix)

	// We use multiple buffers and writers to allow us to write different parts of the file
	// simultaneously and then merge the results

	// This buffer will be used for the final source generation, + writer
	sourceBuffer := new(bytes.Buffer)
	sourceWriter := bufio.NewWriter(sourceBuffer)
	writeCopyright(sourceWriter)

	// This buffer will be used for the struct generation, + writer
	var structBuffer = new(bytes.Buffer)
	var structWriter = bufio.NewWriter(structBuffer)

	// This buffer will be used for the functions generation, + writer
	var funcBuffer = new(bytes.Buffer)
	var funcWriter = bufio.NewWriter(funcBuffer)

	// This set is all imports needed for the file
	// We may add to this later depending on the requirements for the individual types
	// within the file
	importMap := map[string]bool{}

	// Create internal YAML decoder
	internalTypeName := fmt.Sprintf("%sInternal", prefixLC)
	externalTypeName := fmt.Sprintf("%sRegion", prefix)

	// struct header for our "internal" struct
	fmt.Fprintf(structWriter, "type %s struct {\n", internalTypeName)

	// Add each field to the struct
	for _, field := range fields {
		if field.IsDeprecated() {
			continue
		}

		// Figure out what opt struct tag we'll put on this field
		optAnnotations := []string{
			fmt.Sprintf("%s=%s", st.OptTagName, field.Name()),
			fmt.Sprintf("%s=%d", st.OptTagKey, field.Key()),
		}
		if field.Required() == "true" {
			optAnnotations = append(optAnnotations, st.OptTagRequired)
		}
		if field.Required() == "recommended" {
			optAnnotations = append(optAnnotations, st.OptTagRecommended)
		}
		if field.MaxLength() != 0 {
			optAnnotations = append(optAnnotations, fmt.Sprintf("%s=%d", st.OptTagMaxLength, field.MaxLength()))
		}
		if field.Type() == "color_rgba" {
			optAnnotations = append(optAnnotations, st.OptTagRGBA)
		}

		// Compile the finalized opt: struct tag
		optAnnotation := fmt.Sprintf("%s:\"%s\"", st.OptTag, strings.Join(optAnnotations, ","))

		theType, imports := field.GetInternalTypeAndImports()
		for _, item := range imports {
			importMap[item] = true
		}

		// Generate fhe field, including struct tags for cbor, json, yaml and our own opt tag
		additionalYamlTags := ""
		if field.Type() == "color_rgba" {
			additionalYamlTags = ",flow"
		}
		fmt.Fprintf(structWriter, "  %s *%s `cbor:\"%d,keyasint,omitempty\" yaml:\"%s%s,omitempty\" %s`\n", field.GetInternalFieldName(), theType, field.Key(), field.Name(), additionalYamlTags, optAnnotation)

		// Now, into the functions writer, we will generate appropriate setter/getter/clearer functions from our templates
		// the template we use depends on the native type of the field
		// These functions are associated with our external struct

		// Generate a base comment to be used by setters, getters, clear functions
		commentTemplate := fmt.Sprintf("// #type#%s #op# the value of %s (%d)", field.GetInternalFieldName(), field.Name(), field.Key())
		if field.Description() != "" {
			commentTemplate += "\n// " + field.Description()
		}

		// Refine the base comment to be appropriate for getter, setter, clearer cases
		setterComment := strings.ReplaceAll(strings.ReplaceAll(commentTemplate, "#type#", "Set"), "#op#", "Sets")
		getterComment := strings.ReplaceAll(strings.ReplaceAll(commentTemplate, "#type#", "Get"), "#op#", "Gets")
		clearComment := strings.ReplaceAll(strings.ReplaceAll(commentTemplate, "#type#", "Clear"), "#op#", "Clears")

		// Generate a templater with common properties
		templater := NewTemplater().
			WithFieldName(field.GetInternalFieldName()).
			WithStructName(externalTypeName).
			WithSetterComment(setterComment).
			WithGetterComment(getterComment).
			WithClearComment(clearComment)

		// Generate the field using the appropriate template file and any necessary additional parameters
		switch field.Type() {
		case "uint64":
			templater.WithFilename("uint64.template").Generate(funcWriter)
		case "int":
			templater.WithFilename("int.template").Generate(funcWriter)
		case "color_rgba":
			templater.WithFilename("color_rgba.template").Generate(funcWriter)
		case "string":
			templater.WithFilename("string.template").Generate(funcWriter)
		case "timestamp":
			templater.WithFilename("timestamp.template").Generate(funcWriter)
		case "enum":
			templater.WithFilename("enum.template").WithEnumName(field.GetInternalEnumType()).Generate(funcWriter)

		case "enum_array":
			templater.WithFilename("enum_arr.template").WithEnumName(field.GetInternalEnumType()).Generate(funcWriter)
		case "uuid":
			templater.WithFilename("uuid.template").Generate(funcWriter)
		case "number":
			templater.WithFilename("number.template").Generate(funcWriter)

		default:
			panic("Unhanled type " + field.Type())
		}
	}

	// All internal structs have a map of unknown fields
	fmt.Fprintf(structWriter, "  Unknowns map[any]any `cbor:\"-\" yaml:\"other,omitempty\"`")

	// The external struct needs a getter for unknown fields
	NewTemplater().
		WithFilename("unknowns.template").
		WithStructName(externalTypeName).
		Generate(funcWriter)

	// The external struct needs a (private) getter to retrieve internal
	fmt.Fprintf(funcWriter, "func (s %s) getInternal() any {\n", externalTypeName)
	fmt.Fprintf(funcWriter, "  return &s.internal\n")
	fmt.Fprintf(funcWriter, "}\n\n")

	// The external struct needs a region name getter
	fmt.Fprintf(funcWriter, "func (s %s) getRegionName() string {\n", externalTypeName)
	fmt.Fprintf(funcWriter, "  return \"%s\"", name)
	fmt.Fprintf(funcWriter, "}\n\n")

	// Function for retrieving encode options
	fmt.Fprintf(funcWriter, "// RegionOptions accesses encoding options for this region\n")
	fmt.Fprintf(funcWriter, "func (s %s) RegionOptions() *RegionOptions {\n", externalTypeName)
	fmt.Fprintf(funcWriter, "  return s.regionOptions\n")
	fmt.Fprintf(funcWriter, "}\n\n")

	// Function for vendor specific fields, aux region only
	if isAux {
		NewTemplater().WithFilename("vendor.template").
			WithStructName(externalTypeName).
			Generate(funcWriter)

	}

	// We can now close the struct
	fmt.Fprintf(structWriter, "}\n\n")

	// Create external struct, referencing our internal struct and encode options
	fmt.Fprintf(structWriter, "type %s struct {\n", externalTypeName)
	fmt.Fprintf(structWriter, "  internal %s\n", internalTypeName)
	fmt.Fprintf(structWriter, "  regionOptions *RegionOptions\n")
	fmt.Fprintf(structWriter, "}\n")

	// Now that we've completed our structs, we can commit everything into our main writer

	// First the preamble
	sourceWriter.WriteString(Preamble)

	// Next any imports
	if len(importMap) > 0 {
		sourceWriter.WriteString("import (\n")
		for imp := range importMap {
			fmt.Fprintf(sourceWriter, "  \"%s\"\n", imp)
		}
		sourceWriter.WriteString(")\n\n")
	}

	// Finally flush and output the struct bodies
	structWriter.Flush()
	sourceWriter.Write(structBuffer.Bytes())

	// Setters and getters and other functions
	funcWriter.Flush()
	sourceWriter.Write(funcBuffer.Bytes())

	// All written to main buffer
	sourceWriter.Flush()

	// Now we can format for consistency
	formatted := formatCode(filename, sourceBuffer)

	// Write the file
	return writeCodeFile(filename, formatted)
}
