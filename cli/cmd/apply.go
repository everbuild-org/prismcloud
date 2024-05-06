package cmd

import (
	"os"
	"time"

	"github.com/spf13/cobra"
	"prismcloud.dev/cli/resources"
)

var applyForce bool

var applyCmd = &cobra.Command{
	Use:   "apply [file.y(a)ml]",
	Short: "Apply a configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			err := cmd.Help()
			if err != nil {
				return
			}
			return
		}

		filename := args[0]

		contents := make([]string, 0)

		all, err := os.ReadDir(filename)
		if err == nil {
			for _, f := range all {
				if f.IsDir() {
					continue
				}

				content, err := os.ReadFile(f.Name())
				cobra.CheckErr(err)

				text := string(content)
				contents = append(contents, text)
			}
		} else {
			content, err := os.ReadFile(filename)
			cobra.CheckErr(err)

			text := string(content)
			contents = append(contents, text)
		}

		if applyForce {
			for _, text := range contents {
				_ = resources.ParseAndActOnResourceFile(text, apiAddr, resources.Delete, currentNamespace)
			}
			println("[f] waiting for deletion to complete")
			time.Sleep(time.Second * 3)
		}

		for _, text := range contents {
			err = resources.ParseAndActOnResourceFile(text, apiAddr, resources.Apply, currentNamespace)
			cobra.CheckErr(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(applyCmd)

	applyCmd.Flags().BoolVarP(&applyForce, "force", "f", false, "Force apply")
}
