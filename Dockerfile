# Build stage
FROM golang:1.23-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

# Copy go module files first for layer caching
COPY go.mod go.sum ./

# Copy the haruki-cloud dependency (relative path in replace directive)
COPY ../Haruki-Cloud /haruki-cloud

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /haruki-command-parser ./cmd/server

# Runtime stage
FROM alpine:3.20

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

COPY --from=builder /haruki-command-parser .
COPY configs.yaml .

EXPOSE 24001

ENV CONFIG_PATH=/app/configs.yaml

CMD ["/app/haruki-command-parser"]
