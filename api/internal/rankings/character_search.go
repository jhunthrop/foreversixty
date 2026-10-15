package rankings

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
	// searchMinChars is the shortest prefix worth a lookup; a shorter one answers no rows.
	searchMinChars = 2
	// searchMaxChars is where a longer prefix is cut.
	searchMaxChars = 24
	// searchDefaultLimit and searchMaxLimit bound the rows a lookup returns.
	searchDefaultLimit = 8
	searchMaxLimit     = 10
	// searchCacheSeconds is the typeahead's edge cache; it shares the boards' stale window.
	searchCacheSeconds = 60
)

// CharacterHit is one row of the character lookup typeahead.
type CharacterHit struct {
	Key     string  `json:"key"`
	Region  string  `json:"region"`
	Ruleset string  `json:"ruleset"`
	Name    string  `json:"name"`
	Class   *string `json:"class"`
	// Level is the Battle.net import's level (characters.level, the same column Store.identity
	// reads), null when it is unknown. The addon export's level= section, which identity also
	// weighs, lives in a text blob that cannot be decoded in SQL, so it is not consulted here.
	Level *int `json:"level"`
	// Spec is always null: a character's spec is derived from its fights, not one cheap join.
	Spec *string `json:"spec"`
}

// CharacterSearch is a prefix lookup. Prefix is already trimmed and cut (searchTerm); Region
// and Ruleset are empty or valid.
type CharacterSearch struct {
	Prefix  string
	Region  string
	Ruleset string
	Limit   int
}

// searchTerm trims q and cuts it to searchMaxChars characters.
func searchTerm(q string) string {
	q = strings.TrimSpace(q)
	if utf8.RuneCountInString(q) <= searchMaxChars {
		return q
	}
	return string([]rune(q)[:searchMaxChars])
}

// escapeLike makes s match literally inside a LIKE pattern (backslash is LIKE's default escape).
func escapeLike(s string) string {
	return strings.NewReplacer(`\`, `\\`, `%`, `\%`, `_`, `\_`).Replace(s)
}

// searchCharactersSQL: a case-insensitive prefix on the name (served by characters_name_prefix_idx),
// an anonymised account's characters left out as Store.Character does (users.anonymize), the
// exact name first, then name, region us, ruleset normal.
const searchCharactersSQL = `
select c.key, c.region, c.ruleset, c.name, c.class, case when c.level > 0 then c.level end
from characters c
left join users u on u.id = c.user_id
where lower(c.name) like lower($1) || '%'
  and ($2 = '' or c.region = $2)
  and ($3 = '' or c.ruleset = $3)
  and not coalesce(u.anonymize, false)
order by lower(c.name) = lower($4) desc, c.name, c.region = 'us' desc, c.ruleset = 'normal' desc, c.key
limit $5`

// SearchCharacters is the character lookup behind the site's typeahead. It never returns nil.
func (s *Store) SearchCharacters(ctx context.Context, q CharacterSearch) ([]CharacterHit, error) {
	rows, err := s.Pool.Query(ctx, searchCharactersSQL,
		escapeLike(q.Prefix), q.Region, q.Ruleset, q.Prefix, q.Limit)
	if err != nil {
		return nil, fmt.Errorf("rankings: search characters: %w", err)
	}
	defer rows.Close()
	hits := []CharacterHit{}
	for rows.Next() {
		var h CharacterHit
		if err := rows.Scan(&h.Key, &h.Region, &h.Ruleset, &h.Name, &h.Class, &h.Level); err != nil {
			return nil, fmt.Errorf("rankings: scan character search: %w", err)
		}
		hits = append(hits, h)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rankings: search characters: %w", err)
	}
	return hits, nil
}

// searchLimit reads the limit parameter: absent is the default, else 1 to searchMaxLimit.
func searchLimit(v string) (int, bool) {
	if v == "" {
		return searchDefaultLimit, true
	}
	n, err := strconv.Atoi(v)
	return n, err == nil && n >= 1 && n <= searchMaxLimit
}

// searchQuery validates the request's parameters; it writes the 400 itself when one is bad.
func searchQuery(w http.ResponseWriter, r *http.Request) (CharacterSearch, bool) {
	p := r.URL.Query()
	q := CharacterSearch{
		Prefix:  searchTerm(p.Get("q")),
		Region:  strings.ToLower(p.Get("region")),
		Ruleset: strings.ToLower(p.Get("ruleset")),
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
	limit, ok := searchLimit(p.Get("limit"))
	if !ok {
		httpx.WriteError(w, r, http.StatusBadRequest, "bad_request", "limit must be a number from 1 to "+
			strconv.Itoa(searchMaxLimit), map[string]string{"limit": "1 to " + strconv.Itoa(searchMaxLimit)})
		return q, false
	}
	q.Limit = limit
	return q, true
}

func (s *Service) searchCharacters(w http.ResponseWriter, r *http.Request) {
	q, ok := searchQuery(w, r)
	if !ok {
		return
	}
	hits := []CharacterHit{}
	if utf8.RuneCountInString(q.Prefix) >= searchMinChars {
		var err error
		if hits, err = s.Store.SearchCharacters(r.Context(), q); err != nil {
			s.fail(w, r, "character_search", err, "could not search characters just now")
			return
		}
	}
	httpx.CachePublic(w, searchCacheSeconds*time.Second, liveBoardStale)
	httpx.WriteOK(w, r, http.StatusOK, map[string]any{"rows": hits})
}
