// logs/engine/summary/summary_test.go
package summary

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
	"github.com/jhunthrop/foreversixty/logs/engine/fight"
	"github.com/jhunthrop/foreversixty/logs/engine/units"
)

var t0 = time.Date(2026, 9, 26, 20, 10, 0, 0, time.UTC)

const (
	tank   = "Player-4184-000000A1"
	healer = "Player-4184-000000A2"
	mage   = "Player-4184-000000A3"
	hunter = "Player-4184-000000A4"
	pet    = "Pet-0-2085-2284-7855-165189-01000000B1"
	boss   = "Creature-0-2085-2284-7855-169753-0000AA0001"
)

func at(sec float64) time.Time { return t0.Add(time.Duration(sec * float64(time.Second))) }

func opts(t *testing.T) (Options, *units.Registry) {
	t.Helper()
	r := units.NewRegistry(units.Options{ClassBySpec: units.RetailSpecClass})
	o := DefaultOptions()
	o.Registry = r
	o.SpecNames = units.RetailSpecName
	return o, r
}

func dmg(sec float64, src, dst string, spellID int64, spellName string, amount, overkill int64) event.Event {
	name := "SPELL_DAMAGE"
	if spellID == 0 {
		name = "SWING_DAMAGE"
	}
	return event.Event{
		Time: at(sec), Kind: event.Damage, Name: name,
		Source: event.Unit{GUID: src, Flags: 0x512}, Dest: event.Unit{GUID: dst, Flags: 0xa48},
		Spell:    event.Spell{ID: spellID, Name: spellName, School: 0x10},
		Amount:   event.OptInt{V: amount, OK: true},
		Overkill: event.OptInt{V: overkill, OK: true},
	}
}

func heal(sec float64, src, dst string, spellID int64, spellName string, amount, overheal int64) event.Event {
	return event.Event{
		Time: at(sec), Kind: event.Heal, Name: "SPELL_HEAL",
		Source: event.Unit{GUID: src, Flags: 0x512}, Dest: event.Unit{GUID: dst, Flags: 0x512},
		Spell:    event.Spell{ID: spellID, Name: spellName, School: 0x2},
		Amount:   event.OptInt{V: amount, OK: true},
		Overheal: event.OptInt{V: overheal, OK: true},
	}
}

// script is the fixture fight every table test is checked against. The
// expected numbers below are computed by hand from these lines.
func script() []event.Event {
	return []event.Event{
		{Time: at(0), Kind: event.Summon, Name: "SPELL_SUMMON",
			Source: event.Unit{GUID: hunter, Name: "Thalgrit-Nightslayer", Flags: 0x512},
			Dest:   event.Unit{GUID: pet, Name: "Ashfang", Flags: 0x1114}},
		{Time: at(0), Kind: event.CombatantInfo, Name: "COMBATANT_INFO",
			Combatant: &event.Combatant{GUID: tank, SpecID: 73, ItemLevel: 183,
				Talents: []int64{202751},
				Auras:   []event.Aura{{SourceGUID: healer, SpellID: 17}, {SourceGUID: tank, SpellID: 871}}}},
		dmg(1, mage, boss, 116, "Frostbolt", 1000, -1),
		dmg(2, mage, boss, 116, "Frostbolt", 500, -1),
		dmg(2, pet, boss, 0, "", 200, -1),
		dmg(3, boss, tank, 334660, "Anima Lash", 800, -1),
		heal(4, healer, tank, 2060, "Heal", 1000, 400),
		{Time: at(5), Kind: event.CastStart, Name: "SPELL_CAST_START",
			Source: event.Unit{GUID: mage, Flags: 0x512}, Spell: event.Spell{ID: 116, Name: "Frostbolt"}},
		{Time: at(6), Kind: event.CastSuccess, Name: "SPELL_CAST_SUCCESS",
			Source: event.Unit{GUID: mage, Flags: 0x512}, Spell: event.Spell{ID: 116, Name: "Frostbolt"}},
		{Time: at(6), Kind: event.CastFailed, Name: "SPELL_CAST_FAILED",
			Source: event.Unit{GUID: mage, Flags: 0x512}, Spell: event.Spell{ID: 116, Name: "Frostbolt"},
			FailedType: "Not enough mana"},
		{Time: at(7), Kind: event.AuraApplied, Name: "SPELL_AURA_APPLIED",
			Source: event.Unit{GUID: mage, Flags: 0x512}, Dest: event.Unit{GUID: boss, Flags: 0xa48},
			Spell: event.Spell{ID: 122, Name: "Frost Nova"}, AuraType: "DEBUFF"},
		{Time: at(11), Kind: event.AuraRemoved, Name: "SPELL_AURA_REMOVED",
			Source: event.Unit{GUID: mage, Flags: 0x512}, Dest: event.Unit{GUID: boss, Flags: 0xa48},
			Spell: event.Spell{ID: 122, Name: "Frost Nova"}, AuraType: "DEBUFF"},
		{Time: at(12), Kind: event.Interrupt, Name: "SPELL_INTERRUPT",
			Source: event.Unit{GUID: tank, Flags: 0x512}, Dest: event.Unit{GUID: boss, Flags: 0xa48},
			Spell: event.Spell{ID: 6552, Name: "Pummel"}, ExtraSpell: event.Spell{ID: 334653, Name: "Anima Surge"}},
		{Time: at(13), Kind: event.Dispel, Name: "SPELL_DISPEL",
			Source: event.Unit{GUID: healer, Flags: 0x512}, Dest: event.Unit{GUID: mage, Flags: 0x512},
			Spell: event.Spell{ID: 527, Name: "Purify"}, ExtraSpell: event.Spell{ID: 321038, Name: "Wrack Soul"},
			AuraType: "DEBUFF"},
		{Time: at(14), Kind: event.Missed, Name: "SWING_MISSED",
			Source: event.Unit{GUID: boss, Flags: 0xa48}, Dest: event.Unit{GUID: tank, Flags: 0x512},
			MissType: "PARRY"},
		{Time: at(15), Kind: event.Absorbed, Name: "SPELL_ABSORBED",
			Source: event.Unit{GUID: boss, Flags: 0xa48}, Dest: event.Unit{GUID: tank, Flags: 0x512},
			ExtraUnit:  event.Unit{GUID: healer, Flags: 0x512},
			ExtraSpell: event.Spell{ID: 17, Name: "Power Word: Shield"},
			Amount:     event.OptInt{V: 300, OK: true}, Total: event.OptInt{V: 450, OK: true}},
		dmg(15.5, boss, tank, 0, "", 900, -1),
		// The same swing, written again at impact with the tank's health on it.
		// It must not become a second hit anywhere; it only fills in the health.
		{Time: at(15.5), Kind: event.DamageLanded, Name: "SWING_DAMAGE_LANDED",
			Source: event.Unit{GUID: boss, Flags: 0xa48}, Dest: event.Unit{GUID: tank, Flags: 0x512},
			Amount: event.OptInt{V: 900, OK: true}, Overkill: event.OptInt{V: -1, OK: true},
			Adv: event.Advanced{OK: true, InfoGUID: tank, CurrentHP: 5000, MaxHP: 20000}},
		dmg(16, boss, tank, 334660, "Anima Lash", 5000, 1200),
		{Time: at(16), Kind: event.Death, Name: "UNIT_DIED",
			Dest: event.Unit{GUID: tank, Flags: 0x512}},
		{Time: at(19), Kind: event.Energize, Name: "SPELL_ENERGIZE",
			Source: event.Unit{GUID: mage, Flags: 0x512}, Dest: event.Unit{GUID: mage, Flags: 0x512},
			Spell:     event.Spell{ID: 34428, Name: "Victory Rush"},
			Amount:    event.OptInt{V: 50, OK: true},
			PowerType: event.OptInt{V: 0, OK: true}, MaxPower: event.OptInt{V: 1000, OK: true},
			Adv: event.Advanced{OK: true, InfoGUID: mage, PowerType: 0, CurrentPower: 900, MaxPower: 1000}},
		dmg(20, mage, boss, 116, "Frostbolt", 2000, 700),
		{Time: at(20), Kind: event.Death, Name: "UNIT_DIED",
			Dest: event.Unit{GUID: boss, Flags: 0xa48}},
	}
}

func build(t *testing.T) (*Accumulator, fight.Fight, Summary) {
	t.Helper()
	o, reg := opts(t)
	a := New(o)
	a.Start(at(0))
	f := fight.Fight{
		Index: 1, Kind: fight.Encounter, EncounterID: 9001, Name: "Warden Kelthas",
		Difficulty: 8, Size: 5, Kill: true, Start: at(0), End: at(20),
		Players: []string{tank, healer, mage, hunter},
	}
	for _, e := range script() {
		reg.Observe(e)
		a.Add(e)
	}
	return a, f, a.Snapshot(f, "test")
}

// jsonOf renders a value as JSON, so two renderings can be compared as
// bytes.
func jsonOf(t *testing.T, v any) string {
	t.Helper()
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

// jsonIndentOf renders a value as indented JSON, for the committed golden:
// a text file a human can diff is worth the extra bytes.
func jsonIndentOf(t *testing.T, v any) string {
	t.Helper()
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

func actorByGUID(rows []Actor, guid string) (Actor, bool) {
	for _, r := range rows {
		if r.GUID == guid {
			return r, true
		}
	}
	return Actor{}, false
}

func TestDamageDoneCreditsPetsToTheirOwner(t *testing.T) {
	_, _, s := build(t)
	m, ok := actorByGUID(s.DamageDone, mage)
	if !ok {
		t.Fatal("no mage row")
	}
	// 1000 + 500 + 2000
	if m.Effective != 3500 {
		t.Errorf("mage damage = %d, want 3500", m.Effective)
	}
	if len(m.Abilities) != 1 || m.Abilities[0].SpellID != 116 {
		t.Fatalf("mage abilities = %+v", m.Abilities)
	}
	if m.Abilities[0].Overkill != 700 {
		t.Errorf("overkill = %d, want 700 (the -1s clamp to zero)", m.Abilities[0].Overkill)
	}
	if m.Abilities[0].Min != 500 || m.Abilities[0].Max != 2000 {
		t.Errorf("min/max = %d/%d", m.Abilities[0].Min, m.Abilities[0].Max)
	}
	h, ok := actorByGUID(s.DamageDone, hunter)
	if !ok {
		t.Fatal("the pet's damage must appear on the hunter's row")
	}
	if h.Effective != 200 {
		t.Errorf("hunter damage = %d, want the pet's 200", h.Effective)
	}
	if _, isOwnRow := actorByGUID(s.DamageDone, pet); isOwnRow {
		t.Error("the pet must not have its own row")
	}
}

// TestAbilityMinStaysAtATrueZero guards against Ability.Min conflating "no
// hit recorded yet" with "the lowest hit so far was a genuine zero" (a
// fully resisted or otherwise informational hit lands for 0). Once a real
// zero-amount hit is recorded, later larger hits must not overwrite it.
func TestAbilityMinStaysAtATrueZero(t *testing.T) {
	o, reg := opts(t)
	a := New(o)
	a.Start(at(0))
	events := []event.Event{
		dmg(1, mage, boss, 116, "Frostbolt", 0, -1),
		dmg(2, mage, boss, 116, "Frostbolt", 500, -1),
		dmg(3, mage, boss, 116, "Frostbolt", 900, -1),
	}
	for _, e := range events {
		reg.Observe(e)
		a.Add(e)
	}
	s := a.Snapshot(fight.Fight{Index: 1, Start: at(0), End: at(3)}, "test")
	m, ok := actorByGUID(s.DamageDone, mage)
	if !ok {
		t.Fatal("no mage row")
	}
	if len(m.Abilities) != 1 || m.Abilities[0].Min != 0 {
		t.Fatalf("min = %+v, want 0 (a genuine zero-amount hit, not overwritten by later larger hits)", m.Abilities)
	}
	if m.Abilities[0].Max != 900 {
		t.Errorf("max = %d, want 900", m.Abilities[0].Max)
	}
}

func TestPerSecondSeries(t *testing.T) {
	_, _, s := build(t)
	m, _ := actorByGUID(s.DamageDone, mage)
	// Buckets: second 1 = 1000, second 2 = 500, second 20 = 2000.
	if len(m.Series) != 21 {
		t.Fatalf("series has %d buckets, want 21", len(m.Series))
	}
	if m.Series[1] != 1000 || m.Series[2] != 500 || m.Series[20] != 2000 {
		t.Errorf("series = %v", m.Series)
	}
	if m.Series[0] != 0 || m.Series[3] != 0 {
		t.Errorf("quiet seconds must be zero, got %v", m.Series[:4])
	}
}

func TestDamageTakenAndTheAbsorbCredit(t *testing.T) {
	_, _, s := build(t)
	tk, ok := actorByGUID(s.DamageTaken, tank)
	if !ok {
		t.Fatal("no tank row in damage taken")
	}
	if tk.Effective != 6700 { // 800 + 900 (the swing, once, not its landed line too) + 5000
		t.Errorf("tank damage taken = %d, want 6700", tk.Effective)
	}
	if tk.Abilities[0].SpellID != 334660 || tk.Abilities[0].Effective != 5800 {
		t.Errorf("abilities are sorted by effective damage, got %+v", tk.Abilities[0])
	}
	var melee *Ability
	for i := range tk.Abilities {
		if tk.Abilities[i].SpellID == 0 {
			melee = &tk.Abilities[i]
		}
	}
	if melee == nil || melee.Misses["PARRY"] != 1 {
		t.Errorf("the parried swing must appear as a melee miss, got %+v", melee)
	}
	hl, ok := actorByGUID(s.Healing, healer)
	if !ok {
		t.Fatal("no healer row")
	}
	// 1000 heal with 400 overheal is 600 effective, plus a 300 absorb.
	if hl.Effective != 900 {
		t.Errorf("healer effective = %d, want 900", hl.Effective)
	}
	if hl.Overheal != 400 {
		t.Errorf("overheal = %d, want 400", hl.Overheal)
	}
	if hl.Absorbed != 300 {
		t.Errorf("absorbed = %d, want 300", hl.Absorbed)
	}
}

func TestDeathsKeepTheKillingBlowAndTheAurasHeld(t *testing.T) {
	_, _, s := build(t)
	if len(s.Deaths) != 1 {
		t.Fatalf("deaths = %d, want 1 (only players count)", len(s.Deaths))
	}
	d := s.Deaths[0]
	if d.GUID != tank || d.AtMS != 16000 {
		t.Errorf("death = %+v", d)
	}
	if d.KillingBlow == nil || d.KillingBlow.Amount != 5000 || d.KillingBlow.Overkill != 1200 {
		t.Fatalf("killing blow = %+v", d.KillingBlow)
	}
	if len(d.Last) != 3 {
		t.Fatalf("last damage = %d events, want the three hits on the tank (the landed line is not a fourth)", len(d.Last))
	}
	if d.Last[1].HPAfter != 5000 || d.Last[1].MaxHP != 20000 {
		t.Errorf("the swing should carry the health its landed line reported, got %+v", d.Last[1])
	}
	if d.Last[0].SpellName != "Anima Lash" {
		t.Errorf("first recorded hit = %+v", d.Last[0])
	}
	if len(d.Heals) != 1 || d.Heals[0].SpellName != "Heal" || d.Heals[0].Amount != 1000 || d.Heals[0].Overheal != 400 {
		t.Errorf("heals before the death = %+v, want the healer's one Heal on the tank", d.Heals)
	}
	if d.KillingBlow.SpellName != "Anima Lash" || d.KillingBlow.Overkill != 1200 {
		t.Errorf("killing blow = %+v, want the hit that overkilled", d.KillingBlow)
	}
}

func TestAurasUpAtThePullAreCountedFromItsStart(t *testing.T) {
	_, _, s := build(t)
	// The tank's combatant snapshot lists spell 17 (from the healer) and 871 (their own);
	// neither has an APPLIED line, so without seeding they would not exist.
	var seeded *AuraTrack
	for i := range s.Auras {
		if s.Auras[i].TargetGUID == tank && s.Auras[i].SpellID == 871 {
			seeded = &s.Auras[i]
		}
	}
	if seeded == nil {
		t.Fatal("the aura the snapshot said was up at the pull is missing")
	}
	if seeded.Name != "Spell #871" || len(seeded.Segments) != 1 || seeded.Segments[0].StartMS != 0 {
		t.Errorf("seeded track = %+v, want one segment from the pull's start named by its id", *seeded)
	}
	if seeded.Applications != 0 {
		t.Errorf("a seeded aura was not applied during the fight: applications = %d", seeded.Applications)
	}
}

func TestAuraUptimeAndSegments(t *testing.T) {
	_, _, s := build(t)
	var nova *AuraTrack
	for i := range s.Auras {
		if s.Auras[i].SpellID == 122 {
			nova = &s.Auras[i]
		}
	}
	if nova == nil {
		t.Fatal("Frost Nova track missing")
	}
	if nova.UptimeMS != 4000 {
		t.Errorf("uptime = %d ms, want 4000", nova.UptimeMS)
	}
	if len(nova.Segments) != 1 || nova.Segments[0].StartMS != 7000 || nova.Segments[0].EndMS != 11000 {
		t.Errorf("segments = %+v", nova.Segments)
	}
	if nova.Applications != 1 || len(nova.Appliers) != 1 || nova.Appliers[0] != mage {
		t.Errorf("applications = %d appliers = %v", nova.Applications, nova.Appliers)
	}
}

func TestAnAuraStillUpAtTheEndIsClosedAtTheFightEnd(t *testing.T) {
	o, reg := opts(t)
	a := New(o)
	a.Start(at(0))
	e := event.Event{Time: at(2), Kind: event.AuraApplied, Name: "SPELL_AURA_APPLIED",
		Source: event.Unit{GUID: healer, Flags: 0x512}, Dest: event.Unit{GUID: tank, Flags: 0x512},
		Spell: event.Spell{ID: 17, Name: "Power Word: Shield"}, AuraType: "BUFF"}
	reg.Observe(e)
	a.Add(e)
	a.Add(dmg(10, boss, tank, 1, "x", 1, -1))
	s := a.Snapshot(fight.Fight{Index: 1, Start: at(0), End: at(10)}, "test")
	if len(s.Auras) != 1 || s.Auras[0].UptimeMS != 8000 {
		t.Fatalf("auras = %+v", s.Auras)
	}
}

func TestCastsInterruptsDispels(t *testing.T) {
	_, _, s := build(t)
	var frostbolt *CastRow
	for i := range s.Casts {
		if s.Casts[i].GUID == mage && s.Casts[i].SpellID == 116 {
			frostbolt = &s.Casts[i]
		}
	}
	if frostbolt == nil {
		t.Fatal("no frostbolt cast row")
	}
	if frostbolt.Started != 1 || frostbolt.Succeeded != 1 || frostbolt.Failed != 1 {
		t.Errorf("casts = %+v", *frostbolt)
	}
	if frostbolt.CastTimeMS != 1000 {
		t.Errorf("cast time = %d ms, want 1000 from the start-success pair", frostbolt.CastTimeMS)
	}
	if frostbolt.FailReasons["Not enough mana"] != 1 {
		t.Errorf("fail reasons = %v", frostbolt.FailReasons)
	}
	if len(frostbolt.Sequence) != 1 || frostbolt.Sequence[0] != 6000 {
		t.Errorf("sequence = %v", frostbolt.Sequence)
	}
	if len(s.Interrupts) != 1 || s.Interrupts[0].ExtraSpellName != "Anima Surge" {
		t.Errorf("interrupts = %+v", s.Interrupts)
	}
	if len(s.Dispels) != 1 || s.Dispels[0].ExtraSpellName != "Wrack Soul" {
		t.Errorf("dispels = %+v", s.Dispels)
	}
}

// TestExchangeRowsBreakTiesOnTheFullMapKey guards against exchangeRows
// falling back on map iteration order — which Go does not guarantee is
// stable across iterations — when two rows tie on Count, SourceGUID and
// ExtraSpellID. The comparator must also break ties on TargetGUID and
// SpellID, the rest of exchangeKey, so the order is a function of the data
// and not of how the map happened to be walked.
func TestExchangeRowsBreakTiesOnTheFullMapKey(t *testing.T) {
	o, reg := opts(t)
	a := New(o)
	a.Start(at(0))
	// Same source, same extra spell removed, count 1 each: these tie on
	// the comparator's first three fields and differ only by target.
	events := []event.Event{
		{Time: at(1), Kind: event.Dispel, Name: "SPELL_DISPEL",
			Source: event.Unit{GUID: healer, Flags: 0x512}, Dest: event.Unit{GUID: hunter, Flags: 0x512},
			Spell: event.Spell{ID: 527, Name: "Purify"}, ExtraSpell: event.Spell{ID: 999, Name: "Curse"},
			AuraType: "DEBUFF"},
		{Time: at(2), Kind: event.Dispel, Name: "SPELL_DISPEL",
			Source: event.Unit{GUID: healer, Flags: 0x512}, Dest: event.Unit{GUID: mage, Flags: 0x512},
			Spell: event.Spell{ID: 527, Name: "Purify"}, ExtraSpell: event.Spell{ID: 999, Name: "Curse"},
			AuraType: "DEBUFF"},
		// Same source, same target, same extra spell, count 1 each: these
		// tie on the first three fields AND on TargetGUID, and differ only
		// by the primary spell used to cast the dispel.
		{Time: at(3), Kind: event.Dispel, Name: "SPELL_DISPEL",
			Source: event.Unit{GUID: healer, Flags: 0x512}, Dest: event.Unit{GUID: tank, Flags: 0x512},
			Spell: event.Spell{ID: 528, Name: "Mass Dispel"}, ExtraSpell: event.Spell{ID: 999, Name: "Curse"},
			AuraType: "DEBUFF"},
		{Time: at(4), Kind: event.Dispel, Name: "SPELL_DISPEL",
			Source: event.Unit{GUID: healer, Flags: 0x512}, Dest: event.Unit{GUID: tank, Flags: 0x512},
			Spell: event.Spell{ID: 527, Name: "Purify"}, ExtraSpell: event.Spell{ID: 999, Name: "Curse"},
			AuraType: "DEBUFF"},
	}
	for _, e := range events {
		reg.Observe(e)
		a.Add(e)
	}
	f := fight.Fight{Index: 1, Start: at(0), End: at(4)}
	// GUIDs sort tank < mage < hunter (the fixture GUIDs end ...A1, ...A3,
	// ...A4), so with Count, SourceGUID and ExtraSpellID tied throughout,
	// the four distinct (target, spellID) rows must land in this order:
	// tank/527, tank/528, mage/527, hunter/527.
	want := []struct {
		target  string
		spellID int64
	}{
		{tank, 527}, {tank, 528}, {mage, 527}, {hunter, 527},
	}
	for trial := range 20 {
		s := a.Snapshot(f, "test")
		if len(s.Dispels) != len(want) {
			t.Fatalf("trial %d: dispels = %d, want %d", trial, len(s.Dispels), len(want))
		}
		for i, w := range want {
			if s.Dispels[i].TargetGUID != w.target || s.Dispels[i].SpellID != w.spellID {
				t.Fatalf("trial %d: dispels not ordered by the full key, got %+v", trial, s.Dispels)
			}
		}
	}
}

func TestResourcesFromTheAdvancedBlock(t *testing.T) {
	_, _, s := build(t)
	// The mage's energize and advanced blocks, plus the tank's own from the swing that
	// landed on them: a landed line's block describes the target, power included.
	if len(s.Resources) != 2 {
		t.Fatalf("resources = %+v", s.Resources)
	}
	r := s.Resources[1]
	if r.GUID != mage || r.PowerType != 0 || r.Gained != 50 {
		t.Errorf("resource = %+v", r)
	}
	if len(r.Series) != 20 || r.Series[19] != 900 {
		t.Errorf("series len=%d tail=%v", len(r.Series), r.Series[len(r.Series)-1:])
	}
}

func TestThreatReportsItsModelAndSaysWhenItIsIncomplete(t *testing.T) {
	_, _, s := build(t)
	if len(s.Threat) == 0 {
		t.Fatal("no threat rows")
	}
	for _, r := range s.Threat {
		if r.ModelVersion != "base-1" {
			t.Errorf("model = %q", r.ModelVersion)
		}
		if r.Complete {
			t.Error("with no modifier table the threat model must report itself incomplete")
		}
	}
	// The mage did 3500 effective damage, so 3500 threat under the base model.
	for _, r := range s.Threat {
		if r.GUID == mage && r.Threat != 3500 {
			t.Errorf("mage threat = %f, want 3500", r.Threat)
		}
	}
	withTable := BaseThreat{Modifiers: map[int64]float64{116: 0.5}}
	if !withTable.Complete() {
		t.Error("a model with modifiers is complete")
	}
	if got := withTable.Damage(dmg(1, mage, boss, 116, "Frostbolt", 100, -1)); got != 50 {
		t.Errorf("modified threat = %f, want 50", got)
	}
}

func TestRosterAndRankingMetrics(t *testing.T) {
	a, f, s := build(t)
	if len(s.Roster) != 4 {
		t.Fatalf("roster = %d rows, want 4", len(s.Roster))
	}
	byGUID := map[string]RosterRow{}
	for _, r := range s.Roster {
		byGUID[r.GUID] = r
	}
	if byGUID[mage].Role != "dps" || byGUID[healer].Role != "healer" || byGUID[tank].Role != "tank" {
		t.Errorf("roles = mage %q healer %q tank %q",
			byGUID[mage].Role, byGUID[healer].Role, byGUID[tank].Role)
	}
	if byGUID[tank].Deaths != 1 {
		t.Errorf("tank deaths = %d", byGUID[tank].Deaths)
	}
	if byGUID[tank].Class != "Warrior" || byGUID[tank].Spec != "Protection" {
		t.Errorf("tank class/spec = %q %q", byGUID[tank].Class, byGUID[tank].Spec)
	}
	if got := byGUID[mage].DPS; got != 175 { // 3500 over 20 seconds
		t.Errorf("mage dps = %f, want 175", got)
	}

	rows := a.Metrics("report-1", f, s, "test")
	if len(rows) != 4 {
		t.Fatalf("metrics = %d rows", len(rows))
	}
	for _, r := range rows {
		if r.ReportID != "report-1" || r.EncounterID != 9001 || !r.Kill ||
			r.Difficulty != 8 || r.Size != 5 || r.DurationMS != 20000 || r.EngineVersion != "test" {
			t.Fatalf("metric row = %+v", r)
		}
		if r.PlayerGUID == mage && (r.Metric != "dps" || r.Value != 175) {
			t.Errorf("mage metric = %s %f", r.Metric, r.Value)
		}
		if r.PlayerGUID == healer && r.Metric != "hps" {
			t.Errorf("healer metric = %s", r.Metric)
		}
	}

	trash := f
	trash.Kind = fight.Trash
	if got := a.Metrics("report-1", trash, s, "test"); got != nil {
		t.Errorf("trash must not produce ranking rows, got %d", len(got))
	}
}

func TestCombatantRowsCarryGearTalentsAndConsumables(t *testing.T) {
	o, reg := opts(t)
	o.ConsumableSpells = map[int64]string{871: "Shield Wall Elixir"}
	o.RaidBuffSpells = map[int64]string{17: "Power Word: Shield", 1459: "Arcane Intellect"}
	a := New(o)
	a.Start(at(0))
	for _, e := range script() {
		reg.Observe(e)
		a.Add(e)
	}
	s := a.Snapshot(fight.Fight{Index: 1, Start: at(0), End: at(20), Players: []string{tank}}, "test")
	if len(s.Combatants) != 1 {
		t.Fatalf("combatants = %+v", s.Combatants)
	}
	c := s.Combatants[0]
	if c.GUID != tank || c.SpecID != 73 || c.Spec != "Protection" || c.ItemLevel != 183 {
		t.Errorf("combatant = %+v", c)
	}
	if len(c.Talents) != 1 || c.Talents[0] != 202751 {
		t.Errorf("talents = %v", c.Talents)
	}
	if len(c.Consumables) != 1 || c.Consumables[0].Name != "Shield Wall Elixir" {
		t.Errorf("consumables = %+v", c.Consumables)
	}
	if len(c.RaidBuffs) != 1 || c.RaidBuffs[0].SpellID != 17 {
		t.Errorf("raid buffs = %+v", c.RaidBuffs)
	}
	if len(c.MissingBuffs) != 1 || c.MissingBuffs[0] != 1459 {
		t.Errorf("missing buffs = %v", c.MissingBuffs)
	}
}

// TestNilRegistryFallsBackToGUIDs pins the contract that Options.Registry
// may be nil (Task 9's startFight sets it after New, not before): names
// fall back to the raw GUID and pet ownership cannot be resolved, but
// nothing panics.
func TestNilRegistryFallsBackToGUIDs(t *testing.T) {
	o := DefaultOptions()
	// o.Registry is left nil on purpose.
	a := New(o)
	a.Start(at(0))
	events := []event.Event{
		dmg(1, mage, boss, 116, "Frostbolt", 1000, -1),
		dmg(2, pet, boss, 0, "", 200, -1),
	}
	for _, e := range events {
		a.Add(e)
	}
	s := a.Snapshot(fight.Fight{Index: 1, Start: at(0), End: at(2), Players: []string{mage}}, "test")
	m, ok := actorByGUID(s.DamageDone, mage)
	if !ok {
		t.Fatal("no mage row")
	}
	if m.Name != mage {
		t.Errorf("name = %q, want the raw GUID %q with no registry", m.Name, mage)
	}
	// With no registry, owner() cannot resolve the pet to its owner, so the
	// pet's damage lands on its own GUID rather than being credited to the
	// hunter.
	p, ok := actorByGUID(s.DamageDone, pet)
	if !ok {
		t.Fatal("with no registry the pet's damage must land on the pet's own GUID")
	}
	if p.Name != pet {
		t.Errorf("pet name = %q, want the raw GUID", p.Name)
	}
	if len(s.Roster) != 1 || s.Roster[0].Name != mage {
		t.Errorf("roster = %+v, want the mage's row named by its GUID", s.Roster)
	}
}

func TestSnapshotIsDeterministic(t *testing.T) {
	first, _, _ := build(t)
	_, f, _ := build(t)
	a, b := first.Snapshot(f, "test"), first.Snapshot(f, "test")
	ja, err := json.Marshal(a)
	if err != nil {
		t.Fatal(err)
	}
	jb, err := json.Marshal(b)
	if err != nil {
		t.Fatal(err)
	}
	if string(ja) != string(jb) {
		t.Fatal("two snapshots of the same accumulator differ")
	}
}

func TestRecomputingFromTheEventsMatchesTheStreamedSummary(t *testing.T) {
	// The property the spec asks for: rebuilding a fresh Accumulator from
	// the same script of events, 20 times over, reproduces the exact same
	// summary bytes every time as the one built once in build(t). This is
	// deliberately the same event order each trial: order is semantically
	// load-bearing for the per-second series, cast start/success pairing,
	// and the death windows, so a shuffled run would not be a stronger
	// version of this property, only a different and failing one.
	streamed, f, want := build(t)
	_ = streamed

	for trial := range 20 {
		o, reg := opts(t)
		a := New(o)
		a.Start(at(0))
		// Registry order does not affect the summary, so observe first.
		evs := script()
		for _, e := range evs {
			reg.Observe(e)
		}
		for _, e := range evs {
			a.Add(e)
		}
		got := a.Snapshot(f, "test")
		jw, err := json.Marshal(want)
		if err != nil {
			t.Fatal(err)
		}
		jg, err := json.Marshal(got)
		if err != nil {
			t.Fatal(err)
		}
		if string(jw) != string(jg) {
			t.Fatalf("trial %d: recomputed summary differs from the streamed one", trial)
		}
	}
}

// TestSnapshotDoesNotAliasLiveAccumulatorState covers the live-tail use
// Snapshot's doc comment advertises: a snapshot already handed to a caller
// must not change when later events land in the accumulator.
func TestSnapshotDoesNotAliasLiveAccumulatorState(t *testing.T) {
	o, reg := opts(t)
	a := New(o)
	a.Start(at(0))
	f := fight.Fight{Index: 1, Kind: fight.Encounter, Start: at(0), End: at(30),
		Players: []string{tank, healer, mage}}
	for _, e := range script() {
		reg.Observe(e)
		a.Add(e)
	}

	early := a.Snapshot(f, "test")
	before := jsonOf(t, early)

	// Everything the accumulator mutates in place: a per-second bucket an
	// earlier event already wrote, a miss type, a cast failure reason, a
	// cast sequence, and a resource series bucket.
	later := []event.Event{
		dmg(1, mage, boss, 116, "Frostbolt", 9999, -1),
		{Time: at(21), Kind: event.Missed, Name: "SWING_MISSED",
			Source: event.Unit{GUID: boss, Flags: 0xa48}, Dest: event.Unit{GUID: tank, Flags: 0x512},
			MissType: "DODGE"},
		{Time: at(22), Kind: event.CastFailed, Name: "SPELL_CAST_FAILED",
			Source: event.Unit{GUID: mage, Flags: 0x512}, Spell: event.Spell{ID: 116, Name: "Frostbolt"},
			FailedType: "Out of range"},
		{Time: at(23), Kind: event.CastSuccess, Name: "SPELL_CAST_SUCCESS",
			Source: event.Unit{GUID: mage, Flags: 0x512}, Spell: event.Spell{ID: 116, Name: "Frostbolt"}},
		{Time: at(19), Kind: event.Energize, Name: "SPELL_ENERGIZE",
			Source: event.Unit{GUID: mage, Flags: 0x512}, Dest: event.Unit{GUID: mage, Flags: 0x512},
			Spell:     event.Spell{ID: 34428, Name: "Victory Rush"},
			Amount:    event.OptInt{V: 50, OK: true},
			PowerType: event.OptInt{V: 0, OK: true}, MaxPower: event.OptInt{V: 1000, OK: true},
			Adv: event.Advanced{OK: true, InfoGUID: mage, PowerType: 0, CurrentPower: 100, MaxPower: 1000}},
	}
	for _, e := range later {
		reg.Observe(e)
		a.Add(e)
	}

	if after := jsonOf(t, early); after != before {
		t.Fatalf("the snapshot changed after later events landed:\nbefore %s\nafter  %s", before, after)
	}

	// And the new snapshot really does see the new events, so the test is
	// not passing because nothing was folded in.
	late := a.Snapshot(f, "test")
	if jsonOf(t, late) == before {
		t.Fatal("the later events changed nothing at all; the test proves nothing")
	}
}

// TestSnapshotIsSafeToSerialiseWhileTheParseContinues is the race-detector
// half of the same guarantee: the companion renders a snapshot off the
// parse goroutine every few seconds.
func TestSnapshotIsSafeToSerialiseWhileTheParseContinues(t *testing.T) {
	o, reg := opts(t)
	a := New(o)
	a.Start(at(0))
	f := fight.Fight{Index: 1, Kind: fight.Encounter, Start: at(0), End: at(30),
		Players: []string{tank, healer, mage}}
	for _, e := range script() {
		reg.Observe(e)
		a.Add(e)
	}
	snap := a.Snapshot(f, "test")

	done := make(chan string, 1)
	go func() { done <- jsonOf(t, snap) }()
	for i := 0; i < 200; i++ {
		a.Add(dmg(float64(i%25), mage, boss, 116, "Frostbolt", int64(i), -1))
	}
	<-done
}

// TestAnEnergizeWithNoDestinationOpensNoResourceRow covers the guard that
// used to test Source while keying on Dest.
func TestAnEnergizeWithNoDestinationOpensNoResourceRow(t *testing.T) {
	o, _ := opts(t)
	a := New(o)
	a.Start(at(0))
	for _, dest := range []string{"", units.NoGUID} {
		a.Add(event.Event{
			Time: at(1), Kind: event.Energize, Name: "SPELL_ENERGIZE",
			Source:    event.Unit{GUID: mage, Flags: 0x512},
			Dest:      event.Unit{GUID: dest},
			Amount:    event.OptInt{V: 50, OK: true},
			PowerType: event.OptInt{V: 0, OK: true},
		})
	}
	rows := a.Snapshot(fight.Fight{Index: 1, Start: at(0), End: at(2)}, "test").Resources
	for _, r := range rows {
		if r.GUID == "" || r.GUID == units.NoGUID {
			t.Errorf("a resource row was opened for GUID %q", r.GUID)
		}
	}
	if len(rows) != 0 {
		t.Errorf("got %d resource rows, want none", len(rows))
	}
}

// TestActivityPctIsClampedAtOneHundred covers a fight shorter than the
// activity gap: markActive credits a full ActiveGap for an actor's first
// action, so the ratio can exceed one and the table would show a
// percentage the engine invented.
func TestActivityPctIsClampedAtOneHundred(t *testing.T) {
	o, reg := opts(t)
	a := New(o)
	a.Start(at(0))
	e := dmg(0.2, mage, boss, 116, "Frostbolt", 100, -1)
	reg.Observe(e)
	a.Add(e)
	f := fight.Fight{Index: 1, Kind: fight.Encounter, Start: at(0), End: at(0.2),
		Players: []string{mage}}
	s := a.Snapshot(f, "test")
	if len(s.Roster) != 1 {
		t.Fatalf("roster = %+v", s.Roster)
	}
	row := s.Roster[0]
	if row.ActivityPct > 100 {
		t.Errorf("activity = %v%%, want it clamped at 100", row.ActivityPct)
	}
	if row.ActivityPct != 100 {
		t.Errorf("activity = %v%%, want 100 for an actor active for the whole fight", row.ActivityPct)
	}
	// The underlying accounting is untouched: it is what the ranking
	// metric divides by, so clamping it would change a ranked number.
	if row.ActiveMS != o.ActiveGap.Milliseconds() {
		t.Errorf("active_ms = %d, want the full gap %d", row.ActiveMS, o.ActiveGap.Milliseconds())
	}
}
