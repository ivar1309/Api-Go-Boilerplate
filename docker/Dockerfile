# Stage 1: Builder
FROM golang:1.24.4-alpine AS builder

WORKDIR /code

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main cmd/main.go

# Stage 2: Final stage using scratch
FROM scratch

COPY --from=builder /code/main /main

CMD ["./main"]