package activity

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func report(t *testing.T, s *Store, id string, owner int64, visibility, zone, logger string) {
	t.Helper()
	exec(t, s.Pool, `insert into reports (id, owner_id, visibility, zone, status, logging_character, completed_at)
		values ($1, $2, $3, $4, 'complete', $5, $6)`, id, owner, visibility, zone, logger, testNow)
	for i, f := range []struct {
		enc   int
		boss  string
		start int64
	}{{100, "Rethilgore", 0}, {101, "Arugal", 1_200_000}, {0, "", 2_400_000}} {
		var enc any
		if f.enc > 0 {
			enc = f.enc
		}
		exec(t, s.Pool, `insert into fights (report_id, fight_index, encounter_id, name, kill, start_ms, duration_ms)
			values ($1, $2, $3, $4, $5, $6, 60000)`, id, i, enc, f.boss, f.enc > 0, f.start)
	}
}

func rowTexts(t *testing.T, s *Store) []string {
	t.Helper()
	p := read(t, s, Query{Filter: FilterAll})
	out := []string{}
	for _, r := range p.Rows {
		var b strings.Builder
		for _, sg := range r.Segments {
			b.WriteString(sg.Text)
		}
		out = append(out, b.String())
	}
	return out
}

// Ruling P1, P2: a public report names its logger; an anonymised owner's report does not; a
// private one writes nothing; a raid's first kill writes once.
func TestReportWriterWritesPublicUploadsAndFirstKills(t *testing.T) {
	s := testStore(t)
	w := &ReportWriter{Store: s, RaidZones: map[string]bool{"Onyxia's Lair": true}}
	named := user(t, s, "a#1", false)
	hidden := user(t, s, "b#1", true)
	char(t, s, "us/normal/nightfall", "Nightfall", "Warrior", 24, named)
	char(t, s, "us/normal/quiet", "Quiet", "Mage", 24, hidden)
	report(t, s, "r1", named, "public", "Shadowfang Keep", "us/normal/nightfall")
	report(t, s, "r2", hidden, "public", "Shadowfang Keep", "us/normal/quiet")
	report(t, s, "r3", named, "private", "Shadowfang Keep", "us/normal/nightfall")
	report(t, s, "r4", named, "public", "Onyxia's Lair", "us/normal/nightfall")
	report(t, s, "r5", named, "public", "Onyxia's Lair", "us/normal/nightfall")
	for _, id := range []string{"r1", "r2", "r3", "r4", "r5", "r1"} {
		w.ReportCompleted(context.Background(), id)
	}
	got := strings.Join(rowTexts(t, s), "\n")
	for _, want := range []string{
		"Nightfall uploaded Shadowfang Keep, 2 bosses in 41 min",
		"Shadowfang Keep uploaded, 2 bosses in 41 min",
		"Nightfall downed Rethilgore in Onyxia's Lair, the first kill on Forever",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("missing %q in\n%s", want, got)
		}
	}
	if strings.Count(got, "first kill") != 2 || strings.Count(got, "uploaded") != 4 || strings.Contains(got, "Quiet") {
		t.Fatalf("rows:\n%s", got)
	}
}

// Ruling P3: one row per raid, its counts rewritten as members sign up, for the guild only.
func TestGuildRaidWriterKeepsOneRowWithItsCounts(t *testing.T) {
	s := testStore(t)
	u := user(t, s, "o#1", false)
	g := guild(t, s, "Olympus XXVII")
	var raid int64
	if err := s.Pool.QueryRow(context.Background(), `insert into guild_raids (guild_id, zone, title, starts_at, duration_minutes, created_by)
		values ($1, 'raid:onyxias-lair', 'Onyxia''s Lair', $2, 120, $3) returning id`, g, testNow.Add(48*time.Hour), u).Scan(&raid); err != nil {
		t.Fatal(err)
	}
	w := &GuildRaidWriter{Store: s}
	w.RaidChanged(context.Background(), raid)
	exec(t, s.Pool, `insert into guild_raid_signups (raid_id, character_key, name, class, spec, role, state, signed_by, acted_by)
		values ($1, 'us/normal/a', 'A', 'warrior', 'fury', 'dps', 'in', $2, $2), ($1, 'us/normal/b', 'B', 'mage', 'fire', 'dps', 'late', $2, $2)`, raid, u)
	w.RaidChanged(context.Background(), raid)
	p := read(t, s, Query{Filter: FilterGuild, GuildIDs: []int64{g}})
	if len(p.Rows) != 1 || p.Rows[0].Segments[3].Text != ", Mon 9 Nov: 1 in, 1 late" || *p.Rows[0].Icon != "inv_misc_head_dragon_black" {
		t.Fatalf("rows: %+v", p.Rows)
	}
	if len(read(t, s, Query{Filter: FilterAll}).Rows) != 0 {
		t.Fatal("a guild row reached the public feed")
	}
}

func writeBis(t *testing.T, dir, generated, chest, belt string) {
	t.Helper()
	body := `{"spec":"warrior-fury","generated_at":"` + generated + `","bands":[` +
		`{"band":20,"preset":"bare","faction":"horde","set_dps":30,"slots":[{"slot":"chest","item_id":` + chest + `,"item_name":"Chest ` + chest + `"},{"slot":"waist","item_id":` + belt + `,"item_name":"Belt ` + belt + `"}]},` +
		`{"band":20,"preset":"bare","faction":"alliance","set_dps":30,"slots":[{"slot":"chest","item_id":` + chest + `,"item_name":"Chest ` + chest + `"}]}]}`
	path := filepath.Join(dir, "b1", "bis")
	if err := os.MkdirAll(path, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(path, "warrior-fury.json"), []byte(body), 0o600); err != nil {
		t.Fatal(err)
	}
}

// Ruling P5: the first look seeds, one change names the pick, two group into one row.
func TestDiffBisSeedsThenTellsChanges(t *testing.T) {
	s := testStore(t)
	dir := t.TempDir()
	steps := []struct {
		gen, chest, belt string
		rows             int
	}{{"2026-11-05T10:00:00Z", "1", "2", 0}, {"2026-11-06T10:00:00Z", "3", "2", 1}, {"2026-11-07T10:00:00Z", "4", "5", 1}, {"2026-11-07T10:00:00Z", "4", "5", 0}}
	for i, st := range steps {
		writeBis(t, dir, st.gen, st.chest, st.belt)
		n, err := s.DiffBis(context.Background(), dir, "b1")
		if err != nil || n != st.rows {
			t.Fatalf("step %d: %d rows, %v", i, n, err)
		}
	}
	got := rowTexts(t, s)
	if len(got) != 2 || got[0] != "Best in slot refreshed for Fury Warrior 20 to 29: 2 new picks" ||
		got[1] != "Best in slot changed for Fury Warrior 20 to 29: Chest 3 is the new chest" {
		t.Fatalf("rows: %q", got)
	}
}
