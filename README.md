# Kubernetes Deployment Manager
KubeDeployment is a command-line tool for managing Kubernetes deployments using the Kubernetes API. It supports the creation, retrieval, updating, and deletion of deployments in a Kubernetes cluster.

## Installation
- Clone the repository from GitHub:
```bash
git clone https://github.com/akash-kumar-saw/kubeDeployment
```
- Install the required dependencies:
```bash
cd kubeDeployment
go mod download
```
- Run the following command to build the tool:
```bash
go build .
```
- Run the following command to start the tool:
```bash
kubeDeployment
```

## Usage
To use this tool, you must have a kubeconfig file that allows access to your Kubernetes cluster. You can pass the path to your kubeconfig file using the -kubeconfig flag when using the apply subcommand.

```bash
kubeDeployment apply --kubeconfig=<path-to-kubeconfig>
```

## Subcommands
This tool uses subcommands to specify the action to perform. The available subcommands are:

- apply - applies a configuration to a Kubernetes cluster
- create - creates a new deployment
- read - retrieves information about deployments
- update - updates an existing deployment
- delete - deletes an existing deployment

### Apply
To apply a configuration to a Kubernetes cluster, use the apply subcommand and provide the path to your kubeconfig.yaml file using the "kubeconfig" flag.

```bash
kubeDeployment apply --kubeconfig=<path-to-kubeconfig.yaml>
```

### Create
To create a new deployment, use the create subcommand and provide the path to your deployment.yaml file using the "deployment" flag and the namespace using the "namespace" flag.

```bash
kubeDeployment create --deployment=<path-to-deployment.yaml> --namespace=<namespace>
```

### Read
To retrieve information about deployments, use the read subcommand and provide the namespace using the "namespace" flag.

```bash
kubeDeployment read --namespace=<namespace>
```

### Update
To update an existing deployment, use the update subcommand and provide the path to your updated deployment.yaml file using the "deployment" flag and the namespace using the "namespace" flag.

```bash
kubeDeployment update --deployment=<path-to-new-deployment.yaml> --namespace=<namespace>
```

### Delete
To delete an existing deployment, use the delete subcommand and provide the name of the deployment you want to delete using "deployment" flag and the namespace using the "namespace" flag.

```bash
kubeDeployment delete --deployment=<name-of-deployment> --namespace=<namespace>
```

## Dependencies
This tool uses the following dependencies:

k8s.io/client-go - Kubernetes client library for Go
gopkg.in/yaml.v2 - YAML parser for Go
