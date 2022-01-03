FROM golang:1.18beta1-bullseye

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN go test ./... -cover