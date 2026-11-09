package bnetimport

import (
	"context"
	"testing"
)

// The 2026-11-08 shape: a PvP character the import keyed normal, its main,
// export and guild link under the normal key, in a normal guild it created.
func seedMiskeyed(t *testing.T, svc *Service) int64 {
	t.Helper()
	ctx := context.Background()
	uid := seedUser(t, svc.Pool)
	if _, err := svc.Pool.Exec(ctx,
		`insert into guilds (id, region, ruleset, name) values (900, 'us', 'normal', 'Quiet Hours')`); err != nil {
		t.Fatal(err)
	}
	for _, q := range []string{
		`insert into characters (key, region, ruleset, name, class, user_id, source, bnet_profile)
		 values ('us/normal/anon-one', 'us', 'normal', 'Anon One', 'warrior', $1, 'bnet',
		         '{"ruleset":{"id":1,"name":"PvP"}}')`,
		`insert into guild_characters (guild_id, character_key, user_id, rank_index, rank, source)
		 values (900, 'us/normal/anon-one', $1, 0, 'leader', 'bnet')`,
		`update users set main_character_key = 'us/normal/anon-one' where id = $1`,
	} {
		if _, err := svc.Pool.Exec(ctx, q, uid); err != nil {
			t.Fatalf("%s: %v", q, err)
		}
	}
	return uid
}

func TestRepairRulesetsDryRunsThenMovesThenFindsNothing(t *testing.T) {
	svc := foreverService(t, foreverFixture(t))
	uid := seedMiskeyed(t, svc)
	ctx := context.Background()

	dry, err := svc.RepairRulesets(ctx, false)
	if err != nil {
		t.Fatal(err)
	}
	if len(dry.Moves) != 1 || dry.Moves[0].NewKey != "us/pvp/anon-one" || len(dry.Guilds) != 1 {
		t.Fatalf("dry run = %+v, want one move and one guild", dry)
	}
	var ruleset string
	if err := svc.Pool.QueryRow(ctx, `select ruleset from characters where user_id = $1`, uid).Scan(&ruleset); err != nil {
		t.Fatal(err)
	}
	if ruleset != "normal" {
		t.Fatal("the dry run wrote")
	}

	if _, err := svc.RepairRulesets(ctx, true); err != nil {
		t.Fatal(err)
	}
	var main, link, guildRuleset string
	if err := svc.Pool.QueryRow(ctx,
		`select u.main_character_key, gc.character_key, g.ruleset from users u
		   join guild_characters gc on gc.user_id = u.id join guilds g on g.id = gc.guild_id where u.id = $1`, uid).
		Scan(&main, &link, &guildRuleset); err != nil {
		t.Fatal(err)
	}
	if main != "us/pvp/anon-one" || link != "us/pvp/anon-one" || guildRuleset != "pvp" {
		t.Fatalf("main %q, link %q, guild %q; want all moved to pvp", main, link, guildRuleset)
	}

	again, err := svc.RepairRulesets(ctx, true)
	if err != nil {
		t.Fatal(err)
	}
	if len(again.Moves) != 0 || len(again.Guilds) != 0 {
		t.Fatalf("second run = %+v, want nothing", again)
	}
}

// Every column the schema names a character key by is one moveKey moves.
func TestCharacterKeyColumnsCoverTheSchema(t *testing.T) {
	pool := testPool(t)
	rows, err := pool.Query(context.Background(),
		`select c.table_name, c.column_name from information_schema.columns c
		  where c.table_schema = 'public'
		    and c.column_name in ('character_key', 'player_key', 'logging_character', 'main_character_key')
		    and c.table_name not like '%\_20__\___'
		    and not exists (select 1 from information_schema.key_column_usage k
		                      join information_schema.referential_constraints r on r.constraint_name = k.constraint_name
		                     where k.table_name = c.table_name and k.column_name = c.column_name
		                       and r.update_rule = 'CASCADE')`)
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()
	listed := map[string]bool{}
	for _, c := range characterKeyColumns {
		listed[c.table+"."+c.column] = true
	}
	for rows.Next() {
		var table, column string
		if err := rows.Scan(&table, &column); err != nil {
			t.Fatal(err)
		}
		if !listed[table+"."+column] {
			t.Errorf("%s.%s holds a character key that moveKey does not move", table, column)
		}
	}
}
