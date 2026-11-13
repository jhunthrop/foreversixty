package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func writeFile(t *testing.T, path, body string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
}

func TestLoadSpecFindsBySlugAndReportsFailures(t *testing.T) {
	root := t.TempDir()
	if _, err := loadSpec(root, "a"); err == nil || !strings.Contains(err.Error(), "reading") {
		t.Fatalf("missing file: %v", err)
	}
	specs := filepath.Join(root, "data", "curated", "specs.json")
	writeFile(t, specs, `[{"spec":"a","class_slug":"mage","role":"dps","tree_index":2},{"spec":"b","role":"healer"}]`)
	got, err := loadSpec(root, "b")
	if err != nil || got.Role != "healer" {
		t.Fatalf("got %+v %v", got, err)
	}
	if got, _ := loadSpec(root, "a"); got.TreeIndex != 2 || got.ClassSlug != "mage" {
		t.Fatalf("got %+v", got)
	}
	if _, err := loadSpec(root, "zzz"); err == nil || !strings.Contains(err.Error(), `no spec "zzz"`) {
		t.Fatalf("unknown: %v", err)
	}
	writeFile(t, specs, `{`)
	if _, err := loadSpec(root, "a"); err == nil || !strings.Contains(err.Error(), "decoding") {
		t.Fatalf("bad json: %v", err)
	}
}

func TestAPLStateReadsTheStateField(t *testing.T) {
	root := t.TempDir()
	if _, err := aplState(root, "a"); err == nil {
		t.Fatal("missing file accepted")
	}
	path := filepath.Join(root, "data", "curated", "apl", "a.json")
	writeFile(t, path, `{"state":"written"}`)
	if s, err := aplState(root, "a"); err != nil || s != writtenAPL {
		t.Fatalf("got %q %v", s, err)
	}
	writeFile(t, path, `[`)
	if _, err := aplState(root, "a"); err == nil {
		t.Fatal("bad json accepted")
	}
}

func TestLoadBISBandMatchesLevelFactionAndPreset(t *testing.T) {
	dir := t.TempDir()
	if _, err := loadBISBand(dir, "a", 60, "alliance", "bare"); err == nil {
		t.Fatal("missing file accepted")
	}
	path := filepath.Join(dir, "bis", "a.json")
	writeFile(t, path, `{"bands":[
	  {"band":60,"preset":"bare","faction":"horde","race":"Orc"},
	  {"band":60,"preset":"raid","faction":"alliance","race":"Gnome"},
	  {"band":60,"preset":"bare","faction":"alliance","race":"Human","slots":[{"slot":"head","item_id":5},{"slot":"neck","item_id":0}]}]}`)
	band, err := loadBISBand(dir, "a", 60, "alliance", "bare")
	if err != nil || band.Race != "Human" {
		t.Fatalf("got %+v %v", band, err)
	}
	gear := band.gear()
	if len(gear) != 1 || gear[0].Slot != "head" || gear[0].ItemID != 5 {
		t.Fatalf("empty slots must be dropped, got %+v", gear)
	}
	if _, err := loadBISBand(dir, "a", 40, "alliance", "bare"); err == nil || !strings.Contains(err.Error(), "no level-40") {
		t.Fatalf("absent band: %v", err)
	}
	writeFile(t, path, `x`)
	if _, err := loadBISBand(dir, "a", 60, "alliance", "bare"); err == nil {
		t.Fatal("bad json accepted")
	}
}

// A raid band's slots carry the permanent enchant the ranker measured the
// set wearing (enchant_id). A search that read the items alone simmed the
// rogue raid builds 5 to 6% under the published set_dps (2026-11-13,
// spec-breakdown's 2026-11-11 fix), so the search wears what the band wears.
func TestBISBandGearWearsItsPublishedEnchants(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "bis", "a.json"), `{"bands":[
	  {"band":60,"preset":"raid","faction":"alliance","race":"Human","slots":[
	    {"slot":"main_hand","item_id":5,"enchant_id":1900},{"slot":"head","item_id":6}]}]}`)
	band, err := loadBISBand(dir, "a", 60, "alliance", "raid")
	if err != nil {
		t.Fatal(err)
	}
	got := map[string]int{}
	for _, g := range band.gear() {
		got[g.Slot] = g.Enchant
	}
	if got["main_hand"] != 1900 || got["head"] != 0 || len(got) != 2 {
		t.Fatalf("gear enchants %v, want main_hand 1900 and a bare head", got)
	}
}
