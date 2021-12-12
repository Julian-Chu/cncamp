
KIND_CLUSTER := kind
test-healthz:
	@curl -XGET -I http://localhost/healthz -v 

test-response:
	@curl -XGET -I  http://localhost/ -v

run:
	@export VERSION="1.1.1" && go run main.go

docker-build:
	@docker build -f Dockerfile -t julianchu/cncamp-go-server:metrics .

docker-push:
	@docker tag julianchu/cncamp-go-server:metrics 
	@docker image push julianchu/cncamp-go-server:metrics

docker-run:
	@docker run -p 8080:8000 julianchu/cncamp-go-server:metrics



apply:
	kubectl apply -f ./k8s

delete:
	kubectl delete -f ./k8s

debug:
	kubectl exec -it debug -- bash

ingress-install:
	kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml	

kind-up:
	kind create cluster \
			--image kindest/node:v1.21.1\
			--name $(KIND_CLUSTER) \
			--config kind-config.yaml

kind-down:
	kind delete cluster --name $(KIND_CLUSTER)

kind-load:
	kind load docker-image julianchu/cncamp-go-server:metrics


prometheus-install:
	helm install prom-stack -f prometheus-additional.yaml -f values.yaml prometheus-community/kube-prometheus-stack

prometheus-uninstall:
	helm uninstall prom-stack

prometheus-forward:
	kubectl port-forward service/prom-stack-kube-prometheus-prometheus 3001:9090

grafana-forward:
	kubectl port-forward service/prom-stack-grafana 3000:80


helm-upgrade:
	helm upgrade -f prometheus-additional.yaml  -f values.yaml  prom-stack prometheus-community/kube-prometheus-stack
