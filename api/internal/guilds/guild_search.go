// api/internal/guilds/guild_search.go
//
// GET /v1/guilds/search: the guild lookup's "Did you mean" for a stranger. Public, behind the
// router's per-client limit like every other public guild read: a case-insensitive contains
// match on the guild name over every guild that has a page (any guilds row), prefix matches
// first, then by name, at most guildSearchLimit rows.
package guilds

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

const (
	// guildSearchMinChars and guildSearchMaxChars bound the trimmed query; outside them is a 400.
	guildSearchMinChars = 3
	guildSearchMaxChars = 40
	// guildSearchLimit is the most rows a search answers.
	guildSearchLimit = 5
	// guildSearchCacheSeconds is the edge cache, the character typeahead's.
	guildSearchCacheSeconds = 60
)

// GuildSearchHit is one "Did you mean" row: the guild and its page's site path (GuildPagePath).
type GuildSearchHit struct {
	Name    string `json:"name"`
	Region  string `json:"region"`
	Ruleset string `json:"ruleset"`
	Path    string `json:"path"`
}

// GuildSearch is a validated search: Text trimmed, Region and Ruleset empty or valid.
type GuildSearch struct {
	Text, Region, Ruleset string
}

// likeEscaper makes a string match literally inside a LIKE pattern (backslash is the default escape).
var likeEscaper = strings.NewReplacer(`\`, `\\`, `%`, `\%`, `_`, `\_`)

// searchGuildsSQL: a contains match on lower(name), the realm filters when given, prefix
// matches first, then name, region and ruleset for a stable order.
const searchGuildsSQL = `
select g.name, g.region, g.ruleset
from guilds g
where lower(g.name) like '%' || lower($1) || '%'
  and ($2 = '' or g.region = $2)
  and ($3 = '' or g.ruleset = $3)
order by lower(g.name) like lower($1) || '%' desc, lower(g.name), g.region, g.ruleset, g.id
limit $4`

// SearchGuilds answers a search's rows; never nil.
func (s *Store) SearchGuilds(ctx context.Context, q GuildSearch) ([]GuildSearchHit, error) {
	rows, err := s.Pool.Query(ctx, searchGuildsSQL, likeEscaper.Replace(q.Text), q.Region, q.Ruleset, guildSearchLimit)
	if err != nil {
		return nil, fmt.Errorf("guilds: search: %w", err)
	}
	defer rows.Close()
	hits := []GuildSearchHit{}
	for rows.Next() {
		var g Guild
		if err := rows.Scan(&g.Name, &g.Region, &g.Ruleset); err != nil {
			return nil, fmt.Errorf("guilds: search: %w", err)
		}
		hits = append(hits, GuildSearchHit{Name: g.Name, Region: g.Region, Ruleset: g.Ruleset, Path: GuildPagePath(g)})
	}
	return hits, rows.Err()
}

// guildSearchQuery validates the request; it writes the 400 itself when a parameter is bad.
func guildSearchQuery(w http.ResponseWriter, r *http.Request) (GuildSearch, bool) {
	p := r.URL.Query()
	q := GuildSearch{
		Text:    strings.TrimSpace(p.Get("q")),
		Region:  strings.ToLower(p.Get("region")),
		Ruleset: strings.ToLower(p.Get("ruleset")),
	}
	if n := utf8.RuneCountInString(q.Text); n < guildSearchMinChars || n > guildSearchMaxChars {
		span := strconv.Itoa(guildSearchMinChars) + " to " + strconv.Itoa(guildSearchMaxChars) + " characters"
		httpx.WriteError(w, r, http.StatusBadRequest, "bad_request", "q must be "+span, map[string]string{"q": span})
		return q, false
	}
	if q.Region != "" && !character.ValidRegion(q.Region) {
		httpx.WriteError(w, r, http.StatusBadRequest, "bad_request", "that is not a region",
			map[string]string{"region": strings.Join(character.Regions, ", ")})
		return q, false
	}
	if q.Ruleset != "" && !character.ValidRuleset(q.Ruleset) {
		httpx.WriteError(w, r, http.StatusBadRequest, "bad_request", "that is not a ruleset",
			map[string]string{"ruleset": strings.Join(character.Rulesets, ", ")})
		return q, false
	}
	return q, true
}

// searchGuilds is GET /v1/guilds/search.
func (s *Service) searchGuilds(w http.ResponseWriter, r *http.Request) {
	q, ok := guildSearchQuery(w, r)
	if !ok {
		return
	}
	hits, err := s.Store.SearchGuilds(r.Context(), q)
	if err != nil {
		s.fail(w, r, "guild_search", err, "could not search guilds just now")
		return
	}
	httpx.CachePublic(w, guildSearchCacheSeconds*time.Second, 0)
	httpx.WriteOK(w, r, http.StatusOK, map[string]any{"rows": hits})
}
