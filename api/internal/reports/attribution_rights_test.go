package reports

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
)

// A member's attributed report hands its guild's officers nothing: guild_id stays null, so
// mayEdit is the owner's alone.
func TestAttributionGivesTheGuildsOfficersNoEditRight(t *testing.T) {
	h := newHarness(t)
	g := h.seedGuild("Alpha", "hardcore")
	h.seedGuildCharacter(g, h.owner, harnessLoggingKey, true)
	const officer = int64(990001)
	h.addGuildMember(g, officer, "officer")
	id := h.createReport(Public)
	if _, err := h.store.AttributeGuild(t.Context(), id); err != nil {
		t.Fatal(err)
	}
	rep, err := h.store.Get(t.Context(), id)
	if err != nil {
		t.Fatal(err)
	}
	ok, err := h.service.mayEdit(t.Context(), auth.Actor{UserID: officer, Role: "user", Method: "session"}, rep)
	if err != nil {
		t.Fatal(err)
	}
	if ok {
		t.Fatal("an officer may edit a member's attributed report")
	}
}

// The officer PATCH still sets guild_id itself.
func TestOfficerPatchStillSetsGuildID(t *testing.T) {
	h := newHarness(t)
	g := h.seedGuild("Alpha", "hardcore")
	h.addGuildMember(g, h.owner, "officer")
	id := h.createReport(Public)
	res := h.json(http.MethodPatch, "/v1/reports/"+id, `{"guild_id":`+strconv.FormatInt(g, 10)+`}`)
	res.Body.Close()
	if res.StatusCode != http.StatusOK || !sameGuild(h.reportGuild(id), g) {
		t.Fatalf("patch = %d, guild_id = %v, want 200 and %d", res.StatusCode, h.reportGuild(id), g)
	}
}
