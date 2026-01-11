package main

import (
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/cjbearman/openprinttag/cmd/tagtool/vtag"
)

func terminal(note string, err error) {
	if err != nil {
		fmt.Printf("TERMINAL: %s: %v\n", note, err)
		os.Exit(1)
	}
}

type nilCloser struct{}

func (n *nilCloser) Close() error { return nil }

func ioHandle(filename string, write bool) (*os.File, io.Closer, error) {
	if filename == "-" {
		nc := &nilCloser{}
		if write {
			return os.Stdout, nc, nil
		}
		return os.Stdin, nc, nil
	}
	if write {
		f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		return f, f, err

	}
	f, err := os.Open(filename)
	return f, f, err
}

func countOpts(opts ...any) (count int) {
	for _, o := range opts {
		switch v := o.(type) {
		case bool:
			if v {
				count++
			}
		case string:
			if v != "" {
				count++
			}
		}
	}
	return count
}

func main() {

	var debug, id, dumpHex bool
	var read, write string
	var nbytes int

	flag.StringVar(&read, "r", "", "Read data to file or (- stdout)")
	flag.StringVar(&write, "w", "", "Write data from file or (- stdout)")
	flag.BoolVar(&debug, "d", false, "Debug")
	flag.BoolVar(&id, "i", false, "ID the chip")
	flag.IntVar(&nbytes, "n", 0, "Number of bytes to read (if unspecified, entire memory is read)")
	flag.BoolVar(&dumpHex, "hex", false, "Hex dump")
	flag.Parse()

	var err error

	nOpts := countOpts(id, read, write)
	if nOpts == 0 {
		fmt.Fprintf(os.Stderr, "Must use one of -i, -r, -w\n")
		os.Exit(1)
	} else if nOpts > 1 {
		fmt.Fprintf(os.Stderr, "-r, -w, -i can be used at a time\n")
		os.Exit(1)
	}

	vtag.DebugMode = debug
	scanner := vtag.NewScanner()

	defer scanner.Close()

	err = scanner.OnCard(func(session *vtag.Session) error {
		if id {
			uid := session.GetTag().GetUIDHex()
			terminal("uid", err)
			fmt.Printf("UID: %s, Type: %s\n", uid, session.GetTag().GetTagType().String())
		}

		if read != "" {
			f, closer, err := ioHandle(read, true)
			if err != nil {
				terminal("open-read", err)
			}
			defer closer.Close()
			if nbytes == 0 {
				nbytes = session.GetTag().GetAvailableBytes()
			}
			data, err := session.Read(0, nbytes)
			terminal("read", err)
			if dumpHex {
				data = []byte(hex.Dump(data))
			}
			_, err = f.Write(data)
			terminal("write-data", err)
		}
		if write != "" {
			f, closer, err := ioHandle(write, false)
			if err != nil {
				terminal("open-read", err)
			}
			defer closer.Close()
			data, err := io.ReadAll(f)
			terminal("read-data", err)
			if len(data) > int(session.GetTag().GetAvailableBytes()) {
				terminal("data-size", errors.New("data too large for tag"))
			}
			err = session.Write(0, data)
			terminal("write-data-to-tag", err)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}
