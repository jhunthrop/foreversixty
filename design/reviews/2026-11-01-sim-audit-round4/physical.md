# Sim audit, round 4: physical group (2026-11-01)

Arms, Fury, Protection, Assassination, Combat, Subtlety, Beast Mastery, Marksmanship and Survival (melee). Client 1.60.1.70707. Baseline: engine pin 640da215f on origin/main's published data. Fixes are on fork branch `lane/r4-physical` (E) and site branch `wt/r4-physical` (S). Measurements use `sim/cmd/spec-breakdown` at the level-60 raid preset, 10000 iterations, seeds 1 and 50001. The "after" binary is the lane engine with a simdb.bin regenerated locally from this branch's pipeline. That simdb.bin is not committed, because the nightly owns it.

## Cross-cutting

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| X1 | **Which weapon damage is true (assigned item 1).** 225 weapons disagree between `items/<class>.json` and simdb.bin. **Cause 1, all 225:** `pipeline/loot/weapons.py` overwrote every resolved row with the engine fork's `db.json`, which is Classic Era's itemisation. Elemental Mage Staff 944 is 147-221 there; Forever is 109-164 (42.66 DPS). Thunderfury 19019 is 60-145 there; the Forever tooltip says 44-115. Of the 225, 164 match wowhead Forever exactly on the simdb side. **Cause 2, 51 of the 225 plus 122 more weapons where the two files agree:** the client curve missed two rules. QualityModifier is a flat DPS offset that moves the average and not the spread: Thunderfury -20, Benediction -14, the AQ caster maces -10, Torturing Poker -2.5. Bit 0x100 of ItemSparse.Flags_4 marks a caster weapon, which deals 2/3 of the one-hand curve or 29/39 of the two-hand curve. The fitted interval on item level 30 and up is [0.6663, 0.6668] for one-handers and 0.7436 exactly for two-handers. With both rules, the curve reproduces wowhead Forever's damage on 1,027 of 1,036 uncommon-or-better melee weapons (868 before). The other 9 are hotfix rows newer than our 70685 seed (Fiery War Axe: 2.9 s in the seed, 3.6 s on the tooltip). Ranged, wand and thrown weapons carry neither rule. The fork now fills only the weapons that normalize left at zero. **simdb.bin was the closer of the two.** Its weapon rows move on 173 items and nothing else in simdb.bin changes. | wago ItemSparse and ItemDamage* at 70707; `raw/hotfixes/ItemSparse.csv`; `nether.wowhead.com/forever/tooltip/item/{944,19019,21407,18608,15444,2271}` (read 2026-11-01); classic tooltips for comparison; tests `test_a_quality_modifier_moves_the_average_but_not_the_spread`, `test_a_caster_one_hander_deals_two_thirds_of_the_curve`, `test_a_caster_two_hander_deals_29_39ths_of_the_curve`, `test_the_planner_reads_the_caster_bit_and_quality_modifier_like_the_sim`, `test_the_forks_era_numbers_never_replace_damage_normalize_resolved` (all failed first) | CONFIRMED-BUG (two generators) | S 5043869bd |
| X1a | DPS impact of X1. At band 60 raid, no published physical pick moves. Published low-band picks do. Band 30 bare rogues on Swinetusk Shank and Torturing Poker: Combat 66.6 to 60.9, Assassination 58.8 to 52.9, Subtlety 60.8 to 55.6. Protection band 20 on The Stitcher, threat 86.1 to 82.5. Feral bear band 30 on Staff of the Shade, 45.7 to 41.6. Thunderfury with `-gear main_hand:19019` goes from 629.6 to 593.2 on Combat and from 852.5 to 819.8 on Fury, so it would have beaten the published Combat set. The figures are horde, seed 1, 10000 iterations, before the ranker re-ranks. | spec-breakdown | measured | - |
| H1 | **Ranged attack power on items was dropped (fresh).** ItemSparse stat 39 is ITEM_MOD_RANGED_ATTACK_POWER. 37 rows carry it: Core Marksman Rifle, Rhok'delar, the Dawnstalker set, the rank-14 bows and guns, and Hyper Deluxe Sniper Rifle Mk XVII 279273, which is the published BM and MM ranged pick at +21. `gear.py` mapped stat 39 to None and `wowhead_items.py` did the same for `rgdatkpwr`, so simdb.bin had 0 on all 20 items wowhead also lists. | wowhead Forever's `rgdatkpwr` gives the same amounts as the client stat; tests `test_ranged_attack_power_is_kept_not_dropped`, `test_ranged_attack_power_maps_like_the_clients_stat_39` (failed first) | CONFIRMED-BUG | S 22fc3c763 |
| H2 | **"+N Weapon Damage" was dropped (fresh).** ItemSparse stat 83 was mapped to None. The tooltips say "+ 2 Weapon Damage" on Molten Heart of the Mountain 249470, +4 on Might of Cenarius 21189, +1 on Darkmantle Bracers 226830, +3 on Studded Timbermaw Brawlers 227809 and +1 on Sharpened Cutlery 274092. The fork's `db.json` gives Might of Cenarius `bonusPhysicalDamage` 4, and the curve budget reproduces every amount. simdb.bin now carries it on `SimItem.bonus_physical_damage`. The planner has no key for it, so `items/<class>.json` leaves it out. Molten Heart is the alliance trinket in every published physical set at band 60 except Protection, and in Enhancement, Retribution and Feral. | tooltips read 2026-11-01; test `test_item_stat_83_is_flat_weapon_damage` (failed first) | CONFIRMED-BUG | S 57c59ff31 |
| H3 | Other ItemSparse stats mapped to None, for the integrator and the casters lane (not fixed here). Stats 84, 86, 87, 88 and 89 are holy, nature, frost, shadow and arcane spell damage. They are on 19, 9, 40, 57 and 24 rows, and wowhead's `holsplpwr`, `natsplpwr`, `frosplpwr`, `shasplpwr` and `arcsplpwr` match one for one. Published examples: Fel Cape (+15 shadow) and Gloves of the Greatfather (+10 nature, Balance). Wowhead's school-power keys are untracked too: Elemental Mage Staff's +100 Fire and +100 Frost are missing. Stat 47 (spell penetration, 32 rows, `splpen`) is None while 119 and 121 map to spell_penetration. Stat 50 (bonus armor) is already counted in the armor curve, on 317 of 318 rows. Stat 37 (expertise, 9 rows) and stat 36 (haste, 102 rows, none in wowhead's feed) are unconfirmed. | wago ItemSparse 70707 against `raw/wowhead/items.json` | CONFIRMED (casters; needs web `STAT_KEYS`); 36 and 37 OPEN | - |
| X2 | Rage crit factor (assigned item 2): a second Blizzard post confirms it. The 2 October warrior update (Kaivax, `us.forums.blizzard.com/en/wow/t/2369360`) says "100% increased Rage (was 75%)". The development-notes thread (`/t/2360696/4`) repeats it, and Warcraft Tavern reports it. All of these are Blizzard text, so a log check is still open. The same notes say Unbridled Wrath no longer doubles rage for two-handers, which the engine follows (`talents.go:482-490`). They also say Dual Wield Specialization drops its off-hand rage, but the 70707 talent text still reads "+50% Rage generated by your off-hand attacks", and the engine follows the client. Heroic Strike queueing no longer raises off-hand hit, and the engine never removes the dual-wield penalty for a warrior. | forum posts read 2026-11-01 | aligned (Blizzard-only); log OPEN | - |
| X3 | Sunder's 5% AP threat: no second source found. The Forever notes cover only Sunder's instant tap. | notes above | OPEN | - |

## Warrior: Fury

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| F1 | Furious Precision 3/3 gives +10% off-hand hit. The breakdown shows main-hand white 18.8% miss and off-hand white 8.8%. Death Wish and Recklessness holds: 1.09 and 1.00 casts in 180 s. Bloodthirst is 45% AP (round 3). Deep Wounds is 60% of average weapon damage over 12 s; its ticks crit (37%) under the bleed rule. Enrage procs only on hits taken, per its 70707 text, so a DPS warrior gets none. Whirlwind hits with both weapons. | `talents/warrior.json`; breakdown `scratchpad/r4/physical/bd/warrior-fury.txt` | aligned | - |
| F2 | Molten Heart's +2 weapon damage (H2): alliance +0.49%. | H2 | CONFIRMED-BUG | S 57c59ff31 |

## Warrior: Arms

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| A1 | Bloodthrill 5/5 (a 20% Overpower window from main-hand hits on Rend targets): Rend is kept up (7.3 casts) and Overpower is cast 18.3 times at 78.5% crit (Improved Overpower +50%). Blood Craze 2/3 is a heal on being struck, so it adds no DPS. Two-Handed Weapon Specialization is 3%, Impale 20% and Weaponmaster as its client text. | breakdown; talent text | aligned | - |
| A2 | Molten Heart (H2): alliance +0.34%. | H2 | CONFIRMED-BUG | S 57c59ff31 |

## Warrior: Protection

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| P1 | No change at band 60. Band 20's The Stitcher drops from 26-49 to 17-33 (X1). Even then it misses the tooltip's 13-26, because its hotfix row is stale. Bastion 10%, Defiance 15% and Focused Rage -3 match their client text. | X1 | aligned at 60 | S 5043869bd |
| P2 | Sunder's 5% AP threat. | X3 | OPEN | - |

## Rogue: Assassination, Combat, Subtlety

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| R1 | **Wound Poison takes Venom and Vile Poisons (assigned item 3).** Both talents' masks (8192, 65536) exclude Wound Poison's (268435456). The engine scaled it anyway. It deals no damage and no rotation casts it, so nothing moves. | test `TestVenomAndVilePoisonsLeaveWoundPoisonAlone` failed first; `client_values_test.go` expectation corrected | CONFIRMED-BUG (0%) | E 8f654098f |
| R2 | Molten Heart (H2): alliance Assassination +0.52%, Combat +0.66%, Subtlety +0.72%. | H2 | CONFIRMED-BUG | S 57c59ff31 |
| R3 | Checked with no finding: Puncturing Wounds (Backstab crit 61.7% with Hack and Slash's dagger +5%), Thousand Cuts, Lethality 12% a rank, Seal Fate, Relentless Strikes 25 energy, and both Mutilate hands. Set bonuses on the published sets: Shadowcraft 2/3/5, Stormshroud 2/3 and Devilsaur 2 are modelled. Shadowcraft 4 (Crafted Shadows 1302698) is a 5% snare escape when struck, per the client and the Forever tooltip, so it has no sim effect. Rogue Armor Energize runs at the engine's unstated 1 PPM, because the client states ProcChance 100 and no PPM row, so that rate rests on one source. Combat's Slice and Dice is still round 2's C2 HOLD. | breakdowns; SpellAuraOptions 136944, 246754 | aligned | - |
| R4 | Band 30's caster-dagger and qm picks (X1a): -8.6% to -10% until the ranker re-picks. | X1 | CONFIRMED-BUG | S 5043869bd |

## Hunter: Beast Mastery, Marksmanship, Survival

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| B1 | **The pet did not inherit the hunter's hit (assigned item 4, fresh).** Hunter Pet Scaling 415429 carries hit channels (aura 54 and 55, server amounts), the same as Warlock Pet Scaling 416189, and this fork's demons inherit hit (round 3). wowsims/sod gives the pet all of the hunter's hit on the same spell (`sim/hunter/pet.go` `makeStatInheritance`, read 2026-11-01). Horde, published sets: BM 656.9 to 661.1 (+0.6%), MM 660.9 to 665.7 (+0.7%). Haste (aura 319) is not inherited, as in sod, and is still open. | SpellEffect 415429 rows 9-11; test `TestHunterPetInheritsTheHuntersHit` failed first | CONFIRMED-BUG (two sources plus our warlock precedent) | E 8948a94b2 |
| B2 | Ranged attack power on the published rifle (H1): BM +0.55%, MM +0.64% horde. | H1 | CONFIRMED-BUG | S 22fc3c763 |
| B3 | Pet family: still the cat. No Forever raid-pet consensus has been published. | - | OPEN | - |
| B4 | Rhok'delar (18713), with its stat-39 RAP now counted, sims MM at 688.6 against the published rifle's 665.1 (+3.5%). It is not among the ranker's candidates. Its source (the MC quest chain) needs a ranker check. | `-gear ranged:18713` | OPEN (ranker) | - |
| V1 | Survival: Strider Kick 1317257 is normalized 100% weapon damage on an 8 s cooldown (effects 121, 31 bp 100). Raptor Strike uses non-normalized weapon damage plus 70 (effect 58). Mongoose Bite is normalized plus 57 (effect 121, 5 s). Lacerate is aura 3 at 58, and Lacerating Strikes is 40%. All match the engine. Molten Heart (H2): alliance +0.57%. Raptor Strike's queue still removes the dual-wield penalty. Forever removed that for Heroic Strike only, so this is OPEN. | SpellEffect, SpellCooldowns 70707 | aligned; H2 CONFIRMED-BUG | S 57c59ff31 |
| M1 | Trueshot Aura is ranged AP only, 15/22/30/40/50 by rank (aura 124), as in `buff_ranks.go:127`. Beaststalker 4 (Beast Unleashed) is a root escape when struck, so it has no sim effect. Sniper Shot is not taken by the search's build. | SpellEffect 1299346, 20906; tooltip 1302712 | aligned | - |

## Before/after (raid, band 60, 10000 iterations; seed 1 / seed 50001; absolute at seed 1)

| Spec | alliance | horde | alliance before -> after (s1) |
|---|---|---|---|
| Arms | +0.34 / +0.34 | 0.00 / 0.00 | 726.5 -> 729.0 |
| Fury | +0.48 / +0.50 | 0.00 / 0.00 | 764.0 -> 767.7 |
| Protection (threat) | 0.00 / 0.00 | 0.00 / 0.00 | 1038.1 -> 1038.1 |
| Assassination | +0.52 / +0.52 | 0.00 / 0.00 | 601.0 -> 604.1 |
| Combat | +0.66 / +0.66 | 0.00 / 0.00 | 608.1 -> 612.1 |
| Subtlety | +0.72 / +0.71 | 0.00 / 0.00 | 545.3 -> 549.2 |
| Beast Mastery | +1.48 / +1.56 | +1.19 / +1.14 | 640.5 -> 650.0 |
| Marksmanship | +1.39 / +1.42 | +1.36 / +1.44 | 646.2 -> 655.2 |
| Survival | +0.56 / +0.58 | 0.00 / 0.00 | 730.0 -> 734.1 |

Horde sets do not wear Molten Heart, so horde moves only with the hunter fixes. Outside the group, alliance seed 1: Feral +0.70% (582.8 -> 586.9), Retribution +0.15%, Enhancement +0.13%. Balance, Fire and Feral bear do not move at band 60. Lower bands move wherever a caster weapon or a qm weapon was picked (X1a), across most specs. All of these need the next simdb and bis.yml regeneration to reach the published numbers.

## Still open

- Nine weapons (X1) and The Stitcher sit on hotfix rows newer than the 70685 seed. A fresh hotfix-cache dump settles them.
- Caster school spell damage, spell penetration (47), expertise (37) and haste rating (36) are dropped (H3). This needs the casters lane and the web `STAT_KEYS`.
- Rage crit factor and Sunder threat (X2, X3); Rogue Armor Energize's 1 PPM; the off-hand rage constant: each needs a level-60 log.
- Pet family; pet haste inheritance; Raptor Strike's queue lifting the dual-wield penalty.
- Rhok'delar's absence from the ranker's candidates (B4).
- TestRotationLadder goldens move for every hunter and every alliance physical spec once simdb is regenerated. They are left for the integrator.
