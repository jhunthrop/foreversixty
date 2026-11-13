// api/internal/bnetapi/forever_realms_test.go
package bnetapi

import (
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/character"
)

// productionRealmIDs is every realm id the import's needs_ruleset lines
// named in Cloud Run logs from 2026-11-12 to 2026-11-13, per region, with the ruleset the
// production evidence (profile rulesets on rows of the same realm) gives it.
// The table must cover each one, or the next sign-in from that realm skips
// its characters again.
var productionRealmIDs = map[string]map[int64]string{
	"us": {
		6800: character.RulesetPvP,
		6801: character.RulesetPvP,
		6929: character.RulesetNormal,
		6930: character.RulesetPvP,
	},
	"eu": {
		6939: character.RulesetNormal,
		6942: character.RulesetPvP,
		6947: character.RulesetNormal,
		6950: character.RulesetPvP,
	},
}

func TestKnownRealmCoversEveryRealmSeenInProductionLogs(t *testing.T) {
	c := New(Config{Game: "forever"})
	for region, want := range productionRealmIDs {
		for id, ruleset := range want {
			realm, ok := c.KnownRealm(region, id)
			if !ok {
				t.Errorf("%s realm %d is not in the Forever realm table", region, id)
				continue
			}
			if got, ok := RulesetOf(realm); !ok || got != ruleset {
				t.Errorf("%s realm %d resolves %q, %v; want %q", region, id, got, ok, ruleset)
			}
		}
	}
}

func TestKnownRealmMapsEveryTableEntryToARuleset(t *testing.T) {
	c := New(Config{Game: "forever"})
	for region, realms := range foreverRealms {
		for id := range realms {
			realm, ok := c.KnownRealm(region, id)
			if !ok {
				t.Fatalf("%s realm %d: KnownRealm misses its own table", region, id)
			}
			if realm.ID != id {
				t.Errorf("%s realm %d: ID = %d", region, id, realm.ID)
			}
			if _, ok := RulesetOf(realm); !ok {
				t.Errorf("%s realm %d: type %q resolves no ruleset", region, id, realm.Type)
			}
		}
	}
}

func TestKnownRealmReadsTheLiveRealmTypes(t *testing.T) {
	c := New(Config{Game: "forever"})
	cases := []struct {
		region string
		id     int64
		want   string
	}{
		{"us", 6794, character.RulesetNormal}, // CAMWOW US1 Realm 1 PvE, probed 2026-11-03
		{"us", 6798, character.RulesetPvP},    // CAMWOW US1 Realm 2 PvP
		{"us", 6928, character.RulesetRP},     // CAMWOW US1 Realm 10 RP
		{"us", 6927, character.RulesetRP},     // profile ruleset "Roleplay PvP"
		{"eu", 6952, character.RulesetRP},     // profile ruleset "Roleplay"
	}
	for _, tc := range cases {
		realm, ok := c.KnownRealm(tc.region, tc.id)
		if !ok {
			t.Errorf("%s realm %d missing", tc.region, tc.id)
			continue
		}
		if got, _ := RulesetOf(realm); got != tc.want {
			t.Errorf("%s realm %d = %q, want %q", tc.region, tc.id, got, tc.want)
		}
	}
}

func TestKnownRealmIsForeverOnlyAndNeverGuesses(t *testing.T) {
	if _, ok := New(Config{Game: "classic1x"}).KnownRealm("us", 6800); ok {
		t.Fatal("a classic1x client answered a Forever realm id")
	}
	c := New(Config{Game: "forever"})
	if _, ok := c.KnownRealm("us", 6931); ok {
		t.Fatal("realm 6931 has no production evidence yet and must not be in the table")
	}
	if _, ok := c.KnownRealm("kr", 6800); ok {
		t.Fatal("a region with no table answered")
	}
	if _, ok := c.KnownRealm("US", 6800); !ok {
		t.Fatal("the region is matched case-insensitively, like every namespace")
	}
}
