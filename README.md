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

Each service runs in containers and has its own `docker-compose.yaml`.

In this way services can be developed, deployed, and restarted independently.

Current services:

- ### [sysmetrics](./services/sysmetrics)

  Server that exposes CPU and RAM usage.
  
  ![Go](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white)
  ![Docker](https://img.shields.io/badge/Docker-2496ED?style=flat&logo=docker&logoColor=white)

- ### [svcmonitor](./services/svcmonitor)

  Monitor that checks the status of the services and exposes them.

  ![Go](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white)
  ![Docker](https://img.shields.io/badge/Docker-2496ED?style=flat&logo=docker&logoColor=white)

- ### [observability](./services/observability)

  Observability stack running Prometheus and Grafana.
  
  ![Docker](https://img.shields.io/badge/Docker-2496ED?style=flat&logo=docker&logoColor=white)
  ![Prometheus](https://img.shields.io/badge/Prometheus-E6522C?style=flat&logo=prometheus&logoColor=white)
  ![Grafana](https://img.shields.io/badge/Grafana-F46800?style=flat&logo=grafana&logoColor=white)

- ### [adguard](./services/adguard)

  DNS server with ad blocking.

  ![Docker](https://img.shields.io/badge/Docker-2496ED?style=flat&logo=docker&logoColor=white)

- ### [nginx](./services/nginx)

  Web server and reverse proxy.

  ![Docker](https://img.shields.io/badge/Docker-2496ED?style=flat&logo=docker&logoColor=white)

- ### [postgres](./services/postgres)

  PostgreSQL database.

  ![Postgres](https://img.shields.io/badge/PostgreSQL-4169E1?logo=postgresql&logoColor=white)
  ![Docker](https://img.shields.io/badge/Docker-2496ED?style=flat&logo=docker&logoColor=white)

## Network

Since every service has its own `docker-compose.yaml`, communication is enabled through an external Docker network.

This network is created by the `Makefile` before the containers are started.

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

## Makefile

In the `services` directory there is a `Makefile` that can be used to start and stop all the services.

## Scripts

- ### [pull-and-run](./scripts/pull-and-run)
  
  Cron job that pulls the latest Docker images and redeploys all services by running the `Makefile`.

  ![Bash](https://img.shields.io/badge/Bash-4EAA25?style=flat&logo=gnubash&logoColor=white)
