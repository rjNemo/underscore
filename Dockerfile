FROM golang:1.18-alpine

WORKDIR /lib

COPY go.* ./
RUN go mod download

COPY . ./
