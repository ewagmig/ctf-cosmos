package adv

import (
	"github.com/ewagmig/adv/x/adv/keeper"
	"github.com/ewagmig/adv/x/adv/types"
)

const (
	ModuleName   = types.ModuleName
	StoreKey     = types.StoreKey
	RouterKey    = types.RouterKey
	QuerierRoute = types.QuerierRoute
)

var (
	// functions aliases
	NewKeeper           = keeper.NewKeeper
	NewQuerier          = keeper.NewQuerier
	RegisterCodec       = types.RegisterCodec
	NewGenesisState     = types.NewGenesisState
	DefaultGenesisState = types.DefaultGenesisState

	// variable aliases
	ModuleCdc = types.ModuleCdc
)

type (
	Keeper       = keeper.Keeper
	GenesisState = types.GenesisState
)
