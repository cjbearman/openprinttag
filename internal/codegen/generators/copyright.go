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
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func writeCopyright(w io.Writer) {
	cwd, err := os.Getwd()
	if err != nil {
		panic("code gen failed, unable to get working directory")
	}

	pkgDir := filepath.Join(filepath.Dir(filepath.Dir(cwd)))
	lic := filepath.Join(pkgDir, "LICENSE")
	fp, err := os.Open(lic)
	if err != nil {
		panic("failed to read license file from " + lic)
	}
	defer fp.Close()

	scan := bufio.NewScanner(fp)
	for scan.Scan() {
		fmt.Fprintf(w, "// %s\n", scan.Text())
	}
	if err = scan.Err(); err != nil {
		panic("error reading license file: " + err.Error())
	}
}
