// logs/engine/event/decode.go
package event

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/logs/engine/layout"
	"github.com/jhunthrop/foreversixty/logs/engine/lexer"
)

// Decoder turns lexed lines into Events using one layout row. It is not safe
// for concurrent use; one decoder belongs to one session.
type Decoder struct {
	lay       layout.Layout
	prev      time.Time
	seen      bool
	rollovers int
	jumps     int
}

// NewDecoder returns a decoder for one dialect. base seeds the clock: a
// dialect whose timestamps carry no year takes the year from it, so pass
// the log file's modification time for a batch parse and the current time
// for a live tail. A zero base makes every line of a yearless dialect a
// parse error, which is the honest failure rather than a guessed year.
func NewDecoder(l layout.Layout, base time.Time) *Decoder {
	return &Decoder{lay: l, prev: base}
}

// Layout returns the row in use.
func (d *Decoder) Layout() layout.Layout { return d.lay }

// SetLayout swaps the row, for a header that appears mid-file.
func (d *Decoder) SetLayout(l layout.Layout) { d.lay = l }

// Time is the timestamp of the last line decoded.
func (d *Decoder) Time() time.Time { return d.prev }

// SetTime seeds the clock, for Restore.
func (d *Decoder) SetTime(t time.Time) { d.prev = t }

// Seen reports whether at least one timestamped line has been decoded.
func (d *Decoder) Seen() bool { return d.seen }

// SetSeen seeds whether a timestamped line has already been decoded, for
// Restore: a session resuming mid-stream has seen lines before the byte
// range it replays, so the clock-jump check must not treat the first
// replayed line as the stream's first.
func (d *Decoder) SetSeen(seen bool) { d.seen = seen }

// Rollovers counts year rollovers applied so far.
func (d *Decoder) Rollovers() int { return d.rollovers }

// ClockJumps counts lines whose timestamp went backwards.
func (d *Decoder) ClockJumps() int { return d.jumps }

// Decode never fails: a line it cannot read becomes a ParseError event
// carrying the raw text, so a malformed line costs one event, not the file.
func (d *Decoder) Decode(ln lexer.Line) Event {
	e := Event{Offset: ln.Offset, Line: ln.Number, Time: d.prev}
	if len(ln.Params) == 0 || ln.Params[0] == "" {
		e.Kind, e.Raw, e.Error = ParseError, ln.Raw, "empty line"
		return e
	}
	e.Name = ln.Params[0]

	if ln.Stamp != "" {
		t, rolled, err := d.lay.ParseStamp(ln.Stamp, d.prev)
		if err != nil {
			e.Kind, e.Raw, e.Error = ParseError, ln.Raw, err.Error()
			return e
		}
		if rolled {
			d.rollovers++
		}
		if d.seen && t.Before(d.prev) {
			d.jumps++
		}
		d.seen = true
		d.prev, e.Time = t, t
	}

	if h, ok := layout.ParseHeader(ln); ok {
		e.Kind, e.Raw = Header, ln.Raw
		e.Amount = OptInt{V: int64(h.Version), OK: true}
		e.Total = OptInt{V: int64(h.ProjectID), OK: true}
		e.Critical = OptBool{V: h.Advanced, OK: true}
		e.ItemName = h.Build
		return e
	}

	if spec, ok := d.lay.Specials[e.Name]; ok {
		if len(spec.Widths) == 0 {
			// The row lists the event but not its shape. There is nothing
			// to index against, so the line is kept raw and the
			// conformance report picks it up as an unknown event.
			e.Kind, e.Raw = Unknown, ln.Raw
			return e
		}
		if !spec.Accepts(len(ln.Params)) {
			return fail(e, ln, fmt.Sprintf("%s has %d fields, layout %q allows %v",
				e.Name, len(ln.Params), d.lay.Name, spec.Widths))
		}
		return d.decodeSpecial(e, ln)
	}

	prefix, suffix, known := d.lay.Split(e.Name)
	if !known {
		e.Kind, e.Raw = Unknown, ln.Raw
		if len(ln.Params) >= layout.BaseParams {
			readUnits(&e, ln.Params)
		}
		return e
	}
	return d.decodeStandard(e, ln, prefix, suffix)
}

func fail(e Event, ln lexer.Line, msg string) Event {
	e.Kind, e.Raw, e.Error = ParseError, ln.Raw, msg
	return e
}

// readUnits fills the common header from p[1..8].
func readUnits(e *Event, p []string) {
	e.Source = Unit{GUID: p[1], Name: nilless(p[2]), Flags: hex32(p[3]), Raid: hex32(p[4])}
	e.Dest = Unit{GUID: p[5], Name: nilless(p[6]), Flags: hex32(p[7]), Raid: hex32(p[8])}
}

func (d *Decoder) decodeStandard(e Event, ln lexer.Line, prefix, suffix string) Event {
	p := ln.Params
	spec := d.lay.Suffixes[suffix]
	want, advAt := d.lay.Width(prefix, suffix)

	// _MISSED carries three more fields, but only on an absorb; _DAMAGE on
	// the Classic row carries an optional trailing isOffHand.
	switch {
	case spec.AbsorbExtra > 0 && len(p) == want+spec.AbsorbExtra:
		want += spec.AbsorbExtra
	case spec.OffHand && len(p) == want+1:
		want++
	}
	// An aura event may carry a trailing absorb size.
	if strings.HasPrefix(suffix, "_AURA_") && !strings.HasSuffix(suffix, "_DOSE") &&
		suffix != "_AURA_BROKEN_SPELL" && len(p) == want+1 {
		want++
	}
	if len(p) != want {
		return fail(e, ln, fmt.Sprintf("%s has %d fields, layout %q wants %d", e.Name, len(p), d.lay.Name, want))
	}

	// Where the suffix's fields start. The inferred row derives Params
	// from the file itself and can hand back a negative count, so this
	// index and the suffix branch's own requirement are both checked
	// before anything is sliced: len(p) == want above only proves the
	// line matches the row's arithmetic, not that the arithmetic is sane.
	i := layout.BaseParams
	if n := d.lay.Prefixes[prefix]; n == 3 {
		i += 3
	}
	if advAt >= 0 {
		i = advAt + d.lay.Advanced
	}
	if advAt < 0 || advAt > i {
		advAt = -1
	}
	if need := suffixNeeds(suffix, spec); i < layout.BaseParams || i > len(p) || len(p)-i < need {
		return fail(e, ln, fmt.Sprintf("%s has %d fields, layout %q leaves %d for a %s that reads %d",
			e.Name, len(p), d.lay.Name, len(p)-i, suffix, need))
	}

	readUnits(&e, p)
	if d.lay.Prefixes[prefix] == 3 {
		e.Spell = Spell{ID: intOf(p[layout.BaseParams]), Name: nilless(p[layout.BaseParams+1]), School: intOf(p[layout.BaseParams+2])}
	}
	if advAt >= 0 {
		e.Adv = readAdvanced(p[advAt : advAt+d.lay.Advanced])
	}
	rest := p[i:]

	switch suffix {
	case "_DAMAGE", "_DAMAGE_LANDED":
		e.Kind = Damage
		readDamage(&e, rest, spec)
	case "_HEAL":
		e.Kind = Heal
		readHeal(&e, rest, spec)
	case "_MISSED":
		e.Kind = Missed
		e.MissType = rest[0]
		e.OffHand = boolOf(rest[1])
		if len(rest) >= 5 {
			e.Amount, e.BaseAmount, e.Critical = optInt(rest[2]), optInt(rest[3]), boolOf(rest[4])
		} else if len(rest) == 4 {
			e.Amount, e.Critical = optInt(rest[2]), boolOf(rest[3])
		}
	case "_ENERGIZE", "_DRAIN", "_LEECH":
		e.Kind = Energize
		e.Amount, e.OverEnergize = optInt(rest[0]), optInt(rest[1])
		e.PowerType, e.MaxPower = optInt(rest[2]), optInt(rest[3])
	case "_AURA_APPLIED", "_AURA_REMOVED", "_AURA_REFRESH", "_AURA_BROKEN":
		e.Kind = map[string]Kind{
			"_AURA_APPLIED": AuraApplied, "_AURA_REMOVED": AuraRemoved,
			"_AURA_REFRESH": AuraRefresh, "_AURA_BROKEN": AuraBroken,
		}[suffix]
		e.AuraType = rest[0]
		if len(rest) > 1 {
			e.Absorbed = optInt(rest[1])
		}
	case "_AURA_APPLIED_DOSE", "_AURA_REMOVED_DOSE":
		e.Kind = AuraDose
		e.AuraType, e.Stacks = rest[0], optInt(rest[1])
	case "_AURA_BROKEN_SPELL":
		e.Kind = AuraBroken
		e.ExtraSpell = Spell{ID: intOf(rest[0]), Name: nilless(rest[1]), School: intOf(rest[2])}
		e.AuraType = rest[3]
	case "_INTERRUPT":
		e.Kind = Interrupt
		e.ExtraSpell = Spell{ID: intOf(rest[0]), Name: nilless(rest[1]), School: intOf(rest[2])}
	case "_DISPEL", "_STOLEN":
		e.Kind = Dispel
		e.ExtraSpell = Spell{ID: intOf(rest[0]), Name: nilless(rest[1]), School: intOf(rest[2])}
		e.AuraType = rest[3]
	case "_DISPEL_FAILED":
		e.Kind = Dispel
		e.ExtraSpell = Spell{ID: intOf(rest[0]), Name: nilless(rest[1]), School: intOf(rest[2])}
	case "_CAST_START":
		e.Kind = CastStart
	case "_CAST_SUCCESS":
		e.Kind = CastSuccess
	case "_CAST_FAILED":
		e.Kind = CastFailed
		e.FailedType = rest[0]
	case "_SUMMON":
		e.Kind = Summon
	case "_CREATE":
		e.Kind = Create
	case "_RESURRECT":
		e.Kind = Resurrect
	case "_INSTAKILL":
		e.Kind = Instakill
	case "_EXTRA_ATTACKS":
		e.Kind = ExtraAttacks
		e.Amount = optInt(rest[0])
	case "_DURABILITY_DAMAGE":
		e.Kind = Durability
	default:
		e.Kind, e.Raw = Unknown, ln.Raw
	}
	return e
}

// suffixNeeds is how many fields after the prefix and the advanced block
// the switch in decodeStandard reads unconditionally. The layout row says
// how wide the line should be; this says what the code actually indexes,
// and the two are checked separately because an inferred row's widths come
// from the file rather than from a document.
func suffixNeeds(suffix string, spec layout.Suffix) int {
	switch suffix {
	case "_DAMAGE", "_DAMAGE_LANDED":
		// amount, then overkill through crushing; plus baseAmount when
		// the row carries it. The trailing isOffHand is optional and is
		// read only when it is there.
		if spec.BaseAmount {
			return 10
		}
		return 9
	case "_HEAL":
		if spec.HealedToHP {
			return 5
		}
		return 4
	case "_MISSED":
		return 2
	case "_ENERGIZE", "_DRAIN", "_LEECH":
		return 4
	case "_AURA_APPLIED", "_AURA_REMOVED", "_AURA_REFRESH", "_AURA_BROKEN":
		return 1
	case "_AURA_APPLIED_DOSE", "_AURA_REMOVED_DOSE":
		return 2
	case "_AURA_BROKEN_SPELL", "_DISPEL", "_STOLEN":
		return 4
	case "_INTERRUPT", "_DISPEL_FAILED":
		return 3
	case "_CAST_FAILED", "_EXTRA_ATTACKS":
		return 1
	default:
		return 0
	}
}

// readDamage reads the damage suffix. Amount is always the first field and
// is always the damage that landed: it matched the target's HP drop in
// 47,278 events of the retail sample, and the base amount never did.
func readDamage(e *Event, rest []string, spec layout.Suffix) {
	i := 0
	e.Amount = optInt(rest[i])
	i++
	if spec.BaseAmount {
		e.BaseAmount = optInt(rest[i])
		i++
	}
	e.Overkill = optInt(rest[i])
	e.School = optInt(rest[i+1])
	e.Resisted = optInt(rest[i+2])
	e.Blocked = optInt(rest[i+3])
	e.Absorbed = optInt(rest[i+4])
	e.Critical = boolOf(rest[i+5])
	e.Glancing = boolOf(rest[i+6])
	e.Crushing = boolOf(rest[i+7])
	if i+8 < len(rest) {
		e.OffHand = boolOf(rest[i+8])
	}
}

// readHeal reads the heal suffix. Amount is the canonical heal: it includes
// overheal and the part diverted into a shield, and healedToHP + absorbed
// == amount held in 39,592 of 39,599 heals in the retail sample.
func readHeal(e *Event, rest []string, spec layout.Suffix) {
	i := 0
	if spec.HealedToHP {
		e.Total = optInt(rest[0]) // healedToHP
		i = 1
	}
	e.Amount = optInt(rest[i])
	e.Overheal = optInt(rest[i+1])
	e.Absorbed = optInt(rest[i+2])
	e.Critical = boolOf(rest[i+3])
}

func readAdvanced(f []string) Advanced {
	if len(f) < 17 {
		return Advanced{}
	}
	return Advanced{
		OK:           true,
		InfoGUID:     f[0],
		OwnerGUID:    f[1],
		CurrentHP:    intOf(f[2]),
		MaxHP:        intOf(f[3]),
		AttackPower:  intOf(f[4]),
		SpellPower:   intOf(f[5]),
		Armor:        intOf(f[6]),
		Absorb:       intOf(f[7]),
		PowerType:    intOf(f[8]),
		CurrentPower: intOf(f[9]),
		MaxPower:     intOf(f[10]),
		PowerCost:    intOf(f[11]),
		PositionX:    floatOf(f[12]),
		PositionY:    floatOf(f[13]),
		UIMapID:      intOf(f[14]),
		Facing:       floatOf(f[15]),
		Level:        intOf(f[16]),
	}
}

// nilless turns the game's "nil" placeholder into an empty string.
func nilless(s string) string {
	if s == "nil" {
		return ""
	}
	return s
}

// optInt parses a number that may be hex ("0x20"), decimal ("32"), signed
// ("-1"), or a decimal fraction ("50.0000", which energize events use).
func optInt(s string) OptInt {
	if s == "" || s == "nil" {
		return OptInt{}
	}
	if len(s) > 2 && (s[:2] == "0x" || s[:2] == "0X") {
		if v, err := strconv.ParseUint(s[2:], 16, 64); err == nil {
			return OptInt{V: int64(v), OK: true}
		}
		return OptInt{}
	}
	if v, err := strconv.ParseInt(s, 10, 64); err == nil {
		return OptInt{V: v, OK: true}
	}
	if f, err := strconv.ParseFloat(s, 64); err == nil {
		return OptInt{V: int64(f), OK: true}
	}
	return OptInt{}
}

func intOf(s string) int64 { return optInt(s).V }

func floatOf(s string) float64 {
	if s == "" || s == "nil" {
		return 0
	}
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return f
}

// boolOf reads the game's boolean: "nil" is present and false, "1" is true,
// an empty field is absent.
func boolOf(s string) OptBool {
	switch s {
	case "":
		return OptBool{}
	case "nil":
		return OptBool{OK: true}
	default:
		return OptBool{V: s != "0", OK: true}
	}
}

func hex32(s string) uint32 {
	v := optInt(s)
	if !v.OK {
		return 0
	}
	return uint32(v.V)
}
