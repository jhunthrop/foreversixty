// logs/engine/fight/fight.go
// Package fight splits a log into fights. ENCOUNTER_START names a fight;
// everything else is segmented by gaps in hostile combat and labelled
// "Trash" with the zone, which is what makes dungeon logs with no encounter
// events work the same way as raid logs.
package fight

import (
	"sort"
	"time"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
	"github.com/jhunthrop/foreversixty/logs/engine/units"
)

// Kind separates named encounters from everything else.
type Kind string

// The kinds. The strings are stable: they appear in report.json.
const (
	Encounter Kind = "encounter"
	Trash     Kind = "trash"
)

// Marker is a raid target marker seen on a unit during the fight.
type Marker struct {
	GUID string    `json:"guid"`
	Name string    `json:"name"`
	Flag uint32    `json:"flag"`
	Time time.Time `json:"time"`
}

// Fight is one segment of the log.
type Fight struct {
	Index       int       `json:"index"`
	Kind        Kind      `json:"kind"`
	EncounterID int64     `json:"encounter_id,omitempty"`
	Name        string    `json:"name"`
	Difficulty  int64     `json:"difficulty,omitempty"`
	Size        int64     `json:"size,omitempty"`
	Kill        bool      `json:"kill"`
	InProgress  bool      `json:"in_progress"`
	ZoneID      int64     `json:"zone_id,omitempty"`
	Zone        string    `json:"zone,omitempty"`
	Start       time.Time `json:"start"`
	End         time.Time `json:"end"`
	StartLine   int64     `json:"start_line"`
	EndLine     int64     `json:"end_line"`
	StartOffset int64     `json:"start_offset"`
	EndOffset   int64     `json:"end_offset"`
	Players     []string  `json:"players"`
	Markers     []Marker  `json:"markers,omitempty"`
	NPCKills    int       `json:"npc_kills"`
	Deaths      int       `json:"deaths"`

	players map[string]bool
	markers map[string]Marker
}

// Duration is the fight's wall length.
func (f *Fight) Duration() time.Duration {
	if f.End.Before(f.Start) {
		return 0
	}
	return f.End.Sub(f.Start)
}

// Options tunes segmentation.
type Options struct {
	// Gap closes a trash fight when hostile combat has been quiet this long.
	Gap time.Duration
	// MinTrash drops a trash segment shorter than this: a single stray
	// swing in a city is not a fight.
	MinTrash time.Duration
	// Trailing keeps an encounter open after ENCOUNTER_END so that the
	// damage-over-time ticks that land afterwards are counted.
	Trailing time.Duration
}

// DefaultOptions are the values the CLI and the session use. A zero Gap
// takes the default; a zero MinTrash or Trailing disables that behaviour on
// purpose, so pass a negative value to ask for the default instead.
func DefaultOptions() Options {
	return Options{Gap: 5 * time.Second, MinTrash: 3 * time.Second, Trailing: 2 * time.Second}
}

func (o Options) withDefaults() Options {
	d := DefaultOptions()
	if o.Gap <= 0 {
		o.Gap = d.Gap
	}
	if o.MinTrash < 0 {
		o.MinTrash = d.MinTrash
	}
	if o.Trailing < 0 {
		o.Trailing = d.Trailing
	}
	return o
}

// Step is the segmenter's verdict on one event.
type Step struct {
	// Closed is a fight that ended before this event was assigned. It is
	// nil when nothing closed, and nil for a trash segment too short to
	// report.
	Closed *Fight
	// Fight is the fight this event belongs to, or nil when the event fell
	// outside any fight.
	Fight *Fight
	// Opened is true when Fight was opened by this event.
	Opened bool
}

// State is the segmenter in a form that survives a restart.
type State struct {
	Next      int       `json:"next"`
	Open      *Fight    `json:"open,omitempty"`
	Zone      string    `json:"zone,omitempty"`
	ZoneID    int64     `json:"zone_id,omitempty"`
	LastFight time.Time `json:"last_fight,omitempty"`
	Ending    bool      `json:"ending,omitempty"`
	EndAt     time.Time `json:"end_at,omitempty"`
}

// Segmenter turns a stream of events into fights.
type Segmenter struct {
	opt       Options
	next      int
	open      *Fight
	zone      string
	zoneID    int64
	lastFight time.Time
	ending    bool
	endAt     time.Time
}

// NewSegmenter returns a segmenter. Zero fields in o take their defaults.
func NewSegmenter(o Options) *Segmenter {
	return &Segmenter{opt: o.withDefaults(), next: 1}
}

// RestoreSegmenter rebuilds a segmenter from saved state.
func RestoreSegmenter(o Options, s State) *Segmenter {
	sg := NewSegmenter(o)
	sg.next, sg.zone, sg.zoneID = s.Next, s.Zone, s.ZoneID
	sg.lastFight, sg.ending, sg.endAt = s.LastFight, s.Ending, s.EndAt
	if s.Open != nil {
		f := *s.Open
		f.players = map[string]bool{}
		for _, g := range f.Players {
			f.players[g] = true
		}
		f.markers = map[string]Marker{}
		for _, m := range f.Markers {
			f.markers[m.GUID] = m
		}
		sg.open = &f
	}
	if sg.next == 0 {
		sg.next = 1
	}
	return sg
}

// State captures the segmenter for serialisation.
func (s *Segmenter) State() State {
	st := State{Next: s.next, Zone: s.zone, ZoneID: s.zoneID,
		LastFight: s.lastFight, Ending: s.ending, EndAt: s.endAt}
	if s.open != nil {
		f := *s.open
		f.Players, f.Markers = s.open.snapshotPlayers(), s.open.snapshotMarkers()
		st.Open = &f
	}
	return st
}

// Open returns the fight in progress, or nil.
func (s *Segmenter) Open() *Fight {
	if s.open == nil {
		return nil
	}
	f := *s.open
	f.Players, f.Markers = s.open.snapshotPlayers(), s.open.snapshotMarkers()
	f.InProgress = true
	return &f
}

// Feed assigns one event.
func (s *Segmenter) Feed(e event.Event) Step {
	var step Step

	if e.Kind == event.ZoneChange && e.Zone != nil {
		s.zone, s.zoneID = e.Zone.Name, e.Zone.ID
	}

	// An encounter boundary always wins over gap segmentation.
	if e.Kind == event.EncounterStart && e.Encounter != nil {
		step.Closed = s.close(e.Time, e.Line, e.Offset)
		s.open = s.newFight(Encounter, e)
		s.open.EncounterID = e.Encounter.ID
		s.open.Name = e.Encounter.Name
		s.open.Difficulty = e.Encounter.Difficulty
		s.open.Size = e.Encounter.Size
		step.Fight, step.Opened = s.open, true
		s.assign(e)
		return step
	}
	if e.Kind == event.EncounterEnd && e.Encounter != nil && s.open != nil && s.open.Kind == Encounter {
		s.open.Kill = e.Encounter.Kill
		s.assign(e)
		step.Fight = s.open
		if s.opt.Trailing > 0 {
			s.ending, s.endAt = true, e.Time.Add(s.opt.Trailing)
		} else {
			step.Closed = s.close(e.Time, e.Line, e.Offset)
		}
		return step
	}

	combat := hostileCombat(e)

	switch {
	case s.ending && !e.Time.Before(s.endAt):
		// The encounter's trailing window has run out.
		step.Closed = s.close(s.endAt, e.Line, e.Offset)
	case s.open != nil && s.open.Kind == Trash && e.Time.Sub(s.lastFight) > s.opt.Gap:
		// Hostile combat has been quiet for longer than the gap.
		step.Closed = s.close(s.lastFight, e.Line, e.Offset)
	}

	if s.open == nil {
		if !combat {
			return step
		}
		s.open = s.newFight(Trash, e)
		s.open.Name = "Trash"
		step.Opened = true
	}
	if combat {
		s.lastFight = e.Time
	}
	s.assign(e)
	step.Fight = s.open
	return step
}

// Flush closes the fight still open at the end of the stream.
func (s *Segmenter) Flush(at time.Time) *Fight {
	if s.open == nil {
		return nil
	}
	end := at
	if s.open.Kind == Trash && !s.lastFight.IsZero() && s.lastFight.Before(at) {
		end = s.lastFight
	}
	if s.ending && s.endAt.Before(end) {
		end = s.endAt
	}
	return s.close(end, s.open.EndLine, s.open.EndOffset)
}

func (s *Segmenter) newFight(k Kind, e event.Event) *Fight {
	return &Fight{
		Index: s.next, Kind: k,
		ZoneID: s.zoneID, Zone: s.zone,
		Start: e.Time, End: e.Time,
		StartLine: e.Line, EndLine: e.Line,
		StartOffset: e.Offset, EndOffset: e.Offset,
		InProgress: true,
		players:    map[string]bool{},
		markers:    map[string]Marker{},
	}
}

func (s *Segmenter) assign(e event.Event) {
	f := s.open
	if f == nil {
		return
	}
	if e.Time.After(f.End) {
		f.End = e.Time
	}
	f.EndLine, f.EndOffset = e.Line, e.Offset
	for _, u := range [...]event.Unit{e.Source, e.Dest} {
		if u.GUID == "" || u.GUID == units.NoGUID {
			continue
		}
		if units.Parse(u.GUID).Kind == units.KindPlayer {
			f.players[u.GUID] = true
		}
		if u.Raid != 0 {
			if _, seen := f.markers[u.GUID]; !seen {
				f.markers[u.GUID] = Marker{GUID: u.GUID, Name: u.Name, Flag: u.Raid, Time: e.Time}
			}
		}
	}
	switch e.Kind {
	case event.Death:
		if units.Parse(e.Dest.GUID).Kind == units.KindPlayer {
			f.Deaths++
		} else if units.Hostile(e.Dest.Flags) {
			f.NPCKills++
		}
	case event.PartyKill:
		f.NPCKills++
	}
}

func (s *Segmenter) close(end time.Time, line, offset int64) *Fight {
	f := s.open
	s.open, s.ending = nil, false
	if f == nil {
		return nil
	}
	if end.After(f.End) {
		f.End = end
	}
	if line > f.EndLine {
		f.EndLine, f.EndOffset = line, offset
	}
	f.InProgress = false
	f.Players, f.Markers = f.snapshotPlayers(), f.snapshotMarkers()
	if f.Kind == Trash && f.Duration() < s.opt.MinTrash {
		return nil // too short to be a fight; the index is not consumed
	}
	f.Index = s.next
	s.next++
	return f
}

func (f *Fight) snapshotPlayers() []string {
	out := make([]string, 0, len(f.players))
	for g := range f.players {
		out = append(out, g)
	}
	sort.Strings(out)
	return out
}

func (f *Fight) snapshotMarkers() []Marker {
	out := make([]Marker, 0, len(f.markers))
	for _, m := range f.markers {
		out = append(out, m)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].GUID < out[j].GUID })
	if len(out) == 0 {
		return nil
	}
	return out
}

// hostileCombat reports whether an event is combat between opposed sides,
// which is what starts and sustains a trash fight. Buffs cast in a city and
// a player eating do not.
func hostileCombat(e event.Event) bool {
	switch e.Kind {
	case event.Damage, event.Missed, event.Absorbed, event.PartyKill:
	default:
		return false
	}
	src, dst := e.Source.Flags, e.Dest.Flags
	if src == 0 && dst == 0 {
		return false
	}
	return units.Hostile(src) != units.Hostile(dst)
}
