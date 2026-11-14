// api/internal/bnetimport/directory_ruleset_test.go
package bnetimport

import (
	"context"
	"testing"
)

func clearDirectory(t *testing.T, svc *Service) {
	t.Helper()
	if _, err := svc.Pool.Exec(context.Background(), `truncate character_directory, realm_rulesets`); err != nil {
		t.Fatal(err)
	}
}

// A realm the directory learned keys its characters by its ruleset at sign-in: the player is
// never asked.
func TestImportUsesALearnedRealmBeforeAsking(t *testing.T) {
	svc, f, _ := unknownRealmService(t)
	clearDirectory(t, svc)
	if _, err := svc.Pool.Exec(context.Background(),
		`insert into realm_rulesets (region, realm_id, ruleset, evidence) values ('us', $1, 'pvp', 3)`, unknownRealm); err != nil {
		t.Fatal(err)
	}
	if _, key := importOne(t, svc, f, "Newly", "Rolled", 9100, unknownRealm); key != "us/pvp/newly-rolled" {
		t.Fatalf("key = %q, want the learned realm's pvp", key)
	}
}

// A character the directory knows from its guild's roster (not public yet) takes the
// guild's ruleset at sign-in, whatever its realm.
func TestImportUsesTheDirectoryRowOfTheCharacter(t *testing.T) {
	svc, f, _ := unknownRealmService(t)
	clearDirectory(t, svc)
	if _, err := svc.Pool.Exec(context.Background(),
		`insert into character_directory (region, bnet_character_id, name, slug, ruleset, ruleset_source)
		 values ('us', 9100, 'Newly Rolled', 'newly-rolled', 'rp', 'guild')`); err != nil {
		t.Fatal(err)
	}
	if _, key := importOne(t, svc, f, "Newly", "Rolled", 9100, unknownRealm); key != "us/rp/newly-rolled" {
		t.Fatalf("key = %q, want the guild's rp", key)
	}
}

// A realm whose evidence disagrees resolves nothing: the character still waits for its player.
func TestImportIgnoresARealmWhoseEvidenceDisagrees(t *testing.T) {
	svc, f, _ := unknownRealmService(t)
	clearDirectory(t, svc)
	if _, err := svc.Pool.Exec(context.Background(),
		`insert into realm_rulesets (region, realm_id, ruleset, evidence) values ('us', $1, 'pvp', 3), ('us', $1, 'normal', 1)`,
		unknownRealm); err != nil {
		t.Fatal(err)
	}
	if _, key := importOne(t, svc, f, "Newly", "Rolled", 9100, unknownRealm); key != "us/unknown/newly-rolled" {
		t.Fatalf("key = %q, want unknown", key)
	}
}
