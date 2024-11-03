# Docker UI
docker-ui is a REST API built with Go and Gin that provides endpoints to interact with Docker.


## Features
List Containers: Retrieve a list of running Docker containers.
List Images: Retrieve a list of Docker images available on the host.

## Prerequisites
- **Docker**: Make sure Docker is installed and running.
- **Go**: The API is written in Go (go1.23.0).

## Usage
### Endpoints
- **GET** `/containers`: Returns a JSON array of running Docker containers.
- **GET** `/images`: Returns a JSON array of Docker images on the host.
`