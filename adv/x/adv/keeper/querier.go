package keeper

import (
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/ewagmig/adv/x/adv/types"
)

// NewQuerier creates a new querier for adv clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryListAdvertisement:
			return listAdvertisement(ctx, k)
		case types.QueryGetAdvertisement:
			return getAdvertisement(ctx, path[1:], k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown adv query endpoint")
		}
	}
}
