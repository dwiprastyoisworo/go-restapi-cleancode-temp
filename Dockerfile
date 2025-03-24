FROM golang:1.23-alpine AS build

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Copy the user.config file
COPY file file/

# Build the application
RUN go build -o /bin/serve cmd/app/*.go

# Final stage
FROM alpine:latest

WORKDIR /root/

# Copy the built binary from the build stage
COPY --from=build /bin/serve .

# Copy the user.config file
COPY --from=build /app/file file/

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./serve"]