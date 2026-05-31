# nginx

![Docker](https://img.shields.io/badge/Docker-2496ED?style=flat&logo=docker&logoColor=white)

Docker container acting as web server and reverse proxy. It serves the hauslab Web UI as a static file and it routes incoming HTTP requests to the correct services based on the subdomain.

## Virtual hosts (Subdomains)

- `hauslab`: Hauslab Web UI
- `gasmetrics.hauslab`: Gasmetrics server
- `grafana.hauslab`: Grafana Web UI
- `prometheus.hauslab`: Prometheus Web UI
- `adguard.hauslab`: AdGuard Web UI
- `setup.adguard.hauslab`: AdGuard Initial Setup
- `vaultwarden.hauslab`: Vaultwarden

## How to add a virtual host

In `nginx.conf` add a new `server` block:

```nginx
server {
    server_name <subdomain>.hauslab;
    listen 80;
    location / {
        proxy_set_header Host $host;
        proxy_pass http://<service-name>:<port>;
    }
}
```

Add the hostname to AdGuard's DNS rewrites:

```bash
<subdomain>.hauslab 192.168.178.2
```

## Troubleshooting

Test the configuration file:

```bash
docker exec nginx nginx -t
```

Reload the configuration file:

```bash
docker exec nginx nginx -s reload
```
