package did

import (
	"fmt"
	"strings"

	logging "github.com/ipfs/go-log/v2"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/spf13/cobra"
	"github.com/storacha/go-ucanto/did"
	ed25519 "github.com/storacha/go-ucanto/principal/ed25519/verifier"
)

var parseCmd = &cobra.Command{
	Use:   "parse <value>",
	Short: "Parse a DID.",
	Long:  `Parse a DID and display additional information where appropriate.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		logging.SetLogLevel("*", "info")

		d, err := did.Parse(args[0])
		cobra.CheckErr(err)

		fmt.Printf("DID:\t%s\n", d.String())

		if strings.HasPrefix(d.String(), "did:key:") {
			v, err := ed25519.Parse(d.String())
			if err == nil {
				pub, err := crypto.UnmarshalEd25519PublicKey(v.Raw())
				cobra.CheckErr(err)
				peer, err := peer.IDFromPublicKey(pub)
				cobra.CheckErr(err)
				fmt.Printf("PeerID:\t%s\n", peer.String())
			}
		}
	},
}

func init() {
	Cmd.AddCommand(parseCmd)
}
