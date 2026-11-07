package activity

import (
	"os"
	"path/filepath"
	"regexp"
	"testing"
	"time"
)

func TestWeekOfIsTheTuesdayResetAtOrBefore(t *testing.T) {
	cases := map[time.Time]string{
		time.Date(2026, 11, 7, 18, 0, 0, 0, time.UTC):   "2026-11-03", // Saturday
		time.Date(2026, 11, 10, 14, 59, 0, 0, time.UTC): "2026-11-03", // Tuesday before reset
		time.Date(2026, 11, 10, 15, 0, 0, 0, time.UTC):  "2026-11-10", // the reset
		time.Date(2026, 11, 9, 23, 0, 0, 0, time.UTC):   "2026-11-03", // Monday
	}
	for at, want := range cases {
		if got := WeekOf(at).Format("2006-01-02"); got != want {
			t.Errorf("WeekOf(%s) = %s, want %s", at, got, want)
		}
	}
	if PreviousWeek(time.Date(2026, 11, 7, 0, 0, 0, 0, time.UTC)).Format("2006-01-02") != "2026-10-27" {
		t.Fatal("PreviousWeek")
	}
}

func TestTierOfFollowsTheSiteLines(t *testing.T) {
	cases := map[float64]string{0: "S", 4.99: "S", 5: "A", 9.9999999: "B", 10: "B", 19.9: "B", 20: "C", 30: "D", 55: "D"}
	for gap, want := range cases {
		if got := tierOf(gap); got != want {
			t.Errorf("tierOf(%v) = %s, want %s", gap, got, want)
		}
	}
}

// The tier lines are the site's (web/src/lib/tiers/tier-rules.ts TIER_BANDS): read them off
// the file so the two cannot drift.
func TestTierBandsMatchTheSite(t *testing.T) {
	raw, err := os.ReadFile(filepath.Join("..", "..", "..", "web", "src", "lib", "tiers", "tier-rules.ts"))
	if err != nil {
		t.Skip("web checkout absent: ", err)
	}
	found := regexp.MustCompile(`\{ letter: '([SABCD])', from: (\d+) \}`).FindAllStringSubmatch(string(raw), -1)
	if len(found) != len(tierBands) {
		t.Fatalf("site has %d tiers, the API %d", len(found), len(tierBands))
	}
	for i, m := range found {
		if tierBands[i].letter != m[1] || tierBands[i].from != parseFloat(t, m[2]) {
			t.Errorf("tier %d: site %s from %s, API %+v", i, m[1], m[2], tierBands[i])
		}
	}
}

func parseFloat(t *testing.T, s string) float64 {
	t.Helper()
	var f float64
	for _, c := range s {
		f = f*10 + float64(c-'0')
	}
	return f
}

func TestGradeRowsRankEachRoleOnItsOwnFigure(t *testing.T) {
	band := func(spec string, set float64, role string, dtps float64) bisFile {
		b := bisBand{Band: 20, Faction: "horde", Role: role, SetDPS: set}
		if dtps > 0 {
			b.Metrics = &struct {
				DTPS float64 `json:"dtps"`
			}{dtps}
		}
		return bisFile{Spec: spec, Bands: []bisBand{b}}
	}
	rows := gradeRows([]bisFile{
		band("warrior-fury", 100, "", 0), band("rogue-combat", 92, "", 0), band("mage-fire", 70, "", 0),
		band("warrior-protection", 0, "tank", 300), band("paladin-protection", 0, "tank", 250),
	})
	got := map[string]gradeRow{}
	for _, r := range rows {
		got[r.spec] = r
	}
	if *got["warrior-fury"].grade != "S" || *got["rogue-combat"].grade != "A" || *got["mage-fire"].grade != "D" {
		t.Fatalf("dps grades: %+v", got)
	}
	if got["rogue-combat"].place != 2 || got["rogue-combat"].of != 3 {
		t.Fatalf("place: %+v", got["rogue-combat"])
	}
	if got["paladin-protection"].place != 1 || got["warrior-protection"].grade != nil {
		t.Fatalf("tanks sort on damage taken, with no letter: %+v %+v", got["paladin-protection"], got["warrior-protection"])
	}
}

func TestRankingEventsTellATopAndTheClimbsWorthARow(t *testing.T) {
	id := boardID{"warrior-fury", 20}
	board := []BoardRow{{Key: "a", Name: "Grukk", Class: "warrior", Rank: 1}, {Key: "b", Rank: 2}, {Key: "c", Rank: 8},
		{Key: "d", Rank: 9}, {Key: "e", Rank: 12}}
	prev := map[string]int{"b": 1, "a": 2, "c": 14, "d": 12, "e": 30}
	evs := rankingEvents(id, board, prev, testNow)
	kinds := map[string]string{}
	for _, e := range evs {
		kinds[e.CharacterKey] = e.Kind
	}
	// a took #1; b fell; c climbed 6 into the top 10; d only 3 and not into the top 3; e
	// climbed 18 but ended outside the top 10.
	if len(evs) != 2 || kinds["a"] != KindRankingTop || kinds["c"] != KindRankingClimb {
		t.Fatalf("events: %+v", evs)
	}
	if got := evs[0].Segments[1].Text; got != " took #1 for Fury Warrior, 20 to 29" {
		t.Fatalf("top sentence: %q", got)
	}
	if rankingEvents(id, board, nil, testNow) != nil {
		t.Fatal("a board seen for the first time tells no movement")
	}
}
