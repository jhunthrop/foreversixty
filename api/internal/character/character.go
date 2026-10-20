// Package character is the one place the API turns a name into a
// character key, and the one place the combat log's realm segment
// becomes a ruleset.
//
// Forever has no realms: the Deep Dive panel replaced them with four
// rulesets per region. A character key is
// <region>/<ruleset>/<name-slug>.
package character

import (
	"context"
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/jackc/pgx/v5"
)

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

// MaxSlugBytes bounds a name-slug. A character name is at most twelve
// characters and a guild name at most twenty-four; sixty-four bytes is
// room for either written in any script.
const MaxSlugBytes = 64

// ValidSlug reports whether s is a canonical name-slug: non-empty, at
// most MaxSlugBytes, unchanged by Slug, valid UTF-8, and made only of
// letters, digits and hyphens.
//
// The charset matters beyond tidiness. A character key crosses the
// addon inbox into a Lua source file the addon loads, so a slug
// carrying a quote, a bracket, a backslash or a newline is an
// injection waiting for a companion that forgets to escape. Nothing in
// a real name needs an ASCII character outside [a-z0-9-], and
// non-ASCII letters - which kr, tw and cn names are made of - are
// allowed through unharmed.
func ValidSlug(s string) bool {
	if s == "" || len(s) > MaxSlugBytes || !utf8.ValidString(s) || Slug(s) != s {
		return false
	}
	for _, r := range s {
		switch {
		case r == '-' || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9'):
		case r < utf8.RuneSelf || unicode.IsControl(r) || unicode.IsSpace(r):
			return false
		}
	}
	return true
}

// ValidKey reports whether key is a well-formed character key:
// <region>/<ruleset>/<name-slug>, with all three parts valid.
func ValidKey(key string) bool {
	region, rest, ok := strings.Cut(key, "/")
	if !ok {
		return false
	}
	ruleset, slug, ok := strings.Cut(rest, "/")
	if !ok {
		return false
	}
	return ValidRegion(region) && ValidRuleset(ruleset) && ValidSlug(slug)
}

// Key builds a character key. Region and ruleset are lowercased but not
// validated here: the handlers validate what a client sent, and the
// ingest maps what the log carried.
func Key(region, ruleset, name string) string {
	return strings.ToLower(region) + "/" + strings.ToLower(ruleset) + "/" + Slug(name)
}

// SplitUnit separates a combat log's unit into the character's name and the realm segment.
// The client writes a Forever character's two-word name with a hyphen between the words and
// ends the unit with the normalised realm and a trailing hyphen ("Bow-Jackzon-ClassicBetaPvP-"),
// while older logs wrote "Bow-ClassicBetaPvP-" or "Mörk-ClassicBetaPvP". So the trailing
// hyphen is dropped and the last segment is the realm; the segments before it are the name,
// joined with single spaces so Slug gives the export's key ("bow-jackzon"). A unit with one
// segment is all name and has no realm.
func SplitUnit(unit string) (name, realm string) {
	segments := strings.Split(strings.TrimSuffix(strings.TrimSpace(unit), "-"), "-")
	if len(segments) == 1 {
		return segments[0], ""
	}
	last := len(segments) - 1
	return strings.Join(segments[:last], " "), segments[last]
}

// Realm is the realm a character lives on as the client names it, from the first source
// that has one: the Battle.net import's realm name (characters.realm_name), else the addon
// export's who= realm, else the realm half of the newest rated row's logged unit
// ("Mörk-ClassicBetaPvP", which the combat log writes in the client's normalised form).
// Empty when no source has one. The nightly data addon's keys and the character page both
// read a character's realm through this one order.
func Realm(bnetRealm, exportRealm, loggedUnit string) string {
	for _, realm := range []string{bnetRealm, exportRealm} {
		if realm = strings.TrimSpace(realm); realm != "" {
			return realm
		}
	}
	_, realm := SplitUnit(strings.TrimSpace(loggedUnit))
	return realm
}

// RulesetFromRealm maps the realm segment of a logged unit name (or a realm name the
// client or Battle.net reports) onto a ruleset. This is the single mapping the whole API
// uses.
//
// A realm names its ruleset by a word in it: the ruleset itself once Forever's realms are
// its rulesets ("Hardcore"), and the beta's realm names ("Classic Beta PvP", "Classic Beta
// PvE 2"), the same words the web's rulesetFromRealm reads (rulesetWords). The client's
// normalised realm name drops the spaces and hyphens ("ClassicBetaPvP",
// GetNormalizedRealmName and the combat log's Name-Realm), so a single camel-case word is
// read by the ruleset word it ends in, trailing digits aside. A realm with no such word
// takes the fallback - the ruleset the report's logging character declared - and
// RulesetNormal when the report declared none.
func RulesetFromRealm(realm, fallback string) string {
	if r, ok := rulesetOfRealm(realm); ok {
		return r
	}
	if ValidRuleset(fallback) {
		return fallback
	}
	return RulesetNormal
}

// rulesetWords are the words a realm name spells a ruleset with, longest first so a
// camel-case suffix match tries "hardcore" before "hc".
var rulesetWords = []struct{ word, ruleset string }{
	{"hardcore", RulesetHardcore}, {"roleplay", RulesetRP}, {"normal", RulesetNormal},
	{"pvp", RulesetPvP}, {"pve", RulesetNormal}, {"rp", RulesetRP}, {"hc", RulesetHardcore},
}

// rulesetOfRealm is the ruleset the first word of realm that names one spells, if any.
func rulesetOfRealm(realm string) (string, bool) {
	words := strings.FieldsFunc(realm, func(r rune) bool {
		return !(r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9')
	})
	for _, word := range words {
		// Trailing digits number a realm ("PvE 2", "ClassicBetaPvE2"), they name no ruleset.
		trimmed := strings.TrimRight(strings.ToLower(word), "0123456789")
		camel := camelCase(word)
		for _, w := range rulesetWords {
			if trimmed == w.word || camel && strings.HasSuffix(trimmed, w.word) {
				return w.ruleset, true
			}
		}
	}
	return "", false
}

// camelCase reports whether word joins several words in the normalised form: a capital
// letter after a lower-case one ("ClassicBetaPvP", "ForeverRP").
func camelCase(word string) bool {
	for i := 1; i < len(word); i++ {
		if word[i] >= 'A' && word[i] <= 'Z' && word[i-1] >= 'a' && word[i-1] <= 'z' {
			return true
		}
	}
	return false
}

// KeyFromUnit builds the character key for a unit name out of a combat
// log, in a report whose region and ruleset are region and ruleset.
func KeyFromUnit(region, ruleset, unit string) string {
	name, realm := SplitUnit(unit)
	return Key(region, RulesetFromRealm(realm, ruleset), name)
}

// Querier is the one database call LinkedCharacters makes; a pool and a
// transaction both have it.
type Querier interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
}

// Linked is one character an account owns. Main is true for the account's
// main (users.main_character_key); every other linked character is an alt.
type Linked struct {
	Key     string
	Region  string
	Ruleset string
	Name    string
	Class   string
	Main    bool
}

// LinkedCharacters lists the characters an account owns - exactly the
// characters GET /v1/me lists (characters.user_id, which only the account's
// own verified paths ever set and which never moves a character off another
// account) - with the main first, then the alts by key. An account with no
// characters has none, and that is not an error.
func LinkedCharacters(ctx context.Context, q Querier, userID int64) ([]Linked, error) {
	rows, err := q.Query(ctx,
		`select c.key, c.region, c.ruleset, c.name, coalesce(c.class, ''),
		        c.key = coalesce(u.main_character_key, '')
		 from characters c join users u on u.id = c.user_id
		 where c.user_id = $1
		 order by 6 desc, c.key`, userID)
	if err != nil {
		return nil, fmt.Errorf("character: linked characters of %d: %w", userID, err)
	}
	defer rows.Close()
	out := []Linked{}
	for rows.Next() {
		var l Linked
		if err := rows.Scan(&l.Key, &l.Region, &l.Ruleset, &l.Name, &l.Class, &l.Main); err != nil {
			return nil, fmt.Errorf("character: scan linked character: %w", err)
		}
		out = append(out, l)
	}
	return out, rows.Err()
}
