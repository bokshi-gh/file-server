# Goserve

A simple, lightweight HTTP file server written in Go. Supports directory listing and serving files with proper MIME types.

## Features
- Serves static files and lists directory over HTTP.
- Automatic index.html fallback for directories.
- Prevents directory traversal.

## Installation

Clone the repository:

```bash
git clone https://github.com/bokshi-gh/file-server.git
cd file-server
```

Build the server:

```bash
go build -o goserve ./cmd/goserve
```

## Usage

Run the server:

```bash
./goserve --root ./public --host 0.0.0.0 --port 8080 --v
```
