package handlers

import (
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/saygilib/avalanche-project/go-backend/contracts"
	"github.com/saygilib/avalanche-project/go-backend/db"
	"github.com/saygilib/avalanche-project/go-backend/utils"
)

type MintRequest struct {
	TokenType string `json:"tokenType"`
	ToAddress string `json:"toAddress"`
	Amount    string `json:"amount"`
	TokenId   string `json:"tokenId"`
	TokenURI  string `json:"tokenURI"`
}

func RegisterRoutes(r *gin.Engine) {
	r.POST("/mint", mintHandler)
	r.POST("/transfer", transferHandler)
	r.GET("/balance", balanceHandler)
	r.GET("/all-logs", allLogsHandler)
}

func mintHandler(c *gin.Context) {
	var req MintRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		db.InsertAPILog("/mint", "Invalid JSON", `{"error":"invalid request"}`, 400)
		return
	}

	requestPayload := utils.StringifyJSON(req)

	client := utils.GetClient()
	auth := utils.GetTransactOpts(client)
	to := common.HexToAddress(req.ToAddress)

	var txHash string
	var responsePayload string
	var statusCode int

	switch req.TokenType {
	case "erc20":
		contract, err := contracts.NewERC20(common.HexToAddress(os.Getenv("ERC20_ADDRESS")), client)
		if err != nil {
			responsePayload = `{"error":"erc20 init failed"}`
			statusCode = 500
			c.JSON(statusCode, gin.H{"error": "erc20 init failed"})
			break
		}

		amount, ok := new(big.Int).SetString(req.Amount, 10)
		if !ok {
			responsePayload = `{"error":"invalid amount"}`
			statusCode = 400
			c.JSON(statusCode, gin.H{"error": "invalid amount"})
			break
		}

		tx, err := contract.Mint(auth, to, amount)
		if err != nil {
			responsePayload = fmt.Sprintf(`{"error":"%v"}`, err.Error())
			statusCode = 500
			c.JSON(statusCode, gin.H{"error": err.Error()})
			break
		}
		txHash = tx.Hash().Hex()
		responsePayload = fmt.Sprintf(`{"txHash":"%s"}`, txHash)
		statusCode = 200
		c.JSON(statusCode, gin.H{"txHash": txHash})

	case "erc721":
		contract, err := contracts.NewERC721(common.HexToAddress(os.Getenv("ERC721_ADDRESS")), client)
		if err != nil {
			responsePayload = `{"error":"erc721 init failed"}`
			statusCode = 500
			c.JSON(statusCode, gin.H{"error": "erc721 init failed"})
			break
		}

		tx, err := contract.Mint(auth, to, req.TokenURI)
		if err != nil {
			responsePayload = fmt.Sprintf(`{"error":"%v"}`, err.Error())
			statusCode = 500
			c.JSON(statusCode, gin.H{"error": err.Error()})
			break
		}
		txHash = tx.Hash().Hex()
		responsePayload = fmt.Sprintf(`{"txHash":"%s"}`, txHash)
		statusCode = 200
		c.JSON(statusCode, gin.H{"txHash": txHash})

	case "erc1155":
		contract, err := contracts.NewERC1155(common.HexToAddress(os.Getenv("ERC1155_ADDRESS")), client)
		if err != nil {
			responsePayload = `{"error":"erc1155 init failed"}`
			statusCode = 500
			c.JSON(statusCode, gin.H{"error": "erc1155 init failed"})
			break
		}

		amount, _ := new(big.Int).SetString(req.Amount, 10)
		id, _ := new(big.Int).SetString(req.TokenId, 10)

		tx, err := contract.Mint(auth, to, id, amount, req.TokenURI)
		if err != nil {
			responsePayload = fmt.Sprintf(`{"error":"%v"}`, err.Error())
			statusCode = 500
			c.JSON(statusCode, gin.H{"error": err.Error()})
			break
		}
		txHash = tx.Hash().Hex()
		responsePayload = fmt.Sprintf(`{"txHash":"%s"}`, txHash)
		statusCode = 200
		c.JSON(statusCode, gin.H{"txHash": txHash})

	default:
		responsePayload = `{"error":"unknown token type"}`
		statusCode = 400
		c.JSON(statusCode, gin.H{"error": "unknown token type"})
	}

	db.InsertAPILog("/mint", requestPayload, responsePayload, statusCode)
}

func transferHandler(c *gin.Context) {
	var req MintRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		db.InsertAPILog("/transfer", "Invalid JSON", `{"error":"invalid request"}`, 400)
		return
	}

	requestPayload := utils.StringifyJSON(req)

	client := utils.GetClient()
	auth := utils.GetTransactOpts(client)
	to := common.HexToAddress(req.ToAddress)

	var txHash string
	var responsePayload string
	var statusCode int

	switch req.TokenType {
	case "erc20":
		contract, err := contracts.NewERC20(common.HexToAddress(os.Getenv("ERC20_ADDRESS")), client)
		if err != nil {
			responsePayload = `{"error":"erc20 init failed"}`
			statusCode = 500
			c.JSON(statusCode, gin.H{"error": "erc20 init failed"})
			break
		}

		amount, ok := new(big.Int).SetString(req.Amount, 10)
		if !ok {
			responsePayload = `{"error":"invalid amount"}`
			statusCode = 400
			c.JSON(statusCode, gin.H{"error": "invalid amount"})
			break
		}

		tx, err := contract.Transfer(auth, to, amount)
		if err != nil {
			responsePayload = fmt.Sprintf(`{"error":"%v"}`, err.Error())
			statusCode = 500
			c.JSON(statusCode, gin.H{"error": err.Error()})
			break
		}
		txHash = tx.Hash().Hex()
		responsePayload = fmt.Sprintf(`{"txHash":"%s"}`, txHash)
		statusCode = 200
		c.JSON(statusCode, gin.H{"txHash": txHash})

	case "erc721":
		contract, err := contracts.NewERC721(common.HexToAddress(os.Getenv("ERC721_ADDRESS")), client)
		if err != nil {
			responsePayload = `{"error":"erc721 init failed"}`
			statusCode = 500
			c.JSON(statusCode, gin.H{"error": "erc721 init failed"})
			break
		}

		tokenId, _ := new(big.Int).SetString(req.TokenId, 10)

		tx, err := contract.SafeTransferFrom(auth, auth.From, to, tokenId)
		if err != nil {
			responsePayload = fmt.Sprintf(`{"error":"%v"}`, err.Error())
			statusCode = 500
			c.JSON(statusCode, gin.H{"error": err.Error()})
			break
		}
		txHash = tx.Hash().Hex()
		responsePayload = fmt.Sprintf(`{"txHash":"%s"}`, txHash)
		statusCode = 200
		c.JSON(statusCode, gin.H{"txHash": txHash})

	case "erc1155":
		contract, err := contracts.NewERC1155(common.HexToAddress(os.Getenv("ERC1155_ADDRESS")), client)
		if err != nil {
			responsePayload = `{"error":"erc1155 init failed"}`
			statusCode = 500
			c.JSON(statusCode, gin.H{"error": "erc1155 init failed"})
			break
		}

		id, _ := new(big.Int).SetString(req.TokenId, 10)
		amount, _ := new(big.Int).SetString(req.Amount, 10)
		data := []byte{}

		tx, err := contract.SafeTransferFrom(auth, auth.From, to, id, amount, data)
		if err != nil {
			responsePayload = fmt.Sprintf(`{"error":"%v"}`, err.Error())
			statusCode = 500
			c.JSON(statusCode, gin.H{"error": err.Error()})
			break
		}
		txHash = tx.Hash().Hex()
		responsePayload = fmt.Sprintf(`{"txHash":"%s"}`, txHash)
		statusCode = 200
		c.JSON(statusCode, gin.H{"txHash": txHash})

	default:
		responsePayload = `{"error":"unknown token type"}`
		statusCode = 400
		c.JSON(statusCode, gin.H{"error": "unknown token type"})
	}

	db.InsertAPILog("/transfer", requestPayload, responsePayload, statusCode)
}

func balanceHandler(c *gin.Context) {
	tokenType := c.Query("tokenType")
	address := c.Query("walletAddress")

	requestPayload := fmt.Sprintf(`{"tokenType":"%s", "walletAddress":"%s"}`, tokenType, address)
	var responsePayload string
	var statusCode int

	if tokenType == "" || address == "" {
		responsePayload = `{"error":"tokenType and walletAddress are required"}`
		statusCode = 400
		c.JSON(statusCode, gin.H{"error": "tokenType and walletAddress are required"})
		db.InsertAPILog("/balance", requestPayload, responsePayload, statusCode)
		return
	}

	client := utils.GetClient()
	user := common.HexToAddress(address)

	switch tokenType {
	case "erc20":
		contract, err := contracts.NewERC20(common.HexToAddress(os.Getenv("ERC20_ADDRESS")), client)
		if err != nil {
			responsePayload = `{"error":"erc20 init failed"}`
			statusCode = 500
			c.JSON(statusCode, gin.H{"error": "erc20 init failed"})
			break
		}

		balance, err := contract.BalanceOf(nil, user)
		if err != nil {
			responsePayload = fmt.Sprintf(`{"error":"%v"}`, err.Error())
			statusCode = 500
			c.JSON(statusCode, gin.H{"error": err.Error()})
			break
		}

		responsePayload = fmt.Sprintf(`{"balance":"%s"}`, balance.String())
		statusCode = 200
		c.JSON(statusCode, gin.H{"balance": balance.String()})

	case "erc721":
		contract, err := contracts.NewERC721(common.HexToAddress(os.Getenv("ERC721_ADDRESS")), client)
		if err != nil {
			responsePayload = `{"error":"erc721 init failed"}`
			statusCode = 500
			c.JSON(statusCode, gin.H{"error": "erc721 init failed"})
			break
		}

		count, err := contract.BalanceOf(nil, user)
		if err != nil {
			responsePayload = fmt.Sprintf(`{"error":"%v"}`, err.Error())
			statusCode = 500
			c.JSON(statusCode, gin.H{"error": err.Error()})
			break
		}

		responsePayload = fmt.Sprintf(`{"ownedCount":"%s"}`, count.String())
		statusCode = 200
		c.JSON(statusCode, gin.H{"ownedCount": count.String()})

	case "erc1155":
		tokenId := c.Query("tokenId")
		if tokenId == "" {
			responsePayload = `{"error":"tokenId is required for ERC1155"}`
			statusCode = 400
			c.JSON(statusCode, gin.H{"error": "tokenId is required for ERC1155"})
			break
		}

		id, _ := new(big.Int).SetString(tokenId, 10)
		contract, err := contracts.NewERC1155(common.HexToAddress(os.Getenv("ERC1155_ADDRESS")), client)
		if err != nil {
			responsePayload = `{"error":"erc1155 init failed"}`
			statusCode = 500
			c.JSON(statusCode, gin.H{"error": "erc1155 init failed"})
			break
		}

		balance, err := contract.BalanceOf(nil, user, id)
		if err != nil {
			responsePayload = fmt.Sprintf(`{"error":"%v"}`, err.Error())
			statusCode = 500
			c.JSON(statusCode, gin.H{"error": err.Error()})
			break
		}

		responsePayload = fmt.Sprintf(`{"tokenId":"%s","balance":"%s"}`, tokenId, balance.String())
		statusCode = 200
		c.JSON(statusCode, gin.H{"tokenId": tokenId, "balance": balance.String()})

	default:
		responsePayload = `{"error":"unsupported tokenType"}`
		statusCode = 400
		c.JSON(statusCode, gin.H{"error": "unsupported tokenType"})
	}

	db.InsertAPILog("/balance", requestPayload, responsePayload, statusCode)
}

func allLogsHandler(c *gin.Context) {
	logs := db.GetAllLogs()
	responsePayload := utils.StringifyJSON(logs)
	db.InsertAPILog("/all-logs", "{}", responsePayload, 200)
	c.JSON(200, gin.H{"message": logs})
}
