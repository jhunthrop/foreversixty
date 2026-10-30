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

func TestTheNightsReportTellsItsRaiders(t *testing.T) {
	w, _ := eventsWorld(t)
	raid := w.createRaid(t)
	got := w.h.svc.nightReadyEvents(context.Background(), raid.ID, "rep-1")
	if len(got) != 1 || got[0].Kind != events.ReportReady || got[0].Audience.Kind != events.ToRaidInLate ||
		got[0].Data["report_id"] != "rep-1" || got[0].Validate() != nil {
		t.Fatalf("night %+v", got)
	}
}
