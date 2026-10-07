package main

import (
	"encoding/json"
	"os"
	"path/filepath"
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
