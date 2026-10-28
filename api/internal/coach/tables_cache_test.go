package coach

import (
	"context"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/sims"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

type countingSpecs struct{ calls atomic.Int32 }

func (c *countingSpecs) Specs(context.Context) ([]sims.SpecFidelity, error) {
	c.calls.Add(1)
	return []sims.SpecFidelity{{Spec: "warrior-fury", WorstActions: []sims.WorstAction{{Name: "Hamstring"}}}}, nil
}

// writeClassBuild writes a one-spell warrior spellconst with the given cooldown.
func writeClassBuild(t *testing.T, root, build, fetched string, cooldownMS int) {
	t.Helper()
	writeManifest(t, root, build, fetched)
	for dir, body := range map[string]string{
		"spellconst": `{"spells":{"1":{"name":"X","cooldown_ms":` + strconv.Itoa(cooldownMS) + `}}}`,
		"trainables": `{"trainables":[]}`,
	} {
		d := filepath.Join(root, build, dir)
		if err := os.MkdirAll(d, 0o750); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(d, "warrior.json"), []byte(body), 0o600); err != nil {
			t.Fatal(err)
		}
	}
}

func TestTablesFor_NoIOAfterWarmUp(t *testing.T) {
	root, coach := t.TempDir(), coachDirWith(t)
	writeClassBuild(t, root, "1.60.1.2", "2026-10-01T00:00:00Z", 1000)
	specs := &countingSpecs{}
	now := time.Date(2026, 10, 2, 0, 0, 0, 0, time.UTC)
	ft := &FileTables{BuildsDir: root, CoachDir: coach, Specs: specs, Now: func() time.Time { return now }}
	if _, err := ft.TablesFor(context.Background(), furyRow, 0, now); err != nil {
		t.Fatal(err)
	}
	if err := os.RemoveAll(root); err != nil {
		t.Fatal(err)
	}
	if err := os.RemoveAll(coach); err != nil {
		t.Fatal(err)
	}
	tb, err := ft.TablesFor(context.Background(), furyRow, 0, now)
	if err != nil || tb.Rules.Spec != "warrior-fury" || tb.Spells[1].CooldownMS != 1000 || !tb.WorstActions["Hamstring"] {
		t.Fatalf("warm call: %q %v", tb.Rules.Spec, err)
	}
	if n := specs.calls.Load(); n != 1 {
		t.Fatalf("Specs called %d times, want 1", n)
	}
	now = now.Add(specsTTL)
	if _, err := ft.TablesFor(context.Background(), furyRow, 0, now); err != nil || specs.calls.Load() != 2 {
		t.Fatalf("after the ttl: %v, %d calls", err, specs.calls.Load())
	}
}

func TestTablesFor_BuildBoundaryIsInclusive(t *testing.T) {
	root := t.TempDir()
	writeClassBuild(t, root, "1.60.1.2", "2026-10-01T00:00:00Z", 1000)
	writeClassBuild(t, root, "1.60.1.10", "2026-10-20T00:00:00Z", 2000)
	ft := &FileTables{BuildsDir: root, CoachDir: t.TempDir()}
	boundary := time.Date(2026, 10, 20, 0, 0, 0, 0, time.UTC)
	for at, want := range map[time.Time]int64{boundary.Add(-time.Nanosecond): 1000, boundary: 2000} {
		tb, err := ft.TablesFor(context.Background(), furyRow, 0, at)
		if err != nil || tb.Spells[1].CooldownMS != want {
			t.Fatalf("at %s: cooldown %d, want %d (%v)", at, tb.Spells[1].CooldownMS, want, err)
		}
	}
}

func TestTablesFor_GuardianReadsTheFeralBearRules(t *testing.T) {
	dir := t.TempDir()
	rules := strings.ReplaceAll(validRules, "warrior-fury", "druid-feral-bear")
	if err := os.WriteFile(filepath.Join(dir, "druid-feral-bear.json"), []byte(rules), 0o600); err != nil {
		t.Fatal(err)
	}
	ft := &FileTables{BuildsDir: repoBuilds, CoachDir: dir}
	row := summary.RosterRow{Class: "Druid", Spec: "Guardian", Role: "tank", Level: 60}
	tb, err := ft.TablesFor(context.Background(), row, 0, time.Now())
	if err != nil || tb.Rules.Spec != "druid-feral-bear" {
		t.Fatalf("rules %q err %v", tb.Rules.Spec, err)
	}
}
