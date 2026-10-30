package events

import (
	"testing"
	"time"
)

func TestEveryKindHasOneSpec(t *testing.T) {
	seen := map[Kind]bool{}
	for _, s := range Catalogue {
		if seen[s.Kind] {
			t.Fatalf("%s twice", s.Kind)
		}
		seen[s.Kind] = true
	}
	if len(seen) != 20 {
		t.Fatalf("catalogue has %d kinds, want the plan's 20", len(seen))
	}
	if s, _ := Lookup(RaidStarting); !s.QuietExempt {
		t.Fatal("raid.starting must be exempt from quiet hours")
	}
	if s, _ := Lookup(BisChanged); s.DefaultOn {
		t.Fatal("bis.changed is off by default")
	}
}

func TestPrefsDefaultsAndOverrides(t *testing.T) {
	p := Prefs{}
	if !p.Allows(RaidPosted, ChannelPush) || p.Allows(BisChanged, ChannelPush) {
		t.Fatal("defaults")
	}
	p = p.With(RaidPosted, ChannelPush, false).With(BisChanged, ChannelDiscord, true)
	if p.Allows(RaidPosted, ChannelPush) || !p.Allows(RaidPosted, ChannelDiscord) {
		t.Fatal("override off")
	}
	if !p.Allows(BisChanged, ChannelDiscord) || p.Allows(BisChanged, ChannelPush) {
		t.Fatal("override on")
	}
	if p.Allows("nope", ChannelPush) {
		t.Fatal("unknown kind allowed")
	}
}

func TestPrefsWithDoesNotMutate(t *testing.T) {
	a := Prefs{}.With(RaidPosted, ChannelPush, false)
	b := a.With(RaidPosted, ChannelPush, true)
	if a.Allows(RaidPosted, ChannelPush) || !b.Allows(RaidPosted, ChannelPush) {
		t.Fatal("With mutated its receiver")
	}
}

func TestQuietHoursInTheplayersZone(t *testing.T) {
	// 22:00 to 07:00 in New York.
	p := Prefs{TimeZone: "America/New_York", Quiet: QuietHours{Enabled: true, Start: 22 * 60, End: 7 * 60}}
	ny, _ := time.LoadLocation("America/New_York")
	cases := []struct {
		at   time.Time
		want bool
	}{
		{time.Date(2026, 11, 4, 23, 0, 0, 0, ny), true},
		{time.Date(2026, 11, 4, 3, 0, 0, 0, ny), true},
		{time.Date(2026, 11, 4, 7, 0, 0, 0, ny), false},
		{time.Date(2026, 11, 4, 21, 59, 0, 0, ny), false},
		{time.Date(2026, 11, 4, 12, 0, 0, 0, time.UTC), false}, // 07:00 in NY
	}
	for _, c := range cases {
		if got := p.InQuiet(c.at); got != c.want {
			t.Errorf("%v: quiet = %v, want %v", c.at, got, c.want)
		}
	}
	// A same-day window.
	day := Prefs{Quiet: QuietHours{Enabled: true, Start: 9 * 60, End: 17 * 60}}
	if !day.InQuiet(time.Date(2026, 11, 4, 12, 0, 0, 0, time.UTC)) || day.InQuiet(time.Date(2026, 11, 4, 18, 0, 0, 0, time.UTC)) {
		t.Fatal("same-day window in UTC")
	}
	off := Prefs{Quiet: QuietHours{Start: 0, End: 24 * 60}}
	if off.InQuiet(time.Now()) {
		t.Fatal("disabled quiet hours")
	}
}

func TestQuietHoursValidate(t *testing.T) {
	for _, q := range []QuietHours{{Start: -1}, {Start: 24 * 60}, {End: 1440}} {
		if q.Validate() == nil {
			t.Fatalf("%+v accepted", q)
		}
	}
	if (QuietHours{Enabled: true, Start: 60, End: 60}).Validate() == nil {
		t.Fatal("an empty window accepted")
	}
	if ValidTimeZone("Not/AZone") || !ValidTimeZone("Europe/Berlin") || ValidTimeZone("") {
		t.Fatal("time zones")
	}
}
