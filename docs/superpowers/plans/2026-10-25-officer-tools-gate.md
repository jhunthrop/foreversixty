# Officer Tools Gate Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Gate Award, Undo and Nudge behind the guild plan, with an officer-tools-only grant kind that every guild claimed before the gate receives through the first raid night.

**Architecture:** One `features text[]` column on `entitlements` scopes a guild grant to `officer_views`; one SQL predicate makes grants expire; a per-guild `CanForGuild` is the handler gate. The web learns the state from fields the page already fetches (`home.viewer.officer_tools`, `me.guilds[].plan.scope`) and draws the designer's `officer-read-only` state. Two PRs in one lane: PR A (Tasks 1–4) ships the grant kind and the fields; the grandfather run happens from the deployed sha; PR B (Tasks 5–10) ships the gate.

**Tech Stack:** Go (pgx, migrations in `api/internal/db/migrations`), Astro + Svelte 5, vitest, Playwright.

**Spec:** `docs/superpowers/specs/2026-10-25-officer-tools-gate-design.md` (product) and `design/specs/2026-10-25-officer-tools-gate.md` (experience). Both argue from `docs/superpowers/specs/2026-10-24-premium-experience-design.md` §4.7.

## Global Constraints

- Only the deployed API runs migrations; the `grant` and `grandfather-officer-tools` commands migrate at startup, so they run only from the deployed sha, with `DATABASE_URL` read from Secret Manager inside the command and never printed.
- A scoped grant gives members nothing: `Can` for every other feature, `IsSupporter` and the Supporter chip ignore it.
- `status in ('active','trialing','past_due') and (source <> 'grant' or current_period_end is null or current_period_end > now())` is the one active predicate; no Go caller re-derives it from `status` alone.
- Every new player-facing string is in the experience spec §7 and tested verbatim against it.
- Members who are not officers see nothing new. Controls are disabled, never hidden (tenet 4).
- The gate (Task 5 onward) merges only after the grandfather run is verified in production.
- Commit trailer `Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>`; `git commit -m`, never a heredoc.

## Review Focus

1. A personal Premium grant whose `--until` passed: `Can(server_sims)` must now be false (the predicate change reaches personal grants). Pinned in Task 1's store test.
2. A guild whose Stripe plan is `past_due` with `current_period_end` in the past: officers keep Award (Stripe owns that row's life). Pinned in Task 1.
3. A guild on a scoped grant buys the plan: checkout is allowed, the webhook replaces the row with `features = null`, members gain Premium. Pinned in Task 2.
4. The 403 for a non-officer (`forbidden`) and the 403 for a lapsed grant (`guild_plan_required`) must stay distinct on the web; the web never branches on status alone. Pinned in Task 6 (error class) and Task 7 (component).
5. The home endpoint for a signed-out or non-member viewer must carry `officer_tools: false` and never error on the extra query. Pinned in Task 3.

---

## PR A — the grant kind

### Task 1: Migration 0049, the active predicate, scoped reads, `CanForGuild`

**Files:**
- Create: `api/internal/db/migrations/0049_entitlement_features.up.sql`, `0049_entitlement_features.down.sql`
- Modify: `api/internal/entitlements/store.go` (lines 22–120 `activeStatusClause`, `hasVerifiedGuildPlan`, `canOfficerViews`; 160–185 `GuildBilling`), `api/internal/entitlements/entitlements.go` (Feature docs)
- Test: `api/internal/entitlements/store_test.go`

**Interfaces:**
- Produces: `const activeClause` (SQL, replaces every `activeStatusClause` use), `func (s *Store) CanForGuild(ctx, guildID, userID int64, feature Feature) (bool, Reason, error)`, `GuildBilling{…, Features []string, Active bool}` with `func (g GuildBilling) Scoped() bool` (true when `Features` is non-empty), `const FeatureScopeOfficerViews = "officer_views"` (the only value `features` may hold).

- [ ] **Step 1: Write the migration**

`0049_entitlement_features.up.sql`:
```sql
-- A guild grant may cover officer tools alone (premium spec §4.7; design 2026-10-25 §4.1).
-- null = the whole plan (every row before this migration); '{officer_views}' = that feature only.
alter table entitlements add column features text[];
alter table entitlements add constraint entitlements_features_scope check (
  features is null
  or (plan = 'guild' and source = 'grant' and features <@ array['officer_views']::text[])
);
```
Down: drop the constraint, drop the column.

- [ ] **Step 2: Write the failing store tests** (extend `store_test.go`; `seedEntitlement` gains a variant `seedEntitlementUntil(t, pool, subj, plan, status, features []string, until *time.Time)` inserting `source='grant'`)

```go
func TestExpiredGrantIsOff(t *testing.T)              // personal premium grant, until = yesterday → Can(FeatureServerSims) false, ReasonNoPlan; IsSupporter false
func TestPastDueStripeRowStaysOn(t *testing.T)        // seed source='stripe', status='past_due', current_period_end yesterday → officer Can(FeatureOfficerViews) true
func TestScopedGrantGivesMembersNothing(t *testing.T) // guild grant features={officer_views}: verified member Can(FeatureServerSims) false, IsSupporter false; verified officer Can(FeatureOfficerViews) true
func TestCanForGuild(t *testing.T)                    // table: officer+scoped grant → true/entitled; officer+expired scoped grant → false/no_plan; member+grant → false/not_officer; officer of another guild → false/no_plan; unverified officer → false/no_plan; feature != officer_views → error
func TestGuildBillingCarriesFeaturesAndActive(t *testing.T) // scoped grant → Features=[officer_views], Scoped()=true, Active=true; expired → Active=false, still returned
```

- [ ] **Step 3: Run them to see them fail** — `cd api && go test -race -run 'TestExpiredGrant|TestPastDue|TestScopedGrant|TestCanForGuild|TestGuildBillingCarries' ./internal/entitlements/` → compile errors / FAIL.

- [ ] **Step 4: Implement**
  - `activeClause` = the Global Constraints predicate; replace the three `activeStatusClause` uses. Keep `isActiveStatus` for the Stripe upsert conflict rule (it compares statuses, not time).
  - `hasVerifiedGuildPlan`: add `and e.features is null`.
  - `canOfficerViews`: add `and (e.features is null or 'officer_views' = any(e.features))`.
  - `CanForGuild`: one query over `guild_members` joined to `entitlements` for `guildID`, returning the three facts (verified officer rank, verified member, covered active row) and mapping to the reasons in the test table. Non-`officer_views` feature → `fmt.Errorf("entitlements: %s is not a per-guild feature", feature)`.
  - `GuildBilling`: select `features`, and `(<activeClause>) as active`.

- [ ] **Step 5: Run the package tests** — `go test -race -cover ./internal/entitlements/` → PASS.

- [ ] **Step 6: Commit** — `feat(api): entitlement features scope, expiring grants, CanForGuild (migration 0049)`

### Task 2: Stripe upsert clears the scope; checkout allows a scoped guild

**Files:**
- Modify: `api/internal/entitlements/store.go` (`UpsertStripe` ~433–470: `features = null` in the `on conflict … do update`), `api/internal/billing/handler.go` (`checkGuildCheckout` 355–360)
- Test: `api/internal/entitlements/store_test.go`, `api/internal/billing/handler_test.go` (or `webhook_test.go`, whichever holds `checkGuildCheckout` cases)

**Interfaces:**
- Consumes: `GuildBilling.Active`, `GuildBilling.Scoped()` (Task 1).

- [ ] **Step 1: Failing tests**
```go
func TestUpsertStripeReplacesScopedGrant(t *testing.T) // seed scoped grant; UpsertStripe active sub → row source='stripe', features null, audit row old 'active' new 'active'
func TestGuildCheckoutAllowedOnScopedGrant(t *testing.T) // checkGuildCheckout on a guild with an active scoped grant → nil refusal; on a full grant → alreadyOnThePlanRefusal; on an expired full grant → nil refusal
```
- [ ] **Step 2: Run, see FAIL.**
- [ ] **Step 3: Implement** — upsert's update set adds `features = null`; checkout's `active :=` becomes `existing != nil && existing.Active && !existing.Scoped()`.
- [ ] **Step 4: `go test -race ./internal/entitlements/ ./internal/billing/` → PASS.**
- [ ] **Step 5: Commit** — `feat(api): a paid plan replaces a scoped grant; checkout open to scoped guilds`

### Task 3: `/v1/me` plan scope and home `officer_tools`

**Files:**
- Modify: `api/internal/auth/handler.go` (`GuildBillingView` 456–462, `attachGuildPlans` 626–662, delete `isActiveEntitlementStatus` 595–600 if unused after), `api/internal/guilds/home.go` (`ViewerView` 220–224, construction 585), `api/internal/guilds/handler.go` (the Service gains an `Entitlements` dependency exposing `CanForGuild`; wire it in `cmd/api/main.go` where guilds.Service is built)
- Test: `api/internal/auth/handler_test.go` (existing me/guild-plan tests), `api/internal/guilds/home_test.go` (or the harness test that covers `home`)

**Interfaces:**
- Produces: JSON `me.guilds[].plan.scope` = `"plan"` | `"officer_tools"`; `GET /v1/guilds/{id}/home` → `viewer.officer_tools: bool`.
- Produces (Go): `guilds.Service.Entitlements` interface `{ CanForGuild(ctx, guildID, userID int64, f entitlements.Feature) (bool, entitlements.Reason, error) }`, nil-safe (nil → `officer_tools` false).

- [ ] **Step 1: Failing tests** — me: officer of a guild with a scoped grant sees `plan.scope == "officer_tools"`, `billed_by == ""`, `current_period_end` the date; a Stripe plan sees `scope == "plan"`; an expired scoped grant attaches no plan. home: officer with grant → `viewer.officer_tools true`; after expiry → false; member → false; signed-out → false, 200.
- [ ] **Step 2: Run, see FAIL.**
- [ ] **Step 3: Implement** — `Scope string \`json:"scope"\`` from `gb.Scoped()`; `attachGuildPlans` skips when `!gb.Active`; `ViewerView.OfficerTools bool \`json:"officer_tools"\`` set from `CanForGuild` only when `role` is officer/leader and the viewer is verified, else false.
- [ ] **Step 4: `go test -race ./internal/auth/ ./internal/guilds/` → PASS.**
- [ ] **Step 5: Commit** — `feat(api): plan scope on /v1/me and officer_tools on the guild home viewer`

### Task 4: `grant --features` and `grandfather-officer-tools`

**Files:**
- Modify: `api/cmd/api/grant.go` (flags 28–75, `resolveGrantSubject`, the run function), `api/cmd/api/main.go` (dispatch ~111), `api/internal/entitlements/store.go` (`Grant` gains `features []string`; add `ClaimedGuildsWithoutActivePlan(ctx) ([]GuildRef, error)` returning id and name)
- Create: `api/cmd/api/grandfather.go`
- Test: `api/cmd/api/grant_test.go`, `api/cmd/api/grandfather_test.go`, `api/internal/entitlements/store_test.go`
- Docs: `api/README.md` (the grant section: the two new forms)

**Interfaces:**
- Produces: `api grant --guild <id> --features officer_views --until YYYY-MM-DD --note "…"` (`--plan` optional, must be `guild`; `--until` required with `--features`); `api grandfather-officer-tools --until YYYY-MM-DD --note "…" [--apply]`.
- Produces (Go): `Grant(ctx, subj, plan, until, grantedBy, note, actor string, features []string)`.

- [ ] **Step 1: Failing tests**
  - `grant_test.go`: `--features officer_views` without `--until` → error `--until is required with --features`; `--features roster_check` → error; `--features` with `--user` → error; `--features officer_views --guild 3 --until 2026-12-09 --note x` parses with plan defaulted to `guild`.
  - `store_test.go`: `Grant` with features writes `features = '{officer_views}'`; re-grant without features clears it (full plan).
  - `grandfather_test.go` (DB harness): three guilds (claimed/no row, claimed/scoped grant, unclaimed) + one on a Stripe plan: dry run writes nothing and lists `granted`/`skipped` lines; `--apply` grants exactly the first; second `--apply` grants none.
- [ ] **Step 2: Run, see FAIL.**
- [ ] **Step 3: Implement** — flag parsing per the test; `Grant` passes `features` through the insert and the `on conflict` update (`features = excluded.features`); the grandfather command prints `granted <id> <name> until <date>` / `skipped <id> <name>: <reason>` and a totals line; actor `cli_grandfather:<operator> <note>`; `--apply` gates the write.
- [ ] **Step 4: `go test -race ./cmd/api/ ./internal/entitlements/` → PASS; `go vet ./...`.**
- [ ] **Step 5: README** — add both command forms under the grant section with one sentence each and the order (deploy 0049 first, then grandfather from the deployed sha).
- [ ] **Step 6: Commit** — `feat(api): grant --features officer_views and the grandfather-officer-tools command`

**PR A gate:** CI green → fast-forward main → ordered deploy (api run with the migration first; verify `schema_migrations` = 49) → run `grandfather-officer-tools` dry run, then `--apply`, from the deployed sha via the production script pattern (never printing `DATABASE_URL`) → verify every claimed guild's row (`features = '{officer_views}'`, `current_period_end = 2026-12-09`) → memory note. Only then Task 5.

---

## PR B — the gate

### Task 5: The server refuses Award and Undo without the feature

**Files:**
- Modify: `api/internal/guilds/loot.go` (`createLootAward` 593–632, `deleteLootAward` 635–660)
- Test: `api/internal/guilds/loot_test.go` (officer-only award test at ~206 is the pattern)

- [ ] **Step 1: Failing tests** — officer + scoped grant → 200 on award and 200/204 on undo; officer + no row → 403 body `{"code":"guild_plan_required","message":"Award and Undo are part of the guild plan."}`; member → 403 `forbidden` (unchanged, asserted first in the handler); officer + full Stripe plan → 200.
- [ ] **Step 2: Run, see FAIL.**
- [ ] **Step 3: Implement** — after `verifiedOfficerOrLeader`, `ok, _, err := s.Entitlements.CanForGuild(ctx, guildID, userID, entitlements.FeatureOfficerViews)`; `!ok` → `httpx.WriteError(w, r, http.StatusForbidden, "guild_plan_required", "Award and Undo are part of the guild plan.", nil)`.
- [ ] **Step 4: `go test -race ./internal/guilds/` → PASS.**
- [ ] **Step 5: Commit** — `feat(api): Award and Undo need the guild plan (guild_plan_required)`

### Task 6: Web contracts and copy

**Files:**
- Modify: `web/src/lib/guild/api.ts` (`GuildApiError` 13–21: `readonly code: string`; the two `throw` sites 377–381 pass `result.code ?? ''`), `web/src/lib/account/api.ts` (`GuildBillingView` gains `scope: 'plan' | 'officer_tools'`; `GuildHome.viewer` type gains `officer_tools?: boolean`), `web/src/lib/guild/copy.ts`, `web/src/lib/guild/settings-view.ts`
- Test: `web/src/lib/guild/settings-view.test.ts`, `web/src/lib/guild/copy.test.ts` (create if absent), `web/src/lib/guild/api.test.ts`

**Interfaces:**
- Produces (copy.ts): `guildGateCopy = { sentence: 'Award and Nudge are part of the guild plan. ', link: 'Subscribe the guild', awardRefused: "Not awarded. This guild's officer tools have ended.", undoRefused: "Not undone. This guild's officer tools have ended.", readOnly: 'Read-only' }`.
- Produces (settings-view.ts): `BillingState` gains `'officer-tools'`; `billingState` returns it when `billing.scope === 'officer_tools'`; `guildPlanLine` returns `Officer tools · until {d Mon yyyy}` for that scope; `officerToolsExplanation(billing): string` = `Granted by Forever Sixty for Award, Undo and Nudge only. It ends on {d Mon yyyy}; the guild plan keeps them after that.`; `accountGuildPlanLine(billing): string` = `guildPlanLine(billing)` (the account row appends ` · ` and the link itself).
- Produces: `GuildApiError.code`.

- [ ] **Step 1: Failing tests** — every string above read verbatim from `design/specs/2026-10-25-officer-tools-gate.md` §7's table (the Discord copy test's `specCell` reader is the pattern: copy it into a small `web/src/test-support/spec-table.ts` and use it from both tests); `billingState({scope:'officer_tools', billed_by:''})` → `'officer-tools'`; `guildPlanLine` scoped with `2026-12-09T00:00:00Z` → `Officer tools · until 9 Dec 2026`; `GuildApiError` from an envelope `{code:'guild_plan_required'}` carries `code`.
- [ ] **Step 2: Run `npx vitest run src/lib/guild` → FAIL.**
- [ ] **Step 3: Implement.**
- [ ] **Step 4: `npx vitest run src/lib/guild src/lib/account` → PASS; `npm run check`.**
- [ ] **Step 5: Commit** — `feat(web): guild gate copy, scoped plan lines, GuildApiError.code`

### Task 7: `officer-read-only` on Loot and Readiness, and the 403 path

**Files:**
- Modify: `web/src/components/guild/GuildLoot.svelte` (props 31–45, header 67, Undo 149–160, Award 163–175, confirm 178–205), `web/src/components/guild/GuildReadiness.svelte` (props 30–34, `NUDGE_BTN` 67, buttons 184–190 and 250–256, h2), `web/src/components/Guild.svelte` (flag near 245; `onAward`/`onUndo` 571–605; the two mounts 929 and 951)
- Test: `web/src/components/guild/GuildLoot.test.ts`, create `GuildReadiness.test.ts`

**Interfaces:**
- Consumes: `guildGateCopy`, `GuildApiError.code`, `home.viewer.officer_tools`, `guildSettingsHref`.
- Produces: `GuildLoot` props `officerTools: boolean`, `refusedItemId: number | null`, `refusedKind: 'award' | 'undo' | null`; `GuildReadiness` prop `officerTools: boolean`; `Guild.svelte` state `officerToolsEnded`.

- [ ] **Step 1: Failing component tests** (server-render html, as the existing tests do)
  - Loot, `officer && !officerTools`: the sentence `p#officer-tools-gate-loot[data-testid=guild-loot-plan-gate]` precedes the boss picker and the empty state; every Award and Undo has `disabled` and `aria-describedby="officer-tools-gate-loot"`; the `Read-only` span is present; no confirm row renders even with `confirming` set.
  - Loot, member: no sentence, no label change, no controls (unchanged assertions).
  - Loot, `refusedItemId` set: one `p[role=alert][tabindex=-1][data-testid=guild-loot-plan-refused-{id}]` with `awardRefused` (or `undoRefused`) inside that item's card.
  - Readiness, `officer && !officerTools`: sentence `#officer-tools-gate-readiness` under the h2, `Read-only` span, every Nudge `disabled` with the describedby; footnote still names Nudge.
- [ ] **Step 2: Run, see FAIL.**
- [ ] **Step 3: Implement** per the experience spec §3.1–3.4 and §4: the flag `officerTools = (home?.viewer?.officer_tools ?? true) && !officerToolsEnded`; in `onAward`/`onUndo`, `thrown instanceof GuildApiError && thrown.code === 'guild_plan_required'` sets `officerToolsEnded = true`, `refusedItemId`, `refusedKind`, leaves `lootError` empty, and after render focuses the refused line (an `$effect` keyed on `refusedItemId`); disabled classes exactly as §3.4's table; the link gets `py-3` below `md`.
- [ ] **Step 4: `npx vitest run src/components/guild` → PASS; `npm run check`; eslint + prettier on the touched files.**
- [ ] **Step 5: Commit** — `feat(web): officer-read-only Loot and Readiness; the guild_plan_required path`

### Task 8: Guild settings and the account guild row

**Files:**
- Modify: `web/src/components/GuildSettings.svelte` (billing block 221–270), `web/src/components/Account.svelte` (guild row 815–870)
- Test: `web/src/components/GuildSettings.test.ts` (create if absent, server-render), `web/src/components/account/AccountGuildRow.test.ts` if the row is extracted (extract `AccountGuildRow.svelte` if `Account.svelte` would pass 1000 lines; it is at ~900)

**Interfaces:**
- Consumes: `billingState` `'officer-tools'`, `guildPlanLine`, `officerToolsExplanation`, `notOnPlanLines`, `guildSettingsHref`.

- [ ] **Step 1: Failing tests**
  - Settings, claim holder, scoped: `guild-billing-plan` = `Officer tools · until 9 Dec 2026`; `guild-billing-contact` = the explanation; price line; `Subscribe the guild` button with the checkout href; no Manage billing.
  - Settings, `mayBill === false` with a plan: the plan line and explanation render, then `guildMasterOnly`, no button. With no plan: today's single line.
  - Account guild row: `plan !== null` → `p[data-testid=account-guild-plan]` reading `Officer tools · until 9 Dec 2026 · Guild settings` with the link to `guildSettingsHref`; a paid plan reads `Guild plan · renews 31 Dec 2026 · Guild settings`; `plan === null` → no row.
- [ ] **Step 2: Run, see FAIL.**
- [ ] **Step 3: Implement** per the experience spec §5 and §6 (classes and slots as written there).
- [ ] **Step 4: `npx vitest run src/components` → PASS; `npm run check`.**
- [ ] **Step 5: Commit** — `feat(web): scoped grant in guild settings and on the account guild row`

### Task 9: e2e, captures, reviews

**Files:**
- Create: `web/tests/e2e/guild-officer-tools.spec.ts`, `web/tests/captures/officer-tools-gate/{capture.spec.ts,playwright.capture.config.ts}`, `design/reviews/2026-10-25-officer-tools-gate/README.md`
- Modify: `web/src/fixtures/guild/mock-guild.ts` (`buildMockHome` accepts `officerTools?: boolean`), `web/src/test-support/billing-fixtures.ts` (`meWithGuildPlan` accepts `scope`)

- [ ] **Step 1: Write the e2e** — the checks in the experience spec §8 (axe on #2, #7, #10, #13; disabled counts equal enabled counts; a click on a disabled control opens no confirm and writes nothing to the clipboard; the link resolves to `/guild/{region}/{ruleset}/{slug}/settings`; the sentence's box inside the viewport at 360 and 1440 on load; after a routed 403 `guild_plan_required` on Award, a `window` marker survives, the refused line is focused, Readiness is read-only). Desktop and mobile projects.
- [ ] **Step 2: Run against a fixture build** — `cd web && FOREVER_DATA=fixture npm run build && E2E_SKIP_BUILD=1 E2E_PORT=4418 npx playwright test tests/e2e/guild-officer-tools.spec.ts tests/e2e/guild-centre.spec.ts tests/e2e/guild-settings-billing.spec.ts tests/e2e/account-billing.spec.ts` → PASS.
- [ ] **Step 3: Captures** — the 13 states × widths in §8, PNGs under `design/reviews/2026-10-25-officer-tools-gate/`, with the panel heights annotated against §3.5.
- [ ] **Step 4: Reviews** — ux-designer and wow-player on the PNGs (SHIP / SHIP WITH FIXES); fix round if needed; reports in the reviews folder.
- [ ] **Step 5: Commit** — `test(web): officer tools gate e2e, captures and reviews`

### Task 10: Changelog and page truth

**Files:**
- Create: `web/src/content/changelog/2026-10-xx-officer-tools-gate.md` (dated the ship day)
- Modify: `web/src/lib/guild/settings-view.ts` `notOnPlanLines().detail` → `…Officers get the loot council helper and the raid readiness board, with more as they ship.` (the designer's out-of-lane flag; §4.7 says the page names only what ships; ruling: fix it here, it is one string) with its test; confirm `web/src/lib/billing/copy.ts:21-24` already reads "with more as they ship".

- [ ] **Step 1: Update the test for `notOnPlanLines`, see FAIL, fix, PASS.**
- [ ] **Step 2: Write the changelog entry** (player-facing, three sentences: officer tools are part of the guild plan; every guild claimed before today keeps them through 9 December; what to do).
- [ ] **Step 3: Layout spec** — `tests/e2e/layout.spec.ts` if the changelog teaser matches a role/name query (the Discord lane hit this).
- [ ] **Step 4: Commit** — `docs(web): officer tools gate changelog; the guild settings price line names what ships`

**PR B gate:** full web vitest + full Playwright suite (a changelog entry touches the shared layout) + lhci shard → CI green → fast-forward main → ordered deploy (api first, then web) → verify in production with the owner's guild (`officer_tools` on the home response; the settings line) → memory → remove the worktree.
