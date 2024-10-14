FROM golang:1.23.2 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 make


FROM alpine:3 AS release
WORKDIR /
COPY --from=builder /app/bin/main /main
RUN chmod +x /main

ENTRYPOINT [ "./main" ]
EXPOSE 8080