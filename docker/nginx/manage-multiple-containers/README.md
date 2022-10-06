# manage-multiple-containers

The objective is to manipulate several containers, set some environment
variables by passing those as arguments to `docker container run`
command, peek at logs, etc.

`nginx` will be used as a reverse proxy, `mysql` shall randomise root
password on each run, and `httpd` shall serve as a web server. Each of
the mentioned programs ought to be run in the background.

## Without further ado, let’s get started

Run `nginx` listening on the default port and give it a sane name:

```sh
$ docker container run -d -p 80:80 --name proxy nginx
caac69421c0100b756794279983f6fc9e26742b6b02be0d5785c199d7175e3c3
```

Next, run `mysql` on its default port and pass it an environment
variable that shall randomise root password on each run:

```sh
$ docker container run \
>   -d -p 3306:3306 -e MYSQL_RANDOM_ROOT_PASSWORD=yes --name db mysql
Unable to find image 'mysql:latest' locally
latest: Pulling from library/mysql
051f419db9dd: Pulling fs layer
7627573fa82a: Pulling fs layer
a44b358d7796: Pull complete
95753aff4b95: Pull complete
a1fa3bee53f4: Pull complete
f5227e0d612c: Pull complete
b4b4368b1983: Pull complete
f26212810c32: Pull complete
d803d4215f95: Pull complete
d5358a7f7d07: Pull complete
435e8908cd69: Pull complete
Digest: sha256:b9532b1edea72b6cee12d9f5a78547bd3812ea5db842566e17f8b33291ed2921
Status: Downloaded newer image for mysql:latest
3e4ae3671a2e74c75b8e0bebeccdf333917523536934e464ee6c61ee000d52a3
```

Now, let’s obtain the root password `mysql` has randomly generated for
us:

```sh
$ rootpwd="$(docker container logs db | awk '/GENERATED ROOT PASSWORD/ { print $(NF); exit }')"
$ echo "$rootpwd"
YT4EqhdWG55n6M2qfeekAQOo5fSXlBIv
```

Wonderful! Let’s run `httpd` and get it over with already:

```sh
$ docker container run -d -p 8080:80 --name web httpd
Unable to find image 'httpd:latest' locally
latest: Pulling from library/httpd
bd159e379b3b: Pull complete
36d838c2f6d6: Pull complete
b55eda22bb18: Pull complete
f6e6bfa28393: Pull complete
a1b49b7ecb8a: Pull complete
Digest: sha256:4400fb49c9d7d218d3c8109ef721e0ec1f3897028a3004b098af587d565f4ae5
Status: Downloaded newer image for httpd:latest
5341f21c27c6ad2357dffa3dfb4f4f8591130b8faa4bc1d457fbce9a29720c6f
```

Let’s see if we indeed have three of our containers up and running:

```sh
$ docker container ls
CONTAINER ID   IMAGE     COMMAND                  CREATED              STATUS              PORTS                               NAMES
5341f21c27c6   httpd     "httpd-foreground"       About a minute ago   Up About a minute   0.0.0.0:8080->80/tcp                web
3e4ae3671a2e   mysql     "docker-entrypoint.s…"   3 minutes ago        Up 3 minutes        0.0.0.0:3306->3306/tcp, 33060/tcp   db
caac69421c01   nginx     "/docker-entrypoint.…"   3 minutes ago        Up 3 minutes        0.0.0.0:80->80/tcp                  proxy
```

Cool! Let’s stop all the containers now and remove them for good:

```sh
$ docker container stop $(docker container ls -q)
5341f21c27c6
3e4ae3671a2e
caac69421c01
$ docker container rm $(docker container ls -aq)
5341f21c27c6
3e4ae3671a2e
caac69421c01
```

Hooray! Now that we’ve done all this hard work, we might as well have a
snack or two.
