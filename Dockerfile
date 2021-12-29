FROM golang:1.18beta1-bullseye

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY *.go ./

RUN go test ./... -cover