package orders

import (
	"github.com/zond/godip"
	orders2 "github.com/zond/godip/orders"
	"reflect"
)

type FullOrders struct {
	Moves        []Move
	SupportMoves []SupportMove
	SupportHolds []SupportHold
	Holds        Hold
	Convoy       []Convoy
	Builds       []Build
	Disbands     Disband
}

type AppliedOrder struct {
	Move        *Move        `json:",omitempty"`
	SupportMove *SupportMove `json:",omitempty"`
	SupportHold *SupportHold `json:",omitempty"`
	Hold        *Hold        `json:",omitempty"`
	Convoy      *Convoy      `json:",omitempty"`
	Build       *Build       `json:",omitempty"`
	Disband     *Disband     `json:",omitempty"`
}

func FullOrdersFromAdjudicator(a godip.Adjudicator) (o AppliedOrder) {
	var t = a.Targets()
	switch a.Type() {
	case "Move":
		{
			o.Move = &Move{
				Location: t[0],
				To:       t[1],
				Convoy:   false,
			}
		}
	case "Build":
		{
			var typ godip.UnitType
			if reflect.TypeOf(a) == reflect.TypeOf(orders2.BuildOrder) {
				typ = godip.UnitType(reflect.ValueOf(a).Elem().FieldByName("typ").String())

			} else {
				panic("Some type errors on build")
			}
			o.Build = &Build{
				Location: t[0],
				Unit:     typ,
			}
		}
	case "SupportMove":
		{
			o.SupportMove = &SupportMove{
				Location: t[0],
				From:     t[1],
				To:       t[2],
			}
		}
	case "SupportHold":
		{
			o.SupportHold = &SupportHold{
				Location: t[0],
				To:       t[1],
			}
		}
	case "Convoy":
		{
			o.Convoy = &Convoy{
				Location: t[0],
				From:     t[1],
				To:       t[2],
			}
		}
	case "Disband":
		{
			o.Disband = &Disband{Location: t[0]}
		}
	case "Hold":
		{
			o.Hold = &Hold{Location: t[0]}
		}
	}
	return
}
