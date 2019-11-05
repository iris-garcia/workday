+++
title = "API Test Class"
author = ["Iris Garcia"]
lastmod = 2019-11-05T17:06:27+01:00
tags = ["bdd", "test"]
draft = false
weight = 3
asciinema = true
+++

## User story {#user-story}

As an admin employee, I want to be able to:

1.  Add new employees to the system.
2.  Edit already existing employees' details.
3.  List all the employees existing in the system.
4.  Retrieve the details of any employee given its ID.
5.  Delete an employee from the system given its ID.


## Endpoints {#endpoints}

| Method | Endpoint        |
|--------|-----------------|
| GET    | /status         |
| GET    | /employees      |
| POST   | /employees      |
| GET    | /employees/{id} |
| PUT    | /employees/{id} |
| DELETE | /employees/{id} |

-   **GET /status**: Returns a 200 and `{"status": "OK"}` in the body.
-   **GET /employees**: Returns a 200 and the list of employees stored in the database.
-   **POST /employees**: Adds a new employee into the database if the body
    fulfills the requirements (firstname, lastname, role, password),
    otherwise it returns a 204 reponse.
-   **GET /employees/{id}**: Returns a 200 and the details of an employee
    if the id is found otherwise a 404.
-   **DELETE /employees/{id}**: Returns a 200 and an OK message if the
    employee with **{id}** is found otherwise a 500 and the error.
-   **PUT /employees/{id}**: Returns a 200 and an OK message if the
    employee with **{id}** is found, if it is not found a 404 and if any
    other error happens a 500 with the error message.

BDD has been used to test every possible use case to reach a 100%
[code coverage](/coverage.html) it is worth to mention that the database has been mocked so
the current tests are **Unit tests**, leaving the **Integration tests**
for a later iteration.


## Source code {#source-code}

The code which covers the current class can be found [here](https://github.com/iris-garcia/workday/blob/master/api/router%5Ftest.go).


## Demo {#demo}

{{< asciinema key="employees\_test" rows="50" preload="1"
idle-time-limit="1" >}}
