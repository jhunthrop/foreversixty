// api/internal/dataaddon/keys.go
package dataaddon

import (
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/fs1"
)

// characterKey is the addon's lookup key for a character (contract 5): "Name-Realm", the
// name as the game shows it and the client's own realm in its normalised form
// (normalRealm). Score.forCharacter looks a character up as Name-<GetRealmName()>, then as
// Name-<the realm with its spaces and hyphens removed>; the file writes the second form,
// which every spelling of the realm reduces to. Empty when either half is unknown.
func characterKey(name, realm string) string {
	realm = normalRealm(realm)
	if name == "" || realm == "" {
		return ""
	}
	return name + "-" + realm
}

// normalRealm is a realm name with its spaces and hyphens removed ("Classic Beta PvE 2" ->
// "ClassicBetaPvE2"), the client's GetNormalizedRealmName form and the second lookup
// Score.forCharacter makes.
func normalRealm(realm string) string {
	return strings.Map(func(r rune) rune {
		if r == '-' || r == ' ' || r == '\t' {
			return -1
		}
		return r
	}, strings.TrimSpace(realm))
}

// identity is what the database holds about who a character is in game.
type identity struct {
	// Name is characters.name, the display name the site stores.
	Name string
	// RealmName is characters.realm_name, the realm the Battle.net import read.
	RealmName string
	// PlayerName is the newest rated row's player_name, as the combat log printed the unit
	// ("Mörk-ClassicBetaPvP").
	PlayerName string
}

// nameAndRealm is the name and realm a character is keyed by. The name is the stored
// display name, else the name half of the newest rated row's unit. The realm is
// character.Realm's: the Battle.net import's realm name, else the stored export's who=
// realm, else the realm half of the newest rated row's unit (the order the character page
// reads too). Either is empty when no source has it, and the caller writes no record for
// that character.
func nameAndRealm(id identity, export string) (name, realm string) {
	loggedName, _ := character.SplitUnit(strings.TrimSpace(id.PlayerName))
	name = strings.TrimSpace(id.Name)
	if name == "" {
		name = loggedName
	}
	var whoRealm string
	if export != "" {
		if decoded, ok := fs1.Decode(export); ok && decoded.HasWho {
			whoRealm = decoded.Realm
		}
	}
	return name, character.Realm(id.RealmName, whoRealm, id.PlayerName)
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
