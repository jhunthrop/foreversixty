package app

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/rankings"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

// The home answer's reference world: the upgrades package's parity character (a real band and
// item rows of build 1.60.1.70658, the same input the site's upgrades.ts is pinned against),
// worn by Velaren, a level 60 Dwarf Beast Mastery Hunter.
const (
	homeBuild     = "1.60.1.70658"
	homeBisBuild  = "bisbuild"
	homeParity    = "../upgrades/testdata/parity/input.json"
	homeKey       = "us/normal/velaren"
	homeTalents   = "5420001505001250/0053502001000000/5"
	homeContracts = "../../../docs/contracts/fixtures/app"
)

var homeSynced = time.Date(2026, 10, 29, 21, 15, 0, 0, time.UTC)

type parityInput struct {
	Band  json.RawMessage `json:"band"`
	Gear  map[string]int  `json:"gear"`
	Items []trees.Item    `json:"items"`
}

func readParityInput(t *testing.T) parityInput {
	t.Helper()
	raw, err := os.ReadFile(homeParity)
	if err != nil {
		t.Fatal(err)
	}
	var in parityInput
	if err := json.Unmarshal(raw, &in); err != nil {
		t.Fatal(err)
	}
	return in
}

// fakeCharacters serves the character cards by lowercase name.
type fakeCharacters struct {
	cards map[string]rankings.Character
	err   error
}

func (f *fakeCharacters) CharacterCard(_ context.Context, region, ruleset, name string) (rankings.Character, bool, error) {
	if f.err != nil {
		return rankings.Character{}, false, f.err
	}
	if strings.ToLower(region) != "us" || strings.ToLower(ruleset) != "normal" {
		return rankings.Character{}, false, nil
	}
	c, ok := f.cards[strings.ToLower(name)]
	return c, ok, nil
}

// fakeData is the reference build: its BiS file under dir and its hunter item rows.
type fakeData struct {
	dir    string
	items  map[int]trees.Item
	active string
}

func (f *fakeData) ActiveBuild() (string, bool) { return f.active, f.active != "" }
func (f *fakeData) BisBuild() string            { return homeBisBuild }
func (f *fakeData) BisDataDir() string          { return f.dir }
func (f *fakeData) RaceFaction(race string) string {
	if race == "Dwarf" {
		return "alliance"
	}
	return ""
}
func (f *fakeData) BuildItem(version, class string, id int) (trees.Item, bool) {
	it, ok := f.items[id]
	return it, ok && version == homeBuild && class == "hunter"
}

// fakeExports serves the stored exports by character key.
type fakeExports struct {
	rows map[string]Export
	err  error
}

func (f *fakeExports) Export(_ context.Context, key string) (Export, bool, error) {
	if f.err != nil {
		return Export{}, false, f.err
	}
	e, ok := f.rows[key]
	return e, ok, nil
}

type homeWorld struct {
	svc     *Service
	chars   *fakeCharacters
	exports *fakeExports
	in      parityInput
}

func newHomeWorld(t *testing.T) *homeWorld {
	t.Helper()
	in := readParityInput(t)
	dir := t.TempDir()
	bisDir := filepath.Join(dir, homeBisBuild, "bis")
	if err := os.MkdirAll(bisDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(bisDir, "hunter-beast-mastery.json"), []byte(`{"bands": [`+string(in.Band)+`]}`), 0o600); err != nil {
		t.Fatal(err)
	}
	items := map[int]trees.Item{}
	for _, it := range in.Items {
		items[it.ID] = it
	}
	w := &homeWorld{
		chars:   &fakeCharacters{cards: map[string]rankings.Character{"velaren": velaren(in.Gear)}},
		exports: &fakeExports{rows: map[string]Export{homeKey: {Code: velarenExport(in.Gear), Source: "addon", CapturedAt: homeSynced}}},
		in:      in,
	}
	w.svc = &Service{
		Characters: w.chars, Data: &fakeData{dir: dir, items: items, active: homeBuild}, Exports: w.exports,
		Settings: Settings{WebCheckoutRegions: map[string][]string{"ios": {"US"}, "android": {}}, MinVersionIOS: "1.0.0", MinVersionAndroid: "1.0.0"},
		Log:      slog.New(slog.NewTextHandler(io.Discard, nil)),
	}
	return w
}

// velaren is the fixture's public character card: the gear panel's 17 rows over gear.
func velaren(gear map[string]int) rankings.Character {
	var slots []rankings.GearSlot
	for _, key := range []string{"head", "neck", "shoulder", "back", "chest", "wrist", "hands", "waist", "legs", "feet",
		"finger1", "finger2", "trinket1", "trinket2", "main_hand", "off_hand", "ranged"} {
		row := rankings.GearSlot{Slot: key}
		if id := gear[key]; id > 0 {
			row.Worn = &rankings.GearWorn{ID: id}
		}
		slots = append(slots, row)
	}
	level, race, spec := 60, "Dwarf", "hunter-beast-mastery"
	return rankings.Character{
		Character: rankings.CharacterRef{Key: homeKey, Region: "us", Ruleset: "normal", Name: "Velaren", Class: "Hunter",
			Level: &level, Race: &race},
		Gear: slots, GearSpec: &spec,
	}
}

// velarenExport is the addon export the card's gear came from, with its talents.
func velarenExport(gear map[string]int) string {
	var parts []string
	for slot, id := range gear {
		parts = append(parts, slot+"="+strconv.Itoa(id))
	}
	sort.Strings(parts)
	return "FS1:" + homeBuild + ":hunter:dwarf:" + homeTalents + ":" + strings.Join(parts, ",") + "|level=60"
}

var errBoom = errors.New("boom")
