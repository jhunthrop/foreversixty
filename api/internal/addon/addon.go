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
	"github.com/jhunthrop/foreversixty/api/internal/events"
	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/sims"
	"github.com/jhunthrop/foreversixty/api/internal/synclog"
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
	// maxRankIndex is a WoW guild's highest real rank index (10 ranks, 0-9).
	maxRankIndex = 9
	// myExportsPerHour is POST /v1/me/exports's per-IP rate limit (spec
	// §4.5: "rate-limited like POST /v1/guilds/{id}/claim/contest" — the
	// same kind of httpx.RateLimitPer wrap, at guilds' contestPerHour
	// magnitude rather than claim's own 30-day account-level gate, since
	// a signed-in paste has no equivalent per-account store to lean on).
	myExportsPerHour = 20
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
type Store struct {
	Pool *pgxpool.Pool
	// Log is used only for warning about a malformed guild= section
	// during a sync (validateGuildName/rank-bound failures) - never for
	// routine operation. Defaults to slog.Default() when nil.
	Log *slog.Logger
}

func (s *Store) logger() *slog.Logger {
	if s.Log != nil {
		return s.Log
	}
	return slog.Default()
}

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
	dropped, err := s.dropAliasRows(ctx, userID, exports)
	if dropped > 0 {
		s.logger().Info("addon", "op", "drop_aliases", "user_id", userID, "dropped", dropped)
	}
	return err
}

// dropAliasRows removes the rows an older addon build left behind for a
// two-word character: it saved the first word only ("Booty" beside "Booty
// Twerk"), and the companion posted those records with the rest. For each
// export whose name has a space, the same user's addon_exports and
// characters rows on the same region and ruleset whose name slug is the
// slug of its first word go, when the row's class matches the export's (or
// its export carries no level section, which only the old copies lack) and
// the batch does not carry it itself. Returns how many characters went.
func (s *Store) dropAliasRows(ctx context.Context, userID int64, exports []Export) (int, error) {
	inBatch := make(map[string]bool, len(exports))
	for _, e := range exports {
		inBatch[character.Key(e.Region, character.RulesetFromRealm(e.Ruleset, ""), e.Name)] = true
	}
	dropped := 0
	for _, e := range exports {
		first, _, spaced := strings.Cut(strings.TrimSpace(e.Name), " ")
		if !spaced || first == "" {
			continue
		}
		ruleset := character.RulesetFromRealm(e.Ruleset, "")
		aliasKey := character.Key(e.Region, ruleset, first)
		if inBatch[aliasKey] {
			continue
		}
		var aliasExport *string
		var aliasClass *string
		err := s.Pool.QueryRow(ctx,
			`select c.class, x.export
			   from characters c
			   left join addon_exports x on x.character_key = c.key
			  where c.key = $1 and c.user_id = $2`, aliasKey, userID).Scan(&aliasClass, &aliasExport)
		if errors.Is(err, pgx.ErrNoRows) {
			continue
		}
		if err != nil {
			return dropped, fmt.Errorf("addon: read alias %s: %w", aliasKey, err)
		}
		class := ""
		if aliasExport != nil {
			class, _ = ParseFS1Class(*aliasExport)
		}
		if class == "" && aliasClass != nil {
			class = *aliasClass
		}
		want, _ := ParseFS1Class(e.Export)
		hasLevel := aliasExport != nil && strings.Contains("|"+*aliasExport, "|level=")
		if hasLevel && !strings.EqualFold(class, want) {
			continue
		}
		if err := s.deleteCharacter(ctx, userID, aliasKey); err != nil {
			return dropped, err
		}
		dropped++
	}
	return dropped, nil
}

// deleteCharacter removes one of userID's characters and its export in one
// transaction, clearing its guild row the way an unguilded export would.
func (s *Store) deleteCharacter(ctx context.Context, userID int64, key string) error {
	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("addon: drop alias %s: begin: %w", key, err)
	}
	var guildID *int64
	err = tx.QueryRow(ctx,
		`delete from guild_characters where character_key = $1 and user_id = $2 returning guild_id`,
		key, userID).Scan(&guildID)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		_ = tx.Rollback(ctx)
		return fmt.Errorf("addon: drop alias %s: guild: %w", key, err)
	}
	if _, err := tx.Exec(ctx, `delete from addon_exports where character_key = $1 and user_id = $2`, key, userID); err != nil {
		_ = tx.Rollback(ctx)
		return fmt.Errorf("addon: drop alias %s: export: %w", key, err)
	}
	if _, err := tx.Exec(ctx, `delete from characters where key = $1 and user_id = $2`, key, userID); err != nil {
		_ = tx.Rollback(ctx)
		return fmt.Errorf("addon: drop alias %s: character: %w", key, err)
	}
	if guildID != nil {
		if err := guilds.AfterGuildChange(ctx, tx, *guildID, userID); err != nil {
			_ = tx.Rollback(ctx)
			return err
		}
	}
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("addon: drop alias %s: commit: %w", key, err)
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
//
// seen_at moves on every accepted upload, because the inbox ack tells the addon when the site
// last received the character; updated_at moves only when the export text differs, so it still
// means "the export changed" (the ack carries it as export_changed_at).
func (s *Store) putOneExport(ctx context.Context, tx pgx.Tx, userID int64, key, region, ruleset string, e Export) error {
	tag, err := tx.Exec(ctx,
		`insert into addon_exports (character_key, user_id, region, ruleset, name, export, source, captured_at, updated_at, seen_at)
		 values ($1, $2, $3, $4, $5, $6, 'addon', now(), now(), now())
		 on conflict (character_key) do update set
		   user_id = excluded.user_id, export = excluded.export,
		   source = 'addon', captured_at = now(), seen_at = now(),
		   updated_at = case when addon_exports.export is distinct from excluded.export
		                     then now() else addon_exports.updated_at end
		 where addon_exports.user_id = excluded.user_id`,
		key, userID, region, ruleset, e.Name, e.Export)
	if err != nil {
		return fmt.Errorf("addon: store export %s: %w", key, err)
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("%w: %s", ErrCharacterClaimed, key)
	}
	if err := s.putOneCharacter(ctx, tx, userID, key, region, ruleset, e); err != nil {
		return err
	}
	if err := synclog.Record(ctx, tx, key, synclog.SourceAddon, synclog.OutcomeOK); err != nil {
		return err
	}
	return s.syncGuild(ctx, tx, userID, key, region, ruleset, e.Export)
}

// putOneCharacter upserts characters (spec §4.5): the same ownership
// guard as auth.Store.LinkCharacter and the Battle.net import's own
// write — never steal a key from another account. class is read from
// the export's head (ParseFS1Class); when the head carries none, the
// column is left as it was (coalesce), the same "silently degrade,
// never clobber good data with unknown" rule the Battle.net import's
// richer write follows for its own optional fields. This is what makes
// GET /v1/me list a companion (or signed-in paste) user's characters.
func (s *Store) putOneCharacter(ctx context.Context, tx pgx.Tx, userID int64, key, region, ruleset string, e Export) error {
	class, _ := ParseFS1Class(e.Export)
	tag, err := tx.Exec(ctx,
		`insert into characters (key, region, ruleset, name, class, user_id, source, refreshed_at)
		 values ($1, $2, $3, $4, nullif($5, ''), $6, 'export', now())
		 on conflict (key) do update set
		   name = excluded.name, class = coalesce(nullif($5, ''), characters.class),
		   user_id = excluded.user_id,
		   -- A row Battle.net imported stays 'bnet' so the nightly refresh
		   -- keeps re-checking its guild; an export only claims the
		   -- source when nothing stronger holds it.
		   source = case when characters.source = 'bnet' then 'bnet' else 'export' end,
		   refreshed_at = now()
		 where characters.user_id is null or characters.user_id = excluded.user_id`,
		key, region, ruleset, e.Name, class, userID)
	if err != nil {
		return fmt.Errorf("addon: store character %s: %w", key, err)
	}
	if tag.RowsAffected() == 0 {
		// characters and addon_exports are independently owned rows; the
		// addon_exports write above can succeed for userID while a
		// different account already owns this key's characters row (a
		// Battle.net import, or another device). Never steal it.
		return fmt.Errorf("%w: %s", ErrCharacterClaimed, key)
	}
	return nil
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
	if ok {
		var validName bool
		name, validName = validateGuildName(name)
		if !validName || rankIndex < 0 || rankIndex > maxRankIndex {
			s.logger().Warn("addon", "op", "guild_sync", "character_key", key,
				"reason", "invalid guild= section: bad name or rank index outside 0-9")
			ok = false
		}
	}
	if !ok {
		if prevGuildID == nil {
			return nil
		}
		if _, err := tx.Exec(ctx, `delete from guild_characters where character_key = $1`, key); err != nil {
			return fmt.Errorf("addon: clear guild for %s: %w", key, err)
		}
		return guilds.AfterGuildChange(ctx, tx, *prevGuildID, prevUserID)
	}

	guildID, officerMax, err := guilds.ResolveGuild(ctx, tx, region, ruleset, name)
	if err != nil {
		return err
	}
	rank := guilds.DeriveRank(rankIndex, officerMax)

	if prevGuildID != nil && *prevGuildID != guildID {
		// A transfer touches two guilds; lock both, in a fixed
		// ascending order, before any mutation on either - two
		// concurrent opposite-direction transfers between the same
		// pair of guilds would otherwise each lock their own "new"
		// guild first and then deadlock waiting for the other's "old"
		// guild (E, 2026-09-21 second security review response).
		if err := guilds.LockGuilds(ctx, tx, guildID, *prevGuildID); err != nil {
			return err
		}
		if _, err := tx.Exec(ctx,
			`delete from guild_characters where character_key = $1 and guild_id = $2`, key, *prevGuildID); err != nil {
			return fmt.Errorf("addon: clear previous guild for %s: %w", key, err)
		}
	}

	// Reverify: false — a re-sync of the same character in the same
	// guild keeps whatever verification it already earned (including a
	// 'bnet' row this export must not downgrade, spec §4.3). Only a
	// transfer starts a fresh, unverified row.
	if err := guilds.UpsertCharacterMembership(ctx, tx, guilds.MembershipRow{
		GuildID: guildID, CharacterKey: key, UserID: userID,
		RankIndex: rankIndex, Rank: rank, Source: "export", Reverify: false,
	}); err != nil {
		return fmt.Errorf("addon: sync guild membership for %s: %w", key, err)
	}

	if rank == "leader" {
		if err := guilds.AutoConfirmClaimIfPending(ctx, tx, guildID, userID); err != nil {
			return err
		}
	}
	if err := guilds.AfterGuildChange(ctx, tx, guildID, userID); err != nil {
		return err
	}
	if prevGuildID != nil && *prevGuildID != guildID {
		return guilds.AfterGuildChange(ctx, tx, *prevGuildID, prevUserID)
	}
	return nil
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
	// Accounts answers POST /v1/me/exports's response body. Nil is safe
	// (the response carries an empty characters list) for a test
	// harness that does not exercise it.
	Accounts CharacterReader
	// Guilds and Sims produce the Wave C typed inbox messages
	// (messages.go): guild state per character, and the account's
	// upgrade/weights messages from its saved sims. Either or both nil
	// is safe — the inbox still answers with builds and whatever
	// messages the ones that are wired can produce, never an error, the
	// same "degrade, don't fail the sync" rule Builds/Data already
	// follow for a build's code.
	Guilds GuildSource
	Sims   SimSource
	// SimRequests turns an outbox's sim_request into a server run.
	// sims.Service satisfies it. Nil answers every request
	// "unavailable"; the export it rode in with is still stored.
	SimRequests SimRequester
	// Events publishes character.linked and band.entered after a sync; nil publishes nothing.
	Events events.Publisher
	// PublicBaseURL is the site origin the notifications link to.
	PublicBaseURL string
	Log           *slog.Logger
}

// SimRequester is enough of sims.Service to start the run an in-game
// sim request asks for, for one of the caller's own characters.
type SimRequester interface {
	RequestFromAddon(ctx context.Context, userID int64, characterKey string, req sims.AddonSimRequest) (string, error)
}

// Mount registers the addon routes. The companion reads and writes with
// its device token; the site queues a build, and a signed-in paste
// stores its own exports, with a session.
func Mount(mux *http.ServeMux, s *Service, trustedProxyHops int) {
	mux.HandleFunc("POST /v1/addon/exports", auth.RequireDevice(s.putExports))
	mux.HandleFunc("GET /v1/addon/inbox", auth.RequireDevice(s.inbox))
	mux.HandleFunc("POST /v1/addon/inbox", auth.RequireSession(s.queueBuild))
	myExports := httpx.RateLimitPer(myExportsPerHour, time.Hour, trustedProxyHops)
	mux.Handle("POST /v1/me/exports", myExports(auth.RequireSession(s.putMyExports)))
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

// ExportsInput is the body of POST /v1/addon/exports: the per-character
// exports the companion has always sent, the addon's outbox
// (docs/contracts/addon-v2.md, contract 4), or both. Outbox stays raw
// until its format is read: an outbox of a format this API does not know
// is ignored whole, never half-read.
type ExportsInput struct {
	Characters []Export        `json:"characters"`
	Outbox     json.RawMessage `json:"outbox,omitempty"`
}

// validExports reports whether exports is a well-formed batch (between
// one and MaxExports entries, each with a name, an export of at most
// MaxExportLen, a valid region, and a name whose slug is a valid
// character key), writing the first violation to w. Shared by
// putExports and putMyExports so the two request bodies (device sync
// and the signed-in paste, spec §4.5) are validated identically.
func validExports(w http.ResponseWriter, r *http.Request, exports []Export) bool {
	if len(exports) == 0 || len(exports) > MaxExports {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "that is not a sync",
			map[string]string{"characters": "between one and 200 characters"})
		return false
	}
	for _, c := range exports {
		if c.Name == "" || len(c.Export) == 0 || len(c.Export) > MaxExportLen {
			httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "that export cannot be stored",
				map[string]string{"characters": "each needs a name and an export of at most 4096 characters"})
			return false
		}
		if !character.ValidRegion(c.Region) {
			httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "that character cannot be stored",
				map[string]string{"characters": "region must be one of us, eu, kr, tw, cn"})
			return false
		}
		// The name becomes the key's slug and, through the inbox, a
		// Lua string; the same rule ValidKey applies on the way back
		// out is applied on the way in.
		if !character.ValidSlug(character.Slug(c.Name)) {
			httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "that character cannot be stored",
				map[string]string{"characters": "name must be letters, digits, spaces and hyphens"})
			return false
		}
	}
	return true
}

// ExportsResult is the body POST /v1/addon/exports answers with.
type ExportsResult struct {
	// Stored counts the exports written: every character in the list,
	// plus the outbox's own export when it was stored for a character
	// the list did not already carry.
	Stored int           `json:"stored"`
	Outbox *OutboxResult `json:"outbox,omitempty"`
}

func (s *Service) putExports(w http.ResponseWriter, r *http.Request) {
	var in ExportsInput
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxBody)).Decode(&in); err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid",
			"body must be JSON with a characters list", nil)
		return
	}
	hasOutbox := len(in.Outbox) > 0 && string(in.Outbox) != "null"
	// An outbox may arrive alone; without one the list is the sync and
	// must hold at least one character, as it always has.
	if (len(in.Characters) > 0 || !hasOutbox) && !validExports(w, r, in.Characters) {
		return
	}
	var outbox *Outbox
	if hasOutbox {
		var ok bool
		if outbox, ok = readOutbox(w, r, in.Outbox); !ok {
			return
		}
	}
	userID := auth.ActorFrom(r.Context()).UserID
	out := ExportsResult{Stored: len(in.Characters)}
	notifier := s.beforeSync(r.Context(), userID, in.Characters)
	err := s.Store.PutExports(r.Context(), userID, in.Characters)
	if err == nil {
		s.afterSync(r.Context(), notifier, in.Characters)
	}
	if err == nil && hasOutbox {
		var key string
		out.Outbox, key, err = s.applyOutbox(r.Context(), userID, outbox)
		if key != "" && !listNames(in.Characters, key) {
			out.Stored++
		}
	}
	if err != nil {
		if errors.Is(err, ErrCharacterClaimed) {
			httpx.WriteError(w, r, http.StatusConflict, "conflict",
				"one of those characters is already synced from a different account",
				map[string]string{"characters": "already claimed by another account"})
			return
		}
		s.fail(w, r, "exports", err, "could not store those exports just now")
		return
	}
	httpx.WriteOK(w, r, http.StatusOK, out)
}

// listNames reports whether one of the list's exports is stored under key.
func listNames(exports []Export, key string) bool {
	for _, e := range exports {
		if character.Key(e.Region, character.RulesetFromRealm(e.Ruleset, ""), e.Name) == key {
			return true
		}
	}
	return false
}

// MyExportsInput is the body of POST /v1/me/exports (spec §4.5): the
// addon-less way to reach the same characters rows the companion writes.
type MyExportsInput struct {
	Exports []Export `json:"exports"`
}

// CharacterReader answers the /v1/me character objects for a set of
// keys. auth.Store satisfies it; kept narrow so this package need not
// depend on auth's whole Store surface and a test can stub it.
type CharacterReader interface {
	CharactersByKeys(ctx context.Context, keys []string) ([]auth.Character, error)
}

func (s *Service) putMyExports(w http.ResponseWriter, r *http.Request) {
	var in MyExportsInput
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxBody)).Decode(&in); err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid",
			"body must be JSON with an exports list", nil)
		return
	}
	if !validExports(w, r, in.Exports) {
		return
	}
	userID := auth.ActorFrom(r.Context()).UserID
	notifier := s.beforeSync(r.Context(), userID, in.Exports)
	if err := s.Store.PutExports(r.Context(), userID, in.Exports); err != nil {
		if errors.Is(err, ErrCharacterClaimed) {
			httpx.WriteError(w, r, http.StatusConflict, "conflict",
				"one of those characters is already synced from a different account",
				map[string]string{"exports": "already claimed by another account"})
			return
		}
		s.fail(w, r, "my_exports", err, "could not store those exports just now")
		return
	}
	s.afterSync(r.Context(), notifier, in.Exports)
	chars, err := s.writtenCharacters(r.Context(), in.Exports)
	if err != nil {
		s.fail(w, r, "my_exports", err, "could not store those exports just now")
		return
	}
	httpx.WriteOK(w, r, http.StatusOK, map[string]any{"characters": chars})
}

// writtenCharacters reads back the /v1/me character objects for exactly
// the keys putMyExports just wrote. A nil Accounts (a test harness that
// does not wire one) answers an empty list rather than failing the
// request — the write already succeeded.
func (s *Service) writtenCharacters(ctx context.Context, exports []Export) ([]auth.Character, error) {
	if s.Accounts == nil {
		return []auth.Character{}, nil
	}
	keys := make([]string, len(exports))
	for i, e := range exports {
		keys[i] = character.Key(e.Region, character.RulesetFromRealm(e.Ruleset, ""), e.Name)
	}
	return s.Accounts.CharactersByKeys(ctx, keys)
}

func (s *Service) inbox(w http.ResponseWriter, r *http.Request) {
	userID := auth.ActorFrom(r.Context()).UserID
	rows, err := s.Store.Inbox(r.Context(), userID)
	if err != nil {
		s.fail(w, r, "inbox", err, "could not read the inbox just now")
		return
	}
	entries, err := s.entries(r.Context(), rows)
	if err != nil {
		s.fail(w, r, "inbox", err, "could not read the inbox just now")
		return
	}
	// Contract 3's envelope (docs/contracts/addon-v2.md): builds stays as
	// it always was for one release, and each build is repeated as a build
	// message ahead of the typed messages.
	now := time.Now().UTC()
	messages := append(buildMessages(entries), s.messages(r.Context(), userID, now)...)
	if messages == nil {
		messages = []MessageEntry{}
	}
	httpx.WriteOK(w, r, http.StatusOK, InboxResponse{
		Format: InboxFormat, GeneratedAt: now.Format(time.RFC3339), Builds: entries, Messages: messages,
	})
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
			// Only a name the player gave. Describe's race-and-class
			// fallback is for the site's pages; left empty, the addon names
			// the build from its trees ("Beast Mastery build").
			e.Name = b.Title
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

// unsendable says why a saved build cannot be queued, or "" when it can.
// A build with no talent points would reach the addon as an empty code
// it has nothing to follow, and one whose code cannot be rendered would
// be dropped silently, so both are refused here rather than queued to
// vanish. A build id the store does not have is not refused: a deleted
// build is queued and carried without a code, as entries documents. With
// no build source or tree data wired in there is nothing to check.
func (s *Service) unsendable(ctx context.Context, id string) string {
	if s.Builds == nil || s.Data == nil {
		return ""
	}
	records, err := s.Builds.GetMany(ctx, []string{id})
	if err != nil {
		// A read failure is not the build's fault; queue it and let the
		// inbox read report on it.
		s.logger().Warn("addon", "op", "queue", "build", id, "err", err)
		return ""
	}
	b, ok := records[id]
	if !ok {
		return ""
	}
	if len(b.PointOrder) == 0 {
		return "that build has no talent points"
	}
	if _, err := builds.AddonCode(s.Data, b); err != nil {
		s.logger().Warn("addon", "op", "queue", "build", id, "err", err)
		return "that build cannot be sent to the game"
	}
	return ""
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
	if msg := s.unsendable(r.Context(), in.BuildID); msg != "" {
		httpx.WriteError(w, r, http.StatusUnprocessableEntity, "invalid", msg,
			map[string]string{"build_id": "a build with talent points the addon can read"})
		return
	}
	if err := s.Store.AddInbox(r.Context(), auth.ActorFrom(r.Context()).UserID,
		in.CharacterKey, in.BuildID); err != nil {
		s.fail(w, r, "inbox", err, "could not queue that build just now")
		return
	}
	httpx.WriteOK(w, r, http.StatusCreated, map[string]string{"build_id": in.BuildID})
}
