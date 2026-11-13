# Sim accuracy 2026-11-13: Assassination, Combat and Subtlety at 50 and 60 against real fights

Lane brief: `design/reviews/sim-accuracy/brief.md`. Lane `sim-rogues60`. Site from origin/main d730bbcac, fork at the pin 5515ae0ef. Client on the site: 1.60.1.70790; live client 1.60.1.70888 with its hotfix push. Last rogue pass: 2026-11-12 `rogues.md` (0 fixed, MR60).

**Totals: 1 confirmed bug, a tool, fixed: `rotation-search` and `talent-search` simmed every raid band without the published enchants (the same gap the 2026-11-11 casters lane closed in `spec-breakdown`), so every raid rotation and talent search ran 5 to 6% under the set it was judging. No engine bug: every periodic effect in the three rotations rolls crit exactly as its client flag says (the Deep Wounds class of bug does not exist for the rogue), and every log check aligned. No readable rogue fight at 50 or above exists yet: our own 1,178 backfilled fights, foreverlogs.gg and Warcraft Logs hold none. No card, build, order, rank row, guide line or golden changes. Engine pin unchanged: 5515ae0ef.**

## Method

1. **Crit flags.** Every rogue spell (SpellClassSet 8, every rogue talent rank, three hops of trigger spells: 506 spells) with a periodic damage aura or a school damage effect, from the hotfixed 70888 SpellEffect and SpellMisc tables: SPELL_ATTR2_CANT_CRIT (Attributes_2 0x20000000) and SPELL_ATTR8_PERIODIC_CAN_CRIT (Attributes_8 0x200), against the fork's `core/dot.go` rule (`periodicDamageCanCrit`, `NeverCrits`) and each rogue dot's config. Scripts `rogues-60/logs/crit.py`, `crit2.py`.
2. **Our own reports.** The database (read only, through the Cloud SQL proxy): 7 reports, 1,178 fights, uploaded 9 Nov to 13 Nov (the backfill). Each report's `report.json` and the events of every fight with a rogue in it (the two public reports over the site's `/logs-data/`, the three private ones read only from R2): 342 fights, 24,886 rogue events with the advanced fields. Scripts `rogues-60/logs/r2get.py`, `build.py`, `common.py`, `lv.py`, `ipproc.py`. Private reports are counted here in aggregate only; no player from them is named.
3. **foreverlogs.gg.** The public list (26 reports, newest 2854 on 12 Nov): nothing new since the 2026-11-12 rogue and warrior lanes, and nobody at 45 or above (those lanes' event caches).
4. **Warcraft Logs.** The site's v1 key against vanilla, classic, fresh, sod and www: no host lists a Forever zone or partition; `forever.warcraftlogs.com` does not exist. Nothing to read.
5. **Sims** on the Windows box (`tools/remote-test.sh --dir sim --pin-engine`, engine 5515ae0ef, one at a time, `nice -n 15`, GOMAXPROCS=5): `spec-breakdown` at raid 60, bare 60, bare 50 and bare 50 at 55, 5,000 iterations, for the three specs; `rotation-search` and `talent-search` at raid 60 on the raid build (the Arcane lane's fix), before and after the tool fix. Script `rogues-60/run.sh`, output `rogues-60/run.out.txt` and `rogues-60/search.out.txt`.

## No log at 50 or above

| source | highest rogue | note |
|---|---|---|
| our reports (7, 1,178 fights; Desolace, Scarlet Monastery, Hillsbrad, Stormwind, the Wetlands dig) | one rogue with Shank rank 2 (learned at 44) in 46 events, another in 75 | the advanced log's level field is the zone's scaled level, not the character's: a rogue logged at 17 casts Sprint rank 2 (34) and the two Shank rank 2 rogues log 36; the real level comes from the highest rank cast |
| foreverlogs.gg (26 public reports to 2854) | about 43 | unchanged since 12 Nov |
| Warcraft Logs | none | no Forever partition on any host the v1 API serves |

So the 50 and 60 numbers (cast mix, Slice and Dice uptime, Blade Flurry and Adrenaline Rush use, finisher timing, glancing and dodge at +3) stay unmeasured against a log. Every check below is a mechanic that does not depend on the level, read on the richest data we have.

## Crit flags: every periodic effect (the Deep Wounds class)

| effect | client | engine | logs (our reports) | verdict |
|---|---|---|---|---|
| Rupture 1943 to 11275 | PERIODIC_CAN_CRIT on every rank, no Can't Crit | bleed rule: crits on the melee table, x2.0 | 2 critical ticks of 33, each 1.94 to 2.00 of the same rogue's plain tick with an equal base amount | aligned |
| Garrote 703 to 11290 | PERIODIC_CAN_CRIT on every rank | bleed rule, melee table | 21 of 177 ticks crit (11.9%, the melee crit of those rogues), x1.98 to 2.02 | aligned |
| Deadly Poison 2818, 2819, 11353, 11354, 25349 | PERIODIC_CAN_CRIT | a non-channelled dot: crits on the spell table (`DefenseTypeMagic`) | 1 of 72 ticks crit, against the same rogues' 9.6% Sinister Strike crit | aligned (spell table) |
| Death Mark 1316709, 1325756 to 1325758 | Can't Crit (Attributes_2 0x20400000) | the payout is `OutcomeAlwaysHit`; the cap carries the crit chance by the rank text | no log | aligned |
| Devious Strike 1292029 (Defias Leather 4) | PERIODIC_CAN_CRIT | bleed rule | no log | aligned |
| Hemorrhage 16511 | no periodic effect: a Rupture and Garrote amplifier (aura 271) | a debuff, no ticks | 17 hits, 1 crit | nothing to roll |
| Blade Flurry's copy 22482 | | `OutcomeAlwaysHit`, the copied hit's damage | 326 copies, 0 crits | aligned |
| Instant Poison | a direct spell hit | `OutcomeMagicHitAndCrit`: spell crit plus item, talent and buff crit, never Agility | 6 of 416 crit (1.4%) against the same rogues' 9.6% Sinister Strike crit | aligned (the two pools, `core/base_stats.go`) |

No rogue periodic carries Can't Crit except Death Mark, which the engine never rolls. The rogue has no Forever-new bleed: the client's other rogue-set bleeds (Crimson Tempest 412096, Saber Slash 424785, Between the Eyes) are Season of Discovery rows no Forever rogue learns.

## Engine against real fights (levels 17 to about 46)

| # | check | log | engine | verdict |
|---|---|---|---|---|
| L1 | Instant Poison's rate | one Combat rogue, 1,541 landed weapon hits: 18.2% of white hits (main and off hand together; the log sets no off-hand flag on a landed swing), 14.4% of yellow; 18% in all | 20%, plus Improved Poisons | held, as the physical lane held it: poison uptime is not in the log |
| L2 | Flash of Steel | three rogues gain 30 from Shank rank 1 and 40 from rank 2 every time (35 once at the cap), the rest 15, 18 or 25; the shortest gap between two Shanks at 30 or 40 is 7.93 s | 15 or 25 plus 3 a rank, -0.4 s a rank: 5/5 gives +15 and 8.0 s | aligned |
| L3 | Puncturing Wounds | 8 extra combo points (1310710) on 15 landed Backstabs from three rogues | 15% a rank, 45% at 3/3 | consistent (8 of 15 at 45% is a one in three draw) |
| L4 | Relentless Strikes | 1314102 gives 25 Energy | 25 | aligned |
| L5 | combo points per builder | Sinister Strike 1 (1,300), Hemorrhage 1, Garrote 1, Ghostly Strike 1, Cheap Shot 2, Initiative 1, Ruthlessness 1; Sinister Strike's negative and zero gains are finisher spends and the 5-point cap | the same | aligned |
| L6 | glancing, dodge and the boss | not readable: a logged mob's level is not in the log and the rogues' own level field is scaled | Era table at +3 (40% glance, 4.5 to 6.5% dodge) | open, needs a level 60 boss log |

## Sims at 50 and 60 (reference, 5,000 iterations, seed 1)

| spec | raid 60 | bare 60 | bare 50 | bare 50 at 55 |
|---|---|---|---|---|
| Assassination | 917.65 ± 0.64 | 344.31 ± 0.24 | 220.60 ± 0.12 | 250.28 ± 0.14 |
| Combat | 902.21 ± 0.53 | 313.77 ± 0.21 | 217.93 ± 0.12 | 242.81 ± 0.14 |
| Subtlety | 890.81 ± 0.50 | 407.19 ± 0.27 | 236.62 ± 0.14 | 299.46 ± 0.17 |

Raid 60 Assassination: Slice and Dice 91.3% up, Venom 59.6%, Deadly Poison crits 30.4% and Instant Poison 30.0% against 35.3% white (the raid set's item crit, Malice and Moonkin Aura reach both pools, Agility does not), white glance 39.9% and dodge 4.5%. These are the numbers a level 60 log will be checked against.

## T1: the searches wore no enchants (CONFIRMED-BUG, tool, fixed)

**Finding.** `sim/cmd/rotation-search` and `sim/cmd/talent-search` built their gear from the band's `item_id`s alone (`bisBand.gear()`), while the ranker measures every raid set wearing its permanent enchants and the published slots carry them as `enchant_id`. The 2026-11-11 casters lane fixed `spec-breakdown` (T1 there) and left the two searches. The rogue raid searches showed it: the curated baselines read 869.6 (Assassination), 851.3 (Combat) and 841.1 (Subtlety) where `spec-breakdown` on the same bands reads 917.65, 902.21 and 890.81 (-5.2%, -5.6% and -5.6%).

**Fix.** Both tools read `enchant_id` and wear it (`Enchant` on each `api.GearSlot`), as `spec-breakdown` does. Red first: `TestBISBandGearWearsItsPublishedEnchants` in both packages fails at d26f86d06 (`gear enchants map[head:0 main_hand:0]`) and passes at the fix.

**Reach.** No published number moves: the searches only judge rotations and builds. Every raid-60 rotation and talent search before this landing, for every spec, compared candidates on bare items; their deltas were relative and in this lane's rogue re-run they stand (below). A spec whose raid build leans on an enchant-sensitive stat (a weapon proc such as Crusader, an Agility or hit enchant near a cap) should have its raid searches re-run by its next lane; listed under Open.

**The rogue searches on enchanted gear (800 iterations to confirm, seed 7, raid build):**

| search | before (bare items) | after (published enchants) | verdict |
|---|---|---|---|
| Assassination rotation | 869.6 ± 1.6, curated best | 919.0 ± 1.6, curated best (+0.1, ±2.3) | holds |
| Combat rotation | 851.3 ± 1.2, curated best | 901.3 ± 1.3, curated best (+0.0, ±1.9) | holds |
| Subtlety rotation | 841.1 ± 1.2, curated best | 891.6 ± 1.2, curated best (+0.0, ±1.8) | holds |
| Assassination talents | guide best | guide 919.0; best variant +0.2 (±2.2) | holds |
| Combat talents | guide best | guide best (901.3) | holds |
| Subtlety talents | guide best | guide 891.6; best variant +1.3 (±1.7), within error | holds |

The enchanted baselines now sit within error of `spec-breakdown` on the same bands (917.65, 902.21, 890.81 at 5,000 iterations, seed 1).

## Builds, rotations, orders and rank rows

Nothing to change: the raid rotation searches on the raid builds keep every curated rotation, the raid talent searches keep every raid build (above), and no log at 50 or above exists to move a leveling order or a card ruling. Rogues carry no rank rows (every rank costs the same Energy).

## Held

- Instant Poison's rate (L1): 18% of landed hits against 20%; poison uptime unknown.
- Subtlety's third Improved Slice and Dice point (2026-11-12, bare 60 +0.5%, under the 1% bar).

## Open

- Every 50 and 60 check against a log (L6 and the cast mix): needs a rogue at 50 or above in our reports, foreverlogs.gg or a Warcraft Logs Forever partition. The advanced log's level field is the zone's scaled level; a lane reading levels must take the real one from the highest rank cast.
- Other specs' raid rotation and talent searches on enchanted gear (T1 reach).
- Mutilate's off-hand strike, Death Mark's storage and cap, Ghostly Strike: still no log.
- MR60 (Poised under 70888) stands.

## Sim changes in this landing

- **Fork:** none. Pin 5515ae0ef unchanged.
- **Site:** d26f86d06 (red), 18d046e19 (the two search tools wear the published enchants); this report, `rogues-60/` (run script, outputs, log scripts) and the ledger row. No card, guide, coach note, build, order or golden changes, so no wow-player round.

## Sources

- wago.tools 1.60.1.70888 with `useHotfixes=1` (SpellEffect, SpellMisc, SpellClassOptions, SpellLevels, SpellName), the 2026-11-12 lane cache.
- Our reports: `fights` and `reports` (read only); `report.json` and `fights/<n>/events.parquet` for ctxicu4cuyek, 4q5el2xsgoy7 (public) and three private reports.
- foreverlogs.gg `/api/reports/public` (fetched 13 Nov).
- Warcraft Logs v1 `/v1/zones` on vanilla, classic, fresh, sod and www (fetched 13 Nov).
