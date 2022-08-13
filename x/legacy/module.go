package legacy

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"

	"bitbucket.org/decimalteam/go-smart-node/x/legacy/keeper"
	"bitbucket.org/decimalteam/go-smart-node/x/legacy/types"
)

var (
	_ module.AppModule           = AppModule{}
	_ module.AppModuleBasic      = AppModuleBasic{}
	_ module.AppModuleSimulation = AppModule{}
)

////////////////////////////////////////////////////////////////
// AppModuleBasic
////////////////////////////////////////////////////////////////

// AppModuleBasic implements the AppModuleBasic interface for the module.
type AppModuleBasic struct {
	cdc codec.BinaryCodec
}

func NewAppModuleBasic(cdc codec.BinaryCodec) AppModuleBasic {
	return AppModuleBasic{cdc: cdc}
}

// Name returns the module's name.
func (AppModuleBasic) Name() string {
	return types.ModuleName
}

// ConsensusVersion returns the consensus state-breaking version for the module.
func (AppModuleBasic) ConsensusVersion() uint64 {
	return 1
}

// RegisterLegacyAminoCodec performs a no-op as the module doesn't support Amino encoding.
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	//
}

// RegisterInterfaces registers the module's interface types.
func (AppModuleBasic) RegisterInterfaces(interfaceRegistry codectypes.InterfaceRegistry) {
	types.RegisterInterfaces(interfaceRegistry)
}

// DefaultGenesis returns the module's default genesis state.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(types.DefaultGenesisState())
}

// ValidateGenesis performs genesis state validation for the module.
func (b AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, config client.TxEncodingConfig, bz json.RawMessage) error {
	var genesisState types.GenesisState
	if err := cdc.UnmarshalJSON(bz, &genesisState); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", types.ModuleName, err)
	}
	return genesisState.Validate()
}

// RegisterRESTRoutes registers the module's REST service handlers.
func (AppModuleBasic) RegisterRESTRoutes(clientCtx client.Context, rtr *mux.Router) {

}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the module.
func (b AppModuleBasic) RegisterGRPCGatewayRoutes(c client.Context, serveMux *runtime.ServeMux) {
	if err := types.RegisterQueryHandlerClient(context.Background(), serveMux, types.NewQueryClient(c)); err != nil {
		panic(err)
	}
}

// GetTxCmd returns the module's root tx command.
func (AppModuleBasic) GetTxCmd() *cobra.Command {
	return nil
}

// GetQueryCmd returns the module's root query command.
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return nil
}

////////////////////////////////////////////////////////////////
// AppModule
////////////////////////////////////////////////////////////////

// AppModule implements the AppModule interface for the module.
type AppModule struct {
	AppModuleBasic

	keeper         keeper.Keeper
	accountKeeper  authkeeper.AccountKeeper
	bankKeeper     types.BankKeeper
	nftKeeper      types.NftKeeper
	multisigKeeper types.MultisigKeeper
}

// NewAppModule creates a new AppModule Object
func NewAppModule(
	cdc codec.Codec,
	keeper keeper.Keeper,
) AppModule {
	return AppModule{
		AppModuleBasic: NewAppModuleBasic(cdc),
		keeper:         keeper,
	}
}

// Name returns the module's name.
func (AppModule) Name() string {
	return types.ModuleName
}

func (am AppModule) NewHandler() sdk.Handler {
	return NewHandler(&am.keeper)
}

// Route returns the module's message routing key.
func (am AppModule) Route() sdk.Route {
	return sdk.NewRoute(types.RouterKey, am.NewHandler())
}

// QuerierRoute returns the module's query routing key.
func (am AppModule) QuerierRoute() string {
	return types.RouterKey
}

// LegacyQuerierHandler returns the module's Querier.
func (am AppModule) LegacyQuerierHandler(amino *codec.LegacyAmino) sdk.Querier {
	return keeper.NewQuerier(am.keeper, amino)
}

// RegisterServices registers a GRPC query service to respond to the module-specific GRPC queries.
func (am AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterMsgServer(cfg.MsgServer(), am.keeper)
	types.RegisterQueryServer(cfg.QueryServer(), am.keeper)
}

// RegisterInvariants registers the module's invariants.
func (am AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {
	//
}

// InitGenesis performs the module's genesis initialization.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) []abci.ValidatorUpdate {
	var gs types.GenesisState
	cdc.MustUnmarshalJSON(data, &gs)
	InitGenesis(ctx, am.keeper, gs)
	return []abci.ValidatorUpdate{}
}

// ExportGenesis returns the module's exported genesis state as raw JSON bytes.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	gs := ExportGenesis(ctx, am.keeper)
	return cdc.MustMarshalJSON(gs)
}

// ConsensusVersion implements ConsensusVersion.
func (am AppModule) ConsensusVersion() uint64 {
	return am.AppModuleBasic.ConsensusVersion()
}

// BeginBlock executes all ABCI BeginBlock logic respective to the module.
func (am AppModule) BeginBlock(_ sdk.Context, _ abci.RequestBeginBlock) {
	//
}

// EndBlock executes all ABCI EndBlock logic respective to the module.
func (am AppModule) EndBlock(_ sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}