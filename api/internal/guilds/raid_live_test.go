package guilds

import (
	"context"
	"slices"
	"sync"
	"testing"
)

// fakeLiveRaids records which raids started and ended.
type fakeLiveRaids struct {
	mu             sync.Mutex
	started, ended []int64
}

func (f *fakeLiveRaids) RaidStarting(_ context.Context, raidID int64) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.started = append(f.started, raidID)
}

func (f *fakeLiveRaids) RaidEnded(_ context.Context, raidID int64) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.ended = append(f.ended, raidID)
}

// The attendees are one row per account In or Late, by its best answer.
func TestRaidAttendees(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	w.as(w.member)
	w.putSignup(raid.ID, memberKeyA, `{"state":"late"}`)
	w.as(w.member2)
	w.putSignup(raid.ID, member2Key, `{"state":"out"}`)
	got, err := w.h.svc.RaidAttendees(context.Background(), raid.ID)
	if err != nil || len(got) != 1 || got[0].UserID != w.member || got[0].CharacterKey != memberKeyA || got[0].Name != "Kraggor" {
		t.Fatalf("attendees %+v %v", got, err)
	}
}

// raid.starting starts the raid's live surfaces; a hand match ends them.
func TestRaidLiveHooks(t *testing.T) {
	w, pub := eventsWorld(t)
	live := &fakeLiveRaids{}
	w.h.svc.LiveRaids = live
	raid := w.createRaid(t)
	pub.waitFor(t, "raid.posted", 1)
	if _, err := w.h.svc.RunEventSchedule(context.Background(), raid.StartsAt.Add(-startingLead)); err != nil {
		t.Fatal(err)
	}
	if !slices.Equal(live.started, []int64{raid.ID}) {
		t.Fatalf("started %v", live.started)
	}
	w.h.svc.liveRaidEnded(context.Background(), raid.ID)
	if !slices.Equal(live.ended, []int64{raid.ID}) {
		t.Fatalf("ended %v", live.ended)
	}
}
