package delegation

import (
	"fmt"

	logging "github.com/ipfs/go-log/v2"
	"github.com/spf13/cobra"
	"github.com/storacha/debugger/pkg/ipldfmt"
	"github.com/storacha/debugger/pkg/ucanfmt"
	"github.com/storacha/go-ucanto/core/delegation"
)

var parseCmd = &cobra.Command{
	Use:   "parse <value>",
	Short: "Parse a delegation.",
	Long:  `Parse a multibase encoded CID, with an identity multihash that contains delegation data in a CAR file.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		logging.SetLogLevel("*", "info")

		d, err := delegation.Parse(args[0])
		if err != nil {
			panic(fmt.Errorf("parsing delegation: %w", err))
		}

		jsonOutput, _ := cmd.Flags().GetBool("json")
		if jsonOutput {
			for b, err := range d.Blocks() {
				if err != nil {
					panic(fmt.Errorf("iterating delegation blocks: %w", err))
				}
				cmd.Printf("%s\n", b.Link())
				s, err := ipldfmt.FormatDagCBOR(b.Bytes())
				if err != nil {
					panic(fmt.Errorf("formatting block %s: %w", b.Link(), err))
				}
				cmd.Println(s)
				cmd.Println("")
			}
		} else {
			ucanfmt.PrintDelegation(d)
		}
	},
}

func init() {
	parseCmd.Flags().Bool("json", false, "Output DAG JSON")
	Cmd.AddCommand(parseCmd)
}
