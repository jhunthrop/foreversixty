# Battle.net-first API: Blizzard profile becomes a build — Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: superpowers:subagent-driven-development
> (recommended) or superpowers:executing-plans. Steps use checkbox (`- [ ]`) syntax for
> tracking. **Execution note (this lane):** given the lane's model/turn budget, the lane
> controller implements these tasks directly in-session (self-reviewed with `go vet`/`go
> test` at every commit and one final whole-branch review), rather than dispatching one
> fresh subagent per task — see the ledger at `.superpowers/sdd/2026-09-22-bnet-first-api/
> progress.md` for the ruling. Anyone re-executing this plan from scratch should still use
> subagent-driven-development per task.

**Goal:** A signed-in Battle.net character with a profile, equipment and specializations
becomes a simmable FS1 build automatically — at login import and nightly refresh — with no
addon required.

**Architecture:** A new pure package `api/internal/bnetbuild` maps Blizzard's three profile
sub-resources onto the existing FS1 v1 grammar, using the site's own talent/enchant/suffix/
race data (loaded once per process). `bnetimport` calls it after capturing profile,
equipment and specializations, and upserts the result into `addon_exports` with
`source='blizzard'`, newest-`captured_at`-wins. `sims/input.go` and `auth`'s `/v1/me` learn
the new source. No web changes (web is `bnet-first-web`'s lane).

**Tech Stack:** Go 1.25, pgx v5, Postgres, the existing `trees` package for client data.

**Spec:** `docs/superpowers/specs/2026-09-22-battlenet-first-design.md` (sections 1, 2, 5
bind this lane) and `docs/superpowers/specs/2026-09-19-simulator-parity-interfaces.md` §7,
§10.5 (the FS1 grammar) and `web/src/lib/planner/fs1.ts` (the decoder this lane's output
must satisfy).

## Global Constraints

- Vocabulary: a Blizzard-sourced build is "from Battle.net" in any user-facing text; the
  code's internal `armory`/`bnet` naming is unaffected (spec §1.1). This lane writes no
  user-facing copy.
- Honesty: a character is simmable only when the site holds a real build; nothing inferred
  or faked. An unmatched talent, a clamped rank, an unresolved suffix and a skipped slot are
  each reported, never silently guessed (spec §1.3, §2.2).
- No token persistence (spec §1.4, unchanged from the import spec — this lane touches no
  token storage).
- File ownership: this lane owns `api/**` and `openapi.yaml`. It must not touch `web/**`.
- House rules (`.superpowers/journeys/lane-common-go.md`): every subagent on `sonnet`, never
  `opus`; `go vet ./... && go test ./...` (scoped to touched packages, `-p 1`) before every
  commit; commit messages via `printf` + `git commit -F`, never a heredoc combined with
  `git commit`; conventional subjects; the two-line sign-off; never touch the production
  database; never a broad `pkill`.
- Migration `0027` is next (main's highest is `0026`); the `addon_exports` SQL in spec §2.1
  is verbatim and binding.
- `Inputs`/`Report`/`Encode`'s signature (spec §2.2) is binding: `bnetbuild` takes exactly
  `Build, Profile, Equipment, Talents, Talent, Enchants, Suffixes, Races` and returns
  `(code string, report Report, err error)`.

---

## Ruling log (recorded here; mirrored in the SDD ledger)

1. **`TalentTable` wraps `*trees.Build` + a class id**, not a fresh loader. `trees.Load`
   already reads `data/builds/<build>/talents/<class>.json` for every class at startup
   (`internal/trees/trees.go`); duplicating that parser in `bnetbuild` would violate DRY and
   risk drift. `bnetbuild.TalentTable{Build: b, ClassID: id}` is the whole type.
2. **`EnchantTable`/`SuffixTable` are loaded by `bnetbuild` itself** (`enchants.json`,
   `suffixes.json`) since no existing package reads them; `RaceTable` is built from a new
   small `trees.Build.Races()` accessor (mirrors the existing `Classes()`), avoiding a third
   parser for data `trees` already holds.
3. **Enchant ids are passed through unvalidated.** `EnchantTable` is loaded and threaded
   through `Inputs` (required by the frozen signature) but the gear encoder does not gate on
   it: Blizzard's `enchantment_id` is emitted verbatim when `enchantment_slot.type ==
   "PERMANENT"`. Cost if wrong: a stale/removed enchant id could appear in an FS1 string the
   simulator's own `enchants.json` doesn't resolve — no worse than what the addon path
   already allows (it never validates enchant ids from the game client either).
4. **Suffix matching has no item-eligibility gate.** The frozen `Inputs` struct carries no
   `items.json` table, so "the item has suffixes in items.json" (spec §2.2) cannot be
   checked. Matching is: does the item's Blizzard `name` end with `" " + <a SuffixTable
   name>`, longest match wins. A miss is not reported (`Report.NoSuffix` stays empty for it)
   — without item-eligibility data, a miss cannot be distinguished from "this item never had
   a suffix," and flagging every non-matching named item (most loot) would be noise, not
   honesty. Verified against the era-kiloz fixture: zero of its 17 simmable-slot items
   accidentally matches a real `suffixes.json` name, so this produces no false positives on
   the one fixture this lane can check by hand.
5. **`addon_exports.captured_at` has no column default** (per spec §2.1's exact SQL), so
   every existing direct-SQL test insert into `addon_exports` needs an explicit value. Two
   call sites are affected (`internal/sims/input_test.go`, `internal/guilds/home_test.go`)
   and are updated in Task 1.
6. **`/v1/me`'s `characters[]` today has no OpenAPI schema at all** (checked: `openapi.yaml`
   models `/v1/me`'s response only as `{ data: { user: User } }}`, with `characters`
   undocumented already). Spec §2.5 says "OpenAPI too" only for the `sim-input` enum, so
   only that enum is updated; `/v1/me` stays as under-specified as it already was.

---

## Task 1: Migration `0027` and its test-fixture ripple

**Files:**
- Create: `api/internal/db/migrations/0027_export_sources.up.sql`
- Create: `api/internal/db/migrations/0027_export_sources.down.sql`
- Modify: `api/internal/sims/input_test.go` (the `insertExport` helper)
- Modify: `api/internal/guilds/home_test.go` (the `seedExport`-shaped helper)

**Interfaces:**
- Produces: `addon_exports.source text` (`'addon'|'blizzard'`, default `'addon'`),
  `addon_exports.captured_at timestamptz not null` (no default), `characters.bnet_talents
  jsonb` (spec §2.3's new capture column, same shape as `bnet_equipment`).

- [ ] **Step 1: Write the up/down migrations**

`api/internal/db/migrations/0027_export_sources.up.sql`:
```sql
-- api/internal/db/migrations/0027_export_sources.up.sql
-- Battle.net-first (spec docs/superpowers/specs/2026-09-22-battlenet-first-design.md §2.1,
-- §2.3): addon_exports learns which path wrote it and when it was true in the game, and
-- characters gains a raw capture column for the specializations read, alongside
-- bnet_profile/bnet_equipment (migration 0025).
alter table addon_exports
  add column if not exists source text not null default 'addon' check (source in ('addon', 'blizzard')),
  add column if not exists captured_at timestamptz;
update addon_exports set captured_at = updated_at where captured_at is null;
alter table addon_exports alter column captured_at set not null;

alter table characters
  add column if not exists bnet_talents jsonb;
```

`api/internal/db/migrations/0027_export_sources.down.sql`:
```sql
-- api/internal/db/migrations/0027_export_sources.down.sql
alter table characters drop column if exists bnet_talents;
alter table addon_exports drop column if exists captured_at;
alter table addon_exports drop column if exists source;
```

- [ ] **Step 2: Fix the two direct-SQL test inserts the new NOT NULL column breaks**

In `api/internal/sims/input_test.go`, the `insertExport` helper (around line 28) currently
inserts `(character_key, user_id, region, ruleset, name, export, updated_at)`. Add
`captured_at` set to the same `at` value the helper already takes:
```go
	if _, err := h.store.Pool.Exec(h.t.Context(),
		`insert into addon_exports (character_key, user_id, region, ruleset, name, export, captured_at, updated_at)
		 values ($1, $2, $3, $4, $5, $6, $7, $7)
		 on conflict (character_key) do update set export = excluded.export,
		   captured_at = excluded.captured_at, updated_at = excluded.updated_at`,
		key, h.owner, region, ruleset, name, export, at); err != nil {
		h.t.Fatal(err)
	}
```

In `api/internal/guilds/home_test.go`, the helper around line 44 inserts a bare placeholder
export; add `captured_at`:
```go
	if _, err := pool.Exec(context.Background(), `
		insert into addon_exports (character_key, user_id, region, ruleset, name, export, captured_at, updated_at)
		values ($1, $2, $3, $4, $5, '', now(), now())`,
		key, userID, region, ruleset, name); err != nil {
		t.Fatal(err)
	}
```

- [ ] **Step 3: Run the migration and the two touched packages' tests**

```bash
cd api && export GOWORK=off TEST_DATABASE_URL=postgres://forever:forever@localhost:5434/forever_test?sslmode=disable
go vet ./internal/db/... ./internal/sims/... ./internal/guilds/...
go test ./internal/db/... ./internal/sims/... ./internal/guilds/... -p 1
```
Expected: PASS (the migration applies cleanly against a fresh test DB; both fixed tests
pass).

- [ ] **Step 4: Commit**

```bash
cd api && git add internal/db/migrations/0027_export_sources.up.sql internal/db/migrations/0027_export_sources.down.sql internal/sims/input_test.go internal/guilds/home_test.go
```
Message via `printf`/`git commit -F` (see house rules): `feat(db): addon_exports learns its source and capture time (migration 0027)`.

---

## Task 2: `bnetapi` — character race name and the specializations endpoint

**Files:**
- Modify: `api/internal/bnetapi/profile.go`
- Modify: `api/internal/bnetapi/profile_test.go`

**Interfaces:**
- Produces: `CharacterProfile.RaceName string` (Blizzard's `race.name`, title case, e.g.
  "Night Elf" — matches `data/builds/<build>/races.json`'s own `name` field verbatim, spec
  §2.2's race-slug lookup). `Client.Specializations(ctx, region, realmSlug, name)
  (json.RawMessage, error)`, mirroring `Client.Equipment`.

- [ ] **Step 1: Write the failing tests**

Add to `api/internal/bnetapi/profile_test.go` (extend the existing guilded/unguilded
fixtures with a `"race"` field, and add a new test for the endpoint):
```go
// in TestCharacterReadsGuildedAndUnguilded's guilded fixture body, add:
// "race":{"name":"Orc"},
```
Change the guilded fixture body to:
```go
	fs.handlers[http.MethodGet+" /profile/wow/character/whitemane/thoradin?namespace=profile-classic1x-us"] = func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"name":"Thoradin","level":42,"faction":{"type":"ALLIANCE"},
			"character_class":{"name":"Warrior"},"race":{"name":"Orc"},
			"realm":{"slug":"whitemane","name":"Whitemane"},
			"guild":{"name":"Iron Vanguard","id":12},"last_login_timestamp":1700000000000,
			"average_item_level":55,"equipped_item_level":54}`))
	}
```
and after the existing `guilded.AverageItemLevel` assertions:
```go
	if guilded.RaceName != "Orc" {
		t.Fatalf("RaceName = %q, want the verbatim race.name", guilded.RaceName)
	}
```
New test:
```go
func TestSpecializationsReturnsTheRawBody(t *testing.T) {
	fs := newFixtureServer(t)
	fs.json(http.MethodPost, "/token", http.StatusOK, map[string]any{"access_token": "tok", "expires_in": 3600})
	fs.json(http.MethodGet, "/profile/wow/character/whitemane/thoradin/specializations?namespace=profile-classic1x-us",
		http.StatusOK, map[string]any{"specialization_groups": []map[string]any{{"is_active": true}}})

	c := newTestClient(fs)
	raw, err := c.Specializations(context.Background(), "us", "whitemane", "Thoradin")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(raw), "specialization_groups") {
		t.Fatalf("raw = %s, want the verbatim specializations body", raw)
	}
}
```

- [ ] **Step 2: Run to verify failure**

```bash
cd api && export GOWORK=off && go test ./internal/bnetapi/... -run 'TestCharacterReadsGuildedAndUnguilded|TestSpecializationsReturnsTheRawBody' -v
```
Expected: FAIL (`RaceName` field/`Specializations` method do not exist yet).

- [ ] **Step 3: Implement**

In `api/internal/bnetapi/profile.go`, add `RaceName` to `CharacterProfile` and
`characterProfileResponse`, and fill it:
```go
type CharacterProfile struct {
	Name               string
	Level              int
	Faction            string
	ClassSlug          string
	// RaceName is Blizzard's race.name, verbatim (title case, e.g. "Night Elf") — the
	// same string data/builds/<build>/races.json keys its own race rows by (spec
	// docs/superpowers/specs/2026-09-22-battlenet-first-design.md §2.2).
	RaceName           string
	RealmSlug          string
	RealmName          string
	GuildName          string
	HasGuild           bool
	LastLoginTimestamp int64
	AverageItemLevel  *int
	EquippedItemLevel *int
}

type characterProfileResponse struct {
	Name    string `json:"name"`
	Level   int    `json:"level"`
	Faction struct {
		Type string `json:"type"`
	} `json:"faction"`
	CharacterClass struct {
		Name string `json:"name"`
	} `json:"character_class"`
	Race struct {
		Name string `json:"name"`
	} `json:"race"`
	Realm struct {
		Slug string `json:"slug"`
		Name string `json:"name"`
	} `json:"realm"`
	Guild *struct {
		Name string `json:"name"`
	} `json:"guild"`
	LastLoginTimestamp int64 `json:"last_login_timestamp"`
	AverageItemLevel   *int  `json:"average_item_level"`
	EquippedItemLevel  *int  `json:"equipped_item_level"`
}
```
In `Client.Character`, add `RaceName: res.Race.Name` to the `CharacterProfile{...}` literal.

Add, after `Equipment`:
```go
// Specializations reads a character's talent groups verbatim. Not parsed here — the
// Blizzard-to-FS1 talent mapping is bnetbuild's job (spec §2.2) — so the caller stores the
// raw body directly, the same shape Equipment already returns.
func (c *Client) Specializations(ctx context.Context, region, realmSlug, name string) (json.RawMessage, error) {
	token, err := c.AppToken(ctx)
	if err != nil {
		return nil, fmt.Errorf("bnetapi: specializations: %w", err)
	}
	u := c.APIHost(region) + "/profile/wow/character/" + url.PathEscape(realmSlug) + "/" +
		url.PathEscape(strings.ToLower(name)) + "/specializations?namespace=" + c.ProfileNamespace(region)
	body, err := c.getBytes(ctx, "specializations", u, token)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(body), nil
}
```

- [ ] **Step 4: Run to verify pass**

```bash
cd api && export GOWORK=off && go vet ./internal/bnetapi/... && go test ./internal/bnetapi/... -p 1
```
Expected: PASS.

- [ ] **Step 5: Commit** — `feat(bnetapi): character race name and the specializations endpoint`.

---

## Task 3: `trees.Build.Races()` accessor

**Files:**
- Modify: `api/internal/trees/trees.go`
- Modify: `api/internal/trees/trees_test.go`

**Interfaces:**
- Produces: `func (b *Build) Races() []Race`, sorted by id (mirrors `Classes()`).

- [ ] **Step 1: Write the failing test** — add to `trees_test.go` after the existing
`TestLoadFixtureReadsTheTwoTreeClass` race assertion:
```go
	races := b.Races()
	if len(races) == 0 {
		t.Fatal("Races() must list every race the fixture loaded")
	}
	found := false
	for _, race := range races {
		if race.ID == 3 && race.Name == "Dwarf" {
			found = true
		}
	}
	if !found {
		t.Fatal("Races() must include race 3 (Dwarf) from the fixture")
	}
```

- [ ] **Step 2: Run to verify failure** — `cd api && export GOWORK=off && go test ./internal/trees/... -run TestLoadFixtureReadsTheTwoTreeClass -v` → FAIL (no `Races` method).

- [ ] **Step 3: Implement** — in `trees.go`, immediately after `Classes()`:
```go
// Races lists the build's playable races by id. bnetbuild uses it to build a
// Blizzard-race-name-to-slug table without re-parsing races.json (spec
// docs/superpowers/specs/2026-09-22-battlenet-first-design.md §2.2).
func (b *Build) Races() []Race {
	out := make([]Race, 0, len(b.races))
	for _, r := range b.races {
		out = append(out, r)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out
}
```

- [ ] **Step 4: Run to verify pass** — `go vet ./internal/trees/... && go test ./internal/trees/... -p 1` → PASS.

- [ ] **Step 5: Commit** — `feat(trees): list a build's races (Build.Races)`.

---

## Task 4: `bnetbuild` package skeleton — tables and the SLOTS parity test

**Files:**
- Create: `api/internal/bnetbuild/tables.go`
- Create: `api/internal/bnetbuild/tables_test.go`
- Create: `api/internal/bnetbuild/slots.go`
- Create: `api/internal/bnetbuild/slots_test.go`

**Interfaces:**
- Produces: `type TalentTable struct { Build *trees.Build; ClassID int }`,
  `type EnchantTable map[int]bool`, `func LoadEnchantTable(path string) (EnchantTable, error)`,
  `type SuffixTable map[string]int`, `func LoadSuffixTable(path string) (SuffixTable, error)`,
  `type RaceTable map[string]string`, `func RaceTableFrom(b *trees.Build) RaceTable`,
  `type Tables struct { Build string; Trees *trees.Build; Enchants EnchantTable; Suffixes
  SuffixTable; Races RaceTable }`, `func LoadTables(dir string, data *trees.Data) (Tables,
  bool, error)`, `var SLOTS = [...]string{...}` (Go copy of `web/src/lib/planner/types.ts`'s
  `SLOTS`).
- Consumes: `trees.Data`, `trees.Build` (Task 3).

- [ ] **Step 1: Write the failing SLOTS parity test** (`slots_test.go`) — this is the
"copy the list into Go with a test that reads the TS file and compares" spec §2.2 asks for:
```go
// api/internal/bnetbuild/slots_test.go
package bnetbuild

import (
	"os"
	"path/filepath"
	"regexp"
	"testing"
)

// slotsArrayPattern extracts the quoted strings inside web/src/lib/planner/types.ts's
// `export const SLOTS = [...] as const;` block, so this test fails the moment that file's
// slot list changes without this package's copy changing with it.
var slotsArrayPattern = regexp.MustCompile(`export const SLOTS = \[([\s\S]*?)\] as const;`)

func TestSlotsMatchesThePlannerTypesFile(t *testing.T) {
	path := filepath.Join("..", "..", "..", "web", "src", "lib", "planner", "types.ts")
	src, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}
	m := slotsArrayPattern.FindSubmatch(src)
	if m == nil {
		t.Fatalf("could not find the SLOTS array in %s", path)
	}
	quoted := regexp.MustCompile(`'([a-z0-9_]+)'`).FindAllStringSubmatch(string(m[1]), -1)
	var want []string
	for _, q := range quoted {
		want = append(want, q[1])
	}
	if len(want) == 0 {
		t.Fatal("parsed zero slots out of types.ts; the regexp or the file format changed")
	}
	if len(SLOTS) != len(want) {
		t.Fatalf("bnetbuild.SLOTS has %d entries, types.ts has %d: %v vs %v", len(SLOTS), len(want), SLOTS, want)
	}
	for i, slot := range SLOTS {
		if slot != want[i] {
			t.Fatalf("SLOTS[%d] = %q, types.ts has %q at the same position", i, slot, want[i])
		}
	}
}
```

- [ ] **Step 2: Run to verify failure** — `cd api && export GOWORK=off && go test ./internal/bnetbuild/... -v` → FAIL (package/SLOTS do not exist).

- [ ] **Step 3: Implement `slots.go`**
```go
// api/internal/bnetbuild/slots.go
package bnetbuild

// SLOTS is web/src/lib/planner/types.ts's SLOTS constant, copied so an FS1 gear line this
// package writes uses exactly the slot names the planner's decoder (fs1.ts) accepts. A test
// (slots_test.go) reads types.ts and refuses to let the two drift.
var SLOTS = [...]string{
	"head", "neck", "shoulder", "back", "chest", "wrist", "hands", "waist", "legs", "feet",
	"finger1", "finger2", "trinket1", "trinket2", "main_hand", "off_hand", "ranged",
}
```

- [ ] **Step 4: Write the failing tables tests** (`tables_test.go`) — uses the real active
build (`data/builds/1.60.1.69893`), three levels up from this package:
```go
// api/internal/bnetbuild/tables_test.go
package bnetbuild

import (
	"path/filepath"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

const realDataDir = "../../../data/builds"

func TestLoadEnchantTableFindsTheEraKilozHeadEnchant(t *testing.T) {
	table, err := LoadEnchantTable(filepath.Join(realDataDir, "1.60.1.69893", "enchants.json"))
	if err != nil {
		t.Fatal(err)
	}
	if !table[2583] {
		t.Fatal("enchant 2583 (Presence of Might, on the fixture's head slot) must be known")
	}
	if table[999999999] {
		t.Fatal("an id absent from enchants.json must not read as known")
	}
}

func TestLoadSuffixTableFindsAKnownSuffix(t *testing.T) {
	table, err := LoadSuffixTable(filepath.Join(realDataDir, "1.60.1.69893", "suffixes.json"))
	if err != nil {
		t.Fatal(err)
	}
	if table["of the Falcon"] == 0 {
		t.Fatal(`"of the Falcon" must resolve to a non-zero suffix id`)
	}
}

func TestRaceTableFromMapsBlizzardNamesToSlugs(t *testing.T) {
	data, err := trees.Load(realDataDir)
	if err != nil {
		t.Fatal(err)
	}
	build, ok := data.Build("1.60.1.69893")
	if !ok {
		t.Fatal("1.60.1.69893 must load")
	}
	races := RaceTableFrom(build)
	if races["Orc"] != "orc" {
		t.Fatalf(`races["Orc"] = %q, want "orc"`, races["Orc"])
	}
}

func TestLoadTablesResolvesTheActiveBuild(t *testing.T) {
	data, err := trees.Load(realDataDir)
	if err != nil {
		t.Fatal(err)
	}
	tables, ok, err := LoadTables(realDataDir, data)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("LoadTables must resolve a build from the real data/builds directory")
	}
	if tables.Build != "1.60.1.69893" {
		t.Fatalf("Build = %q, want the newest numeric client build", tables.Build)
	}
	if tables.Trees == nil || tables.Enchants == nil || tables.Suffixes == nil || tables.Races == nil {
		t.Fatalf("LoadTables left a table nil: %+v", tables)
	}
}
```

- [ ] **Step 5: Run to verify failure** — FAIL (types/functions absent).

- [ ] **Step 6: Implement `tables.go`**
```go
// api/internal/bnetbuild/tables.go
package bnetbuild

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

// TalentTable is one class's talent data within a loaded trees.Data build (spec
// docs/superpowers/specs/2026-09-22-battlenet-first-design.md §2.2: "loaded once per class
// from data/builds/<build>/talents/<class>.json"). trees.Load already reads that file for
// every class at process startup; this wraps it rather than re-reading it.
type TalentTable struct {
	Build   *trees.Build
	ClassID int
}

// EnchantTable is the set of known permanent-enchant ids (data/builds/<build>/
// enchants.json), keyed by the same id Blizzard's equipment response calls
// enchantment_id (spec §2.2).
type EnchantTable map[int]bool

type enchantRow struct {
	ID int `json:"id"`
}

// LoadEnchantTable reads enchants.json's flat row list into a lookup set. Unknown JSON
// fields are ignored on purpose — the file carries name/icon/slots/stats this package has
// no use for.
func LoadEnchantTable(path string) (EnchantTable, error) {
	var rows []enchantRow
	if err := readJSON(path, &rows); err != nil {
		return nil, err
	}
	out := make(EnchantTable, len(rows))
	for _, r := range rows {
		out[r.ID] = true
	}
	return out, nil
}

// SuffixTable maps a random suffix's display name, exactly as it appears appended to an
// item's base name (e.g. "of the Falcon"), to its id (data/builds/<build>/suffixes.json).
type SuffixTable map[string]int

type suffixRow struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// LoadSuffixTable reads suffixes.json into a name-to-id lookup.
func LoadSuffixTable(path string) (SuffixTable, error) {
	var rows []suffixRow
	if err := readJSON(path, &rows); err != nil {
		return nil, err
	}
	out := make(SuffixTable, len(rows))
	for _, r := range rows {
		out[r.Name] = r.ID
	}
	return out, nil
}

// RaceTable maps Blizzard's race.name (title case, e.g. "Night Elf") to the site's own race
// slug (spec §2.2).
type RaceTable map[string]string

// RaceTableFrom builds a RaceTable from a loaded build's races, rather than re-parsing
// races.json (trees.Load already did, Task 3's Races() exposes it).
func RaceTableFrom(b *trees.Build) RaceTable {
	races := b.Races()
	out := make(RaceTable, len(races))
	for _, r := range races {
		out[r.Name] = r.Slug
	}
	return out
}

// Tables is every data table Encode needs, loaded once per process from the site's active
// client build (spec §2.3: "Data tables are loaded once per process from
// data/builds/<active build>/").
type Tables struct {
	Build    string
	Trees    *trees.Build
	Enchants EnchantTable
	Suffixes SuffixTable
	Races    RaceTable
}

// LoadTables resolves the active build (trees.Data's newest client build, the same rule
// dataaddon's job already uses) and loads its enchant, suffix and race tables. ok is false,
// with a zero Tables, when data holds no build at all — the caller (bnetimport, cmd/api)
// logs and runs with Battle.net-sourced builds disabled, same as when TreeDataDir is empty.
func LoadTables(dir string, data *trees.Data) (Tables, bool, error) {
	active, ok := data.Latest()
	if !ok {
		return Tables{}, false, nil
	}
	buildDir := filepath.Join(dir, active.Version)
	enchants, err := LoadEnchantTable(filepath.Join(buildDir, "enchants.json"))
	if err != nil {
		return Tables{}, false, fmt.Errorf("bnetbuild: load tables %s: %w", active.Version, err)
	}
	suffixes, err := LoadSuffixTable(filepath.Join(buildDir, "suffixes.json"))
	if err != nil {
		return Tables{}, false, fmt.Errorf("bnetbuild: load tables %s: %w", active.Version, err)
	}
	return Tables{
		Build: active.Version, Trees: active, Enchants: enchants, Suffixes: suffixes,
		Races: RaceTableFrom(active),
	}, true, nil
}

func readJSON(path string, v any) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("bnetbuild: open %s: %w", path, err)
	}
	defer f.Close()
	if err := json.NewDecoder(f).Decode(v); err != nil {
		return fmt.Errorf("bnetbuild: decode %s: %w", path, err)
	}
	return nil
}
```

- [ ] **Step 7: Run to verify pass**

```bash
cd api && export GOWORK=off && go vet ./internal/bnetbuild/... && go test ./internal/bnetbuild/... -p 1 -v
```
Expected: PASS (all four table tests, the SLOTS parity test).

- [ ] **Step 8: Commit** — `feat(bnetbuild): data tables and the SLOTS parity test`.

---

## Task 5: `bnetbuild` — talent matching and tree encoding

**Files:**
- Create: `api/internal/bnetbuild/talents.go`
- Create: `api/internal/bnetbuild/talents_test.go`

**Interfaces:**
- Consumes: `TalentTable` (Task 4), `trees.Build.Trees(classID)`, `.Talent(classID, id)`,
  `.TalentBySpellID(classID, spellID)` (all pre-existing).
- Produces: `type blizzardSpecGroups struct{...}` (unexported, decodes the
  `/specializations` body), `func encodeTalents(talentsRaw json.RawMessage, table
  TalentTable) (treeRanks [3][]int, unmatched []string, clamped []string, err error)`.
  `treeRanks[i]` is `nil` for a tree the active group never touches; `encodeTree` (below,
  Task 6) already treats a nil/short slice as all-zero ranks.

- [ ] **Step 1: Write the failing tests** (uses the real era-kiloz fixture and the real
active build):
```go
// api/internal/bnetbuild/talents_test.go
package bnetbuild

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

func loadWarriorTable(t *testing.T) TalentTable {
	t.Helper()
	data, err := trees.Load(realDataDir)
	if err != nil {
		t.Fatal(err)
	}
	build, ok := data.Build("1.60.1.69893")
	if !ok {
		t.Fatal("1.60.1.69893 must load")
	}
	var classID int
	for _, c := range build.Classes() {
		if c.Slug == "warrior" {
			classID = c.ID
		}
	}
	if classID == 0 {
		t.Fatal("warrior class not found")
	}
	return TalentTable{Build: build, ClassID: classID}
}

func readEraKiloz(t *testing.T, file string) []byte {
	t.Helper()
	b, err := os.ReadFile(filepath.Join("..", "bnetapi", "testdata", "era-kiloz", file))
	if err != nil {
		t.Fatal(err)
	}
	return b
}

func TestEncodeTalentsUsesTheActiveGroupAndReportsUnmatched(t *testing.T) {
	table := loadWarriorTable(t)
	raw := readEraKiloz(t, "specializations.json")

	treeRanks, unmatched, clamped, err := encodeTalents(raw, table)
	if err != nil {
		t.Fatal(err)
	}
	// The active group (is_active: true) is Fury + Arms, 51 points total; the site's
	// 1.60.1.69893 talent ids are a different DBC generation than the classic1x profile
	// API's, so most of these 15 talents cannot be id- or spell-matched and are reported —
	// this pins that today's real behaviour, not a bug this lane owes a fix for.
	if len(unmatched) == 0 {
		t.Fatal("this fixture's talent ids are known not to match this build's table; want at least one report")
	}
	if len(clamped) != 0 {
		t.Fatalf("clamped = %v, want none (every rank in the fixture is within max_rank)", clamped)
	}
	totalRanks := 0
	for _, tree := range treeRanks {
		for _, r := range tree {
			totalRanks += r
		}
	}
	if totalRanks == 0 {
		t.Fatal("at least the id-matched talents (single-rank abilities) must land a nonzero rank")
	}
}

func TestEncodeTalentsPicksTheFirstGroupWhenNoneIsActive(t *testing.T) {
	table := loadWarriorTable(t)
	raw := []byte(`{"specialization_groups":[
		{"is_active":false,"specializations":[{"specialization_name":"Fury","spent_points":1,
			"talents":[{"talent":{"id":167},"spell_tooltip":{"spell":{"id":23881,"name":"Bloodthirst"}},"talent_rank":1}]}]},
		{"is_active":false,"specializations":[{"specialization_name":"Arms","spent_points":0,"talents":[]}]}
	]}`)
	treeRanks, _, _, err := encodeTalents(raw, table)
	if err != nil {
		t.Fatal(err)
	}
	total := 0
	for _, tree := range treeRanks {
		for _, r := range tree {
			total += r
		}
	}
	if total == 0 {
		t.Fatal("no group is active: the first group (Bloodthirst rank 1) must still be used")
	}
}
```

- [ ] **Step 2: Run to verify failure** — FAIL (`encodeTalents` undefined).

- [ ] **Step 3: Implement `talents.go`**
```go
// api/internal/bnetbuild/talents.go
package bnetbuild

import (
	"encoding/json"
	"fmt"
)

type blizzardSpecGroups struct {
	SpecializationGroups []blizzardSpecGroup `json:"specialization_groups"`
}

type blizzardSpecGroup struct {
	IsActive        bool                `json:"is_active"`
	Specializations []blizzardSpecEntry `json:"specializations"`
}

type blizzardSpecEntry struct {
	Talents []blizzardTalentEntry `json:"talents"`
}

type blizzardTalentEntry struct {
	Talent struct {
		ID int `json:"id"`
	} `json:"talent"`
	SpellTooltip struct {
		Spell struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"spell"`
	} `json:"spell_tooltip"`
	TalentRank int `json:"talent_rank"`
}

// encodeTalents reads the Blizzard specializations body and returns one rank slice per
// tree, in the class's tree position order (spec §2.2: "the active specialization group is
// used; when none is active, the first"). unmatched and clamped name every talent that
// could not be placed or whose rank exceeded max_rank, by spell name (Report.UnmatchedTalents,
// Report.Clamped).
func encodeTalents(raw json.RawMessage, table TalentTable) (treeRanks [3][]int, unmatched, clamped []string, err error) {
	var body blizzardSpecGroups
	if err := json.Unmarshal(raw, &body); err != nil {
		return treeRanks, nil, nil, fmt.Errorf("bnetbuild: decode specializations: %w", err)
	}
	group := firstActiveOrFirst(body.SpecializationGroups)

	classTrees := table.Build.Trees(table.ClassID)
	for i, tree := range classTrees {
		if i >= 3 {
			break
		}
		treeRanks[i] = make([]int, len(tree.Talents))
	}

	for _, spec := range group.Specializations {
		for _, bt := range spec.Talents {
			ref, ok := table.Build.Talent(table.ClassID, bt.Talent.ID)
			if !ok {
				ref, ok = table.Build.TalentBySpellID(table.ClassID, bt.SpellTooltip.Spell.ID)
			}
			if !ok {
				unmatched = append(unmatched, displayName(bt))
				continue
			}
			rank := bt.TalentRank
			if rank > ref.MaxRank {
				clamped = append(clamped, displayName(bt))
				rank = ref.MaxRank
			}
			treeIndex, talentIndex, found := positionOf(classTrees, ref.TreeID, ref.Talent.ID)
			if !found || treeIndex >= 3 {
				continue
			}
			treeRanks[treeIndex][talentIndex] = rank
		}
	}
	return treeRanks, unmatched, clamped, nil
}

// firstActiveOrFirst is spec §2.2's own words: "the active specialization group is used;
// when none is active, the first."
func firstActiveOrFirst(groups []blizzardSpecGroup) blizzardSpecGroup {
	for _, g := range groups {
		if g.IsActive {
			return g
		}
	}
	if len(groups) == 0 {
		return blizzardSpecGroup{}
	}
	return groups[0]
}

// displayName prefers the spell's own name (what a player recognises) over the bare talent
// id, falling back to the id when Blizzard sent no tooltip name.
func displayName(bt blizzardTalentEntry) string {
	if bt.SpellTooltip.Spell.Name != "" {
		return bt.SpellTooltip.Spell.Name
	}
	return fmt.Sprintf("talent %d", bt.Talent.ID)
}

// positionOf finds a talent's tree index (by tree position, 0-based) and its index within
// that tree's own Talents slice — the two coordinates encodeTree (talent_encode.go, Task 6)
// needs, matching web/src/lib/planner/fs1.ts's own "array index is tab order" contract.
func positionOf(classTrees []treeShape, treeID, talentID int) (treeIndex, talentIndex int, found bool) {
	for i, tree := range classTrees {
		for j, t := range tree.Talents {
			if tree.ID == treeID && t.ID == talentID {
				return i, j, true
			}
		}
	}
	return 0, 0, false
}
```

**Note for the implementer:** `table.Build.Trees(classID)` returns `[]trees.Tree`, not a
locally-defined `treeShape` — `positionOf`'s parameter type must be `[]trees.Tree`, and the
import list needs `"github.com/jhunthrop/foreversixty/api/internal/trees"`. Fix the
signature to `func positionOf(classTrees []trees.Tree, treeID, talentID int) (int, int,
bool)` before compiling (the plan's own placeholder name `treeShape` is not a real type —
correct this on implementation, since `trees.Tree`/`trees.Talent` are the real shapes
already defined in Task 3's package).

- [ ] **Step 4: Run to verify pass**

```bash
cd api && export GOWORK=off && go vet ./internal/bnetbuild/... && go test ./internal/bnetbuild/... -p 1 -v
```
Expected: PASS.

- [ ] **Step 5: Commit** — `feat(bnetbuild): match Blizzard talents onto the active build's trees`.

---

## Task 6: `bnetbuild` — gear encoding (slots, enchants, suffixes)

**Files:**
- Create: `api/internal/bnetbuild/gear.go`
- Create: `api/internal/bnetbuild/gear_test.go`

**Interfaces:**
- Consumes: `SLOTS`, `EnchantTable`, `SuffixTable` (Task 4).
- Produces: `func encodeGear(equipmentRaw json.RawMessage, suffixes SuffixTable) (gear
  string, skippedSlots, noSuffix []string, err error)`.

- [ ] **Step 1: Write the failing tests**
```go
// api/internal/bnetbuild/gear_test.go
package bnetbuild

import (
	"strings"
	"testing"
)

func TestEncodeGearMapsEveryFixtureSlotAndSkipsShirtAndTabard(t *testing.T) {
	raw := readEraKiloz(t, "equipment.json")
	gear, skipped, noSuffix, err := encodeGear(raw, SuffixTable{})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(gear, "head=21329:2583") {
		t.Fatalf("gear = %q, want head=21329:2583 (Conqueror's Crown, Presence of Might)", gear)
	}
	if !strings.Contains(gear, "main_hand=22816:1900") || !strings.Contains(gear, "off_hand=18828:1900") {
		t.Fatalf("gear = %q, want both weapon hands with their Crusader enchant", gear)
	}
	if strings.Contains(gear, "trinket1=11815:") == false {
		t.Fatalf("gear = %q, want trinket1=11815 (Hand of Justice, no enchant)", gear)
	}
	wantSkipped := map[string]bool{"SHIRT": false, "TABARD": false}
	for _, s := range skipped {
		if _, ok := wantSkipped[s]; ok {
			wantSkipped[s] = true
		}
	}
	if !wantSkipped["SHIRT"] || !wantSkipped["TABARD"] {
		t.Fatalf("skippedSlots = %v, want SHIRT and TABARD both reported", skipped)
	}
	if len(noSuffix) != 0 {
		t.Fatalf("noSuffix = %v, want none: no fixture item's name matches a suffix", noSuffix)
	}
}

func TestEncodeGearMatchesASuffixedItemName(t *testing.T) {
	raw := []byte(`{"equipped_items":[{"item":{"id":1234},"slot":{"type":"WAIST"},"name":"Girdle of the Falcon","enchantments":[]}]}`)
	gear, _, noSuffix, err := encodeGear(raw, SuffixTable{"of the Falcon": 14})
	if err != nil {
		t.Fatal(err)
	}
	if gear != "waist=1234::14" {
		t.Fatalf("gear = %q, want waist=1234::14 (item id, no enchant, suffix 14)", gear)
	}
	if len(noSuffix) != 0 {
		t.Fatalf("noSuffix = %v, want none: this item matched", noSuffix)
	}
}

func TestEncodeGearOnlyEmitsThePermanentEnchant(t *testing.T) {
	raw := []byte(`{"equipped_items":[{"item":{"id":999},"slot":{"type":"FEET"},"name":"Test Boots",
		"enchantments":[{"enchantment_id":1,"enchantment_slot":{"type":"TEMPORARY"}},
		                {"enchantment_id":911,"enchantment_slot":{"type":"PERMANENT"}}]}]}`)
	gear, _, _, err := encodeGear(raw, SuffixTable{})
	if err != nil {
		t.Fatal(err)
	}
	if gear != "feet=999:911" {
		t.Fatalf("gear = %q, want feet=999:911 (the PERMANENT enchant only, not the TEMPORARY one)", gear)
	}
}
```

- [ ] **Step 2: Run to verify failure** — FAIL (`encodeGear` undefined).

- [ ] **Step 3: Implement `gear.go`**
```go
// api/internal/bnetbuild/gear.go
package bnetbuild

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type blizzardEquipment struct {
	EquippedItems []blizzardEquippedItem `json:"equipped_items"`
}

type blizzardEquippedItem struct {
	Item struct {
		ID int `json:"id"`
	} `json:"item"`
	Enchantments []struct {
		EnchantmentID   int `json:"enchantment_id"`
		EnchantmentSlot struct {
			Type string `json:"type"`
		} `json:"enchantment_slot"`
	} `json:"enchantments"`
	Slot struct {
		Type string `json:"type"`
	} `json:"slot"`
	Name string `json:"name"`
}

// bnetSlotToOurs maps Blizzard's equipped-item slot.type to this site's SLOTS names (spec
// §2.2). A slot not in this table (SHIRT, TABARD, and anything else the simulator does not
// model) is skipped and named in Report.SkippedSlots.
var bnetSlotToOurs = map[string]string{
	"HEAD": "head", "NECK": "neck", "SHOULDER": "shoulder", "BACK": "back", "CHEST": "chest",
	"WRIST": "wrist", "HANDS": "hands", "WAIST": "waist", "LEGS": "legs", "FEET": "feet",
	"FINGER_1": "finger1", "FINGER_2": "finger2", "TRINKET_1": "trinket1", "TRINKET_2": "trinket2",
	"MAIN_HAND": "main_hand", "OFF_HAND": "off_hand", "RANGED": "ranged",
}

// encodeGear builds the FS1 v1 gear list (spec §2.2, contract 10.5: "<slot>=<item_id>
// [:<enchant[:<suffix>]]", in SLOTS order). suffixes may be empty; a miss is never
// reported (ruling 4 in the plan header — no items.json eligibility table is available to
// distinguish "unknown suffix" from "never had one").
func encodeGear(raw json.RawMessage, suffixes SuffixTable) (gear string, skippedSlots, noSuffix []string, err error) {
	var body blizzardEquipment
	if err := json.Unmarshal(raw, &body); err != nil {
		return "", nil, nil, fmt.Errorf("bnetbuild: decode equipment: %w", err)
	}

	bySlot := make(map[string]blizzardEquippedItem, len(body.EquippedItems))
	for _, it := range body.EquippedItems {
		ourSlot, ok := bnetSlotToOurs[it.Slot.Type]
		if !ok {
			skippedSlots = append(skippedSlots, it.Slot.Type)
			continue
		}
		bySlot[ourSlot] = it
	}

	var parts []string
	for _, slot := range SLOTS {
		it, ok := bySlot[slot]
		if !ok {
			continue
		}
		entry := slot + "=" + strconv.Itoa(it.Item.ID)
		enchant := permanentEnchant(it)
		suffix, matched := matchSuffix(it.Name, suffixes)
		switch {
		case enchant != 0 && matched:
			entry += ":" + strconv.Itoa(enchant) + ":" + strconv.Itoa(suffix)
		case enchant != 0:
			entry += ":" + strconv.Itoa(enchant)
		case matched:
			entry += "::" + strconv.Itoa(suffix)
		}
		parts = append(parts, entry)
	}
	return strings.Join(parts, ","), skippedSlots, noSuffix, nil
}

// permanentEnchant returns the item's PERMANENT-slot enchantment_id, or 0 when it has none
// (spec §2.2: "the enchant is the PERMANENT enchantment_id when present").
func permanentEnchant(it blizzardEquippedItem) int {
	for _, e := range it.Enchantments {
		if e.EnchantmentSlot.Type == "PERMANENT" {
			return e.EnchantmentID
		}
	}
	return 0
}

// matchSuffix finds the longest suffix name in the table that name ends with, preceded by a
// space (ruling 4: an exact trailing match against real suffixes.json strings, not a
// heuristic on "of").
func matchSuffix(name string, suffixes SuffixTable) (id int, matched bool) {
	bestLen := -1
	for suffixName, suffixID := range suffixes {
		trailer := " " + suffixName
		if strings.HasSuffix(name, trailer) && len(suffixName) > bestLen {
			bestLen = len(suffixName)
			id = suffixID
			matched = true
		}
	}
	return id, matched
}
```

- [ ] **Step 4: Run to verify pass**

```bash
cd api && export GOWORK=off && go vet ./internal/bnetbuild/... && go test ./internal/bnetbuild/... -p 1 -v
```
Expected: PASS.

- [ ] **Step 5: Commit** — `feat(bnetbuild): map Blizzard equipment onto the FS1 gear grammar`.

---

## Task 7: `bnetbuild.Encode` — the top-level function, the Report, and the golden fixture

**Files:**
- Create: `api/internal/bnetbuild/encode.go`
- Create: `api/internal/bnetbuild/encode_test.go`
- Create: `api/internal/bnetbuild/testdata/era-kiloz.fs1` (golden file, written in Step 5)

**Interfaces:**
- Consumes: everything from Tasks 4–6, plus `bnetapi.CharacterProfile` (Task 2).
- Produces: `type Inputs struct { Build string; Profile bnetapi.CharacterProfile; Equipment,
  Talents json.RawMessage; Talent TalentTable; Enchants EnchantTable; Suffixes SuffixTable;
  Races RaceTable }`, `type Report struct { UnmatchedTalents, Clamped, NoSuffix,
  SkippedSlots []string }`, `func Encode(in Inputs) (code string, report Report, err
  error)`. This is the spec §2.2 signature verbatim — nothing here may drift from it.

- [ ] **Step 1: Write the encode test (real fixture, error-path tests)**
```go
// api/internal/bnetbuild/encode_test.go
package bnetbuild

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/bnetapi"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

func loadEraKilozInputs(t *testing.T) Inputs {
	t.Helper()
	data, err := trees.Load(realDataDir)
	if err != nil {
		t.Fatal(err)
	}
	tables, ok, err := LoadTables(realDataDir, data)
	if err != nil || !ok {
		t.Fatalf("LoadTables: ok=%v err=%v", ok, err)
	}
	return Inputs{
		Build:     tables.Build,
		Profile:   bnetapi.CharacterProfile{Name: "Kiloz", ClassSlug: "warrior", RaceName: "Orc"},
		Equipment: readEraKiloz(t, "equipment.json"),
		Talents:   readEraKiloz(t, "specializations.json"),
		Talent:    TalentTable{Build: tables.Trees, ClassID: classIDFor(t, tables.Trees, "warrior")},
		Enchants:  tables.Enchants,
		Suffixes:  tables.Suffixes,
		Races:     tables.Races,
	}
}

func classIDFor(t *testing.T, b *trees.Build, slug string) int {
	t.Helper()
	for _, c := range b.Classes() {
		if c.Slug == slug {
			return c.ID
		}
	}
	t.Fatalf("class %q not found", slug)
	return 0
}

func TestEncodeEraKilozHasTheRightHeadAndGear(t *testing.T) {
	code, report, err := Encode(loadEraKilozInputs(t))
	if err != nil {
		t.Fatal(err)
	}
	head := strings.Split(code, ":")
	if len(head) < 6 || head[0] != "FS1" || head[2] != "warrior" || head[3] != "orc" {
		t.Fatalf("code head = %v, want FS1:<build>:warrior:orc:...", head[:min(6, len(head))])
	}
	if !strings.Contains(code, "head=21329:2583") {
		t.Fatalf("code = %q, missing the fixture's head slot", code)
	}
	if len(report.SkippedSlots) != 2 {
		t.Fatalf("SkippedSlots = %v, want exactly SHIRT and TABARD", report.SkippedSlots)
	}
}

func TestEncodeRefusesAnUnknownRace(t *testing.T) {
	in := loadEraKilozInputs(t)
	in.Profile.RaceName = "Ogre"
	if _, _, err := Encode(in); err == nil {
		t.Fatal("an unmapped race must return an error, not a guessed slug")
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```

- [ ] **Step 2: Run to verify failure** — FAIL (`Inputs`/`Encode` undefined).

- [ ] **Step 3: Implement `encode.go`**
```go
// api/internal/bnetbuild/encode.go

// Package bnetbuild turns a Blizzard character profile, equipment and specializations
// response into the FS1 v1 export string the simulator and planner already read (spec
// docs/superpowers/specs/2026-09-22-battlenet-first-design.md §2.2). Every function here is
// pure: no I/O, no database, no HTTP — bnetimport does all three and calls Encode with what
// it captured.
package bnetbuild

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/bnetapi"
)

// Inputs is everything Encode needs, exactly the spec §2.2 signature.
type Inputs struct {
	Build     string
	Profile   bnetapi.CharacterProfile
	Equipment json.RawMessage
	Talents   json.RawMessage
	Talent    TalentTable
	Enchants  EnchantTable
	Suffixes  SuffixTable
	Races     RaceTable
}

// Report is what Encode could not map cleanly — logged at INFO with the character's key by
// the caller (bnetimport), never inside this pure package (spec §2.2).
type Report struct {
	UnmatchedTalents []string
	Clamped          []string
	NoSuffix         []string
	SkippedSlots     []string
}

// Encode builds one FS1 v1 string: "FS1:<build>:<class>:<race>:<t1>/<t2>/<t3>:<gear>" (spec
// §2.2, grammar in docs/superpowers/specs/2026-09-19-simulator-parity-interfaces.md §7,
// §10.5). An unmapped race is the one hard failure — every other mismatch (an unmatched
// talent, a clamped rank, an unresolved suffix, a skipped gear slot) is dropped and named in
// Report instead, per spec §1.3's honesty rule: nothing here is inferred or faked, but a
// partial build is still worth serving.
func Encode(in Inputs) (code string, report Report, err error) {
	raceSlug, ok := in.Races[in.Profile.RaceName]
	if !ok {
		return "", Report{}, fmt.Errorf("bnetbuild: unknown race %q", in.Profile.RaceName)
	}

	treeRanks, unmatched, clamped, err := encodeTalents(in.Talents, in.Talent)
	if err != nil {
		return "", Report{}, err
	}
	gear, skippedSlots, noSuffix, err := encodeGear(in.Equipment, in.Suffixes)
	if err != nil {
		return "", Report{}, err
	}

	trees := make([]string, 3)
	for i := 0; i < 3; i++ {
		trees[i] = encodeTree(treeRanks[i])
	}

	code = strings.Join([]string{
		"FS1", in.Build, in.Profile.ClassSlug, raceSlug, strings.Join(trees, "/"), gear,
	}, ":")
	return code, Report{
		UnmatchedTalents: unmatched, Clamped: clamped, NoSuffix: noSuffix, SkippedSlots: skippedSlots,
	}, nil
}

// encodeTree matches web/src/lib/planner/fs1.ts's encodeTree exactly: one base-36 digit per
// rank in tab order, trailing zeros trimmed, "0" for an empty tree.
func encodeTree(ranks []int) string {
	var b strings.Builder
	for _, r := range ranks {
		if r < 0 {
			r = 0
		}
		if r > 35 {
			r = 35
		}
		b.WriteString(strings.ToLower(fmt.Sprintf("%x", r)))
		if r >= 36 {
			// unreachable: max_rank never exceeds single-digit base-36 in this game, but
			// guards against a future data change silently truncating.
			panic("bnetbuild: talent rank does not fit one base-36 digit")
		}
	}
	trimmed := strings.TrimRight(b.String(), "0")
	if trimmed == "" {
		return "0"
	}
	return trimmed
}
```

- [ ] **Step 4: Run to verify pass**

```bash
cd api && export GOWORK=off && go vet ./internal/bnetbuild/... && go test ./internal/bnetbuild/... -p 1 -v
```
Expected: PASS.

- [ ] **Step 5: Write the golden fixture the web lane's decode test reads**

Write a short throwaway `main.go` under the scratchpad (not committed) that imports
`bnetbuild`, builds `loadEraKilozInputs`-equivalent `Inputs` and prints `Encode`'s `code`;
or add a temporary `t.Log(code)` to `TestEncodeEraKilozHasTheRightHeadAndGear` and run `go
test ./internal/bnetbuild/... -run TestEncodeEraKilozHasTheRightHeadAndGear -v`, copy the
logged string, then:
```bash
printf '%s' '<the printed FS1 string>' > api/internal/bnetbuild/testdata/era-kiloz.fs1
```
Remove the temporary `t.Log`. Then add a byte-exact golden test to `encode_test.go`:
```go
func TestEraKilozFixtureMatchesTheCheckedInFile(t *testing.T) {
	code, _, err := Encode(loadEraKilozInputs(t))
	if err != nil {
		t.Fatal(err)
	}
	want, err := os.ReadFile(filepath.Join("testdata", "era-kiloz.fs1"))
	if err != nil {
		t.Fatal(err)
	}
	if code != string(want) {
		t.Fatalf("Encode(era-kiloz) = %q, want the checked-in fixture %q", code, want)
	}
}
```

- [ ] **Step 6: Run every bnetbuild test once more**

```bash
cd api && export GOWORK=off && go vet ./internal/bnetbuild/... && go test ./internal/bnetbuild/... -p 1 -v
```
Expected: PASS, including the new golden test.

- [ ] **Step 7: Commit** — `feat(bnetbuild): Encode a Blizzard profile into an FS1 build, plus the era-kiloz golden fixture`.

---

## Task 8: `bnetimport` — capture specializations, encode, and write the export

**Files:**
- Modify: `api/internal/bnetimport/capture.go` (`captureEquipment` returns raw;
  `captureSpecializations` new)
- Modify: `api/internal/bnetimport/character.go` (`syncCharacterGuild` returns the typed
  profile too)
- Modify: `api/internal/bnetimport/refresh.go` (thread the new returns through)
- Modify: `api/internal/bnetimport/import.go` (`Service` gains `Tables bnetbuild.Tables`)
- Create: `api/internal/bnetimport/build.go` (`buildAndWriteExport`)
- Modify: `api/internal/bnetimport/import_test.go`, `refresh_test.go`, `harness_test.go` (new
  assertions + fixture wiring)

**Interfaces:**
- Consumes: `bnetbuild.Tables`, `bnetbuild.Inputs`, `bnetbuild.Encode` (Tasks 4–7).
- Produces: `Service.Tables bnetbuild.Tables` (zero value skips encoding entirely — no crash
  when a test or a deployment has no build data).

- [ ] **Step 1: Write the failing test**

Add to `api/internal/bnetimport/import_test.go` (near the existing full-import tests) a test
using `newBlizzardFixture` plus real warrior table data loaded the way `bnetbuild`'s own
tests do:
```go
func TestImportWritesABlizzardSourcedExportWhenAllThreeCapturesAnswer(t *testing.T) {
	pool := testPool(t)
	userID := seedUser(t, pool)
	f := newBlizzardFixture(t)
	f.realms("us", map[string]string{"whitemane": "NORMAL"})
	f.json(http.MethodGet, "/profile/user/wow?namespace=profile-classic1x-us", http.StatusOK,
		map[string]any{"wow_accounts": []map[string]any{{"characters": []map[string]any{{
			"name": "Kiloz", "id": 36745378, "realm": map[string]any{"slug": "whitemane", "name": "Whitemane"},
			"playable_class": map[string]any{"name": "Warrior"}, "playable_race": map[string]any{"name": "Orc"},
			"gender": map[string]any{"type": "MALE"}, "faction": map[string]any{"type": "HORDE"}, "level": 60,
		}}}}})
	profileBody, err := os.ReadFile("../bnetapi/testdata/era-kiloz/profile.json")
	if err != nil {
		t.Fatal(err)
	}
	equipmentBody, err := os.ReadFile("../bnetapi/testdata/era-kiloz/equipment.json")
	if err != nil {
		t.Fatal(err)
	}
	specBody, err := os.ReadFile("../bnetapi/testdata/era-kiloz/specializations.json")
	if err != nil {
		t.Fatal(err)
	}
	f.handlers[http.MethodGet+" /profile/wow/character/whitemane/kiloz?namespace=profile-classic1x-us"] = func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(profileBody)
	}
	f.handlers[http.MethodGet+" /profile/wow/character/whitemane/kiloz/equipment?namespace=profile-classic1x-us"] = func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(equipmentBody)
	}
	f.handlers[http.MethodGet+" /profile/wow/character/whitemane/kiloz/specializations?namespace=profile-classic1x-us"] = func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(specBody)
	}
	f.handlers[http.MethodGet+" /profile/wow/character/whitemane/kiloz/character-media?namespace=profile-classic1x-us"] = func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}

	data, err := trees.Load("../../../data/builds")
	if err != nil {
		t.Fatal(err)
	}
	tables, ok, err := bnetbuild.LoadTables("../../../data/builds", data)
	if err != nil || !ok {
		t.Fatalf("LoadTables: ok=%v err=%v", ok, err)
	}
	svc := newTestService(t, pool, f)
	svc.Tables = tables

	if _, err := svc.ImportAccount(context.Background(), userID, "user-tok"); err != nil {
		t.Fatal(err)
	}

	var source string
	var export string
	if err := pool.QueryRow(context.Background(),
		`select source, export from addon_exports where character_key = 'us/normal/kiloz'`).
		Scan(&source, &export); err != nil {
		t.Fatal(err)
	}
	if source != "blizzard" {
		t.Fatalf("source = %q, want blizzard", source)
	}
	if !strings.HasPrefix(export, "FS1:") || !strings.Contains(export, ":warrior:orc:") {
		t.Fatalf("export = %q, want an FS1 warrior/orc build", export)
	}
}
```
(Add `"os"`, `"strings"`, `"github.com/jhunthrop/foreversixty/api/internal/bnetbuild"`,
`"github.com/jhunthrop/foreversixty/api/internal/trees"` to the test file's imports.)

- [ ] **Step 2: Run to verify failure** — `go test ./internal/bnetimport/... -run TestImportWritesABlizzardSourcedExportWhenAllThreeCapturesAnswer -v` → FAIL (no `Service.Tables` field, no write).

- [ ] **Step 3: Implement**

In `capture.go`, change `captureEquipment` to return the raw bytes too, and add
`captureSpecializations`:
```go
// captureEquipment now returns the raw body alongside writing it, so the caller can feed it
// straight to bnetbuild.Encode without a second Blizzard call (spec §2.3).
func (s *Service) captureEquipment(ctx context.Context, tx pgx.Tx, key, region, realmSlug, name string) (json.RawMessage, error) {
	raw, err := s.Client.Equipment(ctx, region, realmSlug, name)
	if err != nil {
		if errors.Is(err, bnetapi.ErrNotFound) || errors.Is(err, bnetapi.ErrForbidden) {
			return nil, nil
		}
		return nil, fmt.Errorf("bnetimport: equipment %s: %w", key, err)
	}
	capped := s.capCapture("bnet_equipment", key, raw)
	if _, err := tx.Exec(ctx,
		`update characters set bnet_equipment = coalesce($2, bnet_equipment), bnet_captured_at = now() where key = $1`,
		key, rawOrNil(capped)); err != nil {
		return nil, fmt.Errorf("bnetimport: capture equipment %s: %w", key, err)
	}
	return raw, nil
}

// captureSpecializations fetches and stores a character's talent groups (spec §2.3, same
// coalesce-on-partial-data and 403/404-is-not-an-error shape as captureEquipment), and
// returns the raw body for bnetbuild.Encode.
func (s *Service) captureSpecializations(ctx context.Context, tx pgx.Tx, key, region, realmSlug, name string) (json.RawMessage, error) {
	raw, err := s.Client.Specializations(ctx, region, realmSlug, name)
	if err != nil {
		if errors.Is(err, bnetapi.ErrNotFound) || errors.Is(err, bnetapi.ErrForbidden) {
			return nil, nil
		}
		return nil, fmt.Errorf("bnetimport: specializations %s: %w", key, err)
	}
	capped := s.capCapture("bnet_talents", key, raw)
	if _, err := tx.Exec(ctx,
		`update characters set bnet_talents = coalesce($2, bnet_talents), bnet_captured_at = now() where key = $1`,
		key, rawOrNil(capped)); err != nil {
		return nil, fmt.Errorf("bnetimport: capture specializations %s: %w", key, err)
	}
	return raw, nil
}
```

In `character.go`, change `syncCharacterGuild`'s signature to also return the profile:
```go
func (s *Service) syncCharacterGuild(ctx context.Context, tx pgx.Tx, userID int64, region, ruleset, key, realmSlug, name string,
	rosterCache map[string]bnetapi.Roster) (profile bnetapi.CharacterProfile, guildWritten, unavailable bool, err error) {
```
Update its body: every `return false, false, nil` / `return false, false, err` becomes
`return bnetapi.CharacterProfile{}, false, false, nil` / `..., err`, and every
`return true, false, nil` / `return false, true, nil` (etc.) gains the leading
`profile`/`bnetapi.CharacterProfile{}` value matching what was fetched — after
`profile, rawProfile, err := s.Client.Character(...)` and its 404 branch, subsequent success
paths return the real `profile`; the 404/403 branches return the zero value. (Every existing
return statement in the function gets one new leading value; the logic is unchanged.) At the
very end, the two success returns become `return profile, true, false, nil` and
`return profile, false, false, nil` (unguilded).

In `character.go`'s `importOneCharacter`, after the guild sync call, thread the new returns
through and call the encoder:
```go
	profile, guildWritten, unavailable, err := s.syncCharacterGuild(ctx, tx, userID, region, ruleset, key, ch.RealmSlug, ch.Name, rosterCache)
	if err != nil {
		return false, false, false, err
	}
	rawEquipment, err := s.captureEquipment(ctx, tx, key, region, ch.RealmSlug, ch.Name)
	if err != nil {
		return false, false, false, err
	}
	rawSpecializations, err := s.captureSpecializations(ctx, tx, key, region, ch.RealmSlug, ch.Name)
	if err != nil {
		return false, false, false, err
	}
	if err := s.captureMedia(ctx, tx, key, region, ch.RealmSlug, ch.Name); err != nil {
		return false, false, false, err
	}
	if err := s.buildAndWriteExport(ctx, tx, userID, key, region, ruleset, profile, rawEquipment, rawSpecializations); err != nil {
		return false, false, false, err
	}
```

In `refresh.go`'s `refreshOneCharacter`, the same threading:
```go
	profile, _, _, err := s.syncCharacterGuild(ctx, tx, sc.UserID, region, ruleset, key, sc.RealmSlug, name, rosterCache)
	if err != nil {
		return err
	}
	rawEquipment, err := s.captureEquipment(ctx, tx, key, region, sc.RealmSlug, name)
	if err != nil {
		return err
	}
	rawSpecializations, err := s.captureSpecializations(ctx, tx, key, region, sc.RealmSlug, name)
	if err != nil {
		return err
	}
	if err := s.captureMedia(ctx, tx, key, region, sc.RealmSlug, name); err != nil {
		return err
	}
	if err := s.buildAndWriteExport(ctx, tx, sc.UserID, key, region, ruleset, profile, rawEquipment, rawSpecializations); err != nil {
		return err
	}
```

In `import.go`, add the field to `Service`:
```go
type Service struct {
	Pool    *pgxpool.Pool
	Client  *bnetapi.Client
	Regions []string
	Log     *slog.Logger
	// Tables is the active client build's talent/enchant/suffix/race data (spec §2.2,
	// §2.3), loaded once per process by cmd/api's wiring. A zero Tables (Trees nil) means
	// no build data is available and buildAndWriteExport is a no-op — never a crash.
	Tables bnetbuild.Tables
}
```

Create `build.go`:
```go
// api/internal/bnetimport/build.go
package bnetimport

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/bnetapi"
	"github.com/jhunthrop/foreversixty/api/internal/bnetbuild"
)

// buildAndWriteExport runs bnetbuild.Encode over a character's freshly-captured profile,
// equipment and specializations (spec §2.3) and, on success, upserts addon_exports with the
// newest-captured_at-wins rule (spec §2.4). It is a no-op — never an error — whenever any of
// the three captures came back empty (a private profile, a 403/404 equipment or
// specializations read — a Season of Discovery character has none of them) or this process
// holds no active-build data tables: "all three answered" is the gate, and honesty (spec
// §1.3) means a character this cannot build for keeps whatever export it already had.
func (s *Service) buildAndWriteExport(ctx context.Context, tx pgx.Tx, userID int64, key, region, ruleset string,
	profile bnetapi.CharacterProfile, rawEquipment, rawSpecializations json.RawMessage) error {
	if rawEquipment == nil || rawSpecializations == nil || s.Tables.Trees == nil {
		return nil
	}
	classID, ok := classIDFor(s.Tables.Trees, profile.ClassSlug)
	if !ok {
		return nil
	}
	code, report, err := bnetbuild.Encode(bnetbuild.Inputs{
		Build:     s.Tables.Build,
		Profile:   profile,
		Equipment: rawEquipment,
		Talents:   rawSpecializations,
		Talent:    bnetbuild.TalentTable{Build: s.Tables.Trees, ClassID: classID},
		Enchants:  s.Tables.Enchants,
		Suffixes:  s.Tables.Suffixes,
		Races:     s.Tables.Races,
	})
	if err != nil {
		s.logger().Warn("bnetimport", "op", "build_encode", "key", key, "err", err)
		return nil
	}

	capturedAt := time.Now()
	if profile.LastLoginTimestamp > 0 {
		capturedAt = time.UnixMilli(profile.LastLoginTimestamp)
	}
	if _, err := tx.Exec(ctx,
		`insert into addon_exports (character_key, user_id, region, ruleset, name, export, source, captured_at, updated_at)
		 values ($1, $2, $3, $4, $5, $6, 'blizzard', $7, now())
		 on conflict (character_key) do update set
		   user_id = excluded.user_id, region = excluded.region, ruleset = excluded.ruleset, name = excluded.name,
		   export = excluded.export, source = 'blizzard', captured_at = excluded.captured_at, updated_at = now()
		 where addon_exports.user_id = excluded.user_id
		   and (addon_exports.source = 'blizzard' or excluded.captured_at > addon_exports.captured_at)`,
		key, userID, region, ruleset, profile.Name, code, capturedAt); err != nil {
		return fmt.Errorf("bnetimport: write export %s: %w", key, err)
	}
	s.logger().Info("bnetimport", "op", "build_encoded", "key", key,
		"unmatched_talents", len(report.UnmatchedTalents), "clamped", len(report.Clamped),
		"no_suffix", len(report.NoSuffix), "skipped_slots", len(report.SkippedSlots))
	return nil
}

// classIDFor resolves a class slug to its id within the active build's data, the one lookup
// bnetbuild.TalentTable needs that trees.Build indexes by id rather than slug.
func classIDFor(b interface{ Classes() []classRow }, slug string) (int, bool) {
	for _, c := range b.Classes() {
		if c.Slug == slug {
			return c.ID, true
		}
	}
	return 0, false
}
```

**Note for the implementer:** the `classIDFor` signature above using a structural interface
(`interface{ Classes() []classRow }`) is over-engineered for one call site with one real
caller — replace it with the concrete type directly:
```go
func classIDFor(b *trees.Build, slug string) (int, bool) {
	for _, c := range b.Classes() {
		if c.Slug == slug {
			return c.ID, true
		}
	}
	return 0, false
}
```
and add `"github.com/jhunthrop/foreversixty/api/internal/trees"` to `build.go`'s imports.

- [ ] **Step 4: Fix every other call site the two signature changes touch**

`grep -rn "syncCharacterGuild\|captureEquipment(" internal/bnetimport/*.go` and update each
one (character.go's two call sites are done above; check `refresh_test.go` and
`import_test.go` for any direct calls — expected: none, both are unexported and called only
from within the package's own non-test files).

- [ ] **Step 5: Run to verify pass**

```bash
cd api && export GOWORK=off && go vet ./internal/bnetimport/... && go test ./internal/bnetimport/... -p 1 -v
```
Expected: PASS, including the new test. Run the full previously-passing suite too (no
regressions from the signature changes):
```bash
go test ./internal/bnetimport/... -p 1
```

- [ ] **Step 6: Commit** — `feat(bnetimport): capture specializations and write a Blizzard-sourced FS1 build`.

---

## Task 9: `addon.Store.putOneExport` — stamp `source='addon'`, `captured_at=now()`

**Files:**
- Modify: `api/internal/addon/addon.go`
- Modify: `api/internal/addon/addon_test.go` (add one assertion; existing tests must still
  pass unmodified since `source`/`captured_at` are additive columns)

**Interfaces:**
- Produces: no signature change — `putOneExport`'s SQL only.

- [ ] **Step 1: Write the failing test** — add to `addon_test.go` near the existing
`PutExports` happy-path test:
```go
func TestPutExportsStampsSourceAddonAndCapturedAt(t *testing.T) {
	s := testStore(t)
	userID := seedUser(t, s.Pool)
	if err := s.PutExports(context.Background(), userID, []Export{
		{Name: "Kiloz", Ruleset: "normal", Region: "us", Export: "FS1:1.60.1.69893:warrior:orc:0/0/0:"},
	}); err != nil {
		t.Fatal(err)
	}
	var source string
	var capturedAt time.Time
	if err := s.Pool.QueryRow(context.Background(),
		`select source, captured_at from addon_exports where character_key = 'us/normal/kiloz'`).
		Scan(&source, &capturedAt); err != nil {
		t.Fatal(err)
	}
	if source != "addon" {
		t.Fatalf("source = %q, want addon", source)
	}
	if time.Since(capturedAt) > time.Minute {
		t.Fatalf("captured_at = %v, want stamped to roughly now", capturedAt)
	}
}
```
(Add `"time"` to the test file's imports if not already present.)

- [ ] **Step 2: Run to verify failure** — FAIL: `source`/`captured_at` columns exist (Task
1's migration) but this insert doesn't set them, so `source` defaults to `'addon'` already
(this part may already pass) while `captured_at` has no default and the insert as written
would actually fail with a NOT NULL violation — confirming the change is required.

- [ ] **Step 3: Implement** — in `addon.go`'s `putOneExport`:
```go
func (s *Store) putOneExport(ctx context.Context, tx pgx.Tx, userID int64, key, region, ruleset string, e Export) error {
	tag, err := tx.Exec(ctx,
		`insert into addon_exports (character_key, user_id, region, ruleset, name, export, source, captured_at, updated_at)
		 values ($1, $2, $3, $4, $5, $6, 'addon', now(), now())
		 on conflict (character_key) do update set
		   user_id = excluded.user_id, export = excluded.export,
		   source = 'addon', captured_at = now(), updated_at = now()
		 where addon_exports.user_id = excluded.user_id`,
		key, userID, region, ruleset, e.Name, e.Export)
```
(everything below `tag, err :=` is unchanged).

- [ ] **Step 4: Run to verify pass**

```bash
cd api && export GOWORK=off && go vet ./internal/addon/... && go test ./internal/addon/... -p 1
```
Expected: PASS, no regressions.

- [ ] **Step 5: Commit** — `feat(addon): stamp addon-sourced exports with source and captured_at`.

---

## Task 10: `sims.Store.SimInput` — the `blizzard` source

**Files:**
- Modify: `api/internal/sims/input.go`
- Modify: `api/internal/sims/input_test.go`

**Interfaces:**
- Produces: `Input.Source` may now be `"addon"`, `"blizzard"`, or `"fight"`. Selection
  between the export row and the last fight uses `captured_at` (spec §2.5), not
  `updated_at`.

- [ ] **Step 1: Write the failing test** — add to `input_test.go`:
```go
func TestSimInputReportsBlizzardSource(t *testing.T) {
	h := newInputHarness(t)
	defer h.close()
	h.insertExportWithSource(t, "us", "pvp", "kiloz", "FS1:1.60.1.69893:warrior:orc:0/0/0:", "blizzard", pastTime(1))
	in, _, ok, err := h.store.SimInput(context.Background(), "us/pvp/kiloz")
	if err != nil || !ok {
		t.Fatalf("ok=%v err=%v", ok, err)
	}
	if in.Source != "blizzard" {
		t.Fatalf("Source = %q, want blizzard", in.Source)
	}
}
```
Check `input_test.go`'s existing helper names (`newInputHarness`/`pastTime`/`insertExport`)
and match them exactly — if the harness type or helper names differ from this sketch, use
the real ones and add a `source` parameter to the existing `insertExport` helper (defaulting
existing call sites to `"addon"`) rather than adding a parallel helper, to avoid duplicating
the insert SQL.

- [ ] **Step 2: Run to verify failure** — FAIL (no `source` column read, `Source` always
`"addon"` today).

- [ ] **Step 3: Implement** — in `input.go`'s `SimInput`:
```go
	var (
		ref        FightRef
		export     *string
		exportSrc  *string
		exportAt   *time.Time
		...
	)
	err := s.Pool.QueryRow(ctx,
		`select export, source, captured_at from addon_exports where character_key = $1`, key).
		Scan(&export, &exportSrc, &exportAt)
```
and in the `switch`:
```go
	switch {
	case exportAt != nil && (fightAt == nil || exportAt.After(*fightAt)):
		out.Source, out.CapturedAt = *exportSrc, *exportAt
```
(the rest of that branch — the JSON-string quoting of `*export` — is unchanged).

- [ ] **Step 4: Run to verify pass**

```bash
cd api && export GOWORK=off && go vet ./internal/sims/... && go test ./internal/sims/... -p 1
```
Expected: PASS.

- [ ] **Step 5: Commit** — `feat(sims): sim-input reports the blizzard source and uses captured_at`.

---

## Task 11: `openapi.yaml` — `sim-input`'s `source` enum

**Files:**
- Modify: `api/openapi.yaml`

- [ ] **Step 1:** in the `sim-input` response schema (around line 1387), change:
```yaml
                          source: { type: string, enum: [addon, fight] }
```
to:
```yaml
                          source: { type: string, enum: [addon, blizzard, fight] }
```

- [ ] **Step 2:** validate the file parses (no schema linter is wired in this repo per the
earlier scan; a YAML syntax check is enough):
```bash
cd api && python3 -c "import yaml; yaml.safe_load(open('openapi.yaml'))" && echo OK
```
Expected: `OK`.

- [ ] **Step 3: Commit** — `docs(api): sim-input's source enum gains blizzard`.

---

## Task 12: `auth.Store` — `Character.Build` from `addon_exports`

**Files:**
- Modify: `api/internal/auth/store.go`
- Modify: `api/internal/auth/store_test.go`

**Interfaces:**
- Produces: `type CharacterBuild struct { Source string; CapturedAt time.Time }`,
  `Character.Build *CharacterBuild` (`json:"build,omitempty"`), matching spec §5's frozen
  `/v1/me` contract: `"build": { "source": "blizzard", "captured_at": "..." }`, omitted
  entirely when the character has no `addon_exports` row.

- [ ] **Step 1: Write the failing test** — add to `store_test.go`:
```go
func TestCharactersIncludesBuildFromAddonExports(t *testing.T) {
	s := testStore(t)
	ctx := context.Background()
	u, err := s.UpsertEmailUser(ctx, "built@example.com")
	if err != nil {
		t.Fatal(err)
	}
	if err := s.LinkCharacter(ctx, u.ID, Character{
		Key: "us/pvp/kiloz", Region: "us", Ruleset: "pvp", Name: "Kiloz", Class: "warrior",
	}); err != nil {
		t.Fatal(err)
	}
	capturedAt := time.Now().Add(-2 * 24 * time.Hour).Truncate(time.Second)
	if _, err := s.Pool.Exec(ctx,
		`insert into addon_exports (character_key, user_id, region, ruleset, name, export, source, captured_at, updated_at)
		 values ('us/pvp/kiloz', $1, 'us', 'pvp', 'Kiloz', 'FS1:1.60.1.69893:warrior:orc:0/0/0:', 'blizzard', $2, now())`,
		u.ID, capturedAt); err != nil {
		t.Fatal(err)
	}
	if err := s.LinkCharacter(ctx, u.ID, Character{
		Key: "us/pvp/nobuild", Region: "us", Ruleset: "pvp", Name: "Nobuild", Class: "mage",
	}); err != nil {
		t.Fatal(err)
	}

	chars, err := s.Characters(ctx, u.ID)
	if err != nil || len(chars) != 2 {
		t.Fatalf("characters = %v, err = %v", chars, err)
	}
	byKey := map[string]Character{}
	for _, c := range chars {
		byKey[c.Key] = c
	}
	built := byKey["us/pvp/kiloz"]
	if built.Build == nil || built.Build.Source != "blizzard" {
		t.Fatalf("built.Build = %+v, want source blizzard", built.Build)
	}
	if !built.Build.CapturedAt.Equal(capturedAt) {
		t.Fatalf("CapturedAt = %v, want %v", built.Build.CapturedAt, capturedAt)
	}
	if byKey["us/pvp/nobuild"].Build != nil {
		t.Fatal("a character with no addon_exports row must have Build == nil (omitted from JSON)")
	}
}
```

- [ ] **Step 2: Run to verify failure** — FAIL (`Character.Build` undefined).

- [ ] **Step 3: Implement** — in `store.go`, add the type and field:
```go
// CharacterBuild is which path most recently produced a simmable build for this character,
// and when (spec docs/superpowers/specs/2026-09-22-battlenet-first-design.md §2.5, §5).
type CharacterBuild struct {
	Source     string    `json:"source"`
	CapturedAt time.Time `json:"captured_at"`
}

type Character struct {
	... // unchanged fields
	Source string `json:"source"`
	// Build is this character's addon_exports row, omitted when it has none (spec §2.5:
	// "so the account page and the simulator's landing know which characters are simmable
	// without one request per character").
	Build *CharacterBuild `json:"build,omitempty"`
	Guild *CharacterGuild `json:"guild,omitempty"`
}
```
Add `time` to imports if not already present (it likely already is, for other timestamp
fields — check before adding a duplicate import).

Extend the shared query and scan:
```go
const characterColumns = `c.key, c.region, c.ruleset, c.name, coalesce(c.class, ''),
	        coalesce(c.realm_name, ''), c.level, coalesce(c.faction, ''), c.source,
	        coalesce(c.race, ''), coalesce(c.gender, ''), c.equipped_item_level,
	        c.avatar_url, c.render_url,
	        ae.source, ae.captured_at,
	        g.id, g.name, gc.rank, gc.rank_index, gc.verified_at is not null`

const characterFrom = `from characters c
	 left join addon_exports ae on ae.character_key = c.key
	 left join guild_characters gc on gc.character_key = c.key
	 left join guilds g on g.id = gc.guild_id`

func scanCharacterRows(rows pgx.Rows) ([]Character, error) {
	defer rows.Close()
	out := []Character{}
	for rows.Next() {
		var c Character
		var buildSource *string
		var buildCapturedAt *time.Time
		var guildID *int64
		var guildName, rank *string
		var rankIndex *int
		var verified bool
		if err := rows.Scan(&c.Key, &c.Region, &c.Ruleset, &c.Name, &c.Class,
			&c.Realm, &c.Level, &c.Faction, &c.Source,
			&c.Race, &c.Gender, &c.ItemLevel,
			&c.AvatarURL, &c.RenderURL,
			&buildSource, &buildCapturedAt,
			&guildID, &guildName, &rank, &rankIndex, &verified); err != nil {
			return nil, fmt.Errorf("auth: scan characters: %w", err)
		}
		if buildSource != nil && buildCapturedAt != nil {
			c.Build = &CharacterBuild{Source: *buildSource, CapturedAt: *buildCapturedAt}
		}
		if guildID != nil {
			c.Guild = &CharacterGuild{
				ID: *guildID, Name: *guildName, Rank: *rank, RankIndex: rankIndex, Verified: verified,
			}
		}
		out = append(out, c)
	}
	return out, rows.Err()
}
```

- [ ] **Step 4: Run to verify pass**

```bash
cd api && export GOWORK=off && go vet ./internal/auth/... && go test ./internal/auth/... -p 1
```
Expected: PASS, no regressions (the join is a `left join`, so a character with no export row
is unaffected — same row count as before).

- [ ] **Step 5: Commit** — `feat(auth): /v1/me characters carry their build source and capture time`.

---

## Task 13: `cmd/api/main.go` — wire `bnetbuild.LoadTables` into both Service constructions

**Files:**
- Modify: `api/cmd/api/main.go`

- [ ] **Step 1:** in `runBnetRefresh` (spec §2.6, the nightly job), load tree data and the
Battle.net build tables before constructing the service:
```go
func runBnetRefresh(ctx context.Context, log *slog.Logger) error {
	cfg, pool, err := start(ctx)
	if err != nil {
		return err
	}
	defer pool.Close()
	if !cfg.BattleNetConfigured() {
		return fmt.Errorf("%s needs Battle.net credentials", bnetimport.RefreshJobCommand)
	}
	treeData, err := trees.Load(cfg.TreeDataDir)
	if err != nil {
		return fmt.Errorf("trees: %s: %w", cfg.TreeDataDir, err)
	}
	tables, ok, err := bnetbuild.LoadTables(cfg.TreeDataDir, treeData)
	if err != nil {
		return fmt.Errorf("bnetbuild: load tables: %w", err)
	}
	if !ok {
		log.Warn(bnetimport.RefreshJobCommand, "err", "no client build available",
			"effect", "characters are re-synced but no Blizzard-sourced build is (re-)encoded")
	}
	svc := &bnetimport.Service{Pool: pool, Client: bnetClient(cfg, log), Regions: cfg.BnetRegions, Log: log, Tables: tables}
	result, err := svc.RunRefresh(ctx, cfg.BnetProbeGames)
	...
```

- [ ] **Step 2:** in the server-startup block (around line 490, where `accounts.Importer` is
built and `treeData` is already in scope from the earlier `trees.Load` call at ~line 438),
wire the same tables:
```go
	if cfg.BattleNetConfigured() {
		accounts.BNet = auth.NewBattleNet(cfg.BnetClientID, cfg.BnetClientSecret, cfg.BnetRedirectURL)
		bnetSvcClient := bnetClient(cfg, log)
		bnetTables, ok, err := bnetbuild.LoadTables(cfg.TreeDataDir, treeData)
		if err != nil {
			return fmt.Errorf("bnetbuild: load tables: %w", err)
		}
		if !ok {
			log.Warn("auth", "err", "no client build available",
				"effect", "Battle.net sign-in works but no character gets a Blizzard-sourced build")
		}
		accounts.Importer = &bnetimport.Service{
			Pool: pool, Client: bnetSvcClient, Regions: cfg.BnetRegions, Log: log, Tables: bnetTables,
		}
		go bnetSvcClient.WarmRealms(context.Background(), cfg.BnetRegions)
	} else {
		log.Warn("auth", "state", "battle.net is not configured", "effect", "email sign-in only")
	}
```

- [ ] **Step 3:** add the `bnetbuild` import to `main.go`'s import block if not already
present.

- [ ] **Step 4: Build and vet the whole module (this file touches the top-level binary)**

```bash
cd api && export GOWORK=off && go build ./... && go vet ./cmd/... ./internal/bnetimport/...
```
Expected: builds clean.

- [ ] **Step 5: Commit** — `feat(api): wire bnetbuild's data tables into the Battle.net import and refresh jobs`.

---

## Task 14: Final whole-branch review and full scoped test run

**Files:** none (verification only).

- [ ] **Step 1: Full build and vet**

```bash
cd api && export GOWORK=off
go build ./...
go vet ./...
```
Expected: clean.

- [ ] **Step 2: Full test run of every touched package**

```bash
cd api && export GOWORK=off TEST_DATABASE_URL=postgres://forever:forever@localhost:5434/forever_test?sslmode=disable
go test -p 1 ./internal/bnetbuild/... ./internal/bnetapi/... ./internal/bnetimport/... \
  ./internal/addon/... ./internal/sims/... ./internal/auth/... ./internal/trees/... \
  ./internal/guilds/... ./internal/db/... ./cmd/...
```
Expected: all PASS. Record the pass/fail counts in the ledger and the final report.

- [ ] **Step 3: Dispatch one fresh `sonnet` review subagent** over the full diff
(`git diff main...HEAD` in the worktree) against this plan's Global Constraints and the
spec's §1–§2/§5, per the lane's `subagent-driven-development` final-review step. Apply one
fix wave for anything CRITICAL or HIGH; log MEDIUM/LOW in the ledger as accepted debt or
fixed.

- [ ] **Step 4:** re-run Step 2 after any fixes; commit the fix wave separately if non-empty
(`fix(bnetbuild): <what the review caught>` or similar, one commit per logical fix).

- [ ] **Step 5:** write the lane's final report per `.superpowers/journeys/lane-common-go.md`
(head sha; what shipped per spec section; test results; every ruling; anything undone; the
exported types/signatures another lane builds against verbatim).
