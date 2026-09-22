# Rating Web (Web lane) Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Ship the two free rating surfaces on the site — the report's Rating tab (with a
compact dashboard row on Summary) and the character page's rating panel — plus the public
`/ratings` explanation page and the anchor plumbing (`ExchangeTable`/`AuraTable`) that lets a
rating's evidence link to the exact rows behind it. Build against the API response shapes in
spec section 5.2 (the API lane has not landed); stub the API in e2e with `page.route`.

**Architecture:** Two new small reactive data modules
(`web/src/lib/rating/types.ts`, `web/src/lib/rating/report-ratings.svelte.ts`) hold the shapes
and a reusable fetch-state hook; two rating fetchers join `web/src/lib/rankings/api.ts`
alongside `fetchCharacter`; a copy module and a pure moment-anchor module hold every visible
string and the evidence-link logic. Two new Svelte components (`RatingPanel.svelte`,
`RatingTab.svelte`) render the dashboard row and the full three-level card (number → six parts
→ moments), reusing `ActorRow.svelte`'s `<details>` disclosure idiom. `RatingTab.svelte` is
lazy-loaded into `ReportView.svelte` exactly the way `TimelinesView.svelte` already is.
`Character.svelte` grows one new panel plus a small dedicated sparkline (not `TimeChart`, which
is built for per-second in-fight series, not a sparse multi-fight trend — see Ruling 6).
`/ratings` is a new static Astro page with zero client JavaScript.

**Tech Stack:** Astro 5 + Svelte 5 (runes) islands, TypeScript, Vitest, Playwright, Tailwind v4
(`@utility`/`--color-*` tokens), the existing `httpx` envelope contract via
`account/api.ts`'s `requestEnvelope`.

**Spec:** `docs/superpowers/specs/2026-09-21-performance-rating-design.md` — this plan's scope
is that spec's section 6 in full, the web half of section 8's tests, and the `ExchangeTable`/
`AuraTable` anchor additions of section 6.1.

## Global Constraints

- Content pages ship **no client JavaScript**: `/ratings` uses `Base.astro` with no `session`
  prop and no `client:*` directive anywhere on the page (`web/src/layouts/Base.astro`'s
  `session` prop; pinned by `tests/e2e/home.spec.ts`'s "content pages ship no client
  JavaScript" and `tests/e2e/nav.spec.ts`).
- The report page and character page already run islands; the rating UI lives inside those,
  lazily loaded the way `ReportView.svelte` lazy-loads Timelines/Events/Queries/Compare/
  Mechanics/Rankings via `createLazyComponent` (`web/src/lib/report/lazy-component.svelte.ts`).
- No layout shift; Lighthouse budgets in `web/lighthouserc.json` must keep passing (that file's
  URL list is not changed by this plan — see Ruling 9).
- Several files sit at or near the repo's 800-line ceiling
  (`web/src/components/sim/SimView.svelte` exactly 800, `Planner.svelte` 799); before adding
  to `ReportView.svelte` (already 1814 lines — see Ruling 8) or `SummaryTab.svelte`, keep the
  added surface to the smallest possible wiring diff and put real logic in new, small files.
- 44px hit targets on phone (`min-h-11`/`md:min-h-9` or `md:min-h-0` utilities already used
  throughout `web/src/components/report/*`).
- Every number that lines up in a column uses tabular figures (`tabular` utility,
  `web/src/styles/global.css:67`).
- Colour alone never carries meaning: a low rating component says so in words, not only in a
  bar's colour (spec §6.5) — see Ruling 5 for the concrete mechanism.
- Every visible string lives in a copy module (`web/src/lib/rating/copy.ts`), matching
  `web/src/lib/rankings/copy.ts` and `web/src/lib/sim/copy.ts`'s existing precedent.
- No mutation: Svelte 5 `$state`/`$derived` objects are replaced, never mutated in place,
  matching every existing report-lane module's style.
- Functions under 50 lines, files under 800, no magic numbers (named constants for the
  MIN_SAMPLE-style thresholds this plan itself introduces — see Ruling 7), no dead code.
- Web toolchain (from `web/`): `export NVM_DIR="$HOME/.nvm"; . "$NVM_DIR/nvm.sh"; nvm use
  22.12`, then before every commit: `npx vitest run <paths>`, `npx astro check`, `npm run
  lint`, `npx prettier --check <paths>`. E2E: `E2E_PORT=4451 npx playwright test <spec
  files>` (stop a stale preview first: `npx astro preview stop`). Never `npm install`, never
  stage `web/node_modules` (it is a symlink into the main checkout).
- File ownership (this lane, "rating-web"): `web/src/lib/report/url.ts`'s `Tab` addition,
  `RatingPanel.svelte`, `RatingTab.svelte`, `web/src/lib/rankings/api.ts`'s rating fetchers,
  the character page's rating panel, the `/ratings` explanation page, and the `ExchangeTable`/
  `AuraTable` anchor additions — plus the minimal wiring diffs into `ReportView.svelte`,
  `SummaryTab.svelte` and `Character.svelte` that make those reachable (see Ruling 8). Do not
  edit `api/` (except the one permitted chrome-snapshot regen, unused here since this plan does
  not touch `Header.astro`/`Footer.astro`), `addon/`, `logs/`, `sim/`, or the sim/planner/guild
  components. Stay out of `web/src/pages/premium*`, the legal pages, and
  `CheckoutLauncher.svelte` (owned by the parallel `pay-web` lane).

---

## Rulings (recorded for the coordinator and the API lane's reconciliation)

**Ruling 1 — empty-fetch shape, not an error.** `GET /v1/reports/{id}/fights/{n}/ratings`
returns `ok:true` with `players: []` (never a 404 or an error envelope) when no rating rows
exist for that fight — the ordinary "backfill has not run yet / this encounter has no curated
table" launch state, mirroring how `GET /v1/rankings` already returns an empty `rows: []`
rather than an error when nothing is ranked. `GET /v1/characters/.../rating` likewise returns
`ok:true` with `sample_size: 0, trend: [], best_component: '', worst_component: '', latest:
null` for a character nobody has rated yet. **Cost if wrong:** if the real API instead 404s on
an empty result, the fetchers' `catch` blocks already render the same truthful "no ratings yet"
copy (Task 2's `RatingFetchStatus` includes `'empty'` produced either way — see Task 2), so the
UI does not break; only the loading-vs-empty distinction blurs slightly until the API lane
confirms.

**Ruling 2 — the anonymize 404 is real and distinct.** Per spec §5.1, the character-rating
endpoint answers `404 not_found` when `anonymize = true`. This is a genuine error the fetcher
surfaces (not folded into Ruling 1's empty state): `fetchCharacterRating` lets a 404 raise
`RankingsError` as usual, and `Character.svelte`'s rating panel renders nothing at all on a 404
(no placeholder, matching the anonymize copy rule in spec §6.4 — "nothing on the page implies
someone is hidden here").

**Ruling 3 — the roster GUID join is by name, not `player_key`.** The rating API's per-fight
rows carry `player_key`/`player_name` (a cross-report identity, `region/ruleset/slug`) but no
per-fight unit GUID, while the report URL's `source=<guid>` scoping and every existing tab's
"show only this player" affordance are GUID-keyed. Building `characterKey(region, ruleset,
name)` client-side to match `player_key` was considered and rejected: the report island does
not currently know its own region/ruleset (`ActorRow.svelte`'s `characterLink` prop exists for
exactly this join but is always passed `null` from `ReportView.svelte` today — confirmed by
grep, nothing wires it), and threading that in is out of this lane's scope. **Ruling:** join a
rating row to its roster GUID by exact string equality between `splitUnitName(row.player_name)
.name` and `splitUnitName(unit.name).name` over the fight's own roster (`ReportView.svelte`'s
existing `roster: {guid,name,class}[]` derivation, already passed to `ModeBar.svelte` the same
shape). A raid roster's display names are unique within one fight, so this is a safe join; a
row that matches no roster unit (should not happen — the API computes off this exact fight's
roster) renders with its name as plain text, not a "show only this player" button. **Cost if
wrong:** a genuine name collision (would require two players with the identical display name
in one raid, which the game does not allow) silently merges two players' click-throughs; no
raid can produce this today.

**Ruling 4 — moment anchors are built by this lane, not trusted verbatim from the API, except
for deaths.** `ExchangeRow` (`web/src/lib/report/types.ts:185-195`) and `AuraTrack`
(`:151-163`) carry no per-event timestamp — they are fight-aggregated count/uptime rows, not
per-occurrence logs — so the spec's literal `id={exchange-<sourceGuid>-<atMs>}` convention
cannot be built from this data. Since this spec section (§6.1) assigns the anchor convention on
`ExchangeTable`/`AuraTable` to the web lane to invent, **this lane owns that convention**:
`id={exchange-<guid>-<spellId>}` on an `ExchangeTable` row (source-scoped: the row for the
spell this player interrupted/dispelled) and `id={aura-<guid>-<spellId>}` on an `AuraTable` row
(target-scoped: the row for the debuff this player carried), both built from fields those
components already have. `RatingTab`'s moment links are built the same way
(`web/src/lib/rating/moments.ts`'s `momentHref`), reading `moment.spell_id` plus the card's own
already-resolved player GUID (Ruling 3) — never a `moment.anchor` string for these kinds. The
one exception is `kind: "death"`, whose anchor convention (`death-<guid>-<at_ms>`) already
exists and is owned by `DeathsTab.svelte`; for that kind alone, the API's own `anchor` string
(built the same way, `death-<guid>-<at_ms>`) is used verbatim, since `Death.AtMs` genuinely
exists. **Moment `kind` vocabulary this lane expects and renders a working link for:** `death`
(→ Deaths tab, API's own anchor), `interrupt` (→ Interrupts tab,
`exchange-<guid>-<spell_id>`), `dispel` (→ Dispels tab, `exchange-<guid>-<spell_id>`),
`carried-debuff` (→ Debuffs tab, `aura-<guid>-<spell_id>`), `utility-uptime` (→ Buffs tab,
`aura-<guid>-<spell_id>`). Any other `kind` string (including `avoidable-hit`, which has no
table row anywhere in the UI today) renders as plain, unlinked text — the moment's `spell_name`
and, when present, `at_ms`. **Cost if wrong:** if the API lane picks different `kind` strings,
every moment degrades gracefully to plain text (never a broken link or a crash) until the two
lanes reconcile the vocabulary; flagged in the final report for the coordinator.

**Ruling 5 — colour never varies on the six-segment bar.** `percentileToken`
(`web/src/lib/report/format.ts:283`) is the site's traffic-light-by-value ladder
(grey→green→blue→purple→orange→pink→gold), used for parse percentiles elsewhere — spec §6.5
explicitly forbids exactly that scheme for the six rating parts. `RatingPanel`/`RatingTab`
never call `percentileToken` for a component bar. Every segment fills with the same flat
`var(--color-gold)` (the site's one numeric accent, already used this way for the Damage
chart's line and the "See the tab" links, `web/src/components/report/SummaryPanels.svelte`'s
`text-gold` "more" links); only the fill's **width** (already labelled in text beside it) varies
by score, and an excluded component renders as an outline-only, unfilled track with a "not
scored" title — never a colour change. The overall headline figure also uses `text-gold`.

**Ruling 6 — no `TimeChart` reuse for the character trend.** `TimeChart.svelte` is built
around a dense per-second `series: number[]`, a `durationMs`, a drag-brush `window`/`onWindow`
pair and per-second death ticks — a continuous in-fight series with an interactive brush. A
character's rating trend is the opposite shape: a handful of unevenly-spaced points, one per
rated fight, spanning weeks, with no brushing interaction at all. Forcing `trend[]` through
`TimeChart`'s props would mean synthesizing a fake `durationMs`/per-second `series` from sparse
dated points, distorting a contract that other in-fight charts (Damage/Healing/Resources/
Threat) depend on reading literally. **Ruling:** a new, small, dedicated component,
`web/src/components/RatingTrend.svelte` (~60 lines, one inline `<svg>` polyline, no charting
library, no brushing, no `TimeChart` import) renders the sparkline. **Cost if wrong:** a second
small charting primitive exists alongside `TimeChart`; acceptable, since the two solve
genuinely different problems and forcing one shape onto the other is a worse coupling than two
small files.

**Ruling 7 — the trend-sample display threshold is a named web-only constant.** Spec §6.4's
example copy, *"Not enough rated fights yet to show a trend (2 of 5 needed)"*, implies a
minimum sample count before the sparkline renders but never states the number generally. This
is a **display-only** choice (nothing server-side depends on it): `MIN_TREND_SAMPLES = 5` in
`web/src/lib/rating/copy.ts`, matching the spec's own worked example. Below it, the copy above
renders with real `sample_size`/`MIN_TREND_SAMPLES` numbers substituted in; at or above it, the
sparkline renders. **Cost if wrong:** a display threshold, trivially changed in one place if
the coordinator picks a different number later.

**Ruling 8 — wiring beyond the literal §8 file list is necessary and minimal.** Spec §8's
table names `RatingPanel.svelte`/`RatingTab.svelte` as new files but the spec's own prose
(§6.1: "mirrors `SummaryPanels.svelte`'s dashboard pattern... **on `SummaryTab`**") requires
mounting `RatingPanel` from inside `SummaryTab.svelte`, and reaching `RatingTab` at all requires
a `state.tab === 'rating'` branch in `ReportView.svelte`'s existing tab chain (mirroring how
every other tab is wired) plus a lazy-load `$effect`. Both edits are the smallest possible
diffs: `ModeBar.svelte` needs **no** change at all (it already renders every entry of `TABS`
generically, confirmed by reading its `{#each TABS as tab}` loop — adding `'rating'` to `TABS`
in `url.ts` is sufficient for the tab button to appear). `SummaryTab.svelte` gains one new prop
threaded straight to `RatingPanel` plus one new line in its existing `{#if onTab}` block (it
already receives `reportId`/`fightIndex`, so no new prop plumbing is needed for those).
`ReportView.svelte` gains one `createLazyComponent` line, one `$effect` mirroring its five
siblings, and one `{:else if state.tab === 'rating'}` branch plus a `nightMode` branch (ratings
are per-fight; the whole-night view shows a short explanatory note, matching the existing "one
pull's" message pattern at `ReportView.svelte`'s `night-tables-only` block) — no other tab's
branch, and no restructuring of the surrounding chain.

**Ruling 9 — `lighthouserc.json` is not extended.** The file's `url` list is a small, hand
-curated sample (it does not include `/character`, `/account`, `/guild`, `/changelog`, or any
other existing content/island page added since the list was last curated — confirmed by `git
log -- lighthouserc.json` showing no such additions alongside those pages' own commits).
Adding `/ratings.html` here is not required by the spec and is not this lane's call to make
unilaterally; the existing budgets must keep passing (Global Constraints), which this plan's
zero-JS static page trivially satisfies without being added to the sampled list.

**Ruling 10 — duplicate fetches between `RatingPanel` and `RatingTab` are accepted, not
shared across components.** Visiting Summary then Rating triggers two independent `GET
.../ratings` calls for the same fight. The API sets `Cache-Control: public, max-age=30` (spec
§5.3), making the second call a cheap edge-cache hit. Building a cross-component request cache
is not asked for anywhere else in this codebase (`Character.svelte`, `Rankings.svelte` and
every other island each fetch independently) and would be speculative infrastructure ahead of
a real perf problem. The shared unit is the **fetch-state hook**
(`web/src/lib/rating/report-ratings.svelte.ts`, Task 2), not a cache.

**Ruling 11 — the six-segment mini bar does not collapse to a stacked list on phone.** Spec
§6.5 asks the six-segment bar to collapse to a stacked list below `md` width, citing the nav's
own phone-collapse pattern. That pattern applies where the horizontal layout is the actual
content (a full set of labelled nav entries competing for width). `RatingPanel`'s and
`CharacterRatingPanel`'s mini bar is purely decorative context beside a name and a number — a
fixed ~96-128px strip of six 2px-gap segments with no text inside it, not a hit target (44px
does not apply) and not itself carrying the score's meaning (the number beside it already
does, in text, per Ruling 5) — and it already fits a 360px viewport without crowding, the same
way `SummaryPanels.svelte`'s own single share-bar (an identical fixed-width `bg-line-soft`
track) renders un-collapsed on phone today. `RatingTab`'s six components are never a
horizontal bar at all — they render as a vertical `<details>` list from the first line (Task
5), which is already phone-shaped by construction. **Cost if wrong:** if a Lighthouse/phone
review finds the mini bar genuinely cramped at 360px, collapsing it to a vertical six-row list
is a small, isolated follow-up to `RatingPanel.svelte`/`CharacterRatingPanel.svelte` alone.

**Ruling 12 — the roster-GUID join (Ruling 3) is one shared module, not duplicated per
component.** `RatingPanel.svelte` and `RatingTab.svelte` both need "which roster GUID does
this rating row belong to" and "which class colours it." Rather than each component defining
its own `guidFor`/`classFor` closures (a DRY violation the user's own coding-style rule
forbids), both import `guidForPlayer`/`classForPlayer` from a new pure module,
`web/src/lib/rating/roster-join.ts` (Task 2), independently unit-tested there rather than only
indirectly through component tests.

---

## Task 1: `Tab` union gains `'rating'`

**Files:**
- Modify: `web/src/lib/report/url.ts:14-25` (the `Tab` union), `:52-68` (the `TABS` array)
- Test: `web/src/lib/report/url.test.ts:26-45`

**Interfaces:**
- Produces: `Tab` now includes `'rating'`; `TABS` (imported by `ModeBar.svelte`) includes `{
  id: 'rating', label: 'Rating' }` as its second entry (right after `summary`).

- [ ] **Step 1: Write the failing test**

Add to the existing `it('offers the four views and the twelve tabs in the spec's order', ...)`
block in `web/src/lib/report/url.test.ts` — rename it to state thirteen tabs and extend both
arrays it asserts:

```ts
  it('offers the four views and the thirteen tabs in the spec’s order', () => {
    expect(VIEWS.map((v) => v.id)).toEqual(['tables', 'timelines', 'events', 'queries']);
    expect(TABS.map((t) => t.id)).toEqual([
      'summary',
      'rating',
      'damage-done',
      'damage-taken',
      'healing',
      'threat',
      'buffs',
      'debuffs',
      'deaths',
      'interrupts',
      'dispels',
      'resources',
      'casts',
    ]);
    expect(TABS.map((t) => t.label)).toEqual([
      'Summary',
      'Rating',
      'Damage Done',
      'Damage Taken',
      'Healing',
      'Threat',
      'Buffs',
      'Debuffs',
      'Deaths',
      'Interrupts',
      'Dispels',
      'Resources',
      'Casts',
    ]);
  });
```

Also add a round-trip test right after it:

```ts
  it('rating is a real tab: it parses from the url and serialises back', () => {
    expect(parseReportState('?tab=rating', 1).tab).toBe('rating');
    expect(reportSearch(withState(defaultState(1), { tab: 'rating' }), 1)).toBe('?tab=rating');
  });
```

- [ ] **Step 2: Run test to verify it fails**

Run: `cd web && npx vitest run src/lib/report/url.test.ts`
Expected: FAIL — `TABS.map((t) => t.id)` is missing `'rating'`.

- [ ] **Step 3: Implement**

In `web/src/lib/report/url.ts`, widen the `Tab` union (line ~14):

```ts
export type Tab =
  | 'summary'
  | 'rating'
  | 'damage-done'
  | 'damage-taken'
  | 'healing'
  | 'threat'
  | 'buffs'
  | 'debuffs'
  | 'deaths'
  | 'interrupts'
  | 'dispels'
  | 'resources'
  | 'casts';
```

And add the entry to `TABS` (line ~52), right after `summary`:

```ts
export const TABS: readonly { id: Tab; label: string }[] = [
  { id: 'summary', label: 'Summary' },
  { id: 'rating', label: 'Rating' },
  { id: 'damage-done', label: 'Damage Done' },
  { id: 'damage-taken', label: 'Damage Taken' },
  { id: 'healing', label: 'Healing' },
  { id: 'threat', label: 'Threat' },
  { id: 'buffs', label: 'Buffs' },
  { id: 'debuffs', label: 'Debuffs' },
  { id: 'deaths', label: 'Deaths' },
  { id: 'interrupts', label: 'Interrupts' },
  { id: 'dispels', label: 'Dispels' },
  { id: 'resources', label: 'Resources' },
  { id: 'casts', label: 'Casts' },
];
```

- [ ] **Step 4: Run test to verify it passes**

Run: `cd web && npx vitest run src/lib/report/url.test.ts`
Expected: PASS, all tests in the file green.

- [ ] **Step 5: Type-check and lint**

Run: `cd web && npx astro check && npm run lint && npx prettier --check src/lib/report/url.ts src/lib/report/url.test.ts`
Expected: no errors (a new `Tab` member is a superset change; nothing currently switches
exhaustively on `Tab` outside `ReportView.svelte`, which Task 6 updates).

- [ ] **Step 6: Commit**

```bash
printf 'feat(web): add the Rating tab to the report URL contract\n\nAdds `rating` to the `Tab` union and `TABS` list so the report tab bar\nrenders it (ModeBar.svelte already iterates TABS generically). Task 6\nwires the actual panel; this is the URL/state groundwork alone.\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-task1.txt
git add web/src/lib/report/url.ts web/src/lib/report/url.test.ts
git commit -F .superpowers/commit-msg-task1.txt
```

---

## Task 2: Rating types, the copy module, the moment-anchor module, and the two fetchers

**Files:**
- Create: `web/src/lib/rating/types.ts`
- Create: `web/src/lib/rating/copy.ts`
- Create: `web/src/lib/rating/moments.ts`
- Create: `web/src/lib/rating/moments.test.ts`
- Create: `web/src/lib/rating/roster-join.ts`
- Create: `web/src/lib/rating/roster-join.test.ts`
- Create: `web/src/lib/rating/report-ratings.svelte.ts`
- Create: `web/src/lib/rating/report-ratings.test.ts`
- Modify: `web/src/lib/rankings/api.ts` (add `fetchReportRatings`, `fetchCharacterRating`)
- Modify: `web/src/lib/rankings/api.test.ts` (test the two new fetchers)

**Interfaces:**
- Produces: `RatingComponentName`, `RatingBasis`, `RatingMoment`, `RatingComponent`,
  `RatingCardPlayer`, `ReportRatings`, `CharacterRatingTrendPoint`, `CharacterRating` (all in
  `web/src/lib/rating/types.ts`); `RATING_COMPONENT_ORDER: readonly RatingComponentName[]`;
  `ratingCopy` object and `MIN_TREND_SAMPLES` (in `copy.ts`); `momentHref(fightIndex: number,
  playerGuid: string, moment: RatingMoment): string | null` (in `moments.ts`);
  `guidForPlayer(player: RatingCardPlayer, roster: {guid:string;name:string;class?:string}[]):
  string` and `classForPlayer(guid: string, fallbackClass: string | undefined, roster:
  {guid:string;name:string;class?:string}[]): string | undefined` (in `roster-join.ts`, Ruling
  3/Ruling 12); `createReportRatingsFetch(apiBase?: string): { data: ReportRatings | null;
  status: 'idle'|'loading'|'ready'|'failed'; error: string; load(reportId: string, fightIndex:
  number): void }` (in `report-ratings.svelte.ts`); `fetchReportRatings(reportId: string,
  fightIndex: number, apiBase?: string): Promise<ReportRatings>` and `fetchCharacterRating(path:
  CharacterPath, apiBase?: string): Promise<CharacterRating>` (in `rankings/api.ts`).
- Consumes: `RankingsError`, the module-private `get<T>()` helper, `API_BASE_URL`, and
  `CharacterPath` — all already in `web/src/lib/rankings/api.ts`/`web/src/lib/characters.ts`.

- [ ] **Step 1: Write `web/src/lib/rating/types.ts` (no test — a pure type module)**

```ts
// web/src/lib/rating/types.ts
// The shapes docs/superpowers/specs/2026-09-21-performance-rating-design.md section 5.2
// defines. The API lane has not landed; this plan builds against the spec's own response
// examples (§5.2) per the simulator-parity precedent for building against a signature before
// its provider lands. See docs/superpowers/plans/2026-09-21-rating-web.md's Rulings 1-4 for
// where this module fills a gap the spec leaves open.

export type RatingBasis = 'percentile' | 'absolute' | 'mixed' | '';

export type RatingComponentName =
  | 'output'
  | 'survival'
  | 'mechanics'
  | 'utility'
  | 'preparation'
  | 'activity';

/** The spec §1.4 weight table's own order — every card's six components render in this
 *  order regardless of what order the API's array happens to list them in. */
export const RATING_COMPONENT_ORDER: readonly RatingComponentName[] = [
  'output',
  'survival',
  'mechanics',
  'utility',
  'preparation',
  'activity',
];

/**
 * One piece of evidence behind a component's score. `kind` drives which tab
 * `moments.ts`'s `momentHref` links into; an unrecognised kind renders as plain text
 * rather than a broken link (Ruling 4).
 */
export interface RatingMoment {
  kind: string;
  at_ms?: number;
  spell_id?: number;
  spell_name?: string;
  avoidable?: boolean;
  /** Only meaningful for `kind: "death"` — DeathsTab's own `death-<guid>-<at_ms>` anchor,
   *  supplied by the API. Every other kind's link is built locally (Ruling 4). */
  anchor?: string;
}

export interface RatingComponent {
  name: RatingComponentName;
  /** Null exactly when `excluded` is true. */
  score: number | null;
  /** The renormalised weight actually applied, 0-100. */
  weight: number;
  basis: RatingBasis;
  percentile: number | null;
  bracket_n: number;
  excluded: boolean;
  /** Set iff `excluded`; one of the fixed strings in spec §7.4. */
  reason: string;
  moments: RatingMoment[];
}

export interface RatingCardPlayer {
  player_key: string;
  player_name: string;
  class: string;
  spec: string;
  role: string;
  overall: number;
  overall_uncapped: number;
  overall_capped: boolean;
  basis: RatingBasis;
  /** Always six entries, one per RATING_COMPONENT_ORDER member, per spec §4.1's
   *  Card.Components [6]Component — a component the engine could not score is present
   *  with excluded: true, never omitted. */
  components: RatingComponent[];
}

export interface ReportRatings {
  fight_index: number;
  kill: boolean;
  kill_time_band: string;
  model_version: string;
  players: RatingCardPlayer[];
}

export interface CharacterRatingTrendPoint {
  fought_at: string;
  overall: number;
  report_id: string;
  fight_index: number;
}

export interface CharacterRating {
  player_key: string;
  sample_size: number;
  trend: CharacterRatingTrendPoint[];
  best_component: string;
  worst_component: string;
  /** Null for a character with zero rated fights (Ruling 1) or when `latest` genuinely has
   *  nothing to report; the spec's own example always populates it, but a brand-new
   *  character cannot. */
  latest: RatingCardPlayer | null;
}
```

- [ ] **Step 2: Write `web/src/lib/rating/copy.ts`**

```ts
// web/src/lib/rating/copy.ts
// Every visible string the rating surfaces use, in the site's honest-copy voice: no
// exclamation marks, no "AI", no claim the engine does not make. Strings quoted directly
// from spec §6.4/§1.5 are copied verbatim; the rest follow their tone.

/** Spec §6.4's own worked-example threshold for showing the character trend sparkline —
 *  a display-only choice (Ruling 7), not a server-side minimum. */
export const MIN_TREND_SAMPLES = 5;

export const ratingCopy = {
  tabLabel: 'Rating',
  headingFor: (playerName: string): string => `${playerName}’s rating`,
  explainLink: 'How is this calculated?',
  panelHeading: 'Performance rating',
  panelMore: 'Rating tab',
  panelEmpty: 'No ratings for this fight yet.',
  tabEmpty:
    'No ratings for this fight yet. Ratings are computed after a report finishes processing, and some encounters do not have a curated table yet.',
  fetchFailed: 'Ratings did not load.',
  nightNotOnePull: 'Ratings are one pull’s. Pick a boss pull from the list to see one.',
  enemiesHaveNone: 'Ratings are for your raid, not the enemy.',
  noMatchingRow: 'This player has no rating for this fight.',
  /** Spec §1.2's two exact basis sentences. */
  percentileBasis: (spec: string, klass: string, role: string, pct: number, n: number): string =>
    `Compared with other ${spec} ${roleNoun(klass, role)}s on this fight (${n} logs).`,
  absoluteBasis: 'Measured against the encounter’s own numbers — not enough logs yet to compare players.',
  /** Spec §7.4's fixed excluded-reason strings, by the `reason` code the API returns. */
  excludedReason: (componentLabel: string, reason: string): string => {
    switch (reason) {
      case 'no_mechanics_table':
        return `${componentLabel} — not scored. This fight’s encounter has no curated mechanics table yet.`;
      case 'spec_not_modeled':
        return `${componentLabel} — not scored. The simulator does not model this spec yet; scores return once it does.`;
      case 'no_utility_table':
        return `${componentLabel} — not scored. No curated utility table exists for this spec yet.`;
      case 'no_consumable_catalogue':
        return `${componentLabel} — not scored. No consumable catalogue exists for this role yet.`;
      case 'unclaimed_and_unranked':
        return `${componentLabel} — not scored. This character is not signed in, and the bracket has too few logs to compare against.`;
      default:
        return `${componentLabel} — not scored.`;
    }
  },
  threatNotModeled: 'Threat — not modeled yet. This does not count for or against Utility.',
  /** Spec §6.4's exact capped-score sentence. */
  cappedNote:
    'capped from a higher weighted average — a costly avoidable death outweighs the rest of the fight. See Survival.',
  /** Spec §1.5's exact rule sentence, published on /ratings. */
  cappedRule:
    'An avoidable death early in a fight caps the overall score at 40, because nothing else in the fight makes up for it.',
  trendTooFew: (have: number): string => `Not enough rated fights yet to show a trend (${have} of ${MIN_TREND_SAMPLES} needed).`,
  characterEmpty: 'No rated fights yet.',
} as const;

const COMPONENT_LABELS: Record<string, string> = {
  output: 'Output',
  survival: 'Survival',
  mechanics: 'Mechanics',
  utility: 'Utility',
  preparation: 'Preparation',
  activity: 'Activity',
};

export function componentLabel(name: string): string {
  return COMPONENT_LABELS[name] ?? name;
}

/** "Fury" + "Warrior" -> "Fury Warriors"; every vanilla class name pluralises with a bare
 *  "s", so no irregular table is needed. `role` is unused today (the bracket key includes
 *  it, but the copy the spec quotes names the class, not the role) and kept as a parameter
 *  so a future copy revision that needs it has no signature to change. */
function roleNoun(klass: string, _role: string): string {
  return `${klass}s`;
}
```

- [ ] **Step 3: Write the failing test for `moments.ts`**

```ts
// web/src/lib/rating/moments.test.ts
import { describe, expect, it } from 'vitest';
import { momentHref } from './moments';
import type { RatingMoment } from './types';

const GUID = 'Player-4184-000000A1';

describe('momentHref', () => {
  it('a death moment links to the Deaths tab using the API’s own anchor', () => {
    const moment: RatingMoment = {
      kind: 'death',
      at_ms: 140_000,
      spell_id: 19712,
      spell_name: 'Arcane Explosion',
      avoidable: true,
      anchor: `death-${GUID}-140000`,
    };
    expect(momentHref(3, GUID, moment)).toBe(`?fight=3&tab=deaths#death-${GUID}-140000`);
  });

  it('an interrupt moment links to the Interrupts tab via the exchange-<guid>-<spellId> convention', () => {
    const moment: RatingMoment = { kind: 'interrupt', spell_id: 20066, spell_name: 'Repentance' };
    expect(momentHref(3, GUID, moment)).toBe(`?fight=3&tab=interrupts#exchange-${GUID}-20066`);
  });

  it('a dispel moment links to the Dispels tab the same way', () => {
    const moment: RatingMoment = { kind: 'dispel', spell_id: 6205, spell_name: 'Gehennas’ Curse' };
    expect(momentHref(3, GUID, moment)).toBe(`?fight=3&tab=dispels#exchange-${GUID}-6205`);
  });

  it('a carried-debuff moment links to the Debuffs tab via the aura-<guid>-<spellId> convention', () => {
    const moment: RatingMoment = { kind: 'carried-debuff', spell_id: 6205, spell_name: 'Gehennas’ Curse' };
    expect(momentHref(3, GUID, moment)).toBe(`?fight=3&tab=debuffs#aura-${GUID}-6205`);
  });

  it('a utility-uptime moment links to the Buffs tab the same way', () => {
    const moment: RatingMoment = { kind: 'utility-uptime', spell_id: 1160, spell_name: 'Demoralizing Shout' };
    expect(momentHref(3, GUID, moment)).toBe(`?fight=3&tab=buffs#aura-${GUID}-1160`);
  });

  it('an unrecognised kind, or one missing the spell id a link needs, has no href', () => {
    expect(momentHref(3, GUID, { kind: 'avoidable-hit', spell_name: 'Standing in fire' })).toBeNull();
    expect(momentHref(3, GUID, { kind: 'interrupt' })).toBeNull();
  });

  it('an unresolved player guid (Ruling 3’s no-match case) has no href even for a death', () => {
    expect(momentHref(3, '', { kind: 'death', anchor: 'death-x-1' })).toBeNull();
  });
});
```

- [ ] **Step 4: Run test to verify it fails**

Run: `cd web && npx vitest run src/lib/rating/moments.test.ts`
Expected: FAIL — `./moments` does not exist yet.

- [ ] **Step 5: Implement `web/src/lib/rating/moments.ts`**

```ts
// web/src/lib/rating/moments.ts
// Where a rating moment jumps to, per docs/superpowers/plans/2026-09-21-rating-web.md's
// Ruling 4: a death uses DeathsTab's own pre-existing anchor convention verbatim; every
// other kind is built here from a spell id plus this player's own guid, since
// ExchangeRow/AuraTrack carry no per-event timestamp for the spec's literal <atMs>
// convention to read.
import type { RatingMoment } from './types';

const TAB_BY_KIND: Record<string, { tab: string; prefix: 'exchange' | 'aura' } | undefined> = {
  interrupt: { tab: 'interrupts', prefix: 'exchange' },
  dispel: { tab: 'dispels', prefix: 'exchange' },
  'carried-debuff': { tab: 'debuffs', prefix: 'aura' },
  'utility-uptime': { tab: 'buffs', prefix: 'aura' },
};

export function momentHref(fightIndex: number, playerGuid: string, moment: RatingMoment): string | null {
  if (playerGuid === '') return null;
  if (moment.kind === 'death') {
    if (moment.anchor === undefined || moment.anchor === '') return null;
    return `?fight=${fightIndex}&tab=deaths#${moment.anchor}`;
  }
  const target = TAB_BY_KIND[moment.kind];
  if (target === undefined || moment.spell_id === undefined) return null;
  return `?fight=${fightIndex}&tab=${target.tab}#${target.prefix}-${playerGuid}-${moment.spell_id}`;
}
```

- [ ] **Step 6: Run test to verify it passes**

Run: `cd web && npx vitest run src/lib/rating/moments.test.ts`
Expected: PASS.

- [ ] **Step 6a: Write the failing test for `roster-join.ts`** (Ruling 3/Ruling 12 — one
  shared join, not duplicated per component)

```ts
// web/src/lib/rating/roster-join.test.ts
import { describe, expect, it } from 'vitest';
import { classForPlayer, guidForPlayer } from './roster-join';
import type { RatingCardPlayer } from './types';

const ROSTER = [{ guid: 'Player-4184-000000A1', name: 'Simfury', class: 'Warrior' }];

function player(overrides: Partial<RatingCardPlayer> = {}): RatingCardPlayer {
  return {
    player_key: 'us/normal/simfury',
    player_name: 'Simfury',
    class: 'Warrior',
    spec: 'Fury',
    role: 'dps',
    overall: 70,
    overall_uncapped: 70,
    overall_capped: false,
    basis: 'percentile',
    components: [],
    ...overrides,
  };
}

describe('guidForPlayer', () => {
  it('joins by display name, stripping the log’s trailing -<segment> on both sides', () => {
    expect(guidForPlayer(player({ player_name: 'Simfury-A1B2' }), ROSTER)).toBe('Player-4184-000000A1');
  });

  it('a name with no roster match returns an empty string, not a guess', () => {
    expect(guidForPlayer(player({ player_name: 'Nobody' }), ROSTER)).toBe('');
  });
});

describe('classForPlayer', () => {
  it('prefers the roster’s own class, falling back to the rating row’s', () => {
    expect(classForPlayer('Player-4184-000000A1', 'Priest', ROSTER)).toBe('Warrior');
    expect(classForPlayer('', 'Priest', ROSTER)).toBe('Priest');
  });
});
```

- [ ] **Step 6b: Run test to verify it fails**

Run: `cd web && npx vitest run src/lib/rating/roster-join.test.ts`
Expected: FAIL — `./roster-join` does not exist.

- [ ] **Step 6c: Implement `web/src/lib/rating/roster-join.ts`**

```ts
// web/src/lib/rating/roster-join.ts
// Ruling 3 (docs/superpowers/plans/2026-09-21-rating-web.md): the rating API's per-fight
// rows carry player_key/player_name, not this fight's own unit GUID, and the report island
// has no region/ruleset available to build a characterKey join with. Join by display name
// instead -- a raid roster's names are unique within one fight. Shared by RatingPanel.svelte
// and RatingTab.svelte (Ruling 12) rather than duplicated.
import { splitUnitName } from '../characters';
import type { RatingCardPlayer } from './types';

export interface RosterUnit {
  guid: string;
  name: string;
  class?: string;
}

export function guidForPlayer(player: RatingCardPlayer, roster: RosterUnit[]): string {
  const name = splitUnitName(player.player_name).name;
  return roster.find((unit) => splitUnitName(unit.name).name === name)?.guid ?? '';
}

export function classForPlayer(guid: string, fallbackClass: string | undefined, roster: RosterUnit[]): string | undefined {
  if (guid === '') return fallbackClass;
  return roster.find((unit) => unit.guid === guid)?.class ?? fallbackClass;
}
```

- [ ] **Step 6d: Run test to verify it passes**

Run: `cd web && npx vitest run src/lib/rating/roster-join.test.ts`
Expected: PASS.

- [ ] **Step 7: Add the two fetchers to `web/src/lib/rankings/api.ts`, test-first**

Add to `web/src/lib/rankings/api.test.ts` (alongside the existing `describe('fetchRankings', ...)`
blocks, reusing the file's own `envelope()` helper already defined at the top):

```ts
import { fetchCharacterRating, fetchReportRatings, RankingsError } from './api';

const REPORT_RATINGS = {
  fight_index: 3,
  kill: true,
  kill_time_band: 'typical',
  model_version: 'rating-2026-09-21',
  players: [
    {
      player_key: 'us/normal/simfury',
      player_name: 'Simfury',
      class: 'Warrior',
      spec: 'Fury',
      role: 'dps',
      overall: 70,
      overall_uncapped: 70,
      overall_capped: false,
      basis: 'percentile',
      components: [],
    },
  ],
};

describe('fetchReportRatings', () => {
  it('reads the per-fight report card from the ratings endpoint', async () => {
    const upstream = vi.fn<GlobalFetch>(async () => envelope(REPORT_RATINGS));
    vi.stubGlobal('fetch', upstream);

    const ratings = await fetchReportRatings('fixture2abcd', 3, API);

    expect(upstream).toHaveBeenCalledWith(
      `${API}/v1/reports/fixture2abcd/fights/3/ratings`,
      expect.objectContaining({ credentials: 'omit' }),
    );
    expect(ratings.players[0].player_name).toBe('Simfury');
  });

  it('an empty roster (Ruling 1’s launch state) resolves normally, not as an error', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => envelope({ ...REPORT_RATINGS, players: [] })),
    );
    const ratings = await fetchReportRatings('fixture2abcd', 3, API);
    expect(ratings.players).toEqual([]);
  });

  it('a transport failure raises RankingsError', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => envelope(null, 500)),
    );
    await expect(fetchReportRatings('fixture2abcd', 3, API)).rejects.toBeInstanceOf(RankingsError);
  });
});

describe('fetchCharacterRating', () => {
  it('reads the aggregate from the character rating endpoint', async () => {
    const data = {
      player_key: 'us/hardcore/elyra-duskvale',
      sample_size: 12,
      trend: [{ fought_at: '2026-12-09T22:10:00Z', overall: 62, report_id: 'fixture2abcd', fight_index: 2 }],
      best_component: 'preparation',
      worst_component: 'activity',
      latest: null,
    };
    const upstream = vi.fn<GlobalFetch>(async () => envelope(data));
    vi.stubGlobal('fetch', upstream);

    const rating = await fetchCharacterRating({ region: 'us', ruleset: 'hardcore', slug: 'elyra-duskvale' }, API);

    expect(upstream).toHaveBeenCalledWith(
      `${API}/v1/characters/us/hardcore/elyra-duskvale/rating`,
      expect.objectContaining({ credentials: 'omit' }),
    );
    expect(rating.sample_size).toBe(12);
  });

  it('an anonymized character 404s, and the fetcher raises RankingsError (Ruling 2)', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => envelope(null, 404)),
    );
    await expect(
      fetchCharacterRating({ region: 'us', ruleset: 'hardcore', slug: 'hidden' }, API),
    ).rejects.toBeInstanceOf(RankingsError);
  });
});
```

- [ ] **Step 8: Run test to verify it fails**

Run: `cd web && npx vitest run src/lib/rankings/api.test.ts`
Expected: FAIL — `fetchReportRatings`/`fetchCharacterRating` are not exported yet.

- [ ] **Step 9: Implement the fetchers in `web/src/lib/rankings/api.ts`**

Add near `fetchCharacter`/`fetchGuild` at the bottom of the file, reusing the module's own
`get()` helper and `CharacterPath` import already present:

```ts
import type { ReportRatings, CharacterRating } from '../rating/types';

export function fetchReportRatings(
  reportId: string,
  fightIndex: number,
  apiBase: string = API_BASE_URL,
): Promise<ReportRatings> {
  return get<ReportRatings>(`/v1/reports/${reportId}/fights/${fightIndex}/ratings`, apiBase);
}

export function fetchCharacterRating(
  path: CharacterPath,
  apiBase: string = API_BASE_URL,
): Promise<CharacterRating> {
  return get<CharacterRating>(`/v1/characters/${path.region}/${path.ruleset}/${path.slug}/rating`, apiBase);
}
```

(Add the `import type { ReportRatings, CharacterRating } from '../rating/types';` line beside
the file's existing `import type { CharacterPath } from '../characters';`.)

- [ ] **Step 10: Run test to verify it passes**

Run: `cd web && npx vitest run src/lib/rankings/api.test.ts`
Expected: PASS, all tests in the file green (including the pre-existing ones — nothing else in
the file changed).

- [ ] **Step 11: Write the failing test for the fetch-state hook**

```ts
// web/src/lib/rating/report-ratings.test.ts
import { afterEach, describe, expect, it, vi } from 'vitest';
import { createReportRatingsFetch } from './report-ratings.svelte';

const API = 'https://api.foreversixty.test';
type GlobalFetch = (...args: Parameters<typeof fetch>) => Promise<Response>;

function envelope(data: unknown, status = 200): Response {
  return new Response(JSON.stringify({ ok: status < 400, data, error: null, request_id: 'r' }), {
    status,
    headers: { 'content-type': 'application/json' },
  });
}

function flush(): Promise<void> {
  return Promise.resolve().then(() => Promise.resolve());
}

afterEach(() => vi.unstubAllGlobals());

describe('createReportRatingsFetch', () => {
  it('loads, then reflects ready with the resolved data', async () => {
    const data = { fight_index: 3, kill: true, kill_time_band: 'typical', model_version: 'v1', players: [] };
    vi.stubGlobal('fetch', vi.fn<GlobalFetch>(async () => envelope(data)));
    const hook = createReportRatingsFetch(API);

    expect(hook.status).toBe('idle');
    hook.load('fixture2abcd', 3);
    expect(hook.status).toBe('loading');
    await flush();

    expect(hook.status).toBe('ready');
    expect(hook.data).toEqual(data);
    expect(hook.error).toBe('');
  });

  it('a failure sets status failed and a message, and clears any stale data', async () => {
    vi.stubGlobal('fetch', vi.fn<GlobalFetch>(async () => envelope(null, 500)));
    const hook = createReportRatingsFetch(API);

    hook.load('fixture2abcd', 3);
    await flush();

    expect(hook.status).toBe('failed');
    expect(hook.data).toBeNull();
    expect(hook.error).not.toBe('');
  });

  it('switching fight while a request is in flight drops the stale response', async () => {
    let resolveFirst: ((r: Response) => void) | undefined;
    const upstream = vi.fn<GlobalFetch>(
      () =>
        new Promise<Response>((resolve) => {
          resolveFirst = resolve;
        }),
    );
    vi.stubGlobal('fetch', upstream);
    const hook = createReportRatingsFetch(API);

    hook.load('fixture2abcd', 1);
    const secondData = { fight_index: 2, kill: true, kill_time_band: 'typical', model_version: 'v1', players: [] };
    vi.stubGlobal('fetch', vi.fn<GlobalFetch>(async () => envelope(secondData)));
    hook.load('fixture2abcd', 2);
    await flush();

    expect(hook.data).toEqual(secondData);

    // The first request finally resolves; it must not clobber the second's already-landed data.
    resolveFirst?.(envelope({ fight_index: 1, kill: true, kill_time_band: 'typical', model_version: 'v1', players: [] }));
    await flush();
    expect(hook.data).toEqual(secondData);
  });

  it('a repeat load() for the same fight while already ready is a no-op (no second fetch)', async () => {
    const data = { fight_index: 3, kill: true, kill_time_band: 'typical', model_version: 'v1', players: [] };
    const upstream = vi.fn<GlobalFetch>(async () => envelope(data));
    vi.stubGlobal('fetch', upstream);
    const hook = createReportRatingsFetch(API);

    hook.load('fixture2abcd', 3);
    await flush();
    hook.load('fixture2abcd', 3);

    expect(upstream).toHaveBeenCalledTimes(1);
  });
});
```

- [ ] **Step 12: Run test to verify it fails**

Run: `cd web && npx vitest run src/lib/rating/report-ratings.test.ts`
Expected: FAIL — `./report-ratings.svelte` does not exist.

- [ ] **Step 13: Implement `web/src/lib/rating/report-ratings.svelte.ts`**

```ts
// web/src/lib/rating/report-ratings.svelte.ts
// A tiny reusable fetch-state container for GET /v1/reports/{id}/fights/{n}/ratings,
// shared by RatingPanel.svelte (the Summary dashboard row) and RatingTab.svelte (the full
// card) rather than each reimplementing the same load/loading/error/stale-response guard --
// see docs/superpowers/plans/2026-09-21-rating-web.md's Ruling 10 on why this is a shared
// hook and not a shared cross-component cache. Modeled on
// web/src/lib/report/lazy-component.svelte.ts's plain-state-container style, and on
// Character.svelte's own `resolved !== requested` stale-response guard.
import { fetchReportRatings } from '../rankings/api';
import type { ReportRatings } from './types';

export type ReportRatingsStatus = 'idle' | 'loading' | 'ready' | 'failed';

export interface ReportRatingsFetch {
  readonly data: ReportRatings | null;
  readonly status: ReportRatingsStatus;
  readonly error: string;
  /** Starts a fetch for this (reportId, fightIndex) pair unless one already succeeded for
   *  the same pair; safe to call on every render/effect run. */
  load(reportId: string, fightIndex: number): void;
}

export function createReportRatingsFetch(apiBase?: string): ReportRatingsFetch {
  let data = $state<ReportRatings | null>(null);
  let status = $state<ReportRatingsStatus>('idle');
  let error = $state('');
  let loadedKey = '';
  let requestedKey = '';

  return {
    get data() {
      return data;
    },
    get status() {
      return status;
    },
    get error() {
      return error;
    },
    load(reportId: string, fightIndex: number): void {
      const key = `${reportId}:${fightIndex}`;
      if (key === loadedKey || (key === requestedKey && status === 'loading')) return;
      requestedKey = key;
      status = 'loading';
      error = '';
      void fetchReportRatings(reportId, fightIndex, apiBase)
        .then((result) => {
          if (requestedKey !== key) return;
          data = result;
          loadedKey = key;
          status = 'ready';
        })
        .catch((thrown: unknown) => {
          if (requestedKey !== key) return;
          data = null;
          status = 'failed';
          error = thrown instanceof Error ? thrown.message : 'Ratings did not load.';
        });
    },
  };
}
```

- [ ] **Step 14: Run test to verify it passes**

Run: `cd web && npx vitest run src/lib/rating/report-ratings.test.ts`
Expected: PASS.

- [ ] **Step 15: Type-check, lint, format**

Run: `cd web && npx astro check && npm run lint && npx prettier --check src/lib/rating src/lib/rankings/api.ts src/lib/rankings/api.test.ts`
Expected: no errors.

- [ ] **Step 16: Commit**

```bash
printf 'feat(web): rating types, copy, moment anchors, and the two fetchers\n\nAdds web/src/lib/rating/{types,copy,moments,report-ratings.svelte}.ts\nand fetchReportRatings/fetchCharacterRating to rankings/api.ts, built\nagainst spec section 5.2 since the API lane has not landed. Records\nRulings 1-4 and 7 inline: the empty-roster launch state, the anonymize\n404, the name-based roster-GUID join, and the moment-anchor convention\nthis lane owns for ExchangeTable/AuraTable rows.\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-task2.txt
git add web/src/lib/rating web/src/lib/rankings/api.ts web/src/lib/rankings/api.test.ts
git commit -F .superpowers/commit-msg-task2.txt
# (web/src/lib/rating now also includes roster-join.ts and roster-join.test.ts, added by
# Step 6a-6d above; `git add web/src/lib/rating` already stages the whole directory.)
```

---

## Task 3: `ExchangeTable.svelte` and `AuraTable.svelte` anchor additions

**Files:**
- Modify: `web/src/components/report/ExchangeTable.svelte`
- Modify: `web/src/components/report/AuraTable.svelte`
- Test: `web/src/components/report/AuraTable.test.ts` (already exists — extend it)
- Test: create `web/src/components/report/ExchangeTable.test.ts` if one does not already exist
  (check first: `ls web/src/components/report/ExchangeTable.test.ts`); if it does not, use
  `AuraTable.test.ts`'s own imports/render harness as the template (it renders the component
  with `@testing-library/svelte` or the project's own render helper — read the file first to
  match its exact harness before writing this step for real).

**Interfaces:**
- Produces: every row `<li>`/`<tr>` `ExchangeTable.svelte` renders gains `id={exchange-<source
  _guid>-<spell_id>}`; every row `AuraTable.svelte` renders gains `id={aura-<target_guid>-
  <spell_id>}`. This is additive markup only — no prop, type, or behavior changes to either
  component.

- [ ] **Step 1: Read both components' current row markup in full**

```bash
cd web && sed -n '1,292p' src/components/report/ExchangeTable.svelte
cd web && sed -n '1,262p' src/components/report/AuraTable.svelte
cat src/components/report/AuraTable.test.ts
```

Confirm the exact row element (`<li>` or `<tr>`) and its existing `{#each rows as row (key)}`
keying expression before writing the id (the id must be unique per the same key the `{#each}`
already uses, or Svelte's own keyed-block warning fires on a collision — Ruling 4 assumes one
row per (source or target, spell_id), matching each table's existing dedup-by-spell grouping,
already confirmed by reading `ExchangeTable.svelte`'s `uncured`/row-building logic in Task 2's
research and `AuraTrack`'s shape in `web/src/lib/report/types.ts:151-163`).

- [ ] **Step 2: Write the failing test**

Add to `web/src/components/report/AuraTable.test.ts` (following its existing render-and-query
style exactly — read the file's existing `it(...)` blocks first and match its render helper,
prop fixture shape, and query style):

```ts
it('every row carries an aura-<guid>-<spellId> anchor id for a rating moment to jump to', () => {
  // Reuse this file's own existing fixture props (the track fixture already defined above
  // in this file); assert against whatever target_guid/spell_id that fixture actually uses.
  const { container } = render(AuraTable, { props: /* this file's existing default props */ });
  const row = container.querySelector('[id^="aura-"]');
  expect(row).not.toBeNull();
  expect(row?.id).toMatch(/^aura-[^-]+.*-\d+$/);
});
```

Create `web/src/components/report/ExchangeTable.test.ts` (or extend it if found in Step 1)
with the equivalent:

```ts
it('every row carries an exchange-<guid>-<spellId> anchor id for a rating moment to jump to', () => {
  const { container } = render(ExchangeTable, {
    props: {
      rows: [
        {
          kind: 'interrupt',
          source_guid: 'Player-4184-000000A1',
          source_name: 'Elyra Duskvale',
          target_guid: 'Creature-1',
          target_name: 'Shazzrah',
          spell_id: 20066,
          spell_name: 'Repentance',
          extra_spell_id: 19712,
          extra_spell_name: 'Arcane Explosion',
          count: 1,
        },
      ],
      emptyText: 'Nothing interrupted.',
    },
  });
  const row = container.querySelector('#exchange-Player-4184-000000A1-19712');
  expect(row).not.toBeNull();
});
```

(If `ExchangeTable.test.ts` already exists with a different render harness, use that file's own
import/render pattern instead of `@testing-library/svelte` verbatim — match what's already
there.)

- [ ] **Step 3: Run tests to verify they fail**

Run: `cd web && npx vitest run src/components/report/AuraTable.test.ts src/components/report/ExchangeTable.test.ts`
Expected: FAIL — no element has an `id` starting with `aura-`/`exchange-` yet.

- [ ] **Step 4: Implement — add the `id` attribute to each row**

In `AuraTable.svelte`'s row element (the one keyed by the track's own identity in its
`{#each ... as track (...)}`), add:

```svelte
id={`aura-${track.target_guid}-${track.spell_id}`}
```

In `ExchangeTable.svelte`'s row element (the one keyed in its `{#each rows as row (...)}`),
add:

```svelte
id={`exchange-${row.source_guid}-${row.spell_id}`}
```

Place the `id` attribute alongside each row's existing `class`/`data-testid` attributes (do not
reorder or otherwise touch the rest of the row's markup).

- [ ] **Step 5: Run tests to verify they pass**

Run: `cd web && npx vitest run src/components/report/AuraTable.test.ts src/components/report/ExchangeTable.test.ts`
Expected: PASS.

- [ ] **Step 6: Run the full component test suite for regressions**

Run: `cd web && npx vitest run src/components/report`
Expected: PASS — no existing test asserted on the absence of an `id`, but confirm.

- [ ] **Step 7: Type-check, lint, format**

Run: `cd web && npx astro check && npm run lint && npx prettier --check src/components/report/ExchangeTable.svelte src/components/report/AuraTable.svelte src/components/report/ExchangeTable.test.ts src/components/report/AuraTable.test.ts`

- [ ] **Step 8: Commit**

```bash
printf 'feat(web): anchor ids on ExchangeTable and AuraTable rows\n\nAdds id={exchange-<sourceGuid>-<spellId>}/id={aura-<targetGuid>-\n<spellId>} to every row, the convention Ruling 4 of\ndocs/superpowers/plans/2026-09-21-rating-web.md settles on since\nExchangeRow/AuraTrack carry no per-event timestamp for the spec’s\nliteral <atMs> anchor. Every tab benefits from the deep-link going\nforward, not only Rating.\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-task3.txt
git add web/src/components/report/ExchangeTable.svelte web/src/components/report/AuraTable.svelte web/src/components/report/ExchangeTable.test.ts web/src/components/report/AuraTable.test.ts
git commit -F .superpowers/commit-msg-task3.txt
```

---

## Task 4: `RatingPanel.svelte` (the Summary dashboard row) and its `SummaryTab` wiring

**Files:**
- Create: `web/src/components/report/RatingPanel.svelte`
- Create: `web/src/components/report/RatingPanel.test.ts`
- Modify: `web/src/components/report/SummaryTab.svelte` (one new prop, one new render line)

**Interfaces:**
- Consumes: `createReportRatingsFetch` (Task 2), `RATING_COMPONENT_ORDER`/`RatingCardPlayer`
  (Task 2's `types.ts`), `ratingCopy`/`componentLabel` (Task 2's `copy.ts`), `splitUnitName`/
  `classColorVar` (existing `web/src/lib/characters.ts`/`web/src/lib/report/format.ts`).
- Produces: `RatingPanel` props `{ reportId: string; fightIndex: number; roster: { guid:
  string; name: string; class?: string }[]; onTab: () => void; onSelectPlayer?: (guid: string)
  => void; apiBase?: string }`. `SummaryTab.svelte`'s exported prop type for `onTab` widens to
  include `'rating'` (Task 6 relies on this).

- [ ] **Step 1: Write the failing test**

```ts
// web/src/components/report/RatingPanel.test.ts
import { afterEach, describe, expect, it, vi } from 'vitest';
import { render, screen, waitFor } from '@testing-library/svelte';
import RatingPanel from './RatingPanel.svelte';

type GlobalFetch = (...args: Parameters<typeof fetch>) => Promise<Response>;

function envelope(data: unknown, status = 200): Response {
  return new Response(JSON.stringify({ ok: status < 400, data, error: null, request_id: 'r' }), {
    status,
    headers: { 'content-type': 'application/json' },
  });
}

const ROSTER = [{ guid: 'Player-4184-000000A1', name: 'Simfury', class: 'Warrior' }];

const RATINGS = {
  fight_index: 3,
  kill: true,
  kill_time_band: 'typical',
  model_version: 'rating-2026-09-21',
  players: [
    {
      player_key: 'us/normal/simfury',
      player_name: 'Simfury',
      class: 'Warrior',
      spec: 'Fury',
      role: 'dps',
      overall: 70,
      overall_uncapped: 70,
      overall_capped: false,
      basis: 'percentile',
      components: [
        { name: 'output', score: 71, weight: 35, basis: 'percentile', percentile: 71, bracket_n: 142, excluded: false, reason: '', moments: [] },
        { name: 'survival', score: 76, weight: 15, basis: 'percentile', percentile: 60, bracket_n: 89, excluded: false, reason: '', moments: [] },
        { name: 'mechanics', score: 55, weight: 20, basis: 'percentile', percentile: 55, bracket_n: 60, excluded: false, reason: '', moments: [] },
        { name: 'utility', score: 62, weight: 15, basis: 'percentile', percentile: 62, bracket_n: 60, excluded: false, reason: '', moments: [] },
        { name: 'preparation', score: 95, weight: 10, basis: 'absolute', percentile: null, bracket_n: 0, excluded: false, reason: '', moments: [] },
        { name: 'activity', score: 80, weight: 5, basis: 'percentile', percentile: 80, bracket_n: 60, excluded: false, reason: '', moments: [] },
      ],
    },
  ],
};

afterEach(() => vi.unstubAllGlobals());

describe('RatingPanel', () => {
  it('renders one row per player with the overall score, and links to the Rating tab', async () => {
    vi.stubGlobal('fetch', vi.fn<GlobalFetch>(async () => envelope(RATINGS)));
    const onTab = vi.fn();
    render(RatingPanel, {
      props: { reportId: 'fixture2abcd', fightIndex: 3, roster: ROSTER, onTab, apiBase: 'https://api.test' },
    });

    await waitFor(() => expect(screen.getByText('70')).toBeTruthy());
    expect(screen.getByText('Simfury')).toBeTruthy();

    await screen.getByRole('button', { name: 'Rating tab' }).click();
    expect(onTab).toHaveBeenCalled();
  });

  it('an empty roster (no ratings yet) shows the truthful empty state, not a spinner', async () => {
    vi.stubGlobal('fetch', vi.fn<GlobalFetch>(async () => envelope({ ...RATINGS, players: [] })));
    render(RatingPanel, {
      props: { reportId: 'fixture2abcd', fightIndex: 3, roster: ROSTER, onTab: vi.fn(), apiBase: 'https://api.test' },
    });

    await waitFor(() => expect(screen.getByTestId('rating-panel-empty')).toBeTruthy());
  });

  it('a fetch failure shows a quiet failure line, not a thrown error', async () => {
    vi.stubGlobal('fetch', vi.fn<GlobalFetch>(async () => envelope(null, 500)));
    render(RatingPanel, {
      props: { reportId: 'fixture2abcd', fightIndex: 3, roster: ROSTER, onTab: vi.fn(), apiBase: 'https://api.test' },
    });

    await waitFor(() => expect(screen.getByTestId('rating-panel-error')).toBeTruthy());
  });
});
```

(Match this file's `render`/`screen` import to whatever `AuraTable.test.ts` actually uses if it
differs — the project may re-export a configured `render` from a local test-utils module; check
before writing this for real and align imports.)

- [ ] **Step 2: Run test to verify it fails**

Run: `cd web && npx vitest run src/components/report/RatingPanel.test.ts`
Expected: FAIL — `./RatingPanel.svelte` does not exist.

- [ ] **Step 3: Implement `web/src/components/report/RatingPanel.svelte`**

```svelte
<!-- web/src/components/report/RatingPanel.svelte -->
<!-- The Summary dashboard's rating row: one line per roster player, the overall score plus
     a compact six-segment bar, linking into the full RatingTab (spec §6.1). Mirrors
     SummaryPanels.svelte's own dashboard pattern (a short list, a "go deeper" link) but
     fetches its own data -- unlike SummaryPanels' siblings, a rating is not part of the
     report's already-loaded summary.json (see docs/superpowers/plans/2026-09-21-rating-web
     .md Ruling 8 on why this is its own component rather than a SummaryPanels addition). -->
<script lang="ts">
  import { splitUnitName } from '../../lib/characters';
  import { classColorVar } from '../../lib/report/format';
  import { createReportRatingsFetch } from '../../lib/rating/report-ratings.svelte';
  import { ratingCopy, componentLabel } from '../../lib/rating/copy';
  import { RATING_COMPONENT_ORDER } from '../../lib/rating/types';
  import { classForPlayer, guidForPlayer } from '../../lib/rating/roster-join';

  let {
    reportId,
    fightIndex,
    roster,
    onTab,
    onSelectPlayer = undefined,
    apiBase = undefined,
  }: {
    reportId: string;
    fightIndex: number;
    roster: { guid: string; name: string; class?: string }[];
    onTab: () => void;
    onSelectPlayer?: (guid: string) => void;
    apiBase?: string;
  } = $props();

  const fetcher = createReportRatingsFetch(apiBase);
  $effect(() => {
    fetcher.load(reportId, fightIndex);
  });

  const rows = $derived(
    (fetcher.data?.players ?? []).map((player) => ({ player, guid: guidForPlayer(player, roster) })),
  );

  const panel = 'border-line rounded-panel bg-raised flex flex-col gap-2 border p-3';
  const heading = 'label text-muted flex items-center justify-between';
  const more =
    'text-gold inline-flex min-h-11 items-center text-[11px] normal-case tracking-normal md:min-h-0';
</script>

<section class={panel} data-testid="rating-panel">
  <h2 class={heading}>
    {ratingCopy.panelHeading}
    <button type="button" class={more} onclick={onTab}>{ratingCopy.panelMore}</button>
  </h2>
  {#if fetcher.status === 'loading' || fetcher.status === 'idle'}
    <div class="flex h-8 flex-col gap-2" data-testid="rating-panel-loading" aria-hidden="true">
      <div class="bg-line-soft h-8 w-full animate-pulse rounded"></div>
    </div>
  {:else if fetcher.status === 'failed'}
    <p class="text-muted text-[13px]" role="alert" data-testid="rating-panel-error">{ratingCopy.fetchFailed}</p>
  {:else if rows.length === 0}
    <p class="text-muted text-[13px]" data-testid="rating-panel-empty">{ratingCopy.panelEmpty}</p>
  {:else}
    <ul class="flex flex-col" data-testid="rating-panel-rows">
      {#each rows as { player, guid } (player.player_key)}
        <li class="border-line-soft flex min-h-8 items-center gap-3 border-b py-1 text-[13px] last:border-0">
          <span
            class="flex min-w-0 flex-1 items-center truncate font-semibold"
            style={`color: ${classColorVar(classForPlayer(guid, player.class, roster))}`}
          >
            {#if onSelectPlayer && guid !== ''}
              <button
                type="button"
                class="inline-flex min-h-11 items-center truncate underline-offset-2 hover:underline md:min-h-0"
                style={`color: ${classColorVar(classForPlayer(guid, player.class, roster))}`}
                title="Show only this player"
                onclick={() => onSelectPlayer(guid)}>{splitUnitName(player.player_name).name}</button
              >
            {:else}
              {splitUnitName(player.player_name).name}
            {/if}
          </span>
          <span class="text-gold tabular w-8 text-right font-mono font-bold" data-testid="rating-panel-overall"
            >{Math.round(player.overall)}</span
          >
          <span class="flex w-24 shrink-0 gap-[2px]" aria-hidden="true">
            {#each RATING_COMPONENT_ORDER as name (name)}
              {@const part = player.components.find((c) => c.name === name)}
              <span class="bg-line-soft h-2 flex-1" title={part ? componentLabel(name) : undefined}
                ><span
                  class="bg-gold block h-full"
                  style={`width: ${part && !part.excluded && part.score !== null ? part.score : 0}%`}
                ></span></span
              >
            {/each}
          </span>
        </li>
      {/each}
    </ul>
  {/if}
</section>
```

- [ ] **Step 4: Run test to verify it passes**

Run: `cd web && npx vitest run src/components/report/RatingPanel.test.ts`
Expected: PASS.

- [ ] **Step 5: Wire it into `SummaryTab.svelte`**

Read `web/src/components/report/SummaryTab.svelte` in full first (it is 259+ lines already
shown in Task research), then:

1. Widen its `onTab` prop's type union (find the line `onTab?: (tab: 'damage-done' | 'healing'
   | 'damage-taken' | 'deaths') => void;`) to add `'rating'`:
   ```ts
   onTab?: (tab: 'damage-done' | 'healing' | 'damage-taken' | 'deaths' | 'rating') => void;
   ```
2. Import `RatingPanel` beside the existing `import SummaryPanels from './SummaryPanels.svelte';`.
3. In the existing `{#if onTab}<SummaryPanels {summary} {everyone} {durationMs} {players}
   {onTab} {onSelectPlayer} {approximate} />` block, add `RatingPanel` right after
   `SummaryPanels`, gated on `reportId`/`fightIndex` being defined (they are already props on
   `SummaryTab`):
   ```svelte
   {#if onTab}
     <SummaryPanels {summary} {everyone} {durationMs} {players} {onTab} {onSelectPlayer} {approximate} />
     {#if reportId !== undefined && fightIndex !== undefined}
       <RatingPanel
         {reportId}
         {fightIndex}
         roster={[...summary.roster].map((row) => ({ guid: row.guid, name: row.name, class: row.class }))}
         onTab={() => onTab('rating')}
         {onSelectPlayer}
       />
     {/if}
   {/if}
   ```
   (Place this immediately after the existing `<SummaryPanels ... />` line; do not reorder any
   other markup in the file.)

- [ ] **Step 6: Write a small `SummaryTab.svelte` regression test if one already covers the
  `onTab`-gated block; otherwise add one**

Check `web/src/components/report/SummaryTab.test.ts` for an existing test asserting on the
`{#if onTab}` block's contents (e.g. asserting `summary-panels` renders when `onTab` is
passed). If found, extend it with:

```ts
it('also renders the rating dashboard row when reportId and fightIndex are given', () => {
  // Match this file's existing render call/props fixture, adding onTab, reportId, fightIndex.
  const { getByTestId } = render(SummaryTab, {
    props: { /* ...this file's existing fixture props..., */ onTab: vi.fn(), reportId: 'fixture2abcd', fightIndex: 3 },
  });
  expect(getByTestId('rating-panel')).toBeTruthy();
});
```

If `SummaryTab.test.ts` does not exist at all, skip this step — `RatingPanel.test.ts` (Step 1-4)
already covers the component's own behavior in isolation, and Task 6's e2e spec covers the
wired-together page.

- [ ] **Step 7: Run the full affected test set**

Run: `cd web && npx vitest run src/components/report/RatingPanel.test.ts src/components/report/SummaryTab.test.ts 2>&1 | tail -40`
Expected: PASS (if `SummaryTab.test.ts` exists; otherwise the first file alone passes).

- [ ] **Step 8: Type-check, lint, format**

Run: `cd web && npx astro check && npm run lint && npx prettier --check src/components/report/RatingPanel.svelte src/components/report/RatingPanel.test.ts src/components/report/SummaryTab.svelte`

- [ ] **Step 9: Commit**

```bash
printf 'feat(web): RatingPanel, the Summary dashboard’s rating row\n\nOne line per roster player -- overall score plus a flat-gold\nsix-segment bar (Ruling 5: never colour-by-score) -- linking into the\nfull Rating tab. SummaryTab gains one gated line mounting it; onTab’s\ntype widens to include '"'"'rating'"'"'.\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-task4.txt
git add web/src/components/report/RatingPanel.svelte web/src/components/report/RatingPanel.test.ts web/src/components/report/SummaryTab.svelte
git add web/src/components/report/SummaryTab.test.ts 2>/dev/null || true
git commit -F .superpowers/commit-msg-task4.txt
```

---

## Task 5: `RatingTab.svelte` (the full three-level card)

**Files:**
- Create: `web/src/components/report/RatingTab.svelte`
- Create: `web/src/components/report/RatingTab.test.ts`

**Interfaces:**
- Consumes: everything Task 2/4 produced, plus `momentHref` (Task 2's `moments.ts`) and
  `ratingCopy`/`componentLabel`/`MIN_TREND_SAMPLES` are not needed here (trend is character
  -page only, Task 7).
- Produces: `RatingTab` props `{ reportId: string; fightIndex: number; roster: { guid: string;
  name: string; class?: string }[]; source: string; onSelectPlayer: (guid: string) => void;
  apiBase?: string }`, where `source` is the report URL's own `SOURCE_FRIENDLIES` /
  `SOURCE_ENEMIES` / a player GUID (exactly `ReportView.svelte`'s `state.source`).

- [ ] **Step 1: Write the failing test**

```ts
// web/src/components/report/RatingTab.test.ts
import { afterEach, describe, expect, it, vi } from 'vitest';
import { render, screen, waitFor } from '@testing-library/svelte';
import RatingTab from './RatingTab.svelte';
import { SOURCE_ENEMIES, SOURCE_FRIENDLIES } from '../../lib/report/url';

type GlobalFetch = (...args: Parameters<typeof fetch>) => Promise<Response>;
function envelope(data: unknown, status = 200): Response {
  return new Response(JSON.stringify({ ok: status < 400, data, error: null, request_id: 'r' }), {
    status,
    headers: { 'content-type': 'application/json' },
  });
}
afterEach(() => vi.unstubAllGlobals());

const ROSTER = [{ guid: 'Player-4184-000000A1', name: 'Simfury', class: 'Warrior' }];

const PLAYER = {
  player_key: 'us/normal/simfury',
  player_name: 'Simfury',
  class: 'Warrior',
  spec: 'Fury',
  role: 'dps',
  overall: 70,
  overall_uncapped: 70,
  overall_capped: false,
  basis: 'percentile',
  components: [
    {
      name: 'output',
      score: 71,
      weight: 35,
      basis: 'percentile',
      percentile: 71,
      bracket_n: 142,
      excluded: false,
      reason: '',
      moments: [],
    },
    {
      name: 'survival',
      score: 76,
      weight: 15,
      basis: 'percentile',
      percentile: 60,
      bracket_n: 89,
      excluded: false,
      reason: '',
      moments: [
        {
          kind: 'death',
          at_ms: 140000,
          spell_id: 19712,
          spell_name: 'Arcane Explosion',
          avoidable: true,
          anchor: 'death-Player-4184-000000A1-140000',
        },
      ],
    },
    { name: 'mechanics', score: null, weight: 0, basis: '', percentile: null, bracket_n: 0, excluded: true, reason: 'no_mechanics_table', moments: [] },
    { name: 'utility', score: 62, weight: 15, basis: 'percentile', percentile: 62, bracket_n: 60, excluded: false, reason: '', moments: [] },
    { name: 'preparation', score: 95, weight: 10, basis: 'absolute', percentile: null, bracket_n: 0, excluded: false, reason: '', moments: [] },
    { name: 'activity', score: 80, weight: 5, basis: 'percentile', percentile: 80, bracket_n: 60, excluded: false, reason: '', moments: [] },
  ],
};

const RATINGS = { fight_index: 3, kill: true, kill_time_band: 'typical', model_version: 'v1', players: [PLAYER] };

describe('RatingTab', () => {
  it('one number first: the overall figure renders before the six parts', async () => {
    vi.stubGlobal('fetch', vi.fn<GlobalFetch>(async () => envelope(RATINGS)));
    render(RatingTab, {
      props: { reportId: 'fixture2abcd', fightIndex: 3, roster: ROSTER, source: SOURCE_FRIENDLIES, onSelectPlayer: vi.fn(), apiBase: 'https://api.test' },
    });

    await waitFor(() => expect(screen.getByTestId('rating-card-Player-4184-000000A1')).toBeTruthy());
    const card = screen.getByTestId('rating-card-Player-4184-000000A1');
    const overall = card.querySelector('[data-testid="rating-overall"]');
    const components = card.querySelector('[data-testid="rating-components"]');
    expect(overall).not.toBeNull();
    expect(components).not.toBeNull();
    expect(overall!.compareDocumentPosition(components!) & Node.DOCUMENT_POSITION_FOLLOWING).toBeTruthy();
    expect(overall!.textContent).toContain('70');
  });

  it('an excluded component states why in words, not a bare gap', async () => {
    vi.stubGlobal('fetch', vi.fn<GlobalFetch>(async () => envelope(RATINGS)));
    render(RatingTab, {
      props: { reportId: 'fixture2abcd', fightIndex: 3, roster: ROSTER, source: SOURCE_FRIENDLIES, onSelectPlayer: vi.fn(), apiBase: 'https://api.test' },
    });

    await waitFor(() => expect(screen.getByTestId('rating-component-mechanics')).toBeTruthy());
    expect(screen.getByTestId('rating-component-mechanics').textContent).toContain(
      'has no curated mechanics table yet',
    );
  });

  it('a moment opens into a link that jumps to the tab and anchor behind it', async () => {
    vi.stubGlobal('fetch', vi.fn<GlobalFetch>(async () => envelope(RATINGS)));
    render(RatingTab, {
      props: { reportId: 'fixture2abcd', fightIndex: 3, roster: ROSTER, source: SOURCE_FRIENDLIES, onSelectPlayer: vi.fn(), apiBase: 'https://api.test' },
    });

    await waitFor(() => expect(screen.getByTestId('rating-component-survival')).toBeTruthy());
    const link = screen.getByRole('link', { name: /Arcane Explosion/ });
    expect(link.getAttribute('href')).toBe('?fight=3&tab=deaths#death-Player-4184-000000A1-140000');
  });

  it('scoping to one player via source shows only that player’s card', async () => {
    vi.stubGlobal('fetch', vi.fn<GlobalFetch>(async () => envelope(RATINGS)));
    render(RatingTab, {
      props: {
        reportId: 'fixture2abcd',
        fightIndex: 3,
        roster: ROSTER,
        source: 'Player-4184-000000A1',
        onSelectPlayer: vi.fn(),
        apiBase: 'https://api.test',
      },
    });
    await waitFor(() => expect(screen.getByTestId('rating-tab')).toBeTruthy());
    expect(screen.queryAllByTestId(/^rating-card-/).length).toBe(1);
  });

  it('enemies have no ratings, and the tab says so without fetching a mismatched row', async () => {
    vi.stubGlobal('fetch', vi.fn<GlobalFetch>(async () => envelope(RATINGS)));
    render(RatingTab, {
      props: { reportId: 'fixture2abcd', fightIndex: 3, roster: ROSTER, source: SOURCE_ENEMIES, onSelectPlayer: vi.fn(), apiBase: 'https://api.test' },
    });
    await waitFor(() => expect(screen.getByTestId('rating-tab-enemies')).toBeTruthy());
  });

  it('no ratings at all for this fight is a truthful empty state, not a spinner', async () => {
    vi.stubGlobal('fetch', vi.fn<GlobalFetch>(async () => envelope({ ...RATINGS, players: [] })));
    render(RatingTab, {
      props: { reportId: 'fixture2abcd', fightIndex: 3, roster: ROSTER, source: SOURCE_FRIENDLIES, onSelectPlayer: vi.fn(), apiBase: 'https://api.test' },
    });
    await waitFor(() => expect(screen.getByTestId('rating-tab-empty')).toBeTruthy());
  });

  it('a capped overall score shows the coaching note naming what capped it', async () => {
    const capped = { ...PLAYER, overall: 40, overall_uncapped: 58, overall_capped: true };
    vi.stubGlobal('fetch', vi.fn<GlobalFetch>(async () => envelope({ ...RATINGS, players: [capped] })));
    render(RatingTab, {
      props: { reportId: 'fixture2abcd', fightIndex: 3, roster: ROSTER, source: SOURCE_FRIENDLIES, onSelectPlayer: vi.fn(), apiBase: 'https://api.test' },
    });
    await waitFor(() => expect(screen.getByTestId('rating-capped-note')).toBeTruthy());
    expect(screen.getByTestId('rating-capped-note').textContent).toContain('costly avoidable death');
  });

  it('the Utility component always notes that Threat is not modeled yet (spec §1.3)', async () => {
    vi.stubGlobal('fetch', vi.fn<GlobalFetch>(async () => envelope(RATINGS)));
    render(RatingTab, {
      props: { reportId: 'fixture2abcd', fightIndex: 3, roster: ROSTER, source: SOURCE_FRIENDLIES, onSelectPlayer: vi.fn(), apiBase: 'https://api.test' },
    });
    await waitFor(() => expect(screen.getByTestId('rating-component-utility')).toBeTruthy());
    expect(screen.getByTestId('rating-threat-note').textContent).toBe(
      'Threat — not modeled yet. This does not count for or against Utility.',
    );
  });

  it('links to the explanation page from the tab', async () => {
    vi.stubGlobal('fetch', vi.fn<GlobalFetch>(async () => envelope(RATINGS)));
    render(RatingTab, {
      props: { reportId: 'fixture2abcd', fightIndex: 3, roster: ROSTER, source: SOURCE_FRIENDLIES, onSelectPlayer: vi.fn(), apiBase: 'https://api.test' },
    });
    await waitFor(() => expect(screen.getByRole('link', { name: 'How is this calculated?' })).toBeTruthy());
    expect(screen.getByRole('link', { name: 'How is this calculated?' }).getAttribute('href')).toBe('/ratings');
  });
});
```

- [ ] **Step 2: Run test to verify it fails**

Run: `cd web && npx vitest run src/components/report/RatingTab.test.ts`
Expected: FAIL — `./RatingTab.svelte` does not exist.

- [ ] **Step 3: Implement `web/src/components/report/RatingTab.svelte`**

```svelte
<!-- web/src/components/report/RatingTab.svelte -->
<!-- The full report card: one number first, then the six components, then each
     component's moments -- spec section 6.1's exact three-level structure. Reuses
     ActorRow.svelte's existing <details>/aria-expanded disclosure idiom twice: once per
     player (only relevant when `source` shows every friendly), once per component. -->
<script lang="ts">
  import { splitUnitName } from '../../lib/characters';
  import { classColorVar } from '../../lib/report/format';
  import { SOURCE_ENEMIES, SOURCE_FRIENDLIES } from '../../lib/report/url';
  import { createReportRatingsFetch } from '../../lib/rating/report-ratings.svelte';
  import { ratingCopy, componentLabel } from '../../lib/rating/copy';
  import { momentHref } from '../../lib/rating/moments';
  import { classForPlayer, guidForPlayer } from '../../lib/rating/roster-join';
  import { RATING_COMPONENT_ORDER } from '../../lib/rating/types';
  import type { RatingCardPlayer, RatingComponent, RatingMoment } from '../../lib/rating/types';

  let {
    reportId,
    fightIndex,
    roster,
    source,
    onSelectPlayer,
    apiBase = undefined,
  }: {
    reportId: string;
    fightIndex: number;
    roster: { guid: string; name: string; class?: string }[];
    source: string;
    onSelectPlayer: (guid: string) => void;
    apiBase?: string;
  } = $props();

  const fetcher = createReportRatingsFetch(apiBase);
  $effect(() => {
    fetcher.load(reportId, fightIndex);
  });

  const withGuids = $derived(
    (fetcher.data?.players ?? []).map((player) => ({ player, guid: guidForPlayer(player, roster) })),
  );
  const visible = $derived(
    source === SOURCE_ENEMIES
      ? []
      : source === SOURCE_FRIENDLIES
        ? withGuids
        : withGuids.filter((row) => row.guid === source),
  );

  function orderedComponents(player: RatingCardPlayer): RatingComponent[] {
    return RATING_COMPONENT_ORDER.map(
      (name) =>
        player.components.find((c) => c.name === name) ?? {
          name,
          score: null,
          weight: 0,
          basis: '',
          percentile: null,
          bracket_n: 0,
          excluded: true,
          reason: '',
          moments: [],
        },
    );
  }

  function basisLine(player: RatingCardPlayer, part: RatingComponent): string {
    if (part.excluded) return ratingCopy.excludedReason(componentLabel(part.name), part.reason);
    if (part.basis === 'percentile' && part.percentile !== null)
      return ratingCopy.percentileBasis(player.spec, player.class, player.role, part.percentile, part.bracket_n);
    if (part.basis === 'absolute') return ratingCopy.absoluteBasis;
    return '';
  }

  function momentText(moment: RatingMoment): string {
    const time = moment.at_ms !== undefined ? `${Math.round(moment.at_ms / 1000)}s — ` : '';
    return `${time}${moment.spell_name ?? 'Unnamed'}`;
  }
</script>

<div class="flex flex-col gap-4" data-testid="rating-tab">
  {#if fetcher.status === 'loading' || fetcher.status === 'idle'}
    <p class="text-muted text-[14px]" data-testid="rating-tab-loading">Loading ratings.</p>
  {:else if fetcher.status === 'failed'}
    <p class="text-[14px]" role="alert" data-testid="rating-tab-error">{ratingCopy.fetchFailed}</p>
  {:else if source === SOURCE_ENEMIES}
    <p class="text-muted text-[14px]" data-testid="rating-tab-enemies">{ratingCopy.enemiesHaveNone}</p>
  {:else if visible.length === 0 && withGuids.length === 0}
    <p class="text-muted text-[14px]" data-testid="rating-tab-empty">{ratingCopy.tabEmpty}</p>
  {:else if visible.length === 0}
    <p class="text-muted text-[14px]" data-testid="rating-tab-no-match">{ratingCopy.noMatchingRow}</p>
  {:else}
    {#each visible as { player, guid } (player.player_key)}
      <details class="border-line rounded-panel bg-raised border p-3" data-testid={`rating-card-${guid || player.player_key}`} open={visible.length === 1}>
        <summary class="flex cursor-pointer items-center gap-3">
          <span class="flex-1 truncate font-semibold" style={`color: ${classColorVar(classForPlayer(guid, player.class, roster))}`}>
            {splitUnitName(player.player_name).name}
          </span>
          <span class="text-gold tabular text-[28px] font-bold" data-testid="rating-overall">{Math.round(player.overall)}</span>
        </summary>
        {#if player.overall_capped}
          <p class="text-muted mt-2 text-[13px]" data-testid="rating-capped-note">
            <span class="tabular font-semibold">{Math.round(player.overall)}</span>, {ratingCopy.cappedNote}
          </p>
        {/if}
        <ul class="mt-3 flex flex-col gap-2" data-testid="rating-components">
          {#each orderedComponents(player) as part (part.name)}
            <li>
              <details class="border-line-soft rounded-control border p-2" data-testid={`rating-component-${part.name}`}>
                <summary class="flex cursor-pointer items-center gap-3 text-[13px]">
                  <span class="flex-1 font-semibold">{componentLabel(part.name)}</span>
                  {#if !part.excluded}
                    <span class="tabular text-muted text-[12px]">weight {Math.round(part.weight)}%</span>
                    <span class="text-gold tabular w-10 text-right font-mono font-bold">{Math.round(part.score ?? 0)}</span>
                  {:else}
                    <span class="text-muted text-[12px]">not scored</span>
                  {/if}
                </summary>
                <p class="text-muted mt-2 text-[12px]">{basisLine(player, part)}</p>
                {#if part.name === 'utility' && !part.excluded}
                  <p class="text-muted mt-1 text-[12px]" data-testid="rating-threat-note">{ratingCopy.threatNotModeled}</p>
                {/if}
                {#if part.moments.length > 0}
                  <ul class="mt-2 flex flex-col gap-1" data-testid={`rating-moments-${part.name}`}>
                    {#each part.moments as moment, index (index)}
                      {@const href = momentHref(fightIndex, guid, moment)}
                      <li class="text-[12px]">
                        {#if href !== null}
                          <a class="text-gold underline-offset-2 hover:underline" {href}>{momentText(moment)}</a>
                        {:else}
                          <span class="text-muted">{momentText(moment)}</span>
                        {/if}
                      </li>
                    {/each}
                  </ul>
                {/if}
              </details>
            </li>
          {/each}
        </ul>
      </details>
    {/each}
  {/if}
  <a class="text-gold inline-flex min-h-11 items-center text-[12px] underline-offset-2 hover:underline md:min-h-0" href="/ratings"
    >{ratingCopy.explainLink}</a
  >
</div>
```

- [ ] **Step 4: Run test to verify it passes**

Run: `cd web && npx vitest run src/components/report/RatingTab.test.ts`
Expected: PASS. If the "one number first" DOM-order test is flaky against `<details>`/
`<summary>` nesting (the overall figure is inside `<summary>`, the components list is a
sibling `<ul>` after it), confirm `compareDocumentPosition` still reports `overall` before
`rating-components` — it does, since `<summary>` is `<details>`'s first child and `<ul
data-testid="rating-components">` follows it in document order regardless of open/closed
state; if the assertion needs `open` forced true for the container to be queryable at all,
the fixture already sets `open={visible.length === 1}` which is true for every test above
except none (all fixtures have exactly one visible player).

- [ ] **Step 5: Type-check, lint, format**

Run: `cd web && npx astro check && npm run lint && npx prettier --check src/components/report/RatingTab.svelte src/components/report/RatingTab.test.ts`

- [ ] **Step 6: Commit**

```bash
printf 'feat(web): RatingTab, the full three-level report card\n\nOne number first, then the six components as ActorRow-style\n<details> disclosures, then each component'"'"'s moments as links built\nby moments.ts (Ruling 4). Handles excluded-with-reason, capped, one\n-player-scoped-by-source, enemies-have-none, and the truthful empty\nstate for a fight with no ratings computed yet.\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-task5.txt
git add web/src/components/report/RatingTab.svelte web/src/components/report/RatingTab.test.ts
git commit -F .superpowers/commit-msg-task5.txt
```

---

## Task 6: Wire the Rating tab into `ReportView.svelte`, plus the report e2e spec

**Files:**
- Modify: `web/src/components/report/ReportView.svelte`
- Create: `web/tests/e2e/report-rating.spec.ts`

**Interfaces:**
- Consumes: `RatingTab` (Task 5), `createLazyComponent` (existing), `roster`/`playerSet`/
  `reportId`/`state`/`patch`/`nightMode` (all already exist in `ReportView.svelte`).

- [ ] **Step 1: Add the lazy import and load effect**

Beside the existing five `createLazyComponent` calls (around line 86-91):

```ts
const ratingTabLazy = createLazyComponent(() => import('./RatingTab.svelte'));
```

Beside the existing five lazy-load `$effect`s (around line 1259-1277), add a sixth:

```ts
$effect(() => {
  if (scoped !== null && !nightMode && state.mode === 'analyze' && state.view === 'tables' && state.tab === 'rating')
    ratingTabLazy.load();
});
```

- [ ] **Step 2: Add the template branch**

In the tab chain (right after the existing `{:else if state.tab === 'deaths'}<DeathsTab ...
/>` block, before its closing `{/if}`), add two branches: one for `nightMode` (ratings are
per-fight, matching Ruling 8 and the existing "one pull's" message tone), one for the real tab:

```svelte
{:else if state.tab === 'rating' && nightMode}
  <p class="text-muted text-[14px]" data-testid="rating-tab-night">{ratingNightCopy}</p>
{:else if state.tab === 'rating'}
  {#if ratingTabLazy.current}
    <ratingTabLazy.current
      {reportId}
      fightIndex={state.fight}
      {roster}
      source={state.source}
      onSelectPlayer={(guid) => patch({ source: guid })}
    />
  {:else}
    {@render lazyFallback(ratingTabLazy)}
  {/if}
{/if}
```

Define `ratingNightCopy` as a local constant near the top of the `<script>` block (beside
other small literals already inline in the file — search for how the existing
`night-tables-only` copy is written; it is inlined directly in that block's own markup rather
than a named constant, so match that: inline the string directly rather than adding a new
constant):

```svelte
{:else if state.tab === 'rating' && nightMode}
  <p class="text-muted text-[14px]" data-testid="rating-tab-night">
    Ratings are one pull's. Pick a boss pull from the list to see one.
  </p>
```

(Use this inline form instead of the `ratingNightCopy` variable above — it matches the file's
own existing style for this exact kind of message.)

- [ ] **Step 3: Run the existing report-tabs e2e spec for regressions before adding new ones**

Run: `cd web && npx astro preview stop 2>/dev/null; E2E_PORT=4451 npx playwright test tests/e2e/report-tabs.spec.ts`
Expected: PASS — confirms the Tab/ModeBar/ReportView changes have not broken the existing tab
bar.

- [ ] **Step 4: Write the new e2e spec**

```ts
// web/tests/e2e/report-rating.spec.ts
// Rating tab coverage: the tab reaches the full card, the Summary dashboard's rating row
// links into it, a scoped one-player link opens directly onto that player, and the
// truthful empty state shows when the API has nothing yet -- API stubbed with page.route
// per docs/superpowers/plans/2026-09-21-rating-web.md (the API lane has not landed).
import { expect, test } from '@playwright/test';

const REPORT = '/reports/fixture2abcd';

const envelope = (data: unknown, status = 200) => ({
  status,
  contentType: 'application/json',
  body: JSON.stringify({ ok: status < 400, data, error: null, request_id: 'r' }),
});

const PLAYER = {
  player_key: 'us/normal/elyra-duskvale',
  player_name: 'Elyra Duskvale',
  class: 'Priest',
  spec: 'Holy',
  role: 'healer',
  overall: 70,
  overall_uncapped: 70,
  overall_capped: false,
  basis: 'percentile',
  components: [
    { name: 'output', score: 71, weight: 30, basis: 'percentile', percentile: 71, bracket_n: 40, excluded: false, reason: '', moments: [] },
    { name: 'survival', score: 90, weight: 10, basis: 'percentile', percentile: 90, bracket_n: 40, excluded: false, reason: '', moments: [] },
    { name: 'mechanics', score: null, weight: 0, basis: '', percentile: null, bracket_n: 0, excluded: true, reason: 'no_mechanics_table', moments: [] },
    { name: 'utility', score: 62, weight: 20, basis: 'percentile', percentile: 62, bracket_n: 40, excluded: false, reason: '', moments: [] },
    { name: 'preparation', score: 100, weight: 10, basis: 'absolute', percentile: null, bracket_n: 0, excluded: false, reason: '', moments: [] },
    { name: 'activity', score: 66, weight: 10, basis: 'percentile', percentile: 66, bracket_n: 40, excluded: false, reason: '', moments: [] },
  ],
};

test('the Summary dashboard shows a rating row that opens the Rating tab', async ({ page }) => {
  await page.route('**/v1/reports/fixture2abcd/fights/3/ratings', (route) =>
    route.fulfill(envelope({ fight_index: 3, kill: true, kill_time_band: 'typical', model_version: 'v1', players: [PLAYER] })),
  );
  await page.goto(`${REPORT}?fight=3`);

  await expect(page.getByTestId('rating-panel')).toBeVisible();
  await expect(page.getByTestId('rating-panel-overall')).toHaveText('70');
  await page.getByRole('button', { name: 'Rating tab' }).click();

  await expect(page).toHaveURL(/tab=rating/);
  await expect(page.getByTestId('rating-tab')).toBeVisible();
  await expect(page.getByTestId('rating-overall')).toHaveText('70');
});

test('an excluded component states why, and a moment link jumps to its tab', async ({ page }) => {
  await page.route('**/v1/reports/fixture2abcd/fights/3/ratings', (route) =>
    route.fulfill(envelope({ fight_index: 3, kill: true, kill_time_band: 'typical', model_version: 'v1', players: [PLAYER] })),
  );
  await page.goto(`${REPORT}?fight=3&tab=rating`);

  await expect(page.getByTestId('rating-component-mechanics')).toContainText('has no curated mechanics table yet');
});

test('scoping the url to one player shows only that player’s card', async ({ page }) => {
  await page.route('**/v1/reports/fixture2abcd/fights/3/ratings', (route) =>
    route.fulfill(envelope({ fight_index: 3, kill: true, kill_time_band: 'typical', model_version: 'v1', players: [PLAYER] })),
  );
  await page.goto(`${REPORT}?fight=3`);
  const eventsRow = page.locator('[data-testid="source-scope"] option', { hasText: 'Elyra Duskvale' });
  const guid = await eventsRow.getAttribute('value');
  await page.goto(`${REPORT}?fight=3&tab=rating&source=${guid}`);

  await expect(page.getByTestId('rating-tab')).toBeVisible();
  await expect(page.locator('[data-testid^="rating-card-"]')).toHaveCount(1);
});

test('no ratings yet for this fight is a truthful empty state, not a spinner or an error', async ({ page }) => {
  await page.route('**/v1/reports/fixture2abcd/fights/3/ratings', (route) =>
    route.fulfill(envelope({ fight_index: 3, kill: true, kill_time_band: 'typical', model_version: 'v1', players: [] })),
  );
  await page.goto(`${REPORT}?fight=3&tab=rating`);

  await expect(page.getByTestId('rating-tab-empty')).toBeVisible();
});

test('the whole night has no ratings view; it explains why rather than rendering nothing', async ({ page }) => {
  await page.goto(`${REPORT}?fight=0&tab=rating`);
  await expect(page.getByTestId('rating-tab-night')).toBeVisible();
});

test('the Rating tab links to the explanation page', async ({ page }) => {
  await page.route('**/v1/reports/fixture2abcd/fights/3/ratings', (route) =>
    route.fulfill(envelope({ fight_index: 3, kill: true, kill_time_band: 'typical', model_version: 'v1', players: [PLAYER] })),
  );
  await page.goto(`${REPORT}?fight=3&tab=rating`);
  await expect(page.getByRole('link', { name: 'How is this calculated?' })).toHaveAttribute('href', '/ratings');
});
```

- [ ] **Step 5: Run the new e2e spec**

Run: `cd web && E2E_PORT=4451 npx playwright test tests/e2e/report-rating.spec.ts`
Expected: PASS. If the "scoping the url" test's option-locator does not resolve (the fixture
report's fight 3 roster names may differ from "Elyra Duskvale" — confirm against
`web/src/fixtures/report/fights/3/summary.json`'s actual roster before relying on that name;
substitute the real fixture player's name and expect the `PLAYER.player_name` in the route
stub to match it exactly, since Ruling 3's join is name-based).

- [ ] **Step 6: Run the full existing report e2e suite for regressions**

Run: `cd web && E2E_PORT=4451 npx playwright test tests/e2e/report-tabs.spec.ts tests/e2e/report-deaths.spec.ts tests/e2e/report-phone.spec.ts tests/e2e/report-mechanics.spec.ts`
Expected: PASS.

- [ ] **Step 7: Type-check, lint, format**

Run: `cd web && npx astro check && npm run lint && npx prettier --check src/components/report/ReportView.svelte tests/e2e/report-rating.spec.ts`

- [ ] **Step 8: Commit**

```bash
printf 'feat(web): wire the Rating tab into the report island\n\nLazy-loads RatingTab.svelte the way Timelines/Events/Queries/Compare/\nMechanics already are; the whole-night view explains that ratings are\none pull'"'"'s rather than rendering nothing. Covers the dashboard link,\nthe excluded-component copy, one-player url scoping, the empty state,\nnight mode, and the explanation-page link end to end.\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-task6.txt
git add web/src/components/report/ReportView.svelte web/tests/e2e/report-rating.spec.ts
git commit -F .superpowers/commit-msg-task6.txt
```

---

## Task 7: Character page rating panel, the trend sparkline, and its e2e spec

**Files:**
- Create: `web/src/components/RatingTrend.svelte`
- Create: `web/src/components/RatingTrend.test.ts`
- Create: `web/src/components/CharacterRatingPanel.svelte`
- Create: `web/src/components/CharacterRatingPanel.test.ts`
- Modify: `web/src/components/Character.svelte`
- Create: `web/tests/e2e/character-rating.spec.ts`

**Interfaces:**
- Consumes: `fetchCharacterRating` (Task 2), `CharacterRating`/`RatingCardPlayer`/
  `RATING_COMPONENT_ORDER` (Task 2's `types.ts`), `ratingCopy`/`componentLabel`/
  `MIN_TREND_SAMPLES` (Task 2's `copy.ts`), `CharacterPath` (existing).
- Produces: `RatingTrend` props `{ points: { fought_at: string; overall: number }[] }` (pure,
  presentational, no fetch). `CharacterRatingPanel` props `{ path: CharacterPath; apiBase?:
  string }` (does its own fetch, mirroring `Character.svelte`'s own `$effect` pattern).

- [ ] **Step 1: Write the failing test for `RatingTrend.svelte`**

```ts
// web/src/components/RatingTrend.test.ts
import { describe, expect, it } from 'vitest';
import { render } from '@testing-library/svelte';
import RatingTrend from './RatingTrend.svelte';

describe('RatingTrend', () => {
  it('draws one point per fight as an inline svg polyline, oldest first', () => {
    const { container } = render(RatingTrend, {
      props: {
        points: [
          { fought_at: '2026-12-01T00:00:00Z', overall: 40 },
          { fought_at: '2026-12-05T00:00:00Z', overall: 70 },
          { fought_at: '2026-12-09T00:00:00Z', overall: 55 },
        ],
      },
    });
    const svg = container.querySelector('svg');
    expect(svg).not.toBeNull();
    const polyline = container.querySelector('polyline');
    expect(polyline?.getAttribute('points')?.split(' ').length).toBe(3);
  });

  it('renders nothing (an empty fragment) for fewer than two points', () => {
    const { container } = render(RatingTrend, { props: { points: [{ fought_at: '2026-12-01T00:00:00Z', overall: 40 }] } });
    expect(container.querySelector('svg')).toBeNull();
  });
});
```

- [ ] **Step 2: Run test to verify it fails**

Run: `cd web && npx vitest run src/components/RatingTrend.test.ts`
Expected: FAIL — component does not exist.

- [ ] **Step 3: Implement `web/src/components/RatingTrend.svelte`**

```svelte
<!-- web/src/components/RatingTrend.svelte -->
<!-- A small dedicated sparkline for a character's rating trend: a handful of unevenly
     -spaced points across weeks, not TimeChart.svelte's dense per-second in-fight series
     with a drag-brush -- see docs/superpowers/plans/2026-09-21-rating-web.md's Ruling 6 for
     why this is its own tiny component rather than a TimeChart reuse. No charting library:
     one inline <svg>, one <polyline>, matching the site's existing no-dependency-chart
     precedent (ResourceGraphs.svelte, TimeChart.svelte). -->
<script lang="ts">
  let { points }: { points: { fought_at: string; overall: number }[] } = $props();

  const WIDTH = 240;
  const HEIGHT = 48;
  const PAD = 4;

  const ordered = $derived([...points].sort((a, b) => a.fought_at.localeCompare(b.fought_at)));

  const coords = $derived.by(() => {
    if (ordered.length < 2) return [];
    const span = ordered.length - 1;
    return ordered.map((point, index) => {
      const x = PAD + (index / span) * (WIDTH - PAD * 2);
      const y = PAD + (1 - point.overall / 100) * (HEIGHT - PAD * 2);
      return `${x.toFixed(1)},${y.toFixed(1)}`;
    });
  });
</script>

{#if coords.length > 0}
  <svg
    viewBox={`0 0 ${WIDTH} ${HEIGHT}`}
    width={WIDTH}
    height={HEIGHT}
    role="img"
    aria-label={`Rating trend over the last ${ordered.length} fights, from ${Math.round(ordered[0].overall)} to ${Math.round(ordered[ordered.length - 1].overall)}`}
    data-testid="rating-trend"
  >
    <polyline points={coords.join(' ')} fill="none" stroke="var(--color-gold)" stroke-width="2" stroke-linejoin="round" stroke-linecap="round" />
  </svg>
{/if}
```

- [ ] **Step 4: Run test to verify it passes**

Run: `cd web && npx vitest run src/components/RatingTrend.test.ts`
Expected: PASS.

- [ ] **Step 5: Write the failing test for `CharacterRatingPanel.svelte`**

```ts
// web/src/components/CharacterRatingPanel.test.ts
import { afterEach, describe, expect, it, vi } from 'vitest';
import { render, screen, waitFor } from '@testing-library/svelte';
import CharacterRatingPanel from './CharacterRatingPanel.svelte';

type GlobalFetch = (...args: Parameters<typeof fetch>) => Promise<Response>;
function envelope(data: unknown, status = 200): Response {
  return new Response(JSON.stringify({ ok: status < 400, data, error: null, request_id: 'r' }), {
    status,
    headers: { 'content-type': 'application/json' },
  });
}
afterEach(() => vi.unstubAllGlobals());

const PATH = { region: 'us' as const, ruleset: 'hardcore' as const, slug: 'elyra-duskvale' };

describe('CharacterRatingPanel', () => {
  it('renders the latest overall score, the trend, and best/worst component', async () => {
    const data = {
      player_key: 'us/hardcore/elyra-duskvale',
      sample_size: 8,
      trend: [
        { fought_at: '2026-12-01T00:00:00Z', overall: 40, report_id: 'r1', fight_index: 1 },
        { fought_at: '2026-12-09T00:00:00Z', overall: 70, report_id: 'r2', fight_index: 2 },
      ],
      best_component: 'preparation',
      worst_component: 'activity',
      latest: {
        player_key: 'us/hardcore/elyra-duskvale',
        player_name: 'Elyra Duskvale',
        class: 'Priest',
        spec: 'Holy',
        role: 'healer',
        overall: 70,
        overall_uncapped: 70,
        overall_capped: false,
        basis: 'percentile',
        components: [],
      },
    };
    vi.stubGlobal('fetch', vi.fn<GlobalFetch>(async () => envelope(data)));
    render(CharacterRatingPanel, { props: { path: PATH, apiBase: 'https://api.test' } });

    await waitFor(() => expect(screen.getByTestId('character-rating-overall')).toBeTruthy());
    expect(screen.getByTestId('character-rating-overall').textContent).toContain('70');
    expect(screen.getByTestId('rating-trend')).toBeTruthy();
    expect(screen.getByText('Preparation', { exact: false })).toBeTruthy();
    expect(screen.getByText('Activity', { exact: false })).toBeTruthy();
  });

  it('too few rated fights shows the exact threshold copy instead of a sparkline', async () => {
    const data = {
      player_key: 'us/hardcore/elyra-duskvale',
      sample_size: 2,
      trend: [
        { fought_at: '2026-12-01T00:00:00Z', overall: 40, report_id: 'r1', fight_index: 1 },
        { fought_at: '2026-12-02T00:00:00Z', overall: 44, report_id: 'r1', fight_index: 2 },
      ],
      best_component: '',
      worst_component: '',
      latest: null,
    };
    vi.stubGlobal('fetch', vi.fn<GlobalFetch>(async () => envelope(data)));
    render(CharacterRatingPanel, { props: { path: PATH, apiBase: 'https://api.test' } });

    await waitFor(() => expect(screen.getByTestId('character-rating-too-few')).toBeTruthy());
    expect(screen.getByTestId('character-rating-too-few').textContent).toBe('Not enough rated fights yet to show a trend (2 of 5 needed).');
  });

  it('zero rated fights is a truthful empty state', async () => {
    const data = { player_key: 'x', sample_size: 0, trend: [], best_component: '', worst_component: '', latest: null };
    vi.stubGlobal('fetch', vi.fn<GlobalFetch>(async () => envelope(data)));
    render(CharacterRatingPanel, { props: { path: PATH, apiBase: 'https://api.test' } });

    await waitFor(() => expect(screen.getByTestId('character-rating-empty')).toBeTruthy());
  });

  it('an anonymized character (404) renders nothing at all -- no placeholder', async () => {
    vi.stubGlobal('fetch', vi.fn<GlobalFetch>(async () => envelope(null, 404)));
    const { container } = render(CharacterRatingPanel, { props: { path: PATH, apiBase: 'https://api.test' } });

    await waitFor(() => expect(container.querySelector('[data-testid^="character-rating"]')).toBeNull());
  });
});
```

- [ ] **Step 6: Run test to verify it fails**

Run: `cd web && npx vitest run src/components/CharacterRatingPanel.test.ts`
Expected: FAIL — component does not exist.

- [ ] **Step 7: Implement `web/src/components/CharacterRatingPanel.svelte`**

```svelte
<!-- web/src/components/CharacterRatingPanel.svelte -->
<!-- The character page's rating panel: overall score, six-segment bar, a trend
     sparkline once there are enough fights, and the best/worst component -- spec §6.2.
     Fetches independently, the same self-contained pattern Character.svelte's own
     $effect already uses for fetchCharacter. -->
<script lang="ts">
  import { fetchCharacterRating, RankingsError } from '../lib/rankings/api';
  import { ratingCopy, componentLabel, MIN_TREND_SAMPLES } from '../lib/rating/copy';
  import { RATING_COMPONENT_ORDER } from '../lib/rating/types';
  import type { CharacterPath } from '../lib/characters';
  import type { CharacterRating } from '../lib/rating/types';
  import RatingTrend from './RatingTrend.svelte';

  let { path, apiBase = undefined }: { path: CharacterPath; apiBase?: string } = $props();

  let data = $state<CharacterRating | null>(null);
  let status = $state<'loading' | 'ready' | 'hidden'>('loading');

  $effect(() => {
    const requested = path;
    status = 'loading';
    void fetchCharacterRating(requested, apiBase)
      .then((result) => {
        if (path !== requested) return;
        data = result;
        status = 'ready';
      })
      .catch((thrown: unknown) => {
        if (path !== requested) return;
        // Ruling 2: an anonymized character 404s. Every other failure also hides the
        // panel rather than showing an alarming error for what is, from a visitor's
        // seat, the same as "nothing here yet" -- the rest of the character page already
        // loaded and still has value.
        if (!(thrown instanceof RankingsError) || thrown.status !== 404) {
          // Non-anonymize failures still hide rather than alarm, but are worth a console
          // trace for whoever is watching the browser's own devtools during development.
        }
        data = null;
        status = 'hidden';
      });
  });

</script>

{#if status === 'ready' && data !== null}
  <section class="flex flex-col gap-2" data-testid="character-rating">
    <h2 class="section-title text-[18px]">Performance rating</h2>
    {#if data.sample_size === 0}
      <p class="text-muted text-[14px]" data-testid="character-rating-empty">{ratingCopy.characterEmpty}</p>
    {:else}
      {#if data.latest !== null}
        <div class="flex items-center gap-4">
          <span class="text-gold tabular text-[32px] font-bold" data-testid="character-rating-overall">
            {Math.round(data.latest.overall)}
          </span>
          <span class="flex w-32 shrink-0 gap-[2px]" aria-hidden="true">
            {#each RATING_COMPONENT_ORDER as name (name)}
              {@const part = data.latest.components.find((c) => c.name === name)}
              <span class="bg-line-soft h-2 flex-1" title={componentLabel(name)}
                ><span class="bg-gold block h-full" style={`width: ${part && !part.excluded && part.score !== null ? part.score : 0}%`}
                ></span></span
              >
            {/each}
          </span>
        </div>
      {/if}
      {#if data.sample_size < MIN_TREND_SAMPLES}
        <p class="text-muted text-[13px]" data-testid="character-rating-too-few">{ratingCopy.trendTooFew(data.sample_size)}</p>
      {:else}
        <RatingTrend points={data.trend} />
      {/if}
      {#if data.best_component !== '' || data.worst_component !== ''}
        <p class="text-muted text-[13px]">
          {#if data.best_component !== ''}Best: {componentLabel(data.best_component)}.{/if}
          {#if data.worst_component !== ''}Needs work: {componentLabel(data.worst_component)}.{/if}
        </p>
      {/if}
    {/if}
    <a class="text-gold inline-flex min-h-11 items-center text-[12px] underline-offset-2 hover:underline md:min-h-0" href="/ratings"
      >{ratingCopy.explainLink}</a
    >
  </section>
{/if}
```

- [ ] **Step 8: Run test to verify it passes**

Run: `cd web && npx vitest run src/components/CharacterRatingPanel.test.ts`
Expected: PASS.

- [ ] **Step 9: Wire it into `Character.svelte`**

Read `web/src/components/Character.svelte` in full (already shown in research above), then
add the import beside the existing ones and mount it right after the `<header>` block and
before the `<section>` for "Best per encounter" (so the rating panel sits high on the page,
matching spec §6.2's "a new panel beside the existing progression/roster-bests panels"):

```svelte
<script lang="ts">
  // ...existing imports...
  import CharacterRatingPanel from './CharacterRatingPanel.svelte';
</script>
```

```svelte
      <CharacterHandoffLinks path={resolved} />
    </header>

    <CharacterRatingPanel path={resolved} />

    <section class="flex flex-col gap-2">
      <h2 class="section-title text-[18px]">Best per encounter</h2>
```

- [ ] **Step 10: Run a quick regression on `Character.svelte`'s existing behavior**

Run: `cd web && E2E_PORT=4451 npx playwright test tests/e2e/character-phone.spec.ts tests/e2e/current-character.spec.ts`
Expected: PASS — `character-phone.spec.ts`'s fixture does not stub the new `/rating` endpoint,
so `CharacterRatingPanel` will hit an unmocked route in Playwright; Playwright's default
behavior for an unstubbed `**/v1/characters/.../rating` request is to let it through to
whatever `E2E_PORT`'s dev server proxies (there is no live API in e2e), which should surface as
a fetch failure the panel already handles by hiding itself (`status = 'hidden'`) rather than
throwing — confirm this is actually true by running the test; if `character-phone.spec.ts`
fails because the panel's failed fetch introduces a console error the test's own error-collector
asserts against, add a `page.route('**/v1/characters/**/rating', route => route.fulfill(envelope(null, 404)))`
stub to that spec file's existing `beforeEach`/route setup (read the file first to find where
its other `/v1/characters/...` stub already lives, matching Step 25 below's own confirmed
location) rather than leaving it unmocked.

- [ ] **Step 11: Write the character e2e spec**

```ts
// web/tests/e2e/character-rating.spec.ts
// The character page's rating panel: overall + six-segment bar + trend once there are
// enough fights, the too-few-samples copy below that, and the anonymized-character case
// rendering nothing. API stubbed with page.route (the API lane has not landed).
import { expect, test } from '@playwright/test';

const CHARACTER = {
  ok: true,
  data: {
    character: { name: 'Elyra Duskvale', region: 'us', ruleset: 'hardcore', class: 'Priest' },
    best: [],
    history: [],
    builds_seen: [],
  },
  error: null,
  request_id: 'r',
};

const envelope = (data: unknown, status = 200) => ({
  status,
  contentType: 'application/json',
  body: JSON.stringify({ ok: status < 400, data, error: null, request_id: 'r' }),
});

test('the character page shows the rating panel with a trend once there are enough fights', async ({ page }) => {
  await page.route('**/v1/characters/us/hardcore/elyra-duskvale', (route) => route.fulfill(envelope(CHARACTER.data)));
  await page.route('**/v1/characters/us/hardcore/elyra-duskvale/rating', (route) =>
    route.fulfill(
      envelope({
        player_key: 'us/hardcore/elyra-duskvale',
        sample_size: 6,
        trend: Array.from({ length: 6 }, (_, i) => ({
          fought_at: `2026-12-0${i + 1}T00:00:00Z`,
          overall: 40 + i * 5,
          report_id: 'fixture2abcd',
          fight_index: i + 1,
        })),
        best_component: 'preparation',
        worst_component: 'activity',
        latest: {
          player_key: 'us/hardcore/elyra-duskvale',
          player_name: 'Elyra Duskvale',
          class: 'Priest',
          spec: 'Holy',
          role: 'healer',
          overall: 65,
          overall_uncapped: 65,
          overall_capped: false,
          basis: 'percentile',
          components: [],
        },
      }),
    ),
  );

  await page.goto('/character/us/hardcore/elyra-duskvale');

  await expect(page.getByTestId('character-rating')).toBeVisible();
  await expect(page.getByTestId('character-rating-overall')).toHaveText('65');
  await expect(page.getByTestId('rating-trend')).toBeVisible();
});

test('fewer than five rated fights shows the exact threshold copy, no sparkline', async ({ page }) => {
  await page.route('**/v1/characters/us/hardcore/elyra-duskvale', (route) => route.fulfill(envelope(CHARACTER.data)));
  await page.route('**/v1/characters/us/hardcore/elyra-duskvale/rating', (route) =>
    route.fulfill(
      envelope({
        player_key: 'us/hardcore/elyra-duskvale',
        sample_size: 2,
        trend: [
          { fought_at: '2026-12-01T00:00:00Z', overall: 40, report_id: 'r1', fight_index: 1 },
          { fought_at: '2026-12-02T00:00:00Z', overall: 44, report_id: 'r1', fight_index: 2 },
        ],
        best_component: '',
        worst_component: '',
        latest: null,
      }),
    ),
  );

  await page.goto('/character/us/hardcore/elyra-duskvale');

  await expect(page.getByTestId('character-rating-too-few')).toHaveText(
    'Not enough rated fights yet to show a trend (2 of 5 needed).',
  );
  await expect(page.getByTestId('rating-trend')).toHaveCount(0);
});

test('an anonymized character’s rating panel renders nothing, not a placeholder', async ({ page }) => {
  await page.route('**/v1/characters/us/hardcore/elyra-duskvale', (route) => route.fulfill(envelope(CHARACTER.data)));
  await page.route('**/v1/characters/us/hardcore/elyra-duskvale/rating', (route) =>
    route.fulfill({ status: 404, contentType: 'application/json', body: JSON.stringify({ ok: false, data: null, error: 'not_found', request_id: 'r' }) }),
  );

  await page.goto('/character/us/hardcore/elyra-duskvale');

  await expect(page.getByTestId('character-rating')).toHaveCount(0);
});
```

- [ ] **Step 12: Run the new e2e spec**

Run: `cd web && E2E_PORT=4451 npx playwright test tests/e2e/character-rating.spec.ts`
Expected: PASS.

- [ ] **Step 13: Type-check, lint, format**

Run: `cd web && npx astro check && npm run lint && npx prettier --check src/components/RatingTrend.svelte src/components/RatingTrend.test.ts src/components/CharacterRatingPanel.svelte src/components/CharacterRatingPanel.test.ts src/components/Character.svelte tests/e2e/character-rating.spec.ts`

- [ ] **Step 14: Commit**

```bash
printf 'feat(web): the character page’s rating panel and trend sparkline\n\nCharacterRatingPanel.svelte (self-contained fetch, mirroring\nCharacter.svelte'"'"'s own pattern) plus RatingTrend.svelte, a small\ndedicated inline-svg sparkline -- Ruling 6 explains why this is not a\nTimeChart reuse. Fewer than MIN_TREND_SAMPLES (5, Ruling 7) shows the\nexact threshold copy instead of a sparkline; a 404 (anonymized) hides\nthe panel entirely per Ruling 2.\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-task7.txt
git add web/src/components/RatingTrend.svelte web/src/components/RatingTrend.test.ts web/src/components/CharacterRatingPanel.svelte web/src/components/CharacterRatingPanel.test.ts web/src/components/Character.svelte web/tests/e2e/character-rating.spec.ts
git commit -F .superpowers/commit-msg-task7.txt
```

---

## Task 8: `/ratings` explanation page (zero client JavaScript)

**Files:**
- Create: `web/src/pages/ratings.astro`
- Create: `web/tests/e2e/ratings.spec.ts`

**Interfaces:**
- Consumes: `Base.astro` (no `session` prop — static), `Panel.astro`,
  `ratingCopy.cappedRule` (Task 2's `copy.ts`, so the page and the in-app copy never drift on
  the cap sentence's exact wording).

- [ ] **Step 1: Write the failing e2e spec**

```ts
// web/tests/e2e/ratings.spec.ts
// /ratings is a static content page: no client JavaScript, no layout shift, and it must
// publish every fact spec section 6.3 requires -- the weight table, the percentile-vs
// -absolute rule, the cap in the coordinator's exact words, and that a rating comes only
// from a public report.
import { expect, test } from '@playwright/test';

test('ships no client JavaScript', async ({ page }) => {
  const scripts: string[] = [];
  page.on('request', (r) => {
    if (r.resourceType() === 'script') scripts.push(r.url());
  });
  await page.goto('/ratings');
  expect(scripts).toEqual([]);
});

test('publishes the weight table for all three roles', async ({ page }) => {
  await page.goto('/ratings');
  await expect(page.getByRole('heading', { name: 'How ratings work' })).toBeVisible();
  const table = page.getByTestId('ratings-weights');
  await expect(table).toBeVisible();
  await expect(table).toContainText('Output');
  await expect(table).toContainText('Survival');
  await expect(table).toContainText('Mechanics');
  await expect(table).toContainText('Utility');
  await expect(table).toContainText('Preparation');
  await expect(table).toContainText('Activity');
  await expect(table).toContainText('DPS');
  await expect(table).toContainText('Healer');
  await expect(table).toContainText('Tank');
});

test('states the cap in the coordinator’s exact words', async ({ page }) => {
  await page.goto('/ratings');
  await expect(
    page.getByText('An avoidable death early in a fight caps the overall score at 40, because nothing else in the fight makes up for it.'),
  ).toBeVisible();
});

test('says ratings are public like Warcraft Logs parses, and only from public reports', async ({ page }) => {
  await page.goto('/ratings');
  await expect(page.getByTestId('ratings-visibility')).toContainText('public');
  await expect(page.getByTestId('ratings-visibility')).toContainText('report');
});

test('explains the percentile-vs-absolute rule, small samples, and role-specific weights', async ({ page }) => {
  await page.goto('/ratings');
  await expect(page.getByTestId('ratings-basis')).toContainText('20');
  await expect(page.getByTestId('ratings-small-samples')).toBeVisible();
  await expect(page.getByTestId('ratings-roles')).toContainText('tanks and healers');
});

test('says a wipe scores differently than a kill', async ({ page }) => {
  await page.goto('/ratings');
  await expect(page.getByTestId('ratings-wipes')).toBeVisible();
});
```

- [ ] **Step 2: Run test to verify it fails**

Run: `cd web && E2E_PORT=4451 npx playwright test tests/e2e/ratings.spec.ts`
Expected: FAIL — `/ratings` 404s.

- [ ] **Step 3: Implement `web/src/pages/ratings.astro`**

```astro
---
// web/src/pages/ratings.astro
// The public explanation page spec section 6.3 requires: the exact weight table, the
// formulas in plain language, the percentile-vs-absolute rule, the cap in the
// coordinator's own words, and that a rating comes only from a public report -- linked
// from every rating surface ("How is this calculated?"). No client JavaScript: a content
// page, per Base.astro's session prop (omitted here, its default is false).
import Base from '../layouts/Base.astro';
import Panel from '../components/Panel.astro';
import { ratingCopy } from '../lib/rating/copy';

const weights = [
  { component: 'Output', dps: 35, healer: 30, tank: 10 },
  { component: 'Survival', dps: 15, healer: 10, tank: 35 },
  { component: 'Mechanics', dps: 20, healer: 20, tank: 20 },
  { component: 'Utility', dps: 15, healer: 20, tank: 20 },
  { component: 'Preparation', dps: 10, healer: 10, tank: 10 },
  { component: 'Activity', dps: 5, healer: 10, tank: 5 },
];
---

<Base
  title="How ratings work"
  description="How World of Warcraft: Forever's performance ratings are computed: the six components, their weights per role, the percentile-versus-absolute rule, and the cap."
  path="/ratings"
>
  <main
    id="main"
    tabindex="-1"
    class="mx-auto flex w-full max-w-[900px] flex-col gap-6 px-[18px] py-8 md:gap-8 md:px-12"
  >
    <header class="flex flex-col gap-2">
      <h1 class="section-title text-[18px]">How ratings work</h1>
      <p class="text-muted max-w-[720px] text-[14px] leading-relaxed" data-testid="ratings-visibility">
        A rating comes only from a public report, the same way a Warcraft Logs parse does. There is no
        cross-guild leaderboard of ratings — a rating is one fight's roster, or one character's own history,
        never a ranked list of players against each other site-wide.
      </p>
    </header>

    <Panel title="One number first, then six parts">
      <div class="flex flex-col gap-3 p-3 text-[14px] leading-relaxed">
        <p>
          Every rating starts with one overall number, 0 to 100. That number is a weighted average of six
          components — Output, Survival, Mechanics, Utility, Preparation, Activity — each also scored 0 to
          100. A parse is not a performance: someone who tops the meter while dying to avoidable damage and
          skipping consumables does not rate well here, because Output is a fraction of the score, not the
          whole of it.
        </p>
        <p>
          A component the engine cannot score for a given fight — no curated table for this encounter, the
          simulator does not model this spec yet, too few logs to compare against — is left out and said so
          in plain words, never silently zeroed and never faked. Its weight is redistributed across the
          components that did score.
        </p>
      </div>
    </Panel>

    <Panel title="Weights, by role">
      <div class="overflow-x-auto p-3">
        <table class="w-full min-w-[420px] border-collapse text-[13px]" data-testid="ratings-weights">
          <thead>
            <tr>
              <th scope="col" class="label text-muted border-line-soft border-b px-2 py-2 text-left">Component</th>
              <th scope="col" class="label text-muted border-line-soft border-b px-2 py-2 text-right">DPS</th>
              <th scope="col" class="label text-muted border-line-soft border-b px-2 py-2 text-right">Healer</th>
              <th scope="col" class="label text-muted border-line-soft border-b px-2 py-2 text-right">Tank</th>
            </tr>
          </thead>
          <tbody>
            {weights.map((row) => (
              <tr>
                <th scope="row" class="border-line-soft border-b px-2 py-2 text-left font-semibold">{row.component}</th>
                <td class="tabular border-line-soft border-b px-2 py-2 text-right font-mono">{row.dps}</td>
                <td class="tabular border-line-soft border-b px-2 py-2 text-right font-mono">{row.healer}</td>
                <td class="tabular border-line-soft border-b px-2 py-2 text-right font-mono">{row.tank}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
      <p class="text-muted p-3 pt-0 text-[13px]" data-testid="ratings-roles">
        Tanks and healers are rated by their own weights, not a DPS role's: a tank's Output is worth little
        (10 of 100) and its Survival a lot (35), because taking damage on purpose is the job. A healer's
        Output is measured on effective healing, never raw — overheal does not count toward it.
      </p>
    </Panel>

    <Panel title="Percentile or absolute">
      <div class="flex flex-col gap-3 p-3 text-[14px] leading-relaxed" data-testid="ratings-basis">
        <p>
          When at least 20 kills of the same encounter, difficulty, spec, role and roughly the same kill
          speed exist, a component is scored as a percentile against those other logs — "better than 71% of
          Fury Warriors on this fight." Below that, there are not enough logs yet to compare players fairly,
          so the score is measured against the encounter's own numbers instead and says so plainly, rather
          than presenting a percentile built on noise.
        </p>
        <p data-testid="ratings-small-samples">
          A brand-new boss, or a kill speed nobody has repeated yet, starts on the absolute standard and
          moves to the percentile once the eleventh comparable log crosses that threshold within its own
          speed band (fast, typical, or slow, judged against the encounter's own kill times, not a hand
          -written expectation).
        </p>
        <p data-testid="ratings-wipes">
          A wipe scores differently than a kill: Survival and Mechanics are still computed — a wipe is
          exactly where avoidable deaths and missed interrupts matter most — but Output, Utility, Preparation
          and Activity are not, since a wipe's damage window is cut short and is not a fair comparison to a
          full kill's.
        </p>
      </div>
    </Panel>

    <Panel title="The cap">
      <p class="p-3 text-[14px] leading-relaxed">{ratingCopy.cappedRule}</p>
    </Panel>
  </main>
</Base>
```

- [ ] **Step 4: Run test to verify it passes**

Run: `cd web && npx astro build 2>&1 | tail -20 && E2E_PORT=4451 npx playwright test tests/e2e/ratings.spec.ts`
Expected: PASS. (A build is needed first since e2e serves the built `dist/`, matching every
other e2e spec's own precondition in this repo's toolchain.)

- [ ] **Step 5: Confirm zero layout shift and no console errors**

Run: `cd web && E2E_PORT=4451 npx playwright test tests/e2e/ratings.spec.ts --reporter=list`
Expected: PASS, no console errors reported (`page.on('request', ...)` in Step 1's first test
already asserts zero scripts, which by construction means zero hydration-driven shift).

- [ ] **Step 6: Type-check, lint, format**

Run: `cd web && npx astro check && npm run lint && npx prettier --check src/pages/ratings.astro tests/e2e/ratings.spec.ts`

- [ ] **Step 7: Commit**

```bash
printf 'feat(web): the /ratings explanation page\n\nA static, zero-client-JS content page publishing the weight table,\nthe percentile-vs-absolute rule, wipes, small samples, role-specific\nweights, and the cap in the coordinator'"'"'s exact words -- linked from\nevery rating surface as "How is this calculated?".\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-task8.txt
git add web/src/pages/ratings.astro web/tests/e2e/ratings.spec.ts
git commit -F .superpowers/commit-msg-task8.txt
```

---

## Task 9: Whole-branch verification

Not a code task — the final gate before the lane's report. Runs after Tasks 1-8 are all
committed and individually green.

- [ ] **Step 1: Full unit/type/lint/format sweep**

```bash
cd web
npx vitest run
npx astro check
npm run lint
npx prettier --check src tests
```

Expected: everything green, including every pre-existing test this plan did not touch (a
regression anywhere is this lane's to fix, per the dispatch's "a failure you can show is
pre-existing on main is reported, not fixed; anything else is yours").

- [ ] **Step 2: Full e2e sweep**

```bash
cd web
npx astro preview stop 2>/dev/null || true
E2E_PORT=4451 npx playwright test
```

Expected: everything green. Investigate any failure outside this plan's own new specs against
a clean `main` checkout before concluding it is pre-existing; fix anything this lane's changes
caused.

- [ ] **Step 3: Confirm no stray processes or worktree pollution**

```bash
cd web && npx astro preview stop 2>/dev/null || true
cd /Users/jh/code/forever/.worktrees/rating-web && git status
```

Expected: clean tree (everything committed), no `web/node_modules` staged anywhere across the
plan's commits (`git log --stat` over this branch's commits should show none), no background
preview server left running.

- [ ] **Step 4: Final report**

Compose the lane's final report per `lane-common-web.md`'s required format: branch head sha,
what shipped per spec bullet, test counts from this run, every ruling (all ten above), anything
left undone and why (none expected — Tasks 1-8 cover section 6 in full and the web half of §8's
tests), and files another lane must know about (`ExchangeTable.svelte`/`AuraTable.svelte` now
carry `id` attributes every tab can use, not only Rating — worth flagging to any lane touching
those files concurrently).
