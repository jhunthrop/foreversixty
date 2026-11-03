package request

import (
	"slices"
	"testing"
)

// Forever's Juju Might (spell 16329) has two effects in client 1.60.1.70753:
// aura 99, 40 melee attack power (Era's), and aura 124, 40 ranged attack
// power (SpellEffect 870621, Forever's own row). So a bow hunter gains from
// it, and the raid preset gives it to every hunter. Juju Power is Strength,
// which gives a hunter no ranged attack power, so only the melee Survival
// hunter carries it. The landing's preset move (S 252455d15) took Juju Might
// off Beast Mastery and Marksmanship too (round 12, regression-melee lane).
func TestEveryHunterCarriesJujuMightAndOnlySurvivalJujuPower(t *testing.T) {
	all, err := RaidConsumes(loadRealPresets(t), loadRealItemLevels(t))
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct {
		spec      string
		jujuPower bool
	}{
		{"hunter-beast-mastery", false},
		{"hunter-marksmanship", false},
		{"hunter-survival", true},
	} {
		got := ConsumesAtLevel(all[c.spec], 60)
		if !slices.Contains(got, "juju_might") {
			t.Errorf("%s at 60 carries %v, want juju_might among them", c.spec, got)
		}
		if slices.Contains(got, "juju_power") != c.jujuPower {
			t.Errorf("%s at 60: juju_power carried = %v, want %v", c.spec, !c.jujuPower, c.jujuPower)
		}
	}
}
