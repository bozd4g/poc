# ArgoCD
Argo CD is a declarative, GitOps continuous delivery tool for Kubernetes.

It follows the GitOps pattern of using Git repositories as the source of truth for defining the desired application state.

It automates the deployment of the desired application states in the specified target environments. Application deployments can track updates to branches, tags, or be pinned to a specific version of manifests at a Git commit.

## Prerequisites


To install minikube;
```
$ brew install minikube
```

To install kubectl;
```
$ brew install kubectl
```

To download argocd;
```
https://github.com/argoproj/argo-cd/releases/download/v2.0.0/argocd-darwin-amd64
```

To install kustomize;
```
$ brew install kustomize
```

## Minikube
To start minikube;
```
$ minikube start --memory=4096 --cpus=2 --kubernetes-version=v1.20.2 -p gitops
```

To see cluster info;
```
$ kubectl cluster-info
```

To see dashboard;
```
$ minikube dashboard
```

To enable ingress for minikube;
```
$ minikube addons enable ingress -p gitops
```

To create argocd namespace and install components;
```
$ kubectl create namespace argocd
$ kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
```

To open tunnel for minikube;
```
$ minikube tunnel
```

To patch the ArgoCD service from ClusterIP to a LoadBalancer;
```
$ kubectl patch svc argocd-server -n argocd -p '{"spec": {"type": "LoadBalancer"}}'
```


## Argo CD

To see argo cd ui, open `127.0.0.1` address on your browser.
Default username is `admin`.

To get the password;
```
$ argoPass=$(kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d)
echo $argoPass
```

To create an app;
```
$ kubectl apply -f argocd/openshift-gitops-examples/components/applications/bgd-app.yaml
```