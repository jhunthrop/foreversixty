# Sim-accuracy round 5: page lane (2026-11-01)

The /sim landing page, `web/src/lib/sim/settings.ts` and /tiers/healer, against the three new owner rulings. Site branch `wt/r5-page`; no engine change.

## Findings

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| P1 | **The /sim consumables were the page's own two lists** (round 4 I6). The page sent one physical list (Titans flask, Winterfall, Sunfruit, Scorpok, two Elemental stones, rage potion) or one caster list; the published level-60 numbers send the curated raid group for the spec (`data/curated/presets.json`) plus the class kit (`sim/leveling/kit.go`) as `request.Presets.Resolve` and `ResolvedPreset.Layer` combine them. Rogues had no poisons. `sim/request.RaidConsumes` now resolves that set per spec, `go run ./internal/genraidconsumes` writes `sim/request/raid-consumes.json` (held current by `TestRaidConsumesJSONIsCommitted`), and `web/scripts/sync-raid-preset.mjs` copies it into the page's build. The Raid-buffed preset, its summary and the consumable candidates key off the spec; the controls are unchanged. | `settings.test.ts` (defaults equal the generated file for all 28 specs) and `character.test.ts` (the request sends them untouched) failed first; Go `TestRaidConsumesAreThePublishedSet` compares against Resolve+Layer directly; engine table below | CONFIRMED-BUG (page) | S 72c6c5b3e, S cf707e3bf |
| P2 | **A stone beside a rogue's poison was filed with the potions** in the buff panel and the preset summary (`main_hand_stone`, `off_hand_stone` had no section). Now under weapon oils and stones. | `preset-summary.test.ts` failed first | CONFIRMED-BUG (page, cosmetic) | S 72c6c5b3e |
| P3 | **Power Infusion in a default.** Checked: not in `presets.json` (raid buffs or any consumable group), not in the generated per-spec sets, not in `defaultSettings` or `withPreset` for any spec and preset, not in an empty API request (`buffsFor(nil)` gives 0), not in the coach reference (`simRefRequest` sends no buffs). The healing-priest options' `PowerInfusionTarget` is self: the priest's own talent, not an external. The /sim buff panel already lists every IndividualBuffs id from IDS.md, so `power_infusions` is offered, off by default. | `TestNoRaidDefaultCarriesPowerInfusion`, `TestPowerInfusionOnlyWhenTheRequestNamesIt`, settings.test.ts (passed first: no default carried it) | Aligned (tests added) | S 72c6c5b3e |
| P4 | **/tiers/healer ordered healers by any margin.** Healers ranked 1 to 5 on the published HPS with the DPS list's fixed 1% tie mark, and the healer band carried no error. At the current ×1.65 profile the published gear measures Holy Priest 640.0 and Restoration Shaman 641.1 (Alliance) and Discipline 542.8 and Holy Paladin 541.3 (Horde); the ×1.5 files order those pairs by about 7% and 2%, so the order came from the tuned profile. The ranker now publishes `metrics.hps_error` (the same run's standard error); the tier list ties neighbours whose gap is within max(1%, 2·sqrt(e1²+e2²)), chains the tie, gives tied healers one rank and the existing "≈ tie" mark titled by the error, and keeps a gap figure on a healer tied for the top. DPS and tank lists are unchanged. | `tier-list.test.ts`, `tier-copy.test.ts`, `score_heal_test.go` failed first; spec-breakdown -heal, 2000 iterations, seeds 1 and 50001 (spread at most 0.2%) | CONFIRMED-BUG (page) | S 783aa84f6, S 935f197f7 |

## Page against published (band 60 raid, horde, 10,000 iterations)

"page" runs spec-breakdown with `-consumes` set to the page's default list for the spec; "old page" with the list the page sent before P1. The request equality itself is proven by the tests in P1 (same ids for every spec; buffs are the raid list, which already holds every kit buff at 60).

| spec | seed | published | page | old page |
|---|---|---|---|---|
| Combat | 1 / 50001 | 622.0 / 621.4 | 622.0 / 621.4 (0.00%) | -7.69% / -7.60% |
| Survival | 1 / 50001 | 650.2 / 650.0 | 650.2 / 650.0 (0.00%) | +3.80% / +3.79% |
| Fury | 1 / 50001 | 778.8 / 779.1 | 778.8 / 779.1 (0.00%) | +0.61% / +0.65% |
| Fire | 1 / 50001 | 618.9 / 620.0 | 618.9 / 620.0 (0.00%) | +3.89% / +3.94% |
| Destruction | 1 / 50001 | 682.2 / 682.2 | 682.2 / 682.2 (0.00%) | +0.29% / +0.33% |
| Marksmanship | 1 / 50001 | 665.7 / 666.2 | 665.7 / 666.2 (0.00%) | -1.75% / -1.80% |

Standard error at 10,000 iterations is 0.3 to 0.7 DPS.

The planner island grew about 0.7 KB gzipped with the per-spec table; its budget moves from 82 to 83 KB (S cf707e3bf, the real-data build read 84,126 bytes).

## Gates

Full web vitest 0 (702 files, 7,912 tests); astro check 0 (0 errors); `npm run lint` 0; Playwright sim-* (less the real-engine specs), tiers, planner-share-sim and header-finished-sim, desktop and mobile, 0 (484 passed, 14 skipped); real-data build and postbuild 0; Go `sim/request` and `sim/cmd/leveling-bis` 0. No shared layout touched.

## Screenshots

In the lane scratch `r5/page/shots/`: `sim-rogue-combat-{1440,2000}.png` and `sim-hunter-survival-{1440,2000}.png` (the Raid-buffed summary open: the rogue's Deadly and Instant Poison with a stone beside each and Thistle Tea; Survival's ranged group and Major Mana Potion); `tiers-healer-{alliance,horde}-{1440,2000}.png` on the committed (×1.5, no error) files; `tiers-healer-*-x165-demo.png` on a local, uncommitted overlay of the ×1.65 seed-1 measurements with a 0.12% error, which shows the tied pairs (Alliance ranks 1 2≈ 2≈ 4 5, Horde 1 2 3 4≈ 4≈).

## Still open

- The ranker lane owns `sim/cmd/leveling-bis`; this lane added one field there (`hps_error` in `score_heal.go`), additive, for the integrator to merge with the ranker's branch. The tier json picks it up at the next nightly; until then /tiers/healer ties on the 1% floor, which already ties both pairs above once the ×1.65 files publish.
- The hybrids lane's ×1.5/×1.65 sensitivity: if a pair flips between profiles by more than 1% plus error, the floor should be raised to that figure (`TIE_MARGIN_PERCENT` floor in `tier-list.ts`, one constant `HEALER_ERROR_Z` in `tier-rules.ts`).
- The page's preset is the level-60 set at every level. Below 30 a rogue's kit is two Instant Poisons (`KitConsumes`), and the page still sends Deadly Poison; below 60 the published numbers are bare-preset bands, so there is no published level-60 raid figure to agree with there.
- Protection paladin's kit Devotion Aura is not in the raid buff list the page sends (kit buffs only matter for that spec at 60); the tank list is not on /sim's DPS comparison.
