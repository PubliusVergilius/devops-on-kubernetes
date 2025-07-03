FROM golang:1.24.4-alpine3.22 as builder

# Install git for go modules (if needed)
RUN apk add --no-cache git 

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download 

COPY . .

# Build the static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o myapi ./cmd/myapi

# Stage 2: Create minimal runtime image
FROM gcr.io/distroless/static

# Copy binary from build stage
COPY --from=builder /app/myapi /myapi

# Set non-root user (distroless has `nonroot`)
USER nonroot:nonroot

# Run the app
ENTRYPOINT ["API_PORT=8080 /myapi"]
