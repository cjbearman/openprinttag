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
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/cjbearman/openprinttag/internal/codegen/generators"
)

// Code generation
// To regenerate the auto-generated files in the pkg directory, simply run "go generate ./..." from the root directory
// of this project after (if needed) updating the source_repo submodule to the required version of the original prusa source

//go:generate go run generate.go

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		panic("code gen failed, unable to get working directory")
	}

	// All of our generate files go into our main directory
	pkgDir := filepath.Join(filepath.Dir(filepath.Dir(cwd)))

	// Remove all generated files
	fmt.Printf("Evaluating files in %s to determine which are auto-generated\n", pkgDir)
	entries, err := os.ReadDir(pkgDir)
	if err != nil {
		panic(fmt.Sprintf("Failed to read package dir (%s): %v", pkgDir, err))
	}
	for _, entry := range entries {
		fmt.Sprintf("File: %s", entry.Name())
		if !entry.Type().IsRegular() {
			continue
		}
		absName := filepath.Join(pkgDir, entry.Name())
		contents, err := os.ReadFile(absName)
		if err != nil {
			panic(fmt.Sprintf("Failed to read file %s to evaluate if it's auto generated: %v", absName, err))
		}
		if strings.Contains(string(contents), generators.Preamble) {
			fmt.Printf("Will remove %s because it's an auto-generated file\n", entry.Name())
			err := os.Remove(absName)
			if err != nil {
				panic(fmt.Sprintf("Failed to remove %s: %v", absName, err))
			}
		}
	}

	generators.GenerateEnums(pkgDir)
	generators.GenerateStructs(pkgDir)

}
