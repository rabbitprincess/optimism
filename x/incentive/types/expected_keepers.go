package types

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	epochingtypes "github.com/ethereum-optimism/optimism/x/epoching/types"
)

type AccountKeeper interface {
	GetAccount(ctx context.Context, addr sdk.AccAddress) sdk.AccountI
	GetModuleAccount(ctx context.Context, name string) sdk.ModuleAccountI
}

type BankKeeper interface {
	SpendableCoins(ctx context.Context, addr sdk.AccAddress) sdk.Coins
	GetAllBalances(ctx context.Context, addr sdk.AccAddress) sdk.Coins
	SendCoinsFromModuleToModule(ctx context.Context, senderModule string, recipientModule string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
}

type EpochingKeeper interface {
	GetEpoch(ctx context.Context) *epochingtypes.Epoch
}
