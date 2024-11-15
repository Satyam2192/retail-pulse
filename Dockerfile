# Start from golang base image
FROM golang:1.21-alpine3.18

# Add git for fetching dependencies
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o main ./cmd/server

# Expose port 7000
EXPOSE 7000

# Command to run the application
CMD ["./main"]