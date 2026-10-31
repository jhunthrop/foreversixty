package seats

import (
	"context"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/events"
)

type eventRecorder struct{ got []events.Event }

func (r *eventRecorder) Publish(_ context.Context, e events.Event) error {
	if err := e.Validate(); err != nil {
		return err
	}
	r.got = append(r.got, e)
	return nil
}

// giveOne gives slot 1 of a fresh guild to a member, paying for their own Premium or not,
// and returns what was published.
func giveOne(t *testing.T, paying bool) (int64, []events.Event) {
	t.Helper()
	h := newHarness(t)
	rec := &eventRecorder{}
	h.svc.Events = rec
	gid, leader := h.guild(t, "guild")
	amy := h.member(t, gid, "Amy", "member")
	if paying {
		h.payingPremium(t, amy, testNow.Add(5*24*time.Hour))
	}
	if _, err := h.svc.Give(bg, gid, 1, leader, amy); err != nil {
		t.Fatal(err)
	}
	// A refused give (the slot is taken) tells nothing more.
	if _, err := h.svc.Give(bg, gid, 1, leader, leader); err == nil {
		t.Fatal("a taken slot was given")
	}
	return amy, rec.got
}

func TestGivingASeatTellsTheMember(t *testing.T) {
	amy, got := giveOne(t, false)
	if len(got) != 1 || got[0].Kind != events.PremiumSeat || got[0].Audience.UserIDs[0] != amy || got[0].Body != seatCopy.Held {
		t.Fatalf("given %+v", got)
	}
	amy, got = giveOne(t, true)
	if len(got) != 1 || got[0].Audience.UserIDs[0] != amy || got[0].Body != seatCopy.Offered {
		t.Fatalf("offered %+v", got)
	}
}
