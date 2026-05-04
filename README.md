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

- ### [sysmetrics](./services/sysmetrics)

  Containerized server that exposes CPU and RAM usage.
  
  ![Go](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white)
  ![Docker](https://img.shields.io/badge/Docker-2496ED?style=flat&logo=docker&logoColor=white)

- ### [svcmonitor](./services/svcmonitor)

  Containerized monitor that checks the status of the services and exposes them.

  ![Go](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white)
  ![Docker](https://img.shields.io/badge/Docker-2496ED?style=flat&logo=docker&logoColor=white)

- ### [observability](./services/observability)

  Containerized observability stack running Prometheus and Grafana.
  
  ![Docker](https://img.shields.io/badge/Docker-2496ED?style=flat&logo=docker&logoColor=white)
  ![Prometheus](https://img.shields.io/badge/Prometheus-E6522C?style=flat&logo=prometheus&logoColor=white)
  ![Grafana](https://img.shields.io/badge/Grafana-F46800?style=flat&logo=grafana&logoColor=white)

- ### [adguard](./services/adguard)

  Containerized DNS server with ad blocking.

  ![Docker](https://img.shields.io/badge/Docker-2496ED?style=flat&logo=docker&logoColor=white)

- ### [nginx](./services/nginx)

  Web server and reverse proxy.

  ![Docker](https://img.shields.io/badge/Docker-2496ED?style=flat&logo=docker&logoColor=white)

## Networks

Since every service has its own `docker-compose.yaml`, communication is enabled through external Docker networks.

These networks are created by the `Makefile` before the containers are started.

Current networks:
- **monitor**:
  
  Allows **svcmonitor** to send request to all the other services.

- **sysmetrics-observability**:
  
  Allows communication between **sysmetrics** and the **observability** stack.

- **svcmonitor-observability**:
  
  Allows communication between **svcmonitor** and the **observability** stack.

- **proxy**:

  Allows **nginx** to route requests to all services accessible via subdomain.

## Scripts

- ### [pull-and-run](./scripts/pull-and-run)
  
  Automated script that pulls the latest commit and redeploys all services.

  ![Bash](https://img.shields.io/badge/Bash-4EAA25?style=flat&logo=gnubash&logoColor=white)


## IP Address

The homelab uses a static IP address:

```bash
192.168.178.2
```

In this way it is easily accessible from other devices connected to the same network.

## Hostnames

A reverse proxy ([nginx](./services/nginx)) allows each service to be accessible via its own subdomain.

All hostnames points to the same IP and are routed to the correct service by the reverse proxy.

The following entries need to be added to `/etc/hosts` on every device that needs to access the services:

```bash
192.168.178.2   hauslab
192.168.178.2   grafana.hauslab
192.168.178.2   prometheus.hauslab
192.168.178.2   adguard.hauslab
192.168.178.2   setup.adguard.hauslab
```

## Host Ports

The following ports are exposed on the host machine:

| Port | Protocol | Service        |
|------|----------|----------------|
| 22   |  TCP     | SSH            |
| 53   |  TCP     | DNS (AdGuard)  |
| 53   |  UDP     | DNS (AdGuard)  |
| 80   |  TCP     | Nginx          |

## SSH

The homelab runs an SSH server, which means it can be accessed by any device in the local network:

```bash
ssh <username>@hauslab
```
