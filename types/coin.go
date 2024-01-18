package types

import (
	"math/big"

	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// AttoGurux defines the default coin denomination used in Gurux in:
	//
	// - Staking parameters: denomination used as stake in the dPoS chain
	// - Mint parameters: denomination minted due to fee distribution rewards
	// - Governance parameters: denomination used for spam prevention in proposal deposits
	// - Crisis parameters: constant fee denomination used for spam prevention to check broken invariant
	// - EVM parameters: denomination used for running EVM state transitions in Gurux.
	AttoGurux string = "agurux"

	// BaseDenomUnit defines the base denomination unit for Gurux.
	// 1 gurux = 1x10^{BaseDenomUnit} agurux
	BaseDenomUnit = 18

	// DefaultGasPrice is default gas price for evm transactions
	DefaultGasPrice = 20
)

// PowerReduction defines the default power reduction value for staking
var PowerReduction = sdkmath.NewIntFromBigInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(BaseDenomUnit), nil))

// NewGuruxCoin is a utility function that returns an "agurux" coin with the given sdkmath.Int amount.
// The function will panic if the provided amount is negative.
func NewGuruxCoin(amount sdkmath.Int) sdk.Coin {
	return sdk.NewCoin(AttoGurux, amount)
}

// NewGuruxDecCoin is a utility function that returns an "agurux" decimal coin with the given sdkmath.Int amount.
// The function will panic if the provided amount is negative.
func NewGuruxDecCoin(amount sdkmath.Int) sdk.DecCoin {
	return sdk.NewDecCoin(AttoGurux, amount)
}

// NewGuruxCoinInt64 is a utility function that returns an "agurux" coin with the given int64 amount.
// The function will panic if the provided amount is negative.
func NewGuruxCoinInt64(amount int64) sdk.Coin {
	return sdk.NewInt64Coin(AttoGurux, amount)
}
