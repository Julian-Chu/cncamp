# Homework 12: Deploy httpserver with istio
[v] deploy httpserver with istio \
[v] L7 routing rules \
[ ] tracing

## How to start
### install istioctl
https://istio.io/latest/docs/ops/diagnostic-tools/istioctl/#install-hahahugoshortcode-s2-hbhb

### minikube  tunnel
in new terminal, execute `minikube tunnel` \
Now an external ip is assigned to the ingress gateway 

### get ingress ip 
`kubectl get svc -n istio-system`\
export INGRESS_IP = \<external ip\>

### create tls  in secret
`kubectl create -n istio-system secret tls cncamp-credential --key=cncamp.io.key --cert=cncamp.io.crt`

### apply yaml files
`make apply`
### request with https
`curl --resolve cncamp-service.cncamp.io:443:$INGRESS_IP https://cncamp-service.cncamp.io/ -v -k`

### request with L7 routing
`curl --resolve cncamp-service.cncamp.io:443:$INGRESS_IP https://cncamp-service.cncamp.io/module12 -v -k`

`curl --resolve cncamp-service.cncamp.io:443:$INGRESS_IP https://cncamp-service.cncamp.io/nginx -v -k`


