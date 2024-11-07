
FROM golang:1.23.2-alpine AS build

WORKDIR /code
COPY . .

RUN go mod download
RUN go build -o /code/app cmd/server/main.go

FROM alpine:latest as app

WORKDIR /app
COPY --from=build /code/app .

CMD ["./app"]