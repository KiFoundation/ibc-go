package keeper_test

import (
	"github.com/cosmos/ibc-go/v7/modules/core/03-connection/types"
	ibcexported "github.com/cosmos/ibc-go/v7/modules/core/exported"
)

// TestDefaultSetParams tests the default params set are what is expected
func (suite *KeeperTestSuite) TestDefaultSetParams() {
	expParams := types.DefaultParams()

	params := suite.chainA.App.GetIBCKeeper().ConnectionKeeper.GetParams(suite.chainA.GetContext())
	suite.Require().Equal(expParams, params)
}

// TestParams tests that param setting and retrieval works properly
func (suite *KeeperTestSuite) TestParams() {
	testCases := []struct {
		name    string
		input   types.Params
		expPass bool
	}{
		{"success: set default params", types.DefaultParams(), true},
		{"success: valid value for MaxExpectedTimePerBlock", types.NewParams(10), true},
		{"failure: invalid value for MaxExpectedTimePerBlock", types.NewParams(0), false},
	}

	for _, tc := range testCases {
		tc := tc

		suite.Run(tc.name, func() {
			suite.SetupTest() // reset
			ctx := suite.chainA.GetContext()
			err := tc.input.Validate()
			suite.chainA.GetSimApp().IBCKeeper.ConnectionKeeper.SetParams(ctx, tc.input)
			if tc.expPass {
				suite.Require().NoError(err)
				expected := tc.input
				p := suite.chainA.GetSimApp().IBCKeeper.ConnectionKeeper.GetParams(ctx)
				suite.Require().Equal(expected, p)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

// TestUnsetParams tests that trying to get params that are not set panics.
func (suite *KeeperTestSuite) TestUnsetParams() {
	suite.SetupTest()
	ctx := suite.chainA.GetContext()
	store := ctx.KVStore(suite.chainA.GetSimApp().GetKey(ibcexported.StoreKey))
	store.Delete([]byte(types.ParamsKey))

	suite.Require().Equal(suite.chainA.GetSimApp().IBCKeeper.ConnectionKeeper.GetParams(ctx), types.Params{})
}
