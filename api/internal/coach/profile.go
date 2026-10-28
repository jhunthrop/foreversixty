package coach

import "strconv"

// Profile is one player's measurements on one fight, the §11.1 coach_profiles.profile
// jsonb (about 1 KB). The pool's median is a Profile too, which F1–F8 read through the
// accessors. Task 17 builds it from a fight and writes it.
type Profile struct {
	AliveMS         int64              `json:"alive_ms"`
	CountedMS       int64              `json:"counted_ms"`
	ActivePct       float64            `json:"active_pct"`
	CastsPerMin     map[string]float64 `json:"cpm"`    // ability family -> casts per counted minute
	Uptime          map[string]float64 `json:"uptime"` // owned aura spell id -> % (sole applier only)
	CappedPct       float64            `json:"capped_pct"`
	ManaLowMS       int64              `json:"mana_low_ms"`
	ManaLowFirstMS  *int64             `json:"mana_low_first_ms,omitempty"`
	OverhealByRank  map[string]float64 `json:"overheal"` // heal spell id (the rank) -> %
	DefensiveUptime map[string]float64 `json:"defensive"`
	Adherence       *float64           `json:"adherence,omitempty"`
	Consumables     map[string]bool    `json:"consumables"` // category -> present at the pull
}

// CastsPerMinOf is the casts per counted minute of an ability family; ok is false when
// the profile has no value for it.
func (p Profile) CastsPerMinOf(family string) (float64, bool) {
	v, ok := p.CastsPerMin[family]
	return v, ok
}

// UptimeOf is an owned aura's uptime %.
func (p Profile) UptimeOf(spellID int64) (float64, bool) { return bySpellID(p.Uptime, spellID) }

// OverhealOf is a heal rank's overheal %.
func (p Profile) OverhealOf(spellID int64) (float64, bool) {
	return bySpellID(p.OverhealByRank, spellID)
}

// DefensiveOf is a defensive aura's uptime %.
func (p Profile) DefensiveOf(spellID int64) (float64, bool) {
	return bySpellID(p.DefensiveUptime, spellID)
}

// AdherenceOf is the rotation adherence, 0 to 1.
func (p Profile) AdherenceOf() (float64, bool) {
	if p.Adherence == nil {
		return 0, false
	}
	return *p.Adherence, true
}

// SpellKey is the profile maps' key for a spell id (jsonb object keys are strings).
func SpellKey(spellID int64) string { return strconv.FormatInt(spellID, 10) }

func bySpellID(m map[string]float64, spellID int64) (float64, bool) {
	v, ok := m[SpellKey(spellID)]
	return v, ok
}
