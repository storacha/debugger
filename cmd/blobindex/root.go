package blobindex

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "blobindex",
	Short: "Tools for debugging blob indexes",
}
