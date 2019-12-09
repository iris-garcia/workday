+++
title = "Task manager configuration"
author = ["Iris Garcia"]
lastmod = 2019-12-09T18:37:05+01:00
draft = false
weight = 4
asciinema = true
+++

Two new tasks has been added to the task manager:


## Create VM and Provision {#create-vm-and-provision}

```go
// CreateVM creates a vagrant box already provisioned
func CreateVM() error {
        cwd, err := os.Getwd()
        if err != nil {
                fmt.Printf("Error: ", err.Error())
        }
        cmd := exec.Command("packer", "build", "-var", "'version=10.2.0'", "debian10.json")
        cmd.Dir = cwd + "/.packer"
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr
        return cmd.Run()
}
```

A wrapper which creates a VirtualBox VM and then provisions it with
Ansible.


## Provision a host {#provision-a-host}

```go
// ProvisionVM runs an Ansible playbook to provision the configured host.
func ProvisionVM() error {
        cwd, err := os.Getwd()
        if err != nil {
                fmt.Printf("Error: ", err.Error())
        }
        cmd := exec.Command("ansible-playbook", "-i", "env/packer", "workday.yml")
        cmd.Dir = cwd + "/.packer/ansible"
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr
        return cmd.Run()
}
```

This action runs `ansible-playbook` to provision a host previously
configured in the inventory.
