package addon

import (
	"context"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/events"
)

func exportAt(level string) string {
	return "FS1:1.15.8:warrior:orc:0/0/0:1=123|level=" + level
}

func TestSyncEventsLinkAndBand(t *testing.T) {
	exports := []Export{
		{Name: "Kraggor", Region: "us", Ruleset: "normal", Export: exportAt("31")},
		{Name: "Newbie", Region: "us", Ruleset: "normal", Export: exportAt("12")},
		{Name: "Steady", Region: "us", Ruleset: "normal", Export: exportAt("44")},
	}
	if lvl := exportLevel(exports[0].Export); lvl != 31 {
		t.Fatalf("level %d", lvl)
	}
	prior := map[string]int{"us/normal/kraggor": 29, "us/normal/steady": 41}
	got := SyncEvents(7, prior, exports, "https://foreversixty.gg")
	if len(got) != 2 {
		t.Fatalf("events %+v", got)
	}
	band, linked := got[0], got[1]
	if band.Kind != events.BandEntered || band.Title != "Kraggor reached level 30" || band.DedupeKey != "band.entered:us/normal/kraggor:30" {
		t.Fatalf("band %+v", band)
	}
	if linked.Kind != events.CharacterLinked || linked.Link != "https://foreversixty.gg/character/us/normal/newbie" || linked.Audience.UserIDs[0] != 7 {
		t.Fatalf("linked %+v", linked)
	}
	for _, e := range got {
		if err := e.Validate(); err != nil {
			t.Fatal(err)
		}
	}
	// An export with no level, before or after, never tells a band.
	if got := SyncEvents(7, map[string]int{"us/normal/kraggor": 0}, exports[:1], ""); len(got) != 0 {
		t.Fatalf("unknown prior level told %+v", got)
	}
}

func TestBandOf(t *testing.T) {
	for level, want := range map[int]int{1: 0, 19: 0, 20: 20, 39: 30, 60: 60} {
		if got := bandOf(level); got != want {
			t.Errorf("bandOf(%d) = %d", level, got)
		}
	}
}

type syncRecorder struct{ got []events.Event }

func (r *syncRecorder) Publish(_ context.Context, e events.Event) error {
	r.got = append(r.got, e)
	return nil
}

// Through the store: the first sync links, the next across a band tells the band, and a
// repeat tells nothing new.
func TestASyncPublishesLinkedThenBand(t *testing.T) {
	h := newHarness(t)
	rec := &syncRecorder{}
	svc := &Service{Store: h.store, Events: rec, PublicBaseURL: "https://foreversixty.gg"}
	ctx := context.Background()
	sync := func(level string) {
		exports := []Export{{Name: "Kraggor", Region: "us", Ruleset: "normal", Export: exportAt(level)}}
		n := svc.beforeSync(ctx, h.owner, exports)
		if err := h.store.PutExports(ctx, h.owner, exports); err != nil {
			t.Fatal(err)
		}
		svc.afterSync(ctx, n, exports)
	}
	sync("19")
	sync("21")
	sync("22")
	if len(rec.got) != 2 || rec.got[0].Kind != events.CharacterLinked || rec.got[1].Kind != events.BandEntered {
		t.Fatalf("events %+v", rec.got)
	}
}
