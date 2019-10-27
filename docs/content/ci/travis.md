+++
title = "Travis CI"
author = ["Iris Garcia"]
lastmod = 2019-10-18T00:54:18+02:00
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

The following lines tell travis to run the tests against two different
versions of Go (1.13.x and master)

```yaml
go:
- 1.13.x
- master
```

Travis is smart enough to setup all the required dependencies before
it start running the tests, basically it runs:

1.  **go get** which fetches all the dependencies defined in the file `go.mod`
2.  **go test** which runs the tests.

The `tests coverage` stage is exactly the same as the above but adds
the code coverage percentage which is useful to know if we are missing
some code from beign tested.

```yaml
- stage: tests coverage
  script: go test -cover -v
  go: 1.13.x
```

And the last stage `Release artifact to GitHub` runs only when there
is a tag push, it creates a compressed artifact with the source code
and place it in GitHub's releases.

```yaml
- stage: Release artifact to GitHub
  deploy:
    provider: releases
    api_key: $WORKDAY_RELEASE
    file: "workday"
    skip_cleanup: true
    on:
      tags: true
```
