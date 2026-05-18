# Port Scan Helper

A small TCP listener tool for checking whether ports on an internal host are reachable from outside.

The tool opens a range of local TCP ports. You can then scan or connect to the host from an external network. If a connection reaches one of the opened ports, the tool prints the source address.

## Usage

```bash
./port_scan_linux_amd64 -start 1000 -end 2000
```

Windows:

```powershell
.\port_scan_windows_amd64.exe -start 1000 -end 2000
```

## Options

```text
-start int
    start port, default 1000

-end int
    end port, default 2000
```

## Example Output

```text
[*] Starting listeners on ports 1000-2000...
[+] Listening in background on ports 1000-2000.
[*] Scan or connect to this host from outside.
[*] Press Ctrl+C to stop and release all ports.
[!] Port reachable: 1500 | Source: 203.0.113.10:53218
```

## Build

Linux amd64:

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags="-s -w" -o port_scan_linux_amd64 .
```

Windows amd64:

```bash
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -trimpath -ldflags="-s -w" -o port_scan_windows_amd64.exe .
```

## Notes

- TCP only.
- Use only on systems and networks where you have authorization.
- Press `Ctrl+C` to stop the tool and release all opened ports.
