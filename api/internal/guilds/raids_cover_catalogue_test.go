package guilds

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"testing"
)

const (
	raidCoverPath = "../../../data/curated/raid-cover.json"
	specsPath     = "../../../data/curated/specs.json"
	bookRanksPath = "../../../data/curated/book-ranks.json"
)

func coverBuildPath(file string) string {
	return filepath.Join("../../../data/builds", raidCoverBuild, file)
}

func TestRaidCoverCatalogueIsRaidCoverJSON(t *testing.T) {
	var file struct {
		Entries []coverEntry `json:"entries"`
	}
	readJSON(t, raidCoverPath, &file)
	got := make([]coverEntry, len(raidCoverCatalogue))
	for i, e := range raidCoverCatalogue {
		e.Icon = ""
		got[i] = e
	}
	if !reflect.DeepEqual(got, file.Entries) {
		t.Fatal("raids_cover_catalogue.go drifted from raid-cover.json; run go generate ./internal/guilds/")
	}
}

func TestRaidCoverRowsAreWellFormed(t *testing.T) {
	keys := map[string]bool{}
	poolProviders := map[string][]coverProvider{}
	for _, e := range raidCoverCatalogue {
		if keys[e.Key] || e.Name == "" || e.Icon == "" || e.Need.Count < 1 || len(e.Providers) == 0 {
			t.Errorf("%s: duplicate key or empty field: %+v", e.Key, e)
		}
		keys[e.Key] = true
		if e.Kind != "buff" && e.Kind != "debuff" {
			t.Errorf("%s: kind %q", e.Key, e.Kind)
		}
		if e.Need.Per != coverPerRaid && e.Need.Per != coverPerMeleeGroup && e.Need.Per != coverPerCasterGroup {
			t.Errorf("%s: need per %q", e.Key, e.Need.Per)
		}
		if e.Pool == "" {
			continue
		}
		// A pool shares one set of players, so its rows must name the same providers' classes.
		if first, seen := poolProviders[e.Pool]; seen && !sameClasses(first, e.Providers) {
			t.Errorf("%s: pool %s mixes providers", e.Key, e.Pool)
		}
		poolProviders[e.Pool] = e.Providers
	}
	if n := len(raidCoverCatalogue); n < 15 || n > 25 {
		t.Errorf("%d rows, want the 15 to 25 a raid leader checks", n)
	}
}

func sameClasses(a, b []coverProvider) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].Class != b[i].Class || a[i].Spec != b[i].Spec || a[i].Talent != b[i].Talent {
			return false
		}
	}
	return true
}

// Every row's spell is in the build, and every provider is a class and spec of specs.json that
// learns its spell at 60 from a trainer (never an AQ book rank) or a talent it names.
func TestRaidCoverProvidersAreTheBuildsOwn(t *testing.T) {
	spells := buildSpellIDs(t)
	specs := curatedSpecs(t)
	books := bookRankIDs(t)
	learn := map[string]classLearns{}
	for _, e := range raidCoverCatalogue {
		if !spells[e.SpellID] {
			t.Errorf("%s: spell %d is not in %s's spells.json", e.Key, e.SpellID, raidCoverBuild)
		}
		for _, p := range e.Providers {
			if _, ok := learn[p.Class]; !ok {
				learn[p.Class] = loadClassLearns(t, p.Class)
			}
			checkProvider(t, e, p, specs, books, learn[p.Class])
		}
	}
}

func checkProvider(t *testing.T, e coverEntry, p coverProvider, specs map[string]string, books map[int]bool, cl classLearns) {
	t.Helper()
	if !classExists(specs, p.Class) {
		t.Errorf("%s: class %q is in no specs.json spec", e.Key, p.Class)
	}
	if p.Spec != "" && specs[p.Spec] != p.Class {
		t.Errorf("%s: spec %q is not a %s spec of specs.json", e.Key, p.Spec, p.Class)
	}
	if p.Talent != "" && !cl.talents[p.Talent] {
		t.Errorf("%s: %s has no talent %q", e.Key, p.Class, p.Talent)
	}
	spell := p.Spell
	if spell == 0 {
		spell = e.SpellID
	}
	if !cl.spells[spell] || books[spell] {
		t.Errorf("%s: a %s does not learn spell %d from a trainer or talent by 60", e.Key, p.Class, spell)
	}
}

// The icon is the client's: the build's spellicons.json entry when it has one, and a file the
// build ships either way.
func TestRaidCoverIconsAreTheBuildsOwn(t *testing.T) {
	var icons map[string]string
	readJSON(t, coverBuildPath("spellicons.json"), &icons)
	for _, e := range raidCoverCatalogue {
		if stem, ok := icons[strconv.Itoa(e.SpellID)]; ok && stem != e.Icon {
			t.Errorf("%s: icon %q, spellicons.json says %q; run go generate ./internal/guilds/", e.Key, e.Icon, stem)
		}
		if _, err := os.Stat(coverBuildPath(fmt.Sprintf("icons/%s.webp", e.Icon))); err != nil {
			t.Errorf("%s: %v", e.Key, err)
		}
	}
}

func buildSpellIDs(t *testing.T) map[int]bool {
	var rows []struct {
		ID int `json:"id"`
	}
	readJSON(t, coverBuildPath("spells.json"), &rows)
	out := make(map[int]bool, len(rows))
	for _, r := range rows {
		out[r.ID] = true
	}
	return out
}

// curatedSpecs is specs.json's spec key -> class slug.
func curatedSpecs(t *testing.T) map[string]string {
	var rows []struct {
		Spec  string `json:"spec"`
		Class string `json:"class_slug"`
	}
	readJSON(t, specsPath, &rows)
	out := make(map[string]string, len(rows))
	for _, r := range rows {
		out[r.Spec] = r.Class
	}
	return out
}

func classExists(specs map[string]string, class string) bool {
	for _, c := range specs {
		if c == class {
			return true
		}
	}
	return false
}

func bookRankIDs(t *testing.T) map[int]bool {
	var file struct {
		Books []struct {
			ID int `json:"id"`
		} `json:"books"`
	}
	readJSON(t, bookRanksPath, &file)
	out := map[int]bool{}
	for _, b := range file.Books {
		out[b.ID] = true
	}
	return out
}

// classLearns is what a class learns by 60: trainer and talent rank ids, and its talent names.
type classLearns struct {
	spells  map[int]bool
	talents map[string]bool
}

func loadClassLearns(t *testing.T, class string) classLearns {
	var trainables struct {
		Trainables []struct {
			Ranks []struct {
				ID    int `json:"id"`
				Level int `json:"level"`
			} `json:"ranks"`
		} `json:"trainables"`
	}
	readJSON(t, coverBuildPath("trainables/"+class+".json"), &trainables)
	var talents struct {
		Trees []struct {
			Talents []struct {
				Name  string `json:"name"`
				Ranks []struct {
					SpellID int `json:"spell_id"`
				} `json:"ranks"`
			} `json:"talents"`
		} `json:"trees"`
	}
	readJSON(t, coverBuildPath("talents/"+class+".json"), &talents)
	cl := classLearns{spells: map[int]bool{}, talents: map[string]bool{}}
	for _, tr := range trainables.Trainables {
		for _, r := range tr.Ranks {
			if r.Level <= raidLevel {
				cl.spells[r.ID] = true
			}
		}
	}
	for _, tree := range talents.Trees {
		for _, ta := range tree.Talents {
			cl.talents[ta.Name] = true
			for _, r := range ta.Ranks {
				cl.spells[r.SpellID] = true
			}
		}
	}
	return cl
}
