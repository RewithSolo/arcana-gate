# Build Stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod ./
# COPY go.sum ./ 
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-s -w" \
    -o /app/arcana-gate ./cmd/arcana-gate

FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /

COPY --from=builder /app/arcana-gate /arcana-gate

USER nonroot:nonroot

ENTRYPOINT ["/arcana-gate"]