// logs/engine/event/decode_v22_test.go
package event

import (
	"os"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/logs/engine/layout"
	"github.com/jhunthrop/foreversixty/logs/engine/lexer"
)

// v22Base is unused by v22 itself: its timestamps carry a year. It is
// passed anyway because NewDecoder wants a base, and a wrong one must not
// be able to leak into a dated dialect.
var v22Base = time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC)

func decodeV22File(t *testing.T, path string) (map[string][]Event, []Event) {
	t.Helper()
	text, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	d := NewDecoder(layout.RetailV22(), v22Base)
	byName := map[string][]Event{}
	var all []Event
	l := lexer.New()
	emit := func(ln lexer.Line) error {
		e := d.Decode(ln)
		byName[e.Name] = append(byName[e.Name], e)
		all = append(all, e)
		return nil
	}
	if err := l.Feed(text, 0, emit); err != nil {
		t.Fatal(err)
	}
	if err := l.Flush(emit); err != nil {
		t.Fatal(err)
	}
	return byName, all
}

func TestEveryV22LineDecodesWithoutError(t *testing.T) {
	for _, path := range []string{"testdata/v22.log", "testdata/v22-shapes.log"} {
		t.Run(path, func(t *testing.T) {
			_, all := decodeV22File(t, path)
			if len(all) == 0 {
				t.Fatal("no lines decoded")
			}
			for _, e := range all {
				if e.Kind == ParseError {
					t.Errorf("line %d (%s) failed: %s\n%s", e.Line, e.Name, e.Error, e.Raw)
				}
				if e.Kind == Unknown {
					t.Errorf("line %d (%s) was not recognised", e.Line, e.Name)
				}
			}
		})
	}
}

func TestV22SpellDamageReadsTheNineteenFieldAdvancedBlockAndTheScopeTag(t *testing.T) {
	by, _ := decodeV22File(t, "testdata/v22.log")
	e := one(t, by, "SPELL_DAMAGE", 0)
	if e.Kind != Damage {
		t.Fatalf("kind = %v, want Damage", e.Kind)
	}
	if !e.Adv.OK {
		t.Fatal("the advanced block was not read")
	}
	if e.Adv.PositionX == 0 || e.Adv.PositionY == 0 {
		t.Errorf("position = %g, %g; both zero means the block was read at the v16 offsets",
			e.Adv.PositionX, e.Adv.PositionY)
	}
	if e.Adv.Level < 200 || e.Adv.Level > 500 {
		t.Errorf("item level = %d, want a plausible level-80 item level", e.Adv.Level)
	}
	if e.Scope != "ST" && e.Scope != "AOE" {
		t.Errorf("scope = %q, want ST or AOE", e.Scope)
	}
	if e.OffHand.OK {
		t.Error("the scope tag was read as isOffHand")
	}
	if !e.Amount.OK || !e.BaseAmount.OK {
		t.Errorf("amount=%+v baseAmount=%+v, both must be read", e.Amount, e.BaseAmount)
	}
}

func TestV22SwingDamageHasNoScopeTag(t *testing.T) {
	by, _ := decodeV22File(t, "testdata/v22.log")
	e := one(t, by, "SWING_DAMAGE", 0)
	if e.Kind != Damage {
		t.Fatalf("kind = %v, want Damage", e.Kind)
	}
	if e.Scope != "" {
		t.Errorf("scope = %q, want empty: SWING lines carry no tag", e.Scope)
	}
	if !e.Adv.OK {
		t.Error("the advanced block was not read")
	}
}

func TestV22PowerFieldsComeFromTheShiftedOffsets(t *testing.T) {
	by, _ := decodeV22File(t, "testdata/v22.log")
	// Index 0 is a Rake energize, which awards combo points: a different
	// power pool than the character's tracked primary resource, so it
	// does not satisfy the invariant below. Index 1 (Force of Nature,
	// which restores mana) does, as does every other energize event in
	// the excerpt.
	e := one(t, by, "SPELL_ENERGIZE", 1)
	if e.Kind != Energize {
		t.Fatalf("kind = %v, want Energize", e.Kind)
	}
	// The energize suffix names the same power the advanced block
	// describes, which is what pinned the block's offsets in the first
	// place.
	if e.Adv.PowerType != e.PowerType.V {
		t.Errorf("advanced powerType = %d, suffix powerType = %d; the block is misaligned",
			e.Adv.PowerType, e.PowerType.V)
	}
	if e.Adv.MaxPower != e.MaxPower.V {
		t.Errorf("advanced maxPower = %d, suffix maxPower = %d; the block is misaligned",
			e.Adv.MaxPower, e.MaxPower.V)
	}
}

func TestV22MissedReadsTheAbsorbExtrasPastTheTag(t *testing.T) {
	by, _ := decodeV22File(t, "testdata/v22.log")
	var absorbed *Event
	for i := range by["SPELL_MISSED"] {
		if by["SPELL_MISSED"][i].MissType == "ABSORB" {
			absorbed = &by["SPELL_MISSED"][i]
			break
		}
	}
	if absorbed == nil {
		t.Fatal("the excerpt has no absorbed SPELL_MISSED")
	}
	if !absorbed.Amount.OK || !absorbed.BaseAmount.OK {
		t.Errorf("amount=%+v baseAmount=%+v; the absorb extras were not read past the tag",
			absorbed.Amount, absorbed.BaseAmount)
	}
	if absorbed.Scope != "ST" && absorbed.Scope != "AOE" {
		t.Errorf("scope = %q, want ST or AOE", absorbed.Scope)
	}
}

func TestV22BlockedMissCarriesAnAmount(t *testing.T) {
	by, _ := decodeV22File(t, "testdata/v22-shapes.log")
	e := one(t, by, "SPELL_MISSED", 0)
	if e.MissType != "BLOCK" && e.MissType != "RESIST" {
		t.Fatalf("the shapes file's SPELL_MISSED is %q, want the 16-wide BLOCK or RESIST form", e.MissType)
	}
	if !e.Amount.OK {
		t.Error("the amount missed was not read")
	}
	if e.Scope != "ST" && e.Scope != "AOE" {
		t.Errorf("scope = %q, want ST or AOE", e.Scope)
	}
}

func TestV22RangeMissedHasNoScopeTag(t *testing.T) {
	by, _ := decodeV22File(t, "testdata/v22-shapes.log")
	e := one(t, by, "RANGE_MISSED", 0)
	if e.Kind != Missed {
		t.Fatalf("kind = %v, want Missed", e.Kind)
	}
	if e.Scope != "" {
		t.Errorf("scope = %q, want empty: RANGE_MISSED carries no tag", e.Scope)
	}
}

func TestV22SupportEventsNameTheSupporter(t *testing.T) {
	by, _ := decodeV22File(t, "testdata/v22-shapes.log")
	for _, name := range []string{
		"SPELL_DAMAGE_SUPPORT", "SPELL_PERIODIC_DAMAGE_SUPPORT",
		"RANGE_DAMAGE_SUPPORT", "SWING_DAMAGE_LANDED_SUPPORT",
		"SPELL_HEAL_SUPPORT", "SPELL_PERIODIC_HEAL_SUPPORT",
		"SPELL_ABSORBED_SUPPORT",
	} {
		t.Run(name, func(t *testing.T) {
			e := one(t, by, name, 0)
			if e.Kind == ParseError || e.Kind == Unknown {
				t.Fatalf("kind = %v: %s", e.Kind, e.Error)
			}
			if e.Supporter == "" {
				t.Error("the supporting player's GUID was not read")
			}
			if e.Scope != "" {
				t.Errorf("scope = %q; a support line ends in a GUID, not a tag", e.Scope)
			}
		})
	}
}

func TestV22DamageSplitIsDamage(t *testing.T) {
	by, _ := decodeV22File(t, "testdata/v22-shapes.log")
	e := one(t, by, "DAMAGE_SPLIT", 0)
	if e.Kind != DamageSplit {
		t.Fatalf("kind = %v, want DamageSplit", e.Kind)
	}
	if e.Spell.ID == 0 || e.Spell.Name == "" {
		t.Errorf("spell = %+v, want the splitting spell", e.Spell)
	}
	if !e.Adv.OK {
		t.Error("the advanced block was not read")
	}
	if !e.Amount.OK {
		t.Error("the amount was not read")
	}
}

func TestV22EmpowerAndArenaAndStaggerAndMarkers(t *testing.T) {
	by, _ := decodeV22File(t, "testdata/v22-shapes.log")
	for name, want := range map[string]Kind{
		"SPELL_EMPOWER_START":     EmpowerStart,
		"SPELL_EMPOWER_END":       EmpowerEnd,
		"SPELL_EMPOWER_INTERRUPT": EmpowerEnd,
		"ARENA_MATCH_END":         ArenaMatchEnd,
		"STAGGER_CLEAR":           StaggerClear,
		"STAGGER_PREVENTED":       StaggerPrevented,
		"WORLD_MARKER_PLACED":     WorldMarker,
		"WORLD_MARKER_REMOVED":    WorldMarker,
	} {
		t.Run(name, func(t *testing.T) {
			e := one(t, by, name, 0)
			if e.Kind != want {
				t.Fatalf("kind = %v, want %v (error: %s)", e.Kind, want, e.Error)
			}
		})
	}
}

// TestV22ArenaMatchStartDecodes constructs its own line: the shapes
// excerpt starts mid-match, an arena log's normal state after a
// reconnect, so it carries ARENA_MATCH_END but no ARENA_MATCH_START.
func TestV22ArenaMatchStartDecodes(t *testing.T) {
	d := NewDecoder(layout.RetailV22(), v22Base)
	l := lexer.New()
	var got Event
	seen := false
	raw := "5/9/2026 12:00:00.000-5  ARENA_MATCH_START,1509,3,Rated Solo Shuffle,1\n"
	emit := func(ln lexer.Line) error {
		got = d.Decode(ln)
		seen = true
		return nil
	}
	if err := l.Feed([]byte(raw), 0, emit); err != nil {
		t.Fatal(err)
	}
	if err := l.Flush(emit); err != nil {
		t.Fatal(err)
	}
	if !seen {
		t.Fatal("no line decoded")
	}
	if got.Kind != ArenaMatchStart {
		t.Fatalf("kind = %v, want ArenaMatchStart (error: %s)", got.Kind, got.Error)
	}
	if got.Zone == nil || got.Zone.ID != 1509 {
		t.Errorf("zone = %+v, want instance id 1509", got.Zone)
	}
	if !got.Amount.OK || got.Amount.V != 3 {
		t.Errorf("amount (bracket) = %+v, want 3", got.Amount)
	}
	if got.ItemName != "Rated Solo Shuffle" {
		t.Errorf("itemName (match type) = %q, want %q", got.ItemName, "Rated Solo Shuffle")
	}
	if !got.Critical.OK || !got.Critical.V {
		t.Errorf("critical (isRated) = %+v, want true", got.Critical)
	}
}

// v22RawLines lexes path and returns every line, so a test can index a raw
// field directly instead of trusting the same decoder it is checking.
func v22RawLines(t *testing.T, path string) []lexer.Line {
	t.Helper()
	text, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	var out []lexer.Line
	l := lexer.New()
	emit := func(ln lexer.Line) error { out = append(out, ln); return nil }
	if err := l.Feed(text, 0, emit); err != nil {
		t.Fatal(err)
	}
	if err := l.Flush(emit); err != nil {
		t.Fatal(err)
	}
	return out
}

func TestV22CombatantInfoReadsTheShiftedStats(t *testing.T) {
	by, all := decodeV22File(t, "testdata/v22.log")
	e := one(t, by, "COMBATANT_INFO", 0)
	if e.Kind != CombatantInfo {
		t.Fatalf("kind = %v: %s", e.Kind, e.Error)
	}
	c := e.Combatant
	if c == nil {
		t.Fatal("no combatant")
	}
	if c.SpecID <= 0 || c.SpecID > 2000 {
		t.Errorf("specID = %d, want a plausible spec id; the index may still be v16's 24", c.SpecID)
	}
	if len(c.Gear) == 0 {
		t.Error("no gear; the gear index is wrong")
	}
	if c.ItemLevel < 200 || c.ItemLevel > 500 {
		t.Errorf("item level = %d, want a plausible level-80 item level", c.ItemLevel)
	}
	// Ruling: version 22's talent trees write (nodeID, entryID, rank)
	// triples, not the flat spell-id tuple Combatant.Talents is typed to
	// hold, so the decoder emits an empty Talents until named triples
	// exist rather than a flattened, wrong-but-typed mix of the three; see
	// docs/ledger/2026-09-16-retail-v22.md. (v16's non-empty Talents is
	// covered by TestCombatantInfoReadsSpecTalentsGearAndAuras.)
	if len(c.Talents) != 0 {
		t.Errorf("talents = %v, want empty for v22", c.Talents)
	}
	if c.Borrowed != "" {
		t.Errorf("borrowed = %q, want empty: v22 writes no borrowed-power field", c.Borrowed)
	}

	// The advanced block's armor is the same number COMBATANT_INFO reports
	// at field 24, which is what pinned the stat shift, and field 23 (v16's
	// armor index) is not: this is checked for every player in the
	// excerpt, not just the first, by indexing the raw COMBATANT_INFO line
	// directly rather than trusting Combatant.Stats, which reads the same
	// index under test.
	// The advanced block describes the event's source, not its target: a
	// destination match would as often as not pick up whoever the player
	// was fighting instead. Armor moves during the fight (stances, auras,
	// trinket procs), so the cross-check is membership in every armor this
	// player's source events ever reported, not equality with just one of
	// them.
	armorsOf := func(guid string) map[int64]bool {
		set := map[int64]bool{}
		for _, ev := range all {
			if ev.Adv.OK && ev.Source.GUID == guid {
				set[ev.Adv.Armor] = true
			}
		}
		return set
	}
	found23, found24 := 0, 0
	for _, ln := range v22RawLines(t, "testdata/v22.log") {
		if len(ln.Params) == 0 || ln.Params[0] != "COMBATANT_INFO" {
			continue
		}
		guid := ln.Params[1]
		armors := armorsOf(guid)
		if len(armors) == 0 {
			t.Fatalf("no advanced-block event sourced by %s to cross-check armor against", guid)
		}
		if len(ln.Params) <= 24 {
			t.Fatalf("COMBATANT_INFO for %s has only %d fields, want at least 25", guid, len(ln.Params))
		}
		if got := intOf(ln.Params[24]); !armors[got] {
			t.Errorf("%s: CI[24] = %d, not among the armor values %s's own events report", guid, got, guid)
		} else {
			found24++
		}
		if got := intOf(ln.Params[23]); armors[got] {
			found23++
		}
	}
	if found24 == 0 {
		t.Fatal("no player's CI[24] was checked against the advanced block")
	}
	if found23 != 0 {
		t.Errorf("%d player(s) had CI[23] land among the advanced block's armor values too; "+
			"the excerpt no longer distinguishes field 23 from field 24", found23)
	}
}

// TestV22AuraWithTwoTrailingNumbers checks both of a 15-field aura line's
// trailing numbers: the first is the absorb size, read into e.Absorbed; the
// second is counted by the layout but not read into any field (see the
// ledger). The fixture line has 0 for both, which proves too little on its
// own, so a second, hand-built line with two distinct nonzero values pins
// that the first is read as the absorb size and the second does not leak
// into it (or into anything else that would make decoding fail).
func TestV22AuraWithTwoTrailingNumbers(t *testing.T) {
	by, _ := decodeV22File(t, "testdata/v22-shapes.log")
	e := one(t, by, "SPELL_AURA_APPLIED", 0)
	if e.Kind != AuraApplied {
		t.Fatalf("kind = %v: %s", e.Kind, e.Error)
	}
	if e.AuraType != "BUFF" && e.AuraType != "DEBUFF" {
		t.Errorf("auraType = %q", e.AuraType)
	}
	if !e.Absorbed.OK || e.Absorbed.V != 0 {
		t.Errorf("absorbed = %+v, want {0 true}: the fixture's absorb size", e.Absorbed)
	}

	d := NewDecoder(layout.RetailV22(), v22Base)
	raw := `4/29/2026 21:04:30.621-5  SPELL_AURA_APPLIED,Player-1129-0BED7A21,"Nightwarrior-Jaedenar-US",0x10548,0x80000000,Player-1129-0BED7A21,"Nightwarrior-Jaedenar-US",0x10548,0x80000000,458245,"Second Wind",0x1,BUFF,500,999` + "\n"
	l := lexer.New()
	var got Event
	emit := func(ln lexer.Line) error { got = d.Decode(ln); return nil }
	if err := l.Feed([]byte(raw), 0, emit); err != nil {
		t.Fatal(err)
	}
	if err := l.Flush(emit); err != nil {
		t.Fatal(err)
	}
	if got.Kind != AuraApplied {
		t.Fatalf("kind = %v: %s", got.Kind, got.Error)
	}
	if !got.Absorbed.OK || got.Absorbed.V != 500 {
		t.Errorf("absorbed = %+v, want {500 true}: the first trailing number", got.Absorbed)
	}
}

func TestTheV16FixtureStillDecodesTheSameWayAfterTheV22Work(t *testing.T) {
	_, all := decodeAll(t)
	if len(all) != 33 {
		t.Fatalf("decoded %d lines, want 33", len(all))
	}
	for _, e := range all {
		if e.Kind == ParseError || e.Kind == Unknown {
			t.Errorf("line %d (%s): kind=%v %s", e.Line, e.Name, e.Kind, e.Error)
		}
		if e.Scope != "" {
			t.Errorf("line %d (%s) got a scope tag on a v16 line", e.Line, e.Name)
		}
		if e.Supporter != "" {
			t.Errorf("line %d (%s) got a supporter on a v16 line", e.Line, e.Name)
		}
	}
}
