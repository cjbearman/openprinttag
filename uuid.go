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

import (
	"crypto/sha1"
	"errors"

	"github.com/google/uuid"
)

const uuidVersion = 5

var (
	brandNamespace    = uuid.Must(uuid.Parse("5269dfb7-1559-440a-85be-aba5f3eff2d2"))
	materialNamespace = uuid.Must(uuid.Parse("616fc86d-7d99-4953-96c7-46d2836b9be9"))
	packageNamespace  = uuid.Must(uuid.Parse("6f7d485e-db8d-4979-904e-a231cd6602b2"))
	instanceNamespace = uuid.Must(uuid.Parse("31062f81-b5bd-4f86-a5f8-46367e841508"))
)

// BrandUUID returns a properly formed brand UUID for the given brand name
func BrandUUID(brandName string) uuid.UUID {
	return uuid.NewHash(sha1.New(), brandNamespace, []byte(brandName), uuidVersion)
}

// MaterialUUID returns a properly formed material UUID for the given material name.
// The associated brand UUID must also be passed
func MaterialUUID(materialName string, brandUUID uuid.UUID) uuid.UUID {
	content := append(brandUUID[:], []byte(materialName)...)
	return uuid.NewHash(sha1.New(), materialNamespace, content, uuidVersion)
}

// MaterialPackageUUID returns a properly formed package ID from the gtin.
// The associated brand UUID must also be passed
func MaterialPackageUUID(gtin string, brandUUID uuid.UUID) uuid.UUID {
	content := append(brandUUID[:], []byte(gtin)...)
	return uuid.NewHash(sha1.New(), packageNamespace, content, uuidVersion)
}

// MaterialPackageInstanceUUID returns a properly formed instance UUID from
// the NFC tag UID.
// Errors are generated if the NFC Tag UID is not proper by spec (8 bytes starting 0x0e)
func MaterialPackageInstanceUUID(NFCTagUid []byte) (uuid.UUID, error) {
	if len(NFCTagUid) != 8 {
		return uuid.UUID{}, errors.New("NFC Tag UID should be 8 bytes long")
	}
	if NFCTagUid[0] != 0x0e {
		return uuid.UUID{}, errors.New("NFC Tag UUID should start with 0x0E")
	}

	return uuid.NewHash(sha1.New(), instanceNamespace, NFCTagUid, uuidVersion), nil
}
