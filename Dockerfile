# Build stage
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/server

# Run stage
FROM debian:bullseye-slim

RUN apt-get update && apt-get install -y ca-certificates && apt-get clean

RUN useradd -m appuser

WORKDIR /app

COPY --from=builder /app/main .

# Purpose of this command
# Copy template files
# In Dockerfile, I only copy the executable file, not the templates folder
# This results in the templates folder not existing in the container.
COPY --from=builder /app/templates ./templates

# Create logs folder and permissions
# Purpose of this command
# In Dockerfile, I am using non-root user
# And because it is a user right, it does not have the right to create a folder
RUN mkdir -p storages/logs && chown -R appuser:appuser storages

COPY app.env .
COPY local.yaml .

USER appuser

EXPOSE 5004

CMD ["./main"]