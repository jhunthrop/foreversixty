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

## 0B. Coordinator rulings, 2026-11-08 (applied)

- The fixtures agree across pages at a clock of 2026-12-09: November is closed and on the statements page, December is the open month on `/admin/creators` and `/creators/stats`; every count says `all time`, `{Month}` or `paying now`. Grimbrew: 131 returning all time = 64 in November + 67 in December; owed so far $184.20 = December. Lunawick's cap is $50, so both November (over cap $24.40) and December (cap reached) show it.
- `/creators/stats` lists one row per week, not per player.
- No `Settle $0.00`: a month that owes nothing shows `Nothing owed` and no button.
- Payouts sit behind `CREATOR_PAYOUTS`, off by default, with the admin notice (new mock `web-admin-creators-payouts-off`).
- A code gives a free month (owner ruling, section 2); the privacy sentence waits for the owner (section 2).
- The migration takes the next free number on main at landing (0092 is taken by app-beta's `0092_app_testers`).

## 0C. Build reviews, 2026-11-08 (ux-designer and wow-player on the build screenshots; security review)

- ux-designer, SHIP WITH FIXES, all applied:
  - the 390 statements layout (the month field at its text's width, the trail on its own line, a two-column figure grid, a disabled Approve in the secondary tone);
  - the counter `{n} of 500 characters`;
  - 44 px radio rows;
  - the screenshots shot at the fixture clock (2026-12-09).
- wow-player, SHIP WITH FIXES, applied:
  - `Both add up: ...`;
  - the 14-day fill-in line on the public and stats pages;
  - the exact disclosure words;
  - tax details and leaving the program;
  - the disclosure checkbox;
  - `import` for `bring in`;
  - `an estimate until the month closes`;
  - `once we approve it`.

  Left for the owner: whether undisclosed links end a creator's place, a payout day or a minimum, a reapply cool-down, and a stats breakdown.
- Security review: 0 critical. Every high and medium finding is fixed, with tests (review_test.go, creator_code_test.go), except disputes: `charge.dispute.*` is not booked as a negative line yet. That is a known gap, and payouts stay behind the owner's approval.

## 0D. Go-live facts, 2026-11-08 (coordinator)

- **Two keys.** Subscriptions keep `STRIPE_SECRET_KEY`, a restricted `rk_live_…` key in Secret Manager project `foreversixty`, and it is never widened. It also reads players' invoices and refunds for the revenue share.
  - Creator payouts use their own restricted key, `STRIPE_PAYOUTS_KEY`, in the same project. It has exactly four platform-account permissions: Balance read, and write on Accounts (which covers account links), Login Links and Transfers. It has none in connected accounts, so no call ever carries a `Stripe-Account` header. Creator-facing figures come from our own records; the creator's Express dashboard covers the rest.
  - Every Connect and transfer call goes through a gateway built on that key (`creators.NewStripeGateway(secret, payouts)`).
  - In tests, the Connect sandbox key (`STRIPE_SANDBOX_SECRET_KEY`, project `trycommish`) stands for both.
- **Startup check.** With `CREATOR_PAYOUTS` on, the API lists one connected account with the payouts key at startup. A missing key, or one without Connect access:
  - fails loudly: an ERROR log line, `Stripe payouts key missing or lacks Connect permissions. Payouts can't be made until it's set.` on both admin pages, and Approve answering 503 `connect_denied`;
  - never fails silently at payout time.

  Without the key, onboarding answers 503 `payouts_key_missing`.
- **Live Connect is enabled.** An accounts list works, with 0 accounts. The owner is finishing the platform profile. The go-live checklist is in `api/README.md`, "Creator program".

## 0E. Production and the balance, 2026-11-08 (coordinator)

- `CREATOR_PAYOUTS=true` and `STRIPE_PAYOUTS_KEY` are set for production by `api.yml`'s deploy step, merged into the service. Payouts are still one statement at a time, only on the owner's Approve.
- **Balance check before paying.** The Approve dialog reads the platform's available USD balance (`GET /v1/admin/creators/balance`) and shows `Available in Stripe: $X. This payout: $Y.`
  - If X is less than Y, Pay is off and the dialog says `Not enough in your Stripe balance. Add funds in Stripe (Balances → Add to balance) or wait for subscription payments to clear.` (mock `web-admin-creators-confirm-short`).
  - The API checks the balance again before claiming anything.
- **A transfer Stripe refuses with `balance_insufficient`** undoes the claim, so the statement stays unpaid with nothing recorded. The dialog stays open and says `Stripe didn't have enough in your balance for this payout. Nothing was paid; approve it again once the balance covers it.` The idempotency key is the same, so a retry is safe.

## 0F. Program terms, 2026-11-08 (coordinator; on the /creators page, no new mocks)

- Disclosure: the terms require it. The first time a creator's link or code runs without it, we warn them; a second time ends their place in the program (the owner uses Remove; nothing is automatic).
- Statements close on the 1st of the month. The owner approves them during the first week.
- Minimum payout $10 (`MinPayoutCents`). A month under it shows "Under the $10 minimum: it carries over to next month" (status `carry`), cannot be approved (409 `below_minimum`), and its lines stay unsettled so they count toward the next month.
- Reapplying: 30 days after a decline (`ReapplyAfter`). Until then `GET /v1/creators/application` returns `reapply_on` (UTC day), the apply page says "You can apply again from {date}." with no form, and `POST` answers 409 `reapply_too_soon`.
- No per-source stats breakdown for now.

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
2. **Code:** the account redeemed the creator's code at checkout (a free month, section 2) and had no creator yet.

and in both cases: the account was made **on or after the creator's start date**, and it is **not the creator's own account**. One account belongs to one creator only, for good: the first rule that fires wins (a link player who later types another creator's code stays the link creator's).

**A returning player** is a creator's player who
- signed in with Battle.net (`users.bnet_sub` is set; every sign-in is Battle.net since launch),
- has at least one Forever character imported from Battle.net (`characters.source = 'bnet'`),
- and was signed in on a **second UTC calendar day** within **14 days** of the account being made (any signed-in request on a later UTC day than the sign-up day, before `created_at + 14 days`; the API stamps `creator_players.second_day_at` once from the sign-in hook, so nothing per day is stored).

The creator earns **$2.00 once** per returning player, in the month the player qualified (the later of the second day and the statement's check that the character exists).

**Revenue share:** 50% of the **net** of every paid invoice of a subscription billed to the player (Premium, and a Guild or Raid guild plan when the paying officer is the player), for invoices paid within **12 months of the player's first paid invoice**. Net is the payment's Stripe balance transaction `net` (the amount after Stripe's fees), less any tax on the invoice that the fees did not already take, less refunds. A refund lands as a negative line in the month it happens, at 50% of the refunded amount (Stripe keeps its fee on a refund, so the refund line is the refunded amount itself, tax share removed). Invoices of $0 (a trial) are skipped.

**Monthly cap:** default **$500.00** per creator per month, set per creator by the owner. A month's total is `min(cap, max(0, bounties + shares - refunds))`. Anything over the cap is not carried to the next month (the statement shows it as `Over cap`). A month whose refunds outweigh its earnings owes $0, has no action, and its lines stay unsettled, so the refund rolls into the next month's statement.

**Statements and approval:** a month is a UTC calendar month. A statement for a creator and a month lists every unsettled line dated before the month's end. Only a closed month can be approved. Approving is a person pressing Approve on `/admin/creators/payouts`; nothing ever pays automatically. Approving settles the lines (they never appear again), records the payout, and creates one Stripe transfer to the creator's Express account with the idempotency key and transfer group `creator-<id>-<yyyy-mm>`; a second press, a retry or a second tab returns the same payout. A $0 statement has no action (`Nothing owed`). Payouts sit behind a server switch, `CREATOR_PAYOUTS`, off by default: while it is off, statements read normally, Approve is disabled, both admin pages say `Payouts are switched off on the server. Statements can be read but not paid.`, and the API answers 503 `payouts_off`. It stays off until the owner sets up live Connect (transfers under Managed Payments are unverified). A creator whose Stripe account cannot receive transfers yet cannot be approved (no action; the Row says why and carries a `Stripe not ready` Tag; the API answers 409). Approving a month also settles any earlier month that was never approved (its lines are dated earlier), under this month's cap.

**Alts and self-referral:** the creator's own site account (bound by the invite, section 3) never counts, by link or by code, and binding it removes it if it was already counted. One account counts once (the attribution row is keyed by account). One Battle.net account is one site account (`users.bnet_sub` is unique), so a player's alts are characters on one account and count once. A character that already qualified a paid player cannot qualify a second account.

**Privacy:** the creator sees counts only: totals, and one row per week of sign-ups (a Monday, UTC: `Week of 23 Nov: 54 signed up · 36 returning · 10 paying`; the week still running shows only `18 signed up so far`, so "I subbed today" on stream cannot be matched to a change), never a name, a BattleTag, a character or an email. Nothing about a player is in a URL or a log line: logs carry ids only, as the rest of the API does. The owner's statement lists players by their account number.

## 2. Creator codes: a free month, as a Stripe trial

Owner ruling 2026-11-08: a creator's code (`GRIMBREW`, the one the creator says on stream) gives **one month free** on whichever plan the viewer picks: Premium, or a guild plan (Guild or Raid guild), where it turns the guild's standard trial (`catalog.GuildTrialDays`, 7) into 30 days. It is the site's own code, not a Stripe promotion code: the buyer types it at checkout, `POST /v1/billing/checkout` takes `creator_code`, the API checks it is a live creator's code and puts `trial_period_days: 30` on the Checkout Session (`subscription_data`, the same mechanism the guild trial already uses under Managed Payments), with the code in the subscription metadata. A card is still required. One trial per account for Premium (`premium_trials`, written when a trialing subscription lands) and one per guild (`guild_trials`): a code on an account or guild that already had its trial is refused (409 `trial_used`) and the checkout is not made; an unknown code is a 400. On `checkout.session.completed` the program attributes the buyer by the metadata code (the officer, for a guild plan) and records the session's first invoice. Stripe's own Checkout page shows the trial. The revenue share starts at the first paid invoice, after the trial ($0 trial invoices are skipped). Verified in the Connect sandbox on 2026-11-08: a subscription's invoice, payment, charge and balance transaction read for the net (invoice $4.50, fee $0.43, net $4.07; a $2.00 partial refund books -$2.00 with no fee back). Not verified: a live charge under Managed Payments (the sandbox has no tax head office); the first live statement should be checked against the Dashboard before its first approval. The code field on the checkout chooser, with the trial's real dates, is built in the land/premium-trial landing that follows this one (design/specs/2026-11-08-premium-trial.md); until it lands, the API accepts the code but no page sends one. **Draft privacy sentence (pending the owner, privacy.md unchanged):** "If you arrive through a creator's link or use their code, we tell that creator how many players they brought and how many keep playing or subscribe, never who you are, and we pay them for it."

## 3. Surfaces

All four pages are `noindex`, outside the sitemap (`/admin/` is already filtered; `/creators/` is added), and use the site's flat system. The two admin pages are the owner's alone: anyone else gets the same one line `/admin/links` shows (signed out: the sign-in prompt; signed in: `This page is for the site's owner.`); the API answers 401/403 whatever the page shows.

### 3a. `/admin/creators` (owner)

PageHeader `Creators`, lead `Who brought players, what they earned this month, and whether Stripe can pay them.` A split Page:
- **Main:** Section `Your creators` (meta `{n} creators`) holding the page's one Plate: one Row per creator, title the creator's name, sub `/go/{slug} · code {CODE} · since {date}`, foot: five Figures, each saying what it covers: Clicks, Signed up, Returning, Premium net (each `all time`), then `Owed so far` in gold with `{Month} · of {cap} cap` or `{Month} · {cap} cap reached` (the open month); trail: the Stripe status Tag (`Payouts ready` success, `Onboarding` muted, `Needs info` warn, `Not started` muted), then an `Edit` small button. Below the Plate one line link `Monthly statements ›` to the payouts page.
- **Rail:** Section `Add a creator` (or `Edit {name}` when Edit was pressed): Fields Name, Channel (a Twitch or YouTube URL), Link slug (becomes `/go/slug`, new), Promo code (uppercase letters and digits, 4 to 20), Starts on (a date; players who signed up before it never count), Monthly cap (dollars, default 500). Primary `Add creator` / `Save`. In edit mode the slug and code are read-only text (they are live in Stripe and in posts) and two more actions: `Copy onboarding invite` (makes a fresh one-time invite link, 14 days, copies it; the old one stops working) and `Refresh Stripe status`.
- States: loading (Skeleton), signed out, not owner, load error (LoadError with Retry), empty (`No creators yet. Add the first one.`), saving, field errors from the API under each Field, invite copied (`Invite copied. It works once, for 14 days.`), Stripe unavailable (`Stripe isn't configured, so codes and payouts are off.` one line above the Plate).

### 3b. `/admin/creators/payouts` (owner)

PageHeader `Statements`, lead `Each creator's month. Nothing is paid until you approve it.`, action link `‹ Creators`. A split Page:
- **Main:** Section `{n} creators` (meta `{total} to pay`), its action the month `Select` (the current month `December 2026 (open)` and the twelve closed months before it; the page opens on the last closed month). With the payout switch off, the line `Payouts are switched off on the server. Statements can be read but not paid.` (wipe) sits above the Plate and Approve is disabled. The page's one Plate: a Row per creator, title the name, sub the status line (`Ready to approve`, `Month still open`, `Paid {date} · transfer tr_…`, `Nothing owed`, `Stripe can't pay yet: onboarding not finished`); trail a `Stripe not ready` Tag when blocked, the action, and `Lines`; foot five Figures (Total in gold first, Returning with `{n} × $2.00`, Share, Refunds, Over cap). The only action is `Approve {amount}` primary (closed, unpaid, a ready account, total over $0); a month that owes nothing has none. The selected Row (the first by default) is marked current.
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

PageHeader `{creator name}`, meta `foreversixty.gg/go/{slug} · code {CODE}`. A split Page. Main: the one Plate holding four Figures, each saying what it covers (`Earned so far` in gold with `{Month}, estimated, {cap} cap max`; Players `all time`; Returning `all time`; Paying `paying now`), then Section `Your players` (meta `by week · no names, by design`) as RowsSlim, one row per week, newest first, at most 26: `Week of 23 Nov: 54 signed up · 36 returning · 10 paying`, the running week `Week of 7 Dec: 18 signed up so far`. Rail: Section `Payouts`, one RowsSlim line per approved month, `{Month}` `{amount}`; empty `No payouts yet. Each month is paid after it closes and is approved.` While Stripe is not ready a `Set up payouts` button heads the rail (to `/creators/onboarding`). Once it is ready, a `Payouts and tax forms` button closes the rail: it opens the creator's Stripe Express dashboard (payouts, bank details, the 1099) through a login link the API makes on each press (`POST /v1/creators/me/dashboard`, Stripe's `POST /v1/accounts/{id}/login_links`), never stored and never written into the page. A signed-in creator also gets a `Creator stats` item in the profile and account menu (`/v1/me` carries `creator: true`). Read-only; no names (section 1, Privacy). Not a creator: the onboarding page's not-a-creator line.

### 3e. `/creators` (public, indexable, in the sitemap; the footer's About column links it as `Creators`)

PageHeader `Creators`, lead `If you stream or make videos about World of Warcraft: Forever, you can earn from the players you send us.`, action `Apply` primary. A document (`Page reading`, no contents rail): ReadingSections `What you get` (the $2 rule in plain words, the 50% share for 12 months, Premium and the Raid guild plan while in the program), `How players count` (link within 30 days or the code; one account, one creator; older accounts and your own never count), `Payouts` (monthly, through Stripe, the cap, the stats page by week), `Saying it is sponsored` (FTC and Twitch disclosure), and `Apply` with the button again. Every word is `programCopy`. Apply signed in goes to `/creators/apply`; signed out, to Battle.net sign-in with `next=/creators/apply`. Premium and the Raid guild plan are granted by hand with the `grant` CLI after approval (grants are never an HTTP route, RULING 4).

### 3f. `/creators/apply` (signed in; noindex)

`Page layout="form"`, PageHeader `Apply to the creator program`. Section `Your channel`: Fields Channel (a Twitch or YouTube https address; the API checks the host is twitch.tv or youtube.com), Main platform (two radios), Viewers or subscribers (optional whole number), `What do you make about Forever?` (required, 1 to 500 characters, a counter `{n} of 500 characters`), Contact email (optional, only the owner sees it; never in a URL or a log), a required checkbox `I'll say it's an affiliate link or sponsored wherever I share my link or code.` (unticked: `Tick this to apply: the FTC and Twitch require it.`), the consent line `We review every application; we'll reply on your Forever Sixty account.`, primary `Send application`. One open application per account (409 otherwise), 5 sends an hour per address. States after sending, in a Section `Your application` with a Tag: Pending (`Sent {date}. We'll reply here.`), Approved (`You're in. Set up payouts so we can pay you.` and `Set up payouts` to `/creators/onboarding`), Not this time (`Not this time. Thank you for applying.`; the form comes back so they can apply again later). Already a creator: `You're already in the creator program.` and `Your stats ›`. Signed out: the sign-in prompt.

### 3g. Requests on `/admin/creators`

A Section `Requests` (meta `{n} waiting`) above `Your creators`: open Rows (the page keeps one Plate), title the channel without its scheme, sub `{Platform} · about {n} | size not given · {date}`, foot the applicant's own text, trail `Approve` primary and `Decline`. Empty: the Section is not shown. Approve puts the rail in Approve mode, `Approve {channel}`, the add form prefilled (name and slug from the channel path, the code its capitals, starts today, cap 500), with `Their account is already bound; no invite is needed.`; `Add creator` creates the creator bound to the applicant's account and marks the request approved. Decline marks it declined at once (no reason). The contact email shows under the text for the owner only.

**Owner alert.** There is no owner-only alert path today (the events catalogue is per-player preferences shown in every account's settings), so the owner's account menu shows the count as a badge, `{n} creator requests`, linking `/admin/creators`; `/v1/me` carries `creator_requests` for the admin only.

## 4. API (all under the envelope; openapi.yaml gains every route)

| Route | Who | What |
|---|---|---|
| `GET /v1/admin/creators` | owner | creators with this month's figures and Stripe status |
| `POST /v1/admin/creators` | owner | add: makes the go link and reserves the code |
| `PATCH /v1/admin/creators/{id}` | owner | name, channel, cap, start date |
| `POST /v1/admin/creators/{id}/invite` | owner | a fresh one-time invite URL (token hashed at rest, 14 days) |
| `POST /v1/admin/creators/{id}/stripe-status` | owner | re-read the Express account |
| `GET /v1/admin/creators/statements?month=YYYY-MM` | owner | every creator's statement for the month |
| `POST /v1/admin/creators/{id}/statements/{month}/approve` | owner | settle and transfer (idempotent) |
| `POST /v1/creators/invite` | signed in | accept an invite `{token}` |
| `GET /v1/creators/me` | signed in creator | the stats page |
| `POST /v1/creators/me/onboarding` | signed in creator | make the Express account if needed, answer an account link URL |
| `POST /v1/creators/me/dashboard` | signed in creator, Stripe ready | a fresh Express dashboard login link `{url}` |
| `GET /v1/creators/application` | signed in | the account's latest application and whether it is a creator |
| `POST /v1/creators/application` | signed in, rate limited | send one (409 while one is open or for a creator) |
| `GET /v1/admin/creators/requests` | owner | open requests, with suggestions for the Approve form |
| `POST /v1/admin/creators/requests/{id}/approve` | owner | create the creator from the request, bound to the applicant |
| `POST /v1/admin/creators/requests/{id}/decline` | owner | decline |
| `POST /v1/me/arrival` | signed in | gains an optional `at` (the click time, ms) |

The webhook stays `POST /v1/billing/webhook`: `checkout.session.completed` (a creator code -> attribution, then the first invoice's share), `invoice.paid` (shares), `charge.refunded` (refund lines) and `account.updated` from a second, Connect endpoint on the same URL whose secret is `STRIPE_CONNECT_WEBHOOK_SECRET` (the handler accepts either secret). Onboarding status is also re-read on `?done=1` and by the owner's Refresh, so a missed event never strands a creator. Stripe calls run only with the sandbox key in this lane; go-live steps go in `api/README.md`.

Migration `0100_creators` (0094 on the branch; renumbered at landing, 2026-11-11): `creators`, `creator_invites`, `creator_players`, `creator_revenue`, `creator_refunds`, `creator_payouts`, `creator_applications`. Only the deployed API applies it.
