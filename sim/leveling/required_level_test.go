package leveling

import "testing"

func TestEffectiveRequiredLevelRaisesButNeverLowers(t *testing.T) {
	cases := []struct {
		name                     string
		itemRequiredLevel, floor int
		want                     int
	}{
		{"no floor leaves required_level alone", 12, 0, 12},
		{"a quest floor above required_level wins", 0, 20, 20}, // Deadhead Blade shape
		{"a quest floor below required_level never lowers it", 30, 10, 30},
		{"equal floor and required_level is a no-op", 20, 20, 20},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := EffectiveRequiredLevel(tc.itemRequiredLevel, tc.floor); got != tc.want {
				t.Errorf("EffectiveRequiredLevel(%d, %d) = %d, want %d", tc.itemRequiredLevel, tc.floor, got, tc.want)
			}
		})
	}
}

func TestLowestFloorIgnoresZeroAndPicksTheEasiestQuest(t *testing.T) {
	cases := []struct {
		name   string
		floors []int
		want   int
	}{
		{"empty", nil, 0},
		{"all zero", []int{0, 0}, 0},
		{"one quest", []int{40}, 40},
		{"lowest of several quests wins - any one suffices", []int{40, 12, 25}, 12},
		{"a zero entry does not erase a real floor", []int{0, 18}, 18},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := LowestFloor(tc.floors); got != tc.want {
				t.Errorf("LowestFloor(%v) = %d, want %d", tc.floors, got, tc.want)
			}
		})
	}
}

func TestItemLevelProxyRequiredLevelMatchesThePythonFormula(t *testing.T) {
	cases := []struct {
		itemLevel, want int
	}{
		{80, 60}, // Polar Leggings: min(60, 75) = 60
		{44, 39},
		{3, 0}, // never negative
		{65, 60},
	}
	for _, tc := range cases {
		if got := ItemLevelProxyRequiredLevel(tc.itemLevel); got != tc.want {
			t.Errorf("ItemLevelProxyRequiredLevel(%d) = %d, want %d", tc.itemLevel, got, tc.want)
		}
	}
}

func TestQuestFloorRaisesTheAcceptLevelToThreeBelowTheQuestLevel(t *testing.T) {
	// Bride of the Embalmer: accept at 20, quest level 30 -> realistic at 27.
	if got := QuestFloor(20, 30); got != 27 {
		t.Fatalf("QuestFloor(20, 30) = %d, want 27", got)
	}
	// Scramble (Forever): accept at 14, quest level 24 -> 21, so its Silver
	// Star never heads a level-20 list.
	if got := QuestFloor(14, 24); got != 21 {
		t.Fatalf("QuestFloor(14, 24) = %d, want 21", got)
	}
	// Wanted: Murkdeep: accept at 15, level 18 -> the accept level still gates.
	if got := QuestFloor(15, 18); got != 15 {
		t.Fatalf("QuestFloor(15, 18) = %d, want 15", got)
	}
	// Unknown quest level: only the accept level.
	if got := QuestFloor(20, 0); got != 20 {
		t.Fatalf("QuestFloor(20, 0) = %d, want 20", got)
	}
}
