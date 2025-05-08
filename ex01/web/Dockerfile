FROM debian:bookworm-slim

RUN apt update \
    && apt install -y curl nginx

RUN mkdir -p /etc/nginx
COPY ./web/nginx /etc/nginx

CMD ["nginx", "-c", "/etc/nginx/nginx.conf", "-g", "daemon off;"]
