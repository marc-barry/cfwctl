package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/cloudflare/cloudflare-go"
	"github.com/marc-barry/cfwctl/internal/utility"
	"github.com/spf13/cobra"
)

var workerScriptCmd = &cobra.Command{
	Use:              "script",
	Short:            "script commands",
	Long:             `Script commands that can be run`,
	TraverseChildren: true,
}

var workerScriptDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete worker",
	Long:  `Delete a Worker`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		api, err := newCfAPIClient(cloudflare.UsingAccount(cfAccountIDFlag))
		if err != nil {
			log.Fatalf("creating new Cloudflare API client: %s", err.Error())
		}
		res, err := api.DeleteWorker(&cloudflare.WorkerRequestParams{ZoneID: cfZoneIDFlag, ScriptName: args[0]})
		if err != nil {
			log.Fatalf("deleting Worker %s", err.Error())
		}
		b, err := json.MarshalIndent(res, "", " ")
		if err != nil {
			log.Fatalf("marshaling JSON: %s", err.Error())
		}
		fmt.Printf("%s", b)
	},
}

var workerScriptDownloadCmd = &cobra.Command{
	Use:   "download",
	Short: "download worker",
	Long:  `Download a Worker`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		api, err := newCfAPIClient(cloudflare.UsingAccount(cfAccountIDFlag))
		if err != nil {
			log.Fatalf("creating new Cloudflare API client: %s", err.Error())
		}
		res, err := api.DownloadWorker(&cloudflare.WorkerRequestParams{ZoneID: cfZoneIDFlag, ScriptName: args[0]})
		if err != nil {
			log.Fatalf("downloading Worker %s", err.Error())
		}
		b, err := json.MarshalIndent(res.WorkerScript, "", " ")
		if err != nil {
			log.Fatalf("marshaling JSON: %s", err.Error())
		}
		fmt.Printf("%s", b)
	},
}

var workerScriptListCmd = &cobra.Command{
	Use:   "list",
	Short: "list workers",
	Long:  `Fetch a list of uploaded Workers`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		api, err := newCfAPIClient(cloudflare.UsingAccount(cfAccountIDFlag))
		if err != nil {
			log.Fatalf("creating new Cloudflare API client: %s", err.Error())
		}
		res, err := api.ListWorkerScripts()
		if err != nil {
			log.Fatalf("listing worker scripts: %s", err.Error())
		}
		b, err := json.MarshalIndent(res.WorkerList, "", " ")
		if err != nil {
			log.Fatalf("marshaling JSON: %s", err.Error())
		}
		fmt.Printf("%s", b)
	},
}

var workerScriptUploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "upload worker",
	Long:  `Upload a Worker`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		api, err := newCfAPIClient(cloudflare.UsingAccount(cfAccountIDFlag))
		if err != nil {
			log.Fatalf("creating new Cloudflare API client: %s", err.Error())
		}
		if !utility.FileExists(args[1]) {
			log.Fatalf("%s is not a file", args[1])
		}
		script, err := ioutil.ReadFile(args[1])
		if err != nil {
			log.Fatalf("reading script file %s", err.Error())
		}
		res, err := api.UploadWorker(&cloudflare.WorkerRequestParams{ZoneID: cfZoneIDFlag, ScriptName: args[0]}, string(script))
		if err != nil {
			log.Fatalf("uploading Worker %s", err.Error())
		}
		b, err := json.MarshalIndent(res, "", " ")
		if err != nil {
			log.Fatalf("marshaling JSON: %s", err.Error())
		}
		fmt.Printf("%s", b)
	},
}

func init() {
	workerCmd.AddCommand(workerScriptCmd)
	workerScriptCmd.AddCommand(workerScriptDeleteCmd)
	workerScriptCmd.AddCommand(workerScriptDownloadCmd)
	workerScriptCmd.AddCommand(workerScriptListCmd)
	workerScriptCmd.AddCommand(workerScriptUploadCmd)
}
