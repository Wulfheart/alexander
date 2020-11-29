package advance

import (
	"github.com/zond/godip"
	"github.com/zond/godip/state"
)

func OnePhase(s *state.State){
	err := s.Next()
	if err != nil {
		panic(err)
	}
}

func ToPhaseType(s *state.State, p godip.PhaseType){
	for s.Phase().Type() != p {
		OnePhase(s)
	}
}

func ToSeason(s *state.State, j godip.Season){
	for s.Phase().Season() != j {
		OnePhase(s)
	}
}