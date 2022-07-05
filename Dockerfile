# syntax=docker/dockerfile:1

FROM golang:1.18.3-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /docker-gs-ping

CMD [ "/docker-gs-ping" ]