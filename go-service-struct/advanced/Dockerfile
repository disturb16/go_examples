# Dockerfile for go application with multi stage build
# Build stage
FROM golang:1.21.6-alpine3.18 AS build-env
ARG MAIN_FOLDER=api
RUN apk add --no-cache git
WORKDIR /app


# Copy external project dependencies,
# path is relative to the Dockerfile
COPY ["../msgbroker/go.mod", "../msgbroker/msgbroker.go", "/msgbroker/"]

# Copy the entire project and download dependencies
COPY ./advanced .
RUN go mod download -x

# Build the specified project
RUN cd ./$MAIN_FOLDER && go build -o /go/bin/app

# Final stage
FROM alpine:3.18
WORKDIR /app
COPY --from=build-env /go/bin/app /app/
ENTRYPOINT ./app
