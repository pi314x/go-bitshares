package types

import (
	"fmt"

	"github.com/denkhaus/bitshares/util"
	"github.com/denkhaus/logging"
	"github.com/juju/errors"
)

type SametFundID struct {
	ObjectID
}

func (p SametFundID) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(p.Instance())); err != nil {
		return errors.Annotate(err, "encode instance")
	}
	return nil
}

func (p *SametFundID) Unmarshal(dec *util.TypeDecoder) error {
	var instance uint64
	if err := dec.DecodeUVarint(&instance); err != nil {
		return errors.Annotate(err, "decode instance")
	}
	p.number = UInt64((uint64(SpaceTypeProtocol) << 56) | (uint64(ObjectTypeSametFund) << 48) | instance)
	return nil
}

type SametFundIDs []SametFundID

func (p SametFundIDs) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(len(p))); err != nil {
		return errors.Annotate(err, "encode length")
	}
	for _, ex := range p {
		if err := enc.Encode(ex); err != nil {
			return errors.Annotate(err, "encode SametFundID")
		}
	}
	return nil
}

func SametFundIDFromObject(ob GrapheneObject) SametFundID {
	id, ok := ob.(*SametFundID)
	if ok {
		return *id
	}
	p := SametFundID{}
	p.MustFromObject(ob)
	if p.ObjectType() != ObjectTypeSametFund {
		panic(fmt.Sprintf("invalid ObjectType: %q has no ObjectType 'ObjectTypeSametFund'", p.ID()))
	}
	return p
}

func NewSametFundID(id string) GrapheneObject {
	gid := new(SametFundID)
	if err := gid.Parse(id); err != nil {
		logging.Errorf("SametFundID parser error %v", errors.Annotate(err, "Parse"))
		return nil
	}
	if gid.ObjectType() != ObjectTypeSametFund {
		logging.Errorf("SametFundID parser error %s", fmt.Sprintf("%q has no ObjectType 'ObjectTypeSametFund'", id))
		return nil
	}
	return gid
}
