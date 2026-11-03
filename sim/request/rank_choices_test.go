package request

import (
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/enginever"
	"github.com/jhunthrop/foreversixty/sim/internal/spellranks"
	"github.com/jhunthrop/foreversixty/sim/specs"
)

const (
	curatedRankChoices = "../../data/curated/rank-choices.json"
	// activeSpellRanks is the table spellranks embeds (`make simdb` copies it).
	activeSpellRanks = "../internal/spellranks/spellranks.json"
)

func TestEmbeddedRankChoicesMatchCurated(t *testing.T) {
	curated, err := os.ReadFile(filepath.FromSlash(curatedRankChoices))
	if err != nil {
		t.Fatalf("reading the curated rank choices: %v", err)
	}
	if string(curated) != string(rankChoicesJSON) {
		t.Fatal("sim/rankchoices/rank-choices.json differs from data/curated/rank-choices.json; run `make rank-choices-sync`")
	}
	if _, err := embeddedRankChoices(); err != nil {
		t.Fatalf("the embedded rank choices do not load: %v", err)
	}
}

// Every row is a real rank choice the rotation can cast: a spec this build
// sims, an id its rotation names, the spell its name says, levels 1 to 60
// (a row at 60 names its preset; one for both presets is the rotation's own
// rank), and a rank learned by the row's first level that differs from what
// the default rule casts there.
func TestCuratedRankChoicesAreCastableDownranks(t *testing.T) {
	var file rankChoicesFile
	readJSON(t, curatedRankChoices, &file)
	var ranks struct {
		Classes map[string]map[string][]struct{ ID int32 } `json:"classes"`
	}
	readJSON(t, activeSpellRanks, &ranks)
	for spec, spells := range file.Specs {
		s, ok := specs.ByKey[spec]
		if !ok {
			t.Errorf("%s is not a spec", spec)
			continue
		}
		raw, err := aplFS.ReadFile("apl/" + spec + ".apl.json")
		if err != nil {
			t.Fatal(err)
		}
		for _, sp := range spells {
			if !strings.Contains(string(raw), `"spellId": `+strconv.Itoa(int(sp.Authored))) {
				t.Errorf("%s: its rotation does not name %d", spec, sp.Authored)
			}
			named := false
			for _, row := range ranks.Classes[s.ClassSlug][sp.Spell] {
				named = named || row.ID == sp.Authored
			}
			if !named {
				t.Errorf("%s: %d is not a rank of %q", spec, sp.Authored, sp.Spell)
			}
			for _, row := range sp.Bands {
				if row.FromLevel < 1 || row.ToLevel > api.MaxLevel {
					t.Errorf("%s %s: a row over levels %d to %d", spec, sp.Spell, row.FromLevel, row.ToLevel)
					continue
				}
				if row.ToLevel >= api.MaxLevel && row.Preset == "" {
					t.Errorf("%s %s: a row at 60 for both presets belongs in the curated rotation itself", spec, sp.Spell)
				}
				chosen, learned := spellranks.SpellIDOfRank(s.ClassSlug, sp.Authored, row.Rank, row.FromLevel)
				if !learned {
					t.Errorf("%s %s rank %d is not learned at %d", spec, sp.Spell, row.Rank, row.FromLevel)
					continue
				}
				if def, _ := rankResolver(s.ClassSlug, row.FromLevel, nil)(sp.Authored); def == chosen {
					t.Errorf("%s %s rank %d from %d is what the rotation casts anyway", spec, sp.Spell, row.Rank, row.FromLevel)
				}
			}
		}
	}
}

func TestParseRankChoicesRefusesDuplicates(t *testing.T) {
	for _, raw := range []string{
		`{"specs":{"x":[{"authored":1,"bands":[]},{"authored":1,"bands":[]}]}}`,
		`{"specs":{"x":[{"authored":1,"bands":[{"from_level":40,"to_level":49,"rank":2},{"from_level":45,"to_level":49,"rank":3}]}]}}`,
		`{"specs":{"x":[{"authored":1,"bands":[{"from_level":40,"to_level":39,"rank":2}]}]}}`,
		`not json`,
	} {
		if _, err := parseRankChoices([]byte(raw)); err == nil {
			t.Errorf("accepted %s", raw)
		}
	}
	got, err := parseRankChoices([]byte(`{"specs":{"x":[{"authored":7,"bands":[{"from_level":50,"to_level":59,"rank":3},{"from_level":30,"to_level":39,"rank":2}]}]}}`))
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct{ level, rank int }{{29, 0}, {30, 2}, {39, 2}, {40, 0}, {50, 3}, {59, 3}, {60, 0}} {
		rank, ok := got["x"].rankAt(7, c.level)
		if ok != (c.rank != 0) || rank != c.rank {
			t.Errorf("rankAt(7, %d) = (%d, %v), want %d", c.level, rank, ok, c.rank)
		}
	}
}

// A row may name a preset: it then applies only to requests simmed under
// it (round 9: bare 60 and raid 60 run dry differently). A row with no
// preset applies to both.
func TestRankChoicesFollowThePreset(t *testing.T) {
	all, err := parseRankChoices([]byte(`{"specs":{"x":[{"authored":7,"bands":[
		{"from_level":50,"to_level":59,"rank":3},
		{"from_level":60,"to_level":60,"rank":2,"preset":"bare"}]}]}}`))
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct {
		preset    string
		level     int
		rank      int
		wantFound bool
	}{
		{BarePreset, 55, 3, true},
		{RaidPreset, 55, 3, true},
		{BarePreset, 60, 2, true},
		{RaidPreset, 60, 0, false},
	} {
		rank, ok := all["x"].forRequest(c.preset, HordeFaction).rankAt(7, c.level)
		if ok != c.wantFound || rank != c.rank {
			t.Errorf("%s at %d: rankAt = (%d, %v), want (%d, %v)", c.preset, c.level, rank, ok, c.rank, c.wantFound)
		}
	}
	if _, err := parseRankChoices([]byte(`{"specs":{"x":[{"authored":7,"bands":[{"from_level":60,"to_level":60,"rank":2,"preset":"heroic"}]}]}}`)); err == nil {
		t.Error("a row naming an unknown preset was accepted")
	}
}

// At 60 the rotation runs as written unless a row for the request's preset
// names one of its ids; then only that id changes.
func TestRewriteAtSixtyTakesOnlyThePresetsRows(t *testing.T) {
	raw, err := aplFS.ReadFile("apl/shaman-elemental.apl.json")
	if err != nil {
		t.Fatal(err)
	}
	choices := rankChoices{15207: {{FromLevel: 60, ToLevel: 60, Rank: 4, Preset: BarePreset}}}
	bare, err := rewriteRotationRanks(raw, "shaman", api.MaxLevel, choices.forRequest(BarePreset, HordeFaction))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(bare), `"spellId":915`) || strings.Contains(string(bare), `"spellId":15207`) {
		t.Error("bare 60 should cast Lightning Bolt rank 4 (915) in place of 15207")
	}
	raid, err := rewriteRotationRanks(raw, "shaman", api.MaxLevel, choices.forRequest(RaidPreset, HordeFaction))
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(unmarshalGeneric(t, raid), unmarshalGeneric(t, raw)) {
		t.Error("raid 60 has no row and must run the rotation as written")
	}
}

func TestBuildRefusesAnUnknownPreset(t *testing.T) {
	req := api.SimRequest{
		EngineVersion: enginever.Version, Spec: "shaman-elemental",
		Source:    api.CharacterSource{Kind: api.SourceManual},
		Character: api.CharacterSpec{Name: "x", Race: "orc", Class: "shaman", Level: 60},
		Encounter: api.DefaultEncounter(), Iterations: 3000, RandomSeed: 1,
		Preset: "heroic",
	}
	if _, err := Build(req); !errors.Is(err, ErrUnknownPreset) {
		t.Errorf("Build with preset %q: err = %v, want ErrUnknownPreset", req.Preset, err)
	}
}

// A row may name a faction (round 12): it then applies only to a character of
// that faction, where the two factions' best ranks differ beyond error. A row
// with none applies to both, and an Alliance and a Horde row may share a band.
func TestRankChoicesFollowTheFaction(t *testing.T) {
	all, err := parseRankChoices([]byte(`{"specs":{"x":[{"authored":7,"bands":[
		{"from_level":50,"to_level":59,"rank":3},
		{"from_level":60,"to_level":60,"rank":2,"preset":"bare","faction":"alliance"},
		{"from_level":60,"to_level":60,"rank":4,"preset":"bare","faction":"horde"}]}]}}`))
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct {
		preset, faction string
		level, rank     int
		wantFound       bool
	}{
		{BarePreset, AllianceFaction, 55, 3, true},
		{BarePreset, HordeFaction, 55, 3, true},
		{BarePreset, AllianceFaction, 60, 2, true},
		{BarePreset, HordeFaction, 60, 4, true},
		{RaidPreset, AllianceFaction, 60, 0, false},
	} {
		rank, ok := all["x"].forRequest(c.preset, c.faction).rankAt(7, c.level)
		if ok != c.wantFound || rank != c.rank {
			t.Errorf("%s %s at %d: rankAt = (%d, %v), want (%d, %v)", c.preset, c.faction, c.level, rank, ok, c.rank, c.wantFound)
		}
	}
	for _, raw := range []string{
		`{"specs":{"x":[{"authored":7,"bands":[{"from_level":60,"to_level":60,"rank":2,"preset":"bare","faction":"neutral"}]}]}}`,
		`{"specs":{"x":[{"authored":7,"bands":[{"from_level":60,"to_level":60,"rank":2,"preset":"bare","faction":"horde"},{"from_level":60,"to_level":60,"rank":3,"preset":"bare","faction":"horde"}]}]}}`,
		`{"specs":{"x":[{"authored":7,"bands":[{"from_level":60,"to_level":60,"rank":2,"preset":"bare"},{"from_level":60,"to_level":60,"rank":3,"preset":"bare","faction":"horde"}]}]}}`,
	} {
		if _, err := parseRankChoices([]byte(raw)); err == nil {
			t.Errorf("accepted %s", raw)
		}
	}
}

// Every race the request maps has a faction, and Skyborne's two slugs follow
// the faction their racial is chosen for.
func TestFactionOfRaceCoversEveryRace(t *testing.T) {
	want := map[string]string{
		"dwarf": AllianceFaction, "gnome": AllianceFaction, "human": AllianceFaction,
		"night-elf": AllianceFaction, "high-order-skyborne": AllianceFaction,
		"orc": HordeFaction, "tauren": HordeFaction, "troll": HordeFaction,
		"undead": HordeFaction, "windshaper-skyborne": HordeFaction,
	}
	for race := range races {
		got, err := FactionOfRace(race)
		if err != nil || got != want[race] {
			t.Errorf("FactionOfRace(%q) = (%q, %v), want %q", race, got, err, want[race])
		}
	}
	if _, err := FactionOfRace("murloc"); err == nil {
		t.Error("an unknown race has no faction")
	}
}
