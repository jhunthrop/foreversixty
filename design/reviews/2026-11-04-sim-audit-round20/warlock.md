# Sim-accuracy round 20: warlock (2026-11-04)

Specs: Affliction, Demonology, Destruction. Client 1.60.1.70755 with the late hotfix re-layer. Site `wt/r20-warlock` from int/r20 (dcb17fa6b); fork `lane/r20-warlock` from int/r20 (6c96eb479). Measured with spec-breakdown on `$R/bis-frozen`; raid runs use each guide's raid build (`$R/rt`), bare runs the frozen band entry's talents. The round-19 method rules were applied: rank rows level by level on any gain beyond error, a client coefficient of 0 means check the logs, and every lower rank confirmed cast under its own id.

**Totals: 2 CONFIRMED-BUGs, both `old`, both fixed: E 5f37ab425 (the level-50 Succubus swung the level-40 weapon) and S bf09a4c6e (Destruction casts Shadow Bolt rank 4 at 30 and 31).** Every talent the three builds take was re-read against the 70755 SpellEffect masks and the TraitDefinitionEffectPoints curves and matches the engine. The log items stay open: no live log existed before launch (last looked at the time in "Logs" below).

## Seeds

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **Eureka! on Incinerate, demon crit, Soul Siphon, Wrack and Doom on a new log.** No live log: see "Logs". The engine still follows the client row 1259821 (Incinerate in neither mask), the measured demon crit rates (0.7% melee, 6.4% Lash, 2.2% Firebolt) and Soul Siphon's 3-effect cap. | `sim/core/eureka_auto_gen.go:825`; `pet.go:253`; "Logs" | OPEN | - | - |
| S2 | **Affliction Malevolence string at refresh.** The frozen bare 40/50/60 entries still carry Nightfall (`25522300130201351`); carried to "Measure at refresh". | `$R/bis-frozen/warlock-affliction.json` | carried | - | - |
| S3 | **The mage lane's handler-order trap** (an OnCastComplete handler reading state another handler changes): the only warlock OnCastComplete handlers are Shadow Trance (reads `spell.CurCast.CastTime`, which the aura's own OnGain set before the cast started, so activation order cannot bite) and Fel Domination (reads only the spell identity). Shadow and Flame, Decimation and Improved Shadow Bolt trigger from OnSpellHitDealt or inside ApplyEffects. Nothing to fix. | `talents.go:330`, `fel_domination.go:33`, `soul_fire.go:571`, `conflagrate.go:185` | Aligned | - | - |
| S4 | **Low-rank proc chance** (Blizzard, 2 November: only Blizzard keeps the penalty): the warlock engine has no proc penalty and no warlock proc keys to a cast rank (Nightfall rolls per tick, Decimation and Shadow and Flame on any rank), so nothing to add. | mage.md S6; `talents.go` | Aligned | - | - |
| S5 | **`sim/leveling/dead_points_test.go`**: no warlock case to add. Every point the six builds spend acts on something the spec does in play: Grim Reach and Destructive Reach are range, Molten Skin and Demonic Aegis are damage taken, Soul Harvest and Improved Drains are Drain Soul and Drain Life, Improved Health Funnel and Improved Voidwalker are the leveling demon's upkeep. None is an Elemental Devastation (a stat the spec can never use). | `tools/dec.py` decodes; talents/warlock.json | Aligned | - | - |

## Affliction

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | Raid 60 reproduces `before.tsv` to the decimal at round start: 753.1 / 753.6 alliance, 750.6 / 750.8 horde. Fight shape (horde, seed 1): Wrack 47.5% (22.8 casts, 24.9% crit), Corruption 12.9%, Succubus melee 10.9%, Doom 10.4% (2.41 casts), Immolate 9.5%, Siphon Life 4.5%, Agony 3.2%, Lash 1.1%; 1.95 Life Taps; 0 s out of mana. | `out/aff-h-sanity.txt`, `out/after60.tsv` | Aligned | - | - |
| A2 | **Round 18 A1's "5.4 Wrack ticks a cast because a missed cast ends the channel" is right for the wrong reason, and there is no lost tick.** A one-iteration combat log (`out/logfull-corrbolt.txt`) shows every Corruption tick landing at 3, 6, 9, 12, 15 and 18 s; the raw counters over 3,000 iterations (`TestScratchMetrics`, 10.67 casts, 166,087 ticks) give 5.19 a cast only because 8.4% of casts miss (the character has 8.6% spell hit against the 17% boss miss, Shadow Bolt and Lash miss 8.3 to 8.5%) and the last cast of the fight is cut by its end. Wrack's 32.5 casts in a Wrack-only run are 30 channels plus 2.5 misses recast after the global. The breakdown's Miss column on a DoT is diluted by its ticks (Immolate "1.6%" is 1 miss in 12 casts), which is how round 18 read a 1.6% miss as the hit rate. Nothing to fix in the engine; the column is noted for platform below. | `out/log-corrbolt.txt`, `out/logfull-corrbolt.txt`, scratch metrics dump; `core/spell_result.go:218-237` | Aligned (round 18's reading corrected) | - | - |
| A3 | **Improved Drains and Wrack (round 2 A12) closes.** The layered 70755 SpellEffect for 403511 (rec 1050496) now carries mask (0x4008, 0, 0, 0): Drain Life and Drain Soul only, Wrack's bit (word 1 0x40000) is gone and effect 1 (rec 1368037) is Drain Mana. The engine's list matches the text and the mask. | `out/talent-masks.txt` lines 1-5; `talents.go:1014` | Aligned (closed) | - | - |
| A4 | **Grim Reach 2 against Nightfall 2 with a Shadow Trance bolt.** The raid build's two Grim Reach points do nothing in the rotation (range). The one Affliction talent that could spend them for damage is Nightfall (tier-legal: `25500020135201351`) with a "Shadow Bolt when Shadow Trance is up" line before Wrack. Raid 60 horde, 10,000 iterations: 748.3 / 748.1 against 750.6 / 750.8 (-0.3 / -0.4%); an instant bolt (1,050) over 1.5 s of Wrack (700) does not pay for its 370 mana in Life Taps. Round 15 A4 found the same at bare 60. Grim Reach stays. | `out/ties.tsv`, `out/aff-trance.json` | Aligned | - | - |
| A5 | Every talent the builds take, re-read against the client: Suppression (aura 55 and 54, 5; threat -20), Improved Corruption (-2000 ms; +10% DoT), Improved Bane of Agony (curve 5/10 on 0x400), Pandemic (misc 15, +33/67/100 on the DoT list), Malevolence (aura 107 misc 7, 1 a rank; mask covers Agony, Doom, Corruption, Death Coil, the drains, Lash of Pain, Shadow Bolt, Shadowburn, Siphon Life, Wrack), Soul Siphon (curve 4/8/12, dummy), Shadow Mastery (misc 0 on Shadow Bolt, Corruption, Death Coil, Shadowburn; misc 22 on Agony, Doom, Corruption, the drains, Siphon Life, Wrack), Improved Life Tap (10/20), Amplify Curse, Siphon Life (aura 53, 41, 0.05, 3 s), Wrack (36 + 2.5 a level, 0.286, 1 s, aura 271 +10 on 0x402), Bane of Doom (1742, coefficient 4.0, 60 s), Bane of Agony (46, 0.133, 2 s, 24 s), Corruption (57, 0.2, 3 s, 18 s), Immolate (146 + 2.3 a level at 0.2; 52 a tick at 0.13), Life Tap (420 + 1 a level), Curse of the Elements rank 4 (-75, +10%, level 50); Unholy Power (aura 107 misc 8, +10 on the pet passive), Demonic Aegis, Bane (-500 ms; -2000 ms Soul Fire), Aftermath (misc 0 +50 on Immolate, the initial hit). The crit-bonus math is `1 + (1.5 - 1) x (1 + bonus)`: Pandemic 3 doubles a DoT crit, as the text says. | `out/talent-masks.txt`, `out/spell-rows.txt`, the curve dump; `core/spell_outcome.go:1083` | Aligned | - | - |
| A6 | Rank rows 20-39, level by level (bare, horde and alliance, floor and +3/+6/+9 as a screen at 3,000 iterations, every near tie re-run at 10,000, seeds 1 and 50001): Corruption, Immolate, Shadow Bolt, Bane of Agony and Siphon Life all lose a rank down at every level tried (Corruption by 1.0 to 4.6%, Immolate 1.4 to 3.6%, Agony 1.0 to 2.5%, Siphon Life 1.7% at 39, Shadow Bolt 0.1 to 1.6%). Every forced rank was cast under its own id. No row. | `out/sweep-low.tsv`, `out/ties.tsv` | Aligned | - | - |

## Demonology

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| D1 | Raid 60 reproduces `before.tsv`: 734.1 / 735.1 alliance, 724.3 / 724.0 horde. Fight shape (horde Troll, seed 1): Shadow Bolt 48.2% (44.9 casts, 18.1% crit), Doom 12.6%, Corruption 12.3%, Succubus melee 11.2%, Immolate 10.8%, Agony 3.1%, Lash 1.7%; prepull Imp at -60 s, sacrifice at -45 s, Succubus at -44 s, Soul Link at -2 s; Demonic Pact up 99.9%; 15.1 Life Taps, 0 s out of mana. | `out/raid60-shapes.txt`, `out/after60.tsv` | Aligned | - | - |
| D2 | **Decimation's damage half (guide round 9: "0 in the client")**: SpellEffect 440870 effect 3 reads 0 because the rank values live in the curve: TraitDefinitionEffectPoints e3 is (1, 3), (2, 6). The engine's 3/6% on Shadow Bolt and Searing Pain below 35% is the client's. Closed. Cooldown -45/-90 (e1), cast time -20/-40 (440873 e0) likewise. | the curve dump; `soul_fire.go:517-536` | Aligned (closed) | - | - |
| D3 | Improved Sayaad is aura 108 misc 8 (all effects' base points) on Lash of Pain's mask, +10/20/30: the engine multiplies the base 50 only, not the spell-power part, which is what misc 8 does. Master Demonologist (dummies 2 a rank), Demonic Knowledge (33/67/100% of level to both), Soul Link (3% both, 30% taken), Fel Vitality (aura 107 misc 8 +15 on the pet passive; aura 178 +15% mana), Demonic Sacrifice (Imp: Shadow +15%; the engine's 30-minute aura against the text's 2 hours never matters in a fight), Master Summoner (-4000 ms, -40% cost) all match. | `out/talent-masks.txt` | Aligned | - | - |
| D4 | Rank rows 20-39: Shadow Bolt, Corruption, Immolate and Agony lose a rank down at every level, except the levels that already carry a row (Shadow Bolt rank 3 at 20-21 and rank 4 at 28-30, `rank-choices.json`), where the forced rank is the row's and the numbers are identical. Soul Fire is a raid-build skip and a bare 48+ cast; round 18 D2 swept it. | `out/sweep-low.tsv`, `out/ties.tsv` | Aligned | - | - |

## Destruction

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | **Shadow Bolt rank 4 at 30 and 31.** <br>**What was wrong:** at 30 the rotation cast rank 5 (1106, learned at 28, 160 mana). On band-30 gear with the leveling build (Suppression comes at 47) the fight runs short of mana: 2.7 s out of mana and 30 Life Taps in 3 minutes. <br>**Evidence** (bare, band 30, 10,000 iterations, seeds 1 / 50001): rank 4 (1088, 110 mana) against rank 5: at 30 orc 100.0 / 100.0 to 100.2 / 100.2, human 98.7 / 98.7 to 99.0 / 99.0; at 31 orc 101.0 / 100.9 to 101.1 / 101.0, human 99.6 / 99.7 to 99.8 / 99.8; at 32 orc +0.1 to +0.2 but human 101.0 / 100.9 to 101.0 / 101.0 (a tie); at 33 rank 5 wins by 0.1 to 0.2 on both. The forced rank was cast under its own id (the breakdown reads "Shadow Bolt (1088)", 27.1 casts, 2,805 mana). <br>**Fix:** row Shadow Bolt rank 4 from 30 to 31, both presets and races (`data/curated/rank-choices.json` and its embedded copy). `TestWarlockDestructionShadowBoltRank4At30And31` was red, then green. destruction.md "Levels 20 to 60" says so (one sentence) and takes guide round 17's wording for the 40 to 42 ranks. wow-player: see "Guide review". | `out/sweep-low.tsv`, `out/ties.tsv`, `out/sb30.txt` | CONFIRMED-BUG | old (rounds 12 to 14 set Destruction's Shadow Bolt row at 20-23 and Demonology's at 28-30; no round judged Destruction at 30-39, round 17 X1 and round 19 R1 started at 40) | S bf09a4c6e |
| X2 | Raid 60 reproduces `before.tsv`: 795.3 / 794.1 alliance, 792.2 / 792.6 horde. Fight shape (horde, seed 1): Incinerate 31.8% (37.4 casts), Immolate 14.6%, Corruption 12.5%, Conflagrate 10.7% (40.0% crit), Doom 10.4%, Succubus melee 9.4%, Shadowburn 6.8%, Agony 2.6%, Lash 0.9%; 14.8 Life Taps; 0.2 s out of mana; prepull Immolate at -1.5 s. | `out/raid60-shapes.txt`, `out/after60.tsv` | Aligned | - | - |
| X3 | Talents against the client: Improved Shadow Bolt (curve 4 a rank, 12 s, `DamageDoneByCasterMultiplier` so only the warlock's own Shadow damage), Bane, Cataclysm (-3/-6/-10 on the Destruction list with Bane of Havoc), Aftermath (+50 misc 0 on Immolate), Ruin (misc 15 +100 on Conflagrate, Immolate, Incinerate, Rain of Fire, Searing Pain, Shadow Bolt, Shadowburn, Soul Fire), Agonizing Flames (crit on Searing Pain; misc 0 +10 on the eight direct spells; misc 22 +10 on Immolate and Rain of Fire), Fire and Brimstone (aura 107 misc 7 +25 on Conflagrate), Shadow and Flame (dummies: 4 a rank for 20 s, 20 a rank keep), Incinerate (217 at 0.714, +25% dummy), Conflagrate (282 at 0.429, category 672 10 s), Shadowburn (266 at 0.429, 15 s). All match `talents.go` and the spell files. | `out/talent-masks.txt`, `out/spell-rows.txt` | Aligned | - | - |
| X4 | Rank rows 20-39 beyond X1: Corruption, Immolate and Agony lose a rank down at every level; Shadowburn rank 1 against 2 at 26 and 29 is an exact tie at 10,000 iterations on both races (no row; rank 2 stays); Shadow Bolt rank 5 against 6 at 36 is within error (110.5 / 110.6 to 110.6 / 110.6 orc, 109.2 / 109.1 to 109.2 / 109.1 human). Conflagrate below 40 could not be judged this way: the band-30 entry's talents (`2053235100000000`, 21 points) stop before Conflagrate at 34, so a `-level 33` to `39` fight never casts it (the forced rank changed nothing, which is how it showed). Its two sub-40 ranks (1293817 at 25, 100 mana; 1293818 at 32, 130 mana) are both registered. | `out/sweep-low.tsv`, `out/ties.tsv`; `$R/bis-frozen` band 30 | Aligned | - | - |

## Cross-cutting (demons, gear, races)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| C1 | **The level-50 Succubus swung the level-40 weapon.** <br>**What was wrong:** the carried level-50 sheet (`succubus.go`) repeated the level-40 weapon, 41-61, under an upstream "Not updated" comment (wowsims/classic and wowsims/sod carry the same placeholder), so a level-50 demon hit no harder than a level-40 one and then jumped to 95-131 at 60; every other sheet value grows (Strength 74 to 107 to 129). The Succubus is the demon every warlock band sims (`request.go:630`), and at bare 50 its melee is about a tenth of the damage. <br>**Fix:** the level-50 weapon is 68-96, halfway between the carried 40 and 60 sheets, the same linear rule `petSheetAtLevel` already applies to every level without a sheet (30 and 20); marked unconfirmed in the code until a level-50 Forever Succubus log gives the real roll. `TestSuccubusWeaponGrowsBetweenCarriedSheets` was red, then green. The Voidwalker and Felhunter carry the same placeholder (`voidwalker.go:72`, `felhunter.go:71`; the Felhunter's 25 and 40 sheets repeat as well); no site band sims them, so they are left as they are and listed under "Still open". | `succubus.go:74` before the fix; raw.githubusercontent.com wowsims/classic and wowsims/sod `sim/warlock/succubus.go` (read 2026-11-04) | CONFIRMED-BUG | old (upstream placeholder; rounds 2 to 19 checked the demon's crit, hit, scaling and buffs at 60 and the level-20/30 interpolation, never the 50 sheet) | E 5f37ab425 |
| C2 | Published BiS effects at raid 60 (`$R/bis-frozen`): Frozen Heart of the Mountain (round 18 C1); Thunderbrew's Boot Flask (744, the Alliance Affliction and Destruction trinket): ItemEffect 97235 casts 12257 on a 5-minute cooldown, a 5-second aura casting 12278 each second, a direct 75 Fire hit that rolls spell hit and crit, which is the Wowhead Forever tooltip ("Deals 75 Fire damage every 1 sec for 5 sec to all enemies in front of you", 5 min); Rune of the Guard Captain, Uther's Strength, Sageblade, Therazane's Touch, the goggles and Gloves of Spell Mastery are static stats. No set bonus is worn. | `common/item_effects/client_procs.go:137-176`; nether.wowhead.com/forever/tooltip/item/744 | Aligned | - | - |
| C3 | Demon inheritance and hit (horde raid 60, the log's pet lines): the Succubus inherits 4.047 hit (8/17 of the warlock's 8.6%), 65.6 spell damage (a tenth of 656) and 111.5 attack power (17%); Lash of Pain's cast tops the hit up to the warlock's (it misses 8.5%, Shadow Bolt 8.3%); melee crit 0.7% is swallowed by the boss suppression, Lash crits 5.9 to 10.2% by Malevolence. Heart of the Lion is the only player buff on the demon (99.9%). | `out/logfull-corrbolt.txt` lines 38-40; breakdowns | Aligned | - | - |
| C4 | Spell hit: the raid-60 BiS gives 8.6% spell hit (Suppression 5% plus items), so every non-binary spell misses 8.4% and the binary ones (Immolate, Siphon Life, Shadowburn) the same, because `SpellChanceToMiss` reads 17% minus hit for both when the target's resistance is 0 under Curse of the Elements. The breakdown's Miss column is per attempt including ticks, so a DoT's cast miss rate reads low (A2). | `core/spell_result.go:225-237` | Aligned | - | - |
| C5 | Races (published): Human (Ruthless 1% crit, The Human Spirit 10% Spirit into Life Tap), Orc (Blood Fury 2 casts a fight, Command on the demon's passive), Troll (Berserking 1.48 casts, Demonology horde) read in the breakdowns as the client rows say; round 19 D2/X3 measured the orders and nothing in this round moves a race. | breakdowns | Aligned | - | - |
| C6 | Pending owner, not reopened: level-60 trainer book ranks (25307, 25309, 25311), Siphon Life tick crit, the Demonology Horde race, Heart of the Lion 280/276, warlock Orc (O-list). | brief | pending owner | - | - |

## Logs

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| L1 | Before launch the newest foreverlogs.gg `reports/public` entry was still 2811 (2026-10-24, Stockade) and `api.foreversixty.gg/v1/reports/recent` held only z2opt2x2gs5u (2026-10-20); looked at 21:28, 21:59 and again after 23:00 UTC (see the last line of this section). Launch characters are level 1, so nothing this round needs (a 40+ Gnome Destruction log for Eureka! on Incinerate, a 40+ Affliction log for Soul Siphon and Wrack, a level-60 log for Doom and demon crit, a level-50 Succubus log for C1) can exist tonight. The foreverlogs API refuses a bare curl ("API access not allowed") and answers the browser headers of `r3/casters/py/fl.sh`. | fl.sh, api.foreversixty.gg | OPEN | - | - |
| L2 | Warcraft Logs is blocked (gcloud login expired). Recorded, not worked around. | brief | recorded | - | - |

Last looked 2026-11-05 01:44 UTC (2 h 44 min after launch): foreverlogs.gg `reports/public` returned an empty list (no public report at all, the pre-launch 2811 is gone too), and `api.foreversixty.gg/v1/reports/recent` still holds only z2opt2x2gs5u (2026-10-20). No live log settles any L1 item; all stay OPEN.

## Before/after

Raid 60 is unchanged to the decimal on all 12 runs (the level-60 sheet and every level-60 spell are untouched): seed 1 / 50001, alliance then horde: Affliction 753.1 / 753.6 and 750.6 / 750.8; Demonology 734.1 / 735.1 and 724.3 / 724.0; Destruction 795.3 / 794.1 and 792.2 / 792.6 (`out/after60.tsv` against `$R/before.tsv`).

Bare 50 (the Succubus weapon, C1), 10,000 iterations, seeds 1 / 50001, `$R/sb-before` against the lane binary:

| spec | faction | before | after | change |
|---|---|---|---|---|
| Affliction | horde | 268.0 / 268.0 | 276.0 / 276.0 | +3.0% / +3.0% |
| Affliction | alliance | 266.5 / 266.5 | 274.2 / 274.1 | +2.9% / +2.9% |
| Demonology | horde | 229.7 / 229.7 | 237.7 / 237.7 | +3.5% / +3.5% |
| Demonology | alliance | 229.9 / 229.6 | 237.9 / 237.6 | +3.5% / +3.5% |
| Destruction | horde | 278.9 / 279.0 | 286.3 / 286.4 | +2.7% / +2.7% |
| Destruction | alliance | 275.8 / 276.0 | 282.8 / 283.0 | +2.5% / +2.5% |

Bare 30 and 31 (the Shadow Bolt row, X1), Destruction, 10,000 iterations, seeds 1 / 50001: 30 orc 100.0 / 100.0 to 100.2 / 100.2 (+0.2%), human 98.7 / 98.7 to 99.0 / 99.0 (+0.3%); 31 orc 101.0 / 100.9 to 101.1 / 101.0 (+0.1%), human 99.6 / 99.7 to 99.8 / 99.8 (+0.1 to +0.2%). Levels 20 to 29 and 32 to 60 are unchanged. Absolute at raid 60, seed 1, horde: Affliction 750.6, Demonology 724.3, Destruction 792.2.

## Files to regenerate

- sim/request ladder goldens (TestRotationLadder): the warlock cells at 50 to 59 (the Succubus weapon) and Destruction's at 30 and 31 (the rank row).
- addon-data.json and Data.lua: the Destruction rotation lines carry the 30 and 31 row.
- The fork's level-aware goldens (`sim/warlock/dps/*.results`) are level-60 presets and do not move; `TestLevelSmoke` has no golden.

## Measure at refresh

- Carried from round 18: confirm the Affliction bare 40, 50 and 60 BiS entries carry the Malevolence string (`25522300132001351`); until then `-preset bare` runs still show Shadow Trance procs.
- Every warlock bare 50 entry re-ranks on the heavier Succubus (+2.5 to +3.5%); the demon's share of the pick's weights moves with it.

## For other lanes

- **Platform (spec-breakdown, a SHOULD, no DPS):** the Miss column of a DoT row divides misses by every attempt including ticks, so Immolate's "1.6%" is 1 miss in 12 casts (8.4%), and the Hits/iter of a DoT include its ticks, so "ticks a cast" has to be worked out from Casts and Hits. A per-cast miss figure (misses over casts) beside the per-attempt one would have saved round 18 A1 and this round's A2. Nothing in the engine's counters is wrong: `Ticks` already includes `ResistedTicks`.
- No shared code was touched; nothing in this lane moves DPS outside warlock. The Succubus sheet (C1) is read only by warlock bands.

## For the guide lane

- Done in this lane: destruction.md "Levels 20 to 60" (the 30-31 Shadow Bolt sentence, and guide round 17's wording for the 40 to 42 ranks so the two landings agree). No other build, APL, card or race line changed.

## Guide text not done

- None.

## Guide review

- wow-player on the rendered /guides/warlock/destruction "Levels 20 to 60" paragraph (dev build of this branch, 1440 and 390): first pass CHANGES ("feed the bigger bolt" did not say mana, and "at once. Once" read clumsily); its wording was taken verbatim ("Shadowburn joins at 30, and on level 30 gear your mana cannot carry the 160-mana rank 5 as well as the 110-mana rank 4, so cast rank 4 at 30 and 31 and go back to rank 5 at 32"). Re-review on fresh captures at both widths: **SHIP**. `_prose.test.ts` passes.

## Still open

- L1: Eureka! on Incinerate, demon crit at 60, Soul Siphon's count, Wrack and Doom at 60, and the level-50 Succubus weapon (C1): each needs the log named in L1. Last looked: see "Logs".
- The Voidwalker and Felhunter level-50 (and Felhunter 40) weapon placeholders (`voidwalker.go:72`, `felhunter.go:71`): no site band sims them; the same interpolation applies if a band ever does.
- Pending owner (C6), not reopened.

## Sources

- Client 1.60.1.70755, layered with the Forever hotfixes (region 91): `SpellEffect`, `SpellClassOptions`, `SpellMisc`, `SpellPower`, `SpellCooldowns`, `SpellLevels`, `SpellCategories` (cached under `r17/warlock/wago`, read with `r17/warlock/tools/mask.py` and `spell.py`); `TraitDefinition`, `TraitDefinitionEffectPoints`, `CurvePoint` from `c755/fetch/data/builds/1.60.1.70755/raw` (the curve dump in this lane); `data/builds/1.60.1.70755/talents/warlock.json`.
- Engine at the lane head: `sim/warlock/*.go`, `sim/core/spell_result.go`, `sim/core/spell_outcome.go`, `sim/core/eureka_auto_gen.go`, `sim/common/item_effects/client_procs.go`.
- wowsims/classic and wowsims/sod `sim/warlock/succubus.go` (raw.githubusercontent.com, read 2026-11-04): both carry "Not updated" at 50.
- Wowhead Forever tooltip, item 744, https://nether.wowhead.com/forever/tooltip/item/744 (2026-11-04).
- foreverlogs.gg `api/reports/public` and api.foreversixty.gg `/v1/reports/recent` (times in "Logs").
- Lane outputs under `$R/warlock/out/` (`sweep-low.tsv`, `ties.tsv`, `sb30.txt`, `pet50.tsv`, `after60.tsv`, `talent-masks.txt`, `spell-rows.txt`, the combat logs).
