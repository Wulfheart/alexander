package common

import (
	"github.com/zond/godip"
	"github.com/zond/godip/variants/common"
)

type MetaDto struct {
	Provinces map[godip.Province]string
	SupplyCenters []godip.Province
	UnitTypes []godip.UnitType
	Seasons []godip.Season
	Phases []godip.PhaseType
	Nations []godip.Nation
	OrderTypes []godip.OrderType
}

func CreateMetaDtoFromVariant(v common.Variant) MetaDto{
	return MetaDto{
		Provinces:  v.ProvinceLongNames,
		// SupplyCenters: v.S
		UnitTypes:  v.UnitTypes,
		Seasons:    v.Seasons,
		Phases:     v.PhaseTypes,
		Nations:    v.Nations,
		OrderTypes: v.Parser.OrderTypes(),
	}
}
