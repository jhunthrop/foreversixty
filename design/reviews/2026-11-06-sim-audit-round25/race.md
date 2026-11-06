# Sim-accuracy round 25: race lane

Lane `race`, site branch `wt/r25-race` (from int/r25 0a017f71d). No engine change, so there is no fork branch.

## How the call is made

- **Legal races** come from the build's client race-class table (`data/builds/1.60.1.70790/combos.json`, with races.json and classes.json). That table includes Forever's added pairs: Human hunter, Orc mage, Dwarf shaman, Undead paladin, Gnome priest, Troll warlock, and Skyborne for warrior, hunter, rogue, mage and druid (Alliance only) and for warrior, hunter, rogue, shaman and druid (Horde).
- **Harness.** The ranker's own `plainRequest` (ranker encounter, class kit, raid preset and raid enchants on the raid pass) runs on every published band set in `$R/bis-frozen`, using the published talents repositioned to the engine layout. Each race runs 10,000 iterations at seed 1 and again at seed 50001. Healers run the heal score the ranker ranks by, as the quotient of 10,000 over the scale factor, which lands at 10,003 effective iterations. Tanks are judged on **threat per second** (owner rule), not on the tank score.
- **Beyond error** means the margin is larger than the combined standard error of the two races. The rule: a **CONFIRMED-BUG** needs the same race to beat the guide race at both seeds, beyond error, by more than 1%.
- **Lane decision 1: the line is a level-60 pick.** The guide's race line is one permanent choice per faction. It is judged on band 60, using both the raid and bare passes. A race that wins only leveling bands is reported in the table, but it does not move the line.
- **Lane decision 2: Skyborne is reported but does not set the line.** Skyborne needs the paid Skyborne Heroic Pack, and every guide picks from the races every account has. Where Skyborne leads at 60 by more than 1% (Balance and Feral, Horde, raid), the guide prose already says so. I checked that balance.md and feral.md name it as the paid option.
- **Runtime:** the whole sweep is 1,032 sims per seed (28 specs, 12 band entries each, every legal race). Seed 1 took 601 s and seed 50001 took 611 s through `$R/sim.sh`, lock waits included, so about 20 minutes for both. That is far under the 25 minutes per seed the brief estimated.

## Cross-cutting (ranker code)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | The ranker never measured any race but the guide's. A guide race line could lose to another race and nothing would flag it. | `sim/cmd/leveling-bis/main.go`: `measureFinalSet` runs the guide race only | feature (seed) | - | S b49b84ffe |
| X2 | Tank threat had no standard error to judge a race on. | `sim/score/tank.go` `TankFigures` carried TPS without an error | feature | - | S b49b84ffe (adds `TPSError`; additive) |

What b49b84ffe adds:

- **`race_sweep.go`**: `legalRaces`, `sweepRaces` and the standalone `runRaceSweep`.
- **Ranker integration.** In `runSpec`, after `measureFinalSet`, every band and faction's final set is swept at the published figure's count and seed. The guide race's run is a memo hit. The band report gains `race_sweep`: `guide_race`, `best_race`, `margin`, `margin_pct`, `error`, `beyond_error`, `score_unit`, `iterations`, `seed`, and every race's score and error. The markdown report also gets a race sweep line per band.
- **Standalone mode:** `-race-sweep -bis-dir <dir> -out <dir> -race-iterations N -race-seed S` writes `<spec>.races.json` and refuses to run without `-out`.
- **Refactors:** `requestIterationsFor` replaces the healer quotient in `final_set.go`, and `passSpec` is shared with `passesFor`.
- **Fixtures:** the testbuild fixture gains classes.json and combos.json.

Tests: `race_sweep_test.go` has 7 tests. They cover the legal races from the real 70790 tables, best race and margin, the inside-error case, a refused illegal guide race, the healer quotient, tanks judged on threat, and the standalone mode on one spec, one band and low iterations. The package test passes except `TestPublishedBISMatchesGuides`, which was already red (see below).

## Warrior

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| W1 | **Protection, Horde: Tauren loses to Undead on threat at every band.** The guide picked Tauren for Endurance's 5% health. Tanks are judged on threat, and Undead's Touch of the Grave (20% chance on a warrior's attacks to drain 5% of max health, 15 s internal cooldown; client 1260189/1260198, Blizzard 27 Oct notes) is damage, so it is threat. At 60, Undead leads Tauren by +1.65% / +1.97% (raid, seeds 1 / 50001) and +2.29% / +2.42% (bare). The error is ±0.08 points. Orc is second at 60. | `$R/race/seed{1,50001}/warrior-protection.races.json`; web check below | **CONFIRMED-BUG** | old (guide line from c2cece8b7, guide-accuracy round 3's "measured race picks", which judged survival; no sim round swept races) | S f0b23ad12 |
| W2 | Protection, Alliance: Dwarf and Human beat Night Elf on threat by 0.1 to 0.8%. | same files | Aligned (under 1%, stays) | - | - |
| W3 | Arms and Fury: only Skyborne leads, under 1% at 60 (except Arms band 20 at about 1.2%). | same | Aligned | - | - |

W1 fix: `recommendedRaces: [night-elf, undead]`. The Races paragraph in protection.md now names Undead first and gives Touch of the Grave's threat as the reason, and it keeps Tauren as the bigger health bar for when healers need it. The prose test passes.

The guide lane's open "Protection race TMI on one shared preset" item now reads against an Undead pick.

## Every spec, band and faction where the guide race is beaten beyond error at both seeds

Margins are the best race's lead over the guide race, as a percentage of the guide race's score, plus or minus the combined error in the same units. 336 band entries were swept per seed. Rows not listed have the guide race best, or a lead inside error at one seed or more.

| spec | band | preset | faction | guide | best | seed 1 | seed 50001 | unit | verdict |
|---|---|---|---|---|---|---|---|---|---|
| druid-balance | 20 | bare | horde | tauren | windshaper-skyborne | +1.45% ± 0.08 | +1.45% ± 0.08 | dps | paid race; prose covers it |
| druid-balance | 30 | bare | horde | tauren | windshaper-skyborne | +1.55% ± 0.08 | +1.58% ± 0.08 | dps | paid race; prose covers it |
| druid-balance | 40 | bare | horde | tauren | windshaper-skyborne | +1.07% ± 0.08 | +1.04% ± 0.08 | dps | paid race; prose covers it |
| druid-balance | 50 | bare | horde | tauren | windshaper-skyborne | +0.78% ± 0.08 | +0.92% ± 0.08 | dps | paid race, under 1% |
| druid-balance | 60 | bare | horde | tauren | windshaper-skyborne | +0.75% ± 0.08 | +0.85% ± 0.08 | dps | paid race, under 1% |
| druid-balance | 60 | raid | horde | tauren | windshaper-skyborne | +1.23% ± 0.08 | +1.16% ± 0.08 | dps | paid race; prose covers it |
| druid-feral | 20 | bare | horde | tauren | windshaper-skyborne | +0.66% ± 0.05 | +0.63% ± 0.05 | dps | paid race, under 1% |
| druid-feral | 30 | bare | horde | tauren | windshaper-skyborne | +0.82% ± 0.05 | +0.83% ± 0.05 | dps | paid race, under 1% |
| druid-feral | 40 | bare | horde | tauren | windshaper-skyborne | +0.79% ± 0.06 | +0.96% ± 0.06 | dps | paid race, under 1% |
| druid-feral | 50 | bare | horde | tauren | windshaper-skyborne | +0.80% ± 0.06 | +0.89% ± 0.06 | dps | paid race, under 1% |
| druid-feral | 60 | bare | horde | tauren | windshaper-skyborne | +0.84% ± 0.06 | +0.80% ± 0.06 | dps | paid race, under 1% |
| druid-feral | 60 | raid | horde | tauren | windshaper-skyborne | +1.18% ± 0.05 | +1.10% ± 0.05 | dps | paid race; prose covers it |
| druid-feral-bear | 20 | bare | alliance | night-elf | high-order-skyborne | +1.35% ± 0.09 | +1.23% ± 0.09 | threat_per_second | paid race; prose covers it |
| druid-feral-bear | 20 | bare | horde | tauren | windshaper-skyborne | +0.43% ± 0.09 | +0.42% ± 0.09 | threat_per_second | paid race, under 1% |
| druid-feral-bear | 30 | bare | alliance | night-elf | high-order-skyborne | +1.01% ± 0.07 | +1.01% ± 0.08 | threat_per_second | paid race; prose covers it |
| druid-feral-bear | 30 | bare | horde | tauren | windshaper-skyborne | +0.62% ± 0.08 | +0.53% ± 0.08 | threat_per_second | paid race, under 1% |
| druid-feral-bear | 40 | bare | alliance | night-elf | high-order-skyborne | +0.37% ± 0.08 | +0.46% ± 0.08 | threat_per_second | paid race, under 1% |
| druid-feral-bear | 40 | bare | horde | tauren | windshaper-skyborne | +0.28% ± 0.08 | +0.31% ± 0.08 | threat_per_second | paid race, under 1% |
| druid-feral-bear | 50 | bare | horde | tauren | windshaper-skyborne | +0.72% ± 0.09 | +0.73% ± 0.09 | threat_per_second | paid race, under 1% |
| druid-feral-bear | 60 | bare | horde | tauren | windshaper-skyborne | +0.81% ± 0.09 | +0.81% ± 0.09 | threat_per_second | paid race, under 1% |
| druid-feral-bear | 60 | raid | horde | tauren | windshaper-skyborne | +0.56% ± 0.10 | +0.64% ± 0.10 | threat_per_second | paid race, under 1% |
| druid-restoration | 20 | bare | horde | tauren | windshaper-skyborne | +0.11% ± 0.02 | +0.10% ± 0.02 | heal_score | paid race, under 1% |
| druid-restoration | 30 | bare | alliance | night-elf | high-order-skyborne | +1.31% ± 0.06 | +1.18% ± 0.05 | heal_score | paid race; prose covers it |
| druid-restoration | 30 | bare | horde | tauren | windshaper-skyborne | +0.14% ± 0.06 | +0.06% ± 0.06 | heal_score | paid race, under 1% |
| druid-restoration | 40 | bare | alliance | night-elf | high-order-skyborne | +1.40% ± 0.04 | +1.36% ± 0.04 | heal_score | paid race; prose covers it |
| druid-restoration | 40 | bare | horde | tauren | windshaper-skyborne | +0.19% ± 0.04 | +0.19% ± 0.04 | heal_score | paid race, under 1% |
| druid-restoration | 50 | bare | alliance | night-elf | high-order-skyborne | +2.50% ± 0.06 | +2.49% ± 0.06 | heal_score | paid race; prose covers it |
| druid-restoration | 50 | bare | horde | tauren | windshaper-skyborne | +0.66% ± 0.07 | +0.72% ± 0.07 | heal_score | paid race, under 1% |
| druid-restoration | 60 | bare | alliance | night-elf | high-order-skyborne | +1.76% ± 0.05 | +1.81% ± 0.05 | heal_score | paid race; prose covers it |
| druid-restoration | 60 | raid | alliance | night-elf | high-order-skyborne | +0.70% ± 0.03 | +0.70% ± 0.03 | heal_score | paid race, under 1% |
| hunter-beast-mastery | 20 | bare | horde | orc | windshaper-skyborne | +0.94% ± 0.05 | +0.87% ± 0.05 | dps | paid race, under 1% |
| hunter-beast-mastery | 30 | bare | horde | orc | windshaper-skyborne | +0.30% ± 0.05 | +0.25% ± 0.05 | dps | paid race, under 1% |
| hunter-marksmanship | 20 | bare | horde | troll | windshaper-skyborne | +1.34% ± 0.05 | +1.31% ± 0.05 | dps | paid race; prose covers it |
| hunter-marksmanship | 30 | bare | horde | troll | windshaper-skyborne | +1.08% ± 0.05 | +1.13% ± 0.05 | dps | paid race; prose covers it |
| hunter-marksmanship | 40 | bare | horde | troll | windshaper-skyborne | +1.32% ± 0.06 | +1.33% ± 0.06 | dps | paid race; prose covers it |
| hunter-marksmanship | 50 | bare | horde | troll | windshaper-skyborne | +1.09% ± 0.06 | +1.05% ± 0.06 | dps | paid race; prose covers it |
| hunter-marksmanship | 60 | bare | horde | troll | orc | +0.99% ± 0.06 | +0.96% ± 0.06 | dps | under 1%, stays |
| hunter-survival | 60 | bare | horde | orc | windshaper-skyborne | +0.65% ± 0.08 | +0.53% ± 0.08 | dps | paid race, under 1% |
| hunter-survival | 60 | raid | alliance | night-elf | high-order-skyborne | +0.17% ± 0.08 | +0.14% ± 0.08 | dps | paid race, under 1% |
| hunter-survival | 60 | raid | horde | orc | windshaper-skyborne | +0.14% ± 0.08 | +0.32% ± 0.08 | dps | paid race, under 1% |
| mage-fire | 20 | bare | alliance | human | gnome | +3.28% ± 0.12 | +3.31% ± 0.12 | dps | leveling band only; the line is a level-60 pick |
| mage-fire | 30 | bare | alliance | human | gnome | +2.47% ± 0.13 | +2.56% ± 0.13 | dps | leveling band only; the line is a level-60 pick |
| mage-fire | 40 | bare | alliance | human | gnome | +0.56% ± 0.13 | +0.41% ± 0.13 | dps | under 1%, stays |
| mage-fire | 50 | bare | alliance | human | gnome | +1.86% ± 0.14 | +1.86% ± 0.14 | dps | leveling band only; the line is a level-60 pick |
| mage-frost | 60 | raid | alliance | gnome | high-order-skyborne | +0.57% ± 0.08 | +0.67% ± 0.08 | dps | paid race, under 1% |
| paladin-protection | 30 | bare | alliance | dwarf | human | +1.20% ± 0.14 | +1.05% ± 0.14 | threat_per_second | leveling band only; the line is a level-60 pick |
| paladin-protection | 40 | bare | alliance | dwarf | human | +0.32% ± 0.12 | +0.21% ± 0.13 | threat_per_second | under 1%, stays |
| paladin-protection | 50 | bare | alliance | dwarf | human | +0.31% ± 0.14 | +0.22% ± 0.14 | threat_per_second | under 1%, stays |
| paladin-retribution | 60 | raid | alliance | human | dwarf | +0.46% ± 0.10 | +0.43% ± 0.10 | dps | under 1%, stays |
| priest-holy | 30 | bare | alliance | human | night-elf | +4.81% ± 0.04 | +4.85% ± 0.04 | heal_score | leveling band only; the line is a level-60 pick |
| priest-shadow | 20 | bare | alliance | human | gnome | +0.25% ± 0.07 | +0.29% ± 0.07 | dps | under 1%, stays |
| priest-shadow | 30 | bare | alliance | human | gnome | +1.06% ± 0.07 | +1.13% ± 0.07 | dps | leveling band only; the line is a level-60 pick |
| priest-shadow | 40 | bare | alliance | human | gnome | +3.06% ± 0.11 | +3.05% ± 0.11 | dps | leveling band only; the line is a level-60 pick |
| priest-shadow | 50 | bare | alliance | human | gnome | +2.41% ± 0.11 | +2.48% ± 0.11 | dps | leveling band only; the line is a level-60 pick |
| priest-shadow | 60 | bare | alliance | human | gnome | +0.54% ± 0.10 | +0.55% ± 0.10 | dps | under 1%, stays |
| rogue-assassination | 60 | raid | alliance | night-elf | gnome | +0.43% ± 0.06 | +0.38% ± 0.06 | dps | under 1%, stays |
| rogue-combat | 20 | bare | alliance | gnome | night-elf | +1.34% ± 0.07 | +1.41% ± 0.07 | dps | leveling band only; the line is a level-60 pick |
| rogue-combat | 30 | bare | alliance | gnome | night-elf | +0.83% ± 0.07 | +0.86% ± 0.07 | dps | under 1%, stays |
| rogue-combat | 40 | bare | alliance | gnome | night-elf | +1.01% ± 0.07 | +1.10% ± 0.07 | dps | leveling band only; the line is a level-60 pick |
| rogue-combat | 50 | bare | alliance | gnome | night-elf | +0.55% ± 0.06 | +0.51% ± 0.06 | dps | under 1%, stays |
| rogue-subtlety | 60 | raid | alliance | night-elf | gnome | +0.27% ± 0.05 | +0.35% ± 0.05 | dps | under 1%, stays |
| shaman-elemental | 20 | bare | horde | orc | windshaper-skyborne | +4.92% ± 0.12 | +5.08% ± 0.11 | dps | paid race; prose covers it |
| shaman-elemental | 60 | raid | horde | orc | windshaper-skyborne | +0.46% ± 0.07 | +0.44% ± 0.07 | dps | paid race, under 1% |
| shaman-enhancement | 20 | bare | horde | orc | windshaper-skyborne | +0.62% ± 0.09 | +0.68% ± 0.09 | dps | paid race, under 1% |
| shaman-enhancement | 30 | bare | horde | orc | windshaper-skyborne | +1.23% ± 0.10 | +1.03% ± 0.10 | dps | paid race; prose covers it |
| shaman-enhancement | 40 | bare | horde | orc | windshaper-skyborne | +0.92% ± 0.12 | +0.82% ± 0.12 | dps | paid race, under 1% |
| shaman-enhancement | 50 | bare | horde | orc | windshaper-skyborne | +0.70% ± 0.12 | +0.77% ± 0.12 | dps | paid race, under 1% |
| shaman-enhancement | 60 | bare | horde | orc | windshaper-skyborne | +0.40% ± 0.11 | +0.23% ± 0.11 | dps | paid race, under 1% |
| shaman-restoration | 20 | bare | horde | orc | windshaper-skyborne | +0.10% ± 0.02 | +0.15% ± 0.02 | heal_score | paid race, under 1% |
| shaman-restoration | 60 | raid | horde | orc | windshaper-skyborne | +0.31% ± 0.05 | +0.26% ± 0.05 | heal_score | paid race, under 1% |
| warlock-affliction | 20 | bare | horde | orc | undead | +1.13% ± 0.05 | +1.03% ± 0.05 | dps | leveling band only; the line is a level-60 pick |
| warlock-affliction | 30 | bare | alliance | human | gnome | +0.09% ± 0.05 | +0.10% ± 0.05 | dps | under 1%, stays |
| warlock-affliction | 30 | bare | horde | orc | undead | +0.34% ± 0.05 | +0.41% ± 0.05 | dps | under 1%, stays |
| warlock-demonology | 20 | bare | horde | orc | undead | +1.08% ± 0.06 | +1.04% ± 0.06 | dps | leveling band only; the line is a level-60 pick |
| warlock-demonology | 30 | bare | horde | orc | undead | +0.29% ± 0.05 | +0.26% ± 0.06 | dps | under 1%, stays |
| warlock-destruction | 20 | bare | horde | orc | undead | +0.76% ± 0.05 | +0.80% ± 0.05 | dps | under 1%, stays |
| warlock-destruction | 30 | bare | horde | orc | undead | +0.13% ± 0.06 | +0.11% ± 0.06 | dps | under 1%, stays |
| warrior-arms | 20 | bare | alliance | night-elf | high-order-skyborne | +1.27% ± 0.08 | +1.19% ± 0.08 | dps | paid race; prose covers it |
| warrior-arms | 20 | bare | horde | undead | windshaper-skyborne | +1.14% ± 0.07 | +1.17% ± 0.07 | dps | paid race; prose covers it |
| warrior-arms | 60 | raid | horde | undead | windshaper-skyborne | +0.32% ± 0.10 | +0.21% ± 0.10 | dps | paid race, under 1% |
| warrior-fury | 20 | bare | alliance | night-elf | high-order-skyborne | +0.20% ± 0.09 | +0.18% ± 0.09 | dps | paid race, under 1% |
| warrior-fury | 30 | bare | alliance | night-elf | high-order-skyborne | +0.22% ± 0.08 | +0.15% ± 0.08 | dps | paid race, under 1% |
| warrior-fury | 60 | raid | alliance | night-elf | high-order-skyborne | +0.20% ± 0.08 | +0.12% ± 0.08 | dps | paid race, under 1% |
| warrior-protection | 20 | bare | alliance | night-elf | human | +0.74% ± 0.07 | +0.70% ± 0.07 | threat_per_second | under 1%, stays |
| warrior-protection | 20 | bare | horde | tauren | undead | +3.07% ± 0.07 | +3.17% ± 0.07 | threat_per_second | leveling band only; the line is a level-60 pick |
| warrior-protection | 30 | bare | alliance | night-elf | human | +0.38% ± 0.08 | +0.41% ± 0.08 | threat_per_second | under 1%, stays |
| warrior-protection | 30 | bare | horde | tauren | undead | +3.14% ± 0.08 | +3.04% ± 0.08 | threat_per_second | leveling band only; the line is a level-60 pick |
| warrior-protection | 40 | bare | alliance | night-elf | dwarf | +0.69% ± 0.07 | +0.84% ± 0.07 | threat_per_second | under 1%, stays |
| warrior-protection | 40 | bare | horde | tauren | undead | +2.09% ± 0.07 | +2.06% ± 0.07 | threat_per_second | leveling band only; the line is a level-60 pick |
| warrior-protection | 50 | bare | alliance | night-elf | human | +0.10% ± 0.07 | +0.15% ± 0.07 | threat_per_second | under 1%, stays |
| warrior-protection | 50 | bare | horde | tauren | undead | +2.25% ± 0.07 | +2.28% ± 0.07 | threat_per_second | leveling band only; the line is a level-60 pick |
| warrior-protection | 60 | bare | alliance | night-elf | dwarf | +0.50% ± 0.07 | +0.47% ± 0.07 | threat_per_second | under 1%, stays |
| warrior-protection | 60 | bare | horde | tauren | undead | +2.29% ± 0.07 | +2.42% ± 0.07 | threat_per_second | **CONFIRMED-BUG, fixed** |
| warrior-protection | 60 | raid | alliance | night-elf | dwarf | +0.22% ± 0.08 | +0.25% ± 0.08 | threat_per_second | under 1%, stays |
| warrior-protection | 60 | raid | horde | tauren | undead | +1.65% ± 0.08 | +1.97% ± 0.08 | threat_per_second | **CONFIRMED-BUG, fixed** |

The leveling-only losses over 1% are listed here so a later round can decide whether a guide should name a leveling race:

- **Fire mage, Alliance:** Gnome leads at 20, 30 and 50 by 1.9 to 3.3%. Human is best at 60.
- **Shadow priest, Alliance:** Gnome leads at 30 to 50 by 1.1 to 3.1%, and at bare 60 by 0.55%. Human is best in the raid.
- **Holy priest, Alliance:** Night Elf leads at band 30 by 4.8%. This is likely a priest race spell at that band; I did not check which.
- **Protection paladin, Alliance:** Human leads Dwarf at band 30 by 1.1 to 1.2%.
- **Combat rogue, Alliance:** Night Elf leads at 20 and 40 by 1.0 to 1.4%. The bands tie at 60.
- **Warlock, Horde:** Undead leads Orc at band 20 by 0.8 to 1.1%. Warlock Orc is pending with the owner (round-14 O1-O7), so this is not reopened.
- **Marksmanship, Horde:** a free race, Orc, leads Troll by 0.5 to 1.0% at every bare band. That includes 60 bare (+0.99% / +0.96%) and band 50 (+1.00% / +1.02%), where Skyborne is the overall best. It is under 1% at 60, so Troll stays.

## Web check (Forever race consensus), looked 2026-11-06

- **Icy Veins, "Best Race for Every Class in WoW Forever" (2026-09-17, beta data):** for Warrior it gives "Night Elf (Dwarf for Tank)" on the Alliance and "Orc or Tauren" on the Horde. https://www.icy-veins.com/wow-forever/news/best-race-for-every-class-in-wow-forever-watch-our-breakdown/
- **Odealo, "Best Race for Every Class" (updated 2026-10-09, beta build):** gives Warrior Night Elf / Orc and Priest Gnome / Undead. It names Tauren for survivability. https://odealo.com/articles/best-race-for-every-class-world-of-warcraft-forever-guide
- Both sources predate Blizzard's 27 October cut of Touch of the Grave's internal cooldown to 15 s. Neither measured threat. The strong-player sources split between Orc and Tauren on survival. Under the owner's threat rule the measurement decides: Undead first, with Orc second, as both sources rank Orc. Single-source caveat: no Forever-specific tank race write-up after 27 Oct was found.
- Real logs: `api.foreversixty.gg/v1/reports/recent` at 17:12 UTC on 6 Nov has one report (20 Oct, no kills), so nothing can settle a race.

## Before/after

The race line moves no published DPS until the BiS refresh re-ranks Protection on Undead. On the frozen Horde Protection set, Undead against Tauren in threat per second at seed 1 / seed 50001:

- Raid 60: 1672.1 against 1644.9 (+1.65%) / 1674.6 against 1642.2 (+1.97%).
- Bare 60: 917.6 against 897.1 (+2.29%) / 918.0 against 896.3 (+2.42%).

No other spec's sims change. The ranker's extra sweep sims run only inside the ranker.

## Files to regenerate

None by this lane. The next BiS refresh writes `race_sweep` into every `data/builds/1.60.1.70790/bis/<spec>.json` and ranks Protection warrior Horde on Undead. Tiers and app data that carry the published race follow the refresh.

## Measure at refresh

- The published `race_sweep` fields arrive with the next BiS refresh. Its ranker run sweeps at 3,000 iterations on the verify seed. Confirm that the band reports carry `race_sweep` and that no `beyond_error` row over 1% at 60 has a free race beating a guide race, other than ones this report lists.
- Protection warrior Horde gear re-ranked on Undead. **`TestPublishedBISMatchesGuides` (sim/cmd/leveling-bis) is red for warrior-protection Horde** until that refresh. It was already red on int/r25 for priest-discipline Alliance (published Gnome, guide Human, from round 24).
- The ranker now runs about 4 extra sims per band and faction at 3,000 iterations: the legal races minus the memoized guide race, 2 to 5 per faction. Expect a few percent more ranker time, and more for healers.

## For other lanes

- Shared code: `sim/score/tank.go` `TankFigures.TPSError` (additive; no other caller changes) and the bandReport JSON field `race_sweep` (optional). `web/src/lib/bis/types.ts` can add it when the page wants to show it.
- App test fixtures (apps/ios GuidesTestData.swift :33, GuidesLogicTests.swift :140; apps/android GuidesLogicTest.kt :204) use a Tauren Protection warrior. These are self-contained tier fixtures, not the guide, so I left them.
- Guide-accuracy: the leveling-band race leads listed above, if a guide should name a leveling race.

## For the guide lane

- protection.md :7 and :113 changed (f0b23ad12). Re-read the Races section against main's wording when it lands.

## Still open

- Holy priest band 30, where Night Elf leads by 4.8%: which racial or race spell drives it was not traced.
- Vitest suites that import `web/src/data/generated/*` could not load in this worktree, because no generated data is synced. The prose and race-table and race-pill tests ran and pass. CI covers the rest.
