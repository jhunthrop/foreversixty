# Creator program: experience and system spec

Author: build lane `creator-program`, 2026-11-08, from the owner-approved plan. Mock first (tenet 15): the four surfaces match `design/mocks/ws_creators.py` region by region, and `web/tests/e2e/parity-creators.spec.ts` gates it.
System: `design/specs/2026-11-05-web-layout.md` (the primitives, the split Page, one Plate per page) and the shipped `/admin/links` page (`GoLinksAdmin.svelte`, `ws_billing.py` `admin_links`), whose grammar these pages reuse: PageHeader, `Page layout="split"`, a Plate of Rows with `Figure`s in the trail, a rail Section of `Field`s, `Tag` for a status.
Memory rules applied: stripe-billing (Commish LLC's account, Managed Payments; keys in Secret Manager, never printed), owner-account-grant (the owner is the only admin; never named), no-tool-migrations-ahead-of-deploy (only the deployed API migrates), provenance-quiet, one-job-per-page, design-work-lives-in-repo.

## 0A. ux-designer review of the mocks, 2026-11-08: SHIP WITH FIXES (all MUSTs applied)

- M1 privacy: player rows read `Signed up week of {Monday}`; Returning and Paying show only once the player is 7 days old (the API answers null before); sorted by week; meta `by week · no names, by design`.
- M2 creator Rows: `since start` gone; `Owed now` is the one gold Figure, last in the foot. The column is `Premium net` everywhere.
- M3 the ready Tag is `Payouts ready`.
- M4 payouts Section title is `{n} creators` (meta `{total} to pay`); the month lives only in the Select.
- M5 statement foot: Total first, gold. M6 `Revenue share` is `Share`.
- M7 `Lines` selects the Row and moves focus to the Lines rail Section (on a phone it scrolls there).
- M8 a blocked statement has no action and a `Stripe not ready` warn Tag; section 1 says so.
- M9 refund amounts in the wipe colour; amounts stay on the title line.
- S10 placeholders `e.g. Grimbrew`; S11 `estimated, $500 cap max`; S12 `Showing the newest 100.` when capped; S14 the dialog title names the month: `Pay Grimbrew $168.60 for November 2026?`. S15 rejected: the mock clock is 2026-12-02, so 30 Nov and 1 Dec are past.

## 0. Files

| What | Path |
|---|---|
| Copy (every word on the four pages) | `web/src/lib/creators/copy.ts` |
| Arithmetic the pages show (money, status, statement lines) | `web/src/lib/creators/model.ts` |
| Facts reader (bundles the two modules with esbuild over fixed sample data, never retyped) | `design/mocks/creators_facts.mjs` -> `design/mocks/data/creators-facts.json` |
| Generator (composes only from `ws_prims`; passes `assert_no_surface_styles`) | `design/mocks/ws_creators.py` |
| Mocks | `design/mocks/web/web-admin-creators.html`, `web-admin-creators-payouts.html`, `web-creators-onboarding.html` (invite, done, refresh states stacked: one mock per state below), `web-creators-stats.html` |
| Renders (tracked copies) | `design/reviews/2026-11-08-creator-program/` |
| Regenerate | `node design/mocks/creators_facts.mjs && python3 design/mocks/ws_creators.py`, then `node design/mocks/ws_shoot.mjs <mock> 390,1440` |

## 1. The rules (exact; the API is the only place they are computed)

**A creator's player** is an account that is either
1. **Link:** the account's first tracked link (`users.first_touch_slug`, first touch: the first `/go/` link the browser carried, kept up to 30 days on the device and sent with the first signed-in page, recorded only while the account is under an hour old, exactly as today) is the creator's link, and the click was at most **30 days** before the account was made; or
2. **Code:** the account redeemed the creator's Stripe promotion code at checkout and had no creator yet.

and in both cases: the account was made **on or after the creator's start date**, and it is **not the creator's own account**. One account belongs to one creator only, for good: the first rule that fires wins (a link player who later types another creator's code stays the link creator's).

**A returning player** is a creator's player who
- signed in with Battle.net (`users.bnet_sub` is set; every sign-in is Battle.net since launch),
- has at least one Forever character imported from Battle.net (`characters.source = 'bnet'`),
- and was signed in on a **second UTC calendar day** within **14 days** of the account being made (any signed-in request on a later UTC day than the sign-up day, before `created_at + 14 days`; the API stamps `creator_players.second_day_at` once from the sign-in hook, so nothing per day is stored).

The creator earns **$2.00 once** per returning player, in the month the player qualified (the later of the second day and the statement's check that the character exists).

**Revenue share:** 50% of the **net** of every paid invoice of a subscription billed to the player (Premium, and a Guild or Raid guild plan when the paying officer is the player), for invoices paid within **12 months of the player's first paid invoice**. Net is the payment's Stripe balance transaction `net` (the amount after Stripe's fees), less any tax on the invoice that the fees did not already take, less refunds. A refund lands as a negative line in the month it happens, at 50% of the refunded amount (Stripe keeps its fee on a refund, so the refund line is the refunded amount itself, tax share removed). Invoices of $0 (a trial) are skipped.

**Monthly cap:** default **$500.00** per creator per month, set per creator by the owner. A month's total is `min(cap, max(0, bounties + shares - refunds))`. Anything over the cap is not carried to the next month; a month whose refunds outweigh its earnings pays $0 and the difference is not carried. Both are shown on the statement line, never hidden.

**Statements and approval:** a month is a UTC calendar month. A statement for a creator and a month lists every unsettled line dated before the month's end. Only a closed month can be approved. Approving is a person pressing Approve on `/admin/creators/payouts`; nothing ever pays automatically. Approving settles the lines (they never appear again), records the payout, and creates one Stripe transfer to the creator's Express account with the idempotency key and transfer group `creator-<id>-<yyyy-mm>`; a second press, a retry or a second tab returns the same payout. A $0 statement is settled without a transfer. A creator whose Stripe account cannot receive transfers yet cannot be approved (no action; the Row says why and carries a `Stripe not ready` Tag; the API answers 409). Approving a month also settles any earlier month that was never approved (its lines are dated earlier), under this month's cap.

**Alts and self-referral:** the creator's own site account (bound by the invite, section 3) never counts, by link or by code, and binding it removes it if it was already counted. One account counts once (the attribution row is keyed by account). One Battle.net account is one site account (`users.bnet_sub` is unique), so a player's alts are characters on one account and count once. A character that already qualified a paid player cannot qualify a second account.

**Privacy:** the creator sees counts and anonymous rows (the week of sign-up, a Monday; returning and paying only once the player is 7 days old, so "I subbed today" on stream cannot be matched to a row), never a name, a BattleTag, a character or an email. Nothing about a player is in a URL or a log line: logs carry ids only, as the rest of the API does. The owner's statement lists players by their account number.

## 2. Promo codes and Managed Payments

Each creator gets one Stripe promotion code (the code the creator says on stream, `GRIMBREW`), on one shared coupon `fs_creator`: **10% off the first payment** of any plan (`duration: once`). Checkout already sets `allow_promotion_codes` under Managed Payments (`stripe_gateway.go`), so no checkout change is needed. Verified in the Connect sandbox on 2026-11-08: a coupon, a promotion code and a subscription redeeming it (invoice $5.00 -> $4.50, fee $0.43, net $4.07, a $2.00 partial refund books -$2.00 with no fee back). Not verified: a live redemption under Managed Payments (the sandbox has no tax head office, so it runs without Managed Payments). **For the owner:** the discount (10%, once) is a constant (`creators.CouponPercentOff`); the first live statement should be checked against the Dashboard before its first approval.

## 3. Surfaces

All four pages are `noindex`, outside the sitemap (`/admin/` is already filtered; `/creators/` is added), and use the site's flat system. The two admin pages are the owner's alone: anyone else gets the same one line `/admin/links` shows (signed out: the sign-in prompt; signed in: `This page is for the site's owner.`); the API answers 401/403 whatever the page shows.

### 3a. `/admin/creators` (owner)

PageHeader `Creators`, lead `Who brought players, what they earned this month, and whether Stripe can pay them.` A split Page:
- **Main:** Section `Your creators` (meta `{n} creators · this month so far`) holding the page's one Plate: one Row per creator, title the creator's name, sub `/go/{slug} · code {CODE} · since {date}`, foot: five Figures (Clicks, Signed up, Returning, Premium net, then `Owed now` in gold with `of {cap} cap` or `cap reached` under it); trail: the Stripe status Tag (`Payouts ready` success, `Onboarding` muted, `Needs info` warn, `Not started` muted), then an `Edit` small button. Below the Plate one line link `Monthly statements ›` to the payouts page.
- **Rail:** Section `Add a creator` (or `Edit {name}` when Edit was pressed): Fields Name, Channel (a Twitch or YouTube URL), Link slug (becomes `/go/slug`, new), Promo code (uppercase letters and digits, 4 to 20), Starts on (a date; players who signed up before it never count), Monthly cap (dollars, default 500). Primary `Add creator` / `Save`. In edit mode the slug and code are read-only text (they are live in Stripe and in posts) and two more actions: `Copy onboarding invite` (makes a fresh one-time invite link, 14 days, copies it; the old one stops working) and `Refresh Stripe status`.
- States: loading (Skeleton), signed out, not owner, load error (LoadError with Retry), empty (`No creators yet. Add the first one.`), saving, field errors from the API under each Field, invite copied (`Invite copied. It works once, for 14 days.`), Stripe unavailable (`Stripe isn't configured, so codes and payouts are off.` one line above the Plate).

### 3b. `/admin/creators/payouts` (owner)

PageHeader `Statements`, lead `Each creator's month. Nothing is paid until you approve it.`, action link `‹ Creators`. A split Page:
- **Main:** Section titled with the month (`November 2026`, meta `{total} to pay`), its action the month `Select` (the current month `December 2026 (open)` and the twelve closed months before it; the page opens on the last closed month). The page's one Plate: a Row per creator, title the name, sub the status line (`Ready to approve`, `Month still open`, `Paid {date} · transfer tr_…`, `Settled {date} · nothing to transfer`, `Nothing owed`, `Stripe can't pay yet: onboarding not finished`); trail the action and `Lines`; foot five Figures (Returning with `{n} × $2.00`, Revenue share, Refunds, Over cap, Total). The action is `Approve {amount}` primary (closed, unpaid, a ready account, total over $0), `Settle $0.00` (closed, nothing owed), else none. The selected Row (the first by default) is marked current.
- **Rail:** Section `Lines · {name}`: RowsSlim, `Player #{id} · returning · {date}` `$2.00`, `Player #{id} · paid {date}` `{share}`, `Player #{id} · refund {date}` `−{amount}`; empty `No lines this month.` (tenet 4: the lines are shown, never behind a disclosure).
- Approve asks once in the shared Dialog: title `Pay {name} {amount}?`, line `Stripe transfers {amount} to their account now. This can't be undone here.`, actions `Cancel` / `Pay {amount}`.
- States: loading, signed out, not owner, error, approving (the button busy), approved (the sub turns `Paid …`), approve failed (`Stripe didn't take the transfer: {reason}. Nothing was paid; try again.` under the Plate), month open (no actions).

### 3c. `/creators/onboarding` (the creator; Stripe's return_url and refresh_url)

`Page layout="form"`, PageHeader `Creator payouts`. One Section (`Your invite` while an invite is open, else `Stripe`), one job per state:
- `?invite=<token>` signed out: `Sign in with Battle.net to accept the invite from Forever Sixty.` and the sign-in prompt (the invite rides `next`).
- `?invite=<token>` signed in: `You're joining the creator program as {creator name}.` and primary `Accept and set up payouts`; accepting binds the account and continues to Stripe.
- signed in creator, no params or `?refresh=1`: `Stripe handles your payout details. We never see your bank account.` and primary `Continue to Stripe` (a fresh account link each press; `?refresh=1` presses it for them once, and shows the button if the redirect fails).
- `?done=1`: re-reads the account from Stripe. Ready: `You're set up. Payouts arrive after each month is approved.` and `Your stats ›`. Not ready: `Stripe still needs a few details.` and `Continue to Stripe`.
- Not a creator: `This page is for creators in the program. Ask us for an invite.`; invalid or used invite: `That invite has expired or was already used. Ask us for a new one.`

### 3d. `/creators/stats` (the creator, signed in)

PageHeader `{creator name}`, meta `foreversixty.gg/go/{slug} · code {CODE}`. A split Page. Main: the one Plate holding four Figures (This month, the estimated payout with `estimated · {cap} cap`; Players; Returning; Paying), then Section `Your players` (meta `newest first · no names, by design`) as RowsSlim, newest first, at most 100: `Signed up {date}` with Tags `Returning` (success) / `Paying` (gold). Rail: Section `Payouts`, one RowsSlim line per approved month, `{Month}` `{amount}`; empty `No payouts yet. Each month is paid after it closes and is approved.` While Stripe is not ready a `Set up payouts` button heads the rail (to `/creators/onboarding`). Read-only; no names (section 1, Privacy). Not a creator: the onboarding page's not-a-creator line.

## 4. API (all under the envelope; openapi.yaml gains every route)

| Route | Who | What |
|---|---|---|
| `GET /v1/admin/creators` | owner | creators with this month's figures and Stripe status |
| `POST /v1/admin/creators` | owner | add: makes the go link and the promotion code |
| `PATCH /v1/admin/creators/{id}` | owner | name, channel, cap, start date |
| `POST /v1/admin/creators/{id}/invite` | owner | a fresh one-time invite URL (token hashed at rest, 14 days) |
| `POST /v1/admin/creators/{id}/stripe-status` | owner | re-read the Express account |
| `GET /v1/admin/creators/statements?month=YYYY-MM` | owner | every creator's statement for the month |
| `POST /v1/admin/creators/{id}/statements/{month}/approve` | owner | settle and transfer (idempotent) |
| `POST /v1/creators/invite` | signed in | accept an invite `{token}` |
| `GET /v1/creators/me` | signed in creator | the stats page |
| `POST /v1/creators/me/onboarding` | signed in creator | make the Express account if needed, answer an account link URL |
| `POST /v1/me/arrival` | signed in | gains an optional `at` (the click time, ms) |

The webhook stays `POST /v1/billing/webhook`: `checkout.session.completed` (a creator code -> attribution, then the first invoice's share), `invoice.paid` (shares), `charge.refunded` (refund lines) and `account.updated` from a second, Connect endpoint on the same URL whose secret is `STRIPE_CONNECT_WEBHOOK_SECRET` (the handler accepts either secret). Onboarding status is also re-read on `?done=1` and by the owner's Refresh, so a missed event never strands a creator. Stripe calls run only with the sandbox key in this lane; go-live steps go in `api/README.md`.

Migration `0092_creators`: `creators`, `creator_invites`, `creator_players`, `creator_revenue`, `creator_refunds`, `creator_payouts`. Only the deployed API applies it.
