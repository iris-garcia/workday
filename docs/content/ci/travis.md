+++
title = "Travis CI"
author = ["Iris Garcia"]
lastmod = 2019-11-16T19:58:26+01:00
tags = ["ci"]
draft = false
weight = 1
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
dist: bionic
sudo: required

stages:
  - test
  - deploy

jobs:
  include:
    - stage: test
      go: 1.13.x
      before_install:
        - sudo apt-get -y install npm
        - npm install pm2@latest -g
      install:
        - go get github.com/magefile/mage
        - mage build
      addons:
        mariadb: '10.1'
      services:
      - mariadb
      before_script:
      - mysql -u root -e 'CREATE DATABASE workday;'
      - mysql -u root -e "CREATE USER 'workday'@'%' IDENTIFIED BY 'workday';"
      - mysql -u root -e "GRANT ALL ON workday.* TO 'workday'@'%'; FLUSH PRIVILEGES;"
      script:
        - mage start
        - mage test
        - mage stop

    - stage: deploy
      name: "Deploy to OpenShift"
      script: skip
      deploy:
        provider: openshift
        server: https://api.us-east-2.starter.openshift-online.com:6443
        project: workday
        app: api
        edge: true # opt in to dpl v2
        on:
          branch: master
        token:
          secure: dQ/DwmYDyJ2JkhUh++II/1QgnIU/TAlobn//zki+G/Id9+Z4XU0DwGHb+WQuxS+RBqASS79imBkzd0b8uZsSgzf8mEFCEbzikZy3rYGJW/CVFVKygbOBRsM7ms+clAEAr9cet6QqKBeRt6WH3AiPfetcNw0GpjKYr0WGdzzq+sf347NRFrhr/rSiOeugBq2EYqtuXeE6tAzm0ivGLl9C4hDYBdkYiQfJ16hk+/hJrwFRZpVv+7yR9J+WphMVqbCrB0XY3qSnwUlgfMw5QdCFvZAqoZbbiIF0OqEDZ+kwSVSPKPZ/zybpyrE+ty83GGuQ3MymMLM35Upr51HB6VNAcwtpwW8Cf3Bzj2odFKzk26etvUDhaPpXMV8Ow9VgYgweEti9KebdM0esN5emr/7vCmLVe3ppNDhH+tfGGmaVM8dkB+L4d2A4kXoxfHyS59HZPGBVFPLmNrxgwxbVaO7EiqUPlBX7SOMMNKn83HUF96edCOXwqVdznfLaG9Uh1/pvfTj4N1NOO1zTdTuuda4WeXSAyWEpgc15RwNQcYp6smtgXk3zFYKA0ZB9C9jyO01Fvoy96H8llY+wrEVuiUmyzSu3KAk6+86SLPJQUHWsvhSTES7qb6c5oSmoBao7X97b4/3EOGHq86wJLE/6vjrqWlrq3BtXpXqiOcbB5el1a9M=
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
