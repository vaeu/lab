# vagrant

Just playing around with things, nothing to see here.

---

- [SETTING UP VAGRANT](#setting-up-vagrant)

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
session and see if we can access any server that belongs to ‘Hashicorp’:

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
