package peer

import (
	"fmt"

	logging "github.com/ipfs/go-log/v2"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/spf13/cobra"
	ed25519 "github.com/storacha/go-ucanto/principal/ed25519/verifier"
)

var parseCmd = &cobra.Command{
	Use:   "parse <value>",
	Short: "Parse a Peer ID.",
	Long:  `Parse a Peer ID and display additional information where appropriate.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		logging.SetLogLevel("*", "info")

		peer, err := peer.Decode(args[0])
		cobra.CheckErr(err)

		fmt.Printf("PeerID:\t%s\n", peer.String())

		pk, err := peer.ExtractPublicKey()
		if err == nil {
			r, err := pk.Raw()
			cobra.CheckErr(err)

			v, err := ed25519.FromRaw(r)
			if err == nil {
				fmt.Printf("DID:\t%s\n", v.DID().String())
			}
		}
	},
}

func init() {
	Cmd.AddCommand(parseCmd)
}
