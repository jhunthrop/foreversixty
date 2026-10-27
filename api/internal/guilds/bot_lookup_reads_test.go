package guilds

import (
	"context"
	"strings"
	"testing"
)

func TestBotLookupReads(t *testing.T) {
	h := newHTTPHarness(t)
	ctx := context.Background()
	gid := seedGuild(t, h.pool, "Olympus XXVII")
	if _, err := h.pool.Exec(ctx, `update guilds set faction = 'horde' where id = $1`, gid); err != nil {
		t.Fatal(err)
	}
	add := func(name string, verified, anonymised bool) {
		t.Helper()
		user := seedUser(t, h.pool, strings.ToLower(name)+"@example.com")
		key := "us/hardcore/" + strings.ToLower(name)
		if _, err := h.pool.Exec(ctx, `insert into characters (key, region, ruleset, name, class, user_id)
			values ($1, 'us', 'hardcore', $2, 'hunter', $3)`, key, name, user); err != nil {
			t.Fatal(err)
		}
		if _, err := h.pool.Exec(ctx, `update users set anonymize = $2 where id = $1`, user, anonymised); err != nil {
			t.Fatal(err)
		}
		seedCharacter(t, h.pool, gid, user, key, "member", verified)
	}
	add("Zulmara", true, false)
	add("Zarek", true, false)
	add("Pendrel", false, false)
	add("Hidden", true, true)

	real := realTreesStore(t)
	svc := &Service{Store: &Store{Pool: h.pool, DataDir: real.DataDir, Trees: real.Trees}}
	names, err := svc.PublicRosterNames(ctx, gid, "", 25)
	if err != nil || strings.Join(names, ",") != "Zarek,Zulmara" {
		t.Fatalf("names = %v, %v (the unverified and the anonymised must not appear)", names, err)
	}
	if names, _ := svc.PublicRosterNames(ctx, gid, "ULM", 25); strings.Join(names, ",") != "Zulmara" {
		t.Fatalf("typed = %v", names)
	}
	if names, _ := svc.PublicRosterNames(ctx, gid, "", 1); len(names) != 1 {
		t.Fatalf("limit = %v", names)
	}

	g, err := svc.BotGuildFacts(ctx, gid)
	if err != nil || g != (BotGuild{Name: "Olympus XXVII", Region: "us", Ruleset: "hardcore", Faction: "horde"}) {
		t.Fatalf("guild = %+v, %v", g, err)
	}

	build := svc.BisBuild()
	if build == "" {
		t.Fatal("no BiS build in the repo data")
	}
	// Helm of Wrath (16963): a warrior item, not a mage's.
	if it, ok := svc.BuildItem(build, "warrior", 16963); !ok || it.Name != "Helm of Wrath" || it.Icon == "" || it.Quality != 4 {
		t.Fatalf("warrior item = %+v, %v", it, ok)
	}
	if _, ok := svc.BuildItem(build, "mage", 16963); ok {
		t.Fatal("a mage resolved a plate helm")
	}
	if _, ok := svc.BuildItem("0.0.0.0", "warrior", 16963); ok {
		t.Fatal("an unknown build resolved an item")
	}
	if svc.RaceFaction("Orc") != FactionHorde || svc.RaceFaction("Night Elf") != FactionAlliance || svc.RaceFaction("") != "" {
		t.Fatalf("race factions %q %q", svc.RaceFaction("Orc"), svc.RaceFaction("Night Elf"))
	}
}
