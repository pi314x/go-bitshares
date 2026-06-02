package operations

import (
	"encoding/json"

	"github.com/denkhaus/bitshares/types"
	"github.com/denkhaus/bitshares/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeHTLCCreate] = func() types.Operation {
		return &HTLCCreateOperation{}
	}
	types.OperationMap[types.OperationTypeHTLCRedeem] = func() types.Operation {
		return &HTLCRedeemOperation{}
	}
	types.OperationMap[types.OperationTypeHTLCRedeemed] = func() types.Operation {
		return &HTLCRedeemedOperation{}
	}
	types.OperationMap[types.OperationTypeHTLCExtend] = func() types.Operation {
		return &HTLCExtendOperation{}
	}
	types.OperationMap[types.OperationTypeHTLCRefund] = func() types.Operation {
		return &HTLCRefundOperation{}
	}
}

// HTLCHash holds the preimage hash for an HTLC. BitShares supports sha256, ripemd160 and hash160.
// We represent the variant as raw JSON for flexibility.
type HTLCHash struct {
	data json.RawMessage
}

func (p HTLCHash) MarshalJSON() ([]byte, error) { return p.data, nil }
func (p *HTLCHash) UnmarshalJSON(b []byte) error { p.data = b; return nil }

// Marshal writes the hash as a raw byte buffer prefixed by its varint length.
// BitShares encodes the fc::static_variant as type-byte + value, which requires
// knowledge of the hash algorithm. We write a best-effort opaque encoding here.
func (p HTLCHash) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(len(p.data))); err != nil {
		return errors.Annotate(err, "encode hash length")
	}
	if err := enc.Encode([]byte(p.data)); err != nil {
		return errors.Annotate(err, "encode hash data")
	}
	return nil
}

// HTLCCreateOperation (op 49) creates a Hash Time Lock Contract.
type HTLCCreateOperation struct {
	types.OperationFee
	From                types.AccountID  `json:"from"`
	To                  types.AccountID  `json:"to"`
	Amount              types.AssetAmount `json:"amount"`
	PreimageHash        HTLCHash         `json:"preimage_hash"`
	PreimageSize        types.UInt16     `json:"preimage_size"`
	ClaimPeriodSeconds  types.UInt32     `json:"claim_period_seconds"`
	Extensions          types.Extensions `json:"extensions"`
}

func (p HTLCCreateOperation) Type() types.OperationType {
	return types.OperationTypeHTLCCreate
}

func (p HTLCCreateOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode type")
	}
	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}
	if err := enc.Encode(p.From); err != nil {
		return errors.Annotate(err, "encode from")
	}
	if err := enc.Encode(p.To); err != nil {
		return errors.Annotate(err, "encode to")
	}
	if err := enc.Encode(p.Amount); err != nil {
		return errors.Annotate(err, "encode amount")
	}
	if err := enc.Encode(p.PreimageHash); err != nil {
		return errors.Annotate(err, "encode preimage_hash")
	}
	if err := enc.Encode(p.PreimageSize); err != nil {
		return errors.Annotate(err, "encode preimage_size")
	}
	if err := enc.Encode(p.ClaimPeriodSeconds); err != nil {
		return errors.Annotate(err, "encode claim_period_seconds")
	}
	if err := enc.Encode(p.Extensions); err != nil {
		return errors.Annotate(err, "encode extensions")
	}
	return nil
}

// HTLCRedeemOperation (op 50) redeems an HTLC by providing the preimage.
type HTLCRedeemOperation struct {
	types.OperationFee
	HTLCId     types.HTLCID     `json:"htlc_id"`
	Redeemer   types.AccountID  `json:"redeemer"`
	Preimage   types.Buffer     `json:"preimage"`
	Extensions types.Extensions `json:"extensions"`
}

func (p HTLCRedeemOperation) Type() types.OperationType {
	return types.OperationTypeHTLCRedeem
}

func (p HTLCRedeemOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode type")
	}
	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}
	if err := enc.Encode(p.HTLCId); err != nil {
		return errors.Annotate(err, "encode htlc_id")
	}
	if err := enc.Encode(p.Redeemer); err != nil {
		return errors.Annotate(err, "encode redeemer")
	}
	if err := enc.Encode(p.Preimage); err != nil {
		return errors.Annotate(err, "encode preimage")
	}
	if err := enc.Encode(p.Extensions); err != nil {
		return errors.Annotate(err, "encode extensions")
	}
	return nil
}

// HTLCExtendOperation (op 52) extends the claim period of an HTLC.
type HTLCExtendOperation struct {
	types.OperationFee
	HTLCId       types.HTLCID     `json:"htlc_id"`
	UpdateIssuer types.AccountID  `json:"update_issuer"`
	SecondsToAdd types.UInt32     `json:"seconds_to_add"`
	Extensions   types.Extensions `json:"extensions"`
}

func (p HTLCExtendOperation) Type() types.OperationType {
	return types.OperationTypeHTLCExtend
}

func (p HTLCExtendOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode type")
	}
	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}
	if err := enc.Encode(p.HTLCId); err != nil {
		return errors.Annotate(err, "encode htlc_id")
	}
	if err := enc.Encode(p.UpdateIssuer); err != nil {
		return errors.Annotate(err, "encode update_issuer")
	}
	if err := enc.Encode(p.SecondsToAdd); err != nil {
		return errors.Annotate(err, "encode seconds_to_add")
	}
	if err := enc.Encode(p.Extensions); err != nil {
		return errors.Annotate(err, "encode extensions")
	}
	return nil
}

// HTLCRedeemedOperation (op 51) is a virtual operation generated when an HTLC
// is successfully redeemed. It records the revealed preimage for history.
type HTLCRedeemedOperation struct {
	types.OperationFee
	HTLCId        types.HTLCID      `json:"htlc_id"`
	From          types.AccountID   `json:"from"`
	To            types.AccountID   `json:"to"`
	Redeemer      types.AccountID   `json:"redeemer"`
	Amount        types.AssetAmount `json:"amount"`
	PreimageHash  HTLCHash          `json:"htlc_preimage_hash"`
	PreimageSize  types.UInt16      `json:"htlc_preimage_size"`
	Preimage      types.Buffer      `json:"preimage"`
}

func (p HTLCRedeemedOperation) Type() types.OperationType {
	return types.OperationTypeHTLCRedeemed
}

func (p HTLCRedeemedOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode type")
	}
	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}
	if err := enc.Encode(p.HTLCId); err != nil {
		return errors.Annotate(err, "encode htlc_id")
	}
	if err := enc.Encode(p.From); err != nil {
		return errors.Annotate(err, "encode from")
	}
	if err := enc.Encode(p.To); err != nil {
		return errors.Annotate(err, "encode to")
	}
	if err := enc.Encode(p.Redeemer); err != nil {
		return errors.Annotate(err, "encode redeemer")
	}
	if err := enc.Encode(p.Amount); err != nil {
		return errors.Annotate(err, "encode amount")
	}
	if err := enc.Encode(p.PreimageHash); err != nil {
		return errors.Annotate(err, "encode preimage_hash")
	}
	if err := enc.Encode(p.PreimageSize); err != nil {
		return errors.Annotate(err, "encode preimage_size")
	}
	if err := enc.Encode(p.Preimage); err != nil {
		return errors.Annotate(err, "encode preimage")
	}
	return nil
}

// HTLCRefundOperation (op 53) is a virtual operation generated when an HTLC
// expires without being redeemed and the locked funds are returned to the sender.
type HTLCRefundOperation struct {
	types.OperationFee
	HTLCId                types.HTLCID      `json:"htlc_id"`
	To                    types.AccountID   `json:"to"`
	OriginalHTLCRecipient types.AccountID   `json:"original_htlc_recipient"`
	HTLCAmount            types.AssetAmount `json:"htlc_amount"`
	PreimageHash          HTLCHash          `json:"htlc_preimage_hash"`
	PreimageSize          types.UInt16      `json:"htlc_preimage_size"`
}

func (p HTLCRefundOperation) Type() types.OperationType {
	return types.OperationTypeHTLCRefund
}

func (p HTLCRefundOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode type")
	}
	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}
	if err := enc.Encode(p.HTLCId); err != nil {
		return errors.Annotate(err, "encode htlc_id")
	}
	if err := enc.Encode(p.To); err != nil {
		return errors.Annotate(err, "encode to")
	}
	if err := enc.Encode(p.OriginalHTLCRecipient); err != nil {
		return errors.Annotate(err, "encode original_htlc_recipient")
	}
	if err := enc.Encode(p.HTLCAmount); err != nil {
		return errors.Annotate(err, "encode htlc_amount")
	}
	if err := enc.Encode(p.PreimageHash); err != nil {
		return errors.Annotate(err, "encode preimage_hash")
	}
	if err := enc.Encode(p.PreimageSize); err != nil {
		return errors.Annotate(err, "encode preimage_size")
	}
	return nil
}
