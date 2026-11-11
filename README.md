# CashieDB

A lightweight, in-memory key-value database written in Go for TCP-based client-server access.

<picture>
  <source srcset="docs/assets/cashew.jpeg" type="image/avif" />
  <img src="docs/assets/cashew.avif" alt="CashieDB demo" width="40%" style="display:block; margin:0 auto; background-color:#000; border-radius:8px;" />
</picture>

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
