/*
Copyright © Akash Kumar Saw <akashkumarsaw03@gmail.com>
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

// readCmd represents the read command
var readCmd = &cobra.Command{
	Use:   "read",
	Short: "KubeDeployment - Read",
	Long: `To retrieve information about deployments, use the read subcommand and provide the namespace using the "namespace" flag.
	
	Example : kubeDeployment read --namespace=<namespace>
	`,
	Run: func(cmd *cobra.Command, args []string) {
		kubeconfig := getKubeConfig()

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

		readDeployment(clientset, namespace)
	},
}

func init() {
	rootCmd.AddCommand(readCmd)

	readCmd.PersistentFlags().String("namespace", "default", "Namespace for the deployment")
}

func readDeployment(clientset *kubernetes.Clientset, namespace string) {
	deployments, err := clientset.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, deployment := range deployments.Items {
		fmt.Printf("Deployment: %s\n", deployment.Name)
		fmt.Printf("Replicas: %d\n", *deployment.Spec.Replicas)
		fmt.Printf("Image: %s\n", deployment.Spec.Template.Spec.Containers[0].Image)
	}
}
