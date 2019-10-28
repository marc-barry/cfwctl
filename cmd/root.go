package cmd

import (
	"fmt"

	"github.com/cloudflare/cloudflare-go"
	"github.com/spf13/cobra"
)

var (
	// Used for flags
	cfAPITokenFlag  string
	cfAccountIDFlag string
	cfZoneIDFlag    string

	rootCmd = &cobra.Command{
		Use:              "cfwctl",
		Short:            "A CLI for interacting with Cloudflare Workers",
		Long:             "cfwctl is a CLI for interacting with Cloudflare Workers",
		TraverseChildren: true,
	}
)

func newCfAPIClient(opts ...cloudflare.Option) (*cloudflare.API, error) {
	if cfAPITokenFlag != "" {
		return cloudflare.NewWithAPIToken(cfAPITokenFlag, opts...)
	}

	return nil, fmt.Errorf("Cloudflare API token must not be empty")
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
