FROM gcc:12 AS builder
ARG APP=links
ARG VER=2.27
ARG BASE=$APP-$VER
ARG TAR=$BASE.tar.gz
ARG URL=http://$APP.twibright.com/download/$TAR
ARG CONFOPTS=--with-ssl=/usr/lib/x86_64-linux-gnu
RUN curl -LO $URL && tar xzf $TAR && cd $BASE && \
    ./configure $CONFOPTS && make && ldconfig && strip $APP && \
    mkdir /app && mv $APP /app/

FROM scratch
ENV PATH="/app:$PATH"
ENTRYPOINT ["links"]
CMD ["--help"]
COPY --from=builder /app /app
COPY --from=builder /lib64/ld-linux-x86-64.so.2 /lib64/
COPY --from=builder /lib/x86_64-linux-gnu/libbz2.so.1.0 \
  /lib/x86_64-linux-gnu/libc.so.6 \
  /lib/x86_64-linux-gnu/libdl.so.2 \
  /lib/x86_64-linux-gnu/liblzma.so.5 \
  /lib/x86_64-linux-gnu/libpthread.so.0 \
  /lib/x86_64-linux-gnu/libz.so.1 /lib/x86_64-linux-gnu/
COPY --from=builder /usr/lib/x86_64-linux-gnu/libbrotlicommon.so.1 \
  /usr/lib/x86_64-linux-gnu/libbrotlidec.so.1 \
  /usr/lib/x86_64-linux-gnu/libcrypto.so.1.1 \
  /usr/lib/x86_64-linux-gnu/libevent-2.1.so.7 \
  /usr/lib/x86_64-linux-gnu/libssl.so.1.1 /usr/lib/x86_64-linux-gnu/
COPY links.cfg /root/.links/
