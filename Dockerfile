FROM golang:1.18-alpine

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64

RUN apk upgrade --no-cache

WORKDIR /lib

COPY go.* ./
RUN go mod download

COPY . ./
