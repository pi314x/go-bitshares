package operations

import (
	"encoding/json"

	"github.com/denkhaus/bitshares/types"
	"github.com/denkhaus/bitshares/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeCustomAuthorityCreate] = func() types.Operation {
		return &CustomAuthorityCreateOperation{}
	}
	types.OperationMap[types.OperationTypeCustomAuthorityUpdate] = func() types.Operation {
		return &CustomAuthorityUpdateOperation{}
	}
	types.OperationMap[types.OperationTypeCustomAuthorityDelete] = func() types.Operation {
		return &CustomAuthorityDeleteOperation{}
	}
}

// Restriction is a complex recursive type used by custom authority operations.
// We represent it as raw JSON for parsing; binary encoding is not supported for
// building these operations from Go.
type Restriction struct {
	data json.RawMessage
}

func (p Restriction) MarshalJSON() ([]byte, error)  { return p.data, nil }
func (p *Restriction) UnmarshalJSON(b []byte) error { p.data = b; return nil }
func (p Restriction) Marshal(enc *util.TypeEncoder) error {
	return errors.New("Restriction binary encoding not implemented")
}

// CustomAuthorityCreateOperation (op 54) grants another account limited authority
// to perform specific operations on behalf of the granting account.
type CustomAuthorityCreateOperation struct {
	types.OperationFee
	Account       types.AccountID   `json:"account"`
	Enabled       bool              `json:"enabled"`
	ValidFrom     types.Time        `json:"valid_from"`
	ValidTo       types.Time        `json:"valid_to"`
	OperationType types.UInt32      `json:"operation_type"`
	Auth          types.Authority   `json:"auth"`
	Restrictions  []Restriction     `json:"restrictions"`
	Extensions    types.Extensions  `json:"extensions"`
}

func (p CustomAuthorityCreateOperation) Type() types.OperationType {
	return types.OperationTypeCustomAuthorityCreate
}

func (p CustomAuthorityCreateOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode type")
	}
	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}
	return errors.New("CustomAuthorityCreateOperation binary encoding requires restriction serialization — not supported")
}

// CustomAuthorityUpdateOperation (op 55) modifies an existing custom authority.
type CustomAuthorityUpdateOperation struct {
	types.OperationFee
	Account             types.AccountID        `json:"account"`
	AuthorityToUpdate   types.CustomAuthorityID `json:"authority_to_update"`
	NewEnabled          *bool                  `json:"new_enabled,omitempty"`
	NewValidFrom        *types.Time            `json:"new_valid_from,omitempty"`
	NewValidTo          *types.Time            `json:"new_valid_to,omitempty"`
	NewAuth             *types.Authority       `json:"new_auth,omitempty"`
	RestrictionsToAdd   []Restriction          `json:"restrictions_to_add"`
	Extensions          types.Extensions       `json:"extensions"`
}

func (p CustomAuthorityUpdateOperation) Type() types.OperationType {
	return types.OperationTypeCustomAuthorityUpdate
}

func (p CustomAuthorityUpdateOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode type")
	}
	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}
	return errors.New("CustomAuthorityUpdateOperation binary encoding requires restriction serialization — not supported")
}

// CustomAuthorityDeleteOperation (op 56) revokes a custom authority.
type CustomAuthorityDeleteOperation struct {
	types.OperationFee
	Account           types.AccountID        `json:"account"`
	AuthorityToDelete types.CustomAuthorityID `json:"authority_to_delete"`
	Extensions        types.Extensions       `json:"extensions"`
}

func (p CustomAuthorityDeleteOperation) Type() types.OperationType {
	return types.OperationTypeCustomAuthorityDelete
}

func (p CustomAuthorityDeleteOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode type")
	}
	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}
	if err := enc.Encode(p.Account); err != nil {
		return errors.Annotate(err, "encode account")
	}
	if err := enc.Encode(p.AuthorityToDelete); err != nil {
		return errors.Annotate(err, "encode authority_to_delete")
	}
	if err := enc.Encode(p.Extensions); err != nil {
		return errors.Annotate(err, "encode extensions")
	}
	return nil
}
