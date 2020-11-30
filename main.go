package main

import (
	"fmt"
	"github.com/zond/godip"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants"
)

func main() {
	s := scaffoldVariant("Classical")


	var t = s.Phase().Options(s, godip.France)

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
	fleet := godip.Unit{
		Type:   godip.Fleet,
		Nation: godip.France,
	}
	army := godip.Unit{
		Type:   godip.Army,
		Nation: godip.France,
	}
	s.SetUnit("eng", fleet)
	s.SetUnit("mid", fleet)
	s.SetUnit("wal", army)
	s.SetUnit("pic", army)
	s.SetUnit("bur", army)
	s.SetUnit("ruh", army)
	return
}

