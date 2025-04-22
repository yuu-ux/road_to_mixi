FROM debian:bookworm-slim

RUN apt-get update \
    && apt-get install -y wget

RUN wget https://go.dev/dl/go1.24.2.linux-amd64.tar.gz

RUN rm -rf /usr/local/go && tar -C /usr/local -xzf go1.24.2.linux-amd64.tar.gz

ENV PATH=$PATH:/usr/local/go/bin

COPY etc/.bashrc /root/.bashrc

WORKDIR /app

RUN go install github.com/air-verse/air@latest
ENV PATH=$PATH:/root/go/bin

COPY app/.air.toml .

CMD ["air", "-c", ".air.toml"]

# ビルドステージ
# COPY app/go.mod .
# COPY app/go.sum .
#
# RUN go mod download
#
# COPY app .
#
# RUN go build -o server .

# バイナリをコピー
# FROM debian:bookworm-slim
#
# WORKDIR /app
#
# COPY --from=builder /app/server .
#
# EXPOSE 1323
#
# CMD ["./server"]
