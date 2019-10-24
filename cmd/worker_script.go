package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/cloudflare/cloudflare-go"
	"github.com/spf13/cobra"
)

// workerCmd represents the worker command
var workerScriptCmd = &cobra.Command{
	Use:              "script",
	Short:            "script commands",
	Long:             `Script commands that can be run`,
	TraverseChildren: true,
}

var workerScriptListCmd = &cobra.Command{
	Use:   "list",
	Short: "list workers",
	Long:  `Fetch a list of uploaded workers`,
	Run: func(cmd *cobra.Command, args []string) {
		api, err := newCfAPIClient(cloudflare.UsingAccount(cfAccountIDFlag))
		if err != nil {
			log.Fatalf("creating new Cloudflare API client: %s", err.Error())
		}
		res, err := api.ListWorkerScripts()
		if err != nil {
			log.Fatalf("listing worker scripts: %s", err.Error())
		}
		b, err := json.Marshal(res.WorkerList)
		if err != nil {
			log.Fatalf("marshaling JSON: %s", err.Error())
		}
		fmt.Printf("%s", b)
	},
}

func init() {
	workerCmd.AddCommand(workerScriptCmd)
	workerScriptCmd.AddCommand(workerScriptListCmd)
}
