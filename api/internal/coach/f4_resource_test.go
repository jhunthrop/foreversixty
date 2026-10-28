package coach

import (
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/mechanics"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// f4WithDowntime is in with a curated 20 s downtime starting when an aura lands at 40 s.
func f4WithDowntime(in Input) Input {
	in.Fight = withAura(in.Fight, "Creature-0-1", 999, "Lift Off", nil, span(40_000, 41_000))
	in.Tables.Mechanics = &mechanics.Table{Downtime: []mechanics.Downtime{
		{Trigger: mechanics.PhaseStart{SpellID: 999, On: mechanics.OnAuraApplied}, DurationMS: 20_000},
	}}
	return in
}

func TestF4CapLeavesOutCuratedDowntime(t *testing.T) {
	in := f4WithDowntime(f4Capped(t, PowerRage, f4RageMax, 100_000, span(30_000, 65_000)))
	f := f4Resource(in).Finding
	if f == nil || f.Title != "Rage full for 15 seconds" {
		t.Fatalf("10 s before and 5 s after the downtime count: %+v", f)
	}
	if want := "Your rage sat at its cap for 15 s outside the downtime, longest at 0:30–0:40."; f.Detail != want {
		t.Errorf("detail %q, want %q", f.Detail, want)
	}
	if want := Fill(BasisCapOwnBar, map[string]string{"resource": "rage", "downtime": "downtime"}); f.Basis.Label != want {
		t.Errorf("basis %q, want %q", f.Basis.Label, want)
	}
}

func TestF4ManaCountsDowntime(t *testing.T) {
	in := f4WithDowntime(holyManaLow(t, 100_000, 45_000, 58_000))
	f := f4Resource(in).Finding
	if f == nil || f.Title != "Out of mana from 0:45" {
		t.Fatalf("§6.2 sets downtime aside for rage and energy only: %+v", f)
	}
}

func TestPowerNames(t *testing.T) {
	want := map[int64]string{PowerMana: "mana", PowerRage: "rage", PowerEnergy: "energy"}
	if len(PowerNames) != len(want) {
		t.Fatalf("PowerNames %v", PowerNames)
	}
	for id, name := range want {
		if PowerNames[id] != name {
			t.Errorf("power %d = %q, want %q", id, PowerNames[id], name)
		}
	}
	if PowerMana != 0 || PowerRage != 1 || PowerEnergy != 3 {
		t.Errorf("the client enum is mana 0, rage 1, energy 3")
	}
}

func TestResource(t *testing.T) {
	cases := []struct {
		name                  string
		in                    Input
		wantTitle, wantDetail string
		wantSev               Severity
		none                  bool
	}{
		{name: "fury rage capped 14 s of 232 s", in: furyRageCapped(t, 232_000, span(161_000, 170_000), span(60_000, 65_000)),
			wantTitle: "Rage full for 14 seconds", wantSev: SeverityMinor},
		{name: "holy out of mana from 3:05", in: holyManaLow(t, 232_000, 185_000, 200_000),
			wantTitle: "Out of mana from 3:05", wantSev: SeverityMinor, wantDetail: "15 s"},
		{name: "mana low only in the last 10%", in: holyManaLow(t, 232_000, 215_000, 232_000), none: true},
		{name: "out of mana before 75% is major", in: holyManaLow(t, 232_000, 150_000, 190_000),
			wantTitle: "Out of mana from 2:30", wantSev: SeverityMajor},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			f := resource(tc.in, CountedTime(tc.in)).Finding
			if tc.none {
				if f != nil {
					t.Fatalf("unexpected %+v", f)
				}
				return
			}
			if f == nil || f.Title != tc.wantTitle || f.Severity != tc.wantSev || len(f.Anchors) == 0 {
				t.Fatalf("got %+v", f)
			}
			if tc.wantDetail != "" && !strings.Contains(f.Detail, tc.wantDetail) {
				t.Errorf("detail %q", f.Detail)
			}
			text := strings.ToLower(f.Title + f.Detail + f.Basis.Label)
			if strings.Contains(text, "heal more") || strings.Contains(text, "leftover") {
				t.Errorf("healer copy is about running out early: %q", text)
			}
		})
	}
}

func f4Resource(in Input) outcome { return resource(in, CountedTime(in)) }

func TestF4CapThresholds(t *testing.T) {
	cases := []struct {
		name       string
		durationMS int64
		window     summary.Segment
		wantSev    Severity
		none       bool
	}{
		{"9 s is under the 10 s floor", 232_000, span(60_000, 69_000), "", true},
		{"10 s but under 5% of a long fight", 232_000, span(60_000, 70_000), "", true},
		{"10 s of 99 s counted is minor", 100_000, span(60_000, 70_000), SeverityMinor, false},
		{"15% capped is major", 100_000, span(10_000, 25_000), SeverityMajor, false},
		{"just under 15% stays minor", 100_000, span(10_000, 24_000), SeverityMinor, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			f := f4Resource(f4Capped(t, PowerRage, f4RageMax, tc.durationMS, tc.window)).Finding
			if tc.none {
				if f != nil {
					t.Fatalf("unexpected %+v", f)
				}
				return
			}
			if f == nil || f.Severity != tc.wantSev {
				t.Fatalf("want %s, got %+v", tc.wantSev, f)
			}
		})
	}
}

func TestF4EnergyCapTitle(t *testing.T) {
	o := f4Resource(f4Capped(t, PowerEnergy, f4EnergyMax, 100_000, span(40_000, 52_000)))
	if o.Finding == nil || o.Finding.Title != "Energy full for 12 seconds" {
		t.Fatalf("got %+v", o.Finding)
	}
	if !strings.Contains(o.Finding.Detail, "energy") {
		t.Errorf("detail %q", o.Finding.Detail)
	}
}

func TestF4CapDetailAnchorAndWindows(t *testing.T) {
	o := f4Resource(furyRageCapped(t, 232_000, span(161_000, 170_000), span(60_000, 65_000)))
	f := o.Finding
	if f == nil {
		t.Fatal("no finding")
	}
	if want := "Your rage sat at its cap for 14 s, longest at 2:41–2:50."; f.Detail != want {
		t.Errorf("detail %q, want %q", f.Detail, want)
	}
	if want := Fill(BasisOwnBar, map[string]string{"resource": "rage"}); f.Basis.Label != want || f.Basis.Kind != BasisOwnLog {
		t.Errorf("basis %+v", f.Basis)
	}
	if len(o.Anchors) != 1 || o.Anchors[0].Kind != AnchorResource || *o.Anchors[0].AtMS != 161_000 || *o.Anchors[0].ToMS != 170_000 {
		t.Errorf("anchors %+v", o.Anchors)
	}
	if len(f.Anchors) != 1 || f.Anchors[0] != "#0" {
		t.Errorf("finding anchors %v", f.Anchors)
	}
	if len(o.Windows) != 2 || o.Windows[0].FromMS != 60_000 || o.Windows[1].ToMS != 170_000 {
		t.Errorf("windows %+v", o.Windows)
	}
	assertNumbers(t, f.Facts, "14")
	assertTimes(t, f.Facts, "2:41", "2:50")
}

func TestF4CapIgnoresTimeDead(t *testing.T) {
	in := furyRageCapped(t, 100_000, span(60_000, 100_000))
	in.Fight = withDeath(in.Fight, furyGUID, 65_000, 0, "")
	if f := f4Resource(in).Finding; f != nil {
		t.Fatalf("5 s alive at cap is under the floor, got %+v", f)
	}
	in = furyRageCapped(t, 100_000, span(60_000, 100_000))
	in.Fight = withDeath(in.Fight, furyGUID, 75_000, 0, "")
	f := f4Resource(in).Finding
	if f == nil || f.Title != "Rage full for 15 seconds" {
		t.Fatalf("counted only to the death at 75 s: %+v", f)
	}
}

func TestF4CapIgnoresTimeBeforeFirstCast(t *testing.T) {
	in := furyRageCapped(t, 100_000, span(0, 1_000), span(20_000, 29_000))
	if f := f4Resource(in).Finding; f != nil {
		t.Fatalf("the second before the first cast is set aside, leaving 9 s: %+v", f)
	}
}

func TestF4RolesDecideWhichBars(t *testing.T) {
	capped := func(in Input, powerType, max int64) Input {
		in = f4Prepared(in, 100_000)
		return f4WithTrack(in, powerType, max, f4Series(100_000, max/2, max, span(40_000, 60_000)))
	}
	cases := []struct {
		name string
		in   Input
		want bool
	}{
		{"dps rage", capped(furyInput(t), PowerRage, f4RageMax), true},
		{"dps energy", capped(furyInput(t), PowerEnergy, f4EnergyMax), true},
		{"tank rage", capped(protInput(t), PowerRage, f4RageMax), true},
		{"tank energy is not theirs", capped(protInput(t), PowerEnergy, f4EnergyMax), false},
		{"healer rage is not theirs", capped(holyInput(t), PowerRage, f4RageMax), false},
		{"healer energy is not theirs", capped(holyInput(t), PowerEnergy, f4EnergyMax), false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := f4Resource(tc.in).Finding != nil; got != tc.want {
				t.Fatalf("fires = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestF4ManaForTankAndDPS(t *testing.T) {
	for name, base := range map[string]Input{"tank": protInput(t), "dps": furyInput(t)} {
		t.Run(name, func(t *testing.T) {
			in := f4Prepared(base, 232_000)
			in = f4WithTrack(in, PowerMana, f4ManaMax, f4Series(232_000, f4ManaMax/2, 0, span(100_000, 120_000)))
			f := f4Resource(in).Finding
			if f == nil || f.Title != "Out of mana from 1:40" || f.Severity != SeverityMajor {
				t.Fatalf("got %+v", f)
			}
		})
	}
}

func TestF4ManaThresholds(t *testing.T) {
	cases := []struct {
		name     string
		from, to int64
		none     bool
		wantSev  Severity
	}{
		{"9 s low is under the floor", 40_000, 49_000, true, ""},
		{"10 s low early fires major", 40_000, 50_000, false, SeverityMajor},
		{"starts at exactly 90% does not fire", 90_000, 100_000, true, ""},
		{"starts just before 90% is minor", 89_000, 100_000, false, SeverityMinor},
		{"starts at exactly 75% is minor", 75_000, 90_000, false, SeverityMinor},
		{"starts just before 75% is major", 74_000, 90_000, false, SeverityMajor},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			f := f4Resource(holyManaLow(t, 100_000, tc.from, tc.to)).Finding
			if tc.none {
				if f != nil {
					t.Fatalf("unexpected %+v", f)
				}
				return
			}
			if f == nil || f.Severity != tc.wantSev {
				t.Fatalf("want %s, got %+v", tc.wantSev, f)
			}
		})
	}
}

func TestF4ManaFivePercentBoundary(t *testing.T) {
	in := f4Prepared(holyInput(t), 100_000)
	at := func(v int64) Input {
		return f4WithTrack(in, PowerMana, f4ManaMax, f4Series(100_000, f4ManaMax/2, v, span(20_000, 40_000)))
	}
	if f := f4Resource(at(f4ManaMax / 20)).Finding; f != nil {
		t.Errorf("exactly 5%% is not under 5%%: %+v", f)
	}
	if f := f4Resource(at(f4ManaMax/20 - 1)).Finding; f == nil {
		t.Error("just under 5% fires")
	}
}

func TestF4ManaDetailVariants(t *testing.T) {
	died := holyManaLow(t, 232_000, 185_000, 232_000)
	died.Fight = withDeath(died.Fight, priestGUID, 200_000, 0, "")
	cases := []struct {
		name string
		in   Input
		want string
	}{
		{"died while out of mana", died, "Your mana stayed under 5% from 3:05 until you died at 3:20, 15 s in all."},
		{"out to the end", holyManaLow(t, 232_000, 150_000, 232_000), "Your mana stayed under 5% from 2:30, 82 s in all."},
		{"recovered", holyManaLow(t, 232_000, 185_000, 200_000), "Your mana stayed under 5% from 3:05, 15 s in all."},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			o := f4Resource(tc.in)
			if o.Finding == nil || o.Finding.Detail != tc.want {
				t.Fatalf("got %+v, want %q", o.Finding, tc.want)
			}
			assertTimes(t, o.Finding.Facts, Clock(o.Windows[0].FromMS))
			assertNumbers(t, o.Finding.Facts, "5")
		})
	}
}

func TestF4ManaAnchorAndBasis(t *testing.T) {
	o := f4Resource(holyManaLow(t, 232_000, 185_000, 200_000))
	if o.Finding == nil || len(o.Anchors) != 1 || o.Anchors[0].Kind != AnchorResource || *o.Anchors[0].AtMS != 185_000 {
		t.Fatalf("got %+v %+v", o.Finding, o.Anchors)
	}
	if want := Fill(BasisOwnBar, map[string]string{"resource": "mana"}); o.Finding.Basis.Label != want {
		t.Errorf("basis %q", o.Finding.Basis.Label)
	}
	if len(o.Windows) != 1 || o.Windows[0].FromMS != 185_000 || o.Windows[0].ToMS != 200_000 {
		t.Errorf("windows %+v", o.Windows)
	}
}

func TestF4WorstBarWins(t *testing.T) {
	in := f4Prepared(furyInput(t), 100_000)
	in = f4WithTrack(in, PowerRage, f4RageMax, f4Series(100_000, 0, f4RageMax, span(40_000, 52_000)))
	in = f4WithTrack(in, PowerMana, f4ManaMax, f4Series(100_000, f4ManaMax/2, 0, span(50_000, 80_000)))
	f := f4Resource(in).Finding
	if f == nil || f.Title != "Out of mana from 0:50" || f.Severity != SeverityMajor {
		t.Fatalf("a major beats a minor: %+v", f)
	}
}

func TestF4NothingToCheck(t *testing.T) {
	t.Run("no track", func(t *testing.T) {
		if o := f4Resource(f4Prepared(furyInput(t), 100_000)); o.Finding != nil || len(o.Anchors) != 0 {
			t.Fatalf("got %+v", o)
		}
	})
	t.Run("no maximum in the log", func(t *testing.T) {
		in := f4WithTrack(f4Prepared(furyInput(t), 100_000), PowerRage, 0, f4Series(100_000, 0, 0))
		if f := f4Resource(in).Finding; f != nil {
			t.Fatalf("a bar with no cap cannot be at it: %+v", f)
		}
	})
	t.Run("no counted time", func(t *testing.T) {
		in := furyRageCapped(t, 100_000, span(0, 100_000))
		if f := resource(in, Counted{}).Finding; f != nil {
			t.Fatalf("got %+v", f)
		}
	})
}

func TestF4TextIsFilled(t *testing.T) {
	text := textOf(f4Resource(furyRageCapped(t, 232_000, span(161_000, 170_000), span(60_000, 65_000))))
	for _, want := range []string{"Rage full for 14 seconds", "2:41", "2:50"} {
		if !strings.Contains(text, want) {
			t.Errorf("text lacks %q: %q", want, text)
		}
	}
	if strings.Contains(text, "{") {
		t.Errorf("unfilled placeholder: %q", text)
	}
}
