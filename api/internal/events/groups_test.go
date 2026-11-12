package events

import (
	"slices"
	"testing"
)

func TestEveryCatalogueKindHasAGroup(t *testing.T) {
	for _, s := range Catalogue {
		g, ok := kindGroups[s.Kind]
		if !ok {
			t.Errorf("%s has no group", s.Kind)
			continue
		}
		if !KnownGroup(g) {
			t.Errorf("%s is in unknown group %q", s.Kind, g)
		}
	}
	for kind := range kindGroups {
		if _, ok := Lookup(kind); !ok {
			t.Errorf("the group table names %s, which the catalogue does not", kind)
		}
	}
}

func TestOfficerKindsAreTheOfficerGroup(t *testing.T) {
	for _, s := range Catalogue {
		if s.Officer != (GroupOf(s.Kind) == GroupOfficer) {
			t.Errorf("%s: officer %v, group %s", s.Kind, s.Officer, GroupOf(s.Kind))
		}
	}
}

func TestGroupsMatchTheAppsSettings(t *testing.T) {
	cases := map[Group][]Kind{
		GroupRaids:   {RaidPosted, RaidSignupNudge, RaidChanged, RaidRosterSet, RaidRosterMoved, RaidStarting},
		GroupGear:    {BandEntered, BisChanged, LootAwarded, SimFinished, GameMilestone},
		GroupLogs:    {ReportReady, RaidSummary, CoachReady, RankingBest},
		GroupLadder:  {GuildSimPassed},
		GroupTiers:   {TierMoved},
		GroupOfficer: {RosterPending, RaidShort, RaidLateChange, GuildApplicant},
		GroupAccount: {CharacterLinked, GuildApproved, PremiumSeat, BillingFailed, BillingTrialEnding},
	}
	for g, want := range cases {
		if got := KindsOf(g); !slices.Equal(got, want) {
			t.Errorf("%s = %v, want %v", g, got, want)
		}
	}
	if len(cases) != len(Groups) {
		t.Fatalf("groups %v", Groups)
	}
}

func TestAnUnknownKindReadsAsAccount(t *testing.T) {
	if GroupOf("gone.kind") != GroupAccount || KnownGroup("nope") {
		t.Fatal("unknown kind or group")
	}
}
