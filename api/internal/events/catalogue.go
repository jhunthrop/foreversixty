// Package events is the player's notification catalogue as typed events
// (docs/superpowers/plans/2026-10-30-mobile-apps-lanes.md, "The player's lifecycle and its
// notifications"): one Kind per catalogue entry, the per-user and per-channel preferences
// with the catalogue's defaults, and the Router that sends a published Event to each
// recipient's channels. It is a leaf package: the channels (Discord in notify, push in
// push) implement its Channel interface, and the producers call Publish.
package events

import "slices"

// Kind names one catalogue entry.
type Kind string

// The catalogue, in its own order.
const (
	CharacterLinked Kind = "character.linked"
	GuildApproved   Kind = "guild.approved"
	BandEntered     Kind = "band.entered"
	BisChanged      Kind = "bis.changed"
	RaidPosted      Kind = "raid.posted"
	RaidSignupNudge Kind = "raid.signup_nudge"
	RaidChanged     Kind = "raid.changed"
	RaidRosterSet   Kind = "raid.roster_set"
	RaidRosterMoved Kind = "raid.roster_moved"
	RaidStarting    Kind = "raid.starting"
	LootAwarded     Kind = "loot.awarded"
	ReportReady     Kind = "report.ready"
	RaidSummary     Kind = "raid.summary"
	CoachReady      Kind = "coach.ready"
	RankingBest     Kind = "ranking.best"
	SimFinished     Kind = "sim.finished"
	RosterPending   Kind = "roster.pending"
	RaidShort       Kind = "raid.short"
	RaidLateChange  Kind = "raid.late_change"
	PremiumSeat     Kind = "premium.seat"
	BillingFailed   Kind = "billing.failed"
	GameMilestone   Kind = "game.milestone"
)

// Channel is one way a notification reaches a player.
type Channel string

const (
	// ChannelPush is the iOS and Android apps' push notification.
	ChannelPush Channel = "push"
	// ChannelDiscord is Discord: the account's own webhook today, a DM later.
	ChannelDiscord Channel = "discord"
)

// Channels is every channel, in the order the preferences list them.
var Channels = []Channel{ChannelPush, ChannelDiscord}

// Who an event goes to, as the catalogue words it.
const (
	WhoPlayer   = "player"
	WhoMembers  = "guild_members"
	WhoSigned   = "signed_up"
	WhoInLate   = "in_and_late"
	WhoRaiders  = "raiders"
	WhoOfficers = "officers"
	WhoPayer    = "payer"
	WhoEveryone = "everyone"
)

// Spec is one catalogue entry's fixed rules.
type Spec struct {
	Kind Kind
	// Who is the catalogue's audience, for the preferences page.
	Who string
	// DefaultOn is the catalogue's default for every channel.
	DefaultOn bool
	// QuietExempt sends the event inside the player's quiet hours.
	QuietExempt bool
	// Officer marks an event only officers receive.
	Officer bool
	// TimeSensitive breaks through Focus on iOS (interruption-level time-sensitive) and goes
	// at high priority on Android; such an event also ignores quiet hours.
	TimeSensitive bool
	// OffWhenOn names an event that carries this one: with no override, this event defaults
	// to off on a channel where that one is on (raid.summary holds coach.ready and
	// ranking.best).
	OffWhenOn Kind
}

// Catalogue is every event and its defaults. bis.changed is off by default: it is the
// weekly digest a player opts into. loot.awarded is off by default because raid.summary
// carries the night's loot (coordinator ruling 2026-10-30, from the wow-player review).
var Catalogue = []Spec{
	{Kind: CharacterLinked, Who: WhoPlayer, DefaultOn: true},
	{Kind: GuildApproved, Who: WhoPlayer, DefaultOn: true},
	{Kind: BandEntered, Who: WhoPlayer, DefaultOn: true},
	{Kind: BisChanged, Who: WhoPlayer, DefaultOn: false},
	{Kind: RaidPosted, Who: WhoMembers, DefaultOn: true},
	{Kind: RaidSignupNudge, Who: WhoMembers, DefaultOn: true},
	{Kind: RaidChanged, Who: WhoSigned, DefaultOn: true},
	{Kind: RaidRosterSet, Who: WhoSigned, DefaultOn: true},
	{Kind: RaidRosterMoved, Who: WhoPlayer, DefaultOn: true, QuietExempt: true, TimeSensitive: true},
	{Kind: RaidStarting, Who: WhoInLate, DefaultOn: true, QuietExempt: true, TimeSensitive: true},
	{Kind: LootAwarded, Who: WhoPlayer, DefaultOn: false},
	{Kind: ReportReady, Who: WhoRaiders, DefaultOn: true},
	{Kind: RaidSummary, Who: WhoRaiders, DefaultOn: true},
	{Kind: CoachReady, Who: WhoPlayer, DefaultOn: true, OffWhenOn: RaidSummary},
	{Kind: RankingBest, Who: WhoPlayer, DefaultOn: true, OffWhenOn: RaidSummary},
	{Kind: SimFinished, Who: WhoPlayer, DefaultOn: true},
	{Kind: RosterPending, Who: WhoOfficers, DefaultOn: true, Officer: true},
	{Kind: RaidShort, Who: WhoOfficers, DefaultOn: true, Officer: true},
	{Kind: RaidLateChange, Who: WhoOfficers, DefaultOn: true, Officer: true},
	{Kind: PremiumSeat, Who: WhoPlayer, DefaultOn: true},
	{Kind: BillingFailed, Who: WhoPayer, DefaultOn: true},
	{Kind: GameMilestone, Who: WhoEveryone, DefaultOn: true},
}

// Lookup is kind's catalogue entry; ok is false for a kind the catalogue does not name.
func Lookup(kind Kind) (Spec, bool) {
	i := slices.IndexFunc(Catalogue, func(s Spec) bool { return s.Kind == kind })
	if i < 0 {
		return Spec{}, false
	}
	return Catalogue[i], true
}

// KnownChannel reports whether c is one of Channels.
func KnownChannel(c Channel) bool { return slices.Contains(Channels, c) }
