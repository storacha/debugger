package xagentmessage

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "xagentmessage",
	Short: "Tools for working with `X-Agent-Message` HTTP headers",
}

func init() {
	Cmd.AddCommand(parseCmd)
}
