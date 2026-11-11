package discordbot

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/copygen"
	"github.com/jhunthrop/foreversixty/api/internal/upgrades"
)

// Discord's option type numbers the bot uses.
const optionString = 3

// classChoices are the nine classes; slotChoices the 17 gear slots, valued by the BiS data's
// own slot keys (bis.Slot.Slot); phaseChoices the level bands and the Phase 1 raid.
var (
	classChoices = choicesOf("Druid", "Hunter", "Mage", "Paladin", "Priest", "Rogue", "Shaman", "Warlock", "Warrior")
	slotChoices  = slotChoicesOf(upgrades.Slots)
	phaseChoices = []Choice{
		{copygen.DiscordCommandsPhaseLevel20, "20"}, {copygen.DiscordCommandsPhaseLevel30, "30"}, {copygen.DiscordCommandsPhaseLevel40, "40"}, {copygen.DiscordCommandsPhaseLevel50, "50"},
		{copygen.DiscordCommandsPhaseLevel60, "60"}, {copygen.DiscordCommandsPhaseRaid, "p1"},
	}
	statusChoices = []Choice{{"In", "in"}, {"Late", "late"}, {"Bench", "bench"}, {"Out", "out"}}
)

// slotChoicesOf is the gear slots as choices, named by their words and valued by their keys.
func slotChoicesOf(slots []upgrades.Slot) []Choice {
	out := make([]Choice, len(slots))
	for i, s := range slots {
		out[i] = Choice{Name: s.Label, Value: s.Key}
	}
	return out
}

func choicesOf(names ...string) []Choice {
	out := make([]Choice, len(names))
	for i, n := range names {
		out[i] = Choice{Name: n, Value: strings.ToLower(n)}
	}
	return out
}

func stringOption(name, description string, required, autocomplete bool, choices []Choice) CommandOption {
	return CommandOption{Name: name, Description: description, Type: optionString,
		Required: required, Autocomplete: autocomplete, Choices: choices}
}

// Commands is every slash command in the spec's order. The released ones register globally;
// the rest register to the development server until they ship.
func Commands() []Command {
	return []Command{
		{Name: "raid", Description: copygen.DiscordCommandsRaid, Released: true},
		{Name: "signup", Description: copygen.DiscordCommandsSignup, Released: true, Options: []CommandOption{
			stringOption("status", descSignupStatus, false, false, statusChoices),
			stringOption("raid", descRaidOption, false, true, nil),
			stringOption("character", copygen.DiscordCommandsSignupCharacter, false, true, nil),
		}},
		{Name: "char", Description: copygen.DiscordCommandsChar, Released: true, Options: []CommandOption{
			stringOption("name", descName, false, true, nil),
		}},
		{Name: "bis", Description: copygen.DiscordCommandsBis, Released: true, Options: []CommandOption{
			stringOption("class", copygen.DiscordCommandsBisClass, false, false, classChoices),
			stringOption("spec", copygen.DiscordCommandsBisSpec, false, true, nil),
			stringOption("slot", copygen.DiscordCommandsBisSlot, false, false, slotChoices),
			stringOption("phase", copygen.DiscordCommandsBisPhase, false, false, phaseChoices),
		}},
		{Name: "report", Description: copygen.DiscordCommandsReport, Released: true, Options: []CommandOption{
			stringOption("raid", descReportRaidOption, false, true, nil),
		}},
		{Name: "attendance", Description: copygen.DiscordCommandsAttendance, Released: true, Options: []CommandOption{
			stringOption("name", copygen.DiscordCommandsAttendanceName, false, true, nil),
		}},
		{Name: "readiness", Description: copygen.DiscordCommandsReadiness, Released: true, Options: []CommandOption{
			stringOption("raid", descRaidOption, false, true, nil),
		}},
		{Name: "link", Description: copygen.DiscordCommandsLink, Released: true},
		{Name: "help", Description: copygen.DiscordCommandsHelp, Released: true},
	}
}

// commandsOf is the commands whose Released flag equals released.
func commandsOf(cmds []Command, released bool) []Command {
	var out []Command
	for _, c := range cmds {
		if c.Released == released {
			out = append(out, c)
		}
	}
	return out
}

// commandsHashVersion is bumped to make every scope register again (commands deleted in the
// Developer Portal, a change to how they are sent).
const commandsHashVersion = "v1"

// commandsHash is the sha256 of the commands' canonical JSON, under the hash version: what
// Discord is sent.
func commandsHash(cmds []Command) (string, error) {
	raw, err := json.Marshal(cmds)
	if err != nil {
		return "", fmt.Errorf("discordbot: marshal commands: %w", err)
	}
	sum := sha256.Sum256(append([]byte(commandsHashVersion+"\n"), raw...))
	return hex.EncodeToString(sum[:]), nil
}
