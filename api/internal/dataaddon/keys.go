// api/internal/dataaddon/keys.go
package dataaddon

import (
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/character"
)

// characterKey is the addon's lookup key for a character (contract 5): "Name-Realm", the
// name as the game shows it and the realm a client on the character's ruleset sees
// (character.RealmForRuleset, the inverse of the API's one realm-to-ruleset mapping). The
// addon's Score.forCharacter builds the same string from UnitName and GetRealmName. Empty
// when the ruleset has no realm.
func characterKey(name, ruleset string) string {
	realm := character.RealmForRuleset(ruleset)
	if name == "" || realm == "" {
		return ""
	}
	return name + "-" + realm
}

// displayName is the name a character is keyed by: the stored display name (characters.name,
// else the newest rated row's player_name with any "-Realm" the log printed cut off), else
// the key's own name slug when nothing better was ever stored.
func displayName(stored, slug string) string {
	name, _ := character.SplitUnit(strings.TrimSpace(stored))
	if name == "" {
		return slug
	}
	return name
}

// profilePath is the character's page on the site, "/character/<region>/<ruleset>/<slug>",
// which the addon shows only as a copyable address.
func profilePath(region, ruleset, slug string) string {
	return "/character/" + region + "/" + ruleset + "/" + slug
}

// splitCharacterKey reads region/ruleset/name-slug back out of a character key
// ("us/normal/thoradin" -> "us", "normal", "thoradin", true), the inverse of
// character.Key for a well-formed key.
func splitCharacterKey(key string) (region, ruleset, slug string, ok bool) {
	parts := strings.SplitN(key, "/", 3)
	if len(parts) != 3 || parts[0] == "" || parts[1] == "" || parts[2] == "" {
		return "", "", "", false
	}
	return parts[0], parts[1], parts[2], true
}
