package types

import (
	"fmt"

	"github.com/denkhaus/bitshares/util"
	"github.com/denkhaus/logging"
	"github.com/juju/errors"
)

type HTLCID struct {
	ObjectID
}

func (p HTLCID) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(p.Instance())); err != nil {
		return errors.Annotate(err, "encode instance")
	}
	return nil
}

func (p *HTLCID) Unmarshal(dec *util.TypeDecoder) error {
	var instance uint64
	if err := dec.DecodeUVarint(&instance); err != nil {
		return errors.Annotate(err, "decode instance")
	}
	p.number = UInt64((uint64(SpaceTypeProtocol) << 56) | (uint64(ObjectTypeHTLC) << 48) | instance)
	return nil
}

type HTLCIDs []HTLCID

func (p HTLCIDs) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(len(p))); err != nil {
		return errors.Annotate(err, "encode length")
	}
	for _, ex := range p {
		if err := enc.Encode(ex); err != nil {
			return errors.Annotate(err, "encode HTLCID")
		}
	}
	return nil
}

func HTLCIDFromObject(ob GrapheneObject) HTLCID {
	id, ok := ob.(*HTLCID)
	if ok {
		return *id
	}
	p := HTLCID{}
	p.MustFromObject(ob)
	if p.ObjectType() != ObjectTypeHTLC {
		panic(fmt.Sprintf("invalid ObjectType: %q has no ObjectType 'ObjectTypeHTLC'", p.ID()))
	}
	return p
}

func NewHTLCID(id string) GrapheneObject {
	gid := new(HTLCID)
	if err := gid.Parse(id); err != nil {
		logging.Errorf("HTLCID parser error %v", errors.Annotate(err, "Parse"))
		return nil
	}
	if gid.ObjectType() != ObjectTypeHTLC {
		logging.Errorf("HTLCID parser error %s", fmt.Sprintf("%q has no ObjectType 'ObjectTypeHTLC'", id))
		return nil
	}
	return gid
}
