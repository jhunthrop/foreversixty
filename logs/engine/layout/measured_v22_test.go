// logs/engine/layout/measured_v22_test.go
package layout

import (
	"os"
	"path/filepath"
	"sort"
	"testing"
)

// v22Testdata is the committed v22 excerpt and the shapes file beside it.
// They live in the event package's testdata because the decoder tests read
// them too; this package reads them to pin the measurement the retail-v22
// row is built from.
var v22Testdata = []string{
	filepath.Join("..", "event", "testdata", "v22.log"),
	filepath.Join("..", "event", "testdata", "v22-shapes.log"),
}

// v22MeasuredWidths is the field count of every event shape the committed
// testdata must contain. It is the subset of the 89-file corpus measurement
// (26,030,980 lines) that the excerpts were cut to cover, and it is what the
// retail-v22 row's widths are checked against in TestRetailV22Widths.
//
// A shape dropped from the testdata is a shape nothing tests, so this list
// failing is a real failure, not a chore.
var v22MeasuredWidths = map[string][]int{
	"ARENA_MATCH_END":               {5},
	"ARENA_MATCH_START":             {5},
	"COMBATANT_INFO":                {34},
	"COMBAT_LOG_VERSION":            {8},
	"DAMAGE_SHIELD_MISSED":          {15},
	"DAMAGE_SPLIT":                  {42},
	"EMOTE":                         {6},
	"ENCHANT_APPLIED":               {12},
	"ENCHANT_REMOVED":               {12},
	"ENVIRONMENTAL_DAMAGE":          {39},
	"MAP_CHANGE":                    {7},
	"PARTY_KILL":                    {10},
	"RANGE_DAMAGE":                  {42},
	"RANGE_DAMAGE_SUPPORT":          {42},
	"RANGE_MISSED":                  {14, 17},
	"SPELL_ABSORBED":                {19, 22},
	"SPELL_ABSORBED_SUPPORT":        {20, 23},
	"SPELL_AURA_APPLIED":            {13, 14, 15},
	"SPELL_AURA_APPLIED_DOSE":       {14},
	"SPELL_AURA_BROKEN":             {13},
	"SPELL_AURA_BROKEN_SPELL":       {16},
	"SPELL_AURA_REFRESH":            {13, 14},
	"SPELL_AURA_REMOVED":            {13, 14, 15},
	"SPELL_AURA_REMOVED_DOSE":       {14},
	"SPELL_CAST_FAILED":             {13},
	"SPELL_CAST_START":              {12},
	"SPELL_CAST_SUCCESS":            {31},
	"SPELL_CREATE":                  {12},
	"SPELL_DAMAGE":                  {42},
	"SPELL_DAMAGE_SUPPORT":          {42},
	"SPELL_DISPEL":                  {16},
	"SPELL_DRAIN":                   {35},
	"SPELL_EMPOWER_END":             {13},
	"SPELL_EMPOWER_INTERRUPT":       {13},
	"SPELL_EMPOWER_START":           {12},
	"SPELL_ENERGIZE":                {35},
	"SPELL_EXTRA_ATTACKS":           {13},
	"SPELL_HEAL":                    {36},
	"SPELL_HEAL_ABSORBED":           {21},
	"SPELL_HEAL_SUPPORT":            {37},
	"SPELL_INSTAKILL":               {13},
	"SPELL_INTERRUPT":               {15},
	"SPELL_MISSED":                  {15, 16, 18},
	"SPELL_PERIODIC_DAMAGE":         {42},
	"SPELL_PERIODIC_DAMAGE_SUPPORT": {42},
	"SPELL_PERIODIC_ENERGIZE":       {35},
	"SPELL_PERIODIC_HEAL":           {36},
	"SPELL_PERIODIC_HEAL_SUPPORT":   {37},
	"SPELL_PERIODIC_MISSED":         {15, 18},
	"SPELL_RESURRECT":               {12},
	"SPELL_STOLEN":                  {16},
	"SPELL_SUMMON":                  {12},
	"STAGGER_CLEAR":                 {3},
	"STAGGER_PREVENTED":             {4},
	"SWING_DAMAGE":                  {38},
	"SWING_DAMAGE_LANDED":           {38},
	"SWING_DAMAGE_LANDED_SUPPORT":   {42},
	"SWING_MISSED":                  {11, 14},
	"UNIT_DIED":                     {10},
	"WORLD_MARKER_PLACED":           {5},
	"WORLD_MARKER_REMOVED":          {2},
	"ZONE_CHANGE":                   {4},
}

// v22TestdataWidths lexes the committed excerpts and returns every
// (event, width) pair they contain.
func v22TestdataWidths(t *testing.T) map[string][]int {
	t.Helper()
	seen := map[string]map[int]bool{}
	for _, path := range v22Testdata {
		text, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		for _, ln := range lines(t, string(text)) {
			if len(ln.Params) == 0 || ln.Params[0] == "" {
				continue
			}
			ev := ln.Params[0]
			if seen[ev] == nil {
				seen[ev] = map[int]bool{}
			}
			seen[ev][len(ln.Params)] = true
		}
	}
	out := map[string][]int{}
	for ev, ws := range seen {
		for w := range ws {
			out[ev] = append(out[ev], w)
		}
		sort.Ints(out[ev])
	}
	return out
}

func TestTheV22TestdataCoversEveryMeasuredShape(t *testing.T) {
	got := v22TestdataWidths(t)
	for ev, want := range v22MeasuredWidths {
		have := got[ev]
		for _, w := range want {
			found := false
			for _, g := range have {
				if g == w {
					found = true
				}
			}
			if !found {
				t.Errorf("the v22 testdata has no %s line of width %d (it has %v)", ev, w, have)
			}
		}
	}
}

func TestTheV22TestdataHasNoShapeTheMeasurementDoesNotName(t *testing.T) {
	for ev, have := range v22TestdataWidths(t) {
		want, known := v22MeasuredWidths[ev]
		if !known {
			t.Errorf("the v22 testdata contains %s, which the measurement does not name", ev)
			continue
		}
		for _, g := range have {
			found := false
			for _, w := range want {
				if g == w {
					found = true
				}
			}
			if !found {
				t.Errorf("the v22 testdata has a %s line of width %d; measured widths are %v", ev, g, want)
			}
		}
	}
}

func TestTheV22ExcerptHeaderIsVersion22(t *testing.T) {
	text, err := os.ReadFile(v22Testdata[0])
	if err != nil {
		t.Fatal(err)
	}
	ls := lines(t, string(text))
	h, ok := ParseHeader(ls[0])
	if !ok {
		t.Fatal("the first line of the v22 excerpt is not a COMBAT_LOG_VERSION line")
	}
	if h.Version != 22 || !h.Advanced || h.ProjectID != 1 || h.Fields != 8 {
		t.Fatalf("header = %+v, want version 22, advanced, project 1, 8 fields", h)
	}
	if h.Build != "12.0.7" {
		t.Errorf("build = %q, want 12.0.7", h.Build)
	}
}

func TestTheV22ExcerptTimestampsCarryAYearAndAZone(t *testing.T) {
	text, err := os.ReadFile(v22Testdata[0])
	if err != nil {
		t.Fatal(err)
	}
	year, zone := inferStamp(lines(t, string(text))[0].Stamp)
	if !year || !zone {
		t.Fatalf("year=%v zone=%v, want both true", year, zone)
	}
}
