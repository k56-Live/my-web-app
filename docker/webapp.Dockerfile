# Use a lightweight base image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . /app

# Build the web app service
RUN go build -o webapp-service ./webapp

# Expose the port on which the service will run
EXPOSE 8080

# Start the web app service
CMD ["./webapp-service"]

