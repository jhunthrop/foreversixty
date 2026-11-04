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
		"priest-holy":        {139: true, 2054: true},
		"priest-discipline":  {139: true, 3747: true, 2050: true},
		"paladin-holy":       {19941: true, 1311606: true, 19750: true},
		"druid-restoration":  {774: true},
		"shaman-restoration": {332: true},
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

// Discipline's party filler is Lesser Heal, not Heal rank 1 (2054). Once
// heals learned below 20 kept their whole coefficient (engine 074edfadf,
// sim-accuracy round 11), Lesser Heal rank 3's 0.714 at 75 mana out-healed
// Heal rank 1's 0.857 at 155 for a list that runs dry: +1.8 to +2.1% at raid
// 60 and +7.8 to +8.8% at bare 60. Round 14, at the 2.3 heal scale: rank 1
// (2050, level 1) heals more again, raid 60 +0.3 to +0.4% and bare 60 +4.3
// to +4.5%, both factions, seeds 1 and 50001. Below 60 the line keeps what
// it cast: rank 2 from 4 to 9 and rank 3 from 10 to 59. No list line names
// Heal rank 1, and so no rank-choice row may either.
func TestDisciplineFillsWithLesserHealRankOne(t *testing.T) {
	raw, err := os.ReadFile("../../data/curated/apl/priest-discipline.json")
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(raw), `"spellId": 2054`) {
		t.Error("the Discipline list still casts Heal rank 1 (2054)")
	}
	if strings.Contains(string(raw), `"spellId": 2053`) || !strings.Contains(string(raw), `"spellId": 2050`) {
		t.Error("the Discipline list does not author Lesser Heal rank 1 (2050)")
	}
	choices, err := os.ReadFile("../../data/curated/rank-choices.json")
	if err != nil {
		t.Fatal(err)
	}
	var file rankChoicesFile
	if err := json.Unmarshal(choices, &file); err != nil {
		t.Fatal(err)
	}
	for _, sp := range file.Specs["priest-discipline"] {
		if sp.Authored == 2054 {
			t.Error("rank-choices.json still has a Discipline row for Heal rank 1 (2054)")
		}
	}
	var rotation struct {
		Rotation json.RawMessage `json:"rotation"`
	}
	if err := json.Unmarshal(raw, &rotation); err != nil {
		t.Fatal(err)
	}
	lesser := map[int]int{1: 2050, 2: 2052, 3: 2053}
	for _, c := range []struct {
		level  int
		preset string
		rank   int
	}{{4, BarePreset, 2}, {9, BarePreset, 2}, {10, BarePreset, 3}, {20, BarePreset, 3}, {50, BarePreset, 3}, {59, BarePreset, 3}, {60, BarePreset, 1}, {60, RaidPreset, 1}} {
		assertOnlyRank(t, rotation.Rotation, "priest-discipline", "priest", c.preset, "undead", c.level, lesser, c.rank)
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

// Round 14, at the 2.3 heal scale: the Restoration shaman's tank Healing
// Wave is rank 2 (332) at raid 60 as it already was from 40 to 59 and at
// bare 60: +0.7 to +0.8%, both factions, seeds 1 and 50001, and the fight
// still runs dry (162 to 146 s). Below 40 the line keeps the highest rank
// learned.
func TestRestorationShamanHealingWaveRanksByBand(t *testing.T) {
	filler, _ := restorationHealingWaveLines(t)
	for _, c := range []struct {
		level  int
		preset string
		rank   int
	}{{1, BarePreset, 1}, {6, BarePreset, 2}, {12, BarePreset, 3}, {18, BarePreset, 4}, {20, BarePreset, 4}, {24, BarePreset, 5}, {30, BarePreset, 5},
		{32, BarePreset, 6}, {39, BarePreset, 6}, {40, BarePreset, 2}, {50, BarePreset, 2}, {60, BarePreset, 2}, {60, RaidPreset, 2}} {
		assertOnlyRank(t, filler, "shaman-restoration", "shaman", c.preset, "tauren", c.level, healingWaveRanks, c.rank)
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
	}{{20, BarePreset, 4}, {30, BarePreset, 5}, {40, BarePreset, 7}, {50, BarePreset, 8}, {56, BarePreset, 9}, {60, BarePreset, 9}, {60, RaidPreset, 9}} {
		assertOnlyRank(t, swiftness, "shaman-restoration", "shaman", c.preset, "tauren", c.level, healingWaveRanks, c.rank)
	}
}

var healingWaveRanks = map[int]int{1: 331, 2: 332, 3: 547, 4: 913, 5: 939, 6: 959, 7: 8005, 8: 10395, 9: 10396, 10: 25357}

// restorationHealingWaveLines splits the Restoration list's Healing Wave
// lines into two one-line rotations: the tank filler and the cast that
// spends Nature's Swiftness (16188).
func restorationHealingWaveLines(t *testing.T) (filler, swiftness json.RawMessage) {
	t.Helper()
	var list struct {
		Rotation struct {
			PriorityList []json.RawMessage `json:"priorityList"`
		} `json:"rotation"`
	}
	readJSON(t, "../../data/curated/apl/shaman-restoration.json", &list)
	wrap := func(line json.RawMessage) json.RawMessage {
		out, err := json.Marshal(map[string]any{"type": "TypeAPL", "priorityList": []json.RawMessage{line}})
		if err != nil {
			t.Fatal(err)
		}
		return out
	}
	for _, line := range list.Rotation.PriorityList {
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
			swiftness = wrap(line)
		} else {
			filler = wrap(line)
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
// +0.3% both factions, seeds 1 and 50001; the raid keeps rank 1.
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
	assertOnlyRank(t, rotation.Rotation, "paladin-holy", "paladin", RaidPreset, "human", 60, shock, 1)
}

// Round 13: at raid 60 the Restoration druid's mana outlasted the fight by
// more than half again with Rejuvenation rank 4 on everyone, and the party's
// Rejuvenation waited on the mana pace the costly heals follow. The tank's
// Rejuvenation is rank 10 (9841) in a raid; the party's is rank 1 (774),
// the cheapest healing a druid has (heals learned below 20 keep their whole
// coefficient in Forever, round 11 C1), cast whenever someone is hurt. The
// tank's line keeps what it cast before at every other band and at bare 60:
// rank 4 where it is learned, rank 1 at 50 to 59, and below 28 the highest
// rank learned up to rank 4.
func TestRestorationDruidRejuvenationRanksByBand(t *testing.T) {
	const (
		tankRejuv  = 9841
		partyRejuv = 774
	)
	raw, err := os.ReadFile("../../data/curated/apl/druid-restoration.json")
	if err != nil {
		t.Fatal(err)
	}
	var file struct {
		Rotation json.RawMessage `json:"rotation"`
	}
	if err := json.Unmarshal(raw, &file); err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(raw), `"spellId": 2090`) {
		t.Fatal("the Restoration list still authors Rejuvenation rank 4 (2090)")
	}
	rejuv := map[int]int{1: 774, 2: 1058, 3: 1430, 4: 2090, 10: 9841}
	cases := []struct {
		level       int
		preset      string
		tank, party int
	}{
		{10, RaidPreset, 2, 1}, {16, RaidPreset, 3, 1}, {21, RaidPreset, 3, 1}, {22, RaidPreset, 4, 1},
		{30, RaidPreset, 4, 1}, {45, RaidPreset, 4, 1}, {50, RaidPreset, 1, 1}, {59, RaidPreset, 1, 1},
		{60, BarePreset, 4, 1}, {60, RaidPreset, 10, 1},
	}
	for _, c := range cases {
		for _, race := range []string{"night-elf", "tauren"} {
			out, err := RewriteRotationRanks(file.Rotation, "druid-restoration", "druid", c.preset, race, c.level)
			if err != nil {
				t.Fatal(err)
			}
			got := string(out)
			for _, want := range []struct{ authored, rank int }{{tankRejuv, c.tank}, {partyRejuv, c.party}} {
				id := rejuv[want.rank]
				if !strings.Contains(got, `"spellId":`+strconv.Itoa(id)) && !strings.Contains(got, `"spellId": `+strconv.Itoa(id)) {
					t.Errorf("level %d %s %s: no Rejuvenation rank %d (%d) for the line authored %d", c.level, c.preset, race, want.rank, id, want.authored)
				}
			}
		}
	}
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
		if cast == nil || cast.SpellID.SpellID != 774 || cast.Target == nil || cast.Target.Index > 4 {
			continue
		}
		party++
		if strings.Contains(string(line.Action.Condition), "currentManaPercent") {
			t.Errorf("the party Rejuvenation on player %d waits on the mana pace", cast.Target.Index)
		}
	}
	if party != 5 {
		t.Errorf("%d party Rejuvenation rank 1 lines, want 5 (players 0 to 4)", party)
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

// Round 14: both priests' Renew is rank 1 (139, level 8) at 60 as it already
// was at every band below (the rank-choice rows of round 9): raid 60 +0.9 to
// +1.2% for both, bare 60 +0.2 to +1.6% Holy and +4.3% Discipline, both
// factions, seeds 1 and 50001. Heals learned below 20 keep their whole
// coefficient in Forever (round 11 C1), so rank 1 heals the most per mana.
func TestPriestsCastRenewRankOne(t *testing.T) {
	const renewOne, renewThree = 139, 6075
	for _, spec := range []string{"priest-holy", "priest-discipline"} {
		raw, err := os.ReadFile("../../data/curated/apl/" + spec + ".json")
		if err != nil {
			t.Fatal(err)
		}
		if strings.Contains(string(raw), `"spellId": `+strconv.Itoa(renewThree)) {
			t.Errorf("%s still authors Renew rank 3 (%d)", spec, renewThree)
		}
		if !strings.Contains(string(raw), `"spellId": `+strconv.Itoa(renewOne)) {
			t.Errorf("%s does not cast Renew rank 1 (%d)", spec, renewOne)
		}
	}
}

// Round 14: Holy's Heal is rank 1 (2054, level 16) in a raid, +2.0% at raid
// 60 on top of Renew rank 1, both factions, seeds 1 and 50001. Bare it loses
// (bare 30 -3.6%, bare 40 -1.5 to -2.6%, bare 60 -5.2 to -5.4%), so every
// bare band casts what it cast before: the highest rank learned up to rank
// 4 to 59 (rank 4 from 50 to 59 since regcast's 9dba07e78) and rank 4 at
// bare 60.
func TestHolyHealRanksByBand(t *testing.T) {
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
	heal := map[int]int{1: 2054, 2: 2055, 3: 6063, 4: 6064}
	cases := []struct {
		level  int
		preset string
		rank   int
	}{
		{20, BarePreset, 1}, {22, BarePreset, 2}, {27, BarePreset, 2}, {28, BarePreset, 3}, {30, BarePreset, 3},
		{34, BarePreset, 4}, {40, BarePreset, 4}, {49, BarePreset, 4}, {50, BarePreset, 4}, {59, BarePreset, 4},
		{60, BarePreset, 4}, {60, RaidPreset, 1},
	}
	for _, c := range cases {
		out, err := RewriteRotationRanks(file.Rotation, "priest-holy", "priest", c.preset, "troll", c.level)
		if err != nil {
			t.Fatal(err)
		}
		got := string(out)
		for rank, id := range heal {
			has := strings.Contains(got, `"spellId":`+strconv.Itoa(id)) || strings.Contains(got, `"spellId": `+strconv.Itoa(id))
			if has != (rank == c.rank) {
				t.Errorf("level %d %s: Heal rank %d (%d) cast %v, want only rank %d", c.level, c.preset, rank, id, has, c.rank)
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
