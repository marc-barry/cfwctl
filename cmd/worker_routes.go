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
		route := cloudflare.WorkerRoute{Pattern: args[0], Script: args[1]}
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

var workerRoutesDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete worker route",
	Long:  `Delete a Worker route`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		api, err := newCfAPIClient(cloudflare.UsingAccount(cfAccountIDFlag))
		if err != nil {
			log.Fatalf("creating new Cloudflare API client: %s", err.Error())
		}
		res, err := api.DeleteWorkerRoute(cfZoneIDFlag, args[0])
		if err != nil {
			log.Fatalf("deleting Worker route %s", err.Error())
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

var workerRoutesUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "update worker route",
	Long:  `Update a Worker route`,
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		api, err := newCfAPIClient(cloudflare.UsingAccount(cfAccountIDFlag))
		if err != nil {
			log.Fatalf("creating new Cloudflare API client: %s", err.Error())
		}
		route := cloudflare.WorkerRoute{Pattern: args[1], Script: args[2]}
		res, err := api.UpdateWorkerRoute(cfZoneIDFlag, args[0], route)
		if err != nil {
			log.Fatalf("updating worker route: %s", err.Error())
		}
		b, err := json.MarshalIndent(res, "", " ")
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
	workerRoutesCmd.AddCommand(workerRoutesDeleteCmd)
	workerRoutesCmd.AddCommand(workerRoutesListCmd)
	workerRoutesCmd.AddCommand(workerRoutesUpdateCmd)
}
