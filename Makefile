
build:
	docker build -t addozhang/k8s-container-sequence-sidecar:latest . -f docker/Dockerfile.sidecar

load-2-minikube:
	minikube image load addozhang/k8s-container-sequence-sidecar:latest