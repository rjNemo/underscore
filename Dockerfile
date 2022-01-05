FROM golang:1.18beta1-bullseye

WORKDIR /lib

COPY go.* ./
RUN go mod download

COPY . ./
