# syntax=docker/dockerfile:1

# Build Stage
FROM golang:1.22.5-alpine AS builder
WORKDIR /app

# Copy go.mod files
COPY go.mod ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o docker-gs-ping

# Final Stage
FROM alpine:latest
WORKDIR /root/

# Install curl for readiness probe
RUN apk add --no-cache curl

# Copy the executable from the builder stage
COPY --from=builder /app/docker-gs-ping .

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["./docker-gs-ping"]
