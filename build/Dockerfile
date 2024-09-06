# Stage 1: Build the Go app
FROM golang:1.22-alpine AS build

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum from the root of the project (one directory above)
COPY ../go.mod ../go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code from the root directory (one level up)
COPY ../ .

# Build the Go app
RUN go build -o /dalabio ./cmd/server

# Stage 2: Run the Go app
FROM alpine:latest

# Install necessary libraries for running Go and PostgreSQL
RUN apk --no-cache add ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the binary from the build stage
COPY --from=build /dalabio .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./dalabio"]
