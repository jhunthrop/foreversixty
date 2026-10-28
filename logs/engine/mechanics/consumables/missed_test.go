package consumables

import "testing"

func TestMissedAtPull_OnlyAKnownPullOnACompleteListProvesAbsence(t *testing.T) {
	rc := RoleCatalogue{Complete: map[string]bool{CategoryFlask: true, CategoryCombatPotion: true}}
	flask := []Entry{{SpellID: 17626}}
	for _, c := range []struct {
		name     string
		pull     Pull
		category string
		want     bool
	}{
		{"no sheet", Pull{}, CategoryFlask, false},
		{"complete list, nothing up", Pull{Known: true}, CategoryFlask, true},
		{"complete list, a flask up", Pull{Known: true, Present: map[string][]Entry{CategoryFlask: flask}}, CategoryFlask, false},
		{"incomplete list, nothing up", Pull{Known: true}, CategoryFood, false},
		{"complete potion list, nothing drunk", Pull{Known: true}, CategoryCombatPotion, true},
		{"complete potion list, one cast", Pull{Known: true, PotionUses: 1}, CategoryCombatPotion, false},
	} {
		if got := rc.MissedAtPull(c.pull, c.category); got != c.want {
			t.Errorf("%s: MissedAtPull = %v, want %v", c.name, got, c.want)
		}
	}
}
