# Sim-accuracy round 28: integration

Scoped round in ship mode: the leveling openings of six specs, by round 27's method. Each leveling order spent points on talents that do nothing in the fight while a live talent in the same row waited. Bases: site int/r28 from main eb22d9382, fork int/r28 from 978e4c8aa. Engine pin after: **276ecc361**. No engine code changed; the fork carries only the synced rotation notes.

## Confirmed bugs: 17

| lane | fixed | by origin | report |
|---|---|---|---|
| casters | 11 | 11 old | casters.md |
| hybrids | 6 | 6 old | hybrids.md |
| **total** | **17** | **17 old** | |

What changed, in a player's words:
- **Destruction:**
  - Bane at 10 to 14 before Destructive Reach and Molten Skin.
  - Shadowburn at 20.
  - Conflagrate at 31 with Fire and Brimstone right behind it, at 32 to 34.
  - Shadow and Flame at 35 to 38.
  - Two rank rows re-cut.
  - Gains: +1.0 to 3.4% at 10 to 14, +5.3% at 30, +14% at 38.
- **Demonology:**
  - Unholy Power at 10 to 14 and Fel Vitality at 15 to 17, before Improved Health Funnel and Demonic Embrace.
  - Demonic Knowledge before Decimation's second point (+3.2% at 31 to 33).
  - A Gnome Shadow Bolt rank row at 22 and 23.
- **Frost:**
  - Precision at 15 to 17 before Frost Warding.
  - Frost Channeling at 20 to 22 (+8% at 22 to 25).
  - Fingers of Frost at 30, six levels sooner (+11 to 14.6% at 31 to 35).
  - Frostbolt rank 4 rows set per race.
- **Feral:**
  - Sharpened Claws at 23 and 24 before Savage Fury.
  - Leader of the Pack, Blood Frenzy, Natural Instinct and Rend and Tear at 31 to 38, before Predatory Discipline.
  - Feral Swiftness before Feral Instinct is order only: below 20 the druid fights as a caster.
- **Survival:**
  - Improved Tracking is live (+1% damage a point), so the seed is aligned.
  - Predator's Edge at 25 to 29.
  - Resourcefulness, Expose Prey, Survivalist's Discipline and Lightning Reflexes at 31 to 37, before Deterrence and Counterattack, which need the target to swing at you.
  - The Raptor Strike rank 1 row moves to 31.
- **Protection paladin** (judged on threat):
  - Divine Intellect at 10 ties, so the seed is aligned.
  - Shield Specialization at 21 and 22 (+2.3 to 3.2%).
  - Swift Judgement at 25 (+5.2 to 6.2%).
  - Band 20 keeps Templar's Bulwark.

## Before/after

Only the touched specs, level by level, at 10,000 iterations, seeds 1 and 50001, both factions. The tables are in casters.md and hybrids.md. Raid 60 does not move.

## Regenerated (S 655959b33, regen790.sh at pin 276ecc361)

Regenerated: the Feral, Survival, Frost, Demonology and Destruction ladder goldens, addon-builds, dps-per-point, the Survival coach rules and the build and rotation rows in Data.lua. No BiS pick moves. `test_addonrotation`'s Survival Counterattack pin follows the order to 39 (S 610d7eca0).

## Re-rank needed (one batched dispatch)

- **warlock-destruction:** bands 20, 30 and 40.
- **warlock-demonology:** band 20.
- **mage-frost:** bands 20 and 30.
- **hunter-survival:** band 30.

All four are in `knownStaleBISFiles` (S 8883a8965) until the batch re-rank.

## Open, for round 29

- Demonology D4: Master Summoner against Improved Sayaad in the leveling build is a build question.
- Survival's raid build still spends Deterrence and Counterattack.
- The band 20 tank boss kills a bare level 25 to 28 tank almost every time.
- Generator comments in `tools/coach/starter_overrides.py` name old levels: Conflagrate (201), Expose Prey, Counterattack and Deterrence (126, 673, 688), and Initiative (489). These go to the coach owner.
- The icon regen for the 54 engine-only ids, once wago.tools answers.
- Logs: still only the 20 Oct test report at 01:16 UTC 7 Nov.
