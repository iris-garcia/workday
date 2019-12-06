+++
title = "OpenShift"
author = ["Iris Garcia"]
lastmod = 2019-12-06T18:24:52+01:00
tags = ["cd"]
draft = false
weight = 1
asciinema = true
+++

## Requirements {#requirements}


### Create account in RedHat {#create-account-in-redhat}

1.  Create a new account using the [Sign up](https://manage.openshift.com/accounts/auth/keycloak) form.
2.  Sign in into [OpenShift](https://manage.openshift.com/accounts/auth/keycloak).


### Install the OpenShift client {#install-the-openshift-client}

```bash
wget https://mirror.openshift.com/pub/openshift-v4/clients/oc/4.1/linux/oc.tar.gz
tar xvzf oc.tar.gz
sudo cp oc /usr/local/bin/
```

{{% notice tip %}}
**The binary can be copied to any place as long as it is included in the $PATH.**
{{% /notice %}}


### Get the Login command {#get-the-login-command}

{{< figure src="/images/oc_login.png" >}}


### Create a Dockerfile {#create-a-dockerfile}

OpenShift can create a docker image given a source git repository, but
it needs a Dockerfile definition which is shown [here](/docker/dockerfile).


## QuickStart {#quickstart}

This is a summary of the needed commands to get the project working:

```bash
# Login
oc login --token={REDACTED} --server=https://api.us-east-2.starter.openshift-online.com:6443
# New project
oc new-project workday
# Not needed: mariadb service
oc new-app \
        -e MYSQL_USER=workday \
        -e MYSQL_PASSWORD=workday \
        -e MYSQL_DATABASE=workday \
        mariadb
# Our API service
oc new-app deployment/openshift.yml
# Manually trigger a new build
oc start-build api
```

{{% notice tip %}}
**For further explanation keep reading the step by step guide.**
{{% /notice %}}


## Step 1: Login using the CLI {#step-1-login-using-the-cli}

Use the login command from the previous step:

```bash
oc login --token={REDACTED} --server=https://api.us-east-2.starter.openshift-online.com:6443
```


## Step 2: Create a new project {#step-2-create-a-new-project}

```bash
oc new-project workday
```


## Step 3: Create a new MariaDB app {#step-3-create-a-new-mariadb-app}

```bash
oc new-app \
      -e MYSQL_USER=workday \
      -e MYSQL_PASSWORD=workday \
      -e MYSQL_DATABASE=workday \
      mariadb
```

{{% notice warning %}}
**This step is not needed yet, but it will be in a near future to get the service fully working.**
{{% /notice %}}


## Step 4: Create a template for our service {#step-4-create-a-template-for-our-service}

```yaml
kind: Template
apiVersion: v1
metadata:
  name: api
  annotations:
    description: "Workday Api app"
    tags: "workday,api,golang"
    iconClass: "icon-go-gopher"
  labels:
    template: "api"
    app: "api"
objects:
- kind: Service
  apiVersion: v1
  metadata:
    name: api
    annotations:
      description: "Exposes and load balances the application pods"
  spec:
    ports:
    - name: web
      port: 8080
      targetPort: 8080
    selector:
      name: api
- kind: Route
  apiVersion: v1
  metadata:
    name: api
  spec:
    to:
      kind: Service
      name: api
- kind: ImageStream
  apiVersion: v1
  metadata:
    name: api
    annotations:
      description: "Keeps track of changes in the application image"
- kind: Secret
  apiVersion: v1
  metadata:
    name: gh-secret
    creationTimestamp:
  data:
    WebHookSecretKey: "${GITHUB_SECRET}"
- kind: BuildConfig
  apiVersion: v1
  metadata:
    name: api
    annotations:
      description: "Defines how to build the application"
  spec:
    source:
      type: Git
      git:
        uri: "${SOURCE_REPOSITORY_URL}"
        ref: "${SOURCE_REPOSITORY_REF}"
      contextDir: "${CONTEXT_DIR}"
    strategy:
      type: Docker
      dockerStrategy: {}
    output:
      to:
        kind: ImageStreamTag
        name: api:latest
    postCommit:
      script: "GIN_MODE=release go test -v ./..."
    resources:
      limits:
        cpu: 100m
        memory: 1Gi
    triggers:
    - type: "GitHub"
      github:
        secretReference:
          name: "gh-secret"
- kind: DeploymentConfig
  apiVersion: v1
  metadata:
    name: api
    annotations:
      description: "Defines how to deploy the application server"
  spec:
    strategy:
      type: Recreate
    triggers:
    - type: ImageChange
      imageChangeParams:
        automatic: true
        containerNames:
        - api
        from:
          kind: ImageStreamTag
          name: "api:latest"
    - type: ConfigChange
    replicas: 1
    selector:
      name: api
    template:
      metadata:
        name: api
        labels:
          name: api
      spec:
        containers:
        - name: api
          image: api
          ports:
          - containerPort: 8080
          env:
          - name: GIN_MODE
            value: "release"
          - name: WORKDAY_DB_NAME
            value: "workday"
          - name: WORKDAY_DB_USER
            value: "workday"
          - name: WORKDAY_DB_PASSWORD
            value: "workday"

parameters:
- name: SOURCE_REPOSITORY_URL
  description: "The URL of the repository with your application source code"
  value: "https://github.com/iris-garcia/workday.git"
- name: SOURCE_REPOSITORY_REF
  description: "Set this to a branch name, tag or other ref of your repository if you are not using the default branch"
- name: CONTEXT_DIR
  description: "Set this to the relative path to your project if it is not in the root of your repository"
- name: GITHUB_SECRET
  description: "Github webhook secret"
```

<div class="src-block-caption">
  <span class="src-block-number">Code Snippet 1</span>:
  openshift.yml
</div>

The template defines 5 main resources to be created in OpenShift:

1.  **Service**: exposes connectivity to the API container in a port.
2.  **Route**: Creates and endpoint which points to the Service resource.
3.  **ImageStream**: Keeps track of of changes in the application image.
4.  **BuildConfig**: This resource as it is configured basically fetches
    the source code from the git repository then builds and tags a
    docker image and finally runs the tests to make sure it is a
    working image.
5.  **DeploymentConfig**: This resource is the one spawning a new
    container/s with the built image, it has many parameters such as
    the ability to recreate the container automatically when the
    specified image has been changed or the number of replicas, it also
    allows to setup environment variables to the container/s.


## Step 5: Create a new app using the above template {#step-5-create-a-new-app-using-the-above-template}

```bash
oc new-app deployment/openshift.yml -p GITHUB_SECRET=supersecret
```


## Step 6: Trigger the build {#step-6-trigger-the-build}

The build can be triggered manually with the following command:

```bash
oc start-build api
```

But we don't want to be triggering a new build everytime we integrate
a new change in our service, therefore we have also set up a continuous
delivery processes which deploys a new version of the app whenever there
are new changes in the repository.

1.  Using Travis CI ([how-to documentation](/howto/travis-cd/)).
2.  Using GitHub ([how-to documentation](/howto/github-cd/)).


## Demo {#demo}

{{< asciinema key="openshift" rows="50" preload="1"
idle-time-limit="1" >}}
