package main

import (
	"fmt"
	"github.com/zond/godip"
	"github.com/zond/godip/orders"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants"
	"wulfheartalexander/advance"
)

func main() {
	s := scaffoldVariant("Classical")


	advance.ToPhaseType(s, godip.Adjustment)
	var t = s.Phase().Options(s, godip.Austria)
	fmt.Println(t)

}

func scaffoldVariant( variantName string) (s *state.State) {
	variant, found := variants.Variants[variantName]
	if !found {
		panic(fmt.Sprint("Variant", variantName, "not found"))
	}
	s, err := variant.Start()
	if err != nil {
		panic(err.Error())
	}
	fleetFrance := godip.Unit{
		Type:   godip.Fleet,
		Nation: godip.France,
	}
	armyFrance := godip.Unit{
		Type:   godip.Army,
		Nation: godip.France,
	}
	s.SetUnit("eng", fleetFrance)
	s.SetUnit("mid", fleetFrance)
	s.SetUnit("wal", armyFrance)
	s.SetUnit("pic", armyFrance)
	s.SetUnit("bur", armyFrance)
	s.SetUnit("ruh", armyFrance)
	s.SetOrder("bud", orders.Move("bud", "ser"))
	s.SetOrder("vie", orders.Move("vie", "gal"))
	s.SetOrder("tri", orders.Move("tri", "alb"))
	s.SetOrder("bur", orders.Move("bur", "mun"))
	s.SetOrder("ruh", orders.SupportMove("ruh", "bur", "mun"))
	return
}

