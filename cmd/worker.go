package cmd

import (
	"github.com/spf13/cobra"
)

// workerCmd represents the worker command
var workerCmd = &cobra.Command{
	Use:              "worker",
	Short:            "worker commands",
	Long:             `Worker commands that can be run`,
	TraverseChildren: true,
}

func init() {
	rootCmd.AddCommand(workerCmd)
	workerCmd.PersistentFlags().StringVar(&cfAccountIDFlag, CfAccountIDFlag, "", "Cloudflare account ID")
	workerCmd.MarkPersistentFlagRequired(CfAccountIDFlag)
}
