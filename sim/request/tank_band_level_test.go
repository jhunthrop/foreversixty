package request

import "testing"

// Sim-accuracy round 29 (walks-heal): a tank between two bands wears the lower band's gear,
// and the heal profile's band rule (data/curated/heal-profile.json bands.rule) has it fight
// that band: "a level between two bands fights the band at or below it". The tank boss
// scaled its damage and healing by the character's own level instead, so a bare level 25
// to 29 tank in its band 20 gear took up to twice the damage the band 20 boss deals, and
// the Protection paladin and the Feral bear died in every fight (round 28 hybrids P4).
func TestTankProfileFightsTheBandAtOrBelowTheLevel(t *testing.T) {
	cases := []struct{ level, band int }{
		{20, 20}, {25, 20}, {29, 20},
		{30, 30}, {35, 30}, {39, 30},
		{40, 40}, {48, 40}, {50, 50}, {59, 50}, {60, 60},
	}
	for _, c := range cases {
		got, err := TankProfileForLevel(c.level)
		if err != nil {
			t.Fatal(err)
		}
		want, err := TankProfileForLevel(c.band)
		if err != nil {
			t.Fatal(err)
		}
		if got != want {
			t.Errorf("level %d tank fight %+v, want band %d's %+v", c.level, got, c.band, want)
		}
	}
	// Below the first band there is no band gear or band row to fight, so a level keeps its
	// own scaling (the ladder's level 10 rung).
	ten, err := TankProfileForLevel(10)
	if err != nil {
		t.Fatal(err)
	}
	twenty, err := TankProfileForLevel(20)
	if err != nil {
		t.Fatal(err)
	}
	if ten.Boss.MinBaseDamage >= twenty.Boss.MinBaseDamage {
		t.Errorf("level 10 boss %v, want below band 20's %v", ten.Boss.MinBaseDamage, twenty.Boss.MinBaseDamage)
	}
}
