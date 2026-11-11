package discordbot

import "github.com/jhunthrop/foreversixty/api/internal/copygen"

// officerCopy is the words of /report, /attendance and /readiness (experience spec §6.2); the
// tests read each string back out of the spec. The G2 night line, kill order and outcome words
// are the raid page's own (web/src/lib/guild/attendance-copy.ts), below.
var officerCopy = struct {
	// ReportPrefix heads a report's title, before the raid's "{Title} · {Weekday} {d} {Month}".
	ReportPrefix string
	// The report's field names.
	KillOrder, MostDeaths, Consumables, Loot string
	// LootAll is the Loot field when every item follows ("4 items"); otherwise LootCount and
	// LootFirst: "10 items" and ", the best 8 below".
	LootAll, LootCount, LootFirst string
	// LootTo is an award's source line.
	LootTo string
	// OpenReport is the report's link button.
	OpenReport string
	// ReportNone and ReportUnlinked answer no report and an unlinked viewer.
	ReportNone, ReportUnlinked string
	// LoggedPresent opens a logged night's line and LoggedNote follows it: a night no planned
	// raid holds, the /report fallback. LoggedPlanned replaces the note when a planned raid of
	// the log's zone waits for a log that night; LoggedPlannedPage is its {page}, the raid page's link.
	LoggedPresent, LoggedNote, LoggedPlanned, LoggedPlannedPage string
	// LootNone is the Loot field of a night with no awards, the site's words.
	LootNone string
	// AttendanceTitle heads a record; LastRaids its field; Roster its link button.
	AttendanceTitle, LastRaids, Roster string
	// OfficersOnly, NoMember, NoRaids and AttendanceUnlinked are /attendance's answers;
	// AttendanceMembers (Task 14's addition) answers a linked viewer with no verified character.
	OfficersOnly, NoMember, NoRaids, AttendanceUnlinked, AttendanceMembers string
	// ReadinessTitle heads the board summary, before the raid's "{Title} · {Weekday} {d} {Month}".
	ReadinessTitle string
	// ReadinessLead and ReadinessLeadOne say how many signed and how many have something to fix.
	ReadinessLead, ReadinessLeadOne, ReadinessAllReady string
	// ReadinessNobodySigned (the controller's addition) is a raid nobody signed In or Late for.
	ReadinessNobodySigned string
	// ToFix is the failing rows' field; ReadinessMore follows the tenth row; Board the link.
	ToFix, ReadinessMore, Board string
	// The refusals: a member, an unlinked viewer, a frozen claim.
	NotOfficer, ReadinessUnlinked, Frozen string
	// RaidGuild and MoveToRaidGuild (Task 14's addition) answer a Guild-tier guild: the board is
	// part of Raid guild (launch pricing §3.2.8, the site's GuildPlanGate sentence).
	RaidGuild, MoveToRaidGuild string
	// GrantedMove and AskUsOnDiscord answer a Guild plan Forever Sixty granted: no Move to
	// buy, so it asks us (the site's guildGateCopy.grantedMove, word for word).
	GrantedMove, AskUsOnDiscord string
}{
	ReportPrefix:          copygen.DiscordOfficerReportPrefix,
	KillOrder:             copygen.DiscordOfficerKillOrder,
	MostDeaths:            copygen.DiscordOfficerMostDeaths,
	Consumables:           copygen.DiscordOfficerConsumables,
	Loot:                  copygen.DiscordOfficerLoot,
	LootAll:               copygen.DiscordOfficerLootAll,
	LootCount:             copygen.DiscordOfficerLootAll,
	LootFirst:             copygen.DiscordOfficerLootFirst,
	LootTo:                copygen.DiscordOfficerLootTo,
	OpenReport:            copygen.RankMenuReport,
	ReportNone:            copygen.DiscordOfficerReportNone,
	LoggedPresent:         copygen.DiscordOfficerLoggedPresent,
	LoggedNote:            copygen.DiscordOfficerLoggedNote,
	LoggedPlanned:         copygen.DiscordOfficerLoggedPlanned,
	LoggedPlannedPage:     copygen.DiscordOfficerLoggedPlannedPage,
	LootNone:              copygen.DiscordOfficerLootNone,
	ReportUnlinked:        copygen.DiscordOfficerReportUnlinked,
	AttendanceTitle:       copygen.DiscordOfficerAttendanceTitle,
	LastRaids:             copygen.DiscordOfficerLastRaids,
	Roster:                copygen.RaidsRosterRow,
	OfficersOnly:          copygen.DiscordOfficerOfficersOnly,
	NoMember:              copygen.DiscordOfficerNoMember,
	NoRaids:               copygen.DiscordOfficerNoRaids,
	AttendanceUnlinked:    copygen.DiscordOfficerAttendanceUnlinked,
	AttendanceMembers:     copygen.DiscordOfficerAttendanceMembers,
	ReadinessTitle:        copygen.DiscordOfficerReadinessTitle,
	ReadinessLead:         copygen.DiscordOfficerReadinessLead,
	ReadinessLeadOne:      copygen.DiscordOfficerReadinessLeadOne,
	ReadinessAllReady:     copygen.DiscordOfficerReadinessAllReady,
	ReadinessNobodySigned: copygen.DiscordOfficerReadinessNobodySigned,
	ToFix:                 copygen.DiscordOfficerToFix,
	ReadinessMore:         copygen.DiscordOfficerReadinessMore,
	Board:                 copygen.DiscordOfficerBoard,
	NotOfficer:            copygen.DiscordOfficerNotOfficer,
	ReadinessUnlinked:     copygen.DiscordOfficerReadinessUnlinked,
	Frozen:                copygen.DiscordOfficerFrozen,
	RaidGuild:             copygen.DiscordOfficerRaidGuild,
	MoveToRaidGuild:       copygen.DiscordOfficerMoveToRaidGuild,
	GrantedMove:           copygen.DiscordOfficerGrantedMove,
	AskUsOnDiscord:        copygen.DiscordOfficerAskUsOnDiscord,
}

// CommunityDiscordInvite is the Forever Sixty Discord invite, web/src/data/links.json's
// discordInvite (a test holds the two equal).
const CommunityDiscordInvite = "https://discord.gg/NZhBtPwvHt"

// nightCopy is the raid page's night and kill-order words (G2, attendance-copy.ts): the night
// line "22 of 23 came · 1 kill · 3 wipes · 14 deaths · 2 h 41 min" and a kill row
// "1. Onyxia 1:42 after 3 wipes".
var nightCopy = struct {
	Came, FirstPull, After, NotKilled string
	// CameUnsigned and NobodySigned open the line when nobody signed up beforehand.
	CameUnsigned, NobodySigned string
	// ShortBefore and NoFlask are the Consumables clauses; NothingShared stands for both when
	// nobody shared their bags and the log has no buffs to read.
	ShortBefore, NoFlask, NothingShared string
}{
	Came:          copygen.DiscordNightCame,
	CameUnsigned:  copygen.DiscordNightCameUnsigned,
	NobodySigned:  copygen.DiscordNightNobodySigned,
	NothingShared: copygen.DiscordNightNothingShared,
	FirstPull:     copygen.DiscordNightFirstPull,
	After:         copygen.DiscordNightAfter,
	NotKilled:     copygen.DiscordNightNotKilled,
	ShortBefore:   copygen.DiscordNightShortBefore,
	NoFlask:       copygen.DiscordNightNoFlask,
}

// outcomeWords are the attendance outcomes in G2's words (never "no-show").
var outcomeWords = map[string]string{
	"present": copygen.DiscordOutcomeWordsPresent,
	"late":    copygen.IntentStateLate,
	"no_show": copygen.DiscordOutcomeWordsNoShow,
}

// The attendance line's extra words: "Came as {Name}" and the excused suffix.
const (
	cameAsWords   = copygen.DiscordOfficerCameAsWords
	excusedSuffix = copygen.DiscordOfficerExcusedSuffix
)
