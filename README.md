# go-http-server
A simple multithreaded http server written in Go.

## Usage

The server can be started using the following command:

```bash
go run server.go
```

Alternatively,

```bash
go build server.go
./server
```

### Command line flags

This application currently supports the following flags:

| Flag | Description | Default Value | Example |
| --- | --- | --- | --- |
| port | Port on which the server listens | 8000 | `go run server.go -port=8080` or `./server --port 8001`|

## Suggestions?

Please open an issue or code it up and send me a PR.
