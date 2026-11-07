# Premium, checkout, the account page and the shared build in the flat web system: experience spec

Author: ux-designer, 2026-11-08. Mock first (tenet 15): the build matches these mocks region by region. Nothing is built from this round. Ship mode: the coordinator runs the wow-player persona on section 17 and records the rulings as section 0A.
Surfaces: `web/src/pages/premium.astro` (`/pricing` is a 301 to it in `web/public/_redirects`, so there is no second page), `pages/premium/checkout.astro` with `components/CheckoutLauncher.svelte` and `components/checkout/*`, `components/premium/*` (`PlanCard`, `PriceBlock`, `IntervalControl`, `GuildRules`, `PremiumCoachSample`, `model.ts`); `pages/account.astro` with `components/Account.svelte` (mode `account`) and `components/account/*` (`AccountTitle`, `AccountLoading`, `CharacterList`, `CharacterRowLink`, `CoachLastRaidLine`, `AccountPremiumBlock`, `AccountSeatOffer`, `AccountGuildBlock`, `AccountGuildRowStep`, `AccountGuildPlanLine`, `AccountDiscordLink`, `AccountDiscordRole`, `AccountDiscordBlock`); the shared build `/b/:id`, rendered by the API (`api/internal/site/handler.go`, `page.go`: the page template mounts `web/src/planner-island.ts`, the planner read-only; `messagePage` for a missing or failed build) and the Worker's fallback `web/src/pages/b-unavailable.astro`.
System: `design/specs/2026-11-05-web-layout.md` (2 principles; 3.1 to 3.5 tokens, type, the control bar, buttons; 3.6 the character lives in the header selector; 3.7 L3 split; 3A primitives, the guard, parity regions). Rulings that stand: setup (`Steps` with one gold step, `Page layout="form"` for a one-task page, the head answer reads the header selector and never switches it, `Capture`, `Code`), planner (the `/b/:id` island is the planner read-only: section 5.1 "Shared link"; `TalentEditor`, `PointOrder`, `Ordered`, `Switch`; Send is the primary signed in, Copy for the game signed out; G9 every navigation link ends in `›`), guides (`PlateColumns`, `RowsSlim`), sim (the system Button's case; a `<button>` for an action, an `<a>` for a navigation).
Memory rules applied: pricing direction (Premium $6/$50, Guild $15/$150, Raid guild $25/$250; founder $5/$40, $12/$120, $20/$200 until 6 Jan 2027: every figure here comes from `lib/billing/plans.ts`, none is typed), stripe-billing (Stripe is merchant of record; the card never reaches us), owner-account-grant (a grant is a real Premium source with its own words), one-way-to-switch (no character switch on these pages; the account's main is a setting, the header selector is the switch), blizzard-profile-import (Battle.net is the primary character source; the addon is detail), one-crest-language, one-job-per-page, provenance-quiet, wide-screen-check, visual-quality-2026-10-28 (no empty example card; nothing plain), legal-entity (Commish LLC, `help@foreversixty.gg`), design-work-lives-in-repo.
Type: Marcellus once per page (the H1: `--color-strong` on Plans, checkout and account; the class colour on a shared build), Barlow everywhere else, no condensed face, no caps headings (12 px caps labels only). Mono only for literal text a player types or matches (the pairing code). One Plate per page.

## 0A. Rulings 2026-11-08 (owner delegated; the wow-player persona ruled, SHIP WITH FIXES)

- **YES:** PR1, PR3, PR4 (once MUST 2 is fixed), PR5, A1, A2, A3, A5, B1, B3, P1.
- **PR2 YES with change:** founder is said once: the Premium column's `Founder` Tag goes (the notice and the price line stay).
- **K1 YES with change:** a guild checkout in its free trial never says `Stripe took the payment.`; it reads `Card saved. First charge {date}.` A paid Premium checkout keeps `Stripe took the payment.`
- **K2 YES with change:** above the three next-step links, the renewal line `{price} a year · renews {date} · receipt from Stripe by email`.
- **A4 YES with change:** an account that already has Premium reads `Link it and the role turns on.` (MUST 3).
- **A6 YES with change:** the deletion sentence per MUST 4.
- **B2 YES with change:** Fork moves into the first screen and is named `Edit a copy`.
- **Q8:** the claim rule and the sample disagree (`Only a guild's guild master can claim it` vs an Officer holding the claim): the build checks the API's real rule and fixes whichever is wrong. **Q10:** show `Shared by {name}` when the share record has the sharer. **Q12:** the not-found page adds `Ask whoever sent it for a fresh link.` beside `Talent builds in the class guides ›`.
- **MUST 1:** the founder line reads `Founder price for anyone who subscribes by 6 Jan 2027. $50 after that.` (and the same shape on Guild and Raid guild), never `then $50` beside `the price stays`.
- **MUST 2:** the guild-plan rules and refunds.md say the same thing. refunds.md is final legal text (owner, 2026-10-24), so `plans-page-copy.ts`'s `Ask within 14 days of a first charge for a refund.` changes to match it, not the other way round.
- **MUST 3:** the Premium role sentence per A4.
- **MUST 4:** `Email help@foreversixty.gg with your BattleTag and we delete your account and everything on it, within thirty days.`; the pseudonym note `Reports themselves are never deleted or rewritten` is checked against privacy.md and matched to it.
- **MUST 5:** on /b/:id, a build of another class is one plain line above `Copy for the game` (no greyed caps box), inside the rail.
- **SHOULD:** the guild checkout button reads `Start 7 free days · then {price} a year`; the shared build rail heading `Guide build`; HPS without a decimal.
- **For the owner (not decided here; the build leaves the shipped wording):** who is on the card statement under Managed Payments (`Commish LLC d/b/a Forever Sixty` vs Stripe as merchant of record), and which address the 7-day renewal email goes to when Battle.net gives no email.

## 0B. Rulings applied to the mocks (2026-11-08, build round)

Every string below now comes from the copy modules the facts reader bundles; the generator's `PROPOSED` keeps only the two API message-page lines.
- MUST 1: `priceCell().founderLine` = `Founder price for anyone who subscribes by 6 Jan 2027. {list} after that.` on every plan and interval.
- MUST 2: the `Refunds` rule = `Email help@foreversixty.gg for a refund. Cancelling stops the renewal and the plan runs to the end of the period you paid for; there are no part-period refunds. If you were charged twice, or paid for something that did not work, we put it right.` (refunds.md); `REFUND_WINDOW_DAYS` retired.
- MUST 3: `discordRoleCopy.leadOwned` `Link it and the role turns on.` on an account that has Premium.
- MUST 4: `accountPageCopy.deleteBody` verbatim; `pseudonymNote` ends `Your reports are kept until you delete them or your account.` (privacy.md, Data retention).
- MUST 5: /b/:id on another class's build: the `otherClass` sentence as one plain line above `Copy for the game`.
- PR2: no `Founder` Tag on the Premium column. K1: new mock `web-checkout-trial` (`Card saved. First charge 15 Nov 2026.`). K2: `checkoutCopy.renewalLine` above the next steps. B2: `plannerCopy.fork` = `Edit a copy`. Q10: the build record has no sharer, so no `Shared by` line. Q12: `Ask whoever sent it for a fresh link.` on not-found. SHOULD: `Start 7 free days · then $120 a year` (`checkoutGuildCopy.ready.startTrial(price)`), `Guide build` (`bandCompareCopy.sharedHeading`), HPS whole (`scoreUnitCopy.figure`, `bandFigure`).
- Q8: the API (`api/internal/guilds/claim.go`) lets a guild master claim at once and an officer claim with a second officer's (or the guild master's) confirmation; whoever holds the claim subscribes. The copy was wrong: `CLAIM_RULE`, `lead.allClaimed`, `askLeader`, `claimedAsk.member` now say so.
- Found in the build round: Stripe's success URL named no plan (`/premium/checkout?status=success&session_id=…`), which the checkout page refuses as an invalid link; the API now returns `plan`, `interval` and `guild_id` (`successPath` in `api/internal/billing/checkout.go`).

## 0. Files

| What | Path |
|---|---|
| Generator (composes only from `ws_prims`; `/b/:id` reuses `ws_planner`'s planner blocks; every page passes `assert_no_surface_styles`) | `design/mocks/ws_premium.py` (41 mocks; `--list` names them; `compare <before dir>` writes the sheets) |
| Facts reader (copy modules and the pages' own models bundled with esbuild, never retyped; runs in CI without `src/data/generated` through setup's `generatedFallback` plugin) | `design/mocks/premium_facts.mjs` -> `design/mocks/data/premium-facts.json` |
| Before captures (production 2026-11-08, `/v1/me`, `/v1/devices` and `/v1/billing/checkout` answered from the fixtures; the valid `/b/:id` is the API template filled with a Survival record and the live island, as `shared-build.spec.ts` stubs it; writes `measure.json`) | `design/mocks/shoot_premium_before.mjs <out dir>` |
| Primitives added (appended block "added by the premium and account round"; nothing existing changed) | `design/mocks/ws_prims.py`: `Columns`, `Button(tone='danger')`, `PlanHead` (section 3A) |
| Mocks (42) | `design/mocks/web/`: `web-premium` (signed out, yearly, founder window open), `-monthly`, `-closed` (founder window closed), `-free` (Felsnap, free), `-own` (Thandrel: own yearly Premium at the founder price, billing Olympus XXVII on Guild), `-seat` (a seat from the guild), `-grant` (a grant), `-canceled` (back from Stripe's cancel), `-pastdue` (the last payment failed), `-ended` (a seat ended); `web-checkout-redirecting`, `-processing` (back from Stripe, the webhook not in yet), `-confirmed`, `-slow`, `-error`, `-notopen`, `-owned` (already has Premium through a seat), `-guild` (Which guild?), `-invalid`; `web-account` (Felsnap: one Battle.net character, free, no device, Discord not linked), `-out`, `-loading`, `-error`, `-new` (no character), `-several` (three characters from Battle.net and the addon, the main chosen, three devices), `-premium` (Thandrel: own Premium, a guild officer on Guild, a device, Discord linked and a channel connected), `-seat`, `-offer` (a seat offer to a paying member), `-ended`, `-pastdue`, `-claim` (a claim pending, and a guild where the member is not verified), `-grant`, `-confirm` (Leave, Revoke and Unlink asking at once), `-states` (every Plan variant, the offer answered, the pairing code, refreshing, sign out); `web-b` (Survival hunter, a stranger signed out), `-you` (Felsnap signed in, a hunter), `-holy` (Felsnap on a priest build), `-fire` (an untitled build), `-missing` (unknown id), `-error` (the API failed), `-unavailable` (the Worker's 503 page) |
| Renders (gitignored) | `design/mocks/renders/web/<mock>-{390,1440,1920,2000}.png` and `-fold.png` (all 41 at all four widths); `compare-*-<w>.png` |
| Tracked copies (256 colours) | `design/reviews/2026-11-08-premium-account/before/` (key live captures and `measure.json`); `.../after/` (the 390 and 1440 first screens of every mock, the 2000 first screens of the four leads, full pages of `web-premium`, `web-account-premium`, `web-b` at 1440 and 390, the compare sheets, `mock-measure.json`) |
| Regenerate | `node design/mocks/premium_facts.mjs && node design/mocks/planner_facts.mjs && python3 design/mocks/ws_premium.py`, then `node design/mocks/ws_shoot.mjs <mock> 390,1440,1920,2000` per name, then `node design/mocks/shoot_premium_before.mjs <dir> && python3 design/mocks/ws_premium.py compare <dir>`. Needs `npm ci` and `npm run sync` in `web/` |

**Data.** Build 1.60.1.70790. Clock pinned to 2026-11-08T10:00Z, zone UTC. The accounts are the home fixtures (`me-34.json` Felsnap, a level 34 Dwarf Beast Mastery hunter; `me-60.json` Thandrel, a level 60 Human Fury warrior, officer of Olympus XXVII; `me-new.json`) with, per state, the billing, seat, guild, Discord and device fields set in `premium_facts.mjs` (each marked DERIVED there). Two alts for `-several` (Brisa, a level 19 Night Elf druid from Battle.net with no gear yet; Korrin, a level 27 Dwarf Holy priest from the addon) and a second guild for `-claim` (Pinefeather Lodge) are DERIVED. Every price, date and line under a plan is the model's own output (`priceCell`, `buttonLabel`, `costLine`, `accountPlanLine`, `guildRowPlanLine`, `premiumOwnedView`, `offerStatusLine`); every account line is `characterDescriptor`, `setupAnswer`, `deviceStatusLine`, `premiumRowValue` on those accounts. The sample coach review is `coachSampleView` over `src/data/premium-coach-sample.json` with the build's spell icons. The shared builds are the planner facts' own states (`guide` Survival 5/13/33, `holy` 17/34/0, `fire` 14/34/3); their live score is the planner mocks' SAMPLE figure. The API's two message sentences are read from `handler.go`, the Worker page from `b-unavailable.astro`.

## 1. Diagnosis of the live pages (production 2026-11-08; `before/`, `before/measure.json`)

**/premium**
1. **The H1 is a slogan in condensed caps**: `WHAT IS FREE, AND WHAT THE PLANS ADD` over a caps eyebrow `PLANS`. A player who came to pay reads a sentence before a price (tenet 14).
2. **Seven framed blocks** (Free, Premium, the coach sample, Guild, Raid guild, How the guild plans work, Questions), each a gold-edged gradient card with an icon chip on two of them. Nothing is primary; the gold Premium border is the only hint.
3. **The page does not use the width.** The column is 1,248 px at 1440, 1920 and 2000 (the xl `--page-max` is not applied): at 2000 it fills 62% of the screen and the page is the same 3,474 px tall at all three widths.
4. **The answer is late on a phone.** At 390 the page is 6,551 px and the first `Subscribe` is at y 932, below an 844 px screen: the Free card (eight lines) stacks first.
5. **Signed in changes one line.** A Premium account sees a muted `You have Premium through your subscription.` inside the card and the button becomes `See your plan`; the head says nothing about the account, and an officer billing a guild on Guild sees the same `Subscribe a guild · $120 a year` as a stranger.
6. **Five answers hidden** in `Disclosure`s under Questions (cancel, refunds, Battle.net, ads, who am I paying). Each is one or two sentences (tenet 4).
7. **Two signals for one thing**: the `Founder` pill on the band and again on the Premium card, plus a gold founder line in each price block.
8. **What is right and stays:** every price built from `plans.ts`; the founder line and its end date; the per-member and per-raider lines; the guild trial line; the sample coach review (real model output, labelled Sample); the history promise; the ten rules; the honest voice (no urgency, no exclamation marks); the static page with no client script; the anchors `#coach`, `#server-sims`, `#history`, `#notifications`, `#premium`, `#guild-leaders` that contextual links land on.

**/premium/checkout**
9. **The answer is one line in a void.** Back from Stripe the page reads `You're all set.` in a 640 px column centred at x 400 while the band's title starts at x 96 (two left edges); nothing says what Premium now does or where to go (1440: `before/checkout-success-1440.png`).
10. **The guild step's column is centred too** (x 280 to 1,160) under a left-aligned band: three left edges on one page (`checkout-guild-1440.png`).
11. **No progress.** Redirecting, confirming and failing are the same one-line slot; a player cannot tell whether they were charged.

**/account**
12. **Six to eight cards** (Current character, Characters, Your reports, Devices, You, Guild, Plan, Discord). The current character is drawn twice: a hero card and its own row under it, with the same crest, name and descriptor (`account-premium-1440.png`).
13. **Lopsided at every desktop width.** The rail is a 440 px stack of five cards (1,500 px tall at 1440) and the main column ends at y 910: 1,000 px of empty page on the left.
14. **The plan is the last thing.** The Plan card starts at y 1,240 at 1440 (below a 1,000 px screen) and at y 2,124 at 390.
15. **Discord is in two places**: the link and the Premium role inside `You`, the notifications in their own card further down.
16. **Pipeline words**: the device row prints the raw platform `windows/amd64` (`Account.svelte` prints `device.platform`; `deviceStatusLine` with `platformLabel` already says `Windows`).
17. **No way to delete the account** on the page; the privacy page says to email (gap G6).
18. **What is right and stays:** one `/v1/me` read; the Battle.net refresh with its snapshot compare and result line; Set as main; the seat offer's consequence sentence; the plan words from `accountPlanLine` (own, own then a seat, grant, seat, ended, past due); the guild share level, plan line, next step and Leave with its two-step confirm; Revoke asking first; the pseudonym switch; every Discord state machine; the reserve that keeps the loading state at the ready page's height.

**/b/:id**
19. **The stranger's page is the old planner in a page with no column.** The API page links `planner-island.css` but has no `page-column`: the head, the trees and the rail run from x 0 to the screen edge at 1440, 1920 and 2000 (`b-shared-2000.png`), while the API's own header sits in the site column.
20. **The build's title is never shown.** The record says `Survival raid build`; the H1 says `Hunter talents` and the eyebrow `BUILD PLANNER`. The person who shared it named it, and the link preview card says it; the page does not.
21. **The one action is below the fold.** `Fork` is at y 1,077 at 1440 and y 1,918 at 390, at the end of the rail; `Copy for the game`, the action a stranger with the addon wants, is absent signed out.
22. **A disabled class picker and race select** lead the page (read-only), with `Sign in or paste an export to see your character here`: controls the stranger cannot use, in the first screen.
23. **Not found is a bare sentence** (`Build not found`, Cinzel gold, one link) with no way to the guides' builds.

## 2. The one question per page, and the answer in the first screen

| Page | Question | 390 x 844 (mock) | 1440 x 1000 (mock) | 1920 / 2000 (mock) |
|---|---|---|---|---|
| /premium signed out | "What does Premium add, and what does it cost?" | Title, lead, the interval line, the founder line, then the Premium column first: $40 /year and `Subscribe · $40 a year` (primary, y 602 to 646; live y 932) | The Plate y 353: both columns' prices, Subscribe at y 543; the coach sample beside it in the rail | Same; the rail 400 px, the plate ends y 1,043 at 2000 |
| /premium signed in | "What am I on, and what would more give me?" | The head answer: `Plan` Free or Premium, the date and price, `See your plan ›`; then as above with `Your plan` on the account's column | Same, the answer at the right of the title | Same |
| /premium/checkout | "Did it work, and am I charged?" | Three steps (Sign in, Pay with Stripe, Premium on) in the Plate y 219 to 421; the confirmed state's three next links under it | Plate y 197 to 399 in the 560 px form column at the page's left edge (x 96) | Same at x 168 |
| /account | "Is my account set up, and what is my plan?" | The main character (head answer), the Characters Plate from y 390, then Plan | Characters Plate y 282 to 344; the plan's button (`See plans`) y 750; the rail (Devices, Discord, You) beside | Same |
| /b/:id stranger | "What is this build, and how do I use it?" | Title (the build's name), meta, the fork line with `Fork` (y 260), then the trees | Title, the 51 points and the score, `Fork`, the three trees (Plate y 309 to 856) and `Copy for the game` (y 370) beside them | Same, wider trees |

No mock scrolls sideways at 390 (`scrollWidth` 390, `after/mock-measure.json`).

## 3. Reference surfaces, and what is borrowed

- **The Battle.net subscription page and the WarcraftLogs Premium page** (what a WoW player has bought a subscription from): one price figure per plan with the per-month line, one button per plan, the feature list under the price, the current plan marked. Borrowed: price figure first, one button, the `Your plan` mark. Not borrowed: three-tier upsell cards with a "Most popular" ribbon, crossed-out prices, a comparison grid of ticks against tiers.
- **Raider.IO and WarcraftLogs account settings** for /account: one page of rows (characters, connections, subscription) with the action at the right of each row. Borrowed: the row grammar and the action column. Not borrowed: tabs that hide each group.
- **The ElvUI installer steps** (already the setup page's `Steps`) for checkout: numbered steps, the current one marked, a tick on what is done.
- **Wowhead's talent calculator share links** for /b/:id: the shared link opens the tree as it was, with the build's own title, read-only, and one way to make it yours. Borrowed: the title, the read-only tree, one Fork. Not borrowed: ads and the calculator's controls on a read-only view.
- **Our own system as shipped:** `PageHeader` with the answer at the right, `Plate`, open `Section`s, `Row`/`Rows`/`RowsSlim`, `RowGrid`, `PlateColumns`, `Facts`, `Tag`, `Notice`, `Switch`, `Steps`, `Page layout="form"`, the planner's `TalentEditor`, `PointOrder`, `Ordered`, `Page('workbench')`.

## 3A. Primitives (added by this round; a design-review item under tenet 10, decision P1)

All appended to `ws_prims.py`; built in `web/src/components/ui/`.
| Primitive | What | Replaces | Next consumers |
|---|---|---|---|
| `Columns(cols, open_=True, second_first=False)` | `PlateColumns`' columns without the Plate: hairlines between, stacked below 1024, no side padding when open. `open_=False` is the padded form a caller puts in its own `Plate` when each column carries its own head (`PlanHead`). `second_first` stacks the second column first below 1024 (the plan the page sells leads on a phone) | `PlanCard`'s `CARD`/`CARD_QUIET` frames, the guild row's three-column grid | the guild settings' plan choice, the checkout guild step on wide screens |
| `Button(tone='danger')` | The shipped `button('danger')`: wipe text and a `#f0736b66` hairline, never filled; hover wipe border and `#ff8f88` text | the hand-styled Leave and Revoke confirms | every destructive confirm |
| `PlanHead(name, audience, value, unit, lines, tag, gold)` | A plan's head: name Barlow 700 22/28 `--color-strong` with an optional `Tag`, the audience Small, the price `Figure` big (36/40) in gold on the plan the page sells (Premium), `--color-strong` on the others, then its small lines (second line muted, the founder line gold) | `PriceBlock.astro`, `PlanCard`'s name and audience cells | checkout's plan summary on a phone |

`Notice`, `Steps`, `Switch`, `Facts`, `Row`, `RowGrid`, `RowsSlim`, `Prose(ruled)`, `Emblem`, `Figure`, `Tag`, `Page('form')`, `Page('split')`, `Page('workbench')` and the planner's blocks are reused unchanged.

## 4. /premium

`Page layout="split"` (main, then a rail of 340 px from 1024 and 400 from 1760). Region order: `head`, `head-title`, `head-meta` (none), `head-answer` (signed in), `controls`, `notes`, `plate`, `section-for-guild-leaders`, `section-how-the-guild-plans-work`, `section-questions`; rail `section-what-the-coach-says-after-a-raid`.

**4.1 PageHeader.** Title `plansPageCopy.title` (`Plans`), `--color-strong`, Marcellus 40/46 (32/38 below 768). Lead `band.lede` (17/26, 62 ch). The eyebrow and the slogan H1 `band.heading` retire (decision PR1). Answer, signed in only: label `accountPageCopy.planLabel` (`Plan`), the plan name as a `Figure` sm (`Free`, `Premium`, `Guild seat`, from `accountPlanLine(me).head.name`), one Small line `{when.label} {when.date} · {extra}` and the `Founder` Tag when `founder`, link `See your plan ›` (`premiumSourceCopy.buttons.seePlan`, `/account#plan`). Signed out: no answer. Below 900 the answer stacks under the lead, left-aligned.

**4.2 ControlBar** (`Billing interval`, sticky, 56 px): one segmented group `Yearly | Monthly` (`interval.labels`; yearly default, the URL `?interval=monthly` is the truth, written with `replaceState`); `end`: Small `Running a guild?` and the link `Guild and Raid guild are further down ›` to `#guild-leaders`. Phone: the bar's summary line `Yearly` and `Change ▾`. The page stays static: the radios switch with `:has()` as shipped, both intervals in the DOM in fixed cells (no shift).

**4.3 Notes** (`notes`, one `Notice` line each, neutral): the founder line (`Tag Founder` + `founder.open`) while the window is open (the pre-paint `data-founder` attribute as shipped; with the window closed the line is absent and every `Founder` mark goes with it). Then at most one state line: `checkoutCopy.canceled` (`role="status"`, from `?canceled=1`); past due: `Tag Payment failed` (wipe) + `billingBlockCopy.pastDueBanner` (`role="alert"`) + `Manage billing ›`; a seat ended: `accountPlanCopy.seatEnded.text` + `Come back at $40 a year ›` (the account's `resubscribe_url`). The founder Tag on the Premium column and the band's pill merge into this one line (decision PR2).

**4.4 The one Plate: the player's plans** (`Columns(open_=False, second_first=True)`, accessible name `Plans for a player`). Two columns split by a hairline; below 1024 Premium stacks first (decision PR3).
- **Free**: `PlanHead` (`free.name`, `free.audience`, `$0`, `free.note`; the `Your plan` Tag in kill when the account is free); the seven `free.features` as a ruled list (`Prose ruled lane`: 15/22, a 6 px gold dot, hairlines); signed out a foot link `Sign in with Battle.net ›`.
- **Premium**: `PlanHead` (`premium.name`, `premium.audience`, the price from `priceCell` in gold: `$40` `/year`, `$3.33 a month, billed yearly`, the founder line `Founder price until 6 Jan 2027, then $50.` in gold); the page's one primary Button `buttonLabel` (`Subscribe · $40 a year`, to `checkoutHref`); Small `premium.support`; the seven feature lines ruled, each first line of an anchor carrying its id (`#coach`, `#server-sims`, `#history`, `#notifications`); Small `premium.scope`. Owned (own, seat, grant): the Tag `Your plan` replaces Founder, the button becomes secondary `See your plan` under the source line `premiumSourceCopy.premiumCard.{own|seat|grant}` (15/24 strong), and the page has no primary.

**4.5 Section `For guild leaders`** (`guildRegion.heading`, `#guild-leaders`). `guildRegion.lede` (15/24), Small `guildRegion.members`, then `Columns` open: Guild and Raid guild, each `PlanHead` (price strong, not gold; lines: the second line, the founder line, `costLine`), a secondary Button `Subscribe a guild · $120 a year`, the trial line (`trialParts`: the lead strong), the features ruled. An officer whose guild is on that plan sees Tag `Olympus XXVII` (kill) on its column, `checkoutGuildCopy.onPlan.youPay` (`Olympus XXVII is on the Guild plan. You pay for it.`) and `Manage billing` in place of Subscribe (or `onPlan.billedBy` with no button). Under the columns Small `historyPromise`.

**4.6 Section `How the guild plans work`**: the ten `rules` as open Rows (term as title 15/20 700, sentence as sub) in a `RowGrid` of two columns from 768, one below. Replaces the bordered `GuildRules` table.

**4.7 Section `Questions`** (meta Small `plansPageCopy.foot`): the five `faq` entries as open Rows in a two-column `RowGrid`, answer visible (decision PR4). The `Disclosure`s and the Card retire.

**4.8 Rail: Section `What the coach says after a raid`** (`coachSample.label`, meta Tag `Sample` muted, sticky from 1024 is not used: the rail is shorter than main). The sample review: crest 40 (`ClassCrestRing`, Fury warrior) with `Your Fury coach` (16/20 700) over Small `coachSample.context`; the opening (15/23); label `Fix first` and one selected Row (gold ground, 3 px gold bar) with the spell icon 40, the title and the why; label `Then` and its Rows; label `Doing well` and its slim Rows with icons; Small `coachSample.caption`. It is the evidence for Premium's first line, beside it at 1440 and up; below 1024 it follows Questions.

## 5. /premium states

| Mock | When | Head answer | Notes | Premium column | Guild columns |
|---|---|---|---|---|---|
| `web-premium` | signed out, yearly, window open | none | founder line | primary Subscribe $40 a year | Subscribe a guild |
| `-monthly` | `?interval=monthly` | none | founder line | `$5 /month`, `Or $40 a year.` (`priceCell`), Subscribe $5 a month | `$12 /month`, `$0.48 per member a month at 25 members.` |
| `-closed` | after 6 Jan 2027 08:00Z | none | none | `$50 /year`, no founder line, no Founder tag | list prices |
| `-free` | Felsnap, free | `Plan` `Free`, `See your plan ›` | founder line | Free carries `Your plan`; Premium sells | as signed out |
| `-own` | Thandrel, own yearly founder, billing Olympus XXVII on Guild | `Premium`, `Renews 4 Nov 2027 · $40 a year`, Founder Tag | founder line | `Your plan`, `You have Premium through your subscription.`, `See your plan` | Guild: Tag `Olympus XXVII`, `onPlan.youPay`, `Manage billing` |
| `-seat` | a seat from Olympus XXVII | `Guild seat`, `A seat from Olympus XXVII` | founder line | `Your plan`, `premiumCard.seat` | as signed out |
| `-grant` | a grant, no end | `Premium` | founder line | `Your plan`, `premiumCard.grant` | as signed out |
| `-canceled` | `?canceled=1` | as the session | founder line, `Checkout was canceled. Nothing was charged.` | as the session | |
| `-pastdue` | own, `past_due` | `Premium`, `Renews 11 Nov 2026 · $40 a year` | founder line, the payment-failed line with `Manage billing ›` | `Your plan`, `See your plan` | |
| `-ended` | `ended_seat` within 30 days | `Free` | founder line, `Your seat from Olympus XXVII ended 2 Nov.` `Come back at $40 a year ›` | sells | |

Reads: the founder instant (`FOUNDER_WINDOW_CLOSES`, pre-paint), `?interval=`, `?canceled=1`, and signed in `/v1/me` (`entitlements.billing`, `seat`, `ended_seat`, `guilds[].plan`) through one small island (`PremiumAccount`, `client:idle`) that fills the answer, the `Your plan` marks and the state line. Until it answers the pre-paint `data-premium` hint (`fs.premiumSource`, shipped) decides the Premium column's owned look, and the answer reserves its box when `data-session` is set (no shift). Signed out the page stays zero-JS apart from the shipped inline scripts.

## 6. /premium/checkout

`Page layout="form"` (560 px at the page column's left edge: x 96 at 1440, 168 at 2000; decision C1 of setup). Regions `head`, `head-title`, `head-answer`, `plate`.
**PageHeader.** Title: `Premium` (plan name) for a Premium checkout; `checkoutGuildCopy.band.title` (`Subscribe a guild`) for a guild plan with the lead `lead.{state}`; `band.invalidTitle` (`Checkout`) for a bad link. The full-bleed band and its caps label retire (decision K1). Answer: the plan summary, `Figure` gold `$40` `/year` with the `Founder` Tag, Small `$3.33 a month, billed yearly` (guild: the size line, `First 7 days free` strong), link `Change plan ›` (to /premium) shown only before payment.
**The Plate: `Steps`** (accessible name `Checkout`), three rows: 1 `Sign in` (`Signed in as {battletag}`), 2 `Pay with Stripe`, 3 `Premium on` (titles proposed). Signed out the launcher redirects to `/login?next=` as shipped (no row is drawn).

| Mock | Row 1 | Row 2 | Row 3 | Under the Plate |
|---|---|---|---|---|
| `-redirecting` | ok | busy · `checkoutCopy.redirecting` (`role="status"`) + `Stripe handles the card; we never see it.` (P) | todo | |
| `-processing` (`?status=success`, polling) | ok | ok · `Stripe took the payment.` (P) | busy · `checkoutCopy.processing` | |
| `-confirmed` | ok | ok | ok · `checkoutCopy.confirmed` | links `Run a server sim ›` (/sim), `Upload a log for the coach ›` (/logs), `See your plan ›` (decision K2) |
| `-slow` (4 polls, no entitlement) | ok | ok | warn · `confirmationSlow` | primary `Check again` (re-polls; P) |
| `-error` | ok | miss · `genericFailure` (`role="alert"`) · primary `Try again` | todo | |
| `-notopen` | ok | miss · `notOpenYet` | todo | (also `comebackGone` and a conflict's own sentence, same row) |
| `-owned` (409: already Premium) | the Plate is `premiumOwnedView`: `You already have Premium.` (h3 18/24), the line (15/24 strong), line 2 Small, the action (`Manage billing` button or `See your plan` link) | | | portal failure: `genericFailure` under it |
| `-guild` | the Plate is the guild list (`Which guild?`): a Row per guild with `Emblem` 40, the name, the card's sentence (`ready.holdsClaim`, `askLeader`, ...) and its facts line (`facts.verified`, `onRoster`, `noPlan`); trail: the card's one action (primary `Subscribe · $120 a year` on the first ready guild; `Copy link` for a link to send) | | | Small `foot.missingGuild` `Get set up ›`, `foot.premium` `Get Premium ›` |
| `-invalid` | no Plate: `invalidLink.text` `Premium page ›` (`role="alert"`) | | | |

The guild step's states (loading, load error, empty A and B, claim, start claim, confirm claim, checkout open, frozen, on plan, switch plan, the copy field) keep their shipped words and logic (`CheckoutGuildStep`, `GuildStepCard`, `ClaimConfirmBlock`); each card becomes one `Row` in this Plate with its action in the trail and its confirm as the Row's foot. Only the ready and ask states are mocked; the rest are a frame change, not a content change.

## 7. /account

`Page layout="split"`. Region order: `head` (`head-title`, `head-meta`, `head-answer`), main `section-characters` (holding the `plate`), `section-plan`, `section-guild`; rail `section-devices`, `section-discord`, `section-you`. The account's identity, its characters and its plan are main; the connections are the rail (decision A1: the shipped page put the plan last in the rail and left the main column empty below y 910).

**7.1 PageHeader.** Title `Your account` (`accountPageCopy.title`), strong. Meta: the BattleTag, `Signed in with Battle.net` (or `Signed in by email`), and a secondary sm Button `Sign out` (`role` button; it asks first, 7.7). Answer: the account's main character, as /setup's (`setupAnswer`): crest 40, the name 18/22 700 in the class colour, Small `Level 60 Fury Warrior`, links `Open Thandrel ›` and `Best in slot ›`. It reads the main (`main_character_key`, else the newest), never the header pointer, and never switches it. The eyebrow, the identity tile and the `Current character` hero card retire (decision A2: the character was drawn twice).

**7.2 Section `Characters`** (meta Small `Imported from Battle.net 2 days ago.`, action `Refresh from Battle.net ›`; while it runs `Refreshing…` with the rows at 60%; the result line replaces the meta for 8 s; an error stays in wipe with `Try again ›`). **The one Plate**: a Row per character: crest 40 lead; title the name in the class colour with Tags `Current` (gold, the header's character) and `Main` (muted); sub `characterDescriptor` and the guild name on a second line; trail: the build-source Tag (`Addon`, `Battle.net`, or `No gear yet`), `Set as main` (secondary sm, not on the main), `Open ›`. Under the Plate: Small `introBattlenetLine` with its two links, and `Your reports ›` (to /reports; the `Your reports` card and `Your ratings` card retire from this page, decision A3). The coach's last-raid line (`CoachLastRaidLine`) moves into the Plate as its foot row when there is a last raid; with no raid it draws nothing (it drew `Your last raid isn't available right now. Try again` on a failed read).
Empty (`-new`): the Plate holds the gold `characterListCopy.empty` sentence, primary `Get set up` and secondary `Paste an export` (shipped `EmptyState`); the Refresh action stays in the head.

**7.3 Section `Plan`** (`#plan`, the anchor `See your plan` lands on). Top line: the plan name `Figure` sm with the `Founder` Tag, and at the right `Renews` / `Ends` / `Until` over the date (`accountPlanLine.head`). Lines: `head.extra` (`$40 a year`, `Then a seat from Olympus XXVII`), `cancelAnytime` under a renewing own plan, `line2` (grant and seat words), or the past-due line (Tag `Payment failed` wipe + `pastDueBanner`, `role="alert"`). Ended seat: `seatEnded.text` + the link `Come back at $40 a year` in place of the plan lines. A seat offer: one selected Row (gold ground, 3 px gold bar) with `offer.title`, `offer.from`, `offer.consequence` (15/24 text) and in its foot primary `Accept the seat` and secondary `Decline`; answered, it becomes one Small status line (`offer.accepted` / `offer.declined`, `role="status"`). The four benefits as slim Rows: label left, the value right (a gold link `›` when live and linked: `On · 4 runs this month ›`, `Open on your main character ›`, `On · 2 events ›`, `Shown on your characters ›`; muted text when off: `In your browser`, `Sparkline only`, `Premium only`, `Not shown`). Foot: `Manage billing` (secondary; primary when past due) or `Billing history`, or primary `See plans` on a free account.

**7.4 Section `Guild`** (only with guilds). Per guild: a Row with `Emblem` 32, the name with Tags (plan tier `Guild` gold or `No plan` muted, `Not verified` muted, `Claim pending` gold, proposed), sub the rank, trail `Open ›`; slim Rows `Share with this guild` with the `Select` (Roster only, Gear, Gear and bags; `Saved. Applies to 1 character.` for 3 s), `Guild settings` with the plan line (`Renews 4 Oct 2027 ›`), the claim line (`claimPendingYours`, trail `Copy link`), the next step (`Olympus XXVII has no plan.` `See the next step ›`); then `Leave guild` (danger sm). Asking: the confirm lines (`leaveConfirmLine`, `leaveLosesOfficer` or the claim release lines) and danger `Leave guild` with secondary `Cancel` (focus lands on Cancel).

**7.5 Rail: Section `Devices`**: a Row per companion (name 15/20 700, sub `Windows · last seen 2 days ago` from `deviceStatusLine`, never the raw platform; trail `Revoke` sm); asking: `Revoke JHUNT-PC? The companion signs out.` with danger `Revoke` and `Cancel`. None: Small `noDevices`, secondary `Download the companion`, link `Use a pairing code instead ›`; with devices the link is `Pair another computer ›`. Code shown: Small `codeFallbackLine` and the code as `Code big`. The `Updated …` stamp retires (provenance quiet: each row says when it was seen).

**7.6 Rail: Section `Discord`** (`#discord`, one place for everything Discord, decision A4): a Row `Linked to Discord as thandrel` with `Unlink` (or the lead sentence and `Link Discord`; `leadSuggested` when a guild signs up in Discord); a Row `Premium role in Discord` with its lead and `Link Premium role`; label `Notifications` and the shipped block in its state (free: `freeSentence` and `See plans ›`; empty: Connect a Discord channel; saved: `destinationChannel`, three `Switch`es, `Send test message`, `Connect a different channel ›`; warning, switched off, lapsed as shipped).

**7.7 Rail: Section `You`**: the pseudonym `Switch` and its note. Then the last rail Section, `Delete account`: one line and a danger `Delete account` that opens the confirm dialog (G6, closed 2026-11-07). `Sign out` asks first (proposed `Sign out of Forever Sixty on this browser?`, danger `Sign out`, `Cancel`) because it now sits beside the title (decision A5).

## 7A. /account states

| Mock | Account | Characters | Plan | Guild | Devices | Discord |
|---|---|---|---|---|---|---|
| `web-account-out` | none | `Page form` Plate: `accountSignInCopy.reason`, primary `Sign in with Battle.net` | | | | |
| `-loading` | session hinted | Skeleton rows at the final geometry (two 56 px rows) | Skeletons | | | |
| `-error` | `/v1/me` failed (not 401) | Plate: `That did not work; try again` (`ACCOUNT_FAILED`, alert) + primary `Try again` | | | | |
| `-new` | no character | empty Plate | Free, See plans | | none | not linked, free |
| `web-account` | Felsnap, Battle.net, main | one row | Free, See plans | | none | not linked |
| `-several` | Felsnap + Brisa (Battle.net, no gear) + Korrin (addon) | three rows, Set as main on two | Free | | three (Windows, Linux, macOS `paired, not seen yet`) | |
| `-premium` | Thandrel | one row, guild line | Premium, Founder, Renews 4 Nov 2027, $40 a year, Manage billing | Olympus XXVII, Guild, Officer, Renews 4 Oct 2027 | one | linked, channel connected |
| `-seat` | Felsnap | | Guild seat, `Given by Thandrel…`, no billing button | Olympus XXVII, Guild, `Renews … · you hold a seat` | one | |
| `-offer` | Felsnap, own | | the offer Row first (primary Accept) | | | |
| `-ended` | Felsnap | | the ended line, `Billing history` | | | |
| `-pastdue` | Felsnap | | Payment failed line, primary Manage billing | | | |
| `-claim` | Thandrel | | Free | Olympus XXVII (`No plan`, `Claim pending`, the claim line with Copy link, next step), Pinefeather Lodge (`Not verified`, next step) | | |
| `-grant` | Felsnap | | Premium, `Given by Forever Sixty. Nothing is billed to you.` | | | |
| `-confirm` | Thandrel | | | Leave asking | Revoke asking | Unlink asking |
| `-states` | | refreshing | every plan variant (free, monthly, yearly, ending, grant, seat, ended, past due), offer accepted and declined | | the pairing code | sign out asking |

Reads: `GET /v1/me` (user, characters, guilds with `plan`, `claim`, `subscribe`, `can`; entitlements and billing; seat, seat_offer, ended_seat; server_sims; notifications; discord_link), `GET /v1/devices`, the refresh snapshot on `?refreshed=1`, `?signed_in=1` (the one-time banner `signedInBanner` as a `Notice` under the head).

## 8. /b/:id

The API renders it; the build is a change to `api/internal/site/page.go` (the template wraps `main` in the site's `page-column` and passes the record's title) and to the island's read-only props. Regions `head` (`head-title`, `head-meta`, `head-answer`), `notes`, `plate`, then the planner's `section-point-order`, `section-in-the-game`, `section-level-60-build`.
**PageHeader.** Crest 56 and the Title: the record's title (`Survival raid build`), else `{Race} {Class}` (`builds.Describe().Title`), in the class colour (decision B1). Meta: `Shared build` (P) · `Night Elf · Level 60 · Survival Hunter` (the planner's meta) · `5/13/33` (`SplitText`) · `41 views` (record `views`, P). Answer: the planner's own (Points left `0` gold, `51 of 51 spent · level 60`; the score `612` DPS `± 4 · with the level 60 gear set`; `Sim this build ›`).
**Notes**: one line, `plannerCopy.sharedLinkNote` with the secondary Button `Fork` beside it (decision B2: the Fork moves from the end of the rail to the first screen).
**No ControlBar** (decision B3): the class and race are the record's; a disabled picker is a control a stranger cannot use. `READ_ONLY_REASON` stays the status line's refusal when a cell is pressed.
**The Plate**: `TalentEditor` read-only (no per-tree Reset, no Reset… in the foot; the source line stays). Main below it: `Point order`. Rail: `In the game` first (signed out: primary `Copy for the game`, the hint, `Get the addon and companion ›`; signed in: the send state for the header's character: `Send this build to Felsnap (replaces their 23 points)` as primary, or `This is a Priest build; Felsnap is a Hunter` disabled with Copy as the primary), then the level build compare without `Load`.

| Mock | Record | Viewer | Primary |
|---|---|---|---|
| `web-b` | Survival raid build, Night Elf hunter 5/13/33 | signed out | Copy for the game |
| `-you` | same | Felsnap (hunter 34) | Send this build to Felsnap (replaces their 23 points) |
| `-holy` | Holy raid healing, Human priest 17/34/0 (HPS) | Felsnap | Copy for the game (Send disabled, other class) |
| `-fire` | untitled: H1 `Human Mage` | signed out | Copy for the game |
| `-missing` | 404 | any | `Open the build planner` (primary) + `Talent builds in the class guides ›` (P) |
| `-error` | 500 | any | as missing, with `Build unavailable` / `That build could not be loaded just now.` |
| `-unavailable` | the Worker's 503 (`b-unavailable.astro`) | any | as missing, its own two sentences |

The message pages are `Page layout="form"` with the PageHeader (title, the sentence as meta) and the actions under it; Cinzel gold and the bare link retire.

## 9. Every state of every interactive element

All Buttons are the system Button (primary gold 44 px, secondary outline, sm 36 px and 44 below 1024, danger as 3A); links gold 600 with `›` where they navigate. Focus-visible: 2 px gold outline, 2 px offset, everywhere. Active: pressed 1 px down. Disabled: `aria-disabled`, 45%, no hover, the reason in the title. Touch: 44 px targets below 768.
- **Interval segmented control**: default `--nav`; hover S2 and `--strong`; selected gold text, `#e5b9551a`, 2 px gold underline; focus-visible; arrows move and select (radio group); no disabled, loading or error state (static, both prices in the DOM).
- **Subscribe / Subscribe a guild** (links to checkout): default, hover (`--gold-hover` ground on primary, gold border on secondary), focus-visible, active; owned: replaced, never disabled.
- **See your plan, Manage billing, Billing history**: Manage billing is a `<button>` that opens the portal: busy (spinner, `aria-busy`, disabled) while it opens; error: `That did not work; try again` under the foot (`role="alert"`).
- **Checkout Steps actions** (`Try again`, `Check again`): default, hover, focus, busy (`aria-busy`, disabled while the call runs), error (the row's mark `miss`, its sentence in `role="alert"`).
- **Guild card action** (Subscribe, Claim, Copy link, Share): as shipped in `GuildStepCard` (busy words in `checkoutGuildCopy.busy`, errors in `checkoutErrors` / `claimErrors` on the card).
- **Character row**: the row is not a link (it holds `Set as main`); `Open ›` is the link. `Set as main`: busy (`aria-busy`, every Set as main disabled), error `Could not set your main just now.` under the Plate (wipe, alert). Hover on the row: S2.
- **Refresh from Battle.net**: link; refreshing (`Refreshing…`, `aria-disabled`, the rows 60%, a 2 px gold bar at the Plate's top), result line, error with `Try again ›`.
- **Seat offer**: Accept (primary) and Decline (secondary): `Accepting…` while busy, both disabled; failed `That did not save. Try again.` (alert); answered: the status line, the buttons never return.
- **Share select**: `Select` compact: default, hover, focus, open, disabled while saving; saved line 3 s.
- **Leave guild, Revoke, Unlink, Sign out**: secondary sm (Leave and Revoke danger sm); press: the confirm in place with danger confirm and secondary Cancel, focus on Cancel; Escape cancels; busy disables both.
- **Switches** (pseudonym, Discord events): on gold, off `--line`; focus-visible ring; disabled 45% while saving; error line under the section.
- **Fork** (`/b/:id`): secondary; press makes the planner editable in place (the shipped `store.fork()`), the notes line goes, the tree feet gain Reset.
- **Copy for the game / Send**: the planner's states (`sendActionFor`, every state in `web-planner-states`).

## 10. Copy (verbatim)

### 10.1 From the committed modules, unchanged
- /premium: `plansPageCopy` (`title`, `band.lede`, `band.guildLead`, `band.guildLink`, `founder.label`, `founder.open`, `interval.*`, `free.*`, `premium.*` with the feature anchors, `guild.*`, `raidGuild.*`, `price.*` through `priceCell`, `buttons.*` through `buttonLabel`, `trial` through `trialParts`, `coachSample.*`, `historyPromise`, `howItWorks`, `rules`, `foot`, `faq`), `premiumSourceCopy.premiumCard.*` and `buttons.*`, `checkoutCopy.canceled`, `billingBlockCopy.pastDueTag`, `pastDueBanner`, `accountPlanCopy.seatEnded.*`, `checkoutGuildCopy.onPlan.youPay`, `billedBy`, `manageBilling`.
- /premium/checkout: `checkoutCopy` (`redirecting`, `processing`, `confirmed`, `confirmationSlow`, `genericFailure`, `notOpenYet`, `comebackGone`, `retry`, `whichGuild`), `checkoutGuildCopy` (band, planSummary, sizeLine, lead, facts, ready, askLeader, copyField, foot, invalidLink and every card string), `premiumSourceCopy.panelTitle` through `premiumOwnedView`.
- /account: `accountPageCopy` (`title`, `signedInWithBattlenet`, `signedInByEmail`, `signOut`, `planLabel`, `guildLabel`, `devicesLabel`, `youLabel`, `noDevices`, `downloadCompanion`, `pairADevice`, `pseudonymLabel`, `pseudonymNote`, `signedInBanner`), `characterListCopy` (every key), `devicesCopy`, `accountSignInCopy.reason`, `accountPlanCopy` (head, planLine, founder, offer, seatEnded, guildRow), `premiumBlockCopy.rows`, `discordLinkCopy`, `discordRoleCopy`, `notificationsCopy`, `premiumRowCopy`, `guildConsentCopy`, `checkoutGuildCopy.claimPendingYours`, `settingsWords.guildSettingsLink`, `ACCOUNT_FAILED`, `addonCopy.setupOpenCharacter`, `setupBis` (the head answer).
- /b/:id: `plannerCopy.sharedLinkNote`, `fork`, `READ_ONLY_REASON`, every planner key the planner spec lists; `handler.go`'s `Build not found` / `That link does not point at a saved build. It may have been mistyped.` and `Build unavailable` / `That build could not be loaded just now.`; `page.go`'s `Open the build planner`; `b-unavailable.astro`'s two sentences.

### 10.2 Proposed (the generator's `PROPOSED`)
| Key | Text | Where |
|---|---|---|
| `plansPageCopy.yourPlanTag` | `Your plan` | the Tag on the account's plan column |
| `plansPageCopy.faqTitle` | `Questions` | Section title (today a literal in `premium.astro`) |
| `checkoutCopy.steps` | `Sign in`, `Pay with Stripe`, `Premium on` | the checkout Steps |
| `checkoutCopy.signedInAs` | `Signed in as {battletag}` | row 1 |
| `checkoutCopy.stripeLine` | `Stripe handles the card; we never see it.` | row 2 while redirecting (the FAQ's own fact) |
| `checkoutCopy.paid` | `Stripe took the payment.` | row 2 after the return |
| `checkoutCopy.checkAgain` | `Check again` | the slow state's button |
| `checkoutCopy.nextSim`, `nextCoach` | `Run a server sim`, `Upload a log for the coach` | the confirmed state's links |
| `accountPageCopy.reportsLink` | `Your reports` | under the Characters Plate |
| `accountPageCopy.deleteTitle`, `deleteBody` | `Delete your account`; `Email help@foreversixty.gg from the address on your Battle.net account and we delete it with everything on it. Reports you made public stay anonymous.` | You (needs the owner's check of the last sentence against the privacy page) |
| `accountPageCopy.signOutAsk` | `Sign out of Forever Sixty on this browser?` | the sign-out confirm |
| `accountPlanCopy.guildRow.claimPending` | `Claim pending` | the guild Tag |
| `CharacterRowLink` label | `Open` | the row's link (with `›`) |
| page.go `sharedMeta`, `views` | `Shared build`, `{n} views` | /b/:id meta |
| page.go `guidesLink` | `Talent builds in the class guides` | the message pages' second link (to /guides) |
Retired: `plansPageCopy.band.eyebrow`, `band.heading` (as the H1), `accountPageCopy.eyebrow`, `currentCharacterLabel`, `heroLogs`, `noExportTitle` and `noExportBody` (the character page owns the no-export state), `reportsLabel` (as a card), `yourRatingsLabel`, `checkoutGuildCopy.band.label`.

## 11. What each block reads, gaps, bugs

| Shown | Reads | Exists? |
|---|---|---|
| Prices, founder lines, per-member and per-raider | `plans.ts`, `plan-costs.ts` | yes |
| The account's plan on /premium | `/v1/me` entitlements, seat, ended_seat, guilds[].plan | yes (an island /premium does not have today: G1) |
| Checkout progress | the launcher's own states, `/v1/me` poll | yes |
| Characters, main, source, guild | `/v1/me` characters, `main_character_key`, `build.source` | yes |
| Plan, offer, ended seat, past due | `/v1/me` entitlements.billing, seat_offer, ended_seat | yes |
| Guild plan, claim pending, verified, next step | `/v1/me` guilds[].plan, claim, subscribe, verified, can | yes |
| Devices with platform words | `/v1/devices` + `platformLabel` | yes (B2) |
| Discord link, role, notifications | `/v1/me` discord_link, notifications | yes |
| Delete the account | `DELETE /v1/me` | yes (G6 closed 2026-11-07) |
| Shared build title and views | the record (`title`, `views`) | yes; the page does not pass them to the island (G4) |
| The shared build in the site column | `page.go` template | no column today (B4) |

Gaps:
- **G1. /premium has no session read.** The owned marks come from the `fs.premiumSource` hint only; the head answer, the officer's guild column, past due and ended seat need `/v1/me`. A `PremiumAccount` island (`client:idle`, the shared `/v1/me` cache) fills them; signed out it never hydrates.
- **G2. Seat ended on /premium** reads `ended_seat.resubscribe_url` from `/v1/me` (exists); the page has no place for it today.
- **G3. The checkout's paid row** asserts a payment the page only infers from `?status=success`. Stripe returns `success_url` only after a paid session, so the line is true; if the owner prefers, row 2 can read `Back from Stripe.` (question 4).
- **G4. `/b/:id`'s island props** gain `title` and `views` (already in the record) and `hideBar: true`; `page.go` passes the record's title to the template.
- **G5. The coach's last-raid line** shows a failure sentence when its read fails; it should draw nothing (not a load error the player can act on).
- **G6. No self-serve account deletion.** **Closed 2026-11-07 by the account-delete lane:** `DELETE /v1/me` (operation `deleteAccount`) with a typed confirm. The You row that said to email is gone; a `Delete account` Section sits last in the rail (one line, a danger button) and opens a confirm dialog listing what is deleted and what is kept, with the BattleTag (or `delete`) typed out (`AccountDelete.svelte`, `AccountDeleteDialog.svelte`, copy `accountPageCopy.deleteSection` / `deleteDialog`; mock `web-account-delete`, parity state `account-delete`). Email stays the route for export and correction, and an alternative for deletion (privacy page).
- **G7. A parity fixture for /b/:id**: the API page's HTML as a fixture (planner spec G6), now with this template.

Bugs for the build:
- **B1.** The current character is drawn twice on /account (hero and row).
- **B2.** `Account.svelte` prints `device.platform` raw (`windows/amd64`); use `deviceStatusLine`.
- **B3.** /premium's H1 is the slogan, not the page's name; `Founder` shows three times.
- **B4.** `page.go`'s template has no page column: /b/:id runs edge to edge from 1440 up.
- **B5.** Checkout's success and guild step columns are centred under a left-aligned head (three left edges).

## 12. Phone and desktop

- **Never collapses:** the Premium price and Subscribe (first in the Plate below 1024), the founder line, every state line, the account's characters, the plan's name and date, every confirm, `Fork` and `Copy for the game` on /b/:id.
- **Collapses:** the interval bar to its summary line below 768; the plan columns stack below 1024 (Premium first); guild columns stack; the rules and questions grids to one column below 768; the head answer under the title below 900; the rail follows main below 1024 (/premium: the coach sample after Questions; /account: Devices, Discord, You after Guild); a Row's trail drops under its text below 600.
- **Widths:** 1024 to 1759 main + 340 rail in a 1,344 column; from 1760 main + 400 rail, the column 1,792 at 2000 (title, Plate and bar share x 168 at 2000; live: x 376 on /premium and /account, x 0 on /b/:id). /premium/checkout's form column is 560 at the page's left edge; the guild list uses the full column (`ledger`).

## 13. Performance and polish

- /premium stays static and zero-JS signed out; the `PremiumAccount` island is `client:idle` and reserves the answer's box when `data-session` is set; the `Your plan` marks swap inside fixed cells (no shift).
- /account keeps the shipped reserve (`account-left-col`, `account-rail-col` heights stored per account) under the new column split; the reserve keys move with the sections.
- Checkout: the Plate is server-rendered with the three rows in their pre-hydration states (row 2 busy) so nothing moves when the launcher reads the URL.
- /b/:id: the trees' art paints in the server render (the planner's reserve, 575/543/547/615 px); the API page loads Marcellus (the H1) and Barlow from `planner-island.css`.
- Lighthouse (`web/lighthouserc.json` budgets unchanged): /premium, /account performance ≥ 0.95 at 360 px 4x CPU, CLS ≤ 0.05; /b/:id is not in the run (API-rendered; G7).
- The guard: `premium.astro`, `PlanCard`, `PriceBlock`, `IntervalControl`, `GuildRules`, `PremiumCoachSample`, `CheckoutLauncher`, `components/checkout/*`, `Account.svelte`, `components/account/*` leave `surface-guard.allowlist.json`. `plates.spec.ts` gains /premium, /account and /premium/checkout (one Plate each).

## 14. Parity regions and states

/premium: `head`, `head-title`, `head-answer` (signed in), `controls`, `notes`, `plate`, `section-for-guild-leaders`, `section-how-the-guild-plans-work`, `section-questions`, `section-what-the-coach-says-after-a-raid`. Checkout: `head`, `head-title`, `head-answer`, `plate`. /account: `head`, `head-title`, `head-meta`, `head-answer`, `section-characters`, `plate`, `section-plan`, `section-guild`, `section-devices`, `section-discord`, `section-you`. /b/:id: `head`, `head-title`, `head-meta`, `head-answer`, `notes`, `plate`, `section-point-order`, `section-in-the-game`, `section-level-60-build`. Parity states to start with: `premium`, `premium-own`, `checkout-confirmed`, `checkout-guild`, `account`, `account-premium`, `account-new`, at 1440 and 390 (fixtures: the home fixtures with the DERIVED fields written as files under `web/tests/e2e/fixtures/account/`).

## 15. Acceptance screenshots the implementer attaches

At 390, 1024, 1440, 1920 and 2000, full page and first screen, each beside its mock at the same scale, the 2000 capture first:
1. /premium in every state of section 5 (10), the interval switched, the founder window closed (clock past 2027-01-07T08:00Z).
2. /premium/checkout in every state of section 6 (9), the guild step's other cards (claim, start claim, confirm, on plan, frozen, checkout open, empty, load error).
3. /account in every state of section 7A (15), plus each Discord block state and the guild leave confirm for a claim holder.
4. /b/:id in every state of section 8 (7), after Fork, and a press on a cell (the refusal line).
5. Focus-visible on Subscribe, the interval control, a checkout step action, Set as main, Accept the seat, Leave guild, Fork (1440); hover on a plan button and a character row.
6. The parity side-by-sides of section 14, Lighthouse for /premium and /account, the guard diff.

## 16. Decisions (proposed; the coordinator rules after the wow-player round)

- **PR1. /premium's H1 is `Plans`**, the slogan heading retires and its lede becomes the head's lead. Recommended (tenet 14: no slogan hero). Alternative: keep `What is free, and what the plans add` in Marcellus.
- **PR2. One founder line** under the bar instead of three founder marks. Recommended.
- **PR3. The player's two plans are the one Plate; Premium leads on a phone.** Recommended (the first Subscribe moves from y 932 to y 602 at 390). Alternative: Free first, as the "free stays complete" promise reads.
- **PR4. Questions and the guild rules are open rows**, no disclosures. Recommended (tenet 4). Cost: /premium grows to 6,902 px at 390.
- **PR5. Signed in, /premium's head answers with the account's plan** and marks `Your plan`; an officer whose guild is on a plan sees it on that plan's column with Manage billing. Recommended (needs G1).
- **K1. Checkout is a form-column page with `Steps`** (Sign in, Pay with Stripe, Premium on) instead of a band over one line. Recommended.
- **K2. After a confirmed payment, the next step** is three links (a server sim, a log for the coach, your plan). Recommended (tenet 14).
- **A1. /account: Characters, Plan and Guild in main; Devices, Discord and You in the rail.** Recommended (the plan was below the fold in a rail beside an empty column).
- **A2. The hero card retires; the head answer is the main character.** Recommended (B1; one crest per character on the page).
- **A3. `Your reports` and `Your ratings` leave /account** for one link to /reports; the character page owns ratings. Recommended (one job per page).
- **A4. One Discord section** (link, Premium role, notifications). Recommended.
- **A5. Sign out sits in the head's meta and asks first.** Recommended. Alternative: only in the header's account menu.
- **A6. A `Delete your account` row** that says how today (email) until G6. Recommended.
- **B1. /b/:id's title is the build's own name** in the class colour. Recommended.
- **B2. `Fork` moves into the first screen** beside the shared-link note. Recommended.
- **B3. No class and race bar on /b/:id.** Recommended. Alternative: the planner's bar, disabled (planner spec 5.1).
- **P1. Approve the primitives** `Columns`, `Button(tone='danger')`, `PlanHead`.

## 17. Questions for the wow-player persona

1. `web-premium` at 390 and 1440: does the page answer "what does Premium add and cost" before you scroll? Should Free or Premium lead on a phone (PR3)?
2. `web-premium-own` and `web-premium-free`: is `Plan · Premium · Renews 4 Nov 2027 · $40 a year` in the head what you want to see when you come to this page already paying? Is `Your plan` on the column enough?
3. `web-premium`: is the coach sample beside the plans convincing evidence, or would you rather see a real server-sim example (the owner dropped the empty one on 2026-10-28)?
4. `web-checkout-processing` and `-confirmed`: do the three steps make it clear whether you were charged? `Stripe took the payment.` or `Back from Stripe.` (G3)?
5. `web-checkout-confirmed`: are `Run a server sim`, `Upload a log for the coach`, `See your plan` the right next steps after paying?
6. `web-account-premium` at 1440: characters and plan in the main column, connections in the rail: is anything you look for on an account page missing or in the wrong place?
7. `web-account-offer`: is the seat offer clear about what happens to the Premium you pay for?
8. `web-account-claim`: does `Claim pending` with the sentence and `Copy link` tell you what to do next?
9. `web-account`: is losing the `Your reports` and `Your ratings` cards from this page fine with `Your reports ›` under your characters (A3)?
10. `web-b` at 1440 and 390, as someone who clicked a link in Discord: do you see whose build it is, what it does and how to put it in game? Is `Fork` the right word for "make this mine to edit"?
11. `web-b-holy`: signed in on a hunter, a priest build: does the page say clearly that Send does not apply?
12. `web-b-missing`: what would you click next?
