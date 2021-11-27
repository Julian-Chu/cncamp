
KIND_CLUSTER := kind

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
	kind load docker-image julianchu/cncamp-go-server:dev

