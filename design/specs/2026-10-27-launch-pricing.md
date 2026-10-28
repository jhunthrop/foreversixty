# Launch pricing: plans page, guild billing, seats, account (experience spec)

Author: ux-designer, 2026-10-27. Product design (binding): `docs/superpowers/specs/2026-10-27-pricing-design.md` (prices §3, rules §4, builder decisions §8; chip-in is deferred and is not designed here). Experience rules (binding): `docs/superpowers/specs/2026-10-24-premium-experience-design.md` §2 anti-goals and §5 (no upsell interrupts, one honest sentence, `/premium` lists only what is live). Surfaces as they are today: `web/src/pages/premium.astro` (and `/pricing`, a 301 to it), `web/src/components/guild/GuildBillingSection.svelte`, `web/src/components/account/AccountGuildPlanLine.svelte`, the account page's `Plan` `StatePanel` with `AccountPremiumBlock`. Read-only grammar: `design/specs/2026-10-25-officer-tools-gate.md`. Mock: `design/mocks/gen_launch_pricing.py` writes `design/mocks/renders/launch-pricing-*.html` (and `.png` through `render.py`); it imports G1's helpers (`gen_guild_raids`) so chips, buttons and the guild band cannot drift. No new token. No new shared control (tenet 10): every control below is a recipe that exists today, named where it is used.

Prices are never typed in copy. `web/src/lib/billing/plans.ts` gains the plan `raid_guild`, the founder prices and the founder window; every string in §6 that shows money is a function of those constants. The mock does the same (`PRICES` in the generator).

## 1. The one question, and the first screen

| Surface, viewer | Question | Answer visible with no scroll, 360 x 640 and 1280 x 800 |
|---|---|---|
| `/premium`, anyone | "What does each plan cost, and what stays free?" | The band's promise line, the `Yearly / Monthly` control, and the Premium card's price, founder line and `Subscribe` button. 360: the Premium button ends above y 600 (the Free card is a one-row summary above it, §3.1.3). 1280: all four cards' names, prices and buttons end above y 520, and the first four feature lines of each card are visible. |
| Guild settings `#billing`, the subscriber or the guild leader | "What is my guild on, what will it cost next, and who holds the seats?" | The plan line (`Guild · renews 18 Nov 2027 · $120 a year` + `Founder`), the size line when over 25, and the first seat row. 360: the first seat row ends above y 640 when the page opens at `#billing`. 1280: the whole Billing section down to the last seat row. |
| Guild settings `#billing`, an officer who may not bill | "Is the guild covered, and by whom?" | The plan line and `Billed by Kraggor`. Seats read-only. |
| `/account`, a paying member offered a seat | "What happens to my own subscription if I take it?" | The offer block: who offers, the date their own Premium ends, the price they keep a right to, `Accept the seat` and `Decline`. 360 and 1280: the whole block is the first thing in the Plan panel and ends above the fold when the page opens at `#plan`. |
| `/account`, a seat holder | "Am I on Premium, and until when?" | `Premium · a seat from Olympus XXVII` and its one-line why. |

## 2. Reference, and what we borrow

- **The in-game Shop's Game Time panel** (WoW's own subscription screen). Borrowed: the longest term first, its per-month equivalent under the price, one button per offer, no feature pitch on the purchase step. Not borrowed: the art splash.
- **Warcraft Logs' subscription page** (the tier columns raiders already compare). Borrowed: tiers side by side, the price at the top of each column, feature lines under it, one button per column, an "Everything in Guild" line instead of repeating rows. Not borrowed: a "most popular" highlight. No column is pushed; the price figure is each card's primary element (§5).
- **Raid-Helper's and RCLootCouncil's officer settings** (where an officer already manages who has what). Borrowed: a short list of named slots, each with its holder, a change control and the date it can next change.
- **Our guild settings** (`GuildSettings.svelte`): `section-title` h2s, 13 px muted helper lines, `SECONDARY_BUTTON_FIXED` and `PRIMARY_BUTTON_FIXED`, the native `<select>` with `SelectChevron` (the visibility select), the in-place confirm block (`border-line-soft border p-4`, the claim-contest confirm), `role="alert"` error lines.
- **Our account page**: the `Plan` `StatePanel` and `AccountPremiumBlock`'s key/value rows (10 rem muted label, value beside it).
- **Our G1 segmented control** (`SignupSegments`: 44 px segments, `--card-top` on, 2 px inset underline) for the interval control, and the `pill-site` pill (G1 `pill(text, 'site')`) for `Founder`.

## 3. Layout, region by region

Tokens: `--bg #07090d`, `--raised #0d111a`, `--card-top #131824`, `--line #262e40`, `--line-soft #1c2230`, `--line-warm #3a3326`, `--line-warm-strong #4a4030`, `--text #e9e4d8`, `--strong #f2eee4`, `--muted #9a9484`, `--gold #e5b955`, `--gold-hover #f5d27a`, `--kill #7fd48a`, `--wipe #f0736b`. Type: `section-title` 18 px Cinzel uppercase; card titles Cinzel 700 17 px; prices JetBrains Mono 500, 32 px desktop and 28 px phone, tabular; body Barlow 14–15 px; secondary 13 px; `label` 11 px. Gutter 48 / 18. Section gap 32 / 22. Content max 1344 px, centred at every width up to 2000 and beyond.

### 3.1 `/premium`, the plans page (`/pricing` keeps its 301 to it)

Static Astro with zero client JavaScript files, as today. Two `is:inline` scripts only. The first is today's `?canceled=1` script. The second is a six-line head script that sets `data-founder` on the root element to `before`, `open` or `closed` by comparing the clock with the window constants (R-P4), before first paint. A page built on 6 January and read on 7 January is therefore right, and nothing shifts. The `measure-reading` column is replaced by the 1344 px content column.

**3.1.1 Band.** Full bleed (the guild page's band treatment, never a card in the gutters). Inside the 1344 column; padding 28 px top desktop, 22 phone; 24 px bottom.
- `label` `PLANS` in `--gold`.
- h1, Cinzel 700 22 px `--strong`: `What is free, and what the plans add`.
- Promise line, 15 px `--text`, max 720 px (§6.1).
- Founder row, only when `data-founder` is `before` or `open`, 13 px `--text`: `pill-site` `Founder`, 8 px, the line (§6.1), dates in mono.
- Interval control. From 1024 it sits right-aligned on the h1's row; below 1024 it is its own row under the founder row. A radiogroup labelled `Billing interval`: two native radios styled with the G1 segment recipe, 44 px tall at every width, 112 px per segment on desktop, 50 % each on phone (full width). `Yearly` is first and checked by default. On: `--card-top` ground, `--strong` text, a 2 px inset gold underline. CSS `:checked` decides which price set shows. Both sets live in the DOM inside the same fixed-height price cells, so the two states are pixel-identical in height.

**3.1.2 Cards.** From 1024: four columns (`grid-cols-4`, gap 16). From 768 to 1023: two columns, Free and Premium first. Below 768: §3.1.3. Each card is the `card` recipe (`--raised`, 1 px `--line`, 6 px radius, 1 px top highlight `rgba(229,185,85,.35)`), padding 20, a flex column. Each region has a fixed height on desktop, so the four prices and the four buttons sit on one line each.
1. Name, Cinzel 700 17 px `--strong` (24 px line). Under it, 13 px muted, one line saying who it is for (§6.1).
2. Price block, 96 px tall desktop, 88 phone. Line 1: the price in mono 32 px `--strong` (`$40`), then ` /year` 15 px muted. When founder is open, the list price follows 8 px later, struck through, mono 15 px muted (`$50`). Line 2, 13 px muted: the per-month equivalent (yearly) or `Or $40 a year.` (monthly). Line 3, founder open only, 12 px `--gold`: `Founder price, kept while you stay subscribed.` Free shows `$0` and `No card, no trial, no end date.`
3. Button, 44 px, full card width, label centred. Premium: `Subscribe · $40 a year`. Guild and Raid guild: `Subscribe a guild · $120 a year`. Recipe: `SECONDARY_BUTTON_FIXED border-gold text-gold w-full justify-center` (outline gold; never filled, because no column is pushed). Free has no button; its slot holds the 13 px muted line `Sign in with Battle.net from the top of any page.` at the same 44 px, so the feature lists align.
4. Trial line, guild cards only, 13 px `--text`, 40 px slot, with `First 7 days free` in `--strong` 600 (§6.1). Free and Premium keep the 40 px blank so the lists align.
5. Features: a list, 14 px `--text`, 32 px rows divided on `--line-soft`, each led by a 14 px `--kill` tick (the planner's tick glyph, inline SVG, hidden from assistive tech). Raid guild's first line, `Everything in Guild`, is `--strong`. Free's last line is the promise, 13 px `--strong` 600, no tick, 12 px above: `This list never shrinks.` A Premium benefit that is not live renders with ` · ships {Month}` muted after it, or not at all (premium spec §5).

**3.1.3 Phone (below 768).** Order: band; the **Free summary row**; Premium; Guild; Raid guild; the history promise; the **Free card** in full (`id="free"`). The summary row is one `card`, 64 px, the whole row a link to `#free`: `Free` Cinzel 15 px, `$0` mono 20 px at the right, and under the name, 13 px muted, `Everything a player needs. This list never shrinks.` It carries sr-only `See the free list`. The summary keeps the free promise first and Premium's answer above the fold. Cards are full width, gap 16. Feature lists never go behind a disclosure (tenet 4).

**3.1.4 History promise**, under the guild cards. From 1024 it spans the Guild and Raid guild columns; below, full width. 14 px `--text`, a 2 px `--line` left rule, padding-left 14, max 720 (§6.1).

**3.1.5 How the guild plans work.** h2 `section-title`. From 1024 a grid `200px minmax(0,1fr)` of 44 px minimum rows on `--line-soft` dividers; below 1024 stacked (the term, then the 14 px sentence). The term is 12 px `--strong` 600 at every width, as the mock (amended (controller, 2026-10-27 persona round)). Eight rows, verbatim in §6.1: Size, Seats, Moving a seat, Already on Premium, Trial, Renewals, A new guild leader, Refunds. Never a disclosure: these are the rules an officer reads before putting the plan on a card.

**3.1.6 Foot.** One 13 px muted line (§6.1). Then today's FAQ list of disclosures (two edits in §6.1). Then today's no-ads line.

**3.1.7 Checkout guild chooser.** Today the checkout URL for a guild plan with no `guild_id` is invalid. It now shows, under the checkout page's h1, `Which guild?` (15 px `section-title`) and one 56 px row button per guild in `me.guilds` whose `can` holds `billing` and whose `plan` is null. Row: the guild crest at 32 px (the render `GuildCrestSettings` uses), the guild name 15 px `--strong`, `{n} verified members` 13 px muted; at the right, 13 px muted, the plan that fits (`Guild`, or above 25 `Raid guild fits {n} members`). Hover `--card-top`; focus-visible 2 px gold at −2 px. Choosing one continues to Stripe with that `guild_id` and the chosen plan. No billable guild: `You can subscribe a guild you lead or have claimed.` and the inline gold link `Find your guild` to `/guild`.

### 3.2 Guild settings `#billing` (`GuildBillingSection`, rebuilt)

h2 `Billing`, then the blocks below in order, 20 px apart. **Payers** are the subscriber and the guild leader (`can` holds `billing`; seat controls also need the new word `seats`). Every other officer sees the plan line, `Billed by`, the size line and the seat holders, read-only, with no buttons, plus Records. Members never reach settings.

**3.2.1 Plan line** (14 px `--text`), then line 2 (13 px muted). States, verbatim in §6.2:
- Not on a plan: the chooser (3.2.2) and nothing else above Records.
- Trial: `Guild · trial · first charge 11 Dec 2026: $120`, then `Founder` when founder. Line 2: `Cancel before 11 Dec and nothing is charged. Billed to you.`
- Active: `Guild · renews 18 Nov 2027 · $120 a year`, then `Founder`. Line 2: `Billed to you.` or `Billed by Kraggor`.
- Ending (cancel at period end): `Guild · ends 18 Nov 2027`. Line 2: `Officer tools turn read-only then. Records stay, and CSV export stays.`
- Past due: today's `pastDueBanner` as an alert, 13 px `--strong`, under line 2.
- Lapsed: `Guild plan ended 18 Nov 2027.` Line 2: `Officer tools are read-only. Raids, attendance, loot and reports stay readable, and the export below always works.` Then the chooser with list prices and the button `Subscribe again`.
- Officer-tools grant (today's scope): today's line and explanation unchanged, then the chooser.

Payers get, on the plan line's row (right-aligned from 768, under line 2 below): `Manage billing` (`SECONDARY_BUTTON_FIXED`, the Stripe portal as today) and, on Guild only, `Move to Raid guild` (`SECONDARY_BUTTON_FIXED border-gold text-gold`). When the size line shows, `Move to Raid guild` moves there and is not drawn here.

**3.2.2 Chooser** (not on a plan, lapsed, grant). A `Yearly / Monthly` control first (G1 segment recipe, 44 px, `Yearly` on). Then two option cards side by side from 768, stacked below. Each is a label wrapping a native radio, min 96 px, padding 16: name Cinzel 15 px, price mono 20 px for the chosen interval (founder price with the struck list price while open), and one 13 px muted line (§6.2). The plan that fits the verified count is preselected (Raid guild above 25). Selected: 1 px `--gold` border on `--card-top`. Under the cards, 13 px muted: `28 verified members today.` Then the button, `PRIMARY_BUTTON_FIXED px-4` (the section's one fill): `Start the 7-day trial`, or `Subscribe · $120 a year` when the guild has used its trial (and `Subscribe again` when lapsed). The trial line follows unless the trial is used; the founder line follows while open. The button opens Checkout with the plan, the interval and the guild id.

**3.2.3 Size line** (Guild tier, verified count above 25): 14 px `--text` on a 2 px `--gold` left rule, padding-left 14 (§6.2). Payers get `Move to Raid guild` beside it from 768, under it below. At 25 or fewer there is no line. The count is the daily count, leader included (product §8.4).

**3.2.4 Move to Raid guild, confirm**, in place under the button (the claim-contest confirm block: `border-line-soft border p-4`, a flex column with 12 px gap). The 14 px sentence carries today's prorated amount from the preview and the next renewal price (founder lands on the founder Raid guild price, product §8.2). While the preview loads, a 64 x 14 px `Skeleton` stands in the amount's place and the commit button is disabled. Buttons: `Move to Raid guild` (`PRIMARY_BUTTON_FIXED`) and `Keep Guild` (44 px underlined text button), focus on `Keep Guild`. Done: the plan line updates and a 13 px muted status line confirms (§6.2). Failure: an alert line (§6.2).

**3.2.5 Seats.** Sub-head `Premium seats` (15 px `section-title`) with its mono muted count `2 of 3 given`. Helper, 13 px muted (§6.2). One row per seat, 64 px desktop, divided on `--line-soft`, grid `minmax(220px,1fr) minmax(0,1.4fr) auto`, gap 16.
- **Held**: `ClassCrestRing` 32, the name in Cinzel 14 px class colour, `Given 20 Nov` 12 px muted under it. Middle, 13 px: `Can move from 20 Dec` (muted) or `Can move now` (`--text`). Right: `Move seat` (`SECONDARY_BUTTON_FIXED`), aria-disabled at 50 % opacity while locked, described by the middle text.
- **Offered, waiting** (the member pays for Premium): crest, name, `pill-sample` `Waiting`. Middle: `Offered 8 Dec. They pay for Premium, so they accept it first.` Right: `Withdraw` (44 px underlined text button). Withdrawing never starts a lock. No lock date and no `Move seat`.
- **Held, moving to a payer** (amended (controller, 2026-10-27 persona round)): line 1 the holder with `Given {d MMM}`. Middle: `Moving to Mirelle · offered 10 Nov` with `pill-sample` `Waiting`. Right: only `Withdraw`. The count `{given} of {total}` is out of the plan's seat count.
- **Accepted, starts later**: crest, name. Middle: `Accepted · starts 22 Dec, when their own Premium ends`. Right: `Move seat`, locked as Held (the lock runs from acceptance, R-P3).
- **Empty**: a 32 px dashed `--line` circle, `Not given` 14 px muted. Right: `Give seat` (`SECONDARY_BUTTON_FIXED border-gold text-gold`).
- **Returned**: the Empty row with the middle `Free since 2 Dec: Sagebloom left the guild.` for 14 days.

**Give and Move** open in place under the row (the confirm block). `label` `Give this seat to` (or `Move this seat to`), then a native select with `SelectChevron` (44 px, the visibility select's recipe). Options: verified members not holding a seat here, sorted by name, each `Name · Class`; a member who pays for Premium reads `Name · Class · pays for Premium`. When such a member is chosen, 13 px muted under the select: the payer note (§6.2). When moving, 13 px muted: `Lightbrand keeps Premium until the seat moves.` Buttons: `Give seat` or `Move seat` (`PRIMARY_BUTTON_FIXED`) and `Cancel` (text button); focus lands on the select. Officers without `seats` see the rows with no right column.

**3.2.6 Who pays** (sub-head; the subscriber only). 13 px muted helper (§6.2). `Transfer the plan` (`SECONDARY_BUTTON_FIXED`) opens the confirm block: `label` `Transfer to`, a native select of the guild's verified officers who have a Forever Sixty account, the consequence sentence (14 px), then `Send the offer` (`PRIMARY_BUTTON_FIXED`) and `Cancel`. **Offer sent**: the button's place shows `Offered to Bow Jackzon on 8 Dec. Waiting for them.` and `Withdraw`. **Declined**: `Bow Jackzon declined.` muted, and `Transfer the plan` returns. **The receiving officer** sees at the top of `#billing` the confirm block with a 2 px gold left rule: the offer sentence, `Take over the plan` (`PRIMARY_BUTTON_FIXED`, to Checkout in setup mode, first charge at the current period end) and `Not now` (text button; it declines). **Accepted**: the subscriber's plan line reads `Guild · ends 18 Nov 2027 · then billed by Bow Jackzon`.

**3.2.7 Records** (every officer, every state, lapsed included). Sub-head `Records`, a 13 px muted line, then two `SECONDARY_BUTTON_FIXED`: `Download attendance (CSV)` and `Download loot (CSV)`, side by side from 768, stacked full width below.

**3.2.8 Gate sentences on a Guild-tier guild** (officer-tools gate §3.2 grammar, `GuildPlanGate`): the Loot and Readiness tabs read `Loot council and the readiness board are part of Raid guild. ` + `Move to Raid guild` (link to `#billing`). A lapsed or never-subscribed guild keeps today's sentence.

### 3.3 Account page, the `Plan` panel (`id="plan"`)

The panel keeps its `StatePanel` and `AccountPremiumBlock`. Order inside: the offer (when one), the plan line, line 2, the premium rows, the button.
- **Seat offer**, top: the confirm block with a 2 px `--gold` left rule. Line 1, 15 px `--strong`: `Olympus XXVII offers you a Premium seat.` Line 2, 13 px muted: `From Kraggor, 8 Dec.` Line 3, 14 px `--text`: the consequence sentence (§6.3), dates and price in mono. Buttons: `Accept the seat` (`PRIMARY_BUTTON_FIXED`, the page's one fill while the offer stands) and `Decline` (44 px underlined text button). Neither is auto-focused. Accepted: the block becomes a 13 px muted status line and the plan line takes its "ends, then a seat" state. Declined: a status line; the block is gone on reload.
- **A member who does not pay** never sees an offer: a given seat is live at once (R-P2).
- **Plan line**, 14 px: own Premium `Premium · renews 22 Dec 2026 · $5 a month`, then `Founder` when founder; own Premium ending into a seat `Premium · ends 22 Dec 2026 · then a seat from Olympus XXVII`; a seat `Premium · a seat from Olympus XXVII`, with line 2, 13 px muted, `Given by Kraggor. It lasts while the guild's plan does; nothing is billed to you.`; free unchanged (`Plan · Not on Premium.` `See plans`). `Manage billing` shows only with an own subscription, as today.
- **Founder pill**: `pill-site` `Founder`, 8 px after the plan line, with a title and an sr-only tail (§6.3). Account and guild settings only, never a public surface (R-P6).
- **Seat ended** (30 days, only for a member who cancelled their own Premium for that seat): the plan line's place reads the ended sentence with the inline gold link `Come back at $5 a month` (the email's link).
- **Guild rows** (`AccountGuildPlanLine`): `Guild plan` or `Raid guild plan`, then today's verb and date, then ` · you hold a seat` where the account holds one.

### 3.4 Emails

Sender `Forever Sixty`, from help@foreversixty.gg. Plain text first, the site's voice, no exclamation marks. Each email is a subject, a first sentence that carries the whole message, then one link (§6.4). Five emails: renewal (7 days before a yearly renewal), size (14 days before renewal, Guild tier, count above 25 on its send day), seat offered (to a payer), seat given (its twin, to a non-payer), seat removed (to someone who cancelled their own Premium for that seat; its link opens Checkout at their old price and expires 30 days after the email).

## 4. States

| Element | Default | Hover | Focus-visible | Active | Disabled | Loading | Empty | Error |
|---|---|---|---|---|---|---|---|---|
| Interval control (plans page, chooser) | `Yearly` on | off segment `--text` on `--card-top` | 2 px gold ring on the checked segment, 2 px offset | on look | n/a | n/a (static) | n/a | n/a |
| Card `Subscribe` | outline gold | `--gold-hover` border and text, `rgba(229,185,85,.08)` ground | 2 px gold, 2 px offset | `--gold-hover` | n/a | n/a (it navigates) | n/a | the checkout page's own states |
| Free summary row (phone) | `card` | `--card-top` | 2 px gold, −2 px | `#171d2b` | n/a | n/a | n/a | n/a |
| Checkout guild row | row | `--card-top` | 2 px gold, −2 px | `#171d2b` | n/a | 3 `Skeleton` rows of 56 px | the none line (§3.1.7) | `Your guilds did not load.` + `Try again` |
| Chooser option | `--line` border | `--line-warm-strong` border | 2 px gold | gold border on `--card-top` | n/a | 2 `Skeleton` cards of 96 px | n/a | n/a |
| `Start the 7-day trial`, `Subscribe`, `Subscribe again` | gold fill | `--gold-hover` | 2 px `--strong` ring, 2 px offset | `--gold-hover` | while redirecting | `Opening checkout…`, busy | n/a | `Checkout did not open. Try again.` (alert) |
| `Manage billing` | today | today | today | today | busy | busy | n/a | today's `billingUnavailable` + `Try again` |
| `Move to Raid guild` and its confirm | outline gold | as the card button | 2 px gold | `--gold-hover` | commit disabled until the preview lands | amount `Skeleton`; on press `Moving…` | n/a | `That did not go through. Nothing was charged. Try again.` |
| `Give seat`, `Move seat` | outline (gold for Give) | `--line-warm-strong` | 2 px gold | `--card-top` | locked: 50 % opacity, aria-disabled, the date in text | `Giving…`, `Moving…` | no member left: select disabled + `Every verified member already has Premium from this guild.` | the row's alert line (§6.2) |
| Text buttons (`Withdraw`, `Cancel`, `Keep Guild`, `Decline`, `Not now`) | underlined `--text` | `--gold-hover` | 2 px gold | `--gold-hover` | busy | `Withdrawing…` | n/a | the block's alert line |
| `Transfer the plan`, `Send the offer`, `Take over the plan` | outline, fill, fill | as above | as above | as above | no other officer: disabled + `No other officer has a Forever Sixty account yet.` | `Sending…`, `Opening checkout…` | n/a | `The offer did not send. Try again.` |
| `Download … (CSV)` | outline | `--line-warm-strong` | 2 px gold | `--card-top` | n/a | `Preparing…`, busy | no rows: a header-only file still downloads | `The file did not download. Try again.` |
| `Accept the seat` | gold fill | `--gold-hover` | 2 px `--strong` | `--gold-hover` | busy | `Accepting…` | n/a | `That did not save. Try again.` |
| Billing section, whole | per state | n/a | n/a | n/a | n/a | reserves the height this guild's section last rendered at this layout (phone or from 768), else 520 px from 768 and 760 px on a phone (amended, persona round 2: the measured ready height of a paid plan with seats), with `Skeleton` lines until settings answer | n/a | today's `billingUnavailable` |

Touch and keyboard parity: every control is 44 px tall at every width (the `_FIXED` recipes; text buttons carry a 44 px min-height). Nothing is hover-only: a locked `Move seat` states its date in text, never in a tooltip. Radios move with the arrow keys. A confirm block moves focus to its safe choice (`Keep Guild`, `Cancel`) and Escape acts as the safe choice. Tab order in `#billing`: the receiver's offer, the plan line's buttons, the size line's button, the chooser, seats top to bottom, Who pays, Records.

## 5. Hierarchy

- Plans page: in each card the price figure is primary, the button is the action, the features are secondary. No column is emphasised and nothing on the page is filled. Gold appears once per card (the button), plus `Founder` in the band and the founder line in each price block.
- `#billing`: the plan line is primary. One fill at a time in the section: `Start the 7-day trial`, a confirm block's commit, or `Take over the plan`. Seat rows lead with the holder; dates and reasons are 12–13 px muted.
- Account: an open offer is the panel's primary element and the page's only fill.
- Copy states consequences, never urgency: no countdown, no "only N days left", no exclamation marks (premium spec §2).

## 6. Copy, verbatim

Strings live in `web/src/lib/billing/copy.ts` (`premiumCopy`, `checkoutCopy`, `billingBlockCopy`), `web/src/lib/guild/settings-view.ts`, and a new `web/src/lib/billing/seats-copy.ts`. `{price}` is `formatPrice` from `plans.ts` and reads `$40 a year` or `$5 a month` in running text (replacing today's `$4/month`); only the card price block uses ` /year` beside the figure. Dates are mono: `d MMM yyyy` when the date may fall in another year, `d MMM` otherwise; emails use `d MMMM yyyy`.

### 6.1 Plans page and checkout

| Where | String |
|---|---|
| Page | title `Plans`; description `What is free on Forever Sixty for good, what Premium, Guild and Raid guild add, and what each costs.` |
| Band | `PLANS`; `What is free, and what the plans add`; `Everything a player needs is free, and the free list never shrinks. Premium and the guild plans pay for the site. There are no ads for anyone.` |
| Founder row | open: `Founder` + `Founder prices until 6 January 2027. Buy by then and the price stays for as long as you stay subscribed.` (amended (controller, 2026-10-27 persona round): the year added, the date in the sentence's own font); before: `Founder` + `Founder prices from 4 November until 6 January.`; closed: no row |
| Interval | `Yearly`, `Monthly`; group label `Billing interval` |
| Free card | `Free`; `For every player`; `$0`; `No card, no trial, no end date.`; `Sign in with Battle.net from the top of any page.`; features `Addon export and import`, `BiS lists for every spec and level band`, `Talent planner`, `Browser sims`, `Logs, live logging and parses`, `Rankings and the tier list`, `Class guides`; `This list never shrinks.` |
| Free summary (phone) | `Free`; `$0`; `Everything a player needs. This list never shrinks.`; sr-only `See the free list` |
| Premium card | `Premium`; `For one player`; features, live ones only: `Server sims: 5,000 combinations, any precision, no tab left open`, `History charts of your score and parses`, `Discord notifications when a server sim or a log finishes`, `A supporter mark on your characters`; suffix for a benefit not yet live ` · ships {Month}` |
| Guild card | `Guild`; `For a guild of up to 25`; features `Up to 25 verified members`, `Raids and sign-ups, on the site and in Discord`, `Attendance and the raid-night report`, `The guild's Discord channel and bot`, `1 Premium seat to give to a member` |
| Raid guild card | `Raid guild`; `For a raiding guild of any size`; features `Everything in Guild`, `Any number of members`, `Loot council`, `The readiness board, with consumables`, `3 Premium seats to give to members` |
| Price block | ` /year`, ` /month`; the list price as plain text under the founder figure, `{list} after 6 Jan 2027` yearly or `{list} a month after 6 Jan 2027` monthly, no strike (owner-approved 2026-10-27, persona appeal review); yearly `{perMonth} a month, billed yearly` (cents only when not whole: `$3.33`, `$10`); monthly `Or {yearly price}.`; founder `Founder price, kept while you stay subscribed.` |
| Buttons | Premium `Subscribe · {price}`; guild cards `Subscribe a guild · {price}` |
| Trial | `First 7 days free, once per guild. Cancel before day 8 and nothing is charged.` |
| History promise | `If a guild plan ends, its raids, attendance, loot and reports stay readable for good, and officers can export attendance and loot as CSV at any time.` |
| How it works | h2 `How the guild plans work`. `Size`: `Guild counts verified members each day, the leader included. Above 25 the officer is told before renewal and offered Raid guild. Nothing switches off mid-plan.` `Seats`: `The subscriber or the guild leader gives each seat to any verified member. A seat holder has Premium while the plan lasts.` `Moving a seat`: `A seat can move once every 30 days. The old holder keeps Premium until it moves. If a holder leaves the guild, the seat comes back at once.` `Already on Premium`: `A member who pays for Premium accepts a seat first. Their own Premium stops renewing and the seat starts when it ends, so nobody pays twice. If the seat is taken away later, they get a link to come back at the price they had.` `Trial`: `A guild's first 7 days are free, once per guild. The card is added at the start; cancel before day 8 and nothing is charged.` `Renewals`: `We email 7 days before every yearly renewal. Cancel in one click; the plan runs to the end of the period you paid for.` `A new guild leader`: `The subscriber hands the plan to another officer from guild settings. The founder price goes with it.` `Refunds`: `Ask within 14 days of a first charge for a refund. After that there are no part-year refunds.` |
| Foot | `Prices in US dollars. Cancel any time; you keep the plan to the end of the period you paid for.` |
| FAQ edits | `What happens if I cancel?` now answers `You keep everything through the end of the period you already paid for. After that, the account works like a free one, and nothing is deleted. A guild's records stay readable for good.`; in the no-ads answers, `Premium and the guild plan` becomes `Premium and the guild plans` |
| Checkout chooser | `Which guild?`; `{n} verified members`; `Guild`; `Raid guild fits {n} members`; `You can subscribe a guild you lead or have claimed.`; `Find your guild`; `Your guilds did not load.`; `Try again`; added (controller, 2026-10-27 persona round), under the Which guild? heading: `{Plan} · {price}` |

### 6.2 Guild settings `#billing`

| Where | String |
|---|---|
| Plan line | `Guild · trial · first charge {date}: {amount}`; `Guild · renews {date} · {price}`; `Raid guild · renews {date} · {price}`; `Guild · ends {date}`; `Guild · ends {date} · then billed by {Name}`; `Guild plan ended {date}.`; pill `Founder` |
| Line 2 | `Cancel before {d MMM} and nothing is charged. Billed to you.`; `Billed to you.`; `Billed by {Name}`; `Officer tools turn read-only then. Records stay, and CSV export stays.`; `Officer tools are read-only. Raids, attendance, loot and reports stay readable, and the export below always works.` |
| Chooser | `Guild`; `Raid guild`; `Up to 25 verified members · 1 Premium seat`; `Any size · loot council and readiness · 3 Premium seats`; `{n} verified members today.`; `Start the 7-day trial`; `Subscribe · {price}`; `Subscribe again`; the trial line of §6.1; `Founder price until 6 January 2027, kept while the plan stays active.`; `Opening checkout…`; `Checkout did not open. Try again.` |
| Size line | `{n} verified members. Guild covers up to 25; nothing switches off. Raid guild covers any size.`; `Move to Raid guild` |
| Upgrade confirm | `Move to Raid guild? Today you pay {prorated} for the rest of this {year or month}. From {date} it is {price}.`; `Move to Raid guild` (`Moving…`); `Keep Guild`; `Moved. Loot council, readiness and 2 more seats are on now.`; `That did not go through. Nothing was charged. Try again.`; added (controller, 2026-10-27 persona round): `Moved. Stripe is still confirming the charge.`, `Stripe is still confirming the move. Reload in a minute to see the plan.`, `Today's amount was updated. Check it, then move.` |
| Seats | `Premium seats`; `{given} of {total} given`; `A seat gives one verified member Premium while the plan lasts. A seat can move once every 30 days.`; `Given {d MMM}`; `Can move from {d MMM}`; `Can move now`; `Move seat` (`Moving…`); `Waiting`; `Offered {d MMM}. They pay for Premium, so they accept it first.`; `Withdraw` (`Withdrawing…`); `Accepted · starts {d MMM}, when their own Premium ends`; `Not given`; `Give seat` (`Giving…`); `Free since {d MMM}: {Name} left the guild.`; added (controller, 2026-10-27 persona round): `Moving to {Name} · offered {d MMM}` |
| Seat picker | `Give this seat to`; `Move this seat to`; options `{Name} · {Class}` and `{Name} · {Class} · pays for Premium`; `They pay for Premium, so they get an offer to accept. Their own Premium stops renewing and the seat starts when it ends.`; `{Name} keeps Premium until the seat moves.`; `Give seat`; `Move seat`; `Cancel`; `Every verified member already has Premium from this guild.`; errors `That did not save. Try again.`, `That seat can move from {d MMM}.` (409 `seat_locked`), `{Name} is no longer a verified member.` (409 `not_member`) |
| Who pays | `Who pays`; `Hand the plan to another officer, for example when the guild leader changes. The founder price goes with it.`; `Transfer the plan`; `Transfer to`; `{Name} gets an offer to take over. When they add a card, they are billed from {date} at {price}, and your plan stops renewing. Nothing changes until they accept.`; `Send the offer` (`Sending…`); `Cancel`; `Offered to {Name} on {d MMM}. Waiting for them.`; `Withdraw`; `{Name} declined.`; `No other officer has a Forever Sixty account yet.`; `The offer did not send. Try again.`; added (controller, 2026-10-27 persona round): `Seats, raids and records stay with the guild; {Name} manages seats from then.`, `An offer to {Name} is already waiting. Withdraw it first.`, `A plan change is already under way. Reload to see it.` |
| Receiver | `{Name} offers you the guild plan. You would be billed from {date} at {price}, the founder price. Nothing is charged before then.` (without `, the founder price` at a list price); `Take over the plan` (`Opening checkout…`); `Not now`; added (controller, 2026-10-27 persona round; round 2: the receiver reads it about themselves): `Seats, raids and records stay with the guild; you manage seats from then.` |
| Records | `Records`; `Your guild's raids, attendance and loot are kept for good.`; `Download attendance (CSV)`; `Download loot (CSV)`; `Preparing…`; `The file did not download. Try again.`; file names `{guild-slug}-attendance.csv`, `{guild-slug}-loot.csv` |
| Gate, Guild tier | `Loot council and the readiness board are part of Raid guild. ` + `Move to Raid guild` |
| Read-only officer | `The subscriber or the guild leader manages billing and seats.` (amended (controller, 2026-10-27 persona round); hidden while an offer to take over billing waits for the viewer) |

### 6.3 Account page

| Where | String |
|---|---|
| Offer | `{Guild} offers you a Premium seat.`; `From {Name}, {d MMM}.`; `If you accept, your own Premium stops renewing and ends {date}, and the seat starts then. You are not charged twice. If the seat is ever taken away, we email you a link to come back at {price}, the price you pay now.`; `Accept the seat` (`Accepting…`); `Decline`; `Accepted. Your Premium ends {date}; the seat starts then.`; `Declined. Your Premium carries on as before.`; `That did not save. Try again.` |
| Plan line | `Premium · renews {date} · {price}`; `Premium · ends {date} · then a seat from {Guild}`; `Premium · a seat from {Guild}`; `Given by {Name}. It lasts while the guild's plan does; nothing is billed to you.`; free unchanged |
| Founder | pill `Founder`; title and sr-only `Founder price, kept while your plan stays active.` |
| Seat ended | `Your seat from {Guild} ended {d MMM}. ` + `Come back at {price}` |
| Guild row | `Guild plan`; `Raid guild plan`; ` · you hold a seat` |

### 6.4 Emails (subject, first sentence, link)

| Email | Subject | First sentence | Link |
|---|---|---|---|
| Renewal, 7 days before | `Your {Plan} plan renews on {d MMMM}` | `Your {Plan} plan{ for {Guild}} renews on {d MMMM yyyy} for {price}{, your founder price}; to stop it, cancel before then and nothing more is charged.` | `Manage or cancel` (the site's plan page: `/account#plan`, or the guild's settings `#billing`) |
| Size over 25, 14 days before | `{Guild} has {n} members; Guild covers 25` | `{Guild} has {n} verified members and its Guild plan covers up to 25, so when it renews on {d MMMM yyyy} you may want Raid guild at {price}; nothing switches off either way.` | `Look at Raid guild` (guild settings, the Billing section) |
| Seat offered, to a payer | `{Guild} offers you a Premium seat` | `{Name} offered you a Premium seat in {Guild}; if you accept, your own Premium stops renewing and ends on {d MMMM yyyy}, the seat starts then, and you are not charged twice.` | `Accept or decline` (the account page, Plan panel) |
| Seat given, to a non-payer | `You have Premium from {Guild}` | `{Name} gave you a Premium seat in {Guild}, so you have Premium for as long as the guild's plan lasts, with nothing billed to you.` | `See your plan` (the account page, Plan panel) |
| Seat removed | `Your Premium seat from {Guild} has ended` | `Your Premium seat in {Guild} ended on {d MMMM yyyy}, and because you gave up your own Premium for it, this link brings it back at {price}, the price you had.` | `Come back at {price}` (Checkout at the old price, valid 30 days) |

## 7. Phone and desktop

**Never collapses**: any price, a founder line or its date, the trial line, the free promise, the history promise, every feature line, a seat's lock date, the offer's consequence sentence. **Collapses**: below 1024 the cards go to two columns, and below 768 to one column with the Free summary row first (§3.1.3); below 1024 the interval control takes its own row; below 768 chooser options and seat rows stack (a seat row is who on line 1 with the button at the right, the middle text on line 2, at most 112 px), confirm blocks put their buttons full width with the commit first, and How it works stacks. **Wide**: at 1440 and 2000 the 1344 column is centred, the band is full bleed, and each card keeps about 324 px. Heights: price block 96 desktop, 88 phone; seat row 64 / 112; checkout guild row 56; Free summary row 64 minimum, 80 at 360 where its line wraps to two (amended (controller, 2026-10-27 persona round): 64 px cannot hold the two-line summary at 360).

## 8. Rulings

- **R-P1 A transfer is an offer.** A subscription cannot be put on another person's card without them, so the receiving officer accepts and adds a card in Checkout (setup mode, first charge at the current period end). Until then nothing changes and the old plan keeps running.
- **R-P2 Only payers accept seats.** A member who does not pay has the seat at once and gets the "seat given" email. A payer gets an offer, which never expires on its own and can be withdrawn.
- **R-P3 The 30-day lock starts at acceptance**, or at giving for a non-payer, not when an accepted seat starts later. A withdrawn offer starts no lock. A holder leaving the guild frees the seat at once (product §4).
- **R-P4 The founder window is two instants in one constant** shared by the API and `plans.ts`: opens 2026-11-04T08:00:00Z, closes 2027-01-07T08:00:00Z (midnight US Pacific at each end). The API measures it at Checkout session creation (product §8.2); the page's head script uses the same instants.
- **R-P5 The plans page names only live benefits** (premium spec §5). The pricing design's "2-year retention", "saved sims and builds kept", "compare more than two builds" and "guild-wide retention" do not render: the owner retired them on 2026-10-24, and everything is kept for everyone.
- **R-P6 No founder or seat mark on public surfaces.** The Supporter chip (seat holders resolve as supporters) stays the only public mark.
- **R-P7 Guild buttons on the plans page go to the checkout guild chooser**, because a guild plan needs a guild id and the static page cannot know it.
- **R-P8 Before 4 November the page shows list prices** with the "from 4 November" founder line. Whether purchases made before the window should get founder prices is the owner's call (see the lane note); the page follows whatever Checkout does.

## 9. Performance, contract and acceptance

Limits: the plans page stays static with no `_astro` script. Add `http://localhost/premium.html` to the collect list in `web/lighthouserc.json`; the default assertion row applies (performance and accessibility at least 0.95, LCP at most 2400 ms, TBT at most 100 ms, CLS at most 0.05). The interval switch and the founder state change no box size. `#billing` reserves its ready height (§4), so a warm officer load has CLS at most 0.02 at 390 and 1440. Seat crests use `ClassCrestRing` with its reserved box, so no crest flashes in.

**Contract this spec assumes** (the builder confirms or amends it in the lane report). The guild settings read adds `billing.tier` (`guild` or `raid_guild`), `billing.trial_end`, `billing.trial_used`, `billing.founder`, `billing.interval`, `billing.amount_cents`, `verified_count`, `seats` (each: `slot`, `state` of held, offered, accepted or empty, `holder` with character key, name and class or null, `given_at`, `starts_at`, `movable_at`, `returned` with name and date or null) and `transfer` (`to_name`, `at`, `state` of offered, declined or accepted; or null); `can` gains `seats`. Actions: give, move and withdraw a seat; preview and perform the upgrade; send, withdraw, accept and decline a transfer; the two CSV downloads. `/v1/me` adds `seat` (guild name, given by, since; or null), `seat_offer` (id, guild name, from name, date, `own_ends_at`; or null), `billing.founder`, `billing.interval`, `billing.amount_cents`, and `ended_seat` (guild name, date, `resubscribe_url`; or null).

**E2E** (`premium.spec.ts` extended, new `guild-billing.spec.ts` and `account-seat.spec.ts`; API stubbed; clock faked):
1. The plans page on 27 Oct 2026: the founder row reads the before line; prices are list prices with no strike.
2. On 8 Dec 2026: the founder row, struck list prices, founder prices; `Yearly` is checked; the Premium button reads `Subscribe · $40 a year` and links to checkout with plan premium and interval yearly.
3. Switching to `Monthly` changes every price and button; every card's height is unchanged; the arrow keys switch too.
4. A page built on 6 Jan 2027 and read on 7 Jan 2027: no founder row, no strike, list prices.
5. No `_astro` script and no third-party script loads; `/pricing` answers 301 to `/premium`.
6. At 360: the Free summary row precedes Premium; Premium's button bottom is above y 600; `#free` holds the full list and `This list never shrinks.`
7. No text on the plans page contains `retention`, `compare more than`, or an exclamation mark.
8. A guild card button opens the chooser; a guild with 28 verified members reads `Raid guild fits 28 members`; an account with no billable guild sees the none line.
9. `#billing`, not on a plan: the option that fits the count is preselected; `Start the 7-day trial` carries plan, interval and guild id; a used trial shows `Subscribe · …` and no trial line.
10. Trial: the plan line and its cancel-before date.
11. Guild tier with 28 members: the size line and exactly one `Move to Raid guild`; the confirm shows the preview amount, focus is on `Keep Guild`, Escape closes; the commit updates the plan line.
12. Seats: give to a non-payer (Held, its lock date); give to a payer (the picker's payer note, then Waiting); withdraw; a locked `Move seat` is aria-disabled and its date is visible text; 409 `seat_locked` shows its line; a returned seat shows its line.
13. Transfer: send, the waiting state, withdraw; the receiver sees the offer and `Take over the plan` opens checkout; a decline shows to the sender.
14. Lapsed: the read-only line, `Subscribe again`, and both CSV buttons download with the right file name and the CSV content type.
15. A read-only officer sees the plan line, `Billed by`, seats with no buttons, and Records.
16. On a Guild-tier guild the Loot and Readiness gate sentence links to `#billing`.
17. Account: a payer sees the offer with their own end date and price; accept gives the status and `Premium · ends … · then a seat from …`; decline gives its status; a seat holder sees `a seat from` and no `Manage billing`; a founder sees `Founder`; the ended-seat line links to its resubscribe URL.
18. Axe is clean in every state above; every control is at least 44 px tall at 360.

Email cases (unit, the billing mailer's tests): each subject and first sentence equals §6.4 for a fixture; the size email does not send when the count is 25 or fewer on its send day; the renewal email does not send for a monthly plan.

**Screenshots** (fixture build; x = required):

| # | State | 360 | 390 | 1440 | 2000 |
|---|---|---|---|---|---|
| 1 | Plans page, founder open, Yearly | x | x | x | x |
| 2 | Plans page, Monthly | | x | x | |
| 3 | Plans page, before the window and after it | | x | x | |
| 4 | Plans page, focus-visible on the interval control and on a card button; hover on a card button | | | x | |
| 5 | Checkout guild chooser: two guilds; no billable guild | | x | x | |
| 6 | `#billing` not on a plan: trial available, trial used | x | x | x | |
| 7 | `#billing` trial | | x | x | |
| 8 | `#billing` Guild with 28 members: size line, the confirm loading and ready | x | x | x | |
| 9 | `#billing` Raid guild seats: held and locked, waiting, accepted, empty, returned; the picker open with a payer chosen | x | x | x | x |
| 10 | Who pays: send confirm, waiting, the receiver's offer | | x | x | |
| 11 | Lapsed with Records; the read-only officer view | | x | x | |
| 12 | Account: offer, accepted, seat holder, founder, seat ended | x | x | x | |
| 13 | The five emails in Gmail web and iOS Mail | | | | |
| 14 | Same-scale side-by-side with the mock boards at 1440, plus 1024 and 2000 | | | x | x |

**Controller ruling on R-P8 (2026-10-27):** founder prices apply from the day the launch pricing ships until 6 January 2027, 23:59:59 America/Los_Angeles. There is no "from 4 November" line; anyone who buys before launch gets the founder price.

**Amendment (controller, 2026-10-27, from the Task 16 review):** below 768 px the plans page hides the `PLANS` eyebrow, uses band padding 12/12 and a 14 px lede, so Premium's subscribe button sits above y 600 at 360 px. The standalone no-ads line (§3.1.6) is covered by the lede and the "Are there ads?" FAQ.

**Amendment (controller, 2026-10-27, persona round on the launch pricing captures):** the copy changes above are recorded in their §6 rows. Also: the account Premium rows read `In your browser` (Server sims, off) and `Premium only` (Discord notifications, free) in place of `Browser lane only` and `Off · Premium`; class names in the seat picker are capitalised (`Mirelle · Priest · pays for Premium`); a member with an open seat offer is left out of the picker; the seat-ended sentence replaces the plan line with the gold inline link; the answered seat offer is a plain status line with no rule; the founder Guild upgrade lands on the founder Raid guild price and the confirm quotes it (verified in `api/internal/billing/upgrade.go` and `catalog.UpgradeLookupKey`); `Billed by` names the payer's character in the guild, the public name only without one.

**Amendment (controller, 2026-10-27, persona round 2):** the guild plan line's `Founder` is the gold `pill-site` 8 px after the price, as on the account panel; the chooser's struck list figure sits 8 px after the price; the receiver reads `Seats, raids and records stay with the guild; you manage seats from then.`; from 1024 to 1279 the plan cards take 16 px padding and the card buttons no extra tracking, so the guild labels hold one line; the checkout chooser's right-hand label shows only when the plan picked on /premium does not fit the guild (`Raid guild fits {n} members` on a Guild purchase above 25), and every row's label shows only when no plan was picked; while a seat offer stands, the account Discord panel's `Save webhook` is a gold outline (§5: the offer is the page's only fill); How it works terms stay 12 px `--strong` 600 per the mock.
