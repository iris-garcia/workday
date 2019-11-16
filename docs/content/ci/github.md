+++
title = "GitHub Actions"
author = ["Iris Garcia"]
lastmod = 2019-11-16T19:58:26+01:00
tags = ["ci"]
draft = false
weight = 2
asciinema = true
+++

## Description {#description}

GitHub Actions is the new continuous integration and deployment system
built and maintained by the community.


## Configuration {#configuration}

Currently there are two workflows configured:


### Test with verbosity enabled {#test-with-verbosity-enabled}

To make it a bit different that the pipeline configured in Travis CI,
this one will run the tests with verbosity enabled, this way it
outputs every spec and API call done for each test case.

An example of one run can be seen [here](https://github.com/iris-garcia/workday/runs/285177520).

```yaml
on: push

name: Unit tests

jobs:
  checks:
    name: run
    runs-on: ubuntu-latest

    steps:
    - uses: actions/checkout@master

    - name: run
      uses: cedrickring/golang-action@1.4.1
      env:
        GO111MODULE: "on"
      with:
        args: |
          go get github.com/magefile/mage && \
          go get github.com/onsi/ginkgo/ginkgo && \
          mage build && \
          mage testverbose
```

There is just one job configured in this **Action** with the name
`checks` and as stated in the line 8 of the configuration it uses
**Ubuntu** in it latest available version.

There are two steps in this _job_:

The first one Checks out the project repository in its `master` branch.

The second one uses an action to automatically setup a Go workspace
and run arbitrary commands, the documentation can be seen [here](https://github.com/cedrickring/golang-action).
If no args are specified and a `Makefile` is detected, this action will
run `make`. Otherwise `go test` and `go build` will be run.

In this case it is overwritten in order to install `mage` and `ginkgo`
CLIs to allow the build and run of the tests.


### Hugo documentation site {#hugo-documentation-site}

This workflow is not really a typical continuous integration one, but
I think it makes sense to mention it here as it is taking care of
automatically update the documentation site.

```yaml
on:
  push:
    paths:
      - 'docs/**'
      - '.github/workflows/hugo.yml'

name: Hugo

jobs:
  build:
    runs-on: ubuntu-latest

    steps:
    - uses: actions/checkout@v1
    - name: Installs hugo
      run: |
        cd /tmp
        wget https://github.com/gohugoio/hugo/releases/download/v0.57.0/hugo_0.57.0_Linux-64bit.deb
        sudo dpkg -i hugo_0.57.0_Linux-64bit.deb
        hugo version

    - name: Build hugo site
      run: |
        cd docs
        rm -rf public
        git worktree add -b gh-pages public origin/gh-pages
        hugo

    - name: Configure git and deployment key
      env:
        GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
        GITHUB_DEPLOY_KEY: ${{ secrets.WORKDAY_GH }}
      run: |
        mkdir /home/runner/.ssh
        ssh-keyscan -t rsa github.com >/home/runner/.ssh/known_hosts
        echo "${GITHUB_DEPLOY_KEY}" > /home/runner/.ssh/id_rsa && \
        chmod 400 /home/runner/.ssh/id_rsa
        git remote set-url origin git@github.com:iris-garcia/workday.git
        git config --global user.name "GitHub Action"
        git config --global user.email "action@github.com"

    - name: Commit and push changes to gh-pages
      run: |
        cd docs/public
        git add .
        git commit -m "Publishing to gh-pages branch"
        git push origin gh-pages
```

To get this one working there are some requisites explained in a howto
[document](/howto/gh-pages).
