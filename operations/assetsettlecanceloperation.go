package operations

import (
	"github.com/denkhaus/bitshares/types"
	"github.com/denkhaus/bitshares/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeAssetSettleCancel] = func() types.Operation {
		return &AssetSettleCancelOperation{}
	}
}

// AssetSettleCancelOperation (op 42) is a virtual operation generated when a forced
// settlement is cancelled before execution.
type AssetSettleCancelOperation struct {
	types.OperationFee
	Settlement types.ForceSettlementID `json:"settlement"`
	Account    types.AccountID         `json:"account"`
	Amount     types.AssetAmount       `json:"amount"`
}

func (p AssetSettleCancelOperation) Type() types.OperationType {
	return types.OperationTypeAssetSettleCancel
}

func (p AssetSettleCancelOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode type")
	}
	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}
	if err := enc.Encode(p.Settlement); err != nil {
		return errors.Annotate(err, "encode settlement")
	}
	if err := enc.Encode(p.Account); err != nil {
		return errors.Annotate(err, "encode account")
	}
	if err := enc.Encode(p.Amount); err != nil {
		return errors.Annotate(err, "encode amount")
	}
	return nil
}
