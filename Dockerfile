# Use an official Go runtime as the base image
FROM golang:1.18 AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go mod and sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the application
RUN go build -o app

# Use a minimal base image to run the application
FROM debian:buster-slim

# Set the working directory inside the container
WORKDIR /app

# Copy the built binary from the previous stage
COPY --from=build /app/app .

# Expose the port that the application listens on
EXPOSE 8080

# Command to run the application
CMD ["./app"]