# Entitlements, Stripe payments, and the premium page

**Date:** 2026-09-21
**Status:** DRAFT for coordinator review. Implements proposal (`2026-09-21-guild-and-premium-
proposal.md`) sections 4.2 items 1–4 and the player half of the premium tier in §4.1: the
entitlements model, Stripe integration, the premium page, and the legal pages. It gates what
already exists (server sims) and adds the tables and API every later premium feature reads.
It does **not** build officer tools, the performance analyzer, or ratings — those are out of
scope (§9) and read `Can()` when they land.

Every **RULING** below is a call this spec had to make where the proposal, the codebase, or
Stripe's own docs left a gap. Each names its reason and the cost of being wrong, so the
coordinator can overrule it cheaply.

## 0. What exists today, and what this spec assumes from a parallel spec

**Read, not built here** (from the codebase — see the grep in the brief, repeated in full in
the sections below where it matters):

- `users.premium boolean not null default false` (migration `0013_sims.up.sql`) is the only
  gate that exists. It is read by `auth.Store.Premium` (`api/internal/auth/store.go:429`),
  carried on `auth.User.Premium` (`store.go:53`, JSON `premium`), and checked exactly once,
  in `sims.Service.run` (`api/internal/sims/run.go:34`) via the `Premiumer` interface
  (`run.go:19`). It is set by hand today — "no code here ever turns it on" (`store.go:52`).
  It has no expiry, no source, no per-feature granularity, and nothing writes it.
- Web reads it from `GET /v1/me`'s `user.premium` (`web/src/lib/account/api.ts:54`) into
  `sim/store.svelte.ts`'s `premium` state (`store.svelte.ts:171`, set by `SimView.svelte:355`
  and `ToolsView.svelte:106`), which drives `RunControl.svelte`'s "Run on our servers"
  button, `BulkRunBar.svelte`'s cap-lifted note, and `sim/copy.ts`'s `premiumRequired`
  (line 139), `capPremium` (985) and `capPremiumNote` (986) strings. `sim/api.ts:24`'s
  `PREMIUM_REQUIRED_STATUS = 402` is how a failed run is told apart from any other failure —
  the web never infers premium state from a 402, only from `user.premium` on the shared
  `/v1/me` load (`api.ts:13`'s own comment). This status-code contract does not change here.
- The API's conventions this spec follows exactly: the envelope (`httpx.WriteOK`/`WriteError`,
  `api/internal/httpx/envelope.go`), double-submit CSRF on every state-changing session
  request (`api/internal/auth/session.go`'s `Authenticator.Middleware`, cookies `fs_session`/
  `fs_csrf`, header `X-CSRF-Token`), per-route rate limiting (`httpx.RateLimitPer`,
  `middleware.go:189`, the pattern `auth.Mount` already uses for `POST /v1/auth/email`), and
  one binary dispatched by its first `os.Args` entry for every Cloud Run job
  (`api/cmd/api/main.go:58`'s `switch os.Args[1]`, alongside `reports.ParseJobCommand` and
  `sims.SimRunJobCommand`).
- Secrets live in Google Secret Manager, bound to the `api-runtime` service account, and are
  wired into the running service by a **one-time, by-hand** `gcloud run deploy` with
  `--set-secrets` (`api/README.md:399-405`); the GitHub Actions deploy step
  (`.github/workflows/api.yml`, `deploy` job) only ever passes `--image` and `--memory` —
  it never lists secrets or env vars, so a Cloud Run revision **keeps whatever secrets and
  env vars the last by-hand `gcloud run deploy`/`services update` set**, and CI's image-only
  deploys carry them forward automatically. Adding a secret is therefore always: create it in
  Secret Manager, grant `api-runtime` the `secretAccessor` role, then one manual
  `gcloud run services update api --update-secrets ...` — never a workflow-file change.
- Migrations: the newest applied one is `0017_reports_recent_idx`. Two parallel, not-yet-
  merged specs claim the next two numbers: `2026-09-21-guild-membership-design.md` is
  `0018`, and the performance-rating spec is `0019`. **This spec's migration is therefore
  `0020_entitlements`, with a follow-up `0021_drop_users_premium` for §1.2's second step.**
  The coordinator assigns final numbers at merge if the landing order among these three
  specs turns out different from the order they were drafted in — a renumbering `mv` with
  no functional change, since `golang-migrate` orders by the numeric prefix alone and none
  of these three specs' migrations reads a column another one of them adds.
- **What this spec assumes from the guild membership spec**, stated explicitly since that
  spec is a draft and was revised after this spec was drafted against it — this list
  reflects the revision, not the version originally read:
  1. `guilds.id` is the stable identity a guild entitlement attaches to (`guilds` table,
     `0005_logs.up.sql:61`, unchanged by `0018`).
  2. Membership is now tracked **per character**, in a new `guild_characters` table
     (`0018`), not per account directly. `guild_members (guild_id, user_id)` — the table
     this spec's queries read throughout, unchanged in shape from `0005_logs.up.sql:73` —
     is a **derived, per-account** access row: `rank` (`'member' | 'officer' | 'leader'`)
     and `verified_at timestamptz` are computed from that account's own `guild_characters`
     rows, and `verified_at` is set (i.e. the account counts as a verified member) once
     **at least one** of the account's characters in that guild is itself verified.
     A character is verified by any of: two distinct raid nights in the guild's own logs
     within a 30-day window; officer approval; redeeming the guild's invite link; or the
     claim flow. **This spec makes no query against `guild_characters` directly** — every
     `Can()`/checkout check below reads `guild_members.verified_at` exactly as written,
     and that column's derivation is entirely the guild membership spec's concern.
     **A claimed guild has officers**: `guilds.claimed_by` names the account that claimed it
     (`0005_logs.up.sql:66`); rank `officer` or `leader` with `verified_at is not null` is
     that guild's `VerifiedGuildRank` — this spec's own phrase for it, matching
     `guild-membership-design.md §2.6`'s naming, is used identically below.
  3. If `0018` changes these column names or the verification rule again before it merges,
     the one place this spec touches is §1.3's `Can()` guild-membership query and §2.9's
     "who may buy a guild plan" check — both named below so the fix is localized.

## 1. Entitlements model

### 1.1 Migration `0020_entitlements`

```sql
-- 0020_entitlements.up.sql
-- One row per (subject, plan): a user's premium entitlement, or a guild's guild-plan
-- entitlement. Polymorphic by two nullable FKs rather than a single untyped subject_id,
-- so referential integrity is real (a deleted user or guild cascades its own entitlement
-- away) and a query never has to trust an unenforced subject_type string.

create table entitlements (
  id                      bigserial primary key,
  user_id                 bigint references users(id) on delete cascade,
  guild_id                bigint references guilds(id) on delete cascade,
  plan                    text not null check (plan in ('premium', 'guild')),
  source                  text not null check (source in ('stripe', 'grant', 'trial')),
  status                  text not null check (status in
    ('active', 'trialing', 'past_due', 'canceled', 'incomplete', 'incomplete_expired', 'unpaid')),
  current_period_end      timestamptz,
  cancel_at_period_end    boolean not null default false,
  -- grace_until is the retention grace only (proposal §4.1's "2 years, then 90 days after a
  -- 30-day grace"), not a feature-access grace: feature access follows `status` and
  -- `current_period_end` directly. See §2.7.
  grace_until             timestamptz,
  stripe_subscription_id  text unique,
  -- Set only for source = 'stripe': the user whose Stripe Customer backs this subscription.
  -- For plan = 'premium' this is always user_id itself; for plan = 'guild' it is the
  -- officer who bought it (§2.9), which can differ from every guild_members row it benefits.
  billing_user_id         bigint references users(id) on delete set null,
  -- Set only for source = 'grant': who ran the grant CLI and why (§1.5).
  granted_by              bigint references users(id) on delete set null,
  grant_note              text,
  created_at              timestamptz not null default now(),
  updated_at              timestamptz not null default now(),
  constraint entitlements_one_subject check (
    (user_id is not null and guild_id is null and plan = 'premium') or
    (guild_id is not null and user_id is null and plan = 'guild')
  ),
  -- NULLs are distinct in a unique constraint, so a guild-plan row (user_id null) never
  -- collides with a premium row here, and vice versa: each constraint only ever fires
  -- within its own subject type.
  constraint entitlements_user_plan_unique unique (user_id, plan),
  constraint entitlements_guild_plan_unique unique (guild_id, plan)
);
create index entitlements_guild_idx on entitlements (guild_id) where guild_id is not null;
create index entitlements_billing_user_idx on entitlements (billing_user_id)
  where billing_user_id is not null;

-- One Stripe Customer per site account, created lazily on first checkout (§2.4).
create table stripe_customers (
  user_id             bigint primary key references users(id) on delete cascade,
  stripe_customer_id  text not null unique,
  created_at          timestamptz not null default now()
);

-- The webhook's idempotency log and its audit trail in one: every event Stripe ever sent
-- us, keyed by Stripe's own event id, so a redelivery (which Stripe both retries for three
-- days on a non-2xx and lets you resend by hand for up to 30) is detected by a unique-key
-- conflict rather than re-applied. `payload` is the raw event JSON — subscription events
-- never carry card data (§3), so storing it verbatim is the audit record for "what Stripe
-- actually told us," independent of how the handler interpreted it.
create table stripe_events (
  id            text primary key,
  type          text not null,
  payload       jsonb not null,
  received_at   timestamptz not null default now(),
  processed_at  timestamptz
);

-- Our own interpretation of every entitlement change, independent of source: a Stripe
-- webhook, the nightly reconciliation job, or a CLI grant. This is what "audit logging of
-- entitlement changes" (§3) means concretely — stripe_events above is a log of what Stripe
-- sent; this is a log of what we did about it.
create table entitlement_audit (
  id              bigserial primary key,
  entitlement_id  bigint references entitlements(id) on delete set null,
  user_id         bigint references users(id) on delete set null,
  guild_id        bigint references guilds(id) on delete set null,
  plan            text not null,
  old_status      text,
  new_status      text not null,
  -- 'stripe_webhook:<event_id>' | 'stripe_reconcile' | 'cli_grant:<operator note>' |
  -- 'cli_revoke:<operator note>'. A free-text tag rather than a foreign key: the actor is
  -- sometimes not a user row at all (a webhook, a cron job), and this column exists to be
  -- read by a human debugging a billing question, not joined on.
  actor           text not null,
  created_at      timestamptz not null default now()
);
create index entitlement_audit_entitlement_idx on entitlement_audit (entitlement_id, created_at desc);
```

```sql
-- 0020_entitlements.down.sql
drop table if exists entitlement_audit;
drop table if exists stripe_events;
drop table if exists stripe_customers;
drop table if exists entitlements;
```

**RULING 1 — one row per subject+plan, upserted, not an append-only ledger.** A
subscription's lifecycle (trial → active → past_due → canceled) is one Stripe object with
one id; modeling it as a mutated row matches that directly, and `entitlement_audit` already
is the ledger for "what changed and when." *Cost if wrong:* if a subject is ever allowed two
overlapping entitlements of the same plan concurrently (not currently possible — Stripe
enforces one subscription per Customer per Price the way this integration uses it, and the
unique constraints enforce it here too), the schema would need a real ledger; cheap to add
later as `entitlements_history`, not a breaking change to `Can()`'s read path.

### 1.2 Migrating `users.premium`, in two steps (zero-downtime)

A single migration that both creates `entitlements` and drops `users.premium` would break a
rolling deploy: the old binary is still reading `users.premium` while the new schema no
longer has it. This is an **expand-and-contract**, matching how `0013`→`0014`→...  already
added simulator columns incrementally without ever dropping one out from under a running
revision.

1. **`0020_entitlements.up.sql`** (above) also backfills, in the same migration, one
   `entitlements` row per account that has `premium = true` today:
   ```sql
   insert into entitlements (user_id, plan, source, status, granted_by, grant_note)
   select id, 'premium', 'grant', 'active', null, 'migrated from users.premium'
   from users where premium = true
   on conflict (user_id, plan) do nothing;
   ```
   `users.premium` is **left in place** by this migration. The Go code deployed alongside it
   stops reading the column (§1.4) but the column still exists, so a mid-rollout mix of old
   and new revisions both work — the old revision reads a column nobody writes anymore
   (harmless: nothing turned it on before this spec either), the new one reads `entitlements`.
2. **`0021_drop_users_premium`**, a follow-up migration landed only after the entitlements
   code has been deployed and confirmed (the next normal deploy, not a same-day one):
   ```sql
   -- 0021_drop_users_premium.up.sql
   alter table users drop column if exists premium;
   -- 0021_drop_users_premium.down.sql
   alter table users add column if not exists premium boolean not null default false;
   ```
   The down migration cannot restore the *values* (there is no revert path from entitlements
   back to a boolean once other sources — Stripe, guild membership — exist), only the
   column shape; this is stated in the migration's own comment so a future rollback does not
   assume otherwise.

### 1.3 `Can`, the one function every feature check goes through

```go
// package entitlements (new, api/internal/entitlements/)

type Feature string

const (
	// FeatureServerSims gates POST /v1/sims/run (sims.Service.run, run.go:34). The cap
	// itself is unchanged: simapi.Caps[simapi.LaneServer] (5,000 combinations, any
	// precision) — there is one premium tier at launch, so Can does not return a cap, only
	// yes/no; if a second tier is ever added, the cap moves onto EntitlementsView (§1.4)
	// as its own field rather than overloading this signature.
	FeatureServerSims    Feature = "server_sims"
	// FeatureRetention gates the 2-year raw-log/saved-sim retention tier vs. the free
	// 90-day one (§2.7's nightly purge job reads this per-account, not per-request).
	FeatureRetention     Feature = "retention"
	// FeatureMultiCompare gates comparing more than two sims/builds side by side.
	FeatureMultiCompare  Feature = "multi_compare"
	// FeatureHistory gates a character's history chart over time.
	FeatureHistory       Feature = "history"
	// FeatureNotifications gates the Discord-webhook notification channel (vs. email on
	// request, which stays free). A guild entitlement additionally unlocks the *guild's*
	// shared webhook, which is guild-settings surface this spec does not build (§9).
	FeatureNotifications Feature = "notifications"
	// FeatureOfficerViews gates the raid-night sheet, wipe analysis, loot council helper,
	// readiness board and assignments (proposal 3.3) — guild plan AND a verified officer
	// or leader of that specific guild, never granted by an individual premium entitlement
	// alone (§1.3.2's ruling).
	FeatureOfficerViews  Feature = "officer_views"
	// FeatureRosterCheck gates the bulk rating lookup: paste a roster or applicant list,
	// see everyone's rating; another player's trend over weeks; side-by-side comparison.
	// The single-player lookup itself stays free (proposal 3.4) and needs no Can() check.
	FeatureRosterCheck   Feature = "roster_check"
)

type Reason string

const (
	ReasonEntitled       Reason = "entitled"
	ReasonSignInRequired Reason = "sign_in_required" // userID == 0
	ReasonNoPlan         Reason = "no_plan"           // no active entitlement, personal or guild
	ReasonNotOfficer     Reason = "not_officer"        // the guild has the plan; caller is not a verified officer/leader of it
	ReasonGracePeriod    Reason = "grace_period"        // informational only; never returned by Can itself (§2.7)
)

// Can answers whether userID may use feature right now, and why — the reason is for the
// API's own error body and the web's copy, not a control flow branch the caller inspects
// (a feature is either usable or it is not; the reason only changes what sentence is shown).
//
// signed-out (userID == 0) always answers (false, ReasonSignInRequired, nil): every caller
// in this codebase already knows its actor before asking Can, since every mounted route
// that reaches it is behind auth.Require or auth.RequireSession — this is a defensive
// floor, not a path any real caller exercises.
func (s *Store) Can(ctx context.Context, userID int64, feature Feature) (bool, Reason, error)

// IsSupporter is a lighter question than Can: does this account show the supporter mark
// (proposal §4.1's last row) — true for an active personal premium entitlement, or
// verified membership (rank irrelevant — the mark is not an officer perk) in a guild that
// currently has the guild plan. Public surfaces (a character page, a guild page) call this,
// never Can, because a supporter mark is cosmetic and must not require the caller to be the
// subject's own session.
func (s *Store) IsSupporter(ctx context.Context, userID int64) (bool, error)
```

**How `Can` resolves a feature**, in order:

1. `userID == 0` → `(false, ReasonSignInRequired, nil)`.
2. **`FeatureOfficerViews`** is the one feature with its own path: it is true only if the
   caller has a `guild_members` row with `verified_at is not null` and `rank in ('officer',
   'leader')` for some guild that has an `entitlements` row with `plan = 'guild'` and
   `status in ('active', 'trialing', 'past_due')`; otherwise `(false, ReasonNotOfficer, nil)`
   when such a guild exists but the caller is not verified-officer-rank in it, or
   `(false, ReasonNoPlan, nil)` when no guild the caller belongs to has the plan at all — the
   API endpoint decides which of the two to surface (§4).
3. Every other feature: true if either (a) the caller has an `entitlements` row
   `(user_id = userID, plan = 'premium', status in ('active','trialing','past_due'))`, or
   (b) the caller has a `guild_members` row with **`verified_at is not null`** (any rank —
   §1.3.2's ruling) for a guild whose `entitlements` row is `(plan = 'guild', status in
   (...))`. Otherwise `(false, ReasonNoPlan, nil)`.
4. `status` alone decides access, not `current_period_end`: `past_due` still counts (Stripe's
   own smart retries are still trying, and revoking on the first missed payment before the
   grace Stripe itself offers is harsher than Stripe's own guidance recommends — see §2.6's
   citation); `canceled`, `unpaid` and `incomplete_expired` do not. A `cancel_at_period_end`
   subscription stays `active` (Stripe does not flip `status` until the period actually
   ends), which is exactly proposal §4.1's "cancelling keeps access to period end."

#### 1.3.1 Officer-only vs. every-member, restated as a table

| Feature | Personal premium | Guild plan, any verified member | Guild plan, verified officer/leader only |
|---|---|---|---|
| `server_sims`, `retention`, `multi_compare`, `history`, `notifications`, `roster_check` | yes | yes | — |
| `officer_views` | no (proposal 3.3: officer tools are never an individual perk) | no | yes |

**RULING 2 — guild-plan benefits (everything but `officer_views`) require `verified_at is
not null` on the caller's `guild_members` row, the same threshold the guild membership spec
already drew for report visibility, not a looser one.** The alternative — any row, verified
or not — is exactly the abuse this spec's §3 was asked to consider: a self-reported `guild`
FS1 export costs nothing to fabricate (`addon_exports.PutExports`, `RequireDevice`-only, no
server-side proof the exporting character is actually in that guild) — an attacker who names
a real, guild-plan guild on a hand-crafted export would get an unverified `guild_characters`
row for free, and if that alone conferred benefits, free premium features at zero cost, at
scale, for as long as they keep re-exporting. Requiring verification closes it: per the
guild membership spec's revision, a *character* is verified by any of two distinct raid
nights in the guild's own logs within a 30-day window, officer approval, redeeming the
guild's invite link, or the claim flow — each of which costs the attacker something they
cannot fabricate alone (a combat log that already names the character as a participant twice
in a month, a real officer's approval, or an officer's own out-of-band invite link) — and an
account's `guild_members.verified_at` is set only once at least one of its characters clears
one of those bars, per §0's assumption. *Cost if wrong:* a brand-new, legitimate member whose
only character so far has an unverified `guild_characters` row sees "not yet verified" on a
guild that already has the plan, until that character's second raid night, an officer's
approval, or an invite redemption verifies it — typically the same raid week; a real but
temporary UX gap, not a security hole either direction.

**RULING 3 — `officer_views` is guild-plan-and-officer, never unlockable by an individual's
own premium subscription.** The proposal's tiers table marks it "officers" only, under the
guild-plan column, with a blank under personal premium; §1.3.1 above is that table restated
as code paths. *Cost if wrong:* if the intent was ever "any premium officer, even of a
non-paying guild, gets officer views for free," this reads as under-selling the guild plan —
cheap to relax later (drop the guild-entitlement join, keep the rank/verified check), but
starting permissive and then taking a feature away from officers who got used to it is the
worse direction to be wrong in, so this spec starts strict.

### 1.4 `GET /v1/me`'s new shape

`auth.Service.me` (`handler.go:258`) gains an `Entitlements` block, computed server-side so
the web makes zero extra calls to render every premium control's state. `auth.Me` becomes:

```go
type Me struct {
	User         User             `json:"user"`
	Characters   []Character      `json:"characters"`
	Guilds       []Guild          `json:"guilds"`
	Entitlements EntitlementsView `json:"entitlements"`
}

// EntitlementsView is every Can() feature pre-resolved for the caller, plus the caller's
// own personal billing state (nil when they have no personal premium subscription — they
// may still see every feature as true via a guild's plan, which EntitlementsView does not
// distinguish source for: a feature is on or it is not, and Billing is specifically "what
// does *this account's own* subscription look like," for the account page).
type EntitlementsView struct {
	ServerSims    bool          `json:"server_sims"`
	Retention     bool          `json:"retention"`
	MultiCompare  bool          `json:"multi_compare"`
	History       bool          `json:"history"`
	Notifications bool          `json:"notifications"`
	OfficerViews  bool          `json:"officer_views"`
	RosterCheck   bool          `json:"roster_check"`
	SupporterMark bool          `json:"supporter_mark"`
	Billing       *BillingView  `json:"billing"`
}

type BillingView struct {
	Plan              string  `json:"plan"`                 // always "premium" here
	Status            string  `json:"status"`
	CurrentPeriodEnd  *string `json:"current_period_end"`   // RFC 3339, null for a grant with no expiry
	CancelAtPeriodEnd bool    `json:"cancel_at_period_end"`
}
```

`Guild` (`auth.Guild`, already returned per membership — `store.go`'s `Guilds` query) gains
one field, `Plan *GuildBillingView`, non-nil only when that guild has an active/trialing/
past_due `entitlements` row and the caller is a **verified officer or leader** of it (a
non-officer member sees the guild is on the plan implicitly, through the features it already
unlocks for them via §1.3's rule 3(b) — they do not need the billing detail, which is why
this is gated tighter than membership alone, matching the settings-page visibility rule the
guild membership spec already applies to everything billing-adjacent):

```go
type GuildBillingView struct {
	Status            string  `json:"status"`
	CurrentPeriodEnd  *string `json:"current_period_end"`
	CancelAtPeriodEnd bool    `json:"cancel_at_period_end"`
	// BilledBy is who the guild's Stripe Customer belongs to — "you" when it is the
	// caller, otherwise a battletag (or "a former member" if that account has since left
	// the guild — §2.9's transfer ruling). Never an email or a Stripe id.
	BilledBy          string  `json:"billed_by"`
	// YouAreBillingContact is true only for the account whose Stripe Customer backs this
	// subscription — the only account the portal (§4) will open for. Every other verified
	// officer sees the block read-only.
	YouAreBillingContact bool `json:"you_are_billing_contact"`
}
```

This block is **additive** on the guild membership spec's own `GET /v1/guilds/{id}/settings`
response too (that spec's §2.6/§4.4 own the route; this spec only adds the `billing` key to
its response body, not a new route) — named here as the one cross-spec coordination point,
per the brief's instruction to state cross-spec assumptions explicitly.

### 1.5 Manual grants (testers, supporters)

**RULING 4 — a CLI subcommand, not an admin HTTP endpoint.** The image already dispatches on
`os.Args[1]` for every Cloud Run job (`main.go:58`); a `grant`/`revoke` pair is the same
mechanism, run by hand (`docker run <image> grant ...` locally against a `DATABASE_URL`, or
`gcloud run jobs execute` once a tiny throwaway job resource exists) rather than a new
authenticated HTTP surface. *Cost if wrong:* an HTTP endpoint would be more convenient for a
non-engineer to use later (the owner, without a terminal) — cheap to add on top of the same
`entitlements.Store.Grant`/`Revoke` functions this CLI calls, as `POST /v1/admin/entitlements`
behind `role = 'admin'` (the `role` column and `Actor.IsModerator`-style checks already
exist, `session.go`), if that need arises. Starting with a CLI keeps a money-adjacent write
path off the public internet entirely until there is a concrete reason to widen it — the
security-first default.

```
api grant --user <id|battletag> --plan premium [--until 2026-12-31] --note "beta tester"
api revoke --user <id|battletag> --plan premium --note "requested by support"
```

Each upserts an `entitlements` row (`source = 'grant'`, `status = 'active'`,
`current_period_end = --until` or null for open-ended, `granted_by` resolved from the
operator's own shell — `$USER`/`whoami`, recorded in `grant_note` alongside the `--note`
text) and writes one `entitlement_audit` row, `actor = 'cli_grant:<whoami> <note>'` (or
`cli_revoke:...`). `--user` accepts a battletag for convenience and resolves it against
`users.battletag` before touching anything, refusing ambiguity (ties) or a miss with a clear
message rather than guessing.

## 2. Stripe integration

### 2.1 Products and prices, created by an idempotent setup command

**RULING 5 — one more `os.Args[1]` subcommand, `stripe-setup`, not a hand-run Dashboard
click-through.** Consistent with §1.5's CLI-over-Dashboard preference and with the proposal's
own instruction ("created by an idempotent setup command, not by hand"): a Dashboard-created
price has no guaranteed id across test and live mode, while a lookup key set by code is
identical in both (Stripe's Price object: "`lookup_key` — a lookup key used to retrieve
prices dynamically from a static string," unique per mode). Four prices, two products:

| Product | Price | `lookup_key` | Amount |
|---|---|---|---|
| Forever Sixty Premium | monthly | `premium_monthly` | $4.00 / month |
| Forever Sixty Premium | yearly | `premium_yearly` | $40.00 / year |
| Forever Sixty Guild | monthly | `guild_monthly` | $15.00 / month |
| Forever Sixty Guild | yearly | `guild_yearly` | $150.00 / year |

`stripe-setup` lists prices by lookup key first (`prices.List` with `LookupKeys`); creates
the product and price only when missing; never mutates an existing price's amount (Stripe
prices are immutable once created — the amount is fixed at creation). Changing a price later
means: create a new Price with the intended amount and no lookup key yet, clear the
`lookup_key` off the old Price (an update, not a delete — Stripe keeps at most one *active*
price per lookup key, which is exactly why clearing the old one first is required), then set
the new Price's `lookup_key` to the freed value. `stripe-setup` is safe to run repeatedly and
against either mode (it reads `STRIPE_SECRET_KEY` from the environment like everything else
here, so running it against a test key and then a live key sets up both independently).

The four lookup keys are the **only** thing application code hardcodes; every checkout call
resolves a Price by lookup key at request time (with a short in-process cache — the four
values change only when the owner deliberately reprices, not per request) rather than
carrying a price id in the codebase, so test and live mode run identical code paths.

### 2.2 Checkout Sessions

`POST /v1/billing/checkout` (§4) builds a Checkout Session with:

- `mode: "subscription"`
- `line_items: [{price: <resolved price id>, quantity: 1}]`
- `customer: <the caller's Stripe Customer id>` — created lazily on first checkout and
  reused after that (§2.4); **never** `customer_email`, which would let Checkout create a
  second, unlinked Customer object for a returning buyer.
- `client_reference_id`: `"user:<id>"` for a premium purchase, `"guild:<guild id>:user:<id>"`
  for a guild purchase — a cross-check against `metadata` (below) on the confirming webhook,
  not the primary source of truth (metadata is, because it is also copied onto the
  Subscription object itself via `subscription_data.metadata` — see the next bullet and
  §2.6's ordering note).
- `metadata` **and** `subscription_data.metadata`, identical: `{"plan": "premium"|"guild",
  "user_id": "<id>", "guild_id": "<id or omitted>"}`. Both are set, not just one: the
  Checkout Session's own `metadata` is only readable from `checkout.session.completed`, but
  every later `customer.subscription.*` event carries only the *Subscription's* metadata —
  so `subscription_data.metadata` is what makes every later event self-describing without a
  round trip back to the (by-then-expired) Checkout Session.
- `automatic_tax: {enabled: true}` — Stripe Tax, per the proposal's decision; requires
  `billing_address_collection` to default to `auto`, which Checkout already does when
  automatic tax is on (it collects only the address fields tax calculation needs).
- `allow_promotion_codes: true`. **RULING 6 — promotion codes allowed.** The proposal names
  "coupons beyond Stripe promotion codes" as out of scope (§9), implying promotion codes
  themselves are in scope; this is the one flag that turns them on in Checkout, and it costs
  nothing to leave on with no codes issued yet — the owner creates codes in the Dashboard
  whenever wanted, with no code change here. *Cost if wrong:* none either direction; flipping
  it off later is a one-line change if unsolicited code-guessing in the Checkout UI is ever a
  problem (Stripe already rate-limits invalid code attempts itself).
- `success_url`: `{PublicBaseURL}/premium/checkout?status=success&session_id={CHECKOUT_SESSION_ID}`
- `cancel_url`: `{PublicBaseURL}/premium?canceled=1`

Both URLs are built server-side from `cfg.PublicBaseURL` only — nothing in the request body
influences them, so there is no open-redirect surface here (§3).

### 2.3 The guild plan's checkout

`POST /v1/billing/checkout` with `plan: "guild"` additionally requires `guild_id` and:

1. The guild named by `guild_id` exists and `guilds.claimed_by is not null` (a guild must be
   claimed before it can be sold the plan — an unclaimed guild has no accountable officer).
2. The caller is that guild's `VerifiedGuildRank` (§0's assumed definition: `guild_members`
   row, `verified_at is not null`, `rank in ('officer', 'leader')`) — **not** necessarily
   `guilds.claimed_by` itself. **RULING 7 — "an officer of a claimed guild" (proposal §5,
   answered) means any verified officer/leader, not only the claimant.** The claim and the
   subscription are deliberately separate facts: a guild's claim can move (release, re-claim
   by a different officer, §2.4 of the guild membership spec) without touching billing, and
   requiring the literal claimant to be the one who personally pays would make the guild
   plan unbuyable the moment the original claimant goes inactive but stays the recorded
   claimant. *Cost if wrong:* a guild could end up with a claimant and a billing contact who
   are different people and do not coordinate — mitigated by `GuildBillingView.BilledBy`
   (§1.4) being visible to every verified officer, not just the buyer, so this is surfaced
   rather than hidden.
3. The guild does not already have an active/trialing/past_due `entitlements` row for
   `plan = 'guild'` — if it does, the API answers `409 conflict` with a `portal_hint: true`
   field so the web can offer "Manage the guild's billing" (§4) instead of a second
   subscription.

### 2.4 Customer reuse

`stripe_customers` (§1.1) is looked up by `user_id` before every checkout; on a miss, a
Stripe `Customer` is created (`email` from `users.email` when present, otherwise omitted —
a Battle.net-only account has no email on file and Checkout collects one during the flow
if needed) and the mapping is written in the same request, before the Checkout Session is
created, so a crash between the two never leaves an orphaned Customer with no mapping (the
mapping write and the Customer creation are both idempotent to retry: a second attempt that
finds no `stripe_customers` row would create a second Customer only if the first attempt's
write genuinely failed, which is the accepted, rare failure mode — cleaned up by the nightly
reconciliation job noticing a Customer with no matching subscription is harmless clutter, not
by any code path here).

### 2.5 The Customer Portal

`POST /v1/billing/portal` (§4) creates a `billing_portal.Session` with `customer: <the
caller's Stripe Customer id>` and `return_url: {PublicBaseURL}/account`, and answers the
session's `url`. **OWNER:** the portal's feature set (cancel, switch plans, update payment
method, invoice history, tax id collection) is configured once, by hand, in the Stripe
Dashboard's Customer Portal settings before this ships — the API does not create a portal
*configuration*, only *sessions* against whatever the default configuration allows, per
Stripe's own integration order ("configure the portal's features... before you integrate").
Test mode and live mode each need this done once, independently (Stripe keeps them separate).

A guild's billing is managed through this same endpoint by whichever account is
`YouAreBillingContact` (§1.4) for that guild — a portal session is per-Customer, not
per-guild, and the guild's Customer *is* the billing officer's own personal Customer object
(one Customer can hold both a personal premium subscription and a guild subscription they
bought; the portal shows both, distinguished by product name).

### 2.6 The webhook

`POST /v1/billing/webhook` — **no session auth, no CSRF** (Stripe never sends the `fs_session`
cookie, so `Authenticator.Middleware` already treats it as anonymous and skips the
state-changing CSRF check automatically — no special-casing needed in this router, only in
what the handler itself trusts, which is Stripe's signature and nothing else, §3).

```go
const maxWebhookBody = 64 << 10 // Stripe's own Go example uses 65536; a subscription
                                 // event's JSON is a few KB, so this is generous headroom.

func (s *Service) webhook(w http.ResponseWriter, r *http.Request) {
	payload, err := io.ReadAll(http.MaxBytesReader(w, r.Body, maxWebhookBody))
	if err != nil { /* 400, body too large or unreadable */ }

	event, err := webhook.ConstructEvent(payload, r.Header.Get("Stripe-Signature"), s.WebhookSecret)
	if err != nil { /* 400 — bad signature, wrong secret, or the body was mutated in
	                    transit; net/http never touches the body itself, so "mutated" here
	                    would mean a proxy in front of this service rewrote it — Cloud Run's
	                    own front end does not */
	}

	// Idempotency: an INSERT that loses the race (a redelivery, or Stripe's own retry
	// after a slow-but-eventually-200 first attempt) is answered 200 with no reprocessing.
	inserted, err := s.Store.RecordEventOnce(ctx, event.ID, event.Type, payload)
	if err != nil { /* 500 — retry is safe and wanted here */ }
	if !inserted {
		httpx.WriteOK(w, r, http.StatusOK, nil) // already processed; not an error
		return
	}

	switch event.Type {
	case "checkout.session.completed":
	case "customer.subscription.created", "customer.subscription.updated":
	case "customer.subscription.deleted":
	case "invoice.paid":
	case "invoice.payment_failed":
	default:
		// Not in the enabled_events list for this endpoint (Stripe only sends what the
		// endpoint is subscribed to), so this default is unreached in practice; kept as
		// a safety net if the Dashboard/setup config ever drifts from this switch.
	}
	httpx.WriteOK(w, r, http.StatusOK, nil)
}
```

**Why these six event types** (Stripe's own subscriptions-webhooks guide names all six among
its recommended set for "track active subscriptions" and "catch subscription status
changes" — cited at the end of this document):

| Event | What it does to entitlements |
|---|---|
| `checkout.session.completed` | First sight of a new subscription. Reads `session.Metadata` (falling back to `client_reference_id` if metadata is ever empty — defensive only, since both are always set by §2.2); if `session.Customer` is new, writes `stripe_customers`; re-fetches the Subscription fresh by `session.Subscription.ID` (never trusts the session payload's own embedded subscription summary) and calls the shared `upsertFromSubscription` below. |
| `customer.subscription.created` / `.updated` | Re-fetches the Subscription by `event.Data.Object.ID` (the same "always re-fetch, never trust the payload" rule — see below) and calls `upsertFromSubscription`. Covers plan changes made through the portal, renewals, and the trailing state settle after `checkout.session.completed`. |
| `customer.subscription.deleted` | Sets `status = 'canceled'`; sets `grace_until = current_period_end + 30 days` (the retention grace, §2.7) using the *last known* `current_period_end` rather than re-deriving one from a now-deleted subscription. |
| `invoice.paid` | The authoritative "renewed successfully" signal (Stripe's own guide: "you can provision access to your product when you receive this event and the subscription status is active"). Re-fetches the Subscription named by `invoice.Subscription.ID` and calls `upsertFromSubscription` — this is what clears a `past_due` status back to `active` after a retried card succeeds. |
| `invoice.payment_failed` | Re-fetches the Subscription and calls `upsertFromSubscription` — Stripe's own status transition (to `past_due` or `incomplete`) is trusted once re-fetched, not inferred from the invoice event alone. No immediate revocation (§1.3 rule 4); the account page shows a "payment failed, update your card" banner once `status` reads `past_due`. |

**RULING 8 — out-of-order delivery is handled by never trusting an event's own embedded
object state, only by re-fetching the named object fresh from the Stripe API inside the
handler.** Stripe's own docs are explicit that event delivery order is not guaranteed and
recommend exactly this ("you can also use the API to retrieve any missing objects... don't
use `created` to determine event order"). `upsertFromSubscription(sub *stripe.Subscription)`
is the **one** function that writes `entitlements` from Stripe state, called by every event
type above with a freshly `Get`-ed Subscription, never with the fields Stripe happened to
embed in that particular event's JSON. *Cost if wrong (trusting the payload instead):* a
`customer.subscription.updated` delivered before an earlier `.created` (Stripe's own example
ordering shows this can happen) could otherwise downgrade a status the account already
progressed past — a live, exploitable-by-accident correctness bug, not a hypothetical.

```go
// upsertFromSubscription is the only writer of entitlements from Stripe state. plan and
// subject (user_id xor guild_id) come from sub.Metadata — set at Checkout via
// subscription_data.metadata (§2.2) and therefore present on the Subscription object
// itself for the whole of its lifetime, independent of which event triggered the call.
func (s *Service) upsertFromSubscription(ctx context.Context, sub *stripe.Subscription) error
```

### 2.7 Dunning, grace, and refunds

- **Access during dunning**: `past_due` keeps every feature (§1.3 rule 4) while Stripe's own
  Smart Retries continue; the account page shows the failure and links to the portal to
  update the card (`payment_method.attached`/`invoice.paid` on a later retry clears it).
- **Cancel keeps access to period end**: a `cancel_at_period_end` subscription's `status`
  stays `active` until Stripe itself flips it at the period boundary — proposal §4.1's promise
  falls out of §1.3 rule 4 with no special-cased code.
- **The 30-day retention grace** (proposal §4.1: "2 years, then 90 days after a 30-day
  grace") is `entitlements.grace_until`, set only on `customer.subscription.deleted` (§2.6's
  table) to `current_period_end + 30 days`. The nightly job that already exists to manage
  fight-log partitions and retention (this spec does not build a new one — it hands the
  existing retention path one more input) reads: raw log / saved-sim rows older than 90 days
  are purged **unless** the owning account currently has `FeatureRetention` true (§1.3), or
  `grace_until` is set and still in the future. Once `grace_until` passes, retention falls
  back to 90 days exactly as if the account had never been premium.
- **Refunds. RULING 9 — refunds are a manual, owner-run process, not a webhook-automated
  one.** Resolving a `charge.refunded` event back to the subscription it belongs to requires
  a multi-step lookup (Charge → PaymentIntent → the invoice-payments list → Invoice →
  Subscription, per Stripe's own current-API-version guidance, cited below) that exists
  because a charge is not directly linked to a subscription on recent API versions — real,
  but disproportionate machinery for an event this rare on a $4–$15/month product. Instead:
  the refund policy (§6) and the owner's own runbook say to **cancel the subscription in the
  Stripe Dashboard in the same action as issuing the refund**; `customer.subscription.deleted`
  then does the entitlement work exactly as any other cancellation would. *Cost if wrong:* a
  refund issued without also canceling leaves the account entitled for the rest of that
  billing period — a support-process gap, not a code gap, and it costs at most one month's
  (or year's) access on an already-refunded, already rare transaction; cheap to fix by adding
  the Charge→Subscription resolution later if refund volume ever makes the manual step error-
  prone.

**Amendment, 2026-09-21 (security review response).** An independent security review found
`checkGuildCheckout` (§2.3) reading a guild's entitlement with a plain, unserialized SELECT:
two officers of the same guild checking out concurrently could each read "no active plan,"
each create a Checkout Session, and each pay. The second webhook's `upsertFromSubscription`
then silently overwrote the first officer's `stripe_subscription_id` in `entitlements`,
orphaning their subscription — it kept billing with nothing in our database pointing at it —
and §2.8's original, database-→-Stripe-only reconcile could never find it (it has no local
row to start walking from). Three independent layers close this, all landed in the same
`pay-api` branch this spec describes:

1. **A per-guild Postgres advisory lock** (`billing.Store.WithGuildLock`, mirroring the
   webhook's own `WithEventLock`) serializes §2.3's three preconditions through Checkout
   Session creation, plus a `pending_checkouts` row (`guild_id`, `user_id`,
   `stripe_session_id`, `expires_at` mirroring the Checkout Session's own expiry) written
   inside that same lock — a second caller in the same window sees the pending row and gets
   the identical `409 conflict`/`portal_hint` shape an already-active plan would answer with.
   The row is cleared when its `checkout.session.completed` webhook lands, or by
   §2.8's nightly sweep once it expires with no webhook ever landing (an abandoned checkout).
2. **`upsertFromSubscription` refuses to overwrite — except a deliberate transfer.** When it
   finds an existing row for a `(subject, plan)` already `active`/`trialing`/`past_due` and
   pointing at a *different* `stripe_subscription_id`, the row is left untouched, and the
   newcomer is recorded in a new `entitlement_anomalies` table
   (`kind = 'duplicate_subscription'`) instead — see migration `0020`'s amendment. The
   newcomer subscription is then canceled at Stripe with `cancel_at_period_end: true` through
   the `Gateway` interface (never revoked, so whoever paid for it keeps what they already
   paid for through the current period). **Carve-out, found by a second review pass against
   an earlier version of this fix:** §2.9 RULING 10's "Take over billing" handoff
   *deliberately* creates a second, different subscription for a guild that already has one,
   and intends the new subscription to become the entitled one — the checkout endpoint now
   carries `intent: "transfer"` into `metadata`/`subscription_data.metadata` alongside
   `plan`/`user_id`/`guild_id` (§2.2), and `StripeUpsert.Transfer` (read from that metadata)
   makes this one write skip the duplicate guard entirely and overwrite normally, so a
   legitimate handoff is never mistaken for the race this guard exists to catch. The write
   itself is also now serialized per subject (`billing.Store.WithGuildLock`/`WithUserLock`,
   held around the whole call to `entitlements.UpsertStripe`) — a second finding from that
   same review pass: layer 1's advisory lock only ever serializes *creating* a Checkout
   Session, not the webhook deliveries that later land for whatever subscriptions got
   created, so two deliveries for two different subscriptions on the same subject could
   previously still reach `UpsertStripe` concurrently and race its own check-then-act.
3. **`stripe-reconcile` becomes two-directional** (§2.8, revised below): the
   database-→-Stripe pass this section originally specified is unchanged, but the "logs...
   as an anomaly to investigate" promise this section made and never implemented is now real,
   through `Gateway.ListSubscriptions` and the same `entitlement_anomalies` table
   (`kind = 'orphan_subscription'`).

**Operator runbook for a `duplicate_subscription` anomaly**: query `entitlement_anomalies
where kind = 'duplicate_subscription'` (or wait for the eventual admin surface to list them —
not built in this fix). Each row names the plan, the guild or user subject, the newcomer's
`stripe_subscription_id`, and when it happened. In the Stripe Dashboard, find that
subscription (already `cancel_at_period_end` by the time it is recorded — no urgency), and
**refund the second officer's charge**, exactly as the ordinary refund runbook above already
does: refund and cancel in the same action (the newcomer is already scheduled to cancel at
period end, so "cancel" here means canceling it immediately rather than waiting out the
period, if the refund is issued before then). The kept row (`ExistingSubscriptionID` in the
anomaly's sibling `entitlements` row) needs no action — it was never touched.

### 2.8 Reconciliation

A seventh `os.Args[1]` subcommand, `stripe-reconcile`, run nightly on a Cloud Scheduler job
the same way `sim-validate` already is (`api/README.md`'s job-creation pattern), with three
passes:

1. **Database → Stripe** (the original design): every `entitlements` row with
   `source = 'stripe'` and `status` not `canceled` is re-fetched fresh by
   `stripe_subscription_id` and upserted through `upsertFromSubscription` exactly as a
   webhook event would — healing anything Stripe's own three-day retry window never
   successfully delivered (rare, but the backstop this spec's brief asks for explicitly).
2. **The `pending_checkouts` sweep** (added by the amendment below): any row whose Checkout
   Session has expired with no webhook ever landing for it — an abandoned checkout — is
   deleted, the same backstop role this job already plays for webhooks, extended to a
   session nobody ever completed.
3. **Stripe → database** (added by the amendment below): for each of our two Stripe Product
   ids (`billing.ProductIDPremium`, `billing.ProductIDGuild`), `Gateway.ListSubscriptions`
   lists every subscription Stripe currently considers live, and any with no matching
   `entitlements.stripe_subscription_id` is written to `entitlement_anomalies`
   (`kind = 'orphan_subscription'`) for a human to investigate.

**Amendment, 2026-09-21 (security review response): pass 3 is new, and pass 2 is new.**
Originally this job was one-directional (database → Stripe only, pass 1 above): Stripe's own
retry-for-three-days already covers the common case, and the reasoning against a
two-directional sync was that a Stripe subscription with no local row should not exist under
normal operation (every subscription this integration creates carries
`subscription_data.metadata`) and is a signal worth a human looking at, not something to
silently auto-repair. That reasoning still holds — pass 3 does not auto-repair anything, it
only surfaces the anomaly durably (a queryable `entitlement_anomalies` row) instead of the log
line this section originally specified and the code never actually wrote. Pass 3 is exactly
the layer that would have caught the double-billing finding's orphaned first subscription,
which pass 1 alone can never see (it has no local row to start walking from) — see §2.7's own
amendment for the full finding and the other two layers that close it, and for the operator
runbook a `duplicate_subscription` or `orphan_subscription` anomaly row calls for.

### 2.9 Who may buy the guild plan, and what happens on departure

Covered in full in §2.3 (buying) and §1.4's `GuildBillingView` (visibility). Summarized:

- **Buying**: any `VerifiedGuildRank` officer/leader of a *claimed* guild (§2.3, RULING 7).
- **The buyer leaves the guild, or is no longer an officer**: the subscription and the
  `entitlements` row are untouched — Stripe bills whoever's card is on the Customer object
  regardless of that person's current guild membership, and the guild keeps the plan's
  features for as long as the subscription stays active. `GuildBillingView.BilledBy` reads
  the departed account's battletag with no special-casing (a battletag does not become
  invalid by leaving a guild); only `YouAreBillingContact` changes (it is keyed off the
  requester, not off guild membership).
- **Transfer of billing ownership. RULING 10 — no direct Stripe-side transfer; a clean
  handoff is a second subscription, not a moved one.** Stripe does not offer an API call that
  moves a Subscription from one Customer to another. The guild settings page (owned by the
  guild membership spec, this spec adds the button and the endpoint it calls) offers "Take
  over billing" to any verified officer: it starts a **new** Checkout Session for the guild
  plan under the new officer's own Customer, exactly as `POST /v1/billing/checkout` already
  does — §2.3's "already has an active row" conflict check is bypassed only for this one
  flow, by an explicit `intent: "transfer"` field the endpoint requires alongside `guild_id`,
  which the web only ever sends from the "Take over billing" button, never from the normal
  buy flow. Copy on that button says plainly: "Cancel the old subscription first to avoid
  paying twice; ask the previous billing officer, or use the portal once you have taken over
  billing to see both." *Cost if wrong:* a guild can be briefly double-billed across a
  handoff if the old subscription is not canceled promptly — bounded to at most one billing
  period, self-correcting, and stated honestly in the button's own copy rather than hidden.
  If the departed billing officer is unreachable and nobody currently at the guild is
  `YouAreBillingContact`, canceling the orphaned old subscription is an **OWNER**-mediated,
  out-of-band task (Stripe Dashboard), not built into this version.

## 3. Security

This section is deliberately organized as a checklist against the brief's own list, so
nothing on it is silently skipped.

- **Webhook authenticity**: `webhook.ConstructEvent` with `s.WebhookSecret` (from Secret
  Manager, §3's key handling below) on every request to `/v1/billing/webhook`, before any
  other code runs against the payload (§2.6). Stripe's default replay tolerance (5 minutes
  between the signed timestamp and now) is used unmodified — the brief's "replay tolerance"
  ask is satisfied by *not* overriding the library default to something looser, and the
  library's own guidance is explicit that a tolerance of `0` is a misconfiguration, not
  extra safety, so this is stated as a deliberate non-change.
- **No trust in the browser about payment state, anywhere.** The web never sets or reads a
  local "premium" flag except what `GET /v1/me`'s server-computed `Entitlements` (§1.4) says,
  on every load, no-store (`httpx.SetPrivateListCache`-style header, matching the existing
  per-account response convention). Even the post-checkout redirect (§5's `/premium/checkout`
  success state) does not assume success from the URL alone: it re-fetches `/v1/me` and
  shows "processing" with a short bounded poll if the webhook has not landed yet (typically
  sub-second, but Checkout's redirect can race it), never "you're premium now" from the
  presence of `?status=success` by itself — that query parameter is not proof of anything
  and an attacker could hand a victim that exact URL with no effect beyond a false-looking
  loading state that a re-fetch corrects.
- **CSRF on the session-creating endpoints**: `POST /v1/billing/checkout` and
  `POST /v1/billing/portal` are both `auth.RequireSession`, mounted normally, so the existing
  double-submit `fs_csrf`/`X-CSRF-Token` check (`session.go`'s `Authenticator.Middleware`)
  already covers them with no new mechanism — the same reason §5's buy flow is designed as a
  small hydrated island making an ordinary `requestEnvelope`-style POST (`account/api.ts`'s
  existing convention) rather than a bare link: a bare `<a>`/GET to an action endpoint would
  either need CSRF exemption (wrong — this creates state, a pending Checkout Session tied to
  the signed-in account) or would not work at all under the existing CSRF rule, so it was
  never on the table as a design once the no-JS constraint on `/premium` itself (§5) was
  read correctly as "the marketing/FAQ page ships no JS," not "the whole buy flow does."
- **Open-redirect safety of return URLs**: `success_url`/`cancel_url` (§2.2) are built
  server-side from `cfg.PublicBaseURL` alone, never from request input. The one place a
  redirect target *does* come from a query string — `/login?next=...` when a signed-out
  visitor clicks a plan (§5) — reuses `auth.safeNext` (`bnet.go:142`) unmodified: it already
  enforces "a same-site path, with no scheme and no host" for every existing `next` consumer
  (`bnetCallback`, `emailCallback`), and this spec adds no second `next`-like parameter that
  would need its own validation.
- **Rate limits**: `POST /v1/billing/checkout` and `POST /v1/billing/portal` each get
  `httpx.RateLimitPer(20, time.Hour, hops)`, mounted the same way `POST /v1/auth/email`
  scopes its own limiter today (`auth.Mount`'s `emailLimit`) — generous for a signed-in
  member retrying a declined card, bounding a script that would otherwise mint many Stripe
  Checkout/Portal Sessions for no cost to the attacker but real API-quota and Dashboard-
  clutter cost to the account. `POST /v1/billing/webhook` gets **no additional per-route
  limiter beyond the router's existing per-IP budget** (`RateLimitExcept`, 120/min,
  `server.go:57`) — considered and rejected adding a carve-out: Stripe delivers webhooks to
  this endpoint from its own origin, not proxied through shared infrastructure the way
  browser traffic through Cloudflare is, so this account's own delivery volume (in the tens
  per day at this product's scale, bursting briefly around a renewal wave) never approaches
  120/minute; `maxWebhookBody` (§2.6) bounds the one other resource a flood could exhaust.
- **No card data ever touches our servers.** Every payment surface is Stripe-hosted
  (Checkout, the Customer Portal) — nothing in this design collects, transmits, proxies, or
  logs a card number, CVC, or expiry anywhere in this codebase; the request/response bodies
  this service ever handles for billing are Stripe object ids, metadata, and status strings.
- **Secrets**: `STRIPE_SECRET_KEY` and `STRIPE_WEBHOOK_SECRET` are added to
  `config.Config`/`config.Load` exactly like `RESEND_API_KEY` (`config.go`'s existing
  required-env pattern) and to the required-secrets set in `api/README.md`'s "First-time
  setup" (`gcloud secrets create`, an IAM binding for `api-runtime`, then one
  `gcloud run services update api --update-secrets STRIPE_SECRET_KEY=...,
  STRIPE_WEBHOOK_SECRET=...` — never a workflow-file change, per §0). **`STRIPE_PUBLISHABLE_KEY`
  is not a secret** — Stripe's own docs name it explicitly as "safe to expose" — and this
  design in fact never needs it server-side or client-side at all, because every payment
  surface is Stripe-hosted (Checkout Sessions, the Portal): there is no Stripe.js/Elements
  embed anywhere in this spec that would need a publishable key. If that changes later
  (e.g. an embedded Checkout instead of hosted), it ships as a build-time public env var in
  `web/`, never through Secret Manager. Neither secret is ever logged: `slog` call sites in
  this spec's handlers log Stripe object ids (`sub.ID`, `event.ID`) and status strings, never
  a request/response body that could carry a key back out through a log line (Stripe's own
  API responses never echo the key used to make the call, so this is a "never construct a
  log line from the raw config/secret values" discipline rather than a redaction problem).
- **Least-privilege restricted API key.** **OWNER:** create a Restricted API key (`rk_...`,
  not the unrestricted `sk_...` secret key — Stripe's own current guidance recommends RAKs
  for all new integrations) scoped to exactly: write access to Checkout Sessions, Billing
  Portal Sessions, and Webhook Endpoints (setup only); read+write to Customers, Subscriptions,
  and Prices/Products (the last two only for `stripe-setup`, §2.1); read access to Invoices
  and Charges (for `stripe-reconcile` and any future refund tooling); no access to Payouts,
  Balance, Transfers, or anything Connect-related, none of which this integration ever
  touches. This key is `STRIPE_SECRET_KEY` above — the name is kept for parity with Stripe's
  own docs and env-var conventions, even though the value is a restricted key, not an
  unrestricted secret key.
- **Test mode vs. live mode isolation.** A new `config.Config.StripeEnvironment` field
  (`"test" | "live"`, from `STRIPE_ENVIRONMENT`) plus a startup check in `main.go`'s `start`
  (alongside the existing required-env validation): the key's own prefix is authoritative,
  not trusted config — `sk_test_.../rk_test_...` in an environment declared `live`, or
  `sk_live_.../rk_live_...` in an environment declared anything but `live`, refuses to start
  with a clear error, the same "fail fast at startup" pattern `config.Load` already uses for
  a missing required var. This is the concrete mechanism behind the brief's "a startup check
  that refuses a live key in a non-production environment and a test key in production."
  `PublicBaseURL`'s existing scheme check (`strings.HasPrefix(cfg.PublicBaseURL, "https://")`,
  already read for `Authenticator.Secure`) is reused as the signal for "is this production,"
  rather than inventing a second environment name — `STRIPE_ENVIRONMENT=live` is required
  whenever `PublicBaseURL` is `https://foreversixty.gg` and refused everywhere else.
- **PII minimization**: `stripe_customers` stores a Stripe Customer id and nothing else —
  no address, no card details, no tax id (Stripe holds all of that; §2.5's portal is where a
  member manages it directly with Stripe, never round-tripped through this database). The
  only new PII this spec introduces into Postgres at all is the Customer id mapping itself.
- **Audit logging of entitlement changes**: `entitlement_audit` (§1.1), written by every
  path that changes an `entitlements` row — `upsertFromSubscription`, the cancellation path,
  and both CLI grant/revoke commands — with `old_status`/`new_status` and a human-readable
  `actor` string, queryable by a support/debugging session with no separate tool.
- **Abuse: one guild plan shared by an inflated "guild."** Membership itself, and its
  verification, are entirely the guild membership spec's concern (§0's assumption); this
  spec's own exposure is narrower — could an attacker make an *unrelated* account benefit
  from a guild's plan cheaply? RULING 2 (§1.3.1) already closes the cheap version: an
  unverified `guild_characters` row confers no guild-plan benefit at all, so inflating a
  guild's roster with forged exports — the account's characters "named" into the guild but
  never actually raiding with it, never officer-approved, never invited — buys the attacker
  nothing; the only way an account's `guild_members.verified_at` is ever set is one of its
  characters clearing a real bar (two raid nights in the guild's own logs, officer approval,
  or an invite redemption), each of which costs something a bulk-forged roster cannot supply
  at scale. The remaining question is scale among *genuine* members: is there a cap on
  legitimate guild size? **RULING 11 — no hard member cap, because raid guild rosters in
  this game are naturally bounded (a 40-player raid plus alts and inactive members tops out
  in the low hundreds) and the plan's price ($15/month covering "every member") is already
  priced assuming a full raiding guild, not designed to be undercut by a large roster** —
  the proposal's own math ("under 40 cents a head" for 40 members) only gets cheaper per
  head as a guild grows, which is a real economic exposure if guild size were ever
  unbounded and gameable (e.g. an alliance of unrelated players registering one shared
  "guild" purely to split a subscription forty ways). *Cost if wrong:* if this becomes a real
  pattern post-launch (observable as guild rosters with implausible member/character-class
  distributions, or many members whose only site activity is a single addon export naming
  the guild and nothing else — no logs, no reports), the cheapest fix is a soft cap read at
  checkout time (`guild_members` count with `verified_at is not null` above, say, 60 refuses
  the purchase with a message to contact support) added to §2.3's existing precondition list
  with no schema change — deferred rather than built now because there is no guild roster
  data yet to calibrate a threshold against, and an uncalibrated cap risks refusing a
  legitimate large guild on day one, which is the worse failure mode to ship with.

## 4. API endpoints

All in the existing envelope (`httpx.Envelope`); errors use `WriteError`'s `(status, code,
message, fields)` shape throughout.

| Method & path | Auth | Request | Response (`data`) | Errors |
|---|---|---|---|---|
| `POST /v1/billing/checkout` | session | `{plan: "premium"\|"guild", interval: "monthly"\|"yearly", guild_id?: number, intent?: "transfer"}` | `{checkout_url: string}` | 400 `invalid` (bad plan/interval, or `guild_id` missing for `plan:"guild"`); 403 `forbidden` (not a verified officer/leader of `guild_id`, or the guild is unclaimed); 409 `conflict` (`fields.portal_hint = "1"` — guild already on the plan and `intent` is not `"transfer"`); 429 |
| `POST /v1/billing/portal` | session | `{guild_id?: number}` (omitted for the caller's own personal billing) | `{portal_url: string}` | 404 `not_found` (no `stripe_customers` row for the caller, or — with `guild_id` — the caller is not `YouAreBillingContact` for that guild); 429 |
| `POST /v1/billing/webhook` | none (Stripe signature) | raw Stripe event body | `null` | 400 `invalid` (bad/missing signature, or body over `maxWebhookBody`) |

Reads are not a separate route: `GET /v1/me`'s `Entitlements`/`Guilds[].Plan` (§1.4) is the
one place the web asks "what am I entitled to" — a dedicated `GET /v1/billing` would be a
second source of truth for exactly the data `/v1/me` already carries on every page load.

Admin grant/revoke is the CLI (§1.5), not an HTTP endpoint (RULING 4).

## 5. Web

### 5.1 `/premium`: the page itself ships no client JavaScript

**RULING 12 — `/premium` is `session=false` like `/about` (Base.astro's own contract,
`home.spec.ts`'s "content pages ship no client JavaScript" pinning `/about` today, extended
here to `/premium`), and the plan "Subscribe" controls are plain same-origin links to a
second, tiny page that *does* hydrate.** Reconciled as follows, matching the site's existing
divide between a pure content page and a page that only exists to run one action (`/login`
is exactly this today — `session` per `login.astro:11`, `Account.svelte`'s `login` mode does
the real work):

- `/premium` (new `.astro` page, not the generic `pages` content collection — it needs the
  monthly/yearly toggle and plan-specific hrefs a Markdown body cannot carry) is static
  markup: what is free forever, the two plans, the FAQ (§5.2). It links each plan/interval
  combination as a plain `<a href="/premium/checkout?plan=premium&interval=monthly">` styled
  as a secondary button (`design/DESIGN-SYSTEM.md`: "no marketing buttons, secondary buttons
  only" — the buy control is styled identically to every other secondary button on the site,
  never a colored/urgent call-to-action).
- **The monthly/yearly toggle is CSS-only**, the same no-JS-interactivity idiom the header's
  Reference disclosure already uses (native `<details>`, per `one-product-design.md` §3):
  two anchors, `/premium#monthly` and `/premium#yearly`, each targeting a `:target`-shown
  panel (`#monthly:target ~ .yearly-panel { display: none }` and the mirror), defaulting to
  monthly when there is no fragment. No JavaScript, no layout shift (both panels reserve the
  same height; only visibility toggles).
- `/premium/checkout` (new page, `session` — it hydrates, exactly like `/login`) reads its
  own query params (`plan`, `interval`, `guild_id?`) and mounts one small island,
  `CheckoutLauncher.svelte`: on mount, `fetchMeOnce()` — signed out, it immediately
  `window.location.assign` to `/login?next=` + this exact URL (`auth.safeNext`-safe, §3),
  so sign-in returns straight back and fires; signed in, it immediately `POST
  /v1/billing/checkout` via the existing `requestEnvelope` (`account/api.ts`), CSRF header
  included automatically the way every other state-changing call in this codebase already
  gets it, then `window.location.assign(result.checkout_url)`. The page's only visible
  content, the whole time, is "Redirecting to secure checkout…" with the same error-line
  convention (`min-h-[21px]`, `role="alert"`) `Account.svelte` uses everywhere else, for the
  rare case the POST itself fails (e.g. the 409 conflict, §4 — shown as "This guild is
  already on the plan. Manage its billing from guild settings." with a link, not a raw error).
  A `?status=success&session_id=...` return from Stripe (§2.2) is handled by the same page:
  it re-fetches `/v1/me` (never trusts the query string, §3) and shows "You're all set" once
  `Entitlements` confirms it, or keeps a bounded poll (a few seconds) with "This can take a
  moment; refresh if it doesn't update" as the honest fallback.
- The guild plan's buy flow starts from the guild settings page (`GuildSettings.svelte`, the
  guild membership spec's file), not from `/premium` at all: an officer viewing settings for
  a guild not yet on the plan sees a "Subscribe the guild" secondary button linking to
  `/premium/checkout?plan=guild&interval=monthly&guild_id=<id>` (interval chosen the same
  CSS-toggle way, mirrored on that page) — `/premium` itself never needs to know about a
  specific guild id, keeping this spec's and the guild membership spec's files from touching
  the same component.

### 5.2 `/premium` content, in full

Structure, top to bottom (design system: dark, Cinzel section titles, Barlow body, no
emoji, no exclamation marks, honest copy):

1. **What is free, forever** — a short list, verbatim from the proposal's own principle:
   "Planner, browser simulator, logs, live logging, rankings, the addon — always free, no
   ads. Nothing that comes from Battle.net is ever behind a paywall — your characters and
   your sign-in stay free no matter what." One line under it: "A player who never pays never
   hits a wall that makes the site feel broken."
2. **The two plans**, monthly/yearly toggle (§5.1), each as a plain feature list (not a
   comparison-table gimmick — the tiers table in this spec's §1.3.1 is the source, rendered
   as prose per plan):
   - **Premium — $4/month or $40/year.** "Run sims on our servers (5,000 combinations, any
     precision, no tab left open). Two years of log retention instead of ninety days.
     Compare more than two builds at once. Character history charts over time. Discord
     notifications. A supporter mark on your profile."
   - **Guild — $15/month or $150/year, covers every member.** "Everything in Premium, for
     everyone in the guild. Officer tools once they ship: the loot council helper, the raid
     readiness board, and the rest — a claimed guild's officers see what's live today on the
     guild's settings page." (Honest about sequencing — the proposal's own §4.3 ships officer
     tools after the plan itself; this line is written to still be true on day one, when the
     guild plan first goes on sale with only some officer tools live.)
3. **FAQ**, written out in full (copy the implementation ships verbatim, not a paraphrase):
   - **"What happens if I cancel?"** "You keep everything through the end of the period you
     already paid for. After that, the account works exactly like a free one — nothing is
     deleted."
   - **"What happens to my long-retention logs if I stop?"** "They stay at the two-year
     retention for thirty days after your subscription ends, in case you resubscribe. After
     that, retention falls back to ninety days like every free account, and anything older
     than that is not kept."
   - **"Can I get a refund?"** "Email us and we'll sort it out — see the refund policy for
     the details." (links to `/refunds`, §6)
   - **"Is anything from Battle.net ever paid?"** "No. Your characters, your sign-in, and
     everything the game itself tells us about you stays free, always — Blizzard's own rules
     for using their data require this, and we'd want it that way regardless."
   - **"Are there ads?"** "No. Not for anyone, paying or not. The site is paid for by
     Premium and the guild plan, not by ads, and that isn't a perk you're buying — it's true
     for everyone."
   - **"Who am I paying?"** "COMMISH LLC, the company behind Forever Sixty. Stripe handles
     the payment; we never see or store your card."

### 5.3 Every existing premium control, updated

- `web/src/lib/account/api.ts`: `Me.user.premium: boolean` is removed; `Me` gains
  `entitlements: EntitlementsView` mirroring §1.4's Go shape (a hand-written TS interface,
  matching how `Me`/`MeGuild`/`MeCharacter` are already hand-mirrored here, not generated).
- Every current reader of `user.premium` switches to `entitlements.server_sims` for the
  identical boolean it used to read — no behavior change, just the field it comes from:
  `web/src/components/sim/SimView.svelte:355` and `web/src/components/sim/tools/
  ToolsView.svelte:106`'s `store.setPremium(result?.user.premium === true)` become
  `store.setPremium(result?.entitlements.server_sims === true)`; `web/src/lib/sim/
  store.svelte.ts`'s `premium` state (line 171) and `bulk-store.svelte.ts`'s own copy
  (line 144) are unrenamed (the sim module's internal vocabulary stays "premium" — it is
  gating one specific feature, `server_sims`, and renaming every internal identifier in a
  package this spec does not otherwise touch is unnecessary churn).
- `sim/copy.ts`'s `premiumRequired` (139), `capPremium` (985) and `capPremiumNote` (986)
  gain one more line each, a link: `premiumRequired`'s 402 message and the button's own note
  both point to `/premium` (`capPremiumNote`: "Premium lifts the limit to 5,000 combinations
  and any precision. <a href="/premium">See plans</a>.") — the "every premium control links
  to `/premium`" requirement, satisfied at the two places gating server sims lives today.
  `RunControl.svelte`'s `{#if premium}` branch (line 187) additionally renders a "Get
  premium" secondary-button link to `/premium` in the `{:else}` branch it does not currently
  have, for a signed-in, non-premium visitor specifically (a signed-out one sees the existing
  sign-in prompt first, unchanged).
- The **supporter mark**: a public `entitlements.IsSupporter` (§1.3) call from whatever
  handler already serves a character page/card and a guild page (the `rankings` package —
  this spec adds the call, not new rendering conventions; the mark itself is a small badge
  matching the "This site" source-pill's gold treatment per `DESIGN-SYSTEM.md`, placed beside
  the character/guild name, with a one-word label "Supporter" and no further explanation
  inline — the premium page is where "what supporter means" lives, per the honest-copy rule
  against repeating an explanation everywhere a fact appears).

### 5.4 Account page billing block

`Account.svelte`'s `account` mode (`web/src/components/Account.svelte`) gains one section,
placed after "Account" and before "Devices," modelled on the existing `onAnonymize` pattern
(optimistic local state, same error-line convention):

- No personal subscription (`entitlements.billing === null`): "Not on Premium. <a
  href="/premium">See plans</a>." — plain text, no upsell tone (house style).
  A guild plan the account benefits from (but does not bill) is **not** shown here at all —
  it belongs on that guild's own settings page, not the personal account page, keeping this
  block about the account's own billing only.
- A personal subscription exists: plan name, `current_period_end` formatted as a date,
  "Renews" or "Ends" depending on `cancel_at_period_end`, and a "Manage billing" secondary
  button calling `POST /v1/billing/portal` (no `guild_id`) and redirecting to the returned
  `portal_url` — the one new state-changing call this component makes, following its
  existing `run()` wrapper (busy/error handling already shared by every other action here).

## 6. Legal pages

Three new entries in the `pages` content collection (`web/src/content/pages/`, the same
collection `about.md`/`sources.md` already use — `[...slug].astro` and `Content.astro`
render them with no code change needed beyond the new Markdown files and their frontmatter),
each with the `confidence: confirmed` / `sources: [{label: "This site", ...}]` frontmatter
`about.md` uses, since these are the site's own policy, not a claim about the game. Every
fact the owner alone can supply is marked `OWNER:` below; none of it is invented here.

### 6.1 `/terms` — Subscription terms

Outline: what Premium and the guild plan are (server compute and storage, never data);
billing cadence and auto-renewal; how cancellation works (§5.2's FAQ answer, restated
formally); price changes (existing subscribers keep their price until they voluntarily
change plans — **OWNER:** confirm this is the intended policy, or state the alternative);
account termination for abuse (ties to §3's abuse ruling); that the service is provided by
**OWNER: COMMISH LLC's exact registered name and state of formation**; **OWNER: governing
law and venue**; **OWNER: support email**; a note that Blizzard-sourced data and features are
never sold (ties this page back to the proposal's core promise, in the one place a legal
document states it).

### 6.2 `/privacy` — Privacy policy

Outline: what is collected (Battle.net sign-in identity, email for magic-link sign-in,
uploaded combat logs and addon exports, a Stripe Customer id — explicitly **not** card
details, address, or tax id, which Stripe holds directly per §3's PII minimization); how
Battle.net data is used and the 30-day retention rule Blizzard's Developer API Terms of Use
impose on anything sourced from their API (cited plainly, since it is also this site's own
reason nothing from Battle.net is ever paid); Stripe named as the payment processor, with a
link to Stripe's own privacy policy for what Stripe itself collects during Checkout/the
Portal; cookies (`fs_session`, `fs_csrf` — functional only, no tracking/advertising cookies,
consistent with "no ads anywhere, ever"); data retention generally (ties to the 90-day/2-year/
30-day-grace rule, §2.7); **OWNER: support/contact email for a data request**; **OWNER:
whether a formal DPA or GDPR/CCPA-specific section is needed for the site's actual audience**.

### 6.3 `/refunds` — Refund policy

Outline: the plain-language promise from §5.2's FAQ, expanded: how to request one (**OWNER:
support email**), what happens to entitlements on a refund (§2.7's RULING 9 — the owner
cancels the subscription as part of processing the refund; the policy says access ends at
that point, not "immediately upon refund," since the two are the same action taken together
by the owner, and stating it this way keeps the page honest about how it actually works
rather than promising an automated instant cutoff this design does not build); no
partial-period refunds beyond the owner's own discretion (**OWNER: confirm, or state the
actual policy**); COMMISH LLC named again as the entity processing the refund.

### 6.4 Footer and premium-page links

`Footer.astro`'s `footerLinks` (currently `About`, `Sources`, `Changelog`, `Contribute`)
gains `Premium`, `Terms`, `Privacy`, `Refunds` — eight entries total, each already clearing
the 44px hit-target floor on text length alone except possibly `Terms`/`Privacy` on the
narrowest phones, checked against the same rule `About` needed `px-2` for (`Footer.astro`'s
existing `label === 'About' && 'px-2'` becomes a small set membership check rather than a
single equality, no other change to that component's structure).

## 7. Testing

- **Unit** (`api/internal/entitlements/`): `Can` for every `(feature, entitlement state)`
  combination in §1.3's table, including the guild-verification boundary (RULING 2: a
  `guild_members` row with `verified_at = null` must answer `ReasonNoPlan`/false exactly like
  no row at all — the regression test the guild membership spec already writes for report
  visibility, mirrored here for entitlements); `upsertFromSubscription` against hand-built
  `stripe.Subscription` fixtures for each status Stripe defines (§2.6's table), asserting the
  resulting row and one `entitlement_audit` write per call; the migration round-trip
  (`0020` up/down against a seeded `0005`+`0018` database, and the backfill query against a
  seeded `users.premium = true` row, following the existing migration-test harness pattern
  `db_test.go` already uses for `0013`–`0017`).
- **Handler tests, with signed fixture webhooks**: a small helper that builds a raw JSON
  payload and computes a valid `Stripe-Signature` header against a fixture signing secret
  (Stripe's own manual-verification steps, §2.6's citation, are simple enough to replicate in
  a test helper with no network call — HMAC-SHA256 over `"{timestamp}.{payload}"`), covering:
  a first delivery is processed and audited; a redelivery of the same event id is a no-op
  200 (idempotency); a bad signature is `400`; a payload over `maxWebhookBody` is refused
  before signature verification even runs (the size limit is on the reader, not a check
  after decode); each of the six event types drives the right `upsertFromSubscription` call
  with a **stubbed** Stripe client (never a real network call in a unit/handler test — the
  re-fetch-by-id call is an interface this package defines and the test fakes, the same
  `Premiumer`/`fakePremium` pattern `sims` already uses for its own single external check).
- **Stripe test-mode end-to-end checklist**, run once by a human before launch and after any
  change to the checkout/webhook code, against `stripe listen --forward-to
  <local>/v1/billing/webhook` (or the deployed test-mode endpoint) and Stripe's documented
  test cards:
  1. A premium monthly purchase with `4242 4242 4242 4242` completes; `/v1/me` shows
     `server_sims: true` within the bounded poll (§5.1); the account page shows the plan.
  2. The SCA-required test card (`4000 0025 0000 3155`) completes the additional
     authentication step and still lands as `active`.
  3. A forced decline card (`4000 0000 0000 0002`) leaves the account with no entitlement
     and the checkout launcher shows the failure, not a false "processing" state.
  4. A simulated failed renewal (`stripe trigger invoice.payment_failed` against an existing
     test subscription) flips the account to `past_due` and the account page shows the
     banner; access is unchanged.
  5. Canceling through the portal sets `cancel_at_period_end`; the account still shows
     `server_sims: true` until the period end shown on the account page; after
     `stripe trigger customer.subscription.deleted` (simulating the actual boundary),
     access is revoked and `grace_until` is set.
  6. A guild plan purchase, by a verified officer fixture account, attaches to the right
     `guild_id`; a second guild-plan checkout attempt for the same guild gets the `409`
     with `portal_hint`; a non-officer fixture account attempting the same purchase gets
     `403`.
- **Web e2e, API stubbed** (`web/tests/e2e/`, `test-support/`'s existing fixture-API
  convention — `sim-api.ts`'s `setPremium` is the pattern this section's new stub follows):
  `/premium` renders with no `<script>` request (the exact assertion `home.spec.ts` already
  makes for `/about`, extended to this path); the monthly/yearly `:target` toggle works with
  JavaScript disabled (Playwright can assert this directly); clicking a plan while signed out
  round-trips through `/login?next=...` back to `/premium/checkout` and fires the stubbed
  checkout call; clicking while signed in goes straight to the stubbed
  `checkout_url` redirect; the guild "Subscribe the guild" button (from a fixture officer
  account) reaches the same launcher with `guild_id` set; the account page's billing block
  renders each state (`entitlements.billing` null, active, `cancel_at_period_end`, `past_due`)
  from a fixture `/v1/me` response, exercising every render branch without a real Stripe call.
- **A regression test per security rule in §3**: webhook signature required (bad signature
  → 400, no row written); no browser-side trust (a fixture `/v1/me` returning
  `server_sims: false` is asserted to hide the server-sims button even if a stale client-side
  state object were somehow set — i.e. the render is a pure function of the last `/v1/me`
  response, not of any locally-mutated flag); CSRF required on both billing POST routes
  (a request missing `X-CSRF-Token` gets `403 csrf`, the existing router-level test already
  covers the mechanism — this is one more route added to that table-driven test, not new
  logic); `next`/redirect safety (a `/login?next=https://evil.example` is rejected by the
  existing `safeNext` test, extended with one more case at the call site this spec adds if it
  is not already parameterized broadly enough to cover it); the live/test key startup check
  (§3) refuses to start under each of the two wrong-key-for-environment combinations.

## 8. Lane split

Two lanes, consistent with the proposal's own "about one lane, API-heavy" sizing for this
piece (§4.3.2) and split the same way the guild membership spec's own lanes are: by
repository half, not by feature slice, so file ownership never overlaps.

| Lane | Owns |
|---|---|
| **API + Stripe** | `api/internal/db/migrations/0020_entitlements*`, `0021_drop_users_premium*`; new `api/internal/entitlements/` package entirely (`Can`, `IsSupporter`, `Store`, `*_test.go`); new `api/internal/billing/` package (checkout, portal, webhook handlers, `upsertFromSubscription`, Stripe client wiring, `*_test.go`); `api/internal/auth/store.go` (`Premium` method and `User.Premium` field removed; `Guilds` query gains nothing new — `GuildBillingView` is assembled in `auth.Service.me`, not the store); `api/internal/auth/handler.go` (`Me`/`EntitlementsView`/`BillingView`/`GuildBillingView`); `api/internal/sims/run.go` + `handler.go` (`Premiumer` → the new `Can`-based interface); `api/internal/config/config.go` (`STRIPE_SECRET_KEY`, `STRIPE_WEBHOOK_SECRET`, `STRIPE_ENVIRONMENT`, the startup key-prefix check); `api/cmd/api/main.go` (`stripe-setup`, `stripe-reconcile`, `grant`, `revoke` subcommands); `api/internal/db/db_test.go` (the `users.premium` schema-pin test removed once `0021` lands, an `entitlements`/`stripe_customers`/`stripe_events`/`entitlement_audit` one added in its place); `api/README.md` (the new secrets in "First-time setup," the two new scheduled jobs) |
| **Web** | `web/src/pages/premium.astro`, `web/src/pages/premium/checkout.astro`; new `web/src/components/CheckoutLauncher.svelte`; `web/src/content/pages/terms.md`, `privacy.md`, `refunds.md`; `web/src/lib/account/api.ts` (`Me.entitlements`, `Me.user.premium` removed); every call site listed in §5.3 (`SimView.svelte`, `ToolsView.svelte`, `store.svelte.ts`, `bulk-store.svelte.ts`, `RunControl.svelte`, `BulkRunBar.svelte`, `sim/copy.ts`); `web/src/components/Account.svelte` (billing block); `web/src/components/Footer.astro` (new links); `web/tests/e2e/` (new specs, `home.spec.ts`'s no-JS assertion extended); `web/src/test-support/` (a billing stub alongside the existing sim-api one) |

**Order constraints**: API before Web for anything Web reads (`Me.entitlements` shape) —
Web may build `CheckoutLauncher.svelte` against the §1.4/§4 shapes ahead of the API landing,
the same allowance both `one-product-design.md` and the guild membership spec already give
their own web lanes. Within the API lane, `0020_entitlements` lands before the `entitlements`
package's own tests can run against a real database, and `0021_drop_users_premium` lands only
after a deploy running the new `Can`-based code has been confirmed healthy — not the same PR,
not the same day.

**What only the owner can do** (none of it belongs in chat, per the standing rule):

- Create the Stripe account under **COMMISH LLC** — EIN, bank account, a support email, and a
  public business address or registered agent address for receipts.
- Register for Stripe Tax in whatever jurisdictions require it once real revenue starts.
- Configure the Customer Portal's feature set in the Stripe Dashboard (§2.5), once per mode.
- Create the Restricted API key with the exact scopes in §3, in both test and live mode.
- Put `STRIPE_SECRET_KEY` and `STRIPE_WEBHOOK_SECRET` into Google Secret Manager and grant
  `api-runtime` access — never pasted into chat, a terminal session the owner runs alone.
- Fill in every `OWNER:` marker in §6's legal pages before `/terms`, `/privacy`, `/refunds`
  go live (the pages render with a placeholder-obviously-unfinished state until then — an
  Astro content-collection frontmatter field, `draft: true`, excluded from `getStaticPaths`
  the same way an unfinished page would be, is the concrete mechanism, not a promise).
- Decide and set the statement descriptor (`FOREVERSIXTY`, per the proposal) in the Stripe
  Dashboard's business settings.

## 9. Out of scope, recorded

Officer tools themselves (proposal 3.3 — the loot council helper, readiness board,
attendance/performance, raid-night comparison, assignments: this spec builds the
`FeatureOfficerViews` gate they will all check, not the tools). Ratings and the performance
analyzer (3.4/3.5/3.6). In-game features. Gifting a subscription to another account.
Regional pricing (Stripe Tax handles tax, not price localization — Adaptive Pricing, if ever
wanted, is a separate, later decision). Coupons beyond Stripe promotion codes (RULING 6
turns codes on; issuing any is an owner action in the Dashboard, not a build item).

---

## Sources consulted

- Webhook signature verification, raw body handling, and troubleshooting:
  [Resolve webhook signature verification errors](https://docs.stripe.com/webhooks/signature),
  [Receive Stripe events in your webhook endpoint](https://docs.stripe.com/webhooks) (Go
  `ConstructEvent` example, `MaxBytesReader` pattern, default 5-minute replay tolerance,
  event-ordering and idempotency guidance, CSRF-exemption note for the webhook route).
- Recommended subscription webhook events and status handling:
  [Using webhooks with subscriptions](https://docs.stripe.com/billing/subscriptions/webhooks)
  (event table, "track active subscriptions," "catch subscription status changes," the
  refund/dispute Charge→Subscription resolution steps cited in §2.7's RULING 9).
- Checkout Session parameters for subscriptions:
  [Create a Checkout Session (API reference)](https://docs.stripe.com/api/checkout/sessions/create)
  (`mode`, `line_items`, `customer`, `client_reference_id`, `metadata`, `automatic_tax`,
  `allow_promotion_codes`, `success_url`/`cancel_url`).
- Customer Portal configuration and session creation:
  [Integrate the customer portal with the API](https://docs.stripe.com/customer-management/integrate-customer-portal)
  (configure-then-integrate ordering, `billing_portal/sessions` creation, portal webhooks).
- Restricted API keys and key management:
  [API keys](https://docs.stripe.com/keys) (RAK vs. secret key guidance, test/live mode
  isolation, secrets-vault best practice).
- Price `lookup_key`: [The Price object (API reference)](https://docs.stripe.com/api/prices/object)
  (`lookup_key` field definition and uniqueness).
