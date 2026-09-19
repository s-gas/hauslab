# nginx

![Docker](https://img.shields.io/badge/Docker-2496ED?style=flat&logo=docker&logoColor=white)

Docker container acting as web server and reverse proxy. It serves the hauslab Web UI as a static file and it routes incoming HTTP requests to the correct services based on the subdomain.

## Web page

Nginx serves a web page with link to various services:

- Grafana
- Prometheus
- AdGuard
- Cockpit
- Tailscale

The 3 Docker services (Grafana, Prometheus, AdGuard) can be reached only from the same network. 

## Virtual hosts (Subdomains)

Only Docker services are set as virtual hosts.

- `hauslab`: Hauslab Web UI
- `gasmetrics.hauslab`: Gasmetrics server
- `grafana.hauslab`: Grafana Web UI
- `prometheus.hauslab`: Prometheus Web UI
- `adguard.hauslab`: AdGuard Web UI
- `setup.adguard.hauslab`: AdGuard Initial Setup

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

## Troubleshooting

Test the configuration file:

```bash
docker exec nginx nginx -t
```

Reload the configuration file:

```bash
docker exec nginx nginx -s reload
```
