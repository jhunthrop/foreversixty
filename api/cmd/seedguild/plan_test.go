package main

import (
	"testing"
	"time"
)

func testGuild() guildInfo {
	return guildInfo{ID: 2, Region: "us", Ruleset: "pvp", Name: "OLYMPUS XXVII", OfficerMaxRankIndex: 1}
}

func TestBuildSeedPlanCharacterKeysAreUniqueAndWellFormed(t *testing.T) {
	plan := buildSeedPlan(testGuild(), time.Now())
	seen := map[string]bool{}
	for _, pc := range plan.Characters {
		if seen[pc.Key] {
			t.Fatalf("duplicate character key %q", pc.Key)
		}
		seen[pc.Key] = true
		want := "us/pvp/" + lowerHyphen(pc.Mock.Name)
		if pc.Key != want {
			t.Errorf("key = %q, want %q", pc.Key, want)
		}
	}
}

func TestBuildSeedPlanBattletagsAreUniquePerAccount(t *testing.T) {
	plan := buildSeedPlan(testGuild(), time.Now())
	battletagOf := map[string]string{} // account key -> battletag
	seenBattletags := map[string]bool{}
	for _, pc := range plan.Characters {
		if len(pc.Battletag) < 5 || pc.Battletag[:5] != "Seed#" {
			t.Errorf("battletag %q does not carry the Seed# prefix", pc.Battletag)
		}
		if existing, ok := battletagOf[pc.AccountKey]; ok {
			if existing != pc.Battletag {
				t.Fatalf("account %q has two battletags: %q and %q", pc.AccountKey, existing, pc.Battletag)
			}
			continue // a second character of an already-seen account: shares its battletag on purpose
		}
		battletagOf[pc.AccountKey] = pc.Battletag
		if seenBattletags[pc.Battletag] {
			t.Fatalf("battletag %q reused by a different account", pc.Battletag)
		}
		seenBattletags[pc.Battletag] = true
	}
}

func TestBuildSeedPlanHasTwentyTwoAccountsForTwentyFourCharacters(t *testing.T) {
	plan := buildSeedPlan(testGuild(), time.Now())
	if len(plan.Characters) != 24 {
		t.Fatalf("len(Characters) = %d, want 24", len(plan.Characters))
	}
	if got := plan.accountCount(); got != 22 {
		t.Fatalf("accountCount() = %d, want 22 (two accounts each carry a main and an alt)", got)
	}
}

func TestBuildSeedPlanRankAndVerificationMatchRoster(t *testing.T) {
	plan := buildSeedPlan(testGuild(), time.Now())
	for _, pc := range plan.Characters {
		switch {
		case pc.Mock.Officer:
			if pc.Rank != "officer" || pc.RankIndex != rankIndexOfficer {
				t.Errorf("%s: officer rank = %s/%d", pc.Mock.Name, pc.Rank, pc.RankIndex)
			}
		default:
			if pc.Rank != "member" || pc.RankIndex != rankIndexMember {
				t.Errorf("%s: member rank = %s/%d", pc.Mock.Name, pc.Rank, pc.RankIndex)
			}
		}
		if pc.Verified == pc.Mock.Unverifed {
			t.Errorf("%s: verified=%v unverified-mock=%v should disagree", pc.Mock.Name, pc.Verified, pc.Mock.Unverifed)
		}
	}
}

func TestBuildSeedPlanReportsScheduledInRange(t *testing.T) {
	now := time.Date(2026, 10, 4, 12, 0, 0, 0, time.UTC)
	plan := buildSeedPlan(testGuild(), now)
	if len(plan.Reports) != 4 {
		t.Fatalf("len(Reports) = %d, want 4", len(plan.Reports))
	}
	withinWeek := 0
	for _, r := range plan.Reports {
		age := now.Sub(r.CreatedAt)
		if age < 0 || age > 14*24*time.Hour {
			t.Errorf("report %s created_at %v outside the last 14 days", r.ID, r.CreatedAt)
		}
		if age <= 7*24*time.Hour {
			withinWeek++
		}
		if !r.CompletedAt.After(r.CreatedAt) {
			t.Errorf("report %s completed_at not after created_at", r.ID)
		}
	}
	if withinWeek != 2 {
		t.Fatalf("reports within the last 7 days = %d, want 2", withinWeek)
	}
}

func TestSeedTagIsStablePerGuild(t *testing.T) {
	if seedTag(2) != "seedguild-2" {
		t.Fatalf("seedTag(2) = %q, want seedguild-2", seedTag(2))
	}
	if seedTag(2) == seedTag(3) {
		t.Fatal("seedTag must differ per guild")
	}
}

// lowerHyphen mirrors character.Slug's own transform, kept local so this test does not
// need to import the character package just to re-derive the expected key.
func lowerHyphen(s string) string {
	out := make([]rune, 0, len(s))
	for _, r := range s {
		if r == ' ' {
			out = append(out, '-')
			continue
		}
		if r >= 'A' && r <= 'Z' {
			r = r - 'A' + 'a'
		}
		out = append(out, r)
	}
	return string(out)
}
