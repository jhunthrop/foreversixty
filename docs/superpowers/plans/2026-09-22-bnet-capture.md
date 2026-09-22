# Lane bnet-capture Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Fix three production defects in the Battle.net realm/import path (A1-A3), capture
every raw Blizzard payload the site is allowed to keep (B), and amend `/v1/me`'s shape with
`race`, `gender`, `item_level` (C).

**Architecture:** `bnetapi.Client.Realms` moves from serial per-realm detail fetches to a
bounded 8-worker pool and refuses to cache a partial result; `Client.WarmRealms` primes the
cache at service start. `bnetapi.Client.Character` gains a raw-body return alongside its
typed struct; a new `Client.Equipment` fetches the equipment endpoint raw. `bnetimport`
writes the raw bodies (size-capped at 256 KiB) into the new `bnet_account`/`bnet_profile`/
`bnet_equipment`/`bnet_captured_at` columns, logs `profile_unavailable` on a 404 character
profile and counts it, and rekeys a character whose Blizzard id already exists under a
different key for the same account. `auth.Character` and its SQL gain `race`, `gender`,
`item_level`.

**Tech Stack:** Go 1.25, pgx/v5, `httptest` fixtures, golang-migrate-style up/down SQL.

**Spec:** `docs/superpowers/specs/2026-09-22-battlenet-character-import-design.md`, amended
by `.superpowers/bnet-capture-brief.md` (the brief wins where they differ).

## Global Constraints

- Never persist a Battle.net user access token anywhere; never log a token, code, or state.
- Namespace/region stay configuration (`BNET_PROFILE_GAME`, `BNET_REGIONS`), never literals.
- `characters` is the single list of an account's characters; every writer keeps it so.
- Raw JSON capture: refuse to store a body over 256 KiB, log `op=capture_too_large`, and
  leave the previous value in place (a `coalesce($n, column)` write, not a NULL overwrite).
- Migration numbered `0025` (`0024` is main's highest at this lane's start).
- Go from `api/`, `GOWORK=off`; `go vet ./... && go test ./...` for touched packages before
  every commit; `TEST_DATABASE_URL=postgres://forever:forever@localhost:5434/forever_test?sslmode=disable`.
- File ownership: this lane owns `api/**` only; nothing in `web/**` changes.
- Functions <50 lines, files <800 lines, table-driven tests, wrapped errors, no magic
  numbers without a named constant and a comment on why.

---

## Task 1: Migration `0025_bnet_capture`

**Files:**
- Create: `api/internal/db/migrations/0025_bnet_capture.up.sql`
- Create: `api/internal/db/migrations/0025_bnet_capture.down.sql`
- Test: `api/internal/db/migration_order_test.go` (read-only check; no edit needed — `25 >
  18` already passes `TestMigrationNumbersDoNotRegressBelowMergeBase` by construction)

**Interfaces:**
- Produces: columns `characters.race text`, `characters.gender text check (gender in
  ('male','female'))`, `characters.average_item_level int`, `characters.equipped_item_level
  int`, `characters.last_login_at timestamptz`, `characters.bnet_account jsonb`,
  `characters.bnet_profile jsonb`, `characters.bnet_equipment jsonb`,
  `characters.bnet_captured_at timestamptz`.

- [ ] Step 1: Write `0025_bnet_capture.up.sql` with exactly the `alter table characters add
  column if not exists ...` block from brief §B (nine columns above).
- [ ] Step 2: Write `0025_bnet_capture.down.sql` dropping the same nine columns in reverse
  order, each `drop column if exists`.
- [ ] Step 3: Run `cd api && GOWORK=off go test ./internal/db/... -run Migration` — expect
  PASS (the order test never inspects column contents, only file numbering).
- [ ] Step 4: Run `cd api && GOWORK=off TEST_DATABASE_URL=postgres://forever:forever@localhost:5434/forever_test?sslmode=disable go test ./internal/bnetimport/... -run TestNothing` —
  this forces `db.Migrate` to run against the real test database and fail loudly if the SQL
  has a syntax error (any bnetimport test triggers `testPool`, which calls `db.Migrate`).
- [ ] Step 5: Commit: `feat(db): migration 0025 adds Battle.net capture columns`.

## Task 2: `bnetapi.Client.Realms` — concurrent, never-partial-cache, `WarmRealms`

**Files:**
- Modify: `api/internal/bnetapi/realms.go`
- Modify: `api/internal/bnetapi/realms_test.go`

**Interfaces:**
- Produces: `func (c *Client) Realms(ctx, region string) ([]Realm, error)` (same signature,
  new behavior — never returns a list where any detail failed); `func (c *Client)
  WarmRealms(ctx context.Context, regions []string)`.

- [ ] Step 1: Add `const realmWorkerCount = 8` and `const warmRealmsTimeout = 30 *
  time.Second` to `realms.go`; add `"fmt"`, `"sync"`, `"time"` imports.
- [ ] Step 2: Replace the serial detail loop inside `Realms` with a call to a new
  `fetchRealmDetails(ctx, region string, realms []Realm) error` that runs `realmWorkerCount`
  goroutines pulling `{index, slug}` jobs off a channel, writes `Type`/`Category` into
  `realms[job.index]` (safe: each worker owns a disjoint index), logs the existing
  `realm_detail`/`realm_type` WARN lines per failure/unknown-type, and — after every worker
  drains — returns the **first** collected error (or nil). `Realms` returns that error
  immediately (never caching) instead of continuing with an empty `Type`.
- [ ] Step 3: Add `func (c *Client) WarmRealms(ctx context.Context, regions []string)`:
  derives its own `context.WithTimeout(ctx, warmRealmsTimeout)`, calls `c.Realms` per region,
  logs a WARN `op=warm_realms` on error, never returns anything (fire-and-forget).
- [ ] Step 4: Add `TestRealmsNeverCachesAPartialListOnDetailFailure` (index answers 2 realms,
  one detail 200, one 500 — assert `Realms` returns a non-nil error and
  `c.realmCache["us"]` has no entry afterward).
- [ ] Step 5: Add `TestRealmsFetchesDetailsConcurrently` (16 realms, each detail handler
  sleeps 20ms and records concurrent in-flight count under a mutex; assert total elapsed
  time is well under the serial 320ms and max in-flight > 1).
- [ ] Step 6: Run `cd api && GOWORK=off go test ./internal/bnetapi/... -run Realms -v` —
  expect all `TestRealms*` PASS.
- [ ] Step 7: Commit: `fix(bnetapi): fetch realm details concurrently, never cache a partial list`.

## Task 3: main.go — service logger on the Blizzard client, warm the cache at start

**Files:**
- Modify: `api/cmd/api/main.go`

**Interfaces:**
- Consumes: `bnetapi.Config.Log *slog.Logger` (already exists); `Client.WarmRealms` (Task 2).
- Produces: `func bnetClient(cfg config.Config, log *slog.Logger) *bnetapi.Client` (signature
  change — both call sites updated).

- [ ] Step 1: Change `bnetClient(cfg config.Config)` to `bnetClient(cfg config.Config, log
  *slog.Logger) *bnetapi.Client`, adding `Log: log` to the `bnetapi.Config{}` literal.
- [ ] Step 2: Update `runBnetRefresh`'s call site to `bnetClient(cfg, log)`.
- [ ] Step 3: In `serve`, replace the inline `bnetimport.Service{..., Client: bnetClient(cfg),
  ...}` with a local `client := bnetClient(cfg, log)` reused both for `accounts.Importer` and
  a `go client.WarmRealms(context.Background(), cfg.BnetRegions)` call, inside the existing
  `if cfg.BattleNetConfigured() { ... }` block.
- [ ] Step 4: Run `cd api && GOWORK=off go build ./...` — expect success (this is a wiring
  change with no new test; `go vet` catches a signature mismatch).
- [ ] Step 5: Commit: `fix(api): the Blizzard client logs structured JSON and warms its realm cache at start`.

## Task 4: `bnetapi` — raw bodies for `Character`, new `Equipment`, race/gender on `AccountCharacters`

**Files:**
- Modify: `api/internal/bnetapi/client.go` (extract `getBytes`)
- Modify: `api/internal/bnetapi/profile.go`
- Modify: `api/internal/bnetapi/profile_test.go`

**Interfaces:**
- Produces: `func (c *Client) Character(ctx, region, realmSlug, name string) (CharacterProfile,
  json.RawMessage, error)` (signature change from two to three return values — every caller
  in `bnetimport` updates in Task 5); `func (c *Client) Equipment(ctx, region, realmSlug, name
  string) (json.RawMessage, error)`; `AccountCharacter.RaceName, GenderType string` and
  `AccountCharacter.Raw json.RawMessage`; `CharacterProfile.AverageItemLevel,
  EquippedItemLevel *int`.

- [ ] Step 1: In `client.go`, extract `getBytes(ctx context.Context, op, rawURL, token
  string) ([]byte, error)` out of `getJSONWithToken` (the fetch-and-read half, unchanged
  logic), and have `getJSONWithToken` call it then `json.Unmarshal`. `getJSON` is unchanged
  (still resolves the app token then calls `getJSONWithToken`).
- [ ] Step 2: In `profile.go`, change `accountCharactersResponse.WowAccounts[].Characters`
  from a typed slice to `[]json.RawMessage`; add a private `accountCharacterFields` struct
  (name, id, realm{slug,name}, playable_class{name}, playable_race{name}, gender{type},
  faction{type}, level) that `AccountCharacters` unmarshals each raw entry into, building
  `AccountCharacter{..., RaceName: fields.PlayableRace.Name, GenderType:
  strings.ToLower(fields.Gender.Type), Raw: raw}`.
- [ ] Step 3: Add `AverageItemLevel, EquippedItemLevel *int` (json tags
  `average_item_level`, `equipped_item_level`) to `characterProfileResponse` and
  `CharacterProfile`.
- [ ] Step 4: Rewrite `Character` to fetch the app token, call `getBytes`, `json.Unmarshal`
  the body itself into `characterProfileResponse`, build `CharacterProfile` (now including
  the two item-level fields), and return `(profile, json.RawMessage(body), nil)`; every
  error path returns `(CharacterProfile{}, nil, err)`.
- [ ] Step 5: Add `func (c *Client) Equipment(ctx context.Context, region, realmSlug, name
  string) (json.RawMessage, error)`: app token, `GET
  /profile/wow/character/{realm}/{lowercase name}/equipment?namespace=...`, `getBytes`,
  return the raw body.
- [ ] Step 6: Update `TestAccountCharactersUsesTheUsersOwnToken`'s fixture body to add
  `"playable_race":{"name":"Dwarf"},"gender":{"type":"MALE"}`; assert
  `chars[0].RaceName == "Dwarf"`, `chars[0].GenderType == "male"`, and `len(chars[0].Raw) >
  0`.
- [ ] Step 7: Update `TestCharacterReadsGuildedAndUnguilded` for the new three-value return;
  add `"average_item_level":55,"equipped_item_level":54` to the guilded fixture body and
  assert `*guilded.AverageItemLevel == 55`, `*guilded.EquippedItemLevel == 54`, and
  `len(rawGuilded) > 0` and `strings.Contains(string(rawGuilded), "Iron Vanguard")`.
- [ ] Step 8: Add `TestEquipmentReturnsTheRawBody`: fixture answers
  `/profile/wow/character/whitemane/thoradin/equipment?namespace=profile-classic1x-us` with
  a small `{"equipped_items":[...]}` body; assert the returned `json.RawMessage` round-trips
  byte-for-byte (`json.Equal` via `bytes.Equal` after re-marshal, or a direct field check).
- [ ] Step 9: Run `cd api && GOWORK=off go test ./internal/bnetapi/... -v` — expect all PASS
  (this task does not yet update `bnetimport`, which will fail to build until Task 5 — run
  `go vet ./internal/bnetapi/...` here, not `./...`).
- [ ] Step 10: Commit: `feat(bnetapi): Character returns its raw body, add Equipment, capture race/gender on AccountCharacters`.

## Task 5: `bnetimport` — write the capture columns (B), `profile_unavailable` (A2)

**Files:**
- Modify: `api/internal/bnetimport/character.go`
- Modify: `api/internal/bnetimport/import.go`
- Modify: `api/internal/bnetimport/refresh.go`
- Modify: `api/internal/auth/importer.go`
- Modify: `api/internal/auth/handler.go`
- Modify: `api/internal/bnetimport/import_test.go`
- Modify: `api/internal/bnetimport/refresh_test.go`
- Create: `api/internal/bnetimport/capture.go`
- Create: `api/internal/bnetimport/capture_test.go`

**Interfaces:**
- Consumes: `bnetapi.Client.Character` returning `(CharacterProfile, json.RawMessage,
  error)` (Task 4); `bnetapi.Client.Equipment` (Task 4).
- Produces: `func (s *Service) capCapture(field, key string, raw json.RawMessage)
  json.RawMessage`; `func (s *Service) captureAccountRaw(ch bnetapi.AccountCharacter, key
  string) (raw json.RawMessage, race, gender string)`; `func (s *Service) captureProfile(ctx,
  tx pgx.Tx, key string, profile bnetapi.CharacterProfile, raw json.RawMessage) error`;
  `func (s *Service) captureEquipment(ctx, tx pgx.Tx, key, region, realmSlug, name string)
  error`; `syncCharacterGuild(...) (guildWritten, unavailable bool, err error)` (signature
  gains a return value); `auth.ImportSummary.Unavailable int`.

- [ ] Step 1: Create `capture.go` with `const maxCaptureBytes = 256 * 1024` (spec §B: "refuse
  to store a body over 256 KiB") and:

```go
// capCapture returns raw when it is small enough to store verbatim, or
// nil (after logging op=capture_too_large) when it exceeds
// maxCaptureBytes — spec §B. The caller writes nil through a
// coalesce($n, column) UPDATE so a too-large body never blanks out a
// previously captured one.
func (s *Service) capCapture(field, key string, raw json.RawMessage) json.RawMessage {
	if len(raw) <= maxCaptureBytes {
		return raw
	}
	s.logger().Warn("bnetimport", "op", "capture_too_large", "field", field, "key", key, "bytes", len(raw))
	return nil
}

// captureProfile stores a character profile fetch's raw body and the
// three typed fields it carries (spec §B), preserving whatever was
// there before when raw was capped or a field is absent.
func (s *Service) captureProfile(ctx context.Context, tx pgx.Tx, key string, profile bnetapi.CharacterProfile, raw json.RawMessage) error {
	raw = s.capCapture("bnet_profile", key, raw)
	var lastLogin *time.Time
	if profile.LastLoginTimestamp > 0 {
		t := time.UnixMilli(profile.LastLoginTimestamp)
		lastLogin = &t
	}
	if _, err := tx.Exec(ctx,
		`update characters set
		   bnet_profile = coalesce($2, bnet_profile),
		   average_item_level = coalesce($3, average_item_level),
		   equipped_item_level = coalesce($4, equipped_item_level),
		   last_login_at = coalesce($5, last_login_at),
		   bnet_captured_at = now()
		 where key = $1`,
		key, rawOrNil(raw), profile.AverageItemLevel, profile.EquippedItemLevel, lastLogin); err != nil {
		return fmt.Errorf("bnetimport: capture profile %s: %w", key, err)
	}
	return nil
}

// captureEquipment fetches and stores a character's equipment snapshot
// verbatim (spec §B — not parsed here). A private or missing equipment
// page is left exactly as it was, the same rule syncCharacterGuild
// already applies to a private character profile.
func (s *Service) captureEquipment(ctx context.Context, tx pgx.Tx, key, region, realmSlug, name string) error {
	raw, err := s.Client.Equipment(ctx, region, realmSlug, name)
	if err != nil {
		if errors.Is(err, bnetapi.ErrNotFound) || errors.Is(err, bnetapi.ErrForbidden) {
			return nil
		}
		return fmt.Errorf("bnetimport: equipment %s: %w", key, err)
	}
	raw = s.capCapture("bnet_equipment", key, raw)
	if _, err := tx.Exec(ctx,
		`update characters set bnet_equipment = coalesce($2, bnet_equipment), bnet_captured_at = now() where key = $1`,
		key, rawOrNil(raw)); err != nil {
		return fmt.Errorf("bnetimport: capture equipment %s: %w", key, err)
	}
	return nil
}

// rawOrNil turns an empty json.RawMessage into a real nil interface value,
// so pgx sends SQL NULL (and the caller's coalesce(...) keeps whatever
// was already stored) instead of an empty-but-non-nil []byte.
func rawOrNil(raw json.RawMessage) any {
	if raw == nil {
		return nil
	}
	return raw
}
```

  (Imports: `context`, `encoding/json`, `errors`, `fmt`, `time`, `github.com/jackc/pgx/v5`,
  `.../bnetapi`.)

- [ ] Step 2: In `character.go`, change the `characters` upsert inside `importOneCharacter`
  to add `race`, `gender`, `bnet_account`, `bnet_captured_at`:

```go
	race := ch.RaceName
	gender := strings.ToLower(ch.GenderType)
	accountRaw := s.capCapture("bnet_account", key, ch.Raw)
	tag, err := tx.Exec(ctx,
		`insert into characters (key, region, ruleset, name, class, user_id, realm_slug, realm_name,
		    level, faction, bnet_character_id, source, imported_at, refreshed_at,
		    race, gender, bnet_account, bnet_captured_at)
		 values ($1, $2, $3, $4, $5, $6, $7, $8, $9, nullif($10, ''), $11, 'bnet', now(), now(),
		    nullif($12, ''), nullif($13, ''), $14, now())
		 on conflict (key) do update set
		   name = excluded.name, class = excluded.class, user_id = excluded.user_id,
		   realm_slug = excluded.realm_slug, realm_name = excluded.realm_name,
		   level = excluded.level, faction = excluded.faction, bnet_character_id = excluded.bnet_character_id,
		   source = 'bnet', imported_at = now(), refreshed_at = now(),
		   race = excluded.race, gender = excluded.gender,
		   bnet_account = coalesce(excluded.bnet_account, characters.bnet_account), bnet_captured_at = now()
		 where characters.user_id is null or characters.user_id = excluded.user_id`,
		key, region, ruleset, ch.Name, ch.ClassSlug, userID, ch.RealmSlug, ch.RealmName,
		ch.Level, faction, ch.ID, race, gender, rawOrNil(accountRaw))
```

  Note `nullif($12,'')`/`nullif($13,'')` reuse the existing `nullif($10,'')` faction
  pattern for an empty (not just missing) race/gender.

- [ ] Step 3: Still in `character.go`, before that insert, add the rekey guard (spec A3),
  called as `if err := s.rekeyIfNeeded(ctx, tx, userID, key, ch.ID); err != nil { return
  false, false, false, err }`:

```go
// rekeyIfNeeded implements spec A3: a character whose Blizzard id
// already belongs to this account under a *different* key (a realm's
// ruleset resolved differently between imports, or a realm transfer)
// gets moved, not duplicated. The old guild_characters row is deleted
// (running AfterGuildChange for its guild and user) and the old
// characters row is deleted, inside the same transaction the caller
// commits after writing the new row.
func (s *Service) rekeyIfNeeded(ctx context.Context, tx pgx.Tx, userID int64, newKey string, bnetCharacterID int64) error {
	var oldKey string
	err := tx.QueryRow(ctx,
		`select key from characters where bnet_character_id = $1 and user_id = $2 and key != $3`,
		bnetCharacterID, userID, newKey).Scan(&oldKey)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil
	}
	if err != nil {
		return fmt.Errorf("bnetimport: rekey lookup for %s: %w", newKey, err)
	}

	var guildID *int64
	err = tx.QueryRow(ctx, `select guild_id from guild_characters where character_key = $1`, oldKey).Scan(&guildID)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return fmt.Errorf("bnetimport: rekey guild lookup for %s: %w", oldKey, err)
	}
	if guildID != nil {
		if _, err := tx.Exec(ctx, `delete from guild_characters where character_key = $1`, oldKey); err != nil {
			return fmt.Errorf("bnetimport: rekey clear guild for %s: %w", oldKey, err)
		}
		if err := guilds.AfterGuildChange(ctx, tx, *guildID, userID); err != nil {
			return err
		}
	}
	if _, err := tx.Exec(ctx, `delete from characters where key = $1`, oldKey); err != nil {
		return fmt.Errorf("bnetimport: rekey delete old row %s: %w", oldKey, err)
	}
	s.logger().Info("bnetimport", "op", "rekey", "old_key", oldKey, "new_key", newKey)
	return nil
}
```

- [ ] Step 4: In `syncCharacterGuild`, change the signature to `(guildWritten, unavailable
  bool, err error)`. Update the `Character` call site for the new three-return signature;
  on `errors.Is(err, bnetapi.ErrNotFound)`, log `s.logger().Info("bnetimport", "op",
  "profile_unavailable", "key", key, "status", http.StatusNotFound)` and `return false, true,
  nil`; on `ErrForbidden`, unchanged `return false, false, nil`; on any other error, `return
  false, false, fmt.Errorf(...)`. On success, call `s.captureProfile(ctx, tx, key, profile,
  rawProfile)` immediately (before the `HasGuild` branch) and return its error if non-nil (as
  `false, false, err`). Every other `return` in the function gains a `false` for
  `unavailable`. Add `"net/http"` import.
- [ ] Step 5: In `character.go`'s `importOneCharacter`, change its signature to `(wrote,
  guildWritten, unavailable bool, err error)`. After the successful guild sync (`guildWritten,
  unavailable, err = s.syncCharacterGuild(...)`), add `if err := s.captureEquipment(ctx, tx,
  key, region, ch.RealmSlug, ch.Name); err != nil { return false, false, false, err }` before
  `tx.Commit`. The final success return becomes `return true, guildWritten, unavailable,
  nil`.
- [ ] Step 6: In `import.go`'s `importRegion`, update the `importOneCharacter` call to
  4-value: `wrote, guildWritten, unavailable, err := s.importOneCharacter(...)`. After the
  existing `summary.Characters++`/`if guildWritten { summary.Guilds++ }`, add `if unavailable
  { summary.Unavailable++ }`.
- [ ] Step 7: In `auth/importer.go`, add `Unavailable int // spec A2: a 404'd character
  profile (e.g. a classic1x/SoD character)` to `ImportSummary`.
- [ ] Step 8: In `auth/handler.go`'s `runImport`, add `"unavailable", summary.Unavailable` to
  the existing `s.logger().Info(...)` call.
- [ ] Step 9: In `refresh.go`, change `staleCharacter` to `{BnetCharacterID int64; RealmSlug
  string; UserID int64}` and `staleBnetCharacters`'s query to `select coalesce(bnet_character_id,
  0), coalesce(realm_slug, ''), user_id from characters where source = 'bnet' and
  refreshed_at < now() - $1::interval order by refreshed_at` (drop `key`, `region`,
  `ruleset`, `name` from the select and the scan). Rewrite `refreshOneCharacter` to
  re-resolve the *current* row from `sc.BnetCharacterID`/`sc.RealmSlug` inside its own
  transaction rather than trusting a key captured at the batch snapshot (spec A3: "the
  refresh job must select rows for re-sync by bnet_character_id and realm, not by the
  possibly stale key"):

```go
func (s *Service) refreshOneCharacter(ctx context.Context, sc staleCharacter, rosterCache map[string]bnetapi.Roster) error {
	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("bnetimport: refresh begin bnet_character_id=%d: %w", sc.BnetCharacterID, err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	var key, region, ruleset, name string
	err = tx.QueryRow(ctx,
		`select key, region, ruleset, name from characters
		 where bnet_character_id = $1 and coalesce(realm_slug, '') = $2 and source = 'bnet'`,
		sc.BnetCharacterID, sc.RealmSlug).Scan(&key, &region, &ruleset, &name)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil // rekeyed or removed since the batch was snapshotted; nothing to do
	}
	if err != nil {
		return fmt.Errorf("bnetimport: refresh resolve bnet_character_id=%d: %w", sc.BnetCharacterID, err)
	}

	if _, err := tx.Exec(ctx, `update characters set refreshed_at = now() where key = $1`, key); err != nil {
		return fmt.Errorf("bnetimport: refresh stamp %s: %w", key, err)
	}
	if _, _, err := s.syncCharacterGuild(ctx, tx, sc.UserID, region, ruleset, key, sc.RealmSlug, name, rosterCache); err != nil {
		return err
	}
	if err := s.captureEquipment(ctx, tx, key, region, sc.RealmSlug, name); err != nil {
		return err
	}
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("bnetimport: refresh commit %s: %w", key, err)
	}
	return nil
}
```

  Add `"github.com/jackc/pgx/v5"` import to `refresh.go`.

- [ ] Step 10: Update `import_test.go`'s two existing tests: add
  `"playable_race":{"name":"Human"},"gender":{"type":"MALE"}` to the account-characters
  fixture body, add a `GET
  /profile/wow/character/whitemane/thoradin/equipment?namespace=profile-classic1x-us`
  fixture answering `{"equipped_items":[]}`, and assert (in the guilded test)
  `race, gender string; ... select race, gender from characters where key =
  'us/pvp/thoradin'` reads `"human"`/`"male"`... wait, race is the *name* verbatim
  ("Human"), not lowercased (spec §B: "race... come from it (playable_race.name..."). Assert
  `race == "Human"`, `gender == "male"`.
- [ ] Step 11: Update `refresh_test.go`'s three tests to insert an explicit
  `bnet_character_id` (e.g. `101` for `us/pvp/stale`, `201`/`202` for `one`/`two`, `301` for
  `left`) in every `insert into characters (...)` that used to be found by key; update
  `TestStaleBnetCharactersFindsOnlyOldBnetRows`'s assertion from `stale[0].Key` to
  `stale[0].BnetCharacterID == 101`. Add equipment fixture stubs (`GET
  .../equipment?namespace=...` → `{"equipped_items":[]}`) for every character whose profile
  fixture answers 200 (`two`, `left`) — the rate-limited-`one` case never reaches equipment.
- [ ] Step 12: Create `capture_test.go` with `TestCapCaptureRefusesAnOversizedBody` (build a
  `json.RawMessage` of `maxCaptureBytes+1` bytes, assert `capCapture` returns nil) and
  `TestCapCaptureKeepsABodyAtTheLimit` (exactly `maxCaptureBytes` bytes, assert it is
  returned unchanged) — pure unit tests, no database (`&Service{}` with no `Log` uses
  `slog.Default()`).
- [ ] Step 13: Add `TestImportAccountRekeysWhenARealmsRulesetResolvesDifferently` to
  `import_test.go`: seed a `characters` row directly at `us/normal/dottzz` with
  `bnet_character_id = 777`, `user_id = uid`, `source = 'bnet'`, and a `guild_characters` row
  for it in some guild with `source='bnet'`. Run `ImportAccount` against a fixture where the
  realm index/detail resolves `whitemane` (reuse an existing realm slug already in the
  fixture, e.g. rename to a PVP-typed realm) to `pvp` and the account-characters response
  names character id `777` at that realm. Assert: exactly one row exists for
  `bnet_character_id = 777`, its `key = 'us/pvp/dottzz'`, the old
  `us/normal/dottzz` key is gone from `characters` and from `guild_characters`.
- [ ] Step 14: Add `TestImportAccountLogsAndCountsA404CharacterProfileAsUnavailable`: account
  characters fixture names one character; the character-profile fixture answers 404; assert
  `summary.Unavailable == 1` and `summary.Characters == 1` (the characters row is still
  written — only the guild/profile capture is unavailable) and no `guild_characters` row
  exists.
- [ ] Step 15: Run `cd api && GOWORK=off go vet ./... && GOWORK=off
  TEST_DATABASE_URL=postgres://forever:forever@localhost:5434/forever_test?sslmode=disable go
  test -p 1 ./internal/bnetimport/... ./internal/bnetapi/... ./internal/auth/... -v` —
  expect all PASS.
- [ ] Step 16: Commit: `feat(bnetimport): capture raw Blizzard payloads, log profile_unavailable, rekey a moved character`.

## Task 6: `auth` — `race`/`gender`/`item_level` on `/v1/me`

**Files:**
- Modify: `api/internal/auth/store.go`
- Modify: `api/internal/auth/store_test.go`

**Interfaces:**
- Produces: `Character.Race, Gender string` (json `race,omitempty` / `gender,omitempty`),
  `Character.ItemLevel *int` (json `item_level,omitempty`), sourced from
  `characters.race`/`.gender`/`.equipped_item_level`.

- [ ] Step 1: Add to the `Character` struct in `store.go`, near `Faction`:

```go
	// Race and Gender come from the Battle.net account-profile capture
	// (spec §B/C); omitted when the character has never been through
	// the Battle.net import.
	Race   string `json:"race,omitempty"`
	Gender string `json:"gender,omitempty"`
	// ItemLevel is the character's equipped item level from the last
	// Battle.net profile capture; omitted when unknown.
	ItemLevel *int `json:"item_level,omitempty"`
```

- [ ] Step 2: Add `coalesce(c.race, ''), coalesce(c.gender, ''), c.equipped_item_level,` to
  `characterColumns` (after `c.source,`), and `&c.Race, &c.Gender, &c.ItemLevel,` to the
  `rows.Scan` call in `scanCharacterRows`, in the matching position.
- [ ] Step 3: Add `TestCharactersIncludesRaceGenderAndItemLevel` to `store_test.go`: insert a
  `characters` row with `race = 'Dwarf'`, `gender = 'male'`, `equipped_item_level = 54`,
  read it back via `(*Store).Characters`, assert `c.Race == "Dwarf"`, `c.Gender == "male"`,
  `*c.ItemLevel == 54`; add a second row with those columns left null and assert `c.Race ==
  ""`, `c.Gender == ""`, `c.ItemLevel == nil` (so the JSON `omitempty` tags actually omit
  them).
- [ ] Step 4: Run `cd api && GOWORK=off TEST_DATABASE_URL=postgres://forever:forever@localhost:5434/forever_test?sslmode=disable go test ./internal/auth/... -run Character -v` — expect PASS.
- [ ] Step 5: Commit: `feat(auth): /v1/me characters gain race, gender, item_level`.

## Task 7: Whole-branch review pass

- [ ] Step 1: Run `cd api && GOWORK=off go vet ./...`.
- [ ] Step 2: Run `cd api && GOWORK=off TEST_DATABASE_URL=postgres://forever:forever@localhost:5434/forever_test?sslmode=disable go test -p 1 ./...` and record the pass count.
- [ ] Step 3: Re-read the brief's A1–A3, B, C sections against the diff; fix any gap found
  (a fix round, not a new task) before the final report.
- [ ] Step 4: Update the SDD progress ledger with every ruling made during implementation.
