package activity

import (
	"context"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/events"
)

// tierRow writes one tier_grade_weeks row.
func tierRow(t *testing.T, s *Store, week time.Time, spec string, band int, faction, grade string) {
	t.Helper()
	exec(t, s.Pool, `insert into tier_grade_weeks (week, spec, band, faction, role, grade, place, of_n)
		values ($1, $2, $3, $4, 'dps', $5, 1, 9)`, week, spec, band, faction, grade)
}

// mainOf seeds an account whose main is a level-60 alliance warrior simmed as spec.
func mainOf(t *testing.T, s *Store, name, spec string, level int) int64 {
	t.Helper()
	uid := user(t, s, name+"#1", false)
	key := "us/normal/" + name
	char(t, s, key, name, "Warrior", level, uid)
	exec(t, s.Pool, `update characters set faction = 'alliance' where key = $1`, key)
	exec(t, s.Pool, `update users set main_character_key = $1 where id = $2`, key, uid)
	exec(t, s.Pool, `insert into sims (id, user_id, spec, engine_version, lane, dps_mean, dps_error, iterations, result, state)
		values ($1, $2, $3, 'e', 'server', 500, 0, 3000, '{}', 'done')`, "s-"+name, uid, spec)
	exec(t, s.Pool, `insert into character_sims (sim_id, character_key, user_id, kind, spec, role, band, preset)
		values ($1, $2, $3, 'import', $4, 'dps', 60, 'raid')`, "s-"+name, key, uid, spec)
	return uid
}

// A spec whose grade moved since last week tells the accounts whose main is that spec at that
// band and faction, once per grade.
func TestTierMovesTellTheSpecsMains(t *testing.T) {
	s := testStore(t)
	week, last := WeekOf(testNow), PreviousWeek(testNow)
	tierRow(t, s, last, "warrior-fury", 60, "alliance", "A")
	tierRow(t, s, week, "warrior-fury", 60, "alliance", "S")
	tierRow(t, s, last, "warrior-arms", 60, "alliance", "B")
	tierRow(t, s, week, "warrior-arms", 60, "alliance", "B")
	fury := mainOf(t, s, "furyone", "warrior-fury", 60)
	mainOf(t, s, "armsone", "warrior-arms", 60)
	mainOf(t, s, "furylow", "warrior-fury", 44)
	pub := &recordingPublisher{}
	n, err := s.PublishTierMoves(context.Background(), pub, "https://foreversixty.gg", testNow)
	if err != nil || n != 1 || len(pub.got) != 1 {
		t.Fatalf("moves %d %v %+v", n, err, pub.got)
	}
	e := pub.got[0]
	if e.Kind != events.TierMoved || len(e.Audience.UserIDs) != 1 || e.Audience.UserIDs[0] != fury {
		t.Fatalf("event %+v", e)
	}
	if e.Title != "Fury Warrior moved to S" || e.Body != "Up from A last week, level 60, Alliance." ||
		e.Link != "https://foreversixty.gg/tiers" || e.Data["spec"] != "warrior-fury" {
		t.Fatalf("event words %+v", e)
	}
	// The next pass has nothing new to say.
	if n, err := s.PublishTierMoves(context.Background(), pub, "", testNow); err != nil || n != 0 || len(pub.got) != 1 {
		t.Fatalf("second pass %d %v", n, err)
	}
}
