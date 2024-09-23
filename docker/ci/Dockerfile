FROM golang:1.21.0

WORKDIR /build
COPY go.mod ./
COPY go.sum ./
COPY ./cmd/ ./cmd/
COPY ./internal/ ./internal/

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /app ./cmd/app