# Healthy Service

An example service that can be used for testing readiness and liveness probe settings on Kubernetes.
It randomly returns json healthcheck output between `MAXTIME` and `MINTIME`.

## Usage:

➜  go build -v
github.com/rflorenc/randomhealth

➜  echo $MINTIME $MAXTIME
2 8

➜ ./randomhealth

## Kubernetes deployment
➜ kubectl apply -f examples/deployment.yaml


