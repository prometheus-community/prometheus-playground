FROM golang:1.18.3-alpine3.16

RUN apk update \
    && apk upgrade \
    && apk add --no-cache git \
    && go install github.com/prometheus/alertmanager/cmd/amtool@v0.24.0 \
    && apk del git
