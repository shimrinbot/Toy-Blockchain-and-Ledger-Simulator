# Toy Blockchain in Go

## Overview

This project is a simple blockchain implementation written in Go. It demonstrates the core concepts of blockchain technology, including hashing, proof of work, transaction processing, blockchain validation, persistent storage, and a command-line interface.

## Features

- Block creation
- SHA-256 hashing
- Proof of Work mining
- Transaction handling
- Blockchain validation
- Ledger with balances
- JSON persistence
- Command-line interface
- Unit testing

## Project Structure

```
block/
blockchain/
cli/
cmd/
ledger/
storage/
tests/
```

## Installation

```bash
git clone <repository>
cd toy-blockchain
go mod tidy
```

## Running

Print the blockchain:

```bash
go run ./cmd print
```

Get initial funds from the faucet:

```bash
go run ./cmd faucet Alice 100
```

Add a transaction:

```bash
go run ./cmd add Alice Bob 50
```

Mine a block:

```bash
go run ./cmd mine
```

Validate the blockchain:

```bash
go run ./cmd validate
```

View balances:

```bash
go run ./cmd balances
```

Run all tests:

```bash
go test -v ./...
```

## Technologies

- Go
- SHA-256
- JSON
- Proof of Work

## Author

Shimrin