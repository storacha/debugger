package peer

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "peer",
	Short: "Tools for debugging libp2p Peer IDs",
}
