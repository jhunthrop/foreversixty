# One product: current character, hand-offs, grouped nav, logs discoverability

**Date:** 2026-09-21
**Status:** approved ("ship it all"). Source: four journey reviews (fresh leveller, raider,
guild officer, information architect), 2026-09-20. The reviews agree: four or five good
tools in one skin, not one product. Nothing holds a "current character"; reference pages and
character/guild/account pages are dead ends; the addon page (the front door for the whole
character flow) is in the footer; no public report is reachable.

This spec is binding for four lanes that run in parallel. Each lane owns a set of files and
must not edit another lane's files; where two lanes need the same behaviour, the owner is
named here.

## 0. Rules for every lane

- Design system: `design/DESIGN-SYSTEM.md`. No emoji, no marketing buttons, secondary
  buttons only, restrained motion, 44px hit targets on phone, dark only.
- Honest copy: never promise what is not built; say what a page is for and who it is for.
  No exclamation marks.
- No layout shift: anything that appears after hydration reserves its space or sits below
  existing content. Lighthouse budgets in `web/lighthouserc.json` must hold.
- Signed-out first: every feature works without an account; an account only adds.
- localStorage is a convenience, never the source of truth: wrap every read and write in
  try/catch, render correctly without it, never read it during SSR/prerender.
- Tests: unit tests for every pure function, component tests where the repo has them, e2e
  for each new hand-off (desktop and mobile projects). `FOREVER_DATA=fixture npm run sync`
  before web tests. Node 22.12 via nvm.

## 1. The current character (Lane B owns)

**Pointer.** `{ source: SourceKind, ref: string }` is already the sim's `TabSource`
(`web/src/lib/sim/tabs.ts`). Promote it to a shared module `web/src/lib/current-character.ts`:

```ts
export interface CurrentCharacter {
  source: 'addon' | 'build' | 'fight' | 'armory' | 'code';
  /** For 'addon' and 'code' the FS1 string itself; otherwise the existing ref format. */
  ref: string;
  /** What the chip shows: "Simfury · Fury Warrior". Derived once at load, never parsed back. */
  label: string;
  classSlug: string;
  savedAt: string; // ISO
}
export function readCurrent(storage?: Storage): CurrentCharacter | null;
export function writeCurrent(value: CurrentCharacter, storage?: Storage): void;
export function clearCurrent(storage?: Storage): void;
export function plannerHrefFor(current: CurrentCharacter): string;
export function simHrefFor(current: CurrentCharacter, tab?: SimTabId): string;
```

Storage key `fs.currentCharacter`, one JSON value, capped at 16 KB (an FS1 with a full bank
is several KB; over the cap it is not written). Precedence on every tool page: URL
(`?code=`, `?source=&ref=`, `?req=`) first, then the stored pointer, then nothing. A page
that loads a character from its URL or a paste writes the pointer. "Forget" clears it.

**Chip.** `web/src/components/CurrentCharacterChip.svelte`, mounted by `/planner`, every
`/sim*` page, and (by Lane A) `/account`, `/character/*`, `/addon`. It shows the label in
class colour and up to three links: Open in planner, Open in simulator, Copy addon code
(when the planner has a build), plus a Forget control. With nothing loaded it renders
nothing on pages that have their own paste box, and one line ("No character loaded. Paste
an addon export in the planner or the simulator.") elsewhere. Fixed height; it never moves
content when it resolves.

**Sim tabs carry the character.** All six `/sim*` tabs load the current character with no
re-paste: the tools island (`/sim/gear`, `/sim/drops`, `/sim/talents`, `/sim/weights`)
learns to bootstrap from `?code=` and from the stored pointer (the known gap left by
persona round 1: "Tools island can't bootstrap ?code="). The tab strip's links carry the
character whenever the URL can hold it.

**Return visit.** `/sim` and `/planner` opened bare restore the stored character and say so
("Restored your last character. Forget"). The planner restores only when its URL is bare
and its own build is empty.

Lane B also owns, inside the sim and planner components: a "Get the addon" one-line link in
every paste box (`SourceSwitcher`, planner `ImportBox`, tools import) pointing at `/addon`;
"Open in planner" on the Quick Sim result; "Plan it" on each Droptimizer and Top Gear
upgrade row (opens the planner with that item equipped); a one-line framing on `/sim` for
visitors below 60 ("The simulator models level 60 characters. Below 60, plan your build in
the planner and come back."); when a planner build with fewer than 51 points is sent to the
sim, the sim says it is simming it as a level 60 with those talents rather than silently
relabelling it; and Share in the planner asks before it writes (a confirm step naming what
becomes public, offering the unsaved link and the addon code as the no-write alternatives).

## 2. Hand-offs outside the tools (Lane A owns)

Lane A never edits files under `web/src/components/sim/`, `web/src/components/planner/`,
`web/src/lib/sim/`, `web/src/lib/planner/`, `Header.astro`, `Footer.astro`, `logs.astro`,
`rankings*`. It consumes Lane B's `current-character.ts` API by the signatures above (if
Lane B has not landed, Lane A builds against the signatures and its links use the
URL forms directly: `/planner?code=<FS1>` and `/sim?code=<FS1>`).

- Homepage tools grid (`web/src/data/tools.json`, `index.astro`): add Simulator and The
  addon cards; order Planner, Simulator, Logs, Rankings, The addon, then reference.
- `/addon`: add a paste box that takes an export and offers Open in planner / Open in
  simulator; add "what it looks like" copy only if true (no screenshots exist yet: say the
  UI is in beta testing, do not fake one).
- `/account` Characters and `/character/<key>`: per character, Open in simulator and Open
  in planner links where the site has an addon export for it (`/v1/sims/input` answers the
  export for the signed-in owner); otherwise say what is needed ("Log in with the addon
  once to make this character simmable").
- Report pages (`web/src/components/report/`): next to each combatant's existing planner
  link, a "Sim" link for that combatant (`?source=fight&ref=<report>:<fight>` plus the
  combatant: extend the ref as `<report>:<fight>:<guid>`; Lane B's `fromLoggedFight` reads
  the optional third part and falls back to today's first-dps rule when absent. Lane A
  writes the link; Lane B owns `sources.ts` and adds the parsing. Coordinate through this
  sentence: the third part is the combatant GUID exactly as the report's roster has it.)
- Reference pages: every `/guides/<class>` links to the planner for that class; every
  `/dungeons/<slug>` links its zone and, where the dungeon has loot in
  `data/builds/<build>/loot.json`, links "See drops for your character" to `/sim/drops`
  with the instance preselected (`?instance=<slug>`; Lane B reads it). `/zones` says which
  zones are covered and that the rest are coming, rather than silently omitting them.

## 3. Navigation (Lane C owns `Header.astro`, `Footer.astro`, `Base.astro` nav wiring)

Primary nav, desktop:

```
Planner   Simulator   Logs   Rankings   Reference ▾   The addon        [Sign in | name]  Discord
Reference ▾ : Classes, Guides, Zones, Dungeons
```

Decision: tools stay first-level (they are the product and one click matters); the four
reference pages fold into one disclosure; The addon is promoted from the footer; Changelog
moves to the footer. A full "My character / Raiding / Reference" regrouping is deferred
until the current-character chip has shipped and been used; this spec records that as the
next step, not this one.

- The disclosure is a native `<details>` styled to the design system, no client JS on
  content pages, closed by default, fixed header height (no layout shift when it opens:
  the panel is absolutely positioned). Keyboard and screen-reader operable; `aria-current`
  on the active item and on the Reference summary when a child is active.
- Phone (below `md`): the nav row still scrolls horizontally; add a right-edge fade that
  shows while more is scrollable (CSS only: a gradient mask on the scroll container), and
  order the items so the four tools come first. Reference opens as a full-width panel
  under the header.
- "Sign in" (or the signed-in name) is in the header on every page: today several pages
  pass no `session` prop to `Base.astro`. Make the session nav the default and remove the
  opt-in. It must not shift layout: reserve its width.
- A `/` keyboard shortcut focuses search where the page has the search box; otherwise it
  goes to `/search`. Tiny inline script, no dependency.

## 4. Logs discoverability (Lane D owns API `api/`, `logs.astro`, `rankings*`, logs lib)

- **API:** `GET /v1/reports/recent` — public, unauthenticated, rate-limited like the other
  public reads, returns the newest `visibility = 'public'` and `status = 'complete'`
  reports: id, title or zone, created_at, fight_count, kill_count, guild name when the
  report has one, never the owner's identity when the owner has `anonymize` set. Page size
  10, cursor by `created_at,id`. Migration only if an index is needed for the query (add
  one on `(visibility, status, created_at desc)` if `EXPLAIN` shows a seq scan).
- **Web:** a "Recent public reports" panel on `/logs` (above "Your reports") and a compact
  one on the homepage; empty state is honest ("No public reports yet. The first raid logs
  land in December; dungeon logs are welcome now.").
- **Rankings:** the Characters board needs an encounter picker. Add `GET /v1/encounters`
  (or extend an existing endpoint if one already lists encounters with ranked parses) and a
  picker on `/rankings`; with no encounter data yet, the board says so and offers the
  Guilds board, instead of demanding an encounter the visitor cannot choose.
- **Framing:** `/logs` gains one line for visitors who are not raiding yet ("Logs are for
  group content at any level: a dungeon run logs the same way a raid does.").
- **Planner Share "gear invalid":** a journey reviewer's Share of the Simfury test
  character was refused 400 "gear invalid". Find out whether `POST /v1/builds` validates
  item ids against a list that lacks ids the planner itself offers (the Forever re-itemised
  ids), and fix the validator or the planner so the two agree. Add a regression test with
  the failing payload.

## 5. Out of scope, recorded

A guild home and roster; the full "My character / Raiding / Reference" regrouping; armory
import; in-game addon screenshots; a fix in the engine fork for the rage-potion nil
dereference (tracked separately; the web guard shipped 2026-09-20).
