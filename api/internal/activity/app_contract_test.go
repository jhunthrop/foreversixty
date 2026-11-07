package activity

import (
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/contract"
)

// segmentOptional are the segment fields Go omits on a plain run.
var segmentOptional = []string{"data.rows[].segments[].style", "data.rows[].segments[].class",
	"data.rows[].segments[].quality", "data.rows[].segments[].item_id"}

func TestAppContractFeed(t *testing.T) {
	s := testStore(t)
	e := ev(KindLogUpload, "upload", testNow.Add(-time.Minute))
	e.Segments = []Segment{Name("Nightfall", "warlock"), Text(" uploaded Shadowfang Keep, 8 bosses in 41 min")}
	mustAppend(t, s, e)
	b := ev(KindBisChange, "bis", testNow.Add(-2*time.Minute))
	b.Segments = []Segment{Item("Blackened Defias Armor", 10399, 3), Text(" is the new chest")}
	mustAppend(t, s, b)
	rec, _ := serve(t, s, nil, 0, "/v1/feed?limit=1")
	contract.AssertFixture(t, "feed", rec.Body.Bytes(), contract.Options{Optional: segmentOptional})
	rec, _ = serve(t, s, nil, 0, "/v1/feed?filter=logs&cursor=")
	contract.AssertFixture(t, "feed", rec.Body.Bytes(), contract.Options{Optional: segmentOptional})
	empty := testStore(t)
	rec, _ = serve(t, empty, nil, 0, "/v1/feed")
	contract.AssertFixture(t, "feed.empty", rec.Body.Bytes(), contract.Options{})
	rec, _ = serve(t, empty, nil, 0, "/v1/feed?filter=nope")
	contract.AssertFixture(t, "feed.error-invalid", rec.Body.Bytes(), contract.Options{})
}

func TestAppContractHome(t *testing.T) {
	s := testStore(t)
	g, users := guildOf(t, s, 3)
	week := WeekOf(testNow)
	exec(t, s.Pool, `update characters set faction = 'horde'`)
	exec(t, s.Pool, `insert into tier_grade_weeks (week, spec, band, faction, role, grade, place, of_n) values ($1, 'warrior-fury', 60, 'horde', 'dps', 'A', 4, 23)`, week)
	exec(t, s.Pool, `insert into spec_board_weeks (week, spec, band, character_key, name, class, rank, score) values ($1, 'warrior-fury', 60, 'k1', 'Grukk', 'warrior', 3, 80)`, week)
	_ = g
	rec, _ := serve(t, s, nil, users[0], "/v1/home?character=us/normal/ma")
	contract.AssertFixture(t, "home", rec.Body.Bytes(), contract.Options{})

	u := user(t, s, "solo#1", false)
	char(t, s, "us/normal/solo", "Solo", "Warrior", 20, u)
	exec(t, s.Pool, `delete from tier_grade_weeks`)
	rec, _ = serve(t, s, nil, u, "/v1/home?character=us/normal/solo")
	contract.AssertFixture(t, "home.empty", rec.Body.Bytes(), contract.Options{})
}
