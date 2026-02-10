package cid

import (
	"fmt"
	"os"

	"github.com/ipfs/go-cid"
	logging "github.com/ipfs/go-log/v2"
	"github.com/multiformats/go-multibase"
	"github.com/multiformats/go-multicodec"
	"github.com/multiformats/go-multihash"
	"github.com/spf13/cobra"
)

var decodeCmd = &cobra.Command{
	Use:   "decode <path>",
	Short: "Print a byte encoded CID.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		logging.SetLogLevel("*", "info")
		bytes, err := os.ReadFile(args[0])
		cobra.CheckErr(err)
		c, err := cid.Cast(bytes)
		cobra.CheckErr(err)
		printCID(c)
	},
}

func init() {
	Cmd.AddCommand(decodeCmd)
}

func printCID(c cid.Cid) {
	digestBaseStr, _ := multibase.EncodingToStr[multibase.Base58BTC]
	digestStr, _ := multibase.Encode(multibase.Base58BTC, c.Hash())
	mh, err := multihash.Decode(c.Hash())
	if err != nil {
		panic(fmt.Errorf("decoding multihash: %w", err))
	}
	if c.Prefix().Version == 0 {
		cidBaseStr := multibase.EncodingToStr[multibase.Base58BTC]
		cidStr := c.String()
		fmt.Printf("%-12s%s (%s)\n", "CIDv0:", cidStr, cidBaseStr)
		cv1 := cid.NewCidV1(uint64(multicodec.DagPb), c.Hash())
		cidv1BaseStr := multibase.EncodingToStr[multibase.Base32]
		cidv1Str, _ := cv1.StringOfBase(multibase.Base32)
		fmt.Printf("%-12s%s (%s)\n", "CIDv1:", cidv1Str, cidv1BaseStr)
	} else {
		cidBaseStr := multibase.EncodingToStr[multibase.Base32]
		cidStr, _ := c.StringOfBase(multibase.Base32)
		fmt.Printf("%-12s%s (%s)\n", "CID:", cidStr, cidBaseStr)
	}
	fmt.Printf("%-12s%d\n", "Version:", c.Prefix().Version)
	fmt.Printf("%-12s%#x (%s)\n", "IPLD Codec:", c.Prefix().Codec, multicodec.Code(c.Prefix().Codec).String())
	fmt.Printf("%-12s%s (%s)\n", "Digest:", digestStr, digestBaseStr)
	fmt.Printf("%-12s%#x (%s)\n", "Code:", mh.Code, multicodec.Code(mh.Code).String())
	fmt.Printf("%-12s%d\n", "Length:", mh.Length)
}
