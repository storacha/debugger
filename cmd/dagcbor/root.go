package dagcbor

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "dagcbor",
	Short: "Tools for debugging dag-cbor data",
}
