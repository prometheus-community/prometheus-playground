FROM golang:1.11beta1-alpine3.7

RUN apk update \
    && apk upgrade \
    && apk add --no-cache git \
    && go get github.com/prometheus/alertmanager/cmd/amtool \
    && apk del git