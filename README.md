# 🔍 Network Security Scanner (CLI)

A basic TCP port scanner written in Go. Useful for quickly checking open ports on a target IP.

## 🚀 Features

- Scans a single host
- Supports individual port or port range
- Fast with goroutine-based concurrency

## 🛠 Usage

```bash
go run main.go <target IP> <port|port-range>
