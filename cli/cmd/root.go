package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile          string
	apiAddr          string
	currentNamespace string
)

var rootCmd = &cobra.Command{
	Use:   "prism",
	Short: "Manage your prism cloud resources from the command line.",
	Long: `Prism is a cloud platform that allows you to deploy and manage your applications. 
This CLI allows you to interact with the Prism API to manage your resources from the command line. 
You can create, update, delete, and list resources such as applications, services, and deployments.
	
Try running 'prism version' to test the CLI.
	`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.prism.yaml)")
	rootCmd.PersistentFlags().StringVar(&apiAddr, "api-addr", "localhost:18948", "Address of the Prism API server")
	rootCmd.PersistentFlags().StringVar(&currentNamespace, "namespace", "", "Namespace to use")

	err := viper.BindPFlag("api-addr", rootCmd.PersistentFlags().Lookup("api-addr"))
	if err != nil {
		panic(err)
	}

	viper.SetDefault("api-addr", "localhost:18948")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".prism")
	}

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		_, err := fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		if err != nil {
			return
		}
	}
}
