// logs/engine/summary/summary.go
// Package summary turns a fight's events into the tables the report page
// opens without a query. Everything is a running accumulator fed one event
// at a time, so Snapshot during a live fight costs no more than a sort.
package summary

import (
	"sort"
	"time"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
	"github.com/jhunthrop/foreversixty/logs/engine/fight"
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics"
	"github.com/jhunthrop/foreversixty/logs/engine/units"
)

// Options tunes the accumulators. Every lookup table is injected: the
// engine ships no spell ids of its own.
type Options struct {
	// Bucket is the width of the per-second series. Default one second.
	Bucket time.Duration
	// ActiveGap is how long an actor counts as active after each action.
	// Default 1.5 seconds.
	ActiveGap time.Duration
	// DeathHealWindow is how many heals to keep before each death; more
	// than the damage window because a healer's ticks come thick.
	DeathHealWindow int
	// DeathWindow is how many damage events to keep before each death.
	// Default ten, which is what the deaths view shows.
	DeathWindow int
	// DeathAuraWindow is how far back to look for auras applied and
	// removed around a death. Default ten seconds.
	DeathAuraWindow time.Duration
	// Registry names units and resolves pet owners. Required.
	Registry *units.Registry
	// Threat computes threat per point of damage and healing. Default
	// BaseThreat{}, which reports itself as incomplete.
	Threat ThreatModel
	// ConsumableSpells maps a spell id to a consumable name, for the
	// consumables-at-pull table. Empty means the table is empty.
	ConsumableSpells map[int64]string
	// RaidBuffSpells maps a spell id to a raid buff name, for the buffs-at-
	// pull check. Empty means the check is skipped.
	RaidBuffSpells map[int64]string
	// SpecNames maps a COMBATANT_INFO spec id to the spec's name.
	SpecNames map[int64]string
	// Mechanics is the curated table for the fight's encounter; nil when there is none.
	Mechanics *mechanics.Table
}

// DefaultOptions returns the values the session uses. Registry must still
// be set by the caller.
func DefaultOptions() Options {
	return Options{
		Bucket:          time.Second,
		ActiveGap:       1500 * time.Millisecond,
		DeathWindow:     10,
		DeathHealWindow: 40,
		DeathAuraWindow: 10 * time.Second,
		Threat:          BaseThreat{},
	}
}

func (o Options) withDefaults() Options {
	d := DefaultOptions()
	if o.Bucket <= 0 {
		o.Bucket = d.Bucket
	}
	if o.ActiveGap <= 0 {
		o.ActiveGap = d.ActiveGap
	}
	if o.DeathWindow <= 0 {
		o.DeathWindow = d.DeathWindow
	}
	if o.DeathHealWindow <= 0 {
		o.DeathHealWindow = d.DeathHealWindow
	}
	if o.DeathAuraWindow <= 0 {
		o.DeathAuraWindow = d.DeathAuraWindow
	}
	if o.Threat == nil {
		o.Threat = d.Threat
	}
	return o
}

// Summary is everything precomputed for one fight.
type Summary struct {
	EngineVersion string `json:"engine_version"`
	FightIndex    int    `json:"fight_index"`
	DurationMS    int64  `json:"duration_ms"`

	DamageDone   []Actor `json:"damage_done"`
	DamageTaken  []Actor `json:"damage_taken"`
	Healing      []Actor `json:"healing"`
	HealingTaken []Actor `json:"healing_taken"`

	Deaths     []Death         `json:"deaths"`
	Auras      []AuraTrack     `json:"auras"`
	Casts      []CastRow       `json:"casts"`
	Interrupts []ExchangeRow   `json:"interrupts"`
	Dispels    []ExchangeRow   `json:"dispels"`
	Resources  []ResourceTrack `json:"resources"`
	Threat     []ThreatRow     `json:"threat"`
	Combatants []CombatantRow  `json:"combatants"`
	Roster     []RosterRow     `json:"roster"`

	Mechanics MechanicsBlock `json:"mechanics"`
}

// Accumulator folds a fight's events into a Summary.
type Accumulator struct {
	opt   Options
	start time.Time
	end   time.Time

	damageDone   map[string]*actor
	damageTaken  map[string]*actor
	healingDone  map[string]*actor
	healingTaken map[string]*actor

	active      map[string]*activity
	deaths      []*death
	recent      map[string][]DamageRef
	recentHeals map[string][]HealRef
	auras       map[auraKey]*auraTrack
	// spellNames is every spell id an event named, so a track seeded from
	// a combatant snapshot can be named by any line that mentions the
	// spell, not only an aura line on the same target.
	spellNames map[int64]string
	casts      map[castKey]*castRow
	pending    map[castKey]time.Time
	exchanges  map[exchangeKey]*ExchangeRow
	resources  map[resourceKey]*resourceTrack
	threat     map[string]float64
	combatants map[string]*event.Combatant
	// mechanicHits is spell id -> player guid -> the hit tally, folded from
	// the Damage case for every spell the fight's mechanics table lists.
	mechanicHits map[int64]map[string]*MechanicHit
}

// New returns an accumulator for one fight.
func New(o Options) *Accumulator {
	return &Accumulator{
		opt:          o.withDefaults(),
		damageDone:   map[string]*actor{},
		damageTaken:  map[string]*actor{},
		healingDone:  map[string]*actor{},
		healingTaken: map[string]*actor{},
		active:       map[string]*activity{},
		recent:       map[string][]DamageRef{},
		recentHeals:  map[string][]HealRef{},
		auras:        map[auraKey]*auraTrack{},
		spellNames:   map[int64]string{},
		casts:        map[castKey]*castRow{},
		pending:      map[castKey]time.Time{},
		exchanges:    map[exchangeKey]*ExchangeRow{},
		resources:    map[resourceKey]*resourceTrack{},
		threat:       map[string]float64{},
		combatants:   map[string]*event.Combatant{},
		mechanicHits: map[int64]map[string]*MechanicHit{},
	}
}

// Start marks the fight's beginning, which the per-second series and every
// millisecond offset are measured from.
func (a *Accumulator) Start(at time.Time) {
	a.start, a.end = at, at
}

// ms is the offset of t from the fight's start, never negative.
func (a *Accumulator) ms(t time.Time) int64 {
	if t.Before(a.start) {
		return 0
	}
	return t.Sub(a.start).Milliseconds()
}

func (a *Accumulator) bucket(t time.Time) int {
	if t.Before(a.start) {
		return 0
	}
	return int(t.Sub(a.start) / a.opt.Bucket)
}

// name resolves a display name through the registry.
func (a *Accumulator) name(guid string) string {
	// The game's null GUID is the source of falling, drowning and the
	// fight's own hazards; a table that prints it as sixteen zeros tells
	// the reader nothing.
	if guid == units.NoGUID {
		return "Environment"
	}
	if a.opt.Registry == nil {
		return guid
	}
	return a.opt.Registry.Name(guid)
}

// owner resolves a pet to its player so pet damage lands on the player's
// row, which is what every table in the report shows.
func (a *Accumulator) owner(guid string) string {
	if a.opt.Registry == nil {
		return guid
	}
	return a.opt.Registry.Owner(guid)
}

// Add folds one event in. Events must arrive in the order they were logged.
func (a *Accumulator) Add(e event.Event) {
	if a.start.IsZero() {
		a.start = e.Time
	}
	if e.Time.After(a.end) {
		a.end = e.Time
	}
	if e.Spell.ID != 0 && e.Spell.Name != "" {
		a.spellNames[e.Spell.ID] = e.Spell.Name
	}
	a.addDamageAndHealing(e)
	a.addCastsAndExchanges(e)
	a.addAuras(e)
	a.addResources(e)
	a.addDeaths(e)
	if e.Kind == event.CombatantInfo && e.Combatant != nil {
		a.combatants[e.Combatant.GUID] = e.Combatant
		a.seedAuras(e)
	}
}

// seedAuras opens a track for every aura the combatant snapshot says was
// already up at the pull. The log writes no APPLIED line for those, so
// without this a flask, a raid buff or a mitigation cast just before the
// pull read as absent until it dropped, and uptime was a floor. The
// snapshot carries only spell ids: the track is named by the first aura
// event that mentions the spell, and "Spell #id" if none ever does.
func (a *Accumulator) seedAuras(e event.Event) {
	for _, aura := range e.Combatant.Auras {
		if aura.SpellID == 0 {
			continue
		}
		key := auraKey{target: e.Combatant.GUID, spellID: aura.SpellID}
		if _, ok := a.auras[key]; ok {
			continue
		}
		tr := &auraTrack{appliers: map[string]bool{}}
		tr.TargetGUID, tr.TargetName = e.Combatant.GUID, a.name(e.Combatant.GUID)
		tr.SpellID, tr.Name, tr.Type = aura.SpellID, "", "BUFF"
		tr.open, tr.openAt, tr.openStacks = true, e.Time, 1
		tr.openSource = aura.SourceGUID
		if aura.SourceGUID != "" && aura.SourceGUID != units.NoGUID {
			tr.appliers[aura.SourceGUID] = true
		}
		a.auras[key] = tr
	}
}

// Snapshot renders the tables. It sorts and copies but does not mutate the
// accumulator, and every slice and map it hands out is the caller's own, so
// it is safe to call every few seconds during a live fight and to serialise
// the result while the parse goes on.
func (a *Accumulator) Snapshot(f fight.Fight, engineVersion string) Summary {
	// The fight's own wall length, the same figure the fight list shows, so a pull has
	// one length everywhere and every per-second figure divides by it. A fight still
	// open has no end yet, so it runs to the last event seen.
	dur := a.end.Sub(a.start)
	if !f.InProgress && f.End.After(f.Start) {
		dur = f.End.Sub(f.Start)
	}
	if dur < 0 {
		dur = 0
	}
	s := Summary{
		EngineVersion: engineVersion,
		FightIndex:    f.Index,
		DurationMS:    dur.Milliseconds(),
		DamageDone:    a.actors(a.damageDone),
		DamageTaken:   a.actors(a.damageTaken),
		Healing:       a.actors(a.healingDone),
		HealingTaken:  a.actors(a.healingTaken),
		Deaths:        a.deathRows(),
		Auras:         a.auraRows(),
		Casts:         a.castRows(),
		Interrupts:    a.exchangeRows("interrupt"),
		Dispels:       a.exchangeRows("dispel"),
		Resources:     a.resourceRows(),
		Threat:        a.threatRows(),
		Combatants:    a.combatantRows(),
	}
	s.Mechanics = a.mechanicsBlock(s.Deaths)
	s.Roster = a.rosterRows(f, s)
	return s
}

// copySlice returns a copy of s. Snapshot hands its result to a caller
// that may hold it while the accumulator keeps folding events in, so every
// slice that comes out of accumulator state is copied on the way: an
// already-rendered snapshot must never change underneath its reader, and a
// companion serialising one off the parse goroutine must not race.
func copySlice[T any](s []T) []T {
	if s == nil {
		return nil
	}
	return append(make([]T, 0, len(s)), s...)
}

// copyMap is copySlice for the maps Snapshot hands out.
func copyMap[K comparable, V any](m map[K]V) map[K]V {
	if m == nil {
		return nil
	}
	c := make(map[K]V, len(m))
	for k, v := range m {
		c[k] = v
	}
	return c
}

// sortActors orders rows by effective amount descending, then GUID, so two
// runs over the same input produce the same bytes.
func sortActors(rows []Actor) {
	sort.Slice(rows, func(i, j int) bool {
		if rows[i].Effective != rows[j].Effective {
			return rows[i].Effective > rows[j].Effective
		}
		return rows[i].GUID < rows[j].GUID
	})
}
