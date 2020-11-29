package main

import (
	"fmt"
	"github.com/zond/godip"
	"github.com/zond/godip/variants"
	"wulfheartalexander/advance"
)

func main() {
	fmt.Println("Hello World")
	variant, found := variants.Variants["Classical"]
	if !found {
		panic("not found")
	}

	state, err := variant.Start()
	state.SetUnit("eng", godip.Unit{
		Type:   godip.Fleet,
		Nation: godip.England,
	})
	state.SetUnit("wal", godip.Unit{
		Type:   godip.Army,
		Nation: godip.England,
	})
	state.SetSC("spa", godip.France)

	advance.ToPhaseType(state, godip.Adjustment)


	if err != nil {
		panic(err)
	}
	var opt = state.Phase().Options(state, godip.France)
	var n = opt[godip.Province("wal")]
	fmt.Println(n)

}
