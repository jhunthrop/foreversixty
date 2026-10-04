# Guides redesign — build report

Branch `web-guides`, built against `design/specs/2026-10-04-guides.md` (§12 amendments included) and the four mock boards `design/mocks/renders/guides-{index,class,spec,spec-phone}.png`.

## What shipped, per spec section

**§4.A `/guides` index (new, `GuidesClassPicker.astro`)** — `web/src/pages/guides/index.astro`, `web/src/components/guides/GuidesClassPicker.astro`. Nine-crest `grid-cols-3`/`grid-cols-1` class grid, 56px `ClassCrest`, three spec links per card in tree order (`tree_index`, via `readSpecCatalog`). Signed-in callout: server-rendered skeleton + a plain `<script type="module">` bootstrap (`web/src/lib/guides/signed-in-callout.ts`, pure view-model + unit tests) that calls the shared, cached `fetchMeOnce()` and swaps in "Your guide: {spec} {class} →" or removes the block — **no new Svelte island**, per the hard rule limiting new client islands to the existing `GuideBuildTree`.

**§4.B `/guides/<class>` class landing** — `web/src/pages/guides/[class]/index.astro`, `web/src/components/guides/SpecCard.astro`. `ArtPanel`+`ClassHeader` in `minimal` mode with the new `titleOverride` prop (renders `"{className} in Forever"` instead of the default `"{specName} {className}"`). Identity paragraph split from the rest of `index.md`'s body with a new pure splitter (`lib/guides/first-paragraph.ts`); three spec cards in tree order, each with icon (`specTabIcon`, reused), role, description (trimmed at a clause break via `lib/guides/trim-to-clause.ts`), Set DPS (only for a spec with a **real, pipeline-published** BiS file — see "Protection DPS" below), and a "Read the guide" button with a per-card aria-label. "Your spec" pill: a plain script tags the matching card's role label after reading the shared session, same no-new-island rule.

**§4.C spec guide header + rail** — `web/src/pages/guides/[class]/[spec].astro`, `web/src/components/guides/GuideActionRail.astro`. `ArtPanel` wraps `ClassHeader` (new `summaryOverride` prop — the third render mode the spec asks for: `minimal` plus a kept one-line summary) and the new rail, inside the same panel. Build card: a 120px-ish non-interactive thumbnail of the build's **primary tree** (`lib/guides/build-tree.ts`'s `decodeBuildTrees`/`primaryTree`, a build-time FS1 decode sharing `decodeFS1`/the talent-file read with the existing `GuideBuildTree` island) plus `BuildActionButtons`. Rotation card: first four level-60 lines via the existing `rotationEntryFor`/`rotationLinesFor`; the card is omitted when no rotation entry applies (never a zero-line card). Stat priority card: `lib/guides/rail-stats.ts`'s `railStatRows`, normalizing the guide's own `statPriority` order against the band's `weights[]` so the top **significant** stat (which can be melee haste) reads 1.00 — matches the mock's worked Fury numbers exactly (0.15 / n/a / n/a / 0.14 / 0.03 / 1.00), confirmed by a unit test against the real band-60 data.

**§4.D Leveling band strip** — `web/src/components/guides/LevelingBandStrip.astro`, planted between the `## Leveling` heading and the kept prose via the same `proseOnly` pattern Stat priority/Races already use. Five rows (one per band, one faction), each with `{band label} · {points} points · Load in planner`, plus a desktop-only thumbnail of that band's lit talents (`decodeBandTalents`/`litTalents`, a second small decoder for the BiS file's own dash-separated, non-FS1 talent string format). `.prose-forever p`/`ul`/`ol` gained `max-width: 68ch` in `global.css`.

**Distinct accessible names (hard rule)** — `BuildActionButtons.astro` gained optional `loadTestId`/`simTestId`/`loadAriaLabel`/`simAriaLabel` props; the rail's copy uses `"Load this build in the planner (summary)"` / `"Sim this build (summary)"` and keeps its own `guide-rail-load-build`/`guide-rail-sim-build` testids, while the copy under the full tree keeps the original `guide-load-build`/`guide-sim-build` defaults — covered by a Playwright assertion.

**`Content.astro`** gained `hideHeader` (suppresses the plain `<h1>` block, used by both new pages) and a `beforeMain` slot (the `ArtPanel`+`ClassHeader`(+rail) region, rendered full-bleed outside the 820/1120px article container) and `wide` (the class landing's 1344px non-aside width for its identity paragraph + spec-card row).

## A fix in a shared component, named explicitly

`ClassHeader.astro`'s own doc comment and the guides spec §7.B both assert the existing `max-width:1023px` breakpoint already stacks the crest above the h1 on narrow screens — it did not; `.class-header-identity` stayed `flex-direction: row` at every width. Added `flex-direction: column` for that class inside the existing media block. This is a two-line, narrowly-scoped fix that makes the documented behavior true; it also benefits `/bis`'s own phone layout (no test or screenshot asserted the old row behavior there).

## Test results (verbatim)

- `npx astro check` — **0 errors, 0 warnings, 10 hints** (760 files; all hints pre-existing or my own intentional `type="application/json"` script, unrelated to processing).
- `npm run lint` — clean, no output.
- `npm run format:check` — "All matched files use Prettier code style!"
- `npx vitest run` — **297 test files passed, 3058 tests passed, 1 skipped** (0 failed). New/updated suites: `class-dps.test.ts` (6), `rail-stats.test.ts` (2), `build-tree.test.ts` (8), `first-paragraph.test.ts` (3), `trim-to-clause.test.ts` (3), `signed-in-callout.test.ts` (4), `copy.test.ts` (updated, 2).
- `npx playwright test tests/e2e/guides*.spec.ts tests/e2e/nav.spec.ts tests/e2e/layout.spec.ts tests/e2e/home.spec.ts` (desktop + mobile projects) — **68 passed, 6 skipped** (0 failed). The 6 skips are the pre-existing `guides.spec.ts` real-build-load smoke test (both projects, guarded `FOREVER_DATA=real`-only) plus desktop-only phone-nav guards that already existed before this branch.
- New spec `tests/e2e/guides-redesign.spec.ts`: 26 tests, all passing on both projects — nine crests/27 spec links on the index, three spec cards with DPS on exactly the two ranked Warrior specs, four rotation lines with icons + Load/Sim buttons + six stat-priority lines on the Fury rail, distinct accessible names on the two Load-build buttons, five Leveling rows each carrying `?talents=` in its planner href, phone stacking with no dropped lines, thumbnails dropped on phone, and no horizontal overflow at 390/1440 across all three pages.

## Captures

`design/mocks/renders/build/` (git-ignored, not committed): `guides-{1440,390}.png`, `guides-warrior-{1440,390}.png`, `guides-warrior-fury-{1440,390,1024,1920}.png`. Captured against the **real** build data (`node scripts/sync-data.mjs` without `FOREVER_DATA`, which is the script's own default) so every icon and number is real; fixture mode was restored afterward (`FOREVER_DATA=fixture npm run sync`) before the final test/lint gates, matching this repo's own `pretest`/CI convention. Full-page screenshots scroll through the page before shooting so the `client:visible` `GuideBuildTree` island hydrates (a naive screenshot otherwise catches its loading skeleton).

### Side-by-side differences found, and what I did

1. **Spec-tab/spec-card icons 404'd in my first capture pass.** Cause: I had synced `public/data` in `fixture` mode (a ~211-icon curated subset) for an earlier test run; the guide pages' real icon names (`ability_rogue_ambush`, etc.) aren't in it. This is a capture-environment issue, not a product defect — these pages always read `data/builds/<build>/` directly at Astro build time regardless of `FOREVER_DATA`. Fixed by re-syncing in real mode before the capture pass.
2. **`GuideBuildTree` (kept, unchanged island) showed its loading skeleton**, not the ready tree, in a `networkidle`-only screenshot. Fixed the capture script to scroll through the full page (triggering its `client:visible` intersection) before shooting.
3. **Crest sat beside the h1 at phone width on the class landing**, contradicting spec §7.B. Root cause and fix: see "A fix in a shared component" above.
4. **Footer differs from the mock's hand-rolled one** (About/Sources/Changelog/Contribute/Premium/Get set up/Terms/Privacy/Refunds vs. the mock's About/Sources/Discord/GitHub). Expected — the mock's footer was a simplified stand-in; the real site's `Base.astro` footer is unchanged, pre-existing chrome, out of this spec's scope.
5. **A confidence line appears above Sources on real captures** ("Reported by one outlet, not yet confirmed" on the class landing, "Our inference from what has been said; may change" on the spec guide) that the mock boards don't show. Expected — `confidencePlacement="footer"` is exactly what spec §3 asks `Content.astro` to keep doing; the mock simply didn't render every text element.
6. **The spec guide is much taller at 390px than the mock's phone board** (~8400px vs. the mock's ~4456px). Cause: the full, real three-tree "Talents and builds" grid (kept, unchanged `GuideBuildTree`) stacks all three trees vertically on narrow screens — nineteen-plus rows apiece — which the hand-rolled mock never modeled at phone width (it only ever drew the fixed-size preview). Not a defect: no horizontal overflow, every control still clears 44px (confirmed by the e2e sweep), and this is pre-existing component behavior this rebuild did not touch.

## Left out, and why

- **Per-visitor "Your spec" card match and the index signed-in callout are vanilla-script-driven, not server-known or Svelte islands** — the task's hard rule caps new client islands at the existing `GuideBuildTree`; a plain module script reading the shared `/v1/me` cache satisfies the spec's behavior with zero added framework runtime.
- **Lighthouse (`lhci`) was not run** — not in this task's required TESTS list; the budget rows (`web/lighthouserc.json`) are unchanged, and the rail adds no new network round trip beyond `GuideBuildTree`'s existing one, per spec §8.
- **Only Warrior (index/class/fury) was captured**, per the explicit capture list; the other 26 spec guides and 8 other class landings were not individually screenshotted, though `build.ts`/`astro build` succeeded for all 36 guide pages and spot-checked output (DPS figures, card counts, leveling-row counts) on a second class (Protection) confirmed graceful degradation.
- **The Stat priority rail card on a spec with no ranked BiS file** (e.g. Protection) renders with just its label and no rows, rather than being omitted entirely — the spec's state table doesn't rule on this case explicitly; omitting would also have been defensible, but an empty-but-present card never fabricates a number, so this was the lower-risk default. **Superseded in round 1 below** — this was fixed; the card is now omitted.

---

## Round 1 (owner side-by-side fix pass)

Three findings from a side-by-side against `design/mocks/renders/build/guides-warrior-1440.png`.

### 1. [must fix, tenet 4] Spec-card descriptions were clipped mid-sentence

`trimToClause` (90-character cap at a clause break) was applied to the class landing's spec-card description, so Arms/Fury/Protection all ended on a trailing comma instead of their guide's full sentence. Removed the call entirely; `SpecCard` now renders `guide.data.description` verbatim and the card grows to fit (no `max-height`/line-clamp anywhere in its styles, confirmed). The now-unused `lib/guides/trim-to-clause.ts` and its test were deleted outright (dead code, never kept "just in case"). Added `tests/e2e/guides-redesign.spec.ts`'s `"every spec card shows its guide's full frontmatter description, never clipped"` test, which reads each of the three warrior guides' `description` frontmatter straight off disk and asserts the rendered card text equals it exactly.

### 2. [must fix] Stat priority rail card with no rows

Ruling (applied in `GuideActionRail.astro`): a card with nothing real to show is **omitted**, never rendered with just a label. The rail is 3-up when Build + Rotation + Stat priority all have content, 2-up when one of Rotation/Stat priority is empty, 1-up (Build only) when both are. On real data today: **Protection is 1-up**, not 2-up as first assumed — its band-60 BiS file doesn't exist (no Stat priority numbers) *and* every one of its `addon-data.json` rotation entries carries zero lines (no rotation prose has been written for the tank spec yet), so both cards are absent, not just one. Pinned with `tests/e2e/guides-redesign.spec.ts`'s `/guides/warrior/protection` test, which asserts both cards are absent (`toHaveCount(0)`, not just hidden) and the rail's own `grid-template-columns` computes to a single column.

Implementation note: the column count is set through a CSS custom property (`--guide-rail-cols`) read by the stylesheet's `grid-template-columns: repeat(var(--guide-rail-cols, 3), 1fr)`, **not** a literal inline `grid-template-columns` — the first attempt used an inline style, which (correctly, by CSS cascade rules) outranks every stylesheet rule including the phone `@media (max-width: 1023px)` override, silently pinning the rail at 3 columns past that breakpoint and overflowing 390px. Caught by the recapture/test pass, not the side-by-side itself; see "Found while fixing" below.

### 3. [polish] Rail buttons now use the shared secondary-button recipe

`BuildActionButtons.astro`'s two links rendered in sentence case (`text-[14px] font-semibold`) with no relation to `design/DESIGN-SYSTEM.md`'s secondary button (uppercase, 12px, 700 weight, 0.06em tracking — `SECONDARY_BUTTON`/`SECONDARY_BUTTON_FIXED`, `lib/planner/styles.ts`). Switched to `SECONDARY_BUTTON_FIXED` specifically, not the plain `SECONDARY_BUTTON`: this component's own pre-existing unit test (`BuildActionButtons.test.ts`, predates this branch) pins a 44px hit target on **both** links at **every** breakpoint, a guarantee only the `_FIXED` variant (no `md:h-9` desktop shrink) keeps. Updated that test's assertion from the old literal `min-h-11` class-string check to the new recipe's own markers (`h-11`, no `md:h-9`, `uppercase`) rather than leaving a stale, failing assertion.

### Found while fixing (not in the owner's three findings)

- **Phone overflow regression at 390px** on `/guides/warrior/fury`, introduced by fix #2's own first implementation (the inline `grid-template-columns` described above). Caught by `npx playwright test tests/e2e/guides*.spec.ts`, not visually — the "no horizontal overflow" sweep failed with `scrollWidth: 543` against a 390px viewport. Fixed with the CSS-custom-property approach described above.
- **`/guides/warrior/protection`'s rail is 1-up, not 2-up** — my own round-1 test first assumed 2-up (Build + Rotation) by analogy with the owner's finding text; the real `addon-data.json` has zero rotation lines for every warrior-protection entry, so Rotation is also absent. Corrected the test to assert the true 1-up state rather than adjust the product code to match a wrong assumption.

### Re-run results (verbatim)

- `npx astro check` — 0 errors, 0 warnings, 10 hints (758 files).
- `npm run lint` — clean.
- `npm run format:check` — "All matched files use Prettier code style!"
- `npx vitest run` (guides-scoped: `src/lib/guides/`, `src/pages/guides/`, `src/content/guides/`) — 13 test files, 305 tests, all passing.
- `npx vitest run` (full suite) — **296 test files passed, 3055 tests passed, 1 skipped** (0 failed; one file count lower than the first build report's 297/3058 because `trim-to-clause.ts`/`.test.ts` were deleted).
- `npx playwright test tests/e2e/guides*.spec.ts` (desktop + mobile) — **30 passed, 2 skipped** (0 failed; the 2 skips are the pre-existing `guides.spec.ts` real-build-load smoke test, both projects, guarded `FOREVER_DATA=real`-only).

### Recaptures

`guides-warrior-1440.png` and `guides-warrior-fury-1440.png` re-shot against real build data (same `sync`/fixture-restore discipline as the first pass). Descriptions now run the full sentence on all three spec cards (Protection's: "Talents, tanking priority, stat priority, and race picks for Protection Warrior in Forever, with beta-versus-projection called out."); Fury's rail buttons read `LOAD THIS BUILD`/`SIM THIS BUILD` in both the rail and under the full tree.
