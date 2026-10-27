package guilds

import (
	"context"
	"testing"
)

func TestCharacterReadinessWord(t *testing.T) {
	h := newHTTPHarness(t)
	ctx := context.Background()
	gid := seedGuild(t, h.pool, "Forever")
	svc := &Service{Store: &Store{Pool: h.pool}}
	seed := func(name, consent, export string) string {
		t.Helper()
		key := "us/hardcore/" + name
		user := seedUser(t, h.pool, name+"@example.com")
		seedCharacter(t, h.pool, gid, user, key, "member", true)
		if _, err := h.pool.Exec(ctx, `
			insert into addon_exports (character_key, user_id, region, ruleset, name, export, captured_at, updated_at)
			values ($1, $2, 'us', 'hardcore', $3, $4, now(), now())`, key, user, name, export); err != nil {
			t.Fatal(err)
		}
		recomputeMembership(t, h.pool, gid, user)
		if _, err := h.pool.Exec(ctx, `update guild_members set consent = $3 where guild_id = $1 and user_id = $2`,
			gid, user, consent); err != nil {
			t.Fatal(err)
		}
		return key
	}
	const spent = "FS1:1.60.1.70009:warrior:human:555555555/33/0:"
	cases := []struct {
		key, want string
	}{
		{seed("stocked", "gear_bags", spent+"|bags=13510|level=60"), "Stocked"},
		{seed("unshared", "gear", spent+"|level=60"), "Not shared"},
		{seed("short", "gear_bags", "FS1:1.60.1.70009:warrior:human:0/0/0:|bags=13510|level=60"), "Short: 51 unspent talent points"},
		{"us/hardcore/nobody", "Not shared"},
	}
	for _, c := range cases {
		got, err := svc.CharacterReadinessWord(ctx, gid, c.key)
		if err != nil || got != c.want {
			t.Errorf("%s = %q, %v; want %q", c.key, got, err, c.want)
		}
	}
	facts, err := svc.CharacterReadinessFacts(ctx, gid, cases[0].key)
	if err != nil || facts.SyncedAt == nil || !facts.OnRoster {
		t.Fatalf("facts = %+v, %v", facts, err)
	}
	if facts, _ := svc.CharacterReadinessFacts(ctx, gid, "us/hardcore/nobody"); facts.OnRoster || facts.SyncedAt != nil {
		t.Fatalf("off-roster facts = %+v", facts)
	}
}
