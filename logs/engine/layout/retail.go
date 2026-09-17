// logs/engine/layout/retail.go
package layout

// RetailV16 is the retail Shadowlands dialect, COMBAT_LOG_VERSION 16,
// PROJECT_ID 1. Every count below was verified by parsing a real 272,367
// event log (build 9.0.2, advanced logging on, 105 COMBATANT_INFO lines).
// MAP_CHANGE is the one shape absent from that log; its width comes from
// wowcoach.gg/docs/combat-log/spec.yaml.
func RetailV16() Layout {
	l := Layout{
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
			Present:            true,
			Params:             34,
			SpecIndex:          24,
			TalentIndex:        25,
			PvPTalentIndex:     26,
			BorrowIndex:        27,
			GearIndex:          28,
			AuraIndex:          29,
			TalentsAreSpellIDs: true,
			StatIndex: map[string]int{
				"strength": 3, "agility": 4, "stamina": 5, "intellect": 6,
				"dodge": 7, "parry": 8, "block": 9, "crit": 10, "speed": 13,
				"lifesteal": 14, "haste": 15, "avoidance": 18, "mastery": 19,
				"versatility": 20, "armor": 23,
			},
		},
	}
	l.widthCache = buildWidthCache(l)
	return l
}

// RetailV22 is the modern retail dialect, COMBAT_LOG_VERSION 22,
// PROJECT_ID 1. Every count below was measured over 89 real logs,
// 26,030,980 lines, builds 12.0.5, 12.0.7 and 12.1.0, all with advanced
// logging on. The field-by-field derivation is in
// docs/superpowers/plans/2026-09-16-retail-v22-layout.md; the fields that
// could not be identified are in docs/ledger/2026-09-16-retail-v22.md.
//
// Three things changed from v16 and drive everything here: the advanced
// block grew from 17 fields to 19, spell-prefixed damage and miss lines
// gained a trailing single-target / area tag, and COMBATANT_INFO's stat
// fields all moved one to the right.
func RetailV22() Layout {
	l := Layout{
		Name:      "retail-v22",
		Version:   22,
		ProjectID: 1,
		Advanced:  19,
		StampYear: true,
		StampZone: true,
		Verified:  true,
		Prefixes: map[string]int{
			"SWING":          0,
			"RANGE":          3,
			"SPELL_PERIODIC": 3,
			"SPELL_BUILDING": 3,
			"SPELL":          3,
			// DAMAGE_SPLIT is a spell-prefixed damage line whose name
			// happens not to start with SPELL.
			"DAMAGE": 3,
			// DAMAGE_SHIELD_MISSED likewise. The longer prefix wins over
			// "DAMAGE" because Split prefers the one that leaves a known
			// suffix, and "_SHIELD_MISSED" is not a suffix.
			"DAMAGE_SHIELD": 3,
			// The four prefixes below exist only so that the "_SUPPORT"
			// suffix attaches to the right spell-triple offset. A support
			// line names the supporting spell, not the supported one, so
			// even SWING_DAMAGE_LANDED_SUPPORT carries a spell triple and
			// is 42 fields wide where SWING_DAMAGE_LANDED is 38.
			"SPELL_DAMAGE":          3,
			"SPELL_PERIODIC_DAMAGE": 3,
			"RANGE_DAMAGE":          3,
			"SWING_DAMAGE_LANDED":   3,
		},
		Suffixes: map[string]Suffix{
			// amount, baseAmount, overkill, school, resisted, blocked,
			// absorbed, critical, glancing, crushing, then the optional
			// "ST" / "AOE" tag.
			"_DAMAGE":        {Params: 10, Advanced: true, BaseAmount: true, Tag: true},
			"_DAMAGE_LANDED": {Params: 10, Advanced: true, BaseAmount: true, Tag: true},
			"_SPLIT":         {Params: 10, Advanced: true, BaseAmount: true, Tag: true},
			// The same ten damage fields, then the supporting player's
			// GUID in place of the tag.
			"_SUPPORT": {Params: 11, Advanced: true, BaseAmount: true},
			// healedToHP, amount, overheal, absorbed, critical.
			"_HEAL": {Params: 5, Advanced: true, HealedToHP: true},
			// The same five, then the supporting player's GUID.
			"_HEAL_SUPPORT": {Params: 6, Advanced: true, HealedToHP: true},
			// missType, isOffHand; three more on ABSORB, one more on
			// BLOCK and RESIST, and the tag after whichever of those
			// applies.
			"_MISSED": {Params: 2, AbsorbExtra: 3, MissAmount: true, Tag: true},
			// amount, overEnergize, powerType, maxPower.
			"_ENERGIZE": {Params: 4, Advanced: true},
			"_DRAIN":    {Params: 4, Advanced: true},
			"_LEECH":    {Params: 4, Advanced: true},
			// auraType, then the absorb size, then one more number whose
			// meaning is not pinned.
			"_AURA_APPLIED":      {Params: 1, AuraExtra: true},
			"_AURA_REMOVED":      {Params: 1, AuraExtra: true},
			"_AURA_REFRESH":      {Params: 1, AuraExtra: true},
			"_AURA_BROKEN":       {Params: 1, AuraExtra: true},
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
			// Evoker empowered casts: start carries nothing extra, end and
			// interrupt carry the empowerment stage reached.
			"_EMPOWER_START":     {Params: 0},
			"_EMPOWER_END":       {Params: 1},
			"_EMPOWER_INTERRUPT": {Params: 1},
		},
		Specials: map[string]Special{
			"COMBAT_LOG_VERSION":  {Widths: []int{8}},
			"UNIT_DIED":           {Widths: []int{10}},
			"UNIT_DESTROYED":      {Widths: []int{10}},
			"UNIT_DISSIPATES":     {Widths: []int{10}},
			"PARTY_KILL":          {Widths: []int{10}},
			"SPELL_ABSORBED":      {Widths: []int{19, 22}},
			"SPELL_HEAL_ABSORBED": {Widths: []int{21}},
			// The absorb shapes plus the supporting player's GUID.
			"SPELL_ABSORBED_SUPPORT": {Widths: []int{20, 23}},
			"ENCOUNTER_START":        {Widths: []int{6}},
			"ENCOUNTER_END":          {Widths: []int{6}},
			"ZONE_CHANGE":            {Widths: []int{4}},
			"MAP_CHANGE":             {Widths: []int{7}},
			"CHALLENGE_MODE_START":   {Widths: []int{6}},
			"CHALLENGE_MODE_END":     {Widths: []int{5}},
			"ENCHANT_APPLIED":        {Widths: []int{12}},
			"ENCHANT_REMOVED":        {Widths: []int{12}},
			"EMOTE":                  {Widths: []int{6}},
			"COMBATANT_INFO":         {Widths: []int{34}},
			// 9 header + 19 advanced + environmental type + 10 damage.
			"ENVIRONMENTAL_DAMAGE": {Widths: []int{39}},
			// instanceID, bracket, matchType, isRated.
			"ARENA_MATCH_START": {Widths: []int{5}},
			// winningTeam, duration, newRatingTeam1, newRatingTeam2.
			"ARENA_MATCH_END": {Widths: []int{5}},
			// playerGUID, amount cleared.
			"STAGGER_CLEAR": {Widths: []int{3}},
			// playerGUID, spellID, amount prevented.
			"STAGGER_PREVENTED": {Widths: []int{4}},
			// uiMapID, markerIndex, x, y.
			"WORLD_MARKER_PLACED": {Widths: []int{5}},
			// markerIndex.
			"WORLD_MARKER_REMOVED": {Widths: []int{2}},
		},
		Combatant: Combatant{
			Present:        true,
			Params:         34,
			SpecIndex:      25,
			TalentIndex:    26,
			PvPTalentIndex: 27,
			// Version 22 writes no borrowed-power field; the line is the
			// same width as v16's because the stat block gained one field
			// and this one went away.
			BorrowIndex: 0,
			GearIndex:   28,
			AuraIndex:   29,
			// Version 22's talent trees write (nodeID, entryID, rank)
			// triples at this index, not the flat spell-id tuple
			// Combatant.Talents is typed to hold, so TalentsAreSpellIDs
			// stays false and the decoder emits an empty Talents rather
			// than a flattened mix of the three numbers. See the ledger.
			TalentsAreSpellIDs: false,
			// Every stat sits one field later than in v16: v22 writes four
			// defensive ratings where v16 writes dodge, parry and block.
			// All four are zero in every sample, so the fourth is counted
			// and not named; see the ledger.
			StatIndex: map[string]int{
				"strength": 3, "agility": 4, "stamina": 5, "intellect": 6,
				"dodge": 7, "parry": 8, "block": 9, "crit": 11, "speed": 14,
				"lifesteal": 15, "haste": 16, "avoidance": 19, "mastery": 20,
				"versatility": 21, "armor": 24,
			},
		},
		// The flag cross-product in widthsFor is a superset of what the
		// corpus actually writes for these thirteen shapes: the
		// single-target/area tag turns out to be mandatory wherever it
		// appears at all (never optional), and which prefix carries it is
		// not derivable from the Suffix flags a shape shares with other
		// prefixes. Measured over the same 89 logs as the rest of this row;
		// see docs/ledger/2026-09-16-retail-v22.md and
		// measured_v22_test.go's v22MeasuredWidths, which this pins against.
		WidthOverrides: map[string][]int{
			// _DAMAGE, _DAMAGE_LANDED and _SPLIT: the tag is mandatory
			// under a spell-triple prefix and never appears under SWING.
			"SWING_DAMAGE":           {38},
			"SWING_DAMAGE_LANDED":    {38},
			"SPELL_DAMAGE":           {42},
			"SPELL_PERIODIC_DAMAGE":  {42},
			"RANGE_DAMAGE":           {42},
			"DAMAGE_SPLIT":           {42},
			// _MISSED: the tag is mandatory under SPELL and SPELL_PERIODIC,
			// and never appears under RANGE, SWING or DAMAGE_SHIELD. Where
			// it applies, base+1 is tag only, base+2 is tag+amount (BLOCK
			// or RESIST), base+4 is tag+the three ABSORB extras; where it
			// does not, only base and base+3 (ABSORB, untagged) occur.
			"SPELL_MISSED":          {15, 16, 18},
			"SPELL_PERIODIC_MISSED": {15, 18},
			"RANGE_MISSED":          {14, 17},
			"SWING_MISSED":          {11, 14},
			"DAMAGE_SHIELD_MISSED":  {15},
			// _AURA_BROKEN never carries the optional absorb size or its
			// extra trailing number; _AURA_REFRESH carries the absorb size
			// but never the extra number. Both are narrower than
			// AuraCarriesAmount + AuraExtra alone would allow.
			"SPELL_AURA_BROKEN":  {13},
			"SPELL_AURA_REFRESH": {13, 14},
		},
	}
	l.widthCache = buildWidthCache(l)
	return l
}
