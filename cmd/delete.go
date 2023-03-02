/*
Copyright © 2023 Akash Kumar Saw <akashkumarsaw03@gmail.com>
*/
package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "KubeDeployment - Delete",
	Long: `To delete an existing deployment, use the delete subcommand and provide the name of the deployment you want to delete using "deployment" flag and the namespace using the "namespace" flag.
	
	Example : kubeDeployment delete --deployment=<name-of-deployment> --namespace=<namespace>
	`,
	Run: func(cmd *cobra.Command, args []string) {
		kubeconfig := getKubeConfig()

		deployment, err := cmd.Flags().GetString("deployment")
		if err != nil {
			log.Fatal(err)
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

		deleteDeployment(clientset, deployment, namespace)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.PersistentFlags().String("deployment", "", "Name of the Deployment")
	deleteCmd.PersistentFlags().String("namespace", "default", "Namespace for the deployment")
}

func deleteDeployment(clientset *kubernetes.Clientset, deployment string, namespace string) {
	deletePolicy := metav1.DeletePropagationForeground

	err := clientset.AppsV1().Deployments(namespace).Delete(context.TODO(), deployment, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deployment %s deleted successfully\n", deployment)
}
