
KIND_CLUSTER := kind

install-tracing:
	kubectl apply -f jaeger.yaml
	kubectl edit configmap istio -n istio-system set tracing.sampling=100
apply:
	kubectl create ns module12
	kubectl label ns module12 istio-injection=enabled
	kubectl apply -f ./k8s -n module12

delete:
	kubectl delete -f ./k8s -n module12
	kubectl delete ns module12

debug:
	kubectl exec -it debug -n module12 -- bash

kind-up:
	kind create cluster \
			--image kindest/node:v1.21.1\
			--name $(KIND_CLUSTER) \
			--config kind-config.yaml

kind-down:
	kind delete cluster --name $(KIND_CLUSTER)

kind-load:
	kind load docker-image julianchu/cncamp-go-server:dev

