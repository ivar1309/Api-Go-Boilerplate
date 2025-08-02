# Stage 1: Builder
FROM golang:1.24.4-alpine AS builder

# Install sqlc
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

WORKDIR /code

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source
COPY . .

# Generate sqlc (if not pre-generated)
RUN sqlc generate -f ./sql/sqlc.yaml

# Build the binary
RUN go build -o main cmd/main.go

# Stage 2: Final stage using scratch
FROM scratch

COPY --from=builder /code/main /main
COPY --from=builder /code/sql/migrations /sql/migrations

CMD ["./main"]