## Command Guide

CashieDB supports the following commands:

| Command | Description | Example |
|---|---|---|
| `SET` | Store a key-value pair | `SET name Go` |
| `GET` | Retrieve a value by key | `GET name` |
| `DEL` | Delete a key-value pair | `DEL name` |

### Usage

Commands are sent over a TCP connection using newline-delimited input.

```text
SET name Go
GET name
DEL name
```

**Note:** All data is stored in memory and is lost when the server shuts down.