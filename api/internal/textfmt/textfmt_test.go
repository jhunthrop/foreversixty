package textfmt

import (
	"encoding/json"
	"os"
	"testing"
	"time"
)

// fixturePath is the house formats every surface replays (apps/shared/formats.json).
const fixturePath = "../../../apps/shared/formats.json"

type instantCase struct {
	In            string `json:"in"`
	Zone          string `json:"zone"`
	Now           string `json:"now"`
	Locale        string `json:"locale"`
	AlwaysMinutes bool   `json:"always_minutes"`
	Seconds       bool   `json:"seconds"`
	Out           string `json:"out"`
}

type numberCase struct {
	In  float64 `json:"in"`
	Out string  `json:"out"`
}

type durationCase struct {
	MS  int64  `json:"ms"`
	Out string `json:"out"`
}

type fixture struct {
	SiteLocale string `json:"site_locale"`
	Dates      struct {
		DayMonthYear    []instantCase `json:"day_month_year"`
		WeekdayDayMonth []instantCase `json:"weekday_day_month"`
		LongDay         []instantCase `json:"long_day"`
	} `json:"dates"`
	Clock     []instantCase `json:"clock"`
	Durations struct {
		Clock  []durationCase `json:"clock"`
		Spoken []durationCase `json:"spoken"`
	} `json:"durations"`
	Numbers struct {
		Grouped []numberCase `json:"grouped"`
	} `json:"numbers"`
}

func loadFixture(t *testing.T) fixture {
	t.Helper()
	raw, err := os.ReadFile(fixturePath)
	if err != nil {
		t.Fatalf("read the shared formats: %v", err)
	}
	var f fixture
	if err := json.Unmarshal(raw, &f); err != nil {
		t.Fatalf("parse %s: %v", fixturePath, err)
	}
	if f.SiteLocale == "" || len(f.Clock) == 0 || len(f.Durations.Spoken) == 0 || len(f.Numbers.Grouped) == 0 {
		t.Fatalf("%s lost a section this package replays", fixturePath)
	}
	return f
}

func instant(t *testing.T, s string) time.Time {
	t.Helper()
	at, err := time.Parse(time.RFC3339, s)
	if err != nil {
		t.Fatalf("instant %q: %v", s, err)
	}
	return at
}

func zone(t *testing.T, name string) *time.Location {
	t.Helper()
	loc, err := time.LoadLocation(name)
	if err != nil {
		t.Fatalf("zone %q: %v", name, err)
	}
	return loc
}

func check(t *testing.T, what string, c any, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("%s(%+v) = %q, want %q", what, c, got, want)
	}
}

// TestTheHouseFormats replays the shared examples. The API never knows its reader's locale, so
// it replays the clock cases of the site's own locale (a Discord message that needs the
// reader's clock writes <t:unix:t> instead).
func TestTheHouseFormats(t *testing.T) {
	f := loadFixture(t)
	for _, c := range f.Dates.DayMonthYear {
		check(t, "DayMonthYear", c, DayMonthYear(instant(t, c.In), zone(t, c.Zone)), c.Out)
	}
	for _, c := range f.Dates.WeekdayDayMonth {
		check(t, "WeekdayDayMonth", c, WeekdayDayMonth(instant(t, c.In), zone(t, c.Zone), instant(t, c.Now)), c.Out)
	}
	for _, c := range f.Dates.LongDay {
		check(t, "LongDay", c, LongDay(instant(t, c.In), zone(t, c.Zone)), c.Out)
	}
	clocks := 0
	for _, c := range f.Clock {
		if c.Locale != f.SiteLocale || c.AlwaysMinutes {
			continue
		}
		clocks++
		if c.Seconds {
			check(t, "ClockSeconds", c, ClockSeconds(instant(t, c.In), zone(t, c.Zone)), c.Out)
			continue
		}
		check(t, "Clock", c, Clock(instant(t, c.In), zone(t, c.Zone)), c.Out)
	}
	if clocks == 0 {
		t.Errorf("no clock case in the site's locale %s", f.SiteLocale)
	}
	for _, c := range f.Durations.Spoken {
		check(t, "Duration", c, Duration(time.Duration(c.MS)*time.Millisecond), c.Out)
	}
	for _, c := range f.Durations.Clock {
		check(t, "FightClock", c, FightClock(time.Duration(c.MS)*time.Millisecond), c.Out)
	}
	for _, c := range f.Numbers.Grouped {
		check(t, "Grouped", c, Grouped(c.In), c.Out)
	}
}

func TestDurationRoundsToTheSecondAndNeverGoesNegative(t *testing.T) {
	for d, want := range map[time.Duration]string{
		48 * time.Second: "48 s", 192 * time.Second: "3 min 12 s", 65 * time.Minute: "1 h 5 min",
		1500 * time.Millisecond: "2 s", -time.Second: "0 s",
	} {
		check(t, "Duration", d, Duration(d), want)
	}
}

func TestFigureIsWholeFromAHundred(t *testing.T) {
	for _, c := range []struct {
		v, scale float64
		want     string
	}{
		{1234.4, 1234.4, "1,234"}, {99.54, 99.54, "99.5"}, {87.5, 87.5, "87.5"},
		{12.34, 1234, "12"}, {2.25, 87.5, "2.2"}, {-1500, -1500, "-1,500"},
	} {
		check(t, "Figure", c, Figure(c.v, c.scale), c.want)
	}
}

func TestRaidClockNamesTheZone(t *testing.T) {
	chicago := zone(t, "America/Chicago")
	pull := time.Date(2026, 12, 10, 2, 30, 0, 0, time.UTC)
	check(t, "RaidClock", pull, RaidClock(pull, chicago), "8:30 pm CST")
	check(t, "RaidWhen", pull, RaidWhen(pull, chicago), "Wed 9 Dec, 8:30 pm CST")
	check(t, "RaidClock UTC", pull, RaidClock(pull, Location("")), "2:30 am UTC")
}

func TestLocationFallsBackToUTC(t *testing.T) {
	for _, name := range []string{"", "Not/AZone"} {
		if got := Location(name); got != time.UTC {
			t.Errorf("Location(%q) = %v, want UTC", name, got)
		}
	}
	if got := Location("America/Chicago"); got.String() != "America/Chicago" {
		t.Errorf("Location(America/Chicago) = %v", got)
	}
}
