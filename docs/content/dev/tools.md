+++
title = "Toolchain"
author = ["Iris Garcia"]
lastmod = 2019-11-05T13:21:18+01:00
draft = false
weight = 1
asciinema = true
+++

## [Go](https://golang.org/) {#go}

The open source programming language Go will be used to develop the
whole project, mainly because I want to try a new language and this
one is becoming quite popular nowadays.


## [Gin](https://github.com/gin-gonic/gin) {#gin}

Gin is a HTTP web framework written in Go (Golang). It features a
Martini-like API with much better performance -- up to 40 times
faster.

It is very well documented and provides many handy features like
authentication, data validation and a configurable logger out of the
box.


## [Mage](https://magefile.org/) {#mage}

Mage is a make/rake-like build tool using Go. You write plain-old go
functions, and Mage automatically uses them as Makefile-like runnable
targets.

Mage has no dependencies outside the Go standard library; in this
project it is going to be used to automate every possible process
like:

1.  `mage test`: runs the test suite.
2.  `mage testverbose`: runs the test suite with verbosity.
3.  `mage testandcoverage`: runs the test suite generating its code coverage.
4.  `mage build`: builds a binary of the project.
5.  `mage install`: installs the built binary under /usr/local/bin.
6.  `mage start`: Starts the HTTP Server using pm2 as process manager.
7.  `mage startdev`: Starts a dev environment, [WIP].
8.  `mage stop`: Stops the HTTP Server using pm2 as process manager.


## [MariaDB](https://mariadb.com/) {#mariadb}

The relational database engine MariaDB has been choosen to persist the
data, it is OpenSource and fulfills the requirements.


## Test-driven development (TDD & BDD) {#test-driven-development--tdd-and-bdd}

Go has support for testing built in to its toolchain which will be used to cover
unit and integration tests with the help of [testify](https://github.com/stretchr/testify) for the
assertions.

[Ginkgo](https://github.com/onsi/ginkgo) will be used as a BDD testing framework and [Gomega](https://github.com/onsi/gomega) as a
matcher library.


## [OpenAPI](https://github.com/OAI/OpenAPI-Specification/) {#openapi}

This project will follow the OpenAPI Specification to document its API
endpoints, probably using swagger to parse the specifications and
generate a static site.


## [PM2](https://pm2.keymetrics.io/) {#pm2}

As a process manager the tool **PM2** is currently used, because it
allows the usage with any language as long as it has a way to
build/run it.

Ideally this tool will be replace in favor of a Go tool, some of them
has been already tested but none of them provide the same features as
**PM2**.
