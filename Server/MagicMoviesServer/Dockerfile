# syntax=docker/dockerfile:1
FROM golang:1.20-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# Build statically
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o /app/server ./main.go

FROM alpine:3.18
RUN addgroup -S app && adduser -S app -G app
WORKDIR /app
COPY --from=builder /app/server /app/server
# copy env if needed: COPY --from=builder /src/.env /app/.env
USER app
EXPOSE 8080
CMD ["/app/server"]