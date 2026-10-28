package coach

import (
	"strings"
	"testing"
)

func TestIdle(t *testing.T) {
	cases := []struct {
		name      string
		in        Input
		wantSev   Severity
		wantTimes []string
		none      bool
	}{
		{name: "dps 24 s idle of 200 s is major", in: furyIdle(t, 200_000, span(40_000, 52_000), span(90_000, 102_000)),
			wantSev: SeverityMajor, wantTimes: []string{"0:40–0:52", "1:30–1:42"}},
		{name: "dps 9 s idle is under the floor", in: furyIdle(t, 200_000, span(40_000, 49_000)), none: true},
		{name: "downtime windows never count", in: furyIdleInsideDowntime(t), none: true},
		{name: "healer: F3 is off", in: holyIdle(t, 200_000, span(40_000, 80_000)), none: true},
		{name: "tank counts", in: protIdle(t, 200_000, span(40_000, 52_000)), wantSev: SeverityMinor, wantTimes: []string{"0:40–0:52"}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			f := idle(tc.in, CountedTime(tc.in)).Finding
			if tc.none {
				if f != nil {
					t.Fatalf("unexpected %+v", f)
				}
				return
			}
			if f == nil || f.Severity != tc.wantSev || len(f.Anchors) != len(tc.wantTimes) {
				t.Fatalf("got %+v", f)
			}
			assertTimes(t, f.Facts, tc.wantTimes...)
		})
	}
}

func TestIdleShowsTheThreeLongestWindowsAndAnchorsEach(t *testing.T) {
	in := furyIdle(t, 400_000, span(10_000, 14_000), span(50_000, 62_000), span(100_000, 105_000), span(150_000, 160_000))
	o := idle(in, CountedTime(in))
	if o.Finding == nil {
		t.Fatal("no finding")
	}
	if len(o.Finding.Anchors) != IdleWindowsShown || len(o.Anchors) != IdleWindowsShown {
		t.Fatalf("anchors %v / %v", o.Finding.Anchors, o.Anchors)
	}
	assertTimes(t, o.Finding.Facts, "0:50–1:02", "2:30–2:40", "1:40–1:45")
	for _, a := range o.Anchors {
		if a.Kind != AnchorTime || a.AtMS == nil || a.ToMS == nil {
			t.Fatalf("anchor %+v", a)
		}
	}
	if len(o.Windows) != 4 {
		t.Fatalf("windows %v: want every idle window for links", o.Windows)
	}
	if strings.Contains(textOf(o), "0:10") {
		t.Fatalf("the shortest window is shown: %s", textOf(o))
	}
}

func TestIdleSeverityBoundaries(t *testing.T) {
	cases := []struct {
		name    string
		in      Input
		wantSev Severity
		none    bool
	}{
		{"10 s of 200 s is 5%: fires minor", furyIdle(t, 200_000, span(40_000, 50_000)), SeverityMinor, false},
		{"20 s of 200 s is 10%: major", furyIdle(t, 200_000, span(40_000, 60_000)), SeverityMajor, false},
		{"10 s of 300 s is 3.3%: under the percent", furyIdle(t, 300_000, span(40_000, 50_000)), "", true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			f := idle(tc.in, CountedTime(tc.in)).Finding
			if tc.none {
				if f != nil {
					t.Fatalf("unexpected %+v", f)
				}
				return
			}
			if f == nil || f.Severity != tc.wantSev {
				t.Fatalf("got %+v", f)
			}
		})
	}
}

func TestIdleAssignmentWindowsNeverCount(t *testing.T) {
	in := withAssignment(furyIdle(t, 200_000, span(40_000, 52_000)), 38_000, 54_000)
	if f := idle(in, CountedTime(in)).Finding; f != nil {
		t.Fatalf("an officer-assigned stretch counted: %+v", f)
	}
}

func TestIdleBasisIsThePoolMedianElseTheThreshold(t *testing.T) {
	own := furyIdle(t, 200_000, span(40_000, 64_000))
	f := idle(own, CountedTime(own)).Finding
	if f == nil || f.Basis.Kind != BasisOwnLog || !strings.HasPrefix(f.Basis.Label, FromYourLog) {
		t.Fatalf("own-log basis: %+v", f)
	}
	pooled := withPool(own, idlePool(97, 25))
	g := idle(pooled, CountedTime(pooled)).Finding
	if g == nil || g.Basis.Kind != BasisTopLogs || g.Basis.N != 25 {
		t.Fatalf("pool basis: %+v", g)
	}
	assertNumbers(t, g.Facts, "97", "88")
}

func TestIdleNumbersAndText(t *testing.T) {
	in := furyIdle(t, 200_000, span(40_000, 52_000), span(90_000, 102_000))
	o := idle(in, CountedTime(in))
	assertNumbers(t, o.Finding.Facts, "24", "12")
	for _, want := range []string{"Idle for 24 seconds", "Idle outside downtime for 24 s, 12% of the time you were up. Longest: 0:40–0:52 and 1:30–1:42."} {
		if !strings.Contains(textOf(o), want) {
			t.Errorf("text lacks %q: %s", want, textOf(o))
		}
	}
	if o.Finding.Code != CodeIdle {
		t.Errorf("code %q", o.Finding.Code)
	}
}

func TestIdleLeadInBeforeTheFirstCastNeverCounts(t *testing.T) {
	lead := withFirstCast(furyIdle(t, 200_000, span(0, 15_000)), 15_000)
	if f := idle(lead, CountedTime(lead)).Finding; f != nil {
		t.Fatalf("the lead-in counted: %+v", f)
	}
	mixed := withFirstCast(furyIdle(t, 200_000, span(0, 15_000), span(60_000, 75_000)), 15_000)
	o := idle(mixed, CountedTime(mixed))
	if o.Finding == nil {
		t.Fatal("no finding for the 15 s gap after the first cast")
	}
	assertNumbers(t, o.Finding.Facts, "15")
	for _, w := range o.Windows {
		if w.FromMS < 15_000 {
			t.Fatalf("window %+v starts before the first cast", w)
		}
	}
	assertTimes(t, o.Finding.Facts, "1:00–1:15")
}

func TestIdleStatedTotalIsTheSumOfTheWindows(t *testing.T) {
	in := furyIdle(t, 400_000, span(10_000, 14_000), span(50_000, 62_000), span(100_000, 105_000), span(150_000, 160_000))
	o := idle(in, CountedTime(in))
	if o.Finding == nil || IdleTotalMS(o.Windows) != 31_000 {
		t.Fatalf("total %d finding %+v", IdleTotalMS(o.Windows), o.Finding)
	}
	assertNumbers(t, o.Finding.Facts, "31", "12")
	if o.Finding.Title != "Idle for 31 seconds" || !strings.Contains(o.Finding.Detail, "for 31 s,") ||
		!strings.HasSuffix(o.Finding.Detail, "Longest: 0:50–1:02, 1:40–1:45 and 2:30–2:40.") {
		t.Fatalf("detail %q", o.Finding.Detail)
	}
}

// The pool's active share and the player's are one measure over one window (A5): the
// profile stores 100 − counted idle / counted time, the number F3 prints as "You".
func TestIdleBasisComparesTheSameMeasureOverCountedTime(t *testing.T) {
	in := furyIdle(t, 200_000, span(40_000, 64_000))
	in.Row.ActivityPct = 55 // the roster's whole-fight share: never the comparison
	p := BuildProfile(in)
	if p.ActivePct == nil {
		t.Fatal("a DPS profile records its counted active share")
	}
	pooled := withPool(in, idlePool(*p.ActivePct, 25))
	f := idle(pooled, CountedTime(pooled)).Finding
	if f == nil || !strings.HasSuffix(f.Basis.Label, "% active. You: "+PctNumber(*p.ActivePct)+"%.") {
		t.Fatalf("the player's own profile reads as the pool would: %+v", f)
	}
	if holy := BuildProfile(holyIdle(t, 200_000, span(40_000, 64_000))); holy.ActivePct != nil {
		t.Errorf("a healer is never idle-judged, so records no active share: %v", *holy.ActivePct)
	}
}
