# Bushuray-core

Back-end of [Bushuray-tui](https://github.com/amirhosseinkhorshidi/Bushuray-tui). Handles proxy management, subscription updates, and latency testing via a local TCP server.

## Requirements

The following binaries must be present alongside `bushuray-core` or in the `bin/` subdirectory:

| Binary | Purpose |
|---|---|
| `xray` | Proxy engine |
| `v2parser` | URI parser |

## Debugging

```bash
tail -f /tmp/bushuray-core-debug.log
```
