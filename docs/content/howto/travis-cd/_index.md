+++
title = "Continuous Delivery (Travis CI)"
author = ["Iris Garcia"]
lastmod = 2019-12-06T18:24:53+01:00
tags = ["doc", "cd", "openshift"]
draft = false
weight = 2
asciinema = true
+++

Travis offers a deployment configuration which fortunately supports
OpenShift, the following steps shows how to get it working.


## Prerequisites {#prerequisites}


### Travis-CI command line client {#travis-ci-command-line-client}

The command line client has been used to secure our OpenShift token.
First of all we need to install the CLI following its [documentation](https://github.com/travis-ci/travis.rb#installation),
but basically if we already have `ruby` installed in our system we
just need to install the following gem:

```bash
gem install travis
```


### GitHub and Travis integration {#github-and-travis-integration}

Travis must be integrated in GitHub which can be done in a few steps:

1.  Go to <https://travis-ci.org/> and sign in with your GitHub account.
2.  In the upper right corner click on your name (or choose Accounts)
    to open your Travis-ci profile. You'll be presented with the list
    of your GitHub projects (only the ones where you have
    administrative authority)
3.  Toggle the switch on one of your projects and then click on little
    wrench icon to open the project's Integrations & services page on GitHub. To
    get to the same place from your GitHub project, go to Settings >
    Integrations & services.


## Step 1: Get the OpenShift token {#step-1-get-the-openshift-token}

![](/images/oc_login.png)
Save the token given by the Login command.


## Step 2: Get the travis API token {#step-2-get-the-travis-api-token}

In the upper right corner click on your name then Settings and again
the Settings tab as shown in the screenshot, then finally click in
`Copy token`.
![](/images/travis_token.png)


## Step 3: Login using the Tavis CLI {#step-3-login-using-the-tavis-cli}

```bash
travis login --github-token {{TOKEN}} --com
```

{{% notice note %}}
Replace `{{TOKEN}}` with the real token copied in the step 2.
{{% /notice %}}


## Step 4: Add and Secure the OpenShift token {#step-4-add-and-secure-the-openshift-token}

First add an environment variable in Travis:

```bash
travis env set OPENSHIFT_TOKEN {{TOKEN}} --com
```

Then add the encrypted value in our `.travis.yml` file:

```bash
travis encrypt --add deploy.token {{TOKEN}}
```

{{% notice note %}}
Replace `{{TOKEN}}` with the real token copied in the step 1.
{{% /notice %}}


## Step 5: Add deploy stage in `.travis.yml` {#step-5-add-deploy-stage-in-dot-travis-dot-yml}

The following excerpt shows only the **deploy** stage, the full
configuration file can be found [here](https://github.com/iris-garcia/workday/blob/master/.travis.yml).

```yaml
- stage: deploy
    name: Deploy to OpenShift
    script: skip
    deploy:
      provider: openshift
      server: https://api.us-east-2.starter.openshift-online.com:6443
      project: workday
      app: api
      edge: true
      on:
        branch: master
      token:
        secure: dQ/DwmYDyJ2JkhUh++II/1QgnIU/TAlobn//zki+G/Id9+Z4XU0DwGHb+WQuxS+RBqASS79imBkzd0b8uZsSgzf8mEFCEbzikZy3rYGJW/CVFVKygbOBRsM7ms+clAEAr9cet6QqKBeRt6WH3AiPfetcNw0GpjKYr0WGdzzq+sf347NRFrhr/rSiOeugBq2EYqtuXeE6tAzm0ivGLl9C4hDYBdkYiQfJ16hk+/hJrwFRZpVv+7yR9J+WphMVqbCrB0XY3qSnwUlgfMw5QdCFvZAqoZbbiIF0OqEDZ+kwSVSPKPZ/zybpyrE+ty83GGuQ3MymMLM35Upr51HB6VNAcwtpwW8Cf3Bzj2odFKzk26etvUDhaPpXMV8Ow9VgYgweEti9KebdM0esN5emr/7vCmLVe3ppNDhH+tfGGmaVM8dkB+L4d2A4kXoxfHyS59HZPGBVFPLmNrxgwxbVaO7EiqUPlBX7SOMMNKn83HUF96edCOXwqVdznfLaG9Uh1/pvfTj4N1NOO1zTdTuuda4WeXSAyWEpgc15RwNQcYp6smtgXk3zFYKA0ZB9C9jyO01Fvoy96H8llY+wrEVuiUmyzSu3KAk6+86SLPJQUHWsvhSTES7qb6c5oSmoBao7X97b4/3EOGHq86wJLE/6vjrqWlrq3BtXpXqiOcbB5el1a9M=
```

-   The `server` parameter must be set accordingly with your OpenShift
    endpoint.
-   The `project` and `app` parameters must match the ones provided in
    the deployment commands.
-   The `edge` parameter tells Travis to use the deployment version 2
    which is currently in _beta_ (but still recommended by Travis [here](https://docs.travis-ci.com/user/deployment/openshift/)).
-   The `script: skip` parameter tells Travis to don't run any build
    command (we don't have to waste time here).
-   Finally the `on: branch` parameter tells Travis to only deploy when
    there are changes in the **master** branch.
