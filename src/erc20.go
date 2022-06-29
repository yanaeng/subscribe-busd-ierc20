package src

import (
	"context"
	token "erc20/contracts"
	model "erc20/model"

	"fmt"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jon4hz/ethconvert/pkg/ethconvert"
	"github.com/shopspring/decimal"
)

func InsertToDb() {

	// connect_db
	db, ctx := ConnectDB(), Ctx()

	// Set Dial
	client, err := ethclient.Dial("wss://speedy-nodes-nyc.moralis.io/f1b8d020016effdedbe0d713/bsc/testnet/archive/ws")
	if err != nil {
		log.Fatal(err)
	}

	// 0x Protocol (BUSD) token address
	address := []common.Address{common.HexToAddress("0x665aE6c8B332cCE9B1B50d9B2c79d1731516d2fB"), common.HexToAddress("0x2F02f77c2bA5A7cE4924f2a1E5Ecc85580fDD096"), common.HexToAddress("0xae13d989daC2f0dEbFf460aC112a837C89BAa7cd"), common.HexToAddress("0xBb5f03Ea8Bee3cF0F646C5260115fb5a16Ca4Ae4")}
	// contractAddressBusd := common.HexToAddress("0x665aE6c8B332cCE9B1B50d9B2c79d1731516d2fB")
	// contractAddressUsdt := common.HexToAddress("0x2F02f77c2bA5A7cE4924f2a1E5Ecc85580fDD096")

	query := ethereum.FilterQuery{
		Addresses: address,
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

				var transferEvent model.LogTransfer

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
				TableHead := "transfer_log"

				TransferValues := map[string]interface{}{
					"block_number": vLog.BlockNumber,
					"from":         transferEvent.From.Hex(),
					"to":           transferEvent.To.Hex(),
					"tokens":       transferEvent.Tokens.String(),
				}
				InsertData(db, ctx, TableHead, TransferValues)

			case logApprovalSigHash.Hex():
				fmt.Printf("Log Name: Approval\n")

				var approvalEvent model.LogApproval

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
				TableHead := "approve"
				TransferValues := map[string]interface{}{
					"block_number": vLog.BlockNumber,
					"from":         approvalEvent.TokenOwner.Hex(),
					"to":           approvalEvent.Spender.Hex(),
					"tokens":       approvalEvent.Tokens.String(),
				}
				InsertData(db, ctx, TableHead, TransferValues)

			}

		}

		fmt.Printf("\n")

	}

}
