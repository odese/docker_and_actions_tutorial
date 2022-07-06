FROM golang:1.18.3-alpine3.16 AS build-env

RUN mkdir -p /go/src/wow
WORKDIR /go/src/wow

COPY . .

# WORKDIR /go/src/wow/client
RUN ls

# WORKDIR /go/src
RUN ls

RUN go mod init
RUN go mod tidy

RUN ls

RUN go build -o clientApp ./cmd/main.go

# COPY --from=build-env /go/client  /app/client

CMD ["./client"]