FROM debian:bookworm-slim AS base
RUN apt-get update \
    && apt-get install -y wget \
    && rm -rf /var/lib/apt/lists/*
ENV PATH=$PATH:/usr/local/go/bin
WORKDIR /app

FROM base AS dev
RUN wget https://go.dev/dl/go1.24.2.linux-amd64.tar.gz \
    && rm -rf /usr/local/go && tar -C /usr/local -xzf go1.24.2.linux-amd64.tar.gz
COPY etc/.bashrc /root/.bashrc
RUN go install github.com/air-verse/air@latest
ENV PATH=$PATH:/root/go/bin
COPY app .
CMD ["air", "-c", ".air.toml"]

FROM base AS builder
RUN wget https://go.dev/dl/go1.24.2.linux-amd64.tar.gz \
    && rm -rf /usr/local/go && tar -C /usr/local -xzf go1.24.2.linux-amd64.tar.gz
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o server .

FROM base AS runner
COPY . .
COPY --from=builder /app/server .
EXPOSE 1323
CMD ["./server"]
