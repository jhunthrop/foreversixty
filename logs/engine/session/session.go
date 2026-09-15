// logs/engine/session/session.go
// Package session is the engine's public surface. One session is fed a
// byte stream in any chunking, hands back decoded events and closed fights
// as they complete, answers Snapshot for the fight in progress, and can be
// serialised and resumed from its own state and byte offset.
package session

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
	"github.com/jhunthrop/foreversixty/logs/engine/fight"
	"github.com/jhunthrop/foreversixty/logs/engine/layout"
	"github.com/jhunthrop/foreversixty/logs/engine/lexer"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
	"github.com/jhunthrop/foreversixty/logs/engine/units"
)

// Version is the engine version. It travels with every summary, every
// metrics row, and every report, so a report always says which code
// produced it. Bump it whenever decoded output changes.
const Version = "0.2.0"

// Options configures a session.
type Options struct {
	// ReportID labels the metrics rows. Required for ranking output.
	ReportID string
	// Layout forces a dialect. Leave it zero to select one from the
	// header, and set Infer to fall back to counting.
	Layout layout.Layout
	// Infer allows the counting fallback: layout.Infer builds a row from
	// the log's own field counts when no registered row matches the
	// header, and when there is no header at all. With Infer off both
	// cases fall back to layout.RetailV16.
	Infer bool
	// Base seeds the clock for dialects whose timestamps carry no year.
	// Pass the log file's modification time for a batch parse and the
	// current time for a live tail.
	Base time.Time
	// KeepEvents keeps each closed fight's events in memory so they can be
	// written as Parquet. Off for a pure tail, on for anything that
	// publishes.
	KeepEvents bool
	Units      units.Options
	Fight      fight.Options
	Summary    summary.Options
}

// Closed is one fight that finished during a Feed.
type Closed struct {
	Fight   fight.Fight
	Summary summary.Summary
	Metrics []summary.MetricRow
	Events  []event.Event
}

// Result is what one Feed produced.
type Result struct {
	Events []event.Event
	Closed []Closed
	Bytes  int64
}

// Health is the per-report health the spec puts in report.json.
type Health struct {
	EngineVersion   string         `json:"engine_version"`
	Layout          string         `json:"layout"`
	LayoutVerified  bool           `json:"layout_verified"`
	LayoutInferred  bool           `json:"layout_inferred"`
	AdvancedLogging bool           `json:"advanced_logging"`
	MissingHeader   bool           `json:"missing_header"`
	Lines           int64          `json:"lines"`
	ParseErrors     int64          `json:"parse_errors"`
	UnknownEvents   map[string]int `json:"unknown_events"`
	DroppedLines    int            `json:"dropped_lines"`
	ClockJumps      int            `json:"clock_jumps"`
	YearRollovers   int            `json:"year_rollovers"`
	HeaderRestarts  int            `json:"header_restarts"`
}

// cloneHealth deep-copies h's map, so a snapshot taken at one point in time
// is unaffected by counters that keep changing afterwards.
func cloneHealth(h Health) Health {
	c := h
	c.UnknownEvents = make(map[string]int, len(h.UnknownEvents))
	for k, v := range h.UnknownEvents {
		c.UnknownEvents[k] = v
	}
	return c
}

// Session is the engine, wired together.
type Session struct {
	opt Options

	lex  *lexer.Lexer
	dec  *event.Decoder
	reg  *units.Registry
	seg  *fight.Segmenter
	acc  *summary.Accumulator
	open *fight.Fight

	kept       []event.Event
	pending    []lexer.Line // buffered while a layout is still unknown
	inferring  bool
	headerSeen bool

	// openClock and openSeen are the decoder's clock and Seen() as they
	// stood immediately before the line that opened the fight currently in
	// progress was decoded. State() serialises them as the clock a restored
	// session must resume with, since the caller replays that same line.
	openClock time.Time
	openSeen  bool
	// openHealth is health as it stood immediately before that same line,
	// so State() can serialise a baseline that excludes every line the
	// caller is about to replay. Without it, Lines and friends would be
	// counted twice: once by this session before it was serialised, once
	// again by the restored session replaying the open fight's byte range.
	openHealth Health

	health Health
}

// New starts a session.
func New(o Options) *Session {
	s := &Session{opt: o, lex: lexer.New()}
	s.reg = units.NewRegistry(o.Units)
	s.seg = fight.NewSegmenter(o.Fight)
	s.health.EngineVersion = Version
	s.health.UnknownEvents = map[string]int{}

	lay := o.Layout
	if lay.Name == "" {
		s.inferring = true
	} else {
		s.setLayout(lay)
	}
	return s
}

func (s *Session) setLayout(l layout.Layout) {
	s.health.Layout = l.Name
	s.health.LayoutVerified = l.Verified
	s.health.LayoutInferred = l.Name == "inferred"
	s.health.AdvancedLogging = l.Advanced > 0
	if s.dec == nil {
		s.dec = event.NewDecoder(l, s.opt.Base)
	} else {
		s.dec.SetLayout(l)
	}
	s.inferring = false
}

// inferWindow is how many lines to buffer before giving up on finding a
// header and inferring a layout from the sample.
const inferWindow = 2000

// Feed consumes one chunk. Overlapping and duplicate byte ranges are
// ignored, so at-least-once delivery is safe; a chunk that would leave a
// gap returns an error and nothing is consumed.
func (s *Session) Feed(chunk []byte, offset int64) (Result, error) {
	var res Result
	before := s.lex.NextOffset()
	err := s.lex.Feed(chunk, offset, func(ln lexer.Line) error {
		s.line(ln, &res)
		return nil
	})
	if err != nil {
		return res, err
	}
	res.Bytes = s.lex.NextOffset() - before
	return res, nil
}

func (s *Session) line(ln lexer.Line, res *Result) {
	if s.inferring {
		h, isHeader := layout.ParseHeader(ln)
		if !isHeader {
			s.buffer(ln, res)
			return
		}
		switch l, found := layout.Lookup(h); {
		case found:
			s.setLayout(l)
		case s.opt.Infer:
			// No row for this header: buffer and let Infer count fields.
			s.buffer(ln, res)
			return
		default:
			s.setLayout(layout.RetailV16())
		}
		// Anything buffered before the header is logged before it.
		s.replay(res)
	}
	s.handle(ln, res)
}

// buffer appends ln to the inference sample and flushes it the moment the
// window fills, checked per line rather than once per Feed call, so the
// sample layout.Infer sees is the same regardless of how the caller chunks
// the input.
func (s *Session) buffer(ln lexer.Line, res *Result) {
	s.pending = append(s.pending, ln)
	if len(s.pending) >= inferWindow {
		s.flushInference(res)
	}
}

// flushInference settles the layout for a log that presented no header at
// all and replays the lines buffered while waiting for one. Options.Infer
// decides what settles it: the counting fallback, or retail v16, which is
// the same fallback line() takes for a header no row recognises. Either
// way Health.MissingHeader records that there was no header.
func (s *Session) flushInference(res *Result) {
	s.health.MissingHeader = true
	if s.opt.Infer {
		s.setLayout(layout.Infer(s.pending))
	} else {
		s.setLayout(layout.RetailV16())
	}
	s.replay(res)
}

func (s *Session) replay(res *Result) {
	pending := s.pending
	s.pending = nil
	for _, ln := range pending {
		s.handle(ln, res)
	}
}

func (s *Session) handle(ln lexer.Line, res *Result) {
	s.health.Lines++
	prevClock, prevSeen := s.dec.Time(), s.dec.Seen()
	e := s.dec.Decode(ln)

	if e.Kind == event.Header && s.headerSeen {
		// A second header means the logger restarted: a hard boundary.
		s.health.HeaderRestarts++
		s.closeOpen(e.Time, res)
		s.resetFight()
		s.reg = units.NewRegistry(s.opt.Units)
		if h, ok := layout.ParseHeader(ln); ok {
			if l, found := layout.Lookup(h); found {
				s.setLayout(l)
			}
		}
	}

	if e.Kind == event.Header {
		s.headerSeen = true
	}

	switch e.Kind {
	case event.ParseError:
		s.health.ParseErrors++
	case event.Unknown:
		s.health.UnknownEvents[e.Name]++
	}

	s.reg.Observe(e)
	res.Events = append(res.Events, e)

	step := s.seg.Feed(e)
	if step.Closed != nil {
		s.finish(*step.Closed, res)
	}
	if step.Discarded {
		// The segmenter dropped a trash segment too short to report, so
		// nothing calls finish for it. Let go of its state here, or
		// s.acc != nil would stop meaning "a fight is open".
		s.resetFight()
	}
	if step.Fight == nil {
		return
	}
	if step.Opened {
		s.openClock, s.openSeen = prevClock, prevSeen
		// The opening line can only be hostile combat or ENCOUNTER_START
		// (see hostileCombat in package fight), so it is never counted as
		// a ParseError, an Unknown event, or a header restart; only Lines
		// needs undoing to land on the count as of just before this line.
		s.openHealth = cloneHealth(s.health)
		s.openHealth.Lines--
		s.startFight(step.Fight)
	}
	if s.acc != nil {
		s.acc.Add(e)
		if s.opt.KeepEvents {
			s.kept = append(s.kept, e)
		}
	}
}

func (s *Session) startFight(f *fight.Fight) {
	o := s.opt.Summary
	o.Registry = s.reg
	s.acc = summary.New(o)
	s.acc.Start(f.Start)
	s.kept = nil
	s.open = f
}

func (s *Session) finish(f fight.Fight, res *Result) {
	if s.acc == nil {
		return
	}
	sum := s.acc.Snapshot(f, Version)
	c := Closed{Fight: f, Summary: sum, Metrics: s.acc.Metrics(s.opt.ReportID, f, sum, Version)}
	if s.opt.KeepEvents {
		c.Events = s.kept
	}
	res.Closed = append(res.Closed, c)
	s.resetFight()
}

// resetFight drops the state that belongs to one fight. Every path that
// ends a fight goes through it, so s.acc != nil means exactly "a fight is
// open".
func (s *Session) resetFight() {
	s.acc, s.kept, s.open = nil, nil, nil
}

func (s *Session) closeOpen(at time.Time, res *Result) {
	if f := s.seg.Flush(at); f != nil {
		s.finish(*f, res)
	} else {
		s.resetFight()
	}
}

// Snapshot is the running summary of the fight in progress. The second
// return is false when no fight is open.
func (s *Session) Snapshot() (fight.Fight, summary.Summary, bool) {
	open := s.seg.Open()
	if open == nil || s.acc == nil {
		return fight.Fight{}, summary.Summary{}, false
	}
	return *open, s.acc.Snapshot(*open, Version), true
}

// Close flushes the trailing partial line and the fight still open.
func (s *Session) Close() (Result, error) {
	var res Result
	if err := s.lex.Flush(func(ln lexer.Line) error {
		s.line(ln, &res)
		return nil
	}); err != nil {
		return res, err
	}
	if s.inferring {
		s.flushInference(&res)
	}
	s.closeOpen(s.dec.Time(), &res)
	s.health.DroppedLines = s.lex.Dropped()
	return res, nil
}

// Health reports the per-report health.
func (s *Session) Health() Health {
	h := s.health
	h.DroppedLines = s.lex.Dropped()
	if s.dec != nil {
		h.ClockJumps = s.dec.ClockJumps()
		h.YearRollovers = s.dec.Rollovers()
	}
	h.UnknownEvents = map[string]int{}
	for k, v := range s.health.UnknownEvents {
		h.UnknownEvents[k] = v
	}
	return h
}

// Offset is the stream offset of the next byte the session expects.
func (s *Session) Offset() int64 { return s.lex.NextOffset() }

// Units is the registry, for report.json's unit list.
func (s *Session) Units() *units.Registry { return s.reg }

// state is the serialised form of a session. Decoded events and the open
// fight's accumulator are not in it: the caller replays the open fight's
// byte range, which is bounded by one fight.
//
// Layout carries the full row, not just LayoutName: a registered row is
// looked up afresh by name on Restore, but an inferred row has no entry in
// layout.Rows() to look up, so its dynamically derived shape must travel
// in the state itself. layout.Layout's fields are all exported and its
// maps are sorted by encoding/json, so this stays deterministic.
type state struct {
	Version    string        `json:"version"`
	LayoutName string        `json:"layout_name"`
	Layout     layout.Layout `json:"layout"`
	Lexer      lexer.State   `json:"lexer"`
	Units      units.State   `json:"units"`
	Fight      fight.State   `json:"fight"`
	Clock      time.Time     `json:"clock"`
	ClockSeen  bool          `json:"clock_seen"`
	Health     Health        `json:"health"`
	HeaderSeen bool          `json:"header_seen"`
	Replay     int64         `json:"replay_offset"`
}

// State serialises the session. Restore resumes from it, and the caller
// must re-feed from the returned replay offset so the open fight is rebuilt.
func (s *Session) State() ([]byte, error) {
	if s.inferring {
		return nil, fmt.Errorf("session: no layout has been settled yet, so there is " +
			"nothing resumable: the stream has presented no header and fewer than the " +
			"lines inference needs. Feed more input or Close the session first")
	}
	st := state{
		Version:    Version,
		Lexer:      s.lex.State(),
		Units:      s.reg.State(),
		Fight:      s.seg.State(),
		Health:     s.health,
		HeaderSeen: s.headerSeen,
		Replay:     s.lex.NextOffset(),
	}
	if s.dec != nil {
		st.LayoutName = s.dec.Layout().Name
		st.Layout = s.dec.Layout()
		st.Clock = s.dec.Time()
		st.ClockSeen = s.dec.Seen()
	}
	if open := s.seg.Open(); open != nil {
		st.Replay = open.StartOffset
		st.Fight.Open = nil // the open fight is rebuilt by replaying
		st.Lexer = lexer.State{Offset: open.StartOffset, Number: open.StartLine - 1}
		// The clock, and the health counters, must resume as of just
		// before the line that opened this fight, since the caller
		// replays every line from there on: replaying already-counted
		// lines a second time must not count them twice.
		st.Clock = s.openClock
		st.ClockSeen = s.openSeen
		st.Health = s.openHealth
	}
	b, err := json.Marshal(st)
	if err != nil {
		return nil, fmt.Errorf("session: marshal state: %w", err)
	}
	return b, nil
}

// ReplayOffset reads the byte offset a restored session must be fed from.
func ReplayOffset(b []byte) (int64, error) {
	var st state
	if err := json.Unmarshal(b, &st); err != nil {
		return 0, fmt.Errorf("session: unmarshal state: %w", err)
	}
	return st.Replay, nil
}

// Restore rebuilds a session. Feed it from ReplayOffset(b).
func Restore(o Options, b []byte) (*Session, error) {
	var st state
	if err := json.Unmarshal(b, &st); err != nil {
		return nil, fmt.Errorf("session: unmarshal state: %w", err)
	}
	if st.Version != Version {
		return nil, fmt.Errorf("session: state was written by engine %q, this is %q", st.Version, Version)
	}
	s := &Session{opt: o}
	s.lex = lexer.Restore(st.Lexer)
	s.reg = units.RestoreRegistry(o.Units, st.Units)
	s.seg = fight.RestoreSegmenter(o.Fight, st.Fight)
	s.health = st.Health
	s.headerSeen = st.HeaderSeen
	if s.health.UnknownEvents == nil {
		s.health.UnknownEvents = map[string]int{}
	}
	lay, ok := rowByName(st.LayoutName)
	if !ok {
		// Not a registered row: fall back to the layout serialised in the
		// state itself, which is how an inferred layout survives a
		// restore, since it has no entry in layout.Rows() to look up.
		switch {
		case st.Layout.Name != "":
			lay = st.Layout
		case o.Layout.Name != "":
			lay = o.Layout
		default:
			return nil, fmt.Errorf("session: state names layout %q, which is not registered", st.LayoutName)
		}
	}
	s.setLayout(lay)
	s.dec.SetTime(st.Clock)
	s.dec.SetSeen(st.ClockSeen)
	return s, nil
}

func rowByName(name string) (layout.Layout, bool) {
	for _, r := range layout.Rows() {
		if r.Name == name {
			return r, true
		}
	}
	return layout.Layout{}, false
}
