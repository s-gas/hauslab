# sysmetrics

![Go](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white)

Go Server that collects host CPU and RAM usage through `gopsutil` and exposes them at `/metrics` in Prometheus format for scaping.

## How to run

```bash
go build -o sysmetrics .
./sysmetrics
```
