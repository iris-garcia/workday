+++
title = "Travis CI"
author = ["Iris Garcia"]
lastmod = 2019-10-31T09:08:46+01:00
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

dist: bionic

sudo: required

addons:
  mariadb: '10.1'

services:
  - mariadb

before_install:
  - sudo apt-get -y install npm
  - npm install pm2@latest -g

install:
  - go get github.com/magefile/mage
  - mage build

before_script:
  - mysql -u root -e 'CREATE DATABASE workday;'
  - mysql -u root -e "CREATE USER 'workday'@'%' IDENTIFIED BY 'workday';"
  - mysql -u root -e "GRANT ALL ON workday.* TO 'workday'@'%'; FLUSH PRIVILEGES;"

script:
  - mage start
  - mage test
  - mage stop
```

The following lines tell travis to run the tests against two different
versions of Go (1.13.x and master)

```yaml
go:
- 1.13.x
- master
```

And the following excerpt defines **bionic** as the Ubuntu version to be
used also tells the requirement of **sudo** needed to install `npm` and
`pm2`.

```yaml
dist: bionic

sudo: required

addons:
  mariadb: '10.1'

services:
  - mariadb
```

The **addons** and **services** sections are in place to boot a **mariadb**
service which will be used for the integration tests against a test
database.

The **before\_install** and **install** configurations will fetch and setup
all the prerequisites in order to run our **build tool** which is
**mage**.

Finally the **script** will use `mage` to start the environment, run the
tests and stop it once it finishes.
