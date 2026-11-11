# CashieDB

A lightweight, in-memory key-value database written in Go for TCP-based client-server access.

![CashieDB demo](docs/assets/cashew.avif)

## Features

- In-memory key-value storage
- TCP-based client-server communication
- Concurrent client connection handling
- Small, focused Go implementation for key-value operations

## Getting Started

```bash
go run ./main.go
```

## How It Works

CashieDB runs as a TCP server. Clients connect to the server, send commands, and receive responses over the network. Each client connection is handled independently, enabling multiple clients to interact with the database concurrently.

Data is stored in memory and is lost when the server shuts down.

## Tech Stack

- Go
- TCP networking
- In-memory storage

## Project Status

Work in progress.
