package request

import (
	"bytes"
	"os"
	"slices"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
)

// The contract's section 1.6 table, transcribed. It is written out
// here rather than derived from the map it checks, because a table
// that checked itself would agree with any typo.
func TestEveryStyleExpandsToTheContractsFields(t *testing.T) {
	cases := []struct {
		id      string
		targets int
		execute float64
		move    *api.Movement
		overT   []api.TargetCount
		dummy   bool
	}{
		{id: "patchwerk", targets: 1, execute: 0.25},
		{id: "execute", targets: 1, execute: 0.35},
		{id: "light-movement", targets: 1, execute: 0.25,
			move: &api.Movement{IntervalSec: 45, DurationSec: 5, Kind: api.MovementAway}},
		{id: "heavy-movement", targets: 1, execute: 0.25,
			move: &api.Movement{IntervalSec: 20, DurationSec: 5, Kind: api.MovementAway}},
		{id: "cleave-2", targets: 2, execute: 0.25},
		{id: "cleave-3", targets: 3, execute: 0.25},
		{id: "cleave-5", targets: 5, execute: 0.25},
		{id: "dungeon", targets: 1, execute: 0, overT: []api.TargetCount{
			{AtSec: 0, Count: 1}, {AtSec: 40, Count: 3}, {AtSec: 80, Count: 5},
			{AtSec: 130, Count: 3}, {AtSec: 160, Count: 1}}},
		{id: "dummy", targets: 1, execute: 0, dummy: true},
	}
	if len(cases) != len(StyleIDs) {
		t.Fatalf("the contract lists %d styles and StyleIDs has %d: %v", len(cases), len(StyleIDs), StyleIDs)
	}
	for _, c := range cases {
		t.Run(c.id, func(t *testing.T) {
			if !slices.Contains(StyleIDs, c.id) {
				t.Fatalf("StyleIDs does not carry %q", c.id)
			}
			got, ok := ExpandStyle(c.id, api.DefaultEncounter())
			if !ok {
				t.Fatalf("ExpandStyle(%q) is not a style", c.id)
			}
			if got.Style != c.id {
				t.Errorf("style label is %q", got.Style)
			}
			if got.Targets != c.targets {
				t.Errorf("targets = %d, want %d", got.Targets, c.targets)
			}
			if got.ExecuteRatio != c.execute {
				t.Errorf("execute_ratio = %v, want %v", got.ExecuteRatio, c.execute)
			}
			if got.Dummy != c.dummy {
				t.Errorf("dummy = %v, want %v", got.Dummy, c.dummy)
			}
			switch {
			case c.move == nil && got.Movement != nil:
				t.Errorf("movement = %+v, want none", got.Movement)
			case c.move != nil && got.Movement == nil:
				t.Errorf("movement is absent, want %+v", c.move)
			case c.move != nil && *got.Movement != *c.move:
				t.Errorf("movement = %+v, want %+v", *got.Movement, *c.move)
			}
			if len(got.TargetsOverTime) != len(c.overT) {
				t.Fatalf("targets_over_time = %v, want %v", got.TargetsOverTime, c.overT)
			}
			for i := range c.overT {
				if got.TargetsOverTime[i] != c.overT[i] {
					t.Errorf("targets_over_time[%d] = %+v, want %+v", i, got.TargetsOverTime[i], c.overT[i])
				}
			}
		})
	}
	if _, ok := ExpandStyle("naxx", api.DefaultEncounter()); ok {
		t.Error("ExpandStyle accepted a style that is not in the vocabulary")
	}
}

// Every expansion has to survive the envelope's own validation, or the
// page would offer a preset the run then refuses.
func TestEveryStyleExpandsToALegalEncounter(t *testing.T) {
	for _, id := range StyleIDs {
		t.Run(id, func(t *testing.T) {
			req := fury()
			got, ok := ExpandStyle(id, req.Encounter)
			if !ok {
				t.Fatalf("no such style")
			}
			req.Encounter = got
			if err := req.Validate(); err != nil {
				t.Errorf("the %q preset does not validate: %v", id, err)
			}
		})
	}
}

// The fields a style does NOT set are the player's, and an expansion
// that reset them would silently throw away a chosen fight length.
func TestExpandStyleKeepsWhatTheStyleDoesNotSet(t *testing.T) {
	base := api.DefaultEncounter()
	base.DurationSec = 300
	base.Variation = 0.1
	base.TargetLevel = 61
	base.TargetArmor = 2500
	base.TargetType = "undead"
	got, _ := ExpandStyle("cleave-3", base)
	if got.DurationSec != 300 || got.Variation != 0.1 || got.TargetLevel != 61 || got.TargetArmor != 2500 || got.TargetType != "undead" {
		t.Errorf("the style overwrote what it does not own: %+v", got)
	}
}

// styles.json is the web lane's fixture: its own style table is tested
// against this file, so the two cannot drift. A stale copy is a failing
// test here rather than a disagreement nobody notices.
func TestStylesJSONIsCommitted(t *testing.T) {
	want, err := StylesJSON()
	if err != nil {
		t.Fatal(err)
	}
	got, err := os.ReadFile("styles.json")
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(bytes.TrimSpace(got), bytes.TrimSpace(want)) {
		t.Error("sim/request/styles.json is stale; run `go run ./internal/genstyles` from sim/")
	}
}
