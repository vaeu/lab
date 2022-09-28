# napa-ansible

The objective is to create several Ansible roles and go through the
hurdle of setting things up.

## Output

<details>
  <summary><i>information emitted by ‘ansible-playbook’ command</i></summary>
  <pre><code>$ ansible-playbook -i hosts configure-webserver
PLAY [Configure nginx and manage partitions] ***********************************

TASK [Gathering Facts] *********************************************************
fatal: [ngx-webserver]: UNREACHABLE! => {"changed": false, "msg": "Failed to connect to the host via ssh: ssh: Could not resolve hostname ngx-webserver: Name or service not known", "unreachable": true}

PLAY RECAP *********************************************************************
ngx-webserver              : ok=0    changed=0    unreachable=1    failed=0    skipped=0    rescued=0    ignored=0</code></pre>
</details>
