package parse

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "parse",
	Short: "Parse domain objects",
}

func init() {
	Cmd.AddCommand(xagentmessageCmd)
}
