package coach

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/sims"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

const (
	repoBuilds    = "../../../data/builds"
	furyEncounter = 1084
)

type fakeSpecs struct {
	worst map[string][]string
	err   error
}

func (f fakeSpecs) Specs(context.Context) ([]sims.SpecFidelity, error) {
	var out []sims.SpecFidelity
	for spec, names := range f.worst {
		s := sims.SpecFidelity{Spec: spec}
		for _, n := range names {
			s.WorstActions = append(s.WorstActions, sims.WorstAction{Name: n})
		}
		out = append(out, s)
	}
	return out, f.err
}

var furyRow = summary.RosterRow{Class: "Warrior", Spec: "Fury", Role: "dps", Level: 60}

// coachDirWith writes the Fury rule file into a temp dir, as Task 15's data will ship it.
func coachDirWith(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "warrior-fury.json"), []byte(validRules), 0o600); err != nil {
		t.Fatal(err)
	}
	return dir
}

func TestTablesFor_FuryAtTheActiveBuild(t *testing.T) {
	ft := &FileTables{BuildsDir: repoBuilds, CoachDir: coachDirWith(t), Specs: fakeSpecs{worst: map[string][]string{"warrior-fury": {"Hamstring"}}}}
	tb, err := ft.TablesFor(context.Background(), furyRow, furyEncounter, time.Now())
	if err != nil {
		t.Fatal(err)
	}
	if tb.Rules.Spec != "warrior-fury" || tb.Spells[23881].CooldownMS != 6000 || !tb.WorstActions["Hamstring"] {
		t.Fatalf("rules %q bt cd %d worst %v", tb.Rules.Spec, tb.Spells[23881].CooldownMS, tb.WorstActions)
	}
	if _, ok := tb.Trainable["Bloodthirst"]; !ok || tb.Mechanics == nil {
		t.Fatalf("trainables %v mechanics %v", tb.Trainable["Bloodthirst"], tb.Mechanics)
	}
	if len(tb.Consumables.Weights) == 0 || len(tb.RaidBuffs) == 0 || tb.Assignments != nil {
		t.Fatalf("catalogue %+v buffs %d", tb.Consumables, len(tb.RaidBuffs))
	}
}

func writeManifest(t *testing.T, root, build, fetched string) {
	t.Helper()
	dir := filepath.Join(root, build)
	if err := os.MkdirAll(dir, 0o750); err != nil {
		t.Fatal(err)
	}
	body := `{"build":"` + build + `","fetched_at":"` + fetched + `"}`
	if err := os.WriteFile(filepath.Join(dir, "manifest.json"), []byte(body), 0o600); err != nil {
		t.Fatal(err)
	}
}

func buildsFixture(t *testing.T) *FileTables {
	root := t.TempDir()
	writeManifest(t, root, "1.60.1.2", "2026-10-01T00:00:00Z")
	writeManifest(t, root, "1.60.1.10", "2026-10-20T00:00:00Z")
	writeManifest(t, root, "forever-prebeta", "2026-10-05T00:00:00Z")
	return &FileTables{BuildsDir: root}
}

func TestBuildAt_ManifestDatesNotDirectoryOrder(t *testing.T) {
	ft := buildsFixture(t)
	for at, want := range map[string]string{"2026-10-10T00:00:00Z": "1.60.1.2", "2026-10-25T00:00:00Z": "1.60.1.10"} {
		ts, _ := time.Parse(time.RFC3339, at)
		got, err := ft.BuildAt(ts)
		if err != nil || got != want {
			t.Fatalf("BuildAt(%s) = %q, %v; want %q", at, got, err, want)
		}
	}
}

func TestBuildAt_BeforeEveryBuildIsTheOldest(t *testing.T) {
	got, err := buildsFixture(t).BuildAt(time.Date(2026, 9, 1, 0, 0, 0, 0, time.UTC))
	if err != nil || got != "1.60.1.2" {
		t.Fatalf("got %q, %v", got, err)
	}
}

func TestTablesFor_NoRuleFileIsTheZeroValue(t *testing.T) {
	ft := &FileTables{BuildsDir: repoBuilds, CoachDir: t.TempDir()}
	tb, err := ft.TablesFor(context.Background(), furyRow, furyEncounter, time.Now())
	if err != nil || tb.Rules.Spec != "" || len(tb.Spells) == 0 {
		t.Fatalf("rules %+v spells %d err %v", tb.Rules, len(tb.Spells), err)
	}
}

func TestTablesFor_HealerCatalogueAndNoEncounterTable(t *testing.T) {
	ft := &FileTables{BuildsDir: repoBuilds, CoachDir: t.TempDir()}
	row := summary.RosterRow{Class: "Priest", Spec: "Holy", Role: "healer", Level: 60}
	tb, err := ft.TablesFor(context.Background(), row, 0, time.Now())
	if err != nil || tb.Mechanics != nil || len(tb.Consumables.Weights) == 0 {
		t.Fatalf("mechanics %v catalogue %+v err %v", tb.Mechanics, tb.Consumables, err)
	}
}

func TestTablesFor_SpecsErrorIsReturned(t *testing.T) {
	boom := errors.New("boom")
	ft := &FileTables{BuildsDir: repoBuilds, CoachDir: t.TempDir(), Specs: fakeSpecs{err: boom}}
	if _, err := ft.TablesFor(context.Background(), furyRow, furyEncounter, time.Now()); !errors.Is(err, boom) {
		t.Fatalf("err %v", err)
	}
}

func TestTablesFor_NoSpecIsNoTables(t *testing.T) {
	ft := &FileTables{BuildsDir: repoBuilds, CoachDir: t.TempDir()}
	_, err := ft.TablesFor(context.Background(), summary.RosterRow{Class: "Warrior"}, 0, time.Now())
	if !errors.Is(err, ErrNoTables) {
		t.Fatalf("err %v", err)
	}
}

func TestTablesFor_UnreadableSpellconstIsARealError(t *testing.T) {
	root := t.TempDir()
	writeManifest(t, root, "1.60.1.2", "2026-10-01T00:00:00Z")
	dir := filepath.Join(root, "1.60.1.2", "spellconst")
	if err := os.MkdirAll(dir, 0o750); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "warrior.json"), []byte("{nope"), 0o600); err != nil {
		t.Fatal(err)
	}
	ft := &FileTables{BuildsDir: root, CoachDir: t.TempDir()}
	_, err := ft.TablesFor(context.Background(), furyRow, 0, time.Now())
	if err == nil || errors.Is(err, ErrNoTables) {
		t.Fatalf("want a real error, got %v", err)
	}
	mage := summary.RosterRow{Class: "Mage", Spec: "Fire"}
	if _, err := ft.TablesFor(context.Background(), mage, 0, time.Now()); !errors.Is(err, ErrNoTables) {
		t.Fatalf("absent class file: %v", err)
	}
}

func TestTablesFor_ConcurrentCallsAreSafe(t *testing.T) {
	ft := &FileTables{BuildsDir: repoBuilds, CoachDir: coachDirWith(t)}
	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if _, err := ft.TablesFor(context.Background(), furyRow, furyEncounter, time.Now()); err != nil {
				t.Error(err)
			}
		}()
	}
	wg.Wait()
}
