package request

import (
	"encoding/json"
	"strconv"
	"strings"
	"testing"
)

// Sim-accuracy round 15: client 1.60.1.70755 raised Lightning Bolt rank 10
// to 353 base against rank 9's 295 (+20% for 16% more mana; +10% on
// 70753), so the raid's mana pacing onto rank 9 costs damage: Elemental at
// raid 60 casting rank 10 alone is +1.0 to +1.2%, seeds 1 and 50001, both
// factions, out of mana 0.7 to 0.9 s of 180. At bare 60 the top line's
// rank 7 row is gone for the same reason (rank 10 with the rank 4 row is
// +0.33 to +0.38%, both factions, seeds 1 and 50001); bare 50 ties and
// keeps rank 7. Round 15 wave 2 (Forever's level-gap rule): the filler's
// rank-4 row at bare 60 is gone too, since rank 4 (learned at 20) keeps no
// spell power at 60 (+28% without it, both factions), so the filler casts
// its authored rank 9. Round 15 wave 3 re-swept every band on that rule:
// the filler casts rank 8 (learned at 44, full bonus at 60) at bare 60,
// +3.0 to +3.3% over rank 9, both factions, seeds 1 and 50001.
func TestElementalRaidCastsLightningBoltRankTen(t *testing.T) {
	var list struct {
		Rotation json.RawMessage `json:"rotation"`
	}
	readJSON(t, "../../data/curated/apl/shaman-elemental.json", &list)
	bolt := map[int]int{4: 915, 7: 10391, 8: 10392, 9: 15207, 10: 15208}
	for _, c := range []struct {
		preset string
		want   map[int]bool
	}{{RaidPreset, map[int]bool{10: true}}, {BarePreset, map[int]bool{8: true, 10: true}}} {
		out, err := RewriteRotationRanks(list.Rotation, "shaman-elemental", "shaman", c.preset, "troll", 60)
		if err != nil {
			t.Fatal(err)
		}
		for rank, id := range bolt {
			if has := strings.Contains(string(out), `"spellId":`+strconv.Itoa(id)); has != c.want[rank] {
				t.Errorf("%s 60: Lightning Bolt rank %d (%d) cast %v, want %v", c.preset, rank, id, has, c.want[rank])
			}
		}
	}
}

// Sim-accuracy round 15: 70755 raised every Flame Shock rank, so the
// rank 1 and 2 rows judged on 70753 lose. The top rank learned wins at 30
// (+3.5 to +4.2%) and 50 (+1.4 to +1.7%). Round 16: with the filler
// casting Lightning Bolt rank 5 from 39, rank 3 at 40 to 49 loses to rank 4
// at every level measured (40 +0.5 to +0.6%, 45 +1.0%), so the top rank
// learned casts there too.
func TestElementalFlameShockRanksByBand(t *testing.T) {
	var list struct {
		Rotation json.RawMessage `json:"rotation"`
	}
	readJSON(t, "../../data/curated/apl/shaman-elemental.json", &list)
	shock := map[int]int{1: 8050, 2: 8052, 3: 8053, 4: 10447, 5: 10448}
	for level, want := range map[int]int{30: 3, 39: 3, 40: 4, 49: 4, 50: 4, 55: 5} {
		assertOnlyRank(t, list.Rotation, "shaman-elemental", "shaman", BarePreset, "troll", level, shock, want)
	}
}

// Sim-accuracy round 15 wave 3: on Forever's level-gap rule (a rank keeps
// its whole spell power until it is 17 levels behind) the filler line's
// best rank moves to rank 3 at 30 and rank 6 at 50. Round 16 measured the
// rows at every level of their bands, not only the floor: a rank chosen at
// 30 is 8 levels deeper in the gap by 39, and the best rank climbs with the
// level. The filler (line 9, authored at rank 10 so a raid casts it as
// written) casts rank 3 to 33, 4 to 38, 5 to 44, 6 to 50, 7 to 56 and 8 from
// 57 (bare 60 too): at 39 rank 5 is +11.8% over rank 3, at 49 rank 6 +29.7%
// over rank 4, at 59 rank 8 +19.6% over rank 6, and at 40 rank 5 +0.5% over
// rank 4 (seeds 1 and 50001). The top-rank line casts the highest rank
// learned at each level.
func TestElementalLightningBoltRanksByBand(t *testing.T) {
	var list struct {
		Rotation json.RawMessage `json:"rotation"`
	}
	readJSON(t, "../../data/curated/apl/shaman-elemental.json", &list)
	bolt := map[int]int{3: 548, 4: 915, 5: 943, 6: 6041, 7: 10391, 8: 10392, 9: 15207, 10: 15208}
	for _, c := range []struct {
		level int
		want  map[int]bool
	}{
		{30, map[int]bool{3: true, 5: true}}, {33, map[int]bool{3: true, 6: true}},
		{34, map[int]bool{4: true, 6: true}}, {38, map[int]bool{4: true, 7: true}},
		{39, map[int]bool{5: true, 7: true}}, {40, map[int]bool{5: true, 7: true}}, {44, map[int]bool{5: true, 8: true}},
		{45, map[int]bool{6: true, 8: true}}, {50, map[int]bool{6: true, 9: true}},
		{51, map[int]bool{7: true, 9: true}}, {56, map[int]bool{7: true, 10: true}},
		{57, map[int]bool{8: true, 10: true}}, {59, map[int]bool{8: true, 10: true}},
	} {
		out, err := RewriteRotationRanks(list.Rotation, "shaman-elemental", "shaman", BarePreset, "troll", c.level)
		if err != nil {
			t.Fatal(err)
		}
		for rank, id := range bolt {
			if has := strings.Contains(string(out), `"spellId":`+strconv.Itoa(id)); has != c.want[rank] {
				t.Errorf("bare %d: Lightning Bolt rank %d (%d) cast %v, want %v", c.level, rank, id, has, c.want[rank])
			}
		}
	}
}
