package cmd

import (
	"github.com/spf13/cobra"
	"prismcloud.dev/cli/client"
	"prismcloud.dev/protobufs"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version of the apiserver",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := client.NewClient(apiAddr)
		cobra.CheckErr(err)

		defer client.Close()

		version, err := client.Api.Version(client.Ctx, &protobufs.Void{})
		cobra.CheckErr(err)

		cmd.Println("Prismcloud CLI 1.0.0 using Protocol v1")
		cmd.Printf("Prismcloud Cluster %v.%v.%v using Protocol v%v on Kubernetes %v\n", version.Major, version.Minor, version.Patch, version.Api, version.Kubernetes)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
