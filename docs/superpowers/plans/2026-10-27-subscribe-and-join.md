# Subscribe and Join Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** No path from "I want to pay" or "I want my guild on a plan" ends on a dead end. The checkout guild step shows every guild the player belongs to, with their character, rank and verified state, and the one next step for each (pay, claim, confirm, or a copy link to send the person who can). Any account that already has Premium sees "You already have Premium." with its source instead of Stripe or a raw API sentence, and the account Plan panel's headline names the real source (fixing the owner's granted account reading `Not on Premium.`).

**Architecture:**
- **One step rule in Go.** `guilds.SubscribeStep(StepFacts) StepResult`, beside `guilds.Capabilities`, resolves the 13 `step` values first-match in §8 order from the predicates `billing.checkGuildCheckout` and `guilds.Claim` already use. The web never derives a step from a rank.
- **One standing read per guild.** `auth.Service` gains a `Standing GuildStandingSource` (an interface in `auth`, like `Seats SeatFactsSource`). `guilds.Store.GuildStanding` implements it (guilds already imports auth) and fills `realm`, `crest_url`, `faction`, `roster_count`, `claim`, `verify`, `covered` and `subscribe` on each `/v1/me` `guilds[]` entry. Billing facts reach it through a small interface that `billing.Store` satisfies, wired in `cmd/api`.
- **`/v1/me` entitlements** gain `billing.source` and `personal_portal`; the personal checkout 409 gains `fields.source`; the guild home `viewer` gains `verify.has_approver`; the Discord bot's plan links move to the checkout guild step.
- **Web.** `premiumSource(me)` in `plan-line.ts` is the one source rule for the account panel, the checkout "already have it" panel and the `/premium` stored hint. The checkout guild step is rebuilt inside the existing single island (`CheckoutLauncher`). `CopyLinkField` is extracted from `GuildInviteLink`. The claim page, guild page, settings `#billing`, account guild rows and invite joined state gain their hand-off lines.

**Tech Stack:** Go (pgx v5), Astro + Svelte 5, Tailwind, vitest, Playwright + axe, lhci.

**Spec:** `design/specs/2026-10-27-subscribe-and-join.md` — §3.2 states, §4 "you already have it", §5 the checkout guild step, §6 join/verify/claim, §8 API, §9 copy verbatim, §11 performance, §12 acceptance and captures, §13 Decisions (2026-10-27). Read with `design/specs/2026-10-27-launch-pricing.md` (its 2026-10-27b amendment owns the plans page cards, `#billing` and the `Verified members` rule).

## Global Constraints

- **Decisions (§13), applied everywhere:**
  1. No member-started claims; only an officer or guild-master character starts one (`eligibleClaimRank` unchanged). Members get the copy link.
  2. Notifying the guild master is deferred. The anonymous copy link ships. A v2 `POST /v1/guilds/{id}/plan-asks` route is noted for later and is **not** built here.
  3. The hand-off link carries no account id, no token and no sender name.
  4. Billing stays with the claim holder (`Capabilities` unchanged).
  5. No purchase is offered to seat or grant holders; every "you already have it" state shows its source and never a buy button.
  6. `covered` (tier, state, seat counts) goes to any account with a character in the guild; the payer's name only to verified members.
  7. Grant wording is `Given by Forever Sixty.`; the grant note is never shown.
  8. The moderator fallback ships on `ask_leader` and `claimed_ask` cards: `Guild master gone? Ask us on Discord and we will sort the claim by hand.`, `Ask us on Discord` an `INLINE_LINK` to `communityLinks.discordInvite` (`web/src/data/links.json`), new tab, `rel="noopener"`.
  9. The Battle.net rule for claims stays (`ErrNoBattleNetIdentity`, `claim_link_bnet`); verification does not need Battle.net.
  10. The stored hint on `/premium` ships, modelled on `fs.accountReserve` in `web/src/lib/account/layout.ts`.
  11. Capture 1 is taken in both states (the owner as a member; the owner as guild master of OLYMPUS XXVII after an addon save).
- **No change to who may claim, verify or bill. No new tables** (§8).
- **Copy is verbatim from §9.** Every string lives in a copy module whose test reads the spec table through `web/src/test-support/spec-table.ts` (`specCell` / `specString`). Dates `d Mon yyyy` (`18 Nov 2027`) via `utcDayMonthYear`. Guild names render as stored. The figure 25 is always `GUILD_TIER_MEMBER_LIMIT`, never typed. No copy assumes 40-person raids; no text contains `40-person` or `40 players`. No copy ties verification to Battle.net.
- **No API sentence is ever shown** on the checkout page; errors map from the API's error codes to §9.1 lines.
- **One gold fill per page** (§10): the `ready` button when it is the only one, or `Get set up` in Empty A. Two or more `ready` cards are all gold outline. Hand-off cards have no fill.
- **Single island, no `/home` reads** on `/premium/checkout` (§11): crests, faction and counts come from `/v1/me`.
- **CLS** at most 0.02 at 390 and 1440 for the guild step from loading to ready, the "already have it" panel and the account Plan panel; `/premium` CLS unchanged with or without `data-premium`. Skeleton cards reserve 156 px desktop and 260 px phone.
- **Touch and keyboard.** Every control at least 44 px tall at 360. Focus never moves on first render; it moves only after an in-page change (§4.2, §5.3, §5.7). Status lines `role="status"`, refusals `role="alert"`.
- **Layout** (§7). The band is full bleed with its content in the centred 1344 px column at every width up to 2000 and beyond; the body is the centred 880 px column; cards never stretch past 880.
- **API tests.** `-p 1`, each task on a private database: `TEST_DATABASE_URL=postgres://forever:forever@localhost:5434/subjoin_t<N>?sslmode=disable`, created with `psql postgres://forever:forever@localhost:5434/postgres -c 'create database subjoin_t<N>'`. TDD: the failing test first.
- **Migrations.** None in this plan. Only the deployed API ever migrates.
- **File limits.** Go files at or under 800 lines; TS and Svelte at or under 400. Already over: `web/src/lib/account/api.ts` (799), `web/src/components/Account.svelte` (983), `web/src/components/Guild.svelte` (1215), `web/src/components/GuildClaim.svelte` (404). Those gain import and one-line hook edits only; additions go in new files.
- **Web tasks** run the touched vitest files, `npm run check`, `npm run lint` and `npx prettier --check <files>`. PR C runs the e2e and captures; the controller runs `astro build`, the full vitest and Playwright suites and the lhci shard at branch level (memory: shared layout needs the full suite). The worktree needs `npm ci` in `web/` once.
- **Commits.** Conventional types, by path (`git commit -m "<msg>" -- <paths>`), with the trailer the session's attribution reminder gives.

## Review Focus

1. **Overlapping facts in the step rule.** A guild on a plan whose claim is also contested; a claim holder with a checkout already open; a guild-master-rank viewer with no Battle.net link who is also rate-limited; an officer who holds another guild's claim (must be `claim_blocked`, not `start_claim`). First match in §8 order must hold. Pinned in Task 3 (`TestSubscribeStep_FirstMatchOverlaps`).
2. **Names leaking to the wrong viewer.** `claim.holder_name` only to a verified officer of the guild; `covered.billed_by` empty for a grant and for a viewer with no verified character there. Pinned in Task 2 (`TestGuildStanding_HolderNameOnlyForVerifiedOfficer`) and Task 5 (`TestGuildStanding_BilledByHiddenFromUnverified`, `TestGuildStanding_BilledByEmptyForGrant`).
3. **The owner's grant with no Stripe customer.** `billing.source` `grant`, `current_period_end` null, `personal_portal` false; the account panel reads `Premium · granted, no end date` with no billing button, and `Manage billing` never renders for it. Pinned in Task 6 (`TestMe_GrantWithoutStripeRow`) and Task 11 (`AccountPlanPanel.test.ts › "grant open-ended shows no billing button"`).
4. **Untrusted URL inputs.** The hand-off link's `plan` / `interval` parsed against the closed sets of `parseCheckoutParams` (anything else dropped silently, never echoed); `/setup?next=` through `safeNextPath` (no off-site or protocol-relative target); `guild_id` not in `me.guilds` renders the full list with no error. Pinned in Task 14 (`subscribe-step.test.ts › "handoffUrl drops values outside the closed sets"`), Task 20 (`setup-next.test.ts › "rejects off-site and protocol-relative next"`) and Task 18 (`CheckoutGuildStep.test.ts › "unknown guild_id shows the full list"`).
5. **The 409 race and a stale stored hint.** A personal checkout refused anyway must carry `fields.source`, and the launcher must re-read `/v1/me` with `forgetSession()` and render the panel, never `this account already has Premium`. A malformed `fs.premiumSource` must be removed and the signed-out default render. Pinned in Task 7 (`TestPremiumCheckout_ConflictCarriesSource`), Task 12 (`PremiumOwnedPanel.test.ts › "a 409 with fields.source=grant renders the grant panel"`) and Task 9 (`premium-hint.test.ts › "removes a value that is not one of the three words"`).

---

## PR A — the API

### Task 1: Guild identity facts on `/v1/me` `guilds[]`

**Files:**
- Create: `api/internal/auth/guild_standing.go` (view types, the `GuildStandingSource` interface, `attachStanding`), `api/internal/guilds/me_standing.go`, `api/internal/guilds/me_standing_test.go`
- Modify: `api/internal/auth/store.go` (`Guild` struct fields only), `api/internal/auth/handler.go` (`Service.Standing` field and the one-line `attachStanding` call in `me`), `api/cmd/api/main.go` (wire `Standing`)

**Interfaces:**
- Produces (`auth`):
  ```go
  type GuildStanding struct {
      Realm       string
      CrestURL    *string
      Faction     *string // "alliance" | "horde" | nil
      RosterCount *int    // counts.on_roster + counts.in_game_no_account; nil without a snapshot
      Claim       *GuildClaimView
      Verify      *GuildVerifyView
      Covered     *GuildCoveredView
      Subscribe   *GuildSubscribeView
  }
  type GuildStandingSource interface {
      GuildStanding(ctx context.Context, userID int64, guildIDs []int64) (map[int64]GuildStanding, error)
  }
  ```
  `Guild` gains `Realm string json:"realm"`, `CrestURL *string json:"crest_url"`, `Faction *string json:"faction"`, `RosterCount *int json:"roster_count"`, and (filled by Tasks 2 to 5) `Claim *GuildClaimView json:"claim"`, `Verify *GuildVerifyView json:"verify"`, `Covered *GuildCoveredView json:"covered"`, `Subscribe *GuildSubscribeView json:"subscribe"`.
  `func (s *Service) attachStanding(ctx context.Context, userID int64, gs []Guild) []Guild` — nil `Standing` or a failed read logs and returns `gs` unchanged (the older shape; the web draws identity-only cards).
- Produces (`guilds`): `func (s *Store) GuildStanding(ctx context.Context, userID int64, guildIDs []int64) (map[int64]auth.GuildStanding, error)` — one query for realm, faction, crest key (through `CrestURL(s.APIBaseURL, id, key)`) and the roster count from the addon snapshot, for all ids at once (no N+1).

- [ ] **Step 1: Write the failing tests:**
  - `TestGuildStanding_IdentityFacts`: a guild with a crest, faction `horde`, a snapshot of 30 on the roster and 8 in game without an account answers `Realm` (display string), a `CrestURL` with `?v=`, `Faction` `horde`, `RosterCount` 38.
  - `TestGuildStanding_NoSnapshotNoCrest`: `RosterCount` nil, `CrestURL` nil, `Faction` nil when unknown.
  - `TestGuildStanding_OneQueryForManyGuilds`: three ids answer three entries.
  - `TestMe_GuildsCarryIdentity` (auth handler test with a fake `Standing`): the JSON has `realm`, `crest_url`, `faction`, `roster_count`; with `Standing` nil, `/v1/me` answers 200 without them.
- [ ] **Step 2: Run to fail.** `cd api && TEST_DATABASE_URL=…/subjoin_t1… go test -p 1 ./internal/guilds/ ./internal/auth/ -run 'GuildStanding_Identity|NoSnapshotNoCrest|OneQueryForManyGuilds|GuildsCarryIdentity'`. Expected: FAIL.
- [ ] **Step 3: Implement.** Wire `Standing: guildStore` in `cmd/api/main.go`.
- [ ] **Step 4: Run to pass**, then `go build ./... && go test -p 1 ./internal/auth/ ./internal/guilds/`.
- [ ] **Step 5: Commit.** `git commit -m "feat(api): guild identity facts on /v1/me guilds" -- api/internal/auth/guild_standing.go api/internal/auth/store.go api/internal/auth/handler.go api/internal/guilds/me_standing.go api/internal/guilds/me_standing_test.go api/cmd/api/main.go <auth handler test file>`

### Task 2: Claim and verify facts

**Files:**
- Modify: `api/internal/auth/guild_standing.go`, `api/internal/guilds/me_standing.go`, `api/internal/guilds/me_standing_test.go`

**Interfaces:**
- Produces (`auth`):
  ```go
  type GuildClaimView struct {
      State        string  `json:"state"` // "unclaimed" | "pending" | "claimed" | "contested"
      Yours        bool    `json:"yours"`
      PendingYours bool    `json:"pending_yours"`
      HolderName   *string `json:"holder_name"` // holder's or pending claimant's character in the guild; verified officers only
  }
  type GuildVerifyView struct {
      HasApprover bool `json:"has_approver"` // at least one verified officer or leader
  }
  ```
- Produces (`guilds`): `func (s *Store) hasApprover(ctx, guildIDs []int64) (map[int64]bool, error)` — the same verified-officer-or-leader predicate `Capabilities` and `verifiedOfficerOrLeader` read. `claim` reads the same columns `Claim`, `ConfirmClaim` and the contest code read.

- [ ] **Step 1: Write the failing tests:**
  - `TestGuildStanding_ClaimStates`: one case each for `unclaimed`, `pending` (viewer's own: `PendingYours` true), `claimed` (viewer holds it: `Yours` true), `contested`.
  - `TestGuildStanding_HolderNameOnlyForVerifiedOfficer`: a verified officer sees the holder's character name; a verified member, an unverified officer and an unverified member see nil.
  - `TestGuildStanding_HasApprover`: false with zero verified members (the owner's guilds); true with one verified leader; false when the only officer is unverified.
- [ ] **Step 2: Run to fail** on `subjoin_t2`. **Step 3: Implement.** **Step 4: Run to pass** with the full `./internal/guilds/` package.
- [ ] **Step 5: Commit.** `git commit -m "feat(api): claim and verify facts on /v1/me guilds" -- api/internal/auth/guild_standing.go api/internal/guilds/me_standing.go api/internal/guilds/me_standing_test.go`

### Task 3: The one step rule

**Files:**
- Create: `api/internal/guilds/subscribe_step.go`, `api/internal/guilds/subscribe_step_test.go`

**Interfaces:**
- Produces (`guilds`, pure, no I/O):
  ```go
  type Step string
  const (
      StepOnPlan            Step = "on_plan"
      StepCheckoutOpen      Step = "checkout_open"
      StepFrozen            Step = "frozen"
      StepReady             Step = "ready"
      StepClaimedAsk        Step = "claimed_ask"
      StepClaimPendingYours Step = "claim_pending_yours"
      StepConfirmClaim      Step = "confirm_claim"
      StepClaimPendingAsk   Step = "claim_pending_ask"
      StepClaim             Step = "claim"
      StepClaimLinkBnet     Step = "claim_link_bnet"
      StepStartClaim        Step = "start_claim"
      StepClaimBlocked      Step = "claim_blocked"
      StepAskLeader         Step = "ask_leader"
  )
  const (
      BlockedHoldsOther  = "holds_other"
      BlockedRateLimited = "rate_limited"
  )
  type StepFacts struct {
      OnLivePlan          bool       // a whole plan in trial, active, ending or past_due
      CheckoutOpen        bool       // billing.Store.PendingGuildCheckout
      ClaimState          string     // unclaimed | pending | claimed | contested
      ViewerHoldsClaim    bool
      ViewerPendingClaim  bool       // the viewer started the pending claim
      ViewerRank          string     // highest raw rank in the guild: "leader" | "officer" | "member"
      ViewerVerifiedOfficer bool     // the guild_members verified rank Capabilities reads (fix round 1 ruling)
      HasBattleNet        bool
      HoldsOtherClaim     bool
      OtherGuildName      string
      RateLimitedUntil    *time.Time // set when the 30-day claim limit applies
      TrialUsed           bool       // billing.Store.GuildTrialUsed
  }
  type StepResult struct {
      Step           Step
      TrialUsed      *bool      // set only on ready
      AvailableAt    *time.Time // set only on claim_blocked with rate_limited
      Reason         string     // "" | BlockedHoldsOther | BlockedRateLimited
      OtherGuildName string     // set only with holds_other
  }
  func SubscribeStep(f StepFacts) StepResult
  ```
  Resolution, first match in §8 order: `on_plan` (live plan); `checkout_open`; `frozen` (contested); `ready` (claimed, viewer holds it, viewer verified officer or leader — the `CapBilling` facts); `claimed_ask` (claimed, not the holder); `claim_pending_yours` (pending, viewer started it); `confirm_claim` (pending, viewer is officer or leader and not the claimant — the facts `ConfirmClaim` checks); `claim_pending_ask` (pending, anyone else); `claim` (unclaimed, leader, Battle.net linked, not blocked); `claim_link_bnet` (unclaimed, leader, no Battle.net, not blocked); `start_claim` (unclaimed, officer, not blocked); `claim_blocked` (unclaimed, leader or officer, blocked: `holds_other` before `rate_limited`); `ask_leader` (everything else).

- [ ] **Step 1: Write the failing table test** `TestSubscribeStep` with one case per step value (13 cases), each giving the facts that produce it and asserting `Step`, plus:
  - `ready` with `TrialUsed: false` answers `TrialUsed` pointing at false; every other step answers `TrialUsed` nil.
  - `claim_blocked` rate-limited answers `AvailableAt` equal to `RateLimitedUntil` and `Reason` `rate_limited`; holds-other answers `Reason` `holds_other` and `OtherGuildName`.
  - `TestSubscribeStep_OwnerFacts`: the owner in both guilds (`ViewerRank` member, unverified, `ClaimState` unclaimed, no plan) answers `ask_leader`.
  - `TestSubscribeStep_OwnerAsGuildMaster` (capture 1b): `ViewerRank` leader, Battle.net linked, not blocked, unclaimed answers `claim`.
  - `TestSubscribeStep_FirstMatchOverlaps`: on a plan and contested → `on_plan`; holder with a checkout open → `checkout_open`; leader without Battle.net and rate-limited → `claim_link_bnet`; officer holding another claim → `claim_blocked` with `holds_other` (not `start_claim`); lapsed plan with the viewer holding the claim → `ready`; lapsed plan, viewer not the holder → `claimed_ask`.
- [ ] **Step 2: Run to fail.** `go test ./internal/guilds/ -run SubscribeStep`. Expected: FAIL (undefined).
- [ ] **Step 3: Implement.** One small predicate per step, each under 15 lines; `SubscribeStep` walks them in order.
- [ ] **Step 4: Run to pass.**
- [ ] **Step 5: Commit.** `git commit -m "feat(guilds): the one subscribe step rule" -- api/internal/guilds/subscribe_step.go api/internal/guilds/subscribe_step_test.go`

### Task 4: Feed the step from billing and wire `subscribe` into `/v1/me`

**Files:**
- Modify: `api/internal/auth/guild_standing.go`, `api/internal/guilds/me_standing.go`, `api/internal/guilds/me_standing_test.go`, `api/internal/guilds/store.go` (a `Billing StepBillingFacts` field, or the struct that holds `GuildStanding`'s dependencies), `api/cmd/api/main.go`

**Interfaces:**
- Consumes: `SubscribeStep` (Task 3); `billing.Store.PendingGuildCheckout(ctx, guildID) (bool, error)` and `billing.Store.GuildTrialUsed(ctx, guildID) (bool, error)` (both exist); `entitlements.GuildBilling`; the claim rate-limit and holds-other facts `checkClaimRateLimit` and `Claim` read; the account's Battle.net identity.
- Produces:
  - (`guilds`) `type StepBillingFacts interface { PendingGuildCheckout(ctx context.Context, guildID int64) (bool, error); GuildTrialUsed(ctx context.Context, guildID int64) (bool, error) }`.
  - (`auth`) `type GuildSubscribeView struct { Step string json:"step"; TrialUsed *bool json:"trial_used"; AvailableAt *string json:"available_at"; Reason *string json:"reason"; OtherGuildName *string json:"other_guild_name" }`.
  - `checkGuildCheckout` is unchanged: no rule about who may bill moves.

- [ ] **Step 1: Write the failing tests** (DB-backed, through `GuildStanding`):
  - `TestGuildStanding_SubscribeFromBilling`: a pending checkout row gives `checkout_open`; a used trial gives `ready` with `trial_used: true`.
  - `TestGuildStanding_SubscribeRateLimited`: a leader who claimed 10 days ago gives `claim_blocked`, `reason` `rate_limited`, `available_at` the 30-day mark in RFC 3339.
  - `TestGuildStanding_OwnerBothGuilds`: two guilds, member, unverified, unclaimed gives `ask_leader` for both.
  - `TestGuildStanding_NilBillingFactsLeavesSubscribeOut`: no billing source wired, `Subscribe` nil (the older shape).
- [ ] **Step 2: Run to fail** on `subjoin_t4`. **Step 3: Implement**; wire `Billing: billingStore` in `main.go`. **Step 4: Run to pass**, `go build ./...`, full `./internal/guilds/ ./internal/billing/ ./internal/auth/`.
- [ ] **Step 5: Commit.** `git commit -m "feat(api): subscribe step on /v1/me guilds" -- api/internal/auth/guild_standing.go api/internal/guilds/me_standing.go api/internal/guilds/me_standing_test.go api/internal/guilds/store.go api/cmd/api/main.go`

### Task 5: `covered` on `/v1/me` `guilds[]`

**Files:**
- Modify: `api/internal/auth/guild_standing.go`, `api/internal/guilds/me_standing.go`, `api/internal/guilds/me_standing_test.go`

**Interfaces:**
- Consumes: `entitlements.GuildBilling` (tier, status, period end, cancel flag, source), the seat rows (`premium_seats`) and `catalog.SeatsForTier`.
- Produces:
  ```go
  type GuildCoveredView struct {
      Tier        string  `json:"tier"`  // "guild" | "raid_guild"
      State       string  `json:"state"` // "trial" | "active" | "ending" | "past_due" | "lapsed"
      EndsAt      *string `json:"ends_at"`
      BilledBy    string  `json:"billed_by"` // "" for a grant and for a viewer with no verified character
      YouPay      bool    `json:"you_pay"`
      YouHoldSeat bool    `json:"you_hold_seat"`
      SeatsGiven  int     `json:"seats_given"` // slots within the tier's count whose state is held or accepted
      SeatsTotal  int     `json:"seats_total"`
  }
  ```
  `const coveredLapsedShown = 90 * 24 * time.Hour` — `lapsed` shows until 90 days after the end, then `covered` is null. Sent to any account with a character in the guild. Today's officer-only `plan` keeps its shape.

- [ ] **Step 1: Write the failing tests:**
  - `TestGuildStanding_CoveredStates`: trial, active, ending (`cancel_at_period_end`), past_due, lapsed at day 89, null at day 91.
  - `TestGuildStanding_BilledByHiddenFromUnverified`: an unverified member sees `billed_by` ""; a verified member sees the payer's character name.
  - `TestGuildStanding_BilledByEmptyForGrant`: a granted guild answers `billed_by` "".
  - `TestGuildStanding_SeatCounts`: a Raid guild with 2 held and 1 empty answers 2 of 3; the viewer holding one answers `you_hold_seat` true; a Guild tier with a stale slot 2 answers 1 total.
  - `TestGuildStanding_PlanShapeUnchanged`: the officer-only `plan` JSON is byte-identical to before for a verified officer.
- [ ] **Step 2: Run to fail** on `subjoin_t5`. **Step 3: Implement.** **Step 4: Run to pass** with the full `./internal/auth/` and `./internal/guilds/` packages.
- [ ] **Step 5: Commit.** `git commit -m "feat(api): covered plan facts on /v1/me guilds" -- api/internal/auth/guild_standing.go api/internal/guilds/me_standing.go api/internal/guilds/me_standing_test.go`

### Task 6: `billing.source` and `personal_portal` on `/v1/me` entitlements

**Files:**
- Modify: `api/internal/auth/me_billing.go`, `api/internal/auth/me_billing_test.go`, `api/internal/auth/handler.go` (`EntitlementsView.PersonalPortal` field and its one-line fill)
- Modify if needed: `api/internal/entitlements/store.go` (`HasPersonalStripeRow(ctx, userID) (bool, error)`)

**Interfaces:**
- Consumes: `entitlements.Billing.Source` (exists from the launch pricing lane).
- Produces:
  - `BillingView` gains `Source string json:"source"` (`stripe` or `grant`). For a grant, `current_period_end` is the grant's end or null.
  - `EntitlementsView` gains `PersonalPortal bool json:"personal_portal"` — true when the account has a personal entitlement row with `source = 'stripe'` in any status.
  - `func (s *Store) HasPersonalStripeRow(ctx context.Context, userID int64) (bool, error)` in `entitlements`.

- [ ] **Step 1: Write the failing tests:**
  - `TestMe_GrantWithoutStripeRow`: an open-ended grant (`source='grant'`, `current_period_end` null) and no Stripe row answers `billing.source` `grant`, `current_period_end` null, `personal_portal` false.
  - `TestMe_GrantUntilDate`: `current_period_end` is the grant's end.
  - `TestMe_StripeSourceAndPortal`: a live Stripe row answers `source` `stripe`, `personal_portal` true.
  - `TestMe_CanceledStripeRowKeepsPortal`: a canceled Stripe row and a seat answers `personal_portal` true.
- [ ] **Step 2: Run to fail** on `subjoin_t6`. **Step 3: Implement.** **Step 4: Run to pass** with the full `./internal/auth/ ./internal/entitlements/`.
- [ ] **Step 5: Commit.** `git commit -m "feat(api): billing source and personal portal on /v1/me" -- api/internal/auth/me_billing.go api/internal/auth/me_billing_test.go api/internal/auth/handler.go api/internal/entitlements/store.go <entitlements test file>`

### Task 7: The personal checkout 409 carries `fields.source`

**Files:**
- Modify: `api/internal/billing/handler.go` (`alreadyOnPremiumRefusal`), `api/internal/billing/checkout.go` (its one call site)
- Create: `api/internal/billing/premium_refusal_test.go`

**Interfaces:**
- Produces: `func alreadyOnPremiumRefusal(source string) *checkoutRefusal` — `status 409`, `code "conflict"`, `fields{"portal_hint": "1", "source": source}`. The call site passes `current.Source` (`stripe` or `grant`). (The task list names `alreadyOnThePlanRefusal`; that is the guild refusal and stays unchanged. The personal 409 is `alreadyOnPremiumRefusal`.)

- [ ] **Step 1: Write the failing tests:**
  - `TestPremiumCheckout_ConflictCarriesSource`: table over a live Stripe row (`source` `stripe`) and a live grant (`source` `grant`); each answers 409 with `fields.source` equal to it and `fields.portal_hint` `1`.
  - `TestPremiumCheckout_SeatOnlyIsNotRefused`: a seat holder with no personal row is not refused by the API (who may buy is unchanged; the web blocks it per §4.2).
- [ ] **Step 2: Run to fail** on `subjoin_t7`. **Step 3: Implement.** **Step 4: Run to pass** with the full `./internal/billing/`.
- [ ] **Step 5: Commit.** `git commit -m "feat(billing): personal checkout conflict names the Premium source" -- api/internal/billing/handler.go api/internal/billing/checkout.go api/internal/billing/premium_refusal_test.go`

### Task 8: Home `verify.has_approver`, the officer-claim verify test, and the bot's plan links

**Files:**
- Modify: `api/internal/guilds/home.go` (`ViewerView.Verify`), its test; `api/internal/guilds/claim_test.go` (or the confirm-claim test file); `api/internal/discordbot/cmd_readiness.go`, `api/internal/discordbot/cmd_readiness_test.go`

**Interfaces:**
- Consumes: `hasApprover` (Task 2).
- Produces:
  - `ViewerView` gains `Verify auth.GuildVerifyView json:"verify"` (or a guilds-local `VerifyView{HasApprover bool json:"has_approver"}` with the same JSON) on `GET /v1/guilds/{id}/home`.
  - (`discordbot`) `func checkoutGuildURL(base string, guildID int64) string` returning `base + "/premium/checkout?plan=raid_guild&interval=yearly&guild_id=" + strconv.FormatInt(guildID, 10)`.
  - **Note: `cmd_readiness.go`'s `billingAnchor` link moves to the checkout guild step.** The `billingAnchor = "/settings#billing"` constant and the `path+billingAnchor` argument to `readinessRefusal` are replaced by `checkoutGuildURL(s.d.PublicBaseURL, guildID)`, so both `planLink(officerCopy.Subscribe)` and `planLink(officerCopy.MoveToRaidGuild)` point there. **`cmd_readiness_test.go` changes with it:** its `strings.HasSuffix(b.URL, "/settings#billing")` assertion becomes an equality with the checkout URL. The bot's sentences are unchanged.

- [ ] **Step 1: Write the failing tests:**
  - `TestHome_ViewerVerifyHasApprover`: false for a guild with no verified officer, true with one.
  - `TestConfirmClaim_VerifiesClaimant`: after a second officer confirms an officer's pending claim, the claimant's guild membership is verified (§5.4: claiming verifies the claimant). If it fails, fix it in `ConfirmClaim` with the same `setVerifiedForAccount` call `Claim` uses, and say so in the lane report.
  - In `cmd_readiness_test.go`, the lapsed and Guild-tier cases assert `b.URL == "https://<base>/premium/checkout?plan=raid_guild&interval=yearly&guild_id=<id>"`.
- [ ] **Step 2: Run to fail** on `subjoin_t8` (`go test -p 1 ./internal/guilds/ ./internal/discordbot/ -run 'ViewerVerify|VerifiesClaimant|Readiness'`). **Step 3: Implement.** **Step 4: Run to pass** with the full `./internal/guilds/ ./internal/discordbot/`.
- [ ] **Step 5: Commit.** `git commit -m "feat(api): home verify facts; bot plan links open the checkout guild step" -- api/internal/guilds/home.go <home test file> api/internal/guilds/claim_test.go api/internal/discordbot/cmd_readiness.go api/internal/discordbot/cmd_readiness_test.go`

---

## PR B — the web

PR B codes against the PR A contract with stubbed APIs, so web waves do not wait on API waves.

### Task 9: Contract, normalize and the session hints

**Files:**
- Modify: `web/src/lib/billing/contract.ts`, `web/src/lib/billing/normalize.ts`, `web/src/lib/billing/normalize.test.ts`; `web/src/lib/account/api.ts` (imports, type fields, and one-line calls to the hint writer where `fetchMeOnce` resolves and in `forgetSession`)
- Create: `web/src/lib/account/premium-hint.ts`, `web/src/lib/account/premium-hint.test.ts`

**Interfaces:**
- Produces:
  - `contract.ts`: `GuildClaimView`, `GuildVerifyView`, `GuildCoveredView`, `GuildSubscribeView`, `SubscribeStep` (the 13-word union), `ClaimState`, `CoveredState`, mirroring Tasks 1 to 5 exactly; and `export type PremiumSource = 'own' | 'own_then_seat' | 'grant' | 'seat' | 'fallback' | 'none';` (§4.1), which Task 10's `premiumSource` returns. `MeGuild` gains `realm?`, `crest_url?`, `faction?`, `roster_count?`, `claim?`, `verify?`, `covered?`, `subscribe?`. `BillingView` gains `source: 'stripe' | 'grant'`. `EntitlementsView` gains `personal_portal: boolean`. `BillingApiError.fields` exposes `source`.
  - `normalize.ts`: `export function normalizeMeGuild(g: unknown): MeGuild` — an entry missing `subscribe` gets `subscribe: null` ("unknown"); a `step` outside the union becomes null.
  - `premium-hint.ts` (storage through the guarded `storage()` from `web/src/lib/layout-reserve.ts`):
    - `export const PREMIUM_SOURCE_KEY = 'fs.premiumSource'; export const ME_GUILD_COUNT_KEY = 'fs.meGuildCount';`
    - `export type StoredPremiumSource = 'own' | 'seat' | 'grant';`
    - `export function writePremiumHint(source: PremiumSource, doc?: Document): void` — `own_then_seat` stores `own`; `fallback` and `none` remove the key and `data-premium`; sets `document.documentElement.dataset.premium`.
    - `export function readPremiumHint(): StoredPremiumSource | null` — removes and answers null for any other value.
    - `export function clearPremiumHint(doc?: Document): void` (on sign-out).
    - `export function writeMeGuildCount(n: number): void`; `export function readMeGuildCount(): number` — clamped 1 to 4, else 2.

- [ ] **Step 1: Write the failing tests:**
  ```ts
  it('stores own for own_then_seat', () => { writePremiumHint('own_then_seat', document); expect(localStorage.getItem(PREMIUM_SOURCE_KEY)).toBe('own'); expect(document.documentElement.dataset.premium).toBe('own'); });
  it.each(['fallback', 'none'] as const)('%s removes the key and the attribute', (s) => { localStorage.setItem(PREMIUM_SOURCE_KEY, 'grant'); writePremiumHint(s, document); expect(localStorage.getItem(PREMIUM_SOURCE_KEY)).toBeNull(); expect(document.documentElement.dataset.premium).toBeUndefined(); });
  it('removes a value that is not one of the three words', () => { localStorage.setItem(PREMIUM_SOURCE_KEY, 'premium'); expect(readPremiumHint()).toBeNull(); expect(localStorage.getItem(PREMIUM_SOURCE_KEY)).toBeNull(); });
  it('guild count clamps to 1..4, else 2', () => { localStorage.setItem(ME_GUILD_COUNT_KEY, '9'); expect(readMeGuildCount()).toBe(4); localStorage.setItem(ME_GUILD_COUNT_KEY, 'x'); expect(readMeGuildCount()).toBe(2); });
  it('survives a throwing storage', () => { /* stub storage() to throw; read answers null / 2, write is a no-op */ });
  // normalize.test.ts
  it('an older guild entry normalizes to unknown', () => expect(normalizeMeGuild({ id: 1, region: 'us', ruleset: 'forever', name: 'X', verified: false, verified_count: 0, can: [] }).subscribe).toBeNull());
  it('an unknown step becomes null', () => expect(normalizeMeGuild({ ...G, subscribe: { step: 'pay_now' } }).subscribe).toBeNull());
  ```
- [ ] **Step 2: Run to fail.** **Step 3: Implement** (the session-module call to `writePremiumHint(premiumSource(me))` lands in Task 10, once `premiumSource` exists; this task wires `writeMeGuildCount(me.guilds.length)` and `clearPremiumHint` on sign-out). **Step 4: Run to pass**, check, lint, prettier.
- [ ] **Step 5: Commit.** `git commit -m "feat(web): subscribe-and-join contract, normalize and session hints" -- web/src/lib/billing/contract.ts web/src/lib/billing/normalize.ts web/src/lib/billing/normalize.test.ts web/src/lib/account/api.ts web/src/lib/account/premium-hint.ts web/src/lib/account/premium-hint.test.ts`

### Task 10: `premiumSource` and the plan line by source

**Files:**
- Modify: `web/src/lib/account/plan-line.ts`, `web/src/lib/account/plan-line.test.ts`, `web/src/lib/billing/account-plan-copy.ts`, `web/src/lib/billing/account-plan-copy.test.ts`, `web/src/lib/account/api.ts` (one-line `writePremiumHint(premiumSource(me))` where `fetchMeOnce` resolves)

**Interfaces:**
- Produces:
  - `export function premiumSource(me: Me, now: Date): PremiumSource` (`PremiumSource` from `contract.ts`, Task 9; `now` decides `own_then_seat`) — §4.1 in order: Stripe live (`active`, `trialing`, `past_due`) → `own`, or `own_then_seat` when `me.seat.starts_at` is in the future; grant live → `grant`; `me.seat` → `seat`; `entitlements.server_sims === true` → `fallback`; else `none`.
  - `AccountPlanLine` gains `source: PremiumSource` and `billingButton: 'manage' | 'history' | null` (replacing `manageBilling: boolean`): `manage` only for `own` / `own_then_seat`; `history` for every other source when `personal_portal`; else null.
  - `accountPlanLine(me, now, zone?)` answers §4.3's plan line and line 2 per source. A guild-plan member whose Supporter row is On stays `Not on Premium.` (no seat, no row).
  - `accountPlanCopy` gains the §9.2 grant lines (`grantNoEnd`, `grantUntil(date)`, `grantGivenNoEnd`, `grantGivenUntil`), `seatSource(guild)`, `seatGiven(name)`, `ownThenSeat(ends, guild)`, `fallback`, the three `/premium` card lines, the buttons `Manage billing`, `Billing history`, `See your plan`, and the panel title `You already have Premium.`

- [ ] **Step 1: Write the failing tests:**
  ```ts
  it.each([
    ['own renewing', OWN_RENEWING, 'own'], ['own ending', OWN_ENDING, 'own'], ['past due', PAST_DUE, 'own'],
    ['own then seat', OWN_THEN_SEAT, 'own_then_seat'], ['grant open-ended', GRANT_OPEN, 'grant'],
    ['grant until', GRANT_UNTIL, 'grant'], ['seat', SEAT, 'seat'], ['fallback', FALLBACK, 'fallback'], ['free', FREE, 'none'],
  ])('premiumSource %s', (_n, me, want) => expect(premiumSource(me, NOW)).toBe(want));
  it('grant open-ended line and line 2', () => {
    const l = accountPlanLine(GRANT_OPEN, NOW, 'UTC');
    expect(l.line).toBe(specString(SPEC, '### 9.2', 'Grant lines', 0));          // Premium · granted, no end date
    expect(l.line2).toBe(specString(SPEC, '### 9.2', 'Grant lines', 2));         // Given by Forever Sixty. Nothing is billed to you.
    expect(l.billingButton).toBeNull();
  });
  it('grant until a date', () => expect(accountPlanLine(GRANT_UNTIL, NOW, 'UTC').line).toBe(specString(SPEC, '### 9.2', 'Grant lines', 1, { date: '31 Mar 2027' })));
  it('seat with personal_portal shows Billing history', () => expect(accountPlanLine({ ...SEAT, entitlements: { ...SEAT.entitlements, personal_portal: true } }, NOW).billingButton).toBe('history'));
  it('Manage billing only for own and own_then_seat', () => { /* every source; manage iff own|own_then_seat */ });
  it('On rows never read Not on Premium.', () => expect(accountPlanLine(FALLBACK, NOW).line).toBe('Premium'));
  it('§4.3 rows verbatim', () => { /* own renewing `Premium · renews 22 Dec 2026 · $5 a month`, own ending, own_then_seat, seat line + line 2 (`Given by Kraggor. It lasts while the guild's plan does; nothing is billed to you.`), read from the §4.3 table cells */ });
  it('/premium card lines', () => expect([accountPlanCopy.premiumCard.own, accountPlanCopy.premiumCard.seat, accountPlanCopy.premiumCard.grant]).toEqual(specCell(SPEC, '### 9.2', '`/premium` card line')));
  ```
  §4.3 is a four-column table; read its cells with a local `tableRow(spec, '### 4.3', '`seat`')` helper in the test that returns all backticked segments of the row.
- [ ] **Step 2: Run to fail.** **Step 3: Implement.** Update `AccountPremiumBlock.svelte` and any `manageBilling` reader to `billingButton === 'manage'` (import-level edit only; Task 11 renders the button rule). **Step 4: Run to pass**, check, lint, prettier.
- [ ] **Step 5: Commit.** `git commit -m "fix(web): plan line names its Premium source; grants no longer read Not on Premium" -- web/src/lib/account/plan-line.ts web/src/lib/account/plan-line.test.ts web/src/lib/billing/account-plan-copy.ts web/src/lib/billing/account-plan-copy.test.ts web/src/lib/account/api.ts web/src/components/account/AccountPremiumBlock.svelte`

### Task 11: Account Plan panel by source

**Files:**
- Create: `web/src/components/account/AccountPlanHeadline.svelte`, `web/src/components/account/AccountPlanPanel.test.ts`
- Modify: `web/src/components/account/AccountPremiumBlock.svelte` (render the headline component), `web/src/components/Account.svelte` (one-line hook if the headline lives there today)

**Interfaces:**
- Consumes: `accountPlanLine` (Task 10), `openPortal` (`web/src/lib/billing/api.ts`), `SECONDARY_BUTTON_FIXED`.
- Produces: `<AccountPlanHeadline me now />` — plan line 14 px, `pill-site` `Founder` 8 px after it when founder, line 2 13 px muted, `pastDueBanner` as `role="alert"` for past due, then the button: `Manage billing` (portal), `Billing history` (`SECONDARY_BUTTON_FIXED`, the portal), or none. `See plans` inline with `Not on Premium.`

- [ ] **Step 1: Write the failing tests:**
  ```ts
  it('grant open-ended shows no billing button', () => { render(AccountPlanHeadline, { props: { me: GRANT_OPEN, now: NOW } }); expect(screen.getByText('Premium · granted, no end date')).toBeTruthy(); expect(screen.queryByRole('button', { name: /billing/i })).toBeNull(); });
  it('seat with personal_portal shows Billing history', () => { /* button name 'Billing history' */ });
  it('own shows Manage billing and the founder pill', () => { /* … */ });
  it('past due shows the alert', () => { /* role=alert */ });
  it('free shows See plans', () => { /* link to /premium */ });
  ```
- [ ] **Step 2: Run to fail.** **Step 3: Implement.** **Step 4: Run to pass**, check, lint, prettier.
- [ ] **Step 5: Commit.** `git commit -m "feat(web): account plan headline and billing button by source" -- web/src/components/account/AccountPlanHeadline.svelte web/src/components/account/AccountPlanPanel.test.ts web/src/components/account/AccountPremiumBlock.svelte web/src/components/Account.svelte`

### Task 12: Checkout `plan=premium` — "You already have Premium."

**Files:**
- Create: `web/src/components/checkout/PremiumOwnedPanel.svelte`, `web/src/components/checkout/PremiumOwnedPanel.test.ts`
- Modify: `web/src/components/CheckoutLauncher.svelte`

**Interfaces:**
- Consumes: `premiumSource`, `accountPlanLine` (Task 10), `fetchMeOnce`, `forgetSession`, `startCheckout`, `openPortal`, `BillingApiError` (with `fields.source`).
- Produces:
  - Launcher flow for `plan=premium`: read `/v1/me` first; when `premiumSource(me, now) !== 'none'` render the panel and never call `POST /v1/billing/checkout`. On a 409 whose `fields.source` is set, call `forgetSession()`, re-read `/v1/me`, render the panel. No API sentence ever renders.
  - `<PremiumOwnedPanel me now />` (§4.2): `card`, max 640 px, centred, padding 20, gap 8; title `You already have Premium.` 15 px `--strong` (tabindex −1, focused when it replaces the redirecting message); plan line 14 px and line 2 13 px muted from `accountPlanLine`; `Founder` pill; past-due alert; button 12 px below line 2: `Manage billing` for `own` / `own_then_seat`, `See your plan` (link to `/account#plan`) for `seat`, `grant`, `fallback`. Band h1 `Premium`. Phone: full width inside 18 px gutters, button full width.

- [ ] **Step 1: Write the failing tests:**
  ```ts
  it.each(SOURCES)('%s: title, line, line 2, button; checkout never called', async (name, me, button) => {
    const start = vi.fn(); render(CheckoutLauncher, { props: { search: '?plan=premium&interval=yearly', deps: { fetchMe: async () => me, startCheckout: start } } });
    expect(await screen.findByText('You already have Premium.')).toBeTruthy();
    expect(screen.getByRole(button.role, { name: button.label })).toBeTruthy();
    expect(start).not.toHaveBeenCalled();
  });
  it('a 409 with fields.source=grant renders the grant panel', async () => {
    const forget = vi.fn(); let calls = 0;
    const fetchMe = async () => (calls++ === 0 ? FREE : GRANT_OPEN);
    const start = vi.fn().mockRejectedValue(new BillingApiError(409, 'conflict', 'this account already has Premium', { source: 'grant' }));
    render(CheckoutLauncher, { props: { search: '?plan=premium&interval=yearly', deps: { fetchMe, startCheckout: start, forgetSession: forget } } });
    expect(await screen.findByText('Premium · granted, no end date')).toBeTruthy();
    expect(forget).toHaveBeenCalled();
    expect(document.body.textContent).not.toContain('this account already has Premium');
  });
  ```
  `SOURCES` covers own renewing, own ending, past due, seat, grant open-ended and grant until (§12 case 14). Build the `BillingApiError` with its real constructor (`web/src/lib/billing/api.ts`); the shape above shows only status, code and `fields.source`. If the launcher takes no `deps` today, add an optional `deps` prop with production defaults (dependency injection, no globals).
- [ ] **Step 2: Run to fail.** **Step 3: Implement.** **Step 4: Run to pass**, check, lint, prettier.
- [ ] **Step 5: Commit.** `git commit -m "fix(web): checkout shows You already have Premium instead of the API refusal" -- web/src/components/checkout/PremiumOwnedPanel.svelte web/src/components/checkout/PremiumOwnedPanel.test.ts web/src/components/CheckoutLauncher.svelte`

### Task 13: Extract `CopyLinkField`

**Files:**
- Create: `web/src/components/ui/CopyLinkField.svelte`, `web/src/components/ui/CopyLinkField.test.ts`
- Modify: `web/src/components/guild/GuildInviteLink.svelte`

**Interfaces:**
- Produces: `<CopyLinkField url: string label: string shareTitle: string shareText: string testid: string />` — the read-only 44 px field (`border-line-warm-strong bg-raised rounded-control`, 13 px mono, truncating, `readonly`, select-all on focus, sr-only `label`), `Copy link` reading `Copied` for 2 s, `Share` where `navigator.share` exists, and the failure line `Could not copy. Select the link and copy it.` Buttons `SECONDARY_BUTTON_FIXED px-4`. Visuals unchanged from `GuildInviteLink`.

- [ ] **Step 1: Write the failing tests:** copy writes `url` to the clipboard and the button reads `Copied`, back to `Copy link` after 2 s (fake timers); a rejected clipboard shows the failure line; `Share` renders only with `navigator.share`; focus selects the whole value; `GuildInviteLink`'s existing tests still pass unchanged.
- [ ] **Step 2: Run to fail.** **Step 3: Extract**, and make `GuildInviteLink` render `CopyLinkField` with its own label and share copy. **Step 4: Run to pass** with `GuildInviteLink`'s tests, check, lint, prettier. (Capture 17, the pixel diff, runs in Task 27.)
- [ ] **Step 5: Commit.** `git commit -m "refactor(web): extract CopyLinkField from GuildInviteLink" -- web/src/components/ui/CopyLinkField.svelte web/src/components/ui/CopyLinkField.test.ts web/src/components/guild/GuildInviteLink.svelte`

### Task 14: Checkout guild-step logic and copy

**Files:**
- Create: `web/src/lib/billing/subscribe-step.ts`, `web/src/lib/billing/subscribe-step.test.ts`, `web/src/lib/billing/checkout-guild-copy.ts`, `web/src/lib/billing/checkout-guild-copy.test.ts`
- Modify: `web/src/lib/billing/checkout-params.ts` (export the closed-set parsers `parsePlan`, `parseInterval` it already applies), `web/src/lib/billing/copy.ts` (remove the chooser strings `whichGuild`, `noBillableGuild`, `findYourGuild`, `guildsDidNotLoad` once Task 16 stops reading them)

**Interfaces:**
- Consumes: `GUILD_TIER_MEMBER_LIMIT`, `priceFor`, `formatPrice`, `formatAmount` (`plans.ts`); `MeGuild`, `SubscribeStep` (Task 9); `utcDayMonthYear`.
- Produces:
  - `checkoutGuildCopy` — every §9.1 row, with a function per `{…}` slot: `band`, `sizeLine.guild`, `sizeLine.raidGuild`, `lead.ready`, `lead.allUnclaimed(n: number)`, `lead.someClaimed`, `lead.allClaimed`, `rank`, `verifyMark`, `more(n)`, `facts.*`, one entry per step (`ready.trialUnused(guild)` …), `confirm.*`, `done.*`, `verify.approver`, `verify.noApprover(guild)`, `verify.short`, `otherCharacter(guild)`, `moderatorFallback`, `copyField.*`, `share.*`, `busy.*`, `claimErrors.*`, `checkoutErrors.*`, `portalError`, `loading`, `loadError`, `emptyA`, `emptyB`, `foot.*`, `invalidLink`. `sizeLine.guild` interpolates `GUILD_TIER_MEMBER_LIMIT`.
  - `subscribe-step.ts`:
    - `export const STEP_ORDER: readonly SubscribeStep[]` — §5.1's card order: `ready`, `claim`, `confirm_claim`, `start_claim`, `claim_link_bnet`, `claimed_ask`, `ask_leader`, `claim_pending_yours`, `claim_pending_ask`, `checkout_open`, `frozen`, `claim_blocked`, `on_plan`.
    - `export function orderGuildCards(guilds: readonly MeGuild[], firstId?: number): MeGuild[]` — by `STEP_ORDER`, ties in `me.guilds` order; `firstId` (a `guild_id` that is in the list) first.
    - `export function leadSentence(guilds: readonly MeGuild[]): string` — §9.1's lead rows (one/two/three-plus unclaimed forms; some claimed; all claimed or on a plan; any `ready`).
    - `export function fitsPlan(verifiedCount: number): PlanKey` — `raid_guild` above `GUILD_TIER_MEMBER_LIMIT`, else `guild`.
    - `export function switchNote(plan: PlanKey, verifiedCount: number): { text: string; link: string; to: PlanKey } | null`.
    - `export function verifySubLine(g: MeGuild, viewerVerified: boolean): string | null` — §5.4.
    - `export function handoffUrl(g: Pick<MeGuild, 'region' | 'ruleset'> & { slug: string }, step: SubscribeStep, plan: string, interval: string, origin = 'https://foreversixty.gg'): string` — the claim URL for `ask_leader`, `claim_pending_yours`, `claim_pending_ask`; the settings `#billing` URL for `claimed_ask`; `plan` / `interval` kept only when inside the closed sets.
    - `export function copyFieldLabel(g: MeGuild, step: SubscribeStep, viewerIsOfficer: boolean): string`.

- [ ] **Step 1: Write the failing tests:**
  ```ts
  // checkout-guild-copy.test.ts — each §9.1 row through specCell/specString
  it('Size line, Guild', () => expect(copy.sizeLine.guild).toBe(specString(SPEC, '### 9.1', 'Size line, Guild')));
  it('lead, two unclaimed', () => expect(copy.lead.allUnclaimed(2)).toBe(specCell(SPEC, '### 9.1', 'Lead, none `ready`, every guild unclaimed')[1]));
  it('ask_leader', () => expect(copy.askLeader('OLYMPUS XXVII')).toBe(specString(SPEC, '### 9.1', '`ask_leader`', 0, { Guild: 'OLYMPUS XXVII' })));
  it('moderator fallback', () => expect(copy.moderatorFallback).toBe(specCell(SPEC, '### 9.1', '`ask_leader` and `claimed_ask`, moderator fallback')[0]));
  it('no verification copy names Battle.net', () => expect([copy.verify.approver, copy.verify.noApprover('X'), copy.verify.short, copy.emptyA.body, copy.emptyB.body].join(' ')).not.toMatch(/Battle\.net/));
  it('no 40-person', () => expect(JSON.stringify(allStrings(copy))).not.toMatch(/40-person|40 players/));
  // …one test per remaining §9.1 row
  // subscribe-step.test.ts
  it('handoffUrl drops values outside the closed sets', () => {
    expect(handoffUrl(G, 'ask_leader', 'premium', 'weekly')).toBe('https://foreversixty.gg/guild/us/forever/olympus-xxvii/claim');
    expect(handoffUrl(G, 'ask_leader', 'guild', 'yearly')).toBe('https://foreversixty.gg/guild/us/forever/olympus-xxvii/claim?plan=guild&interval=yearly');
    expect(handoffUrl(G, 'claimed_ask', 'raid_guild', 'monthly')).toBe('https://foreversixty.gg/guild/us/forever/olympus-xxvii/settings?plan=raid_guild&interval=monthly#billing');
  });
  it('handoffUrl carries no account id, token or name', () => expect(handoffUrl(G, 'ask_leader', 'guild', 'yearly')).not.toMatch(/user|token|from|by=/));
  it('orders ready first, on_plan last, ties by me.guilds', () => { /* … */ });
  it('guild_id in the list goes first', () => { /* … */ });
  it('fitsPlan at the limit', () => { expect(fitsPlan(GUILD_TIER_MEMBER_LIMIT)).toBe('guild'); expect(fitsPlan(GUILD_TIER_MEMBER_LIMIT + 1)).toBe('raid_guild'); });
  it('switch notes', () => { expect(switchNote('guild', 28)?.link).toBe('Switch to Raid guild'); expect(switchNote('raid_guild', 12)?.link).toBe('Switch to Guild'); expect(switchNote('guild', 12)).toBeNull(); });
  it('verify sub-line by approver and step', () => { /* approver → approver line; no approver + claimed_ask → noApprover; ask_leader → short; claim/start_claim/confirm_claim → null; verified → null */ });
  it('lead: owner fixture reads the two-guild form', () => expect(leadSentence(OWNER.guilds)).toBe(copy.lead.allUnclaimed(2)));
  ```
  Where the guild slug comes from: `MeGuild` carries `name`; use the same slug helper the guild links already use (`rg "slug" web/src/lib/guild`), never a new one.
- [ ] **Step 2: Run to fail.** **Step 3: Implement.** Each file within 400 lines (split the copy into `checkout-guild-copy.ts` and `checkout-guild-errors.ts` if needed). **Step 4: Run to pass**, check, lint, prettier.
- [ ] **Step 5: Commit.** `git commit -m "feat(web): checkout guild step rules and copy" -- web/src/lib/billing/subscribe-step.ts web/src/lib/billing/subscribe-step.test.ts web/src/lib/billing/checkout-guild-copy.ts web/src/lib/billing/checkout-guild-copy.test.ts web/src/lib/billing/checkout-params.ts`

### Task 15: The checkout page frame and band

**Files:**
- Create: `web/src/components/checkout/CheckoutBand.svelte`, `web/src/components/checkout/CheckoutBand.test.ts`
- Modify: `web/src/pages/premium/checkout.astro` (drop `measure-form` for the guild step; reserve the band height), `web/src/components/CheckoutLauncher.svelte`

**Interfaces:**
- Consumes: `checkoutGuildCopy.band`, `sizeLine`, `PLANS`, `priceFor`, `founderOpen`, `parseCheckoutParams`.
- Produces: `<CheckoutBand plan interval founder onSwitch? />`, rendered at once from the URL, before `/v1/me` answers (§5.1): full bleed, the `/premium` band markup, padding 28/24 desktop and 12/12 phone, content in the centred 1344 column. `CHECKOUT` label in `--gold`, hidden below 768. h1 Cinzel 700 22 px `Subscribe a guild`. Plan summary right-aligned on the h1's row from 1024, its own row below: plan name Cinzel 700 15 px; 8 px; price JetBrains Mono 500 20 px tabular then ` a year`/` a month` 14 px muted; `pill-site` `Founder` 8 px after while founder; size line 13 px; line 3 `First 7 days free` (`--strong` 600) when the plan carries a trial, ` · `, underlined `Change plan` to `/premium`. `onSwitch(plan)` rewrites `plan` in the URL with `history.replaceState`; no reload.

- [ ] **Step 1: Write the failing tests:** Guild yearly founder renders `Guild`, `$120`, ` a year`, `Founder`, the Guild size line, `First 7 days free`, `Change plan` → `/premium`; Raid guild renders the Raid guild size line; `CHECKOUT` carries `hidden md:block`; switching calls `replaceState` with `plan=raid_guild` and re-renders the summary without navigation; no text matches `/40-person|40 players/`.
- [ ] **Step 2: Run to fail.** **Step 3: Implement.** **Step 4: Run to pass**, check, lint, prettier.
- [ ] **Step 5: Commit.** `git commit -m "feat(web): checkout guild step band and plan summary" -- web/src/components/checkout/CheckoutBand.svelte web/src/components/checkout/CheckoutBand.test.ts web/src/pages/premium/checkout.astro web/src/components/CheckoutLauncher.svelte`

### Task 16: The guild card, skeletons, load error and empty states

**Files:**
- Create: `web/src/components/checkout/CheckoutGuildStep.svelte` (list, lead, foot, loading, error, empty), `web/src/components/checkout/GuildStepCard.svelte` (identity, you row, facts), `web/src/components/checkout/GuildStepBlock.svelte` (the step block for hand-off and read-only steps), `web/src/components/checkout/CheckoutGuildStep.test.ts`, `web/src/components/checkout/GuildStepCard.test.ts`
- Delete: `web/src/components/checkout/CheckoutGuildChooser.svelte` and its test
- Modify: `web/src/components/CheckoutLauncher.svelte`

**Interfaces:**
- Consumes: Task 14's functions and copy; `CopyLinkField` (Task 13); `ClassCrestRing`, `FactionMark`, `Skeleton`, `LoadError`, `EmptyState`, `INLINE_LINK`; `readMeGuildCount` (Task 9); `communityLinks.discordInvite`.
- Produces:
  - `GuildStepCard` (§5.2): `li` `card` recipe, padding 20 / 16, `id="guild-{id}"`, `scroll-margin-top: 96px`, `aria-labelledby` its name; from 768 grid `minmax(0,1fr) 320px`, gap 24. Identity row: 40 px crest circle with a 2 px `--gold` ring (`guildMarkSrc`, 40 x 40 `object-cover`), else `FactionMark` 24 px, else empty ring; name Cinzel 700 16 px truncating; realm and region 13 px muted. You row: up to two of the viewer's characters in the guild, main first, then `+{n} more`; each `ClassCrestRing` 20 px, name in class colour, rank word, verify mark (12 px `--kill` tick + `Verified`, or `Not verified` muted). Facts 13 px muted, numbers mono: `{n} verified` / `No one verified yet`, `{m} on the roster` (left out when null), the plan word.
  - `GuildStepBlock` for `ask_leader`, `claimed_ask`, `claim_pending_yours`, `claim_pending_ask` (sentence + `CopyLinkField` with `handoffUrl`), `on_plan` read-only lines + `Open the guild`, `checkout_open` (`Check again`), `frozen` and `claim_blocked` (`Open the guild`); the verify sub-line; the other-character line under `ask_leader` (`Get set up` to `/setup?next=<this URL>#guild-{id}`); the moderator fallback last on `ask_leader` and `claimed_ask`. Phone: one column, the block after a 1 px `--line-soft` rule with 14 px above and below; controls full width.
  - `CheckoutGuildStep`: the lead (15 px, 14 px below 768, max 720, 16 px below), the `ul` (gap 12) in `orderGuildCards` order, the foot (24 px below, 13 px muted: missing-guild line + `Get set up`; the personal Premium line + `Get Premium` only when `premiumSource` is `none`). Loading: lead slot reserves 21 px (42 px below 768); `Skeleton` cards 156 / 260 px, count `readMeGuildCount()`; sr-only `Loading your guilds`; `aria-busy="true"`; rows replace the skeleton once `/v1/me` answered and every crest decoded or failed, capped at 1000 ms. Error: `LoadError` `Your guilds did not load.` + `Try again`; focus to the lead after a successful retry. Empty A (no characters) and Empty B (none in a guild) per §5.7, `Get set up` the page's one fill, no foot missing-guild line in Empty A. An older `/v1/me` (subscribe null) draws identity rows and `Open the guild` only.

- [ ] **Step 1: Write the failing tests:**
  ```ts
  it('owner fixture: two ask_leader cards, OLYMPUS XXVII first', async () => {
    render(CheckoutGuildStep, { props: { me: OWNER, plan: 'guild', interval: 'yearly' } });
    const cards = await screen.findAllByRole('listitem');
    expect(cards.map((c) => within(c).getByRole('heading').textContent)).toEqual(['OLYMPUS XXVII', 'Employed Millennials']);
    const first = within(cards[0]);
    expect(first.getByText('Member')).toBeTruthy(); expect(first.getByText('Not verified')).toBeTruthy();
    expect(first.getByText(/No one verified yet/)).toBeTruthy(); expect(first.getByText(/No plan/)).toBeTruthy();
    expect((first.getByRole('textbox') as HTMLInputElement).value).toBe('https://foreversixty.gg/guild/us/forever/olympus-xxvii/claim?plan=guild&interval=yearly');
    expect(first.getByText(copy.verify.short)).toBeTruthy();
    expect(first.getByRole('link', { name: 'Ask us on Discord' }).getAttribute('href')).toBe(links.discordInvite);
    expect(first.getByRole('link', { name: 'Ask us on Discord' }).getAttribute('rel')).toContain('noopener');
    expect(screen.getByText(copy.lead.allUnclaimed(2))).toBeTruthy();
    expect(screen.queryByText(/Get Premium/)).toBeNull(); // grant
  });
  it('claimed_ask names the holder to an officer, not to a member', () => { /* … */ });
  it('older /v1/me draws identity rows and Open the guild only', () => { /* subscribe null */ });
  it('skeleton count from fs.meGuildCount, aria-busy', () => { /* … */ });
  it('Empty A and Empty B: Get set up with next back here, no Battle.net', () => { /* … */ });
  it('load error then retry moves focus to the lead', async () => { /* … */ });
  ```
  Use the Task 24 fixtures if they have landed; otherwise define local fixtures in the test and move them in Task 24.
- [ ] **Step 2: Run to fail.** **Step 3: Implement.** **Step 4: Run to pass**, check, lint, prettier.
- [ ] **Step 5: Commit.** `git commit -m "feat(web): checkout guild cards with the next step for every guild" -- web/src/components/checkout/CheckoutGuildStep.svelte web/src/components/checkout/GuildStepCard.svelte web/src/components/checkout/GuildStepBlock.svelte web/src/components/checkout/CheckoutGuildStep.test.ts web/src/components/checkout/GuildStepCard.test.ts web/src/components/checkout/CheckoutGuildChooser.svelte web/src/components/checkout/CheckoutGuildChooser.test.ts web/src/components/CheckoutLauncher.svelte web/src/lib/billing/copy.ts`

### Task 17: In-place claim, start and confirm

**Files:**
- Create: `web/src/components/checkout/ClaimConfirmBlock.svelte`, `web/src/components/checkout/ClaimConfirmBlock.test.ts`
- Modify: `web/src/components/checkout/GuildStepBlock.svelte`, `web/src/components/checkout/CheckoutGuildStep.svelte`

**Interfaces:**
- Consumes: `claimGuild(guildId)`, `confirmClaim(guildId)` (`web/src/lib/guild/api.ts`); `forgetSession`, `fetchMeOnce`; `checkoutGuildCopy.{claim, startClaim, confirmClaim, confirm, done, busy, claimErrors}`.
- Produces: openers for `claim` (`Claim {Guild}`), `start_claim` (`Start a claim`), `confirm_claim` (`Confirm the claim`) — `SECONDARY_BUTTON_FIXED border-gold text-gold w-full justify-center`. `ClaimConfirmBlock` (the `GuildClaim.svelte` recipe, `border-line-soft border p-4`, gap 12): the confirm sentence, the commit `PRIMARY_BUTTON_FIXED` (`Yes, claim it` / `Yes, start the claim` / `Yes, confirm it`; busy `Claiming…` / `Starting…` / `Confirming…` with `aria-busy`), and `Cancel` (44 px underlined text button). Focus moves to `Cancel` on open; Escape closes and returns focus to the opener; one confirm open at a time (other openers disabled). On success: `forgetSession()`, re-read `/v1/me`, the card re-renders in its new step, a 13 px muted `role="status"` line states the outcome, focus to the card's new action or the status line. On refusal: the §9.1 claim-error line mapped from the API error code (`role="alert"`, focus on it), block stays open; already-claimed and pending refusals re-read.
  - `export function claimErrorLine(code: string, guild: string): string` (in `checkout-guild-copy.ts`), mapping the codes the claim routes answer for `ErrAlreadyClaimed`, `ErrClaimPending`, `ErrNotEligible`, `ErrNoBattleNetIdentity`, `ErrClaimRateLimited`, `ErrAlreadyClaimsAnotherGuild`, `ErrSameAccount`; anything else `That did not work; try again.`

- [ ] **Step 1: Write the failing tests:**
  ```ts
  it('claim: confirm opens with focus on Cancel; Escape closes and returns focus', async () => { /* … */ });
  it('Yes, claim it posts, re-reads, card becomes ready with status line', async () => {
    /* claimGuild resolves; second fetchMe returns the guild with step ready; expect status 'Claimed. You hold OLYMPUS XXVII.' and focus on 'Start the 7-day trial' */
  });
  it('start_claim then pending: card becomes claim_pending_yours with a copy field', async () => { /* … */ });
  it.each(['already_claimed', 'not_eligible', 'rate_limited', 'holds_other'])('refusal %s shows its §9.1 line, never the API sentence', async (code) => { /* … */ });
  it('only one confirm open at a time', async () => { /* … */ });
  ```
  Read the claim routes' actual error codes from `api/internal/guilds` (the handler that maps the `Err*` values) before writing the table; do not invent codes.
- [ ] **Step 2: Run to fail.** **Step 3: Implement.** **Step 4: Run to pass**, check, lint, prettier.
- [ ] **Step 5: Commit.** `git commit -m "feat(web): claim, start and confirm in place on the checkout card" -- web/src/components/checkout/ClaimConfirmBlock.svelte web/src/components/checkout/ClaimConfirmBlock.test.ts web/src/components/checkout/GuildStepBlock.svelte web/src/components/checkout/CheckoutGuildStep.svelte web/src/lib/billing/checkout-guild-copy.ts web/src/lib/billing/checkout-guild-copy.test.ts`

### Task 18: `ready` and `on_plan` actions, checkout errors and `guild_id`

**Files:**
- Modify: `web/src/components/checkout/GuildStepBlock.svelte`, `web/src/components/checkout/CheckoutGuildStep.svelte`, `web/src/components/CheckoutLauncher.svelte`, `web/src/lib/billing/checkout-params.ts` (allow `guild_id` on the guild step without auto-launch), `web/src/components/checkout/CheckoutGuildStep.test.ts`

**Interfaces:**
- Consumes: `startCheckout({plan, interval, guildId})`, `openPortal({guildId})`, `switchNote` (Task 14), `CheckoutBand.onSwitch` (Task 15).
- Produces:
  - `ready`: `PRIMARY_BUTTON_FIXED w-full justify-center` labelled `Start the 7-day trial`, `Subscribe · {price}` or `Subscribe again · {price}`; when two or more cards are `ready`, every `ready` button is `SECONDARY_BUTTON_FIXED border-gold text-gold`. Busy `Redirecting to secure checkout…`, other card buttons disabled. The switch note above the button (`Switch to Raid guild` / `Switch to Guild`) rewrites the URL and updates band and button in place.
  - `on_plan`: you pay, tier current or lower → `Manage billing` (`POST /v1/billing/portal` with `guild_id`; error `Billing did not open. Try again.`); you pay on Guild with Raid guild chosen → `Move to Raid guild` (gold outline) linking to the guild's settings with `upgrade=1#billing`; not paying → seat lines + `Open the guild`.
  - Checkout errors: 503 `Purchases are not open yet. Check back soon — nothing was charged.`; 409 on a guild `{Guild} is already on a plan.` + re-read; 403 re-reads and shows the new step; else `That did not work; try again.` + `Try again`.
  - `guild_id`: a `ready` guild launches Stripe directly (today); a non-ready guild renders the step with that card first and `scrollIntoView`; an id not in `me.guilds` renders the full list with no error.

- [ ] **Step 1: Write the failing tests:**
  ```ts
  it('one ready guild: the only filled button posts plan, interval, guild id', async () => { /* … */ });
  it('two ready guilds: both gold outline, no fill', async () => { /* querySelectorAll('.bg-gold') length 0 */ });
  it('switch note rewrites the plan without reload', async () => { /* replaceState spy; band reads Raid guild */ });
  it('on_plan: Manage billing; Move to Raid guild links to settings?upgrade=1#billing; seat held; seats free; seats full', () => { /* … */ });
  it('checkout_open and frozen show no buy button', () => { /* … */ });
  it('guild_id of a non-ready guild goes first and scrolls into view', () => { /* scrollIntoView spy */ });
  it('guild_id of a ready guild launches Stripe', () => { /* startCheckout called on mount */ });
  it('unknown guild_id shows the full list', () => { /* no error text, all cards */ });
  it('403 re-reads and shows the new step, never the API sentence', async () => { /* … */ });
  ```
- [ ] **Step 2: Run to fail.** **Step 3: Implement.** **Step 4: Run to pass**, check, lint, prettier.
- [ ] **Step 5: Commit.** `git commit -m "feat(web): pay, manage and upgrade from the checkout card; guild_id arrivals" -- web/src/components/checkout/GuildStepBlock.svelte web/src/components/checkout/CheckoutGuildStep.svelte web/src/components/CheckoutLauncher.svelte web/src/lib/billing/checkout-params.ts web/src/components/checkout/CheckoutGuildStep.test.ts <checkout-params test>`

### Task 19: `/premium` Premium card "you have it" state

**Files:**
- Modify: `web/src/pages/premium.astro` (the third head `is:inline` script), `web/src/components/premium/PlanCard.astro` (both versions of two slots), `web/src/pages/_premium.test.ts` (create if absent)

**Interfaces:**
- Consumes: `PREMIUM_SOURCE_KEY` (Task 9), `accountPlanCopy.premiumCard` and `See your plan` (Task 10).
- Produces (§4.4): a head `is:inline` script of at most six lines that copies a stored `own` / `seat` / `grant` onto `document.documentElement.dataset.premium` before first paint (and nothing for any other value). The Premium card holds both versions of the 13 px audience line and the 44 px button slot in fixed-height cells; CSS on `html[data-premium=…]` picks one: the default audience / `Subscribe · {price}`, or the source line and a neutral `SECONDARY_BUTTON_FIXED` link `See your plan` (`border-line-warm-strong text-strong`, never gold) to `/account#plan`. Price block unchanged; guild cards unchanged (R-P7). Zero `_astro` scripts stay.

- [ ] **Step 1: Write the failing render tests:** the page has exactly one more `is:inline` head script than before this task (count them in the test's setup, don't hard-code three: today there is one head inline script and one in the body) and the new one references `fs.premiumSource` and only the three words; the Premium card contains all three source lines and `See your plan` alongside the default line and button, each in a cell with a fixed height class; the CSS selectors `html[data-premium="grant"]` etc. exist; no `_astro` script.
- [ ] **Step 2: Run to fail.** **Step 3: Implement.** **Step 4: Run to pass**, check, lint, prettier. (Case 16, height equality on first paint, runs in Task 26.)
- [ ] **Step 5: Commit.** `git commit -m "feat(web): /premium Premium card shows your plan when you have it" -- web/src/pages/premium.astro web/src/components/premium/PlanCard.astro web/src/pages/_premium.test.ts`

### Task 20: Claim page and `/setup?next=`

**Files:**
- Create: `web/src/components/claim/ClaimPlanLine.svelte`, `web/src/components/claim/ClaimMemberHandoff.svelte`, `web/src/components/claim/claim-handoff.test.ts`, `web/src/lib/account/setup-next.test.ts`
- Modify: `web/src/components/GuildClaim.svelte` (one-line mounts only; it is at 404 lines), `web/src/lib/guild/claim-view.ts` and its test, `web/src/lib/guild/copy.ts` (the §9.3 claim and setup strings), `web/src/pages/setup.astro` and the setup island it mounts (accept `next`, `Back to {Guild}'s claim`), `web/src/lib/account/safe-next.ts` (only if `safeNextPath` needs a claim-path allowance)

**Interfaces:**
- Consumes: `parsePlan`, `parseInterval` (Task 14), `CopyLinkField` (Task 13), `safeNextPath` (`web/src/lib/account/safe-next.ts`), the §5.3 `ready` button labels.
- Produces (§6.2):
  - With `plan`: `ClaimPlanLine` — 14 px `--text` on a 2 px `--gold` left rule, padding-left 14: `This link is for subscribing {Guild} to the {Plan} plan ({price}). Claim the guild first; then the plan is one step.`
  - After a confirmed claim with `plan`: `Next` gains the `ready` button first (`Start the 7-day trial` or `Subscribe · {price}`), the page's one fill; `Approve your officers and members` and `Make an invite link` follow.
  - Signed in, no character in the guild: `notAMember` + inline `Get set up` with `next` set to this claim URL.
  - A member viewer: `ClaimMemberHandoff` — `You are a member of {Guild}, so only its guild master can claim it. Send them this page's link.` and a `CopyLinkField` holding the claim URL with the same `plan`.
  - `/setup` reads `next` through `safeNextPath`; after a save, `Back to {Guild}'s claim` is its one primary.

- [ ] **Step 1: Write the failing tests:** `claim-handoff.test.ts` — plan line verbatim (§9.3 via `specString`); a tampered `plan=premium` renders no plan line; the member hand-off text and the field's URL keep `plan`; after a claim the subscribe button is first in `Next`. `setup-next.test.ts`:
  ```ts
  it('rejects off-site and protocol-relative next', () => {
    expect(safeNextPath('https://evil.example/claim')).toBeNull();
    expect(safeNextPath('//evil.example/claim')).toBeNull();
    expect(safeNextPath('/guild/us/forever/olympus-xxvii/claim?plan=guild&interval=yearly')).toBe('/guild/us/forever/olympus-xxvii/claim?plan=guild&interval=yearly');
  });
  ```
  (Match `safeNextPath`'s real return shape; read it first.)
- [ ] **Step 2: Run to fail.** **Step 3: Implement.** **Step 4: Run to pass** with `claim-view`'s and `GuildClaim`'s tests, check, lint, prettier.
- [ ] **Step 5: Commit.** `git commit -m "feat(web): claim page plan line, member hand-off and setup return" -- <the paths above>`

### Task 21: Guild page, invite and settings hand-offs

**Files:**
- Create: `web/src/components/guild/GuildUnverifiedNote.svelte`, `web/src/components/guild/BillingHandoff.svelte`, `web/src/components/guild/guild-handoffs.test.ts`
- Modify: `web/src/components/Guild.svelte` (one-line mounts), `web/src/lib/guild/copy.ts` (`notVerifiedNote` by approver; `Get the guild claimed`; invite joined plan sentences), `web/src/lib/guild/roster-gate.ts` (`claimEntryFor`: the member's slot reads `Get the guild claimed`), `web/src/lib/guild/api.ts` (home `viewer.verify`), `web/src/components/GuildJoin.svelte`, `web/src/components/guild/GuildBillingSection.svelte` (mount `BillingHandoff`; preselect `plan` and `interval` from the URL; open the upgrade confirm on `upgrade=1`)

**Interfaces:**
- Consumes: `viewer.verify.has_approver` (Task 8), `fitsPlan` and `handoffUrl` (Task 14), `CopyLinkField`, `me.guilds[].covered` (Task 9).
- Produces:
  - `GuildUnverifiedNote`: with an approver, today's sentence; without, `You are not verified yet. Nobody in {Guild} can approve members until its guild master claims it.` + the inline gold link `Send your guild master a link` to the claim page.
  - The header claim slot for a member of an unclaimed guild reads `Get the guild claimed` and links to the claim page.
  - `GuildJoin` joined state: under `You joined the guild.`, 13 px muted, `{Guild} is on the {Tier} plan.` or `{Guild} has no plan yet.` (from `covered` after `forgetSession()` and a `/v1/me` re-read), then `Open the guild`.
  - `BillingHandoff` (§6.4): under today's `Only the guild master can subscribe the guild, release the claim or change billing.`, when the guild has no live plan: `Send them this link.` and a `CopyLinkField` holding the settings `#billing` link with `plan=fitsPlan(verified_count)&interval=yearly`.
  - Settings preselect: `?plan&interval` preselect the launch-pricing §3.2.2 chooser; `upgrade=1` opens the §3.2.4 confirm.

- [ ] **Step 1: Write the failing tests:** note by approver (both forms, verbatim §9.3); the member's claim slot text and href; the joined state's two plan sentences; `BillingHandoff` renders only with no live plan and its URL uses `fitsPlan` (12 verified → `plan=guild`, 28 → `plan=raid_guild`); settings with `?plan=raid_guild&interval=monthly` preselects both; `upgrade=1` opens the confirm.
- [ ] **Step 2: Run to fail.** **Step 3: Implement.** **Step 4: Run to pass** with the guild component tests touched, check, lint, prettier.
- [ ] **Step 5: Commit.** `git commit -m "feat(web): guild page, invite and settings hand-off lines" -- <the paths above>`

### Task 22: Account guild rows

**Files:**
- Create: `web/src/components/account/AccountGuildRowStep.svelte`, `web/src/components/account/AccountGuildRowStep.test.ts`
- Modify: `web/src/components/Account.svelte` (one-line mount inside `account-guild-row`)

**Interfaces:**
- Consumes: `fitsPlan` (Task 14), `checkoutGuildCopy` or the §9.3 account strings in `account-plan-copy.ts`.
- Produces (§6.5): after the guild name, `Not verified` 13 px muted when the account has no verified character there. When `subscribe.step !== 'on_plan'`: `{Guild} has no plan.` 13 px muted and the inline link `See the next step` to `/premium/checkout?plan={fits}&interval=yearly#guild-{id}`. On a plan: today's `AccountGuildPlanLine`.

- [ ] **Step 1: Write the failing tests:** unverified row shows `Not verified`; a no-plan row links to `…?plan=guild&interval=yearly#guild-7` at 12 verified and `plan=raid_guild` at 28; an `on_plan` row renders `AccountGuildPlanLine`; a row with `subscribe` null (older API) renders today's row.
- [ ] **Step 2: Run to fail.** **Step 3: Implement.** **Step 4: Run to pass**, check, lint, prettier.
- [ ] **Step 5: Commit.** `git commit -m "feat(web): account guild rows show the next step" -- web/src/components/account/AccountGuildRowStep.svelte web/src/components/account/AccountGuildRowStep.test.ts web/src/components/Account.svelte`

### Task 23: Lighthouse collect list

**Files:**
- Modify: `web/lighthouserc.json`

**Interfaces:**
- Produces: `http://localhost/premium/checkout.html?plan=guild&interval=yearly` in the collect list, under the default assertion row (performance and accessibility at least 0.95, LCP at most 2400 ms, TBT at most 100 ms, CLS at most 0.05).

- [ ] **Step 1: Add the URL.** **Step 2: Run** `npx prettier --check lighthouserc.json`. The controller runs the shard (Task 28).
- [ ] **Step 3: Commit.** `git commit -m "ci(web): lighthouse covers the checkout guild step" -- web/lighthouserc.json`

---

## PR C — e2e and captures

### Task 24: Stubbed API fixtures

**Files:**
- Modify: `web/src/test-support/billing-fixtures.ts` (or create `web/src/test-support/subscribe-fixtures.ts` if it would pass 400 lines)

**Interfaces:**
- Produces (typed against Task 9's contract):
  - `OWNER_MEMBER` — **exactly as production holds the owner's account** (capture 1): header `Obnoxious Yell` (Fury Warrior, 21); guilds OLYMPUS XXVII and Employed Millennials; rank member in both; not verified; 0 verified in each; unclaimed; no plan; `subscribe.step` `ask_leader` for both; open-ended operator grant (`billing.source` `grant`, `current_period_end` null); no Stripe customer (`personal_portal` false).
  - `OWNER_GUILD_MASTER` (capture 1b): the same after an addon save that makes the owner guild master of OLYMPUS XXVII: that guild `claim`, Employed Millennials `ask_leader`; and `OWNER_GUILD_MASTER_CLAIMED`, OLYMPUS XXVII `ready` with the trial unused.
  - `MIXED` — one `ready` (trial), one `on_plan` member with seats free.
  - One fixture per `step` (13), including `claimed_ask` as officer and as member, both `claim_blocked` reasons, the four `on_plan` variants, and `ready` with 12 verified on Raid guild and 28 on Guild.
  - One per premium source: own renewing with founder, own ending, past due, own then seat, seat (with and without `personal_portal`), grant open-ended, grant until `2027-03-31`, fallback, free.
  - `EMPTY_A` (no characters), `EMPTY_B` (characters, no guild), `OLDER_ME` (no `subscribe`).

- [ ] **Step 1: Write the fixtures** and a type-only test that each one satisfies `Me` (`npm run check`). Move any local fixtures from Tasks 16 to 22 here and re-point those tests.
- [ ] **Step 2: Run** the touched vitest files, check, lint, prettier.
- [ ] **Step 3: Commit.** `git commit -m "test(web): subscribe-and-join fixtures, the owner in both states" -- web/src/test-support/billing-fixtures.ts <re-pointed tests>`

### Task 25: `checkout-guild.spec.ts` — §12 cases 1 to 13 and 18

**Files:**
- Modify: `web/tests/e2e/checkout-guild.spec.ts` (rewritten for the guild step; the launch-pricing chooser cases it held are superseded)

- [ ] **Step 1: Write the e2e cases** with the API stubbed from Task 24's fixtures, one `test` per §12 case:
  1. Owner fixture: two cards in order, `Member · Not verified`, `No one verified yet`, `No plan`, the `ask_leader` sentence, the copy field value `…/claim?plan=guild&interval=yearly`, the short verify line, the other-character line, the Discord fallback linking to the site invite; the two-guild lead; no `Get Premium`.
  2. `Copy link` writes the clipboard and reads `Copied` for 2 s; a denied clipboard shows the failure line.
  3. One `ready` (trial unused): `Start the 7-day trial` is the only filled button; clicking posts plan, interval and guild id.
  4. Two `ready`: both gold outline; no fill on the page.
  5. `claim`: confirm in place, focus on `Cancel`; Escape returns focus; `Yes, claim it` posts, re-reads, card `ready`, focus on its button, status `Claimed. You hold {Guild}.`
  6. `start_claim` then pending: `claim_pending_yours` with a copy field.
  7. Claim refusals (already claimed, not eligible, rate limited, holds another) each show their §9.1 line; no API sentence anywhere.
  8. `claimed_ask` names the holder to an officer, not to a member.
  9. `on_plan`: you pay; you pay on Guild with Raid guild chosen (`upgrade=1`); seat held; seats free; seats full.
  10. `checkout_open` and `frozen`: their sentences, no buy button.
  11. `guild_id`: non-ready first and in view; ready launches Stripe; unknown shows the list.
  12. Size lines: Guild's 10-person line; 12 verified on Raid guild shows `Switch to Guild`, 28 on Guild `Switch to Raid guild`; switching changes band, button and URL with no reload; no `40-person` or `40 players`.
  13. Empty A and B: `Get set up` with `next` back here; no `Battle.net`.
  18. Axe clean in every state above; every control at least 44 px tall at 360; CLS at most 0.02 from loading to ready at 390 and 1440 (a `PerformanceObserver` for `layout-shift` across the skeleton→ready swap).
- [ ] **Step 2: The task runs** `npm run check && npm run lint && npx prettier --check tests/e2e/checkout-guild.spec.ts`; the controller runs Playwright (Task 28).
- [ ] **Step 3: Commit.** `git commit -m "test(web): checkout guild step e2e" -- web/tests/e2e/checkout-guild.spec.ts`

### Task 26: §12 cases 14 to 17 in the existing specs

**Files:**
- Modify: `web/tests/e2e/checkout-launcher.spec.ts`, `web/tests/e2e/premium.spec.ts`, `web/tests/e2e/account-billing.spec.ts`, `web/tests/e2e/account-seat.spec.ts`, `web/tests/e2e/guild-claim.spec.ts`, and the guild page spec that covers the unverified note (`rg -l notVerified\|not-verified web/tests/e2e`)

- [ ] **Step 1: Write the cases:**
  14. `plan=premium` for own renewing, own ending, past due, seat, grant open-ended, grant until: panel title, line, line 2 and button match §4; `POST /v1/billing/checkout` never called (route spy). A stubbed 409 with `fields.source=grant` renders the grant panel and never `this account already has Premium`.
  15. Account: the owner's grant reads `Premium · granted, no end date` with no billing button; a seat holder with `personal_portal` sees `Billing history`; no fixture whose Premium rows are On reads `Not on Premium.`
  16. `/premium` with `fs.premiumSource=grant` set by `addInitScript` before load: the Premium card reads the grant line and `See your plan` on first paint; its height equals the signed-out card's; a stored `premium` (malformed) renders the default and the key is gone.
  17. Guild page, unverified, no approver: the new note and link. Claim page as a member: the copy field. Claim page with `plan`: the plan line, and after a claim the subscribe button first in `Next`.
  Axe and the 44 px rule in each.
- [ ] **Step 2: The task runs** check, lint and prettier on the touched specs.
- [ ] **Step 3: Commit.** `git commit -m "test(web): already-have-it, account, premium hint and claim e2e" -- <the specs above>`

### Task 27: Fixture-build captures, §12 rows 1 to 17

**Files:**
- Create: `web/tests/captures/subscribe-and-join/capture.spec.ts`, `web/tests/captures/subscribe-and-join/playwright.capture.config.ts`

- [ ] **Step 1: Write the capture script** for §12's table at the widths it marks (360, 390, 1440, 2000), from Task 24's fixtures:
  - Row 1: `OWNER_MEMBER` on `/premium/checkout?plan=guild&interval=yearly`, exactly as production holds the account.
  - Row 1b: `OWNER_GUILD_MASTER` (OLYMPUS XXVII `claim`, sorted first; Employed Millennials `ask_leader`; the some-claimed lead), and the card after `Yes, claim it` (`ready`, `Start the 7-day trial`).
  - Rows 2 to 4: the owner on `plan=premium`, `/account#plan`, `/premium`.
  - Rows 5 to 16 as the table lists; row 13 drives focus-visible and hover.
  - Row 17: `GuildInviteLink` before (from `main`'s build) and after the extraction, with a pixel diff that must be zero.
- [ ] **Step 2: The task runs** check, lint and prettier. The controller runs the captures at branch level; ux-designer and wow-player review every row on the rendered page (memory: design before build, two reviewers).
- [ ] **Step 3: Commit.** `git commit -m "test(web): subscribe-and-join captures, the owner in both states" -- web/tests/captures/subscribe-and-join/capture.spec.ts web/tests/captures/subscribe-and-join/playwright.capture.config.ts`

### Task 28: Controller at branch level, then ship

The controller does this by hand; no subagent.

- [ ] **Step 1:** `cd api && go build ./... && go vet ./...` and the touched packages with `-p 1` on a private database.
- [ ] **Step 2:** `cd web && npm run build`, the full `npx vitest run`, the full `npx playwright test`, and the lhci shard (the guild step URL from Task 23; `/premium` unchanged). These pages share the layout and rebuild the checkout page, so the full suites run (memory: shared layout needs the full suite).
- [ ] **Step 3:** Run Task 27's captures; collect ux-designer and wow-player SHIP on every row, starting from capture 1 and 1b.
- [ ] **Step 4:** Merge PR A; deploy the API; confirm in production that user 78's `/v1/me` answers `billing.source` `grant`, `personal_portal` false, and `subscribe.step` `ask_leader` for both guilds. Then merge and deploy PR B and PR C.
- [ ] **Step 5:** Open the Discord bot's `/readiness` refusal on a test guild and confirm its button opens `/premium/checkout?plan=raid_guild&interval=yearly&guild_id={id}`.

---

## Parallel waves

| Wave | API | Web |
|---|---|---|
| 1 | 1, 3 | 9, 13, 23 |
| 2 | 2 | 10, 14 |
| 3 | 4, 6, 7 | 11, 12, 15 |
| 4 | 5, 8 | 16, 19, 20, 22 |
| 5 | | 17, 21 |
| 6 | | 18, 24 |
| 7 | | 25, 26, 27 |
| last | 28 (controller) | |

- Tasks 1, 2, 4 and 5 all edit `auth/guild_standing.go` and `guilds/me_standing.go`, so they run in sequence. Task 3 is a new pure file and runs beside Task 1.
- Tasks 1 and 4 both touch `cmd/api/main.go`; they are in different waves.
- Tasks 12, 15, 16 and 18 all edit `CheckoutLauncher.svelte`; 12 and 15 share wave 3 only if the launcher edits are split (12 the `plan=premium` branch, 15 the band mount) — otherwise run 15 after 12. Tasks 16, 17 and 18 edit the checkout step components in sequence.
- Tasks 10 and 11 edit `AccountPremiumBlock.svelte` in sequence; Tasks 11 and 22 both mount into `Account.svelte` in different waves.
- Tasks 9 and 10 both touch `web/src/lib/account/api.ts` (one-line hooks) in sequence.

## Conflicts ruled

- **Personal refusal name.** The task list says `alreadyOnThePlanRefusal`; that is the guild 409 and stays unchanged. The personal 409 is `alreadyOnPremiumRefusal`, which gains the `source` argument (Task 7).
- **`fields.source` = `seat`.** §8 lists `stripe`, `grant` or `seat`. The API's refusal fires only on a live personal row, so it emits `stripe` or `grant`; who may buy is unchanged (§8, "No change to who may … bill"). The web keeps `seat` in the closed type and blocks seat holders before calling checkout (§4.2, Decision 5).
- **Package boundaries.** `guilds` and `billing` import `auth`, so `auth` cannot import either. The standing read is an interface in `auth` (`GuildStandingSource`, like `SeatFactsSource`), implemented by `guilds.Store`, with billing's predicates behind `guilds.StepBillingFacts` that `billing.Store` already satisfies. `SubscribeStep` itself is pure and sits beside `Capabilities`, as §8 asks.
- **`claim_blocked` vs `start_claim` / `claim` / `claim_link_bnet`.** §8 orders `claim`, `claim_link_bnet`, `start_claim` before `claim_blocked`; each of the first three requires "not blocked", so a blocked officer or guild master answers `claim_blocked` (Review Focus 1).
- **`seats_given`.** Counts slots within the tier's seat count whose state is `held` or `accepted`; an `offered` seat is not yet given.
- **The bot's billing link.** `cmd_readiness.go`'s `billingAnchor` (`/settings#billing`) moves to the checkout guild step URL, and `cmd_readiness_test.go` changes with it (Task 8). The bot's sentences are unchanged (§6.6).
- **Plan-asks route.** Deferred to v2 (Decision 2): `POST /v1/guilds/{id}/plan-asks` is not in this plan.
- **`checkout-guild.spec.ts`** exists from the launch pricing lane; Task 25 rewrites it for the guild step rather than adding a second file.

## Pre-flight corrections (controller, 2026-10-27)

From `.superpowers/sdd/2026-10-27-subscribe-and-join/preflight.md`; implementers follow these over the task text where they differ.
- Task 9: the `Me` types live in `web/src/lib/account/api.ts` (799 lines, at the 800 cap), not `contract.ts`. Put the new `guilds[]` and `entitlements` field types in a new `web/src/lib/account/me-guild-types.ts` and import them; one-line edits only in `api.ts`.
- Task 10: the `manageBilling` reader is in `Account.svelte`, not `AccountPremiumBlock.svelte`; reuse Premium-value Task 3's spec-table helper (`specKeyCell`) and the shared `allStrings` helper instead of a local `tableRow`.
- Task 14: `parsePlan` / `parseInterval` do not exist yet — extract them from the checkout page's inline parsing into `web/src/lib/billing/plan-params.ts` with tests; `fitsPlan` is the existing `guildFitsRaidTier` (reuse it); delete `billableGuilds` in Task 16 when the chooser goes. Build every `7` in `First 7 days free` / `Start the 7-day trial` from `GUILD_TRIAL_DAYS`.
- Task 19: runs AFTER Premium-value Tasks 11–16 and reuses their `renderPremium()`; the Premium card's doubled slots stay single subgrid cells beside PV 12's per-plan button split.
- Task 20: `safeNextPath` returns the fallback, never `null` — test that.
- Task 23 and PV 21 edit `lighthouserc.json` one at a time; Task 26 follows PV 21 in `premium.spec.ts`; Task 27's `/premium` capture follows PV 16 and PV 19.
- Both plans' branch-level verification (build, vitest, Playwright, lhci) runs ONCE, by the controller; deploy the API (PR A) first — the web handles the older `/v1/me` shape.
