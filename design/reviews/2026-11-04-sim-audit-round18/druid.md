# Sim-accuracy round 18: Druid

Lane `druid`:
- Site: wt/r18-druid, from int/r18 141cffea8.
- Fork: lane/r18-druid, from b2c4a2d05.
- Client: 1.60.1.70755 with the late re-layer.
- Specs: Balance, Feral cat, Feral bear (judged on threat), Restoration.
- Method: measured on guide builds (`$R/rt`) and the frozen BiS (`$R/bis-frozen`).

**Summary: 4 confirmed, all fixed.**
- Three are Restoration rank-choice gaps, found by sweeping the ranks of every spell no round had judged level by level.
- One is Read Leyline adding onto Innervate. Blizzard's notes settle it: Innervate replaces Read Leyline.
- Raid 60 is unchanged for all four specs.
- The fork test `sim/druid/tank` was red on int/r18. The cause is a stale golden from round 17. It is regenerated.

**Logs:** there are still no launch logs.
- foreverlogs.gg's newest report is 2814 ("Test", 2026-11-04 12:24 UTC, 0 encounters). 2815 and up return nothing.
- api.foreversixty.gg/v1/reports/recent has one report, from 2026-10-20.
- Warcraft Logs is blocked (gcloud); I did not use it.

## Seed items

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **Eclipse on a missed Wrath.** The engine banks 2 charges on any finished Wrath (`druid/talents.go:712-725`). `balance.md` says a miss banks none. No new source settles it:<br>- The Forever tooltip only says "Your Wrath spell reduces..." (nether.wowhead.com/forever/tooltip/spell/408248, read 2026-11-04). Its "0.03 sec" is a scrape artifact: talents/druid.json reads 0.17/0.33/0.50.<br>- Season of Discovery sources do not cover misses (warcraft.wiki.gg/wiki/Eclipse_(druid_rune)).<br>- No Forever log exists. | as above; round 17 L1 | OPEN (a launch Balance log with a resisted Wrath followed by Starfire) | - | - |
| S2 | **Cat Form ±20% spread; the 54.8 unarmed fallback.** No log. Every paw ability is linear in the paw's average. Every published cat and bear band wears a weapon. | round 17 S1/S2 | OPEN (a launch cat log's white min/max at a known weapon) | - | - |
| S3 | **Read Leyline stacking with Innervate.** The engine added them: Innervate +4 plus Energized +0.5 on `SpiritRegenMultiplier`, 4.5 together.<br>Blizzard's Forever development notes (Kaivax, 26 October, repeated in the 3 November full update, https://us.forums.blizzard.com/en/wow/t/wow-classic-forever-beta-patch-notes/2360696) say: "This ability cannot be used while Innervate is active, and will be replaced by Innervate as a more powerful effect." The same notes say Attuned "is removed upon entering a raid". TweakTown (https://www.tweaktown.com/news/114319/) reports the same.<br>Fixed:<br>- an Innervate landing ends Energized (`core/buffs.go` InnervateAura);<br>- Read Leyline is not castable under any Innervate (`core/racial_read_leyline.go`).<br>Test `TestReadLeylineGivesWayToInnervate`: red, then green.<br>Effect: Alliance Skyborne Restoration, raid 60, -0.13% (949.2 to 948.0, seeds 1 and 50001). No published race is Skyborne.<br>**Shared core:** this reaches every Alliance Skyborne of any class. | Blizzard notes as quoted | CONFIRMED-BUG | old (round 16 S3 added Read Leyline; round 17 X2 left this open) | E 960e3fc55; prose S bd63ff021 |
| S4 | **Attuned in raids (guide text).** `balance.md` and `restoration.md` said Attuned "doubles it wherever you are". Blizzard's notes remove it when you enter a raid. The engine already models the base 50% only. | as S3 | Guide text fixed | - | S bd63ff021 |

## Restoration

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | **Healing Touch filler never judged by level.** Round 16 wave 2 swept Moonfire, Wrath, Insect Swarm and Rejuvenation only.<br>I swept the HT filler (line 9, the tank under 55%) every level from 26 to 53. The Nature's Swiftness line (4) stayed at the top rank. At 10,000 iterations, seeds 1 and 50001:<br>- rank 3 beats the top rank by 1.2 to 1.5% at 30 to 33 (tauren 1.1 to 1.5%);<br>- 34 to 39 are under 1%;<br>- rank 5 beats the top rank by 1.1 to 1.3% at 40 to 44 (tauren 1.1 to 1.3%);<br>- from 45 it is under 1%, and the top rank wins from 50.<br>Fix: line-scoped rows (`lines: [9]`). The Ranks and Levels 20 to 60 prose and the line note are updated. Test `TestRestorationHealingTouchFillerRanksByLevel`: red, then green.<br>Measured with the real rows at bare 30 / 33 / 40 / 44: +1.4/+1.5%, +1.2/+1.3%, +1.3/+1.3%, +1.2/+1.2%. | `w/htline.out`, `htv1.out`, `htv2.out`, `htv-horde.out` (lane dir) | CONFIRMED-BUG | old (no round swept HT ranks) | S ddd06e732, E 7986a6911 (note) |
| R2 | **Party Rejuvenation at 40 and 41.** The 10-59 rank-1 row loses to rank 2 at 40 and 41:<br>- night elf: +1.4/+1.4% at 40 and +2.1/+2.0% at 41;<br>- tauren: +0.8% at 40 and +1.4% at 41.<br>It is level again at 39 (-0.9%) and 42 (+0.8%). This is measured with R1's rows in. The cause is real: Regrowth rank 6 arrives at 42 and costs more.<br>Round 16 W2-2 saw "+1.3% at 40 only" and kept rank 1. The rule adopts a row that gains beyond error at both seeds.<br>Fix: a 40-41 row at rank 2 for every race. For tauren at 40, rank 2 is still the better rank, just under the bar. Prose updated ("Rank 1 stays on the party, apart from rank 2 at 40 and 41"). Test `TestRestorationPartyRejuvenationAtFortyAndFortyOne`: red, then green. | `w/party40.py` output, `party3.py` | CONFIRMED-BUG | old (round 16 W2-2 looked and kept it) | S 0e8819586 |
| R3 | Regrowth (20-60) and Tranquility (30-60) swept by level: the top rank wins or ties at every level. | `w/rg.out`, `w/tq.out` | Aligned | - | - |
| R4 | Dormant Heart of the Mountain (raid trinket) is modelled in `common/item_effects/client_procs.go:310` (Eternal Power on every Healing Touch rank). The comment at `common/item_effects.go:3756-3771` still says it is "NOT implemented". That is a stale comment, not a sim bug. | engine | Aligned (comment stale) | - | - |
| R5 | Bare 60 Healing Touch rank 11 (25297) beats rank 10 by 6.4%. Pending owner (level-60 book ranks). Not reopened. | `w/ht.out` | Pending owner | - | - |

## Balance

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| L1 | **Starfire swept by level** (never swept before), every 3 levels from 20 to 59: the top rank wins or ties everywhere. At 36 and 39, rank 2 is within 0.12%. | `w/sf.out` | Aligned | - | - |
| L2 | Insect Swarm, Moonfire and Wrath swept at 40-59, where round 16 judged mostly 20-39. The top rank holds; the largest gap is Wrath rank 6 at 46, +0.6%. | `w/is.out`, `mf.out`, `wr.out` | Aligned | - | - |
| L3 | The 70755 re-layer (43dc0f074), checked rank by rank against the engine:<br>- Starfire 113/154/201/277/372/471/533, with per-level growth;<br>- Rip per tick 4/5/8/10/14/19 plus the per-point step (31.9 at rank 6; tooltip 305..1071 over 12 s matches);<br>- Claw 140%, Shred 200%, Ravage 435%;<br>- Rake 27/39/56/76 up front and 20/26/32/42 a tick;<br>- Bite base and 3.1 an energy (tooltip 215-279 / 855-919 matches).<br>The flat bonus is added before the weapon percent, which is the emulator order (effect 58 index 0, then 31). | spellconst diff 43dc0f074^..43dc0f074; Wowhead Forever tooltips 2026-11-04 | Aligned | - | - |
| L4 | Bare bands 20-60 breakdowns: the cast mix and mana fit the list (Wrath when behind pace; Eclipse 89-91% from 40). "Improved Wrath" mana at 40/50 is Talons of Wrath (1302521), not the talent. | `out/balance-bare*.txt` | Aligned | - | - |
| L5 | Ironfeather 2-piece and the Heart trinkets (Frozen Heart's Use is Frost/Shadow only, so Balance gets its hit) are implemented as the client says. | `core/client_sets_gen.go`, `common/item_effects.go:3732` | Aligned | - | - |

## Feral cat

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| C1 | **Bite energy pooling** (guide rounds 10-14, "Sims requested 2", never run). Raid 60, 3000 iterations:<br>- shipped list: 886.4;<br>- no Shred at 5 points: 888.1 (+0.2%);<br>- pooling to 50 / 70 / 90 energy: 888.6 / 885.9 / 877.0.<br>Nothing clears 1%. "Do not pool energy for it" holds. | `w/pool.py` | Aligned (closes the carried request) | - | - |
| C2 | Bare 20-60 breakdowns are plausible:<br>- Claw before 24 and Shred after;<br>- free Faerie Fire in form;<br>- Devious Strike from the band-20 Defias set;<br>- Manual Crowd Pummeler on its 3 min cooldown;<br>- Thunderbrew's Boot Flask as Breath of Fire;<br>- Poised spent on the Ravage opener.<br>Rake stays out (round 3 F5; still a loss on 200% Shred). | `out/feral-bare*.txt` | Aligned | - | - |
| C3 | Rend and Tear is on every special's direct hit (Claw, Shred, Ravage, Rake initial, Bite, Maul, Swipe, Lacerate, Primal Bite). It is not on bleed ticks. The client row is a dummy aura (1223246). | `druid/feral_talents.go:84-118` | Aligned | - | - |

## Feral bear

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| B1 | **Fork test `sim/druid/tank` TestFeralTank was red on int/r18.** The stat-weight golden's Hit moved from 0 to 0.118. Round 17's 7c8f710d0 keeps measuring hit past the white cap while spells still miss, and the bear's Demoralizing Roar and Faerie Fire do. With the pre-7c8f710d0 `hitprofile.go` the old golden passes. Only that line changed. Golden regenerated; every `./sim/druid/...` package is green. | `go test -tags=with_db ./sim/druid/tank/` | Stale golden (not a sim bug) | R17 7c8f710d0 | E 403dab609 |
| B2 | Bare 20-60 tank fights: Lacerate ranks use Forever's ids (1328617/1328616/414644/1235826/1235827). Primal Bite runs on its 6 s cooldown, Demoralizing Roar is refreshed under 3 s, and Barkskin and Frenzied Regeneration come in from 44/36. Raid-60 values are unchanged from round 17 B1/B2. | `out/feral-bear-bare*.txt` | Aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | Round 17's shared changes reach druids only through B1. Heart of the Lion on pets: no druid pet takes buffs. The rank-choice `lines` scope is used by R1. | fork diff f637ec32c..b2c4a2d05 | - | - | - |

## Before/after

**Raid 60:** `$R/before.tsv` against this branch, using the guide builds and the frozen BiS, 10,000 iterations. All 16 rows are identical (+0.00%), e.g.:
- Balance alliance s1: 652.8;
- cat: 886.1;
- bear: 1200.3;
- Restoration: 939.5.

**Below 60** (Restoration bare, alliance, effective HPS, seed 1; seed 50001 within 0.1):

| Level | Before | After | Change |
|---|---|---|---|
| 30 | 119.8 | 121.5 | +1.4% |
| 33 | 119.2 | 120.6 | +1.2% |
| 40 | 206.3 | 211.9 | about +2.7%: HT row +1.3% and party row +1.4% |
| 41 | 203.7 | 210.5 | about +3.3% |
| 44 | 205.8 | 208.3 | +1.2% |

**Alliance Skyborne Restoration, raid 60:** 949.2 to 948.0 (-0.13%), from S3.

## Files to regenerate

- `data/builds/1.60.1.70755/addon-data.json` and `addon/ForeverSixty/Data.lua`: the addon rotation reads `rank-choices.json` (R1, R2) and the APL notes.
- The Restoration ladder golden (`sim/request/testdata/ladder/druid-restoration.golden.md`), if it prints sub-60 ranks or HPS.
- Fork golden regenerated here: `sim/druid/tank/TestFeralTank.results` (B1).

## Measure at refresh

- Restoration healer picks at bands 30 and 40 on the new Healing Touch and party Rejuvenation rows.

## For other lanes

- **All lanes / integrator:** E 960e3fc55 touches `sim/core/buffs.go` (InnervateAura) and `racial_read_leyline.go`. It moves only an Alliance Skyborne who is also Innervated: any class taking a raid Innervate, a druid's own. No published race is Skyborne.
- **Mage guide lane:** `arcane.md:98` and `fire.md:98` mention Read Leyline but not Attuned or Innervate. They are accurate as written.
- **Items:** the stale "NOT implemented" comment on Dormant Heart at `common/item_effects.go:3756`.

## For the guide lane

Done here, in `restoration.md` and `balance.md`:
- R1 and R2 rank sentences (Ranks paragraph, Levels 20 to 60);
- Read Leyline under Innervate (Mana and Races);
- Attuned in raids (both Races paragraphs).

wow-player on the rendered pages (1440, captures in `$R/druid/cap/`):
- First pass: HOLD, two MUSTs, both fixed in S 0bda57567:
  - "After Nature's Swiftness" read as "after you pop it"; it is now "With Nature's Swiftness".
  - "doubles it" had no clear subject, and it did not say that entering a raid removes Attuned. Both are now stated.
- SHOULDs applied: the Read Leyline wording in both guides, and "rank 4 from 38 to 49, then higher".
- Not applied: a table for the Ranks paragraph. The guides keep prose as their house style.
- Second pass: **SHIP**.
- Optional nit left open: Balance's "Innervate replaces it" could read "Innervate overrides it".

## Still open

- S1 Eclipse on a missed Wrath, S2 form spread and unarmed paw: need launch logs.
- Pending owner, not reopened:
  - low-rank penalty parameters;
  - Rejuvenation 1 tooltip;
  - Healing Touch / Rejuvenation 11 and Ferocious Bite 6 as level-60 book ranks.
