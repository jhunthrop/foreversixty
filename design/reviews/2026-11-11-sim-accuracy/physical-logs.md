# Sim accuracy 2026-11-11: physical specs against real fights

Lane brief: `design/reviews/sim-accuracy/brief.md`. Site from origin/main 5b4dc79a1, fork from master 2b10bae6c (the pin). Client on the site: 1.60.1.70790. Method: the warlock lane's "Logs against the engine" (`warlock.md`), repeated for warrior (Arms, Fury), rogue (Combat, Assassination, Subtlety), hunter (Beast Mastery, Marksmanship, Survival) and Feral cat.

**Totals: 0 confirmed bugs. 12 checks aligned, 2 held, 2 open (one of them no data at all), 1 not comparable. Engine pin unchanged (2b10bae6c); nothing to re-rank.**

## Method

1. **Our own API, advanced logging.** Report gxrgvz4yzsi2 (8 Nov, 104 fights, 6 boss kills) serves each fight's full event table (`logs-data/reports/<id>/fights/<n>/events.parquet`) with the advanced fields: the unit's current power and level, attack power, max health, the glancing flag, the off-hand flag on misses, and `base_amount`. `base_amount` is the hit before armor and before the crit or glance multiplier (a crit's amount over its base is twice a normal hit's ratio). The fight summaries carry COMBATANT_INFO gear. The logging group is five players, two of them in scope:
   - Bùckk, a warrior with Hamhock's Cleaver (one-hand axe, 2.2) and Arctic Buckler, tanking. Ranks cast: Rend 4, Cleave 2, Thunder Clap 3.
   - Pink, a dual-wield rogue with Cruel Barb (sword, 2.8) and Bite of Serra'kis (dagger, 1.3). Ranks cast: Sinister Strike 1759, Eviscerate 6762, Kick 2, Instant Poison 1.

   The log's own level field reads 21 for both in the dungeon fights and 23 later, below the level their ranks need (30 and 26): either the zone scales them or the field is not the character level. Checks that depend on the player's level (glancing, rage from damage taken) are read at both.

   The 20 Oct report (z2opt2x2gs5u) holds only levels 1 to 7 and adds nothing here.
2. **foreverlogs.gg.** All 23 public reports (2815 to 2849, 5 to 11 Nov, Ragefire Chasm to Zul'Farrak plus City of Dalaran and Excavation Site: Wetlands): every damage, cast, avoidance, resource and buff event from each warrior, rogue, hunter and druid, and from every pet (owner from the event), via the report API with browser headers. The API rate-limits hard (HTTP 429 with a 5 to 9 minute Retry-After); one paced fetcher took about two hours. These logs carry no glancing flag, no off-hand flag, no power and no level, so they settle rates, ratios and timings, not absolute per-hit formulas.
3. **Engine.** The fork's rage bar, energy bar, attack table, Shank, Sinister Strike, poisons, Windfury Totem, Touch of the Grave, Deadly Aspects, Serpent Sting and Volley, Omen of Clarity; client rows (SpellEffect, SpellAuraOptions, SpellLevels) for 1759, 1329283, 13165 to 25296 and 6150, 71.
4. **Sims** on the Windows box (`tools/remote-test.sh --pin-engine`, one `spec-breakdown` at a time, `nice -n 15`, GOMAXPROCS=6, 2,000 iterations, horde): Beast Mastery bare at levels 25, 30, 35 and 40, and Combat bare band 20 at level 28 with the logged rogue's two weapons.

## Logs against the engine

| # | finding | evidence | confirmed? |
|---|---|---|---|
| L1 | **Normalized rage is exact on a one-hander.** Every landed white swing of the 2.2 axe adds 7.6 rage (291 of 406 swings read 76 tenths, 42 read 77): 2.2 x 3.46 = 7.61. A crit adds 15.2 (29 of 40): the 2x crit. Read as the rage change between the swing and the warrior's previous snapshot, under a second apart, with no cast between. | own 8 Nov, adv power on SWING_DAMAGE | Aligned (`normalizedRagePerSpeedOneHand`, `normalizedRageCritMultiplier`; first log confirmation of round 1's measured constants) |
| L2 | **Energy regenerates at 10 a second, continuously.** Cast-to-cast accounting (the snapshot at each cast before its cost, minus Shank's 15 back, windows with no other cast): slope 9.8 a second over 427 windows, 10.6 to 10.75 over windows of 3 to 15 s. The bar moves in 1 to 11 point steps, never in 20-point ticks. | own 8 Nov, adv power on the rogue's casts | Aligned (`EnergyRegenPerSecond`; the engine comment said no rate was published) |
| L3 | **Crit multipliers.** From base to amount: white 2.00 (rogue 190 crits, warrior 57), Sinister Strike 2.00, Shank 1.96, Revenge 2.02; Thunder Clap 1.51 (magic defense, `DefenseTypeMagic`); Instant Poison 1.53. Across foreverlogs (median of each player's crit over non-crit): see the table under "foreverlogs aggregates". | own 8 Nov; foreverlogs 23 reports | Aligned |
| L4 | **Shank is the engine's Shank.** Cost 5 (adv power cost on all 137 casts), 15 Energy back (134 of 136), no miss, dodge or parry (one evade on an evading mob), crits (15 of 136), and its base is half the off-hand white hit's (13.0 against 26.7), so 25% of the raw off-hand weapon, as `shankWeaponDamage` times `OHWeaponDamage`'s 0.5. It never procs Instant Poison: 0 of 136, against 17 to 19% on every other landed hit. The client's SUPPRESS_CASTER_PROCS reading (`ProcMaskEmpty`) holds. | own 8 Nov | Aligned |
| L5 | **Off-hand white damage is half.** Fitting MH base = 43.5 + 0.2 AP per attack power snapshot and putting that AP into the dagger's 23 + 0.093 AP gives an off-hand factor of 0.50 to 0.53 in 8 of 10 snapshots (0.37 and 0.40 in two short fights, likely a weapon swap). | own 8 Nov | Aligned (no Dual Wield Specialization visible) |
| L6 | **Sinister Strike sits 12.0 ± 0.8 above the same main-hand white hit.** The client and engine agree on the spell (1759: effect 121 normalized, +15, one combo point, 45 Energy). At the attack power the white hits imply (about 340), normalized +15 predicts SS 117.3 against the white 112.1. The log's 124.1 is 6% higher. Aggression 3/3 (+6% to Sinister Strike) closes it exactly. The log carries no talents. | own 8 Nov | Held (talents unknown) |
| L7 | **Instant Poison.** Both hands proc it (main hand 15.0%, off hand 16.7%, Sinister Strike 16.9% of the landed hits right before a proc). 322 procs in all, 17 to 19% of the eligible hits, against the engine's 20% plus Improved Poisons. The rogue applied poison 4 times in 2.8 hours, so uptime explains a shortfall. Rank 1's damage, 13 to 21 before modifiers (mean 17.8), is the engine's 15 ± 2 plus 0.01 of the implied attack power. | own 8 Nov | Rate held (uptime unknown); damage aligned |
| L8 | **Windfury Totem procs on about 20% of landed main-hand hits.** Rogue: 56 extra attacks, 19% of landed main-hand white hits and Sinister Strikes inside the fights where the totem was up. Warrior and paladin: higher, but their windows are drawn from the procs themselves, which biases them up. | own 8 Nov, SPELL_EXTRA_ATTACKS 8516 | Aligned (20%, white or yellow) |
| L9 | **Touch of the Grave's 15 s internal cooldown.** See the aggregates below. | foreverlogs | Aligned |
| L10 | **Deadly Aspects (Quick Shots, 6150) stays within 2% a rank.** In the client, Aspect of the Hawk carries the 6150 proc at chance 0, and talent 19552 adds 2% a rank. See the aggregates below. | foreverlogs; SpellEffect 690401, SpellAuraOptions 125507 | Aligned |
| L11 | **Pet share.** Boss damage only, hunters with a pet out: see the aggregates below. The sim's Beast Mastery bare band 20 at level 25: the cat does 28.4% (white 24.0, Bite 2.7, Claw 1.7). | foreverlogs; box sim | Aligned |
| L12 | **Rage from damage taken.** 653 hits taken (no block) against the warrior's next rage snapshot: 0.0356 rage a point of damage, conversion c = 70.2, the same at 1,429, 1,539 and 2,000 max health, so it is damage-based, not health-based. The engine's c at the logged level 21 is 66.2 (73.6 with Defensive Stance's 90% from forever-bugs #409); at level 30 it is 100. The log matches the engine at its own level field within 5%. | own 8 Nov | Aligned at the logged level (the warrior's real level is uncertain; see Method 1) |
| L13 | **Glancing at and above the player's level.** Rogue: glances 1% against level 27, 6% against 28, 19% against 29 and 27% against the level 31 boss (22 of 81). Warrior: 1 of 37 against the boss. The glance penalty against the boss is 0.67 to 0.84 of a normal hit (mean 0.75). The engine at +3 gives 40% and 0.55 to 0.75. Its +2 gives 30% and 0.80 to 0.90. The rogue's ranks put him at 26 to 29, the log's level field says 21, and the mix of a sword and a dagger muddies the penalty. | own 8 Nov | Open (needs a log with a known level against a +3 boss; at 60, the canonical 40% against level 63 is unchallenged) |
| L14 | **Dual-wield misses by hand.** Main hand 17.0% (103 of 605 swings), off hand 21.8% (233 of 1,068), against level 25 to 31 mobs. The engine gives both hands the same 19% penalty. A dagger skill below the sword's explains the gap. | own 8 Nov | Not comparable (weapon skills unknown) |
| L15 | **Arms, Fury and Survival in melee: next to no data.** Of 141 warrior rows, every one with 80 or more swings tanks (Shield Block, Revenge, Sunder) except Garry, a dual-wield warrior at 26 or more who mixes Heroic Strike (220), Execute (82), Overpower (56) and Slam (43) with Revenge: his white swings miss 18.8% (of 954), the dual-wield 19% less a little hit or level, and show no Unbridled Wrath. Three more Arms or Fury rows have 10 to 16 abilities each. Doublewide shows Unbridled Wrath on 11.6% of white hits (rank unknown, 60% at 5/5). No hunter fights in melee: the most Raptor Strikes, 293, come with 6,837 Auto Shots. | foreverlogs | Open (no usable logs) |
| L16 | **Omen of Clarity in Cat Form.** The cat-only druid (Valdo, 1,953 landed melee hits) gains Clearcasting 46 times, 2.4% of hits. That fits 2 a minute on the 1.0 s paw under the 10 s proc cooldown, the engine's existing fit (10 beta cats, 346 gains). | foreverlogs | Aligned |
| L17 | **Auto Shot keeps its timer through Aimed Shot's 2 s cast and Multi-Shot.** For 14 hunters, the gap between two landed Auto Shots on one target is the same with an Aimed Shot cast inside it as without: 1.00 of the hunter's own median over 581 such gaps, and 1.00 over 651 with a Multi-Shot inside. | foreverlogs; SpellMisc 19434 (cast index 5, 2,000 ms) | Aligned (round 21's Auto Shot through casts) |

## foreverlogs aggregates

All 23 reports: 374,190 events from 635 warrior, rogue, hunter and druid rows (165 druids, 141 warriors, 108 hunters, 88 rogues with 50 or more events) and 44,774 pet events.

| check | figure | engine |
|---|---|---|
| Crit over non-crit (L3), median of each player's ratio | warrior white 2.04 (18 players, 671 crits), Revenge 2.01, Rend ticks 2.00; rogue white 1.99 (16, 1,960), Sinister Strike 1.98 (13, 458), Backstab 2.04, Eviscerate 1.91; cat white 2.00 (19, 992), Claw 2.04, Rake 2.02, Shred 1.97, Rip 2.05; Auto Shot 2.02 (23, 1,027), Arcane Shot 2.00, Multi-Shot 2.02, Serpent Sting ticks 2.01 (18, 681); magic-table hits 1.40 to 1.55 (Thunder Clap, Moonfire, Explosive Trap, Instant Poison) | physical and ranged 2.0, magic 1.5; Forever's DoT ticks crit on their spell's table (`critMultiplierAs`) |
| Touch of the Grave (L9) | 7 Undead (5 rogues, 2 warriors), 317 gaps between procs inside one fight: the shortest is 14.99 s (two 0 s gaps are one proc's partial-resist row) | 15 s (`touchOfTheGraveICD`, 27 Oct notes) |
| Deadly Aspects (L10) | 13 hunters with Quick Shots: 582 on 6,836 landed Auto Shots (8.5%); the highest are Pun 11.0% (112 of 1,015), Sigmund 11.0% (72 of 652), Owen 11.0% (59 of 534), one standard error over 10% | 2% a rank, 10% at 5/5 |
| Pet share on bosses (L11) | Beast Mastery-leaning hunters 25 to 48% (median 32%, 9 hunters, levels 18 to 40); Marksmanship with a pet out 18 to 30% | Beast Mastery bare: 28.4% at 25, 27.6% at 30, 29.6% at 35, 31.1% at 40 |
| Unbridled Wrath (L15) | Doublewide: 66 on 570 landed white swings (11.6%) | 12% a rank, white swings only |
| Omen in Cat Form (L16) | Valdo: 46 Clearcasting on 1,953 landed melee hits | 2 a minute, 10 s cooldown |

The two highest pet shares (Owen 48% at 28 or more, Draaxil 44% at 34 or more) sit well above the sim's 28 to 31%. A hunter's own damage in a dungeon falls with movement, mana and pulls while the pet keeps hitting, and the sim's bare hunter has no gear either way. The other seven sit at 25 to 35%. Not a stable gap.

## Sims (box)

| run | figure | log |
|---|---|---|
| hunter-beast-mastery, bare, band 20 at level 25, band 30 at 30 and 35, band 40 at 40, horde, 2,000 iterations | 101.4, 135.9, 153.6 and 209.6 DPS; the cat does 28.4, 27.6, 29.6 and 31.1%; at 25, Deadly Aspects procs 7.38 times a fight (about 9% of 81 Auto Shots) | boss pet shares 25 to 48%, median 32% (L11); Quick Shots up to 11% (L10) |
| rogue-combat, bare band 20, level 28, Cruel Barb and Bite of Serra'kis, horde, 2,000 iterations | 55.8 DPS; white miss 24.6% both hands at +3, glance 40%; Instant Poison II from 28 (the logged rogue still used rank 1) | L13, L14 |

## Confirmed and fixed

None.

## Held

- L6: Sinister Strike 6% above the client's normalized +15, the size of Aggression 3/3. Needs the rogue's talents.
- L7: Instant Poison's rate, 17 to 19% against 20% plus Improved Poisons. Needs an armory capture with the poison up all fight.

## Open

- L13: glancing at +2 and +3 below 60. Needs a known level against a +3 boss.
- Arms and Fury: no Arms or Fury warrior with a usable fight in any public Forever log (one dual-wield hybrid, Garry, without power data). Two-hand rage (4.50 a speed) and the off-hand rage factor (0.5, plus Dual Wield Specialization's +50%) are still unlogged.
- Survival in melee: no melee hunter in any public log. Lacerating Strikes and the queued Raptor Strike's off-hand miss (round 19 H4) stay open.
- Hunters' Ranged Weapon Specialization on Serpent Sting (round 19 H3): foreverlogs carries no attack power, and our own logs have no hunter.
- Level 60: every public report is a level 15 to 45 dungeon.

## Measure at refresh

Nothing. No engine, rotation or rank row moved.

## For the owner

- Our own report pipeline (advanced logging, parquet per fight) is the best accuracy source we have. A level-60 Arms or Fury warrior, a Survival hunter and a rogue with a known talent build, uploading one raid or dungeon night with advanced logging on, would close L6, L13 and the open rows above.
- foreverlogs rate-limits its report API at a few hundred requests. A lane repeating this method should fetch once into a shared folder.

## Sources

- foreversixty API reports gxrgvz4yzsi2 and z2opt2x2gs5u (`/v1/reports/<id>`, `logs-data/reports/<id>/fights/<n>/events.parquet` and `summary.json`).
- foreverlogs.gg reports 2815 to 2849 (`/api/reports/public`, `/api/reports/<id>/combatants-roster`, `/meta_actors`, `/meta_events`).
- Client 1.60.1.70790: SpellEffect, SpellAuraOptions, SpellLevels.
- forever-bugs #409 (Defensive Stance and rage from damage taken).
