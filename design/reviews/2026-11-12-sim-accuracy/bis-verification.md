# BiS verification round, 2026-11-12

Lane `bis-verify`, ledger step 3 ("A verification round audits the regenerated output"). Scope: the first full re-rank since the refresh hold lifted, `b84481ef5` "chore(bis): nightly leveling BiS for 1.60.1.70790 @ 4f8bb748d" (58 files under `data/builds/1.60.1.70790/bis/`), against the previous nightly `2de65052d` (engine 0e04b1284, 10 Nov). Part 2 (below) covers the second ranking, GitHub Actions bis run 43269156099 on `c14f5c58a` (engine pin 0fb4dcebb: the Shadow wand windup, the paladin Judgement fixes, the Frost build).

## Method

1. **Offline checks** over every spec's `bis/<spec>.json` and `bis/<spec>.md` (28 specs, 336 band entries, 20,121 published picks and alternatives), with a script that re-derives each fact from the build's own data: `items.json` (class, subclass, inventory type, required level, faction restriction), `items/<class>.json` (the class mask), `sim/leveling/armor_proficiency.go` (mail at 40 for shaman and hunter, plate at 40 for warrior and paladin), the ranker's `classWeaponSubclasses` table and the owner's rules (shamans never dual wield; Survival fights in melee).
2. **The ranker's own guards** read from the files: `kept_alternative` (the previous published set or a pre-swap set re-measured against the new set at every band), `set_dps` monotonic across bands and raid above bare, and the per-slot `alternatives` deltas.
3. **The ranker's published-file tests** on the Windows box: `TestPublishedBISFilesPassSanityChecks` and `TestPublishedBISMatchesGuides` (`sim/cmd/leveling-bis`), engine pinned.
4. **Spot sims** on the box through `tools/remote-test.sh --engine-sha 4f8bb748d --dir sim` (the engine the files were ranked on, so a figure reproduces the published one), `bis-verification/jobs.sh`, one `spec-breakdown` at a time, `nice -n 15`, GOMAXPROCS=5, the ranker's own 3,000 iterations and seed 7: the previous band's published set simmed at the next band (`prevband.txt`, 224 jobs), and the published set with one sampled slot emptied (`emptied.txt`, 691 jobs: every trinket, ranged, relic and off-hand pick the ranker decided by sim or that had no competing alternative).
5. **Measure at refresh** rows MR47 to MR71: the new picks, talents, weights and metrics against the lane's expected move, old file against new.

## Part 1: b84481ef5

**Verdict: every check passes; no ranker bug confirmed; nothing to re-rank.** Slot legality Pass, weapon picks Pass, no set below its own check Pass (the ranker's guards, 200 previous-band and 578 emptied-slot sims plus the tanks on the full score), tier order Pass, the published-file tests Pass. MR47 to MR71: 17 rows confirmed in the files, 5 not visible in the bis files (rank rows and the kit live in the addon data), 3 not applicable to a re-rank (the 70888 move, Poised and Pounce need logs). The four `knownStaleBISFiles` entries this re-rank landed are dropped (Protection warrior, Arms, Beast Mastery, Arcane); mage-frost stays until part 2. The picks are ready for an addon tag.

### Slot legality: Pass

Every published pick and every listed alternative (20,121) is in its class's equippable set, sits in a slot its inventory type fills, carries no armor subclass before the level the class trains it (no mail on a hunter or shaman and no plate on a warrior or paladin below band 40), no weapon subclass the class cannot train, no faction-restricted item on the other side, and no required level above the band. No finger or trinket pair wears the same item twice.

### Weapon picks: Pass

Every band has a main hand. No two-hander is published beside an off hand. No shaman entry wears a weapon in the off hand (shields and held items only). Every Survival main hand is a melee weapon and its weights are the melee table (the sanity test's own check). The dual-wield off hands are on warriors, rogues and hunters only. Ranged slots hold bows, guns, crossbows, thrown, wands or relics the class carries.

### No set scores below its own check

- **Across bands (ranker check): Pass.** `set_dps` rises band to band on every faction for all 28 specs, and raid 60 is above bare 60 everywhere.
- **The previous published set (ranker guard): Pass, 33 holds.** At 33 band entries the ranker's final set measured below a set it had already considered and the guard kept the better one (`kept_alternative` in the json). The holds worth a look, because the new weights-ranked set was clearly worse than the incumbent: Elemental bare 40 (new -4.4% against the kept set, both factions), bear bare 50 (-3.9 / -4.4%), Marksmanship bare 60 (-1.6 / -1.7%), Protection paladin bare 60 Horde (-2.0%), Arms raid 60 Alliance (the pre-swap set kept, -2.0%). None is a published defect (the published set is the better one, re-measured); they are the weights stage under-ranking the incumbent, which the guard exists for.
- **The previous band's set re-simmed at the new level (spot sims):** see the spot-sim section.
- **The set with one slot emptied (spot sims):** see the spot-sim section.

### Spot sims: the previous band's set and the emptied slot

Run on the box at b85c3e3a2, engine 4f8bb748d, 915 sims in 395 s (`spot.log` in the lane's scratch, the job lines in `bis-verification/prevband.txt` and `emptied.txt`).

- **The previous band's published set at the next band: Pass for every DPS and healer entry** (200 of 200). At every bare band 30 to 60, both factions, the band below's set (with the slots it lacks emptied) scores under the band's own published figure by more than the two runs' combined error. The tank entries (24) are compared on threat and chance of death, since `spec-breakdown` prints the figures and not the score: 22 pass; Protection warrior bare 40 on both factions reads more threat a second on band 30's set (463 against 440 Alliance, 474 against 452 Horde, death 0 both ways). The tank score is effective health times the quarter power of threat times the TMI guard, so a lighter set with more rage can out-threat the heavier set and still score lower; the full figures are in the tank section below.
- **The published set with one slot emptied: Pass for every DPS and healer entry** (578 of 578 sampled trinket, ranged, relic and off-hand picks: none scores above its own set when removed). The 18 hunter `ranged` jobs sim to NaN (a hunter with no ranged weapon has no fight) and are not a finding. The tank entries (95) pass on threat and death except four Protection paladin trinket jobs (bare 40 and 50 Alliance, Rune of Duty and the Talismans of Arathor): the set without the trinket reads 275.4 and 398.5 threat a second, which is the published set's own `spec-breakdown` baseline (275.5, 398.x), so the trinkets add no threat and are worn for the survival terms; see the tank section.

### Tank entries on the full score

The six tank spot sims whose threat alone rose were re-run with the full figures (`tankfull.txt`) and scored with `sim/score/tank.go`'s formula (effective health times the quarter power of threat over 100 times exp(-0.005 TMI); the boss's raw damage cancels in the ratio):

| entry | set | health | damage taken /s | TMI | threat /s | score, relative |
|---|---|---|---|---|---|---|
| Protection warrior bare 40 Alliance | published | 2,919 | 93.0 | 20.6 | 440 | 1.00 |
| | band 30's set | 2,619 | 129.3 | 33.4 | 463 | 0.61 |
| Protection warrior bare 40 Horde | published | 2,909 | 91.8 | 20.4 | 451 | 1.00 |
| | band 30's set | 2,639 | 131.0 | 33.5 | 474 | 0.61 |
| Protection paladin bare 40 Alliance | published | 2,881 | 123.8 | 29.1 | 275.5 | 1.00 |
| | without a trinket | 2,801 | 123.7 | 30.0 | 275.4 | 0.97 |
| Protection paladin bare 50 Alliance | published | 3,936 | 220.1 | 40.6 | 398.8 | 1.00 |
| | without a trinket | 3,836 | 219.8 | 41.8 | 398.5 | 0.97 |

Every published tank set scores above the lighter set: the previous band's set and the emptied-slot checks pass for the tanks too. Both Protection paladin trinkets at 40 and 50 (Rune of Duty, Talisman of Arathor, Volunteer's Lucky Seal) are stamina pieces worth 3% of score each, nothing to threat.

### Tier order: Pass

Within every band the ranked alternatives per slot descend in DPS delta, in the json and in the md tables (5,712 slots), and no alternative scores above its pick (ties at +0.00 are the sim-decided picks' equals).

### The ranker's published-file tests: Pass

On the box at b85c3e3a2 (engine 4f8bb748d): `TestPublishedBISMatchesGuides` passes for 27 specs and skips mage-frost (its build moved at 0f7c39cc8, 15:35 UTC, after the b84481ef5 ranking had started; part 2 lands it). `TestPublishedBISFilesPassSanityChecks` (gated on `FOREVER_BIS_SANITY=1`, as the nightly runs it): all 28 active-build subtests pass (eligibility, source, faction, weapon shape, the armor gate, no duplicated pair, the weight lists, Survival's melee weights, set_dps monotonic). The same test fails for the two older builds still carrying bis files (1.60.1.70753, 1.60.1.70755: their PvP source labels predate MR2's in-game rank labels and the bear's weight list predates feral_attack_power); CI runs that step warning-only, and those files are not what the site reads. Not a finding against this re-rank; whether the old builds' bis directories should be pruned is a housekeeping question for the integrator.

### Calibration of the spot sims

`spec-breakdown` at 3,000 iterations and seed 7 on engine 4f8bb748d reproduces the published figures to the hundredth: Arms bare 30 Alliance 65.66 (published 65.66), Holy priest bare 40 Alliance 175.91 HPS (published 175.91). A tank entry is close but not identical: Protection paladin bare 40 Alliance threat 275.5 a second against the published 272.1 (the tank pipeline's final-set run is a separate request), so tank spot sims are judged on a 1% threat margin with no worse chance of death. `-gear head:0` empties a slot (Arms bare 30 without its head: 63.82).

### Measure at refresh, MR47 to MR71

| row | expected | in b84481ef5 | verdict |
|---|---|---|---|
| MR47 | tanks 20 to 50 on the dungeon-table boss: picks, weights, threat, death, TMI | published `metrics` moved on every tank band 20 to 50: damage taken a second roughly halved (warrior 62 to 39 at 20, 184 to 93 at 40; paladin 232 to 124 at 40; bear 272 to 149 at 40), TMI down by a third to a half, chance of death 0 on every band 20 to 50 (bear 30: 3.8%, the lane's 2 to 4%); 45 Protection paladin slots, 30 Protection warrior slots and 4 bear slots re-picked | confirmed |
| MR48 | Protection warrior band 30 talents; the bear's Maul gate | warrior-protection band 30 talents moved on both factions (`T` in the diff); the bear's APL is not in the file, its band figures moved (below) | confirmed (talents); the gate shows only in the numbers |
| MR49 | round 30 rank rows | rank rows live in `addon-data.json`, not the bis files (regenerated at 2f3faa827); no bis field carries them | not visible here |
| MR50 | tanks on the 48 s pull: published metrics | the same `metrics` move as MR47 (the two lanes compound in this one re-rank); paladin threat a second +38 to +42% at 20 to 50, warrior -5 to -15%, bear -1 to -9% | confirmed |
| MR51 | bear Maul gate 25 to 29, no band entry moves | no bear band entry moved for it (bear 20 is 22.8 to 22.2 tank score on the boss change) | confirmed |
| MR52 | Retribution +5 to +6% bare, +2% raid; Protection threat +1 to +7% | Retribution bare +5.1 to +6.3% on every band and faction, raid 60 +1.9 / +3.5%; Protection threat is inside the tank lanes' larger move. The ranker's Protection paladin weight set carries no spell power or Intellect (`data/curated/specs.json`), so the fork's 0.304 to 0.346 spell power move cannot show in the bis weights; spell power pieces reach a Protection set only through the sim-decided tournaments | confirmed (Retribution); Protection weights: open, see below |
| MR53 | Protection band 40 on the top Holy Strike rank | the published entry carries no casts; the addon data carries the row | not visible here |
| MR54 | Balance 50 and 60 after the Insect Swarm row | rows are in the addon data; Balance bare 50 Horde moved -1.5% on three re-picked slots (feet and trinkets), 60 unchanged | not visible here |
| MR55 | Enhancement band 20: Rockbiter against Flametongue in the kit | the imbue is the leveling kit's (`sim/leveling/kit.go`), not a bis field; Enhancement bare 20 moved -3.9 / -4.1% (the 11 Nov hotfix lane's -5.2% at 20, net of the kit change) | not visible here |
| MR56 | Idol of Synthesis at druid-restoration 60 | the ranged slot was empty in 2de65052d and is Idol of Synthesis at bare 60 and raid 60, both factions | confirmed |
| MR57 | Tidal Totem at shaman-restoration 60 | the ranged slot was empty and is Tidal Totem at bare 60 and raid 60, both factions; raid 60 +2.8 / +3.1%, bare 60 +2.5% | confirmed |
| MR58 | the 70888 client move | the site's data is still 70790 | not applicable to this re-rank |
| MR59 | Swiftmend's source, raid 60 +0.05% | resto druid raid 60 moved +0.8% (1019.7 to 1028.3 Alliance) with the idol and three re-picked slots; +0.05% is under the entry's own error | confirmed in direction |
| MR60 | Poised on a 70888 log | needs a log, not a re-rank | not applicable |
| MR61 | Arms talents from 17; bare 60 +1.3 to +1.7% | talents moved on every bare band 20 to 60, both factions; bare 60 +3.3 / +3.1% with six re-picked slots, bare 30 +4.1 / +3.8%, bare 40 +2.3 / +2.5% | confirmed; stale entry dropped |
| MR62 | Beast Mastery talents from 46; +1.2 to +1.6% | talents moved at bare 50 and 60, both factions; the bands fell -5.2% (30), -4.9 / -5.8% (40), -3.2% (50), -0.9 / -0.3% (60) because the 11 Nov hotfix lane's Arcane Shot and Summon Hawk mana (-2.1 / -5.9 / -7.3 / -6.5 / -5.4% at 20 to 60, MR44) lands in the same re-rank; the build's gain is the gap between that table and these | confirmed; stale entry dropped |
| MR63 | Survival bare 40 -2.2%, 50 -1.5% | bare 40 -2.0 / -2.6%, bare 50 -1.6 / -1.8% | confirmed |
| MR64 | Marksmanship sting rows | rows are in the addon data; Marksmanship bands moved -1.3 to -3.2% (the hotfix lane's MM column, MR44) | not visible here |
| MR65 | rage refund: Arms raid -0.4%, Fury -0.1 to -0.5% | Arms raid 60 +0.7 / -0.4% (Alliance kept the pre-swap set, four slots differ), Fury raid 60 -0.1 / -0.1%, Fury bare 60 -0.2% Alliance (Horde +1.2% on four re-picked slots) | confirmed |
| MR66 | Demonology ladder -0.4 to +0.3%; raid unchanged | raid 60 +1.0 / +0.5% on eight re-picked slots (the Imp's Firebolt cadence and crit from the same lane); bare 60 +2.7 / +2.5% | consistent; the lane's "unchanged" was the raid number on fixed gear |
| MR67 | mage walks at 23 to 27; no band entry moves | no band entry moved for it; Arcane bare 20 -12.7% is the hotfix lane's Arcane Concentration split (-13.1% at 20, MR45) | confirmed |
| MR68 | Arcane talents from 44; bare 60 +0.4 to +1.0% | talents moved at bare 50 and 60, both factions; bare 60 -1.9 / -2.4% because the hotfix lane's -3.0% at 60 lands in the same re-rank (bare 50 -0.8 / -0.4% against its -3.1%) | confirmed; stale entry dropped |
| MR69 | cat on Rake: +19 to +20% bare, +10.6% raid; attack power reaches Rake | bare 30 to 60 +19.1 to +20.7% on both factions, raid 60 +14.5 / +14.7%; bare 60 weights: attack power 0.388 to 0.446 and Strength 0.854 to 0.982 of an Agility point | confirmed |
| MR70 | bear threat +10 to +12% at 25 to 35, +7 to +9% at 40 to 60, +5.4% raid | bear bare 60 +6.4 / +7.7% tank score (threat a second 806 to 858), raid 60 +4.8 / +4.9% (1,740 to 1,820); below 60 the tank lanes' smaller pull and boss (-2 to -9% threat) land in the same re-rank: bare 40 -6.4 / -8.6%, bare 50 -0.7 / -2.7%, bare 30 +1.1 / -1.5% | confirmed at 60 and raid; compounded below 60 |
| MR71 | Pounce on a log | needs a log | not applicable |

### Other moves in the diff, explained

- Enhancement -16.5 to -18% at bare 30 to 60 and -7.0 / -7.6% raid 60: the 11 Nov hotfix lane (MR42: Stormstrike, Mental Dexterity, Maelstrom Weapon; -19.2 to -19.9% bare, -9% raid on fixed gear). Affliction -9 to -10.6% from 40: the 11 Nov warlock lane (Wrack -30%; -11% raid on fixed gear). Protection warrior -13 to -18% at bare 40 and 50: the tank lanes' boss and pull (less rage from damage taken) plus the rage refund.
- Every caster's band 40 feet changed and the band fell 1.1 to 4.7% (Arcane, Fire, Frost, Balance, Shadow, the three warlocks): the 11 Nov data's Earthen Silk Slippers (MR74: 6 spell damage and 7 Spirit where the 10 Nov rank read 24 spell damage). 49 entries wore them in 2de65052d; two do now (Elemental bare 40, a sim-decided tie with Furen's Boots and the Skycaller boots on a band whose weights put Spirit at 1.55 of an Intellect point and mp5 at 9.2), and Shadow raid 60 wears Omnicast Boots on both factions. MR74 is therefore already confirmed in this re-rank; part 2 re-checks it.
- Elemental bare 40 -2.8% with no slot change: the same slippers, kept, on the corrected stats.

### Open for the owning lanes (not ranker bugs)

- **Protection paladin weight set.** `paladin-protection` weighs stamina, armor, defense, avoidance, block, Strength, Agility, attack power, hit, crit and expertise, never spell power or Intellect, while Holy Strike, Holy Shield and Consecration scale with spell power (MR52). The guard and the tournaments can still seat a spell power piece, but the weights stage cannot rank one. For the paladins lane: whether spell power belongs in the Protection weight set.
- **Protection paladin bare 60 chance of death 34%** (28% before). Not new, published, for the tanks lane.

## Part 2: run 43269156099 (c14f5c58a, engine 0fb4dcebb)

Pending: the run was still in progress when part 1 pushed. The checks above repeat for priest-shadow, the three mages, paladin-retribution and paladin-protection, plus MR72 to MR80, once it lands on main.
