package orders

import (
	"github.com/zond/godip"
	"github.com/zond/godip/orders"
	"time"
)

type OrderInterface interface {
	ToOrder() (p godip.Province, a godip.Adjudicator)
}

type Move struct {
	Location godip.Province
	To godip.Province
	Convoy bool
}

func (m Move) ToOrder() (godip.Province, godip.Adjudicator) {
	return m.Location, orders.Move(m.Location, m.To)
}

type SupportMove struct {
	Location godip.Province
	From godip.Province
	To godip.Province
}

func (s SupportMove) ToOrder() (godip.Province, godip.Adjudicator) {
	return s.Location, orders.SupportMove(s.Location, s.From, s.To)
}

type SupportHold struct {
	Location godip.Province
	To godip.Province
}

func (s SupportHold) ToOrder() (godip.Province, godip.Adjudicator) {
	return s.Location, orders.SupportHold(s.Location, s.To)
}

type Hold struct {
	Location godip.Province
}

func (h Hold) ToOrder() (godip.Province, godip.Adjudicator) {
	return h.Location, orders.Hold(h.Location)
}

type Convoy struct {
	Location godip.Province
	From godip.Province
	To godip.Province
}

func (c Convoy) ToOrder() (godip.Province, godip.Adjudicator) {
	return c.Location, orders.Convoy(c.Location, c.From, c.To)
}

type Disband struct {
	Location godip.Province

}

func (d Disband) ToOrder() (p godip.Province, a godip.Adjudicator) {
	return d.Location, orders.Disband(d.Location, time.Now())
}

type Build struct {
	Location godip.Province
	Unit godip.UnitType
}

func (b Build) ToOrder() (p godip.Province, a godip.Adjudicator) {
	return b.Location, orders.Build(b.Location, b.Unit, time.Now())
}

