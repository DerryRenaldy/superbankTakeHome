FROM golang:latest AS build

WORKDIR /app

# Copy the Go module files
COPY go.mod .
COPY go.sum .

# Download the Go module dependencies
RUN go mod download

# Copy the source code (main.go is in the root directory)
COPY . .

# Build the Go application to a relative path
RUN go build -o myapp .

# Use a smaller image for running the app
FROM alpine:latest AS run

# Install any necessary packages
RUN apk add --no-cache libc6-compat

# Copy the compiled binary from the build stage
COPY --from=build /app/myapp /myapp

# Copy the configuration file into the image
COPY --from=build /app/configs/config.yaml /app/configs/config.yaml

WORKDIR /app
EXPOSE 8090

# Ensure the binary is executable
RUN chmod +x /myapp

# Run the application
CMD ["/myapp"]
