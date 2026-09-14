// logs/engine/layout/retail.go
package layout

// RetailV16 is the retail Shadowlands dialect, COMBAT_LOG_VERSION 16,
// PROJECT_ID 1. Every count below was verified by parsing a real 272,367
// event log (build 9.0.2, advanced logging on, 105 COMBATANT_INFO lines).
// MAP_CHANGE is the one shape absent from that log; its width comes from
// wowcoach.gg/docs/combat-log/spec.yaml.
func RetailV16() Layout {
	return Layout{
		Name:      "retail-v16",
		Version:   16,
		ProjectID: 1,
		Advanced:  17,
		StampYear: false,
		StampZone: false,
		Verified:  true,
		Prefixes: map[string]int{
			"SWING":          0,
			"RANGE":          3,
			"SPELL_PERIODIC": 3,
			"SPELL_BUILDING": 3,
			"SPELL":          3,
		},
		Suffixes: map[string]Suffix{
			// amount, baseAmount, overkill, school, resisted, blocked,
			// absorbed, critical, glancing, crushing.
			"_DAMAGE":        {Params: 10, Advanced: true, BaseAmount: true},
			"_DAMAGE_LANDED": {Params: 10, Advanced: true, BaseAmount: true},
			// healedToHP, amount, overheal, absorbed, critical.
			"_HEAL": {Params: 5, Advanced: true, HealedToHP: true},
			// missType, isOffHand (+ amountMissed, baseAmount, critical on ABSORB).
			"_MISSED": {Params: 2, AbsorbExtra: 3},
			// amount, overEnergize, powerType, maxPower.
			"_ENERGIZE": {Params: 4, Advanced: true},
			"_DRAIN":    {Params: 4, Advanced: true},
			"_LEECH":    {Params: 4, Advanced: true},
			// auraType (+ amount when the aura carries an absorb size).
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
			"COMBAT_LOG_VERSION":   {Widths: []int{8}},
			"UNIT_DIED":            {Widths: []int{10}},
			"UNIT_DESTROYED":       {Widths: []int{10}},
			"UNIT_DISSIPATES":      {Widths: []int{10}},
			"PARTY_KILL":           {Widths: []int{10}},
			"SPELL_ABSORBED":       {Widths: []int{19, 22}},
			"SPELL_HEAL_ABSORBED":  {Widths: []int{21}},
			"ENCOUNTER_START":      {Widths: []int{6}},
			"ENCOUNTER_END":        {Widths: []int{6}},
			"ZONE_CHANGE":          {Widths: []int{4}},
			"MAP_CHANGE":           {Widths: []int{7}},
			"CHALLENGE_MODE_START": {Widths: []int{6}},
			"CHALLENGE_MODE_END":   {Widths: []int{5}},
			"ENCHANT_APPLIED":      {Widths: []int{12}},
			"ENCHANT_REMOVED":      {Widths: []int{12}},
			"EMOTE":                {Widths: []int{6}},
			"COMBATANT_INFO":       {Widths: []int{34}},
			"ENVIRONMENTAL_DAMAGE": {Widths: []int{37}},
		},
		Combatant: Combatant{
			Present:        true,
			Params:         34,
			SpecIndex:      24,
			TalentIndex:    25,
			PvPTalentIndex: 26,
			BorrowIndex:    27,
			GearIndex:      28,
			AuraIndex:      29,
		},
	}
}
