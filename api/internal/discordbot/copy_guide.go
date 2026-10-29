package discordbot

// guideStrings is the guided cards' copy (spec 2026-10-28 §G.7). Bold is Discord markdown.
// {Class}, {Spec}, {Name}, {Guild}, {typed}, {Slot} and {command} are escaped member data.
type guideStrings struct {
	BisAuthor                 string
	BisTitleNone              string
	BisTitleClass             string
	BisTitleReady             string
	BisLeadUnlinked           string
	BisLeadNone               string
	BisLeadMain               string
	BisLeadClass              string
	BisLeadReady              string
	BisLeadSlot               string
	BisLeadUnknownSpec        string
	BisLeadUnknownSpecNoClass string
	BisNudgeSpec              string
	BisNudgeBoth              string
	BisPlaceholderClass       string
	BisPlaceholderSpecLocked  string
	BisPlaceholderSpec        string
	BisPlaceholderRange       string
	// BisRangeDesc is one line per phaseChoices entry, in order.
	BisRangeDesc [6]string
	// BisRole is DPS, Tank, Healer.
	BisRole    [3]string
	BisShow    string
	BisUseMain string
	BisChange  string

	CharLeadUnlinked string
	CharLeadNone     string
	CharUnknown      string
	CharSomeoneElse  string
	CharTypeName     string
	CharSelect       string
	CharSelectDesc   string

	ModalCharTitle       string
	ModalAttendanceTitle string
	ModalLabel           string
	ModalPlaceholder     string

	SignupPlaceholderRaid      string
	SignupRaidLabel            string
	SignupRaidDesc             string
	SignupRaidDescNoSize       string
	SignupPlaceholderCharacter string

	AttendancePlaceholder string
	AttendanceUnknown     string
	ReadinessPlaceholder  string
	ReportPlaceholder     string

	CardStale   string
	CardFailure string

	// The spec card, "Which spec tonight?" (spec 2026-10-28 §P.10). SpecKept starts with a space:
	// it is appended to the outcome sentence (alone, the space is trimmed).
	SpecTitle       string
	SpecLead        string
	SpecLeadChange  string
	SpecPlaceholder string
	SpecKept        string
	SpecChange      string
	SpecStale       string
}

var guideCopy = guideStrings{
	BisAuthor:                 "Best in slot",
	BisTitleNone:              "Choose a class",
	BisTitleClass:             "{Class} · choose a spec",
	BisTitleReady:             "{Spec} {Class} · {band label}",
	BisLeadUnlinked:           "Choose a class, a spec and a range, then Show BiS. Link Discord and this starts from your character.",
	BisLeadNone:               "Choose a class, a spec and a range, then Show BiS.",
	BisLeadMain:               "Picked from your main, **{Name}**. Show BiS, or change a pick.",
	BisLeadClass:              "Now choose a {Class} spec.",
	BisLeadReady:              "Ready. Show BiS, or change a pick.",
	BisLeadSlot:               "For {Slot} only. Choose a class, a spec and a range, then Show BiS.",
	BisLeadUnknownSpec:        "No {Class} spec by that name. Choose one.",
	BisLeadUnknownSpecNoClass: "No spec by that name. Choose one.",
	BisNudgeSpec:              "**Choose a spec first.** Then Show BiS.",
	BisNudgeBoth:              "**Choose a class and a spec first.** Then Show BiS.",
	BisPlaceholderClass:       "Choose a class",
	BisPlaceholderSpecLocked:  "Choose a class first",
	BisPlaceholderSpec:        "Choose a spec",
	BisPlaceholderRange:       "Choose a range",
	BisRangeDesc: [6]string{
		"Levels 20 to 29", "Levels 30 to 39", "Levels 40 to 49", "Levels 50 to 59",
		"Level 60, no buffs or consumables", "Level 60, raid buffs and consumables",
	},
	BisRole:    [3]string{"DPS", "Tank", "Healer"},
	BisShow:    "Show BiS",
	BisUseMain: "Use {Name}",
	BisChange:  "Change picks",

	CharLeadUnlinked: "Which character? Type a name, or link Discord to start from yours.",
	CharLeadNone:     "Which character? Type a name.",
	CharUnknown:      "No public character named {typed}. Try another name.",
	CharSomeoneElse:  "Someone else",
	CharTypeName:     "Type a name",
	CharSelect:       "Another of your characters",
	CharSelectDesc:   "{Spec} {Class} · Level {n}",

	ModalCharTitle:       "Look up a character",
	ModalAttendanceTitle: "Look up attendance",
	ModalLabel:           "Character name",
	ModalPlaceholder:     "Zulmara",

	SignupPlaceholderRaid:      "Choose a raid",
	SignupRaidLabel:            "{Title} · {Ddd d Mmm}",
	SignupRaidDesc:             "{Guild time} · {signed} of {size} signed up",
	SignupRaidDescNoSize:       "{n} signed up",
	SignupPlaceholderCharacter: "Choose a character",

	AttendancePlaceholder: "Look up a member",
	AttendanceUnknown:     "No member of {Guild} named {typed}.",
	ReadinessPlaceholder:  "Another raid",
	ReportPlaceholder:     "Another night",

	CardStale:   "This card is out of date. Run /{command} again.",
	CardFailure: commonCopy.Failure,

	SpecTitle:       "Which spec tonight?",
	SpecLead:        "The pick signs **{Name}** up as {State} for {Title}, {Date}, and is kept for the next raid.",
	SpecLeadChange:  "The pick changes the spec **{Name}** brings to {Title}, {Date}, and is kept for the next raid.",
	SpecPlaceholder: "Choose a spec",
	SpecKept:        " {Spec} is kept for {Name}.",
	SpecChange:      "Change spec",
	SpecStale:       "This card is out of date. Press the button on the raid post again.",
}
