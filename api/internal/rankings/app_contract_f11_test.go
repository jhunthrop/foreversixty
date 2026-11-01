package rankings

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/contract"
	"github.com/jhunthrop/foreversixty/api/internal/engine"
)

// rankingRowOptional are the ranking row fields an answer leaves out or null: the guild of a
// report filed to none, a fight not yet scored.
var rankingRowOptional = []string{
	"data.rows[].guild", "data.character_rank.row.guild", "data.character_rank",
	"data.rows[].player.class", "data.rows[].player.spec",
	"data.character_rank.row.player.class", "data.character_rank.row.player.spec",
}

func TestAppContractRankings(t *testing.T) {
	h := newHarness(t)
	serveContract(t, h)
	_, _, key := h.seedRankedFight(t)

	res := h.get(fmt.Sprintf("/v1/rankings?encounter=%d&distinct=character&character=%s", executionEncounterID, key))
	if got, want := res.Header.Get("Cache-Control"), "public, max-age=30, stale-while-revalidate=300"; got != want {
		t.Errorf("Cache-Control = %q, want %q", got, want)
	}
	contract.AssertFixture(t, "rankings.list", readAll(t, res, http.StatusOK),
		contract.Options{Optional: rankingRowOptional})
}

func TestAppContractRankingsEmpty(t *testing.T) {
	h := newHarness(t)
	serveContract(t, h)
	h.seedRankedFight(t)

	contract.AssertFixture(t, "rankings.list.empty",
		readAll(t, h.get(fmt.Sprintf("/v1/rankings?encounter=%d&distinct=character&region=eu", executionEncounterID)), http.StatusOK),
		contract.Options{Optional: rankingRowOptional})
}

func TestAppContractRankingPercentile(t *testing.T) {
	h := newHarness(t)
	serveContract(t, h)
	h.seedRankedFight(t)

	contract.AssertFixture(t, "rankings.percentile",
		readAll(t, h.get(fmt.Sprintf("/v1/rankings/percentile?encounter=%d&difficulty=8&value=100", executionEncounterID)), http.StatusOK),
		contract.Options{})
}

func TestAppContractRankingEncounters(t *testing.T) {
	h := newHarness(t)
	serveContract(t, h)
	h.seedRankedFight(t)
	if _, err := h.pool.Exec(t.Context(), `update reports set log_zone = 'Blackrock Depths'`); err != nil {
		t.Fatal(err)
	}

	contract.AssertFixture(t, "rankings.encounters",
		readAll(t, h.get("/v1/rankings/encounters"), http.StatusOK), contract.Options{})
}

func TestAppContractWatching(t *testing.T) {
	h := newHarness(t)
	serveAs(t, h, appActor(h.owner))
	contract.AssertFixture(t, "watching.list.empty", h.sent(http.MethodGet, "/v1/me/watching", http.StatusOK), contract.Options{})

	h.seedOwnedCharacter("baelgrim", &h.owner)
	contract.AssertFixture(t, "watching.put", h.sent(http.MethodPut, "/v1/me/watching/us/hardcore/baelgrim", http.StatusOK), contract.Options{})
	contract.AssertFixture(t, "watching.list", h.sent(http.MethodGet, "/v1/me/watching", http.StatusOK), contract.Options{})
	contract.AssertFixture(t, "watching.put.error-not-found",
		h.sent(http.MethodPut, "/v1/me/watching/us/hardcore/nobody", http.StatusNotFound), contract.Options{})
}

func TestAppContractWatchingLimit(t *testing.T) {
	h := newHarness(t)
	serveAs(t, h, appActor(h.owner))
	for i := range WatchLimit {
		if _, err := h.pool.Exec(t.Context(),
			`insert into watched_characters (user_id, character_key) values ($1, $2)`,
			h.owner, fmt.Sprintf("us/hardcore/filler-%02d", i)); err != nil {
			t.Fatal(err)
		}
	}
	h.seedOwnedCharacter("baelgrim", &h.owner)
	contract.AssertFixture(t, "watching.put.error-limit",
		h.sent(http.MethodPut, "/v1/me/watching/us/hardcore/baelgrim", http.StatusConflict), contract.Options{})
}

func TestAppContractCharacterCardRanked(t *testing.T) {
	h := newHarness(t)
	serveContract(t, h)
	h.seedReport("report-one")
	h.seedFight("report-one", 1, engine.FixtureBase, nil)
	h.seedOwnedCharacter("baelgrim", &h.owner)

	body := readAll(t, h.get("/v1/characters/us/hardcore/baelgrim"), http.StatusOK)
	for _, name := range []string{"character.card.ranked", "character.card.yours"} {
		contract.AssertFixture(t, name, body, contract.Options{Optional: characterOptional})
	}
}

// sent sends one request and answers its body, failing unless the status is want.
func (h *harness) sent(method, path string, want int) []byte {
	h.t.Helper()
	res, body := h.send(method, path)
	if res.StatusCode != want {
		h.t.Fatalf("%s %s = %d, want %d: %s", method, path, res.StatusCode, want, body)
	}
	return body
}
