# gohttp

A small HTTP file server for sharing the current directory.

## Usage

Windows:

```powershell
.\gohttp-windows-amd64.exe
```

Linux:

```bash
chmod +x gohttp-linux-amd64
./gohttp-linux-amd64
```

Default port is `8000`. Use `-p` to set another port:

```bash
./gohttp-linux-amd64 -p 8080
```

The server listens on `0.0.0.0`, so it can be accessed from the local network.

## Output

```text
Dir: C:\path\to\folder
Listen: 0.0.0.0:8000
Local: http://localhost:8000
LAN: http://192.168.1.10:8000
Stop: Ctrl+C
```

## Build

Windows:

```powershell
$env:CGO_ENABLED='0'
$env:GOOS='windows'
$env:GOARCH='amd64'
go build -trimpath -ldflags='-s -w -buildid=' -o gohttp-windows-amd64.exe .
.\upx.exe --best --lzma gohttp-windows-amd64.exe
```

Linux:

```powershell
$env:CGO_ENABLED='0'
$env:GOOS='linux'
$env:GOARCH='amd64'
go build -trimpath -ldflags='-s -w -buildid=' -o gohttp-linux-amd64 .
.\upx.exe --best --lzma gohttp-linux-amd64
```
