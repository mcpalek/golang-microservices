# Build stage
FROM golang:1.23-alpine AS build

WORKDIR /app

# Copy go.mod and go.sum from the correct directory
COPY db_service/go.mod db_service/go.sum ./

# Copy the configloader directory into the build context (shared module)
COPY configloader /app/configloader

# Download the Go module dependencies
RUN go mod download

# Copy the rest of the application code
COPY db_service/ ./

# Build the db-service binary
RUN go build -o db-service .

# Final stage
FROM alpine:latest

# Install libc6 compatibility
RUN apk --no-cache add libc6-compat

WORKDIR /root/

# Copy the built binary from the build stage
COPY --from=build /app/db-service .

EXPOSE 8000

CMD ["./db-service"]





