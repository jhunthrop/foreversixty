# Sim-accuracy round 17: integration (2026-11-04)

Round 17 started from the landed round 16 (site d7fc51fe7, fork pin f637ec32c) on client 1.60.1.70755 with the late hotfix re-layer. One lane per class plus platform. The lane reports sit beside this file.

**Not a zero round, but the count keeps falling:** about 77 bugs in round 15, about 43 in round 16, about 15 in round 17. Round 18 is needed.

## Landed ahead of the round: the gear score (alone, 3ad1bedf4)

The addon and the API multiplied item rating points by per-1% weights, so every rating stat scored 5 to 15 times too high and upgrade order flipped in every slot with rating gear. For example, Primal Batskin Gloves showed 647 and ranked first for Fury, when the correct score is 71. The fix converts rating to percent at score time through one table (docs/contracts/fixtures/rating-per-percent.json), and Lua, Go and TS are each pinned to it with a real-example test. Addon rating haste now feeds casters' spell haste. It shipped in addon 2.0.24.

As a side effect, every player's Forever Score gear component shifted toward correct, because the full-gear denominator dropped (e.g. warrior-arms from 5344 to 995).

## Branches and regen

- **Site `int/r17`:** every lane, plus origin/main through 2f3ff2a08.
- **Fork `int/r17`:** at **b2c4a2d05**, pushed. Fork master is not fast-forwarded until landing.
- **Pin:** d06dac7b6.
- **Regen once at the pin (38b890465):**
  - 8 ladder goldens;
  - addon-builds, addon-data and Data.lua;
  - spellicons;
  - coach rules and starters;
  - dps-per-point.
- **Not regenerated, left to the data job:**
  - **spelltips:** the Rebirth/Revive `$q` fill needs the current layered raw. The only local raw predates the late re-layer and reverted the druid tooltips, so I reverted that regen.
  - **items/<class>.json:** for the expertise line.
- **Checks:**
  - apl-check is clean;
  - web vitest passes in full (8738), with astro check and prettier clean;
  - Go: vet, plus the sim packages and the api sims package;
  - data: test_apl and test_addonrotation pass.
- **Stale test fixed by the integrator (ee71342d2):** the API Fury rotation test now reads round 16's per-level addon bands. Level 25 now correctly carries Execute.

## Before/after, raid 60

Before: round-start engine f637ec32c on site d7fc51fe7. After: int/r17 at the pin. Both use the same frozen published BiS and each side's guide raid builds, at 10,000 iterations, seeds 1 and 50001, both factions. The table lists specs moving more than 1%: seed 1 value, then % at seed 1 / 50001.

| Spec | Faction | Seed 1 | % |
|---|---|---|---|
| hunter-beast-mastery | alliance | 730.6 to 782.8 | +7.14 / +7.17 |
| hunter-beast-mastery | horde | 728.0 to 779.6 | +7.09 / +7.09 |
| rogue-subtlety | alliance | 751.3 to 765.4 | +1.88 / +2.01 |
| rogue-subtlety | horde | 746.0 to 760.1 | +1.89 / +1.89 |
| warlock-affliction | alliance | 736.7 to 753.1 | +2.23 / +2.25 |
| warlock-affliction | horde | 733.1 to 750.6 | +2.39 / +2.37 |
| warlock-demonology | alliance | 717.1 to 734.1 | +2.37 / +2.37 |
| warlock-demonology | horde | 707.2 to 724.3 | +2.42 / +2.45 |
| warlock-destruction | alliance | 780.2 to 795.3 | +1.94 / +1.91 |
| warlock-destruction | horde | 776.4 to 792.2 | +2.04 / +2.01 |

- **Beast Mastery and the warlocks:** pets and demons now take the raid hunter's Heart of the Lion. The 70755 client does not flag that aura "can only target players", unlike every other group buff. Beta logs show the similarly unflagged Trueshot Aura landing on hunters' pets.
- **Subtlety:** the dead Improved Ambush point moved to Dirty Deeds.
- **Every other spec** is within 1%. Below 60 the moves are:
  - Assassination and Subtlety Backstab before Mutilate/Hemorrhage (bare 20 +9.7%);
  - Restoration shaman tank Healing Wave rows at 38-46;
  - hunter Aimed Shot rank 1 rows;
  - Affliction leveling Malevolence over Nightfall (+1.5% at 40-60);
  - the Arcane band-20 Missiles line;
  - Discipline's pre-pull Renew.

## Confirmed by lane

| Lane | Fixed | Main ones |
|---|---|---|
| platform | 4 | the gear score (landed alone); Rebirth/Revive tooltip `$q`; expertise on item tooltips; per-band `lines` support for rank rows; `/statistics` stored on import |
| hunter | 2 | the pet never took its hunter's Heart of the Lion; Aimed Shot rank 1 rows at 28-59 |
| warlock | 2 | Affliction leveling Nightfall dead from 40; demons take Heart of the Lion |
| rogue | 2 | no Backstab before Mutilate/Hemorrhage; Subtlety's dead Improved Ambush point |
| shaman | 2 | tank Healing Wave rows at 38-46 (R16); Enhancement hit weight 0 past the weapon cap (weights only) |
| mage | 2 | Arcane "Fire points" text reverted by the round 16 merge (R16); Missile Barrage line without its talent check |
| priest | 1 | Discipline pre-pull Renew |
| paladin, druid, warrior | 0 | aligned |
| **total** | **about 15** | |

## Calls decided this round

| Call | Decision | Reason |
|---|---|---|
| Oath of Light and the low-rank penalty | Holy Light rows stay as modelled; the owner check goes on the post-launch list | The rows matter only from level 47, so there is no day-one risk. Per-band `lines` now makes the hedge or fix a one-line row edit. A hedge would cost -0.8% at 47 up to -17% at 59 if Oath is not penalized. |
| Mage Fire toll points (Improved Channeling 5, Arcane Resilience 3) | Keep | They are the required path to Arcane Meditation, and no swap gains. |
| Combat Eviscerate at four | Keep five | Four loses at bands 20-40 and gains only on top gear; the Forever guides say five. |
| Heroic Strike as a white swing for procs | Keep the engine (it counts) | The newer, log-based Season of Discovery evidence is from the same server line. If wrong, Fury loses about 18 rage a fight. |

## Owner's post-launch list

1. **Oath of Light** (needs a level-25+ Holy paladin with Oath; not day one). Cast top-rank and rank-1 Holy Light on the Oath target and on someone without it, and read the non-crit amounts. Expected values are in paladin.md. Then edit the Holy Light rows to suit.
2. **Low-rank penalty parameters:** the top-rank exemption, damage spells, levels above 30. The test is in round 16 priest.md, section 1.
3. **Summon Hawk:** ten swings at 60 on a level-60+ target. Both beta estimates now project above the 20-damage threshold, so it will probably need modelling.
4. **Pet buff bar:** check that Heart of the Lion shows on the pet in game.
5. **`/dump GetItemStats("item:23072")`:** shows whether the addon sees Fists of the Unrelenting's equip hit/crit.
6. **A naked level-60 Skyborne:** /statistics or `/dump UnitStat`, for the base stats.

## Open, for round 18

1. Everything on the owner list above.
2. **Log-dependent items, once level-60 Forever logs exist** (none yet: foreverlogs has only empty test reports):
   - Exorcism/Holy Shield crit, Death Mark storage, Mutilate/Shank off-hand;
   - Heroic Strike procs;
   - Eureka! on Incinerate, demon crit, Soul Siphon;
   - Eclipse on a missed Wrath, the cat form spread.
3. **Per-band talent lists:** each guide's per-band talent list shows the band's starting build, so a level-36 reader misses points taken at 34-35 (all guides; platform).
4. **iOS/Android item tooltips:** they need the expertise line.
5. **Arcane opener card:** it says Frostbolt 3 s before the pull, while the prose says 2.5 s.
6. **Fork test tags:** the "missing item" fork tests (TestP1Mage, the hunter sharpening-stone and ranged-hit tests) need `-tags=with_db`. That is a test-invocation note, not a bug.
