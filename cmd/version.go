package cmd

import (
	"fmt"

	"github.com/marc-barry/cfwctl/internal/version"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of cfwctl",
	Long:  `Version returns the current version of cfctl that you are running`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s", currentVersion())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func currentVersion() string {
	return version.CurrentVersion()
}
