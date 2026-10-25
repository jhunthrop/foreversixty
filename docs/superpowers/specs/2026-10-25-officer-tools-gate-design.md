# Officer tools gate — design (premium lane 4.7)

Status: design, 2026-10-25. Implements §4.7 of
`docs/superpowers/specs/2026-10-24-premium-experience-design.md` (the owner's rulings there
are binding: build the officer-tools-only grant kind, grandfather every guild claimed before
the gate, ship the gate last). This document settles what §4.7 leaves open, from a read of
the code as it is today. The experience half (what an officer sees in the read-only state,
how the grant shows in guild settings and on the account page) is the ux-designer's spec at
`design/specs/2026-10-25-officer-tools-gate.md`.

## 1. Intent

Officers of a claimed guild have had the loot council helper and the readiness board free
since 2026-10-04, gated by rank alone. The guild plan promises them as its headline. From
this lane on, their controls work only for a guild on the plan, with one exception the owner
chose: every guild claimed before the gate ships keeps them through the first raid night
(2026-12-09) by an **officer-tools grant** that gives the guild nothing else.

## 2. Goals, non-goals, anti-goals

Goals

- A grant kind that turns on `officer_views` for one guild and nothing else: no Premium for
  its members, no Supporter chip, no server sims, no history.
- Grants expire. An `--until` date that has passed turns the feature off without a job.
- Award and Undo refuse on the server for a guild without the feature. Nudge, which never
  reaches the server, is disabled on the web by the same fact.
- The gate ships after the grant kind and after the grandfather run, in that order.

Non-goals

- Roster approvals and invite links stay rank-gated (free). §4.7 names only the loot helper
  and the readiness board.
- Loot proposals: §4.7 lists them, but no proposal exists in the API or the web. Nothing to
  gate; nothing built here. (Ruling: drop the word from the gate; the loot lane that adds
  proposals gates them on the same check.)
- No self-serve way to get a grant. It is the `api grant` CLI, run by the owner.

Anti-goals

- A scoped grant that reads as a paid plan anywhere `Can`, `IsSupporter` or billing looks.
- Gating a non-officer out of something they could see yesterday. Members see nothing new.
- A guild on a grant that cannot buy the plan ("already on the plan" from the grant row).

## 3. Constraints

- `entitlements` is one row per (guild, plan) with `plan in ('premium','guild')` and
  `unique (guild_id, plan)` (migration 0020). A guild therefore holds one `guild` row,
  whether it is a Stripe subscription or a grant.
- `Can(ctx, userID, feature)` is per user; `canOfficerViews` is true for an officer of *any*
  guild on the plan. Handlers that act on one guild need the guild in the question.
- Only the deployed API runs migrations; the `grant` CLI migrates at startup too, so the
  grandfather run happens from the deployed sha, after the API that carries 0049 is live.
- The web reads the plan from `/v1/me` (`guilds[].plan`, `guilds[].can`) and the guild page
  from `/v1/guilds/{id}/home` (`viewer.role`). The gate must be answerable from what the
  page already fetches; no extra request per tab.

## 4. Decisions

### 4.1 Schema: `features` on the entitlement row (migration 0049)

- `alter table entitlements add column features text[]` (nullable). `null` means the whole
  plan, as every row today. `'{officer_views}'` means that feature alone.
- Check: `features is null or (plan = 'guild' and source = 'grant' and
  features <@ array['officer_views'])`. Only a guild grant can be scoped, and only to this
  feature for now; a later feature widens the array literal in a later migration.
- No change to uniqueness. A guild with a scoped grant that buys the plan gets its row
  **replaced** by the Stripe upsert (`on conflict (guild_id, plan) do update … source =
  'stripe'`, and this lane adds `features = null` to that update). The audit row records the
  change. If the subscription later cancels, the grant is gone with it; the owner re-grants
  by hand if that ever matters (it should not: grants end 2026-12-09).
- `entitlement_audit` gains nothing; the actor string already carries the operator note.

### 4.2 Active means not expired, for grants

One predicate, used everywhere `activeStatusClause` is used today (`Can`, `IsSupporter`,
`hasVerifiedGuildPlan`, `canOfficerViews`, billing's `checkGuildCheckout`, the `/v1/me`
views):

```
status in ('active','trialing','past_due')
and (source <> 'grant' or current_period_end is null or current_period_end > now())
```

A Stripe row's life is Stripe's: `past_due` is inside the retry window and stays on. A
grant has no Stripe behind it, so its own `--until` is the end. Today an expired grant
stays on forever; this closes that for every grant, personal ones included (spec
2026-09-21 §1.3 rule 4 already says access follows status and period end). `GuildBilling`
and `Billing` keep returning the row with no filter, so settings pages can say "ended".

### 4.3 Scope-aware reads

- `hasVerifiedGuildPlan` (feeds Premium features and the Supporter mark) adds
  `and e.features is null`. A scoped grant gives members nothing.
- `canOfficerViews` adds `and (e.features is null or 'officer_views' = any(e.features))`.
- New `CanForGuild(ctx, guildID, userID, feature) (bool, Reason, error)`: for
  `FeatureOfficerViews`, true when the user is a verified officer or leader of `guildID`
  and that guild's row is active (4.2) and covers the feature (`features is null or
  'officer_views' = any(features)`). `ReasonNotOfficer` when verified in that guild below
  officer rank; `ReasonNoPlan` otherwise. Other features are not per guild and return an
  error. This is the handler gate.
- `GuildBilling` gains `Features []string` and a derived `Scoped() bool`.

### 4.4 The gate (server)

- `POST` and `DELETE /v1/guilds/{id}/loot/awards`: after the existing officer check, call
  `CanForGuild(id, user, FeatureOfficerViews)`; refuse with 403 and the envelope error code
  `guild_plan_required` (message `Award and Undo are part of the guild plan.`). The existing
  403 for non-officers is unchanged and checked first, so a member never learns the plan
  state from the code.
- `GET …/loot` and `GET …/readiness` stay as they are for officers: the board and the helper
  stay visible (§4.7, "both tools stay visible to every officer").
- Nudge copies `nudge_text` client-side; the web disables it on the same flag (4.5). The
  text stays in the payload: the board is readable, the one-click copy is the plan's.

### 4.5 What the web learns

- `/v1/guilds/{id}/home` `viewer` gains `officer_tools: bool` — `CanForGuild` for the
  viewer, false for non-officers. Guild.svelte already passes `officer` into the Loot and
  Readiness tabs; it passes `officerTools` beside it. No extra request.
- `/v1/me` `guilds[].plan` (`GuildBillingView`) gains `scope: "plan" | "officer_tools"`.
  Guild settings renders the scoped grant as the designer specs it (`Officer tools · until
  9 Dec 2026` is the working line; the designer's spec is binding).
- `guilds[].can` is unchanged (`award` stays rank-derived; it answers "may this person
  award in principle", the plan answers "today").
- The Premium page's officer-tools line already reads "with more as they ship" (lane 1).
  The account page: the designer decides whether a guild line appears there; the API
  already exposes what it would need through `guilds[].plan`.

### 4.6 The CLI and the grandfather run

- `api grant --guild <id> --features officer_views --until 2026-12-09 --note "…"`.
  `--features` accepts only `officer_views`, requires `--guild`, and implies `--plan guild`
  (`--plan` may be given and must be `guild`). `--until` is required with `--features` (a
  scoped grant is a bridge, never open-ended). `api revoke --guild <id> --plan guild` is
  unchanged and ends either kind.
- `api grandfather-officer-tools --until 2026-12-09 --note "claimed before the gate"`:
  every guild with `claimed_at is not null` and no active `guild` row (4.2) gets the scoped
  grant; a guild already on the plan or on a grant is skipped and listed. Prints one line
  per guild (`granted 12 Olympus XXVII until 2026-12-09`, `skipped 7 … on the plan`), then
  the totals. Idempotent: a second run grants nothing. Dry run by default; `--apply` writes.
- Audit actor strings: `cli_grant:<operator> <note>` as today; the grandfather run uses
  `cli_grandfather:<operator> <note>`.

### 4.7 Order within the lane

1. Migration 0049, `features`, the active predicate (4.2), the scoped reads (4.3),
   `CanForGuild`, `GuildBilling.Features`, the `/v1/me` `scope` and home `officer_tools`
   fields, the CLI flag and the grandfather command, with tests. Deploy. Verify
   `schema_migrations` reads 49.
2. The owner runs `api grandfather-officer-tools --until 2026-12-09 --note "…"` from the
   deployed sha (dry run first, then `--apply`), from a shell with `DATABASE_URL` from Secret
   Manager and never printed. Verify: every claimed guild has a row with
   `features = '{officer_views}'`, `current_period_end = 2026-12-09`.
3. The gate: the handler refusals (4.4), the web read-only state (designer's spec), the
   e2e of §4.7's acceptance. Deploy.

Steps 1 and 3 are two PRs in one lane, so the gate cannot land ahead of the grants.

## 5. Acceptance (from §4.7, made concrete)

- API: a verified officer of a guild with a scoped grant awards (201) and undoes (204); the
  same officer after `current_period_end` passes gets 403 `guild_plan_required`; a verified
  member of that guild has no Premium (`Can(server_sims)` false) and no Supporter mark; an
  officer of a guild on the Stripe plan is unchanged.
- CLI: `grant --features officer_views` without `--until` is refused; the grandfather dry
  run lists and writes nothing; `--apply` writes once; a second `--apply` writes nothing.
- Web e2e (`guild-officer-tools.spec.ts`): a claimed guild, an officer, `officer_tools`
  toggled in the home fixture: Award, Undo and Nudge enabled with it, disabled and the
  designer's one sentence shown without it, the sentence linking to guild settings; a member
  sees neither the controls nor the sentence. Guild settings shows the scoped grant line and
  its date.

## 6. Left to the builder

Package placement of `CanForGuild` (it belongs in `entitlements`, next to `Can`); whether
`GuildBillingView.scope` is derived in `auth` or returned by `entitlements`; the grandfather
command's file (`api/cmd/api/grandfather.go` beside `grant.go`); test fixtures.
