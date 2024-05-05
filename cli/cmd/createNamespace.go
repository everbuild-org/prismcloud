package cmd

import (
	"github.com/spf13/cobra"
	"prismcloud.dev/cli/client"
	"prismcloud.dev/protobufs"
)

var namespaceCreateRam string

// createNamespaceCmd represents the createNamespace command
var createNamespaceCmd = &cobra.Command{
	Use:   "create-namespace",
	Short: "Create a namespace",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cobra.CheckErr(cmd.Help())
			return
		}

		name := args[0]
		c, err := client.NewClient(apiAddr)
		cobra.CheckErr(err)

		defer c.Close()

		ramInt, err := client.ParseRam(namespaceCreateRam)
		cobra.CheckErr(err)

		namespace, err := c.Api.CreateNamespace(c.Ctx, &protobufs.NamespaceCreateRequest{
			Name:     name,
			RamLimit: ramInt,
		})

		cobra.CheckErr(err)

		cmd.Printf("Namespace '%v' created with a RAM Limit of %v\n", namespace.Name, client.FmtRam(namespace.RamLimit))
	},
}

func init() {
	rootCmd.AddCommand(createNamespaceCmd)

	createNamespaceCmd.Flags().StringVarP(&namespaceCreateRam, "ram", "r", "1Gi", "The amount of RAM to allocate to the namespace")
}
