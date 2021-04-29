
## About this project

This project is used to simulate sidecar container and app container starting sequence in Kubernetes.

### Description

1. `sidecar` will start a http server listening on port `8080`, and hold `10s` before status updated to `success` by responding `200` via `/ready` endpoint.

2. `entrypoint wait` will trigger `http://localhost:8080/ready`

### How to build

Running command `make build` will output an image as sidecar image.

### How to deploy(simulating)

Run command `kubectl apply -f deploy.yaml`. Before this, you can use `stern` to monitor pod log.

More details can be located [here]().