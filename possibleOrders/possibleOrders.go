package possibleOrders

import "github.com/zond/godip"
const (
	Move godip.OrderType = "Move"
	MoveViaConvoy godip.OrderType = "Move"
	SupportMove godip.OrderType = "Move"
	SupportHold godip.OrderType = "Move"
	Convoy godip.OrderType = "Move"
	Build godip.OrderType = "Move"
)
type Possibility map[godip.OrderType][]Payload

type Payload struct {
	source godip.Province
	target godip.Province
}
