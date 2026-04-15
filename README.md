# hauslab

This repository contains all the services I've built and deployed on my self-hosted Linux lab, which I use as a personal engineering environment to experiment with backend development and DevOps workflows in a real, self-managed infrastructure setup.

The services that I develop myself are written in Go and are integrated with DevOps tools such as Docker, Prometheus, and Grafana for deployment, containerization, and observability.

## Goal

Since I learn best by doing, the goal of this homelab is to improve my understanding of backend systems, Go, Linux and DevOps, while also having something real and practical that I could use every day.

## Hardware

Lenovo ThinkCentre M710Q Tiny
- **CPU**: Intel Core i3 3.2GHz
- **Storage**: 256GB SSD
- **RAM**: 8GB
- **OS**: Debian

## Services

- ### [sysmetrics](./sysmetrics)

  Server that exposes CPU and RAM usage.
  
  ![Go](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white)

- ### [svcmonitor](./svcmonitor)

  Monitor that sends HTTP requests to the services to check if they are running.

  ![Go](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white)

- ### [observability](./observability)

  Containerized observability stack running Prometheus and Grafana via Docker Compose.
  
  ![Docker](https://img.shields.io/badge/Docker-2496ED?style=flat&logo=docker&logoColor=white)
  ![Prometheus](https://img.shields.io/badge/Prometheus-E6522C?style=flat&logo=prometheus&logoColor=white)
  ![Grafana](https://img.shields.io/badge/Grafana-F46800?style=flat&logo=grafana&logoColor=white)
