/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// namespaceCmd represents the namespace command
var namespaceCmd = &cobra.Command{
	Use:   "namespace",
	Short: "KubeDeployment - Namespace",
	Long: `To play with namespaces, use the namespace subcommand and provide the name for the namespace using the "name" flag and the action to perform using the "action" flag.
	
	Example : kubeDeployment namespace --action=<create,delete> --name=<name for the namespace>
	`,
	Run: func(cmd *cobra.Command, args []string) {
		action, err := cmd.Flags().GetString("action")
		if err != nil {
			log.Fatal(err)
		}

		name, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Fatal(err)
		}

		configname, err := cmd.Flags().GetString("configname")
		if err != nil {
			log.Fatal(err)
		}

		kubeconfig := getKubeConfig(configname)

		config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			log.Fatal(err)
		}

		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			log.Fatal(err)
		}

		if action == "" && name == "" {
			displayNamespace(clientset)
		} else if action == "" {
			createNamespace(clientset, name)
		} else if name == "" {
			log.Fatal("Please provide a name for your namespace action using 'name' option")
		} else {
			switch action {
			case "create":
				createNamespace(clientset, name)
			case "delete":
				deleteNamespace(clientset, name)
			default:
				log.Fatal("Invalid action provided. Please provide 'create' or 'delete' option")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(namespaceCmd)

	namespaceCmd.PersistentFlags().String("action", "", "Action to perform with the namespace")
	namespaceCmd.PersistentFlags().String("name", "", "Name for the namespace")
	namespaceCmd.PersistentFlags().String("configname", getdefaultConfig(), "Name of the kubeconfig file")
}

func createNamespace(clientset *kubernetes.Clientset, name string) {
	newNamespace := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}

	_, err := clientset.CoreV1().Namespaces().Create(context.Background(), newNamespace, metav1.CreateOptions{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Namespace %s created successfully.\n", name)
}

func deleteNamespace(clientset *kubernetes.Clientset, name string) {

	err := clientset.CoreV1().Namespaces().Delete(context.Background(), name, metav1.DeleteOptions{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Namespace %s deleted successfully.\n", name)
}

func displayNamespace(clientset *kubernetes.Clientset) {

	namespaces, err := clientset.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("List of Namespaces:")
	for _, ns := range namespaces.Items {
		fmt.Println(ns.ObjectMeta.Name)
	}
}
