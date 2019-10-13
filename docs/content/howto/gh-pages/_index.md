+++
title = "GitHub Pages"
author = ["Iris Garcia"]
lastmod = 2019-10-13T13:31:25+02:00
tags = ["ci", "hugo", "doc"]
draft = false
weight = 1
asciinema = true
+++

This document ilustrates how to setup [GitHub Pages](https://pages.github.com/) using [Hugo](https://gohugo.io/) as
website generator and [GitHub's Actions](https://github.com/features/actions) to automate the deployment
process.

There are different alternatives to setup GitHub pages, the one used
in here is a project pages using `gh-pages` branch, the advantages
are:

-   It keeps your source and generated website in different branches and
    therefore maintains version control history for both.
-   It uses the default Hugo's **public** folder.

So basically this project's repository has the following branches:

-   **master**: Hosts the source code of the project.
-   **hugo**: Hosts the source code of the documentation website.
-   **gh-pages**: Hosts the static assets generates by hugo.


## Step 1: Initializing branches {#step-1-initializing-branches}

```bash
# Hugo branch
git checkout -b hugo
echo "public" >> .gitignore
git add .gitignore
git commit -m "Adds .gitignore"
hugo new site .
git add .
git commit -m "Adds initial hugo site"
git push origin hugo

# gh-pages branch
git checkout --orphan gh-pages
git reset --hard
git commit --allow-empty -m "Initializing gh-pages branch"
git push origin gh-pages
```


## Step 2: Generate a SSH key. {#step-2-generate-a-ssh-key-dot}

```bash
ssh-keygen -t rsa -f hugo -q -N ""
```

{{% notice note %}}
This will generate the files: `hugo` and `hugo.pub` which will be
needed for the next steps.
{{% /notice %}}


## Step 3: Add a deployment key {#step-3-add-a-deployment-key}

Navigate to your GitHub's repository settings and under **Deploy keys**
and add a new one using the content of the `hugo` SSH private key
generated in the previous step.

{{< figure src="/images/deploy_key.png" >}}

{{% notice warning %}}
Make sure the **Allow write access** is checked, otherwise the GitHub's
Action won't be able to push changes.
{{% /notice %}}


## Step 4: Add the GitHub's Action. {#step-4-add-the-github-s-action-dot}

Create the needed directory in the **hugo** branch:

```bash
git checkout hugo
mkdir -p .github/workflows
```

Add a new file in the path `.github/workflows/gh_pages.yml` with the
following content:

{{< highlight yaml "hl_lines=35" >}}
name: gh_pages

on:
  push:
    branches:
      - hugo

jobs:
  build:

    runs-on: ubuntu-18.04

    steps:
    - uses: actions/checkout@v1
    - name: Installs hugo
      run: |
        cd /tmp
        wget https://github.com/gohugoio/hugo/releases/download/v0.57.0/hugo_0.57.0_Linux-64bit.deb
        sudo dpkg -i hugo_0.57.0_Linux-64bit.deb
        hugo version
    - name: Build hugo site
      run: |
        rm -rf public
        git worktree add -b gh-pages public origin/gh-pages
        hugo

    - name: Configure git and deployment key
      env:
        GITHUB_DEPLOY_KEY: ${{ secrets.GITHUB_DEPLOY_KEY }}
      run: |
        mkdir /home/runner/.ssh
        ssh-keyscan -t rsa github.com >/home/runner/.ssh/known_hosts
        echo "${GITHUB_DEPLOY_KEY}" > /home/runner/.ssh/id_rsa && \
        chmod 400 /home/runner/.ssh/id_rsa
        git remote set-url origin git@github.com:iris-garcia/webhooks-handler.git
        git config --global user.name "GitHub Action"
        git config --global user.email "action@github.com"

    - name: Commit and push changes to gh-pages
      run: |
        cd public
        git add --all
        git commit -m "Publishing to gh-pages branch"
        cd ..
        git push origin gh-pages
{{< /highlight >}}

{{% notice note %}}
Replace the origin's remote with your repository.
{{% /notice %}}

Finally commit and push the changes (which should trigger already the
Action).

```bash
git add .github/workflows/gh_pages.yml
git commit -m "Adds GitHub's Action to build hugo site."
git push origin hugo
```


## Step 5: Verify the Action {#step-5-verify-the-action}

If everything went well you should already have your site updated and a
new commit to the `gh-pages` branch.

You can also see the output of the Action navigating to the **Actions**
section of your repository.

{{< figure src="/images/gh_action.png" >}}
