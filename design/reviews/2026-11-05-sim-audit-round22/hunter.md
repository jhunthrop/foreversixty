# Round 22, lane hunter: Beast Mastery, Marksmanship, Survival (melee)

Client 1.60.1.70790 (no hotfix rows). Site wt/r22-hunter from int/r22 (3337ca488, later merged with int/r22 591df0a19 for the druid lane's crit change); fork lane/r22-hunter from int/r22 (69c6325fb, later merged with a03f8512c).
Method: spec-breakdown through `$R/sim.sh`; guide builds (`$R/rt`) and `-bis-dir $R/bis-frozen` at raid 60; bare bands with `-level` for rank rows. 2000-3000 iterations to explore, 10000 at seeds 1 and 50001 for every reported number. Raid 60 measured against this branch's own int/r22 baseline (`$R/hunter/sb-base`), not `$R/before.tsv`, after the druid lane's Leader of the Pack change.

**Result: 4 CONFIRMED-BUGs, all fixed: 1 `R21`, 2 `R20`, 1 `old`.** The largest is a rank finding: since round 20 put an attack power share on Arcane Shot and Serpent Sting at every rank, a low rank keeps most of its hit for a fraction of the mana, and a leveling hunter who fires both on cooldown was running dry for over a minute of every 180 s fight on the newest ranks. Round 20 asserted the share "cannot flip" a rank and never swept. Bare bands move up to +10.6% (Marksmanship 50, with the Aimed Shot row) and +9.9% (Beast Mastery 38); raid 60 does not move. The seed (Beast Mastery's order) is decided with no change: no gate separates the cases, and the raid gain turns into a loss on long fights.

## Seed: Beast Mastery's order

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| B0 | **Multi-Shot above Aimed Shot: the talent builds do not separate the cases, the gear and mana do.** Horde, 3000 it, seed 1: raid preset with the raid build 890.9 -> 898.7 (+0.9%), raid preset with the leveling build 876.1 -> 884.3 (+0.9%); bare 60 with the raid build 434.0 -> 413.2 (-4.8%), with the leveling build 433.0 -> 409.9 (-5.3%). A raid-only talent gate (Rapid Killing, Mortal Shots 4) would be a false proxy. Why it differs: at raid 60 the hunter never runs dry and casts rank 6 Aimed Shot (310 mana, 2 s cast, 1110 average) whose cast costs about five Arcane Shots a fight; bare 60 runs dry for 61 s and casts rank 1 (75 mana). Gates tried: mana floors 40-80% (raid gain and bare loss trade off monotonically, none separates); a cost gate (Aimed Shot's current cost above Multi-Shot's, which the rank rows already flip, plus level 60) separates exactly (bare bands identical to the digit, raid +0.9%). But that line wins only on short fights: raid 60 at 90 s +0.2%, 180 s +0.9%, 420 s -0.7% (with a 40% / 60% mana floor: 180 s +0.7% / +0.1%, 420 s -0.4% / -0.3%). Re-run after the merge of the druid lane's crit change: 180 s +0.8%, 420 s -0.6%. A rule that loses on a long boss is not one to print on the raid card. | `$R/hunter/rot/bm-{now,mf,mg*,cost,c60,c60m*}.json`, `cmp.sh` | Decided: keep Aimed Shot first | - | - |

## Beast Mastery

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| B1 | **Arcane Shot and Serpent Sting ranks while leveling.** Blizzard's 3 November notes: Arcane Shot "Attack Power coefficient raised from 0.1 to 0.15", Serpent Sting "Attack Power coefficient per tick increased to 5%"; the engine adds both on every rank (`attack_power_scaling.go`, round 20 E 628f99a7c), and the October beta logs fit the share on ranks 3 and 4 (round 20 H4). So rank 1 Arcane Shot (25 mana) keeps most of the damage of rank 8 (190 mana). Every rank swept at 20-60 (2000 it, Horde), then the rows at 10000 it, seeds 1 and 50001: Arcane Shot rank 1 at 26-49, rank 3 (trained at 20) at 50-59 and bare 60; Serpent Sting rank 1 at 26-39, the newest rank from 40 (it wins there by 1 to 3%). At 20-25 every rank ties (no row). The breakdown confirms 3044 and 1978 are cast; bare 60 dry time falls from about 60 s to a few seconds. No race differs (Alliance seed 1 at 26, 38, 50, 60 within 0.4% of Horde's gain). Round 20 M3 wrote "downranks cannot flip: the attack power share is the same for every rank" and did not sweep; the same sameness is why they flip. The round 14 test's three Arcane Shot cases follow. | `$R/hunter/arc-all.tsv`, `sting-all.tsv`, `ba-bm-h.tsv`; `sim/request/rank_choices_hunter_r22_test.go` | CONFIRMED-BUG (rank) | R20 628f99a7c (round 20 M3 asserted, did not measure) | S aae25e43b |
| B2 | Aimed Shot rank 1 rows (28-59, bare 60) re-checked after round 21's A2: the ungated Aimed Shot makes rank 1 stronger, not weaker (BM 50 damage per mana: rank 1 3.33, Arcane Shot 2.78). | `dpm.py` runs | Aligned | - | - |
| B3 | Hunter's Mark ranks against Wowhead Forever tooltips: 26/59/98/143 (1130, 14323, 14324, 14325); the 3 November "rank 4 less than rank 3" fix is in. Aspect of the Hawk 20/35/50/70/90/110 (13165-14322) match the 3 November line and the engine. | nether.wowhead.com/forever/tooltip/spell/<id>, 5 Nov | Aligned | - | - |

## Marksmanship

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| M1 | **Aimed Shot at 50 and 51 cast rank 4 and ran the hunter dry for 90 s of 180.** Rounds 17 and 18 measured rank 1 and rank 4 even there while Aimed Shot waited for an auto shot; round 21's A2 ungated it, so it takes the shared cooldown far more often and rank 4 (210 mana) drains the bar. Rank 1 now wins for every race at both levels: 3000 it, eight races +4.8 to +5.2%; 10000 it, seeds 1 / 50001: Horde 50 291.8 -> 306.9 / 292.0 -> 307.0 (+5.2/+5.1%), 51 +5.1/+5.1%; Alliance 50 +4.9/+4.9%, 51 +4.9/+4.9%. Failing test first (`TestRound22MarksmanshipAimedShotRankOneAt50And51`, red on 20902), then the row runs 28-59 unbroken; the round 17/18 tests that pinned rank 4 at 50-51 follow. Guide: keep rank 1 to 60 out in the world, drag it back at 28, 36, 44, 52 and 60. | `$R/hunter/mm-aimed-5051-races.tsv`, `mm5051-after2.tsv` | CONFIRMED-BUG (rank) | R21 4abba08a0 (A2 made the round 18 tie stale) | S 0a6bbcacc |
| M2 | **Arcane Shot and Serpent Sting ranks while leveling** (as B1). Rows: Arcane Shot rank 1 at 20-29, rank 3 at 30-39, rank 1 at 40-59 and bare 60; Serpent Sting rank 1 at 20-29, newest from 30 (it wins at 30-60 by 1 to 3%). 10000 it, Horde, seeds 1 / 50001 against the landed rows: 20 +2.5/+2.5%, 26 +5.9/+5.9%, 29 +8.5/+8.6%, 30 +1.7/+1.7%, 33 +0.9/+0.9%, 36 +3.4/+3.4%, 39 +2.8/+2.8%, 40 +4.5/+4.4%, 43 +5.4/+5.3%, 46 +5.6/+5.6%, 49 +4.9/+4.7%, 50 +10.6/+10.4% (with M1), 53 +5.1/+5.0%, 56 +4.5/+4.4%, 59 +5.9/+5.8%, bare 60 +7.4/+7.3%. Alliance seed 1 at 20, 33, 50, 60: +2.5, +2.1, +9.7, +6.9%. Why 30-39 differs from Beast Mastery: the band 30 gear and Rapid Recuperation at 33-34 leave more mana, so the sting's newest rank and Arcane Shot rank 3 pay; the sweep decides, not a story. | `$R/hunter/arc-all.tsv`, `sting-all.tsv`, `ba-mm-h.tsv` | CONFIRMED-BUG (rank) | R20 628f99a7c | S aae25e43b |
| M3 | Sims requested by guide round 22 (damage per mana, 10000 it, Horde, bare, before the M1/M2 rows): 40 rank 1 Aimed 4.40, sting 4.48, Arcane 3.13; 50 rank 4 Aimed 2.45 (M1), sting 4.36, Arcane 3.52; 52 Aimed 5.07, sting 4.40, Arcane 3.37; 58 Aimed 5.33, sting 4.37, Arcane 3.46; bare 60 Aimed 7.68, sting 5.37, Arcane 4.09. Beast Mastery 40 2.73 / 3.07 / 2.22; 50 3.33 / 3.90 / 2.78; 58 3.59 / 3.96 / 2.76. Raid 60 Auto Shot share: 46% (round 21 M1 unchanged in shape). The notes' "does more for its mana than Arcane Shot" holds everywhere. | `dpm.py` | Answered | - | - |

## Survival (melee)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | Ranks of every spell the card casts with a choice, swept (2000 it, Horde): Immolation Trap at 20-60, Lacerate at 33-60, Mongoose Bite at 33-60. The newest rank wins at every level (Immolation Trap by 1.5 to 3.5% over the next; Survival does not run dry the way the shooters do). No row. | `$R/hunter/sv-sweep.tsv` | Aligned | - | - |
| S2 | Raptor Strike rank rows (round 21, 26-29 and 48-49): the Forever low-rank penalty is a spell power cut (`core.levelPenalty`; Blizzard 1 October "Spell Damage and Spell Healing"), physical strikes keep everything, so the rows were judged under the right rule. Rows in sync with `sim/rankchoices`. | `sim/core/spell.go:795` | Aligned | - | - |
| S3 | Immolation Trap's 2% attack power a tick is still single-source (Blizzard 16 Oct "Fire traps now correctly scale with Attack Power"; no number). The Forever tooltips (13795, 14302, 14305, 13797) show flat damage only. | nether.wowhead.com, 5 Nov | OPEN | - | - |

## Blizzard's Forever notes, every hunter line (24 Sep to 4 Nov, us.forums.blizzard.com/en/wow/t/2360696, fetched 5 Nov)

| # | note line | engine | verdict |
|---|---|---|---|
| N1 | 16 Oct, Pets: "pets were only ever inheriting the player's Melee Critical Strike chance instead of whichever is higher, melee or ranged" | `pet.go` gave the pet the melee crit stat only; a scope's ranged-only crit (8720, +2%) never reached it. Test first `TestHunterPetInheritsTheHigherRangedCrit` (red 0.00, want 2); core `Unit.RangedOnlyCritRating` sums `AddBonusRangedCritRating`; the pet's melee channel adds it. **CONFIRMED-BUG, `old`** (no hunter round checked this line), E d4a1e2ca0. No published number moves: no preset carries a scope and the site's enchant table has none. |
| N2 | 16 Oct: Auto Shot needs 0.5 s standing still; still fires through Aimed/Sniper Shot | round 21 A1, re-read: `shootDuringCast` fires without the wind-up, instants keep the 0.5 s push | Aligned |
| N3 | 16 Oct: fire traps scale with attack power; one fire and one frost trap | S3 | OPEN (single source) |
| N4 | 16 Oct: Trueshot Aura raid wide; Sniper Shot gains MM talents; Resourcefulness on melee and ranged ability crits only, not traps | round 15-21 reports | Aligned (re-read) |
| N5 | 3 Nov: Hunter's Mark ranks no longer stack; rank 4 above rank 3 | B3 | Aligned |
| N6 | 3 Nov: all traps share 30 s | client category, round 21 | Aligned |
| N7 | 3 Nov: Aimed Shot +110%, Arcane Shot +30% and 0.15, Serpent Sting 5% and +20%, Hawk 20-110, Barrage 5/10/15, Sniper Shot +100%, Trueshot 20-100, Lightning Reflexes 3-15%, Lacerating Strikes overwrite | rounds 13, 20; B3 | Aligned; B1/M2 are the rank consequence |
| N8 | 27 Oct: Orc Command +5% hunter pet damage | `racials.go:233` | Aligned |
| N9 | 8 Oct: Unleashed Fury reaches the pet; Expose Prey 10 s; Disengage threat | rounds 15-20 | Aligned |
| N10 | 16 Oct: pets benefit from Expertise | no hunter Expertise source on any BiS list | Aligned (nothing to inherit) |
| N11 | 24 Sep: Furious Howl -40%, Lone Wolf -4% fix; 16 Oct pet family ability changes (Dismember, Sonic Blast, Mine!) | the sim's pet is a cat (Bite, Claw) | Not reached |

## Round 21's own changes, re-read

| # | change | check | verdict |
|---|---|---|---|
| R1 | Auto Shot through casts (fork 723a3ccaf) | code read; the logs' ~60% of casts with a shot inside; raid unchanged against baseline | Aligned |
| R2 | Low-mana rule at every level (S 9045ac705) | B2/M3 damage per mana; no stall in any sweep | Aligned |
| R3 | Hunter's Mark pre-pull (fork fdc82f78b) | breakdown: one mark cast at -20 s, 0.9 hits (10% resist) | Aligned |
| R4 | Raptor Strike rank rows (S 7b3cec608) | S2 | Aligned |
| R5 | A2 ungated Aimed Shot | made the Marksmanship 50-51 tie stale | M1 (CONFIRMED, R21) |

## Before/after

Raid 60, guide builds, frozen BiS, 10000 it, seeds 1 / 50001, both factions, against this branch's int/r22 baseline (after the druid lane's crit change): every hunter spec identical to the digit (Beast Mastery 868.3/868.2 Alliance, 865.5/865.3 Horde; Marksmanship 959.3/960.6, 960.0/959.2; Survival 851.1/850.4, 850.7/851.2). The rows are bare-only at 60 and no preset carries a scope.

Bare bands, Horde, 10000 it, seed 1 / 50001, against `$R/sb-before` (landed d5cc4d6b4; bare bands carry no Leader of the Pack, so the druid change does not reach them):

| spec | level: seed 1 (abs) / seed 50001 |
|---|---|
| Marksmanship | 20 87.3 -> 89.5 +2.5/+2.5%; 30 145.6 -> 148.0 +1.7/+1.7%; 40 217.9 -> 227.6 +4.5/+4.4%; 50 291.8 -> 322.6 +10.6/+10.4%; bare 60 424.3 -> 455.7 +7.4/+7.3% |
| Beast Mastery | 20-25 unchanged; 26 99.7 -> 104.1 +4.4/+4.3%; 30s up to +9.9% (38); 40 209.7 -> 224.2 +6.9/+6.9%; 50 304.8 -> 317.5 +4.2/+4.2%; bare 60 432.9 -> 459.5 +6.1/+6.0% |
| Survival | unchanged |

Full tables: `$R/hunter/ba-mm-h.tsv`, `ba-bm-h.tsv`, `mm5051-after2.tsv`.

## Files to regenerate

- addon-data / Data.lua: the BM and MM rank lines (Arcane Shot, Serpent Sting, Aimed Shot 50-51) on every leveling band; coach starters for both.
- Ladder goldens (`sim/request TestRotationLadder`) for Beast Mastery and Marksmanship bare bands 20-60.
- dps-per-point and planner vectors for BM and MM leveling bands (more mana in the fight: Intellect and mp5 worth less).
- `sim/rankchoices/rank-choices.json` synced here with `make rank-choices-sync`. No APL changed. Fork: no golden.

## Measure at refresh

- BM and MM leveling BiS picks and weights on the new rows (the fights no longer run dry, so Intellect and mp5 drop).
- Round 21's list stands.

## For other lanes

- **Platform (shared core, flagged):** `sim/core/unit.go` gains `Unit.RangedOnlyCritRating` (summed by `AddBonusRangedCritRating`). Only the hunter pet reads it; no other spec moves.
- **Every mana class:** any ranked spell whose damage is mostly a share of attack power or a share that is the same on every rank behaves like Arcane Shot here: a low rank keeps most of the damage for far less mana. Worth a sweep wherever a rank row was removed by assertion rather than measurement.

## For the guide lane

- BM:159 says rank 1 Aimed Shot hits for "well over half" of the newest rank, MM:165 says "most": the two pages disagree (wow-player SHOULD 6, not applied here).

## Guide text

Done in this lane: beast-mastery.md and marksmanship.md "Levels 20 to 60" (the Arcane Shot and Serpent Sting rank paragraph on both; Marksmanship's Aimed Shot paragraph to 60). `_prose.test.ts` passes (18). wow-player on the rendered pages (astro dev on this branch, 1440, `$R/hunter/shots/*-leveling.png`): SHIP WITH FIXES, two MUSTs ("train the newest Serpent Sting" at 30/40 named no new rank; Marksmanship's Aimed Shot paragraph said 59 then 60) and five SHOULDs; both MUSTs and SHOULDs 3, 4, 5, 7 applied; 6 handed to the guide lane. Second pass on `*-leveling2.png`: **SHIP**, no MUST. One optional polish not applied (the drag-back clause could name Arcane Shot explicitly; the paragraph already ends the sting's low rank at 30 and 40).

## Logs

Last looked 07:58 UTC 5 Nov: api.foreversixty.gg `/v1/reports/recent` lists only `z2opt2x2gs5u` (20 Oct); foreverlogs.gg `/api/reports/public` answers 403 (platform lane, 07:26); Warcraft Logs v1 `/v1/zones` on www (66), classic (57), fresh (21), vanilla (29), sod (12) lists no Forever zone.

## Still open

- Round 20/21 log items: the hawk's swing, Heart of the Lion on the pet, Focused Fire snapshot, Ranged Weapon Specialization on the sting, Lacerating Strikes, Immolation Trap's 2% (S3), whether multipliers reach the attack power share (round 21 X1). Each needs a live Forever log: a hunter with a known sheet ranged attack power (Lone Wolf for X1), a BM hawk, a Survival trap.
- B1/M2 rest on the share reaching a low rank in full. Blizzard's notes give one coefficient for "all ranks" and their low-rank rule names spell damage only; a live log of a rank 1 Arcane Shot at 40+ (hit = 13ish base x 1.3 + 15% sheet RAP) would confirm it.
- Pending owner, not reopened: Arcane Shot with spell power, level-60 book ranks, Heart of the Lion 280/276.
