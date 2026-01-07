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

	"gopkg.in/yaml.v3"
)

const (
	ConfigNFCV   = "config_nfcv.yaml"
	ConfigNoRoot = "config_noroot.yaml"
)

type ConfigYAML struct {
	MimeType           string `yaml:"mime_type"`
	Root               string `yaml:"root"`
	MetaFieldsFileName string `yaml:"meta_fields"`
	MainFieldsFileName string `yaml:"main_fields"`
	AuxFieldsFileName  string `yaml:"aux_fields"`
}

type Config struct {
	filename string
	yaml     ConfigYAML
	meta     []Field
	main     []Field
	aux      []Field
}

func (c Config) MetaFields() []Field {
	return c.meta
}

func (c Config) MainFields() []Field {
	return c.main
}

func (c Config) AuxFields() []Field {
	return c.aux
}

func (c Config) MimeType() string {
	return c.yaml.MimeType
}

func (c Config) Root() string {
	return c.yaml.Root
}

func GetConfig(filename string) (*Config, error) {
	c := new(Config)
	c.filename = filename
	return c, c.load()
}

func (c *Config) load() error {
	fn := dataFileName(c.filename)
	f, err := os.Open(fn)
	if err != nil {
		return fmt.Errorf("failed to find config %s: %w", c.filename, err)
	}
	defer f.Close()

	var y ConfigYAML
	err = yaml.NewDecoder(f).Decode(&y)
	if err != nil {
		return fmt.Errorf("failed to decode config %s: %w", c.filename, err)
	}
	c.yaml = y

	c.meta, err = loadFields(c.yaml.MetaFieldsFileName)
	if err != nil {
		return fmt.Errorf("failed to load meta fields: %w", err)
	}
	c.main, err = loadFields(c.yaml.MainFieldsFileName)
	if err != nil {
		return fmt.Errorf("failed to load main fields: %w", err)
	}
	c.aux, err = loadFields(c.yaml.AuxFieldsFileName)
	if err != nil {
		return fmt.Errorf("failed to load aux fields: %w", err)
	}

	return nil
}

func (c *Config) String() string {
	return fmt.Sprintf("Config: %s, Mime: %s, Root: %s, Meta: %s, Main: %s, Aux: %s",
		c.filename, c.yaml.MimeType, c.yaml.Root, c.yaml.MetaFieldsFileName, c.yaml.MainFieldsFileName, c.yaml.AuxFieldsFileName)
}
