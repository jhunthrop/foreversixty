# Guild Raids and Sign-ups (Lane G1) Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Officers plan raids on the guild's Raids tab; members sign up in one tap from the raid page or the guild band; an officer's post can go to the guild's Discord channel.

**Architecture:** Two new tables (`guild_raids`, `guild_raid_signups`) and `guilds.raid_timezone` in migration 0050, served by new files in `api/internal/guilds` (store, view, handlers) and read by three web surfaces: the Raids tab's `Planned` section with the create form, the raid page `GuildRaid.svelte` under the guild shell, and the band line from `home.next_raid`. Discord posting is PR G1b: a guild-scoped webhook (`guild_notification_webhooks`, migration 0051) connected from guild settings, two embeds, posted from the raid handlers through the existing `notify` poster.

**Tech Stack:** Go (pgx; goldmark + bluemonday for notes), Astro + Svelte 5 runes, vitest, Playwright.

**Spec:** `design/specs/2026-10-26-guild-raids.md` (experience, binding for the web; its §9 contract is the API's contract) and `docs/superpowers/specs/2026-10-26-guild-suite-design.md` Lane G1 (product).

## Global Constraints

- Stored times are UTC `timestamptz`; `guilds.raid_timezone` is an IANA name; the API renders nothing in a zone (the web does), except the Discord `Guild time` field and the copied list.
- Visibility (spec R1): planned raids and the raid route exist only for verified members and officers of the guild; everyone else gets 404 `not_found` from the raid route and `planned: []` from the list.
- Every officer write is behind `CanForGuild(guildID, userID, FeatureOfficerViews)` after the rank check, answering 403 `guild_plan_required` `Planning and running raids are part of the guild plan.`; member sign-ups are free.
- One character per account at `in` or `late` (R3), enforced in the store's transaction.
- At most 12 `planned`/`locked` raids per guild (`A guild plans at most 12 raids ahead.`, 422 `raid_cap`).
- Notes: markdown subset (paragraphs, bold, italic, lists, links), sanitised server-side; `notes_html` is the only form the web renders; never raw HTML.
- Copy: every new string in `web/src/lib/guild/raid-copy.ts`, tested verbatim against spec §6.1 through `web/src/test-support/spec-table.ts`.
- Commit trailer `Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>`; `git commit -m`, never a heredoc; API tests with `-p 1`.

## Review Focus

1. A member who owns two characters in the guild sets the second to `in`: the first must move to `bench` atomically, and a concurrent second request must not leave both `in` (Task 1's transaction test with two goroutines).
2. A raid at 20:00 `America/Chicago` repeated weekly across 14 March 2027 (US DST start) keeps 20:00 guild wall time; the stored instants differ by 1 h that week (Task 3's repeat test and Task 5's `forYouLine` test).
3. An officer locks the raid between a member's page load and their tap: the PUT answers 409 `raid_locked`, the web reverts the segment, shows the race line focused, and renders the locked state without a reload (Task 3 + Task 8).
4. A signed-out visitor or an unverified member requests `GET /raids/{id}`: 404, never 403, and `GET /raids` carries no `planned` key for them (Task 3).
5. Notes containing `<script>`, an image, or a `javascript:` link render as text or are dropped; a plain URL becomes a link (Task 2).

---

## PR G1a — raids and sign-ups

### Task 1: Zones, migration 0050, the raids store

**Files:**
- Create: `data/curated/raids/first-tier.json`, `api/internal/guilds/raids_zones.go`, `api/internal/db/migrations/0050_guild_raids.up.sql`, `0050_guild_raids.down.sql`, `api/internal/guilds/raids_store.go`
- Test: `api/internal/guilds/raids_store_test.go`, `api/internal/guilds/raids_zones_test.go`

**Interfaces:**
- Produces: `type Zone struct{ Key, Name string; Size *int; LogZone *string }`, `func Zones() []Zone` (the curated order), `func ZoneByKey(key string) (Zone, bool)`.
- Produces: `type Raid struct{ ID int64; GuildID int64; Zone string; Title string; StartsAt time.Time; DurationMinutes int; Notes string; Status string; ReportID *string; PostDefault bool; CreatedBy int64; CreatedAt, UpdatedAt time.Time }`, statuses `RaidPlanned|RaidLocked|RaidDone|RaidCancelled`.
- Produces: `type Signup struct{ RaidID int64; CharacterKey, Name, Class, Spec, Role, State, Note string; SignedBy int64; UpdatedAt time.Time }`, roles `tank|healer|dps`, states `in|late|bench|out`.
- Produces (Store): `CreateRaid(ctx, in RaidInput) (Raid, error)` (ErrRaidCap when 12 planned/locked exist), `UpdateRaid(ctx, id int64, in RaidInput) (Raid, error)`, `SetRaidStatus(ctx, id int64, status string) (Raid, error)`, `GetRaid(ctx, guildID, id int64) (Raid, bool, error)`, `PlannedRaids(ctx, guildID int64, now time.Time) ([]Raid, error)` (status planned/locked/cancelled with `starts_at + duration + 6h > now`, by starts_at), `NextRaid(ctx, guildID int64, now time.Time) (Raid, bool, error)` (first planned/locked ending after now), `UpsertSignup(ctx, s Signup) (moved *string, err error)` (R3: returns the name of the own character moved to bench, or nil; ErrRaidLocked / ErrRaidCancelled from the raid's status read in the same tx), `RaidSignups(ctx, raidID int64) ([]Signup, error)`, `SetGuildRaidTimezone(ctx, guildID int64, tz string) error`, `GuildRaidTimezone(ctx, guildID int64) (string, error)`.

- [ ] **Step 1: Write the curated zones file**

`data/curated/raids/first-tier.json`:
```json
{
  "notes": "The first raid tier (opens 9 December 2026). size and log_zone are sourced facts: Onyxia's Lair is 40 and the combat log names it \"Onyxia's Lair\"; the other two stay null until the client data names them (tenet 8).",
  "zones": [
    { "key": "raid:barrow-deeps", "name": "Barrow Deeps", "size": null, "log_zone": null },
    { "key": "raid:hyjal-summit", "name": "Hyjal Summit", "size": null, "log_zone": null },
    { "key": "raid:onyxias-lair", "name": "Onyxia's Lair", "size": 40, "log_zone": "Onyxia's Lair" }
  ]
}
```
`raids_zones.go` embeds it (`//go:embed` via a copy the build already makes for `data/curated`? — check how `loot.go` reads `data/builds/.../loot.json` (`TREE_DATA_DIR`) and follow the same loader; if curated files are not shipped in the image, embed the JSON with `go:embed` from a package-local copy `api/internal/guilds/first-tier.json` kept in sync by a test that compares it to the data file).

- [ ] **Step 2: Write the migration**

```sql
alter table guilds add column raid_timezone text;
create table guild_raids (
  id bigserial primary key,
  guild_id bigint not null references guilds(id) on delete cascade,
  zone text not null,
  title text not null check (char_length(title) between 1 and 80),
  starts_at timestamptz not null,
  duration_minutes int not null check (duration_minutes between 60 and 240),
  notes text not null default '' check (char_length(notes) <= 2000),
  status text not null default 'planned' check (status in ('planned','locked','done','cancelled')),
  report_id text references reports(id) on delete set null,
  post_default boolean not null default false,
  created_by bigint not null references users(id) on delete restrict,
  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now()
);
create index guild_raids_guild_starts_idx on guild_raids (guild_id, starts_at);
create table guild_raid_signups (
  raid_id bigint not null references guild_raids(id) on delete cascade,
  character_key text not null,
  name text not null, class text not null, spec text not null,
  role text not null check (role in ('tank','healer','dps')),
  state text not null check (state in ('in','late','bench','out')),
  note text not null default '' check (char_length(note) <= 140),
  signed_by bigint not null references users(id) on delete restrict,
  updated_at timestamptz not null default now(),
  primary key (raid_id, character_key)
);
```
Down: drop the two tables, drop the column. (`reports(id)` is text — confirm against `0005_logs`.)

- [ ] **Step 3: Failing store tests** (DB harness from `harness_test.go`)
```go
func TestZonesAreTheCuratedThree(t *testing.T)
func TestCreateRaidRefusesTheThirteenth(t *testing.T)          // 12 planned → ErrRaidCap; a cancelled one does not count
func TestPlannedRaidsWindow(t *testing.T)                       // a raid 7 h past its end is gone; cancelled stays until its start
func TestUpsertSignupMovesTheOtherOwnCharacterToBench(t *testing.T) // two chars of one user: second `in` → first `bench`, moved == first's name; `out` moves nothing
func TestUpsertSignupIsAtomicUnderRace(t *testing.T)            // two goroutines set each char `in`; after both, exactly one is `in`
func TestUpsertSignupRefusesLockedAndCancelled(t *testing.T)    // ErrRaidLocked / ErrRaidCancelled
func TestNextRaidSkipsEnded(t *testing.T)
```
- [ ] **Step 4: Run, see FAIL.** `cd api && TEST_DATABASE_URL=… go test -race -p 1 -run 'Raid|Zones' ./internal/guilds/`
- [ ] **Step 5: Implement** (small functions; R3 as `select … for update` on the user's signups in this raid, then the two writes).
- [ ] **Step 6: Run → PASS; `go vet ./...`; gofmt.**
- [ ] **Step 7: Commit** — `feat(api): guild raids and sign-ups store, zones, migration 0050`

### Task 2: Notes rendering

**Files:**
- Create: `api/internal/guilds/raids_notes.go`
- Test: `api/internal/guilds/raids_notes_test.go`
- Modify: `api/go.mod` (add `github.com/yuin/goldmark` and `github.com/microcosm-cc/bluemonday`, pinned current versions; `go mod tidy`)

**Interfaces:**
- Produces: `func RenderNotes(markdown string) string` — goldmark (CommonMark, no raw HTML passthrough, `Linkify` extension on) then bluemonday with a policy allowing only `p, strong, em, ul, ol, li, a[href rel target]` with `http`/`https` hrefs, `rel="nofollow noopener"`, `target="_blank"`.

- [ ] **Step 1: Failing table test** — `**bold**` → `<p><strong>bold</strong></p>`; `- a\n- b` → a `<ul>`; `<script>x</script>` → escaped text, no tag; `![x](u)` → no `<img>`; `[x](javascript:alert(1))` → the text without an href; `https://foreversixty.gg` → an `<a>`; 2 000 chars renders under 10 ms.
- [ ] **Step 2–4: red, implement, green.**
- [ ] **Step 5: Commit** — `feat(api): raid notes rendered from a markdown subset and sanitised`

### Task 3: Raid views and handlers

**Files:**
- Create: `api/internal/guilds/raids_view.go`, `api/internal/guilds/raids_handler.go`, `api/internal/guilds/raids_signup_handler.go`
- Modify: `api/internal/guilds/raids.go` (`RaidsPage` gains `Planned []RaidView`, `Zones []ZoneView`, `GuildTimezone string` for verified members only; absent keys otherwise — `omitempty` on the three), `api/internal/guilds/handler.go` (routes), `api/internal/guilds/home.go` (`HomeView.NextRaid *NextRaidView` with `omitempty`)
- Test: `api/internal/guilds/raids_handler_test.go`, `raids_signup_handler_test.go`, `home_test.go` (or the existing home tests' file)

**Interfaces:**
- Produces JSON (spec §9 contract, verbatim keys): `RaidView{id, zone, zone_name, title, starts_at, duration_minutes, status, size, guild_timezone, notes_html, counts{in,late,bench,out,tanks,healers,dps}, signups[{character_key,name,class,spec,role,state,note,own}], not_signed[{character_key,name,class,spec}], viewer{characters[{character_key,name,class,spec,verified,state,role}]}, discord{connected,post_default}}` (`discord.connected` false until G1b), `ZoneView{key,name,size}`, `NextRaidView{id, zone, zone_name, title, starts_at, size, guild_timezone, counts{in,late,tanks,healers,dps}, own{character_key,state,role}|null}`.
- Routes: `GET /v1/guilds/{id}/raids/{raid}` (session optional like `/raids`; 404 unless verified member), `POST /v1/guilds/{id}/raids` (body `{zone, title, starts_at, duration_minutes, notes, timezone, repeat_weeks (0|3), post_to_discord}`; 201 `{raids: [RaidView…]}`), `PATCH /v1/guilds/{id}/raids/{raid}` (same fields minus repeat; 200), `POST …/raids/{raid}/lock|unlock|cancel` (200 RaidView), `PUT …/raids/{raid}/signups/{character_key}` (body `{state, role, note}`; 200 `{raid: RaidView, moved: name|null}`; 409 `raid_locked`/`raid_cancelled`; 403 `forbidden` when the character is not the caller's and the caller is not an officer; 422 `invalid` on a bad state/role/note). Rate limit the PUT at 30/min per user.
- Errors: 403 `guild_plan_required` (officer writes without the feature), 422 `raid_cap`, 422 `invalid` with field names (`zone`, `starts_at`, `title`, `timezone`, `duration_minutes`).
- `starts_at` arrives as RFC 3339 (the web converts the guild wall clock with the zone before sending); `timezone` sets `guilds.raid_timezone` when it changes (officer only).

- [ ] **Step 1: Failing handler tests** (table-driven over the harness): the Review Focus items 3 and 4; officer creates with `repeat_weeks: 3` at 20:00 Chicago starting 2027-03-07 → four raids whose wall time in Chicago is 20:00 each (the 14 March one is 01:00 UTC later); cap → 422 `raid_cap`; member PUT own character `in` with a second own character already `in` → `moved` names it; member PUT another member's character → 403 `forbidden`; officer PUT anyone → 200; lock then member PUT → 409 `raid_locked`; officer without the feature → 403 `guild_plan_required` on POST/PATCH/lock/cancel, and PUT on their own character still 200; `GET /raids` for a member carries `planned`, `zones`, `guild_timezone`; for a visitor none of the three keys; `home.next_raid` present for a member with a planned raid, `own` null when unsigned, absent when none.
- [ ] **Step 2: Run, see FAIL.**
- [ ] **Step 3: Implement** — views built from the store rows plus `RenderNotes`; the character list from `guild_characters` for the caller (verified only can sign; unverified listed with `verified:false`); counts in Go; `not_signed` = verified roster rows without a signup.
- [ ] **Step 4: `go test -race -p 1 ./internal/guilds/` → PASS; vet; gofmt.**
- [ ] **Step 5: Commit** — `feat(api): planned raids, the raid page view, sign-ups and the band's next_raid`

### Task 4: Web contracts, copy and time helpers

**Files:**
- Create: `web/src/lib/guild/raid-copy.ts`, `web/src/lib/guild/raid-time.ts`, tests for both; `web/src/lib/guild/raids-api.ts` (the fetch/mutation calls: `fetchGuildRaid`, `createRaid`, `updateRaid`, `setRaidStatus`, `putSignup`, with the `GuildApiError.code` path)
- Modify: `web/src/lib/guild/api.ts` (types `GuildRaidView`, `GuildZoneView`, `NextRaidView`; `GuildRaidsPage` gains the optional `planned`, `zones`, `guild_timezone`; `GuildHome.next_raid?`), `web/src/lib/characters.ts` (`parseGuildRaidPath(pathname): { guild: CharacterPath; raidId: string } | null` next to `parseGuildSettingsPath`)
- Test: `web/src/lib/guild/raid-copy.test.ts`, `raid-time.test.ts`, `raids-api.test.ts`, `web/src/lib/characters.test.ts` (extend)

**Interfaces:**
- `raidCopy` (every §6.1 string as constants or functions: `bandLine(args)`, `whenPhrase(date, now, locale, guildTz)`, `countdown(startsAt, endsAt, now)`, `countLine(n, size)`, `roleBreakdown(t,h,d)`, `guildTimeLabel(startsAt, tz)`, `showAll(n)`, `copyList(view, tz)`), tested verbatim via `spec-table.ts` (`## 6. Copy, verbatim` / `### 6.1 On the page`).
- `raid-time.ts`: `wallClockToInstant(date, time, tz) → Date` and `instantToWallClock(d, tz)` (Intl-based, no library), `forYouLine(date, time, guildTz, viewerTz)`, `offsetsDiffer(d, tzA, tzB)`, `nextWeekday(lastStart, now)`.

- [ ] **Step 1: Failing tests** — the spec strings; `whenPhrase` cases (`Tonight 8 pm`, `Tomorrow 8 pm`, `Thursday 8 pm`, `Thu 17 Dec, 8 pm`, ` (9 pm guild time)` when offsets differ, `8:30 pm`); `countdown` cases incl. `starting now`, `ended 2 h ago`; `wallClockToInstant('2027-03-14','20:00','America/Chicago')` → `2027-03-15T01:00:00Z` and `2027-03-07` → `02:00Z`; `forYouLine` with viewer `America/New_York` → `For you: Thu 10 Dec, 9:00 pm.`, same zone → `For you: same time.`; `parseGuildRaidPath('/guild/us/pvp/olympus-xxvii/raids/12')` → `{…, raidId: '12'}`, `/raids/abc` → null, `/raids` → null.
- [ ] **Step 2–4: red, implement, green;** `npm run check`, eslint, prettier.
- [ ] **Step 5: Commit** — `feat(web): raid copy, time helpers, raid API calls and the raid path`

### Task 5: Sign-up controls

**Files:**
- Create: `web/src/components/guild/SignupSegments.svelte`, `web/src/components/guild/RaidSignupStrip.svelte`, `web/src/components/guild/RaidChip.svelte`
- Test: `web/src/components/guild/SignupSegments.test.ts`, `RaidSignupStrip.test.ts`, `RaidChip.test.ts`

**Interfaces:**
- `SignupSegments` props `{ value: 'in'|'late'|'bench'|'out'|null, name: string, size: 'band'|'strip', disabled?: boolean, describedBy?: string, onChange: (state) => void }` — `role="radiogroup"`, `aria-label` `{name}'s sign-up`, four 44 px segments (36 px in `band` from `md`), arrow-key movement, Space/Enter select, classes per spec §3.3.
- `RaidSignupStrip` props `{ raid: GuildRaidView, locked: boolean, ended: boolean, cancelled: boolean, onSignup: (characterKey, state, role, note) => Promise<{ ok: true; moved: string|null } | { ok: false; code: string; message: string }> }`; per-row status line (`Saved.` 4 s, `{Name} moved to Bench.`, the race lines with `role=alert tabindex=-1` focused).
- `RaidChip` props `{ kind: 'status'|'viewer', status?, state?, role? }` → the §3.6 table.

- [ ] **Step 1: Failing component tests** (server render): four segments with the right checked classes per value; `disabled` + `aria-describedby` propagate; the strip renders one row per character in the header-first order, the unverified row muted with the waiting sentence and no controls, the note counter at 120+, the locked line above the rows; chips text and classes per §3.6.
- [ ] **Step 2–4: red, implement, green;** check/eslint/prettier.
- [ ] **Step 5: Commit** — `feat(web): raid sign-up segments, the sign-up strip and the raid chips`

### Task 6: The Raids tab's Planned section and the form

**Files:**
- Create: `web/src/components/guild/RaidForm.svelte`, `web/src/components/guild/PlannedRaidRow.svelte`, `web/src/lib/guild/raid-form.ts` (pure: defaults from the last raid, validation returning the §6.1 field errors, the repeat helper text under the cap, the payload builder calling `wallClockToInstant`)
- Modify: `web/src/components/guild/GuildRaids.svelte` (header row with `Plan a raid`, the form slot, `Planned` with show-all, `Raid nights` sub-head; props gain `planned`, `zones`, `guildTimezone`, `officer`, `officerTools`, `onPlan`, `settingsHref`), `web/src/components/Guild.svelte` (pass the new props; `officerTools` from the existing derivation; the create call through `raids-api.ts`; refresh the raids read after a create), `web/src/lib/guild/copy.ts` only if a shared string moves
- Test: `web/src/lib/guild/raid-form.test.ts`, `web/src/components/guild/RaidForm.test.ts`, `PlannedRaidRow.test.ts`, `GuildRaids.test.ts` (extend or create)

- [ ] **Step 1: Failing tests** — form defaults (last zone/weekday/time/duration; first raid `20:00`, no date, `Choose a raid`); validation messages; `For you` line; the cap helper `Only 2 more fit…`; the Discord line when `discord.connected` is false (no checkbox); planned row cells at desktop and phone (date block, title/zone, when + guild time only when offsets differ, count/bar or `23 signed up` without a size, chips); `Show all 7 planned raids` after four; the empty lines (member vs officer); officer-read-only: `Plan a raid` disabled with the gate sentence `tab = raids` (reuse `GuildPlanGate.svelte` with a `sentence` prop); at 12 planned the cap line under the row and the button disabled.
- [ ] **Step 2–4: red, implement, green;** check/eslint/prettier.
- [ ] **Step 5: Commit** — `feat(web): planned raids on the Raids tab, Plan a raid and the raid form`

### Task 7: The raid page

**Files:**
- Create: `web/src/components/GuildRaid.svelte` (the page: band, officer row, tabs, strip, notes, roster; ≤ 400 lines — split `guild/RaidRoster.svelte`, `guild/RaidOfficerRow.svelte` (incl. the cancel confirm and `Copy the sign-up list`), `web/src/lib/guild/raid-page-controller.svelte.ts` (load, 10 s poll while visible and before end + 6 h, optimistic sign-up, the race handling, the three-failure pause line))
- Modify: `web/src/components/GuildShell.svelte` (route `parseGuildRaidPath` before `guildPath`), `web/src/lib/guild/api.ts` cache invalidation for the raid key
- Test: `web/src/components/GuildRaid.test.ts`, `guild/RaidRoster.test.ts`, `guild/RaidOfficerRow.test.ts`, `web/src/lib/guild/raid-page-controller.test.ts` (fake timers: poll cadence, pause after three failures, visibility gating, optimistic row kept until its own answer)

- [ ] **Step 1: Failing tests** — band lines (when line, guild time only when offsets differ, countdown, count line, status chip); officer row buttons and their disabled set in read-only with `Copy` live; the cancel confirm (focus on `Keep it`, Escape keeps); roster grouping/order (role → class order → name), the Late/Bench/Out/Not signed up groups, `No tanks yet.`, officer inline select with the `—` first option; notes empty lines; page states (locked line + disabled segments, ended, cancelled, done with `Open the night`); the not-visible page (R1) for signed-out and unverified; `copyList` text in guild time; the loading reserve from `localStorage` key `fs.guildRaid.rows` (wrapped in try/catch) — one row when absent.
- [ ] **Step 2–4: red, implement, green;** check/eslint/prettier.
- [ ] **Step 5: Commit** — `feat(web): the raid page: band, sign-up strip, notes, roster, officer row, live poll`

### Task 8: The band line and the needs line

**Files:**
- Modify: `web/src/components/Guild.svelte` (the raid line between the standing line and the needs line, from `home.next_raid`; a press calls `putSignup` for the current character and rewrites the sentence in place; the needs line's weekday), `web/src/lib/guild/copy.ts` (the needs-line weekday function if it lives there)
- Test: `web/src/components/Guild.test.ts` (extend; or a small `guild/RaidBandLine.svelte` with its own test if Guild.svelte would pass 1 100 lines — prefer the extraction)

- [ ] **Step 1: Failing tests** — sentence variants from §6.1 (In/Late/Bench/Out/unsigned/unknown size/two own characters), the none-planned lines (member vs officer with `Plan a raid` → `#raids`), the link `{Zone} →`, the segments at 36 px from `md`, the needs line `before Thursday` / `before the next raid`.
- [ ] **Step 2–4: red, implement, green.**
- [ ] **Step 5: Commit** — `feat(web): the next-raid line in the guild band`

### Task 9: e2e, captures, reviews, changelog

**Files:**
- Create: `web/tests/e2e/guild-raids.spec.ts`, `web/tests/e2e/support/guild-raids.ts` (stubs: a raid view builder, the routes for `/raids`, `/raids/{id}`, PUT signups, POST raids), `web/tests/captures/guild-raids/{capture.spec.ts,playwright.capture.config.ts}` (port 4441), `design/reviews/2026-10-26-guild-raids/README.md`, `web/src/content/changelog/2026-10-27-guild-raids.md` (the §6.1 changelog line)
- Modify: `web/src/fixtures/guild/mock-guild.ts` (`buildMockHome` option `nextRaid`), `web/tests/e2e/layout.spec.ts` only if the teaser collides

- [ ] **Step 1: Write the e2e** per spec §9 (every bullet), desktop + mobile.
- [ ] **Step 2: Run** against a fixture build with the full guild set: `E2E_SKIP_BUILD=1 E2E_PORT=4442 npx playwright test tests/e2e/guild-raids.spec.ts tests/e2e/guild-centre.spec.ts tests/e2e/guild-officer-tools.spec.ts tests/e2e/layout.spec.ts`.
- [ ] **Step 3: Captures** — spec §9's table (1–16, 18); heights annotated against §7's caps; the README with heights and "Noticed while capturing".
- [ ] **Step 4: Reviews** — ux-designer and wow-player on the PNGs; one fix round; reports in the reviews folder. (Capture 17, the live Discord clients, is PR G1b's.)
- [ ] **Step 5: Commit** — `test(web): guild raids e2e, captures, reviews and changelog`

**PR G1a gate:** full web vitest + full Playwright + lhci (GuildShell routing and the band change the guild page) → CI green → fast-forward main → ordered deploy (api with 0050 first, then web) → verify: `schema_migrations` 50; `GET /v1/guilds/1/raids` anonymously carries no `planned`; the live guild bundle carries `Plan a raid` → memory.

---

## PR G1b — posting raids to Discord

### Task 10: Spec amendment and the guild webhook

**Files:**
- Spec: the ux-designer adds `§3.8 Guild settings · Discord` to `design/specs/2026-10-26-guild-raids.md` (the connect block for a guild: `Connect a Discord channel` → the guild connect route, the connected line `Posts raids to a Discord channel · Channel ….`, `Send test message`, `Remove`; officer only; mirrors the account block's connected/paste states but without events) before this task's implementer starts.
- Create: `api/internal/db/migrations/0051_guild_notification_webhooks.{up,down}.sql` (`guild_notification_webhooks`: guild_id pk → guilds, token_box bytea, channel_id text, discord_guild_id text, source text, failures int, disabled_reason text, last_error text, connected_by bigint, created_at, updated_at; `notify_connect_states`: state text pk, user_id, guild_id, expires_at), `api/internal/notify/guild_store.go`, `api/internal/notify/guild_connect_handler.go`, `api/internal/notify/guild_handler.go` (test/remove), web: `web/src/components/guild/GuildDiscordSection.svelte` in settings, `web/src/lib/guild/settings-view.ts` words
- Routes: `GET /v1/guilds/{id}/notifications/discord/connect` (officer; writes a `notify_connect_states` row; 302 to Discord with the state), the shared callback reads the state row (user match, not expired, officer rank re-checked) and saves to the guild table; `POST …/test`, `DELETE …`; guild settings `GET` gains `discord{connected, channel_id, disabled_reason}`.
- Tests: store, handlers (state misuse → `failed`), settings e2e.

- [ ] Steps as the Discord connect lane's Tasks (reuse `Sealer`, `Poster`, `Lookup`); commit `feat: guild Discord channel connect, test and remove`.

### Task 11: The two embeds and posting from the raid handlers

**Files:**
- Create: `api/internal/notify/raid_embed.go` (`RaidPostedEmbed(publicBaseURL, in RaidPosted)`, `RaidChangedEmbed(…)` for `moved` and `cancelled` per spec §6.2: Discord timestamps `<t:unix:F> (<t:unix:R>)`, `Guild time` in the guild's zone, thumbnail the crest when absolute), `api/internal/notify/guild_deliver.go` (`GuildDeliverer.RaidPosted/RaidChanged` with the five-failure rule on the guild row)
- Modify: `api/internal/guilds/raids_handler.go` (post when `post_to_discord` is true on create — one embed per raid, soonest first — on a time change, and on cancel with `post_cancellation`; remember `post_default` on the guild), `raids_view.go` (`discord.connected`, `post_default`)
- Tests: embed goldens; handler posts through a recording poster; no post on a title-only edit.

- [ ] Steps: red/green/commit `feat: raid_posted and raid_changed Discord embeds posted from the raid handlers`.
- [ ] Capture 17 by hand in a staging channel; wow-player checks the embed.

**PR G1b gate:** api tests, guild settings e2e, CI, ordered deploy (0051), verify the connect route 302s for an officer and 404s for a member, memory.
