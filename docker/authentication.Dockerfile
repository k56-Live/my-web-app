# Use a lightweight base image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . /app

# Build the authentication service
RUN go build -o authentication-service ./authentication

# Expose the port on which the service will run
EXPOSE 8080

# Start the authentication service
CMD ["./authentication-service"]

