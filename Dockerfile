# syntax=docker/dockerfile:1

## Build
FROM golang:1.20-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY main.go ./
COPY ./pkg/. ./pkg/
RUN GOOS=linux GOARCH=amd64 go build -o /wsperf

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /wsperf /wsperf

EXPOSE 8080

USER nonroot:nonroot

CMD ["/wsperf"]
