package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/cloudflare/cloudflare-go"
	"github.com/spf13/cobra"
)

var workerRoutesCmd = &cobra.Command{
	Use:              "routes",
	Short:            "routes commands",
	Long:             `Routes commands that can be run`,
	TraverseChildren: true,
}

var workerRoutesCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create worker route",
	Long:  `Create a Worker route`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		api, err := newCfAPIClient(cloudflare.UsingAccount(cfAccountIDFlag))
		if err != nil {
			log.Fatalf("creating new Cloudflare API client: %s", err.Error())
		}
		enabled, err := cmd.Flags().GetBool(EnabledFlag)
		if err != nil {
			log.Fatalf("reading %s flag: %s", EnabledFlag, err.Error())
		}
		route := cloudflare.WorkerRoute{Pattern: args[0], Enabled: enabled, Script: args[1]}
		res, err := api.CreateWorkerRoute(cfZoneIDFlag, route)
		if err != nil {
			log.Fatalf("creating worker route: %s", err.Error())
		}
		b, err := json.MarshalIndent(res, "", " ")
		if err != nil {
			log.Fatalf("marshaling JSON: %s", err.Error())
		}
		fmt.Printf("%s", b)
	},
}

var workerRoutesListCmd = &cobra.Command{
	Use:   "list",
	Short: "list workers routes",
	Long:  `Fetch a list of routes for Workers`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		api, err := newCfAPIClient(cloudflare.UsingAccount(cfAccountIDFlag))
		if err != nil {
			log.Fatalf("creating new Cloudflare API client: %s", err.Error())
		}
		res, err := api.ListWorkerRoutes(cfZoneIDFlag)
		if err != nil {
			log.Fatalf("listing worker routes: %s", err.Error())
		}
		b, err := json.MarshalIndent(res.Routes, "", " ")
		if err != nil {
			log.Fatalf("marshaling JSON: %s", err.Error())
		}
		fmt.Printf("%s", b)
	},
}

func init() {
	workerCmd.AddCommand(workerRoutesCmd)
	workerRoutesCmd.PersistentFlags().StringVar(&cfZoneIDFlag, CfZoneIDFlag, "", "Cloudflare zone ID")
	workerRoutesCmd.MarkPersistentFlagRequired(CfZoneIDFlag)
	workerRoutesCmd.AddCommand(workerRoutesCreateCmd)
	workerRoutesCreateCmd.Flags().Bool(EnabledFlag, false, "Worker enabled")
	workerRoutesCmd.AddCommand(workerRoutesListCmd)
}
