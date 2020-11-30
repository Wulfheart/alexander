package orders

import (
	"fmt"
	"github.com/zond/godip"
)

func ParseOptions(o godip.Options, g godip.Graph) (dto DTO, err error) {
	for province, value := range o {
		if _, ok := province.(godip.Province); ok {
			_ = value
		} else {
			return nil, fmt.Errorf("key is not a province")
		}
	}
	return nil, nil
}

func ParseMovements(o godip.Options, g godip.Graph) (movements []Move) {
	orders, ok := o[godip.OrderType("Move")]
	if !ok {
		return []Move{}
	}
	for src, targets := range orders {
		if val, k := src.(godip.SrcProvince); k {
			for provs, _ := range targets {
				if tar, ok2 := provs.(godip.Province); ok2 {
					for key := range g.Edges(godip.Province(val), false) {
						if key == tar {
							movements = append(movements, Move{
								Location: godip.Province(val),
								To:       tar,
								Convoy:   false,
							})
						}
					}
				}
			}
		}
	}

	// Parsing Movements via convoy
	orders, ok = o[godip.OrderType("MoveViaConvoy")]
	if !ok {
		return
	}
	for src, targets := range orders {
		if val, k := src.(godip.SrcProvince); k {
			for provs, _ := range targets {
				if tar, ok2 := provs.(godip.Province); ok2 {

					movements = append(movements, Move{
						Location: godip.Province(val),
						To:       tar,
						Convoy:   true,
					})
				}
			}
		}
	}

	return
}


func ParseSupports(o godip.Options, g godip.Graph) (shs []SupportHold, sms []SupportMove) {
	orders, ok := o[godip.OrderType("Support")]
	if !ok {
		return []SupportHold{}, []SupportMove{}
	}
	for src, targets := range orders {
		if val, k := src.(godip.SrcProvince); k {
			loc := godip.Province(val)
			for base, tos := range targets {
				if baseProvince, ok2 := base.(godip.Province); ok2 {
					for ts := range tos {
						if pts, ok3 := ts.(godip.Province); ok3 {
							if baseProvince == pts {
								shs = append(shs, SupportHold{
									Location: loc,
									To:       baseProvince,
								})
							} else {
								sms = append(sms, SupportMove{
									Location: loc,
									From:     baseProvince,
									To:       pts,
								})
							}
						}
					}
				}
			}
		}
	}
	return
}

func ParseHold(o godip.Options, g godip.Graph) (h Hold) {
	orders, ok := o[godip.OrderType("Hold")]
	if !ok {
		return Hold{}
	}
	for o := range orders {
		if src, ok := o.(godip.SrcProvince); ok {
			return Hold{Location: godip.Province(src)}
		}
	}
	return
}

func ParseConvoy(o godip.Options, g godip.Graph) (cvy []Convoy){
	orders, ok := o[godip.OrderType("Convoy")]
	if !ok {
		return []Convoy{}
	}
	for src, targets := range orders {
		if val, k := src.(godip.SrcProvince); k {
			loc := godip.Province(val)
			for base, tos := range targets {
				if baseProvince, ok2 := base.(godip.Province); ok2 {
					for ts := range tos {
						if pts, ok3 := ts.(godip.Province); ok3 {

								cvy = append(cvy, Convoy{
									Location: loc,
									From:     baseProvince,
									To:       pts,
								})

						}
					}
				}
			}
		}
	}
	return
}

func ParseDisband(o godip.Options, g godip.Graph) (h Disband) {
	orders, ok := o[godip.OrderType("Disband")]
	if !ok {
		return Disband{}
	}
	for o := range orders {
		if src, ok := o.(godip.SrcProvince); ok {
			return Disband{Location: godip.Province(src)}
		}
	}
	return
}