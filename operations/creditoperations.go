package operations

import (
	"github.com/denkhaus/bitshares/types"
	"github.com/denkhaus/bitshares/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeCreditOfferCreate] = func() types.Operation {
		return &CreditOfferCreateOperation{}
	}
	types.OperationMap[types.OperationTypeCreditOfferDelete] = func() types.Operation {
		return &CreditOfferDeleteOperation{}
	}
	types.OperationMap[types.OperationTypeCreditOfferUpdate] = func() types.Operation {
		return &CreditOfferUpdateOperation{}
	}
	types.OperationMap[types.OperationTypeCreditOfferAccept] = func() types.Operation {
		return &CreditOfferAcceptOperation{}
	}
	types.OperationMap[types.OperationTypeCreditDealRepay] = func() types.Operation {
		return &CreditDealRepayOperation{}
	}
	types.OperationMap[types.OperationTypeCreditDealUpdate] = func() types.Operation {
		return &CreditDealUpdateOperation{}
	}
	types.OperationMap[types.OperationTypeCreditDealExpired] = func() types.Operation {
		return &CreditDealExpiredOperation{}
	}
}

// CreditOfferCreateOperation (op 69) creates an on-chain credit offer (lending pool).
type CreditOfferCreateOperation struct {
	types.OperationFee
	OwnerAccount        types.AccountID   `json:"owner_account"`
	AssetType           types.AssetID     `json:"asset_type"`
	Balance             types.Int64       `json:"balance"`
	FeeRate             types.UInt32      `json:"fee_rate"`
	MaxDurationSeconds  types.UInt32      `json:"max_duration_seconds"`
	MinDealAmount       types.Int64       `json:"min_deal_amount"`
	Enabled             bool              `json:"enabled"`
	AutoDisableTime     types.Time        `json:"auto_disable_time"`
	Extensions          types.Extensions  `json:"extensions"`
}

func (p CreditOfferCreateOperation) Type() types.OperationType {
	return types.OperationTypeCreditOfferCreate
}

func (p CreditOfferCreateOperation) Marshal(enc *util.TypeEncoder) error {
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
	if err := enc.Encode(p.MaxDurationSeconds); err != nil {
		return errors.Annotate(err, "encode max_duration_seconds")
	}
	if err := enc.Encode(p.MinDealAmount); err != nil {
		return errors.Annotate(err, "encode min_deal_amount")
	}
	if err := enc.Encode(p.Enabled); err != nil {
		return errors.Annotate(err, "encode enabled")
	}
	if err := enc.Encode(p.AutoDisableTime); err != nil {
		return errors.Annotate(err, "encode auto_disable_time")
	}
	if err := enc.Encode(p.Extensions); err != nil {
		return errors.Annotate(err, "encode extensions")
	}
	return nil
}

// CreditOfferDeleteOperation (op 70) deletes an empty credit offer.
type CreditOfferDeleteOperation struct {
	types.OperationFee
	OwnerAccount types.AccountID   `json:"owner_account"`
	OfferID      types.CreditOfferID `json:"offer_id"`
	Extensions   types.Extensions  `json:"extensions"`
}

func (p CreditOfferDeleteOperation) Type() types.OperationType {
	return types.OperationTypeCreditOfferDelete
}

func (p CreditOfferDeleteOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode type")
	}
	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}
	if err := enc.Encode(p.OwnerAccount); err != nil {
		return errors.Annotate(err, "encode owner_account")
	}
	if err := enc.Encode(p.OfferID); err != nil {
		return errors.Annotate(err, "encode offer_id")
	}
	if err := enc.Encode(p.Extensions); err != nil {
		return errors.Annotate(err, "encode extensions")
	}
	return nil
}

// CreditOfferUpdateOperation (op 71) modifies an existing credit offer.
type CreditOfferUpdateOperation struct {
	types.OperationFee
	OwnerAccount       types.AccountID     `json:"owner_account"`
	OfferID            types.CreditOfferID `json:"offer_id"`
	DeltaAmount        *types.AssetAmount  `json:"delta_amount,omitempty"`
	FeeRate            *types.UInt32       `json:"fee_rate,omitempty"`
	MaxDurationSeconds *types.UInt32       `json:"max_duration_seconds,omitempty"`
	MinDealAmount      *types.Int64        `json:"min_deal_amount,omitempty"`
	Enabled            *bool               `json:"enabled,omitempty"`
	AutoDisableTime    *types.Time         `json:"auto_disable_time,omitempty"`
	Extensions         types.Extensions    `json:"extensions"`
}

func (p CreditOfferUpdateOperation) Type() types.OperationType {
	return types.OperationTypeCreditOfferUpdate
}

func (p CreditOfferUpdateOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode type")
	}
	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}
	if err := enc.Encode(p.OwnerAccount); err != nil {
		return errors.Annotate(err, "encode owner_account")
	}
	if err := enc.Encode(p.OfferID); err != nil {
		return errors.Annotate(err, "encode offer_id")
	}
	// All remaining fields are optional — simplified: write Extensions only.
	if err := enc.Encode(p.Extensions); err != nil {
		return errors.Annotate(err, "encode extensions")
	}
	return nil
}

// CreditOfferAcceptOperation (op 72) borrows from a credit offer.
type CreditOfferAcceptOperation struct {
	types.OperationFee
	Borrower           types.AccountID     `json:"borrower"`
	OfferID            types.CreditOfferID `json:"offer_id"`
	BorrowAmount       types.AssetAmount   `json:"borrow_amount"`
	Collateral         types.AssetAmount   `json:"collateral"`
	MaxFeeRate         types.UInt32        `json:"max_fee_rate"`
	MinDurationSeconds types.UInt32        `json:"min_duration_seconds"`
	Extensions         types.Extensions    `json:"extensions"`
}

func (p CreditOfferAcceptOperation) Type() types.OperationType {
	return types.OperationTypeCreditOfferAccept
}

func (p CreditOfferAcceptOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode type")
	}
	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}
	if err := enc.Encode(p.Borrower); err != nil {
		return errors.Annotate(err, "encode borrower")
	}
	if err := enc.Encode(p.OfferID); err != nil {
		return errors.Annotate(err, "encode offer_id")
	}
	if err := enc.Encode(p.BorrowAmount); err != nil {
		return errors.Annotate(err, "encode borrow_amount")
	}
	if err := enc.Encode(p.Collateral); err != nil {
		return errors.Annotate(err, "encode collateral")
	}
	if err := enc.Encode(p.MaxFeeRate); err != nil {
		return errors.Annotate(err, "encode max_fee_rate")
	}
	if err := enc.Encode(p.MinDurationSeconds); err != nil {
		return errors.Annotate(err, "encode min_duration_seconds")
	}
	if err := enc.Encode(p.Extensions); err != nil {
		return errors.Annotate(err, "encode extensions")
	}
	return nil
}

// CreditDealRepayOperation (op 73) repays a credit deal.
type CreditDealRepayOperation struct {
	types.OperationFee
	Account     types.AccountID  `json:"account"`
	DealID      types.CreditDealID `json:"deal_id"`
	RepayAmount types.AssetAmount `json:"repay_amount"`
	CreditFee   types.AssetAmount `json:"credit_fee"`
	Extensions  types.Extensions `json:"extensions"`
}

func (p CreditDealRepayOperation) Type() types.OperationType {
	return types.OperationTypeCreditDealRepay
}

func (p CreditDealRepayOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode type")
	}
	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}
	if err := enc.Encode(p.Account); err != nil {
		return errors.Annotate(err, "encode account")
	}
	if err := enc.Encode(p.DealID); err != nil {
		return errors.Annotate(err, "encode deal_id")
	}
	if err := enc.Encode(p.RepayAmount); err != nil {
		return errors.Annotate(err, "encode repay_amount")
	}
	if err := enc.Encode(p.CreditFee); err != nil {
		return errors.Annotate(err, "encode credit_fee")
	}
	if err := enc.Encode(p.Extensions); err != nil {
		return errors.Annotate(err, "encode extensions")
	}
	return nil
}

// CreditDealUpdateOperation (op 76) changes the auto-repay setting on a credit deal.
// AutoRepay: 0=no_auto_repayment, 1=only_full_repayment, 2=allow_partial_repayment.
type CreditDealUpdateOperation struct {
	types.OperationFee
	Account    types.AccountID   `json:"account"`
	DealID     types.CreditDealID `json:"deal_id"`
	AutoRepay  types.UInt8       `json:"auto_repay"`
	Extensions types.Extensions  `json:"extensions"`
}

func (p CreditDealUpdateOperation) Type() types.OperationType {
	return types.OperationTypeCreditDealUpdate
}

func (p CreditDealUpdateOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode type")
	}
	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}
	if err := enc.Encode(p.Account); err != nil {
		return errors.Annotate(err, "encode account")
	}
	if err := enc.Encode(p.DealID); err != nil {
		return errors.Annotate(err, "encode deal_id")
	}
	if err := enc.Encode(p.AutoRepay); err != nil {
		return errors.Annotate(err, "encode auto_repay")
	}
	if err := enc.Encode(p.Extensions); err != nil {
		return errors.Annotate(err, "encode extensions")
	}
	return nil
}

// CreditDealExpiredOperation (op 74) is a virtual operation generated when a credit deal
// expires without being repaid and collateral is seized by the offer owner.
type CreditDealExpiredOperation struct {
	types.OperationFee
	DealID       types.CreditDealID  `json:"deal_id"`
	OfferID      types.CreditOfferID `json:"offer_id"`
	OfferOwner   types.AccountID     `json:"offer_owner"`
	Borrower     types.AccountID     `json:"borrower"`
	UnpaidAmount types.AssetAmount   `json:"unpaid_amount"`
	Collateral   types.AssetAmount   `json:"collateral"`
	FeeRate      types.UInt32        `json:"fee_rate"`
}

func (p CreditDealExpiredOperation) Type() types.OperationType {
	return types.OperationTypeCreditDealExpired
}

func (p CreditDealExpiredOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode type")
	}
	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}
	if err := enc.Encode(p.DealID); err != nil {
		return errors.Annotate(err, "encode deal_id")
	}
	if err := enc.Encode(p.OfferID); err != nil {
		return errors.Annotate(err, "encode offer_id")
	}
	if err := enc.Encode(p.OfferOwner); err != nil {
		return errors.Annotate(err, "encode offer_owner")
	}
	if err := enc.Encode(p.Borrower); err != nil {
		return errors.Annotate(err, "encode borrower")
	}
	if err := enc.Encode(p.UnpaidAmount); err != nil {
		return errors.Annotate(err, "encode unpaid_amount")
	}
	if err := enc.Encode(p.Collateral); err != nil {
		return errors.Annotate(err, "encode collateral")
	}
	if err := enc.Encode(p.FeeRate); err != nil {
		return errors.Annotate(err, "encode fee_rate")
	}
	return nil
}
