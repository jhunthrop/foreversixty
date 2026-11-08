package guilds

import (
	"context"
	"net/http"
	"slices"
	"strconv"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/contract"
)

// fakeCoach answers scripted notes for every night.
type fakeCoach struct{ notes []string }

func (f fakeCoach) CoachNotes(context.Context, int64, string, string) ([]string, error) {
	return f.notes, nil
}

// raidNightOf reads the raid night as uid; status is the answer's code.
func (w raidWorld) raidNightOf(t *testing.T, uid, raidID int64) (RaidNightSummary, int) {
	t.Helper()
	w.as(uid)
	res := w.h.do(http.MethodGet, "/v1/me/raid-nights/"+strconv.FormatInt(raidID, 10), "")
	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		return RaidNightSummary{}, res.StatusCode
	}
	var v RaidNightSummary
	w.h.data(res, &v)
	return v, res.StatusCode
}

// Before the pull: the countdown's times, your answer with its minutes, who is coming by role
// with the late notes, and the ready check; nothing from a log yet.
func TestRaidNightBeforeThePull(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	w.as(w.member)
	w.putSignup(raid.ID, memberKeyA, `{"state":"late","late_minutes":15,"note":"work runs late"}`)
	w.as(w.member2)
	w.putSignup(raid.ID, member2Key, `{"state":"in","role":"healer"}`)

	v, code := w.raidNightOf(t, w.member, raid.ID)
	if code != http.StatusOK {
		t.Fatalf("raid night = %d", code)
	}
	if v.Phase != RaidNightBefore || v.ZoneName != "Onyxia's Lair" || !v.EndsAt.Equal(v.StartsAt.Add(180*time.Minute)) {
		t.Fatalf("head %+v", v)
	}
	if v.Character.CharacterKey != memberKeyA || v.You.State == nil || *v.You.State != SignupLate ||
		v.You.LateMinutes == nil || *v.You.LateMinutes != 15 {
		t.Fatalf("you %+v / %+v", v.Character, v.You)
	}
	if v.Coming.In != 1 || v.Coming.Late != 1 || len(v.Coming.Healers) != 1 || v.Coming.Healers[0].Name != "Mirelle" {
		t.Fatalf("coming %+v", v.Coming)
	}
	late := slices.IndexFunc(v.Coming.Damage, func(p RaidNightPerson) bool { return p.Name == "Kraggor" })
	if late < 0 || v.Coming.Damage[late].State != SignupLate || v.Coming.Damage[late].Note != "work runs late" {
		t.Fatalf("damage %+v", v.Coming.Damage)
	}
	if v.Ready == nil || v.ReportID != nil || v.Kills != nil || v.Wipes != nil || len(v.Loot) != 0 || v.CoachNotes != nil {
		t.Fatalf("log fields before the pull %+v", v)
	}
	w.as(w.member)
	contract.AssertFixture(t, "me.raid-night.before",
		guildBody(t, w.h.do(http.MethodGet, "/v1/me/raid-nights/"+strconv.FormatInt(raid.ID, 10), ""), http.StatusOK), contract.Options{})
}

// After the log: kills and wipes, the loot you got, your bests on its kills and how many beat
// your earlier ones, and the coach teaser; the ready check is gone.
func TestRaidNightAfterTheLog(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	w.as(w.member)
	w.putSignup(raid.ID, memberKeyA, `{"state":"in"}`)
	night := time.Now().Add(-20 * time.Hour).UTC().Truncate(time.Minute)
	me := seedPlayer{key: memberKeyA, name: "Kraggor", class: "warrior", spec: "Fury", role: "dps"}
	seedNight(t, w.h.pool, w.gid, "night-old", "Onyxia's Lair", night.Add(-7*24*time.Hour),
		[]seedFight{{encounterID: 1084, kill: true, duration: 5 * time.Minute, players: []seedPlayer{me}}})
	seedNight(t, w.h.pool, w.gid, "night-new", "Onyxia's Lair", night, []seedFight{
		{encounterID: 1084, kill: false, duration: 2 * time.Minute, players: []seedPlayer{me}},
		{encounterID: 1084, kill: true, offset: 10 * time.Minute, duration: 4 * time.Minute, players: []seedPlayer{me}},
	})
	ctx := context.Background()
	exec := func(sql string, args ...any) {
		t.Helper()
		if _, err := w.h.pool.Exec(ctx, sql, args...); err != nil {
			t.Fatal(err)
		}
	}
	exec(`update fight_metrics set metric_dps = 80 where report_id = 'night-old'`)
	exec(`update fight_metrics set metric_dps = 712.4 where report_id = 'night-new' and kill`)
	exec(`update guild_raids set report_id = 'night-new', status = 'done', starts_at = $2 where id = $1`, raid.ID, night)
	exec(`insert into loot_awards (guild_id, encounter_id, item_id, character_key, awarded_by, report_id)
		values ($1, 1084, 17078, $2, $3, 'night-new')`, w.gid, memberKeyA, w.officer)
	w.h.svc.Coach = fakeCoach{notes: []string{"Bloodthirst: 9 casts lost."}}

	v, code := w.raidNightOf(t, w.member, raid.ID)
	if code != http.StatusOK {
		t.Fatalf("raid night = %d", code)
	}
	if v.Phase != RaidNightAfter || v.ReportID == nil || *v.ReportID != "night-new" || v.Ready != nil {
		t.Fatalf("head %+v", v)
	}
	if v.Kills == nil || *v.Kills != 1 || v.Wipes == nil || *v.Wipes != 1 {
		t.Fatalf("kills %v wipes %v", v.Kills, v.Wipes)
	}
	if len(v.Bests) != 1 || v.Bests[0].DPS != 712.4 || v.Bests[0].KillMs != 240000 || v.NewBests == nil || *v.NewBests != 1 {
		t.Fatalf("bests %+v new %v", v.Bests, v.NewBests)
	}
	if len(v.Loot) != 1 || v.Loot[0].ItemID != 17078 || v.Loot[0].AwardedBy != "Thornhide" {
		t.Fatalf("loot %+v", v.Loot)
	}
	if len(v.CoachNotes) != 1 {
		t.Fatalf("coach %+v", v.CoachNotes)
	}
	w.as(w.member)
	contract.AssertFixture(t, "me.raid-night",
		guildBody(t, w.h.do(http.MethodGet, "/v1/me/raid-nights/"+strconv.FormatInt(raid.ID, 10), ""), http.StatusOK), contract.Options{})
}

func TestRaidNightIsForTheGuildsMembers(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	stranger := seedUser(t, w.h.pool, "stranger@example.com")
	if _, code := w.raidNightOf(t, stranger, raid.ID); code != http.StatusNotFound {
		t.Fatalf("stranger = %d", code)
	}
	if _, code := w.raidNightOf(t, w.member, 999999); code != http.StatusNotFound {
		t.Fatalf("no such raid = %d", code)
	}
}
