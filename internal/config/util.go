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
	"path/filepath"
)

func dataFileName(filename string) string {
	absoluteFilename := filepath.Join(rootDir(), filename)
	if _, err := os.Stat(absoluteFilename); err == nil {
		return absoluteFilename
	}
	panic(fmt.Sprintf("Unable to stat %s", absoluteFilename))
}

func rootDir() string {
	// The CWD will be the directory where the go generate annotation is, which is internal/codegen
	// This is guaranteed by go generate spec..
	cwd, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("Cannot find working directory: %v", err))
	}

	// We want the source_repo/data directory from the second parent of this CWD
	parent := filepath.Dir(filepath.Dir(cwd))
	dataDir := filepath.Join(parent, "source_repo", "data")

	// Make sure this is actually a directory
	st, err := os.Stat(dataDir)
	if err != nil {
		panic(fmt.Sprintf("Failed to stat source repo data directory (%s), was it checked out?: %v", dataDir, err))
	}
	if !st.IsDir() {
		panic(fmt.Sprintf("Source repo data directory (%s) does not appear to be a directory", dataDir))
	}

	return dataDir
}
