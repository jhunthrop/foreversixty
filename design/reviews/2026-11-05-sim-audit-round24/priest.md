# Sim-accuracy round 24: Priest

Site branch wt/r24-priest (on int/r24 7f7b29fa5, which carries origin/main c52b4d51d). Fork lane/r24-priest has no priest commits this round (merged int/r24 only). Engine and site built at the merge; `spec-breakdown -level N` uses the player's talents at N (platform S 50235426e) and a forced `-rank` wins over every row (S fcceada57).

Method for every rank row: `spec-breakdown -band <floor> -level N -preset bare -race <race> -rank <authored id>=<rank> -heal|-json -bis-dir $R/bis-frozen`, 10,000 iterations, seeds 1 and 50001. A row is taken on a gain beyond error at both seeds (heals: above 0.1%, the printed error is 0.05 HPS on 90 to 300; Shadow: above 0.15%, the two seeds differ by up to 0.1%). Unkeyed rows were judged on the published races (Human, Undead) and then on the other four races wherever a change was made; race-keyed rows on their own races. Rows that interact at a level (Greater Heal with Renew at 58 and 59; Mind Blast with Devouring Plague at 41) were judged jointly. Screen first (3,000 iterations, sampled levels, `$R/priest/screen3k.tsv`), then every level of each suspect range at 10,000 (`$R/priest/fill10k.tsv`, `combo5859.tsv`, `combo2.tsv`, `combo3.tsv`; tools in `$R/priest/tools/`, `verdict.py` prints each comparison).

## Discipline

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| D1 | **Alliance race pick is Human, not Gnome.** Human heals more at every band: bare 20 +1.1%, 30 +3.6%, 40 +2.9%, 50 +4.6% (198.57 against 189.86), 60 +0.7%, raid 60 +0.1% (530.42 against 530.00, a tie). The Gnome pick dates from 3 Nov, before the level-gap penalty took Lesser Heal rank 1 out of the leveling fight; at bare 50 the Human Spirit's 10% Spirit carries the mana (Human 30.8 mana/s income against 29.4, 57.6 s against 54.8 s of mana). Raid 60 ties, so the leveling bands decide. Guide front matter, the race section and index.md follow. | `$R/priest/race-disc.txt`, `race50.txt` (10,000 its, seeds 1 and 50001, frozen BiS, guide build) | CONFIRMED-BUG | old (guides r19 pick; healbisect lane flagged it this round) | S 846b26e2e |
| D2 | **Renew 34 to 35: rank 2, not rank 3.** At 34 every race heals more with rank 2 (+0.8% Human to +1.6% Troll); at 35 every race but a Human (+0.2% to +0.6%; a Human ties). Rank 3 from 36 (rank 2 -0.6% to -1.7%). | fill10k.tsv | CONFIRMED-BUG | old (round 16 row, band-floor talents) | S 92ba8e5db |
| D3 | **Renew 40 to 49 by race.** At 42 a Human (+0.4%) and Undead (+0.1%) cast rank 5; a Troll heals most on the default rank 6 (+1.3%). At 44 rank 5 for a Dwarf, Gnome, Night Elf or Undead (+0.4% to +0.8% over rank 4), a Troll rank 7 (+1.1%). From 45 to 47 a Gnome or Troll casts rank 7 (+0.2% to +0.5%), an Undead rank 6 at 46 and 47 (+0.1% to +0.3%). At 48 and 49 a Human (+0.4% to +0.6%) and Troll (+0.3%) cast rank 7. Rank 4 at 40, 41 and 43 holds for every race. | fill10k.tsv | CONFIRMED-BUG | old (round 18/19 rows) | S 92ba8e5db |
| D4 | **Renew 56 to 58: rank 8 for a Dwarf or Night Elf.** Dwarf +0.2% to +0.6% at 56 to 58, Night Elf +0.2% to +0.4% at 56 and 58, Gnome +0.2% at 57; a Human keeps rank 9 (rank 8 -0.5% to -1.1%). 59 holds (rank 9; within error for Dwarf, Night Elf, Undead). | fill10k.tsv, combo2.tsv (58 and 59 with Greater Heal rank 1) | CONFIRMED-BUG | old (round 19 rows) | S 92ba8e5db |
| D5 | **Heal at 36: an Undead keeps rank 1** (+0.3% both seeds); every other race loses 0.1% to 0.4% on rank 1. 37 to 39 rank 3 holds (Undead rank 2 +0.1%/+0.06% at 39, within error). | fill10k.tsv | CONFIRMED-BUG | old (round 22 row) | S 92ba8e5db |
| D6 | **Greater Heal rank 1 through 59, not rank 4 from 58** (rank sweep at the two levels after a new rank): at 58 rank 1 beats rank 4 by +0.5%/+0.6% (Human) and +0.4%/+0.5% (Undead); at 59 by +0.3% and +0.2%. Not on the suspect list: no row covered 58 and 59. discipline.md:124 said rank 1 "until you learn rank 4 at 58". | combo5859.tsv | CONFIRMED-BUG | old (round 18 stopped the row at 57) | S 92ba8e5db |
| D7 | Prayer of Healing 40 to 49 rank 1, Greater Heal 46 to 57 rank 1, Heal 22 to 35 rank 1, Renew 14 to 33, Power Word: Shield 36 to 40 and 54 to 55 race rows: hold on the level's talents (screen within 0.1%, no challenger beyond error). | screen3k.tsv, fill10k.tsv | Aligned | - | - |
| D8 | Guide-r23 sim 3, the shield with no health gate, re-measured: raid 60 530.42 / 530.05 (Alliance) and 525.95 / 526.19 (Horde) drop to 525.89 / 525.80 and 521.09 / 521.27, -0.8% to -0.9%. The 90% gate stays. | `$R/priest/g24.txt` | Aligned | - | - |

## Holy

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| H1 | **Renew 28 to 35 race by race.** Rank 1 holds through 27 for every race. Rank 2 takes over as the talents arrive: Undead at 28 and 29 (+0.5%, +1.7%), Dwarf at 29 (+2.4%), Dwarf, Gnome, Night Elf and Troll at 30 (+2.8% to +13.4%: on the band-30 talents rank 1 runs dry), Troll and Undead at 31 (+0.6% to +3.9%; a Human rank 3, +0.5%), every race but Undead at 32 (+0.5% to +5.1%), every race but Dwarf and Human at 33 (+0.1% to +3.1% over rank 3), Human and Troll at 34 (+0.4% to +0.7%), Human at 35 (+0.6%). Rank 3 from 36. | fill10k.tsv | CONFIRMED-BUG | old (round 16 rows, band-floor talents) | S 92ba8e5db |
| H2 | **58 and 59: Greater Heal rank 1 and Renew rank 7.** Judged jointly (Greater Heal 1/3/4 x Renew 6/7/8/9). At 58 Greater Heal rank 1 beats the top rank by +1.6% (Human 287.77 against 283.37 with Renew 6; Undead 301.07 against 296.34), and a Human gains another +1.5% with Renew rank 7 (292.12); the other races keep Renew rank 6 (within error or a loss). At 59 rank 1 with Renew rank 7 heals +4.0% (Human 308.08 against 296.04) and +4.1% (Undead 321.27 against 308.67) over the round-18 pair (Greater Heal 3, Renew 6); Dwarf, Gnome, Night Elf and Troll gain +3.0% to +3.4% from Renew rank 7 too. holy.md:128 said rank 1 "until you learn rank 4 at 58"; the card note said "rank 1 until 57". | combo5859.tsv, combo3.tsv | CONFIRMED-BUG | old (round 18 rows) | S 92ba8e5db |
| H3 | Renew 40 to 57, Heal 22 to 39 race rows, Prayer of Healing 40 to 59 rank 1, Flash Heal 32 Undead: hold on the level's talents. Renew rank 7 at 55 to 57 loses 5% to 8%. | screen3k.tsv, fill10k.tsv | Aligned | - | - |
| H4 | Guide-r23 sim 4, Greater Heal rank 1 behind the pace at 60, re-measured: raid 60 771.17 / 771.14 (Alliance) and 766.02 / 766.90 (Horde) drop to 767.68 / 767.99 and 760.75 / 761.52 (-0.4% to -0.7%). The top rank stays. (Round 23's 788 Horde base predates a round-23 change; before.tsv has 766.0, so nothing moved this round.) | `$R/priest/g24.txt` | Aligned | - | - |

## Shadow

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **Mind Blast 40 to 43 by race.** Rank 5 beats the row's rank 4 at 40 for a Human (+0.5%), Gnome (+0.3% over the top rank) and Undead (+1.2% to +1.3%); at 41 for a Dwarf (+0.6%), Gnome (+1.1%) and Human (+2.4%), with an Undead on the top rank 6 (+0.8%); at 42 for a Gnome (+1.0%) and Undead (+0.2%). Rank 4 holds elsewhere; a Gnome takes rank 6 at 43. | fill10k.tsv | CONFIRMED-BUG | old (round 21/22 rows) | S 9c67b212c |
| S2 | **Mind Blast 50 to 53.** At 50 and 51 a Dwarf, Night Elf or Troll casts rank 5 (+0.3% to +0.8% over rank 7), a Dwarf at 52 too (+0.2% to +0.3% over the new rank 8). The Undead rank 6 row holds at 50 but loses at 51 (rank 7 +0.3%); the Dwarf rank 6 row at 51 loses to rank 5 (+0.7% to +0.8%). A Human casts rank 6 at 53 (+0.3%/+0.4%). Rank 5 loses 2% to 3% for a Dwarf at 47 to 49 (band 40 gear). | fill10k.tsv | CONFIRMED-BUG | old (round 21 rows) | S 9c67b212c |
| S3 | **Devouring Plague: Troll rank 2 at 29** (+0.3%/+0.4%) and **Human rank 3 at 49** (+0.3%). A Human's rank 2 at 41 wins alone (+0.3%) but with Mind Blast rank 5 at 41 the pair loses (93.41 against 94.42), so it is not taken. The other Devouring Plague rows, Shadow Word: Pain 26 to 39 and Shadow Word: Death hold. | fill10k.tsv; joint check at 41 | CONFIRMED-BUG (two rows) | old | S 9c67b212c |
| S4 | Guide-r23 sim 2, Shadow Word: Pain at 6 against 10 seconds left: raid 60 724.3 / 724.8 (Alliance), 717.5 / 717.2 (Horde) against 724.4 / 724.9, 717.6 / 717.3: within error, no change. | `$R/priest/g24.txt` | Aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | Every rank row changed here was set on band-floor talents (the tool bug the hunter lane found this round), so their origin is `old`. The new rows cluster where a talent arrives between band floors (Holy 28 to 35, Discipline 42 to 49), as the integrator's decision predicted. | platform-suspect-rank-rows.tsv | - | - | - |
| X2 | Logs: api.foreversixty.gg/v1/reports/recent (07:42 UTC 6 Nov) lists one report from 20 Oct; foreverlogs.gg answers 200 with no priest data. Round 22's open items (heals scaling with target level, Eureka! on a Gnome Shadow log, Lightwell Renew crit, Dark Inspiration Spirit base) still need a level 40+ priest log. | curl | OPEN | - | - |

## Before/after

All changes are leveling rows (no row at 60 moves), so raid 60 and bare 60 are unchanged. Per changed level, relative gain at seeds 1 / 50001 (absolute at seed 1 in the tables above and in fill10k.tsv): Discipline +0.1% to +1.6%, Holy +0.5% to +13.4% (Renew at 30) and +4.0% at 59, Shadow +0.2% to +2.4%. Each new row was checked to apply by default (no -rank): Disc 42 Troll 133.62, Disc 36 Undead 94.54, Disc 58 Dwarf 204.04, Holy 59 Human 308.08, Holy 58 Human 292.12, Holy 30 Night Elf 69.34, Shadow 41 Human 94.42 / 94.45, Shadow 50 Dwarf 208.59, Shadow 49 Human 178.51.

## Files to regenerate

- addon-data (rotations/priest-* rank lines and race_spell_ids follow rank-choices.json).
- coach starters (persona_starter.py --all): data/curated/coach/priest-discipline.md:32 repeats the Heal note, now "(36 as an Undead)".
- `sim/rankchoices/rank-choices.json` is copied by hand (make rank-choices-sync); `sim/request/apl/priest-{holy,discipline}.apl.json` notes edited by hand to match the curated APLs.

## Measure at refresh

- The Discipline race pick at the BiS refresh (Human at every band; raid 60 ties).
- Holy and Discipline 58 to 59 on refreshed gear (Greater Heal rank 1 margins are 0.2% to 1.6%).

## For other lanes

- Platform: `TestCuratedRankChoicesAreCastableDownranks` still fails on int/r24 for paladin-protection Consecration rows 49 and 58 (not priest).

## For the guide lane

- `tools/coach/starter_overrides.py` Discipline "ranks" drill (round 26 R26-6 text) and the Holy persona `data/curated/coach/priest-holy.md` (R26-7/R26-8 hand edits) say Greater Heal rank 1 until rank 4 at 58; rank 1 now holds to 60.
- Round 26's optional measurements are the four routed sims (Prayer of Mending per cast, Shadow Word: Pain at 6 seconds, Discipline shield with no health gate, Holy Greater Heal rank 1 behind the pace); the last three were re-measured here (g24.txt), Prayer of Mending carries round 23's answer (no priest engine change since).

## Still open

- X2 (logs).
- Rank rows for races outside Human and Undead on unchanged unkeyed rows were not re-judged where the screen showed no loss for the published races.
