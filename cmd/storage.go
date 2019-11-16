package cmd

import (
	"github.com/spf13/cobra"
)

// workerCmd represents the worker command
var storageCmd = &cobra.Command{
	Use:              "storage",
	Short:            "storage commands",
	Long:             `Storage commands that can be run`,
	TraverseChildren: true,
}

func init() {
	rootCmd.AddCommand(storageCmd)
	storageCmd.PersistentFlags().StringVar(&cfAPITokenFlag, CfAPITokenFlag, "", "Cloudflare API token")
	storageCmd.MarkPersistentFlagRequired(CfAPITokenFlag)
	storageCmd.PersistentFlags().StringVar(&cfAccountIDFlag, CfAccountIDFlag, "", "Cloudflare account ID")
	storageCmd.MarkPersistentFlagRequired(CfAccountIDFlag)
}
