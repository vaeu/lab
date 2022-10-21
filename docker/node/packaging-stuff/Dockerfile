FROM node:18-alpine
EXPOSE 3000
WORKDIR /usr/src/app
COPY . .
RUN apk add --no-cache tini && npm install && npm cache clean --force
ENTRYPOINT ["/sbin/tini"]
CMD ["--", "node", "./bin/www"]
