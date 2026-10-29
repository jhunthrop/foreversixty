package discordbot

// commandCopy is the words of /raid and /help (experience spec §6.2); the tests read each
// string back out of the spec.
var commandCopy = struct {
	// RaidUnlinked answers /raid for an unlinked Discord account; RaidNone when nothing is planned.
	RaidUnlinked, RaidNone string
	// The viewer's state sentences on /raid.
	RaidIn, RaidLate, RaidBench, RaidOut, RaidUnsigned string
	// The same sentences when the signed character cannot be named.
	RaidInNameless, RaidLateNameless, RaidBenchNameless, RaidOutNameless string
	// HelpTitle heads /help; HelpLine is one command's line.
	HelpTitle, HelpLine string
}{
	RaidUnlinked:      "Link your Discord account to see your guild's next raid.",
	RaidNone:          "No raid planned yet.",
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

// The option descriptions that more than one command shares.
const (
	descRaidOption       = "Which raid. The next one if left out."
	descReportRaidOption = "Which raid. The last one if left out."
	descSignupStatus     = "In, Late, Bench or Out. Leave out to pick on a card."
	descName             = "A character name. Yours if left out."
)
