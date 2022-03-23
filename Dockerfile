FROM golang:1.18-alpine

ENV CGO_ENABLED 0
ENV GOOS linux

RUN apk update --no-cache

WORKDIR /lib

COPY go.* ./
RUN go mod download

COPY . ./
