package operations

import (
	"github.com/denkhaus/bitshares/types"
	"github.com/denkhaus/bitshares/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeLiquidityPoolCreate] = func() types.Operation {
		return &LiquidityPoolCreateOperation{}
	}
	types.OperationMap[types.OperationTypeLiquidityPoolDelete] = func() types.Operation {
		return &LiquidityPoolDeleteOperation{}
	}
	types.OperationMap[types.OperationTypeLiquidityPoolDeposit] = func() types.Operation {
		return &LiquidityPoolDepositOperation{}
	}
	types.OperationMap[types.OperationTypeLiquidityPoolWithdraw] = func() types.Operation {
		return &LiquidityPoolWithdrawOperation{}
	}
	types.OperationMap[types.OperationTypeLiquidityPoolExchange] = func() types.Operation {
		return &LiquidityPoolExchangeOperation{}
	}
	types.OperationMap[types.OperationTypeLiquidityPoolUpdate] = func() types.Operation {
		return &LiquidityPoolUpdateOperation{}
	}
}

// LiquidityPoolCreateOperation (op 59) creates a new AMM liquidity pool.
type LiquidityPoolCreateOperation struct {
	types.OperationFee
	Account            types.AccountID      `json:"account"`
	AssetA             types.AssetID        `json:"asset_a"`
	AssetB             types.AssetID        `json:"asset_b"`
	ShareAsset         types.AssetID        `json:"share_asset"`
	TakerFeePercent    types.UInt16         `json:"taker_fee_percent"`
	WithdrawalFeePercent types.UInt16       `json:"withdrawal_fee_percent"`
	Extensions         types.Extensions     `json:"extensions"`
}

func (p LiquidityPoolCreateOperation) Type() types.OperationType {
	return types.OperationTypeLiquidityPoolCreate
}

func (p LiquidityPoolCreateOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode type")
	}
	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}
	if err := enc.Encode(p.Account); err != nil {
		return errors.Annotate(err, "encode account")
	}
	if err := enc.Encode(p.AssetA); err != nil {
		return errors.Annotate(err, "encode asset_a")
	}
	if err := enc.Encode(p.AssetB); err != nil {
		return errors.Annotate(err, "encode asset_b")
	}
	if err := enc.Encode(p.ShareAsset); err != nil {
		return errors.Annotate(err, "encode share_asset")
	}
	if err := enc.Encode(p.TakerFeePercent); err != nil {
		return errors.Annotate(err, "encode taker_fee_percent")
	}
	if err := enc.Encode(p.WithdrawalFeePercent); err != nil {
		return errors.Annotate(err, "encode withdrawal_fee_percent")
	}
	if err := enc.Encode(p.Extensions); err != nil {
		return errors.Annotate(err, "encode extensions")
	}
	return nil
}

// LiquidityPoolDeleteOperation (op 60) deletes an empty liquidity pool.
type LiquidityPoolDeleteOperation struct {
	types.OperationFee
	Account    types.AccountID      `json:"account"`
	Pool       types.LiquidityPoolID `json:"pool"`
	Extensions types.Extensions     `json:"extensions"`
}

func (p LiquidityPoolDeleteOperation) Type() types.OperationType {
	return types.OperationTypeLiquidityPoolDelete
}

func (p LiquidityPoolDeleteOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode type")
	}
	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}
	if err := enc.Encode(p.Account); err != nil {
		return errors.Annotate(err, "encode account")
	}
	if err := enc.Encode(p.Pool); err != nil {
		return errors.Annotate(err, "encode pool")
	}
	if err := enc.Encode(p.Extensions); err != nil {
		return errors.Annotate(err, "encode extensions")
	}
	return nil
}

// LiquidityPoolDepositOperation (op 61) adds liquidity to a pool.
type LiquidityPoolDepositOperation struct {
	types.OperationFee
	Account    types.AccountID      `json:"account"`
	Pool       types.LiquidityPoolID `json:"pool"`
	AmountA    types.AssetAmount    `json:"amount_a"`
	AmountB    types.AssetAmount    `json:"amount_b"`
	Extensions types.Extensions     `json:"extensions"`
}

func (p LiquidityPoolDepositOperation) Type() types.OperationType {
	return types.OperationTypeLiquidityPoolDeposit
}

func (p LiquidityPoolDepositOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode type")
	}
	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}
	if err := enc.Encode(p.Account); err != nil {
		return errors.Annotate(err, "encode account")
	}
	if err := enc.Encode(p.Pool); err != nil {
		return errors.Annotate(err, "encode pool")
	}
	if err := enc.Encode(p.AmountA); err != nil {
		return errors.Annotate(err, "encode amount_a")
	}
	if err := enc.Encode(p.AmountB); err != nil {
		return errors.Annotate(err, "encode amount_b")
	}
	if err := enc.Encode(p.Extensions); err != nil {
		return errors.Annotate(err, "encode extensions")
	}
	return nil
}

// LiquidityPoolWithdrawOperation (op 62) removes liquidity from a pool by burning share tokens.
type LiquidityPoolWithdrawOperation struct {
	types.OperationFee
	Account     types.AccountID      `json:"account"`
	Pool        types.LiquidityPoolID `json:"pool"`
	ShareAmount types.AssetAmount    `json:"share_amount"`
	Extensions  types.Extensions     `json:"extensions"`
}

func (p LiquidityPoolWithdrawOperation) Type() types.OperationType {
	return types.OperationTypeLiquidityPoolWithdraw
}

func (p LiquidityPoolWithdrawOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode type")
	}
	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}
	if err := enc.Encode(p.Account); err != nil {
		return errors.Annotate(err, "encode account")
	}
	if err := enc.Encode(p.Pool); err != nil {
		return errors.Annotate(err, "encode pool")
	}
	if err := enc.Encode(p.ShareAmount); err != nil {
		return errors.Annotate(err, "encode share_amount")
	}
	if err := enc.Encode(p.Extensions); err != nil {
		return errors.Annotate(err, "encode extensions")
	}
	return nil
}

// LiquidityPoolExchangeOperation (op 63) swaps an asset into a pool.
// This is the primary operation used by the arbitrage bot.
type LiquidityPoolExchangeOperation struct {
	types.OperationFee
	Account      types.AccountID      `json:"account"`
	Pool         types.LiquidityPoolID `json:"pool"`
	AmountToSell types.AssetAmount    `json:"amount_to_sell"`
	MinToReceive types.AssetAmount    `json:"min_to_receive"`
	Extensions   types.Extensions     `json:"extensions"`
}

func (p LiquidityPoolExchangeOperation) Type() types.OperationType {
	return types.OperationTypeLiquidityPoolExchange
}

func (p LiquidityPoolExchangeOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode type")
	}
	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}
	if err := enc.Encode(p.Account); err != nil {
		return errors.Annotate(err, "encode account")
	}
	if err := enc.Encode(p.Pool); err != nil {
		return errors.Annotate(err, "encode pool")
	}
	if err := enc.Encode(p.AmountToSell); err != nil {
		return errors.Annotate(err, "encode amount_to_sell")
	}
	if err := enc.Encode(p.MinToReceive); err != nil {
		return errors.Annotate(err, "encode min_to_receive")
	}
	if err := enc.Encode(p.Extensions); err != nil {
		return errors.Annotate(err, "encode extensions")
	}
	return nil
}

// LiquidityPoolUpdateOperation (op 75) changes fee percentages on an existing pool.
type LiquidityPoolUpdateOperation struct {
	types.OperationFee
	Account              types.AccountID       `json:"account"`
	Pool                 types.LiquidityPoolID  `json:"pool"`
	NewTakerFeePercent   *types.UInt16         `json:"taker_fee_percent,omitempty"`
	NewWithdrawalFee     *types.UInt16         `json:"withdrawal_fee_percent,omitempty"`
	Extensions           types.Extensions      `json:"extensions"`
}

func (p LiquidityPoolUpdateOperation) Type() types.OperationType {
	return types.OperationTypeLiquidityPoolUpdate
}

func (p LiquidityPoolUpdateOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode type")
	}
	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}
	if err := enc.Encode(p.Account); err != nil {
		return errors.Annotate(err, "encode account")
	}
	if err := enc.Encode(p.Pool); err != nil {
		return errors.Annotate(err, "encode pool")
	}
	// optional taker_fee_percent
	if p.NewTakerFeePercent != nil {
		if err := enc.Encode(true); err != nil {
			return errors.Annotate(err, "encode taker_fee_percent present")
		}
		if err := enc.Encode(*p.NewTakerFeePercent); err != nil {
			return errors.Annotate(err, "encode taker_fee_percent")
		}
	} else {
		if err := enc.Encode(false); err != nil {
			return errors.Annotate(err, "encode taker_fee_percent absent")
		}
	}
	// optional withdrawal_fee_percent
	if p.NewWithdrawalFee != nil {
		if err := enc.Encode(true); err != nil {
			return errors.Annotate(err, "encode withdrawal_fee_percent present")
		}
		if err := enc.Encode(*p.NewWithdrawalFee); err != nil {
			return errors.Annotate(err, "encode withdrawal_fee_percent")
		}
	} else {
		if err := enc.Encode(false); err != nil {
			return errors.Annotate(err, "encode withdrawal_fee_percent absent")
		}
	}
	if err := enc.Encode(p.Extensions); err != nil {
		return errors.Annotate(err, "encode extensions")
	}
	return nil
}
