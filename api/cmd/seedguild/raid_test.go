package main

import (
	"testing"
	"time"
)

func TestRaidPlansHaveFourReportsWithinRange(t *testing.T) {
	plans := raidPlans()
	if len(plans) != 4 {
		t.Fatalf("len(plans) = %d, want 4", len(plans))
	}
	for _, p := range plans {
		n := len(p.Fights)
		if n < 6 || n > 10 {
			t.Errorf("%s: %d fights, want 6-10", p.Tag, n)
		}
		if p.wipes() < 1 || p.wipes() > 2 {
			t.Errorf("%s: %d wipes, want 1-2", p.Tag, p.wipes())
		}
		if p.kills() == 0 {
			t.Errorf("%s: no kills at all", p.Tag)
		}
	}
}

func TestRaidPlansTwoAreWithinTheLastWeek(t *testing.T) {
	within7 := 0
	for _, p := range raidPlans() {
		if p.DaysAgo <= 7 {
			within7++
		}
		if p.DaysAgo < 0 || p.DaysAgo > 14 {
			t.Errorf("%s: DaysAgo=%d outside 0-14", p.Tag, p.DaysAgo)
		}
	}
	if within7 != 2 {
		t.Fatalf("reports within the last 7 days = %d, want 2", within7)
	}
}

func TestScheduleOrdersFightsSequentially(t *testing.T) {
	now := time.Date(2026, 10, 4, 12, 0, 0, 0, time.UTC)
	p := raidPlans()[0]
	createdAt, completedAt, starts := p.schedule(now)
	if len(starts) != len(p.Fights) {
		t.Fatalf("len(starts) = %d, want %d", len(starts), len(p.Fights))
	}
	if !starts[0].Equal(createdAt) {
		t.Errorf("first fight start %v != createdAt %v", starts[0], createdAt)
	}
	for i := 1; i < len(starts); i++ {
		if !starts[i].After(starts[i-1]) {
			t.Fatalf("fight %d does not start after fight %d", i, i-1)
		}
	}
	if !completedAt.After(starts[len(starts)-1]) {
		t.Errorf("completedAt %v not after the last fight's start %v", completedAt, starts[len(starts)-1])
	}
}

func TestEncounterIDsMatchMoltenCoreAndOnyxiaZones(t *testing.T) {
	for _, p := range raidPlans() {
		for _, f := range p.Fights {
			if f.EncounterID == 0 {
				continue // trash
			}
			switch p.Zone {
			case zoneMoltenCore:
				if f.EncounterID < encLucifron || f.EncounterID > encMajordomoExecutus {
					t.Errorf("%s: encounter %d outside Molten Core's id range", p.Tag, f.EncounterID)
				}
			case zoneOnyxia:
				if f.EncounterID != encOnyxia {
					t.Errorf("%s: encounter %d is not Onyxia", p.Tag, f.EncounterID)
				}
			default:
				t.Errorf("unexpected zone %q", p.Zone)
			}
		}
	}
}
