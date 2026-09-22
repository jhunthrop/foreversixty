# Lane `account-visual` plan

## Global Constraints (copied from the brief)

- Binding spec: `.superpowers/account-visual-brief.md`; design source of truth
  `design/DESIGN-SYSTEM.md` (Character row / State panel definitions, Color, Typography,
  Spacing). Vocabulary: addon in game, companion on the desktop. Motion budget: nothing new
  animates beyond `.reveal`. Lighthouse CLS budget on every page.
- House rules: `.superpowers/journeys/lane-common-web.md`, `lane-common-go.md` — sonnet
  only, never opus; commit rules (printf message file, `git commit -F`); scoped checks
  (`npx vitest run`, `npx astro check`, `npm run lint`, `npx prettier --check`, Go
  `go vet ./... && go test ./...`); never a broad `pkill`.
- File ownership: `api/**` (migration 0026, `bnetapi`, `bnetimport`, `auth/store.go`,
  `rankings/query.go`); `web/**` limited to `Account.svelte`, `components/account/**`,
  `Character.svelte` (render only), `components/ui/StatePanel.svelte` (new, the one
  exception), `lib/account/**`, fixtures under `web/src/fixtures/me-bnet.ts` and two new
  tiny local images under `web/public/fixtures/`. Never touch `Header.astro`,
  `Footer.astro`, other `components/ui/**`/`lib/ui/**` files.
- `TEST_DATABASE_URL=postgres://forever:forever@localhost:5434/forever_test?sslmode=disable`,
  `export GOWORK=off` in `api/`. `E2E_PORT=4451`.
- Ruling (process): given the 90-minute lane budget, tasks below are implemented directly
  by the lane controller (self-reviewed against the review package listed per task) rather
  than dispatched to separate implementer/reviewer subagents — logged in
  `.superpowers/sdd/2026-09-22-account-visual/progress.md` as Ruling 0.

## Tasks

### A1. Migration 0026_character_media
Add `avatar_url text`, `render_url text`, `bnet_media jsonb` to `characters`, `if not
exists`, matching 0024/0025's style. Up + down. Run `go run ./cmd/migrate` or rely on
`db.Migrate` in tests. Review: matches sibling migrations' column style; down drops in
reverse order.

### A2. `bnetapi.Client.CharacterMedia`
New file `api/internal/bnetapi/media.go`: `CharacterMedia(ctx, region, realmSlug, name)
(Media, json.RawMessage, error)`, `Media{AvatarURL, RenderURL string}` from `assets[].key
in {avatar, main-raw}`, falling back to `main` for the render when `main-raw` is absent.
Any asset URL not under `https://render.worldofwarcraft.com/` is dropped and logged
(`op=media_host_rejected`) rather than returned. Table-driven test `media_test.go`:
avatar+main-raw present; main-raw absent falls back to main; untrusted host dropped and
logged; 404 propagates as `ErrNotFound`.

### A3. `bnetimport` capture + wiring
`captureMedia` in `capture.go` (same coalesce/cap shape as `captureProfile`/
`captureEquipment`); called from `importOneCharacter` (character.go) and
`refreshOneCharacter` (refresh.go) right after `captureEquipment`, 404/403 → no-op, no
extra log (spec: `profile_unavailable` already covers it). Update existing
`import_test.go`/`refresh_test.go` fixtures that reach `captureEquipment` to also stub
`character-media` (three import tests, two refresh tests — see research notes). New test:
media URLs land in `avatar_url`/`render_url`/`bnet_media`.

### A4. `/v1/me` and public character endpoint gain the two fields
`auth/store.go`: `Character` struct gains `AvatarURL, RenderURL *string
json:"avatar_url,omitempty"/"render_url,omitempty"`; `characterColumns` adds
`c.avatar_url, c.render_url`; `scanCharacterRows` scans them. `rankings/query.go`:
`CharacterRef` gains the same two fields; the `select name, class from characters where
key = $1` query adds the two columns. Tests: `store_test.go` (or existing) asserts the
fields round-trip; `query_test.go` asserts the public endpoint carries them when present
and omits them when null.

### A5. Go verification
`export GOWORK=off && go vet ./... && go test ./...` from `api/`, scoped then full for the
touched packages (`bnetapi`, `bnetimport`, `auth`, `rankings`).

### B1. `web/src/components/ui/StatePanel.svelte` (new — the one `components/ui/` exception)
A Svelte twin of the design system's "State panel": label row with the glowing gold dot,
slot for key/value rows, optional `updated` stamp on the right. Generalizes
`AccountPanel.svelte`'s box treatment (`bg-raised border-line rounded-panel`) plus the
label-row markup `CharacterList.svelte`/`Account.svelte` currently hand-roll per section,
so Devices/You/Guilds-and-plan can share one component instead of three copies. SSR render
test: label, dot, rows, updated stamp present/absent.

### B2. `web/src/lib/account/layout.ts` — new two-column, panel layout constants
New skeleton heights for the 2-column layout (measured against the fixture at 360 and
1280). Keep the existing constants' names where the shape is unchanged; add new ones for
the rail panels and the hero band's reserved height.

### B3. Fixtures: local placeholder images + `me-bnet.ts`
Generate two tiny local PNGs/JPGs under `web/public/fixtures/` (36×36 avatar, 400×300
render) with a short Node/canvas-free script (raw byte-level BMP→PNG is overkill — use a
minimal valid PNG/JPEG built by a small script, or Node's built-in zlib to hand-roll a
1x1-scaled PNG, whichever is fastest and produces a valid, tiny file). Add
`avatar_url`/`render_url` (pointing at `/fixtures/...`) to Thoradin in `me-bnet.ts`.

### B4. `Account.svelte` — page header, two columns, rail panels
Title + identity line (BattleTag/email, "Signed in with Battle.net"/"Signed in by email",
Sign out) responsive per spec B1. `lg:grid lg:grid-cols-12 lg:gap-8`: main
(`lg:col-span-8`, Characters + Your reports), rail (`lg:col-span-4`, Devices / You /
Guilds-and-plan as three `StatePanel`s). Devices and Guilds panels show an `Updated` stamp
when a timestamp exists. Billing collapses to the one-row form the brief specifies when
`billing === null`.

### B5. Current-character hero band
Read `readCurrent()` in `Account.svelte`'s `$effect` (mirrors `CurrentCharacterBar`'s own
read); when `source === 'armory'` and `ref` matches a listed character with `render_url`,
render the hero band above Characters: render image (`object-contain max-h-[280px]
lg:max-h-[360px]`, `--bg` background, right-aligned) with name/descriptor/actions on the
left. No band when there is no match or no `render_url`.

### B6. `CharacterList.svelte` rows — avatar square, class-colour letter fallback
36px square: `avatar_url` image (`rounded-[3px] object-cover loading=lazy alt=""`) or a
class-coloured square with the class's first letter. Keep existing descriptor/guild-line/
actions; move the panel into the new `StatePanel` label-row treatment; footer copy line
stays once.

### B7. `Character.svelte` header render
`lg`: render on the right when `render_url` present; else 44px avatar when `avatar_url`
present; else nothing. Needs `CharacterPage.character` to carry `avatar_url`/`render_url`
in `web/src/lib/rankings/api.ts`'s `CharacterPage` type.

### B8. States, e2e, Lighthouse
Re-measure skeleton heights (Task B2) against `FOREVER_DATA=fixture` build at 360; update
`account-visual`/account e2e specs that assert on layout; run `npm run lhci` and quote the
CLS per URL touched.

### B9. Web verification
Scoped: `npx vitest run src/components/Account* src/components/account/** src/components/ui/StatePanel* src/components/Character*`,
`npx astro check`, `npm run lint`, `npx prettier --check <touched paths>`. Full suite once
at the end per house rules: `E2E_PORT=4451 npx playwright test`.

### C. Screenshots + report
`chromium.launch()` against `astro preview --port 4451` with `/v1/me` routed to the
fixture; phone (390×844) and desktop (1280×900) screenshots of the signed-in account page
saved to `.superpowers/shots/`. Final report per the brief's format C.
