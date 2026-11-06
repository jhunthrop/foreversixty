package guilds

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/contract"
)

// profileOptional are the roster row's omitempty fields, absent on a seeded row.
var profileOptional = []string{
	"data.member.class", "data.member.spec", "data.member.role", "data.member.spec_name",
	"data.member.spec_icon", "data.member.item_level", "data.member.attendance_30",
	"data.member.best_parse", "data.member.rating", "data.member.professions",
}

func TestRecruitResponsesKeepTheirFixtures(t *testing.T) {
	h := newHTTPHarness(t)
	h.svc.ApplicationBot = &fakeAppBot{grant: OutcomeNotLinked, dm: OutcomeSent}
	gid, officer := recruitGuild(t, h, "Contract")
	h.as(officer)
	res := h.do(http.MethodPatch, fmt.Sprintf("/v1/guilds/%d/recruit", gid),
		`{"open":true,"wanted":[{"class":"priest","spec":"priest-holy"}],"discord_role_id":"123456789012345678"}`)
	contract.AssertFixture(t, "guild.recruit.settings", guildBody(t, res, http.StatusOK), contract.Options{})
	h.as(0)
	res = h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/recruit", gid), "")
	contract.AssertFixture(t, "guild.recruit", guildBody(t, res, http.StatusOK), contract.Options{})

	uid, key := applicant(t, h, "contract-app@example.com", "contractchar")
	h.as(uid)
	res = h.do(http.MethodPost, fmt.Sprintf("/v1/guilds/%d/applications", gid), applyBody(key, true))
	raw := guildBody(t, res, http.StatusCreated)
	contract.AssertFixture(t, "guild.application", raw, contract.Options{})
	res = h.do(http.MethodPost, fmt.Sprintf("/v1/guilds/%d/applications", gid), applyBody(key, true))
	contract.AssertFixture(t, "guild.application.error-already-applied", guildBody(t, res, http.StatusConflict), contract.Options{})
	res = h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/applications/mine", gid), "")
	contract.AssertFixture(t, "guild.applications.mine", guildBody(t, res, http.StatusOK), contract.Options{})

	apps, err := h.store.MyApplications(context.Background(), gid, uid)
	if err != nil || len(apps) != 1 {
		t.Fatalf("applications = %v, %v", apps, err)
	}
	app := fmt.Sprintf("/v1/guilds/%d/applications/%d", gid, apps[0].ID)
	h.as(officer)
	res = h.do(http.MethodPut, app+"/vote", `{"vote":"yes"}`)
	contract.AssertFixture(t, "guild.application.votes", guildBody(t, res, http.StatusOK), contract.Options{})
	res = h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/applications", gid), "")
	contract.AssertFixture(t, "guild.applications", guildBody(t, res, http.StatusOK), contract.Options{})
	res = h.do(http.MethodPost, app+"/accept", "")
	contract.AssertFixture(t, "guild.application.accept", guildBody(t, res, http.StatusOK), contract.Options{})
}

func TestMemberResponsesKeepTheirFixtures(t *testing.T) {
	h := newHTTPHarness(t)
	gid, leader, _, member := notesGuild(t, h)
	h.as(leader)
	notes := fmt.Sprintf("/v1/guilds/%d/members/%s/notes", gid, accountKeyFor(member))
	guildBody(t, h.do(http.MethodPost, notes, `{"body":"Always brings flasks."}`), http.StatusCreated)
	res := h.do(http.MethodGet, notes, "")
	contract.AssertFixture(t, "guild.member.notes", guildBody(t, res, http.StatusOK), contract.Options{})
	res = h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/members/%s/profile", gid, accountKeyFor(member)), "")
	contract.AssertFixture(t, "guild.member.profile", guildBody(t, res, http.StatusOK),
		contract.Options{Optional: profileOptional})
}
