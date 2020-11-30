package common

import (
	godipInfluence "github.com/wulfheart/godip-influence"
	"github.com/zond/godip"
	orders2 "github.com/zond/godip/orders"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/common"
	"strings"
	"time"
)

type RequestDTO struct {
	Season string
	Year int
	Type string
	Units map[godip.Province]godip.Unit
	SupplyCenters map[godip.Province]godip.Nation
	Influence godipInfluence.Influence
	Orders map[godip.Province]OrderRequestDTO
}

func (r *RequestDTO) State(v common.Variant) *state.State{
	var o = make(map[godip.Province]godip.Adjudicator)
	for province, or := range r.Orders{
		o[province] = or.Parse()
	}
	return v.Blank(v.Phase(r.Year, godip.Season(r.Season), godip.PhaseType(r.Type))).Load(r.Units, r.SupplyCenters, nil, nil, nil, o)
}

type OrderRequestDTO struct {
	Type string
	Payload struct{
		Location godip.Province
		From godip.Province
		To godip.Province
		Convoy bool
		Unit godip.UnitType
	}
}

func (o *OrderRequestDTO) Parse() godip.Adjudicator{
	switch strings.ToLower(o.Type) {
	case "move":
		return orders2.Move(o.Payload.Location, o.Payload.To)
	case "supporthold":
		return orders2.SupportHold(o.Payload.Location, o.Payload.To)
	case "supportmove":
		return orders2.SupportMove(o.Payload.Location, o.Payload.From, o.Payload.To)
	case "hold":
		return orders2.Hold(o.Payload.Location)
	case "disband":
		return orders2.Disband(o.Payload.Location, time.Now())
	case "build":
		return orders2.Build(o.Payload.Location, o.Payload.Unit, time.Now())
	case "convoy":
		return orders2.Convoy(o.Payload.Location, o.Payload.From, o.Payload.To)
	}
	return nil
}
