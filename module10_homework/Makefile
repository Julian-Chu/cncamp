
KIND_CLUSTER := kind
test-healthz:
	@curl -XGET -I http://localhost:8000/healthz

test-response:
	@curl -XGET -I  http://localhost:8000 -H "MyHeader: test"

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

kind-up:
	kind create cluster \
			--image kindest/node:v1.21.1\
			--name $(KIND_CLUSTER) \
			--config kind-config.yaml

kind-down:
	kind delete cluster --name $(KIND_CLUSTER)

kind-load:
	kind load docker-image julianchu/cncamp-go-server:metrics

