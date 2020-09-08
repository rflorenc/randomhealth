# Random Health Service

An example service that can be used for testing readiness and liveness probe settings on Kubernetes.

It randomly sleeps between `MAXTIME` and `MINTIME` and returns basic json healthcheck output.

## Build

>$ go build

## Usage

>$  export MINTIME=2; export MAXTIME=8

>$ ./randomhealth

>$ curl 127.0.0.1:9337/health
```
/health endpoint called
Sleeping 4 seconds...
{"healthMax:":7,"healthMin:":3,"healthy":true,"n":4,"remoteAddr":"127.0.0.1:59057"}%
```


>$ curl 127.0.0.1:9337/health
```
/health endpoint called
Sleeping 5 seconds...
{"healthMax:":7,"healthMin:":3,"healthy":true,"n":5,"remoteAddr":"127.0.0.1:59064"}%
```


## Kubernetes deployment

>$ kubectl apply -f examples/deployment.yaml

The time range environment variables can be modified in the controller spec.
```
    env:
    - name: MINTIME
      value: "2"
    - name: MAXTIME
      value: "6"
```

## Docker

>$ docker build -t rfplay/randomhealth .

>$ docker run -e MINTIME=3 -e MAXTIME=7 -it rfplay/randomhealth:latest
```
INFO:main.go:73: MAXTIME: 7 MINTIME: 3
```
