package delegation

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "delegation",
	Short: "Tools for debugging UCAN delegations",
}
