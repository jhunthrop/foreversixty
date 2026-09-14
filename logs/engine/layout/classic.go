// logs/engine/layout/classic.go
package layout

// ClassicWiki is the shape warcraft.wiki.gg documents for COMBAT_LOG_EVENT
// without the retail Shadowlands additions: no baseAmount on damage, no
// healedToHP on heals, amountMissed always present on _MISSED, and a
// trailing isOffHand on swings. It is NOT verified against a real Classic
// log, so Verified is false: the decoder reports a width mismatch instead of
// mis-reading a field, and the conformance command lists what it saw.
//
// COMBATANT_INFO is deliberately absent: no document I could source
// describes a Classic layout for it, and guessing the field order would put
// wrong gear and wrong talents on a report.
//
// The version number is 0 because no source states which
// COMBAT_LOG_VERSION Classic clients write; Lookup therefore never selects
// this row automatically. Select it with the CLI's -layout flag, or let the
// inferred row take over.
func ClassicWiki() Layout {
	return Layout{
		Name:      "classic-wiki",
		Version:   0,
		ProjectID: 0,
		Advanced:  17, // present only when the header says ADVANCED_LOG_ENABLED,1
		StampYear: false,
		StampZone: false,
		Verified:  false,
		Prefixes: map[string]int{
			"SWING":          0,
			"RANGE":          3,
			"SPELL_PERIODIC": 3,
			"SPELL_BUILDING": 3,
			"SPELL":          3,
		},
		Suffixes: map[string]Suffix{
			// amount, overkill, school, resisted, blocked, absorbed,
			// critical, glancing, crushing (+ isOffHand).
			"_DAMAGE":        {Params: 9, Advanced: true, OffHand: true},
			"_DAMAGE_LANDED": {Params: 9, Advanced: true, OffHand: true},
			// amount, overhealing, absorbed, critical.
			"_HEAL": {Params: 4, Advanced: true},
			// missType, isOffHand (+ amountMissed, critical on ABSORB).
			"_MISSED":            {Params: 2, AbsorbExtra: 2},
			"_ENERGIZE":          {Params: 4, Advanced: true},
			"_DRAIN":             {Params: 4, Advanced: true},
			"_LEECH":             {Params: 4, Advanced: true},
			"_AURA_APPLIED":      {Params: 1},
			"_AURA_REMOVED":      {Params: 1},
			"_AURA_REFRESH":      {Params: 1},
			"_AURA_BROKEN":       {Params: 1},
			"_AURA_APPLIED_DOSE": {Params: 2},
			"_AURA_REMOVED_DOSE": {Params: 2},
			"_AURA_BROKEN_SPELL": {Params: 4},
			"_INTERRUPT":         {Params: 3},
			"_DISPEL_FAILED":     {Params: 3},
			"_DISPEL":            {Params: 4},
			"_STOLEN":            {Params: 4},
			"_CAST_START":        {Params: 0},
			"_CAST_SUCCESS":      {Params: 0, Advanced: true},
			"_CAST_FAILED":       {Params: 1},
			"_SUMMON":            {Params: 0},
			"_CREATE":            {Params: 0},
			"_RESURRECT":         {Params: 0},
			"_INSTAKILL":         {Params: 1},
			"_EXTRA_ATTACKS":     {Params: 1},
			"_DURABILITY_DAMAGE": {Params: 0},
		},
		Specials: map[string]Special{
			"COMBAT_LOG_VERSION": {},
			"UNIT_DIED":          {Widths: []int{9, 10}},
			"UNIT_DESTROYED":     {Widths: []int{9, 10}},
			"UNIT_DISSIPATES":    {Widths: []int{9, 10}},
			"PARTY_KILL":         {Widths: []int{9, 10}},
			// The wiki's suffix stops at absorbedAmount and marks
			// totalAmount as a later addition, so both widths are allowed.
			"SPELL_ABSORBED":       {Widths: []int{18, 19, 21, 22}},
			"SPELL_HEAL_ABSORBED":  {Widths: []int{20, 21}},
			"ENCOUNTER_START":      {},
			"ENCOUNTER_END":        {},
			"ZONE_CHANGE":          {},
			"MAP_CHANGE":           {},
			"ENCHANT_APPLIED":      {Widths: []int{12}},
			"ENCHANT_REMOVED":      {Widths: []int{12}},
			"EMOTE":                {},
			"ENVIRONMENTAL_DAMAGE": {},
		},
		Combatant: Combatant{Present: false},
	}
}
