# Homework 8: Deploy httpserver
## How to start 
### start kind cluster if need
`make kind-up`
### Install Ingress NGINX into k8s cluster
https://kind.sigs.k8s.io/docs/user/ingress/#using-ingress
```shell
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml

kubectl wait --namespace ingress-nginx \
  --for=condition=ready pod \
  --selector=app.kubernetes.io/component=controller \
  --timeout=90s
```
### deploy and test
```shell
make apply                 // deploy app
curl -H "Host: cncamp.com" https://127.0.0.1 -v -k // version header should be loaded from configmap
curl -H "Host: cncamp.com" https://127.0.0.1/healthz -v -k  // health check
make delete                // remove app
```

### remove kind cluster 
`make kind-down`


## k8s yaml
under `/k8s`
- http-server:
  - deployment.yaml
  - config.yaml
  - service.yaml
  - ingress.yaml
- dummy pod for debug
  - debug-pod.yaml


## kind with Ingress nginx
kind config: `kind-config`
## Create kind cluster
```shell
make kind-up
```

### hwo to create TLS in secret
kubectl create secret tls cncamp-tls --cert=./tls.crt --key=./tls.key


