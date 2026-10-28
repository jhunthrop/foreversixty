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
	if n := len(d.Find("the whirlwind is lowercase prose and flask of the titans too")); n != 0 {
		t.Errorf("lowercase prose found %d names", n)
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
