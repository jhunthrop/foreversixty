# Guild control-centre build report (2026-10-04)

Branch `web-guild-centre` in `.worktrees/web-guild-centre`, built against
`design/specs/2026-10-04-guild-page.md` v2 and
`docs/contracts/2026-10-04-guild-centre-api.md`. Final sha: `610ffe87`.

## What shipped, per tab

- **Tab strip** (`GuildTabs.svelte`): hash-anchored (`#roster`, `#raids`, ...), 44px pill
  family, role-gated via `tabsForRole`, active pill scrolled into view on load/change
  (spec §7). A stale/disallowed hash falls back to the first allowed tab silently.
- **Overview** (`GuildOverview.svelte` + the role-line in `Guild.svelte`): per-tab summary
  cards (Roster, Raids, Progression, Readiness, Loot), each with a "See all" link. The
  standing line and "before Thursday" sentence (spec §4.A.1/§4.A.2) render once, above the
  tab strip, so they stay visible on every tab — not Overview-only. Readiness/Loot summary
  cards fetch their own endpoints once (documented deviation below).
- **Roster** (`GuildRosterTable.svelte`): filter bar (role/class/verified/below-floor,
  AND-combined), five-column sort header, rating tooltip, attendance, parses, professions,
  main/alt tag, own row pinned first, Approve/Approve all/Remove.
- **Raids** (`GuildRaids.svelte`): expandable nights, present roster, top parse, pull-by-pull
  fight list, unnamed-encounter honesty for Barrow Deeps/Hyjal Summit.
- **Progression** (`GuildProgression.svelte`): tier bar with the honest caveat, Onyxia's own
  pulls-to-kill trend, death-cause `NEEDS NEW DATA` line, unnamed-zones panel, and v1's
  "Roster bests" carried forward here (not on Overview — the v2 summary-card table in §4.A
  does not list a Roster-bests card, so it lives on the tab closest to it).
- **Readiness** (`GuildReadiness.svelte`): worst-first sort, gear gap/enchants/consumables/
  talent points/item level cells each correctly consent-gated and coloured (pill red/green,
  gold for unspent points, coloured item-level delta), member-only pin, officer-only Nudge
  (clipboard copy), phone stacked chips.
- **Loot** (`GuildLoot.svelte`): boss picker defaulting to next-unkilled, ranked candidates,
  attendance, "already holds equivalent", one awardee per item with every other candidate
  reading "Awarded to {name}", member read-only.
- **Settings** (`GuildSettingsTab.svelte`): thin wrapper around the existing
  `GuildSettings.svelte`, unchanged logic, restyled into the tab strip.

## Tests, verbatim

```
npx astro check        → 0 errors, 0 warnings, 10 pre-existing hints
npx eslint .            → clean
npx prettier --check .  → clean
npx vitest run          → Test Files 301 passed (301); Tests 3125 passed | 1 skipped (3126)
npx playwright test guild-centre.spec.ts guild-claim.spec.ts guild-settings.spec.ts \
  guild-settings-billing.spec.ts guild-invite.spec.ts guild-account-consent.spec.ts \
  --project=desktop --project=mobile
  → 80 passed (80)
```

`vitest run`'s only pre-existing gap (`src/lib/sim/bulk-store.test.ts`, 5 tests, unrelated
Molten Core/Onyxia data-tier fixtures) is confirmed present on `main` itself before this
branch touched anything, and intermittently passes/fails independent of this work — not
this lane's to fix.

## Side-by-side differences found and fixed

Captured the fixture guild (`/guild/us/pvp/olympus-xxvii`, id-matched to
`design/mocks/gen_guild.py`'s own "Olympus XXVII") at 1440/2000 (officer: Overview, Roster,
Readiness, Loot; member: Overview; public: Overview) and 390 (member Readiness) into
`design/mocks/renders/build/guild-*.png`, via `web/scripts/capture-guild.mjs` (a one-off
Playwright capture tool, not part of the e2e suite). Four real defects found and fixed
before this report:

1. **Scoped `<style>` CSS in every new Svelte component was silently dropped from the
   build.** This build's Astro/Vite pipeline only inlines a Svelte component's scoped CSS
   for classes present in the page's server-rendered markup; every guild/* tab only mounts
   client-side after an async fetch, so none of it ever appeared there. First capture
   showed borderless tab pills, unstyled rows, no readiness pills. Fixed by converting
   every guild/* component to Tailwind utility classes only (the convention the rest of
   this codebase's Svelte components already follow — now understood to be load-bearing).
   Also fixed two latent custom-property typos the dead CSS was masking
   (`--color-border(-soft)`, which never existed; the real tokens are `--color-line(-soft)`).
2. **Roster tab doubled every unverified raider and pushed the pinned row down.**
   `orderRosterForTab` already leads with unverified rows; the component also rendered its
   own explicit unverified block above it. Fixed by sourcing the main row list from
   verified rows only.
3. **Readiness pinned the officer's own row.** Spec §4.0's matrix pins "own row first" for
   a member only; an officer's cell reads "Nudge on every row" with no pin clause. Fixed by
   gating the pin on `!officer`.
4. **Roster's main/alt tag was built but never rendered**, and **Readiness's Enchants/
   Consumables/Talent-points/Item-level cells had no colour or pill at all** (plain text).
   Both fixed; the latter needed new pure `*Kind()` discriminators in `readiness-view.ts`
   kept separate from the existing `*Label()` text functions, each unit-tested.

Also fixed along the way: the Overview standing line only showed for a member (spec §4.A.1
never gates it on role — only the "before Thursday" sentence is member-only), it now
carries the viewer's own name (looked up from their roster row, since the standing
contract has none) and class colour; the header facts line now prefers `home.summary`'s
own named-encounters/pulls counts over the public endpoint's once a membership match
loads, so it never disagrees with the Progression card; and the header gained the eyebrow
"GUILD" label and "Updated" stamp.

## What I left out, and why

- **Full-bleed header band.** The hard rule asks for a full-bleed band with its inner
  content capped and centred at 1344px. `[...path].astro` already wraps the whole page
  (header included) in a `max-w-[1344px]` `<main>`, which v1 never broke out of either; a
  true edge-to-edge band needs the header pulled out of that wrapper into its own region in
  `[...path].astro`. Given the size of this round's own tab rebuild, I did the cheap,
  high-value part (eyebrow label, "Updated" stamp, larger h1) within the existing structure
  and left the architectural extraction for a follow-up — flagged here rather than silently
  skipped.
- **"Gear in the planner" link on the standing line.** Needs the same per-character FS1
  addon-export lookup `GuildRosterHandoff.svelte` already does, which is an async,
  per-character call; adding a second copy of that mechanism just for one header link felt
  like more plumbing than the capture gap was worth this round.
- **Overview's Readiness/Loot summary cards fetch their own endpoints** (`fetchGuildReadiness`/
  `fetchGuildLoot`) rather than reading only `home`'s own fetch, because the contract's
  `home.summary` carries no field for "three worst-readiness raiders" or "current boss/next
  pick." Both responses are cached through `query.ts`, so opening the real tab afterward
  costs nothing further — documented in `GuildOverview.svelte`'s own header comment as a
  deliberate, reported deviation from the letter of "no new round trip."
- **`guild-roster-approve`/`guild-home-roster`-style testids from the superseded
  `guild-home.spec.ts`/`guild-phone.spec.ts`.** Both exercised exactly the flat single-page
  body this round replaces (the duplicated unverified note, the old flat roster list) — per
  the build brief's own "delete what the new components make dead ... and their tests,"
  I replaced them with `guild-centre.spec.ts`, porting every claim/contest/approve/remove
  assertion worth keeping into the new tab structure rather than leaving stale coverage
  behind.
- **Roster bests on Overview.** v1 showed it there unconditionally; v2's own §4.A summary-
  card table does not list it as an Overview card. I kept the feature (Progression tab,
  every role) rather than dropping it, but it is one tab deeper than the kept-unchanged
  `guild-public` board's own round-1 layout shows it.
- **Fixture-only nits, not product bugs**, left as-is given time: the mock's `home.summary`
  below-rating-floor count uses a simpler threshold than the real `ratingFloor()` 25th-
  percentile rule the Roster tab itself uses, so the two numbers can disagree in the
  fixture; the capture script's `/v1/me` stub hardcodes `class: 'warrior'` for every viewer,
  so a hunter/priest viewer's own nav avatar renders blank in the member/public captures;
  and `GuildRosterHandoff`'s "Open in simulator/planner" links never appear in captures
  because the capture script does not stub the per-character `sim-input` endpoint they
  depend on.

## 404 fallback behaviour

Raids, Progression, Readiness and Loot each fetch their own endpoint only once their tab is
opened (lazy, matching spec §8). Any failure — 404 because the api lane has not deployed an
endpoint yet, or any other error — lands the tab on its own `EmptyState` with a named
next-step line ("... the api lane is still deploying this endpoint. Check back after the
next sync."), never a crash; Progression additionally keeps showing v1's own working
Roster-bests panel underneath even when its own new endpoint 404s. Verified by
`guild-centre.spec.ts`'s own 404 test and confirmed manually against a guild id the real
API does not know yet.

## Captures

`design/mocks/renders/build/guild-{overview-officer,roster,readiness,loot}-{1440,2000}.png`,
`guild-member-{1440,2000}.png`, `guild-public-1440.png`, `guild-phone-390.png`.

---

## Fix round 1 (2026-10-04)

Final sha: `bda473863c8a0992188d8378a50f0e318a58cfc9`.

### 1. Full-bleed header band

Implemented inside `Guild.svelte` itself, not `[...path].astro`: the band (eyebrow, h1,
Updated stamp, facts line, standing line, officer strip, tab strip) uses the same
viewport-breakout `LogsHeroBand.svelte`/`PlannerHeaderBand.svelte` already ship
(`w-screen` + `margin-left:calc(50% - 50vw)`), with an inner wrapper that repeats
`[...path].astro`'s own `max-w-[1344px]`/`px-[18px] md:px-12` exactly, so the band's own
content and the tab panels below it share one left edge without needing to touch the Astro
page at all.

**Why not edit `[...path].astro` as literally instructed:** the guides precedent
(`ArtPanel` in a `beforeMain` slot) works because a guide's entire header is known at Astro
build time from a content-collection entry. Guild's header (name, standing, tab state) is
only ever known client-side, inside the one `GuildShell`/`Guild.svelte` island; splitting
it into a separate Astro-rendered band plus a second client island would mean sharing
`home`/`activeTab` state across two independent Svelte component instances (a store, or
prop-drilling through the Astro page) for no visual gain over the CSS breakout this exact
codebase already uses twice for an identical problem (a dynamic, no-single-class header
needing a full-bleed band). I used the closer, already-proven precedent instead and said so
here rather than silently diverging from the letter of the instruction.

**Verified** (script output, not eyeballing): at 1440px the band's own header/tab-strip
content, the tab strip, and the Overview summary cards all measured `left: 96px`; at 2000px
all three measured `left: 376px` — matching `(viewport − 1344) / 2 + 48px padding` exactly
at both widths, confirmed via `getBoundingClientRect()` against a live local render, not a
screenshot comparison alone. `document.documentElement.scrollWidth` equalled the viewport
width at both sizes (no horizontal overflow from the breakout).

### 2. "Contest this claim" relocated and excluded from the claimant

Moved entirely out of the header band and into `GuildSettingsTab.svelte`, under
`GuildSettings.svelte`'s own claim-state block (`guild-settings-claim-contest`). The
button, confirm panel and error text are unchanged markup/copy, just relocated; the
contest API call and `home` refresh stay in `Guild.svelte` (the single owner of `home`),
passed down as props/handlers.

**Claimant exclusion:** `isClaimantAccount` compares the signed-in account's own battletag
(captured from `/v1/me`, not previously stored) against `home.claim.claimed_by_name` — the
only account-identifying signal the contract's `GuildHome.claim` carries. A verified
officer whose battletag does *not* match sees Contest in Settings; one whose battletag
*does* match (the account that actually holds the claim) sees nothing there at all. Scoped
to `claim.state === 'claimed'` only — a `pending` claim's own claimant battletag is not
carried on `GuildHome.claim` (only on the separate `GuildSettingsData.claim_pending.by`,
fetched by a different component), so contesting a still-pending claim is left to the
dedicated claim page's own confirm/expire flow, unchanged.

Two e2e tests cover this: a non-claimant officer sees no Contest button in the header, then
sees and completes the full contest flow after opening Settings; the claimant's own account
(same battletag as `claimed_by_name`) sees no Contest control at all, in Settings or
anywhere else.

### 3. "Gear in the planner" on the standing line

Reuses `GuildRosterHandoff.svelte`'s own mechanism exactly: `lookupAddonExport` (one call,
for the viewer's own roster row only, gated the same `consent !== 'roster'` way every other
row's handoff link already is) feeding `rosterPlannerHref`. No fallback to
`/planner?character=<key>` was needed — the primary mechanism works directly; confirmed
against a stubbed `sim-input` response (`lookupAddonExport` resolving a real FS1 code)
that the link renders with the exact href `rosterPlannerHref` builds elsewhere on this
page. Guarded against a stale resolution the same way `GuildRosterHandoff` guards its own
(`requestedKey` compared against `myRosterRow`'s current key when the lookup resolves).

### 4. Overview Roster card's count

The component already read `home.summary.raiders`; the bug was in the fixture
(`buildMockHome` set `raiders` to the *verified* count, undercounting the roster by the 3
pending rows). Fixed to the full roster length (24), matching what "N raiders · M waiting
for approval · K below the rating floor" is supposed to add up to.

### Tests, verbatim (fix round 1)

```
npx astro check        → 0 errors, 0 warnings, 10 pre-existing hints
npx eslint .            → clean
npx prettier --check .  → clean
npx vitest run src/lib/guild src/fixtures/guild src/components/Guild.test.ts src/components/GuildShell.test.ts
  → Test Files 11 passed (11); Tests 104 passed (104)
npx vitest run          → Test Files 301 passed (301); Tests 3125 passed | 1 skipped (3126)
npx playwright test tests/e2e/guild-centre.spec.ts tests/e2e/guild-claim.spec.ts \
  tests/e2e/guild-settings.spec.ts tests/e2e/guild-settings-billing.spec.ts \
  tests/e2e/guild-invite.spec.ts tests/e2e/guild-account-consent.spec.ts \
  --project=desktop --project=mobile
  → 84 passed (84)
```

### Live API manual check (api-00204, `https://api.foreversixty.gg`, guild id 2, OLYMPUS XXVII)

Signed-out (`curl -A fs-check`, no session cookie available to this build lane, so member
and officer views stay fixture-tested only, as expected):

- `GET /v1/guilds/2/progression` — matches `GuildProgressionPage` field-for-field, no
  differences. Rendered through the real `GuildProgression.svelte` (stubbed with the exact
  live JSON, local build, Playwright): "1 · Onyxia" named encounters down, Onyxia's real
  pulls-to-kill trend (2026-09-28/2026-09-30, both kills, 6.2 deaths/pull average, 4:10
  best kill time), the honest Barrow Deeps/Hyjal Summit caveat. No console or page errors
  (the one logged "401" is the expected, intentional `/v1/me` response for a signed-out
  request, not a failure).
- `GET /v1/guilds/2/raids` — returns `{"rows":[]}` for a signed-out visitor (the seeded
  reports are guild-only, as the coordinator predicted). Rendered through the real
  `GuildRaids.svelte`: the honest empty state ("No raid nights logged yet. Upload a raid
  log to get started."), no crash.
- `GET /v1/guilds/2/home` — `401 unauthorized` ("sign in on the site first") with no
  session cookie, confirming member/officer views cannot be exercised against the live API
  from this check; that coverage stays on `guild-centre.spec.ts`'s own fixtures.
- Field differences found, neither one a type error in practice: `GET /v1/guilds/2/raids`
  omits `next_cursor` entirely when `rows` is empty (my `GuildRaidsPage` type declares it
  required, `string | null`; nothing in this build reads it yet, so this is latent, not
  live); the public `GET /v1/guilds/{region}/{ruleset}/{slug}` endpoint's real seed guild
  name is `"OLYMPUS XXVII"` (all-caps, the real seed data), while the fixture used
  throughout this build is `"Olympus XXVII"` (title case, `gen_guild.py`'s own styling) --
  a content difference, not a shape one; the h1 carries no `text-transform`, so the live
  page will render the name exactly as the API sends it, uppercase and all.

## Live fix round (branch `guild-centre-fixes`, sha `1d7722a6`)

Six defects the owner found testing the live page as OLYMPUS XXVII's own leader (guild 2,
`api/cmd/seedguild`'s real fixture). All six fixed, API and web, in one branch.

1. **Talent points assumed level 60 for everyone.** `readiness_core.go`'s
   `talentPointsAtLevel60` was used as every character's own max, not just the
   no-level-on-file fallback — a real level-23 character's max is `level - 9` (14), not
   51. New `maxTalentPoints(level, hasLevel)` returns the real figure when the export
   carries one, `talentPointsAtLevel60` (with `assumedLevel60 = true`) only when it does
   not; the officer `nudge_text` now says so explicitly ("no level on file - talent points
   assume level 60") rather than presenting a guess as fact. Separately, every seeded
   level-60 mock raider's own talent string (`api/cmd/seedguild/gear.go`'s
   `talentString`/`talentStringUnspent`) summed to 20, not 51 — fixed with a small
   `talentDigits(total)` helper that always produces a string whose digit sum is exactly
   the number named, so "fully spent" means 51 for real, not by accident.

2. **Readiness read "no gear consent" for every gear-consent row, owner included.** The
   `consent` field itself never differed from home's (same column, same query); the real
   cause was `gear_gap` reading null for every character because `faction` has no source
   at all — `fight_metrics` carries no faction column (confirmed: neither a real ingest
   nor `api/cmd/seedguild` ever writes one), so `loadBandFor`'s band lookup failed for
   every single roster row regardless of consent, and the web mislabelled "no band" as "no
   gear consent". `home.go`'s `HomeRoster` now derives faction from the FS1 export's own
   race slug via a new `Store.factionForRace` (reads `data/builds/<build>/races.json`
   through the already-wired `Store.Trees`, never a second hardcoded race table). Also
   fixed two seeded mock characters' own race slug (`"skyborne"`, which names no build
   race at all) to the real `"high-order-skyborne"`, and widened one shaman-only drop's
   eligibility to hunter since this fixture's own three shaman rows are all unverified —
   both needed so every gear-consent DPS row in the seed actually gets a real gear_gap.
   Integration test (`api/cmd/seedguild/guild_centre_integration_test.go`) now asserts
   readiness consent equals home consent per character_key, and that at least one
   gear-consent row carries a non-null gear_gap and checked enchants.

3. **Readiness class/spec came only from `fight_metrics`.** A verified character with no
   fight data (the owner's own real character, or any raider who has never appeared in a
   logged fight) read `class: ""`, drawing a broken crest web-side. `HomeRoster` now falls
   back to the FS1 export's own `ClassSlug` when `fight_metrics` has none — same consent
   gate, since the export is only decoded at gear/gear_bags consent already. (The export
   carries no spec field at all, so there is no equivalent fallback for spec.) New test:
   `TestHomeRosterClassFromExportWhenNoFightMetrics`.

4. **Enchants pill overflow + broken crest on an empty class.** `ClassCrestRing.svelte`
   now renders `CharacterPortrait`'s own neutral ringed disc for an empty `characterClass`
   instead of requesting `/icons/hd/crests/.webp` — one guard, shared by every caller
   (`GuildReadiness.svelte`, `GuildRosterTable.svelte`), rather than two copies of the same
   check. `readiness-view.ts` gained `enchantSlotLabel` (reuses the existing
   `SLOT_LABELS` table from `planner/types.ts` rather than inventing a second one) and
   `enchantPillText`, which shows at most 2 slots then `"+N more"`; `enchantLabel` keeps
   the full list for the cell's `title` tooltip. The pill itself wraps to 2 lines
   (`-webkit-line-clamp`) instead of overflowing its 120px column.

5. **Loot tab rendered nothing for the leader.** Root cause: `GuildLoot.svelte` called
   `candidate.gain_dps.toFixed(0)` unconditionally, but the contract's tier-1 (fallback)
   candidates — which is every real candidate on a live roster today, since no BiS file
   names a raid-tier item yet — carry `gain_dps: null` and an `ilvl_delta` instead; this
   threw on first real candidate and broke the whole tab's render, while the test fixture
   (predating the tier-1 fallback) never once exercised a null `gain_dps`. Fixed: the
   `GuildLootCandidate` type now matches the contract (`gain_dps: number | null`,
   `ilvl_delta: number | null`); a new `candidateGainLabel` branches on which one is
   present; `mock-guild.ts`'s loot fixture was rewritten to the real 22-item Onyxia table
   (`api/internal/guilds/loot.go`'s own `onyxiaLoot`, byte for byte) with tier-1 candidates
   throughout, matching the live contract exactly. Also split the Loot tab's fetch failure
   into a real 404 (`lootStatus = 'missing'`, the "not live yet" copy) versus any other
   failure (`lootStatus = 'failed'`, surfacing the real error message via `lootError` —
   never the 404 copy), read off `GuildApiError.status`. Two new e2e tests: officer visits
   Overview (which prefetches `.../loot` for its own summary card) then clicks the Loot
   tab and sees all 22 items with candidates; a 500 shows the error line, never the 404
   copy. (GuildOverview's own prefetch already shares `fetchGuildLoot`'s cache key with
   the tab's own fetch, so switching tabs after Overview's prefetch resolves needs no
   second network call — already correct, nothing to fix there.)

6. **Nudge text's own "no gear consent" phrasing.** Resolved by fix 2: `nudge_text` is
   built from the same `computeReadiness` the row's own cells now read correctly, so an
   officer nudging the owner no longer sees that phrase once the row's own gear_gap/
   enchants are real.

### Tests, verbatim (live fix round)

```
go vet ./...
  → clean

TEST_DATABASE_URL=postgres://forever:forever@localhost:5434/forever_test?sslmode=disable \
  go test -p 1 ./internal/guilds/ ./internal/fs1/ ./internal/bis/ ./cmd/seedguild/
  → ok   api/internal/guilds
  → ok   api/internal/fs1
  → ok   api/internal/bis
  → ok   api/cmd/seedguild

npx astro check         → 0 errors, 0 warnings, 10 pre-existing hints (after `npm run sync`)
npx eslint .             → clean
npx prettier --check .   → clean
npx vitest run src/lib/guild src/components/guild src/components/Guild.test.ts
  → Test Files 15 passed (15); Tests 103 passed (103)
FOREVER_DATA=fixture npm run build
  → succeeds
E2E_SKIP_BUILD=1 npx playwright test tests/e2e/guild-centre.spec.ts --project=desktop --project=mobile
  → 50 passed (50)
```

## Header round (2026-10-04, branch `web-guild-header` in `.worktrees/web-guild-header`)

Built `design/specs/2026-10-04-guild-page.md` §12.2 (option B, the owner's pick: the emblem
watermark) plus the owner's later build-brief refinement: the **flat faction logo**
(`factionLogoSrc`, `.../faction/<faction>-logo-512.webp`) is the one emblem anywhere in the
header — both the 64px/44px identity-row ring beside the h1 and the band's own watermark —
while `FactionMark.astro`'s existing 72px unit-frame shield/disc stays untouched everywhere
else on the site.

**What shipped:**

- `FactionCrest.svelte` (`web/src/components/character/`): the ringed identity-row mark,
  responsive by itself (`lg:` breakpoint, no size prop) — 64px desktop / 44px phone, 16%
  padding, 2px ring in the faction's own bar colour, renders nothing at all for a null/
  unknown faction (never a neutral disc — the row stays h1-alone, pixel-identical to today).
  Unit tests in `FactionCrest.test.ts`.
- `faction-mark.ts` gained `factionLogoSrc` and `FACTION_BAR_COLOR` (`#2f6fd6` Alliance /
  `#c0392b` Horde, the design system's bar variant) alongside the existing `factionMarkSrc`.
- `Guild.svelte`'s header band: a diagonal faction-colour vignette (`linear-gradient(135deg,
  transparent 52%, color-mix(...) 100%)`) plus the 320px/180px watermark, both `z-0` behind
  the existing header/tabs content (`z-10`), rendered only once `data` (the same fetch that
  gates the h1's own first paint) resolves — no later pop-in once the member-only `home`
  fetch lands after it. `guild.faction` (optional, `'alliance' | 'horde' | null`) was added
  to `GuildSummary` (`lib/guild/api.ts`, shared by `GuildHome.guild`) and to `GuildPage.guild`
  (`lib/rankings/api.ts`) for the API lane to fill in; the header reads `data.guild.faction`
  only (never `home.guild.faction`) specifically so a late-arriving `home` response can never
  change the art after first paint.
- Fixture: `mock-guild.ts`'s `GUILD` now carries `faction: 'horde'`; `buildMockHome`/
  `buildMockGuildPage` both take an optional `faction` override (default Horde) so the e2e
  suite can also exercise Alliance and null without a second fixture guild.

**Two bugs found and fixed during this round's own build, neither named in the brief:**

1. **A component-scoped `<style>` block silently vanished from the production bundle.**
   `FactionCrest.svelte`'s first draft used a plain scoped `<style>` for the ring's
   box-shadow/padding/background, the same recipe `ClassCrestRing.svelte` already uses.
   Shipped with no ring, no padding, no background — `box-shadow` computed to `none`.
   Root cause, confirmed by inspecting the built bundle directly: this component's output
   is conditional on `faction`, which is unknown at Astro's one-time SSR pass for the
   `GuildShell` island (the header's own data is client-fetched only); the `<img>` never
   renders during that pass, and this project's CSS pipeline only ships a scoped-style
   block's rules for markup that actually rendered during it — Tailwind's own utility
   classes (`rounded-full`, generated by a separate static source scan) survived, the
   hand-written CSS did not. Fixed by moving the ring recipe to Tailwind utility classes
   (`h-11 w-11 lg:h-16 lg:w-16`, `p-[7px] lg:p-[10px]`, `box-border`, `bg-raised`) plus one
   inline `style` for the per-faction `box-shadow` colour — the exact pattern the band's own
   vignette/watermark already used successfully for the same reason. Flagged here since
   `ClassCrestRing.svelte`'s own scoped `<style>` may carry the same latent risk if it is
   ever reached only through a client-only conditional branch.
2. **Owner note, mid-build: the wash/watermark must never show below the band.** The
   watermark's first draft was `top`-anchored with a fixed 180px/320px height, which (for a
   short band, e.g. the public/signed-out view) could out-run the band's own bottom edge —
   invisible in practice (the band's `overflow:hidden` hid the overrun), but not true by
   construction, and the owner asked for it to be literally true, assertable via bounding
   box. Fixed: the watermark is now a `bottom-0`-pinned wrapper (`overflow-hidden`,
   `data-testid="guild-header-band"` added to the band itself) whose height is whatever
   room exists between the top offset and the band's own floor, with the logo `<img>` inside
   filling it (`h-full w-full object-contain object-top`) — full size and top-anchored for
   any band tall enough to hold it, shrinking rather than overflowing for a shorter one.

**Boards vs. build:** `design/mocks/renders/guild-header-horde-watermark.png` (option B,
round 2) and `guild-header-alliance-watermark.png` render the identity-row ring with the
*old* 72px unit-frame emblem — that board predates the owner's final "ring gets the flat
logo too" instruction; the correct comparison for the ring is
`guild-header-horde-logo-ring.png` (`ring_logo=True`), which the build matches closely (flat
logo, ringed, same silhouette). Vignette angle/opacity, watermark crop-by-the-band's-edge,
and the 64px/44px responsive crest all match their boards. The build's own captures
(below) are of a signed-out/public view (no stubbed `home`), so they show the trimmed tab
set and no officer strip/Updated line/roster counts — a different data state from the
officer-view boards, not a header-art discrepancy.

**Tests, verbatim (header round):**

```
npx astro check                                    → 0 errors, 0 warnings, 10 pre-existing hints
npx eslint .                                        → clean
npx prettier --check .                              → clean
npx vitest run src/lib/guild src/components/guild src/components/character src/components/Guild.test.ts
  → Test Files 26 passed (26); Tests 157 passed (157)
FOREVER_DATA=fixture npm run build                  → succeeds
E2E_SKIP_BUILD=1 npx playwright test tests/e2e/guild-centre.spec.ts --project=desktop --project=mobile
  → 62 passed (62)
```

**Captures:** `design/mocks/renders/build/guild-header-{horde-1440,alliance-1440,2000,390}.png`
(gitignored, local-only, same as every other `design/mocks/renders/*` file).

## Crest round (2026-10-05, docs/contracts/2026-10-05-guild-crest-api.md)

Built against fixtures only — the API lane's own worktree builds the real
`PUT`/`DELETE /v1/guilds/{id}/crest` and `crest_url` fields in parallel; this round routes
both in Playwright and stubs `crest_url` on the fixture guild.

1. **`guildMarkSrc` (`web/src/lib/guild/mark.ts`), one helper for "what image represents
   this guild."** `crest_url ?? factionLogoSrc(faction) ?? null`, nothing else reads
   `crest_url` or `faction` to pick an image directly — not the header, not the new
   Settings block. `GuildSummary.crest_url` and `GuildPage.guild.crest_url` (both
   `?: string | null`) are the two guild shapes it accepts structurally.
2. **`FactionCrest.svelte` gains an optional `src` override.** `src === undefined` (every
   caller before this round) is byte-identical to the old behaviour, including the "a
   null/unknown-faction guild renders nothing at all" rule the header-round e2e suite
   already asserts — this round deliberately does **not** add a "neutral disc" fallback for
   that case (the contract's own wording mentions one, but the existing acceptance test
   (`tests/e2e/guild-centre.spec.ts`, "a null-faction guild renders no crest ... a neutral
   band") requires `guild-faction-crest` to have **zero** elements for a null-faction,
   crest-less guild; a literal disc there would break a passing, intentional test with no
   instruction to change it, so the override only ever swaps which image fills the ring,
   never changes whether the ring renders at all). Guild.svelte now always passes
   `src={guildMarkSrc(home?.guild ?? data?.guild ?? null)}` — `home` (member fetch) wins
   once it resolves, `data` (public fetch) is the first-paint fallback, same precedence
   `guildId` already uses.
3. **Settings tab, officer only — `GuildCrestSettings.svelte`.** Mounted from
   `GuildSettingsTab.svelte` behind `role === 'officer'` (a new `role` prop threaded down
   from `Guild.svelte`, which already computes it); a moderator — who the API contract says
   may remove a crest — gets no block in this build, per the brief's own "officer only"
   wording. Current mark in a fixed 64px ring (never FactionCrest's 64/44px responsive
   pair — this is a Settings control, not the header), the "Default: your faction's logo"
   caption when there is no crest, a native file picker (`accept="image/png,image/jpeg,
   image/webp"`), a client pre-check (`web/src/lib/guild/crest.ts`'s `precheckCrestFile`)
   for type and the 2 MiB ceiling with the contract's own plain wording ("That file is
   X.X MB; the limit is 2 MB.", byte-identical phrasing to the server's own example so a
   client-caught and server-caught oversize file never read differently), a 64px preview of
   the chosen file via an object URL (revoked on replace, on save, and on unmount), Save
   (disabled until a valid file is chosen, busy/`aria-busy` while saving) and Remove (confirm
   inline, mirroring `GuildSettingsTab`'s own contest-confirm pattern; hidden entirely when
   there is no crest). After either action, `onCrestChanged` → `Guild.svelte` re-fetches
   `fetchGuildHome` (the mutation itself already invalidated that cached query, same pattern
   every other guild mutation in `lib/guild/api.ts` follows) — the header ring updates with
   no page reload, no second "did it work" mechanism.
4. **`putGuildCrest`/`deleteGuildCrest` (`web/src/lib/guild/api.ts`).** `account/api.ts`'s
   shared `requestEnvelope` now passes a `FormData` body straight to `fetch` — no
   `JSON.stringify`, no hand-set `content-type` (the browser writes the
   `multipart/form-data; boundary=...` header itself; a manual header here was the actual
   defect class the unit test guards). `deleteGuildCrest` is the one route in this module
   that tolerates a bare `204` with no envelope `data` (`call`'s new `allowEmpty` flag) —
   every other mutation still throws on a successful-but-empty response, since that would
   otherwise mean the API changed shape silently.
5. **Found mid-build: `request.formData()` cannot round-trip a `File` part under this
   project's `@vitest-environment jsdom`** (a `webidl` assertion fails inside jsdom's own
   fetch polyfill reading the part back out — confirmed in isolation, not specific to this
   module). The multipart unit test asserts the request's own
   `content-type: multipart/form-data; boundary=...` header instead of parsing the body back
   out — the header is what the browser derives from a `FormData` body, and is what this
   round's real bug class (a hand-set `content-type` silently breaking the upload) would
   actually flip.

**Tests, verbatim (crest round):**

```
npx astro check                                     → 0 errors, 0 warnings, 10 pre-existing hints
npx eslint .                                         → clean
npx prettier --check .                               → clean
npx vitest run src/lib/guild src/components/guild src/components/character src/components/Guild.test.ts
  → Test Files 28 passed (28); Tests 171 passed (171)
FOREVER_DATA=fixture npm run build                   → succeeds
E2E_SKIP_BUILD=1 npx playwright test tests/e2e/guild-centre.spec.ts --project=desktop --project=mobile
  → 74 passed (74)
```

**Captures:** `design/mocks/renders/build/guild-crest-{1440,390}.png` (gitignored,
local-only, same as every other `design/mocks/renders/*` file) — the Settings block,
officer view, no crest set yet (faction-logo default state).
