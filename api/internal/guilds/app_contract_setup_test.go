package guilds

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/contract"
)

// The Guild tab's officer Setup row (design/specs/2026-11-10-guild-tab.md 2.3) reads the officer's
// checklist; the fixture is the floor the apps decode.
func TestAppContractGuildSetup(t *testing.T) {
	w := newRaidWorld(t, true)
	w.as(w.officer)

	res := w.h.do(http.MethodGet, w.path("/setup"), "")
	contract.AssertFixture(t, "guild.setup", guildBody(t, res, http.StatusOK), contract.Options{})
}

// The Guild tab's Loot section (spec 2026-11-10-guild-tab.md 2.2) reads the last night's drops and
// who won them; a member's copy carries no vote counts.
func TestAppContractGuildLootNights(t *testing.T) {
	w := newRaidWorld(t, true)
	w.as(w.officer)
	if res := w.putVote("", memberKeyA, `"yes"`); res.StatusCode != http.StatusOK {
		t.Fatalf("vote = %d: %s", res.StatusCode, readBody(res))
	}
	if res := w.h.do(http.MethodPost, w.path("/loot/awards"), fmt.Sprintf(
		`{"encounter_id":1084,"item_id":%d,"character_key":%q,"ms_os":"ms"}`, councilItem, memberKeyA)); res.StatusCode != http.StatusOK {
		t.Fatalf("award = %d: %s", res.StatusCode, readBody(res))
	}
	w.as(w.member)
	res := w.h.do(http.MethodGet, w.path("/loot/nights")+"?limit=1", "")
	contract.AssertFixture(t, "guild.loot.nights", guildBody(t, res, http.StatusOK), contract.Options{})
}
