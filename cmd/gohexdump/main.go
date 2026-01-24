package main

import (
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func main() {
	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", hex.Dump(bytes))
	os.Exit(0)
}
