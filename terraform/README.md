# Terraform
Terraform is an open-source infrastructure as code software tool that provides a consistent CLI workflow to manage hundreds of cloud services. Terraform codifies cloud APIs into declarative configuration files.

## Prerequisites

To install helm;
```
$ brew install terraform
```

To install minikube;
```
$ brew install minikube
```

## Minikube

To start minikube;
```
$ minikube start
```

To see cluster info;
```
$ kubectl cluster-info
```

To see dashboard;
```
$ minikube dashboard
```

## Terraform

To download the specified version of K8S provider from tf file;
```
$ terraform init
```

To display a list of resources to be created, and highlight any possible unknown attributes at apply time;
```
$ terraform plan
```

To deploy your application;
```
$ terraform apply --auto-approve
```

To access the application
```
$ kubectl --namespace terraform port-forward <pod-name> 8080:80
```

To delete the application
```
$ terraform destroy
```