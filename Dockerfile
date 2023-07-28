# Build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go build -o /go/bin/server -v ./cmd

# Final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /go/bin/server .
ENTRYPOINT ["./server"]

LABEL Name=atass-payments Version=0.0.1
EXPOSE 8080