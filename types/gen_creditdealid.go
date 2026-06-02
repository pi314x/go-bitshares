package types

import (
	"fmt"

	"github.com/denkhaus/bitshares/util"
	"github.com/denkhaus/logging"
	"github.com/juju/errors"
)

type CreditDealID struct {
	ObjectID
}

func (p CreditDealID) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(p.Instance())); err != nil {
		return errors.Annotate(err, "encode instance")
	}
	return nil
}

func (p *CreditDealID) Unmarshal(dec *util.TypeDecoder) error {
	var instance uint64
	if err := dec.DecodeUVarint(&instance); err != nil {
		return errors.Annotate(err, "decode instance")
	}
	p.number = UInt64((uint64(SpaceTypeProtocol) << 56) | (uint64(ObjectTypeCreditDeal) << 48) | instance)
	return nil
}

type CreditDealIDs []CreditDealID

func (p CreditDealIDs) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(len(p))); err != nil {
		return errors.Annotate(err, "encode length")
	}
	for _, ex := range p {
		if err := enc.Encode(ex); err != nil {
			return errors.Annotate(err, "encode CreditDealID")
		}
	}
	return nil
}

func CreditDealIDFromObject(ob GrapheneObject) CreditDealID {
	id, ok := ob.(*CreditDealID)
	if ok {
		return *id
	}
	p := CreditDealID{}
	p.MustFromObject(ob)
	if p.ObjectType() != ObjectTypeCreditDeal {
		panic(fmt.Sprintf("invalid ObjectType: %q has no ObjectType 'ObjectTypeCreditDeal'", p.ID()))
	}
	return p
}

func NewCreditDealID(id string) GrapheneObject {
	gid := new(CreditDealID)
	if err := gid.Parse(id); err != nil {
		logging.Errorf("CreditDealID parser error %v", errors.Annotate(err, "Parse"))
		return nil
	}
	if gid.ObjectType() != ObjectTypeCreditDeal {
		logging.Errorf("CreditDealID parser error %s", fmt.Sprintf("%q has no ObjectType 'ObjectTypeCreditDeal'", id))
		return nil
	}
	return gid
}
