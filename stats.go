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

type RootStat struct {
	DataSize        int `yaml:"data_size"`
	PayloadSize     int `yaml:"payload_size"`
	Overhead        int `yaml:"overhead"`
	PayloadUsedSize int `yaml:"payload_used_size"`
	TotalUsedSize   int `yaml:"total_used_size"`
}

type RegionStat struct {
	PayloadOffset  int `yaml:"payload_offset"`
	AbsoluteOffset int `yaml:"absolute_offset"`
	Size           int `yaml:"size"`
	UsedSize       int `yaml:"used_size"`
}

type Stats struct {
	// Root provides information about the root tag
	Root RootStat

	// Meta provides information about the meta region
	Meta RegionStat

	// Main provides information about the main region
	Main RegionStat

	// Aux provides information about the aux region and will be nil
	// if there is no aux region
	Aux *RegionStat
}

// GetStats will return the stat data from the last
// encode operation and true
// If the tag has not been encoded, it will return nil, false
func (o *OpenPrintTag) GetStats() (*Stats, bool) {
	return o.stats, o.stats != nil
}
