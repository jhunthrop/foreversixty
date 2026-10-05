package rankings

import (
	"context"
	"strconv"
	"testing"
)

// TestGuildPageCarriesTheCrestURL checks the public guild page builds
// crest_url from guilds.crest_key the same way GET .../home does
// (guilds.CrestURL), and reads null when no crest is set - the contract's
// own "`crest_url ?? faction logo`" rule, shared by both readers.
func TestGuildPageCarriesTheCrestURL(t *testing.T) {
	h := newHarness(t)
	h.store.APIBaseURL = "https://api.forever.test"
	serve(t, h)

	res := h.get("/v1/guilds/us/hardcore/forever-sixty")
	var before struct {
		Guild struct {
			CrestURL *string `json:"crest_url"`
		} `json:"guild"`
	}
	h.data(res, &before)
	if before.Guild.CrestURL != nil {
		t.Fatalf("crest_url with no crest set = %v, want null", before.Guild.CrestURL)
	}

	if _, err := h.pool.Exec(context.Background(),
		`update guilds set crest_key = 'guilds/1/crest/abcdef0123456789.webp' where id = $1`, h.guildID); err != nil {
		t.Fatal(err)
	}

	res = h.get("/v1/guilds/us/hardcore/forever-sixty")
	var after struct {
		Guild struct {
			CrestURL *string `json:"crest_url"`
		} `json:"guild"`
	}
	h.data(res, &after)
	want := "https://api.forever.test/v1/guilds/" + strconv.FormatInt(h.guildID, 10) + "/crest.webp?v=abcdef0123456789"
	if after.Guild.CrestURL == nil || *after.Guild.CrestURL != want {
		t.Fatalf("crest_url = %v, want %q", after.Guild.CrestURL, want)
	}
}
