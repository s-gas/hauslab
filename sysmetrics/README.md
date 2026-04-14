# sysmetrics

![Go](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white)

Server that collects host-level resources usage and exposes the data for external services.

## Host-level resources

The server collects through `gopsutil` the following resources:
- CPU usage
- RAM usage

## Endpoint

The server listens on port `1024`.

Data is exposed at `/metrics`.

## Exposition format

The data is exposed in the Prometheus Text Format, including the metadata:

```bash
# HELP cpu_usage_percent Current CPU usage as a percentage
# TYPE cpu_usage_percent gauge
cpu_usage_percent 12.34

# HELP ram_usage_percent Current RAM usage as a percentage
# TYPE ram_usage_percent gauge
ram_usage_percent 56.78
```

## How to run

```bash
go build -o sysmetrics .
./sysmetrics
```
