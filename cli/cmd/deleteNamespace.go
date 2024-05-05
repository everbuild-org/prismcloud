package cmd

import (
	"github.com/spf13/cobra"
	"prismcloud.dev/cli/client"
	"prismcloud.dev/protobufs"
)

var deleteNamespaceCmd = &cobra.Command{
	Use:   "delete-namespace [name]",
	Short: "Delete a namespace",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cobra.CheckErr(cmd.Help())
			return
		}

		name := args[0]
		c, err := client.NewClient(apiAddr)
		cobra.CheckErr(err)

		defer c.Close()

		_, err = c.Api.DeleteNamespace(c.Ctx, &protobufs.NamespaceDeleteRequest{
			Name: name,
		})

		cobra.CheckErr(err)

		cmd.Printf("Namespace '%v' deleted\n", name)
	},
}

func init() {
	rootCmd.AddCommand(deleteNamespaceCmd)
}
