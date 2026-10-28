package coach

import (
	"os"
	"path/filepath"
	"slices"
	"sync"
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/mechanics"
)

const mechanicsTablesDir = "../../../logs/engine/mechanics/tables"

var (
	namesOnce sync.Once
	namesDict *NameDict
	namesErr  error
)

// buildNames is the dictionary of the newest build and every mechanics table.
func buildNames(t *testing.T) *NameDict {
	t.Helper()
	build := activeBuild(t)
	namesOnce.Do(func() {
		var tables []mechanics.Table
		paths, _ := filepath.Glob(filepath.Join(mechanicsTablesDir, "*.json"))
		for _, p := range paths {
			data, err := os.ReadFile(p)
			if err != nil {
				namesErr = err
				return
			}
			table, err := mechanics.Parse(data)
			if err != nil {
				namesErr = err
				return
			}
			tables = append(tables, table)
		}
		namesDict, namesErr = LoadNames(filepath.Join(buildsDir, build), tables)
	})
	if namesErr != nil {
		t.Fatal(namesErr)
	}
	return namesDict
}

func TestNameDict_FindsSpellsItemsBosses(t *testing.T) {
	d := buildNames(t)
	got := d.Find("Open with Whirlwind, bring a Flask of the Titans, and Onyxia waits.")
	for _, want := range []string{"Whirlwind", "Flask of the Titans", "Onyxia"} {
		if !slices.Contains(got, want) {
			t.Errorf("%q not found in %v", want, got)
		}
	}
}

func TestNameDict_LowercasePhrasesAreFoundSingleWordsAreNot(t *testing.T) {
	d := buildNames(t)
	got := d.Find("drink a flask of the titans and keep mortal strike up")
	for _, want := range []string{"Flask of the Titans", "Mortal Strike"} {
		if !slices.Contains(got, want) {
			t.Errorf("%q not found in %v", want, got)
		}
	}
	if got := d.Find("the unused charge and rage, a burn phase, fear and haste"); len(got) != 0 {
		t.Errorf("ordinary lowercase words found: %v", got)
	}
}

// A lowercase single-word class spell is still the spell ("press whirlwind"): the validator
// must catch it as it catches "Whirlwind".
func TestNameDict_LowercaseSingleWordAbilitiesAreNames(t *testing.T) {
	d := buildNames(t)
	for _, probe := range []struct{ prose, want string }{
		{"Press whirlwind on cooldown.", "Whirlwind"},
		{"then bloodthirst", "Bloodthirst"},
		{"use execute below 20", "Execute"},
		{"weave a slam", "Slam"},
		{"keep hamstring up", "Hamstring"},
		{"refresh rend", "Rend"},
		{"two whirlwinds", "Whirlwind"},
	} {
		if got := d.Find(probe.prose); !slices.Contains(got, probe.want) {
			t.Errorf("%q: %v, want %q", probe.prose, got, probe.want)
		}
	}
}

// ordinaryProseWords is a reviewed list: every entry is a single-word class spell or talent
// name of the build (else it guards nothing and has rotted), written in lowercase.
func TestOrdinaryProseWords_AreSingleWordAbilities(t *testing.T) {
	d := buildNames(t)
	for w := range ordinaryProseWords {
		if !d.abilities[w] {
			t.Errorf("%q is not a single-word ability name of the build", w)
		}
		if got := d.Find(w); len(got) != 0 {
			t.Errorf("%q lowercase is prose, found %v", w, got)
		}
	}
}

func TestNameDict_OrdinaryLowercaseWordsAreProse(t *testing.T) {
	d := buildNames(t)
	if got := d.Find("spend rage before it caps, the fire and frost mark the unused rage"); len(got) != 0 {
		t.Errorf("ordinary prose found %v", got)
	}
}

func TestNameDict_LongestMatchWins(t *testing.T) {
	d := &NameDict{names: map[string]string{}}
	d.add("Heroic", "Heroic Strike")
	got := d.Find("Heroic Strike waits")
	if !slices.Equal(got, []string{"Heroic Strike"}) {
		t.Fatalf("%v", got)
	}
}

func TestNameDict_TalentsFoundDeprecatedSkipped(t *testing.T) {
	d := buildNames(t)
	if got := d.Find("DEPRECATED"); len(got) != 0 {
		t.Errorf("deprecated item kept: %v", got)
	}
	if got := d.Find("Improved Heroic Strike"); !slices.Contains(got, "Improved Heroic Strike") {
		t.Errorf("talent not found: %v", got)
	}
}

func TestLoadNames_MissingBuildIsAnError(t *testing.T) {
	if _, err := LoadNames(t.TempDir(), nil); err == nil {
		t.Fatal("empty build dir accepted")
	}
}
