package guilds

import "testing"

// is_main marks the character its account chose as main, so the roster names an alt's main.
func TestHomeRosterMarksTheMain(t *testing.T) {
	h := newHTTPHarness(t)
	gid := seedGuild(t, h.pool, "Mains")
	uid := seedUser(t, h.pool, "mains@example.com")
	for _, name := range []string{"main", "alt"} {
		key := "us/hardcore/" + name
		seedCharacter(t, h.pool, gid, uid, key, "member", true)
		seedExport(t, h.pool, uid, key, "us", "hardcore", name)
	}
	recomputeMembership(t, h.pool, gid, uid)
	if _, err := h.pool.Exec(t.Context(), `update users set main_character_key = 'us/hardcore/main' where id = $1`, uid); err != nil {
		t.Fatal(err)
	}
	roster, err := h.store.HomeRoster(t.Context(), gid, Guild{}, 0, false, false)
	if err != nil {
		t.Fatal(err)
	}
	got := map[string]bool{}
	for _, row := range roster {
		got[row.CharacterKey] = row.IsMain
	}
	if len(got) != 2 || !got["us/hardcore/main"] || got["us/hardcore/alt"] {
		t.Fatalf("is_main = %v", got)
	}
}
