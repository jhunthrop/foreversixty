# Guild membership from the addon, and the free guild home

**Date:** 2026-09-21
**Status:** DRAFT for coordinator review. Implements proposal sections 1, 2, 3.1, 3.2 and
3.6 only (`2026-09-21-guild-and-premium-proposal.md`). Payments, officer tools (3.3), the
performance analyzer (3.4) and in-game ratings (3.5) are out of scope — section 8.

Every **RULING** below is a call this spec had to make where the proposal or the existing
code left a gap. Each names its reason and the cost of being wrong, so the coordinator can
overrule it cheaply.

## 0. What exists today (read, not built here)

- `guilds` and `guild_members` tables exist and nothing writes to `guild_members`
  (`api/internal/db/migrations/0005_logs.up.sql`). `guilds` has `claimed_by`,
  `default_visibility` (default `'public'`), a `ruleset` check
  (`normal|pvp|rp|hardcore`). `guild_members` has `(guild_id, user_id)` primary key,
  `rank text default 'member'`, `refreshed_at`.
- `addon_exports` (`character_key` primary key, `user_id`, `region`, `ruleset`, `name`,
  `export`, `updated_at`) is the only server-side store of an export string. It is written
  by exactly one path: `Store.PutExports` (`api/internal/addon/addon.go:87`), reached by
  `POST /v1/addon/exports`, mounted `auth.RequireDevice` (`api/internal/addon/addon.go:198`)
  — a paired companion install only. `MaxExportLen = 4096` per character
  (`api/internal/addon/addon.go:35`); a guild section is ~80 bytes worst case, no change
  needed. The planner's paste-and-decode path (`fs1.ts`) never uploads anything; it is
  client-only. This is load-bearing for section 3's trust rule.
- `GET /v1/me` already returns `guilds: []Guild` with `Rank` per guild
  (`api/internal/auth/handler.go:270`, `auth.Store.Guilds`, `api/internal/auth/store.go:391`).
- Report visibility (`api/internal/reports/report.go:16`): `public | unlisted | private |
  guild`. Today, `mayView` (`api/internal/reports/handler.go:521`) grants a `guild`-visible
  report to **any** row in `guild_members` for that guild, regardless of rank —
  `s.Accounts.GuildRank(ctx, guildID, userID)` only checks existence
  (`api/internal/auth/store.go:413`). `mayEdit` (`handler.go:543`) requires rank `officer`
  or `leader` (`handler.go:29`). Attaching a report to a guild (`PATCH .guild_id`,
  `handler.go:307-326`) uses the same raw `GuildRank` check. **None of this has ever been
  exercised in production because `guild_members` has always been empty** — this spec is
  what activates it, which is why section 3's security review is not optional.
- The public guild page (`api/internal/rankings/guilds.go:177`, `GET
  /v1/guilds/{region}/{ruleset}/{name}`) already builds progression and roster-bests from
  every report with that `guild_id` and `visibility <> 'private'` — i.e. `guild`-visible
  reports already contribute kill counts and roster-best rows to the **public** page today
  once any exist. This spec does not change that; it is noted in section 4 so nobody
  mistakes it for new exposure.
- FS1 version 2 (`docs/superpowers/specs/2026-09-19-simulator-parity-interfaces.md` §7,
  `web/src/lib/planner/fs1.ts`, `addon/ForeverSixty/Codec.lua`) is a fixed-order sequence of
  optional pipe-delimited sections after the version-1 head; an unknown section is ignored
  and named in `ignored`. `bags`, `bank`, `sets`, `loadouts`, `professions` exist. This is
  where `guild` joins.
- `web/src/lib/current-character.ts` is landing on `main` from another lane
  (`2026-09-21-one-product-design.md` Lane B) around the same time as this spec.
  `web/src/lib/handoff-links.ts` exists on `main` today and is the URL-building module this
  spec falls back to until then. §4.1 specifies one helper that picks whichever is present,
  so the roster's own code never branches on it.
- `/guild/<region>/<ruleset>/<name>` is served by the Worker's OG-rewriting shell
  (`web/src/worker.ts:358`, prefix `/guild/` → `/guild.html`), not by
  `web/src/pages/guild/[...path].astro` (fixture-only, for Playwright). New guild sub-paths
  in section 5 reuse this same prefix rule; the shell's `shellHead` (`worker.ts:387`) needs
  to recognise them (small, explicit task, not assumed here).

## 1. The addon: FS1 gains a `guild` section

### 1.1 The WoW API call

`GetGuildInfo("player")` returns `guildName, guildRankName, guildRankIndex[, guildRealm]`
(unchanged across every WoW client generation, so no beta-only risk); it returns `nil` when
the unit is not in a guild. `guildRankIndex` is 0-based; **0 is always the guild master**
(server-authoritative — a player cannot make themselves rank 0 without being GM). Rank
**names** (`guildRankName`) are player-chosen free text and are **not** used for anything —
only the index is trustworthy and stable.

`Export.guildInfo()` (new, `addon/ForeverSixty/Export.lua`, beside `Export.professionSlugs`):

```lua
function Export.guildInfo()
  local name, _, rankIndex = GetGuildInfo("player")
  if name == nil then
    return nil
  end
  return { name = name, rankIndex = rankIndex }
end
```

`Export.string` passes `guild = Export.guildInfo()` into `Codec.encodeFS1`, beside the
existing `professions = Export.professionSlugs()` line. An unguilded character produces no
`guild` field, exactly as an empty `professions` list is omitted — the encoder must never
write a `guild=` section with an empty name.

### 1.2 Wire syntax (FS1 version 2, section 6 of the fixed order)

```
FS1:<build>:<class>:<race>:<t1>/<t2>/<t3>:<gear>|bags=<items>|bank=<items>|sets=<...>|loadouts=<...>|professions=<slugs>|guild=<name>:<rank-index>
```

- New section name `guild`, appended **after** `professions` — the contract's order is
  "fixed as written" and unknown sections are ignored by name, so appending is the only
  change that cannot break an already-shipped decoder reading a version-1 or pre-`guild`
  version-2 string (the existing "unknown section is ignored and named" case,
  `addon/tests/vectors_spec.lua` via `codec-vectors.json`'s `pets=1,2` vector, already
  proves this generically; no new decoder-compatibility test is needed for that half).
- Payload is `<url-encoded-guild-name>:<rank-index>` — **not** the `name=payload;...`
  named-entries grammar `sets`/`loadouts` use, because a guild is one fact, not a
  collection. `<name>` is percent-encoded with the same `urlEncode`/`encodeURIComponent`
  functions `sets`/`loadouts` already use (`Codec.lua`'s `urlEncode`,
  `fs1.ts`'s `encodeURIComponent`), which both already escape `:` — RFC 3986's unreserved
  set is alnum, `-`, `.`, `_`, `~` only, so a guild name containing a literal colon is
  guaranteed pre-escaped to `%3A` on both sides before this format is ever split on the
  first unescaped `:`. `<rank-index>` is digits-only (the existing `isDigits` grammar item
  ids and profession-free-of-punctuation checks already use); no rank index is ever
  negative.
- Decode: split the section's payload on the **first** `:`; URL-decode the left half with
  the existing `decodeName`/`urlDecode` helper (never throws, per `fs1.ts`'s own comment on
  why); the right half must satisfy `isDigits` (Lua) / `/^\d+$/` (TS) or the whole code is
  refused with a new message (`L.codecGuildRank` in `Locale.lua`, mirroring
  `L.codecGearEntry`'s "unreadable X entry" phrasing) — **not** silently dropped, because a
  malformed `guild` section is evidence of a hand-edited string, and the existing codec's
  rule throughout is that a malformed *known* section refuses the whole code (only an
  *unrecognised section name* is forgiven).
- `FS1Build.guild?: { name: string; rankIndex: number }` (TS) / `build.guild` (Lua, `nil` or
  a table) — **always optional, never defaulted to an empty object** when absent, unlike
  `bags`/`bank`/`sets`/`loadouts`/`professions`, which default to `[]`. "No guild" and "an
  empty guild" are not the same fact and must not be conflated.
- `MAX_CODE_LENGTH` (16,384) is untouched; a guild section is under 100 bytes.

### 1.3 Codec changes, file by file

| File | Change |
|---|---|
| `addon/ForeverSixty/Export.lua` | `Export.guildInfo()` (new); `Export.string` passes `guild = Export.guildInfo()` |
| `addon/ForeverSixty/Codec.lua` | `Codec.encodeFS1`: append a `guild=` section when `build.guild` is set, after `professions`. `Codec.decodeFS1`: new branch in the section loop, `name == "guild"` |
| `addon/ForeverSixty/Locale.lua` | `L.codecGuildRank = "That code has an unreadable guild rank: %s."` (or the locale's phrasing convention) |
| `web/src/lib/planner/fs1.ts` | `FS1Build.guild?`; `encodeFS1V2` appends the section after `professions`; `decodeFS1`'s section loop gains a `guild` branch. `DecodedFS1Build` does **not** add `guild` to its `Required<Pick<...>>` union — it stays optional on the decoded result too, since absence is meaningful |
| `tools/gen-codec-vectors.mjs` | Two new FS1 vectors (below); regenerate both fixture copies |

### 1.4 Codec vectors to add

`tools/gen-codec-vectors.mjs`'s `FS1_CODES` (valid) array:

```js
['version 2, guild', 'FS1:1.60.1.69893:warrior:tauren:0/0/0:|guild=Iron%20Vanguard:2'],
```

`INVALID` array (a malformed **known** section refuses the whole code):

```js
['a non-numeric guild rank', 'FS1:1.60.1.69893:warrior:tauren:0/0/0:|guild=Iron%20Vanguard:officer'],
```

Run `node tools/gen-codec-vectors.mjs`; both `addon/tests/fixtures/codec-vectors.json` and
`web/src/fixtures/addon/codec-vectors.json` must come out byte-identical
(`addon/tests/vectors_spec.lua`'s existing "byte-identical in both lanes" test already
guards this — no new test needed there). `addon/tests/codec_fs1_spec.lua` and
`web/src/lib/planner/fs1.test.ts` each get one new `it`/`describe` reading the shared
vectors for the `guild` case, matching how `professions` is tested today.

### 1.5 Spike checklist addition (`addon/README.md`, after row 22)

| # | Check | Run | Constant it fixes |
|---|---|---|---|
| 23 | Guild info shape while in a guild | `/dump GetGuildInfo("player")` — expect `name, rankName, rankIndex[, realm]`; confirm `rankIndex` is `0` for the guild master | `Export.guildInfo()` reads positions 1 and 3; if the shape differs, fix there, not a flag |
| 23a | Guild info while unguilded | `/dump GetGuildInfo("player")` on a character with no guild — expect `nil` | Confirms `Export.guildInfo()` returns `nil` and `Export.string` writes no `guild=` section |
| 23b | Round trip | `/fs export` while in a guild, paste the code into the planner's import box on the site, confirm the guild name and rank index shown there match what `/dump GetGuildInfo("player")` reported | End-to-end check that `Codec.encodeFS1` and `fs1.ts`'s `decodeFS1` agree, beyond the fixture vectors |

Add one row to the **Manual checklist**: "The Export tab's export includes a `|guild=`
section for a guilded character and none for an unguilded one (`/fs diag` or `/dump` the
saved string)."

## 2. The API: membership, verification, claim, invite, settings

### 2.1 Model (coordinator direction, 2026-09-21 review — overrules the first draft's
RULING 1/2, "roster is per account with a sticky character anchor")

The roster is a list of **characters**, because that is what officers actually manage and
what the rating, readiness and loot features (3.3/3.4, out of scope here but not far off)
will key on. Two tables, not one:

- **`guild_characters`** (new) — one row per character currently reporting membership in a
  guild. Primary key `(guild_id, character_key)`, **plus a table-wide `unique
  (character_key)`** — a character belongs to at most one guild at a time, exactly as the
  game itself enforces, and this is what makes "find this character's current guild, if
  any" an index lookup rather than a scan. Each character's export governs **only its own
  row**: no sticky anchor, no batch-order dependency, and an unguilded alt's sync never
  touches its main's row, because they are different rows entirely.
- **`guild_members`** (existing, migration 0005) — stays exactly what `GuildRank`,
  `mayView` and `mayEdit` already read: one row per **account** per guild. It is now a
  **derived** row, recomputed from that account's `guild_characters` rows every time one of
  them changes (§2.2's `recomputeMembership`), never written directly by an export. An
  account is a member of a guild while it has at least one `guild_characters` row there,
  of any verification status; its `rank` is the highest rank among its **verified**
  characters only (an unverified officer-rank character does not make the account read as
  an officer — that would defeat §3.3's raised corroboration bar); it is verified
  (`verified_at is not null`) when at least one of its characters is. Consent
  (§3.2) stays here, per account per guild, untouched by recomputation.

`api/internal/db/migrations/0018_guild_membership.up.sql` (paired `.down.sql` drops
`guild_characters` and the added columns/indexes in reverse):

```sql
-- One row per character currently reporting membership in a guild. rank_index is the raw
-- GetGuildInfo value (kept so an officer-threshold change, §2.3, can re-derive rank with
-- no addon round trip); rank is the derived label; source records how the row was made;
-- verified_at is null until corroborated (§3.3).
create table if not exists guild_characters (
  guild_id     bigint not null references guilds (id) on delete cascade,
  character_key text not null,
  user_id      bigint not null references users (id) on delete cascade,
  rank_index   smallint,
  rank         text not null default 'member' check (rank in ('member', 'officer', 'leader')),
  source       text not null default 'export' check (source in ('export', 'invite', 'claim')),
  verified_at  timestamptz,
  refreshed_at timestamptz not null default now(),
  primary key (guild_id, character_key)
);
-- A character belongs to at most one guild at a time - this is both that invariant and
-- the index PutExports uses to find a character's *previous* guild row (§2.2 step 2b).
create unique index if not exists guild_characters_character_idx on guild_characters (character_key);
create index if not exists guild_characters_user_idx on guild_characters (guild_id, user_id);
create index if not exists guild_characters_unverified_idx on guild_characters (guild_id)
  where verified_at is null;

-- guild_members stays the account-level access row every existing GuildRank/mayView/
-- mayEdit call site reads; it gains per-account-per-guild consent and the derived
-- verified_at recomputeMembership (§2.2) writes. Its `rank` column already exists
-- (migration 0005); this spec stops writing it directly and starts deriving it.
alter table guild_members add column if not exists consent text not null default 'gear'
  check (consent in ('roster', 'gear', 'gear_bags'));
alter table guild_members add column if not exists verified_at timestamptz;
create index if not exists guild_members_user_idx on guild_members (user_id);

-- guilds gains: the officer-rank threshold (a claimed guild's own setting, §2.3), the
-- pending half of the two-step claim (§2.4), and the invite link (§2.5, hashed like
-- login_tokens and devices - never stored in the clear).
alter table guilds add column if not exists officer_max_rank_index integer not null default 1;
alter table guilds add column if not exists claim_pending_by bigint references users (id) on delete set null;
alter table guilds add column if not exists claim_requested_at timestamptz;
alter table guilds add column if not exists invite_token_hash bytea;
alter table guilds add column if not exists invite_token_rotated_at timestamptz;
```

One consequence worth stating: an account can now be a verified member of **two different
guilds at once** through two different characters — genuinely correct (an alt in a second
guild is a real, common case), and something the first draft's single-sticky-character
model could not represent at all.

### 2.2 Creating and refreshing membership from an export

New file `api/internal/addon/guild.go`, unchanged from the first draft — a small, targeted
parser, not a Go port of the whole FS1 codec:

```go
// ParseFS1Guild reads only the `guild=` section of an FS1 v2 export string, ignoring
// everything else. It exists because the API otherwise never parses an export (input.go's
// own comment: "an opaque string this repository never parses") — this is the one field
// that must cross that line, and it crosses it minimally.
func ParseFS1Guild(export string) (name string, rankIndex int, ok bool)
```

Called from `Store.PutExports` (`api/internal/addon/addon.go:87`), once per character in
the batch, **after** that character's `addon_exports` row is written. Per character:

1. Look up this character's current row, if any: `select guild_id, user_id from
   guild_characters where character_key = $1` (the unique index makes this O(1); at most
   one row exists, per the invariant).
2. **The export names a guild.** Resolve/create the guild row (as before: `insert into
   guilds (region, ruleset, name) ... on conflict do nothing`, then `select id,
   officer_max_rank_index`), derive `rank` text from `rankIndex` and
   `officer_max_rank_index` (unchanged formula: `0` → `leader`; `1..threshold` → `officer`;
   else `member`), then:

   ```sql
   insert into guild_characters (guild_id, character_key, user_id, rank_index, rank, source, refreshed_at)
   values ($1, $2, $3, $4, $5, 'export', now())
   on conflict (guild_id, character_key) do update set
     user_id = excluded.user_id, rank_index = excluded.rank_index, rank = excluded.rank,
     refreshed_at = now();
   -- verified_at is deliberately not in this SET list: a re-sync of the same character in
   -- the same guild keeps whatever verification it already earned. Only a transfer (below)
   -- starts a fresh, unverified row - only §3.3's corroboration paths ever set verified_at.
   ```

   If step 1 found a **different** `guild_id` for this character (a transfer): `delete from
   guild_characters where character_key = $1 and guild_id = $previous_guild_id` — the old
   row is not left behind as a ghost.
3. **The export names no guild** (or step 1 found a row and this export is unguilded):
   `delete from guild_characters where character_key = $1`.
4. **`recomputeMembership(guildID, userID)`** — run for the new `guild_id` (step 2) and,
   whenever a transfer or removal happened, for the *previous* `guild_id` too, so both
   accounts' derived rows stay correct in the same transaction as the character write:

   ```sql
   -- Recompute the account-level guild_members row for (guild_id, user_id) from its
   -- current guild_characters rows. $2 may be a specific user_id, or NULL to recompute
   -- every account in the guild at once (used by §2.3's officer-threshold change).
   insert into guild_members (guild_id, user_id, rank, verified_at, refreshed_at)
   select $1, gc.user_id,
          case max(case when gc.verified_at is not null then
                      case gc.rank when 'leader' then 3 when 'officer' then 2 else 1 end
                    end)
            when 3 then 'leader' when 2 then 'officer' else 'member' end,
          min(gc.verified_at) filter (where gc.verified_at is not null),
          now()
   from guild_characters gc
   where gc.guild_id = $1 and ($2::bigint is null or gc.user_id = $2)
   group by gc.user_id
   on conflict (guild_id, user_id) do update set
     rank = excluded.rank, verified_at = excluded.verified_at, refreshed_at = now();
   -- consent is not in this SET list either - recomputation never touches it.

   -- An account with no guild_characters row left in this guild is not "a member with
   -- rank member" - it is not a member at all.
   delete from guild_members m
   where m.guild_id = $1 and ($2::bigint is null or m.user_id = $2)
     and not exists (
       select 1 from guild_characters gc where gc.guild_id = m.guild_id and gc.user_id = m.user_id
     );
   ```

   `verified_at` is the **earliest** verification among the account's verified characters
   (`min`, not `max`) — it answers "how long has this account been trustworthy for this
   guild," which is the more stable and more meaningful reading; it does not move around as
   individual characters re-verify.
5. **Claim release**, checked *after* recomputation: if the guild's `claimed_by = user_id`
   and the post-recompute `guild_members` row for `(guild_id, user_id)` is now absent or
   `verified_at is null`, clear `claimed_by` on `guilds` in the same transaction — leaving
   (or losing verification in) a guild you lead releases your claim, even if you still have
   another, unrelated character elsewhere. This closes the same gap 3.6's last sentence
   closes for report access.

**Ageing** now applies to `guild_characters` rows: one whose `refreshed_at` is older than
**45 days** (RULING 3, unchanged reasoning — kept per the coordinator's instruction) is
removed by a scheduled sweep (same shape as `sim_specs`'s nightly validation job), followed
by `recomputeMembership` for the affected `(guild_id, user_id)` pairs — "the account row
follows": it downgrades or disappears exactly as if that character's export had gone
unguilded.

### 2.3 Officer detection

Unchanged in substance from the first draft, now scoped to `guild_characters`.
`rankIndex == 0` is always the guild master (`'leader'`) — server-authoritative, not
configurable. **`guilds.officer_max_rank_index` defaults to `1`** (RULING 4, kept). A
claimed guild's settings page (§2.6) can widen or narrow it; doing so re-derives every
character's `rank` from its stored `rank_index` in one statement, then recomputes every
affected account:

```sql
update guild_characters set rank = case
    when rank_index = 0 then 'leader'
    when rank_index <= $2 then 'officer'
    else 'member'
  end
where guild_id = $1 and rank_index is not null;
-- followed by recomputeMembership($1, null) - every account in the guild at once.
```

### 2.4 The claim flow

The eligibility check is now **raw**, read directly from `guild_characters` (bypassing the
verified-only `guild_members`/`GuildRank`, exactly as the first draft intended, but now
correctly scoped to a character rather than an account-wide raw rank that no longer
exists): `select rank from guild_characters where guild_id = $1 and user_id = $2 and rank
in ('officer', 'leader') order by (rank = 'leader') desc limit 1`. Claiming is itself the
corroboration mechanism, so it does not require `verified_at` beforehand.

- **A character at `rankIndex == 0` (the guild master) claims immediately**, no
  confirmation needed. `guilds.claimed_by = user_id`; **every one of that account's
  `guild_characters` rows for this guild** gets `verified_at = coalesce(verified_at,
  now())` (usually just the one GM character, but any alts the account also has in the
  guild are trusted along with it), then `recomputeMembership`.
- **An officer (not GM) claims and it goes pending**: `claim_pending_by = user_id`,
  `claim_requested_at = now()`. Confirmed by either:
  - a **second, distinct** account with a `guild_characters` row in this guild at rank
    `officer` or `leader` calling confirm, or
  - the guild master's own account reaching rank `leader` for this guild at any point while
    the claim is pending (checked at the moment any export names this guild with
    `rankIndex == 0` — auto-confirms with no separate API call, matching the proposal's
    "failing that, by the guild master's export").
  A pending claim **expires after 14 days** unverified (RULING 5, kept). On confirmation,
  `claimed_by = claim_pending_by`, the pending fields clear, and every one of the
  **original claimant's** `guild_characters` rows for this guild gets `verified_at =
  coalesce(verified_at, now())`, then `recomputeMembership` for that account (the
  confirmer's own rows are unaffected — they only vouch).
- **Dispute/release**: the current `claimed_by` account, or a moderator, may release the
  claim (`claimed_by = null`); also released automatically per §2.2 step 5.

#### Amendment, 2026-09-21 (security review response)

An independent security review of the `guild-api` branch found the guild-master-immediate
branch above exploitable as written: nothing distinguished a genuine `GetGuildInfo`
`rankIndex == 0` from a hand-typed one in a forged `POST /v1/addon/exports` body, so a
single self-posted string let an attacker instantly become a verified leader of any
**unclaimed** guild — which, at ship time, is every real guild, since `claimed_by` starts
empty for all of them. This defeated §3.3's "raised corroboration bar" entirely for the one
path that bootstraps a guild's whole claim/officer structure. The review also found
`AutoConfirmClaimIfPending` (below) never checked the confirming signal was a distinct
account from the pending claimant, letting a single account self-confirm its own pending
claim with a second forged character.

The coordinator's ruling: **the guild-master-immediate branch stays the bootstrap** — an
unclaimed guild has no verified member who could vouch for a second signal, so requiring
one here would make an unclaimed real guild unclaimable by its real leader. Instead the
branch is hardened, and — new — made recoverable if it is still abused:

- **Battle.net identity required.** The claiming account must have a linked Battle.net
  identity (`users.bnet_sub` is not empty) before the guild-master branch will act on it.
  Forever supports both Battle.net and email magic-link sign-in (`api/internal/auth/handler.go`
  mounts both unconditionally), so this is a real, enforced check — not a fact already true
  of every account — and it raises the floor from "any signed-in account with one paired
  device" to "an account that has actually authenticated through Battle.net," which is a
  real person's game account, not a disposable email address.
- **Rate-limited to one claim (successful or pending) per account per rolling 30 days**,
  tracked in a new `guild_claim_attempts` table, and **at most one currently-claimed guild
  per account at a time** (checked against `guilds.claimed_by` directly). A successful
  attack no longer scales past one guild per attacker per month.
- **Claims are contestable.** `POST /v1/guilds/{id}/claim/contest` (§2.6): a caller with a
  raw `guild_characters` row in that guild at `rankIndex == 0` or officer rank, on a
  *different* account from the current claimant/pending claimant, flags the claim as
  disputed (`guilds.claim_contested_at`/`claim_contested_by`) and **freezes** the
  claimant's officer powers for that guild — approve, remove, settings, invite rotate, and
  (via the unchanged `auth.Store.GuildRank`/`mayEdit` path) attaching a report — until a
  moderator resolves it with `POST /v1/guilds/{id}/claim/resolve`, outcome `uphold`
  (dismiss the contest), `release` (clear the claim and un-verify every character whose
  *only* verification source was the claim), or `transfer` (move the claim, and the same
  claim-sourced verification, to the contesting account). A genuine guild master who is
  attacked no longer has no way back in — they contest, and a moderator settles it.
- **`verified_at` now records its source** (`guild_characters.verified_by`, one of
  `claim`/`officer`/`invite`/`logs`), so a claim `release` can un-verify precisely the rows
  the claim itself vouched for and nothing else — a character an officer separately
  approved, or that corroborated by logs, keeps its verification even if the claim that
  first brought its account into the guild is later released.
- `GET /v1/guilds/{id}/settings` and `GET /v1/guilds/{id}/home` now expose `claim: {
  state: "unclaimed"|"pending"|"claimed"|"contested", since }` so the web can show it (the
  review separately found `SettingsView.ClaimedBy`/`ClaimPending` declared but never
  populated in the shipped code — fixed alongside this change, not a spec gap).
- `AutoConfirmClaimIfPending` now requires the triggering `guild_characters` row's account
  to differ from `claim_pending_by`, mirroring `ConfirmClaim`'s existing `ErrSameAccount`
  check for the explicit confirm endpoint.

**Residual risk, restated:** a Battle.net-linked account can still, once every 30 days,
instantly claim one currently-unclaimed real guild by forging a single `rankIndex == 0`
export naming it. This is deliberately not closed outright (closing it would make a
genuinely unclaimed guild unclaimable by its real leader without a second signal that does
not exist yet); it is made rare (30-day, one-guild-at-a-time), attributable (a real
Battle.net identity, not a disposable one), and reversible (contest + moderator resolve)
instead.

#### Second amendment, 2026-09-21 (second security review response)

A scoped re-review of the fix round above marked six of seven original findings fixed with
real tests, but found the contest mechanism *itself* exploitable exactly the way the claim
flow it hardens against once was: `POST /v1/guilds/{id}/claim/contest` had no Battle.net
requirement, no rate limit (despite the endpoint table already saying "rate-limited like
claim"), and an upheld contest left the losing contester free to re-contest at once — a
free, email-only account with one forged officer-rank export could freeze any legitimately
claimed guild's officer tools, indefinitely, and many guilds at once. Fixed, mirroring the
claim flow's own hardening:

- **Contesting needs a linked Battle.net identity**, exactly as a leader claim does
  (`ErrNoBattleNetIdentity`, checked first).
- **Contest attempts are rate-limited per account in the database**: one contest per
  account per rolling 30 days across every guild, and at most one *open* contest per
  account at a time (`guild_claim_attempts` gained a `kind` column shared with claim
  attempts, so both share one 30-day-window query shape). The route is also wrapped in the
  same per-IP `httpx.RateLimitPer` the invite-accept route uses, at a much lower ceiling
  (5/hour/IP) given how much one successful contest can suspend.
- **An uphold is final for that `(guild, contesting account)` pair.** Every resolution is
  recorded in a new `guild_claim_resolutions` table (`guild_id`, `contester_id`, `outcome`,
  `resolved_at`, `moderator_id`); a repeat contest of the same guild by the same account
  after an uphold is refused (409); and an uphold **deletes the contester's own unverified**
  `guild_characters` rows for that guild — a *verified* row of theirs survives, since that
  account is then understood to be a real member who lost a dispute, not an impostor who
  needs removing.
- **A contest freezes officer tools only for a young or uncorroborated claim.** A new
  `guilds.claimed_at` column (distinct from `claim_requested_at`, which only ever times a
  *pending* claim) records when the currently active claim was established — by the
  guild-master-immediate branch, a confirm, an auto-confirm, or a contest `transfer`. A
  contest freezes officer tools when that claim is **less than 14 days old**, or when the
  guild has **no character verified by `logs`** other than the claimant's own; otherwise
  (an established claim with independently log-verified members) the contest is still
  recorded and queued for a moderator, but nothing freezes. `claim: {state, since, frozen}`
  (both `GET .../settings` and `GET .../home`) carries the new `frozen` boolean so the web
  shows the right thing; every freeze check in the codebase now reads "contested AND
  frozen," extracted into one shared `Service.freezeCheck` helper (it takes the
  moderator-exemption as a parameter, since `PATCH .../settings`'s moderator bypass must
  also bypass the freeze, while a route with no moderator standing blocks everyone alike
  once frozen).
- **The freeze now also covers report edit rights.** While a guild's claim is contested
  *and* frozen, the disputed claimant's officer-derived edit right over that guild's
  reports (`reports.mayEdit`) is suspended too, through a small `GuildClaims.FrozenClaimant`
  hook `reports.Service` now holds — their own reports, every other verified officer, and
  every moderator are unaffected. Previously a frozen claimant could still retitle,
  re-scope, or attach/detach the guild's reports through the one route this spec had left
  untouched.
- **Guild name validation now rejects Unicode category Cf** (format characters — zero-width
  space, zero-width joiner, right-to-left override, the byte-order mark) alongside the
  pre-existing Cc (control) check, since none of them can appear in a real WoW guild name
  and a bidi override in particular can make a guild's displayed name misleading about what
  it actually contains.
- **A character's guild transfer now locks both guilds it touches, in a fixed ascending
  order, before any mutation.** Two concurrent transfers moving characters in opposite
  directions between the same two guilds previously could each lock their own "new" guild
  first and then deadlock waiting for the other's "old" guild; `guilds.LockGuilds` rules
  this out and is covered by a concurrent-opposite-transfers regression test.

Guild identity stays `(region, ruleset, lower(name))` rather than adding realm to the key:
Forever merges each ruleset's original realms into one shared roster and leaderboard, so a
guild name is only ever ambiguous within a `(region, ruleset)` pair, never within a single
original realm, and the public guild page (`rankings/guilds.go`) already keys the same way.

### 2.5 The invite link

Unchanged mechanics from the first draft (a random 32-byte token, shown once, stored only
as its SHA-256 hash on `guilds.invite_token_hash`; only a verified officer/leader may
rotate it; no fixed TTL — RULING 6, kept; redemption rate-limited at 20/hour/IP).

Redeeming needs one adjustment for the character-keyed model: an invite-joined account has
no character at all yet, so it cannot anchor a normal `guild_characters` row by
`character_key`. It gets a **synthetic key**, `'account:' || user_id` — a namespace that
never collides with a real `character.Key` output (which is always a real
`region/ruleset/name-slug`, never prefixed `account:`). The row is created with `source =
'invite'`, `verified_at = now()` immediately (an officer sharing a secret link **is** the
corroboration), `rank = 'member'`, `rank_index = null`, then `recomputeMembership`. If that
account later syncs a real, guilded character via the addon, it gets its **own** ordinary
row alongside the synthetic one — the roster (§4.1) hides synthetic `account:`-prefixed
rows from its character list, since they represent no real character to show class/spec/
item level for; they still count correctly toward the account's derived `guild_members`
row. An invite-sourced row is never promoted to officer/leader; only an export carrying a
real `rankIndex` can do that.

### 2.6 Endpoints

New package `api/internal/guilds` (mutations; `rankings` stays the public read-only guild
page, `auth` keeps owning `/v1/me`). Every response uses the existing envelope
(`httpx.WriteOK`/`WriteError`); every write is `RequireSession`. Parameterised SQL only,
per the existing store style throughout this codebase.

| Method & path | Auth | Body | 200 body | Errors |
|---|---|---|---|---|
| `POST /v1/guilds/{id}/claim` | session | — | `{status: "confirmed"\|"pending", expires_at?}` | 403 `forbidden` (no character at rank officer/leader in this guild); 404 `not_found`; 409 `conflict` (already claimed, or already pending) |
| `POST /v1/guilds/{id}/claim/confirm` | session | — | `{status: "confirmed", claimed_by: {battletag}}` | 403 `forbidden` (no officer/leader character, or same account as the pending claimant); 404 `not_found` (no pending claim, or expired) |
| `POST /v1/guilds/{id}/claim/release` | session | — | `{status: "released"}` | 403 `forbidden` (not `claimed_by`, not moderator) |
| `POST /v1/guilds/{id}/claim/contest` (2026-09-21 amendment; hardened by the second amendment) | session, Battle.net identity required, rate-limited per-account (30 days, one open contest) and per-IP (5/hour) | — | `{status: "contested"}` | 403 `forbidden` (no officer/leader-or-rank-0 character, same account as the claimant, or no Battle.net identity); 404 `not_found`; 409 `conflict` (no active claim, already contested, already upheld against this account, or an open contest already exists elsewhere for this account); 429 `rate_limited` |
| `POST /v1/guilds/{id}/claim/resolve` (2026-09-21 amendment; records to `guild_claim_resolutions` per the second amendment) | session, moderator only | `{outcome: "uphold"\|"release"\|"transfer"}` | `{status: "resolved", outcome}` | 400 `invalid`; 403 `forbidden` (not a moderator); 404 `not_found`; 409 `conflict` (no contested claim) |
| `GET /v1/guilds/{id}/settings` | session, verified officer/leader or moderator | — | `{default_visibility, officer_max_rank_index, claimed_by, claim_pending, claim: {state, since?, frozen} (`frozen` added by the second amendment), invite: {rotated_at}}` | 403 `forbidden` |
| `PATCH /v1/guilds/{id}/settings` | session, verified officer/leader or moderator | `{default_visibility?, officer_max_rank_index?}` | updated settings | 400 `invalid` (`default_visibility` must be `public`, `unlisted` or `guild` — never `private` for a guild default); 403 `forbidden`; 409 `claim_contested` (only when the claim is contested AND frozen — second amendment; a moderator bypasses this too) |
| `POST /v1/guilds/{id}/invite/rotate` | session, verified officer/leader | — | `{token, url, rotated_at}` (token shown once) | 403 `forbidden`; rate-limited; 409 `claim_contested` (contested AND frozen — second amendment) |
| `POST /v1/guilds/invite/{token}/accept` | session | — | `{guild: {...}, rank: "member"}` | 404 `not_found` (unknown or revoked token — never distinguished); rate-limited 20/hour/IP |
| `POST /v1/guilds/{id}/characters/{character_key}/approve` | session, verified officer/leader | — | updated character row (`verified_at`/`verified_by='officer'` set) | 403 `forbidden`; 404 `not_found` (no such character row); 409 `claim_contested` (contested AND frozen — second amendment) |
| `DELETE /v1/guilds/{id}/characters/{character_key}` | session, the character's own account **or** a verified officer/leader — **rank protects rank (2026-09-21 amendment)**: an `officer`-rank row also requires the account currently holding the claim; a `leader`-rank row only its own account or a moderator | — | `{status: "removed"}` | 403 `forbidden`; 404 `not_found`; 409 `claim_contested` (contested AND frozen — second amendment; never on the caller's own row or a moderator's call) |
| `PATCH /v1/guilds/{id}/members/me` | session | `{consent: "roster"\|"gear"\|"gear_bags"}` | updated member row | 400 `invalid`; 404 `not_found` (no membership) |
| `DELETE /v1/guilds/{id}/members/me` | session | — | `{status: "left"}` (removes every one of the caller's own `guild_characters` rows in this guild) | 404 `not_found` |
| `GET /v1/guilds/{id}/home` | session, any member (verified or not — §3.2) | — | this week's **verified-or-public-or-own** reports (2026-09-21 amendment), character roster, who-logged (§4.1), `claim: {state, since?, frozen}` (`frozen` added by the second amendment) | 403 `forbidden` (not a member); 404 `not_found` |

Report editing (`api/internal/reports/handler.go`'s `mayEdit`) also observes the contested-
AND-frozen freeze now (second amendment): while frozen, the disputed claimant's
officer-derived edit right over their guild's reports is suspended through a new
`reports.Service.Guilds` (`GuildClaims.FrozenClaimant`) hook — their own reports, every
other verified officer, and every moderator are unaffected.

Both new mutating endpoints (`approve`, the officer branch of `DELETE .../characters/...`)
call `recomputeMembership` for the affected account afterward, same as every other write in
this section.

`GET /v1/me` gains, per `Guild` entry (`api/internal/auth/store.go:105`): `Consent string`
and `Verified bool` (`verified_at is not null`). `auth.Store.Guilds`'s `ORDER BY g.name`
changes to `ORDER BY m.refreshed_at DESC` — most-recently-active membership first, which is
what "My guild" in the header (§4.3) reads position `[0]` from. (`store_test.go`'s existing
`TestGuildRank`-adjacent assertion only checks the single-guild case — this ordering change
does not touch it; confirmed by reading `api/internal/auth/store_test.go:227-230`.) `Me`
stays account-level, one row per guild the account has any character in; the per-character
breakdown lives only in `GET /v1/guilds/{id}/home`.

## 3. Access control

### 3.1 Who can see a `guild`-visible report, precisely

Today (`mayView`, `api/internal/reports/handler.go:521`): the report's owner, a moderator,
or **any** row in `guild_members` for that guild (rank irrelevant). This spec **narrows**
that to: the owner, a moderator, or a `guild_members` row with `verified_at is not null`.

This narrowing needs **no interface or call-site change at all** in the reports package.
`mayView` and `mayEdit` (`handler.go:521`, `:543`) keep calling
`s.Accounts.GuildRank(ctx, guildID, userID)` exactly as today, with its existing signature
(`string, bool, error`) — because, per §2.1, `guild_members` is now a **derived** row whose
`rank` only ever reflects the account's **verified** characters. The one change is inside
`auth.Store.GuildRank` itself (`api/internal/auth/store.go:413-424`), which gains a single
clause so an unverified-but-present row does not answer `ok = true`:

```sql
select rank from guild_members where guild_id = $1 and user_id = $2 and verified_at is not null
```

`mayEdit`'s rank check and the `PATCH .guild_id` "must be an officer of that guild" check
(`handler.go:317`) inherit the same tightening automatically, since they read the same
method. The raw, unverified signal (needed only for §2.4's claim precondition, which is
deliberately allowed to run before verification) never goes through `GuildRank` at all —
the claim flow queries `guild_characters` directly, as specified there.

A second, new method serves a different purpose — letting an **unverified** member into the
free home shell (§4.1), which `GuildRank` must not do:

```go
// IsMember reports whether this account has any guild_characters row in this guild,
// verified or not - the gate for GET /v1/guilds/{id}/home, never for report visibility
// or edit rights (those stay on GuildRank, which is verified-only).
IsMember(ctx context.Context, guildID, userID int64) (bool, error)
```

owned by the new `api/internal/guilds` package's own store, not exposed to `reports`.

### 3.2 Consent (proposal 3.6)

Stored on `guild_members.consent` (§2.1), one of `roster | gear | gear_bags`, default
`gear` (matches the proposal's stated default exactly). It is **per membership**, not a
global account setting — a player's trust in one guild's officers is not necessarily their
trust in another's, and the column lives on the row that already scopes "this account, this
guild."

**Every guild-home query honours it** by never reading gear or bag contents from
`addon_exports`/`fight_metrics` for a roster row unless that row's `consent` allows it:

- `roster` — name, class, spec only (from the latest ranked fight or export; item level and
  gear are withheld even though the data exists in `addon_exports`).
- `gear` (default) — adds item level and equipped gear (what the roster table already shows
  publicly today via `RosterBest`, extended to the signed-in home's fuller roster).
- `gear_bags` — adds bag contents, needed for the (out-of-scope, 3.3) readiness board later;
  built here only so the setting exists and the column is honoured, not consumed by
  anything yet.

Officers see the *choice* a member made (e.g. "Simfury: gear and bags"), never data the
choice withholds — enforced the same way as everything else in §2.6's `/home` endpoint: the
query itself does not select withheld columns for a row below the caller's own consent
level, rather than fetching everything and filtering in the handler (fail-closed by
construction, not by remembering to check).

### 3.3 Security

**Spoofing (the central risk).** A player can hand-edit their own `SavedVariables` file (or
bypass the game entirely and call `POST /v1/addon/exports` directly with a forged export
string) to claim membership in any guild, since `auth.RequireDevice` proves *"this is a
paired device on my own account,"* not *"this string came from `GetGuildInfo` on a real
character in that guild."* §3.1's narrowing is the answer: a bare `guild_characters` row
from an export **never** grants access to a `guild`-visible report, edit rights, or the
"officer" consequences of §2.3–2.4 by itself. It grants only `verified_at is null`
membership — visible on the free home as an unverified roster entry, nothing more (§4.1
marks these distinctly, with a one-click approve for a verified officer viewer).
`verified_at` on a `guild_characters` row is set by any **one** of four paths (raised from
the first draft's single-appearance bar, per coordinator review — a lone appearance in a
guild's log is not enough on its own):

1. **Corroboration by the guild's own uploaded logs, on two distinct raid nights.** The
   character appears in `fights.players` for fights belonging to reports with this
   `guild_id`, on **at least two distinct report dates within any trailing 30-day window**
   (not two fights in the same night, and not two nights more than 30 days apart). Checked
   whenever a report's `guild_id` is set (`handler.go:307`) or a new fight is ingested for
   an already-guild-attached report:

   ```sql
   update guild_characters gc
   set verified_at = coalesce(gc.verified_at, now())
   where gc.guild_id = $1
     and gc.character_key = any($2)  -- character keys this report/fight touched
     and (
       select count(distinct r2.created_at::date)
       from fights f2 join reports r2 on r2.id = f2.report_id
       where r2.guild_id = $1
         and gc.character_key = any(f2.players)
         and r2.created_at >= now() - interval '30 days'
     ) >= 2;
   ```

   A single forged export cannot satisfy this at all — it needs two genuinely separate
   nights of that guild's own real combat-log uploads naming the same character, which is a
   materially higher, sustained bar than a one-off appearance.
2. **An officer of the claimed guild approves it** from the roster's pending list
   (§2.6's new `POST .../characters/{character_key}/approve`) — a human, social-trust
   decision by someone who already has standing in the guild, available immediately with no
   wait for a second raid night.
3. **The claim flow** (§2.4), which has its own two-signal corroboration (a second
   independent officer-rank character, or the GM).
4. **The invite link** (§2.5) — an officer's own out-of-band action is the corroboration.

*Residual risk, stated plainly, under the raised bar:* a player who (a) has a paired,
signed-in companion device on their own account, (b) is technically capable of a direct API
call, and (c) has genuinely appeared as a player in the target guild's own uploaded reports
on **two separate raid nights within a 30-day window** (e.g. filled a pugged slot twice for
the same guild in a month) can forge a `guild=` section naming that guild afterward and
become `verified_at`-corroborated from those two real appearances, gaining read access to
that guild's private reports going forward. This is a substantially higher bar than the
first draft's single-appearance version — it requires sustained, repeated, genuine overlap
with the guild's actual raiding, not a coincidence — but it is not zero. **Mitigation, not
closure**: officers see every character on their roster, verified or not, with a one-click
remove next to each (§2.6's `DELETE .../characters/{character_key}`, officer branch);
removing a row does not auto-re-verify on the character's next export unless it earns
verification again through one of the four paths above, so a wrongly-verified character an
officer notices and removes stays removed. Fully closing this residual risk would require
signing the export string itself in the game client (out of reach — no network access
in-game) or trusting Blizzard's API for guild roster (forbidden by the terms proposal §1
already rules out). **RULING 7 (revised)**: accept this narrower residual risk rather than
block the whole feature on an unreachable cryptographic guarantee; *cost if wrong*: a
guild's private logs leak to someone who genuinely raided with them twice in a month and
later fabricates a permanent membership claim — mitigated by the officer remove action and
by the fact that two-night overlap already means the guild's own officers likely recognise
the name.

**Enumeration.** `GET /v1/guilds/{region}/{ruleset}/{name}` (public, existing) already lets
anyone probe guild names one at a time; this spec adds no new enumeration surface beyond it.
`POST /v1/guilds/invite/{token}/accept` answers the same 404 for "no such token" and "token
was rotated away" — never distinguishing them — so a guessed-and-failed token teaches
nothing.

**Invite-link leakage.** A leaked invite link lets anyone who is signed in join as a
`verified` member (§2.5 — trusted by construction, since an officer chose to share it). This
is accepted as the feature's own design (a link is meant to admit people); the mitigation is
rotation, which is one click and invalidates the leaked link at once. The settings page
(§5.4) states this plainly: "Anyone with this link can join as a member. Rotate it if it
leaks."

**Leaving.** Deleting one's own characters' rows (`DELETE
/v1/guilds/{id}/members/me`, or one character at a time via `DELETE
.../characters/{character_key}`) triggers `recomputeMembership`, which removes the account's
`guild_members` row the moment no `guild_characters` row remains — `mayView` re-reads
`GuildRank` on every request, so there is nothing cached to invalidate. The same happens
automatically the moment a character's export stops naming the guild (§2.2 step 3), and
clears `claimed_by` too when the leaving account held it (§2.2 step 5). This matches 3.6's
closing sentence exactly: "Leaving the guild, or an export that names a different guild,
removes access at once" — now true per character, which is the more precise reading of
"an export" than the first draft's single account-wide anchor gave it.

#### Amendment, 2026-09-21 (security review response)

Beyond the claim-flow hardening recorded in §2.4's amendment, the same review found four
more gaps, all fixed as part of the same response:

- **Rank protects rank.** `DELETE .../characters/{character_key}` (§2.6) let *any* verified
  officer remove *any* character on the roster, including the guild master's own —
  combined with §2.2 step 5's automatic claim release, an officer could strip the real
  guild master of membership and immediately claim the now-unclaimed guild themselves.
  Fixed: a `member`-rank row is removable by its own account, a verified officer/leader, or
  a moderator (unchanged); an `officer`-rank row adds only the account currently holding
  the guild's claim (not "any officer") to that list; a `leader`-rank row is removable only
  by its own account or a moderator — never by another officer, and never by the account
  holding the claim either, since a second `leader`-rank row belongs to a different real
  character than the claimant's own.
- **The guild home's report list now respects visibility.** `HomeReports` (§4.1) is
  reachable by any member with even a single, unverified, freshly-forged `guild_characters`
  row (`IsMember`, deliberately not `GuildRank` — that carve-out for the free home shell is
  unchanged), and previously returned every report with the guild's `guild_id` regardless
  of `visibility`, leaking a `private` or `unlisted` report's title, creation time, and
  fight/kill counts to anyone who could forge membership at all. Fixed: the list now shows
  a report the caller owns (any visibility), every `public` report, and a `guild`-visible
  report only once the caller is **verified** (`GuildRank`-equivalent, checked once per
  request and passed down) — never a `private` or `unlisted` row that is not the caller's
  own, even to a verified member; an unverified member's home shows public and their own
  reports only.
- **Guild identity is case-insensitive.** `resolveGuild` matched guild names on exact text
  while the pre-existing public guild page (`rankings/guilds.go`, untouched by this
  amendment) already matched case-insensitively, so two exports differing only in casing
  (`Iron Vanguard` vs. `IRON VANGUARD`) could mint two distinct `guilds` rows for what the
  public page treats as one guild — a cheap way to spawn a same-named decoy guild to claim.
  Fixed: one guild per `(region, ruleset, lower(name))` (a new unique index; the
  pre-existing exact-text constraint from migration 0005 stays and is subsumed by it),
  `resolveGuild` looks up and inserts through the case-insensitive index, and the first
  writer's casing is kept as the guild's display name. Guild names are also now normalised
  to Unicode NFC and trimmed, and rejected (making that one character's guild sync a silent,
  logged no-op — never a 500, never aborting the rest of the export batch) if they exceed
  the game's own 24-character guild name limit or contain a control character.
- **Rank index and guild name are validated at the boundary.** A hand-crafted
  `rankIndex` outside `0-9` (WoW's real range) previously reached
  `guild_characters.rank_index smallint` unbounded, and an invalid-UTF-8 decoded name
  previously reached `guilds.name text` unvalidated; either overflow/encoding failure
  surfaced as an unhandled database error that aborted the *rest* of that `PutExports`
  batch and answered a generic 500. Fixed: both are validated before any database write,
  and a single character's malformed `guild=` section is now a synced-as-unguilded no-op
  for that character alone, with a logged reason, never a batch-wide failure.
- **`RecomputeMembership` now takes a transaction-scoped advisory lock** keyed on the
  guild, serialising every concurrent caller (`PutExports`, `ApproveCharacter`,
  `RemoveCharacter`, `Claim`, `UpdateSettings`, and the ageing/corroboration sweep jobs)
  against each other for that guild, closing the theoretical stale-snapshot race the
  original design left untested (§5's own "concurrent PutExports-plus-approve" case, now
  covered by a real concurrency test).

#### Second amendment, 2026-09-21 (second security review response)

Two further gaps, both fixed as part of the contest-hardening response recorded in §2.4's
second amendment:

- **A character's guild transfer now locks both guilds it touches** — the new one and the
  one being left — in a fixed ascending order before either guild's `guild_characters` row
  is mutated (`guilds.LockGuilds`, called from `addon.Store.syncGuild`). Previously each
  transfer locked only its "new" guild first (via `RecomputeMembership`'s own per-guild
  advisory lock) and then its "old" one; two concurrent transfers moving characters in
  opposite directions between the same two guilds could lock in opposite orders and
  deadlock each other. Covered by a concurrent-opposite-transfers regression test.
- **Guild name validation now rejects Unicode category Cf** (format characters), not only
  Cc (control): zero-width space (U+200B), zero-width joiner (U+200D), the right-to-left
  override (U+202E), and the byte-order mark (U+FEFF) can none of them appear in a real WoW
  guild name, and a bidi override in particular can make a guild's displayed name
  misleading about what it actually contains.

## 4. The web side

### 4.1 The signed-in guild home

`/guild/<region>/<ruleset>/<name>` grows a signed-in section, fetched from the new
`GET /v1/guilds/{id}/home` (§2.6), rendered only when `GET /v1/me`'s `guilds` includes this
guild (any row, verified or not — an unverified member still sees the home, just as they
still appear on the roster; §3.1's narrowing governs *report* access, not the home shell
itself, since the home's own reports list already filters through `mayView` server-side).
`GET /v1/guilds/{id}/home` itself is gated by the new `IsMember` check (§3.1), not
`GuildRank`.

- **This week's reports** — reports with this `guild_id` created in the trailing 7 days,
  newest first, kill/wipe counts, keyset-paginated the same way `GET /v1/reports/recent`
  already is (`api/internal/reports/recent.go`'s cursor shape, reused verbatim for this
  endpoint's `reports` array). **RULING 8**: "this week" is a trailing 7-day window, not a
  server-specific weekly-reset timestamp — no reset concept exists anywhere in this
  codebase to anchor to (`grep` across `api/` and `docs/` turns up none), and inventing one
  here would be a guess dressed as precision. *Cost if wrong:* a report from 6 days ago that
  crossed an actual in-game reset reads as "this week" when a raid leader would call it
  "last week." Cheap to fix later if a real reset schedule surfaces — one `WHERE` clause.
  The game's own weekly raid reset (Tuesday in US regions, region-shifted elsewhere) is the
  obvious real anchor once the site has a region-aware reset schedule to read from
  anywhere; that is a follow-up, not built here.
- **Who logged** — per roster row, whether `addon_exports.updated_at` for that character's
  `character_key` falls within the last 24 hours (**RULING 9**, same reasoning as above: no
  existing "raid night" concept to anchor to; 24 hours is a simple, legible proxy for "ran
  the companion recently"). Same follow-up as RULING 8: a real, region-aware reset schedule
  would make "since the last reset" the more natural window here too.
- **Roster** — one row per `guild_characters` character (verified and unverified, visually
  distinguished; a verified officer viewer sees a one-click Approve next to each unverified
  row calling §2.6's `approve` endpoint, and Remove next to any row). Synthetic
  `account:`-prefixed rows from invite joins with no character yet (§2.5) are not listed
  here — they contribute to the account-level home access but have no class/spec/item
  level to show. Each real row shows class, spec (from that character's latest ranked
  fight or export), `addon_exports.updated_at`, item level, gated by the account's
  `consent` (§3.2, read from `guild_members` for that character's `user_id`). Each row with
  gear-level consent or above gets "Open in simulator" / "Open in planner" links, built
  through **one helper** so the roster's own code never branches on which hand-off module
  is available: a small wrapper (owned by the Web lane, §6) that calls
  `current-character.ts`'s `plannerHrefFor`/`simHrefFor` when that module is present on
  `main` (it is landing from another lane around the same time as this spec) and falls back
  to `handoff-links.ts`'s `plannerCodeHref`/`simCodeHref` (on `main` today) otherwise — the
  same compatibility rule `CharacterHandoffLinks.svelte` already follows for `/account` and
  `/character/<key>`.
- **Progression and rankings** — reuses the existing public `Guild.svelte` data
  (`GuildPage` from `rankings/api.ts`) unchanged; the signed-in sections are additive panels
  above/below it, not a replacement.
- **Empty states**: no reports this week → "No reports this week yet." (honest, matches the
  house style in `2026-09-21-one-product-design.md` §4's empty-state pattern). No roster
  beyond the viewer → "You're the only member the site knows about. Share the invite link to
  bring the rest of the guild in." with the link to `/settings` if the viewer is a verified
  officer, otherwise nothing (they cannot rotate or view it).
- **Signed-out / non-member state**: the page shows exactly what it shows today — public
  progression, roster-bests, reports list — with no hint that a signed-in view exists for
  members (no "sign in to see more" nag inside the tools, per the design system's honest-copy
  rule).

### 4.2 What stays public for non-members

Unchanged from today: progression per boss, pull counts, kill dates, roster-bests built from
`fight_metrics` for **every** report with this `guild_id` and `visibility <> 'private'` —
this already includes `guild`-visible reports' contribution to kill counts and best-parse
rows (§0's note). This spec does not add or remove anything from that surface; it is called
out here so it is not mistaken for new exposure introduced by making membership real.

### 4.3 "My guild" in the header and homepage

`SessionNav.svelte` (currently a single "Sign in" / battletag link, no dropdown — there is
no existing "account menu" component to extend) gains a second link when `me.guilds.length >
0`: "My guild" → the first entry's `/guild/<region>/<ruleset>/<name>` (the `ORDER BY
m.refreshed_at DESC` change in §2.6 makes this the most recently active membership). A
player in more than one guild sees only their most-recent one here; the full list, if ever
needed, lives on `/account`. Homepage (`index.astro`): the same "My guild" link appears in
the signed-in state of whatever hero/tools-grid section already reads `GET /v1/me`, styled
as a secondary button per the design system.

### 4.4 Claim and settings pages

New routes, all served by the Worker's existing `/guild/` prefix shell (`worker.ts:358`, no
new prefix needed — `shellHead` (`worker.ts:387`) needs its `parseGuildPath`-based branch
extended to recognise these three suffixes and answer a non-indexable, guild-scoped `<title>`
for each rather than falling through to a generic/blank head):

- `/guild/<region>/<ruleset>/<name>/claim` — shows current claim state (unclaimed / pending
  / claimed-by-you / claimed-by-someone-else); a claim button for an officer/leader-rank
  viewer; a confirm button for a second officer when one is pending; a release button for
  the current claimant.
- `/guild/<region>/<ruleset>/<name>/settings` — verified-officer-or-leader only (a
  non-officer visiting sees "You need to be a verified officer of this guild to see its
  settings," never a 403 page with no explanation, honest-copy rule); default visibility,
  officer rank line, invite link with rotate and the leak-warning copy from §3.3.
- `/guild/invite/<token>` — a minimal landing page: guild name, "Join as a member," and the
  sign-in prompt (`SignInPrompt.svelte`, reused, the same component `Account.svelte`
  composes with today) when signed out.

### 4.5 The consent control on the account page

`Account.svelte`'s `account` mode gains a "My guilds" block, modelled directly on its
existing `onAnonymize` pattern (`Account.svelte:14-20`): one row per `me.guilds` entry, a
select (`roster | gear | gear and bags`) calling `PATCH /v1/guilds/{id}/members/me`, and a
"Leave" action calling `DELETE .../members/me`, with the same optimistic local update
(`me = {...me, guilds: ...}`) the anonymize toggle already uses.

## 5. Testing

- **Addon (`addon/tests/`)**: `codec_fs1_spec.lua` — the two new shared vectors (§1.4) via
  `spec_helper.vectors()`, plus one hand-written case for "malformed rank refuses the whole
  code" (distinct from "unknown section is ignored"). `export_spec.lua` —
  `Export.guildInfo()` returns `nil` when `GetGuildInfo` is stubbed to return `nil`
  (`wow_mock.lua` gains a `GetGuildInfo` stub), and the right table when stubbed guilded.
- **Web (`web/src/lib/planner/fs1.test.ts`)**: the same two shared vectors; `guild` absent
  stays `undefined` through a round trip, never `{}`.
- **API**:
  - `api/internal/addon/addon_test.go`: `PutExports` creates a `guild_characters` row per
    character from a guilded export, and the derived `guild_members` row follows (§2.2's
    `recomputeMembership`); **two characters of one account, both guilded into the same
    guild, both appear** as their own `guild_characters` rows and the account's single
    `guild_members` row reflects the higher of the two ranks; **an unguilded alt in the
    same batch leaves the main's row alone** — asserted directly, since with no sticky
    anchor and no batch-order dependency this is now true by construction (each
    character's export writes only its own row), not by a rule that could be gotten wrong;
    **a transfer moves only that character** — character A's row moves from guild 1 to
    guild 2, character B of the same account stays in guild 1 untouched, and the account
    ends up with two `guild_members` rows (one per guild).
  - `api/internal/reports/handler_test.go`: **the spoofing regression test** — an account
    with an unverified `guild_characters` row (`verified_at = null`) must not see a
    `guild`-visible report (`mayView` returns false, via `auth.Store.GuildRank`'s new
    `verified_at is not null` clause) and must not attach a report to the guild (`mayEdit`
    false); the row does not become verified merely by existing. A single appearance in the
    guild's own uploaded report does **not** verify it (the raised bar, §3.3); a second
    appearance on a distinct report date within 30 days does; a second appearance more than
    30 days after the first does not. Existing tests at `handler_test.go:166-180` and
    `:211-260` (officer attach) are updated to seed `verified_at` explicitly (via a direct
    `guild_characters`/`guild_members` seed, not two report uploads), since they currently
    pass with the pre-spec unverified behaviour this spec removes.
  - New `api/internal/guilds/*_test.go`: claim happy path (GM immediate), claim pending +
    second-officer confirm, claim pending + GM auto-confirm, claim expiry, release on leave,
    invite rotate + accept + revoked-token rejection (including the synthetic
    `account:`-keyed row an invite-only join creates), officer approve sets `verified_at`
    immediately regardless of appearance count, officer remove clears a character's row and
    is not auto-restored by a later unguilded-then-reguilded sync, consent PATCH gates a
    `/home` query (seed two characters at different consent levels — consent lives on the
    account, so this means two *accounts* — assert the response omits gear for the
    `roster`-consent one), officer-threshold PATCH re-derives existing character ranks and
    recomputes every account's `guild_members` row in one pass.
  - Migration test: `0018` up/down round-trips cleanly against a seeded `0005` database
    (the existing migration test harness pattern, whatever it is named under
    `api/internal/db/migrations`), including the `guild_characters_character_idx` unique
    constraint (inserting the same `character_key` under two different `guild_id`s must
    fail).
- **E2E** (`web/tests/e2e/`): sign in with a fixture export carrying a `guild` section →
  `/guild/.../` shows the signed-in home; claim → settings → rotate invite → accept as a
  second account → new member appears unverified on roster; consent change hides/shows gear
  on the home for that row.

## 6. Lane split

Three lanes, file ownership stated so they cannot collide; order constraints below the
table.

| Lane | Owns |
|---|---|
| **Addon+codec** | `addon/ForeverSixty/Export.lua`, `Codec.lua`, `Locale.lua`; `web/src/lib/planner/fs1.ts` (the `guild` branch only — Lane B/one-product-design owns the rest of that file, no conflict, different lines); `tools/gen-codec-vectors.mjs`; both `codec-vectors.json` fixture copies; `addon/README.md` (spike rows + manual checklist line); `addon/tests/codec_fs1_spec.lua`, `wow_mock.lua`; `web/src/lib/planner/fs1.test.ts` |
| **API** | `api/internal/db/migrations/0018_*`; `api/internal/addon/addon.go` + new `guild.go`; `api/internal/auth/store.go` (the `GuildRank` SQL clause, `Guilds` ordering, `Guild.Consent`/`Verified`), `handler.go` (`Me` shape); new `api/internal/guilds/` package entirely (claim, settings, invite, approve/remove, `/home`, `IsMember`); all new/updated `*_test.go` in these packages. **`api/internal/reports/handler.go` and `store.go` are not touched** — `mayView`/`mayEdit` tighten automatically because `auth.Store.GuildRank`'s data source changed underneath them, not their own code (§3.1) |
| **Web** | `web/src/components/Guild.svelte` (signed-in sections: per-character roster with approve/remove for officers), new `GuildClaim.svelte`, `GuildSettings.svelte`, `GuildJoin.svelte`; new small helper (e.g. `web/src/lib/guild/roster-links.ts`) that picks `current-character.ts` when present and `handoff-links.ts` otherwise, per §4.1; `web/src/pages/guild/**` (new routes + fixture paths); `web/src/components/SessionNav.svelte`, `web/src/pages/index.astro` ("My guild"); `web/src/components/Account.svelte` (consent block); `web/src/worker.ts` (`shellHead` branch for the three new sub-paths); `web/src/lib/characters.ts` (any new path-parsing helpers); e2e specs |

**Order constraints:**

1. **API before Web** for anything Web reads (`/home`, `/settings`, `/claim`, `Me.guilds`
   shape) — Web may build against the shapes in §2.6 ahead of the API landing, same
   allowance `one-product-design.md` §2 gives Lane A for `current-character.ts`.
2. **Addon+codec is independent of both** — it only touches the wire format and the shared
   TS decoder function signature (`FS1Build.guild?`), which API's `guild.go` parser (§2.2)
   is written against directly (the wire syntax in §1.2), not against `fs1.ts` — so API does
   not block on Addon+codec landing, and vice versa.
3. **Migration `0018` is the one hard sequencing point**: API's own handlers, tests, and
   Web's `/settings` and `/claim` pages all assume its columns exist. Land it first within
   the API lane, before the rest of that lane's PRs.

## 7. Out of scope, recorded

Payments and entitlements (proposal §4.2); officer tools (3.3: raid readiness board, loot
council helper, attendance/performance, raid-night comparison, assignments); the
performance analyzer (3.4, its scoring model, percentiles, per-role weights, the three
pages); ratings in game (3.5, the nightly snapshot job, companion auto-update,
CurseForge/Wago data-addon release). This spec's `guild_members.consent` column and
`gear_bags` value exist so 3.3's later readiness board has somewhere to read consent from
without a further migration — nothing here builds what reads it.
