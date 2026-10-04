// api/cmd/seedguild/plan.go
//
// seedPlan is the whole seed, computed once as pure data from the roster and raid plans
// above plus the target guild's own region/ruleset/id. Both the dry-run printer and the
// real database writer build and walk the exact same plan, so what a dry run shows is
// what an apply writes, byte for byte.
package main

import (
	"fmt"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/character"
)

// region is the region every mock character and report is seeded under. The guild this
// tool seeds is itself region "us" (the owner's own guild); a mock character's realm-less
// key format (api/internal/character) only has room for one region/ruleset pair per
// character, so this tool seeds entirely inside the guild's own region/ruleset rather
// than scattering mock characters across others the guild could never actually draw from.
type guildInfo struct {
	ID                  int64
	Region, Ruleset     string
	Name                string
	ClaimedBy           *int64
	ClaimedAt           *time.Time
	OfficerMaxRankIndex int
}

// seedOfficerMaxRankIndex is the officer-rank threshold this seed sets on the guild
// (spec: "officer_max_rank_index 1") — rank_index 1 is "officer", 0 is always the guild
// master regardless of this setting (guilds.DeriveRank).
const seedOfficerMaxRankIndex = 1

const (
	rankIndexLeader  = 0
	rankIndexOfficer = 1
	rankIndexMember  = 2
)

// plannedCharacter is one mock roster entry with every value this tool will actually
// write already computed: its account's battletag, its character key, its membership
// rank, and when its export was last "synced".
type plannedCharacter struct {
	Mock mockCharacter
	// AccountKey groups two characters under one account (mockCharacter.accountKey) -
	// Battletag is shared by every plannedCharacter with the same AccountKey.
	AccountKey string
	Battletag  string
	Key        string
	Rank       string
	RankIndex  int
	Verified   bool
	ExportedAt time.Time
}

// plannedReport is one raid night with its id and every fight's own scheduled start,
// alongside the reportPlan it was built from.
type plannedReport struct {
	Plan        reportPlan
	ID          string
	Title       string
	CreatedAt   time.Time
	CompletedAt time.Time
	FightStarts []time.Time
}

// seedPlan is the whole seed.
type seedPlan struct {
	Tag        string
	Guild      guildInfo
	Characters []plannedCharacter
	Reports    []plannedReport
}

// seedTag is the one string that names every row this tool's --remove must find again -
// tied to the guild id, so seeding (and un-seeding) one guild never touches another.
func seedTag(guildID int64) string {
	return fmt.Sprintf("seedguild-%d", guildID)
}

// battletagFor is the account tag every mock user gets - unmistakable at a glance in any
// admin view, and never a value a real Battle.net account could hold (real battletags
// never contain the literal word "Seed").
func battletagFor(i int) string {
	return fmt.Sprintf("Seed#%04d", i+1)
}

// exportedAtFor spreads addon_exports.updated_at across the last ten days (so
// HomeRoster's "logged recently" 24-hour check reads true for some rows and false for
// others), deterministically per roster index.
func exportedAtFor(now time.Time, i int) time.Time {
	daysBack := time.Duration((i*37)%10) * 24 * time.Hour
	hoursBack := time.Duration((i*13)%24) * time.Hour
	return now.Add(-daysBack - hoursBack)
}

// buildSeedPlan computes the whole plan for guild from the fixed roster and raid plans.
func buildSeedPlan(guild guildInfo, now time.Time) seedPlan {
	roster := mockRoster()
	plan := seedPlan{Tag: seedTag(guild.ID), Guild: guild}

	// Two characters share an account (roster.go's Account field) when they share an
	// accountKey; the account gets one battletag, assigned the first time its key is
	// seen, in roster order - so a dry run and a real apply agree on which battletag
	// goes with which account regardless of how many characters share it.
	accountBattletags := map[string]string{}
	nextAccount := 0
	for i, m := range roster {
		rank, rankIndex := "member", rankIndexMember
		if m.Officer {
			rank, rankIndex = "officer", rankIndexOfficer
		}
		key := m.accountKey()
		battletag, ok := accountBattletags[key]
		if !ok {
			battletag = battletagFor(nextAccount)
			accountBattletags[key] = battletag
			nextAccount++
		}
		plan.Characters = append(plan.Characters, plannedCharacter{
			Mock: m, AccountKey: key, Battletag: battletag,
			Key:  character.Key(guild.Region, guild.Ruleset, m.Name),
			Rank: rank, RankIndex: rankIndex,
			Verified:   !m.Unverifed,
			ExportedAt: exportedAtFor(now, i),
		})
	}

	for _, rp := range raidPlans() {
		createdAt, completedAt, starts := rp.schedule(now)
		plan.Reports = append(plan.Reports, plannedReport{
			Plan: rp, ID: fmt.Sprintf("%s-report-%s", plan.Tag, rp.Tag), Title: rp.titleFor(createdAt),
			CreatedAt: createdAt, CompletedAt: completedAt, FightStarts: starts,
		})
	}
	return plan
}

// mockRosterFrom projects a plan's characters back to the []mockCharacter shape
// buildFightMetrics and talentSplitFor want, in the same order chars is given in - the
// plan always carries plannedCharacter in mockRoster()'s own order, so this is a cheap
// projection rather than a second data source.
func mockRosterFrom(chars []plannedCharacter) []mockCharacter {
	out := make([]mockCharacter, len(chars))
	for i, c := range chars {
		out[i] = c.Mock
	}
	return out
}

// officerCount, unverifiedCount and memberCount are the dry-run summary's own tallies,
// kept here so the summary and the writer read the same numbers.
func (p seedPlan) officerCount() int {
	return countBy(mockRoster(), func(c mockCharacter) bool { return c.Officer })
}
func (p seedPlan) unverifiedCount() int {
	return countBy(mockRoster(), func(c mockCharacter) bool { return c.Unverifed })
}

// accountCount is how many distinct accounts the roster resolves to - 22 today, two
// pairs of mockRoster entries sharing an Account.
func (p seedPlan) accountCount() int {
	seen := map[string]bool{}
	for _, pc := range p.Characters {
		seen[pc.AccountKey] = true
	}
	return len(seen)
}

func (p seedPlan) fightCount() int {
	n := 0
	for _, r := range p.Reports {
		n += len(r.Plan.Fights)
	}
	return n
}
