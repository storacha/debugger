package message

import (
	"bytes"
	"fmt"

	logging "github.com/ipfs/go-log/v2"
	"github.com/multiformats/go-multibase"
	"github.com/spf13/cobra"
	"github.com/storacha/debugger/pkg/ipldfmt"
	"github.com/storacha/debugger/pkg/ucanfmt"
	"github.com/storacha/go-ucanto/core/car"
	"github.com/storacha/go-ucanto/core/dag/blockstore"
	"github.com/storacha/go-ucanto/core/message"
)

var parseCmd = &cobra.Command{
	Use:   "parse <value>",
	Short: "Parse a message.",
	Long:  `Parse multibase encoded agent message data in a CAR file.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		logging.SetLogLevel("*", "info")

		// Decode the multibase encoded value
		_, data, err := multibase.Decode(args[0])
		if err != nil {
			panic(fmt.Errorf("decoding multibase: %w", err))
		}

		// Decode CAR file
		roots, blocks, err := car.Decode(bytes.NewReader(data))
		if err != nil {
			panic(fmt.Errorf("decoding CAR: %w", err))
		}
		if len(roots) != 1 {
			panic(fmt.Errorf("unexpected number of roots: %d, expected: 1", len(roots)))
		}

		// Create blockstore from blocks
		bstore, err := blockstore.NewBlockReader(blockstore.WithBlocksIterator(blocks))
		if err != nil {
			panic(fmt.Errorf("creating blockstore: %w", err))
		}

		// Create message from root and blockstore
		msg, err := message.NewMessage(roots[0], bstore)
		if err != nil {
			panic(fmt.Errorf("creating message: %w", err))
		}

		jsonOutput, _ := cmd.Flags().GetBool("json")
		if jsonOutput {
			for b, err := range msg.Blocks() {
				if err != nil {
					panic(fmt.Errorf("iterating message blocks: %w", err))
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
			ucanfmt.PrintMessage(msg)
		}
	},
}

func init() {
	parseCmd.Flags().Bool("json", false, "Output DAG JSON")
	Cmd.AddCommand(parseCmd)
}
