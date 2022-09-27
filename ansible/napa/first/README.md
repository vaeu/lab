# first-role

This is the first Ansible role that requires one to install `nginx`
binary via HTTP from [some remote
server](https://example.com/registry/nginx_1.0.0), place said binary
into `/bin/` and add it as a `systemd` system service.

## Measures I’ve undertaken to complete this task (all in present tense)

1. Create master VM and copy my public SSH key onto it:

```sh
$ ssh-copy-id vm@192.168...
```

1. Disable password authentication for master VM:

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

1. Change hosts of both VMs:

```sh
# connecting to target VM via ssh and then changing its hosts
$ printf '127.0.0.1\tlocalhost t\n::1\t\tlocalhost t\n' | sudo tee /etc/hosts
# connecting to controller VM via ssh and then changing its hosts
$ printf '127.0.0.1\tlocalhost c\n::1\t\tlocalhost c\n' | sudo tee /etc/hosts
```

1. Reboot both VMs for all changes to take place:

```sh
$ sudo /sbin/shutdown -r now
```
