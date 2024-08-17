# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.23.0-alpine AS build-stage
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /api ./cmd/goPatterns

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM alpine:3.20.2 AS build-release-stage
WORKDIR /
COPY --from=build-stage /api /api

EXPOSE 8080

ENTRYPOINT ["/api"]