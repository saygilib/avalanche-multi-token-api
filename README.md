# Quick Start

```docker compose up --build```

# Overview

This project includes:
- 3 Smart Contracts (ERC20, ERC721, ERC1155) deployed on Avalanche Fuji.
- Go backend API with Gin.
- PostgreSQL for API logging.
- Dockerized setup.
- Unit tests (Truffle + Mocha/Chai).

# Smart Contract Deployment

**Truffle commands used:**

```
truffle compile
truffle migrate --network fuji
```

**Contracts deployed to:**
```
Contract	Address
ERC20	0x7BcC0Ee8628c0abE0D60424a53081773792531b8
ERC721	0xc5fCF19357B076Fa5efbE444E23610D12A400797
ERC1155	0xCe6ca22e42128bC9576B3b877aD1e5b365692df2
```

# Go Backend API Endpoints

```/mint [POST]
{
  "tokenType": "erc20 | erc721 | erc1155",
  "toAddress": "0x...",
  "amount": "1000",         // erc20 / erc1155
  "tokenId": "1",           // erc1155 / erc721
  "tokenURI": "https://..." // erc721 / erc1155
}

/transfer [POST]
Transfer tokens.

Same input as /mint.

/balance [GET]
Get balances:

bash
Kopyala
Düzenle
/balance?tokenType=erc20&walletAddress=0x...
/balance?tokenType=erc721&walletAddress=0x...
/balance?tokenType=erc1155&walletAddress=0x...&tokenId=1

/all-logs [GET]
Fetch all API logs from the database.

```


# ABI Integration in Go

Used abigen to generate Go bindings for ERC20, ERC721, ERC1155.

ABIs are stored inside go-backend/abi/.

Contracts loaded via environment variables:

ERC20_ADDRESS

ERC721_ADDRESS

ERC1155_ADDRESS

# PostgreSQL Schema

**api_logs table:**
```
Field	Type
endpoint	text
request_payload	text
response	text
status	integer
created_at	timestamp
```

Example row:
```
endpoint	request_payload	response	status	created_at
/mint	{...}	{"txHash":"0x..."}	200	2025-05-01 10:55:56
```

# Docker
**Build & run:**

```docker-compose up --build```

**Services:**

postgres: PostgreSQL 15.

go-backend: Go API (Gin) on port 8080.

**Environment variables:**

Stored in go-backend/.env.

**Port mapping:**

API → localhost:8080

PostgreSQL → localhost:5433 (if you changed the port)

# Project Structure 

```
BBCode-PROJECT/
├── contracts/         # Solidity contracts
├── migrations/        # Truffle migrations
├── test/              # Smart contract tests
├── go-backend/
│   ├── abi/
│   ├── contracts/
│   ├── handlers/
│   ├── db/
│   ├── utils/
│   ├── main.go
│   ├── .env
│   └── Dockerfile
├── docker-compose.yml
├── test.sh
├── README.md
``` 