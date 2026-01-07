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
	"reflect"
	"strconv"
	"strings"

	st "github.com/cjbearman/openprinttag/structtags"
)

// Validate checks for missing required (error) or recommended (warnings)
// fields in all regions
func (o *OpenPrintTag) Validate() (errors []string, warnings []string) {
	if o.meta != nil {
		e, w := validateRegion(o.meta)
		errors = append(errors, e...)
		warnings = append(warnings, w...)
	}
	if o.main != nil {
		e, w := validateRegion(o.main)
		errors = append(errors, e...)
		warnings = append(warnings, w...)
	}
	if o.aux != nil {
		e, w := validateRegion(o.aux)
		errors = append(errors, e...)
		warnings = append(warnings, w...)
	}
	return
}

// Validate will return any errors or warnings for a specific region
func validateRegion(region Region) (errors, warnings []string) {
	defer func() {
		// Just in case something goes catastrophically wrong
		if r := recover(); r != nil {
			errors = append(errors, fmt.Sprintf("panic during validation: %v", r))
		}
	}()

	// This is an interface, get it's value
	obj := reflect.ValueOf(region).Elem()

	// It must have an internal field, get that
	internal := obj.FieldByName("internal")

	// Iterate through all fields in the internal struct
	for i := 0; i < internal.NumField(); i++ {
		// The field name
		name := internal.Type().Field(i).Name

		// The field (as reflect.Field)
		field := internal.Type().Field(i)

		// The opt tag set
		tag := field.Tag.Get(st.OptTag)
		if tag == "" {
			// No opt tag, since all our fields have tags, we can skip this field as being of no interest
			continue
		}

		// Decode the opt tag into it's constituent parts
		tagMap := decodeOptTag(tag)

		// The native field name
		nativeName := tagMap[st.OptTagName]

		// The native field key (string encoded integer)
		key := tagMap[st.OptTagKey]

		// The field instance value
		fieldValue := internal.Field(i)

		// Decode the value, which is a pointer to the actual value where set
		var valueIsNil = true
		if !fieldValue.IsNil() {
			valueIsNil = false
		}

		// Is this required, if so error if nil
		_, isRequired := tagMap[st.OptTagRequired]
		if isRequired && valueIsNil {
			errors = append(errors, genErrorOrWarning(name, key, nativeName, "is required"))
		}

		// Is this recommended, if so warning if nil
		_, isRecommended := tagMap[st.OptTagRecommended]
		if isRecommended && valueIsNil {
			warnings = append(warnings, genErrorOrWarning(name, key, nativeName, "is recommended"))
		}
	}
	return
}

// OptCheck checks options for warnings and errors in all present regions
func (o *OpenPrintTag) OptCheck() (errors []string, warnings []string) {
	if o.meta != nil {
		e, w := optCheck(o.meta)
		errors = append(errors, e...)
		warnings = append(warnings, w...)
	}
	if o.main != nil {
		e, w := optCheck(o.main)
		errors = append(errors, e...)
		warnings = append(warnings, w...)
	}
	if o.aux != nil {
		e, w := optCheck(o.aux)
		errors = append(errors, e...)
		warnings = append(warnings, w...)
	}
	return
}

// optCheck returns errors and warnings for a specific region
func optCheck(region Region) (errors, warnings []string) {
	defer func() {
		// Just in case something goes catastrophically wrong
		if r := recover(); r != nil {
			errors = append(errors, fmt.Sprintf("panic during validation: %v", r))
		}
	}()

	// This is an interface, get it's value
	obj := reflect.ValueOf(region).Elem()

	// It must have an internal field, get that
	internal := obj.FieldByName("internal")

	// Iterate through all fields in the internal struct
	for i := 0; i < internal.NumField(); i++ {
		// The field name
		name := internal.Type().Field(i).Name

		// The field (as reflect.Field)
		field := internal.Type().Field(i)

		// The opt tag set
		tag := field.Tag.Get(st.OptTag)
		if tag == "" {
			// No opt tag, since all our fields have tags, we can skip this field as being of no interest
			continue
		}

		// Decode the opt tag into it's constituent parts
		tagMap := decodeOptTag(tag)

		// The native field name
		nativeName := tagMap[st.OptTagName]

		// The native field key (string encoded integer)
		key := tagMap[st.OptTagKey]

		// The field instance value
		fieldValue := internal.Field(i)

		// Decode the value, which is a pointer to the actual value where set
		var value reflect.Value
		if fieldValue.IsNil() {
			continue
		}
		value = fieldValue.Elem()

		// Check for max length
		maxLenStr, found := tagMap[st.OptTagMaxLength]
		if found {
			maxLen, err := strconv.Atoi(maxLenStr)
			if err == nil {
				var length int
				// max_length is applicable only to string and slice fields
				if value.Kind() == reflect.String || value.Kind() == reflect.Slice {
					length = value.Len()
					if length > maxLen {
						errors = append(errors, genErrorOrWarning(name, key, nativeName, "has length %d which exceeds maximum length of %d for this field", length, maxLen))
					}
				}
			}
		}

		// Check for rgba
		_, rgba := tagMap[st.OptTagRGBA]
		if rgba {
			length := value.Len()
			if length != 3 && length != 4 {
				errors = append(errors, genErrorOrWarning(name, key, nativeName, "has length %d which is not valid for RGBA fields (must be 3 or 4)", length))
			}
		}

	}
	customErrors, customWarnings := getCustomErrorsAndWarnings(region)
	errors = append(errors, customErrors...)
	warnings = append(warnings, customWarnings...)

	return
}

// GetCustomErrorsAndWarnings returns errors/warnings that cannot be auto-derived and must be hand coded
func getCustomErrorsAndWarnings(region Region) (errors, warnings []string) {
	if main, ok := region.(*MainRegion); ok {
		brand, brandSet := main.GetBrandName()
		brandUUID, brandUUIDSet := main.GetBrandUuid()

		if brandSet && brandUUIDSet {
			calculatedBrandUUID := BrandUUID(brand)
			if calculatedBrandUUID.String() == brandUUID.String() {
				warnings = append(warnings, "brand_uuid is identical to the auto-generated version, and thus can be omitted to save space")
			}
		}

		material, materialSet := main.GetMaterialName()
		materialUUID, materialUUIDSet := main.GetMaterialUuid()
		if materialSet && materialUUIDSet && (brandSet || brandUUIDSet) {
			if !brandUUIDSet {
				brandUUID = BrandUUID(brand)
			}
			calculatedMaterialUUID := MaterialUUID(material, brandUUID)
			if calculatedMaterialUUID.String() == materialUUID.String() {
				warnings = append(warnings, "material_uuid is identical to the auto-generated version, and thus can be omitted to save space")
			}
		}
		// TODO: Gtin and package need to be added
	}
	return
}

// genErrorOrWarning will return a suitably formatted field based error or warning
// given the name, key and native name of the field (for consistency)
// and additional format and args
func genErrorOrWarning(name, key, nativeName, format string, v ...any) string {
	preamble := fmt.Sprintf("field %s (%s/%s) ", name, nativeName, key)
	return fmt.Sprintf(preamble+format, v...)
}

// decodeOptTag takes the value of the opt struct tag and decodes it to a map
// containing key (the subtag name) and value (the optional subtag value)
func decodeOptTag(optTag string) map[string]string {
	tagBits := strings.Split(optTag, ",")
	tagMap := make(map[string]string, len(tagBits))

	for _, tagBit := range tagBits {
		tagSplit := strings.Split(tagBit, "=")
		if len(tagSplit) == 1 {
			tagMap[tagSplit[0]] = ""
		} else if len(tagSplit) == 2 {

			tagMap[tagSplit[0]] = tagSplit[1]
		}
	}
	return tagMap
}
