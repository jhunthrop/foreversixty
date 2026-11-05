package guilds

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
)

// searchHits runs GET /v1/guilds/search with params, signed out, and answers its rows.
func searchHits(t *testing.T, h *httpHarness, params url.Values) []GuildSearchHit {
	t.Helper()
	res := h.do(http.MethodGet, "/v1/guilds/search?"+params.Encode(), "")
	var out struct {
		Rows []GuildSearchHit `json:"rows"`
	}
	h.data(res, &out)
	return out.Rows
}

func hitNames(hits []GuildSearchHit) string {
	names := make([]string, len(hits))
	for i, h := range hits {
		names[i] = h.Name + "@" + h.Region + "/" + h.Ruleset
	}
	return strings.Join(names, ",")
}

// A stranger's "Did you mean": a case-insensitive contains match, prefix matches first, then
// by name, at most five, each with the guild page's path; region and ruleset narrow it.
func TestGuildSearchSuggestsGuildsByName(t *testing.T) {
	h := newHTTPHarness(t)
	h.actor = auth.Actor{}
	for _, g := range []struct{ region, ruleset, name string }{
		{"us", "pvp", "Lifecycle Test"}, {"us", "hardcore", "The Lifecycle Club"}, {"eu", "pvp", "lifecycle"},
		{"us", "normal", "Life Cycle"}, {"us", "pvp", "Cycle of Life"}, {"us", "pvp", "Zz Lifecycle A"},
		{"us", "pvp", "Zz Lifecycle B"}, {"us", "pvp", "Zz Lifecycle C"}, {"us", "pvp", "100%_Lifecycle"},
	} {
		if _, err := h.pool.Exec(t.Context(), `insert into guilds (region, ruleset, name) values ($1, $2, $3)`,
			g.region, g.ruleset, g.name); err != nil {
			t.Fatal(err)
		}
	}
	hits := searchHits(t, h, url.Values{"q": {"  LIFECYCLE "}})
	if got, want := hitNames(hits), "lifecycle@eu/pvp,Lifecycle Test@us/pvp,100%_Lifecycle@us/pvp,The Lifecycle Club@us/hardcore,Zz Lifecycle A@us/pvp"; got != want {
		t.Fatalf("rows = %s, want %s", got, want)
	}
	if hits[1].Path != "/guild/us/pvp/lifecycle-test" {
		t.Fatalf("path = %q, want the site's guild path", hits[1].Path)
	}
	if got := hitNames(searchHits(t, h, url.Values{"q": {"lifecycle"}, "region": {"US"}, "ruleset": {"hardcore"}})); got != "The Lifecycle Club@us/hardcore" {
		t.Fatalf("narrowed rows = %s", got)
	}
	if got := searchHits(t, h, url.Values{"q": {"0%_l"}}); len(got) != 1 || got[0].Name != "100%_Lifecycle" {
		t.Fatalf("a literal %%_ = %s, want 100%%_Lifecycle", hitNames(got))
	}
	if got := searchHits(t, h, url.Values{"q": {"e%_c"}}); len(got) != 0 {
		t.Fatalf("a LIKE wildcard matched as a wildcard: %s", hitNames(got))
	}
	if got := searchHits(t, h, url.Values{"q": {"nothing like it"}}); got == nil || len(got) != 0 {
		t.Fatalf("no match = %v, want []", got)
	}
}

// The handler's row has exactly the contract fixture's keys.
func TestGuildSearchMatchesItsFixture(t *testing.T) {
	h := newHTTPHarness(t)
	h.actor = auth.Actor{}
	seedGuild(t, h.pool, "Lifecycle Test")
	rowKeys := func(raw []byte) string {
		var body struct {
			Data struct {
				Rows []map[string]any `json:"rows"`
			} `json:"data"`
		}
		if err := json.Unmarshal(raw, &body); err != nil || len(body.Data.Rows) == 0 {
			t.Fatalf("body %s: %v", raw, err)
		}
		keys := make([]string, 0, len(body.Data.Rows[0]))
		for k := range body.Data.Rows[0] {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		return strings.Join(keys, ",")
	}
	fixture, err := os.ReadFile("../../../docs/contracts/fixtures/app/guild.search.json")
	if err != nil {
		t.Fatal(err)
	}
	res := h.do(http.MethodGet, "/v1/guilds/search?q=cycle", "")
	defer res.Body.Close()
	got, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	if rowKeys(got) != rowKeys(fixture) {
		t.Fatalf("row keys %s, fixture %s", rowKeys(got), rowKeys(fixture))
	}
}

func TestGuildSearchRefusesABadQuery(t *testing.T) {
	h := newHTTPHarness(t)
	h.actor = auth.Actor{}
	for name, params := range map[string]url.Values{
		"too short":   {"q": {" ab "}},
		"too long":    {"q": {strings.Repeat("a", 41)}},
		"missing":     {},
		"bad region":  {"q": {"life"}, "region": {"mars"}},
		"bad ruleset": {"q": {"life"}, "ruleset": {"easy"}},
	} {
		res := h.do(http.MethodGet, "/v1/guilds/search?"+params.Encode(), "")
		res.Body.Close()
		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("%s = %d, want 400", name, res.StatusCode)
		}
	}
}
