# Billing takeover and webhook hardening Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** A "Take over billing" subscription can only replace a strictly older one, a dead newcomer never clobbers or cancels anything, the webhook writes from a subscription fetched inside the subject lock, and anomalies are recorded once per subscription.

**Architecture:** `entitlements.UpsertStripe` grows an outcome enum and a `created` comparison backed by a new `stripe_subscription_created` column; the duplicate guard orders its checks dead newcomer → transfer (newer wins, stale is ignored) → duplicate. `billing.upsertFromSubscription` re-fetches the subscription inside `WithGuildLock`/`WithUserLock` and acts on the outcome. A unique index makes `entitlement_anomalies` idempotent per `(kind, stripe_subscription_id)`.

**Tech Stack:** Go 1.2x, pgx v5, stripe-go v82, Postgres migrations under `api/internal/db/migrations`, the package test harnesses against `TEST_DATABASE_URL`.

**Spec:** `docs/superpowers/specs/2026-09-21-entitlements-and-payments-design.md` §2.6 (RULING 8, the 2026-09-21 amendment's layer 2), §2.8, §2.9 RULING 10. Tracking issue: GitHub #3 (items H2, M3, M4, L1 — the rest stay open in the issue).

## Global Constraints

- Only the deployed API applies migrations; the lane adds `0044_*.up.sql`/`.down.sql` and never runs them against production.
- No secrets in code, tests, logs or chat. Log lines carry subscription ids and plan names only (webhook.go:293 convention).
- Tests run with `TEST_DATABASE_URL='postgres://forever:forever@localhost:5434/forever_test?sslmode=disable'` and `-p 1` (shared truncating harness); `go vet` and `gofmt` clean.
- Every write of `entitlements` from Stripe state still goes through `UpsertStripe` (RULING 8); never trust an event's embedded object for the write.
- Commit messages: conventional type, trailer `Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>`.

## Review Focus

1. A transfer subscription B's later `customer.subscription.updated` (cancel_at_period_end=true) arriving after transfer C replaced it: the row must still point at C, untouched, and B must not be canceled at Stripe or recorded as an anomaly (Task 1 store test, Task 2 webhook test).
2. A newcomer with status `canceled`/`incomplete`/`incomplete_expired`/`unpaid` for a different id than the active row: no write, no anomaly, no `CancelSubscriptionAtPeriodEnd` call, HTTP 200 (Task 1, Task 2).
3. `CancelSubscriptionAtPeriodEnd` failing on a genuine duplicate: the delivery still fails (Stripe retries), and the retry records no second anomaly row (Task 1 dedupe test, Task 2 retry test).
4. Two deliveries for the same subscription where the second fetch is newer than the first: the row ends in the state of the fetch taken inside the lock (Task 2 hook test).
5. A row written before migration 0044 (`stripe_subscription_created` null) receiving a transfer: the transfer wins (Task 1 test).

---

### Task 1: UpsertStripe outcomes, the created comparison, dead newcomers, anomaly dedupe

**Files:**
- Create: `api/internal/db/migrations/0044_subscription_created_and_anomaly_dedupe.up.sql`
- Create: `api/internal/db/migrations/0044_subscription_created_and_anomaly_dedupe.down.sql`
- Modify: `api/internal/entitlements/store.go:315-327` (recordAnomaly), `:339-420` (StripeUpsert, UpsertResult, UpsertStripe)
- Test: `api/internal/entitlements/store_test.go` (existing tests at `:266`, `:378` adapt to the new result shape; new tests appended)

**Interfaces:**
- Consumes: migration 0020's `entitlements` and `entitlement_anomalies` tables.
- Produces:
  ```go
  type UpsertOutcome string
  const (
      OutcomeWritten       UpsertOutcome = "written"
      OutcomeDuplicate     UpsertOutcome = "duplicate"      // existing active row, different id, live non-transfer newcomer: anomaly recorded, caller cancels newcomer
      OutcomeStaleTransfer UpsertOutcome = "stale_transfer" // transfer newcomer not strictly newer than the existing row's subscription: nothing written
      OutcomeDeadNewcomer  UpsertOutcome = "dead_newcomer"  // newcomer status is not live and the existing active row points elsewhere: nothing written
  )
  type UpsertResult struct {
      Outcome                UpsertOutcome
      ExistingSubscriptionID string // set for every outcome but Written
  }
  // StripeUpsert gains: SubscriptionCreated time.Time (sub.Created; zero when unknown)
  ```
  `Duplicate bool` is removed from `UpsertResult`; `recordAnomaly` becomes idempotent per `(kind, stripe_subscription_id)`.

- [ ] **Step 1: Write the migration**

`0044_subscription_created_and_anomaly_dedupe.up.sql`:
```sql
-- The Stripe subscription's own created timestamp, so a "Take over billing" transfer
-- (spec §2.9 RULING 10) can only replace a strictly older subscription (issue #3 H2).
alter table entitlements add column stripe_subscription_created timestamptz;
-- One anomaly row per (kind, subscription): a webhook retry after a failed Stripe cancel,
-- or the nightly orphan pass, no longer re-records the same subscription (issue #3 L1).
create unique index entitlement_anomalies_kind_sub_unique
  on entitlement_anomalies (kind, stripe_subscription_id);
```
Down: drop the index, drop the column. Before writing the down file, check how the existing paired files handle a dedupe index whose duplicates already exist in production: the up must first `delete` duplicate anomaly rows keeping the lowest id per `(kind, stripe_subscription_id)` (`delete from entitlement_anomalies a using entitlement_anomalies b where a.kind = b.kind and a.stripe_subscription_id = b.stripe_subscription_id and a.id > b.id;`) and then create the index.

- [ ] **Step 2: Write the failing store tests**

Adapt the two existing tests to the new shape: `TestUpsertStripeKeepsAnExistingActiveRowOnADifferentSubscriptionID` asserts `res.Outcome == OutcomeDuplicate`; `TestUpsertStripeTransferOverwritesInsteadOfFlaggingADuplicate` sets `SubscriptionCreated` on both writes (old = now-2h, new = now-1h) and asserts `OutcomeWritten`. Add:

```go
func TestUpsertStripeStaleTransferLeavesTheNewerRowUntouched(t *testing.T)
// seed row: sub_c, active, SubscriptionCreated now-1h, Transfer true
// write: sub_b, active, CancelAtPeriodEnd true, SubscriptionCreated now-2h, Transfer true
// assert: Outcome == OutcomeStaleTransfer, ExistingSubscriptionID == "sub_c",
//   row still points at sub_c with cancel_at_period_end false, anomaly rows == 0, audit rows == 1

func TestUpsertStripeTransferWinsOverARowWithUnknownCreated(t *testing.T)
// seed row: sub_old written with zero SubscriptionCreated (column null)
// write: sub_new, Transfer true, SubscriptionCreated now
// assert: OutcomeWritten, row points at sub_new, stripe_subscription_created is not null

func TestUpsertStripeEqualCreatedIsNotNewer(t *testing.T)
// both created at the same instant → OutcomeStaleTransfer

func TestUpsertStripeDeadNewcomerIsIgnoredWithoutAnAnomaly(t *testing.T)
// table-driven over status in canceled, incomplete, incomplete_expired, unpaid:
// seed active sub_a; write sub_x with that status (Transfer false)
// assert: OutcomeDeadNewcomer, ExistingSubscriptionID == "sub_a", row unchanged, anomaly rows == 0

func TestUpsertStripeStoresTheSubscriptionCreated(t *testing.T)
// first write with SubscriptionCreated → column equals it (UTC, second precision)

func TestRecordAnomalyIsIdempotentPerKindAndSubscription(t *testing.T)
// RecordAnomaly twice with the same kind and id → exactly one row; a different kind for the same id → second row
```

- [ ] **Step 3: Run the tests to verify they fail**

Run: `cd api && TEST_DATABASE_URL=... go test -p 1 ./internal/entitlements/ -run 'UpsertStripe|RecordAnomaly' -v`
Expected: compile failures on `Outcome`/`SubscriptionCreated`, then FAIL.

- [ ] **Step 4: Implement**

In `store.go`:
- `StripeUpsert.SubscriptionCreated time.Time`; `UpsertOutcome` consts and the new `UpsertResult` replace the `Duplicate` field.
- `existingRow` gains `stripeSubscriptionCreated *time.Time`; `existingEntitlement` selects it.
- `UpsertStripe` guard, in this order, when `!isNew && isActiveStatus(existing.status) && existing.stripeSubscriptionID != nil && *existing.stripeSubscriptionID != p.StripeSubscriptionID`:
  1. `!isActiveStatus(p.Status)` → return `UpsertResult{Outcome: OutcomeDeadNewcomer, ExistingSubscriptionID: ...}` with no write, no anomaly.
  2. `p.Transfer`: newer when `existing.stripeSubscriptionCreated == nil || p.SubscriptionCreated.After(*existing.stripeSubscriptionCreated)`; not newer → `OutcomeStaleTransfer`, no write, no anomaly; newer → fall through to the write.
  3. otherwise → `OutcomeDuplicate` and `recordAnomaly` (unchanged).
- The insert/upsert writes `stripe_subscription_created` (nil when `p.SubscriptionCreated.IsZero()`, else UTC).
- `recordAnomaly`: `on conflict (kind, stripe_subscription_id) do nothing`.
- Update the doc comments on `StripeUpsert.Transfer` and `UpsertResult` to describe the three non-written outcomes (keep them short).

- [ ] **Step 5: Run the package tests**

Run: `cd api && TEST_DATABASE_URL=... go test -p 1 ./internal/entitlements/ && go vet ./internal/entitlements/`
Expected: PASS. (The billing package will not compile until Task 2; that is expected at this commit — run only the entitlements package.)

- [ ] **Step 6: Commit**

```bash
git add api/internal/db/migrations/0044_* api/internal/entitlements/store.go api/internal/entitlements/store_test.go
git commit -m "fix(entitlements): a transfer only replaces an older subscription; dead newcomers and repeat anomalies are ignored"
```

---

### Task 2: The webhook acts on outcomes and writes from a fetch taken inside the lock

**Files:**
- Modify: `api/internal/billing/webhook.go:140-173` (event handlers), `:226-301` (upsertFromSubscription, applyUpsert)
- Modify: `api/internal/billing/reconcile.go:46-55` (Reconcile's loop)
- Modify: `api/internal/billing/fake_gateway.go:104-115` (GetSubscription)
- Modify: `api/internal/billing/webhook_test.go:44-80` (fixtures), `:241-404` (existing tests adapt), new tests appended
- Modify: `docs/superpowers/specs/2026-09-21-entitlements-and-payments-design.md` (an "Amendment, 2026-10-24" paragraph after the duplicate-subscription operator runbook in §2.7's amendment block, ~line 730)
- Test: `api/internal/billing/webhook_test.go`, `api/internal/billing/reconcile_test.go` (compile only)

**Interfaces:**
- Consumes: Task 1's `UpsertOutcome`, `UpsertResult`, `StripeUpsert.SubscriptionCreated`.
- Produces:
  ```go
  // upsertSubscription fetches id, picks the subject lock from its metadata, then inside the
  // lock fetches it AGAIN and writes from that second object (issue #3 M3).
  func (s *Service) upsertSubscription(ctx context.Context, subscriptionID, actor string) error
  // FakeGateway gains:
  GetSubscriptionCalls []string                 // ids passed to GetSubscription, in order
  OnGetSubscription    func(id string, call int) // optional; runs after each call, before the return; call counts from 1
  ```

- [ ] **Step 1: Write the failing webhook tests**

Fixtures: `fixtureSubscription` sets `Created: time.Now().Add(-time.Hour).Unix()`; `entitlementsStripeUpsertFor` copies `SubscriptionCreated: time.Unix(sub.Created, 0).UTC()` and `Transfer: sub.Metadata["intent"] == "transfer"`. Existing tests: the transfer test gives `newSub.Created` a later value than `oldSub.Created`.

```go
func TestWebhookStaleTransferEventDoesNotOverwriteTheNewerTakeover(t *testing.T)
// seed row from transfer C (Created now-1h) via UpsertStripe; register B (Created now-2h, intent=transfer,
// CancelAtPeriodEnd true, active) in the gateway; send customer.subscription.updated for B
// assert 200; row still sub_C with cancel_at_period_end false; no CanceledAtPeriodEnd; anomaly rows 0

func TestWebhookDeadNewcomerIsAcknowledgedWithoutCancelOrAnomaly(t *testing.T)
// seed active sub_a; register sub_x status canceled for another officer; send customer.subscription.updated
// assert 200; row sub_a; CanceledAtPeriodEnd empty; anomaly rows 0

func TestWebhookDuplicateCancelFailureRetriesWithoutASecondAnomaly(t *testing.T)
// seed active sub_a; register live sub_b; set gateway.Err only for CancelSubscriptionAtPeriodEnd
//   (add FakeGateway.CancelErr error used by CancelSubscriptionAtPeriodEnd alone); first delivery → 500;
//   clear CancelErr; redeliver the SAME event (build the signed request once, keep its body bytes and
//   Stripe-Signature header, and post a second request from them so the event id matches) → 200;
//   anomaly rows == 1; CanceledAtPeriodEnd == [sub_b]

func TestWebhookWritesFromTheFetchTakenInsideTheLock(t *testing.T)
// register sub_p status past_due; OnGetSubscription: on call 1 flip the registered object to status active;
// send invoice.paid for it; assert row status active and GetSubscriptionCalls == [sub_p, sub_p]

func TestWebhookUpdatedAfterDeletedStaysCanceled(t *testing.T)
// seed active sub_d; register sub_d with status canceled; send customer.subscription.deleted (row canceled, grace set);
// then send customer.subscription.updated for sub_d; assert status still canceled and grace_until unchanged
```
Also assert in the existing concurrency test that each delivery fetched twice (`len(GetSubscriptionCalls) == 4`).

- [ ] **Step 2: Run them to verify they fail**

Run: `cd api && TEST_DATABASE_URL=... go test -p 1 ./internal/billing/ -run 'Webhook' -v`
Expected: compile failures on `Outcome`/hooks, then FAIL.

- [ ] **Step 3: Implement**

`webhook.go`:
- `onCheckoutCompleted`: keep its own first fetch (it needs the customer id); after `SaveCustomerID` call `s.upsertSubscription(ctx, subID, actor)`.
- `onSubscriptionEvent`, `onInvoiceEvent`: call `s.upsertSubscription(ctx, id, actor)` directly (no fetch of their own).
- `upsertSubscription(ctx, id, actor)`: fetch once; decide plan/subject exactly as today (the three "not ours" warnings stay); inside `WithGuildLock`/`WithUserLock` fetch again, verify `plan`, `user_id` and `guild_id` metadata equal the first fetch's (if not: warn `"billing: subscription metadata changed between fetches"` and return nil), build `StripeUpsert` from the second object with `SubscriptionCreated: time.Unix(sub.Created, 0).UTC()`, then `applyUpsert`.
- `applyUpsert` switches on `result.Outcome`: `Written` → nil; `Duplicate` → the existing Error log + `CancelSubscriptionAtPeriodEnd` (error still returned); `StaleTransfer` → `s.logger().Info("billing", "op", "stale_transfer_ignored", "subscription_id", sub.ID, "kept_subscription_id", ...)` and nil; `DeadNewcomer` → `Info` with `"op", "dead_newcomer_ignored"` and nil.
- Keep `upsertFromSubscription(ctx, sub, actor)` only if Reconcile still needs it; otherwise Reconcile calls `upsertSubscription(ctx, id, reconcileActor)` and drops its own `GetSubscription` (one fewer fetch there, two inside). Delete the dead function.

`fake_gateway.go`: `GetSubscriptionCalls`, `OnGetSubscription`, `CancelErr` as specified; `GetSubscription` appends the id, invokes the hook with the 1-based call count, then returns the (possibly mutated) registered object.

Spec amendment (one paragraph): the duplicate guard's order (dead newcomer, transfer by `created`, duplicate), the stale-transfer outcome being silent by design (RULING 10 leaves the old subscription for the guild to cancel; a stale event is that cancellation arriving), the in-lock re-fetch, and the anomaly unique key; cite issue #3 H2/M3/M4/L1.

- [ ] **Step 4: Run the billing packages**

Run: `cd api && TEST_DATABASE_URL=... go test -p 1 -race ./internal/billing/ ./internal/entitlements/ && go vet ./... && gofmt -l ./internal`
Expected: PASS, no vet output, no unformatted files.

- [ ] **Step 5: Commit**

```bash
git add api/internal/billing docs/superpowers/specs/2026-09-21-entitlements-and-payments-design.md
git commit -m "fix(billing): write from a fetch inside the subject lock; act on stale transfers and dead newcomers without canceling"
```
