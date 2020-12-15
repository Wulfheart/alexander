package common

import (
	godipInfluence "github.com/wulfheart/godip-influence"
	"github.com/zond/godip"
	orders2 "github.com/zond/godip/orders"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/common"
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
	switch o.Type {
	case "Move":
		return orders2.Move(o.Payload.Location, o.Payload.To)
	case "SupportHold":
		return orders2.SupportHold(o.Payload.Location, o.Payload.To)
	case "SupportMove":
		return orders2.SupportMove(o.Payload.Location, o.Payload.From, o.Payload.To)
	case "Hold":
		return orders2.Hold(o.Payload.Location)
	case "Disband":
		return orders2.Disband(o.Payload.Location, time.Now())
	case "Build":
		return orders2.Build(o.Payload.Location, o.Payload.Unit, time.Now())
	case "Convoy":
		return orders2.Convoy(o.Payload.Location, o.Payload.From, o.Payload.To)
	}
	return nil
}
