// logs/engine/summary/threat_test.go
package summary

import (
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
	"github.com/jhunthrop/foreversixty/logs/engine/fight"
)

// Two enemies, one healer. Damage threat lands on the unit hit; healing threat is
// spread over the enemies engaged in the last ten seconds; a taunt is recorded.
func TestThreatIsCreditedPerTargetAndTauntsAreKept(t *testing.T) {
	o, reg := opts(t)
	a := New(o)
	a.Start(at(0))
	events := []event.Event{
		dmg(1, tank, boss, 1, "Melee", 1000, -1),
		dmg(2, mage, bossB, 116, "Frostbolt", 500, -1),
		heal(3, healer, tank, 2050, "Holy Light", 400, 0),
		cast(4, tank, boss, 355, "Taunt"), // add a cast() helper if summary_test.go lacks one
	}
	for _, e := range events {
		reg.Observe(e)
		a.Add(e)
	}
	s := a.Snapshot(fight.Fight{Index: 1, Kind: fight.Encounter, Start: at(0), End: at(5), Players: []string{tank, mage, healer}}, "test")

	pair := func(guid, target string) float64 {
		for _, p := range s.ThreatByTarget {
			if p.GUID == guid && p.TargetGUID == target {
				return p.Threat
			}
		}
		return -1
	}
	if got := pair(tank, boss); got != 1000 {
		t.Errorf("tank on boss = %v, want 1000 (a point of damage is a point of threat)", got)
	}
	if got := pair(mage, bossB); got != 500 {
		t.Errorf("mage on bossB = %v, want 500", got)
	}
	// 400 effective healing at 0.5 = 200 threat, over the two engaged enemies: 100 each.
	if got := pair(healer, boss); got != 100 {
		t.Errorf("healer on boss = %v, want 100", got)
	}
	if got := pair(healer, bossB); got != 100 {
		t.Errorf("healer on bossB = %v, want 100", got)
	}
	// The per-player totals are untouched by attribution.
	for _, row := range s.Threat {
		if row.GUID == healer && row.Threat != 200 {
			t.Errorf("healer total = %v, want 200", row.Threat)
		}
	}
	if len(s.Taunts) != 1 || s.Taunts[0].SourceGUID != tank || s.Taunts[0].TargetGUID != boss || s.Taunts[0].AtMS != 4000 {
		t.Fatalf("taunts = %+v", s.Taunts)
	}
}
