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
// its authored rank 9.
func TestElementalRaidCastsLightningBoltRankTen(t *testing.T) {
	var list struct {
		Rotation json.RawMessage `json:"rotation"`
	}
	readJSON(t, "../../data/curated/apl/shaman-elemental.json", &list)
	bolt := map[int]int{4: 915, 7: 10391, 9: 15207, 10: 15208}
	for _, c := range []struct {
		preset string
		want   map[int]bool
	}{{RaidPreset, map[int]bool{10: true}}, {BarePreset, map[int]bool{9: true, 10: true}}} {
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
// (+3.5 to +4.2%) and 50 (+1.4 to +1.7%); rank 3 wins at 40 (+0.8 to
// +0.9% over rank 1, ahead of rank 4), both factions, seeds 1 and 50001.
func TestElementalFlameShockRanksByBand(t *testing.T) {
	var list struct {
		Rotation json.RawMessage `json:"rotation"`
	}
	readJSON(t, "../../data/curated/apl/shaman-elemental.json", &list)
	shock := map[int]int{1: 8050, 2: 8052, 3: 8053, 4: 10447, 5: 10448}
	for level, want := range map[int]int{30: 3, 39: 3, 40: 3, 49: 3, 50: 4, 55: 5} {
		assertOnlyRank(t, list.Rotation, "shaman-elemental", "shaman", BarePreset, "troll", level, shock, want)
	}
}
