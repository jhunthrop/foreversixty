# Pricing for launch (Premium, two guild tiers, seats, founder offer)

Status: draft for the owner's sign-off, 2026-10-27. Nothing here changes Stripe or the
entitlements code until the owner approves. Live today: Premium $4/$40, Guild $15/$150, and
the guild plan grants every member Premium (`docs/superpowers/specs/2026-09-21-guild-and-premium-proposal.md` §4.1).

## 1. Intent

Price for the audience Forever actually has: players in their 30s and 40s with jobs, who pay
for time saved and dislike being squeezed. Capture more than today's prices do, keep the free
site exactly as generous, and make the guild plan easy for an officer to say yes to.

Inputs: the owner's direction (one launch chance; mature audience; multiple guild tiers are
fine), five adversarial pricing personas (solo player, parse chaser, raid leader, small guild
leader, SaaS analyst), and five feature-gap personas
(`design/reviews/2026-10-27-premium-differentiators/README.md`).

## 2. Goals, non-goals, anti-goals

Goals
- One Premium price everyone pays; no "solo tax".
- Two guild tiers so a 15-person friends guild and a 40-person raiding guild both have a yes.
- The guild plan pays for running the guild; Premium seats are a perk with a fixed count.
- Founder pricing that rewards buying in the first 60 days, monthly as well as yearly.
- Promises that make an officer comfortable putting the plan on their card.

Non-goals (this lane)
- The Premium differentiators (log coaching, account-wide alt view) — their own lanes; this
  lane prices what exists and leaves room for a higher tier later.
- A Raid Leader personal tier (~$12) — a candidate once wipe review exists.
- Multiple raid teams under one plan — a later upsell.

Anti-goals
- Any paywall on the funnel: the addon export, BiS, the planner, browser sims, logs, live
  logging, rankings and guides stay free, stated publicly as never shrinking.
- Per-head pricing, lifetime plans, ads.
- A rule an officer has to explain twice in guild chat.

## 3. The prices

| Plan | Monthly | Yearly | Includes |
|---|---|---|---|
| Free | $0 | $0 | Everything a player needs today; the free list never shrinks |
| Premium | $6 | $50 | Server sims, saved sims and builds kept, 2-year retention, history charts, compare more than two builds, Discord notifications, supporter mark |
| Guild | $15 | $150 | Up to 25 verified members: raids and sign-ups (site and Discord), attendance, the raid-night report, the guild Discord channel and bot, guild-wide retention, 1 Premium seat |
| Raid guild | $25 | $250 | Any size: everything in Guild plus loot council, the readiness board and consumables, 3 Premium seats |

Founder prices, for purchases between 4 November 2026 and 3 January 2027, locked while the
subscription stays active (Stripe's card-retry window counts as active):

| Plan | Founder monthly | Founder yearly |
|---|---|---|
| Premium | $5 | $40 |
| Guild | $12 | $120 |
| Raid guild | $20 | $200 |

Net after Managed Payments fees (~7% + 30¢ per invoice), for the page's ordering: yearly
first everywhere, because the 30¢ fee costs small monthly charges most and yearly hedges the
spring population drop.

## 4. Rules

| Situation | What happens |
|---|---|
| Guild tier size | Guild counts verified members at checkout and at each renewal; at 26+ the officer is told before renewal and offered Raid guild. Nothing is switched off mid-period. |
| Who holds a seat | The subscriber or the guild leader gives seats to any verified member. |
| Moving a seat | Once every 30 days per seat. The old holder keeps Premium until the move. |
| A seat holder leaves the guild | The seat returns to the pool at once (the 30-day lock does not apply). |
| A paying Premium player is offered a seat | They accept it first. On accepting, their own subscription is set to cancel at period end and the seat starts when that paid period ends, so nobody pays twice and no refund is needed. |
| A seat is taken away from someone who cancelled their own Premium for it | They get an email with a one-click link to resubscribe at the price they had. |
| Plan lapses | Officer tools become read-only; attendance, loot and reports stay readable forever; CSV export of attendance and loot always available. Seats end at period end. |
| Guild leader changes | The subscriber can transfer the plan to another officer from guild settings; the founder price travels with it. |
| Members chip in | Any member can add a recurring contribution toward the guild's plan from the guild page; contributions reduce the subscriber's invoice; the subscriber sees totals, not names, unless the member opts in. |
| Trial | A guild's first raid night on a plan is free: the plan starts 7 days before the first charge, enough for one Onyxia and its report. One trial per guild. |
| Renewals | Email 7 days before every yearly renewal; one-click cancel; cancellation at period end. |
| Refunds | Goodwill refund in the first 14 days of any first charge; otherwise no partial-year refunds except the seat case above. |

## 5. What changes in the code

- **Stripe** (`api/cmd/api/stripe_setup.go`, `api/internal/billing`): new products and lookup
  keys for Guild and Raid guild, founder prices (monthly and yearly per plan), the trial
  period, and the portal configuration. The old $4/$15 prices stay archived; no live
  subscription exists yet, so no migration of subscribers.
- **Entitlements** (`api/internal/entitlements`): a guild plan no longer grants Premium to
  every member. Plans gain a tier (`guild` / `raid_guild`); `officer_views` features split by
  tier (loot council, readiness and consumables for `raid_guild`); a `premium_seats` table
  (guild, holder, assigned_at) with the 30-day move rule; seat holders resolve as Premium.
- **Guild settings**: seats list, transfer the plan, contributions view, export CSV.
- **Guild page**: chip-in for members.
- **Premium and pricing pages**: the new table, founder badge until 3 January, the free promise.
- **Reconcile job**: tier size checks before renewal, seat expiry, founder lock checks.

## 6. Open for the owner

1. The founder window end: 3 January 2027 (60 days) or the first raid night plus 4 weeks.
2. Chip-in: build for launch, or right after (it is the largest piece here).
3. Trial: 7 days, or "until the guild's first matched raid log, at most 14 days".

## 7. Left to the builder

Product and lookup key names, table names beyond those given, email wording within the site's
copy rules, the seat UI within the guild settings design system.

## 8. Decisions from the builder check (2026-10-27, the controller's calls)

1. **Chip-in is deferred until after launch** (recommended answer to open question 2). Stripe
   has no cross-customer credit and Stripe is merchant of record under Managed Payments, so a
   member's contribution is a sale to that member; the mechanism needs its own design.
2. **Founder prices are separate price objects** per plan and interval, offered by Checkout
   only while the window is open (the window is measured at Checkout session creation) and
   kept active, unlisted, afterwards. Founder customers get their own portal configuration
   listing only founder prices, so monthly↔yearly keeps the lock. The lock lasts while the
   subscription stays continuous; a lapse and resubscribe pays list price. Guild → Raid guild
   for a founder uses an upgrade endpoint of ours (the portal cannot switch products) and
   lands on the founder Raid guild price.
3. **Seats for paying players** — see §4: accept first, own subscription cancels at period
   end, the seat starts then. No automated refunds; the 14-day goodwill refund stays manual
   (the existing refund policy).
4. **Guild size is a soft cap.** Verified members, including the leader, are counted daily.
   Above 25 the officer sees a line in guild settings and gets an email 14 days before
   renewal offering Raid guild. Nothing is switched or cancelled automatically. Upgrading
   mid-period uses our endpoint with Stripe's proration.
5. **Trial** is Stripe's `trial_period_days` on Checkout with the card collected up front;
   it applies to every price (founder and yearly included) and grants the tier's full
   features including seats. The 14-day goodwill window starts at the first charge. "One
   trial per guild" is enforced in our own table, so transfers and new subscribers do not get
   a second one. The length stays open question 3.
6. **Lapse** is derived, not a new state: no plan with access → officer tools read-only;
   seats end when the plan's access ends (`past_due` keeps access as today). **Transfer**
   creates the new officer's subscription at the same price (founder prices are still active)
   starting when the current period ends, and sets the old one to cancel at period end.
7. **No subscribers to migrate**: billing went live 2026-10-24 and no real purchase has gone
   through (memory: forever-stripe-billing); the builder verifies with a Stripe list call
   before archiving. The new prices take new lookup keys; the old keys' prices are archived.
