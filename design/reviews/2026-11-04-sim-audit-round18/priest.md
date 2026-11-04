# Sim-accuracy round 18: priest

Lane priest. Site wt/r18-priest and fork lane/r18-priest, both cut from int/r18 (site 4600b7d82, fork b2c4a2d05). Client 1.60.1.70755 with the late re-layer.

Specs: Discipline, Holy and Shadow, plus the shared low-rank penalty (`core.SpellLevelPenalty`, MaxLevel + 12).

Every sim went through `$R/sim.sh` on frozen BiS (`$R/bis-frozen`). Raid 60 uses the guide raid build. Exploration ran at 3,000 iterations. Every number in a row note or in this report is 10,000 iterations at seeds 1 and 50001, both factions, unless it says otherwise.

**How a rank was judged.** I sliced level by level with `spec-breakdown -level` at the floor, +3, +6 and +9 (every level where it was close). I forced each candidate rank by authoring that rank's own id, and replaced every reference to the authored id, conditions included, so the rows did not interfere. A first pass that swapped only the cast id gave false Shadow Word: Pain results. Those numbers were thrown away.

## 1. The low-rank penalty: no new source settles anything

| Source (checked 2026-11-04, about 16:10 UTC) | What it holds |
|---|---|
| foreverlogs.gg `/api/reports/<id>` | It now answers `403 "API access not allowed"` for every id: 2814, 2850, 2900, 3000, 3100. |
| api.foreversixty.gg/v1/reports/recent | Nothing new since round 17. |
| HealRankAdvisor (github.com/mendeleevforever/HealRankAdvisor) | The last commit is still 2026-10-24. |
| ForeverRank (github.com/PhilForever/ForeverRank, pushed 2026-10-15) | This is new to our reports. It is an addon that reads live tooltips and ranks them by healing per mana. It holds no measurements and no formula. |
| ForeverChanges /downranking | Still "updated October 16". It still says "Measured on heals in the beta, up to its level 30 cap. Above level 30, and for damage spells, I assume the same rule." |
| r/wowforever, plus a web search | No thread with numbers. EpicCarry's downranking post (403 from here) repeats the 18-level rule as a claim. |
| Warcraft Logs | Blocked (expired login). Recorded, not worked around. |

**Verdict: OPEN, unchanged.** That covers the top-rank exemption, damage spells and levels above 30. The owner's in-game test (round 16 priest.md, section 1) still settles all three. There is no code change.

### Two side leads from the sources, both OPEN

**P1. The penalty also cuts proc chances. The engine models only the spell-power part.**
- Blizzard's beta notes of 1 October 2026 say low ranks "have a reduced chance to activate your Class abilities and Talents (examples: Frostbite and Omen of Clarity)". The developer note adds: "Using a Rank 1 Frostbolt at level 60 will result in a 0% chance to get a Frostbite activation."
- The 8 October notes say the same for Improved Wing Clip.
- The 16 October notes take Touch of the Dead out of it.
- The engine (`sim/core/spell.go:787-829`) scales only `BonusCoefficient`. No proc chance anywhere reads the rank's level.
- The curve is unknown. Blizzard gives one endpoint, 0% at a rank whose spell-power share is also 0.
- **For priest the effect is negligible:**
  - Raid lists cast no share-0 rank into a chance proc. Lesser Heal 3 at 60 feeds only crit-triggered effects (Inspiration, Divine Aegis), which have no activation roll.
  - Shadow Weaving is 3/3 (100%) wherever the band rows cast low ranks.
- **OPEN** until a curve exists. I left it as is, because a guessed curve would move other classes. See "For other lanes".

**P2. Heals may scale with the target's level.**
- The source is the Blizzard forum thread `us.forums.blizzard.com/en/wow/t/2371681`, 4 to 19 October, with no blue reply.
- A level-26 paladin's Flash of Light rank 1 healed a level-24 target for 90-96 against the tooltip's 101-108.
- One poster reports about -20% inside dungeons.
- This is single-source and unexplained. Every sim fight heals same-level targets, so no published number is affected. **OPEN.**

## 2. Regression target: Discipline's pre-pull Renew (round 17)

| check | result |
|---|---|
| The list and its copies | `priest-discipline.json` prepull is Lightwell at -2.5 s, then Renew 10927 rank 7 on player 5 at -1 s. The site copy and the fork copy match the curated list. `TestDisciplineRenewsTheTankBeforeThePull` passes. |
| Raid 60, horde, seed 1 | With the line: 495.7. Without it (`-rotation` with prepull[1] removed): 494.5. That is +0.24%, which matches round 17 exactly. |
| Bands | At band 30 the pre-pull Renew follows the 30-33 row (6074, rank 2). At band 40 it casts rank 6 before this round, and rank 4 now (new row, see Discipline). |

**Aligned.**

## Shadow

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **The Mind Blast rows at 40-49 were wrong past 43.**<br>- **Before:** rank 3 (Dwarf, Night Elf) or rank 4 (everyone else) at every level from 40 to 49.<br>- **Why it fails:** under the penalty, rank 4 (learned 28, MaxLevel 33) loses spell power from 46, and rank 3 from 40.<br>- **Level by level:** rank 4 beats the highest rank learned at 40 to 43 (Human +0.3% / +2.0% at 40 / 43, Undead +0.8% / +1.3%, Troll +2.1% / +1.9%, Dwarf +0.9% / +0.5%, Night Elf +0.9% / +0.3%). From 44 the top rank wins: Human -1.0% with rank 4 at 44 and -3.9% at 49; Dwarf -2.9% at 49.<br>- **Rank 3 never beat rank 4:** a Dwarf or Night Elf loses 1.0% to 1.9% with it at 40 and 43.<br>- **Gnome:** rank 4 is +0.1% at 40 and -0.4% at 43, so a Gnome keeps the default.<br>- **Undead:** rank 4 still won 44-47 on its own, but measured jointly with Devouring Plague, the top Mind Blast with Devouring Plague rank 2 beats rank 4 with rank 1 by +0.5% (44) to +2.1% (49). So Undead follows the same row.<br>**Fix:** one row, rank 4 at 40-43 for Human, Dwarf, Night Elf, Undead and Troll. | `mb-10k.txt`, combos (lane scratch); `rank-choices.json` | CONFIRMED | old (round-13 rows measured before round 15's penalty; rounds 15-16 re-judged Renew and Devouring Plague 52+, not these) | S b44b9c521 |
| S2 | **Devouring Plague rank 1 at 44-49.** The row held rank 1 to 49. Rank 1 (learned 20) loses spell power from 38. With the top Mind Blast from 44, rank 2 wins from 44 for every race: Human +1.6% at 44 and +1.2% at 49, Troll +3.5% / +2.9%, Gnome +0.4% / +2.0%, Dwarf +1.0% / +3.2%. Rank 3 and rank 4 trail rank 2 (Gnome rank 3 +0.2% at 49 only).<br>**Fix:** rank 1 at 40-43, rank 2 at 44-49. | `dp-10k-c.txt` | CONFIRMED | old (as S1) | S b44b9c521 |
| S3 | **Devouring Plague at 50-51 had a round-13 race split** (rank 1 for five races, rank 2 for Night Elf). Rank 2 now beats or ties rank 1 at 50 for every race: Gnome +1.2%, Troll +1.7%, Human +0.6%, Dwarf +0.6%, Night Elf +0.5%, Undead a tie. At 51 a Human (-0.3%) and a Night Elf (-0.2%) lose a hair, too little for a one-level key.<br>**Fix:** rank 2 at 50-51 for every race. | `s50.txt` | CONFIRMED | old | S bfacc6b65 |
| S4 | **Shadow Word: Death rank 2 at 50-59 skipped the Human** (round 13: -0.4%). Now a Human gains +2.2% at 50, +0.3% at 53, +1.5% at 56 and +0.4% at 59 over the top rank learned. Every other race gains or ties (Undead +1.9% at 56).<br>**Fix:** the race key is dropped. | `s50.txt`, true rank 4 through a build without the row | CONFIRMED | old | S bfacc6b65 |
| S5 | **Rows that hold** (level by level):<br>- Shadow Word: Pain rank 3 at 30-39 beats ranks 4 and 5 by 2-4%.<br>- Devouring Plague rank 1 at 30-39 (rank 3 ties at 36).<br>- Devouring Plague rank 2 at 52-59.<br>- Bare 60: Devouring Plague rank 2 for Dwarf, Night Elf and Troll (+0.8%); Human +0.2% and Gnome -0.2% stay on the default; Undead -2.0%. Shadow Word: Death rank 2 at bare 60 is +1.0% to +1.6% for every race.<br>- Mind Blast at 50-56: rank 6 beats the default for Human and Undead (+0.3% to +0.8%) and loses for Gnome and Troll (-0.1% to -3.3%). That is mixed, so no row. | `shadow-sweep2.txt`, `s50.txt` | Aligned | - | - |
| S6 | **Fresh re-audit, raid 60** (Undead, guide build):<br>- Damage shares: Mind Flay 39%, Shadow Word: Pain 20%, Mind Blast 19%, Devouring Plague 12%, Shadow Word: Death 9%.<br>- No time out of mana. Inner Focus 1.3 casts. Vampiric Embrace 6 casts.<br>- Conformance: priest has no damage "differs" row.<br>- The round-17 S1-S3 items stand. | spec-breakdown; conformance golden | Aligned | - | - |

## Holy

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| H1 | **The tank Renew at 40-59 cast the highest rank learned** (the list authors rank 7). Holy is mana-bound on band gear, and a cheaper rank that keeps its whole bonus heals far more over the fight. Gains over rank 7, Horde / Alliance:<br>- Rank 4: +4.3% / +3.5% at 40, +6.1% / +6.3% at 44.<br>- Rank 5 (learned 32): +5.8% / +5.6% at 45, +6.0% / +5.9% at 50. At 51 it splits by faction against rank 6 (-0.4% / +0.7%).<br>- Rank 6 (learned 38): +1.7% / +0.9% over rank 5 at 52, +6.6% / +6.0% over rank 7 at 53, +2.2% / +2.6% at 59.<br>- At 60 (bare and raid) rank 7 wins, so the list is right at 60.<br>**Fix:** rows rank 4 at 40-44, rank 5 at 45-51, rank 6 at 52-59. Leveling prose names them. | `renew-sweep.txt`, `renew-10k.txt` | CONFIRMED | old (round 15 authored rank 7 for 60; rows were judged only to 39) | S 0b4ba8103, 6ef187a33 |
| H2 | **Greater Heal (top-rank lines) at 46-57:** rank 1 beats the highest rank learned. Horde: +1.2% at 46, +0.9% at 49, +0.3% at 52 and 55. Alliance: +0.3% to +0.4%. At 58-59 ranks 1, 3 and 4 are within 0.2%.<br>**Fix:** a row, rank 1 at 46-57. | `gh-10k.txt` | CONFIRMED | old (round 15 W6) | S 9b74c9dba |
| H3 | **The step-down below the top Greater Heal was Heal rank 4,** which keeps 55% of the bonus at 60 (MaxLevel 39). Greater Heal rank 1 keeps 85%.<br>- **Raid 60, with a rank-1 Greater Heal line before each Heal line** (same target and condition): Alliance 773.8 to 780.6 (+0.88% / +0.98%), Horde 793.5 to 797.8 (+0.54% / +0.63%).<br>- **Bare 60:** +0.39% / +0.28% Alliance, +0.77% / +0.72% Horde.<br>- Using rank 2 or 3 as the step-down loses about 0.6%.<br>- Replacing the Heal lines outright gains more on bare gear (+1.6% / +3.2%) but drops Heal before 40, so I did not take it.<br>**Fix:** ten lines (2060 rank 1) ahead of the ten Heal lines. They drop out before 40. Notes, the filler / Overheal / Low ranks prose, and the fork copy all follow. | combos; `TestHolySteppingDownCastsGreaterHealRankOneBeforeHeal` failed, then passed; `make apl-check ENGINE_DIR=<lane engine>` clean | CONFIRMED (rotation) | old (round 15 W6 kept Heal 4 as the step-down under the round-15 penalty) | S 6eacb213b, 53c75f4d7; E 442acc4b2 |
| H4 | **Prayer of Healing rank 1 at bare 60** keeps 55% of its bonus there. Rank 2 (learned 40, MaxLevel 49) keeps all of it. Rank 2 beats the top rank by +0.5% / +0.6% (Horde / Alliance) and rank 1 by +0.4% / +0.1%.<br>**Fix:** row to rank 2. The guide already says to downrank within the forties ranks. | `sb-nopoh` sweep | CONFIRMED | old (round 13 row) | S 2fee3af1d |
| H5 | **Rows that hold:**<br>- Prayer of Healing rank 1 at 40-49 and 50-59 (ties rank 2 at 59).<br>- Renew rank 1 at 14-32 and rank 3 at 33-39 (re-swept at 20/23/26/29/30/33/36/39).<br>- Flash Heal ranks at 46-58 tie.<br>Raid 60 (Undead):<br>- Prayer of Mending is the best healing per mana (10.6).<br>- Renew rank 7 is 7.1 healing per mana.<br>- Heal 4 was cast 17 times a fight before H3. | sweeps; spec-breakdown | Aligned | - | - |

## Discipline

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| D1 | **The tank Renew at 40-49 cast the highest rank learned.**<br>- Rank 4: +2.0% / +1.9% at 40 and +1.9% / +1.8% at 44 over rank 7.<br>- Rank 5: +1.4% / +1.3% at 45 and +2.2% / +1.7% at 49.<br>- From 50, ranks 6 to 8 tie and rank 7 stands. Rank 9 at 56-59 is within 0.5% (Alliance 0.0% at 59), so no row.<br>**Fix:** rows rank 4 at 40-44 and rank 5 at 45-49. They also govern the pre-pull Renew. | `renew-10k.txt` | CONFIRMED | old (as H1) | S 0b4ba8103, 6ef187a33 |
| D2 | **The tank's Greater Heal (10965) at 46-57:** rank 1 beats it.<br>- +0.9% / +0.8% at 46.<br>- +0.5% / +0.4% at 49.<br>- +0.6% / +0.4% at 52.<br>- +0.2% / +0.5% at 55.<br>**Fix:** a row, rank 1 at 46-57. | `gh-10k.txt` | CONFIRMED | old | S 9b74c9dba |
| D3 | **Fresh re-audit, raid 60.** Penance 16.9 casts. Power Word: Shield rank 10. The pre-pull Renew (section 2). Lesser Heal 3 is cast 57 times at 0% bonus as the deeper-hole top-up; that is the pending-owner low-rank item and is not reopened. Mana runs dry at 78 s, as known (round 16 D2). | spec-breakdown | Aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | The penalty's open parameters (section 1) | sources above | OPEN | - | - |
| X2 | The proc-chance half of the penalty is unmodeled (P1) | Blizzard beta notes, 1 and 8 October | OPEN (curve unknown) | - | - |
| X3 | Target-level heal scaling (P2) | forum thread 2371681 | OPEN (single source) | - | - |
| X4 | **Test method:** a rank-swap harness must rewrite every reference to the authored id, not only the cast. Otherwise a DoT condition watches the wrong rank, and the swapped line recasts on every global (Shadow Word: Pain rank 4 showed -29% that way). | lane harness | Note | - | - |

## Before/after

Before = `$R/before.tsv` (raid 60) and `$R/sb-before` from `$R/site-before/sim` (bands). After = this branch at 53c75f4d7.

**Raid 60** (guide raid builds, frozen BiS):

| spec | alliance s1 | alliance s50001 | horde s1 | horde s50001 |
|---|---|---|---|---|
| priest-discipline | 506.2 (0.00%) | 0.00% | 495.7 (0.00%) | 0.00% |
| priest-holy | 773.8 → 780.6 (+0.88%) | +0.98% | 793.5 → 797.8 (+0.54%) | +0.63% |
| priest-shadow | 742.6 (0.00%) | 0.00% | 735.5 (0.00%) | 0.00% |

**Bare bands.** The value is the band at the level shown. The 60 rows include H3 for Holy; the others predate H3, which moves only 58+.

| spec | level | alliance s1 | alliance s50001 | horde s1 | horde s50001 |
|---|---|---|---|---|---|
| holy | 40 | 162.0 → 167.6 (+3.46%) | +3.52% | 170.7 → 178.1 (+4.34%) | +4.27% |
| holy | 46 | 162.9 → 171.7 (+5.40%) | +5.53% | 170.6 → 181.2 (+6.21%) | +6.28% |
| holy | 50 | 232.2 → 246.3 (+6.07%) | +6.03% | 245.1 → 260.6 (+6.32%) | +6.20% |
| holy | 56 | 243.2 → 249.7 (+2.67%) | +2.59% | 250.4 → 263.3 (+5.15%) | +5.24% |
| holy | 60 | 387.2 → 388.7 (+0.39%) | +0.28% | 402.9 → 406.0 (+0.77%) | +0.72% |
| discipline | 40 | 122.8 → 125.1 (+1.87%) | +1.79% | 128.3 → 130.9 (+2.03%) | +1.95% |
| discipline | 46 | 118.4 → 120.7 (+1.94%) | +2.03% | 123.9 → 126.7 (+2.26%) | +2.26% |
| discipline | 50 | 135.3 → 135.8 (+0.37%) | +0.30% | 144.2 → 144.8 (+0.42%) | +0.42% |
| discipline | 56 | 143.2 → 143.9 (+0.49%) | +0.49% | 152.1 → 152.3 (+0.13%) | +0.07% |
| discipline | 60 | 201.8 (0.00%) | 0.00% | 200.8 (0.00%) | 0.00% |
| shadow | 40 | 150.0 → 149.8 (-0.13%) | -0.13% | 169.1 (0.00%) | 0.00% |
| shadow | 46 | 162.3 → 167.9 (+3.45%) | +3.45% | 183.7 → 185.4 (+0.93%) | +0.87% |
| shadow | 50 | 213.0 → 215.5 (+1.17%) | +1.17% | 242.1 → 242.0 (-0.04%) | 0.00% |
| shadow | 56 | 228.0 (0.00%) | 0.00% | 256.2 (0.00%) | 0.00% |
| shadow | 60 | 389.4 (0.00%) | 0.00% | 423.6 (0.00%) | 0.00% |

Shadow Alliance at 40 is the published Gnome, which now casts the top Mind Blast. Rank 4 was a tie at 40 and -0.4% at 43, so -0.13% at the floor is the per-level trade.

**No other class moves.** Every change is a priest row, the priest-holy APL, or priest prose.

## Files to regenerate

- `data/builds/1.60.1.70755/addon-data.json`, then `addon/ForeverSixty/Data.lua`. These carry the new rank rows (Mind Blast, Devouring Plague, Shadow Word: Death, Renew, Greater Heal, Prayer of Healing) and Holy's ten Greater Heal rank 1 lines with their notes. The guide's Priority card renders from these.
- TestRotationLadder goldens: priest-holy (40, 50, 60 bare and raid), priest-discipline (40, 50), priest-shadow (40, 50).
- dps-per-point: all three priest specs.
- Coach starters, if they read the rotation lines.
- Copied by hand, not regenerated: `sim/rankchoices/rank-choices.json` (`make rank-choices-sync`, a plain `cp`), `sim/request/apl/priest-holy.apl.json`, and the fork's `ui/healing_priest/apls/forever_holy.apl.json` (E 442acc4b2).

## Measure at refresh

- priest-holy and priest-discipline at every band 40-60. The rows and the step-down were judged on frozen gear, and a re-ranked band gear set may move the Renew and Greater Heal switch levels by a level or two.
- priest-shadow 40 and 50 (the per-race rows).

## For other lanes

- **Every caster/healer lane (shared penalty, P1):** the engine does not reduce proc chances for low ranks. Blizzard names Frostbite and Omen of Clarity (1 October) and Improved Wing Clip (8 October). Nothing to change until a curve exists. Watch any row that casts a share-0 rank for its proc (mage Frostbolt rows, druid Wrath for Omen).
- **Every healer lane (method):** the tank-heal rows at 40-59 were never judged level by level. Holy and Discipline gained 2-6% from cheaper ranks there. Restoration druid, Restoration shaman and Holy paladin likely have the same gap (Rejuvenation, Healing Wave, Flash of Light at 40-59).
- **Platform:** the rank-swap harness note (X4).

## For the guide lane

Done in this lane, with wow-player SHIP:
- holy.md leveling (Renew ranks 4/5/6, Greater Heal rank 1 until 58);
- holy.md rotation (filler step-down to Greater Heal rank 1, the Overheal and Low ranks lines);
- discipline.md leveling (Renew ranks 4/5, then 7 or 8; Greater Heal rank 1 until 58);
- shadow.md leveling (the per-level ranks).

**Pending after the regen:** the Holy Priority card. It renders from addon-data, so the new "Greater Heal rank 1" row and the rewritten Heal note show only after the integrator's regen. The reviewer asked to re-capture the card then and to check that the longer notes do not clip.

## Still open

- The penalty: the top-rank exemption, damage spells and levels above 30 (owner test, round 16 section 1). Also the proc-chance curve (P1) and target-level scaling (P2).
- Eureka!'s periodic mask and Lightwell Renew's crit source (round 17), with no change.
- foreverlogs.gg now refuses its API (403). Warcraft Logs is blocked.
- Disc Renew rank 9 at 56-59 and Greater Heal rank 1/3 at 58-59 are within 0.5%. No row.

## Sources

- Blizzard, WoW Forever Beta Development Notes (1, 8 and 16 October 2026): https://us.forums.blizzard.com/en/wow/t/wow-forever-beta-development-notes-%E2%80%93-updated-october-8/2360696/1
- Blizzard forums, "[bug?] Healing spells are scaling based on the targets level", 4 to 19 October 2026: https://us.forums.blizzard.com/en/wow/t/2371681
- ForeverChanges /downranking (updated 16 October 2026, fetched 4 November): https://foreverchanges.pro/downranking
- HealRankAdvisor: https://github.com/mendeleevforever/HealRankAdvisor (last commit 2026-10-24)
- ForeverRank: https://github.com/PhilForever/ForeverRank (pushed 2026-10-15)
- EpicCarry, "WoW Forever Downranking Spell Nerf" (403 from here): https://epiccarry.com/blogs/?p=41909
- Client: `data/builds/1.60.1.70755/spellranks.json`; MaxLevel exceptions in `sim/core/spell_max_level_auto_gen.go`
- Lane scratch, the tables behind every number: `$R/priest/{mb-10k,dp-10k-c,s50,renew-10k,gh-10k,bands}.txt|tsv`

## wow-player verdict

**SHIP**, over three passes on the rendered pages (dev server, 1440 wide). The captures are in `priest-captures/`.

- **Holy, Discipline and Shadow leveling:** SHIP with no MUSTs. The SHOULDs I applied:
  - switch "when you learn rank 4 at 58";
  - no overclaim on Greater Heal;
  - Discipline's Renew from 50 is "rank 7 or 8";
  - the Shadow sentence is split, with "then the top rank".
- **Holy rotation prose:** SHIP. The SHOULDs I applied:
  - the Overheal and Low ranks paragraphs agree with the step-down;
  - "leave ranks 2 and 3 off your bars".
- **The one MUST:** re-capture the Priority card after the regen (see "For the guide lane"). The prose and notes are signed off.
- **Not applied:** the tool bar over "LOAD IN PLANNER" is Astro's dev toolbar, not the page. A note on putting old ranks on the bar is optional and left out.
