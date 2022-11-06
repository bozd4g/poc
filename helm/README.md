# Helm
Helm helps you manage Kubernetes applications — Helm Charts help you define, install, and upgrade even the most complex Kubernetes application.

## Prerequisites

To install helm;
```
$ brew install helm
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

## Helm

To create a new chart;
```
$ helm create helloworld
```

To deploy the application;
```
$ helm install helloworld ./helloworld
```

To access the application
```
$ export POD_NAME=$(kubectl get pods --namespace default -l "app.kubernetes.io/name=helloworld,app.kubernetes.io/instance=helloworld" -o jsonpath="{.items[0].metadata.name}")

$ export CONTAINER_PORT=$(kubectl get pod --namespace default $POD_NAME -o jsonpath="{.spec.containers[0].ports[0].containerPort}")

$ kubectl --namespace default port-forward $POD_NAME 8080:$CONTAINER_PORT
```


To update the application
```
$ helm upgrade helloworld ./helloworld
```

To delete the application
```
$ helm delete helloworld
```
