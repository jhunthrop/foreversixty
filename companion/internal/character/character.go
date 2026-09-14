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
