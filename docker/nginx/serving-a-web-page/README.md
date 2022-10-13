# serving-a-web-page

We are about to serve a simple web page using the latest version of
`nginx` image.

---

Let’s build a local Docker image based on `Dockerfile` file we’ve just
created:

```sh
$ docker build -t myimage .
Sending build context to Docker daemon  4.096kB
Step 1/3 : FROM nginx:latest
 ---> 2d389e545974
Step 2/3 : WORKDIR /usr/share/nginx/html
 ---> Running in 7c5efbb9718a
Removing intermediate container 7c5efbb9718a
 ---> 644d23f75461
Step 3/3 : COPY index.html .
 ---> ad1f8c28b7e2
Successfully built ad1f8c28b7e2
Successfully tagged myimage:latest
```

It is time for us to fire up the container based on our local image and
expose the default ports so we can access the web server:

```sh
$ docker container run --rm -d -p 80:80 myimage
e931e320eaf67cc4415c45451d3bf43650cf6b35501ad3feb501272286d40d35
```

Let’s see if our custom web page is being shown upon querying localhost:

```sh
$ curl -X GET localhost
<!DOCTYPE html>
<html lang="en-GB">
  <head>
    <title>Hooray!</title>
    <meta charset="UTF-8">
  </head>

  <body>
    <main>
      <h1>This shite works!</h1>
    </main>
  </body>
</html>
```

And as always, let’s clean up after ourselves:

```sh
$ docker container stop $(docker container ls -q)
e931e320eaf6
```
