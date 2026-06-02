package operations

import (
	"github.com/denkhaus/bitshares/types"
	"github.com/denkhaus/bitshares/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeFBADistribute] = func() types.Operation {
		return &FBADistributeOperation{}
	}
}

// FBADistributeOperation (op 44) is a virtual operation that distributes fees
// collected by a Fee-Backed Asset (FBA) to its asset holders.
type FBADistributeOperation struct {
	types.OperationFee
	AccountID types.AccountID `json:"account_id"`
	FBAID     types.ObjectID  `json:"fba_id"`
	Amount    types.Int64     `json:"amount"`
}

func (p FBADistributeOperation) Type() types.OperationType {
	return types.OperationTypeFBADistribute
}

func (p FBADistributeOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode type")
	}
	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}
	if err := enc.Encode(p.AccountID); err != nil {
		return errors.Annotate(err, "encode account_id")
	}
	if err := enc.Encode(p.FBAID); err != nil {
		return errors.Annotate(err, "encode fba_id")
	}
	if err := enc.Encode(p.Amount); err != nil {
		return errors.Annotate(err, "encode amount")
	}
	return nil
}
