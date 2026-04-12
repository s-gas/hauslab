# traces

![Screenshot](assets/screenshot-dashboard.png)

## Description

System monitor that collects CPU and RAM usage and displays the metrics through Prometheus and Grafana.

### Services

- **Go App**  
  Runs natively on the machine and exposes the data at `/metrics`.

- **Prometheus**  
  Runs in a Docker container and scrapes the endpoint every 15 seconds and stores the data.

- **Grafana**  
  Runs in a Docker container and displays the data in a dashboard.

### Volumes

Two persistent volumes store Prometheus and Grafana data.

## How to run

*Note: to run this project you need to have [Docker](https://www.docker.com/) and [Go](https://go.dev/) installed.*

Clone this repository:

```bash
git clone https://github.com/s-gas/traces.git
```

Change to the project directory:

```bash
cd traces
```

Run the app and the containers through the Makefile:

```bash
make run
```

To stop the app press `Ctrl-C`.

To stop the containers:

```bash
docker compose down
```

To wipe the data in the named volumes:

```bash
docker compose down -v
```
