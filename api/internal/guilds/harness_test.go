// api/internal/guilds/harness_test.go
package guilds

import (
	"context"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/db"
)

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
	if _, err := pool.Exec(context.Background(), `truncate users, guilds, reports, fight_metrics cascade`); err != nil {
		t.Fatal(err)
	}
	return pool
}

// seedUser inserts a bare account and returns its id.
func seedUser(t *testing.T, pool *pgxpool.Pool, email string) int64 {
	t.Helper()
	var id int64
	if err := pool.QueryRow(context.Background(),
		`insert into users (email) values ($1) returning id`, email).Scan(&id); err != nil {
		t.Fatal(err)
	}
	return id
}

// seedGuild inserts a guild and returns its id.
func seedGuild(t *testing.T, pool *pgxpool.Pool, name string) int64 {
	t.Helper()
	var id int64
	if err := pool.QueryRow(context.Background(),
		`insert into guilds (region, ruleset, name) values ('us', 'hardcore', $1) returning id`, name).Scan(&id); err != nil {
		t.Fatal(err)
	}
	return id
}

// seedCharacter inserts a guild_characters row directly.
func seedCharacter(t *testing.T, pool *pgxpool.Pool, guildID, userID int64, key, rank string, verified bool) {
	t.Helper()
	if _, err := pool.Exec(context.Background(), `
		insert into guild_characters (guild_id, character_key, user_id, rank, verified_at)
		values ($1, $2, $3, $4, case when $5 then now() else null end)`,
		guildID, key, userID, rank, verified); err != nil {
		t.Fatal(err)
	}
}
