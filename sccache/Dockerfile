FROM ubuntu:18.04

ENV SCCACHE_VER 0.2.6

RUN apt-get update && \
    apt-get install -y wget && \
    rm -rf /var/lib/apt/lists/*

RUN wget -o sccache.tar.gz https://github.com/mozilla/sccache/releases/download/${SCCACHE_VER}/sccache-${SCCACHE_VER}-x86_64-unknown-linux-musl.tar.gz
RUN tar --strip-components=1 -xvf sccache.tar.gz