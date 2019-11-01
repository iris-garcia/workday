+++
title = "What is Workday?"
author = ["Iris Garcia"]
lastmod = 2019-11-01T19:47:27+01:00
draft = false
weight = 1
asciinema = true
+++

Workday is a RESTful API project that allows the management of
workers' day registration, for this purpose it makes use of a
relational database where such data persists.

It allows the creation of different types of users (roles), initially
there are two main roles:

1.  `employee`: The employee has permissions to:
    -   Register checks in and checks out.
    -   Update their own password.
    -   Update their own schedule.
2.  `HR`: The Human Resources role is the one with almost full control,
    it has the permissions of a regular employee plus:
    -   Retrieve every employee's schedule.
    -   Register employees.
    -   Remove employees.
    -   Reset employees' password.

The final goal is to build clients consuming this API to improve the
user experience, for example:

An Android/iOS app which automatically registers the checks in when the
GPS location is near the Office's location and the checks out when it
gets away from the Office's location.
