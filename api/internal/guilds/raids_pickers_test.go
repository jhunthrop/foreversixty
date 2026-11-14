package guilds

import "testing"

// A member's raid view carries no picker names; they are an officer's, who frees a wrong pick.
func TestHidePickersClearsEveryRow(t *testing.T) {
	views := []RaidView{{
		Signups:   []SignupView{{Name: "Vela", DiscordPickedBy: "brightfeather"}},
		NotSigned: []RaidCharacter{{Name: "Orrin", DiscordPickedBy: "stonejaw"}},
	}}
	hidePickers(views)
	if views[0].Signups[0].DiscordPickedBy != "" || views[0].NotSigned[0].DiscordPickedBy != "" {
		t.Fatalf("picker names left on a member's view: %+v", views[0])
	}
}
