package notify

import (
	"context"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/events"
)

// memRouterStore is an events.RouterStore over plain maps: the router's own reads are
// covered in package events.
type memRouterStore struct{ claims map[string]bool }

func (m *memRouterStore) Recipients(_ context.Context, a events.Audience) ([]int64, error) {
	return a.UserIDs, nil
}

func (m *memRouterStore) PrefsOf(context.Context, []int64) (map[int64]events.Prefs, error) {
	return nil, nil
}

func (m *memRouterStore) Claim(_ context.Context, uid int64, key string, ch events.Channel) (bool, error) {
	k := key + string(ch)
	if m.claims[k] {
		return false, nil
	}
	m.claims[k] = true
	return true, nil
}

func (m *memRouterStore) Unclaim(_ context.Context, _ int64, key string, ch events.Channel) error {
	delete(m.claims, key+string(ch))
	return nil
}

type recordingSink struct{ got []events.Event }

func (s *recordingSink) Deliver(_ context.Context, _ int64, e events.Event) error {
	s.got = append(s.got, e)
	return nil
}

func routedRig(t *testing.T) (*deliverRig, *EventProducer, *recordingSink) {
	t.Helper()
	r := newDeliverRig(t, map[EventName]bool{EventSimFinished: true, EventUploadParsed: true, EventUploadPrivate: true})
	push := &recordingSink{}
	router := &events.Router{Store: &memRouterStore{claims: map[string]bool{}}, Sinks: map[events.Channel]events.Sink{
		events.ChannelDiscord: &DiscordSink{Deliverer: r.d}, events.ChannelPush: push,
	}}
	return r, &EventProducer{Publisher: router, PublicBaseURL: r.d.PublicBaseURL}, push
}

func sampleSim(uid int64) SimFinished {
	at := time.Date(2026, 11, 4, 20, 0, 0, 0, time.UTC)
	return SimFinished{UserID: uid, SimID: "sim-1", Spec: "warrior-fury", Character: "Kraggor", DPS: 1234.5,
		Title: "BiS check", CreatedAt: at.Add(-time.Minute), FinishedAt: at}
}

func sampleUpload(uid int64, private bool) UploadParsed {
	return UploadParsed{UserID: uid, ReportID: "rep-1", Title: "MC night", Zone: "Molten Core", Visibility: "private",
		Private: private, Fights: 10, Kills: 9, FinishedAt: time.Date(2026, 11, 4, 23, 0, 0, 0, time.UTC)}
}

// The webhook post through the router is byte for byte the post the Deliverer makes directly.
func TestTheWebhookPostThroughTheRouterIsUnchanged(t *testing.T) {
	direct := newDeliverRig(t, map[EventName]bool{EventSimFinished: true, EventUploadParsed: true, EventUploadPrivate: true})
	direct.d.SimFinished(context.Background(), sampleSim(direct.uid))
	direct.d.UploadParsed(context.Background(), sampleUpload(direct.uid, true))

	routed, producer, push := routedRig(t)
	producer.SimFinished(context.Background(), sampleSim(routed.uid))
	producer.UploadParsed(context.Background(), sampleUpload(routed.uid, true))

	want, got := direct.poster.Posts(), routed.poster.Posts()
	if len(want) != 2 || !reflect.DeepEqual(want, got) {
		t.Fatalf("direct %+v\nrouted %+v", want, got)
	}
	if len(push.got) != 2 || push.got[0].Kind != events.SimFinished || push.got[1].Kind != events.ReportReady {
		t.Fatalf("push %+v", push.got)
	}
}

func TestThePrivateUploadGateStillHolds(t *testing.T) {
	r := newDeliverRig(t, map[EventName]bool{EventUploadParsed: true})
	router := &events.Router{Store: &memRouterStore{claims: map[string]bool{}}, Sinks: map[events.Channel]events.Sink{
		events.ChannelDiscord: &DiscordSink{Deliverer: r.d},
	}}
	producer := &EventProducer{Publisher: router, PublicBaseURL: r.d.PublicBaseURL}
	producer.UploadParsed(context.Background(), sampleUpload(r.uid, true))
	if len(r.poster.Posts()) != 0 {
		t.Fatal("a private report posted without upload_private")
	}
	producer.UploadParsed(context.Background(), UploadParsed{UserID: r.uid, ReportID: "rep-2", Visibility: "public"})
	if len(r.poster.Posts()) != 1 {
		t.Fatal("a public report did not post")
	}
}

func TestDiscordSinkIgnoresOtherKindsAndOtherAccounts(t *testing.T) {
	r := newDeliverRig(t, bothEvents)
	sink := &DiscordSink{Deliverer: r.d}
	for _, e := range []events.Event{
		{Kind: events.RaidPosted},
		{Kind: events.SimFinished, Source: sampleSim(r.uid + 1)},
	} {
		if err := sink.Deliver(context.Background(), r.uid, e); err != nil {
			t.Fatal(err)
		}
	}
	if len(r.poster.Posts()) != 0 {
		t.Fatal("posted")
	}
}

func TestSimAndUploadEventsReadPlainly(t *testing.T) {
	sim := SimFinishedEvent("https://foreversixty.gg", sampleSim(4))
	if sim.Title != "Kraggor's Fury Warrior sim finished" || sim.Body != "1,235 DPS" || sim.Link != "https://foreversixty.gg/sim/sim-1" {
		t.Fatalf("sim %+v", sim)
	}
	if sim.DedupeKey != "sim.finished:sim-1" || sim.Audience.UserIDs[0] != 4 {
		t.Fatalf("sim %+v", sim)
	}
	up := UploadParsedEvent("https://foreversixty.gg", sampleUpload(4, false))
	if up.Title != "MC night finished parsing" || up.Body != "Molten Core · 10 boss fights · 9 kills" || !strings.HasSuffix(up.Link, "/reports/rep-1") {
		t.Fatalf("upload %+v", up)
	}
	named := SimFinished{UserID: 1, SimID: "s", Character: `**Bold**`, Spec: "nope"}
	if e := SimFinishedEvent("", named); e.Title != "Your sim finished" {
		t.Fatalf("unknown spec %+v", e)
	}
}
