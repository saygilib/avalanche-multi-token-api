#!/bin/bash

BASE_URL=http://localhost:8080

echo "🔵 Testing /mint (ERC20)..."
curl -X POST $BASE_URL/mint \
-H "Content-Type: application/json" \
-d '{
    "tokenType": "erc20",
    "toAddress": "0x8DAf1D28716c98106AAd7368D24cD3FEad66Cad0",
    "amount": "1000"
}'
echo -e "\n"

echo "🔵 Testing /mint (ERC721)..."
curl -X POST $BASE_URL/mint \
-H "Content-Type: application/json" \
-d '{
    "tokenType": "erc721",
    "toAddress": "0x8DAf1D28716c98106AAd7368D24cD3FEad66Cad0",
    "tokenURI": "https://example.com/metadata/1.json"
}'
echo -e "\n"

echo "🔵 Testing /mint (ERC1155)..."
curl -X POST $BASE_URL/mint \
-H "Content-Type: application/json" \
-d '{
    "tokenType": "erc1155",
    "toAddress": "0x8DAf1D28716c98106AAd7368D24cD3FEad66Cad0",
    "amount": "10",
    "tokenId": "1",
    "tokenURI": "https://example.com/metadata/1155/1.json"
}'
echo -e "\n"

echo "🔵 Testing /balance (ERC20)..."
curl -X GET "$BASE_URL/balance?tokenType=erc20&walletAddress=0x8DAf1D28716c98106AAd7368D24cD3FEad66Cad0"
echo -e "\n"

echo "🔵 Testing /balance (ERC721)..."
curl -X GET "$BASE_URL/balance?tokenType=erc721&walletAddress=0x8DAf1D28716c98106AAd7368D24cD3FEad66Cad0"
echo -e "\n"

echo "🔵 Testing /balance (ERC1155)..."
curl -X GET "$BASE_URL/balance?tokenType=erc1155&walletAddress=0x8DAf1D28716c98106AAd7368D24cD3FEad66Cad0&tokenId=1"
echo -e "\n"

echo "🔵 Testing /transfer (ERC20)..."
curl -X POST $BASE_URL/transfer \
-H "Content-Type: application/json" \
-d '{
    "tokenType": "erc20",
    "toAddress": "0x8DAf1D28716c98106AAd7368D24cD3FEad66Cad0",
    "amount": "10"
}'
echo -e "\n"

echo "🔵 Testing /all-logs..."
curl -X GET $BASE_URL/all-logs
echo -e "\n"
