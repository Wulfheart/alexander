package main

import (
	"fmt"
	"github.com/zond/godip"
	"github.com/zond/godip/variants"
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
	// state.SetOrder("wal", orders.Move("wal", "pic"))
	// state.SetOrder("eng", orders.Convoy("eng", "wal", "pic"))
	// state.SetOrder("mar", orders.Move("mar", "spa"))
	// state.SetOrder("bre", orders.Move("bre", "mid"))
	// state.SetOrder("par", orders.Move("par", "gas"))
	// advance.ToPhaseType(state, godip.)


	if err != nil {
		panic(err)
	}
	var opt = state.Phase().Options(state, godip.France)

	fmt.Println(opt)

}
