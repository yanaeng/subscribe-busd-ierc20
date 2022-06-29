package src

// import (
// 	token "erc20/contracts"
// 	"fmt"
// 	"log"

// 	"github.com/ethereum/go-ethereum/accounts/abi/bind"
// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/ethclient"
// 	// "github.com/ethereum/go-ethereum/"
// )

// func GetNameCoins(client *ethclient.Client, coin common.Address) (string, string) {
// 	// tokenAddress := common.HexToAddress("0xa74476443119A942dE498590Fe1f2454d7D4aC0d")
// 	instance, err := token.NewToken(coin, client)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// address := common.HexToAddress("0x0536806df512d6cdde913cf95c9886f65b1d3462")
// 	// bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
// 	// if err != nil {
// 	//     log.Fatal(err)
// 	// }

// 	name, err := instance.Name(&bind.CallOpts{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	symbol, err := instance.Symbol(&bind.CallOpts{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	decimals, err := instance.Decimals(&bind.CallOpts{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Printf("name: %s\n", name)     // "name: Golem Network"
// 	fmt.Printf("symbol: %s\n", symbol) // "symbol: GNT"
// 	fmt.Printf("decimals: %v\n", decimals)
// 	return name, symbol
// }
