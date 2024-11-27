FROM golang:1.23.2 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN make build

FROM alpine:3 AS release
WORKDIR /
COPY --from=builder /app/bin/main /main

COPY .env /.env

ENTRYPOINT [ "./main" ]
EXPOSE 8080