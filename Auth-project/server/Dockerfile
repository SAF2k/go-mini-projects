# Use the official Golang image as the base image
FROM golang:1.16-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go binary inside the container
RUN go build -o main .

# Expose port 8000 for the application
EXPOSE 8000

# Start the application when the container starts
CMD ["./main"]