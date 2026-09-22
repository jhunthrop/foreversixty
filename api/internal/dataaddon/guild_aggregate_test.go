// api/internal/dataaddon/guild_aggregate_test.go
package dataaddon

import "testing"

func TestMemberNameOfReadsTheNameSegmentOfACharacterKey(t *testing.T) {
	if got := memberNameOf("us/normal/thoradin"); got != "thoradin" {
		t.Errorf("memberNameOf = %q, want thoradin", got)
	}
	if got := memberNameOf("not-a-character-key"); got != "not-a-character-key" {
		t.Errorf("memberNameOf on an unparseable key should return it unchanged, got %q", got)
	}
}

func TestBuildGuildRowSortsMembersAndComputesProgress(t *testing.T) {
	g := guildIdentity{ID: 1, Region: "us", Ruleset: "normal", Name: "Iron Vanguard"}
	row := buildGuildRow(g, []string{"us/normal/thoradin", "us/normal/o'malley"}, 2, 2, 3)
	if row.Name != "Iron Vanguard" {
		t.Errorf("name = %q", row.Name)
	}
	if row.Progress != "2/3" {
		t.Errorf("progress = %q, want 2/3", row.Progress)
	}
	if row.Nights != 2 {
		t.Errorf("nights = %d, want 2", row.Nights)
	}
	if row.Roster != 2 {
		t.Errorf("roster = %d, want 2", row.Roster)
	}
	want := []string{"o'malley", "thoradin"}
	for i, name := range want {
		if row.Members[i] != name {
			t.Errorf("members[%d] = %q, want %q", i, row.Members[i], name)
		}
	}
}

func TestBuildGuildRowOmitsProgressWithNoAttemptedEncounters(t *testing.T) {
	g := guildIdentity{ID: 1, Region: "us", Ruleset: "normal", Name: "Iron Vanguard"}
	row := buildGuildRow(g, []string{"us/normal/thoradin"}, 0, 0, 0)
	if row.Progress != "" {
		t.Errorf("progress = %q, want empty (omitted)", row.Progress)
	}
}
