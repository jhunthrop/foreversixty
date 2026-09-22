# Pay API Security Fix Round Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development
> to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Fix one MEDIUM finding from an independent security review of the `pay-api`
branch, required before merge: `checkGuildCheckout` (`api/internal/billing/handler.go`)
reads a guild's entitlement with a plain SELECT and, unserialized against a second
concurrent request for the same guild, can let two officers each create a Stripe Checkout
Session and each pay — the second webhook's upsert on `(guild_id, plan)` then overwrites the
first subscription id, orphaning the first officer's subscription (it keeps billing with
nothing in our database pointing at it), and `stripe-reconcile` never finds it because it
only walks rows we already have.

**Architecture:** Three independent layers, each closing one part of the gap:
1. A Postgres advisory lock per guild id (`billing.Store.WithGuildLock`, mirroring the
   existing `WithEventLock`) serializes `checkGuildCheckout`'s read through Checkout Session
   creation, plus a `pending_checkouts` row written inside that lock so a second caller in
   the same window is refused exactly like an already-active plan.
2. `entitlements.Store.UpsertStripe` refuses to overwrite an ACTIVE row with a different
   `stripe_subscription_id`: it keeps the existing row, records the newcomer in a new
   `entitlement_anomalies` table, and returns enough for `billing.Service` to cancel the
   newcomer subscription at Stripe (at period end) through the `Gateway` interface.
3. `stripe-reconcile` becomes two-directional: `Gateway.ListSubscriptions` lists Stripe's own
   live subscriptions for our two product ids, and anything with no matching `entitlements`
   row is recorded as an `orphan_subscription` anomaly — spec §2.8's last paragraph,
   never implemented.

All in the same packages the branch already built (`api/internal/billing`,
`api/internal/entitlements`), plus one migration amendment (`0020` is unreleased and may be
edited in place, not renumbered — confirmed by the coordinator's dispatch).

**Tech Stack:** Go 1.25.11, pgx/v5, Postgres 16, stripe-go v82.

**Spec:** `docs/superpowers/specs/2026-09-21-entitlements-and-payments-design.md` — this plan
also amends §2.6 and §2.7 in place (operator runbook for an anomaly) as part of Task 6.

**Prior plan this extends:** the original `pay-api` lane build (already implemented,
reviewed, and merged into this branch at `adba349`, its own plan/ledger not committed — see
`.superpowers/sdd/.gitignore`).

## Global Constraints

- `go vet ./... && go test -p 1 ./...` (from `api/`) must pass before every commit; `-p 1`
  avoids spurious deadlocks from parallel packages sharing the test Postgres container.
  `TEST_DATABASE_URL` points at this lane's own throwaway database on the already-running
  `localhost:5434` server (created before this plan started; dropped at the end — never the
  shared `forever_test` database other lanes may also be using).
- File ownership: this lane owns `api/internal/billing/*.go`, `api/internal/entitlements/*.go`,
  `api/internal/db/migrations/0020_*`, `cmd/api/stripe_setup.go`,
  `docs/superpowers/specs/2026-09-21-entitlements-and-payments-design.md`. Nothing outside
  `api/` and this one spec doc.
- Migration `0020` is amended in place, not renumbered (unreleased, coordinator-confirmed).
- Commits: conventional subjects, message written to a file under this worktree's
  `.superpowers/` with `printf`, committed with `git commit -F <file>` as its own Bash
  command with nothing else in it, ending with exactly:
  `Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>`
  `Claude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5`
- Parameterised SQL only. Functions under 50 lines, files under 800, table-driven tests where
  the existing files already use them, errors wrapped `fmt.Errorf("<package>: <what>: %w", err)`.
- Failing test first for every behavioral change — this is a security fix round.
- **New shapes this plan introduces** (verbatim for another lane building against them):
  `billing.CheckoutSession{URL, SessionID, ExpiresAt}` (was a bare `string`);
  `billing.Gateway.CancelSubscriptionAtPeriodEnd(ctx, subscriptionID) error`;
  `billing.Gateway.ListSubscriptions(ctx, productID) ([]StripeSubscriptionSummary, error)`;
  `billing.ProductIDPremium`/`ProductIDGuild` constants;
  `entitlements.Store.UpsertStripe` now returns `(UpsertResult, error)`, where
  `UpsertResult{Duplicate bool, ExistingSubscriptionID string}`;
  `entitlements.Store.RecordAnomaly(ctx, userID, guildID *int64, plan, kind, stripeSubscriptionID, actor string) error`.

---

## Task 1: Migration 0020 amendment — `pending_checkouts`, `entitlement_anomalies`

**Files:** Modify `api/internal/db/migrations/0020_entitlements.up.sql`, `.down.sql`.

- [ ] Add `pending_checkouts (id, guild_id, user_id, stripe_session_id unique, created_at,
  expires_at)` with an index on `guild_id`.
- [ ] Add `entitlement_anomalies (id, entitlement_id nullable, user_id nullable,
  guild_id nullable, plan, kind check (kind in ('duplicate_subscription',
  'orphan_subscription')), stripe_subscription_id, actor, created_at)`.
- [ ] Down migration drops both, in dependency order, before the existing drops.
- [ ] `go vet ./...` and a throwaway `db.Migrate` round-trip (up via `testPool`, any existing
  test) confirm the migration applies cleanly.

## Task 2: `entitlements.Store.UpsertStripe` refuses to overwrite a live different subscription

**Files:** Modify `api/internal/entitlements/store.go`; update every caller
(`api/internal/billing/webhook.go`, `reconcile.go`, `handler.go`'s `Entitlements` interface);
update `api/internal/entitlements/store_test.go` and `api/internal/billing/*_test.go` call
sites for the new two-value return.

- [ ] Failing test first: `TestUpsertStripeKeepsAnExistingActiveRowOnADifferentSubscriptionID`
  — seed an active guild-plan row for `sub_a`, call `UpsertStripe` with `sub_b` for the same
  `(guild_id, plan)`, assert the row still reads `sub_a`, `UpsertResult.Duplicate == true`,
  `ExistingSubscriptionID == "sub_a"`, and one `entitlement_anomalies` row with
  `kind = 'duplicate_subscription'`.
- [ ] Implement: `UpsertResult` type; `UpsertStripe` reads the existing row inside its
  existing transaction, and when it is active (`activeStatusClause`) with a non-nil
  `stripe_subscription_id` different from `p.StripeSubscriptionID`, skips the write, records
  the anomaly, and returns early — no audit row for the skipped write (the anomaly row is
  the record of what happened instead).
- [ ] Add `RecordAnomaly` (pool-level, for `stripe-reconcile`'s orphan case in Task 5) reusing
  the same insert via a small `execer` interface shared with the transactional path (DRY).
- [ ] Update every existing `UpsertStripe` call site for the new `(UpsertResult, error)`
  return (tests only check `err` today — keep doing that, just discard the first value with
  `_` where a test does not care).

## Task 3: `billing.Store` guild advisory lock + pending-checkout bookkeeping

**Files:** Modify `api/internal/billing/store.go`; add cases to `api/internal/billing/store_test.go`.

- [ ] Failing test first: `TestWithGuildLockSerializesTwoCallers` — two goroutines call
  `WithGuildLock` for the same guild id with a channel-gated `fn`; assert they never run
  concurrently (a shared counter that must never exceed 1) and both eventually complete.
- [ ] Refactor `WithEventLock` into a shared `withAdvisoryLock(ctx, key string, fn) error`
  helper; add `WithGuildLock(ctx, guildID int64, fn) error` calling it with a
  `"guild-checkout:<id>"` key (distinct from a raw Stripe event id, so the two never collide
  under `hashtext`).
- [ ] Add `RecordPendingCheckout(ctx, guildID, userID int64, stripeSessionID string,
  expiresAt time.Time) error`, `PendingGuildCheckout(ctx, guildID int64) (bool, error)`
  (unexpired rows only), `DeletePendingCheckout(ctx, stripeSessionID string) error`,
  `SweepExpiredPendingCheckouts(ctx) (int64, error)`. Table-driven or per-behavior tests for
  each: record→exists→delete→gone; an expired row does not count as pending; sweep removes
  only expired rows.
- [ ] Update `store_test.go`'s `testPool` truncate list to include the two new tables.

## Task 4: `Gateway` interface — `CheckoutSession` return, `CancelSubscriptionAtPeriodEnd`, `ListSubscriptions`

**Files:** Modify `api/internal/billing/gateway.go`, `stripe_gateway.go`, `fake_gateway.go`,
`fake_gateway_test.go`.

- [ ] `CreateCheckoutSession` returns `CheckoutSession{URL, SessionID, ExpiresAt}` instead of
  a bare URL string, on both the interface and both implementations. `FakeGateway` gains a
  session-id counter and an optional `NextCheckoutExpiresAt` override (default `now +24h`).
- [ ] `CancelSubscriptionAtPeriodEnd(ctx, subscriptionID string) error` on the interface;
  `StripeGateway` calls `V1Subscriptions.Update` with `CancelAtPeriodEnd: true`; `FakeGateway`
  records every call in an exported `CanceledAtPeriodEnd []string` slice.
- [ ] `StripeSubscriptionSummary{ID, Status}` and
  `ListSubscriptions(ctx, productID string) ([]StripeSubscriptionSummary, error)` on the
  interface; `StripeGateway` lists non-canceled subscriptions and filters client-side by
  line-item product id (Stripe's List API has no product filter, only price); `FakeGateway`
  answers from an exported `ListSubscriptionsResult map[string][]StripeSubscriptionSummary`
  keyed by product id.
- [ ] `ProductIDPremium = "prod_fs_premium"`, `ProductIDGuild = "prod_fs_guild"` constants in
  `gateway.go`; `cmd/api/stripe_setup.go`'s `priceSpecs` references them instead of the
  string literals it currently duplicates (DRY — one source of truth for the two ids).
- [ ] Update `fake_gateway_test.go`'s `TestFakeGatewaySatisfiesGateway` for the new return
  shape; add `TestFakeGatewayRecordsCancelAndListSubscriptions`.

## Task 5: `handler.go` — serialize the guild checkout window, write the pending row

**Files:** Modify `api/internal/billing/handler.go`; extend `api/internal/billing/handler_test.go`.

- [ ] Failing test first: `TestCheckoutGuildSerializesConcurrentRequestsForTheSameGuild` —
  two verified officers of the same claimed, not-yet-subscribed guild fire
  `POST /v1/billing/checkout` concurrently (built before launching goroutines; only
  `h.do` runs inside them, per Go's `t.FailNow`-from-one-goroutine rule); assert exactly one
  `200` and one `409` (with `portal_hint`), and exactly one `CreateCheckoutSession` call on
  the fake gateway.
- [ ] Implement: extract the guild-plan branch of `checkout` into `guildCheckout`, wholly
  wrapped in `s.Store.WithGuildLock`: `checkGuildCheckout` (extended with a
  `PendingGuildCheckout` check returning the same `409 conflict`/`portal_hint` shape as an
  active plan) → `customerFor` → `PriceIDForLookupKey` → `CreateCheckoutSession` →
  `RecordPendingCheckout` with the session's own `ExpiresAt`. A small unexported
  `checkoutRefusal` error type carries the status/code/message/fields back out of the locked
  closure without conflating a refusal with an actual failure (`s.fail`'s 500 path).
- [ ] The non-guild (personal premium) path keeps its existing shape, just reading
  `session.URL` from the new `CheckoutSession` return.
- [ ] Re-run the full existing `handler_test.go` suite — every prior checkout/portal test
  must still pass unchanged in behavior (only the internal return type moved).

## Task 6: `webhook.go` — cancel the newcomer, expire the pending row; spec runbook

**Files:** Modify `api/internal/billing/webhook.go`; extend `api/internal/billing/webhook_test.go`;
amend `docs/superpowers/specs/2026-09-21-entitlements-and-payments-design.md` §2.7.

- [ ] Failing test first: `TestWebhookDuplicateSubscriptionForTheSameGuildIsKeptAndTheNewcomerCanceled`
  — seed an active guild-plan row for `sub_a` (as today's webhook would have written it),
  then deliver a `customer.subscription.created` for `sub_b` carrying the same
  `guild_id`/`plan` metadata; assert `sub_a` still reads in `entitlements`, one
  `entitlement_anomalies` row (`kind = 'duplicate_subscription'`, `stripe_subscription_id =
  'sub_b'`), and `sub_b` appears in the fake gateway's `CanceledAtPeriodEnd`.
- [ ] Implement: `upsertFromSubscription` inspects `UpsertStripe`'s `UpsertResult`; when
  `Duplicate`, calls `s.Gateway.CancelSubscriptionAtPeriodEnd(ctx, sub.ID)` and logs at error
  level with `sub.ID`/the kept subscription id only (no metadata dump — nothing here is a
  secret, but the discipline matches gateway.go's own "never log a body that could carry a
  key back out").
- [ ] `onCheckoutCompleted` reads the Checkout Session id (`event.Data.Object["id"]`, already
  what a `checkout.session.completed` event's object *is*) and calls
  `s.Store.DeletePendingCheckout` after a successful upsert (an aborted/duplicate upsert
  still had a real Checkout Session complete, so the pending row is cleared either way — a
  duplicate is caught at the entitlements layer, not by leaving the pending row dangling).
- [ ] Spec amendment: append an "Amendment, 2026-09-21 (security review response)" block
  under §2.7 stating the operator runbook for a `duplicate_subscription` anomaly — check
  `entitlement_anomalies`, refund the second (canceled-at-period-end) charge by hand in the
  Stripe Dashboard, matching the existing refund runbook's own pattern — and noting the two
  other layers (the advisory lock, two-directional reconcile) by name.

## Task 7: `reconcile.go` — sweep expired pending rows, two-directional orphan check

**Files:** Modify `api/internal/billing/reconcile.go`; extend `api/internal/billing/reconcile_test.go`.

- [ ] Failing test first: `TestReconcileReportsAStripeSubscriptionWithNoLocalRowAsAnOrphan`
  — populate the fake gateway's `ListSubscriptionsResult` for `ProductIDGuild` with a
  subscription id no `entitlements` row references; run `Reconcile`; assert one
  `entitlement_anomalies` row with `kind = 'orphan_subscription'` and that subscription id.
- [ ] Failing test first: `TestReconcileSweepsExpiredPendingCheckouts` — insert one expired
  and one live `pending_checkouts` row; run `Reconcile`; assert only the expired one is gone.
- [ ] Implement: `Reconcile` keeps its existing database→Stripe pass unchanged, then calls
  `s.Store.SweepExpiredPendingCheckouts`, then for each of `ProductIDPremium`/`ProductIDGuild`
  calls `s.Gateway.ListSubscriptions`, and for every returned id with no matching
  `entitlements.stripe_subscription_id`, calls `s.Entitlements.RecordAnomaly` with
  `kind = "orphan_subscription"`, `actor = "stripe_reconcile"`. Needs a small `Entitlements`
  interface addition (`RecordAnomaly`) and a `StripeSubscriptionIDSet`-style helper (or just
  a `map[string]bool` built from the existing `StripeSubscriptionIDs` read) to check
  membership without a second query per subscription.

## Task 8: Whole-branch verification and final review

- [ ] `go vet ./...` from `api/` — zero findings.
- [ ] `go test -p 1 ./...` from `api/` against this lane's throwaway database — full pass,
  including every new test above and every pre-existing one (nothing in `internal/billing`
  or `internal/entitlements`'s existing behavior may have changed observably except the
  `UpsertStripe`/`CreateCheckoutSession` return shapes callers already needed to follow).
- [ ] One `code-reviewer` pass and one `security-reviewer` pass over the full diff (`git diff
  adba349`), fresh subagents with no prior context, each given the finding text verbatim and
  asked to re-verify all three layers close it. One fix wave for anything CRITICAL/HIGH they
  raise; record every ruling in the ledger.
- [ ] Final report per `lane-common-go.md`: branch head sha, what shipped per spec section,
  final test run results, every ruling, anything left undone and why, exported
  types/signatures verbatim.
