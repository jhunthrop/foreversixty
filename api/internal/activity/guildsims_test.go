package activity

import (
	"context"
	"strings"
	"testing"
	"time"
)

func sim(t *testing.T, s *Store, id, key string, userID int64, kind string, dps float64, band int, preset string, at time.Time) {
	t.Helper()
	exec(t, s.Pool, `insert into sims (id, user_id, spec, engine_version, lane, dps_mean, dps_error, iterations, result, state, created_at, finished_at)
		values ($1, $2, 'warrior-fury', 'e', 'server', $3, 0, 3000, '{}', 'done', $4, $4)`, id, userID, dps, at)
	exec(t, s.Pool, `insert into character_sims (sim_id, character_key, user_id, kind, spec, role, band, preset, created_at)
		values ($1, $2, $3, $4, 'warrior-fury', 'dps', $5, $6, $7)`, id, key, userID, kind, band, preset, at)
}

// guildOf makes a guild of n warriors who each simmed, best first by name order: m0 simmed
// the most.
func guildOf(t *testing.T, s *Store, n int) (int64, []int64) {
	t.Helper()
	g := guild(t, s, "Olympus XXVII")
	users := make([]int64, n)
	for i := 0; i < n; i++ {
		name := "M" + string(rune('a'+i))
		key := "us/normal/" + strings.ToLower(name)
		users[i] = user(t, s, name+"#1", false)
		char(t, s, key, name, "Warrior", 60, users[i])
		member(t, s, g, key, users[i])
		sim(t, s, "s"+name, key, users[i], "import", float64(600-10*i), 60, "raid", testNow.Add(-time.Duration(i+1)*time.Hour))
	}
	return g, users
}

// Ruling P7: a member's sim posts with its place from 5 members on, never to another guild.
func TestGuildSimWriterPostsTheMembersPlace(t *testing.T) {
	s := testStore(t)
	g, users := guildOf(t, s, 5)
	w := &GuildSimWriter{Store: s, Now: func() time.Time { return testNow }}
	w.SimFinished(context.Background(), "sMc")
	p := read(t, s, Query{Filter: FilterGuild, GuildIDs: []int64{g}})
	if len(p.Rows) != 1 {
		t.Fatalf("rows: %+v", p.Rows)
	}
	var b strings.Builder
	for _, sg := range p.Rows[0].Segments {
		b.WriteString(sg.Text)
	}
	if want := "Mc simmed 580 DPS, #3 of 5 DPS in Olympus XXVII (Fury Warrior, 60 raid)"; b.String() != want {
		t.Fatalf("sentence %q, want %q", b.String(), want)
	}
	// The same figure again posts nothing new.
	sim(t, s, "sMc2", "us/normal/mc", users[2], "import", 581, 60, "raid", testNow)
	w.SimFinished(context.Background(), "sMc2")
	if n := len(read(t, s, Query{Filter: FilterGuild, GuildIDs: []int64{g}}).Rows); n != 1 {
		t.Fatalf("a 0.2%% change posted: %d rows", n)
	}
	if len(read(t, s, Query{Filter: FilterAll}).Rows) != 0 {
		t.Fatal("a guild sim reached the public feed")
	}
}

// A sim that measured nothing (0.0) posts no row, moves nobody and leaves posted_value null.
func TestGuildSimWriterSkipsASimThatMeasuredNothing(t *testing.T) {
	s := testStore(t)
	g, users := guildOf(t, s, 5)
	w := &GuildSimWriter{Store: s, Now: func() time.Time { return testNow }}
	sim(t, s, "sZero", "us/normal/mc", users[2], "import", 0, 60, "raid", testNow)
	w.SimFinished(context.Background(), "sZero")
	if n := len(read(t, s, Query{Filter: FilterGuild, GuildIDs: []int64{g}}).Rows); n != 0 {
		t.Fatalf("a 0.0 sim posted: %d rows", n)
	}
	var posted *float64
	if err := s.Pool.QueryRow(context.Background(), `select posted_value from character_sims where sim_id = 'sZero'`).Scan(&posted); err != nil || posted != nil {
		t.Fatalf("posted_value %v %v, want null", posted, err)
	}
}

// A member below the first band plays against its list without being in it: the row names the
// level, never "20 to 29".
func TestGuildSimWriterNamesALevelBelowTheFirstBand(t *testing.T) {
	s := testStore(t)
	g, _ := guildOf(t, s, 5)
	low := user(t, s, "Low#1", false)
	char(t, s, "us/normal/pump", "Pump", "Paladin", 13, low)
	member(t, s, g, "us/normal/pump", low)
	sim(t, s, "sPump", "us/normal/pump", low, "import", 41.5, 20, "bare", testNow)
	w := &GuildSimWriter{Store: s, Now: func() time.Time { return testNow }}
	w.SimFinished(context.Background(), "sPump")
	p := read(t, s, Query{Filter: FilterGuild, GuildIDs: []int64{g}})
	if len(p.Rows) != 1 {
		t.Fatalf("rows: %+v", p.Rows)
	}
	var b strings.Builder
	for _, sg := range p.Rows[0].Segments {
		b.WriteString(sg.Text)
	}
	if want := "Pump simmed 41.5 DPS in Olympus XXVII (Fury Warrior, level 13)"; b.String() != want {
		t.Fatalf("sentence %q, want %q", b.String(), want)
	}
}

// Ruling R2: hidden below 3 members; anonymised members and old sims drop out.
func TestGuildLadderRules(t *testing.T) {
	s := testStore(t)
	g, users := guildOf(t, s, 4)
	exec(t, s.Pool, `update users set anonymize = true where id = $1`, users[0])
	exec(t, s.Pool, `update sims set finished_at = $1 where id = 'sMb'`, testNow.Add(-31*24*time.Hour))
	ladder, err := s.GuildLadder(context.Background(), LadderGroup{g, "dps", 60, "raid"}, testNow)
	if err != nil || len(ladder) != 2 || ladder[0].Name != "Mc" || ladder[0].Rank != 1 {
		t.Fatalf("ladder %+v %v", ladder, err)
	}
	h := &Home{Store: s}
	answer, err := h.answer(context.Background(), homeCharacter{key: "us/normal/mc", spec: "warrior-fury", level: 60, faction: "horde", userID: users[2]}, testNow)
	if err != nil || answer.GuildSims != nil {
		t.Fatalf("a 2-member ladder showed: %+v %v", answer.GuildSims, err)
	}
}
