+++
title = "Docker Hub"
author = ["Iris Garcia"]
lastmod = 2019-12-09T18:37:04+01:00
draft = false
weight = 3
asciinema = true
+++

Now that we have a working Docker image we want to automate the
building process, so it automatically builds and publishes the image with
every new code change.


## Step 1: Register in [Docker Hub](https://hub.docker.com) {#step-1-register-in-docker-hub}

Just click in `Sing up` and fill up the following form:

{{< figure src="/images/dockerhub_1.png" >}}


## Step 2: Create a new Repository {#step-2-create-a-new-repository}

Click in **Create Repository +** and name it accordingly, under the
**Build Settings** section select the **Connected** option with a GitHub
logo to autobuild a new image with every `git push` event as long as
it matches the build rules, which in our case are the default ones
(the Dockerfile is located in the root directory of the project).

{{< figure src="/images/dockerhub_2.png" >}}

Click in **Create and Build** to start the first build of the image,
then in the **Builds** tab we can see the process of the build and its
logs as shown in the screenshot below:

{{< figure src="/images/dockerhub_3.png" >}}

{{% notice tip %}}
**Now with every new change detected in our GitHub repository a new build will be triggered in Docker Hub.**
{{% /notice %}}
