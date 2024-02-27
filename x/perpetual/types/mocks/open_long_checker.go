// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	math "cosmossdk.io/math"
	ammtypes "github.com/elys-network/elys/x/amm/types"

	mock "github.com/stretchr/testify/mock"

	perpetualtypes "github.com/elys-network/elys/x/perpetual/types"

	types "github.com/cosmos/cosmos-sdk/types"
)

// OpenLongChecker is an autogenerated mock type for the OpenLongChecker type
type OpenLongChecker struct {
	mock.Mock
}

type OpenLongChecker_Expecter struct {
	mock *mock.Mock
}

func (_m *OpenLongChecker) EXPECT() *OpenLongChecker_Expecter {
	return &OpenLongChecker_Expecter{mock: &_m.Mock}
}

// Borrow provides a mock function with given fields: ctx, collateralAmount, custodyAmount, mtp, ammPool, pool, eta, baseCurrency
func (_m *OpenLongChecker) Borrow(ctx types.Context, collateralAmount math.Int, custodyAmount math.Int, mtp *perpetualtypes.MTP, ammPool *ammtypes.Pool, pool *perpetualtypes.Pool, eta math.LegacyDec, baseCurrency string) error {
	ret := _m.Called(ctx, collateralAmount, custodyAmount, mtp, ammPool, pool, eta, baseCurrency)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, math.Int, math.Int, *perpetualtypes.MTP, *ammtypes.Pool, *perpetualtypes.Pool, math.LegacyDec, string) error); ok {
		r0 = rf(ctx, collateralAmount, custodyAmount, mtp, ammPool, pool, eta, baseCurrency)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OpenLongChecker_Borrow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Borrow'
type OpenLongChecker_Borrow_Call struct {
	*mock.Call
}

// Borrow is a helper method to define mock.On call
//   - ctx types.Context
//   - collateralAmount math.Int
//   - custodyAmount math.Int
//   - mtp *perpetualtypes.MTP
//   - ammPool *ammtypes.Pool
//   - pool *perpetualtypes.Pool
//   - eta math.LegacyDec
//   - baseCurrency string
func (_e *OpenLongChecker_Expecter) Borrow(ctx interface{}, collateralAmount interface{}, custodyAmount interface{}, mtp interface{}, ammPool interface{}, pool interface{}, eta interface{}, baseCurrency interface{}) *OpenLongChecker_Borrow_Call {
	return &OpenLongChecker_Borrow_Call{Call: _e.mock.On("Borrow", ctx, collateralAmount, custodyAmount, mtp, ammPool, pool, eta, baseCurrency)}
}

func (_c *OpenLongChecker_Borrow_Call) Run(run func(ctx types.Context, collateralAmount math.Int, custodyAmount math.Int, mtp *perpetualtypes.MTP, ammPool *ammtypes.Pool, pool *perpetualtypes.Pool, eta math.LegacyDec, baseCurrency string)) *OpenLongChecker_Borrow_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(math.Int), args[2].(math.Int), args[3].(*perpetualtypes.MTP), args[4].(*ammtypes.Pool), args[5].(*perpetualtypes.Pool), args[6].(math.LegacyDec), args[7].(string))
	})
	return _c
}

func (_c *OpenLongChecker_Borrow_Call) Return(_a0 error) *OpenLongChecker_Borrow_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *OpenLongChecker_Borrow_Call) RunAndReturn(run func(types.Context, math.Int, math.Int, *perpetualtypes.MTP, *ammtypes.Pool, *perpetualtypes.Pool, math.LegacyDec, string) error) *OpenLongChecker_Borrow_Call {
	_c.Call.Return(run)
	return _c
}

// CalcMTPConsolidateCollateral provides a mock function with given fields: ctx, mtp, baseCurrency
func (_m *OpenLongChecker) CalcMTPConsolidateCollateral(ctx types.Context, mtp *perpetualtypes.MTP, baseCurrency string) error {
	ret := _m.Called(ctx, mtp, baseCurrency)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, *perpetualtypes.MTP, string) error); ok {
		r0 = rf(ctx, mtp, baseCurrency)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OpenLongChecker_CalcMTPConsolidateCollateral_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CalcMTPConsolidateCollateral'
type OpenLongChecker_CalcMTPConsolidateCollateral_Call struct {
	*mock.Call
}

// CalcMTPConsolidateCollateral is a helper method to define mock.On call
//   - ctx types.Context
//   - mtp *perpetualtypes.MTP
//   - baseCurrency string
func (_e *OpenLongChecker_Expecter) CalcMTPConsolidateCollateral(ctx interface{}, mtp interface{}, baseCurrency interface{}) *OpenLongChecker_CalcMTPConsolidateCollateral_Call {
	return &OpenLongChecker_CalcMTPConsolidateCollateral_Call{Call: _e.mock.On("CalcMTPConsolidateCollateral", ctx, mtp, baseCurrency)}
}

func (_c *OpenLongChecker_CalcMTPConsolidateCollateral_Call) Run(run func(ctx types.Context, mtp *perpetualtypes.MTP, baseCurrency string)) *OpenLongChecker_CalcMTPConsolidateCollateral_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(*perpetualtypes.MTP), args[2].(string))
	})
	return _c
}

func (_c *OpenLongChecker_CalcMTPConsolidateCollateral_Call) Return(_a0 error) *OpenLongChecker_CalcMTPConsolidateCollateral_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *OpenLongChecker_CalcMTPConsolidateCollateral_Call) RunAndReturn(run func(types.Context, *perpetualtypes.MTP, string) error) *OpenLongChecker_CalcMTPConsolidateCollateral_Call {
	_c.Call.Return(run)
	return _c
}

// CheckMinLiabilities provides a mock function with given fields: ctx, collateralTokenAmt, eta, ammPool, borrowAsset, baseCurrency
func (_m *OpenLongChecker) CheckMinLiabilities(ctx types.Context, collateralTokenAmt types.Coin, eta math.LegacyDec, ammPool ammtypes.Pool, borrowAsset string, baseCurrency string) error {
	ret := _m.Called(ctx, collateralTokenAmt, eta, ammPool, borrowAsset, baseCurrency)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, types.Coin, math.LegacyDec, ammtypes.Pool, string, string) error); ok {
		r0 = rf(ctx, collateralTokenAmt, eta, ammPool, borrowAsset, baseCurrency)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OpenLongChecker_CheckMinLiabilities_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CheckMinLiabilities'
type OpenLongChecker_CheckMinLiabilities_Call struct {
	*mock.Call
}

// CheckMinLiabilities is a helper method to define mock.On call
//   - ctx types.Context
//   - collateralTokenAmt types.Coin
//   - eta math.LegacyDec
//   - ammPool ammtypes.Pool
//   - borrowAsset string
//   - baseCurrency string
func (_e *OpenLongChecker_Expecter) CheckMinLiabilities(ctx interface{}, collateralTokenAmt interface{}, eta interface{}, ammPool interface{}, borrowAsset interface{}, baseCurrency interface{}) *OpenLongChecker_CheckMinLiabilities_Call {
	return &OpenLongChecker_CheckMinLiabilities_Call{Call: _e.mock.On("CheckMinLiabilities", ctx, collateralTokenAmt, eta, ammPool, borrowAsset, baseCurrency)}
}

func (_c *OpenLongChecker_CheckMinLiabilities_Call) Run(run func(ctx types.Context, collateralTokenAmt types.Coin, eta math.LegacyDec, ammPool ammtypes.Pool, borrowAsset string, baseCurrency string)) *OpenLongChecker_CheckMinLiabilities_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(types.Coin), args[2].(math.LegacyDec), args[3].(ammtypes.Pool), args[4].(string), args[5].(string))
	})
	return _c
}

func (_c *OpenLongChecker_CheckMinLiabilities_Call) Return(_a0 error) *OpenLongChecker_CheckMinLiabilities_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *OpenLongChecker_CheckMinLiabilities_Call) RunAndReturn(run func(types.Context, types.Coin, math.LegacyDec, ammtypes.Pool, string, string) error) *OpenLongChecker_CheckMinLiabilities_Call {
	_c.Call.Return(run)
	return _c
}

// CheckSameAssetPosition provides a mock function with given fields: ctx, msg
func (_m *OpenLongChecker) CheckSameAssetPosition(ctx types.Context, msg *perpetualtypes.MsgOpen) *perpetualtypes.MTP {
	ret := _m.Called(ctx, msg)

	var r0 *perpetualtypes.MTP
	if rf, ok := ret.Get(0).(func(types.Context, *perpetualtypes.MsgOpen) *perpetualtypes.MTP); ok {
		r0 = rf(ctx, msg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*perpetualtypes.MTP)
		}
	}

	return r0
}

// OpenLongChecker_CheckSameAssetPosition_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CheckSameAssetPosition'
type OpenLongChecker_CheckSameAssetPosition_Call struct {
	*mock.Call
}

// CheckSameAssetPosition is a helper method to define mock.On call
//   - ctx types.Context
//   - msg *perpetualtypes.MsgOpen
func (_e *OpenLongChecker_Expecter) CheckSameAssetPosition(ctx interface{}, msg interface{}) *OpenLongChecker_CheckSameAssetPosition_Call {
	return &OpenLongChecker_CheckSameAssetPosition_Call{Call: _e.mock.On("CheckSameAssetPosition", ctx, msg)}
}

func (_c *OpenLongChecker_CheckSameAssetPosition_Call) Run(run func(ctx types.Context, msg *perpetualtypes.MsgOpen)) *OpenLongChecker_CheckSameAssetPosition_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(*perpetualtypes.MsgOpen))
	})
	return _c
}

func (_c *OpenLongChecker_CheckSameAssetPosition_Call) Return(_a0 *perpetualtypes.MTP) *OpenLongChecker_CheckSameAssetPosition_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *OpenLongChecker_CheckSameAssetPosition_Call) RunAndReturn(run func(types.Context, *perpetualtypes.MsgOpen) *perpetualtypes.MTP) *OpenLongChecker_CheckSameAssetPosition_Call {
	_c.Call.Return(run)
	return _c
}

// EstimateSwap provides a mock function with given fields: ctx, leveragedAmtTokenIn, borrowAsset, ammPool
func (_m *OpenLongChecker) EstimateSwap(ctx types.Context, leveragedAmtTokenIn types.Coin, borrowAsset string, ammPool ammtypes.Pool) (math.Int, error) {
	ret := _m.Called(ctx, leveragedAmtTokenIn, borrowAsset, ammPool)

	var r0 math.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, types.Coin, string, ammtypes.Pool) (math.Int, error)); ok {
		return rf(ctx, leveragedAmtTokenIn, borrowAsset, ammPool)
	}
	if rf, ok := ret.Get(0).(func(types.Context, types.Coin, string, ammtypes.Pool) math.Int); ok {
		r0 = rf(ctx, leveragedAmtTokenIn, borrowAsset, ammPool)
	} else {
		r0 = ret.Get(0).(math.Int)
	}

	if rf, ok := ret.Get(1).(func(types.Context, types.Coin, string, ammtypes.Pool) error); ok {
		r1 = rf(ctx, leveragedAmtTokenIn, borrowAsset, ammPool)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OpenLongChecker_EstimateSwap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EstimateSwap'
type OpenLongChecker_EstimateSwap_Call struct {
	*mock.Call
}

// EstimateSwap is a helper method to define mock.On call
//   - ctx types.Context
//   - leveragedAmtTokenIn types.Coin
//   - borrowAsset string
//   - ammPool ammtypes.Pool
func (_e *OpenLongChecker_Expecter) EstimateSwap(ctx interface{}, leveragedAmtTokenIn interface{}, borrowAsset interface{}, ammPool interface{}) *OpenLongChecker_EstimateSwap_Call {
	return &OpenLongChecker_EstimateSwap_Call{Call: _e.mock.On("EstimateSwap", ctx, leveragedAmtTokenIn, borrowAsset, ammPool)}
}

func (_c *OpenLongChecker_EstimateSwap_Call) Run(run func(ctx types.Context, leveragedAmtTokenIn types.Coin, borrowAsset string, ammPool ammtypes.Pool)) *OpenLongChecker_EstimateSwap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(types.Coin), args[2].(string), args[3].(ammtypes.Pool))
	})
	return _c
}

func (_c *OpenLongChecker_EstimateSwap_Call) Return(_a0 math.Int, _a1 error) *OpenLongChecker_EstimateSwap_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OpenLongChecker_EstimateSwap_Call) RunAndReturn(run func(types.Context, types.Coin, string, ammtypes.Pool) (math.Int, error)) *OpenLongChecker_EstimateSwap_Call {
	_c.Call.Return(run)
	return _c
}

// EstimateSwapGivenOut provides a mock function with given fields: ctx, tokenOutAmount, tokenInDenom, ammPool
func (_m *OpenLongChecker) EstimateSwapGivenOut(ctx types.Context, tokenOutAmount types.Coin, tokenInDenom string, ammPool ammtypes.Pool) (math.Int, error) {
	ret := _m.Called(ctx, tokenOutAmount, tokenInDenom, ammPool)

	var r0 math.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, types.Coin, string, ammtypes.Pool) (math.Int, error)); ok {
		return rf(ctx, tokenOutAmount, tokenInDenom, ammPool)
	}
	if rf, ok := ret.Get(0).(func(types.Context, types.Coin, string, ammtypes.Pool) math.Int); ok {
		r0 = rf(ctx, tokenOutAmount, tokenInDenom, ammPool)
	} else {
		r0 = ret.Get(0).(math.Int)
	}

	if rf, ok := ret.Get(1).(func(types.Context, types.Coin, string, ammtypes.Pool) error); ok {
		r1 = rf(ctx, tokenOutAmount, tokenInDenom, ammPool)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OpenLongChecker_EstimateSwapGivenOut_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EstimateSwapGivenOut'
type OpenLongChecker_EstimateSwapGivenOut_Call struct {
	*mock.Call
}

// EstimateSwapGivenOut is a helper method to define mock.On call
//   - ctx types.Context
//   - tokenOutAmount types.Coin
//   - tokenInDenom string
//   - ammPool ammtypes.Pool
func (_e *OpenLongChecker_Expecter) EstimateSwapGivenOut(ctx interface{}, tokenOutAmount interface{}, tokenInDenom interface{}, ammPool interface{}) *OpenLongChecker_EstimateSwapGivenOut_Call {
	return &OpenLongChecker_EstimateSwapGivenOut_Call{Call: _e.mock.On("EstimateSwapGivenOut", ctx, tokenOutAmount, tokenInDenom, ammPool)}
}

func (_c *OpenLongChecker_EstimateSwapGivenOut_Call) Run(run func(ctx types.Context, tokenOutAmount types.Coin, tokenInDenom string, ammPool ammtypes.Pool)) *OpenLongChecker_EstimateSwapGivenOut_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(types.Coin), args[2].(string), args[3].(ammtypes.Pool))
	})
	return _c
}

func (_c *OpenLongChecker_EstimateSwapGivenOut_Call) Return(_a0 math.Int, _a1 error) *OpenLongChecker_EstimateSwapGivenOut_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OpenLongChecker_EstimateSwapGivenOut_Call) RunAndReturn(run func(types.Context, types.Coin, string, ammtypes.Pool) (math.Int, error)) *OpenLongChecker_EstimateSwapGivenOut_Call {
	_c.Call.Return(run)
	return _c
}

// GetAmmPool provides a mock function with given fields: ctx, poolId, tradingAsset
func (_m *OpenLongChecker) GetAmmPool(ctx types.Context, poolId uint64, tradingAsset string) (ammtypes.Pool, error) {
	ret := _m.Called(ctx, poolId, tradingAsset)

	var r0 ammtypes.Pool
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, uint64, string) (ammtypes.Pool, error)); ok {
		return rf(ctx, poolId, tradingAsset)
	}
	if rf, ok := ret.Get(0).(func(types.Context, uint64, string) ammtypes.Pool); ok {
		r0 = rf(ctx, poolId, tradingAsset)
	} else {
		r0 = ret.Get(0).(ammtypes.Pool)
	}

	if rf, ok := ret.Get(1).(func(types.Context, uint64, string) error); ok {
		r1 = rf(ctx, poolId, tradingAsset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OpenLongChecker_GetAmmPool_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAmmPool'
type OpenLongChecker_GetAmmPool_Call struct {
	*mock.Call
}

// GetAmmPool is a helper method to define mock.On call
//   - ctx types.Context
//   - poolId uint64
//   - tradingAsset string
func (_e *OpenLongChecker_Expecter) GetAmmPool(ctx interface{}, poolId interface{}, tradingAsset interface{}) *OpenLongChecker_GetAmmPool_Call {
	return &OpenLongChecker_GetAmmPool_Call{Call: _e.mock.On("GetAmmPool", ctx, poolId, tradingAsset)}
}

func (_c *OpenLongChecker_GetAmmPool_Call) Run(run func(ctx types.Context, poolId uint64, tradingAsset string)) *OpenLongChecker_GetAmmPool_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(uint64), args[2].(string))
	})
	return _c
}

func (_c *OpenLongChecker_GetAmmPool_Call) Return(_a0 ammtypes.Pool, _a1 error) *OpenLongChecker_GetAmmPool_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OpenLongChecker_GetAmmPool_Call) RunAndReturn(run func(types.Context, uint64, string) (ammtypes.Pool, error)) *OpenLongChecker_GetAmmPool_Call {
	_c.Call.Return(run)
	return _c
}

// GetMaxLeverageParam provides a mock function with given fields: ctx
func (_m *OpenLongChecker) GetMaxLeverageParam(ctx types.Context) math.LegacyDec {
	ret := _m.Called(ctx)

	var r0 math.LegacyDec
	if rf, ok := ret.Get(0).(func(types.Context) math.LegacyDec); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(math.LegacyDec)
	}

	return r0
}

// OpenLongChecker_GetMaxLeverageParam_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMaxLeverageParam'
type OpenLongChecker_GetMaxLeverageParam_Call struct {
	*mock.Call
}

// GetMaxLeverageParam is a helper method to define mock.On call
//   - ctx types.Context
func (_e *OpenLongChecker_Expecter) GetMaxLeverageParam(ctx interface{}) *OpenLongChecker_GetMaxLeverageParam_Call {
	return &OpenLongChecker_GetMaxLeverageParam_Call{Call: _e.mock.On("GetMaxLeverageParam", ctx)}
}

func (_c *OpenLongChecker_GetMaxLeverageParam_Call) Run(run func(ctx types.Context)) *OpenLongChecker_GetMaxLeverageParam_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context))
	})
	return _c
}

func (_c *OpenLongChecker_GetMaxLeverageParam_Call) Return(_a0 math.LegacyDec) *OpenLongChecker_GetMaxLeverageParam_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *OpenLongChecker_GetMaxLeverageParam_Call) RunAndReturn(run func(types.Context) math.LegacyDec) *OpenLongChecker_GetMaxLeverageParam_Call {
	_c.Call.Return(run)
	return _c
}

// GetPool provides a mock function with given fields: ctx, poolId
func (_m *OpenLongChecker) GetPool(ctx types.Context, poolId uint64) (perpetualtypes.Pool, bool) {
	ret := _m.Called(ctx, poolId)

	var r0 perpetualtypes.Pool
	var r1 bool
	if rf, ok := ret.Get(0).(func(types.Context, uint64) (perpetualtypes.Pool, bool)); ok {
		return rf(ctx, poolId)
	}
	if rf, ok := ret.Get(0).(func(types.Context, uint64) perpetualtypes.Pool); ok {
		r0 = rf(ctx, poolId)
	} else {
		r0 = ret.Get(0).(perpetualtypes.Pool)
	}

	if rf, ok := ret.Get(1).(func(types.Context, uint64) bool); ok {
		r1 = rf(ctx, poolId)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// OpenLongChecker_GetPool_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPool'
type OpenLongChecker_GetPool_Call struct {
	*mock.Call
}

// GetPool is a helper method to define mock.On call
//   - ctx types.Context
//   - poolId uint64
func (_e *OpenLongChecker_Expecter) GetPool(ctx interface{}, poolId interface{}) *OpenLongChecker_GetPool_Call {
	return &OpenLongChecker_GetPool_Call{Call: _e.mock.On("GetPool", ctx, poolId)}
}

func (_c *OpenLongChecker_GetPool_Call) Run(run func(ctx types.Context, poolId uint64)) *OpenLongChecker_GetPool_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(uint64))
	})
	return _c
}

func (_c *OpenLongChecker_GetPool_Call) Return(_a0 perpetualtypes.Pool, _a1 bool) *OpenLongChecker_GetPool_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OpenLongChecker_GetPool_Call) RunAndReturn(run func(types.Context, uint64) (perpetualtypes.Pool, bool)) *OpenLongChecker_GetPool_Call {
	_c.Call.Return(run)
	return _c
}

// GetSafetyFactor provides a mock function with given fields: ctx
func (_m *OpenLongChecker) GetSafetyFactor(ctx types.Context) math.LegacyDec {
	ret := _m.Called(ctx)

	var r0 math.LegacyDec
	if rf, ok := ret.Get(0).(func(types.Context) math.LegacyDec); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(math.LegacyDec)
	}

	return r0
}

// OpenLongChecker_GetSafetyFactor_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSafetyFactor'
type OpenLongChecker_GetSafetyFactor_Call struct {
	*mock.Call
}

// GetSafetyFactor is a helper method to define mock.On call
//   - ctx types.Context
func (_e *OpenLongChecker_Expecter) GetSafetyFactor(ctx interface{}) *OpenLongChecker_GetSafetyFactor_Call {
	return &OpenLongChecker_GetSafetyFactor_Call{Call: _e.mock.On("GetSafetyFactor", ctx)}
}

func (_c *OpenLongChecker_GetSafetyFactor_Call) Run(run func(ctx types.Context)) *OpenLongChecker_GetSafetyFactor_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context))
	})
	return _c
}

func (_c *OpenLongChecker_GetSafetyFactor_Call) Return(_a0 math.LegacyDec) *OpenLongChecker_GetSafetyFactor_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *OpenLongChecker_GetSafetyFactor_Call) RunAndReturn(run func(types.Context) math.LegacyDec) *OpenLongChecker_GetSafetyFactor_Call {
	_c.Call.Return(run)
	return _c
}

// IsPoolEnabled provides a mock function with given fields: ctx, poolId
func (_m *OpenLongChecker) IsPoolEnabled(ctx types.Context, poolId uint64) bool {
	ret := _m.Called(ctx, poolId)

	var r0 bool
	if rf, ok := ret.Get(0).(func(types.Context, uint64) bool); ok {
		r0 = rf(ctx, poolId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// OpenLongChecker_IsPoolEnabled_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsPoolEnabled'
type OpenLongChecker_IsPoolEnabled_Call struct {
	*mock.Call
}

// IsPoolEnabled is a helper method to define mock.On call
//   - ctx types.Context
//   - poolId uint64
func (_e *OpenLongChecker_Expecter) IsPoolEnabled(ctx interface{}, poolId interface{}) *OpenLongChecker_IsPoolEnabled_Call {
	return &OpenLongChecker_IsPoolEnabled_Call{Call: _e.mock.On("IsPoolEnabled", ctx, poolId)}
}

func (_c *OpenLongChecker_IsPoolEnabled_Call) Run(run func(ctx types.Context, poolId uint64)) *OpenLongChecker_IsPoolEnabled_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(uint64))
	})
	return _c
}

func (_c *OpenLongChecker_IsPoolEnabled_Call) Return(_a0 bool) *OpenLongChecker_IsPoolEnabled_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *OpenLongChecker_IsPoolEnabled_Call) RunAndReturn(run func(types.Context, uint64) bool) *OpenLongChecker_IsPoolEnabled_Call {
	_c.Call.Return(run)
	return _c
}

// SetMTP provides a mock function with given fields: ctx, mtp
func (_m *OpenLongChecker) SetMTP(ctx types.Context, mtp *perpetualtypes.MTP) error {
	ret := _m.Called(ctx, mtp)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, *perpetualtypes.MTP) error); ok {
		r0 = rf(ctx, mtp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OpenLongChecker_SetMTP_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetMTP'
type OpenLongChecker_SetMTP_Call struct {
	*mock.Call
}

// SetMTP is a helper method to define mock.On call
//   - ctx types.Context
//   - mtp *perpetualtypes.MTP
func (_e *OpenLongChecker_Expecter) SetMTP(ctx interface{}, mtp interface{}) *OpenLongChecker_SetMTP_Call {
	return &OpenLongChecker_SetMTP_Call{Call: _e.mock.On("SetMTP", ctx, mtp)}
}

func (_c *OpenLongChecker_SetMTP_Call) Run(run func(ctx types.Context, mtp *perpetualtypes.MTP)) *OpenLongChecker_SetMTP_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(*perpetualtypes.MTP))
	})
	return _c
}

func (_c *OpenLongChecker_SetMTP_Call) Return(_a0 error) *OpenLongChecker_SetMTP_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *OpenLongChecker_SetMTP_Call) RunAndReturn(run func(types.Context, *perpetualtypes.MTP) error) *OpenLongChecker_SetMTP_Call {
	_c.Call.Return(run)
	return _c
}

// SetPool provides a mock function with given fields: ctx, pool
func (_m *OpenLongChecker) SetPool(ctx types.Context, pool perpetualtypes.Pool) {
	_m.Called(ctx, pool)
}

// OpenLongChecker_SetPool_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetPool'
type OpenLongChecker_SetPool_Call struct {
	*mock.Call
}

// SetPool is a helper method to define mock.On call
//   - ctx types.Context
//   - pool perpetualtypes.Pool
func (_e *OpenLongChecker_Expecter) SetPool(ctx interface{}, pool interface{}) *OpenLongChecker_SetPool_Call {
	return &OpenLongChecker_SetPool_Call{Call: _e.mock.On("SetPool", ctx, pool)}
}

func (_c *OpenLongChecker_SetPool_Call) Run(run func(ctx types.Context, pool perpetualtypes.Pool)) *OpenLongChecker_SetPool_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(perpetualtypes.Pool))
	})
	return _c
}

func (_c *OpenLongChecker_SetPool_Call) Return() *OpenLongChecker_SetPool_Call {
	_c.Call.Return()
	return _c
}

func (_c *OpenLongChecker_SetPool_Call) RunAndReturn(run func(types.Context, perpetualtypes.Pool)) *OpenLongChecker_SetPool_Call {
	_c.Call.Return(run)
	return _c
}

// TakeInCustody provides a mock function with given fields: ctx, mtp, pool
func (_m *OpenLongChecker) TakeInCustody(ctx types.Context, mtp perpetualtypes.MTP, pool *perpetualtypes.Pool) error {
	ret := _m.Called(ctx, mtp, pool)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, perpetualtypes.MTP, *perpetualtypes.Pool) error); ok {
		r0 = rf(ctx, mtp, pool)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OpenLongChecker_TakeInCustody_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TakeInCustody'
type OpenLongChecker_TakeInCustody_Call struct {
	*mock.Call
}

// TakeInCustody is a helper method to define mock.On call
//   - ctx types.Context
//   - mtp perpetualtypes.MTP
//   - pool *perpetualtypes.Pool
func (_e *OpenLongChecker_Expecter) TakeInCustody(ctx interface{}, mtp interface{}, pool interface{}) *OpenLongChecker_TakeInCustody_Call {
	return &OpenLongChecker_TakeInCustody_Call{Call: _e.mock.On("TakeInCustody", ctx, mtp, pool)}
}

func (_c *OpenLongChecker_TakeInCustody_Call) Run(run func(ctx types.Context, mtp perpetualtypes.MTP, pool *perpetualtypes.Pool)) *OpenLongChecker_TakeInCustody_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(perpetualtypes.MTP), args[2].(*perpetualtypes.Pool))
	})
	return _c
}

func (_c *OpenLongChecker_TakeInCustody_Call) Return(_a0 error) *OpenLongChecker_TakeInCustody_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *OpenLongChecker_TakeInCustody_Call) RunAndReturn(run func(types.Context, perpetualtypes.MTP, *perpetualtypes.Pool) error) *OpenLongChecker_TakeInCustody_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateMTPHealth provides a mock function with given fields: ctx, mtp, ammPool, baseCurrency
func (_m *OpenLongChecker) UpdateMTPHealth(ctx types.Context, mtp perpetualtypes.MTP, ammPool ammtypes.Pool, baseCurrency string) (math.LegacyDec, error) {
	ret := _m.Called(ctx, mtp, ammPool, baseCurrency)

	var r0 math.LegacyDec
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, perpetualtypes.MTP, ammtypes.Pool, string) (math.LegacyDec, error)); ok {
		return rf(ctx, mtp, ammPool, baseCurrency)
	}
	if rf, ok := ret.Get(0).(func(types.Context, perpetualtypes.MTP, ammtypes.Pool, string) math.LegacyDec); ok {
		r0 = rf(ctx, mtp, ammPool, baseCurrency)
	} else {
		r0 = ret.Get(0).(math.LegacyDec)
	}

	if rf, ok := ret.Get(1).(func(types.Context, perpetualtypes.MTP, ammtypes.Pool, string) error); ok {
		r1 = rf(ctx, mtp, ammPool, baseCurrency)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OpenLongChecker_UpdateMTPHealth_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateMTPHealth'
type OpenLongChecker_UpdateMTPHealth_Call struct {
	*mock.Call
}

// UpdateMTPHealth is a helper method to define mock.On call
//   - ctx types.Context
//   - mtp perpetualtypes.MTP
//   - ammPool ammtypes.Pool
//   - baseCurrency string
func (_e *OpenLongChecker_Expecter) UpdateMTPHealth(ctx interface{}, mtp interface{}, ammPool interface{}, baseCurrency interface{}) *OpenLongChecker_UpdateMTPHealth_Call {
	return &OpenLongChecker_UpdateMTPHealth_Call{Call: _e.mock.On("UpdateMTPHealth", ctx, mtp, ammPool, baseCurrency)}
}

func (_c *OpenLongChecker_UpdateMTPHealth_Call) Run(run func(ctx types.Context, mtp perpetualtypes.MTP, ammPool ammtypes.Pool, baseCurrency string)) *OpenLongChecker_UpdateMTPHealth_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(perpetualtypes.MTP), args[2].(ammtypes.Pool), args[3].(string))
	})
	return _c
}

func (_c *OpenLongChecker_UpdateMTPHealth_Call) Return(_a0 math.LegacyDec, _a1 error) *OpenLongChecker_UpdateMTPHealth_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OpenLongChecker_UpdateMTPHealth_Call) RunAndReturn(run func(types.Context, perpetualtypes.MTP, ammtypes.Pool, string) (math.LegacyDec, error)) *OpenLongChecker_UpdateMTPHealth_Call {
	_c.Call.Return(run)
	return _c
}

// UpdatePoolHealth provides a mock function with given fields: ctx, pool
func (_m *OpenLongChecker) UpdatePoolHealth(ctx types.Context, pool *perpetualtypes.Pool) error {
	ret := _m.Called(ctx, pool)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, *perpetualtypes.Pool) error); ok {
		r0 = rf(ctx, pool)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OpenLongChecker_UpdatePoolHealth_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdatePoolHealth'
type OpenLongChecker_UpdatePoolHealth_Call struct {
	*mock.Call
}

// UpdatePoolHealth is a helper method to define mock.On call
//   - ctx types.Context
//   - pool *perpetualtypes.Pool
func (_e *OpenLongChecker_Expecter) UpdatePoolHealth(ctx interface{}, pool interface{}) *OpenLongChecker_UpdatePoolHealth_Call {
	return &OpenLongChecker_UpdatePoolHealth_Call{Call: _e.mock.On("UpdatePoolHealth", ctx, pool)}
}

func (_c *OpenLongChecker_UpdatePoolHealth_Call) Run(run func(ctx types.Context, pool *perpetualtypes.Pool)) *OpenLongChecker_UpdatePoolHealth_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(*perpetualtypes.Pool))
	})
	return _c
}

func (_c *OpenLongChecker_UpdatePoolHealth_Call) Return(_a0 error) *OpenLongChecker_UpdatePoolHealth_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *OpenLongChecker_UpdatePoolHealth_Call) RunAndReturn(run func(types.Context, *perpetualtypes.Pool) error) *OpenLongChecker_UpdatePoolHealth_Call {
	_c.Call.Return(run)
	return _c
}

// NewOpenLongChecker creates a new instance of OpenLongChecker. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOpenLongChecker(t interface {
	mock.TestingT
	Cleanup(func())
}) *OpenLongChecker {
	mock := &OpenLongChecker{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
