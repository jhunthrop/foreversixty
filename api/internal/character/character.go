// Package character is the one place the API turns a name into a
// character key, and the one place the combat log's realm segment
// becomes a ruleset.
//
// Forever has no realms: the Deep Dive panel replaced them with four
// rulesets per region. A character key is
// <region>/<ruleset>/<name-slug>.
package character

import "strings"

// The four rulesets, per the Phase 3 contract.
const (
	RulesetNormal   = "normal"
	RulesetPvP      = "pvp"
	RulesetRP       = "rp"
	RulesetHardcore = "hardcore"
)

// Rulesets is every valid ruleset, in the order the site lists them.
var Rulesets = []string{RulesetNormal, RulesetPvP, RulesetRP, RulesetHardcore}

// Regions is every valid region.
var Regions = []string{"us", "eu", "kr", "tw", "cn"}

// ValidRuleset reports whether s is one of the four rulesets.
func ValidRuleset(s string) bool { return in(Rulesets, s) }

// ValidRegion reports whether s is one of the five regions.
func ValidRegion(s string) bool { return in(Regions, s) }

func in(all []string, s string) bool {
	for _, v := range all {
		if v == s {
			return true
		}
	}
	return false
}

// Slug is the key form of a name: lowercased, with spaces as hyphens.
// Guild names are several words; character names are one, and take the
// same path so there is only ever one rule.
func Slug(name string) string {
	return strings.ReplaceAll(strings.ToLower(strings.TrimSpace(name)), " ", "-")
}

// Key builds a character key. Region and ruleset are lowercased but not
// validated here: the handlers validate what a client sent, and the
// ingest maps what the log carried.
func Key(region, ruleset, name string) string {
	return strings.ToLower(region) + "/" + strings.ToLower(ruleset) + "/" + Slug(name)
}

// SplitUnit separates a combat log's "Name-Realm" into its two halves.
// A unit with no hyphen is all name.
func SplitUnit(unit string) (name, realm string) {
	name, realm, _ = strings.Cut(unit, "-")
	return name, realm
}

// RulesetFromRealm maps the realm segment of a logged unit name onto a
// ruleset. This is the single mapping the whole API uses.
//
// The combat log prints units as Name-Realm today. What Forever writes
// in that segment is settled by the first beta log on Sept 17: if it
// writes a ruleset name, the first branch already handles it; if it
// writes something else, that translation belongs here and nowhere else.
// Until then an unrecognised segment takes the fallback - the ruleset
// the report's logging character declared - and RulesetNormal when the
// report declared none.
func RulesetFromRealm(realm, fallback string) string {
	if r := strings.ToLower(strings.TrimSpace(realm)); ValidRuleset(r) {
		return r
	}
	if ValidRuleset(fallback) {
		return fallback
	}
	return RulesetNormal
}

// KeyFromUnit builds the character key for a unit name out of a combat
// log, in a report whose region and ruleset are region and ruleset.
func KeyFromUnit(region, ruleset, unit string) string {
	name, realm := SplitUnit(unit)
	return Key(region, RulesetFromRealm(realm, ruleset), name)
}
