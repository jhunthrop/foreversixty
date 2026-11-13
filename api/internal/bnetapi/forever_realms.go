// api/internal/bnetapi/forever_realms.go
package bnetapi

import "strings"

// foreverGame is the profile-game segment of Forever's namespaces, the only
// game the committed realm table below applies to.
const foreverGame = "forever"

// foreverRealms is the committed Forever realm table: region, then realm id
// (the id the account profile's protected-character link carries), to the
// realm's Blizzard type in the vocabulary RulesetOf reads.
//
// Blizzard's game-data API lists no Forever realm at all: the realm index and
// the connected-realm index answer 404 in both regions, as does every realm
// by id or slug, and /data/wow/search/realm answers an empty page (probed
// with the app's own credentials 2026-11-12; the index had answered until
// 2026-11-05). The character profile names a ruleset only once the character
// is public, which a fresh character is not, so without this table every
// such character was skipped at sign-in (needs_ruleset, 40 to 80 a day).
//
// Evidence per row: the three named realms are the live index probed
// 2026-11-03 (testdata/forever/realm-index-us.json); every other type is the
// ruleset Blizzard's own profile answered for characters of that realm in
// production (characters.bnet_profile.ruleset.name, read 2026-11-12), or,
// for 6922, the key the realm index itself gave its rows before 2026-11-05.
// US 6929 and EU 6942 are the profile rulesets of characters on those realms,
// re-read live from Blizzard's character profile 2026-11-13.
// A realm id not listed here resolves nothing (KnownRealm reports !ok) and
// the import falls through to the profile, then to its players' answers
// (bnetimport/unknown_realm.go); the needs_ruleset log line names the id so
// it can be added once a profile of that realm has answered.
var foreverRealms = map[string]map[int64]Realm{
	"us": {
		6794: {ID: 6794, Slug: "camwow-us1-realm-1-pve", Name: "CAMWOW US1 Realm 1 PvE", Type: "NORMAL"},
		6796: {ID: 6796, Type: "NORMAL"},
		6798: {ID: 6798, Slug: "camwow-us1-realm-2-pvp", Name: "CAMWOW US1 Realm 2 PvP", Type: "PVP"},
		6800: {ID: 6800, Type: "PVP"},
		6801: {ID: 6801, Type: "PVP"},
		6921: {ID: 6921, Type: "NORMAL"},
		6922: {ID: 6922, Type: "PVP"},
		6923: {ID: 6923, Type: "NORMAL"},
		6927: {ID: 6927, Type: "PVP_RP"},
		6928: {ID: 6928, Slug: "camwow-us1-realm-10-rp", Name: "CAMWOW US1 Realm 10 RP", Type: "RP"},
		6929: {ID: 6929, Type: "NORMAL"},
		6930: {ID: 6930, Type: "PVP"},
		6932: {ID: 6932, Type: "PVP"},
		6934: {ID: 6934, Type: "PVP"},
		6935: {ID: 6935, Type: "PVP_RP"},
	},
	"eu": {
		6841: {ID: 6841, Type: "NORMAL"},
		6855: {ID: 6855, Type: "NORMAL"},
		6938: {ID: 6938, Type: "PVP"},
		6939: {ID: 6939, Type: "NORMAL"},
		6940: {ID: 6940, Type: "PVP"},
		6942: {ID: 6942, Type: "PVP"},
		6945: {ID: 6945, Type: "NORMAL"},
		6946: {ID: 6946, Type: "PVP"},
		6947: {ID: 6947, Type: "NORMAL"},
		6949: {ID: 6949, Type: "NORMAL"},
		6950: {ID: 6950, Type: "PVP"},
		6952: {ID: 6952, Type: "RP"},
	},
}

// KnownRealm is the committed realm table's entry for a realm id in region,
// for a client whose game is Forever; ok is false for any other game, a
// region the table does not hold, or an id it has no evidence for. It is the
// realm-id source rulesetFor reads after Realms, which lists nothing for
// Forever (see foreverRealms).
func (c *Client) KnownRealm(region string, id int64) (Realm, bool) {
	if c.Game != foreverGame {
		return Realm{}, false
	}
	realm, ok := foreverRealms[strings.ToLower(region)][id]
	return realm, ok
}
