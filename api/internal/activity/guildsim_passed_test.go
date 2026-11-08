package activity

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/events"
)

// recordingPublisher keeps every published event.
type recordingPublisher struct {
	mu  sync.Mutex
	got []events.Event
}

func (p *recordingPublisher) Publish(_ context.Context, e events.Event) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.got = append(p.got, e)
	return nil
}

// A sim that moves a member above others tells each one passed, and only them. guildOf: Ma
// 600, Mb 590, Mc 580, Md 570, Me 560; Me now sims 585, passing Mc and Md, not Ma or Mb.
func TestGuildSimPassedTellsEachMemberPassed(t *testing.T) {
	s := testStore(t)
	_, users := guildOf(t, s, 5)
	sim(t, s, "sMe2", "us/normal/me", users[4], "import", 585, 60, "raid", testNow)
	pub := &recordingPublisher{}
	w := &GuildSimWriter{Store: s, Events: pub, PublicBaseURL: "https://foreversixty.gg", Now: func() time.Time { return testNow }}
	w.SimFinished(context.Background(), "sMe2")
	if len(pub.got) != 2 {
		t.Fatalf("events %+v", pub.got)
	}
	byUser := map[int64]events.Event{}
	for _, e := range pub.got {
		if e.Kind != events.GuildSimPassed || e.Audience.Kind != events.ToUsers || len(e.Audience.UserIDs) != 1 {
			t.Fatalf("event %+v", e)
		}
		byUser[e.Audience.UserIDs[0]] = e
	}
	mc, md := byUser[users[2]], byUser[users[3]]
	if mc.Title != "Me passed you in the guild ladder" || mc.Body != "585 DPS, +5 over your 580 (Fury Warrior, 60 raid)." {
		t.Fatalf("Mc's event %+v", mc)
	}
	if md.Body != "585 DPS, +15 over your 570 (Fury Warrior, 60 raid)." {
		t.Fatalf("Md's event %+v", md)
	}
	if mc.Data["guild_id"] == "" || mc.Link != "https://foreversixty.gg/app/guild/"+mc.Data["guild_id"] {
		t.Fatalf("Mc's link or icon %+v", mc)
	}
}

// A first sim in a group passes nobody (there was no place to move from), and a sim that
// falls passes nobody.
func TestGuildSimPassedNeedsAMoveUp(t *testing.T) {
	s := testStore(t)
	_, users := guildOf(t, s, 5)
	pub := &recordingPublisher{}
	w := &GuildSimWriter{Store: s, Events: pub, Now: func() time.Time { return testNow }}
	w.SimFinished(context.Background(), "sMe")
	sim(t, s, "sMa2", "us/normal/ma", users[0], "import", 550, 60, "raid", testNow)
	w.SimFinished(context.Background(), "sMa2")
	if len(pub.got) != 0 {
		t.Fatalf("events %+v", pub.got)
	}
}
