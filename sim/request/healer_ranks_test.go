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
		"druid-restoration":  {2090: true},
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
