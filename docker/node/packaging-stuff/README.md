# packaging-stuff

---

- [DOCKERISING NODE](#dockerising-node)

## DOCKERISING NODE

Having all the source files, we’re going to use the lightweight version
of the official `node` image as well as
[tini](https://github.com/krallin/tini) for starting `node` as its child
process.

```sh
$ docker build -t vaeu/nodetest:latest .
Sending build context to Docker daemon  238.6kB
Step 1/7 : FROM node:18-alpine
 ---> 9f6ca4d52527
Step 2/7 : EXPOSE 3000
 ---> Using cache
 ---> 189229e99766
Step 3/7 : WORKDIR /usr/src/app
 ---> Using cache
 ---> c5311f8e7dcd
Step 4/7 : COPY . .
 ---> Using cache
 ---> b65ff7a909a8
Step 5/7 : RUN apk add --no-cache tini && npm install && npm cache clean --force
 ---> Using cache
 ---> 22dd61500a85
Step 6/7 : ENTRYPOINT ["/sbin/tini"]
 ---> Using cache
 ---> 88a8dc2b4b0b
Step 7/7 : CMD ["--", "node", "./bin/www"]
 ---> Using cache
 ---> 35152bfe9d01
Successfully built 35152bfe9d01
Successfully tagged vaeu/nodetest:latest
```

Let’s see whether we can access the web page using the port we have
exposed in the `Dockerfile` file:

```sh
$ docker run -d --rm --name node vaeu/nodetest
1d09d128a6dc0936baf8350eab185b7b09d9ef50a7b3a7a0a1de27be7959eb76
$
$ docker container inspect -f '{{ .NetworkSettings.IPAddress }}' node
172.17.0.2
$
$ contip="$(docker container inspect -f '{{ .NetworkSettings.IPAddress }}' node)"
$
$ curl -X GET "$contip:3000"
<!DOCTYPE html>
<html>
  <head>
    <title>test</title>
  </head>
  <body>
    <h1>test</h1>
<p>It works.</p>

<img src="/images/varg.gif" />

  </body>
</html>
```

Wonderful! Let’s stop the container, push the image onto Docker Hub and
remove local image from cache:

```sh
$ docker image push vaeu/nodetest:latest
The push refers to repository [docker.io/vaeu/nodetest]
59b136d17ec7: Pushed
13dc8a9baad7: Pushed
157286095852: Pushed
bb2ebe54bcf0: Mounted from library/node
cd2258b7ca68: Mounted from library/node
fbf63a621b59: Mounted from library/node
994393dc58e7: Mounted from library/node
latest: digest: sha256:ee0991151bade0c94809983838625a6a4f9e1b96a7f85d3506bf444cfcce6da6 size: 1786
$
$ docker image rm $(docker image ls -q | head -1)
Untagged: vaeu/nodetest:latest
Untagged: vaeu/nodetest@sha256:ee0991151bade0c94809983838625a6a4f9e1b96a7f85d3506bf444cfcce6da6
Deleted: sha256:35152bfe9d01631daf9ea2431cde429f199b56cb538eead28983de414242dbc3
Deleted: sha256:88a8dc2b4b0bd918a95ea48d07415ebf8809c301aeba1405104e6fe3658240bd
Deleted: sha256:22dd61500a854871a820dd3ad0491cf439a98b08f4432ea86dd9725eea16ef30
Deleted: sha256:2090d44f6fbeb1940b5466f8c7a22a122ab8027ffc0c9bad3d880e66cf013b71
Deleted: sha256:b65ff7a909a85e0ba98e758dc04ce9fd40d7621dcd593fa5b7a44683dff51cb5
Deleted: sha256:4855bb31d711916246429ab451ed7a5f48aa3a4affe2a1b4a3630578eecb9362
Deleted: sha256:c5311f8e7dcd7576b0f2fb3b0675a34479a770eb794070e3086b8c64551e0e9b
Deleted: sha256:61915c1bb812c24f4277ffc453864c2554f83313464928d081f9975e2c17109e
Deleted: sha256:189229e997660ed29b72c49f0572a87df99034f4830fd5cc5c2c3a38a53a0780
```

OK, let’s run the container based on the remote image we have just
uploaded to Docker Hub:

```sh
$ docker container run -d --rm --name node vaeu/nodetest
Unable to find image 'vaeu/nodetest:latest' locally
latest: Pulling from vaeu/nodetest
213ec9aee27d: Already exists
4cf055c45671: Already exists
bb15f8897be6: Already exists
8d429ea46e68: Already exists
e9239b15c908: Pull complete
9a31360f194d: Pull complete
2a63dcc0b7e0: Pull complete
Digest: sha256:ee0991151bade0c94809983838625a6a4f9e1b96a7f85d3506bf444cfcce6da6
Status: Downloaded newer image for vaeu/nodetest:latest
8fd545454de854d78f03fc920ebfa9fba85c51df1689fc8ee13ce922e2b62f83
$
$ contip="$(docker container inspect -f '{{ .NetworkSettings.IPAddress }}' node)"
$
$ curl -X GET "$contip:3000"
<!DOCTYPE html>
<html>
  <head>
    <title>test</title>
  </head>
  <body>
    <h1>test</h1>
<p>It works.</p>

<img src="/images/varg.gif" />

  </body>
</html>
```

Gorgeous!
