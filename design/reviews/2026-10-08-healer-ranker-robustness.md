# Healer ranker robustness, 2026-10-08

Lane heal-rank. Follows design/reviews/2026-10-08-heart-of-the-lion.md item 6 (restoration druid's new pick published a lower HPS than its old set) and 2026-10-08-ranker-set-completion.md. Two commits: the first made the guarded score honest and added the incumbent check; the second, after the coordinator's ruling, removed the guard.

## 1. Reproduction (restoration druid, band 60 raid, engine f3a673d10)

The committed set is `data/builds/1.60.1.70009/bis/druid-restoration.json` (built on engine 35be2e187); "before" is the unmodified ranker at main. Both sets simmed on the same engine, talents and seeds (7, 11, 13 agree to the first decimal), 2000 iterations, one standard error.

| Faction | Set | Effective HPS | mana_lasts_sec | Overheal | Guarded score (the old objective) |
|---|---|---|---|---|---|
| alliance | committed set | 622.2 +/- 0.3 | 275.0 | 3.7% | 523.0 +/- 0.3 |
| alliance | before pick | 598.6 +/- 0.3 | 280.9 | 3.5% | 524.8 +/- 0.3 |
| horde | committed set | 619.3 +/- 0.3 | 272.9 | 3.7% | 512.5 +/- 0.3 |
| horde | before pick | 583.6 +/- 0.3 | 284.4 | 3.4% | 524.5 +/- 0.3 |

## 2. Cause

1. **The objective double counted mana.** The old objective was effective HPS times the squared share of the 300 s fight the mana lasted. Effective HPS over the whole fight already pays for running dry (a set that empties at 270 s heals nothing for the last 30 s and its average falls), so the guard counted mana twice. By that objective the before pick was not worse (524.8 against 523.0, and 524.5 against 512.5); by the quantity a player reads it healed 3.8% and 5.7% less. Ruling, 2026-10-08: the healer objective is effective HPS alone. The guard's original purpose, separating sets when nobody ran dry, is gone under the retuned profile, where mana is a live constraint for every healer.
2. **The guarded score's error bar was too small.** It scaled only the healing's error, not the mana-lasts term (about 17 s of spread per iteration), so a 300-iteration score showed +/-0.7 when its real error was 3.7 to 4.5. The horde lead looked like 4% and was 2.3%. Moot once the guard is gone, but it is why the ranker's order held only by luck of margin.
3. **No incumbent check.** Nothing compared the new set with the set already published, so a chain of locally better picks could end below it.

## 3. Changes

- score_heal.go: `healEngine` returns effective HPS and its standard error; `lastingFactor`, `healingScore` and the mana-lasts error are deleted. `mana_lasts_sec`, overheal, raw HPS and healing per mana stay published as metrics. The package comment says why mana is not a second term. Nothing on the site described the squared guard (the healer panel only prints the mana metric), and data/curated/heal-profile.json already states the retune rule in terms of `mana_lasts_sec` as a constraint, so no copy changed.
- `healIterationScale = 7`: a healer's ranking runs use 7 times the requested iterations (300 becomes 2100, 100 becomes 700). Healing's own error is about 0.05 to 0.1% of the score at that count; published metrics run 2000 iterations.
- incumbent.go, `keepIncumbent`: after the last pass a healer band reads the previously published set for its band, preset and faction from the committed report (before this run can overwrite it), simulates it and the new set under the final harness, and keeps the incumbent when it beats the new set by more than the two errors combined. Unreproducible or identical incumbents cost no run. A kept band publishes `kept_incumbent` (both effective HPS and errors, the differing slots). Healers only.
- Tests (synthetic): the incumbent kept ahead beyond error, not inside it or when worse, no run when identical or unreproducible, a slot the incumbent left empty, report loading; a set that runs dry is scored on its effective healing alone; the error is under half the adoption margin at the healer harness and the iterations are scaled.

## 4. Results, band 60, all five healers

Cells are published HPS / mana_lasts_sec. Committed is the repo file (engine 35be2e187), before is the ranker at main, after is this branch; before and after are the same engine, data and seeds. The last column counts slots whose item differs from the committed set.

| Healer | Preset | Faction | Committed | Before | After | Kept incumbent | Slots changed |
|---|---|---|---|---|---|---|---|
| druid-restoration | raid | alliance | 600.0 / 271 | 598.5 / 281 | 622.2 / 275 | kept | 0 |
| druid-restoration | raid | horde | 596.5 / 269 | 583.3 / 285 | 619.3 / 273 | kept | 0 |
| druid-restoration | bare | alliance | 324.7 / 228 | 324.7 / 228 | 324.6 / 227 | kept | 0 |
| druid-restoration | bare | horde | 297.7 / 241 | 297.7 / 241 | 321.8 / 215 |  | 5 |
| paladin-holy | raid | alliance | 510.0 / 239 | 518.0 / 249 | 525.4 / 235 |  | 7 |
| paladin-holy | raid | horde | 508.8 / 236 | 516.5 / 249 | 523.2 / 234 |  | 5 |
| paladin-holy | bare | alliance | 175.3 / 75 | 173.1 / 76 | 181.0 / 73 |  | 3 |
| paladin-holy | bare | horde | 173.6 / 74 | 173.6 / 74 | 179.3 / 72 |  | 3 |
| priest-discipline | raid | alliance | 583.9 / 188 | 628.3 / 209 | 618.6 / 201 |  | 7 |
| priest-discipline | raid | horde | 580.3 / 184 | 622.6 / 209 | 613.3 / 196 |  | 7 |
| priest-discipline | bare | alliance | 373.4 / 101 | 373.4 / 101 | 373.5 / 101 | kept | 0 |
| priest-discipline | bare | horde | 367.7 / 100 | 367.7 / 100 | 367.8 / 100 | kept | 0 |
| priest-holy | raid | alliance | 659.2 / 231 | 707.6 / 257 | 702.8 / 255 |  | 1 |
| priest-holy | raid | horde | 650.7 / 227 | 682.9 / 244 | 679.7 / 243 | kept | 0 |
| priest-holy | bare | alliance | 331.2 / 96 | 331.2 / 96 | 330.9 / 96 | kept | 0 |
| priest-holy | bare | horde | 323.2 / 92 | 323.2 / 92 | 323.2 / 92 | kept | 0 |
| shaman-restoration | raid | alliance | 568.4 / 291 | 576.2 / 333 | 582.6 / 356 | kept | 0 |
| shaman-restoration | raid | horde | 567.8 / 295 | 573.5 / 313 | 582.3 / 358 | kept | 0 |
| shaman-restoration | bare | alliance | 211.7 / 79 | 211.7 / 79 | 212.9 / 78 |  | 2 |
| shaman-restoration | bare | horde | 211.2 / 78 | 211.2 / 78 | 210.9 / 78 | kept | 0 |

Under the final harness (2100-iteration effective HPS, same seed, the committed set simmed on the current engine against the after set) no row of the 20 is below the committed set beyond error (two paladin raid rows gain 0.8 and 0.2, inside it). The largest gains: priest-holy raid alliance +10.3, priest-discipline raid +13.7 and +14.2, restoration druid bare horde +24.3, paladin-holy bare +5.9 and +5.8. The rows that kept the incumbent are the incumbent's own set, so they equal it by construction. The incumbent check fired on 11 rows; on restoration druid raid, both factions, it kept a set 15 and 13 effective HPS above the pick the ranker had just made (622.3 against 606.8, 619.2 against 605.8), the regression that started this lane. The published figures of the rows that kept the committed set differ from the committed numbers only by engine (the committed druid raid file says 600.0 and 596.5; the same sets measure 622.2 and 619.3 on the current engine).

Honest limit: priest-discipline raid publishes 618.6 and 613.3, below the 628.3 and 622.6 the earlier guarded run reached (a set that carried more mana, which lifts effective HPS here). The per-slot swap pass promotes a runner-up only at +1% beyond error, so a chain of sub-1% swaps toward mana stats is not followed. Both are well above the committed 583.9 and 580.3. Lowering the swap margin for healers is the open follow-up.

## 5. Spot check, tanks and DPS

druid-feral-bear, paladin-protection, warrior-protection, hunter-marksmanship, mage-fire and rogue-combat, band 60, binary at main against the first commit of this branch: all four bands of every spec publish identical set_dps and identical items in every slot. The damage and tank path is unchanged by construction; the only shared edit is the `kept_incumbent` field, omitted when absent. (The second commit touches only the healer engine.)
