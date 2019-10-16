+++
title = "Travis CI"
author = ["Iris Garcia"]
lastmod = 2019-10-16T20:40:18+02:00
draft = false
weight = 3
asciinema = true
+++

## Description {#description}

Travis is a hosted continuous integration service used to build and
test software projects hosted at GitHub.

It provides a free plan for open source projects which is very
convinient for our use case.

The whole configuration is set up in a single file `.travis.yml` which
must be placed in the root directory of the project.


## Configuration {#configuration}

To avoid replicating exactly the same workflows in GitHub Actions and
Travis, there are some little changes like testing two different Go
versions: `v1.13.x` and `master` (which is the latest available at any
given time).

It also releases artifacts when a new **tag** is pushed to the
repository.

```yaml
language: go

go:
- 1.13.x
- master

jobs:
  include:
    - stage: tests coverage
      script: go test -cover -v
      go: 1.13.x

    - stage: Release artifact to GitHub
      deploy:
        provider: releases
        api_key: $WORKDAY_RELEASE
        file: "workday"
        skip_cleanup: true
        on:
          tags: true
```
