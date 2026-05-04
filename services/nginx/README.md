# nginx

![Docker](https://img.shields.io/badge/Docker-2496ED?style=flat&logo=docker&logoColor=white)

Docker container running a reverse proxy that routes incoming HTTP requests to the correct services based on the subdomain.

## Virtual hosts (Subdomains)

- `grafana.hauslab`
- `prometheus.hauslab`

## How to add a virtual host

In `nginx.conf` add a new `server` block:

```nginx
server {
    server_name <subdomain>.hauslab;
    listen 80;
    location / {
        proxy_set_header Host $host;
        proxy_pass http://<container_name>:<port>;
    }
}
```

Add the hostname to `/etc/hosts` on every device that needs access:

```bash
192.168.178.2    <subdomain>.hauslab
```
