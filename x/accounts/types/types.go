package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (m *MsgDeploy) ValidateBasic() error {
	if m.InitMessage == nil {
		return fmt.Errorf("invalid init msg")
	}

	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return err
	}
	// TODO
	return nil
}

func (m *MsgDeploy) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Sender)}
}

func (m *MsgExecute) ValidateBasic() error {
	if m.Message == nil {
		return fmt.Errorf("invalid msg")
	}

	if _, err := sdk.AccAddressFromBech32(m.Address); err != nil {
		return err
	}

	// TODO
	return nil
}

func (m *MsgExecute) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Address)}
}
