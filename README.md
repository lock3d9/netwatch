# netwatch
(educational project)
## usage

```bash
./netwatch -ip 1.1.1.1 -p 443 -c 5
```

### flags

| Flag | Description | Default |
|---|---|---|
| `-ip` | Target IP address or hostname | `google.com` |
| `-p` | Target TCP port | `80` |
| `-c` | Number of connection attempts | `3` |
| `-vb` | Enable verbose logging | `false` |

## Build
```bash
CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o netwatch .
```

## project Status

**EAL (End of Active Life)**
## License

See [LICENSE](LICENSE) for details.
