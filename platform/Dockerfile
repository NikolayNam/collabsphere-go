# Stage 1: Build
FROM golang:1.26.0-bookworm AS builder

WORKDIR /platform

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go vet -v ./...
RUN go test -v ./...

# Компилируем бинарник и сохраняем в /platform/platform
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o ./platform ./cmd/api

# Stage 2: Final image
FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y

# Копируем только бинарник из предыдущего этапа
COPY --from=builder /platform/platform /app
USER 65532:65532
CMD ["/app"]