FROM golang:1.18 as builder
WORKDIR /app

# Retrieve application dependencies using go modules.
# Allows container builds to reuse downloaded dependencies.
COPY rdp_bot/go.* ./
RUN go mod download

# Copy local code to the container image.
# Skip .env

COPY rdp_bot ./


# Build the binary.
# -o myapp specifies the output name of the binary.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o rdp_bot

# Use the official Alpine image for a lean production container.
FROM alpine:latest
RUN apk --no-cache add ca-certificates

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/rdp_bot /serve_rdp_bot

# Run the web service on container startup.
CMD ["/serve_rdp_bot"]