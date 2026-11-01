# Sim audit round 4: casters (Priest and Mage), 2026-11-01

Specs: priest-shadow, priest-discipline, priest-holy, mage-arcane, mage-fire, mage-frost.

## Method

- **Engine.** Lane branch `lane/r4-casters` from the pin `640da215f`. "Before" runs a second worktree at the pin against origin/main's site data.
- **Client.** wago.tools 1.60.1.70707 tables (SpellEffect, SpellMisc, SpellCooldowns, SpellCategories, SpellPower, SpellLevels, SkillLineAbility, SpellClassOptions, SpellAuraOptions, Spell), read 2026-11-01. `?hotfixes=1` returns the same SpellEffect rows. 70707 (2026-10-30) is still the newest `wowf` build.
- **Logs.** foreverlogs.gg reports 2740 to 2813. The beta closed on 23 October (report 2811, "Stocks as Beta Ends"); 2812 and 2813 are empty. No log exists for any build after 70532.
- **Runs.** `sim/cmd/spec-breakdown`, raid preset, band 60, 10,000 iterations, seeds 1 and 50001, both factions. Talent searches: `sim/cmd/talent-search -preset raid -faction horde -keep <raid-keeps.json>`.
- **Tests.** Every confirmed bug got a failing test first, seen failing, then passing.

## Cross-cutting

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| X1 | Item 1: a second source for Ice Lance's 8 s on every rank. | None exists. 70685 was built 2026-10-28, five days after the beta closed, so no log can show it. The shipped tables still give ranks 2 to 6 (400640, 1240044-1240047) Category 0 and no recovery; only rank 1 (1312002) is in category 19 with 8000 ms. `?hotfixes=1` changes nothing. Wowhead Forever (spell pages and tooltips, read 2026-11-01) shows "8 sec cooldown" on every rank, with no comments. Web search (Blizzard forums, Reddit, kami-labs, guides) has nothing either way. | OPEN, single source. Kept, because Wowhead's overlay is newer than the shipped tables. Settled in game: two Ice Lances within 8 s at rank 2 or higher. Costs about 0.8% if wrong. | - |
| X2 | Item 3: Power Infusion for healers. | `priest/power_infusion.go` puts it on the ally the player picks. The site passes an empty target (`sim/request/request.go:717`), so a healer's own number never counts it. The client row 10060 is +20% healing and +20% magic damage (auras 136 and 79), 15 s, 3 min. Raid practice is to give it to a damage caster. | Settled as policy, not a sim bug: published healer numbers leave out Power Infusion, because it goes to a DPS. One self-cast per 180 s would add at most about 1.7% (15 s at +20%). The owner can overrule. | - |
| X3 | Item 4: Frost Blast (1326053-1326055). | It is learnable: no skill-line row, but three beta mages cast rank 1 (reports 2795 and 2811). The engine does not register it. At 60 it is 77 + 0.6 a level over 48 base, 0.271 coefficient, 1.5 s (1.0 s with Improved Frostbolt, still on the GCD), 90 mana: about 40% of Frostbolt's damage a second. | Aligned. No published list would cast it, so it moves no number. Still unregistered. | - |
| X4 | Item 4: Violet Sorcerer's Vestments (set 2134). | Its pieces are items 273044, 273051 and 286987-286989 (sets.json). No published BiS set wears one, for any spec, band or preset (data/builds/1.60.1.70707/bis/*.json). | Aligned | - |
| X5 | Arcane Concentration's Clearcasting lasted its full 15 s. | Client 12536 has ProcCharges 1 and a -100% cost on the damage-spell mask. The engine's consume (`mage/talents.go:185`) checked the spell's *current* cost, which Clearcasting itself had already made 0, so the charge was never spent. Frostbolt cost 132 a cast against 221 (260 less Frost Channeling). It also paid for Ice Barrier, which is outside the mask, and Arcane Concentration ignored 11213's 1 s proc cooldown (ProcCategoryRecovery 1000). Tests: `mage/clearcasting_test.go` (3 tests, failed, then passed). | CONFIRMED-BUG (fresh re-audit; client row, and "your next damage spell") | E c6d0c1f46 |
| X6 | Arcane Blast's stack ramp replaced the other cost mods. | `setArcaneBlastCostMultiplier` set Cost.Multiplier outright, which dropped Arcane Power's +30% (12042 effect 1, misc 14) and, after X5, Clearcasting's -100%. It now adds. Test: `mage/arcane_power_cost_test.go` (200 against 230, then passed). | CONFIRMED-BUG (fresh re-audit) | E c6d0c1f46 |
| X7 | Missile Barrage replaced the other cost mods. | It set Arcane Missiles' multiplier to 0, then back to 100. After X5, a Barrage that came and went under Clearcasting left Arcane Missiles at double cost. It now takes 100 off and gives it back (400589 effect 2). Test: TestMissileBarrageAddsToTheOtherCostMods (200 against 100, then passed). | CONFIRMED-BUG (fresh re-audit) | E db121d9c0 |
| X8 | Arcane Missiles rank 8 (25345) is an Ahn'Qiraj book. | It is taught by Tome of Arcane Missiles VIII, which drops in Ruins of Ahn'Qiraj: warcraft.wiki.gg "Arcane Missiles", and wowclassicdb's AQ20 loot (npc 15369, 15348). It shares level 56 with trainer rank 7, the book signature. Arcane's and Frost's Missile Barrage lines cast it. Now: the engine gates it on core.IncludeAQ (test changed first, failed, then passed). The lists cast rank 7 (10212). `book-ranks.json` names it (test_spellranks floor failed, then passed). `spellranks.json` flags it, and the regeneration's only change is that row plus its manifest hash. | CONFIRMED-BUG (fresh re-audit, two sources) | E 7a380d6fe, S a081e3d6c |

## Priest: Shadow

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| P1 | Item 2: Shadowfiend learnability. | No new evidence. The beta closed 23 October, and no launch trainer data exists yet. In 70707, row 49886 still has no learn level (401977 SpellLevels BaseLevel 0). | OPEN (unchanged; revert S 09848d599 and E 72aefbb89 if a launch trainer shows it) | - |
| P2 | Fresh: Shadowform's +100% crit damage and -50% cost mask (15473 effects 3 and 4, mask 41984016/2). | Mind Flay, Mind Blast, Shadow Word: Pain, Devouring Plague and Shadow Word: Death are all in the mask. The engine's school-wide class-spell mod covers the same set. Touch of the Grave is not a class spell, so it is left out. | Aligned | - |
| P3 | Fresh: Mind Flay 18807 and Improved Mind Flay. | Client: 130 a tick, 0.167, 1 s period, 3 ticks. Improved Mind Flay 2/2 compounds +20% a tick (1225139 text), and `priest/talents.go:562` does the same. Shadow Focus 5/5 is +5% hit, and Mind Blast misses 10.3% against level 63 with about 1.7% gear hit. | Aligned | - |
| P4 | Raid talent search, lane engine. | Best +0.2% (±0.2%). | Hold | - |

## Priest: Discipline and Holy

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| H1 | Power Infusion. | See X2. | Policy | - |
| H2 | Fresh: Inner Focus consume. | It adds and removes 100 on Cost.Multiplier and is consumed by the next priest spell. The cooldown starts on use. It does not share X5's current-cost check. | Aligned | - |
| H3 | Healer numbers after the lane. | No priest engine or data line moved: 0.0% at both seeds and both factions. | - | - |

## Mage: Arcane

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| A1 | X5, X6, X7, X8. | Clearcasting (X5) now makes one cast free, not 15 s of casts. Missile Barrage's Arcane Missiles are rank 7. Arcane Power's +30% cost now holds across Arcane Blast stacks. | CONFIRMED-BUG | E c6d0c1f46, db121d9c0, 7a380d6fe; S a081e3d6c |
| A2 | Evocation (12051, trainable at 20, 8 min). | No mage list cast it. Arcane with Evocation below 5, 10 or 20% mana measured -0.6% to -1.0%, so it stays out. | Aligned (no gain) | - |
| A3 | Raid talent search after the fixes. | The guide is the best build found. | Hold | - |

## Mage: Fire

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| F1 | X5 left the Fire list out of mana 49 s into the 180 s fight. | Adding Evocation below 5% mana gains +17%. 5, 10 and 20% are within error of each other. Gating Fire Blast on the mana pace (-0.4%), dropping Fire Blast (-6%) and a paced Fireball filler (no casts) all measured at or below that. Test: `data/tests/test_apl.py` test_a_rotation_that_runs_dry_casts_its_mana_cooldown (failed, then passed). Fork copy synced. | CONFIRMED-BUG (consequence of X5) | S 98a9dad7e, c8619ad2d; E b5436192e |
| F2 | Raid build after X5. | Arcane Meditation 2/3 and Arcane Resilience 3/3 come from Accelerant 2, Magic Absorption 1 and one rank of Improved Fire Blast. Measured +5.2% (610.8 / 611.1 to 642.1 / 643.2 horde; alliance 622.1 / 624.2 to 649.8 / 650.8). It keeps every unmodeled guide talent, keeps Fire at 31 and drops nothing the list casts. The guide's raidBuild and prose are updated. A re-search from the new build finds +0.9%, under the 1% bar. | Adopted | S 65fa2ffb1 |
| F3 | Fresh: Heating Up and Master of Elements. | Heating Up: 25% a stack, 3 stacks, a crit from the feeder mask, and Fire Blast adds one on every hit (400624 text). Master of Elements: 30% of base cost on Fire and Frost crits. The engine skips a refund on a free cast, which is unverified for Forever. | Aligned (refund on a free cast is OPEN, single-source) | - |

## Mage: Frost

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| R1 | X5 and X8. | Frost lost the extra free Frostbolts and the rank 8 missiles. | CONFIRMED-BUG | as X5, X8 |
| R2 | Fresh: Ice Barrier's +15% / +5% rests on one source. | Ice Barrier (`mage/ice_barrier.go:27-28`) uses Wowhead's 15% (spell 1326891 "Frigid Power") and 5%. The shipped tables, with or without hotfixes, give 3 and 3 (1326891 effect 0; 11426 and 13033 effect 2). Kami-labs (22 October, Venruki) reported 3% + 3% on the 16 October build. Measured with 3/3: 543.9 to 494.0, -9.2%. | OPEN, single source (Wowhead's newer overlay is kept). In game: the Frigid Power buff tooltip, or a Frostbolt's hit with and without the shield. | - |
| R3 | Fresh: Ice Lance's coefficient. | The client column is 0 for every rank, and the engine uses 0.4286 (marked unconfirmed). Beta logs show spell power scaling: rank 1 non-frozen hits of 42-45 against a 34 ceiling (reports 2747, 2762, 2776), and rank 2 hits of 51-58 against 44 (2786). The value itself cannot be fitted without the casters' spell power. | Aligned that it scales; the 0.43 is unconfirmed | - |
| R4 | Raid talent search after the fixes. | Best +0.2% (±0.3%). | Hold | - |

## Before/after (raid, band 60, 10,000 iterations; pin and origin/main against lane engine plus site branch)

| Spec | Faction | Before (seed 1 / 50001) | After | Change | Cause |
|---|---|---|---|---|---|
| Arcane | alliance | 663.6 / 664.8 | 599.7 / 599.8 | -9.7% | X5, X6, X7, X8 |
| Arcane | horde (orc) | 665.7 / 666.4 | 600.9 / 600.7 | -9.8% | X5, X6, X7, X8 |
| Fire, published build | alliance | 714.7 / 717.0 | 622.1 / 624.2 | -13.0% | X5, F1 |
| Fire, published build | horde (orc) | 710.9 / 712.5 | 610.8 / 611.1 | -14.1% | X5, F1 |
| Fire, new raid build (F2) | alliance | 714.7 / 717.0 | 649.8 / 650.8 | -9.2% | X5, F1, F2 |
| Fire, new raid build (F2) | horde | 710.9 / 712.5 | 642.1 / 643.2 | -9.7% | X5, F1, F2 |
| Frost | alliance | 559.3 / 559.1 | 547.4 / 548.1 | -2.0% | X5, X8 |
| Frost | horde (troll) | 557.3 / 556.7 | 543.3 / 543.6 | -2.4% | X5, X8 |
| Shadow, Discipline, Holy | both | - | - | 0.0% | - |

Without the Evocation line, Fire fell to 521.5 (horde, seed 1).

## Still open

- Ice Lance's 8 s on ranks 2 to 6 (X1) and Ice Barrier's 15% / 5% (R2) rest on Wowhead's overlay alone. R2 is worth 9% of Frost. One in-game look settles both.
- Shadowfiend at launch (P1).
- Master of Elements on a Clearcasting cast (F3).
- Frost Blast is unregistered (X3); it would not be cast.
- Ladder goldens: the three mage rotations moved. TestRotationLadder is left for the integrator.
- `spellranks.json` now flags 25345. If data.yml regenerates before the pin, it writes the same row, because `book-ranks.json` names it.

## Sources

- wago.tools db2 CSVs, build 1.60.1.70707 (read 2026-11-01); build list https://wago.tools/api/builds
- Wowhead Forever tooltips (read 2026-11-01): https://nether.wowhead.com/forever/tooltip/spell/ for 1240047, 400640, 1312002, 13033, 11426, 1326891, 1326055; https://www.wowhead.com/forever/spell=1240047
- warcraft.wiki.gg, Arcane Missiles: https://warcraft.wiki.gg/wiki/Arcane_Missiles
- wowclassicdb Ruins of Ahn'Qiraj loot: https://wowclassicdb.com/tbc/npc/15369, https://wowclassicdb.com/tbc/npc/15348
- kami-labs, the hidden October 20 Frost hotfix (22 October 2026): https://kami-labs.fr/en/wow-forever/mage-givre-wow-forever-correctif-discret/
- Blizzard forums, "RIP Frost Mage" (16 October 2026): https://us.forums.blizzard.com/en/wow/t/2384052
- foreverlogs.gg public reports 2740 to 2813
