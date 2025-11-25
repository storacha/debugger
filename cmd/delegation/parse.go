package delegation

import (
	"fmt"

	logging "github.com/ipfs/go-log/v2"
	"github.com/spf13/cobra"
	"github.com/storacha/debugger/pkg/ucanfmt"
	"github.com/storacha/go-ucanto/core/delegation"
)

var parseCmd = &cobra.Command{
	Use:   "parse <value>",
	Short: "Parse a delegation.",
	Long: `Parse a multibase encoded CID, with an identity multihash that contains delegation data in a CAR file.
`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		logging.SetLogLevel("*", "info")

		d, err := delegation.Parse(args[0])
		if err != nil {
			panic(fmt.Errorf("parsing delegation: %w", err))
		}

		ucanfmt.PrintDelegation(d)
	},
}
