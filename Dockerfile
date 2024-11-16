# Build image
FROM golang:1.23.2-alpine AS build

WORKDIR /code

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /code/app cmd/server/main.go

# Main image with reduced size
FROM alpine:latest as app

WORKDIR /app
COPY --from=build /code/app .

CMD ["./app"]