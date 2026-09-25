# Home v2: the character is the page

**Date:** 2026-09-24
**Status:** APPROVED by the owner ("I still don't love the layout of the site, specifically the
landing page" → the character-hub direction, "use your judgement" on the dates panel). One
lane: `home-v2`, `web/**` only. Supersedes the layout half of
`2026-09-23-landing-page-design.md`; that spec's honesty, voice and Lighthouse rules still bind.

## 0. What is wrong today (the owner's screenshot, 2026-09-24)

The hero is four stacked rows of text (eyebrow, headline, a dark strip with a 40 px avatar and
four links, chips) inside a sky band with empty space below them, while the "Right now" dates
card on the right is bigger and better structured than the hero, so the eye lands on launch
dates rather than on the visitor's own character. Under the hero, four equal-weight panels of
different heights hold link lists (nine class tiles; twenty-one spec pills; two sentences; two
sentences) and no live number. About 100 px of dead space separates hero from panels. The
page reads as four boxes of links under a slogan.

## 1. Binding constraints

1. `design/DESIGN-SYSTEM.md` in full (Stone and Gold, Cinzel display at 22 px max, reference
   voice, 44 px targets, the sky band as the one reveal, the panel and card treatments).
2. The states model (`2026-09-22-account-and-island-states-design.md` §1): every island has a
   skeleton sized to its ready state, retry, empty state, one `.reveal`; nothing moves.
3. Islands read the session through `createQueryState` and never call `fetchMe()` or
   `fetchMeOnce()` from an effect that re-runs on the answer (the 2026-09-23 loop).
4. Character identity is drawn only through `components/character/*` (`CharacterPortrait`,
   `CharacterIdentity`, `CharacterRow`), spec `2026-09-23-character-identity-design.md`. No
   private copies.
5. Honesty: every figure shown is real, dated and sourced, or absent. No invented numbers,
   no fake screenshots, no "10,000 players". Empty states are real empty states.
6. Vocabulary: addon in game, companion on the desktop, "from Battle.net".
7. Lighthouse (`web/lighthouserc.json`): the home page is in the strict bucket (performance
   0.95, LCP 2000 ms, TBT 100 ms, CLS 0.05). The LCP element is server-rendered text or a
   static image, never an island's output; the character render is `loading="lazy"` and never
   the LCP. `npm run lhci` before the final report, every number quoted. Local Lighthouse on a
   loaded Mac reads high on every page; compare against `main` under the same load before
   concluding anything.
8. Not touched: `Header.astro`, `Footer.astro`, `components/ui/**`, `lib/ui/**`, `lib/data/**`,
   `api/**`. (Header changes are a separate decision.)
9. Lane rules: `.superpowers/journeys/lane-common-web.md`; commit trailer
   `Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>` then
   `Claude-Session: https://claude.ai/code/session_01DAHb7A9Ytd6qpZu2UujXgY`.

## 2. Information architecture, top to bottom

Section names are the design's, not headings to print. Vertical rhythm: 48 px between
sections at `lg`, 32 px on a phone; the hero band's height is its content plus the band's own
padding, never a fixed viewport fraction.

### 2.1 Hero (sky band)

**Signed in** (`HomeAccountPanel.svelte`, grown into the whole hero; it already owns the
`/v1/me` read and the grid-overlay swap with the server-rendered signed-out hero):

- Left, at `lg`: the main character's render (`render_url`) at up to 320 px tall, bottom-aligned
  to the band, `loading="lazy"`, `alt=""`, hidden below `lg`. When there is no render (addon-only
  import), `CharacterPortrait` at `lg` takes the slot beside the name instead and the band is
  simply shorter.
- Right of it: `CharacterIdentity` at `lg` with the full descriptor (this is the one place the
  display face is used for a character name); under it the guild line when one exists.
- Actions, one row: two buttons in the design system's primary and secondary treatments,
  **Sim {name}** (to `/sim` with the armory source for the character, the same href
  `sim/LandingState.svelte`'s `hrefFor` builds) and **Plan talents** (to `/planner`, which
  restores the pointer); two text links, **Logs** (`/logs`) and **Your characters**
  (`/account`). When the character has no build, the first button becomes **Get the build**
  and links to `/account#characters` (the paste and refresh live there); nothing pretends to
  sim.
- Under the actions: the alts as chips (`CharacterPortrait` sm + name + level), the existing
  `home-character-chip` behaviour (a chip switches the current character), `+N more`.
- The latest rating figure, when one exists, sits at the end of the identity block as today
  (`home-hero-rating`); nothing shown until it resolves.
- The eyebrow "World of Warcraft: Forever" stays as the band's first line; the pitch headline
  does not render when signed in.

**Signed out** (server-rendered, unchanged copy): eyebrow, the headline "Your character,
planned, simmed, logged and ranked.", the one sentence, the Battle.net button and the muted
"or paste an addon export" link. Right of it at `lg`: a sample render, a static image under
`web/public/images/home/` chosen by the lane from the site's existing art or omitted if none
exists that is ours to use (never a copyrighted screenshot). The signed-out block keeps
`data-session-hide` and the same grid cell as the signed-in hero so nothing moves.

### 2.2 The timeline (replaces the "Right now" card)

One row directly under the sky band, full width, the design system's panel treatment at its
quietest: the five dates from `src/data/dates.json` as a horizontal timeline (date in mono,
label, note), the next upcoming date carrying the gold dot, past dates muted. On a phone it
becomes a two-column list. The "Updated" stamp stays at the row's end. After launch the same
row shows "This week" (out of scope; the component takes rows, not a meaning).

### 2.3 Next steps (signed in) / The four things (signed out)

One row of four cards, equal height (`grid-rows` from the tallest; `items-stretch`), each the
state panel treatment with the gold dot, a title, one line, one live element and one action.
Every live element is a real number from an endpoint that exists, or the card's empty state.

Signed in, each card is keyed to the main character:

| Card | Live element | Action |
|---|---|---|
| Simulator | The latest saved sim's DPS for this character from `listMySims` (kind and figure, dated), else "No sim yet." | Run / Sim {name} (same href as the hero) |
| Planner | Points spent from the character's build (`decodeFS1` on the build code from `fetchSimInput`, "24 of 51 points", else "No build yet.") | Continue planning / Get the build |
| Logs | The latest report of the visitor's from `listMyReports(1)` (title, date), else "No logs yet." | Open / Upload a log |
| Rankings | The character's rating figure (already fetched for the hero; share the value), else "Not rated yet." | Rankings for {class} → `/rankings?class=<slug>` |

Signed out, the same four cards carry their product sentence from today's copy and their
empty states, with the links to `/planner`, `/sim`, `/logs`, `/rankings`. The class tile row
and the spec pill wall are removed from the home page (both live on `/planner` and
`/sim/specs`).

### 2.4 Around the site

Two compact tables side by side at `lg`: **Recent reports** (`RecentReports` compact, as today)
and **Top guilds** (`HomeTopGuilds`, as today), each five rows with a "See all" link. Their
skeletons are five rows tall.

### 2.5 Reference

The quieter band: the search box with the `/` hint and "Every fact dated and sourced.", the
four reference tiles (Classes, Guides, Zones, Dungeons from `data/tools.json`'s `community`
entries), the "Popular" links row. The Guides tile links to `/guides`; when the guides lane
lands, this section's Guides tile is where its class guides surface, no other change here.

### 2.6 Your guild, the addon and the companion, What changed, Still unknown, Help build this

As today (`HomeGuildLink`, the two small cards, the changelog feed, the unknowns, the subscribe
box), in this order, with the new rhythm.

## 3. What leaves

The "Right now" card (becomes 2.2), the class tile row and the spec pill wall (2.3), the
`HomeProductPanel` 2×2 grid, the dark hero strip.

## 4. Tests and evidence

- SSR render tests for the timeline row (next date carries the dot; past muted) and the four
  cards' signed-out states.
- `home-panel.spec.ts` updated: signed-in hero shows the render when present and the portrait
  when not, the button labels by build state, chips switch, the returning-visitor snapshot
  test; a new test that the four signed-in cards render their live element with stubbed
  `sims`, `reports` and `rating` routes and their empty states without.
- `home.spec.ts`, `subscribe.spec.ts`, `handoffs-home.spec.ts` pass; the `/` shortcut still
  focuses the moved search.
- `npm run lhci` with every number quoted; screenshots of the signed-out and signed-in home at
  360 and 1280 saved under `.superpowers/shots/` in the worktree, with `/v1/me` stubbed to a
  fixture that has a `render_url`.
- The final report names every copy string it added.
