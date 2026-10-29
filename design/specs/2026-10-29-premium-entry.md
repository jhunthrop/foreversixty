# Premium in the flow: the nav tag, the plan row, the contextual links (experience spec)

Author: ux-designer, 2026-10-29. Authority: `docs/tenets.md` (9, 10, 11, 13, 14), `design/specs/2026-10-28-polish.md` §1 (direction B header, account panel), `design/specs/2026-10-27-launch-pricing.md` (the /premium page).

Owner direction, verbatim: “make the premium changes if it doesn't turn off the personas”.

Mocks: `design/mocks/2026-10-29-premium-entry/`. The generator `gen_premium_entry.py` imports the polish boards' `theme.py` and `parts.py`. The renderer is `render.mjs`. Re-render from the repo root: run `python3 design/mocks/2026-10-29-premium-entry/gen_premium_entry.py`, then `node design/mocks/2026-10-29-premium-entry/render.mjs`.

## 0. Decisions

| # | Decision | Reference |
|---|---|---|
| P1 | Premium is not a fifth door. It is one gold `pill-site` tag, `Premium`, after `Get set up` at the right end of `#primary-nav`. It has no panel and no chevron. | Warcraft Logs (`Subscribe` tag at the bar's end) and Raider.IO (`Premium` tag beside the nav): a small tag, never a menu group. |
| P2 | The tag is hidden for anyone with Premium from any source: own plan, seat, grant or fallback (`premiumSource(me, now) !== 'none'`). Signed-out visitors and free accounts see it. | Warcraft Logs drops its `Subscribe` tag for subscribers. |
| P3 | A free account's `Plan` row in the account panel shows a gold `Upgrade` word before its chevron. The row stays one link. Paid, seat, grant and fallback rows are unchanged. | Battle.net launcher account menu: the plan line carries its one action. |
| P4 | Every contextual link lands on its own line of the Premium card, through an anchor, and that line is highlighted. No contextual link lands on the page top. | Wowhead comment links: the linked item is scrolled into view and tinted. |
| P5 | A contextual link appears only where the player has just met the feature: a coach finding, the server run, a sim's cap, an officer's plan choice. Nothing else on the site gains a Premium mention. | Details! and WeakAuras ask nothing until you reach the feature. |
| P6 | Guild members are never shown a guild plan. The officer gates keep linking to their own guild's billing, because only that link carries the guild. | Members never pay (launch-pricing §B3.3). |

## 1. The question and the answer in view

| Viewer | Question | Answer with no scroll |
|---|---|---|
| Signed out or free, 1280 to 2000 | “What do I get if I pay, and is the free site complete?” | The `Premium` tag at the right end of the doors. One click lands on `/premium`, whose band says `Everything a player needs is free, and the free list never shrinks.` |
| Free, 360 | Same | The Menu sheet's last row is the `Premium` tag. The account sheet's `Plan` row reads `Free` with `Upgrade`. |
| Free, on a report, the sim or guild billing | “What would Premium do for this?” | The link beside the feature opens `/premium` with that feature's line highlighted at one third of the viewport. |
| Premium from any source | Nothing to ask | No tag, no `Upgrade`, no contextual link. The plan row shows the plan as today. |

## 2. Hard limits (the anti-nag rules)

1. No modal, banner, toast, interstitial, countdown or dismissible promo anywhere, ever.
2. No lock icon on any free feature.
3. At most one Premium link per region.
4. No Premium link above a page's answer. On a report the coach teaser sits where it sits today.
5. No gold-filled button for Premium outside `/premium` and `/premium/checkout`. Contextual links are outlined, a tag, or an inline link.
6. The words are `Premium` and `Upgrade`. Never `Go Premium`, `Unlock`, `Pro`, `Limited time` or an exclamation mark.
7. Free stays free. Nothing that is free today gains a Premium mention.

## 3. The nav tag (desktop, from 1150)

Board: `renders/premium-entry-header-1440.png`, `premium-entry-header-2000.png`.

**Markup.** In `Header.astro`, after the `Get set up` link and inside `<nav id="primary-nav">`: an `<a class="site-door site-door-premium" data-door="premium" href="/premium">`. It holds one `<span class="pill pill-site">Premium</span>`. The item is a new `PREMIUM_NAV_ITEM: NavItem = { label: 'Premium', href: '/premium' }` in `lib/nav.ts`, beside `SETUP_NAV_ITEM`. `aria-current="page"` comes from `isNavItemCurrent` (so `/premium/checkout` marks it too).

**Box.** The link is the door's full 40 px height, `align-self: center`, padding 0 4 px. The tag inside is the shared `.pill.pill-site`: 22 px tall, 0 8 px padding, `--radius-pill`, 11 px 700, `.08em`, uppercase, `--color-gold` on `rgba(229,185,85,.14)`, 1 px `rgba(229,185,85,.35)`. No dot: the glowing dot means “you have it” in the account head (polish §1.3), and this tag means “see the plans”. The gap before it is the doors' gap (22 px at 1440, 20 px from 1280, 13 px below 1100).

**Position.** The doors stay one centred group between the wordmark and the right cluster (`justify-content: space-between`, today). The tag is the group's last child. At 2000 everything sits in the 1344 column (polish §1.3).

**States.**

| State | Look |
|---|---|
| Default | as Box |
| Hover (fine pointer) | tag ground `rgba(229,185,85,.22)`, border `rgba(229,185,85,.6)`, text `--color-gold-hover`, 120 ms |
| Focus-visible | 2 px `--color-gold` outline on the link, 2 px offset, radius `--radius-control` |
| Active | tag ground `rgba(229,185,85,.30)` |
| Current (`aria-current="page"`) | tag border `--color-gold`, ground `rgba(229,185,85,.22)` |
| Disabled, loading, empty, error | none: it is a static link in the build-time header |

Keyboard: one Tab stop after `Get set up`, before the search button. Enter follows the link. The door-panel wiring ignores it (`data-door="premium"` has no toggle and no panel).

**Width.** From 1024 to 1149 the tag is `display: none` in the bar. The 1024 bar holds the wordmark, five doors and a full-name account control with no slack. The implementer measures 1150 signed in as `Lightningbolltz` (15 characters). If the bar overflows, raise the breakpoint to the first width that fits and write it in `nav.css` as one named custom property. Premium stays reachable there through the account panel's `Plan` row and the footer's `Premium` link.

## 4. Hiding it for Premium accounts, with no layout shift

The tag is in the static header of every page. Its visibility is decided before first paint and never toggled by `display` after paint.

1. **Pre-paint.** `Base.astro`'s head script reads `localStorage['fs.premiumSource']` (`PREMIUM_SOURCE_KEY`, values `own`, `seat`, `grant`) on every page and sets `html[data-premium]`. The same script already sets `html[data-server-sims]` from the cached `/v1/me` when signed in. Move `/premium`'s own copy of the `data-premium` script into Base and delete it there (one source; PlanCard's CSS keeps working). Then Base sets `html[data-premium-nav="hide"]` when either attribute is present. CSS: `html[data-premium-nav='hide'] .site-door-premium { display: none }`.
2. **After `/v1/me`.** The header wiring reads the session the selector publishes (the `publishSessionGuild` pattern in `session-guild-dom.ts`; add the plan to what is published). If the account has Premium and the tag is showing, set `data-premium-nav="hold"`: `visibility: hidden`, the box kept. If the tag was hidden and the account is now free, it stays hidden for this page view. The next page's pre-paint gets it right, because `writePremiumHint` has already written or cleared the hint.
3. Never key the tag on `html[data-premium]` directly: `writePremiumHint` sets that attribute after paint, and a `display` change then would shift the doors.

CLS on the bar stays 0. The Lighthouse budget (`cumulative-layout-shift` 0.05, `web/lighthouserc.json`) is unchanged.

## 5. The phone Menu sheet (below 1024)

Board: `renders/premium-entry-phone.png`.

The Menu keeps search, the four doors and `Get set up` (polish §1.3). The `Premium` link is its last row, after `Get set up`. Row: 52 px minimum, 18 px side padding, the `.pill.pill-site` tag `Premium` at the left, a 16 px chevron-right in `--color-muted` at the right. The whole row is the link. Hover and active: `--row-hover` ground. Focus-visible: 2 px gold ring, inset 2 px. The same pre-paint rule hides the row. It is the same element as the bar's tag; `nav.css` restyles it below 1024.

## 6. The account panel's Plan row

Board: `renders/premium-entry-account-1440.png` (free beside Premium), and the phone board's account sheet.

| Account | Sub | Trailing | Href |
|---|---|---|---|
| Free (`premiumSource === 'none'`) | `Free` | `Upgrade` in `--color-gold` 13 px 600, then the 16 px chevron in `--color-gold` | `/premium#premium` |
| Own plan | `Premium · $50/yr` over `Renews 26 Oct 2027` (today) | `Founder` tag when founder, then the muted chevron (today) | `/account#plan` (today) |
| Seat, grant, fallback | `accountPlanLine(me, now).line` (today) | muted chevron (today) | `/account#plan` (today) |

The row stays one `<a>` with one Tab stop (`SettingsRow`, polish §4.2). `Upgrade` is text inside it, not a second control. Its accessible name reads `Plan, Free, Upgrade` in that order. States are the row's (polish §1.4): hover `--row-hover`; focus-visible 2 px gold ring inset 2 px; active `rgba(229,185,85,.10)`. Hover also turns `Upgrade` to `--color-gold-hover`. The signed-out panel gains nothing: signing in is its job.

Code: `PlanRow` in `account-panel.ts` gains `upgrade: boolean`, true only for `source === 'none'`. `selectorCopy.account` gains `upgrade: 'Upgrade'`. `freePlanHref` becomes `/premium#premium`.

## 7. Contextual links

Board: `renders/premium-entry-upsell-1440.png`.

Every existing `/premium` link in a feature gets its anchor. Two change their look. None changes its visibility rule except the sim's, which is unchanged in who sees it.

| Where | Shown to | Wording (verbatim) | Look | Href |
|---|---|---|---|---|
| Report coach teaser, `CoachTeaser.svelte` | free viewer with at least one finding (today) | `Read the full review with Premium` (today) | outlined gold button (today) | `/premium#coach` |
| Reports list coach strip, `CoachRowStrip.svelte` | free viewer (today) | the free row lines in `coachCopy` (today) | gold underlined link (today) | `/premium#coach` |
| Sim run control, `RunControl.svelte` | signed in, free (today's `Get premium` slot); signed out sees nothing (today) | `Run on our servers`, then the tag `Premium` | the payer's `Run on our servers` button recipe (`border-line-warm`, `text-nav`, `label`, 44 px, 36 px from md) as a link, with a 4 px gap and the `.pill.pill-site` tag after the label | `/premium#server-sims` |
| Sim premium-required message | after a refused server run (today) | `See plans` (today) | underlined (today) | `/premium#server-sims` |
| Bulk sim cap note, `BulkRunBar.svelte` | at the browser cap (today) | `Premium lifts the limit to 5,000 combinations and any precision.` `See plans` (today) | underlined (today) | `/premium#server-sims` |
| Character History, `HistoryMount.svelte` | owner without the entitlement (today) | today's words | today's look | `/premium#history` |
| Discord notifications, `DiscordPasteForm`, `DiscordStatusLine`, `DiscordSummary` | free (today) | `See plans` (today) | today's look | `/premium#notifications` |
| Guild billing chooser, `GuildPlanChooser.svelte` | an officer choosing a plan | `Compare Guild and Raid guild` | `ActionLink` (gold 14 px 600, 44 px tall, trailing chevron), one line under the tier cards and above the subscribe button | `/premium#guild-leaders` |

The sim link replaces `Get premium`. Delete `simCopy.getPremium`. Its accessible name is `Run on our servers, Premium`. States: hover border `rgba(229,185,85,.45)` and text `--color-strong`; focus-visible 2 px gold outline, 2 px offset; active ground `rgba(229,185,85,.10)`; while a browser run is in progress it stays a live link (it starts nothing).

The officer gates (`GuildPlanGate.svelte`) are unchanged: `Subscribe the guild` and `Move to Raid guild` go to that guild's `#billing`, where the checkout already carries `guild_id`. A link to `/premium` there would lose the guild (P6).

## 8. /premium anchors to add

| Id | Element | Used by |
|---|---|---|
| `premium` | the Premium `PlanCard` `<article>` (today only `free` has an id) | account panel `Plan` row |
| `coach` | Premium feature line 1 (`Coach: after every logged raid…`) | coach teaser, coach strip |
| `server-sims` | Premium feature line 2; lines 2 to 4 carry `data-anchor="server-sims"` | sim run control, premium-required message, bulk cap note |
| `history` | Premium feature line 5 (`See if you are improving…`) | History |
| `notifications` | Premium feature line 6 (`A message in your own Discord channel…`) | Discord notifications |
| `guild-leaders` | exists | guild billing chooser, footer `Guild plans` |

Code: `plansPageCopy.premium.features` becomes `{ text: string; anchor: 'coach' | 'server-sims' | 'history' | 'notifications' | null }[]`. The first line of an anchor carries `id`; every line carries `data-anchor` when it has one. Feature text is unchanged.

**Landing.** Each anchored element has `scroll-margin-top: 33vh`, so the card's name and price stay above the line on a 900 px desktop. The highlighted lines are the lines whose `data-anchor` matches `:target` (`.features:has(#server-sims:target) [data-anchor='server-sims']`, one rule per id). Highlight: ground `rgba(229,185,85,.08)`, an inset 2 px `--color-gold` left rule (`box-shadow: inset 2px 0 0 var(--color-gold)`, no border, so no reflow), `--radius-control`, 6 px 8 px padding applied to every feature line at all times so the highlight changes no box. No animation. The page ships no new JavaScript (it ships none today).

## 9. Hierarchy

The doors are the bar's primary elements. The tag is secondary: 11 px against the doors' 13 px, and the only gold in the doors. When the account control is open it is the bar's one gold-bordered element (polish §1.5); the tag's border is 35% gold and never glows. In the account panel `Upgrade` is the only gold below the character list. In a feature, the feature's own control stays primary; the Premium link is the last element of its row.

## 10. Copy

| Key | String |
|---|---|
| `PREMIUM_NAV_ITEM.label` (`lib/nav.ts`) | `Premium` |
| `selectorCopy.account.upgrade` | `Upgrade` |
| `selectorCopy.account.freePlanHref` | `/premium#premium` |
| `simCopy.runOnServersPremiumTag` | `Premium` |
| `simCopy.runOnServersFreeAria` | `Run on our servers, Premium` |
| `simCopy.getPremium` | deleted |
| `guildBillingCopy.chooser.compare` (the chooser's copy object) | `Compare Guild and Raid guild` |

No other string changes.

## 11. Phone and desktop

Never collapses: the tag from 1150 for a free or signed-out viewer; the Menu row; `Upgrade` on a free Plan row at every width. Collapses: from 1024 to 1149 the tag leaves the bar. Below 1024 it is the Menu's last row. At 2000 the tag sits inside the 1344 column with the doors, never at the window edge.

## 12. Acceptance (screenshots the implementer attaches)

| # | Viewport | State |
|---|---|---|
| A1 | 1440 | signed out, top of `/`: tag after `Get set up` |
| A2 | 1440 | free, signed in: tag visible; account panel open showing `Plan` · `Free` · `Upgrade` |
| A3 | 1440 | Premium (owner's grant account): no tag; account panel open with today's plan row |
| A4 | 2000 | free: tag in the 1344 column |
| A5 | 1150 | signed in as a 15-character name: no overflow, no wrap |
| A6 | 1100 | tag absent from the bar |
| A7 | 1440 | tag hover, focus-visible and current (`/premium`) |
| A8 | 390 and 360 | free: Menu open with the `Premium` row last; account sheet with `Upgrade` |
| A9 | 390 | Premium: Menu with no `Premium` row |
| A10 | 1440 and 390 | `/premium#coach`, `#server-sims` (three lines tinted), `#premium` landings |
| A11 | 1440 | `/sim` free signed in: `Run on our servers` + `Premium` tag; Premium: today's button |
| A12 | 1440 | report coach teaser, free: button lands on `/premium#coach` |
| A13 | 1440 | guild billing chooser with `Compare Guild and Raid guild` |
| A14 | n/a | Lighthouse: CLS 0 on the bar on `/`, `/sim`, `/premium` with a stored `fs.premiumSource=grant`, and on a fresh profile signed in to a granted account (the `hold` path) |

Checklist:

- [ ] Tag is `.pill.pill-site`, no dot, after `Get set up`, inside `#primary-nav`.
- [ ] Hidden for own, seat, grant and fallback; shown signed out and free.
- [ ] Hidden before first paint from the hint; `hold` (visibility) after paint; never a `display` toggle after paint.
- [ ] `/premium`'s duplicate `data-premium` head script deleted; Base owns it.
- [ ] Menu row last, 52 px, chevron.
- [ ] Plan row `Upgrade` only for `none`; one link; href `/premium#premium`.
- [ ] Every link in §7 carries its anchor; sim `Get premium` replaced.
- [ ] Anchors in §8 exist; highlight changes no box; no new JavaScript on `/premium`.
- [ ] No string outside §10 changed; none of §2's forbidden words appear.
- [ ] `Header.test.ts` asserts `PREMIUM_NAV_ITEM` order; `account-panel.test.ts` covers `upgrade` for all six sources.

## 13. Owner decisions

1. Signed-out visitors see the tag (Warcraft Logs shows `Subscribe` to everyone). The alternative is free accounts only.
2. A guild officer whose guild pays but who holds no seat has no personal Premium, so the tag shows for them.
3. From 1024 to 1149 the tag leaves the bar rather than squeezing the doors.

## 14. Mocks

| Board | Width | What it proves |
|---|---|---|
| `renders/premium-entry-header-1440.png` | 1440 | signed out, free and Premium bars; the tag's hover, focus and current states |
| `renders/premium-entry-header-2000.png` | 2000 | the tag inside the centred column |
| `renders/premium-entry-phone.png` | 390 | Menu free and Premium; the account sheet free with `Upgrade` |
| `renders/premium-entry-account-1440.png` | 1440 | the account panel free beside Premium |
| `renders/premium-entry-upsell-1440.png` | 1440 | the sim run control free and Premium; the `/premium#server-sims` landing |

Mock notes: the character name and BattleTag are illustrative. The Premium feature lines are `plansPageCopy.premium.features` with today's figures written out (`CAPS` 5,000 and 400, `LANE_ITERATION_CEILING` 100,000 and 30,000); the page renders them from `lib/sim/precision.ts`.
