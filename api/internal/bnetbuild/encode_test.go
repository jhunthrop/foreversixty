// api/internal/bnetbuild/encode_test.go
package bnetbuild

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/bnetapi"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

func loadEraKilozInputs(t *testing.T) Inputs {
	t.Helper()
	data, err := trees.Load(realDataDir)
	if err != nil {
		t.Fatal(err)
	}
	tables, ok, err := LoadTables(realDataDir, data)
	if err != nil || !ok {
		t.Fatalf("LoadTables: ok=%v err=%v", ok, err)
	}
	return Inputs{
		Build:     tables.Build,
		Profile:   bnetapi.CharacterProfile{Name: "Kiloz", ClassSlug: "warrior", RaceName: "Orc"},
		Equipment: readEraKiloz(t, "equipment.json"),
		Talents:   readEraKiloz(t, "specializations.json"),
		Talent:    TalentTable{Build: tables.Trees, ClassID: classIDFor(t, tables.Trees, "warrior")},
		Enchants:  tables.Enchants,
		Suffixes:  tables.Suffixes,
		Races:     tables.Races,
	}
}

func TestEncodeEraKilozHasTheRightHeadAndGear(t *testing.T) {
	code, report, err := Encode(loadEraKilozInputs(t))
	if err != nil {
		t.Fatal(err)
	}
	parts := strings.Split(code, ":")
	if len(parts) < 6 || parts[0] != "FS1" || parts[2] != "warrior" || parts[3] != "orc" {
		t.Fatalf("code head = %v, want FS1:<build>:warrior:orc:...", parts)
	}
	if !strings.Contains(code, "head=21329:2583") {
		t.Fatalf("code = %q, missing the fixture's head slot", code)
	}
	if len(report.SkippedSlots) != 2 {
		t.Fatalf("SkippedSlots = %v, want exactly SHIRT and TABARD", report.SkippedSlots)
	}
}

func TestEncodeRefusesAnUnknownRace(t *testing.T) {
	in := loadEraKilozInputs(t)
	in.Profile.RaceName = "Ogre"
	if _, _, err := Encode(in); err == nil {
		t.Fatal("an unmapped race must return an error, not a guessed slug")
	}
}

func TestEraKilozFixtureMatchesTheCheckedInFile(t *testing.T) {
	code, _, err := Encode(loadEraKilozInputs(t))
	if err != nil {
		t.Fatal(err)
	}
	want, err := os.ReadFile(filepath.Join("testdata", "era-kiloz.fs1"))
	if err != nil {
		t.Fatal(err)
	}
	if code != string(want) {
		t.Fatalf("Encode(era-kiloz) = %q, want the checked-in fixture %q", code, want)
	}
}

// encodeEraKilozFrom is the era-kiloz code against the builds under dir, as cmd/api wires it.
func encodeEraKilozFrom(t *testing.T, dir string) (Tables, string, error) {
	t.Helper()
	data, err := trees.Load(dir)
	if err != nil {
		t.Fatal(err)
	}
	tables, ok, err := LoadTables(dir, data)
	if err != nil || !ok {
		t.Fatalf("LoadTables: ok=%v err=%v", ok, err)
	}
	in := loadEraKilozInputs(t)
	in.Build, in.Enchants, in.Suffixes, in.Races = tables.Build, tables.Enchants, tables.Suffixes, tables.Races
	in.Talent = TalentTable{Build: tables.Trees, ClassID: classIDFor(t, tables.Trees, "warrior")}
	code, _, err := Encode(in)
	return tables, code, err
}

// A build data.yml fetched but no activation moved to (70685 on disk, active-build.json still
// naming 70658) must not change the stamp; activating it does.
func TestEncodeStampsTheActiveBuildNotTheNewestFolder(t *testing.T) {
	const active, fetched = "1.60.1.70658", "1.60.1.70685"
	root := t.TempDir()
	if err := trees.LinkBuilds(root, realDataDir, active, fetched); err != nil {
		t.Fatal(err)
	}
	if err := trees.WriteActiveBuildFile(root, active); err != nil {
		t.Fatal(err)
	}
	tables, code, err := encodeEraKilozFrom(t, root)
	if err != nil || tables.Build != active || !strings.HasPrefix(code, "FS1:"+active+":") {
		t.Fatalf("build %q code %q err %v, want the active %s", tables.Build, code, err, active)
	}

	if err := trees.WriteActiveBuildFile(root, fetched); err != nil {
		t.Fatal(err)
	}
	tables, code, err = encodeEraKilozFrom(t, root)
	if tables.Build != fetched {
		t.Fatalf("tables after activation = %q, want %s", tables.Build, fetched)
	}
	if err == nil && !strings.HasPrefix(code, "FS1:"+fetched+":") {
		t.Fatalf("code after activation = %q, want the %s stamp", code, fetched)
	}
}
