package character

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/db"
)

func TestKeyIsRegionRulesetAndNameSlug(t *testing.T) {
	for _, tc := range []struct{ region, ruleset, name, want string }{
		{"us", "normal", "Baelgrim", "us/normal/baelgrim"},
		{"US", "Hardcore", "Baelgrim", "us/hardcore/baelgrim"},
		{"eu", "rp", "Lady Sunwick", "eu/rp/lady-sunwick"},
	} {
		if got := Key(tc.region, tc.ruleset, tc.name); got != tc.want {
			t.Errorf("Key(%q, %q, %q) = %q, want %q", tc.region, tc.ruleset, tc.name, got, tc.want)
		}
	}
}

func TestValidRulesetAndRegion(t *testing.T) {
	for _, r := range []string{"normal", "pvp", "rp", "hardcore"} {
		if !ValidRuleset(r) {
			t.Errorf("%q should be a ruleset", r)
		}
	}
	if ValidRuleset("nightslayer") {
		t.Error("a realm name is not a ruleset")
	}
	if !ValidRegion("eu") || ValidRegion("mars") {
		t.Error("regions are the five the contract names")
	}
}

func TestRulesetFromRealmTakesTheSegmentThenTheFallback(t *testing.T) {
	if got := RulesetFromRealm("Hardcore", "normal"); got != "hardcore" {
		t.Errorf("a segment that names a ruleset wins: %q", got)
	}
	if got := RulesetFromRealm("Nightslayer", "pvp"); got != "pvp" {
		t.Errorf("an unknown segment takes the report's ruleset: %q", got)
	}
	if got := RulesetFromRealm("Nightslayer", ""); got != RulesetNormal {
		t.Errorf("with no fallback the default is normal: %q", got)
	}
}

// The realm table RulesetFromRealm reads is the four ruleset names; each ruleset is reached
// from exactly one realm in it, and RealmForRuleset names that realm.
func TestRealmForRulesetIsTheOneRealmEachRulesetHas(t *testing.T) {
	realms := map[string][]string{}
	for _, realm := range Rulesets {
		if r := strings.ToLower(realm); ValidRuleset(r) {
			ruleset := RulesetFromRealm(realm, "")
			realms[ruleset] = append(realms[ruleset], realm)
		}
	}
	seen := map[string]string{}
	for _, ruleset := range Rulesets {
		if len(realms[ruleset]) != 1 {
			t.Fatalf("ruleset %q has realms %v, want exactly one", ruleset, realms[ruleset])
		}
		realm := RealmForRuleset(ruleset)
		if !strings.EqualFold(realm, realms[ruleset][0]) {
			t.Errorf("RealmForRuleset(%q) = %q, want %q", ruleset, realm, realms[ruleset][0])
		}
		if got := RulesetFromRealm(realm, RulesetHardcore); got != ruleset {
			t.Errorf("RulesetFromRealm(RealmForRuleset(%q)) = %q", ruleset, got)
		}
		if other, dup := seen[realm]; dup {
			t.Errorf("rulesets %q and %q share the realm %q", other, ruleset, realm)
		}
		seen[realm] = ruleset
	}
	for ruleset, want := range map[string]string{"normal": "Normal", "pvp": "Pvp", "rp": "Rp", "hardcore": "Hardcore"} {
		if got := RealmForRuleset(ruleset); got != want {
			t.Errorf("RealmForRuleset(%q) = %q, want %q", ruleset, got, want)
		}
	}
	if got := RealmForRuleset("Nightslayer"); got != "" {
		t.Errorf("a name that is not a ruleset has no realm: %q", got)
	}
}

func TestKeyFromUnitSplitsTheLoggedName(t *testing.T) {
	if got := KeyFromUnit("us", "hardcore", "Baelgrim-Nightslayer"); got != "us/hardcore/baelgrim" {
		t.Errorf("KeyFromUnit = %q", got)
	}
	if got := KeyFromUnit("us", "pvp", "Baelgrim"); got != "us/pvp/baelgrim" {
		t.Errorf("a unit with no realm segment is all name: %q", got)
	}
}

func TestValidKeyAcceptsTheContractsShapeAndNothingElse(t *testing.T) {
	for _, c := range []struct {
		key  string
		want bool
	}{
		{"us/normal/baelgrim", true},
		{"eu/rp/lady-sunwick", true},
		{"kr/hardcore/" + "사실", true},
		{"us/normal/bael9", true},
		{"", false},
		{"us/normal", false},
		{"us/normal/", false},
		{"mars/normal/baelgrim", false},
		{"us/nightslayer/baelgrim", false},
		{"us/normal/Baelgrim", false},
		{"us/normal/bael grim", false}, // Slug would have hyphenated it
		{"us/normal/bael/grim", false},
		// The slug crosses into a Lua file the addon loads, so the
		// characters that would close a string or a long bracket there
		// are refused at the door.
		{`us/normal/bael"grim`, false},
		{"us/normal/bael'grim", false},
		{`us/normal/bael\grim`, false},
		{"us/normal/bael]]grim", false},
		{"us/normal/bael\ngrim", false},
		{"us/normal/bael\x00grim", false},
		{"us/normal/bael\u00a0grim", false}, // a non-breaking space is still a space
		{"us/normal/bael\xffgrim", false},
		{"us/normal/" + strings.Repeat("a", MaxSlugBytes), true},
		{"us/normal/" + strings.Repeat("a", MaxSlugBytes+1), false},
	} {
		if got := ValidKey(c.key); got != c.want {
			t.Errorf("ValidKey(%q) = %v, want %v", c.key, got, c.want)
		}
	}
}

// linkedTestPool is a migrated pool on TEST_DATABASE_URL, or a skip without one.
func linkedTestPool(t *testing.T) *pgxpool.Pool {
	t.Helper()
	url := os.Getenv("TEST_DATABASE_URL")
	if url == "" {
		t.Skip("TEST_DATABASE_URL not set; start docker-compose.test.yml")
	}
	if err := db.Migrate(url); err != nil {
		t.Fatal(err)
	}
	pool, err := db.Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(pool.Close)
	return pool
}

// LinkedCharacters is /v1/me's list: every character the account owns, its main
// first and flagged, then the alts by key - and nobody else's characters.
func TestLinkedCharactersListsTheAccountsMainThenItsAlts(t *testing.T) {
	pool := linkedTestPool(t)
	ctx := context.Background()
	var owner, stranger int64
	for email, id := range map[string]*int64{
		"linked-owner@example.com": &owner, "linked-stranger@example.com": &stranger,
	} {
		if err := pool.QueryRow(ctx, `insert into users (email) values ($1)
			on conflict (email) do update set email = excluded.email returning id`, email).Scan(id); err != nil {
			t.Fatal(err)
		}
	}
	chars := []struct {
		key, name string
		user      int64
	}{
		{"us/normal/linkedalta", "Linkedalta", owner},
		{"us/normal/linkedmain", "Linkedmain", owner},
		{"eu/pvp/linkedaltb", "Linkedaltb", owner},
		{"us/normal/linkedother", "Linkedother", stranger},
	}
	for _, c := range chars {
		region, rest, _ := strings.Cut(c.key, "/")
		ruleset, _, _ := strings.Cut(rest, "/")
		if _, err := pool.Exec(ctx, `insert into characters (key, region, ruleset, name, class, user_id)
			values ($1, $2, $3, $4, 'Warrior', $5)
			on conflict (key) do update set user_id = excluded.user_id`, c.key, region, ruleset, c.name, c.user); err != nil {
			t.Fatal(err)
		}
	}
	if _, err := pool.Exec(ctx, `update users set main_character_key = 'us/normal/linkedmain' where id = $1`, owner); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		pool.Exec(context.Background(), `delete from characters where key like '%/linked%'`)
		pool.Exec(context.Background(), `delete from users where email in ('linked-owner@example.com', 'linked-stranger@example.com')`)
	})

	got, err := LinkedCharacters(ctx, pool, owner)
	if err != nil {
		t.Fatal(err)
	}
	want := []Linked{
		{Key: "us/normal/linkedmain", Region: "us", Ruleset: "normal", Name: "Linkedmain", Class: "Warrior", Main: true},
		{Key: "eu/pvp/linkedaltb", Region: "eu", Ruleset: "pvp", Name: "Linkedaltb", Class: "Warrior"},
		{Key: "us/normal/linkedalta", Region: "us", Ruleset: "normal", Name: "Linkedalta", Class: "Warrior"},
	}
	if len(got) != len(want) {
		t.Fatalf("LinkedCharacters = %+v, want %+v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("LinkedCharacters[%d] = %+v, want %+v", i, got[i], want[i])
		}
	}

	none, err := LinkedCharacters(ctx, pool, -1)
	if err != nil {
		t.Fatal(err)
	}
	if len(none) != 0 {
		t.Errorf("an account with no characters links %+v, want none", none)
	}
}
