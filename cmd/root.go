package cmd

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/storacha/debugger/cmd/blobindex"
	"github.com/storacha/debugger/cmd/cid"
	"github.com/storacha/debugger/cmd/delegation"
	"github.com/storacha/debugger/cmd/flatfs"
	"github.com/storacha/debugger/cmd/ipni"
	"github.com/storacha/debugger/cmd/xagentmessage"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "debugger",
	Short: "A debugger for Storacha",
	Long:  `Various tools and commands that can help debugging the Storacha Network.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	_ = godotenv.Load()

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(cid.Cmd)
	rootCmd.AddCommand(blobindex.Cmd)
	rootCmd.AddCommand(delegation.Cmd)
	rootCmd.AddCommand(flatfs.Cmd)
	rootCmd.AddCommand(ipni.Cmd)
	rootCmd.AddCommand(xagentmessage.Cmd)
}
