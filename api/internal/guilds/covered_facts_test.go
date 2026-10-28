// api/internal/guilds/covered_facts_test.go
package guilds

import (
	"context"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
)

const coveredPlanSQL = `insert into entitlements
	(guild_id, tier, plan, source, status, current_period_end, trial_end, cancel_at_period_end, billing_user_id)
	values ($1, $2, 'guild', $3, $4, $5, $6, $7, $8)`

// coveredOf is guildID's covered view for viewer (nil when the guild answers none).
func coveredOf(t *testing.T, h *httpHarness, viewer, guildID int64) *auth.GuildCoveredView {
	t.Helper()
	got, err := h.store.GuildStanding(context.Background(), viewer, []int64{guildID})
	if err != nil {
		t.Fatal(err)
	}
	return got[guildID].Covered
}

func seedStripePlan(t *testing.T, h *httpHarness, gid int64, tier, status string, end *time.Time, cancel bool, payer *int64) {
	t.Helper()
	var trialEnd *time.Time
	if status == "trialing" {
		trialEnd = end
	}
	setGuild(t, h, coveredPlanSQL, gid, tier, "stripe", status, end, trialEnd, cancel, payer)
}

func daysFromNow(d int) *time.Time {
	at := time.Now().Add(time.Duration(d) * 24 * time.Hour)
	return &at
}

func TestGuildStanding_CoveredStates(t *testing.T) {
	cases := []struct {
		name   string
		status string
		end    *time.Time
		cancel bool
		want   string // "" = null
	}{
		{"trial", "trialing", daysFromNow(10), false, "trial"},
		{"active", "active", daysFromNow(20), false, "active"},
		{"ending", "active", daysFromNow(5), true, "ending"},
		{"past due", "past_due", daysFromNow(1), false, "past_due"},
		{"lapsed day 89", "canceled", daysFromNow(-89), false, "lapsed"},
		{"null day 91", "canceled", daysFromNow(-91), false, ""},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			h := withStepSources(newHTTPHarness(t))
			gid := seedGuild(t, h.pool, "Covered")
			viewer := seedRanked(t, h, gid, "Viewer", "member", true)
			seedStripePlan(t, h, gid, "guild", c.status, c.end, c.cancel, nil)
			got := coveredOf(t, h, viewer, gid)
			if c.want == "" {
				if got != nil {
					t.Fatalf("covered = %+v, want null", got)
				}
				return
			}
			if got == nil || got.State != c.want || got.Tier != "guild" || got.EndsAt == nil {
				t.Fatalf("covered = %+v, want state %s tier guild with ends_at", got, c.want)
			}
		})
	}
}

func TestGuildStanding_CoveredNullWithoutPlan(t *testing.T) {
	h := withStepSources(newHTTPHarness(t))
	gid := seedGuild(t, h.pool, "Bare")
	viewer := seedRanked(t, h, gid, "Unverified", "member", false)
	if got := coveredOf(t, h, viewer, gid); got != nil {
		t.Fatalf("covered = %+v, want null with no plan", got)
	}
}

func TestGuildStanding_BilledByHiddenFromUnverified(t *testing.T) {
	h := withStepSources(newHTTPHarness(t))
	gid := seedGuild(t, h.pool, "Billed")
	payer := seedRanked(t, h, gid, "Payer", "leader", true)
	unverified := seedRanked(t, h, gid, "Newcomer", "member", false)
	verified := seedRanked(t, h, gid, "Regular", "member", true)
	seedStripePlan(t, h, gid, "guild", "active", daysFromNow(20), false, &payer)

	if got := coveredOf(t, h, unverified, gid); got == nil || got.BilledBy != "" || got.YouPay {
		t.Fatalf("unverified member sees %+v, want covered with billed_by empty", got)
	}
	if got := coveredOf(t, h, verified, gid); got == nil || got.BilledBy != "Payer" || got.YouPay {
		t.Fatalf("verified member sees %+v, want billed_by Payer", got)
	}
	if got := coveredOf(t, h, payer, gid); got == nil || !got.YouPay {
		t.Fatalf("payer sees %+v, want you_pay", got)
	}
}

func TestGuildStanding_BilledByEmptyForGrant(t *testing.T) {
	h := withStepSources(newHTTPHarness(t))
	gid := seedGuild(t, h.pool, "Granted")
	viewer := seedRanked(t, h, gid, "Viewer", "member", true)
	setGuild(t, h, guildTierSQL, gid)
	got := coveredOf(t, h, viewer, gid)
	if got == nil || got.BilledBy != "" || got.State != "active" {
		t.Fatalf("grant covered = %+v, want active with billed_by empty", got)
	}
}

func TestGuildStanding_SeatCounts(t *testing.T) {
	h := withStepSources(newHTTPHarness(t))
	gid := seedGuild(t, h.pool, "Seated")
	holder := seedRanked(t, h, gid, "Holder", "member", true)
	other := seedRanked(t, h, gid, "Other", "member", true)
	seedStripePlan(t, h, gid, "raid_guild", "active", daysFromNow(20), false, nil)
	for slot, uid := range map[int]int64{1: holder, 2: other} {
		setGuild(t, h, `insert into premium_seats (guild_id, slot, state, holder_user_id, given_at, movable_at)
			values ($1, $2, 'held', $3, now(), now())`, gid, slot, uid)
	}
	got := coveredOf(t, h, holder, gid)
	if got == nil || got.SeatsGiven != 2 || got.SeatsTotal != 3 || !got.YouHoldSeat {
		t.Fatalf("raid guild covered = %+v, want 2 of 3, you_hold_seat", got)
	}

	small := seedGuild(t, h.pool, "Small")
	member := seedRanked(t, h, small, "Solo", "member", true)
	seedStripePlan(t, h, small, "guild", "active", daysFromNow(20), false, nil)
	second := seedRanked(t, h, small, "Second", "member", true)
	setGuild(t, h, `insert into premium_seats (guild_id, slot, state, holder_user_id, given_at, movable_at)
		values ($1, 1, 'held', $2, now(), now()), ($1, 2, 'held', $3, now(), now())`, small, member, second)
	got = coveredOf(t, h, member, small)
	if got == nil || got.SeatsTotal != 1 || got.SeatsGiven != 1 {
		t.Fatalf("guild tier with a stale slot 2 = %+v, want 1 of 1", got)
	}
}
