package operations

import (
	"github.com/denkhaus/bitshares/types"
	"github.com/denkhaus/bitshares/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeSametFundCreate] = func() types.Operation {
		return &SametFundCreateOperation{}
	}
	types.OperationMap[types.OperationTypeSametFundDelete] = func() types.Operation {
		return &SametFundDeleteOperation{}
	}
	types.OperationMap[types.OperationTypeSametFundUpdate] = func() types.Operation {
		return &SametFundUpdateOperation{}
	}
	types.OperationMap[types.OperationTypeSametFundBorrow] = func() types.Operation {
		return &SametFundBorrowOperation{}
	}
	types.OperationMap[types.OperationTypeSametFundRepay] = func() types.Operation {
		return &SametFundRepayOperation{}
	}
}

// SametFundCreateOperation (op 64) creates a same-type asset lending fund (flash loans).
type SametFundCreateOperation struct {
	types.OperationFee
	OwnerAccount types.AccountID   `json:"owner_account"`
	AssetType    types.AssetID     `json:"asset_type"`
	Balance      types.Int64       `json:"balance"`
	FeeRate      types.UInt32      `json:"fee_rate"`
	Extensions   types.Extensions  `json:"extensions"`
}

func (p SametFundCreateOperation) Type() types.OperationType {
	return types.OperationTypeSametFundCreate
}

func (p SametFundCreateOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode type")
	}
	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}
	if err := enc.Encode(p.OwnerAccount); err != nil {
		return errors.Annotate(err, "encode owner_account")
	}
	if err := enc.Encode(p.AssetType); err != nil {
		return errors.Annotate(err, "encode asset_type")
	}
	if err := enc.Encode(p.Balance); err != nil {
		return errors.Annotate(err, "encode balance")
	}
	if err := enc.Encode(p.FeeRate); err != nil {
		return errors.Annotate(err, "encode fee_rate")
	}
	if err := enc.Encode(p.Extensions); err != nil {
		return errors.Annotate(err, "encode extensions")
	}
	return nil
}

// SametFundDeleteOperation (op 65) deletes an empty SameT fund.
type SametFundDeleteOperation struct {
	types.OperationFee
	OwnerAccount types.AccountID   `json:"owner_account"`
	FundID       types.SametFundID `json:"fund_id"`
	Extensions   types.Extensions  `json:"extensions"`
}

func (p SametFundDeleteOperation) Type() types.OperationType {
	return types.OperationTypeSametFundDelete
}

func (p SametFundDeleteOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode type")
	}
	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}
	if err := enc.Encode(p.OwnerAccount); err != nil {
		return errors.Annotate(err, "encode owner_account")
	}
	if err := enc.Encode(p.FundID); err != nil {
		return errors.Annotate(err, "encode fund_id")
	}
	if err := enc.Encode(p.Extensions); err != nil {
		return errors.Annotate(err, "encode extensions")
	}
	return nil
}

// SametFundUpdateOperation (op 66) changes the balance or fee rate of a SameT fund.
type SametFundUpdateOperation struct {
	types.OperationFee
	OwnerAccount types.AccountID    `json:"owner_account"`
	FundID       types.SametFundID  `json:"fund_id"`
	DeltaAmount  *types.AssetAmount `json:"delta_amount,omitempty"`
	NewFeeRate   *types.UInt32      `json:"new_fee_rate,omitempty"`
	Extensions   types.Extensions   `json:"extensions"`
}

func (p SametFundUpdateOperation) Type() types.OperationType {
	return types.OperationTypeSametFundUpdate
}

func (p SametFundUpdateOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode type")
	}
	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}
	if err := enc.Encode(p.OwnerAccount); err != nil {
		return errors.Annotate(err, "encode owner_account")
	}
	if err := enc.Encode(p.FundID); err != nil {
		return errors.Annotate(err, "encode fund_id")
	}
	if p.DeltaAmount != nil {
		if err := enc.Encode(true); err != nil {
			return errors.Annotate(err, "encode delta_amount present")
		}
		if err := enc.Encode(*p.DeltaAmount); err != nil {
			return errors.Annotate(err, "encode delta_amount")
		}
	} else {
		if err := enc.Encode(false); err != nil {
			return errors.Annotate(err, "encode delta_amount absent")
		}
	}
	if p.NewFeeRate != nil {
		if err := enc.Encode(true); err != nil {
			return errors.Annotate(err, "encode new_fee_rate present")
		}
		if err := enc.Encode(*p.NewFeeRate); err != nil {
			return errors.Annotate(err, "encode new_fee_rate")
		}
	} else {
		if err := enc.Encode(false); err != nil {
			return errors.Annotate(err, "encode new_fee_rate absent")
		}
	}
	if err := enc.Encode(p.Extensions); err != nil {
		return errors.Annotate(err, "encode extensions")
	}
	return nil
}

// SametFundBorrowOperation (op 67) takes a flash loan from a SameT fund.
type SametFundBorrowOperation struct {
	types.OperationFee
	Borrower     types.AccountID   `json:"borrower"`
	FundID       types.SametFundID `json:"fund_id"`
	BorrowAmount types.AssetAmount `json:"borrow_amount"`
	Extensions   types.Extensions  `json:"extensions"`
}

func (p SametFundBorrowOperation) Type() types.OperationType {
	return types.OperationTypeSametFundBorrow
}

func (p SametFundBorrowOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode type")
	}
	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}
	if err := enc.Encode(p.Borrower); err != nil {
		return errors.Annotate(err, "encode borrower")
	}
	if err := enc.Encode(p.FundID); err != nil {
		return errors.Annotate(err, "encode fund_id")
	}
	if err := enc.Encode(p.BorrowAmount); err != nil {
		return errors.Annotate(err, "encode borrow_amount")
	}
	if err := enc.Encode(p.Extensions); err != nil {
		return errors.Annotate(err, "encode extensions")
	}
	return nil
}

// SametFundRepayOperation (op 68) repays a SameT fund flash loan plus fees.
type SametFundRepayOperation struct {
	types.OperationFee
	Account      types.AccountID   `json:"account"`
	FundID       types.SametFundID `json:"fund_id"`
	RepayAmount  types.AssetAmount `json:"repay_amount"`
	FundFee      types.AssetAmount `json:"fund_fee"`
	Extensions   types.Extensions  `json:"extensions"`
}

func (p SametFundRepayOperation) Type() types.OperationType {
	return types.OperationTypeSametFundRepay
}

func (p SametFundRepayOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode type")
	}
	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}
	if err := enc.Encode(p.Account); err != nil {
		return errors.Annotate(err, "encode account")
	}
	if err := enc.Encode(p.FundID); err != nil {
		return errors.Annotate(err, "encode fund_id")
	}
	if err := enc.Encode(p.RepayAmount); err != nil {
		return errors.Annotate(err, "encode repay_amount")
	}
	if err := enc.Encode(p.FundFee); err != nil {
		return errors.Annotate(err, "encode fund_fee")
	}
	if err := enc.Encode(p.Extensions); err != nil {
		return errors.Annotate(err, "encode extensions")
	}
	return nil
}
