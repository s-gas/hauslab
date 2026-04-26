# hauslab

This repository contains all the services running on my self-hosted Linux lab. I use it as a personal engineering environment to experiment with backend development and DevOps workflows in a real, self-managed infrastructure setup.

The services that I develop myself are written in Go and are integrated with DevOps tools such as Docker, Prometheus, and Grafana for deployment, containerization, and observability.

## Goal

Since I learn best by doing, the goal of this homelab is to improve my understanding of backend systems, Go, Linux, and DevOps, while also having something real and practical that I could use every day.

## Hardware

Lenovo ThinkCentre M710Q Tiny
- **CPU**: Intel Core i5 2.4GHz
- **Storage**: 256GB SSD
- **RAM**: 8GB
- **OS**: Debian

## Services

Each service has its own `docker-compose.yaml`.

In this way services can be developed, deployed, and restarted independently.

Current services:

- ### [sysmetrics](./sysmetrics)

  Containerized server that exposes CPU and RAM usage.
  
  ![Go](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white)
  ![Docker](https://img.shields.io/badge/Docker-2496ED?style=flat&logo=docker&logoColor=white)

- ### [svcmonitor](./svcmonitor)

  Containerized monitor that sends HTTP requests to the services to check if they are running.

  ![Go](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white)
  ![Docker](https://img.shields.io/badge/Docker-2496ED?style=flat&logo=docker&logoColor=white)

- ### [observability](./observability)

  Containerized observability stack running Prometheus and Grafana.
  
  ![Docker](https://img.shields.io/badge/Docker-2496ED?style=flat&logo=docker&logoColor=white)
  ![Prometheus](https://img.shields.io/badge/Prometheus-E6522C?style=flat&logo=prometheus&logoColor=white)
  ![Grafana](https://img.shields.io/badge/Grafana-F46800?style=flat&logo=grafana&logoColor=white)

- ### [adguard](./adguard)

  Containerized DNS server with ad blocking.

  ![Docker](https://img.shields.io/badge/Docker-2496ED?style=flat&logo=docker&logoColor=white)

## Networks

Since every service has its own `docker-compose.yaml`, communication is enabled through external Docker networks.

These networks are created by the `Makefile` before the containers are started.

Current networks:
- **monitor**:
  
  Allows **svcmonitor** to send request to all the other services

- **sysmetrics-observability**:
  
  Allows communication between **sysmetrics** and the **observability** stack.

## IP Address

The homelab uses a static IP address:

```bash
192.168.178.2
```

In this way it is easily accessible from other devices connected to the same network.

The following entry has been added to `etc/hosts`:

```bash
192.168.178.2 	hauslab
```

In this way the hostname `hauslab` can be used instead of the IP address.

## Host Ports

The following ports are exposed on the host machine:

| Port | Protocol | Service        |
|------|----------|----------------|
| 22   |  TCP     | SSH            |
| 53   |  TCP     | DNS (AdGuard)  |
| 53   |  UDP     | DNS (AdGuard)  |
| 1024 |  TCP     | Grafana        |
| 1025 |  TCP     | AdGuard Web    |
| 1026 |  TCP     | AdGuard Setup  |
