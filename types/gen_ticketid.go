package types

import (
	"fmt"

	"github.com/denkhaus/bitshares/util"
	"github.com/denkhaus/logging"
	"github.com/juju/errors"
)

type TicketID struct {
	ObjectID
}

func (p TicketID) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(p.Instance())); err != nil {
		return errors.Annotate(err, "encode instance")
	}
	return nil
}

func (p *TicketID) Unmarshal(dec *util.TypeDecoder) error {
	var instance uint64
	if err := dec.DecodeUVarint(&instance); err != nil {
		return errors.Annotate(err, "decode instance")
	}
	p.number = UInt64((uint64(SpaceTypeProtocol) << 56) | (uint64(ObjectTypeTicket) << 48) | instance)
	return nil
}

type TicketIDs []TicketID

func (p TicketIDs) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(len(p))); err != nil {
		return errors.Annotate(err, "encode length")
	}
	for _, ex := range p {
		if err := enc.Encode(ex); err != nil {
			return errors.Annotate(err, "encode TicketID")
		}
	}
	return nil
}

func TicketIDFromObject(ob GrapheneObject) TicketID {
	id, ok := ob.(*TicketID)
	if ok {
		return *id
	}
	p := TicketID{}
	p.MustFromObject(ob)
	if p.ObjectType() != ObjectTypeTicket {
		panic(fmt.Sprintf("invalid ObjectType: %q has no ObjectType 'ObjectTypeTicket'", p.ID()))
	}
	return p
}

func NewTicketID(id string) GrapheneObject {
	gid := new(TicketID)
	if err := gid.Parse(id); err != nil {
		logging.Errorf("TicketID parser error %v", errors.Annotate(err, "Parse"))
		return nil
	}
	if gid.ObjectType() != ObjectTypeTicket {
		logging.Errorf("TicketID parser error %s", fmt.Sprintf("%q has no ObjectType 'ObjectTypeTicket'", id))
		return nil
	}
	return gid
}
