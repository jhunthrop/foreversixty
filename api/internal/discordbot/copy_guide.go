package discordbot

import "github.com/jhunthrop/foreversixty/api/internal/copygen"

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
	BisAuthor:                 copygen.BisTitle,
	BisTitleNone:              copygen.DiscordGuideBisTitleNone,
	BisTitleClass:             copygen.DiscordGuideBisTitleClass,
	BisTitleReady:             copygen.DiscordGuideBisTitleReady,
	BisLeadUnlinked:           copygen.DiscordGuideBisLeadUnlinked,
	BisLeadNone:               "",
	BisLeadMain:               copygen.DiscordGuideBisLeadMain,
	BisLeadClass:              copygen.DiscordGuideBisLeadClass,
	BisLeadReady:              copygen.DiscordGuideBisLeadReady,
	BisLeadSlot:               copygen.DiscordGuideBisLeadSlot,
	BisLeadUnknownSpec:        copygen.DiscordGuideBisLeadUnknownSpec,
	BisLeadUnknownSpecNoClass: copygen.DiscordGuideBisLeadUnknownSpecNoClass,
	BisNudgeSpec:              copygen.DiscordGuideBisNudgeSpec,
	BisNudgeBoth:              copygen.DiscordGuideBisNudgeBoth,
	BisPlaceholderClass:       copygen.DiscordGuideBisTitleNone,
	BisPlaceholderSpecLocked:  copygen.DiscordGuideBisPlaceholderSpecLocked,
	BisPlaceholderSpec:        copygen.WidgetChooseSpec,
	BisPlaceholderRange:       copygen.DiscordGuideBisPlaceholderRange,
	BisRangeDesc: [6]string{
		copygen.DiscordGuideBisRangeDesc1, copygen.DiscordGuideBisRangeDesc2, copygen.DiscordGuideBisRangeDesc3, copygen.DiscordGuideBisRangeDesc4,
		copygen.DiscordGuideBisRangeDesc5, copygen.DiscordGuideBisRangeDesc6,
	},
	BisRole:    [3]string{copygen.GuidesRoleDps, copygen.RaidRoleTank, copygen.RaidRoleHealer},
	BisShow:    copygen.DiscordGuideBisShow,
	BisUseMain: copygen.DiscordGuideBisUseMain,
	BisChange:  copygen.DiscordGuideBisChange,

	CharLeadUnlinked: copygen.DiscordGuideCharLeadUnlinked,
	CharLeadNone:     copygen.DiscordGuideCharLeadNone,
	CharUnknown:      copygen.DiscordGuideCharUnknown,
	CharSomeoneElse:  copygen.DiscordGuideCharSomeoneElse,
	CharTypeName:     copygen.DiscordGuideCharTypeName,
	CharSelect:       copygen.DiscordGuideCharSelect,
	CharSelectDesc:   copygen.DiscordGuideCharSelectDesc,

	ModalCharTitle:       copygen.DiscordGuideModalCharTitle,
	ModalAttendanceTitle: copygen.DiscordGuideModalAttendanceTitle,
	ModalLabel:           copygen.DiscordGuideModalLabel,
	ModalPlaceholder:     copygen.DiscordGuideModalPlaceholder,

	SignupPlaceholderRaid:      copygen.DiscordGuideSignupPlaceholderRaid,
	SignupRaidLabel:            copygen.DiscordGuideSignupRaidLabel,
	SignupRaidDesc:             copygen.DiscordGuideSignupRaidDesc,
	SignupRaidDescNoSize:       copygen.DiscordGuideSignupRaidDescNoSize,
	SignupPlaceholderCharacter: copygen.PushActionChoose,

	AttendancePlaceholder: copygen.DiscordGuideAttendancePlaceholder,
	AttendanceUnknown:     copygen.DiscordGuideAttendanceUnknown,
	ReadinessPlaceholder:  copygen.DiscordGuideReadinessPlaceholder,
	ReportPlaceholder:     copygen.DiscordGuideReportPlaceholder,

	CardStale:   copygen.DiscordGuideCardStale,
	CardFailure: commonCopy.Failure,

	SpecTitle:       copygen.SpecTitle,
	SpecLead:        copygen.DiscordGuideSpecLead,
	SpecLeadChange:  copygen.DiscordGuideSpecLeadChange,
	SpecPlaceholder: copygen.WidgetChooseSpec,
	SpecKept:        copygen.DiscordGuideSpecKept,
	SpecChange:      copygen.SpecChange,
	SpecStale:       copygen.DiscordGuideSpecStale,
}
