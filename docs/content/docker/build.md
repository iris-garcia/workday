+++
title = "Building the Docker image"
author = ["Iris Garcia"]
lastmod = 2019-12-06T18:28:28+01:00
draft = false
weight = 2
asciinema = true
+++

To build the Docker image locally we can run the following command
from the root directory of the project:

```bash
$ docker build . -t workday

Sending build context to Docker daemon  118.8kB
Step 1/11 : FROM golang:1.13 as builder
 ---> a2e245db8bd3
Step 2/11 : RUN mkdir -p /go/src/github.com/iris-garcia/workday
 ---> Running in 61717fb25087
Removing intermediate container 61717fb25087
 ---> eaf6e4d73eb4
Step 3/11 : WORKDIR /go/src/github.com/iris-garcia/workday
 ---> Running in 41ddb2a701bd
Removing intermediate container 41ddb2a701bd
 ---> d5b6af426ca6
Step 4/11 : COPY . /go/src/github.com/iris-garcia/workday/
 ---> d64c634bdf5d
Step 5/11 : RUN go get github.com/magefile/mage
 ---> Running in 390e199e05f6
go: finding github.com/magefile/mage v1.9.0
go: downloading github.com/magefile/mage v1.9.0
go: extracting github.com/magefile/mage v1.9.0
Removing intermediate container 390e199e05f6
 ---> a8cf88a8c7fe
Step 6/11 : RUN mage build
 ---> Running in b2de42b4f44e
Installing Deps...
Building...
Removing intermediate container b2de42b4f44e
 ---> cdb82f1fa6e2
Step 7/11 : FROM busybox:latest
 ---> 020584afccce
Step 8/11 : WORKDIR /root
 ---> Running in 4fcc3c788a2c
Removing intermediate container 4fcc3c788a2c
 ---> 3f616f2908d0
Step 9/11 : COPY --from=builder /go/src/github.com/iris-garcia/workday/api_server .
 ---> f3a03d0b5dd7
Step 10/11 : COPY --from=builder /go/src/github.com/iris-garcia/workday/db_config.toml .
 ---> 1267efc3b59c
Step 11/11 : CMD ["./api_server"]
 ---> Running in 01b44065e48f
Removing intermediate container 01b44065e48f
 ---> d7a90aacccb6
Successfully built d7a90aacccb6
Successfully tagged workday:latest
```

This is basically telling docker to build a Docker image reading the
`Dockerfile` from the directory `.` (current working directory) and
name it (`-t`) `workday`.

As we did not especify any tag after the name Docker defaults to
`latest`.

To check our newly created image we can run the command:

```bash
docker image ls

REPOSITORY      TAG                 IMAGE ID            CREATED             SIZE
workday         latest              d7a90aacccb6        3 minutes ago       17.2MB
<none>          <none>              cdb82f1fa6e2        3 minutes ago       910MB
```

{{% notice tip %}}
**Thanks to the multi stage build we have our app in a 17.2MB container.**
{{% /notice %}}

And finally we can run our image as follows:

```bash
docker run -e PORT=8080 -p 9090:8080 workday
```

This will set the environment variable `PORT` to the value `8080` (this
is internally used by our app), then the `-p 9090:8080` will publish the
container port **8080** to the host port **9090**, therefore using
<http://localhost:9090> should hit our container app.
