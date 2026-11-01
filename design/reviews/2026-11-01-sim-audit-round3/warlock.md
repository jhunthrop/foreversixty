# Sim-accuracy round 3: warlock (2026-11-01)

Specs: Affliction, Demonology, Destruction. Client 1.60.1.70707. Pin c8d29ea33; lane engine `lane/r3-warlock` (fork); site branch `wt/r3-warlock`.

Method: spec-breakdown, horde raid preset, band 60, troll, 10000 iterations at seeds 1 and 50001. Affliction is measured on its guide `raidBuild` (`2552010013500105/00052/0550004`), because the published BiS json still carries the pre-round-2 build (nightly-owned, not regenerated here). Talent searches: `sim/cmd/talent-search -preset raid -faction horde -keep <raid-keeps.json>` on the lane engine; the reports are in `warlock-ts/` beside this file. Their header prints the pin sha, but GOWORK pointed them at the lane engine. Client rows are from wago SpellEffect for build 1.60.1.70707. Forever tooltips come from `nether.wowhead.com/forever/tooltip/spell/<id>` (read 2026-10-31).

## Pet scaling: sources and the figures taken (X2)

The client's Warlock Pet Scaling 416189 (flag "Owner Power Scaling") has 20 effects, all with base 0. They are aura 230/29 (Stamina, Intellect), 99 (melee AP), 13 misc 126 (damage done, all magic), 22 (five resistances), the physical-resistance (armor) row, 54/55 (hit), 319/65 (haste), 52/57 (crit), 85 (mana regen) and 135 (healing). The server computes every amount. Hunter Pet Scaling 415429 has the same list. No Forever source publishes a warlock figure.

| Channel | Taken | Sources | Single source? |
|---|---|---|---|
| Stamina | 75% | wowsims/sod `sim/warlock/pet.go` (378fec876, 2024-04-17); warcraft.wiki.gg Warlock minion page (Classic). TBC's 30% is not taken: Forever is Classic. | no |
| Intellect | 30% | wowsims/sod; warcraft.wiki.gg | no |
| Armor | 35% | wowsims/sod; TBC 35% | no |
| Resistances | 40% | TBC only. The client has the channels. Sims never use it. | **yes** |
| Attack power | 56.5% of spell power | wowsims/sod 70f6b1327 (tested 56.5%); warcraft.wiki.gg and TBC say 57% | no |
| Spell damage | 15% of spell power | wowsims/sod 61102b576; warcraft.wiki.gg; TBC | no |
| Hit | 100% of the warlock's Hit stat (Suppression included) | wowsims/sod only. Its own comment says WotLK testing found Suppression does not count. | **yes** |
| Crit | not inherited | wowsims/sod says 100%. Kargoz says "exactly like hunter pets", and hunter pets take 100% (Warcraft Tavern). But the Forever logs contradict it (see the next paragraph). | OPEN |
| Haste | not inherited | wowsims/sod inherits cast speed. This engine keeps haste as speed multipliers, not stats. | OPEN |

"Spell power" here is all-school spell power plus spell damage plus the higher of Fire and Shadow power, chosen once per iteration so the map stays linear (`pet.go:243`). The Forever texts agree that pets scale, but not on how much. The Blizzard warlock deep dive (2026-10-07) says "Pets now scale from their summoner's stats". Kargoz (kami-labs.fr, 2026-10-25) says "exactly like hunter pets" and gives a pet share of "around 35%". Warcraft Tavern's Forever warlock guide (beta 69913) says "Demons now scale with your gear". None gives a number, so the Classic set is used.

Why crit is not inherited: in the 29 foreverlogs.gg warlock logs (levels 13 to 30), Imp Firebolt crits 2.2% of 2,383 hits and demon melee crits 1.4% of 2,432. The warlock's own Shadow Bolt crits 5.6% and its DoTs 6%. A 100% share would push both pet figures above the warlock's. Lash of Pain crits 7.1% of 324, which fits either model. The cross-check against hunter pets in the same logs could not be run: foreverlogs.gg's per-character endpoints returned 403.

## Affliction

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| A1 | Raid search on the lane engine. The only winners beyond error drop Grim Reach, an unmodeled guide talent. The best variant that keeps every unmodeled talent (1 Improved Drains to Improved Bane of Agony) gains +0.7% at 10000 iterations (591.2 to 595.3 at both seeds), below the 1% bar. The build holds. | `warlock-ts/warlock-affliction-raid.md`; runs `after-aff-iboa-*` | Aligned | - |
| A2 | Re-check of round 2's fixes. The Wrack channel holds (Wrack stays out of the APL). Nightfall triggers from Corruption, Drain Life, Drain Soul and Wrack (`talents.go:290`). The Shadow Mastery masks match 18271. Siphon Life crit is unchanged (owner hold). | client 18094, 18271; engine `talents.go` | Aligned | - |

## Demonology

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| D1 | Raid search with inheriting demons: the guide build is still the best (658.5 ± 1.5). Master Demonologist, Unholy Power, Soul Link, Demonic Knowledge and Demonic Pact all keep their points. | `warlock-ts/warlock-demonology-raid.md` | Aligned | - |
| D2 | Re-check of Decimation's gate, cast time and cost: unchanged. Soul Fire casts only under the buff (37% uptime). | run `after-demonology-1` | Aligned | - |
| D3 | Demonic Brand (3 ranks: the demon's next 6 attacks add `(60-26)*1.5+14..17 + 0.078 x Shadow SP`) is unmodeled. No guide raid build takes it. | talents json 1.60.1.70707 | OPEN (moot for the raid builds) | - |

## Destruction

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| S1 | Held round-2 build `2550000002/0/2053345103101351` (Pandemic 2 and Cataclysm 3 for Unholy Power 5), re-run with inheritance. The raid search still picks it (+1.8% at 800 iterations). At 10000 iterations: 676.3 to 686.8 at seed 1 (+1.55%), 676.1 to 686.8 at 50001 (+1.58%). It keeps Destructive Reach and Molten Skin, Destruction holds 37 points, and the APL casts nothing it drops. It passes the adoption rule. The commit message's "+1.9%" is the round-2 figure; the figure now is +1.6%. | `warlock-ts/warlock-destruction-raid.md`; runs `after-destro-newbuild-*` | CONFIRMED (build) | S ab9433481 (cherry-pick of 5d013a9c0) |

## Cross-cutting (pets)

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| X2 | **Demons inherited nothing.** They now take the Classic shares in the table above. | client 416189; wowsims/sod; warcraft.wiki.gg; deep dive 2026-10-07. Test `pet_inheritance_test.go` | CONFIRMED-BUG (amounts: Classic figures, see the table) | E 94a45a77b |
| X3 | Demons refuse player buffs. Cherry-picked from dfbb63ae3, with its test `pet_buffs_test.go`. It now lands together with X2. | round 2 X3 | CONFIRMED (with X2) | E e3bdac097 |
| P1 | **Lash of Pain dealt Era's amounts (rank 6: 99).** Forever halved them to 16/22/30/36/43/50. The coefficient (0.429) and the 12 s cooldown match. | SpellEffect 7814-11780 base 16..50; Wowhead Forever 11780 "causing 50 Shadow damage"; Kargoz "16 (33 in Classic)" | CONFIRMED-BUG | E f6d77f02a |
| P2 | Firebolt dealt a flat 44. The client roll for rank 7 at level 60 is 42.6 to 47.8 (centre 44 at level 58, 0.1136 wide, +0.6 a level). The coefficient 0.571 matches. | spellconst/warlock.json 11763; Wowhead Forever "45 to 50" (rendered at a higher level); Kargoz "42 to 48" | CONFIRMED-BUG | E f6d77f02a |
| P3 | **Unholy Power reached only the demon's main-hand swing** (and Lash of Pain by copy), never Firebolt. Its text is "all damage done by your ... pets", and the Firebolt tooltip applies it (x1.1). It is now one pet-wide multiplier. | talents json 18769; Wowhead Forever 11763 | CONFIRMED-BUG | E f6d77f02a |
| P4 | **No demon sheet at levels 20 and 30.** The sheets exist only at 25/40/50/60, so at bands 20 and 30 a demon had no stats, mana or weapon and dealt 0%. Firebolt and Lash also fell back to rank 1 outside those four levels. The sheet is now interpolated by level (unconfirmed amounts; Forever's creature tables are server-side), and each ability uses the top rank its level has learned. | `pet.go:280`, `pet.go:308`; tests `pet_levels_test.go` | CONFIRMED-BUG | E d1e2536b9 |
| P5 | Kargoz says Lash "benefits from 42.9% of your spell power", which would be the warlock's own spell power. The client gives Lash only an SP coefficient and no owner flag, and Wowhead lists "SP mod 0.429". With the 15% share, the band-30 Lash averages 32 to 33, against 30 to 31 in the level-28-plus logs. | Wowhead 11780 details; logs | OPEN (single source; a level-60 Lash hit against a known warlock spell power would settle it) | - |
| P6 | Pet hit includes Suppression's 5%. At band 60 Lash of Pain misses 10.4% and demon melee 2.4%, against 17% and 8% at the pin, where the demon had no hit. | wowsims/sod only | OPEN (single source) | - |

## Before/after (relative, horde raid, band 60, 10000 iterations)

| Spec | Step | Seed 1 | Seed 50001 |
|---|---|---|---|
| Affliction | demons refuse buffs (E e3bdac097) | -10.8% | -11.0% |
| Affliction | + inheritance (E 94a45a77b) | +5.1% | +5.2% |
| Affliction | + Lash, Firebolt, Unholy Power (E f6d77f02a) | -0.8% | -0.8% |
| Demonology | buffs / inheritance / spells | -10.1% / +5.1% / -0.9% | -10.1% / +5.2% / -0.9% |
| Destruction | buffs / inheritance / spells | -9.6% / +4.3% / -0.6% | -9.5% / +4.2% / -0.7% |
| Destruction | + adopted raid build (S ab9433481) | +1.55% | +1.58% |
| **All steps** | Affliction / Demonology / Destruction | **-7.0% / -6.3% / -4.8%** | **-7.1% / -6.3% / -4.8%** |

The level fix (E d1e2536b9) does not move band 60. Absolute DPS at seed 1, for the integrator: Affliction 635.9 to 591.2, Demonology 702.4 to 658.1, Destruction 721.7 to 686.8. The order is still Destruction, Demonology, Affliction. SixtySim and MythicSim rank Demonology first.

**Pet share.** Band 60 raid: Affliction 21.2 to 15.3%, Demonology 20.4 to 15.0%, Destruction 18.6 to 11.7% (the new build drops Unholy Power). The pin's share came from player buffs (Windfury, Blessing of Might and others), not from scaling. Kargoz's "around 35%" is a leveling figure, Warcraft Tavern gives none, and it stays far above the sim.

**Low bands against the logs** (bare preset, seed 1, 10000 iterations). At band 20 the pin gave the pet 0% in every spec, and at band 30 also 0% (no mana to cast). The lane gives 20.6 / 26.0 / 22.5% at band 20 and 18.1 / 22.3 / 19.0% at band 30 (Affliction / Demonology / Destruction). The 29 logs show 1 to 30%, and the eight Succubus logs have a median of about 14%. The logs include AoE trash, where Hellfire alone deals most of the damage. Per hit the match is close: demon melee averages 35 to 36 in the sim at band 30 against 34 to 45 in the logs, and Lash of Pain rank 2 averages 32 to 33 against 30 to 31. Lash crits 6.4 to 6.8% against 7.1% in the logs. Demon melee crits 2.9% against 1.4% in the logs, a gap left by the demon's own Agility crit.

## Still open

- Pet crit and haste inheritance (X2 table), pet hit and Suppression (P6), and Lash scaling from the warlock's full spell power (P5): each needs a level-60 log with the warlock's sheet. The foreverlogs.gg hunter-pet cross-check is blocked (the endpoints return 403).
- The demon sheets at levels 20 and 30 are interpolated, not sourced (P4).
- Demonic Brand is unmodeled (D3). Siphon Life tick crit stays held for the owner.
- The published BiS json still carries Affliction's pre-round-2 raid build; the nightly regenerates it. sim/request TestRotationLadder goldens are not regenerated, as the brief requires.
