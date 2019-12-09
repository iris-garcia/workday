+++
title = "Ansible"
author = ["Iris Garcia"]
lastmod = 2019-12-09T18:37:05+01:00
draft = false
weight = 2
asciinema = true
+++

## Structure {#structure}

{{< figure src="/images/ansible.png" caption="Figure 1: directory structure" >}}


## Playbook {#playbook}

The following playbook has been created:

```yaml
- name: Base setup
  hosts: api
  gather_facts: no
  become: yes
  tags:
    - base

  vars:
    USER: "workday"
    PACKAGES:
      - git
      - wget

  pre_tasks:
    - name: Install system packages
      apt:
        name: "{{ PACKAGES }}"
        update_cache: yes

  roles:
    - role: user

- name: Setup Golang
  hosts: api
  gather_facts: no
  become: yes
  tags:
    - go

  vars:
    GO_VERSION: "1.13.3"

  roles:
    - role: go

- name: Setup build tool Mage
  hosts: api
  gather_facts: no
  become: yes
  tags:
    - buildtool

  vars:
    USER: "workday"

  roles:
    - role: mage

- name: Setup process manager pm2
  hosts: api
  gather_facts: no
  become: yes
  tags:
    - pm2

  roles:
    - role: pm2
```

<div class="src-block-caption">
  <span class="src-block-number">Code Snippet 1</span>:
  .packer/ansible/workday.yml
</div>

It has four main tasks, each of them run a role and has a `tag`
defined to allow running a one or more tasks instead of the whole
playbook.

{{% notice tip %}}
**The use of `become: yes` is needed because the playbook is ran using
the `vagrant` user, therefore privileged rights are needed.**
{{% /notice %}}


## Roles {#roles}

1.  `Base setup`: This role installs some base packages needed and
    creates an APP user (**workday**) using a variable, this allows to
    easily change the user name without having to change the role.

    ```yaml
    ​- name: Create App user
      user:
        name: "{{ USER }}"
        shell: "/bin/bash"
        create_home: True
    ```

    <div class="src-block-caption">
      <span class="src-block-number">Code Snippet 2</span>:
      .packer/ansible/roles/user/tasks/main.yml
    </div>

2.  `Setup Golang`: This role downloads and installs a specific Go
    version given as a variable, it also creates a symlink of the Go
    binary to **/usr/local/bin/go**.

    ```yaml
    ​- name: Check if go is already installed
      stat:
        path: "/usr/local/bin/go"
      register: go_binary

    - name: Download Go archive
      get_url:
        url: "https://dl.google.com/go/go{{ GO_VERSION }}.linux-amd64.tar.gz"
        dest: "/tmp/go{{ GO_VERSION }}.linux-amd.tar.gz"
      when: not go_binary.stat.exists

    - name: Unarchive Go archive
      unarchive:
        src: "/tmp/go{{ GO_VERSION }}.linux-amd.tar.gz"
        dest: "/usr/local"
        remote_src: yes
      when: not go_binary.stat.exists

    - name: Symlink Go binary
      file:
        src: "/usr/local/go/bin/go"
        dest: "/usr/local/bin/go"
        state: link
      when: not go_binary.stat.exists
    ```

    <div class="src-block-caption">
      <span class="src-block-number">Code Snippet 3</span>:
      .packer/ansible/roles/go/tasks/main.yml
    </div>

3.  `Setup Mage`: This is the build tool used in this project, this
    role installs Mage in the _GOROOT_ path of the user given as a
    variable.

    ```yaml
    ​- name: Create GOROOT for user
      file:
        path: "/home/{{ USER }}/go/src"
        state: directory
        owner: "{{ USER }}"
        group: "{{ USER }}"

    - name: Check if mage is installed
      stat:
        path: "/home/{{ USER }}/go/bin/mage"
      register: mage

    - name: Install mage
      environment:
        GOROOT: "/home/{{ USER }}/go"
      shell: "sudo -u {{ USER }} go get github.com/magefile/mage"
      when: not mage.stat.exists
    ```

    <div class="src-block-caption">
      <span class="src-block-number">Code Snippet 4</span>:
      .packer/ansible/roles/mage/tasks/main.yml
    </div>

4.  `Setup Process Manager`: A role to install npm and pm2, which will
    be used once we deploy the project.

    ```yaml
    ​- name: Install npm
      apt:
        name: npm

    - name: Check if pm2 is installed
      stat:
        path: "/usr/local/bin/pm2"
      register: pm2

    - name: Install pm2
      shell: "npm install -g pm2"
      when: not pm2.stat.exists
    ```

    <div class="src-block-caption">
      <span class="src-block-number">Code Snippet 5</span>:
      .packer/ansible/roles/pm2/tasks/main.yml
    </div>

{{% notice tip %}}
**Every role has checks to make sure the tasks are idempotent, this
means we can run the playbook as many times as we want and we will get the same result.**
{{% /notice %}}


## Inventory {#inventory}

The inventory created is for a `packer` environment, that's why the
ansible host points to localhost.

```yaml
---
all:
  hosts:
    api:
      ansible_host: "127.0.0.1"
```

<div class="src-block-caption">
  <span class="src-block-number">Code Snippet 6</span>:
  .packer/ansible/env/packer/hosts.yml
</div>
