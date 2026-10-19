// companion/internal/character/character.go
// Package character holds the one attribution triple the config, the
// report state and the wire all share, so the ruleset field is
// spelled once.
package character

import "strings"

// Character is region, ruleset and name.
//
// Forever has no realms: the second segment of a character key is the
// ruleset (normal, pvp, rp, hardcore). The combat log and the addon
// still identify units as Name-Realm, so the companion carries that
// segment through untouched and the API maps it once the Sept 17 beta
// log settles what it holds. Nothing here validates the value against
// a list.
type Character struct {
	Region  string `json:"region"`
	Ruleset string `json:"ruleset"`
	Name    string `json:"name"`
}

// Valid reports whether all three segments are filled in.
func (c Character) Valid() bool {
	return strings.TrimSpace(c.Region) != "" &&
		strings.TrimSpace(c.Ruleset) != "" &&
		strings.TrimSpace(c.Name) != ""
}

// Key is the contract's character key: region, ruleset and the
// two-part name lowercased with spaces turned into hyphens.
func (c Character) Key() string {
	slug := strings.ToLower(strings.ReplaceAll(strings.TrimSpace(c.Name), " ", "-"))
	return strings.ToLower(c.Region) + "/" + strings.ToLower(c.Ruleset) + "/" + slug
}

// RulesetFromRealm is the API's mapping of a realm segment onto a
// ruleset, kept in step with api/internal/character.RulesetFromRealm:
// a realm names its ruleset by a word in it ("Hardcore", "Classic Beta
// PvP", the normalised "ClassicBetaPvP2"), and a realm with no such
// word is a normal one.
func RulesetFromRealm(realm string) string {
	words := strings.FieldsFunc(realm, func(r rune) bool {
		return !(r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9')
	})
	for _, word := range words {
		// Trailing digits number a realm ("PvE 2", "ClassicBetaPvE2").
		trimmed := strings.TrimRight(strings.ToLower(word), "0123456789")
		camel := camelCase(word)
		for _, w := range rulesetWords {
			if trimmed == w.word || camel && strings.HasSuffix(trimmed, w.word) {
				return w.ruleset
			}
		}
	}
	return "normal"
}

// rulesetWords are the words a realm spells a ruleset with, longest
// first so a camel-case suffix match tries "hardcore" before "hc".
var rulesetWords = []struct{ word, ruleset string }{
	{"hardcore", "hardcore"}, {"roleplay", "rp"}, {"normal", "normal"},
	{"pvp", "pvp"}, {"pve", "normal"}, {"rp", "rp"}, {"hc", "hardcore"},
}

// camelCase reports whether word joins several words in the normalised
// form: a capital after a lower-case letter ("ClassicBetaPvP").
func camelCase(word string) bool {
	for i := 1; i < len(word); i++ {
		if word[i] >= 'A' && word[i] <= 'Z' && word[i-1] >= 'a' && word[i-1] <= 'z' {
			return true
		}
	}
	return false
}
