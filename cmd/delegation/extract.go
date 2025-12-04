package delegation

import (
	"fmt"
	"io"
	"os"

	logging "github.com/ipfs/go-log/v2"
	"github.com/spf13/cobra"
	"github.com/storacha/debugger/pkg/ipldfmt"
	"github.com/storacha/debugger/pkg/ucanfmt"
	"github.com/storacha/go-ucanto/core/delegation"
)

var extractCmd = &cobra.Command{
	Use:   "extract [car-file]",
	Short: "Extract a delegation from a CAR.",
	Long:  "Extract a delegation that has been archived to a CAR. You can pipe directly to this command.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		logging.SetLogLevel("*", "info")

		var archive []byte
		var err error
		if len(args) > 0 {
			archive, err = os.ReadFile(args[0])
		} else {
			archive, err = io.ReadAll(cmd.InOrStdin())
		}
		if err != nil {
			panic(err)
		}

		d, err := delegation.Extract(archive)
		if err != nil {
			panic(fmt.Errorf("extracting delegation: %w", err))
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
	extractCmd.Flags().Bool("json", false, "Output DAG JSON")
	Cmd.AddCommand(extractCmd)
}
