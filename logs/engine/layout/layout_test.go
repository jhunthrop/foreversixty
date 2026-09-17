// logs/engine/layout/layout_test.go
package layout

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/logs/engine/lexer"
)

func lines(t *testing.T, text string) []lexer.Line {
	t.Helper()
	l := lexer.New()
	var out []lexer.Line
	emit := func(ln lexer.Line) error { out = append(out, ln); return nil }
	if err := l.Feed([]byte(text), 0, emit); err != nil {
		t.Fatal(err)
	}
	if err := l.Flush(emit); err != nil {
		t.Fatal(err)
	}
	return out
}

func TestParseHeaderReadsTheRetailV16Line(t *testing.T) {
	ls := lines(t, "9/26 20:10:00.000  COMBAT_LOG_VERSION,16,ADVANCED_LOG_ENABLED,1,BUILD_VERSION,9.0.2,PROJECT_ID,1\n")
	h, ok := ParseHeader(ls[0])
	if !ok {
		t.Fatal("ParseHeader returned false")
	}
	if h.Version != 16 || !h.Advanced || h.Build != "9.0.2" || h.ProjectID != 1 || h.Fields != 8 {
		t.Fatalf("header = %+v", h)
	}
}

func TestParseHeaderRejectsAnyOtherLine(t *testing.T) {
	ls := lines(t, "9/26 20:10:00.000  ZONE_CHANGE,2284,\"Sanguine Depths\",8\n")
	if _, ok := ParseHeader(ls[0]); ok {
		t.Fatal("ParseHeader accepted a ZONE_CHANGE line")
	}
}

func TestLookupPicksRetailV16AndHonoursTheAdvancedFlag(t *testing.T) {
	on, ok := Lookup(Header{Version: 16, ProjectID: 1, Advanced: true})
	if !ok || on.Name != "retail-v16" || on.Advanced != 17 {
		t.Fatalf("advanced on: %+v ok=%v", on.Name, ok)
	}
	off, ok := Lookup(Header{Version: 16, ProjectID: 1, Advanced: false})
	if !ok || off.Advanced != 0 {
		t.Fatalf("advanced off: advanced=%d ok=%v", off.Advanced, ok)
	}
	if _, ok := Lookup(Header{Version: 23, ProjectID: 1, Advanced: true}); ok {
		t.Fatal("Lookup matched a version with no row; v23 is not implemented")
	}
}

func TestLookupNeverAutoSelectsTheZeroVersionClassicRow(t *testing.T) {
	if _, ok := Lookup(Header{}); ok {
		t.Fatal("Lookup(Header{}) matched a row; a zero-value header must fall through to inference")
	}
	// A header whose COMBAT_LOG_VERSION field failed to parse also comes
	// through as Version: 0 and must not silently select classic-wiki.
	if _, ok := Lookup(Header{Version: 0, ProjectID: 1, Advanced: true}); ok {
		t.Fatal("Lookup matched the zero-version row for an unparseable version")
	}
}

func TestRetailV16WidthsMatchTheVerifiedCounts(t *testing.T) {
	l := RetailV16()
	for _, tc := range []struct {
		event     string
		wantWidth int
		wantAdvAt int
	}{
		{"SPELL_DAMAGE", 39, 12},
		{"SPELL_PERIODIC_DAMAGE", 39, 12},
		{"RANGE_DAMAGE", 39, 12},
		{"SWING_DAMAGE", 36, 9},
		{"SWING_DAMAGE_LANDED", 36, 9},
		{"SPELL_HEAL", 34, 12},
		{"SPELL_PERIODIC_HEAL", 34, 12},
		{"SPELL_ENERGIZE", 33, 12},
		{"SPELL_CAST_SUCCESS", 29, 12},
		{"SPELL_CAST_START", 12, -1},
		{"SPELL_CAST_FAILED", 13, -1},
		{"SPELL_AURA_APPLIED", 13, -1},
		{"SPELL_AURA_APPLIED_DOSE", 14, -1},
		{"SPELL_INTERRUPT", 15, -1},
		{"SPELL_DISPEL", 16, -1},
		{"SPELL_AURA_BROKEN_SPELL", 16, -1},
		{"SPELL_SUMMON", 12, -1},
		{"SPELL_INSTAKILL", 13, -1},
		{"SWING_MISSED", 11, -1},
		{"SPELL_MISSED", 14, -1},
	} {
		t.Run(tc.event, func(t *testing.T) {
			prefix, suffix, ok := l.Split(tc.event)
			if !ok {
				t.Fatalf("Split(%q) = %q %q, not known", tc.event, prefix, suffix)
			}
			w, advAt := l.Width(prefix, suffix)
			if w != tc.wantWidth || advAt != tc.wantAdvAt {
				t.Fatalf("width = %d advAt = %d, want %d and %d", w, advAt, tc.wantWidth, tc.wantAdvAt)
			}
		})
	}
}

func TestRetailV16SpecialWidths(t *testing.T) {
	l := RetailV16()
	for event, widths := range map[string][]int{
		"UNIT_DIED":            {10},
		"PARTY_KILL":           {10},
		"SPELL_ABSORBED":       {19, 22},
		"SPELL_HEAL_ABSORBED":  {21},
		"ENCOUNTER_START":      {6},
		"ENCOUNTER_END":        {6},
		"ZONE_CHANGE":          {4},
		"MAP_CHANGE":           {7},
		"COMBATANT_INFO":       {34},
		"ENVIRONMENTAL_DAMAGE": {37},
		"CHALLENGE_MODE_START": {6},
		"CHALLENGE_MODE_END":   {5},
		"ENCHANT_APPLIED":      {12},
		"EMOTE":                {6},
	} {
		s, ok := l.Specials[event]
		if !ok {
			t.Errorf("%s is not a special on the retail row", event)
			continue
		}
		for _, w := range widths {
			if !s.Accepts(w) {
				t.Errorf("%s does not accept width %d", event, w)
			}
		}
		if s.Accepts(widths[0] + 100) {
			t.Errorf("%s accepted an absurd width", event)
		}
	}
}

func TestClassicRowDiffersFromRetailWhereTheWikiSaysItDoes(t *testing.T) {
	r, c := RetailV16(), ClassicWiki()
	if r.Suffixes["_DAMAGE"].Params != 10 || c.Suffixes["_DAMAGE"].Params != 9 {
		t.Errorf("_DAMAGE retail=%d classic=%d, want 10 and 9",
			r.Suffixes["_DAMAGE"].Params, c.Suffixes["_DAMAGE"].Params)
	}
	if !c.Suffixes["_DAMAGE"].OffHand {
		t.Error("the Classic row must allow a trailing isOffHand on swings")
	}
	if r.Suffixes["_HEAL"].Params != 5 || c.Suffixes["_HEAL"].Params != 4 {
		t.Errorf("_HEAL retail=%d classic=%d, want 5 and 4",
			r.Suffixes["_HEAL"].Params, c.Suffixes["_HEAL"].Params)
	}
	if r.Suffixes["_MISSED"].AbsorbExtra != 3 || c.Suffixes["_MISSED"].AbsorbExtra != 2 {
		t.Errorf("_MISSED ABSORB extras retail=%d classic=%d, want 3 and 2",
			r.Suffixes["_MISSED"].AbsorbExtra, c.Suffixes["_MISSED"].AbsorbExtra)
	}
	if !r.Suffixes["_DAMAGE"].BaseAmount || c.Suffixes["_DAMAGE"].BaseAmount {
		t.Error("only the retail row carries baseAmount on damage")
	}
	if !r.Suffixes["_HEAL"].HealedToHP || c.Suffixes["_HEAL"].HealedToHP {
		t.Error("only the retail row carries healedToHP on heals")
	}
	if !r.Combatant.Present {
		t.Error("retail v16 must decode COMBATANT_INFO")
	}
	if c.Combatant.Present {
		t.Error("the Classic row must not claim a COMBATANT_INFO layout")
	}
	if !r.Verified || c.Verified {
		t.Error("retail v16 is verified against a real log; the Classic row is not")
	}
}

func TestParseStampWithoutAYearRollsOverAndReportsIt(t *testing.T) {
	l := RetailV16()
	prev := time.Date(2026, 12, 31, 23, 59, 0, 0, time.UTC)
	got, rolled, err := l.ParseStamp("1/1 00:00:01.250", prev)
	if err != nil {
		t.Fatal(err)
	}
	if !rolled {
		t.Error("rolled = false, want true across new year")
	}
	want := time.Date(2027, 1, 1, 0, 0, 1, 250000000, time.UTC)
	if !got.Equal(want) {
		t.Fatalf("got %s, want %s", got, want)
	}
}

func TestParseStampKeepsASmallBackwardsJumpInTheSameYear(t *testing.T) {
	l := RetailV16()
	prev := time.Date(2026, 9, 26, 2, 0, 0, 0, time.UTC)
	got, rolled, err := l.ParseStamp("9/26 01:00:00.000", prev)
	if err != nil {
		t.Fatal(err)
	}
	if rolled {
		t.Error("a one-hour backwards jump must not roll the year")
	}
	if got.Year() != 2026 || got.Hour() != 1 {
		t.Fatalf("got %s", got)
	}
}

func TestParseStampRejectsAMismatchedDate(t *testing.T) {
	l := RetailV16()
	if _, _, err := l.ParseStamp("9/26/2026 01:00:00.000", time.Date(2026, 9, 26, 0, 0, 0, 0, time.UTC)); err == nil {
		t.Fatal("a row with StampYear false must reject a dated timestamp")
	}
	withYear := RetailV16()
	withYear.StampYear = true
	if _, _, err := withYear.ParseStamp("9/26 01:00:00.000", time.Time{}); err == nil {
		t.Fatal("a row with StampYear true must reject an undated timestamp")
	}
}

func TestParseStampReadsAYearAndAZone(t *testing.T) {
	l := RetailV16()
	l.StampYear, l.StampZone = true, true
	got, _, err := l.ParseStamp("9/26/2026 01:02:03.400-4", time.Time{})
	if err != nil {
		t.Fatal(err)
	}
	_, off := got.Zone()
	if off != -4*3600 {
		t.Fatalf("zone offset = %d seconds, want -14400", off)
	}
	if got.UTC().Hour() != 5 {
		t.Fatalf("got %s, want 05:02:03 UTC", got.UTC())
	}
}

func TestParseStampNeedsAPreviousLineWhenThereIsNoYear(t *testing.T) {
	l := RetailV16()
	if _, _, err := l.ParseStamp("9/26 01:00:00.000", time.Time{}); err == nil {
		t.Fatal("want an error when there is neither a year in the line nor a previous line")
	}
}

func TestInferCountsParametersFromTheLogItself(t *testing.T) {
	text := strings.Join([]string{
		`9/26 20:10:00.000  COMBAT_LOG_VERSION,99,ADVANCED_LOG_ENABLED,1,BUILD_VERSION,4.0.0,PROJECT_ID,7`,
		`9/26 20:10:01.000  SPELL_CAST_SUCCESS,Player-1-A,"Baelgrim",0x511,0x0,Creature-0-1-2-3-4-5,"Sentinel",0xa48,0x0,116,"Frostbolt",0x10,Player-1-A,0000000000000000,1,2,3,4,5,6,7,8,9,10,1.0,2.0,11,3.0,12`,
		`9/26 20:10:02.000  SPELL_DAMAGE,Player-1-A,"Baelgrim",0x511,0x0,Creature-0-1-2-3-4-5,"Sentinel",0xa48,0x0,116,"Frostbolt",0x10,Creature-0-1-2-3-4-5,0000000000000000,1,2,3,4,5,6,7,8,9,10,1.0,2.0,11,3.0,12,500,-1,16,0,0,0,nil,nil,nil`,
		`9/26 20:10:03.000  SPELL_AURA_APPLIED,Player-1-A,"Baelgrim",0x511,0x0,Player-1-A,"Baelgrim",0x511,0x0,17,"Shield",0x2,BUFF`,
		`9/26 20:10:04.000  MADE_UP_EVENT,a,b,c`,
		"",
	}, "\n")
	l := Infer(lines(t, text))
	if l.Name != "inferred" || l.Verified {
		t.Fatalf("name=%q verified=%v", l.Name, l.Verified)
	}
	if l.Version != 99 || l.ProjectID != 7 {
		t.Errorf("version=%d project=%d", l.Version, l.ProjectID)
	}
	if l.Advanced != 17 {
		t.Fatalf("advanced = %d, want 17 from the SPELL_CAST_SUCCESS width", l.Advanced)
	}
	if got := l.Suffixes["_DAMAGE"].Params; got != 9 {
		t.Errorf("_DAMAGE params = %d, want 9", got)
	}
	if got := l.Suffixes["_AURA_APPLIED"].Params; got != 1 {
		t.Errorf("_AURA_APPLIED params = %d, want 1", got)
	}
	if s, ok := l.Specials["MADE_UP_EVENT"]; !ok || !s.Accepts(4) {
		t.Errorf("unknown event was not recorded as a special: %+v", l.Specials["MADE_UP_EVENT"])
	}
	if l.StampYear || l.StampZone {
		t.Error("the sample has no year and no zone")
	}
}

func TestInferDetectsAYearAndAZoneInTheTimestamp(t *testing.T) {
	l := Infer(lines(t, "9/26/2026 20:10:00.000-4  ZONE_CHANGE,1,\"Zone\",0\n"))
	if !l.StampYear || !l.StampZone {
		t.Fatalf("year=%v zone=%v, want both true", l.StampYear, l.StampZone)
	}
}

func TestEventNamesIsSortedAndIncludesSpecials(t *testing.T) {
	got := RetailV16().EventNames()
	if len(got) == 0 {
		t.Fatal("no event names")
	}
	for i := 1; i < len(got); i++ {
		if got[i-1] >= got[i] {
			t.Fatalf("not sorted at %d: %q then %q", i, got[i-1], got[i])
		}
	}
	found := false
	for _, n := range got {
		if n == "COMBATANT_INFO" {
			found = true
		}
	}
	if !found {
		t.Error("COMBATANT_INFO missing from EventNames")
	}
}

// twoEventsSharingTheDamageSuffix is the shape that used to make Infer
// nondeterministic: SPELL_CAST_SUCCESS fixes the advanced block at 17,
// SPELL_DAMAGE then derives _DAMAGE{Params: 11, Advanced: true}, and
// SWING_DAMAGE is too narrow to hold an advanced block so it falls back to
// _DAMAGE{Params: 11, Advanced: false}. The two agree on Params and
// disagree on Advanced, so the reduction's Params guard does not fire and
// whichever event the loop visited last used to win.
const twoEventsSharingTheDamageSuffix = `9/26 20:10:01.000  SPELL_CAST_SUCCESS,Player-4184-000000A1,"Baelgrim-Nightslayer",0x511,0x0,Creature-0-2085-2284-7855-169753-0000AA0001,"Hollow Sentinel",0xa48,0x0,116,"Frostbolt",0x10,Player-4184-000000A1,0000000000000000,1,2,3,4,5,6,7,8,9,10,1.0,2.0,11,3.0,12
9/26 20:10:02.000  SPELL_DAMAGE,Player-4184-000000A1,"Baelgrim-Nightslayer",0x511,0x0,Creature-0-2085-2284-7855-169753-0000AA0001,"Hollow Sentinel",0xa48,0x0,116,"Frostbolt",0x10,Creature-0-2085-2284-7855-169753-0000AA0001,0000000000000000,1,2,3,4,5,6,7,8,9,10,1.0,2.0,11,3.0,12,500,499,-1,16,0,0,0,nil,nil,nil,nil
9/26 20:10:03.000  SWING_DAMAGE,Player-4184-000000A1,"Baelgrim-Nightslayer",0x511,0x0,Creature-0-2085-2284-7855-169753-0000AA0001,"Hollow Sentinel",0xa48,0x0,300,299,-1,1,0,0,0,nil,nil,nil,nil
`

// renderSuffixes writes the suffix map out in key order, so two layouts can
// be compared as text without depending on map order.
func renderSuffixes(l Layout) string {
	var b strings.Builder
	for _, k := range sortedKeys(l.Suffixes) {
		fmt.Fprintf(&b, "%s=%+v\n", k, l.Suffixes[k])
	}
	return b.String()
}

func TestInferIsDeterministicWhenTwoEventsShareASuffix(t *testing.T) {
	ls := lines(t, twoEventsSharingTheDamageSuffix)

	// Sanity: this input really does drive both branches into the same
	// Params, which is the precondition for the bug.
	first := Infer(ls)
	if first.Advanced != 17 {
		t.Fatalf("advanced = %d, want 17", first.Advanced)
	}
	if got := first.Suffixes["_DAMAGE"].Params; got != 11 {
		t.Fatalf("_DAMAGE params = %d, want 11 from both events", got)
	}

	seen := map[string]int{}
	advanced := map[bool]int{}
	for i := 0; i < 100; i++ {
		l := Infer(ls)
		seen[renderSuffixes(l)]++
		advanced[l.Suffixes["_DAMAGE"].Advanced]++
	}
	if len(seen) != 1 {
		t.Fatalf("Infer produced %d distinct suffix maps over 100 runs of the same input: %v", len(seen), seen)
	}
	if len(advanced) != 1 {
		t.Fatalf("_DAMAGE.Advanced flipped across runs: %v", advanced)
	}
	if _, ok := advanced[false]; !ok {
		t.Errorf("_DAMAGE.Advanced = true, want the value the alphabetically last writer (SWING_DAMAGE) sets")
	}
}

func TestInferIsDeterministicOverTheWholeRow(t *testing.T) {
	ls := lines(t, twoEventsSharingTheDamageSuffix+
		"9/26 20:10:04.000  MADE_UP_EVENT,a,b,c\n"+
		"9/26 20:10:05.000  MADE_UP_EVENT,a,b,c,d\n")
	want, err := json.Marshal(Infer(ls))
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 100; i++ {
		got, err := json.Marshal(Infer(ls))
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(got, want) {
			t.Fatalf("run %d produced a different row:\n got %s\nwant %s", i, got, want)
		}
	}
}

func TestLookupPicksRetailV22(t *testing.T) {
	on, ok := Lookup(Header{Version: 22, ProjectID: 1, Advanced: true})
	if !ok || on.Name != "retail-v22" || on.Advanced != 19 {
		t.Fatalf("advanced on: name=%q advanced=%d ok=%v", on.Name, on.Advanced, ok)
	}
	if !on.Verified {
		t.Error("retail-v22 must be a verified row")
	}
	if !on.StampYear || !on.StampZone {
		t.Errorf("stampYear=%v stampZone=%v, want both true", on.StampYear, on.StampZone)
	}
	off, ok := Lookup(Header{Version: 22, ProjectID: 1, Advanced: false})
	if !ok || off.Advanced != 0 {
		t.Fatalf("advanced off: advanced=%d ok=%v", off.Advanced, ok)
	}
	// v16 must still win for version 16.
	v16, ok := Lookup(Header{Version: 16, ProjectID: 1, Advanced: true})
	if !ok || v16.Name != "retail-v16" || v16.Advanced != 17 {
		t.Fatalf("v16 selection regressed: name=%q advanced=%d", v16.Name, v16.Advanced)
	}
}

// TestRetailV22WidthsMatchTheVerifiedCounts is the v22 twin of
// TestRetailV16WidthsMatchTheVerifiedCounts. wantWidths is the complete set
// the row accepts; the measured widths from the corpus must all be in it.
func TestRetailV22WidthsMatchTheVerifiedCounts(t *testing.T) {
	l := RetailV22()
	for _, tc := range []struct {
		event      string
		wantWidths []int
		wantAdvAt  int
	}{
		{"SPELL_DAMAGE", []int{41, 42}, 12},
		{"SPELL_PERIODIC_DAMAGE", []int{41, 42}, 12},
		{"RANGE_DAMAGE", []int{41, 42}, 12},
		{"DAMAGE_SPLIT", []int{41, 42}, 12},
		{"SWING_DAMAGE", []int{38, 39}, 9},
		{"SWING_DAMAGE_LANDED", []int{38, 39}, 9},
		{"SPELL_DAMAGE_SUPPORT", []int{42}, 12},
		{"SPELL_PERIODIC_DAMAGE_SUPPORT", []int{42}, 12},
		{"RANGE_DAMAGE_SUPPORT", []int{42}, 12},
		{"SWING_DAMAGE_LANDED_SUPPORT", []int{42}, 12},
		{"SPELL_HEAL", []int{36}, 12},
		{"SPELL_PERIODIC_HEAL", []int{36}, 12},
		{"SPELL_HEAL_SUPPORT", []int{37}, 12},
		{"SPELL_PERIODIC_HEAL_SUPPORT", []int{37}, 12},
		{"SPELL_ENERGIZE", []int{35}, 12},
		{"SPELL_PERIODIC_ENERGIZE", []int{35}, 12},
		{"SPELL_DRAIN", []int{35}, 12},
		{"SPELL_CAST_SUCCESS", []int{31}, 12},
		{"SPELL_CAST_START", []int{12}, -1},
		{"SPELL_CAST_FAILED", []int{13}, -1},
		{"SPELL_EMPOWER_START", []int{12}, -1},
		{"SPELL_EMPOWER_END", []int{13}, -1},
		{"SPELL_EMPOWER_INTERRUPT", []int{13}, -1},
		{"SPELL_AURA_APPLIED", []int{13, 14, 15}, -1},
		{"SPELL_AURA_REMOVED", []int{13, 14, 15}, -1},
		{"SPELL_AURA_REFRESH", []int{13, 14, 15}, -1},
		{"SPELL_AURA_APPLIED_DOSE", []int{14}, -1},
		{"SPELL_AURA_BROKEN_SPELL", []int{16}, -1},
		{"SPELL_INTERRUPT", []int{15}, -1},
		{"SPELL_DISPEL", []int{16}, -1},
		{"SPELL_STOLEN", []int{16}, -1},
		{"SPELL_SUMMON", []int{12}, -1},
		{"SPELL_CREATE", []int{12}, -1},
		{"SPELL_RESURRECT", []int{12}, -1},
		{"SPELL_INSTAKILL", []int{13}, -1},
		{"SPELL_EXTRA_ATTACKS", []int{13}, -1},
		{"SWING_MISSED", []int{11, 12, 13, 14, 15}, -1},
		{"RANGE_MISSED", []int{14, 15, 16, 17, 18}, -1},
		{"SPELL_MISSED", []int{14, 15, 16, 17, 18}, -1},
		{"SPELL_PERIODIC_MISSED", []int{14, 15, 16, 17, 18}, -1},
		{"DAMAGE_SHIELD_MISSED", []int{14, 15, 16, 17, 18}, -1},
	} {
		t.Run(tc.event, func(t *testing.T) {
			prefix, suffix, ok := l.Split(tc.event)
			if !ok {
				t.Fatalf("Split(%q) = %q %q, not known", tc.event, prefix, suffix)
			}
			_, advAt := l.Width(prefix, suffix)
			if advAt != tc.wantAdvAt {
				t.Errorf("advAt = %d, want %d", advAt, tc.wantAdvAt)
			}
			got := l.Widths(prefix, suffix)
			if len(got) != len(tc.wantWidths) {
				t.Fatalf("widths = %v, want %v", got, tc.wantWidths)
			}
			for i := range got {
				if got[i] != tc.wantWidths[i] {
					t.Fatalf("widths = %v, want %v", got, tc.wantWidths)
				}
			}
		})
	}
}

// TestRetailV22AcceptsEveryMeasuredWidth ties the row back to the corpus
// measurement: every (event, width) pair seen in 26,030,980 real lines must
// be a width the row accepts, whether the event is a prefix/suffix shape or
// a special.
func TestRetailV22AcceptsEveryMeasuredWidth(t *testing.T) {
	l := RetailV22()
	for event, widths := range v22MeasuredWidths {
		if event == "COMBAT_LOG_VERSION" {
			continue // the header is parsed before the row is consulted
		}
		for _, w := range widths {
			if s, ok := l.Specials[event]; ok {
				if !s.Accepts(w) {
					t.Errorf("special %s does not accept measured width %d (accepts %v)", event, w, s.Widths)
				}
				continue
			}
			prefix, suffix, known := l.Split(event)
			if !known {
				t.Errorf("%s is neither a special nor a known prefix/suffix on retail-v22", event)
				continue
			}
			found := false
			for _, got := range l.Widths(prefix, suffix) {
				if got == w {
					found = true
				}
			}
			if !found {
				t.Errorf("%s width %d is not accepted (row accepts %v)", event, w, l.Widths(prefix, suffix))
			}
		}
	}
}

// TestRetailV16WidthsAreUnchangedByTheWidthsHelper guards the refactor: the
// set of widths the v16 row accepts must be exactly what the decoder
// accepted before Widths existed.
func TestRetailV16WidthsAreUnchangedByTheWidthsHelper(t *testing.T) {
	l := RetailV16()
	for _, tc := range []struct {
		event string
		want  []int
	}{
		{"SPELL_DAMAGE", []int{39}},
		{"SWING_DAMAGE", []int{36}},
		{"SPELL_HEAL", []int{34}},
		{"SPELL_ENERGIZE", []int{33}},
		{"SPELL_CAST_SUCCESS", []int{29}},
		{"SPELL_MISSED", []int{14, 17}},
		{"SWING_MISSED", []int{11, 14}},
		{"SPELL_AURA_APPLIED", []int{13, 14}},
		{"SPELL_AURA_APPLIED_DOSE", []int{14}},
		{"SPELL_AURA_BROKEN_SPELL", []int{16}},
	} {
		t.Run(tc.event, func(t *testing.T) {
			prefix, suffix, ok := l.Split(tc.event)
			if !ok {
				t.Fatalf("Split(%q) failed", tc.event)
			}
			got := l.Widths(prefix, suffix)
			if len(got) != len(tc.want) {
				t.Fatalf("widths = %v, want %v", got, tc.want)
			}
			for i := range got {
				if got[i] != tc.want[i] {
					t.Fatalf("widths = %v, want %v", got, tc.want)
				}
			}
		})
	}
}

// TestSplitPrefersThePrefixThatLeavesAKnownSuffix is the rule that lets the
// v22 row register both SWING and SWING_DAMAGE_LANDED without the longer
// one swallowing the shorter one's events.
func TestSplitPrefersThePrefixThatLeavesAKnownSuffix(t *testing.T) {
	l := RetailV22()
	for _, tc := range []struct{ event, prefix, suffix string }{
		{"SWING_DAMAGE_LANDED", "SWING", "_DAMAGE_LANDED"},
		{"SWING_DAMAGE_LANDED_SUPPORT", "SWING_DAMAGE_LANDED", "_SUPPORT"},
		{"SPELL_DAMAGE", "SPELL", "_DAMAGE"},
		{"SPELL_DAMAGE_SUPPORT", "SPELL_DAMAGE", "_SUPPORT"},
		{"SPELL_PERIODIC_DAMAGE", "SPELL_PERIODIC", "_DAMAGE"},
		{"SPELL_PERIODIC_DAMAGE_SUPPORT", "SPELL_PERIODIC_DAMAGE", "_SUPPORT"},
		{"RANGE_DAMAGE", "RANGE", "_DAMAGE"},
		{"RANGE_DAMAGE_SUPPORT", "RANGE_DAMAGE", "_SUPPORT"},
		{"SPELL_HEAL", "SPELL", "_HEAL"},
		{"SPELL_HEAL_SUPPORT", "SPELL", "_HEAL_SUPPORT"},
		{"DAMAGE_SPLIT", "DAMAGE", "_SPLIT"},
		{"DAMAGE_SHIELD_MISSED", "DAMAGE_SHIELD", "_MISSED"},
	} {
		t.Run(tc.event, func(t *testing.T) {
			prefix, suffix, ok := l.Split(tc.event)
			if !ok || prefix != tc.prefix || suffix != tc.suffix {
				t.Fatalf("Split(%q) = %q %q ok=%v, want %q %q true",
					tc.event, prefix, suffix, ok, tc.prefix, tc.suffix)
			}
		})
	}
}

// TestSplitIsUnchangedForRetailV16 is the other half: the new rule must not
// move a single v16 event.
func TestSplitIsUnchangedForRetailV16(t *testing.T) {
	l := RetailV16()
	for _, tc := range []struct {
		event, prefix, suffix string
		ok                    bool
	}{
		{"SPELL_DAMAGE", "SPELL", "_DAMAGE", true},
		{"SPELL_PERIODIC_DAMAGE", "SPELL_PERIODIC", "_DAMAGE", true},
		{"SPELL_BUILDING_DAMAGE", "SPELL_BUILDING", "_DAMAGE", true},
		{"SWING_DAMAGE_LANDED", "SWING", "_DAMAGE_LANDED", true},
		{"SPELL_AURA_APPLIED_DOSE", "SPELL", "_AURA_APPLIED_DOSE", true},
		{"SPELL_MADE_UP", "SPELL", "_MADE_UP", false},
		{"NOT_AN_EVENT", "", "", false},
	} {
		t.Run(tc.event, func(t *testing.T) {
			prefix, suffix, ok := l.Split(tc.event)
			if ok != tc.ok || prefix != tc.prefix || suffix != tc.suffix {
				t.Fatalf("Split(%q) = %q %q ok=%v, want %q %q %v",
					tc.event, prefix, suffix, ok, tc.prefix, tc.suffix, tc.ok)
			}
		})
	}
}
