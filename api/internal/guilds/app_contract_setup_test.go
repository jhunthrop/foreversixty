package guilds

import (
	"net/http"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/contract"
)

// The Guild tab's Setup row (design/specs/2026-11-10-guild-tab.md 2.3) reads the officer's
// checklist; the fixture is the floor the apps decode.
func TestAppContractGuildSetup(t *testing.T) {
	w := newRaidWorld(t, true)
	w.as(w.officer)

	res := w.h.do(http.MethodGet, w.path("/setup"), "")
	contract.AssertFixture(t, "guild.setup", guildBody(t, res, http.StatusOK), contract.Options{})
}
