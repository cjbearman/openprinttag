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
package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/cjbearman/openprinttag"
)

var load, out, imprt, setURI string
var soft, useYaml, optcheck, validate, uuids, root, regions, uri, all,
	discardAux, hexForm, b64, hexDump, testMode, nocc bool
var initTag, auxSize, metaSize, blockSize int

func cmdLine() {
	// Command line
	flag.StringVar(&load, "load", "", "Loads an existing open print tag from a file (or specify \"-\" to load from STDIN)")
	flag.StringVar(&out, "out", "", "Outputs the completed tag to a file (or specify \"-\" to output to STDOUT)")
	flag.StringVar(&imprt, "data", "", "Import YAML encoded data and apply to tag")
	flag.BoolVar(&soft, "soft", false, "When importing data to a tag, do not overwrite fields already set in the tag")
	flag.BoolVar(&useYaml, "yaml", false, "output as YAML instead of binary tag")
	flag.BoolVar(&optcheck, "opt-check", false, "Run opt-check, requires -yaml")
	flag.BoolVar(&validate, "validate", false, "Validate required/recommended fields, requires -yaml")
	flag.BoolVar(&uuids, "uuids", false, "Output defined/calculated UUIDs, require -yaml")
	flag.BoolVar(&root, "root", false, "Output root information, requires -yaml")
	flag.BoolVar(&regions, "regions", false, "Output region information, requires -yaml")
	flag.BoolVar(&uri, "uri", false, "Output URI information, requires -yaml")
	flag.BoolVar(&all, "all", false, "Output all possible YAML information, requires -yaml")
	flag.IntVar(&initTag, "init", 0, "Initialize a new tag with the provided size")
	flag.IntVar(&auxSize, "aux-size", 0, "Set size of aux section")
	flag.IntVar(&metaSize, "meta-size", 0, "Set size of meta section")
	flag.IntVar(&blockSize, "block-size", 0, "Set block size")
	flag.StringVar(&setURI, "set-uri", "", "Set URI")
	flag.BoolVar(&discardAux, "discard-aux", false, "Discard the AUX region")
	flag.BoolVar(&hexForm, "hex", false, "Output tag in hex format, -out required")
	flag.BoolVar(&b64, "base-64", false, "Output tag in base64 format, -out required")
	flag.BoolVar(&hexDump, "hex-dump", false, "Output tag in hex dump format, -out required")
	flag.BoolVar(&testMode, "test-mode", false, "Sets parameters used for integration test")
	flag.BoolVar(&nocc, "no-cc", false, "Disable capability container encoding/decoding")

	flag.Parse()

	if optcheck && !useYaml {
		terminal(errors.New("-opt-check flag requires -yaml flag"))
	}
	if validate && !useYaml {
		terminal(errors.New("-validate flag requires -yaml flag"))
	}
	if uuids && !useYaml {
		terminal(errors.New("-uuids flag requires -yaml flag"))
	}
	if root && !useYaml {
		terminal(errors.New("-uuids flag requires -yaml flag"))
	}
	if regions && !useYaml {
		terminal(errors.New("-uuids flag requires -yaml flag"))
	}
	if uri && !useYaml {
		terminal(errors.New("-uuids flag requires -yaml flag"))
	}
	if all && !useYaml {
		terminal(errors.New("-all flag requires -yaml flag"))
	}
	if initTag != 0 && load != "" {
		terminal(errors.New("-init and -load cannot be used together"))
	}
	if initTag == 0 && load == "" {
		terminal(errors.New("must use either -init or -load"))
	}
	if hexForm && useYaml {
		terminal(errors.New("-hex flag and -yaml are mutually exclusive"))
	}
	if hexDump && useYaml {
		terminal(errors.New("-hex-dump and -yaml are mutually exclusive"))
	}
	if b64 && useYaml {
		terminal(errors.New("-base-64 and -yaml are mutually exclusive"))
	}
	if hexForm && hexDump {
		terminal(errors.New("-hex and -hex-dump are mutually exclusive"))
	}
	if hexForm && b64 {
		terminal(errors.New("-hex and -base-64 are mutually exclusive"))
	}
	if hexDump && b64 {
		terminal(errors.New("-hex-dump and -base-64 are mutually exclusive"))
	}

}

func main() {
	cmdLine()

	// Set up encode/decode options
	ecOpts := []openprinttag.EncodeDecodeOption{}
	if nocc {
		ecOpts = append(ecOpts, openprinttag.WithoutCapabilityContainer)
	}

	// First step we create our working tag, either as new or from file
	var tag *openprinttag.OpenPrintTag
	if initTag != 0 {
		tag = openprinttag.NewOpenPrintTag().WithSize(initTag)
	} else {
		tagData := loadTag(load)
		var err error
		tag, err = openprinttag.Decode(tagData, ecOpts...)
		if err != nil {
			terminal(fmt.Errorf("failed to decode tag: %w", err))
		}
	}

	if testMode {
		tag.MainRegion().RegionOptions().SetFloatMaxPrecision(openprinttag.FloatMaxPrecision16)
	}

	if auxSize != 0 {
		tag.WithAuxRegionSize(auxSize)
	}

	if metaSize != 0 {
		tag.WithMetaRegionSize(metaSize)
	}
	if blockSize != 0 {
		tag.WithBlockSize(blockSize)
	}
	if setURI != "" {
		tag.WithURIRecord(setURI)
	}
	if discardAux {
		tag.RemoveAuxRegion().WithAuxRegionSize(0)
	}

	// Import stage
	if imprt != "" {
		imported := loadRecords(imprt)
		tag.Merge(imported, !soft)
	}

	// Output stage
	if useYaml {
		options := []openprinttag.YAMLOption{}
		includeIf := func(condition bool, option openprinttag.YAMLOption) {
			if condition {
				options = append(options, option)
			}
		}
		includeIf(optcheck, openprinttag.IncludeOptCheck)
		includeIf(validate, openprinttag.IncludeValidation)
		includeIf(uuids, openprinttag.IncludeUUIDs)
		includeIf(root, openprinttag.IncludeRootStats)
		includeIf(regions, openprinttag.IncludeRegionStats)
		includeIf(uri, openprinttag.IncludeURI)
		includeIf(all, openprinttag.IncludeAll)
		yamlData, err := tag.ToYAML(options...)
		if err != nil {
			terminal(fmt.Errorf("failed to format tag as YAML: %w\n", err))
		}
		writeOutput(out, []byte(yamlData))
	} else {
		bintag, err := tag.Encode(ecOpts...)
		if err != nil {
			terminal(fmt.Errorf("failed to encode tag: %w\n", err))
		}
		writeOutput(out, bintag)
	}
	os.Exit(0)
}

func loadTag(filename string) []byte {
	if filename == "-" {
		// From stdin
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			terminal(fmt.Errorf("failed to read tag from stdin: %w", err))
		}
		return data
	}

	// From file
	data, err := os.ReadFile(filename)
	if err != nil {
		terminal(fmt.Errorf("failed to read tag from %s: %w", filename, err))
	}
	return data
}

func loadRecords(filename string) *openprinttag.OpenPrintTag {
	data, err := os.ReadFile(filename)
	if err != nil {
		terminal(fmt.Errorf("failed to read from %s: %w", filename, err))
	}
	tag, err := openprinttag.FromYAML(string(data))
	if err != nil {
		terminal(fmt.Errorf("failed to parse yaml data from %s: %w", filename, err))
	}
	return tag
}

func writeOutput(filename string, output []byte) []byte {

	var finalized []byte = output
	if hexForm {
		finalized = []byte(hex.EncodeToString(output) + "\n")
	} else if hexDump {
		buf := &bytes.Buffer{}
		writer := hex.Dumper(buf)
		_, err := writer.Write(output)
		writer.Close()
		if err != nil {
			terminal(fmt.Errorf("failed to hex dump output: %w", err))
		}
		finalized = buf.Bytes()
	} else if b64 {
		finalized = []byte(base64.StdEncoding.EncodeToString(output) + "\n")
	}

	if filename == "" {
		// To stdout
		br := bufio.NewWriter(os.Stdout)
		_, err := br.Write(finalized)
		if err != nil {
			terminal(fmt.Errorf("failed to write output to STDOUT: %w", err))
		}
		err = br.Flush()
		if err != nil {
			terminal(fmt.Errorf("failed to flush output data to STDOUT: %w", err))
		}

	} else {
		// To file
		err := os.WriteFile(filename, finalized, 0644)
		if err != nil {
			terminal(fmt.Errorf("faildd to write output to %s: %w", filename, err))

		}
	}
	return nil
}

// terminal will receive an error and if not nil will output the error to stderr
// and exit with code 1
func terminal(err error) {
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error()+"\n")
		os.Exit(1)
	}
}
