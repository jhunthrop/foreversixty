package request

import (
	"encoding/json"
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

// The site runs a healer's list as written at band 60, so a ranked spell
// below the top rank a trainer teaches is cast at that lower rank. Every
// such rank must be one the list chose on purpose (its notes state the rank
// rule's picks); round 6 found Restoration shaman on Mana Tide Totem rank 3
// (17359), a rank below the Forever rank 4 (1325256) learned at 58.
func TestHealerListsCastTheTopTrainerRankUnlessChosen(t *testing.T) {
	const build = "../../data/builds/1.60.1.70707/spellranks.json"
	chosen := map[string]map[int]bool{ // spec -> spell id the rank rule picked
		"priest-holy":        {10927: true, 2060: true}, // round 18: the step down from the top Greater Heal
		"priest-discipline":  {10927: true, 2060: true},
		"paladin-holy":       {19941: true, 1311606: true, 19942: true},
		"druid-restoration":  {8910: true},
		"shaman-restoration": {},
	}
	var ranks struct {
		Classes map[string]map[string][]struct {
			ID, Level, Rank int
		} `json:"classes"`
	}
	readJSON(t, build, &ranks)
	var books struct{ Books []struct{ ID int } }
	readJSON(t, "../../data/curated/book-ranks.json", &books)
	book := map[int]bool{}
	for _, b := range books.Books {
		book[b.ID] = true
	}
	ids := regexp.MustCompile(`"spellId": (\d+)`)
	for spec, picked := range chosen {
		class, _, _ := strings.Cut(spec, "-")
		raw, err := os.ReadFile("../../data/curated/apl/" + spec + ".json")
		if err != nil {
			t.Fatal(err)
		}
		for _, m := range ids.FindAllStringSubmatch(string(raw), -1) {
			id, _ := strconv.Atoi(m[1])
			for name, family := range ranks.Classes[class] {
				mine, top := 0, 0
				for _, r := range family {
					if r.ID == id {
						mine = r.Rank
					}
					if r.Level <= 60 && !book[r.ID] && r.Rank > top {
						top = r.Rank
					}
				}
				if mine > 0 && mine < top && !picked[id] {
					t.Errorf("%s casts %s rank %d (%d), below the top trainer rank %d, and its notes do not choose it", spec, name, mine, id, top)
				}
			}
		}
	}
}

func readJSON(t *testing.T, path string, into any) {
	t.Helper()
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(raw, into); err != nil {
		t.Fatalf("%s: %v", path, err)
	}
}

// A healer list's note that names the rank its line pays for ("Rank 4 is
// the rank this profile pays for") must name the rank the line casts at
// band 60. Round 9 found Discipline's Prayer of Healing note still saying
// rank 1 after the line moved to the trainer's rank 4 (10961) on
// 2026-10-21.
func TestHealerListNotesNameTheRankTheyCast(t *testing.T) {
	var ranks struct {
		Classes map[string]map[string][]struct {
			ID, Level, Rank int
		} `json:"classes"`
	}
	readJSON(t, "../../data/builds/1.60.1.70707/spellranks.json", &ranks)
	named := regexp.MustCompile(`Rank (\d+) is the rank`)
	for _, spec := range []string{"priest-holy", "priest-discipline", "paladin-holy", "druid-restoration", "shaman-restoration"} {
		class, _, _ := strings.Cut(spec, "-")
		var list struct {
			Rotation struct {
				PriorityList []struct {
					Notes  string
					Action map[string]json.RawMessage
				} `json:"priorityList"`
			}
		}
		readJSON(t, "../../data/curated/apl/"+spec+".json", &list)
		for _, entry := range list.Rotation.PriorityList {
			m := named.FindStringSubmatch(entry.Notes)
			if m == nil {
				continue
			}
			want, _ := strconv.Atoi(m[1])
			var cast struct {
				SpellID struct {
					SpellID int `json:"spellId"`
				} `json:"spellId"`
			}
			if raw, ok := entry.Action["castSpell"]; !ok || json.Unmarshal(raw, &cast) != nil {
				t.Errorf("%s: a note names rank %d on a line that casts no spell", spec, want)
				continue
			}
			for name, family := range ranks.Classes[class] {
				for _, r := range family {
					if r.ID == cast.SpellID.SpellID && r.Rank != want {
						t.Errorf("%s: the %s line's note names rank %d but it casts rank %d (%d)", spec, name, want, r.Rank, r.ID)
					}
				}
			}
		}
	}
}

// assertOnlyRank checks that the rewritten list casts the family's rank and
// no other rank of it.
func assertOnlyRank(t *testing.T, rotation json.RawMessage, spec, class, preset, race string, level int, family map[int]int, want int) {
	t.Helper()
	out, err := RewriteRotationRanks(rotation, spec, class, preset, race, level)
	if err != nil {
		t.Fatal(err)
	}
	got := string(out)
	for rank, id := range family {
		has := strings.Contains(got, `"spellId":`+strconv.Itoa(id)) || strings.Contains(got, `"spellId": `+strconv.Itoa(id))
		if has != (rank == want) {
			t.Errorf("%s level %d %s: rank %d (%d) cast %v, want only rank %d", spec, level, preset, rank, id, has, want)
		}
	}
}

// Sim-accuracy round 15: the Healing Wave cast after Nature's Swiftness is
// the save, and Nature's Swiftness only takes the cast time away, so it is
// the highest trainer rank learned at every band (rank 9, 10396, at 60;
// rank 10 is a book rank). Round 14's rank-2 sweep (cf5fe63bb) moved this
// line to rank 2 with the tank filler: +1.0% effective HPS at raid 60 back,
// both factions, seeds 1 and 50001.
func TestRestorationShamanSwiftnessHealingWaveIsTheTopRank(t *testing.T) {
	_, swiftness := restorationHealingWaveLines(t)
	for _, c := range []struct {
		level  int
		preset string
		rank   int
	}{{44, BarePreset, 7}, {49, BarePreset, 8}, {50, BarePreset, 8}, {56, BarePreset, 9}, {60, BarePreset, 9}, {60, RaidPreset, 9}} {
		assertOnlyRank(t, swiftness, "shaman-restoration", "shaman", c.preset, "tauren", c.level, healingWaveRanks, c.rank)
	}
}

// Sim-accuracy round 15 wave 3: on Forever's level-gap rule a Healing Wave
// rank keeps its whole healing bonus until it is 17 levels behind, so the
// tank filler pays for a lower rank below 50. Round 16 measured every level,
// not only each band's floor, and added rank 2 and the ranks between: the
// filler casts rank 2 to 25 (+3.0% over rank 3 at 20), 3 to 36, 4 at 37, 5
// to 43 (+3.5% over rank 3 at 39), 6 to 56 (+5.3% over rank 5 at 49, +2.0%
// over rank 8 at 50), 7 at 57 and 58, 8 at 59, and the top rank at 60. Each
// fight still runs dry; seeds 1 and 50001. Round 17 re-measured the rows on
// the gated Nature's Swiftness save (round 16 picked 37 to 47 before the
// save needed the talent): rank 4 holds to 41 (+3.3% over rank 5 at 38) and
// rank 5 to 46 (+3.9% over rank 6 at 44). The rows are scoped to the
// filler's line, so the Nature's Swiftness save on the same id keeps its top
// rank (TestRestorationShamanSwiftnessHealingWaveIsTheTopRank).
// Round 19: with Chain Heal rank 1 through 59, rank 7 starts at 56 (+0.4% to
// +0.7% over rank 6). Round 24, on the talents a player has at each level:
// rank 1 at 20 and 21 (+0.6% to +4.1% over rank 2) and at 22 for a tauren,
// orc or troll, rank 3 at 23 and 24, rank 2 at 25 for the horde races (a
// dwarf casts rank 3), and rank 6 from 46 (+0.2% to +0.4% over rank 5). Round 29,
// on the new leveling order: rank 2 at 20 to 22 (TestRound29HealerRankRows).
func TestRestorationShamanFillerHealingWaveByBand(t *testing.T) {
	filler, _ := restorationHealingWaveLines(t)
	for _, c := range []struct {
		level  int
		preset string
		rank   int
	}{
		{20, BarePreset, 2}, {21, BarePreset, 2}, {22, BarePreset, 2}, {23, BarePreset, 3}, {24, BarePreset, 3},
		{25, BarePreset, 2}, {26, BarePreset, 3}, {30, BarePreset, 3}, {36, BarePreset, 3},
		{37, BarePreset, 6}, {38, BarePreset, 6}, {41, BarePreset, 5}, {42, BarePreset, 5}, {44, BarePreset, 6},
		{45, BarePreset, 6}, {46, BarePreset, 6}, {47, BarePreset, 6}, {50, BarePreset, 6}, {55, BarePreset, 6}, {56, BarePreset, 7}, {57, BarePreset, 7}, {58, BarePreset, 7}, {59, BarePreset, 8},
		{60, BarePreset, 9}, {60, RaidPreset, 9},
	} {
		assertOnlyRank(t, filler, "shaman-restoration", "shaman", c.preset, "tauren", c.level, healingWaveRanks, c.rank)
	}
}

var healingWaveRanks = map[int]int{1: 331, 2: 332, 3: 547, 4: 913, 5: 939, 6: 959, 7: 8005, 8: 10395, 9: 10396, 10: 25357}

// restorationHealingWaveLines splits the Restoration list's Healing Wave
// lines into two one-line rotations: the tank filler and the cast that
// spends Nature's Swiftness (16188). Each keeps its line's place in the list
// (earlier entries empty), so a rank choice scoped to lines still finds it.
func restorationHealingWaveLines(t *testing.T) (filler, swiftness json.RawMessage) {
	t.Helper()
	var list struct {
		Rotation struct {
			PriorityList []json.RawMessage `json:"priorityList"`
		} `json:"rotation"`
	}
	readJSON(t, "../../data/curated/apl/shaman-restoration.json", &list)
	wrap := func(index int, line json.RawMessage) json.RawMessage {
		entries := make([]json.RawMessage, index+1)
		for i := range index {
			entries[i] = json.RawMessage(`{}`)
		}
		entries[index] = line
		out, err := json.Marshal(map[string]any{"type": "TypeAPL", "priorityList": entries})
		if err != nil {
			t.Fatal(err)
		}
		return out
	}
	for index, line := range list.Rotation.PriorityList {
		var entry struct {
			Action struct {
				CastSpell *struct {
					SpellID struct {
						SpellID int `json:"spellId"`
					} `json:"spellId"`
				} `json:"castSpell"`
				Condition json.RawMessage `json:"condition"`
			} `json:"action"`
		}
		if err := json.Unmarshal(line, &entry); err != nil {
			t.Fatal(err)
		}
		cast := entry.Action.CastSpell
		if cast == nil || !isHealingWave(cast.SpellID.SpellID) {
			continue
		}
		if strings.Contains(string(entry.Action.Condition), `"auraIsActive"`) && strings.Contains(string(entry.Action.Condition), "16188") {
			swiftness = wrap(index, line)
		} else {
			filler = wrap(index, line)
		}
	}
	if filler == nil || swiftness == nil {
		t.Fatal("the Restoration list lacks the tank Healing Wave or the Nature's Swiftness Healing Wave")
	}
	return filler, swiftness
}

func isHealingWave(id int) bool {
	for _, wave := range healingWaveRanks {
		if wave == id {
			return true
		}
	}
	return false
}

// Round 14, at the 2.3 heal scale: Holy Shock is rank 4 (20930) at bare 60,
// +0.3% both factions, seeds 1 and 50001. Round 15 wave 3: under the
// level-gap rule rank 1 (learned at 30) keeps only 35% of its healing bonus
// at 60, and rank 4 wins the raid too, +0.4 to +0.6% both factions and seeds.
// The rotation now names rank 4 (20930); below 60 the top rank learned ties
// rank 1 (bands 40 and 50 within 0.1%), so no row holds it back.
func TestHolyPaladinHolyShockAtBareSixty(t *testing.T) {
	raw, err := os.ReadFile("../../data/curated/apl/paladin-holy.json")
	if err != nil {
		t.Fatal(err)
	}
	var rotation struct {
		Rotation json.RawMessage `json:"rotation"`
	}
	if err := json.Unmarshal(raw, &rotation); err != nil {
		t.Fatal(err)
	}
	shock := map[int]int{1: 1311606, 2: 20473, 3: 20929, 4: 20930}
	assertOnlyRank(t, rotation.Rotation, "paladin-holy", "paladin", BarePreset, "human", 60, shock, 4)
	assertOnlyRank(t, rotation.Rotation, "paladin-holy", "paladin", RaidPreset, "human", 60, shock, 4)
}

// Round 13: Prayer of Mending is Holy's cheapest healing (its line says so:
// "every time it is ready"), so its line is not held back by the mana pace
// the costlier heals follow. Gated, it sat out whole stretches of a fight
// the mana could not keep pace with: raid 60 +1.4 to +1.5% ungated, bare 60
// +22 to +23%, bare 40 and 50 +10 to +11%.
func TestHolyCastsPrayerOfMendingWheneverReady(t *testing.T) {
	const prayerOfMending = 1240827
	raw, err := os.ReadFile("../../data/curated/apl/priest-holy.json")
	if err != nil {
		t.Fatal(err)
	}
	var file struct {
		Rotation struct {
			PriorityList []json.RawMessage `json:"priorityList"`
		} `json:"rotation"`
	}
	if err := json.Unmarshal(raw, &file); err != nil {
		t.Fatal(err)
	}
	found := false
	for _, item := range file.Rotation.PriorityList {
		var line struct {
			Action struct {
				CastSpell *struct {
					SpellID struct {
						SpellID int `json:"spellId"`
					} `json:"spellId"`
				} `json:"castSpell"`
				Condition json.RawMessage `json:"condition"`
			} `json:"action"`
		}
		if err := json.Unmarshal(item, &line); err != nil {
			t.Fatal(err)
		}
		if line.Action.CastSpell == nil || line.Action.CastSpell.SpellID.SpellID != prayerOfMending {
			continue
		}
		found = true
		if strings.Contains(string(line.Action.Condition), "currentManaPercent") {
			t.Errorf("Prayer of Mending waits on the mana pace: %s", line.Action.Condition)
		}
	}
	if !found {
		t.Fatal("the Holy list does not cast Prayer of Mending (1240827)")
	}
}

// The party's Rejuvenation is the druid's cheapest healing, so it does not
// wait on the mana pace (round 13).
func TestRestorationDruidPartyRejuvenationIsNotPaced(t *testing.T) {
	raw, err := os.ReadFile("../../data/curated/apl/druid-restoration.json")
	if err != nil {
		t.Fatal(err)
	}
	var file struct {
		Rotation struct {
			PriorityList []struct {
				Action struct {
					CastSpell *struct {
						SpellID struct {
							SpellID int `json:"spellId"`
						} `json:"spellId"`
						Target *struct {
							Index int `json:"index"`
						} `json:"target"`
					} `json:"castSpell"`
					Condition json.RawMessage `json:"condition"`
				} `json:"action"`
			} `json:"priorityList"`
		} `json:"rotation"`
	}
	if err := json.Unmarshal(raw, &file); err != nil {
		t.Fatal(err)
	}
	party := 0
	for _, line := range file.Rotation.PriorityList {
		cast := line.Action.CastSpell
		if cast == nil || cast.SpellID.SpellID != 8910 || cast.Target == nil || cast.Target.Index > 4 {
			continue
		}
		party++
		if strings.Contains(string(line.Action.Condition), "currentManaPercent") {
			t.Errorf("the party Rejuvenation on player %d waits on the mana pace", cast.Target.Index)
		}
	}
	if party != 5 {
		t.Errorf("%d party Rejuvenation rank 7 lines, want 5 (players 0 to 4)", party)
	}
}

// Round 14 (regcast): Holy's Heal rank 1 row at 50 to 59 (round 11) was
// adopted while Prayer of Mending waited on the mana pace. Once Prayer of
// Mending runs whenever it is ready (round 13), the list's own rank 4 heals
// more at band 50 on the published build and on the ladder build: effective
// HPS +12% horde and alliance, seeds 1 and 50001, and the fight still runs
// dry. So Heal at 50 to 59 casts the highest rank learned up to the list's
// rank 4 (6064), as its note says, for every race.
func TestHolyHealAtFiftyCastsTheListsRankFour(t *testing.T) {
	const (
		healRank1 = 2054
		healRank4 = 6064
	)
	raw, err := os.ReadFile("../../data/curated/apl/priest-holy.json")
	if err != nil {
		t.Fatal(err)
	}
	var file struct {
		Rotation json.RawMessage `json:"rotation"`
	}
	if err := json.Unmarshal(raw, &file); err != nil {
		t.Fatal(err)
	}
	cast := regexp.MustCompile(`"spellId":\s*(\d+)`)
	for _, level := range []int{50, 55, 59} {
		for _, race := range []string{"human", "dwarf", "night-elf", "gnome", "undead", "troll"} {
			out, err := RewriteRotationRanks(file.Rotation, "priest-holy", "priest", BarePreset, race, level)
			if err != nil {
				t.Fatal(err)
			}
			ids := map[int]bool{}
			for _, m := range cast.FindAllStringSubmatch(string(out), -1) {
				id, _ := strconv.Atoi(m[1])
				ids[id] = true
			}
			if ids[healRank1] {
				t.Errorf("level %d %s: Heal casts rank 1 (%d)", level, race, healRank1)
			}
			if !ids[healRank4] {
				t.Errorf("level %d %s: no Heal rank 4 (%d)", level, race, healRank4)
			}
		}
	}
}

// Round 14: the Holy priest's tank Flash Heal was rank 2 (9472) at raid 60,
// round 9's pick for a list that then ran dry at 137 s. Later rounds raised
// Holy's mana (Prayer of Mending ungated, Litany of Light), and the list
// ended raid fights with mana to spare (lasts about 440 s at the 1.85 heal
// scale, 530 s at 2.3) while the tank sat under 35%. Rank 7 (10917) heals
// +10.3 to +10.9% at raid 60, both factions, seeds 1 and 50001. At every
// published band below 60 the line casts what its rank-choice rows made it
// cast before (rank 1 at 20, 2 at 30, 4 at 40, 6 at 50, 7 at bare 60), which
// is now the highest rank learned, so the line needs no rows.
func TestHolyTankFlashHealIsTheTopRankInARaid(t *testing.T) {
	raw, err := os.ReadFile("../../data/curated/apl/priest-holy.json")
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(raw), `"spellId": 9472`) {
		t.Fatal("the Holy list still authors Flash Heal rank 2 (9472)")
	}
	var file struct {
		Rotation json.RawMessage `json:"rotation"`
	}
	if err := json.Unmarshal(raw, &file); err != nil {
		t.Fatal(err)
	}
	flash := map[int]int{1: 2061, 2: 9472, 3: 9473, 4: 9474, 5: 10915, 6: 10916, 7: 10917}
	cases := []struct {
		level  int
		preset string
		rank   int
	}{
		{20, BarePreset, 1}, {30, BarePreset, 2}, {40, BarePreset, 4}, {50, BarePreset, 6},
		{60, BarePreset, 7}, {60, RaidPreset, 7},
	}
	for _, c := range cases {
		for _, race := range []string{"gnome", "troll"} {
			out, err := RewriteRotationRanks(file.Rotation, "priest-holy", "priest", c.preset, race, c.level)
			if err != nil {
				t.Fatal(err)
			}
			got := string(out)
			for rank, id := range flash {
				has := strings.Contains(got, `"spellId":`+strconv.Itoa(id)) || strings.Contains(got, `"spellId": `+strconv.Itoa(id))
				if rank == c.rank && !has {
					t.Errorf("level %d %s %s: no Flash Heal rank %d (%d)", c.level, c.preset, race, rank, id)
				}
				if rank != c.rank && rank != 7 && has {
					t.Errorf("level %d %s %s: Flash Heal rank %d (%d) cast, want rank %d", c.level, c.preset, race, rank, id, c.rank)
				}
			}
		}
	}
}

// Sim-accuracy round 15: 70755 made Riptide a heal over time only (no
// direct heal, 12 s cooldown, 150/180/230 mana), and the bare-60 row that
// held it at rank 2 now loses to rank 3 by 1.5% effective HPS, both
// factions, seeds 1 and 50001. Every band casts the highest rank learned.
func TestRestorationShamanRiptideIsTheTopRank(t *testing.T) {
	var list struct {
		Rotation json.RawMessage `json:"rotation"`
	}
	readJSON(t, "../../data/curated/apl/shaman-restoration.json", &list)
	riptide := map[int]int{1: 408521, 2: 1239242, 3: 1239243}
	for _, c := range []struct {
		level  int
		preset string
		rank   int
	}{{40, BarePreset, 1}, {50, BarePreset, 2}, {60, BarePreset, 3}, {60, RaidPreset, 3}} {
		assertOnlyRank(t, list.Rotation, "shaman-restoration", "shaman", c.preset, "tauren", c.level, riptide, c.rank)
	}
}

// Sim-accuracy round 15: once 70755 made Riptide a 230-mana heal over time
// (about 4.6 healed per mana against Lesser Healing Wave's 3.6), keeping it
// on the tank comes before the tank's Lesser Healing Wave and Healing Wave
// lines, right after the Nature's Swiftness save: +0.8% effective HPS at
// raid 60 (mana lasts 139 to 160 s), bare 50 +1.4%, bare 60 +0.8%, both
// factions, seeds 1 and 50001.
func TestRestorationShamanRiptidesTheTankBeforeItsWaves(t *testing.T) {
	var list struct {
		Rotation struct {
			PriorityList []json.RawMessage `json:"priorityList"`
		} `json:"rotation"`
	}
	readJSON(t, "../../data/curated/apl/shaman-restoration.json", &list)
	first := map[string]int{}
	for i, line := range list.Rotation.PriorityList {
		var entry struct {
			Action struct {
				CastSpell *struct {
					SpellID struct {
						SpellID int `json:"spellId"`
					} `json:"spellId"`
					Target struct {
						Index int `json:"index"`
					} `json:"target"`
				} `json:"castSpell"`
			} `json:"action"`
		}
		if err := json.Unmarshal(line, &entry); err != nil {
			t.Fatal(err)
		}
		cast := entry.Action.CastSpell
		if cast == nil || cast.Target.Index != 5 {
			continue
		}
		key := strconv.Itoa(cast.SpellID.SpellID)
		if _, seen := first[key]; !seen {
			first[key] = i
		}
	}
	riptide, ok := first["1239243"]
	if !ok {
		t.Fatal("no tank Riptide line")
	}
	for _, wave := range []string{"10468", "332"} {
		if at, ok := first[wave]; ok && at < riptide {
			t.Errorf("the tank's %s line (priority %d) comes before its Riptide (priority %d)", wave, at, riptide)
		}
	}
}

// Sim-accuracy round 16: the Nature's Swiftness save asked only whether the
// aura was active, and the engine drops a condition on an aura the character
// does not have, so below 44 (no Nature's Swiftness yet) the line fired as a
// plain Healing Wave whenever the tank fell under 25%: a heal no leveling
// shaman's rotation has (+5.5% effective HPS at bare 40 on that phantom line).
// The line now also needs the talent known.
func TestRestorationShamanSwiftnessSaveNeedsTheTalent(t *testing.T) {
	_, swiftness := restorationHealingWaveLines(t)
	if !strings.Contains(strings.ReplaceAll(string(swiftness), " ", ""), `"auraIsKnown":{"auraId":{"spellId":16188}}`) {
		t.Errorf("the Nature's Swiftness Healing Wave does not ask that Nature's Swiftness is known: %s", swiftness)
	}
}
