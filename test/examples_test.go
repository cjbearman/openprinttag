package test

import (
	"fmt"
	"io"

	"github.com/cjbearman/openprinttag"
)

func x() {

	tag := openprinttag.NewOpenPrintTag().
		WithAuxRegionSize(32).
		WithSize(304)

	tag.MainRegion().
		SetBrandName("Awesome Filaments").
		SetMaterialName("Fancy PLA Yellow").
		SetPrimaryColor(openprinttag.MustNewColor("#FFFF00")).
		SetMaterialClass(openprinttag.MaterialClassFFF).
		SetMaterialType(openprinttag.MaterialTypePLA)

	// tagData, err := tag.Encode()

}

func readTag(r io.Reader) *openprinttag.OpenPrintTag {

	tagBytes, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}

	tag, err := openprinttag.Decode(tagBytes)
	if err != nil {
		panic("failed to read tag")
	}

	if brandName, found := tag.MainRegion().GetBrandName(); found {
		fmt.Printf("This tag is from brand: %s\n", brandName)
	} else {
		fmt.Printf("No brand name found in this tag")
	}

	return tag
}
