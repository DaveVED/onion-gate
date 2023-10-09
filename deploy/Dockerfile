# Build stage
FROM golang:1.20 AS builder

WORKDIR /src
COPY . .
RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o /app ./cmd/onion_gate/

# Final stage
FROM alpine:latest

# Add CA certificates and potential required libraries
RUN apk --no-cache add ca-certificates

# Copy app binary
COPY --from=builder /app /app

# Copy entire public directory (this includes templates and any other files/dirs inside public)
COPY --from=builder /src/public /public

EXPOSE 8080

ENTRYPOINT ["/app"]
