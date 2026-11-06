package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/leveling"
	"github.com/jhunthrop/foreversixty/sim/request"
	"github.com/jhunthrop/foreversixty/sim/specs"
)

// Round 25 platform: -level 10 to 19 runs the band-20 entry, and every slot
// whose item needs a higher level used to be left empty. Rogues fought
// unarmed below 19 (Blackfang 20, Assassin's Blade 19) and hunters had no bow
// below 20 (Ranger Bow 20), so no melee or hunter rank row at 10 to 19 could
// be judged. A slot the level cannot wear now takes a weapon it can.

type itemKind struct {
	ClassID    int `json:"class_id"`
	SubclassID int `json:"subclass_id"`
}

func readItemKinds(t *testing.T) map[int]itemKind {
	t.Helper()
	dir, err := buildDirectory(realRoot)
	if err != nil {
		t.Fatal(err)
	}
	raw, err := os.ReadFile(filepath.Join(dir, "items.json"))
	if err != nil {
		t.Fatal(err)
	}
	var rows []struct {
		ID int `json:"id"`
		itemKind
	}
	if err := json.Unmarshal(raw, &rows); err != nil {
		t.Fatal(err)
	}
	out := make(map[int]itemKind, len(rows))
	for _, r := range rows {
		out[r.ID] = r.itemKind
	}
	return out
}

const (
	kindWeapon   = 2
	kindArmor    = 4
	kindShield   = 6
	subclassWand = 19
)

// launchers are the ranged weapons Auto Shot fires: bow, gun, crossbow.
var launchers = map[int]bool{2: true, 3: true, 18: true}

func slotItems(gear []api.GearSlot) map[string]int {
	out := make(map[string]int, len(gear))
	for _, g := range gear {
		out[g.Slot] = g.ItemID
	}
	return out
}

// armed names the slots of full (the band-20 entry) that hold a weapon, a
// shield or a wand: the ones a lower level must still fill.
func armed(full map[string]int, kinds map[int]itemKind) []string {
	var out []string
	for _, slot := range []string{"main_hand", "off_hand", "ranged"} {
		k, ok := kinds[full[slot]]
		if !ok {
			continue
		}
		if k.ClassID == kindWeapon || (k.ClassID == kindArmor && k.SubclassID == kindShield) {
			out = append(out, slot)
		}
	}
	return out
}

func TestNoSpecBelowTheLowestBandIsUnarmedOrBowless(t *testing.T) {
	build, err := leveling.ReadActiveBuild(realRoot)
	if err != nil {
		t.Fatal(err)
	}
	kinds := readItemKinds(t)
	for _, spec := range specs.All {
		gear, err := request.LoadLevelGear(realRoot, build, spec.Spec, "horde")
		if err != nil {
			t.Fatal(err)
		}
		for _, faction := range []string{"horde", "alliance"} {
			o := realOptions(spec.Spec)
			o.band, o.preset, o.faction, o.level = 20, "bare", faction, 20
			fullReq, err := buildRequest(o)
			if err != nil {
				t.Fatal(err)
			}
			full := slotItems(fullReq.Character.Gear)
			for level := levelFloor; level < lowestBand; level++ {
				o.level = level
				req, err := buildRequest(o)
				if err != nil {
					t.Fatalf("%s %s %d: %v", spec.Spec, faction, level, err)
				}
				worn := slotItems(req.Character.Gear)
				for _, g := range req.Character.Gear {
					if need := gear.RequiredLevel(g.ItemID); need > level {
						t.Errorf("%s %s %d wears %s %d, which needs %d", spec.Spec, faction, level, g.Slot, g.ItemID, need)
					}
				}
				for _, slot := range armed(full, kinds) {
					if worn[slot] == 0 {
						t.Errorf("%s %s %d: %s is empty (band 20 wears %d)", spec.Spec, faction, level, slot, full[slot])
					}
				}
				if spec.ClassSlug == "hunter" && !launchers[kinds[worn["ranged"]].SubclassID] {
					t.Errorf("%s %s %d: no bow, gun or crossbow (ranged %d)", spec.Spec, faction, level, worn["ranged"])
				}
			}
		}
	}
}

// A quest reward the client lists at required level 0 is gated by its quest:
// Forsaken Greataxe (251533) comes from The Wrath of Rath'mael (quest level
// 22), so a level-12 Retribution paladin cannot hold it.
func TestALevelBelowTheBandHoldsNoQuestRewardItCannotHaveYet(t *testing.T) {
	o := realOptions("paladin-retribution")
	o.band, o.preset, o.faction, o.level = 20, "bare", "horde", 12
	req, err := buildRequest(o)
	if err != nil {
		t.Fatal(err)
	}
	worn := slotItems(req.Character.Gear)
	if worn["main_hand"] == 251533 {
		t.Error("level 12 holds Forsaken Greataxe, a level-22 quest's reward")
	}
	if worn["main_hand"] == 0 {
		t.Error("level 12 holds no main hand")
	}
}

// A wand entry keeps a wand: a caster below 20 still has its Shoot.
func TestACasterBelowTheBandKeepsAWand(t *testing.T) {
	kinds := readItemKinds(t)
	o := realOptions("mage-fire")
	o.band, o.preset, o.faction, o.level = 20, "bare", "horde", 12
	req, err := buildRequest(o)
	if err != nil {
		t.Fatal(err)
	}
	ranged := slotItems(req.Character.Gear)["ranged"]
	if k := kinds[ranged]; k.ClassID != kindWeapon || k.SubclassID != subclassWand {
		t.Errorf("level-12 Fire mage ranged = %d (%+v), want a wand", ranged, k)
	}
}

// A refill is one the character's faction can earn: Deepblaze (279896) is the
// reward of an Alliance quest (Old Ironforge Incursion), so a level-13 Horde
// priest is never handed it.
func TestARefillIsOneTheFactionCanEarn(t *testing.T) {
	o := realOptions("priest-shadow")
	o.band, o.preset, o.faction, o.level = 20, "bare", "horde", 13
	req, err := buildRequest(o)
	if err != nil {
		t.Fatal(err)
	}
	worn := slotItems(req.Character.Gear)
	if worn["ranged"] == 279896 {
		t.Error("a level-13 Horde priest holds Deepblaze, an Alliance quest's reward")
	}
	if worn["ranged"] == 0 {
		t.Error("a level-13 Horde priest holds no wand")
	}
}
