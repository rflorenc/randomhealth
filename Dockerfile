FROM openshift/origin-release:golang-1.14

RUN GOBIN=/usr/bin go get github.com/rflorenc/randomhealth

CMD ["/usr/bin/randomhealth"]
