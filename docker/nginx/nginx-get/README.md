# nginx-get

The goal is to run `nginx` and send a `GET` request to the container via
`curl` to ensure it is running and is listening on port 80.

---

Let’s run the container based on the official image, shall we?

```sh
$ docker container run --rm -p 80:80 nginx
latest: Pulling from library/nginx
31b3f1ad4ce1: Pull complete
fd42b079d0f8: Pull complete
30585fbbebc6: Pull complete
18f4ffdd25f4: Pull complete
9dc932c8fba2: Pull complete
600c24b8ba39: Pull complete
Digest: sha256:0b970013351304af46f322da1263516b188318682b2ab1091862497591189ff1
Status: Downloaded newer image for nginx:latest
```

Next, let’s try sending a `GET` request to our localhost:

```sh
$ curl -X GET localhost
<!DOCTYPE html>
<html>
<head>
<title>Welcome to nginx!</title>
<style>
html { color-scheme: light dark; }
body { width: 35em; margin: 0 auto;
font-family: Tahoma, Verdana, Arial, sans-serif; }
</style>
</head>
<body>
<h1>Welcome to nginx!</h1>
<p>If you see this page, the nginx web server is successfully installed and
working. Further configuration is required.</p>

<p>For online documentation and support please refer to
<a href="http://nginx.org/">nginx.org</a>.<br/>
Commercial support is available at
<a href="http://nginx.com/">nginx.com</a>.</p>

<p><em>Thank you for using nginx.</em></p>
</body>
</html>
```

Great, seems like the server is running and is listening on the default
port! Finally, let’s take a look at logs `nginx` has produced inside the
container:

```sh
$ docker container logs "$(docker container ps | awk '/nginx/ { print $(NF) }')"
/docker-entrypoint.sh: /docker-entrypoint.d/ is not empty, will attempt to perform configuration
/docker-entrypoint.sh: Looking for shell scripts in /docker-entrypoint.d/
/docker-entrypoint.sh: Launching /docker-entrypoint.d/10-listen-on-ipv6-by-default.sh
10-listen-on-ipv6-by-default.sh: info: Getting the checksum of /etc/nginx/conf.d/default.conf
10-listen-on-ipv6-by-default.sh: info: Enabled listen on IPv6 in /etc/nginx/conf.d/default.conf
/docker-entrypoint.sh: Launching /docker-entrypoint.d/20-envsubst-on-templates.sh
/docker-entrypoint.sh: Launching /docker-entrypoint.d/30-tune-worker-processes.sh
/docker-entrypoint.sh: Configuration complete; ready for start up
2022/10/01 14:13:12 [notice] 1#1: using the "epoll" event method
2022/10/01 14:13:12 [notice] 1#1: nginx/1.23.1
2022/10/01 14:13:12 [notice] 1#1: built by gcc 10.2.1 20210110 (Debian 10.2.1-6)
2022/10/01 14:13:12 [notice] 1#1: OS: Linux 5.10.0-18-amd64
2022/10/01 14:13:12 [notice] 1#1: getrlimit(RLIMIT_NOFILE): 1048576:1048576
2022/10/01 14:13:12 [notice] 1#1: start worker processes
2022/10/01 14:13:12 [notice] 1#1: start worker process 32
2022/10/01 14:13:12 [notice] 1#1: start worker process 33
2022/10/01 14:13:12 [notice] 1#1: start worker process 34
2022/10/01 14:13:12 [notice] 1#1: start worker process 35
172.17.0.1 - - [01/Oct/2022:14:13:23 +0000] "GET / HTTP/1.1" 200 615 "-" "curl/7.74.0" "-"
```

Brilliant! Now, let’s take a tea break; we deserved it after all this
hard work.
