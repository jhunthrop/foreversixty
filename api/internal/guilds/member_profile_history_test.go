package guilds

import (
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/bis"
)

func profileNightOn(day int, outcome string, excused bool) ProfileNight {
	n := ProfileNight{StartsAt: time.Date(2026, 12, day, 2, 0, 0, 0, time.UTC), Excused: excused}
	if outcome != "" {
		n.Outcome = &outcome
	}
	return n
}

func TestAttendanceStreak(t *testing.T) {
	cases := []struct {
		name   string
		nights []ProfileNight
		want   int
	}{
		{"none", nil, 0},
		{"every night", []ProfileNight{profileNightOn(16, "present", false), profileNightOn(9, "late", false)}, 2},
		{"a miss ends it", []ProfileNight{profileNightOn(16, "present", false), profileNightOn(9, "absent", false), profileNightOn(2, "present", false)}, 1},
		{"excused is skipped", []ProfileNight{profileNightOn(16, "present", false), profileNightOn(9, "absent", true), profileNightOn(2, "present", false)}, 2},
		{"no row ends it", []ProfileNight{profileNightOn(16, "", false), profileNightOn(9, "present", false)}, 0},
	}
	for _, c := range cases {
		if got := attendanceStreak(c.nights); got != c.want {
			t.Errorf("%s: streak = %d, want %d", c.name, got, c.want)
		}
	}
}

func TestNightsSinceAward(t *testing.T) {
	nights := []ProfileNight{profileNightOn(16, "present", false), profileNightOn(9, "absent", false), profileNightOn(2, "late", false)}
	awarded := time.Date(2026, 12, 2, 3, 0, 0, 0, time.UTC)
	if got := nightsSinceAward(nights, &awarded); got != 1 {
		t.Errorf("since an award on the 2nd = %d, want 1 (the 9th was missed)", got)
	}
	if got := nightsSinceAward(nights, nil); got != 2 {
		t.Errorf("with no award = %d, want every night they came, 2", got)
	}
}

func TestMedianOf(t *testing.T) {
	cases := []struct {
		values []int
		want   int
	}{{nil, 0}, {[]int{3}, 3}, {[]int{0, 2, 1}, 1}, {[]int{0, 1, 3, 2}, 1}}
	for _, c := range cases {
		if got := medianOf(c.values); got != c.want {
			t.Errorf("medianOf(%v) = %d, want %d", c.values, got, c.want)
		}
	}
}

func TestProfileUpgradesLeadWithTheBiggestKnownGain(t *testing.T) {
	gain := func(v float64) *float64 { return &v }
	got := profileUpgrades([]bis.GapPiece{
		{Slot: "head", ItemID: 1, ItemName: "Helm", Source: "Onyxia"},
		{Slot: "back", ItemID: 2, ItemName: "Drape", Source: "Onyxia", GainDps: gain(4.2)},
		{Slot: "hands", ItemID: 3, ItemName: "Gloves", Source: "Lucifron", GainDps: gain(11.7)},
	})
	want := []string{"Gloves", "Drape", "Helm"}
	for i, name := range want {
		if got[i].Name != name {
			t.Fatalf("order = %+v, want %v", got, want)
		}
	}
	if got[2].GainDps != nil || got[0].Source != "Lucifron" {
		t.Fatalf("pieces = %+v", got)
	}
	if len(profileUpgrades(nil)) != 0 || profileUpgrades(nil) == nil {
		t.Fatal("no gap reads an empty list, never null")
	}
}
