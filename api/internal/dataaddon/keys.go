// api/internal/dataaddon/keys.go
package dataaddon

import (
	"regexp"
	"strings"
)

// whitespaceRun matches the same class the addon's own slug function does
// (Ratings.lua: tostring(text):lower():gsub("%s+", "-")). Lua's %s class is
// ASCII whitespace (space, tab, newline, CR, FF, VT); Go's \s matches it
// identically for the ASCII range every real character or guild name is
// drawn from.
var whitespaceRun = regexp.MustCompile(`\s+`)

// slug lowercases s and collapses every run of whitespace into a single
// hyphen -- byte-for-byte what the addon's own Ratings.lua computes
// client-side from GetRealmName(), a character name, and a guild name.
//
// This cannot reuse api/internal/character.Slug: that helper trims the
// string first and replaces only the literal space character one-for-one,
// which is a different function for a name with a tab or a doubled space
// (neither occurs in a real character name, but a guild name can have a
// doubled space, and the addon's own gsub is the contract this job has to
// match exactly, not approximate).
func slug(s string) string {
	return whitespaceRun.ReplaceAllString(strings.ToLower(s), "-")
}

// characterKey builds the addon's lookup key for a character:
// "<region>:<ruleset>:<name-slug>". The dispatch's "realm" position is
// filled with the site's ruleset -- see the plan's "Key-format
// reconciliation" section for why.
func characterKey(region, ruleset, name string) string {
	return slug(region) + ":" + slug(ruleset) + ":" + slug(name)
}

// guildKey is characterKey's twin for a guild: the same three parts, the
// guild's own name in the third.
func guildKey(region, ruleset, name string) string {
	return characterKey(region, ruleset, name)
}
