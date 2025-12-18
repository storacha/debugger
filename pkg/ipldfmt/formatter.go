package ipldfmt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/alecthomas/chroma/v2/quick"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
	"golang.org/x/term"
)

// FormatDagCBOR formats dag-cbor encoded data to a dag-json encoded string.
func FormatDagCBOR(buf []byte) (string, error) {
	n, err := ipld.Decode(buf, dagcbor.Decode)
	if err != nil {
		return "", fmt.Errorf("decoding CBOR: %w", err)
	}
	return FormatNode(n, "")
}

// FormatNode formats an ipld-node to a dag-json encoded string.
func FormatNode(n ipld.Node, prefix string) (string, error) {
	jsonData, err := ipld.Encode(n, dagjson.Encode)
	if err != nil {
		return "", fmt.Errorf("encoding JSON: %w", err)
	}
	var indentedJSON bytes.Buffer
	err = json.Indent(&indentedJSON, jsonData, prefix, "  ")
	if err != nil {
		return "", fmt.Errorf("indenting JSON: %w", err)
	}
	if !term.IsTerminal(int(os.Stdout.Fd())) {
		return indentedJSON.String(), nil
	}
	var highlightedJSON bytes.Buffer
	err = quick.Highlight(&highlightedJSON, indentedJSON.String(), "json", "terminal16m", "doom-one2")
	if err != nil {
		panic(err)
	}
	return highlightedJSON.String(), nil
}
