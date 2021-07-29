package keeper

import (
	"context"
	"testing"

	"github.com/MasterPi-2124/haya/x/haya/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	keeper, ctx := setupKeeper(t)
	return NewMsgServerImpl(*keeper), sdk.WrapSDKContext(ctx)
}
