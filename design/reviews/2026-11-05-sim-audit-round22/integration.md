# Sim-accuracy round 22: integration (2026-11-05)

Round 22 started from the landed round 21 (site d5cc4d6b4, fork 170a59c39) on the live client 1.60.1.70790. Lanes: platform (four follow-ups), items, warrior, druid, hunter, mage, paladin, priest, rogue, shaman, warlock, at most two at once. Leads came from round 21's open list, guide-accuracy rounds 22 and 24, the MythicSim comparison, the sim view lane's reading of real Fury runs, and the guild loop (rogue axes). Guild round 2, guides rounds 23 and 24 and the sim view branch landed on main during the round and are merged in.

**Not a zero round: 33 confirmed bugs, all fixed** (round 21: 41, round 20: 31).

| Lane | Confirmed | Origin | What |
|---|---|---|---|
| platform | 12 | 10 old, 1 70753, 1 R22 | Feint, Vanish, Counterspell and Tremor Totem from their trainer levels; Rend and Tear's assume-bleed wired (raid only); the Fire raid card's Fireball row; raid-buff scrolls ignored since an upstream change; the BiS panel reads the published race's rank row; talent spells learn at 10 + 5 x tier (trainables, spellranks) and card rows start where the build takes the talent (U1 corrects T1); card legend wording; an unknown aura no longer drops an APL gate (cross-class code, no published number moves); the raid card uses raid rank rows at 60; the card keeps a later step with its own note |
| druid | 5 | 4 old, 1 new-data | Leader of the Pack and Moonkin Aura exclusive (3% raid crit, not 6%: every spec moves); Rip and Ferocious Bite +25% attack power (3 Nov); no Cat Form below 20; caster Faerie Fire resets the swing |
| hunter | 4 | 2 R20, 1 R21, 1 old | Beast Mastery and Marksmanship low ranks of Arcane Shot and Serpent Sting while leveling; Aimed Shot rank 1 at 50-51; the pet takes the hunter's ranged crit when higher |
| priest | 4 | 4 old | Discipline tank Heal rank 1 at 22-35 and Lesser Heal below 16; Holy Heal rank 1 filler and race rows at 22-39; Vampiric Embrace cannot miss (16 Oct) |
| rogue | 3 | 3 old | Rogues use one-hand axes (proficiency.py, the ranker table, the addon); Assassination raid 31/20/0 with a three-point Slice and Dice; Subtlety Hemorrhage never at five points |
| paladin | 2 | 2 old | Holy party Flash of Light rank 5 at raid 60; Holy Light rank 1 at 18-19 |
| warlock | 2 | 2 old | Affliction raid build (Imp sacrifice, Nightfall 2, Improved Shadow Bolt 5, no Immolate); Demonology Soul Fire in Decimation at 56-59 |
| mage | 1 | 1 old | Fire raid build 14/34/3 with a Fireball filler |
| warrior, items, shaman | 0 | | Revenge 6, Battle Shout 7 and Heroic Strike 9 are book ranks (owner); weapon damage ranges match the client on every compared weapon (8 stale by post-run hotfixes, none on BiS); the Elemental stat-weight test expectation was stale |

By origin: 27 old, 1 new-data (the crit auras), 1 70753 (talent learn levels), 2 R20, 1 R21, 1 R22 (the platform lane's own T1, corrected in-round). Lane reports carry the evidence.

## Decided this round, not bugs

- **Forever has no sub-20 spell-power cut for heals** (three sources); damage stays OPEN-low. Rank rows that skipped ranks learned below 20 were the bug (priest, paladin).
- **Weapon speeds and damage ranges** match the client ItemSparse and ItemDamage tables on all 1,534 client-backed weapons; a permanent parity test skips without local client tables.
- **Auto Shot keeps firing through Sniper Shot** (same attributes as Aimed Shot; Blizzard 16 Oct names it).
- **The sim view findings on a real Fury run** (off-hand miss gap, boss armor 471, Berserking and Death Wish uptimes, Heart of the Lion, Sunder vs Expose) are all Forever-correct, except the scrolls (fixed).
- **Not adopted:** Shadow without Vampiric Embrace (+3.9%, but its group healing is why it is cast); Balance +2.35% and Enhancement leads (adoption rule / loses); Combat Slice and Dice; the Beast Mastery Multi-Shot order (no gate separates raid from leveling); bear Swipe first on packs (no level-by-level gate; guide line to scope).

## Integration

- **Fork:** int/r22 = bff826207 (pushed), includes fork master 996826ddf (sim view sample data). Fork goldens regenerated for the 3% crit, the scrolls and Cat Form; conformance keeps amd64 values on rounding-only rows.
- **Site pin:** 1134fa394 (engine bff826207).
- **Merges from main:** 178 commits (guild r2, guides r23 and r24, sim view), 16 conflicts. Lane rules kept on main's wording; proficiency.py takes main's public `VESTIGIAL_CLIENT_BITS` with the rogue axe bit dropped, and equip_skills.py now says the two tables agree; request.go keeps both assumeRaidBleed and RotationSHA.
- **Data regenerated here (data.yml is disabled):** spellranks.json from the committed spell constants and talents; trainables/ from the 70790 client tables (fetched into the ignored raw/). Every change is a talent-taught ability moving to its tier level (93 abilities); no other field moved. Manifest: only these files' hashes.
- **Regen at the pin (7a138c265, then 5dbffb0bf and 928ddcf14 for card notes):** make simdb, raid-consumes, addon-builds (Fire, Affliction, Assassination raid orders), dps-per-point, ladder goldens, addon-data and Data.lua, coach starters via persona_starter.py --all.
- **Integrator fixes:** healer rank tests and the rank-choices test follow the new rows; smoke-build entries for the Fire raid build's Arcane spells and the Affliction raid talents; the Affliction Wrack gate's rank; the Fury card test reads its note case-blind; Protection's Strengths back to four bullets (a round 21 regression on main); engine-pin now rewrites result-full.json too (its stale engine_version failed every warnings test); healer card notes rewritten so the first sentence (all the card prints) is right at 60.
- **Checks run:** sim request, leveling, rankchoices and internal packages; sim and api vet; data test_apl, test_addonrotation, test_addondata, addon-data and coach-rules --check; luac on Data.lua; web vitest for guides, sim and guide libs (with the fixture sync); prettier; apl-check. LAUNCH_MODE is on.
- **Failures that predate this round (also red on origin/main):** data test_build_stat_keys (70755 items emit `expertise`), test_loot_merge_parity, test_wago (tables list); fork `sim/web` needs binary_dist. The weapon parity test and test_wago also fail when the client tables sit in raw/ locally (they were moved aside for the run).
- **Not regenerated:** the BiS (its own run). Every BiS pick and weight was measured with 6% raid crit and must be re-ranked on this pin; rogue picks gain one-hand axes; `items/rogue.json` and simdb.bin's 8 hotfix-stale weapons refresh with the next data run.
- **Measuring tool note:** the scratch `rt` helper falls back to the leveling build when a raidBuild has an empty third tree, so Assassination was re-measured by hand on its raid string.

## Before/after, raid 60

Before: round 21's landed engine and site (`before.tsv`). After: int/r22 at the pin. Same frozen published BiS, each side's guide raid builds, 10,000 iterations at seeds 1 and 50001, both factions. Seed 1 value, then the change at seed 1 / seed 50001. Most of the drop is the 3% raid crit (Leader of the Pack and Moonkin Aura exclusive).

| Spec | Faction | Seed 1 | % |
|---|---|---|---|
| rogue-assassination | alliance | 777.3 to 780.6 | +0.42 / +0.49 |
| druid-feral | alliance | 886.1 to 888.0 | +0.21 / +0.21 |
| warlock-affliction | alliance | 753.1 to 753.7 | +0.08 / +0.05 |
| druid-feral | horde | 870.6 to 871.2 | +0.07 / +0.07 |
| rogue-assassination | horde | 780.7 to 778.8 | -0.24 / -0.38 |
| warlock-affliction | horde | 750.6 to 745.0 | -0.75 / -0.81 |
| mage-fire | horde | 721.0 to 712.4 | -1.19 / -1.28 |
| priest-holy (HPS) | alliance | 780.6 to 771.2 | -1.20 / -1.32 |
| druid-restoration (HPS) | horde | 935.3 to 923.9 | -1.22 / -1.30 |
| priest-holy (HPS) | horde | 797.8 to 788.0 | -1.23 / -1.27 |
| druid-restoration (HPS) | alliance | 939.5 to 927.3 | -1.30 / -1.28 |
| mage-fire | alliance | 735.6 to 725.0 | -1.44 / -1.49 |
| priest-discipline (HPS) | alliance | 539.9 to 530.4 | -1.76 / -1.83 |
| rogue-subtlety | alliance | 810.5 to 795.8 | -1.81 / -1.81 |
| priest-discipline (HPS) | horde | 536.2 to 526.0 | -1.90 / -1.88 |
| rogue-subtlety | horde | 805.3 to 789.0 | -2.02 / -2.01 |
| mage-frost | alliance | 632.3 to 619.2 | -2.07 / -2.10 |
| mage-frost | horde | 624.1 to 610.9 | -2.12 / -2.10 |
| warlock-destruction | alliance | 795.3 to 778.3 | -2.14 / -2.18 |
| shaman-restoration (HPS) | horde | 757.1 to 740.6 | -2.18 / -2.18 |
| warlock-destruction | horde | 792.2 to 774.7 | -2.21 / -2.21 |
| shaman-restoration (HPS) | alliance | 759.4 to 742.5 | -2.23 / -2.20 |
| rogue-combat | alliance | 798.1 to 780.2 | -2.24 / -2.23 |
| paladin-protection (threat) | horde | 1150.0 to 1123.8 | -2.28 / -2.28 |
| paladin-protection (threat) | alliance | 1151.9 to 1125.6 | -2.28 / -2.29 |
| druid-feral-bear (threat) | horde | 1750.6 to 1710.5 | -2.29 / -2.46 |
| shaman-elemental | alliance | 669.4 to 653.8 | -2.33 / -2.30 |
| warrior-protection (threat) | horde | 1488.0 to 1453.2 | -2.34 / -2.35 |
| shaman-elemental | horde | 666.1 to 650.5 | -2.34 / -2.33 |
| druid-feral-bear (threat) | alliance | 1756.3 to 1715.1 | -2.35 / -2.32 |
| warrior-protection (threat) | alliance | 1501.1 to 1465.5 | -2.37 / -2.35 |
| rogue-combat | horde | 800.9 to 781.6 | -2.41 / -2.43 |
| priest-shadow | horde | 735.5 to 717.5 | -2.45 / -2.42 |
| paladin-retribution | alliance | 792.3 to 772.9 | -2.45 / -2.45 |
| paladin-retribution | horde | 799.4 to 779.8 | -2.45 / -2.43 |
| priest-shadow | alliance | 742.6 to 724.3 | -2.46 / -2.44 |
| mage-arcane | alliance | 729.5 to 711.5 | -2.47 / -2.50 |
| mage-arcane | horde | 716.8 to 698.8 | -2.51 / -2.43 |
| warlock-demonology | alliance | 734.1 to 715.6 | -2.52 / -2.52 |
| warlock-demonology | horde | 724.3 to 705.6 | -2.58 / -2.60 |
| druid-balance | alliance | 652.8 to 635.7 | -2.62 / -2.63 |
| hunter-beast-mastery | alliance | 893.3 to 868.3 | -2.80 / -2.85 |
| druid-balance | horde | 646.9 to 628.7 | -2.81 / -2.77 |
| shaman-enhancement | alliance | 858.0 to 833.7 | -2.83 / -2.74 |
| hunter-beast-mastery | horde | 891.0 to 865.5 | -2.86 / -2.88 |
| warrior-arms | alliance | 853.0 to 828.5 | -2.87 / -2.76 |
| shaman-enhancement | horde | 864.5 to 839.5 | -2.89 / -2.83 |
| hunter-survival | alliance | 876.7 to 851.1 | -2.92 / -2.93 |
| warrior-arms | horde | 855.5 to 830.4 | -2.93 / -2.89 |
| hunter-survival | horde | 876.5 to 850.7 | -2.94 / -2.95 |
| hunter-marksmanship | alliance | 988.5 to 959.3 | -2.95 / -2.94 |
| hunter-marksmanship | horde | 989.5 to 960.0 | -2.98 / -3.01 |
| warrior-fury | alliance | 917.6 to 886.9 | -3.35 / -3.43 |
| warrior-fury | horde | 909.8 to 878.2 | -3.47 / -3.51 |
| paladin-holy (HPS) | horde | 709.3 to 684.6 | -3.48 / -3.62 |
| paladin-holy (HPS) | alliance | 727.1 to 701.6 | -3.51 / -3.53 |

## Open, for round 23

- **Live logs:** still none anywhere (site API: only the 20 Oct report; foreverlogs.gg's API now answers 403 "API access not allowed"; the Warcraft Logs v1 key in Secret Manager returns 401 "Invalid key specified": owner). Each lane's "Still open" names the log it needs.
- **Platform:** a mana-income APL value (specced in platform.md) for the Retribution and Protection pacing items; the sub-20 damage-spell penalty (a tooltip read on a level 25 to 30 character); Arcanite Champion's proc rate; Skyborne offsets; the `rt` helper's empty-tree parse.
- **Mage:** the Fire leveling build at 60 (+2.3% bare with the raid build's talents; needs the leveling order re-walked).
- **Guide lane:** feral-bear.md:83 "Swipe moves to the top on a pack" holds only in a raid at 60 on three or more targets; the Subtlety Death Mark row wording; the Holy paladin and Hunter rank SHOULDs in the lane reports.
- **Measure at refresh:** every spec's BiS and weights on 3% raid crit; rogue BiS with one-hand axes; Fire, Affliction and Assassination raid builds; Retribution gear under the paced twist.
- **Owner list:** Defensive Stance's tooltip needs data.yml's spelltips step (data.yml is disabled; not turned on); the Warcraft Logs key; level-60 trainer book ranks (Revenge 6, Battle Shout 7, Heroic Strike 9, Backstab 9, Eviscerate 9, Feint 5, Deadly Poison V, Greater Heal 5, Holy Light 9, Renew 10).

## Rendered cards after the regen (wow-player)

Twelve cards changed this round were captured at 1440 on int/r22 after the regen. Fire, Affliction, Assassination, Subtlety, Enhancement and Survival: SHIP on the first pass (Subtlety's Death Mark row wording is a SHOULD for the guide lane). Holy priest, Holy paladin and Discipline took two to three passes: the card printed "Rank 4 is your filler until 33" at 60 (note rewritten), the raid card showed party Flash of Light rank 4 and merged rows with different conditions (platform V1 and V2), and the mana-short rows led with the action instead of the condition (rewritten). Final verdicts: all SHIP. Checked on the level-60 view only; the lower-level changes were read against the guide text.
