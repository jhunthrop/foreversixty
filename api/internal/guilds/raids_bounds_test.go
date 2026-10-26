package guilds

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"testing"
	"time"
)

const farStartMessage = "Pick a date within the next year."

func TestCreateRaidsCapCountsOpenRaidsOnly(t *testing.T) {
	f := newRaidFixture(t)
	ctx := context.Background()
	for i := 0; i < RaidCap; i++ {
		f.createRaid(t, time.Now().Add(-time.Duration(48+i)*time.Hour)) // long ended, still planned
	}
	if _, err := f.store.CreateRaid(ctx, f.input(time.Now().Add(24*time.Hour))); err != nil {
		t.Fatalf("ended raids counted toward the cap: %v", err)
	}
	g := newRaidFixture(t)
	for i := 0; i < RaidCap; i++ {
		g.createRaid(t, time.Now().Add(time.Duration(24+i)*time.Hour))
	}
	if _, err := g.store.CreateRaid(ctx, g.input(time.Now().Add(24*time.Hour))); !errors.Is(err, ErrRaidCap) {
		t.Fatalf("12 open raids + one: err = %v, want ErrRaidCap", err)
	}
}

func TestRaidStartHasAYearBound(t *testing.T) {
	w := newRaidWorld(t, true)
	w.as(w.officer)
	far := time.Now().Add(maxRaidAhead + 48*time.Hour).UTC().Format(time.RFC3339)
	e := expectError(t, w.h.do(http.MethodPost, w.path("/raids"), raidBody("9999-01-01T00:00:00Z", 0)), http.StatusUnprocessableEntity, "invalid")
	if e.Fields["starts_at"] != farStartMessage {
		t.Fatalf("fields = %v", e.Fields)
	}
	expectError(t, w.h.do(http.MethodPost, w.path("/raids"), raidBody(far, 0)), http.StatusUnprocessableEntity, "invalid")
	lastWeek := time.Now().Add(maxRaidAhead - 10*24*time.Hour).UTC().Format(time.RFC3339)
	e = expectError(t, w.h.do(http.MethodPost, w.path("/raids"), raidBody(lastWeek, 3)), http.StatusUnprocessableEntity, "invalid")
	if e.Fields["starts_at"] != farStartMessage {
		t.Fatalf("repeat fields = %v", e.Fields)
	}
	ok := time.Now().Add(maxRaidAhead - 40*24*time.Hour).UTC().Format(time.RFC3339)
	if got := w.createRaids(t, raidBody(ok, 3)); len(got) != 4 {
		t.Fatalf("created %d", len(got))
	}
}

func TestControlCharactersAreRefused(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	post := func(title, notes string) *http.Response {
		w.as(w.officer)
		body, err := json.Marshal(map[string]any{"zone": raidZoneOny, "title": title, "starts_at": futureStart(),
			"duration_minutes": 180, "notes": notes, "timezone": "America/Chicago"})
		if err != nil {
			t.Fatal(err)
		}
		return w.h.do(http.MethodPost, w.path("/raids"), string(body))
	}
	for _, tc := range []struct{ name, title, notes, field string }{
		{"title nul", "a\u0000b", "", "title"},
		{"title newline", "a\nb", "", "title"},
		{"notes nul", "", "a\u0000b", "notes"},
		{"notes escape", "", "a\u001bb", "notes"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			e := expectError(t, post(tc.title, tc.notes), http.StatusUnprocessableEntity, "invalid")
			if _, ok := e.Fields[tc.field]; !ok {
				t.Fatalf("fields = %v", e.Fields)
			}
		})
	}
	res := post("", "line one\n\tline two")
	if res.StatusCode != http.StatusCreated {
		t.Fatalf("newline and tab in notes = %d: %s", res.StatusCode, readBody(res))
	}
	w.as(w.member)
	body, err := json.Marshal(map[string]any{"state": "in", "note": "a\u0000b"})
	if err != nil {
		t.Fatal(err)
	}
	e := expectError(t, w.putSignup(raid.ID, memberKeyA, string(body)), http.StatusUnprocessableEntity, "invalid")
	if _, ok := e.Fields["note"]; !ok {
		t.Fatalf("fields = %v", e.Fields)
	}
}

func TestOfficerRaidRoutesAreLimitedPerAccount(t *testing.T) {
	w := newRaidWorld(t, true)
	w.as(w.officer)
	for i := 1; i <= officerRaidWritesPerMinute; i++ {
		res := w.h.do(http.MethodPost, w.path("/raids"), `{}`)
		res.Body.Close()
		if res.StatusCode != http.StatusUnprocessableEntity {
			t.Fatalf("call %d = %d", i, res.StatusCode)
		}
	}
	expectError(t, w.h.do(http.MethodPost, w.path("/raids/1/lock"), ""), http.StatusTooManyRequests, "rate_limited")
	expectError(t, w.h.do(http.MethodPatch, w.path("/raids/1"), `{}`), http.StatusTooManyRequests, "rate_limited")
}
