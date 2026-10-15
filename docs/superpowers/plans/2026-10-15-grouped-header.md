# Grouped Header Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Replace the eight-door bar with four doors (Tools, Logs, Guild, Guides), each opening a hover panel that lists its pages, with the same four groups in the phone Menu.

**Architecture:** `lib/nav.ts` becomes the one source of the groups; `Header.astro` server-renders every panel (hidden) from it, the Guides grid from `classes.json` plus the spec catalogue at build time; one plain inline script in the header handles hover, click, keyboard and touch (no island); `nav.css` keeps every breakpoint; the selector island tells the header the session's guild through one window event.

**Tech Stack:** Astro 5, Svelte 5, TypeScript, vitest, Playwright. Work in `web/`.

**Spec:** `design/specs/2026-10-15-grouped-header.md` (mock: https://claude.ai/artifact/QMUea9sAGREktD7bPZhFmB).

## Global Constraints

- Doors, in order: `Tools` (`/planner`), `Logs` (`/logs`), `Guild` (`/guild`), `Guides` (`/guides`). Rows, hrefs and descriptions verbatim from spec section 2; foots `All guides` and `Guild settings`; the Guild first row is `Your guild` until the session fills it.
- Hover opens after 120ms, closes 180ms after the pointer leaves door and panel; one panel open at a time; Escape closes and refocuses the door; the bar never changes height or shifts the doors when a panel opens.
- No new island and no framework code in the header: the panel script is a processed `<script>` in `Header.astro` next to the existing menu script, plain DOM, no global `keydown` beyond what the header already registers.
- Tokens only (`--color-*` from `tokens.css`); class colours from `classes.json`; crests through `classCrestSrc`. No new images.
- The rendered bar must not scroll sideways at 1024, 1279, 1280, 1439, 1440 and 2000 (nav-selector spec), and the lookup's "/" hint shows from 1280 again.
- Format with Prettier, lint clean, `astro check` clean. Run the whole web vitest suite and the full Playwright suite locally before reporting (the header is shared layout: `feedback-shared-layout-full-suite`).
- Never commit generated data; never touch `addon/`, `data/`, `api/`.

## Review Focus

1. At 1280, signed in with "Obnoxious Yell", hover each door: the panel opens under it, nothing in the bar moves, and `document.documentElement.scrollWidth === clientWidth` with the Guides panel open.
2. Keyboard only: Tab to Logs, ArrowDown opens and focuses "Upload a log", ArrowDown twice reaches Rankings, Escape returns focus to the Logs door with the panel closed.
3. A signed-out visitor on `/tiers`: the Logs door is `aria-current="page"`, its panel's Tier list row links `/tiers`, and the Guild panel links `/guild` everywhere.

---

### Task 1: The nav groups in `lib/nav.ts`

**Files:**
- Modify: `web/src/lib/nav.ts`
- Modify: `web/src/lib/nav.test.ts`
- Modify: `web/src/pages/_setup.test.ts` (only if it breaks)

**Interfaces:**
- Produces: `interface NavRow { label: string; href: string; description: string }`, `interface NavGroup { label: string; href: string; rows: readonly NavRow[]; foot?: NavItem }`, `NAV_GROUPS: readonly NavGroup[]` (four groups, spec section 2, Guides with rows `[]` because the header renders its grid from `classes.json`), `GUILD_GROUP_INDEX = 2`, `isNavGroupCurrent(group, path)`. `SETUP_NAV_ITEM` and `isNavItemCurrent` stay (GuildClaim.svelte and `_setup.test.ts` import them). `PRIMARY_NAV_ITEMS` and `TRAILING_NAV_ITEMS` are deleted with their tests; grep for other importers first and move them to `NAV_GROUPS`.

- [ ] **Step 1: Write the failing tests** in `nav.test.ts`: the four group labels and hrefs in order; the Tools rows are Planner `/planner`, BiS `/bis`, Simulator `/sim`, Get set up `/setup`; the Logs rows are Upload a log `/logs`, Your reports `/reports`, Rankings `/rankings`, Tier list `/tiers`; the Guild rows are Your guild `/guild`, Roster `/guild#roster`, Loot council `/guild#loot`, Guild logs `/guild#raids` with foot Guild settings `/guild#settings`; Guides has no rows and foot All guides `/guides`; every description is non-empty; `isNavGroupCurrent(logs, '/tiers/healer')` is true, `isNavGroupCurrent(tools, '/logs')` is false, `isNavGroupCurrent(guild, '/guild/us/pvp/olympus-xxvii')` is true.
- [ ] **Step 2: Run** `npx vitest run src/lib/nav.test.ts` → FAIL.
- [ ] **Step 3: Implement** the types, `NAV_GROUPS` (descriptions verbatim from the spec), `isNavGroupCurrent` (the group href or any row href by `isNavItemCurrent`, hashes stripped), delete the two old arrays.
- [ ] **Step 4: Run** the file → PASS; `npx vitest run src/pages/_setup.test.ts` → PASS.
- [ ] **Step 5: Commit** `feat(web): the nav groups behind the four doors`.

### Task 2: The header renders the doors, panels and phone groups

**Files:**
- Modify: `web/src/components/Header.astro`
- Create: `web/src/components/HeaderPanel.astro` (one group's panel: rows or the Guides grid, the foot)
- Modify: `web/src/styles/nav.css`
- Modify: `web/src/components/Header.test.ts`

**Interfaces:**
- Consumes: `NAV_GROUPS`, `isNavGroupCurrent`, `classCrestSrc` (`lib/class-crest.ts`), `classes.json`, `readSpecCatalog` (`lib/bis/load.ts`, build time, as `pages/guides/index.astro` does; spec names in `tree_index` order).
- Produces: markup contract for Task 3 and 4: each door `a.site-door[data-door="{tools|logs|guild|guides}"]` with `aria-haspopup="true" aria-expanded="false" aria-controls="site-panel-{key}"`, a `button.site-door-toggle` (chevron, `aria-label="Open the {Door} list"`) beside the label, the panel `div#site-panel-{key}.site-panel[role="menu"][hidden]` with `a.site-panel-row[role="menuitem"]` rows carrying `data-testid="site-panel-row-{slug}"`; the guild rows additionally `data-guild-tab="{overview|roster|loot|raids|settings}"` and the first row's title `span[data-guild-name]`. `data-testid="primary-nav"` stays on the doors nav; `data-testid="menu-button"` stays.

- [ ] **Step 1: Write the failing tests** in `Header.test.ts` (render with `experimental_AstroContainer` as the file already does): four doors in order with those hrefs; each panel present, `hidden`, and holding its rows with the spec hrefs; the Guides panel holds nine `a[href^="/guides/"]` in `classes.json` order, each with an `img[src="/icons/hd/crests/{slug}.webp"]` and the class name; `/tiers/healer` marks the Logs door `aria-current="page"`; the old "renders the doors and Get set up" test is rewritten to the four doors; the phone Menu holds the four group headings in order.
- [ ] **Step 2: Run** `npx vitest run src/components/Header.test.ts` → FAIL.
- [ ] **Step 3: Implement.** `Header.astro`: replace the doors loop with the groups loop (label link + toggle button + `<HeaderPanel group=… />`). `HeaderPanel.astro`: rows for a group with rows; for `guides`, the grid built from `classes.json` and `readSpecCatalog()` (import `readSpecCatalog` from `../lib/bis/load`; this runs at build only, as the guides index does). The phone Menu: the same groups as headed sections (`.site-menu-group`, `h2.site-menu-heading`) reusing the row markup, Guides as crest tiles. `nav.css`: door layout (`.site-door` 44px tall, chevron, open and current states, the 2px rule as `::after`), `.site-panel` (spec section 3: position, widths, raised, line, radius 6, shadow, rows 44px, hover `#101624`), `.site-panel-grid` (3 columns), the phone sections (headings 11px uppercase muted, two-column grid, Guides three columns), and restore the lookup's "/" hint from 1280 (`@media (max-width: 1279px)` on `.site-search-kbd` and the button padding, reverting 2ec19219's 1439 rule) with the 1280 tier's gaps back to the nav-selector spec's 18/20/12 now that four doors fit. `prefers-reduced-motion` kills the chevron transition.
- [ ] **Step 4: Run** the Header tests → PASS; `npm run check`, `npm run lint`, `npm run format:check` → clean.
- [ ] **Step 5: Commit** `feat(web): four doors with server-rendered panels and phone groups`.

### Task 3: Hover, click, keyboard, touch, and the guild fill

**Files:**
- Modify: `web/src/components/Header.astro` (the processed `<script>`)
- Create: `web/src/lib/nav-panels-dom.ts` (pure DOM wiring, testable under jsdom)
- Create: `web/src/lib/nav-panels-dom.test.ts`
- Modify: `web/src/lib/character-selector/config.ts` (add `SESSION_GUILD_EVENT = 'fs:session-guild'`)
- Modify: `web/src/components/nav/CharacterSelector.svelte` (dispatch it once the session resolves)

**Interfaces:**
- Produces: `wireNavPanels(header: HTMLElement, opts?: { openDelay?: number; closeDelay?: number; now?: () => number })` returning `{ open(key), close(), destroy() }`; `applySessionGuild(header, guild: { name: string; href: string } | null)` rewriting the guild rows' hrefs to `${href}#${tab}` (overview: no hash) and the `[data-guild-name]` text; the selector dispatches `window.dispatchEvent(new CustomEvent(SESSION_GUILD_EVENT, { detail }))` with `detail = { name, href: guildHref(region, ruleset, name) }` for the first verified guild in `Me.guilds`, or `null`, each time its session query settles.
- Consumes: Task 2's markup contract.

- [ ] **Step 1: Write the failing jsdom tests** in `nav-panels-dom.test.ts` with fake timers: pointerenter on a door opens its panel after 120ms (`hidden` false, `aria-expanded` true), pointerleave closes after 180ms, entering a second door closes the first; a toggle click opens without navigation and a second closes; ArrowDown on a focused door opens and focuses the first row, ArrowDown walks, Escape closes and focuses the door; a click outside closes; `applySessionGuild` rewrites four hrefs and the name, and `null` restores `Your guild` and `/guild` hrefs.
- [ ] **Step 2: Run** `npx vitest run src/lib/nav-panels-dom.test.ts` → FAIL.
- [ ] **Step 3: Implement** `nav-panels-dom.ts` (timers per door, one open key, `matchMedia('(hover: none)')` makes the first tap on a door open instead of navigate, `Tab` out of the last row closes) and wire it in `Header.astro`'s processed script next to the search wiring; the selector dispatches `SESSION_GUILD_EVENT` from the `$effect` that already reads its session state; the header listens and calls `applySessionGuild`.
- [ ] **Step 4: Run** the new tests and `src/lib/character-selector` tests → PASS; lint, check, format.
- [ ] **Step 5: Commit** `feat(web): hover panels with keyboard and touch parity, guild rows follow the session`.

### Task 4: End-to-end coverage and the old specs

**Files:**
- Create: `web/tests/e2e/nav-panels.spec.ts`
- Modify: `web/tests/e2e/nav.spec.ts`, `web/tests/e2e/nav-selector.spec.ts` (`DOORS`), `web/tests/e2e/tiers.spec.ts` (the "Tier List sits between Rankings and Guides" test becomes "Tier list is a Logs panel row")
- Check: `web/tests/e2e/layout.spec.ts`, `footer-links.spec.ts`, `handoffs-home.spec.ts`, `home.spec.ts` (no change expected)

- [ ] **Step 1: Write `nav-panels.spec.ts`** (desktop project, 1280 and 1440; one phone test): hover Tools opens its panel with the four rows in order and `aria-expanded="true"`, the bar's bounding box is unchanged, `scrollWidth === clientWidth`; moving to Logs closes Tools and opens Logs; mouse out of the header closes within 300ms; the toggle button opens without navigation (URL unchanged); keyboard flow from Review Focus 2; `hasTouch` context: a tap on Guild opens the panel and a second tap navigates to `/guild`; the Guides panel shows nine crests and `/guides/warrior` first; signed in with a verified guild fixture (route `**/v1/me` as nav-selector.spec does, with a `guilds` entry), the Guild panel's first row names the guild and Roster links `guildHref(...) + '#roster'`; on the phone (390) the Menu lists the headings Tools, Logs, Guild, Guides in order with 44px rows.
- [ ] **Step 2: Update** the three old specs to the four doors; keep every width measurement in `nav-selector.spec.ts`.
- [ ] **Step 3: Run** `E2E_SKIP_BUILD` unset (one local build): `npx playwright test tests/e2e/nav-panels.spec.ts tests/e2e/nav.spec.ts tests/e2e/nav-selector.spec.ts tests/e2e/tiers.spec.ts` → PASS; then the full `npx playwright test` and the full `npx vitest run` → PASS (use Node 22.12.0 via nvm, as `web/.node-version` pins).
- [ ] **Step 4: Commit** `test(web): the four doors, their panels and the phone groups`.

### Task 5: Spec cross-references

**Files:**
- Modify: `design/specs/2026-10-09-nav-character-selector.md` (one line under section 2.2: superseded by `2026-10-15-grouped-header.md`)
- Modify: `web/src/lib/nav.ts` header comment (names the new spec)

- [ ] **Step 1:** Add the two references. **Step 2: Commit** `docs: the nav spec points at the grouped header`.
