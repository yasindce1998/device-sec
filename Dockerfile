# Dockerfile.server
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/server .
COPY --from=builder /app/config.yaml .
EXPOSE 8080
CMD ["./server"]

# Dockerfile.agent
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o agent ./cmd/agent

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/agent .
COPY --from=builder /app/config.yaml .
CMD ["./agent"]

# docker-compose.yml
version: '3.8'

services:
  postgres:
    image: postgres:latest
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: postgres
      POSTGRES_DB: device_management
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data

  rabbitmq:
    image: rabbitmq:management
    ports:
      - "5672:5672"
      - "15672:15672"

  server:
    build:
      context: .
      dockerfile: Dockerfile.server
    depends_on:
      - postgres
      - rabbitmq
    ports:
      - "8080:8080"

  agent:
    build:
      context: .
      dockerfile: Dockerfile.agent
    depends_on:
      - server

volumes:
  postgres_data: