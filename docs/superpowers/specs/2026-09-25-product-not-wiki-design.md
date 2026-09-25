# The product, not the wiki: one character, five doors, no reference

**Date:** 2026-09-25
**Status:** APPROVED in direction by the owner (the whole-site UX review of 2026-09-25 and
"we can drop the reference all together"; "write the spec for all of this"). Five lanes in
three phases, listed in section 8. Sections 1 and 2 bind every lane.

## 0. Why

The site has three identities stacked on top of each other: a wiki (Reference, Classes,
Zones, Dungeons, search, "every fact dated and sourced"), a tool suite (Planner, Simulator,
Logs, Rankings, each with its own sub-navigation and none carrying the visitor's character to
the next), and, since home v2, a character hub. A visitor who lands in one never sees the
others. The wiki will not beat Wowhead and tells visitors the site is something it is not.

The direction: the site is organised around one object, the visitor's character. Every page
opens already pointed at it. The reference leaves; the spec guides stay because they end in a
running build, not in prose; the changelog stays because it is the site's honesty.

## 1. Binding constraints (every lane)

1. `design/DESIGN-SYSTEM.md` in full (Stone and Gold, Cinzel display at 22 px max, reference
   voice: state the thing and stop; no pitch, no superlatives, no emoji; 44 px targets; the
   sky band is the one page reveal; no primary marketing button).
2. The states model (`2026-09-22-account-and-island-states-design.md` §1): every island has a
   skeleton sized to its ready state, a retry, an empty state, one `.reveal`; nothing moves.
3. Islands read the session through `createQueryState` (`lib/data/query.svelte.ts`) and never
   call `fetchMe()`/`fetchMeOnce()` from an effect that re-runs on the answer.
4. Character identity only through `components/character/*`.
5. Honesty: every figure real, dated, or absent. No invented numbers, no fake screenshots.
6. Vocabulary: addon in game, companion on the desktop, "from Battle.net". One character
   descriptor everywhere: `characterDescriptor` (race, class, level, realm and ruleset).
7. Lighthouse budgets in `web/lighthouserc.json` hold; `npm run lhci` before every final
   report with numbers quoted, compared against `main` under the same load.
8. Old URLs never 404: every removed route gets a permanent redirect in `web/public/_redirects`
   (Cloudflare Pages), and `web/tests/e2e/redirects.spec.ts` (new, lane A) proves each one.
9. Lane rules: `.superpowers/journeys/lane-common-web.md`; commit trailer
   `Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>` then
   `Claude-Session: https://claude.ai/code/session_01DAHb7A9Ytd6qpZu2UujXgY`. Never end a
   turn to wait; ledger every ruling.

## 2. The information architecture

Five doors in the header, in this order: **Plan** (`/planner`), **Sim** (`/sim`), **Logs**
(`/logs`), **Rankings** (`/rankings`), **Guides** (`/guides`). A sixth, quieter item at the
end of the row: **Get set up** (`/setup`, the addon and companion in one page, section 3.4).
Premium moves to the footer and the account menu. Discord stays as the icon button.

The account is an avatar chip (the main character's portrait, else the Battletag) that opens a
menu: the character list (switch current), Your account, Devices, Plan, Sign out. Signed out
it is the Battle.net button, as today.

The current character is the spine (section 4): a slim bar under the header on every tool
and guide page, and every tool opens already pointed at it.

## 3. Lane A: the cut, the header and the redirects (`web/**`)

### 3.1 Removed

- Routes and their pages: `/classes`, `/zones`, `/zones/*`, `/dungeons`, `/dungeons/*`,
  `/search`, and the three wiki entries of the `pages` collection (`skyborne`,
  `everything-we-know`, `editions`). The `zones` and `dungeons` content collections and
  `src/content/pages/{skyborne,everything-we-know,editions}.md` are deleted with their
  tests. The `pages` collection itself stays for `about`, `sources`, `roadmap`, `privacy`,
  `refunds`, `terms`.
- The search box component, the `/` shortcut, the search index build step, and the "Popular"
  links (`data/tools.json`'s `community` entries and the popular row).
- The home page's Reference band and the "Still unknown" panel (`data/unknowns.json`).
- The Reference dropdown and the `isReferenceCurrent` machinery in `lib/nav.ts`/`Header.astro`.
- The `classes` link from every class guide landing page that pointed at `/classes`; the
  guides index carries the class list now.

### 3.2 Redirects (`web/public/_redirects`, 301)

| From | To |
|---|---|
| `/classes` | `/guides` |
| `/zones`, `/zones/*`, `/dungeons`, `/dungeons/*` | `/guides` |
| `/search` (any query) | `/guides` |
| `/skyborne`, `/everything-we-know`, `/editions` | `/changelog` |
| `/addon` | `/setup` |

### 3.3 Header

`lib/nav.ts` becomes the five doors plus `SETUP_NAV_ITEM`; `PREMIUM_NAV_ITEM` leaves the
header (footer keeps its Premium link; the account menu gains "Plan"). On a phone the nav is
a wrapping two-row grid of the six items, not a horizontal scroll that cuts off (the CLS
reservation the old comment describes moves to a fixed two-row height). `SessionNav.svelte`
becomes the avatar chip and menu (`components/account/AccountMenu.svelte`): portrait
(`CharacterPortrait` sm of the main, else the initial), the Battletag, a native
`<details>`-based menu with the items in section 2, 44 px rows, closes on outside click and
Escape. Test ids: `account-menu`, `account-menu-character-<key>`, `account-menu-signout`.

### 3.4 `/setup` (replaces `/addon`; `pages/addon.astro` becomes `pages/setup.astro`)

One page, three numbered steps, each a panel: **1. Sign in with Battle.net** (imports your
characters; one line saying which realm types Blizzard serves today, from
`accountPageCopy.noBattlenetDataForRealm`'s family of copy), **2. The addon** (install links,
`/fs export`, the paste box, the in-game window paragraph, no screenshot claim until one
exists), **3. The companion** (downloads, pairing, `/combatlog`). The paste box keeps its
`#paste` anchor: every "Paste export" link on the site points at `/setup#paste`. The logs
page's companion column links here instead of repeating the download panel (the logs page
keeps the upload form and its own one-paragraph companion pointer).

### 3.5 Home, 404, footer

Home: the Reference band, Popular row and Still unknown panel are removed; the addon and
companion row becomes one "Get set up" card linking `/setup`. The 404 page offers Guides,
Plan, Sim and Logs. Footer: About, Sources, Changelog, Contribute, Premium, Terms, Privacy,
Refunds (unchanged), plus "Get set up".

### 3.6 Bugs fixed in this lane

The logs page's "Loading your reports" for a signed-out visitor (render the section only
when signed in); the account page's empty "Your ratings" panel (an `EmptyState`: "Nothing
rated yet. Ratings appear after your first ranked fight."); the simulator page's 600 px
floor under "Your sims" (the shell's `min-h` applies to the pre-hydration shell only; the
island's own height governs after).

## 4. Lane B: the character spine (`web/**`, after lane A merges)

### 4.1 The bar

`components/character/CurrentCharacterBar.svelte` grows into the spine and renders on
`/planner`, `/sim*`, `/logs`, `/rankings*`, `/guides/*`, `/character/*`: one row under the
header, `CharacterIdentity` at `md` with the realm descriptor, then four links (Plan, Sim,
Logs, Rankings for {class}) each carrying the character, and a "Switch" control that opens
the same character list the account menu shows. It reads the current pointer
(`lib/current-character.ts`) and, signed in, the session; the pointer wins, the main is the
fallback. No pointer and signed out: the bar renders one line, "Sign in with Battle.net or
paste an export to point the site at your character", with the two links. The `.chip-slot`
pre-paint rule (2026-09-25) extends to the bar: reserved only while a pointer or a session
hint exists.

### 4.2 Every door opens on the character

- `/planner`: opens on the current character's class and build (today's restore), and when
  there is none, on the main's class.
- `/sim`: opens on the current character's sim (the armory source) when it has a build; the
  landing list only when it does not.
- `/logs`: "Your reports" first for a signed-in visitor, filtered to reports the character
  appears in, then the upload form.
- `/rankings`: pre-filtered to the character's class and ruleset, with the character's own
  row pinned at the top when it is ranked.
- `/guides/<class>/<spec>`: the class landing links to the character's class first; a spec
  guide's "Load this build" (lane C) loads it for the current character.

### 4.3 Vocabulary

One descriptor string everywhere (constraint 6). The ruleset labels (`rulesetLabel`) and
the region are rendered by `CharacterIdentity` only; every other hand-written "PvP · US"
is removed.

## 5. Lane C: guides that end in a build (`web/**`, after lane A)

Every spec guide's **Talents and builds** section embeds the planner's tree component
(`components/planner/TreeGrid.svelte` in read-only mode, one column per tree) with the
guide's recommended build lit, from a `build:` frontmatter field holding an FS1 code the
guides lane's writers produce from the guide's own point list (cross-checked against
`talents/<class>.json`). Under the trees, one button **Load this build** (`/planner?code=`
for a signed-out visitor; for a signed-in visitor it also writes the pointer so Sim and Logs
follow) and one **Sim this build** (`/sim?code=`). **Stat priority** renders as an ordered
row of pills with the reasoning beneath; **Races** renders as a row of race portraits with
the recommended ones marked. Gear stays prose until raid loot is itemised. Each guide's
"may change" caveat moves from under the title to a quiet line above Sources.

## 6. Lane D: empty states, caveats and the after-sim sentence (`web/**`, after lane A)

- Every empty state says what will appear and offers the one action that gets there:
  rankings with no data ("No ranked fights yet for this filter. Rankings fill in as reports
  are uploaded." + Upload a log), the simulator's "Your sims", the account's reports and
  ratings, the guild page's roster.
- Caveats move after the thing: the simulator's two intro sentences become one line under
  the character list ("Damage specs at level 60; healing and tanking specs are not simulated
  yet."); the planner's build line moves under the trees; the guide caveat as in lane C.
- The report page's three tab rows collapse to two: the mode row (Analyze, Compare,
  Rankings, Mechanics) stays; the view row (Tables, Timelines, Events, Queries) becomes a
  segmented control inside the Analyze mode; the thirteen table tabs become a select on a
  phone and a wrapping pill row at `lg`.
- After a sim: the results card's first line is the one best next action from the data the
  run already produced ("Upgrade: Bracers of X from Blackfathom Deeps, +14 DPS" from the
  Droptimizer's top row when it exists; "Run Top Gear to find your next upgrade" when not).
- The planner's summary bar labels its figures (Level, Points left, Split, Spent, DPS) and
  drops "Level 9" for a bare build (level shown only when a character is loaded).

## 7. Phase 3 (specs to follow, not lanes yet)

- **Progress over time**: a rating and item-level line per character from ranked fights
  (`/v1/characters/{key}`'s history), on the character page and the home hero.
- **First-run checklist**: after the first sign-in, three ticks on the account page and the
  home hero (characters imported, a build in, a first log or sim) that tick themselves.
- **Guild activity**: "your guild raided last night" on the home page from the reports feed
  filtered to the guild, with the night's best parses.
- **Share cards**: Open Graph images for builds, fights and characters through the existing
  `/og/` route.
- **Tooltips**: every stat, talent and enchant name gets the same tooltip everywhere.

## 8. Lanes, order and ownership

| Lane | Sections | Owns | After |
|---|---|---|---|
| A `cut` | 3 | `lib/nav.ts`, `Header.astro`, `SessionNav.svelte` → `AccountMenu.svelte`, `pages/{classes,zones,dungeons,search,addon}*`, `content/{zones,dungeons}`, the three wiki pages, `public/_redirects`, `pages/index.astro`'s reference band, `404.astro`, `Footer.astro`, `logs.astro`/`Account.svelte`/`SimView.svelte` for 3.6 only | — |
| B `spine` | 4 | `CurrentCharacterBar.svelte`, `CurrentCharacterChip.svelte`, `lib/current-character*.ts`, the pages' bar mounts, `Rankings.svelte` prefilter, `logs` "Your reports" order, `sim` bootstrap | A |
| C `guides-build` | 5 | `content/guides/**` (frontmatter `build:`), `pages/guides/**`, `components/guides/*` (new), `TreeGrid.svelte` read-only mode | A |
| D `states` | 6 | copy modules, empty states, `ReportView.svelte` tabs, `SummaryBar.svelte`, `SimResults` first line | A |

B, C and D run in parallel after A merges; they own disjoint files. Each lane: plan, ledger,
task reviews, whole e2e suite once, Lighthouse quoted, final report under 25 lines.

## 9. Tests and evidence

- Lane A: `redirects.spec.ts` (every row of 3.2 answers 301 to its target); header specs
  updated for the five doors and the menu (open, switch character, sign out, Escape closes);
  the phone nav shows all six items without horizontal scroll; the classes-page overflow is
  gone with the page; `home.spec.ts` for the removed bands; the three bug fixes each have a
  spec.
- Lane B: a spec per door proving it opens on the current character, and one proving the
  bar's signed-out line.
- Lane C: an SSR test that every spec guide has a `build:` that decodes and matches its
  class; a spec that Load this build lands in the planner with the points lit.
- Lane D: SSR tests for each new empty state; the report tabs spec; the after-sim sentence
  with a stubbed Droptimizer row.
