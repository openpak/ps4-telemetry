# PS4 Telemetry Sink

OpenPak's PlayStation 4 telemetry sink.

| | |
| --- | --- |
| Network | PlayStation 4 |
| Role | Accepts telemetry/usage-logging requests (`nsx.sec.np.dl.playstation.net` family), logs them, answers 200 and discards the body |
| Depends on | Nothing (Go standard library only) |
| Used by | PlayStation 4 consoles via OpenPak DNS |
| Image | `ghcr.io/openpak/ps4-telemetry` |
| Licence | AGPL-3.0-only |

## What it serves

| Endpoint | What |
| --- | --- |
| any path | 200, empty body; method, host and path are logged |
| `GET /healthz` | Health check |

## Run

```sh
cp example.env .env
set -a; . ./.env; set +a
GOWORK=off go run .
```

Tests: `GOWORK=off go test ./...`

## Docs

- [Protocol](docs/protocol.md)
- [Deploy](docs/deploy.md)

## Licence

AGPL-3.0-only. See [LICENSE](LICENSE).
