// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// TokenMetaData contains all meta data concerning the Token contract.
var TokenMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"oldRank\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"newRank\",\"type\":\"uint8\"}],\"name\":\"ActivateRank\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"loanId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isAdd\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralAdjustAmount\",\"type\":\"uint256\"}],\"name\":\"AdjustCollateral\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"ApprovedForRouter\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldDuration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDuration\",\"type\":\"uint256\"}],\"name\":\"AvailablePeriodChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"loanId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrowTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collateralTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"owedPerDay\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minInterest\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"newLoan\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"rolloverTimestamp\",\"type\":\"uint64\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"burner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"BurnIfpToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"burner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"BurnItpToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"burner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BurnPToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"CallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"predecessor\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"CallScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"Cancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interestForwClaimed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interestForwBonus\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"burnedIfp\",\"type\":\"uint256\"}],\"name\":\"ClaimForwInterest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interestTokenClaimed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interestTokenBonus\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"burnedItp\",\"type\":\"uint256\"}],\"name\":\"ClaimTokenInterest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"loanId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowPaid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interestPaid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralAmountwithdraw\",\"type\":\"uint256\"}],\"name\":\"CloseLoan\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintedP\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintedItp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintedIfp\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"}],\"name\":\"DeprecateStakeInfo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"FaucetTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"}],\"name\":\"GlobalPricingPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"coreAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"interestVaultAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"membershipAddress\",\"type\":\"address\"}],\"name\":\"Initialize\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"loanId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"swapPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmountFromSwap\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bountyReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bountyRewardTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenSentBackToUser\",\"type\":\"uint256\"}],\"name\":\"Liquidate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldDuration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDuration\",\"type\":\"uint256\"}],\"name\":\"MinDelayChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldDuration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDuration\",\"type\":\"uint256\"}],\"name\":\"MinProxyAdminDelayChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldDuration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDuration\",\"type\":\"uint256\"}],\"name\":\"MinUrgentDelayChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"MintIfpToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"MintItpToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MintPToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"forwAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"OwnerApprove\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldProxyAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newProxyAdmin\",\"type\":\"address\"}],\"name\":\"ProxyAdminChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"}],\"name\":\"RegisterNewPool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"loanId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"closeLoan\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowPaid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interestPaid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralAmountwithdraw\",\"type\":\"uint256\"}],\"name\":\"Repay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"loanId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bountyHunter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delayInterest\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bountyReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bountyRewardTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newInterestOwedPerDay\",\"type\":\"uint256\"}],\"name\":\"Rollover\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"SetAdvancedInterestDuration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"rates\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"utils\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"targetSupply\",\"type\":\"uint256\"}],\"name\":\"SetBorrowInterestParams\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldValue\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"SetCoreBorrowingAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"SetDecimals\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldValue\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"SetFeeController\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"SetFeeSpread\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldValue\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"SetForwAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldValue\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"SetForwDistributorAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"targetBlock\",\"type\":\"uint256\"}],\"name\":\"SetFowPerBlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collateralTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"safeLTV\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxLTV\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liqLTV\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bountyFeeRate\",\"type\":\"uint256\"}],\"name\":\"SetLoanConfig\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"SetLoanDuration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldValue\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"SetMembershipAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldValue\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"SetPoolBorrowingAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldValue\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"SetPoolLendingAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"SetPoolStartTimestamp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"feeds\",\"type\":\"address[]\"}],\"name\":\"SetPriceFeed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldValue\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"SetPriceFeedAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldValue\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"SetProtocolAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldValue\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"SetRouterAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldValue\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"SetTokenAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldValue\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"SetWETHHandler\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"coreAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"interestVaultAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"forwDistributionAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"forwTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SettleForwInterest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimableTokenInterest\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"heldTokenInterest\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimableForwInterest\",\"type\":\"uint256\"}],\"name\":\"SettleInterest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrowTokenAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldSafeLTV\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldMaxLTV\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldLiquidationLTV\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldBountyFeeRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newSafeLTV\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newLMaxLTV\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newLiquidationLTV\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBountyFeeRate\",\"type\":\"uint256\"}],\"name\":\"SetupLoanConfig\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Stake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UnStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"urgentFunction\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"set\",\"type\":\"bool\"}],\"name\":\"UrgentFunctionUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"burnedP\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"burnedItp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"burnedIfp\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"profitWithdraw\",\"type\":\"uint256\"}],\"name\":\"WithdrawActualProfit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimable\",\"type\":\"uint256\"}],\"name\":\"WithdrawForwInterest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimable\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bonus\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"profit\",\"type\":\"uint256\"}],\"name\":\"WithdrawTokenInterest\",\"type\":\"event\"}]",
}

// TokenABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenMetaData.ABI instead.
var TokenABI = TokenMetaData.ABI

// Token is an auto generated Go binding around an Ethereum contract.
type Token struct {
	TokenCaller     // Read-only binding to the contract
	TokenTransactor // Write-only binding to the contract
	TokenFilterer   // Log filterer for contract events
}

// TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSession struct {
	Contract     *Token            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallerSession struct {
	Contract *TokenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransactorSession struct {
	Contract     *TokenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRaw struct {
	Contract *Token // Generic contract binding to access the raw methods on
}

// TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallerRaw struct {
	Contract *TokenCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransactorRaw struct {
	Contract *TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// NewTokenCaller creates a new read-only instance of Token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{contract: contract}, nil
}

// NewTokenTransactor creates a new write-only instance of Token, bound to a specific deployed contract.
func NewTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransactor, error) {
	contract, err := bindToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTransactor{contract: contract}, nil
}

// NewTokenFilterer creates a new log filterer instance of Token, bound to a specific deployed contract.
func NewTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenFilterer, error) {
	contract, err := bindToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFilterer{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Token.Contract.TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.contract.Transact(opts, method, params...)
}

// TokenActivateRankIterator is returned from FilterActivateRank and is used to iterate over the raw logs and unpacked data for ActivateRank events raised by the Token contract.
type TokenActivateRankIterator struct {
	Event *TokenActivateRank // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenActivateRankIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenActivateRank)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenActivateRank)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenActivateRankIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenActivateRankIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenActivateRank represents a ActivateRank event raised by the Token contract.
type TokenActivateRank struct {
	Owner   common.Address
	NftId   *big.Int
	OldRank uint8
	NewRank uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterActivateRank is a free log retrieval operation binding the contract event 0xf62d2ef9723120c1b43dc6ee49b2d6673ebcb929fa1e7027086843236c4d6ef1.
//
// Solidity: event ActivateRank(address indexed owner, uint256 indexed nftId, uint8 oldRank, uint8 newRank)
func (_Token *TokenFilterer) FilterActivateRank(opts *bind.FilterOpts, owner []common.Address, nftId []*big.Int) (*TokenActivateRankIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "ActivateRank", ownerRule, nftIdRule)
	if err != nil {
		return nil, err
	}
	return &TokenActivateRankIterator{contract: _Token.contract, event: "ActivateRank", logs: logs, sub: sub}, nil
}

// WatchActivateRank is a free log subscription operation binding the contract event 0xf62d2ef9723120c1b43dc6ee49b2d6673ebcb929fa1e7027086843236c4d6ef1.
//
// Solidity: event ActivateRank(address indexed owner, uint256 indexed nftId, uint8 oldRank, uint8 newRank)
func (_Token *TokenFilterer) WatchActivateRank(opts *bind.WatchOpts, sink chan<- *TokenActivateRank, owner []common.Address, nftId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "ActivateRank", ownerRule, nftIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenActivateRank)
				if err := _Token.contract.UnpackLog(event, "ActivateRank", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseActivateRank is a log parse operation binding the contract event 0xf62d2ef9723120c1b43dc6ee49b2d6673ebcb929fa1e7027086843236c4d6ef1.
//
// Solidity: event ActivateRank(address indexed owner, uint256 indexed nftId, uint8 oldRank, uint8 newRank)
func (_Token *TokenFilterer) ParseActivateRank(log types.Log) (*TokenActivateRank, error) {
	event := new(TokenActivateRank)
	if err := _Token.contract.UnpackLog(event, "ActivateRank", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenAdjustCollateralIterator is returned from FilterAdjustCollateral and is used to iterate over the raw logs and unpacked data for AdjustCollateral events raised by the Token contract.
type TokenAdjustCollateralIterator struct {
	Event *TokenAdjustCollateral // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenAdjustCollateralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenAdjustCollateral)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenAdjustCollateral)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenAdjustCollateralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenAdjustCollateralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenAdjustCollateral represents a AdjustCollateral event raised by the Token contract.
type TokenAdjustCollateral struct {
	Owner                  common.Address
	NftId                  *big.Int
	LoanId                 *big.Int
	IsAdd                  bool
	CollateralAdjustAmount *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterAdjustCollateral is a free log retrieval operation binding the contract event 0x20728dd89b8e8dacc7d7539dce3bd05d87c4174dbee8f921d4a6e12c0e9a3ae9.
//
// Solidity: event AdjustCollateral(address indexed owner, uint256 indexed nftId, uint256 indexed loanId, bool isAdd, uint256 collateralAdjustAmount)
func (_Token *TokenFilterer) FilterAdjustCollateral(opts *bind.FilterOpts, owner []common.Address, nftId []*big.Int, loanId []*big.Int) (*TokenAdjustCollateralIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}
	var loanIdRule []interface{}
	for _, loanIdItem := range loanId {
		loanIdRule = append(loanIdRule, loanIdItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "AdjustCollateral", ownerRule, nftIdRule, loanIdRule)
	if err != nil {
		return nil, err
	}
	return &TokenAdjustCollateralIterator{contract: _Token.contract, event: "AdjustCollateral", logs: logs, sub: sub}, nil
}

// WatchAdjustCollateral is a free log subscription operation binding the contract event 0x20728dd89b8e8dacc7d7539dce3bd05d87c4174dbee8f921d4a6e12c0e9a3ae9.
//
// Solidity: event AdjustCollateral(address indexed owner, uint256 indexed nftId, uint256 indexed loanId, bool isAdd, uint256 collateralAdjustAmount)
func (_Token *TokenFilterer) WatchAdjustCollateral(opts *bind.WatchOpts, sink chan<- *TokenAdjustCollateral, owner []common.Address, nftId []*big.Int, loanId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}
	var loanIdRule []interface{}
	for _, loanIdItem := range loanId {
		loanIdRule = append(loanIdRule, loanIdItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "AdjustCollateral", ownerRule, nftIdRule, loanIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenAdjustCollateral)
				if err := _Token.contract.UnpackLog(event, "AdjustCollateral", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAdjustCollateral is a log parse operation binding the contract event 0x20728dd89b8e8dacc7d7539dce3bd05d87c4174dbee8f921d4a6e12c0e9a3ae9.
//
// Solidity: event AdjustCollateral(address indexed owner, uint256 indexed nftId, uint256 indexed loanId, bool isAdd, uint256 collateralAdjustAmount)
func (_Token *TokenFilterer) ParseAdjustCollateral(log types.Log) (*TokenAdjustCollateral, error) {
	event := new(TokenAdjustCollateral)
	if err := _Token.contract.UnpackLog(event, "AdjustCollateral", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenApprovedForRouterIterator is returned from FilterApprovedForRouter and is used to iterate over the raw logs and unpacked data for ApprovedForRouter events raised by the Token contract.
type TokenApprovedForRouterIterator struct {
	Event *TokenApprovedForRouter // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenApprovedForRouterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenApprovedForRouter)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenApprovedForRouter)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenApprovedForRouterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenApprovedForRouterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenApprovedForRouter represents a ApprovedForRouter event raised by the Token contract.
type TokenApprovedForRouter struct {
	Sender common.Address
	Asset  common.Address
	Router common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterApprovedForRouter is a free log retrieval operation binding the contract event 0x0655f25cdc029feb280ae84f7e17f7920ffd72612c8817e4bc67d53ab2a0fd68.
//
// Solidity: event ApprovedForRouter(address indexed sender, address asset, address router)
func (_Token *TokenFilterer) FilterApprovedForRouter(opts *bind.FilterOpts, sender []common.Address) (*TokenApprovedForRouterIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "ApprovedForRouter", senderRule)
	if err != nil {
		return nil, err
	}
	return &TokenApprovedForRouterIterator{contract: _Token.contract, event: "ApprovedForRouter", logs: logs, sub: sub}, nil
}

// WatchApprovedForRouter is a free log subscription operation binding the contract event 0x0655f25cdc029feb280ae84f7e17f7920ffd72612c8817e4bc67d53ab2a0fd68.
//
// Solidity: event ApprovedForRouter(address indexed sender, address asset, address router)
func (_Token *TokenFilterer) WatchApprovedForRouter(opts *bind.WatchOpts, sink chan<- *TokenApprovedForRouter, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "ApprovedForRouter", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenApprovedForRouter)
				if err := _Token.contract.UnpackLog(event, "ApprovedForRouter", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovedForRouter is a log parse operation binding the contract event 0x0655f25cdc029feb280ae84f7e17f7920ffd72612c8817e4bc67d53ab2a0fd68.
//
// Solidity: event ApprovedForRouter(address indexed sender, address asset, address router)
func (_Token *TokenFilterer) ParseApprovedForRouter(log types.Log) (*TokenApprovedForRouter, error) {
	event := new(TokenApprovedForRouter)
	if err := _Token.contract.UnpackLog(event, "ApprovedForRouter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenAvailablePeriodChangeIterator is returned from FilterAvailablePeriodChange and is used to iterate over the raw logs and unpacked data for AvailablePeriodChange events raised by the Token contract.
type TokenAvailablePeriodChangeIterator struct {
	Event *TokenAvailablePeriodChange // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenAvailablePeriodChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenAvailablePeriodChange)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenAvailablePeriodChange)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenAvailablePeriodChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenAvailablePeriodChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenAvailablePeriodChange represents a AvailablePeriodChange event raised by the Token contract.
type TokenAvailablePeriodChange struct {
	OldDuration *big.Int
	NewDuration *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAvailablePeriodChange is a free log retrieval operation binding the contract event 0x4615d5aa9c478ed77feb40a3514b2515d90f82aa2626c15d97ebfa8c4add984b.
//
// Solidity: event AvailablePeriodChange(uint256 oldDuration, uint256 newDuration)
func (_Token *TokenFilterer) FilterAvailablePeriodChange(opts *bind.FilterOpts) (*TokenAvailablePeriodChangeIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "AvailablePeriodChange")
	if err != nil {
		return nil, err
	}
	return &TokenAvailablePeriodChangeIterator{contract: _Token.contract, event: "AvailablePeriodChange", logs: logs, sub: sub}, nil
}

// WatchAvailablePeriodChange is a free log subscription operation binding the contract event 0x4615d5aa9c478ed77feb40a3514b2515d90f82aa2626c15d97ebfa8c4add984b.
//
// Solidity: event AvailablePeriodChange(uint256 oldDuration, uint256 newDuration)
func (_Token *TokenFilterer) WatchAvailablePeriodChange(opts *bind.WatchOpts, sink chan<- *TokenAvailablePeriodChange) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "AvailablePeriodChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenAvailablePeriodChange)
				if err := _Token.contract.UnpackLog(event, "AvailablePeriodChange", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAvailablePeriodChange is a log parse operation binding the contract event 0x4615d5aa9c478ed77feb40a3514b2515d90f82aa2626c15d97ebfa8c4add984b.
//
// Solidity: event AvailablePeriodChange(uint256 oldDuration, uint256 newDuration)
func (_Token *TokenFilterer) ParseAvailablePeriodChange(log types.Log) (*TokenAvailablePeriodChange, error) {
	event := new(TokenAvailablePeriodChange)
	if err := _Token.contract.UnpackLog(event, "AvailablePeriodChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the Token contract.
type TokenBorrowIterator struct {
	Event *TokenBorrow // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenBorrow)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenBorrow)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenBorrow represents a Borrow event raised by the Token contract.
type TokenBorrow struct {
	Owner                  common.Address
	NftId                  *big.Int
	LoanId                 *big.Int
	BorrowTokenAddress     common.Address
	CollateralTokenAddress common.Address
	BorrowAmount           *big.Int
	CollateralAmount       *big.Int
	OwedPerDay             *big.Int
	MinInterest            *big.Int
	NewLoan                uint8
	RolloverTimestamp      uint64
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0x7f4e96a700fa57f9743854f3f859d60cc80e55aaa986732b6e86a8b350f0483a.
//
// Solidity: event Borrow(address indexed owner, uint256 indexed nftId, uint256 indexed loanId, address borrowTokenAddress, address collateralTokenAddress, uint256 borrowAmount, uint256 collateralAmount, uint256 owedPerDay, uint256 minInterest, uint8 newLoan, uint64 rolloverTimestamp)
func (_Token *TokenFilterer) FilterBorrow(opts *bind.FilterOpts, owner []common.Address, nftId []*big.Int, loanId []*big.Int) (*TokenBorrowIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}
	var loanIdRule []interface{}
	for _, loanIdItem := range loanId {
		loanIdRule = append(loanIdRule, loanIdItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Borrow", ownerRule, nftIdRule, loanIdRule)
	if err != nil {
		return nil, err
	}
	return &TokenBorrowIterator{contract: _Token.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0x7f4e96a700fa57f9743854f3f859d60cc80e55aaa986732b6e86a8b350f0483a.
//
// Solidity: event Borrow(address indexed owner, uint256 indexed nftId, uint256 indexed loanId, address borrowTokenAddress, address collateralTokenAddress, uint256 borrowAmount, uint256 collateralAmount, uint256 owedPerDay, uint256 minInterest, uint8 newLoan, uint64 rolloverTimestamp)
func (_Token *TokenFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *TokenBorrow, owner []common.Address, nftId []*big.Int, loanId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}
	var loanIdRule []interface{}
	for _, loanIdItem := range loanId {
		loanIdRule = append(loanIdRule, loanIdItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Borrow", ownerRule, nftIdRule, loanIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenBorrow)
				if err := _Token.contract.UnpackLog(event, "Borrow", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBorrow is a log parse operation binding the contract event 0x7f4e96a700fa57f9743854f3f859d60cc80e55aaa986732b6e86a8b350f0483a.
//
// Solidity: event Borrow(address indexed owner, uint256 indexed nftId, uint256 indexed loanId, address borrowTokenAddress, address collateralTokenAddress, uint256 borrowAmount, uint256 collateralAmount, uint256 owedPerDay, uint256 minInterest, uint8 newLoan, uint64 rolloverTimestamp)
func (_Token *TokenFilterer) ParseBorrow(log types.Log) (*TokenBorrow, error) {
	event := new(TokenBorrow)
	if err := _Token.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenBurnIfpTokenIterator is returned from FilterBurnIfpToken and is used to iterate over the raw logs and unpacked data for BurnIfpToken events raised by the Token contract.
type TokenBurnIfpTokenIterator struct {
	Event *TokenBurnIfpToken // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenBurnIfpTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenBurnIfpToken)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenBurnIfpToken)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenBurnIfpTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenBurnIfpTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenBurnIfpToken represents a BurnIfpToken event raised by the Token contract.
type TokenBurnIfpToken struct {
	Burner common.Address
	NftId  *big.Int
	Amount *big.Int
	Price  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurnIfpToken is a free log retrieval operation binding the contract event 0x2f44022f379bc2af79f4a7c6e725279a309ebd13c16e4bb1c849b45144bdfa59.
//
// Solidity: event BurnIfpToken(address indexed burner, uint256 indexed nftId, uint256 amount, uint256 price)
func (_Token *TokenFilterer) FilterBurnIfpToken(opts *bind.FilterOpts, burner []common.Address, nftId []*big.Int) (*TokenBurnIfpTokenIterator, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "BurnIfpToken", burnerRule, nftIdRule)
	if err != nil {
		return nil, err
	}
	return &TokenBurnIfpTokenIterator{contract: _Token.contract, event: "BurnIfpToken", logs: logs, sub: sub}, nil
}

// WatchBurnIfpToken is a free log subscription operation binding the contract event 0x2f44022f379bc2af79f4a7c6e725279a309ebd13c16e4bb1c849b45144bdfa59.
//
// Solidity: event BurnIfpToken(address indexed burner, uint256 indexed nftId, uint256 amount, uint256 price)
func (_Token *TokenFilterer) WatchBurnIfpToken(opts *bind.WatchOpts, sink chan<- *TokenBurnIfpToken, burner []common.Address, nftId []*big.Int) (event.Subscription, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "BurnIfpToken", burnerRule, nftIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenBurnIfpToken)
				if err := _Token.contract.UnpackLog(event, "BurnIfpToken", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBurnIfpToken is a log parse operation binding the contract event 0x2f44022f379bc2af79f4a7c6e725279a309ebd13c16e4bb1c849b45144bdfa59.
//
// Solidity: event BurnIfpToken(address indexed burner, uint256 indexed nftId, uint256 amount, uint256 price)
func (_Token *TokenFilterer) ParseBurnIfpToken(log types.Log) (*TokenBurnIfpToken, error) {
	event := new(TokenBurnIfpToken)
	if err := _Token.contract.UnpackLog(event, "BurnIfpToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenBurnItpTokenIterator is returned from FilterBurnItpToken and is used to iterate over the raw logs and unpacked data for BurnItpToken events raised by the Token contract.
type TokenBurnItpTokenIterator struct {
	Event *TokenBurnItpToken // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenBurnItpTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenBurnItpToken)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenBurnItpToken)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenBurnItpTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenBurnItpTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenBurnItpToken represents a BurnItpToken event raised by the Token contract.
type TokenBurnItpToken struct {
	Burner common.Address
	NftId  *big.Int
	Amount *big.Int
	Price  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurnItpToken is a free log retrieval operation binding the contract event 0x80dea9c95ee9c0c2f35c9bf18f12ab7b88c0567a38ffad4e499dbc2e3854874a.
//
// Solidity: event BurnItpToken(address indexed burner, uint256 indexed nftId, uint256 amount, uint256 price)
func (_Token *TokenFilterer) FilterBurnItpToken(opts *bind.FilterOpts, burner []common.Address, nftId []*big.Int) (*TokenBurnItpTokenIterator, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "BurnItpToken", burnerRule, nftIdRule)
	if err != nil {
		return nil, err
	}
	return &TokenBurnItpTokenIterator{contract: _Token.contract, event: "BurnItpToken", logs: logs, sub: sub}, nil
}

// WatchBurnItpToken is a free log subscription operation binding the contract event 0x80dea9c95ee9c0c2f35c9bf18f12ab7b88c0567a38ffad4e499dbc2e3854874a.
//
// Solidity: event BurnItpToken(address indexed burner, uint256 indexed nftId, uint256 amount, uint256 price)
func (_Token *TokenFilterer) WatchBurnItpToken(opts *bind.WatchOpts, sink chan<- *TokenBurnItpToken, burner []common.Address, nftId []*big.Int) (event.Subscription, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "BurnItpToken", burnerRule, nftIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenBurnItpToken)
				if err := _Token.contract.UnpackLog(event, "BurnItpToken", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBurnItpToken is a log parse operation binding the contract event 0x80dea9c95ee9c0c2f35c9bf18f12ab7b88c0567a38ffad4e499dbc2e3854874a.
//
// Solidity: event BurnItpToken(address indexed burner, uint256 indexed nftId, uint256 amount, uint256 price)
func (_Token *TokenFilterer) ParseBurnItpToken(log types.Log) (*TokenBurnItpToken, error) {
	event := new(TokenBurnItpToken)
	if err := _Token.contract.UnpackLog(event, "BurnItpToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenBurnPTokenIterator is returned from FilterBurnPToken and is used to iterate over the raw logs and unpacked data for BurnPToken events raised by the Token contract.
type TokenBurnPTokenIterator struct {
	Event *TokenBurnPToken // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenBurnPTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenBurnPToken)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenBurnPToken)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenBurnPTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenBurnPTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenBurnPToken represents a BurnPToken event raised by the Token contract.
type TokenBurnPToken struct {
	Burner common.Address
	NftId  *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurnPToken is a free log retrieval operation binding the contract event 0x2ce12fc54f500a8c2b3a6734e919554a5d7e02bf39129905fdcb75ae9ec75c2a.
//
// Solidity: event BurnPToken(address indexed burner, uint256 indexed nftId, uint256 amount)
func (_Token *TokenFilterer) FilterBurnPToken(opts *bind.FilterOpts, burner []common.Address, nftId []*big.Int) (*TokenBurnPTokenIterator, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "BurnPToken", burnerRule, nftIdRule)
	if err != nil {
		return nil, err
	}
	return &TokenBurnPTokenIterator{contract: _Token.contract, event: "BurnPToken", logs: logs, sub: sub}, nil
}

// WatchBurnPToken is a free log subscription operation binding the contract event 0x2ce12fc54f500a8c2b3a6734e919554a5d7e02bf39129905fdcb75ae9ec75c2a.
//
// Solidity: event BurnPToken(address indexed burner, uint256 indexed nftId, uint256 amount)
func (_Token *TokenFilterer) WatchBurnPToken(opts *bind.WatchOpts, sink chan<- *TokenBurnPToken, burner []common.Address, nftId []*big.Int) (event.Subscription, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "BurnPToken", burnerRule, nftIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenBurnPToken)
				if err := _Token.contract.UnpackLog(event, "BurnPToken", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBurnPToken is a log parse operation binding the contract event 0x2ce12fc54f500a8c2b3a6734e919554a5d7e02bf39129905fdcb75ae9ec75c2a.
//
// Solidity: event BurnPToken(address indexed burner, uint256 indexed nftId, uint256 amount)
func (_Token *TokenFilterer) ParseBurnPToken(log types.Log) (*TokenBurnPToken, error) {
	event := new(TokenBurnPToken)
	if err := _Token.contract.UnpackLog(event, "BurnPToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenCallExecutedIterator is returned from FilterCallExecuted and is used to iterate over the raw logs and unpacked data for CallExecuted events raised by the Token contract.
type TokenCallExecutedIterator struct {
	Event *TokenCallExecuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenCallExecuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenCallExecuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenCallExecuted represents a CallExecuted event raised by the Token contract.
type TokenCallExecuted struct {
	Id     [32]byte
	Index  *big.Int
	Target common.Address
	Value  *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCallExecuted is a free log retrieval operation binding the contract event 0xc2617efa69bab66782fa219543714338489c4e9e178271560a91b82c3f612b58.
//
// Solidity: event CallExecuted(bytes32 indexed id, uint256 indexed index, address target, uint256 value, bytes data)
func (_Token *TokenFilterer) FilterCallExecuted(opts *bind.FilterOpts, id [][32]byte, index []*big.Int) (*TokenCallExecutedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "CallExecuted", idRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &TokenCallExecutedIterator{contract: _Token.contract, event: "CallExecuted", logs: logs, sub: sub}, nil
}

// WatchCallExecuted is a free log subscription operation binding the contract event 0xc2617efa69bab66782fa219543714338489c4e9e178271560a91b82c3f612b58.
//
// Solidity: event CallExecuted(bytes32 indexed id, uint256 indexed index, address target, uint256 value, bytes data)
func (_Token *TokenFilterer) WatchCallExecuted(opts *bind.WatchOpts, sink chan<- *TokenCallExecuted, id [][32]byte, index []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "CallExecuted", idRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenCallExecuted)
				if err := _Token.contract.UnpackLog(event, "CallExecuted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCallExecuted is a log parse operation binding the contract event 0xc2617efa69bab66782fa219543714338489c4e9e178271560a91b82c3f612b58.
//
// Solidity: event CallExecuted(bytes32 indexed id, uint256 indexed index, address target, uint256 value, bytes data)
func (_Token *TokenFilterer) ParseCallExecuted(log types.Log) (*TokenCallExecuted, error) {
	event := new(TokenCallExecuted)
	if err := _Token.contract.UnpackLog(event, "CallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenCallScheduledIterator is returned from FilterCallScheduled and is used to iterate over the raw logs and unpacked data for CallScheduled events raised by the Token contract.
type TokenCallScheduledIterator struct {
	Event *TokenCallScheduled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenCallScheduledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenCallScheduled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenCallScheduled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenCallScheduledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenCallScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenCallScheduled represents a CallScheduled event raised by the Token contract.
type TokenCallScheduled struct {
	Id          [32]byte
	Index       *big.Int
	Target      common.Address
	Value       *big.Int
	Data        []byte
	Predecessor [32]byte
	Delay       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCallScheduled is a free log retrieval operation binding the contract event 0x4cf4410cc57040e44862ef0f45f3dd5a5e02db8eb8add648d4b0e236f1d07dca.
//
// Solidity: event CallScheduled(bytes32 indexed id, uint256 indexed index, address target, uint256 value, bytes data, bytes32 predecessor, uint256 delay)
func (_Token *TokenFilterer) FilterCallScheduled(opts *bind.FilterOpts, id [][32]byte, index []*big.Int) (*TokenCallScheduledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "CallScheduled", idRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &TokenCallScheduledIterator{contract: _Token.contract, event: "CallScheduled", logs: logs, sub: sub}, nil
}

// WatchCallScheduled is a free log subscription operation binding the contract event 0x4cf4410cc57040e44862ef0f45f3dd5a5e02db8eb8add648d4b0e236f1d07dca.
//
// Solidity: event CallScheduled(bytes32 indexed id, uint256 indexed index, address target, uint256 value, bytes data, bytes32 predecessor, uint256 delay)
func (_Token *TokenFilterer) WatchCallScheduled(opts *bind.WatchOpts, sink chan<- *TokenCallScheduled, id [][32]byte, index []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "CallScheduled", idRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenCallScheduled)
				if err := _Token.contract.UnpackLog(event, "CallScheduled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCallScheduled is a log parse operation binding the contract event 0x4cf4410cc57040e44862ef0f45f3dd5a5e02db8eb8add648d4b0e236f1d07dca.
//
// Solidity: event CallScheduled(bytes32 indexed id, uint256 indexed index, address target, uint256 value, bytes data, bytes32 predecessor, uint256 delay)
func (_Token *TokenFilterer) ParseCallScheduled(log types.Log) (*TokenCallScheduled, error) {
	event := new(TokenCallScheduled)
	if err := _Token.contract.UnpackLog(event, "CallScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenCancelledIterator is returned from FilterCancelled and is used to iterate over the raw logs and unpacked data for Cancelled events raised by the Token contract.
type TokenCancelledIterator struct {
	Event *TokenCancelled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenCancelled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenCancelled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenCancelled represents a Cancelled event raised by the Token contract.
type TokenCancelled struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCancelled is a free log retrieval operation binding the contract event 0xbaa1eb22f2a492ba1a5fea61b8df4d27c6c8b5f3971e63bb58fa14ff72eedb70.
//
// Solidity: event Cancelled(bytes32 indexed id)
func (_Token *TokenFilterer) FilterCancelled(opts *bind.FilterOpts, id [][32]byte) (*TokenCancelledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Cancelled", idRule)
	if err != nil {
		return nil, err
	}
	return &TokenCancelledIterator{contract: _Token.contract, event: "Cancelled", logs: logs, sub: sub}, nil
}

// WatchCancelled is a free log subscription operation binding the contract event 0xbaa1eb22f2a492ba1a5fea61b8df4d27c6c8b5f3971e63bb58fa14ff72eedb70.
//
// Solidity: event Cancelled(bytes32 indexed id)
func (_Token *TokenFilterer) WatchCancelled(opts *bind.WatchOpts, sink chan<- *TokenCancelled, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Cancelled", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenCancelled)
				if err := _Token.contract.UnpackLog(event, "Cancelled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCancelled is a log parse operation binding the contract event 0xbaa1eb22f2a492ba1a5fea61b8df4d27c6c8b5f3971e63bb58fa14ff72eedb70.
//
// Solidity: event Cancelled(bytes32 indexed id)
func (_Token *TokenFilterer) ParseCancelled(log types.Log) (*TokenCancelled, error) {
	event := new(TokenCancelled)
	if err := _Token.contract.UnpackLog(event, "Cancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenClaimForwInterestIterator is returned from FilterClaimForwInterest and is used to iterate over the raw logs and unpacked data for ClaimForwInterest events raised by the Token contract.
type TokenClaimForwInterestIterator struct {
	Event *TokenClaimForwInterest // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenClaimForwInterestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenClaimForwInterest)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenClaimForwInterest)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenClaimForwInterestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenClaimForwInterestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenClaimForwInterest represents a ClaimForwInterest event raised by the Token contract.
type TokenClaimForwInterest struct {
	Owner               common.Address
	NftId               *big.Int
	InterestForwClaimed *big.Int
	InterestForwBonus   *big.Int
	BurnedIfp           *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterClaimForwInterest is a free log retrieval operation binding the contract event 0x7306e65aafc988ed7e1a7a3546c7a79dace207c76c10db147dc0e3a1a218e0af.
//
// Solidity: event ClaimForwInterest(address indexed owner, uint256 indexed nftId, uint256 interestForwClaimed, uint256 interestForwBonus, uint256 burnedIfp)
func (_Token *TokenFilterer) FilterClaimForwInterest(opts *bind.FilterOpts, owner []common.Address, nftId []*big.Int) (*TokenClaimForwInterestIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "ClaimForwInterest", ownerRule, nftIdRule)
	if err != nil {
		return nil, err
	}
	return &TokenClaimForwInterestIterator{contract: _Token.contract, event: "ClaimForwInterest", logs: logs, sub: sub}, nil
}

// WatchClaimForwInterest is a free log subscription operation binding the contract event 0x7306e65aafc988ed7e1a7a3546c7a79dace207c76c10db147dc0e3a1a218e0af.
//
// Solidity: event ClaimForwInterest(address indexed owner, uint256 indexed nftId, uint256 interestForwClaimed, uint256 interestForwBonus, uint256 burnedIfp)
func (_Token *TokenFilterer) WatchClaimForwInterest(opts *bind.WatchOpts, sink chan<- *TokenClaimForwInterest, owner []common.Address, nftId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "ClaimForwInterest", ownerRule, nftIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenClaimForwInterest)
				if err := _Token.contract.UnpackLog(event, "ClaimForwInterest", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClaimForwInterest is a log parse operation binding the contract event 0x7306e65aafc988ed7e1a7a3546c7a79dace207c76c10db147dc0e3a1a218e0af.
//
// Solidity: event ClaimForwInterest(address indexed owner, uint256 indexed nftId, uint256 interestForwClaimed, uint256 interestForwBonus, uint256 burnedIfp)
func (_Token *TokenFilterer) ParseClaimForwInterest(log types.Log) (*TokenClaimForwInterest, error) {
	event := new(TokenClaimForwInterest)
	if err := _Token.contract.UnpackLog(event, "ClaimForwInterest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenClaimTokenInterestIterator is returned from FilterClaimTokenInterest and is used to iterate over the raw logs and unpacked data for ClaimTokenInterest events raised by the Token contract.
type TokenClaimTokenInterestIterator struct {
	Event *TokenClaimTokenInterest // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenClaimTokenInterestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenClaimTokenInterest)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenClaimTokenInterest)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenClaimTokenInterestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenClaimTokenInterestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenClaimTokenInterest represents a ClaimTokenInterest event raised by the Token contract.
type TokenClaimTokenInterest struct {
	Owner                common.Address
	NftId                *big.Int
	InterestTokenClaimed *big.Int
	InterestTokenBonus   *big.Int
	BurnedItp            *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterClaimTokenInterest is a free log retrieval operation binding the contract event 0x199a7a8ae450da3700d2c3bd80f41a0f29b853dbde55f79199d8956ec3267dd6.
//
// Solidity: event ClaimTokenInterest(address indexed owner, uint256 indexed nftId, uint256 interestTokenClaimed, uint256 interestTokenBonus, uint256 burnedItp)
func (_Token *TokenFilterer) FilterClaimTokenInterest(opts *bind.FilterOpts, owner []common.Address, nftId []*big.Int) (*TokenClaimTokenInterestIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "ClaimTokenInterest", ownerRule, nftIdRule)
	if err != nil {
		return nil, err
	}
	return &TokenClaimTokenInterestIterator{contract: _Token.contract, event: "ClaimTokenInterest", logs: logs, sub: sub}, nil
}

// WatchClaimTokenInterest is a free log subscription operation binding the contract event 0x199a7a8ae450da3700d2c3bd80f41a0f29b853dbde55f79199d8956ec3267dd6.
//
// Solidity: event ClaimTokenInterest(address indexed owner, uint256 indexed nftId, uint256 interestTokenClaimed, uint256 interestTokenBonus, uint256 burnedItp)
func (_Token *TokenFilterer) WatchClaimTokenInterest(opts *bind.WatchOpts, sink chan<- *TokenClaimTokenInterest, owner []common.Address, nftId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "ClaimTokenInterest", ownerRule, nftIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenClaimTokenInterest)
				if err := _Token.contract.UnpackLog(event, "ClaimTokenInterest", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClaimTokenInterest is a log parse operation binding the contract event 0x199a7a8ae450da3700d2c3bd80f41a0f29b853dbde55f79199d8956ec3267dd6.
//
// Solidity: event ClaimTokenInterest(address indexed owner, uint256 indexed nftId, uint256 interestTokenClaimed, uint256 interestTokenBonus, uint256 burnedItp)
func (_Token *TokenFilterer) ParseClaimTokenInterest(log types.Log) (*TokenClaimTokenInterest, error) {
	event := new(TokenClaimTokenInterest)
	if err := _Token.contract.UnpackLog(event, "ClaimTokenInterest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenCloseLoanIterator is returned from FilterCloseLoan and is used to iterate over the raw logs and unpacked data for CloseLoan events raised by the Token contract.
type TokenCloseLoanIterator struct {
	Event *TokenCloseLoan // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenCloseLoanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenCloseLoan)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenCloseLoan)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenCloseLoanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenCloseLoanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenCloseLoan represents a CloseLoan event raised by the Token contract.
type TokenCloseLoan struct {
	Owner                    common.Address
	NftId                    *big.Int
	LoanId                   *big.Int
	BorrowPaid               *big.Int
	InterestPaid             *big.Int
	CollateralAmountwithdraw *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterCloseLoan is a free log retrieval operation binding the contract event 0x412fc0526cd5dcefcd8bd9058d40171270ec1eeb1a86605d1fd36ed006b4344d.
//
// Solidity: event CloseLoan(address indexed owner, uint256 indexed nftId, uint256 indexed loanId, uint256 borrowPaid, uint256 interestPaid, uint256 collateralAmountwithdraw)
func (_Token *TokenFilterer) FilterCloseLoan(opts *bind.FilterOpts, owner []common.Address, nftId []*big.Int, loanId []*big.Int) (*TokenCloseLoanIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}
	var loanIdRule []interface{}
	for _, loanIdItem := range loanId {
		loanIdRule = append(loanIdRule, loanIdItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "CloseLoan", ownerRule, nftIdRule, loanIdRule)
	if err != nil {
		return nil, err
	}
	return &TokenCloseLoanIterator{contract: _Token.contract, event: "CloseLoan", logs: logs, sub: sub}, nil
}

// WatchCloseLoan is a free log subscription operation binding the contract event 0x412fc0526cd5dcefcd8bd9058d40171270ec1eeb1a86605d1fd36ed006b4344d.
//
// Solidity: event CloseLoan(address indexed owner, uint256 indexed nftId, uint256 indexed loanId, uint256 borrowPaid, uint256 interestPaid, uint256 collateralAmountwithdraw)
func (_Token *TokenFilterer) WatchCloseLoan(opts *bind.WatchOpts, sink chan<- *TokenCloseLoan, owner []common.Address, nftId []*big.Int, loanId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}
	var loanIdRule []interface{}
	for _, loanIdItem := range loanId {
		loanIdRule = append(loanIdRule, loanIdItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "CloseLoan", ownerRule, nftIdRule, loanIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenCloseLoan)
				if err := _Token.contract.UnpackLog(event, "CloseLoan", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCloseLoan is a log parse operation binding the contract event 0x412fc0526cd5dcefcd8bd9058d40171270ec1eeb1a86605d1fd36ed006b4344d.
//
// Solidity: event CloseLoan(address indexed owner, uint256 indexed nftId, uint256 indexed loanId, uint256 borrowPaid, uint256 interestPaid, uint256 collateralAmountwithdraw)
func (_Token *TokenFilterer) ParseCloseLoan(log types.Log) (*TokenCloseLoan, error) {
	event := new(TokenCloseLoan)
	if err := _Token.contract.UnpackLog(event, "CloseLoan", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Token contract.
type TokenDepositIterator struct {
	Event *TokenDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenDeposit represents a Deposit event raised by the Token contract.
type TokenDeposit struct {
	Owner         common.Address
	NftId         *big.Int
	DepositAmount *big.Int
	MintedP       *big.Int
	MintedItp     *big.Int
	MintedIfp     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xf943cf10ef4d1e3239f4716ddecdf546e8ba8ab0e41deafd9a71a99936827e45.
//
// Solidity: event Deposit(address indexed owner, uint256 indexed nftId, uint256 depositAmount, uint256 mintedP, uint256 mintedItp, uint256 mintedIfp)
func (_Token *TokenFilterer) FilterDeposit(opts *bind.FilterOpts, owner []common.Address, nftId []*big.Int) (*TokenDepositIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Deposit", ownerRule, nftIdRule)
	if err != nil {
		return nil, err
	}
	return &TokenDepositIterator{contract: _Token.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xf943cf10ef4d1e3239f4716ddecdf546e8ba8ab0e41deafd9a71a99936827e45.
//
// Solidity: event Deposit(address indexed owner, uint256 indexed nftId, uint256 depositAmount, uint256 mintedP, uint256 mintedItp, uint256 mintedIfp)
func (_Token *TokenFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *TokenDeposit, owner []common.Address, nftId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Deposit", ownerRule, nftIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenDeposit)
				if err := _Token.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposit is a log parse operation binding the contract event 0xf943cf10ef4d1e3239f4716ddecdf546e8ba8ab0e41deafd9a71a99936827e45.
//
// Solidity: event Deposit(address indexed owner, uint256 indexed nftId, uint256 depositAmount, uint256 mintedP, uint256 mintedItp, uint256 mintedIfp)
func (_Token *TokenFilterer) ParseDeposit(log types.Log) (*TokenDeposit, error) {
	event := new(TokenDeposit)
	if err := _Token.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenDeprecateStakeInfoIterator is returned from FilterDeprecateStakeInfo and is used to iterate over the raw logs and unpacked data for DeprecateStakeInfo events raised by the Token contract.
type TokenDeprecateStakeInfoIterator struct {
	Event *TokenDeprecateStakeInfo // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenDeprecateStakeInfoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenDeprecateStakeInfo)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenDeprecateStakeInfo)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenDeprecateStakeInfoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenDeprecateStakeInfoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenDeprecateStakeInfo represents a DeprecateStakeInfo event raised by the Token contract.
type TokenDeprecateStakeInfo struct {
	Sender common.Address
	NftId  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeprecateStakeInfo is a free log retrieval operation binding the contract event 0x607fc08dc081fa4af4cb2b87cc9e94e4ec45e79f54c446c254c497387ed85194.
//
// Solidity: event DeprecateStakeInfo(address indexed sender, uint256 indexed nftId)
func (_Token *TokenFilterer) FilterDeprecateStakeInfo(opts *bind.FilterOpts, sender []common.Address, nftId []*big.Int) (*TokenDeprecateStakeInfoIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "DeprecateStakeInfo", senderRule, nftIdRule)
	if err != nil {
		return nil, err
	}
	return &TokenDeprecateStakeInfoIterator{contract: _Token.contract, event: "DeprecateStakeInfo", logs: logs, sub: sub}, nil
}

// WatchDeprecateStakeInfo is a free log subscription operation binding the contract event 0x607fc08dc081fa4af4cb2b87cc9e94e4ec45e79f54c446c254c497387ed85194.
//
// Solidity: event DeprecateStakeInfo(address indexed sender, uint256 indexed nftId)
func (_Token *TokenFilterer) WatchDeprecateStakeInfo(opts *bind.WatchOpts, sink chan<- *TokenDeprecateStakeInfo, sender []common.Address, nftId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "DeprecateStakeInfo", senderRule, nftIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenDeprecateStakeInfo)
				if err := _Token.contract.UnpackLog(event, "DeprecateStakeInfo", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeprecateStakeInfo is a log parse operation binding the contract event 0x607fc08dc081fa4af4cb2b87cc9e94e4ec45e79f54c446c254c497387ed85194.
//
// Solidity: event DeprecateStakeInfo(address indexed sender, uint256 indexed nftId)
func (_Token *TokenFilterer) ParseDeprecateStakeInfo(log types.Log) (*TokenDeprecateStakeInfo, error) {
	event := new(TokenDeprecateStakeInfo)
	if err := _Token.contract.UnpackLog(event, "DeprecateStakeInfo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenFaucetTransferIterator is returned from FilterFaucetTransfer and is used to iterate over the raw logs and unpacked data for FaucetTransfer events raised by the Token contract.
type TokenFaucetTransferIterator struct {
	Event *TokenFaucetTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenFaucetTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenFaucetTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenFaucetTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenFaucetTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenFaucetTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenFaucetTransfer represents a FaucetTransfer event raised by the Token contract.
type TokenFaucetTransfer struct {
	To           common.Address
	TokenAddress common.Address
	Value        *big.Int
	Timestamp    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFaucetTransfer is a free log retrieval operation binding the contract event 0x0bdd60bb057d785e4f01a1ca616f02e4300168ddc229e8f8dae2251c54a2cadd.
//
// Solidity: event FaucetTransfer(address indexed to, address tokenAddress, uint256 value, uint256 timestamp)
func (_Token *TokenFilterer) FilterFaucetTransfer(opts *bind.FilterOpts, to []common.Address) (*TokenFaucetTransferIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "FaucetTransfer", toRule)
	if err != nil {
		return nil, err
	}
	return &TokenFaucetTransferIterator{contract: _Token.contract, event: "FaucetTransfer", logs: logs, sub: sub}, nil
}

// WatchFaucetTransfer is a free log subscription operation binding the contract event 0x0bdd60bb057d785e4f01a1ca616f02e4300168ddc229e8f8dae2251c54a2cadd.
//
// Solidity: event FaucetTransfer(address indexed to, address tokenAddress, uint256 value, uint256 timestamp)
func (_Token *TokenFilterer) WatchFaucetTransfer(opts *bind.WatchOpts, sink chan<- *TokenFaucetTransfer, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "FaucetTransfer", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenFaucetTransfer)
				if err := _Token.contract.UnpackLog(event, "FaucetTransfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFaucetTransfer is a log parse operation binding the contract event 0x0bdd60bb057d785e4f01a1ca616f02e4300168ddc229e8f8dae2251c54a2cadd.
//
// Solidity: event FaucetTransfer(address indexed to, address tokenAddress, uint256 value, uint256 timestamp)
func (_Token *TokenFilterer) ParseFaucetTransfer(log types.Log) (*TokenFaucetTransfer, error) {
	event := new(TokenFaucetTransfer)
	if err := _Token.contract.UnpackLog(event, "FaucetTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenGlobalPricingPausedIterator is returned from FilterGlobalPricingPaused and is used to iterate over the raw logs and unpacked data for GlobalPricingPaused events raised by the Token contract.
type TokenGlobalPricingPausedIterator struct {
	Event *TokenGlobalPricingPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenGlobalPricingPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenGlobalPricingPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenGlobalPricingPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenGlobalPricingPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenGlobalPricingPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenGlobalPricingPaused represents a GlobalPricingPaused event raised by the Token contract.
type TokenGlobalPricingPaused struct {
	Sender   common.Address
	IsPaused bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterGlobalPricingPaused is a free log retrieval operation binding the contract event 0xd72570f3e3995824c4448eca84a22aed922dab8c58e23e9c5e60e74d0714f7f1.
//
// Solidity: event GlobalPricingPaused(address indexed sender, bool isPaused)
func (_Token *TokenFilterer) FilterGlobalPricingPaused(opts *bind.FilterOpts, sender []common.Address) (*TokenGlobalPricingPausedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "GlobalPricingPaused", senderRule)
	if err != nil {
		return nil, err
	}
	return &TokenGlobalPricingPausedIterator{contract: _Token.contract, event: "GlobalPricingPaused", logs: logs, sub: sub}, nil
}

// WatchGlobalPricingPaused is a free log subscription operation binding the contract event 0xd72570f3e3995824c4448eca84a22aed922dab8c58e23e9c5e60e74d0714f7f1.
//
// Solidity: event GlobalPricingPaused(address indexed sender, bool isPaused)
func (_Token *TokenFilterer) WatchGlobalPricingPaused(opts *bind.WatchOpts, sink chan<- *TokenGlobalPricingPaused, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "GlobalPricingPaused", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenGlobalPricingPaused)
				if err := _Token.contract.UnpackLog(event, "GlobalPricingPaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGlobalPricingPaused is a log parse operation binding the contract event 0xd72570f3e3995824c4448eca84a22aed922dab8c58e23e9c5e60e74d0714f7f1.
//
// Solidity: event GlobalPricingPaused(address indexed sender, bool isPaused)
func (_Token *TokenFilterer) ParseGlobalPricingPaused(log types.Log) (*TokenGlobalPricingPaused, error) {
	event := new(TokenGlobalPricingPaused)
	if err := _Token.contract.UnpackLog(event, "GlobalPricingPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenInitializeIterator is returned from FilterInitialize and is used to iterate over the raw logs and unpacked data for Initialize events raised by the Token contract.
type TokenInitializeIterator struct {
	Event *TokenInitialize // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenInitializeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenInitialize)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenInitialize)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenInitializeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenInitializeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenInitialize represents a Initialize event raised by the Token contract.
type TokenInitialize struct {
	Manager              common.Address
	CoreAddress          common.Address
	InterestVaultAddress common.Address
	MembershipAddress    common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterInitialize is a free log retrieval operation binding the contract event 0x420f11aae030b1b71b7480842c183b50e595c7c74203eb478d183ad7d3fdbd26.
//
// Solidity: event Initialize(address indexed manager, address indexed coreAddress, address interestVaultAddress, address membershipAddress)
func (_Token *TokenFilterer) FilterInitialize(opts *bind.FilterOpts, manager []common.Address, coreAddress []common.Address) (*TokenInitializeIterator, error) {

	var managerRule []interface{}
	for _, managerItem := range manager {
		managerRule = append(managerRule, managerItem)
	}
	var coreAddressRule []interface{}
	for _, coreAddressItem := range coreAddress {
		coreAddressRule = append(coreAddressRule, coreAddressItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Initialize", managerRule, coreAddressRule)
	if err != nil {
		return nil, err
	}
	return &TokenInitializeIterator{contract: _Token.contract, event: "Initialize", logs: logs, sub: sub}, nil
}

// WatchInitialize is a free log subscription operation binding the contract event 0x420f11aae030b1b71b7480842c183b50e595c7c74203eb478d183ad7d3fdbd26.
//
// Solidity: event Initialize(address indexed manager, address indexed coreAddress, address interestVaultAddress, address membershipAddress)
func (_Token *TokenFilterer) WatchInitialize(opts *bind.WatchOpts, sink chan<- *TokenInitialize, manager []common.Address, coreAddress []common.Address) (event.Subscription, error) {

	var managerRule []interface{}
	for _, managerItem := range manager {
		managerRule = append(managerRule, managerItem)
	}
	var coreAddressRule []interface{}
	for _, coreAddressItem := range coreAddress {
		coreAddressRule = append(coreAddressRule, coreAddressItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Initialize", managerRule, coreAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenInitialize)
				if err := _Token.contract.UnpackLog(event, "Initialize", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialize is a log parse operation binding the contract event 0x420f11aae030b1b71b7480842c183b50e595c7c74203eb478d183ad7d3fdbd26.
//
// Solidity: event Initialize(address indexed manager, address indexed coreAddress, address interestVaultAddress, address membershipAddress)
func (_Token *TokenFilterer) ParseInitialize(log types.Log) (*TokenInitialize, error) {
	event := new(TokenInitialize)
	if err := _Token.contract.UnpackLog(event, "Initialize", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenLiquidateIterator is returned from FilterLiquidate and is used to iterate over the raw logs and unpacked data for Liquidate events raised by the Token contract.
type TokenLiquidateIterator struct {
	Event *TokenLiquidate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenLiquidateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenLiquidate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenLiquidate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenLiquidateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenLiquidateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenLiquidate represents a Liquidate event raised by the Token contract.
type TokenLiquidate struct {
	Owner                    common.Address
	NftId                    *big.Int
	LoanId                   *big.Int
	Liquidator               common.Address
	SwapPrice                *big.Int
	TokenAmountFromSwap      *big.Int
	BountyReward             *big.Int
	BountyRewardTokenAddress common.Address
	TokenSentBackToUser      *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterLiquidate is a free log retrieval operation binding the contract event 0xf84d71ecc273b9e194eb1012e8fcf5ea8e113119298b775745839901c486ce26.
//
// Solidity: event Liquidate(address indexed owner, uint256 indexed nftId, uint256 indexed loanId, address liquidator, uint256 swapPrice, uint256 tokenAmountFromSwap, uint256 bountyReward, address bountyRewardTokenAddress, uint256 tokenSentBackToUser)
func (_Token *TokenFilterer) FilterLiquidate(opts *bind.FilterOpts, owner []common.Address, nftId []*big.Int, loanId []*big.Int) (*TokenLiquidateIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}
	var loanIdRule []interface{}
	for _, loanIdItem := range loanId {
		loanIdRule = append(loanIdRule, loanIdItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Liquidate", ownerRule, nftIdRule, loanIdRule)
	if err != nil {
		return nil, err
	}
	return &TokenLiquidateIterator{contract: _Token.contract, event: "Liquidate", logs: logs, sub: sub}, nil
}

// WatchLiquidate is a free log subscription operation binding the contract event 0xf84d71ecc273b9e194eb1012e8fcf5ea8e113119298b775745839901c486ce26.
//
// Solidity: event Liquidate(address indexed owner, uint256 indexed nftId, uint256 indexed loanId, address liquidator, uint256 swapPrice, uint256 tokenAmountFromSwap, uint256 bountyReward, address bountyRewardTokenAddress, uint256 tokenSentBackToUser)
func (_Token *TokenFilterer) WatchLiquidate(opts *bind.WatchOpts, sink chan<- *TokenLiquidate, owner []common.Address, nftId []*big.Int, loanId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}
	var loanIdRule []interface{}
	for _, loanIdItem := range loanId {
		loanIdRule = append(loanIdRule, loanIdItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Liquidate", ownerRule, nftIdRule, loanIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenLiquidate)
				if err := _Token.contract.UnpackLog(event, "Liquidate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLiquidate is a log parse operation binding the contract event 0xf84d71ecc273b9e194eb1012e8fcf5ea8e113119298b775745839901c486ce26.
//
// Solidity: event Liquidate(address indexed owner, uint256 indexed nftId, uint256 indexed loanId, address liquidator, uint256 swapPrice, uint256 tokenAmountFromSwap, uint256 bountyReward, address bountyRewardTokenAddress, uint256 tokenSentBackToUser)
func (_Token *TokenFilterer) ParseLiquidate(log types.Log) (*TokenLiquidate, error) {
	event := new(TokenLiquidate)
	if err := _Token.contract.UnpackLog(event, "Liquidate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenMinDelayChangeIterator is returned from FilterMinDelayChange and is used to iterate over the raw logs and unpacked data for MinDelayChange events raised by the Token contract.
type TokenMinDelayChangeIterator struct {
	Event *TokenMinDelayChange // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenMinDelayChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenMinDelayChange)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenMinDelayChange)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenMinDelayChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenMinDelayChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenMinDelayChange represents a MinDelayChange event raised by the Token contract.
type TokenMinDelayChange struct {
	OldDuration *big.Int
	NewDuration *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMinDelayChange is a free log retrieval operation binding the contract event 0x11c24f4ead16507c69ac467fbd5e4eed5fb5c699626d2cc6d66421df253886d5.
//
// Solidity: event MinDelayChange(uint256 oldDuration, uint256 newDuration)
func (_Token *TokenFilterer) FilterMinDelayChange(opts *bind.FilterOpts) (*TokenMinDelayChangeIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "MinDelayChange")
	if err != nil {
		return nil, err
	}
	return &TokenMinDelayChangeIterator{contract: _Token.contract, event: "MinDelayChange", logs: logs, sub: sub}, nil
}

// WatchMinDelayChange is a free log subscription operation binding the contract event 0x11c24f4ead16507c69ac467fbd5e4eed5fb5c699626d2cc6d66421df253886d5.
//
// Solidity: event MinDelayChange(uint256 oldDuration, uint256 newDuration)
func (_Token *TokenFilterer) WatchMinDelayChange(opts *bind.WatchOpts, sink chan<- *TokenMinDelayChange) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "MinDelayChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenMinDelayChange)
				if err := _Token.contract.UnpackLog(event, "MinDelayChange", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMinDelayChange is a log parse operation binding the contract event 0x11c24f4ead16507c69ac467fbd5e4eed5fb5c699626d2cc6d66421df253886d5.
//
// Solidity: event MinDelayChange(uint256 oldDuration, uint256 newDuration)
func (_Token *TokenFilterer) ParseMinDelayChange(log types.Log) (*TokenMinDelayChange, error) {
	event := new(TokenMinDelayChange)
	if err := _Token.contract.UnpackLog(event, "MinDelayChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenMinProxyAdminDelayChangeIterator is returned from FilterMinProxyAdminDelayChange and is used to iterate over the raw logs and unpacked data for MinProxyAdminDelayChange events raised by the Token contract.
type TokenMinProxyAdminDelayChangeIterator struct {
	Event *TokenMinProxyAdminDelayChange // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenMinProxyAdminDelayChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenMinProxyAdminDelayChange)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenMinProxyAdminDelayChange)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenMinProxyAdminDelayChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenMinProxyAdminDelayChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenMinProxyAdminDelayChange represents a MinProxyAdminDelayChange event raised by the Token contract.
type TokenMinProxyAdminDelayChange struct {
	OldDuration *big.Int
	NewDuration *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMinProxyAdminDelayChange is a free log retrieval operation binding the contract event 0x85c44ec58b94d3ae342506f4bde4fd3a5f24bc2e131377b4b12cb0c16a56e51e.
//
// Solidity: event MinProxyAdminDelayChange(uint256 oldDuration, uint256 newDuration)
func (_Token *TokenFilterer) FilterMinProxyAdminDelayChange(opts *bind.FilterOpts) (*TokenMinProxyAdminDelayChangeIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "MinProxyAdminDelayChange")
	if err != nil {
		return nil, err
	}
	return &TokenMinProxyAdminDelayChangeIterator{contract: _Token.contract, event: "MinProxyAdminDelayChange", logs: logs, sub: sub}, nil
}

// WatchMinProxyAdminDelayChange is a free log subscription operation binding the contract event 0x85c44ec58b94d3ae342506f4bde4fd3a5f24bc2e131377b4b12cb0c16a56e51e.
//
// Solidity: event MinProxyAdminDelayChange(uint256 oldDuration, uint256 newDuration)
func (_Token *TokenFilterer) WatchMinProxyAdminDelayChange(opts *bind.WatchOpts, sink chan<- *TokenMinProxyAdminDelayChange) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "MinProxyAdminDelayChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenMinProxyAdminDelayChange)
				if err := _Token.contract.UnpackLog(event, "MinProxyAdminDelayChange", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMinProxyAdminDelayChange is a log parse operation binding the contract event 0x85c44ec58b94d3ae342506f4bde4fd3a5f24bc2e131377b4b12cb0c16a56e51e.
//
// Solidity: event MinProxyAdminDelayChange(uint256 oldDuration, uint256 newDuration)
func (_Token *TokenFilterer) ParseMinProxyAdminDelayChange(log types.Log) (*TokenMinProxyAdminDelayChange, error) {
	event := new(TokenMinProxyAdminDelayChange)
	if err := _Token.contract.UnpackLog(event, "MinProxyAdminDelayChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenMinUrgentDelayChangeIterator is returned from FilterMinUrgentDelayChange and is used to iterate over the raw logs and unpacked data for MinUrgentDelayChange events raised by the Token contract.
type TokenMinUrgentDelayChangeIterator struct {
	Event *TokenMinUrgentDelayChange // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenMinUrgentDelayChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenMinUrgentDelayChange)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenMinUrgentDelayChange)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenMinUrgentDelayChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenMinUrgentDelayChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenMinUrgentDelayChange represents a MinUrgentDelayChange event raised by the Token contract.
type TokenMinUrgentDelayChange struct {
	OldDuration *big.Int
	NewDuration *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMinUrgentDelayChange is a free log retrieval operation binding the contract event 0xd0c2913b1a9168880e0e0a94ced6acb16e4c14aa67d9d138ee0eedc3fd423d63.
//
// Solidity: event MinUrgentDelayChange(uint256 oldDuration, uint256 newDuration)
func (_Token *TokenFilterer) FilterMinUrgentDelayChange(opts *bind.FilterOpts) (*TokenMinUrgentDelayChangeIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "MinUrgentDelayChange")
	if err != nil {
		return nil, err
	}
	return &TokenMinUrgentDelayChangeIterator{contract: _Token.contract, event: "MinUrgentDelayChange", logs: logs, sub: sub}, nil
}

// WatchMinUrgentDelayChange is a free log subscription operation binding the contract event 0xd0c2913b1a9168880e0e0a94ced6acb16e4c14aa67d9d138ee0eedc3fd423d63.
//
// Solidity: event MinUrgentDelayChange(uint256 oldDuration, uint256 newDuration)
func (_Token *TokenFilterer) WatchMinUrgentDelayChange(opts *bind.WatchOpts, sink chan<- *TokenMinUrgentDelayChange) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "MinUrgentDelayChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenMinUrgentDelayChange)
				if err := _Token.contract.UnpackLog(event, "MinUrgentDelayChange", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMinUrgentDelayChange is a log parse operation binding the contract event 0xd0c2913b1a9168880e0e0a94ced6acb16e4c14aa67d9d138ee0eedc3fd423d63.
//
// Solidity: event MinUrgentDelayChange(uint256 oldDuration, uint256 newDuration)
func (_Token *TokenFilterer) ParseMinUrgentDelayChange(log types.Log) (*TokenMinUrgentDelayChange, error) {
	event := new(TokenMinUrgentDelayChange)
	if err := _Token.contract.UnpackLog(event, "MinUrgentDelayChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenMintIfpTokenIterator is returned from FilterMintIfpToken and is used to iterate over the raw logs and unpacked data for MintIfpToken events raised by the Token contract.
type TokenMintIfpTokenIterator struct {
	Event *TokenMintIfpToken // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenMintIfpTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenMintIfpToken)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenMintIfpToken)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenMintIfpTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenMintIfpTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenMintIfpToken represents a MintIfpToken event raised by the Token contract.
type TokenMintIfpToken struct {
	Minter common.Address
	NftId  *big.Int
	Amount *big.Int
	Price  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMintIfpToken is a free log retrieval operation binding the contract event 0x714d551cd67ff1c6f5f83ee980fb6b1d8b94e5864c8c0240aa115340e8af243d.
//
// Solidity: event MintIfpToken(address indexed minter, uint256 indexed nftId, uint256 amount, uint256 price)
func (_Token *TokenFilterer) FilterMintIfpToken(opts *bind.FilterOpts, minter []common.Address, nftId []*big.Int) (*TokenMintIfpTokenIterator, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "MintIfpToken", minterRule, nftIdRule)
	if err != nil {
		return nil, err
	}
	return &TokenMintIfpTokenIterator{contract: _Token.contract, event: "MintIfpToken", logs: logs, sub: sub}, nil
}

// WatchMintIfpToken is a free log subscription operation binding the contract event 0x714d551cd67ff1c6f5f83ee980fb6b1d8b94e5864c8c0240aa115340e8af243d.
//
// Solidity: event MintIfpToken(address indexed minter, uint256 indexed nftId, uint256 amount, uint256 price)
func (_Token *TokenFilterer) WatchMintIfpToken(opts *bind.WatchOpts, sink chan<- *TokenMintIfpToken, minter []common.Address, nftId []*big.Int) (event.Subscription, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "MintIfpToken", minterRule, nftIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenMintIfpToken)
				if err := _Token.contract.UnpackLog(event, "MintIfpToken", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMintIfpToken is a log parse operation binding the contract event 0x714d551cd67ff1c6f5f83ee980fb6b1d8b94e5864c8c0240aa115340e8af243d.
//
// Solidity: event MintIfpToken(address indexed minter, uint256 indexed nftId, uint256 amount, uint256 price)
func (_Token *TokenFilterer) ParseMintIfpToken(log types.Log) (*TokenMintIfpToken, error) {
	event := new(TokenMintIfpToken)
	if err := _Token.contract.UnpackLog(event, "MintIfpToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenMintItpTokenIterator is returned from FilterMintItpToken and is used to iterate over the raw logs and unpacked data for MintItpToken events raised by the Token contract.
type TokenMintItpTokenIterator struct {
	Event *TokenMintItpToken // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenMintItpTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenMintItpToken)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenMintItpToken)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenMintItpTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenMintItpTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenMintItpToken represents a MintItpToken event raised by the Token contract.
type TokenMintItpToken struct {
	Minter common.Address
	NftId  *big.Int
	Amount *big.Int
	Price  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMintItpToken is a free log retrieval operation binding the contract event 0xbc166a5440f6f3afa832b0aab4f37288d7f3a2fa7cadf9a325e0ffe57ad8a059.
//
// Solidity: event MintItpToken(address indexed minter, uint256 indexed nftId, uint256 amount, uint256 price)
func (_Token *TokenFilterer) FilterMintItpToken(opts *bind.FilterOpts, minter []common.Address, nftId []*big.Int) (*TokenMintItpTokenIterator, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "MintItpToken", minterRule, nftIdRule)
	if err != nil {
		return nil, err
	}
	return &TokenMintItpTokenIterator{contract: _Token.contract, event: "MintItpToken", logs: logs, sub: sub}, nil
}

// WatchMintItpToken is a free log subscription operation binding the contract event 0xbc166a5440f6f3afa832b0aab4f37288d7f3a2fa7cadf9a325e0ffe57ad8a059.
//
// Solidity: event MintItpToken(address indexed minter, uint256 indexed nftId, uint256 amount, uint256 price)
func (_Token *TokenFilterer) WatchMintItpToken(opts *bind.WatchOpts, sink chan<- *TokenMintItpToken, minter []common.Address, nftId []*big.Int) (event.Subscription, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "MintItpToken", minterRule, nftIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenMintItpToken)
				if err := _Token.contract.UnpackLog(event, "MintItpToken", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMintItpToken is a log parse operation binding the contract event 0xbc166a5440f6f3afa832b0aab4f37288d7f3a2fa7cadf9a325e0ffe57ad8a059.
//
// Solidity: event MintItpToken(address indexed minter, uint256 indexed nftId, uint256 amount, uint256 price)
func (_Token *TokenFilterer) ParseMintItpToken(log types.Log) (*TokenMintItpToken, error) {
	event := new(TokenMintItpToken)
	if err := _Token.contract.UnpackLog(event, "MintItpToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenMintPTokenIterator is returned from FilterMintPToken and is used to iterate over the raw logs and unpacked data for MintPToken events raised by the Token contract.
type TokenMintPTokenIterator struct {
	Event *TokenMintPToken // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenMintPTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenMintPToken)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenMintPToken)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenMintPTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenMintPTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenMintPToken represents a MintPToken event raised by the Token contract.
type TokenMintPToken struct {
	Minter common.Address
	NftId  *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMintPToken is a free log retrieval operation binding the contract event 0x2593172dfe1b54f3821b7b56adb5b36ef5267d46187d2b431013bafa8287a3e6.
//
// Solidity: event MintPToken(address indexed minter, uint256 indexed nftId, uint256 amount)
func (_Token *TokenFilterer) FilterMintPToken(opts *bind.FilterOpts, minter []common.Address, nftId []*big.Int) (*TokenMintPTokenIterator, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "MintPToken", minterRule, nftIdRule)
	if err != nil {
		return nil, err
	}
	return &TokenMintPTokenIterator{contract: _Token.contract, event: "MintPToken", logs: logs, sub: sub}, nil
}

// WatchMintPToken is a free log subscription operation binding the contract event 0x2593172dfe1b54f3821b7b56adb5b36ef5267d46187d2b431013bafa8287a3e6.
//
// Solidity: event MintPToken(address indexed minter, uint256 indexed nftId, uint256 amount)
func (_Token *TokenFilterer) WatchMintPToken(opts *bind.WatchOpts, sink chan<- *TokenMintPToken, minter []common.Address, nftId []*big.Int) (event.Subscription, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "MintPToken", minterRule, nftIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenMintPToken)
				if err := _Token.contract.UnpackLog(event, "MintPToken", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMintPToken is a log parse operation binding the contract event 0x2593172dfe1b54f3821b7b56adb5b36ef5267d46187d2b431013bafa8287a3e6.
//
// Solidity: event MintPToken(address indexed minter, uint256 indexed nftId, uint256 amount)
func (_Token *TokenFilterer) ParseMintPToken(log types.Log) (*TokenMintPToken, error) {
	event := new(TokenMintPToken)
	if err := _Token.contract.UnpackLog(event, "MintPToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenOwnerApproveIterator is returned from FilterOwnerApprove and is used to iterate over the raw logs and unpacked data for OwnerApprove events raised by the Token contract.
type TokenOwnerApproveIterator struct {
	Event *TokenOwnerApprove // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenOwnerApproveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenOwnerApprove)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenOwnerApprove)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenOwnerApproveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenOwnerApproveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenOwnerApprove represents a OwnerApprove event raised by the Token contract.
type TokenOwnerApprove struct {
	Sender       common.Address
	TokenAddress common.Address
	ForwAddress  common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnerApprove is a free log retrieval operation binding the contract event 0x70761c05296a92277eb19043412d0c73e9c7cc51ae48f4df2617d0b5f6944bd3.
//
// Solidity: event OwnerApprove(address indexed sender, address tokenAddress, address forwAddress, uint256 amount)
func (_Token *TokenFilterer) FilterOwnerApprove(opts *bind.FilterOpts, sender []common.Address) (*TokenOwnerApproveIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "OwnerApprove", senderRule)
	if err != nil {
		return nil, err
	}
	return &TokenOwnerApproveIterator{contract: _Token.contract, event: "OwnerApprove", logs: logs, sub: sub}, nil
}

// WatchOwnerApprove is a free log subscription operation binding the contract event 0x70761c05296a92277eb19043412d0c73e9c7cc51ae48f4df2617d0b5f6944bd3.
//
// Solidity: event OwnerApprove(address indexed sender, address tokenAddress, address forwAddress, uint256 amount)
func (_Token *TokenFilterer) WatchOwnerApprove(opts *bind.WatchOpts, sink chan<- *TokenOwnerApprove, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "OwnerApprove", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenOwnerApprove)
				if err := _Token.contract.UnpackLog(event, "OwnerApprove", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnerApprove is a log parse operation binding the contract event 0x70761c05296a92277eb19043412d0c73e9c7cc51ae48f4df2617d0b5f6944bd3.
//
// Solidity: event OwnerApprove(address indexed sender, address tokenAddress, address forwAddress, uint256 amount)
func (_Token *TokenFilterer) ParseOwnerApprove(log types.Log) (*TokenOwnerApprove, error) {
	event := new(TokenOwnerApprove)
	if err := _Token.contract.UnpackLog(event, "OwnerApprove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenProxyAdminChangeIterator is returned from FilterProxyAdminChange and is used to iterate over the raw logs and unpacked data for ProxyAdminChange events raised by the Token contract.
type TokenProxyAdminChangeIterator struct {
	Event *TokenProxyAdminChange // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenProxyAdminChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenProxyAdminChange)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenProxyAdminChange)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenProxyAdminChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenProxyAdminChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenProxyAdminChange represents a ProxyAdminChange event raised by the Token contract.
type TokenProxyAdminChange struct {
	OldProxyAdmin common.Address
	NewProxyAdmin common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterProxyAdminChange is a free log retrieval operation binding the contract event 0x03f1c187037dd01b9806846c1efb47f2737c75a6c44b4b3cb78628a40ea4ee9c.
//
// Solidity: event ProxyAdminChange(address oldProxyAdmin, address newProxyAdmin)
func (_Token *TokenFilterer) FilterProxyAdminChange(opts *bind.FilterOpts) (*TokenProxyAdminChangeIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "ProxyAdminChange")
	if err != nil {
		return nil, err
	}
	return &TokenProxyAdminChangeIterator{contract: _Token.contract, event: "ProxyAdminChange", logs: logs, sub: sub}, nil
}

// WatchProxyAdminChange is a free log subscription operation binding the contract event 0x03f1c187037dd01b9806846c1efb47f2737c75a6c44b4b3cb78628a40ea4ee9c.
//
// Solidity: event ProxyAdminChange(address oldProxyAdmin, address newProxyAdmin)
func (_Token *TokenFilterer) WatchProxyAdminChange(opts *bind.WatchOpts, sink chan<- *TokenProxyAdminChange) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "ProxyAdminChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenProxyAdminChange)
				if err := _Token.contract.UnpackLog(event, "ProxyAdminChange", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProxyAdminChange is a log parse operation binding the contract event 0x03f1c187037dd01b9806846c1efb47f2737c75a6c44b4b3cb78628a40ea4ee9c.
//
// Solidity: event ProxyAdminChange(address oldProxyAdmin, address newProxyAdmin)
func (_Token *TokenFilterer) ParseProxyAdminChange(log types.Log) (*TokenProxyAdminChange, error) {
	event := new(TokenProxyAdminChange)
	if err := _Token.contract.UnpackLog(event, "ProxyAdminChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenRegisterNewPoolIterator is returned from FilterRegisterNewPool and is used to iterate over the raw logs and unpacked data for RegisterNewPool events raised by the Token contract.
type TokenRegisterNewPoolIterator struct {
	Event *TokenRegisterNewPool // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenRegisterNewPoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenRegisterNewPool)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenRegisterNewPool)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenRegisterNewPoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenRegisterNewPoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenRegisterNewPool represents a RegisterNewPool event raised by the Token contract.
type TokenRegisterNewPool struct {
	Sender      common.Address
	PoolAddress common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRegisterNewPool is a free log retrieval operation binding the contract event 0xc12160e153d12997b24933aaecb5ea1fa7ffa03fdf71e674f0325aaea766541f.
//
// Solidity: event RegisterNewPool(address indexed sender, address poolAddress)
func (_Token *TokenFilterer) FilterRegisterNewPool(opts *bind.FilterOpts, sender []common.Address) (*TokenRegisterNewPoolIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "RegisterNewPool", senderRule)
	if err != nil {
		return nil, err
	}
	return &TokenRegisterNewPoolIterator{contract: _Token.contract, event: "RegisterNewPool", logs: logs, sub: sub}, nil
}

// WatchRegisterNewPool is a free log subscription operation binding the contract event 0xc12160e153d12997b24933aaecb5ea1fa7ffa03fdf71e674f0325aaea766541f.
//
// Solidity: event RegisterNewPool(address indexed sender, address poolAddress)
func (_Token *TokenFilterer) WatchRegisterNewPool(opts *bind.WatchOpts, sink chan<- *TokenRegisterNewPool, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "RegisterNewPool", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenRegisterNewPool)
				if err := _Token.contract.UnpackLog(event, "RegisterNewPool", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRegisterNewPool is a log parse operation binding the contract event 0xc12160e153d12997b24933aaecb5ea1fa7ffa03fdf71e674f0325aaea766541f.
//
// Solidity: event RegisterNewPool(address indexed sender, address poolAddress)
func (_Token *TokenFilterer) ParseRegisterNewPool(log types.Log) (*TokenRegisterNewPool, error) {
	event := new(TokenRegisterNewPool)
	if err := _Token.contract.UnpackLog(event, "RegisterNewPool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenRepayIterator is returned from FilterRepay and is used to iterate over the raw logs and unpacked data for Repay events raised by the Token contract.
type TokenRepayIterator struct {
	Event *TokenRepay // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenRepayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenRepay)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenRepay)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenRepayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenRepayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenRepay represents a Repay event raised by the Token contract.
type TokenRepay struct {
	Owner                    common.Address
	NftId                    *big.Int
	LoanId                   *big.Int
	CloseLoan                bool
	BorrowPaid               *big.Int
	InterestPaid             *big.Int
	CollateralAmountwithdraw *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterRepay is a free log retrieval operation binding the contract event 0x546807c7ba535bd956244fac7a3ccaa2e8877dcc277a84508261e0d4a817718e.
//
// Solidity: event Repay(address indexed owner, uint256 indexed nftId, uint256 indexed loanId, bool closeLoan, uint256 borrowPaid, uint256 interestPaid, uint256 collateralAmountwithdraw)
func (_Token *TokenFilterer) FilterRepay(opts *bind.FilterOpts, owner []common.Address, nftId []*big.Int, loanId []*big.Int) (*TokenRepayIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}
	var loanIdRule []interface{}
	for _, loanIdItem := range loanId {
		loanIdRule = append(loanIdRule, loanIdItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Repay", ownerRule, nftIdRule, loanIdRule)
	if err != nil {
		return nil, err
	}
	return &TokenRepayIterator{contract: _Token.contract, event: "Repay", logs: logs, sub: sub}, nil
}

// WatchRepay is a free log subscription operation binding the contract event 0x546807c7ba535bd956244fac7a3ccaa2e8877dcc277a84508261e0d4a817718e.
//
// Solidity: event Repay(address indexed owner, uint256 indexed nftId, uint256 indexed loanId, bool closeLoan, uint256 borrowPaid, uint256 interestPaid, uint256 collateralAmountwithdraw)
func (_Token *TokenFilterer) WatchRepay(opts *bind.WatchOpts, sink chan<- *TokenRepay, owner []common.Address, nftId []*big.Int, loanId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}
	var loanIdRule []interface{}
	for _, loanIdItem := range loanId {
		loanIdRule = append(loanIdRule, loanIdItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Repay", ownerRule, nftIdRule, loanIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenRepay)
				if err := _Token.contract.UnpackLog(event, "Repay", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRepay is a log parse operation binding the contract event 0x546807c7ba535bd956244fac7a3ccaa2e8877dcc277a84508261e0d4a817718e.
//
// Solidity: event Repay(address indexed owner, uint256 indexed nftId, uint256 indexed loanId, bool closeLoan, uint256 borrowPaid, uint256 interestPaid, uint256 collateralAmountwithdraw)
func (_Token *TokenFilterer) ParseRepay(log types.Log) (*TokenRepay, error) {
	event := new(TokenRepay)
	if err := _Token.contract.UnpackLog(event, "Repay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenRolloverIterator is returned from FilterRollover and is used to iterate over the raw logs and unpacked data for Rollover events raised by the Token contract.
type TokenRolloverIterator struct {
	Event *TokenRollover // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenRolloverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenRollover)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenRollover)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenRolloverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenRolloverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenRollover represents a Rollover event raised by the Token contract.
type TokenRollover struct {
	Owner                    common.Address
	NftId                    *big.Int
	LoanId                   *big.Int
	BountyHunter             common.Address
	DelayInterest            *big.Int
	BountyReward             *big.Int
	BountyRewardTokenAddress common.Address
	NewInterestOwedPerDay    *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterRollover is a free log retrieval operation binding the contract event 0x8996e3197cdf53f0ecd3d2cd081917e64c75d04cf176ddafafd1733bb691d169.
//
// Solidity: event Rollover(address indexed owner, uint256 indexed nftId, uint256 indexed loanId, address bountyHunter, uint256 delayInterest, uint256 bountyReward, address bountyRewardTokenAddress, uint256 newInterestOwedPerDay)
func (_Token *TokenFilterer) FilterRollover(opts *bind.FilterOpts, owner []common.Address, nftId []*big.Int, loanId []*big.Int) (*TokenRolloverIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}
	var loanIdRule []interface{}
	for _, loanIdItem := range loanId {
		loanIdRule = append(loanIdRule, loanIdItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Rollover", ownerRule, nftIdRule, loanIdRule)
	if err != nil {
		return nil, err
	}
	return &TokenRolloverIterator{contract: _Token.contract, event: "Rollover", logs: logs, sub: sub}, nil
}

// WatchRollover is a free log subscription operation binding the contract event 0x8996e3197cdf53f0ecd3d2cd081917e64c75d04cf176ddafafd1733bb691d169.
//
// Solidity: event Rollover(address indexed owner, uint256 indexed nftId, uint256 indexed loanId, address bountyHunter, uint256 delayInterest, uint256 bountyReward, address bountyRewardTokenAddress, uint256 newInterestOwedPerDay)
func (_Token *TokenFilterer) WatchRollover(opts *bind.WatchOpts, sink chan<- *TokenRollover, owner []common.Address, nftId []*big.Int, loanId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}
	var loanIdRule []interface{}
	for _, loanIdItem := range loanId {
		loanIdRule = append(loanIdRule, loanIdItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Rollover", ownerRule, nftIdRule, loanIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenRollover)
				if err := _Token.contract.UnpackLog(event, "Rollover", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRollover is a log parse operation binding the contract event 0x8996e3197cdf53f0ecd3d2cd081917e64c75d04cf176ddafafd1733bb691d169.
//
// Solidity: event Rollover(address indexed owner, uint256 indexed nftId, uint256 indexed loanId, address bountyHunter, uint256 delayInterest, uint256 bountyReward, address bountyRewardTokenAddress, uint256 newInterestOwedPerDay)
func (_Token *TokenFilterer) ParseRollover(log types.Log) (*TokenRollover, error) {
	event := new(TokenRollover)
	if err := _Token.contract.UnpackLog(event, "Rollover", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenSetAdvancedInterestDurationIterator is returned from FilterSetAdvancedInterestDuration and is used to iterate over the raw logs and unpacked data for SetAdvancedInterestDuration events raised by the Token contract.
type TokenSetAdvancedInterestDurationIterator struct {
	Event *TokenSetAdvancedInterestDuration // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenSetAdvancedInterestDurationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenSetAdvancedInterestDuration)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenSetAdvancedInterestDuration)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenSetAdvancedInterestDurationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenSetAdvancedInterestDurationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenSetAdvancedInterestDuration represents a SetAdvancedInterestDuration event raised by the Token contract.
type TokenSetAdvancedInterestDuration struct {
	Sender   common.Address
	OldValue *big.Int
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetAdvancedInterestDuration is a free log retrieval operation binding the contract event 0x27d95b47aff3e3fd4ab379ce195ff563567a9c9458e42f659cf5e4927064b0ce.
//
// Solidity: event SetAdvancedInterestDuration(address indexed sender, uint256 oldValue, uint256 newValue)
func (_Token *TokenFilterer) FilterSetAdvancedInterestDuration(opts *bind.FilterOpts, sender []common.Address) (*TokenSetAdvancedInterestDurationIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "SetAdvancedInterestDuration", senderRule)
	if err != nil {
		return nil, err
	}
	return &TokenSetAdvancedInterestDurationIterator{contract: _Token.contract, event: "SetAdvancedInterestDuration", logs: logs, sub: sub}, nil
}

// WatchSetAdvancedInterestDuration is a free log subscription operation binding the contract event 0x27d95b47aff3e3fd4ab379ce195ff563567a9c9458e42f659cf5e4927064b0ce.
//
// Solidity: event SetAdvancedInterestDuration(address indexed sender, uint256 oldValue, uint256 newValue)
func (_Token *TokenFilterer) WatchSetAdvancedInterestDuration(opts *bind.WatchOpts, sink chan<- *TokenSetAdvancedInterestDuration, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "SetAdvancedInterestDuration", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenSetAdvancedInterestDuration)
				if err := _Token.contract.UnpackLog(event, "SetAdvancedInterestDuration", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetAdvancedInterestDuration is a log parse operation binding the contract event 0x27d95b47aff3e3fd4ab379ce195ff563567a9c9458e42f659cf5e4927064b0ce.
//
// Solidity: event SetAdvancedInterestDuration(address indexed sender, uint256 oldValue, uint256 newValue)
func (_Token *TokenFilterer) ParseSetAdvancedInterestDuration(log types.Log) (*TokenSetAdvancedInterestDuration, error) {
	event := new(TokenSetAdvancedInterestDuration)
	if err := _Token.contract.UnpackLog(event, "SetAdvancedInterestDuration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenSetBorrowInterestParamsIterator is returned from FilterSetBorrowInterestParams and is used to iterate over the raw logs and unpacked data for SetBorrowInterestParams events raised by the Token contract.
type TokenSetBorrowInterestParamsIterator struct {
	Event *TokenSetBorrowInterestParams // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenSetBorrowInterestParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenSetBorrowInterestParams)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenSetBorrowInterestParams)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenSetBorrowInterestParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenSetBorrowInterestParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenSetBorrowInterestParams represents a SetBorrowInterestParams event raised by the Token contract.
type TokenSetBorrowInterestParams struct {
	Sender       common.Address
	Rates        []*big.Int
	Utils        []*big.Int
	TargetSupply *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSetBorrowInterestParams is a free log retrieval operation binding the contract event 0xb04a151596a26d09f2788b960f91f4e9445689e900b062f9f8bd143c114e0fab.
//
// Solidity: event SetBorrowInterestParams(address indexed sender, uint256[] rates, uint256[] utils, uint256 targetSupply)
func (_Token *TokenFilterer) FilterSetBorrowInterestParams(opts *bind.FilterOpts, sender []common.Address) (*TokenSetBorrowInterestParamsIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "SetBorrowInterestParams", senderRule)
	if err != nil {
		return nil, err
	}
	return &TokenSetBorrowInterestParamsIterator{contract: _Token.contract, event: "SetBorrowInterestParams", logs: logs, sub: sub}, nil
}

// WatchSetBorrowInterestParams is a free log subscription operation binding the contract event 0xb04a151596a26d09f2788b960f91f4e9445689e900b062f9f8bd143c114e0fab.
//
// Solidity: event SetBorrowInterestParams(address indexed sender, uint256[] rates, uint256[] utils, uint256 targetSupply)
func (_Token *TokenFilterer) WatchSetBorrowInterestParams(opts *bind.WatchOpts, sink chan<- *TokenSetBorrowInterestParams, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "SetBorrowInterestParams", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenSetBorrowInterestParams)
				if err := _Token.contract.UnpackLog(event, "SetBorrowInterestParams", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetBorrowInterestParams is a log parse operation binding the contract event 0xb04a151596a26d09f2788b960f91f4e9445689e900b062f9f8bd143c114e0fab.
//
// Solidity: event SetBorrowInterestParams(address indexed sender, uint256[] rates, uint256[] utils, uint256 targetSupply)
func (_Token *TokenFilterer) ParseSetBorrowInterestParams(log types.Log) (*TokenSetBorrowInterestParams, error) {
	event := new(TokenSetBorrowInterestParams)
	if err := _Token.contract.UnpackLog(event, "SetBorrowInterestParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenSetCoreBorrowingAddressIterator is returned from FilterSetCoreBorrowingAddress and is used to iterate over the raw logs and unpacked data for SetCoreBorrowingAddress events raised by the Token contract.
type TokenSetCoreBorrowingAddressIterator struct {
	Event *TokenSetCoreBorrowingAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenSetCoreBorrowingAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenSetCoreBorrowingAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenSetCoreBorrowingAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenSetCoreBorrowingAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenSetCoreBorrowingAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenSetCoreBorrowingAddress represents a SetCoreBorrowingAddress event raised by the Token contract.
type TokenSetCoreBorrowingAddress struct {
	Sender   common.Address
	OldValue common.Address
	NewValue common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetCoreBorrowingAddress is a free log retrieval operation binding the contract event 0xf903b347045e8554edd091377f04860add48bc1ec60624de790df9e69a5a75f5.
//
// Solidity: event SetCoreBorrowingAddress(address indexed sender, address oldValue, address newValue)
func (_Token *TokenFilterer) FilterSetCoreBorrowingAddress(opts *bind.FilterOpts, sender []common.Address) (*TokenSetCoreBorrowingAddressIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "SetCoreBorrowingAddress", senderRule)
	if err != nil {
		return nil, err
	}
	return &TokenSetCoreBorrowingAddressIterator{contract: _Token.contract, event: "SetCoreBorrowingAddress", logs: logs, sub: sub}, nil
}

// WatchSetCoreBorrowingAddress is a free log subscription operation binding the contract event 0xf903b347045e8554edd091377f04860add48bc1ec60624de790df9e69a5a75f5.
//
// Solidity: event SetCoreBorrowingAddress(address indexed sender, address oldValue, address newValue)
func (_Token *TokenFilterer) WatchSetCoreBorrowingAddress(opts *bind.WatchOpts, sink chan<- *TokenSetCoreBorrowingAddress, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "SetCoreBorrowingAddress", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenSetCoreBorrowingAddress)
				if err := _Token.contract.UnpackLog(event, "SetCoreBorrowingAddress", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetCoreBorrowingAddress is a log parse operation binding the contract event 0xf903b347045e8554edd091377f04860add48bc1ec60624de790df9e69a5a75f5.
//
// Solidity: event SetCoreBorrowingAddress(address indexed sender, address oldValue, address newValue)
func (_Token *TokenFilterer) ParseSetCoreBorrowingAddress(log types.Log) (*TokenSetCoreBorrowingAddress, error) {
	event := new(TokenSetCoreBorrowingAddress)
	if err := _Token.contract.UnpackLog(event, "SetCoreBorrowingAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenSetDecimalsIterator is returned from FilterSetDecimals and is used to iterate over the raw logs and unpacked data for SetDecimals events raised by the Token contract.
type TokenSetDecimalsIterator struct {
	Event *TokenSetDecimals // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenSetDecimalsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenSetDecimals)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenSetDecimals)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenSetDecimalsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenSetDecimalsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenSetDecimals represents a SetDecimals event raised by the Token contract.
type TokenSetDecimals struct {
	Sender common.Address
	Tokens []common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetDecimals is a free log retrieval operation binding the contract event 0xe71dc4253b36cab4de9843f1b202d9c4b6205c00bfa23006461372c0786536d8.
//
// Solidity: event SetDecimals(address indexed sender, address[] tokens)
func (_Token *TokenFilterer) FilterSetDecimals(opts *bind.FilterOpts, sender []common.Address) (*TokenSetDecimalsIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "SetDecimals", senderRule)
	if err != nil {
		return nil, err
	}
	return &TokenSetDecimalsIterator{contract: _Token.contract, event: "SetDecimals", logs: logs, sub: sub}, nil
}

// WatchSetDecimals is a free log subscription operation binding the contract event 0xe71dc4253b36cab4de9843f1b202d9c4b6205c00bfa23006461372c0786536d8.
//
// Solidity: event SetDecimals(address indexed sender, address[] tokens)
func (_Token *TokenFilterer) WatchSetDecimals(opts *bind.WatchOpts, sink chan<- *TokenSetDecimals, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "SetDecimals", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenSetDecimals)
				if err := _Token.contract.UnpackLog(event, "SetDecimals", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetDecimals is a log parse operation binding the contract event 0xe71dc4253b36cab4de9843f1b202d9c4b6205c00bfa23006461372c0786536d8.
//
// Solidity: event SetDecimals(address indexed sender, address[] tokens)
func (_Token *TokenFilterer) ParseSetDecimals(log types.Log) (*TokenSetDecimals, error) {
	event := new(TokenSetDecimals)
	if err := _Token.contract.UnpackLog(event, "SetDecimals", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenSetFeeControllerIterator is returned from FilterSetFeeController and is used to iterate over the raw logs and unpacked data for SetFeeController events raised by the Token contract.
type TokenSetFeeControllerIterator struct {
	Event *TokenSetFeeController // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenSetFeeControllerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenSetFeeController)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenSetFeeController)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenSetFeeControllerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenSetFeeControllerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenSetFeeController represents a SetFeeController event raised by the Token contract.
type TokenSetFeeController struct {
	Sender   common.Address
	OldValue common.Address
	NewValue common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetFeeController is a free log retrieval operation binding the contract event 0x03675447460d2a30f368ec273be49f07bfc1fb9b8e0a857fd830f134e7a6a449.
//
// Solidity: event SetFeeController(address indexed sender, address oldValue, address newValue)
func (_Token *TokenFilterer) FilterSetFeeController(opts *bind.FilterOpts, sender []common.Address) (*TokenSetFeeControllerIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "SetFeeController", senderRule)
	if err != nil {
		return nil, err
	}
	return &TokenSetFeeControllerIterator{contract: _Token.contract, event: "SetFeeController", logs: logs, sub: sub}, nil
}

// WatchSetFeeController is a free log subscription operation binding the contract event 0x03675447460d2a30f368ec273be49f07bfc1fb9b8e0a857fd830f134e7a6a449.
//
// Solidity: event SetFeeController(address indexed sender, address oldValue, address newValue)
func (_Token *TokenFilterer) WatchSetFeeController(opts *bind.WatchOpts, sink chan<- *TokenSetFeeController, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "SetFeeController", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenSetFeeController)
				if err := _Token.contract.UnpackLog(event, "SetFeeController", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetFeeController is a log parse operation binding the contract event 0x03675447460d2a30f368ec273be49f07bfc1fb9b8e0a857fd830f134e7a6a449.
//
// Solidity: event SetFeeController(address indexed sender, address oldValue, address newValue)
func (_Token *TokenFilterer) ParseSetFeeController(log types.Log) (*TokenSetFeeController, error) {
	event := new(TokenSetFeeController)
	if err := _Token.contract.UnpackLog(event, "SetFeeController", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenSetFeeSpreadIterator is returned from FilterSetFeeSpread and is used to iterate over the raw logs and unpacked data for SetFeeSpread events raised by the Token contract.
type TokenSetFeeSpreadIterator struct {
	Event *TokenSetFeeSpread // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenSetFeeSpreadIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenSetFeeSpread)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenSetFeeSpread)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenSetFeeSpreadIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenSetFeeSpreadIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenSetFeeSpread represents a SetFeeSpread event raised by the Token contract.
type TokenSetFeeSpread struct {
	Sender   common.Address
	OldValue *big.Int
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetFeeSpread is a free log retrieval operation binding the contract event 0xdae6e940fbc41dbcabe383eda5cc6429cadb587f6249781ebf917c0056707103.
//
// Solidity: event SetFeeSpread(address indexed sender, uint256 oldValue, uint256 newValue)
func (_Token *TokenFilterer) FilterSetFeeSpread(opts *bind.FilterOpts, sender []common.Address) (*TokenSetFeeSpreadIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "SetFeeSpread", senderRule)
	if err != nil {
		return nil, err
	}
	return &TokenSetFeeSpreadIterator{contract: _Token.contract, event: "SetFeeSpread", logs: logs, sub: sub}, nil
}

// WatchSetFeeSpread is a free log subscription operation binding the contract event 0xdae6e940fbc41dbcabe383eda5cc6429cadb587f6249781ebf917c0056707103.
//
// Solidity: event SetFeeSpread(address indexed sender, uint256 oldValue, uint256 newValue)
func (_Token *TokenFilterer) WatchSetFeeSpread(opts *bind.WatchOpts, sink chan<- *TokenSetFeeSpread, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "SetFeeSpread", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenSetFeeSpread)
				if err := _Token.contract.UnpackLog(event, "SetFeeSpread", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetFeeSpread is a log parse operation binding the contract event 0xdae6e940fbc41dbcabe383eda5cc6429cadb587f6249781ebf917c0056707103.
//
// Solidity: event SetFeeSpread(address indexed sender, uint256 oldValue, uint256 newValue)
func (_Token *TokenFilterer) ParseSetFeeSpread(log types.Log) (*TokenSetFeeSpread, error) {
	event := new(TokenSetFeeSpread)
	if err := _Token.contract.UnpackLog(event, "SetFeeSpread", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenSetForwAddressIterator is returned from FilterSetForwAddress and is used to iterate over the raw logs and unpacked data for SetForwAddress events raised by the Token contract.
type TokenSetForwAddressIterator struct {
	Event *TokenSetForwAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenSetForwAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenSetForwAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenSetForwAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenSetForwAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenSetForwAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenSetForwAddress represents a SetForwAddress event raised by the Token contract.
type TokenSetForwAddress struct {
	Sender   common.Address
	OldValue common.Address
	NewValue common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetForwAddress is a free log retrieval operation binding the contract event 0x9408dd080d4dfc13f9a203a2fb36089a7c08db7ac37f83da7b473ebfa5881729.
//
// Solidity: event SetForwAddress(address indexed sender, address oldValue, address newValue)
func (_Token *TokenFilterer) FilterSetForwAddress(opts *bind.FilterOpts, sender []common.Address) (*TokenSetForwAddressIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "SetForwAddress", senderRule)
	if err != nil {
		return nil, err
	}
	return &TokenSetForwAddressIterator{contract: _Token.contract, event: "SetForwAddress", logs: logs, sub: sub}, nil
}

// WatchSetForwAddress is a free log subscription operation binding the contract event 0x9408dd080d4dfc13f9a203a2fb36089a7c08db7ac37f83da7b473ebfa5881729.
//
// Solidity: event SetForwAddress(address indexed sender, address oldValue, address newValue)
func (_Token *TokenFilterer) WatchSetForwAddress(opts *bind.WatchOpts, sink chan<- *TokenSetForwAddress, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "SetForwAddress", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenSetForwAddress)
				if err := _Token.contract.UnpackLog(event, "SetForwAddress", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetForwAddress is a log parse operation binding the contract event 0x9408dd080d4dfc13f9a203a2fb36089a7c08db7ac37f83da7b473ebfa5881729.
//
// Solidity: event SetForwAddress(address indexed sender, address oldValue, address newValue)
func (_Token *TokenFilterer) ParseSetForwAddress(log types.Log) (*TokenSetForwAddress, error) {
	event := new(TokenSetForwAddress)
	if err := _Token.contract.UnpackLog(event, "SetForwAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenSetForwDistributorAddressIterator is returned from FilterSetForwDistributorAddress and is used to iterate over the raw logs and unpacked data for SetForwDistributorAddress events raised by the Token contract.
type TokenSetForwDistributorAddressIterator struct {
	Event *TokenSetForwDistributorAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenSetForwDistributorAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenSetForwDistributorAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenSetForwDistributorAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenSetForwDistributorAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenSetForwDistributorAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenSetForwDistributorAddress represents a SetForwDistributorAddress event raised by the Token contract.
type TokenSetForwDistributorAddress struct {
	Sender   common.Address
	OldValue common.Address
	NewValue common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetForwDistributorAddress is a free log retrieval operation binding the contract event 0x211bc2b0eb66f6f356f4de49131c7093e7099446899240bf9dffccde4da1f5e5.
//
// Solidity: event SetForwDistributorAddress(address indexed sender, address oldValue, address newValue)
func (_Token *TokenFilterer) FilterSetForwDistributorAddress(opts *bind.FilterOpts, sender []common.Address) (*TokenSetForwDistributorAddressIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "SetForwDistributorAddress", senderRule)
	if err != nil {
		return nil, err
	}
	return &TokenSetForwDistributorAddressIterator{contract: _Token.contract, event: "SetForwDistributorAddress", logs: logs, sub: sub}, nil
}

// WatchSetForwDistributorAddress is a free log subscription operation binding the contract event 0x211bc2b0eb66f6f356f4de49131c7093e7099446899240bf9dffccde4da1f5e5.
//
// Solidity: event SetForwDistributorAddress(address indexed sender, address oldValue, address newValue)
func (_Token *TokenFilterer) WatchSetForwDistributorAddress(opts *bind.WatchOpts, sink chan<- *TokenSetForwDistributorAddress, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "SetForwDistributorAddress", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenSetForwDistributorAddress)
				if err := _Token.contract.UnpackLog(event, "SetForwDistributorAddress", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetForwDistributorAddress is a log parse operation binding the contract event 0x211bc2b0eb66f6f356f4de49131c7093e7099446899240bf9dffccde4da1f5e5.
//
// Solidity: event SetForwDistributorAddress(address indexed sender, address oldValue, address newValue)
func (_Token *TokenFilterer) ParseSetForwDistributorAddress(log types.Log) (*TokenSetForwDistributorAddress, error) {
	event := new(TokenSetForwDistributorAddress)
	if err := _Token.contract.UnpackLog(event, "SetForwDistributorAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenSetFowPerBlockIterator is returned from FilterSetFowPerBlock and is used to iterate over the raw logs and unpacked data for SetFowPerBlock events raised by the Token contract.
type TokenSetFowPerBlockIterator struct {
	Event *TokenSetFowPerBlock // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenSetFowPerBlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenSetFowPerBlock)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenSetFowPerBlock)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenSetFowPerBlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenSetFowPerBlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenSetFowPerBlock represents a SetFowPerBlock event raised by the Token contract.
type TokenSetFowPerBlock struct {
	Sender      common.Address
	Amount      *big.Int
	TargetBlock *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetFowPerBlock is a free log retrieval operation binding the contract event 0xe309d824d705b373af9664078f1d0a958c5809603c53155c69610b532597f075.
//
// Solidity: event SetFowPerBlock(address indexed sender, uint256 amount, uint256 targetBlock)
func (_Token *TokenFilterer) FilterSetFowPerBlock(opts *bind.FilterOpts, sender []common.Address) (*TokenSetFowPerBlockIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "SetFowPerBlock", senderRule)
	if err != nil {
		return nil, err
	}
	return &TokenSetFowPerBlockIterator{contract: _Token.contract, event: "SetFowPerBlock", logs: logs, sub: sub}, nil
}

// WatchSetFowPerBlock is a free log subscription operation binding the contract event 0xe309d824d705b373af9664078f1d0a958c5809603c53155c69610b532597f075.
//
// Solidity: event SetFowPerBlock(address indexed sender, uint256 amount, uint256 targetBlock)
func (_Token *TokenFilterer) WatchSetFowPerBlock(opts *bind.WatchOpts, sink chan<- *TokenSetFowPerBlock, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "SetFowPerBlock", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenSetFowPerBlock)
				if err := _Token.contract.UnpackLog(event, "SetFowPerBlock", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetFowPerBlock is a log parse operation binding the contract event 0xe309d824d705b373af9664078f1d0a958c5809603c53155c69610b532597f075.
//
// Solidity: event SetFowPerBlock(address indexed sender, uint256 amount, uint256 targetBlock)
func (_Token *TokenFilterer) ParseSetFowPerBlock(log types.Log) (*TokenSetFowPerBlock, error) {
	event := new(TokenSetFowPerBlock)
	if err := _Token.contract.UnpackLog(event, "SetFowPerBlock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenSetLoanConfigIterator is returned from FilterSetLoanConfig and is used to iterate over the raw logs and unpacked data for SetLoanConfig events raised by the Token contract.
type TokenSetLoanConfigIterator struct {
	Event *TokenSetLoanConfig // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenSetLoanConfigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenSetLoanConfig)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenSetLoanConfig)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenSetLoanConfigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenSetLoanConfigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenSetLoanConfig represents a SetLoanConfig event raised by the Token contract.
type TokenSetLoanConfig struct {
	Sender                 common.Address
	CollateralTokenAddress common.Address
	SafeLTV                *big.Int
	MaxLTV                 *big.Int
	LiqLTV                 *big.Int
	BountyFeeRate          *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetLoanConfig is a free log retrieval operation binding the contract event 0x75e3d3b9b9cdefcd628c1650fde4c55854e59789c227cb538bc3ada6a48a78c1.
//
// Solidity: event SetLoanConfig(address indexed sender, address collateralTokenAddress, uint256 safeLTV, uint256 maxLTV, uint256 liqLTV, uint256 bountyFeeRate)
func (_Token *TokenFilterer) FilterSetLoanConfig(opts *bind.FilterOpts, sender []common.Address) (*TokenSetLoanConfigIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "SetLoanConfig", senderRule)
	if err != nil {
		return nil, err
	}
	return &TokenSetLoanConfigIterator{contract: _Token.contract, event: "SetLoanConfig", logs: logs, sub: sub}, nil
}

// WatchSetLoanConfig is a free log subscription operation binding the contract event 0x75e3d3b9b9cdefcd628c1650fde4c55854e59789c227cb538bc3ada6a48a78c1.
//
// Solidity: event SetLoanConfig(address indexed sender, address collateralTokenAddress, uint256 safeLTV, uint256 maxLTV, uint256 liqLTV, uint256 bountyFeeRate)
func (_Token *TokenFilterer) WatchSetLoanConfig(opts *bind.WatchOpts, sink chan<- *TokenSetLoanConfig, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "SetLoanConfig", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenSetLoanConfig)
				if err := _Token.contract.UnpackLog(event, "SetLoanConfig", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetLoanConfig is a log parse operation binding the contract event 0x75e3d3b9b9cdefcd628c1650fde4c55854e59789c227cb538bc3ada6a48a78c1.
//
// Solidity: event SetLoanConfig(address indexed sender, address collateralTokenAddress, uint256 safeLTV, uint256 maxLTV, uint256 liqLTV, uint256 bountyFeeRate)
func (_Token *TokenFilterer) ParseSetLoanConfig(log types.Log) (*TokenSetLoanConfig, error) {
	event := new(TokenSetLoanConfig)
	if err := _Token.contract.UnpackLog(event, "SetLoanConfig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenSetLoanDurationIterator is returned from FilterSetLoanDuration and is used to iterate over the raw logs and unpacked data for SetLoanDuration events raised by the Token contract.
type TokenSetLoanDurationIterator struct {
	Event *TokenSetLoanDuration // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenSetLoanDurationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenSetLoanDuration)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenSetLoanDuration)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenSetLoanDurationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenSetLoanDurationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenSetLoanDuration represents a SetLoanDuration event raised by the Token contract.
type TokenSetLoanDuration struct {
	Sender   common.Address
	OldValue *big.Int
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetLoanDuration is a free log retrieval operation binding the contract event 0xc7ae87c625a7efdb53ed6e0019d6d293f8c858c2ef60668f9b1611a25f242e19.
//
// Solidity: event SetLoanDuration(address indexed sender, uint256 oldValue, uint256 newValue)
func (_Token *TokenFilterer) FilterSetLoanDuration(opts *bind.FilterOpts, sender []common.Address) (*TokenSetLoanDurationIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "SetLoanDuration", senderRule)
	if err != nil {
		return nil, err
	}
	return &TokenSetLoanDurationIterator{contract: _Token.contract, event: "SetLoanDuration", logs: logs, sub: sub}, nil
}

// WatchSetLoanDuration is a free log subscription operation binding the contract event 0xc7ae87c625a7efdb53ed6e0019d6d293f8c858c2ef60668f9b1611a25f242e19.
//
// Solidity: event SetLoanDuration(address indexed sender, uint256 oldValue, uint256 newValue)
func (_Token *TokenFilterer) WatchSetLoanDuration(opts *bind.WatchOpts, sink chan<- *TokenSetLoanDuration, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "SetLoanDuration", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenSetLoanDuration)
				if err := _Token.contract.UnpackLog(event, "SetLoanDuration", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetLoanDuration is a log parse operation binding the contract event 0xc7ae87c625a7efdb53ed6e0019d6d293f8c858c2ef60668f9b1611a25f242e19.
//
// Solidity: event SetLoanDuration(address indexed sender, uint256 oldValue, uint256 newValue)
func (_Token *TokenFilterer) ParseSetLoanDuration(log types.Log) (*TokenSetLoanDuration, error) {
	event := new(TokenSetLoanDuration)
	if err := _Token.contract.UnpackLog(event, "SetLoanDuration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenSetMembershipAddressIterator is returned from FilterSetMembershipAddress and is used to iterate over the raw logs and unpacked data for SetMembershipAddress events raised by the Token contract.
type TokenSetMembershipAddressIterator struct {
	Event *TokenSetMembershipAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenSetMembershipAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenSetMembershipAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenSetMembershipAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenSetMembershipAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenSetMembershipAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenSetMembershipAddress represents a SetMembershipAddress event raised by the Token contract.
type TokenSetMembershipAddress struct {
	Sender   common.Address
	OldValue common.Address
	NewValue common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetMembershipAddress is a free log retrieval operation binding the contract event 0xbedf1133392cf3b0f12efdb8143d5251179e92b6801df7c742789914196956f5.
//
// Solidity: event SetMembershipAddress(address indexed sender, address oldValue, address newValue)
func (_Token *TokenFilterer) FilterSetMembershipAddress(opts *bind.FilterOpts, sender []common.Address) (*TokenSetMembershipAddressIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "SetMembershipAddress", senderRule)
	if err != nil {
		return nil, err
	}
	return &TokenSetMembershipAddressIterator{contract: _Token.contract, event: "SetMembershipAddress", logs: logs, sub: sub}, nil
}

// WatchSetMembershipAddress is a free log subscription operation binding the contract event 0xbedf1133392cf3b0f12efdb8143d5251179e92b6801df7c742789914196956f5.
//
// Solidity: event SetMembershipAddress(address indexed sender, address oldValue, address newValue)
func (_Token *TokenFilterer) WatchSetMembershipAddress(opts *bind.WatchOpts, sink chan<- *TokenSetMembershipAddress, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "SetMembershipAddress", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenSetMembershipAddress)
				if err := _Token.contract.UnpackLog(event, "SetMembershipAddress", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetMembershipAddress is a log parse operation binding the contract event 0xbedf1133392cf3b0f12efdb8143d5251179e92b6801df7c742789914196956f5.
//
// Solidity: event SetMembershipAddress(address indexed sender, address oldValue, address newValue)
func (_Token *TokenFilterer) ParseSetMembershipAddress(log types.Log) (*TokenSetMembershipAddress, error) {
	event := new(TokenSetMembershipAddress)
	if err := _Token.contract.UnpackLog(event, "SetMembershipAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenSetPoolBorrowingAddressIterator is returned from FilterSetPoolBorrowingAddress and is used to iterate over the raw logs and unpacked data for SetPoolBorrowingAddress events raised by the Token contract.
type TokenSetPoolBorrowingAddressIterator struct {
	Event *TokenSetPoolBorrowingAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenSetPoolBorrowingAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenSetPoolBorrowingAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenSetPoolBorrowingAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenSetPoolBorrowingAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenSetPoolBorrowingAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenSetPoolBorrowingAddress represents a SetPoolBorrowingAddress event raised by the Token contract.
type TokenSetPoolBorrowingAddress struct {
	Sender   common.Address
	OldValue common.Address
	NewValue common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetPoolBorrowingAddress is a free log retrieval operation binding the contract event 0xe7b1881df53f9e26241f97afbc45ffed3bbeb29b402e83cd10e6bdc51a3ece81.
//
// Solidity: event SetPoolBorrowingAddress(address indexed sender, address oldValue, address newValue)
func (_Token *TokenFilterer) FilterSetPoolBorrowingAddress(opts *bind.FilterOpts, sender []common.Address) (*TokenSetPoolBorrowingAddressIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "SetPoolBorrowingAddress", senderRule)
	if err != nil {
		return nil, err
	}
	return &TokenSetPoolBorrowingAddressIterator{contract: _Token.contract, event: "SetPoolBorrowingAddress", logs: logs, sub: sub}, nil
}

// WatchSetPoolBorrowingAddress is a free log subscription operation binding the contract event 0xe7b1881df53f9e26241f97afbc45ffed3bbeb29b402e83cd10e6bdc51a3ece81.
//
// Solidity: event SetPoolBorrowingAddress(address indexed sender, address oldValue, address newValue)
func (_Token *TokenFilterer) WatchSetPoolBorrowingAddress(opts *bind.WatchOpts, sink chan<- *TokenSetPoolBorrowingAddress, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "SetPoolBorrowingAddress", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenSetPoolBorrowingAddress)
				if err := _Token.contract.UnpackLog(event, "SetPoolBorrowingAddress", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetPoolBorrowingAddress is a log parse operation binding the contract event 0xe7b1881df53f9e26241f97afbc45ffed3bbeb29b402e83cd10e6bdc51a3ece81.
//
// Solidity: event SetPoolBorrowingAddress(address indexed sender, address oldValue, address newValue)
func (_Token *TokenFilterer) ParseSetPoolBorrowingAddress(log types.Log) (*TokenSetPoolBorrowingAddress, error) {
	event := new(TokenSetPoolBorrowingAddress)
	if err := _Token.contract.UnpackLog(event, "SetPoolBorrowingAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenSetPoolLendingAddressIterator is returned from FilterSetPoolLendingAddress and is used to iterate over the raw logs and unpacked data for SetPoolLendingAddress events raised by the Token contract.
type TokenSetPoolLendingAddressIterator struct {
	Event *TokenSetPoolLendingAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenSetPoolLendingAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenSetPoolLendingAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenSetPoolLendingAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenSetPoolLendingAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenSetPoolLendingAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenSetPoolLendingAddress represents a SetPoolLendingAddress event raised by the Token contract.
type TokenSetPoolLendingAddress struct {
	Sender   common.Address
	OldValue common.Address
	NewValue common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetPoolLendingAddress is a free log retrieval operation binding the contract event 0x803ded723024ae6142163988a2710414ee9871abc2eb2dcd18c995638ac84fce.
//
// Solidity: event SetPoolLendingAddress(address indexed sender, address oldValue, address newValue)
func (_Token *TokenFilterer) FilterSetPoolLendingAddress(opts *bind.FilterOpts, sender []common.Address) (*TokenSetPoolLendingAddressIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "SetPoolLendingAddress", senderRule)
	if err != nil {
		return nil, err
	}
	return &TokenSetPoolLendingAddressIterator{contract: _Token.contract, event: "SetPoolLendingAddress", logs: logs, sub: sub}, nil
}

// WatchSetPoolLendingAddress is a free log subscription operation binding the contract event 0x803ded723024ae6142163988a2710414ee9871abc2eb2dcd18c995638ac84fce.
//
// Solidity: event SetPoolLendingAddress(address indexed sender, address oldValue, address newValue)
func (_Token *TokenFilterer) WatchSetPoolLendingAddress(opts *bind.WatchOpts, sink chan<- *TokenSetPoolLendingAddress, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "SetPoolLendingAddress", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenSetPoolLendingAddress)
				if err := _Token.contract.UnpackLog(event, "SetPoolLendingAddress", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetPoolLendingAddress is a log parse operation binding the contract event 0x803ded723024ae6142163988a2710414ee9871abc2eb2dcd18c995638ac84fce.
//
// Solidity: event SetPoolLendingAddress(address indexed sender, address oldValue, address newValue)
func (_Token *TokenFilterer) ParseSetPoolLendingAddress(log types.Log) (*TokenSetPoolLendingAddress, error) {
	event := new(TokenSetPoolLendingAddress)
	if err := _Token.contract.UnpackLog(event, "SetPoolLendingAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenSetPoolStartTimestampIterator is returned from FilterSetPoolStartTimestamp and is used to iterate over the raw logs and unpacked data for SetPoolStartTimestamp events raised by the Token contract.
type TokenSetPoolStartTimestampIterator struct {
	Event *TokenSetPoolStartTimestamp // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenSetPoolStartTimestampIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenSetPoolStartTimestamp)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenSetPoolStartTimestamp)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenSetPoolStartTimestampIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenSetPoolStartTimestampIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenSetPoolStartTimestamp represents a SetPoolStartTimestamp event raised by the Token contract.
type TokenSetPoolStartTimestamp struct {
	Sender    common.Address
	Timestamp uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetPoolStartTimestamp is a free log retrieval operation binding the contract event 0x4d4ba4883f5abdfaeed8a27b79e47a0d023000105b13c3c4006f4cf2d66a0a7a.
//
// Solidity: event SetPoolStartTimestamp(address indexed sender, uint64 indexed timestamp)
func (_Token *TokenFilterer) FilterSetPoolStartTimestamp(opts *bind.FilterOpts, sender []common.Address, timestamp []uint64) (*TokenSetPoolStartTimestampIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "SetPoolStartTimestamp", senderRule, timestampRule)
	if err != nil {
		return nil, err
	}
	return &TokenSetPoolStartTimestampIterator{contract: _Token.contract, event: "SetPoolStartTimestamp", logs: logs, sub: sub}, nil
}

// WatchSetPoolStartTimestamp is a free log subscription operation binding the contract event 0x4d4ba4883f5abdfaeed8a27b79e47a0d023000105b13c3c4006f4cf2d66a0a7a.
//
// Solidity: event SetPoolStartTimestamp(address indexed sender, uint64 indexed timestamp)
func (_Token *TokenFilterer) WatchSetPoolStartTimestamp(opts *bind.WatchOpts, sink chan<- *TokenSetPoolStartTimestamp, sender []common.Address, timestamp []uint64) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "SetPoolStartTimestamp", senderRule, timestampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenSetPoolStartTimestamp)
				if err := _Token.contract.UnpackLog(event, "SetPoolStartTimestamp", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetPoolStartTimestamp is a log parse operation binding the contract event 0x4d4ba4883f5abdfaeed8a27b79e47a0d023000105b13c3c4006f4cf2d66a0a7a.
//
// Solidity: event SetPoolStartTimestamp(address indexed sender, uint64 indexed timestamp)
func (_Token *TokenFilterer) ParseSetPoolStartTimestamp(log types.Log) (*TokenSetPoolStartTimestamp, error) {
	event := new(TokenSetPoolStartTimestamp)
	if err := _Token.contract.UnpackLog(event, "SetPoolStartTimestamp", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenSetPriceFeedIterator is returned from FilterSetPriceFeed and is used to iterate over the raw logs and unpacked data for SetPriceFeed events raised by the Token contract.
type TokenSetPriceFeedIterator struct {
	Event *TokenSetPriceFeed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenSetPriceFeedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenSetPriceFeed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenSetPriceFeed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenSetPriceFeedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenSetPriceFeedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenSetPriceFeed represents a SetPriceFeed event raised by the Token contract.
type TokenSetPriceFeed struct {
	Sender common.Address
	Tokens []common.Address
	Feeds  []common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetPriceFeed is a free log retrieval operation binding the contract event 0xc9316e2440af9aa4bb5f3306905f06f644ab8ab8dd739f4003784ceab951ddde.
//
// Solidity: event SetPriceFeed(address indexed sender, address[] tokens, address[] feeds)
func (_Token *TokenFilterer) FilterSetPriceFeed(opts *bind.FilterOpts, sender []common.Address) (*TokenSetPriceFeedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "SetPriceFeed", senderRule)
	if err != nil {
		return nil, err
	}
	return &TokenSetPriceFeedIterator{contract: _Token.contract, event: "SetPriceFeed", logs: logs, sub: sub}, nil
}

// WatchSetPriceFeed is a free log subscription operation binding the contract event 0xc9316e2440af9aa4bb5f3306905f06f644ab8ab8dd739f4003784ceab951ddde.
//
// Solidity: event SetPriceFeed(address indexed sender, address[] tokens, address[] feeds)
func (_Token *TokenFilterer) WatchSetPriceFeed(opts *bind.WatchOpts, sink chan<- *TokenSetPriceFeed, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "SetPriceFeed", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenSetPriceFeed)
				if err := _Token.contract.UnpackLog(event, "SetPriceFeed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetPriceFeed is a log parse operation binding the contract event 0xc9316e2440af9aa4bb5f3306905f06f644ab8ab8dd739f4003784ceab951ddde.
//
// Solidity: event SetPriceFeed(address indexed sender, address[] tokens, address[] feeds)
func (_Token *TokenFilterer) ParseSetPriceFeed(log types.Log) (*TokenSetPriceFeed, error) {
	event := new(TokenSetPriceFeed)
	if err := _Token.contract.UnpackLog(event, "SetPriceFeed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenSetPriceFeedAddressIterator is returned from FilterSetPriceFeedAddress and is used to iterate over the raw logs and unpacked data for SetPriceFeedAddress events raised by the Token contract.
type TokenSetPriceFeedAddressIterator struct {
	Event *TokenSetPriceFeedAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenSetPriceFeedAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenSetPriceFeedAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenSetPriceFeedAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenSetPriceFeedAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenSetPriceFeedAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenSetPriceFeedAddress represents a SetPriceFeedAddress event raised by the Token contract.
type TokenSetPriceFeedAddress struct {
	Sender   common.Address
	OldValue common.Address
	NewValue common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetPriceFeedAddress is a free log retrieval operation binding the contract event 0x2d69e04509dea85debffb6b25e9d18765603974aab0ec79abc2e4824d6fd4fe8.
//
// Solidity: event SetPriceFeedAddress(address indexed sender, address oldValue, address newValue)
func (_Token *TokenFilterer) FilterSetPriceFeedAddress(opts *bind.FilterOpts, sender []common.Address) (*TokenSetPriceFeedAddressIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "SetPriceFeedAddress", senderRule)
	if err != nil {
		return nil, err
	}
	return &TokenSetPriceFeedAddressIterator{contract: _Token.contract, event: "SetPriceFeedAddress", logs: logs, sub: sub}, nil
}

// WatchSetPriceFeedAddress is a free log subscription operation binding the contract event 0x2d69e04509dea85debffb6b25e9d18765603974aab0ec79abc2e4824d6fd4fe8.
//
// Solidity: event SetPriceFeedAddress(address indexed sender, address oldValue, address newValue)
func (_Token *TokenFilterer) WatchSetPriceFeedAddress(opts *bind.WatchOpts, sink chan<- *TokenSetPriceFeedAddress, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "SetPriceFeedAddress", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenSetPriceFeedAddress)
				if err := _Token.contract.UnpackLog(event, "SetPriceFeedAddress", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetPriceFeedAddress is a log parse operation binding the contract event 0x2d69e04509dea85debffb6b25e9d18765603974aab0ec79abc2e4824d6fd4fe8.
//
// Solidity: event SetPriceFeedAddress(address indexed sender, address oldValue, address newValue)
func (_Token *TokenFilterer) ParseSetPriceFeedAddress(log types.Log) (*TokenSetPriceFeedAddress, error) {
	event := new(TokenSetPriceFeedAddress)
	if err := _Token.contract.UnpackLog(event, "SetPriceFeedAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenSetProtocolAddressIterator is returned from FilterSetProtocolAddress and is used to iterate over the raw logs and unpacked data for SetProtocolAddress events raised by the Token contract.
type TokenSetProtocolAddressIterator struct {
	Event *TokenSetProtocolAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenSetProtocolAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenSetProtocolAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenSetProtocolAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenSetProtocolAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenSetProtocolAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenSetProtocolAddress represents a SetProtocolAddress event raised by the Token contract.
type TokenSetProtocolAddress struct {
	Sender   common.Address
	OldValue common.Address
	NewValue common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetProtocolAddress is a free log retrieval operation binding the contract event 0xe48b1893efaee62da0749dd98dbf0ea62ee472b37f7205114aca36cc9de11228.
//
// Solidity: event SetProtocolAddress(address indexed sender, address oldValue, address newValue)
func (_Token *TokenFilterer) FilterSetProtocolAddress(opts *bind.FilterOpts, sender []common.Address) (*TokenSetProtocolAddressIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "SetProtocolAddress", senderRule)
	if err != nil {
		return nil, err
	}
	return &TokenSetProtocolAddressIterator{contract: _Token.contract, event: "SetProtocolAddress", logs: logs, sub: sub}, nil
}

// WatchSetProtocolAddress is a free log subscription operation binding the contract event 0xe48b1893efaee62da0749dd98dbf0ea62ee472b37f7205114aca36cc9de11228.
//
// Solidity: event SetProtocolAddress(address indexed sender, address oldValue, address newValue)
func (_Token *TokenFilterer) WatchSetProtocolAddress(opts *bind.WatchOpts, sink chan<- *TokenSetProtocolAddress, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "SetProtocolAddress", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenSetProtocolAddress)
				if err := _Token.contract.UnpackLog(event, "SetProtocolAddress", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetProtocolAddress is a log parse operation binding the contract event 0xe48b1893efaee62da0749dd98dbf0ea62ee472b37f7205114aca36cc9de11228.
//
// Solidity: event SetProtocolAddress(address indexed sender, address oldValue, address newValue)
func (_Token *TokenFilterer) ParseSetProtocolAddress(log types.Log) (*TokenSetProtocolAddress, error) {
	event := new(TokenSetProtocolAddress)
	if err := _Token.contract.UnpackLog(event, "SetProtocolAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenSetRouterAddressIterator is returned from FilterSetRouterAddress and is used to iterate over the raw logs and unpacked data for SetRouterAddress events raised by the Token contract.
type TokenSetRouterAddressIterator struct {
	Event *TokenSetRouterAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenSetRouterAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenSetRouterAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenSetRouterAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenSetRouterAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenSetRouterAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenSetRouterAddress represents a SetRouterAddress event raised by the Token contract.
type TokenSetRouterAddress struct {
	Sender   common.Address
	OldValue common.Address
	NewValue common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetRouterAddress is a free log retrieval operation binding the contract event 0x8bddec4e99c28379744f9203cc15872c7a677c12475433523bc81568721159f4.
//
// Solidity: event SetRouterAddress(address indexed sender, address oldValue, address newValue)
func (_Token *TokenFilterer) FilterSetRouterAddress(opts *bind.FilterOpts, sender []common.Address) (*TokenSetRouterAddressIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "SetRouterAddress", senderRule)
	if err != nil {
		return nil, err
	}
	return &TokenSetRouterAddressIterator{contract: _Token.contract, event: "SetRouterAddress", logs: logs, sub: sub}, nil
}

// WatchSetRouterAddress is a free log subscription operation binding the contract event 0x8bddec4e99c28379744f9203cc15872c7a677c12475433523bc81568721159f4.
//
// Solidity: event SetRouterAddress(address indexed sender, address oldValue, address newValue)
func (_Token *TokenFilterer) WatchSetRouterAddress(opts *bind.WatchOpts, sink chan<- *TokenSetRouterAddress, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "SetRouterAddress", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenSetRouterAddress)
				if err := _Token.contract.UnpackLog(event, "SetRouterAddress", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetRouterAddress is a log parse operation binding the contract event 0x8bddec4e99c28379744f9203cc15872c7a677c12475433523bc81568721159f4.
//
// Solidity: event SetRouterAddress(address indexed sender, address oldValue, address newValue)
func (_Token *TokenFilterer) ParseSetRouterAddress(log types.Log) (*TokenSetRouterAddress, error) {
	event := new(TokenSetRouterAddress)
	if err := _Token.contract.UnpackLog(event, "SetRouterAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenSetTokenAddressIterator is returned from FilterSetTokenAddress and is used to iterate over the raw logs and unpacked data for SetTokenAddress events raised by the Token contract.
type TokenSetTokenAddressIterator struct {
	Event *TokenSetTokenAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenSetTokenAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenSetTokenAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenSetTokenAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenSetTokenAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenSetTokenAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenSetTokenAddress represents a SetTokenAddress event raised by the Token contract.
type TokenSetTokenAddress struct {
	Sender   common.Address
	OldValue common.Address
	NewValue common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetTokenAddress is a free log retrieval operation binding the contract event 0xe1284317042ff70cf6435810b203435d61e9b2b4ea6dcd1b4545662d7a246390.
//
// Solidity: event SetTokenAddress(address indexed sender, address oldValue, address newValue)
func (_Token *TokenFilterer) FilterSetTokenAddress(opts *bind.FilterOpts, sender []common.Address) (*TokenSetTokenAddressIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "SetTokenAddress", senderRule)
	if err != nil {
		return nil, err
	}
	return &TokenSetTokenAddressIterator{contract: _Token.contract, event: "SetTokenAddress", logs: logs, sub: sub}, nil
}

// WatchSetTokenAddress is a free log subscription operation binding the contract event 0xe1284317042ff70cf6435810b203435d61e9b2b4ea6dcd1b4545662d7a246390.
//
// Solidity: event SetTokenAddress(address indexed sender, address oldValue, address newValue)
func (_Token *TokenFilterer) WatchSetTokenAddress(opts *bind.WatchOpts, sink chan<- *TokenSetTokenAddress, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "SetTokenAddress", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenSetTokenAddress)
				if err := _Token.contract.UnpackLog(event, "SetTokenAddress", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetTokenAddress is a log parse operation binding the contract event 0xe1284317042ff70cf6435810b203435d61e9b2b4ea6dcd1b4545662d7a246390.
//
// Solidity: event SetTokenAddress(address indexed sender, address oldValue, address newValue)
func (_Token *TokenFilterer) ParseSetTokenAddress(log types.Log) (*TokenSetTokenAddress, error) {
	event := new(TokenSetTokenAddress)
	if err := _Token.contract.UnpackLog(event, "SetTokenAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenSetWETHHandlerIterator is returned from FilterSetWETHHandler and is used to iterate over the raw logs and unpacked data for SetWETHHandler events raised by the Token contract.
type TokenSetWETHHandlerIterator struct {
	Event *TokenSetWETHHandler // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenSetWETHHandlerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenSetWETHHandler)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenSetWETHHandler)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenSetWETHHandlerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenSetWETHHandlerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenSetWETHHandler represents a SetWETHHandler event raised by the Token contract.
type TokenSetWETHHandler struct {
	Sender   common.Address
	OldValue common.Address
	NewValue common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetWETHHandler is a free log retrieval operation binding the contract event 0xaf1a9693e87f7b470cf82b45afae69388908f2a8a0ae6b8a613206a197cd3230.
//
// Solidity: event SetWETHHandler(address indexed sender, address oldValue, address newValue)
func (_Token *TokenFilterer) FilterSetWETHHandler(opts *bind.FilterOpts, sender []common.Address) (*TokenSetWETHHandlerIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "SetWETHHandler", senderRule)
	if err != nil {
		return nil, err
	}
	return &TokenSetWETHHandlerIterator{contract: _Token.contract, event: "SetWETHHandler", logs: logs, sub: sub}, nil
}

// WatchSetWETHHandler is a free log subscription operation binding the contract event 0xaf1a9693e87f7b470cf82b45afae69388908f2a8a0ae6b8a613206a197cd3230.
//
// Solidity: event SetWETHHandler(address indexed sender, address oldValue, address newValue)
func (_Token *TokenFilterer) WatchSetWETHHandler(opts *bind.WatchOpts, sink chan<- *TokenSetWETHHandler, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "SetWETHHandler", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenSetWETHHandler)
				if err := _Token.contract.UnpackLog(event, "SetWETHHandler", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetWETHHandler is a log parse operation binding the contract event 0xaf1a9693e87f7b470cf82b45afae69388908f2a8a0ae6b8a613206a197cd3230.
//
// Solidity: event SetWETHHandler(address indexed sender, address oldValue, address newValue)
func (_Token *TokenFilterer) ParseSetWETHHandler(log types.Log) (*TokenSetWETHHandler, error) {
	event := new(TokenSetWETHHandler)
	if err := _Token.contract.UnpackLog(event, "SetWETHHandler", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenSettleForwInterestIterator is returned from FilterSettleForwInterest and is used to iterate over the raw logs and unpacked data for SettleForwInterest events raised by the Token contract.
type TokenSettleForwInterestIterator struct {
	Event *TokenSettleForwInterest // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenSettleForwInterestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenSettleForwInterest)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenSettleForwInterest)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenSettleForwInterestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenSettleForwInterestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenSettleForwInterest represents a SettleForwInterest event raised by the Token contract.
type TokenSettleForwInterest struct {
	CoreAddress             common.Address
	InterestVaultAddress    common.Address
	ForwDistributionAddress common.Address
	ForwTokenAddress        common.Address
	Amount                  *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterSettleForwInterest is a free log retrieval operation binding the contract event 0x9e85554bb610e8a83aef1532606338530ce65077ca3b596c95e72aba0ec3fd95.
//
// Solidity: event SettleForwInterest(address indexed coreAddress, address indexed interestVaultAddress, address forwDistributionAddress, address forwTokenAddress, uint256 amount)
func (_Token *TokenFilterer) FilterSettleForwInterest(opts *bind.FilterOpts, coreAddress []common.Address, interestVaultAddress []common.Address) (*TokenSettleForwInterestIterator, error) {

	var coreAddressRule []interface{}
	for _, coreAddressItem := range coreAddress {
		coreAddressRule = append(coreAddressRule, coreAddressItem)
	}
	var interestVaultAddressRule []interface{}
	for _, interestVaultAddressItem := range interestVaultAddress {
		interestVaultAddressRule = append(interestVaultAddressRule, interestVaultAddressItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "SettleForwInterest", coreAddressRule, interestVaultAddressRule)
	if err != nil {
		return nil, err
	}
	return &TokenSettleForwInterestIterator{contract: _Token.contract, event: "SettleForwInterest", logs: logs, sub: sub}, nil
}

// WatchSettleForwInterest is a free log subscription operation binding the contract event 0x9e85554bb610e8a83aef1532606338530ce65077ca3b596c95e72aba0ec3fd95.
//
// Solidity: event SettleForwInterest(address indexed coreAddress, address indexed interestVaultAddress, address forwDistributionAddress, address forwTokenAddress, uint256 amount)
func (_Token *TokenFilterer) WatchSettleForwInterest(opts *bind.WatchOpts, sink chan<- *TokenSettleForwInterest, coreAddress []common.Address, interestVaultAddress []common.Address) (event.Subscription, error) {

	var coreAddressRule []interface{}
	for _, coreAddressItem := range coreAddress {
		coreAddressRule = append(coreAddressRule, coreAddressItem)
	}
	var interestVaultAddressRule []interface{}
	for _, interestVaultAddressItem := range interestVaultAddress {
		interestVaultAddressRule = append(interestVaultAddressRule, interestVaultAddressItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "SettleForwInterest", coreAddressRule, interestVaultAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenSettleForwInterest)
				if err := _Token.contract.UnpackLog(event, "SettleForwInterest", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSettleForwInterest is a log parse operation binding the contract event 0x9e85554bb610e8a83aef1532606338530ce65077ca3b596c95e72aba0ec3fd95.
//
// Solidity: event SettleForwInterest(address indexed coreAddress, address indexed interestVaultAddress, address forwDistributionAddress, address forwTokenAddress, uint256 amount)
func (_Token *TokenFilterer) ParseSettleForwInterest(log types.Log) (*TokenSettleForwInterest, error) {
	event := new(TokenSettleForwInterest)
	if err := _Token.contract.UnpackLog(event, "SettleForwInterest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenSettleInterestIterator is returned from FilterSettleInterest and is used to iterate over the raw logs and unpacked data for SettleInterest events raised by the Token contract.
type TokenSettleInterestIterator struct {
	Event *TokenSettleInterest // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenSettleInterestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenSettleInterest)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenSettleInterest)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenSettleInterestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenSettleInterestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenSettleInterest represents a SettleInterest event raised by the Token contract.
type TokenSettleInterest struct {
	Sender                 common.Address
	ClaimableTokenInterest *big.Int
	HeldTokenInterest      *big.Int
	ClaimableForwInterest  *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSettleInterest is a free log retrieval operation binding the contract event 0xb5b31e0ec63ec6588d8e4132641011a6ef764401d1023187b22cb4a76f15f1cb.
//
// Solidity: event SettleInterest(address indexed sender, uint256 claimableTokenInterest, uint256 heldTokenInterest, uint256 claimableForwInterest)
func (_Token *TokenFilterer) FilterSettleInterest(opts *bind.FilterOpts, sender []common.Address) (*TokenSettleInterestIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "SettleInterest", senderRule)
	if err != nil {
		return nil, err
	}
	return &TokenSettleInterestIterator{contract: _Token.contract, event: "SettleInterest", logs: logs, sub: sub}, nil
}

// WatchSettleInterest is a free log subscription operation binding the contract event 0xb5b31e0ec63ec6588d8e4132641011a6ef764401d1023187b22cb4a76f15f1cb.
//
// Solidity: event SettleInterest(address indexed sender, uint256 claimableTokenInterest, uint256 heldTokenInterest, uint256 claimableForwInterest)
func (_Token *TokenFilterer) WatchSettleInterest(opts *bind.WatchOpts, sink chan<- *TokenSettleInterest, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "SettleInterest", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenSettleInterest)
				if err := _Token.contract.UnpackLog(event, "SettleInterest", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSettleInterest is a log parse operation binding the contract event 0xb5b31e0ec63ec6588d8e4132641011a6ef764401d1023187b22cb4a76f15f1cb.
//
// Solidity: event SettleInterest(address indexed sender, uint256 claimableTokenInterest, uint256 heldTokenInterest, uint256 claimableForwInterest)
func (_Token *TokenFilterer) ParseSettleInterest(log types.Log) (*TokenSettleInterest, error) {
	event := new(TokenSettleInterest)
	if err := _Token.contract.UnpackLog(event, "SettleInterest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenSetupLoanConfigIterator is returned from FilterSetupLoanConfig and is used to iterate over the raw logs and unpacked data for SetupLoanConfig events raised by the Token contract.
type TokenSetupLoanConfigIterator struct {
	Event *TokenSetupLoanConfig // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenSetupLoanConfigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenSetupLoanConfig)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenSetupLoanConfig)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenSetupLoanConfigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenSetupLoanConfigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenSetupLoanConfig represents a SetupLoanConfig event raised by the Token contract.
type TokenSetupLoanConfig struct {
	Sender                 common.Address
	BorrowTokenAddress     common.Address
	CollateralTokenAddress common.Address
	OldSafeLTV             *big.Int
	OldMaxLTV              *big.Int
	OldLiquidationLTV      *big.Int
	OldBountyFeeRate       *big.Int
	NewSafeLTV             *big.Int
	NewLMaxLTV             *big.Int
	NewLiquidationLTV      *big.Int
	NewBountyFeeRate       *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetupLoanConfig is a free log retrieval operation binding the contract event 0x34c34d33b7836f5da8474a70f9c0dc1345bbcf0fc7a192f12ee36ad7b2ca3dac.
//
// Solidity: event SetupLoanConfig(address indexed sender, address indexed borrowTokenAddress, address indexed collateralTokenAddress, uint256 oldSafeLTV, uint256 oldMaxLTV, uint256 oldLiquidationLTV, uint256 oldBountyFeeRate, uint256 newSafeLTV, uint256 newLMaxLTV, uint256 newLiquidationLTV, uint256 newBountyFeeRate)
func (_Token *TokenFilterer) FilterSetupLoanConfig(opts *bind.FilterOpts, sender []common.Address, borrowTokenAddress []common.Address, collateralTokenAddress []common.Address) (*TokenSetupLoanConfigIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var borrowTokenAddressRule []interface{}
	for _, borrowTokenAddressItem := range borrowTokenAddress {
		borrowTokenAddressRule = append(borrowTokenAddressRule, borrowTokenAddressItem)
	}
	var collateralTokenAddressRule []interface{}
	for _, collateralTokenAddressItem := range collateralTokenAddress {
		collateralTokenAddressRule = append(collateralTokenAddressRule, collateralTokenAddressItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "SetupLoanConfig", senderRule, borrowTokenAddressRule, collateralTokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &TokenSetupLoanConfigIterator{contract: _Token.contract, event: "SetupLoanConfig", logs: logs, sub: sub}, nil
}

// WatchSetupLoanConfig is a free log subscription operation binding the contract event 0x34c34d33b7836f5da8474a70f9c0dc1345bbcf0fc7a192f12ee36ad7b2ca3dac.
//
// Solidity: event SetupLoanConfig(address indexed sender, address indexed borrowTokenAddress, address indexed collateralTokenAddress, uint256 oldSafeLTV, uint256 oldMaxLTV, uint256 oldLiquidationLTV, uint256 oldBountyFeeRate, uint256 newSafeLTV, uint256 newLMaxLTV, uint256 newLiquidationLTV, uint256 newBountyFeeRate)
func (_Token *TokenFilterer) WatchSetupLoanConfig(opts *bind.WatchOpts, sink chan<- *TokenSetupLoanConfig, sender []common.Address, borrowTokenAddress []common.Address, collateralTokenAddress []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var borrowTokenAddressRule []interface{}
	for _, borrowTokenAddressItem := range borrowTokenAddress {
		borrowTokenAddressRule = append(borrowTokenAddressRule, borrowTokenAddressItem)
	}
	var collateralTokenAddressRule []interface{}
	for _, collateralTokenAddressItem := range collateralTokenAddress {
		collateralTokenAddressRule = append(collateralTokenAddressRule, collateralTokenAddressItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "SetupLoanConfig", senderRule, borrowTokenAddressRule, collateralTokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenSetupLoanConfig)
				if err := _Token.contract.UnpackLog(event, "SetupLoanConfig", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetupLoanConfig is a log parse operation binding the contract event 0x34c34d33b7836f5da8474a70f9c0dc1345bbcf0fc7a192f12ee36ad7b2ca3dac.
//
// Solidity: event SetupLoanConfig(address indexed sender, address indexed borrowTokenAddress, address indexed collateralTokenAddress, uint256 oldSafeLTV, uint256 oldMaxLTV, uint256 oldLiquidationLTV, uint256 oldBountyFeeRate, uint256 newSafeLTV, uint256 newLMaxLTV, uint256 newLiquidationLTV, uint256 newBountyFeeRate)
func (_Token *TokenFilterer) ParseSetupLoanConfig(log types.Log) (*TokenSetupLoanConfig, error) {
	event := new(TokenSetupLoanConfig)
	if err := _Token.contract.UnpackLog(event, "SetupLoanConfig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenStakeIterator is returned from FilterStake and is used to iterate over the raw logs and unpacked data for Stake events raised by the Token contract.
type TokenStakeIterator struct {
	Event *TokenStake // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenStake)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenStake)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenStake represents a Stake event raised by the Token contract.
type TokenStake struct {
	Sender common.Address
	NftId  *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStake is a free log retrieval operation binding the contract event 0x5af417134f72a9d41143ace85b0a26dce6f550f894f2cbc1eeee8810603d91b6.
//
// Solidity: event Stake(address indexed sender, uint256 indexed nftId, uint256 amount)
func (_Token *TokenFilterer) FilterStake(opts *bind.FilterOpts, sender []common.Address, nftId []*big.Int) (*TokenStakeIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Stake", senderRule, nftIdRule)
	if err != nil {
		return nil, err
	}
	return &TokenStakeIterator{contract: _Token.contract, event: "Stake", logs: logs, sub: sub}, nil
}

// WatchStake is a free log subscription operation binding the contract event 0x5af417134f72a9d41143ace85b0a26dce6f550f894f2cbc1eeee8810603d91b6.
//
// Solidity: event Stake(address indexed sender, uint256 indexed nftId, uint256 amount)
func (_Token *TokenFilterer) WatchStake(opts *bind.WatchOpts, sink chan<- *TokenStake, sender []common.Address, nftId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Stake", senderRule, nftIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenStake)
				if err := _Token.contract.UnpackLog(event, "Stake", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStake is a log parse operation binding the contract event 0x5af417134f72a9d41143ace85b0a26dce6f550f894f2cbc1eeee8810603d91b6.
//
// Solidity: event Stake(address indexed sender, uint256 indexed nftId, uint256 amount)
func (_Token *TokenFilterer) ParseStake(log types.Log) (*TokenStake, error) {
	event := new(TokenStake)
	if err := _Token.contract.UnpackLog(event, "Stake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenUnStakeIterator is returned from FilterUnStake and is used to iterate over the raw logs and unpacked data for UnStake events raised by the Token contract.
type TokenUnStakeIterator struct {
	Event *TokenUnStake // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenUnStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenUnStake)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenUnStake)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenUnStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenUnStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenUnStake represents a UnStake event raised by the Token contract.
type TokenUnStake struct {
	Sender common.Address
	NftId  *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnStake is a free log retrieval operation binding the contract event 0x54a9763035584fc4fcad1bc4e0e7a83f93e016f50ae32bd527530a77257393ee.
//
// Solidity: event UnStake(address indexed sender, uint256 indexed nftId, uint256 amount)
func (_Token *TokenFilterer) FilterUnStake(opts *bind.FilterOpts, sender []common.Address, nftId []*big.Int) (*TokenUnStakeIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "UnStake", senderRule, nftIdRule)
	if err != nil {
		return nil, err
	}
	return &TokenUnStakeIterator{contract: _Token.contract, event: "UnStake", logs: logs, sub: sub}, nil
}

// WatchUnStake is a free log subscription operation binding the contract event 0x54a9763035584fc4fcad1bc4e0e7a83f93e016f50ae32bd527530a77257393ee.
//
// Solidity: event UnStake(address indexed sender, uint256 indexed nftId, uint256 amount)
func (_Token *TokenFilterer) WatchUnStake(opts *bind.WatchOpts, sink chan<- *TokenUnStake, sender []common.Address, nftId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "UnStake", senderRule, nftIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenUnStake)
				if err := _Token.contract.UnpackLog(event, "UnStake", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnStake is a log parse operation binding the contract event 0x54a9763035584fc4fcad1bc4e0e7a83f93e016f50ae32bd527530a77257393ee.
//
// Solidity: event UnStake(address indexed sender, uint256 indexed nftId, uint256 amount)
func (_Token *TokenFilterer) ParseUnStake(log types.Log) (*TokenUnStake, error) {
	event := new(TokenUnStake)
	if err := _Token.contract.UnpackLog(event, "UnStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenUrgentFunctionUpdateIterator is returned from FilterUrgentFunctionUpdate and is used to iterate over the raw logs and unpacked data for UrgentFunctionUpdate events raised by the Token contract.
type TokenUrgentFunctionUpdateIterator struct {
	Event *TokenUrgentFunctionUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenUrgentFunctionUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenUrgentFunctionUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenUrgentFunctionUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenUrgentFunctionUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenUrgentFunctionUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenUrgentFunctionUpdate represents a UrgentFunctionUpdate event raised by the Token contract.
type TokenUrgentFunctionUpdate struct {
	UrgentFunction [4]byte
	Set            bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUrgentFunctionUpdate is a free log retrieval operation binding the contract event 0x514c1cd5b948b0c2794705911eaf30633a7de72469f47eb8619abea755e9c68c.
//
// Solidity: event UrgentFunctionUpdate(bytes4 urgentFunction, bool set)
func (_Token *TokenFilterer) FilterUrgentFunctionUpdate(opts *bind.FilterOpts) (*TokenUrgentFunctionUpdateIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "UrgentFunctionUpdate")
	if err != nil {
		return nil, err
	}
	return &TokenUrgentFunctionUpdateIterator{contract: _Token.contract, event: "UrgentFunctionUpdate", logs: logs, sub: sub}, nil
}

// WatchUrgentFunctionUpdate is a free log subscription operation binding the contract event 0x514c1cd5b948b0c2794705911eaf30633a7de72469f47eb8619abea755e9c68c.
//
// Solidity: event UrgentFunctionUpdate(bytes4 urgentFunction, bool set)
func (_Token *TokenFilterer) WatchUrgentFunctionUpdate(opts *bind.WatchOpts, sink chan<- *TokenUrgentFunctionUpdate) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "UrgentFunctionUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenUrgentFunctionUpdate)
				if err := _Token.contract.UnpackLog(event, "UrgentFunctionUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUrgentFunctionUpdate is a log parse operation binding the contract event 0x514c1cd5b948b0c2794705911eaf30633a7de72469f47eb8619abea755e9c68c.
//
// Solidity: event UrgentFunctionUpdate(bytes4 urgentFunction, bool set)
func (_Token *TokenFilterer) ParseUrgentFunctionUpdate(log types.Log) (*TokenUrgentFunctionUpdate, error) {
	event := new(TokenUrgentFunctionUpdate)
	if err := _Token.contract.UnpackLog(event, "UrgentFunctionUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Token contract.
type TokenWithdrawIterator struct {
	Event *TokenWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenWithdraw represents a Withdraw event raised by the Token contract.
type TokenWithdraw struct {
	Owner          common.Address
	NftId          *big.Int
	WithdrawAmount *big.Int
	BurnedP        *big.Int
	BurnedItp      *big.Int
	BurnedIfp      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xb7ee00edafdb5db3a1b52a5789a6b562eb48268842d113dbfe860b6f86e5f595.
//
// Solidity: event Withdraw(address indexed owner, uint256 indexed nftId, uint256 withdrawAmount, uint256 burnedP, uint256 burnedItp, uint256 burnedIfp)
func (_Token *TokenFilterer) FilterWithdraw(opts *bind.FilterOpts, owner []common.Address, nftId []*big.Int) (*TokenWithdrawIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Withdraw", ownerRule, nftIdRule)
	if err != nil {
		return nil, err
	}
	return &TokenWithdrawIterator{contract: _Token.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xb7ee00edafdb5db3a1b52a5789a6b562eb48268842d113dbfe860b6f86e5f595.
//
// Solidity: event Withdraw(address indexed owner, uint256 indexed nftId, uint256 withdrawAmount, uint256 burnedP, uint256 burnedItp, uint256 burnedIfp)
func (_Token *TokenFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *TokenWithdraw, owner []common.Address, nftId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Withdraw", ownerRule, nftIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenWithdraw)
				if err := _Token.contract.UnpackLog(event, "Withdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdraw is a log parse operation binding the contract event 0xb7ee00edafdb5db3a1b52a5789a6b562eb48268842d113dbfe860b6f86e5f595.
//
// Solidity: event Withdraw(address indexed owner, uint256 indexed nftId, uint256 withdrawAmount, uint256 burnedP, uint256 burnedItp, uint256 burnedIfp)
func (_Token *TokenFilterer) ParseWithdraw(log types.Log) (*TokenWithdraw, error) {
	event := new(TokenWithdraw)
	if err := _Token.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenWithdrawActualProfitIterator is returned from FilterWithdrawActualProfit and is used to iterate over the raw logs and unpacked data for WithdrawActualProfit events raised by the Token contract.
type TokenWithdrawActualProfitIterator struct {
	Event *TokenWithdrawActualProfit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenWithdrawActualProfitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenWithdrawActualProfit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenWithdrawActualProfit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenWithdrawActualProfitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenWithdrawActualProfitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenWithdrawActualProfit represents a WithdrawActualProfit event raised by the Token contract.
type TokenWithdrawActualProfit struct {
	Sender         common.Address
	ProfitWithdraw *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdrawActualProfit is a free log retrieval operation binding the contract event 0x6fcb4d7b632ae52561ee85b9e54cd5417be51edaeb5534dba41effd91ce1bcb7.
//
// Solidity: event WithdrawActualProfit(address indexed sender, uint256 profitWithdraw)
func (_Token *TokenFilterer) FilterWithdrawActualProfit(opts *bind.FilterOpts, sender []common.Address) (*TokenWithdrawActualProfitIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "WithdrawActualProfit", senderRule)
	if err != nil {
		return nil, err
	}
	return &TokenWithdrawActualProfitIterator{contract: _Token.contract, event: "WithdrawActualProfit", logs: logs, sub: sub}, nil
}

// WatchWithdrawActualProfit is a free log subscription operation binding the contract event 0x6fcb4d7b632ae52561ee85b9e54cd5417be51edaeb5534dba41effd91ce1bcb7.
//
// Solidity: event WithdrawActualProfit(address indexed sender, uint256 profitWithdraw)
func (_Token *TokenFilterer) WatchWithdrawActualProfit(opts *bind.WatchOpts, sink chan<- *TokenWithdrawActualProfit, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "WithdrawActualProfit", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenWithdrawActualProfit)
				if err := _Token.contract.UnpackLog(event, "WithdrawActualProfit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawActualProfit is a log parse operation binding the contract event 0x6fcb4d7b632ae52561ee85b9e54cd5417be51edaeb5534dba41effd91ce1bcb7.
//
// Solidity: event WithdrawActualProfit(address indexed sender, uint256 profitWithdraw)
func (_Token *TokenFilterer) ParseWithdrawActualProfit(log types.Log) (*TokenWithdrawActualProfit, error) {
	event := new(TokenWithdrawActualProfit)
	if err := _Token.contract.UnpackLog(event, "WithdrawActualProfit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenWithdrawForwInterestIterator is returned from FilterWithdrawForwInterest and is used to iterate over the raw logs and unpacked data for WithdrawForwInterest events raised by the Token contract.
type TokenWithdrawForwInterestIterator struct {
	Event *TokenWithdrawForwInterest // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenWithdrawForwInterestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenWithdrawForwInterest)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenWithdrawForwInterest)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenWithdrawForwInterestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenWithdrawForwInterestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenWithdrawForwInterest represents a WithdrawForwInterest event raised by the Token contract.
type TokenWithdrawForwInterest struct {
	Sender    common.Address
	Claimable *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawForwInterest is a free log retrieval operation binding the contract event 0xbb7e1dd12387ea2ea599b52a762e1a3a0405cdf3482b43c44e656517a3eb7e83.
//
// Solidity: event WithdrawForwInterest(address indexed sender, uint256 claimable)
func (_Token *TokenFilterer) FilterWithdrawForwInterest(opts *bind.FilterOpts, sender []common.Address) (*TokenWithdrawForwInterestIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "WithdrawForwInterest", senderRule)
	if err != nil {
		return nil, err
	}
	return &TokenWithdrawForwInterestIterator{contract: _Token.contract, event: "WithdrawForwInterest", logs: logs, sub: sub}, nil
}

// WatchWithdrawForwInterest is a free log subscription operation binding the contract event 0xbb7e1dd12387ea2ea599b52a762e1a3a0405cdf3482b43c44e656517a3eb7e83.
//
// Solidity: event WithdrawForwInterest(address indexed sender, uint256 claimable)
func (_Token *TokenFilterer) WatchWithdrawForwInterest(opts *bind.WatchOpts, sink chan<- *TokenWithdrawForwInterest, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "WithdrawForwInterest", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenWithdrawForwInterest)
				if err := _Token.contract.UnpackLog(event, "WithdrawForwInterest", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawForwInterest is a log parse operation binding the contract event 0xbb7e1dd12387ea2ea599b52a762e1a3a0405cdf3482b43c44e656517a3eb7e83.
//
// Solidity: event WithdrawForwInterest(address indexed sender, uint256 claimable)
func (_Token *TokenFilterer) ParseWithdrawForwInterest(log types.Log) (*TokenWithdrawForwInterest, error) {
	event := new(TokenWithdrawForwInterest)
	if err := _Token.contract.UnpackLog(event, "WithdrawForwInterest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenWithdrawTokenInterestIterator is returned from FilterWithdrawTokenInterest and is used to iterate over the raw logs and unpacked data for WithdrawTokenInterest events raised by the Token contract.
type TokenWithdrawTokenInterestIterator struct {
	Event *TokenWithdrawTokenInterest // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenWithdrawTokenInterestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenWithdrawTokenInterest)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenWithdrawTokenInterest)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenWithdrawTokenInterestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenWithdrawTokenInterestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenWithdrawTokenInterest represents a WithdrawTokenInterest event raised by the Token contract.
type TokenWithdrawTokenInterest struct {
	Sender    common.Address
	Claimable *big.Int
	Bonus     *big.Int
	Profit    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawTokenInterest is a free log retrieval operation binding the contract event 0x5204ccba4ecb08e13cda3f108f4efe1b98ff8443251ff98749cdd2dabb52f37a.
//
// Solidity: event WithdrawTokenInterest(address indexed sender, uint256 claimable, uint256 bonus, uint256 profit)
func (_Token *TokenFilterer) FilterWithdrawTokenInterest(opts *bind.FilterOpts, sender []common.Address) (*TokenWithdrawTokenInterestIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "WithdrawTokenInterest", senderRule)
	if err != nil {
		return nil, err
	}
	return &TokenWithdrawTokenInterestIterator{contract: _Token.contract, event: "WithdrawTokenInterest", logs: logs, sub: sub}, nil
}

// WatchWithdrawTokenInterest is a free log subscription operation binding the contract event 0x5204ccba4ecb08e13cda3f108f4efe1b98ff8443251ff98749cdd2dabb52f37a.
//
// Solidity: event WithdrawTokenInterest(address indexed sender, uint256 claimable, uint256 bonus, uint256 profit)
func (_Token *TokenFilterer) WatchWithdrawTokenInterest(opts *bind.WatchOpts, sink chan<- *TokenWithdrawTokenInterest, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "WithdrawTokenInterest", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenWithdrawTokenInterest)
				if err := _Token.contract.UnpackLog(event, "WithdrawTokenInterest", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawTokenInterest is a log parse operation binding the contract event 0x5204ccba4ecb08e13cda3f108f4efe1b98ff8443251ff98749cdd2dabb52f37a.
//
// Solidity: event WithdrawTokenInterest(address indexed sender, uint256 claimable, uint256 bonus, uint256 profit)
func (_Token *TokenFilterer) ParseWithdrawTokenInterest(log types.Log) (*TokenWithdrawTokenInterest, error) {
	event := new(TokenWithdrawTokenInterest)
	if err := _Token.contract.UnpackLog(event, "WithdrawTokenInterest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
