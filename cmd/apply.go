/*
Copyright © 2023 Akash Kumar Saw <akashkumarsaw03@gmail.com>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"k8s.io/client-go/util/homedir"
)

// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "KubeDeployment - Apply",
	Long: `To apply a configuration to a Kubernetes cluster, use the apply subcommand and provide the path to your kubeconfig.yaml file using the "kubeconfig" flag.
	
	Example : kubeDeployment apply --kubeconfig=<path-to-kubeconfig.yaml>
	`,
	Run: func(cmd *cobra.Command, args []string) {
		kubeconfig, err := cmd.Flags().GetString("kubeconfig")
		if err != nil {
			log.Fatal(err)
		}

		err = os.WriteFile("kubeconfig.txt", []byte(kubeconfig), 0644)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Kubeconfig applied successfully\n")
	},
}

func init() {
	rootCmd.AddCommand(applyCmd)

	applyCmd.PersistentFlags().String("kubeconfig", getDefaultKubeconfigPath(), "Path to the kubeconfig file")
}

func getDefaultKubeconfigPath() string {
	home := homedir.HomeDir()
	return fmt.Sprintf("%s/.kube/config", home)
}
