package operations

import (
	"github.com/denkhaus/bitshares/types"
	"github.com/denkhaus/bitshares/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeBlindTransfer] = func() types.Operation {
		return &BlindTransferOperation{}
	}
}

// BlindTransferOperation (op 40) transfers assets between confidential blinded balances.
// Both inputs and outputs are Pedersen commitments; no account IDs are revealed.
type BlindTransferOperation struct {
	types.OperationFee
	Inputs  types.BlindInputs  `json:"inputs"`
	Outputs types.BlindOutputs `json:"outputs"`
}

func (p BlindTransferOperation) Type() types.OperationType {
	return types.OperationTypeBlindTransfer
}

func (p BlindTransferOperation) MarshalFeeScheduleParams(params types.M, enc *util.TypeEncoder) error {
	if fee, ok := params["fee"]; ok {
		if err := enc.Encode(types.UInt64(fee.(float64))); err != nil {
			return errors.Annotate(err, "encode Fee")
		}
	}
	if ppk, ok := params["price_per_output"]; ok {
		if err := enc.Encode(types.UInt32(ppk.(float64))); err != nil {
			return errors.Annotate(err, "encode PricePerOutput")
		}
	}
	return nil
}

func (p BlindTransferOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode type")
	}
	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}
	if err := enc.Encode(p.Inputs); err != nil {
		return errors.Annotate(err, "encode inputs")
	}
	if err := enc.Encode(p.Outputs); err != nil {
		return errors.Annotate(err, "encode outputs")
	}
	return nil
}
