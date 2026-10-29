# Subscribe and join: the checkout guild step, claim and verify, and "you already have it" (experience spec)

Author: ux-designer, 2026-10-27. Supersedes launch pricing §3.1.7 (the checkout guild chooser) and amends launch pricing §3.3's plan line and §6.3. Where Premium's value and the guild leader's view meet this spec, the amendment "2026-10-27b: Premium value and guild leaders" in `design/specs/2026-10-27-launch-pricing.md` owns the plans page cards, `#billing` and the `Verified members` rule; this spec owns the checkout page, the claim and verify path, the "you already have it" states and the account Plan panel's headline. Read with `design/specs/2026-10-25-officer-tools-gate.md` (gate grammar), `design/specs/2026-10-26-guild-raids.md` (verified-only surfaces), `design/specs/2026-10-27-discord-bot.md` (bot links) and `docs/tenets.md`.

Code read for this spec: `web/src/components/CheckoutLauncher.svelte`, `web/src/components/checkout/CheckoutGuildChooser.svelte`, `web/src/lib/billing/{checkout-params,copy,plans,contract,normalize,account-plan-copy}.ts`, `web/src/lib/account/{api,plan-line}.ts`, `web/src/components/account/AccountPremiumBlock.svelte`, `web/src/lib/guild/copy.ts`, `web/src/components/GuildClaim.svelte`, `web/src/components/GuildLanding.svelte`, `web/src/components/guild/GuildInviteLink.svelte`, `api/internal/guilds/{claim,capabilities,invite,membership}.go`, `api/internal/billing/{handler,checkout}.go`, `api/internal/auth/{store,me_billing}.go`, `api/internal/entitlements/store.go`.

**Verified member** means what the code and the amendment's `Verified members` rule say, in the amendment's words: a verified member has a Forever Sixty account and a character the guild has confirmed: an officer approved it, they joined by the guild's invite link, or they turned up in the guild's logs. Each player counts once, however many alts they have. Discord members without an account do not count. An email-link account counts the same as any other. No copy in this spec ties verification to Battle.net.

**Raid size.** Forever runs 10-, 20- and 40-person raids. No copy here assumes 40. A 10-person raid team, and a 20-person one with a small bench, fits under Guild's 25-member cap; the checkout says so where the plan is shown (§5.1, §9.1).

## 0. What we got wrong before

1. Launch pricing §3.1.7 listed only guilds the viewer could already bill (`can` holds `billing`, `plan` null). Billing needs a claimed guild and a verified claim holder. On launch day almost no guild is claimed. So the list was empty for nearly everyone, and the "edge case" line was the page.
2. The empty line told the player a rule ("a guild you lead or have claimed") and sent them to `/guild`. `/guild` forwards only verified members. The owner, unverified in both guilds, landed on a name lookup. It was a dead end.
3. We already knew both guilds, the owner's rank in each, that nobody in either was verified, and that neither was claimed. The page showed none of it.
4. The page dropped the full-bleed band, the crests, the character and the plan summary. It was a narrow form column with three lines of text. That fails tenets 1, 11 and 14.
5. `plan=premium` for an account that has Premium printed the API's raw lowercase refusal, `this account already has Premium`. We never designed "you already have it".
6. The account Plan panel built its headline from `current_period_end`. An operator grant has none, so a granted account read `Not on Premium.` above four rows that said On. `Manage billing` showed for any billing row, including a grant with no Stripe customer, so the button could only fail.
7. Every member-side sentence assumed a verified officer exists to approve them (`An officer can approve you from the roster.`). With zero verified members that sentence is false.

## 1. The one question, and the first screen

| Surface, viewer | Question | Answer visible with no scroll at 360 x 640 and 1280 x 800 |
|---|---|---|
| `/premium/checkout?plan=guild…`, any signed-in player | "Can I subscribe my guild, and if not, what is the one thing that unblocks it?" | The band with the plan and price, the lead sentence, and the first guild card: crest, name, the viewer's character and rank, verified state, and the card's one next step (a button, or a sentence plus `Copy link`). 360: the first card's action ends above y 640. 1280: the first two cards end above y 800. |
| Same, the owner's account (two guilds, member, unverified, unclaimed) | "Why can't I pay, and what do I do?" | The lead `Only a guild's guild master can claim it here and subscribe it.`, then OLYMPUS XXVII's card with `Send them this link.` and its `Copy link` button. |
| `/premium/checkout?plan=premium`, an account that has Premium | "Do I need to buy this?" | `You already have Premium.` and the source line (`Premium · granted, no end date`), with its one action. |
| `/account#plan`, any account with Premium | "Am I on Premium, from where, and until when?" | The plan line naming the source and the end (or `no end date`), and line 2 saying who pays. |
| `/premium`, an account that has Premium | "Is this for me?" | The Premium card says `You have Premium through …` and its button reads `See your plan`. |

## 2. References, and what we borrow

- **The in-game Guild Control panel and the roster's rank column.** Players already read "rank, then what that rank may do". Borrowed: the rank word beside the character (`Guild master`, `Officer`, `Member`) and one capability per rank. Not borrowed: the permission checkbox grid.
- **The Battle.net shop's "Owned" state.** A product the account owns never shows Buy. It shows that it is owned, the source, and a link to where it is managed. Borrowed for every "you already have it" state (§4).
- **Discord's server invite screen.** One server card: icon, name, member count, then one primary action. Borrowed: the guild card's crest, name, counts, then one action.
- **Raid-Helper's "ask an officer" reply** (a member who cannot post a raid is told who can). Borrowed: name who can act, and hand the member something to send them.
- **Our own surfaces:** the `/premium` band (launch pricing §3.1.1), the `card` recipe, `GuildInviteLink`'s copy-and-share block, `ClassCrestRing`, `FactionMark`, the claim confirm block in `GuildClaim.svelte`, `StatePanel` and the key/value rows of `AccountPremiumBlock`, `PRIMARY_BUTTON_FIXED`, `SECONDARY_BUTTON_FIXED`, `INLINE_LINK`, `pill-site`, `Skeleton`, `LoadError`, `EmptyState`.

## 3. Journey map

### 3.1 Entries into "pay" and "join, verify or claim"

| # | Entry | Where it lands after this spec |
|---|---|---|
| E1 | `/premium` Guild or Raid guild card button | `/premium/checkout?plan={guild or raid_guild}&interval={…}`, the guild step (§5) |
| E2 | `/premium` Premium card button | `/premium/checkout?plan=premium&interval={…}`: Stripe, or "you already have it" (§4.2) |
| E3 | Guild page gate sentences (`Subscribe the guild`, `Move to Raid guild`) | Officers: settings `#billing` as today. Members and unverified viewers: the checkout guild step with `guild_id`, which shows their step (§5.6) |
| E4 | Guild page claim slot and the unverified note | `/guild/{r}/{rs}/{slug}/claim` (§6.2) |
| E5 | Guild settings `#billing`, an officer who does not hold the claim | The hand-off line with a copy link for the claim holder (§6.4) |
| E6 | Account page, guild row with no plan | `See the next step`, to the checkout guild step scrolled to that guild's card (§6.5) |
| E7 | Account page Plan panel | `See plans`, `Manage billing`, `Billing history` or nothing, by source (§4.3) |
| E8 | Discord bot replies with `Subscribe the guild` or `Move to Raid guild` | The checkout guild step with `plan=raid_guild&interval=yearly&guild_id={id}` (§6.6) |
| E9 | An officer's invite link `/guild/invite/{token}` | Join, verified at once (today's flow); the joined state gains the guild's plan line (§6.3) |
| E10 | The hand-off link a member copies on the checkout step | The guild's claim page, or the claim holder's settings `#billing`, with `?plan&interval`; it ends in Checkout for the person who can pay (§6.1, §6.2) |

### 3.2 Every state, and what each surface shows

`step` is the new API field `me.guilds[].subscribe.step` (§8). The web never derives it from a rank.

| State | `step` | Checkout guild step (§5) | Guild page / claim page | Account page |
|---|---|---|---|---|
| S0 Signed out | n/a | Redirect to `/login?next=` (today) | Claim page: `Sign in to claim this guild.` | n/a |
| S1 Signed in, no characters yet | n/a | Empty state A: `Get set up` | Claim page: `notAMember` with `Get set up` | No guild rows |
| S2 Characters, none in a guild | n/a | Empty state B | n/a | No guild rows |
| S3 Member, not verified, guild unclaimed, nobody verified (**the owner, in both guilds**) | `ask_leader` | Card: `Only {Guild}'s guild master can claim it and subscribe it. Send them this link.` + copy link + `Lead {Guild} on another character?` | Unverified note, no-approver form, with `Send your guild master a link` | Row: `Not verified` + `See the next step` |
| S4 Member, not verified, a verified officer exists | follows the claim (S10, S11, S12) | That step + the verify sub-line (officer approval, invite link, logs) | Today's note `An officer can approve you from the roster.` | Row: `Not verified` |
| S5 Verified member, guild not on a plan | `claimed_ask` or `ask_leader` | Card: hand-off to the claim holder or guild master, copy link | Gate sentences as today | Row: `See the next step` |
| S6 Officer rank, guild unclaimed | `start_claim` | Card: `Start a claim` (in-place confirm) | Claim page as today | Row: `See the next step` |
| S7a Officer, own claim pending | `claim_pending_yours` | Card: waiting sentence + copy link for a second officer or the guild master | Claim page pending line | Same |
| S7b Officer or guild master, another's claim pending | `confirm_claim` | Card: `Confirm the claim` (in-place) | Claim page `Confirm this claim` | Same |
| S7c Member, a claim pending | `claim_pending_ask` | Card: pending sentence + copy link to the claim page | Pending line | Same |
| S8a Guild master rank, unclaimed, claim allowed | `claim` | Card: `Claim {Guild}` (in-place); the card then becomes S11 | Claim page as today, plus the plan carried by `?plan` | Same |
| S8b Guild master rank, no Battle.net link (the claim rule in `claim.go`, `ErrNoBattleNetIdentity`; not a verification rule) | `claim_link_bnet` | Card: `Link Battle.net` | Claim page rule line | Same |
| S8c Guild master rank, already holds another guild's claim | `claim_blocked` (`holds_other`) | Card: one-claim sentence, no action | Claim page refusal | Same |
| S8d Guild master rank, made a claim in the last 30 days | `claim_blocked` (`rate_limited`) | Card: `You can claim {Guild} from {date}.` | Same | Same |
| S9 Claim contested | `frozen` | Card: contested sentence, `Open the guild` | `frozenNotice` | Same |
| S10 Claimed by someone else (viewer is an officer or member) | `claimed_ask` | Card: names the holder to officers, "the guild master who holds the claim" to members; copy link to settings `#billing` with the plan | Today | Same |
| S11a Claim holder, never subscribed, trial unused | `ready` | Card: `Start the 7-day trial` (the page's one fill) | `#billing` chooser | Row: `See the next step` |
| S11b Claim holder, trial used | `ready` | Card: `Subscribe · $120 a year` | `#billing` chooser | Same |
| S11c Claim holder, plan lapsed | `ready` | Card: `Subscribe again · $120 a year` | `#billing` lapsed | Row: today's lapsed line |
| S12a On a plan, viewer pays, chosen tier is current or lower | `on_plan` | Card: `You pay for it.` + `Manage billing` | `#billing` | Row: today's plan line |
| S12b On Guild, viewer pays, chose Raid guild | `on_plan` | Card: `Move to Raid guild`, to the `#billing` confirm | `#billing` | Same |
| S12c On a plan, viewer does not pay, holds a seat | `on_plan` | Card: `You hold one of its Premium seats.` | Today | Row: ` · you hold a seat` |
| S12d On a plan, viewer does not pay, seats free | `on_plan` | Card: `{given} of {total} Premium seats are given. Ask {Name} for one.` | Today | Row: plan line |
| S12e On a plan, all seats given | `on_plan` | Card: `All {total} of its Premium seats are given.` | Today | Same |
| S13 A checkout for the guild is already open | `checkout_open` | Card: open-checkout sentence + `Check again` | n/a | n/a |
| S14 Lapsed, viewer is not the claim holder | `claimed_ask` | As S10, plan word `Plan ended {date}` | Read-only gate | Same |
| S15a Own Premium subscription (active, trial, founder) | n/a | `plan=premium`: §4.2 panel, `Manage billing` | n/a | `Premium · renews …` + `Manage billing` |
| S15b Own Premium ending or past due | n/a | Same, `Manage billing` | n/a | `Premium · ends …` or the past-due alert, `Manage billing` |
| S15c Premium from a guild seat | n/a | §4.2 panel, `See your plan` | n/a | `Premium · a seat from {Guild}` + `Given by {Name}…` |
| S15d Operator grant, open-ended (**the owner**) | n/a | §4.2 panel, `See your plan` | n/a | `Premium · granted, no end date` + `Given by Forever Sixty…` |
| S15e Operator grant until a date | n/a | §4.2 panel | n/a | `Premium · granted until {date}` |
| S15f Free | n/a | Stripe as today | n/a | `Not on Premium.` + `See plans` |
| S16 Own Premium and offered a seat | n/a | §4.2 panel | n/a | Launch pricing §3.3 offer block, unchanged |

## 4. "You already have it"

### 4.1 One source rule

`premiumSource(me)` in `web/src/lib/account/plan-line.ts` returns exactly one of `own`, `own_then_seat`, `grant`, `seat`, `fallback`, `none`, in this order:
1. `billing.source === 'stripe'` and the status is live (`active`, `trialing`, `past_due`): `own`, or `own_then_seat` when `me.seat.starts_at` is in the future.
2. `billing.source === 'grant'` and the row is live: `grant`.
3. `me.seat` present: `seat`.
4. `entitlements.server_sims === true` with none of the above: `fallback`. The line reads `Premium` with no button. "Not on Premium" above On rows becomes impossible.
5. Otherwise `none`.

`Manage billing` appears only for `own` and `own_then_seat`. `Billing history` (the same `SECONDARY_BUTTON_FIXED`, opening the Stripe portal) appears for every other source when `me.entitlements.personal_portal === true` (the account has had its own Stripe subscription). Otherwise no billing button renders. A grant with no Stripe customer shows no button.

### 4.2 `/premium/checkout?plan=premium`

The launcher reads `/v1/me` before it calls `POST /v1/billing/checkout`. When `premiumSource` is not `none`, it never calls Stripe and renders the panel below. If the API refuses anyway (a race), the 409 carries `fields.source` (§8); the launcher re-reads `/v1/me` with `forgetSession()` and renders the same panel. No API sentence is ever shown on this page.

The band carries the h1 `Premium` alone and hugs it: padding 28/24 desktop and 12/12 phone, no plan-summary reserve (51 px phone, 103 px from 768). The head script marks a `plan=premium` URL before first paint (`html[data-checkout-band='premium']`), so the static band and the hydrated one are the same box (gate fix, row 2, 2026-10-28).

```
1280 ──────────────────────────────────────────────────────────────── full-bleed band
 CHECKOUT                                                   (label, --gold)
 Premium                                                    (h1 Cinzel 700 22)
──────────────────────────────────────────────────────────────────────────────
            ┌──────────────────────────────── 640 ─┐
            │ You already have Premium.            │  15 --strong
            │ Premium · granted, no end date       │  14 --text
            │ Given by Forever Sixty. Nothing is   │  13 --muted
            │ billed to you.                       │
            │ [ SEE YOUR PLAN ]                    │  SECONDARY_BUTTON_FIXED px-4
            └──────────────────────────────────────┘
```

Panel: `card` recipe, max 640 px, centred, padding 20, flex column gap 8, the button 12 px below line 2. The plan line and line 2 come from `accountPlanLine(me, now)` (§4.3), so the two surfaces cannot disagree. Button by source: `own` and `own_then_seat`, `Manage billing` (the portal); `seat`, `grant` and `fallback`, `See your plan` (link to `/account#plan`). Founder: `pill-site` `Founder` 8 px after the plan line, as on the account panel. Past due: the account panel's `pastDueBanner` as `role="alert"` under line 2. Focus: the title `p` takes focus (tabindex −1) when the panel replaces the redirecting message. Phone: the panel is full width inside 18 px gutters, and the button is full width.

### 4.3 Account `Plan` panel headline (amends launch pricing §3.3)

| Source | Plan line (14 px) | Line 2 (13 px muted) | Button |
|---|---|---|---|
| `own`, renewing | `Premium · renews 22 Dec 2026 · $5 a month` + `Founder` | none | `Manage billing` |
| `own`, ending | `Premium · ends 22 Dec 2026` | none | `Manage billing` |
| `own`, past due | the renewing line | `pastDueBanner` as an alert | `Manage billing` |
| `own_then_seat` | `Premium · ends 22 Dec 2026 · then a seat from Olympus XXVII` | none | `Manage billing` |
| `seat` | `Premium · a seat from Olympus XXVII` | `Given by Kraggor. It lasts while the guild's plan does; nothing is billed to you.` | `Billing history` if `personal_portal`, else none |
| `grant`, open-ended | `Premium · granted, no end date` | `Given by Forever Sixty. Nothing is billed to you.` | `Billing history` if `personal_portal`, else none |
| `grant`, until a date | `Premium · granted until 31 Mar 2027` | `Given by Forever Sixty. It ends on its own; nothing is billed to you.` | same |
| `fallback` | `Premium` | none | none |
| `none` | `Not on Premium.` + `See plans` | none | `Billing history` if `personal_portal`, else none |
| Seat ended | launch pricing §3.3's ended sentence, unchanged | | |

The rows under the line (`AccountPremiumBlock`) are unchanged. A guild-plan member's Supporter row reading On does not make the plan line say Premium. It is a guild benefit, so line 1 stays `Not on Premium.` for them.

### 4.4 `/premium` Premium card for an account that has Premium

The static page cannot read the session at build time. It reuses the founder pattern (launch pricing §3.1, `data-founder`):
- Storage follows the account page's `fs.accountReserve` hint in `web/src/lib/account/layout.ts` (Decision 10): one key, read and written through the same guarded `storage()` helper, removed on sign-out and whenever the stored value is not one of the three words. A stale or malformed hint only ever shows the signed-out default.
- The session module (`web/src/lib/account/api.ts`, where `fetchMeOnce` resolves) writes `localStorage['fs.premiumSource']` as `own`, `seat` or `grant` (`own_then_seat` stores `own`; `fallback` and `none` remove the key). It also sets `document.documentElement.dataset.premium` to the same value. The header's `CharacterSelector` already calls it on every page, `/premium` included, so the attribute reconciles within one `/v1/me` round-trip.
- A third head `is:inline` script on `/premium`, at most six lines, copies the stored value onto `data-premium` before first paint.
- The Premium card holds both versions of two slots in the DOM at identical heights. CSS on `html[data-premium]` picks one:
  - The 13 px line under the card name: the default "who it is for" line, or `You have Premium through your subscription.` / `You have Premium through a seat from your guild.` / `You have Premium through a grant from Forever Sixty.`
  - The 44 px button slot: the default `Subscribe · $40 a year`, or a neutral `SECONDARY_BUTTON_FIXED` link (`border-line-warm-strong text-strong`, never gold) `See your plan` to `/account#plan`.
- The price block stays. The guild cards do not change (R-P7: a player on Premium may still subscribe a guild).
- No layout shift: both versions sit in the same fixed-height cells. Lighthouse runs signed out, so the default renders.

## 5. The checkout guild step (`plan=guild` or `plan=raid_guild`)

### 5.1 Page frame

`web/src/pages/premium/checkout.astro` drops `measure-form` for this step. Order:
1. **Band**, full bleed, the `/premium` band markup (launch pricing §3.1.1), padding 28/24 desktop and 12/12 phone, its content in the 1344 px column centred at every width up to 2000 and beyond.
   - `label` `CHECKOUT` in `--gold` (hidden below 768, as on `/premium`). From 1024 the band row is top-aligned (`align-items: flex-start`), so the label's top lines up with the plan summary's first line, as `/premium` top-aligns its interval control with `PLANS` (gate fix, row 1).
   - h1, Cinzel 700 22 px `--strong`: `Subscribe a guild`.
   - **Plan summary.** From 1024 it is right-aligned on the h1's row; below 1024 it is its own row under the h1. Line 1: the plan name, Cinzel 700 15 px `--strong` (`Guild`); 8 px; the price, JetBrains Mono 500 20 px `--strong`, tabular (`$120`), then ` a year` 14 px muted; then `pill-site` `Founder` 8 px after it while the founder window is open. Line 2, 13 px `--text`: the size line (§9.1). Guild: `Up to 25 verified members. A 10-person raid team fits, and a 20-person one with a small bench.` Raid guild: `Any size. For guilds over 25 verified members.` Line 3, 13 px `--text`: `First 7 days free` (`--strong` 600) when the plan carries a trial, then ` · ` and the underlined `--text` link `Change plan` to `/premium`.
2. **Body**, a centred column, max 880 px, padding 32 px top desktop and 20 px phone, side gutters 48 / 18.
   - **Lead sentence**, 15 px `--text` (14 px below 768), max 720 px, 16 px below it. Chosen from the set of steps (§9.1), in this order (gate fix G6, 2026-10-28): any card `ready` gives the `ready` lead; else any card `claim`, `start_claim`, `confirm_claim`, `claim_link_bnet` or `claim_pending_yours` (the viewer claims, starts, confirms or holds the pending claim) gives the can-claim lead; else a `claim_blocked` card gives the some-claimed form; else the §9.1 rows as written (gate round 2, 2026-10-28). A lead never tells a player who can act that the next step is someone else's.
   - **Guild cards**, a `ul` with gap 12. Order: `ready` first, then `claim`, `confirm_claim`, `start_claim`, `claim_link_bnet`, then the hand-off steps (`claimed_ask`, `ask_leader`, `claim_pending_yours`, `claim_pending_ask`), then `checkout_open`, `frozen`, `claim_blocked`, then `on_plan`. Ties keep the order of `me.guilds`.
   - **Foot**, 24 px below the last card, 13 px muted, two lines at most: the missing-guild line with `Get set up`; then, only when `premiumSource(me)` is `none`, the personal Premium line with `Get Premium`.

### 5.2 The guild card

`li` with the `card` recipe (`--raised`, 1 px `--line`, 6 px radius, 1 px top highlight `rgba(229,185,85,.35)`), padding 20 desktop and 16 phone, `id="guild-{id}"`, `scroll-margin-top: 96px`, `aria-labelledby` its guild name.

Desktop (from 768): grid `minmax(0,1fr) 320px`, gap 24, `items-start`.

Left column, three lines:
1. **Identity row**, flex, gap 12, items centred. The guild crest is a 40 px circle with a 2 px `--gold` ring holding the `guildMarkSrc` image at 40 x 40 `object-cover` (the render `GuildCrestSettings` uses). With no crest, `FactionMark` at 24 px sits centred in the ring. With no faction either, the ring is empty. Then the name, Cinzel 700 16 px `--strong`, truncating; under it the ruleset and region, 13 px muted (`Forever · US`).
2. **You row**, 8 px below, flex wrap, gap 8, 14 px. One entry per character of the viewer in this guild (`me.characters` whose `guild.id` matches), main character first, at most two, then `+{n} more` muted. Each entry: `ClassCrestRing` at 20 px; the character name in its class colour (Barlow 600); the rank word (`Guild master`, `Officer`, `Member`) in `--text`; the verify mark. Verified: a 12 px `--kill` tick (the plans page tick glyph, `aria-hidden`) and `Verified` in `--kill`. Not verified: `Not verified` in `--muted`. Separator ` · ` muted. The name and the rank word never break inside (`white-space: nowrap`); on a narrow phone the verify mark, with its separator, drops to its own line (gate round 2, row 1b at 360).
3. **Guild facts**, 6 px below, 13 px muted, numbers in JetBrains Mono: `{n} verified` (or `No one verified yet` at 0), then `{m} on the roster` (left out when the guild has no addon snapshot), then the plan word (§9.1), joined by ` · `.

Right column, the **step block**, flex column, gap 10:
- The step sentence, 14 px `--text`, at most four lines at 320 px.
- The one action: a button, the copy-link field, a text link, or nothing (§5.3).
- The verify sub-line and the other-character line, 13 px muted, when §5.4 says so.

Phone (below 768): one column. The step block follows the facts line, after a 1 px `--line-soft` rule with 14 px above and below. Buttons and the copy field go full width.

The card has no hover or focus state; it is not a control. Only the controls inside it have states.

### 5.3 Steps: action and behaviour

Copy is verbatim in §9.1.

| `step` | Action control | Behaviour |
|---|---|---|
| `ready` | `PRIMARY_BUTTON_FIXED w-full justify-center` | Starts Checkout with `plan`, `interval` and `guild_id` (today's `startCheckout`). The label states trial, price, or again. One fill per page: when two or more cards are `ready`, every `ready` button is `SECONDARY_BUTTON_FIXED border-gold text-gold`. |
| `ready`, plan `guild`, `verified_count` above 25 | The same button, with a 13 px line above it ending in the inline link `Switch to Raid guild` | The link rewrites the URL's `plan` in place (replaceState) and updates the band and the button; no reload. |
| `ready`, plan `raid_guild`, `verified_count` 25 or fewer | The same button, with a 13 px line above it ending in the inline link `Switch to Guild` | Same mechanism. It is the honest note for a 10- or 20-person raid team that picked the bigger plan. |
| `claim`, `start_claim`, `confirm_claim` | `SECONDARY_BUTTON_FIXED border-gold text-gold w-full justify-center` | Opens the claim confirm block in place under the button, the `GuildClaim.svelte` recipe (`border-line-soft border p-4`, gap 12): the confirm sentence, the commit (`PRIMARY_BUTTON_FIXED`) and `Cancel` (44 px underlined text button). Focus moves to `Cancel`. Escape closes the block and returns focus to the opener. One confirm is open at a time; the other claim buttons are disabled meanwhile. On success the launcher re-reads `/v1/me` with `forgetSession`, the card re-renders in its new step, and a 13 px muted status line (`role="status"`) states the outcome. Focus moves to the card's new action, or to the status line when there is none. |
| `claim_link_bnet` | `SECONDARY_BUTTON_FIXED w-full justify-center` link | To the Battle.net link start, with `next` set to this checkout URL plus `#guild-{id}`. |
| `ask_leader`, `claimed_ask` (except `claim.yours` — the holder's own card shows no copy field and no moderator fallback; controller ruling 2026-10-27), `claim_pending_yours`, `claim_pending_ask` | `CopyLinkField` (§5.5) | Holds the hand-off URL (§6.1). |
| `on_plan`, viewer pays, chosen tier is current or lower | `SECONDARY_BUTTON_FIXED w-full justify-center` `Manage billing` | `POST /v1/billing/portal` with `guild_id`, as `#billing` does. |
| `on_plan`, viewer pays, on Guild, chose Raid guild | `SECONDARY_BUTTON_FIXED border-gold text-gold w-full justify-center` `Move to Raid guild` | Link to the guild's settings with `upgrade=1` and `#billing`, which opens launch pricing §3.2.4's confirm. |
| `on_plan`, viewer does not pay | 44 px underlined text link `Open the guild` | To the guild page. |
| `checkout_open` | 44 px underlined text button `Check again` | Re-reads `/v1/me`. |
| `frozen`, `claim_blocked` | 44 px underlined text link `Open the guild` | |

### 5.4 The verify sub-line, and the other-character line

The verify sub-line shows under the step when the viewer has no verified character in the guild. It names the three routes of the `Verified members` rule:
- The guild has a verified officer (`verify.has_approver` true): `Not verified yet. An officer approves you from the roster, the guild's invite link verifies you at once, or you are verified when you turn up in the guild's logs.`
- No verified officer, the guild claimed: `Not verified yet. You are verified when you turn up in the guild's logs, or once the claim holder approves you.` (gate fix, 2026-10-28: a claimed guild is never waiting on its guild master to claim it).
- No verified officer, the guild not claimed, and the step is not `ask_leader` or `claim_pending_ask`: `Not verified yet. You are verified when you turn up in the guild's logs, or once the guild master claims {Guild} and approves you.`
- On `ask_leader` and `claim_pending_ask` with no verified officer, the step sentence already says the guild master acts first, so the short form shows: `Not verified yet. Turning up in the guild's logs also verifies you.`
- Never on `claim_blocked`: its viewer is the guild master, who is never told to wait for the guild master (gate fix, 2026-10-28).
- Never on `claim`, `start_claim` or `confirm_claim`: claiming verifies the claimant (`setVerifiedForAccount` in `claim.go`; the API lane confirms an officer's claim does the same once confirmed).

Under the `ask_leader` step only, a second 13 px muted line: `Lead {Guild} on another character? Save that character with the addon, then come back.` with the inline gold link `Get set up` (to `/setup` with `next` set to this URL plus `#guild-{id}`).

On `ask_leader` and `claimed_ask` cards, last in the step block, 13 px muted: `Guild master gone? Ask us on Discord and we will sort the claim by hand.`, with `Ask us on Discord` as an `INLINE_LINK` to `communityLinks.discordInvite` (`web/src/data/links.json`), new tab, `rel="noopener"` (Decision 8).

### 5.5 `CopyLinkField` (an extraction, not a new control)

`GuildInviteLink.svelte` already draws a read-only field holding a URL, a `Copy link` button that reads `Copied` for 2 s, a `Share` button where the browser offers sharing, and a failure line. Extract that half into `web/src/components/ui/CopyLinkField.svelte` (props: `url`, `label` for the sr-only field label, `shareTitle`, `shareText`, `testid`, and an optional `note` snippet rendered at the end of the button row, which `GuildInviteLink` uses for its "Shown once" line), and make `GuildInviteLink` use it. The visuals stay: field 44 px, `border-line-warm bg-raised rounded-control` (today's border, so the captures stay pixel-identical), 13 px mono, truncating, `readonly`, select-all on focus; buttons `SECONDARY_BUTTON_FIXED px-4`. In the 320 px step column the field takes its own row and the buttons sit under it, side by side. Gate fixes G3 to G5 (2026-10-28) add three props whose defaults keep `GuildInviteLink`'s render: `labelVisible` (false: the label is `sr-only`; set on the checkout card and the claim page member hand-off), `onCard` (the field's ground is `--bg`, so it reads as a field against the card's `--raised`; set on the checkout card) and `phoneFullWidth` (below 768 the button row is a one-column grid, two columns with `Share`, and each button is `w-full justify-center`; from 768 `w-fit`; set on the checkout card and the claim page hand-off, whose field also stops at 640 px). It is an extraction of a reviewed control, so it needs no separate design review; the reviewer checks that the `GuildInviteLink` captures stay pixel-identical (capture 17).

### 5.6 Arriving with a `guild_id`

`guild_id` is allowed on the guild step from now on (E3, E8). When that guild is `ready`, today's direct launch to Stripe stays. Otherwise the step renders with that guild's card first, scrolled into view, and the other cards below. An id that is not in `me.guilds` renders the full list with no error.

### 5.7 Empty, loading and error

- **Loading.** The band renders from the first paint, since it needs only the URL. The lead slot reserves one line (21 px; 42 px below 768). The body holds `Skeleton` cards at the ready height, 156 px desktop and 260 px phone, their count read from `localStorage['fs.meGuildCount']` (written by the session module) and clamped to 1 to 4, else 2. sr-only `Loading your guilds`; the `ul` carries `aria-busy="true"`. Rows replace the skeleton once `/v1/me` has answered and every crest image has decoded or failed, capped at 1000 ms after the answer, so no crest flashes in. That is today's rule, now fed by `me.guilds[].crest_url`, with no `/home` reads.
- **Error.** `LoadError` with `Your guilds did not load.` and `Try again`, in the body's place. After a successful retry, focus moves to the lead sentence.
- **Empty A** (no characters; S1). `EmptyState` in a `card`, max 640, left in the 880 column where the foot sits too (gate fix, row 11): title 15 px `--strong`, body 14 px `--text`, then `Get set up` (`PRIMARY_BUTTON_FIXED px-4`, the page's one fill, to `/setup` with `next` set to this URL; full width below 768, `w-fit` from 768). The foot's missing-guild line is not drawn, since it repeats the action.
- **Empty B** (characters, none in a guild; S2). The same panel with its own copy and `Get set up`. Here too the foot's missing-guild line is not drawn.
- **Signed out.** Today's redirect to `/login` with `next` stays.

### 5.8 Wireframes

Desktop 1440, the owner's account:

```
═══════════════════════════════════════════════════════════════════════════════════════════ full bleed
   CHECKOUT
   Subscribe a guild                                    Guild  $120 a year  [FOUNDER]
                                                        Up to 25 verified members. A 10-person
                                                        raid team fits, and a 20-person one
                                                        with a small bench.
                                                        First 7 days free · Change plan
═══════════════════════════════════════════════════════════════════════════════════════════
              Only a guild's guild master can claim it here and subscribe it. Neither of
              your guilds has been claimed yet, so the next step is theirs.

              ┌──────────────────────────────────────────────────────────────────────────┐
              │ (◎) OLYMPUS XXVII                          │ Only OLYMPUS XXVII's guild   │
              │     Forever · US                           │ master can claim it and      │
              │                                            │ subscribe it. Send them this │
              │ (●) Obnoxious Yell · Member · Not verified │ link.                        │
              │                                            │ [foreversixty.gg/guild/u…]   │
              │ No one verified yet · 38 on the roster ·   │ [ COPY LINK ] [ SHARE ]      │
              │ No plan                                    │ Not verified yet. Turning up │
              │                                            │ in the guild's logs also     │
              │                                            │ verifies you.                │
              │                                            │ Lead OLYMPUS XXVII on another│
              │                                            │ character? Save that         │
              │                                            │ character with the addon,    │
              │                                            │ then come back. Get set up   │
              └──────────────────────────────────────────────────────────────────────────┘
              ┌──────────────────────────────────────────────────────────────────────────┐
              │ (◎) Employed Millennials   …same shape…                                  │
              └──────────────────────────────────────────────────────────────────────────┘

              A guild missing? Save the character that is in it with the addon. Get set up
```

There is no personal Premium line in the foot because the owner has a grant. `38 on the roster` is illustrative; it renders from `roster_count` or is left out.

Desktop 1280, a mixed account (claim holder of one guild, member of another that is on a plan):

```
   Choose the guild this plan is for.
   ┌──────────────────────────────────────────────────────────────────────────┐
   │ (◎) Ashen Vow · Forever · US              │ You hold Ashen Vow's claim.   │
   │ (●) Kraggor · Guild master · ✓ Verified   │ [ START THE 7-DAY TRIAL ]     │  the one fill
   │ 18 verified · 24 on the roster · No plan  │                               │
   └──────────────────────────────────────────────────────────────────────────┘
   ┌──────────────────────────────────────────────────────────────────────────┐
   │ (◎) OLYMPUS XXVII                         │ OLYMPUS XXVII is on the Raid  │
   │ (●) Mirelle · Member · ✓ Verified         │ guild plan, billed by Bow     │
   │ 31 verified · 38 on the roster ·          │ Jackzon. 2 of 3 Premium seats │
   │ Raid guild plan                           │ are given. Ask Bow Jackzon    │
   │                                           │ for one.   Open the guild     │
   └──────────────────────────────────────────────────────────────────────────┘
```

Phone 360, the owner's account:

```
┌──────────────────────────────┐ full bleed
│ Subscribe a guild            │ (CHECKOUT label hidden below 768)
│ Guild  $120 a year [FOUNDER]  │
│ Up to 25 verified members. A │
│ 10-person raid team fits, …  │
│ First 7 days free ·          │
│ Change plan                  │
├──────────────────────────────┤
│ Only a guild's guild master  │
│ can claim it here and        │
│ subscribe it. Neither of     │
│ your guilds has been claimed │
│ yet, so the next step is     │
│ theirs.                      │
│┌────────────────────────────┐│
││(◎) OLYMPUS XXVII           ││
││    Forever · US            ││
││(●) Obnoxious Yell · Member ││
││    · Not verified          ││
││No one verified yet · 38 on ││
││the roster · No plan        ││
││────────────────────────────││
││Only OLYMPUS XXVII's guild  ││
││master can claim it and     ││
││subscribe it. Send them this││
││link.                       ││
││[foreversixty.gg/guild/us/…]││
││[ COPY LINK ]  [ SHARE ]    ││  ends above y 640
││Not verified yet. Turning … ││
││Lead OLYMPUS XXVII on …     ││
│└────────────────────────────┘│
```

At 360 the band is about 150 px and the lead about 120 px, so the copy buttons land near y 600. If the build measures past 640, the size line moves below the lead as a 13 px muted line in the body (never removed), before anything else changes.

Phone 360, empty state A:

```
│ Subscribe a guild            │
│ Guild  $120 a year            │
│┌────────────────────────────┐│
││We do not know your guild   ││
││yet.                        ││
││Your guilds come from your  ││
││characters. Save a character││
││with the addon on Get set   ││
││up, and its guild shows     ││
││here.                       ││
││[        GET SET UP       ] ││
│└────────────────────────────┘│
```

Account Plan panel, the owner (1440, inside today's `StatePanel`):

```
 PLAN
 Premium · granted, no end date
 Given by Forever Sixty. Nothing is billed to you.
 Server sims            On · 4 runs this month
 History charts         Open on your main character
 Discord notifications  On
 Supporter mark         Shown on your characters
                                                   (no billing button)
```

### 5.9 States of every interactive element

| Element | Default | Hover | Focus-visible | Active | Disabled | Loading | Error |
|---|---|---|---|---|---|---|---|
| `ready` button | gold fill (gold outline when two or more are `ready`) | `--gold-hover` fill; `--card-top` ground on outline | 2 px `--gold` outline, 2 px offset | outline `#171d2b` ground | while another card is redirecting | `aria-busy`, label `Redirecting to secure checkout…`; the other card buttons are disabled | the card's alert line (§9.1), focus on it |
| Claim, start and confirm openers | gold outline | `--card-top` | as above | `#171d2b` | while another confirm is open | n/a | n/a |
| Confirm commit | gold fill | `--gold-hover` | as above | as above | while busy | `aria-busy`, `Claiming…`, `Starting…` or `Confirming…` | alert line inside the block, focus on it; the block stays open |
| `Cancel` text button | underlined `--text` | `--strong` | as above | n/a | while busy | n/a | n/a |
| `CopyLinkField` | `GuildInviteLink` today | today | today | today | n/a | n/a | `Could not copy. Select the link and copy it.` |
| `Manage billing` | `SECONDARY_BUTTON_FIXED` | today | today | today | n/a | `aria-busy` | `Billing did not open. Try again.` |
| `Check again` | underlined `--text` | `--strong` | as above | n/a | while busy | `aria-busy` | the load error line |
| Inline links (`Switch to Raid guild`, `Switch to Guild`, `Get set up`) | `INLINE_LINK` (gold, underlined) | `--gold-hover` | as above | n/a | n/a | n/a | n/a |
| `Change plan`, `Open the guild` | underlined `--text` | `--strong` | as above | n/a | n/a | n/a | n/a |

Touch: every control is at least 44 px tall at 360. Keyboard: tab order follows the card order, then the foot. The step sentence is the action's `aria-describedby`. The verify mark is text, never colour alone. Status lines use `role="status"` and refusals `role="alert"`.

## 6. Join, verify, claim: end to end

### 6.1 The hand-off link

- To an unclaimed guild's guild master or officers (`ask_leader`, `claim_pending_yours`, `claim_pending_ask`): `https://foreversixty.gg/guild/{region}/{ruleset}/{slug}/claim?plan={plan}&interval={interval}`.
- To the claim holder (`claimed_ask`): `https://foreversixty.gg/guild/{region}/{ruleset}/{slug}/settings?plan={plan}&interval={interval}#billing`. The settings page preselects that plan and interval in the launch pricing §3.2.2 chooser.
- The link carries no account id, no token and nothing that names the sender. `plan` and `interval` are parsed against the same closed sets as `parseCheckoutParams`; anything else is dropped silently.
- Share sheet copy: §9.1.

### 6.2 Claim page (`GuildClaim.svelte`) changes

- **With `plan`**: a 14 px `--text` line under the heading, on a 2 px `--gold` left rule, padding-left 14 (§9.3), only while the viewer does not hold the claim (gate fix, row 16: "Claim the guild first" is never shown to the player who just claimed it).
- **After a confirmed claim with `plan`**: the `Next` list gains, first, the §5.3 `ready` button (`Start the 7-day trial` or `Subscribe · {price}`) as the page's one fill. `Approve your officers and members` and `Make an invite link` follow, as today, 12 px apart (the button's failure line is drawn only after a failed click).
- **Signed in, no character of theirs in the guild**: `notAMember` gains the inline link `Get set up`, with `next` set to this claim URL. `/setup` accepts `next` (validated by `safeNextPath`) and, after a save, shows `Back to {Guild}'s claim` as its one primary.
- **A member viewer**: `memberCannotClaim` is replaced by the §9.3 sentence and a `CopyLinkField` holding the claim URL with the same `plan` it was opened with (sr-only label, max 640 px, full-width buttons below 768; §5.5).
- **Signed out**: today's `signInLine`; the sign-in keeps `next`.

### 6.3 Guild page and invite

- The unverified note (`Guild.svelte`, `guildHomeCopy.notVerifiedNote`) picks by `viewer.verify.has_approver` (§8). With an approver: today's sentence. Without one: the §9.3 sentence and the inline gold link `Send your guild master a link` to the claim page.
- The header claim slot for a member of an unclaimed guild reads `Get the guild claimed` and links to the claim page. It no longer offers `Claim this guild` to someone who cannot claim.
- Invite link, joined state: under `You joined the guild.`, 13 px muted, the guild's plan sentence (§9.3), then `Open the guild`.

### 6.4 Guild settings `#billing`, an officer who does not hold the claim

Under today's `Only the guild master can subscribe the guild, release the claim or change billing.` (gate spec §5.2), when the guild has no live plan: `Send them this link.` and a `CopyLinkField` holding the §6.1 settings link, with `plan` set to the plan that fits the verified count and `interval=yearly`.

### 6.5 Account page guild rows (`account-guild-row`)

After the guild name: `Not verified`, 13 px muted, when the account has no verified character there. When `subscribe.step` is not `on_plan`: `{Guild} has no plan.` 13 px muted and the inline link `See the next step` to the checkout guild step with `plan={fits}&interval=yearly` and `#guild-{id}`, where `{fits}` is `raid_guild` above 25 verified members, else `guild`. On a plan: today's `AccountGuildPlanLine`. Polish 2026-10-29: the row is a card block (name in the heading face, a plan `Tag` (`Guild`, `Raid guild` or `No plan`), a warn `Not verified` tag, then settings rows: `Share with this guild` as a `Select`, `Guild settings` (value = the date, `Renews 31 Dec 2026`; the tier is only in the tag) or `See the next step` (sub = the no-plan sentence) as link rows, `Leave guild` as `BUTTON_DANGER`, the Unlink Discord size).

### 6.6 Discord bot

Every bot `Subscribe the guild` and `Move to Raid guild` link (discord-bot spec, `/readiness` and the read-only rows) points to `https://foreversixty.gg/premium/checkout?plan=raid_guild&interval=yearly&guild_id={id}`. The checkout step resolves for whoever opens it (§5.6). The bot's sentences are unchanged.

### 6.7 The paths, step by step

A member gets their guild subscribed (the owner's case):
1. The member opens the checkout step. Their card says only the guild master can claim, and gives them a link.
2. They send the link (Discord, guild chat, a whisper).
3. The guild master opens it and signs in.
4. If the site does not know their guild-master character yet, the claim page's `Get set up` has them save it with the addon, then brings them back.
5. The claim page shows `Claim this guild`. A guild-master claim is confirmed at once and verifies them.
6. `Next` shows `Start the 7-day trial`. Stripe. The guild is on the plan.
7. The guild master approves members from the roster, or rotates and shares an invite link; members who turn up in the guild's logs are verified too. The member is now verified, and their checkout card reads `on_plan` with its seat line.

An officer of an unclaimed guild: `Start a claim` on the card. The claim is pending. The card gives a copy link for a second officer or the guild master. A second officer presses `Confirm the claim` (on their own card, the claim page, or the guild page slot), or the guild master's next export confirms it (`AutoConfirmClaimIfPending`). The claimant's card turns `ready`.

No state on these paths ends on a page without a next step. The one wait we cannot remove is a person (the guild master or a second officer), and the page gives the member the link to send them.

## 7. Phone and desktop

- **Never collapses:** the plan name and price, the founder pill, the size line, the lead sentence, any step sentence, the copy field's URL (it truncates but stays selectable and copyable), the verify mark, the other-character line.
- **Collapses:** below 1024 the plan summary leaves the h1's row for its own row. Below 768 the `CHECKOUT` label hides; the card goes to one column with the step block under a rule; buttons go full width; at most two own characters show, then `+{n} more`.
- **Wide:** at 1440 and 2000 the band is full bleed with its content in the centred 1344 column; the body is the centred 880 column; cards never stretch past 880.
- **Heights (ready):** a card is 156 px desktop with a two-line step and a button, up to 240 px with the copy field and both 13 px lines; 260 to 360 px on phone. The skeleton reserves 156 / 260, and the bound is CLS of at most 0.02. Real cards carry no minimum height; each sizes to its content (gate fix G1, 2026-10-28). The page's pre-paint height reserve (at least the viewport) holds only while the step loads, and the site footer is not drawn meanwhile, so a card taller or shorter than its skeleton never moves anything on screen; once rows or an empty state render, the reserve is released and the footer is first drawn at its final place (G2). The load error keeps the reserve and only shows the footer, below the fold: released there, the footer would be first drawn in view seconds after load and become the page's LCP element (Lighthouse's signed-out run reaches this state; measured 3039 ms against the 2400 ms cap).

## 8. API this spec needs

`GET /v1/me`, per `guilds[]` entry. These keys are new. The web treats an older response as unknown, and draws a card with only the identity rows and `Open the guild`.
- `ruleset_label` (the ruleset word: Normal, PvP, Hardcore — shipped in place of the `realm` the first draft named; the web renders `{ruleset} · {REGION}`), `slug` (the guild page's URL slug, sent raw; the web escapes it), `crest_url` (string or null), `faction` (`alliance`, `horde` or null), `roster_count` (the home's `counts.on_roster + counts.in_game_no_account`; null without a snapshot).
- `claim`: `{ state: 'unclaimed'|'pending'|'claimed'|'contested', yours: boolean, pending_yours: boolean, holder_name: string|null }`. `holder_name` is the claim holder's or pending claimant's character in the guild, sent only to a verified officer of the guild; null otherwise.
- `verify`: `{ has_approver: boolean }`, true when the guild has at least one verified officer or leader.
- `covered`: null, or `{ tier, state: 'trial'|'active'|'ending'|'past_due'|'lapsed', ends_at: string|null, billed_by: string, you_pay: boolean, you_hold_seat: boolean, seats_given: number, seats_total: number }`, for any account with a character in the guild. `billed_by` is empty for a grant, and for a viewer with no verified character there. `lapsed` stays for 90 days after the end. This is separate from today's officer-only `plan`, which keeps its shape.
- `subscribe`: `{ step, trial_used: boolean|null, available_at: string|null, reason: 'holds_other'|'rate_limited'|null, other_guild_name: string|null }`. `step` is one of `on_plan`, `checkout_open`, `frozen`, `ready`, `claimed_ask`, `claim_pending_yours`, `confirm_claim`, `claim_pending_ask`, `claim`, `claim_link_bnet`, `start_claim`, `claim_blocked`, `ask_leader`, resolved first match in that order. It is computed in one Go function beside `guilds.Capabilities`, from the predicates `billing.checkGuildCheckout` and `guilds.Claim` already use. `trial_used` is set only on `ready`. `available_at` is set only on `claim_blocked` with `rate_limited`.
- `verified_count` keeps its meaning: one count per player, under the `Verified members` rule.

`GET /v1/me` `entitlements`:
- `billing.source`: `stripe` or `grant`. For a grant, `current_period_end` is the grant's end, or null.
- `personal_portal`: boolean, true when the account has a personal entitlement row with `source = 'stripe'` in any status (the portal has something to show).

`GET /v1/guilds/{id}/home` `viewer` gains `verify.has_approver` (for §6.3).

`POST /v1/billing/checkout`, personal 409 `conflict`, gains `fields.source`: `stripe`, `grant` or `seat`.

Web-only additions: `localStorage['fs.premiumSource']`; `localStorage['fs.meGuildCount']`; `plan` and `interval` on the claim and settings pages; `next` on `/setup`; `upgrade=1` on settings; `guild_id` allowed on the guild step without auto-launch when the guild is not `ready`.

No new tables. No change to who may claim, verify or bill.

## 9. Copy, verbatim

Prices, dates and counts come from `plans.ts`, `GUILD_TIER_MEMBER_LIMIT` and the API; `{…}` marks a value. Dates use `d Mon yyyy` (`18 Nov 2027`). Guild names render as stored (`OLYMPUS XXVII`). A possessive on a name ending in s takes the apostrophe alone: `{Guild}'s` reads `Employed Millennials'` (gate fix, 2026-10-28; `web/src/lib/possessive.ts`). The figure 25 is always `GUILD_TIER_MEMBER_LIMIT`, never typed.

### 9.1 Checkout guild step

| Key | Copy |
|---|---|
| Band label | `CHECKOUT` |
| h1 | `Subscribe a guild` |
| h1, invalid link | `Checkout` (ruling 2026-10-27: invalid link h1) |
| Plan summary | `{Plan}` · `{amount}` + ` a year` or ` a month`; `Founder` (pill); `First 7 days free`; `Change plan` |
| Size line, Guild | `Up to 25 verified members. A 10-person raid team fits, and a 20-person one with a small bench.` |
| Size line, Raid guild | `Any size. For guilds over 25 verified members.` |
| Lead, one or more `ready` | `Choose the guild this plan is for.` |
| Lead, none `ready`, every guild unclaimed | one guild: `Only a guild's guild master can claim it here and subscribe it. Your guild has not been claimed yet, so the next step is theirs.`; two: `Only a guild's guild master can claim it here and subscribe it. Neither of your guilds has been claimed yet, so the next step is theirs.`; three or more: `Only a guild's guild master can claim it here and subscribe it. None of your guilds has been claimed yet, so the next step is theirs.` |
| Lead, none `ready`, a card the viewer can claim from (`claim`, `start_claim`, `confirm_claim`, `claim_link_bnet`, `claim_pending_yours`) | `You can claim a guild here. Each card below has its next step.` |
| Lead, none `ready`, some claimed and some not | `Only a guild's guild master can claim it here and subscribe it. Here is the next step for each of your guilds.` |
| Lead, none `ready`, every guild claimed or on a plan | `The guild master who holds a guild's claim subscribes it. Here is where each of your guilds stands.` |
| Rank words | `Guild master`; `Officer`; `Member` |
| Verify mark | `Verified`; `Not verified` |
| More characters | `+{n} more` |
| Facts | `{n} verified`; `No one verified yet`; `{m} on the roster`; `No plan`; `Guild plan`; `Raid guild plan`; `{Tier} plan · trial`; `{Tier} plan · ends {date}`; `Plan ended {date}` |
| `ready`, trial unused | `You hold {Guild}'s claim.`; button `Start the 7-day trial` |
| `ready`, trial used | `You hold {Guild}'s claim.`; button `Subscribe · {price}` |
| `ready`, lapsed | `{Guild}'s plan ended {date}. You hold its claim.`; button `Subscribe again · {price}` |
| `ready`, Guild chosen, over 25 verified | `{n} verified members, more than Guild's 25.` + `Switch to Raid guild` |
| `ready`, Raid guild chosen, 25 or fewer verified | `{n} verified members. Guild covers up to 25 for less.` + `Switch to Guild` |
| `claim` | `You lead {Guild}. Claim it here, then subscribe it.`; button `Claim {Guild}` |
| `claim` confirm | `Claim {Guild}? You may attempt one claim every 30 days, hold one claimed guild at a time, and a claim can be contested.`; `Yes, claim it`; `Cancel` |
| `claim` done | `Claimed. You hold {Guild}.` |
| `claim_link_bnet` | `You lead {Guild}. Claiming as guild master needs your Battle.net account linked.`; button `Link Battle.net` |
| `start_claim` | `Nobody holds {Guild}'s claim yet. As an officer you can start one; a second officer or the guild master confirms it.`; button `Start a claim` |
| `start_claim` confirm | `Start a claim on {Guild}? A second officer or the guild master confirms it. You may attempt one claim every 30 days.`; `Yes, start the claim`; `Cancel` |
| `start_claim` done | `Claim started. It waits for a second officer or the guild master.` |
| `claim_pending_yours` | `Your claim on {Guild} waits for a second officer or the guild master. Send them this link.` |
| `confirm_claim` | `{Name} has started a claim on {Guild}. As an officer you can confirm it.`; button `Confirm the claim` |
| `confirm_claim` confirm | `Confirm {Name}'s claim on {Guild}? They will run the guild here and can subscribe it.`; `Yes, confirm it`; `Cancel` |
| `confirm_claim` done | `Confirmed. {Name} holds {Guild}.` |
| `confirm_claim`, no `holder_name` (gate fix, 2026-10-28: a rank word never stands in for `{Name}`) | `An officer has started a claim on {Guild}. As an officer you can confirm it.` |
| `confirm_claim` confirm, no `holder_name` | `Confirm the claim on {Guild}? The officer who started it will run the guild here and can subscribe it.` |
| `confirm_claim` done, no `holder_name` | `Confirmed. The officer who started the claim holds {Guild}.` |
| `claim_pending_ask` | `An officer has started a claim on {Guild}. Once a second officer or the guild master confirms it, they can subscribe it. Send them this link.` |
| `claimed_ask`, viewer is an officer | `{Name} holds {Guild}'s claim and subscribes it. Send them this link.` |
| `claimed_ask`, viewer is a member | `The guild master who holds {Guild}'s claim subscribes it. Send them this link.` |
| `claimed_ask`, viewer holds the claim (`claim.yours` true; controller ruling 2026-10-27) | `You hold {Guild}'s claim. Subscribing it needs a verified officer or guild-master character.` |
| `ask_leader` | `Only {Guild}'s guild master can claim it and subscribe it. Send them this link.` |
| `ask_leader`, other-character line | `Lead {Guild} on another character? Save that character with the addon, then come back.` + `Get set up` |
| `ask_leader` and `claimed_ask`, moderator fallback | `Guild master gone? Ask us on Discord and we will sort the claim by hand.` (link text `Ask us on Discord`, to the site Discord invite) |
| `claim_blocked`, `holds_other` | `You already hold {Other}'s claim. One account holds one guild's claim, so an officer of {Guild} claims it instead.` |
| `claim_blocked`, `rate_limited` | `You made a claim in the last 30 days. You can claim {Guild} from {date}.` |
| `frozen` | `{Guild}'s claim is contested. Nobody can subscribe it until a moderator resolves the contest.` + `Open the guild` |
| `checkout_open` | `Someone already has a checkout open for {Guild}. It closes on its own if they do not finish.` + `Check again` |
| `on_plan`, you pay, chosen tier is current or lower | `{Guild} is on the {Tier} plan. You pay for it.` + `Manage billing` |
| `on_plan`, you pay, on Guild, chose Raid guild | `{Guild} is on the Guild plan. You pay for it.` + `Move to Raid guild` |
| `on_plan`, someone else pays, named | `{Guild} is on the {Tier} plan, billed by {Name}.` |
| `on_plan`, no payer name | `{Guild} is on the {Tier} plan.` |
| `on_plan`, seat lines | `You hold one of its Premium seats.`; `{given} of {total} Premium seats are given. Ask {Name} for one.`; `{given} of {total} Premium seats are given. Ask the guild master for one.` (no payer name); `All {total} of its Premium seats are given.` |
| `on_plan`, link | `Open the guild` |
| Verify sub-lines | approver exists: `Not verified yet. An officer approves you from the roster, the guild's invite link verifies you at once, or you are verified when you turn up in the guild's logs.`; no approver: `Not verified yet. You are verified when you turn up in the guild's logs, or once the guild master claims {Guild} and approves you.`; short form: `Not verified yet. Turning up in the guild's logs also verifies you.`; claimed, no approver: `Not verified yet. You are verified when you turn up in the guild's logs, or once the claim holder approves you.` |
| Copy field | sr label `Link for {Guild}'s guild master`, `Link for {Name}` or `Link for a second officer`; `Copy link`; `Copied`; `Share`; `Could not copy. Select the link and copy it.` |
| Share sheet | title `Subscribe {Guild} on Forever Sixty`; text, claim link: `{Guild} could use the {Plan} plan on Forever Sixty. Claim the guild here, then subscribe it.`; text, settings link: `{Guild} could use the {Plan} plan on Forever Sixty. Subscribe it from the guild's billing.` |
| Busy | `Redirecting to secure checkout…`; `Claiming…`; `Starting…`; `Confirming…` |
| Claim errors (from the API's error codes, never its sentence) | already claimed: `Someone claimed {Guild} a moment ago.` (the card re-reads); pending: `An officer started a claim on {Guild} a moment ago.` (re-reads); not eligible: `Your character in {Guild} is not an officer or the guild master.`; (no-Battle.net line removed 2026-10-27: the API answers 403 forbidden for both cases; the claim_link_bnet step covers it); rate limited: `You made a claim in the last 30 days.`; holds another: `You already hold another guild's claim.`; same account: `You started this claim, so a second officer confirms it.`; anything else: `That did not work; try again.` |
| Checkout errors | 503: `Purchases are not open yet. Check back soon — nothing was charged.`; 409 on a guild: `{Guild} is already on a plan.` (the card re-reads); 403: the card re-reads `/v1/me` and shows its new step, never the API sentence; anything else: `That did not work; try again.` + `Try again` |
| Portal error | `Billing did not open. Try again.` |
| Loading | sr `Loading your guilds` |
| Load error | `Your guilds did not load.` + `Try again` |
| Empty A | title `We do not know your guild yet.`; body `Your guilds come from your characters. Save a character with the addon on Get set up, and its guild shows here.`; button `Get set up` |
| Empty B | title `None of your characters is in a guild we know.`; body `If you joined a guild since, save your character with the addon again. An officer's invite link also adds you at once.`; button `Get set up` |
| Foot, missing guild | `A guild missing? Save the character that is in it with the addon.` + `Get set up` |
| Foot, personal Premium (only `none`) | `Or Premium for yourself, {price}.` + `Get Premium` |
| Invalid link (today) | `Something is wrong with that link. Start again from the premium page.` + `Premium page` |

### 9.2 "You already have it"

| Key | Copy |
|---|---|
| Checkout h1, premium | `Premium` |
| Panel title | `You already have Premium.` |
| Plan lines and line 2 | §4.3, verbatim |
| Grant lines | `Premium · granted, no end date`; `Premium · granted until {date}`; `Given by Forever Sixty. Nothing is billed to you.`; `Given by Forever Sixty. It ends on its own; nothing is billed to you.` |
| Fallback line | `Premium` |
| Buttons | `Manage billing`; `Billing history`; `See your plan` |
| `/premium` card line | `You have Premium through your subscription.`; `You have Premium through a seat from your guild.`; `You have Premium through a grant from Forever Sixty.` |
| `/premium` card button | `See your plan` |

### 9.3 Claim, guild page, settings, account, invite

| Key | Copy |
|---|---|
| Claim page, with `plan` | `This link is for subscribing {Guild} to the {Plan} plan ({price}). Claim the guild first; then the plan is one step.` |
| Claim page, member | `You are a member of {Guild}, so only its guild master can claim it. Send them this page's link.` |
| Claim page, not a member | today's `notAMember` + `Get set up` |
| `/setup` after a save with `next` to a claim | `Back to {Guild}'s claim` |
| Guild page, unverified, no approver | `You are not verified yet. Nobody in {Guild} can approve members until its guild master claims it.` + `Send your guild master a link` |
| Guild page claim slot, member of an unclaimed guild | `Get the guild claimed` |
| Settings, not the claim holder, no plan | today's sentence + `Send them this link.` |
| Account guild row | `Not verified`; `{Guild} has no plan.`; `See the next step` |
| Invite joined | `{Guild} is on the {Tier} plan.`; `{Guild} has no plan yet.` |

## 10. Hierarchy

- Band: the h1 is primary. The plan summary is secondary but always visible, because the price is what the player came to pay.
- Body: the lead sentence frames the list. Each card's step block is the card's primary element. The page has at most one gold fill: the `ready` button when it is the only one, or `Get set up` in Empty A. Hand-off cards have no fill; the copy button is the action and stays neutral. An opened confirm block's commit button is the one fill while it is open; the resting page keeps at most one (ruling 2026-10-27).
- Inside a card, the guild name is the strongest type, the viewer's character second, and the facts are muted.
- Secondary lines (the other-character line, the verify sub-line, the foot) are 13 px muted and never above the action.

## 11. Performance and polish

- `/premium/checkout` stays a single island (`CheckoutLauncher`). The guild step makes no `/home` reads; crests, faction and counts come from `/v1/me`.
- CLS of at most 0.02 at 390 and 1440 for the guild step from loading to ready, the premium "already have it" panel, and the account Plan panel. `/premium` CLS is unchanged whether `data-premium` is set or not.
- No crest flash (§5.7). `ClassCrestRing` keeps its reserved 20 px box.
- `web/lighthouserc.json`: add `http://localhost/premium/checkout.html?plan=guild&interval=yearly` to the collect list. The default assertion row applies (performance and accessibility at least 0.95, LCP at most 2400 ms, TBT at most 100 ms, CLS at most 0.05). It runs signed out, so it measures the redirect shell; the signed-in states are held by the CLS e2e checks above.
- Focus never moves on first render. It moves only after an in-page change (claim done, retry, the panel replacing the redirect message), as §4.2, §5.3 and §5.7 say.

## 12. Acceptance

**E2E** (`web/tests/e2e/checkout-guild.spec.ts` new; `premium.spec.ts` and the account spec extended; API stubbed):
1. The owner fixture (capture 1): two cards, OLYMPUS XXVII then Employed Millennials. Each shows the owner's character in it with `Member · Not verified`, `No one verified yet`, `No plan`, the `ask_leader` sentence, a copy field whose value is the claim URL with `plan=guild&interval=yearly`, the short verify line, the other-character line, and the Discord fallback line linking to the site invite. The lead reads the two-guild form. The foot has no Premium line (grant).
2. `Copy link` writes the URL to the clipboard and reads `Copied` for 2 s; a refused clipboard shows the failure line.
3. A `ready` guild with the trial unused shows `Start the 7-day trial` as the only filled button; clicking posts checkout with plan, interval and guild id.
4. Two `ready` guilds: both buttons are gold outline; the page has no fill.
5. `claim`: the confirm opens in place with focus on `Cancel`; Escape closes it and returns focus; `Yes, claim it` posts the claim, re-reads `/v1/me`, and the card becomes `ready` with focus on its button and the status line `Claimed. You hold {Guild}.`
6. `start_claim`, then pending: the card becomes `claim_pending_yours` with a copy field.
7. Claim refusals (already claimed, not eligible, rate limited, holds another) each show their §9.1 line; no API sentence appears anywhere on the page.
8. `claimed_ask` as an officer names the holder; as a member it does not.
9. `on_plan`: you pay (`Manage billing`); you pay on Guild with Raid guild chosen (`Move to Raid guild` links to settings with `upgrade=1`); seat held; seats free; seats full.
10. `checkout_open` and `frozen` show their sentences and no buy button.
11. `guild_id` of a non-ready guild puts that card first and scrolls it into view; of a ready guild it launches Stripe as today; an unknown id shows the list.
12. Plan size lines: Guild shows the 10-person line; a `ready` guild with 12 verified on Raid guild shows `Switch to Guild`, and with 28 verified on Guild shows `Switch to Raid guild`; switching changes the band, the button and the URL without a reload. No text on the page contains `40-person` or `40 players`.
13. Empty A and Empty B show `Get set up` with `next` back to this URL; neither mentions Battle.net.
14. `plan=premium` for each source (own renewing, own ending, past due, seat, grant open-ended, grant until): the panel's title, line, line 2 and button match §4; `POST /v1/billing/checkout` is never called. A stubbed 409 with `fields.source=grant` renders the grant panel.
15. Account: the grant fixture reads `Premium · granted, no end date` and has no billing button; a seat holder with `personal_portal` sees `Billing history`; any fixture whose Premium rows are On never reads `Not on Premium.`
16. `/premium` with `fs.premiumSource=grant` stored before load: the Premium card reads the grant line and `See your plan` on first paint; the card's height equals the signed-out card's.
17. Guild page, unverified, no approver: the new note and link. Claim page as a member: the copy field. Claim page with `plan`: the plan line, and after a claim the subscribe button first in `Next`.
18. Axe is clean in every state above; every control is at least 44 px tall at 360; CLS is at most 0.02 for the guild step from loading to ready at 390 and 1440.

**API tests** (Go, beside `capabilities_test.go`): one table case per `step` value with the guild and account facts that produce it, including the owner's facts (`ask_leader` for both guilds); `covered` hides `billed_by` from a viewer with no verified character; `billing.source` and `personal_portal` for a grant with no Stripe row; the personal 409 carries `fields.source`.

**Screenshots** (fixture build; x = required). ux-designer and wow-player review every row on the rendered page.

| # | State | 360 | 390 | 1440 | 2000 |
|---|---|---|---|---|---|
| 1 | **The owner's account, exactly as production holds it**: signed in, header `Obnoxious Yell` (Fury Warrior, 21), guilds OLYMPUS XXVII and Employed Millennials, rank member in both, not verified, 0 verified members in each, unclaimed, no plan, open-ended operator Premium grant, no Stripe customer. `/premium/checkout?plan=guild&interval=yearly` | x | x | x | x |
| 1b | The same account after an addon save that makes the owner guild master of OLYMPUS XXVII (Decision 11): that card reads `claim` and sorts first; Employed Millennials stays `ask_leader`; the lead reads `Choose the guild this plan is for.` only once a card is `ready`, so here it reads the can-claim lead (`You can claim a guild here. Each card below has its next step.`, gate round 2). Also the card after `Yes, claim it` (`ready`, `Start the 7-day trial`) | x | x | x | x |
| 2 | Same account, `/premium/checkout?plan=premium&interval=yearly` (grant panel) | x | x | x | |
| 3 | Same account, `/account#plan` (grant headline, rows On, no billing button) | x | x | x | |
| 4 | Same account, `/premium` (Premium card in its grant state) | x | x | x | |
| 5 | Mixed account: one `ready` (trial), one `on_plan` member with seats free | | x | x | x |
| 6 | `claim` with the confirm open; the same card after success (`ready` plus status line) | x | x | x | |
| 7 | `start_claim`; `claim_pending_yours`; `confirm_claim`; `claim_pending_ask` | | x | x | |
| 8 | `claimed_ask` as officer and as member; `frozen`; `checkout_open`; `claim_blocked`, both reasons | | x | x | |
| 9 | `on_plan`: you pay; you pay on Guild with Raid guild chosen; seat held; all seats given | | x | x | |
| 10 | `ready` with `Switch to Guild` (12 verified, Raid guild chosen) and `Switch to Raid guild` (28 verified, Guild chosen) | | x | x | |
| 11 | Empty A and Empty B | x | x | x | |
| 12 | Loading skeleton and load error | | x | x | |
| 13 | Focus-visible on a `ready` button, a confirm `Cancel` and `Copy link`; hover on `Copy link` | | | x | |
| 14 | `plan=premium` for own renewing with founder, own ending, past due, seat | | x | x | |
| 15 | Account Plan panel: own, seat with `Billing history`, grant until a date, fallback, free | x | x | x | |
| 16 | Claim page: member with copy field; with `plan` before and after claiming. Guild page unverified note with no approver | x | x | x | |
| 17 | `GuildInviteLink` before and after the `CopyLinkField` extraction, pixel diff | | x | x | |

## 13. Decisions (2026-10-27)

Ruled by the owner through the coordinator on 2026-10-27. Each is applied in the sections named.

1. **No member-started claims.** Only an officer or guild-master character starts a claim (`eligibleClaimRank`, unchanged). Members get the copy link (§5.3, §6.1).
2. **Notifying the guild master is deferred.** The anonymous copy link ships. A v2 `POST /v1/guilds/{id}/plan-asks` route (stores who asked, notifies the claim holder or guild master by email or Discord where reachable) is noted for later and is not in §8.
3. **The hand-off link stays anonymous.** It carries no account id, no token and no sender name (§6.1).
4. **Billing stays with the claim holder.** A verified officer without the claim cannot pay (`Capabilities` unchanged; §6.4 keeps today's sentence plus the copy link).
5. **No purchase is offered to seat or grant holders.** Every "you already have it" state shows its source and never a buy button (§4).
6. **Plan visibility as specified.** `covered` (tier, state, seat counts) goes to any account with a character in the guild; the payer's name only to verified members (§8).
7. **Grant wording is `Given by Forever Sixty.`** The grant note is never shown (§4.3, §9.2).
8. **The moderator fallback ships.** On `ask_leader` and `claimed_ask` cards, a 13 px muted line under everything else in the step block: `Guild master gone? Ask us on Discord and we will sort the claim by hand.`, with `Ask us on Discord` as an `INLINE_LINK` to `communityLinks.discordInvite` (`web/src/data/links.json`, the link the account page and header already use), opening in a new tab with `rel="noopener"` (§5.4, §9.1).
9. **The Battle.net rule for claims stays.** A claim gives control of a guild, so a guild-master claim still needs a linked Battle.net account (`ErrNoBattleNetIdentity`). Verification does not (§3.2 S8b, `claim_link_bnet`).
10. **The stored hint on `/premium` ships**, modelled on the account page's `fs.accountReserve` hint in `web/src/lib/account/layout.ts` (§4.4).
11. **Capture 1 is taken in both states:** the owner as a member (today's production data), and after an addon save that makes the owner guild master of OLYMPUS XXVII, where that card reads `claim` (§12, captures 1 and 1b).
