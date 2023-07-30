/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// defaultCmd represents the default command
var defaultCmd = &cobra.Command{
	Use:   "default",
	Short: "KubeDeployment - Default",
	Long: `To set a default kubeconfig file, use the default subcommand and provide the name of the kubeconfig file using the "configname" option
	
	Example : kubeDeployment namespace --action=<create,delete> --name=<name for the namespace>
	`,
	Run: func(cmd *cobra.Command, args []string) {
		configname, err := cmd.Flags().GetString("configname")
		if err != nil {
			log.Fatal(err)
		}

		if configname == "" {
			log.Fatal("Please provide a name for the kubeconfig file using 'configname' option")
		}

		err = os.WriteFile("./config/default.txt", []byte(strings.Trim(configname, " \t\n")), 0644)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Kubeconfig default applied successfully\n")
	},
}

func init() {
	rootCmd.AddCommand(defaultCmd)

	applyCmd.PersistentFlags().String("configname", "", "Name for the kubeconfig file")
}
