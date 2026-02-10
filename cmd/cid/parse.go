package cid

import (
	"github.com/ipfs/go-cid"
	logging "github.com/ipfs/go-log/v2"
	"github.com/spf13/cobra"
)

var parseCmd = &cobra.Command{
	Use:   "parse <path>",
	Short: "Print a string encoded CID.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		logging.SetLogLevel("*", "info")
		c, err := cid.Parse(args[0])
		cobra.CheckErr(err)
		printCID(c)
	},
}

func init() {
	Cmd.AddCommand(parseCmd)
}
