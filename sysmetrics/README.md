# sysmetrics

![Screenshot](assets/screenshot-dashboard.png)

## Description

System monitor that collects CPU and RAM usage and displays the metrics through Prometheus and Grafana.

### Diagram of the infrastructure:

```mermaid
flowchart TB
    A([WWW])
    B[Grafana]
    C[Prometheus]
    D[Go Server]
    E[(Volume)]
    F[(Volume)]

    A <-->|3000| B
    B <-->|9090| C
    C <-->|1024| D

    B <----> E
    C <---> F
```

### Services

- **Go Server**  
  Runs in a Docker container and exposes the data at `/metrics`.
  In order to read the metrics of the host machine, the host`/proc` directory is mounted into the container, overwriting its `/proc` directory.

- **Prometheus**  
  Runs in a Docker container and scrapes the endpoint every 10 seconds and stores the data.

- **Grafana**  
  Runs in a Docker container and displays the data in a dashboard.

### Volumes

Two persistent volumes store Prometheus and Grafana data.

## How to run

*Note: to run this service you need to have [Docker](https://www.docker.com/) and [Go](https://go.dev/) installed.*

Clone the hauslab repo:

```bash
git clone https://github.com/s-gas/hauslab.git
```

Change to the service directory:

```bash
cd hauslab/sysmetrics
```

Run the containers:

```bash
docker compose up
```

To stop the containers:

```bash
docker compose down
```

To wipe the data in the named volumes:

```bash
docker compose down -v
```
