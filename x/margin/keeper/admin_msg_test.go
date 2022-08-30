//go:build FEATURE_TOGGLE_MARGIN_CLI_ALPHA
// +build FEATURE_TOGGLE_MARGIN_CLI_ALPHA

package keeper_test

import (
	sifapp "github.com/Sifchain/sifnode/app"
	admintypes "github.com/Sifchain/sifnode/x/admin/types"
	clptypes "github.com/Sifchain/sifnode/x/clp/types"
	"github.com/Sifchain/sifnode/x/margin/keeper"
	"github.com/Sifchain/sifnode/x/margin/test"
	"github.com/Sifchain/sifnode/x/margin/types"
	tokenregistrytypes "github.com/Sifchain/sifnode/x/tokenregistry/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMsgServer_AdminClose(t *testing.T) {
	table := []struct {
		msgOpen                       types.MsgOpen
		msgForceClose                 types.MsgForceClose
		msgAdminClose                 types.MsgAdminClose
		health                        sdk.Dec
		forceCloseThreshold           sdk.Dec
		name                          string
		poolAsset                     string
		token                         string
		overrideSigner                string
		overrideForceCloseThreadshold string
		err                           error
		errString                     error
		err2                          error
		errString2                    error
		poolEnabled                   bool
		fundedAccount                 bool
		mtpCreateDisabled             bool
		repay                         bool
	}{
		{
			name: "admin close all mtps automatically",
			msgForceClose: types.MsgForceClose{
				Signer:     "sif1azpar20ck9lpys89r8x7zc8yu0qzgvtp48ng5v",
				MtpAddress: "sif1azpar20ck9lpys89r8x7zc8yu0qzgvtp48ng5v",
				Id:         1,
			},
			msgAdminClose: types.MsgAdminClose{Signer: "sif1azpar20ck9lpys89r8x7zc8yu0qzgvtp48ng5v", CloseAll: true},
			msgOpen: types.MsgOpen{
				CollateralAsset: "rowan",
				BorrowAsset:     "xxx",
				Position:        types.Position_LONG,
			},
			health:              sdk.NewDecWithPrec(1, 2),
			forceCloseThreshold: sdk.OneDec(),
			poolAsset:           "xxx",
			token:               "xxx",
			poolEnabled:         true,
			fundedAccount:       true,
			err2:                types.ErrMTPDoesNotExist,
			repay:               false,
		},
		{
			name: "admin update repay only",
			msgForceClose: types.MsgForceClose{
				Signer:     "sif1azpar20ck9lpys89r8x7zc8yu0qzgvtp48ng5v",
				MtpAddress: "sif1azpar20ck9lpys89r8x7zc8yu0qzgvtp48ng5v",
				Id:         1,
			},
			msgAdminClose: types.MsgAdminClose{Signer: "sif1azpar20ck9lpys89r8x7zc8yu0qzgvtp48ng5v", CloseAll: false},
			msgOpen: types.MsgOpen{
				CollateralAsset: "rowan",
				BorrowAsset:     "xxx",
				Position:        types.Position_LONG,
			},
			health:              sdk.NewDecWithPrec(1, 2),
			forceCloseThreshold: sdk.OneDec(),
			poolAsset:           "xxx",
			token:               "xxx",
			poolEnabled:         true,
			fundedAccount:       true,
			err2:                nil,
			repay:               false,
		},
	}
	for _, tt := range table {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			asset := clptypes.Asset{Symbol: tt.poolAsset}

			ctx, app := test.CreateTestAppMarginFromGenesis(false, func(app *sifapp.SifchainApp, genesisState sifapp.GenesisState) sifapp.GenesisState {
				gs1 := &admintypes.GenesisState{
					AdminAccounts: []*admintypes.AdminAccount{
						{
							AdminType:    admintypes.AdminType_MARGIN,
							AdminAddress: tt.msgAdminClose.Signer,
						},
						{
							AdminType:    admintypes.AdminType_CLPDEX,
							AdminAddress: tt.msgAdminClose.Signer,
						},
						{
							AdminType:    admintypes.AdminType_TOKENREGISTRY,
							AdminAddress: tt.msgAdminClose.Signer,
						},
					},
				}
				bz, _ := app.AppCodec().MarshalJSON(gs1)
				genesisState["admin"] = bz

				gs2 := &tokenregistrytypes.GenesisState{
					Registry: &tokenregistrytypes.Registry{
						Entries: []*tokenregistrytypes.RegistryEntry{
							{Denom: tt.token, BaseDenom: tt.token, Decimals: 18, Permissions: []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP}},
						},
					},
				}
				bz, _ = app.AppCodec().MarshalJSON(gs2)
				genesisState["tokenregistry"] = bz

				gs3 := &types.GenesisState{
					Params: &types.Params{
						LeverageMax:                              sdk.NewDec(2),
						InterestRateMax:                          sdk.NewDec(1),
						InterestRateMin:                          sdk.ZeroDec(),
						InterestRateIncrease:                     sdk.NewDecWithPrec(1, 1),
						InterestRateDecrease:                     sdk.NewDecWithPrec(1, 1),
						HealthGainFactor:                         sdk.NewDecWithPrec(1, 2),
						EpochLength:                              0,
						ForceCloseThreshold:                      tt.forceCloseThreshold,
						RemovalQueueThreshold:                    sdk.ZeroDec(),
						Pools:                                    []string{},
						ForceCloseFundPercentage:                 sdk.NewDecWithPrec(1, 1),
						ForceCloseInsuranceFundAddress:           "sif1syavy2npfyt9tcncdtsdzf7kny9lh777yqc2nd",
						IncrementalInterestPaymentFundPercentage: sdk.NewDecWithPrec(1, 1),
						IncrementalInterestPaymentInsuranceFundAddress: "sif1syavy2npfyt9tcncdtsdzf7kny9lh777yqc2nd",
						IncrementalInterestPaymentEnabled:              false,
						PoolOpenThreshold:                              sdk.NewDecWithPrec(1, 1),
						MaxOpenPositions:                               10000,
						SqModifier:                                     sdk.MustNewDecFromStr("10000000000000000000000000"),
						SafetyFactor:                                   sdk.MustNewDecFromStr("0.0"),
					},
				}

				if tt.poolEnabled {
					gs3.Params.Pools = []string{
						tt.poolAsset,
					}
				}

				bz, _ = app.AppCodec().MarshalJSON(gs3)
				genesisState["margin"] = bz

				nativeAsset := tt.msgOpen.CollateralAsset
				externalAsset := clptypes.Asset{Symbol: tt.msgOpen.BorrowAsset}

				nativeCoin := sdk.NewCoin(nativeAsset, sdk.Int(sdk.NewUint(1000000000000)))
				externalCoin := sdk.NewCoin(externalAsset.Symbol, sdk.Int(sdk.NewUint(1000000000000)))

				balances := []banktypes.Balance{
					{
						Address: tt.msgAdminClose.Signer,
						Coins: sdk.Coins{
							nativeCoin,
							externalCoin,
						},
					},
				}

				gs4 := banktypes.DefaultGenesisState()
				gs4.Balances = append(gs4.Balances, balances...)
				bz, _ = app.AppCodec().MarshalJSON(gs4)
				genesisState["bank"] = bz

				gs5 := &clptypes.GenesisState{
					Params: clptypes.Params{
						MinCreatePoolThreshold: 100,
					},
					AddressWhitelist: []string{
						tt.msgAdminClose.Signer,
					},
					PoolList: []*clptypes.Pool{
						{
							ExternalAsset:        &asset,
							NativeAssetBalance:   sdk.NewUint(1000000000),
							NativeLiabilities:    sdk.NewUint(1000000000),
							ExternalCustody:      sdk.NewUint(1000000000),
							ExternalAssetBalance: sdk.NewUint(1000000000),
							ExternalLiabilities:  sdk.NewUint(1000000000),
							NativeCustody:        sdk.NewUint(1000000000),
							PoolUnits:            sdk.NewUint(1),
							Health:               sdk.NewDec(1),
						},
					},
					LiquidityProviders: []*clptypes.LiquidityProvider{
						{
							Asset:                    &clptypes.Asset{Symbol: tt.poolAsset},
							LiquidityProviderAddress: tt.msgAdminClose.Signer,
							LiquidityProviderUnits:   sdk.NewUint(1000000000),
						},
					},
				}
				bz, _ = app.AppCodec().MarshalJSON(gs5)
				genesisState["clp"] = bz

				return genesisState
			})
			marginKeeper := app.MarginKeeper
			msgServer := keeper.NewMsgServerImpl(marginKeeper)

			if tt.poolEnabled {
				marginKeeper.SetEnabledPools(ctx, []string{tt.poolAsset})
			}

			var address string

			address = tt.msgAdminClose.Signer

			msg := tt.msgAdminClose
			msg.Signer = address

			var signer = msg.Signer
			if tt.overrideSigner != "" {
				signer = tt.overrideSigner
			}

			if !tt.mtpCreateDisabled {
				addMTPKey(t, ctx, app, marginKeeper, tt.msgOpen.CollateralAsset, tt.msgOpen.BorrowAsset, signer, tt.msgOpen.Position, 1, sdk.NewDec(20))
			}

			_, got := msgServer.AdminClose(sdk.WrapSDKContext(ctx), &msg)
			balanceOriginal, _ := app.BankKeeper.Balance(sdk.WrapSDKContext(ctx), &banktypes.QueryBalanceRequest{
				Address: signer,
				Denom:   tt.msgOpen.CollateralAsset,
			})
			if tt.errString != nil {
				require.EqualError(t, got, tt.errString.Error())
			} else if tt.err == nil {
				require.NoError(t, got)
			} else {
				require.ErrorIs(t, got, tt.err)
			}

			marginKeeper.BeginBlocker(ctx)

			_, got2 := marginKeeper.GetMTP(ctx, signer, 1)
			balanceAfter, _ := app.BankKeeper.Balance(sdk.WrapSDKContext(ctx), &banktypes.QueryBalanceRequest{
				Address: signer,
				Denom:   tt.msgOpen.CollateralAsset,
			})
			assert.NotEqual(t, tt.repay, balanceAfter.Balance.IsEqual(*balanceOriginal.Balance))

			if tt.errString2 != nil {
				require.EqualError(t, got2, tt.errString2.Error())
			} else if tt.err2 == nil {
				require.NoError(t, got2)
			} else {
				require.ErrorIs(t, got2, tt.err2)
			}
		})
	}
}
