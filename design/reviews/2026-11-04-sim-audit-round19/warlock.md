# Sim-accuracy round 19: warlock (2026-11-04)

Specs: Affliction, Demonology, Destruction. Client 1.60.1.70755 with the late hotfix re-layer. Site `wt/r19-warlock` from int/r19 (99ef61f0c); fork `lane/r19-warlock` from int/r19, with no fork change. Measured with spec-breakdown on `$R/bis-frozen`. Raid runs use each guide's raid build (`$R/rt`). Bare runs use the frozen band entry's talents.

**Totals: 1 CONFIRMED-BUG (old), fixed in S 6ce860617.** The coefficient lead is clean. Every log item is still open: at 19:12 UTC no launch log existed, and none at level 40+ can exist on launch night.

## Seed 1: rank rows at 40-59, level by level

Method: `-level` 40, 43, 46, 49, 50, 53, 56 and 59. At each level the top learned rank was set against the two ranks below it, by forcing every reference with `-rank` (casts and conditions). Bare preset, orc, 3,000 iterations, seed 1, as a screen (`$R/warlock/out/sweep1.txt`). Every near tie was then re-run at 10,000 iterations, seeds 1 and 50001, on both factions (`ties.txt`, `destro40*.txt`). The bar is any gain beyond error, not 1%.

**Each lower rank was cast under its own id** (integrator's trap): Bane of Agony 11711, Siphon Life 18879, Immolate 11665 and 2941, Corruption 7648 and 6223, Shadow Bolt 1106 and 11659, Soul Fire 6353, Conflagrate 17962, Shadowburn 18869, Incinerate 412758. No rank is unregistered.

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | **Destruction at 40 and 41 should cast a cheaper Immolate and Corruption.** <br>**What was wrong:** the rotation cast the newest ranks at 40 and 41. At 40, Incinerate, Conflagrate rank 3 and Immolate rank 5 all arrive together. On band-40 gear the fight then runs short of mana: about 33 Life Taps in 3 minutes, and 3.3 s out of mana at 41. <br>**Evidence:** 10,000 iterations, seeds 1 / 50001. <br>- Immolate rank 4 at 40: orc 169.5 / 169.5 → 169.8 / 169.7; human 167.9 / 167.8 → 168.4 / 168.3. <br>- Corruption rank 3 at 41: orc 170.7 / 170.5 → 170.9 / 170.9; human 169.0 / 168.9 → 169.6 / 169.6. <br>- Both rows at the same level are no better than one alone. At 42 the top ranks win again: Corruption rank 4 by 0.2 to 0.3%, and Immolate rank 5 ties or wins by 0.1. <br>**Fix:** rows Immolate rank 4 at 40 and Corruption rank 3 at 41 (`data/curated/rank-choices.json`). Test `TestWarlockDestructionCheaperDoTsAt40And41` was red, then green. The production rewrite reproduces the `-rank` numbers to the decimal. destruction.md "Levels 20 to 60" says so. wow-player: SHIP, and its wording was applied. | `out/destro40.txt`, `destro40a.txt`, `ties.txt` | CONFIRMED-BUG | old (round 17 X1 called Immolate at 40 a tie, at 3,000 iterations on one faction; it never tried 41) | S 6ce860617 |
| R2 | Every other ranked cast loses to the top learned rank at every level tried. The margins at 59, top rank against the next: <br>- Affliction: Corruption 1.5%, Bane of Agony 1.9%, Immolate 1.2%, Siphon Life 0.9%. Shadow Bolt is not cast from 40 (Wrack). <br>- Demonology: Shadow Bolt 2.2%, Corruption 1.7%, Agony 2.3%, Immolate 2.2%, Soul Fire 0.4%. <br>- Destruction: Corruption 1.3%, Agony 2.1%, Immolate 3.2%, Conflagrate 1.4%, Shadowburn 0.8%, Incinerate 2.4%. <br>The near ties: <br>- Destruction Shadowburn rank 3 against 4 at 40, 43 and 44: within ±0.2, never beyond error on both seeds. The same holds for Corruption rank 3 at 40 on orc (+0.1 at seed 1). <br>- Demonology Shadow Bolt rank 5 against 6 at 40 and 43: rank 6 wins by 0.1 to 0.4 on both seeds. | `out/sweep1.txt`, `ties.txt` | Aligned | - | - |
| R3 | **Pet spells.** The Imp's Firebolt and the Succubus's Lash of Pain take the top rank learned by level (Firebolt 1/8/18/28/38/48/58; Lash 20/28/36/44/52/60). Mana, cooldowns (Lash 12 s on a 1.5 s GCD) and cast time (Firebolt 2 s) match SpellPower, SpellCooldowns and SpellMisc. Pets have no rank choice. | `imp.go:85-135`, `succubus.go:105-150`; wago 70755 layered | Aligned | - | - |
| R4 | Shared watch (priest P1, low ranks proccing less): the new rows cast Immolate rank 4 and Corruption rank 3. The Destruction leveling build has no proc keyed to either spell (no Nightfall; Shadow and Flame procs from Conflagrate and Shadowburn). Nothing to flag. | build decode | Aligned | - | - |

## Seed 2: coefficients (Ice Lance lead)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| K1 | Every spell-power-scaled warlock and demon spell the engine casts has a non-zero client EffectBonusCoefficient on every rank, and the engine carries that value. <br>- Direct: Shadow Bolt .486/.629/.8/.857; Immolate .2 direct; Searing Pain, Conflagrate, Shadowburn and Lash of Pain .429; Incinerate .714; Soul Fire 1.0; Death Coil .214. <br>- Per tick: Corruption .2, Bane of Agony .133, Immolate .13, Siphon Life .05, Drain Life and Drain Soul .1, Wrack .286, Rain of Fire .083. <br>- Bane of Doom 4.0. <br>- Imp Firebolt .164/.314/.529/.571. <br>None of these reads a generator default. The "unconfirmed, vanilla convention" constants in `constants_auto_gen.go` that the warlock package uses are none of these (it reads only `WrackSpellCoeff`). The others belong to uncast or non-damage spells: Torment, Consume Shadows, Haunt, Unstable Affliction, the tier-1 bonuses. Firestone's proc carries 0, but warlocks do not melee. | `out/coef.txt` (wago 70755 SpellEffect, layered); engine `grep BonusCoefficient` | Aligned | - | - |
| K2 | Mana costs of every cast rank match SpellPower: Shadow Bolt, Corruption, Agony, Immolate, Siphon Life, Shadowburn, Conflagrate, Incinerate, Soul Fire, Searing Pain, Curse of the Elements, Wrack 200, Bane of Doom 300. | wago SpellPower; engine `manaCost` arrays | Aligned | - | - |

## Seed 3: launch logs

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| L1 | **No launch log yet.** At 19:12 UTC (before the 23:00 launch), the newest foreverlogs.gg `reports/public` entry was still 2811 (2026-10-24), and `api.foreversixty.gg/v1/reports/recent` held only z2opt2x2gs5u (2026-10-20). Characters start at level 1 at launch, so no level-40+ log can exist tonight. Several items need such a log: Eureka! on Incinerate (a level-40+ Gnome Destruction log), Soul Siphon's count and Wrack (40+), and Doom and demon crit at 60. All carried. | foreverlogs API; api.foreversixty.gg | OPEN | - | - |
| L2 | Warcraft Logs is blocked (gcloud login expired). Recorded and not worked around. | brief | recorded | - | - |

## Affliction (fresh re-audit)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | Raid 60 reproduces `before.tsv` to the decimal on int/r19 (with the druid, shaman, platform, mage and paladin merges): 753.1 / 753.6 alliance, 750.6 / 750.8 horde. | `out/raid60.tsv` | Aligned | - | - |
| A2 | The leveling build (`25522300132001351/200522/005`) decodes to points that each act on something: Grim Reach is range, Soul Harvest is regen after a Drain Soul kill, and the rest act on cast spells. Unchanged since round 18 S4. | `tools/dec.py` | Aligned | - | - |

## Demonology (fresh re-audit)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| D1 | Raid 60 reproduces `before.tsv`: 734.1 / 735.1 alliance, 724.3 / 724.0 horde. | `out/raid60.tsv` | Aligned | - | - |
| D2 | **Races** (guide round 15, carried request). Raid 60, guide raid build, frozen BiS, 10,000 iterations, seeds 1 / 50001. <br>- Alliance: Human 737.0 / 737.8, Gnome 734.1 / 735.1 (-0.4%). <br>- Horde: Orc 730.1 / 730.9, Undead 726.1 / 727.3 (-0.5%), Troll 724.3 / 724.0 (-0.8%). <br>The guide's order holds (demonology.md:110): Human and Orc, Gnome a step behind, Troll behind Orc, Undead close behind Orc. The published BiS race (Gnome; Troll on Horde) is within 1%, so it stays. The Horde race is pending owner (O-list). | `out/races.txt` | Aligned | - | - |

## Destruction (fresh re-audit)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | Rank rows at 40 and 41: R1. | | CONFIRMED-BUG | old | S 6ce860617 |
| X2 | Raid 60 reproduces `before.tsv`: 795.3 / 794.1 alliance, 792.2 / 792.6 horde. The new rows reach only 40 and 41. | `out/raid60.tsv` | Aligned | - | - |
| X3 | **Races.** Alliance: Human 798.8 / 798.2, Gnome 795.3 / 794.1 (-0.5%). Horde: Orc 792.2 / 792.6, Undead 790.2 / 790.0 (-0.3%), Troll 782.7 / 783.4 (-1.2%). This matches destruction.md:110 (Human and Orc, Gnome a step behind, Undead close, Troll a step behind). The published picks are within 1%. | `out/races.txt` | Aligned | - | - |
| X4 | Bane of Havoc 1 in both builds is Incinerate's prerequisite (rounds 14 and 16). | | Aligned | - | - |

## Before/after

Only the Destruction bare fights at 40 and 41 move. Seeds 1 / 50001, 10,000 iterations:

| level | race | before | after | change |
|---|---|---|---|---|
| 40 | orc | 169.5 / 169.5 | 169.8 / 169.7 | +0.2% / +0.1% |
| 40 | human | 167.9 / 167.8 | 168.4 / 168.3 | +0.3% / +0.3% |
| 41 | orc | 170.7 / 170.5 | 170.9 / 170.9 | +0.1% / +0.2% |
| 41 | human | 169.0 / 168.9 | 169.6 / 169.6 | +0.4% / +0.4% |

Raid 60 and every other level are unchanged. Absolute at raid 60, seed 1, horde: Affliction 750.6, Demonology 724.3, Destruction 792.2.

## Files to regenerate

- sim/request ladder goldens (TestRotationLadder): only where the Destruction 40 and 41 cells read the rotation's ranks.
- addon-data.json and Data.lua: the Destruction rotation lines carry the 40 and 41 rank rows.

## Measure at refresh

- Carried from round 18: confirm that the Affliction bare 40, 50 and 60 BiS entries carry the Malevolence string.

## For other lanes

- None. No shared code touched; nothing moves DPS outside warlock.

## For the guide lane

- Done in this lane: destruction.md "Levels 20 to 60" now names the 40 and 41 ranks (wow-player SHIP). Race order on all three pages re-checked; no change needed.

## Guide text not done

- None.

## Still open

- L1: Eureka! on Incinerate, demon crit, Soul Siphon's count, and Wrack and Doom at 60. These need level-40+ and level-60 Forever logs, which will take days after launch. Last looked at 19:12 UTC, 2026-11-04.
- Pending owner, not reopened: level-60 trainer book ranks, Siphon Life tick crit, the Demonology Horde race, Heart of the Lion 280/276, warlock Orc (O-list).

## Sources

- wago.tools 1.60.1.70755 SpellEffect, SpellPower, SpellCooldowns and SpellMisc, layered with hotfixes (cached `r17/warlock/wago`, linked from `r19/warlock/wago`)
- foreverlogs.gg `api/reports/public`, 2026-11-04 19:12 UTC; api.foreversixty.gg `/v1/reports/recent`, same time
- Lane outputs under `$R/warlock/out/`
