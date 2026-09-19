package phase

import (
	"encoding/json"
	"os"
	"testing"
	"time"
)

// TestBoundariesMatchTheCuratedFile is the reason the table in phase.go
// may be a compiled constant. The dates live in data/curated/phases.json,
// which the pipeline also emits to the site; this table is a copy, and a
// copy nothing compares is a copy that drifts. A phase that moved in one
// place and not the other re-buckets stored rankings silently, which is
// exactly what this catches.
func TestBoundariesMatchTheCuratedFile(t *testing.T) {
	b, err := os.ReadFile("../../../data/curated/phases.json")
	if err != nil {
		t.Fatal(err)
	}
	var file []struct {
		Name  string `json:"name"`
		Start string `json:"start"`
	}
	if err := json.Unmarshal(b, &file); err != nil {
		t.Fatalf("phases.json is not [{name,start}]: %v", err)
	}
	if len(file) != len(Boundaries) {
		t.Fatalf("%d phases in the file, %d in the table", len(file), len(Boundaries))
	}
	for i, want := range file {
		got := Boundaries[i]
		if got.Name != want.Name {
			t.Errorf("phase %d: name %q, want %q", i, got.Name, want.Name)
		}
		// An empty or absent start is the zero time: pre-beta has no
		// opening instant, it is simply everything before beta.
		var start time.Time
		if want.Start != "" {
			start, err = time.Parse(time.RFC3339, want.Start)
			if err != nil {
				t.Fatalf("phase %s: start %q is not RFC3339: %v", want.Name, want.Start, err)
			}
		}
		if !got.Start.Equal(start) {
			t.Errorf("phase %s: start %s, want %s", want.Name, got.Start, start)
		}
	}
}
