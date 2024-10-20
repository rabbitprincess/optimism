package keeper_test

import (
	"testing"

	testkeeper "github.com/ethereum-optimism/optimism/x/testutil/keeper"
	"github.com/ethereum-optimism/optimism/x/btclightclient/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.BTCLightClientKeeper(t)
	// using nil as empty params list as, default proto decoder deserializes empty list as nil
	params := types.NewParams(nil)

	err := k.SetParams(ctx, params)
	require.NoError(t, err)

	retrievedParams := k.GetParams(ctx)
	require.EqualValues(t, params, retrievedParams)
}
