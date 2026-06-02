package operations

import (
	"github.com/denkhaus/bitshares/types"
	"github.com/denkhaus/bitshares/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeAssetClaimPool] = func() types.Operation {
		return &AssetClaimPoolOperation{}
	}
}

// AssetClaimPoolOperation (op 47) claims accumulated fees from an asset's fee pool.
type AssetClaimPoolOperation struct {
	types.OperationFee
	Issuer        types.AccountID  `json:"issuer"`
	AssetID       types.AssetID    `json:"asset_id"`
	AmountToClaim types.AssetAmount `json:"amount_to_claim"`
	Extensions    types.Extensions `json:"extensions"`
}

func (p AssetClaimPoolOperation) Type() types.OperationType {
	return types.OperationTypeAssetClaimPool
}

func (p AssetClaimPoolOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode type")
	}
	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}
	if err := enc.Encode(p.Issuer); err != nil {
		return errors.Annotate(err, "encode issuer")
	}
	if err := enc.Encode(p.AssetID); err != nil {
		return errors.Annotate(err, "encode asset_id")
	}
	if err := enc.Encode(p.AmountToClaim); err != nil {
		return errors.Annotate(err, "encode amount_to_claim")
	}
	if err := enc.Encode(p.Extensions); err != nil {
		return errors.Annotate(err, "encode extensions")
	}
	return nil
}
