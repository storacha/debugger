package message

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "message",
	Short: "Tools for debugging UCAN messages",
}
