# Premium Value Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** `/premium` answers a solo player's question in Region A (Free, Premium with six benefit lines that say what each does for the character, and a real server Top Gear run as proof) and a guild leader's question in Region B (`#guild-leaders`: Guild and Raid guild with computed per-member and per-raider cost lines, the history promise, and How the guild plans work with `Verified members` first). `/sim`'s invented `1,245 DPS` example is replaced by the same real capture.

**Architecture:**
- `plans.ts` gains the raid sizes, the trial reminder days and the two cost-line functions; every figure on the page is computed from it, `lib/sim/precision.ts` (`CAPS`, `LANE_ITERATION_CEILING`) or the example data file.
- One data file, `web/src/data/premium-examples.json`, written by one capture script, feeds both pages through one reader (`premium-examples.ts`) and one image source (`example-image.ts`). The reader validates the file and picks the caption template by `preset`.
- `plans-page-copy.ts` and `landing-copy.ts` hold every B5 string; a spec-table test pins each to the amended spec.
- `premium.astro` is re-laid into Region A (three subgrid columns) and Region B (`#guild-leaders`). `PlanCard`, `PriceBlock` and `GuildRules` change; `PremiumExampleCard.astro` is new. The page stays static.
- The example is a ship blocker: a controller step (Task 19) runs the real server sim on the owner's account and writes the data file.

**Tech Stack:** Astro + Svelte 5, Tailwind, vitest (Astro container renders), Playwright + axe, the Python mock generator `design/mocks/gen_launch_pricing.py`.

**Spec:** `design/specs/2026-10-27-launch-pricing.md`, section "Amendment 2026-10-27b: Premium value and guild leaders" (B1 to B10), as amended by Task 0 with the controller's preset ruling. B5 is the copy, verbatim; B8 the acceptance checks; B9 capture rows 15 to 22.

**Controller ruling (2026-10-27), applied to Tasks 0, 2, 4, 7, 10, 17, 18, 19 and 21:** the example is a real server Top Gear run on a **level-60 Fury Warrior in the site's preset gear**, run from the owner's account (the owner's own warrior, `Obnoxious Yell`, is level 21). The data file carries `simId`, `characterName`, `level`, `specLabel`, `preset`, `combinations`, `ranAt`. When `preset` is true the caption reads `A level-{level} {specLabel} in the site's preset gear: {combinations} combinations on our servers, {ranAt as d MMM yyyy}.`; when false it uses the spec's own-gear form `{characterName}, {specLabel}: {combinations} combinations on our servers, {ranAt as d MMM yyyy}.` Level and spec are read from the data file, never typed in copy. The run itself is a controller step that needs the owner's session.

## Global Constraints

- **Figures are never typed** (B8 check 2, the standing rule that prose never quotes generated numbers). Prices come from `web/src/lib/billing/plans.ts`; `CAPS.server` (5,000), `CAPS.browser` (400), `LANE_ITERATION_CEILING.server` (100,000) and `.browser` (30,000) from `web/src/lib/sim/precision.ts`; `GUILD_TIER_MEMBER_LIMIT` (25), `GUILD_TRIAL_DAYS` (7), `RAID_SIZES` (`[40, 20, 10]`) and `TRIAL_REMINDER_DAYS` (3) from `plans.ts`; every example word and number from `web/src/data/premium-examples.json`.
- **Copy is verbatim from B5.** Every string goes through `plans-page-copy.ts` or `landing-copy.ts`, and a spec-table test reads it from the spec (`web/src/test-support/spec-table.ts`). Nothing is retyped from memory.
- **Only live benefits render** (B5 "Not on the page, on purpose"; R-P5). The words `coach`, `chip in`, `retention`, `compare more than`, and an exclamation mark never render on `/premium`. Until the trial email ships, no text contains `email you` near `trial`. The held coach line (`premium.featuresHeld.coach`) stays out of `premium.features`.
- **Static page.** `/premium` ships no island besides the layout's `CharacterSelector` (the existing e2e asserts exactly one; the subscribe-and-join lane's stored Premium hint depends on it). The example image is the only new request; the AVIF is at most 45 KB. No new design token, no new shared control: every control is a recipe the spec names (`card`, the tick list, `SECONDARY_BUTTON_FIXED`, the inline gold link, `section-title`).
- **Layout** (B3, B7). The 1344 px column is centred at every width up to 2000 and beyond; the band is full bleed. Three columns from 1024, two from 768 to 1023 (the example or the history promise spans underneath, the example capped at 480 px), one below 768. One-up rows below 768 keep the column width.
- **No layout shift** (B8 check 4). Both intervals and both founder states sit in fixed cells: the guild price block is 152 px from 768 and 144 px on phone, the cost line a 54 px box, the trial slot at least 40 px, and the example image carries width 440 and height 330 in a reserved 4:3 box on `--card-top`.
- **Gold** (B6) is Premium's button, the founder pill, the founder lines, the band's guild link and `Open this result`. Nothing is filled. Guild cards have no gold top highlight and neutral buttons; the example card has no highlight.
- **Touch and keyboard** (B4). Every link and button is at least 44 px tall at 360. Focus rings are 2 px gold at 2 px offset. Tab order: band link, interval control, Premium button, `Open this result`, Guild button, Raid guild button, the FAQ summaries. The image is not focusable.
- **No Battle.net wording** in `Verified members` (B10 ruling 5).
- **Ship blocker.** The page never ships while the data file holds the placeholder (B3.2; Task 19).
- **Tasks run vitest, `npm run check`, `npm run lint` and `npx prettier --check <files>` only.** The controller runs `astro build`, Playwright, the lhci shard and the captures at branch level (Tasks 19 and 21). The worktree needs `npm ci` in `web/` once.
- **File limits.** TS and Svelte files stay at or under 400 lines; split before passing it.
- **Spec state.** `design/specs/2026-10-27-launch-pricing.md` carries the amendment uncommitted in this worktree; Task 0 commits it with the ruling.
- **Commits.** Conventional types, commit by path (`git commit -m "<msg>" -- <paths>`), with the trailer the session's attribution reminder gives.

## Review Focus

1. **The data file's `preset` flag.** A file with `preset` missing, `"true"` as a string, or `null` must throw at import naming the field, not fall back to the own-gear caption (which would print a preset run as if it were a named character). Pinned in Task 2 (`premium-examples.test.ts › "rejects a missing or non-boolean preset"`, `"picks the preset template when preset is true"`) and Task 10 / Task 17 (both fixtures, alt equals caption).
2. **The $0.625 rounding case.** Raid guild monthly list `$25 / 40` is `$0.625`; `formatAmount` must give `$0.63`, and the order must be 40, 20, 10. Pinned in Task 1 (`plans.test.ts › "perRaiderLine monthly list rounds $0.625 to $0.63"`).
3. **`ranAt` across the UTC date line.** A run saved at `2026-10-27T23:30:00-07:00` is 28 Oct in UTC; the caption, the alt text and the `/sim` card must all read the same date, through `utcDayMonthYear`. Pinned in Task 2 (`premium-examples.test.ts › "formats ranAt in UTC across the date line"`).
4. **The image failing to load.** A blocked request must leave the reserved 4:3 box on `--card-top`, the alt text, the caption and `Open this result`, with no height change. Pinned in Task 10 (`PremiumExampleCard.test.ts › "renders caption and link with the image absent"`) and Task 21 (capture row 21, blocked image request).
5. **The per-raider line at its widest.** The three-line Raid guild cost line at 360, 1024 and 1280 must stay inside its 54 px box in both intervals and both founder states, so the 152/144 px price block never grows. Pinned in Task 12 (`PriceBlock.test.ts › "cost line sits in a 54 px box in all four forms"`) and Task 21 (`premium.spec.ts › "guild price blocks hold 152/144 px across interval and founder"`).

---

### Task 0: Record the preset ruling in the spec

**Files:**
- Modify: `design/specs/2026-10-27-launch-pricing.md` (§B3.2, §B5 example rows and the simulator landing row, §B8 check 7, §B10 ruling 4)

**Interfaces:**
- Produces (spec text the spec-table tests read):
  - §B3.2 "What the image is (ruled)": the run is one finished server Top Gear run on a level-60 Fury Warrior in the site's preset gear, run from the owner's Premium account; the sidecar lists `simId`, `characterName`, `level`, `specLabel`, `preset`, `combinations`, `ranAt`. Marked `(controller ruling 2026-10-27)`.
  - §B5: the single `examples.serverRun.caption` row is replaced by two rows, keys `examples.serverRun.caption.preset` and `examples.serverRun.caption.ownGear` (both "also the image alt"):
    - preset: `A level-{level} {specLabel} in the site's preset gear: {combinations} combinations on our servers, {ranAt as d MMM yyyy}.`
    - ownGear: `{characterName}, {specLabel}: {combinations} combinations on our servers, {ranAt as d MMM yyyy}.`
  - §B5 simulator landing row: "the caption uses the same sidecar and templates as `examples.serverRun.caption.preset` / `.ownGear`".
  - §B8 check 7: `The example's caption words equal the data file's values, and the data file carries a finished server run with preset true or a character on the owner's account. Open this result resolves to a saved result that answers 200 signed out.` Marked `(controller ruling 2026-10-27)`.
  - §B10 ruling 4: the server-sim example is a level-60 Fury Warrior in the site's preset gear, run from the owner's account.

- [ ] **Step 1: Edit the four places.** Keep the B5 table's three columns (`Where | Key | String`), each template in one backtick pair, so Task 3's helper reads it.
- [ ] **Step 2: Check the table parses.** Run `cd web && node -e "const {readFileSync}=require('fs');const s=readFileSync('../design/specs/2026-10-27-launch-pricing.md','utf8');for(const k of ['examples.serverRun.caption.preset','examples.serverRun.caption.ownGear']){const r=s.split('\n').find(l=>l.includes('\`'+k+'\`'));if(!r)throw new Error(k);console.log(r.split('|')[3].trim())}"`. Expected: both templates print.
- [ ] **Step 3: Commit.** `git commit -m "docs(spec): premium example is a level-60 preset Fury Warrior (controller ruling)" -- design/specs/2026-10-27-launch-pricing.md`

### Task 1: Raid sizes, the trial reminder and the cost lines in `plans.ts`

**Files:**
- Modify: `web/src/lib/billing/plans.ts`, `web/src/lib/billing/plans.test.ts`
- Create: `web/src/lib/billing/plan-costs.ts`, `web/src/lib/billing/plan-costs.test.ts`

**Interfaces:**
- Consumes: `priceFor(plan, interval, founder)`, `formatAmount(usd)`, `GUILD_TIER_MEMBER_LIMIT`, `type Interval`.
- Produces:
  - `export const RAID_SIZES = [40, 20, 10] as const;` and `export type RaidSize = (typeof RAID_SIZES)[number];`
  - `export const TRIAL_REMINDER_DAYS = 3;`
  - `export function perRaiderAmounts(interval: Interval, founder: boolean): readonly string[]` — the Raid guild price divided by each size, through `formatAmount`, in `RAID_SIZES` order.
  - `export function perMemberAmount(interval: Interval, founder: boolean): string` — the Guild price divided by `GUILD_TIER_MEMBER_LIMIT`, through `formatAmount`.
  - PRE-FLIGHT RULING: both functions live in a NEW file `web/src/lib/billing/plan-costs.ts` (tested in `plan-costs.test.ts`), importing `PLANS`, `RAID_SIZES`, `GUILD_TIER_MEMBER_LIMIT`, `priceFor` and `formatAmount` from `./plans` — never inside `plans.ts`. Task 7's fixture-swap test mocks `./plans`; a helper defined inside `plans.ts` would stay bound to the real values and the test could not pass. Task 4 imports from `./plan-costs`.
  - (The sentences are built in Task 4's `price.perRaider` / `price.perMember`; the function names the task list gives, `perRaiderLine` / `perMemberLine`, are Task 4's copy functions, re-exported from `plans-page-copy.ts`.)

- [ ] **Step 1: Write the failing tests** in `plans.test.ts`:
  ```ts
  describe('raid sizes and cost amounts (B3.3)', () => {
    it('RAID_SIZES leads with 40', () => expect(RAID_SIZES).toEqual([40, 20, 10]));
    it('TRIAL_REMINDER_DAYS is 3', () => expect(TRIAL_REMINDER_DAYS).toBe(3));
    it.each([
      ['yearly', true, ['$5', '$10', '$20']],
      ['yearly', false, ['$6.25', '$12.50', '$25']],
      ['monthly', true, ['$0.50', '$1', '$2']],
      ['monthly', false, ['$0.63', '$1.25', '$2.50']],
    ] as const)('perRaiderAmounts %s founder=%s', (interval, founder, want) => {
      expect(perRaiderAmounts(interval, founder)).toEqual(want);
    });
    it('perRaiderLine monthly list rounds $0.625 to $0.63', () => {
      expect(perRaiderAmounts('monthly', false)[0]).toBe('$0.63');
    });
    it.each([
      ['yearly', true, '$4.80'],
      ['yearly', false, '$6'],
      ['monthly', true, '$0.48'],
      ['monthly', false, '$0.60'],
    ] as const)('perMemberAmount %s founder=%s', (interval, founder, want) => {
      expect(perMemberAmount(interval, founder)).toBe(want);
    });
  });
  ```
  The expected amounts are B3.3's "Raid guild values today" and "Guild values today", copied.
- [ ] **Step 2: Run the tests to verify they fail.** `cd web && npx vitest run src/lib/billing/plans.test.ts`. Expected: FAIL (no export).
- [ ] **Step 3: Implement.** If `formatAmount` gives `$0.62` for `0.625` (binary float), round on cents first (`Math.round(usd * 100 + Number.EPSILON) / 100`) inside `formatAmount` and re-run the existing `$4.17` / `$12.50` cases.
- [ ] **Step 4: Run the tests to verify they pass**, then `npm run check && npm run lint && npx prettier --check src/lib/billing/plans.ts src/lib/billing/plans.test.ts`.
- [ ] **Step 5: Commit.** `git commit -m "feat(web): raid sizes, trial reminder days and the guild cost amounts" -- web/src/lib/billing/plans.ts web/src/lib/billing/plans.test.ts`

### Task 2: The example data file and its reader

**Files:**
- Create: `web/src/data/premium-examples.json`, `web/src/lib/billing/premium-examples.ts`, `web/src/lib/billing/premium-examples.test.ts`

**Interfaces:**
- Consumes: `utcDayMonthYear` from `web/src/lib/dates.ts`.
- Produces:
  - Data file (placeholder until Task 19):
    ```json
    { "serverRun": { "simId": "placeholder", "characterName": "", "level": 60, "specLabel": "Fury Warrior", "preset": true, "combinations": 0, "ranAt": "2026-10-27T00:00:00Z" } }
    ```
  - `export interface ServerRunExample { simId: string; characterName: string; level: number; specLabel: string; preset: boolean; combinations: number; ranAt: string }`
  - `export interface CaptionFields { characterName: string; level: string; specLabel: string; combinations: string; ranAt: string }` (display strings)
  - `export interface CaptionTemplates { preset: (f: CaptionFields) => string; ownGear: (f: CaptionFields) => string }`
  - `export const PLACEHOLDER_SIM_ID = 'placeholder';`
  - `export function parseServerRunExample(raw: unknown): ServerRunExample` — throws `Error('premium-examples.json: serverRun.<field> …')` on a missing or mistyped field: `simId` non-empty string; `characterName` string (may be empty only when `preset` is true); `level` integer 1..60; `specLabel` non-empty string; `preset` boolean; `combinations` positive integer unless placeholder; `ranAt` a parseable RFC 3339 instant.
  - `export const serverRunExample: ServerRunExample` — the parsed file, at import (fail fast at build).
  - `export function isPlaceholder(ex: ServerRunExample): boolean` — `ex.simId === PLACEHOLDER_SIM_ID`.
  - `export function captionFields(ex: ServerRunExample): CaptionFields` — `combinations.toLocaleString('en-US')`, `level` as a string, `ranAt` through `utcDayMonthYear` (`d MMM yyyy`).
  - `export function exampleCaption(ex: ServerRunExample, templates: CaptionTemplates): string` — `ex.preset ? templates.preset(f) : templates.ownGear(f)`.
  - `export function exampleHref(ex: ServerRunExample): string` — `` `/sim/${encodeURIComponent(ex.simId)}` ``.

- [ ] **Step 1: Write the failing tests:**
  ```ts
  const VALID = { simId: 'abc123', characterName: 'Kraggor', level: 60, specLabel: 'Fury Warrior', preset: true, combinations: 4812, ranAt: '2026-10-27T18:00:00Z' };
  const T: CaptionTemplates = { preset: (f) => `P|${f.level}|${f.specLabel}|${f.combinations}|${f.ranAt}`, ownGear: (f) => `O|${f.characterName}|${f.specLabel}|${f.combinations}|${f.ranAt}` };
  it('parses a valid file', () => expect(parseServerRunExample({ serverRun: VALID })).toEqual(VALID));
  it.each(['simId', 'level', 'specLabel', 'combinations', 'ranAt'])('rejects a file missing %s', (k) => {
    const { [k]: _gone, ...rest } = VALID;
    expect(() => parseServerRunExample({ serverRun: rest })).toThrow(k);
  });
  it('rejects a missing or non-boolean preset', () => {
    const { preset: _p, ...rest } = VALID;
    expect(() => parseServerRunExample({ serverRun: rest })).toThrow('preset');
    expect(() => parseServerRunExample({ serverRun: { ...VALID, preset: 'true' } })).toThrow('preset');
    expect(() => parseServerRunExample({ serverRun: { ...VALID, preset: null } })).toThrow('preset');
  });
  it('rejects an empty characterName on an own-gear run', () => {
    expect(() => parseServerRunExample({ serverRun: { ...VALID, preset: false, characterName: '' } })).toThrow('characterName');
  });
  it('picks the preset template when preset is true', () => expect(exampleCaption(VALID, T)).toBe('P|60|Fury Warrior|4,812|27 Oct 2026'));
  it('picks the own-gear template when preset is false', () => expect(exampleCaption({ ...VALID, preset: false }, T)).toBe('O|Kraggor|Fury Warrior|4,812|27 Oct 2026'));
  it('formats ranAt in UTC across the date line', () => expect(captionFields({ ...VALID, ranAt: '2026-10-27T23:30:00-07:00' }).ranAt).toBe('28 Oct 2026'));
  it('flags the placeholder', () => {
    expect(isPlaceholder({ ...VALID, simId: PLACEHOLDER_SIM_ID })).toBe(true);
    expect(isPlaceholder(VALID)).toBe(false);
  });
  it('the shipped file parses', () => expect(() => serverRunExample).not.toThrow());
  it('links to the saved result', () => expect(exampleHref(VALID)).toBe('/sim/abc123'));
  ```
- [ ] **Step 2: Run the tests to verify they fail.** `npx vitest run src/lib/billing/premium-examples.test.ts`. Expected: FAIL (module missing).
- [ ] **Step 3: Implement** the file and the reader. Each validator is a small function; the module stays under 120 lines.
- [ ] **Step 4: Run the tests to verify they pass**, then check, lint and prettier on the three files.
- [ ] **Step 5: Commit.** `git commit -m "feat(web): premium example data file and its validating reader" -- web/src/data/premium-examples.json web/src/lib/billing/premium-examples.ts web/src/lib/billing/premium-examples.test.ts`

### Task 3: Premium copy

**Files:**
- Modify: `web/src/lib/billing/plans-page-copy.ts`, `web/src/lib/billing/plans-page-copy.test.ts`, `web/src/test-support/spec-table.ts`

**Interfaces:**
- Consumes: `CAPS`, `LANE_ITERATION_CEILING` from `web/src/lib/sim/precision.ts`.
- Produces:
  - `spec-table.ts`: `export function specKeyCell(specPath: URL, heading: string, key: string): string[]` — the backticked segments of the String column (the third) of the row whose Key column contains `` `key` ``. B5's table is `Where | Key | String`, so `specCell` (which reads the second column) cannot read it.
  - `plansPageCopy.premium.support: string` — `Free stays complete. Premium pays for the servers and keeps the site going.`
  - `plansPageCopy.premium.features: readonly string[]` — the six B5 benefit lines, figures interpolated with `toLocaleString('en-US')`.
  - `plansPageCopy.premium.scope: string` — B5 `premium.scope`.
  - `plansPageCopy.premium.featuresHeld.coach: string` — B5's held line, verbatim.
  - The four §6.1 Premium features are deleted.

- [ ] **Step 1: Write the failing tests** (in `plans-page-copy.test.ts`):
  ```ts
  const B5 = '### B5. Copy, verbatim';
  const key = (k: string) => specKeyCell(SPEC, B5, k);
  const fill = (s: string, m: Record<string, string>) => Object.entries(m).reduce((t, [k, v]) => t.replaceAll(`{${k}}`, v), s);
  it('premium.support', () => expect(copy.premium.support).toBe(key('premium.support')[0]));
  it('premium.features[0] from CAPS', () => expect(copy.premium.features[0]).toBe(
    fill(key('premium.features[0]')[0], { 'CAPS.server': CAPS.server.toLocaleString('en-US'), 'CAPS.browser': CAPS.browser.toLocaleString('en-US') })));
  it('premium.features[1]', () => expect(copy.premium.features[1]).toBe(key('premium.features[1]')[0]));
  it('premium.features[2] from LANE_ITERATION_CEILING', () => expect(copy.premium.features[2]).toBe(
    fill(key('premium.features[2]')[0], {
      'LANE_ITERATION_CEILING.server': LANE_ITERATION_CEILING.server.toLocaleString('en-US'),
      'LANE_ITERATION_CEILING.browser': LANE_ITERATION_CEILING.browser.toLocaleString('en-US'),
    })));
  it.each([3, 4, 5])('premium.features[%i]', (i) => expect(copy.premium.features[i]).toBe(key(`premium.features[${i}]`)[0]));
  it('has exactly six features', () => expect(copy.premium.features).toHaveLength(6));
  it('premium.scope', () => expect(copy.premium.scope).toBe(key('premium.scope')[0]));
  it('held coach line is not in premium.features', () => {
    expect(copy.premium.features).not.toContain(copy.premium.featuresHeld.coach);
    expect(copy.premium.features.join(' ')).not.toMatch(/coach/i);
  });
  it('old Premium features are gone', () => expect(allStrings(copy)).not.toContain('History charts of your score and parses'));
  ```
  The held line is checked against the spec by `expect(copy.premium.featuresHeld.coach).toBe(<the backticked line under "Held until built">)` read with a small `specLineAfter(SPEC, '**Held until built')` helper in the test file.
- [ ] **Step 2: Run to fail.** `npx vitest run src/lib/billing/plans-page-copy.test.ts`. Expected: FAIL.
- [ ] **Step 3: Implement** `specKeyCell` and the copy. Update the existing `Premium card` §6.1 test row to the new shape (the §6.1 Premium row is overridden by B5).
- [ ] **Step 4: Run to pass**, check, lint, prettier.
- [ ] **Step 5: Record B8 check 1** in the lane report: one file or route per benefit line (`lib/sim/precision.ts` + `bulk-server-run.ts` and the job's cap check; `ServerRuns.svelte`, `FinishedSimLine.svelte`; `CharacterHistory.svelte`; the account Discord block and the notify job; `SupporterChip.svelte` on the character header and in `GuildRosterTable.svelte`). A line with no proof comes off.
- [ ] **Step 6: Commit.** `git commit -m "feat(web): Premium benefit lines, support line and scope note" -- web/src/lib/billing/plans-page-copy.ts web/src/lib/billing/plans-page-copy.test.ts web/src/test-support/spec-table.ts`

### Task 4: Guild-region copy and the example strings

**Files:**
- Modify: `web/src/lib/billing/plans-page-copy.ts`, `web/src/lib/billing/plans-page-copy.test.ts`

**Interfaces:**
- Consumes: `perRaiderAmounts`, `perMemberAmount`, `RAID_SIZES`, `GUILD_TIER_MEMBER_LIMIT`, `GUILD_TRIAL_DAYS` (Task 1); `CaptionTemplates`, `exampleCaption`, `serverRunExample` (Task 2).
- Produces:
  - `band.guildLead` `Running a guild?`, `band.guildLink` `Guild and Raid guild are further down`.
  - `guildRegion: { heading, lede, members }` — B5 verbatim.
  - `guild.audience` (from `GUILD_TIER_MEMBER_LIMIT`), `guild.features` (five lines), `raidGuild.audience` (unchanged), `raidGuild.features` (five lines, loot council first, `Everything in Guild: …` last).
  - `price.perMember(interval: Interval, founder: boolean): string` and `price.perRaider(interval: Interval, founder: boolean): string` — the B5 templates filled from Task 1's amounts; `RAID_SIZES` gives the `40-person`, `20` and `10` words.
  - `export const perMemberLine = plansPageCopy.price.perMember; export const perRaiderLine = plansPageCopy.price.perRaider;`
  - `trial: string` — B5 "Trial slot (guild cards), today".
  - `rules`: `Verified members` first (B5 sentence), `Paying together` directly after `Size`, `Trial` changed to B5's "today" sentence. The `Everything in Guild` lead row is deleted.
  - `examples.serverRun: { label: string; caption: CaptionTemplates; link: string }` with `caption.preset` and `caption.ownGear` from Task 0's two B5 rows.
  - `export function serverRunCaption(ex: ServerRunExample = serverRunExample): string` — `exampleCaption(ex, plansPageCopy.examples.serverRun.caption)`.

- [ ] **Step 1: Write the failing tests:**
  ```ts
  it('band guild link', () => expect([copy.band.guildLead, copy.band.guildLink]).toEqual(key('band.guildLead')));
  it.each(['heading', 'lede', 'members'] as const)('guildRegion.%s', (k) => expect(copy.guildRegion[k]).toBe(key(`guildRegion.${k}`)[0]));
  it('guild.audience', () => expect(copy.guild.audience).toBe(fill(key('guild.audience')[0], { GUILD_TIER_MEMBER_LIMIT: '25' })));
  it('guild.features', () => expect(copy.guild.features).toEqual(key('guild.features').map((s) => fill(s, { GUILD_TIER_MEMBER_LIMIT: String(GUILD_TIER_MEMBER_LIMIT) }))));
  it('raidGuild.features, loot council first', () => {
    expect(copy.raidGuild.features).toEqual(key('raidGuild.features'));
    expect(copy.raidGuild.features[0]).toMatch(/^Loot council:/);
  });
  it.each([['yearly', true, 1], ['monthly', true, 3]] as const)('perMember %s matches the B5 example', (i, f, idx) =>
    expect(perMemberLine(i, f)).toBe(key('price.perMember(interval, founder)')[idx]));
  it.each([['yearly', true, 1], ['monthly', true, 3]] as const)('perRaider %s matches the B5 example', (i, f, idx) =>
    expect(perRaiderLine(i, f)).toBe(key('price.perRaider(interval, founder)')[idx]));
  it('perRaider monthly list carries $0.63', () => expect(perRaiderLine('monthly', false)).toBe('$0.63 per raider a month in a 40-person raid, $1.25 in a 20, $2.50 in a 10.'));
  it('trial', () => expect(copy.trial).toBe(fill(fill(key('trial')[0], { GUILD_TRIAL_DAYS: '7' }), { 'GUILD_TRIAL_DAYS + 1': '8' })));
  it('rule order: Verified members first, Paying together after Size', () => {
    const terms = copy.rules.map((r) => r.term);
    expect(terms[0]).toBe('Verified members');
    expect(terms[terms.indexOf('Size') + 1]).toBe('Paying together');
  });
  it('Verified members rule is B5 verbatim and has no Battle.net', () => {
    const rule = copy.rules[0];
    expect([rule.term, rule.sentence]).toEqual(key('rules[Verified members]'));
    expect(rule.sentence).not.toMatch(/Battle\.net/);
  });
  it('Paying together rule', () => expect(['Paying together', copy.rules.find((r) => r.term === 'Paying together')?.sentence]).toEqual(key('rules[Paying together]')));
  it('Trial rule today', () => expect(copy.rules.find((r) => r.term === 'Trial')?.sentence).toBe(fill(fill(key('rules[Trial]')[0], { GUILD_TRIAL_DAYS: '7' }), { 'GUILD_TRIAL_DAYS + 1': '8' })));
  it('no Everything in Guild lead row', () => expect(copy.raidGuild.features[0]).not.toBe('Everything in Guild'));
  it('example label and link', () => {
    expect(copy.examples.serverRun.label).toBe(key('examples.serverRun.label')[0]);
    expect(copy.examples.serverRun.link).toBe(key('examples.serverRun.link')[0]);
  });
  it('caption.preset and caption.ownGear follow the B5 templates', () => {
    const f = { characterName: 'Kraggor', level: '60', specLabel: 'Fury Warrior', combinations: '4,812', ranAt: '27 Oct 2026' };
    const slots = { ...f, 'ranAt as d MMM yyyy': f.ranAt };
    expect(copy.examples.serverRun.caption.preset(f)).toBe(fill(key('examples.serverRun.caption.preset')[0], slots));
    expect(copy.examples.serverRun.caption.ownGear(f)).toBe(fill(key('examples.serverRun.caption.ownGear')[0], slots));
  });
  ```
  `fill` here (and in Task 3's file, where it is defined once) replaces every occurrence of each `{slot}` (`replaceAll`).
- [ ] **Step 2: Run to fail.** Expected: FAIL.
- [ ] **Step 3: Implement.** The B5 per-raider string `{a40} … {a20} … {a10}` is filled from `perRaiderAmounts`, with the `40-person` word from `RAID_SIZES[0]`.
- [ ] **Step 4: Run to pass**, check, lint, prettier. Fix any `PlanCard`/`GuildRules` compile error from removed keys with the minimal rename only; layout changes belong to Tasks 11 to 16.
- [ ] **Step 5: Commit.** `git commit -m "feat(web): guild-region copy, cost lines and the example caption templates" -- web/src/lib/billing/plans-page-copy.ts web/src/lib/billing/plans-page-copy.test.ts`

### Task 5: The trial-email sentence is defined but not appended

**Files:**
- Modify: `web/src/lib/billing/plans-page-copy.ts`, `web/src/lib/billing/plans-page-copy.test.ts`

**Interfaces:**
- Consumes: `TRIAL_REMINDER_DAYS` (Task 1).
- Produces: `plansPageCopy.trialEmail: string` — B5 `trialEmail`, `We email you {TRIAL_REMINDER_DAYS} days before the trial ends.` filled to `3`. Appended to nothing (B10 ruling 1).

- [ ] **Step 1: Write the failing tests:**
  ```ts
  it('trialEmail is B5 verbatim', () => expect(copy.trialEmail).toBe(fill(key('trialEmail')[0], { TRIAL_REMINDER_DAYS: String(TRIAL_REMINDER_DAYS) })));
  it('trial line and Trial rule make no email promise until the trial email ships', () => {
    expect(copy.trial).not.toContain('email you');
    expect(copy.rules.find((r) => r.term === 'Trial')?.sentence).not.toContain('email you');
  });
  ```
  The trial-email lane flips the second test and appends the sentence; this lane does neither.
- [ ] **Step 2: Run to fail** (`trialEmail` undefined). **Step 3: Implement.** **Step 4: Run to pass**, check, lint, prettier.
- [ ] **Step 5: Commit.** `git commit -m "feat(web): trial email sentence, held until the email ships" -- web/src/lib/billing/plans-page-copy.ts web/src/lib/billing/plans-page-copy.test.ts`

### Task 6: The sim landing copy

**Files:**
- Modify: `web/src/lib/sim/landing-copy.ts` and its test (`web/src/lib/sim/landing-copy.test.ts`; create it if absent)

**Interfaces:**
- Produces: `landingCopy.exampleCaption` `EXAMPLE · A REAL SERVER RUN`; `landingCopy.exampleSpec` and `landingCopy.exampleDps` deleted; `exampleHint` unchanged. The caption sentence is not in `landingCopy`: the card reads `serverRunCaption()` (Task 4).

- [ ] **Step 1: Write the failing tests:**
  ```ts
  it('label is the B5 label', () => expect(landingCopy.exampleCaption).toBe(specKeyCell(SPEC, '### B5. Copy, verbatim', 'landingCopy.exampleCaption')[0]));
  it('no invented figure keys', () => {
    expect(landingCopy).not.toHaveProperty('exampleSpec');
    expect(landingCopy).not.toHaveProperty('exampleDps');
  });
  ```
- [ ] **Step 2: Run to fail.** **Step 3: Implement.** `ExampleResultCard.svelte` stops compiling: replace its two removed reads with `serverRunCaption()` as a one-line holding change; Task 17 does the card.
- [ ] **Step 4: Run to pass** with `src/components/sim/ExampleResultCard.test.ts` (update only its `1,245 DPS` expectation to the caption), check, lint, prettier.
- [ ] **Step 5: Commit.** `git commit -m "feat(web): sim landing example label; drop the invented figure" -- web/src/lib/sim/landing-copy.ts web/src/lib/sim/landing-copy.test.ts web/src/components/sim/ExampleResultCard.svelte web/src/components/sim/ExampleResultCard.test.ts`

### Task 7: The fixture-swap test (B8 check 2)

**Files:**
- Create: `web/src/lib/billing/plans-page-copy.fixtures.test.ts`

**Interfaces:**
- Consumes: `plans-page-copy.ts`, `landing-copy.ts`, `web/src/components/premium/model.ts`, `precision.ts`, `plans.ts`, `premium-examples.json`.

- [ ] **Step 1: Write the test** (it fails if any figure is typed):
  ```ts
  beforeEach(() => vi.resetModules());
  it('every string follows swapped figures', async () => {
    vi.doMock('../sim/precision', async (orig) => ({ ...(await orig()), CAPS: { browser: 333, server: 7777 }, LANE_ITERATION_CEILING: { browser: 11111, server: 222222 } }));
    // The copy module must compute every amount from the imported PLANS / RAID_SIZES values (never
    // through plans.ts's own helpers, which stay bound to the real module under a mock), so swapping
    // the values here is enough to move every string. Tasks 1 and 4 build the lines that way.
    vi.doMock('./plans', async (orig) => {
      const real = await orig<typeof import('./plans')>();
      return { ...real, RAID_SIZES: [30, 15, 5] as const, TRIAL_REMINDER_DAYS: 4,
        PLANS: { ...real.PLANS, raid_guild: { ...real.PLANS.raid_guild, founderYearly: 300 } } };
    });
    vi.doMock('../../data/premium-examples.json', () => ({ default: { serverRun: { simId: 'swap1', characterName: 'Swapname', level: 58, specLabel: 'Arms Warrior', preset: true, combinations: 1234, ranAt: '2027-02-03T12:00:00Z' } } }));
    const { plansPageCopy: c, serverRunCaption, perRaiderLine } = await import('./plans-page-copy');
    expect(c.premium.features[0]).toContain('7,777');
    expect(c.premium.features[0]).toContain('333');
    expect(c.premium.features[2]).toContain('222,222');
    expect(c.trialEmail).toContain('4 days');
    expect(perRaiderLine('yearly', true)).toBe('$10 per raider a year in a 30-person raid, $20 in a 15, $60 in a 5.');
    expect(serverRunCaption()).toContain('level-58 Arms Warrior');
    expect(serverRunCaption()).toContain('1,234');
    expect(serverRunCaption()).toContain('3 Feb 2027');
    const { costLine } = await import('../../components/premium/model');
    expect(costLine('raid_guild', 'yearly', true)).toContain('30-person');
  });
  ```
  The `perRaiderLine` expectation is built from B5's per-raider template with the swapped figures; adjust the `PLANS` field name to whatever `plans.ts` names the founder yearly price.
- [ ] **Step 2: Run.** `npx vitest run src/lib/billing/plans-page-copy.fixtures.test.ts`. Expected: FAIL until Task 8's `costLine` exists (run after Task 8 in the same wave, or mark the `costLine` assertion with the task order in Parallel waves).
- [ ] **Step 3: Fix any string that does not follow** by moving its figure to the import. **Step 4: Run to pass**, check, lint, prettier.
- [ ] **Step 5: Commit.** `git commit -m "test(web): every plans-page figure follows its source (B8 check 2)" -- web/src/lib/billing/plans-page-copy.fixtures.test.ts`

### Task 8: `model.ts` cost lines, trial split and held coach

**Files:**
- Modify: `web/src/components/premium/model.ts`, `web/src/components/premium/model.test.ts`

**Interfaces:**
- Consumes: `perMemberLine`, `perRaiderLine`, `plansPageCopy.trial` (Task 4).
- Produces:
  - `export function costLine(plan: PlanKey, interval: Interval, founder: boolean): string | null` — `perMemberLine` for `guild`, `perRaiderLine` for `raid_guild`, `null` for `premium`.
  - `trialParts(): { lead: string; rest: string }` — `lead` `First 7 days free` (from `GUILD_TRIAL_DAYS`), `rest` the remainder of the new `trial` line.
  - `planBlock('premium').features` is `plansPageCopy.premium.features`.

- [ ] **Step 1: Write the failing tests:**
  ```ts
  it('costLine per plan', () => {
    expect(costLine('premium', 'yearly', true)).toBeNull();
    expect(costLine('guild', 'yearly', true)).toBe(perMemberLine('yearly', true));
    expect(costLine('raid_guild', 'monthly', false)).toBe(perRaiderLine('monthly', false));
  });
  it('trialParts splits the new trial line', () => {
    const { lead, rest } = trialParts();
    expect(lead).toBe(`First ${GUILD_TRIAL_DAYS} days free`);
    expect(lead + rest).toBe(plansPageCopy.trial);
  });
  it('planBlock premium never includes the held coach line', () =>
    expect(planBlock('premium').features).not.toContain(plansPageCopy.premium.featuresHeld.coach));
  ```
- [ ] **Step 2: Run to fail.** **Step 3: Implement.** **Step 4: Run to pass** with Task 7's file, check, lint, prettier.
- [ ] **Step 5: Commit.** `git commit -m "feat(web): plan model cost lines and the new trial split" -- web/src/components/premium/model.ts web/src/components/premium/model.test.ts`

### Task 9: One source for the example image

**Files:**
- Create: `web/src/lib/billing/example-image.ts`, `web/src/lib/billing/example-image.test.ts`

**Interfaces:**
- Consumes: `serverRunExample`, `ServerRunExample` (Task 2); `serverRunCaption` (Task 4).
- Produces:
  - `export const EXAMPLE_IMAGE_WIDTH = 440; export const EXAMPLE_IMAGE_HEIGHT = 330;`
  - `export interface ExampleImage { avif: string; webp: string; width: 440; height: 330; alt: string; loading: 'lazy'; decoding: 'async' }`
  - `export function exampleImage(ex: ServerRunExample = serverRunExample): ExampleImage` — `avif` `/premium/example-server-run.avif`, `webp` `/premium/example-server-run.webp`, `alt` `serverRunCaption(ex)`.

- [ ] **Step 1: Write the failing tests:**
  ```ts
  it('image source, size and loading', () => {
    const img = exampleImage(VALID);
    expect(img).toMatchObject({ avif: '/premium/example-server-run.avif', webp: '/premium/example-server-run.webp', width: 440, height: 330, loading: 'lazy', decoding: 'async' });
  });
  it.each([true, false])('alt equals the caption (preset=%s)', (preset) =>
    expect(exampleImage({ ...VALID, preset }).alt).toBe(serverRunCaption({ ...VALID, preset })));
  ```
- [ ] **Step 2: Run to fail.** **Step 3: Implement.** **Step 4: Run to pass**, check, lint, prettier.
- [ ] **Step 5: Commit.** `git commit -m "feat(web): one image source for the premium example" -- web/src/lib/billing/example-image.ts web/src/lib/billing/example-image.test.ts`

### Task 10: `PremiumExampleCard.astro`

**Files:**
- Create: `web/src/components/premium/PremiumExampleCard.astro`, `web/src/components/premium/PremiumExampleCard.test.ts`

**Interfaces:**
- Consumes: `exampleImage` (Task 9), `plansPageCopy.examples.serverRun` and `serverRunCaption` (Task 4), `exampleHref` (Task 2).
- Produces: `<PremiumExampleCard example?: ServerRunExample />` (defaults to `serverRunExample`), B3.2's example card:
  - `card` recipe, no gold top highlight;
  - `label` 11 px `--muted` `EXAMPLE · A REAL SERVER RUN`;
  - 8 px, a `picture` (AVIF then WebP `source`), `img` with `width="440" height="330"`, `aspect-ratio: 4 / 3`, `width: 100%`, 1 px `--line-soft` border, 4 px radius, `loading="lazy" decoding="async"`, alt the caption, inside a box on `--card-top`;
  - 8 px, the caption, 13 px `--muted`, numbers in mono (`data-testid="premium-example-caption"`);
  - the inline gold link `Open this result` in a 44 px min-height row to `/sim/{simId}` (`data-testid="premium-example-link"`).

- [ ] **Step 1: Write the failing container-render tests** (`createPageContainer` from `web/src/test-support/page-container.ts`):
  ```ts
  it.each([[true], [false]])('renders label, picture, caption, link in order (preset=%s)', async (preset) => {
    const ex = { ...VALID, preset };
    const html = await (await createPageContainer()).renderToString(PremiumExampleCard, { props: { example: ex } });
    const doc = new JSDOM(html).window.document;
    const order = [...doc.querySelectorAll('[data-part]')].map((n) => n.getAttribute('data-part'));
    expect(order).toEqual(['label', 'image', 'caption', 'link']);
    const img = doc.querySelector('img')!;
    expect(img.getAttribute('width')).toBe('440');
    expect(img.getAttribute('height')).toBe('330');
    expect(img.getAttribute('alt')).toBe(serverRunCaption(ex));
    expect(doc.querySelector('[data-testid=premium-example-caption]')!.textContent!.trim()).toBe(serverRunCaption(ex));
    expect(doc.querySelector('[data-testid=premium-example-link]')!.getAttribute('href')).toBe('/sim/abc123');
  });
  it('renders caption and link with the image absent', async () => { /* render, remove <picture>, assert caption + link + the reserved box (aspect-ratio 4 / 3 on the wrapper) remain */ });
  it('has no gold top highlight', async () => { /* assert the card root lacks the gold highlight class the Premium card uses */ });
  ```
- [ ] **Step 2: Run to fail.** **Step 3: Implement.** The reserved box holds the space with `aspect-ratio`, so a failed image keeps the height.
- [ ] **Step 4: Run to pass**, check, lint, prettier.
- [ ] **Step 5: Commit.** `git commit -m "feat(web): the premium example card" -- web/src/components/premium/PremiumExampleCard.astro web/src/components/premium/PremiumExampleCard.test.ts`

### Task 11: The Premium card

**Files:**
- Modify: `web/src/components/premium/PlanCard.astro`
- Create: `web/src/components/premium/PlanCard.test.ts`

**Interfaces:**
- Consumes: `plansPageCopy.premium.{support, features, scope}` (Task 3).
- Produces (B3.2 Premium card): the 40 px slot holds the support line (13 px `--text`, weight 400, never gold; `data-testid="premium-support"`); the six benefits in the tick-list recipe (14 px `--text`, 32 px minimum rows on `--line-soft`, 14 px `--kill` tick), each allowed to wrap; the scope note 12 px above, 13 px `--muted`, no tick (`data-testid="premium-scope"`). Free's slot stays blank. Gold top highlight kept on Premium.

- [ ] **Step 1: Write the failing render tests:**
  ```ts
  it('premium slot holds the support line, not gold', async () => { /* render PlanCard plan='premium'; [data-testid=premium-support] text === copy.premium.support; class list has no text-gold */ });
  it('six benefits then the scope note, no tick on the note', async () => { /* li count 6, texts equal copy.premium.features; scope after the list; scope has no tick svg */ });
  it('free slot stays blank', async () => { /* render plan='free'; slot element empty */ });
  ```
- [ ] **Step 2: Run to fail.** **Step 3: Implement.** **Step 4: Run to pass**, check, lint, prettier.
- [ ] **Step 5: Commit.** `git commit -m "feat(web): Premium card support line, six benefits and scope note" -- web/src/components/premium/PlanCard.astro web/src/components/premium/PlanCard.test.ts`

### Task 12: The guild cards, visibly secondary

**Files:**
- Modify: `web/src/components/premium/PlanCard.astro`, `web/src/components/premium/PriceBlock.astro`, `web/src/components/premium/PlanCard.test.ts`
- Create: `web/src/components/premium/PriceBlock.test.ts`

**Interfaces:**
- Consumes: `costLine` (Task 8), `trialParts` (Task 8), `SECONDARY_BUTTON_FIXED` (`web/src/lib/planner/styles.ts`).
- Produces (B3.3 guild cards):
  - top border `--line`, no gold top highlight;
  - price figure mono 500 at 24 px desktop, 22 px phone (Premium keeps 32 / 28);
  - price block lines in order: price; list-price line while founder; per-month line (yearly) or `Or {yearly}.` (monthly); founder line while founder; the cost line, 13 px `--text`, 18 px line height, in a 54 px box (`data-testid="cost-line"`);
  - the whole block fixed at 152 px from 768 and 144 px on phone, in both intervals and both founder states;
  - button `Subscribe a guild · {price}` in plain `SECONDARY_BUTTON_FIXED` (border `--line-warm-strong`, text `--text`), `w-full justify-center`;
  - trial slot 13 px `--text`, `First 7 days free` in `--strong` 600, minimum 40 px, on a shared subgrid row;
  - five feature lines; no `Everything in Guild` lead row or its `--strong` style.

- [ ] **Step 1: Write the failing render tests:**
  ```ts
  it.each([['yearly', true], ['yearly', false], ['monthly', true], ['monthly', false]] as const)(
    'cost line sits in a 54 px box in all four forms (%s, founder=%s)', async (interval, founder) => {
      /* render PriceBlock plan='raid_guild'; the cell for (interval, founder) has [data-testid=cost-line] with class h-[54px] and text costLine('raid_guild', interval, founder) */
    });
  it('guild price block is fixed 144 px phone, 152 px from 768', async () => { /* root class contains h-[144px] and md:h-[152px] */ });
  it('guild price figure 22/24 px, premium 28/32', async () => { /* class assertions */ });
  it('guild button is neutral', async () => { /* button class contains SECONDARY_BUTTON_FIXED tokens and border-line-warm-strong; no border-gold */ });
  it('guild cards have no gold top highlight', async () => { /* … */ });
  it('trial slot min 40 px with the strong lead', async () => { /* … */ });
  ```
- [ ] **Step 2: Run to fail.** **Step 3: Implement.** **Step 4: Run to pass**, check, lint, prettier.
- [ ] **Step 5: Commit.** `git commit -m "feat(web): guild cards secondary, with fixed price blocks and cost lines" -- web/src/components/premium/PlanCard.astro web/src/components/premium/PriceBlock.astro web/src/components/premium/PlanCard.test.ts web/src/components/premium/PriceBlock.test.ts`

### Task 13: The band's guild link

**Files:**
- Modify: `web/src/pages/premium.astro`
- Create: `web/src/pages/_premium.test.ts`

**Interfaces:**
- Consumes: `plansPageCopy.band.{guildLead, guildLink}` (Task 4).
- Produces (B3.1): one row under the founder row, from 768 only (`hidden md:flex`): 13 px `--text`, 44 px tap row, `Running a guild?`, 6 px, the inline gold link `Guild and Raid guild are further down` to `#guild-leaders` (`--gold`, `--gold-hover` and underline on hover, 2 px gold focus ring at 2 px offset; `data-testid="band-guild-link"`).

- [ ] **Step 1: Write the failing render test:**
  ```ts
  it('band guild link from 768 only', async () => {
    const doc = await renderPremium();
    const link = doc.querySelector('[data-testid=band-guild-link]')!;
    expect(link.getAttribute('href')).toBe('#guild-leaders');
    expect(link.textContent!.trim()).toBe(plansPageCopy.band.guildLink);
    expect(link.closest('[data-band-guild-row]')!.className).toMatch(/\bhidden\b.*\bmd:flex\b|\bmd:flex\b.*\bhidden\b/);
  });
  ```
- [ ] **Step 2: Run to fail.** **Step 3: Implement.** **Step 4: Run to pass**, check, lint, prettier.
- [ ] **Step 5: Commit.** `git commit -m "feat(web): band link to the guild plans" -- web/src/pages/premium.astro web/src/pages/_premium.test.ts`

### Task 14: Region A layout

**Files:**
- Modify: `web/src/pages/premium.astro`, `web/src/pages/_premium.test.ts`

**Interfaces:**
- Consumes: `PlanCard` (Tasks 11, 12), `PremiumExampleCard` (Task 10).
- Produces (B3.2): Region A, no heading. From 1024 `grid-template-columns: repeat(3, minmax(0, 1fr))`, gap 16: Free, Premium, example; Free and Premium one subgrid row set, the example stretching to the row's height. From 768 to 1023 two columns, the example spanning both underneath at max 480 px, left-aligned. Below 768 one column: Free, Premium, example. The 1344 px column centred up to 2000 and beyond; the band full bleed. The four-column grid and the guild-card order rules come out. Free appears once.

- [ ] **Step 1: Write the failing render tests:**
  ```ts
  it('region A holds Free, Premium, example in that order', async () => { /* [data-region=a] children data-plan order ['free','premium','example'] */ });
  it('Free renders once', async () => expect((await renderPremium()).querySelectorAll('[data-plan=free]')).toHaveLength(1));
  it('region A grid classes', async () => { /* lg:grid-cols-3, md:grid-cols-2, example md:col-span-2 md:max-w-[480px] lg:col-span-1 lg:max-w-none */ });
  it('no four-column grid', async () => expect((await renderPremium()).body.innerHTML).not.toMatch(/grid-cols-4/));
  ```
- [ ] **Step 2: Run to fail.** **Step 3: Implement.** **Step 4: Run to pass**, check, lint, prettier.
- [ ] **Step 5: Commit.** `git commit -m "feat(web): plans page region A - Free, Premium and the real example" -- web/src/pages/premium.astro web/src/pages/_premium.test.ts`

### Task 15: Region B layout

**Files:**
- Modify: `web/src/pages/premium.astro`, `web/src/pages/_premium.test.ts`

**Interfaces:**
- Consumes: `plansPageCopy.guildRegion` (Task 4), `PlanCard` guild variants (Task 12), the history promise (§3.1.4 recipe).
- Produces (B3.3): `<section id="guild-leaders">` with `scroll-margin-top: 16px`, 40 px above on desktop and 28 px on phone; h2 `section-title` `For guild leaders`; 8 px, the lede 14 px `--text` max 720 px; 4 px, the members line 13 px `--muted` max 720 px; 16 px, the guild row: from 1024 three equal columns, gap 16 (Guild, Raid guild, the history promise top-aligned with its 2 px `--line` left rule, padding-left 14); 768 to 1023 two columns with the history promise spanning underneath; below 768 one column.

- [ ] **Step 1: Write the failing render tests:**
  ```ts
  it('#guild-leaders holds heading, lede, members, Guild, Raid guild, history promise', async () => {
    const s = (await renderPremium()).getElementById('guild-leaders')!;
    expect(s.querySelector('h2')!.textContent!.trim()).toBe(plansPageCopy.guildRegion.heading);
    expect(s.textContent).toContain(plansPageCopy.guildRegion.lede);
    expect(s.textContent).toContain(plansPageCopy.guildRegion.members);
    expect([...s.querySelectorAll('[data-guild-row] > *')].map((n) => n.getAttribute('data-plan'))).toEqual(['guild', 'raid_guild', 'history']);
  });
  it('scroll margin and gaps', async () => { /* scroll-mt-4, mt-7 md:mt-10 */ });
  it('no loot-council example', async () => expect((await renderPremium()).body.textContent).not.toContain('EXAMPLE · A REAL LOOT COUNCIL'));
  ```
- [ ] **Step 2: Run to fail.** **Step 3: Implement.** **Step 4: Run to pass**, check, lint, prettier.
- [ ] **Step 5: Commit.** `git commit -m "feat(web): plans page region B - for guild leaders" -- web/src/pages/premium.astro web/src/pages/_premium.test.ts`

### Task 16: `GuildRules.astro` moves inside the guild section

**Files:**
- Modify: `web/src/components/premium/GuildRules.astro`, `web/src/pages/premium.astro`
- Create: `web/src/components/premium/GuildRules.test.ts`

**Interfaces:**
- Consumes: `plansPageCopy.howItWorks`, `plansPageCopy.rules` (Task 4).
- Produces: `GuildRules` renders inside `#guild-leaders`, its heading an h3 at the `section-title` size (`How the guild plans work`), rows in `rules` order.

- [ ] **Step 1: Write the failing render test:**
  ```ts
  it('h3, Verified members first, Paying together after Size, no Battle.net', async () => {
    const doc = /* render GuildRules */;
    expect(doc.querySelector('h3')!.textContent!.trim()).toBe(plansPageCopy.howItWorks);
    const terms = [...doc.querySelectorAll('[data-rule-term]')].map((n) => n.textContent!.trim());
    expect(terms[0]).toBe('Verified members');
    expect(terms[terms.indexOf('Size') + 1]).toBe('Paying together');
    expect(doc.body.textContent).not.toContain('Battle.net');
  });
  ```
  Plus a `_premium.test.ts` case: `GuildRules` is a descendant of `#guild-leaders`.
- [ ] **Step 2: Run to fail.** **Step 3: Implement.** **Step 4: Run to pass**, check, lint, prettier.
- [ ] **Step 5: Commit.** `git commit -m "feat(web): how the guild plans work, inside the guild section" -- web/src/components/premium/GuildRules.astro web/src/components/premium/GuildRules.test.ts web/src/pages/premium.astro web/src/pages/_premium.test.ts`

### Task 17: `ExampleResultCard.svelte` shows the real capture

**Files:**
- Modify: `web/src/components/sim/ExampleResultCard.svelte`, `web/src/components/sim/ExampleResultCard.test.ts`

**Interfaces:**
- Consumes: `exampleImage` (Task 9), `serverRunCaption` (Task 4), `landingCopy.exampleCaption`, `landingCopy.exampleHint` (Task 6).
- Produces (B3.4): recipe kept (`border-line bg-card-top rounded-panel border p-4`, gap 8): the label, the same picture element, image files and reserved 4:3 box, the caption, the existing hint. No `Open this result` link. Prop `example?: ServerRunExample` for tests.

- [ ] **Step 1: Write the failing tests:**
  ```ts
  it.each([true, false])('caption and alt follow the data file (preset=%s)', (preset) => {
    const ex = { ...VALID, preset };
    render(ExampleResultCard, { props: { example: ex } });
    expect(screen.getByText(serverRunCaption(ex))).toBeTruthy();
    expect(screen.getByRole('img').getAttribute('alt')).toBe(serverRunCaption(ex));
  });
  it('img carries width and height', () => { /* 440 / 330 */ });
  it('no invented figure and no link', () => {
    render(ExampleResultCard);
    expect(document.body.textContent).not.toContain('1,245 DPS');
    expect(screen.queryByRole('link')).toBeNull();
  });
  it('label is the B5 label', () => { /* EXAMPLE · A REAL SERVER RUN */ });
  ```
- [ ] **Step 2: Run to fail.** **Step 3: Implement.** **Step 4: Run to pass**, check, lint, prettier.
- [ ] **Step 5: Commit.** `git commit -m "feat(web): sim landing shows the real server run" -- web/src/components/sim/ExampleResultCard.svelte web/src/components/sim/ExampleResultCard.test.ts`

### Task 18: The capture script for the example image

**Files:**
- Create: `web/tests/captures/premium-examples/capture.spec.ts`, `web/tests/captures/premium-examples/playwright.capture.config.ts`

**Interfaces:**
- Consumes: `SIM_ID` (env), `API_BASE` (env, default production), `GET /v1/sims/{id}` (public read).
- Produces:
  - Fails fast unless `SIM_ID` is set, the result answers 200 signed out, and the result is a **finished server Top Gear run** (status finished, lane server, kind Top Gear; read the fields from the sim response the result page already uses).
  - Opens `/sim/{id}` at a 440 x 900 CSS px viewport, DPR 2; clips the top 440 x 330 box (character strip, combinations and precision, winning set DPS with error, first swaps) to `web/public/premium/example-server-run.webp` and `.avif` at 880 x 660 (AVIF via `sharp`, quality stepped down until at most 45 KB; fail if it cannot reach it).
  - Writes `web/src/data/premium-examples.json` with all seven fields: `simId`, `characterName`, `level`, `specLabel`, `preset`, `combinations`, `ranAt`, read from the sim response (`preset` true when the run's gear source is the site's preset; the script fails if it cannot tell). Then re-imports the file through `parseServerRunExample` and fails if it throws.
  - The config mirrors `web/tests/captures/discord-bot/playwright.capture.config.ts`.

- [ ] **Step 1: Write the script and config.** No unit test: the controller runs it in Task 19.
- [ ] **Step 2: Run** `npm run check && npm run lint && npx prettier --check tests/captures/premium-examples/`. Expected: PASS.
- [ ] **Step 3: Commit.** `git commit -m "test(web): capture script for the premium example image and data file" -- web/tests/captures/premium-examples/capture.spec.ts web/tests/captures/premium-examples/playwright.capture.config.ts`

### Task 19 (controller step): Controller: run the example sim on the owner's account, save the result id into the data file

The controller does this by hand; it needs the owner's signed-in session (prod user 78, open-ended Premium grant).

- [ ] **Step 1:** On production, signed in as the owner, open the simulator with a **level-60 Fury Warrior in the site's preset gear** (the owner's own warrior, `Obnoxious Yell`, is level 21, so it is not used). Run a **server** Top Gear sim and wait for it to finish. Note the result id.
- [ ] **Step 2:** Confirm `GET /v1/sims/{id}` answers 200 signed out.
- [ ] **Step 3:** Run `cd web && SIM_ID=<id> npx playwright test -c tests/captures/premium-examples/playwright.capture.config.ts`. Expected: the two images (AVIF at most 45 KB) and the data file with `preset: true`, `level: 60`, `specLabel` the Fury Warrior label the result shows.
- [ ] **Step 4:** Run `npx vitest run src/lib/billing/` (the shipped file parses, `isPlaceholder` false). Look at the image beside the result page.
- [ ] **Step 5:** Commit: `git commit -m "feat(web): the premium example is a real server Top Gear run" -- web/public/premium/example-server-run.avif web/public/premium/example-server-run.webp web/src/data/premium-examples.json`.
- [ ] **Ship blocker:** the branch does not merge while the data file holds `simId: "placeholder"`. Task 21's e2e asserts it.

### Task 20: The `premium-value` mock board

**Files:**
- Modify: `design/mocks/gen_launch_pricing.py`

**Interfaces:**
- Produces: a `premium-value` board at 1024, 1440 and 2000, laid out as B3.5's wireframes (band with the guild link; Region A Free, Premium, example; Region B heading, lede, members line, Guild, Raid guild, history promise; How the guild plans work). Copy is imported from the same strings B5 lists (typed once in the generator's copy block, matched to B5). The comparison board for capture row 22.

- [ ] **Step 1: Add the board.** **Step 2: Run** `python3 design/mocks/gen_launch_pricing.py` and open the three outputs. Expected: the three PNG/HTML files under the generator's output folder.
- [ ] **Step 3: Commit.** `git commit -m "docs(design): premium-value mock board" -- design/mocks/gen_launch_pricing.py <generated board paths>`

### Task 21 (final): e2e spec updates and capture rows 15 to 22

**Files:**
- Modify: `web/tests/e2e/premium.spec.ts`, the `/sim` landing e2e (`web/tests/e2e/sim.spec.ts` or the spec that covers `ExampleResultCard`; `rg -l ExampleResultCard\|example-result web/tests/e2e`), `web/tests/captures/launch-pricing/capture.spec.ts`, `web/lighthouserc.json` (confirm `/premium.html` and `/sim.html` are collected)

**Interfaces:**
- Consumes: `serverRunExample`, `isPlaceholder`, `serverRunCaption`, `exampleHref`; `plansPageCopy`.

- [ ] **Step 1: Write the e2e cases** in `premium.spec.ts` (B8 checks 4 to 14):
  - `guild price blocks hold 152/144 px across interval and founder`: at 360, 768, 1024 and 1440, with `page.clock.install` on both sides of `FOUNDER_WINDOW_CLOSES`, every card's height and every guild price block (144 px below 768, 152 px from 768) and `cost-line` box (54 px, `scrollHeight <= clientHeight`) are unchanged when the interval switches.
  - `first screen at 1280 x 800`: Free's and Premium's buttons end above y 520; the band link is above y 300; activating it puts `For guild leaders` at the top of the viewport, and both guild buttons and trial lines end above its bottom.
  - `guild leaders section`: `#guild-leaders` contains Guild, Raid guild, the history promise and How the guild plans work; `Verified members` is the first rule and equals the copy; no `Battle.net`; `Paying together` follows `Size`.
  - `forbidden words`: no text contains `coach`, `chip in`, `retention`, `compare more than`, or `!`; no `email you` within the trial slot or the Trial rule.
  - `held coach line not rendered`, `no loot-council example`, `Free card once at every width`, axe clean, every link and button at least 44 px tall at 360, exactly one island (the layout's `CharacterSelector`) — run with and without `data-premium` set, and the "buttons above y 520" check passes in both.
  - `example caption reads the flag`: the caption equals `serverRunCaption(serverRunExample)`, which is the preset template when `serverRunExample.preset` is true and the own-gear template otherwise; `isPlaceholder(serverRunExample)` is false (ship blocker).
  - `Open this result answers 200 signed out`: `request.get(API + '/v1/sims/' + simId)` is 200 with no cookie (run against the stubbed API in CI; the controller checks production in Task 19).
- [ ] **Step 2: `/sim` landing e2e:** the real card shows the label and the caption from the data file, `1,245 DPS` appears nowhere, and the img has `width="440"` and `height="330"`.
- [ ] **Step 3: Capture rows 15 to 22** in `web/tests/captures/launch-pricing/capture.spec.ts` at the B9 widths. Row 21 routes the image request to abort (`page.route('**/premium/example-server-run.*', (r) => r.abort())`) and captures the reserved box, alt, caption and link. Row 22 is the full page side by side with Task 20's `premium-value` board at 1024, 1440 and 2000.
- [ ] **Step 4: The task runs** `npm run check && npm run lint && npx prettier --check <the touched specs>` only. The controller then runs `npm run build`, the full `npx vitest run`, the full `npx playwright test`, the lhci shard (`/premium` and `/sim`: performance and accessibility at least 0.95, LCP at most 2400 ms, CLS at most 0.05), and the captures, at branch level. ux-designer and wow-player review rows 15 to 22 on the rendered page, starting from row 22 (memory: side-by-side with the mock).
- [ ] **Step 5: Commit.** `git commit -m "test(web): premium value e2e and capture rows 15 to 22" -- web/tests/e2e/premium.spec.ts <the /sim e2e path> web/tests/captures/launch-pricing/capture.spec.ts web/lighthouserc.json`

---

## Parallel waves

Tasks in a wave touch disjoint files.

| Wave | Tasks |
|---|---|
| 0 | 0 (spec) |
| 1 | 1, 2, 18, 20 |
| 2 | 3 (plans-page-copy, spec-table) |
| 3 | 4 |
| 4 | 5, 6, 8, 9 |
| 5 | 7, 10, 11 |
| 6 | 12, 13, 17 |
| 7 | 14 |
| 8 | 15 |
| 9 | 16 |
| controller | 19 (after 18; before 21) |
| last | 21 |

- Tasks 3, 4 and 5 all edit `plans-page-copy.ts` and its test, so they run in sequence.
- Tasks 11 and 12 both edit `PlanCard.astro`; Tasks 13 to 16 all edit `premium.astro`. Each set runs in sequence.
- Task 6 needs `serverRunCaption` (Task 4) for its one-line holding change to `ExampleResultCard.svelte`, which Task 17 replaces.
- Task 7's `costLine` assertion needs Task 8, so Task 7 runs after Task 8.

## Conflicts ruled

- **Preset example vs B3.2's "owner's own Fury Warrior".** The controller ruled (2026-10-27) a level-60 Fury Warrior in the site's preset gear from the owner's account, because the owner's warrior is level 21. Task 0 writes the ruling into §B3.2, §B5, §B8 check 7 and §B10 ruling 4 before any copy test reads it.
- **`perRaiderLine` / `perMemberLine` placement.** The task list puts them in `plans.ts`; B5 keys them as `price.perRaider` / `price.perMember` in the copy module. The plan keeps the arithmetic in `plans.ts` (`perRaiderAmounts`, `perMemberAmount`) and the sentences in `plans-page-copy.ts`, re-exported under the task list's names, so no sentence lives in `plans.ts`.
- **Caption ownership.** The reader picks the template by `preset` (`exampleCaption(ex, templates)`); the templates live in the copy module. Strings stay in one place and the reader stays testable with injected templates.
- **Spec-table reader.** B5's table has a `Where | Key | String` shape; `specCell` reads the second column. Task 3 adds `specKeyCell` rather than reshaping the spec table.
- **`model.ts` path.** It lives at `web/src/components/premium/model.ts`, not under `lib/billing`.
