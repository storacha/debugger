package ipni

import (
	"github.com/spf13/cobra"
	"github.com/storacha/debugger/cmd/ipni/metadata"
)

var Cmd = &cobra.Command{
	Use:   "ipni",
	Short: "Tools for debugging IPNI",
}

func init() {
	Cmd.AddCommand(metadata.Cmd)
}
