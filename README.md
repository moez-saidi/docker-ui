# Docker UI
docker-ui is a REST API built with Go and Gin that provides endpoints to interact with Docker.

## Prerequisites
- **Docker**: Make sure Docker is installed and running.
- **Go**: The API is written in Go (go1.23.*).

## Usage
### Endpoints
- **GET** `api/<version>/containers`: Returns a JSON array of running Docker containers.
- **GET** `api/<version>/images`: Returns a JSON array of Docker images on the host.

## Getting started
To start the server:

```bash
go run ./cmd/server
```