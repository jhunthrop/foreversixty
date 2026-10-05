# Guild control centre — API lane report

Contract: `docs/contracts/2026-10-04-guild-centre-api.md`. All 7 build steps shipped, one
commit each on branch `api-guild-centre`.

**Follow-up (same branch):** loot candidates no longer read empty on real data. Added a
tier-1 fallback to `loot.go`'s `lootCandidates`: when a roster character has no tier-0
(BiS-band) match for an item, they still qualify when their class can use it at all (checked
against the build's own per-class item table, `data/builds/<build>/items/<class>.json`, read
through `api/internal/trees` — the API's existing per-class item loader, wired onto
`guilds.Store` as a new `Trees *trees.Data` field) and the item's own item level beats what
they wear in that slot (or the slot is empty). Tier-1 candidates carry `gain_dps: null`,
`not_sim_checked: true`, and a new `ilvl_delta` field; tier 0 always ranks ahead of tier 1.
`docs/contracts/2026-10-04-guild-centre-api.md`'s loot section documents both tiers and the
new field. The integration test now asserts every real (non-slotless) Onyxia drop has at
least one candidate from the seeded 24-character roster — it does, entirely via tier 1,
since (per the original report below) no BiS band in this repo names a raid-tier item yet.

## Endpoints and where each number comes from

| Endpoint | Source |
|---|---|
| `GET /home` | `home.go` + `roster_extended.go`. Roster base query joins `guild_characters`/`addon_exports`/`guild_members` + a lateral `fight_metrics` row (class/spec/role/ilvl/faction, consent-gated). `professions`/`gear`/`enchants`/`bags`/talent points come from `fs1.Decode(ae.export)`, same consent gate. `attendance`/`best_parse`/`rating`/`last_report_at` are four small batch queries over the guild's own reports (`roster_extended.go`). `summary.below_rating_floor` is the 25th percentile of `rating.Overall` **guild-wide**, not per-role (see Deviations). `summary.pulls_this_tier` counts every fight ever logged for the guild — no tier-start timestamp exists anywhere queryable (see Deviations). `standing` is derived in-memory from the already-built roster slice; `claim.claimed_by_name` resolves via the new `Store.Accounts`. |
| `GET /raids` | `raids.go`. Cursor-paginated `reports`, visibility clause copied verbatim from `HomeReports`. Present-roster name/class read from `addon_exports`/`fight_metrics` directly (report data, never gated by guild consent). `top_parse` is `fight_metrics`' single best kill-fight row for that report. |
| `GET /progression` | `progression.go`. Fully public, not filtered by report visibility (matches `rankings.Store.Guild`'s own existing choice for this aggregate). Tier/raid names are hardcoded constants (see Deviations); Onyxia's own depth is a live query over `fights`/`fight_metrics`. |
| `GET /readiness` | `readiness.go` + `readiness_core.go`. One row per **verified** roster character; `computeReadiness` is the exact function the home endpoint's `standing.needs_before_next_raid` also calls. Measured latency below. |
| `GET /loot`, `POST`/`DELETE /loot/awards` | `loot.go`. Onyxia's 22 real drops are a hand-curated Go literal (see Deviations for why, and for the candidate-ranking data gap). Awards persist in the new `loot_awards` table (migration 0030). |
| `POST /roster/approve-all` | `roster.go`'s `ApproveAll` — one transaction, bulk `ApproveCharacter` + `RecomputeMembership`. |

## What is null, and why

- `class`/`spec`/`role`/`item_level`/`faction`/`gear` (internal) — null without `gear`/`gear_bags` consent, unchanged from the pre-existing roster privacy model.
- `standing` — null when the viewer has no verified, gear-consented character with a known class/spec in this guild.
- `readiness.rows[].gear_gap` — null without gear consent, *or* when the character's spec has no leveling-BiS band at all (every tank/healer spec — the catalogue only ever simmed DPS specs).
- `enchants.checked` — false without gear/gear_bags consent (see Deviations: never gated on BiS enchant data, because none exists).
- `consumables.state` — `"unknown"` without `gear_bags` consent specifically (stricter floor than gear_gap/enchants, per the membership spec's own three-tier consent).
- `best_parse`/`rating`/`last_report_at` — null/omitted when the character has no qualifying row yet.
- `loot.items[].awarded_to` — null until an officer awards that drop.

## Deviations from the contract's literal text (named, as asked)

1. **Enchant check is not gated on "the band's BiS lists an enchant for it."** Checked every `band-60-alliance` entry in every file under `data/builds/1.60.1.70009/bis/` — **no slot in any file carries an `enchant` field at all**. Implemented the design spec's own narrower EXISTS claim instead: an enchantable slot (`main_hand`/`chest`/`back`/`feet`) with gear equipped but no enchant, independent of any BiS recommendation. `readiness_core.go`'s own doc comment names this.
2. **Loot candidate ranking will read empty for every item today.** Checked every `band-60-alliance` entry in every BiS file against Onyxia's own 22 item ids — **zero matches**. The leveling-BiS catalogue only covers gear up to level 60, not raid-tier drops; no raid-tier BiS/sim data exists in this repo yet. The ranking logic itself (`namesItem`, `lootCandidateLess`) is unit-tested against a hand-built band so it's provably correct once that data exists — this is a data gap, not a code gap.
3. **`gear_gap`'s gain rule is the contract's own 3-branch shorthand, not `upgrades.ts`'s full rule.** `upgrades.ts` also falls back to a `scoreItem` stat-weight diff for a non-weapon slot with no listed alternative; the contract's own text ("pick == worn → none; worn in alternatives → gain; otherwise not sim-checked") omits that branch, and this port implements exactly the three branches stated, not the TypeScript module's fuller one.
4. **`below_rating_floor` is guild-wide, not "percentile 25 within role."** A 20–30 raider guild split by role leaves most buckets too small (2–4 raiders) for a percentile to mean anything; used one guild-wide 25th percentile instead.
5. **`pulls_this_tier` has no tier-start filter.** No raid-tier-start timestamp exists anywhere queryable (the design doc's own 9 December 2026 tier open is a future calendar date, not a stored marker) — counts every fight ever logged for the guild.
6. **Tier/raid names (`Barrow Deeps`/`Hyjal Summit`/`Onyxia's Lair`, encounter 1084) are hardcoded constants**, not read from `data/curated/loot/forever-raid-phases.json` — that file is not copied into the API's Docker image (`api/Dockerfile` only copies `data/builds/`) and doesn't carry raid display names anyway. Mirrors `api/cmd/seedguild/raid.go`'s own existing choice.
7. **`data/curated/specs.json` is read via the generated `sim/specs` Go package**, not a runtime file read — same reason as #6 (not in the Docker image), and the same pattern `api/internal/sims/score.go` already uses.
8. **BiS build is hardcoded to `1.60.1.70009`.** No "active build" config exists anywhere in this service; inventing one was out of scope. Mirrors `api/cmd/seedguild/gear.go`'s own `bisBuild` constant.
9. **Onyxia's loot table is a hand-curated Go literal**, not a per-request read of `loot.json`/`items.json`/`items/<class>.json` — those three files total over 40MB for this one build; a per-request scan for 22 fixed ids would have defeated the latency goal for no reason.
10. **Consumables "stocked"/"short" is presence-of-any-bagged-item**, not flask/food/potion identification — no curated consumables-by-item-id catalogue exists in this repo (the one that does, `logs/engine/mechanics/consumables`, is keyed by spell id for buff/cast matching, not item id, so it can't resolve an FS1 `bags=` item list).

## Readiness latency (measured, local harness, 25 raiders)

`readiness_latency_test.go`: **~15ms** for 25 raiders with full gear/enchants/professions/bags
data and `gear_bags` consent — well under the 300ms target. Logged verbatim by the test
(`go test -run TestReadinessLatencyWith25Raiders -v`).

## Store changes

`guilds.Store` gained two fields, wired in `api/cmd/api/main.go`:
- `Accounts` — resolves `claimed_by_name` (home) and loot-award `by_name` (loot).
- `DataDir` — `api/internal/bis` band reads (same `TreeDataDir` every other build-scoped reader already uses).

## Test summary (verbatim)

```
$ go vet ./...
(clean)

$ TEST_DATABASE_URL=postgres://forever:forever@localhost:5434/forever_test?sslmode=disable go test ./... -p 1
ok  	github.com/jhunthrop/foreversixty/api/cmd/api	0.336s
ok  	github.com/jhunthrop/foreversixty/api/cmd/seedguild	1.340s
ok  	github.com/jhunthrop/foreversixty/api/internal/addon	4.130s
ok  	github.com/jhunthrop/foreversixty/api/internal/auth	8.535s
ok  	github.com/jhunthrop/foreversixty/api/internal/billing	3.454s
ok  	github.com/jhunthrop/foreversixty/api/internal/bis	0.224s
ok  	github.com/jhunthrop/foreversixty/api/internal/bnetapi	0.301s
ok  	github.com/jhunthrop/foreversixty/api/internal/bnetbuild	2.923s
ok  	github.com/jhunthrop/foreversixty/api/internal/bnetimport	2.590s
ok  	github.com/jhunthrop/foreversixty/api/internal/builds	1.627s
ok  	github.com/jhunthrop/foreversixty/api/internal/card	0.516s
ok  	github.com/jhunthrop/foreversixty/api/internal/character	0.228s
ok  	github.com/jhunthrop/foreversixty/api/internal/config	0.265s
ok  	github.com/jhunthrop/foreversixty/api/internal/dataaddon	0.604s
ok  	github.com/jhunthrop/foreversixty/api/internal/db	3.177s
ok  	github.com/jhunthrop/foreversixty/api/internal/digest	0.225s
ok  	github.com/jhunthrop/foreversixty/api/internal/engine	0.290s
ok  	github.com/jhunthrop/foreversixty/api/internal/entitlements	1.959s
ok  	github.com/jhunthrop/foreversixty/api/internal/fs1	0.236s
ok  	github.com/jhunthrop/foreversixty/api/internal/guilds	~20s (per-run, see below)
ok  	github.com/jhunthrop/foreversixty/api/internal/httpx	0.234s
ok  	github.com/jhunthrop/foreversixty/api/internal/jobs	0.325s
ok  	github.com/jhunthrop/foreversixty/api/internal/mail	0.253s
ok  	github.com/jhunthrop/foreversixty/api/internal/metrics	0.360s
ok  	github.com/jhunthrop/foreversixty/api/internal/parse	2.722s
ok  	github.com/jhunthrop/foreversixty/api/internal/phase	0.228s
ok  	github.com/jhunthrop/foreversixty/api/internal/r2	0.325s
ok  	github.com/jhunthrop/foreversixty/api/internal/rankings	7.988s
ok  	github.com/jhunthrop/foreversixty/api/internal/rating	1.130s
ok  	github.com/jhunthrop/foreversixty/api/internal/reports	11.640s
ok  	github.com/jhunthrop/foreversixty/api/internal/server	0.540s
ok  	github.com/jhunthrop/foreversixty/api/internal/sims	19.568s
ok  	github.com/jhunthrop/foreversixty/api/internal/site	0.284s
ok  	github.com/jhunthrop/foreversixty/api/internal/spec	0.222s
ok  	github.com/jhunthrop/foreversixty/api/internal/subscribe	0.521s
ok  	github.com/jhunthrop/foreversixty/api/internal/textx	0.229s
ok  	github.com/jhunthrop/foreversixty/api/internal/trees	0.242s
ok  	github.com/jhunthrop/foreversixty/api/internal/zstdx	0.307s
```

Every package passes, including `sims` under `-p 1` (no deadlock).
