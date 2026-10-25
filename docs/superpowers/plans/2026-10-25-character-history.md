# Character history charts Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** A Premium owner's character page gains a `History` panel under the Forever Score card: the score over time on tier bands with dated points, a parse-by-boss chart of the `output` percentile over the member's kills, and the fights behind the points; a free owner sees today's sparkline plus one upsell sentence; a visitor sees the card alone.

**Architecture:** One owner-only API route (`GET /v1/characters/{region}/{ruleset}/{name}/history`) reads the Forever Score snapshots the rater already rolls onto `rating_scores` and the `output` component's percentile per kill, grouped by boss, plus the gaps with a closed reason set and the fight rows behind every point. The web collapses score points to one per UTC date, draws two inline-SVG charts inside the character island gated on `entitlements.history`, and reuses `CharacterFightTable`'s row for the lists. The public payload and sparkline are untouched.

**Tech Stack:** Go (pgx v5), Svelte 5 inline SVG (no chart library), vitest, Playwright.

**Spec:** Product: `docs/superpowers/specs/2026-10-24-premium-experience-design.md` §4.4, §5. Experience: `design/specs/2026-10-24-character-history.md` (ux-designer; §3 layout, §4 states, §6 copy and §10 contract bind; mock `design/mocks/gen_character_history.py`).

## Global Constraints

- No new table and no migration: everything comes from `rating_scores` + `reports` + the public `history` rows' query (product spec §4.4 ruling).
- The owner-only query is never called by the public character payload; the visitor's sparkline stays `rankings.Store.scoreHistory` (public, unflagged, per UTC date, newest 30).
- Flagged reports are excluded everywhere (`presence.NotFlagged`). The owner's history includes `public`, `unlisted` (reads as `public`), `guild` and `private` reports; a visitor's never includes anything but public.
- Score scale is 0–2100 (`score_v2.go`); tiers are the rating module's `common … legendary` with band lines at 0/500/1000/1500/2000.
- Charts: inline SVG in a reserved aspect box, gold line on the raised ground, tier bands, tabular numerals, no legend for one series, a visible resting state, gap segments drawn as breaks with the API's reason; copy only from the copy module, verbatim from the experience spec §6.
- Tests: API with `TEST_DATABASE_URL` and `-p 1`; web `npm run check`, vitest for touched files, the lane's e2e plus `character-identity.spec.ts`, `character-phone.spec.ts`, `account-billing.spec.ts`; the character page is a shared surface, so the lane close runs the full web vitest and full Playwright suites plus the lhci shard.
- Commit trailer `Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>`.

## Review Focus

1. The owner has a private-report fight with a snapshot: it appears on their score chart as a ring (visibility `private`) and in the lists; a visitor's sparkline and public `history` rows never show it (Task 1 store test, Task 2 route test, Task 6 e2e).
2. A free owner never triggers the history request and sees the sparkline + the one sentence; the sentence is absent when the owner has no score (Task 5, Task 6).
3. A boss killed once shows a point, not a line; two or more kills draw the line (Task 4 view-model test, Task 5).
4. A 14-day stretch with a flagged report in it reads `Report flagged`; one with NULL snapshots reads `Score not kept`; otherwise `No rated fights` (Task 1 gap tests, Task 4 renders the API's words only).
5. An anonymised character, a non-owner, an unentitled owner: 404 / 404 / 402 from the route and nothing drawn (Task 2, Task 5).

---

### Task 1: The owner's history query

**Files:**
- Create: `api/internal/rating/history.go`, `api/internal/rating/history_test.go`
- Read first: `api/internal/rating/store.go:299-372` (`ReadCharacterRating`: the join), `api/internal/rankings/character_extras.go:200-235` (`scoreHistory`), `api/internal/rankings/query.go:289-305` (`CharacterFight`) and the query that fills the page's `History` rows (grep `History =` / `historyRows` in `api/internal/rankings`), `api/internal/rating/score_v2.go:18-40`, `logs/engine/rating/rating.go:135-145` (`Component.Percentile`; names `output|survival|mechanics|utility|preparation|activity`), `api/internal/rating/cards.go` (components decoding), `api/internal/rating/presence/presence.go:43-50`.

**Interfaces:**
- Produces (experience spec §10 verbatim):
  ```go
  type HistoryPoint struct { At time.Time; Score int; Tier string; ReportID string; FightIndex int; Encounter string; EncounterID int64; Difficulty string; ParsePercentile *float64; Kill bool; Visibility string }
  type HistoryEncounter struct { Encounter string; EncounterID int64; Difficulty string; Points []HistoryParsePoint }
  type HistoryParsePoint struct { At time.Time; ParsePercentile float64; ReportID string; FightIndex int }
  type HistoryGap struct { From, To string; Reason string } // YYYY-MM-DD; Reason ∈ GapNoRatedFights|GapReportFlagged|GapScoreNotKept
  const GapNoRatedFights = "No rated fights"; GapReportFlagged = "Report flagged"; GapScoreNotKept = "Score not kept"; GapDays = 14; HistoryMaxFights = 1000
  type History struct { AsOf string; Score []HistoryPoint; Encounters []HistoryEncounter; Gaps []HistoryGap; Fights []rankings.CharacterFight /* + Visibility */ }
  func (s *Store) ReadOwnerHistory(ctx context.Context, playerKey string) (History, error)
  ```
  If importing `rankings` from `rating` creates a cycle, declare `HistoryFight` in `rating` with the same JSON fields as `CharacterFight` plus `visibility`, and note it.
  Rules: `Score` = every rated fight with `score is not null and tier is not null`, report not flagged, any visibility (`unlisted` → `public`), oldest first, newest 1000; `ParsePercentile` = the `output` component's `Percentile` (nil when absent/excluded); `Difficulty` as the page names it (reuse the page's difficulty naming helper). `Encounters` = kills with a parse percentile, grouped by (encounter, difficulty), points oldest first, groups ordered by newest kill first. `Gaps` = per §10: consecutive UTC dates with a score point more than 14 days apart; reason precedence flagged → NULL snapshot → none. `Fights` = the rows behind every point, newest first, deduplicated by (report_id, fight_index), in the public `history` row shape plus `visibility` (same query as the public rows but without the public-only predicate).

- [ ] **Step 1: failing tests** (`history_test.go`, seeding through the rating test helpers and `reports` rows of each visibility, one flagged): `TestReadOwnerHistoryIncludesGuildAndPrivateButNeverFlagged` (visibility values as sent; unlisted reads public), `TestReadOwnerHistoryParsePointsComeFromTheOutputComponentOnKillsOnly`, `TestReadOwnerHistoryGapsUseTheClosedReasonSetInPrecedence` (three gaps, one of each reason, a 13-day stretch is no gap), `TestReadOwnerHistoryFightsAreDedupedNewestFirstWithVisibility`, `TestReadOwnerHistoryCapsAtTheNewestThousand` (seed 1001 → 1000, the oldest dropped).
- [ ] **Step 2: RED.** **Step 3: implement** (keep each query under 50 lines; one pass over rows builds score, encounters and fights; gaps from the dated score list). **Step 4: PASS, vet, gofmt.** **Step 5: commit** `feat(rating): the owner's full score and parse history`.

### Task 2: `GET /v1/characters/{region}/{ruleset}/{name}/history`

**Files:**
- Modify: `api/internal/rating/handler.go` (`Service` gains `Owners interface { Characters(ctx, userID int64) ([]auth.Character, error) }` or a narrower `OwnsCharacter(ctx, userID, playerKey) (bool, error)` on `auth.Store` if cheaper — pick one, nil-safe → route unmounted; `Entitlements Entitlementer` like `sims`; `Mount` adds the route under `auth.RequireSession`; new `characterHistory`), `api/internal/rating/dto.go` (`historyDTO` with the §10 JSON names: `as_of`, `score[{at, score, tier, report_id, fight_index, encounter, difficulty, parse_percentile?, kill, visibility}]`, `encounters[{encounter, encounter_id, difficulty, points[{at, parse_percentile, report_id, fight_index}]}]`, `gaps[{from, to, reason}]`, `fights[…]`), `api/cmd/api/main.go` (wire the auth store and entitlements), `api/openapi.yaml`, `api/README.md`
- Test: `api/internal/rating/handler_test.go` (pattern after `characterRating`'s tests)

Rules: not the owner → 404 `not_found` (never reveals the character); anonymised → 404; owner without `FeatureHistory` → 402 `premium_required`; anonymous → 401; `Cache-Control: private, no-store`; `as_of` today's UTC date.

- [ ] Tests: owner+entitled → 200 with every top-level key present (empty arrays, never null); owner unentitled → 402; another user → 404; anonymous → 401; anonymised → 404; `openapi_test.go` requires the path. Commit `feat(rating): owner-only character history route`.

### Task 3: The account's History row goes live

**Files:** `web/src/components/account/AccountPremiumBlock.svelte`, `web/src/lib/billing/copy.ts` (the History row per experience spec §3.10 and §6: `Open on your main character` linking to the main character's page `#history` when `entitlements.history` and a main exists; `On` when entitled with no main; `Sparkline only` otherwise — and `/premium`'s History line loses its `(ships November)` per §5 of the product spec when the lane ships; `shipsMonth.history` stays for Discord-style future rows only if still referenced), tests (`AccountPremiumBlock.test.ts`, `premium` copy tests, `account-billing.spec.ts` assertion).
- [ ] Commit `feat(web): the account's History row goes live; /premium lists history charts as shipped`.

### Task 4: Web types, copy, client and the chart view model

**Files:**
- Create: `web/src/lib/character/history-api.ts` (types mirroring the DTO; `fetchCharacterHistory(path)` via the existing envelope helper; private, no cache), `web/src/lib/character/history-copy.ts` (every §6 string, functions over data), `web/src/lib/character/history-chart.ts` (pure: collapse score points to one per UTC date keeping the last fight and a fight count; x scale proportional to time with the spec's axis tick rule; y scale 0–2100 with tier band lines; gap breaks from the API's gaps; the parse y bounds following the colour ladder (0/25/50/75 → 100); the "≥ 2 kills draws a line" rule; the newest-boss default; phone axis thinning; the readout text `1830 Epic · 11 Dec 2026 · 2 fights` per §6), `web/src/lib/character/history-select.ts` (selection/hover state → marked rows), tests for each with the spec's examples (§3.5, §3.6, §3.8, §6).
- Modify: `web/src/test-support` (a fake `/history` route for e2e and vitest).
- [ ] Commit `feat(web): character history types, copy and chart geometry`.

### Task 5: The History panel

**Files:**
- Create: `web/src/components/character/CharacterHistory.svelte` (the section per §3.2: two columns from 1280, stacked below; owner+entitled → `ScoreChart.svelte` + `ParseChart.svelte` (+ the boss tab pills in the `MyReports` pattern) each with its fights list; free owner with a score → the sentence under the card (§3.9); visitor → nothing), `ScoreChart.svelte`, `ParseChart.svelte`, `HistoryFights.svelte` (8 rows, `Show all N` to 100, reusing `CharacterFightTable`'s row with the four optional props §11.10 adds: `heading`, `columns`, `lead`, `selected`).
- Modify: `web/src/components/character/CharacterFightTable.svelte` (the four optional props, defaults preserve today's rendering), `web/src/components/Character.svelte` (mount under the Forever Score card, before Gear; the reserved box per §3.11 — the lane 3 pattern: a stored rendered height `fs.history.reserve` behind the `data-session` hint and a `data-history` hint stamped from the cached `/v1/me`'s `entitlements.history`, fallbacks per §3.11 re-measured), tests per state (owner with data, one point, a gap, a boss with one kill, loading, error, free owner, visitor, anonymised).
- [ ] Commit `feat(web): the History panel on the character page`.

### Task 6: e2e, changelog, lane close

- `web/tests/e2e/character-history.spec.ts`: premium owner with seeded points, a gap of each reason, two bosses (one with a single kill) → both charts, the pills, the lists, the readout, selection marks rows, `Show all`; free owner → sparkline + sentence and no `/history` request; visitor → card alone; anonymised → nothing; CLS ≤ 0.02 at 390/1440 hinted.
- Changelog entry (past tense, the ship day).
- Lane close (controller): full web vitest + full Playwright + lhci shard; captures per the experience spec §9 into `design/reviews/2026-10-25-character-history/`; ux-designer + wow-player reviews; one fix round; merge by fast-forward; deploy; verify.

## Rulings made by the plan (owner to check)

- The experience spec's fifteen open decisions (§11) are accepted as written, with one amendment: rings mark **private** fights only; guild and unlisted reports are evidence for the score (`presence.EvidenceReport`) and draw as ordinary points.
- "Parse percentile" is the rating engine's `output` component percentile; fights where that component was excluded are left off the parse chart and carry no `parse_percentile`.
- The history route is paid at the API (402) even though the free owner's page never calls it.
- The score chart collapses to one point per UTC date (the readout says how many fights), matching the sparkline, while the API returns per-fight rows.
- The payload is capped at the newest 1000 fights; no `since` parameter in this lane.
