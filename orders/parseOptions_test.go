package orders

import (
	"github.com/stretchr/testify/assert"
	"github.com/zond/godip"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants"
	"strings"
	"testing"
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

func scaffoldVariant(t *testing.T, variantName string) (s *state.State) {
	variant, found := variants.Variants[variantName]
	if !found {
		t.Fatal("Variant", variantName, "not found")
	}
	s, err := variant.Start()
	if err != nil {
		t.Fatal(err.Error())
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

func destringify(s string) (p []godip.Province) {
	n := strings.ReplaceAll(s, " ", "")
	splitted := strings.Split(n, ",")
	for _, o := range splitted {
		p = append(p, godip.Province(o))
	}
	return
}