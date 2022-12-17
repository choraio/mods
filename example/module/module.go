package module

import (
	"context"
	"encoding/json"

	"github.com/choraio/mods/example"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	abci "github.com/tendermint/tendermint/abci/types"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/types/ormjson"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/choraio/mods/example/client"
	"github.com/choraio/mods/example/server"
	"github.com/choraio/mods/example/types/v1"
)

var (
	_ module.AppModule = &Module{}
)

const (
	// ConsensusVersion is the module consensus version.
	ConsensusVersion = 1
)

// Module implements the AppModule interface.
type Module struct {
	key storetypes.StoreKey
	srv server.Server
}

// NewModule returns a new example module.
func NewModule(key storetypes.StoreKey, server server.Server) Module {
	return Module{
		key: key,
		srv: server,
	}
}

// ConsensusVersion implements AppModule/ConsensusVersion.
func (m Module) ConsensusVersion() uint64 {
	return ConsensusVersion
}

// Name implements AppModule/Name.
func (m Module) Name() string {
	return example.ModuleName
}

// Route implements AppModule/Route.
func (m Module) Route() sdk.Route {
	return sdk.Route{}
}

// QuerierRoute implements AppModule/QuerierRoute.
func (m Module) QuerierRoute() string {
	return example.ModuleName
}

// RegisterInvariants implements AppModule/RegisterInvariants.
func (m Module) RegisterInvariants(registry sdk.InvariantRegistry) {
	m.srv.RegisterInvariants(registry)
}

// RegisterInterfaces implements AppModule/RegisterTypes.
func (m Module) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	v1.RegisterTypes(registry)
}

// RegisterServices implements AppModule/RegisterServices.
func (m Module) RegisterServices(cfg module.Configurator) {
	v1.RegisterMsgServer(cfg.MsgServer(), m.srv)
	v1.RegisterQueryServer(cfg.QueryServer(), m.srv)
}

// RegisterGRPCGatewayRoutes implements AppModule/RegisterGRPCGatewayRoutes.
func (m Module) RegisterGRPCGatewayRoutes(clientCtx sdkclient.Context, mux *runtime.ServeMux) {
	err := v1.RegisterQueryHandlerClient(context.Background(), mux, v1.NewQueryClient(clientCtx))
	if err != nil {
		panic(err)
	}
}

// RegisterLegacyAminoCodec implements AppModule/RegisterLegacyAminoCodec.
func (m Module) RegisterLegacyAminoCodec(amino *codec.LegacyAmino) {
	v1.RegisterLegacyAminoCodec(amino)
}

// InitGenesis implements AppModule/InitGenesis.
func (m Module) InitGenesis(ctx sdk.Context, jsonCodec codec.JSONCodec, message json.RawMessage) []abci.ValidatorUpdate {
	update, err := m.srv.InitGenesis(ctx, jsonCodec, message)
	if err != nil {
		panic(err)
	}
	return update
}

// ExportGenesis implements AppModule/ExportGenesis.
func (m Module) ExportGenesis(ctx sdk.Context, jsonCodec codec.JSONCodec) json.RawMessage {
	json, err := m.srv.ExportGenesis(ctx, jsonCodec)
	if err != nil {
		panic(err)
	}
	return json
}

// DefaultGenesis implements AppModule/DefaultGenesis.
func (m Module) DefaultGenesis(_ codec.JSONCodec) json.RawMessage {
	return nil
}

// ValidateGenesis implements AppModule/ValidateGenesis.
func (m Module) ValidateGenesis(_ codec.JSONCodec, _ sdkclient.TxEncodingConfig, bz json.RawMessage) error {
	db, err := ormdb.NewModuleDB(&example.ModuleSchema, ormdb.ModuleDBOptions{})
	if err != nil {
		return err
	}

	src, err := ormjson.NewRawMessageSource(bz)
	if err != nil {
		return err
	}

	err = db.ValidateJSON(src)
	if err != nil {
		return err
	}

	return m.srv.ValidateGenesis(bz)
}

// GetTxCmd implements AppModule/GetTxCmd.
func (m Module) GetTxCmd() *cobra.Command {
	return client.TxCmd()
}

// GetQueryCmd implements AppModule/GetQueryCmd.
func (m Module) GetQueryCmd() *cobra.Command {
	return client.QueryCmd()
}

// LegacyQuerierHandler implements AppModule/LegacyQuerierHandler.
func (m Module) LegacyQuerierHandler(_ *codec.LegacyAmino) sdk.Querier {
	return nil
}