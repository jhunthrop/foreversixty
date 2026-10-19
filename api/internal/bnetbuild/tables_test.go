// api/internal/bnetbuild/tables_test.go
package bnetbuild

import (
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

// realDataDir is the repo's real client data, used (not a synthetic fixture) so this
// package's tests verify against the same tables bnetimport loads in production — the era-
// kiloz fixture (api/internal/bnetapi/testdata/era-kiloz/) is a real character captured
// against this same active build.
const realDataDir = "../../../data/builds"

func TestLoadEnchantTableFindsTheEraKilozHeadEnchant(t *testing.T) {
	table, err := LoadEnchantTable(filepath.Join(realDataDir, "1.60.1.69893", "enchants.json"))
	if err != nil {
		t.Fatal(err)
	}
	if !table[2583] {
		t.Fatal("enchant 2583 (Presence of Might, on the fixture's head slot) must be known")
	}
	if table[999999999] {
		t.Fatal("an id absent from enchants.json must not read as known")
	}
}

func TestLoadSuffixTableFindsAKnownSuffix(t *testing.T) {
	table, err := LoadSuffixTable(filepath.Join(realDataDir, "1.60.1.69893", "suffixes.json"))
	if err != nil {
		t.Fatal(err)
	}
	if table["of the Falcon"] == 0 {
		t.Fatal(`"of the Falcon" must resolve to a non-zero suffix id`)
	}
}

func TestRaceTableFromMapsBlizzardNamesToSlugs(t *testing.T) {
	data, err := trees.Load(realDataDir)
	if err != nil {
		t.Fatal(err)
	}
	build, ok := data.Build("1.60.1.69893")
	if !ok {
		t.Fatal("1.60.1.69893 must load")
	}
	races := RaceTableFrom(build)
	if races["Orc"] != "orc" {
		t.Fatalf(`races["Orc"] = %q, want "orc"`, races["Orc"])
	}
}

func TestLoadTablesResolvesTheActiveBuild(t *testing.T) {
	data, err := trees.Load(realDataDir)
	if err != nil {
		t.Fatal(err)
	}
	tables, ok, err := LoadTables(realDataDir, data)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("LoadTables must resolve a build from the real data/builds directory")
	}
	if want := newestClientBuildDir(t); tables.Build != want {
		t.Fatalf("Build = %q, want the newest numeric client build, %q", tables.Build, want)
	}
	if tables.Trees == nil || tables.Enchants == nil || tables.Suffixes == nil || tables.Races == nil {
		t.Fatalf("LoadTables left a table nil: %+v", tables)
	}
}

// newestClientBuildDir is the highest-numbered client build directory under realDataDir,
// the build LoadTables must resolve: the data pipeline lands a new client's tables there
// before the site activates them, so the expectation is read from the directory, not pinned.
func newestClientBuildDir(t *testing.T) string {
	t.Helper()
	entries, err := os.ReadDir(realDataDir)
	if err != nil {
		t.Fatal(err)
	}
	clientBuild := regexp.MustCompile(`^\d+(\.\d+)+$`)
	var newest string
	var newestParts []int
	for _, entry := range entries {
		if !entry.IsDir() || !clientBuild.MatchString(entry.Name()) {
			continue
		}
		parts := buildParts(entry.Name())
		if newest == "" || buildLess(newestParts, parts) {
			newest, newestParts = entry.Name(), parts
		}
	}
	if newest == "" {
		t.Fatalf("no client build directory under %s", realDataDir)
	}
	return newest
}

func buildParts(version string) []int {
	fields := strings.Split(version, ".")
	parts := make([]int, len(fields))
	for i, field := range fields {
		parts[i], _ = strconv.Atoi(field)
	}
	return parts
}

func buildLess(a, b []int) bool {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			return a[i] < b[i]
		}
	}
	return len(a) < len(b)
}
