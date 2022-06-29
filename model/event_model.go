package model

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

//struct borrow
type LogBorrow struct {
	Owner                  common.Address
	NftId                  *big.Int
	LoanId                 *big.Int
	BorrowTokenAddress     common.Address
	CollateralTokenAddress common.Address
	BorrowAmount           *big.Int
	CollateralAmount       *big.Int
	OwedPerDay             *big.Int
	MinInterest            *big.Int
	NewLoan                *big.Int
	RolloverTimestamp      *big.Int
	CreatedTime            common.Address
	ContractAddress        common.Address
	TxHash                 common.Address
}

type LogCloseLoan struct {
	Owner                    common.Address
	NftId                    *big.Int
	LoanId                   *big.Int
	BorrowPaid               *big.Int
	InterestPaid             *big.Int
	CollateralAmountwithdraw *big.Int
	CreatedTime              common.Address
	ContractAddress          common.Address
	TxHash                   common.Address
}

type LogRepay struct {
	Owner                    common.Address
	NftId                    *big.Int
	LoanId                   *big.Int
	CloseLoan                bool
	BorrowPaid               *big.Int
	InterestPaid             *big.Int
	CollateralAmountwithdraw *big.Int
	CreatedTime              common.Address
	ContractAddress          common.Address
	TxHash                   common.Address
}

type LogAdjustCollateral struct {
	Owner                  common.Address
	NftId                  *big.Int
	LoanId                 *big.Int
	IsAdd                  bool
	CollateralAdjustAmount *big.Int
	CreatedTime            common.Address
	ContractAddress        common.Address
	TxHash                 common.Address
}

type LogRollover struct {
	Owner                    common.Address
	NftId                    *big.Int
	LoanId                   *big.Int
	BountyHunter             common.Address
	DelayInterest            *big.Int
	BountyReward             *big.Int
	BountyRewardTokenAddress common.Address
	NewInterestOwedPerDay    common.Address
	CreatedTime              common.Address
	ContractAddress          common.Address
	TxHash                   common.Address
}

type LogLiquidate struct {
	Owner                    common.Address
	NftId                    *big.Int
	LoanId                   *big.Int
	Liquidator               common.Address
	SwapPrice                *big.Int
	TokenAmountFromSwap      *big.Int
	BountyReward             *big.Int
	BountyRewardTokenAddress common.Address
	CreatedTime              common.Address
	ContractAddress          common.Address
	TxHash                   common.Address
}

type LogSettleForwInterest struct {
	CoreAddress             common.Address
	InterestVaultAddress    common.Address
	ForwDistributionAddress common.Address
	ForwTokenAddress        common.Address
	Amount                  *big.Int
	CreatedTime             common.Address
	ContractAddress         common.Address
	TxHash                  common.Address
}

type LogSetMembershipAddress struct {
	Sender          common.Address
	OldValue        common.Address
	NewValue        common.Address
	CreatedTime     common.Address
	ContractAddress common.Address
	TxHash          common.Address
}

type LogSetPriceFeedAddress struct {
	Sender          common.Address
	OldValue        common.Address
	NewValue        common.Address
	CreatedTime     common.Address
	ContractAddress common.Address
	TxHash          common.Address
}

type LogSetRouterAddress struct {
	Sender          common.Address
	OldValue        common.Address
	NewValue        common.Address
	CreatedTime     common.Address
	ContractAddress common.Address
	TxHash          common.Address
}

type LogSetCoreBorrowingAddress struct {
	Sender          common.Address
	OldValue        common.Address
	NewValue        common.Address
	CreatedTime     common.Address
	ContractAddress common.Address
	TxHash          common.Address
}

type LogSetFeeController struct {
	Sender          common.Address
	OldValue        common.Address
	NewValue        common.Address
	CreatedTime     common.Address
	ContractAddress common.Address
	TxHash          common.Address
}

type LogSetForwDistributorAddress struct {
	Sender          common.Address
	OldValue        common.Address
	NewValue        common.Address
	CreatedTime     common.Address
	ContractAddress common.Address
	TxHash          common.Address
}

type LogSetWETHHandler struct {
	Sender          common.Address
	OldValue        common.Address
	NewValue        common.Address
	CreatedTime     common.Address
	ContractAddress common.Address
	TxHash          common.Address
}

type LogSetLoanDuration struct {
	Sender          common.Address
	OldValue        *big.Int
	NewValue        *big.Int
	CreatedTime     common.Address
	ContractAddress common.Address
	TxHash          common.Address
}

type LogSetAdvancedInterestDuration struct {
	Sender          common.Address
	OldValue        *big.Int
	NewValue        *big.Int
	CreatedTime     common.Address
	ContractAddress common.Address
	TxHash          common.Address
}

type LogSetFeeSpread struct {
	Sender          common.Address
	OldValue        *big.Int
	NewValue        *big.Int
	CreatedTime     common.Address
	ContractAddress common.Address
	TxHash          common.Address
}

type LogRegisterNewPool struct {
	Sender          common.Address
	PoolAddress     common.Address
	CreatedTime     common.Address
	ContractAddress common.Address
	TxHash          common.Address
}

type LogSetupLoanConfig struct {
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
	CreatedTime            common.Address
	ContractAddress        common.Address
	TxHash                 common.Address
}

type LogSetFowPerBlock struct {
	Sender          common.Address
	Amount          *big.Int
	TargetBlock     *big.Int
	CreatedTime     common.Address
	ContractAddress common.Address
	TxHash          common.Address
}

type LogApprovedForRouter struct {
	Sender          common.Address
	Asset           common.Address
	Router          common.Address
	CreatedTime     common.Address
	ContractAddress common.Address
	TxHash          common.Address
}

type LogCallScheduled struct {
	Timelock_id     common.Address
	Index           *big.Int
	Target          common.Address
	Value           *big.Int
	Data            common.Address
	Predecessor     common.Address
	Delay           common.Address
	CreatedTime     common.Address
	ContractAddress common.Address
	TxHash          common.Address
}

type LogCallExecuted struct {
	Timelock_id     common.Address
	Index           *big.Int
	Target          common.Address
	Value           *big.Int
	CreatedTime     common.Address
	ContractAddress common.Address
	TxHash          common.Address
}

type LogCancelled struct {
	Timelock_id     common.Address
	CreatedTime     common.Address
	ContractAddress common.Address
	TxHash          common.Address
}

type LogProxyAdminChange struct {
	OldProxyAdmin   common.Address
	NewProxyAdmin   common.Address
	CreatedTime     common.Address
	ContractAddress common.Address
	TxHash          common.Address
}

type LogUrgentFunctionUpdate struct {
	UrgentFunction  common.Address
	Set             bool
	CreatedTime     common.Address
	ContractAddress common.Address
	TxHash          common.Address
}

type LogMinDelayChange struct {
	OldDuration     *big.Int
	NewDuration     *big.Int
	CreatedTime     common.Address
	ContractAddress common.Address
	TxHash          common.Address
}

type LogMinUrgentDelayChange struct {
	OldDuration     *big.Int
	NewDuration     *big.Int
	CreatedTime     common.Address
	ContractAddress common.Address
	TxHash          common.Address
}

type LogMinProxyAdminDelayChange struct {
	OldDuration     *big.Int
	NewDuration     *big.Int
	CreatedTime     common.Address
	ContractAddress common.Address
	TxHash          common.Address
}

type LogAvailablePeriodChange struct {
	OldDuration     *big.Int
	NewDuration     *big.Int
	CreatedTime     common.Address
	ContractAddress common.Address
	TxHash          common.Address
}

type LogSetTokenAddress struct {
	Sender          common.Address
	OldValue        common.Address
	NewValue        common.Address
	CreatedTime     common.Address
	ContractAddress common.Address
	TxHash          common.Address
}

type LogSetForwAddress struct {
	Sender          common.Address
	OldValue        common.Address
	NewValue        common.Address
	CreatedTime     common.Address
	ContractAddress common.Address
	TxHash          common.Address
}

type LogSetProtocolAddress struct {
	Sender          common.Address
	OldValue        common.Address
	NewValue        common.Address
	CreatedTime     common.Address
	ContractAddress common.Address
	TxHash          common.Address
}

type LogOwnerApprove struct {
	Sender          common.Address
	TokenAddress    common.Address
	ForwAddress     common.Address
	Amount          *big.Int
	CreatedTime     common.Address
	ContractAddress common.Address
	TxHash          common.Address
}

type LogSettleInterest struct {
	Sender                 common.Address
	ClaimableTokenInterest *big.Int
	HeldTokenInterest      *big.Int
	ClaimableForwInterest  *big.Int
	CreatedTime            common.Address
	ContractAddress        common.Address
	TxHash                 common.Address
}

type LogWithdrawTokenInterest struct {
	Sender          common.Address
	Claimable       *big.Int
	Bonus           *big.Int
	Profit          *big.Int
	CreatedTime     common.Address
	ContractAddress common.Address
	TxHash          common.Address
}

type LogWithdrawForwInterest struct {
	Sender          common.Address
	Claimable       *big.Int
	CreatedTime     common.Address
	ContractAddress common.Address
	TxHash          common.Address
}

type LogWithdrawActualProfit struct {
	Sender          common.Address
	ProfitWithdraw  *big.Int
	CreatedTime     common.Address
	ContractAddress common.Address
	TxHash          common.Address
}

type LogDeposit struct {
	Owner           common.Address
	NftId           *big.Int
	DepositAmount   *big.Int
	MintedP         *big.Int
	MintedItp       *big.Int
	MintedIfp       *big.Int
	CreatedTime     common.Address
	ContractAddress common.Address
	TxHash          common.Address
}

type LogWithdraw struct {
	Owner                common.Address
	NftId                *big.Int
	InterestTokenClaimed *big.Int
	InterestTokenBonus   *big.Int
	BurnedItp            *big.Int
	CreatedTime          common.Address
	ContractAddress      common.Address
	TxHash               common.Address
}

type LogClaimTokenInterest struct {
	Owner                common.Address
	NftId                *big.Int
	InterestTokenClaimed *big.Int
	InterestTokenBonus   *big.Int
	BurnedItp            *big.Int
	CreatedTime          common.Address
	ContractAddress      common.Address
	TxHash               common.Address
}

type LogClaimForwInterest struct {
	Owner               common.Address
	NftId               *big.Int
	InterestForwClaimed *big.Int
	InterestForwBonus   *big.Int
	BurnedIfp           *big.Int
	CreatedTime         common.Address
	ContractAddress     common.Address
	TxHash              common.Address
}

type LogActivateRank struct {
	Owner           common.Address
	NftId           *big.Int
	OldRank         *big.Int
	NewRank         *big.Int
	BurnedIfp       *big.Int
	CreatedTime     common.Address
	ContractAddress common.Address
	TxHash          common.Address
}

type LogSetBorrowInterestParams struct {
	Sender          common.Address
	Rates           []*big.Int
	Utils           []*big.Int
	TargetSupply    common.Address
	CreatedTime     common.Address
	ContractAddress common.Address
	TxHash          common.Address
}

type LogSetLoanConfig struct {
	Sender                 common.Address
	CollateralTokenAddress common.Address
	SafeLTV                *big.Int
	MaxLTV                 *big.Int
	LiqLTV                 *big.Int
	BountyFeeRate          *big.Int
	CreatedTime            common.Address
	ContractAddress        common.Address
	TxHash                 common.Address
}

type LogSetPoolLendingAddress struct {
	Sender          common.Address
	OldValue        common.Address
	NewValue        common.Address
	CreatedTime     common.Address
	ContractAddress common.Address
	TxHash          common.Address
}

type LogSetPoolBorrowingAddress struct {
	Sender          common.Address
	OldValue        common.Address
	NewValue        common.Address
	CreatedTime     common.Address
	ContractAddress common.Address
	TxHash          common.Address
}

type LogInitialize struct {
	Manager              common.Address
	CoreAddress          common.Address
	InterestVaultAddress common.Address
	MembershipAddress    common.Address
	CreatedTime          common.Address
	ContractAddress      common.Address
	TxHash               common.Address
}

type LogMintPToken struct {
	Minter          common.Address
	NftId           *big.Int
	Amount          *big.Int
	CreatedTime     common.Address
	ContractAddress common.Address
	TxHash          common.Address
}

type LogMintItpToken struct {
	Minter          common.Address
	NftId           *big.Int
	Amount          *big.Int
	Price           *big.Int
	CreatedTime     common.Address
	ContractAddress common.Address
	TxHash          common.Address
}

type LogMintIfpToken struct {
	Minter          common.Address
	NftId           *big.Int
	Amount          *big.Int
	Price           *big.Int
	CreatedTime     common.Address
	ContractAddress common.Address
	TxHash          common.Address
}

type LogBurnPToken struct {
	Burner          common.Address
	NftId           *big.Int
	Amount          *big.Int
	CreatedTime     common.Address
	ContractAddress common.Address
	TxHash          common.Address
}

type LogBurnItpToken struct {
	Burner          common.Address
	NftId           *big.Int
	Amount          *big.Int
	Price           *big.Int
	CreatedTime     common.Address
	ContractAddress common.Address
	TxHash          common.Address
}

type LogBurnIfpToken struct {
	Burner          common.Address
	NftId           *big.Int
	Amount          *big.Int
	Price           *big.Int
	CreatedTime     common.Address
	ContractAddress common.Address
	TxHash          common.Address
}

type LogStake struct {
	Sender          common.Address
	NftId           *big.Int
	Amount          *big.Int
	CreatedTime     common.Address
	ContractAddress common.Address
	TxHash          common.Address
}

type LogUnStake struct {
	Sender          common.Address
	NftId           *big.Int
	Amount          *big.Int
	CreatedTime     common.Address
	ContractAddress common.Address
	TxHash          common.Address
}

type LogDeprecateStakeInfo struct {
	Sender          common.Address
	NftId           *big.Int
	CreatedTime     common.Address
	ContractAddress common.Address
	TxHash          common.Address
}

type LogSetPoolStartTimestamp struct {
	Sender          common.Address
	CreatedTime     common.Address
	ContractAddress common.Address
	TxHash          common.Address
}

type LogFaucetTransfer struct {
	To              common.Address
	TokenAddress    common.Address
	Value           *big.Int
	CreatedTime     common.Address
	ContractAddress common.Address
	TxHash          common.Address
}

type LogGlobalPricingPaused struct {
	Sender          common.Address
	IsPaused        bool
	CreatedTime     common.Address
	ContractAddress common.Address
	TxHash          common.Address
}

type LogSetPriceFeed struct {
	Sender          common.Address
	Tokens          []common.Address
	Feeds           []common.Address
	CreatedTime     common.Address
	ContractAddress common.Address
	TxHash          common.Address
}

type LogSetDecimals struct {
	Sender          common.Address
	Tokens          []common.Address
	CreatedTime     common.Address
	ContractAddress common.Address
	TxHash          common.Address
}
