FROM golang:alpine AS builder

WORKDIR /app
COPY . /app

RUN go mod download && go build -o /swag-server

ENTRYPOINT ["/swag-server"]

