package server

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	v1 "github.com/choraio/mods/geonode/types/v1"
)

// Node implements the Query/Node method.
func (s Server) Node(ctx context.Context, req *v1.QueryNodeRequest) (*v1.QueryNodeResponse, error) {

	// get node from node table
	node, err := s.ss.NodeTable().Get(ctx, req.Id)
	if err != nil {
		if ormerrors.NotFound.Is(err) {
			return nil, sdkerrors.ErrNotFound.Wrapf("node with id %d", req.Id)
		}
		return nil, err // internal error
	}

	// get account from account bytes
	curator := sdk.AccAddress(node.Curator)

	// return query response
	return &v1.QueryNodeResponse{
		Id:       node.Id,
		Curator:  curator.String(),
		Metadata: node.Metadata,
	}, nil
}
