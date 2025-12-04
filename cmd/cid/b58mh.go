package cid

import (
	"github.com/ipfs/go-cid"
	logging "github.com/ipfs/go-log/v2"
	"github.com/multiformats/go-multibase"
	"github.com/spf13/cobra"
)

var b58mhCmd = &cobra.Command{
	Use:   "b58mh <cid>",
	Short: "Print the multibase base58btc encoded multihash from a CID.",
	Long:  "Extract the multihash from a CID and print the multibase base58btc encoded string.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		logging.SetLogLevel("*", "info")
		c, err := cid.Parse(args[0])
		cobra.CheckErr(err)
		str, _ := multibase.Encode(multibase.Base58BTC, c.Hash())
		cmd.Println(str)
	},
}

func init() {
	Cmd.AddCommand(b58mhCmd)
}
