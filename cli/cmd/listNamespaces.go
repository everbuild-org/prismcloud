package cmd

import (
	"github.com/spf13/cobra"
	"prismcloud.dev/cli/client"
	"prismcloud.dev/protobufs"
)

var listNamespacesCmd = &cobra.Command{
	Use:   "list-namespace",
	Short: "List all namespaces",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := client.NewClient(apiAddr)
		cobra.CheckErr(err)

		defer c.Close()

		namespaces, err := c.Api.GetNamespaces(c.Ctx, &protobufs.Void{})
		cobra.CheckErr(err)

		cmd.Printf("Namespaces:\n")
		for _, namespace := range namespaces.Namespaces {
			cmd.Printf("- %v (%v Ram)\n", namespace.Name, client.FmtRam(namespace.RamLimit))
		}
	},
}

func init() {
	rootCmd.AddCommand(listNamespacesCmd)
}
