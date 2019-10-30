+++
title = "Requirements"
author = ["Iris Garcia"]
lastmod = 2019-10-27T18:54:25+01:00
draft = false
weight = 2
asciinema = true
+++

## Go {#go}

-   **Version used**: 1.13.1
-   **Setup**: Most of the linux distributions provide a package to
    install Go, but it is also possible to download its binary.

    ```bash
    wget https://dl.google.com/go/go1.13.3.src.tar.gz
    tar -C /usr/local -xzf go1.13.3.src.tar.gz
    export PATH=$PATH:/usr/local/go/bin
    ```


## Mage {#mage}

Once go is installed, we can install mage as follows:

```bash
go get github.com/magefile/mage
```


## MariaDB {#mariadb}

-   **Version used**: 10.4
-   **Setup**: Most of the linux distrubutions provide a package to
    install MariaDB, but if docker is installed the following command
    will spawn a MariaDB container:

    ```bash
    mage startdev
    ```


## (Production) Nodejs {#production--nodejs}

-   **Version used**: 12.11.1
-   Nodejs is only needed to install **PM2** and therefore it is not a
    requirement, however the current state of the project uses it to run
    the service in a production environment.