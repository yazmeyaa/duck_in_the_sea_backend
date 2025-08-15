FROM golang:1.24-alpine AS builder

WORKDIR /build

ENV CGO_ENABLED=0 \
    GOPATH=/go \
    GOCACHE=/go-build

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod/cache \
    go mod download

COPY . .

RUN --mount=type=cache,target=/go/pkg/mod/cache \
    --mount=type=cache,target=/go-build \
    go build -o /build/backend ./cmd/http_server

FROM alpine:latest
WORKDIR /
COPY --from=builder /build/backend .
EXPOSE 8080
CMD ["./backend"]