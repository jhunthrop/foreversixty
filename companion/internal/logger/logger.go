// companion/internal/logger/logger.go
// Package logger works out whose combat log it is. The game marks
// every unit in the log with affiliation flags, and the logging
// player's own units carry COMBATLOG_OBJECT_AFFILIATION_MINE: the
// player-kind unit with that bit is the logger, with no guessing and
// from the first fight. The log names it "Name-Realm-" (the realm with
// its spaces removed, then a trailing hyphen), and a two-word Forever
// name arrives with its words joined by a hyphen too:
// "Bow-Jackzon-ClassicBetaPvP-". The addon's exports, which the
// companion already reads, say which region and ruleset that name
// lives on.
package logger

import (
	"strings"

	"github.com/jhunthrop/foreversixty/companion/internal/addon"
	"github.com/jhunthrop/foreversixty/companion/internal/character"
	"github.com/jhunthrop/foreversixty/logs/engine/units"
)

// Detected is who the combat log says is logging.
type Detected struct {
	// Name is the character's name with its words separated by single
	// spaces ("Bow Jackzon", though the log writes "Bow-Jackzon"),
	// without the realm.
	Name string
	// Realm is the log's realm segment with its spaces removed; empty
	// when the unit carried none.
	Realm string
	// Class is the engine's class for the unit when COMBATANT_INFO
	// supplied it, empty otherwise.
	Class string
}

// Detect finds the logger among the players a fight's registry holds.
// A unit with an owner is a pet or a guardian wearing a player's GUID
// and is never the logger. When a log spans two characters the one
// seen last wins.
func Detect(players []units.Unit) (Detected, bool) {
	var best *units.Unit
	for i := range players {
		u := &players[i]
		if u.Kind != units.KindPlayer || u.OwnerGUID != "" || u.Flags&units.FlagAffiliationMine == 0 {
			continue
		}
		if best == nil || u.LastSeen.After(best.LastSeen) {
			best = u
		}
	}
	if best == nil {
		return Detected{}, false
	}
	name, realm := Split(best.Name)
	if name == "" {
		return Detected{}, false
	}
	d := Detected{Name: name, Realm: realm}
	if best.ClassSource == "combatant_info" {
		d.Class = best.Class
	}
	return d, true
}

// Split separates a log unit name into the character's name and the
// normalised realm. The current client writes "Xuen-ClassicBetaPvP-"
// for a one-word name and "Bow-Jackzon-ClassicBetaPvP-" for a two-word
// one: the words of the name, the realm, a trailing hyphen. The realm
// is the last segment once the trailing hyphen is gone, with its spaces
// removed; the name is every segment before it, joined with single
// spaces, which is how the addon spells it. A unit with no hyphen is
// all name.
func Split(unit string) (name, realm string) {
	unit = strings.TrimRight(strings.TrimSpace(unit), "-")
	parts := strings.Split(unit, "-")
	if len(parts) == 1 {
		return strings.TrimSpace(unit), ""
	}
	words := make([]string, 0, len(parts)-1)
	for _, w := range parts[:len(parts)-1] {
		if w = strings.TrimSpace(w); w != "" {
			words = append(words, w)
		}
	}
	return strings.Join(words, " "), strings.ReplaceAll(parts[len(parts)-1], " ", "")
}

// Record is one character the addon has exported.
type Record struct {
	Name string
	// Realm is what the addon wrote as the second segment: a realm
	// name or a ruleset.
	Realm  string
	Region string
	// Class is the export string's class token, empty when it has none.
	Class string
	// Preferred marks a record from the watched install's account
	// files, which wins a tie.
	Preferred bool
}

// Records turns exports into records. preferred is whether they came
// from the watched install.
func Records(exports []addon.Export, preferred bool) []Record {
	out := make([]Record, 0, len(exports))
	for _, e := range exports {
		out = append(out, Record{Name: e.Name, Realm: e.Ruleset, Region: e.Region,
			Class: exportClass(e.Export), Preferred: preferred})
	}
	return out
}

// exportClass is the class token of "FS1:<build>:<class>:...".
func exportClass(export string) string {
	parts := strings.SplitN(export, ":", 4)
	if len(parts) < 4 || parts[0] != "FS1" {
		return ""
	}
	return parts[2]
}

// Reason is why a log could not be credited to a character.
type Reason int

const (
	// Matched means the log resolved to one character.
	Matched Reason = iota
	// NoRecord means no addon record carries the name: the addon has
	// not exported that character.
	NoRecord
	// NoMatch means records carry the name but none, or more than one,
	// fits the log's class and realm (or has a region to credit).
	NoMatch
)

// Unmatched is the sentence that says what the addon has not sent for
// a log naming name, given why Judge found no character: the Status
// row, the tray and the companion log all repeat it. It is empty for
// Matched.
func (r Reason) Unmatched(name string) string {
	switch r {
	case NoRecord:
		return "the log names " + name + ", but the addon has not exported that character — " +
			"log out once with the addon loaded on it"
	case NoMatch:
		return "the log names " + name + ", but no addon record carries that realm — " +
			"log out once with the addon loaded"
	}
	return ""
}

// Resolve names the character the log is for out of the addon's
// records; ok is false when it cannot. See Judge for the rules.
func Resolve(d Detected, records []Record) (c character.Character, ok bool) {
	c, why := Judge(d, records)
	return c, why == Matched
}

// Judge names the character the log is for out of the addon's
// records, and says why when it cannot, so the pipeline's log line and
// the Status row give the same reason. It matches on the name
// (case-insensitively, spaces kept), then on class where both sides
// have one, then on the realm where that narrows, then on the watched
// install; only exactly one survivor with a region is a character.
// Anything else is no answer, because crediting a report to the wrong
// character is worse than to none.
//
// Forever characters can carry a two-word name ("Bow Jackzon"), which
// the addon records in full from UnitName. The current client's log
// carries the full name too ("Bow-Jackzon-ClassicBetaPvP-", which Split
// reads as Bow Jackzon), but an older log and the WTF folder carry only
// the first word ("Bow-ClassicBetaPvP-", seen on 19 Oct 2026). Both
// shapes must work, so a one-word log name also matches a record whose
// first word it is, while a two-word log name matches only the full
// name and is never matched by its first word alone.
func Judge(d Detected, records []Record) (character.Character, Reason) {
	name := strings.Join(strings.Fields(d.Name), " ")
	oneWord := !strings.ContainsRune(name, ' ')
	var found []Record
	for _, r := range records {
		full := strings.Join(strings.Fields(r.Name), " ")
		if strings.EqualFold(full, name) || (oneWord && strings.EqualFold(firstWord(full), name)) {
			found = append(found, r)
		}
	}
	if len(found) == 0 {
		return character.Character{}, NoRecord
	}
	if d.Class != "" {
		found = keep(found, func(r Record) bool { return r.Class == "" || sameWord(r.Class, d.Class) })
	}
	if d.Realm != "" {
		found = narrow(found, func(r Record) bool { return sameWord(r.Realm, d.Realm) })
	}
	found = narrow(found, func(r Record) bool { return r.Preferred })
	if len(found) != 1 {
		return character.Character{}, NoMatch
	}
	r := found[0]
	realm := r.Realm
	if realm == "" {
		realm = d.Realm
	}
	c := character.Character{
		Region:  strings.ToLower(strings.TrimSpace(r.Region)),
		Ruleset: character.RulesetFromRealm(realm),
		Name:    strings.TrimSpace(r.Name),
	}
	if !c.Valid() {
		return character.Character{}, NoMatch
	}
	return c, Matched
}

func keep(rs []Record, f func(Record) bool) []Record {
	var out []Record
	for _, r := range rs {
		if f(r) {
			out = append(out, r)
		}
	}
	return out
}

// narrow applies f only when something survives it: a tie-breaker
// that matches nobody must not veto the candidates.
func narrow(rs []Record, f func(Record) bool) []Record {
	if len(rs) < 2 {
		return rs
	}
	if kept := keep(rs, f); len(kept) > 0 {
		return kept
	}
	return rs
}

// sameWord compares two spellings of a class or realm: case, spaces
// and punctuation do not count.
func sameWord(a, b string) bool { return flat(a) == flat(b) }

// firstWord is the part of a name before its first space: what the
// combat log and the WTF folder call a two-word Forever character.
func firstWord(name string) string {
	name = strings.TrimSpace(name)
	if i := strings.IndexRune(name, ' '); i >= 0 {
		return name[:i]
	}
	return name
}

func flat(s string) string {
	var b strings.Builder
	for _, r := range strings.ToLower(s) {
		if r >= 'a' && r <= 'z' || r >= '0' && r <= '9' || r > 127 {
			b.WriteRune(r)
		}
	}
	return b.String()
}
