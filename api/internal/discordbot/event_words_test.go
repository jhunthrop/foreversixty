package discordbot

import (
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
)

// A guild event has no roles anywhere (events spec §3.5; guild round 3 read "In as Healer").
func TestAGuildEventNamesNoRole(t *testing.T) {
	c := guilds.SignupCandidate{CharacterKey: "us/pvp/ashenveil", Name: "Ashenveil", Class: "paladin", Spec: "Holy", Role: "healer"}
	if got := optionDescription(c, true); got != "Holy Paladin" {
		t.Errorf("event option = %q, want no role", got)
	}
	if got := optionDescription(c, false); got != "Holy Paladin · Healer" {
		t.Errorf("raid option = %q", got)
	}
	own := &guilds.NextRaidOwn{CharacterKey: c.CharacterKey, State: guilds.SignupIn, Role: "healer"}
	if got := ownStateSentence(own, []guilds.SignupCandidate{c}, true); got != "Ashenveil is In." {
		t.Errorf("event own state = %q", got)
	}
	if got := ownStateSentence(own, []guilds.SignupCandidate{c}, false); got != "Ashenveil is In as Healer." {
		t.Errorf("raid own state = %q", got)
	}
}
