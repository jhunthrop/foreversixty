package guilds

import (
	"context"
	"net/http"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/events"
)

func TestApprovingACharacterTellsItsAccount(t *testing.T) {
	w, pub := eventsWorld(t)
	w.as(w.officer)
	res := w.h.do(http.MethodPost, w.path("/characters/%s/approve", memberKeyC), "")
	if res.StatusCode != http.StatusOK {
		t.Fatalf("approve %d %s", res.StatusCode, readBody(res))
	}
	got := pub.waitFor(t, events.GuildApproved, 1)[0]
	if got.Audience.UserIDs[0] != w.member || got.Title != "Approved to Olympus" ||
		got.Body != "An officer approved Unverifiedalt to the roster." || !strings.Contains(got.Link, "/guild/") {
		t.Fatalf("approved %+v", got)
	}
}

func TestALootAwardTellsTheWinner(t *testing.T) {
	w, pub := eventsWorld(t)
	w.h.svc.publishLootAwarded(context.Background(), w.gid, 41, memberKeyA, 17078)
	got := pub.waitFor(t, events.LootAwarded, 1)[0]
	if got.Audience.UserIDs[0] != w.member || got.DedupeKey != "loot.awarded:41" || !strings.HasPrefix(got.Body, "Kraggor was awarded") {
		t.Fatalf("loot %+v", got)
	}
	// An award to no roster character tells nobody.
	w.h.svc.publishLootAwarded(context.Background(), w.gid, 42, "invite:99", 17078)
	pub.waitFor(t, events.LootAwarded, 1)
}

func TestTheNightsSummaryTellsEachRaiderTheirOwnNight(t *testing.T) {
	w, _ := eventsWorld(t)
	raid := w.createRaid(t)
	w.as(w.member)
	w.signup(t, raid.ID, memberKeyA, `{"state":"in"}`)
	w.as(w.member2)
	w.signup(t, raid.ID, member2Key, `{"state":"out"}`)
	exec := func(q string, args ...any) {
		if _, err := w.h.pool.Exec(context.Background(), q, args...); err != nil {
			t.Fatal(err)
		}
	}
	exec(`insert into reports (id, owner_id, title, visibility, status) values ('rep-1', $1, 'Ony', 'guild', 'complete')`, w.officer)
	exec(`insert into loot_awards (guild_id, encounter_id, item_id, character_key, report_id) values ($1, 1084, 17078, $2, 'rep-1')`, w.gid, memberKeyA)
	got := w.h.svc.summaryEvents(context.Background(), raid.ID, "rep-1")
	if len(got) != 1 {
		t.Fatalf("summaries %+v", got)
	}
	e := got[0]
	if e.Kind != events.RaidSummary || e.Audience.UserIDs[0] != w.member || e.DedupeKey != "raid.summary:rep-1" ||
		!strings.HasPrefix(e.Body, summaryCopy.Report) || e.Data["loot_item_ids"] != "17078" || e.Validate() != nil {
		t.Fatalf("summary %+v", e)
	}
	if !strings.Contains(e.Body, "You won") {
		t.Fatalf("no loot line: %q", e.Body)
	}
}
