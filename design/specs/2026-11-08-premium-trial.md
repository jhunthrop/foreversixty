# Free trials, told plainly: experience and system spec

Author: build lane `creator-program` (landing `land/premium-trial`, after `land/creator-program`), 2026-11-08, from the owner's approved change. Mock first (tenet 15): the surfaces below match `design/mocks/ws_premium_trial.py`. The owner asked for full transparency: every place a trial is offered, started, running or about to charge says the real end date and the real price, and cancelling is one click and one confirm.
System: `design/specs/2026-11-05-web-layout.md`; the shipped premium, checkout and account surfaces (`design/specs/2026-11-08-premium-account.md`, mocks `ws_premium.py`), whose blocks the trial mocks import. The creator code's API side is `design/specs/2026-11-08-creator-program.md` section 2.

## 0. Coordinator rulings, 2026-11-08

1. Reminders and receipts go to the Stripe customer's email, read from Stripe at send time, never stored and never logged, sent through Resend; the in-app notification goes too.
2. A Stripe customer with no email gets the in-app notification alone; the miss is logged with the Stripe customer id only.
3. Stripe's own "trial ending" reminder (Billing settings, customer emails) is a backstop the owner can turn on; ours stays the primary reminder.
4. The guild trial reminder ships in this landing, so the existing guild trial copy (`plans-page-copy.ts` `trialEmail`, "We email you 3 days before the trial ends.") becomes true and is appended to the trial line, as its comment asks.
5. Order: `land/creator-program` first, then `land/premium-trial`.

## 0B. Coordinator copy rulings, 2026-11-08 (apply in the build, mocks follow)

1. The chooser names the exact first charge for the chosen interval and its date, in one line under the existing sentence: `On 8 Dec 2026 you'll be charged $40 for a year of Premium` (yearly) or `On 8 Dec 2026 you'll be charged $5, then $5 each month` (monthly).
2. While the founder price is open, a trial start says: `Start before 6 Jan 2027 and the founder price stays for as long as you stay subscribed, even though your first charge comes after the trial.` The build must prove it: the subscription's price is the founder price fixed at checkout (catalog.CheckoutLookupKey at the session's creation), and a test runs a trial that ends after the founder window closes and checks the first invoice is the founder amount.
3. The reminder email names the plan and interval (`Premium, yearly: $40`); a yearly trial says `you'll be charged $40 for the next 12 months`. The email ends with one line saying why it was sent (`You're getting this because you started a free Premium trial on Forever Sixty.`) and the company name, Commish LLC. Never the owner's name.

## 0C. The email split, decided against the owner's live Stripe settings (coordinator, 2026-11-08)

- **Stripe sends:**
  - the 7-day trial-ending reminder;
  - upcoming renewals;
  - expiring cards;
  - failed payments;
  - refunds;
  - the manage-subscription link in its emails.
- **We send:**
  - our 3-day trial reminder, plus the in-app notification;
  - the guild trial reminder;
  - the first-charge receipt.

  Stripe's "Successful payments" receipt stays off.
- **No change to renewal email recipients.** Our renewal emails do not move to the Stripe customer's email, so the earlier side effect is dropped. Stripe owns renewals. Our renewal email (billing-daily's yearly renewal notice) sits behind a server switch, `OUR_RENEWAL_EMAILS`, which is off by default. Off, billing-daily sends no renewal notice. The spec and `api/README.md` both say so.
- **Our reminder never contradicts Stripe's.** It gives the same dates and the same price, and says once: `You may also get a reminder from Stripe, our payment processor.`

## 0D. Server side as built, 2026-11-09 (UI waits for the design system)

The coordinator split the landing: the server side is built now on `premium-trial` (off `creator-program`); the chooser, the account trial state and the confirmation are built on the new design system (`docs/design-system.md`) once `land/design-system` is on main, against the trial mocks the design lane redraws (batch 4). The old look is not built.

- **Spent trial:** an account's Premium trial is spent when `premium_trials` has it **or** it already paid for Premium (a Stripe Premium row, any status): its first Premium was that subscription. A seat-holder's comeback checkout never trials. A code on a spent trial is 409 `trial_used`.
- **Routes:** `POST /v1/billing/checkout/terms` (session; its own 120/hour budget); `POST /v1/billing/trial/cancel` and `/trial/keep` `{guild_id?}` (browser session; the caller's trialing Premium, or a guild's when the caller is its payer; 409 `not_trialing`, 403 for a non-payer); `POST /v1/billing/cancel-link/preview` and `/cancel-link` `{token}` (no sign-in; 410 `link_expired` with the 4.3 sentence). Each change re-reads the subscription into the entitlement at once.
- **Reminder:** only trials recorded in `premium_trials` or `guild_trials` (a billing transfer's trial never gets one). The notice is claimed, the in-app `billing.trial_ending` goes (dedupe key per trial), then the email with a fresh token; a failed send releases the claim and drops the token. A free trial is never also sent our renewal notice.
- **Receipt:** `invoice.paid` with `amount_paid > 0` and `billing_reason = subscription_cycle` on a recorded free trial, claimed once (`trial_receipt`); a failed send fails the event so Stripe redelivers.
- **Email copy contract:** `api/internal/billingmail/testdata/trial-emails.json` (rewrite with `go test ./internal/billingmail -update`); `trial-copy.ts` is held to it when the UI lands.
- **Migration** `0095_trial_reminders`: `billing_cancel_tokens` and the two notice kinds.

## 1. Files

| What | Path |
|---|---|
| Copy and the date and price helpers (every trial word; prices from `plans.ts` through `plan-costs.ts`, dates from `dates.ts`) | `web/src/lib/billing/trial-copy.ts` |
| Facts reader (esbuild bundle of `trial-copy.ts` at the premium mocks' clock, 2026-11-08T10:00Z, founder window open) | `design/mocks/premium_trial_facts.mjs` -> `design/mocks/data/premium-trial-facts.json` |
| Generator (ws_prims and ws_premium's own blocks) | `design/mocks/ws_premium_trial.py` (13 mocks, `--list`) |
| Renders (390 and 1440, full page) | `design/reviews/2026-11-08-premium-trial/mock/` |
| Regenerate | `node design/mocks/premium_trial_facts.mjs && python3 design/mocks/ws_premium_trial.py`, then `node design/mocks/ws_shoot.mjs <mock> 390,1440` |

## 2. The rules

- **Premium:** every first Premium checkout trials **7 days** (`PREMIUM_TRIAL_DAYS`); with a creator's code, **30 days** (`CODE_TRIAL_DAYS`). One Premium trial per account, ever (`premium_trials`, written when a trialing subscription lands, from the creator-program landing). A guild trial does not count against it. An account that had its trial checks out at the full price, no trial; a code on it is refused (409 `trial_used`) with the chooser's sentence, and the buyer can still subscribe.
- **Guild plans:** the existing 7-day guild trial (`GUILD_TRIAL_DAYS`), 30 days with a code, one per guild (`guild_trials`); a code on a guild that had its trial is refused the same way.
- **A card is always required.** Stripe Checkout collects it (`payment_method_collection` stays `always`), and the session's `trial_period_days` makes Stripe's own page show the trial and its end date. Nothing is charged until the end.
- **The server decides.** The checkout decides the trial (`api/internal/billing/checkout.go`, `creator_code.go`); a client can't ask for days. `POST /v1/billing/checkout/terms` (new, read-only) answers the chooser: `{ trial_days, ends_at, price_cents, interval, founder, code: "applied" | "unknown" | "trial_used" | null }` for a plan, interval and optional code, so the chooser shows the dates the checkout will really use. The checkout itself re-checks everything.
- **Dates:** a trial ends `days * 24h` after the session starts (Stripe's own `trial_end`). Every surface shows that instant's UTC day in the house style (`15 Nov 2026`), and "Cancel any time before {date}" means before that day begins, so it is true in every time zone: the cutoff is at most 24 hours before the charge.
- **Cancelling** a trial sets `cancel_at_period_end` (`Gateway.CancelSubscriptionAtPeriodEnd`, which exists): the plan stays on to the trial end and Stripe never charges. Undo ("Keep Premium after the trial") clears it (`ResumeSubscription`, exists). No retention step, no survey, no offer.

## 3. Surfaces

### 3.1 The trial line (one sentence, everywhere it is offered)

`trialLine` in `trial-copy.ts`: `Free for 7 days, then $5 a month (or $40 a year). Cancel any time before 15 Nov 2026 and you pay nothing.` While the founder window is open it adds `Founder price for anyone who subscribes by 6 Jan 2027.`; once it closes the prices are the list prices (`$6 a month (or $50 a year)`). A guild plan says `... (or $12 a month) for the whole guild. ...`. The date is computed from now on the plans page and the chooser (a trial started now would end then), and from the real `trial_end` everywhere after the start.

### 3.2 `/premium`, the Premium column (`web-trial-card`)

The column keeps its head (price, `Or $40 a year.`) and gains a `Free trial` Tag beside the name, the trial line under the price, and the primary button `Start 7 free days` (to the chooser). For an account that already had its trial the line and Tag are gone and the button is the shipped `Subscribe · {price}`. The guild columns keep the shipped trial line, with `trialEmail` now appended (ruling 4) and the trial line's date added.

### 3.3 The checkout chooser (`web-trial-chooser`, `-code`, `-code-unknown`, `-used`)

A new first step on `/premium/checkout` before Stripe (today the page redirects at once). PageHeader `Start your free trial`, the shipped band (price, founder Tag, `Change plan ›`). One Plate: `Billing` (Yearly, Monthly segmented), the trial line for that interval, a `Creator code` Field (optional; hint `Optional. A creator's code makes the trial a free month.`) with `Apply`, the code's status line, `Stripe asks for a card now and shows the same dates before you confirm.`, and the primary action:
- no code: `Start 7 free days`;
- a valid code: `GRIMBREW applied: your first month is free. Remove`, the line changes to 30 days and `8 Dec 2026`, the action `Start 30 free days`;
- an unknown code: `That isn't a creator's code. Check the spelling, or start without one.` (wipe), the 7-day line and action stay;
- the trial already used: no trial line, `This account already had its free trial, so a code can't give another. You can still subscribe.` and `Subscribe · $40 a year`.
A creator's link can prefill the code (`/premium/checkout?code=GRIMBREW` puts it in the Field, still checked by Apply). The guild checkout's `Which guild?` step gets the same Field and line, with the guild's sentence for a used trial.

### 3.4 The confirmation (`web-trial-confirmed`)

Back from Stripe, the shipped steps with row 2 `Card saved. First charge {date}.` (ruling K1, already shipped for guilds) and row 3 `Trial ends 15 Nov 2026, then $5 a month.`; under the Plate the trial line again, `Nothing else to do. We remind you 3 days before the first charge.`, and `Cancel the trial` (opens the confirm in 3.5). The guild confirmation reads `{guild}'s free trial is on` with the guild's line.

### 3.5 The account page and guild settings (`web-trial-account`, `-confirm`, `-canceled`, `-guild`)

The account's plan Section, while trialing, always shows `Premium` with the `Free trial` Tag, `Trial ends 15 Nov 2026, then $5 a month.`, and `Cancel trial`. The confirm is the shared Dialog: `Cancel your free trial?`, `Premium stays on until 15 Nov 2026 and you won't be charged.`, `Keep my trial` / `Cancel trial`. After it: `Trial ends 15 Nov 2026. You won't be charged.` and `Keep Premium after the trial`. Guild settings' plan row (the billing officer only) reads `Guild · {guild}`, `Trial ends 15 Nov 2026, then $120 a year.`, `Cancel trial` with the same Dialog (`{guild} keeps the plan until {date} and you won't be charged.`); the account page's guild row line says the same.

### 3.6 The emails and the notification (`web-trial-email-reminder`, `-email-guild`, `-email-receipt`)

Plain text, as every billing email today (`api/internal/billingmail`, sent by `mail.Resend` from billing-daily):
- **Reminder** (Premium): subject `Your free Premium ends on 15 Nov 2026`; `Your free Premium ends on 15 Nov 2026. On that day you'll be charged $5 a month. Cancel in one click: {link}` then `To keep Premium, do nothing.`
- **Reminder** (guild): `{guild}'s free Guild ends on {date}. On that day you'll be charged {price}. Cancel in one click: {link}`.
- **In-app notification** (a new events kind `billing.trial_ending` to the payer, in the inbox and as a push once the apps carry it): title `Free Premium ends 15 Nov`, body `On that day you'll be charged $5 a month. Tap to cancel.`, opening the account page's plan.
- **Receipt** after the first charge (`invoice.paid` with an amount after a trial): subject `Receipt: Premium, $5 a month`; `Your free trial ended and we charged $5 a month on 15 Nov 2026. Your receipt: {Stripe hosted invoice URL}` then `The next charge is on 15 Dec 2026. Change or cancel any time: {account link}`.
`billingmail` holds the API's copy of these sentences; a Go test compares them with `trial-copy.ts`'s output for fixed inputs (the contract fixtures pattern), so the site and the email never drift.

## 4. The system

### 4.1 Who the email goes to

Battle.net accounts have no `users.email` (the shipped renewal email skips them: `notices.go` logs `no email`). Ruling 1: the reminder and the receipt go to the **Stripe customer's email**, which Checkout collects, read with one Stripe call (`Gateway.CustomerEmail`) at send time, used for that one send, never written to our database and never put in a log; the in-app notification always goes too. With no email on the customer, only the notification goes and the job logs `billing notice skipped` with the entitlement id and the Stripe customer id (ruling 2). The yearly renewal and size emails switch to the same lookup in this landing, so they reach Battle.net accounts for the first time.

### 4.2 When the reminder goes

Billing-daily (daily at the scheduled hour) sends `trial_ending` for every trialing Stripe row whose `trial_end` is within the lead: **3 days** for every trial of 6 days or more (the 7-day, the guild's 7-day and the 30-day), **2 days** for a shorter one (none today; `reminderLeadDays`). 3 days for a 7-day trial means the reminder lands on day 4, which leaves the buyer three full days to decide and keeps the promise the guild copy already makes. Each reminder goes once per trial: claimed in `billing_notices` (kind `trial_ending`, entitlement, period end) as the renewal email is, released on a send failure for the next run. A trial cancelled before the run gets none. Because the job runs once a day, a reminder is sent between 3 and 4 days before the charge; the email says the date, never "in 3 days".

### 4.3 The one-click cancel link, made safe

`https://foreversixty.gg/premium/cancel?t=<token>`. The token is 32 random bytes, stored only as its SHA-256 in `billing_cancel_tokens` (entitlement id, purpose `cancel_trial`, expires at the trial end, used_at), one per reminder. Opening the link never cancels anything (mail scanners and link previews open links): the page asks the API for what the token names (`POST /v1/billing/cancel-link/preview {token}` -> plan, guild name, end date, price) and shows one confirm, `Cancel your free trial?` with the 3.5 sentence and one button; that button's `POST /v1/billing/cancel-link {token}` sets `cancel_at_period_end` and marks the token used. No sign-in is needed (the email is the proof, as Stripe's own portal links), it cannot do anything but cancel at period end the one subscription it names, it expires at the trial end, it works once, and the token travels in POST bodies to the API, never in an API URL. A used or expired link says `This link has expired or was already used. You can cancel from your account page.`

### 4.4 Stripe and Managed Payments

- Trials under Managed Payments already ship (the guild trial), so `trial_period_days` on Premium is the same call.
- Stripe's Checkout page shows the trial ("7 days free, then ...") itself.
- Stripe's own customer emails (a trial-ending reminder, receipts) may also go out under Managed Payments, where Stripe is the merchant of record and owns some customer communication. The owner checks the Dashboard's customer email settings in live before launch: Stripe's trial reminder stays off unless the owner wants the backstop (ruling 3); if Stripe already sends a receipt as merchant of record, ours still goes because it carries the next date and our cancel link, and the owner may turn Stripe's off. Not verifiable from the sandbox (it has no tax head office, so no Managed Payments).
- The receipt link is Stripe's `hosted_invoice_url` for the paid invoice, read at send time.

### 4.5 API changes (land/premium-trial)

`checkout.go`: Premium trial of 7 days when `premium_trials` has none for the account (the code path from creator-program stays). `POST /v1/billing/checkout/terms` (3.3). `notices.go`: kinds `trial_ending` (Premium and guild) and `trial_receipt`; recipient per 4.1. `webhook.go`: `invoice.paid` with an amount on a subscription that trialed sends the receipt once (claimed like a notice). `cancel_link.go`: the token table, preview and cancel routes. `events`: kind `billing.trial_ending`. `GET /v1/me`: the Premium billing block gains `trial_end` (it has `current_period_end`; a trialing row's period end is the trial end) so the account page can say "Trial ends". Migration: `billing_cancel_tokens`.

## 5. Tests the build carries

- One trial per account: a second Premium checkout after a trial has no trial and a code on it is 409; a guild trial does not block a Premium trial; one per guild for guild plans (creator-program's tests plus the 7-day default).
- The reminder: fires once, on the run 3 days before (and not the run before), not again the next day, not for a cancelled trial, 2 days for a 5-day trial; the guild reminder the same.
- Cancelling during the trial: `cancel_at_period_end` set, the entitlement stays active to the trial end, the subscription ends with no invoice charged (FakeGateway records no charge); the cancel link cancels only after the POST, only once, only its own subscription, and not after it expires; a GET of the page changes nothing.
- The copy shows real dates: `trial-copy.ts` unit tests at fixed clocks (7 and 30 days, both intervals, founder open and closed, a start at 23:30 UTC), the facts reader's output matching the mocks, and the Go email copy matching `trial-copy.ts`.
- No PII in logs: the job's and the webhook's logs carry ids only; a test captures them and checks the Stripe customer's email and the token never appear.
- The web: unit tests for the chooser's states and the account page's trial row; parity specs for the chooser and the account page in trial.
