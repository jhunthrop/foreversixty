package guilds

import (
	"net/http"
	"testing"
	"time"
)

// recruitingOf reads the recruiting list for a query.
func (w raidWorld) recruitingOf(t *testing.T, query string) []RecruitingGuild {
	t.Helper()
	w.as(0)
	res := w.h.do(http.MethodGet, "/v1/guilds/recruiting?"+query, "")
	if res.StatusCode != http.StatusOK {
		t.Fatalf("recruiting = %d: %s", res.StatusCode, readBody(res))
	}
	var out struct {
		Rows []RecruitingGuild `json:"rows"`
	}
	w.h.data(res, &out)
	return out.Rows
}

func (w raidWorld) recruit(t *testing.T, gid int64, open bool, faction, nights, wanted string) {
	t.Helper()
	if _, err := w.h.pool.Exec(t.Context(), `update guilds set recruit_open = $2, faction = $3, recruit_nights = $4,
		recruit_wanted = $5::jsonb, raid_timezone = 'America/Chicago' where id = $1`, gid, open, faction, nights, wanted); err != nil {
		t.Fatal(err)
	}
}

// A guild shows when it recruits, on the realm and faction asked, and its wanted list names the
// spec, the class with no spec, or nothing (open to all); needs is the spec's role.
func TestRecruitingListsTheGuildsThatWantTheSpec(t *testing.T) {
	w := newRaidWorld(t, true)
	w.recruit(t, w.gid, true, "alliance", "Wed and Sat, 8:00 pm CST", `[{"class":"warrior","spec":"warrior-fury","note":""}]`)
	other := seedGuild(t, w.h.pool, "Wants Priests")
	w.recruit(t, other, true, "alliance", "", `[{"class":"priest","spec":null,"note":""}]`)
	anyone := seedGuild(t, w.h.pool, "Takes Anyone")
	w.recruit(t, anyone, true, "alliance", "", `[]`)
	closed := seedGuild(t, w.h.pool, "Closed")
	w.recruit(t, closed, false, "alliance", "", `[]`)
	horde := seedGuild(t, w.h.pool, "Horde Side")
	w.recruit(t, horde, true, "horde", "", `[]`)

	rows := w.recruitingOf(t, "spec=warrior-fury&faction=alliance&region=us&ruleset=hardcore")
	names := map[string]RecruitingGuild{}
	for _, r := range rows {
		names[r.Name] = r
	}
	if len(rows) != 2 || names["Olympus"].GuildID != w.gid || names["Takes Anyone"].GuildID != anyone {
		t.Fatalf("rows %+v", rows)
	}
	if o := names["Olympus"]; o.Schedule != "Wed and Sat, 8:00 pm CST" || o.Needs != "damage" || o.Faction != "alliance" {
		t.Fatalf("Olympus %+v", o)
	}
	if rows := w.recruitingOf(t, "spec=priest-holy&faction=alliance&region=us&ruleset=hardcore"); len(rows) != 2 {
		t.Fatalf("priest rows %+v", rows)
	} else if rows[0].Needs != "healer" {
		t.Fatalf("priest needs %q", rows[0].Needs)
	}
}

// With no words of its own, the schedule is the planned raids' nights and time in the raid zone.
func TestRecruitingScheduleFromThePlannedRaids(t *testing.T) {
	w := newRaidWorld(t, true)
	w.recruit(t, w.gid, true, "alliance", "", `[]`)
	chicago, _ := time.LoadLocation("America/Chicago")
	now := time.Now().In(chicago)
	next := func(day time.Weekday) string {
		d := time.Date(now.Year(), now.Month(), now.Day(), 20, 0, 0, 0, chicago)
		for d.Weekday() != day || !d.After(now) {
			d = d.AddDate(0, 0, 1)
		}
		return d.UTC().Format(time.RFC3339)
	}
	w.createRaids(t, raidBody(next(time.Saturday), 0))
	w.createRaids(t, raidBody(next(time.Wednesday), 0))
	rows := w.recruitingOf(t, "spec=warrior-fury&faction=alliance&region=us&ruleset=hardcore")
	if len(rows) != 1 || (rows[0].Schedule != "Wed and Sat, 8 pm CST" && rows[0].Schedule != "Wed and Sat, 8 pm CDT") {
		t.Fatalf("rows %+v", rows)
	}
}

func TestRecruitingValidatesItsQuery(t *testing.T) {
	w := newRaidWorld(t, true)
	for _, q := range []string{"spec=nope&faction=alliance", "spec=warrior-fury&faction=goblin", "faction=alliance"} {
		res := w.h.do(http.MethodGet, "/v1/guilds/recruiting?"+q, "")
		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("%s = %d", q, res.StatusCode)
		}
	}
}
