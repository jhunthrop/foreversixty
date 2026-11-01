package rankings

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// WatchLimit is how many characters one account may watch (spec W3-C7).
const WatchLimit = 50

// ErrWatchLimit is returned when a new watch would pass WatchLimit.
var ErrWatchLimit = errors.New("rankings: watch limit reached")

// WatchedCharacter is one row of GET /v1/me/watching: who the character is, as far as the
// site knows. class and level are null when no source knows them.
type WatchedCharacter struct {
	Key     string    `json:"key"`
	Region  string    `json:"region"`
	Ruleset string    `json:"ruleset"`
	Name    string    `json:"name"`
	Class   *string   `json:"class"`
	Level   *int      `json:"level"`
	AddedAt time.Time `json:"added_at"`
}

// watchedQuery names each watched key from its characters row, else its newest ranking row,
// and leaves out a character whose account is anonymous. $1 is the user; $2, when the query
// is narrowed, one key.
const watchedQuery = `
	select w.character_key, coalesce(c.name, m.player_name, split_part(w.character_key, '/', 3)),
	       coalesce(c.class, m.class), w.added_at
	from watched_characters w
	left join characters c on c.key = w.character_key
	left join users u on u.id = c.user_id
	left join lateral (
	  select fm.player_name, fm.class from fight_metrics fm
	  where fm.player_key = w.character_key order by fm.fought_at desc limit 1
	) m on true
	where w.user_id = $1 and not coalesce(u.anonymize, false)`

// Watching lists the characters userID watches, newest first.
func (s *Store) Watching(ctx context.Context, userID int64) ([]WatchedCharacter, error) {
	return s.readWatched(ctx, watchedQuery+` order by w.added_at desc, w.character_key`, userID)
}

func (s *Store) readWatched(ctx context.Context, query string, args ...any) ([]WatchedCharacter, error) {
	rows, err := s.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("rankings: read watching: %w", err)
	}
	defer rows.Close()
	out := []WatchedCharacter{}
	for rows.Next() {
		var w WatchedCharacter
		if err := rows.Scan(&w.Key, &w.Name, &w.Class, &w.AddedAt); err != nil {
			return nil, fmt.Errorf("rankings: read watching: %w", err)
		}
		parts := strings.SplitN(w.Key, "/", 3)
		w.Region, w.Ruleset = parts[0], parts[1]
		out = append(out, w)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rankings: read watching: %w", err)
	}
	return out, s.addLevels(ctx, out)
}

// addLevels fills each row's level by the character page's own rule (CharacterLevels).
func (s *Store) addLevels(ctx context.Context, rows []WatchedCharacter) error {
	if len(rows) == 0 {
		return nil
	}
	keys := make([]string, len(rows))
	for i, r := range rows {
		keys[i] = r.Key
	}
	levels, err := CharacterLevels(ctx, s.Pool, keys)
	if err != nil {
		return err
	}
	for i := range rows {
		rows[i].Level = levels[rows[i].Key]
	}
	return nil
}

// Watchable reports whether key names a character the site shows: one with a characters row
// or a ranking row, whose account is not anonymous.
func (s *Store) Watchable(ctx context.Context, key string) (bool, error) {
	anon, err := s.anonymized(ctx, key)
	if err != nil || anon {
		return false, err
	}
	var ok bool
	err = s.Pool.QueryRow(ctx, `select exists (select 1 from characters where key = $1)
		or exists (select 1 from fight_metrics where player_key = $1)`, key).Scan(&ok)
	if err != nil {
		return false, fmt.Errorf("rankings: watchable %s: %w", key, err)
	}
	return ok, nil
}

// Watch adds key to userID's list and answers the row. Watching again is a no-op; a new key
// past WatchLimit is ErrWatchLimit. The count and the insert run in one transaction locked on
// the account, so two quick watches cannot both pass the limit.
func (s *Store) Watch(ctx context.Context, userID int64, key string) (WatchedCharacter, error) {
	err := pgx.BeginFunc(ctx, s.Pool, func(tx pgx.Tx) error {
		if _, err := tx.Exec(ctx, `select 1 from users where id = $1 for update`, userID); err != nil {
			return err
		}
		var n int
		var already bool
		if err := tx.QueryRow(ctx, `select count(*), coalesce(bool_or(character_key = $2), false)
			from watched_characters where user_id = $1`, userID, key).Scan(&n, &already); err != nil {
			return err
		}
		if already {
			return nil
		}
		if n >= WatchLimit {
			return ErrWatchLimit
		}
		_, err := tx.Exec(ctx, `insert into watched_characters (user_id, character_key) values ($1, $2)
			on conflict do nothing`, userID, key)
		return err
	})
	if errors.Is(err, ErrWatchLimit) {
		return WatchedCharacter{}, err
	}
	if err != nil {
		return WatchedCharacter{}, fmt.Errorf("rankings: watch %s: %w", key, err)
	}
	rows, err := s.readWatched(ctx, watchedQuery+` and w.character_key = $2`, userID, key)
	if err != nil {
		return WatchedCharacter{}, err
	}
	if len(rows) == 0 {
		return WatchedCharacter{}, fmt.Errorf("rankings: watch %s: the row was not read back", key)
	}
	return rows[0], nil
}

// Unwatch removes key from userID's list; a key that was not there is no error.
func (s *Store) Unwatch(ctx context.Context, userID int64, key string) error {
	if _, err := s.Pool.Exec(ctx,
		`delete from watched_characters where user_id = $1 and character_key = $2`, userID, key); err != nil {
		return fmt.Errorf("rankings: unwatch %s: %w", key, err)
	}
	return nil
}

// watchKey reads the path's character as a player key; false when the region or ruleset is
// not one the site knows.
func watchKey(r *http.Request) (string, bool) {
	region, ruleset := strings.ToLower(r.PathValue("region")), strings.ToLower(r.PathValue("ruleset"))
	name := character.Slug(r.PathValue("name"))
	if !character.ValidRegion(region) || !character.ValidRuleset(ruleset) || name == "" {
		return "", false
	}
	return region + "/" + ruleset + "/" + name, true
}

func (s *Service) watching(w http.ResponseWriter, r *http.Request) {
	rows, err := s.Store.Watching(r.Context(), auth.ActorFrom(r.Context()).UserID)
	if err != nil {
		s.fail(w, r, "watching", err, "could not read who you watch just now")
		return
	}
	httpx.CachePrivate(w)
	httpx.WriteOK(w, r, http.StatusOK, map[string]any{"rows": rows})
}

func (s *Service) watch(w http.ResponseWriter, r *http.Request) {
	key, ok := watchKey(r)
	if ok {
		var err error
		if ok, err = s.Store.Watchable(r.Context(), key); err != nil {
			s.fail(w, r, "watch", err, "could not watch that character just now")
			return
		}
	}
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such character", nil)
		return
	}
	row, err := s.Store.Watch(r.Context(), auth.ActorFrom(r.Context()).UserID, key)
	if errors.Is(err, ErrWatchLimit) {
		httpx.WriteError(w, r, http.StatusConflict, "limit", fmt.Sprintf("you can watch up to %d characters", WatchLimit),
			map[string]string{"character": fmt.Sprintf("stop watching one of your %d first", WatchLimit)})
		return
	}
	if err != nil {
		s.fail(w, r, "watch", err, "could not watch that character just now")
		return
	}
	httpx.CachePrivate(w)
	httpx.WriteOK(w, r, http.StatusOK, row)
}

func (s *Service) unwatch(w http.ResponseWriter, r *http.Request) {
	if key, ok := watchKey(r); ok {
		if err := s.Store.Unwatch(r.Context(), auth.ActorFrom(r.Context()).UserID, key); err != nil {
			s.fail(w, r, "unwatch", err, "could not stop watching just now")
			return
		}
	}
	w.WriteHeader(http.StatusNoContent)
}
