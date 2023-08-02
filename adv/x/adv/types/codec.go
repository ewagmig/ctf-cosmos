package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreateAdvertisement{}, "adv/CreateAdvertisement", nil)
	cdc.RegisterConcrete(MsgDeleteAdvertisement{}, "adv/DeleteAdvertisement", nil)
	cdc.RegisterConcrete(MsgDeposit{}, "adv/deposit", nil)
	cdc.RegisterConcrete(MsgWithdraw{}, "adv/withdraw", nil)
	cdc.RegisterConcrete(MsgCaptureTheFlag{}, "adv/CaptureTheFlag", nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
