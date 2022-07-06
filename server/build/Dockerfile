# syntax=docker/dockerfile:1

FROM golang:1.18.3-alpine

WORKDIR /app

# COPY go.mod ./
# COPY go.sum ./

# RUN go mod download

COPY server .

RUN ls

# WORKDIR /app/server
WORKDIR /server

RUN go mod init `server`
RUN go mod tidy

RUN ls

RUN go build -o my-server /app/server/cmd/main.go

EXPOSE 8080

CMD [ "my-server" ]