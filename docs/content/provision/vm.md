+++
title = "Creation of VM"
author = ["Iris Garcia"]
lastmod = 2019-12-09T18:37:05+01:00
draft = false
weight = 3
asciinema = true
+++

The creation of the VM has been done using [packer](https://www.packer.io/) and adapting a
packer's template from this [project](https://github.com/geerlingguy/packer-debian-10).

```json
{
    "variables": {
        "version": ""
    },
    "provisioners": [
        {
            "type": "shell",
            "execute_command": "echo 'vagrant' | {{.Vars}} sudo -S -E bash '{{.Path}}'",
            "script": "scripts/setup.sh"
        },
        {
            "type": "shell",
            "execute_command": "echo 'vagrant' | {{.Vars}} sudo -S -E bash '{{.Path}}'",
            "script": "scripts/ansible.sh"
        },
        {
            "type": "ansible-local",
            "playbook_file": "ansible/workday.yml",
            "inventory_file": "ansible/env/packer/hosts.yml",
            "role_paths": [
                "ansible/roles/go",
                "ansible/roles/mage",
                "ansible/roles/pm2",
                "ansible/roles/user"
            ]
        }
    ],
    "builders": [
        {
            "type": "virtualbox-iso",
            "boot_command": [
                "<esc><wait>",
                "install <wait>",
                " preseed/url=http://{{ .HTTPIP }}:{{ .HTTPPort }}/preseed.cfg <wait>",
                "debian-installer=en_US.UTF-8 <wait>",
                "auto <wait>",
                "locale=en_US.UTF-8 <wait>",
                "kbd-chooser/method=us <wait>",
                "keyboard-configuration/xkb-keymap=us <wait>",
                "netcfg/get_hostname={{ .Name }} <wait>",
                "netcfg/get_domain=vagrantup.com <wait>",
                "fb=false <wait>",
                "debconf/frontend=noninteractive <wait>",
                "console-setup/ask_detect=false <wait>",
                "console-keymaps-at/keymap=us <wait>",
                "grub-installer/bootdev=/dev/sda <wait>",
                "<enter><wait>"
            ],
            "boot_wait": "5s",
            "disk_size": 81920,
            "guest_os_type": "Debian_64",
            "headless": true,
            "http_directory": "http",
            "iso_urls": [
                "iso/debian-10.2.0-amd64-netinst.iso",
                "https://cdimage.debian.org/debian-cd/current/amd64/iso-cd/debian-10.2.0-amd64-netinst.iso"
            ],
            "iso_checksum_type": "sha256",
            "iso_checksum": "e43fef979352df15056ac512ad96a07b515cb8789bf0bfd86f99ed0404f885f5",
            "ssh_username": "vagrant",
            "ssh_password": "vagrant",
            "ssh_port": 22,
            "ssh_wait_timeout": "10000s",
            "shutdown_command": "echo 'vagrant'|sudo -S shutdown -P now",
            "guest_additions_path": "VBoxGuestAdditions_{{.Version}}.iso",
            "virtualbox_version_file": ".vbox_version",
            "vm_name": "packer-debian-10-amd64",
            "vboxmanage": [
                [
                    "modifyvm",
                    "{{.Name}}",
                    "--memory",
                    "2048"
                ],
                [
                    "modifyvm",
                    "{{.Name}}",
                    "--cpus",
                    "2"
                ]
            ]
        }
    ],
    "post-processors": [
        [
            {
                "output": "builds/{{.Provider}}-workday.box",
                "type": "vagrant"
            }
        ]
    ]
}
```

<div class="src-block-caption">
  <span class="src-block-number">Code Snippet 1</span>:
  Packer template
</div>

A new provisioner has been added to run our Ansible playbook.

Running the command `mage createvm`
starts the build of the VM and does the following:

-   Create a VirtualBox VM
-   Download the Debian 10.20 ISO
-   Boot the VM and install the Operating System
-   Wait until the SSH service is running
-   SSH into the machine and run the provisioners.
-   Export the machine as type `vagrant` box.


## Demo {#demo}

{{< asciinema key="packer" rows="50" preload="1"
idle-time-limit="1" >}}
