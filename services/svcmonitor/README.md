# svcmonitor

![Go](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-2496ED?style=flat&logo=docker&logoColor=white)

Containerized monitoring tool that continuously checks the health of the services running in the homelab.

The health is checked by sending HTTP requests and logging their status.

## Behaviour

The program runs continuosly in a loop with an interval of 1 minute.

On each iteration, the services defined in a static in-memory map are checked via an HTTP request.

The services can be considered:
- `UP`: the HTTP request succeeds and returns status code `200`
- `DOWN`: the HTTP request fails or returns a status code different than `200`

## Logging

The logs are written in `svcmonitor.log`, bind mounted from the host into the container.

Each entry follows the following format:

```bash
YYYY/MM/DD HH:MM:SS <service> UP|DOWN
```

## How to run

Before running, make sure `svcmonitor.log` exists:

```bash
touch svcmonitor.log
```

Start the container:

```bash
docker compose up -d
```

## How to stop

```bash
docker compose down
```

## Troubleshooting

### Log file

Make sure `svcmonitor.log` exists on the host before starting the container. If the file is missing, Docker will create it as a directory, causing the bind mount to fail.

For now, I don't have a rotating feature for the log file, which means that with time it will get very big. The solution for now is to stop the container, delete the log file and create a new one.
