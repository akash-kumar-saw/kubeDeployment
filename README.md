# Kubernetes Deployment Manager
KubeDeployment is a powerfull command-line tool for managing Kubernetes deployments in multiple Kubernetes Cluster at once. It supports the creation, retrieval, updating, and deletion of deployments in a Kubernetes cluster.

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
./kubeDeployment
```

## Usage
To use this tool, you must have at least one kubeconfig file that allows access to your Kubernetes cluster. You can pass the path to your kubeconfig file using the --kubeconfig flag and associate it with a name using the --configname flag while using the apply subcommand.

```bash
kubeDeployment apply --kubeconfig=<path-to-kubeconfig> --configname=<name-for-kubeconfig>
```

## Subcommands
This tool uses subcommands to specify the action to perform. The available subcommands are:

- apply - add a new kubeconfig file
- default - select a default kubeconfig
- create - creates a new deployment
- read - retrieves information about deployments
- update - updates an existing deployment
- delete - deletes an existing deployment
- namespace - manages a namespace

### Apply
To add a new kubeconfig file, use the apply subcommand and provide the path to your kubeconfig.yaml file using the "kubeconfig" flag and associate it with a name using the "configname" flag.

```bash
kubeDeployment apply --kubeconfig=<path-to-kubeconfig.yaml> --configname=<name-for-kubeconfig>
```

### Default
To select a default kubeconfig file, use the default subcommand and provide the name of the kubeconfig added via apply subcommand using the "configname" flag.

```bash
kubeDeployment default --configname=<name-of-kubeconfig>
```

### Create
To create a new deployment, use the create subcommand and provide the path to your deployment.yaml file using the "deployment" flag, the namespace using the "namespace" flag and the kubeconfig using "configname" flag.

```bash
kubeDeployment create --deployment=<path-to-deployment.yaml> --namespace=<namespace> --configname=<name-of-kubeconfig>
```

### Read
To retrieve information about deployments, use the read subcommand and provide the namespace using the "namespace" flag and the kubeconfig using "configname" flag.

```bash
kubeDeployment read --namespace=<namespace> --configname=<name-of-kubeconfig>
```

### Update
To update an existing deployment, use the update subcommand and provide the path to your updated deployment.yaml file using the "deployment" flag, the namespace using the "namespace" flag and the kubeconfig using "configname" flag.

```bash
kubeDeployment update --deployment=<path-to-new-deployment.yaml> --namespace=<namespace> --configname=<name-of-kubeconfig>
```

### Delete
To delete an existing deployment, use the delete subcommand and provide the name of the deployment you want to delete using "deployment" flag, the namespace using the "namespace" flag and the kubeconfig using "configname" flag.

```bash
kubeDeployment delete --deployment=<name-of-deployment> --namespace=<namespace> --configname=<name-of-kubeconfig>
```

### Namespace
To view the namespace present in the kubernetes cluster, use the namespace subcommand and provide the kubeconfig using "configname" flag.

```bash
kubeDeployment namespace --configname=<name-of-kubeconfig>
```

To create or delete a namespace in the kubernetes cluster, use the namespace subcommand and provide action (either 'create' or 'delete') using "action" flag, name of the namespace using "name" flag and the kubeconfig using "configname".

```bash
kubeDeployment namespace --action=<create or delete> --name=<name-of-namespace> --configname=<name-of-kubeconfig>
```

## Contributing

I welcome contributions and feedback! If you'd like to contribute to this portfolio or suggest improvements, please follow these steps:

1. Fork and clone the repository
2. Create a new branch: `git checkout -b YourGithubId/YourFeatureName`
3. Make your changes and commit them: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin YourGithubId/YourFeatureName`
5. Create a pull request
