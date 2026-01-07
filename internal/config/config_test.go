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
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfigLoad(t *testing.T) {
	require := require.New(t)
	cfg, err := GetConfig(ConfigNFCV)
	require.NoError(err)
	require.NotNil(cfg)
	t.Logf("Config Object: %s", cfg)

	dumpFields(t, cfg)
}

func TestTagCategoriesEnumeration(t *testing.T) {
	for _, item := range TagCategories() {
		t.Logf("%s", item)
	}
}

func dumpFields(t *testing.T, cfg *Config) {

	t.Logf("Meta Fields:")
	for _, item := range cfg.MetaFields() {
		t.Logf("%s", item)
		dumpEnumerations(t, item)
	}
	t.Logf("Main Fields:")
	for _, item := range cfg.MainFields() {
		t.Logf("%s", item)
		dumpEnumerations(t, item)
	}
	t.Logf("Aux Fields:")
	for _, item := range cfg.AuxFields() {
		t.Logf("%s", item)
		dumpEnumerations(t, item)
	}
}

func dumpEnumerations(t *testing.T, f Field) {
	if f.HasEnumeration() {
		for _, item := range f.EnumeratedValues() {
			t.Logf("%s", item)
		}
	}
}
