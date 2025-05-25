# ------------------------ BUILD STAGE
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
WORKDIR /app/cmd/api
RUN go build -o api

# ------------------------ RUNTIME STAGE
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/cmd/api/api .

ENTRYPOINT ["./api"]
