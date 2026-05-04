# adguard

![Docker](https://img.shields.io/badge/Docker-2496ED?style=flat&logo=docker&logoColor=white)

Docker container running AdGuard Home, a network ad-blocking DNS server.

## Web UI

```bash
http://hauslab:1024
```

`hauslab` is a hostname mapped to `192.168.178.2` in `/etc/hosts`.

## Upstream DNS servers

```bash
https://dns.quad9.net/dns-query
https://dns.cloudflare.com/dns-query
```

## Troubleshooting

### DNS lookup

```bash
dig @192.168.178.2 google.com
```

It should show `status: NOERROR`. If it shows `status: SERVFAIL` means that the upstream DNS server is unreachable and it should be changed via the Web UI.

### DNS port

Check who is listening on port 53:

```bash
ss -tulnp | grep :53
```

The output should have something like:

```bash
tcp   LISTEN 0      4096                                     0.0.0.0:53        0.0.0.0:*
tcp   LISTEN 0      4096                                        [::]:53           [::]:*
```

The process name is not visible because it runs in a Docker container.

### Container logs

To check the container logs, first find the container ids:

```bash
docker ps
```

Then check its logs:

```bash
docker logs <container id>
```
