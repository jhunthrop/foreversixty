// api/internal/directory/refresh_test.go
package directory

import (
	"context"
	"encoding/json"
	"testing"
)

func seedGuildRow(t *testing.T, s *Service, name, ruleset string) {
	t.Helper()
	if _, err := s.Pool.Exec(context.Background(),
		`insert into guilds (region, ruleset, name) values ('us', $1, $2)`, ruleset, name); err != nil {
		t.Fatal(err)
	}
}

type dirRow struct {
	name, class, ruleset, source, spec string
	public                             bool
	guildID                            *int64
	itemLevel                          *int
	equipment                          []EquippedItem
	professions                        []Profession
	sources                            []string
}

func readRow(t *testing.T, s *Service, id int64) (dirRow, bool) {
	t.Helper()
	var r dirRow
	var eq, prof []byte
	err := s.Pool.QueryRow(context.Background(),
		`select name, coalesce(class, ''), coalesce(ruleset, ''), coalesce(ruleset_source, ''), coalesce(spec, ''),
		        public, guild_bnet_id, item_level, equipment, professions, sources
		   from character_directory where region = 'us' and bnet_character_id = $1`, id).
		Scan(&r.name, &r.class, &r.ruleset, &r.source, &r.spec, &r.public, &r.guildID, &r.itemLevel, &eq, &prof, &r.sources)
	if err != nil {
		return dirRow{}, false
	}
	if eq != nil {
		if err := json.Unmarshal(eq, &r.equipment); err != nil {
			t.Fatal(err)
		}
	}
	if prof != nil {
		if err := json.Unmarshal(prof, &r.professions); err != nil {
			t.Fatal(err)
		}
	}
	return r, true
}

// A known guild's roster names every member under the guild's ruleset; a public member gets
// its profile, gear, spec and professions; a member who is not public is kept bare; a
// member Blizzard says is invalid is deleted.
func TestPassFillsTheDirectoryFromAGuildRoster(t *testing.T) {
	pool := testPool(t)
	b := newBlizzard(t)
	s := b.service(pool)
	seedGuildRow(t, s, "Ember Wardens", "pvp")
	b.guild("ember-wardens", "Ember Wardens", "PvP", 77,
		rosterMember{id: 101, name: "Vela Thornbrook", level: 41, classID: 8},
		rosterMember{id: 102, name: "Orrin Ashmantle", level: 12, classID: 1, rank: 3},
		rosterMember{id: 103, name: "Sel Duskwater", level: 30, classID: 5})
	b.publicCharacter(101, "Vela Thornbrook", "PvP", 77, "ember-wardens", "Ember Wardens")
	b.on("/profile/wow/character/sel-duskwater/status", map[string]any{"id": 103, "is_valid": false})

	res, err := s.Pass(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if res.Rosters != 1 || res.RosterMembers != 3 || res.Profiles != 1 || res.Private != 1 || res.Deleted != 1 {
		t.Fatalf("pass = %+v", res)
	}
	vela, ok := readRow(t, s, 101)
	if !ok || !vela.public || vela.ruleset != "pvp" || vela.source != "profile" || vela.spec != "Fire" || vela.class != "mage" {
		t.Fatalf("public member = %+v", vela)
	}
	if vela.itemLevel == nil || *vela.itemLevel != 38 || vela.guildID == nil || *vela.guildID != 77 {
		t.Fatalf("public member item level or guild = %+v", vela)
	}
	if len(vela.equipment) != 1 || vela.equipment[0].ItemID != 9001 || vela.equipment[0].Level != 40 {
		t.Fatalf("equipment = %+v", vela.equipment)
	}
	if len(vela.professions) != 1 || vela.professions[0].Skill != 180 {
		t.Fatalf("professions = %+v", vela.professions)
	}
	orrin, ok := readRow(t, s, 102)
	if !ok || orrin.public || orrin.ruleset != "pvp" || orrin.source != "guild" || orrin.class != "warrior" {
		t.Fatalf("private member = %+v", orrin)
	}
	if _, ok := readRow(t, s, 103); ok {
		t.Fatal("an invalid character stayed in the directory")
	}

	// A second pass right away reads nothing: the roster and profiles are fresh.
	before := b.called("/data/wow/guild/ember-wardens/roster")
	if _, err := s.Pass(context.Background()); err != nil {
		t.Fatal(err)
	}
	if b.called("/data/wow/guild/ember-wardens/roster") != before || b.called("/profile/wow/character/vela-thornbrook/status") != 1 {
		t.Fatal("a fresh directory was read again")
	}
}

// A character whose name now answers with another id is a new character: the old row goes.
func TestPassReplacesACharacterWhoseNameMoved(t *testing.T) {
	pool := testPool(t)
	b := newBlizzard(t)
	s := b.service(pool)
	if _, err := pool.Exec(context.Background(),
		`insert into character_directory (region, bnet_character_id, name, slug, sources, roster_seen_at)
		 values ('us', 500, 'Vela Thornbrook', 'vela-thornbrook', '{roster}', now())`); err != nil {
		t.Fatal(err)
	}
	b.publicCharacter(501, "Vela Thornbrook", "Normal", 0, "", "")
	if _, err := s.Pass(context.Background()); err != nil {
		t.Fatal(err)
	}
	if _, ok := readRow(t, s, 500); ok {
		t.Fatal("the old character stayed")
	}
	if r, ok := readRow(t, s, 501); !ok || r.ruleset != "normal" {
		t.Fatalf("the new character = %+v, %v", r, ok)
	}
}

// A name the site knows only from an addon row is looked up once; a 404 is remembered so the
// next pass does not ask again, and a name that answers joins the directory with its source.
func TestPassSeedsNamesAndRemembersMisses(t *testing.T) {
	pool := testPool(t)
	b := newBlizzard(t)
	s := b.service(pool)
	for _, name := range []string{"Vela Thornbrook", "Nobody Here"} {
		if _, err := pool.Exec(context.Background(),
			`insert into characters (key, region, ruleset, name) values ($1, 'us', 'normal', $2)`,
			"us/normal/"+bnetSlug(name), name); err != nil {
			t.Fatal(err)
		}
	}
	b.publicCharacter(101, "Vela Thornbrook", "Normal", 0, "", "")
	res, err := s.Pass(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if res.Profiles != 1 || res.Misses != 1 {
		t.Fatalf("pass = %+v", res)
	}
	if r, ok := readRow(t, s, 101); !ok || len(r.sources) != 1 || r.sources[0] != SourceAddon {
		t.Fatalf("seeded row = %+v", r)
	}
	if _, err := s.Pass(context.Background()); err != nil {
		t.Fatal(err)
	}
	if n := b.called("/profile/wow/character/nobody-here/status"); n != 1 {
		t.Fatalf("a missed name was asked %d times, want 1", n)
	}
}

// Every row no source confirmed in 30 days is deleted, whatever else it holds.
func TestPassExpiresRowsPastThirtyDays(t *testing.T) {
	pool := testPool(t)
	b := newBlizzard(t)
	s := b.service(pool)
	if _, err := pool.Exec(context.Background(),
		`insert into character_directory (region, bnet_character_id, name, slug, refreshed_at, profile_refreshed_at)
		 values ('us', 1, 'Old Timer', 'old-timer', now() - interval '31 days', now()),
		        ('us', 2, 'New Comer', 'new-comer', now() - interval '29 days', now())`); err != nil {
		t.Fatal(err)
	}
	res, err := s.Pass(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if res.Expired != 1 {
		t.Fatalf("expired = %d, want 1", res.Expired)
	}
	if _, ok := readRow(t, s, 2); !ok {
		t.Fatal("a row inside 30 days was expired")
	}
}
