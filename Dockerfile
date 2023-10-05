# Use the official Go image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.16 as builder

# Copy local code to the container image.
WORKDIR /go/src/github.com/your-organization/product-service
COPY . .

# Build the command inside the container.
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -v -o product-service ./cmd/main.go

# Use a Docker multi-stage build to create a lean production image.
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine:3.14

# Copy the binary to the production image from the builder stage.
COPY --from=builder /go/src/github.com/your-organization/product-service/product-service /product-service

# Run the web service on container startup.
CMD ["/product-service"]
