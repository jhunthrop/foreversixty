package engagement

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
)

type mark struct {
	userID  int64
	day     string
	surface Surface
}

type fakeMarker struct {
	marks []mark
	err   error
}

func (f *fakeMarker) MarkActive(_ context.Context, userID int64, day time.Time, surface Surface) error {
	f.marks = append(f.marks, mark{userID, day.Format(time.DateOnly), surface})
	return f.err
}

var (
	browser   = auth.Actor{UserID: 5, Role: "user", Method: "session"}
	companion = auth.Actor{UserID: 5, Role: "user", Method: "device", DeviceID: "d1", DeviceKind: auth.DeviceKindCompanion}
	phone     = auth.Actor{UserID: 5, Role: "user", Method: "device", DeviceID: "d2", DeviceKind: auth.DeviceKindApp}
)

// newRecorder runs the writes in place on a clock the test moves.
func newRecorder(f *fakeMarker, now *time.Time) *Recorder {
	return &Recorder{Store: f, Now: func() time.Time { return *now }, run: func(g func()) { g() }}
}

func TestSurfaceOfNamesEachWayIn(t *testing.T) {
	if SurfaceOf(browser) != SurfaceSite || SurfaceOf(companion) != SurfaceAddon || SurfaceOf(phone) != SurfaceApp {
		t.Fatalf("surfaces: %d %d %d", SurfaceOf(browser), SurfaceOf(companion), SurfaceOf(phone))
	}
}

func TestRecorderWritesOncePerAccountDayAndSurface(t *testing.T) {
	f := &fakeMarker{}
	now := time.Date(2026, 11, 13, 9, 0, 0, 0, time.UTC)
	r := newRecorder(f, &now)
	for range 3 {
		r.Active(browser)
	}
	r.Active(companion)
	r.Active(companion)
	r.Active(phone)
	r.Active(auth.Actor{})
	now = now.Add(14 * time.Hour) // 23:00, the same UTC day
	r.Active(browser)
	if len(f.marks) != 3 {
		t.Fatalf("marks = %+v, want one per surface", f.marks)
	}
	want := []mark{{5, "2026-11-13", SurfaceSite}, {5, "2026-11-13", SurfaceAddon}, {5, "2026-11-13", SurfaceApp}}
	for i, m := range want {
		if f.marks[i] != m {
			t.Fatalf("mark %d = %+v, want %+v", i, f.marks[i], m)
		}
	}
	now = now.Add(2 * time.Hour) // past UTC midnight
	r.Active(browser)
	if len(f.marks) != 4 || f.marks[3] != (mark{5, "2026-11-14", SurfaceSite}) {
		t.Fatalf("the next day's first request: marks = %+v", f.marks)
	}
}

func TestRecorderTriesAgainAfterAFailedWrite(t *testing.T) {
	f := &fakeMarker{err: errors.New("db down")}
	now := time.Date(2026, 11, 13, 9, 0, 0, 0, time.UTC)
	r := newRecorder(f, &now)
	r.Active(browser)
	f.err = nil
	r.Active(browser)
	r.Active(browser)
	if len(f.marks) != 2 {
		t.Fatalf("marks = %+v, want the failed write and one retry", f.marks)
	}
}

func TestRecorderReadsTheDayInUTC(t *testing.T) {
	f := &fakeMarker{}
	now := time.Date(2026, 11, 13, 20, 0, 0, 0, time.FixedZone("CST", -6*3600)) // 02:00 UTC on the 14th
	newRecorder(f, &now).Active(browser)
	if f.marks[0].day != "2026-11-14" {
		t.Fatalf("day = %s, want the UTC day", f.marks[0].day)
	}
}
