package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/big"
	"strings"

	token "erc20/contracts"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	bun "github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"

	"github.com/jon4hz/ethconvert/pkg/ethconvert"
	"github.com/shopspring/decimal"
)

// LogTransfer ..
type LogTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
}

// LogApproval ..
type LogApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
}

func main() {
	// Set Dial
	db, ctx := ConnectDB(), Ctx()

	client, err := ethclient.Dial("wss://speedy-nodes-nyc.moralis.io/f1b8d020016effdedbe0d713/bsc/testnet/archive/ws")
	if err != nil {
		log.Fatal(err)
	}

	// 0x Protocol (BUSD) token address
	contractAddress := common.HexToAddress("0x665aE6c8B332cCE9B1B50d9B2c79d1731516d2fB")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs := make(chan types.Log)

	// Subscribe Filter
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	// Contract ABI
	contractAbi, err := abi.JSON(strings.NewReader(string(token.TokenABI)))
	if err != nil {
		log.Fatal(err)
	}

	logTransferSig := []byte("Transfer(address,address,uint256)")
	LogApprovalSig := []byte("Approval(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
	logApprovalSigHash := crypto.Keccak256Hash(LogApprovalSig)

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			switch vLog.Topics[0].Hex() {
			case logTransferSigHash.Hex():

				var transferEvent LogTransfer

				err := contractAbi.UnpackIntoInterface(&transferEvent, "Transfer", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}

				transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
				transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())

				// Convert BigInt to Decimal
				tokenInDecimal := decimal.NewFromBigInt(transferEvent.Tokens, 0)
				// Convert from Wei to Ether
				tokenInEther, _ := ethconvert.Convert(tokenInDecimal, "wei", "ether")

				fmt.Printf("Block Number: %v\n", vLog.BlockNumber)
				fmt.Printf("From: %s\n", transferEvent.From.Hex())
				fmt.Printf("To: %s\n", transferEvent.To.Hex())
				fmt.Printf("Tokens(Wei): %s\n", transferEvent.Tokens.String())
				fmt.Println("Tokens:", tokenInEther)

				TransferValues := map[string]interface{}{
					"block_number": vLog.BlockNumber,
					"from":         transferEvent.From.Hex(),
					"to":           transferEvent.To.Hex(),
					"tokens":       tokenInEther,
				}
				InsertData(db, ctx, TransferValues)

			case logApprovalSigHash.Hex():
				fmt.Printf("Log Name: Approval\n")

				var approvalEvent LogApproval

				err := contractAbi.UnpackIntoInterface(&approvalEvent, "Approval", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}

				// Convert BigInt to Decimal
				tokenInDecimal := decimal.NewFromBigInt(approvalEvent.Tokens, 0)
				// Convert from Wei to Ether
				tokenInEther, _ := ethconvert.Convert(tokenInDecimal, "wei", "ether")

				approvalEvent.TokenOwner = common.HexToAddress(vLog.Topics[1].Hex())
				approvalEvent.Spender = common.HexToAddress(vLog.Topics[2].Hex())

				fmt.Printf("Block Number: %v\n", vLog.BlockNumber)
				fmt.Printf("Token Owner: %s\n", approvalEvent.TokenOwner.Hex())
				fmt.Printf("Spender: %s\n", approvalEvent.Spender.Hex())
				fmt.Printf("Tokens (Wei): %s\n", approvalEvent.Tokens.String())
				fmt.Println("Tokens:", tokenInEther)

				TransferValues := map[string]interface{}{
					"block_number": vLog.BlockNumber,
					"from":         approvalEvent.TokenOwner.Hex(),
					"to":           approvalEvent.Spender.Hex(),
					"tokens":       tokenInEther,
				}
				InsertData(db, ctx, TransferValues)

			}

		}

		fmt.Printf("\n")

	}

}

func InsertData(db *bun.DB, ctx context.Context, v map[string]interface{}) {
	_, err := db.NewInsert().Model(&v).TableExpr("transfer_log").Exec(ctx)
	if err != nil {
		panic(err)
	}
}

func Ctx() context.Context {
	ctx := context.Background()
	return ctx
}

func ConnectDB() *bun.DB {

	// Open a PostgreSQL database.
	dsn := "postgres://postgres:mysecretpassword@localhost:5432/test2?sslmode=disable"
	pgdb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	// Create a Bun db on top of it.
	db := bun.NewDB(pgdb, pgdialect.New())

	// Print all queries to stdout.
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	return db
}

type Transfer struct {
	bun.BaseModel `bun:"borrow,alias:u"`

	block_number int    `json:"block_number"`
	from         string `json:"from"`
	to           string `json:"to"`
	tokens       string `json:"tokens"`
}

// test
