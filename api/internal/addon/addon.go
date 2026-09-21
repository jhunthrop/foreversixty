// Package addon is the two thin stores the in-game addon syncs through
// the companion: the character exports it writes after a logout, and
// the inbox of builds chosen on the site that it reads at startup.
package addon

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/builds"
	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

// ErrCharacterClaimed is returned by PutExports when a character_key
// already belongs to a different account. Forever merges original
// realms into one ruleset, so two real characters can collide on the
// same region/ruleset/name-slug; the row is not moved, the same way
// auth.LinkCharacter refuses to move a characters row it does not own.
var ErrCharacterClaimed = errors.New("addon: character already claimed by another account")

const (
	// maxBody is the ceiling on an export or inbox body.
	maxBody = 1 << 20
	// MaxExportLen bounds one character's export string.
	MaxExportLen = 4096
	// MaxExports is how many characters one sync may carry.
	MaxExports = 200
	// InboxLimit is how many builds the addon's inbox file holds. The
	// store also prunes to this many rows per user on every write, so
	// the table itself never grows past what a read can ever return.
	InboxLimit = 50
)

// Export is one character's addon export.
type Export struct {
	Name    string `json:"name"`
	Ruleset string `json:"ruleset"`
	Region  string `json:"region"`
	Export  string `json:"export"`
}

// InboxEntry is one build waiting to be read in game: everything the
// companion needs to write the inbox file without asking again.
type InboxEntry struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Character string `json:"character,omitempty"`
	Code      string `json:"code"`
}

// queued is one inbox row as it is stored.
type queued struct {
	BuildID      string
	CharacterKey string
	CreatedAt    time.Time
}

// Store is the addon's two tables.
type Store struct{ Pool *pgxpool.Pool }

// PutExports replaces a user's exports for the characters named, and syncs
// each character's guild membership (guild_characters, then the derived
// guild_members row) from its export's guild= section, all in one
// transaction per character so a crash between the export write and the
// guild sync never leaves one without the other.
//
// The ruleset is normalised rather than validated: the addon writes
// whatever the game client gave it, and the companion forwards it
// untouched, so an unrecognised value goes through the API's single
// realm-to-ruleset mapping like any other realm segment would.
func (s *Store) PutExports(ctx context.Context, userID int64, exports []Export) error {
	for _, e := range exports {
		ruleset := character.RulesetFromRealm(e.Ruleset, "")
		region := strings.ToLower(e.Region)
		key := character.Key(e.Region, ruleset, e.Name)

		tx, err := s.Pool.Begin(ctx)
		if err != nil {
			return fmt.Errorf("addon: store export %s: begin: %w", key, err)
		}
		if err := s.putOneExport(ctx, tx, userID, key, region, ruleset, e); err != nil {
			_ = tx.Rollback(ctx)
			return err
		}
		if err := tx.Commit(ctx); err != nil {
			return fmt.Errorf("addon: store export %s: commit: %w", key, err)
		}
	}
	return nil
}

// putOneExport writes one character's export row — the same claim-guard
// upsert PutExports always used, now scoped to tx — then syncs its guild
// membership from the export's guild= section (or clears it) in the same
// transaction.
//
// character_key is the table's only key, with no user_id in the conflict
// target, so a plain upsert would let any device silently reassign and
// overwrite a row it does not own the moment two real characters collide
// on the same region/ruleset/name-slug. The WHERE guard only claims a row
// that is unclaimed or already userID's own, and a claim that touches no
// row (RowsAffected() == 0) is refused as ErrCharacterClaimed rather than
// stealing it.
func (s *Store) putOneExport(ctx context.Context, tx pgx.Tx, userID int64, key, region, ruleset string, e Export) error {
	tag, err := tx.Exec(ctx,
		`insert into addon_exports (character_key, user_id, region, ruleset, name, export, updated_at)
		 values ($1, $2, $3, $4, $5, $6, now())
		 on conflict (character_key) do update set
		   user_id = excluded.user_id, export = excluded.export, updated_at = now()
		 where addon_exports.user_id = excluded.user_id`,
		key, userID, region, ruleset, e.Name, e.Export)
	if err != nil {
		return fmt.Errorf("addon: store export %s: %w", key, err)
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("%w: %s", ErrCharacterClaimed, key)
	}
	return s.syncGuild(ctx, tx, userID, key, region, ruleset, e.Export)
}

// syncGuild reads the export's guild= section (if any) and makes
// guild_characters agree with it: a named guild gets an upserted row (a
// transfer removes the old guild's row first, so the table-wide unique
// index on character_key is never violated), an unguilded export removes
// whatever row existed. Either way, every account the change touches gets
// RecomputeMembership and a lost-claim check, in the same transaction as
// the character write above — this is the one place the API ever parses
// an export's contents.
func (s *Store) syncGuild(ctx context.Context, tx pgx.Tx, userID int64, key, region, ruleset, export string) error {
	var prevGuildID *int64
	var prevUserID int64
	err := tx.QueryRow(ctx,
		`select guild_id, user_id from guild_characters where character_key = $1`, key).
		Scan(&prevGuildID, &prevUserID)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return fmt.Errorf("addon: read previous guild for %s: %w", key, err)
	}

	name, rankIndex, ok := ParseFS1Guild(export)
	if !ok {
		if prevGuildID == nil {
			return nil
		}
		if _, err := tx.Exec(ctx, `delete from guild_characters where character_key = $1`, key); err != nil {
			return fmt.Errorf("addon: clear guild for %s: %w", key, err)
		}
		return afterGuildChange(ctx, tx, *prevGuildID, prevUserID)
	}

	guildID, officerMax, err := resolveGuild(ctx, tx, region, ruleset, name)
	if err != nil {
		return err
	}
	rank := deriveRank(rankIndex, officerMax)

	if prevGuildID != nil && *prevGuildID != guildID {
		if _, err := tx.Exec(ctx,
			`delete from guild_characters where character_key = $1 and guild_id = $2`, key, *prevGuildID); err != nil {
			return fmt.Errorf("addon: clear previous guild for %s: %w", key, err)
		}
	}

	// verified_at is deliberately not in this SET list: a re-sync of the
	// same character in the same guild keeps whatever verification it
	// already earned. Only a transfer starts a fresh, unverified row.
	if _, err := tx.Exec(ctx,
		`insert into guild_characters (guild_id, character_key, user_id, rank_index, rank, source, refreshed_at)
		 values ($1, $2, $3, $4, $5, 'export', now())
		 on conflict (guild_id, character_key) do update set
		   user_id = excluded.user_id, rank_index = excluded.rank_index, rank = excluded.rank,
		   refreshed_at = now()`,
		guildID, key, userID, rankIndex, rank); err != nil {
		return fmt.Errorf("addon: sync guild membership for %s: %w", key, err)
	}

	if rank == "leader" {
		if err := guilds.AutoConfirmClaimIfPending(ctx, tx, guildID); err != nil {
			return err
		}
	}
	if err := afterGuildChange(ctx, tx, guildID, userID); err != nil {
		return err
	}
	if prevGuildID != nil && *prevGuildID != guildID {
		return afterGuildChange(ctx, tx, *prevGuildID, prevUserID)
	}
	return nil
}

// resolveGuild finds or creates the guild an export names, returning its
// id and its current officer-rank threshold.
func resolveGuild(ctx context.Context, tx pgx.Tx, region, ruleset, name string) (id int64, officerMax int, err error) {
	if _, err := tx.Exec(ctx,
		`insert into guilds (region, ruleset, name) values ($1, $2, $3) on conflict (region, ruleset, name) do nothing`,
		region, ruleset, name); err != nil {
		return 0, 0, fmt.Errorf("addon: create guild %s: %w", name, err)
	}
	err = tx.QueryRow(ctx,
		`select id, officer_max_rank_index from guilds where region = $1 and ruleset = $2 and name = $3`,
		region, ruleset, name).Scan(&id, &officerMax)
	if err != nil {
		return 0, 0, fmt.Errorf("addon: read guild %s: %w", name, err)
	}
	return id, officerMax, nil
}

// deriveRank turns a raw GetGuildInfo rank index into the label
// officer detection uses. Index 0 is always the guild master
// (server-authoritative, never configurable).
func deriveRank(rankIndex, officerMax int) string {
	switch {
	case rankIndex == 0:
		return "leader"
	case rankIndex <= officerMax:
		return "officer"
	default:
		return "member"
	}
}

// afterGuildChange runs RecomputeMembership and the lost-claim check for
// one account in one guild, the pair of calls every branch of syncGuild
// needs after it changes a guild_characters row.
func afterGuildChange(ctx context.Context, tx pgx.Tx, guildID, userID int64) error {
	if err := guilds.RecomputeMembership(ctx, tx, guildID, &userID); err != nil {
		return err
	}
	return guilds.ReleaseClaimIfLost(ctx, tx, guildID, userID)
}

// Exports lists a user's stored exports, newest first.
func (s *Store) Exports(ctx context.Context, userID int64) ([]Export, error) {
	rows, err := s.Pool.Query(ctx,
		`select name, ruleset, region, export from addon_exports
		 where user_id = $1 order by updated_at desc`, userID)
	if err != nil {
		return nil, fmt.Errorf("addon: list exports: %w", err)
	}
	defer rows.Close()
	out := []Export{}
	for rows.Next() {
		var e Export
		if err := rows.Scan(&e.Name, &e.Ruleset, &e.Region, &e.Export); err != nil {
			return nil, fmt.Errorf("addon: list exports: %w", err)
		}
		out = append(out, e)
	}
	return out, rows.Err()
}

// AddInbox queues a build for the addon to read in game, then prunes the
// user's inbox back to the newest InboxLimit rows.
//
// The write is reachable by any signed-in session for its own account
// only - userID always comes from the actor the session middleware
// resolved, never from the request body, so one user can never fill
// another's inbox. That still leaves a single account able to queue
// build ids all day; without a bound the table would grow without
// limit while Inbox only ever returns the newest InboxLimit anyway, so
// every write trims back to that many rows for the user.
func (s *Store) AddInbox(ctx context.Context, userID int64, characterKey, buildID string) error {
	if _, err := s.Pool.Exec(ctx,
		`insert into addon_inbox (user_id, character_key, build_id) values ($1, $2, $3)
		 on conflict (user_id, character_key, build_id) do update set created_at = now()`,
		userID, characterKey, buildID); err != nil {
		return fmt.Errorf("addon: queue build %s: %w", buildID, err)
	}
	if _, err := s.Pool.Exec(ctx,
		`delete from addon_inbox where user_id = $1 and id not in (
		   select id from addon_inbox where user_id = $1 order by created_at desc limit $2)`,
		userID, InboxLimit); err != nil {
		return fmt.Errorf("addon: prune inbox for %d: %w", userID, err)
	}
	return nil
}

// Inbox lists the rows waiting for a user, newest first. The service
// turns them into entries; the store does not know what a build is.
func (s *Store) Inbox(ctx context.Context, userID int64) ([]queued, error) {
	rows, err := s.Pool.Query(ctx,
		`select build_id, character_key, created_at from addon_inbox
		 where user_id = $1 order by created_at desc limit $2`, userID, InboxLimit)
	if err != nil {
		return nil, fmt.Errorf("addon: read inbox: %w", err)
	}
	defer rows.Close()
	out := []queued{}
	for rows.Next() {
		var q queued
		if err := rows.Scan(&q.BuildID, &q.CharacterKey, &q.CreatedAt); err != nil {
			return nil, fmt.Errorf("addon: read inbox: %w", err)
		}
		out = append(out, q)
	}
	return out, rows.Err()
}

// BuildSource reads the builds an inbox names. builds.Store satisfies
// it; the service holds it as an interface so the addon tests need no
// planner data of their own.
type BuildSource interface {
	GetMany(ctx context.Context, ids []string) (map[string]builds.Build, error)
}

// Service serves the addon routes.
type Service struct {
	Store *Store
	// Builds and Data render each queued build's addon code. With
	// either missing the inbox still answers, carrying the build ids
	// and no codes, because a companion that cannot write gear data is
	// better than one that cannot start.
	Builds BuildSource
	Data   *trees.Data
	Log    *slog.Logger
}

// Mount registers the addon routes. The companion reads and writes with
// its device token; the site queues a build with a session, which is
// the only way an inbox row is ever created.
func Mount(mux *http.ServeMux, s *Service) {
	mux.HandleFunc("POST /v1/addon/exports", auth.RequireDevice(s.putExports))
	mux.HandleFunc("GET /v1/addon/inbox", auth.RequireDevice(s.inbox))
	mux.HandleFunc("POST /v1/addon/inbox", auth.RequireSession(s.queueBuild))
}

func (s *Service) logger() *slog.Logger {
	if s.Log != nil {
		return s.Log
	}
	return slog.Default()
}

func (s *Service) fail(w http.ResponseWriter, r *http.Request, op string, err error, message string) {
	s.logger().Error("addon", "id", httpx.RequestIDFrom(r.Context()), "op", op, "err", err)
	httpx.WriteError(w, r, http.StatusInternalServerError, "internal", message, nil)
}

// ExportsInput is the body of POST /v1/addon/exports.
type ExportsInput struct {
	Characters []Export `json:"characters"`
}

func (s *Service) putExports(w http.ResponseWriter, r *http.Request) {
	var in ExportsInput
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxBody)).Decode(&in); err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid",
			"body must be JSON with a characters list", nil)
		return
	}
	if len(in.Characters) == 0 || len(in.Characters) > MaxExports {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "that is not a sync",
			map[string]string{"characters": "between one and 200 characters"})
		return
	}
	for _, c := range in.Characters {
		if c.Name == "" || len(c.Export) == 0 || len(c.Export) > MaxExportLen {
			httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "that export cannot be stored",
				map[string]string{"characters": "each needs a name and an export of at most 4096 characters"})
			return
		}
		if !character.ValidRegion(c.Region) {
			httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "that character cannot be stored",
				map[string]string{"characters": "region must be one of us, eu, kr, tw, cn"})
			return
		}
		// The name becomes the key's slug and, through the inbox, a
		// Lua string; the same rule ValidKey applies on the way back
		// out is applied on the way in.
		if !character.ValidSlug(character.Slug(c.Name)) {
			httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "that character cannot be stored",
				map[string]string{"characters": "name must be letters, digits, spaces and hyphens"})
			return
		}
	}
	if err := s.Store.PutExports(r.Context(), auth.ActorFrom(r.Context()).UserID, in.Characters); err != nil {
		if errors.Is(err, ErrCharacterClaimed) {
			httpx.WriteError(w, r, http.StatusConflict, "conflict",
				"one of those characters is already synced from a different account",
				map[string]string{"characters": "already claimed by another account"})
			return
		}
		s.fail(w, r, "exports", err, "could not store those exports just now")
		return
	}
	httpx.WriteOK(w, r, http.StatusOK, map[string]int{"stored": len(in.Characters)})
}

func (s *Service) inbox(w http.ResponseWriter, r *http.Request) {
	rows, err := s.Store.Inbox(r.Context(), auth.ActorFrom(r.Context()).UserID)
	if err != nil {
		s.fail(w, r, "inbox", err, "could not read the inbox just now")
		return
	}
	entries, err := s.entries(r.Context(), rows)
	if err != nil {
		s.fail(w, r, "inbox", err, "could not read the inbox just now")
		return
	}
	httpx.WriteOK(w, r, http.StatusOK, map[string]any{"builds": entries})
}

// entries turns stored rows into what the companion writes into the
// game: the build's id, its name, the character it was sent to, and the
// addon code. A build that has since been deleted, or one whose code
// cannot be rendered, is carried without a code rather than dropped -
// the companion writes what it has and the addon ignores the rest.
func (s *Service) entries(ctx context.Context, rows []queued) ([]InboxEntry, error) {
	out := make([]InboxEntry, 0, len(rows))
	var records map[string]builds.Build
	if s.Builds != nil {
		ids := make([]string, 0, len(rows))
		for _, q := range rows {
			ids = append(ids, q.BuildID)
		}
		var err error
		if records, err = s.Builds.GetMany(ctx, ids); err != nil {
			return nil, err
		}
	}
	for _, q := range rows {
		e := InboxEntry{ID: q.BuildID, Character: q.CharacterKey}
		if b, ok := records[q.BuildID]; ok {
			e.Name = builds.Describe(s.Data, b).Title
			if code, err := builds.AddonCode(s.Data, b); err == nil {
				e.Code = code
			} else {
				s.logger().Warn("addon", "op", "inbox", "build", q.BuildID, "err", err)
			}
		}
		out = append(out, e)
	}
	return out, nil
}

// QueueInput is the body of POST /v1/addon/inbox.
type QueueInput struct {
	BuildID      string `json:"build_id"`
	CharacterKey string `json:"character_key"`
}

// queueBuild is the site's half of the inbox: the contract names only
// the device's read, but a store with no writer would always be empty,
// and "builds chosen on the site appear in-game" is what the inbox is
// for. The build lands only in the caller's own inbox: userID comes
// from the session actor, never from the request, so this can never
// write into an account the caller does not already own.
//
// Both fields are checked for shape before they are stored, because
// GET /v1/addon/inbox hands them straight back and the companion
// writes them into ForeverSixtyInbox.lua - a Lua source file the addon
// loads. Escaping that file is the companion's job and stays its job;
// this route's part of the bargain is that what it stores is a build
// id and a character key and can be nothing else. See the route's
// description in openapi.yaml.
func (s *Service) queueBuild(w http.ResponseWriter, r *http.Request) {
	var in QueueInput
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxBody)).Decode(&in); err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "body must be JSON with a build_id", nil)
		return
	}
	if !builds.ValidID(in.BuildID) {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "that is not a build",
			map[string]string{"build_id": "the id of a saved build"})
		return
	}
	// An empty key is the whole point of the column's default: a build
	// can be sent to the account rather than to one character.
	if in.CharacterKey != "" && !character.ValidKey(in.CharacterKey) {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "that is not a character",
			map[string]string{"character_key": "<region>/<ruleset>/<name-slug>, or omitted"})
		return
	}
	if err := s.Store.AddInbox(r.Context(), auth.ActorFrom(r.Context()).UserID,
		in.CharacterKey, in.BuildID); err != nil {
		s.fail(w, r, "inbox", err, "could not queue that build just now")
		return
	}
	httpx.WriteOK(w, r, http.StatusCreated, map[string]string{"build_id": in.BuildID})
}
