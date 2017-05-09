FROM golang:latest

WORKDIR /usr/local/src
COPY journald-proxy.go .

RUN go get -d
RUN go build -o journald-proxy

