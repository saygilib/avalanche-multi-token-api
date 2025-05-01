# Quick Start

```docker compose up --build```

**For extracting abi's**
``` node extract-abis.js```

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
ERC20_ADDRESS=0x91E88CbD04b794B7F74CF96211AC1AB5a3639b19
ERC721_ADDRESS=0x6F6AbE23b9C17E22fD01B998195f8e3Cf699c188
ERC1155_ADDRESS=0x8676762eafA005b1eF59ce883Ee9A06d47e086fA
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


/balance?tokenType=erc20&walletAddress=0x...
/balance?tokenType=erc721&walletAddress=0x...
/balance?tokenType=erc1155&walletAddress=0x...&tokenId=1

/all-logs [GET]
Fetch all API logs from the database.

```
# Curl Examples
```
curl -X POST http://localhost:8080/mint \
-H "Content-Type: application/json" \
-d '{
    "tokenType": "erc20",
    "toAddress": "0x8DAf1D28716c98106AAd7368D24cD3FEad66Cad0",
    "amount": "1000"
}'
--------------------------------------------------------------------
curl -X POST http://localhost:8080/mint \
-H "Content-Type: application/json" \
-d '{
    "tokenType": "erc721",
    "toAddress": "0x8DAf1D28716c98106AAd7368D24cD3FEad66Cad0",
    "tokenURI": "https://example.com/metadata/1.json"
}'
--------------------------------------------------------------------
curl -X POST http://localhost:8080/mint \
-H "Content-Type: application/json" \
-d '{
    "tokenType": "erc1155",
    "toAddress": "0x8DAf1D28716c98106AAd7368D24cD3FEad66Cad0",
    "amount": "10",
    "tokenId": "1",
    "tokenURI": "https://example.com/metadata/1155/1.json"
}'

--------------------------------------------------------------------
curl -X GET "http://localhost:8080/balance?tokenType=erc20&walletAddress=0x8DAf1D28716c98106AAd7368D24cD3FEad66Cad0"
--------------------------------------------------------------------

curl -X GET "http://localhost:8080/balance?tokenType=erc721&walletAddress=0x8DAf1D28716c98106AAd7368D24cD3FEad66Cad0"
--------------------------------------------------------------------
curl -X GET "http://localhost:8080/balance?tokenType=erc1155&walletAddress=0x8DAf1D28716c98106AAd7368D24cD3FEad66Cad0&tokenId=1"
--------------------------------------------------------------------

curl -X POST http://localhost:8080/transfer \
-H "Content-Type: application/json" \
-d '{
    "tokenType": "erc20",
    "toAddress": "0x8DAf1D28716c98106AAd7368D24cD3FEad66Cad0",
    "amount": "10"
}'
--------------------------------------------------------------------
curl -X GET http://localhost:8080/all-logs
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
1	"/mint"	"{""tokenType"":""erc20"",""toAddress"":""0x8DAf1D28716c98106AAd7368D24cD3FEad66Cad0"",""amount"":""1000"",""tokenId"":"""",""tokenURI"":""""}"	"{""txHash"":""0xf19705111c266d9fb4746fd1f180fb61b136eb1a2e71d908a9985462a4e0241c""}"	200	"2025-05-01 13:14:18.975929"
2	"/mint"	"{""tokenType"":""erc721"",""toAddress"":""0x8DAf1D28716c98106AAd7368D24cD3FEad66Cad0"",""amount"":"""",""tokenId"":"""",""tokenURI"":""https://example.com/metadata/1.json""}"	"{""txHash"":""0x1b6cfb0c43a3142f7bfbf566dd991c1611f885873a59e5e03a388536767a2623""}"	200	"2025-05-01 13:14:20.33422"
3	"/mint"	"{""tokenType"":""erc1155"",""toAddress"":""0x8DAf1D28716c98106AAd7368D24cD3FEad66Cad0"",""amount"":""10"",""tokenId"":""1"",""tokenURI"":""https://example.com/metadata/1155/1.json""}"	"{""txHash"":""0x85ef4ac4fa8306a1cfb4646b04fef5403ff22d761aaa6f2a03c94e928234d160""}"	200	"2025-05-01 13:14:21.671814"
4	"/balance"	"{""tokenType"":""erc20"", ""walletAddress"":""0x8DAf1D28716c98106AAd7368D24cD3FEad66Cad0""}"	"{""balance"":""1000000000000000000000000000000200001000""}"	200	"2025-05-01 13:14:21.895341"
5	"/balance"	"{""tokenType"":""erc721"", ""walletAddress"":""0x8DAf1D28716c98106AAd7368D24cD3FEad66Cad0""}"	"{""ownedCount"":""2""}"	200	"2025-05-01 13:14:22.375788"
6	"/balance"	"{""tokenType"":""erc1155"", ""walletAddress"":""0x8DAf1D28716c98106AAd7368D24cD3FEad66Cad0""}"	"{""tokenId"":""1"",""balance"":""100""}"	200	"2025-05-01 13:14:22.612128"
7	"/transfer"	"{""tokenType"":""erc20"",""toAddress"":""0x8DAf1D28716c98106AAd7368D24cD3FEad66Cad0"",""amount"":""10"",""tokenId"":"""",""tokenURI"":""""}"	"{""txHash"":""0x7466066efe1b051c513cf5ea39e3656684cf73ba3dea82571b1da3bc11bbd055""}"	200	"2025-05-01 13:14:23.91637"
8	"/all-logs"	"{}"	"[{""action"":""mint"",""tokenType"":""erc20"",""to"":""0x8DAf1D28716c98106AAd7368D24cD3FEad66Cad0"",""amount"":""100000000"",""tokenId"":"""",""tokenURI"":"""",""txHash"":""0xb7af550d54e2bcc55391926091f183ed1f3c1f661cf0e0f13aedbf6a33bd1e4f""}]"	200	"2025-05-01 13:14:23.970924"
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