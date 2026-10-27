package guilds

import (
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestRaidPostFactsCarryCountsNightAndKills(t *testing.T) {
	w, raid := matchedNight(t)
	w.h.svc.PublicBaseURL = "https://site.test"
	w.as(w.officer)
	view := w.getRaid(t, raid.ID)
	if view.Night == nil {
		t.Fatal("the matched night has no night figures")
	}

	facts, ok, err := w.h.svc.RaidPostFacts(t.Context(), raid.ID)
	if err != nil || !ok {
		t.Fatalf("RaidPostFacts = %v, %v", ok, err)
	}
	if facts.GuildID != w.gid || facts.Status != view.Status || !facts.EndsAt.Equal(raid.EndsAt()) {
		t.Fatalf("facts = %+v (view status %s)", facts, view.Status)
	}
	p := facts.Post
	if p.RaidID != raid.ID || p.GuildName != "Olympus" || p.Title != raid.Title || !p.StartsAt.Equal(raid.StartsAt) {
		t.Fatalf("post = %+v", p)
	}
	if p.Counts.In != view.Counts.In || p.Counts.Bench != view.Counts.Bench || p.Counts.Late != view.Counts.Late {
		t.Fatalf("counts = %+v, view %+v", p.Counts, view.Counts)
	}
	if p.Came != view.Night.Came || p.CameOf != view.Night.Signed || p.Came == 0 {
		t.Fatalf("came = %d of %d, night %+v", p.Came, p.CameOf, view.Night)
	}
	if view.Report == nil || len(view.Report.KillOrder) == 0 {
		t.Fatal("the matched night has no kill order")
	}
	if len(p.Kills) != 1 || p.Kills[0].Boss != view.Report.KillOrder[len(view.Report.KillOrder)-1].Name ||
		p.Kills[0].Duration != 5*time.Minute {
		t.Fatalf("kills = %+v", p.Kills)
	}
	if !strings.HasPrefix(facts.RaidURL, "https://site.test/guild/") || !strings.HasSuffix(facts.RaidURL, "/raids/"+strconv.FormatInt(raid.ID, 10)) {
		t.Fatalf("raid url = %q", facts.RaidURL)
	}
	if facts.NightURL != facts.RaidURL {
		t.Fatalf("night url = %q, want the done raid's page %q", facts.NightURL, facts.RaidURL)
	}
}

func TestRaidPostFactsOfAPlannedRaidAndAMissingOne(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.raidAt(t, raidZoneOny, time.Now().Add(48*time.Hour).UTC().Truncate(time.Minute))
	facts, ok, err := w.h.svc.RaidPostFacts(t.Context(), raid.ID)
	if err != nil || !ok {
		t.Fatalf("RaidPostFacts = %v, %v", ok, err)
	}
	if facts.Status != RaidPlanned || facts.Post.Came != 0 || facts.Post.Kills != nil || facts.NightURL != "" {
		t.Fatalf("planned facts = %+v", facts)
	}
	if _, ok, err := w.h.svc.RaidPostFacts(t.Context(), raid.ID+1000); err != nil || ok {
		t.Fatalf("missing raid = %v, %v", ok, err)
	}
}

func TestAStartChangeMarksTheMoveTickedOrNot(t *testing.T) {
	w := newRaidWorld(t, true)
	marker := &fakeMarker{}
	w.h.svc.RaidPosts = marker
	raid := w.createRaid(t)
	w.as(w.officer)
	if res := w.h.do(http.MethodPatch, w.path("/raids/%d", raid.ID), raidBody(raid.StartsAt.Format(time.RFC3339), 0)); res.StatusCode != http.StatusOK {
		t.Fatalf("patch = %d: %s", res.StatusCode, readBody(res))
	}
	moved := time.Now().Add(9 * 24 * time.Hour).UTC().Truncate(time.Minute).Format(time.RFC3339)
	if res := w.h.do(http.MethodPatch, w.path("/raids/%d", raid.ID), raidBody(moved, 0)); res.StatusCode != http.StatusOK {
		t.Fatalf("move = %d: %s", res.StatusCode, readBody(res))
	}
	was, ok := marker.moves[raid.ID]
	if !ok || !was.Equal(raid.StartsAt) || marker.count() != 2 {
		t.Fatalf("moves = %v (want %v), marks %v", marker.moves, raid.StartsAt, marker.ids)
	}
}
