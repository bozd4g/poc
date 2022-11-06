# Istio
Istio is a service mesh—a modernized service networking layer that provides a transparent and language-independent way to flexibly and easily automate application network functions. It is a popular solution for managing the different microservices that make up a cloud-native application. Istio service mesh also supports how those microservices communicate and share data with one another.

## Prerequisites

To install minikube;
```
$ brew install minikube
```

To install istio;
```
$ curl -L https://istio.io/downloadIstio | sh -
```

To configure the demonstration profile of istio;
```
$ cd ~/istio-xxxx (version)

$ istioctl install --set profile=demo -y
```

## Minikube

To start minikube;
```
$ minikube start --memory=12000 --cpus=4
```

To see cluster info;
```
$ kubectl cluster-info
```

To see dashboard;
```
$ minikube dashboard
```

To create a new namespace;
```
$ kubectl create namespace <namespace> --dry-run=client -o yaml | kubectl apply -f -
```

## Istio

To enable istio on the namespace;
```
$ kubectl label namespace <namespace> istio-injection=enabled
```

To deploy an example applications;
```
$ kubectl apply --namespace=<namespace> -f bookinfo.yaml 
```

For external traffic;
```
$ kubectl apply --namespace=<namespace> -f bookinfo-gateway.yaml 
```

To ensure that there are not any issues with the configuration;
```
$ cd ~/istio-xxxx (version)

$ istioctl analyze
```

To see services;
```
$ kubectl get services --namespace=<namespace>
```

To access the service, create tunnel feature of minikube;
```
$ minikube service <service-name> --namespace=<namespace>
```