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
		"priest-holy":        {6075: true, 9472: true},
		"priest-discipline":  {6075: true, 3747: true, 2053: true},
		"paladin-holy":       {19941: true, 1311606: true, 19750: true},
		"druid-restoration":  {774: true},
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

// Discipline's party filler is Lesser Heal rank 3 (2053, level 10), not
// Heal rank 1 (2054). Once heals learned below 20 kept their whole
// coefficient (engine 074edfadf, sim-accuracy round 11), rank 3's 0.714 at
// 75 mana out-healed Heal rank 1's 0.857 at 155 for a list that runs dry:
// +1.8 to +2.1% at raid 60 and +7.8 to +8.8% at bare 60, both factions,
// seeds 1 and 50001, 10,000 iterations. No list line names Heal rank 1, and
// so no rank-choice row may either.
func TestDisciplineFillsWithLesserHealRankThree(t *testing.T) {
	raw, err := os.ReadFile("../../data/curated/apl/priest-discipline.json")
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(raw), `"spellId": 2054`) {
		t.Error("the Discipline list still casts Heal rank 1 (2054)")
	}
	if !strings.Contains(string(raw), `"spellId": 2053`) {
		t.Error("the Discipline list does not cast Lesser Heal rank 3 (2053)")
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
