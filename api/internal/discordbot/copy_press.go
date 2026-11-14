package discordbot

import "github.com/jhunthrop/foreversixty/api/internal/copygen"

// pressCopy is what a press on the raid post, the character select and /signup answer
// (experience spec §3.3.2, §6.2 "Press: …"); the tests read each string back out of the spec.
// {Name}-style markers are filled by fill.
var pressCopy = struct {
	NotLinked string
	// PickFirst answers a Discord-only account with no character picked yet.
	PickFirst string
	// Refusals before the sign-up.
	Locked, Cancelled, Ended string
	NotMember, SetupButton   string
	WaitingOne, WaitingMany  string
	RateLimited, Failure     string
	// The character select.
	SelectPrompt, SelectPlaceholder, OptionDescription, OptionMain, Showing string
	// Outcomes. InSaved plus Count is the spec's In sentence; CountOpen is the count of a raid
	// whose size is unknown. {Date} is the raid's day, "Thu 10 Dec", in the guild's raid zone,
	// so a guild with two raids of one title a week can tell which the press landed on.
	InSaved, Count, CountOpen, LateSaved, NoteButton, BenchSaved, OutSaved, Moved string
	// Already is for In and Late; AlreadyBench and AlreadyOut have no role.
	Already, AlreadyBench, AlreadyOut string
	// NoRaid answers /signup with nothing planned.
	NoRaid string
}{
	NotLinked:         copygen.DiscordPressNotLinked,
	PickFirst:         copygen.DiscordPressPickFirst,
	Locked:            copygen.RaidLocked,
	Cancelled:         copygen.RaidCancelled,
	Ended:             copygen.RaidEnded,
	NotMember:         copygen.DiscordPressNotMember,
	SetupButton:       copygen.SharedHomeSetupAction,
	WaitingOne:        copygen.DiscordPressWaitingOne,
	WaitingMany:       copygen.DiscordPressWaitingMany,
	RateLimited:       copygen.RaidRateLimited,
	Failure:           copygen.FirstPickMainFailed,
	SelectPrompt:      copygen.DiscordPressSelectPrompt,
	SelectPlaceholder: copygen.PushActionChoose,
	OptionDescription: copygen.DiscordPressOptionDescription,
	OptionMain:        copygen.DiscordPressOptionMain,
	Showing:           copygen.DiscordPressShowing,
	InSaved:           copygen.DiscordPressInSaved,
	Count:             copygen.DiscordPressCount,
	CountOpen:         copygen.DiscordPressCountOpen,
	LateSaved:         copygen.DiscordPressLateSaved,
	NoteButton:        copygen.DiscordPressNoteButton,
	BenchSaved:        copygen.DiscordPressBenchSaved,
	OutSaved:          copygen.DiscordPressOutSaved,
	Moved:             copygen.DiscordPressMoved,
	Already:           copygen.DiscordPressAlready,
	AlreadyBench:      copygen.DiscordPressAlreadyBench,
	AlreadyOut:        copygen.DiscordPressAlreadyOut,
	NoRaid:            copygen.RaidsEmpty,
}

// roleWords are the sign-up roles as the sentences say them.
var roleWords = map[string]string{"tank": copygen.RaidRoleTank, "healer": copygen.RaidRoleHealer, "dps": copygen.GuidesRoleDps}

// pickCopy is the sign-up by character pick (design/specs/2026-10-27-discord-bot.md §6.2, the
// Pick rows): a Discord user with no site account picks their character from the guild's
// directory roster under /signup's character option.
var pickCopy = struct {
	// Choice is one autocomplete choice: "Vela Thornbrook · 41 Mage".
	Choice string
	// Claimed leads the first sign-up of a newly picked character.
	Claimed string
	// Taken answers a pick another account already holds.
	Taken string
	// OwnAccount answers a pick asked by an account with a Battle.net sign-in.
	OwnAccount string
	// TooMany answers a pick past directory.MaxDiscordPicks; {n} is the cap.
	TooMany string
}{
	Choice:     copygen.DiscordPickChoice,
	Claimed:    copygen.DiscordPickClaimed,
	Taken:      copygen.DiscordPickTaken,
	OwnAccount: copygen.DiscordPickOwnAccount,
	TooMany:    copygen.DiscordPickTooMany,
}
