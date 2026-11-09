# CashieDB

A lightweight, in-memory key-value database written in Go.

CashieDB is a small database server that stores key-value pairs in memory and allows clients to interact with it over TCP.

## Features

- In-memory key-value storage
- TCP-based client-server communication
- Concurrent client connection handling
- Lightweight and written entirely in Go

## Getting Started

### Run the server

```bash
go run ./main.go
```

## How It Works

CashieDB runs as a TCP server. Clients connect to the server and send requests to perform operations on the in-memory key-value store.

Data is stored in memory and is lost when the server shuts down.

## Tech Stack

- **Language:** Go
- **Networking:** TCP
- **Storage:** In-memory key-value store

## Project Status

Work in progress 🚧
