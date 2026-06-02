package types

import (
	"fmt"

	"github.com/denkhaus/bitshares/util"
	"github.com/denkhaus/logging"
	"github.com/juju/errors"
)

type CustomAuthorityID struct {
	ObjectID
}

func (p CustomAuthorityID) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(p.Instance())); err != nil {
		return errors.Annotate(err, "encode instance")
	}
	return nil
}

func (p *CustomAuthorityID) Unmarshal(dec *util.TypeDecoder) error {
	var instance uint64
	if err := dec.DecodeUVarint(&instance); err != nil {
		return errors.Annotate(err, "decode instance")
	}
	p.number = UInt64((uint64(SpaceTypeProtocol) << 56) | (uint64(ObjectTypeCustomAuthority) << 48) | instance)
	return nil
}

type CustomAuthorityIDs []CustomAuthorityID

func (p CustomAuthorityIDs) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(len(p))); err != nil {
		return errors.Annotate(err, "encode length")
	}
	for _, ex := range p {
		if err := enc.Encode(ex); err != nil {
			return errors.Annotate(err, "encode CustomAuthorityID")
		}
	}
	return nil
}

func CustomAuthorityIDFromObject(ob GrapheneObject) CustomAuthorityID {
	id, ok := ob.(*CustomAuthorityID)
	if ok {
		return *id
	}
	p := CustomAuthorityID{}
	p.MustFromObject(ob)
	if p.ObjectType() != ObjectTypeCustomAuthority {
		panic(fmt.Sprintf("invalid ObjectType: %q has no ObjectType 'ObjectTypeCustomAuthority'", p.ID()))
	}
	return p
}

func NewCustomAuthorityID(id string) GrapheneObject {
	gid := new(CustomAuthorityID)
	if err := gid.Parse(id); err != nil {
		logging.Errorf("CustomAuthorityID parser error %v", errors.Annotate(err, "Parse"))
		return nil
	}
	if gid.ObjectType() != ObjectTypeCustomAuthority {
		logging.Errorf("CustomAuthorityID parser error %s", fmt.Sprintf("%q has no ObjectType 'ObjectTypeCustomAuthority'", id))
		return nil
	}
	return gid
}
