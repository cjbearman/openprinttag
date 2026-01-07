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

// ** THIS FILE IS AUTO-GENERATED, DO NOT MODIFY **

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

type MaterialCertifications uint64

const (
	// MaterialCertificationsUl2818
	MaterialCertificationsUl2818 MaterialCertifications = 0

	// MaterialCertificationsUl94V0
	MaterialCertificationsUl94V0 MaterialCertifications = 1

	// MaterialCertificationsUl2904
	// Certifies that a 3D printing filament produces VOC and ultrafine particle emissions below safe thresholds when printed, making it safer for indoor use.
	MaterialCertificationsUl2904 MaterialCertifications = 2
)

var MaterialCertificationsMap = map[uint64]string{
	0: "ul_2818",
	1: "ul_94_v0",
	2: "ul_2904",
}

func (e MaterialCertifications) String() string {
	return MaterialCertificationsMap[uint64(e)]
}

func (e MaterialCertifications) MarshalYAML() (any, error) {
	if str, ok := MaterialCertificationsMap[uint64(e)]; ok {
		return str, nil
	}
	return nil, fmt.Errorf("unknown enumeration: %d", e)
}

func (e *MaterialCertifications) UnmarshalYAML(value *yaml.Node) error {
	var str string
	if err := value.Decode(&str); err != nil {
		return err
	}

	// Hardly efficient, but this is not critical here
	for key, name := range MaterialCertificationsMap {
		if name == str {
			*e = MaterialCertifications(key)
			return nil
		}
	}
	return fmt.Errorf("unknown enumeration: %s", str)
}
