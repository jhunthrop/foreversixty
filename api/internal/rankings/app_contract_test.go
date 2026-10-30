package rankings

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/contract"
	"github.com/jhunthrop/foreversixty/api/internal/engine"
)

// panelReader answers the Forever Score and gear panel a character with logs, an export and a
// spec has, so the contract sees every block the app reads filled.
type panelReader struct{}

func (panelReader) ScoreAndGear(context.Context, string) (ForeverScore, CharacterGear, error) {
	score, tier, asOf := 74, "common", "2026-10-30"
	gap, chance := 12.6, 0.12
	worn, level := "Green Lens", 49
	pick := 18814
	return ForeverScore{
		Score: &score, Tier: &tier, AsOf: &asOf,
		Components: &ScoreComponents{Gear: 20, Experience: 14, Parses: 22, Attendance: 12, Discipline: 6},
		Experience: &ScoreExperience{Dungeons: 9, RaidKills: 4, Alts: 1},
	}, CharacterGear{
		Slots: []GearSlot{
			{Slot: "head", Worn: &GearWorn{ID: 10504, Name: &worn, ItemLevel: &level},
				Pick: &GearPick{ID: pick, Name: "Choker of the Fire Lord", Source: "Molten Core: Ragnaros", Chance: &chance}, Gap: &gap},
			{Slot: "neck", Worn: nil, Pick: nil, Gap: nil},
		},
		GapTotal: &gap, Spec: "warrior-fury", SpecSource: "talents",
	}, nil
}

func serveContract(t *testing.T, h *harness) {
	t.Helper()
	mux := http.NewServeMux()
	Mount(mux, &Service{Store: h.store, Log: slog.New(slog.NewTextHandler(io.Discard, nil)),
		Now: func() time.Time { return engine.FixtureBase }, ScoreGear: panelReader{}})
	h.server = httptest.NewServer(mux)
	t.Cleanup(h.server.Close)
}

func readAll(t *testing.T, res *http.Response, want int) []byte {
	t.Helper()
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != want {
		t.Fatalf("%s = %d, want %d: %s", res.Request.URL.Path, res.StatusCode, want, body)
	}
	return body
}

// characterOptional are the character card fields Go omits when empty.
var characterOptional = []string{
	"data.character.class", "data.character.avatar_url", "data.character.render_url",
	"data.best[].percentile", "data.best[].spec", "data.history[].percentile", "data.history[].spec",
	"data.history[].talent_split", "data.builds_seen[].spec",
}

func TestAppContractCharacterCard(t *testing.T) {
	h := newHarness(t)
	serveContract(t, h)
	h.seedReport("report-one")
	h.seedFight("report-one", 1, engine.FixtureBase, nil)
	h.seedOwnedCharacter("baelgrim", &h.owner)

	res := h.get("/v1/characters/us/hardcore/baelgrim")
	if got, want := res.Header.Get("Cache-Control"), "public, max-age=30, stale-while-revalidate=300"; got != want {
		t.Errorf("Cache-Control = %q, want %q", got, want)
	}
	contract.AssertFixture(t, "character.card", readAll(t, res, http.StatusOK),
		contract.Options{Optional: characterOptional})
}

func TestAppContractCharacterCardWithoutLogs(t *testing.T) {
	h := newHarness(t)
	serve(t, h)
	h.seedOwnedCharacter("freshling", &h.owner)

	contract.AssertFixture(t, "character.card.empty", readAll(t, h.get("/v1/characters/us/hardcore/freshling"), http.StatusOK),
		contract.Options{Optional: characterOptional})
}

func TestAppContractCharacterCardNotFound(t *testing.T) {
	h := newHarness(t)
	serve(t, h)

	contract.AssertFixture(t, "character.card.error-not-found",
		readAll(t, h.get("/v1/characters/us/hardcore/nobody"), http.StatusNotFound), contract.Options{})
}

func TestAppContractCharacterSearch(t *testing.T) {
	h := newHarness(t)
	serve(t, h)
	class, level := "warrior", 58
	h.seedSearchCharacter("us", "pvp", "Obnoxious Yell", &class, &level, nil)

	res := h.get("/v1/characters/search?q=obn")
	if got, want := res.Header.Get("Cache-Control"), "public, max-age=60, stale-while-revalidate=300"; got != want {
		t.Errorf("Cache-Control = %q, want %q", got, want)
	}
	contract.AssertFixture(t, "character.search", readAll(t, res, http.StatusOK), contract.Options{})
}

func TestAppContractCharacterSearchEmpty(t *testing.T) {
	h := newHarness(t)
	serve(t, h)

	contract.AssertFixture(t, "character.search.empty",
		readAll(t, h.get("/v1/characters/search?q=zz"), http.StatusOK), contract.Options{})
}

func TestAppContractCharacterSearchBadRegion(t *testing.T) {
	h := newHarness(t)
	serve(t, h)

	contract.AssertFixture(t, "character.search.error-bad-request",
		readAll(t, h.get("/v1/characters/search?q=obn&region=xx"), http.StatusBadRequest), contract.Options{})
}
