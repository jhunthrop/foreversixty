package sims

import (
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

func TestRecordedBuffsBecomeEngineIDs(t *testing.T) {
	got := BuffIDs([]summary.AuraRef{
		{SpellID: 20217, Name: "Blessing of Kings"},
		{SpellID: 25289, Name: "Battle Shout"},
		{SpellID: 999999, Name: "Some Forever Aura Nobody Mapped"},
	}, slog.New(slog.NewTextHandler(io.Discard, nil)))
	want := []string{"battle_shout", "blessing_of_kings"}
	if len(got) != len(want) {
		t.Fatalf("ids %v, want %v (the unmapped one is dropped)", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("ids %v, want %v sorted", got, want)
		}
	}
}

func TestTheSameBuffTwiceIsOneID(t *testing.T) {
	quiet := slog.New(slog.NewTextHandler(io.Discard, nil))
	got := BuffIDs([]summary.AuraRef{
		{SpellID: 20217, Name: "Blessing of Kings"},
		{SpellID: 20217, Name: "Blessing of Kings"},
	}, quiet)
	if len(got) != 1 {
		t.Fatalf("ids %v, want one", got)
	}
	if got := BuffIDs(nil, quiet); got == nil || len(got) != 0 {
		t.Fatalf("no auras must give an empty list, never null: %v", got)
	}
}

// TestEveryMappedIDIsInTheEnginesVocabulary is the tripwire on the one
// thing this package cannot check at compile time: sim/request owns the
// id vocabulary and imports the engine, so the api module reads the
// generated document instead.
func TestEveryMappedIDIsInTheEnginesVocabulary(t *testing.T) {
	b, err := os.ReadFile(filepath.Join("..", "..", "..", "sim", "request", "IDS.md"))
	if err != nil {
		t.Fatalf("sim/request/IDS.md is not readable: %v", err)
	}
	doc := string(b)
	for spellID, id := range buffVocabulary {
		if !strings.Contains(doc, "`"+id+"`") {
			t.Errorf("spell %d maps to %q, which is not an id in IDS.md", spellID, id)
		}
	}
	if len(buffVocabulary) == 0 {
		t.Fatal("the buff table is empty; the sim page would run every character unbuffed")
	}
}
