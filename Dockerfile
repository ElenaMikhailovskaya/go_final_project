FROM ubuntu:latest

RUN apt update && apt install golang -y ca-certificates && update-ca-certificates

WORKDIR /build
COPY go.mod ./
COPY go.sum ./
COPY ./cmd/ ./cmd/
COPY ./internal/ ./internal/
COPY ./web/ ./web/

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /app ./cmd/app