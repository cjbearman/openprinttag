# Open Print Tag for Golang
This is a pure golang based implementation of the Open Print Tag standard initiated by Prusa Research, as documented at https://openprinttag.org/ .

## Intent
The design will encode/decode open print tags in accordance with the published specification above, and tries to get as close as possible to the exact encodings used in the python examples from the primary project repository.

Provision is made for:
1. A fully featured golang module that can be used programatically for the generation, modification and interpretation of open print tags.
2. A binary that can be used for these same functions

Code generation is heavily used to generate the specific encoders and decoders, allowing the library to adhere ridgidly and automatically to the published data formats.

## License

Licensed under the MIT license

```
MIT License

Copyright (c) 2026 Christopher J Bearman

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

## GO Version
This module supports go1.21+.

MAC users should be aware of issues that may result in the following linker error:
```
dyld[3723]: missing LC_UUID load command
```
If you encounter this issue, please switch to go1.24 or higher to resolve.



## Using as a golang module
```
go get github.com/cjbearman/openprinttag
```

## Documentation
Refer to https://pkg.go.dev/github.com/cjbearman/openprinttag for API documentation.

### General Usage
The API for the openprintag module provides a fluent methodology for creating, decoding, modifying and outputting tags.

Make sure and import the module:
```
import "github.com/cjbearman/openprinttag"
```

### Create a new blank tag
```golang
	tag := openprinttag.NewOpenPrintTag().
        WithAuxRegionSize(32).
        WithSize(304)
```
The newly created tag variable represents the open print tag and provides methods for accessing the three tag regions, encoding the tag and so forth.

### Full example of creating a tag
```golang
	tag := openprinttag.NewOpenPrintTag().
		WithAuxRegionSize(32).
		WithSize(304)

	tag.MainRegion().
		SetBrandName("Awesome Filaments").
		SetMaterialName("Fancy PLA Yellow").
		SetPrimaryColor(openprinttag.MustNewColor("#FFFF00")).
		SetMaterialClass(openprinttag.MaterialClassFFF).
		SetMaterialType(openprinttag.MaterialTypePLA)
	
	tagData, err:= tag.Encode()
```
N.B. Each region provides three functions for each possible parameter. A "Set" function is used to set the value of the parameter. A "Get" function is used to retrieve the value of the parameter, and includes a second return value indicating whether the parameter was present. A "Clear" function can be used to erase the current content of the parameter.


### Example for reading an existing tag
```golang
func readTag(r io.Reader) *openprinttag.OpenPrintTag {

	tagBytes, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}

	tag, err := openprinttag.Decode(tagBytes)
	if err != nil {
		panic("failed to read tag")
	}

	if brandName, found := tag.MainRegion().GetBrandName(); found {
		fmt.Printf("This tag is from brand: %s\n", brandName)
	} else {
		fmt.Printf("No brand name found in this tag")
	}

	return tag
}
```
This example implements a function to read a tag from a reader, print out the manufacturer information if present, and return the loaded tag to the caller.

### UUID encoding
Four functions are provided to encode brand, material, package and instance UUIDs according to specification:
```
func BrandUUID(brandName string) uuid.UUID
func MaterialUUID(materialName string, brandUUID uuid.UUID) uuid.UUID
func MaterialPackageUUID(gtin string, brandUUID uuid.UUID) uuid.UUID
func MaterialPackageInstanceUUID(NFCTagUid []byte) (uuid.UUID, error)
```

## Command line tool
The optional "optag" binary is provided as an example, as well as a useful tool for creating and modifying tags.

### Installation:
```
go install github.com/cjbearman/openprinttag
```

### General Operation
The optag tool follows a three stage process:
1. Initialize a blank tag, or, load an existing tag from binary data
2. Modify the tag by applying data to it from a YAML data file
3. Encode the tag and output it either as a binary tag, or a YAML document

Various options control this process.

### Example: Initialize a new blank tag
```
$ optag -init 304 -aux-size 12
.. binary data output (not shown)
```

### Example: Load a tag from a file and decode/debug in YAML format
This example loads an existing (binary) tag from a file, and outputs it's data and statistics.
```
$ optag -load ../../test_outputs/TestStats.tag -yaml -all 
regions:
    meta:
        payload_offset: 0
        absolute_offset: 42
        size: 4
        used_size: 4
    main:
        payload_offset: 4
        absolute_offset: 46
        size: 222
        used_size: 119
    aux:
        payload_offset: 226
        absolute_offset: 268
        size: 35
        used_size: 1
root:
    data_size: 304
    payload_size: 261
    overhead: 43
    payload_used_size: 124
    total_used_size: 167
data:
    meta:
        aux_region_offset: 226
    main:
        instance_uuid: 473bb8cd-e129-45b8-9fcf-da1c3add9c47
        brand_specific_material_id: "1"
        material_class: FFF
        material_type: PLA
        material_name: PLA Galaxy Black
        brand_name: Prusament
        manufactured_date: 1739371290
        nominal_netto_full_weight: 1000
        actual_netto_full_weight: 1012
        empty_container_weight: 100
        primary_color: '#3d3e3d'
        transmission_distance: 0.199951171875
        tags:
            - glitter
        min_print_temperature: 205
        max_print_temperature: 220
        preheat_temperature: 170
        min_bed_temperature: 40
        max_bed_temperature: 60
        max_chamber_temperature: 40
        chamber_temperature: 20
        container_width: 75
    aux: {}
validate:
    warnings:
        - field Gtin (gtin/4) is recommended
        - field Density (density/29) is recommended
    errors: []
opt_check:
    warnings: []
    errors: []
    notes: []
uuids:
    brand_uuid: ae5ff34e-298e-50c9-8f77-92a97fb30b09
    material_uuid: 6e774110-9aa4-5ab2-a269-456918dad9b1
    package_uuid: null
    instance_uuid: 473bb8cd-e129-45b8-9fcf-da1c3add9c47
```

### Example: Initialize a blank tag and import some fields
The YAML structure used for field declaration is the same as provided by the reference implementation at https://openprinttag.org/.
```
$ cat some_fields.yaml
data:
    meta:
        aux_region_offset: 226
    main:
        instance_uuid: 473bb8cd-e129-45b8-9fcf-da1c3add9c47
        brand_specific_material_id: "1"
        material_class: FFF
        material_type: PLA
        material_name: PLA Galaxy Black
        brand_name: Prusament
        tags:
            - glitter


$ optag -init 304 -aux-size 32 -data some_fields.yaml -yaml     
data:
    meta: {}
    main:
        instance_uuid: 473bb8cd-e129-45b8-9fcf-da1c3add9c47
        brand_specific_material_id: "1"
        material_class: FFF
        material_type: PLA
        material_name: PLA Galaxy Black
        brand_name: Prusament
        tags:
            - glitter
    aux: {}
```
N.B. Omitting the -yaml option would have output the binary form of the tag

### Notes
The opttag utility can also be used with pipes. We could, for example, pipe in an existing tag, add some fields and output the resulting tag.
When loading an existing tag from standard input, use "-load -"
```
cat existing_tag.bin | optag -load - -data addfields.yaml > new_tag.bin
```

### All Options
The usage message can be obtained by using the -h option:
```
Usage of optag:
  -all
    	Output all possible YAML information, requires -yaml
  -aux-size int
    	Set size of aux section
  -base-64
    	Output tag in base64 format, -out required
  -block-size int
    	Set block size
  -data string
    	Import YAML encoded data and apply to tag
  -discard-aux
    	Discard the AUX region
  -hex
    	Output tag in hex format, -out required
  -hex-dump
    	Output tag in hex dump format, -out required
  -init int
    	Initialize a new tag with the provided size
  -load string
    	Loads an existing open print tag from a file (or specify "-" to load from STDIN)
  -meta-size int
    	Set size of meta section
  -opt-check
    	Run opt-check, requires -yaml
  -out string
    	Outputs the completed tag to a file (or specify "-" to output to STDOUT)
  -regions
    	Output region information, requires -yaml
  -root
    	Output root information, requires -yaml
  -set-uri string
    	Set URI
  -soft
    	When importing data to a tag, do not overwrite fields already set in the tag
  -uri
    	Output URI information, requires -yaml
  -uuids
    	Output defined/calculated UUIDs, require -yaml
  -validate
    	Validate required/recommended fields, requires -yaml
  -yaml
    	output as YAML instead of binary tag
```

## Development
This library is designed to automatically generate much of it's own code from the reference implementation at https://github.com/prusa3d/OpenPrintTag. To this end, this reference repository is included herein as a git submodule at the source_repo path.

To checkout with submodule
```
go clone --recurse-submodules git@github.com:cjbearman/openprinttag.git
```

To update this project to a later (future) version of the reference:
1. Update the submodule
2. Regenerate the auto-generated code files using go generate as follows:
```
go generate ./...
```
The code generation process is driven by the [codegen](https://github.com/cjbearman/openprinttag/tree/main/internal/codegen) and [config](https://github.com/cjbearman/openprinttag/tree/main/internal/config) packages. The codegen package first uses the config package to load the YAML files describing the open print tag data formats. Once these are read and processed, the codegen package uses it to generate code files in the root directory. Code files are generated for each open print tag region and each enumerated type.

Auto-generated files all containg the following comment:
```
// ** THIS FILE IS AUTO-GENERATED, DO NOT MODIFY **
```

Should you update to a new version of the reference (source_repo) repository and run auto-generation, there is a possibility that changes to the data files will cause compilation issues in other (non-generated) files.  Changes made to expected things, such as the introduction of new fields, deprecation of old fields, introduction of new enumerations and so forth should be handled gracefully, however, there is always the prospect that some change to the reference implementation may need some adjustments in the config loader or code generator packages. It is more likely that updates to the reference implementation and subsequent regeneration of code will require tweaks in the unit tests to bring them into compliance with those changes.

To install binary tools from local clone:
```
go install ./...
```
Tools are installed to your local GOPATH bin directory, enusre this is in your path.

## Automated Tests
Use standard golang unit testing procedures:
```
go test -v ./...
```
Some of the tests will write binary and YAML tag data to the test_outputs directory. These outputs can be used for manual testing and comparison against equivalents produced by reference implementation.

## Integration Tests
The integration_tests directory contains some tests to create the same tags using the reference python code and this module and compare the binary outputs, reporting any discrepancies as failures.

This is somewhat tentative as there are many reasons why we would generate slightly different binary representations (field order and so forth), so failures here after updating the reference architecture are highly likely and not necessarily indicative of a problem.

## Development Scripts
The scripts directory contain a few shell scripts that may be useful when developing on Mac/Linux platforms (sorry I don't do Windows).

### scripts/build_all.sh
Runs code generation, rebuilds all binaries and creates new development docker container.

### scripts/docs.sh
Launches a local doc server on port 8080

### scripts/pyvirtual.sh
Creates a python virtual environment under (project root)/.venv and initializes it with all dependencies needed by the (python based) source_repo. Run this script using:
```
source scripts/pyvirtual.sh
```
to activate the environment. 

This simply produces an isolated environment for running the reference scripts, which is useful when you need to compare the output of this module with output produced by the reference impelmentation.

## Docker development container
The docker_test directory contains a docker build and helper scripts for running the reference scripts from source_repo via docker, avoiding the need to mess around with virtual environments locally. The docker container is used by the integration tests.
