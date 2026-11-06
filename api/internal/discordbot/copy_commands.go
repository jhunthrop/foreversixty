package discordbot

import "github.com/jhunthrop/foreversixty/api/internal/guilds"

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
	RaidUnlinked:      "Link your Discord account to see your guild's next raid.",
	RaidNone:          "No raid planned yet.",
	RaidFallback:      "No raid planned yet. Your guild's next {kind}:",
	RaidIn:            "{Name} is In as {Role}.",
	RaidLate:          "{Name} is Late as {Role}.",
	RaidBench:         "{Name} is on the Bench.",
	RaidOut:           "{Name} is Out.",
	RaidUnsigned:      "You have not signed up yet.",
	RaidInNameless:    "You are In as {Role}.",
	RaidLateNameless:  "You are Late as {Role}.",
	RaidBenchNameless: "You are on the Bench.",
	RaidOutNameless:   "You are Out.",
	HelpTitle:         "Forever Sixty commands",
	HelpLine:          "**/{name}** {description}",
}

// kindWords are the site's words for a planned night by its zone kind (web raid-copy.ts,
// guild events spec §3.5): a raid, a dungeon run, a guild event.
var kindWords = map[string]string{
	guilds.ZoneRaid:    "raid",
	guilds.ZoneDungeon: "run",
	guilds.ZoneEvent:   "event",
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
	descRaidOption       = "Which raid. The next one if left out."
	descReportRaidOption = "Which raid. The last one if left out."
	descSignupStatus     = "In, Late, Bench or Out. Leave out to pick on a card."
	descName             = "A character name. Yours if left out."
)
