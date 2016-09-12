FROM golang:1.7.1
MAINTAINER Atsushi Nagase<a@ngs.io>

VOLUME ["/var/run/docker.sock"]

RUN go get "github.com/docker/docker/api/types" && \
    go get "github.com/docker/docker/client" && \
    go get "github.com/nlopes/slack" && \
    go get "golang.org/x/net/context" && \
    go get "github.com/olekukonko/tablewriter"

ADD rabot.go /go/src/github.com/ngs/rabot/rabot.go
RUN go install github.com/ngs/rabot

ENTRYPOINT ["/go/bin/rabot"]

