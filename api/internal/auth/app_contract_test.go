package auth

import (
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/contract"
)

// meOptional are the /v1/me fields Go omits when empty or unknown.
var meOptional = []string{
	"data.bnet_imported_at", "data.main_character_key", "data.server_sims", "data.notifications",
	"data.seat", "data.seat_offer", "data.ended_seat",
	"data.characters[].class", "data.characters[].realm", "data.characters[].level",
	"data.characters[].faction", "data.characters[].race", "data.characters[].gender",
	"data.characters[].item_level", "data.characters[].avatar_url", "data.characters[].render_url",
	"data.characters[].build", "data.characters[].guild", "data.characters[].spec",
	"data.characters[].guild.rank_index", "data.characters[].guild.attendance",
	"data.guilds[].rank", "data.guilds[].consent", "data.guilds[].plan",
	"data.guilds[].ruleset_label", "data.guilds[].slug",
}

func appBody(t *testing.T, res *http.Response, want int) []byte {
	t.Helper()
	defer res.Body.Close()
	raw, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != want {
		t.Fatalf("%s = %d, want %d: %s", res.Request.URL.Path, res.StatusCode, want, raw)
	}
	return raw
}

// seedRaider gives an account the character, guild seat and battletag the app's home reads.
func (h *harness) seedRaider(t *testing.T, email string) int64 {
	t.Helper()
	uid := h.seedUser(t, email)
	pool := h.svc.Store.Pool
	if _, err := pool.Exec(t.Context(), `update users set battletag = 'Thrall#1234' where id = $1`, uid); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(t.Context(),
		`insert into characters (key, region, ruleset, name, class, user_id, realm_name, level, faction, source, imported_at, refreshed_at)
		 values ('us/pvp/obnoxious-yell', 'us', 'pvp', 'Obnoxious Yell', 'warrior', $1, 'Whitemane', 58, 'horde', 'bnet', now(), now())`,
		uid); err != nil {
		t.Fatal(err)
	}
	gid := h.seedGuild(t, "OLYMPUS XXVII")
	h.seedGuildMember(t, gid, uid, "officer", true)
	if _, err := pool.Exec(t.Context(),
		`insert into guild_characters (guild_id, character_key, user_id, rank, rank_index, source, verified_by, verified_at)
		 values ($1, 'us/pvp/obnoxious-yell', $2, 'officer', 1, 'bnet', 'bnet', now())`, gid, uid); err != nil {
		t.Fatal(err)
	}
	return uid
}

func TestAppContractMe(t *testing.T) {
	h := newHarness(t)
	uid := h.seedRaider(t, "contract-me@example.com")

	res := h.sessionRequest(t, http.MethodGet, "/v1/me", uid)
	if got := res.Header.Get("Cache-Control"); got != "private, no-cache" {
		t.Errorf("Cache-Control = %q, want private, no-cache", got)
	}
	contract.AssertFixture(t, "me", appBody(t, res, http.StatusOK), contract.Options{Optional: meOptional})
}

func TestAppContractMePremium(t *testing.T) {
	h := newHarness(t)
	uid := h.seedRaider(t, "contract-premium@example.com")
	h.seedStripePremium(t, uid, "premium_founder_yearly", 4000, time.Now().Add(24*time.Hour))

	contract.AssertFixture(t, "me.premium",
		appBody(t, h.sessionRequest(t, http.MethodGet, "/v1/me", uid), http.StatusOK),
		contract.Options{Optional: meOptional})
}

func TestAppContractMeNewAccount(t *testing.T) {
	h := newHarness(t)
	uid := h.seedUser(t, "contract-new@example.com")

	contract.AssertFixture(t, "me.empty",
		appBody(t, h.sessionRequest(t, http.MethodGet, "/v1/me", uid), http.StatusOK),
		contract.Options{Optional: meOptional})
}

func TestAppContractMeSignedOut(t *testing.T) {
	h := newHarness(t)

	res, err := http.Get(h.server.URL + "/v1/me")
	if err != nil {
		t.Fatal(err)
	}
	contract.AssertFixture(t, "error-unauthorized", appBody(t, res, http.StatusUnauthorized), contract.Options{})
}

func TestAppContractPatchMe(t *testing.T) {
	h := newHarness(t)
	h.signIn(t, "contract-patch@example.com")

	contract.AssertFixture(t, "me.patch",
		appBody(t, h.do(t, http.MethodPatch, "/v1/me", `{"anonymize":true}`), http.StatusOK),
		contract.Options{Optional: meOptional})
	contract.AssertFixture(t, "me.patch.error-invalid",
		appBody(t, h.do(t, http.MethodPatch, "/v1/me", `{}`), http.StatusBadRequest), contract.Options{})
	contract.AssertFixture(t, "me.patch.error-not-your-character",
		appBody(t, h.do(t, http.MethodPatch, "/v1/me", `{"main_character_key":"us/pvp/nobody"}`), http.StatusBadRequest),
		contract.Options{})
}
