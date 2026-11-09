package events

import "slices"

// Group is one filter of the app's Activity tab, matching its notification settings.
type Group string

// The groups. The app shows the first six as filters; account rows appear under All only.
const (
	GroupRaids   Group = "raids"
	GroupGear    Group = "gear"
	GroupLogs    Group = "logs"
	GroupLadder  Group = "ladder"
	GroupTiers   Group = "tiers"
	GroupOfficer Group = "officer"
	GroupAccount Group = "account"
)

// Groups is every group, in the order the app lists them.
var Groups = []Group{GroupRaids, GroupGear, GroupLogs, GroupLadder, GroupTiers, GroupOfficer, GroupAccount}

// kindGroups is the one table of which group each catalogue kind belongs to.
var kindGroups = map[Kind]Group{
	RaidPosted:         GroupRaids,
	RaidSignupNudge:    GroupRaids,
	RaidChanged:        GroupRaids,
	RaidRosterSet:      GroupRaids,
	RaidRosterMoved:    GroupRaids,
	RaidStarting:       GroupRaids,
	LootAwarded:        GroupGear,
	BisChanged:         GroupGear,
	BandEntered:        GroupGear,
	GameMilestone:      GroupGear,
	SimFinished:        GroupGear,
	ReportReady:        GroupLogs,
	CoachReady:         GroupLogs,
	RankingBest:        GroupLogs,
	RaidSummary:        GroupLogs,
	RaidShort:          GroupOfficer,
	RaidLateChange:     GroupOfficer,
	RosterPending:      GroupOfficer,
	GuildApplicant:     GroupOfficer,
	CharacterLinked:    GroupAccount,
	GuildApproved:      GroupAccount,
	PremiumSeat:        GroupAccount,
	BillingFailed:      GroupAccount,
	BillingTrialEnding: GroupAccount,
	GuildSimPassed:     GroupLadder,
	TierMoved:          GroupTiers,
}

// GroupOf is kind's group. A kind the catalogue no longer names reads as account, so an old
// row still shows under All.
func GroupOf(kind Kind) Group {
	if g, ok := kindGroups[kind]; ok {
		return g
	}
	return GroupAccount
}

// KnownGroup reports whether g is one of Groups.
func KnownGroup(g Group) bool { return slices.Contains(Groups, g) }

// KindsOf is every catalogue kind in group g, in catalogue order.
func KindsOf(g Group) []Kind {
	var out []Kind
	for _, s := range Catalogue {
		if kindGroups[s.Kind] == g {
			out = append(out, s.Kind)
		}
	}
	return out
}
