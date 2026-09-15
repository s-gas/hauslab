# tailscale

![Docker](https://img.shields.io/badge/Docker-2496ED?style=flat&logo=docker&logoColor=white)

Docker container running Tailscale.

## How to run

Generate an authentication key:

```
https://console.tailscale.com/admin/settings/keys
```

Store the key in a file under this directory:

```bash
echo -n <auth-key> > ts-authkey.txt
```

Run the container:

```bash
docker compose up --build
```
