// api/internal/directory/harness_test.go
package directory

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/bnetapi"
	"github.com/jhunthrop/foreversixty/api/internal/db"
)

// testPool is a migrated database with the directory, guilds, characters and users emptied.
func testPool(t *testing.T) *pgxpool.Pool {
	t.Helper()
	url := os.Getenv("TEST_DATABASE_URL")
	if url == "" {
		t.Skip("TEST_DATABASE_URL not set; start docker-compose.test.yml")
	}
	if err := db.Migrate(url); err != nil {
		t.Fatal(err)
	}
	pool, err := db.Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(pool.Close)
	if _, err := pool.Exec(context.Background(),
		`truncate character_directory, directory_guilds, directory_name_misses, users, guilds, characters cascade`); err != nil {
		t.Fatal(err)
	}
	return pool
}

// blizzard is an httptest stand-in for the public Forever profile API, keyed by path (the
// namespace and locale are checked once, then dropped). An unstubbed path answers 404.
type blizzard struct {
	t      *testing.T
	server *httptest.Server
	mu     sync.Mutex
	routes map[string]any
	calls  []string
}

const testNamespace = "namespace=profile-forever-us"

func newBlizzard(t *testing.T) *blizzard {
	t.Helper()
	b := &blizzard{t: t, routes: map[string]any{}}
	b.server = httptest.NewServer(http.HandlerFunc(b.serve))
	t.Cleanup(b.server.Close)
	return b
}

func (b *blizzard) serve(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost && r.URL.Path == "/token" {
		_ = json.NewEncoder(w).Encode(map[string]any{"access_token": "app", "expires_in": 3600})
		return
	}
	if !strings.Contains(r.URL.RawQuery, testNamespace) {
		b.t.Errorf("request %s carries no Forever profile namespace", r.URL)
	}
	b.mu.Lock()
	b.calls = append(b.calls, r.URL.Path)
	body, ok := b.routes[r.URL.Path]
	b.mu.Unlock()
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(body)
}

func (b *blizzard) on(path string, body any) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.routes[path] = body
}

func (b *blizzard) called(path string) int {
	b.mu.Lock()
	defer b.mu.Unlock()
	n := 0
	for _, c := range b.calls {
		if c == path {
			n++
		}
	}
	return n
}

func (b *blizzard) service(pool *pgxpool.Pool) *Service {
	c := bnetapi.New(bnetapi.Config{
		TokenURL: b.server.URL + "/token", APIHost: func(string) string { return b.server.URL },
		Game: "forever", Regions: []string{"us"},
	})
	c.RetryBackoff = time.Millisecond
	return &Service{Pool: pool, Client: c, Regions: []string{"us"}, Budget: NewBudget(3_600_000), Workers: 2}
}

// guild stubs a guild summary and roster. Members are (id, "First Last", level, classID).
func (b *blizzard) guild(slug, name, ruleset string, id int64, members ...rosterMember) {
	b.on("/data/wow/guild/"+slug, map[string]any{
		"id": id, "name": name, "slug": slug, "faction": map[string]string{"type": "HORDE"},
		"ruleset": map[string]any{"id": 1, "name": ruleset}, "member_count": len(members),
	})
	rows := []map[string]any{}
	for _, m := range members {
		first, last, _ := strings.Cut(m.name, " ")
		rows = append(rows, map[string]any{
			"rank": m.rank,
			"character": map[string]any{
				"id": m.id, "name": first, "surname": last, "level": m.level,
				"playable_class": map[string]any{"id": m.classID}, "playable_race": map[string]any{"id": 2},
				"faction": map[string]string{"type": "HORDE"},
			},
		})
	}
	b.on("/data/wow/guild/"+slug+"/roster", map[string]any{
		"guild": map[string]any{"id": id, "name": name}, "members": rows,
	})
}

type rosterMember struct {
	id             int64
	name           string
	level, classID int
	rank           int
}

// publicCharacter stubs every public endpoint of one character.
func (b *blizzard) publicCharacter(id int64, name, ruleset string, guildID int64, guildSlug, guildName string) {
	first, last, _ := strings.Cut(name, " ")
	base := "/profile/wow/character/" + bnetapi.CharacterSlug(name)
	b.on(base+"/status", map[string]any{"id": id, "is_valid": true})
	summary := map[string]any{
		"id": id, "name": first, "surname": last, "level": 41, "faction": map[string]string{"type": "HORDE"},
		"character_class": map[string]any{"name": "Mage"}, "race": map[string]any{"name": "Undead"},
		"ruleset": map[string]any{"name": ruleset}, "last_login_timestamp": int64(1_763_000_000_000),
		"equipped_item_level": 38,
	}
	if guildSlug != "" {
		summary["guild"] = map[string]any{"id": guildID, "name": guildName, "slug": guildSlug}
	}
	b.on(base, summary)
	b.on(base+"/equipment", map[string]any{"equipped_items": []map[string]any{{
		"item": map[string]any{"id": 9001}, "slot": map[string]any{"type": "HEAD"},
		"quality": map[string]any{"type": "RARE"}, "name": "Test Circlet", "level": map[string]any{"value": 40},
	}}})
	b.on(base+"/specializations", map[string]any{"specialization_groups": []map[string]any{{
		"is_active": true, "specializations": []map[string]any{
			{"specialization_name": "Fire", "spent_points": 21}, {"specialization_name": "Frost", "spent_points": 5},
		},
	}}})
	b.on(base+"/professions", map[string]any{"primaries": []map[string]any{{
		"profession": map[string]any{"name": "Tailoring"}, "tiers": []map[string]any{{"skill_points": 180, "max_skill_points": 225}},
	}}})
	b.on(base+"/character-media", map[string]any{"assets": []map[string]any{
		{"key": "avatar", "value": "https://render.worldofwarcraft.com/us/avatar.jpg"},
	}})
}
