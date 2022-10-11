# networking

Exploring networking in Docker.

---

- [CREATING OUR OWN NETWORK](#creating-our-own-network)
  - [MANAGING DNS](#managing-dns)
    - [GETTING ROUND-ROBIN DNS TECHNIQUE TO WORK](#getting-round-robin-dns-technique-to-work)

## CREATING OUR OWN NETWORK

For starters, let’s see what networks are available to us and then
create our own:

```sh
$ docker network ls
NETWORK ID     NAME      DRIVER    SCOPE
76b66aa4d21b   bridge    bridge    local
66129ed1bceb   host      host      local
9156199efe6c   none      null      local
$
$ docker network create nginx_net
8f5d255f70f4349f8780a9f64959198b77fd011cdc868cee857bdd3cf9d1fe7a
$
$ docker network ls
NETWORK ID     NAME        DRIVER    SCOPE
76b66aa4d21b   bridge      bridge    local
66129ed1bceb   host        host      local
8f5d255f70f4   nginx_net   bridge    local
9156199efe6c   none        null      local
```

Let’s run `nginx` in the background inside our newly baked `nginx_net`
virtual network:

```sh
$ docker container run --rm -d --name web --network nginx_net nginx:alpine
Unable to find image 'nginx:alpine' locally
alpine: Pulling from library/nginx
213ec9aee27d: Already exists
6779501a69ba: Pull complete
f294ffcdfaa8: Pull complete
56424afbb509: Pull complete
9a1e8d85723a: Pull complete
5056d2fafbf2: Pull complete
Digest: sha256:b87c350e6c69e0dc7069093dcda226c4430f3836682af4f649f2af9e9b5f1c74
Status: Downloaded newer image for nginx:alpine
e060924d3f151ae0c90806e24e73dad2da060498a8189246de0e5dfd2b981f07
```

Wonderful! It’s time to inspect `nginx_net` network and see if it indeed
does have our new container placed in its relevant group:

```sh
$ docker network inspect nginx_net -f '{{ .Containers }}'
map[e060924d3f151ae0c90806e24e73dad2da060498a8189246de0e5dfd2b981f07:{web 43f8f8baccbff6f9af4e6821763e1a73c2b1697ee8b84098127328b04d9b6925 02:42:ac:12:00:02 172.18.0.2/16 }]
```

OK, it is in there, but let’s pipe the `docker` command through `sed`
for clarity’s sake in lieu of using Go formatting:

```sh
$ docker network inspect nginx_net | sed -n '/Containers/,/},$/p'
        "Containers": {
            "e060924d3f151ae0c90806e24e73dad2da060498a8189246de0e5dfd2b981f07": {
                "Name": "web",
                "EndpointID": "43f8f8baccbff6f9af4e6821763e1a73c2b1697ee8b84098127328b04d9b6925",
                "MacAddress": "02:42:ac:12:00:02",
                "IPv4Address": "172.18.0.2/16",
                "IPv6Address": ""
            }
        },
```

### MANAGING DNS

Let’s run a second `nginx` container akin to the first one, appending
‘2’ to its name:

```sh
$ docker container run --rm -d --name web2 --network nginx_net nginx:alpine
178f4867b5b072866a78fa6636591124e7f38989024253592fd0a5a1dc237238
$
$ docker network inspect nginx_net | sed -n '/Containers/,/^\s\{,8\}},$/p'
        "Containers": {
            "178f4867b5b072866a78fa6636591124e7f38989024253592fd0a5a1dc237238": {
                "Name": "web2",
                "EndpointID": "7cfdd4dcc270b0c2ab05aa69892c6a4679eccc43b07013d64523e1cb80703f2c",
                "MacAddress": "02:42:ac:12:00:03",
                "IPv4Address": "172.18.0.3/16",
                "IPv6Address": ""
            },
            "e060924d3f151ae0c90806e24e73dad2da060498a8189246de0e5dfd2b981f07": {
                "Name": "web",
                "EndpointID": "43f8f8baccbff6f9af4e6821763e1a73c2b1697ee8b84098127328b04d9b6925",
                "MacAddress": "02:42:ac:12:00:02",
                "IPv4Address": "172.18.0.2/16",
                "IPv6Address": ""
            }
        },
```

Good, two containers are up and running and belong to `nginx_net`
network, but can they really communicate with each other? Let’s find
out:

```sh
$ docker container exec -it web2 ping -c 3 web
PING web (172.18.0.2): 56 data bytes
64 bytes from 172.18.0.2: seq=0 ttl=64 time=0.087 ms
64 bytes from 172.18.0.2: seq=1 ttl=64 time=0.100 ms
64 bytes from 172.18.0.2: seq=2 ttl=64 time=0.107 ms

--- web ping statistics ---
3 packets transmitted, 3 packets received, 0% packet loss
round-trip min/avg/max = 0.087/0.098/0.107 ms
```

Spectacular! Thanks to Docker’s DNS magic, two containers can talk to
each other on the same virtual network without us intervening!

Let’s stop all the containers and remove our `nginx_net` network:

```sh
$ docker container stop $(docker container ls -q)
178f4867b5b0
e060924d3f15
$
$ docker network rm nginx_net
nginx_net
```

#### GETTING ROUND-ROBIN DNS TECHNIQUE TO WORK

First and foremost, let’s create a brand-new virtual network:

```sh
$ docker network create rrdns
cf3a6b7440edd1c6162eae819a61a0d0c1095736f8a6f6cff64921bc7789a2e4
```

Now, let’s fire up *two* containers based on ‘Elasticsearch’ image:

```sh
$ docker container run --rm -d --network-alias search --net rrdns elasticsearch:2
Unable to find image 'elasticsearch:2' locally
2: Pulling from library/elasticsearch
05d1a5232b46: Pull complete
5cee356eda6b: Pull complete
89d3385f0fd3: Pull complete
65dd87f6620b: Pull complete
78a183a01190: Pull complete
1a4499c85f97: Pull complete
2c9d39b4bfc1: Pull complete
1b1cec2222c9: Pull complete
59ff4ce9df68: Pull complete
1976bc3ee432: Pull complete
a27899b7a5b5: Pull complete
b0fc7d2c927a: Pull complete
6d94b96bbcd0: Pull complete
6f5bf40725fd: Pull complete
2bf2a528ae9a: Pull complete
Digest: sha256:41ed3a1a16b63de740767944d5405843db00e55058626c22838f23b413aa4a39
Status: Downloaded newer image for elasticsearch:2
8024aee739f0e680af38cf17728e949898d198c8fbc3f93da5aee1451dacf224
$
$ docker container run --rm -d --network-alias search --net rrdns elasticsearch:2
421abcdff8fbf21be9d1d9652b6c084ffc1aea89c6b4e25423d7ede28ba469a6
```

Let’s see if two of our containers resolve to the same name:

```sh
$ docker container run --rm --net rrdns alpine nslookup search
Unable to find image 'alpine:latest' locally
latest: Pulling from library/alpine
Digest: sha256:bc41182d7ef5ffc53a40b044e725193bc10142a1243f395ee852a8d9730fc2ad
Status: Downloaded newer image for alpine:latest
Server:         127.0.0.11
Address:        127.0.0.11:53

Non-authoritative answer:

Non-authoritative answer:
Name:   search
Address: 172.19.0.2
Name:   search
Address: 172.19.0.3
```

As our final step, let’s ensure ‘Elasticsearch’ gives the same response
with randomised name and cluster UUID upon each new request:

```sh
$ docker container run --rm --net rrdns centos:7 curl -s search:9200
Unable to find image 'centos:7' locally
7: Pulling from library/centos
2d473b07cdd5: Pull complete
Digest: sha256:c73f515d06b0fa07bb18d8202035e739a494ce760aa73129f60f4bf2bd22b407
Status: Downloaded newer image for centos:7
{
  "name" : "Sluk",
  "cluster_name" : "elasticsearch",
  "cluster_uuid" : "MT8ZXuVUTdyKGx5p8fn0gQ",
  "version" : {
    "number" : "2.4.6",
    "build_hash" : "5376dca9f70f3abef96a77f4bb22720ace8240fd",
    "build_timestamp" : "2017-07-18T12:17:44Z",
    "build_snapshot" : false,
    "lucene_version" : "5.5.4"
  },
  "tagline" : "You Know, for Search"
}
$
$ docker container run --rm --net rrdns centos:7 curl -s search:9200
{
  "name" : "Carolyn Trainer",
  "cluster_name" : "elasticsearch",
  "cluster_uuid" : "RJi6h51JSsOfG0HpFKhfeQ",
  "version" : {
    "number" : "2.4.6",
    "build_hash" : "5376dca9f70f3abef96a77f4bb22720ace8240fd",
    "build_timestamp" : "2017-07-18T12:17:44Z",
    "build_snapshot" : false,
    "lucene_version" : "5.5.4"
  },
  "tagline" : "You Know, for Search"
}
```

This is it, everything works as expected!
