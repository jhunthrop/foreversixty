package guilds

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/refusal"
)

const (
	guildTierSQL = `insert into entitlements (guild_id, tier, plan, source, status, current_period_end)
		values ($1, 'guild', 'guild', 'grant', 'active', now() + interval '30 days')`
	raidGuildTierSQL = `insert into entitlements (guild_id, tier, plan, source, status, current_period_end)
		values ($1, 'raid_guild', 'guild', 'grant', 'active', now() + interval '30 days')`
)

// awardFixture is an officer of a guild with the given plan row ("" for none) and a target
// character to award to.
func awardFixture(t *testing.T, name, planSQL string) (*httpHarness, int64, string) {
	t.Helper()
	h := newHTTPHarness(t)
	h.svc.Entitlements = &entitlements.Store{Pool: h.pool}
	gid := seedGuild(t, h.pool, name)
	target := fmt.Sprintf("us/hardcore/gatetarget%d", gid)
	seedHunterCharacter(t, h, gid, fmt.Sprintf("gate-target-%d@example.com", gid), target, 21524)
	officer := seedUser(t, h.pool, fmt.Sprintf("gate-officer-%d@example.com", gid))
	seedVerifiedMember(t, h, gid, officer, fmt.Sprintf("gateofficer%d", gid), "officer")
	if planSQL != "" {
		if _, err := h.pool.Exec(t.Context(), planSQL, gid); err != nil {
			t.Fatal(err)
		}
	}
	h.actor = auth.Actor{UserID: officer, Role: "user", Method: "session"}
	return h, gid, target
}

func awardBody(target string) string {
	return fmt.Sprintf(`{"encounter_id":1084,"item_id":16963,"character_key":%q}`, target)
}

func TestLootAward_GuildTierRefused(t *testing.T) {
	h, gid, target := awardFixture(t, "GuildTierGate", guildTierSQL)
	res := h.do(http.MethodPost, fmt.Sprintf("/v1/guilds/%d/loot/awards", gid), awardBody(target))
	assertRefusal(t, res, "raid_guild_required", refusal.AwardPlanRequired)

	res = h.do(http.MethodDelete, fmt.Sprintf("/v1/guilds/%d/loot/awards/1", gid), "")
	assertRefusal(t, res, "raid_guild_required", refusal.AwardPlanRequired)
}

func TestLootAward_NoPlanKeepsGuildPlanRequired(t *testing.T) {
	h, gid, target := awardFixture(t, "NoPlanGate", "")
	res := h.do(http.MethodPost, fmt.Sprintf("/v1/guilds/%d/loot/awards", gid), awardBody(target))
	assertRefusal(t, res, "guild_plan_required", refusal.AwardPlanRequired)
}

func TestLootGate_NoPlanAndLapsedPlan(t *testing.T) {
	lapsed := `insert into entitlements (guild_id, tier, plan, source, status, current_period_end)
		values ($1, 'raid_guild', 'guild', 'grant', 'active', now() - interval '1 day')`
	for name, planSQL := range map[string]string{"no plan": "", "lapsed plan": lapsed} {
		t.Run(name, func(t *testing.T) {
			h, gid, target := awardFixture(t, "Gate "+name, planSQL)
			res := h.do(http.MethodPost, fmt.Sprintf("/v1/guilds/%d/loot/awards", gid), awardBody(target))
			assertRefusal(t, res, "guild_plan_required", refusal.AwardPlanRequired)
			res = h.do(http.MethodDelete, fmt.Sprintf("/v1/guilds/%d/loot/awards/1", gid), "")
			assertRefusal(t, res, "guild_plan_required", refusal.AwardPlanRequired)
		})
	}
}

func TestLootAward_RaidGuildAllowed(t *testing.T) {
	h, gid, target := awardFixture(t, "RaidGuildGate", raidGuildTierSQL)
	res := h.do(http.MethodPost, fmt.Sprintf("/v1/guilds/%d/loot/awards", gid), awardBody(target))
	res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("award = %d, want 200", res.StatusCode)
	}
}

func TestHome_RaidGuildToolsFlag(t *testing.T) {
	tests := []struct {
		name         string
		planSQL      string
		wantRaid     bool
		wantOfficers bool
	}{
		{"raid guild", raidGuildTierSQL, true, true},
		{"guild tier", guildTierSQL, false, true},
		{"no plan", "", false, false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			h, gid, _ := awardFixture(t, "Flag "+tc.name, tc.planSQL)
			res := h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/home", gid), "")
			var view HomeView
			h.data(res, &view)
			if view.Viewer.RaidGuildTools != tc.wantRaid || view.Viewer.OfficerTools != tc.wantOfficers {
				t.Fatalf("raid_guild_tools=%v officer_tools=%v, want %v/%v",
					view.Viewer.RaidGuildTools, view.Viewer.OfficerTools, tc.wantRaid, tc.wantOfficers)
			}
		})
	}
}
