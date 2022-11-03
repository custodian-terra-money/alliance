package alliance

import (
	"alliance/x/alliance/keeper"
	"alliance/x/alliance/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

func NewAllianceProposalHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.CreateAllianceProposal:
			return k.CreateAlliance(ctx, c)
		case *types.UpdateAllianceProposal:
			return k.UpdateAlliance(ctx, c)
		case *types.DeleteAllianceProposal:
			return k.DeleteAlliance(ctx, c)

		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized alliance proposal content type: %T", c)
		}
	}
}
