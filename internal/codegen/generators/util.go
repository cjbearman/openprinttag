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
	"bytes"
	"fmt"
	"go/format"
	"os"
)

const (
	// Preamble is written to every code file
	Preamble = "package openprinttag\n\n// ** THIS FILE IS AUTO-GENERATED, DO NOT MODIFY **\n\n"
)

// writeCodeFile writes code content to the specified file and outputs a message
// indicating success case (error returned it if fails)
func writeCodeFile(filename string, content []byte) error {
	err := os.WriteFile(filename, content, 0644)
	if err != nil {
		return fmt.Errorf("failed to write generated code file %s: %w", filename, err)
	}
	fmt.Printf("Successfully wrote %s\n", filename)
	return nil
}

// formatCode will take a byte buffer containing a proposed code file
// and format it. If the format fails, the code is returned as provided
// after a warning message is output
func formatCode(filename string, code *bytes.Buffer) []byte {
	// Now we can format for consistency
	formatted, err := format.Source(code.Bytes())
	if err != nil {
		fmt.Printf("[WARNING] Failed to properly format %s: %v\n", filename, err)
		return code.Bytes()
	}
	return formatted
}
