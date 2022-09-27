# first-role

This is the first Ansible role that requires one to install `nginx`
binary via HTTP from [some remote
server](https://example.com/registry/nginx_1.0.0), place said binary
into `/bin/` and add it as a `systemd` system service.

## Measures I’ve undertaken to complete this task (all in present tense)

1. Create both the controller and the target Ansible VMs just for fun.
1. Copy my public SSH key onto the machines to avoid entering passwords
   all the time:

```sh
for h in 192.168.1.{32,242}; do ssh-copy-id vm@${h}; done
```

1. Disable password authentication for both machines:

```sh
# connecting to each machine and changing the value
$ sudo sed -i '/^PasswordAuth/ s/yes/no/' /etc/ssh/sshd_config
# restarting the sshd service
$ sudo systemctl restart sshd
```

1. Change hostnames of both VMs for clarity’s sake:

```sh
# connecting to target VM via ssh
$ sudo sed -i 's/.*/t/' /etc/hostname # ‘t’ stands for ‘target’
# connecting to controller VM via ssh
$ sudo sed -i 's/.*/c/' /etc/hostname # ‘c’ stands for ‘controller’
```

1. Change hosts of both VMs:

```sh
# connecting to target VM via ssh
$ printf '127.0.0.1\tt\n::1\t\tt\n' | sudo tee /etc/hosts
# connecting to controller VM via ssh
$ printf '127.0.0.1\tc\n::1\t\tc\n' | sudo tee /etc/hosts
```
