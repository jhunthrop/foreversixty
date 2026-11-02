package request

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/internal/spellranks"
	"github.com/jhunthrop/foreversixty/sim/specs"
)

const (
	curatedRankChoices = "../../data/curated/rank-choices.json"
	// activeSpellRanks is the table spellranks embeds (`make simdb` copies it).
	activeSpellRanks = "../internal/spellranks/spellranks.json"
)

func TestEmbeddedRankChoicesMatchCurated(t *testing.T) {
	curated, err := os.ReadFile(filepath.FromSlash(curatedRankChoices))
	if err != nil {
		t.Fatalf("reading the curated rank choices: %v", err)
	}
	if string(curated) != string(rankChoicesJSON) {
		t.Fatal("sim/request/rank-choices.json differs from data/curated/rank-choices.json; run `make rank-choices-sync`")
	}
	if _, err := embeddedRankChoices(); err != nil {
		t.Fatalf("the embedded rank choices do not load: %v", err)
	}
}

// Every row is a real downrank the rotation can cast: a spec this build
// sims, an id its rotation names, the spell its name says, a band below 60
// and a rank learned by the band's first level and below the rank the
// default rule would cast there.
func TestCuratedRankChoicesAreCastableDownranks(t *testing.T) {
	var file rankChoicesFile
	readJSON(t, curatedRankChoices, &file)
	var ranks struct {
		Classes map[string]map[string][]struct{ ID int32 } `json:"classes"`
	}
	readJSON(t, activeSpellRanks, &ranks)
	for spec, spells := range file.Specs {
		s, ok := specs.ByKey[spec]
		if !ok {
			t.Errorf("%s is not a spec", spec)
			continue
		}
		raw, err := aplFS.ReadFile("apl/" + spec + ".apl.json")
		if err != nil {
			t.Fatal(err)
		}
		for _, sp := range spells {
			if !strings.Contains(string(raw), `"spellId": `+strconv.Itoa(int(sp.Authored))) {
				t.Errorf("%s: its rotation does not name %d", spec, sp.Authored)
			}
			named := false
			for _, row := range ranks.Classes[s.ClassSlug][sp.Spell] {
				named = named || row.ID == sp.Authored
			}
			if !named {
				t.Errorf("%s: %d is not a rank of %q", spec, sp.Authored, sp.Spell)
			}
			for _, row := range sp.Bands {
				if row.FromLevel < 1 || row.FromLevel >= api.MaxLevel {
					t.Errorf("%s %s: a row from level %d; level 60 runs the rotation as written", spec, sp.Spell, row.FromLevel)
					continue
				}
				chosen, learned := spellranks.SpellIDOfRank(s.ClassSlug, sp.Authored, row.Rank, row.FromLevel)
				if !learned {
					t.Errorf("%s %s rank %d is not learned at %d", spec, sp.Spell, row.Rank, row.FromLevel)
					continue
				}
				if def, _ := spellranks.LearnedSpellID(s.ClassSlug, sp.Authored, row.FromLevel); def == chosen {
					t.Errorf("%s %s rank %d from %d is what the rotation casts anyway", spec, sp.Spell, row.Rank, row.FromLevel)
				}
			}
		}
	}
}

func TestParseRankChoicesRefusesDuplicates(t *testing.T) {
	for _, raw := range []string{
		`{"specs":{"x":[{"authored":1,"bands":[]},{"authored":1,"bands":[]}]}}`,
		`{"specs":{"x":[{"authored":1,"bands":[{"from_level":40,"rank":2},{"from_level":40,"rank":3}]}]}}`,
		`not json`,
	} {
		if _, err := parseRankChoices([]byte(raw)); err == nil {
			t.Errorf("accepted %s", raw)
		}
	}
	got, err := parseRankChoices([]byte(`{"specs":{"x":[{"authored":7,"bands":[{"from_level":50,"rank":3},{"from_level":40,"rank":2}]}]}}`))
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct{ level, rank int }{{39, 0}, {40, 2}, {49, 2}, {50, 3}} {
		rank, ok := got["x"].rankAt(7, c.level)
		if ok != (c.rank != 0) || rank != c.rank {
			t.Errorf("rankAt(7, %d) = (%d, %v), want %d", c.level, rank, ok, c.rank)
		}
	}
}
