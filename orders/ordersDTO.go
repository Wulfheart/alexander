package orders

import "github.com/zond/godip"

type DTO map[godip.Province]FullOrders

type FullOrders struct{
	Moves []Move
	MovesViaConvoy []MoveViaConvoy
	SupportMoves []SupportMove
	SupportHolds []SupportHold
	Holds Hold
	Convoy []Convoy
	Builds []Build
	Disbands []Disband
}
