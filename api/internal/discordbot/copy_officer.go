package discordbot

// officerCopy is the words of /report, /attendance and /readiness (experience spec §6.2); the
// tests read each string back out of the spec. The G2 night line, kill order and outcome words
// are the raid page's own (web/src/lib/guild/attendance-copy.ts), below.
var officerCopy = struct {
	// ReportPrefix heads a report's title, before the raid's "{Title} · {Weekday} {d} {Month}".
	ReportPrefix string
	// The report's field names.
	KillOrder, MostDeaths, Consumables, Loot string
	// LootCount and LootFirst are the Loot field: "6 items." and, when items follow, " The first 6:".
	LootCount, LootFirst string
	// LootTo is an award's source line.
	LootTo string
	// OpenReport is the report's link button.
	OpenReport string
	// ReportNone and ReportUnlinked answer no report and an unlinked viewer.
	ReportNone, ReportUnlinked string
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
	// The refusals: a member, an unlinked viewer, a lapsed plan (with Subscribe), a frozen claim.
	NotOfficer, ReadinessUnlinked, PlanLapsed, Subscribe, Frozen string
	// RaidGuild and MoveToRaidGuild (Task 14's addition) answer a Guild-tier guild: the board is
	// part of Raid guild (launch pricing §3.2.8, the site's GuildPlanGate sentence).
	RaidGuild, MoveToRaidGuild string
}{
	ReportPrefix:          "Report: ",
	KillOrder:             "Kill order",
	MostDeaths:            "Most deaths",
	Consumables:           "Consumables",
	Loot:                  "Loot",
	LootCount:             "{n} items.",
	LootFirst:             " The first {k}:",
	LootTo:                "to {Name} · {Boss}",
	OpenReport:            "Open the report",
	ReportNone:            "No raid logged yet. A raid shows here once its log is matched.",
	ReportUnlinked:        "Link your Discord account to see your guild's raids.",
	AttendanceTitle:       "{Name} · attendance",
	LastRaids:             "Last 5 raids",
	Roster:                "Roster",
	OfficersOnly:          "Only officers can look up another member's attendance.",
	NoMember:              "No member of {Guild} by that name.",
	NoRaids:               "No raids yet.",
	AttendanceUnlinked:    "Link your Discord account to see your attendance.",
	AttendanceMembers:     "Attendance is for members of {Guild}.",
	ReadinessTitle:        "Readiness · ",
	ReadinessLead:         "{n} signed In or Late. {k} have something to fix.",
	ReadinessLeadOne:      "{n} signed In or Late. 1 has something to fix.",
	ReadinessAllReady:     "Everyone signed In or Late is ready.",
	ReadinessNobodySigned: "No one has signed In or Late yet.",
	ToFix:                 "To fix, most checks first",
	ReadinessMore:         "+{n} more on the board",
	Board:                 "Readiness board",
	NotOfficer:            "Readiness is for officers of {Guild}.",
	ReadinessUnlinked:     "Readiness is for officers of {Guild}. Link your Discord account if you are one.",
	PlanLapsed:            "Planning and running raids are part of the guild plan. An officer can subscribe the guild.",
	Subscribe:             "Subscribe the guild",
	Frozen:                "This guild's claim is contested. Officer actions are frozen until a moderator resolves it.",
	RaidGuild:             "The readiness board is part of Raid guild. An officer can move the guild to Raid guild.",
	MoveToRaidGuild:       "Move to Raid guild",
}

// nightCopy is the raid page's night and kill-order words (G2, attendance-copy.ts): the night
// line "22 of 23 came · 1 kill · 3 wipes · 14 deaths · 2 h 41 min" and a kill row
// "1. Onyxia 1:42 after 3 wipes".
var nightCopy = struct {
	Came, FirstPull, After, NotKilled string
	// ShortBefore and NoFlask are the Consumables clauses.
	ShortBefore, NoFlask string
}{
	Came:        "{came} of {signed} came",
	FirstPull:   "first pull",
	After:       "after {wipes}",
	NotKilled:   "Not killed",
	ShortBefore: "{n} short before the raid",
	NoFlask:     "{n} with no flask in the log",
}

// outcomeWords are the attendance outcomes in G2's words (never "no-show").
var outcomeWords = map[string]string{
	"present": "Present",
	"late":    "Late",
	"no_show": "Not in the log",
}

// The attendance line's extra words: "Came as {Name}" and the excused suffix.
const (
	cameAsWords   = "Came as {Name}"
	excusedSuffix = ", excused"
)
