# Rogue energy model, rogue talents and Dark Sacrifice, 2026-10-08

Lane `energy`. Fork branch `energy` (from `forever`), site branch `energy` (from `main`). Evidence: design/reviews/2026-10-08-beta-evidence.md, item 1 (Blizzard's 30 September class-highlights video: energy "will no longer return in clearly separated intervals and will instead regenerate continuously"; a tester bug report on GitHub agrees; no per-second rate is published; Vigor raises the maximum).

## 1. Energy

What changed (`sim/core/energy.go`): the 2.02 s tick (20.2 Energy per tick, phase drawn at random per iteration) is gone. Regeneration is continuous at the same mean rate, 10 Energy per second (`EnergyRegenPerSecond`). No rate change is modelled, because the evidence states none; 10 per second is the Classic mean carried over, and the evidence file already calls it an inference.

Choice: **on-demand accrual**, not a periodic accrual. Every read (`CurrentEnergy`), spend, gain and multiplier change first brings the stored value up to `sim.CurrentTime` at the current rate, so the value is exact at any instant and a spend at 0.35 s leaves the fractional 3.5 intact. The engine's scheduler cannot do this by itself (a task runs only when it asks to), so a small wake task remains for one job: telling the APL when energy has reached the next decision threshold. For a rogue the thresholds are the precomputed energy costs and `currentEnergy` comparisons of its APL, so the wake lands on the exact nanosecond the threshold is reached (plus 1 ns against float error). Units without precomputed thresholds (cat form) are polled every 100 ms (`EnergyPollInterval`). A full bar, or a bar with no APL, schedules nothing. Reads accrue silently and never call into the APL, so there is no re-entrancy.

Kept: maximum 100 plus Vigor (5 per rank, tested), a set bonus's +10 maximum, Relentless Strikes (25 on a per-combo-point chance), Thistle Tea, refunds. Adrenaline Rush is a regen multiplier (+100%, `AddEnergyRegenMultiplier`) that flushes accrual at the old rate before it changes. Druid cat form shares the bar and follows.

Dropped: `ResetEnergyTick`, `NextEnergyTickAt`, `EnergyTickDuration`, `EnergyPerTick` (no tick exists). The APL value "Time to Next Energy Tick" stays in the proto for old saved rotations and is dropped with a validation warning; no curated rotation uses it. The legacy feral rotation (`rotation.go`, not in any curated APL) used the tick constants as if one tick were one Energy (a leftover from the 100 ms, 1 Energy tick of the original engine, so its waits were 20 times too long); it now uses `EnergyForTime` and `TimeForEnergy`, and plans one poll step ahead.

Tests (`sim/core/energy_test.go`, written to the spec): 10.0 ±0.1 after 1.0 s from empty, 3.5 after 0.35 s, never above the maximum, Adrenaline Rush doubles the rate (and halves back), a spend at 0.35 s keeps the fractional accrual, the wake lands on the next threshold (and halves under double regen), no wake when full or without an APL, the 20.2 per 2.02 s conversion round trip.

## 2. Rogue talents against the live tree (build 1.60.1.70009)

An earlier lane (fork commits c043f79cc and e4b87c1e1, `sim/rogue/client_values.go`) had already moved Dual Wield Specialization (5% per rank, 25% at 5/5, was 50%) and Improved Eviscerate (7/13/20%, 20% at 3/3, was 15%) to the live text, with tests; both are confirmed against the ranks' text here and need no change. The pass over every other talent found one open mismatch, Improved Expose Armor, fixed here (tests first: the cost test failed before the change).

| Talent (ranks) | Live value | Engine value | Status |
|---|---|---|---|
| Improved Gouge (3) | +0.5 s Gouge per rank | none (no Gouge spell) | n/a, no DPS lever |
| Remorseless Attacks (2) | +20/40% crit after a kill | not modelled (no kill hook) | documented gap |
| Malice (5) | +1% crit per rank | +1 per rank | right |
| Ruthlessness (3) | 20% per rank, extra combo point | 0.2 per rank | right |
| Murder (2) | +2/4% vs Humanoid and Giant | 2% per rank, those types | right |
| Improved Slice and Dice (3) | +15% duration per rank | 1.15/1.30/1.45 | right |
| Relentless Strikes (1) | 20% per combo point, 25 Energy | 0.2 per point, 25 | right |
| Improved Expose Armor (2) | -5/-10 Energy, refund 1/2 combo points at 5 | was Classic +25/+50% armor reduction, cost 25 | **fixed**: cost 25-5r, refund r points at five, no armor bonus |
| Lethality (5) | +4% crit damage per rank | 0.04 per rank | right |
| Vile Poisons (5) | +4% poison damage per rank | 1.04..1.20 | right |
| Cold Blood (1) | +100% crit on next of five named strikes | 100% crit, the five | right |
| Improved Poisons (5) | +2% apply chance per rank; 10% per rank to keep a charge | +2% per rank; charges not modelled | chance right, no-consume n/a (poisons have no charges) |
| Vigor (2) | +5/+10 maximum Energy | 5 per rank | right |
| Mutilate (1) | 75% weapon + flat per weapon, +20% poisoned, 2 points | 75%, +20%, 2 points; flat from the client effect rows (noted in `mutilate.go`) | right |
| Improved Kidney Shot (2) | +5/10% damage on a stunned target | not modelled (no Kidney Shot) | documented gap |
| Seal Fate (5) | 20% per rank | 0.2 per rank | right |
| Venom (1) | +30% poison damage, +10% apply chance | 0.30, 0.10 | right |
| Improved Eviscerate (3) | 7/13/20% | 1.07/1.13/1.20 | already right (earlier lane) |
| Improved Sinister Strike (2) | -3/-5 Energy | 45, 42, 40 | right |
| Lightning Reflexes, Deflection, Precision (5, 3, 3) | +1% dodge, +2% parry, +1% hit per rank | same | right |
| Puncturing Wounds (3) | Backstab +10% crit, Mutilate +5%, 15% extra point, per rank | same | right |
| Endurance, Improved Sprint, Improved Kick (2 each) | cooldown and utility | not modelled | no DPS lever |
| Riposte (1) | 150% weapon damage | 1.5 | right |
| Flawless Execution (1) | -10 Eviscerate Energy | 10 | right |
| Dual Wield Specialization (5) | +5% off-hand per rank, 25% at 5/5 | 0.05 per rank | already right (earlier lane) |
| Blade Flurry (1) | +20% melee speed, 15 s | 1.2, 15 s (cost 25, text silent) | right |
| Hack and Slash (5) | 1% extra attack / 1% crit / 3% armor per rank by weapon | same | right |
| Weapon Expertise (2) | -1/-2% dodge and parry | 1/2 | right |
| Aggression (3) | +2% per rank to Sinister Strike, Backstab, Eviscerate | 0.02 per rank, all three | right |
| Adrenaline Rush (1) | +100% Energy regeneration, 15 s | regen multiplier +1, 15 s | right (now continuous) |
| Camouflage, Master of Deception, Improved Distract, Heightened Senses, Dirty Tricks (Sap and Blind) | stealth and utility | not modelled | no DPS lever |
| Opportunity (2) | +5/10% Backstab, Garrote, Ambush, Mutilate | 1.05/1.10 | right |
| Setup (3), Initiative (3) | 33/67/100% | 0.33/0.67/1.0 | right |
| Elusiveness (2) | -45/-90 s Vanish and Blind | Vanish only | right for Vanish |
| Improved Ambush (3) | +15% crit per rank | 15 per rank | right |
| Ghostly Strike (1) | 125% (180% dagger) | same | right |
| Premeditation (1) | 2 combo points, 20 s | 2 points; the 20 s forfeiture is not modelled | documented |
| Serrated Blades (3) | 3% armor ignored and +10% Rupture per rank | both per rank | right |
| Dirty Deeds (2) | -10/-20 Garrote and Cheap Shot | Garrote 50-10r | right (no Cheap Shot spell) |
| Hemorrhage (1) | 100% (145% dagger), +15% Rupture damage | same | right |
| Quietus (5) | +2% per rank under 35% health, three strikes | same | right |
| Cutthroat (5) | 3% per rank, 10 s | same | right |
| Thousand Cuts (1) | -3 Energy per stack, five stacks, 10 s | same | right |

Residual, not changed: the raid-panel debuff "Improved Expose Armor" in `sim/core/debuffs.go` still multiplies the armor reduction by 1.5 (a core toggle for another player's talent; the live text has no armor bonus). It does not affect the rogue's own Expose Armor any more; the toggle should be retired with the next debuff-panel pass.

## 3. Dark Sacrifice

Confirmed against the client rows (spells 1277324 to 1277328, ranks at 20, 30, 40, 50, 60). Text: "Cannibalize $o1 of your own Health over $d to gain ${$o2+$SPI} Mana." SpellEffect: aura 226 (health) and aura 24 (mana), both period 3000 ms, base 80 / 136 / ... / 320 per tick over a 15 s duration (5 ticks): 400 at rank 1, 1600 at rank 5, same on both sides. `${$o2+$SPI}` is the mana total plus 100% of Spirit, once. The engine already carried the Spirit term (one fifth per tick); no value changed. The split is now a function, `darkSacrificeManaPerTick`, with tests for the 400 and 1600 totals and for Spirit summing to exactly once over the cast.

## 4. Measurement

Fork: `go test --tags=with_db ./sim/rogue/... ./sim/druid/... ./sim/core/ ./sim/priest/...` passes. No `.results` file moved: the rogue suites (TestCombatSinisterStrike, TestCombatDaggers) and TestP1Feral are skipped awaiting the Forever talent rewrite (`SkipAwaitingForeverTalentRewrite`) and their recordings are marked UNVALIDATED; nothing was regenerated or adopted. Site: `go test ./...` under `sim/` passes.

Rotation search (`rotation-search -preset bare|raid`, seed 7, 800-iteration baselines), curated rotation DPS, before (engine `forever`) and after (this lane's engine):

| Spec | Bare before | Bare after | Raid before | Raid after |
|---|---|---|---|---|
| Assassination | 300.8 | 302.0 | 646.2 | 646.5 |
| Combat | 270.0 | 269.8 | 585.6 | 585.0 |
| Subtlety | 266.3 | 273.7 | 570.2 | 589.6 |
| Feral (cat) | 303.9 | 303.0 | 563.1 | 562.0 |

The errors are ±0.4 to ±1.2. Assassination, Combat and Feral are inside error. Subtlety gains 2.8% bare and 3.4% raid: it was the energy-starved line, and the old search's winners for it ("remove Eviscerate", +6.3 bare and +20.8 raid) were artefacts of tick phase and were already marked NOT ADOPTABLE (they drop a finisher); after the change the curated line reaches the old winner's DPS (589.6 against 591.0) and the search finds nothing for Subtlety on either preset. Energy pooling lines did not need to move. The only remaining finding is Combat raid: swapping Eviscerate and Slice and Dice, +8.8 DPS (+1.5%) beyond the ±1.3 error at 800 iterations (before the change: +6.7). **Held, not adopted:** the bare search finds nothing (+0.0), and the standing rule for the last raid-search commit was a gain on both characters. One swap in `data/curated/apl/rogue-combat.json` if the owner wants it on the raid result alone.

Reports regenerated in place: `design/reviews/rotation-search/rogue-{assassination,combat,subtlety}{,-raid}.md`. `make apl-sync` / `make apl-check` (ENGINE_DIR the fork worktree): no rotation changed, all copies match, the fork tree stayed clean. Ladder goldens regenerated for the four specs (`sim/request/testdata/ladder/{rogue-assassination,rogue-combat,rogue-subtlety,druid-feral}.golden.md`); the level 10 to 60 rows move by about 1% in DPS (more at the lowest levels) and shift cast mixes.

Ranker (`leveling-bis -spec <s> -bands 60`, scratch output, set DPS verified, before to after):

| Spec | Bare A / H | Raid A / H |
|---|---|---|
| Assassination | 288.4 to 289.6 / 290.1 to 289.8 | 682.3 to 680.2 / 676.1 to 675.3 |
| Combat | 265.1 to 261.4 / 264.7 to 265.6 | 659.4 to 658.9 / 654.8 to 652.4 |
| Subtlety | 256.4 to 263.5 / 257.5 to 262.6 | 547.5 to 571.7 / 547.3 to 567.2 |
| Feral | 303.1 to 302.9 / 305.2 to 304.6 | 628.3 to 625.7 / 629.6 to 627.3 |

Picks changed only among equal-value choices (trinkets and an off-hand within the sim error), except Feral raid: the Feralheart set completion is no longer adopted (set DPS 628.3 to 625.7). That is the set-completion screen flipping inside noise; the nightly ranker regenerates the published picks, and this lane commits none (nightly-owned data).

## 5. To do on merge

The site pins the engine by commit (`make engine-pin`); this lane did not touch the pin or `sim/go.mod` (the worktree's replace points at the fork worktree locally and is not committed). After the fork's `energy` branch merges, pin it and re-run the nightly ranker.
