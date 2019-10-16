+++
title = "GitHub Actions"
author = ["Iris Garcia"]
lastmod = 2019-10-16T20:41:58+02:00
draft = false
weight = 3
asciinema = true
+++

## Description {#description}

GitHub Actions is the new continuous integration and deployment system
built and maintained by the community.


## Configuration {#configuration}

Currently there are two workflows configured:


### Test and coverage {#test-and-coverage}

This workflow takes care of the Unit tests and shows the total
coverage of the project.

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
        args: go build && go test -cover -v
```


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
