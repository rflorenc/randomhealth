# Healthy Service

An example service that can be used for testing readiness and liveness probe settings on Kubernetes.
It randomly returns json healthcheck output between `MAXTIME` and `MINTIME`.

## Usage

➜  go build -v

➜  export MINTIME=2; export MAXTIME=8

➜ ./randomhealth


## Kubernetes deployment
➜ kubectl apply -f examples/deployment.yaml

The time range environment variables can be modified in the deployment.

    env:
    - name: MINTIME
      value: "2"
    - name: MAXTIME
      value: "6"


