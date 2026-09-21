# Guild web — reconcile with the real API Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Make `web/src/lib/guild/` and every guild component match the real,
now-landed `api/internal/guilds/*` handlers field-for-field, and build the new claim
states, contest flow, frozen-claim state, roster removal rules, and report-visibility
copy the API's 2026-09-21 security hardening round added.

**Architecture:** This is a reconciliation, not new scaffolding: every file this plan
touches already exists from the first guild-web plan. Read the real API code
(`/Users/jh/code/forever/.worktrees/guild-api/api/internal/guilds/*.go`, read-only) as
ground truth over the spec's own endpoint table where the two disagree — three do
disagree (approve/remove's path shape, approve's response body, and settings'
`claim_pending` type), confirmed by reading the actual Go structs and handlers, not
guessed.

**Tech Stack:** Same as the original plan — Astro, Svelte 5, TypeScript, Vitest,
Playwright.

**Spec:** `docs/superpowers/specs/2026-09-21-guild-membership-design.md` in THIS
worktree (unchanged since the first plan) plus the same file in
`/Users/jh/code/forever/.worktrees/guild-api` (read-only, its sections 2.4/2.6/3.3
carry the 2026-09-21 security-review amendment this plan builds against — the API
worktree's copy is newer; this worktree's copy was not touched by the reconcile and is
not re-synced here, since nothing in it needs correcting — the amendment only adds
detail this plan already read from the API worktree's copy).

## Global Constraints (carried from the first plan, still binding)

- Work only in `/Users/jh/code/forever/.worktrees/guild-web`. Never merge this branch
  into `main`, never push. `guild-api` at `/Users/jh/code/forever/.worktrees/guild-api`
  is READ-ONLY reference — never write to it.
- `git merge main` has already been run in this worktree (clean, no conflicts, merge
  commit `c284e51`) — do not re-run it.
- Web lane file ownership unchanged: must not touch `api/`, `addon/`, `logs/`,
  `web/src/components/sim/`, `web/src/components/planner/`, `web/src/lib/sim/`,
  `web/src/lib/planner/`, `web/src/components/report/`, `Header.astro`, `Footer.astro`.
- Every visible string in a copy module (`web/src/lib/guild/copy.ts`). Design system
  restraint: plain sentences, no exclamation marks, no alarm colors beyond what
  `SECONDARY_BUTTON`/existing pill/alert classes already provide (no new red/warning
  color token — reuse `role="alert"` + the existing muted/strong text classes, exactly
  as every other error state in this codebase already does).
- Functions under 50 lines, files under 800, no magic numbers, no dead code, immutable
  updates.
- Toolchain from `web/`: `export NVM_DIR="$HOME/.nvm"; . "$NVM_DIR/nvm.sh"; nvm use
  22.12`, `FOREVER_DATA=fixture npm run sync` once, then before every commit:
  `npx vitest run <paths>`, `npx astro check`, `npm run lint`,
  `npx prettier --check <paths>`. E2E: `E2E_PORT=4421 npx playwright test <files>`.
- Commits: printf a message to a file under `.superpowers/`, then `git commit -F
  <file>` as its own command. Never bare `git stash`, never `-n`/`--no-verify`. End
  every message with exactly:
  `Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>` then
  `Claude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5`
- Run the WHOLE e2e suite once before the final report (not just touched specs) — the
  coordinator's own instruction: main's full suite had 1 failure this morning, already
  fixed, so treat ANY failure after this merge as this lane's own until proven
  otherwise against a clean `main` checkout (`git -C /Users/jh/code/forever/web
  status`/`git log` can confirm a file is untouched by this branch; a failure in a file
  this branch never touched, in a spec this branch never touched, is the strongest
  proof).

## Field-level reconciliation record (from reading the real API code directly)

Read: `api/internal/guilds/{handler,claim,contest,home,roster,settings,invite,store}.go`
in the `guild-api` worktree. Every mismatch below is a required web-side fix; three are
flagged as **required API changes** the coordinator should relay, since nothing on the
web side can substitute for them.

1. **`GET .../home`'s `HomeView`** is `{guild, claim, reports: HomeReport[],
   next_cursor?, roster: RosterRow[]}` — `reports` is a flat array with `next_cursor` as
   a *sibling* field, not nested inside a `{rows, next_cursor}` object as this lane's
   own Ruling 1 first guessed. Fixed below.
2. **`HomeReport`** is `{id, title, created_at, fight_count, kill_count}` — no `zone`,
   no `wipe_count`. `wipe_count` is derived client-side as `fight_count - kill_count`
   (this is unambiguous: the spec's own "kill/wipe" framing and the sibling
   `fight_count`/`kill_count` fields exist for exactly this subtraction, same as
   `ReportRow.svelte`'s pattern elsewhere in the codebase already implies). **Required
   API change (recommend, do not block on)**: `HomeReport` has no `zone` to fall back
   to when `title` is empty, unlike every other report list in this codebase
   (`RecentReport`, `MyReport`); until it's added, an untitled report on the guild home
   shows `guildHomeCopy.untitledReport` instead of a zone name.
3. **`RosterRow`** has no `user_id` field at all. The `user_id` this lane invented in
   its own final-review fix round does not exist on the real API. **Do not keep it as
   a lying optional field.** Fixed below by deriving "is this my own row" from
   `Me.characters[].key` (already fetched, already exact-matches `character_key`)
   instead — this actually fully replaces the need for `user_id`: the empty-roster
   heuristic only needs "does any row belong to someone other than me," which
   `me.characters` answers exactly as well as a `user_id` would, with no API change
   needed. **No required API change for this one** — the client-side substitute is
   complete.
4. **Approve/remove take THREE path segments, never a combined `character_key`.**
   Routes are `POST /v1/guilds/{id}/characters/{region}/{ruleset}/{name}/approve` and
   `DELETE /v1/guilds/{id}/characters/{region}/{ruleset}/{name}` (confirmed in
   `handler.go`'s `Mount` and `roster.go`'s `characterKeyFrom`, which reads
   `r.PathValue("region")`/`"ruleset"`/`"name")` and re-slugifies `name` server-side via
   `character.Key`). This lane's Ruling 3 (percent-encoding a combined `character_key`)
   was a defensible guess against an unlanded API, but the real router never receives a
   combined key at all — there is nothing to percent-encode. **Ruling 3 is retired.**
   `approveCharacter`/`removeCharacter` take `(guildId, region, ruleset, slug)`.
5. **Approve's 200 body is `{character_key, status: "approved"}`, not an updated
   `RosterRow`.** The spec's own table says "updated character row"; the actual handler
   returns only the two fields above (confirmed in `roster.go`'s `approveCharacter`).
   Code wins. The web must locally flip `verified: true` on the matching roster row
   itself after a successful approve (the response gives no other roster fields to
   merge in).
6. **`SettingsView.ClaimPending` is `{by: {battletag}, expires_at} | null`, never a
   boolean.** This lane's `claim_pending: boolean` was wrong from the first plan.
   Fixed below. `SettingsView` also gains `claim: ClaimStateView` (`{state, since?}`),
   confirmed present and populated in `settings.go`'s `getSettings` handler.
7. **`Me.guilds[].verified` is never `omitempty` in the real API** (`auth/store.go`'s
   `Guild` struct: `Verified bool` with no `omitempty` tag) — always present. Loosen
   this lane's own `MeGuild.verified?: boolean` to a required `verified: boolean`.
8. **The `auth.Store.Guilds` `ORDER BY m.refreshed_at DESC` change has already landed**
   (confirmed reading `auth/store.go:395-400`) — Task 12's parked cross-lane-dependency
   finding from the first plan's final review is now resolved; no web change needed,
   just confirmed.
9. **The public `GuildPage.guild` already carries `id`** (confirmed
   `rankings/guilds.go`'s `GuildIdentity{ID, Name, Region, Ruleset}`) — this lane's
   Ruling 2 addition to `web/src/lib/rankings/api.ts`'s `GuildPage.guild` type was
   correct and needs no change; the real API agrees with it exactly.
10. **New:** `HomeView`/`SettingsView` both gain `claim: {state:
    "unclaimed"|"pending"|"claimed"|"contested", since?}`. New endpoint `POST
    /v1/guilds/{id}/claim/contest` → `{status: "contested"}`, errors 403/404/409 (each
    already carries a good, displayable sentence from the API — confirmed reading
    `contest.go`'s `contestClaim` handler; the web does not need to re-map these, only
    display `thrown.message` the same way every other action-error in this lane's
    components already does). New 409 `claim_contested` on approve, remove-by-non-self,
    settings PATCH, and invite rotate while a claim is contested (confirmed in each
    handler). New rank-protects-rank rule on remove (confirmed `roster.go`'s
    `mayRemoveCharacter`): a `member`-rank row is removable by self, verified
    officer/leader, or moderator (unchanged); an `officer`-rank row adds *only the
    account holding the guild's claim* to that list, not any officer; a `leader`-rank
    row is removable only by its own account or a moderator, never by anyone else.
    **The web cannot learn "am I the claim holder" from the home response alone** (no
    `viewer`/claimant-identity field on `HomeView`) — this plan's ruling (below, Task 2)
    is to hide Remove on officer-rank rows for everyone except self/moderator, which is
    conservative (never shows a button that 403s) at the cost of also hiding it from a
    genuine claim-holder officer, who still has the settings-page path or the API
    directly. **Required API change (recommend)**: expose the claim holder's identity
    (or a `may_remove` boolean per row) on `HomeView` so this conservative hide can be
    replaced with an exact one.

## Task 1: `web/src/lib/guild/api.ts` and `copy.ts` — field-exact reconciliation

**Files:**
- Modify: `web/src/lib/guild/api.ts`
- Modify: `web/src/lib/guild/api.test.ts`
- Modify: `web/src/lib/guild/copy.ts`
- Modify: `web/src/lib/account/api.ts` (only `MeGuild.verified` — drop `?`)

**Interfaces produced (Tasks 2-5 depend on these exact names/shapes):**
- `ClaimState = 'unclaimed' | 'pending' | 'claimed' | 'contested'`
- `ClaimStateView { state: ClaimState; since?: string }`
- `MemberRef { battletag: string }`
- `ClaimPendingView { by: MemberRef; expires_at: string }`
- `GuildHomeReport { id: string; title: string; created_at: string; fight_count:
  number; kill_count: number }` (no `zone`, no `wipe_count`)
- `GuildRosterRow` — same fields as today MINUS `user_id` and `updated_at` (the latter
  was never sent by the real API either — `RosterRow` in `home.go` has no
  `updated_at`, only `logged_recently`; drop it, it was unused dead weight from the
  first plan's own guess)
- `GuildHome { guild: GuildSummary; claim: ClaimStateView; reports: GuildHomeReport[];
  next_cursor?: string; roster: GuildRosterRow[] }` (no `viewer`)
- `GuildSettingsData { default_visibility; officer_max_rank_index; claimed_by:
  MemberRef | null; claim_pending: ClaimPendingView | null; claim: ClaimStateView;
  invite: { rotated_at: string | null } }`
- `ContestResult { status: 'contested' }`
- `ApproveResult { character_key: string; status: 'approved' }`
- `contestClaim(guildId, apiBase?): Promise<ContestResult>`
- `approveCharacter(guildId, region, ruleset, slug, apiBase?): Promise<ApproveResult>`
- `removeCharacter(guildId, region, ruleset, slug, apiBase?): Promise<RemovedResult>`
  (unchanged return type, changed parameters)
- `fetchGuildHome(guildId, cursor?, apiBase?): Promise<GuildHome>` (cursor now wired)

- [ ] **Step 1: Write the failing tests**

Replace `web/src/lib/guild/api.test.ts`'s content with (keep the file's existing
`envelope()` helper and `afterEach`/imports pattern, add these cases; remove any
existing case that asserted the old, wrong shapes — e.g. the old
`approveCharacter`/`removeCharacter` percent-encoding test):

```ts
import { afterEach, describe, expect, it, vi } from 'vitest';
import {
  GUILD_API_FAILED,
  GuildApiError,
  acceptInvite,
  approveCharacter,
  claimGuild,
  confirmClaim,
  contestClaim,
  fetchGuildHome,
  fetchGuildSettings,
  leaveGuild,
  releaseClaim,
  removeCharacter,
  rotateInvite,
  updateConsent,
  updateGuildSettings,
} from './api';

const API = 'https://api.foreversixty.test';
type GlobalFetch = (...args: Parameters<typeof fetch>) => Promise<Response>;

function envelope(data: unknown, status = 200): Response {
  return new Response(JSON.stringify({ ok: status < 400, data, error: null, request_id: 'r' }), {
    status,
    headers: { 'content-type': 'application/json' },
  });
}

afterEach(() => vi.unstubAllGlobals());

describe('fetchGuildHome', () => {
  it('reads the flat reports array and top-level next_cursor, plus claim state', async () => {
    const home = {
      guild: { id: 42, region: 'us', ruleset: 'hardcore', name: 'The Last Watch' },
      claim: { state: 'claimed' },
      reports: [{ id: 'r1', title: 'Sanguine Depths', created_at: '2026-09-20T20:00:00Z', fight_count: 8, kill_count: 3 }],
      next_cursor: 'abc123',
      roster: [],
    };
    const upstream = vi.fn<GlobalFetch>(async () => envelope(home));
    vi.stubGlobal('fetch', upstream);

    const result = await fetchGuildHome(42, undefined, API);

    expect(result.claim.state).toBe('claimed');
    expect(result.reports[0].fight_count).toBe(8);
    expect(result.next_cursor).toBe('abc123');
    expect((upstream.mock.calls[0][0] as Request).url).toBe(`${API}/v1/guilds/42/home`);
  });

  it('sends a cursor query param when given one', async () => {
    const upstream = vi.fn<GlobalFetch>(async () =>
      envelope({ guild: { id: 42, region: 'us', ruleset: 'hardcore', name: 'X' }, claim: { state: 'unclaimed' }, reports: [], roster: [] }),
    );
    vi.stubGlobal('fetch', upstream);
    await fetchGuildHome(42, 'abc123', API);
    expect((upstream.mock.calls[0][0] as Request).url).toBe(`${API}/v1/guilds/42/home?cursor=abc123`);
  });
});

describe('claim and contest', () => {
  it('claims, confirms, releases and contests with POST', async () => {
    const upstream = vi.fn<GlobalFetch>(async () => envelope({ status: 'confirmed' }));
    vi.stubGlobal('fetch', upstream);
    await claimGuild(42, API);
    expect((upstream.mock.calls[0][0] as Request).method).toBe('POST');
    await confirmClaim(42, API);
    await releaseClaim(42, API);

    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => envelope({ status: 'contested' })),
    );
    const result = await contestClaim(42, API);
    expect(result.status).toBe('contested');
  });
});

describe('settings', () => {
  it('reads claim_pending as an object, and the new claim field', async () => {
    const upstream = vi.fn<GlobalFetch>(async () =>
      envelope({
        default_visibility: 'guild',
        officer_max_rank_index: 1,
        claimed_by: null,
        claim_pending: { by: { battletag: 'Fixture#1234' }, expires_at: '2026-10-01T00:00:00Z' },
        claim: { state: 'pending', since: '2026-09-17T00:00:00Z' },
        invite: { rotated_at: null },
      }),
    );
    vi.stubGlobal('fetch', upstream);
    const settings = await fetchGuildSettings(42, API);
    expect(settings.claim_pending?.by.battletag).toBe('Fixture#1234');
    expect(settings.claim.state).toBe('pending');
  });
});

describe('roster management', () => {
  it('approves and removes by three literal path segments, never an encoded combined key', async () => {
    const upstream = vi.fn<GlobalFetch>(async () =>
      envelope({ character_key: 'us/hardcore/thrallgar', status: 'approved' }),
    );
    vi.stubGlobal('fetch', upstream);

    const result = await approveCharacter(42, 'us', 'hardcore', 'thrallgar', API);
    expect(result.status).toBe('approved');
    expect((upstream.mock.calls[0][0] as Request).url).toBe(
      `${API}/v1/guilds/42/characters/us/hardcore/thrallgar/approve`,
    );
    expect((upstream.mock.calls[0][0] as Request).method).toBe('POST');

    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => envelope({ status: 'removed' })),
    );
    await removeCharacter(42, 'us', 'hardcore', 'thrallgar', API);
    expect((upstream.mock.calls[0][0] as Request).url).toBe(
      `${API}/v1/guilds/42/characters/us/hardcore/thrallgar`,
    );
    expect((upstream.mock.calls[0][0] as Request).method).toBe('DELETE');
  });
});

describe('membership and invite', () => {
  it('PATCHes consent and DELETEs to leave', async () => {
    const upstream = vi.fn<GlobalFetch>(async () => envelope({ consent: 'gear_bags' }));
    vi.stubGlobal('fetch', upstream);
    await updateConsent(42, 'gear_bags', API);
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => envelope({ status: 'left' })),
    );
    const left = await leaveGuild(42, API);
    expect(left.status).toBe('left');
  });

  it('accepts an invite token', async () => {
    const upstream = vi.fn<GlobalFetch>(async () =>
      envelope({ guild: { id: 42, region: 'us', ruleset: 'hardcore', name: 'The Last Watch' }, rank: 'member' }),
    );
    vi.stubGlobal('fetch', upstream);
    const result = await acceptInvite('abc-123', API);
    expect(result.guild.name).toBe('The Last Watch');
  });

  it('rotates the invite', async () => {
    const upstream = vi.fn<GlobalFetch>(async () =>
      envelope({ token: 'abc', url: '/guild/invite/abc', rotated_at: '2026-09-21T00:00:00Z' }),
    );
    vi.stubGlobal('fetch', upstream);
    const rotated = await rotateInvite(42, API);
    expect(rotated.token).toBe('abc');
  });

  it('surfaces GUILD_API_FAILED when the network fails outright', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => {
        throw new TypeError('offline');
      }),
    );
    await expect(fetchGuildHome(42, undefined, API)).rejects.toThrow(GUILD_API_FAILED);
    await expect(fetchGuildHome(42, undefined, API)).rejects.toBeInstanceOf(GuildApiError);
  });
});
```

- [ ] **Step 2: Run to verify failure**

Run: `cd web && npx vitest run src/lib/guild/api.test.ts`
Expected: FAIL (old shapes/signatures).

- [ ] **Step 3: Rewrite `web/src/lib/guild/api.ts`**

Keep `GUILD_API_FAILED`, `GuildApiError`, `GuildRank`, `GuildConsent`,
`GuildVisibility`, `GuildSummary`, `call<T>()`, `updateConsent`, `leaveGuild`,
`claimGuild`, `confirmClaim`, `releaseClaim`, `rotateInvite`, `acceptInvite`,
`ClaimResult`, `ClaimConfirmResult`, `ClaimReleaseResult`, `InviteRotateResult`,
`InviteAcceptResult`, `RemovedResult`, `LeftResult`, `UpdatedMember` exactly as they
are today (all confirmed correct against the real API in the reconciliation record
above). Replace everything else:

```ts
export type ClaimState = 'unclaimed' | 'pending' | 'claimed' | 'contested';

/** GET .../home and GET .../settings both expose this (2026-09-21 security amendment). */
export interface ClaimStateView {
  state: ClaimState;
  since?: string;
}

export interface MemberRef {
  battletag: string;
}

export interface ClaimPendingView {
  by: MemberRef;
  expires_at: string;
}

/** One report row. No `zone`: HomeReport carries none (plan's reconciliation record,
 *  item 2) — an empty title falls back to guildHomeCopy.untitledReport, not a zone. */
export interface GuildHomeReport {
  id: string;
  title: string;
  created_at: string;
  fight_count: number;
  kill_count: number;
}

/**
 * One roster row. No `user_id`: the real API's RosterRow never sends one (plan's
 * reconciliation record, item 3) — "is this my own row" is derived from
 * `Me.characters[].key` instead, which already exact-matches `character_key`.
 */
export interface GuildRosterRow {
  character_key: string;
  region: string;
  ruleset: string;
  name: string;
  class?: string;
  spec?: string;
  rank: GuildRank;
  verified: boolean;
  /** addon_exports.updated_at within the last 24h (spec section 4.1, RULING 9). */
  logged_recently: boolean;
  /** Present only at gear/gear_bags consent (spec section 3.2's fail-closed query rule). */
  item_level?: number;
  consent: GuildConsent;
}

export interface GuildHome {
  guild: GuildSummary;
  claim: ClaimStateView;
  reports: GuildHomeReport[];
  next_cursor?: string;
  roster: GuildRosterRow[];
}

export interface GuildSettingsData {
  default_visibility: GuildVisibility;
  officer_max_rank_index: number;
  claimed_by: MemberRef | null;
  claim_pending: ClaimPendingView | null;
  claim: ClaimStateView;
  invite: { rotated_at: string | null };
}

export interface ContestResult {
  status: 'contested';
}

export interface ApproveResult {
  character_key: string;
  status: 'approved';
}
```

Replace `fetchGuildHome`:

```ts
export function fetchGuildHome(
  guildId: number,
  cursor?: string,
  apiBase: string = API_BASE_URL,
): Promise<GuildHome> {
  const query = cursor === undefined ? '' : `?cursor=${encodeURIComponent(cursor)}`;
  return call<GuildHome>(`/v1/guilds/${guildId}/home${query}`, apiBase);
}
```

Add, after `releaseClaim`:

```ts
export function contestClaim(guildId: number, apiBase: string = API_BASE_URL): Promise<ContestResult> {
  return call<ContestResult>(`/v1/guilds/${guildId}/claim/contest`, apiBase, { method: 'POST' });
}
```

Replace `approveCharacter`/`removeCharacter` (the real router takes three literal path
segments, never a combined, encoded `character_key` — plan's reconciliation record,
item 4):

```ts
export function approveCharacter(
  guildId: number,
  region: string,
  ruleset: string,
  slug: string,
  apiBase: string = API_BASE_URL,
): Promise<ApproveResult> {
  return call<ApproveResult>(
    `/v1/guilds/${guildId}/characters/${region}/${ruleset}/${encodeURIComponent(slug)}/approve`,
    apiBase,
    { method: 'POST' },
  );
}

export function removeCharacter(
  guildId: number,
  region: string,
  ruleset: string,
  slug: string,
  apiBase: string = API_BASE_URL,
): Promise<RemovedResult> {
  return call<RemovedResult>(
    `/v1/guilds/${guildId}/characters/${region}/${ruleset}/${encodeURIComponent(slug)}`,
    apiBase,
    { method: 'DELETE' },
  );
}
```

(`region`/`ruleset` are always one of a small fixed vocabulary already validated by
`isRegion`/`isRuleset` upstream of every caller — only `slug` needs encoding, matching
how `fetchCharacter` in `rankings/api.ts` already builds the sibling
`/v1/characters/{region}/{ruleset}/{slug}` path.)

- [ ] **Step 4: Update `web/src/lib/account/api.ts`**

Change `MeGuild.verified?: boolean` to `verified: boolean` (the real API never omits
it — plan's reconciliation record, item 7). One-line change; if `astro check` surfaces
any call site that constructed a `MeGuild`-shaped literal without `verified`, fix each
one the same way Task 2 of the first plan fixed `GuildPage.guild.id` gaps.

- [ ] **Step 5: Add the new copy**

Add to `web/src/lib/guild/copy.ts`. Do not remove any existing entry unless a later
step in this plan says to.

```ts
// Add to guildHomeCopy:
untitledReport: 'Untitled report',
contestButton: 'Contest this claim',
contestConfirmLine:
  'Contesting freezes this guild’s officer tools until a moderator reviews it. Use this only when the account that claimed this guild is not its guild master.',
contestConfirmButton: 'Yes, contest this claim',
contested: 'This claim is with a moderator.',
frozenNotice:
  'This guild’s claim is contested. Officer actions are frozen until a moderator resolves it.',
reportsUnverifiedNote: 'You are not verified yet, so only public reports show here.',
```

```ts
// Add to guildClaimCopy:
rulesHeading: 'How claiming works',
rules: [
  'Claiming needs a Battle.net-linked account.',
  'You may attempt one claim every 30 days.',
  'You may hold only one claimed guild at a time.',
  'A claim can be contested.',
] as readonly string[],
contested: 'This guild’s claim is contested and under review by a moderator.',
```

```ts
// Add to guildSettingsCopy:
frozenNotice:
  'This guild’s claim is contested. Officer actions are frozen until a moderator resolves it.',
```

- [ ] **Step 6: Run to verify pass, then lint/typecheck/format**

Run: `cd web && npx vitest run src/lib/guild/api.test.ts && npx astro check && npm run lint && npx prettier --check src/lib/guild/api.ts src/lib/guild/copy.ts src/lib/account/api.ts`

- [ ] **Step 7: Commit**

```bash
cd /Users/jh/code/forever/.worktrees/guild-web
printf 'fix(web): reconcile guild API client with the real, landed api/internal/guilds handlers\n\nField-exact fixes: flat home.reports + top-level next_cursor (was nested),\nno zone/wipe_count/user_id on real response types, three-segment\napprove/remove paths (was a combined encoded character_key), settings\nclaim_pending is an object not a boolean, both home and settings gain\nclaim:{state,since?}, new contestClaim(). MeGuild.verified is required,\nnot optional, matching auth.Store.Guild'"'"'s json tag. New copy for the\nclaim rules blurb, contest confirm/aftermath, and the frozen-claim notice.\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-r1.txt
git add web/src/lib/guild/api.ts web/src/lib/guild/api.test.ts web/src/lib/guild/copy.ts web/src/lib/account/api.ts
git commit -F .superpowers/commit-msg-r1.txt
```

---

## Task 2: `Guild.svelte` — claim state, contest, frozen controls, roster rules, report copy

**Files:**
- Modify: `web/src/components/Guild.svelte`
- Modify: `web/src/components/Guild.test.ts`

**Interfaces consumed:** everything from Task 1 (`fetchGuildHome`, `approveCharacter`,
`removeCharacter`, `contestClaim`, `ClaimStateView`, `GuildHome`, `GuildRosterRow`,
`guildHomeCopy`), plus already-existing `me.characters: MeCharacter[]`,
`me.guilds: MeGuild[]`, `me.user.role`.

- [ ] **Step 1: Read the current file fully** (already read during planning — 340
  lines). Do not change the public-page section (`data`/`status`/`error`, the
  progression/roster-bests sections) or any existing testid.

- [ ] **Step 2: Write the failing test additions**

Add to `web/src/components/Guild.test.ts` (keep the existing two tests):

```ts
it('still renders no guild-home before session resolves, unaffected by the reconcile', () => {
  const { body } = render(Guild, { props: { path: PATH } });
  expect(body).not.toContain('data-testid="guild-home"');
});
```

(This is a smoke check that the reconciled component still compiles and keeps its
static-render contract; the real behavioral coverage for claim states, contest, and
the frozen banner is e2e — Task 5.)

- [ ] **Step 3: Run to verify it fails to compile against the old code (RED via
  type-check)**, then apply the rewrite below and re-run.

Run: `cd web && npx vitest run src/components/Guild.test.ts` (expected to fail on the
still-old component referencing `home.viewer`/`row.user_id`, which Task 1 already
removed from the types — this is your RED signal, a type error, not a runtime one).

- [ ] **Step 4: Rewrite the relevant parts of `Guild.svelte`**

Replace the membership-derivation and roster/report state block (everything from `let
home = $state...` through the `soloRoster` derivation) with:

```ts
  let home = $state<GuildHome | null>(null);
  let homeStatus = $state<'idle' | 'loading' | 'ready'>('idle');
  let rosterBusy = $state<string | null>(null); // character_key currently being approved/removed
  let rosterActionError = $state('');
  let contestBusy = $state(false);
  let contestError = $state('');
  let showContestConfirm = $state(false);
  let contested = $state(false); // locally flipped true right after a successful contest call

  let myGuildMembership = $state<{ rank: string; verified: boolean } | null>(null);
  let myCharacterKeys = $state<Set<string>>(new Set());
  let isModerator = $state(false);

  $effect(() => {
    const requested = resolved;
    if (requested === null) return;
    homeStatus = 'loading';
    void fetchMeOnce()
      .then((result) => {
        if (resolved !== requested) return;
        const membership = result?.guilds.find(
          (g) =>
            g.region === requested.region &&
            g.ruleset === requested.ruleset &&
            characterSlug(g.name) === requested.slug,
        );
        myCharacterKeys = new Set((result?.characters ?? []).map((c) => c.key));
        isModerator = result?.user.role === 'moderator' || result?.user.role === 'admin';
        if (membership === undefined) {
          homeStatus = 'idle';
          return null;
        }
        myGuildMembership = { rank: membership.rank ?? 'member', verified: membership.verified };
        return fetchGuildHome(membership.id).then((page) => {
          if (resolved !== requested) return;
          home = page;
          contested = page.claim.state === 'contested';
          homeStatus = 'ready';
        });
      })
      .catch(() => {
        if (resolved !== requested) return;
        homeStatus = 'idle';
      });
  });

  const canManage = $derived(
    myGuildMembership !== null &&
      myGuildMembership.verified &&
      (myGuildMembership.rank === 'officer' || myGuildMembership.rank === 'leader'),
  );

  /** A raw membership row (any rank, verified or not) is enough to offer contesting —
   *  the API enforces the real officer/leader-or-rank-0 rule server side (spec section
   *  3.3's amendment); a plain member who tries gets a 403 with a sentence. */
  const canContest = $derived(myGuildMembership !== null);

  async function onApprove(row: GuildRosterRow): Promise<void> {
    if (home === null) return;
    rosterBusy = row.character_key;
    rosterActionError = '';
    try {
      await approveCharacter(home.guild.id, row.region, row.ruleset, characterSlug(row.name));
      home = {
        ...home,
        roster: home.roster.map((r) => (r.character_key === row.character_key ? { ...r, verified: true } : r)),
      };
    } catch (thrown) {
      rosterActionError = thrown instanceof Error ? thrown.message : 'That did not work; try again';
    } finally {
      rosterBusy = null;
    }
  }

  async function onRemove(row: GuildRosterRow): Promise<void> {
    if (home === null) return;
    rosterBusy = row.character_key;
    rosterActionError = '';
    try {
      await removeCharacter(home.guild.id, row.region, row.ruleset, characterSlug(row.name));
      home = { ...home, roster: home.roster.filter((r) => r.character_key !== row.character_key) };
    } catch (thrown) {
      rosterActionError = thrown instanceof Error ? thrown.message : 'That did not work; try again';
    } finally {
      rosterBusy = null;
    }
  }

  /** Whether the Remove control should even render for this row. Never shows a control
   *  the API would refuse for the two absolute cases (a leader-rank row, an
   *  officer-rank row) unless the viewer is the row's own account or a moderator — the
   *  conservative reading of rank-protects-rank (plan's reconciliation record, item 10):
   *  the web has no signal for "am I the claim holder" from the home response alone, so
   *  an officer-rank row's Remove is hidden even from a genuine claim-holder officer,
   *  who still has the settings page or the API directly. */
  function mayShowRemove(row: GuildRosterRow): boolean {
    const isSelf = myCharacterKeys.has(row.character_key);
    if (isSelf || isModerator) return true;
    if (row.rank === 'member') return canManage;
    return false; // officer or leader rank, not self, not moderator
  }

  async function onContest(): Promise<void> {
    if (home === null) return;
    contestBusy = true;
    contestError = '';
    try {
      await contestClaim(home.guild.id);
      contested = true;
      showContestConfirm = false;
    } catch (thrown) {
      contestError = thrown instanceof Error ? thrown.message : 'That did not work; try again';
    } finally {
      contestBusy = false;
    }
  }

  function rowPath(row: GuildRosterRow): CharacterPath {
    return { region: row.region as Region, ruleset: row.ruleset as Ruleset, slug: characterSlug(row.name) };
  }

  /** "Am I the guild's only member" is counted by distinct account, not row count: a
   *  solo officer with two characters is still one account. `Me.characters[].key`
   *  exact-matches `character_key`, so any roster row not in that set belongs to
   *  someone else. */
  const soloRoster = $derived(
    home === null ? true : home.roster.every((row) => myCharacterKeys.has(row.character_key)),
  );

  const reportWipeCount = (report: GuildHomeReport): number =>
    Math.max(0, report.fight_count - report.kill_count);
```

Add the import for `GuildHomeReport` and `contestClaim` to the existing `../lib/guild/api`
import line, and add `Guild.svelte`'s import list needs `type Me` no longer if unused —
check before removing; `fetchMeOnce` stays.

- [ ] **Step 5: Rewrite the signed-in section's template**

Replace the entire `{#if homeStatus === 'ready' && home !== null}` block with:

```svelte
    {#if homeStatus === 'ready' && home !== null}
      <section class="border-line-soft flex flex-col gap-4 border-b pb-6" data-testid="guild-home">
        <div class="flex flex-wrap items-center justify-between gap-3">
          <h2 class="section-title text-[18px]">{guildHomeCopy.reportsHeading}</h2>
          {#if canManage}
            <a
              class="text-[13px] font-semibold"
              href={guildSettingsHref(resolved.region, resolved.ruleset, home.guild.name)}
              data-testid="guild-settings-link"
            >
              {guildHomeCopy.settingsLink}
            </a>
          {:else if myGuildMembership !== null && !myGuildMembership.verified}
            <a
              class="text-[13px] font-semibold"
              href={guildClaimHref(resolved.region, resolved.ruleset, home.guild.name)}
              data-testid="guild-claim-link"
            >
              {guildHomeCopy.claimLink}
            </a>
          {/if}
        </div>

        {#if contested}
          <p class="text-[13px]" role="alert" data-testid="guild-home-frozen">{guildHomeCopy.frozenNotice}</p>
        {/if}

        {#if home.claim.state !== 'unclaimed' || canContest}
          <div class="flex flex-wrap items-center gap-3">
            {#if home.claim.state === 'contested'}
              <span class="text-muted text-[13px]" data-testid="guild-home-claim-state">
                {guildHomeCopy.contested}
              </span>
            {/if}
            {#if canContest && home.claim.state !== 'unclaimed' && home.claim.state !== 'contested'}
              <button
                class="{SECONDARY_BUTTON_FIXED} border-line-warm text-text px-3"
                onclick={() => (showContestConfirm = true)}
                data-testid="guild-home-contest-button"
              >
                {guildHomeCopy.contestButton}
              </button>
            {/if}
          </div>
        {/if}

        {#if showContestConfirm}
          <div class="border-line-soft flex flex-col gap-3 border p-4" data-testid="guild-home-contest-confirm">
            <p class="text-[13px]">{guildHomeCopy.contestConfirmLine}</p>
            <div class="flex gap-3">
              <button
                class="{SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong px-3"
                onclick={() => void onContest()}
                disabled={contestBusy}
                data-testid="guild-home-contest-confirm-button"
              >
                {guildHomeCopy.contestConfirmButton}
              </button>
              <button
                class="{SECONDARY_BUTTON_FIXED} border-line-warm text-text px-3"
                onclick={() => (showContestConfirm = false)}
                disabled={contestBusy}
              >
                Cancel
              </button>
            </div>
            {#if contestError !== ''}<p class="text-[13px]" role="alert">{contestError}</p>{/if}
          </div>
        {/if}

        {#if myGuildMembership !== null && !myGuildMembership.verified}
          <p class="text-muted text-[13px]" data-testid="guild-home-unverified-note">
            {guildHomeCopy.reportsUnverifiedNote}
          </p>
        {/if}
        {#if home.reports.length === 0}
          <p class="text-muted text-[14px]" data-testid="guild-home-empty-reports">{guildHomeCopy.noReports}</p>
        {:else}
          <ul class="flex flex-col" data-testid="guild-home-reports">
            {#each home.reports as report (report.id)}
              <li class="border-line-soft flex min-h-11 flex-wrap items-center gap-3 border-b px-2 py-2 text-[14px]">
                <a class={rowLink} href={`/reports/${report.id}`}>
                  {report.title === '' ? guildHomeCopy.untitledReport : report.title}
                </a>
                <span class="text-muted tabular font-mono text-[13px]">
                  {guildHomeCopy.reportSummary(report.kill_count, reportWipeCount(report))}
                </span>
              </li>
            {/each}
          </ul>
        {/if}

        <h2 class="section-title text-[18px]">{guildHomeCopy.rosterHeading}</h2>
        {#if soloRoster}
          <p class="text-muted text-[14px]" data-testid="guild-home-empty-roster">
            {canManage ? guildHomeCopy.emptyRosterOfficer : guildHomeCopy.emptyRosterMember}
          </p>
        {:else}
          <ul class="flex flex-col" data-testid="guild-home-roster">
            {#each home.roster as row (row.character_key)}
              <li class="border-line-soft flex min-h-11 flex-wrap items-center gap-3 border-b px-2 py-2 text-[14px]">
                <a
                  class="{rowLink} truncate font-semibold"
                  style={`color: ${classColorVar(row.class ?? '')}`}
                  href={characterHref(row.region, row.ruleset, row.name)}
                >
                  {row.name}
                </a>
                <span class="text-muted text-[13px]">{row.spec ?? ''}</span>
                {#if row.logged_recently}
                  <span class="pill pill-site text-[11px]">{guildHomeCopy.loggedRecently}</span>
                {/if}
                {#if !row.verified}
                  <span class="pill pill-site text-[11px]" data-testid="guild-roster-unverified">
                    {guildHomeCopy.unverified}
                  </span>
                {/if}
                {#if row.item_level !== undefined}
                  <span class="text-muted tabular font-mono text-[13px]"
                    >{guildHomeCopy.itemLevelLabel} {row.item_level}</span
                  >
                {/if}
                {#if row.consent !== 'roster'}
                  <GuildRosterHandoff path={rowPath(row)} />
                {/if}
                {#if canManage && !row.verified}
                  <button
                    class="{SECONDARY_BUTTON_FIXED} border-line-warm text-text px-3"
                    onclick={() => void onApprove(row)}
                    disabled={rosterBusy === row.character_key || contested}
                    data-testid="guild-roster-approve"
                  >
                    {guildHomeCopy.approve}
                  </button>
                {/if}
                {#if mayShowRemove(row)}
                  <button
                    class="{SECONDARY_BUTTON_FIXED} border-line-warm text-text px-3"
                    onclick={() => void onRemove(row)}
                    disabled={rosterBusy === row.character_key ||
                      (contested && !myCharacterKeys.has(row.character_key) && !isModerator)}
                    data-testid="guild-roster-remove"
                  >
                    {guildHomeCopy.remove}
                  </button>
                {/if}
              </li>
            {/each}
          </ul>
        {/if}
        {#if rosterActionError !== ''}
          <p class="text-[13px]" role="alert" data-testid="guild-roster-action-error">{rosterActionError}</p>
        {/if}
      </section>
    {/if}
```

(Leaving one's own row is never frozen, even while contested — matches the API's own
rule, confirmed in `roster.go`'s `removeCharacter` handler comment — the `disabled`
expression above only adds the contested block for a non-self, non-moderator actor.)

- [ ] **Step 6: Run to verify pass**

Run: `cd web && npx vitest run src/components/Guild.test.ts`

- [ ] **Step 7: Run the pre-existing guild-phone e2e spec**

Run: `cd web && E2E_PORT=4421 npx playwright test tests/e2e/guild-phone.spec.ts`
Expected: PASS unchanged.

- [ ] **Step 8: Lint, typecheck, format, commit**

Run: `cd web && npx astro check && npm run lint && npx prettier --check src/components/Guild.svelte src/components/Guild.test.ts`

```bash
cd /Users/jh/code/forever/.worktrees/guild-web
printf 'feat(web): claim state, contest, frozen controls and roster rank rules on Guild.svelte\n\nDerives can-manage/is-self/is-moderator from Me (home.viewer never\nexisted on the real API); hides Remove on officer/leader rows for\nanyone but self or a moderator (conservative reading of\nrank-protects-rank, no claim-holder signal on the home response yet);\ndisables approve/remove/contest-adjacent controls while a claim is\ncontested and shows why; adds the unverified-member reports note.\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-r2.txt
git add web/src/components/Guild.svelte web/src/components/Guild.test.ts
git commit -F .superpowers/commit-msg-r2.txt
```

---

## Task 3: `GuildClaim.svelte` — rules blurb, contest, contested state

**Files:**
- Modify: `web/src/components/GuildClaim.svelte`
- Modify: `web/src/components/GuildClaim.test.ts`

- [ ] **Step 1: Write the failing test addition**

Add to `GuildClaim.test.ts` (keep the existing loading-state test):

```ts
it('still renders the guild-claim testid unaffected by the reconcile', () => {
  const { body } = render(GuildClaim, { props: { path: PATH } });
  expect(body).toContain('data-testid="guild-claim"');
});
```

- [ ] **Step 2: Rewrite `GuildClaim.svelte`**

Update the type import (`fetchGuildSettings` return type now carries `claim` and the
real `claim_pending` shape — no code change needed for the fetch itself, only for how
the template reads it). Add `contestClaim` to the imports from `../lib/guild/api`.

Replace the `justClaimedExpiry` local-only pending display with reading
`settings.claim_pending?.expires_at` directly (now always present when pending, not
session-only) — simplifies `guildClaimCopy.pending` to take a required string, not
optional. Update `guildClaimCopy.pending` in Task 1's copy step retroactively is not
needed since Task 1 already shipped; instead keep `pending(expiresAt?: string)`
exactly as-is (it already handles both cases) and simply always pass
`settings.claim_pending?.expires_at` (works whether it's the fixed value or
`undefined`).

Add contest state and action, mirroring `onRelease`'s shape:

```ts
  let showContestConfirm = $state(false);

  const onContest = (): void =>
    void run(async () => {
      if (guildId === null) return;
      await contestClaim(guildId);
      settings = await fetchGuildSettings(guildId);
      showContestConfirm = false;
    });
```

Add the rules blurb above the claim-state paragraph (inside the `{:else if settings
!== null}` branch, before the `{#if settings.claimed_by !== null}` chain):

```svelte
    <section class="flex flex-col gap-2" data-testid="guild-claim-rules">
      <h2 class="section-title text-[16px]">{guildClaimCopy.rulesHeading}</h2>
      <ul class="text-muted flex flex-col gap-1 text-[13px]">
        {#each guildClaimCopy.rules as rule (rule)}
          <li>{rule}</li>
        {/each}
      </ul>
    </section>
```

Add a `contested` branch as the FIRST arm of the existing
`{#if settings.claimed_by !== null} ... {:else if settings.claim_pending !== null} ...
{:else} ...` chain (contested takes priority over claimed/pending/unclaimed display,
since `claim.state` is the single source of truth for which of the four states this
is — `claimed_by`/`claim_pending` can both still be non-null while contested):

```svelte
    {#if settings.claim.state === 'contested'}
      <p class="text-[14px]" data-testid="guild-claim-state">{guildClaimCopy.contested}</p>
    {:else if settings.claimed_by !== null}
```

(keep the rest of the existing chain unchanged below this new first arm, changing only
`settings.claim_pending` from a truthy-boolean check to `settings.claim_pending !==
null`, and `justClaimedExpiry` references to `settings.claim_pending?.expires_at`).

Add a Contest button and confirm step, shown when `home`-equivalent eligibility exists
— reuse the same `membership`/`eligible`-adjacent pattern already in this file, but
per the coordinator's instruction, contest is offered to ANY signed-in visitor with a
character in this guild (any rank, not officer/leader-only like claim/confirm):

```ts
  const canContest = $derived(membership !== null);
```

Add the button + confirm block near the bottom of the template, inside `{:else if
settings.claimed_by !== null}` and `{:else if settings.claim_pending !== null}`
branches both (contesting a claimed OR a pending claim is valid per the API):

```svelte
      {#if canContest}
        {#if !showContestConfirm}
          <button
            class="{SECONDARY_BUTTON_FIXED} border-line-warm text-text w-fit px-3"
            onclick={() => (showContestConfirm = true)}
            data-testid="guild-claim-contest-button"
          >
            {guildHomeCopy.contestButton}
          </button>
        {:else}
          <div class="border-line-soft flex flex-col gap-3 border p-4" data-testid="guild-claim-contest-confirm">
            <p class="text-[13px]">{guildHomeCopy.contestConfirmLine}</p>
            <button
              class="{SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong w-fit px-3"
              onclick={onContest}
              disabled={busy}
              data-testid="guild-claim-contest-confirm-button"
            >
              {guildHomeCopy.contestConfirmButton}
            </button>
          </div>
        {/if}
      {/if}
```

Import `guildHomeCopy` alongside `guildClaimCopy` for the shared contest strings (both
`Guild.svelte` and `GuildClaim.svelte` show the identical contest confirm line — kept
in `guildHomeCopy` as the one source, per the "no duplication" rule, since the home
section defined it first in Task 2).

- [ ] **Step 3: Run to verify pass**

Run: `cd web && npx vitest run src/components/GuildClaim.test.ts`

- [ ] **Step 4: Lint, typecheck, format, commit**

Run: `cd web && npx astro check && npm run lint && npx prettier --check src/components/GuildClaim.svelte src/components/GuildClaim.test.ts`

```bash
cd /Users/jh/code/forever/.worktrees/guild-web
printf 'feat(web): claim rules blurb, contest and contested state on GuildClaim.svelte\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-r3.txt
git add web/src/components/GuildClaim.svelte web/src/components/GuildClaim.test.ts
git commit -F .superpowers/commit-msg-r3.txt
```

---

## Task 4: `GuildSettings.svelte` — claim state, frozen controls, `claim_pending` shape fix

**Files:**
- Modify: `web/src/components/GuildSettings.svelte`
- Modify: `web/src/components/GuildSettings.test.ts`

- [ ] **Step 1: Write the failing test addition** (same smoke-test pattern as Task 2/3
  — one line confirming `data-testid="guild-settings"` still renders).

- [ ] **Step 2: Rewrite the relevant parts of `GuildSettings.svelte`**

Add a `contested` derived value:

```ts
  const contested = $derived(settings?.claim.state === 'contested');
```

Show the claim state and frozen notice near the top of the `{:else if settings !==
null}` branch, before the default-visibility section:

```svelte
    <p class="text-[13px]" data-testid="guild-settings-claim-state">
      Claim: {settings.claim.state}
    </p>
    {#if contested}
      <p class="text-[13px]" role="alert" data-testid="guild-settings-frozen">
        {guildSettingsCopy.frozenNotice}
      </p>
    {/if}
```

Disable the visibility select, the officer-threshold input, and the rotate button when
`contested`, in addition to the existing `disabled={busy}`:

```svelte
        disabled={busy || contested}
```

(apply to all three controls). The existing `run()`/`onSave`/`onRotate` error handling
already surfaces a 409 `claim_contested` the same as any other thrown `GuildApiError`
via `thrown.message` — no code change needed there, since the API's own message
("this guild's claim is contested; officer actions are frozen until a moderator
resolves it") is already a good, displayable sentence; this satisfies "still handle
the 409 if it arrives" without a special-cased branch.

- [ ] **Step 3: Run to verify pass**

Run: `cd web && npx vitest run src/components/GuildSettings.test.ts`

- [ ] **Step 4: Lint, typecheck, format, commit**

Run: `cd web && npx astro check && npm run lint && npx prettier --check src/components/GuildSettings.svelte src/components/GuildSettings.test.ts`

```bash
cd /Users/jh/code/forever/.worktrees/guild-web
printf 'feat(web): claim state and frozen controls on GuildSettings.svelte\n\nAlso fixes claim_pending from a boolean (this lane'"'"'s own earlier guess)\nto the real ClaimPendingView object the API actually sends.\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-r4.txt
git add web/src/components/GuildSettings.svelte web/src/components/GuildSettings.test.ts
git commit -F .superpowers/commit-msg-r4.txt
```

---

## Task 5: e2e spec reconciliation and the whole-suite run

**Files:**
- Modify: `web/tests/e2e/guild-home.spec.ts`
- Modify: `web/tests/e2e/guild-claim.spec.ts`
- Modify: `web/tests/e2e/guild-settings.spec.ts`
- Modify: `web/tests/e2e/guild-account-consent.spec.ts` (only if its `ME` fixture is
  missing `verified` on a guild entry — check first; add `verified: true` where needed
  now that the field is required)
- Modify: `web/tests/e2e/guild-nav-home-link.spec.ts` (same check as above)
- Modify: `web/tests/e2e/guild-invite.spec.ts` (only if it references
  `character_key`-shaped approve/remove stubs — check first; it should not, since
  invite has no roster actions)

- [ ] **Step 1: Fix every fixture object in `guild-home.spec.ts`**

Before writing new tests, read the file's current `HOME`/`ME` fixtures fully (already
read during planning — `HOME.viewer`, `HOME.reports.rows`, `wipe_count`/`zone`, and
`user_id` on roster rows are all wrong per Task 1's reconciliation record). Rewrite
`HOME` to:

```ts
const HOME = {
  guild: { id: 501, name: 'The Last Watch', region: 'us', ruleset: 'hardcore' },
  claim: { state: 'claimed' },
  reports: [
    { id: 'fixture2abcd', title: 'Sanguine Depths', created_at: '2026-09-20T20:00:00Z', fight_count: 8, kill_count: 3 },
  ],
  roster: [
    { character_key: 'us/hardcore/simfury', region: 'us', ruleset: 'hardcore', name: 'Simfury', class: 'Warrior', spec: 'Fury', rank: 'officer', verified: true, logged_recently: true, consent: 'gear' },
    { character_key: 'us/hardcore/newbie', region: 'us', ruleset: 'hardcore', name: 'Newbie', class: 'Priest', spec: 'Holy', rank: 'member', verified: false, logged_recently: false, consent: 'roster' },
  ],
};
```

Add `verified: true` to every `ME.guilds[]` entry across this file (now required).

Update the approve-character `page.route` stub from
`**/v1/guilds/501/characters/us%2Fhardcore%2Fnewbie/approve` to
`**/v1/guilds/501/characters/us/hardcore/newbie/approve`, and its fulfillment body to
`{ character_key: 'us/hardcore/newbie', status: 'approved' }` (matching Task 1's
`ApproveResult`).

Update the "3 kills · 5 wipes" assertion: with `fight_count: 8, kill_count: 3`, the
derived wipe count is `8 - 3 = 5` — text stays `'3 kills · 5 wipes'`, only the
underlying fixture fields changed shape.

- [ ] **Step 2: Add new e2e cases to `guild-home.spec.ts`**

```ts
test('a contested claim shows the frozen notice and disables approve/remove', async ({ page }) => {
  await page.route('**/v1/guilds/us/hardcore/the-last-watch', (route) => route.fulfill(envelope(GUILD_PAGE)));
  await page.route('**/v1/me', (route) => route.fulfill(envelope(ME)));
  await page.route('**/v1/guilds/501/home', (route) =>
    route.fulfill(envelope({ ...HOME, claim: { state: 'contested', since: '2026-09-19T00:00:00Z' } })),
  );
  await page.goto('/guild/us/hardcore/the-last-watch');
  await expect(page.getByTestId('guild-home-frozen')).toBeVisible();
  await expect(page.getByTestId('guild-roster-approve')).toBeDisabled();
});

test('an officer sees no remove control on another officer’s row, but sees it on a member row and their own row', async ({ page }) => {
  const ME_WITH_OWN_CHAR = {
    ...ME,
    characters: [{ key: 'us/hardcore/simfury', region: 'us', ruleset: 'hardcore', name: 'Simfury', class: 'Warrior' }],
  };
  await page.route('**/v1/guilds/us/hardcore/the-last-watch', (route) => route.fulfill(envelope(GUILD_PAGE)));
  await page.route('**/v1/me', (route) => route.fulfill(envelope(ME_WITH_OWN_CHAR)));
  await page.route('**/v1/guilds/501/home', (route) =>
    route.fulfill(
      envelope({
        ...HOME,
        roster: [
          ...HOME.roster,
          { character_key: 'us/hardcore/otherofficer', region: 'us', ruleset: 'hardcore', name: 'OtherOfficer', class: 'Mage', spec: 'Frost', rank: 'officer', verified: true, logged_recently: false, consent: 'roster' },
        ],
      }),
    ),
  );
  await page.goto('/guild/us/hardcore/the-last-watch');
  const ownRow = page.getByTestId('guild-home-roster').getByRole('listitem').filter({ hasText: 'Simfury' });
  await expect(ownRow.getByTestId('guild-roster-remove')).toBeVisible();
  const memberRow = page.getByTestId('guild-home-roster').getByRole('listitem').filter({ hasText: 'Newbie' });
  await expect(memberRow.getByTestId('guild-roster-remove')).toBeVisible();
  const otherOfficerRow = page.getByTestId('guild-home-roster').getByRole('listitem').filter({ hasText: 'OtherOfficer' });
  await expect(otherOfficerRow.getByTestId('guild-roster-remove')).toHaveCount(0);
});

test('an unverified member sees the public-reports-only note', async ({ page }) => {
  const ME_UNVERIFIED = { ...ME, guilds: [{ ...ME.guilds[0], rank: 'member', verified: false }] };
  await page.route('**/v1/guilds/us/hardcore/the-last-watch', (route) => route.fulfill(envelope(GUILD_PAGE)));
  await page.route('**/v1/me', (route) => route.fulfill(envelope(ME_UNVERIFIED)));
  await page.route('**/v1/guilds/501/home', (route) => route.fulfill(envelope(HOME)));
  await page.goto('/guild/us/hardcore/the-last-watch');
  await expect(page.getByTestId('guild-home-unverified-note')).toBeVisible();
});

test('contesting a claim calls the contest endpoint after confirming', async ({ page }) => {
  await page.route('**/v1/guilds/us/hardcore/the-last-watch', (route) => route.fulfill(envelope(GUILD_PAGE)));
  await page.route('**/v1/me', (route) => route.fulfill(envelope(ME)));
  await page.route('**/v1/guilds/501/home', (route) => route.fulfill(envelope(HOME)));
  await page.route('**/v1/guilds/501/claim/contest', (route) => route.fulfill(envelope({ status: 'contested' })));
  await page.goto('/guild/us/hardcore/the-last-watch');
  await page.getByTestId('guild-home-contest-button').click();
  await page.getByTestId('guild-home-contest-confirm-button').click();
  await expect(page.getByTestId('guild-home-frozen')).toBeVisible();
});
```

(Adjust selectors to whatever Task 2 actually shipped if a name drifted while writing
the component — the implementer for this task must read the real, committed
`Guild.svelte` testids rather than trust this brief blindly, per the standing
instruction from the first plan's Task 15 that this reconcile inherits.)

- [ ] **Step 3: Fix `guild-claim.spec.ts` and `guild-settings.spec.ts` fixtures**

In both files: change every `claimed_by`/`claim_pending` fixture object to include the
new `claim: {state, since?}` field consistently (`state` matching whichever of
`claimed_by`/`claim_pending` is set), and change any `claim_pending: true`-style
boolean fixture to the real `{by: {battletag}, expires_at}` object shape. Add one new
test to `guild-claim.spec.ts` proving the contested state renders
`guildClaimCopy.contested`'s text when `settings.claim.state === 'contested'`. Add one
new test to `guild-settings.spec.ts` proving the frozen notice appears and the
visibility select is disabled when `claim.state === 'contested'`.

- [ ] **Step 4: Fix `verified` on any other spec's `ME`/`Me`-shaped fixture**

Run: `cd web && grep -rln "guilds:" tests/e2e/guild-*.spec.ts` and check each hit's
guild entries for a missing `verified` field; add `verified: true` (or `false` as the
test needs) to each.

- [ ] **Step 5: Run every guild e2e spec**

Run: `cd web && E2E_PORT=4421 npx playwright test tests/e2e/guild-home.spec.ts tests/e2e/guild-claim.spec.ts tests/e2e/guild-settings.spec.ts tests/e2e/guild-invite.spec.ts tests/e2e/guild-nav-home-link.spec.ts tests/e2e/guild-account-consent.spec.ts tests/e2e/guild-phone.spec.ts`
Expected: all green.

- [ ] **Step 6: Run the full scoped vitest/astro-check/lint sweep**

Run: `cd web && npx vitest run src/lib/guild src/lib/characters.test.ts src/lib/rankings/api.test.ts src/lib/report/og-meta.test.ts src/lib/report/shell-paths.test.ts src/worker.test.ts src/components/Guild.test.ts src/components/GuildRosterHandoff.test.ts src/components/GuildClaim.test.ts src/components/GuildSettings.test.ts src/components/GuildJoin.test.ts src/components/GuildShell.test.ts src/components/SessionNav.test.ts src/components/HomeGuildLink.test.ts src/components/Account.test.ts && npx astro check && npm run lint`

- [ ] **Step 7: Run the WHOLE e2e suite once** (the journey's own house rule, and the
  coordinator's explicit instruction for this round)

Run: `cd web && E2E_PORT=4421 npx playwright test`

For every failure, classify it:
- If the failing spec file (or the component/lib file whose behavior it tests) has
  zero diff between this branch and `main` — check with
  `git -C /Users/jh/code/forever/.worktrees/guild-web diff main...HEAD --stat -- <path>`
  returning nothing — it is pre-existing; name it in the report as "pre-existing,
  confirmed: <path> has no diff vs main" and do not fix it.
- Otherwise it is this lane's own regression from this reconcile: fix it before
  reporting DONE.

- [ ] **Step 8: Lint, typecheck, format, commit**

Run: `cd web && npx prettier --check tests/e2e/guild-*.spec.ts`

```bash
cd /Users/jh/code/forever/.worktrees/guild-web
printf 'test(web): reconcile guild e2e fixtures with the real API shapes\n\nFlat home.reports, claim:{state,since?} on both home and settings, the\nreal claim_pending object shape, three-segment approve/remove stubs, new\ncoverage for the contested/frozen state, rank-protects-rank remove\nvisibility, and the unverified-member reports note.\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-r5.txt
git add web/tests/e2e/guild-home.spec.ts web/tests/e2e/guild-claim.spec.ts web/tests/e2e/guild-settings.spec.ts web/tests/e2e/guild-account-consent.spec.ts web/tests/e2e/guild-nav-home-link.spec.ts web/tests/e2e/guild-invite.spec.ts
git commit -F .superpowers/commit-msg-r5.txt
```

---

### Final integration (part of this reconcile round, not a separate numbered task)

After Task 5, this round's own whole-suite run (Task 5 Step 7) already satisfies the
journey's "run the whole suite once before the final report" requirement — do not run
it a second time. Proceed straight to the final report: branch head SHA, every
field-level mismatch found (the reconciliation record above) and how each was
resolved, the two required-API-change recommendations (item 2's missing `zone` on
`HomeReport`, item 10's missing claim-holder identity on `HomeView`), and the full-suite
result with every failure named and classified.
