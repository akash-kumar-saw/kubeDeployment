/*
Copyright © 2023 Akash Kumar Saw <akashkumarsaw03@gmail.com>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/ghodss/yaml"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "KubeDeployment - Update",
	Long: `To update an existing deployment, use the update subcommand and provide the path to your updated deployment.yaml file using the "deployment" flag and the namespace using the "namespace" flag.
	
	Example : kubeDeployment update --deployment=<path-to-new-deployment.yaml> --namespace=<namespace>
	`,
	Run: func(cmd *cobra.Command, args []string) {

		configname, err := cmd.Flags().GetString("configname")
		if err != nil {
			log.Fatal(err)
		}

		kubeconfig := getKubeConfig(configname)

		deployment, err := cmd.Flags().GetString("deployment")
		if err != nil {
			log.Fatal(err)
		}
		if deployment == "" {
			log.Fatal("Please provide the deployment.yaml file")
		}

		namespace, err := cmd.Flags().GetString("namespace")
		if err != nil {
			log.Fatal(err)
		}

		config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			log.Fatal(err)
		}

		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			log.Fatal(err)
		}

		updateDeployment(clientset, deployment, namespace)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.PersistentFlags().String("configname", getdefaultConfig(), "Name for kubeconfig file")
	updateCmd.PersistentFlags().String("deployment", "", "Path to the YAML file for the deployment")
	updateCmd.PersistentFlags().String("namespace", "default", "Namespace for the deployment")
}

func updateDeployment(clientset *kubernetes.Clientset, deploymentFile string, namespace string) {
	deployment := &appsv1.Deployment{}

	yamlData, err := os.ReadFile(deploymentFile)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal([]byte(yamlData), &deployment)
	if err != nil {
		log.Fatal(err)
	}

	_, err = clientset.AppsV1().Deployments(namespace).Update(context.TODO(), deployment, metav1.UpdateOptions{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deployment %s updated successfully\n", deployment.Name)
}
