package ipldfmt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/efekarakus/termcolor"

	"github.com/alecthomas/chroma/v2/quick"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
	"golang.org/x/term"
)

func detectFormatter(fd termcolor.FileDescriptor) (string, error) {
	switch l := termcolor.SupportLevel(fd); l {
	case termcolor.Level16M:
		return "terminal16m", nil
	case termcolor.Level256:
		return "terminal256", nil
	case termcolor.LevelBasic:
		return "terminal16", nil
	default:
		return "", fmt.Errorf("no terminal color support detected")
	}
}

// FormatDagCBOR formats dag-cbor encoded data to a dag-json encoded string.
func FormatDagCBOR(buf []byte) (string, error) {
	n, err := ipld.Decode(buf, dagcbor.Decode)
	if err != nil {
		return "", fmt.Errorf("decoding CBOR: %w", err)
	}
	jsonData, err := ipld.Encode(n, dagjson.Encode)
	if err != nil {
		return "", fmt.Errorf("encoding JSON: %w", err)
	}
	var indentedJSON bytes.Buffer
	err = json.Indent(&indentedJSON, jsonData, "", "  ")
	if err != nil {
		return "", fmt.Errorf("indenting JSON: %w", err)
	}
	if !term.IsTerminal(int(os.Stdout.Fd())) {
		return indentedJSON.String(), nil
	}
	formatter, err := detectFormatter(os.Stderr)
	if err != nil {
		return indentedJSON.String(), nil
	}
	var highlightedJSON bytes.Buffer
	err = quick.Highlight(&highlightedJSON, indentedJSON.String(), "json", formatter, "doom-one2")
	if err != nil {
		panic(err)
	}
	return highlightedJSON.String(), nil
}
