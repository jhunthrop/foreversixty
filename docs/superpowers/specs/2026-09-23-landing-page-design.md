# The landing page is the product, not the wiki

**Date:** 2026-09-23
**Status:** APPROVED by the owner ("We need to redesign the landing page to align with the
new features of the site. We started as just a static wiki but that's not our first tier
function anymore."). One lane: `landing`, `web/**` only.

## 0. Why

The home page today: eyebrow "fan reference, every fact dated and sourced", a search box
as the first control, "Popular" wiki links, the launch-dates panel, then tool cards, "What
changed", recent reports, class tiles. It reads as a wiki with tools attached. The product
is now: sign in with Battle.net, and your character is planned, simmed, logged and ranked
here, with your guild. The reference is the second tier: it supports those, it is not the
front door.

## 1. Binding constraints

1. `design/DESIGN-SYSTEM.md` in full: Stone and Gold, Cinzel display, reference voice ("state
   the thing and stop", no pitch, no superlatives, no emoji), the sky band as the one page
   reveal, the panel and card treatments, the 44 px targets. This is the same site; it
   changes what it says first, not what it looks like.
2. The vocabulary rule: addon in game, companion on the desktop, "from Battle.net" for
   Blizzard-sourced data.
3. The states model (`2026-09-22-account-and-island-states-design.md`): every island has a
   skeleton sized to its ready state, a retry, an empty state, one `.reveal`.
4. Lighthouse (`web/lighthouserc.json`): the home page is in the strict bucket (performance
   0.95, LCP 1800 ms, TBT 100 ms, CLS 0.05). Islands hydrate `client:visible` or later; the
   page's LCP element must be server-rendered text or a static image, never an island's
   output. Measure with `npm run lhci` before the final report and quote every number.
5. Honesty: every figure shown is real, dated and sourced, or absent. Nothing invented,
   no "10,000 players" copy, no fake screenshots. Where a live number needs an endpoint that
   does not exist, show the panel's real empty or "sample" state as the design system's
   state panel does today.
6. `Header.astro`, `Footer.astro`, `components/ui/**`, `lib/ui/**` are not touched.
7. Lane rules (`.superpowers/journeys/lane-common-web.md`): sonnet only, never opus; commit
   rules; scoped checks; no broad `pkill`; stop servers.

## 2. Information architecture

Top to bottom, signed out. Section names are the design's, not headings to print.

1. **Header** (unchanged, session-aware).
2. **Hero (sky band).** Eyebrow: `World of Warcraft: Forever`. Headline (Cinzel, the one
   place on the site above 22 px is still not allowed: 22 px max per the design system, so
   the headline is a two-line 22 px display line): "Your character, planned, simmed, logged
   and ranked." One sentence under it: "Sign in with Battle.net and your characters arrive
   with their gear, talents and guild." The Battle.net button (`next=/account?signed_in=1`),
   beside it the muted alternative "or paste an addon export" linking to `/addon#paste`.
   Right column at `lg`: the **Right now** dates panel stays (pre-launch, the dates are the
   most useful fact on the site); after launch it becomes a "This week" panel, out of scope.
   The search box moves out of the hero to the Reference section (4); the `/` shortcut still
   focuses it from anywhere on the page.
3. **The four things, with real data.** A 2×2 grid at `lg` (one column on a phone) of
   product panels, each the design system's state panel with the gold dot, a title, one
   sentence of what it does, one live element, and the link:
   - **Planner**: "Every talent tree for 1.60, every race and class, shared by link." Live
     element: a class row of the nine class tiles at 36 px (server-rendered, links to
     `/planner?class=<slug>`).
   - **Simulator**: "Your DPS, upgrades from any loot table, stat weights." Live element: a
     server-rendered list of the damage specs the simulator offers (from the same data the
     sim's spec grid uses), as small pills linking to `/sim?spec=<slug>`.
   - **Logs**: "Upload a night or log live with the companion; every fight, every parse."
     Live element: the existing `RecentReports` island (`client:visible`, its skeleton and
     empty state as they are).
   - **Rankings**: "Guild progression and character parses, per boss." Live element: the top
     five guilds by progression from `GET /v1/rankings/guilds` if the endpoint exists
     (check `web/src/lib/rankings/`), as a `client:visible` island with a skeleton of five
     rows; otherwise the panel's real empty state "No kills ranked yet." with the link.
4. **Reference.** A quieter band: the search box (moved here, same component, same `/`
   hint) with the sentence "Every fact dated and sourced." and the four reference tiles
   Classes, Guides, Zones, Dungeons (from `data/tools.json`'s `community` entries), plus
   the "Popular" links row under the search.
5. **Your guild.** One panel: "Claim your guild with a Battle.net sign-in: a home page,
   roster, progression and officer tools." with the link to `/guilds` (or wherever the guild
   entry point lives today; the lane finds it). Includes the existing `HomeGuildLink`
   behaviour for a signed-in member.
6. **The addon and the companion**, as one row of two small cards: "The addon adds bags,
   bank and professions to your Battle.net build and shows the next talent point in game."
   → `/addon`; "The companion logs live from your desktop and sends your exports." →
   `/logs#companion`. Reference voice: what each does, one line each.
7. **What changed** and **Still unknown** stay, below, as they are; **Help build this**
   with the subscribe box stays last.

Signed in, sections 2 and 5 change and the rest stays:

2. **Hero, signed in.** The sky band holds the hub summary instead of the pitch: the main
   character (the same `hero-character` rule the account page uses): avatar or class
   square, name in class colour, descriptor, the four actions `Open in planner`,
   `Open in simulator`, `Logs`, `Your characters`, and the latest rating figure when one
   exists. The dates panel stays on the right. This replaces `HomeAccountPanel`'s strip:
   that component grows into the signed-in hero (it already owns the `/v1/me` read and the
   grid-overlay swap; the server-rendered signed-out hero stays in place under it until it
   mounts, same cell, same height, so nothing moves).
5. **Your guild, signed in**: the panel shows the member's guild name, progression and a link
   to its home (the data `HomeGuildLink` already reads).

## 3. What leaves

- The eyebrow "fan reference · every fact dated and sourced" as the page's first line (it
  moves to the Reference section as its sentence).
- The tool-card grid as a section (its content is section 3).
- The class tiles as a standalone section (they are the planner panel's live element).

## 4. Tests and evidence

- SSR render tests for the new panels' server-rendered content; the existing
  `home-panel.spec.ts` and `home.spec.ts` updated (the signed-in hero, the moved search, the
  `/` shortcut still focusing it, no client JavaScript on the content pages that never had
  it).
- `npm run lhci` with every number quoted; screenshots of the signed-out and signed-in
  home at 360 and 1280 saved under `.superpowers/shots/` in the worktree, taken with
  `/v1/me` and the rankings route stubbed.
- The final report names every copy string it added, so the owner can read the page's
  words in one place.
