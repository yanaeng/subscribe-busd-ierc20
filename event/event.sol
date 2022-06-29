pragma solidity ^0.8.14;

contract CoreBorrowingEvent {
// core borrowing event
    event Borrow(
        address indexed owner,
        uint256 indexed nftId,
        uint256 indexed loanId,
        address borrowTokenAddress,
        address collateralTokenAddress,
        uint256 borrowAmount,
        uint256 collateralAmount,
        uint256 owedPerDay,
        uint256 minInterest,
        uint8 newLoan,
        uint64 rolloverTimestamp
    );

    event CloseLoan(
        address indexed owner,
        uint256 indexed nftId,
        uint256 indexed loanId,
        uint256 borrowPaid,
        uint256 interestPaid,
        uint256 collateralAmountwithdraw
    );
    event Repay(
        address indexed owner,
        uint256 indexed nftId,
        uint256 indexed loanId,
        bool closeLoan,
        uint256 borrowPaid,
        uint256 interestPaid,
        uint256 collateralAmountwithdraw
    );
    event AdjustCollateral(
        address indexed owner,
        uint256 indexed nftId,
        uint256 indexed loanId,
        bool isAdd,
        uint256 collateralAdjustAmount
    );
    event Rollover(
        address indexed owner,
        uint256 indexed nftId,
        uint256 indexed loanId,
        address bountyHunter,
        uint256 delayInterest,
        uint256 bountyReward,
        address bountyRewardTokenAddress,
        uint256 newInterestOwedPerDay
    );

    event Liquidate(
        address indexed owner,
        uint256 indexed nftId,
        uint256 indexed loanId,
        address liquidator,
        uint256 swapPrice,
        uint256 tokenAmountFromSwap,
        uint256 bountyReward,
        address bountyRewardTokenAddress,
        uint256 tokenSentBackToUser
    );
    // core event 

    event SettleForwInterest(
        address indexed coreAddress,
        address indexed interestVaultAddress,
        address forwDistributionAddress,
        address forwTokenAddress,
        uint256 amount
    );

    // core setting event
    event SetMembershipAddress(address indexed sender, address oldValue, address newValue); //same event as in pool setting
    event SetPriceFeedAddress(address indexed sender, address oldValue, address newValue);
    event SetRouterAddress(address indexed sender, address oldValue, address newValue);
    event SetCoreBorrowingAddress(address indexed sender, address oldValue, address newValue);
    event SetFeeController(address indexed sender, address oldValue, address newValue);
    event SetForwDistributorAddress(address indexed sender, address oldValue, address newValue);

    event SetWETHHandler(address indexed sender, address oldValue, address newValue); //same event as in pool setting

    event SetLoanDuration(address indexed sender, uint256 oldValue, uint256 newValue);
    event SetAdvancedInterestDuration(address indexed sender, uint256 oldValue, uint256 newValue);
    event SetFeeSpread(address indexed sender, uint256 oldValue, uint256 newValue);

    event RegisterNewPool(address indexed sender, address poolAddress);
    event SetupLoanConfig(
        address indexed sender,
        address indexed borrowTokenAddress,
        address indexed collateralTokenAddress,
        uint256 oldSafeLTV,
        uint256 oldMaxLTV,
        uint256 oldLiquidationLTV,
        uint256 oldBountyFeeRate,
        uint256 newSafeLTV,
        uint256 newLMaxLTV,
        uint256 newLiquidationLTV,
        uint256 newBountyFeeRate
    );
    event SetFowPerBlock(address indexed sender, uint256 amount, uint256 targetBlock);
    event ApprovedForRouter(address indexed sender, address asset, address router);

    // timelock
    event CallScheduled(
        bytes32 indexed id,
        uint256 indexed index,
        address target,
        uint256 value,
        bytes data,
        bytes32 predecessor,
        uint256 delay
    );
    event CallExecuted(
        bytes32 indexed id,
        uint256 indexed index,
        address target,
        uint256 value,
        bytes data
    );
    event Cancelled(bytes32 indexed id);
    event ProxyAdminChange(address oldProxyAdmin, address newProxyAdmin);
    event UrgentFunctionUpdate(bytes4 urgentFunction, bool set);
    event MinDelayChange(uint256 oldDuration, uint256 newDuration);
    event MinUrgentDelayChange(uint256 oldDuration, uint256 newDuration);
    event MinProxyAdminDelayChange(uint256 oldDuration, uint256 newDuration);
    event AvailablePeriodChange(uint256 oldDuration, uint256 newDuration);

    // interest vault event
    event SetTokenAddress(address indexed sender, address oldValue, address newValue);
    event SetForwAddress(address indexed sender, address oldValue, address newValue);
    event SetProtocolAddress(address indexed sender, address oldValue, address newValue);

    event OwnerApprove(
        address indexed sender,
        address tokenAddress,
        address forwAddress,
        uint256 amount
    );

    event SettleInterest(
        address indexed sender,
        uint256 claimableTokenInterest,
        uint256 heldTokenInterest,
        uint256 claimableForwInterest
    );

    event WithdrawTokenInterest(
        address indexed sender,
        uint256 claimable,
        uint256 bonus,
        uint256 profit
    );

    event WithdrawForwInterest(address indexed sender, uint256 claimable);

    event WithdrawActualProfit(address indexed sender, uint256 profitWithdraw);
    // pool lending event
    event Deposit(
        address indexed owner,
        uint256 indexed nftId,
        uint256 depositAmount,
        uint256 mintedP,
        uint256 mintedItp,
        uint256 mintedIfp
    );
    event Withdraw(
        address indexed owner,
        uint256 indexed nftId,
        uint256 withdrawAmount,
        uint256 burnedP,
        uint256 burnedItp,
        uint256 burnedIfp
    );
    event ClaimTokenInterest(
        address indexed owner,
        uint256 indexed nftId,
        uint256 interestTokenClaimed,
        uint256 interestTokenBonus,
        uint256 burnedItp
    );
    event ClaimForwInterest(
        address indexed owner,
        uint256 indexed nftId,
        uint256 interestForwClaimed,
        uint256 interestForwBonus,
        uint256 burnedIfp
    );
    event ActivateRank(address indexed owner, uint256 indexed nftId, uint8 oldRank, uint8 newRank);

    // pool setting event
    event SetBorrowInterestParams(
        address indexed sender,
        uint256[] rates,
        uint256[] utils,
        uint256 targetSupply
    );

    event SetLoanConfig(
        address indexed sender,
        address collateralTokenAddress,
        uint256 safeLTV,
        uint256 maxLTV,
        uint256 liqLTV,
        uint256 bountyFeeRate
    );

    event SetPoolLendingAddress(address indexed sender, address oldValue, address newValue);

    event SetPoolBorrowingAddress(address indexed sender, address oldValue, address newValue);


    event Initialize(
        address indexed manager,
        address indexed coreAddress,
        address interestVaultAddress,
        address membershipAddress
    );

    // pool token
    event MintPToken(address indexed minter, uint256 indexed nftId, uint256 amount);
    event MintItpToken(
        address indexed minter,
        uint256 indexed nftId,
        uint256 amount,
        uint256 price
    );
    event MintIfpToken(
        address indexed minter,
        uint256 indexed nftId,
        uint256 amount,
        uint256 price
    );

    event BurnPToken(address indexed burner, uint256 indexed nftId, uint256 amount);
    event BurnItpToken(
        address indexed burner,
        uint256 indexed nftId,
        uint256 amount,
        uint256 price
    );
    event BurnIfpToken(
        address indexed burner,
        uint256 indexed nftId,
        uint256 amount,
        uint256 price
    );

    // stake pool
    event Stake(address indexed sender, uint256 indexed nftId, uint256 amount);
    event UnStake(address indexed sender, uint256 indexed nftId, uint256 amount);
    event DeprecateStakeInfo(address indexed sender, uint256 indexed nftId);
    event SetPoolStartTimestamp(address indexed sender, uint64 indexed timestamp);

    // utils
    event FaucetTransfer(
        address indexed to,
        address tokenAddress,
        uint256 value,
        uint256 timestamp
    );
    event GlobalPricingPaused(address indexed sender, bool isPaused);
    event SetPriceFeed(address indexed sender, address[] tokens, address[] feeds);
    event SetDecimals(address indexed sender, address[] tokens);
}