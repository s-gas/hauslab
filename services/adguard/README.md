# adguard

![Docker](https://img.shields.io/badge/Docker-2496ED?style=flat&logo=docker&logoColor=white)

Docker container running AdGuard Home, a network ad-blocking DNS server.

## Web UI

```
http://adguard.hauslab
```

`adguard.hauslab` is a hostname mapped to `192.168.178.2` via DNS rewrites.

The request is routed to AdGuard by a reverse proxy based on the hostname.

Login as `s-gas`.

## Upstream DNS servers

```
https://dns.quad9.net/dns-query
https://dns.cloudflare.com/dns-query
```

## DNS rewrites

All the domain names for the homelab are resolved through DNS rewrites, to add a new entry go to:

```
Filters->DNS rewrites
```

## FRITZ!Box setup

To set up AdGuard as DNS server, go to the FrizBox homepage:

```
http://192.168.178.1
```

To change what DNS server the Fritzbox uses, go to:

```
Internet->Zugangsart->DNS-Server
```

Set the DNS server IPv4 to the IP address of AdGuard, which is:

```
192.168.178.2
```

To change what DNS the Fritzbox tells the devices to use via DHCP, go to:

```
Heimnetz->Netzwerkeinstellungen->IP-Adressen->IPv4-Einstellungen
```

Set Lokaler DNS-Server to:

```
192.168.178.2
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
