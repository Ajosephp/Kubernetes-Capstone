### README:

# Assignment 01 - Go Application with Docker

## Description

This is a simple Go application that supports three endpoints:

- `GET /foo`: Returns "bar".
- `POST /hello`: Accepts JSON `{"name": "YourName"}` and returns "Hello YourName!".
- `GET /kill`: Shuts down the server.

## Prerequisites

- Docker installed on your machine.
- Access to the internet to pull the Docker image.

## Docker Instructions

### Pull the Docker Image

To pull the Docker image, use the following command:
```bash
docker pull apeterson30/assignment01
```
Run the container and map port 8080:
```bash
docker run -p 8080:8080 apeterson30/assignment01
```

## The application will be accessible at http://localhost:8080.

Testing the Endpoints
1. GET /foo
```bash
curl http://localhost:8080/foo
```

2. POST /hello
```bash
curl -H "Accept: application/json" -H "Content-Type: application/json" -X POST --data '{"name": "YourName"}' http://localhost:8080/hello
```

3. GET /kill
```bash
curl http://localhost:8080/kill
```
