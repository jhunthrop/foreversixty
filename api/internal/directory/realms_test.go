// api/internal/directory/realms_test.go
package directory

import (
	"context"
	"testing"
)

// A log GUID's realm id takes the ruleset of the guild of the character its hex names (in the
// report's guild's region), stamps that character's realm id, and types a realm-mate the
// directory knew without a ruleset.
func TestLearnRealmsFromLogGUIDs(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	if _, err := pool.Exec(ctx, `truncate realm_rulesets, reports cascade`); err != nil {
		t.Fatal(err)
	}
	b := newBlizzard(t)
	s := b.service(pool)
	var guildID int64
	if err := pool.QueryRow(ctx,
		`insert into guilds (region, ruleset, name) values ('us', 'pvp', 'Ember Wardens') returning id`).Scan(&guildID); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx, `
		insert into character_directory (region, bnet_character_id, name, slug, ruleset, ruleset_source) values
		  ('us', 255, 'Vela Thornbrook', 'vela-thornbrook', 'pvp', 'guild'),
		  ('us', 4096, 'Orrin Ashmantle', 'orrin-ashmantle', null, null),
		  ('eu', 255, 'Sel Duskwater', 'sel-duskwater', 'normal', 'profile')`); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx,
		`insert into reports (id, guild_id, visibility, status) values ('r1', $1, 'public', 'complete')`, guildID); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx, `insert into fights (report_id, fight_index, players)
		values ('r1', 1, array['Player-6924-000000FF', 'Player-6924-00001000', 'Creature-0-1'])`); err != nil {
		t.Fatal(err)
	}
	learned, err := s.learnRealms(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if learned != 1 {
		t.Fatalf("learned = %d, want the realm-mate typed", learned)
	}
	if ruleset, ok, err := RealmRuleset(ctx, pool, "us", 6924); err != nil || !ok || ruleset != "pvp" {
		t.Fatalf("realm 6924 = %q %v %v, want pvp", ruleset, ok, err)
	}
	if _, ok, _ := RealmRuleset(ctx, pool, "eu", 6924); ok {
		t.Fatal("an eu character with the same id typed a us realm")
	}
	if ruleset, ok, err := CharacterRuleset(ctx, pool, "us", 4096); err != nil || !ok || ruleset != "pvp" {
		t.Fatalf("realm-mate = %q %v %v, want pvp from its realm", ruleset, ok, err)
	}
}
