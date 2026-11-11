# Sim accuracy 2026-11-11: the three tanks from 20 to 45

Lane `tanks-below-45`, brief `design/reviews/sim-accuracy/brief.md`. Scope: Protection warrior, Protection paladin and Feral bear between 20 and 45, against real logs. Round 30 had left two items open for the tanks (round30-open.md T1b): a bare bear at 30 to 35 died in 93 to 97% of fights, and a paladin at 30 to 39 in up to 49%.

- Site branch from origin/main 5b4dc79a1.
- Fork branch `lane/tanks-below-45` from master 2b10bae6c, the pin at the start.
- Client 1.60.1.70790.

**Totals: 2 confirmed bugs, both fixed. Engine pin 2b10bae6c to e586233d7.**

1. The tank fight below 60 lasted 180 seconds against a dungeon band's hardest pull. Every logged dungeon fight lasts about half a minute.
2. The bear's Maul gate at 25 to 29 had been tuned on that 180 second fight.

Two items are held, and nothing else in the fight is wrong by the logs.

## Method

1. **Logs.**
   - foreverlogs.gg has 23 public reports since launch. The 20 that hold a dungeon above Ragefire Chasm were read: every boss kill's timeline, combatant info (talent points, gear, stats) and summary, plus a seeded sample of 45 trash fights from each of the five reports whose tank is 30 or more (RFK, Gnomeregan, City of Dalaran, Excavation Site, Zul'Farrak).
   - The group's tank is the warrior, paladin or druid who takes the most creature melee. Its level is its tree points plus 9.
   - This gives 156 tank boss kills at 20 to 43 (61 paladin, 56 warrior, 39 bear) and 167 trash fights.
   - Our API holds one new public report with fights (8 Nov). It is a 17-player open-world log with no dungeon tank, so it is not used.
2. **Sims.**
   - Every sim ran on the Windows box through `tools/remote-test.sh`, one batch at a time, one sim at a time, under `nice -n 15` with `GOMAXPROCS=6`.
   - Bare band gear (the committed BiS files) and the guide's leveling walk (`spec-breakdown -tank -level N`).
   - Tank fight, 5,000 iterations for the before/after tables, 10,000 for every adoption call. Both factions' default races.
   - Seeds 1 and 50001 for every adoption call, plus 100001 and 150001 where a cell was within the error.
3. **Judging.** Changes are judged by the adoption rule. A tank is judged on threat with the round 29 tank rule: chance of death may rise by under 5 points, absolute, at every level a change touches.

## The logs against the sim

Logged tanks on boss kills. Each cell gives the lower quartile / median / upper quartile. Damage taken per second is everything creatures did to the tank. Healing is everything the tank received.

| class | band | levels | kills | fight (s) | damage taken /s | healing /s | healing over damage (median) | tank deaths |
|---|---|---|---|---|---|---|---|---|
| bear | 20 | 20-27 | 27 | 25 / 36 / 41 | 18 / 21 / 35 | 10 / 18 / 28 | 0.73 | 0 |
| bear | 30 | 31 | 6 | 17 / 30 / 37 | 29 / 43 / 130 | 36 / 50 / 94 | 0.95 | 1 |
| bear | 40 | 43 | 6 | 14 / 26 / 39 | 58 / 66 / 129 | 16 / 50 / 142 | 1.07 | 0 |
| paladin | 20 | 21-29 | 48 | 23 / 32 / 41 | 28 / 39 / 57 | 21 / 35 / 48 | 0.87 | 0 |
| paladin | 30 | 30-34 | 13 | 32 / 38 / 50 | 43 / 58 / 81 | 44 / 70 / 73 | 0.93 | 0 |
| warrior | 20 | 20-28 | 45 | 23 / 31 / 48 | 16 / 27 / 38 | 10 / 16 / 29 | 0.70 | 0 |
| warrior | 30 | 32-34 | 11 | 16 / 29 / 36 | 21 / 36 / 42 | 0 / 21 / 33 | 0.69 | 0 |

The trash sample shows the same picture:

| tank | trash fights | fight (s) | damage taken /s | deaths |
|---|---|---|---|---|
| bear at 31 (RFK) | 28 | 24 / 44 / 68 | 28 / 35 / 46 | 0 |
| bear at 43 (Zul'Farrak) | 32 | 26 / 33 / 57 | 50 / 76 / 104 | 0 |
| warrior at 34 (Gnomeregan) | 48 | 18 / 28 / 39 | 20 / 29 / 38 | 0 |
| paladin at 34 (Dalaran) | 23 | 20 / 30 / 56 | 12 / 18 / 37 | 2 |

Over all 15 reports whose tank is 20 or more, the fight lengths were:

| fights | count | median (s) | upper quartile (s) |
|---|---|---|---|
| boss kills | 182 | 31.0 | 40.3 |
| trash fights of 10 s or more | 1,141 | 29.3 | 49.5 |
| all | 1,323 | 29.5 | 47.6 |

Real tanks at 20 to 43 die in about 1 boss kill in 150 and in a few of the sampled trash fights.

| # | finding | evidence | verdict |
|---|---|---|---|
| L1 | **Fight length.** The sim's tank fight below 60 lasted the ranker's 180 seconds. No logged dungeon fight comes near that: the upper quartile is 48 s. | the tables above | CONFIRMED-BUG, fixed (T1) |
| L2 | **Damage taken.** On the new fight, the sim's bare tanks take about the upper quartile of logged boss damage. At 30 to 39: bear 106 to 112 a second (logged q3 130), paladin 90 to 95 (81), warrior 66 to 70 (42). At 20 to 29: bear 55 to 64 (35), paladin 57 to 61 (57), warrior 39 to 42 (38). At 40 to 45: bear 135 to 151 (129; Zul'Farrak trash q3 104). The band damage is meant as a 75th-percentile pull (the heal profile's table), so this matches. The warrior sits above the logged quartile because the sim boss is level +3 and hits harder per swing (one boss for a whole pull). | sim `ba-after48` | Aligned (design) |
| L3 | **Healing.** Sim healing over damage taken is 0.73 to 0.95 below 60. Logged medians are 0.69 to 1.07. | sim profile, logs | Aligned |
| L4 | **Gear.** The band gear is not thin. Logged tanks carry less Stamina than the band sets at the same levels: bear 31 has 118 (sim at 32: 180), bear 43 has 214 (sim at 42: 238), paladin 29 to 30 has 151 to 156 (sim at 32: 196), warrior 32 to 34 has 146 to 165 (sim at 32: 207). Logged warriors and paladins wear more Strength. | combatant info against `-tank` stats | Aligned (no change) |
| L5 | **Avoidance table.** Logged bears never parry (as in the engine), and dodge and miss on 4 to 14% of swings. Logged warriors block 27 to 34% of swings against the sim's 73% at 32. The log marks only full blocks as block events, and a partial block arrives as a damage event, so the two numbers are not comparable. | timelines | Not comparable |
| L6 | **Casts.** Warrior per minute on bosses at 30 to 34, logs against sim on the new fight: Sunder 15 against 23, Shield Block 9 against 11, Revenge 8 against 10, Heroic Strike 4 against 5, Thunder Clap 3.5 against 4.6. That is the same shape at a player's lower APM. Paladin: Consecration 3.4 a minute in logs against 0.19 on the old 180 s fight, where Holy Shield's mana gate starved it, and 4.3 on the new fight. Bear: logged bears at 31 Swipe adds on most bosses (15 a minute). The bear at 43 casts Lacerate 15 and Maul 7 a minute against the sim's 35 and 9 single target. That is a human below the global cooldown cap, not a cost error: Lacerate is 15 rage less Ferocity, as audited before. | casts | Aligned (the paladin aligned by T1) |

## T1: the tank fight below 60 is a dungeon pull (CONFIRMED-BUG, fixed)

**Finding.** Below 60 the tank fought its band's hardest pull: the heal profile's 75th-percentile pull, scaled from the level-60 stand-in. It fought it for the ranker's default 180 seconds plus or minus 20%, a raid boss's length. The band healing (round 30 T1) is a little below the band damage on purpose: a dungeon healer covers about 0.7 to 0.95 of what the tank takes (L3), and the tank's health covers the rest. Over 180 seconds that gap bleeds out any tank whose damage taken sits above its healing. A bear at 30 to 35 takes about 102 a second against 80 healing until Frenzied Regeneration at 36, so it died in 96% of fights. A real bear clears Gnomeregan and RFK, and the logs (L1) show why: no pull lasts three minutes.

**Fix.** `data/curated/tank-encounter.json` gets `dungeon_fight.duration_sec`, 48 s, the upper quartile of the 1,323 logged fights. The upper quartile goes with the band's damage, which is its 75th-percentile pull.

- Below the reference level, `applyTankFight` sets the encounter to that length, varied by the request's own share.
- Level 60 keeps the request's raid length, so no raid figure moves.
- `request.FightSeconds` names the fight the engine ran. spec-breakdown divides per-fight totals by it, not by the request's 180.
- The sim card (`role_metrics`) for a tank below 60 now runs the pull too, as the heal profile's fixed fight does for a healer.

Red first: `TestTankFightBelowTheReferenceLevelIsADungeonPull` and `TestFightSecondsIsThePullForATankBelowTheReferenceLevel` (S ac22b04ae red, b660c0c48 and 45da4b9e9 green). The bear's Primal Bite test now counts casts against the casts its cooldown allows in the fight the engine ran.

**Before/after**, bare, the guide's walk, 5,000 iterations, seed 1. Chance of death in percent, and the change in threat:

| level | bear H / A death | bear threat H / A | paladin H / A death | paladin threat H / A | warrior threat H / A |
|---|---|---|---|---|---|
| 20 | 0.2 / 0.1 to 0 / 0 | +8.1 / +9.9% | 3.2 / 2.5 to 0 / 0 | +20.8 / +22.4% | -1.3 / +0.1% |
| 25 | 13.5 / 11.0 to 0 / 0 | -1.8 / 0.0% | 10.0 / 9.4 to 0 / 0 | +22.4 / +23.6% | -3.1 / -1.8% |
| 29 | 26.3 / 23.5 to 0 / 0 | -1.3 / +1.4% | 9.5 / 8.1 to 0 / 0 | +24.1 / +25.3% | -2.6 / -1.0% |
| 30 | 95.9 / 95.2 to 2.1 / 3.7 | -4.2 / -1.8% | 47.7 / 47.6 to 0.1 / 0 | +25.5 / +29.1% | +0.4 / +1.1% |
| 32 | 96.7 / 95.9 to 1.9 / 2.9 | -4.2 / -2.1% | 20.6 / 19.8 to 0 / 0 | +23.8 / +27.0% | -0.1 / +0.6% |
| 35 | 97.8 / 97.6 to 2.0 / 2.9 | -4.8 / -2.4% | 7.0 / 7.8 to 0 / 0 | +24.2 / +26.9% | -1.2 / -0.3% |
| 36 | 67.6 / 69.4 to 0 / 0 | -5.2 / -3.2% | 7.2 / 7.0 to 0 / 0 | +24.0 / +26.0% | -1.2 / -0.3% |
| 39 | 58.8 / 61.6 to 0 / 0 | -4.8 / -2.6% | 6.4 / 5.9 to 0 / 0 | +21.9 / +24.4% | -0.7 / 0.0% |
| 40 | 14.4 / 14.4 to 0 / 0 | -4.9 / -2.8% | 0.4 / 0.6 to 0 / 0 | +30.0 / +31.2% | -4.7 / -3.3% |
| 45 | 0.1 / 0.2 to 0 / 0 | -4.5 / -2.1% | 1.0 / 1.2 to 0 / 0 | +31.0 / +32.1% | -4.4 / -3.1% |
| 50 | 0 / 0 to 0 / 0 | -5.5 / -3.3% | 2.8 / 2.1 to 0 / 0 | +33.1 / +36.2% | -5.6 / -4.2% |
| 59 | 0 / 0 to 0 / 0 | -5.4 / -3.2% | 3.1 / 3.2 to 0.1 / 0 | +30.5 / +32.9% | -5.9 / -4.2% |
| bare and raid 60 | unchanged | 0 | unchanged | 0 | 0 |

- The warrior dies in no fight before or after.
- The paladin gains threat because it no longer rations mana over three minutes: Consecration comes back (L6).
- The bear and the warrior lose 2 to 6% because a pull is shorter than the rage they build in three minutes.
- Damage taken per second rises 3 to 10%, because the boss spends less of the fight under Demoralizing Shout or Roar.
- Durations from 45 to 180 s at 30 to 42 (`dur` batch) show the old deaths were time, not spikes. At 60 s a bear at 30 to 35 dies in 8% of fights, at 90 s in about 40%, and at 120 s in about 70%.

## B2: the bear's Maul gate on the dungeon pull (CONFIRMED-BUG, fixed)

**Finding.** Round 30 tuned the gate by level on the 180 second fight. On the pull, every gate from 15 to 55 was swept at 20 to 59 (16 levels, then every level 24 to 36, 39 and 46 to 59), 10,000 iterations, seeds 1 and 50001, both factions.

- **25 to 29: 35 rage beats 20 beyond the error in every cell.** Mean/least threat gain:

  | level | gain |
  |---|---|
  | 25 | +1.44 / +1.36% |
  | 26 | +0.94 / +0.74% |
  | 27 | +1.11 / +0.95% |
  | 28 | +1.15 / +0.91% |
  | 29 | +1.34 / +1.18% |

  45 also wins, by less from 26. Chance of death is 0 either way.
- **30 to 35 and 50 to 59 (now 45): 35 gains +0.1 to +0.7% everywhere.** It is within the error in one cell at 33 and one at 56. Four more seeds at each confirm the gain is positive but small (+0.10 to +0.64%), so the rows hold.
- **36 to 39 (25):** 30 and 35 tie 25 within the error, and 45 loses. Death is 0 at every gate now, so round 30's cap no longer binds, but nothing beats 25.
- **40 to 49 (45):** 35 ties.
- **Below 25 (15):** every higher gate loses.

**Fix.** Maul at 35 rage from 25 to 29. Red first: `TestBearMaulGateFollowsTheLevel` (S d93bf9495 red), green with the APL change (S 2beef4b97). The fork carries the synced rotation (F e586233d7) and the pin moves to it. The guide's Rage paragraph and the Maul card note give the new range. The wow-player SHIPped them after one optional wording change, which was taken.

## Held

| # | item | why |
|---|---|---|
| P1 | **Templar's Bulwark first at 20 (owner ruling, round 14).** Owner ruling: the round 14 reservation put Bulwark ahead of the second Sacred Duty point because "chance of death falls from about 77% to 66%". On the pull a paladin at 17 to 27 dies in 0% of fights either way, and the plain walk gains +1.53% threat at 20 (least +1.43, error 0.20) and +0.81% at 21; 22 ties. | The rotation casts Bulwark, and the adoption rule refuses an order that delays a cast talent (round 30 A1 reading). The ruling's premise no longer holds on the sim, so it is the owner's call. |
| B3 | **Bear Maul 35 rage at 30 to 35 and 50 to 59.** | It gains +0.1 to +0.7% everywhere but is within the error at 33 and 56 (B2). |

## Other tank-rule steps re-checked

- **Warrior.** Round 30 already put the threat talents first. Death is 0 at every level before and after, so nothing changes.
- **Bear.** Round 29 rejected three steps. Naturalist before Natural Reaction failed on death, but round 30 found it loses threat at 43 to 49, so it fails the adoption rule anyway. Heart of the Wild and Naturalist orders lose threat at 43 to 51 (round 30). Feral Swiftness 2 after 40 delays Primal Bite. None is held by death alone, so none moves.
- **Paladin.** No step of the paladin's order besides P1 was taken for survival.

## Measure at refresh

| what | specs / bands | why |
|---|---|---|
| every tank figure below 60 on the dungeon pull: picks, tank weights, threat, chance of death, TMI, the published metrics | warrior-protection, paladin-protection, druid-feral-bear, bands 20 to 50 | T1 (fight length 180 to 48 s below 60) |
| the bear's Maul gate at 25 to 29 | druid-feral-bear band 20 (its levels 25 to 29 are walk levels; the band entry at 20 is unchanged) | B2 |

## Gates

On the box, one run at a time:
- **sim:** `go vet ./...`; `go test` for request, leveling, score, rankchoices, cmd/leveling-bis, cmd/addon-builds and cmd/spec-breakdown.
- **Ladder goldens:** the three tanks' goldens were regenerated at pin e586233d7 (S 9783c14dc). Every per-fight count below 60 is now a 48 s fight's.
- **web:** the guide prose, sections and builds tests; the sim version and fixture tests. `astro check`.
- **apl-check** against the fork branch.

## Shas

- **Site:** ac22b04ae (red), b660c0c48, 45da4b9e9 (T1); d93bf9495 (red), 2beef4b97 (B2); a8f879ed2 (pin); 9783c14dc (goldens); this report.
- **Fork:** `lane/tanks-below-45` e586233d7 (the bear rotation sync). Master is fast-forwarded to it.
- **Engine pin after:** e586233d7.
