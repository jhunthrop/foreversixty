# Sim-accuracy round 24: Druid

## Lane

- Site: wt/r24-druid from int/r24 f05c281fe (int/r24 merged in at the end). Fork: lane/r24-druid from int/r24 3b4d6b054 (int/r24 merged in at the end).
- Client 1.60.1.70790. Blizzard's notes thread 2360696 re-read 5 Nov ~14:20 UTC: still 10 posts, the last 4 Nov 19:35 UTC, nothing new for druids. "Info & Known Issues" (2396179): 4 posts, no class lines.
- Wago hotfix pushes on 70790, region 91, every SpellEffect row (fetched 5 Nov ~14:30 UTC, 933 records), checked against the cached 70790 client rows for every druid spell and talent: the only difference is spell 21178 (Bear Form's passive threat aura), 50 to 115, push 113326. The engine already has Bear Form at 115% (round 21, Blizzard 3 Nov). Aligned. The committed `data/builds/1.60.1.70790/hotfix-push.json` is stale: every `search_totals` entry is 0 (last written e65ee3982, 4 Nov), so it cannot be used for this check (integrator).
- Logs (5 Nov ~14:24 UTC): api.foreversixty.gg lists only the 20 Oct report; foreverlogs.gg 403; Warcraft Logs v1 401 on classic, fresh and www.
- Method: spec-breakdown on frozen BiS (`$R/bis-frozen`) and the shipped leveling order at each level (`ladder.py`), bare preset, `-level N`, 10,000 iterations, seeds 1 and 50001, night elf (alliance) and tauren (horde). Exploration at 3,000. Ranks forced with `-rank id=r`; where two spells share a mana bar (Wrath and Moonfire; the four Restoration heals) they were judged together on a joint grid, as round 20 did. Every new row was then confirmed at seed 50001 by putting the old rank back (leave-one-out).

The rank sweep (round 24 seed: the first two levels after every new rank of every spell the cards cast) is where every finding came from. Below 20 nothing had been swept level by level before: the fights run out of mana, so the cheap ranks win, and the factions' band-20 gear differs enough to split some rows.

## Balance

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| B1 | Wrath and Moonfire below 20 cast the newest rank. The fight runs out of mana at 10 to 19, so the cheap ranks deal more: Wrath rank 1 at 10 to 12 (+3 to 24%), Wrath rank 2 at 14 to 18 (+3 to 29%), Moonfire rank 2 at 16 and rank 1 at 17 to 19 (+5.6 to 7.6% over rank 3). | `sw/bal-grid.tsv` (joint Wrath x Moonfire grid), trainables Wrath 1/6/14/22, Moonfire 4/10/16/22 | CONFIRMED-BUG (rank rows) | site rank-choices (never swept below 20) | S af0fb2481 |
| B2 | Moonfire rank 4 row started at 36; rank 4 already beats rank 3 at 35 (+0.3 to 0.6%, both seeds and factions). | `sw/mid.tsv` L33 to 37 | CONFIRMED-BUG | site rank-choices (round 19 row edge) | S af0fb2481 |
| B3 | Every other Balance rank row at the first two levels after each new rank (Wrath, Moonfire, Starfire, Insect Swarm, 20 to 60). | `sw/x3b.tsv`, `sw/mid.tsv` | aligned | - | - |

## Feral (cat)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| F1 | The card's pre-Cat Form lines (Moonfire, then Wrath, 10 to 19) cast the newest ranks and run dry. Wrath rank 1 and Moonfire rank 1 at 10 to 13 gain 8.6 to 11.6%; the newest ranks win at 14 and 15; Moonfire rank 2 beats the rank 3 learned at 16 by 3.9 to 13.7% at 16 to 19. | `sw/fer-grid.tsv` (joint grid), `sw/low.tsv` | CONFIRMED-BUG (rank rows) | site rank-choices | S 2ae0fe3b8, E f66d818b3 (APL copy) |
| F2 | Cat melee spells (Claw, Rake, Shred, Rip, Ferocious Bite, Ravage): each rank costs the same energy and does strictly more, so the newest rank wins; the casts are the newest rank at every band. | `sw/feral-casts.txt`, spellconst | aligned | - | - |
| F3 | Wolfshead Helm re-audit: Enrage is effect 2 = instant 10 rage (+5 with the helm); Shifting Power 40 energy (+5). | client SpellMisc/SpellEffect rows | aligned | - | - |

## Feral bear

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| T1 | Card row [8] (Swipe) note: "on a pack of three it is your main threat". On three targets on the shipped order Maul is the top threat at every band: bare 30 Maul 36.9% vs Swipe 36.1%, bare 50 39.3% vs 19.8%, raid 60 51.9% vs 18.7% (seeds 1 and 50001 within 0.1). Guide-accuracy round 26 R26-2 (LIKELY) asked for this breakdown. New note: "It hits up to three, but Maul stays your biggest threat on a pack too; on a single boss, put the rage into Maul instead, and into Lacerate from 25 and Primal Bite from 40." | `sw/bear3.txt` (10,000 iterations, `-targets 3 -tank`, alliance) | CONFIRMED (card text; no sim change) | site card note | S 66eae8d64, E 2abc28803 |
| T2 | Bear melee spells (Maul, Swipe, Demoralizing Roar, Lacerate): newest rank at every band, same reasoning as F2. | `sw/bear-casts.txt` | aligned | - | - |

## Restoration

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | Below 20 the four heals (tank Rejuvenation, party Rejuvenation, Healing Touch filler, Regrowth) cast ranks no one had swept. Judged together on a joint grid: party Rejuvenation keeps the newest rank (rank 2) at 10 to 12, not rank 1 (+0.6 to 1.4% night elf, +13 to 16% tauren); from 13 to 16 a night elf takes rank 1 on the tank and the party, while a tauren keeps the newest party rank (rank 2 to 15, rank 3 at 16: +22%); Healing Touch rank 2 on the hole at 14 to 16; from 17 to 19 rank 1 of everything, Regrowth rank 1 included once rank 2 arrives at 18 (+3.7 to 4.7% alone). The set gains 0.6 to 25% by level (Before/after). | `sw/res-grid.tsv` (seed 1), `sw/loo50001.txt` (seed 50001: reverting any new row loses 0.4 to 20%), `sw/ba-resto.txt` | CONFIRMED-BUG (rank rows) | site rank-choices | S 9ba42a787, E 951f3d051 (APL copy) |
| R2 | A tauren's Healing Touch filler: rank 4 beats rank 5 (newest, learned 26) at 26 to 28 by 0.2 to 0.5%. | `sw/mid.tsv` L26 to 29 | CONFIRMED-BUG | site rank-choices | S af0fb2481, E 4f8231714 |
| R3 | Healing Touch filler at 32 and 33 (rank 6 new at 32): the card's rank 3 holds; rank 4 -0.2 to -0.4%, rank 5 about -1%, rank 6 -1.2 to -1.5%, both seeds and factions. Also 38/39, 44/45, 50/51, 56/57: card holds. | `sw/extra.tsv`, `sw/x3b.tsv` | aligned | - | - |
| R4 | Tranquility at 30/31, 40/41, 50/51 and 60 (bare and raid): the top rank wins or ties (lower ranks -0.3 to -2.1% where it fires). | `sw/extra.tsv` | aligned | - | - |
| R5 | Rejuvenation and Regrowth at the first two levels after every rank from 20 to 60. | `sw/x3b.tsv` | aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | Hotfix pushes: only Bear Form's threat aura (21178) differs from the client, already in the engine. | wago region 91, push 113326 | aligned | - | - |
| X2 | `hotfix-push.json` on the branch shows 0 rows for every table: stale, not a sim bug. | file `search_totals` | for the integrator | data | - |

## Before/after

Bare preset, shipped leveling order, frozen BiS, 10,000 iterations. Before = `$R/sb-before` on f05c281fe; after = this branch. Relative % for night elf (A) and tauren (H); seeds 1 and 50001 agree within 0.2 points everywhere. Absolute at seed 1 (DPS, or effective HPS for Restoration), before -> after. Full table: `r24/druid/sw/ba-all.txt`.

| spec | level | A | H | A seed 1 | H seed 1 |
|---|---|---|---|---|---|
| Balance | 10 | +22.0% | +24.3% | 4.81 -> 5.86 | 3.58 -> 4.45 |
| Balance | 11 | +15.8% | +17.7% | 5.44 -> 6.30 | 4.10 -> 4.82 |
| Balance | 12 | +4.1% | +3.2% | 6.53 -> 6.80 | 5.26 -> 5.43 |
| Balance | 13 | 0 | 0 | 7.48 | 6.11 |
| Balance | 14 | +25.5% | +28.6% | 6.65 -> 8.35 | 5.76 -> 7.40 |
| Balance | 15 | +22.7% | +24.9% | 9.51 -> 11.68 | 8.44 -> 10.54 |
| Balance | 16 | +26.5% | +25.2% | 10.09 -> 12.76 | 8.61 -> 10.77 |
| Balance | 17 | +23.2% | +16.8% | 13.49 -> 16.62 | 12.69 -> 14.82 |
| Balance | 18 | +12.3% | +12.3% | 18.14 -> 20.37 | 16.40 -> 18.41 |
| Balance | 19 | +7.5% | +7.6% | 21.92 -> 23.56 | 19.97 -> 21.48 |
| Balance | 35 | +0.4 to 0.6% | +0.3% | 86.48 -> 86.98 | 85.87 -> 86.16 |
| Feral | 10 | +11.1% | +11.6% | 4.46 -> 4.95 | 4.37 -> 4.88 |
| Feral | 11 | +9.8% | +10.6% | 4.64 -> 5.09 | 4.53 -> 5.01 |
| Feral | 12 | +8.7% | +9.3% | 4.86 -> 5.28 | 4.74 -> 5.18 |
| Feral | 13 | +8.6% | +8.7% | 4.97 -> 5.40 | 4.87 -> 5.29 |
| Feral | 14, 15 | 0 | 0 | 6.42, 6.66 | 6.43, 6.66 |
| Feral | 16 | +4.1% | +4.7% | 6.49 -> 6.76 | 6.47 -> 6.78 |
| Feral | 17 | +6.7% | +7.9% | 9.19 -> 9.81 | 9.27 -> 10.00 |
| Feral | 18 | +13.6% | +12.0% | 12.80 -> 14.53 | 12.99 -> 14.55 |
| Feral | 19 | +11.8% | +9.9% | 13.46 -> 15.05 | 13.68 -> 15.03 |
| Restoration | 10 | +1.4% | +13.7 to 15.1% | 13.8 -> 14.0 | 7.3 -> 8.3 |
| Restoration | 11 | +0.7% | +16.0% | 14.2 -> 14.3 | 7.5 -> 8.7 |
| Restoration | 12 | +0.6% | +12.8% | 15.4 -> 15.5 | 8.6 -> 9.7 |
| Restoration | 13 | +1.3% | +13.6% | 15.8 -> 16.0 | 8.8 -> 10.0 |
| Restoration | 14 | +1.9% | +11.0% | 16.1 -> 16.4 | 9.1 -> 10.1 |
| Restoration | 15 | +3.9% | +1.6% | 20.6 -> 21.4 | 12.8 -> 13.0 |
| Restoration | 16 | +5.3% | +24.7% | 24.5 -> 25.8 | 15.0 -> 18.7 |
| Restoration | 17 | +12.0% | +7.3% | 39.9 -> 44.7 | 24.6 -> 26.4 |
| Restoration | 18 | +18.5% | +17.5% | 43.2 -> 51.2 | 40.5 -> 47.6 |
| Restoration | 19 | +17.8% | +17.1% | 43.7 -> 51.5 | 41.0 -> 48.0 |
| Restoration | 26 to 28 | 0 | +0.3 to 0.5% | 76.1 | 75.8 -> 76.0 |

Feral bear: no change (note text only).

Raid 60 is unchanged for every druid spec: no row at 60 moved and the fork commits carry only APL note text.

## Files to regenerate

- `data/builds/1.60.1.70790/addon-data.json` (rotations/druid-balance, druid-feral, druid-feral-bear, druid-restoration: new rank rows and note text).
- The fork's APL copies are committed by hand (ui/*_druid/apls); apl-sync should find nothing to change.
- Coach starters only if the integrator's persona_starter.py --all picks up the card notes (no druid persona line names these ranks).

## Measure at refresh

- None from this lane: no BiS-dependent pick changed.

## For other lanes

- No commit here moves DPS for any spec outside the druid class.
- Every caster lane: below 20 the bare fights run out of mana, and the faction's band-20 gear differs enough to split rows by race. Sweep 10 to 19 level by level, with spells that share the mana bar judged together.

## For the guide lane

- Guide-accuracy round 26 R26-1 (feral-bear.md:83, "Swipe moves to the top on a pack") is still on the page; round 22 G2's numbers stand. Not touched here.
- R26-2 is closed by T1 (card note rewritten).

## Still open

- Exclusivity in practice (Leader of the Pack and Moonkin Aura): a log with both druids in range.
- Bear threat constants (Maul, Swipe, Lacerate, Primal Bite): a live tank threat capture.
- Eclipse on a missed Wrath; the unarmed paw: a live Balance log at 40+ and a bare-handed cat.
- The exact Rip and Bite attack power shares.
- Logs: none reachable (above).
- Pending owner, not reopened: the level-60 book ranks (Ferocious Bite 6, Healing Touch 11, Rejuvenation 11, Starfire 7).

## Sources

- Blizzard, "WoW Forever Beta Development Notes", https://us.forums.blizzard.com/en/wow/t/2360696; "Info & Known Issues", https://us.forums.blizzard.com/en/wow/t/2396179.
- wago.tools build 1.60.1.70790 hotfix pushes (region 91), SpellEffect.
- `data/builds/1.60.1.70790/trainables/druid.json`, `spellconst/druid.json`, `addon-builds.json`.
- Guide-accuracy round 26 druid report (gr26 scratchpad, `design/reviews/guide-accuracy/round26/druid.md`).
- Sweep data in the lane scratchpad `r24/druid/sw/`.
