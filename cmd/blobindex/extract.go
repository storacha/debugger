package blobindex

import (
	"bytes"
	"fmt"
	"io"
	"os"

	logging "github.com/ipfs/go-log/v2"
	"github.com/spf13/cobra"
	"github.com/storacha/debugger/pkg/ipldfmt"
	"github.com/storacha/go-libstoracha/blobindex"
	"github.com/storacha/go-libstoracha/digestutil"
	"github.com/storacha/go-ucanto/core/car"
)

var extractCmd = &cobra.Command{
	Use:   "extract [car-file]",
	Short: "Extract a sharded DAG index from a CAR.",
	Long:  "Extract a sharded DAG index that has been archived to a CAR. You can pipe directly to this command.",
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

		jsonOutput, _ := cmd.Flags().GetBool("json")
		if jsonOutput {
			_, blocks, err := car.Decode(bytes.NewReader(archive))
			if err != nil {
				panic(fmt.Errorf("decoding CAR: %w", err))
			}
			for b, err := range blocks {
				if err != nil {
					panic(fmt.Errorf("iterating blob index blocks: %w", err))
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
			index, err := blobindex.Extract(bytes.NewReader(archive))
			if err != nil {
				panic(fmt.Errorf("extracting blob index: %w", err))
			}
			cmd.Println("Content:")
			cmd.Printf("  %s\n", index.Content())
			fmt.Printf("Shards (%d):\n", index.Shards().Size())
			for shard, slices := range index.Shards().Iterator() {
				fmt.Printf("  %s\n", digestutil.Format(shard))
				fmt.Printf("    Slices (%d):\n", slices.Size())
				for digest, position := range slices.Iterator() {
					fmt.Printf("      %s @ %d-%d\n", digestutil.Format(digest), position.Offset, position.Offset+position.Length-1)
				}
			}
		}
	},
}

func init() {
	extractCmd.Flags().Bool("json", false, "Output DAG JSON")
	Cmd.AddCommand(extractCmd)
}
