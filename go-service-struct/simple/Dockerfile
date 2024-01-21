# Dockerfile for go application with multi stage build
# Build stage
FROM golang:1.13.4-alpine3.10 AS build-env
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go build -o /go/bin/app

# Final stage
FROM alpine:3.10
WORKDIR /app
COPY --from=build-env /go/bin/app /app/
ENTRYPOINT ./app
