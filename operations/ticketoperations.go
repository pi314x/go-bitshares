package operations

import (
	"github.com/denkhaus/bitshares/types"
	"github.com/denkhaus/bitshares/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeTicketCreate] = func() types.Operation {
		return &TicketCreateOperation{}
	}
	types.OperationMap[types.OperationTypeTicketUpdate] = func() types.Operation {
		return &TicketUpdateOperation{}
	}
}

// TicketCreateOperation (op 57) locks an amount into a governance participation ticket.
// TargetType: 0=liquid, 1=lock_180_days, 2=lock_360_days, 3=lock_720_days, 4=lock_forever.
type TicketCreateOperation struct {
	types.OperationFee
	Account    types.AccountID   `json:"account"`
	TargetType types.UInt32      `json:"target_type"`
	Amount     types.AssetAmount `json:"amount"`
	Extensions types.Extensions  `json:"extensions"`
}

func (p TicketCreateOperation) Type() types.OperationType {
	return types.OperationTypeTicketCreate
}

func (p TicketCreateOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode type")
	}
	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}
	if err := enc.Encode(p.Account); err != nil {
		return errors.Annotate(err, "encode account")
	}
	if err := enc.Encode(p.TargetType); err != nil {
		return errors.Annotate(err, "encode target_type")
	}
	if err := enc.Encode(p.Amount); err != nil {
		return errors.Annotate(err, "encode amount")
	}
	if err := enc.Encode(p.Extensions); err != nil {
		return errors.Annotate(err, "encode extensions")
	}
	return nil
}

// TicketUpdateOperation (op 58) changes the lock type of an existing ticket.
type TicketUpdateOperation struct {
	types.OperationFee
	Ticket             types.TicketID    `json:"ticket"`
	Account            types.AccountID   `json:"account"`
	TargetType         types.UInt32      `json:"target_type"`
	AmountForNewTarget *types.AssetAmount `json:"amount_for_new_target,omitempty"`
	Extensions         types.Extensions  `json:"extensions"`
}

func (p TicketUpdateOperation) Type() types.OperationType {
	return types.OperationTypeTicketUpdate
}

func (p TicketUpdateOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode type")
	}
	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}
	if err := enc.Encode(p.Ticket); err != nil {
		return errors.Annotate(err, "encode ticket")
	}
	if err := enc.Encode(p.Account); err != nil {
		return errors.Annotate(err, "encode account")
	}
	if err := enc.Encode(p.TargetType); err != nil {
		return errors.Annotate(err, "encode target_type")
	}
	// optional amount_for_new_target
	if p.AmountForNewTarget != nil {
		if err := enc.Encode(true); err != nil {
			return errors.Annotate(err, "encode amount_for_new_target present flag")
		}
		if err := enc.Encode(*p.AmountForNewTarget); err != nil {
			return errors.Annotate(err, "encode amount_for_new_target")
		}
	} else {
		if err := enc.Encode(false); err != nil {
			return errors.Annotate(err, "encode amount_for_new_target absent flag")
		}
	}
	if err := enc.Encode(p.Extensions); err != nil {
		return errors.Annotate(err, "encode extensions")
	}
	return nil
}
