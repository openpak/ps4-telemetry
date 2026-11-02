# Deploy

## Run

```sh
cp example.env .env
set -a; . ./.env; set +a
GOWORK=off go run .
```

Quadlet: `deploy/openpak-ps4-telemetry.container`; image `ghcr.io/openpak/ps4-telemetry`; listens on `:8541`.

## Environment

| Variable | Default | |
| --- | --- | --- |
| `PS4_TELEMETRY_ADDR` | `:8541` | Listen address |

## DNS

OpenPak DNS / Traefik must route the PS4 telemetry hosts
(`nsx.sec.np.dl.playstation.net`) to this service. The TLS route is added once
a console capture confirms how the host is called.
