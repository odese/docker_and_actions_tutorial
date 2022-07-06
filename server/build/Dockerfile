# syntax=docker/dockerfile:1

FROM golang:1.18.3-buster
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o my-server /app/server/cmd/main.go
EXPOSE 8080
CMD [ "my-server" ]