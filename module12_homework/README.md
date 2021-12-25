# Homework 12: Deploy httpserver with istio
[v] deploy httpserver with istio\ 
[]
[]

## install istioctl
https://istio.io/latest/docs/ops/diagnostic-tools/istioctl/#install-hahahugoshortcode-s2-hbhb

## minikube 
in new terminal, execute `minikube tunnel`
Now an external ip is assigned to the ingress gateway 

### deploy with https
`kubectl create -n istio-system secret tls cncamp-credential --key=cncamp.io.key --cert=cncamp.io.crt`
`curl --resolve cncamp-service.cncamp.io:443:$INGRESS_IP https://cncamp-service.cncamp.io/ -v -k`

### L7 routing
`curl --resolve cncamp-service.cncamp.io:443:$INGRESS_IP https://cncamp-service.cncamp.io/module12 -v -k`

`curl --resolve cncamp-service.cncamp.io:443:$INGRESS_IP https://cncamp-service.cncamp.io/nginx -v -k`

--- under construction ---
## How to start 
### start kind cluster if need
`make kind-up`
### install istio with istioctl
`istioctl install`
...
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


