# Use a lightweight Go image to build the frontend service
FROM golang:1.20-alpine AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go source code into the container
COPY . .

# Download Go dependencies and build the Go app
RUN go mod tidy
RUN go build -o frontend-service .

# Use a smaller image for the final runtime container
FROM alpine:latest

# Install necessary dependencies (e.g., libc)
RUN apk --no-cache add libc6-compat

# Set the working directory inside the container
WORKDIR /root/

# Copy the built binary from the build stage
COPY --from=build /app/frontend-service .

# Expose the port the app will run on
EXPOSE 8082

# Command to run the app
CMD ["./frontend-service"]
