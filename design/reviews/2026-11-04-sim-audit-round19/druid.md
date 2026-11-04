# Sim-accuracy round 19: Druid

Lane `druid`:
- Site: wt/r19-druid, from int/r19 f91e9d3f9.
- Fork: lane/r19-druid, from int/r19 6ea675594.
- Client: 1.60.1.70755 with the late re-layer. 70790 is known to the integrator (the live client, 70755 with its hotfixes baked in); I did not touch it.
- Specs: Balance, Feral cat, Feral bear (judged on threat), Restoration.
- Method: rank rows judged level by level with `spec-breakdown -level` (floor, +3, +6, +9, and every level where it was close), 10,000 iterations, seeds 1 and 50001, night elf and tauren, frozen BiS (`$R/bis-frozen`). Each candidate rank was forced by authoring that rank's own id and rewriting every reference to the authored id, conditions included (round 18 priest X4). Where a forced id collided with another line's authored id (tank Rejuvenation rank 7 = 8910, the party line's id; rank 10 = 9841 under its own rows), that cell was discarded.

**Summary: 7 confirmed, all fixed. All are rank rows below 60; raid 60 is unchanged for all four specs.**

The common cause: earlier rounds held rank rows to a 1% bar, but the ledger (round 13/14 ruling, `ledger.md:119`) says rank rows are adopted beyond error and the 1% bar is for builds. Restoration's tank Rejuvenation also never had its 53-57 stretch re-judged after the healer rows around it moved.

**Logs:** none from launch yet. Last looked 18:52 UTC: foreverlogs.gg's API answers "API access not allowed" from here; api.foreversixty.gg/v1/reports/recent has one report, from 2026-10-20. Warcraft Logs is blocked (gcloud); not used.

## Seed items

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | Restoration 40-59 rank rows re-judged level by level: see R1-R3, R5. | `w/tank-a.out`, `tank-h.out`, `party-a.out`, `party-h.out`, `ht-a.out`, `ht-a2.out`, `ht-h.out`, `rg-a.out` (lane dir) | see below | - | - |
| S2 | **Dormant Heart of the Mountain** (249473). It is implemented: `common/item_effects/client_procs.go:303-330` registers Eternal Power (1249119, +40 healing and +13 damage for 15 s) on every rank of Greater Heal, Healing Touch, Healing Wave and Holy Light. The client rows (raw 70755 SpellEffect 1252398: aura 42, trigger 1249119; 1252399/1252400: aura 135 = 40, aura 13 = 13, flat, school mask 126) match. `priest/healing/dormant_heart_test.go` covers it. The comment at `common/item_effects.go:3756` that said "NOT implemented" was stale; it is replaced with a pointer. Comment only. | as stated | Aligned (stale comment removed) | - | E c3abe2881 |
| S3 | Guide-accuracy round 15 druid "Sims requested": none new. The carried Bite pooling was closed by round 18 C1. | round15/druid.md:150 | Closed | - | - |
| S4 | **Cross-class lead (Ice Lance, integrator):** every spell-power-scaled druid spell the rotations cast has a non-zero client EffectBonusCoefficient, and the engine matches it: Wrath 0.429/0.486/0.571, Starfire 1, Moonfire 0.15 direct and 0.13 a tick, Insect Swarm 0.158 a tick, Rejuvenation 0.2 a tick, Regrowth 0.286 and 0.071 a tick, Healing Touch 0.429/0.571/0.714/0.857 then 1, Tranquility 0.067, Hurricane 0.03. The only zero rows are Swiftmend (it heals for the consumed HoT), Thorns and Entangling Roots' root effect, none cast for damage. `TestHealingTablesMatchTheClient` passes. | raw 70755 SpellEffect; `druid/moonfire.go:19-20`, `insect_swarm.go:26`, `wrath.go:34`, `starfire.go:28`, `healing_tables.go` | Aligned | - | - |
| S5 | **Shared watch (P1, low ranks proc less):** the rows that cast a share-0 rank are Restoration's party Rejuvenation rank 1 at bare 60 (pending owner) and below 60 where rank 1 is far under the level. Omen of Clarity rolls on heals (4%), so those casts may proc less in Forever. Unmodelled; no code change. The new rows (R1-R3, L1-L2) cast ranks that keep most or all of their bonus. | round 18 priest P1 | Flag only | - | - |

## Restoration

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | **Tank Rejuvenation at 44-59.** The rows were rank 4 at 38-49, rank 6 at 50-52, the top rank at 53-57, rank 8 at 58-59 (round 16). Level by level, both factions, both seeds:<br>- rank 4 still wins at 40-43;<br>- rank 5 beats rank 4 by 0.3 to 1.0% at 44-48;<br>- rank 6 beats rank 4 by 1.2 to 1.4% at 49, and the top rank by 2.1 to 2.7% at 53, 1.5 to 2.1% at 55, 1.0 to 1.6% at 56;<br>- rank 8 beats rank 9 by 0.8 to 1.0% at 57.<br>**What was wrong:** while leveling, a druid's mana runs out first, and the newest Rejuvenation costs much more than an old rank that keeps nearly all its healing bonus. Round 16 called 53-57 a tie at its 1% bar.<br>**Fix:** rows rank 4 at 38-43, rank 5 at 44-48, rank 6 at 49-56, rank 8 at 57-59. Tests `TestRestorationTankRejuvenationRanksByLevel` (red, then green) and `TestRestorationDruidRejuvenationKeepsItsHealingBonusInARaid` (updated). | `w/tank-a.out`, `w/tank-h.out` | CONFIRMED-BUG | old (round 16 W2 rows; 53-57 held by the 1% bar) | S 493e88c6e |
| R2 | **Party Rejuvenation at 42.** Rank 2 beats rank 1 by 0.8% (night elf) and 0.2 to 0.3% (tauren), beyond error at both seeds. Round 18 R2 left 42 at rank 1 as "under the bar". Rank 1 wins again at 43 and 39.<br>**Fix:** the rank-2 row runs 40-42. Test `TestRestorationPartyRejuvenationAtFortyToFortyTwo` (red, then green). | `w/party-a.out`, `w/party-h.out` | CONFIRMED-BUG | old (round 18 R2 looked and kept 42 at the 1% bar) | S 493e88c6e |
| R3 | **Healing Touch filler (line 9) at 34-39 and 45-49.** Round 18 R1 adopted rank 3 at 30-33 and rank 5 at 40-44 but left the levels where a lower rank won by under 1%:<br>- rank 4 beats the top rank by 0.4 to 0.9% at 34-39 (rank 3 ties it at 34);<br>- rank 5 by 1.0 to 1.1% at 45;<br>- rank 6 by 0.7 to 0.9% at 46-49 (rank 5 ties at 46);<br>- the top rank wins from 50 to 59.<br>The Nature's Swiftness line keeps the top rank (the cast is free).<br>**Fix:** rows rank 4 at 34-39, rank 5 at 40-45, rank 6 at 46-49. Test `TestRestorationHealingTouchFillerRanksByLevel` (red, then green). | `w/ht-a.out`, `ht-a2.out`, `ht-h.out` | CONFIRMED-BUG | old (round 18 R1 looked and kept these levels at the 1% bar) | S aafe8f0ad |
| R4 | Regrowth re-swept with the new rows at 40/44/46/49/50/53/56/59: the top rank wins at every level by 0.2 to 1.5%. | `w/rg-a.out` | Aligned | - | - |
| R5 | Raid 60 and bare 60 are unchanged: the rows end at 59. | `after60.tsv` | - | - | - |
| R6 | **Guide and notes.** The tank Rejuvenation and Healing Touch line notes, the rotation notes, the request copy and the fork copy name the new rows (the fork's tank Rejuvenation note had drifted from the site; synced). `restoration.md` Ranks becomes three paragraphs (rule at 60, ranks while leveling, Healing Touch), and Levels 20 to 60 names each tank rank switch. | wow-player on the rendered page: HOLD (one MUST: the Nature's Swiftness clause read as if it applied only outside 30-49), fixed, then **SHIP** (`$R/druid/cap2/`) | Guide text done | - | S aafe8f0ad, 161b24f6b; E f40898e69 |

## Balance

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| L1 | **Wrath rank 6 at 46-48.** Round 18 L2 swept every three levels at 1,500 iterations and saw rank 6 +0.6% at 46, called within error. At 10,000 iterations, both seeds, both factions: rank 6 beats rank 7 by 0.6 to 0.7% at 46, 0.3 to 0.4% at 47 and 0.1 to 0.2% at 48; they tie at 49. Rank 7 (trained at 46) costs more than its extra damage repays while the leveling fight runs short of mana.<br>**Fix:** a row, rank 6 at 46-48, every race. Test `TestBalanceRankRowsFortyToFiftyNine` (red, then green). | `w/bal-a.out`, `bal-h.out` | CONFIRMED-BUG | old (round 18 L2 looked at 1,500 iterations) | S 8523920a0 |
| L2 | **Insect Swarm rank 3 at 54-58.** Rank 3 beats rank 4 by 0.2 to 0.6% at 54-58 (night elf) and 0.1 to 0.2% at 54-56 (tauren, ties at 57-58). They tie at 52-53 and 59. Round 18 L2 saw "+0.35% at 56".<br>**Fix:** a row, rank 3 at 54-58, every race (no race loses at any level; a race key needs keyed rows that differ beyond error). Same test. | as L1 | CONFIRMED-BUG | old (as L1) | S 8523920a0 |
| L3 | Moonfire and Starfire at 40/43/46/49/50/53/56/59: the top rank wins everywhere, by 0.6% or more. | `w/bal-a.out` | Aligned | - | - |
| L4 | `balance.md` Leveling names the two rows in their own paragraph; "Insect Swarm stays at rank 1 until 39". | wow-player SHIP (as R6) | Guide text done | - | S 8523920a0, 161b24f6b |
| L5 | **Insect Swarm hit.** The raid breakdown shows 12.9 Insect Swarm casts a fight and 0% miss, which looked like a spell that never misses. A probe (Insect Swarm against Starfire, 4,000 applications on the same target) lands both at 95.7/95.8%: it rolls the normal spell hit. The 0% is the report's tick-weighted column (`OutcomeMagicHitNoHitCounter`). | lane probe, not committed | Aligned | - | - |
| L6 | Builds: the Balance Subtlety and Nature's Focus points are on the keep list (round 15 L2, round 17 L5). Not reopened. | - | Aligned (keeps) | - | - |

## Feral cat

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| C1 | **Shifting Power** (1322605): 70755's own tables now say 8 s (SpellCooldowns 101953 RecoveryTime 8000), 55% of base mana (SpellPower 314986), 40 energy, 1 s GCD (StartRecoveryTime 1000), and the engine matches (`druid/shifting_power.go`). Its comment still says "the client tables still say 16 s"; stale, harmless. Raid 60: 22.3 casts a fight, fed by Judgement of Wisdom. | raw 70755 rows | Aligned | - | - |
| C2 | Raid 60 breakdown: Shred 66.6 casts, Rip 13.5 (refresh under 2 s), Bite 3.2, Ravage opener, Omen of Clarity 8.5 procs a fight (2 PPM with its 10 s category cooldown, round 13), Thistle Tea through round 18's 1aa12229b (unchanged on a 100-energy bar). | `w/raid-feral.txt` | Aligned | - | - |
| C3 | The 70755 re-layer (Rip, Claw, Rake, Shred, Ravage, Starfire, Ferocious Bite in 43dc0f074) was checked rank by rank in round 18 L3; the spellconst diff names no other druid spell. | spellconst diff 43dc0f074^..43dc0f074 | Aligned | - | - |
| C4 | Cat and bear damage abilities cost energy or rage, not mana; the top rank wins at every level, so they need no rank rows. | - | Aligned | - | - |

## Feral bear

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| B1 | Raid 60 breakdown: Maul 73 a fight, Primal Bite on its 6 s cooldown (28.8), Lacerate 30.1, Demoralizing Roar 7.2, Swipe as the rage dump above 40, Barkskin on cooldown, Frenzied Regeneration under half health. Threat per second identical to the round-18 baseline. | `w/raid-feral-bear.txt`, `after60.tsv` | Aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | **Method note for every lane:** a rank-swap harness that authors another rank's id can collide with an id another line authors (Restoration's tank Rejuvenation rank 7 is 8910, the party line's authored id) or with the line's own id under its rows. Those cells read the other line's rows, not the rank. Discard them or run without the rows. | `w/tank-a.out` (r7 at 40-45, r10 at 58-59 equal the row ranks) | Note | - | - |

## Before/after

**Raid 60** (`after60.tsv` against `$R/before.tsv`, guide builds and frozen BiS, 10,000 iterations): all 16 rows identical, e.g. Balance alliance s1 652.8, cat 886.1, bear 1200.3, Restoration 939.5.

**Below 60**, bare, frozen BiS, 10,000 iterations (`$R/sb-before` from `$R/site-before/sim` against this branch; `w/ba.out`). Seed 1 absolute, alliance (night elf) / horde (tauren); seed 50001 within 0.15 points:

| Spec | Level | Alliance | Horde |
|---|---|---|---|
| Restoration (effective HPS) | 36 | 120.1 → 120.9 (+0.7%) | 119.2 → 119.9 (+0.6%) |
| | 42 | 207.5 → 209.2 (+0.8%) | 209.9 → 210.5 (+0.3%) |
| | 45 | 206.4 → 210.5 (+2.0%) | 209.0 → 212.7 (+1.8%) |
| | 47 | 207.6 → 211.2 (+1.7%) | 210.2 → 213.6 (+1.6%) |
| | 49 | 210.2 → 214.8 (+2.2%) | 213.0 → 217.0 (+1.9%) |
| | 53 | 325.3 → 332.2 (+2.1%) | 317.7 → 326.0 (+2.6%) |
| | 55 | 333.3 → 338.4 (+1.5%) | 325.6 → 332.0 (+2.0%) |
| | 57 | 339.5 → 342.1 (+0.8%) | 332.2 → 335.4 (+1.0%) |
| Balance (DPS) | 46 | 179.8 → 180.8 (+0.6%) | 177.2 → 178.4 (+0.7%) |
| | 47 | 180.6 → 181.4 (+0.4%) | 178.2 → 178.8 (+0.3%) |
| | 54 | 228.0 → 229.3 (+0.6%) | 225.5 → 225.8 (+0.1%) |
| | 56 | 231.1 → 231.9 (+0.3%) | 227.9 → 228.4 (+0.2%) |

No commit moves any spec outside the druid class (S2 is a comment).

## Files to regenerate

- `data/builds/1.60.1.70755/addon-data.json`, then `addon/ForeverSixty/Data.lua`: the addon rotation reads `rank-choices.json` and the line notes (R1-R3, L1-L2).
- The Restoration and Balance ladder goldens (`sim/request/testdata/ladder/druid-restoration.golden.md`, `druid-balance.golden.md`) if they print sub-60 ranks or values.
- Copied by hand, not regenerated: `sim/rankchoices/rank-choices.json` (= curated), `sim/request/apl/druid-restoration.apl.json`, and the fork's `ui/restoration_druid/apls/forever_restoration.apl.json` (E f40898e69).
- No fork goldens touched.

## Measure at refresh

- Restoration bands 40 and 50 and Balance bands 40 and 50 on the new rows. The rows were judged on frozen band gear; a re-ranked set can move a switch level by one.

## For other lanes

- **Every healer and caster lane:** earlier rounds held rank rows to a 1% bar the ledger does not set for rank rows (`ledger.md:119`). Rows judged "under the bar" are worth re-reading.
- **Every lane with a rank-swap harness:** X1 (id collisions).
- **Integrator:** S4 answer for druids: no zero-coefficient spell is cast; nothing to settle from logs.

## For the guide lane

Done here: `restoration.md` (Ranks, Ranks while leveling, the Healing Touch paragraph, Levels 20 to 60) and `balance.md` (Levels 20 to 60). wow-player: HOLD then SHIP on the rendered pages (`$R/druid/cap2/`). None left.

## Still open

- Round 18 S1 (Eclipse on a missed Wrath) and S2 (Cat Form spread, unarmed paw): need launch logs; none at 18:52 UTC.
- P1: the low-rank proc curve (S5).
- Pending owner, not reopened: the low-rank penalty parameters, Rejuvenation 1 at bare 60, Healing Touch / Rejuvenation 11 and Ferocious Bite 6 as level-60 book ranks.
- Stale comment, not a bug: `druid/shifting_power.go` still says the client tables give 16 s (70755 gives 8 s).
