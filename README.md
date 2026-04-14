# hauslab

My homelab, the playground where I experiment with Backend and DevOps concepts by self-hosting services on a Linux machine. 

## Hardware
Lenovo ThinkCentre M710Q Tiny
- **CPU**: Intel Core i3 3.2GHz
- **Storage**: 256GB SSD
- **RAM**: 8GB
- **OS**: Debian

## Services
- ### [sysmetrics](./sysmetrics)

  Server that collects CPU and RAM usage and exposes them to `/metrics`.
  
  ![Go](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white)

- ### [svcmonitor](./svcmonitor)

  Service monitor that sends HTTP requests to the services to check if they are running.

  ![Go](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white)

- ### [observability](./observability)

  Containerized observability stack running Prometheus and Grafana via Docker Compose.
  
  ![Docker](https://img.shields.io/badge/Docker-2496ED?style=flat&logo=docker&logoColor=white)
  ![Prometheus](https://img.shields.io/badge/Prometheus-E6522C?style=flat&logo=prometheus&logoColor=white)
  ![Grafana](https://img.shields.io/badge/Grafana-F46800?style=flat&logo=grafana&logoColor=white)
