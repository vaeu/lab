# first-role

This is the first Ansible role that requires one to install `nginx`
binary via HTTP from [some remote
server](https://example.com/registry/nginx_1.0.0), place said binary
into `/bin/` and add it as a `systemd` system service.

## Measures I’ve undertaken to complete this task (all in present tense)

Create master VM and copy my public SSH key onto it:

```sh
$ ssh-copy-id vm@192.168...
```

Disable password authentication for master VM:

```sh
$ ssh vm@192.168...
$ sudo sed -i '/^PasswordAuth/ s/yes/no/' /etc/ssh/sshd_config
$ sudo systemctl restart sshd  # restarting the sshd service
```

1. Create both the controller and the target Ansible VMs by cloning our
   master VM.
1. Generate new MAC addresses for each cloned VM so that new local IPs
   get assigned to these machines.
1. Change hostnames of both VMs for clarity’s sake:

```sh
# connecting to target VM via ssh and then changing its hostname
$ sudo sed -i 's/.*/t/' /etc/hostname # ‘t’ stands for ‘target’
# connecting to controller VM via ssh and then changing its hostname
$ sudo sed -i 's/.*/c/' /etc/hostname # ‘c’ stands for ‘controller’
```

Change hosts of both VMs:

```sh
# connecting to target VM via ssh and then changing its hosts
$ printf '127.0.0.1\tlocalhost t\n::1\t\tlocalhost t\n' | sudo tee /etc/hosts
# connecting to controller VM via ssh and then changing its hosts
$ printf '127.0.0.1\tlocalhost c\n::1\t\tlocalhost c\n' | sudo tee /etc/hosts
```

Reboot both VMs for all changes to take place:

```sh
$ sudo /sbin/shutdown -r now
```

Manipulate both machines easily with the help of an alias:

```sh
cat >> ~/.ssh/config <<EOF
Host controller
  Hostname 192.168...
  User c
Host target
  Hostname 192.168...
  User t
EOF
```

Update repos and install Ansible on our controller VM:

```sh
$ sudo yum -y update
# add EPEL repo that comes with not-so-outdated Ansible package
$ sudo yum -y install epel-release
$ sudo yum -y install ansible
```

Create an inventory file on our controller VM for target one (I’m aware
of Ansible Vault, don’t care for now):

```sh
$ ssh controller
$ mkdir nginx
$ cd !$
$ echo "t ansible_host=192.168... ansible_ssh_pass=crap" > inventory.yml
```

Test whether our target machine can be accessed by Ansible:

```sh
$ ansible t -m ping -i inventory.yml
t | SUCCESS => {
    "ansible_facts": {
        "discovered_interpreter_python": "/usr/bin/python"
    },
    "changed": false,
    "ping": "pong"
}
```
