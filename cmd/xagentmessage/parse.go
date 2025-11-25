package xagentmessage

import (
	"fmt"

	logging "github.com/ipfs/go-log/v2"
	"github.com/spf13/cobra"
	"github.com/storacha/debugger/pkg/ipldfmt"
	"github.com/storacha/go-ucanto/transport/headercar/message"
)

var parseCmd = &cobra.Command{
	Use:   "parse <value>",
	Short: "Parse an X-Agent-Message.",
	Long: `Parse a multibase encoded, gzipped, X-Agent-Message header.
`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		logging.SetLogLevel("*", "info")

		msg, err := message.DecodeHeader(args[0])
		if err != nil {
			panic(fmt.Errorf("decoding X-Agent-Message header: %w", err))
		}
		cmd.Printf("successfully decoded X-Agent-Message header (%d bytes)\n", len(args[0]))
		if len(msg.Invocations()) != 1 {
			panic(fmt.Errorf("unexpected number of invocations in message: %d", len(msg.Invocations())))
		}

		cmd.Println("blocks:")
		for b, err := range msg.Blocks() {
			if err != nil {
				panic(fmt.Errorf("iterating invocation blocks: %w", err))
			}
			cmd.Printf("%s\n", b.Link())
			s, err := ipldfmt.FormatDagCBOR(b.Bytes())
			if err != nil {
				panic(fmt.Errorf("formatting block %s: %w", b.Link(), err))
			}
			cmd.Println(s)
			cmd.Println("")
		}
	},
}
