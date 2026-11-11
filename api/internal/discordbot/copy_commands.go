package discordbot

import (
	"github.com/jhunthrop/foreversixty/api/internal/copygen"
	"github.com/jhunthrop/foreversixty/api/internal/guilds"
)

// commandCopy is the words of /raid and /help (experience spec §6.2); the tests read each
// string back out of the spec.
var commandCopy = struct {
	// RaidUnlinked answers /raid for an unlinked Discord account; RaidNone when nothing is planned.
	RaidUnlinked, RaidNone string
	// RaidFallback leads the answer when no raid is planned and the next night is a dungeon run
	// or a guild event; {kind} is its kind word from kindWords.
	RaidFallback string
	// The viewer's state sentences on /raid.
	RaidIn, RaidLate, RaidBench, RaidOut, RaidUnsigned string
	// The same sentences when the signed character cannot be named.
	RaidInNameless, RaidLateNameless, RaidBenchNameless, RaidOutNameless string
	// HelpTitle heads /help; HelpLine is one command's line.
	HelpTitle, HelpLine string
}{
	RaidUnlinked:      copygen.DiscordCommandRaidUnlinked,
	RaidNone:          copygen.RaidsEmpty,
	RaidFallback:      copygen.DiscordCommandRaidFallback,
	RaidIn:            copygen.DiscordCommandRaidIn,
	RaidLate:          copygen.DiscordCommandRaidLate,
	RaidBench:         copygen.DiscordCommandRaidBench,
	RaidOut:           copygen.DiscordCommandRaidOut,
	RaidUnsigned:      copygen.DiscordCommandRaidUnsigned,
	RaidInNameless:    copygen.DiscordCommandRaidInNameless,
	RaidLateNameless:  copygen.DiscordCommandRaidLateNameless,
	RaidBenchNameless: copygen.DiscordCommandRaidBenchNameless,
	RaidOutNameless:   copygen.DiscordCommandRaidOutNameless,
	HelpTitle:         copygen.DiscordCommandHelpTitle,
	HelpLine:          copygen.DiscordCommandHelpLine,
}

// kindWords are the site's words for a planned night by its zone kind (web raid-copy.ts,
// guild events spec §3.5): a raid, a dungeon run, a guild event.
var kindWords = map[string]string{
	guilds.ZoneRaid:    copygen.DiscordKindWordsZoneRaid,
	guilds.ZoneDungeon: copygen.DiscordKindWordsZoneDungeon,
	guilds.ZoneEvent:   copygen.DiscordKindWordsZoneEvent,
}

// kindWord is kindWords' word for a zone kind, "raid" for a kind it does not name.
func kindWord(zoneKind string) string {
	if word, ok := kindWords[zoneKind]; ok {
		return word
	}
	return kindWords[guilds.ZoneRaid]
}

// The option descriptions that more than one command shares.
const (
	descRaidOption       = copygen.DiscordCommandsDescRaidOption
	descReportRaidOption = copygen.DiscordCommandsDescReportRaidOption
	descSignupStatus     = copygen.DiscordCommandsDescSignupStatus
	descName             = copygen.DiscordCommandsDescName
)
