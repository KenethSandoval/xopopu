# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /src

COPY ./src/go.mod ./
# RUN go mod download

COPY src/*.go ./

RUN go build -o /src/xopopu src/cmd/main.go