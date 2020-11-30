package orders

import (
	"github.com/stretchr/testify/assert"
	"github.com/zond/godip"
	"github.com/zond/godip/orders"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants"
	"strings"
	"testing"
	"wulfheartalexander/advance"
)

func TestParseMovements(t *testing.T) {
	s := scaffoldVariant(t, "Classical")
	opts := s.Phase().Options(s, godip.France)
	wal := opts[godip.Province("wal")]
	mvmts := ParseMovements(wal, s.Graph())
	provincesWithoutConvoy := "lon,lvp,yor"
	provincesWithConvoy := "bel,pic,bre,gas,spa,por,naf,lon"
	for _, p := range destringify(provincesWithoutConvoy){
		assert.Contains(t, mvmts, Move{
			Location: "wal",
			To:       p,
			Convoy:   false,
		})
	}
	for _, p := range destringify(provincesWithConvoy){
		assert.Contains(t, mvmts, Move{
			Location: "wal",
			To:       p,
			Convoy:   true,
		})
	}

	provincesWithoutPossibleConvoy := "pic,bre,bur,gas"
	for _, p := range destringify(provincesWithoutPossibleConvoy){
		assert.Contains(t, ParseMovements(opts[godip.Province("par")], s.Graph()), Move{
			Location: "par",
			To:       p,
			Convoy:   false,
		})
	}
}

func TestParseSupports(t *testing.T){
	s := scaffoldVariant(t, "Classical")
	opts := s.Phase().Options(s, godip.France)
	bur := opts[godip.Province("bur")]
	shs, sms := ParseSupports(bur, s.Graph())
	supportHoldTest := []godip.Province{"mun", "ruh", "par", "mar", "pic"}
	for _, p := range supportHoldTest {
		assert.Contains(t, shs, SupportHold{
			Location: "bur",
			To:       p,
		})
	}

	supportMoveTest := map[godip.Province]string{
		"mun": "ruh",
		"ruh": "bel, mun",
		"pic": "bel, par",
		"eng": "pic, bel",
		"wal": "pic, bel",
	}
	for from, tos := range supportMoveTest {
		for _, to := range destringify(tos){
			assert.Contains(t, sms, SupportMove{
				Location: "bur",
				From:     from,
				To:       to,
			})
		}
	}
}

func TestParseHold(t *testing.T) {
	s := scaffoldVariant(t, "Classical")
	opts := s.Phase().Options(s, godip.France)
	bur := opts[godip.Province("bur")]
	h := ParseHold(bur, s.Graph())
	assert.Equal(t, h, Hold{Location: "bur"})
}

func TestParseConvoy(t *testing.T) {
	s := scaffoldVariant(t, "Classical")
	opts := s.Phase().Options(s, godip.France)
	mid := opts[godip.Province("mid")]
	c := ParseConvoy(mid, s.Graph())
	convoyTest := map[godip.Province]string{
		"wal": "por,gas,spa,naf,bre",
		"pic": "por,gas,spa,naf,bre",

	}
	for key, value := range convoyTest {
		for _, prov := range destringify(value){
			assert.Contains(t,c,Convoy{
				Location: "mid",
				From:     key,
				To:       prov,
			})
		}
	}

}

func TestParseDisband(t *testing.T) {
	s := scaffoldVariant(t, "Classical")
	advance.ToPhaseType(s, godip.Retreat)
	opts := s.Phase().Options(s, godip.Germany)
	bur := opts[godip.Province("mun")]
	h := ParseDisband(bur, s.Graph())
	assert.Equal(t, h, Disband{Location: "mun"})
}

func TestRetreatMovements(t *testing.T) {
	s := scaffoldVariant(t, "Classical")
	advance.ToPhaseType(s, godip.Retreat)
	opts := s.Phase().Options(s, godip.Germany)
	mun := opts[godip.Province("mun")]
	mvmts := ParseMovements(mun, s.Graph())
	provincesWithoutConvoy := "tyr,boh,sil"
	for _, p := range destringify(provincesWithoutConvoy){
		assert.Contains(t, mvmts, Move{
			Location: "mun",
			To:       p,
			Convoy:   false,
		})
	}

}

func TestParseBuilds(t *testing.T) {
	s := scaffoldVariant(t, "Classical")
	advance.ToPhaseType(s, godip.Adjustment)
	opts := s.Phase().Options(s, godip.Austria)
	bur := opts[godip.Province("tri")]
	h := ParseBuild(bur, s.Graph())
	assert.Contains(t, h, Build{Location: "tri", Unit: godip.Fleet})
	assert.Contains(t, h, Build{Location: "tri", Unit: godip.Army})
}

func scaffoldVariant(t *testing.T, variantName string) (s *state.State) {
	variant, found := variants.Variants[variantName]
	if !found {
		t.Fatal("Variant", variantName, "not found")
	}
	s, err := variant.Start()
	if err != nil {
		t.Fatal(err.Error())
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

func destringify(s string) (p []godip.Province) {
	n := strings.ReplaceAll(s, " ", "")
	splitted := strings.Split(n, ",")
	for _, o := range splitted {
		p = append(p, godip.Province(o))
	}
	return
}
