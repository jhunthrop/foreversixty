// logs/engine/event/decode_test.go
package event

import (
	"os"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/logs/engine/layout"
	"github.com/jhunthrop/foreversixty/logs/engine/lexer"
)

// fixtureBase is the year the fixture's yearless timestamps belong to. A
// real parse takes this from the log file's modification time.
var fixtureBase = time.Date(2026, 9, 26, 0, 0, 0, 0, time.UTC)

// decodeAll runs the whole fixture through one decoder and indexes the
// result by event name, keeping every occurrence.
func decodeAll(t *testing.T) (map[string][]Event, []Event) {
	t.Helper()
	text, err := os.ReadFile("testdata/v16.log")
	if err != nil {
		t.Fatal(err)
	}
	d := NewDecoder(layout.RetailV16(), fixtureBase)
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

// one returns the nth event of a name, failing when it is missing.
func one(t *testing.T, by map[string][]Event, name string, n int) Event {
	t.Helper()
	es := by[name]
	if len(es) <= n {
		t.Fatalf("fixture has %d %s events, wanted index %d", len(es), name, n)
	}
	return es[n]
}

func TestEveryFixtureLineDecodesWithoutError(t *testing.T) {
	_, all := decodeAll(t)
	if len(all) != 33 {
		t.Fatalf("decoded %d lines, want 33", len(all))
	}
	for _, e := range all {
		if e.Kind == ParseError {
			t.Errorf("line %d (%s) failed: %s\n%s", e.Line, e.Name, e.Error, e.Raw)
		}
		if e.Kind == Unknown {
			t.Errorf("line %d (%s) was not recognised", e.Line, e.Name)
		}
	}
}

func TestHeaderLine(t *testing.T) {
	by, _ := decodeAll(t)
	e := one(t, by, "COMBAT_LOG_VERSION", 0)
	if e.Kind != Header || e.Amount.V != 16 || !e.Critical.V || e.ItemName != "9.0.2" || e.Total.V != 1 {
		t.Fatalf("header = %+v", e)
	}
}

func TestSpellDamageReadsAmountBaseAmountAndTheAdvancedBlock(t *testing.T) {
	by, _ := decodeAll(t)
	e := one(t, by, "SPELL_DAMAGE", 0)
	if e.Kind != Damage {
		t.Fatalf("kind = %s", e.Kind)
	}
	if e.Source.Name != "Morrowlyn-Nightslayer" || e.Dest.Name != "Hollow Sentinel" {
		t.Errorf("units = %q -> %q", e.Source.Name, e.Dest.Name)
	}
	if e.Source.Flags != 0x512 || e.Dest.Flags != 0xa48 {
		t.Errorf("flags = %#x %#x", e.Source.Flags, e.Dest.Flags)
	}
	if e.Spell.ID != 116 || e.Spell.Name != "Frostbolt" || e.Spell.School != 0x10 {
		t.Errorf("spell = %+v", e.Spell)
	}
	if e.Amount.V != 1484 || e.BaseAmount.V != 1390 {
		t.Errorf("amount = %d base = %d, want 1484 and 1390", e.Amount.V, e.BaseAmount.V)
	}
	if e.Overkill.V != -1 || e.School.V != 16 {
		t.Errorf("overkill = %d school = %d", e.Overkill.V, e.School.V)
	}
	if !e.Critical.OK || !e.Critical.V {
		t.Errorf("critical = %+v, want a present true", e.Critical)
	}
	if !e.Glancing.OK || e.Glancing.V {
		t.Errorf("glancing = %+v, want a present false from nil", e.Glancing)
	}
	if !e.Adv.OK || e.Adv.CurrentHP != 41320 || e.Adv.MaxHP != 44000 || e.Adv.Level != 45 {
		t.Errorf("advanced = %+v", e.Adv)
	}
	if e.Adv.UIMapID != 1675 || e.Adv.PositionX != -1487.02 {
		t.Errorf("position = %f %f map %d", e.Adv.PositionX, e.Adv.PositionY, e.Adv.UIMapID)
	}
	if e.Effective() != 1484 {
		t.Errorf("effective = %d", e.Effective())
	}
}

func TestSwingDamageHasNoSpellPrefixAndItsAdvancedBlockDescribesTheNamedUnit(t *testing.T) {
	by, _ := decodeAll(t)
	swing := one(t, by, "SWING_DAMAGE", 0)
	landed := one(t, by, "SWING_DAMAGE_LANDED", 0)
	if swing.Spell.ID != 0 || swing.Spell.Name != "" {
		t.Errorf("a swing must have no spell prefix, got %+v", swing.Spell)
	}
	if swing.Amount.V != 812 || landed.Amount.V != 812 {
		t.Errorf("the pair must report the same swing: %d and %d", swing.Amount.V, landed.Amount.V)
	}
	if swing.Adv.InfoGUID != swing.Source.GUID {
		t.Errorf("SWING_DAMAGE advanced block describes the source, got %q", swing.Adv.InfoGUID)
	}
	if landed.Adv.InfoGUID != landed.Dest.GUID {
		t.Errorf("SWING_DAMAGE_LANDED advanced block describes the target, got %q", landed.Adv.InfoGUID)
	}
	if landed.Adv.Level != 183 || swing.Adv.Level != 45 {
		t.Errorf("level field: player item level %d, creature level %d", landed.Adv.Level, swing.Adv.Level)
	}
}

func TestHealKeepsAmountOverhealAndTheShieldPortion(t *testing.T) {
	by, _ := decodeAll(t)
	e := one(t, by, "SPELL_HEAL", 0)
	if e.Kind != Heal {
		t.Fatalf("kind = %s", e.Kind)
	}
	if e.Amount.V != 1618 || e.Overheal.V != 806 || e.Absorbed.V != 0 {
		t.Errorf("amount=%d overheal=%d absorbed=%d", e.Amount.V, e.Overheal.V, e.Absorbed.V)
	}
	if e.Total.V != 1618 {
		t.Errorf("healedToHP = %d", e.Total.V)
	}
	if e.Effective() != 812 {
		t.Errorf("effective heal = %d, want amount minus overheal", e.Effective())
	}
	if !e.Critical.V {
		t.Error("the fixture heal is a crit")
	}
	periodic := one(t, by, "SPELL_PERIODIC_HEAL", 0)
	if periodic.Kind != Heal || periodic.Amount.V != 240 || periodic.Effective() != 0 {
		t.Errorf("periodic heal = %+v", periodic)
	}
}

func TestEnergizeReadsDecimalAmountsAndThePowerType(t *testing.T) {
	by, _ := decodeAll(t)
	e := one(t, by, "SPELL_ENERGIZE", 0)
	if e.Kind != Energize {
		t.Fatalf("kind = %s", e.Kind)
	}
	if e.Amount.V != 50 || e.OverEnergize.V != 0 || e.PowerType.V != 1 || e.MaxPower.V != 1000 {
		t.Errorf("energize = %d over %d type %d max %d", e.Amount.V, e.OverEnergize.V, e.PowerType.V, e.MaxPower.V)
	}
}

func TestMissesCarryTheAbsorbExtrasOnlyOnAbsorb(t *testing.T) {
	by, _ := decodeAll(t)
	parry := one(t, by, "SWING_MISSED", 0)
	absorb := one(t, by, "SWING_MISSED", 1)
	if parry.MissType != "PARRY" || parry.Amount.OK {
		t.Errorf("parry = %+v, want no amount", parry.MissType)
	}
	if absorb.MissType != "ABSORB" || absorb.Amount.V != 640 || absorb.BaseAmount.V != 905 {
		t.Errorf("absorb miss = %s %d %d", absorb.MissType, absorb.Amount.V, absorb.BaseAmount.V)
	}
	spell := one(t, by, "SPELL_MISSED", 0)
	if spell.MissType != "ABSORB" || spell.Amount.V != 1200 || spell.Spell.ID != 334660 {
		t.Errorf("spell miss = %+v", spell)
	}
}

func TestAuraEventsAndTheirOptionalAbsorbSize(t *testing.T) {
	by, _ := decodeAll(t)
	shield := one(t, by, "SPELL_AURA_APPLIED", 0)
	if shield.Kind != AuraApplied || shield.AuraType != "BUFF" || shield.Absorbed.V != 2210 {
		t.Errorf("shield aura = %+v", shield)
	}
	nova := one(t, by, "SPELL_AURA_APPLIED", 1)
	if nova.Kind != AuraApplied || nova.AuraType != "DEBUFF" || nova.Absorbed.OK {
		t.Errorf("nova aura = kind %s type %q absorbed %+v", nova.Kind, nova.AuraType, nova.Absorbed)
	}
	dose := one(t, by, "SPELL_AURA_APPLIED_DOSE", 0)
	if dose.Kind != AuraDose || dose.Stacks.V != 3 {
		t.Errorf("dose = %+v", dose)
	}
	removed := one(t, by, "SPELL_AURA_REMOVED", 0)
	if removed.Kind != AuraRemoved || removed.Spell.ID != 122 {
		t.Errorf("removed = %+v", removed)
	}
}

func TestCastsInterruptsAndDispels(t *testing.T) {
	by, _ := decodeAll(t)
	start := one(t, by, "SPELL_CAST_START", 0)
	success := one(t, by, "SPELL_CAST_SUCCESS", 0)
	if start.Kind != CastStart || success.Kind != CastSuccess {
		t.Fatalf("cast kinds = %s %s", start.Kind, success.Kind)
	}
	if !success.Adv.OK {
		t.Error("SPELL_CAST_SUCCESS carries the advanced block in v16")
	}
	if start.Adv.OK {
		t.Error("SPELL_CAST_START does not carry the advanced block")
	}
	failed := one(t, by, "SPELL_CAST_FAILED", 0)
	if failed.Kind != CastFailed || failed.FailedType != "Not enough mana" {
		t.Errorf("failed = %+v", failed)
	}
	interrupt := one(t, by, "SPELL_INTERRUPT", 0)
	if interrupt.Kind != Interrupt || interrupt.Spell.ID != 6552 || interrupt.ExtraSpell.ID != 334653 {
		t.Errorf("interrupt = %+v / %+v", interrupt.Spell, interrupt.ExtraSpell)
	}
	dispel := one(t, by, "SPELL_DISPEL", 0)
	if dispel.Kind != Dispel || dispel.ExtraSpell.Name != "Wrack Soul" || dispel.AuraType != "DEBUFF" {
		t.Errorf("dispel = %+v", dispel)
	}
}

func TestUnknownEventKeepsItsRawText(t *testing.T) {
	d := NewDecoder(layout.RetailV16(), fixtureBase)
	ln := lexer.Line{
		Stamp:  "9/26 20:10:00.000",
		Params: lexer.SplitParams(`SPELL_EMPOWER_END,Player-1-A,"Baelgrim",0x511,0x0,Player-1-A,"Baelgrim",0x511,0x0,1,"X",0x1,3`),
		Raw:    `9/26 20:10:00.000  SPELL_EMPOWER_END,...`,
	}
	e := d.Decode(ln)
	if e.Kind != Unknown {
		t.Fatalf("kind = %s, want unknown", e.Kind)
	}
	if e.Raw == "" {
		t.Error("an unknown event must keep its raw text")
	}
	if e.Source.Name != "Baelgrim" {
		t.Errorf("an unknown event with a common header still names its units, got %q", e.Source.Name)
	}
}

func TestMalformedLinesBecomeParseErrors(t *testing.T) {
	d := NewDecoder(layout.RetailV16(), fixtureBase)
	for _, tc := range []struct {
		name string
		line lexer.Line
	}{
		{
			name: "damage with a field missing",
			line: lexer.Line{Stamp: "9/26 20:10:00.000", Raw: "truncated",
				Params: lexer.SplitParams(`SPELL_DAMAGE,Player-1-A,"Baelgrim",0x511,0x0,Creature-0-1-2-3-4-5,"Sentinel",0xa48,0x0,116,"Frostbolt",0x10,Creature-0-1-2-3-4-5`)},
		},
		{
			name: "encounter start with the wrong width",
			line: lexer.Line{Stamp: "9/26 20:10:00.000", Raw: "bad encounter",
				Params: lexer.SplitParams(`ENCOUNTER_START,9001,"Warden Kelthas"`)},
		},
		{
			name: "combatant info with the wrong width",
			line: lexer.Line{Stamp: "9/26 20:10:00.000", Raw: "bad combatant",
				Params: lexer.SplitParams(`COMBATANT_INFO,Player-1-A,0,1`)},
		},
		{
			name: "unreadable timestamp",
			line: lexer.Line{Stamp: "not a time", Raw: "bad stamp",
				Params: lexer.SplitParams(`UNIT_DIED,0000000000000000,nil,0x0,0x0,Creature-0-1-2-3-4-5,"S",0xa48,0x0,0`)},
		},
		{
			name: "empty line",
			line: lexer.Line{Params: []string{""}, Raw: ""},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			e := d.Decode(tc.line)
			if e.Kind != ParseError {
				t.Fatalf("kind = %s, want parse_error", e.Kind)
			}
			if e.Error == "" {
				t.Error("a parse error must say what went wrong")
			}
			if e.Raw != tc.line.Raw {
				t.Errorf("raw = %q, want %q", e.Raw, tc.line.Raw)
			}
		})
	}
}

func TestTheDecoderCountsRolloversAndClockJumps(t *testing.T) {
	d := NewDecoder(layout.RetailV16(), time.Date(2026, 12, 31, 23, 59, 0, 0, time.UTC))
	mk := func(stamp string) lexer.Line {
		return lexer.Line{Stamp: stamp, Params: lexer.SplitParams(`ZONE_CHANGE,1,"Z",0`), Raw: stamp}
	}
	d.Decode(mk("1/1 00:00:01.000"))
	if d.Rollovers() != 1 {
		t.Fatalf("rollovers = %d, want 1", d.Rollovers())
	}
	d.Decode(mk("1/1 00:00:00.500"))
	if d.ClockJumps() != 1 {
		t.Fatalf("clock jumps = %d, want 1", d.ClockJumps())
	}
}
