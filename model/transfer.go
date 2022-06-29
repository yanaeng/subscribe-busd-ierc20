package model

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// import "github.com/uptrace/bun"

// type Transfer struct {
// 	bun.BaseModel `bun:"borrow,alias:u"`

// 	block_number int    `json:"block_number"`
// 	from         string `json:"from"`
// 	to           string `json:"to"`
// 	tokens       string `json:"tokens"`
// }

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
