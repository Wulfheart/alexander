package advance

import (
	"github.com/stretchr/testify/assert"
	"github.com/zond/godip"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants"
	"testing"
)

func TestOnePhase(t *testing.T){
	s := scaffoldVariant(t, "Classical")
	OnePhase(s)
	assert.Equal(t, godip.Retreat, s.Phase().Type())
	assert.Equal(t, godip.Spring, s.Phase().Season())
}

func TestToPhaseType(t *testing.T){
	s := scaffoldVariant(t, "Classical")
	ToPhaseType(s, godip.Retreat)
	assert.Equal(t, godip.Retreat, s.Phase().Type())
	assert.Equal(t, godip.Spring, s.Phase().Season())
	ToPhaseType(s, godip.Adjustment)
	assert.Equal(t, godip.Adjustment, s.Phase().Type())
	assert.Equal(t, godip.Fall, s.Phase().Season())
}

func TestToSeason(t *testing.T) {
	s := scaffoldVariant(t, "Classical")
	ToSeason(s, godip.Fall)
	assert.Equal(t, godip.Movement, s.Phase().Type())
	assert.Equal(t, godip.Fall, s.Phase().Season())
}


func scaffoldVariant(t *testing.T, variantName string) (s *state.State){
	variant, found := variants.Variants[variantName]
	if !found {
		t.Fatal("Variant", variantName, "not found")
	}
	s, err := variant.Start()
	if err != nil {
		t.Fatal(err.Error())
	}
	return
}