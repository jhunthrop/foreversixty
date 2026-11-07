# Sim-accuracy round 29: integration

Scoped round in ship mode. It finishes the leveling-walk sweep (round 27's method) over the 18 specs rounds 27 and 28 had not covered, and works round 28's open build items. With it, every spec's leveling order has been walked. Bases: site int/r29 from main 75d6264a5, fork int/r29 from 276ecc361. Engine pin after: **fe29dfbd2**. No engine code changed; the fork carries only the synced Survival rotation.

## Confirmed bugs: 23

| lane | fixed | by origin | report |
|---|---|---|---|
| walks-dps | 13 | 13 old | walksdps.md |
| walks-heal | 10 | 9 old, 1 R28 (a stale Survival rank test) | walksheal.md |
| **total** | **23** | **22 old, 1 R28** | |

What changed, in a player's words:
- **Every one of the 13 DPS and tank specs had a dead point ahead of a live one in its leveling order, and all are fixed:** Fury, Protection warrior, Arcane, Fire, Affliction, Beast Mastery, Marksmanship, Balance, Feral bear, Retribution, Shadow, Elemental and Enhancement. The biggest gains:
  - Enhancement's instant Lightning Bolt arrives at 39 instead of 42 (+24.6% at 39).
  - Fire takes Precision at 41 instead of 58 (+2.3 to 19.5% at 41 to 59).
  - Balance takes Heart of the Wild and Naturalist before Nature's Focus (+0.3 to 10% at 42 to 58).
- **Healers:** Restoration druid (+6 to 14% from 30), Holy paladin (+9 to 16% at 43 to 58), Holy priest, Discipline and Restoration shaman all stop spending early points on talents that do nothing for a healer. The Holy paladin and Restoration shaman rank rows that ride on those levels moved with them.
- **Survival raid:** Deterrence and Counterattack did nothing against a boss that never swings at the hunter. Their points go to Rapid Killing (+0.6 to 0.7% at raid 60), and the Counterattack rotation line is gone.
- **The tank fight between bands:** a tank between band levels fought a boss scaled to its own level in the lower band's gear. It now fights the band at or below its level, as the heal profile already did. Chance of death at 25 and 28 falls from 100% to 65 to 74%.
- **Demonology D4** (Master Summoner against Improved Sayaad) is aligned: Improved Sayaad gains under 0.5% and loses at 40 and 60.
- **Shared ladder change:** a reservation may now take a deeper talent in another tree once its rows and prerequisite are legal. Both lanes needed it, for Balance's Naturalist and Restoration druid's Nature's Splendor. The merge kept walks-heal's version, which checks full tree legality on each rank. Every leveling walk test passes on it.

## Ruling: tank leveling orders and the chance of death

The coordinator's ruling (the owner delegated product calls), recorded in the talent adoption rule: leveling tanks also quest solo and run dungeons with weak healers. A threat talent therefore goes ahead of a survival talent only when the sim fight's chance of death rises by less than 5 points absolute at every level the swap touches. Otherwise the survival talent stays until a level where the swap passes.

Round 29 applied it to its two tank orders:
- **Protection warrior:** Defiance at 20 to 22 and Master of Defense at 23 and 24 stay. Anticipation moves to 25 to 30 with Last Stand at 28, before Improved Sunder Armor. Shield Slam and Bastion return to their old levels. Threat +17 to 30% at 20 to 30 and +1 to 11% at 31 to 38; chance of death rises by 1.8 points at most. The first draft had gone from 2% to 14% at 35.
- **Feral bear:** the step that failed at 50 was Naturalist before Natural Reaction, not Thick Hide, which both versions take at 16 to 18. The order is now Natural Reaction, Heart of the Wild, then Naturalist, with Feral Swiftness's second point at 39 (level 40 would otherwise go from 81% to 89%). Threat +2 to 8% at 22 to 39; chance of death rises by 1 point at most.

## Before/after

Only the touched specs, level by level, at 10,000 iterations, seeds 1 and 50001, both factions. The tables are in walksdps.md and walksheal.md; some exploratory levels were confirmed at 3,000 iterations where the gain was above 2%. Raid 60 moves only for Survival (+0.6 to 0.7%).

## Regenerated (S 57de8313b, regen790.sh at pin fe29dfbd2)

20 ladder goldens, addon-builds, dps-per-point, the Survival and Fire coach rules, and the build and rotation rows in Data.lua. No BiS pick moves. `make apl-sync` changed only the fork's Survival copy, and `apl-check` passes on every spec.

The integrator also fixed tests that pinned levels the new orders deliberately moved (S a513a858f). Fire's Arcane Blast note named Missile Barrage a level before the player has it, and now says "Once you have Missile Barrage".

## Re-rank needed (one batched dispatch)

19 specs, all in `knownStaleBISFiles` (S e338874fd) until the batch re-rank:
- druid-balance, druid-feral-bear, druid-restoration;
- hunter-beast-mastery, hunter-marksmanship, hunter-survival (including bare 60 and raid 60);
- mage-arcane, mage-fire;
- paladin-holy, paladin-retribution;
- priest-discipline, priest-holy, priest-shadow;
- shaman-elemental, shaman-enhancement, shaman-restoration;
- warlock-affliction;
- warrior-fury, warrior-protection.

## Open, for round 30

- Rank rows on the moved levels were screened, not re-judged on the new talents (both reports list the cells).
- Below 60 the stand-in tank boss is 1.4 to 2 times harder than the dungeon table, so a bare paladin or bear at 30 to 39 still dies 76 to 100% of the time.
- Master Summoner's sim value rests on the pre-pull summon's mana coming out of the fight.
- Arcane would gain from Precision before 40 only with a build that spends fewer Arcane points.
- The bear's 30-rage Maul gate at 40 to 53 was tuned on the old order.
- Generator comments in `tools/coach/starter_overrides.py` still name old levels (coach owner).
- The icon regen for the 54 engine-only ids, once wago.tools answers.
- Logs: still only the 20 Oct test report at 03:10 UTC 7 Nov.
