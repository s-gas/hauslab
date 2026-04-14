# sysmetrics

![Go](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white)

Server that collects host CPU and RAM usage through `gopsutil`.

The server listens on port `1024` and exposes data at `/metrics` in Prometheus format for scraping.

## How to run

```bash
go build -o sysmetrics .
./sysmetrics
```
