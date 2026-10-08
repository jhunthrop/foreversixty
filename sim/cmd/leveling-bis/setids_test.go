package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/internal/inproc"
	"github.com/jhunthrop/foreversixty/sim/leveling"
	"github.com/wowsims/classic/sim/core"
)

// setidsFile is the generated file this test rewrites.
const setidsFile = "setids_generated.go"

// setidsPerLine is how many entries the generated map prints per line.
const setidsPerLine = 10

// setsFileRow is the part of data/builds/<build>/sets.json this test reads.
type setsFileRow struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// engineSetIDs is every sets.json id the engine has a registered set for.
func engineSetIDs(t *testing.T) []int {
	t.Helper()
	inproc.Register()
	build, err := leveling.ReadActiveBuild(publishedRepoRoot)
	if err != nil {
		t.Fatal(err)
	}
	raw, err := os.ReadFile(filepath.Join(publishedRepoRoot, "data", "builds", build, "sets.json"))
	if err != nil {
		t.Fatal(err)
	}
	var sets []setsFileRow
	if err := json.Unmarshal(raw, &sets); err != nil {
		t.Fatal(err)
	}

	registered := core.RegisteredItemSets()
	var ids []int
	for _, set := range sets {
		for _, engineSet := range registered {
			if int(engineSet.ID) == set.ID || engineSet.Name == set.Name || engineSet.AlternativeName == set.Name {
				ids = append(ids, set.ID)
				break
			}
		}
	}
	sort.Ints(ids)
	return ids
}

func renderSetIDs(ids []int) string {
	var b strings.Builder
	b.WriteString("var engineImplementedSetIDs = map[int]bool{\n")
	for i := 0; i < len(ids); i += setidsPerLine {
		end := min(i+setidsPerLine, len(ids))
		entries := make([]string, 0, setidsPerLine)
		for _, id := range ids[i:end] {
			entries = append(entries, fmt.Sprintf("%d: true,", id))
		}
		b.WriteString("\t" + strings.Join(entries, " ") + "\n")
	}
	b.WriteString("}\n")
	return b.String()
}

// TestEngineImplementedSetIDs holds the generated set list to the engine.
func TestEngineImplementedSetIDs(t *testing.T) {
	ids := engineSetIDs(t)
	want := renderSetIDs(ids)

	source, err := os.ReadFile(setidsFile)
	if err != nil {
		t.Fatal(err)
	}
	text := string(source)
	start := strings.Index(text, "var engineImplementedSetIDs")
	end := strings.Index(text, "func setEffectImplemented")
	if start < 0 || end < 0 {
		t.Fatalf("%s lost its map or its accessor", setidsFile)
	}
	have := strings.TrimRight(text[start:end], "\n") + "\n"

	if have == want {
		return
	}
	if os.Getenv("FOREVER_UPDATE_SETIDS") == "1" {
		rewritten := text[:start] + want + "\n" + text[end:]
		if err := os.WriteFile(setidsFile, []byte(rewritten), 0o644); err != nil {
			t.Fatal(err)
		}
		return
	}
	t.Errorf("%s is stale against the engine's item set registrations; run FOREVER_UPDATE_SETIDS=1 go test ./cmd/leveling-bis -run TestEngineImplementedSetIDs and review the diff", setidsFile)
}
