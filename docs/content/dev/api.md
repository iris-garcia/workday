+++
title = "API Test Class"
author = ["Iris Garcia"]
lastmod = 2019-10-31T09:41:22+01:00
draft = false
weight = 2
asciinema = true
+++

## Description {#description}

Currently the API allows the following actions:

| Method | Endpoint        |
|--------|-----------------|
| GET    | /status         |
| GET    | /employees      |
| POST   | /employees      |
| GET    | /employees/{id} |
| DELETE | /employees/{id} |
| PUT    | /employees/{id} |

BDD has been used to test every possible use case to reach a 100%
[code coverage](/coverage.html) it is worth to mention that the database has been mocked so
the current tests are **Unit tests**, leaving the **Integration tests**
for a later iteration.


## Test code {#test-code}

The code which covers the current class can be found [here](https://github.com/iris-garcia/workday/blob/master/api/router%5Ftest.go).


## Demo {#demo}

{{< asciinema key="coverage" rows="50" preload="1"
idle-time-limit="1" >}}
