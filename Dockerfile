FROM golang:1.18beta1-alpine

WORKDIR /lib

COPY go.* ./
RUN go mod download

COPY . ./
