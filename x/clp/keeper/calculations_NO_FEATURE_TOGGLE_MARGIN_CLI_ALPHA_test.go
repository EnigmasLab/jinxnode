//go:build !FEATURE_TOGGLE_MARGIN_CLI_ALPHA
// +build !FEATURE_TOGGLE_MARGIN_CLI_ALPHA

package keeper_test

import (
	"errors"
	"testing"

	sifapp "github.com/Sifchain/sifnode/app"

	clpkeeper "github.com/Sifchain/sifnode/x/clp/keeper"
	tokenregistrytypes "github.com/Sifchain/sifnode/x/tokenregistry/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/Sifchain/sifnode/x/clp/test"
	"github.com/Sifchain/sifnode/x/clp/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func FEATURE_TOGGLE_MARGIN_CLI_ALPHA_SwapOne(ctx sdk.Context,
	k clpkeeper.Keeper,
	sentAsset types.Asset,
	sentAmount sdk.Uint,
	nativeAsset types.Asset,
	inPool types.Pool,
	pmtpCurrentRunningRate sdk.Dec) (sdk.Uint, sdk.Uint, sdk.Uint, types.Pool, error) {
	return clpkeeper.SwapOne(sentAsset, sentAmount, nativeAsset, inPool, pmtpCurrentRunningRate)
}

func FEATURE_TOGGLE_MARGIN_CLI_ALPHA_GetSwapFee(ctx sdk.Context,
	k clpkeeper.Keeper,
	ReceivedAsset *types.Asset,
	liquidityFeeNative sdk.Uint,
	outPool types.Pool,
	pmtpCurrentRunningRate sdk.Dec) sdk.Uint {
	return clpkeeper.GetSwapFee(liquidityFeeNative, *ReceivedAsset, outPool, pmtpCurrentRunningRate)
}

func TestKeeper_SwapOneFromGenesis(t *testing.T) {
	const address = "sif1syavy2npfyt9tcncdtsdzf7kny9lh777yqc2nd"
	SwapPriceNative := sdk.ZeroDec()
	SwapPriceExternal := sdk.ZeroDec()

	testcases := []struct {
		name                   string
		poolAsset              string
		address                string
		calculateWithdraw      bool
		adjustExternalToken    bool
		nativeBalance          sdk.Int
		externalBalance        sdk.Int
		wBasis                 sdk.Int
		asymmetry              sdk.Int
		nativeAssetAmount      sdk.Uint
		externalAssetAmount    sdk.Uint
		poolUnits              sdk.Uint
		swapAmount             sdk.Uint
		swapResult             sdk.Uint
		liquidityFee           sdk.Uint
		priceImpact            sdk.Uint
		normalizationFactor    sdk.Dec
		pmtpCurrentRunningRate sdk.Dec
		from                   types.Asset
		to                     types.Asset
		expectedPool           types.Pool
		err                    error
		errString              error
	}{
		{
			name:                   "successful swap with single pool units",
			poolAsset:              "eth",
			address:                address,
			nativeBalance:          sdk.NewInt(10000),
			externalBalance:        sdk.NewInt(10000),
			nativeAssetAmount:      sdk.NewUint(998),
			externalAssetAmount:    sdk.NewUint(998),
			poolUnits:              sdk.OneUint(),
			calculateWithdraw:      true,
			wBasis:                 sdk.NewInt(1000),
			asymmetry:              sdk.NewInt(10000),
			pmtpCurrentRunningRate: sdk.OneDec(),
			swapResult:             sdk.NewUint(19),
			liquidityFee:           sdk.NewUint(978),
			priceImpact:            sdk.ZeroUint(),
			expectedPool: types.Pool{
				ExternalAsset:                 &types.Asset{Symbol: "eth"},
				NativeAssetBalance:            sdk.NewUint(100598),
				ExternalAssetBalance:          sdk.NewUint(979),
				PoolUnits:                     sdk.NewUint(1),
				SwapPriceNative:               &SwapPriceNative,
				SwapPriceExternal:             &SwapPriceExternal,
				RewardPeriodNativeDistributed: sdk.ZeroUint(),
			},
		},
		{
			name:                   "successful swap with equal amount of pool units",
			poolAsset:              "eth",
			address:                address,
			nativeBalance:          sdk.NewInt(10000),
			externalBalance:        sdk.NewInt(10000),
			nativeAssetAmount:      sdk.NewUint(998),
			externalAssetAmount:    sdk.NewUint(998),
			poolUnits:              sdk.NewUint(998),
			calculateWithdraw:      true,
			wBasis:                 sdk.NewInt(1000),
			asymmetry:              sdk.NewInt(10000),
			pmtpCurrentRunningRate: sdk.OneDec(),
			swapResult:             sdk.NewUint(165),
			liquidityFee:           sdk.NewUint(8),
			priceImpact:            sdk.ZeroUint(),
			expectedPool: types.Pool{
				ExternalAsset:                 &types.Asset{Symbol: "eth"},
				NativeAssetBalance:            sdk.NewUint(1098),
				ExternalAssetBalance:          sdk.NewUint(833),
				PoolUnits:                     sdk.NewUint(998),
				SwapPriceNative:               &SwapPriceNative,
				SwapPriceExternal:             &SwapPriceExternal,
				RewardPeriodNativeDistributed: sdk.ZeroUint(),
			},
		},
		{
			name:                   "failed swap with empty pool",
			poolAsset:              "eth",
			address:                address,
			nativeBalance:          sdk.NewInt(10000),
			externalBalance:        sdk.NewInt(10000),
			nativeAssetAmount:      sdk.NewUint(0),
			externalAssetAmount:    sdk.NewUint(0),
			poolUnits:              sdk.NewUint(0),
			calculateWithdraw:      false,
			normalizationFactor:    sdk.NewDec(0),
			adjustExternalToken:    true,
			swapAmount:             sdk.NewUint(0),
			pmtpCurrentRunningRate: sdk.OneDec(),
			swapResult:             sdk.NewUint(166),
			liquidityFee:           sdk.NewUint(8),
			priceImpact:            sdk.ZeroUint(),
			expectedPool: types.Pool{
				ExternalAsset:                 &types.Asset{Symbol: "eth"},
				NativeAssetBalance:            sdk.NewUint(1098),
				ExternalAssetBalance:          sdk.NewUint(833),
				PoolUnits:                     sdk.NewUint(998),
				SwapPriceNative:               &SwapPriceNative,
				SwapPriceExternal:             &SwapPriceExternal,
				RewardPeriodNativeDistributed: sdk.ZeroUint(),
			},
			errString: errors.New("not enough received asset tokens to swap"),
		},
		{
			name:                   "successful swap by inversing from/to assets",
			poolAsset:              "eth",
			address:                address,
			nativeBalance:          sdk.NewInt(10000),
			externalBalance:        sdk.NewInt(10000),
			nativeAssetAmount:      sdk.NewUint(998),
			externalAssetAmount:    sdk.NewUint(998),
			poolUnits:              sdk.NewUint(998),
			calculateWithdraw:      true,
			wBasis:                 sdk.NewInt(1000),
			asymmetry:              sdk.NewInt(10000),
			from:                   types.Asset{Symbol: "eth"},
			to:                     types.Asset{Symbol: "rowan"},
			pmtpCurrentRunningRate: sdk.OneDec(),
			swapResult:             sdk.NewUint(41),
			liquidityFee:           sdk.NewUint(8),
			priceImpact:            sdk.ZeroUint(),
			expectedPool: types.Pool{
				ExternalAsset:                 &types.Asset{Symbol: "eth"},
				NativeAssetBalance:            sdk.NewUint(957),
				ExternalAssetBalance:          sdk.NewUint(1098),
				PoolUnits:                     sdk.NewUint(998),
				SwapPriceNative:               &SwapPriceNative,
				SwapPriceExternal:             &SwapPriceExternal,
				RewardPeriodNativeDistributed: sdk.ZeroUint(),
			},
		},
		{
			name:                   "successful swap with pmtp current running rate value at 0.0",
			poolAsset:              "eth",
			address:                address,
			nativeBalance:          sdk.NewInt(10000),
			externalBalance:        sdk.NewInt(10000),
			nativeAssetAmount:      sdk.NewUint(998),
			externalAssetAmount:    sdk.NewUint(998),
			poolUnits:              sdk.NewUint(998),
			calculateWithdraw:      true,
			wBasis:                 sdk.NewInt(1000),
			asymmetry:              sdk.NewInt(10000),
			pmtpCurrentRunningRate: sdk.MustNewDecFromStr("0.0"),
			swapResult:             sdk.NewUint(82),
			liquidityFee:           sdk.NewUint(8),
			priceImpact:            sdk.ZeroUint(),
			expectedPool: types.Pool{
				ExternalAsset:                 &types.Asset{Symbol: "eth"},
				NativeAssetBalance:            sdk.NewUint(1098),
				ExternalAssetBalance:          sdk.NewUint(916),
				PoolUnits:                     sdk.NewUint(998),
				SwapPriceNative:               &SwapPriceNative,
				SwapPriceExternal:             &SwapPriceExternal,
				RewardPeriodNativeDistributed: sdk.ZeroUint(),
			},
		},
		{
			name:                   "successful swap with pmtp current running rate value at 0.1",
			poolAsset:              "eth",
			address:                address,
			nativeBalance:          sdk.NewInt(10000),
			externalBalance:        sdk.NewInt(10000),
			nativeAssetAmount:      sdk.NewUint(998),
			externalAssetAmount:    sdk.NewUint(998),
			poolUnits:              sdk.NewUint(998),
			calculateWithdraw:      true,
			wBasis:                 sdk.NewInt(1000),
			asymmetry:              sdk.NewInt(10000),
			pmtpCurrentRunningRate: sdk.MustNewDecFromStr("0.1"),
			swapResult:             sdk.NewUint(90),
			liquidityFee:           sdk.NewUint(8),
			priceImpact:            sdk.ZeroUint(),
			expectedPool: types.Pool{
				ExternalAsset:                 &types.Asset{Symbol: "eth"},
				NativeAssetBalance:            sdk.NewUint(1098),
				ExternalAssetBalance:          sdk.NewUint(908),
				PoolUnits:                     sdk.NewUint(998),
				SwapPriceNative:               &SwapPriceNative,
				SwapPriceExternal:             &SwapPriceExternal,
				RewardPeriodNativeDistributed: sdk.ZeroUint(),
			},
		},
		{
			name:                   "successful swap with pmtp current running rate value at 0.2",
			poolAsset:              "eth",
			address:                address,
			nativeBalance:          sdk.NewInt(10000),
			externalBalance:        sdk.NewInt(10000),
			nativeAssetAmount:      sdk.NewUint(998),
			externalAssetAmount:    sdk.NewUint(998),
			poolUnits:              sdk.NewUint(998),
			calculateWithdraw:      true,
			wBasis:                 sdk.NewInt(1000),
			asymmetry:              sdk.NewInt(10000),
			pmtpCurrentRunningRate: sdk.MustNewDecFromStr("0.2"),
			swapResult:             sdk.NewUint(99),
			liquidityFee:           sdk.NewUint(8),
			priceImpact:            sdk.ZeroUint(),
			expectedPool: types.Pool{
				ExternalAsset:                 &types.Asset{Symbol: "eth"},
				NativeAssetBalance:            sdk.NewUint(1098),
				ExternalAssetBalance:          sdk.NewUint(899),
				PoolUnits:                     sdk.NewUint(998),
				SwapPriceNative:               &SwapPriceNative,
				SwapPriceExternal:             &SwapPriceExternal,
				RewardPeriodNativeDistributed: sdk.ZeroUint(),
			},
		},
		{
			name:                   "successful swap with pmtp current running rate value at 0.3",
			poolAsset:              "eth",
			address:                address,
			nativeBalance:          sdk.NewInt(10000),
			externalBalance:        sdk.NewInt(10000),
			nativeAssetAmount:      sdk.NewUint(998),
			externalAssetAmount:    sdk.NewUint(998),
			poolUnits:              sdk.NewUint(998),
			calculateWithdraw:      true,
			wBasis:                 sdk.NewInt(1000),
			asymmetry:              sdk.NewInt(10000),
			pmtpCurrentRunningRate: sdk.MustNewDecFromStr("0.3"),
			swapResult:             sdk.NewUint(107),
			liquidityFee:           sdk.NewUint(8),
			priceImpact:            sdk.ZeroUint(),
			expectedPool: types.Pool{
				ExternalAsset:                 &types.Asset{Symbol: "eth"},
				NativeAssetBalance:            sdk.NewUint(1098),
				ExternalAssetBalance:          sdk.NewUint(891),
				PoolUnits:                     sdk.NewUint(998),
				SwapPriceNative:               &SwapPriceNative,
				SwapPriceExternal:             &SwapPriceExternal,
				RewardPeriodNativeDistributed: sdk.ZeroUint(),
			},
		},
		{
			name:                   "successful swap with pmtp current running rate value at 0.4",
			poolAsset:              "eth",
			address:                address,
			nativeBalance:          sdk.NewInt(10000),
			externalBalance:        sdk.NewInt(10000),
			nativeAssetAmount:      sdk.NewUint(998),
			externalAssetAmount:    sdk.NewUint(998),
			poolUnits:              sdk.NewUint(998),
			calculateWithdraw:      true,
			wBasis:                 sdk.NewInt(1000),
			asymmetry:              sdk.NewInt(10000),
			pmtpCurrentRunningRate: sdk.MustNewDecFromStr("0.4"),
			swapResult:             sdk.NewUint(115),
			liquidityFee:           sdk.NewUint(8),
			priceImpact:            sdk.ZeroUint(),
			expectedPool: types.Pool{
				ExternalAsset:                 &types.Asset{Symbol: "eth"},
				NativeAssetBalance:            sdk.NewUint(1098),
				ExternalAssetBalance:          sdk.NewUint(883),
				PoolUnits:                     sdk.NewUint(998),
				SwapPriceNative:               &SwapPriceNative,
				SwapPriceExternal:             &SwapPriceExternal,
				RewardPeriodNativeDistributed: sdk.ZeroUint(),
			},
		},
		{
			name:                   "successful swap with pmtp current running rate value at 0.5",
			poolAsset:              "eth",
			address:                address,
			nativeBalance:          sdk.NewInt(10000),
			externalBalance:        sdk.NewInt(10000),
			nativeAssetAmount:      sdk.NewUint(998),
			externalAssetAmount:    sdk.NewUint(998),
			poolUnits:              sdk.NewUint(998),
			calculateWithdraw:      true,
			wBasis:                 sdk.NewInt(1000),
			asymmetry:              sdk.NewInt(10000),
			pmtpCurrentRunningRate: sdk.MustNewDecFromStr("0.5"),
			swapResult:             sdk.NewUint(123),
			liquidityFee:           sdk.NewUint(8),
			priceImpact:            sdk.ZeroUint(),
			expectedPool: types.Pool{
				ExternalAsset:                 &types.Asset{Symbol: "eth"},
				NativeAssetBalance:            sdk.NewUint(1098),
				ExternalAssetBalance:          sdk.NewUint(875),
				PoolUnits:                     sdk.NewUint(998),
				SwapPriceNative:               &SwapPriceNative,
				SwapPriceExternal:             &SwapPriceExternal,
				RewardPeriodNativeDistributed: sdk.ZeroUint(),
			},
		},
		{
			name:                   "successful swap with pmtp current running rate value at 0.6",
			poolAsset:              "eth",
			address:                address,
			nativeBalance:          sdk.NewInt(10000),
			externalBalance:        sdk.NewInt(10000),
			nativeAssetAmount:      sdk.NewUint(998),
			externalAssetAmount:    sdk.NewUint(998),
			poolUnits:              sdk.NewUint(998),
			calculateWithdraw:      true,
			wBasis:                 sdk.NewInt(1000),
			asymmetry:              sdk.NewInt(10000),
			pmtpCurrentRunningRate: sdk.MustNewDecFromStr("0.6"),
			swapResult:             sdk.NewUint(132),
			liquidityFee:           sdk.NewUint(8),
			priceImpact:            sdk.ZeroUint(),
			expectedPool: types.Pool{
				ExternalAsset:                 &types.Asset{Symbol: "eth"},
				NativeAssetBalance:            sdk.NewUint(1098),
				ExternalAssetBalance:          sdk.NewUint(866),
				PoolUnits:                     sdk.NewUint(998),
				SwapPriceNative:               &SwapPriceNative,
				SwapPriceExternal:             &SwapPriceExternal,
				RewardPeriodNativeDistributed: sdk.ZeroUint(),
			},
		},
		{
			name:                   "successful swap with pmtp current running rate value at 0.7",
			poolAsset:              "eth",
			address:                address,
			nativeBalance:          sdk.NewInt(10000),
			externalBalance:        sdk.NewInt(10000),
			nativeAssetAmount:      sdk.NewUint(998),
			externalAssetAmount:    sdk.NewUint(998),
			poolUnits:              sdk.NewUint(998),
			calculateWithdraw:      true,
			wBasis:                 sdk.NewInt(1000),
			asymmetry:              sdk.NewInt(10000),
			pmtpCurrentRunningRate: sdk.MustNewDecFromStr("0.7"),
			swapResult:             sdk.NewUint(140),
			liquidityFee:           sdk.NewUint(8),
			priceImpact:            sdk.ZeroUint(),
			expectedPool: types.Pool{
				ExternalAsset:                 &types.Asset{Symbol: "eth"},
				NativeAssetBalance:            sdk.NewUint(1098),
				ExternalAssetBalance:          sdk.NewUint(858),
				PoolUnits:                     sdk.NewUint(998),
				SwapPriceNative:               &SwapPriceNative,
				SwapPriceExternal:             &SwapPriceExternal,
				RewardPeriodNativeDistributed: sdk.ZeroUint(),
			},
		},
		{
			name:                   "successful swap with pmtp current running rate value at 0.8",
			poolAsset:              "eth",
			address:                address,
			nativeBalance:          sdk.NewInt(10000),
			externalBalance:        sdk.NewInt(10000),
			nativeAssetAmount:      sdk.NewUint(998),
			externalAssetAmount:    sdk.NewUint(998),
			poolUnits:              sdk.NewUint(998),
			calculateWithdraw:      true,
			wBasis:                 sdk.NewInt(1000),
			asymmetry:              sdk.NewInt(10000),
			pmtpCurrentRunningRate: sdk.MustNewDecFromStr("0.8"),
			swapResult:             sdk.NewUint(148),
			liquidityFee:           sdk.NewUint(8),
			priceImpact:            sdk.ZeroUint(),
			expectedPool: types.Pool{
				ExternalAsset:                 &types.Asset{Symbol: "eth"},
				NativeAssetBalance:            sdk.NewUint(1098),
				ExternalAssetBalance:          sdk.NewUint(850),
				PoolUnits:                     sdk.NewUint(998),
				SwapPriceNative:               &SwapPriceNative,
				SwapPriceExternal:             &SwapPriceExternal,
				RewardPeriodNativeDistributed: sdk.ZeroUint(),
			},
		},
		{
			name:                   "successful swap with pmtp current running rate value at 0.9",
			poolAsset:              "eth",
			address:                address,
			nativeBalance:          sdk.NewInt(10000),
			externalBalance:        sdk.NewInt(10000),
			nativeAssetAmount:      sdk.NewUint(998),
			externalAssetAmount:    sdk.NewUint(998),
			poolUnits:              sdk.NewUint(998),
			calculateWithdraw:      true,
			wBasis:                 sdk.NewInt(1000),
			asymmetry:              sdk.NewInt(10000),
			pmtpCurrentRunningRate: sdk.MustNewDecFromStr("0.9"),
			swapResult:             sdk.NewUint(156),
			liquidityFee:           sdk.NewUint(8),
			priceImpact:            sdk.ZeroUint(),
			expectedPool: types.Pool{
				ExternalAsset:                 &types.Asset{Symbol: "eth"},
				NativeAssetBalance:            sdk.NewUint(1098),
				ExternalAssetBalance:          sdk.NewUint(842),
				PoolUnits:                     sdk.NewUint(998),
				SwapPriceNative:               &SwapPriceNative,
				SwapPriceExternal:             &SwapPriceExternal,
				RewardPeriodNativeDistributed: sdk.ZeroUint(),
			},
		},
		{
			name:                   "successful swap with pmtp current running rate value at 1.0",
			poolAsset:              "eth",
			address:                address,
			nativeBalance:          sdk.NewInt(10000),
			externalBalance:        sdk.NewInt(10000),
			nativeAssetAmount:      sdk.NewUint(998),
			externalAssetAmount:    sdk.NewUint(998),
			poolUnits:              sdk.NewUint(998),
			calculateWithdraw:      true,
			wBasis:                 sdk.NewInt(1000),
			asymmetry:              sdk.NewInt(10000),
			pmtpCurrentRunningRate: sdk.MustNewDecFromStr("1.0"),
			swapResult:             sdk.NewUint(165),
			liquidityFee:           sdk.NewUint(8),
			priceImpact:            sdk.ZeroUint(),
			expectedPool: types.Pool{
				ExternalAsset:                 &types.Asset{Symbol: "eth"},
				NativeAssetBalance:            sdk.NewUint(1098),
				ExternalAssetBalance:          sdk.NewUint(833),
				PoolUnits:                     sdk.NewUint(998),
				SwapPriceNative:               &SwapPriceNative,
				SwapPriceExternal:             &SwapPriceExternal,
				RewardPeriodNativeDistributed: sdk.ZeroUint(),
			},
		},
		{
			name:                   "successful swap with pmtp current running rate value at 2.0",
			poolAsset:              "eth",
			address:                address,
			nativeBalance:          sdk.NewInt(10000),
			externalBalance:        sdk.NewInt(10000),
			nativeAssetAmount:      sdk.NewUint(998),
			externalAssetAmount:    sdk.NewUint(998),
			poolUnits:              sdk.NewUint(998),
			calculateWithdraw:      true,
			wBasis:                 sdk.NewInt(1000),
			asymmetry:              sdk.NewInt(10000),
			pmtpCurrentRunningRate: sdk.MustNewDecFromStr("2.0"),
			swapResult:             sdk.NewUint(247),
			liquidityFee:           sdk.NewUint(8),
			priceImpact:            sdk.ZeroUint(),
			expectedPool: types.Pool{
				ExternalAsset:                 &types.Asset{Symbol: "eth"},
				NativeAssetBalance:            sdk.NewUint(1098),
				ExternalAssetBalance:          sdk.NewUint(751),
				PoolUnits:                     sdk.NewUint(998),
				SwapPriceNative:               &SwapPriceNative,
				SwapPriceExternal:             &SwapPriceExternal,
				RewardPeriodNativeDistributed: sdk.ZeroUint(),
			},
		},
		{
			name:                   "successful swap with pmtp current running rate value at 3.0",
			poolAsset:              "eth",
			address:                address,
			nativeBalance:          sdk.NewInt(10000),
			externalBalance:        sdk.NewInt(10000),
			nativeAssetAmount:      sdk.NewUint(998),
			externalAssetAmount:    sdk.NewUint(998),
			poolUnits:              sdk.NewUint(998),
			calculateWithdraw:      true,
			wBasis:                 sdk.NewInt(1000),
			asymmetry:              sdk.NewInt(10000),
			pmtpCurrentRunningRate: sdk.MustNewDecFromStr("3.0"),
			swapResult:             sdk.NewUint(330),
			liquidityFee:           sdk.NewUint(8),
			priceImpact:            sdk.ZeroUint(),
			expectedPool: types.Pool{
				ExternalAsset:                 &types.Asset{Symbol: "eth"},
				NativeAssetBalance:            sdk.NewUint(1098),
				ExternalAssetBalance:          sdk.NewUint(668),
				PoolUnits:                     sdk.NewUint(998),
				SwapPriceNative:               &SwapPriceNative,
				SwapPriceExternal:             &SwapPriceExternal,
				RewardPeriodNativeDistributed: sdk.ZeroUint(),
			},
		},
		{
			name:                   "successful swap with pmtp current running rate value at 4.0",
			poolAsset:              "eth",
			address:                address,
			nativeBalance:          sdk.NewInt(10000),
			externalBalance:        sdk.NewInt(10000),
			nativeAssetAmount:      sdk.NewUint(998),
			externalAssetAmount:    sdk.NewUint(998),
			poolUnits:              sdk.NewUint(998),
			calculateWithdraw:      true,
			wBasis:                 sdk.NewInt(1000),
			asymmetry:              sdk.NewInt(10000),
			pmtpCurrentRunningRate: sdk.MustNewDecFromStr("4.0"),
			swapResult:             sdk.NewUint(413),
			liquidityFee:           sdk.NewUint(8),
			priceImpact:            sdk.ZeroUint(),
			expectedPool: types.Pool{
				ExternalAsset:                 &types.Asset{Symbol: "eth"},
				NativeAssetBalance:            sdk.NewUint(1098),
				ExternalAssetBalance:          sdk.NewUint(585),
				PoolUnits:                     sdk.NewUint(998),
				SwapPriceNative:               &SwapPriceNative,
				SwapPriceExternal:             &SwapPriceExternal,
				RewardPeriodNativeDistributed: sdk.ZeroUint(),
			},
		},
		{
			name:                   "successful swap with pmtp current running rate value at 5.0",
			poolAsset:              "eth",
			address:                address,
			nativeBalance:          sdk.NewInt(10000),
			externalBalance:        sdk.NewInt(10000),
			nativeAssetAmount:      sdk.NewUint(998),
			externalAssetAmount:    sdk.NewUint(998),
			poolUnits:              sdk.NewUint(998),
			calculateWithdraw:      true,
			wBasis:                 sdk.NewInt(1000),
			asymmetry:              sdk.NewInt(10000),
			pmtpCurrentRunningRate: sdk.MustNewDecFromStr("5.0"),
			swapResult:             sdk.NewUint(495),
			liquidityFee:           sdk.NewUint(8),
			priceImpact:            sdk.ZeroUint(),
			expectedPool: types.Pool{
				ExternalAsset:                 &types.Asset{Symbol: "eth"},
				NativeAssetBalance:            sdk.NewUint(1098),
				ExternalAssetBalance:          sdk.NewUint(503),
				PoolUnits:                     sdk.NewUint(998),
				SwapPriceNative:               &SwapPriceNative,
				SwapPriceExternal:             &SwapPriceExternal,
				RewardPeriodNativeDistributed: sdk.ZeroUint(),
			},
		},
		{
			name:                   "successful swap with pmtp current running rate value at 6.0",
			poolAsset:              "eth",
			address:                address,
			nativeBalance:          sdk.NewInt(10000),
			externalBalance:        sdk.NewInt(10000),
			nativeAssetAmount:      sdk.NewUint(998),
			externalAssetAmount:    sdk.NewUint(998),
			poolUnits:              sdk.NewUint(998),
			calculateWithdraw:      true,
			wBasis:                 sdk.NewInt(1000),
			asymmetry:              sdk.NewInt(10000),
			pmtpCurrentRunningRate: sdk.MustNewDecFromStr("6.0"),
			swapResult:             sdk.NewUint(578),
			liquidityFee:           sdk.NewUint(8),
			priceImpact:            sdk.ZeroUint(),
			expectedPool: types.Pool{
				ExternalAsset:                 &types.Asset{Symbol: "eth"},
				NativeAssetBalance:            sdk.NewUint(1098),
				ExternalAssetBalance:          sdk.NewUint(420),
				PoolUnits:                     sdk.NewUint(998),
				SwapPriceNative:               &SwapPriceNative,
				SwapPriceExternal:             &SwapPriceExternal,
				RewardPeriodNativeDistributed: sdk.ZeroUint(),
			},
		},
		{
			name:                   "successful swap with pmtp current running rate value at 7.0",
			poolAsset:              "eth",
			address:                address,
			nativeBalance:          sdk.NewInt(10000),
			externalBalance:        sdk.NewInt(10000),
			nativeAssetAmount:      sdk.NewUint(998),
			externalAssetAmount:    sdk.NewUint(998),
			poolUnits:              sdk.NewUint(998),
			calculateWithdraw:      true,
			wBasis:                 sdk.NewInt(1000),
			asymmetry:              sdk.NewInt(10000),
			pmtpCurrentRunningRate: sdk.MustNewDecFromStr("7.0"),
			swapResult:             sdk.NewUint(660),
			liquidityFee:           sdk.NewUint(8),
			priceImpact:            sdk.ZeroUint(),
			expectedPool: types.Pool{
				ExternalAsset:                 &types.Asset{Symbol: "eth"},
				NativeAssetBalance:            sdk.NewUint(1098),
				ExternalAssetBalance:          sdk.NewUint(338),
				PoolUnits:                     sdk.NewUint(998),
				SwapPriceNative:               &SwapPriceNative,
				SwapPriceExternal:             &SwapPriceExternal,
				RewardPeriodNativeDistributed: sdk.ZeroUint(),
			},
		},
		{
			name:                   "successful swap with pmtp current running rate value at 8.0",
			poolAsset:              "eth",
			address:                address,
			nativeBalance:          sdk.NewInt(10000),
			externalBalance:        sdk.NewInt(10000),
			nativeAssetAmount:      sdk.NewUint(998),
			externalAssetAmount:    sdk.NewUint(998),
			poolUnits:              sdk.NewUint(998),
			calculateWithdraw:      true,
			wBasis:                 sdk.NewInt(1000),
			asymmetry:              sdk.NewInt(10000),
			pmtpCurrentRunningRate: sdk.MustNewDecFromStr("8.0"),
			swapResult:             sdk.NewUint(743),
			liquidityFee:           sdk.NewUint(8),
			priceImpact:            sdk.ZeroUint(),
			expectedPool: types.Pool{
				ExternalAsset:                 &types.Asset{Symbol: "eth"},
				NativeAssetBalance:            sdk.NewUint(1098),
				ExternalAssetBalance:          sdk.NewUint(255),
				PoolUnits:                     sdk.NewUint(998),
				SwapPriceNative:               &SwapPriceNative,
				SwapPriceExternal:             &SwapPriceExternal,
				RewardPeriodNativeDistributed: sdk.ZeroUint(),
			},
		},
		{
			name:                   "successful swap with pmtp current running rate value at 9.0",
			poolAsset:              "eth",
			address:                address,
			nativeBalance:          sdk.NewInt(10000),
			externalBalance:        sdk.NewInt(10000),
			nativeAssetAmount:      sdk.NewUint(998),
			externalAssetAmount:    sdk.NewUint(998),
			poolUnits:              sdk.NewUint(998),
			calculateWithdraw:      true,
			wBasis:                 sdk.NewInt(1000),
			asymmetry:              sdk.NewInt(10000),
			pmtpCurrentRunningRate: sdk.MustNewDecFromStr("9.0"),
			swapResult:             sdk.NewUint(826),
			liquidityFee:           sdk.NewUint(8),
			priceImpact:            sdk.ZeroUint(),
			expectedPool: types.Pool{
				ExternalAsset:                 &types.Asset{Symbol: "eth"},
				NativeAssetBalance:            sdk.NewUint(1098),
				ExternalAssetBalance:          sdk.NewUint(172),
				PoolUnits:                     sdk.NewUint(998),
				SwapPriceNative:               &SwapPriceNative,
				SwapPriceExternal:             &SwapPriceExternal,
				RewardPeriodNativeDistributed: sdk.ZeroUint(),
			},
		},
		{
			name:                   "successful swap with pmtp current running rate value at 10.0",
			poolAsset:              "eth",
			address:                address,
			nativeBalance:          sdk.NewInt(10000),
			externalBalance:        sdk.NewInt(10000),
			nativeAssetAmount:      sdk.NewUint(998),
			externalAssetAmount:    sdk.NewUint(998),
			poolUnits:              sdk.NewUint(998),
			calculateWithdraw:      true,
			wBasis:                 sdk.NewInt(1000),
			asymmetry:              sdk.NewInt(10000),
			pmtpCurrentRunningRate: sdk.MustNewDecFromStr("10.0"),
			swapResult:             sdk.NewUint(908),
			liquidityFee:           sdk.NewUint(8),
			priceImpact:            sdk.ZeroUint(),
			expectedPool: types.Pool{
				ExternalAsset:                 &types.Asset{Symbol: "eth"},
				NativeAssetBalance:            sdk.NewUint(1098),
				ExternalAssetBalance:          sdk.NewUint(90),
				PoolUnits:                     sdk.NewUint(998),
				SwapPriceNative:               &SwapPriceNative,
				SwapPriceExternal:             &SwapPriceExternal,
				RewardPeriodNativeDistributed: sdk.ZeroUint(),
			},
		},
		{
			name:                   "failed swap with bigger pmtp current running rate value",
			poolAsset:              "eth",
			address:                address,
			nativeBalance:          sdk.NewInt(10000),
			externalBalance:        sdk.NewInt(10000),
			nativeAssetAmount:      sdk.NewUint(998),
			externalAssetAmount:    sdk.NewUint(998),
			poolUnits:              sdk.NewUint(998),
			calculateWithdraw:      true,
			wBasis:                 sdk.NewInt(1000),
			asymmetry:              sdk.NewInt(10000),
			pmtpCurrentRunningRate: sdk.NewDec(20),
			errString:              errors.New("not enough received asset tokens to swap"),
		},
		{
			name:                   "failed swap with bigger pmtp current running rate value",
			poolAsset:              "eth",
			address:                address,
			nativeBalance:          sdk.NewInt(10000),
			externalBalance:        sdk.NewInt(10000),
			nativeAssetAmount:      sdk.NewUint(998),
			externalAssetAmount:    sdk.NewUint(998),
			poolUnits:              sdk.NewUint(998),
			calculateWithdraw:      true,
			wBasis:                 sdk.NewInt(1000),
			asymmetry:              sdk.NewInt(10000),
			pmtpCurrentRunningRate: sdk.NewDec(20),
			errString:              errors.New("not enough received asset tokens to swap"),
		},
	}

	for _, tc := range testcases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx, app := test.CreateTestAppClpFromGenesis(false, func(app *sifapp.SifchainApp, genesisState sifapp.GenesisState) sifapp.GenesisState {
				trGs := &tokenregistrytypes.GenesisState{
					Registry: &tokenregistrytypes.Registry{
						Entries: []*tokenregistrytypes.RegistryEntry{
							{Denom: tc.poolAsset, BaseDenom: tc.poolAsset, Decimals: 18, Permissions: []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP}},
							{Denom: "rowan", BaseDenom: "rowan", Decimals: 18, Permissions: []tokenregistrytypes.Permission{tokenregistrytypes.Permission_CLP}},
						},
					},
				}
				bz, _ := app.AppCodec().MarshalJSON(trGs)
				genesisState["tokenregistry"] = bz

				balances := []banktypes.Balance{
					{
						Address: tc.address,
						Coins: sdk.Coins{
							sdk.NewCoin(tc.poolAsset, tc.externalBalance),
							sdk.NewCoin("rowan", tc.nativeBalance),
						},
					},
				}
				bankGs := banktypes.DefaultGenesisState()
				bankGs.Balances = append(bankGs.Balances, balances...)
				bz, _ = app.AppCodec().MarshalJSON(bankGs)
				genesisState["bank"] = bz

				pools := []*types.Pool{
					{
						ExternalAsset:                 &types.Asset{Symbol: tc.poolAsset},
						NativeAssetBalance:            tc.nativeAssetAmount,
						ExternalAssetBalance:          tc.externalAssetAmount,
						PoolUnits:                     tc.poolUnits,
						SwapPriceNative:               &SwapPriceNative,
						SwapPriceExternal:             &SwapPriceExternal,
						RewardPeriodNativeDistributed: sdk.ZeroUint(),
					},
				}
				lps := []*types.LiquidityProvider{
					{
						Asset:                    &types.Asset{Symbol: tc.poolAsset},
						LiquidityProviderAddress: tc.address,
						LiquidityProviderUnits:   tc.nativeAssetAmount,
					},
				}
				clpGs := types.DefaultGenesisState()
				clpGs.Params = types.Params{
					MinCreatePoolThreshold: 100,
				}
				clpGs.AddressWhitelist = append(clpGs.AddressWhitelist, tc.address)
				clpGs.PoolList = append(clpGs.PoolList, pools...)
				clpGs.LiquidityProviders = append(clpGs.LiquidityProviders, lps...)
				bz, _ = app.AppCodec().MarshalJSON(clpGs)
				genesisState["clp"] = bz

				return genesisState
			})

			pool, _ := app.ClpKeeper.GetPool(ctx, tc.poolAsset)
			lp, _ := app.ClpKeeper.GetLiquidityProvider(ctx, tc.poolAsset, tc.address)

			SwapPriceNative := sdk.ZeroDec()
			SwapPriceExternal := sdk.ZeroDec()

			require.Equal(t, pool, types.Pool{
				ExternalAsset:                 &types.Asset{Symbol: tc.poolAsset},
				NativeAssetBalance:            tc.nativeAssetAmount,
				ExternalAssetBalance:          tc.externalAssetAmount,
				PoolUnits:                     tc.poolUnits,
				SwapPriceNative:               &SwapPriceNative,
				SwapPriceExternal:             &SwapPriceExternal,
				RewardPeriodNativeDistributed: sdk.ZeroUint(),
			})

			var swapAmount sdk.Uint

			if tc.calculateWithdraw {
				_, _, _, swapAmount = clpkeeper.CalculateWithdrawal(
					pool.PoolUnits,
					pool.NativeAssetBalance.String(),
					pool.ExternalAssetBalance.String(),
					lp.LiquidityProviderUnits.String(),
					tc.wBasis.String(),
					tc.asymmetry,
				)
			} else {
				swapAmount = tc.swapAmount
			}

			from := tc.from
			if from == (types.Asset{}) {
				from = types.GetSettlementAsset()
			}
			to := tc.to
			if to == (types.Asset{}) {
				to = types.Asset{Symbol: tc.poolAsset}
			}

			swapResult, liquidityFee, priceImpact, newPool, err := FEATURE_TOGGLE_MARGIN_CLI_ALPHA_SwapOne(
				ctx,
				app.ClpKeeper,
				from,
				swapAmount,
				to,
				pool,
				tc.pmtpCurrentRunningRate,
			)

			if tc.errString != nil {
				require.EqualError(t, err, tc.errString.Error())
				return
			}
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, swapResult, tc.swapResult, "swapResult")
			require.Equal(t, liquidityFee, tc.liquidityFee)
			require.Equal(t, priceImpact, tc.priceImpact)
			require.Equal(t, newPool, tc.expectedPool)
		})
	}
}