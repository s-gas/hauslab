# svcmonitor

![Go](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white)

Monitoring tool that continuously checks the health of the services running in the homelab.

The health is checked by sending HTTP requests and logging their status.

## Behaviour

The program runs continously in a loop with an interval of 1 minute.

On each iteration, the services defined in a static in-memory map are checked via an HTTP request.

The services can be considered:
- `up`: the HTTP request succeeds and returns status code `200`
- `down`: the HTTP request fails or returns a statuc code different than `200`

## Logging

The logs are written in `svcmonitor.log`.

Each entry follows the following format:

```bash
YYYY/MM/DD HH:MM:SS <service> is up|down
```

## How to run

```bash
go build -o svcmonitor .
./svcmonitor
```
