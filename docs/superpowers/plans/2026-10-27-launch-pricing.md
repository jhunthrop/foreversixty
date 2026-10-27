# Launch Pricing Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Checkout sells the launch prices: Premium $6/$50, Guild $15/$150, Raid guild $25/$250, with founder prices until 6 January 2027. A guild plan stops granting every member Premium; it gives a fixed number of Premium seats instead. `/premium`, guild settings `#billing`, the account `Plan` panel and five emails follow the experience spec.

**Architecture:**
- A leaf Go package, `billing/catalog`, holds every lookup key and amount, the founder end instant, and the tier and seat rules. Checkout, the portal, the webhook, the emails and `/v1/me` all read prices from it.
- Migration 0054 adds four price facts to `entitlements` (`tier`, `price_lookup_key`, `unit_amount_cents`, `trial_end`) and five tables: seats, trials, transfers, comebacks and notices.
- `entitlements` resolves Premium from a personal row or a held seat, never from guild membership. It gains a per-guild feature, `raid_guild_tools`, which gates the loot council and readiness actions.
- A new `seats` package owns giving, moving, offering and releasing seats.
- `billing` gains the guild billing read, the Guild→Raid guild upgrade, the transfer offer, and a `billing-daily` job that sends the renewal and size notices and sweeps seats.
- The web rebuilds `/premium` as a static four-card page; a head script sets its founder state. `GuildBillingSection` is rebuilt on a new `GET /v1/guilds/{id}/billing` read. The account `Plan` panel is extended from new `/v1/me` fields.

**Tech Stack:** Go (pgx v5, golang-migrate, stripe-go v82), Astro + Svelte 5, Tailwind, vitest, Playwright + axe.

**Spec:** The plan implements two specs:
- `design/specs/2026-10-27-launch-pricing.md`: experience, copy verbatim, states, rulings R-P1..R-P8, and the §9 acceptance cases.
- `docs/superpowers/specs/2026-10-27-pricing-design.md`: prices in §3, rules in §4, builder decisions in §8.

Chip-in is deferred and is not in this plan.

## Global Constraints

- **Prices.** These come from product §3 and live only in `api/internal/billing/catalog/catalog.go` (Go) and `web/src/lib/billing/plans.ts` (web). Copy never types a price.

  | Plan | List (month / year) | Founder (month / year) |
  |---|---|---|
  | Premium | $6 / $50 | $5 / $40 |
  | Guild | $15 / $150 | $12 / $120 |
  | Raid guild | $25 / $250 | $20 / $200 |
- **Founder window.** It has one end instant, `2027-01-07T08:00:00Z` (6 January 2027 23:59:59 America/Los_Angeles), exclusive. The controller's ruling on R-P8 is that the window opens when this lane ships. There is no opening instant, no "before" state and no "from 4 November" line.
  - The API measures the window when it creates the Checkout session.
  - The page's head script uses the same instant.
- **Founder lock.** The founder price holds while the subscription stays continuous. A transfer, a monthly↔yearly switch in the portal, and the Guild→Raid guild upgrade all keep it. A lapse followed by a resubscribe pays list price (product §8.2).
- **Trial.** 7 days via Stripe `trial_period_days`, card up front, guild plans only, one per guild. The `guild_trials` table enforces the one-per-guild rule (product §8.5).
- **Tiers.**
  - Guild covers up to 25 verified members. This is a soft cap: nothing switches off. It has 1 seat.
  - Raid guild covers any size and has 3 seats.
  - A seat moves at most once every 30 days. A returned seat shows for 14 days. A comeback link lasts 30 days.
- **Stripe secrets.** Stripe keys never appear in files, logs or responses. Tests use `FakeGateway` only. Log lines carry subscription and price ids, never metadata dumps.
- **Migrations.** Only the deployed API applies migrations. No tool, test helper or CLI run against production migrates.
  - This lane's migration is `0054_launch_pricing`. The discord-bot lane holds 0053.
  - golang-migrate never applies a version below the live one, so 0054 must not deploy before 0053 (Task 22).
- **API tests.** Run with `-p 1`, each agent on a private database:
  - `TEST_DATABASE_URL=postgres://forever:forever@localhost:5434/<db>?sslmode=disable`, where `<db>` is `pricing_t<N>` for task N.
  - Create the database with `psql postgres://forever:forever@localhost:5434/postgres -c 'create database pricing_t<N>'`.
- **Web copy.** All web copy goes through a copy module whose test reads the spec table (`web/src/test-support/spec-table.ts`, `specCell` / `specString`) against `design/specs/2026-10-27-launch-pricing.md`.
- **File limits.** Go files stay at or under 800 lines; Svelte and TS files at or under 400. Two files are already over: `api/internal/auth/handler.go` (835) and `web/src/lib/account/api.ts` (781). Additions go in new files; those two only gain import or one-line hook edits.
- **Web setup.** The new worktree needs `npm ci` in `web/` before any web test.
- **Owner's grant.** The owner's account (prod user 78) holds an open-ended personal Premium grant (`source = 'grant'`, `current_period_end` null). It must keep every Premium feature through every change here.
- **Experience rules** (premium spec §2 and §5):
  - No upsell interrupts, no urgency, no countdowns, no exclamation marks.
  - `/premium` lists only live benefits (R-P5): the words "retention" and "compare more than" never render.
  - No founder or seat mark on any public surface (R-P6).
- **Static page.** `/premium` stays static: zero `_astro` scripts and two `is:inline` scripts only. No new design token. No new shared control: every control is a recipe the spec names.
- **Commits.** The trailer is exactly `Co-Authored-By: Claude Opus 5.5 <noreply@anthropic.com>`. Use conventional commit types.

## Review Focus

1. **The founder boundary.**
   - A Checkout session created at `2027-01-07T07:59:59Z` uses a `*_founder_*` key; one created at `08:00:00Z` uses the list key.
   - The page's head script flips `data-founder` at the same instant.
   - Pinned in Task 2 (`TestCheckoutLookupKey_FounderBoundary`) and Task 16 (e2e with a faked clock on each side of the instant).
2. **Premium after the guild-grant removal.**
   - A verified member of a guild on a whole plan, with no seat, loses Premium and the supporter mark.
   - The owner-style open-ended personal grant still answers `Can` true for every feature, including while the account also holds a seat elsewhere.
   - Pinned in Task 4 (`TestCan_MemberWithoutSeatHasNoPremium`, `TestCan_OpenEndedPersonalGrantSurvives`).
3. **A seat holder leaves the guild.**
   - Their Premium ends at once, because resolution joins verified membership.
   - The next read shows the slot `empty` with `returned {name, at}`.
   - If they gave up their own Premium for the seat, there is exactly one comeback row and one "seat removed" email, even when the sweep runs twice.
   - Pinned in Task 4 and Task 7 (`TestSweep_DepartedHolderReleasedOnce`).
4. **A transfer accepted close to period end.**
   - Stripe refuses a Checkout `trial_end` less than 48 hours away. So when the outgoing period ends sooner than that, the new subscription's `trial_end` is `max(periodEnd, now+49h)`.
   - The outgoing subscription is canceled at period end, and `stripe-reconcile` does not record it as an orphan.
   - Pinned in Task 5 (`TestTransferCheckout_TrialEndFloor`) and Task 6 (`TestReconcile_OutgoingTransferIsNotOrphan`).
5. **A declined card on the upgrade, and a double-submitted give.**
   - `ChangePrice` uses `payment_behavior=error_if_incomplete`. On a declined card the subscription stays on Guild, the entitlement tier is unchanged, and the route answers 409 `upgrade_failed`. The web shows `That did not go through. Nothing was charged. Try again.`
   - A give submitted twice on the same slot leaves one holder; the second request gets 409.
   - Pinned in Task 8 (`TestUpgrade_DeclinedLeavesGuild`) and Task 7 (`TestGive_ConcurrentSameSlot`).

---

## PR A — the API

### Task 1: Migration 0054 and the entitlement row's price facts

**Files:**
- Create: `api/internal/db/migrations/0054_launch_pricing.up.sql` and `.down.sql`, `api/internal/db/launch_pricing_migration_test.go`, `api/internal/entitlements/price_facts_test.go`
- Modify: `api/internal/entitlements/entitlements.go` (tier constants), `api/internal/entitlements/store.go` (`StripeUpsert`, `UpsertStripe`, `Billing`, `PersonalBilling`, `GuildBilling`, `Grant`)

**Interfaces:**
- Produces (schema):
  - `entitlements` gains `tier text`, `price_lookup_key text`, `unit_amount_cents integer` and `trial_end timestamptz`.
    - New constraint `entitlements_tier`: `(plan='premium' and tier is null) or (plan='guild' and tier in ('guild','raid_guild'))`.
    - The migration sets `tier='raid_guild'` on every existing guild row, so grants and the grandfathering scoped grants keep the loot council and readiness.
  - `premium_seats`:
    - Columns: `guild_id bigint` references guilds on delete cascade; `slot smallint check (slot between 1 and 3)`; `state text check in ('empty','held','offered','accepted')`; `holder_user_id bigint` references users on delete set null; `given_by bigint` references users on delete set null; `given_at`, `offered_at`, `starts_at`, `movable_at timestamptz`; `holder_gave_up_lookup_key text`; `returned_name text`; `returned_at timestamptz`.
    - Primary key `(guild_id, slot)`.
    - Check `(state='empty') = (holder_user_id is null)`.
    - Unique index `(guild_id, holder_user_id) where holder_user_id is not null`.
  - `guild_trials`: `guild_id bigint primary key` references guilds on delete cascade, `stripe_subscription_id text not null`, `started_at timestamptz not null default now()`.
  - `plan_transfers`: `guild_id bigint primary key` references guilds on delete cascade, `from_user_id` and `to_user_id bigint not null` referencing users, `from_subscription_id text not null`, `state text check in ('offered','declined','accepted')`, `offered_at timestamptz not null default now()`, `decided_at timestamptz`.
  - `seat_comebacks`: `id bigserial primary key`, `user_id bigint not null` references users on delete cascade, `guild_name text not null`, `ended_at timestamptz not null`, `price_lookup_key text not null`, `expires_at timestamptz not null`, `used_at timestamptz`.
  - `billing_notices`: `kind text check in ('renewal','size')`, `entitlement_id bigint` references entitlements on delete cascade, `period_end timestamptz`, `sent_at timestamptz not null default now()`, primary key `(kind, entitlement_id, period_end)`.
- Produces (Go, `entitlements`):
  - `const TierGuild = "guild"; const TierRaidGuild = "raid_guild"`.
  - `StripeUpsert` gains `Tier string`, `PriceLookupKey string`, `UnitAmountCents int64` and `TrialEnd *time.Time`. `UpsertStripe` writes all four. `Tier` is ignored for premium rows. A guild row with an empty `Tier` fails fast with an error.
  - `Billing` gains `Source string`, `StripeSubscriptionID string`, `PriceLookupKey string` and `UnitAmountCents int64`.
  - `GuildBilling` gains `ID int64`, `Tier string`, `PriceLookupKey string`, `UnitAmountCents int64`, `TrialEnd *time.Time` and `StripeSubscriptionID string`.
  - `Grant` writes `tier = TierRaidGuild` for every guild row, whole-plan or scoped.

- [ ] **Step 1: Write the failing tests.**
  - `TestLaunchPricingMigration_TierBackfill`: seed a guild grant row and a premium row at 0053, then migrate to 0054. The guild row has `tier='raid_guild'`, the premium row has `tier is null`, and down then up is clean.
  - `TestUpsertStripe_WritesPriceFacts`: a guild upsert with `Tier: TierGuild, PriceLookupKey: "guild_founder_yearly", UnitAmountCents: 12000, TrialEnd: &t` reads back the same four fields through `GuildBilling`.
  - `TestUpsertStripe_GuildWithoutTierFails`
  - `TestPersonalBilling_ReturnsSubscriptionAndKey`
  - `TestGrant_GuildRowIsRaidGuild`
- [ ] **Step 2: Run the tests to verify they fail.** Run `cd api && TEST_DATABASE_URL=…/pricing_t1… go test -p 1 ./internal/db/ ./internal/entitlements/ -run 'LaunchPricing|WritesPriceFacts|WithoutTier|ReturnsSubscription|GuildRowIsRaidGuild'`. Expected: FAIL (no migration, unknown fields).
- [ ] **Step 3: Write the migration and the store changes.**
  - Put a comment at the top of the migration noting that 0053 belongs to the discord-bot lane.
  - Do not add `0054` to any `mergeBaseMigrations` list.
  - Billing's webhook, the existing `UpsertStripe` caller, still compiles: the new fields stay zero-valued until Task 6.
- [ ] **Step 4: Run the tests to verify they pass.** Run them together with the full `./internal/entitlements/` and `./internal/db/` packages. Expected: PASS.
- [ ] **Step 5: Commit.** `feat(entitlements): migration 0054 - tier, price facts, seats, trials, transfers, comebacks, notices`

### Task 2: The price catalog, the founder window, the gateway surface and stripe-setup's portals

**Files:**
- Create: `api/internal/billing/catalog/catalog.go`, `catalog_test.go`
- Modify: `api/internal/billing/gateway.go`, `stripe_gateway.go`, `fake_gateway.go`, `fake_gateway_test.go`; the call sites of the moved constants and of `CreatePortalSession` in `handler.go`, `webhook.go` and `reconcile.go`; `api/cmd/api/stripe_setup.go`, `stripe_setup_test.go`

**Interfaces:**
- Produces (`catalog`, a leaf package that imports nothing internal):
  - **Offer.** `type Offer string` with `OfferPremium = "premium"`, `OfferGuild = "guild"`, `OfferRaidGuild = "raid_guild"`.
    - `func ParseOffer(s string) (Offer, bool)`.
    - `func (o Offer) EntitlementPlan() string` returns "premium" or "guild".
    - `func (o Offer) Tier() string` returns "", "guild" or "raid_guild".
    - `func (o Offer) ProductID() string`.
  - **Interval.** `type Interval string` with `Monthly = "monthly"`, `Yearly = "yearly"`.
  - **Price.** `type Price struct { LookupKey string; Offer Offer; Interval Interval; Founder bool; AmountCents int64; Sold bool }`.
    - `var Prices []Price` lists all fourteen lookup keys.
    - `premium_monthly` and `premium_yearly` (the old Premium keys) are `Sold: false`.
    - `guild_monthly` and `guild_yearly` are the Guild list prices and are `Sold: true`.
  - **Founder window.** `var FounderWindowCloses = time.Date(2027, 1, 7, 8, 0, 0, 0, time.UTC)`. `func FounderOpen(now time.Time) bool` returns `now.Before(FounderWindowCloses)`.
  - **Lookups.**
    - `func CheckoutLookupKey(o Offer, i Interval, now time.Time) string` returns the founder key while the window is open and the list key after it closes.
    - `func Lookup(lookupKey string) (Price, bool)`.
    - `func UpgradeLookupKey(fromLookupKey string) (string, bool)` maps a Guild key to the Raid guild key with the same founder flag and interval. Anything else answers false.
    - `func PortalFamily(lookupKey string) (string, bool)` returns one of `premium`, `premium_founder`, `guild`, `guild_founder`, `raid_guild`, `raid_guild_founder`. The two legacy Premium keys map to `premium`.
    - `func FamilyLookupKeys(family string) []string` returns that family's two keys, monthly then yearly.
  - **Rule constants.** `const GuildTierMemberLimit = 25`, `GuildTrialDays = 7`, `SeatsGuild = 1`, `SeatsRaidGuild = 3`. `SeatMoveLock = 30 * 24 * time.Hour`, `SeatReturnedShown = 14 * 24 * time.Hour`, `ComebackValid = 30 * 24 * time.Hour`. `func SeatsForTier(tier string) int`.
  - **Moved constants.** Every `LookupKey*` and `ProductID*` constant moves from `gateway.go` into `catalog`. Billing refers to them as `catalog.X`, with no alias left behind.
- Produces (`billing.Gateway`):
  - `CheckoutParams` gains `TrialPeriodDays int64` and `TrialEnd *time.Time`. At most one is set; the real gateway puts it on `subscription_data`.
  - `CreatePortalSession(ctx, customerID, returnURL, configurationKey string) (string, error)`. The real gateway resolves the configuration whose metadata `fs_portal` equals the key. A missing configuration is an error.
  - `PortalSpec` gains `Key string`. `EnsurePortalConfiguration` creates the configuration tagged `fs_portal=<Key>` when none carries that key, and does nothing otherwise.
  - `PreviewPriceChange(ctx, subscriptionID, priceID string) (PriceChangePreview, error)`, with `PriceChangePreview{AmountDueCents, NextAmountCents int64; PeriodEnd time.Time}`. Implemented with `Invoices.CreatePreview`, the item swapped and `proration_behavior=always_invoice`.
  - `ChangePrice(ctx, subscriptionID, priceID string) error` updates the single item's price with `proration_behavior=always_invoice` and `payment_behavior=error_if_incomplete`.
  - `ArchivePrice(ctx, lookupKey string) error` sets `active=false`.
  - `StripeSubscriptionSummary` gains `PriceLookupKey string`.
  - `FakeGateway` records every new call (`PortalCall.ConfigurationKey`, `PriceChanges []PriceChangeCall`, `Archived []string`). Tests can make it fail (`ChangePriceErr error`) and return a canned `PriceChangePreview`.
- Produces (`cmd/api`):
  - `priceSpecs` is derived from `catalog.Prices`. Product names: Premium "Forever Sixty Premium", Guild "Forever Sixty Guild", Raid guild "Forever Sixty Raid Guild".
  - `func portalSpecs(publicBaseURL string) []billing.PortalSpec` returns six specs, one per family. Each lists exactly that family's two keys under its one product.
  - `runStripeSetup` ensures all six.

- [ ] **Step 1: Write the failing tests.**
  - `TestPrices_MatchProductSpec`: table-driven, the twelve launch amounts from product §3, by lookup key.
  - `TestCheckoutLookupKey_FounderBoundary`:
    - Premium yearly at `07:59:59Z` gives `premium_founder_yearly`; at `08:00:00Z` it gives `premium_yearly_launch`.
    - Raid guild monthly gives `raid_guild_founder_monthly`, then `raid_guild_monthly`.
  - `TestCheckoutLookupKey_NeverLegacy`: no offer, interval or time returns `premium_monthly` or `premium_yearly`.
  - `TestUpgradeLookupKey`: `guild_founder_yearly` gives `raid_guild_founder_yearly`; `guild_monthly` gives `raid_guild_monthly`; a Premium key answers false.
  - `TestPortalFamily`
  - `TestPortalSpecs_SixFamiliesOneProductEach`: the founder families list only founder keys.
  - `TestPriceSpecs_FromCatalog`
  - Fake gateway tests: the fake records `ConfigurationKey`, `TrialPeriodDays` and price changes.
- [ ] **Step 2: Run the tests to verify they fail.** Run `cd api && go test ./internal/billing/... ./cmd/api/ -run 'Prices|LookupKey|PortalFamily|PortalSpecs|PriceSpecs|Fake'`. Expected: FAIL (the package does not exist).
- [ ] **Step 3: Implement the catalog and the gateway changes.** Update every call site. Callers pass `""` as the portal configuration key for now; Task 5 sets the real key.
- [ ] **Step 4: Run `go build ./... && go test ./internal/billing/... ./cmd/api/`.** Expected: PASS.
- [ ] **Step 5: Commit.** `feat(billing): price catalog, founder window, portal families and the upgrade and archive gateway calls`

### Task 3: The billing emails

**Files:**
- Create: `api/internal/billingmail/billingmail.go`, `billingmail_test.go`, `price.go`, `price_test.go`

**Interfaces:**
- Consumes: `mail.Message{To, Subject, Text}`.
- Produces:
  - **Price formatting.** `func FormatPrice(cents int64, interval string) string`:
    - interval "monthly" or "yearly" gives `$5 a month` or `$40 a year`;
    - interval "" gives the bare amount, such as `$3.33`, `$10` or `$12.50`. Cents show only when the amount is not whole.
  - **Dates** (UTC). `func EmailDate(t time.Time) string` gives `6 January 2027` (`d MMMM yyyy`). `func EmailDayMonth(t time.Time) string` gives `6 January` (`d MMMM`).
  - **Emails**, one function per email:

    | Input struct | Function |
    |---|---|
    | `Renewal{To, Plan, Guild string; RenewsOn time.Time; AmountCents int64; Interval string; Founder bool; ManageURL string}`. `Guild` is empty for Premium. | `RenewalEmail(Renewal) mail.Message` |
    | `Size{To, Guild string; Members int; RenewsOn time.Time; RaidGuildCents int64; Interval string; SettingsURL string}` | `SizeEmail` |
    | `SeatOffered{To, Guild, From string; OwnEndsOn time.Time; AccountURL string}` | `SeatOfferedEmail` |
    | `SeatGiven{To, Guild, From, AccountURL string}` | `SeatGivenEmail` |
    | `SeatRemoved{To, Guild string; EndedOn time.Time; AmountCents int64; Interval string; ComebackURL string}` | `SeatRemovedEmail` |
  - **Body shape.** The first sentence, a blank line, then `<link label>: <url>`. Plain text, no exclamation marks. The sender is `cfg.MailFrom`; the controller confirms it reads `Forever Sixty <help@foreversixty.gg>` (Task 22).

- [ ] **Step 1: Write the failing tests.**
  - Each of the five emails' subject and first sentence equals spec §6.4 for one fixture: Guild `Olympus XXVII`, name `Kraggor`, date 18 November 2027, `$120 a year` founder.
  - For example, `RenewalEmail` gives the subject `Your Guild plan renews on 18 November` and the first line `Your Guild plan for Olympus XXVII renews on 18 November 2027 for $120 a year, your founder price; to stop it, cancel before then and nothing more is charged.`
  - The Premium renewal has no ` for {Guild}`, and has no founder tail at a list price.
  - `TestFormatPrice`: `$3.33`, `$10`, `$12.50`, `$5 a month`, `$250 a year`.
  - `TestNoExclamationMarks` across all five emails.
- [ ] **Step 2: Run the tests to verify they fail.** Run `go test ./internal/billingmail/`. Expected: FAIL.
- [ ] **Step 3: Implement.** One function per email, each under 50 lines.
- [ ] **Step 4: Run the tests to verify they pass.**
- [ ] **Step 5: Commit.** `feat(billingmail): renewal, size and seat emails`

### Task 4: Entitlement resolution — seats, not membership; the raid-guild tools feature

**Files:**
- Create: `api/internal/entitlements/launch_resolution_test.go`
- Modify: `api/internal/entitlements/entitlements.go`; `api/internal/entitlements/store.go` (`hasStandardEntitlement`, `canOfficerViews`, `CanForGuild`); `api/internal/entitlements/store_test.go` (rewrite only the cases asserting "guild plan gives every member Premium"); test fixtures in `api/internal/guilds/*_test.go` and `api/internal/auth/*_test.go` that relied on member Premium

**Interfaces:**
- Consumes: the Task 1 schema and `TierRaidGuild`.
- Produces:
  - `const FeatureRaidGuildTools Feature = "raid_guild_tools"`, a per-guild feature.
  - `CanForGuild(ctx, guildID, userID, feature)` accepts `FeatureOfficerViews` and `FeatureRaidGuildTools`. For `FeatureRaidGuildTools` the row must be covered for officer views *and* have `tier='raid_guild'`. Reasons are unchanged.
  - Standard features and `IsSupporter` resolve true for an active personal premium row, or for a held seat. A held seat is a `premium_seats` row that meets all of these:
    - `holder_user_id=$1`;
    - `state in ('held','accepted')` and `(starts_at is null or starts_at <= now())`;
    - `slot` no greater than the seat count of the guild's active whole-plan row (`features is null`): 3 for raid_guild, 1 for guild, rendered into the SQL from `SeatsGuild` / `SeatsRaidGuild`;
    - the holder is still a verified member of that guild.
  - Guild membership alone grants nothing.
  - `canOfficerViews` keeps answering `ReasonNotOfficer` for a verified non-officer member of a covered guild. It gets there through a guild-only check, not the standard path.

- [ ] **Step 1: Write the failing tests.**
  - `TestCan_MemberWithoutSeatHasNoPremium`: a verified member of a guild on an active Stripe Guild plan has no `server_sims` and no supporter mark.
  - `TestCan_HeldSeatGrantsPremium`
  - `TestCan_AcceptedSeatStartsAtStartsAt`: false before `starts_at`, true after.
  - `TestCan_SeatBeyondTierCountIgnored`: slot 2 on a Guild tier.
  - `TestCan_SeatEndsWithPlan`: plan canceled, no Premium.
  - `TestCan_SeatHolderNoLongerVerified`: `verified_at` null, no Premium.
  - `TestCan_ScopedGrantGivesNoSeats`
  - `TestCan_OpenEndedPersonalGrantSurvives`: `source='grant'`, `current_period_end` null; the user also holds a seat elsewhere and is a member of a guild with a plan. Every standard feature is true.
  - `TestCanForGuild_RaidGuildTools`: Guild tier false/`no_plan`; Raid guild true; scoped grant true; lapsed false.
- [ ] **Step 2: Run the tests to verify they fail.** Run `go test -p 1 ./internal/entitlements/ -run 'Can_|CanForGuild_RaidGuildTools'` on `pricing_t4`. Expected: FAIL.
- [ ] **Step 3: Implement the SQL changes.** Keep `activeClause` as the one live-row predicate.
- [ ] **Step 4: Run the full `./internal/entitlements/` package.** Then run `./internal/guilds/ ./internal/auth/`. Fix fixtures that relied on member Premium by giving that account a seat or a personal row. Expected: PASS.
- [ ] **Step 5: Commit.** `feat(entitlements): Premium from a personal plan or a held seat; raid_guild_tools per guild`

### Task 5: Checkout and the portal on the launch prices

**Files:**
- Create: `api/internal/billing/checkout_store.go`, `api/internal/billing/checkout_launch_test.go`, and `api/internal/billing/checkout_guild.go` if `handler.go` would pass 800 lines
- Modify: `api/internal/billing/handler.go` (`checkoutRequest`, remove `lookupKeyFor`, `checkout`, `premiumCheckout`, `guildCheckout`, `checkGuildCheckout`, `portal`, `Service.Now`), `api/internal/billing/harness_test.go` (seed helpers for trials, transfers and comebacks)

**Interfaces:**
- Consumes: `catalog` (Task 2) and the `GuildBilling` / `Billing` price facts (Task 1).
- Produces:
  - **Request body.** `{plan: "premium"|"guild"|"raid_guild", interval, guild_id?, intent?: "transfer", comeback_id?}`. Any other plan answers 400.
  - **Clock.** `Service.Now func() time.Time`; nil means `time.Now`.
  - **Lookup key.** The first rule that applies wins:
    1. Comeback: use the comeback row's `price_lookup_key`. The row must belong to the actor, be unexpired and be unused; otherwise 404. It is marked used when the premium webhook lands (Task 6).
    2. Transfer: use the guild row's current `price_lookup_key`.
    3. Everything else: `catalog.CheckoutLookupKey(offer, interval, s.now())`.
  - **Metadata.** `plan = offer.EntitlementPlan()`, so Raid guild sends `"guild"`. Also `comeback_id` and `from_subscription_id` when present.
  - **Guild trial.** Set `TrialPeriodDays = catalog.GuildTrialDays` when the guild has no `guild_trials` row and this is not a transfer.
  - **Transfer.**
    - Refused with 409 `no_transfer_offer` unless `plan_transfers` has `state='offered', to_user_id=actor`.
    - `TrialEnd = max(row.CurrentPeriodEnd, now+49h)`, with `transferTrialFloor = 49 * time.Hour`.
    - This retires the old any-officer "Take over billing" path.
  - **Portal.** The configuration key is `catalog.PortalFamily(row.PriceLookupKey)` of the personal row, or of the guild row when `guild_id` is given. A row with no key uses its offer's list family.
  - **Store** (`checkout_store.go`):
    - `(*Store) GuildTrialUsed(ctx, guildID int64) (bool, error)`
    - `(*Store) Comeback(ctx, id, userID int64, now time.Time) (lookupKey string, ok bool, err error)`
    - `(*Store) OpenTransferTo(ctx, guildID, toUserID int64) (fromSubscriptionID string, ok bool, err error)`

- [ ] **Step 1: Write the failing tests** in `checkout_launch_test.go`, using a fake clock and `FakeGateway`:
  - `TestPremiumCheckout_FounderKeyWhileOpen` (`premium_founder_yearly`)
  - `TestPremiumCheckout_ListKeyAfterClose` (`premium_yearly_launch`)
  - `TestCheckout_NeverSellsLegacyPremium`
  - `TestRaidGuildCheckout_MetadataPlanGuild`: key `raid_guild_founder_monthly`, `metadata.plan == "guild"`.
  - `TestGuildCheckout_TrialOnce`: 7 days the first time, 0 when a `guild_trials` row exists.
  - `TestTransferCheckout_RequiresOffer`: 409.
  - `TestTransferCheckout_KeepsFounderAfterClose`: the old row is on `guild_founder_yearly` and the clock is after the window closes. The session uses `guild_founder_yearly` with no trial days.
  - `TestTransferCheckout_TrialEndFloor`: the period ends in 10 hours, so `TrialEnd == now+49h`.
  - `TestComebackCheckout_UsesOldKey`; 404 for another user's comeback or an expired one.
  - `TestPortal_FounderConfiguration`: a row on `premium_founder_monthly` gives `ConfigurationKey == "premium_founder"`.
- [ ] **Step 2: Run the tests to verify they fail.** Run `go test -p 1 ./internal/billing/ -run 'Checkout|Portal'` on `pricing_t5`. Expected: FAIL.
- [ ] **Step 3: Implement.** Keep the existing locks and refusals.
- [ ] **Step 4: Run the full `./internal/billing/` package.** Update existing tests that asserted `premium_monthly` or the old transfer path. Expected: PASS.
- [ ] **Step 5: Commit.** `feat(billing): checkout sells launch and founder prices, Raid guild, the 7-day guild trial, offer-only transfers, comebacks; portal by family`

### Task 6: The webhook and reconcile learn tiers, trials and transfers

**Files:**
- Create: `api/internal/billing/webhook_store.go`, `api/internal/billing/webhook_launch_test.go` (its own seeds; does not edit `harness_test.go`)
- Modify: `api/internal/billing/webhook.go` (`stripeUpsertFor`, `applyUpsert`), `api/internal/billing/reconcile.go`, `api/internal/billing/reconcile_test.go`

**Interfaces:**
- Consumes: `catalog.Lookup`, the `StripeUpsert` price fields (Task 1), `Gateway.CancelSubscriptionAtPeriodEnd`.
- Produces:
  - **`stripeUpsertFor`** fills four fields:
    - `PriceLookupKey` from `sub.Items.Data[0].Price.LookupKey`;
    - `UnitAmountCents` from `Price.UnitAmount`;
    - `TrialEnd` from `sub.TrialEnd` when it is non-zero;
    - for guild rows, `Tier` from `catalog.Lookup(key).Offer.Tier()`. An unknown key on a guild subscription logs a warning with the price id and uses `TierGuild`.
  - **After a non-skipped upsert:**
    - a guild subscription with `TrialEnd` that is not a transfer calls `RecordGuildTrial`;
    - a transfer with metadata `from_subscription_id` calls `CancelSubscriptionAtPeriodEnd(from)` and `MarkTransferAccepted`, both idempotent;
    - a premium subscription with metadata `comeback_id` calls `MarkComebackUsed`.
  - **Store** (`webhook_store.go`):
    - `RecordGuildTrial(ctx, guildID int64, subscriptionID string) error`, on conflict do nothing.
    - `MarkTransferAccepted(ctx, guildID int64, fromSubscriptionID string) error`.
    - `MarkComebackUsed(ctx, id int64) error`.
    - `OutgoingTransferSubscriptionIDs(ctx) ([]string, error)` returns the `from_subscription_id` of accepted transfers.
  - **Reconcile.** `reconcileProducts` gains `catalog.ProductIDRaidGuild`; its orphans are recorded with plan `guild`. Outgoing-transfer ids count as known.

- [ ] **Step 1: Write the failing tests.**
  - `TestWebhook_RaidGuildTierFromPrice`
  - `TestWebhook_FounderFactsStored`: key and 12000 cents.
  - `TestWebhook_TrialRecordedOnce`: two deliveries leave one `guild_trials` row.
  - `TestWebhook_TransferCancelsOutgoing`: the fake records exactly one cancel of `from` across two deliveries, and the transfer row is `accepted`.
  - `TestWebhook_ComebackMarkedUsed`
  - `TestReconcile_ScansRaidGuild`
  - `TestReconcile_OutgoingTransferIsNotOrphan`
- [ ] **Step 2: Run the tests to verify they fail.** Run `go test -p 1 ./internal/billing/ -run 'Webhook_|Reconcile_'` on `pricing_t6`. Expected: FAIL.
- [ ] **Step 3: Implement.**
- [ ] **Step 4: Run the full `./internal/billing/` package.** Expected: PASS.
- [ ] **Step 5: Commit.** `feat(billing): webhook stores tier and price facts, records trials, completes transfers; reconcile covers Raid guild`

### Task 7: The seats package

**Files:**
- Create:
  - `api/internal/seats/store.go`, `service.go`, `release.go`, `handler.go`, `me.go`
  - tests `store_test.go`, `service_test.go`, `handler_test.go`, `harness_test.go`
  - `api/internal/auth/seat_facts.go`: the struct only, exactly as in Task 9's Interfaces, so the two tasks compile independently
- Modify: `api/cmd/api/main.go` (construct `seats.Service`, call `seats.Mount`)

**Interfaces:**
- Consumes:
  - `catalog`: seat counts, locks, `Lookup`.
  - `billingmail` (Task 3).
  - `entitlements.Store.PersonalBilling` with the Task 1 fields.
  - `auth.RequireSession`, `auth.ActorFrom`, `httpx`.
  - `Canceller { CancelSubscriptionAtPeriodEnd(ctx, id string) error }`, which `billing.Gateway` satisfies.
  - `FrozenClaimant`, the same shape as billing's.
- Produces:
  - **Types.**
    - `type State string`: `StateEmpty`, `StateHeld`, `StateOffered`, `StateAccepted`.
    - `type Holder struct { UserID int64; CharacterKey, Name, Class string }`. The character is the holder's main if it is in the guild, otherwise their first verified character there.
    - `type Seat struct { Slot int; State State; Holder *Holder; GivenAt, StartsAt, MovableAt *time.Time; Returned *Returned }`.
    - `Returned{Name string; At time.Time}` is set while `returned_at > now - SeatReturnedShown`.
    - Seat JSON is snake_case: `slot`, `state`, `holder{user_id,character_key,name,class}`, `given_at`, `starts_at`, `movable_at`, `returned{name,at}`.
    - `type Candidate struct { UserID int64; Name, Class string; PaysPremium bool }`.
  - **Store reads.**
    - `(*Store) List(ctx, guildID int64, now time.Time) ([]Seat, error)` returns slots 1..`SeatsForTier` of the active whole-plan row; a slot with no row reads as empty. No plan or a scoped grant gives an empty slice.
    - `(*Store) IsPayer(ctx, guildID, userID int64) (bool, error)`: true for a verified officer or leader who holds the claim or is the row's `billing_user_id`.
    - `(*Store) IsSubscriber(ctx, guildID, userID int64) (bool, error)`.
    - `(*Store) Candidates(ctx, guildID int64) ([]Candidate, error)` returns verified members holding no seat here, sorted by name. `PaysPremium` means an active `source='stripe'` personal premium row.
  - **Service actions.**

    | Method | Behavior |
    |---|---|
    | `(*Service) Give(ctx, guildID int64, slot int, actorID, toUserID int64) (Seat, error)` | A payer recipient gets an offer: state `offered`, `offered_at=now`, and the seat-offered email. A non-payer gets the seat at once: state `held`, `given_at=now`, `movable_at=now+SeatMoveLock`, and the seat-given email (R-P2, R-P3). |
    | `(*Service) Move(ctx, guildID int64, slot int, actorID, toUserID int64) (Seat, error)` | Refused with `ErrLocked{MovableAt}` while locked. Otherwise releases the old holder, then gives. |
    | `(*Service) Withdraw(ctx, guildID int64, slot int, actorID int64) (Seat, error)` | Only from `offered`. Starts no lock. |
    | `(*Service) Accept(ctx, guildID, userID int64) error` | If the user's own Stripe row is active: call `CancelSubscriptionAtPeriodEnd` (unless already set), set `starts_at = row.CurrentPeriodEnd`, and store `holder_gave_up_lookup_key = row.PriceLookupKey`. If they no longer pay, the seat starts now. Either way the state becomes `accepted` and `movable_at = now + SeatMoveLock`. |
    | `(*Service) Decline(ctx, guildID, userID int64) error` | The slot becomes empty. |
    | `(*Service) Sweep(ctx, now time.Time) (int, error)` and `(*Service) SweepGuild(ctx, guildID int64, now time.Time) error` | Release holders who are no longer verified members, recording `returned_name` and `returned_at`. Release every seat of a guild whose plan has no access. |
  - **Release** (`release.go`) runs under a guild advisory lock. When `holder_gave_up_lookup_key` is set, it inserts one `seat_comebacks` row (`expires_at = now + ComebackValid`) and sends the seat-removed email. The email's link is `{PublicBaseURL}/premium/checkout?plan=premium&interval=<key interval>&comeback=<id>`.
  - **Errors.**

    | Error | Response |
    |---|---|
    | `ErrNotPayer` | 403 |
    | `ErrNoPlan` | 409 `no_plan` |
    | `ErrNoSuchSlot` | 404 |
    | `ErrNotMember` | 409 `not_member` |
    | `ErrLocked` | 409 `seat_locked`, with field `movable_at` |
    | `ErrTaken` | 409 `conflict` |
    | `ErrNoOffer` | 404 |
    | Frozen claimant | 409 `claim_contested` |

    Every action refuses a frozen claimant.
  - **Routes.** All are wrapped in `auth.RequireSession` with `seatRateLimitPerHour = 60` per IP.
    - `GET /v1/guilds/{id}/seats/candidates` (payers)
    - `POST /v1/guilds/{id}/seats/{slot}/give` with body `{user_id}`
    - `POST /v1/guilds/{id}/seats/{slot}/move` with body `{user_id}`
    - `POST /v1/guilds/{id}/seats/{slot}/withdraw`
    - `POST /v1/guilds/{id}/seats/offer/accept` and `.../decline` (the offeree)
  - **Account facts** (`me.go`). `(*Store) SeatFactsForUser(ctx, userID int64, now time.Time) (auth.SeatFacts, error)`.

- [ ] **Step 1: Write the failing tests.**
  - `TestGive_NonPayerHeldWithLock`: `movable_at = now+30d`, and the seat-given email is sent.
  - `TestGive_PayerOffered`: no lock, seat-offered email.
  - `TestGive_ConcurrentSameSlot`: two goroutines; one holder, one `ErrTaken`.
  - `TestGive_RefusesNonPayerActor`
  - `TestGive_RefusesNonMember`
  - `TestMove_LockedReturnsMovableAt`
  - `TestMove_OldHolderKeepsUntilMove`
  - `TestWithdraw_StartsNoLock`
  - `TestAccept_CancelsOwnAtPeriodEnd`: the fake records the cancel; `starts_at` equals the period end; the gave-up key is stored; the lock starts at acceptance.
  - `TestDecline_EmptiesSlot`
  - `TestSweep_DepartedHolderReleasedOnce`: returned name and date; across two sweeps, exactly one comeback and one email.
  - `TestSweep_LapsedPlanReleasesSeats`
  - `TestList_SlotsByTier`: 1 and 3.
  - `TestList_ReturnedShownFor14Days`
  - `TestCandidates_PaysPremiumFlag`
  - Handler tests for each status code in the errors table.
- [ ] **Step 2: Run the tests to verify they fail.** Run `go test -p 1 ./internal/seats/` on `pricing_t7`. Expected: FAIL.
- [ ] **Step 3: Implement and wire `main.go`.**
  - The Stripe gateway is the `Canceller`.
  - With Stripe unconfigured, `Accept` for a payer answers 503 `billing_unavailable`.
  - The mailer is the existing `mail.NewResend`.
- [ ] **Step 4: Run the tests to verify they pass, plus `go build ./...`.**
- [ ] **Step 5: Commit.** `feat(seats): give, move, offer, accept, withdraw and release Premium seats`

### Task 8: The guild billing read, the upgrade and the transfer offer

**Files:**
- Create: `api/internal/billing/guild_view.go`, `upgrade.go`, `transfer.go`, `guild_routes.go`; tests `guild_view_test.go`, `upgrade_test.go`, `transfer_test.go`
- Modify: `api/internal/billing/handler.go` (`Service` gains `Seats SeatReader`), `api/cmd/api/main.go` (wire `Seats`, call `billing.MountGuild`)

**Interfaces:**
- Consumes:
  - `SeatReader { List(ctx, guildID int64, now time.Time) ([]seats.Seat, error); IsPayer(ctx, guildID, userID int64) (bool, error); IsSubscriber(ctx, guildID, userID int64) (bool, error); SweepGuild(ctx, guildID int64, now time.Time) error }`
  - `catalog`, `Gateway.PreviewPriceChange`, `Gateway.ChangePrice`, `upsertSubscription`.
- Produces:
  - **`GET /v1/guilds/{id}/billing`.**
    - Verified officers and leaders only; everyone else gets 404. It calls `SweepGuild` first.
    - Response:
      ```
      { billing: { scope, tier, status, current_period_end, cancel_at_period_end, trial_end,
                   founder, interval, amount_cents, billed_by, you_are_billing_contact } | null,
        trial_used, verified_count, seats: [Seat], transfer: TransferView | null,
        transfer_candidates: [{user_id, name}], can: [] }
      ```
    - `trial_used` is top-level because `billing` is null before a plan (a contract amendment; see Conflicts ruled).
    - `verified_count` is the live count of distinct verified accounts, leader included.
    - `can` holds `billing` and `seats` for a payer, plus `transfer` for the subscriber.
    - `transfer_candidates` lists verified officers other than the subscriber, filled for the subscriber only.
    - `TransferView{to_name, to_you, from_name, at, state, starts_at, amount_cents, interval, founder}`.
    - `founder` and `interval` come from `catalog.Lookup(price_lookup_key)`.
  - **`GET /v1/guilds/{id}/billing/upgrade-preview`.** Payers only. Answers `{amount_due_cents, next_amount_cents, interval, renews_at}`. Only a Guild-tier Stripe row can upgrade; anything else answers 409 `not_upgradable`.
  - **`POST /v1/guilds/{id}/billing/upgrade`.** Payers only.
    - Calls `ChangePrice` to `catalog.UpgradeLookupKey(row.PriceLookupKey)`, then `upsertSubscription`, so the tier is `raid_guild` before the response.
    - A gateway error answers 409 `upgrade_failed` and changes nothing.
    - Success answers with the full billing read.
  - **`POST /v1/guilds/{id}/billing/transfer`** with body `{to_user_id}`. Subscriber only. The target must be a verified officer of this guild other than the subscriber; otherwise 409 `not_member`. It upserts a `plan_transfers` row in state `offered` with the row's `stripe_subscription_id`.
  - **`POST .../billing/transfer/withdraw`** (the subscriber) and **`POST .../billing/transfer/decline`** (the receiver) both set the state to `declined`.
  - All of these routes use `billingRateLimitPerHour` and refuse a frozen claimant.

- [ ] **Step 1: Write the failing tests.**
  - `TestGuildBilling_NotOnPlan`: billing null, `trial_used` false, seats empty.
  - `TestGuildBilling_Trial`
  - `TestGuildBilling_FounderActive`: `founder: true`, `interval: "yearly"`, `amount_cents: 12000`.
  - `TestGuildBilling_ReadOnlyOfficerHasNoCan`
  - `TestGuildBilling_MemberGets404`
  - `TestGuildBilling_VerifiedCountIncludesLeader`
  - `TestUpgradePreview_GuildOnly`
  - `TestUpgrade_FounderLandsOnFounderRaidGuild`: the fake records `raid_guild_founder_yearly`, and the response has tier `raid_guild`.
  - `TestUpgrade_DeclinedLeavesGuild`
  - `TestTransfer_SendWithdrawDecline`
  - `TestTransfer_TargetMustBeOfficer`
  - `TestTransfer_ReceiverSeesToYou`
- [ ] **Step 2: Run the tests to verify they fail.** Run `go test -p 1 ./internal/billing/ -run 'GuildBilling_|Upgrade|Transfer_'` on `pricing_t8`. Expected: FAIL.
- [ ] **Step 3: Implement and wire.**
- [ ] **Step 4: Run the full `./internal/billing/` package and `go build ./...`.** Expected: PASS.
- [ ] **Step 5: Commit.** `feat(billing): guild billing read, Guild to Raid guild upgrade, plan transfer offers`

### Task 9: `/v1/me` — price facts, the seat, the offer, the ended seat, guild tier

**Files:**
- Create: `api/internal/auth/me_billing.go`, `api/internal/auth/me_billing_test.go`
- Modify: `api/internal/auth/handler.go` (move `BillingView` and `GuildBillingView` into `me_billing.go`; `Me` gains three fields; `me` calls `s.attachSeatFacts`), `api/internal/auth/store.go` (`Guilds` reads `verified_count`), `api/cmd/api/main.go` (wire `Seats`)

**Interfaces:**
- Consumes: `seats.Store.SeatFactsForUser` and `catalog`. `auth` imports `catalog`, never `billing`, to avoid an import cycle.
- Produces:
  - `auth/seat_facts.go`, created in Task 7:
    ```go
    type SeatFacts struct {
        Seat  *HeldSeat  // guild_name, given_by, since, starts_at
        Offer *SeatOffer // guild_id, guild_name, from_name, offered_at
        Ended *EndedSeat // guild_name, ended_at, comeback_id, price_lookup_key (shown 30 days)
    }
    ```
  - `Service.Seats interface{ SeatFactsForUser(ctx, userID int64, now time.Time) (SeatFacts, error) }`. When nil, the seat facts are left out.
  - **JSON additions:**
    - `entitlements.billing` gains `founder`, `interval` and `amount_cents`.
    - `seat`: `{guild_name, given_by, since, starts_at}` or null.
    - `seat_offer`: `{guild_id, guild_name, from_name, offered_at, own_ends_at, own_amount_cents, own_interval}` or null. The `own_*` fields come from the caller's own premium row.
    - `ended_seat`: `{guild_name, ended_at, resubscribe_url, amount_cents, interval}` or null, with `resubscribe_url = "/premium/checkout?plan=premium&interval=<i>&comeback=<id>"`.
    - Each guild entry gains `verified_count`. Its `plan` gains `tier` and `you_hold_seat`.

- [ ] **Step 1: Write the failing tests.**
  - `TestMe_BillingFounderFacts`
  - `TestMe_SeatHolder`
  - `TestMe_SeatOfferCarriesOwnEnd`
  - `TestMe_EndedSeatWithin30Days`
  - `TestMe_EndedSeatGoneAfter30Days`
  - `TestMe_GuildPlanTierAndSeat`
  - `TestMe_GuildVerifiedCount`
  - `TestMe_OwnerGrantStillPremium`: a user with an open-ended grant shows `server_sims` true and `billing.plan` premium.
- [ ] **Step 2: Run the tests to verify they fail.** Run `go test -p 1 ./internal/auth/ -run 'Me_'` on `pricing_t9`. Expected: FAIL.
- [ ] **Step 3: Implement.** `handler.go` must end smaller than it started.
- [ ] **Step 4: Run the full `./internal/auth/` package and `go build ./...`.** Expected: PASS.
- [ ] **Step 5: Commit.** `feat(auth): /v1/me carries founder facts, the seat, the seat offer and the ended seat`

### Task 10: Guild CSV exports and the raid-guild gate

**Files:**
- Create: `api/internal/guilds/export.go`, `api/internal/guilds/export_test.go`, `api/internal/guilds/raid_guild_gate_test.go`
- Modify:
  - `api/internal/guilds/handler.go`: two routes.
  - `api/internal/guilds/loot.go`: `requireGuildPlan` takes the feature; Award and Undo pass `FeatureRaidGuildTools`.
  - `api/internal/guilds/home.go`: `Viewer.RaidGuildTools bool \`json:"raid_guild_tools"\``.
  - `api/internal/guilds/supporter.go`: a `raidGuildTools` helper beside `officerTools`.
  - `api/internal/refusal/refusal.go`: `AwardPlanRequired` becomes `Loot council and the readiness board are part of Raid guild.`

**Interfaces:**
- Consumes: `entitlements.FeatureRaidGuildTools` (Task 4).
- Produces:
  - **Routes.** `GET /v1/guilds/{id}/export/attendance.csv` and `GET /v1/guilds/{id}/export/loot.csv`.
    - Session required. Open to verified officers, leaders and moderators, in any plan state, lapsed included.
    - `Content-Type: text/csv; charset=utf-8`.
    - `Content-Disposition: attachment; filename="{slug}-attendance.csv"` (or `-loot.csv`). The slug is the guild name lower-cased, with each run of non-alphanumerics replaced by `-`, then trimmed.
  - **Attendance CSV.** Header `date,raid,character,class,signed,outcome,excused,pulls_present,pulls_total`. Rows come from `guild_raid_attendance` joined to `guild_raids`, ordered by raid start, then name.
  - **Loot CSV.** Header `date,raid,boss,item_id,item,character`. Rows come from `loot_awards`; item and boss names are resolved the way `Store.Loot` resolves them.
  - **Safety and empty output.** A cell starting with `=`, `+`, `-` or `@` gets a `'` prefix, to block formula injection. An empty guild gives the header only.
  - **Unchanged gates.** The Raids, Discord and raid-write paths keep `FeatureOfficerViews`, which the Guild tier covers.

- [ ] **Step 1: Write the failing tests.**
  - `TestExportAttendance_HeaderAndRows`
  - `TestExportLoot_HeaderOnlyWhenEmpty`
  - `TestExport_LapsedGuildStillExports`
  - `TestExport_MemberGets404`
  - `TestExport_FormulaCellEscaped`
  - `TestExport_FilenameSlug`: `Olympus XXVII` gives `olympus-xxvii-loot.csv`.
  - `TestLootAward_GuildTierRefused`: 403 `guild_plan_required` with the Raid guild sentence.
  - `TestLootAward_RaidGuildAllowed`
  - `TestHome_RaidGuildToolsFlag`
- [ ] **Step 2: Run the tests to verify they fail.** Run `go test -p 1 ./internal/guilds/ -run 'Export|LootAward_|RaidGuildTools'` on `pricing_t10`. Expected: FAIL.
- [ ] **Step 3: Implement.**
- [ ] **Step 4: Run the full `./internal/guilds/` package.** Expected: PASS.
- [ ] **Step 5: Commit.** `feat(guilds): attendance and loot CSV export; loot council and readiness gated on Raid guild`

### Task 11: The `billing-daily` job — renewal and size notices, seat sweep

**Files:**
- Create: `api/internal/billing/notices.go`, `notices_store.go`, `notices_test.go`, `api/cmd/api/billing_daily.go`
- Modify: `api/cmd/api/main.go` (`case "billing-daily"`), `.github/workflows/api.yml` (add `billing-daily` to the job image list near line 229)

**Interfaces:**
- Consumes: `billingmail` (Task 3), `catalog`, `seats.Service.Sweep` (Task 7), `UserReader`.
- Produces:
  - **`(*Service) SendNotices(ctx, now time.Time) (int, error)`** sends two kinds of notice:
    - **Renewal.** For Stripe rows that are `active`, not `cancel_at_period_end`, yearly per `catalog.Lookup`, with `current_period_end` in `(now, now+7d]`.
      - Premium rows link to `{PublicBaseURL}/account#plan`.
      - Guild rows link to the guild's settings `#billing`. Build the URL from the guild's region, ruleset and name with the same path shape as the web's `guildSettingsHref` (`web/src/lib/characters`), as `billing.guildSettingsURL`.
    - **Size.** For guild rows with `tier='guild'`, `active` and not canceling, where `verified_count > catalog.GuildTierMemberLimit` on the send day and the period ends in `(now, now+14d]`. The price is Raid guild at the row's interval and founder flag (`catalog.UpgradeLookupKey`).
  - **Once per period.** Each send first claims `billing_notices (kind, entitlement_id, period_end)`; a conflict means already sent, so skip. A billing user with no email is logged at info and skipped.
  - **Store** (`notices_store.go`):
    - `DueRenewals(ctx, now) ([]DueNotice, error)`
    - `DueSizeChecks(ctx, now) ([]DueNotice, error)`
    - `ClaimNotice(ctx, kind string, entitlementID int64, periodEnd time.Time) (bool, error)`
  - **`runBillingDaily`** runs `seats.Service.Sweep`, then `SendNotices`. It logs counts only.

- [ ] **Step 1: Write the failing tests.**
  - `TestNotices_RenewalYearlyOnly`: a monthly row sends nothing.
  - `TestNotices_RenewalSkipsCanceling`
  - `TestNotices_RenewalOncePerPeriod`
  - `TestNotices_SizeAbove25Only`: 25 members sends nothing; 26 sends one, showing the Raid guild founder price for a founder row.
  - `TestNotices_SizeRaidGuildTierNever`
  - `TestNotices_NoEmailSkipped`
  - Subjects equal Task 3's fixtures, checked through `FakeMailer`.
- [ ] **Step 2: Run the tests to verify they fail.** Run `go test -p 1 ./internal/billing/ -run Notices_` on `pricing_t11`. Expected: FAIL.
- [ ] **Step 3: Implement, wire the subcommand, and add the job to `api.yml`.**
- [ ] **Step 4: Run the tests to verify they pass, plus `go build ./...`.**
- [ ] **Step 5: Commit.** `feat(billing): billing-daily job sends renewal and size notices and sweeps seats`

### Task 12: `stripe-archive-prices` (the controller runs it last)

**Files:**
- Create: `api/cmd/api/stripe_archive.go`, `api/cmd/api/stripe_archive_test.go`
- Modify: `api/cmd/api/main.go` (`case "stripe-archive-prices"`)

**Interfaces:**
- Consumes: `catalog.Prices` entries with `Sold == false`, `Gateway.ListSubscriptions` with `PriceLookupKey`, `Gateway.ArchivePrice`.
- Produces: `api stripe-archive-prices [--apply]`.
  - It lists live subscriptions on `prod_fs_premium`.
  - If any live subscription uses an unsold key, it exits non-zero and prints those subscription ids.
  - Otherwise it prints the keys it would archive, and archives them with `--apply`.
  - The logic is `func archiveUnsoldPrices(ctx, gw billing.Gateway, apply bool, out io.Writer) error`.

- [ ] **Step 1: Write the failing tests.**
  - `TestArchive_DryRunArchivesNothing`
  - `TestArchive_RefusesWhenLegacySubscriptionLive`
  - `TestArchive_ApplyArchivesOnlyUnsold`: exactly `premium_monthly` and `premium_yearly`, never `guild_monthly`.
- [ ] **Step 2: Run the tests to verify they fail.** Run `go test ./cmd/api/ -run Archive_`. Expected: FAIL.
- [ ] **Step 3: Implement.**
- [ ] **Step 4: Run the tests to verify they pass.**
- [ ] **Step 5: Commit.** `feat(billing): stripe-archive-prices retires the unsold Premium prices after a live-subscription check`

---

## PR B — the web

Every web task:
- Run `cd web && npm ci` once per worktree.
- Run the touched unit files with `npx vitest run <files>` and the task's e2e spec with `npx playwright test <spec>`.
- Run `npm run lint` and `npm run check` before committing.
- Tasks 16, 18 and 19 rebuild a page or a shared settings surface. They also run the full `npx vitest run`, the full `npx playwright test` and the lhci shard (memory: shared layout needs the full suite).

### Task 13: `plans.ts`, the plans-page copy and the checkout copy

**Files:**
- Modify: `web/src/lib/billing/plans.ts`, `plans.test.ts`, `web/src/lib/billing/copy.ts`, `copy.test.ts`, and every other `formatPrice` caller that `rg formatPrice web/src` finds
- Create: `web/src/lib/billing/plans-page-copy.ts`, `plans-page-copy.test.ts`

**Interfaces:**
- Produces:
  - **Types and data.**
    - `type PlanKey = 'premium' | 'guild' | 'raid_guild'`; `type Interval = 'monthly' | 'yearly'`.
    - `PLANS: Record<PlanKey, { key; name: 'Premium' | 'Guild' | 'Raid guild'; monthly; yearly; founderMonthly; founderYearly; guild: boolean; seats: 0 | 1 | 3 }>`.
    - `FOUNDER_WINDOW_CLOSES = '2027-01-07T08:00:00Z'`, `GUILD_TIER_MEMBER_LIMIT = 25`, `GUILD_TRIAL_DAYS = 7`.
  - **Functions.**

    | Function | Returns |
    |---|---|
    | `founderOpen(now: Date): boolean` | Whether the founder window is open at `now` |
    | `priceFor(plan, interval, founder: boolean): number` | The price in dollars |
    | `formatAmount(usd: number): string` | `$3.33`, `$10`, `$12.50` |
    | `formatPrice(usd, interval): string` | `$40 a year`, `$5 a month` |
    | `perMonthLine(yearlyUsd): string` | `$3.33 a month, billed yearly` |
    | `formatCents(cents, interval): string` | The same format as `formatPrice`, from cents |
  - **`plansPageCopy`** holds every §6.1 string except the checkout chooser:
    - page title and description, band, the founder open line, interval;
    - the four cards, with features as arrays;
    - price-block functions and buttons (`subscribe(price)`, `subscribeGuild(price)`);
    - trial, history promise, the eight `rules` as `{term, sentence}`, foot;
    - `faq`: today's list with the two §6.1 edits.
  - **`checkoutCopy`** gains `whichGuild`, `verifiedMembers(n)`, `fitsRaidGuild(n)`, `noBillableGuild`, `findYourGuild`, `guildsDidNotLoad` and `tryAgain`.
  - The old `premiumCopy` stays until Task 16 deletes it.

- [ ] **Step 1: Write the failing tests.**
  - `plans.test.ts`:
    - the twelve amounts;
    - `founderOpen` at `07:59:59Z` and at `08:00:00Z`;
    - `formatAmount(50/12) === '$4.17'`, `formatAmount(150/12) === '$12.50'`, `formatPrice(40,'yearly') === '$40 a year'`.
  - `plans-page-copy.test.ts`:
    - each §6.1 row through `specCell(spec, '### 6.1', '<label>')`: Band, the Founder row (open line), the Free, Premium, Guild and Raid guild cards, Trial, History promise, How it works (all eight sentences) and Foot;
    - a guard that no string contains `retention`, `compare more than` or `!`.
  - `copy.test.ts`: the Checkout chooser row.
- [ ] **Step 2: Run the tests to verify they fail.** Run `npx vitest run src/lib/billing/`. Expected: FAIL.
- [ ] **Step 3: Implement.** Each file stays within 400 lines.
- [ ] **Step 4: Run the tests to verify they pass.**
- [ ] **Step 5: Commit.** `feat(web): launch prices, founder window and the plans-page copy`

### Task 14: Guild billing, seats and account copy

**Files:**
- Create: `web/src/lib/billing/guild-billing-copy.ts`, `web/src/lib/billing/seats-copy.ts`, `web/src/lib/billing/account-plan-copy.ts`, each with a test
- Modify:
  - `web/src/lib/guild/settings-view.ts` and its test: `guildPlanLine` and `billedLine` move into `guild-billing-copy.ts`.
  - `web/src/components/guild/GuildBillingSection.svelte` and `web/src/components/account/AccountGuildPlanLine.svelte`: import updates only.

**Interfaces:**
- Consumes: `formatPrice`, `formatCents`, `PLANS` (Task 13); the existing UTC date helpers (`utcDayMonthYear` and its `d MMM` sibling).
- Produces:
  - **`guildBillingCopy`** holds the §6.2 Plan line, Line 2, Chooser, Size line, Upgrade confirm, Records and Gate rows:
    - `planLine(state: GuildPlanLineState): string`. `GuildPlanLineState` is a union:
      - `{kind:'trial', tier, firstCharge, amountCents, interval}`
      - `{kind:'active', tier, renews, amountCents, interval}`
      - `{kind:'ending', tier, ends}`
      - `{kind:'ending-transfer', tier, ends, toName}`
      - `{kind:'lapsed', ended}`
    - `line2(...)`, `upgradeSentence(prorated, interval, from, next)`, `records`.
  - **`seatsCopy`** holds the §6.2 Seats, Seat picker, Who pays and Receiver rows, with a function for every `{…}` slot.
  - **`accountPlanCopy`** holds the §6.3 rows: Offer, Plan line, Founder, Seat ended, Guild row.

- [ ] **Step 1: Write the failing tests.** Check every §6.2 and §6.3 row through `specString` with fixture fills. Examples:
  - `planLine({kind:'trial', tier:'guild', firstCharge:'2026-12-11', amountCents:12000, interval:'yearly'}) === 'Guild · trial · first charge 11 Dec 2026: $120'`
  - `seatsCopy.lockedError('2026-12-20') === 'That seat can move from 20 Dec.'`
- [ ] **Step 2: Run the tests to verify they fail.**
- [ ] **Step 3: Implement.**
- [ ] **Step 4: Run the tests to verify they pass, along with every test that imports `settings-view`.**
- [ ] **Step 5: Commit.** `feat(web): guild billing, seats and account plan copy`

### Task 15: Contract types and API clients

**Files:**
- Create: `web/src/lib/billing/contract.ts`, `web/src/lib/billing/guild-billing-api.ts` (+test), `web/src/lib/billing/seats-api.ts` (+test)
- Modify:
  - `web/src/lib/billing/api.ts` and its test: `CheckoutRequest.comebackId?: number`, sent as `comeback_id`; `BillingApiError` gains `readonly code: string | null` from the envelope.
  - `web/src/lib/account/api.ts`, imports and type fields only, with the types imported from `contract.ts`: `Me` gains `seat`, `seat_offer` and `ended_seat`; `BillingView` gains `founder`, `interval` and `amount_cents`; `GuildBillingView` gains `tier` and `you_hold_seat`; the guild entry gains `verified_count`.
  - `web/src/lib/guild/api.ts`: the home viewer gains `raid_guild_tools: boolean`.
  - `web/src/test-support/billing-fixtures.ts`: fixtures for every §9 e2e state.

**Interfaces:**
- Produces:
  - **`contract.ts`** mirrors Tasks 8 and 9 exactly: `GuildBillingSettings`, `GuildPlanView`, `SeatView`, `SeatState`, `TransferView`, `TransferCandidate`, `SeatCandidate`, `UpgradePreview`, `MeSeat`, `MeSeatOffer`, `MeEndedSeat`, `BillingFacts`.
  - **`guild-billing-api.ts`:** `getGuildBilling(guildId)`, `previewUpgrade(guildId)`, `upgrade(guildId)`, `sendTransfer(guildId, toUserId)`, `withdrawTransfer(guildId)`, `declineTransfer(guildId)`, `exportHref(guildId, kind: 'attendance'|'loot'): string`.
  - **`seats-api.ts`:** `seatCandidates(guildId)`, `giveSeat(guildId, slot, userId)`, `moveSeat(guildId, slot, userId)`, `withdrawSeat(guildId, slot)`, `acceptSeatOffer(guildId)`, `declineSeatOffer(guildId)`.
  - **Errors.** Every client function throws `BillingApiError` with `status` and `code` (for example `seat_locked`, `not_member`, `upgrade_failed`).

- [ ] **Step 1: Write the failing tests.**
  - For each client, a fetch stub checks the method, path and body.
  - `code` propagates for a 409 `seat_locked`.
  - `startCheckout` sends `comeback_id`.
- [ ] **Step 2: Run the tests to verify they fail.**
- [ ] **Step 3: Implement.**
- [ ] **Step 4: Run the tests to verify they pass, then `npm run check`.**
- [ ] **Step 5: Commit.** `feat(web): billing contract types and the guild billing and seats clients`

### Task 16: The `/premium` plans page

**Files:**
- Create: `web/src/components/premium/PlanCard.astro`, `PriceBlock.astro`, `FreeSummaryRow.astro`, `GuildRules.astro`, `IntervalControl.astro`
- Modify: `web/src/pages/premium.astro` (rebuilt), `web/lighthouserc.json` (add `http://localhost/premium.html` to the collect list), `web/tests/e2e/premium.spec.ts`, `web/src/lib/billing/copy.ts` (delete `premiumCopy`)

**Interfaces:**
- Consumes: `plansPageCopy`, `PLANS`, `priceFor`, `formatPrice`, `perMonthLine`, `FOUNDER_WINDOW_CLOSES` (Task 13).
- Produces:
  - **Layout** (spec §3.1):
    - four cards from 1024px, two from 768px;
    - below 768px the order is: the Free summary row, Premium, Guild, Raid guild, the history promise, then the full Free card with `id="free"`.
  - **No-JS state switching.** Both intervals and both founder states are always in the DOM, inside fixed-height price cells. CSS `:checked` picks the interval, and the root's `data-founder="open"|"closed"` picks founder or list prices.
  - **Head script.** An `is:inline` script of about six lines sets `data-founder` from `Date.now() < Date.parse('<FOUNDER_WINDOW_CLOSES>')`, interpolated at build time. There is no `before` state (the R-P8 ruling).
  - **Buttons** link to `/premium/checkout?plan=<plan>&interval=<i>`. Guild cards carry no `guild_id`, so they open the chooser (R-P7).
  - Today's `?canceled=1` script stays as it is.

- [ ] **Step 1: Write the failing e2e tests** in `premium.spec.ts`, with `page.clock.install`. They cover §9 cases 1–7 and 18, adjusted for the R-P8 ruling:
  - On 27 Oct 2026, the founder row reads the open line, prices are founder prices, and the list price is struck through.
  - On 8 Dec 2026, the same holds; `Yearly` is checked; the Premium button reads `Subscribe · $40 a year` and its href has `plan=premium&interval=yearly`.
  - Choosing `Monthly` changes every price and button, every card's bounding height stays the same, and the arrow keys also switch.
  - At 7 Jan 2027 08:00 UTC there is no founder row and no strike; prices are list (`Subscribe · $50 a year`). At 07:59:59 UTC the page still shows founder prices.
  - The page loads no `_astro` script, and `/pricing` answers 301.
  - At 360px wide, the Free summary row comes before Premium; Premium's button bottom is above y 600; `#free` contains `This list never shrinks.`
  - No text contains `retention`, `compare more than` or `!`.
  - axe reports no violations, and every control is at least 44px tall at 360px.
- [ ] **Step 2: Run the tests to verify they fail.** Run `npx playwright test tests/e2e/premium.spec.ts`. Expected: FAIL.
- [ ] **Step 3: Implement the components and the page.** No price literal in the markup. Each component stays within 400 lines.
- [ ] **Step 4: Run the spec, the full vitest suite, the full Playwright suite and the lhci shard.** The lhci budgets are performance and accessibility at least 0.95 and CLS at most 0.05. Expected: PASS.
- [ ] **Step 5: Commit.** `feat(web): the launch plans page - four plans, founder prices, the free promise`

### Task 17: The checkout launcher and the guild chooser

**Files:**
- Create: `web/src/components/checkout/CheckoutGuildChooser.svelte`, `CheckoutGuildChooser.test.ts`, `web/src/lib/billing/checkout-params.ts` (+test), `web/tests/e2e/checkout-guild.spec.ts`
- Modify: `web/src/components/CheckoutLauncher.svelte`

**Interfaces:**
- Consumes: `checkoutCopy` (Task 13); `startCheckout` with `comebackId` (Task 15); `me.guilds[].can`, `.plan` and `.verified_count`.
- Produces:
  - **`parseCheckoutParams(search: string): CheckoutParams | null`.** It accepts:
    - `plan` in `PlanKey`, and `interval`;
    - an optional `guild_id` (a positive integer);
    - an optional `comeback` (a positive integer, Premium only).

    Anything else returns null.
  - **The chooser.** A guild plan with no `guild_id` renders `CheckoutGuildChooser`, listing guilds whose `can` includes `billing` and whose `plan` is null.
    - Rows follow spec §3.1.7. The text on the right is `Guild` at 25 members or fewer, otherwise `Raid guild fits {n} members`.
    - Choosing a row continues with that guild id and the plan from the URL.

- [ ] **Step 1: Write the failing tests.**
  - `parseCheckoutParams` rejects `plan=gold`, `guild_id=-1`, and `comeback` on a guild plan.
  - The chooser component:
    - two guilds give two 56px rows;
    - 28 members gives `Raid guild fits 28 members`;
    - no billable guild gives the none line and the `Find your guild` link;
    - a load error gives its line and `Try again`.
  - E2E: §9 case 8 and case 18.
- [ ] **Step 2: Run the tests to verify they fail.**
- [ ] **Step 3: Implement.**
- [ ] **Step 4: Run the tests to verify they pass.**
- [ ] **Step 5: Commit.** `feat(web): checkout chooses the guild for a guild plan; comeback checkout`

### Task 18: Guild settings `#billing` — plan line, chooser, upgrade, size line, records

**Files:**
- Create:
  - in `web/src/components/guild/billing/`: `GuildPlanLine.svelte`, `GuildPlanChooser.svelte`, `IntervalSegments.svelte` (the G1 `SignupSegments` recipe on native radios), `GuildUpgradeConfirm.svelte`, `GuildSizeLine.svelte`, `GuildRecords.svelte`, each with a `.test.ts`;
  - `web/src/lib/guild/billing-state.ts` (+test);
  - `web/tests/e2e/guild-billing.spec.ts`.
- Modify:
  - `web/src/components/guild/GuildBillingSection.svelte` and its test. It becomes the orchestrator: it fetches `getGuildBilling`, reserves 420px (desktop) / 640px (phone) with `Skeleton` lines while loading, and renders the blocks in §3.2 order.
  - `web/src/components/GuildSettings.svelte`: pass `guildId`; drop the `billing` and `mayBill` props.

**Interfaces:**
- Consumes: Tasks 14 and 15.
- Produces:
  - `billingState(view: GuildBillingSettings | null, now: Date): GuildPlanLineState | {kind:'none'} | {kind:'grant'}`.
  - `fitsTier(count): 'guild' | 'raid_guild'`.
  - `showSizeLine(view): boolean` (tier guild and more than 25 members).
  - The section leaves two mount points, the seats block and the Who pays block, for Task 19 to fill.

- [ ] **Step 1: Write the failing tests.**
  - `billingState`: one unit test per kind.
  - Component tests:
    - the chooser preselects Raid guild at 28 members;
    - when the trial is used, the button reads `Subscribe · $120 a year` and there is no trial line;
    - the founder line shows while the window is open;
    - the upgrade confirm shows a `Skeleton` and keeps the commit button disabled until the preview lands; focus is on `Keep Guild`; Escape closes it;
    - a 409 `upgrade_failed` shows its line;
    - the size line has exactly one `Move to Raid guild`.
  - E2E: §9 cases 9, 10, 11 and 14, plus case 18. Case 14 is the lapsed state: `Subscribe again`, and both CSV downloads with the right file names and `text/csv`.
- [ ] **Step 2: Run the tests to verify they fail.**
- [ ] **Step 3: Implement.**
- [ ] **Step 4: Run them, then the full vitest and Playwright suites and the lhci shard.** `#billing` CLS must be at most 0.02 at 390 and 1440. Expected: PASS.
- [ ] **Step 5: Commit.** `feat(web): guild billing - plan line, plan chooser with the trial, Raid guild upgrade, size line, CSV records`

### Task 19: Seats and Who pays

**Files:**
- Create in `web/src/components/guild/billing/`: `GuildSeats.svelte`, `GuildSeatRow.svelte`, `GuildSeatPicker.svelte`, `GuildWhoPays.svelte`, `GuildTransferOffer.svelte`, each with a `.test.ts`
- Modify: `web/src/components/guild/GuildBillingSection.svelte` (mount points only), `web/tests/e2e/guild-billing.spec.ts` (add cases)

**Interfaces:**
- Consumes: `seatsCopy` (Task 14); `seats-api.ts` and `guild-billing-api.ts` (Task 15); the section's `view` and `refresh()`.
- Produces:
  - Seat rows in all five §3.2.5 states.
  - A picker that loads `seatCandidates` and shows the payer note when a `PaysPremium` member is chosen.
  - Error lines: `seat_locked` gives `That seat can move from {d MMM}.`; `not_member` gives `{Name} is no longer a verified member.`
  - **Who pays.** The subscriber sees send, waiting, withdraw and declined. The receiver block shows when `transfer.to_you` is set. `Take over the plan` calls `startCheckout({plan: view.billing.tier, interval: view.billing.interval, guildId, intent: 'transfer'})`.
  - Officers whose `can` lacks `seats` see the rows without the right-hand column.

- [ ] **Step 1: Write the failing tests.**
  - Component tests for each state:
    - a locked `Move seat` is `aria-disabled` and shows its date as visible text;
    - an empty picker reads `Every verified member already has Premium from this guild.`;
    - focus lands on the select, and Escape cancels;
    - with no transfer candidates, the transfer button is disabled and shows its line.
  - E2E: §9 cases 12, 13 and 15, plus case 18.
- [ ] **Step 2: Run the tests to verify they fail.**
- [ ] **Step 3: Implement.**
- [ ] **Step 4: Run them, then the full vitest and Playwright suites.**
- [ ] **Step 5: Commit.** `feat(web): Premium seats and plan transfer in guild billing`

### Task 20: The account `Plan` panel

**Files:**
- Create: `web/src/components/account/AccountSeatOffer.svelte` (+test), `web/src/lib/account/plan-line.ts` (+test), `web/tests/e2e/account-seat.spec.ts`
- Modify: `web/src/components/account/AccountPremiumBlock.svelte` (+test), `web/src/components/account/AccountGuildPlanLine.svelte` (+test)

**Interfaces:**
- Consumes: `accountPlanCopy` (Task 14); `acceptSeatOffer` and `declineSeatOffer` (Task 15); the `Me` additions.
- Produces:
  - **`accountPlanLine(me: Me, now: Date)`** returns `{line: string; line2: string | null; founder: boolean; manageBilling: boolean; ended: {text, href} | null}`. Its states are: own Premium, own Premium ending into a seat, seat, free, and seat ended.
  - **The offer block** comes first in `#plan`. It auto-focuses nothing, and becomes a status line after accept or decline.
  - **The `Founder` `pill-site`** carries a title and an sr-only tail, and appears on account surfaces only.
  - **Guild rows** read `Guild plan` or `Raid guild plan`, with ` · you hold a seat` added when the account holds one.

- [ ] **Step 1: Write the failing tests.**
  - Unit: one test per `accountPlanLine` state.
  - Component:
    - accepting shows `Accepted. Your Premium ends {date}; the seat starts then.`;
    - declining shows its status line;
    - a failed call shows the error line.
  - E2E: §9 case 17 and case 18.
- [ ] **Step 2: Run the tests to verify they fail.**
- [ ] **Step 3: Implement.** `AccountPremiumBlock.svelte` stays within 400 lines.
- [ ] **Step 4: Run the tests to verify they pass, plus `account-billing.spec.ts`.**
- [ ] **Step 5: Commit.** `feat(web): account plan panel - seat offer, seat holder, founder mark, ended seat`

### Task 21: The Guild-tier gate sentence

**Files:**
- Modify:
  - `web/src/lib/guild/copy.ts` and its test: add `guildGateCopy.raidGuildSentence = 'Loot council and the readiness board are part of Raid guild. '` and `raidGuildLink = 'Move to Raid guild'`, pinned to the §6.2 `Gate, Guild tier` row.
  - `web/src/components/guild/GuildLoot.svelte`, `web/src/components/guild/GuildReadiness.svelte`, and their tests.
- Create: `web/tests/e2e/guild-raid-guild-gate.spec.ts`

**Interfaces:**
- Consumes: `viewer.officer_tools` and `viewer.raid_guild_tools` (the type is in Task 15, the field in Task 10's API).
- Produces:
  - On Loot and Readiness, Award, Undo and Nudge are enabled only when `raid_guild_tools` is true.
  - When `officer_tools` is true and `raid_guild_tools` is false, `GuildPlanGate` shows the Raid guild sentence with its link to `#billing`.
  - A lapsed or never-subscribed guild keeps today's sentence.

- [ ] **Step 1: Write the failing tests.** Component tests for the three gate states on each tab. E2E: §9 case 16.
- [ ] **Step 2: Run the tests to verify they fail.**
- [ ] **Step 3: Implement.**
- [ ] **Step 4: Run the tests to verify they pass.**
- [ ] **Step 5: Commit.** `feat(web): Guild-tier gate sentence on Loot and Readiness`

---

### Task 22: Ship, then retire the old Premium prices (the controller, by hand)

The controller does this task by hand; no subagent.

- [ ] **Step 1:** Confirm that the discord-bot lane's 0053 is merged and deployed before PR A deploys. If it is not, PR A waits.
- [ ] **Step 2:** Confirm that `MAIL_FROM` on the API service reads `Forever Sixty <help@foreversixty.gg>`.
- [ ] **Step 3:** Merge PR A. Run `stripe-setup` from the new image against sandbox. Expected: six `fs_portal` configurations; the price lookups are unchanged.
- [ ] **Step 4:** In sandbox, run Checkout for Premium yearly and for Raid guild monthly on a test guild. Confirm the keys are `premium_founder_yearly` and `raid_guild_founder_monthly`, the 7-day trial applies, and the founder portal opens.
- [ ] **Step 5:** Deploy PR A, then run `stripe-setup` against live. Portal sessions answer an error until it has run, which is acceptable because no live subscriber exists yet.
- [ ] **Step 6:** Create the `billing-daily` Cloud Run job and its daily Cloud Scheduler trigger, shaped like `stripe-reconcile`. Then deploy PR B.
- [ ] **Step 7:** Confirm in production that user 78's `/v1/me` still shows Premium.
- [ ] **Step 8:** Run `api stripe-archive-prices` (dry run) against live. It must list only `premium_monthly` and `premium_yearly`, and must find no live subscription on them. Then run it with `--apply` against live, then against sandbox.

---

## Parallel waves

Tasks in the same wave touch disjoint files. Two rules shape the waves:

- `api/cmd/api/main.go` is touched by Tasks 7, 8, 9, 11 and 12, so no two of those share a wave.
- PR B codes against the Tasks 8 and 9 contract with stubbed APIs, so web waves do not wait on API waves.

| Wave | API | Web |
|---|---|---|
| 1 | 1, 2, 3 | 13, 15 |
| 2 | 4, 5, 6, 12 | 14, 16, 21 |
| 3 | 7, 10 | 17, 18, 20 |
| 4 | 8 | 19 |
| 5 | 9 | |
| 6 | 11 | |
| last | 22 (controller) | |

Two wave pairs share a directory but no files:
- Wave 2: Tasks 5 and 6 are both in `internal/billing`. Only Task 5 edits `harness_test.go`; Task 6 keeps its seeds in its own test file.
- Wave 2: Task 4 edits test fixtures in `internal/guilds` and `internal/auth`. Task 10 (`guilds`) and Task 9 (`auth`) run in later waves, on top of it.

## Conflicts ruled

- **R-P8 and R-P4 vs the controller.** The founder window has no opening instant. `data-founder` is `open` or `closed` only, and the `before:` copy is not rendered. E2E case 1 asserts founder prices on 27 Oct 2026. Product §3's "between 4 November 2026 and 6 January 2027" is superseded.
- **Portal configurations.** Product §8.2 gives founder customers their own configuration. The plan uses six configurations, one per product and founder flag. That keeps the founder lock and also closes cross-product switches, including today's Guild→Premium switch in the single default configuration. Every portal session names its configuration.
- **Transfer checkout mode.** R-P1 says setup mode, but setup mode cannot create a subscription. The plan uses subscription mode with `trial_end` at the current period end, floored at now+49h. Nothing is charged before then, which is what the copy promises.
- **Contract amendments** (reported in the lane report, as the spec asks):
  - `trial_used` is top-level in the guild billing read, because `billing` is null before a plan.
  - The billing read carries its own `can` (`billing`, `seats`, `transfer`). Payers are the claim holder or the billing user. `/v1/me`'s guild `can` keeps today's meaning.
  - The guild size is counted live at read and send time, not stored as a daily snapshot.
- **Existing guild rows and grants** become `raid_guild`, so comps and the grandfathering grants keep the loot council. Whole-plan grants hold 3 seats; scoped grants hold none. Members of granted guilds lose member-wide Premium, as product §5 intends.
- **"Take over billing"** (RULING 10, any officer) is retired. `intent=transfer` now requires an offer addressed to the caller.
- **Email links** go to `/account#plan` or the guild settings `#billing`, not to a portal session, because a portal session URL expires before the reader opens the email.
- **E2E file placement.** Case 8 lives in a new `checkout-guild.spec.ts` and case 16 in `guild-raid-guild-gate.spec.ts`, so no two parallel tasks edit the same spec file.
- **Migration 0054** must not deploy before the discord-bot lane's 0053 (Task 22 step 1).
