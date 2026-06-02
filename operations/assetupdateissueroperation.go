package operations

import (
	"github.com/denkhaus/bitshares/types"
	"github.com/denkhaus/bitshares/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeAssetUpdateIssuer] = func() types.Operation {
		return &AssetUpdateIssuerOperation{}
	}
}

// AssetUpdateIssuerOperation (op 48) transfers an asset's issuer role to a new account.
type AssetUpdateIssuerOperation struct {
	types.OperationFee
	Issuer        types.AccountID  `json:"issuer"`
	AssetToUpdate types.AssetID    `json:"asset_to_update"`
	NewIssuer     types.AccountID  `json:"new_issuer"`
	Extensions    types.Extensions `json:"extensions"`
}

func (p AssetUpdateIssuerOperation) Type() types.OperationType {
	return types.OperationTypeAssetUpdateIssuer
}

func (p AssetUpdateIssuerOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode type")
	}
	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}
	if err := enc.Encode(p.Issuer); err != nil {
		return errors.Annotate(err, "encode issuer")
	}
	if err := enc.Encode(p.AssetToUpdate); err != nil {
		return errors.Annotate(err, "encode asset_to_update")
	}
	if err := enc.Encode(p.NewIssuer); err != nil {
		return errors.Annotate(err, "encode new_issuer")
	}
	if err := enc.Encode(p.Extensions); err != nil {
		return errors.Annotate(err, "encode extensions")
	}
	return nil
}
