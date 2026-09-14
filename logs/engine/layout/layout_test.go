// logs/engine/layout/layout_test.go
package layout

import (
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
	if _, ok := Lookup(Header{Version: 22, ProjectID: 1, Advanced: true}); ok {
		t.Fatal("Lookup matched a version with no row; v22 is not implemented")
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
