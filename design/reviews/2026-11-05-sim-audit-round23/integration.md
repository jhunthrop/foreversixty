# Sim-accuracy round 23: integration (2026-11-05)

Round 23 started from the landed round 22 (site 25432c0f1, fork bff826207) on the live client 1.60.1.70790. Lanes, at most two at once: platform, druid, paladin, hunter, warlock, mage, priest, rogue, shaman, warrior, items. The owner switched to build mode during the round (no wow-player gate, no full local suites). A standalone fix (land/addon-builds-fix, main b40f57073) landed mid-round and is merged in, with guides round 25.

**Not a zero round: 19 confirmed bugs, all fixed** (round 22: 33, round 21: 41, round 20: 31).

## Raid crit sourcing (round 22's change, moves every spec 1 to 4%)

**Stated by Blizzard, not inferred.** Blizzard's Forever notes thread (https://us.forums.blizzard.com/en/wow/t/2360696), post #9, posted 2026-11-03 03:20 UTC (2 Nov Pacific; the thread title reads "Updated November 2"):
- "Moonkin Form: The critical strike bonus from this talent is again an increase to all types of critical strikes and is again exclusive with Leader of the Pack."
- "Leader of the Pack: The critical strike bonus from this talent is again an increase to all types of critical strikes and is again exclusive with Moonkin Aura."

These reverse the 16 Oct lines that had split the two into spell crit and melee/ranged crit. The 70790 client rows agree on the amount and type: SpellEffect 701335 (Moonkin Aura 24907) and 702727 (Leader of the Pack 24932) are both a raid area aura of aura 290 (crit of every kind), base 3; hotfix push 113283 re-sends the same values in the Forever region. A client table cannot record exclusivity, so the client neither confirms nor contradicts that part. Wowhead's Forever tooltips say 3% for each and nothing about stacking. Warcraft Tavern's Forever guide calls them exclusive; Conquest Capped's says they stack but was last updated on 2 Nov, before the notes. The one inferred step is that a raid with both druids gets exactly 3%: both auras are 3%, so whichever applies gives the same number. No other aura-290 buff is in the raid preset. Pets: the client flags both auras "only target players"; the hunter pet gets the 3% once through the hunter's crit, and warlock demons get none of it (they take only the warlock's gear crit, below).

## Confirmed bugs

| Lane | Confirmed | Origin | What |
|---|---|---|---|
| priest | 6 | 6 old | Rank rows right after a new rank arrives: Shadow Word: Pain 26-29, Devouring Plague race rows, Mind Blast rank 4 at 34-36, Shadow Word: Death rank 2 from 48; Discipline shield rows; Holy Undead Flash Heal rank 2 at 32 |
| platform | 2 | 1 old, 1 R22 | A number times a duration in an APL crashed or used nanoseconds; the site's talent parser read a build ending in an empty tree as the leveling build (Assassination's raid build: addon raid order, dps-per-point, the BiS ranker's raid pass, talent-search). Also built: the APL value manaIncomePerSecond |
| mage | 2 | 2 old | Fire's level 60 build and its leveling order (+2.3 to +3.8% bare 60); Arcane raid-order Presence of Mind onto Arcane Wrath at 32-53 |
| warlock | 2 | 2 old | Demons take the warlock's gear crit (forever-bugs #331, single source); demon health 2 per Stamina (3 Nov) |
| items | 2 | 2 old | Robe of the Archmage shares no cooldown with mana gems or runes; Ankh of Life shares the 2-minute consumable cooldown |
| druid | 1 | 1 old | Wolfshead Helm's energy comes on Shifting Power, not on a shift (8 Oct) |
| paladin | 1 | 1 old | Retribution Consecration on a mana budget built on manaIncomePerSecond (166 gains, no losses, level 20 to 60), with the race rank rows re-swept |
| hunter | 1 | 1 R22 | The pet re-took the scope's crit on every hunter stat change |
| rogue | 1 | 1 R22 | Assassination's Rupture line lost 1.5% on round 22's raid build; removed |
| shaman | 1 | 1 old | Card lines gated on a talent's aura started at the spell's own level (Restoration's Healing Wave after Nature's Swiftness from 10, Enhancement's five-stack bolt from 10); they start where the build takes the talent |
| warrior | 0 | | Rank sweep aligned; Rend and Thunder Clap attack power scaling (30 Sep deep dive) recorded OPEN on one beta measurement |

By origin: 16 old, 3 R22. Lane reports carry the evidence.

## Integration

- **Fork:** int/r23 = 3b4d6b054 (pushed). The fork suite's only failures are the rounding-only conformance rows on this arm64 host and `sim/web` (binary_dist), both predating the round; no fork golden moved.
- **Site pin:** 80339ec3a (engine 3b4d6b054), with `python -m pipeline simproto` re-vendoring apl.proto and apl_pb2.py for manaIncomePerSecond.
- **Standalone landing mid-round:** land/addon-builds-fix (main b40f57073): the parser fix, a guard test that every addon raid order ends on the guide's raidBuild and is no copy of the leveling order, and the regenerated Assassination raid order and dps-per-point. Affliction, Demonology and Fire were not affected.
- **Regen at the pin (de0dfa8bf):** simdb, raid-consumes, addon-builds (Fire level 60 order), dps-per-point, ladder goldens, addon-data and Data.lua, coach starters via persona_starter.py --all.
- **Integrator fix:** the Assassination smoke build no longer expects the Blood Splatter gate the removed Rupture line read.
- **Checks run (build mode, touched only):** sim request, rankchoices, leveling, addon-builds; data test_apl, test_addonrotation, test_addondata; addon-data --check; luac; apl-check. CI runs the rest.
- **Not regenerated:** the BiS. The queued re-rank must run on this pin: 3% raid crit (round 22), Assassination's raid build now read correctly, the Fire level 60 build, demon crit, the robe and Ankh cooldowns.

## Before/after, raid 60

Before: round 22's landed engine and site (`before.tsv`). After: int/r23 at the pin. Same frozen BiS, guide raid builds (read with the fixed parser), 10,000 iterations at seeds 1 and 50001, both factions. Only rows that moved are listed.

| Spec | Faction | Seed 1 | % |
|---|---|---|---|
| rogue-assassination | alliance | 780.6 to 792.8 | +1.56 / +1.56 |
| rogue-assassination | horde | 778.8 to 790.2 | +1.46 / +1.57 |
| warlock-demonology | alliance | 715.6 to 720.4 | +0.67 / +0.67 |
| warlock-demonology | horde | 705.6 to 708.0 | +0.34 / +0.28 |
| warlock-destruction | alliance | 778.3 to 780.8 | +0.32 / +0.49 |
| warlock-destruction | horde | 774.7 to 776.3 | +0.21 / +0.22 |
| mage-frost | horde | 610.9 to 611.6 | +0.11 / +0.00 |
| mage-frost | alliance | 619.2 to 619.8 | +0.10 / +0.16 |
| paladin-retribution | alliance | 772.9 to 772.7 | -0.03 / +0.00 |
| mage-fire | alliance | 725.0 to 724.4 | -0.08 / -0.03 |
| mage-arcane | alliance | 711.5 to 710.8 | -0.10 / -0.08 |
| mage-fire | horde | 712.4 to 711.2 | -0.17 / +0.01 |
| mage-arcane | horde | 698.8 to 697.5 | -0.19 / -0.09 |
| druid-restoration (HPS) | alliance | 927.3 to 918.6 | -0.94 / -0.94 |
| druid-restoration (HPS) | horde | 923.9 to 915.1 | -0.95 / -0.96 |
| priest-holy (HPS) | horde | 788.0 to 766.0 | -2.79 / -2.67 |

## Open, for round 24

- **Live logs:** none anywhere (site API only the 20 Oct report; foreverlogs.gg API 403; the Warcraft Logs key 401: owner). Each lane's "Still open" names the log it needs. Single-source items to re-check on a log: demon crit from gear (forever-bugs #331), Rend and Thunder Clap attack power scaling (one beta measurement), Demonic Brand damage.
- **Platform:** a "mana spent per second" value for Protection paladin's Consecration gate; the leveling-order code spends the main tree before an off-tree point (Fire's Wand Specialization lands at 44 to 45, -0.2 to -0.6% at 41 to 43); Enhancement's Fire Nova row still shows from 12 (Clearcasting aura is not a talent's own spell); cost reductions add or multiply (mage and priest).
- **Guide lane:** protection.md:69 "Your biggest button" (Revenge hits 16% harder than Shield Slam at raid 60); feral-bear.md:83 Swipe on packs; an optional "does not stack" line on the crit-aura guide lines; the hunter R24-3/R24-4 note text.
- **Measure at refresh:** every spec's BiS and weights on this pin (see above), Fire bare-band BiS, Demonology and Destruction crit weights.
- **Owner list:** the Warcraft Logs key; level-60 trainer book ranks; Defensive Stance's tooltip needs data.yml's spelltips step (data.yml disabled).
