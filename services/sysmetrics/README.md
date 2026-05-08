# sysmetrics

![Go](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-2496ED?style=flat&logo=docker&logoColor=white)

Containerized server that collects host-level resources usage and exposes the data for external services.

## Host-level resources

The server collects through `gopsutil` the following resources:
- CPU usage
- RAM usage

In order to read host-level resources, the host's `/proc` directory is bind mounted into the container.

## Endpoint

This service does **not expose ports to the host system**.

The HTTP server listens on port `1024` and is reachable by the **Observability** stack through an external Docker network. 

Metrics are exposed at:

```
/metrics
```

## Exposition format

The data is exposed in the Prometheus Text Format, including the metadata:

```yaml
# HELP cpu_usage_percent Current CPU usage as a percentage
# TYPE cpu_usage_percent gauge
cpu_usage_percent 12.34

# HELP ram_usage_percent Current RAM usage as a percentage
# TYPE ram_usage_percent gauge
ram_usage_percent 56.78
```

## How to run

```bash
docker compose up -d
```

## How to stop

```bash
docker compose down
```
