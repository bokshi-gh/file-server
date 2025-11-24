# Goserve

A simple, lightweight HTTP file server written in Go. Supports directory listing and serving files with proper MIME types.

## Features
- Lists directories
- Serves static files over HTTP
- Automatic index.html fallback for directories
- Prevents directory traversal

## Installation

### Install from source (build on your machine)

Clone the repository:

```bash
git clone https://github.com/bokshi-gh/file-server.git
cd file-server
```

Build the server:

```bash
go build -o goserve ./cmd/goserve
```

### Install using build scripts

For Unix-like systems (Linux, macOS):

```bash
curl -fsSL https://raw.githubusercontent.com/youruser/file-server/main/scripts/build.sh | bash
```

For Windows:

```powershell
irm https://raw.githubusercontent.com/bokshi-gh/file-server/main/scripts/build.ps1 | iex
```

### Install a specific version

Replace vX.Y.Z with the version you want

Unix-like:

```bash
curl -fsSL https://raw.githubusercontent.com/bokshi-gh/file-server/vX.Y.Z/scripts/build.sh | bash
```

Windows:
```powershell
irm https://raw.githubusercontent.com/bokshi-gh/file-server/vX.Y.Z/scripts/build.ps1 | iex
```

## Usage

Run the server:

```bash
goserve --root ./public --host 0.0.0.0 --port 8080 --v
```
