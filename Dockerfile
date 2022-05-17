# syntax=docker/dockerfile:1

FROM golang:1.18.2-alpine

WORKDIR /app

# Setup dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Setup project code
COPY *.go ./
RUN go build -o /duogo

EXPOSE 8080

CMD [ "/duogo" ]