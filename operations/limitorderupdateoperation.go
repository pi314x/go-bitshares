package operations

import (
	"github.com/denkhaus/bitshares/types"
	"github.com/denkhaus/bitshares/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeLimitOrderUpdate] = func() types.Operation {
		return &LimitOrderUpdateOperation{}
	}
}

// LimitOrderUpdateOperation (op 77) modifies price, amount, or expiration of an open limit order.
type LimitOrderUpdateOperation struct {
	types.OperationFee
	Seller             types.AccountID    `json:"seller"`
	Order              types.LimitOrderID `json:"order"`
	NewPrice           *types.Price       `json:"new_price,omitempty"`
	DeltaAmountToSell  *types.AssetAmount `json:"delta_amount_to_sell,omitempty"`
	NewExpiration      *types.Time        `json:"new_expiration,omitempty"`
	Extensions         types.Extensions   `json:"extensions"`
}

func (p LimitOrderUpdateOperation) Type() types.OperationType {
	return types.OperationTypeLimitOrderUpdate
}

func (p LimitOrderUpdateOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode type")
	}
	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}
	if err := enc.Encode(p.Seller); err != nil {
		return errors.Annotate(err, "encode seller")
	}
	if err := enc.Encode(p.Order); err != nil {
		return errors.Annotate(err, "encode order")
	}
	// optional new_price
	if p.NewPrice != nil {
		if err := enc.Encode(true); err != nil {
			return errors.Annotate(err, "encode new_price present")
		}
		if err := enc.Encode(*p.NewPrice); err != nil {
			return errors.Annotate(err, "encode new_price")
		}
	} else {
		if err := enc.Encode(false); err != nil {
			return errors.Annotate(err, "encode new_price absent")
		}
	}
	// optional delta_amount_to_sell
	if p.DeltaAmountToSell != nil {
		if err := enc.Encode(true); err != nil {
			return errors.Annotate(err, "encode delta_amount_to_sell present")
		}
		if err := enc.Encode(*p.DeltaAmountToSell); err != nil {
			return errors.Annotate(err, "encode delta_amount_to_sell")
		}
	} else {
		if err := enc.Encode(false); err != nil {
			return errors.Annotate(err, "encode delta_amount_to_sell absent")
		}
	}
	// optional new_expiration
	if p.NewExpiration != nil {
		if err := enc.Encode(true); err != nil {
			return errors.Annotate(err, "encode new_expiration present")
		}
		if err := enc.Encode(*p.NewExpiration); err != nil {
			return errors.Annotate(err, "encode new_expiration")
		}
	} else {
		if err := enc.Encode(false); err != nil {
			return errors.Annotate(err, "encode new_expiration absent")
		}
	}
	if err := enc.Encode(p.Extensions); err != nil {
		return errors.Annotate(err, "encode extensions")
	}
	return nil
}
