package guilds

import (
	"net/http"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/contract"
)

// guildBody reads a response, failing the test on any other status than want.
func guildBody(t *testing.T, res *http.Response, want int) []byte {
	t.Helper()
	raw := []byte(readBody(res))
	if res.StatusCode != want {
		t.Fatalf("%s = %d, want %d: %s", res.Request.URL.Path, res.StatusCode, want, raw)
	}
	return raw
}

func assertGuildCache(t *testing.T, res *http.Response, want string) {
	t.Helper()
	if got := res.Header.Get("Cache-Control"); got != want {
		t.Errorf("Cache-Control = %q, want %q", got, want)
	}
}

const privateCache = "private, no-cache"

// raidViewOptions are the options for a raid view found at prefix ("data" or "data.raid").
func raidViewOptions(prefix string) contract.Options {
	return contract.Options{Optional: append([]string{
		prefix + ".viewer.characters[].raid_spec",
	}, raidViewPending(prefix)...)}
}

func TestAppContractGuildHome(t *testing.T) {
	w, _, _ := reportNight(t)
	seedExport(t, w.h.pool, w.member, memberKeyC, "us", "hardcore", "Unverifiedalt")
	w.createRaid(t)
	w.as(w.officer)

	res := w.h.do(http.MethodGet, w.path("/home"), "")
	assertGuildCache(t, res, privateCache)
	contract.AssertFixture(t, "guild.home", guildBody(t, res, http.StatusOK), contract.Options{Optional: homeOptional})
}

func TestAppContractGuildHomeForbidden(t *testing.T) {
	w := newRaidWorld(t, true)
	w.as(seedUser(t, w.h.pool, "stranger@example.com"))

	contract.AssertFixture(t, "guild.home.error-forbidden",
		guildBody(t, w.h.do(http.MethodGet, w.path("/home"), ""), http.StatusForbidden), contract.Options{})
	w.as(0)
	contract.AssertFixture(t, "error-unauthorized",
		guildBody(t, w.h.do(http.MethodGet, w.path("/home"), ""), http.StatusUnauthorized), contract.Options{})
}

func TestAppContractRaidsList(t *testing.T) {
	w := newRaidWorld(t, true)
	w.createRaid(t)
	w.as(w.member)

	res := w.h.do(http.MethodGet, w.path("/raids"), "")
	assertGuildCache(t, res, privateCache)
	contract.AssertFixture(t, "guild.raids", guildBody(t, res, http.StatusOK), contract.Options{
		Optional: append([]string{"data.next_cursor", "data.guild_timezone"}, raidListOptional...),
	})
}

func TestAppContractRaidsListForANonMember(t *testing.T) {
	w := newRaidWorld(t, true)
	w.createRaid(t)
	w.as(0)

	contract.AssertFixture(t, "guild.raids.public",
		guildBody(t, w.h.do(http.MethodGet, w.path("/raids"), ""), http.StatusOK), contract.Options{
			Optional: []string{"data.next_cursor"},
		})
}

func TestAppContractRaidDetail(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	w.as(w.member)
	w.signup(t, raid.ID, memberKeyA, `{"state":"in","role":"tank","note":"bringing flasks"}`)

	res := w.h.do(http.MethodGet, w.path("/raids/%d", raid.ID), "")
	assertGuildCache(t, res, privateCache)
	contract.AssertFixture(t, "guild.raid", guildBody(t, res, http.StatusOK), raidViewOptions("data"))

	contract.AssertFixture(t, "guild.raid.error-not-found",
		guildBody(t, w.h.do(http.MethodGet, w.path("/raids/999999"), ""), http.StatusNotFound), contract.Options{})
	w.as(seedUser(t, w.h.pool, "outsider@example.com"))
	contract.AssertFixture(t, "guild.raid.error-not-found",
		guildBody(t, w.h.do(http.MethodGet, w.path("/raids/%d", raid.ID), ""), http.StatusNotFound), contract.Options{})
}

func TestAppContractRaidDetailAfterTheNight(t *testing.T) {
	w, night, _ := reportNight(t)
	w.as(w.officer)

	contract.AssertFixture(t, "guild.raid.done",
		guildBody(t, w.h.do(http.MethodGet, w.path("/raids/%d", night.ID), ""), http.StatusOK),
		contract.Options{Optional: doneRaidOptional})
}

func TestAppContractSignup(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	w.as(w.member)

	res := w.putSignup(raid.ID, memberKeyA, `{"state":"in","role":"tank","note":"bringing flasks"}`)
	assertGuildCache(t, res, privateCache)
	contract.AssertFixture(t, "guild.signup", guildBody(t, res, http.StatusOK), raidViewOptions("data.raid"))

	contract.AssertFixture(t, "guild.signup.moved",
		guildBody(t, w.putSignup(raid.ID, memberKeyB, `{"state":"in"}`), http.StatusOK), raidViewOptions("data.raid"))
	contract.AssertFixture(t, "guild.signup.error-forbidden",
		guildBody(t, w.putSignup(raid.ID, member2Key, `{"state":"in"}`), http.StatusForbidden), contract.Options{})
	contract.AssertFixture(t, "guild.signup.error-invalid",
		guildBody(t, w.putSignup(raid.ID, memberKeyA, `{"state":"maybe"}`), http.StatusUnprocessableEntity), contract.Options{})
	contract.AssertFixture(t, "guild.signup.error-character-not-found",
		guildBody(t, w.putSignup(raid.ID, "us/hardcore/nobody", `{"state":"in"}`), http.StatusNotFound), contract.Options{})
}

func TestAppContractSignupOnALockedRaid(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	w.as(w.officer)
	guildBody(t, w.h.do(http.MethodPost, w.path("/raids/%d/lock", raid.ID), ""), http.StatusOK)
	w.as(w.member)

	contract.AssertFixture(t, "guild.signup.error-raid-locked",
		guildBody(t, w.putSignup(raid.ID, memberKeyA, `{"state":"in"}`), http.StatusConflict), contract.Options{})
}

func TestAppContractPatchRaid(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	w.as(w.officer)
	body := raidBody(futureStart(), 0)

	res := w.h.do(http.MethodPatch, w.path("/raids/%d", raid.ID), body)
	assertGuildCache(t, res, privateCache)
	contract.AssertFixture(t, "guild.raid.patch", guildBody(t, res, http.StatusOK), raidViewOptions("data"))
	contract.AssertFixture(t, "guild.raid.patch.error-invalid",
		guildBody(t, w.h.do(http.MethodPatch, w.path("/raids/%d", raid.ID), `{"zone":"nowhere"}`), http.StatusUnprocessableEntity),
		contract.Options{})

	w.as(w.member)
	contract.AssertFixture(t, "guild.raid.patch.error-forbidden",
		guildBody(t, w.h.do(http.MethodPatch, w.path("/raids/%d", raid.ID), body), http.StatusForbidden), contract.Options{})
}

func TestAppContractRaidStatusActions(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	w.as(w.officer)

	contract.AssertFixture(t, "guild.raid.patch",
		guildBody(t, w.h.do(http.MethodPost, w.path("/raids/%d/lock", raid.ID), ""), http.StatusOK), raidViewOptions("data"))
	contract.AssertFixture(t, "guild.raid.error-conflict",
		guildBody(t, w.h.do(http.MethodPost, w.path("/raids/%d/lock", raid.ID), ""), http.StatusConflict), contract.Options{})
}

func TestAppContractLoot(t *testing.T) {
	w := newRaidWorld(t, true)
	w.as(w.member)

	res := w.h.do(http.MethodGet, w.path("/loot"), "")
	assertGuildCache(t, res, privateCache)
	contract.AssertFixture(t, "guild.loot", guildBody(t, res, http.StatusOK), contract.Options{Optional: lootOptional})

	w.as(seedUser(t, w.h.pool, "loot-stranger@example.com"))
	contract.AssertFixture(t, "guild.home.error-forbidden",
		guildBody(t, w.h.do(http.MethodGet, w.path("/loot"), ""), http.StatusForbidden), contract.Options{})
}
