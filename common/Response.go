package common

import (
	godipInfluence "github.com/wulfheart/godip-influence"
	"github.com/wulfheart/godip-influence/influenceCalculators"
	"github.com/zond/godip"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/common"
	"wulfheartalexander/orders"
)

type ResponseDTO struct {
	Season string
	Year int
	Type string
	Units map[godip.Province]godip.Unit
	SupplyCenters map[godip.Province]godip.Nation
	Influence godipInfluence.Influence
	PossibleOrders map[godip.Province]orders.FullOrders
	PreviouslyAppliedOrders map[godip.Province]orders.AppliedOrder
	Resolutions map[godip.Province]string
}

func NewResponseDTOfromState(s *state.State, old godipInfluence.Influence, v common.Variant) (r ResponseDTO) {
	r.Season = string(s.Phase().Season())
	var m = s.PreviouslyAppliedOrders()
	r.PreviouslyAppliedOrders = make(map[godip.Province]orders.AppliedOrder)
	for province, adjudicator := range m {
		r.PreviouslyAppliedOrders[province] = orders.FullOrdersFromAdjudicator(adjudicator)
	}
	r.Year = s.Phase().Year()
	r.Type = string(s.Phase().Type())
	var resolutions map[godip.Province]error
	r.Units, r.SupplyCenters, _, _, _, resolutions = s.Dump()
	r.Resolutions = make(map[godip.Province]string)
	for prov, err := range resolutions {
		if err == nil {
			r.Resolutions[prov] = "OK"
		} else {
			r.Resolutions[prov] = err.Error()
		}
	}
	// Remove all godip.Neutral
	r.Influence = influenceCalculators.WebdiplomacyClassic(old, s)
	for key, value := range r.Influence {
		if value == godip.Neutral {
			delete(r.Influence, key)
		}
	}
	r.PossibleOrders = make(map[godip.Province]orders.FullOrders)
	for province, unit := range r.Units {
		opt := s.Phase().Options(s, unit.Nation)
		r.PossibleOrders[province] = orders.ParseOptions(opt[province], s.Graph())
	}


	return
}