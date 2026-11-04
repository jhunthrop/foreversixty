# Sim-accuracy round 19: priest

Lane priest. Site wt/r19-priest and fork lane/r19-priest, both cut from int/r19 (site f91e9d3f9, fork 6ea675594). int/r19 merged into the site branch at 2fffcbb74. Client 1.60.1.70755 with the late hotfix re-layer.

Every sim went through `$R/sim.sh` on frozen BiS (`$R/bis-frozen`), at 10,000 iterations, seeds 1 and 50001. Raid 60 used the guide build (`$R/rt`). Rank rows were judged with `spec-breakdown -level` and a build with the priest rows taken out (`sb-norow`), so each candidate rank was authored outright. A swapped rank counted only when its line was seen to cast it.

**Error.** spec-breakdown gives a DPS error of about 0.27 at 10,000 iterations, about 0.07%. The healer output prints none. For the same configuration, seeds 1 and 50001 differ by at most 0.3 HPS (about 0.12%), so a heal row was taken only on a gain of more than 0.3 HPS at both seeds. Per the integrator's correction, a rank row is adopted on any gain beyond error, and the 1% bar applies to talent builds only.

## Seeds

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P2 | **Heals that scale with the target's level.** Still one Blizzard forum thread, `us.forums.blizzard.com/en/wow/t/2371681`, last post 19 October, with no blue reply. Posters measured a level-26 Flash of Light rank 1 healing 90-96 on a level-24 target against 101-108 on a level-26 one. One poster saw about -20% inside dungeons. Others say HoTs are hit less.<br>No formula, no second source, and no log we can reach. Every sim fight heals targets of the healer's own level. A flat dungeon cut would not change any rank choice. | forum (fetched 19:10 UTC); web search | OPEN (single source) | - | - |
| P3 | **Eureka!'s periodic mask.** The client row 1259823 has three effects: damage/healing (misc 0), periodic damage (misc 22, mask 10485760/98304) and cost (misc 14). The periodic mask reaches only the channels: Mind Flay, Starshards and Penance's pulses. No DoT or HoT is reached (Renew, Shadow Word: Pain and Devouring Plague are absent). That reads as Blizzard counting channels among the "non-periodic" abilities of the tooltip. `sim/core/eureka_auto_gen.go` lists exactly these, and the Forever Shadow Word: Death ranks (1309595 to 1309636, mask_1 bit 2) are in the damage and cost lists.<br>**Log search:** the events endpoint filtered by spell 1259823 over all 544 encounters of the 9 cached beta reports with a Shadow priest (2668 to 2785) found no Eureka! event. No beta Gnome Shadow priest used it. | client SpellEffect / SpellClassOptions 70755; `$R/priest/eu/scan.txt` | Aligned with the client; OPEN for a log | - | - |
| P4 | **Lightwell Renew's crit source.** The well's renew (7001 / 27873 / 27874) has EffectBonusCoefficient 0 on the client, and the engine adds no spell power (`sim/priest/lightwell.go`). Whose crit it rolls (the priest's, or that of the member who clicks) cannot be read from the client. Lightwell is a level-40 talent, so no beta log (cap 35) has it. It is worth 16 ticks of 320 a fight. | client rows; no Lightwell in the r10 healer timelines | OPEN (negligible) | - | - |
| P5 | **Coefficient lead (Ice Lance pattern).** I checked every priest spell the engine scales with spell power against the 70755 client SpellEffect EffectBonusCoefficient. None is 0 on the client:<br>- **Damage:** Smite (0.429/0.571/0.714), Holy Fire (0.75, DoT 0.05), Mind Blast 0.429, Mind Flay 0.167, Shadow Word: Pain 0.2 (10894; 27605 0.167), Devouring Plague 0.1, Shadow Word: Death 0.429, Starshards 0.167, Holy Nova 0.107.<br>- **Healing:** Lesser Heal (0.429/0.571/0.714), Heal 0.857, Flash Heal 0.429, Greater Heal 0.857, Prayer of Healing 0.286, Renew 0.2, Power Word: Shield 0.1, Penance 0.19, Prayer of Mending 0.429, Desperate Prayer 0.429, Holy Nova heal 0.107, Binding Heal 0.429.<br>- **Unscaled:** Lightwell Renew is 0 on the client and unscaled in the engine.<br>The generated `PenanceSpellCoeff` (0.1333/0.4286, convention defaults) and the PoM/PW:S/Renew convention slots are never read: the heals use `heal_ranks.go`. Every row of that file matches the 70755 client for amount, per-level, coefficient, cost, level and max level (`tools/ladder.py`). The damage files read only ranks 1+ of the generated arrays, all client values. | `$R/priest/coef.txt`; `heal_ranks.go`; `constants_auto_gen.go` | Aligned | - | - |
| P6 | **Low-rank proc watch (flag only).** No priest row casts a rank with a spell-power share of 0 into a chance proc. The lowest is Devouring Plague rank 2 at 52-60 (share 0.65 down to 0.25). Its only chance proc is Shadow Weaving at 3/3 (100%), which Mind Flay and Mind Blast keep stacked anyway. Discipline's raid Lesser Heal 3 (share 0) feeds only crit-triggered effects (Divine Aegis, Inspiration), which have no activation roll. | share table, `$R/priest` | Flag: none | - | - |
| C1 | **Holy card wording (wow-player SHOULDs).** Greater Heal filler: "at your top trainer rank" (rank 5 is a book). Renew: "Keep Renew rank 7 or higher on the tank all fight", for Holy and Discipline both, as holy.md and discipline.md say. The card was previewed with those strings: **SHIP** on the two notes. | `shots/holy-priority.*`, `shots/discipline-priority.*` | Done | - | S 64d90332c |

## Discipline

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| D1 | **No tank filler before 40.** The tank filler is Greater Heal 10965 (learned 40), with no fallback. Below 40 the tank got Penance (from 30), Renew and the shield only, while the party got Lesser Heal and Flash Heal. The guide's own leveling text calls Discipline "a Renew, Heal and shield healer" there.<br>**Fix:** a Heal rank 4 line behind each tank Greater Heal line, with the same target and condition, gated on Greater Heal rank 1 not yet learned.<br>**Results (Alliance / Horde):**<br>- Bare 20: +7.4% / +9.1%.<br>- Bare 30: +2.7% / +2.1%.<br>- Bands 40-60 and raid 60: unchanged to 0.1 HPS.<br>Ungated, the Heal line cost 0.8% to 1.1% at raid 60.<br>Card note and leveling prose added. | `$R/priest/dheal.txt`, `ba-disc.txt`; `TestDisciplineHealsTheTankWithHealBeforeGreaterHeal` failed, then passed | CONFIRMED | old (integrator lead from the wow-player card review) | S 9afb58196; E e08bc5b05 |
| D2 | **The pre-pull Renew held the global into the fight.** Lightwell at -2.5 s (1.5 s cast) ends at -1 s, and the round-17 Renew at -1 s held the global to +0.5 s. The fork's `TestInnerFireDoesNotCutIntoThePrepullCasts` was failing on int/r19 (6ea675594) for that reason.<br>**Fix:** Lightwell at -3 s and Renew at -1.5 s. The opener card now reads 3 s and 1.5 s, and the prose reads "just before he engages". Raid 60 and every band are unchanged to 0.1 HPS. | fork test failed on base and on the lane, then passed | CONFIRMED | old (round 17 5dd716a48 added the -1 s Renew) | S 9afb58196; E e08bc5b05 |
| D3 | **Lesser Heal rank 3 from 40 to 59.** Discipline's Lesser Heal line fires when Greater Heal rank 1 is out of mana reach. Lesser Heal 3 (learned 10) has lost its whole healing bonus by 47 under the level-gap rule, so each cast heals for little and restarts the five-second rule.<br>**Measured:** swapping it for Heal rank 4 gained hugely. Heal 4 itself was cast only 0.1 times a fight (it needs the mana Greater Heal rank 1 lacked), so the whole gain is in not casting Lesser Heal.<br>**Fix (integrator ruling: 40-59 only; the owner keeps 60):** the Lesser Heal lines are scoped out of 40-59 with the platform `currentLevel` gate.<br>**Results (Alliance / Horde), both seeds:**<br>- +7.7% / +6.3% at 40, +13% / +11% at 43, +20% / +17% at 46, +21% / +17% at 49.<br>- +46% / +39% at 50, +45% / +38% at 53, +45% / +38% at 56, +44% / +38% at 59.<br>- Bare 20/30/60 and raid 60 unchanged.<br>Lesser Heal casts drop to 0 in the range.<br>**Trap caught on the way:** a first gate on "not known Power Word: Shield rank 10" never fired (the rank rewrite resolves it) and measured 0.0. | `d3f.txt`, `var-disc.txt`, `lv-disc.txt`; `TestDisciplineSkipsLesserHealFromFortyToFiftyNine` failed, then passed | CONFIRMED | old (round 18 D3 saw the casts and held them) | S 44c6090ac; E f12739f55 |
| D4 | **Renew rows re-judged on the new list.**<br>- **Rank 9 (learned 56) against rank 7:** Human +1.5% to +1.8%, Gnome +0.7% to +0.9%, Dwarf +0.2% to +1.8%, Night Elf +0.3% to +1.8% at 56-59. Undead +0.3% at 59 only (within error at 56 and 58, -0.2% at 57). Troll -0.1% to -0.3%.<br>- **A Human with the top rank learned (rank 7 from 44)** beats the round-18 rows at 44 (+0.6%), 45 (+0.4%), 46 (+0.3%) and 47 (+0.5%), and ties from 48.<br>- **Held:** Greater Heal rank 1 at 46-57 (the top rank loses 0.2% to 0.8%) and Prayer of Healing rank 1 at 40-49 (rank 2 loses 0.4% to 0.8%) for Human, Gnome and Undead.<br>**Fix:** keyed rows. (An earlier set of rank 9 rows, judged on the list with Lesser Heal, was replaced.) | `out7.txt`, `out8.txt`, `out9.txt` | CONFIRMED | old (round 18 kept rank 9 "within 0.5%") | S b05dd9586, 44c6090ac |
| D5 | **Race pick (guide round 15 request), raid 60, guide build:**<br>- Human 506.2 / 505.9.<br>- Gnome 502.4 / 502.2 (Human +0.8%).<br>- Undead 495.7 / 495.9.<br>- Troll 481.9 / 482.0 (Undead +2.9%).<br>The guide's Gnome is within 1% of Human, so it stays. | `race.txt` | Aligned (within 1%) | - | - |
| D6 | **Fresh re-audit, raid 60:** heal ladders match 70755 (P5). Penance's four ranks have the right helper ids (402289, 1240723, 1240724, 1316991) at 0.19. The build has no dead points. | as P5 | Aligned | - | - |

## Holy

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| H1 | **Renew rows by race at 51-53.**<br>- At 51 an Undead gains +0.5% / +0.4% with rank 6. Rank 5 stays for the rest (Human +0.8%, Gnome +2.2%, Dwarf +1.7%, Troll +2.0%).<br>- At 52 a Gnome (+0.3% / +0.5%), Dwarf (+1.0%) and Night Elf (+0.3% / +0.5%) keep rank 5, while Human, Undead and Troll take rank 6.<br>- From 53 rank 6 wins for all (+1.2% to +1.9%).<br>Round 18 had judged only the published Human and Undead. | `out2.txt`, `out4.txt`, `out6.txt` | CONFIRMED | old | S b05dd9586 |
| H2 | **Greater Heal rank 3 at 59** beats rank 4 (learned 58): Human, Undead, Troll and Dwarf +0.2%, Gnome +0.3%, Night Elf +0.1%. At 58 the gain is within error for Human and Troll, so the row is 59 only. | `out1.txt`, `out2.txt`, `out6.txt` | CONFIRMED | old (round 18 "within 0.2%") | S b05dd9586 |
| H3 | **No Power Word: Shield row** (wow-player SHOULD). I added Discipline's tank shield line to Holy:<br>- After Renew: -2.8% raid 60, -1.0% / -0.6% bare 60, -1.7% to -7% at bands 20-40.<br>- As the last line: -0.7% raid 60, still a loss everywhere.<br>Holy has no shield talents, and Weakened Soul locks the shield out for 15 s. **Deliberate.** The holy.md "On the tank" paragraph says why. | `var-holy.txt` | Aligned (now explained) | - | S 237bea092 |
| H4 | **Fresh re-audit:** heal ladders match 70755 (P5). Prayer of Mending heals at 0.429. The step-down to Greater Heal rank 1 (round 18) is unchanged. | as P5 | Aligned | - | - |

## Shadow

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **Devouring Plague at 50-51 for a Human.** Judged jointly with Mind Blast, rank 1 and the top Mind Blast beat the row's rank 2 by +0.9% at 50 and +0.8% at 51, both seeds. Round 18's "+0.6% for rank 2 at 50" no longer holds on this engine. Every other race keeps rank 2 (Gnome +1.2%, Troll +1.7%, Dwarf +0.6%; Night Elf and Undead within error).<br>**Fix:** a Human-keyed rank 1 row. | `out2.txt`, `out3.txt` | CONFIRMED | old (round 18 S3 kept the Human loss at 51 as too small) | S b05dd9586 |
| S2 | **Mind Blast rank 6 at 50-51.**<br>- An Undead (with Devouring Plague rank 2) gains +0.6% at 50 and +0.4% at 51.<br>- A Dwarf gains +0.4% at 51 only.<br>- A Human gains at 51 only jointly with Devouring Plague rank 2 (+0.8%), but rank 1 with the top Mind Blast is as good, so S1 covers him.<br>- A Gnome or Troll loses.<br>**Fix:** keyed rows. | `out1.txt`, `out2.txt`, `out3.txt` | CONFIRMED | old (round 18 S5 "mixed, no row") | S b05dd9586 |
| S3 | **Bare 60 Devouring Plague rank 2 for a Human:** +0.23% / +0.13% (0.9 / 0.5 DPS against an error of about 0.27 per run). Not beyond error at seed 50001, so there is no row. A Gnome loses 0.2%. | `out1.txt` | Aligned | - | - |
| S4 | **Race lines (guide round 15 request), raid 60, raid build:**<br>- Human 745.5 / 745.6.<br>- Night Elf 745.2 / 745.5.<br>- Dwarf 745.0.<br>- Gnome 742.6 / 742.9.<br>- Undead 735.5 / 735.0.<br>- Troll 727.6 / 727.3.<br>Alliance is within 0.4%, so the guide's Human stands. Undead beats Troll by 1.1%. | `race.txt` | Aligned | - | - |
| S5 | **Fresh re-audit:** every damage coefficient matches the client (P5). Eureka! masks match the client (P3). The raid build has no dead point: Dark Inspiration and Silent Resolve are party and threat utility, as round 17 found. | as P5 | Aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | **Rank swaps through conditions:** a ranked id inside `spellIsKnown` (or any condition) is rewritten by the site to the highest rank learned, so it cannot act as a level gate. A change that should move numbers but measures 0.0 is this trap. | D3 | Note (platform) | - | - |
| X2 | **Spot-check of round 18's 40-59 rows at +3/+6:**<br>- Holy Renew at 52/55/58 for every race (H1).<br>- Greater Heal rank 1 at 55/57 against rank 3: rank 1 holds (+0.2% to +0.3%).<br>- Discipline Renew rank 7 against 8 at 56/59: rank 8 never wins.<br>- Shadow Devouring Plague at 50/51/52 for every race (S1).<br>- Mind Blast at 50/51/53/56 for every race (S2): from 53 the top rank wins for all (rank 6 -0.2% to -3.2%).<br>Everything else holds. | out1-out6 | Aligned except D4, H1, H2, S1, S2 | - | - |

## Before/after

Before = `$R/before.tsv` (raid 60) and `$R/sb-before` from `$R/site-before/sim` (bands). After = this branch before the int/r19 merge.

**Raid 60:** no priest change.
- Discipline: 506.2 / 495.7 (Alliance / Horde), 0.00%.
- Holy: 780.6 / 797.8, 0.00%.
- Shadow: 742.6 / 735.5, 0.00%.
- The pre-pull retime moved nothing to 0.1.

**Bare bands (published race), Discipline:**

| band | alliance s1 | alliance s50001 | horde s1 | horde s50001 |
|---|---|---|---|---|
| 20 | 43.2 → 46.4 (+7.4%) | +7.4% | 41.7 → 45.5 (+9.1%) | +9.1% |
| 30 | 89.8 → 92.2 (+2.7%) | +2.7% | 90.6 → 92.5 (+2.1%) | +2.1% |
| 40 | 125.2 → 134.9 (+7.7%) | +7.7% | 131.0 → 139.3 (+6.3%) | +6.3% |
| 50 | 135.9 → 198.5 (+46%) | +46% | 144.8 → 200.8 (+39%) | +39% |
| 60 | 0.0% | 0.0% | 0.0% | 0.0% |

Bands 40 and 50 are D3 and the Renew rows together.

**Rank rows** move only the levels they name.
- **Published band floors:** the only one that moves is Shadow Horde (Undead) at band 50: 242.0 → 243.4 (+0.58%) / 241.8 → 243.2 (+0.58%).
- **Inside bands (level by level):** the gains are in the tables above.

**Other classes:** none. Every change is a priest APL, a priest rank row, a priest test or priest prose. E e08bc5b05 touches only `ui/healing_priest` and a `sim/priest/healing` test.

## Files to regenerate

- `data/builds/1.60.1.70755/addon-data.json`, then `addon/ForeverSixty/Data.lua`. They carry:
  - Discipline's two tank Heal lines and their note;
  - the new pre-pull times;
  - the rank rows (Renew, Greater Heal, Devouring Plague, Mind Blast);
  - the card notes from S 64d90332c and S 237bea092.
  
  The Priority cards render from addon-data.
- TestRotationLadder goldens: priest-discipline (20, 30, and every band for the pre-pull times), priest-holy and priest-shadow (50), if they print the rank rows.
- Coach starters: `data/curated/coach/priest-discipline.{md,json}` (the Renew note and the new Heal lines).
- dps-per-point: priest-discipline (bands 20/30).
- Copied by hand, not regenerated: `sim/rankchoices/rank-choices.json` (`make rank-choices-sync`), `sim/request/apl/priest-discipline.apl.json` and the fork's `ui/healing_priest/apls/forever_discipline.apl.json` (`make apl-sync`).

## Measure at refresh

- **Every rank row of this round:** they were judged on frozen BiS, and the re-ranked band gear may move the switch levels by a level. This applies most to the 0.2-0.5% race rows (D4, H1, H2, S2).
- **Discipline bands 20 and 30:** the Heal line.

## For other lanes

- **Platform:** thanks for `currentLevel` (bf23e4136 / a7fc87f61). Note X1: a ranked id inside a condition is resolved like a cast.
- **Guide lane:** after the regen, re-capture the Discipline Priority card, which gains the tank Heal row "Heal on the tank until you learn Greater Heal at 40: it is your biggest heal until then.", and check the 3 s / 1.5 s opener.

## For the guide lane

This lane wrote the following, all with wow-player review:
- holy.md: Renew "rank 6 from about 52"; Greater Heal ranks 3 and 4 from 58; leave the shield to a Discipline priest.
- discipline.md: pre-pull Renew; the leveling order (Heal as the tank filler until 40); Renew rank 9 from 56; leave Lesser Heal off the bars in your forties and fifties.
- shadow.md: ranks unchanged, and the race exceptions stay in the rows only.

## Still open

- **D3 at 40-59:** shipped (S 44c6090ac, E f12739f55).
- **D3 at 60 (owner, pending low-rank heal item).** Taking Discipline's Lesser Heal rank 3 out at 60 too, by removing the level gate's upper bound (`currentLevel < 60`), is +30% / +24% at bare 60 and +6.0% / +7.5% at raid 60 (Heal rank 4 in its place, same measure). It is ready to ship once the owner rules that the level-gap penalty applies to heals at 60. discipline.md's "Low ranks" paragraph already says Lesser Heal at 60 heals for little more than its base.
- P2 target-level healing (single source), P3 Eureka! on a live Gnome Shadow log, P4 Lightwell Renew crit.
- The penalty's open parameters (round 16 section 1) and the proc-chance curve.
- **Logs:** foreverlogs.gg answers its API again with browser headers. I last looked at 19:05 UTC, before launch, and no live (level 60) reports exist yet. api.foreversixty.gg/v1/reports/recent holds one 20 October report. Warcraft Logs is blocked (expired login).

## Sources

- Blizzard forums, "[bug?] Healing spells are scaling based on the targets level", 4-19 October 2026: https://us.forums.blizzard.com/en/wow/t/2371681 (fetched 2026-11-04 19:10 UTC)
- Wowhead Forever tooltips: https://nether.wowhead.com/forever/tooltip/spell/1259823 (Eureka!), /6788 (Weakened Soul), /1299631
- Client 70755 tables: SpellEffect, SpellClassOptions, SpellPower, SpellLevels, SpellMisc (scratchpad b70790/70755)
- Lane scratch: `$R/priest/` (coef.txt, out1-out6.txt, race.txt, dheal.txt, ba-disc.txt, var-*.txt, lv-disc.txt, eu/scan.txt, shots/)

## wow-player verdicts

- **Card notes (C1):** SHIP. Two MUSTs on the Discipline card were outside this change: 24 uncollapsed rows (platform fixed it, 9 rows) and the missing Heal fallback before 40 (D1).
- **Prose, first pass:** NO-SHIP. The opener card was stale in the capture (it now renders 3 s / 1.5 s), and the card collapse was outstanding. Its SHOULDs were applied:
  - shorter Renew clause;
  - Greater Heal 58-59 wording;
  - Discipline paragraph order;
  - Renew rank 9 wording;
  - Shadow race clauses dropped;
  - shorter tank Heal note.
- **Prose, second pass:** **SHIP** on all prose (Holy, Discipline, Shadow) and on the opener against the prose.
  - **Its MUST 1:** it read the stale 24-row Discipline capture. The rendered page on this branch has 9 rows (the platform collapse), as checked in the page HTML.
  - **Its MUST 2, applied:** the Discipline tank Greater Heal note now reads "Greater Heal at the top rank as the tank's filler at 60", so it does not clash with the leveling prose's rank 1.
  - **Its SHOULDs, left as they are:**
    - Holy card row 6 does not name the leveling rank 1. The card is the 60 list.
    - Holy and Discipline Renew advice differs by spec. Each is measured.
  - Its doubt about Spirit Tap is outside this round.
- **D3 text (discipline.md leveling, Lesser Heal card note):** SHIP with fixes, all applied.
  - **Its MUST:** name the Renew rank at 50. It now reads "From 50, keep rank 7: rank 8 heals barely more for more mana. From 56, take rank 9." The reviewer's learn levels (46/52) were off. The client has rank 7 at 44, rank 8 at 50 and rank 9 at 56.
  - **Its SHOULDs:**
    - "late forties" and "costs mana regeneration" (Meditation keeps half): applied.
    - The level-60 card still lists Lesser Heal: that is the pending owner item.
