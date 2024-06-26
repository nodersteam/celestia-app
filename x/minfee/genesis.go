package minfee

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	params "github.com/cosmos/cosmos-sdk/x/params/keeper"
)

// DefaultGenesis returns the default genesis state.
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		GlobalMinGasPrice: DefaultGlobalMinGasPrice,
	}
}

// ValidateGenesis performs basic validation of genesis data returning an error for any failed validation criteria.
func ValidateGenesis(genesis *GenesisState) error {
	if genesis.GlobalMinGasPrice.IsNegative() || genesis.GlobalMinGasPrice.IsZero() {
		return fmt.Errorf("global min gas price cannot be negative: %g", genesis.GlobalMinGasPrice)
	}

	return nil
}

// ExportGenesis returns the minfee module's exported genesis.
func ExportGenesis(ctx sdk.Context, k params.Keeper) *GenesisState {
	subspace, exists := k.GetSubspace(ModuleName)
	if !exists {
		panic("minfee subspace not set")
	}
	subspace = RegisterMinFeeParamTable(subspace)

	var globalMinGasPrice sdk.Dec
	subspace.Get(ctx, KeyGlobalMinGasPrice, &globalMinGasPrice)

	return &GenesisState{GlobalMinGasPrice: globalMinGasPrice}
}
