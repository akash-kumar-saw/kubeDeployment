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

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "KubeDeployment - Create",
	Long: `To create a new deployment, use the create subcommand and provide the path to your deployment.yaml file using the "deployment" flag and the namespace using the "namespace" flag.
	
	Example : kubeDeployment create --deployment=<path-to-deployment.yaml> --namespace=<namespace>
	`,
	Run: func(cmd *cobra.Command, args []string) {
		kubeconfig := getKubeConfig()

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

		createDeployment(clientset, deployment, namespace)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.PersistentFlags().String("deployment", "", "Path to the YAML file for the deployment")
	createCmd.PersistentFlags().String("namespace", "default", "Namespace for the deployment")
}

func getKubeConfig() string {
	kubeconfig, err := os.ReadFile("kubeconfig.txt")
	if err != nil {
		log.Fatal(err)
	}

	return string(kubeconfig)
}

func createDeployment(clientset *kubernetes.Clientset, deploymentFile string, namespace string) {
	deployment := &appsv1.Deployment{}

	deploymentData, err := os.ReadFile(deploymentFile)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal([]byte(deploymentData), &deployment)
	if err != nil {
		log.Fatal(err)
	}

	_, err = clientset.AppsV1().Deployments(namespace).Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deployment %s created successfully\n", deployment.Name)
}
