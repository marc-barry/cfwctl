package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/cloudflare/cloudflare-go"
	"github.com/spf13/cobra"
)

var kvCmd = &cobra.Command{
	Use:              "kv",
	Short:            "kv commands",
	Long:             `KV commands that can be run`,
	TraverseChildren: true,
}

var namespacesCmd = &cobra.Command{
	Use:              "namespaces",
	Short:            "namespaces commands",
	Long:             `Namespaces commands that can be run`,
	TraverseChildren: true,
}

var namespacesCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create namespace",
	Long:  `Create a Namespace`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		api, err := newCfAPIClient(cloudflare.UsingAccount(cfAccountIDFlag))
		if err != nil {
			log.Fatalf("creating new Cloudflare API client: %s", err.Error())
		}
		req := &cloudflare.WorkersKVNamespaceRequest{Title: args[0]}
		res, err := api.CreateWorkersKVNamespace(context.Background(), req)
		if err != nil {
			log.Fatalf("creating namespace: %s", err.Error())
		}
		b, err := json.MarshalIndent(res, "", " ")
		if err != nil {
			log.Fatalf("marshaling JSON: %s", err.Error())
		}
		fmt.Printf("%s", b)
	},
}

var namespacesDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete namespace",
	Long:  `Delete a Namespace`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		api, err := newCfAPIClient(cloudflare.UsingAccount(cfAccountIDFlag))
		if err != nil {
			log.Fatalf("creating new Cloudflare API client: %s", err.Error())
		}
		res, err := api.DeleteWorkersKVNamespace(context.Background(), args[0])
		if err != nil {
			log.Fatalf("deleting namespace %s", err.Error())
		}
		b, err := json.MarshalIndent(res, "", " ")
		if err != nil {
			log.Fatalf("marshaling JSON: %s", err.Error())
		}
		fmt.Printf("%s", b)
	},
}

var namespacesListCmd = &cobra.Command{
	Use:   "list",
	Short: "list namespaces",
	Long:  `Fetch a list of Namespaces`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		api, err := newCfAPIClient(cloudflare.UsingAccount(cfAccountIDFlag))
		if err != nil {
			log.Fatalf("creating new Cloudflare API client: %s", err.Error())
		}
		res, err := api.ListWorkersKVNamespaces(context.Background())
		if err != nil {
			log.Fatalf("listing namespaces: %s", err.Error())
		}
		b, err := json.MarshalIndent(res.Result, "", " ")
		if err != nil {
			log.Fatalf("marshaling JSON: %s", err.Error())
		}
		fmt.Printf("%s", b)
	},
}

var namespacesUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "update namespace",
	Long:  `Update a Namepsaces`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		api, err := newCfAPIClient(cloudflare.UsingAccount(cfAccountIDFlag))
		if err != nil {
			log.Fatalf("creating new Cloudflare API client: %s", err.Error())
		}
		req := &cloudflare.WorkersKVNamespaceRequest{Title: args[1]}
		res, err := api.UpdateWorkersKVNamespace(context.Background(), args[0], req)
		if err != nil {
			log.Fatalf("updating namespace: %s", err.Error())
		}
		b, err := json.MarshalIndent(res, "", " ")
		if err != nil {
			log.Fatalf("marshaling JSON: %s", err.Error())
		}
		fmt.Printf("%s", b)
	},
}

func init() {
	storageCmd.AddCommand(kvCmd)
	kvCmd.AddCommand(namespacesCmd)
	namespacesCmd.AddCommand(namespacesCreateCmd)
	namespacesCmd.AddCommand(namespacesDeleteCmd)
	namespacesCmd.AddCommand(namespacesListCmd)
	namespacesCmd.AddCommand(namespacesUpdateCmd)
}
