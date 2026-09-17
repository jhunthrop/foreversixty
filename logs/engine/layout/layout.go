// logs/engine/layout/layout.go
// Package layout holds one row per combat-log dialect. A row says how many
// parameters each event carries, where the advanced block sits, and how the
// timestamp is written, so the decoder never guesses. When Forever's first
// beta log arrives the work is one new row here.
package layout

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/logs/engine/lexer"
)

// rolloverDays is how far a yearless timestamp may fall behind the
// previous line before it is read as the next year rather than as a clock
// that went backwards. Half a year is the only threshold that cannot be
// wrong in either direction: a log spanning more than six months does not
// exist, and a New Year's Eve raid crosses the boundary by hours.
const rolloverDays = 180

// BaseParams is the common header shared by every event with a source and a
// target: event, sourceGUID, sourceName, sourceFlags, sourceRaidFlags,
// destGUID, destName, destFlags, destRaidFlags.
const BaseParams = 9

// Suffix describes the parameters that follow the prefix.
type Suffix struct {
	Params      int  // parameters after the prefix and after the advanced block
	Advanced    bool // the event carries the advanced block
	AbsorbExtra int  // extra parameters present only when missType == "ABSORB"
	OffHand     bool // a trailing isOffHand that may be absent
	BaseAmount  bool // the damage suffix carries the unmodified base amount at index 1
	HealedToHP  bool // the heal suffix carries healedToHP at index 0
	// Tag marks a suffix whose line may end in the single-target / area
	// tag, the two-valued field combat-log version 22 writes as "ST" or
	// "AOE". It is optional because the same suffix is written with the
	// tag under a SPELL prefix and without it under SWING, and because
	// RANGE writes it on _DAMAGE but not on _MISSED. No prefix rule fits
	// all four cases, so the decoder recognises the tag by its value; see
	// isScopeTag in the event package.
	Tag bool
	// MissAmount marks a _MISSED suffix that carries one extra field, the
	// amount missed, when the miss type is BLOCK or RESIST. ABSORB's three
	// extras are counted by AbsorbExtra and are unaffected.
	MissAmount bool
	// AuraExtra marks an aura suffix that may carry a second trailing
	// number after the absorb size. Version 22 writes two where version 16
	// writes one; the second one's meaning is not pinned, so it is
	// counted and not read.
	AuraExtra bool
}

// Special describes an event that does not follow the prefix/suffix pattern.
// Widths are total field counts including the event name.
//
// An empty Widths means the row does not know this event's shape, and the
// decoder then keeps the line raw rather than indexing into it. Absence of
// information is ignorance, not permission: a row that does not state a
// width has no business telling the decoder where a field sits.
type Special struct {
	Widths []int
}

// Accepts reports whether n is one of the widths this special allows. A
// special with no declared widths accepts nothing.
func (s Special) Accepts(n int) bool {
	for _, w := range s.Widths {
		if w == n {
			return true
		}
	}
	return false
}

// Combatant locates the parts of a COMBATANT_INFO line.
type Combatant struct {
	Present        bool
	Params         int
	SpecIndex      int
	TalentIndex    int
	PvPTalentIndex int
	// BorrowIndex is the borrowed-power field, or 0 when the dialect
	// writes none. Zero is unambiguous: field 0 is always the event name.
	BorrowIndex int
	GearIndex   int
	AuraIndex   int
	// StatIndex maps a stat name to the field that holds it. The names are
	// the keys the decoder publishes in Combatant.Stats. A dialect that
	// does not write a stat leaves it out of the map rather than pointing
	// it at a field that means something else.
	StatIndex map[string]int
	// TalentsAreSpellIDs is true for a dialect whose talent field is the
	// flat spell-id tuple event.Combatant.Talents is typed to hold. A
	// dialect whose talent trees write something else (version 22's
	// (nodeID, entryID, rank) triples) leaves this false, and the decoder
	// emits an empty Talents rather than a flattened, wrong-but-typed mix
	// of the three; see docs/ledger/2026-09-16-retail-v22.md.
	TalentsAreSpellIDs bool
}

// Layout is one dialect.
type Layout struct {
	Name      string
	Version   int
	ProjectID int // 0 matches any project
	Advanced  int // advanced-block field count, 0 when the dialect has none
	StampYear bool
	StampZone bool
	Prefixes  map[string]int // prefix name to extra parameter count
	Suffixes  map[string]Suffix
	Specials  map[string]Special
	Combatant Combatant
	Verified  bool // the row was checked against a real log of this dialect

	// WidthOverrides pins the exact accepted widths for one event (keyed by
	// prefix+suffix, i.e. the event name Split would reassemble them into),
	// replacing what the flag cross-product in widthsFor would otherwise
	// compute. It exists because a suffix's optional fields do not always
	// occur in every combination the flags allow for every prefix that
	// shares it: version 22's single-target/area tag, for instance, is
	// mandatory on SPELL_DAMAGE and absent on SWING_DAMAGE even though both
	// use the "_DAMAGE" suffix and its Tag flag. A dialect with no
	// exceptions leaves this nil.
	WidthOverrides map[string][]int

	// widthCache is the accepted-widths set for every (prefix, suffix) pair
	// this row knows, built once when the row is finalised (by Lookup, or
	// by a constructor that has no header to wait for) so that Widths, on
	// the decoder's per-event path, is a map lookup rather than a slice
	// built and sorted on every call. A row built by hand (a test's literal
	// Layout{}, or layout.Infer's result) leaves it nil, and Widths falls
	// back to computing the answer on the spot.
	widthCache map[string][]int
}

// buildWidthCache computes the accepted-widths set for every (prefix,
// suffix) pair l's Prefixes and Suffixes can combine into. It must be
// called only after every field the computation reads (Advanced,
// Prefixes, Suffixes, WidthOverrides) has its final value: Lookup zeroes
// Advanced for a header that turns advanced logging off, and a cache built
// before that would keep the wrong widths.
func buildWidthCache(l Layout) map[string][]int {
	cache := make(map[string][]int, len(l.Prefixes)*len(l.Suffixes))
	for prefix := range l.Prefixes {
		for suffix := range l.Suffixes {
			cache[prefix+suffix] = widthsFor(l, prefix, suffix)
		}
	}
	return cache
}

// Header is the COMBAT_LOG_VERSION line.
type Header struct {
	Version   int
	Advanced  bool
	Build     string
	ProjectID int
	Fields    int
}

// ParseHeader reads a COMBAT_LOG_VERSION line. The line is a flat sequence of
// key, value pairs, so unknown keys are skipped rather than shifting anything.
func ParseHeader(ln lexer.Line) (Header, bool) {
	if len(ln.Params) == 0 || ln.Params[0] != "COMBAT_LOG_VERSION" {
		return Header{}, false
	}
	h := Header{Fields: len(ln.Params)}
	p := ln.Params
	if len(p) > 1 {
		h.Version, _ = strconv.Atoi(strings.TrimSpace(p[1]))
	}
	for i := 2; i+1 < len(p); i += 2 {
		v := strings.TrimSpace(p[i+1])
		switch strings.TrimSpace(p[i]) {
		case "ADVANCED_LOG_ENABLED":
			h.Advanced = v == "1"
		case "BUILD_VERSION":
			h.Build = v
		case "PROJECT_ID":
			h.ProjectID, _ = strconv.Atoi(v)
		}
	}
	return h, true
}

// rows is the table, most specific first.
var rows = []Layout{RetailV22(), RetailV16(), ClassicWiki()}

// Lookup returns the row for a header, and whether one matched.
func Lookup(h Header) (Layout, bool) {
	for _, r := range rows {
		if r.Version == 0 || r.Version != h.Version {
			continue
		}
		if r.ProjectID != 0 && h.ProjectID != 0 && r.ProjectID != h.ProjectID {
			continue
		}
		l := r
		if !h.Advanced {
			l.Advanced = 0
		}
		l.widthCache = buildWidthCache(l)
		return l, true
	}
	return Layout{}, false
}

// Rows returns every registered row, for the conformance command.
func Rows() []Layout { return append([]Layout(nil), rows...) }

// Split separates an event name into its prefix and suffix. The winner is
// the longest registered prefix whose remainder is a suffix this row knows,
// which is what lets a row register both "SWING" and "SWING_DAMAGE_LANDED"
// without the longer one swallowing the shorter one's events. When no
// prefix leaves a known suffix the longest prefix match is reported with
// ok false, so the caller's error names the closest thing the row knows.
// Events handled by Specials must be checked first.
func (l Layout) Split(event string) (prefix, suffix string, ok bool) {
	best, bestSuffix := "", ""
	for p := range l.Prefixes {
		if len(p) <= len(best) || !strings.HasPrefix(event, p) {
			continue
		}
		if _, known := l.Suffixes[event[len(p):]]; known {
			best, bestSuffix = p, event[len(p):]
		}
	}
	if best != "" {
		return best, bestSuffix, true
	}
	for p := range l.Prefixes {
		if strings.HasPrefix(event, p) && len(p) > len(best) {
			best = p
		}
	}
	if best == "" {
		return "", "", false
	}
	return best, event[len(best):], false
}

// Width returns the total field count an event of this shape must have, and
// the index at which the advanced block starts (-1 when there is none).
func (l Layout) Width(prefix, suffix string) (width, advAt int) {
	pre := l.Prefixes[prefix]
	s := l.Suffixes[suffix]
	advAt = -1
	width = BaseParams + pre
	if s.Advanced && l.Advanced > 0 {
		advAt = width
		width += l.Advanced
	}
	return width + s.Params, advAt
}

// Widths returns every total field count this row accepts for a shape, in
// ascending order. When the row was finalised by Lookup, this is a lookup
// into a cache built once for the whole row; a row built by hand computes
// the answer on the spot, which is fine off the decoder's per-event path.
func (l Layout) Widths(prefix, suffix string) []int {
	if l.widthCache != nil {
		if w, ok := l.widthCache[prefix+suffix]; ok {
			return w
		}
	}
	return widthsFor(l, prefix, suffix)
}

// widthsFor is the computation Widths caches. WidthOverrides, when it names
// this event, is the exact answer: it exists precisely because the flag
// cross-product below is a superset of what the dialect actually writes for
// some (prefix, suffix) pairs. Absent an override, Width gives the base
// count the row's arithmetic produces, and the optional trailing fields a
// dialect may or may not write (an absorb's extras, a Classic isOffHand, an
// aura's absorb size, version 22's single-target tag) turn that one number
// into a small set.
func widthsFor(l Layout, prefix, suffix string) []int {
	if w, ok := l.WidthOverrides[prefix+suffix]; ok {
		return w
	}
	base, _ := l.Width(prefix, suffix)
	s := l.Suffixes[suffix]
	widths := []int{base}
	if s.AbsorbExtra > 0 {
		widths = append(widths, base+s.AbsorbExtra)
	}
	if s.MissAmount {
		widths = append(widths, base+1)
	}
	if s.OffHand {
		widths = append(widths, base+1)
	}
	if AuraCarriesAmount(suffix) {
		widths = append(widths, base+1)
		if s.AuraExtra {
			widths = append(widths, base+2)
		}
	}
	if s.Tag {
		for _, w := range append([]int(nil), widths...) {
			widths = append(widths, w+1)
		}
	}
	sort.Ints(widths)
	out := widths[:0]
	for i, w := range widths {
		if i == 0 || w != widths[i-1] {
			out = append(out, w)
		}
	}
	return out
}

// AuraCarriesAmount reports whether an aura suffix may be followed by the
// size of the absorb the aura provides. The dose suffixes carry a stack
// count in a field of their own and _AURA_BROKEN_SPELL carries the
// breaking spell, so neither takes the optional amount.
func AuraCarriesAmount(suffix string) bool {
	return strings.HasPrefix(suffix, "_AURA_") &&
		!strings.HasSuffix(suffix, "_DOSE") &&
		suffix != "_AURA_BROKEN_SPELL"
}

// ParseStamp turns a timestamp into a time. prev is the time of the previous
// line and is used only to roll the year over on a dialect that does not
// write one; it is never used to invent a value. The returned bool reports
// whether a year rollover was applied.
func (l Layout) ParseStamp(stamp string, prev time.Time) (time.Time, bool, error) {
	s := strings.TrimSpace(stamp)
	if s == "" {
		return time.Time{}, false, fmt.Errorf("layout: empty timestamp")
	}
	date, clock, ok := strings.Cut(s, " ")
	if !ok {
		return time.Time{}, false, fmt.Errorf("layout: timestamp %q has no date and time", s)
	}
	zone := time.UTC
	if l.StampZone {
		var off string
		if i := strings.IndexAny(clock, "+-"); i > 0 {
			off, clock = clock[i:], clock[:i]
			mins, err := parseZone(off)
			if err != nil {
				return time.Time{}, false, err
			}
			zone = time.FixedZone("", mins*60)
		}
	}
	dp := strings.Split(date, "/")
	if (l.StampYear && len(dp) != 3) || (!l.StampYear && len(dp) != 2) {
		return time.Time{}, false, fmt.Errorf("layout: timestamp date %q does not match the row", date)
	}
	month, err := strconv.Atoi(dp[0])
	if err != nil {
		return time.Time{}, false, fmt.Errorf("layout: month in %q: %w", date, err)
	}
	day, err := strconv.Atoi(dp[1])
	if err != nil {
		return time.Time{}, false, fmt.Errorf("layout: day in %q: %w", date, err)
	}
	year := prev.Year()
	if l.StampYear {
		if year, err = strconv.Atoi(dp[2]); err != nil {
			return time.Time{}, false, fmt.Errorf("layout: year in %q: %w", date, err)
		}
	} else if year == 1 {
		return time.Time{}, false, fmt.Errorf("layout: timestamp %q has no year and no previous line to take one from", s)
	}
	h, m, sec, ns, err := parseClock(clock)
	if err != nil {
		return time.Time{}, false, err
	}
	t := time.Date(year, time.Month(month), day, h, m, sec, ns, zone)
	rolled := false
	if !l.StampYear && !prev.IsZero() && t.Before(prev.AddDate(0, 0, -rolloverDays)) {
		t = t.AddDate(1, 0, 0)
		rolled = true
	}
	return t, rolled, nil
}

func parseClock(clock string) (h, m, s, ns int, err error) {
	parts := strings.Split(clock, ":")
	if len(parts) != 3 {
		return 0, 0, 0, 0, fmt.Errorf("layout: timestamp clock %q is not h:m:s", clock)
	}
	if h, err = strconv.Atoi(parts[0]); err != nil {
		return 0, 0, 0, 0, fmt.Errorf("layout: hour in %q: %w", clock, err)
	}
	if m, err = strconv.Atoi(parts[1]); err != nil {
		return 0, 0, 0, 0, fmt.Errorf("layout: minute in %q: %w", clock, err)
	}
	secText, frac, _ := strings.Cut(parts[2], ".")
	if s, err = strconv.Atoi(secText); err != nil {
		return 0, 0, 0, 0, fmt.Errorf("layout: second in %q: %w", clock, err)
	}
	if frac != "" {
		for len(frac) < 9 {
			frac += "0"
		}
		if ns, err = strconv.Atoi(frac[:9]); err != nil {
			return 0, 0, 0, 0, fmt.Errorf("layout: fraction in %q: %w", clock, err)
		}
	}
	return h, m, s, ns, nil
}

func parseZone(off string) (int, error) {
	sign := 1
	if off[0] == '-' {
		sign = -1
	}
	body := off[1:]
	hours, mins := body, "0"
	if h, m, ok := strings.Cut(body, ":"); ok {
		hours, mins = h, m
	}
	h, err := strconv.Atoi(hours)
	if err != nil {
		return 0, fmt.Errorf("layout: zone hours in %q: %w", off, err)
	}
	m, err := strconv.Atoi(mins)
	if err != nil {
		return 0, fmt.Errorf("layout: zone minutes in %q: %w", off, err)
	}
	return sign * (h*60 + m), nil
}

// EventNames lists every event this row can decode, sorted, for the
// conformance report.
func (l Layout) EventNames() []string {
	seen := map[string]bool{}
	for p := range l.Prefixes {
		for s := range l.Suffixes {
			seen[p+s] = true
		}
	}
	for s := range l.Specials {
		seen[s] = true
	}
	out := make([]string, 0, len(seen))
	for n := range seen {
		out = append(out, n)
	}
	sort.Strings(out)
	return out
}
