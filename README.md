# Toy Blockchain in Go

## Overview

This project is a simple blockchain implementation written in Go. It demonstrates the core concepts of blockchain technology, including hashing, proof of work, transaction processing, blockchain validation, persistent storage, and a command-line interface.

## Features

- Block creation
- SHA-256 hashing & Merkle Trees
- Proof of Work mining (Concurrent Goroutines)
- Difficulty Retargeting
- Transaction handling (ECDSA Signatures & Nonces)
- Blockchain validation (Longest Valid Chain sync)
- Ledger with balances
- JSON persistence (AES-GCM Encrypted Wallet)
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

> [!IMPORTANT]
> **Wallet Encryption**: The first time you run any command, the CLI will prompt you to create a password. It uses this password to generate an AES-GCM encrypted `wallet.json`. You must enter this password every time you run a command to unlock your digital identity!

Print the blockchain:

```bash
go run ./cmd print
```

Mine a block (this will grant your wallet 50 coins as a mining reward!):

```bash
go run ./cmd mine
```
> [!NOTE]  
> **How do I get money?**  
> We removed the `faucet` command to prevent infinite money glitches. Now, just like real Bitcoin, the *only* way to get money is to mine it! When you run the `mine` command, the system automatically creates a special "Coinbase" transaction that grants your public key `50` coins as a reward for expending CPU power to secure the network.

Add a transaction (sending money to another address):

```bash
go run ./cmd add Bob 50
```

Mine the transaction into the blockchain:

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

Sync with a competing chain (Fork Resolution):

```bash
go run ./cmd sync other_chain.json
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