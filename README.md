# netwatch

a lightweight TCP connection monitoring utility for Linux
netwatch checks TCP connectivity to a specified host and port, measures connection latency, and displays basic connection statistics.

## Features

- TCP connectivity checks
- Connection latency measurement
- Success and failure statistics
- Custom host, port, and attempt count
- Verbose logging
- Lightweight, statically compiled binary

## Usage

```bash
./netwatch -ip 1.1.1.1 -p 443 -c 5
```

### Flags

| Flag | Description | Default |
|---|---|---|
| `-ip` | Target IP address or hostname | `google.com` |
| `-p` | Target TCP port | `80` |
| `-c` | Number of connection attempts | `3` |
| `-vb` | Enable verbose logging | `false` |

## Build

Requires Go.

```bash
CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o netwatch .
```

## Project Status

**EAL (End of Active Life)**

Netwatch is no longer under active development. No new features or major updates are planned.

Only bug fixes and maintenance patches may be released.

## License

See [LICENSE](LICENSE) for details.
