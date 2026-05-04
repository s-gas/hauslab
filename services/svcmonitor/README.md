# svcmonitor

![Go](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-2496ED?style=flat&logo=docker&logoColor=white)

Containerized monitoring tool that continuously checks the status of the services running in the homelab, which are exposed for external services.

## Concurrency

The program uses two concurrent patterns:

- HTTP Server: Exposes the current status of all tracked services.

- Background Loop: `goroutine` that runs an infinite loop with a 30-seconds interval, checking services and updating their status.

Data races are prevented by using **mutexes**.

## Status

The services can be considered:
- `UP`: the HTTP request succeeds and returns status code `200`
- `DOWN`: the HTTP request fails or returns a status code different than `200`

## Endpoint

This service does **not expose ports to the host system**.

The HTTP server listens on port `1024` and is reachable by the **Observability** stack through an external Docker network called `svcmonitor-observability`.

Metrics are exposed at:

```
/metrics
```

## Exposition format

The data is exposed in the Prometheus Text Format, including the metadata:

```bash
# HELP service_up Status of the service (1 for UP, 0 for DOWN)
# TYPE service_up gauge
service_up{service="grafana"} 1
```

## Logging

Each entry follows the following format:

```bash
YYYY/MM/DD HH:MM:SS <service> UP|DOWN
```

## How to run

Start the container:

```bash
docker compose up -d
```

## How to stop

```bash
docker compose down
```

## Troubleshooting

### Container logs

To check the container logs, first find the container ids:

```bash
docker ps
```

Then check its logs:

```bash
docker logs <container id>
```
