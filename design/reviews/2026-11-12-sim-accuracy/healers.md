# Sim accuracy 2026-11-12: Restoration druid and Restoration shaman

Lane brief: `design/reviews/sim-accuracy/brief.md`. Lane `sim-healers`. Site from origin/main 36a21b4ac, fork from master 8dacbe217 (the pin). Client on the site: 1.60.1.70790 with its hotfix push; live client 1.60.1.70888.

**Totals: 6 confirmed bugs, all fixed (2 new data from client 70888 and the launch logs, 1 old engine, 3 old engine gaps on relics); 2 stale card notes fixed; 1 held; 4 open.** Engine pin moved to 5f1b39450 (the lane's four commits rebased on 03b13a6a5). No regens: the relic picks and the Wild Growth question go on the "Measure at refresh" table.

## Method

1. **Blizzard since 6 Nov.** "WoW: Forever Hotfixes" (t/2403040, posts 4 to 6) and "Class Tuning Incoming" (t/2404092), cached by the warlock lane. Healer lines: Nature's Swiftness also makes Wild Growth instant (10 Nov, landed by the hotfixes lane), Bloom now heals its target and Swiftmend's tooltip no longer says it consumes the effect (6 Nov, both already modelled). Nothing on the shaman's heals.
2. **Client and push.** wago.tools CSVs for 1.60.1.70790 and 1.60.1.70888 with the push layered (`useHotfixes=1`): SpellEffect, SpellMisc, SpellCooldowns, SpellAuraOptions, SpellPower, SpellClassOptions, SpellLevels, SpellCategories, SkillLineAbility. Every spell, talent and relic the two cards and builds touch was read against the site's `spellconst`, `trainables`, `talents` and `items` tables (the full table is below). 70790 against 70888 on those rows: two differences, Nature's Swiftness's mask (already a hand row) and Wild Growth's coefficient (new, below).
3. **Engine against the tables**, line by line: `sim/druid/{healing,healing_tables,rejuvenation,regrowth,healing_touch,wild_growth,swiftmend,tranquility,bloom,innervate,talents_restoration}.go`, the Restoration half of `talents.go`, `items.go`; `sim/shaman/{healing_spell,healing_wave,lesser_healing_wave,chain_heal,riptide,water_shield,talents_restoration,mana_tide_totem,water_totems,items}.go` and the Restoration half of `talents.go`; `core/buffs.go` for Innervate; `core/raid_damage.go` and `sim/request/healprofile*.go` for the heal profile.
4. **Real fights.** Our API holds two public reports (z2opt2x2gs5u, 20 Oct; gxrgvz4yzsi2, 8 Nov open world) with no druid or shaman healer of use (the 8 Nov druids and shamans are passers-by with a handful of heals each). foreverlogs.gg has 24 public reports (2815 to 2850, 5 to 11 Nov): their participant lists name four Restoration druids with events (Surgeon, 2847 and 2848, 27,148 events; Super, 2823; Premo, 2838; Komi, 2840 and 2845) and one Restoration shaman (Eltato, 2834), all between levels 20 and 31 by the ranks they cast; every event each did was read through the report API (`meta_events?sourceCharacterIds=<id>`, 500 a page, paced). No Resto shaman in any public log casts Chain Heal, Riptide or Earth Shield yet (none is 40), so those rest on the client rows. The tank lane's 935 encounter summaries give each healer's effective healing against the group's damage taken.
5. **Sims** on the Windows box through `tools/remote-test.sh` (`--dir sim`, `--engine-sha` for the lane's fork branch, one job at a time, `nice -n 15`, GOMAXPROCS=6): the fork's druid, shaman and core suites (`-tags with_db`), the conformance goldens, and `spec-breakdown -heal` before and after at raid 60 and bare 60 (both factions, seeds 1 and 50001) and the bare bands 20 to 50 (seed 1), 10,000 iterations. The job lists and the Wild Growth rotation variants are in `design/reviews/2026-11-12-sim-accuracy/healers/`.

## Engine against the client (1.60.1.70790 with its push, and 70888)

| # | spell or talent | engine | client | verdict |
|---|---|---|---|---|
| D1 | Healing Touch ranks 1 to 10 (5185 to 9889) | amounts, per level, variance, costs, cast times, coefficients 0.429 to 1 | identical rows; rank 11 (25297) is the book rank, out under `core.IncludeAQ` | aligned |
| D2 | Regrowth ranks 1 to 9 (8936 to 9858) | heal 0.286, tick 0.071, 7 ticks every 3 s, 21 s | identical; the 436937 to 436946 reissues have no learn row | aligned |
| D3 | Rejuvenation ranks 1 to 10 (774 to 9841) | 4 ticks every 3 s at 0.2 | identical; 417057 to 417068 unlearnable, 25299 the book rank | aligned |
| D4 | Tranquility ranks 1 to 4 (740 to 9863) | party aura, 5 ticks every 2 s, 0.067, 5 min | identical (effect 35) | aligned |
| D5 | **Wild Growth ranks 1 to 3 (408120, 1238214, 1238215): tick coefficient** | 0.033 | 70790: 0.033; **70888: 0.066 on every rank** (SpellEffect 1072065, 1235327, 1235329) | **fixed, new data (W1)** |
| D6 | Wild Growth cast 3 s, 6 s cooldown, 7 ticks a second, 550/755/1050 mana | hand 3 s cast since the 70707 push | 70790 and 70888 carry 3000 ms, 6000 ms, 7000 ms, the costs | aligned |
| D7 | Nature's Swiftness (17116) mask | Wrath, Healing Touch, Regrowth, Wild Growth | 70888 mask_1 0x6000000 (the hotfixes lane's hand row) | aligned |
| D8 | Swiftmend (18562) | 20% of base mana, 15 s, leaves the effect, heals the full duration | cost_pct 20, cd 15000, one heal effect | aligned on the row; **which effect it reads: fixed (S1, logs)** |
| D9 | Bloom (1328621, 1328622) | 100% of Spirit on a natural expiry, crits, no healing power | effect 0 amt 100, effect 1 amt 50 (dispel), coefficient 0 | aligned |
| D10 | Innervate (29166) | 5% of base mana, 6 min, 20 s, +400% Spirit regen, full regen while casting | cost_pct 5, cd 360000, dur 20000, aura 110 amt 400, aura 134 amt 100 | aligned |
| D11 | Gift of Nature, Improved Rejuvenation, Improved Regrowth, Tranquil Spirit, Improved Tranquility, Nature's Splendor, Genesis, Naturalist, Reflection, Living Spirit, Nature's Focus | 2%/rank (not Tranquility), 5%/rank, 10% crit/rank, 2% cost/rank, 60% cd at 2/2, +3/+6 s, 2%/rank periodic, 0.1 s/rank, 17/33/50%, 5/10/15%, 14%/rank | the talent rows (17104, 17111, 17074, 24968, 17123, 1223083, 1223081, 17069, 17063) agree | aligned |
| D12 | Nature's Grace (16886) | 10% haste for 3 s off a heal crit, GCD cut | aura 65 amt 10, dur 3000, aura 108 misc 21 | aligned |
| D13 | Omen of Clarity (16864, Clearcasting 16870) | 4% a cast in the mask, 10 s proc cooldown, 15 s charge | trainable at 20; 208 of 5,082 logged casts in the mask procced (4.1% +/- 0.3%) | aligned (see L5) |
| D14 | Druid Spirit regen | 7.5 + Spirit/10 a second (15 + Spirit/5 a tick) | the Classic formula the healer logs do not contradict | aligned (no power snapshots in the public logs) |
| D15 | **Idol of Synthesis (272429, Pix Xizzix)** | not modelled; the ranker left the druid's idol slot empty at 60 (`no_dps_value`) | "Each of your heal over time effects on the target reduces the remaining cooldown on Swiftmend by 1 sec when you cast it" | **fixed, old (R1)** |
| D16 | Idol of Swiftness (279250) | Swiftmend 3 s shorter (registered) | Enchanting, raids-1 (opens 9 Dec) | aligned; not yet obtainable |
| S1 | Healing Wave ranks 1 to 9 (331 to 10396) | generated rows, 0.857 from rank 4 | identical; rank 10 (25357) is the book rank (held, H1) | aligned |
| S2 | Lesser Healing Wave ranks 1 to 6 (8004 to 10468) | rank 6 is 10468 (820), not 27624 | identical | aligned |
| S3 | Chain Heal ranks 1 to 3 (1064, 10622, 10623) | two jumps at half, three targets | SpellEffect EffectChainTargets 3, EffectChainAmplitude 0.5 (first read of the columns the constants leave out) | aligned |
| S4 | Riptide ranks 1 to 3 (408521, 1239242, 1239243) | 5 ticks every 3 s, 0.1, no direct heal, 12 s, 150/180/230, Chain Heal +25% | identical | aligned |
| S5 | Earth Shield (974) | not learnable | no SkillLineAbility row in 70790 or 70888 (only the rune 408514, AcquireMethod 3); the trainables file lists it as a `class_spell` candidate at 50 | aligned; open until a shaman learns it in a log (O1) |
| S6 | Water Shield (408510, 408511) | 3 globes, 2% of max mana, 3.5 s between globes, hits and heal crits, 10 min | SpellAuraOptions 223904 ProcCharges 3, ProcCategoryRecovery 3500; effect 137 amt 2 | aligned (and L7) |
| S7 | Mana Tide Totem ranks 1 to 4 | 13 s, 4 ticks every 3 s (88/133/197/290), 5 min | identical (1325256, 1325257) | aligned |
| S8 | Healing Stream, Mana Spring totems | 5/6/7/9/11 every 2 s at 0.022; 4/6/8/10 every 2 s; 5 min | identical | aligned |
| S9 | Tidal Focus, Tidal Mastery, Purification, Healing Way, Improved Healing Wave, Restorative Totems, Mindfulness, Healing Focus, Ancestral Healing | 1% cost and hit/rank, 1% crit/rank, 2%/rank, 8/17/25%, 0.1 s/rank, 5% and 10%/rank, 17/33/50%, pushback, armor (unmodelled, reasoned) | the talent rows agree | aligned |
| S10 | Nature's Swiftness (16188) | every Nature spell with a cast time | mask_0 0x9C3: Lightning Bolt, Chain Lightning, Healing Wave, Lesser Healing Wave, Chain Heal | aligned |
| S11 | **Tidal Totem (272431, Pix Xizzix)** | generated mod present, never registered; the ranker left the totem slot empty | "Reduces the Mana cost of your Healing Wave ability by 5%" (1291076) | **fixed, old (R2)** |
| S12 | **Totem of Sustaining (23200, Scholomance)** | not modelled | "Increases the maximum number of charges on your Water Shield by 1" (28856, Forever's redesign) | **fixed, old (R3)** |
| S13 | Totem of Urgency (279249) | Lesser Healing Wave +4% crit (registered) | Enchanting, raids-1 | aligned; not yet obtainable |

## Logs against the engine

| # | finding | evidence | verdict |
|---|---|---|---|
| L1 | **Swiftmend heals for the larger of the two effects' full durations.** With Rejuvenation and Regrowth both up, 15 of 15 decidable non-critical Swiftmends equal the bigger total: 11 the Rejuvenation's four ticks (where Regrowth's seven ticks were smaller, 133 to 196), 4 the Regrowth's seven (273 to 325, where they were bigger). The effect with less time left explains 10 of the 15 and the later-cast effect none. The engine read Regrowth whenever it was up (from 16 beta Swiftmends at levels where Regrowth's total was the larger). | Super (2823, levels 22 to 27, 33 Swiftmends), Premo (2838) | **confirmed, fixed (S1)** |
| L2 | **Rejuvenation ticks 4 times over 12 s, Regrowth 7 times over 21 s, every 3 s.** 1,557 Rejuvenations that ran out: 1,361 with 4 ticks (the rest cut by death or a recast), median 12.00 s; 617 Regrowths: 447 with 7 ticks, median 21.00 s. | all four druids | aligned |
| L3 | **Regrowth heals up front and over time.** 803 non-critical heals within 400 ms of an application (median 134 to 204 by level) against 3,307 ticks (median 19 to 45). | all four druids | aligned |
| L4 | **Bloom fires on a natural expiry, never on a refresh.** 1,864 of 2,174 HoT removals have a Bloom within 150 ms (the rest are deaths and recasts); 2 of 184 refreshes do. Bloom crits (5 to 9%) at 1.5x. | all four druids | aligned |
| L5 | **Omen of Clarity.** 208 Clearcasting gains on 5,082 eligible casts of spells in the mask (no gain in the prior 10 s): 4.1% +/- 0.3%, the engine's 4%. Wrath, outside the mask, 2 of 474. Per spell the rate spreads (Rejuvenation 5.1 +/- 0.5, Healing Touch 4.9 +/- 0.8, Starfire 5.4 +/- 1.2, Moonfire 4.4 +/- 1.4, Regrowth 2.5 +/- 0.4). | all four druids | aligned on the whole; Regrowth's low rate open (O2) |
| L6 | **HoT ticks crit at spell crit for 1.5x.** Rejuvenation crits 4.2 to 5.9%, 1.50x (72/48, 62/41, 57/38); Regrowth ticks 5.1 to 6.6%, 1.5x. | all four druids | aligned |
| L7 | **Water Shield returns 2% of the mana pool a globe, one every 3.5 s or more, off heal crits.** 21 gains of 29 to 30 mana (about 1,480 max mana at 28 to 31); 11 of the 13 heal crits are followed by a gain within 200 ms (the other 2 inside the 3.5 s); the shortest gap between gains is 3.6 s. | Eltato (2834) | aligned |
| L8 | **Direct heal crits are 1.5x.** Healing Wave 187/128, Lesser Healing Wave 422/274; Healing Touch 239/152. | Eltato, Super | aligned |
| L9 | **Healers cover 57 to 79% of their group's damage taken** (effective healing over the group's damage taken, every logged encounter): Surgeon 0.79 over 97 encounters, Super 0.73, Komi 0.72, Eltato 0.64, Premo 0.57; the healer does 85 to 97% of the group's healing. The heal profile's bands 20 and 30 are tuned so a bare healer covers 42 to 90%. | tank lane's 935 encounter summaries | aligned |
| L10 | **The logged shaman downranks Healing Wave to rank 2 at 28 to 31** (207 of 208 landed Healing Waves are rank 2, one rank 5), one rank under the card's rank 3 row for 26 to 36. | Eltato | consistent with the mana-bound rows; a player's choice, no row moves |
| L11 | Chain Heal's jumps, Riptide, Mana Tide's party mana, Earth Shield, Wild Growth, Tranquility, Innervate, Nature's Swiftness: no logged Restoration druid or shaman is 40 or above, and totem ticks are sourced from the totem, not the shaman. | participant lists | held (no log) |

## Confirmed and fixed

| # | bug | origin | fix | evidence |
|---|---|---|---|---|
| W1 | Wild Growth's ticks took 0.033 of healing power; client 70888 carries 0.066 on every rank | new data (client 70888) | `wild_growth.go` hand coefficient with its reason, `healing_client_test.go` and the conformance report take the hotfixed value for the three ids | wago.tools SpellEffect, both builds |
| S1 | Swiftmend read Regrowth whenever it was up; Forever heals for the larger of the two effects | old (beta reading at levels where Regrowth was the larger) | `swiftmend.go` picks the larger full duration | L1 |
| R1 | Idol of Synthesis unmodelled, the druid's only open-content healing idol at 60 | old | `items.go`, `swiftmend.go`: 1 s off the remaining cooldown per own HoT on the target (Rejuvenation, Regrowth, Wild Growth) | client item text |
| R2 | Tidal Totem's generated mod never registered | old | `shaman/items.go` registers it through `relicClassMasks` (Healing Wave's mask) | client spell 1291076 |
| R3 | Totem of Sustaining unmodelled (Forever makes it a fourth Water Shield globe) | old | `shaman/items.go`, `water_shield.go`: the gear's extra globes add to the three | client item text, spell 28856 |
| C1 | The Restoration shaman card's Healing Wave note still said "4 until 41, 5 from 42 to 45, 6 from 46 to 55" after the round 30 rows (rank 6 at 37 to 39, 5 at 40 to 43, 6 from 44); the Restoration druid card's notes and tank and party Rejuvenation notes still said "rank 6 at 49 to 56, rank 8 at 57 to 59" and lacked the 44 to 48 party rows | round 30 (the lane updated the guides, not the card notes) | the three notes follow the rows; the guides already did | `data/curated/rank-choices.json` |

## Before and after

Bare band gear (the committed BiS files), the guide's build and card as published, 10,000 iterations; the pin 8dacbe217 against the lane engine 4027cf9a8 (the same four commits before the rebase onto 03b13a6a5; the shamans lane's commits touch no druid code and no Resto shaman heal). Effective HPS, relative, with the combined error.

**S1, Swiftmend reads the larger effect.** At raid 60 the druid's Rejuvenation (rank 10, Improved Rejuvenation, Nature's Splendor's fifth tick) now and then outheals the Regrowth it sat beside, so the fix is worth +0.05% at raid 60 on both factions and both seeds (errors 0.03 to 0.04, beyond). On every bare band (20, 30, 40, 50 and 60, both factions) Regrowth's seven ticks are the larger whenever both are up and nothing moves (0.00%). W1 moves nothing on the published card (the guide build has no Wild Growth). Casts are unchanged.

| cell | before | after | gain | error |
|---|---|---|---|---|
| raid 60 Alliance s1 | 1019.73 | 1020.23 | +0.05% | 0.03 |
| raid 60 Alliance s50001 | 1019.41 | 1019.90 | +0.05% | 0.04 |
| raid 60 Horde s1 | 1016.07 | 1016.62 | +0.05% | 0.04 |
| raid 60 Horde s50001 | 1015.58 | 1016.13 | +0.05% | 0.04 |
| bare 20 to 60, both factions (10 cells) | as before | as before | 0.00% | 0.02 to 0.08 |

**R1 to R3, the relics**, each worn in the published set's empty relic slot against that set, on the lane engine (both factions, seeds 1 and 50001):

| relic | raid 60 | bare 60 |
|---|---|---|
| Idol of Synthesis (druid) | +0.40 to +0.44% (error 0.03 to 0.04); Swiftmend 14.4 to 15.8 casts a fight | +0.10 to +0.16% (error 0.05 to 0.06) |
| Tidal Totem (shaman) | +2.51 to +2.65% (error 0.05); the mana lasts 228 s instead of 217 and Water Shield is recast 15.5 times instead of 11.7 | +1.80 to +1.87% (error 0.04); the mana lasts 64 s instead of 60 |
| Totem of Sustaining (shaman) | +1.07 to +1.18% (error 0.05); 44 globes a fight instead of 38 | +0.13 to +0.20% (error 0.04) |

The ranker picks between them at the refresh (MR56, MR57); no BiS file moves here.

## Wild Growth under the adoption rule

The question the brief asked: with Nature's Swiftness making Wild Growth instant (10 Nov) and its ticks taking twice the healing power (client 70888), does the guide build take the capstone? Measured on the lane engine at band 60, both factions, seeds 1 and 50001, 10,000 iterations, against the published build and card: the Wild Growth build the guide names (Improved Tranquility's two points to Wild Growth and a second point of Improved Regrowth) with three rotation lines.

| line | raid 60 | bare 60 | why |
|---|---|---|---|
| (a) Wild Growth on your party ahead of Tranquility, on the Tranquility condition (two of the five under 70%) | -8.6 to -8.9% | -5.7 to -6.1% | 10.7 casts a fight at raid 60 and the mana runs out at 233 s instead of 275; Tranquility falls from 3 casts to 1 |
| (b) Wild Growth after the tank lines when any party member is under 80%, on the mana pace | -7.2 to -7.4% | -1.7 to -2.1% | 9.2 casts; the party Rejuvenations it replaces heal more for the mana |
| (c) as (b) at 90% with no pace, above 15% mana | -18.7 to -18.9% | -12.3 to -12.9% | 17 casts a fight, dry at 207 s |

Every variant loses beyond error at every cell, so the build and the card stand (the card never names Wild Growth, and the guide's "worth trying on heavy group damage" paragraph keeps its hedge). The profile's pulses hit three of five members every 5 s for about 7% of a member's health; a 1,050-mana cast whose seven ticks mostly land on full members cannot beat rank 7 Rejuvenation's 195 mana. Nature's Swiftness stays on the Healing Touch save. Rotation variants: `design/reviews/2026-11-12-sim-accuracy/healers/wg-{a,b,c}.rotation.json`.

## Held

- **H1, Healing Wave rank 10 (25357).** The shaman registers it (the generated table carries all ten ranks); the druid cuts its book ranks with `core.MaxTrainerRank`. The card names rank 9 everywhere and the rank rewriting caps at the named rank, so nothing casts it and no number moves. Left for a lane that also re-judges the Resto shaman tests that index `HealingWave[10]`.
- **H2, the heal profile below 60 still lasts 300 s** while the tank lane's dungeon pull is 48 s (owner ruling for the tank fight; the healer bands fight "the same fight at its own level", round 10 ruling). Not touched here.

## Open

- **O1, Earth Shield.** `trainables/shaman.json` lists 974 at level 50 as a `class_spell` (a class-family spell with a level and no learn row); neither client carries a SkillLineAbility row for it. Settled when a level-50 shaman logs it, either way.
- **O2, Omen of Clarity on Regrowth** procs at 2.5 +/- 0.4% a cast against 5% on Rejuvenation and Healing Touch (L5). The aggregate matches the engine's flat 4%; a per-spell model waits on a reason.
- **O3, the druid's Spirit regen formula** and the shaman's 17 + Spirit/5 have no public log with power snapshots yet.
- **O4, Chain Heal's jump targets and falloff in a real fight** (no logged shaman has it).

## Guide and card

Changed in this landing (notes-only card edits need no apl-sync):

- **Restoration druid guide (`druid/restoration.md`):** Swiftmend's talent line (the larger of the two effects), the Wild Growth paragraph (Nature's Swiftness makes it instant; the hedge stays), an idol paragraph under Best in slot (Idol of Synthesis and Idol of Swiftness, the list picks), the Swiftmend rotation line aligned with the card (under about three quarters, mana keeping pace). Card: the top-level note's rank sentences and the tank and party Rejuvenation line notes follow the round 30 rows; the Wild Growth measurement is recorded.
- **Restoration shaman guide (`shaman/restoration.md`):** a totem paragraph under Best in slot (Tidal Totem, Totem of Sustaining, Totem of Urgency, the list picks), "four with Totem of Sustaining" beside both three-globe sentences. Card: the Healing Wave line note follows the round 30 rows.

**wow-player review: SHIP WITH FIXES for both classes, every fix taken.** The reviewer judged the text (no rendered page; the changes are prose and card notes). Druid: a false rank sentence in Rotation ("below the newest rank you know all the way to 59": rank 9 is the newest at 56 and 57, so it now reads "from 49 to 55 and again at 58 and 59"), the guide's Swiftmend line against the card's gate (now the card's), the idol paragraph's fragment opening and missing pointer to the list, the 10 to 12 and 17 to 19 leveling sentences, the stat-priority sentence on healing power, and "rank 6 at 46 to 49" in the card's top-level note (46 to 48). Shaman: "The list above picks" (now "The best in slot list picks"), the two three-globe sentences against the fourth globe, Water Shield's point at 20 missing from the leveling list (added with the point of Ancestral Healing at 24 the order takes), and Totem of Sustaining's Forever text said plainly. Both guides' `updated` stamps read today. The reviewer's blocker, the card's top-level `notes` field reading as a developer log, does not apply: `web/scripts/rotation-apl.mjs` renders the priority lines' notes only, and guards them against engineering terms; the top-level field is the maintainers' record every card carries. The figures the reviewer could not confirm are the client's: Wild Growth's 6 s cooldown (SpellCooldowns 6000), Idol of Swiftness's 3 s (spell 1306488) and Totem of Urgency's 4% (spell 1306448). Elune's Light and Heart of the Wild's Intellect predate the lane and are Forever's rows (talents/druid.json, the race tables).

## Gates

- Fork, on the box (`--engine-sha 5f1b39450`, the lane branch rebased on the shamans lane's pin 03b13a6a5; `go test -tags with_db`): sim/druid, sim/druid/restoration, sim/druid/balance, sim/druid/feral, sim/shaman, sim/shaman/restoration, sim/core, sim/conformance and sim/common/... green, and the conformance goldens regenerated there match the committed ones. The red commit (324087e73) does not build on the pin: the new tests name the constants and relics the fix adds. sim/druid/tank, sim/shaman/elemental and sim/shaman/enhancement cannot run from the box's module cache (their suites write `*.results.tmp` beside the test, which is read-only there); TestFeralTank also prints the same DPS and TPS mismatches at the pin 8dacbe217 as at the lane sha (every subtest, identical figures), so they are not this lane's. Conformance: the druid golden's six Wild Growth rows now read 0.033 to 0.066 under the client-hotfix allowance, regenerated on the box; the shaman golden is unchanged.
- Site, on the box (`--pin-engine` at the new pin): `go test ./request/ ./rankchoices/ ./specs/ ./leveling/` green, and `FOREVER_UPDATE_GOLDEN=1 go test ./request/ -run TestRotationLadder` rewrites no golden (the druid-restoration rungs run 300 iterations, where the Swiftmend rule moves nothing). `npm test -- src/lib/sim src/content/guides src/lib/guides/{rotation-lists,copy,leveling-builds}.test.ts src/lib/bis/{apl-gates,rotation-icons}.test.ts`: 102 files, 2,158 tests green (the pin's version and fixture assertions, the guide prose and section guards, the rotation lists); `npm run lint` 0 errors; `astro check` 0 errors over 2,307 files. No Playwright surface changed, so no e2e spec ran; the card edits are notes only, so apl-check is unaffected.

## Measure at refresh

- **MR56, the druid's idol slot at 60.** druid-restoration band 60, bare and raid: Idol of Synthesis is now a candidate the engine models (+0.4% raid, +0.1 to +0.2% bare over the empty slot); Idol of Swiftness joins when the raid recipes open (its source is `raids-1`).
- **MR57, the shaman's totem slot at 60.** shaman-restoration band 60, bare and raid: Tidal Totem (+2.5% raid, +1.8% bare) against Totem of Sustaining (+1.1% raid, +0.2% bare), both now modelled; Totem of Urgency when the raid recipes open. The weights move with the totem (Healing Wave's cost).
- **MR58, the client move to 70888 (issue #21).** Wild Growth's coefficient (`wild_growth.go`'s hand value, the `healing_client_test.go` and `conformance/damage.go` allowances) reads from the generated tables again once the site's data carries 0.066; drop the three entries then.
- **MR59, Swiftmend's source.** druid-restoration raid 60: +0.05% from reading the larger effect; every bare band unchanged. No pick should move; the raid set_hps re-reads it.

## Sources

- Blizzard, "WoW: Forever Hotfixes" (t/2403040, posts 4 to 6) and "Class Tuning Incoming" (t/2404092).
- wago.tools db2 CSVs for 1.60.1.70790 and 1.60.1.70888 (`useHotfixes=1`), read 2026-11-11.
- foreverlogs.gg public reports 2821, 2823, 2834, 2838, 2840, 2845, 2847, 2848 (report API, read 2026-11-11); the tank lane's encounter summaries for 2815 to 2849.
- The site's `data/builds/1.60.1.70790/{spellconst,trainables,talents,items,bis,loot}`.
