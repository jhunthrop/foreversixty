package main

import (
	"math"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/internal/simdb"
)

// The ranker scores a weapon's DPS off items/<class>.json, but every
// verification sim swings the engine's own simdb.bin row. Elemental Mage
// Staff (944) is the repro: items/druid.json carries classic-db's
// 147-221 (57.5 DPS) while simdb.bin carries wowhead's 109-164 (42.66
// DPS), so score() ranked it the best launch-day feral weapon while the
// engine measured it level with the 46-DPS Warden Staff.
func TestEngineWeaponDamageReplacesTheItemFileDamage(t *testing.T) {
	staff := candidate{ID: 944, Name: "Elemental Mage Staff", DamageMin: 147, DamageMax: 221, Speed: 3.2, DPS: 57.5, TwoHand: true}
	ring := candidate{ID: 1, Name: "A Ring"}
	unknown := candidate{ID: 2, Name: "Not In The Engine", DamageMin: 10, DamageMax: 20, Speed: 2, DPS: 7.5}
	lookup := func(id int32) (simdb.WeaponDamage, bool) {
		if id == 944 {
			return simdb.WeaponDamage{Min: 109, Max: 164, Speed: 3.2}, true
		}
		return simdb.WeaponDamage{}, false
	}
	in := []candidate{staff, ring, unknown}
	out := withEngineWeaponDamage(in, lookup)

	if got := out[0]; got.DamageMin != 109 || got.DamageMax != 164 || got.Speed != 3.2 || math.Abs(got.DPS-42.66) > 0.005 {
		t.Fatalf("staff = %v-%v @%v (%v dps), want the engine's 109-164 @3.2 (42.66 dps)", got.DamageMin, got.DamageMax, got.Speed, got.DPS)
	}
	if out[1].DPS != 0 || out[1].Speed != 0 {
		t.Fatalf("ring gained weapon damage: %+v", out[1])
	}
	if out[2].DPS != 7.5 {
		t.Fatalf("an item the engine lacks keeps the item file's dps, got %v", out[2].DPS)
	}
	if in[0].DPS != 57.5 {
		t.Fatalf("input mutated: %v", in[0].DPS)
	}
}

// The real embedded database must answer for the repro item, so the
// ranker's scores follow the engine's swing.
func TestEmbeddedSimDBStatesElementalMageStaffDamage(t *testing.T) {
	w, ok := simdb.Weapon(944)
	if !ok {
		t.Skip("Elemental Mage Staff is not in this build's simdb.bin")
	}
	if w.Speed <= 0 || w.Max < w.Min || w.Min <= 0 {
		t.Fatalf("Weapon(944) = %+v, want a real damage range and speed", w)
	}
	if !nonWeaponIsAbsent(t) {
		t.Fatal("Weapon reported damage for a non-weapon")
	}
}

func nonWeaponIsAbsent(t *testing.T) bool {
	t.Helper()
	const handOfJustice = 11815 // a trinket
	_, ok := simdb.Weapon(handOfJustice)
	return !ok
}
