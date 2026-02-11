package dagcbor

import (
	"os"

	logging "github.com/ipfs/go-log/v2"
	"github.com/spf13/cobra"
	"github.com/storacha/debugger/pkg/ipldfmt"
)

var decodeCmd = &cobra.Command{
	Use:   "decode <path>",
	Short: "Print a dag-cbor data structure.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		logging.SetLogLevel("*", "info")
		bytes, err := os.ReadFile(args[0])
		cobra.CheckErr(err)
		out, err := ipldfmt.FormatDagCBOR(bytes)
		cobra.CheckErr(err)
		cmd.Println(out)
	},
}

func init() {
	Cmd.AddCommand(decodeCmd)
}
