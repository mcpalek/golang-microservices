# # Build stage
FROM golang:1.23-alpine AS build

# Set working directory inside container
WORKDIR /app

# Copy go.mod and go.sum to handle dependencies
COPY web_service/go.mod web_service/go.sum ./

# Copy the configloader module as a shared dependency
COPY configloader /app/configloader

# Download dependencies after copying configloader
RUN go mod download

# Copy the application code (web_service code)
COPY web_service /app/web_service

# Build the web-service binary
RUN go build -o web-service ./web_service/main.go


# Final image stage
FROM alpine:latest

# Install dependencies for downloading and verifying Microsoft ODBC and tools
RUN apk add --no-cache \
    curl \
    gnupg \
    bash \
    libc6-compat \
    busybox-extras

# Define the architecture for package downloads (assuming amd64)
ARG architecture="amd64"

# Download Microsoft ODBC and SQL Tools
RUN curl -O https://download.microsoft.com/download/7/6/d/76de322a-d860-4894-9945-f0cc5d6a45f8/msodbcsql18_18.4.1.1-1_$architecture.apk && \
    curl -O https://download.microsoft.com/download/7/6/d/76de322a-d860-4894-9945-f0cc5d6a45f8/mssql-tools18_18.4.1.1-1_$architecture.apk && \
    # Optional: Verify signature
    curl -O https://download.microsoft.com/download/7/6/d/76de322a-d860-4894-9945-f0cc5d6a45f8/msodbcsql18_18.4.1.1-1_$architecture.sig && \
    curl -O https://download.microsoft.com/download/7/6/d/76de322a-d860-4894-9945-f0cc5d6a45f8/mssql-tools18_18.4.1.1-1_$architecture.sig && \
    curl https://packages.microsoft.com/keys/microsoft.asc | gpg --import - && \
    gpg --verify msodbcsql18_18.4.1.1-1_$architecture.sig msodbcsql18_18.4.1.1-1_$architecture.apk && \
    gpg --verify mssql-tools18_18.4.1.1-1_$architecture.sig mssql-tools18_18.4.1.1-1_$architecture.apk && \
    # Install the package(s)
    apk add --no-cache --allow-untrusted msodbcsql18_18.4.1.1-1_$architecture.apk && \
    apk add --no-cache --allow-untrusted mssql-tools18_18.4.1.1-1_$architecture.apk

# Set working directory
WORKDIR /root/

# Copy the built binary from the build stage
COPY --from=build /app/web-service .

# Expose the port used by the web-service
EXPOSE 8081

ENV PATH="/opt/mssql-tools18/bin:$PATH"


# Run the web-service application
CMD ["./web-service"]

