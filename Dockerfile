# syntax=docker/dockerfile:1

FROM golang:1.23 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN make

FROM alpine:latest
WORKDIR /app
COPY --from=builder ./bin/main .
CMD ["./main"]
