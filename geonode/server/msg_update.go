package server

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	v1 "github.com/choraio/mods/geonode/types/v1"
)

// Update implements the Msg/Update method.
func (s Server) Update(ctx context.Context, req *v1.MsgUpdate) (*v1.MsgUpdateResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// get curator account from bech32
	curator, err := sdk.AccAddressFromBech32(req.Curator)
	if err != nil {
		return nil, err // internal error
	}

	// get node from node table
	node, err := s.ss.NodeTable().Get(ctx, req.Id)
	if err != nil {
		if ormerrors.NotFound.Is(err) {
			return nil, sdkerrors.ErrNotFound.Wrapf(
				"node with id %d: %s", req.Id, err,
			)
		}
		return nil, err // internal error
	}

	// convert node curator to account
	nodeCurator := sdk.AccAddress(node.Curator)

	// verify curator is node curator
	if !nodeCurator.Equals(curator) {
		return nil, sdkerrors.ErrUnauthorized.Wrapf(
			"curator %s: node curator %s", curator, nodeCurator.String(),
		)
	}

	// set new node metadata
	node.Metadata = req.NewMetadata

	// update node in node table
	err = s.ss.NodeTable().Update(ctx, node)
	if err != nil {
		return nil, err // internal error
	}

	// emit update event
	if err = sdkCtx.EventManager().EmitTypedEvent(&v1.EventUpdate{
		Id: node.Id,
	}); err != nil {
		return nil, err // internal error
	}

	// return update response
	return &v1.MsgUpdateResponse{
		Id: node.Id,
	}, nil
}