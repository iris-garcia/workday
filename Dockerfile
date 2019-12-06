# Use a golang image as base builder image
FROM golang:1.13 as builder

# Create the project's directory under the default GOPATH
RUN mkdir -p /go/src/github.com/iris-garcia/workday

# Use this directory as working directory
WORKDIR /go/src/github.com/iris-garcia/workday

# Copy the needed files to build the binary
COPY . /go/src/github.com/iris-garcia/workday/

# Install our build tool Mage
RUN go get github.com/magefile/mage

# Run the build stage
RUN mage build

# Use a small image to run the binary
FROM busybox:latest

# Use /root as working directory
WORKDIR /root

# Copy the built binary and its default config file
COPY --from=builder /go/src/github.com/iris-garcia/workday/api_server .
COPY --from=builder /go/src/github.com/iris-garcia/workday/db_config.toml .

# Run the API server
CMD ["./api_server"]
