# Sim audit round 13: casters, 2026-11-03

Specs: mage-arcane, mage-fire, mage-frost, shaman-elemental and druid-balance, every band.

## Method

- **Branches.** Engine `lane/r13-casters` from fork int/r13 (076a5a517). Site `wt/r13-casters` from int/r13 (93c5d2510), with `make simdb` and the druid apl-sync (not committed).
- **Before.** `r13/sb-before` run from `r13/site-before/sim`, plus `r13/before.tsv`. **After** is the lane engine on the lane site. Each after run passes `-rotation data/curated/apl/<spec>.json`, because sim/request/apl is the integrator's apl-sync copy.
- **Client.** 1.60.1.70753, the layered copy in `r11/c70753/builds/1.60.1.70753/raw`. The 70753 `hotfix-push.json` has 0 rows in every table.
- **Logs.** foreverlogs.gg beta timelines, fetched 2026-11-03 with `r3/casters/py/fl.sh`:
  - 40 Arcane Blast mage encounters and 156 Wrath druid encounters (`r13/casters/tl/`);
  - 60 shaman encounters with a Searing Totem (`r13/casters/tls/`).

  The scripts are `r13/casters/{mb,cc,ecl,searing3}.py`. The timeline carries `miss`, `cast`, `buff_applied` and `damage` events, so a missed cast and a proc's exact moment can both be read.
- **Re-audit.** Two read-only research passes (mage; Elemental and Balance) read the engine against the client rows. Every candidate below was then checked by hand, by log, or by a measured run.
- **Tests.** Every fix has a failing test first.
  - Fork, run with `go test -tags=with_db`: `sim/druid/...`, `sim/mage`, `sim/shaman/...`, `sim` (Sample|Dummy).
  - Site: `go test ./request/` (EveryWrittenRotation, the new Balance test), `./leveling/`, `./cmd/leveling-bis/`, and `web/src/content/guides/_prose.test.ts`.
  - `web/src/pages/guides/_spec.test.ts` could not load: it needs `web/src/data/generated/combos.json`, which is a sync.
- **Runs.** spec-breakdown, 10,000 iterations, seeds 1 and 50001, both factions.

## 1. Owner ruling 1: Tidal Focus for Elemental at bare 50 and 60

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| T1 | **Adopted.** The guide build `551033130000135/03/503352` becomes `.../053352`. Restoration's five tier-1 points move from Improved Healing Wave (a Healing Wave cast-time cut) to Tidal Focus (1% spell hit a point). Mindfulness 3, Natural Grace 3, Totemic Focus 5 and Improved Reincarnation 2 stay. Bands 20 to 40 are unchanged. Band 50 becomes `-0530000000000000` and band 60 `-0533520000000000` (the leveling walk, checked with `LadderTalentString`). | talents/shaman.json 104739 and 104729 rank texts. The APL casts 10438, 10497, 10587, 10605, 1326047, 15207, 15208, 29228 and 408345, never Healing Wave. | ruling | - | S 11f258619 |
| T2 | Guide prose follows: the leveling paragraph names Tidal Focus and says why, and the hit line no longer says only the raid build takes Tidal Focus. The prose test passes. | `web/src/content/guides/shaman/elemental.md` | - | - | S 11f258619 |

Bare 50: +5.80 / +5.76% alliance and +5.68 / +5.69% horde (seeds 1 / 50001). Bare 60: +5.68 / +5.88% alliance and +5.71 / +5.95% horde. Horde seed 1 absolute: 50 goes 155.65 to 164.49, 60 goes 191.94 to 202.89.

## 2. Round 12 opens and handoffs

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| O1 | Chain Lightning is a 2.0 s base cast and 1.5 s with Elemental Alacrity 3/3. The engine registers 2000 ms and applies Alacrity. The guide no longer states "2.0 s" for the talented spell. | `shaman/chain_lightning.go:56`; guide lines 54, 58, 81 | Aligned | - | - |
| O2 | Frost's leveling build stays as kept in round 12. Its BiS talents wait for the refresh (MR18). | - | Aligned (carried) | - | - |
| O3 | Regcast's handoff: **Frostfire Bolt flew at 28 yd/s.** SpellMisc.Speed is 24 on 401502, 1237312 and 1237313, the same as Fireball 10149 and 25306. Frostbolt 116 is 28. The test failed at 28, then passed. It is latent: no published rotation casts Frostfire Bolt. Found by the regcast lane. | SpellMisc rows | CONFIRMED-BUG (no published number) | old (2026-10-07) | E 99b329919 |
| O4 | Guides handoffs: two APL notes. Arcane Missiles 8 (25345) and Fireball 12 (25306) both have SkillLineAbility rows (13246 and 13245, AcquireMethod 0). Era's Ahn'Qiraj books have the same rows, and `data/curated/book-ranks.json` and spellranks mark both as books. The Fire APL casts Fireball 10151 (rank 11), whose SpellPower ManaCost is 395, so "Fireball's 395" is that rank. Both notes are right, and the sim casts the trainer ranks. | client rows; book-ranks.json | Aligned | - | - |
| O5 | Core lane's report: bare Elemental still drops about 0.3% from 0 to +0.01% haste with the regen phase randomised (`e88a33c97`). A paired probe on that commit (cast speed set by env, not committed) reads: −1% 191.10, −0.1% 191.94, 0% 192.08, +0.01% 191.37, +0.1% 191.44, +1% 192.28. Lava Burst casts fall from 12.66 to 12.05. See E1. | probe runs | (explained, E1) | - | - |

## 3. Elemental

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| E1 | **Any haste cost Lava Bursts.** With Elemental Alacrity, Lightning Bolt and Lava Burst are both 2.0 s casts. Five bolts therefore end exactly as Lava Burst's 10 s cooldown returns. A sliver of haste ended the fifth bolt 1 ms early; the rotation started a sixth bolt, and Lava Burst slipped 2 s every cycle (about 5% of its casts). A player waits that fraction of a second, so the rotation now does too: a `wait` line before the Lightning Bolt lines, taken when Lava Burst is ready within 0.1 s and Flame Shock is up. Published numbers (no haste) are identical to the decimal at raid and bare 60, both seeds and factions. +1% haste now reads 192.71 (was 192.28), so the haste weight rises. | `TestASliverOfHasteDoesNotCostLavaBursts` failed (13.40 against 13.84 casts), then passed. The fork APL copy is synced by hand. | CONFIRMED-BUG (haste weight only) | old (Lava Burst line, 2026-10-31 caster tier lane) | E 25fca77c0, S 457a602df |
| E2 | **Searing Totem attacked every 2.5 s; the logs say 2.437 s.** The client's Attack is a 2.2 s cast on all six ranks (SpellMisc CastingTimeIndex 24, 3606 to 10436). 1,072 landed-to-landed gaps over ranks 1 to 3 have medians of 2,433 to 2,438 ms. The cadence does not change with the bolt's flight (0 to 1.4 s flights all read 2.43 to 2.44 s), so upstream's "waits for its missile" stand-in is wrong. Raid 60 +0.11 to +0.20%, bare 60 +0.11 to +0.18%. | `TestSearingTotemAttacksAtTheLoggedCadence` failed (2.5 s), then passed. `shaman/fire_totems.go`. Two sources: the client row and the logs. | CONFIRMED-BUG | old (upstream; round 11 E5 left it open for a second source) | E 5ca93cfa3 |
| E3 | Fresh, aligned with the client rows: <ul><li>Lightning Bolt ranks 1 to 10 (base, per level, coefficient 0.714, cost, 2500 ms, speed 20); the rank-10 overload 408477 is half damage.</li><li>Chain Lightning ranks 1 to 4; Lava Burst ranks 1 to 4 (+20% on Flame Shock).</li><li>Flame Shock 4 to 6 (0.214 direct, 0.1 a tick, 12 + 9 s).</li><li>Magma Totem 4; Fire Nova 408345 / 408428; Mana Spring 10497.</li><li>Clearcasting 16246's cost mask.</li><li>Convection, Concussion, Call of Flame, Elemental Alacrity, Lightning Overload (3/7/10%), Elemental Fury, Elemental Devastation (melee only).</li><li>Thundering Strikes, Ancestral Knowledge, Tidal Focus, Mindfulness (17/33/50%), Totemic Focus.</li><li>Blood Fury; Artisan's Precision; Ironfeather 2-piece; The Elements 3-piece.</li></ul> | engine against SpellEffect, SpellMisc, SpellPower | Aligned | - | - |
| E4 | Elemental Mastery 16166: one ProcCharge in SpellAuraOptions, against the engine's 15 s buff (round 12 E2). Beta logs stop near level 30, and Mastery needs 30 Elemental points. | - | OPEN (a level-40+ log or a tooltip test) | old | - |
| E5 | Lightning Overload rolls on a missed bolt. That is consistent with M1 below: the same 0x10000 proc family procs on misses in the logs. | - | Aligned (by M1) | - | - |
| E6 | Build questions measured, not adopted: <ul><li>Natural Grace 3 (threat) to Restorative Totems 3 at bare 60: +0.64 to +0.67%, under the 1% bar. Held.</li><li>Ancestral Knowledge 3 in place of Thundering Strikes 3: −1.5% at bare 50 and 60. The guide's Thundering Strikes is right.</li></ul> | runs | Held / Aligned | - | - |

## 4. Balance

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| B1 | **Nature's Grace from a Wrath crit started when the cast ended. On this server it starts when the bolt lands.** 16880's ProcTypeMask is 0x14000 (on hit). In the beta timelines, all 31 Wrath crits with a separable flight (182 to 1398 ms, median 667) applied Nature's Grace on the damage event, and none on the cast. The buff then lasts its 3 s (removed 3,008 ms later). Because cast time is fixed when a cast starts, a crit's haste now reaches the cast after next. Raid Balance −0.28 to −0.33%. | `TestNaturesGraceFromAWrathCritStartsWhenTheBoltLands` failed, then passed. `druid/wrath.go` and `talents.go:157`. Two sources: the proc flag and the logs. | CONFIRMED-BUG | old (upstream "NG procs when the cast finishes") | E a86de2495 |
| B2 | **Balance had no pre-pull cast.** Every other caster rotation pre-casts (Frostbolt, Pyroblast, Lightning Bolt). Now a Wrath at −1.5 s (1.5 s with Improved Wrath 5/5) ends on the pull and lands after its flight, after the tank's first hit. A −3 s Wrath, landing on the pull, measures the same. Raid 60 +0.69 to +0.84%; bare 60 +0.34 to +0.61%. | `TestBalanceOpensWithAPrePullWrath` failed (0 lines), then passed. The fork APL copy is synced by hand. | CONFIRMED-BUG (rotation gap) | old (round 12 regcast D1 saw no pre-pull and did not ask) | S e25400173, E ed156274a |
| B3 | **Leveling build: two dead talents replaced** under the Shadow-build precedent (owner rulings 1 and 2: a point that does nothing in the rotation may be replaced). <ul><li>Improved Entangling Roots 1 moves to Nature's Grace 1 (no Roots is cast).</li><li>Subtlety 3 (threat only) moves to Naturalist 4 and 5 plus Subtlety 1, which keeps tier 3 open.</li></ul> The build `5232221115100151/0/553303` becomes `5232220115101151/0/555301`. The tree holds and nothing the rotation casts is dropped. Before the rotation changes: bare 30 +1.1%, 40 +1.5%, 50 +1.6%, 60 +5.2%, at both seeds and factions. The guide's build line, talent prose and opener follow. **For the wow-player review:** Roots kiting while leveling is a player habit the sim cannot see. | runs; talents/druid.json texts | Adopted (owner may overrule) | - | S b7bc3b144 |
| B4 | Eclipse 408248 banks its charges at the end of Wrath's cast, hit or miss. The client proc mask is 0x10000. The two log analogues disagree. Nature's Grace (B1) procs on impact. Missile Barrage, a rune from the same ability family, procs at the cast on misses and on Arcane Wrath (M1, M3). No beta log shows Eclipse (it needs 20 Balance points). | logs; `druid/talents.go:739` | OPEN (one level-29+ Balance log settles it) | - | - |
| B5 | Fresh, aligned with the client rows: <ul><li>Wrath 8; Starfire 6 and 7 (7 is a book); Moonfire 10 (4 + 1 ticks); Insect Swarm 5.</li><li>Moonfury (school multiplier), Vengeance, Moonglow, Improved Wrath, Genesis, Nature's Majesty, Nature's Reach, Nature's Splendor, Naturalist (2% a point).</li><li>Nature's Grace (10% cast speed, −150 ms global, 3 s); Eclipse amounts.</li><li>Omen of Clarity (Moonkin doubling, Wrath excluded); Innervate; Elune's Light; the Ironfeather and Magister's sets.</li></ul> | engine against client rows | Aligned | - | - |
| B6 | Carried: Improved Moonfire on base damage only (round 11 B1); Improved Wrath and Moonglow cost cuts add (round 9). | - | OPEN (unchanged) | old | - |

## 5. Mage (Arcane, Fire, Frost)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| M1 | Missile Barrage procs on a miss (round 6 A2, open since). Logs: 10 missed Arcane Blasts (`miss` events) gave 4 Missile Barrages, against 56 of 360 landed. The engine's OnCastComplete roll is right. | `r13/casters/mb.py`; e.g. report 2721 encounter 7775 at 35.721 s | Aligned (log) | - | - |
| M2 | Clearcasting survives a Missile Barrage Arcane Missiles (the engine's choice, which had no source). Logs: in all 9 channels started with both buffs up, Clearcasting stayed and was spent by the next spell, e.g. a Fire Blast at 39.889 s in 2721/7775. | `r13/casters/cc.py` | Aligned (log) | - | - |
| M3 | Missile Barrage from a travelling bolt rolls at the cast: 3 of 3 Arcane Wrath procs in the logs land on the cast event, as the engine does. A small sample. | `r13/casters` log scan | Aligned (small sample) | - | - |
| M4 | Fresh, aligned with the client rows: <ul><li>Arcane: Arcane Blast 1239700 and stack 400573; Missile Barrage 400588 / 400589; Arcane Missiles 10212 / 10274; Arcane Concentration; Arcane Power; Presence of Mind.</li><li>Arcane talents: Arcane Impact, Arcane Mind, Arcane Instability, Arcane Meditation.</li><li>Fire talents: Critical Mass, Incineration, Accelerant, Improved Fire Blast.</li><li>Fire spells: Ignite 412538; Improved Scorch 22959; Combustion; Heating Up 400625; Pyroblast 18809.</li><li>Frost talents: Improved Frostbolt, Precision, Ice Shards, Frost Channeling, Fingers of Frost, Shatter, Winter's Chill 12579.</li><li>Frost spells: Ice Lance's multiplier and 8 s category; Ice Barrier with Frigid Power.</li><li>Mana and gear: mana gems (shared 2 min category 1153); Mage Armor; Blood Fury and Eureka; Magister's and Devout set bonuses.</li><li>Openers, as in round 12 M2.</li></ul> | engine against client rows | Aligned | - | - |
| M5 | Percentage cost mods add rather than multiply: Arcane Blast stacks and Arcane Power, and Arcane Power on a free Missile Barrage. One in-game readout of a stacked Blast's cost under Arcane Power settles it. | `arcane_blast.go`; `talents.go` | OPEN (low confidence) | old | - |
| M6 | Master of Elements refunds nothing on a crit that Clearcasting made free. The text says "base mana cost". Fire raid only, under 0.3%. | `talents.go` applyMasterOfElements | OPEN (carried from round 8) | old | - |
| M7 | Carried: Ice Lance's coefficient (the client column is 0; two unexplained dummy effects, base 20 and base 5); Fire Power on DoTs applied after flight; Ice Lance's category-less cooldown (round 11 M4). | - | OPEN | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | A proc's moment depends on the proc, not on one rule. Nature's Grace (an original talent) procs on impact. Missile Barrage (a rune) procs at the cast, misses included. Each travelling-bolt proc needs its own evidence. | B1, M1, M3, B4 | - | - | - |
| X2 | Skyborne races (Wind Blessed, a permanent 2% haste) are never measured for Balance or Elemental. spec-breakdown has no race flag, so this is a ranker question. | combos.json; `core/racials.go:192` | OPEN (ranks lane, Measure at refresh) | - | - |
| X3 | The 70753 hotfix push is still empty (0 rows in every table). | `hotfix-push.json` | - | - | - |

## Before/after

10,000 iterations. Alliance and horde are each shown as seed 1 / seed 50001.

| Spec, band | horde s1 before → after | alliance | horde | cause |
|---|---|---|---|---|
| druid-balance raid 60 | 572.2 → 575.07 | +0.51 / +0.48% | +0.50 / +0.38% | B2 +0.8, B1 −0.3 |
| druid-balance bare 30 | 72.75 → 74.00 | +1.87 / +1.73% | +1.72 / +1.66% | B3, B2 |
| druid-balance bare 40 | 150.98 → 154.30 | +2.33 / +2.25% | +2.20 / +1.99% | B3, B2 |
| druid-balance bare 50 | 191.53 → 195.64 | +2.34 / +2.34% | +2.15 / +1.98% | B3, B2 |
| druid-balance bare 60 | 280.36 → 296.54 | +5.82 / +5.91% | +5.77 / +5.60% | B3, B2 |
| druid-balance bare 20 | 25.01 → 25.11 (h s1) | - | +0.4% | B2 |
| shaman-elemental raid 60 | 479.6 → 480.23 | +0.11 / +0.20% | +0.13 / +0.13% | E2 |
| shaman-elemental bare 50 | 155.65 → 164.68 | +5.96 / +5.78% | +5.80 / +5.84% | T1, E2 |
| shaman-elemental bare 60 | 191.94 → 203.45 | +5.81 / +6.09% | +6.00 / +6.12% | T1, E2 |
| shaman-elemental bare 20 / 30 / 40 | 22.16 / 48.76 / 94.84 → 22.22 / 48.78 / 94.93 (h s1) | - | +0.3 / +0.0 / +0.1% | E2 |
| mage-arcane, mage-fire, mage-frost raid 60 | 619.8, 673.0, 632.0 → same | ±0.00% | ±0.00% | - |

**Outside the group:** shaman-enhancement raid 60 moves −0.14 to +0.09% from E2 (within error; its Searing Totem). TestEnhancement's golden moves.

## Files to regenerate

- **Fork goldens, already regenerated for E2 only:** `sim/shaman/elemental/TestElemental.results` and `sim/shaman/enhancement/TestEnhancement.results` (E 5ca93cfa3). The core lane's regen-phase commit also moves both; regenerate after the merge.
- **apl-sync:** `sim/request/apl/shaman-elemental.apl.json` (E1) and `druid-balance.apl.json` (B2).
- **sim/request TestRotationLadder goldens:** shaman-elemental (T1, E1, E2) and druid-balance (B1, B2, B3); shaman-enhancement if it casts Searing Totem.
- **addon-builds and Data.lua:** the Elemental and Balance leveling builds (T1, B3).
- **BiS files:** shaman-elemental bare 50 and 60 talents; druid-balance bare 30 to 60 talents.
- **dps-per-point:** Elemental haste (E1) and Balance (B1).
- **Copies of the old guide build:** `apps/shared/fixtures/planner-rules.json` and `docs/contracts/fixtures/app/cdn.guides.shaman-elemental.json` carry the 70707 Elemental guide code. They are snapshot fixtures; refresh them with the app data.

## Measure at refresh

- Elemental bare 50 and 60 picks with Tidal Focus (5% hit moves the hit value).
- Balance bare 30 to 60 picks on the new leveling build (Naturalist 5, Nature's Grace).
- Elemental and Balance haste scale factors after E1 and B1.
- The Skyborne race for Balance and Elemental (X2).

## For other lanes

- **Guides:** the Balance guide's "a resisted Wrath is two lost Eclipse charges" (Stat priority) is unproven (B4). The opener and build prose was edited here (S b7bc3b144); please read it in your voice. The two APL notes you raised are right (O4).
- **Core:** the Elemental haste dip is E1 (an APL timing tie); fixed.
- **Hybrids:** Enhancement's Searing Totem now attacks every 2.437 s (E2).
- **Ranks / ranker:** X2.
- **Owner:** B3 replaces Improved Entangling Roots and Subtlety in Balance's leveling build under your two precedent rulings; overrule if Roots kiting should stay. Natural Grace → Restorative Totems for Elemental (+0.65%) is held under the 1% bar.

## Still open

- E4 Elemental Mastery's charge; B4 Eclipse timing; B6 Improved Moonfire and Balance cost cuts; M5 cost-mod stacking; M6 Master of Elements on free crits; M7 Ice Lance coefficient, Fire Power on DoTs, Ice Lance's category.

## Sources

- **1.60.1.70753 client tables:**
  - SpellMisc (Speed and CastingTimeIndex for 116, 10149, 25306, 401502, 1237312, 1237313, 3606 to 10436);
  - SpellCastTimes 24 (2200);
  - SpellPower (10151 395, 25306 410, 18809 440);
  - SkillLineAbility 5436, 5468, 13245, 13246;
  - talents/shaman.json, druid.json, mage.json;
  - data/curated/book-ranks.json.
- **foreverlogs.gg beta timelines,** read 2026-11-03: Arcane Blast encounters in reports 2721, 2726, 2764, 2807 and 2809; Wrath druid encounters from the round-12 cache list; 60 Searing Totem encounters (`r13/casters/searcands.txt`).
