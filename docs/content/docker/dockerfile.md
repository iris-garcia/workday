+++
title = "Dockerfile"
author = ["Iris Garcia"]
lastmod = 2019-12-09T18:37:03+01:00
tags = ["docker"]
draft = false
weight = 1
asciinema = true
+++

To create a Docker image, we have added two new files to our project:

-   `Dockerfile`: Contains the needed commands to create the image.
-   `.dockerignore`: It works like `.gitignore` for git, it allows to
    define which file/s should not be included in the Docker image when
    the command `ADD` or `COPY` is/are invoked.

<!--listend-->

```dockerfile
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
```

<div class="src-block-caption">
  <span class="src-block-number">Code Snippet 1</span>:
  Dockerfile
</div>

The Dockerfile has comments for every command but it is worth to mention
that we are using a multi stage build which allows to use different `FROM`
instructions, each of them begins a new stage of the build.

Then we can copy artifacts from one stage to another, leaving behind
everything we don't want in the final image.

```bash
docs
deployment
.git
.github
internal
api_server
.idea
coverage.html
workday.coverprofile
.travis.yml
LICENSE
Dockerfile
README.org
.packer/.vagrant
.packer/packer_cache
.packer/builds
```

<div class="src-block-caption">
  <span class="src-block-number">Code Snippet 2</span>:
  .dockerignore
</div>
