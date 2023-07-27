# My Microservices

This repository contains the source code for the authentication service and web app service in a microservices architecture.

### Authentication Service

The authentication service handles user authentication and authorization. It provides endpoints for user login, registration, and token generation.

### Web App Service

The web app service is a front-end application that interacts with the authentication service to authenticate users and display content.

## Getting Started

To run the microservices, you can use Docker to build and deploy the containers.

### Build and Run Authentication Service

docker build -t authentication-service -f docker/authentication.Dockerfile .
docker run -p 8080:8080 authentication-service

### Build and Run Web App Service
docker build -t webapp-service -f docker/webapp.Dockerfile .
docker run -p 8080:8080 webapp-service

You can access the web app service at http://localhost:8080.

With this example code, you have a basic microservices architecture with two services: the authentication service and the web app service. The .github/workflows/go.yml file sets up a GitHub Actions workflow for building and testing the services on each push to the main branch. The docker directory contains Dockerfiles to build Docker images for each service, and the go.mod and go.sum files manage the Go dependencies. The README.md provides an overview of the microservices and instructions for building and running them using Docker.
