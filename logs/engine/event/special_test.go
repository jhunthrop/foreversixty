// logs/engine/event/special_test.go
package event

import (
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/layout"
	"github.com/jhunthrop/foreversixty/logs/engine/lexer"
)

func TestBothAbsorbedShapes(t *testing.T) {
	by, _ := decodeAll(t)
	self := one(t, by, "SPELL_ABSORBED", 0)  // 19 fields, a swing was absorbed
	other := one(t, by, "SPELL_ABSORBED", 1) // 22 fields, a spell was absorbed
	if self.Kind != Absorbed || other.Kind != Absorbed {
		t.Fatalf("kinds = %s %s", self.Kind, other.Kind)
	}
	if self.Spell.ID != 0 {
		t.Errorf("the 19-field shape has no damage spell, got %d", self.Spell.ID)
	}
	if self.ExtraUnit.Name != "Sunwick-Nightslayer" || self.ExtraSpell.ID != 17 {
		t.Errorf("absorber = %q shield = %d", self.ExtraUnit.Name, self.ExtraSpell.ID)
	}
	if self.Amount.V != 640 || self.Total.V != 905 {
		t.Errorf("self shape amounts = %d of %d", self.Amount.V, self.Total.V)
	}
	if other.Spell.ID != 334660 || other.Spell.Name != "Anima Lash" {
		t.Errorf("the 22-field shape carries the damage spell, got %+v", other.Spell)
	}
	if other.ExtraUnit.Name != "Sunwick-Nightslayer" || other.ExtraSpell.ID != 17 {
		t.Errorf("absorber = %q shield = %d", other.ExtraUnit.Name, other.ExtraSpell.ID)
	}
	if other.Amount.V != 1200 || other.Total.V != 1610 {
		t.Errorf("other shape amounts = %d of %d", other.Amount.V, other.Total.V)
	}
}

func TestAbsorbedWithAnImpossibleWidthIsAParseError(t *testing.T) {
	l := layout.RetailV16()
	l.Specials["SPELL_ABSORBED"] = layout.Special{} // accept any width at the gate
	d := NewDecoder(l, fixtureBase)
	e := d.Decode(lexer.Line{
		Stamp:  "9/26 20:10:00.000",
		Raw:    "short absorbed",
		Params: lexer.SplitParams(`SPELL_ABSORBED,A,"a",0x0,0x0,B,"b",0x0,0x0,C,"c"`),
	})
	if e.Kind != ParseError {
		t.Fatalf("kind = %s, want parse_error", e.Kind)
	}
}

func TestHealAbsorbed(t *testing.T) {
	by, _ := decodeAll(t)
	e := one(t, by, "SPELL_HEAL_ABSORBED", 0)
	if e.Kind != HealAbsorbed {
		t.Fatalf("kind = %s", e.Kind)
	}
	if e.Spell.Name != "Necrotic Wound" || e.ExtraSpell.Name != "Heal" {
		t.Errorf("spells = %q and %q", e.Spell.Name, e.ExtraSpell.Name)
	}
	if e.ExtraUnit.Name != "Sunwick-Nightslayer" {
		t.Errorf("healer = %q", e.ExtraUnit.Name)
	}
	if e.Amount.V != 412 || e.Total.V != 412 {
		t.Errorf("absorbed = %d of %d", e.Amount.V, e.Total.V)
	}
}

func TestEnvironmentalDamagePutsTheTypeAfterTheAdvancedBlock(t *testing.T) {
	by, _ := decodeAll(t)
	e := one(t, by, "ENVIRONMENTAL_DAMAGE", 0)
	if e.Kind != Damage || e.EnvType != "Falling" {
		t.Fatalf("kind = %s envType = %q", e.Kind, e.EnvType)
	}
	if e.Source.GUID != "0000000000000000" || e.Dest.Name != "Thalgrit-Nightslayer" {
		t.Errorf("units = %q -> %q", e.Source.GUID, e.Dest.Name)
	}
	if e.Amount.V != 1140 || e.BaseAmount.V != 1140 {
		t.Errorf("amount = %d base = %d", e.Amount.V, e.BaseAmount.V)
	}
	if !e.Adv.OK || e.Adv.CurrentHP != 7100 {
		t.Errorf("advanced = %+v", e.Adv)
	}
}

func TestDeathsAndKills(t *testing.T) {
	by, _ := decodeAll(t)
	died := one(t, by, "UNIT_DIED", 0)
	if died.Kind != Death || died.Dest.Name != "Hollow Sentinel" {
		t.Errorf("death = %s %q", died.Kind, died.Dest.Name)
	}
	kill := one(t, by, "PARTY_KILL", 0)
	if kill.Kind != PartyKill || kill.Source.Name != "Morrowlyn-Nightslayer" {
		t.Errorf("party kill = %s %q", kill.Kind, kill.Source.Name)
	}
	instakill := one(t, by, "SPELL_INSTAKILL", 0)
	if instakill.Kind != Instakill || instakill.Dest.Name != "Ashfang" {
		t.Errorf("instakill = %s %q", instakill.Kind, instakill.Dest.Name)
	}
}

func TestEncounterStartAndEnd(t *testing.T) {
	by, _ := decodeAll(t)
	start := one(t, by, "ENCOUNTER_START", 0)
	if start.Kind != EncounterStart || start.Encounter == nil {
		t.Fatalf("start = %s %v", start.Kind, start.Encounter)
	}
	if start.Encounter.ID != 9001 || start.Encounter.Name != "Warden Kelthas" ||
		start.Encounter.Difficulty != 8 || start.Encounter.Size != 5 || start.Encounter.InstanceID != 2284 {
		t.Errorf("encounter = %+v", *start.Encounter)
	}
	end := one(t, by, "ENCOUNTER_END", 0)
	if end.Kind != EncounterEnd || end.Encounter == nil || !end.Encounter.Kill {
		t.Fatalf("end = %s %+v", end.Kind, end.Encounter)
	}
}

func TestZoneAndMapChange(t *testing.T) {
	by, _ := decodeAll(t)
	z := one(t, by, "ZONE_CHANGE", 0)
	if z.Kind != ZoneChange || z.Zone.ID != 2284 || z.Zone.Name != "Sanguine Depths" || z.Zone.Difficulty != 8 {
		t.Errorf("zone = %s %+v", z.Kind, z.Zone)
	}
	m := one(t, by, "MAP_CHANGE", 0)
	if m.Kind != MapChange || m.Zone.ID != 1675 {
		t.Fatalf("map = %s %+v", m.Kind, m.Zone)
	}
	if m.Zone.MaxX != -1300 || m.Zone.MinX != -1900 || m.Zone.MaxY != 6700 || m.Zone.MinY != 6100 {
		t.Errorf("map bounds = %+v", *m.Zone)
	}
}

func TestCombatantInfoReadsSpecTalentsGearAndAuras(t *testing.T) {
	by, _ := decodeAll(t)
	e := one(t, by, "COMBATANT_INFO", 0)
	if e.Kind != CombatantInfo || e.Combatant == nil {
		t.Fatalf("kind = %s combatant = %v", e.Kind, e.Combatant)
	}
	c := e.Combatant
	if c.GUID != "Player-4184-000000A1" || c.Faction != 0 || c.SpecID != 73 {
		t.Errorf("guid=%q faction=%d spec=%d", c.GUID, c.Faction, c.SpecID)
	}
	if c.Stats["strength"] != 1180 || c.Stats["stamina"] != 2790 || c.Stats["armor"] != 4120 {
		t.Errorf("stats = %v", c.Stats)
	}
	if len(c.Talents) != 7 || c.Talents[0] != 202751 {
		t.Errorf("talents = %v", c.Talents)
	}
	if len(c.PvPTalents) != 4 {
		t.Errorf("pvp talents = %v", c.PvPTalents)
	}
	if c.Borrowed != "[0,1,[],[],[]]" {
		t.Errorf("borrowed power is kept raw, got %q", c.Borrowed)
	}
	if len(c.Gear) != 3 {
		t.Fatalf("gear has %d slots, want 3", len(c.Gear))
	}
	if c.Gear[0].ID != 175850 || c.Gear[0].ItemLevel != 183 {
		t.Errorf("first slot = %+v", c.Gear[0])
	}
	if got := c.Gear[0].BonusIDs; len(got) != 3 || got[0] != 6788 {
		t.Errorf("bonus ids = %v", got)
	}
	if c.Gear[2].ID != 0 {
		t.Errorf("an empty slot must stay in the list as item 0, got %+v", c.Gear[2])
	}
	if c.ItemLevel != 183 {
		t.Errorf("item level = %d, want the mean of the filled slots", c.ItemLevel)
	}
	if len(c.Auras) != 2 || c.Auras[0].SpellID != 17 || c.Auras[1].SourceGUID != "Player-4184-000000A1" {
		t.Errorf("auras = %+v", c.Auras)
	}
}

func TestCombatantInfoStaysRawOnARowThatDoesNotDocumentIt(t *testing.T) {
	l := layout.ClassicWiki()
	l.Specials["COMBATANT_INFO"] = layout.Special{}
	d := NewDecoder(l, fixtureBase)
	e := d.Decode(lexer.Line{
		Stamp:  "9/26 20:10:00.000",
		Raw:    "combatant info on a row with no layout for it",
		Params: lexer.SplitParams(`COMBATANT_INFO,Player-1-A,0,1,2,3`),
	})
	if e.Kind != Unknown {
		t.Fatalf("kind = %s, want unknown rather than a guessed layout", e.Kind)
	}
	if e.Raw == "" {
		t.Error("the raw line must be kept")
	}
}

func TestEnchantAndEmote(t *testing.T) {
	by, _ := decodeAll(t)
	en := one(t, by, "ENCHANT_APPLIED", 0)
	if en.Kind != Enchant || en.Spell.Name != "Shadowcore Oil" || en.ItemID.V != 178473 ||
		en.ItemName != "Sentinel's Bulwark" {
		t.Errorf("enchant = %+v", en)
	}

	d := NewDecoder(layout.RetailV16(), fixtureBase)
	e := d.Decode(lexer.Line{
		Stamp:  "9/26 20:10:00.000",
		Raw:    "emote",
		Params: lexer.SplitParams(`EMOTE,Creature-0-1-2-3-4-5,"Hollow Sentinel",Player-4184-000000A1,"Baelgrim-Nightslayer","The Sentinel roars!"`),
	})
	if e.Kind != Emote {
		t.Fatalf("kind = %s", e.Kind)
	}
	if e.Source.Name != "Hollow Sentinel" || e.Dest.Name != "Baelgrim-Nightslayer" {
		t.Errorf("emote units = %q -> %q", e.Source.Name, e.Dest.Name)
	}
	if e.Raw != "" {
		t.Error("emote text must be dropped at parse, not kept in Raw")
	}
}

func TestChallengeModeEvents(t *testing.T) {
	d := NewDecoder(layout.RetailV16(), fixtureBase)
	start := d.Decode(lexer.Line{
		Stamp:  "9/26 20:10:00.000",
		Raw:    "cm start",
		Params: lexer.SplitParams(`CHALLENGE_MODE_START,"Sanguine Depths",2284,380,6,[9,123]`),
	})
	if start.Kind != ChallengeModeStart || start.Zone.Name != "Sanguine Depths" || start.Amount.V != 6 {
		t.Errorf("challenge start = %s %+v keystone %d", start.Kind, start.Zone, start.Amount.V)
	}
	end := d.Decode(lexer.Line{
		Stamp:  "9/26 20:40:00.000",
		Raw:    "cm end",
		Params: lexer.SplitParams(`CHALLENGE_MODE_END,2284,1,6,1800000`),
	})
	if end.Kind != ChallengeModeEnd || !end.Critical.V || end.Amount.V != 6 {
		t.Errorf("challenge end = %s success %v keystone %d", end.Kind, end.Critical.V, end.Amount.V)
	}
}
