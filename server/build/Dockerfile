# syntax=docker/dockerfile:1

FROM golang:1.18.3-alpine

WORKDIR /app

# COPY go.mod ./
# COPY go.sum ./

# RUN go mod download

COPY . .

# RUN go mod init
RUN go mod tidy

RUN ls

RUN go build -o my-server /app/server/cmd/main.go

EXPOSE 8080

CMD [ "my-server" ]