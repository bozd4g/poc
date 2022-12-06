# Consul
Consul is a service networking solution that enables teams to manage secure network connectivity between services and across multi-cloud environments and runtimes.

## Prerequisites


To install minikube;
```
$ brew install minikube
```

To install kubectl;
```
$ brew install kubectl
```

## Minikube
To start minikube;
```
minikube start --profile dc1 --memory 4096 --kubernetes-version=v1.22.0
```

To see cluster info;
```
$ kubectl cluster-info
```

To see dashboard;
```
$ minikube dashboard
```

To open tunnel for minikube;
```
$ minikube tunnel
```


## Consul

To install consul extensions;
```sh
$ brew tap hashicorp/tap

$ brew install hashicorp/tap/consul-k8s
```

To install consul on minikube;
```sh
$ consul-k8s install -config-file=values.yaml -set global.image=hashicorp/consul:1.14.0
```

To set environment variables;
```sh
$ export CONSUL_HTTP_TOKEN=$(kubectl get --namespace consul secrets/consul-bootstrap-acl-token --template={{.data.token}} | base64 -d)

$ export CONSUL_HTTP_ADDR=https://127.0.0.1:8501

$ export CONSUL_HTTP_SSL_VERIFY=false
```

To create port forward;
```sh
$ kubectl port-forward svc/consul-ui --namespace consul 8501:443
```

To create servicess;
```sh
$ kubectl apply -f services/counting.yaml && kubectl apply -f services/dashboard.yaml
```

To create intentions for service mesh;
```sh
$ kubectl apply --filename intentions.yaml
```