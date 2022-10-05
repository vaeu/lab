# links

Run recent version of wonderful [links](http://links.twibright.com/) web
browser with ease.

## The ins and outs of this project

1. Build latest [links](http://links.twibright.com/) release from
   source.
1. Take advantage of Docker’s multi-stage build feature by copying
   shared libraries along with dynamically linked `links` binary over to
   our brand-new empty image.
1. Use Docker’s ‘scratch’ image to shrink the size of our final image
   even further.
