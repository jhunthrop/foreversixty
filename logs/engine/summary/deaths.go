// logs/engine/summary/deaths.go
package summary

import (
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
	AurasHeld   []AuraRef   `json:"auras_held"`
	AurasLost   []AuraRef   `json:"auras_lost"`
	ReleaseMS   int64       `json:"release_ms,omitempty"`
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
	TargetGUID   string    `json:"target_guid"`
	TargetName   string    `json:"target_name"`
	SpellID      int64     `json:"spell_id"`
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
	GUID        string           `json:"guid"`
	Name        string           `json:"name"`
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

// addDeaths records deaths, the last damage before each one, and the auras
// the player was holding and had just lost.
func (a *Accumulator) addDeaths(e event.Event) {
	if e.Kind == event.Damage && e.Dest.GUID != "" {
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
	if n := len(d.Last); n > 0 {
		kb := d.Last[n-1]
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
	a.deaths = append(a.deaths, d)
	a.recent[e.Dest.GUID] = nil
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
		a.auras[key] = tr
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
	if e.Kind == event.Energize && e.Source.GUID != "" {
		k := resourceKey{guid: e.Dest.GUID, powerType: e.PowerType.V}
		a.resource(k, e.Time).Gained += e.Amount.V
	}
	if !e.Adv.OK || e.Adv.InfoGUID == "" || e.Adv.InfoGUID == units.NoGUID {
		return
	}
	k := resourceKey{guid: e.Adv.InfoGUID, powerType: e.Adv.PowerType}
	tr := a.resource(k, e.Time)
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
	tr.lastVal, tr.lastSeen, tr.haveLast = e.Adv.CurrentPower, e.Time, true
	b := a.bucket(e.Time)
	for len(tr.Series) <= b {
		tr.Series = append(tr.Series, 0)
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
		out = append(out, d.Death)
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].AtMS != out[j].AtMS {
			return out[i].AtMS < out[j].AtMS
		}
		return out[i].GUID < out[j].GUID
	})
	return out
}

func (a *Accumulator) auraRows() []AuraTrack {
	out := make([]AuraTrack, 0, len(a.auras))
	for _, tr := range a.auras {
		row := tr.AuraTrack
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
			row.Segments = append(append([]Segment(nil), row.Segments...),
				Segment{StartMS: start, EndMS: end, Stacks: stacks, SourceGUID: tr.openSource})
			row.UptimeMS += end - start
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

func (a *Accumulator) castRows() []CastRow {
	out := make([]CastRow, 0, len(a.casts))
	for _, r := range a.casts {
		row := r.CastRow
		if row.Sequence == nil {
			row.Sequence = []int64{}
		}
		out = append(out, row)
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].GUID != out[j].GUID {
			return out[i].GUID < out[j].GUID
		}
		return out[i].SpellID < out[j].SpellID
	})
	return out
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
		return out[i].ExtraSpellID < out[j].ExtraSpellID
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
		if row.Series == nil {
			row.Series = []int64{}
		}
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
