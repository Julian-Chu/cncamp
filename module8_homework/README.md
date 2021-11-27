# Homework 8: Deploy httpserver
## deploy and test
```shell
make apply                 // deploy app
curl localhost -v          // version header should be loaded from configmap
curl localhost/healthz -v  // health check
make delete                // remove app
```
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
make kind-down
```
### Ingress NGINX
https://kind.sigs.k8s.io/docs/user/ingress/#using-ingress
```shell
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml

kubectl wait --namespace ingress-nginx \
  --for=condition=ready pod \
  --selector=app.kubernetes.io/component=controller \
  --timeout=90s
```


