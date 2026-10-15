package rankings

import (
	"context"
	"net/http"
	"strings"
	"testing"
)

type searchRow struct {
	Key     string  `json:"key"`
	Region  string  `json:"region"`
	Ruleset string  `json:"ruleset"`
	Name    string  `json:"name"`
	Class   *string `json:"class"`
	Level   *int    `json:"level"`
	Spec    *string `json:"spec"`
}

type searchBody struct {
	Rows []searchRow `json:"rows"`
}

func (h *harness) seedSearchCharacter(region, ruleset, name string, class *string, level *int, userID *int64) {
	h.t.Helper()
	key := region + "/" + ruleset + "/" + strings.ToLower(name)
	if _, err := h.pool.Exec(context.Background(),
		`insert into characters (key, region, ruleset, name, class, level, user_id)
		 values ($1, $2, $3, $4, $5, $6, $7)`, key, region, ruleset, name, class, level, userID); err != nil {
		h.t.Fatal(err)
	}
}

func (h *harness) search(query string) searchBody {
	h.t.Helper()
	res := h.get("/v1/characters/search?" + query)
	if res.StatusCode != http.StatusOK {
		h.t.Fatalf("search %q: status %d", query, res.StatusCode)
	}
	var out searchBody
	h.data(res, &out)
	return out
}

func assertNames(t *testing.T, got searchBody, want ...string) {
	t.Helper()
	names := make([]string, len(got.Rows))
	for i, r := range got.Rows {
		names[i] = r.Region + "/" + r.Ruleset + "/" + r.Name
	}
	if g, w := strings.Join(names, ","), strings.Join(want, ","); g != w {
		t.Errorf("rows = %s, want %s", g, w)
	}
}

func TestSearchRouteIsNotTheCharacterPage(t *testing.T) {
	h := newHarness(t)
	serve(t, h)
	res := h.get("/v1/characters/search?q=ba")
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want 200 (the literal route must win over the three-segment one)", res.StatusCode)
	}
}

func TestSearchShortQueryAnswersEmptyRows(t *testing.T) {
	h := newHarness(t)
	serve(t, h)
	h.seedSearchCharacter("us", "normal", "Baelgrim", nil, nil, nil)
	for _, q := range []string{"", "b", "%20b%20", "%20%20"} {
		body := h.search("q=" + q)
		if body.Rows == nil || len(body.Rows) != 0 {
			t.Errorf("q=%q rows = %#v, want empty non-nil", q, body.Rows)
		}
	}
}

func TestSearchMatchesPrefixCaseInsensitively(t *testing.T) {
	h := newHarness(t)
	serve(t, h)
	class, level := "warrior", 42
	h.seedSearchCharacter("us", "normal", "Baelgrim", &class, &level, nil)
	h.seedSearchCharacter("us", "normal", "Abael", nil, nil, nil)
	got := h.search("q=%20BAE")
	assertNames(t, got, "us/normal/Baelgrim")
	row := got.Rows[0]
	if row.Key != "us/normal/baelgrim" || row.Class == nil || *row.Class != "warrior" ||
		row.Level == nil || *row.Level != 42 || row.Spec != nil {
		t.Errorf("row = %+v", row)
	}
}

func TestSearchTreatsLikeWildcardsLiterally(t *testing.T) {
	h := newHarness(t)
	serve(t, h)
	h.seedSearchCharacter("us", "normal", "Baelgrim", nil, nil, nil)
	h.seedSearchCharacter("us", "normal", "Ba_x", nil, nil, nil)
	assertNames(t, h.search("q=ba_"), "us/normal/Ba_x")
	assertNames(t, h.search("q=%25%25"))
}

func TestSearchCutsLongQueryTo24Characters(t *testing.T) {
	h := newHarness(t)
	serve(t, h)
	name := strings.Repeat("a", 24)
	h.seedSearchCharacter("us", "normal", name, nil, nil, nil)
	assertNames(t, h.search("q="+name+"zzzz"), "us/normal/"+name)
}

func TestSearchOrdersExactThenNameThenRegionThenRuleset(t *testing.T) {
	h := newHarness(t)
	serve(t, h)
	h.seedSearchCharacter("us", "normal", "Bobby", nil, nil, nil)
	h.seedSearchCharacter("us", "normal", "Bob", nil, nil, nil)
	h.seedSearchCharacter("eu", "normal", "Bob", nil, nil, nil)
	h.seedSearchCharacter("us", "hardcore", "Bob", nil, nil, nil)
	h.seedSearchCharacter("us", "normal", "Bobzed", nil, nil, nil)
	h.seedSearchCharacter("us", "normal", "Bobalpha", nil, nil, nil)
	assertNames(t, h.search("q=bob"),
		"us/normal/Bob", "us/hardcore/Bob", "eu/normal/Bob",
		"us/normal/Bobalpha", "us/normal/Bobby", "us/normal/Bobzed")
}

func TestSearchExactMatchBeatsShorterAlphabeticalPeers(t *testing.T) {
	h := newHarness(t)
	serve(t, h)
	h.seedSearchCharacter("us", "normal", "Boba", nil, nil, nil)
	h.seedSearchCharacter("us", "normal", "Bo", nil, nil, nil)
	h.seedSearchCharacter("us", "normal", "Bob", nil, nil, nil)
	assertNames(t, h.search("q=bob"), "us/normal/Bob", "us/normal/Boba")
	assertNames(t, h.search("q=bo"), "us/normal/Bo", "us/normal/Bob", "us/normal/Boba")
}

func TestSearchFiltersByRegionAndRuleset(t *testing.T) {
	h := newHarness(t)
	serve(t, h)
	h.seedSearchCharacter("us", "normal", "Bob", nil, nil, nil)
	h.seedSearchCharacter("eu", "normal", "Bob", nil, nil, nil)
	h.seedSearchCharacter("eu", "hardcore", "Bob", nil, nil, nil)
	assertNames(t, h.search("q=bob&region=EU"), "eu/normal/Bob", "eu/hardcore/Bob")
	assertNames(t, h.search("q=bob&region=eu&ruleset=hardcore"), "eu/hardcore/Bob")
}

func TestSearchLimit(t *testing.T) {
	h := newHarness(t)
	serve(t, h)
	for _, n := range []string{"Bob1", "Bob2", "Bob3", "Bob4", "Bob5", "Bob6", "Bob7", "Bob8", "Bob9", "BobA"} {
		h.seedSearchCharacter("us", "normal", n, nil, nil, nil)
	}
	if got := len(h.search("q=bob").Rows); got != 8 {
		t.Errorf("default limit gave %d rows, want 8", got)
	}
	if got := len(h.search("q=bob&limit=10").Rows); got != 10 {
		t.Errorf("limit=10 gave %d rows", got)
	}
	if got := len(h.search("q=bob&limit=1").Rows); got != 1 {
		t.Errorf("limit=1 gave %d rows", got)
	}
}

func TestSearchRejectsBadParameters(t *testing.T) {
	h := newHarness(t)
	serve(t, h)
	for _, q := range []string{"q=bob&region=xx", "q=bob&ruleset=nope", "q=bob&limit=0",
		"q=bob&limit=11", "q=bob&limit=x", "q=b&region=xx"} {
		res := h.get("/v1/characters/search?" + q)
		res.Body.Close()
		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("%s: status = %d, want 400", q, res.StatusCode)
		}
	}
}

func TestSearchHidesAnonymisedAccounts(t *testing.T) {
	h := newHarness(t)
	serve(t, h)
	var anonID int64
	if err := h.pool.QueryRow(context.Background(),
		`insert into users (email, anonymize) values ('hidden@example.com', true) returning id`).Scan(&anonID); err != nil {
		t.Fatal(err)
	}
	h.seedSearchCharacter("us", "normal", "Bobhidden", nil, nil, &anonID)
	h.seedSearchCharacter("us", "normal", "Bobshown", nil, nil, &h.owner)
	h.seedSearchCharacter("us", "normal", "Bobfree", nil, nil, nil)
	assertNames(t, h.search("q=bob"), "us/normal/Bobfree", "us/normal/Bobshown")
}

func TestSearchSetsPublicCacheControl(t *testing.T) {
	h := newHarness(t)
	serve(t, h)
	res := h.get("/v1/characters/search?q=bob")
	res.Body.Close()
	want := "public, max-age=60, stale-while-revalidate=300"
	if got := res.Header.Get("Cache-Control"); got != want {
		t.Errorf("Cache-Control = %q, want %q", got, want)
	}
}

func TestSearchTermNormalisation(t *testing.T) {
	cases := map[string]string{" ab ": "ab", strings.Repeat("é", 30): strings.Repeat("é", 24), "a": "a"}
	for in, want := range cases {
		if got := searchTerm(in); got != want {
			t.Errorf("searchTerm(%q) = %q, want %q", in, got, want)
		}
	}
	if got := escapeLike(`a%b_c\d`); got != `a\%b\_c\\d` {
		t.Errorf("escapeLike = %q", got)
	}
}
