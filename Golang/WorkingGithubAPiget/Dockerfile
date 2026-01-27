# Builder
FROM golang:1.22-alpine AS builder
WORKDIR /app

# Download dependencies first
COPY go.mod ./
RUN go mod download

# Copy all source files
COPY . .

# Build static Linux binary including all Go files
RUN CGO_ENABLED=0 GOOS=linux GOARCH=$(go env GOARCH) go build -o server .

# Final image
FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=builder /app/server /server

EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/server"]

