// logs/engine/event/special.go
package event

import (
	"fmt"
	"strings"

	"github.com/jhunthrop/foreversixty/logs/engine/layout"
	"github.com/jhunthrop/foreversixty/logs/engine/lexer"
)

// specialNeeds is how many fields each branch of decodeSpecial reads
// unconditionally, keyed by event name. The layout row's Widths say what
// shape the dialect writes; this says what the code indexes. Both are
// checked, because the inferred row derives its widths from the file, so a
// width can be declared and still be too short for the branch that reads
// it.
var specialNeeds = map[string]int{
	"UNIT_DIED":              layout.BaseParams,
	"UNIT_DESTROYED":         layout.BaseParams,
	"UNIT_DISSIPATES":        layout.BaseParams,
	"PARTY_KILL":             layout.BaseParams,
	"SPELL_ABSORBED":         layout.BaseParams,
	"SPELL_HEAL_ABSORBED":    20,
	"ENVIRONMENTAL_DAMAGE":   layout.BaseParams,
	"ENCOUNTER_START":        5,
	"ENCOUNTER_END":          5,
	"ZONE_CHANGE":            3,
	"MAP_CHANGE":             3,
	"CHALLENGE_MODE_START":   5,
	"CHALLENGE_MODE_END":     4,
	"ENCHANT_APPLIED":        12,
	"ENCHANT_REMOVED":        12,
	"EMOTE":                  5,
	"COMBATANT_INFO":         2,
	"SPELL_ABSORBED_SUPPORT": layout.BaseParams,
	"ARENA_MATCH_START":      5,
	"ARENA_MATCH_END":        5,
	"STAGGER_CLEAR":          3,
	"STAGGER_PREVENTED":      4,
	"WORLD_MARKER_PLACED":    5,
	"WORLD_MARKER_REMOVED":   2,
}

// decodeSpecial handles the events that do not follow the prefix/suffix
// pattern. The width has already been checked against the layout row's
// declared widths; specialNeeds checks it again against what each branch
// below actually reads.
func (d *Decoder) decodeSpecial(e Event, ln lexer.Line) Event {
	p := ln.Params
	if need, ok := specialNeeds[e.Name]; ok && len(p) < need {
		return fail(e, ln, fmt.Sprintf("%s has %d fields, the decoder reads %d",
			e.Name, len(p), need))
	}
	switch e.Name {
	case "UNIT_DIED", "UNIT_DESTROYED", "UNIT_DISSIPATES":
		readUnits(&e, p)
		e.Kind = Death
		if len(p) > layout.BaseParams {
			e.Critical = boolOf(p[layout.BaseParams])
		}
	case "PARTY_KILL":
		readUnits(&e, p)
		e.Kind = PartyKill
	case "SPELL_ABSORBED":
		return d.readAbsorbed(e, ln)
	case "SPELL_HEAL_ABSORBED":
		readUnits(&e, p)
		e.Kind = HealAbsorbed
		e.Spell = Spell{ID: intOf(p[9]), Name: nilless(p[10]), School: intOf(p[11])}
		e.ExtraUnit = Unit{GUID: p[12], Name: nilless(p[13]), Flags: hex32(p[14]), Raid: hex32(p[15])}
		e.ExtraSpell = Spell{ID: intOf(p[16]), Name: nilless(p[17]), School: intOf(p[18])}
		e.Amount = optInt(p[19])
		// The Classic row allows width 20: the wiki's suffix stops at
		// absorbed and marks totalAmount as a later addition, so at that
		// width the field is absent and Total stays unset rather than
		// being read from off the end of the line.
		if len(p) > 20 {
			e.Total = optInt(p[20])
		}
	case "ENVIRONMENTAL_DAMAGE":
		return d.readEnvironmental(e, ln)
	case "ENCOUNTER_START":
		e.Kind = EncounterStart
		e.Encounter = &Encounter{
			ID: intOf(p[1]), Name: nilless(p[2]),
			Difficulty: intOf(p[3]), Size: intOf(p[4]),
		}
		if len(p) > 5 {
			e.Encounter.InstanceID = intOf(p[5])
		}
	case "ENCOUNTER_END":
		e.Kind = EncounterEnd
		e.Encounter = &Encounter{
			ID: intOf(p[1]), Name: nilless(p[2]),
			Difficulty: intOf(p[3]), Size: intOf(p[4]),
		}
		if len(p) > 5 {
			e.Encounter.Kill = intOf(p[5]) == 1
		}
	case "ZONE_CHANGE":
		e.Kind = ZoneChange
		e.Zone = &Zone{ID: intOf(p[1]), Name: nilless(p[2])}
		if len(p) > 3 {
			e.Zone.Difficulty = intOf(p[3])
		}
	case "MAP_CHANGE":
		e.Kind = MapChange
		z := &Zone{ID: intOf(p[1]), Name: nilless(p[2])}
		if len(p) >= 7 {
			z.MaxX, z.MinX = floatOf(p[3]), floatOf(p[4])
			z.MaxY, z.MinY = floatOf(p[5]), floatOf(p[6])
		}
		e.Zone = z
	case "CHALLENGE_MODE_START":
		e.Kind = ChallengeModeStart
		e.Zone = &Zone{Name: nilless(p[1]), ID: intOf(p[2])}
		e.Amount = optInt(p[4])
	case "CHALLENGE_MODE_END":
		e.Kind = ChallengeModeEnd
		e.Zone = &Zone{ID: intOf(p[1])}
		e.Critical = OptBool{V: intOf(p[2]) == 1, OK: true}
		e.Amount = optInt(p[3])
	case "ENCHANT_APPLIED", "ENCHANT_REMOVED":
		readUnits(&e, p)
		e.Kind = Enchant
		e.Spell = Spell{Name: nilless(p[9])}
		e.ItemID, e.ItemName = optInt(p[10]), nilless(p[11])
	case "EMOTE":
		// Chat and emote text is dropped at parse: spec section 7,
		// minimization. Only the units are kept.
		e.Kind = Emote
		e.Source = Unit{GUID: p[1], Name: nilless(p[2])}
		e.Dest = Unit{GUID: p[3], Name: nilless(p[4])}
	case "COMBATANT_INFO":
		return d.readCombatant(e, ln)
	case "SPELL_ABSORBED_SUPPORT":
		// The same shape as SPELL_ABSORBED with the supporting player's
		// GUID appended; readAbsorbed walks from the front and stops at
		// the fields it knows, so the GUID is taken off first.
		short := ln
		short.Params = p[:len(p)-1]
		out := d.readAbsorbed(e, short)
		out.Supporter = p[len(p)-1]
		out.Raw = ln.Raw
		return out
	case "ARENA_MATCH_START":
		e.Kind = ArenaMatchStart
		e.Zone = &Zone{ID: intOf(p[1])}
		e.Amount = optInt(p[2])    // bracket
		e.ItemName = nilless(p[3]) // match type, e.g. "Rated Solo Shuffle"
		e.Critical = OptBool{V: intOf(p[4]) == 1, OK: true}
	case "ARENA_MATCH_END":
		e.Kind = ArenaMatchEnd
		e.Amount = optInt(p[1]) // winning team
		e.Total = optInt(p[2])  // duration in seconds
	case "STAGGER_CLEAR":
		e.Kind = StaggerClear
		e.Source = Unit{GUID: p[1]}
		e.Amount = optInt(p[2])
	case "STAGGER_PREVENTED":
		e.Kind = StaggerPrevented
		e.Source = Unit{GUID: p[1]}
		e.Spell = Spell{ID: intOf(p[2])}
		e.Amount = optInt(p[3])
	case "WORLD_MARKER_PLACED":
		e.Kind = WorldMarker
		e.Zone = &Zone{ID: intOf(p[1])}
		e.Amount = optInt(p[2])
		e.Critical = OptBool{V: true, OK: true} // placed
	case "WORLD_MARKER_REMOVED":
		e.Kind = WorldMarker
		e.Amount = optInt(p[1])
		e.Critical = OptBool{V: false, OK: true} // removed
	default:
		e.Kind, e.Raw = Unknown, ln.Raw
	}
	return e
}

// readAbsorbed disambiguates the two SPELL_ABSORBED shapes by counting
// fields, which is the only correct method: a 19-field line is a self
// shield and carries no damage spell, a 22-field line carries one.
func (d *Decoder) readAbsorbed(e Event, ln lexer.Line) Event {
	p := ln.Params
	e.Kind = Absorbed
	readUnits(&e, p)
	i := layout.BaseParams
	// After the common header the line holds, optionally the damage spell
	// (3), then the absorbing unit (4), the shield spell (3), the absorbed
	// amount, optionally the full attempted amount, and the crit flag.
	var hasDamageSpell bool
	switch len(p) - layout.BaseParams {
	case 13, 12:
		hasDamageSpell = true
	case 10, 9:
		hasDamageSpell = false
	default:
		return fail(e, ln, fmt.Sprintf("SPELL_ABSORBED has %d fields, which matches neither shape", len(p)))
	}
	if hasDamageSpell {
		e.Spell = Spell{ID: intOf(p[i]), Name: nilless(p[i+1]), School: intOf(p[i+2])}
		i += 3
	}
	e.ExtraUnit = Unit{GUID: p[i], Name: nilless(p[i+1]), Flags: hex32(p[i+2]), Raid: hex32(p[i+3])}
	i += 4
	e.ExtraSpell = Spell{ID: intOf(p[i]), Name: nilless(p[i+1]), School: intOf(p[i+2])}
	i += 3
	e.Amount = optInt(p[i])
	i++
	if len(p)-i >= 2 {
		e.Total = optInt(p[i])
		i++
	}
	if i < len(p) {
		e.Critical = boolOf(p[i])
	}
	return e
}

// EnvironmentalSpellID is the spell id every ENVIRONMENTAL_DAMAGE line is filed
// under. The log gives it none; zero is the melee swing, and a fall filed as a
// melee hit was the biggest "melee hit" on a raid leader's wipe.
const EnvironmentalSpellID int64 = -1

// environmentalSpell names the environment as the spell a fall or a fire is
// filed under, with the school the damage is: falling, drowning and fatigue
// are physical, fire and lava are fire, slime is nature.
func environmentalSpell(envType string) Spell {
	school := int64(1)
	switch strings.ToUpper(envType) {
	case "FIRE", "LAVA":
		school = 4
	case "SLIME":
		school = 8
	}
	name := envType
	if name == "" {
		name = "Environment"
	} else {
		name = strings.ToUpper(name[:1]) + strings.ToLower(name[1:])
	}
	return Spell{ID: EnvironmentalSpellID, Name: name, School: school}
}

// readEnvironmental reads ENVIRONMENTAL_DAMAGE: the common header, the
// advanced block describing the target, the environmental type, then the
// damage suffix.
func (d *Decoder) readEnvironmental(e Event, ln lexer.Line) Event {
	p := ln.Params
	e.Kind = Damage
	readUnits(&e, p)
	i := layout.BaseParams
	if d.lay.Advanced > 0 {
		if len(p) < i+d.lay.Advanced+1 {
			return fail(e, ln, fmt.Sprintf("ENVIRONMENTAL_DAMAGE has %d fields, too few for the advanced block", len(p)))
		}
		e.Adv = readAdvanced(p[i : i+d.lay.Advanced])
		i += d.lay.Advanced
	}
	if i >= len(p) {
		return fail(e, ln, fmt.Sprintf("ENVIRONMENTAL_DAMAGE has %d fields, too few for the environmental type", len(p)))
	}
	e.EnvType = p[i]
	e.Spell = environmentalSpell(e.EnvType)
	i++
	rest := p[i:]
	spec := d.lay.Suffixes["_DAMAGE"]
	want := max(spec.Params, suffixNeeds("_DAMAGE", spec))
	if len(rest) < want {
		return fail(e, ln, fmt.Sprintf("ENVIRONMENTAL_DAMAGE has %d damage fields, layout %q wants %d",
			len(rest), d.lay.Name, want))
	}
	readDamage(&e, rest, spec)
	return e
}

// readCombatant reads COMBATANT_INFO using the indexes in the layout row.
// A row with no COMBATANT_INFO layout keeps the line raw rather than
// guessing where the gear list starts.
func (d *Decoder) readCombatant(e Event, ln lexer.Line) Event {
	c := d.lay.Combatant
	if !c.Present {
		e.Kind, e.Raw = Unknown, ln.Raw
		return e
	}
	p := ln.Params
	if len(p) != c.Params {
		return fail(e, ln, fmt.Sprintf("COMBATANT_INFO has %d fields, layout %q wants %d",
			len(p), d.lay.Name, c.Params))
	}
	for _, at := range []int{c.SpecIndex, c.TalentIndex, c.PvPTalentIndex, c.GearIndex, c.AuraIndex} {
		if at <= 0 || at >= len(p) {
			return fail(e, ln, fmt.Sprintf("COMBATANT_INFO has %d fields, layout %q indexes field %d",
				len(p), d.lay.Name, at))
		}
	}
	// BorrowIndex is zero on a dialect that writes no borrowed-power
	// field, and field zero is always the event name, so zero is an
	// unambiguous "absent" rather than a missing bounds check.
	if c.BorrowIndex < 0 || c.BorrowIndex >= len(p) {
		return fail(e, ln, fmt.Sprintf("COMBATANT_INFO has %d fields, layout %q indexes field %d",
			len(p), d.lay.Name, c.BorrowIndex))
	}
	stats := make(map[string]int64, len(c.StatIndex))
	for name, at := range c.StatIndex {
		if at <= 0 || at >= len(p) {
			return fail(e, ln, fmt.Sprintf("COMBATANT_INFO has %d fields, layout %q reads %s from field %d",
				len(p), d.lay.Name, name, at))
		}
		stats[name] = intOf(p[at])
	}
	e.Kind = CombatantInfo
	info := &Combatant{
		GUID:       p[1],
		Faction:    intOf(p[2]),
		SpecID:     intOf(p[c.SpecIndex]),
		Stats:      stats,
		PvPTalents: intList(p[c.PvPTalentIndex]),
		Gear:       gearList(p[c.GearIndex]),
		Auras:      auraList(p[c.AuraIndex]),
	}
	// TalentsAreSpellIDs is false for a dialect (version 22) whose talent
	// field is not the flat spell-id tuple Combatant.Talents is typed to
	// hold: intList would still flatten it into leaf integers, but those
	// are an interleaved (nodeID, entryID, rank) mix, a wrong-but-typed
	// value worse than leaving the field empty. See the ledger.
	if c.TalentsAreSpellIDs {
		info.Talents = intList(p[c.TalentIndex])
	}
	if c.BorrowIndex > 0 {
		info.Borrowed = p[c.BorrowIndex]
	}
	var sum, n int64
	for _, it := range info.Gear {
		if it.ID != 0 && it.ItemLevel > 0 {
			sum += it.ItemLevel
			n++
		}
	}
	if n > 0 {
		info.ItemLevel = sum / n
	}
	e.Source = Unit{GUID: info.GUID}
	e.Combatant = info
	return e
}

// trimGroup removes one layer of surrounding brackets or parentheses.
func trimGroup(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 2 && (s[0] == '[' && s[len(s)-1] == ']' || s[0] == '(' && s[len(s)-1] == ')') {
		return s[1 : len(s)-1]
	}
	return s
}

// splitGroup splits a group's contents on top-level commas.
func splitGroup(s string) []string {
	s = trimGroup(s)
	if s == "" {
		return nil
	}
	return lexer.SplitParams(s)
}

// intList reads a flat list of integers such as v16's talent tuple
// "(202751,262111,...)". Version 22's talent field nests one level deeper,
// a list of (nodeID, entryID, rank) triples for the new talent trees, e.g.
// "[(90326,112183,1),(90328,112185,1),...]"; a part that is not itself a
// number is recursed into so every leaf integer still lands in one flat
// list, and a v16 line, which nests nothing, recurses zero times.
func intList(s string) []int64 {
	var out []int64
	for _, p := range splitGroup(s) {
		p = strings.TrimSpace(p)
		if v := optInt(p); v.OK {
			out = append(out, v.V)
			continue
		}
		out = append(out, intList(p)...)
	}
	return out
}

// gearList reads [(itemID, itemLevel, (enchants), (bonusIDs), (gems)), ...].
// One entry per equipment slot; an empty slot is itemID 0 and is kept so the
// slot indexes stay meaningful.
func gearList(s string) []Item {
	entries := splitGroup(s)
	out := make([]Item, 0, len(entries))
	for _, entry := range entries {
		f := splitGroup(entry)
		if len(f) < 2 {
			continue
		}
		it := Item{ID: intOf(strings.TrimSpace(f[0])), ItemLevel: intOf(strings.TrimSpace(f[1]))}
		if len(f) > 2 {
			it.Enchants = intList(f[2])
		}
		if len(f) > 3 {
			it.BonusIDs = intList(f[3])
		}
		if len(f) > 4 {
			it.Gems = intList(f[4])
		}
		out = append(out, it)
	}
	if len(out) == 0 {
		return nil
	}
	return out
}

// auraList reads the auras-at-pull list, which is flat: sourceGUID, spellID,
// sourceGUID, spellID, and so on. An odd trailing field is dropped.
func auraList(s string) []Aura {
	f := splitGroup(s)
	out := make([]Aura, 0, len(f)/2)
	for i := 0; i+1 < len(f); i += 2 {
		out = append(out, Aura{
			SourceGUID: strings.TrimSpace(f[i]),
			SpellID:    intOf(strings.TrimSpace(f[i+1])),
		})
	}
	if len(out) == 0 {
		return nil
	}
	return out
}
