# services

This directory contains all the services running in the homelab.

## Makefile

The `Makefile` can be used to:
- start all the services:
```bash
make up
```
- stop all the services:
```bash
make down
```

When adding a new service, remember to add the relevant `docker compose` commands to both the `up` and `down` targets.
