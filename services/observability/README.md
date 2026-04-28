# observability

![Prometheus](https://img.shields.io/badge/Prometheus-E6522C?style=flat&logo=prometheus&logoColor=white)
![Grafana](https://img.shields.io/badge/Grafana-F46800?style=flat&logo=grafana&logoColor=white)

Containerized observability stack running Prometheus and Grafana via Docker Compose.

## Services

- **Prometheus**  
  Metrics scraping and storage.

- **Grafana**  
  Metrics visualization.

## How to run

Create the file to store the admin password necessary to access the dashboard:

```bash
mkdir secrets
printf '<password>' > secrets/grafana_password.txt
```

Run the containers:

```bash
docker compose up
```

## Grafana Dashboard

```bash
http://hauslab:1024
```

`hauslab` is a hostname mapped to `192.168.178.2` in `/etc/hosts`.

`192.168.178.2` is the static IP address of the homelab.

Login as `admin` using the password stored in `secrets/grafana_password.txt`.

## How to add scrape targets

In `prometheus/prometheus.yml` add a new entry under `scrape_configs`:

```yaml
  - job_name: "<scrape-target>"
    static_configs:
      - targets: ["<service-name>:<port>"]
```

## How to stop

To stop the containers:

```bash
docker compose down
```

## Troubleshooting

### Wrong password

A "wrong password" issue could be caused by the fact that the secret is set only on first boot. If the grafana volume already exists, it will be ignored. If the volume was created before the secret was added, you can login using `admin`/`admin`. Otherwise the solution is to delete the volumes.

To wipe all the data in the volumes and restart:

```bash
docker compose down -v && docker compose up
```

Another possible reason for "wrong password" is having `\n` in the password. To avoid this, always use `printf`.

### Container logs

To check the container logs, first find the container ids:

```bash
docker ps
```

Then check its logs:

```bash
docker logs <container id>
```
