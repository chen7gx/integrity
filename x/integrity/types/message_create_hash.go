package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateHash{}

func NewMsgCreateHash(creator string, details string, hash string) *MsgCreateHash {
	return &MsgCreateHash{
		Creator: creator,
		Details: details,
		Hash:    hash,
	}
}

func (msg *MsgCreateHash) Route() string {
	return RouterKey
}

func (msg *MsgCreateHash) Type() string {
	return "CreateHash"
}

func (msg *MsgCreateHash) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateHash) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateHash) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
