# Battle.net first: sign in, and your characters are ready to sim

**Date:** 2026-09-22
**Status:** APPROVED by the owner ("hyper optimize this to be the best possible user
experience and the simplest path of least resistance"; "optimize the site itself for this
workflow"). Two lanes: `bnet-first-api` (section 2) and `bnet-first-web` (sections 3 and 4).
Section 1 binds both; section 5 is the frozen contract between them.

## 0. The path today, and the path after

Today: hear about the site → sign in → install the addon → log in with the character →
type `/fs export` → copy → paste on the site → name the character → sim. Six steps, one
context switch into the game, and the drop-off is at the export.

After: hear about the site → **sign in with Battle.net** → your hub, with your characters
ready for the planner, the simulator, logs, rankings and your guild. One required step. The
addon and the companion become upgrades offered where they matter, never prerequisites.

What makes this possible (verified 2026-09-22 on a Classic Era character, fixture at
`api/internal/bnetapi/testdata/era-kiloz/`): Blizzard's profile API serves, for a signed-in
account's character, every equipped item with its `enchantment_id`, and every talent with
`talent.id`, `spell_tooltip.spell.id` and `talent_rank`, in two specialization groups with
`is_active`. Our data keys talents by Blizzard's talent id and carries each talent's
`spell_id` (`data/builds/<build>/talents/<class>.json`), and our enchants are keyed by the
same `SpellItemEnchantment` id Blizzard sends (`data/builds/<build>/enchants.json`). So a
Blizzard profile maps onto the export string the simulator and planner already read,
without translation tables. Professions, bags and bank are not exposed; those stay the
addon's contribution. Season of Discovery characters have none of it (Blizzard serves no
sub-endpoints for them); Era characters have all of it; Forever's namespace is expected to
behave like Era's.

## 1. Global constraints (both lanes)

1. **Vocabulary.** The addon runs in the game; the companion is the desktop app. A
   character whose gear came from Blizzard is described as "from Battle.net", never "from
   the armory" in copy (the code's `armory` source name is internal).
2. **Voice and design.** `design/DESIGN-SYSTEM.md`: reference, not pitch, even on the home
   page. Every visible string in a copy module. The states model of
   `2026-09-22-account-and-island-states-design.md` applies to anything new.
3. **Honesty.** A character is shown as simmable only when the site holds a build for it
   (an export or a Blizzard build). Nothing is inferred or faked; a Season of Discovery
   character stays "no build yet" with the paste fallback.
4. **No token persistence** (unchanged from the import spec).
5. **Lane ownership.** `bnet-first-api` owns `api/**` and the OpenAPI file. `bnet-first-web`
   owns `web/**`. The contract in section 5 is frozen; the web lane stubs it.
6. Lane house rules (`.superpowers/journeys/lane-common-*.md`): sonnet only, never opus,
   including the final review; commit rules; scoped checks; no broad `pkill`; stop servers.

## 2. Lane `bnet-first-api`: a Blizzard profile becomes a build

### 2.1 Migration `0027_export_sources`

```sql
alter table addon_exports
  add column if not exists source text not null default 'addon' check (source in ('addon', 'blizzard')),
  add column if not exists captured_at timestamptz;
update addon_exports set captured_at = updated_at where captured_at is null;
alter table addon_exports alter column captured_at set not null;
```

`captured_at` is when the build was true in the game: the upload time for an addon export
(the addon wrote it moments before), Blizzard's `last_login_timestamp` for a Blizzard build
(Blizzard's profile reflects the character as of its last session). `updated_at` stays the
row's write time. The primary key stays `character_key`: one build per character, the
freshest by `captured_at` wins (2.4).

### 2.2 `bnetbuild` package: profile + equipment + specializations → FS1

`api/internal/bnetbuild`, pure functions over the fixture shapes, no I/O:

```go
type Inputs struct {
  Build      string              // data build, e.g. "1.60.1.69893" (the site's active build)
  Profile    bnetapi.CharacterProfile
  Equipment  json.RawMessage     // /equipment
  Talents    json.RawMessage     // /specializations
  Talent     TalentTable         // loaded once per class from data/builds/<build>/talents/<class>.json
  Enchants   EnchantTable        // ids from enchants.json
  Suffixes   SuffixTable         // name → id from suffixes.json
  Races      RaceTable           // Blizzard race name → our race slug, from races.json
}
func Encode(in Inputs) (code string, report Report, err error)
```

- Head: `FS1:<Build>:<classSlug>:<raceSlug>:<t1>/<t2>/<t3>:<gear>`, exactly the grammar in
  `docs/superpowers/specs/2026-09-19-simulator-parity-interfaces.md` §7 and §10.5. Talent
  ranks are one base-36 digit per talent in each tree's tab order, trailing zeros trimmed,
  `0` for an empty tree. The **active** specialization group is used; when none is active,
  the first. A Blizzard talent is matched by `talent.id` against our talent `id`, falling
  back to `spell_tooltip.spell.id` against `spell_id`/`ranks[].spell_id`; an unmatched
  talent is dropped and named in `Report.UnmatchedTalents`. A rank above `max_rank` is
  clamped and reported.
- Gear: `<slot>=<item_id>[:<enchant>[:<suffix>]]` in `SLOTS` order
  (`web/src/lib/planner/types.ts`; copy the list into Go with a test that reads the TS file
  and compares). Blizzard slot types map to our slots (`HEAD→head`, `FINGER_1→finger1`,
  `TRINKET_1→trinket1`, `MAIN_HAND→main_hand`, `OFF_HAND→off_hand`, `RANGED→ranged`, and so
  on; tabard, shirt and other slots the simulator does not model are skipped). The
  enchant is the PERMANENT `enchantment_id` when present. The suffix is found by matching
  the item's `name` against `<base name> of the <suffix name>` from `suffixes.json` when
  the item has suffixes in `items.json`; no match → no suffix, reported.
- The class slug comes from `character_class.name` lowercased; the race slug from the race
  table (`Night Elf → night-elf`); an unknown race → `err`.
- `Report{UnmatchedTalents []string; Clamped []string; NoSuffix []string; SkippedSlots []string}`
  is logged at INFO with the key.
- Tests: the Era fixture encodes to a string that `web`'s decoder accepts. Add a Go test
  that shells out to `node` only if the lane finds a light way (a `vitest` test in
  `web/src/lib/planner/fs1.blizzard.test.ts` that decodes a checked-in expected string is
  acceptable: the lane writes `api/internal/bnetbuild/testdata/era-kiloz.fs1` and the web
  lane's test reads it). Table tests for every mapping rule above.

### 2.3 Capture: at import and refresh

`bnetimport` already fetches profile and equipment; it now also fetches `/specializations`
(stored raw as `bnet_talents jsonb`, same cap and stamp as the other captures, migration
0027 adds the column) and, when all three answered, runs `bnetbuild.Encode` and writes an
`addon_exports` row `{character_key, user_id, region, ruleset, name, export, source:
'blizzard', captured_at: last_login_at, updated_at: now()}` through the same ownership
guard as the addon path. Data tables are loaded once per process from
`data/builds/<active build>/` (the API image already ships `/data`; see
`dataaddon`'s build discovery for how the active build is found).

### 2.4 Newest wins

`putOneExport` (the addon and paste path) upserts with `source='addon', captured_at=now()`.
The Blizzard path upserts **only when** its `captured_at` is newer than the row's
current `captured_at` or the row is absent or its source is `blizzard`
(`on conflict ... where addon_exports.source = 'blizzard' or excluded.captured_at > addon_exports.captured_at`).
So an addon export taken after the character's last Blizzard session keeps winning until
the player logs in again, and a Blizzard build never overwrites a fresher addon export.
Guild sync (`syncGuild`) runs only for addon-sourced writes; the Blizzard path already
writes membership from the roster.

### 2.5 `sim-input` and `/v1/me`

`GET /v1/characters/{region}/{ruleset}/{name}/sim-input` gains `source: "blizzard"` in its
enum (OpenAPI too); its `captured_at` is the row's `captured_at`; selection between the
export row and the last ranked fight stays "newest wins" using `captured_at`.

`GET /v1/me` `characters[]` gains `build: {source: "addon"|"blizzard", captured_at}`
(omitted when no row), so the account page and the simulator's landing know which
characters are simmable without one request per character.

### 2.6 Nightly refresh

`bnet-refresh` already re-captures; with 2.3 it re-encodes every night, so a character
that played yesterday has today's gear without touching the site. The refresh logs
`op=build_encoded` per character with the report counts.

## 3. Lane `bnet-first-web`: the site around the new path

### 3.1 Sign in lands on your hub

The site is the planner, the simulator, logs, rankings, guilds and the reference together;
sims are one of six things a signed-in player does. So sign-in lands on **the account
page as the player's hub**, not on any one tool.

- `/login` and every sign-in link default `next` to `/account?signed_in=1` (the header's
  Sign in link, `SignInPrompt`, the home panel).
- On arrival with `signed_in=1`, the account page picks the **main character** (the
  simmable character with the latest `build.captured_at`; ties by level; none simmable →
  the highest level; none at all → nothing), sets the current-character pointer to it so
  every tool opens with it, and shows it in the hero band with the full set of actions:
  `Open in simulator`, `Open in planner`, `Logs` (`/logs`), and the character's own page.
  A one-line banner reads "Signed in. <Name> is your current character; change it from
  any row below." and the URL is cleaned with `history.replaceState`. No simmable
  character → the hero band still shows the main character, with its character page link,
  `Logs`, and the paste fallback in place of the simulator and planner actions, plus the
  sentence that Blizzard serves no data for this realm type when no character has a build.
- The hub's main column, in order: the hero band; Characters (rows carry the build source
  pill `Battle.net · 2 days ago` / `Addon · today` / `No build yet`); **Your ratings**: the
  main character's latest rating card summary with a link to its character page (the
  existing `CharacterRatingPanel`, reused, only when a rating exists); Your reports. The
  rail stays: Devices, You, Guilds and plan; the Guilds panel links to the guild home.
- The simulator's own landing keeps its "Your characters" list with the source pills;
  nothing is auto-loaded there beyond the current-character pointer it already honours.

### 3.2 The home page leads with the path

Signed out, the hero keeps the search (it is the reference site's first control) and gains,
above the tool cards, one panel: the sentence "Sign in with Battle.net and your characters
arrive with their gear, talents and guild: plan, sim, log and rank them from here." with
the Sign in button (`next=/account?signed_in=1`). Reference, not pitch: one sentence, one
button, no list of benefits.

Signed in, that panel becomes the current character strip: the chip, `Open in planner`,
`Open in simulator`, `Logs`, and "Your characters" linking to `/account`.

### 3.3 The addon and the companion become upgrades, in context

- **Simulator gear tab** and Top Gear: when the loaded character's source is `blizzard`,
  the bags/bank slot shows "Bags and bank come from the addon. Install it to include them."
  with the link to `/addon`, replacing `bagsNeedAddon`'s current wording.
- **Account page Characters intro** becomes: "Gear and talents come from Battle.net and
  refresh nightly. Install the addon to include bags and bank and to update right after a
  session; paste an export for a character Battle.net has no data for." with links.
- **Logs page**: "Upload a log" stays first; "Log live with the companion" stays the second
  panel. No change beyond checking the copy never calls the companion the addon.
- **The addon page** (`/addon`): the opening line states what the addon adds on top of
  Battle.net (bags, bank, professions, the in-game build guide, instant refresh) and keeps
  the paste box for the fallback case.
- **Stale copy removed**: `landingSourceNote`, `armoryNotYet`, `armorySignIn`,
  `noCharactersYet`, `sourceAddonBody`, `sourceAccountTitle` and every other string that
  says Blizzard has no profile API or that the addon is required for a sim are rewritten
  to the new truth. The lane greps `web/src` for "no character profile", "armory", "Install
  the addon" and reviews every hit.

### 3.4 Hand-off links and the character page

`CharacterHandoffLinks` and the account rows use `build` from `/v1/me` where available (no
per-row `sim-input` lookup on the account page); the character page keeps its own lookup.
`fromStoredCharacter` and `lookupAddonExport` treat `source === 'blizzard'` exactly like
`addon` (an FS1 string). The source pill copy (`sourcePill`) gains the Battle.net case.

### 3.5 Tests

Vitest: main-character selection (`lib/account/main-character.ts`, pure), the hub's
signed-in arrival, the home panel both ways, copy changes; the FS1 decode test over the API
lane's checked-in `era-kiloz.fs1` (skip with a clear message if the file is absent on this
branch). Playwright: sign-in → hub with the main character in the hero band and the pointer set,
then `Open in simulator` loading it, with `/v1/me` and `sim-input` stubbed (`source:
"blizzard"`), the home panel signed out and signed in. `npm run lhci`
before the final report with CLS per URL (the home page is in the strict bucket).

## 4. Order of work in the web lane

1. Contract types and normalisers (5).
2. The hub's signed-in arrival and the main-character rule (`lib/account/main-character.ts`).
3. Sign-in `next` defaults.
4. Home panel.
5. Copy pass (3.3) and hand-offs (3.4).
6. Tests and Lighthouse.

## 5. The frozen contract

`GET /v1/me` character:
```json
{ "key": "us/pvp/kiloz", "name": "Kiloz", "class": "warrior", "level": 60, ...,
  "build": { "source": "blizzard", "captured_at": "2026-09-21T03:14:00Z" } }
```
`build` omitted when the site holds no build for the character.

`GET /v1/characters/{region}/{ruleset}/{name}/sim-input`:
```json
{ "spec": "fury", "gear": "FS1:1.60.1.69893:warrior:orc:...", "talents": "t1/t2/t3",
  "buffs": [], "captured_at": "2026-09-21T03:14:00Z", "source": "blizzard" }
```
`source` is `addon`, `blizzard` or `fight`; for `addon` and `blizzard`, `gear` is the FS1
string.

## 6. Out of scope

Professions from Blizzard (not exposed); Season of Discovery data (Blizzard serves none);
class icon sets; the companion's own onboarding; Forever's namespace (one config change
when it appears).
