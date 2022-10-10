# vagrant

Just playing around with things, nothing to see here.

---

- [FAMILIARISING MYSELF WITH VAGRANT](#familiarising-myself-with-vagrant)
  - [SETTING UP VAGRANT](#setting-up-vagrant)
  - [DEALING WITH ANSIBLE](#dealing-with-ansible)
- [SAILING ON](#sailing-on)
  - [CRAFTING VAGRANTFILE](#crafting-vagrantfile)
    - [INSTALLING VBOX GUEST ADDITIONS](#installing-vbox-guest-additions)

## FAMILIARISING MYSELF WITH VAGRANT

This is the starting point of my journey into great depths of VM
management and infrastructure automation.

### SETTING UP VAGRANT

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
$ . torsocks off
Tor mode deactivated. Command will NOT go through Tor anymore.
$
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

### DEALING WITH ANSIBLE

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

## SAILING ON

The struggle continues.

### CRAFTING VAGRANTFILE

Let’s continue working with the same CentOS box:

```sh
$ vagrant init centos/7
A `Vagrantfile` has been placed in this directory. You are now
ready to `vagrant up` your first virtual environment! Please read
the comments in the Vagrantfile as well as documentation on
`vagrantup.com` for more information on using Vagrant.
```

Now, let’s populate our `Vagrantfile` file with seemingly unsafe
options, but it should be fine since we’re just testing things:

```sh
$ sed -i '/^  *config[.]vm[.]box/a\
> \
>   config.ssh.insert_key = false\
>   config.vm.synced_folder ".", "/vagrant", disabled: true\
>   config.vm.provider :virtualbox do |vb|\
>     vb.memory = 512\
>     vb.linked_clone = true\
>   end\
> \
>   # first VM\
>   config.vm.define "vm1" do |app|\
>     app.vm.hostname = "vm1.test"\
>     app.vm.network :private_network, ip: "192.168.56.163"\
>   end\
> \
>   # second VM\
>   config.vm.define "vm2" do |app|\
>     app.vm.hostname = "vm2.test"\
>     app.vm.network :private_network, ip: "192.168.56.164"\
>   end\
> \
>   # database server\
>   config.vm.define "db" do |db|\
>     db.vm.hostname = "db.test"\
>     db.vm.network :private_network, ip: "192.168.56.165"\
>   end' Vagrantfile
```

The above configuration disables the copying of current working
directory into `/vagrant` inside the VM; it also caps memory usage at
512M and uses the linked clone feature of ‘VirtualBox’ that shares
virtual disks with the main VM, making the machine spin up faster!

Moreover, we specify three dummy servers, one of which acts as some sort
of database server; each VM gets assigned a manual IP address despite
‘Vagrant’ being able to automatically generate those by itself.

OK, let’s rock:

```sh
$ vagrant up
Bringing machine 'vm1' up with 'virtualbox' provider...
Bringing machine 'vm2' up with 'virtualbox' provider...
Bringing machine 'db' up with 'virtualbox' provider...
==> vm1: Checking if box 'centos/7' version '2004.01' is up to date...
==> vm1: Clearing any previously set network interfaces...
==> vm1: Preparing network interfaces based on configuration...
    vm1: Adapter 1: nat
    vm1: Adapter 2: hostonly
==> vm1: Forwarding ports...
    vm1: 22 (guest) => 2222 (host) (adapter 1)
==> vm1: Running 'pre-boot' VM customizations...
==> vm1: Booting VM...
==> vm1: Waiting for machine to boot. This may take a few minutes...
    vm1: SSH address: 127.0.0.1:2222
    vm1: SSH username: vagrant
    vm1: SSH auth method: private key
==> vm1: Machine booted and ready!
==> vm1: Checking for guest additions in VM...
    vm1: No guest additions were detected on the base box for this VM! Guest
    vm1: additions are required for forwarded ports, shared folders, host only
    vm1: networking, and more. If SSH fails on this machine, please install
    vm1: the guest additions and repackage the box to continue.
    vm1:
    vm1: This is not an error message; everything may continue to work properly,
    vm1: in which case you may ignore this message.
==> vm1: Setting hostname...
==> vm1: Configuring and enabling network interfaces...
==> vm2: Cloning VM...
==> vm2: Matching MAC address for NAT networking...
==> vm2: Checking if box 'centos/7' version '2004.01' is up to date...
==> vm2: Setting the name of the VM: vagrant_vm2_1665318834078_99605
==> vm2: Fixed port collision for 22 => 2222. Now on port 2200.
==> vm2: Clearing any previously set network interfaces...
==> vm2: Preparing network interfaces based on configuration...
    vm2: Adapter 1: nat
    vm2: Adapter 2: hostonly
==> vm2: Forwarding ports...
    vm2: 22 (guest) => 2200 (host) (adapter 1)
==> vm2: Running 'pre-boot' VM customizations...
==> vm2: Booting VM...
==> vm2: Waiting for machine to boot. This may take a few minutes...
    vm2: SSH address: 127.0.0.1:2200
    vm2: SSH username: vagrant
    vm2: SSH auth method: private key
==> vm2: Machine booted and ready!
==> vm2: Checking for guest additions in VM...
    vm2: No guest additions were detected on the base box for this VM! Guest
    vm2: additions are required for forwarded ports, shared folders, host only
    vm2: networking, and more. If SSH fails on this machine, please install
    vm2: the guest additions and repackage the box to continue.
    vm2:
    vm2: This is not an error message; everything may continue to work properly,
    vm2: in which case you may ignore this message.
==> vm2: Setting hostname...
==> vm2: Configuring and enabling network interfaces...
==> db: Cloning VM...
==> db: Matching MAC address for NAT networking...
==> db: Checking if box 'centos/7' version '2004.01' is up to date...
==> db: Setting the name of the VM: vagrant_db_1665318860465_97790
==> db: Fixed port collision for 22 => 2222. Now on port 2201.
==> db: Clearing any previously set network interfaces...
==> db: Preparing network interfaces based on configuration...
    db: Adapter 1: nat
    db: Adapter 2: hostonly
==> db: Forwarding ports...
    db: 22 (guest) => 2201 (host) (adapter 1)
==> db: Running 'pre-boot' VM customizations...
==> db: Booting VM...
==> db: Waiting for machine to boot. This may take a few minutes...
    db: SSH address: 127.0.0.1:2201
    db: SSH username: vagrant
    db: SSH auth method: private key
==> db: Machine booted and ready!
==> db: Checking for guest additions in VM...
    db: No guest additions were detected on the base box for this VM! Guest
    db: additions are required for forwarded ports, shared folders, host only
    db: networking, and more. If SSH fails on this machine, please install
    db: the guest additions and repackage the box to continue.
    db:
    db: This is not an error message; everything may continue to work properly,
    db: in which case you may ignore this message.
==> db: Setting hostname...
==> db: Configuring and enabling network interfaces...
```

#### INSTALLING VBOX GUEST ADDITIONS

Seems like we don’t have VBox Guest Additions installed. Let’s fix it
right away:

```sh
$ . torsocks on
Tor mode activated. Every command will be torified for this shell.
$ vagrant plugin install vagrant-vbguest
Installing the 'vagrant-vbguest' plugin. This can take a few minutes...
Fetching micromachine-3.0.0.gem
Fetching vagrant-vbguest-0.30.0.gem
Installed the plugin 'vagrant-vbguest (0.30.0)'!
$ . torsocks off
Tor mode deactivated. Command will NOT go through Tor anymore.
$
$ vagrant destroy
    db: Are you sure you want to destroy the 'db' VM? [y/N] y
==> db: Forcing shutdown of VM...
==> db: Destroying VM and associated drives...
    vm2: Are you sure you want to destroy the 'vm2' VM? [y/N] y
==> vm2: Forcing shutdown of VM...
==> vm2: Destroying VM and associated drives...
    vm1: Are you sure you want to destroy the 'vm1' VM? [y/N] y
==> vm1: Forcing shutdown of VM...
==> vm1: Destroying VM and associated drives...
$
$ vagrant up
<...>
No package kernel-devel-3.10.0-1127.el7.x86_64 available.
Error: Nothing to do
Unmounting Virtualbox Guest Additions ISO from: /mnt
umount: /mnt: not mounted
```

A little bit of searching, and [this
discussion](https://github.com/dotless-de/vagrant-vbguest/discussions/401)
pops up. Let’s try to make changes to our `Vagrantfile` file and see if
the proposed solution solves anything:

```sh
$ sed -i '/^  *app[.]vm[.]network/a\
>     app.vbguest.installer_hooks[:before_install] = ["yum install -y epel-release", "sleep 1"]\
>     app.vbguest.installer_options = { allow_kernel_upgrade: false, enablerepo: true }' Vagrantfile
$
$ sed -i '/^  *db[.]vm[.]network/a\
>     db.vbguest.installer_hooks[:before_install] = ["yum install -y epel-release", "sleep 1"]\
>     db.vbguest.installer_options = { allow_kernel_upgrade: false, enablerepo: true }' Vagrantfile
$
$ vagrant up
<...>
==> db: Machine booted and ready!
[db] No Virtualbox Guest Additions installation found.
==> db: Executing pre-install hooks
<...>
Complete!
Copy iso file /usr/share/virtualbox/VBoxGuestAdditions.iso into the box /tmp/VBoxGuestAdditions.iso
Mounting Virtualbox Guest Additions ISO to: /mnt
mount: /dev/loop0 is write-protected, mounting read-only
Installing Virtualbox Guest Additions 6.1.38 - guest version is unknown
Verifying archive integrity... All good.
Uncompressing VirtualBox 6.1.38 Guest Additions for Linux........
VirtualBox Guest Additions installer
Copying additional installer modules ...
Installing additional modules ...
/opt/VBoxGuestAdditions-6.1.38/bin/VBoxClient: error while loading shared libraries: libX11.so.6: cannot open shared object file: No such file or directory
/opt/VBoxGuestAdditions-6.1.38/bin/VBoxClient: error while loading shared libraries: libX11.so.6: cannot open shared object file: No such file or directory
VirtualBox Guest Additions: Starting.
VirtualBox Guest Additions: Building the VirtualBox Guest Additions kernel
modules.  This may take a while.
VirtualBox Guest Additions: To build modules for other installed kernels, run
VirtualBox Guest Additions:   /sbin/rcvboxadd quicksetup <version>
VirtualBox Guest Additions: or
VirtualBox Guest Additions:   /sbin/rcvboxadd quicksetup all
VirtualBox Guest Additions: Building the modules for kernel
3.10.0-1127.el7.x86_64.
Redirecting to /bin/systemctl start vboxadd.service
Redirecting to /bin/systemctl start vboxadd-service.service
Unmounting Virtualbox Guest Additions ISO from: /mnt
==> db: Checking for guest additions in VM...
==> db: Setting hostname...
==> db: Configuring and enabling network interfaces...
```

Everything went according to plan, or so it seems.
