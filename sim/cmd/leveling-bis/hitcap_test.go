package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/wowsims/classic/sim/core"
)

func TestHitToCapOfACasterIsAbsent(t *testing.T) {
	if got := hitToCapFromProfile(core.HitProfile{Hit: 3}); got != nil {
		t.Fatalf("a caster has no physical miss table to cap against, got %+v", got)
	}
}

func TestHitToCapOfASingleWeaponUserHasNoWhiteFigure(t *testing.T) {
	got := hitToCapFromProfile(core.HitProfile{Physical: true, Hit: 2, Suppression: 1, SpecialCap: 9, WhiteCap: 9})
	if got == nil || got.Baseline != 2 || got.Specials != 7 || got.White != nil {
		t.Fatalf("hit_to_cap = %+v, want baseline 2, specials 7, no white figure", got)
	}
}

func TestHitToCapOfADualWielderCarriesTheWhiteCap(t *testing.T) {
	got := hitToCapFromProfile(core.HitProfile{Physical: true, Hit: 4, Suppression: 1, SpecialCap: 9, WhiteCap: 28, DualWielding: true})
	if got == nil || got.Specials != 5 || got.White == nil || *got.White != 24 {
		t.Fatalf("hit_to_cap = %+v, want specials 5, white 24", got)
	}
}

func TestHitToCapNeverGoesNegative(t *testing.T) {
	got := hitToCapFromProfile(core.HitProfile{Physical: true, Hit: 12, SpecialCap: 9, WhiteCap: 9})
	if got.Specials != 0 {
		t.Fatalf("specials past the cap = %v, want 0", got.Specials)
	}
}

func TestRunSpecPublishesHitToCapFromTheEngine(t *testing.T) {
	outDir := t.TempDir()
	fake := &fakeEngine{
		DefaultDPS: 500,
		WeightsResult: map[string]api.StatWeight{
			"ranged_attack_power": {Stat: "ranged_attack_power", Weight: 1.0},
			"agility":             {Stat: "agility", Weight: 1.8},
		},
		HitProfile: core.HitProfile{Physical: true, Hit: 1, Suppression: 1, SpecialCap: 9, WhiteCap: 9},
	}
	err := runSpec(fake, repoRootFixture, buildDirFixture(), "testbuild", outDir, "hunter-marksmanship", []int{20}, 5, identityTalentLayout)
	if err != nil {
		t.Fatalf("runSpec: %v", err)
	}
	b, err := os.ReadFile(filepath.Join(outDir, "hunter-marksmanship.json"))
	if err != nil {
		t.Fatal(err)
	}
	var out struct {
		Bands []struct {
			HitToCap *struct {
				Baseline float64  `json:"baseline"`
				Specials float64  `json:"specials"`
				White    *float64 `json:"white"`
			} `json:"hit_to_cap"`
		} `json:"bands"`
	}
	if err := json.Unmarshal(b, &out); err != nil {
		t.Fatal(err)
	}
	if len(out.Bands) == 0 {
		t.Fatal("no bands written")
	}
	for i, band := range out.Bands {
		if band.HitToCap == nil || band.HitToCap.Baseline != 1 || band.HitToCap.Specials != 8 || band.HitToCap.White != nil {
			t.Errorf("band %d hit_to_cap = %+v, want baseline 1, specials 8", i, band.HitToCap)
		}
	}
}

func TestHitToCapOfACasterSpecIsAbsentEvenWithAWeapon(t *testing.T) {
	physical := core.HitProfile{Physical: true, Hit: 1}
	if got := hitToCapFor("druid-balance", physical); got != nil {
		t.Fatalf("druid-balance never stands in melee, got %+v", got)
	}
	if got := hitToCapFor("shaman-elemental", physical); got != nil {
		t.Fatalf("shaman-elemental never stands in melee, got %+v", got)
	}
	if got := hitToCapFor("warrior-fury", physical); got == nil {
		t.Fatal("a melee spec keeps its figure")
	}
}

func casterProfile(spell core.SpellHitProfile) core.HitProfile {
	return core.HitProfile{Hit: spell.Hit, Spell: true, SpellProfile: spell}
}

func TestSpellHitToCapOfACasterIsTheDistanceToSixteen(t *testing.T) {
	got, ok := hitToCapFor("shaman-elemental", casterProfile(core.SpellHitProfile{Hit: 6, SchoolHit: 6, Cap: 16})).(*spellHitToCap)
	if !ok || got.Kind != "spell" || got.Baseline != 6 || got.Spell != 10 || got.School != nil {
		t.Fatalf("hit_to_cap = %+v, want kind spell, baseline 6, spell 10, no school", got)
	}
}

func TestSpellHitToCapCarriesTheSchoolFigureWhenATalentAddsHit(t *testing.T) {
	profile := casterProfile(core.SpellHitProfile{Hit: 6, SchoolHit: 11, SchoolNames: []string{"Fire", "Frost"}, Cap: 16})
	got, ok := hitToCapFor("mage-fire", profile).(*spellHitToCap)
	if !ok || got.Spell != 10 || got.School == nil {
		t.Fatalf("hit_to_cap = %+v, want spell 10 with a school figure", got)
	}
	if got.School.ToCap != 5 || got.School.Hit != 11 || len(got.School.Names) != 2 || got.School.Names[0] != "Fire" {
		t.Errorf("school = %+v, want fire and frost at 11%%, 5 to cap", got.School)
	}
}

func TestSpellHitToCapNeverGoesNegative(t *testing.T) {
	got := spellHitToCapFromProfile(casterProfile(core.SpellHitProfile{Hit: 14, SchoolHit: 19, SchoolNames: []string{"Shadow"}, Cap: 16}))
	if got.Spell != 2 || got.School.ToCap != 0 {
		t.Fatalf("hit_to_cap = %+v / %+v, want spell 2, school 0", got, got.School)
	}
}

func TestSpellHitToCapOfACasterWithoutSpellsIsAbsent(t *testing.T) {
	if got := hitToCapFor("warlock-affliction", core.HitProfile{Hit: 3}); got != nil {
		t.Fatalf("a profile with no spells publishes nothing, got %+v", got)
	}
}

func TestMeleeSpecKeepsTheMeleeShapeEvenWithSpells(t *testing.T) {
	profile := core.HitProfile{Physical: true, Hit: 2, SpecialCap: 9, WhiteCap: 9, Spell: true, SpellProfile: core.SpellHitProfile{Hit: 2, SchoolHit: 2, Cap: 16}}
	got, ok := hitToCapFor("shaman-enhancement", profile).(*hitToCap)
	if !ok || got.Specials != 7 {
		t.Fatalf("hit_to_cap = %+v, want the melee figure with specials 7", got)
	}
}

func TestPublishedShapesAreToldApartByKind(t *testing.T) {
	melee, _ := json.Marshal(hitToCapFromProfile(core.HitProfile{Physical: true, Hit: 2, SpecialCap: 9, WhiteCap: 9}))
	if string(melee) != `{"baseline":2,"specials":7}` {
		t.Errorf("melee shape = %s, want it unchanged", melee)
	}
	spell, _ := json.Marshal(spellHitToCapFromProfile(casterProfile(core.SpellHitProfile{Hit: 6, SchoolHit: 11, SchoolNames: []string{"Fire"}, Cap: 16})))
	want := `{"kind":"spell","baseline":6,"spell":10,"school":{"names":["Fire"],"hit":11,"to_cap":5}}`
	if string(spell) != want {
		t.Errorf("spell shape = %s, want %s", spell, want)
	}
}

func TestPublishedHitToCapRoundTripsBothShapes(t *testing.T) {
	cases := map[string]hitToCapFigure{
		`{"baseline":2,"specials":7,"white":24}`:   &hitToCap{Baseline: 2, Specials: 7, White: ptr(24.0)},
		`{"kind":"spell","baseline":0,"spell":16}`: &spellHitToCap{Kind: "spell", Spell: 16},
	}
	for raw, want := range cases {
		var got publishedHitToCap
		if err := json.Unmarshal([]byte(raw), &got); err != nil {
			t.Fatalf("%s: %v", raw, err)
		}
		out, err := json.Marshal(&got)
		if err != nil || string(out) != raw {
			t.Errorf("round trip of %s = %s, %v", raw, out, err)
		}
		if reflect.TypeOf(got.hitToCapFigure) != reflect.TypeOf(want) {
			t.Errorf("%s decoded as %T, want %T", raw, got.hitToCapFigure, want)
		}
	}
	var bad publishedHitToCap
	if err := json.Unmarshal([]byte(`{"kind":"mystery"}`), &bad); err == nil {
		t.Error("an unknown kind must fail loudly")
	}
}

func ptr(v float64) *float64 { return &v }
