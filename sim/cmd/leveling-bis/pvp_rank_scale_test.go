package main

import "testing"

// Round 11: loot.json's pvp `rank` and `title_rank` are the client's
// RequiredPVPRank (5 to 18), which is the in-game honor rank (1 to 14)
// plus 4: field 10 is Knight / Stone Guard, rank 6; field 14 is
// Lieutenant Commander / Champion, rank 10; field 18 is Grand Marshal /
// High Warlord, rank 14 (data/pipeline/loot/pvp_faction.py's ladder,
// warcraft.wiki.gg "Honor system (pre-2.0)"). The owner's cap is in-game
// rank 10, and the label names the in-game rank.
func TestThePvpRankCapReadsTheInGameRank(t *testing.T) {
	cases := []struct {
		name  string
		field int
		want  bool
	}{
		{"field 10, Knight / Stone Guard (rank 6, the cap)", 10, false},
		{"field 11, Knight-Lieutenant / Blood Guard (rank 7)", 11, true},
		{"field 14, Lieutenant Commander / Champion (rank 10)", 14, true},
		{"field 15, Commander / Lieutenant General (rank 11)", 15, true},
		{"field 18, Grand Marshal / High Warlord (rank 14)", 18, true},
	}
	for _, tc := range cases {
		if got := pvpRankExceedsCap(itemSource{Kind: "pvp", Rank: tc.field}); got != tc.want {
			t.Errorf("%s: rank exceeds cap = %v, want %v", tc.name, got, tc.want)
		}
		if got := pvpRankExceedsCap(itemSource{Kind: "pvp", TitleRank: tc.field}); got != tc.want {
			t.Errorf("%s as title_rank: exceeds cap = %v, want %v", tc.name, got, tc.want)
		}
	}
}

func TestThePvpLabelNamesTheInGameRank(t *testing.T) {
	cases := []struct {
		field          int
		title, faction string
		want           string
	}{
		{18, "Grand Marshal", "alliance", "PvP rank 14 · Grand Marshal · Alliance"},
		{10, "Stone Guard", "horde", "PvP rank 6 · Stone Guard · Horde"},
		{8, "Master Sergeant", "alliance", "PvP rank 4 · Master Sergeant · Alliance"},
	}
	for _, tc := range cases {
		if got := pvpSourceLabel(tc.field, tc.title, tc.faction); got != tc.want {
			t.Errorf("pvpSourceLabel(%d, %q) = %q, want %q", tc.field, tc.title, got, tc.want)
		}
	}
}
