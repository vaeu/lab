# vagrant

Just playing around with things, nothing to see here.

---

- [SETTING UP VAGRANT](#setting-up-vagrant)
- [DEALING WITH ANSIBLE](#dealing-with-ansible)

## SETTING UP VAGRANT

Let’s fetch CentOS 7 box and set it up:

```sh
$ vagrant init centos/7
A `Vagrantfile` has been placed in this directory. You are now
ready to `vagrant up` your first virtual environment! Please read
the comments in the Vagrantfile as well as documentation on
`vagrantup.com` for more information on using Vagrant.
$
$ vagrant up
Bringing machine 'default' up with 'virtualbox' provider...
==> default: Box 'centos/7' could not be found. Attempting to find and install...
    default: Box Provider: virtualbox
    default: Box Version: >= 0
The box 'centos/7' could not be found or
could not be accessed in the remote catalog. If this is a private
box on HashiCorp's Vagrant Cloud, please verify you're logged in via
`vagrant login`. Also, please double-check the name. The expanded
URL and error message are shown below:

URL: ["https://vagrantcloud.com/centos/7"]
Error: The requested URL returned error: 404
```

Turns out ‘Vagrant Cloud’ refuses to establish any sort of connection
with my IP address.

Let’s install `vagrant-proxyconf` plugin to work around this issue:

```sh
$ vagrant plugin install vagrant-proxyconf
Installing the 'vagrant-proxyconf' plugin. This can take a few minutes...
Vagrant failed to load a configured plugin source. This can be caused
by a variety of issues including: transient connectivity issues, proxy
filtering rejecting access to a configured plugin source, or a configured
plugin source not responding correctly. Please review the error message
below to help resolve the issue:

  bad response Not Found 404 (https://gems.hashicorp.com/specs.4.8.gz)

Source: https://gems.hashicorp.com/
```

Interesting… I tried accessing several websites that belong to
‘HashiCorp’, and none of them is available in my region!

Fear not, good old Tor is there to the rescue! Let’s torify the entire
session and see if we can access any server that belongs to ‘HashiCorp’:

```sh
$ . torsocks on
Tor mode activated. Every command will be torified for this shell.
```

Let’s try fetching the box again:

```sh
$ vagrant up
Bringing machine 'default' up with 'virtualbox' provider...
==> default: Box 'centos/7' could not be found. Attempting to find and install...
    default: Box Provider: virtualbox
    default: Box Version: >= 0
==> default: Loading metadata for box 'centos/7'
    default: URL: https://vagrantcloud.com/centos/7
==> default: Adding box 'centos/7' (v2004.01) for provider: virtualbox
    default: Downloading: https://vagrantcloud.com/centos/boxes/7/versions/2004.01/providers/virtualbox.box
Download redirected to host: cloud.centos.org
Progress: 1% (Rate: 745k/s, Estimated time remaining: 0:18:05)
```

Magnificent! Let’s wait for this crap to be downloaded at the speed of a
snail, shall we?

Three hours have passed, and the following error has popped up:

```sh
    default: Calculating and comparing box checksum...
==> default: Successfully added box 'centos/7' (v2004.01) for 'virtualbox'!
==> default: Importing base box 'centos/7'...
==> default: Matching MAC address for NAT networking...
There was an error while executing `VBoxManage`, a CLI used by Vagrant
for controlling VirtualBox. The command and stderr is shown below.

Command: ["list", "hostonlyifs"]

Stderr: VBoxManage: error: Code NS_ERROR_FAILURE (0x80004005) - Operation failed (extended info not available)
VBoxManage: error: Context: "FindHostNetworkInterfacesOfType(HostNetworkInterfaceType_HostOnly, ComSafeArrayAsOutParam(hostNetworkInterfaces))" at line 137 of file VBoxManageList.cpp

1665152152 WARNING torsocks[212763]: [syscall] Unsupported syscall number 315. Denying the call (in tsocks_syscall() at syscall.c:604)
```

Seems like we have to stop routing our traffic through Tor for this
shell session and rerun the same command once again:

```sh
$ exit
$ vagrant up
Bringing machine 'default' up with 'virtualbox' provider...
==> default: Checking if box 'centos/7' version '2004.01' is up to date...
==> default: There was a problem while downloading the metadata for your box
==> default: to check for updates. This is not an error, since it is usually due
==> default: to temporary network problems. This is just a warning. The problem
==> default: encountered was:
==> default:
==> default: The requested URL returned error: 404
==> default:
==> default: If you want to check for box updates, verify your network connection
==> default: is valid and try again.
==> default: Setting the name of the VM: vagrant_default_1665152286975_1895
==> default: Clearing any previously set network interfaces...
==> default: Preparing network interfaces based on configuration...
    default: Adapter 1: nat
==> default: Forwarding ports...
    default: 22 (guest) => 2222 (host) (adapter 1)
==> default: Booting VM...
==> default: Waiting for machine to boot. This may take a few minutes...
    default: SSH address: 127.0.0.1:2222
    default: SSH username: vagrant
    default: SSH auth method: private key
    default:
    default: Vagrant insecure key detected. Vagrant will automatically replace
    default: this with a newly generated keypair for better security.
    default:
    default: Inserting generated public key within guest...
    default: Removing insecure key from the guest if it's present...
    default: Key inserted! Disconnecting and reconnecting using new SSH key...
==> default: Machine booted and ready!
==> default: Checking for guest additions in VM...
    default: No guest additions were detected on the base box for this VM! Guest
    default: additions are required for forwarded ports, shared folders, host only
    default: networking, and more. If SSH fails on this machine, please install
    default: the guest additions and repackage the box to continue.
    default:
    default: This is not an error message; everything may continue to work properly,
    default: in which case you may ignore this message.
==> default: Configuring proxy environment variables...
==> default: Configuring proxy for Yum...
```

Now, let’s set Ansible as the default VM provisioner in our
`Vagrantfile` file and specify which Ansible playbook file shall be used
for further manipulations:

```sh
$ sed -i '/# config.vm.provision/a\
>  config.vm.provision "ansible" do |ansible|\
>    ansible.playbook = "playbook.yml"\
>  end' Vagrantfile
```

## DEALING WITH ANSIBLE

Wonderful. Let’s create a simple Ansible playbook now:

```sh
$ cat <<EOF > playbook.yml
---
- hosts: all
  become: yes
  tasks:
    - name: make sure network time protocol is installed
      ansible.builtin.yum: name=ntp state=present
    - name: ensure network time protocol is up and running
      ansible.builtin.service: name=ntpd enabled=yes state=started
EOF
```

We can go ahead and try to make Vagrant run our playbook against the VM:

```sh
$ vagrant provision
==> default: Configuring proxy environment variables...
==> default: Configuring proxy for Yum...
==> default: Running provisioner: ansible...
    default: Running ansible-playbook...

PLAY [all] *********************************************************************

TASK [Gathering Facts] *********************************************************
ok: [default]

TASK [make sure network time protocol is installed] ****************************
changed: [default]

TASK [ensure network time protocol is up and running] **************************
changed: [default]

PLAY RECAP *********************************************************************
default                    : ok=3    changed=2    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0
```

Excellent! Now is a good time to destroy the VM:

```
$ vagrant destroy
```
