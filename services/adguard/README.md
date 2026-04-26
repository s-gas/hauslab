# adguard

Docker container running AdGuard Home, a network ad-blocking DNS server.

## Updating the software

From time to time run:

```bash
docker compose pull && docker compose up -d
```

In this way Docker fetches the latest image and it recreates the container if the image has changed.
