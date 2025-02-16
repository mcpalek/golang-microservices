FROM golang:1.23-alpine AS build

# Install curl
RUN apk --no-cache add curl

# Set the working directory inside the container
WORKDIR /app

# Copy Go module files from the root of the project
COPY frontend_service/go.mod ./

# Download Go dependencies
RUN go mod download

# Copy the rest of the source code
COPY frontend_service ./

# Ensure templates directory is copied
COPY frontend_service/templates templates/

# Build the Go application
RUN go build -o frontend-service .

# Use a minimal Alpine Linux image for the final runtime container
FROM alpine:latest

# Install necessary dependencies (e.g., libc)
RUN apk --no-cache add libc6-compat

# Set the working directory inside the container
WORKDIR /app

# Copy the built binary from the build stage
COPY --from=build /app/frontend-service .

# Copy the templates directory from the build stage
COPY --from=build /app/templates templates/
RUN ls -lah /app/templates
# Set environment variable (adjust if needed)
ENV API_URL=http://web_service:8081/users

# Expose the port the app will run on
EXPOSE 8082

# Command to run the application
CMD ["./frontend-service"]
