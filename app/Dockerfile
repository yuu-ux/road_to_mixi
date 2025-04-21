FROM debian:bookworm-slim

RUN apt-get update \
    && apt-get install -y wget

RUN wget https://go.dev/dl/go1.24.2.linux-amd64.tar.gz

RUN rm -rf /usr/local/go && tar -C /usr/local -xzf go1.24.2.linux-amd64.tar.gz

COPY ./etc/.bashrc /root/.bashrc

WORKDIR /root/app
ENV PATH=$PATH:/usr/local/go/bin
