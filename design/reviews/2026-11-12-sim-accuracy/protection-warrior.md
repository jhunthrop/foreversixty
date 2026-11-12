# Sim accuracy 2026-11-12: Protection warrior

Lane `protection-warrior`, brief `design/reviews/sim-accuracy/brief.md`. Scope: Protection warrior, every band 20 to 60 and the raid 60 preset, end to end: the engine against the client tables under `data/builds/1.60.1.70790` (plus the push layer and the live 70888 client's rows), Blizzard's Forever notes to 12 November, and the public Forever logs; the build, the rotation and the rank rows under the adoption rule and the tank death-chance rule. Arms and Fury only where a shared warrior rule moves them.

- Site branch `lane/sim-protwar` from origin/main aafe382b5.
- Fork branch `lane/sim-protwar` from master 0fb4dcebb, the pin at the start.
- Client on the site: 1.60.1.70790. Live client: 1.60.1.70888 (wago.tools, hotfix-layered rows fetched for every spell and talent the spec casts or takes).

**Totals: 3 confirmed bugs, all old, all fixed. Engine pin 0fb4dcebb to 837cca30a.**

1. Rage from damage taken was Kalgan's formula on the swing after the tank's own armor. Forever computes it on the swing before armor, at an armor Blizzard assumes for the level (20% to 40%), so a tank in plate was short of rage and a point of armor cost rage it does not cost in the game. Shared core (`core/rage.go`): the bear moves too.
2. Bloodrage charged no health. The client's row costs a fifth of base health a cast (338 at 60), and the cast refuses below it.
3. Thunder Clap's damage was the rank's flat amount. Blizzard's 30 September deep dive gives it "significant damage scaling from Attack Power" and eleven logged Protection warriors show about a sixteenth of attack power a hit.

Protection warrior threat +2.2 to +5.5% at every band, chance of death 0 before and after, damage taken unchanged. Nothing on the card moved: the Heroic Strike gate at 40 rage holds against 30, 50 and 60 on the fixed engine. The walk's first points move: Shield Specialization before Improved Bloodrage, +3 to +15% threat at 11 to 18 with chance of death 0 either way.

## Method

1. **Client.** Every spell, rank, talent, aura and proc the Protection card casts or the build and leveling order take, read line by line in the fork (worktree at 0fb4dcebb) against the site's 70790 `spellconst/warrior.json` and `talents/warrior.json`, then against the live 1.60.1.70888 rows fetched from wago.tools for the same ids (SpellEffect, SpellMisc, SpellCooldowns, SpellPower, SpellAuraOptions, SpellCategories, SpellLevels, SpellClassOptions, Spell, SpellXDescriptionVariables and SpellDescriptionVariables), the two builds diffed row by row. The push layer is the one the site already carries (Defensive Stance 80%, Defiance 7/13/20).
2. **Blizzard.** The "WoW: Forever Hotfixes" thread read to post 8 (12 November 00:48 UTC): the only warrior lines since 6 November are the Reckless Swings tooltip, Thunder Clap's Cooldown Manager display and Sweeping Strikes' charges, none a number this card uses. The 30 September "Class Deep Dives: Priest and Warrior" and the 8 October beta notes carry the three rules this lane acted on (rage from damage taken, Thunder Clap, Sunder Armor's attack power threat); the 3 November notes carry the stance and Defiance figures the engine already has.
3. **Real fights.** foreverlogs.gg: all 25 public reports (2815 to 2853, 5 to 12 November). 15 of them hold a Protection warrior tanking: 26 distinct tanks at 15 to 43, 1,834 encounters, 220,726 events from and against them (damage, casts, resources, buffs, healing, deaths) plus 15,986 avoidance events (block with the blocked amount, dodge, parry, miss). The warrior's level is its talent points plus 9, its attack power 3 x level less 20 plus twice the Strength its combatant snapshot shows, its weapon the snapshot's main hand. Our own API holds nothing newer than the 8 November advanced report the physical lane already read (a warrior at a logged level of 21 to 23); no Protection warrior at 44 or above has logged a fight anywhere yet.
4. **Sims.** Every sim on the Windows box through `tools/remote-test.sh`, one at a time, `nice -n 15`, `GOMAXPROCS=5`: `spec-breakdown -tank` on the published band entries (bare 20 to 60, raid 60, both factions at 60) and the walk levels 26 and 34, 10,000 iterations, seed 1 for the before/after tables and 50001 where a call was within the error. The fork's own suites ran in the box's fork clone.

## Engine against the client

Every row the card touches, at 70790 with the push and at 70888:

| spell or talent | client | engine | verdict |
|---|---|---|---|
| Sunder Armor 7386 to 11597 | -90 to -450 armor a stack, flat threat 34/75/117/158/206 (effect 63), 15 rage, 30 s | the same, plus 5% of attack power threat (server, sixtysim's beta measurement) | aligned; the attack power term is unconfirmed, nothing in the client or a log can read threat |
| Shield Slam 23922 to 23925 | 430/505/580/655 with variance 0.043 to 0.046 (421 to 439 at rank 1, 640 to 670 at rank 4), 20 rage, 6 s category cooldown, a dispel; "increased by your Block Value", no number | the roll plus one times block value; flat threat 508 | aligned on damage; the block value multiple and the 508 are unconfirmed (L7) |
| Revenge 6572 to 25288 | 22/34/48/82/121/153 with variance 0.15 to 0.21, 5 rage, 5 s; description variable 1024: `$APR=$?a71[0.225][0.25]`, a quarter of attack power shown as 0.225 inside Defensive Stance's -10% damage | the roll plus 0.25 x attack power, the stance's 0.9 on top | aligned (the client's own variable; L2 confirms it in logs) |
| Improved Revenge 12797 | aura 108 (damage) +20/40/60, mask 1024 = Revenge | +20% a rank on Revenge | aligned |
| Shield Block 2565 | +75 block (aura 51) for 7 s, ProcCharges 2, 10 rage, 5 s cooldown; "will only block $n attacks" | 75%, 7 s, 2 charges | aligned (L3) |
| Bloodrage 2687 and 29131 | 10 rage at once (effect 30, 100 tenths) and 1 a second for 10 s (aura 24, period 1000); SpellPower PowerType -2, PowerCostPct 20 | the rage; no health cost | CONFIRMED-BUG, fixed (B2) |
| Improved Bloodrage 12301 | aura 108 misc 8 (every effect) +25/50, mask 256 = 2687 and 29131 | x1.25 / x1.5 on both | aligned |
| Last Stand 12975 | 3 min cooldown, a dummy effect; the talent text's 30% for 20 s | 30% of maximum health for 20 s, lost after | aligned |
| Concussion Blow 12809 | 5 s stun, 10 rage, 45 s, no damage | not registered (the Shield Slam prerequisite) | aligned |
| Shield Wall 871 | -60% all schools for 12 s, 15 min, on the global cooldown | the same | aligned |
| Improved Shield Wall 12312 | -330 / -660 s of cooldown (aura 107 misc 11, mask 8192) | the same | aligned |
| Defensive Stance passive 7376 | -10% taken, -10% dealt, +80% threat (SpellEffect 686475: the 70790 client row reads 30 with push 113326 to 80; the 70888 client row reads 80 itself) | 0.9, 0.9, 1.8 | aligned; the push dependence ends at the client move (MR83) |
| Defiance 12792 | aura 10 +15 at rank 3 in both clients; push 113326 and Blizzard's 3 November notes: 7/13/20 | 7/13/20, multiplied with the stance's 80%, shield only | aligned (hand row until a client carries it) |
| Shield Specialization 12298 | +5% block, aura 42 at 100% to 1310318 (5 rage) | 1% block and a 20% chance of 5 rage a rank | aligned (L4) |
| Master of Defense 1310316 | aura 42 to 23602 (5 rage) on dodge or parry with a shield | 50/100% of 5 rage | aligned (L4) |
| Anticipation 12297 | aura 30 misc 95 (Defense) +20 at rank 5 | 4 a point | aligned |
| Bastion 16538 | aura 79 +10 on school mask 1 (physical) at rank 5 | +2% a point on the physical school with a shield | aligned |
| Focused Rage 29787 | aura 107 misc 14 -30 tenths, mask {1852722926, 1057316, 1}: Sunder Armor, Shield Slam, Revenge, Concussion Blow, Heroic Strike, Thunder Clap, Demoralizing Shout in; Shield Block, Bloodrage, Battle Shout and Shield Wall out | -1 a point on every `SpellFlagOffensive` spell, the same set | aligned |
| Improved Sunder Armor 12308, Improved Thunder Clap 12287 | -30 and -60 tenths (masks 16384 and 128) | -1 and -2 a point | aligned |
| Thunder Clap 6343 to 11581 | 10/23/37/55/82/103 flat, -20% attack speed, 20 rage, 6 s category, Battle or Defensive Stance; no attack power coefficient in either client | the flat roll, threat x2.5 | CONFIRMED-BUG on damage, fixed (B3); the threat multiplier is vanilla's and unconfirmed |
| Demoralizing Shout 1160 to 11556 | -49 to -196 attack power and 1.4 a level, threat 11 to 43 and 0.8 a level (effect 63) | the same | aligned |
| Heroic Strike 11567 | weapon damage +138 (effect 17), 15 rage; no threat row in the client | 138, threat 20 to 175 by rank from the Classic server's spell_threat table | aligned on damage; the threat rows are vanilla's and unconfirmed |
| Battle Shout 11551 | +111 attack power and 0.6 a level, 3 min | the same | aligned |
| Tactical Mastery 1310185 | 10 rage kept, baseline at 14 | the same | aligned |

The 70790 to 70888 diff of these rows moves one number: Defensive Stance's threat effect (30 to 80), which the push already carried.

## Logs against the engine

| # | finding | evidence | verdict |
|---|---|---|---|
| L1 | **Thunder Clap scales with attack power.** Each tank's non-crit Thunder Clap median read against its own Revenge median (Revenge's roll plus a quarter of attack power, times Improved Revenge where the tree says so, gives the target's armor and the stance in one factor): the Thunder Clap hit before mitigation is the rank's amount plus 0.051 to 0.071 of attack power, median 0.0625, in eleven tanks at 15 to 33 (Gee, Geemee, Graghon at 18 and 27, Blinky at 15 and 18, Doublewide, Miniwaz, Adriän at 22 and 25, Corve; 40 to 803 hits each). The engine's flat 23 at rank 2 was 50% under the logged hits at 27. | war2815 to war2853 | CONFIRMED-BUG, fixed (B3) |
| L2 | **Revenge is the roll plus a quarter of attack power, times 1.6 under Improved Revenge.** Solving each tank's Revenge median for the attack power share at the tree's Improved Revenge rank gives 0.20 to 0.28 (median 0.23) for the nine tanks with the talent and 0.22 to 0.24 without it. | the same logs | aligned (the client's variable 1024) |
| L3 | **Shield Block: two blocks or 7 s.** Each cast's buff shows two "removed" rows in the log, one at each block (the parser writes a charge drop as a removal): 455 casts with 704 removals on Miniwaz, the same ratio on every tank; a window with no block runs 6.8 to 7.3 s. Inside a window the tanks block 67 to 77% of swings (base block plus 75%, less the dodges, parries and misses), outside 10 to 32%. | war2839 and the avoidance stream | aligned |
| L4 | **Shield Specialization and Master of Defense pay 5 rage.** 1310318 "Shield Specialization" 5 rage on 7,286 blocks, 23602 "Shield Specialization" (Master of Defense's energize) 5 on dodges and parries; 2 to 4 when the bar was within 5 of full. | resources stream | aligned |
| L5 | **Bloodrage.** 10 then 1 a second for 10 s; 13 then 1 at Improved Bloodrage 1 (12.5 and 1.25 shown rounded), 15 then 2 at rank 2 (15 and 1.5 rounded). | resources stream | aligned (the engine's 1.25 and 1.5 are the client's aura 108 on every effect; the log rounds its display) |
| L6 | **Casts a minute on boss kills.** Logged medians at 15 to 24 / 25 to 29 / 30 to 34 (50 / 48 / 15 kills): Sunder Armor 6.4 / 8.8 / 14.3, Revenge 6.8 / 7.3 / 7.3, Shield Block 7.6 / 7.7 / 10.1, Thunder Clap 3.5 / 3.9 / 2.2, Heroic Strike 0 / 0 / 3.1, Demoralizing Shout 2.2 / 1.6 / 1.0. The sim on the 48 s pull at 20 / 26 / 30 / 34: Sunder 13.4 / 15.6 / 20.3 / 22.7, Revenge 10.9 / 10.9 / 10.3 / 10.1, Shield Block 11.4 / 11.4 / 11.0 / 11.0, Thunder Clap 4.3 / 3.9 / 4.6 / 4.6, Heroic Strike 0.1 / 0.1 / 2.2 / 3.7, Demoralizing Shout 2.4. The same shape at a player's lower pace, as the tanks lane found at 30 to 34. | casts stream, boss kills of 15 s or more | aligned |
| L7 | **Shield Slam.** Six landed hits in the public logs, two warriors at 40 to 43 (Eck 322 and 379, Piff 251, 326 and 329 after armor; a 95). Read against their Revenge hits the hit before mitigation is the rank 1 roll plus about one times a level 40 shield's block value; six hits cannot tell one times from two. | war2834 | open (n = 6) |
| L8 | **Damage taken a second on boss kills.** Median 19 / 27 / 36 at 15 to 24 / 25 to 29 / 30 to 34, upper quartile 28 / 39 / 50. The sim's bare tank takes 39 / 41 / 67 / 67 at 20 / 26 / 30 / 34, the band's 75th-percentile pull against a boss three levels up, as the tanks lane set it. | taken stream | aligned (design) |
| L9 | **Crit multipliers.** Thunder Clap 1.48 to 1.50 over a hit (the magic table's 1.5), Revenge 2.0. | damage stream | aligned |
| L10 | **Rage from damage taken.** These logs carry no power snapshot on a hit, so the rule cannot be read here. The one measurement is the physical lane's 8 November advanced log (653 hits on a tank at a logged level of 21 to 23, 0.0356 rage a point of damage after armor); at the warrior's own level and mail (about a quarter off each swing) the fixed rule and the old one land within 5% of it either way, which is why the old rule looked right at that level. A log with power snapshots on a plate tank at 40 or above is what pins the middle of Blizzard's table. | physical-logs.md L12 | open (B1's table) |

## B1: rage from damage taken ignores the tank's armor (CONFIRMED-BUG, fixed)

**Finding.** `core/rage.go` turned each hit taken into rage with Kalgan's formula on the damage after armor (2.5 x damage / c at the receiver's level). Blizzard's 30 September deep dive: "Rage gain from damage is calculated by ignoring Armor. So, wearing Armor will no longer reduce the Rage you gain. We are using the old formula as if you always had 50% damage reduction from Armor"; the 8 October beta notes: "Previously, we were balancing Rage gained around Warriors having 50% Armor ... We have rebalanced Rage gains based on appropriate values per level in the 20% to 40% range, varying by level", and the extra damage of critical strikes and crushing blows now counts. A raid tank at 60 sits at 55 to 60% armor reduction against a level 63 boss, so the engine paid it two thirds of the rage the game pays; and every point of armor the ranker weighed cost rage in the sim that it does not cost in the game.

**Fix.** `damageTakenRage`: a physical hit's rage is Kalgan's formula on the hit before the receiver's armor (`result.Damage / result.ResistanceMultiplier`, so critical and crushing multipliers stay in, as dealt), times one less the armor Forever assumes for the level. The table is unpublished: `assumedArmorReduction` is a straight line from 20% at level 1 to 40% at 60, the two ends Blizzard named, a hand row until a log pins the middle (L10). Partial magic resists are not armor and stay in the damage. Red first: `TestRageFromDamageTakenIgnoresTheReceiversArmor`, `TestAssumedArmorReductionRunsFromTwentyToFortyPercent` and `TestRageFromSpellDamageTakenKeepsThePartialResist` (F f770ad173: the core and warrior packages fail to build, the names are new; `TestBloodrageCostsAFifthOfBaseHealth` red on its own), green at F 313699c2a.

Shared core: the bear's rage from damage taken moves the same way (its suite golden `TestFeralTank` regenerated with the warrior's; its published numbers wait for the refresh, MR81).

## B2: Bloodrage costs a fifth of base health (CONFIRMED-BUG, fixed)

**Finding.** SpellPower for 2687: PowerType -2 (health), PowerCostPct 20. The engine granted the rage for nothing. At 60 the warrior's base health is 1,689, so each cast spends 338; the card casts it before the pull and whenever rage is under 60, about 3.5 times a raid fight.

**Fix.** `RemoveHealth` of a fifth of `GetBaseStats()[stats.Health]` on the cast, and the cast refuses below that health, as the client refuses a health cost it cannot pay. The health loss enters the TMI list (it is damage the healers cover) and not the boss's damage-taken figure. `TestBloodrageMatchesClient` now reads the cost type and share off the client row.

## B3: Thunder Clap takes a sixteenth of attack power (CONFIRMED-BUG, fixed)

**Finding.** L1. The client rows carry no coefficient (a server value), Blizzard says "significant damage scaling from Attack Power", and the logs say 0.0625 of attack power a hit.

**Fix.** `thunderClapBaseDamage`: the rank's roll plus 0.0625 x attack power (a hand row from the logs). The threat multiplier stays vanilla's 2.5 and is marked unconfirmed: Blizzard's same line says the threat was "rebalanced (reduced)" with no number, and no log carries threat. Thunder Clap is one cast in 30 s on the card, so the multiplier moves under 1% of threat either way; it is not offered as a filler for that reason.

## Before and after

Published band entries, the dungeon pull below 60 and the raid boss at 60, 10,000 iterations, seed 1, Alliance (Horde at raid 60). Threat a second and its change; damage taken a second, TMI and chance of death before and after:

| entry | threat before | after | change | damage taken /s | TMI | death |
|---|---|---|---|---|---|---|
| raid 60 Alliance | 1631.1 | 1682.9 | +3.2% | 228.3 to 228.9 | 23.0 to 23.5 | 0 / 0 |
| raid 60 Horde | 1601.7 | 1651.4 | +3.1% | 228.9 to 229.0 | 24.1 to 24.6 | 0 / 0 |
| bare 60 | 890.8 | 923.5 | +3.7% | 273.0 to 273.3 | 42.3 to 42.9 | 0 / 0 |
| bare 50 | 640.3 | 670.1 | +4.7% | 167.7 to 167.4 | 28.0 | 0 / 0 |
| bare 40 | 439.4 | 463.7 | +5.5% | 93.2 to 92.9 | 20.7 to 20.6 | 0 / 0 |
| bare 30 | 218.0 | 225.3 | +3.4% | 67.0 | 19.4 | 0 / 0 |
| bare 20 | 113.0 | 116.4 | +3.0% | 39.0 to 38.8 | 18.0 to 17.9 | 0 / 0 |
| walk 34 | 294.3 | 301.2 | +2.3% | 67.1 to 67.2 | 18.1 | 0 / 0 |
| walk 26 | 182.4 | 186.4 | +2.2% | 41.2 to 41.0 | 16.7 | 0 / 0 |

The threat is Heroic Strike: 39 to 49 casts a raid fight at 60, 27 to 36 bare, 1.1 to 2.0 on the 48 s pull at 40, and Shield Slam no longer waits for rage at 40 (6.3 to 6.9 casts a pull). Thunder Clap's own threat doubles from a small base (18 to 37 a second at raid 60). TMI rises half a point at 60 for Bloodrage's cost; chance of death is 0 everywhere before and after.

## The card and the walk on the fixed engine

**Heroic Strike's rage gate.** The card queues Heroic Strike at 40 rage. Against 30, 50 and 60 (threat a second, 10,000 iterations):

| entry | 30 | 40 (card) | 50 | 60 |
|---|---|---|---|---|
| raid 60 Alliance, seeds 1 / 50001 | 1675.6 | 1682.9 / 1681.7 | 1681.7 / 1680.0 | 1676.5 |
| bare 60 | 915.1 | 923.5 | 924.7 | |
| bare 50 | | 670.1 | 667.0 | |
| walk 45 | | 533.4 | 538.0 | |
| bare 40 Alliance, seeds 1 / 50001 | 449.0 | 463.7 / 463.7 | 470.7 / 471.1 | 472.0 |
| bare 40 Horde | | 477.6 | 484.2 | |
| bare 30 | | 225.3 | 225.2 | |
| bare 20 | | 116.4 | 116.3 | |

40 holds. 30 loses 0.4 to 3.2% everywhere it differs. 50 gains +1.5 to +1.6% at bare 40 (both seeds, both factions) and +0.9% at 45, ties at 20, 30, 60 and raid 60, and loses 0.5% at bare 50 (three combined errors), so it is not a gain at every level it touches and the adoption rule keeps 40; 60 is the same picture (+1.8% at 40, -0.4% at raid 60). A band-40 tank is rage-poor enough that a higher gate keeps rage for Shield Slam and Revenge; at 50 and above the gate no longer binds.

**The walk's first points (adopted).** The order took Improved Bloodrage at 10 and 11, Shield Specialization 12 to 16 and Improved Revenge 17 to 19. Now that Bloodrage costs health and rage from damage taken ignores armor, Shield Specialization first (10 to 14), Improved Revenge 15 to 17 and Improved Bloodrage 18 and 19 was tried level by level on the band 20 gear and the 48 s pull, 10,000 iterations (threat a second, the order / the swap; chance of death 0 in every cell):

| level | order | swap | gain |
|---|---|---|---|
| 11 (seeds 1 / 50001) | 27.1 / 27.1 | 31.1 / 31.1 | +14.8% |
| 13 Alliance / Horde | 38.1 / 41.6 | 41.9 / 45.3 | +10.1% / +9.0% |
| 14 (seeds 1 / 50001) | 75.9 / 75.7 | 79.7 / 79.5 | +5.1% |
| 15 | 82.6 | 87.5 | +5.9% |
| 16 | 86.2 | 92.5 | +7.2% |
| 17 | 92.5 | 99.0 | +7.0% |
| 18 (seeds 1 / 50001) | 102.9 / 102.8 | 106.1 / 106.1 | +3.1% |
| 19 | 108.2 | 108.2 | identical |

Every cell is beyond the error, Defiance still opens at 20 on ten points, and nothing from 20 on moves. Improved Thunder Clap at 18 and 19 instead, with Improved Bloodrage at 28 and 29, loses: 105.0 against 106.1 at 18, 106.2 against 108.2 at 19, 190.5 against 193.0 at 28, 192.6 against 194.7 at 29. The order, the guide's leveling list and the addon's walk follow (`TestProtectionWarriorTakesDefianceBeforeAnticipation` pins 14, 17 and 19). No band entry moves (band 20's talents at 20 are the same string); the walk's per-level figures at 10 to 18 wait for the refresh (MR84).

**Rank rows.** No Protection rank row moves: the three fixes change no rank's damage order (Thunder Clap's attack power term is the same at every rank, Revenge and Sunder Armor are untouched), so the round 30 rows hold.

## Held and open

| # | item | why |
|---|---|---|
| O1 | Shield Slam's flat threat (508) and Revenge's threat (2.25 x damage plus 2.25 x 2 x the rank's level) | vanilla-shaped figures the fork has always used; no client row and no log carries threat. Sixtysim's beta card puts Revenge and Heroic Strike at 30% of threat each and Shield Slam at 18%; ours at raid 60 reads Revenge 38%, Shield Slam 21%, Heroic Strike 16%, Sunder 7%. A Forever threat meter reading is what settles it |
| O2 | Sunder Armor's 5% of attack power threat | server side (Blizzard's 24 September line); sixtysim's beta measurement is the only figure |
| O3 | Thunder Clap's threat multiplier | Blizzard: "rebalanced (reduced)"; vanilla's 2.5 stays, under 1% of threat on the card |
| O4 | The assumed-armor table's middle (B1) | a 40+ plate tank's log with power snapshots |
| O5 | Shield Slam's block value multiple (L7) | six logged hits; a 40+ log with Shield Slam and the shield's block value |
| O6 | Heroic Strike at 50 rage on band 40 alone | +1.5 to +1.6% at 40, +0.9% at 45, -0.5% at bare 50: one card, one gate; a per-band gate is a card design question, not a sim bug |

## Gates

GATES_PENDING

## Measure at refresh

| what | specs / bands | why |
|---|---|---|
| every tank figure on rage from the swing before armor: picks, tank weights (armor and Stamina no longer cost rage), threat, the published metrics | warrior-protection every band and raid 60; druid-feral-bear every band and raid 60 (shared core) | B1 (MR81) |
| Protection warrior on Bloodrage's cost and Thunder Clap's attack power | warrior-protection every band and raid 60: TMI, threat; Arms and Fury cast Bloodrage too but their fights carry no death or health figure, so nothing of theirs moves (their suite goldens are unchanged) | B2, B3 (MR82) |
| the Protection walk's first points | warrior-protection leveling walk 10 to 18 (any published per-level figure there; the addon's walk); no band entry moves | the swap (MR84) |
| Defensive Stance's threat row at the client move to 70888 | the engine's 0.8 reads the push today; the 70888 row carries it, so the conformance server-hotfix entry for 686475 can drop then; Defiance 12792 still needs the push | MR83 |

## Shas

- **Fork `lane/sim-protwar` on master 0fb4dcebb:** f770ad173 (tests, red), 313699c2a (the three fixes), 837cca30a (the tank warrior and bear suite goldens). Master fast-forwarded to 837cca30a.
- **Site:** SHAS_PENDING
- **Engine pin after:** 837cca30a.
- **wow-player:** NO-SHIP on the first wording (four MUSTs: "fire" is sim talk, the Bloodrage rule is health below the cost, Thunder Clap's attack power share is small and the note should say so; three SHOULDs), every item taken verbatim, then SHIP on the second pass with one nit, taken (`protection-warrior/wow-player.md`).
