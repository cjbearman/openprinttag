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
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type Templater struct {
	structName    string
	fieldName     string
	enumName      string
	filename      string
	setterComment string
	getterComment string
	clearComment  string
	mapName       string
}

func NewTemplater() *Templater {
	return &Templater{}
}

func (t *Templater) WithFilename(filename string) *Templater {
	t.filename = filename
	return t
}

func (t *Templater) WithStructName(structName string) *Templater {
	t.structName = structName
	return t
}

func (t *Templater) WithFieldName(fieldName string) *Templater {
	t.fieldName = fieldName
	return t
}

func (t *Templater) WithEnumName(enumName string) *Templater {
	t.enumName = enumName
	return t
}

func (t *Templater) WithSetterComment(setterComment string) *Templater {
	t.setterComment = setterComment
	return t
}

func (t *Templater) WithGetterComment(getterComment string) *Templater {
	t.getterComment = getterComment
	return t
}

func (t *Templater) WithClearComment(clearComment string) *Templater {
	t.clearComment = clearComment
	return t
}

func (t *Templater) WithMapName(mapName string) *Templater {
	t.mapName = mapName
	return t
}

func (t *Templater) Generate(writer io.Writer) {
	if t.filename == "" {
		panic("templater has no filename set")
	}
	data, err := os.ReadFile(filepath.Join("templates", t.filename))
	if err != nil {
		panic(fmt.Sprintf("failed struct setter/getter template: %s: %v", t.filename, err))
	}
	datastr := string(data)
	datastr = strings.ReplaceAll(datastr, "#FIELD#", t.fieldName)
	datastr = strings.ReplaceAll(datastr, "#TYPE#", t.structName)
	datastr = strings.ReplaceAll(datastr, "#ENUMTYPE#", t.enumName)
	datastr = strings.ReplaceAll(datastr, "#SETTER_COMMENT#", t.setterComment)
	datastr = strings.ReplaceAll(datastr, "#GETTER_COMMENT#", t.getterComment)
	datastr = strings.ReplaceAll(datastr, "#CLEAR_COMMENT#", t.clearComment)
	datastr = strings.ReplaceAll(datastr, "#MAP#", t.mapName)
	datastr = strings.TrimSpace(datastr) + "\n\n"
	writer.Write([]byte(datastr))
}
