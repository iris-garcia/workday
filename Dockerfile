# Use the alpine version of golang which is way smaller.
FROM golang:1.13

# Create a new user api
RUN useradd -ms /bin/bash api

# Use our user instead of root
USER api

# Create the project's directory under the default GOPATH
RUN mkdir -p /go/src/github.com/iris-garcia/workday

# Use this directory as working directory
WORKDIR /go/src/github.com/iris-garcia/workday

# Copy the needed files to build the project (using .dockerignore to ignore unneeded)
COPY --chown=api . /go/src/github.com/iris-garcia/workday

# Install our build tool Mage
RUN go get github.com/magefile/mage

# Run the build stage
RUN mage build

# Run the startapi stage
CMD ["mage", "startapi"]
