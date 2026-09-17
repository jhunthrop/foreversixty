// logs/engine/summary/deaths.go
package summary

import (
	"fmt"
	"slices"
	"sort"
	"time"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
	"github.com/jhunthrop/foreversixty/logs/engine/units"
)

// DamageRef is one damage event as the deaths view shows it.
type DamageRef struct {
	AtMS       int64  `json:"at_ms"`
	SourceGUID string `json:"source_guid"`
	SourceName string `json:"source_name"`
	SpellID    int64  `json:"spell_id"`
	SpellName  string `json:"spell_name"`
	Amount     int64  `json:"amount"`
	Overkill   int64  `json:"overkill,omitempty"`
	Absorbed   int64  `json:"absorbed,omitempty"`
	HPAfter    int64  `json:"hp_after,omitempty"`
	MaxHP      int64  `json:"max_hp,omitempty"`
}

// HealRef is one heal on the dying player as the deaths view shows it: the
// question a healer asks of every death is whether anyone was healing them,
// and the damage list alone cannot answer it.
type HealRef struct {
	AtMS       int64  `json:"at_ms"`
	SourceGUID string `json:"source_guid"`
	SourceName string `json:"source_name"`
	SpellID    int64  `json:"spell_id"`
	SpellName  string `json:"spell_name"`
	Amount     int64  `json:"amount"`
	Overheal   int64  `json:"overheal,omitempty"`
	Absorbed   int64  `json:"absorbed,omitempty"`
	// HPAfter and MaxHP are the target's health after the heal, from the
	// advanced fields, so a death card can show the step a heal made.
	HPAfter int64 `json:"hp_after,omitempty"`
	MaxHP   int64 `json:"max_hp,omitempty"`
}

// AuraRef is one aura as the deaths and combatants views show it.
type AuraRef struct {
	SpellID    int64  `json:"spell_id"`
	Name       string `json:"name"`
	SourceGUID string `json:"source_guid,omitempty"`
	Type       string `json:"type,omitempty"`
	AtMS       int64  `json:"at_ms"`
	Stacks     int64  `json:"stacks,omitempty"`
}

// Death is one player death.
type Death struct {
	GUID        string      `json:"guid"`
	Name        string      `json:"name"`
	Class       string      `json:"class,omitempty"`
	AtMS        int64       `json:"at_ms"`
	KillingBlow *DamageRef  `json:"killing_blow,omitempty"`
	Last        []DamageRef `json:"last"`
	// Heals is the last DeathWindow heals landed on the player before the
	// death, in order, alongside Last's damage.
	Heals     []HealRef `json:"heals"`
	AurasHeld []AuraRef `json:"auras_held"`
	AurasLost []AuraRef `json:"auras_lost"`
	ReleaseMS int64     `json:"release_ms,omitempty"`
}

type death struct {
	Death
	at time.Time
}

// Segment is one span during which an aura was up.
type Segment struct {
	StartMS    int64  `json:"start_ms"`
	EndMS      int64  `json:"end_ms"`
	Stacks     int64  `json:"stacks"`
	SourceGUID string `json:"source_guid,omitempty"`
}

// AuraTrack is one aura on one target across the fight.
type AuraTrack struct {
	TargetGUID string `json:"target_guid"`
	TargetName string `json:"target_name"`
	SpellID    int64  `json:"spell_id"`
	// School is the spell's school mask, so a debuff can say Magic from Physical.
	School       int64     `json:"school,omitempty"`
	Name         string    `json:"name"`
	Type         string    `json:"type"`
	Applications int64     `json:"applications"`
	MaxStacks    int64     `json:"max_stacks"`
	UptimeMS     int64     `json:"uptime_ms"`
	Segments     []Segment `json:"segments"`
	Appliers     []string  `json:"appliers"`
}

type auraKey struct {
	target  string
	spellID int64
}

type auraTrack struct {
	AuraTrack
	openAt     time.Time
	open       bool
	openStacks int64
	openSource string
	appliers   map[string]bool
}

// CastRow is one caster's use of one spell.
type CastRow struct {
	GUID string `json:"guid"`
	Name string `json:"name"`
	// OwnerGUID is the caster's owner when the caster is a pet or a guardian,
	// and the caster's own GUID otherwise, so the Casts tab can keep a pet's
	// rows under the player who owns it the way the damage tables already keep
	// a pet's damage. The row itself stays the pet's, name and all: the web
	// prints "via Ashfang" and needs the two apart.
	OwnerGUID   string           `json:"owner_guid"`
	SpellID     int64            `json:"spell_id"`
	SpellName   string           `json:"spell_name"`
	Started     int64            `json:"started"`
	Succeeded   int64            `json:"succeeded"`
	Failed      int64            `json:"failed"`
	FailReasons map[string]int64 `json:"fail_reasons,omitempty"`
	CastTimeMS  int64            `json:"cast_time_ms"`
	Sequence    []int64          `json:"sequence"`
}

type castKey struct {
	guid    string
	spellID int64
}

type castRow struct{ CastRow }

// ExchangeRow is one interrupt or dispel relationship.
type ExchangeRow struct {
	Kind           string `json:"kind"`
	SourceGUID     string `json:"source_guid"`
	SourceName     string `json:"source_name"`
	TargetGUID     string `json:"target_guid"`
	TargetName     string `json:"target_name"`
	SpellID        int64  `json:"spell_id"`
	SpellName      string `json:"spell_name"`
	ExtraSpellID   int64  `json:"extra_spell_id"`
	ExtraSpellName string `json:"extra_spell_name"`
	Count          int64  `json:"count"`
}

type exchangeKey struct {
	kind             string
	source, target   string
	spellID, extraID int64
}

// ResourceTrack is one actor's power over the fight.
type ResourceTrack struct {
	GUID      string  `json:"guid"`
	Name      string  `json:"name"`
	PowerType int64   `json:"power_type"`
	Series    []int64 `json:"series"`
	Gained    int64   `json:"gained"`
	Spent     int64   `json:"spent"`
	ZeroMS    int64   `json:"zero_ms"`
	// Max is the largest maximum the log reported for this power: the cap the
	// graph draws a line at. Zero when no line ever carried one.
	Max int64 `json:"max"`
	// AtMaxMS is the whole seconds the reading sat at Max, times 1000, on the
	// same buckets Series uses: the time a rage bar or an energy bar was full
	// and everything poured into it was poured away.
	AtMaxMS int64 `json:"at_max_ms"`
	// Wasted is the power the client says was gained past the cap, summed over
	// this track's energize lines.
	Wasted int64 `json:"wasted"`
}

type resourceKey struct {
	guid      string
	powerType int64
}

type resourceTrack struct {
	ResourceTrack
	lastSeen time.Time
	lastVal  int64
	haveLast bool
}

// addDeaths records deaths, the last damage and heals before each one, and
// the auras the player was holding and had just lost.
//
// A melee swing arrives twice: SWING_DAMAGE, whose advanced block describes
// the attacker, and SWING_DAMAGE_LANDED, whose block describes the target
// and so carries the health the hit left. Counting both doubled every
// melee hit in the recap; the landed line now only fills in the health on
// the swing it repeats.
func (a *Accumulator) addDeaths(e event.Event) {
	if e.Kind == event.DamageLanded && e.Dest.GUID != "" {
		// A swing a shield ate in full lands for nothing, and the client writes the
		// target's block at full health on that line whatever they were on: not a
		// reading to keep. The recap leaves that hit's health blank.
		if e.Adv.OK && e.Adv.InfoGUID == e.Dest.GUID && e.Amount.V > 0 {
			q := a.recent[e.Dest.GUID]
			for i := len(q) - 1; i >= 0; i-- {
				if q[i].SourceGUID == e.Source.GUID && q[i].Amount == e.Amount.V && q[i].SpellID == e.Spell.ID {
					if q[i].HPAfter == 0 && q[i].MaxHP == 0 {
						q[i].HPAfter, q[i].MaxHP = e.Adv.CurrentHP, e.Adv.MaxHP
					}
					break
				}
			}
		}
		return
	}
	if e.Kind == event.Heal && e.Dest.GUID != "" {
		ref := HealRef{
			AtMS:       a.ms(e.Time),
			SourceGUID: e.Source.GUID,
			SourceName: a.name(e.Source.GUID),
			SpellID:    e.Spell.ID,
			SpellName:  e.Spell.Name,
			Amount:     e.Amount.V,
			Overheal:   max(e.Overheal.V, 0),
			Absorbed:   e.Absorbed.V,
			HPAfter:    e.Adv.CurrentHP,
			MaxHP:      e.Adv.MaxHP,
		}
		q := append(a.recentHeals[e.Dest.GUID], ref)
		if len(q) > a.opt.DeathHealWindow {
			q = q[len(q)-a.opt.DeathHealWindow:]
		}
		a.recentHeals[e.Dest.GUID] = q
	}
	// A hit a shield ate in full is a miss of type ABSORB carrying the amount on one
	// client; on the recap it is a hit that landed for nothing, with what was absorbed,
	// the same as a swing that landed for 0 on another client.
	soaked := e.Kind == event.Missed && e.MissType == "ABSORB" && e.Amount.OK
	if (e.Kind == event.Damage || soaked) && e.Dest.GUID != "" {
		ref := DamageRef{
			AtMS:       a.ms(e.Time),
			SourceGUID: e.Source.GUID,
			SourceName: a.name(e.Source.GUID),
			SpellID:    e.Spell.ID,
			SpellName:  e.Spell.Name,
			Amount:     e.Amount.V,
			Overkill:   max(e.Overkill.V, 0),
			Absorbed:   e.Absorbed.V,
		}
		if soaked {
			ref.Amount, ref.Absorbed = 0, e.Amount.V
		}
		if e.Adv.OK && e.Adv.InfoGUID == e.Dest.GUID {
			ref.HPAfter, ref.MaxHP = e.Adv.CurrentHP, e.Adv.MaxHP
		}
		q := append(a.recent[e.Dest.GUID], ref)
		if len(q) > a.opt.DeathWindow {
			q = q[len(q)-a.opt.DeathWindow:]
		}
		a.recent[e.Dest.GUID] = q
	}

	if e.Kind != event.Death || units.Parse(e.Dest.GUID).Kind != units.KindPlayer {
		return
	}
	d := &death{at: e.Time}
	d.GUID, d.Name, d.AtMS = e.Dest.GUID, a.name(e.Dest.GUID), a.ms(e.Time)
	if a.opt.Registry != nil {
		if u, ok := a.opt.Registry.Get(e.Dest.GUID); ok {
			d.Class = u.Class
		}
	}
	d.Last = append([]DamageRef(nil), a.recent[e.Dest.GUID]...)
	// Heals are kept only from the first kept hit onward: the ten hits may span
	// a minute and the ten heals an earlier one, and a heal list that starts
	// before the damage list reads as healing that never came.
	heals := a.recentHeals[e.Dest.GUID]
	if len(d.Last) > 0 {
		for len(heals) > 0 && heals[0].AtMS < d.Last[0].AtMS {
			heals = heals[1:]
		}
	}
	d.Heals = append([]HealRef(nil), heals...)
	// The killing blow is the last hit that overkilled; a death the log never
	// showed a lethal hit for (the hit landed after UNIT_DIED was written, or
	// was not logged at all) falls back to the last hit, which the view can
	// tell apart by its missing overkill.
	if n := len(d.Last); n > 0 {
		kb := d.Last[n-1]
		for i := n - 1; i >= 0; i-- {
			if d.Last[i].Overkill > 0 {
				kb = d.Last[i]
				break
			}
		}
		d.KillingBlow = &kb
	}
	cutoff := e.Time.Add(-a.opt.DeathAuraWindow)
	for key, tr := range a.auras {
		if key.target != e.Dest.GUID {
			continue
		}
		if tr.open {
			d.AurasHeld = append(d.AurasHeld, AuraRef{
				SpellID: key.spellID, Name: tr.Name, SourceGUID: tr.openSource,
				Type: tr.Type, AtMS: a.ms(tr.openAt), Stacks: tr.openStacks,
			})
			continue
		}
		if n := len(tr.Segments); n > 0 {
			last := tr.Segments[n-1]
			if a.start.Add(time.Duration(last.EndMS) * time.Millisecond).After(cutoff) {
				d.AurasLost = append(d.AurasLost, AuraRef{
					SpellID: key.spellID, Name: tr.Name, SourceGUID: last.SourceGUID,
					Type: tr.Type, AtMS: last.EndMS, Stacks: last.Stacks,
				})
			}
		}
	}
	sortAuraRefs(d.AurasHeld)
	sortAuraRefs(d.AurasLost)
	if d.AurasHeld == nil {
		d.AurasHeld = []AuraRef{}
	}
	if d.AurasLost == nil {
		d.AurasLost = []AuraRef{}
	}
	if d.Last == nil {
		d.Last = []DamageRef{}
	}
	if d.Heals == nil {
		d.Heals = []HealRef{}
	}
	a.deaths = append(a.deaths, d)
	a.recent[e.Dest.GUID] = nil
	a.recentHeals[e.Dest.GUID] = nil
}

func sortAuraRefs(rs []AuraRef) {
	sort.Slice(rs, func(i, j int) bool {
		if rs[i].AtMS != rs[j].AtMS {
			return rs[i].AtMS < rs[j].AtMS
		}
		return rs[i].SpellID < rs[j].SpellID
	})
}

// addAuras opens and closes aura segments and tracks stacks.
func (a *Accumulator) addAuras(e event.Event) {
	switch e.Kind {
	case event.AuraApplied, event.AuraRefresh, event.AuraDose,
		event.AuraRemoved, event.AuraBroken:
	default:
		return
	}
	if e.Dest.GUID == "" || e.Spell.ID == 0 {
		return
	}
	key := auraKey{target: e.Dest.GUID, spellID: e.Spell.ID}
	tr, ok := a.auras[key]
	if !ok {
		tr = &auraTrack{appliers: map[string]bool{}}
		tr.TargetGUID, tr.TargetName = e.Dest.GUID, a.name(e.Dest.GUID)
		tr.SpellID, tr.Name, tr.Type = e.Spell.ID, e.Spell.Name, e.AuraType
		tr.School = e.Spell.School
		a.auras[key] = tr
	}
	// A track seeded from the combatant snapshot knows its spell only by id
	// until an event names it.
	if tr.Name == "" && e.Spell.Name != "" {
		tr.Name = e.Spell.Name
		tr.School = e.Spell.School
		if e.AuraType != "" {
			tr.Type = e.AuraType
		}
	}
	if e.Source.GUID != "" && e.Source.GUID != units.NoGUID {
		tr.appliers[e.Source.GUID] = true
	}
	switch e.Kind {
	case event.AuraApplied, event.AuraRefresh:
		if !tr.open {
			tr.open, tr.openAt, tr.openStacks = true, e.Time, 1
			tr.openSource = e.Source.GUID
			tr.Applications++
		}
		if tr.openStacks > tr.MaxStacks {
			tr.MaxStacks = tr.openStacks
		}
	case event.AuraDose:
		if !tr.open {
			tr.open, tr.openAt = true, e.Time
			tr.openSource = e.Source.GUID
			tr.Applications++
		}
		tr.closeSegment(a, e.Time)
		tr.open, tr.openAt = true, e.Time
		tr.openStacks = e.Stacks.V
		tr.openSource = e.Source.GUID
		if tr.openStacks > tr.MaxStacks {
			tr.MaxStacks = tr.openStacks
		}
	case event.AuraRemoved, event.AuraBroken:
		tr.closeSegment(a, e.Time)
		tr.open = false
	}
}

func (t *auraTrack) closeSegment(a *Accumulator, at time.Time) {
	if !t.open {
		return
	}
	start, end := a.ms(t.openAt), a.ms(at)
	if end < start {
		end = start
	}
	stacks := t.openStacks
	if stacks == 0 {
		stacks = 1
	}
	t.Segments = append(t.Segments, Segment{
		StartMS: start, EndMS: end, Stacks: stacks, SourceGUID: t.openSource,
	})
	t.UptimeMS += end - start
}

// addCastsAndExchanges records casts, cast times from start-success pairs,
// and the interrupt and dispel relationships.
func (a *Accumulator) addCastsAndExchanges(e event.Event) {
	switch e.Kind {
	case event.CastStart:
		row := a.cast(e)
		row.Started++
		a.pending[castKey{guid: e.Source.GUID, spellID: e.Spell.ID}] = e.Time
	case event.CastSuccess:
		row := a.cast(e)
		row.Succeeded++
		row.Sequence = append(row.Sequence, a.ms(e.Time))
		k := castKey{guid: e.Source.GUID, spellID: e.Spell.ID}
		if started, ok := a.pending[k]; ok {
			if d := e.Time.Sub(started); d > 0 {
				row.CastTimeMS += d.Milliseconds()
			}
			delete(a.pending, k)
		}
		a.markActive(a.owner(e.Source.GUID), e.Time)
	case event.CastFailed:
		row := a.cast(e)
		row.Failed++
		if row.FailReasons == nil {
			row.FailReasons = map[string]int64{}
		}
		row.FailReasons[e.FailedType]++
		delete(a.pending, castKey{guid: e.Source.GUID, spellID: e.Spell.ID})
	case event.Interrupt:
		a.exchange("interrupt", e)
	case event.Dispel:
		a.exchange("dispel", e)
	}
}

func (a *Accumulator) cast(e event.Event) *castRow {
	k := castKey{guid: e.Source.GUID, spellID: e.Spell.ID}
	row, ok := a.casts[k]
	if !ok {
		row = &castRow{}
		row.GUID, row.Name = e.Source.GUID, a.name(e.Source.GUID)
		row.OwnerGUID = a.owner(e.Source.GUID)
		row.SpellID, row.SpellName = e.Spell.ID, e.Spell.Name
		a.casts[k] = row
	}
	return row
}

func (a *Accumulator) exchange(kind string, e event.Event) {
	k := exchangeKey{kind: kind, source: e.Source.GUID, target: e.Dest.GUID,
		spellID: e.Spell.ID, extraID: e.ExtraSpell.ID}
	row, ok := a.exchanges[k]
	if !ok {
		row = &ExchangeRow{
			Kind:       kind,
			SourceGUID: e.Source.GUID, SourceName: a.name(e.Source.GUID),
			TargetGUID: e.Dest.GUID, TargetName: a.name(e.Dest.GUID),
			SpellID: e.Spell.ID, SpellName: e.Spell.Name,
			ExtraSpellID: e.ExtraSpell.ID, ExtraSpellName: e.ExtraSpell.Name,
		}
		a.exchanges[k] = row
	}
	row.Count++
}

// addResources builds the power series from the advanced block and from
// energize events, and counts the time an actor spent at zero power.
func (a *Accumulator) addResources(e event.Event) {
	// Guard the GUID the track is keyed on, not the other one: an energize
	// with no destination would otherwise open a resource row keyed by the
	// empty string and show up in resourceRows as a real actor.
	if e.Kind == event.Energize && e.Dest.GUID != "" && e.Dest.GUID != units.NoGUID {
		k := resourceKey{guid: e.Dest.GUID, powerType: e.PowerType.V}
		tr := a.resource(k, e.Time)
		tr.Gained += e.Amount.V
		// What the client says was gained past the cap. Negative would be a
		// malformed line, and a negative waste is not a thing to report.
		tr.Wasted += max(e.OverEnergize.V, 0)
		if e.MaxPower.V > tr.Max {
			tr.Max = e.MaxPower.V
		}
	}
	if !e.Adv.OK || e.Adv.InfoGUID == "" || e.Adv.InfoGUID == units.NoGUID {
		return
	}
	k := resourceKey{guid: e.Adv.InfoGUID, powerType: e.Adv.PowerType}
	tr := a.resource(k, e.Time)
	if e.Adv.MaxPower > tr.Max {
		tr.Max = e.Adv.MaxPower
	}
	if tr.haveLast {
		if drop := tr.lastVal - e.Adv.CurrentPower; drop > 0 {
			tr.Spent += drop
		}
		if tr.lastVal == 0 && e.Adv.CurrentPower == 0 {
			if gap := e.Time.Sub(tr.lastSeen); gap > 0 {
				tr.ZeroMS += gap.Milliseconds()
			}
		}
	}
	// Buckets with no reading carry the last one forward: the log only
	// reports power on the actor's own events, and a zero in every quiet
	// second drew a mana bar that sat on the floor while ZeroMS said
	// "never empty".
	carry := int64(0)
	if tr.haveLast {
		carry = tr.lastVal
	}
	tr.lastVal, tr.lastSeen, tr.haveLast = e.Adv.CurrentPower, e.Time, true
	b := a.bucket(e.Time)
	for len(tr.Series) <= b {
		tr.Series = append(tr.Series, carry)
	}
	tr.Series[b] = e.Adv.CurrentPower
	if e.Adv.PowerCost > 0 {
		tr.Spent += e.Adv.PowerCost
	}
}

func (a *Accumulator) resource(k resourceKey, at time.Time) *resourceTrack {
	tr, ok := a.resources[k]
	if !ok {
		tr = &resourceTrack{lastSeen: at}
		tr.GUID, tr.Name, tr.PowerType = k.guid, a.name(k.guid), k.powerType
		a.resources[k] = tr
	}
	return tr
}

func (a *Accumulator) deathRows() []Death {
	out := make([]Death, 0, len(a.deaths))
	for _, d := range a.deaths {
		row := d.Death
		row.Last = copySlice(d.Last)
		row.Heals = copySlice(d.Heals)
		row.AurasHeld = a.nameAuraRefs(copySlice(d.AurasHeld))
		row.AurasLost = a.nameAuraRefs(copySlice(d.AurasLost))
		if kb := d.KillingBlow; kb != nil {
			blow := *kb
			row.KillingBlow = &blow
		}
		out = append(out, row)
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].AtMS != out[j].AtMS {
			return out[i].AtMS < out[j].AtMS
		}
		return out[i].GUID < out[j].GUID
	})
	return out
}

// nameAuraRefs names a reference to an aura seeded from a combatant
// snapshot the way the Buffs tab names it: by any line that mentioned the
// spell, and "Spell #id" only if none ever did.
func (a *Accumulator) nameAuraRefs(refs []AuraRef) []AuraRef {
	for i := range refs {
		if refs[i].Name == "" {
			refs[i].Name = a.spellNames[refs[i].SpellID]
		}
		if refs[i].Name == "" {
			refs[i].Name = fmt.Sprintf("Spell #%d", refs[i].SpellID)
		}
	}
	return refs
}

func (a *Accumulator) auraRows() []AuraTrack {
	out := make([]AuraTrack, 0, len(a.auras))
	for _, tr := range a.auras {
		row := tr.AuraTrack
		if row.Name == "" {
			row.Name = a.spellNames[row.SpellID]
		}
		if row.Name == "" {
			row.Name = fmt.Sprintf("Spell #%d", row.SpellID)
		}
		if tr.open {
			// The fight ended with the aura still up: close it at the end
			// so uptime is not silently short.
			end := a.ms(a.end)
			start := a.ms(tr.openAt)
			if end < start {
				end = start
			}
			stacks := tr.openStacks
			if stacks == 0 {
				stacks = 1
			}
			row.Segments = append(copySlice(row.Segments),
				Segment{StartMS: start, EndMS: end, Stacks: stacks, SourceGUID: tr.openSource})
			row.UptimeMS += end - start
		} else {
			row.Segments = copySlice(row.Segments)
		}
		row.Appliers = make([]string, 0, len(tr.appliers))
		for g := range tr.appliers {
			row.Appliers = append(row.Appliers, g)
		}
		sort.Strings(row.Appliers)
		if row.Segments == nil {
			row.Segments = []Segment{}
		}
		out = append(out, row)
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].TargetGUID != out[j].TargetGUID {
			return out[i].TargetGUID < out[j].TargetGUID
		}
		return out[i].SpellID < out[j].SpellID
	})
	return out
}

// castRows emits one row per owner, pet name and spell.
//
// The owner is read here rather than when the row was opened. A guardian
// can cast before anything has told the registry whose it is: the shaman's
// Greater Fire Elemental starts its first Fire Blast on the same tick it is
// summoned, and that summon names a unit the log has so far flagged as a
// neutral NPC, so it is not yet ownable. By the end of the fight the
// advanced block on its own cast lines has named the owner, and a summary
// is written then, so the late reading is the right one. Resolved at the
// first cast, the elemental owned itself, and its rows fell off the Casts
// tab entirely: no player's scope claimed them and it is offered as no
// source of its own.
//
// Rows are then folded by owner, caster name and spell. A re-summoned pet
// is a new GUID every time, so a healer who dropped three Jade Serpent
// Statues read "Soothing Mist · via Jade Serpent Statue" on three separate
// rows of 15, 2 and 2 and had to add them up by hand. One pet by that name,
// one row: the counts sum, the cast times sum, the fail reasons merge and
// the sequence is put back in order. A player's own rows fold into
// themselves and are unchanged, since the owner is their own GUID and their
// name is their own.
func (a *Accumulator) castRows() []CastRow {
	out := make([]CastRow, 0, len(a.casts))
	for _, r := range a.casts {
		row := r.CastRow
		row.OwnerGUID = a.owner(row.GUID)
		row.Sequence = copySlice(r.Sequence)
		row.FailReasons = copyMap(r.FailReasons)
		if row.Sequence == nil {
			row.Sequence = []int64{}
		}
		out = append(out, row)
	}
	// Sorted before the fold so the row that survives a merge, and so its
	// GUID, is the same one on every run over the same log.
	sort.Slice(out, func(i, j int) bool {
		if out[i].GUID != out[j].GUID {
			return out[i].GUID < out[j].GUID
		}
		return out[i].SpellID < out[j].SpellID
	})
	type ownedKey struct {
		owner   string
		name    string
		spellID int64
	}
	folded := make([]CastRow, 0, len(out))
	at := map[ownedKey]int{}
	for _, row := range out {
		k := ownedKey{owner: row.OwnerGUID, name: row.Name, spellID: row.SpellID}
		i, ok := at[k]
		if !ok {
			at[k] = len(folded)
			folded = append(folded, row)
			continue
		}
		into := &folded[i]
		into.Started += row.Started
		into.Succeeded += row.Succeeded
		into.Failed += row.Failed
		into.CastTimeMS += row.CastTimeMS
		into.Sequence = append(into.Sequence, row.Sequence...)
		slices.Sort(into.Sequence)
		for reason, n := range row.FailReasons {
			if into.FailReasons == nil {
				into.FailReasons = map[string]int64{}
			}
			into.FailReasons[reason] += n
		}
	}
	sort.Slice(folded, func(i, j int) bool {
		if folded[i].GUID != folded[j].GUID {
			return folded[i].GUID < folded[j].GUID
		}
		return folded[i].SpellID < folded[j].SpellID
	})
	return folded
}

func (a *Accumulator) exchangeRows(kind string) []ExchangeRow {
	var out []ExchangeRow
	for k, r := range a.exchanges {
		if k.kind == kind {
			out = append(out, *r)
		}
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].Count != out[j].Count {
			return out[i].Count > out[j].Count
		}
		if out[i].SourceGUID != out[j].SourceGUID {
			return out[i].SourceGUID < out[j].SourceGUID
		}
		if out[i].ExtraSpellID != out[j].ExtraSpellID {
			return out[i].ExtraSpellID < out[j].ExtraSpellID
		}
		if out[i].TargetGUID != out[j].TargetGUID {
			return out[i].TargetGUID < out[j].TargetGUID
		}
		return out[i].SpellID < out[j].SpellID
	})
	if out == nil {
		return []ExchangeRow{}
	}
	return out
}

func (a *Accumulator) resourceRows() []ResourceTrack {
	out := make([]ResourceTrack, 0, len(a.resources))
	for _, tr := range a.resources {
		row := tr.ResourceTrack
		row.Series = copySlice(tr.Series)
		if row.Series == nil {
			row.Series = []int64{}
		}
		row.AtMaxMS = atMaxMS(row.Series, row.Max, a.opt.Bucket)
		out = append(out, row)
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].GUID != out[j].GUID {
			return out[i].GUID < out[j].GUID
		}
		return out[i].PowerType < out[j].PowerType
	})
	return out
}

// atMaxMS is the time the series sat at the cap: one bucket per second whose
// reading is the maximum. Read off the finished series rather than counted as
// the events arrive, so it means what the drawn line means -- a second with no
// reading carries the last one forward, and a bar that was full through a quiet
// stretch was full. A track the log never gave a maximum for has no cap to be at.
func atMaxMS(series []int64, maximum int64, bucket time.Duration) int64 {
	if maximum <= 0 {
		return 0
	}
	var total int64
	for _, value := range series {
		if value >= maximum {
			total += bucket.Milliseconds()
		}
	}
	return total
}
