# Sim-accuracy tools lane, 2026-11-12

Scope: two measuring-tool bugs the audit lanes reported (the `spec-breakdown` seed, from the shamans lane 798e6452e; raid default builds at 50 points, from the addon tour lane bab7005d5), plus the site's spec-from-points tie rule (forever-coord followups). All Go tests and sims on the Windows box through `tools/remote-test.sh` (engine 5f1b39450, `nice -n 15`, GOMAXPROCS=5, one sim at a time).

## 1. The seed reaches the engine (not a seed bug; a print-precision bug)

**Report.** `spec-breakdown -seed 50001` printed the same DPS as `-seed 1` for every race in the shamans lane's Elemental 39 runs (`shamans/group_g_lb39_seed2.sh`).

**Trace.** `-seed` → `options.seed` → `api.SimRequest.RandomSeed` (`sim/cmd/spec-breakdown/main.go` buildRequest) → `request.BuildWith` → `proto.SimOptions.RandomSeed` (`sim/request/request.go`) → the engine's `SplitSimRequestForConcurrency` offsets each split's seed by the iterations before it, and the run loop reseeds every iteration at `RandomSeed + i` (fork `sim/core/sim_concurrent.go`, `sim/core/sim.go` at 5f1b39450). Nothing drops it.

**Measured** (shaman-elemental, bare, band 30, level 39, Horde, Lightning Bolt rank 4, 10,000 iterations, `-json`):

| Race | seed 1 | seed 50001 | seed 1 again |
|---|---|---|---|
| Windshaper Skyborne | 72.482 ± 0.069 | 72.465 ± 0.070 | 72.482 |
| Tauren | 74.198 ± 0.064 | 74.211 ± 0.066 | 74.198 |

Tauren over eight seeds (1, 50001, 100001 ... 600001): 74.13 to 74.26, spread 0.04 against a reported error of 0.065. The seeds are independent and the error bar is honest (slightly conservative).

**Cause of the misread.** The markdown header printed `DPS %.1f +/- %.1f`. Two seeds 0.02 DPS apart printed 72.5 and 72.5, and the lane's script greps that header.

**Fix.** The header prints DPS and its error to two decimals (`shaman-elemental  DPS 74.21 +/- 0.07`). A new test, `sim/cmd/spec-breakdown/seed_test.go` `TestSeedReachesTheEngine`, runs the whole flag-to-report path at 40 iterations for rogue-combat raid 60 and the shamans lane's Elemental 39: seed 1 twice must match, seeds 1 and 50001 must differ. It passes at the current pin (it passed before the fix too: the seed was never dropped, so this is a guard, not a red test).

**Past measurements.** None is invalidated. Every "mean of seeds 1 and 50001" in the audit reports was two real seeds. Where a report says two seeds "agree to the decimal", read it as agreeing within the one-decimal print, which is inside the sim's own error. No golden, ranked output or ledger figure moves: the change is a print format in one CLI.

## 2. Raid default builds total 51 points; the tour mock dropped one

**Report.** The tour boards ("Paladin raid builds" 31/11/8, 33/10/7, 31/12/7; the druid rows) totalled 50 points at 60.

**Trace.** Every source carries 51:
- `data/builds/1.60.1.70790/addon-builds.json`: every spec's leveling and raid `order` has 51 triples, none over its talent's max rank (checked against `talents/<class>.json`).
- `addon/ForeverSixty/Data.lua` (9c8f96ec5): every `defaultBuilds` order has 51 triples.
- The BiS raid-60 entries (`bis/*.json`, both factions): every talent string sums to 51.
- Guarded already: `data/tests/test_addondata.py::test_every_weighted_spec_has_a_51_point_leveling_and_raid_build` and `web/src/content/guides/builds.test.ts` ("spends exactly 51 points").

**Cause.** `design/mocks/addon_tour.py` read each order with a lazy `order = \{(.*?)\} \}`, which stops inside the last triple (`...{3,1,2} } }`) and drops it. Paladin Holy raid is 31/12/8 in the data; the mock showed 31/11/8.

**Fix.** One `order_triples` helper reads the list triple by triple to its own brace and raises unless it holds 51 points; `default_builds` and `default_order` both use it. `builds-dialog-60`, `builds-dialog-druid` and `builds-dialog-own` re-rendered (now 31/12/8, 33/10/8, 31/12/8).

No talent choice is involved, so no sim and nothing for the owner. Data.lua was not regenerated or committed; `addon-data --check` is unaffected.

## 3. The site reads a spec off points by the addon's rule

The web's `specFromPoints` and the API's `character.SpecFor` named the tree with the most points and tied to the first, so a 2 Holy / 2 Retribution paladin read Holy. The addon's `Talents.decideTree` (owner 2026-11-11: "a wrong silent guess is worse than one question") asks instead.

**Fix.** Ported to both: a tree is named when at least 2 points are spent and all sit in it, or when it holds more than half the points and leads every other tree by 5 (`SURE_LEAD`, `SURE_MIN_POINTS`, mirrored as constants with the same test cases).
- Web (`web/src/lib/character-spec.ts`): `decideTree`; `specFromPoints` returns null when the points ask, so every surface falls back to the spec the addon reported (the player's own answer), as it already did with nothing spent.
- API (`api/internal/character/specfor.go`): `DecideTree`; `SpecFor` returns the new source `"ask"` with the class's default spec as a placeholder for callers that need one (gear picks, sim input, home). The guild readiness board (`exportSpecName`) names no spec for `"ask"`. `gear_spec_source` in the web types gains `'ask'` (nothing renders it today).
- One existing web test moved with the rule: a seen build with 1 point (`1/0/0`) now keeps its filed spec ("Holy") instead of reading Discipline.

Out of scope and unchanged: the iOS and Android apps' own derivations, if they have one.

## Gates

On the box: `go test ./cmd/spec-breakdown/` (sim, pinned engine); `go test ./internal/character/ ./internal/guilds/ ./internal/rating/ ./internal/sims/ ./internal/activity/ ./internal/rankings/` (api): all pass except `guilds` `TestAppContractRaidTabReads` (`character.ready` items lack `icon`), which fails identically on origin/main 9c8f96ec5 and is not this lane's; vitest over `src/lib/character*`, `account`, `bis`, `guides`, `rankings`, `home`, `gear` (2,369 tests); eslint and tsc clean on the touched web files.

## Shas

See the ledger row.
