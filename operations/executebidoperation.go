package operations

import (
	"github.com/denkhaus/bitshares/types"
	"github.com/denkhaus/bitshares/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeExecuteBid] = func() types.Operation {
		return &ExecuteBidOperation{}
	}
}

// ExecuteBidOperation (op 46) is a virtual operation generated when a collateral
// bid is executed during a global settlement revocation.
type ExecuteBidOperation struct {
	types.OperationFee
	Bidder     types.AccountID   `json:"bidder"`
	Debt       types.AssetAmount `json:"debt"`
	Collateral types.AssetAmount `json:"collateral"`
}

func (p ExecuteBidOperation) Type() types.OperationType {
	return types.OperationTypeExecuteBid
}

func (p ExecuteBidOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode type")
	}
	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}
	if err := enc.Encode(p.Bidder); err != nil {
		return errors.Annotate(err, "encode bidder")
	}
	if err := enc.Encode(p.Debt); err != nil {
		return errors.Annotate(err, "encode debt")
	}
	if err := enc.Encode(p.Collateral); err != nil {
		return errors.Annotate(err, "encode collateral")
	}
	return nil
}
