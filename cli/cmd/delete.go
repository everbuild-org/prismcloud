package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"prismcloud.dev/cli/resources"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [file.y(a)ml]",
	Short: "Delete a configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			err := cmd.Help()
			if err != nil {
				return
			}
			return
		}

		filename := args[0]

		all, err := os.ReadDir(filename)
		if err == nil {
			for _, f := range all {
				if f.IsDir() {
					continue
				}

				content, err := os.ReadFile(f.Name())
				cobra.CheckErr(err)

				text := string(content)
				err = resources.ParseAndActOnResourceFile(text, apiAddr, resources.Delete, currentNamespace)
				cobra.CheckErr(err)
			}

			return
		}

		content, err := os.ReadFile(filename)
		cobra.CheckErr(err)

		text := string(content)
		err = resources.ParseAndActOnResourceFile(text, apiAddr, resources.Delete, currentNamespace)
		cobra.CheckErr(err)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
